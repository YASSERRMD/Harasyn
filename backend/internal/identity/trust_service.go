package identity

import (
	"time"
)

type TrustService struct {
	repo       Repository
	ctxRepo    UserContextRepository
}

func NewTrustService(repo Repository, ctxRepo UserContextRepository) *TrustService {
	return &TrustService{
		repo:    repo,
		ctxRepo: ctxRepo,
	}
}

type BuildContextRequest struct {
	UserID       string  `json:"user_id"`
	TenantID     string  `json:"tenant_id"`
	MFAVerified  bool    `json:"mfa_verified"`
	IPAddress    string  `json:"ip_address"`
	Country      string  `json:"country,omitempty"`
	Region       string  `json:"region,omitempty"`
	City         string  `json:"city,omitempty"`
	IsVPN        bool    `json:"is_vpn"`
	IsTor        bool    `json:"is_tor"`
	IsProxy      bool    `json:"is_proxy"`
	FailedLogins int     `json:"failed_logins"`
}

func (s *TrustService) BuildContext(req BuildContextRequest) (*UserContext, error) {
	now := time.Now()

	loc := &LocationInfo{
		Country: req.Country,
		Region:  req.Region,
		City:    req.City,
		IsVPN:   req.IsVPN,
		IsTor:   req.IsTor,
		IsProxy: req.IsProxy,
	}

	ipRep := &IPReputation{
		IP:         req.IPAddress,
		Reputation: "unknown",
		IsKnownBad: false,
	}

	hour := now.Hour()
	dayOfWeek := now.Weekday().String()
	isWeekend := now.Weekday() == time.Saturday || now.Weekday() == time.Sunday
	isBusinessHours := !isWeekend && hour >= 9 && hour <= 17

	timeCtx := &TimeContext{
		IsBusinessHours: isBusinessHours,
		DayOfWeek:       dayOfWeek,
		HourOfDay:       hour,
		IsWeekend:       isWeekend,
		Timezone:        now.Location().String(),
	}

	var riskSignals []*RiskSignal

	if req.IsTor {
		riskSignals = append(riskSignals, &RiskSignal{
			SignalType: "tor_network",
			Severity:   "high",
			Score:      30,
			Source:     "location",
		})
	}
	if req.IsVPN {
		riskSignals = append(riskSignals, &RiskSignal{
			SignalType: "vpn_detected",
			Severity:   "medium",
			Score:      10,
			Source:     "location",
		})
	}
	if req.IsProxy {
		riskSignals = append(riskSignals, &RiskSignal{
			SignalType: "proxy_detected",
			Severity:   "medium",
			Score:      10,
			Source:     "location",
		})
	}
	if !isBusinessHours {
		riskSignals = append(riskSignals, &RiskSignal{
			SignalType: "outside_business_hours",
			Severity:   "low",
			Score:      5,
			Source:     "time",
		})
	}
	if req.FailedLogins > 3 {
		riskSignals = append(riskSignals, &RiskSignal{
			SignalType: "multiple_failed_logins",
			Severity:   "high",
			Score:      25,
			Source:     "authentication",
		})
	} else if req.FailedLogins > 0 {
		riskSignals = append(riskSignals, &RiskSignal{
			SignalType: "failed_login",
			Severity:   "low",
			Score:      5,
			Source:     "authentication",
		})
	}

	trustScore := CalculateUserTrustScore(req.MFAVerified, riskSignals)

	ctx := &UserContext{
		UserID:       req.UserID,
		TenantID:     req.TenantID,
		MFAVerified:  req.MFAVerified,
		Location:     loc,
		IPReputation: ipRep,
		TimeContext:  timeCtx,
		RiskSignals:  riskSignals,
		TrustScore:   trustScore,
		CalculatedAt: now,
	}

	return ctx, nil
}

func (s *TrustService) GetLatestContext(userID string) (*UserContext, error) {
	return s.ctxRepo.GetLatest(userID)
}

func CalculateUserTrustScore(mfaVerified bool, signals []*RiskSignal) int {
	score := 60

	if mfaVerified {
		score += 20
	}

	for _, signal := range signals {
		score -= signal.Score
	}

	if score < 0 {
		score = 0
	}
	if score > 100 {
		score = 100
	}

	return score
}

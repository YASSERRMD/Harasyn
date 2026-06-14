package risk

import (
	"time"
)

type RiskSignalType string

const (
	RiskImpossibleTravel    RiskSignalType = "impossible_travel"
	RiskSuspiciousIP        RiskSignalType = "suspicious_ip"
	RiskUnusualTime         RiskSignalType = "unusual_time_access"
	RiskNewDevice           RiskSignalType = "new_device"
	RiskFailedAuthorization RiskSignalType = "failed_authorization"
	RiskHighRiskCountry     RiskSignalType = "high_risk_country"
	RiskAnomalousBehavior   RiskSignalType = "anomalous_behavior"
)

type RiskSignalTaxonomy struct {
	Type        RiskSignalType `json:"type"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	BaseScore   int            `json:"base_score"`
	Severity    string         `json:"severity"`
}

var DefaultTaxonomy = []RiskSignalTaxonomy{
	{RiskImpossibleTravel, "Impossible Travel", "User logged in from two distant locations in impossible time", 40, "high"},
	{RiskSuspiciousIP, "Suspicious IP", "Access from known malicious IP", 35, "high"},
	{RiskUnusualTime, "Unusual Time Access", "Access outside normal business hours", 10, "low"},
	{RiskNewDevice, "New Device", "Access from previously unseen device", 15, "medium"},
	{RiskFailedAuthorization, "Failed Authorization", "Multiple failed authorization attempts", 25, "medium"},
	{RiskHighRiskCountry, "High Risk Country", "Access from high-risk geographic location", 20, "medium"},
	{RiskAnomalousBehavior, "Anomalous Behavior", "Unusual access pattern detected", 30, "high"},
}

type RiskAggregation struct {
	ID          string          `json:"id"`
	TenantID    string          `json:"tenant_id"`
	UserID      string          `json:"user_id"`
	DeviceID    string          `json:"device_id"`
	TotalScore  int             `json:"total_score"`
	Signals     []*RiskSignal   `json:"signals"`
	DecayedAt   time.Time       `json:"decayed_at"`
	CalculatedAt time.Time      `json:"calculated_at"`
}

type RiskEngineService struct {
	repo Repository
}

func NewRiskEngineService(repo Repository) *RiskEngineService {
	return &RiskEngineService{repo: repo}
}

type AddRiskSignalRequest struct {
	TenantID   string          `json:"tenant_id"`
	UserID     string          `json:"user_id"`
	DeviceID   string          `json:"device_id"`
	SignalType RiskSignalType  `json:"signal_type"`
	Source     string          `json:"source"`
	Details    map[string]string `json:"details,omitempty"`
}

func (s *RiskEngineService) AddSignal(req AddRiskSignalRequest) (*RiskSignal, error) {
	taxonomy := s.getTaxonomy(req.SignalType)

	signal := &RiskSignal{
		TenantID:   req.TenantID,
		UserID:     &req.UserID,
		DeviceID:   &req.DeviceID,
		SignalType: string(req.SignalType),
		Severity:   taxonomy.Severity,
		Score:      taxonomy.BaseScore,
		Source:     req.Source,
		CreatedAt:  time.Now(),
	}

	if err := s.repo.Create(signal); err != nil {
		return nil, err
	}

	return signal, nil
}

func (s *RiskEngineService) AggregateRisk(userID, deviceID string) (*RiskAggregation, error) {
	userSignals, _ := s.repo.ListByUser(userID)
	deviceSignals, _ := s.repo.ListByDevice(deviceID)

	allSignals := append(userSignals, deviceSignals...)

	totalScore := 0
	for _, sig := range allSignals {
		totalScore += sig.Score
	}

	if totalScore > 100 {
		totalScore = 100
	}

	aggregation := &RiskAggregation{
		UserID:       userID,
		DeviceID:     deviceID,
		TotalScore:   totalScore,
		Signals:      allSignals,
		CalculatedAt: time.Now(),
	}

	return aggregation, nil
}

func (s *RiskEngineService) ApplyDecay(score int, elapsed time.Duration) int {
	hours := elapsed.Hours()
	decayFactor := 1.0 - (hours / 168.0)
	if decayFactor < 0.1 {
		decayFactor = 0.1
	}
	decayed := int(float64(score) * decayFactor)
	if decayed < 0 {
		decayed = 0
	}
	return decayed
}

func (s *RiskEngineService) ExplainRisk(aggregation *RiskAggregation) []string {
	explanations := make([]string, 0)
	for _, sig := range aggregation.Signals {
		explanations = append(explanations, "Risk signal: "+sig.SignalType+" (score: "+string(rune(sig.Score))+")")
	}
	if aggregation.TotalScore > 70 {
		explanations = append(explanations, "Overall risk is HIGH - consider requiring MFA or blocking access")
	} else if aggregation.TotalScore > 40 {
		explanations = append(explanations, "Overall risk is MEDIUM - additional verification recommended")
	}
	return explanations
}

func (s *RiskEngineService) getTaxonomy(signalType RiskSignalType) RiskSignalTaxonomy {
	for _, t := range DefaultTaxonomy {
		if t.Type == signalType {
			return t
		}
	}
	return RiskSignalTaxonomy{BaseScore: 10, Severity: "low"}
}

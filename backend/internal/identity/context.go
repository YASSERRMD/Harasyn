package identity

import (
	"encoding/json"
	"time"
)

type UserContext struct {
	UserID         string          `json:"user_id"`
	TenantID       string          `json:"tenant_id"`
	MFAVerified    bool            `json:"mfa_verified"`
	Location       *LocationInfo   `json:"location,omitempty"`
	IPReputation   *IPReputation   `json:"ip_reputation,omitempty"`
	TimeContext    *TimeContext     `json:"time_context,omitempty"`
	RiskSignals    []*RiskSignal   `json:"risk_signals,omitempty"`
	TrustScore     int             `json:"trust_score"`
	CalculatedAt   time.Time       `json:"calculated_at"`
}

type LocationInfo struct {
	Country   string  `json:"country"`
	Region    string  `json:"region"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	IsVPN     bool    `json:"is_vpn"`
	IsTor     bool    `json:"is_tor"`
	IsProxy   bool    `json:"is_proxy"`
}

type IPReputation struct {
	IP          string  `json:"ip"`
	Reputation  string  `json:"reputation"`
	ThreatScore float64 `json:"threat_score"`
	IsKnownBad  bool    `json:"is_known_bad"`
	ISP         string  `json:"isp"`
	ASN         string  `json:"asn"`
}

type TimeContext struct {
	IsBusinessHours bool   `json:"is_business_hours"`
	DayOfWeek       string `json:"day_of_week"`
	HourOfDay       int    `json:"hour_of_day"`
	IsWeekend       bool   `json:"is_weekend"`
	Timezone        string `json:"timezone"`
}

type RiskSignal struct {
	SignalType string          `json:"signal_type"`
	Severity   string          `json:"severity"`
	Score      int             `json:"score"`
	Source     string          `json:"source"`
	Details    json.RawMessage `json:"details,omitempty"`
}

type UserContextRepository interface {
	Save(ctx *UserContext) error
	GetLatest(userID string) (*UserContext, error)
}

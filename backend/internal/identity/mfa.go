package identity

import (
	"time"
)

type MFARequirement struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	UserID      string    `json:"user_id"`
	Required    bool      `json:"required"`
	Reason      string    `json:"reason"`
	RiskScore   int       `json:"risk_score"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type MFAChallenge struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Method      string    `json:"method"`
	Status      string    `json:"status"`
	Code        string    `json:"code,omitempty"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
	VerifiedAt  *time.Time `json:"verified_at,omitempty"`
}

type MFAProvider interface {
	SendChallenge(userID, method string) (*MFAChallenge, error)
	VerifyChallenge(challengeID, code string) (bool, error)
	GetAvailableMethods(userID string) ([]string, error)
}

type AdaptiveMFAService struct {
	mfaProvider MFAProvider
}

func NewAdaptiveMFAService(provider MFAProvider) *AdaptiveMFAService {
	return &AdaptiveMFAService{mfaProvider: provider}
}

type MFARequirementRequest struct {
	UserID      string `json:"user_id"`
	TenantID    string `json:"tenant_id"`
	RiskScore   int    `json:"risk_score"`
	IsNewDevice bool   `json:"is_new_device"`
	IsSensitive bool   `json:"is_sensitive_resource"`
}

func (s *AdaptiveMFAService) EvaluateMFARequirement(req MFARequirementRequest) *MFARequirement {
	required := false
	reason := ""

	if req.RiskScore > 60 {
		required = true
		reason = "high risk score"
	} else if req.IsNewDevice {
		required = true
		reason = "new device detected"
	} else if req.IsSensitive {
		required = true
		reason = "sensitive resource access"
	} else if req.RiskScore > 40 {
		required = true
		reason = "elevated risk"
	}

	return &MFARequirement{
		UserID:    req.UserID,
		TenantID:  req.TenantID,
		Required:  required,
		Reason:    reason,
		RiskScore: req.RiskScore,
		ExpiresAt: time.Now().Add(15 * time.Minute),
		CreatedAt: time.Now(),
	}
}

func (s *AdaptiveMFAService) ChallengeUser(userID, method string) (*MFAChallenge, error) {
	return s.mfaProvider.SendChallenge(userID, method)
}

func (s *AdaptiveMFAService) VerifyChallenge(challengeID, code string) (bool, error) {
	return s.mfaProvider.VerifyChallenge(challengeID, code)
}

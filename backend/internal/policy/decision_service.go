package policy

import (
	"fmt"
)

type DecisionService struct {
	repo       Repository
	condRepo   ConditionRepository
	evaluator  *Evaluator
}

func NewDecisionService(repo Repository, condRepo ConditionRepository, evaluator *Evaluator) *DecisionService {
	return &DecisionService{
		repo:      repo,
		condRepo:  condRepo,
		evaluator: evaluator,
	}
}

type AccessDecisionRequest struct {
	TenantID           string `json:"tenant_id"`
	UserID             string `json:"user_id"`
	DeviceID           string `json:"device_id"`
	ResourceID         string `json:"resource_id"`
	UserTrustScore     int    `json:"user_trust_score"`
	DeviceTrustScore   int    `json:"device_trust_score"`
	ResourceSensitivity string `json:"resource_sensitivity"`
	Location           string `json:"location"`
	IsBusinessHours    bool   `json:"is_business_hours"`
	RiskScore          int    `json:"risk_score"`
	MFAVerified        bool   `json:"mfa_verified"`
	SessionActive      bool   `json:"session_active"`
}

func (s *DecisionService) Evaluate(req AccessDecisionRequest) (*EvaluationResult, error) {
	policies, err := s.repo.ListByTenant(req.TenantID)
	if err != nil {
		return nil, fmt.Errorf("failed to list policies: %w", err)
	}

	ctx := &EvaluationContext{
		UserTrustScore:     req.UserTrustScore,
		DeviceTrustScore:   req.DeviceTrustScore,
		ResourceSensitivity: req.ResourceSensitivity,
		Location:           req.Location,
		IsBusinessHours:    req.IsBusinessHours,
		RiskScore:          req.RiskScore,
		MFAVerified:        req.MFAVerified,
		SessionActive:      req.SessionActive,
	}

	for _, policy := range policies {
		if !policy.Enabled {
			continue
		}

		conditions, err := s.condRepo.ListByPolicy(policy.ID)
		if err != nil {
			continue
		}

		result, err := s.evaluator.Evaluate(policy, conditions, ctx)
		if err != nil {
			continue
		}

		if len(conditions) > 0 {
			return result, nil
		}
	}

	return &EvaluationResult{
		Allowed: false,
		Reasons: []string{"no matching policy found"},
	}, nil
}

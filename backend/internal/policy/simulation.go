package policy

import (
	"time"
)

type SimulationRequest struct {
	ID                string            `json:"id"`
	TenantID          string            `json:"tenant_id"`
	PolicyID          string            `json:"policy_id,omitempty"`
	UserTrustScore    int               `json:"user_trust_score"`
	DeviceTrustScore  int               `json:"device_trust_score"`
	ResourceSensitivity string          `json:"resource_sensitivity"`
	Location          string            `json:"location"`
	IsBusinessHours   bool              `json:"is_business_hours"`
	RiskScore         int               `json:"risk_score"`
	MFAVerified       bool              `json:"mfa_verified"`
	SessionActive     bool              `json:"session_active"`
	ContextOverrides  map[string]string `json:"context_overrides,omitempty"`
	CreatedAt         time.Time         `json:"created_at"`
}

type SimulationResult struct {
	RequestID    string            `json:"request_id"`
	Allowed      bool              `json:"allowed"`
	PolicyID     string            `json:"policy_id,omitempty"`
	Reasons      []string          `json:"reasons"`
	Conditions   []ConditionResult `json:"conditions"`
	Explanation  string            `json:"explanation"`
	DryRun       bool              `json:"dry_run"`
	CreatedAt    time.Time         `json:"created_at"`
}

type PolicySimulationService struct {
	decisionService *DecisionService
	evaluator       *Evaluator
}

func NewPolicySimulationService(ds *DecisionService, ev *Evaluator) *PolicySimulationService {
	return &PolicySimulationService{
		decisionService: ds,
		evaluator:       ev,
	}
}

type SimulateAccessRequest struct {
	TenantID            string `json:"tenant_id"`
	UserTrustScore      int    `json:"user_trust_score"`
	DeviceTrustScore    int    `json:"device_trust_score"`
	ResourceSensitivity string `json:"resource_sensitivity"`
	Location            string `json:"location"`
	IsBusinessHours     bool   `json:"is_business_hours"`
	RiskScore           int    `json:"risk_score"`
	MFAVerified         bool   `json:"mfa_verified"`
	SessionActive       bool   `json:"session_active"`
}

func (s *PolicySimulationService) Simulate(req SimulateAccessRequest) (*SimulationResult, error) {
	result, err := s.decisionService.Evaluate(AccessDecisionRequest{
		TenantID:            req.TenantID,
		UserTrustScore:      req.UserTrustScore,
		DeviceTrustScore:    req.DeviceTrustScore,
		ResourceSensitivity: req.ResourceSensitivity,
		Location:            req.Location,
		IsBusinessHours:     req.IsBusinessHours,
		RiskScore:           req.RiskScore,
		MFAVerified:         req.MFAVerified,
		SessionActive:       req.SessionActive,
	})
	if err != nil {
		return nil, err
	}

	explanation := s.buildExplanation(result)

	return &SimulationResult{
		Allowed:     result.Allowed,
		PolicyID:    result.PolicyID,
		Reasons:     result.Reasons,
		Conditions:  result.Conditions,
		Explanation: explanation,
		DryRun:      true,
		CreatedAt:   time.Now(),
	}, nil
}

func (s *PolicySimulationService) buildExplanation(result *EvaluationResult) string {
	if result.Allowed {
		return "Access would be GRANTED based on current policy evaluation"
	}
	return "Access would be DENIED based on current policy evaluation"
}

func (s *PolicySimulationService) WhatIf(req SimulateAccessRequest, overrides map[string]interface{}) (*SimulationResult, error) {
	return s.Simulate(req)
}

package http

import (
	"encoding/json"
	"net/http"

	"github.com/YASSERRMD/harasyn/backend/internal/policy"
)

type PolicyHandler struct {
	decisionService *policy.DecisionService
}

func NewPolicyHandler(decisionService *policy.DecisionService) *PolicyHandler {
	return &PolicyHandler{
		decisionService: decisionService,
	}
}

type AccessDecisionRequest struct {
	TenantID            string `json:"tenant_id"`
	UserID              string `json:"user_id"`
	DeviceID            string `json:"device_id"`
	ResourceID          string `json:"resource_id"`
	UserTrustScore      int    `json:"user_trust_score"`
	DeviceTrustScore    int    `json:"device_trust_score"`
	ResourceSensitivity string `json:"resource_sensitivity"`
	Location            string `json:"location"`
	IsBusinessHours     bool   `json:"is_business_hours"`
	RiskScore           int    `json:"risk_score"`
	MFAVerified         bool   `json:"mfa_verified"`
	SessionActive       bool   `json:"session_active"`
}

func (h *PolicyHandler) EvaluateAccess(w http.ResponseWriter, r *http.Request) {
	var req AccessDecisionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.decisionService.Evaluate(policy.AccessDecisionRequest{
		TenantID:            req.TenantID,
		UserID:              req.UserID,
		DeviceID:            req.DeviceID,
		ResourceID:          req.ResourceID,
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

package http

import (
	"encoding/json"
	"net/http"

	"github.com/YASSERRMD/harasyn/backend/internal/identity"
)

type UserHandler struct {
	trustService *identity.TrustService
}

func NewUserHandler(trustService *identity.TrustService) *UserHandler {
	return &UserHandler{
		trustService: trustService,
	}
}

type BuildContextRequest struct {
	UserID       string `json:"user_id"`
	TenantID     string `json:"tenant_id"`
	MFAVerified  bool   `json:"mfa_verified"`
	IPAddress    string `json:"ip_address"`
	Country      string `json:"country,omitempty"`
	Region       string `json:"region,omitempty"`
	City         string `json:"city,omitempty"`
	IsVPN        bool   `json:"is_vpn"`
	IsTor        bool   `json:"is_tor"`
	IsProxy      bool   `json:"is_proxy"`
	FailedLogins int    `json:"failed_logins"`
}

func (h *UserHandler) BuildContext(w http.ResponseWriter, r *http.Request) {
	var req BuildContextRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	ctx, err := h.trustService.BuildContext(identity.BuildContextRequest{
		UserID:       req.UserID,
		TenantID:     req.TenantID,
		MFAVerified:  req.MFAVerified,
		IPAddress:    req.IPAddress,
		Country:      req.Country,
		Region:       req.Region,
		City:         req.City,
		IsVPN:        req.IsVPN,
		IsTor:        req.IsTor,
		IsProxy:      req.IsProxy,
		FailedLogins: req.FailedLogins,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ctx)
}

func (h *UserHandler) GetContext(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id required", http.StatusBadRequest)
		return
	}

	ctx, err := h.trustService.GetLatestContext(userID)
	if err != nil {
		http.Error(w, "context not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ctx)
}

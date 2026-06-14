package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/YASSERRMD/harasyn/backend/internal/session"
)

type SessionHandler struct {
	service *session.Service
}

func NewSessionHandler(service *session.Service) *SessionHandler {
	return &SessionHandler{
		service: service,
	}
}

type CreateSessionRequest struct {
	TenantID   string `json:"tenant_id"`
	UserID     string `json:"user_id"`
	DeviceID   string `json:"device_id"`
	ResourceID string `json:"resource_id"`
	PolicyID   string `json:"policy_id,omitempty"`
	Duration   int    `json:"duration_minutes"`
}

func (h *SessionHandler) CreateSession(w http.ResponseWriter, r *http.Request) {
	var req CreateSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	duration := time.Duration(req.Duration) * time.Minute
	if duration == 0 {
		duration = 30 * time.Minute
	}

	sess, err := h.service.CreateSession(session.CreateSessionRequest{
		TenantID:   req.TenantID,
		UserID:     req.UserID,
		DeviceID:   req.DeviceID,
		ResourceID: req.ResourceID,
		PolicyID:   req.PolicyID,
		Duration:   duration,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sess)
}

func (h *SessionHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "session id required", http.StatusBadRequest)
		return
	}

	sess, err := h.service.GetSession(id)
	if err != nil {
		http.Error(w, "session not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sess)
}

func (h *SessionHandler) RevokeSession(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "session id required", http.StatusBadRequest)
		return
	}

	var req struct {
		Reason string `json:"reason"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		req.Reason = "manual revocation"
	}

	if err := h.service.RevokeSession(id, req.Reason); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "revoked"})
}

func (h *SessionHandler) ListActiveSessions(w http.ResponseWriter, r *http.Request) {
	tenantID := r.URL.Query().Get("tenant_id")
	userID := r.URL.Query().Get("user_id")

	var sessions []*session.AccessSession
	var err error

	if tenantID != "" {
		sessions, err = h.service.ListActiveSessionsByTenant(tenantID)
	} else if userID != "" {
		sessions, err = h.service.ListActiveSessionsByUser(userID)
	} else {
		http.Error(w, "tenant_id or user_id required", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

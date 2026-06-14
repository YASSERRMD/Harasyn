package http

import (
	"encoding/json"
	"net/http"

	"github.com/YASSERRMD/harasyn/backend/internal/access"
)

type AccessHandler struct {
	service *access.Service
}

func NewAccessHandler(service *access.Service) *AccessHandler {
	return &AccessHandler{
		service: service,
	}
}

type CreateRequestRequest struct {
	TenantID        string `json:"tenant_id"`
	UserID          string `json:"user_id"`
	DeviceID        string `json:"device_id"`
	ResourceID      string `json:"resource_id"`
	RequestType     string `json:"request_type"`
	Justification   string `json:"justification,omitempty"`
	DurationMinutes int    `json:"duration_minutes"`
}

func (h *AccessHandler) CreateRequest(w http.ResponseWriter, r *http.Request) {
	var req CreateRequestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	accessReq, err := h.service.CreateRequest(access.CreateRequestRequest{
		TenantID:        req.TenantID,
		UserID:          req.UserID,
		DeviceID:        req.DeviceID,
		ResourceID:      req.ResourceID,
		RequestType:     req.RequestType,
		Justification:   req.Justification,
		DurationMinutes: req.DurationMinutes,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(accessReq)
}

func (h *AccessHandler) GetRequest(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "request id required", http.StatusBadRequest)
		return
	}

	accessReq, err := h.service.GetRequest(id)
	if err != nil {
		http.Error(w, "request not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessReq)
}

func (h *AccessHandler) ListPendingRequests(w http.ResponseWriter, r *http.Request) {
	tenantID := r.URL.Query().Get("tenant_id")
	if tenantID == "" {
		http.Error(w, "tenant_id required", http.StatusBadRequest)
		return
	}

	requests, err := h.service.ListPendingRequests(tenantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}

type ApproveRequestRequest struct {
	RequestID  string `json:"request_id"`
	ReviewerID string `json:"reviewer_id"`
	Reason     string `json:"reason,omitempty"`
}

func (h *AccessHandler) ApproveRequest(w http.ResponseWriter, r *http.Request) {
	var req ApproveRequestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	decision, err := h.service.ApproveRequest(access.ApproveRequestRequest{
		RequestID:  req.RequestID,
		ReviewerID: req.ReviewerID,
		Reason:     req.Reason,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(decision)
}

type RejectRequestRequest struct {
	RequestID  string `json:"request_id"`
	ReviewerID string `json:"reviewer_id"`
	Reason     string `json:"reason"`
}

func (h *AccessHandler) RejectRequest(w http.ResponseWriter, r *http.Request) {
	var req RejectRequestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	decision, err := h.service.RejectRequest(access.RejectRequestRequest{
		RequestID:  req.RequestID,
		ReviewerID: req.ReviewerID,
		Reason:     req.Reason,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(decision)
}

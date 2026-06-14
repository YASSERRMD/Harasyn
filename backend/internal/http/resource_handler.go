package http

import (
	"encoding/json"
	"net/http"

	"github.com/YASSERRMD/harasyn/backend/internal/resource"
)

type ResourceHandler struct {
	service *resource.Service
}

func NewResourceHandler(service *resource.Service) *ResourceHandler {
	return &ResourceHandler{
		service: service,
	}
}

type RegisterResourceRequest struct {
	TenantID     string `json:"tenant_id"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	ResourceType string `json:"resource_type"`
	Endpoint     string `json:"endpoint,omitempty"`
	Port         int    `json:"port,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
	Sensitivity  string `json:"sensitivity"`
}

func (h *ResourceHandler) RegisterResource(w http.ResponseWriter, r *http.Request) {
	var req RegisterResourceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	res, err := h.service.RegisterResource(resource.RegisterResourceRequest{
		TenantID:     req.TenantID,
		Name:         req.Name,
		Description:  req.Description,
		ResourceType: req.ResourceType,
		Endpoint:     req.Endpoint,
		Port:         req.Port,
		Protocol:     req.Protocol,
		Sensitivity:  req.Sensitivity,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *ResourceHandler) GetResource(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "resource id required", http.StatusBadRequest)
		return
	}

	res, err := h.service.GetResource(id)
	if err != nil {
		http.Error(w, "resource not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *ResourceHandler) ListResources(w http.ResponseWriter, r *http.Request) {
	tenantID := r.URL.Query().Get("tenant_id")
	if tenantID == "" {
		http.Error(w, "tenant_id required", http.StatusBadRequest)
		return
	}

	resources, err := h.service.ListResourcesByTenant(tenantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources)
}

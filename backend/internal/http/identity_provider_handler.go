package http

import (
	"encoding/json"
	"net/http"

	"github.com/YASSERRMD/harasyn/backend/internal/identity"
)

type IdentityProviderHandler struct {
	oidcService *identity.OIDCProviderService
}

func NewIdentityProviderHandler(oidcService *identity.OIDCProviderService) *IdentityProviderHandler {
	return &IdentityProviderHandler{oidcService: oidcService}
}

func (h *IdentityProviderHandler) HandleCreateProvider(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TenantID string `json:"tenant_id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Issuer   string `json:"issuer"`
		ClientID string `json:"client_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	provider, err := h.oidcService.CreateProvider(r.Context(), req.TenantID, req.Name, req.Issuer, req.ClientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(provider)
}

func (h *IdentityProviderHandler) HandleGetProvider(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	provider, err := h.oidcService.GetProvider(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(provider)
}

func (h *IdentityProviderHandler) HandleListProviders(w http.ResponseWriter, r *http.Request) {
	tenantID := r.URL.Query().Get("tenant_id")
	if tenantID == "" {
		http.Error(w, "tenant_id is required", http.StatusBadRequest)
		return
	}

	providers, err := h.oidcService.ListProviders(tenantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(providers)
}

func (h *IdentityProviderHandler) HandleInitiateLogin(w http.ResponseWriter, r *http.Request) {
	var req identity.LoginInitiationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.oidcService.InitiateLogin(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *IdentityProviderHandler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	var req identity.CallbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.oidcService.HandleCallback(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *IdentityProviderHandler) HandleUpdateTrustStatus(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID          string `json:"id"`
		TrustStatus string `json:"trust_status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	provider, err := h.oidcService.UpdateTrustStatus(req.ID, req.TrustStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(provider)
}

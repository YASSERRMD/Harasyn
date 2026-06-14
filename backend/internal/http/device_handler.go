package http

import (
	"encoding/json"
	"net/http"

	"github.com/YASSERRMD/harasyn/backend/internal/device"
)

type DeviceHandler struct {
	service        *device.Service
	postureService *device.PostureService
}

func NewDeviceHandler(service *device.Service, postureService *device.PostureService) *DeviceHandler {
	return &DeviceHandler{
		service:        service,
		postureService: postureService,
	}
}

type RegisterDeviceRequest struct {
	TenantID     string `json:"tenant_id"`
	UserID       string `json:"user_id"`
	Name         string `json:"name"`
	Fingerprint  string `json:"fingerprint"`
	OS           string `json:"os"`
	OSVersion    string `json:"os_version,omitempty"`
	DeviceType   string `json:"device_type,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Model        string `json:"model,omitempty"`
}

func (h *DeviceHandler) RegisterDevice(w http.ResponseWriter, r *http.Request) {
	var req RegisterDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	deviceReq := device.RegisterDeviceRequest{
		TenantID:     req.TenantID,
		UserID:       req.UserID,
		Name:         req.Name,
		Fingerprint:  req.Fingerprint,
		OS:           req.OS,
		OSVersion:    req.OSVersion,
		DeviceType:   req.DeviceType,
		Manufacturer: req.Manufacturer,
		Model:        req.Model,
	}

	d, err := h.service.RegisterDevice(deviceReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(d)
}

func (h *DeviceHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "device id required", http.StatusBadRequest)
		return
	}

	d, err := h.service.GetDevice(id)
	if err != nil {
		http.Error(w, "device not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)
}

func (h *DeviceHandler) ListDevices(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	tenantID := r.URL.Query().Get("tenant_id")

	var devices []*device.Device
	var err error

	if userID != "" {
		devices, err = h.service.ListDevicesByUser(userID)
	} else if tenantID != "" {
		devices, err = h.service.ListDevicesByTenant(tenantID)
	} else {
		http.Error(w, "user_id or tenant_id required", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devices)
}

package http

import (
	"net/http"
)

type Router struct {
	mux             *http.ServeMux
	deviceHandler   *DeviceHandler
	userHandler     *UserHandler
	resourceHandler *ResourceHandler
	healthHandler   *HealthHandler
}

func NewRouter(deviceHandler *DeviceHandler, userHandler *UserHandler, resourceHandler *ResourceHandler) *Router {
	return &Router{
		mux:             http.NewServeMux(),
		deviceHandler:   deviceHandler,
		userHandler:     userHandler,
		resourceHandler: resourceHandler,
		healthHandler:   &HealthHandler{},
	}
}

func (r *Router) Setup() {
	r.mux.HandleFunc("/health", r.healthHandler.Handle)
	r.mux.HandleFunc("/api/v1/devices", r.handleDevices)
	r.mux.HandleFunc("/api/v1/devices/", r.handleDeviceByID)
	r.mux.HandleFunc("/api/v1/users/context", r.handleUserContext)
	r.mux.HandleFunc("/api/v1/resources", r.handleResources)
	r.mux.HandleFunc("/api/v1/resources/", r.handleResourceByID)
}

func (r *Router) handleDevices(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		r.deviceHandler.ListDevices(w, req)
	case http.MethodPost:
		r.deviceHandler.RegisterDevice(w, req)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (r *Router) handleDeviceByID(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		r.deviceHandler.GetDevice(w, req)
		return
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (r *Router) handleUserContext(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		r.userHandler.GetContext(w, req)
	case http.MethodPost:
		r.userHandler.BuildContext(w, req)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (r *Router) handleResources(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		r.resourceHandler.ListResources(w, req)
	case http.MethodPost:
		r.resourceHandler.RegisterResource(w, req)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (r *Router) handleResourceByID(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		r.resourceHandler.GetResource(w, req)
		return
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

package http

import (
	"net/http"
)

type Router struct {
	mux             *http.ServeMux
	deviceHandler   *DeviceHandler
	userHandler     *UserHandler
	resourceHandler *ResourceHandler
	policyHandler   *PolicyHandler
	sessionHandler  *SessionHandler
	accessHandler   *AccessHandler
	healthHandler   *HealthHandler
}

func NewRouter(
	deviceHandler *DeviceHandler,
	userHandler *UserHandler,
	resourceHandler *ResourceHandler,
	policyHandler *PolicyHandler,
	sessionHandler *SessionHandler,
	accessHandler *AccessHandler,
) *Router {
	return &Router{
		mux:             http.NewServeMux(),
		deviceHandler:   deviceHandler,
		userHandler:     userHandler,
		resourceHandler: resourceHandler,
		policyHandler:   policyHandler,
		sessionHandler:  sessionHandler,
		accessHandler:   accessHandler,
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
	r.mux.HandleFunc("/api/v1/policies/evaluate", r.handlePolicyEvaluate)
	r.mux.HandleFunc("/api/v1/sessions", r.handleSessions)
	r.mux.HandleFunc("/api/v1/sessions/", r.handleSessionByID)
	r.mux.HandleFunc("/api/v1/access-requests", r.handleAccessRequests)
	r.mux.HandleFunc("/api/v1/access-requests/", r.handleAccessRequestByID)
	r.mux.HandleFunc("/api/v1/access-requests/approve", r.handleApproveRequest)
	r.mux.HandleFunc("/api/v1/access-requests/reject", r.handleRejectRequest)
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

func (r *Router) handlePolicyEvaluate(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		r.policyHandler.EvaluateAccess(w, req)
		return
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (r *Router) handleSessions(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		r.sessionHandler.ListActiveSessions(w, req)
	case http.MethodPost:
		r.sessionHandler.CreateSession(w, req)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (r *Router) handleSessionByID(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		r.sessionHandler.GetSession(w, req)
	case http.MethodDelete:
		r.sessionHandler.RevokeSession(w, req)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (r *Router) handleAccessRequests(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		r.accessHandler.ListPendingRequests(w, req)
	case http.MethodPost:
		r.accessHandler.CreateRequest(w, req)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (r *Router) handleAccessRequestByID(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		r.accessHandler.GetRequest(w, req)
		return
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (r *Router) handleApproveRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		r.accessHandler.ApproveRequest(w, req)
		return
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (r *Router) handleRejectRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		r.accessHandler.RejectRequest(w, req)
		return
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

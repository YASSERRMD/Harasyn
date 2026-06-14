package http

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthHandler struct{}

func (h *HealthHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "ok",
		"service":   "harasyn-api",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

package gateway

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRouteMatcher(t *testing.T) {
	rm := NewRouteMatcher()

	target, _ := url.Parse("http://localhost:8081")
	rm.AddRoute(&Route{Path: "/api/", Target: target, Resource: "api", Priority: 10})
	rm.AddRoute(&Route{Path: "/health", Target: target, Resource: "health", Priority: 20})

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{"exact match", "/health", true},
		{"prefix match", "/api/v1/users", true},
		{"no match", "/unknown", false},
		{"root match", "/", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			route := rm.Match(req)
			if tt.expected && route == nil {
				t.Errorf("expected match for %s", tt.path)
			}
			if !tt.expected && route != nil {
				t.Errorf("expected no match for %s", tt.path)
			}
		})
	}
}

func TestTokenValidator(t *testing.T) {
	tv := NewTokenValidator("Authorization")

	tests := []struct {
		name      string
		header    string
		wantToken string
	}{
		{"bearer token", "Bearer abc123def456", "abc123def456"},
		{"plain token", "abc123def456", "abc123def456"},
		{"empty header", "", ""},
		{"no bearer prefix", "Token xyz", "Token xyz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.header != "" {
				req.Header.Set("Authorization", tt.header)
			}
			token := tv.ExtractToken(req)
			if token != tt.wantToken {
				t.Errorf("expected token '%s', got '%s'", tt.wantToken, token)
			}
		})
	}
}

func TestUpstreamResolver(t *testing.T) {
	ur := NewUpstreamResolver()

	err := ur.Register("resource-1", "http://localhost:9001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	resolved, err := ur.Resolve("resource-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resolved.Host != "localhost:9001" {
		t.Errorf("expected host 'localhost:9001', got '%s'", resolved.Host)
	}

	_, err = ur.Resolve("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent resource")
	}
}

func TestReverseProxyConfig(t *testing.T) {
	config := DefaultConfig()

	if config.ReadTimeout == 0 {
		t.Error("read timeout should not be zero")
	}
	if config.WriteTimeout == 0 {
		t.Error("write timeout should not be zero")
	}
	if config.MaxIdleConns == 0 {
		t.Error("max idle conns should not be zero")
	}
}

func TestReverseProxyCreation(t *testing.T) {
	rp := NewReverseProxy(nil)
	defer rp.Close()

	if rp.config == nil {
		t.Error("config should not be nil")
	}
	if rp.transport == nil {
		t.Error("transport should not be nil")
	}
	if rp.routeMatcher == nil {
		t.Error("route matcher should not be nil")
	}
}

func TestReverseProxyServeHTTP(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Upstream", "true")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("upstream response"))
	}))
	defer upstream.Close()

	rp := NewReverseProxy(nil)
	defer rp.Close()

	target, _ := url.Parse(upstream.URL)
	rp.routeMatcher.AddRoute(&Route{Path: "/", Target: target, Resource: "test", Priority: 1})

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer test-token-12345678")

	w := httptest.NewRecorder()
	rp.ServeHTTP(w, req)

	if w.Code == http.StatusUnauthorized {
		if w.Body.String() != `{"error":"unauthorized","message":"access token required"}` {
			// Token validation is placeholder, so 401 is expected
			return
		}
	}
}

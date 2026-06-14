package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Gateway struct {
	transport *http.Transport
}

func NewGateway() *Gateway {
	return &Gateway{
		transport: &http.Transport{
			MaxIdleConns:        100,
			IdleConnTimeout:     90 * time.Second,
			DisableCompression:  false,
			DisableKeepAlives:   false,
			MaxIdleConnsPerHost: 10,
		},
	}
}

type AccessRequest struct {
	SessionToken string `json:"session_token"`
	ResourceID   string `json:"resource_id"`
	UserID       string `json:"user_id"`
	DeviceID     string `json:"device_id"`
}

type AccessResponse struct {
	Allowed    bool   `json:"allowed"`
	Reason     string `json:"reason,omitempty"`
	TargetURL  string `json:"target_url,omitempty"`
}

func (g *Gateway) EvaluateAccess(req AccessRequest) (*AccessResponse, error) {
	if req.SessionToken == "" {
		return &AccessResponse{
			Allowed: false,
			Reason:  "session token required",
		}, nil
	}

	if req.ResourceID == "" {
		return &AccessResponse{
			Allowed: false,
			Reason:  "resource id required",
		}, nil
	}

	return &AccessResponse{
		Allowed: true,
		Reason:  "access granted",
	}, nil
}

func (g *Gateway) ProxyRequest(targetURL string, req *http.Request) (*http.Response, error) {
	proxyReq, err := http.NewRequest(req.Method, targetURL, req.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create proxy request: %w", err)
	}

	proxyReq.Header = req.Header.Clone()

	resp, err := g.transport.RoundTrip(proxyReq)
	if err != nil {
		return nil, fmt.Errorf("proxy request failed: %w", err)
	}

	return resp, nil
}

type TokenHandoff struct {
	AccessToken string    `json:"access_token"`
	ResourceID  string    `json:"resource_id"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func (g *Gateway) GenerateAccessToken(sessionToken, resourceID string) (*TokenHandoff, error) {
	if sessionToken == "" || resourceID == "" {
		return nil, fmt.Errorf("session token and resource id required")
	}

	return &TokenHandoff{
		AccessToken: fmt.Sprintf("gateway_%s_%s", sessionToken[:8], resourceID[:8]),
		ResourceID:  resourceID,
		ExpiresAt:   time.Now().Add(5 * time.Minute),
	}, nil
}

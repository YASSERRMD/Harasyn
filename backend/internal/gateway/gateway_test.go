package gateway

import (
	"testing"
	"time"
)

func TestGatewayAccessEvaluation(t *testing.T) {
	gw := NewGateway()

	tests := []struct {
		name      string
		req       AccessRequest
		wantAllow bool
	}{
		{
			name: "valid request",
			req: AccessRequest{
				SessionToken: "valid-token-12345678",
				ResourceID:   "resource-12345678",
				UserID:       "user-1",
				DeviceID:     "device-1",
			},
			wantAllow: true,
		},
		{
			name: "missing session token",
			req: AccessRequest{
				SessionToken: "",
				ResourceID:   "resource-12345678",
			},
			wantAllow: false,
		},
		{
			name: "missing resource id",
			req: AccessRequest{
				SessionToken: "valid-token-12345678",
				ResourceID:   "",
			},
			wantAllow: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := gw.EvaluateAccess(tt.req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if resp.Allowed != tt.wantAllow {
				t.Errorf("expected allowed=%v, got %v", tt.wantAllow, resp.Allowed)
			}
		})
	}
}

func TestTokenHandoff(t *testing.T) {
	gw := NewGateway()

	token, err := gw.GenerateAccessToken("session-token-12345678", "resource-id-12345678")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if token.AccessToken == "" {
		t.Error("access token should not be empty")
	}
	if token.ExpiresAt.Before(time.Now()) {
		t.Error("token should not be expired")
	}
}

package gateway

import (
	"context"
	"fmt"
	"strings"
)

type SessionInfo struct {
	SessionID  string
	UserID     string
	DeviceID   string
	ResourceID string
	TenantID   string
	ExpiresAt  int64
}

type TokenValidator struct {
	headerName string
}

func NewTokenValidator(headerName string) *TokenValidator {
	if headerName == "" {
		headerName = "Authorization"
	}
	return &TokenValidator{
		headerName: headerName,
	}
}

func (tv *TokenValidator) ExtractToken(req *http.Request) string {
	authHeader := req.Header.Get(tv.headerName)
	if authHeader == "" {
		return ""
	}

	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}

	return authHeader
}

func (tv *TokenValidator) ValidateToken(ctx context.Context, token string) (*SessionInfo, error) {
	if token == "" {
		return nil, fmt.Errorf("empty token")
	}

	if len(token) < 10 {
		return nil, fmt.Errorf("invalid token format")
	}

	return &SessionInfo{
		SessionID:  "session-" + token[:8],
		UserID:     "user-" + token[:4],
		DeviceID:   "device-" + token[4:8],
		ResourceID: "resource-" + token[:8],
		TenantID:   "default",
	}, nil
}

func (tv *TokenValidator) EnrichContext(ctx context.Context, session *SessionInfo) context.Context {
	ctx = context.WithValue(ctx, "session_id", session.SessionID)
	ctx = context.WithValue(ctx, "user_id", session.UserID)
	ctx = context.WithValue(ctx, "device_id", session.DeviceID)
	ctx = context.WithValue(ctx, "resource_id", session.ResourceID)
	ctx = context.WithValue(ctx, "tenant_id", session.TenantID)
	return ctx
}

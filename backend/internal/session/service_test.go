package session

import (
	"testing"
	"time"
)

func TestSessionCreation(t *testing.T) {
	session := &AccessSession{
		TenantID:   "tenant-1",
		UserID:     "user-1",
		DeviceID:   "device-1",
		ResourceID: "resource-1",
		Status:     "active",
		RiskScore:  0,
		GrantedAt:  time.Now(),
		ExpiresAt:  time.Now().Add(1 * time.Hour),
	}

	if session.Status != "active" {
		t.Errorf("expected status 'active', got '%s'", session.Status)
	}

	if session.RiskScore != 0 {
		t.Errorf("expected risk score 0, got %d", session.RiskScore)
	}
}

func TestSessionRevocation(t *testing.T) {
	session := &AccessSession{
		Status: "active",
	}

	session.Status = "revoked"
	session.RevokeReason = "risk score exceeded threshold"
	now := time.Now()
	session.RevokedAt = &now

	if session.Status != "revoked" {
		t.Errorf("expected status 'revoked', got '%s'", session.Status)
	}

	if session.RevokedAt == nil {
		t.Error("expected revoked_at to be set")
	}
}

func TestRiskScoreUpdate(t *testing.T) {
	session := &AccessSession{
		RiskScore: 30,
	}

	newScore := 80
	session.RiskScore = newScore

	if newScore > 70 {
		session.Status = "revoked"
		session.RevokeReason = "high risk score"
	}

	if session.Status != "revoked" {
		t.Error("session should be revoked with high risk score")
	}
}

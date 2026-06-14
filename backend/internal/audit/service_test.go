package audit

import (
	"testing"
	"time"
)

func TestAuditEventCreation(t *testing.T) {
	event := &AuditEvent{
		TenantID:     "tenant-1",
		EventType:    "access_decision",
		Action:       "access",
		Status:       "allowed",
		CreatedAt:    time.Now(),
	}

	if event.TenantID != "tenant-1" {
		t.Errorf("expected tenant_id 'tenant-1', got '%s'", event.TenantID)
	}

	if event.EventType != "access_decision" {
		t.Errorf("expected event_type 'access_decision', got '%s'", event.EventType)
	}
}

func TestAuditEventTypes(t *testing.T) {
	eventTypes := []string{
		"access_decision",
		"device",
		"session",
		"policy",
		"authentication",
		"authorization",
	}

	for _, et := range eventTypes {
		if et == "" {
			t.Error("event type should not be empty")
		}
	}
}

func TestAuditEventStatuses(t *testing.T) {
	statuses := []string{"success", "failure", "allowed", "denied", "pending"}
	for _, s := range statuses {
		if s == "" {
			t.Error("status should not be empty")
		}
	}
}

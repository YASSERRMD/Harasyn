package access

import (
	"testing"
	"time"
)

func TestAccessRequestCreation(t *testing.T) {
	req := &AccessRequest{
		TenantID:        "tenant-1",
		UserID:          "user-1",
		DeviceID:        "device-1",
		ResourceID:      "resource-1",
		RequestType:     "standard",
		Status:          "pending",
		DurationMinutes: 60,
		RequestedAt:     time.Now(),
	}

	if req.Status != "pending" {
		t.Errorf("expected status 'pending', got '%s'", req.Status)
	}

	if req.DurationMinutes != 60 {
		t.Errorf("expected duration 60, got %d", req.DurationMinutes)
	}
}

func TestEmergencyAccessRequest(t *testing.T) {
	req := &AccessRequest{
		RequestType: "emergency",
	}

	if req.RequestType != "emergency" {
		t.Error("expected emergency request type")
	}
}

func TestApprovalDecision(t *testing.T) {
	decision := &ApprovalDecision{
		RequestID:  "req-1",
		ReviewerID: "reviewer-1",
		Decision:   "approved",
		DecidedAt:  time.Now(),
	}

	if decision.Decision != "approved" {
		t.Errorf("expected decision 'approved', got '%s'", decision.Decision)
	}
}

func TestRejectionDecision(t *testing.T) {
	decision := &ApprovalDecision{
		RequestID:  "req-1",
		ReviewerID: "reviewer-1",
		Decision:   "rejected",
		Reason:     "insufficient justification",
		DecidedAt:  time.Now(),
	}

	if decision.Decision != "rejected" {
		t.Errorf("expected decision 'rejected', got '%s'", decision.Decision)
	}

	if decision.Reason == "" {
		t.Error("expected rejection reason")
	}
}

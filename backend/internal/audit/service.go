package audit

import (
	"encoding/json"
	"fmt"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

type LogEventRequest struct {
	TenantID     string          `json:"tenant_id"`
	EventType    string          `json:"event_type"`
	ActorID      string          `json:"actor_id,omitempty"`
	ActorType    string          `json:"actor_type,omitempty"`
	ResourceType string          `json:"resource_type,omitempty"`
	ResourceID   string          `json:"resource_id,omitempty"`
	Action       string          `json:"action"`
	Status       string          `json:"status"`
	Details      json.RawMessage `json:"details,omitempty"`
	IPAddress    string          `json:"ip_address,omitempty"`
	UserAgent    string          `json:"user_agent,omitempty"`
}

func (s *Service) LogEvent(req LogEventRequest) (*AuditEvent, error) {
	event := &AuditEvent{
		TenantID:     req.TenantID,
		EventType:    req.EventType,
		Action:       req.Action,
		Status:       req.Status,
		Details:      req.Details,
		IPAddress:    req.IPAddress,
		UserAgent:    req.UserAgent,
		CreatedAt:    time.Now(),
	}

	if req.ActorID != "" {
		event.ActorID = &req.ActorID
	}
	if req.ResourceID != "" {
		event.ResourceID = &req.ResourceID
	}
	if req.ActorType != "" {
		event.ActorType = req.ActorType
	}
	if req.ResourceType != "" {
		event.ResourceType = req.ResourceType
	}

	if err := s.repo.Create(event); err != nil {
		return nil, fmt.Errorf("failed to log audit event: %w", err)
	}

	return event, nil
}

func (s *Service) GetEvent(id string) (*AuditEvent, error) {
	return s.repo.GetByID(id)
}

func (s *Service) ListEventsByTenant(tenantID string, limit, offset int) ([]*AuditEvent, error) {
	if limit <= 0 {
		limit = 100
	}
	return s.repo.ListByTenant(tenantID, limit, offset)
}

func (s *Service) ListEventsByResource(resourceType, resourceID string) ([]*AuditEvent, error) {
	return s.repo.ListByResource(resourceType, resourceID)
}

func (s *Service) ListEventsByActor(actorID string) ([]*AuditEvent, error) {
	return s.repo.ListByActor(actorID)
}

func (s *Service) ListEventsByType(eventType string, limit int) ([]*AuditEvent, error) {
	if limit <= 0 {
		limit = 100
	}
	return s.repo.ListByType(eventType, limit)
}

func (s *Service) LogAccessDecision(tenantID, actorID, resourceID, decision, reason string) error {
	details, _ := json.Marshal(map[string]string{
		"decision": decision,
		"reason":   reason,
	})

	_, err := s.LogEvent(LogEventRequest{
		TenantID:     tenantID,
		EventType:    "access_decision",
		ActorID:      actorID,
		ActorType:    "user",
		ResourceType: "resource",
		ResourceID:   resourceID,
		Action:       "access",
		Status:       decision,
		Details:      details,
	})
	return err
}

func (s *Service) LogDeviceRegistration(tenantID, deviceID, action string) error {
	_, err := s.LogEvent(LogEventRequest{
		TenantID:     tenantID,
		EventType:    "device",
		ActorID:      deviceID,
		ActorType:    "device",
		ResourceType: "device",
		ResourceID:   deviceID,
		Action:       action,
		Status:       "success",
	})
	return err
}

func (s *Service) LogSessionEvent(tenantID, sessionID, userID, action string) error {
	_, err := s.LogEvent(LogEventRequest{
		TenantID:     tenantID,
		EventType:    "session",
		ActorID:      userID,
		ActorType:    "user",
		ResourceType: "session",
		ResourceID:   sessionID,
		Action:       action,
		Status:       "success",
	})
	return err
}

func (s *Service) LogPolicyChange(tenantID, policyID, action string) error {
	_, err := s.LogEvent(LogEventRequest{
		TenantID:     tenantID,
		EventType:    "policy",
		ResourceType: "policy",
		ResourceID:   policyID,
		Action:       action,
		Status:       "success",
	})
	return err
}

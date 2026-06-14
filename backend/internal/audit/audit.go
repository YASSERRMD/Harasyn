package audit

import (
	"encoding/json"
	"time"
)

type AuditEvent struct {
	ID           string          `json:"id"`
	TenantID     string          `json:"tenant_id"`
	EventType    string          `json:"event_type"`
	ActorID      *string         `json:"actor_id,omitempty"`
	ActorType    string          `json:"actor_type,omitempty"`
	ResourceType string          `json:"resource_type,omitempty"`
	ResourceID   *string         `json:"resource_id,omitempty"`
	Action       string          `json:"action"`
	Status       string          `json:"status"`
	Details      json.RawMessage `json:"details,omitempty"`
	IPAddress    string          `json:"ip_address,omitempty"`
	UserAgent    string          `json:"user_agent,omitempty"`
	CreatedAt    time.Time       `json:"created_at"`
}

type Repository interface {
	Create(e *AuditEvent) error
	GetByID(id string) (*AuditEvent, error)
	ListByTenant(tenantID string, limit, offset int) ([]*AuditEvent, error)
	ListByResource(resourceType, resourceID string) ([]*AuditEvent, error)
	ListByActor(actorID string) ([]*AuditEvent, error)
	ListByType(eventType string, limit int) ([]*AuditEvent, error)
}

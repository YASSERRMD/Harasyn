package resource

import (
	"encoding/json"
	"time"
)

type Resource struct {
	ID           string          `json:"id"`
	TenantID     string          `json:"tenant_id"`
	Name         string          `json:"name"`
	Description  string          `json:"description,omitempty"`
	ResourceType string          `json:"resource_type"`
	Endpoint     string          `json:"endpoint,omitempty"`
	Port         int             `json:"port,omitempty"`
	Protocol     string          `json:"protocol,omitempty"`
	Sensitivity  string          `json:"sensitivity"`
	Status       string          `json:"status"`
	Metadata     json.RawMessage `json:"metadata,omitempty"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

type Connector struct {
	ID             string          `json:"id"`
	ResourceID     string          `json:"resource_id"`
	ConnectorType  string          `json:"connector_type"`
	TargetHost     string          `json:"target_host"`
	TargetPort     int             `json:"target_port,omitempty"`
	TLSRequired    bool            `json:"tls_required"`
	AuthMethod     string          `json:"auth_method"`
	Config         json.RawMessage `json:"config,omitempty"`
	Status         string          `json:"status"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

type Repository interface {
	Create(r *Resource) error
	GetByID(id string) (*Resource, error)
	Update(r *Resource) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*Resource, error)
}

type ConnectorRepository interface {
	Create(c *Connector) error
	GetByID(id string) (*Connector, error)
	GetByResourceID(resourceID string) (*Connector, error)
	Update(c *Connector) error
	Delete(id string) error
}

package policy

import (
	"encoding/json"
	"time"
)

type AccessPolicy struct {
	ID          string          `json:"id"`
	TenantID    string          `json:"tenant_id"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	PolicyType  string          `json:"policy_type"`
	Priority    int             `json:"priority"`
	Enabled     bool            `json:"enabled"`
	Effect      string          `json:"effect"`
	ResourceID  *string         `json:"resource_id,omitempty"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type PolicyCondition struct {
	ID            string          `json:"id"`
	PolicyID      string          `json:"policy_id"`
	ConditionType string          `json:"condition_type"`
	Operator      string          `json:"operator"`
	Value         string          `json:"value"`
	Metadata      json.RawMessage `json:"metadata,omitempty"`
	CreatedAt     time.Time       `json:"created_at"`
}

type Decision struct {
	Allowed  bool     `json:"allowed"`
	PolicyID string   `json:"policy_id,omitempty"`
	Reasons  []string `json:"reasons"`
}

type Repository interface {
	Create(p *AccessPolicy) error
	GetByID(id string) (*AccessPolicy, error)
	Update(p *AccessPolicy) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*AccessPolicy, error)
	ListByResource(resourceID string) ([]*AccessPolicy, error)
}

type ConditionRepository interface {
	Create(c *PolicyCondition) error
	GetByID(id string) (*PolicyCondition, error)
	ListByPolicy(policyID string) ([]*PolicyCondition, error)
	Delete(id string) error
}

package risk

import (
	"encoding/json"
	"time"
)

type RiskSignal struct {
	ID         string          `json:"id"`
	TenantID   string          `json:"tenant_id"`
	UserID     *string         `json:"user_id,omitempty"`
	DeviceID   *string         `json:"device_id,omitempty"`
	SignalType string          `json:"signal_type"`
	Severity   string          `json:"severity"`
	Score      int             `json:"score"`
	Source     string          `json:"source,omitempty"`
	Details    json.RawMessage `json:"details,omitempty"`
	CreatedAt  time.Time       `json:"created_at"`
}

type TrustScore struct {
	ID           string          `json:"id"`
	TenantID     string          `json:"tenant_id"`
	EntityType   string          `json:"entity_type"`
	EntityID     string          `json:"entity_id"`
	Score        int             `json:"score"`
	Components   json.RawMessage `json:"components,omitempty"`
	CalculatedAt time.Time       `json:"calculated_at"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

type Repository interface {
	Create(r *RiskSignal) error
	GetByID(id string) (*RiskSignal, error)
	ListByUser(userID string) ([]*RiskSignal, error)
	ListByDevice(deviceID string) ([]*RiskSignal, error)
}

type TrustScoreRepository interface {
	Create(t *TrustScore) error
	GetByEntity(entityType, entityID string) (*TrustScore, error)
	Update(t *TrustScore) error
}

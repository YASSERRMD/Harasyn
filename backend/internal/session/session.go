package session

import "time"

type AccessSession struct {
	ID                 string     `json:"id"`
	TenantID           string     `json:"tenant_id"`
	UserID             string     `json:"user_id"`
	DeviceID           string     `json:"device_id"`
	ResourceID         string     `json:"resource_id"`
	PolicyID           *string    `json:"policy_id,omitempty"`
	Token              string     `json:"token"`
	Status             string     `json:"status"`
	RiskScore          int        `json:"risk_score"`
	GrantedAt          time.Time  `json:"granted_at"`
	ExpiresAt          time.Time  `json:"expires_at"`
	LastRevalidatedAt  time.Time  `json:"last_revalidated_at"`
	RevokedAt          *time.Time `json:"revoked_at,omitempty"`
	RevokeReason       string     `json:"revoke_reason,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
}

type Repository interface {
	Create(s *AccessSession) error
	GetByID(id string) (*AccessSession, error)
	GetByToken(token string) (*AccessSession, error)
	Update(s *AccessSession) error
	Revoke(id string, reason string) error
	ListActiveByUser(userID string) ([]*AccessSession, error)
	ListActiveByTenant(tenantID string) ([]*AccessSession, error)
	ListExpired() ([]*AccessSession, error)
}

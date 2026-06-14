package access

import "time"

type AccessRequest struct {
	ID              string     `json:"id"`
	TenantID        string     `json:"tenant_id"`
	UserID          string     `json:"user_id"`
	DeviceID        string     `json:"device_id"`
	ResourceID      string     `json:"resource_id"`
	RequestType     string     `json:"request_type"`
	Justification   string     `json:"justification,omitempty"`
	Status          string     `json:"status"`
	DurationMinutes int        `json:"duration_minutes"`
	RequestedAt     time.Time  `json:"requested_at"`
	ExpiresAt       *time.Time `json:"expires_at,omitempty"`
	ResolvedAt      *time.Time `json:"resolved_at,omitempty"`
	ResolvedBy      *string    `json:"resolved_by,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
}

type ApprovalDecision struct {
	ID          string    `json:"id"`
	RequestID   string    `json:"request_id"`
	ReviewerID  string    `json:"reviewer_id"`
	Decision    string    `json:"decision"`
	Reason      string    `json:"reason,omitempty"`
	DecidedAt   time.Time `json:"decided_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type Repository interface {
	Create(r *AccessRequest) error
	GetByID(id string) (*AccessRequest, error)
	Update(r *AccessRequest) error
	ListByTenant(tenantID string) ([]*AccessRequest, error)
	ListPendingByTenant(tenantID string) ([]*AccessRequest, error)
}

type ApprovalRepository interface {
	Create(a *ApprovalDecision) error
	GetByRequestID(requestID string) ([]*ApprovalDecision, error)
}

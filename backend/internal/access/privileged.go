package access

import (
	"time"
)

type PrivilegedResource struct {
	ID          string    `json:"id"`
	ResourceID  string    `json:"resource_id"`
	TenantID    string    `json:"tenant_id"`
	IsPrivileged bool     `json:"is_privileged"`
	Classification string `json:"classification"`
	MaxSessionDuration int `json:"max_session_duration_minutes"`
	RequiresApproval bool  `json:"requires_approval"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BreakGlassAccess struct {
	ID          string     `json:"id"`
	RequestID   string     `json:"request_id"`
	UserID      string     `json:"user_id"`
	ResourceID  string     `json:"resource_id"`
	Reason      string     `json:"reason"`
	ApprovedBy  string     `json:"approved_by"`
	ExpiresAt   time.Time  `json:"expires_at"`
	UsedAt      *time.Time `json:"used_at,omitempty"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
}

type PrivilegedAccessService struct {
	repo Repository
}

func NewPrivilegedAccessService(repo Repository) *PrivilegedAccessService {
	return &PrivilegedAccessService{repo: repo}
}

type RequestPrivilegedAccessRequest struct {
	UserID     string `json:"user_id"`
	ResourceID string `json:"resource_id"`
	Reason     string `json:"reason"`
	Duration   int    `json:"duration_minutes"`
}

func (s *PrivilegedAccessService) RequestPrivilegedAccess(req RequestPrivilegedAccessRequest) (*AccessRequest, error) {
	accessReq := &AccessRequest{
		UserID:          req.UserID,
		ResourceID:      req.ResourceID,
		RequestType:     "privileged",
		Justification:   req.Reason,
		Status:          "pending",
		DurationMinutes: req.Duration,
		RequestedAt:     time.Now(),
		CreatedAt:       time.Now(),
	}

	if req.Duration > 60 {
		accessReq.DurationMinutes = 60
	}

	if err := s.repo.Create(accessReq); err != nil {
		return nil, err
	}

	return accessReq, nil
}

func (s *PrivilegedAccessService) EnforceSessionExpiry(sessionID string) error {
	session, err := s.repo.GetByID(sessionID)
	if err != nil {
		return err
	}

	if session.Status != "active" {
		return nil
	}

	maxDuration := 60
	duration := time.Since(session.GrantedAt).Minutes()
	if duration > float64(maxDuration) {
		session.Status = "expired"
		session.RevokeReason = "privileged session max duration exceeded"
		now := time.Now()
		session.RevokedAt = &now
		return s.repo.Update(session)
	}

	return nil
}

func (s *PrivilegedAccessService) GetPrivilegedAccessAuditTrail(resourceID string) ([]*AccessRequest, error) {
	return s.repo.ListByTenant("default")
}

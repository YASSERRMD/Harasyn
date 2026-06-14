package session

import (
	"crypto/rand"
	"encoding/hex"
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

type CreateSessionRequest struct {
	TenantID   string `json:"tenant_id"`
	UserID     string `json:"user_id"`
	DeviceID   string `json:"device_id"`
	ResourceID string `json:"resource_id"`
	PolicyID   string `json:"policy_id,omitempty"`
	Duration   time.Duration `json:"duration"`
}

func (s *Service) CreateSession(req CreateSessionRequest) (*AccessSession, error) {
	token, err := generateToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	session := &AccessSession{
		TenantID:          req.TenantID,
		UserID:            req.UserID,
		DeviceID:          req.DeviceID,
		ResourceID:        req.ResourceID,
		Token:             token,
		Status:            "active",
		RiskScore:         0,
		GrantedAt:         time.Now(),
		ExpiresAt:         time.Now().Add(req.Duration),
		LastRevalidatedAt: time.Now(),
	}

	if req.PolicyID != "" {
		session.PolicyID = &req.PolicyID
	}

	if err := s.repo.Create(session); err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return session, nil
}

func (s *Service) GetSession(id string) (*AccessSession, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetSessionByToken(token string) (*AccessSession, error) {
	return s.repo.GetByToken(token)
}

func (s *Service) RevokeSession(id, reason string) error {
	return s.repo.Revoke(id, reason)
}

func (s *Service) ListActiveSessionsByUser(userID string) ([]*AccessSession, error) {
	return s.repo.ListActiveByUser(userID)
}

func (s *Service) ListActiveSessionsByTenant(tenantID string) ([]*AccessSession, error) {
	return s.repo.ListActiveByTenant(tenantID)
}

func (s *Service) RevalidateSession(session *AccessSession) error {
	if session.Status != "active" {
		return fmt.Errorf("session is not active")
	}

	if time.Now().After(session.ExpiresAt) {
		session.Status = "expired"
		return s.repo.Update(session)
	}

	session.LastRevalidatedAt = time.Now()
	return s.repo.Update(session)
}

func (s *Service) UpdateRiskScore(sessionID string, riskScore int) error {
	session, err := s.repo.GetByID(sessionID)
	if err != nil {
		return err
	}

	session.RiskScore = riskScore

	if riskScore > 70 {
		session.Status = "revoked"
		session.RevokeReason = "high risk score detected"
		now := time.Now()
		session.RevokedAt = &now
	}

	return s.repo.Update(session)
}

func (s *Service) GetExpiredSessions() ([]*AccessSession, error) {
	return s.repo.ListExpired()
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

package gateway

import (
	"time"
)

type BrowserIsolationConfig struct {
	ID            string    `json:"id"`
	TenantID      string    `json:"tenant_id"`
	ResourceID    string    `json:"resource_id"`
	IsEnabled     bool      `json:"is_enabled"`
	ClipboardCtrl string    `json:"clipboard_control"`
	DownloadCtrl  string    `json:"download_control"`
	ScreenshotCtrl string   `json:"screenshot_control"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type IsolatedSession struct {
	ID            string     `json:"id"`
	TenantID      string     `json:"tenant_id"`
	UserID        string     `json:"user_id"`
	ResourceID    string     `json:"resource_id"`
	ConfigID      string     `json:"config_id"`
	Status        string     `json:"status"`
	StartedAt     time.Time  `json:"started_at"`
	EndedAt       *time.Time `json:"ended_at,omitempty"`
	URL           string     `json:"url,omitempty"`
}

type BrowserIsolationRepository interface {
	Create(c *BrowserIsolationConfig) error
	GetByID(id string) (*BrowserIsolationConfig, error)
	GetByResourceID(resourceID string) (*BrowserIsolationConfig, error)
	Update(c *BrowserIsolationConfig) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*BrowserIsolationConfig, error)
}

type IsolatedSessionRepository interface {
	Create(s *IsolatedSession) error
	GetByID(id string) (*IsolatedSession, error)
	Update(s *IsolatedSession) error
	ListByTenant(tenantID string) ([]*IsolatedSession, error)
	ListByUser(userID string) ([]*IsolatedSession, error)
}

type BrowserIsolationService struct {
	configRepo BrowserIsolationRepository
	sessionRepo IsolatedSessionRepository
}

func NewBrowserIsolationService(cr BrowserIsolationRepository, sr IsolatedSessionRepository) *BrowserIsolationService {
	return &BrowserIsolationService{
		configRepo:  cr,
		sessionRepo: sr,
	}
}

func (s *BrowserIsolationService) EnableIsolation(tenantID, resourceID string) (*BrowserIsolationConfig, error) {
	config := &BrowserIsolationConfig{
		TenantID:       tenantID,
		ResourceID:     resourceID,
		IsEnabled:      true,
		ClipboardCtrl:  "read-only",
		DownloadCtrl:   "block",
		ScreenshotCtrl: "block",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.configRepo.Create(config); err != nil {
		return nil, err
	}

	return config, nil
}

func (s *BrowserIsolationService) GetConfig(id string) (*BrowserIsolationConfig, error) {
	return s.configRepo.GetByID(id)
}

func (s *BrowserIsolationService) GetConfigByResource(resourceID string) (*BrowserIsolationConfig, error) {
	return s.configRepo.GetByResourceID(resourceID)
}

func (s *BrowserIsolationService) UpdateClipboardControl(id, control string) error {
	config, err := s.configRepo.GetByID(id)
	if err != nil {
		return err
	}

	config.ClipboardCtrl = control
	config.UpdatedAt = time.Now()

	return s.configRepo.Update(config)
}

func (s *BrowserIsolationService) UpdateDownloadControl(id, control string) error {
	config, err := s.configRepo.GetByID(id)
	if err != nil {
		return err
	}

	config.DownloadCtrl = control
	config.UpdatedAt = time.Now()

	return s.configRepo.Update(config)
}

func (s *BrowserIsolationService) UpdateScreenshotControl(id, control string) error {
	config, err := s.configRepo.GetByID(id)
	if err != nil {
		return err
	}

	config.ScreenshotCtrl = control
	config.UpdatedAt = time.Now()

	return s.configRepo.Update(config)
}

func (s *BrowserIsolationService) StartIsolatedSession(tenantID, userID, resourceID string) (*IsolatedSession, error) {
	config, err := s.configRepo.GetByResourceID(resourceID)
	if err != nil {
		return nil, err
	}

	session := &IsolatedSession{
		TenantID:   tenantID,
		UserID:     userID,
		ResourceID: resourceID,
		ConfigID:   config.ID,
		Status:     "active",
		StartedAt:  time.Now(),
	}

	if err := s.sessionRepo.Create(session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *BrowserIsolationService) EndIsolatedSession(sessionID string) error {
	session, err := s.sessionRepo.GetByID(sessionID)
	if err != nil {
		return err
	}

	session.Status = "ended"
	now := time.Now()
	session.EndedAt = &now

	return s.sessionRepo.Update(session)
}

func (s *BrowserIsolationService) GetSession(id string) (*IsolatedSession, error) {
	return s.sessionRepo.GetByID(id)
}

func (s *BrowserIsolationService) ListSessions(tenantID string) ([]*IsolatedSession, error) {
	return s.sessionRepo.ListByTenant(tenantID)
}

func (s *BrowserIsolationService) DisableIsolation(id string) error {
	config, err := s.configRepo.GetByID(id)
	if err != nil {
		return err
	}

	config.IsEnabled = false
	config.UpdatedAt = time.Now()

	return s.configRepo.Update(config)
}

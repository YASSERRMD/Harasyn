package gateway

import (
	"time"
)

type SSHConnector struct {
	ID         string            `json:"id"`
	TenantID   string            `json:"tenant_id"`
	ResourceID string            `json:"resource_id"`
	Host       string            `json:"host"`
	Port       int               `json:"port"`
	Username   string            `json:"username"`
	Status     string            `json:"status"`
	Options    map[string]string `json:"options,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

type SSHSession struct {
	ID          string     `json:"id"`
	TenantID    string     `json:"tenant_id"`
	UserID      string     `json:"user_id"`
	ResourceID  string     `json:"resource_id"`
	ConnectorID string     `json:"connector_id"`
	Status      string     `json:"status"`
	StartedAt   time.Time  `json:"started_at"`
	EndedAt     *time.Time `json:"ended_at,omitempty"`
	CommandsRun int        `json:"commands_run"`
}

type SSHCommandMetadata struct {
	ID        string    `json:"id"`
	SessionID string    `json:"session_id"`
	Command   string    `json:"command"`
	Duration  int       `json:"duration"`
	ExitCode  int       `json:"exit_code"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type SSHConnectorRepository interface {
	Create(c *SSHConnector) error
	GetByID(id string) (*SSHConnector, error)
	Update(c *SSHConnector) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*SSHConnector, error)
}

type SSHSessionRepository interface {
	Create(s *SSHSession) error
	GetByID(id string) (*SSHSession, error)
	Update(s *SSHSession) error
	ListByTenant(tenantID string) ([]*SSHSession, error)
	ListByUser(userID string) ([]*SSHSession, error)
}

type SSHAccessService struct {
	connectorRepo SSHConnectorRepository
	sessionRepo   SSHSessionRepository
}

func NewSSHAccessService(cr SSHConnectorRepository, sr SSHSessionRepository) *SSHAccessService {
	return &SSHAccessService{
		connectorRepo: cr,
		sessionRepo:   sr,
	}
}

func (s *SSHAccessService) CreateConnector(tenantID, resourceID, host string, port int, username string) (*SSHConnector, error) {
	connector := &SSHConnector{
		TenantID:   tenantID,
		ResourceID: resourceID,
		Host:       host,
		Port:       port,
		Username:   username,
		Status:     "active",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.connectorRepo.Create(connector); err != nil {
		return nil, err
	}

	return connector, nil
}

func (s *SSHAccessService) GetConnector(id string) (*SSHConnector, error) {
	return s.connectorRepo.GetByID(id)
}

func (s *SSHAccessService) ListConnectors(tenantID string) ([]*SSHConnector, error) {
	return s.connectorRepo.ListByTenant(tenantID)
}

func (s *SSHAccessService) StartSession(tenantID, userID, resourceID, connectorID string) (*SSHSession, error) {
	session := &SSHSession{
		TenantID:    tenantID,
		UserID:      userID,
		ResourceID:  resourceID,
		ConnectorID: connectorID,
		Status:      "active",
		StartedAt:   time.Now(),
	}

	if err := s.sessionRepo.Create(session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *SSHAccessService) EndSession(sessionID string) error {
	session, err := s.sessionRepo.GetByID(sessionID)
	if err != nil {
		return err
	}

	session.Status = "ended"
	now := time.Now()
	session.EndedAt = &now

	return s.sessionRepo.Update(session)
}

func (s *SSHAccessService) GetSession(id string) (*SSHSession, error) {
	return s.sessionRepo.GetByID(id)
}

func (s *SSHAccessService) ListSessions(tenantID string) ([]*SSHSession, error) {
	return s.sessionRepo.ListByTenant(tenantID)
}

func (s *SSHAccessService) CaptureCommandMetadata(sessionID, command string, duration, exitCode int, status string) (*SSHCommandMetadata, error) {
	metadata := &SSHCommandMetadata{
		SessionID: sessionID,
		Command:   command,
		Duration:  duration,
		ExitCode:  exitCode,
		Status:    status,
		CreatedAt: time.Now(),
	}

	session, err := s.sessionRepo.GetByID(sessionID)
	if err != nil {
		return nil, err
	}

	session.CommandsRun++
	s.sessionRepo.Update(session)

	return metadata, nil
}

func (s *SSHAccessService) DeleteConnector(id string) error {
	return s.connectorRepo.Delete(id)
}

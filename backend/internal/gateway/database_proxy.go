package gateway

import (
	"time"
)

type DatabaseConnector struct {
	ID         string            `json:"id"`
	TenantID   string            `json:"tenant_id"`
	ResourceID string            `json:"resource_id"`
	Type       string            `json:"type"`
	Host       string            `json:"host"`
	Port       int               `json:"port"`
	Database   string            `json:"database"`
	Username   string            `json:"username"`
	Status     string            `json:"status"`
	Options    map[string]string `json:"options,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

type DatabaseSession struct {
	ID           string     `json:"id"`
	TenantID     string     `json:"tenant_id"`
	UserID       string     `json:"user_id"`
	ResourceID   string     `json:"resource_id"`
	ConnectorID  string     `json:"connector_id"`
	Status       string     `json:"status"`
	StartedAt    time.Time  `json:"started_at"`
	EndedAt      *time.Time `json:"ended_at,omitempty"`
	QueriesRun   int        `json:"queries_run"`
}

type DatabaseQueryMetadata struct {
	ID        string    `json:"id"`
	SessionID string    `json:"session_id"`
	Query     string    `json:"query"`
	Duration  int       `json:"duration"`
	RowsAffect int     `json:"rows_affected"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type DatabaseConnectorRepository interface {
	Create(c *DatabaseConnector) error
	GetByID(id string) (*DatabaseConnector, error)
	Update(c *DatabaseConnector) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*DatabaseConnector, error)
}

type DatabaseSessionRepository interface {
	Create(s *DatabaseSession) error
	GetByID(id string) (*DatabaseSession, error)
	Update(s *DatabaseSession) error
	ListByTenant(tenantID string) ([]*DatabaseSession, error)
	ListByUser(userID string) ([]*DatabaseSession, error)
}

type DatabaseAccessService struct {
	connectorRepo DatabaseConnectorRepository
	sessionRepo   DatabaseSessionRepository
}

func NewDatabaseAccessService(cr DatabaseConnectorRepository, sr DatabaseSessionRepository) *DatabaseAccessService {
	return &DatabaseAccessService{
		connectorRepo: cr,
		sessionRepo:   sr,
	}
}

func (s *DatabaseAccessService) CreateConnector(tenantID, resourceID, dbType, host string, port int, database, username string) (*DatabaseConnector, error) {
	connector := &DatabaseConnector{
		TenantID:   tenantID,
		ResourceID: resourceID,
		Type:       dbType,
		Host:       host,
		Port:       port,
		Database:   database,
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

func (s *DatabaseAccessService) GetConnector(id string) (*DatabaseConnector, error) {
	return s.connectorRepo.GetByID(id)
}

func (s *DatabaseAccessService) ListConnectors(tenantID string) ([]*DatabaseConnector, error) {
	return s.connectorRepo.ListByTenant(tenantID)
}

func (s *DatabaseAccessService) StartSession(tenantID, userID, resourceID, connectorID string) (*DatabaseSession, error) {
	session := &DatabaseSession{
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

func (s *DatabaseAccessService) EndSession(sessionID string) error {
	session, err := s.sessionRepo.GetByID(sessionID)
	if err != nil {
		return err
	}

	session.Status = "ended"
	now := time.Now()
	session.EndedAt = &now

	return s.sessionRepo.Update(session)
}

func (s *DatabaseAccessService) GetSession(id string) (*DatabaseSession, error) {
	return s.sessionRepo.GetByID(id)
}

func (s *DatabaseAccessService) ListSessions(tenantID string) ([]*DatabaseSession, error) {
	return s.sessionRepo.ListByTenant(tenantID)
}

func (s *DatabaseAccessService) CaptureQueryMetadata(sessionID, query string, duration, rowsAffected int, status string) (*DatabaseQueryMetadata, error) {
	metadata := &DatabaseQueryMetadata{
		SessionID:   sessionID,
		Query:       query,
		Duration:    duration,
		RowsAffect:  rowsAffected,
		Status:      status,
		CreatedAt:   time.Now(),
	}

	session, err := s.sessionRepo.GetByID(sessionID)
	if err != nil {
		return nil, err
	}

	session.QueriesRun++
	s.sessionRepo.Update(session)

	return metadata, nil
}

func (s *DatabaseAccessService) DeleteConnector(id string) error {
	return s.connectorRepo.Delete(id)
}

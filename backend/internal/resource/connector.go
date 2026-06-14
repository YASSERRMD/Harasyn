package resource

import (
	"encoding/json"
	"time"
)

type ConnectorType string

const (
	ConnectorHTTP    ConnectorType = "http"
	ConnectorTCP     ConnectorType = "tcp"
	ConnectorSSH     ConnectorType = "ssh"
	ConnectorDatabase ConnectorType = "database"
)

type ConnectorHealth struct {
	ID          string    `json:"id"`
	ConnectorID string    `json:"connector_id"`
	Status      string    `json:"status"`
	Latency     int64     `json:"latency_ms"`
	CheckedAt   time.Time `json:"checked_at"`
	Error       string    `json:"error,omitempty"`
}

type ConnectorConfig struct {
	Headers     map[string]string `json:"headers,omitempty"`
	TLSVersion  string           `json:"tls_version,omitempty"`
	Timeout     int              `json:"timeout_seconds,omitempty"`
	MaxRetries  int              `json:"max_retries,omitempty"`
}

type ConnectorHealthRepository interface {
	Create(h *ConnectorHealth) error
	GetLatest(connectorID string) (*ConnectorHealth, error)
	ListByConnector(connectorID string) ([]*ConnectorHealth, error)
}

type ConnectorHealthService struct {
	repo ConnectorHealthRepository
}

func NewConnectorHealthService(repo ConnectorHealthRepository) *ConnectorHealthService {
	return &ConnectorHealthService{repo: repo}
}

func (s *ConnectorHealthService) CheckHealth(connectorID string) (*ConnectorHealth, error) {
	health := &ConnectorHealth{
		ConnectorID: connectorID,
		Status:      "healthy",
		Latency:     5,
		CheckedAt:   time.Now(),
	}

	if err := s.repo.Create(health); err != nil {
		return nil, err
	}

	return health, nil
}

func (s *ConnectorHealthService) GetLatestHealth(connectorID string) (*ConnectorHealth, error) {
	return s.repo.GetLatest(connectorID)
}

type HTTPConnector struct {
	ConnectorID string          `json:"connector_id"`
	TargetURL   string          `json:"target_url"`
	Headers     json.RawMessage `json:"headers,omitempty"`
	TLSRequired bool            `json:"tls_required"`
}

type TCPConnector struct {
	ConnectorID string `json:"connector_id"`
	TargetHost  string `json:"target_host"`
	TargetPort  int    `json:"target_port"`
	TLSRequired bool   `json:"tls_required"`
}

type SSHConnector struct {
	ConnectorID string `json:"connector_id"`
	TargetHost  string `json:"target_host"`
	TargetPort  int    `json:"target_port"`
	Username    string `json:"username,omitempty"`
}

type DatabaseConnector struct {
	ConnectorID   string `json:"connector_id"`
	TargetHost    string `json:"target_host"`
	TargetPort    int    `json:"target_port"`
	DatabaseName  string `json:"database_name"`
	DatabaseType  string `json:"database_type"`
	TLSRequired   bool   `json:"tls_required"`
}

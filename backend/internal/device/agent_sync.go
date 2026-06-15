package device

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type AgentPolicySync struct {
	ID               string    `json:"id"`
	TenantID         string    `json:"tenant_id"`
	DeviceID         string    `json:"device_id"`
	PolicyVersion    string    `json:"policy_version"`
	Checksum         string    `json:"checksum"`
	Status           string    `json:"status"`
	LastSyncAt       *time.Time `json:"last_sync_at,omitempty"`
	LastPullAt       *time.Time `json:"last_pull_at,omitempty"`
	PendingChanges   int       `json:"pending_changes"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type AgentConfiguration struct {
	ID            string            `json:"id"`
	TenantID      string            `json:"tenant_id"`
	Version       string            `json:"version"`
	Configuration map[string]interface{} `json:"configuration"`
	Checksum      string            `json:"checksum"`
	IsActive      bool              `json:"is_active"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}

type AgentSyncStatus struct {
	DeviceID      string     `json:"device_id"`
	IsOnline      bool       `json:"is_online"`
	LastHeartbeat *time.Time `json:"last_heartbeat,omitempty"`
	PolicyVersion string     `json:"policy_version"`
	SyncStatus    string     `json:"sync_status"`
}

type AgentPolicySyncRepository interface {
	Create(s *AgentPolicySync) error
	GetByID(id string) (*AgentPolicySync, error)
	GetByDeviceID(deviceID string) (*AgentPolicySync, error)
	Update(s *AgentPolicySync) error
	ListByTenant(tenantID string) ([]*AgentPolicySync, error)
	ListStale(tenantID string, threshold time.Time) ([]*AgentPolicySync, error)
}

type AgentConfigurationRepository interface {
	Create(c *AgentConfiguration) error
	GetByID(id string) (*AgentConfiguration, error)
	GetActiveByTenant(tenantID string) (*AgentConfiguration, error)
	Update(c *AgentConfiguration) error
	ListByTenant(tenantID string) ([]*AgentConfiguration, error)
}

type AgentSyncService struct {
	syncRepo   AgentPolicySyncRepository
	configRepo AgentConfigurationRepository
}

func NewAgentSyncService(sr AgentPolicySyncRepository, cr AgentConfigurationRepository) *AgentSyncService {
	return &AgentSyncService{
		syncRepo:   sr,
		configRepo: cr,
	}
}

func (s *AgentSyncService) RegisterDevice(tenantID, deviceID string) (*AgentPolicySync, error) {
	config, err := s.configRepo.GetActiveByTenant(tenantID)
	if err != nil {
		return nil, err
	}

	sync := &AgentPolicySync{
		TenantID:       tenantID,
		DeviceID:       deviceID,
		PolicyVersion:  config.Version,
		Checksum:       config.Checksum,
		Status:         "synced",
		PendingChanges: 0,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.syncRepo.Create(sync); err != nil {
		return nil, err
	}

	return sync, nil
}

func (s *AgentSyncService) PullPolicy(deviceID string) (*AgentConfiguration, error) {
	sync, err := s.syncRepo.GetByDeviceID(deviceID)
	if err != nil {
		return nil, err
	}

	config, err := s.configRepo.GetActiveByTenant(sync.TenantID)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	sync.LastPullAt = &now
	sync.PolicyVersion = config.Version
	sync.Checksum = config.Checksum
	sync.Status = "synced"
	sync.UpdatedAt = now

	s.syncRepo.Update(sync)

	return config, nil
}

func (s *AgentSyncService) GetSyncStatus(deviceID string) (*AgentSyncStatus, error) {
	sync, err := s.syncRepo.GetByDeviceID(deviceID)
	if err != nil {
		return nil, err
	}

	isOnline := false
	if sync.LastPullAt != nil {
		threshold := time.Now().Add(-5 * time.Minute)
		isOnline = sync.LastPullAt.After(threshold)
	}

	return &AgentSyncStatus{
		DeviceID:      sync.DeviceID,
		IsOnline:      isOnline,
		LastHeartbeat: sync.LastPullAt,
		PolicyVersion: sync.PolicyVersion,
		SyncStatus:    sync.Status,
	}, nil
}

func (s *AgentSyncService) DetectStaleAgents(tenantID string, threshold time.Duration) ([]*AgentPolicySync, error) {
	cutoff := time.Now().Add(-threshold)
	return s.syncRepo.ListStale(tenantID, cutoff)
}

func (s *AgentSyncService) CreateConfiguration(tenantID string, config map[string]interface{}) (*AgentConfiguration, error) {
	data, _ := json.Marshal(config)
	checksum := sha256.Sum256(data)

	agentConfig := &AgentConfiguration{
		TenantID:      tenantID,
		Version:       time.Now().Format("20060102150405"),
		Configuration: config,
		Checksum:      string(checksum[:]),
		IsActive:      true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.configRepo.Create(agentConfig); err != nil {
		return nil, err
	}

	return agentConfig, nil
}

func (s *AgentSyncService) GetConfiguration(id string) (*AgentConfiguration, error) {
	return s.configRepo.GetByID(id)
}

func (s *AgentSyncService) ListConfigurations(tenantID string) ([]*AgentConfiguration, error) {
	return s.configRepo.ListByTenant(tenantID)
}

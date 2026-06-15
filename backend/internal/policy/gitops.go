package policy

import (
	"time"
)

type GitRepository struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	Branch      string    `json:"branch"`
	Path        string    `json:"path"`
	Status      string    `json:"status"`
	LastSyncAt  *time.Time `json:"last_sync_at,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GitOpsSyncConfig struct {
	ID              string    `json:"id"`
	RepositoryID    string    `json:"repository_id"`
	AutoSync        bool      `json:"auto_sync"`
	SyncInterval    int       `json:"sync_interval_minutes"`
	Branch          string    `json:"branch"`
	Path            string    `json:"path"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type GitOpsSyncStatus struct {
	ID             string     `json:"id"`
	ConfigID       string     `json:"config_id"`
	Status         string     `json:"status"`
	LastSyncAt     *time.Time `json:"last_sync_at,omitempty"`
	NextSyncAt     *time.Time `json:"next_sync_at,omitempty"`
	PoliciesSynced int        `json:"policies_synced"`
	Errors         []string   `json:"errors,omitempty"`
}

type GitOpsDrift struct {
	ID          string    `json:"id"`
	ConfigID    string    `json:"config_id"`
	Type        string    `json:"type"`
	Resource    string    `json:"resource"`
	Expected    string    `json:"expected"`
	Actual      string    `json:"actual"`
	DetectedAt  time.Time `json:"detected_at"`
}

type GitRepositoryRepository interface {
	Create(r *GitRepository) error
	GetByID(id string) (*GitRepository, error)
	Update(r *GitRepository) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*GitRepository, error)
}

type GitOpsSyncConfigRepository interface {
	Create(c *GitOpsSyncConfig) error
	GetByID(id string) (*GitOpsSyncConfig, error)
	GetByRepositoryID(repoID string) (*GitOpsSyncConfig, error)
	Update(c *GitOpsSyncConfig) error
	Delete(id string) error
}

type GitOpsSyncStatusRepository interface {
	Create(s *GitOpsSyncStatus) error
	GetByID(id string) (*GitOpsSyncStatus, error)
	GetByConfigID(configID string) (*GitOpsSyncStatus, error)
	Update(s *GitOpsSyncStatus) error
}

type GitOpsDriftRepository interface {
	Create(d *GitOpsDrift) error
	GetByID(id string) (*GitOpsDrift, error)
	Update(d *GitOpsDrift) error
	ListByConfig(configID string) ([]*GitOpsDrift, error)
}

type GitOpsService struct {
	repoRepo    GitRepositoryRepository
	configRepo  GitOpsSyncConfigRepository
	statusRepo  GitOpsSyncStatusRepository
	driftRepo   GitOpsDriftRepository
}

func NewGitOpsService(rr GitRepositoryRepository, cr GitOpsSyncConfigRepository, sr GitOpsSyncStatusRepository, dr GitOpsDriftRepository) *GitOpsService {
	return &GitOpsService{
		repoRepo:   rr,
		configRepo: cr,
		statusRepo: sr,
		driftRepo:  dr,
	}
}

func (s *GitOpsService) RegisterRepository(tenantID, name, url, branch, path string) (*GitRepository, error) {
	repo := &GitRepository{
		TenantID: tenantID,
		Name:     name,
		URL:      url,
		Branch:   branch,
		Path:     path,
		Status:   "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repoRepo.Create(repo); err != nil {
		return nil, err
	}

	return repo, nil
}

func (s *GitOpsService) GetRepository(id string) (*GitRepository, error) {
	return s.repoRepo.GetByID(id)
}

func (s *GitOpsService) ListRepositories(tenantID string) ([]*GitRepository, error) {
	return s.repoRepo.ListByTenant(tenantID)
}

func (s *GitOpsService) ConfigureSync(repositoryID string, autoSync bool, syncInterval int, branch, path string) (*GitOpsSyncConfig, error) {
	config := &GitOpsSyncConfig{
		RepositoryID: repositoryID,
		AutoSync:     autoSync,
		SyncInterval: syncInterval,
		Branch:       branch,
		Path:         path,
		Status:       "active",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.configRepo.Create(config); err != nil {
		return nil, err
	}

	return config, nil
}

func (s *GitOpsService) GetSyncConfig(id string) (*GitOpsSyncConfig, error) {
	return s.configRepo.GetByID(id)
}

func (s *GitOpsService) GetSyncConfigByRepo(repoID string) (*GitOpsSyncConfig, error) {
	return s.configRepo.GetByRepositoryID(repoID)
}

func (s *GitOpsService) UpdateSyncConfig(id string, autoSync bool, syncInterval int) error {
	config, err := s.configRepo.GetByID(id)
	if err != nil {
		return err
	}

	config.AutoSync = autoSync
	config.SyncInterval = syncInterval
	config.UpdatedAt = time.Now()

	return s.configRepo.Update(config)
}

func (s *GitOpsService) DetectDrift(configID string) ([]*GitOpsDrift, error) {
	return s.driftRepo.ListByConfig(configID)
}

func (s *GitOpsService) GetSyncStatus(configID string) (*GitOpsSyncStatus, error) {
	return s.statusRepo.GetByConfigID(configID)
}

func (s *GitOpsService) SyncNow(configID string) (*GitOpsSyncStatus, error) {
	status := &GitOpsSyncStatus{
		ConfigID:       configID,
		Status:         "syncing",
		PoliciesSynced: 0,
		DetectedAt:     time.Now(),
	}

	if err := s.statusRepo.Create(status); err != nil {
		return nil, err
	}

	now := time.Now()
	status.Status = "synced"
	status.LastSyncAt = &now
	s.statusRepo.Update(status)

	return status, nil
}

func (s *GitOpsService) DeleteRepository(id string) error {
	return s.repoRepo.Delete(id)
}

func (s *GitOpsService) DeleteSyncConfig(id string) error {
	return s.configRepo.Delete(id)
}

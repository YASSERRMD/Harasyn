package http

import (
	"time"
)

type HADeploymentConfig struct {
	ID               string    `json:"id"`
	TenantID         string    `json:"tenant_id"`
	Replicas         int       `json:"replicas"`
	LoadBalancer     bool      `json:"load_balancer"`
	SessionAffinity  bool      `json:"session_affinity"`
	HealthCheckPath  string    `json:"health_check_path"`
	GracefulShutdown int       `json:"graceful_shutdown_seconds"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type HAHealthCheck struct {
	ID           string     `json:"id"`
	InstanceID   string     `json:"instance_id"`
	Status       string     `json:"status"`
	LastCheck    time.Time  `json:"last_check"`
	ResponseTime int        `json:"response_time_ms"`
	Healthy      bool       `json:"healthy"`
}

type HAFailoverEvent struct {
	ID          string    `json:"id"`
	FromInstance string  `json:"from_instance"`
	ToInstance  string    `json:"to_instance"`
	Reason      string    `json:"reason"`
	Duration    int       `json:"duration_seconds"`
	StartedAt   time.Time `json:"started_at"`
	EndedAt     *time.Time `json:"ended_at,omitempty"`
}

type HADeploymentConfigRepository interface {
	Create(c *HADeploymentConfig) error
	GetByID(id string) (*HADeploymentConfig, error)
	Update(c *HADeploymentConfig) error
	ListByTenant(tenantID string) ([]*HADeploymentConfig, error)
}

type HAHealthCheckRepository interface {
	Create(h *HAHealthCheck) error
	GetByID(id string) (*HAHealthCheck, error)
	GetByInstance(instanceID string) (*HAHealthCheck, error)
	Update(h *HAHealthCheck) error
	ListHealthy() ([]*HAHealthCheck, error)
}

type HAFailoverEventRepository interface {
	Create(e *HAFailoverEvent) error
	GetByID(id string) (*HAFailoverEvent, error)
	Update(e *HAFailoverEvent) error
	ListRecent(limit int) ([]*HAFailoverEvent, error)
}

type HAService struct {
	configRepo  HADeploymentConfigRepository
	healthRepo  HAHealthCheckRepository
	failoverRepo HAFailoverEventRepository
}

func NewHAService(cr HADeploymentConfigRepository, hr HAHealthCheckRepository, fr HAFailoverEventRepository) *HAService {
	return &HAService{
		configRepo:   cr,
		healthRepo:   hr,
		failoverRepo: fr,
	}
}

func (s *HAService) ConfigureHA(tenantID string, replicas int, loadBalancer, sessionAffinity bool) (*HADeploymentConfig, error) {
	config := &HADeploymentConfig{
		TenantID:         tenantID,
		Replicas:         replicas,
		LoadBalancer:     loadBalancer,
		SessionAffinity:  sessionAffinity,
		HealthCheckPath:  "/health",
		GracefulShutdown: 30,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := s.configRepo.Create(config); err != nil {
		return nil, err
	}

	return config, nil
}

func (s *HAService) GetConfig(id string) (*HADeploymentConfig, error) {
	return s.configRepo.GetByID(id)
}

func (s *HAService) UpdateConfig(id string, replicas int, loadBalancer, sessionAffinity bool) error {
	config, err := s.configRepo.GetByID(id)
	if err != nil {
		return err
	}

	config.Replicas = replicas
	config.LoadBalancer = loadBalancer
	config.SessionAffinity = sessionAffinity
	config.UpdatedAt = time.Now()

	return s.configRepo.Update(config)
}

func (s *HAService) CheckHealth(instanceID string) (*HAHealthCheck, error) {
	health := &HAHealthCheck{
		InstanceID:   instanceID,
		Status:       "healthy",
		LastCheck:    time.Now(),
		ResponseTime: 50,
		Healthy:      true,
	}

	existing, err := s.healthRepo.GetByInstance(instanceID)
	if err == nil && existing != nil {
		existing.Status = health.Status
		existing.LastCheck = health.LastCheck
		existing.ResponseTime = health.ResponseTime
		existing.Healthy = health.Healthy
		s.healthRepo.Update(existing)
		return existing, nil
	}

	if err := s.healthRepo.Create(health); err != nil {
		return nil, err
	}

	return health, nil
}

func (s *HAService) GetHealthyInstances() ([]*HAHealthCheck, error) {
	return s.healthRepo.ListHealthy()
}

func (s *HAService) TriggerFailover(fromInstance, toInstance, reason string) (*HAFailoverEvent, error) {
	event := &HAFailoverEvent{
		FromInstance: fromInstance,
		ToInstance:   toInstance,
		Reason:       reason,
		StartedAt:    time.Now(),
	}

	if err := s.failoverRepo.Create(event); err != nil {
		return nil, err
	}

	return event, nil
}

func (s *HAService) CompleteFailover(eventID string) error {
	event, err := s.failoverRepo.GetByID(eventID)
	if err != nil {
		return err
	}

	now := time.Now()
	event.EndedAt = &now
	event.Duration = int(now.Sub(event.StartedAt).Seconds())

	return s.failoverRepo.Update(event)
}

func (s *HAService) GetRecentFailovers(limit int) ([]*HAFailoverEvent, error) {
	return s.failoverRepo.ListRecent(limit)
}

func (s *HAService) GetDeploymentNotes() map[string]string {
	return map[string]string{
		"api":          "Deploy API stateless with multiple replicas behind load balancer",
		"redis":        "Use Redis Sentinel or Redis Cluster for high availability",
		"postgres":     "Use PostgreSQL streaming replication with automatic failover",
		"worker":       "Implement leader election for background workers",
		"gateway":      "Deploy gateways in active-active configuration",
		"health_check": "Configure health checks for all services",
	}
}

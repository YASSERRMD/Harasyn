package tenant

import (
	"time"
)

type TenantPlan struct {
	ID              string    `json:"id"`
	TenantID        string    `json:"tenant_id"`
	Name            string    `json:"name"`
	MaxUsers        int       `json:"max_users"`
	MaxDevices      int       `json:"max_devices"`
	MaxResources    int       `json:"max_resources"`
	MaxPolicies     int       `json:"max_policies"`
	MaxSessions     int       `json:"max_sessions"`
	Features        []string  `json:"features"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type TenantUsageCounter struct {
	ID              string    `json:"id"`
	TenantID        string    `json:"tenant_id"`
	ResourceType    string    `json:"resource_type"`
	Count           int       `json:"count"`
	LastUpdated     time.Time `json:"last_updated"`
}

type ResourceUsage struct {
	ID           string    `json:"id"`
	TenantID     string    `json:"tenant_id"`
	ResourceType string    `json:"resource_type"`
	ResourceID   string    `json:"resource_id"`
	Action       string    `json:"action"`
	Timestamp    time.Time `json:"timestamp"`
}

type SessionUsage struct {
	ID           string    `json:"id"`
	TenantID     string    `json:"tenant_id"`
	UserID       string    `json:"user_id"`
	ResourceID   string    `json:"resource_id"`
	StartTime    time.Time `json:"start_time"`
	EndTime      *time.Time `json:"end_time,omitempty"`
	Duration     int       `json:"duration_seconds"`
}

type TenantPlanRepository interface {
	Create(p *TenantPlan) error
	GetByID(id string) (*TenantPlan, error)
	GetByTenantID(tenantID string) (*TenantPlan, error)
	Update(p *TenantPlan) error
	Delete(id string) error
}

type TenantUsageCounterRepository interface {
	Create(c *TenantUsageCounter) error
	GetByTenantAndType(tenantID, resourceType string) (*TenantUsageCounter, error)
	Increment(tenantID, resourceType string) error
	GetAllByTenant(tenantID string) ([]*TenantUsageCounter, error)
}

type ResourceUsageRepository interface {
	Create(r *ResourceUsage) error
	ListByTenant(tenantID string, limit int) ([]*ResourceUsage, error)
}

type SessionUsageRepository interface {
	Create(s *SessionUsage) error
	GetByID(id string) (*SessionUsage, error)
	Update(s *SessionUsage) error
	ListByTenant(tenantID string, limit int) ([]*SessionUsage, error)
}

type TenantAdminService struct {
	planRepo      TenantPlanRepository
	counterRepo   TenantUsageCounterRepository
	resourceRepo  ResourceUsageRepository
	sessionRepo   SessionUsageRepository
}

func NewTenantAdminService(pr TenantPlanRepository, cr TenantUsageCounterRepository, rr ResourceUsageRepository, sr SessionUsageRepository) *TenantAdminService {
	return &TenantAdminService{
		planRepo:     pr,
		counterRepo:  cr,
		resourceRepo: rr,
		sessionRepo:  sr,
	}
}

func (s *TenantAdminService) CreatePlan(tenantID, name string, maxUsers, maxDevices, maxResources, maxPolicies, maxSessions int, features []string) (*TenantPlan, error) {
	plan := &TenantPlan{
		TenantID:     tenantID,
		Name:         name,
		MaxUsers:     maxUsers,
		MaxDevices:   maxDevices,
		MaxResources: maxResources,
		MaxPolicies:  maxPolicies,
		MaxSessions:  maxSessions,
		Features:     features,
		Status:       "active",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.planRepo.Create(plan); err != nil {
		return nil, err
	}

	return plan, nil
}

func (s *TenantAdminService) GetPlan(tenantID string) (*TenantPlan, error) {
	return s.planRepo.GetByTenantID(tenantID)
}

func (s *TenantAdminService) UpdatePlan(tenantID, name string, maxUsers, maxDevices, maxResources, maxPolicies, maxSessions int) error {
	plan, err := s.planRepo.GetByTenantID(tenantID)
	if err != nil {
		return err
	}

	plan.Name = name
	plan.MaxUsers = maxUsers
	plan.MaxDevices = maxDevices
	plan.MaxResources = maxResources
	plan.MaxPolicies = maxPolicies
	plan.MaxSessions = maxSessions
	plan.UpdatedAt = time.Now()

	return s.planRepo.Update(plan)
}

func (s *TenantAdminService) CheckLimits(tenantID string) (map[string]bool, error) {
	plan, err := s.planRepo.GetByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	counters, err := s.counterRepo.GetAllByTenant(tenantID)
	if err != nil {
		return nil, err
	}

	usage := make(map[string]int)
	for _, c := range counters {
		usage[c.ResourceType] = c.Count
	}

	limits := map[string]bool{
		"users":     usage["users"] < plan.MaxUsers,
		"devices":   usage["devices"] < plan.MaxDevices,
		"resources": usage["resources"] < plan.MaxResources,
		"policies":  usage["policies"] < plan.MaxPolicies,
		"sessions":  usage["sessions"] < plan.MaxSessions,
	}

	return limits, nil
}

func (s *TenantAdminService) TrackResourceUsage(tenantID, resourceType, resourceID, action string) error {
	usage := &ResourceUsage{
		TenantID:     tenantID,
		ResourceType: resourceType,
		ResourceID:   resourceID,
		Action:       action,
		Timestamp:    time.Now(),
	}

	if err := s.resourceRepo.Create(usage); err != nil {
		return err
	}

	return s.counterRepo.Increment(tenantID, resourceType)
}

func (s *TenantAdminService) StartSessionUsage(tenantID, userID, resourceID string) (*SessionUsage, error) {
	usage := &SessionUsage{
		TenantID:   tenantID,
		UserID:     userID,
		ResourceID: resourceID,
		StartTime:  time.Now(),
	}

	if err := s.sessionRepo.Create(usage); err != nil {
		return nil, err
	}

	return usage, nil
}

func (s *TenantAdminService) EndSessionUsage(sessionID string) error {
	usage, err := s.sessionRepo.GetByID(sessionID)
	if err != nil {
		return err
	}

	now := time.Now()
	usage.EndTime = &now
	usage.Duration = int(now.Sub(usage.StartTime).Seconds())

	return s.sessionRepo.Update(usage)
}

func (s *TenantAdminService) GetUsageStats(tenantID string) (map[string]interface{}, error) {
	counters, err := s.counterRepo.GetAllByTenant(tenantID)
	if err != nil {
		return nil, err
	}

	stats := make(map[string]interface{})
	for _, c := range counters {
		stats[c.ResourceType] = c.Count
	}

	return stats, nil
}

func (s *TenantAdminService) GetRecentResourceUsage(tenantID string, limit int) ([]*ResourceUsage, error) {
	return s.resourceRepo.ListByTenant(tenantID, limit)
}

func (s *TenantAdminService) GetRecentSessionUsage(tenantID string, limit int) ([]*SessionUsage, error) {
	return s.sessionRepo.ListByTenant(tenantID, limit)
}

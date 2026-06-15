package resource

import (
	"time"
)

type SensitivityLevel string

const (
	SensitivityPublic      SensitivityLevel = "public"
	SensitivityInternal    SensitivityLevel = "internal"
	SensitivityConfidential SensitivityLevel = "confidential"
	SensitivityRestricted  SensitivityLevel = "restricted"
)

type ResourceSensitivity struct {
	ID         string           `json:"id"`
	TenantID   string           `json:"tenant_id"`
	ResourceID string           `json:"resource_id"`
	Level      SensitivityLevel `json:"level"`
	AssignedBy string           `json:"assigned_by"`
	AssignedAt time.Time        `json:"assigned_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
}

type SensitivityInheritance struct {
	ID              string           `json:"id"`
	ParentID        string           `json:"parent_id"`
	ChildID         string           `json:"child_id"`
	Level           SensitivityLevel `json:"level"`
	InheritedAt     time.Time        `json:"inherited_at"`
}

type SensitivityPolicyCondition struct {
	ID          string           `json:"id"`
	TenantID    string           `json:"tenant_id"`
	Type        string           `json:"type"`
	Level       SensitivityLevel `json:"level"`
	RequireMFA  bool             `json:"require_mfa"`
	RequireApproval bool         `json:"require_approval"`
	MaxRiskScore int             `json:"max_risk_score"`
}

type ResourceSensitivityRepository interface {
	Create(r *ResourceSensitivity) error
	GetByID(id string) (*ResourceSensitivity, error)
	GetByResourceID(resourceID string) (*ResourceSensitivity, error)
	Update(r *ResourceSensitivity) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*ResourceSensitivity, error)
	ListByLevel(tenantID string, level SensitivityLevel) ([]*ResourceSensitivity, error)
}

type SensitivityInheritanceRepository interface {
	Create(i *SensitivityInheritance) error
	GetByID(id string) (*SensitivityInheritance, error)
	GetByChildID(childID string) (*SensitivityInheritance, error)
	Delete(id string) error
	ListByParent(parentID string) ([]*SensitivityInheritance, error)
}

type SensitivityPolicyConditionRepository interface {
	Create(c *SensitivityPolicyCondition) error
	GetByID(id string) (*SensitivityPolicyCondition, error)
	Update(c *SensitivityPolicyCondition) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*SensitivityPolicyCondition, error)
}

type ResourceSensitivityService struct {
	sensitivityRepo ResourceSensitivityRepository
	inheritanceRepo SensitivityInheritanceRepository
	conditionRepo   SensitivityPolicyConditionRepository
}

func NewResourceSensitivityService(sr ResourceSensitivityRepository, ir SensitivityInheritanceRepository, cr SensitivityPolicyConditionRepository) *ResourceSensitivityService {
	return &ResourceSensitivityService{
		sensitivityRepo: sr,
		inheritanceRepo: ir,
		conditionRepo:   cr,
	}
}

func (s *ResourceSensitivityService) AssignSensitivity(tenantID, resourceID string, level SensitivityLevel, assignedBy string) (*ResourceSensitivity, error) {
	sensitivity := &ResourceSensitivity{
		TenantID:   tenantID,
		ResourceID: resourceID,
		Level:      level,
		AssignedBy: assignedBy,
		AssignedAt: time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.sensitivityRepo.Create(sensitivity); err != nil {
		return nil, err
	}

	return sensitivity, nil
}

func (s *ResourceSensitivityService) GetSensitivity(resourceID string) (*ResourceSensitivity, error) {
	return s.sensitivityRepo.GetByResourceID(resourceID)
}

func (s *ResourceSensitivityService) ListSensitivities(tenantID string) ([]*ResourceSensitivity, error) {
	return s.sensitivityRepo.ListByTenant(tenantID)
}

func (s *ResourceSensitivityService) ListByLevel(tenantID string, level SensitivityLevel) ([]*ResourceSensitivity, error) {
	return s.sensitivityRepo.ListByLevel(tenantID, level)
}

func (s *ResourceSensitivityService) UpdateSensitivity(resourceID string, level SensitivityLevel) error {
	sensitivity, err := s.sensitivityRepo.GetByResourceID(resourceID)
	if err != nil {
		return err
	}

	sensitivity.Level = level
	sensitivity.UpdatedAt = time.Now()

	return s.sensitivityRepo.Update(sensitivity)
}

func (s *ResourceSensitivityService) InheritSensitivity(parentID, childID string) error {
	parent, err := s.sensitivityRepo.GetByResourceID(parentID)
	if err != nil {
		return err
	}

	inheritance := &SensitivityInheritance{
		ParentID:    parentID,
		ChildID:     childID,
		Level:       parent.Level,
		InheritedAt: time.Now(),
	}

	if err := s.inheritanceRepo.Create(inheritance); err != nil {
		return err
	}

	child := &ResourceSensitivity{
		TenantID:   parent.TenantID,
		ResourceID: childID,
		Level:      parent.Level,
		AssignedBy: "inheritance",
		AssignedAt: time.Now(),
		UpdatedAt:  time.Now(),
	}

	return s.sensitivityRepo.Create(child)
}

func (s *ResourceSensitivityService) GetInheritedSensitivity(childID string) (*SensitivityInheritance, error) {
	return s.inheritanceRepo.GetByChildID(childID)
}

func (s *ResourceSensitivityService) EvaluateHighSensitivityAccess(tenantID, resourceID string, hasMFA bool) (bool, string) {
	sensitivity, err := s.sensitivityRepo.GetByResourceID(resourceID)
	if err != nil {
		return false, "resource not found"
	}

	if sensitivity.Level == SensitivityRestricted || sensitivity.Level == SensitivityConfidential {
		if !hasMFA {
			return false, "MFA required for high sensitivity access"
		}
	}

	return true, ""
}

func (s *ResourceSensitivityService) CreatePolicyCondition(tenantID, conditionType string, level SensitivityLevel, requireMFA, requireApproval bool, maxRiskScore int) (*SensitivityPolicyCondition, error) {
	condition := &SensitivityPolicyCondition{
		TenantID:        tenantID,
		Type:            conditionType,
		Level:           level,
		RequireMFA:      requireMFA,
		RequireApproval: requireApproval,
		MaxRiskScore:    maxRiskScore,
	}

	if err := s.conditionRepo.Create(condition); err != nil {
		return nil, err
	}

	return condition, nil
}

func (s *ResourceSensitivityService) GetPolicyCondition(id string) (*SensitivityPolicyCondition, error) {
	return s.conditionRepo.GetByID(id)
}

func (s *ResourceSensitivityService) ListPolicyConditions(tenantID string) ([]*SensitivityPolicyCondition, error) {
	return s.conditionRepo.ListByTenant(tenantID)
}

func (s *ResourceSensitivityService) EvaluatePolicyConditions(tenantID string, level SensitivityLevel) []SensitivityPolicyCondition {
	conditions, err := s.conditionRepo.ListByTenant(tenantID)
	if err != nil {
		return nil
	}

	var matching []SensitivityPolicyCondition
	for _, c := range conditions {
		if c.Level == level {
			matching = append(matching, *c)
		}
	}

	return matching
}

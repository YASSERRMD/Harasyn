package policy

import (
	"encoding/json"
	"fmt"
	"time"
)

type PolicyDocument struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   PolicyMetadata    `json:"metadata"`
	Spec       PolicySpec        `json:"spec"`
}

type PolicyMetadata struct {
	Name        string            `json:"name"`
	TenantID    string            `json:"tenant_id"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type PolicySpec struct {
	Description string         `json:"description,omitempty"`
	Effect      string         `json:"effect"`
	Resources   []string       `json:"resources"`
	Actions     []string       `json:"actions"`
	Conditions  []PolicyCondition `json:"conditions"`
}

type PolicyCondition struct {
	Type     string      `json:"type"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type PolicyImportResult struct {
	Imported  int      `json:"imported"`
	Conflicts []string `json:"conflicts,omitempty"`
	Errors    []string `json:"errors,omitempty"`
}

type PolicyExportResult struct {
	Policies  []PolicyDocument `json:"policies"`
	Format    string           `json:"format"`
	ExportedAt time.Time       `json:"exported_at"`
}

type PolicyRepository interface {
	Create(p *PolicyDocument) error
	GetByID(id string) (*PolicyDocument, error)
	GetByName(tenantID, name string) (*PolicyDocument, error)
	Update(p *PolicyDocument) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*PolicyDocument, error)
}

type PolicyAsCodeService struct {
	repo PolicyRepository
}

func NewPolicyAsCodeService(repo PolicyRepository) *PolicyAsCodeService {
	return &PolicyAsCodeService{repo: repo}
}

func (s *PolicyAsCodeService) ImportPoliciesJSON(data []byte, tenantID string) (*PolicyImportResult, error) {
	var documents []PolicyDocument
	if err := json.Unmarshal(data, &documents); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	result := &PolicyImportResult{}

	for _, doc := range documents {
		doc.Metadata.TenantID = tenantID

		existing, _ := s.repo.GetByName(tenantID, doc.Metadata.Name)
		if existing != nil {
			result.Conflicts = append(result.Conflicts, doc.Metadata.Name)
			continue
		}

		if err := s.repo.Create(&doc); err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("%s: %v", doc.Metadata.Name, err))
			continue
		}

		result.Imported++
	}

	return result, nil
}

func (s *PolicyAsCodeService) ImportPoliciesYAML(data []byte, tenantID string) (*PolicyImportResult, error) {
	return s.ImportPoliciesJSON(data, tenantID)
}

func (s *PolicyAsCodeService) ExportPoliciesJSON(tenantID string) (*PolicyExportResult, error) {
	policies, err := s.repo.ListByTenant(tenantID)
	if err != nil {
		return nil, err
	}

	result := &PolicyExportResult{
		Policies:   make([]PolicyDocument, 0),
		Format:     "json",
		ExportedAt: time.Now(),
	}

	for _, p := range policies {
		result.Policies = append(result.Policies, *p)
	}

	return result, nil
}

func (s *PolicyAsCodeService) DryRunImport(data []byte, tenantID string) (*PolicyImportResult, error) {
	return s.ImportPoliciesJSON(data, tenantID)
}

func (s *PolicyAsCodeService) DetectConflicts(data []byte, tenantID string) ([]string, error) {
	var documents []PolicyDocument
	if err := json.Unmarshal(data, &documents); err != nil {
		return nil, err
	}

	var conflicts []string
	for _, doc := range documents {
		existing, _ := s.repo.GetByName(tenantID, doc.Metadata.Name)
		if existing != nil {
			conflicts = append(conflicts, doc.Metadata.Name)
		}
	}

	return conflicts, nil
}

func (s *PolicyAsCodeService) GetPolicy(id string) (*PolicyDocument, error) {
	return s.repo.GetByID(id)
}

func (s *PolicyAsCodeService) ListPolicies(tenantID string) ([]*PolicyDocument, error) {
	return s.repo.ListByTenant(tenantID)
}

func (s *PolicyAsCodeService) DeletePolicy(id string) error {
	return s.repo.Delete(id)
}

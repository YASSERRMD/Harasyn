package identity

import (
	"time"
)

type SAMLProvider struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	Name        string    `json:"name"`
	EntityID    string    `json:"entity_id"`
	MetadataURL string    `json:"metadata_url,omitempty"`
	Status      string    `json:"status"`
	TrustStatus string    `json:"trust_status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SAMLMetadata struct {
	ID         string    `json:"id"`
	ProviderID string    `json:"provider_id"`
	EntityID   string    `json:"entity_id"`
	SSOURL     string    `json:"sso_url"`
	Certificate string   `json:"certificate"`
	Format     string    `json:"format"`
	CreatedAt  time.Time `json:"created_at"`
}

type SAMLAssertion struct {
	ID           string            `json:"id"`
	Issuer       string            `json:"issuer"`
	Subject      string            `json:"subject"`
	NameID       string            `json:"name_id"`
	NameIDFormat string            `json:"name_id_format"`
	Attributes   map[string]string `json:"attributes,omitempty"`
	Conditions   *SAMLConditions   `json:"conditions,omitempty"`
	AuthnContext string            `json:"authn_context,omitempty"`
}

type SAMLConditions struct {
	NotBefore    time.Time  `json:"not_before"`
	NotOnOrAfter time.Time  `json:"not_on_or_after"`
	Audience     string     `json:"audience"`
}

type SAMLMappingRule struct {
	ID         string            `json:"id"`
	ProviderID string            `json:"provider_id"`
	Source     string            `json:"source"`
	Target     string            `json:"target"`
	Attribute  string            `json:"attribute"`
	Default    string            `json:"default,omitempty"`
	Condition  string            `json:"condition,omitempty"`
}

type SAMLProviderRepository interface {
	Create(p *SAMLProvider) error
	GetByID(id string) (*SAMLProvider, error)
	Update(p *SAMLProvider) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*SAMLProvider, error)
}

type SAMLMetadataRepository interface {
	Create(m *SAMLMetadata) error
	GetByID(id string) (*SAMLMetadata, error)
	GetByProviderID(providerID string) (*SAMLMetadata, error)
	Update(m *SAMLMetadata) error
	Delete(id string) error
}

type SAMLMappingRuleRepository interface {
	Create(r *SAMLMappingRule) error
	GetByID(id string) (*SAMLMappingRule, error)
	Update(r *SAMLMappingRule) error
	Delete(id string) error
	ListByProvider(providerID string) ([]*SAMLMappingRule, error)
}

type SAMLService struct {
	providerRepo SAMLProviderRepository
	metadataRepo SAMLMetadataRepository
	mappingRepo  SAMLMappingRuleRepository
}

func NewSAMLService(pr SAMLProviderRepository, mr SAMLMetadataRepository, mapRepo SAMLMappingRuleRepository) *SAMLService {
	return &SAMLService{
		providerRepo: pr,
		metadataRepo: mr,
		mappingRepo:  mapRepo,
	}
}

func (s *SAMLService) CreateProvider(tenantID, name, entityID string) (*SAMLProvider, error) {
	provider := &SAMLProvider{
		TenantID:    tenantID,
		Name:        name,
		EntityID:    entityID,
		Status:      "active",
		TrustStatus: "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.providerRepo.Create(provider); err != nil {
		return nil, err
	}

	return provider, nil
}

func (s *SAMLService) GetProvider(id string) (*SAMLProvider, error) {
	return s.providerRepo.GetByID(id)
}

func (s *SAMLService) ListProviders(tenantID string) ([]*SAMLProvider, error) {
	return s.providerRepo.ListByTenant(tenantID)
}

func (s *SAMLService) ValidateSAMLAssertion(assertion *SAMLAssertion, providerID string) (bool, error) {
	provider, err := s.providerRepo.GetByID(providerID)
	if err != nil {
		return false, err
	}

	_ = provider

	if assertion.Conditions != nil {
		now := time.Now()
		if now.Before(assertion.Conditions.NotBefore) || now.After(assertion.Conditions.NotOnOrAfter) {
			return false, nil
		}
	}

	return true, nil
}

func (s *SAMLService) ApplyMappingRules(providerID string, attributes map[string]string) (map[string]string, error) {
	rules, err := s.mappingRepo.ListByProvider(providerID)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, rule := range rules {
		if val, ok := attributes[rule.Attribute]; ok {
			result[rule.Target] = val
		} else if rule.Default != "" {
			result[rule.Target] = rule.Default
		}
	}

	return result, nil
}

func (s *SAMLService) GetMetadata(providerID string) (*SAMLMetadata, error) {
	return s.metadataRepo.GetByProviderID(providerID)
}

func (s *SAMLService) StoreMetadata(providerID, entityID, ssoURL, certificate, format string) (*SAMLMetadata, error) {
	metadata := &SAMLMetadata{
		ProviderID:  providerID,
		EntityID:    entityID,
		SSOURL:      ssoURL,
		Certificate: certificate,
		Format:      format,
		CreatedAt:   time.Now(),
	}

	if err := s.metadataRepo.Create(metadata); err != nil {
		return nil, err
	}

	return metadata, nil
}

func (s *SAMLService) DeleteProvider(id string) error {
	return s.providerRepo.Delete(id)
}

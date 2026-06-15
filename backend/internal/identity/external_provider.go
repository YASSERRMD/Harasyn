package identity

import (
	"time"
)

type ExternalProvider struct {
	ID          string            `json:"id"`
	TenantID    string            `json:"tenant_id"`
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	Issuer      string            `json:"issuer"`
	ClientID    string            `json:"client_id"`
	Status      string            `json:"status"`
	TrustStatus string            `json:"trust_status"`
	Config      map[string]string `json:"config,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type OIDCConfig struct {
	Issuer                string   `json:"issuer"`
	AuthorizationEndpoint string   `json:"authorization_endpoint"`
	TokenEndpoint         string   `json:"token_endpoint"`
	UserinfoEndpoint      string   `json:"userinfo_endpoint"`
	JWKSEndpoint          string   `json:"jwks_endpoint"`
	Scopes                []string `json:"scopes"`
}

type OIDCDiscoveryMetadata struct {
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
	JWKSEndpoint          string `json:"jwks_uri"`
	RegistrationEndpoint  string `json:"registration_endpoint,omitempty"`
}

type ExternalUserMapping struct {
	ID           string    `json:"id"`
	TenantID     string    `json:"tenant_id"`
	ProviderID   string    `json:"provider_id"`
	ExternalID   string    `json:"external_id"`
	InternalID   string    `json:"internal_id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Groups       []string  `json:"groups,omitempty"`
	Attributes   map[string]string `json:"attributes,omitempty"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ExternalProviderRepository interface {
	Create(p *ExternalProvider) error
	GetByID(id string) (*ExternalProvider, error)
	Update(p *ExternalProvider) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*ExternalProvider, error)
}

type ExternalUserMappingRepository interface {
	Create(m *ExternalUserMapping) error
	GetByID(id string) (*ExternalUserMapping, error)
	Update(m *ExternalUserMapping) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*ExternalUserMapping, error)
	ListByProvider(providerID string) ([]*ExternalUserMapping, error)
	FindByExternalID(providerID, externalID string) (*ExternalUserMapping, error)
}

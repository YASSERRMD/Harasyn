package resource

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type ServiceIdentity struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	Name        string    `json:"name"`
	ServiceType string    `json:"service_type"`
	TrustScore  int       `json:"trust_score"`
	Status      string    `json:"status"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ServiceCredential struct {
	ID           string    `json:"id"`
	ServiceID    string    `json:"service_id"`
	CredentialType string `json:"credential_type"`
	Secret       string   `json:"secret,omitempty"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
	RevokedAt    *time.Time `json:"revoked_at,omitempty"`
}

type ServiceIdentityRepository interface {
	Create(s *ServiceIdentity) error
	GetByID(id string) (*ServiceIdentity, error)
	Update(s *ServiceIdentity) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*ServiceIdentity, error)
}

type ServiceCredentialRepository interface {
	Create(c *ServiceCredential) error
	GetByID(id string) (*ServiceCredential, error)
	GetByServiceID(serviceID string) (*ServiceCredential, error)
	Revoke(id string) error
}

type ServiceAccessService struct {
	identityRepo  ServiceIdentityRepository
	credentialRepo ServiceCredentialRepository
}

func NewServiceAccessService(ir ServiceIdentityRepository, cr ServiceCredentialRepository) *ServiceAccessService {
	return &ServiceAccessService{
		identityRepo:   ir,
		credentialRepo: cr,
	}
}

type RegisterServiceRequest struct {
	TenantID    string `json:"tenant_id"`
	Name        string `json:"name"`
	ServiceType string `json:"service_type"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

func (s *ServiceAccessService) RegisterService(req RegisterServiceRequest) (*ServiceIdentity, error) {
	identity := &ServiceIdentity{
		TenantID:    req.TenantID,
		Name:        req.Name,
		ServiceType: req.ServiceType,
		TrustScore:  50,
		Status:      "active",
		Metadata:    req.Metadata,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.identityRepo.Create(identity); err != nil {
		return nil, err
	}

	return identity, nil
}

func (s *ServiceAccessService) IssueCredential(serviceID string) (*ServiceCredential, error) {
	secret, err := generateSecret()
	if err != nil {
		return nil, err
	}

	credential := &ServiceCredential{
		ServiceID:      serviceID,
		CredentialType: "api_key",
		Secret:         secret,
		ExpiresAt:      time.Now().Add(90 * 24 * time.Hour),
		CreatedAt:      time.Now(),
	}

	if err := s.credentialRepo.Create(credential); err != nil {
		return nil, err
	}

	return credential, nil
}

func (s *ServiceAccessService) ValidateCredential(serviceID, secret string) (bool, error) {
	cred, err := s.credentialRepo.GetByServiceID(serviceID)
	if err != nil {
		return false, err
	}

	if cred.Secret != secret {
		return false, nil
	}

	if time.Now().After(cred.ExpiresAt) {
		return false, nil
	}

	if cred.RevokedAt != nil {
		return false, nil
	}

	return true, nil
}

func (s *ServiceAccessService) GetServiceTrustScore(serviceID string) (int, error) {
	identity, err := s.identityRepo.GetByID(serviceID)
	if err != nil {
		return 0, err
	}
	return identity.TrustScore, nil
}

func (s *ServiceAccessService) ListServices(tenantID string) ([]*ServiceIdentity, error) {
	return s.identityRepo.ListByTenant(tenantID)
}

func generateSecret() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

package identity

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OIDCProviderService struct {
	providerRepo ExternalProviderRepository
	httpClient   *http.Client
}

func NewOIDCProviderService(repo ExternalProviderRepository) *OIDCProviderService {
	return &OIDCProviderService{
		providerRepo: repo,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *OIDCProviderService) GetDiscoveryMetadata(ctx context.Context, issuer string) (*OIDCDiscoveryMetadata, error) {
	discoveryURL := issuer + "/.well-known/openid-configuration"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, discoveryURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch discovery metadata: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var metadata OIDCDiscoveryMetadata
	if err := json.Unmarshal(body, &metadata); err != nil {
		return nil, fmt.Errorf("failed to unmarshal metadata: %w", err)
	}

	return &metadata, nil
}

func (s *OIDCProviderService) CreateProvider(ctx context.Context, tenantID, name, issuer, clientID string) (*ExternalProvider, error) {
	provider := &ExternalProvider{
		TenantID:    tenantID,
		Name:        name,
		Type:        "oidc",
		Issuer:      issuer,
		ClientID:    clientID,
		Status:      "active",
		TrustStatus: "pending",
		Config:      make(map[string]string),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.providerRepo.Create(provider); err != nil {
		return nil, err
	}

	return provider, nil
}

func (s *OIDCProviderService) GetProvider(id string) (*ExternalProvider, error) {
	return s.providerRepo.GetByID(id)
}

func (s *OIDCProviderService) ListProviders(tenantID string) ([]*ExternalProvider, error) {
	return s.providerRepo.ListByTenant(tenantID)
}

func (s *OIDCProviderService) UpdateTrustStatus(id, trustStatus string) (*ExternalProvider, error) {
	provider, err := s.providerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	provider.TrustStatus = trustStatus
	provider.UpdatedAt = time.Now()

	if err := s.providerRepo.Update(provider); err != nil {
		return nil, err
	}

	return provider, nil
}

func (s *OIDCProviderService) DeactivateProvider(id string) error {
	provider, err := s.providerRepo.GetByID(id)
	if err != nil {
		return err
	}

	provider.Status = "inactive"
	provider.UpdatedAt = time.Now()

	return s.providerRepo.Update(provider)
}

func (s *OIDCProviderService) DeleteProvider(id string) error {
	return s.providerRepo.Delete(id)
}

type LoginInitiationRequest struct {
	ProviderID string `json:"provider_id"`
	RedirectURI string `json:"redirect_uri"`
	State       string `json:"state"`
}

type LoginInitiationResponse struct {
	AuthorizationURL string `json:"authorization_url"`
	State            string `json:"state"`
}

func (s *OIDCProviderService) InitiateLogin(ctx context.Context, req LoginInitiationRequest) (*LoginInitiationResponse, error) {
	provider, err := s.providerRepo.GetByID(req.ProviderID)
	if err != nil {
		return nil, err
	}

	metadata, err := s.GetDiscoveryMetadata(ctx, provider.Issuer)
	if err != nil {
		return nil, fmt.Errorf("failed to get discovery metadata: %w", err)
	}

	authURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=openid+profile+email&state=%s",
		metadata.AuthorizationEndpoint,
		provider.ClientID,
		req.RedirectURI,
		req.State,
	)

	return &LoginInitiationResponse{
		AuthorizationURL: authURL,
		State:            req.State,
	}, nil
}

type CallbackRequest struct {
	ProviderID string `json:"provider_id"`
	Code       string `json:"code"`
	State      string `json:"state"`
}

type CallbackResponse struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	SessionToken string `json:"session_token"`
}

func (s *OIDCProviderService) HandleCallback(ctx context.Context, req CallbackRequest) (*CallbackResponse, error) {
	provider, err := s.providerRepo.GetByID(req.ProviderID)
	if err != nil {
		return nil, err
	}

	_ = provider

	return &CallbackResponse{
		UserID:    "ext-user-1",
		Email:     "user@example.com",
		Name:      "External User",
		SessionToken: "session-token-placeholder",
	}, nil
}

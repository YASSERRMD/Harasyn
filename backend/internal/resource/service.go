package resource

import (
	"fmt"
)

type Service struct {
	repo            Repository
	connectorRepo   ConnectorRepository
}

func NewService(repo Repository, connectorRepo ConnectorRepository) *Service {
	return &Service{
		repo:          repo,
		connectorRepo: connectorRepo,
	}
}

type RegisterResourceRequest struct {
	TenantID     string `json:"tenant_id"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	ResourceType string `json:"resource_type"`
	Endpoint     string `json:"endpoint,omitempty"`
	Port         int    `json:"port,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
	Sensitivity  string `json:"sensitivity"`
}

func (s *Service) RegisterResource(req RegisterResourceRequest) (*Resource, error) {
	res := &Resource{
		TenantID:     req.TenantID,
		Name:         req.Name,
		Description:  req.Description,
		ResourceType: req.ResourceType,
		Endpoint:     req.Endpoint,
		Port:         req.Port,
		Protocol:     req.Protocol,
		Sensitivity:  req.Sensitivity,
		Status:       "active",
	}

	if err := s.repo.Create(res); err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	return res, nil
}

func (s *Service) GetResource(id string) (*Resource, error) {
	return s.repo.GetByID(id)
}

func (s *Service) ListResourcesByTenant(tenantID string) ([]*Resource, error) {
	return s.repo.ListByTenant(tenantID)
}

func (s *Service) UpdateResource(id string, updates map[string]interface{}) (*Resource, error) {
	res, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if name, ok := updates["name"].(string); ok {
		res.Name = name
	}
	if desc, ok := updates["description"].(string); ok {
		res.Description = desc
	}
	if endpoint, ok := updates["endpoint"].(string); ok {
		res.Endpoint = endpoint
	}
	if sensitivity, ok := updates["sensitivity"].(string); ok {
		res.Sensitivity = sensitivity
	}
	if status, ok := updates["status"].(string); ok {
		res.Status = status
	}

	if err := s.repo.Update(res); err != nil {
		return nil, fmt.Errorf("failed to update resource: %w", err)
	}

	return res, nil
}

func (s *Service) DeleteResource(id string) error {
	return s.repo.Delete(id)
}

type RegisterConnectorRequest struct {
	ResourceID    string `json:"resource_id"`
	ConnectorType string `json:"connector_type"`
	TargetHost    string `json:"target_host"`
	TargetPort    int    `json:"target_port"`
	TLSRequired   bool   `json:"tls_required"`
	AuthMethod    string `json:"auth_method"`
}

func (s *Service) RegisterConnector(req RegisterConnectorRequest) (*Connector, error) {
	conn := &Connector{
		ResourceID:    req.ResourceID,
		ConnectorType: req.ConnectorType,
		TargetHost:    req.TargetHost,
		TargetPort:    req.TargetPort,
		TLSRequired:   req.TLSRequired,
		AuthMethod:    req.AuthMethod,
		Status:        "active",
	}

	if err := s.connectorRepo.Create(conn); err != nil {
		return nil, fmt.Errorf("failed to create connector: %w", err)
	}

	return conn, nil
}

func (s *Service) GetConnector(resourceID string) (*Connector, error) {
	return s.connectorRepo.GetByResourceID(resourceID)
}

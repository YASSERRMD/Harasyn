package gateway

import (
	"time"
)

type Region struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	IsPrimary   bool      `json:"is_primary"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RegionalGateway struct {
	ID          string    `json:"id"`
	RegionID    string    `json:"region_id"`
	Address     string    `json:"address"`
	Port        int       `json:"port"`
	Status      string    `json:"status"`
	Load        int       `json:"load"`
	MaxCapacity int       `json:"max_capacity"`
	LastHealth  *time.Time `json:"last_health,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RegionalResourcePlacement struct {
	ID           string    `json:"id"`
	ResourceID   string    `json:"resource_id"`
	RegionID     string    `json:"region_id"`
	IsReplicated bool      `json:"is_replicated"`
	Replicas     []string  `json:"replicas,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RegionHealthStatus struct {
	ID           string     `json:"id"`
	RegionID     string     `json:"region_id"`
	Status       string     `json:"status"`
	ActiveConns  int        `json:"active_connections"`
	Healthy      bool       `json:"healthy"`
	LastCheck    time.Time  `json:"last_check"`
}

type RegionRepository interface {
	Create(r *Region) error
	GetByID(id string) (*Region, error)
	Update(r *Region) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*Region, error)
	GetPrimary(tenantID string) (*Region, error)
}

type RegionalGatewayRepository interface {
	Create(g *RegionalGateway) error
	GetByID(id string) (*RegionalGateway, error)
	Update(g *RegionalGateway) error
	Delete(id string) error
	ListByRegion(regionID string) ([]*RegionalGateway, error)
}

type RegionalResourcePlacementRepository interface {
	Create(p *RegionalResourcePlacement) error
	GetByID(id string) (*RegionalResourcePlacement, error)
	GetByResource(resourceID string) (*RegionalResourcePlacement, error)
	Update(p *RegionalResourcePlacement) error
	Delete(id string) error
	ListByRegion(regionID string) ([]*RegionalResourcePlacement, error)
}

type RegionHealthStatusRepository interface {
	Create(h *RegionHealthStatus) error
	GetByID(id string) (*RegionHealthStatus, error)
	GetByRegion(regionID string) (*RegionHealthStatus, error)
	Update(h *RegionHealthStatus) error
}

type MultiRegionService struct {
	regionRepo     RegionRepository
	gatewayRepo    RegionalGatewayRepository
	placementRepo  RegionalResourcePlacementRepository
	healthRepo     RegionHealthStatusRepository
}

func NewMultiRegionService(rr RegionRepository, gr RegionalGatewayRepository, pr RegionalResourcePlacementRepository, hr RegionHealthStatusRepository) *MultiRegionService {
	return &MultiRegionService{
		regionRepo:    rr,
		gatewayRepo:   gr,
		placementRepo: pr,
		healthRepo:    hr,
	}
}

func (s *MultiRegionService) CreateRegion(tenantID, name, code, location string, isPrimary bool) (*Region, error) {
	region := &Region{
		TenantID:  tenantID,
		Name:      name,
		Code:      code,
		Location:  location,
		Status:    "active",
		IsPrimary: isPrimary,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.regionRepo.Create(region); err != nil {
		return nil, err
	}

	return region, nil
}

func (s *MultiRegionService) GetRegion(id string) (*Region, error) {
	return s.regionRepo.GetByID(id)
}

func (s *MultiRegionService) ListRegions(tenantID string) ([]*Region, error) {
	return s.regionRepo.ListByTenant(tenantID)
}

func (s *MultiRegionService) GetPrimaryRegion(tenantID string) (*Region, error) {
	return s.regionRepo.GetPrimary(tenantID)
}

func (s *MultiRegionService) RegisterGateway(regionID, address string, port, maxCapacity int) (*RegionalGateway, error) {
	gateway := &RegionalGateway{
		RegionID:    regionID,
		Address:     address,
		Port:        port,
		Status:      "active",
		Load:        0,
		MaxCapacity: maxCapacity,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.gatewayRepo.Create(gateway); err != nil {
		return nil, err
	}

	return gateway, nil
}

func (s *MultiRegionService) GetGateway(id string) (*RegionalGateway, error) {
	return s.gatewayRepo.GetByID(id)
}

func (s *MultiRegionService) ListGateways(regionID string) ([]*RegionalGateway, error) {
	return s.gatewayRepo.ListByRegion(regionID)
}

func (s *MultiRegionService) PlaceResource(resourceID, regionID string, isReplicated bool, replicas []string) (*RegionalResourcePlacement, error) {
	placement := &RegionalResourcePlacement{
		ResourceID:   resourceID,
		RegionID:     regionID,
		IsReplicated: isReplicated,
		Replicas:     replicas,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.placementRepo.Create(placement); err != nil {
		return nil, err
	}

	return placement, nil
}

func (s *MultiRegionService) GetResourcePlacement(resourceID string) (*RegionalResourcePlacement, error) {
	return s.placementRepo.GetByResource(resourceID)
}

func (s *MultiRegionService) ListResourcePlacements(regionID string) ([]*RegionalResourcePlacement, error) {
	return s.placementRepo.ListByRegion(regionID)
}

func (s *MultiRegionService) EvaluateRegionCondition(tenantID, regionCode string) bool {
	regions, err := s.regionRepo.ListByTenant(tenantID)
	if err != nil {
		return false
	}

	for _, region := range regions {
		if region.Code == regionCode && region.Status == "active" {
			return true
		}
	}

	return false
}

func (s *MultiRegionService) UpdateRegionHealth(regionID string, activeConns int) (*RegionHealthStatus, error) {
	health, err := s.healthRepo.GetByRegion(regionID)
	if err != nil {
		health = &RegionHealthStatus{
			RegionID: regionID,
		}
	}

	health.Status = "healthy"
	health.ActiveConns = activeConns
	health.Healthy = true
	health.LastCheck = time.Now()

	if err := s.healthRepo.Create(health); err != nil {
		s.healthRepo.Update(health)
	}

	return health, nil
}

func (s *MultiRegionService) GetRegionHealth(regionID string) (*RegionHealthStatus, error) {
	return s.healthRepo.GetByRegion(regionID)
}

func (s *MultiRegionService) SelectBestRegion(tenantID string) (*Region, error) {
	regions, err := s.regionRepo.ListByTenant(tenantID)
	if err != nil {
		return nil, err
	}

	for _, region := range regions {
		if region.IsPrimary && region.Status == "active" {
			return region, nil
		}
	}

	if len(regions) > 0 {
		return regions[0], nil
	}

	return nil, nil
}

func (s *MultiRegionService) DeleteRegion(id string) error {
	return s.regionRepo.Delete(id)
}

func (s *MultiRegionService) DeleteGateway(id string) error {
	return s.gatewayRepo.Delete(id)
}

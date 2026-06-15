package identity

import (
	"time"
)

type GeoLocationContext struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	Country     string    `json:"country"`
	Region      string    `json:"region"`
	City        string    `json:"city"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	ISP         string    `json:"isp"`
	IsVPN       bool      `json:"is_vpn"`
	IsTor       bool      `json:"is_tor"`
	RiskScore   int       `json:"risk_score"`
	CreatedAt   time.Time `json:"created_at"`
}

type GeoPolicyCondition struct {
	ID           string   `json:"id"`
	TenantID     string   `json:"tenant_id"`
	Type         string   `json:"type"`
	Allowlist    []string `json:"allowlist,omitempty"`
	Denylist     []string `json:"denylist,omitempty"`
	RiskThreshold int     `json:"risk_threshold,omitempty"`
}

type ImpossibleTravelEvent struct {
	ID           string    `json:"id"`
	TenantID     string    `json:"tenant_id"`
	UserID       string    `json:"user_id"`
	Location1    string    `json:"location1"`
	Location2    string    `json:"location2"`
	TimeDiff     int       `json:"time_diff_minutes"`
	Distance     float64   `json:"distance_km"`
	IsImpossible bool      `json:"is_impossible"`
	CreatedAt    time.Time `json:"created_at"`
}

type GeoLocationRepository interface {
	Create(g *GeoLocationContext) error
	GetByID(id string) (*GeoLocationContext, error)
	ListByTenant(tenantID string) ([]*GeoLocationContext, error)
}

type GeoPolicyConditionRepository interface {
	Create(c *GeoPolicyCondition) error
	GetByID(id string) (*GeoPolicyCondition, error)
	Update(c *GeoPolicyCondition) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*GeoPolicyCondition, error)
}

type GeoAccessService struct {
	geoRepo     GeoLocationRepository
	policyRepo  GeoPolicyConditionRepository
}

func NewGeoAccessService(gr GeoLocationRepository, pr GeoPolicyConditionRepository) *GeoAccessService {
	return &GeoAccessService{
		geoRepo:    gr,
		policyRepo: pr,
	}
}

func (s *GeoAccessService) EvaluateCountryAllowlist(tenantID, country string) (bool, error) {
	conditions, err := s.policyRepo.ListByTenant(tenantID)
	if err != nil {
		return false, err
	}

	for _, condition := range conditions {
		if condition.Type == "country_allowlist" {
			for _, allowed := range condition.Allowlist {
				if allowed == country {
					return true, nil
				}
			}
			return false, nil
		}
	}

	return true, nil
}

func (s *GeoAccessService) EvaluateCountryDenylist(tenantID, country string) (bool, error) {
	conditions, err := s.policyRepo.ListByTenant(tenantID)
	if err != nil {
		return false, err
	}

	for _, condition := range conditions {
		if condition.Type == "country_denylist" {
			for _, denied := range condition.Denylist {
				if denied == country {
					return false, nil
				}
			}
		}
	}

	return true, nil
}

func (s *GeoAccessService) EvaluateRegionRisk(tenantID, country, region string) (bool, int, error) {
	conditions, err := s.policyRepo.ListByTenant(tenantID)
	if err != nil {
		return false, 0, err
	}

	for _, condition := range conditions {
		if condition.Type == "region_risk" {
			for _, denied := range condition.Denylist {
				if denied == country || denied == region {
					return false, 100, nil
				}
			}
		}
	}

	return true, 0, nil
}

func (s *GeoAccessService) DetectImpossibleTravel(userID string, locations []string) (*ImpossibleTravelEvent, error) {
	if len(locations) < 2 {
		return nil, nil
	}

	event := &ImpossibleTravelEvent{
		TenantID:     "default",
		UserID:       userID,
		Location1:    locations[0],
		Location2:    locations[1],
		TimeDiff:     30,
		Distance:     10000,
		IsImpossible: true,
		CreatedAt:    time.Now(),
	}

	return event, nil
}

func (s *GeoAccessService) CreateGeoContext(tenantID, country, region, city, isp string, isVPN, isTor bool) (*GeoLocationContext, error) {
	riskScore := 0
	if isVPN {
		riskScore += 30
	}
	if isTor {
		riskScore += 50
	}

	geo := &GeoLocationContext{
		TenantID:  tenantID,
		Country:   country,
		Region:    region,
		City:      city,
		ISP:       isp,
		IsVPN:     isVPN,
		IsTor:     isTor,
		RiskScore: riskScore,
		CreatedAt: time.Now(),
	}

	if err := s.geoRepo.Create(geo); err != nil {
		return nil, err
	}

	return geo, nil
}

func (s *GeoAccessService) GetGeoContext(id string) (*GeoLocationContext, error) {
	return s.geoRepo.GetByID(id)
}

func (s *GeoAccessService) ListGeoContexts(tenantID string) ([]*GeoLocationContext, error) {
	return s.geoRepo.ListByTenant(tenantID)
}

func (s *GeoAccessService) CreatePolicyCondition(tenantID, conditionType string, allowlist, denylist []string, riskThreshold int) (*GeoPolicyCondition, error) {
	condition := &GeoPolicyCondition{
		TenantID:      tenantID,
		Type:          conditionType,
		Allowlist:     allowlist,
		Denylist:      denylist,
		RiskThreshold: riskThreshold,
	}

	if err := s.policyRepo.Create(condition); err != nil {
		return nil, err
	}

	return condition, nil
}

func (s *GeoAccessService) GetPolicyCondition(id string) (*GeoPolicyCondition, error) {
	return s.policyRepo.GetByID(id)
}

func (s *GeoAccessService) ListPolicyConditions(tenantID string) ([]*GeoPolicyCondition, error) {
	return s.policyRepo.ListByTenant(tenantID)
}

func (s *GeoAccessService) UpdatePolicyCondition(id string, allowlist, denylist []string, riskThreshold int) error {
	condition, err := s.policyRepo.GetByID(id)
	if err != nil {
		return err
	}

	condition.Allowlist = allowlist
	condition.Denylist = denylist
	condition.RiskThreshold = riskThreshold

	return s.policyRepo.Update(condition)
}

func (s *GeoAccessService) DeletePolicyCondition(id string) error {
	return s.policyRepo.Delete(id)
}

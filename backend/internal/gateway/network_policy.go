package gateway

import (
	"time"
)

type NetworkZone struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	Name        string    `json:"name"`
	CIDR        string    `json:"cidr"`
	ZoneType    string    `json:"zone_type"`
	Sensitivity string    `json:"sensitivity"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NetworkContext struct {
	SourceIP      string `json:"source_ip"`
	SourceZone    string `json:"source_zone"`
	DestIP        string `json:"dest_ip"`
	DestZone      string `json:"dest_zone"`
	IsRiskyPath   bool   `json:"is_risky_path"`
	RiskScore     int    `json:"risk_score"`
}

type NetworkPolicyService struct {
	zones []*NetworkZone
}

func NewNetworkPolicyService() *NetworkPolicyService {
	return &NetworkPolicyService{
		zones: make([]*NetworkZone, 0),
	}
}

func (s *NetworkPolicyService) RegisterZone(zone *NetworkZone) {
	zone.CreatedAt = time.Now()
	zone.UpdatedAt = time.Now()
	s.zones = append(s.zones, zone)
}

func (s *NetworkPolicyService) GetZoneByID(id string) (*NetworkZone, bool) {
	for _, z := range s.zones {
		if z.ID == id {
			return z, true
		}
	}
	return nil, false
}

func (s *NetworkPolicyService) EvaluateNetworkContext(sourceIP, destIP string) *NetworkContext {
	ctx := &NetworkContext{
		SourceIP:    sourceIP,
		DestIP:      destIP,
		IsRiskyPath: false,
		RiskScore:   0,
	}

	for _, zone := range s.zones {
		if zone.CIDR != "" {
			if sourceIP != "" {
				ctx.SourceZone = zone.Name
			}
			if destIP != "" {
				ctx.DestZone = zone.Name
			}
		}
	}

	if ctx.SourceZone != ctx.DestZone && ctx.SourceZone != "" && ctx.DestZone != "" {
		ctx.IsRiskyPath = true
		ctx.RiskScore = 20
	}

	return ctx
}

func (s *NetworkPolicyService) GetRecommendation(ctx *NetworkContext) string {
	if ctx.RiskScore > 50 {
		return "block"
	}
	if ctx.RiskScore > 20 {
		return "require_mfa"
	}
	return "allow"
}

func (s *NetworkPolicyService) ListZones(tenantID string) []*NetworkZone {
	result := make([]*NetworkZone, 0)
	for _, z := range s.zones {
		if z.TenantID == tenantID {
			result = append(result, z)
		}
	}
	return result
}

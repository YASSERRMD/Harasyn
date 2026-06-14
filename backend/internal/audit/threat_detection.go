package audit

import (
	"time"
)

type ThreatRule struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	RuleType    string    `json:"rule_type"`
	Condition   string    `json:"condition"`
	Severity    string    `json:"severity"`
	Enabled     bool      `json:"enabled"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ThreatIncident struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	RuleID      string    `json:"rule_id"`
	UserID      string    `json:"user_id,omitempty"`
	DeviceID    string    `json:"device_id,omitempty"`
	Severity    string    `json:"severity"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	ResolvedAt  *time.Time `json:"resolved_at,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type ThreatRuleRepository interface {
	Create(r *ThreatRule) error
	GetByID(id string) (*ThreatRule, error)
	Update(r *ThreatRule) error
	ListEnabled(tenantID string) ([]*ThreatRule, error)
}

type ThreatIncidentRepository interface {
	Create(i *ThreatIncident) error
	GetByID(id string) (*ThreatIncident, error)
	Update(i *ThreatIncident) error
	ListByTenant(tenantID string, limit int) ([]*ThreatIncident, error)
}

type ThreatDetectionService struct {
	ruleRepo      ThreatRuleRepository
	incidentRepo  ThreatIncidentRepository
}

func NewThreatDetectionService(rr ThreatRuleRepository, ir ThreatIncidentRepository) *ThreatDetectionService {
	return &ThreatDetectionService{
		ruleRepo:     rr,
		incidentRepo: ir,
	}
}

type DetectionEvent struct {
	TenantID  string            `json:"tenant_id"`
	UserID    string            `json:"user_id,omitempty"`
	DeviceID  string            `json:"device_id,omitempty"`
	EventType string            `json:"event_type"`
	Details   map[string]string `json:"details,omitempty"`
}

func (s *ThreatDetectionService) EvaluateEvent(event DetectionEvent) ([]*ThreatIncident, error) {
	rules, err := s.ruleRepo.ListEnabled(event.TenantID)
	if err != nil {
		return nil, err
	}

	incidents := make([]*ThreatIncident, 0)

	for _, rule := range rules {
		if s.matchesRule(rule, event) {
			incident := &ThreatIncident{
				TenantID:    event.TenantID,
				RuleID:      rule.ID,
				UserID:      event.UserID,
				DeviceID:    event.DeviceID,
				Severity:    rule.Severity,
				Description: rule.Description,
				Status:      "open",
				CreatedAt:   time.Now(),
			}

			if err := s.incidentRepo.Create(incident); err != nil {
				continue
			}

			incidents = append(incidents, incident)
		}
	}

	return incidents, nil
}

func (s *ThreatDetectionService) matchesRule(rule *ThreatRule, event DetectionEvent) bool {
	switch rule.RuleType {
	case "repeated_denial":
		return event.EventType == "authorization_denied"
	case "new_device":
		return event.EventType == "device_registered"
	case "unusual_time":
		return event.EventType == "access_outside_hours"
	case "high_risk_country":
		return event.EventType == "access_high_risk_country"
	default:
		return false
	}
}

func (s *ThreatDetectionService) GetIncidents(tenantID string, limit int) ([]*ThreatIncident, error) {
	return s.incidentRepo.ListByTenant(tenantID, limit)
}

func (s *ThreatDetectionService) ResolveIncident(incidentID string) error {
	incident, err := s.incidentRepo.GetByID(incidentID)
	if err != nil {
		return err
	}

	incident.Status = "resolved"
	now := time.Now()
	incident.ResolvedAt = &now

	return s.incidentRepo.Update(incident)
}

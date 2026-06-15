package gateway

import (
	"time"
)

type DLPPolicy struct {
	ID          string            `json:"id"`
	TenantID    string            `json:"tenant_id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Rules       []DLPRule         `json:"rules"`
	Status      string            `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type DLPRule struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Action      string   `json:"action"`
	Conditions  []string `json:"conditions"`
	Sensitivity string   `json:"sensitivity"`
}

type DLPIncident struct {
	ID         string    `json:"id"`
	TenantID   string    `json:"tenant_id"`
	PolicyID   string    `json:"policy_id"`
	UserID     string    `json:"user_id"`
	ResourceID string    `json:"resource_id"`
	Type       string    `json:"type"`
	Content    string    `json:"content"`
	RiskScore  int       `json:"risk_score"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type DLPDecision struct {
	Allowed    bool     `json:"allowed"`
	PolicyID   string   `json:"policy_id,omitempty"`
	RuleID     string   `json:"rule_id,omitempty"`
	RiskScore  int      `json:"risk_score"`
	Reason     string   `json:"reason,omitempty"`
}

type ContentInspectionResult struct {
	HasSensitiveContent bool     `json:"has_sensitive_content"`
	Keywords           []string `json:"keywords,omitempty"`
	Categories         []string `json:"categories,omitempty"`
	RiskScore          int      `json:"risk_score"`
}

type DLPPolicyRepository interface {
	Create(p *DLPPolicy) error
	GetByID(id string) (*DLPPolicy, error)
	Update(p *DLPPolicy) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*DLPPolicy, error)
}

type DLPIncidentRepository interface {
	Create(i *DLPIncident) error
	GetByID(id string) (*DLPIncident, error)
	Update(i *DLPIncident) error
	ListByTenant(tenantID string) ([]*DLPIncident, error)
	ListByUser(userID string) ([]*DLPIncident, error)
}

type DLPService struct {
	policyRepo  DLPPolicyRepository
	incidentRepo DLPIncidentRepository
}

func NewDLPService(pr DLPPolicyRepository, ir DLPIncidentRepository) *DLPService {
	return &DLPService{
		policyRepo:   pr,
		incidentRepo: ir,
	}
}

func (s *DLPService) CreatePolicy(tenantID, name, description string, rules []DLPRule) (*DLPPolicy, error) {
	policy := &DLPPolicy{
		TenantID:    tenantID,
		Name:        name,
		Description: description,
		Rules:       rules,
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.policyRepo.Create(policy); err != nil {
		return nil, err
	}

	return policy, nil
}

func (s *DLPService) GetPolicy(id string) (*DLPPolicy, error) {
	return s.policyRepo.GetByID(id)
}

func (s *DLPService) ListPolicies(tenantID string) ([]*DLPPolicy, error) {
	return s.policyRepo.ListByTenant(tenantID)
}

func (s *DLPService) InspectContent(content string) *ContentInspectionResult {
	sensitiveKeywords := []string{
		"password", "secret", "api_key", "credit_card",
		"social_security", "ssn", "confidential",
	}

	result := &ContentInspectionResult{
		HasSensitiveContent: false,
		Keywords:           []string{},
		Categories:         []string{},
		RiskScore:          0,
	}

	for _, keyword := range sensitiveKeywords {
		if contains(content, keyword) {
			result.HasSensitiveContent = true
			result.Keywords = append(result.Keywords, keyword)
			result.RiskScore += 25
		}
	}

	if result.RiskScore > 100 {
		result.RiskScore = 100
	}

	return result
}

func (s *DLPService) DetectFileDownloadRisk(userID, fileName string, fileSize int) int {
	riskScore := 0

	if fileSize > 10*1024*1024 {
		riskScore += 30
	}

	sensitiveExtensions := []string{".csv", ".xlsx", ".pdf", ".docx"}
	for _, ext := range sensitiveExtensions {
		if contains(fileName, ext) {
			riskScore += 20
			break
		}
	}

	if riskScore > 100 {
		riskScore = 100
	}

	return riskScore
}

func (s *DLPService) DetectCopyPasteRisk(content string) int {
	result := s.InspectContent(content)
	return result.RiskScore
}

func (s *DLPService) EvaluateDLPPolicy(tenantID, userID, resourceID, action, content string) (*DLPDecision, error) {
	policies, err := s.policyRepo.ListByTenant(tenantID)
	if err != nil {
		return nil, err
	}

	for _, policy := range policies {
		if policy.Status != "active" {
			continue
		}

		for _, rule := range policy.Rules {
			if rule.Action == action {
				inspection := s.InspectContent(content)
				if inspection.HasSensitiveContent {
					return &DLPDecision{
						Allowed:   false,
						PolicyID:  policy.ID,
						RuleID:    rule.ID,
						RiskScore: inspection.RiskScore,
						Reason:    "Content contains sensitive data",
					}, nil
				}
			}
		}
	}

	return &DLPDecision{
		Allowed:   true,
		RiskScore: 0,
	}, nil
}

func (s *DLPService) CreateIncident(tenantID, policyID, userID, resourceID, incidentType, content string, riskScore int) (*DLPIncident, error) {
	incident := &DLPIncident{
		TenantID:   tenantID,
		PolicyID:   policyID,
		UserID:     userID,
		ResourceID: resourceID,
		Type:       incidentType,
		Content:    content,
		RiskScore:  riskScore,
		Status:     "open",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.incidentRepo.Create(incident); err != nil {
		return nil, err
	}

	return incident, nil
}

func (s *DLPService) GetIncident(id string) (*DLPIncident, error) {
	return s.incidentRepo.GetByID(id)
}

func (s *DLPService) ListIncidents(tenantID string) ([]*DLPIncident, error) {
	return s.incidentRepo.ListByTenant(tenantID)
}

func (s *DLPService) UpdateIncidentStatus(id, status string) error {
	incident, err := s.incidentRepo.GetByID(id)
	if err != nil {
		return err
	}

	incident.Status = status
	incident.UpdatedAt = time.Now()

	return s.incidentRepo.Update(incident)
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && (s[0:len(substr)] == substr || contains(s[1:], substr)))
}

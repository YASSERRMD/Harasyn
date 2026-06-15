package device

import (
	"time"
)

type TrustScoreExplanation struct {
	ID           string              `json:"id"`
	TenantID     string              `json:"tenant_id"`
	EntityType   string              `json:"entity_type"`
	EntityID     string              `json:"entity_id"`
	Score        int                 `json:"score"`
	Factors      []TrustFactor       `json:"factors"`
	Contributions []FactorContribution `json:"contributions"`
	GeneratedAt  time.Time           `json:"generated_at"`
}

type TrustFactor struct {
	Name        string  `json:"name"`
	Value       int     `json:"value"`
	Weight      float64 `json:"weight"`
	Contribution int    `json:"contribution"`
}

type FactorContribution struct {
	Factor      string  `json:"factor"`
	Contribution int    `json:"contribution"`
	Percentage  float64 `json:"percentage"`
}

type TrustScoreHistory struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	EntityType string  `json:"entity_type"`
	EntityID  string    `json:"entity_id"`
	Score     int       `json:"score"`
	Factors   []TrustFactor `json:"factors"`
	RecordedAt time.Time `json:"recorded_at"`
}

type TrustScoreExplanationRepository interface {
	Create(e *TrustScoreExplanation) error
	GetByID(id string) (*TrustScoreExplanation, error)
	GetByEntity(entityType, entityID string) (*TrustScoreExplanation, error)
	ListByEntity(entityType, entityID string) ([]*TrustScoreExplanation, error)
}

type TrustScoreHistoryRepository interface {
	Create(h *TrustScoreHistory) error
	GetByID(id string) (*TrustScoreHistory, error)
	ListByEntity(entityType, entityID string) ([]*TrustScoreHistory, error)
}

type TrustExplainService struct {
	explanationRepo TrustScoreExplanationRepository
	historyRepo     TrustScoreHistoryRepository
}

func NewTrustExplainService(er TrustScoreExplanationRepository, hr TrustScoreHistoryRepository) *TrustExplainService {
	return &TrustExplainService{
		explanationRepo: er,
		historyRepo:     hr,
	}
}

func (s *TrustExplainService) ExplainDeviceTrust(tenantID, deviceID string, deviceScore int, factors map[string]int) (*TrustScoreExplanation, error) {
	trustFactors := make([]TrustFactor, 0)
	totalContribution := 0

	for name, value := range factors {
		weight := 1.0
		contribution := int(float64(value) * weight)
		totalContribution += contribution

		trustFactors = append(trustFactors, TrustFactor{
			Name:         name,
			Value:        value,
			Weight:       weight,
			Contribution: contribution,
		})
	}

	contributions := make([]FactorContribution, 0)
	for _, factor := range trustFactors {
		percentage := float64(factor.Contribution) / float64(totalContribution) * 100
		contributions = append(contributions, FactorContribution{
			Factor:       factor.Name,
			Contribution: factor.Contribution,
			Percentage:   percentage,
		})
	}

	explanation := &TrustScoreExplanation{
		TenantID:      tenantID,
		EntityType:    "device",
		EntityID:      deviceID,
		Score:         deviceScore,
		Factors:       trustFactors,
		Contributions: contributions,
		GeneratedAt:   time.Now(),
	}

	if err := s.explanationRepo.Create(explanation); err != nil {
		return nil, err
	}

	history := &TrustScoreHistory{
		TenantID:   tenantID,
		EntityType: "device",
		EntityID:   deviceID,
		Score:      deviceScore,
		Factors:    trustFactors,
		RecordedAt: time.Now(),
	}
	s.historyRepo.Create(history)

	return explanation, nil
}

func (s *TrustExplainService) ExplainUserTrust(tenantID, userID string, userScore int, factors map[string]int) (*TrustScoreExplanation, error) {
	trustFactors := make([]TrustFactor, 0)
	totalContribution := 0

	for name, value := range factors {
		weight := 1.0
		contribution := int(float64(value) * weight)
		totalContribution += contribution

		trustFactors = append(trustFactors, TrustFactor{
			Name:         name,
			Value:        value,
			Weight:       weight,
			Contribution: contribution,
		})
	}

	contributions := make([]FactorContribution, 0)
	for _, factor := range trustFactors {
		percentage := float64(factor.Contribution) / float64(totalContribution) * 100
		contributions = append(contributions, FactorContribution{
			Factor:       factor.Name,
			Contribution: factor.Contribution,
			Percentage:   percentage,
		})
	}

	explanation := &TrustScoreExplanation{
		TenantID:      tenantID,
		EntityType:    "user",
		EntityID:      userID,
		Score:         userScore,
		Factors:       trustFactors,
		Contributions: contributions,
		GeneratedAt:   time.Now(),
	}

	if err := s.explanationRepo.Create(explanation); err != nil {
		return nil, err
	}

	history := &TrustScoreHistory{
		TenantID:   tenantID,
		EntityType: "user",
		EntityID:   userID,
		Score:      userScore,
		Factors:    trustFactors,
		RecordedAt: time.Now(),
	}
	s.historyRepo.Create(history)

	return explanation, nil
}

func (s *TrustExplainService) ExplainSessionTrust(tenantID, sessionID string, sessionScore int, factors map[string]int) (*TrustScoreExplanation, error) {
	trustFactors := make([]TrustFactor, 0)
	totalContribution := 0

	for name, value := range factors {
		weight := 1.0
		contribution := int(float64(value) * weight)
		totalContribution += contribution

		trustFactors = append(trustFactors, TrustFactor{
			Name:         name,
			Value:        value,
			Weight:       weight,
			Contribution: contribution,
		})
	}

	contributions := make([]FactorContribution, 0)
	for _, factor := range trustFactors {
		percentage := float64(factor.Contribution) / float64(totalContribution) * 100
		contributions = append(contributions, FactorContribution{
			Factor:       factor.Name,
			Contribution: factor.Contribution,
			Percentage:   percentage,
		})
	}

	explanation := &TrustScoreExplanation{
		TenantID:      tenantID,
		EntityType:    "session",
		EntityID:      sessionID,
		Score:         sessionScore,
		Factors:       trustFactors,
		Contributions: contributions,
		GeneratedAt:   time.Now(),
	}

	if err := s.explanationRepo.Create(explanation); err != nil {
		return nil, err
	}

	history := &TrustScoreHistory{
		TenantID:   tenantID,
		EntityType: "session",
		EntityID:   sessionID,
		Score:      sessionScore,
		Factors:    trustFactors,
		RecordedAt: time.Now(),
	}
	s.historyRepo.Create(history)

	return explanation, nil
}

func (s *TrustExplainService) GetExplanation(entityType, entityID string) (*TrustScoreExplanation, error) {
	return s.explanationRepo.GetByEntity(entityType, entityID)
}

func (s *TrustExplainService) GetHistory(entityType, entityID string) ([]*TrustScoreHistory, error) {
	return s.historyRepo.ListByEntity(entityType, entityID)
}

func (s *TrustExplainService) GetRiskFactorContribution(entityType, entityID string) ([]FactorContribution, error) {
	explanation, err := s.explanationRepo.GetByEntity(entityType, entityID)
	if err != nil {
		return nil, err
	}

	return explanation.Contributions, nil
}

package gateway

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type Agent struct {
	ID           string    `json:"id"`
	TenantID     string    `json:"tenant_id"`
	Name         string    `json:"name"`
	Version      string    `json:"version"`
	OS           string    `json:"os"`
	IPAddress    string    `json:"ip_address"`
	Status       string    `json:"status"`
	TrustScore   int       `json:"trust_score"`
	LastHeartbeat time.Time `json:"last_heartbeat"`
	EnrolledAt   time.Time `json:"enrolled_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AgentEnrollmentToken struct {
	ID        string    `json:"id"`
	Token     string    `json:"token"`
	AgentID   string    `json:"agent_id"`
	TenantID  string    `json:"tenant_id"`
	Used      bool      `json:"used"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type AgentResourceMapping struct {
	ID        string `json:"id"`
	AgentID   string `json:"agent_id"`
	ResourceID string `json:"resource_id"`
	TenantID  string `json:"tenant_id"`
	Status    string `json:"status"`
}

type AgentRepository interface {
	Create(a *Agent) error
	GetByID(id string) (*Agent, error)
	Update(a *Agent) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*Agent, error)
}

type AgentService struct {
	repo AgentRepository
}

func NewAgentService(repo AgentRepository) *AgentService {
	return &AgentService{repo: repo}
}

type EnrollAgentRequest struct {
	TenantID string `json:"tenant_id"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	OS       string `json:"os"`
	IP       string `json:"ip_address"`
}

func (s *AgentService) EnrollAgent(req EnrollAgentRequest) (*Agent, error) {
	agent := &Agent{
		TenantID:   req.TenantID,
		Name:       req.Name,
		Version:    req.Version,
		OS:         req.OS,
		IPAddress:  req.IP,
		Status:     "active",
		TrustScore: 50,
		EnrolledAt: time.Now(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.repo.Create(agent); err != nil {
		return nil, err
	}

	return agent, nil
}

func (s *AgentService) GetAgent(id string) (*Agent, error) {
	return s.repo.GetByID(id)
}

func (s *AgentService) ListAgents(tenantID string) ([]*Agent, error) {
	return s.repo.ListByTenant(tenantID)
}

func (s *AgentService) Heartbeat(agentID string) error {
	agent, err := s.repo.GetByID(agentID)
	if err != nil {
		return err
	}

	agent.LastHeartbeat = time.Now()
	agent.UpdatedAt = time.Now()
	return s.repo.Update(agent)
}

func GenerateEnrollmentToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

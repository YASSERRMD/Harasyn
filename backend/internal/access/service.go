package access

import (
	"fmt"
	"time"
)

type Service struct {
	repo         Repository
	approvalRepo ApprovalRepository
}

func NewService(repo Repository, approvalRepo ApprovalRepository) *Service {
	return &Service{
		repo:         repo,
		approvalRepo: approvalRepo,
	}
}

type CreateRequestRequest struct {
	TenantID        string `json:"tenant_id"`
	UserID          string `json:"user_id"`
	DeviceID        string `json:"device_id"`
	ResourceID      string `json:"resource_id"`
	RequestType     string `json:"request_type"`
	Justification   string `json:"justification,omitempty"`
	DurationMinutes int    `json:"duration_minutes"`
}

func (s *Service) CreateRequest(req CreateRequestRequest) (*AccessRequest, error) {
	duration := time.Duration(req.DurationMinutes) * time.Minute
	if duration == 0 {
		duration = 60 * time.Minute
	}

	accessReq := &AccessRequest{
		TenantID:        req.TenantID,
		UserID:          req.UserID,
		DeviceID:        req.DeviceID,
		ResourceID:      req.ResourceID,
		RequestType:     req.RequestType,
		Justification:   req.Justification,
		Status:          "pending",
		DurationMinutes: req.DurationMinutes,
		RequestedAt:     time.Now(),
	}

	if req.RequestType == "emergency" {
		accessReq.DurationMinutes = 30
	}

	if err := s.repo.Create(accessReq); err != nil {
		return nil, fmt.Errorf("failed to create access request: %w", err)
	}

	return accessReq, nil
}

func (s *Service) GetRequest(id string) (*AccessRequest, error) {
	return s.repo.GetByID(id)
}

func (s *Service) ListPendingRequests(tenantID string) ([]*AccessRequest, error) {
	return s.repo.ListPendingByTenant(tenantID)
}

func (s *Service) ListRequestsByTenant(tenantID string) ([]*AccessRequest, error) {
	return s.repo.ListByTenant(tenantID)
}

type ApproveRequestRequest struct {
	RequestID  string `json:"request_id"`
	ReviewerID string `json:"reviewer_id"`
	Reason     string `json:"reason,omitempty"`
}

func (s *Service) ApproveRequest(req ApproveRequestRequest) (*ApprovalDecision, error) {
	accessReq, err := s.repo.GetByID(req.RequestID)
	if err != nil {
		return nil, fmt.Errorf("request not found: %w", err)
	}

	if accessReq.Status != "pending" {
		return nil, fmt.Errorf("request is not pending")
	}

	decision := &ApprovalDecision{
		RequestID:  req.RequestID,
		ReviewerID: req.ReviewerID,
		Decision:   "approved",
		Reason:     req.Reason,
		DecidedAt:  time.Now(),
	}

	if err := s.approvalRepo.Create(decision); err != nil {
		return nil, fmt.Errorf("failed to create approval decision: %w", err)
	}

	accessReq.Status = "approved"
	now := time.Now()
	accessReq.ResolvedAt = &now
	accessReq.ResolvedBy = &req.ReviewerID

	expiresAt := now.Add(time.Duration(accessReq.DurationMinutes) * time.Minute)
	accessReq.ExpiresAt = &expiresAt

	if err := s.repo.Update(accessReq); err != nil {
		return nil, fmt.Errorf("failed to update request: %w", err)
	}

	return decision, nil
}

type RejectRequestRequest struct {
	RequestID  string `json:"request_id"`
	ReviewerID string `json:"reviewer_id"`
	Reason     string `json:"reason"`
}

func (s *Service) RejectRequest(req RejectRequestRequest) (*ApprovalDecision, error) {
	accessReq, err := s.repo.GetByID(req.RequestID)
	if err != nil {
		return nil, fmt.Errorf("request not found: %w", err)
	}

	if accessReq.Status != "pending" {
		return nil, fmt.Errorf("request is not pending")
	}

	decision := &ApprovalDecision{
		RequestID:  req.RequestID,
		ReviewerID: req.ReviewerID,
		Decision:   "rejected",
		Reason:     req.Reason,
		DecidedAt:  time.Now(),
	}

	if err := s.approvalRepo.Create(decision); err != nil {
		return nil, fmt.Errorf("failed to create approval decision: %w", err)
	}

	accessReq.Status = "rejected"
	now := time.Now()
	accessReq.ResolvedAt = &now
	accessReq.ResolvedBy = &req.ReviewerID

	if err := s.repo.Update(accessReq); err != nil {
		return nil, fmt.Errorf("failed to update request: %w", err)
	}

	return decision, nil
}

func (s *Service) GetDecisions(requestID string) ([]*ApprovalDecision, error) {
	return s.approvalRepo.GetByRequestID(requestID)
}

package access

import (
	"time"
)

type AccessReviewCampaign struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AccessReviewItem struct {
	ID          string    `json:"id"`
	CampaignID  string    `json:"campaign_id"`
	UserID      string    `json:"user_id"`
	ResourceID  string    `json:"resource_id"`
	ReviewerID  string    `json:"reviewer_id"`
	Status      string    `json:"status"`
	Decision    string    `json:"decision,omitempty"`
	Reason      string    `json:"reason,omitempty"`
	ReviewedAt  *time.Time `json:"reviewed_at,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type AccessReviewCampaignRepository interface {
	Create(c *AccessReviewCampaign) error
	GetByID(id string) (*AccessReviewCampaign, error)
	Update(c *AccessReviewCampaign) error
	ListByTenant(tenantID string) ([]*AccessReviewCampaign, error)
}

type AccessReviewItemRepository interface {
	Create(i *AccessReviewItem) error
	GetByID(id string) (*AccessReviewItem, error)
	Update(i *AccessReviewItem) error
	ListByCampaign(campaignID string) ([]*AccessReviewItem, error)
	ListByReviewer(reviewerID string) ([]*AccessReviewItem, error)
}

type AccessReviewService struct {
	campaignRepo AccessReviewCampaignRepository
	itemRepo     AccessReviewItemRepository
}

func NewAccessReviewService(cr AccessReviewCampaignRepository, ir AccessReviewItemRepository) *AccessReviewService {
	return &AccessReviewService{
		campaignRepo: cr,
		itemRepo:     ir,
	}
}

type CreateCampaignRequest struct {
	TenantID    string `json:"tenant_id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CreatedBy   string `json:"created_by"`
	DurationDays int   `json:"duration_days"`
}

func (s *AccessReviewService) CreateCampaign(req CreateCampaignRequest) (*AccessReviewCampaign, error) {
	campaign := &AccessReviewCampaign{
		TenantID:    req.TenantID,
		Name:        req.Name,
		Description: req.Description,
		Status:      "active",
		StartDate:   time.Now(),
		EndDate:     time.Now().AddDate(0, 0, req.DurationDays),
		CreatedBy:   req.CreatedBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.campaignRepo.Create(campaign); err != nil {
		return nil, err
	}

	return campaign, nil
}

func (s *AccessReviewService) GetCampaign(id string) (*AccessReviewCampaign, error) {
	return s.campaignRepo.GetByID(id)
}

func (s *AccessReviewService) ListCampaigns(tenantID string) ([]*AccessReviewCampaign, error) {
	return s.campaignRepo.ListByTenant(tenantID)
}

type ReviewAccessRequest struct {
	ItemID   string `json:"item_id"`
	Decision string `json:"decision"`
	Reason   string `json:"reason,omitempty"`
}

func (s *AccessReviewService) ReviewAccess(req ReviewAccessRequest) (*AccessReviewItem, error) {
	item, err := s.itemRepo.GetByID(req.ItemID)
	if err != nil {
		return nil, err
	}

	item.Decision = req.Decision
	item.Reason = req.Reason
	item.Status = "reviewed"
	now := time.Now()
	item.ReviewedAt = &now

	if err := s.itemRepo.Update(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *AccessReviewService) GetReviewProgress(campaignID string) (int, int, error) {
	items, err := s.itemRepo.ListByCampaign(campaignID)
	if err != nil {
		return 0, 0, err
	}

	total := len(items)
	reviewed := 0
	for _, item := range items {
		if item.Status == "reviewed" {
			reviewed++
		}
	}

	return reviewed, total, nil
}

func (s *AccessReviewService) GetItemsForReviewer(reviewerID string) ([]*AccessReviewItem, error) {
	return s.itemRepo.ListByReviewer(reviewerID)
}

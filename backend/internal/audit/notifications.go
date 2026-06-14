package audit

import (
	"time"
)

type NotificationProvider interface {
	Send(notification *Notification) error
	GetName() string
}

type Notification struct {
	ID        string            `json:"id"`
	TenantID  string            `json:"tenant_id"`
	Type      string            `json:"type"`
	Channel   string            `json:"channel"`
	Subject   string            `json:"subject"`
	Body      string            `json:"body"`
	Recipient string            `json:"recipient"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	Status    string            `json:"status"`
	SentAt    *time.Time        `json:"sent_at,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
}

type NotificationTemplate struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Channel   string `json:"channel"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type NotificationRoute struct {
	ID        string   `json:"id"`
	TenantID  string   `json:"tenant_id"`
	EventType string   `json:"event_type"`
	Channels  []string `json:"channels"`
	Recipients []string `json:"recipients"`
}

type NotificationRepository interface {
	Create(n *Notification) error
	GetByID(id string) (*Notification, error)
	Update(n *Notification) error
	ListByTenant(tenantID string, limit int) ([]*Notification, error)
}

type NotificationService struct {
	repo      NotificationRepository
	providers map[string]NotificationProvider
}

func NewNotificationService(repo NotificationRepository) *NotificationService {
	return &NotificationService{
		repo:      repo,
		providers: make(map[string]NotificationProvider),
	}
}

func (s *NotificationService) RegisterProvider(provider NotificationProvider) {
	s.providers[provider.GetName()] = provider
}

type SendNotificationRequest struct {
	TenantID  string            `json:"tenant_id"`
	Type      string            `json:"type"`
	Channel   string            `json:"channel"`
	Subject   string            `json:"subject"`
	Body      string            `json:"body"`
	Recipient string            `json:"recipient"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

func (s *NotificationService) SendNotification(req SendNotificationRequest) (*Notification, error) {
	notification := &Notification{
		TenantID:  req.TenantID,
		Type:      req.Type,
		Channel:   req.Channel,
		Subject:   req.Subject,
		Body:      req.Body,
		Recipient: req.Recipient,
		Metadata:  req.Metadata,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(notification); err != nil {
		return nil, err
	}

	provider, ok := s.providers[req.Channel]
	if ok {
		if err := provider.Send(notification); err != nil {
			notification.Status = "failed"
			s.repo.Update(notification)
			return notification, err
		}

		notification.Status = "sent"
		now := time.Now()
		notification.SentAt = &now
		s.repo.Update(notification)
	} else {
		notification.Status = "queued"
		s.repo.Update(notification)
	}

	return notification, nil
}

func (s *NotificationService) GetNotifications(tenantID string, limit int) ([]*Notification, error) {
	return s.repo.ListByTenant(tenantID, limit)
}

type WebhookProvider struct {
	URL    string
	Secret string
}

func (p *WebhookProvider) Send(notification *Notification) error {
	return nil
}

func (p *WebhookProvider) GetName() string {
	return "webhook"
}

type EmailProvider struct {
	SMTPHost string
	SMTPPort int
	Username string
	Password string
}

func (p *EmailProvider) Send(notification *Notification) error {
	return nil
}

func (p *EmailProvider) GetName() string {
	return "email"
}

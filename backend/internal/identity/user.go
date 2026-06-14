package identity

import "time"

type User struct {
	ID          string     `json:"id"`
	TenantID    string     `json:"tenant_id"`
	Email       string     `json:"email"`
	Username    string     `json:"username,omitempty"`
	DisplayName string     `json:"display_name,omitempty"`
	AvatarURL   string     `json:"avatar_url,omitempty"`
	MFAEnabled  bool       `json:"mfa_enabled"`
	MFAMethod   string     `json:"mfa_method,omitempty"`
	Status      string     `json:"status"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Repository interface {
	Create(u *User) error
	GetByID(id string) (*User, error)
	GetByEmail(tenantID, email string) (*User, error)
	Update(u *User) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*User, error)
}

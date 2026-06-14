package tenant

import "time"

type Tenant struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Slug      string            `json:"slug"`
	Domain    string            `json:"domain,omitempty"`
	Settings  map[string]string `json:"settings,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type Repository interface {
	Create(t *Tenant) error
	GetByID(id string) (*Tenant, error)
	GetBySlug(slug string) (*Tenant, error)
	Update(t *Tenant) error
	Delete(id string) error
	List() ([]*Tenant, error)
}

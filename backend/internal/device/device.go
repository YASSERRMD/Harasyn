package device

import "time"

type Device struct {
	ID           string    `json:"id"`
	TenantID     string    `json:"tenant_id"`
	UserID       string    `json:"user_id"`
	Name         string    `json:"name"`
	Fingerprint  string    `json:"fingerprint"`
	OS           string    `json:"os"`
	OSVersion    string    `json:"os_version,omitempty"`
	DeviceType   string    `json:"device_type,omitempty"`
	Manufacturer string    `json:"manufacturer,omitempty"`
	Model        string    `json:"model,omitempty"`
	Status       string    `json:"status"`
	TrustScore   int       `json:"trust_score"`
	LastSeenAt   *time.Time `json:"last_seen_at,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Posture struct {
	ID                string    `json:"id"`
	DeviceID          string    `json:"device_id"`
	Encrypted         bool      `json:"encrypted"`
	Jailbroken        bool      `json:"jailbroken"`
	Rooted            bool      `json:"rooted"`
	Patched           bool      `json:"patched"`
	AntivirusEnabled  bool      `json:"antivirus_enabled"`
	FirewallEnabled   bool      `json:"firewall_enabled"`
	DiskEncrypted     bool      `json:"disk_encrypted"`
	OSPatchLevel      string    `json:"os_patch_level,omitempty"`
	ComplianceStatus  string    `json:"compliance_status"`
	EvaluatedAt       time.Time `json:"evaluated_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type Repository interface {
	Create(d *Device) error
	GetByID(id string) (*Device, error)
	GetByFingerprint(tenantID, fingerprint string) (*Device, error)
	Update(d *Device) error
	Delete(id string) error
	ListByUser(userID string) ([]*Device, error)
	ListByTenant(tenantID string) ([]*Device, error)
	UpdateLastSeen(id string) error
}

type PostureRepository interface {
	Create(p *Posture) error
	GetLatestByDeviceID(deviceID string) (*Posture, error)
	Update(p *Posture) error
}

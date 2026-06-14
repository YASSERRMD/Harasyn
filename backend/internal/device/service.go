package device

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Service struct {
	repo     Repository
	postureRepo PostureRepository
}

func NewService(repo Repository, postureRepo PostureRepository) *Service {
	return &Service{
		repo:      repo,
		postureRepo: postureRepo,
	}
}

type RegisterDeviceRequest struct {
	TenantID     string `json:"tenant_id"`
	UserID       string `json:"user_id"`
	Name         string `json:"name"`
	Fingerprint  string `json:"fingerprint"`
	OS           string `json:"os"`
	OSVersion    string `json:"os_version,omitempty"`
	DeviceType   string `json:"device_type,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Model        string `json:"model,omitempty"`
}

func (s *Service) RegisterDevice(req RegisterDeviceRequest) (*Device, error) {
	existing, _ := s.repo.GetByFingerprint(req.TenantID, req.Fingerprint)
	if existing != nil {
		now := time.Now()
		existing.LastSeenAt = &now
		existing.Name = req.Name
		existing.OS = req.OS
		existing.OSVersion = req.OSVersion
		if err := s.repo.Update(existing); err != nil {
			return nil, fmt.Errorf("failed to update device: %w", err)
		}
		return existing, nil
	}

	device := &Device{
		TenantID:     req.TenantID,
		UserID:       req.UserID,
		Name:         req.Name,
		Fingerprint:  req.Fingerprint,
		OS:           req.OS,
		OSVersion:    req.OSVersion,
		DeviceType:   req.DeviceType,
		Manufacturer: req.Manufacturer,
		Model:        req.Model,
		Status:       "active",
		TrustScore:   50,
	}

	if err := s.repo.Create(device); err != nil {
		return nil, fmt.Errorf("failed to create device: %w", err)
	}

	return device, nil
}

func (s *Service) GetDevice(id string) (*Device, error) {
	return s.repo.GetByID(id)
}

func (s *Service) ListDevicesByUser(userID string) ([]*Device, error) {
	return s.repo.ListByUser(userID)
}

func (s *Service) ListDevicesByTenant(tenantID string) ([]*Device, error) {
	return s.repo.ListByTenant(tenantID)
}

func GenerateFingerprint(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func CalculateTrustScore(p *Posture, device *Device) int {
	score := 50

	if p.Encrypted {
		score += 15
	}
	if !p.Jailbroken && !p.Rooted {
		score += 15
	}
	if p.Patched {
		score += 10
	}
	if p.AntivirusEnabled {
		score += 5
	}
	if p.FirewallEnabled {
		score += 5
	}

	if p.Jailbroken || p.Rooted {
		score -= 40
	}
	if !p.Encrypted {
		score -= 10
	}

	if device.LastSeenAt != nil {
		hoursSinceSeen := time.Since(*device.LastSeenAt).Hours()
		if hoursSinceSeen > 720 {
			score -= 10
		}
	}

	if score < 0 {
		score = 0
	}
	if score > 100 {
		score = 100
	}

	return score
}

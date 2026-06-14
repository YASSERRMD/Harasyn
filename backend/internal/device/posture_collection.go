package device

import (
	"time"
)

type PostureSignal struct {
	ID        string          `json:"id"`
	DeviceID  string          `json:"device_id"`
	SignalType string         `json:"signal_type"`
	Value     string          `json:"value"`
	Category  string          `json:"category"`
	Severity  string          `json:"severity"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	CollectedAt time.Time     `json:"collected_at"`
	CreatedAt time.Time       `json:"created_at"`
}

type PostureSignalType string

const (
	SignalOSVersion       PostureSignalType = "os_version"
	SignalDiskEncryption  PostureSignalType = "disk_encryption"
	SignalAntivirus       PostureSignalType = "antivirus"
	SignalFirewall        PostureSignalType = "firewall"
	SignalJailbreak       PostureSignalType = "jailbreak"
	SignalPatchLevel      PostureSignalType = "patch_level"
)

type PostureSignalRepository interface {
	Create(s *PostureSignal) error
	GetByID(id string) (*PostureSignal, error)
	ListByDevice(deviceID string) ([]*PostureSignal, error)
	ListByDeviceAndType(deviceID string, signalType PostureSignalType) ([]*PostureSignal, error)
}

type PostureCollectionService struct {
	repo PostureSignalRepository
}

func NewPostureCollectionService(repo PostureSignalRepository) *PostureCollectionService {
	return &PostureCollectionService{repo: repo}
}

type SubmitPostureRequest struct {
	DeviceID    string          `json:"device_id"`
	SignalType  PostureSignalType `json:"signal_type"`
	Value       string          `json:"value"`
	Category    string          `json:"category"`
	Severity    string          `json:"severity"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

func (s *PostureCollectionService) SubmitSignal(req SubmitPostureRequest) (*PostureSignal, error) {
	signal := &PostureSignal{
		DeviceID:    req.DeviceID,
		SignalType:  string(req.SignalType),
		Value:       req.Value,
		Category:    req.Category,
		Severity:    req.Severity,
		Metadata:    req.Metadata,
		CollectedAt: time.Now(),
		CreatedAt:   time.Now(),
	}

	if err := s.repo.Create(signal); err != nil {
		return nil, err
	}

	return signal, nil
}

func (s *PostureCollectionService) GetSignalsByDevice(deviceID string) ([]*PostureSignal, error) {
	return s.repo.ListByDevice(deviceID)
}

func (s *PostureCollectionService) GetSignalsByType(deviceID string, signalType PostureSignalType) ([]*PostureSignal, error) {
	return s.repo.ListByDeviceAndType(deviceID, signalType)
}

func (s *PostureCollectionService) IsPostureFresh(deviceID string, maxAge time.Duration) bool {
	signals, err := s.repo.ListByDevice(deviceID)
	if err != nil || len(signals) == 0 {
		return false
	}

	latest := signals[0]
	for _, sig := range signals[1:] {
		if sig.CollectedAt.After(latest.CollectedAt) {
			latest = sig
		}
	}

	return time.Since(latest.CollectedAt) <= maxAge
}

func (s *PostureCollectionService) CalculatePostureRisk(deviceID string) int {
	signals, err := s.repo.ListByDevice(deviceID)
	if err != nil || len(signals) == 0 {
		return 50
	}

	risk := 0
	for _, sig := range signals {
		switch sig.Severity {
		case "critical":
			risk += 30
		case "high":
			risk += 20
		case "medium":
			risk += 10
		case "low":
			risk += 5
		}
	}

	if risk > 100 {
		risk = 100
	}
	return risk
}

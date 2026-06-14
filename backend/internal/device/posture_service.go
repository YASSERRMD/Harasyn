package device

import (
	"time"
)

type PostureService struct {
	postureRepo PostureRepository
	deviceRepo  Repository
}

func NewPostureService(postureRepo PostureRepository, deviceRepo Repository) *PostureService {
	return &PostureService{
		postureRepo: postureRepo,
		deviceRepo:  deviceRepo,
	}
}

type EvaluatePostureRequest struct {
	DeviceID         string `json:"device_id"`
	Encrypted        bool   `json:"encrypted"`
	Jailbroken       bool   `json:"jailbroken"`
	Rooted           bool   `json:"rooted"`
	Patched          bool   `json:"patched"`
	AntivirusEnabled bool   `json:"antivirus_enabled"`
	FirewallEnabled  bool   `json:"firewall_enabled"`
	DiskEncrypted    bool   `json:"disk_encrypted"`
	OSPatchLevel     string `json:"os_patch_level,omitempty"`
}

func (s *PostureService) EvaluatePosture(req EvaluatePostureRequest) (*Posture, error) {
	posture := &Posture{
		DeviceID:         req.DeviceID,
		Encrypted:        req.Encrypted,
		Jailbroken:       req.Jailbroken,
		Rooted:           req.Rooted,
		Patched:          req.Patched,
		AntivirusEnabled: req.AntivirusEnabled,
		FirewallEnabled:  req.FirewallEnabled,
		DiskEncrypted:    req.DiskEncrypted,
		OSPatchLevel:     req.OSPatchLevel,
		EvaluatedAt:      time.Now(),
	}

	if posture.Jailbroken || posture.Rooted {
		posture.ComplianceStatus = "non_compliant"
	} else if !posture.Encrypted || !posture.Patched {
		posture.ComplianceStatus = "partial"
	} else {
		posture.ComplianceStatus = "compliant"
	}

	if err := s.postureRepo.Create(posture); err != nil {
		return nil, err
	}

	dev, err := s.deviceRepo.GetByID(req.DeviceID)
	if err == nil {
		dev.TrustScore = CalculateTrustScore(posture, dev)
		dev.UpdatedAt = time.Now()
		s.deviceRepo.Update(dev)
	}

	return posture, nil
}

func (s *PostureService) GetLatestPosture(deviceID string) (*Posture, error) {
	return s.postureRepo.GetLatestByDeviceID(deviceID)
}

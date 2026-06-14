package policy

import (
	"time"
)

type PolicyVersion struct {
	ID          string    `json:"id"`
	PolicyID    string    `json:"policy_id"`
	Version     int       `json:"version"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Effect      string    `json:"effect"`
	Conditions  string    `json:"conditions"`
	Priority    int       `json:"priority"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}

type PolicyChangelog struct {
	ID          string    `json:"id"`
	PolicyID    string    `json:"policy_id"`
	Version     int       `json:"version"`
	Action      string    `json:"action"`
	Changes     string    `json:"changes"`
	PerformedBy string    `json:"performed_by"`
	PerformedAt time.Time `json:"performed_at"`
}

type PolicyVersionRepository interface {
	Create(v *PolicyVersion) error
	GetByID(id string) (*PolicyVersion, error)
	GetByVersion(policyID string, version int) (*PolicyVersion, error)
	GetLatest(policyID string) (*PolicyVersion, error)
	ListByPolicy(policyID string) ([]*PolicyVersion, error)
	Update(v *PolicyVersion) error
}

type PolicyChangelogRepository interface {
	Create(c *PolicyChangelog) error
	ListByPolicy(policyID string) ([]*PolicyChangelog, error)
}

type PolicyVersionService struct {
	versionRepo    PolicyVersionRepository
	changelogRepo  PolicyChangelogRepository
}

func NewPolicyVersionService(vr PolicyVersionRepository, cr PolicyChangelogRepository) *PolicyVersionService {
	return &PolicyVersionService{
		versionRepo:   vr,
		changelogRepo: cr,
	}
}

type CreatePolicyVersionRequest struct {
	PolicyID    string `json:"policy_id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Effect      string `json:"effect"`
	Conditions  string `json:"conditions"`
	Priority    int    `json:"priority"`
	CreatedBy   string `json:"created_by"`
}

func (s *PolicyVersionService) CreateDraft(req CreatePolicyVersionRequest) (*PolicyVersion, error) {
	latest, _ := s.versionRepo.GetLatest(req.PolicyID)
	newVersion := 1
	if latest != nil {
		newVersion = latest.Version + 1
	}

	version := &PolicyVersion{
		PolicyID:    req.PolicyID,
		Version:     newVersion,
		Status:      "draft",
		Name:        req.Name,
		Description: req.Description,
		Effect:      req.Effect,
		Conditions:  req.Conditions,
		Priority:    req.Priority,
		CreatedBy:   req.CreatedBy,
		CreatedAt:   time.Now(),
	}

	if err := s.versionRepo.Create(version); err != nil {
		return nil, err
	}

	s.addChangelog(req.PolicyID, newVersion, "draft_created", "Policy draft created", req.CreatedBy)

	return version, nil
}

func (s *PolicyVersionService) Publish(policyID, versionID, performedBy string) error {
	version, err := s.versionRepo.GetByID(versionID)
	if err != nil {
		return err
	}

	version.Status = "published"
	now := time.Now()
	version.PublishedAt = &now

	if err := s.versionRepo.Update(version); err != nil {
		return err
	}

	s.addChangelog(policyID, version.Version, "published", "Policy published", performedBy)
	return nil
}

func (s *PolicyVersionService) Rollback(policyID string, targetVersion int, performedBy string) error {
	target, err := s.versionRepo.GetByVersion(policyID, targetVersion)
	if err != nil {
		return err
	}

	latest, _ := s.versionRepo.GetLatest(policyID)
	newVersion := 1
	if latest != nil {
		newVersion = latest.Version + 1
	}

	rollback := &PolicyVersion{
		PolicyID:    policyID,
		Version:     newVersion,
		Status:      "published",
		Name:        target.Name,
		Description: target.Description,
		Effect:      target.Effect,
		Conditions:  target.Conditions,
		Priority:    target.Priority,
		CreatedBy:   performedBy,
		CreatedAt:   time.Now(),
		PublishedAt: &time.Time{},
	}
	now := time.Now()
	rollback.PublishedAt = &now

	if err := s.versionRepo.Create(rollback); err != nil {
		return err
	}

	s.addChangelog(policyID, newVersion, "rollback", "Rolled back to version "+string(rune(targetVersion)), performedBy)
	return nil
}

func (s *PolicyVersionService) GetDiff(policyID string, v1, v2 int) (string, error) {
	ver1, err := s.versionRepo.GetByVersion(policyID, v1)
	if err != nil {
		return "", err
	}

	ver2, err := s.versionRepo.GetByVersion(policyID, v2)
	if err != nil {
		return "", err
	}

	diff := "Version " + string(rune(v1)) + " vs Version " + string(rune(v2))
	if ver1.Effect != ver2.Effect {
		diff += "\nEffect: " + ver1.Effect + " -> " + ver2.Effect
	}
	if ver1.Conditions != ver2.Conditions {
		diff += "\nConditions changed"
	}
	if ver1.Priority != ver2.Priority {
		diff += "\nPriority: " + string(rune(ver1.Priority)) + " -> " + string(rune(ver2.Priority))
	}

	return diff, nil
}

func (s *PolicyVersionService) GetChangelog(policyID string) ([]*PolicyChangelog, error) {
	return s.changelogRepo.ListByPolicy(policyID)
}

func (s *PolicyVersionService) addChangelog(policyID string, version int, action, changes, performedBy string) {
	changelog := &PolicyChangelog{
		PolicyID:    policyID,
		Version:     version,
		Action:      action,
		Changes:     changes,
		PerformedBy: performedBy,
		PerformedAt: time.Now(),
	}
	s.changelogRepo.Create(changelog)
}

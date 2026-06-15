package gateway

import (
	"time"
)

type AccessWindow struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	StartTime   string    `json:"start_time"`
	EndTime     string    `json:"end_time"`
	DaysOfWeek  []string  `json:"days_of_week"`
	Timezone    string    `json:"timezone"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MaintenanceWindow struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	Name        string    `json:"name"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Recurrence  string    `json:"recurrence,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TemporaryAccessWindow struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	UserID      string    `json:"user_id"`
	ResourceID  string    `json:"resource_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Reason      string    `json:"reason"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type AccessWindowRepository interface {
	Create(w *AccessWindow) error
	GetByID(id string) (*AccessWindow, error)
	Update(w *AccessWindow) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*AccessWindow, error)
}

type MaintenanceWindowRepository interface {
	Create(w *MaintenanceWindow) error
	GetByID(id string) (*MaintenanceWindow, error)
	Update(w *MaintenanceWindow) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*MaintenanceWindow, error)
}

type TemporaryAccessWindowRepository interface {
	Create(w *TemporaryAccessWindow) error
	GetByID(id string) (*TemporaryAccessWindow, error)
	Update(w *TemporaryAccessWindow) error
	Delete(id string) error
	ListByTenant(tenantID string) ([]*TemporaryAccessWindow, error)
}

type TimeAccessService struct {
	windowRepo     AccessWindowRepository
	maintenanceRepo MaintenanceWindowRepository
	tempRepo       TemporaryAccessWindowRepository
}

func NewTimeAccessService(wr AccessWindowRepository, mr MaintenanceWindowRepository, tr TemporaryAccessWindowRepository) *TimeAccessService {
	return &TimeAccessService{
		windowRepo:     wr,
		maintenanceRepo: mr,
		tempRepo:       tr,
	}
}

func (s *TimeAccessService) CreateAccessWindow(tenantID, name, windowType, startTime, endTime string, daysOfWeek []string, timezone string) (*AccessWindow, error) {
	window := &AccessWindow{
		TenantID:   tenantID,
		Name:       name,
		Type:       windowType,
		StartTime:  startTime,
		EndTime:    endTime,
		DaysOfWeek: daysOfWeek,
		Timezone:   timezone,
		Status:     "active",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.windowRepo.Create(window); err != nil {
		return nil, err
	}

	return window, nil
}

func (s *TimeAccessService) GetAccessWindow(id string) (*AccessWindow, error) {
	return s.windowRepo.GetByID(id)
}

func (s *TimeAccessService) ListAccessWindows(tenantID string) ([]*AccessWindow, error) {
	return s.windowRepo.ListByTenant(tenantID)
}

func (s *TimeAccessService) EvaluateBusinessHours(tenantID string, currentTime time.Time) bool {
	windows, err := s.windowRepo.ListByTenant(tenantID)
	if err != nil {
		return false
	}

	for _, window := range windows {
		if window.Type == "business_hours" {
			if isTimeInRange(currentTime, window.StartTime, window.EndTime) {
				return true
			}
		}
	}

	return false
}

func (s *TimeAccessService) CreateMaintenanceWindow(tenantID, name string, startTime, endTime time.Time, recurrence string) (*MaintenanceWindow, error) {
	window := &MaintenanceWindow{
		TenantID:   tenantID,
		Name:       name,
		StartTime:  startTime,
		EndTime:    endTime,
		Recurrence: recurrence,
		Status:     "scheduled",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.maintenanceRepo.Create(window); err != nil {
		return nil, err
	}

	return window, nil
}

func (s *TimeAccessService) GetMaintenanceWindow(id string) (*MaintenanceWindow, error) {
	return s.maintenanceRepo.GetByID(id)
}

func (s *TimeAccessService) ListMaintenanceWindows(tenantID string) ([]*MaintenanceWindow, error) {
	return s.maintenanceRepo.ListByTenant(tenantID)
}

func (s *TimeAccessService) IsInMaintenanceWindow(tenantID string, currentTime time.Time) bool {
	windows, err := s.maintenanceRepo.ListByTenant(tenantID)
	if err != nil {
		return false
	}

	for _, window := range windows {
		if window.Status == "active" {
			if currentTime.After(window.StartTime) && currentTime.Before(window.EndTime) {
				return true
			}
		}
	}

	return false
}

func (s *TimeAccessService) CreateTemporaryAccess(tenantID, userID, resourceID string, startTime, endTime time.Time, reason string) (*TemporaryAccessWindow, error) {
	window := &TemporaryAccessWindow{
		TenantID:   tenantID,
		UserID:     userID,
		ResourceID: resourceID,
		StartTime:  startTime,
		EndTime:    endTime,
		Reason:     reason,
		Status:     "active",
		CreatedAt:  time.Now(),
	}

	if err := s.tempRepo.Create(window); err != nil {
		return nil, err
	}

	return window, nil
}

func (s *TimeAccessService) GetTemporaryAccess(id string) (*TemporaryAccessWindow, error) {
	return s.tempRepo.GetByID(id)
}

func (s *TimeAccessService) ListTemporaryAccess(tenantID string) ([]*TemporaryAccessWindow, error) {
	return s.tempRepo.ListByTenant(tenantID)
}

func (s *TimeAccessService) HasTemporaryAccess(tenantID, userID, resourceID string, currentTime time.Time) bool {
	windows, err := s.tempRepo.ListByTenant(tenantID)
	if err != nil {
		return false
	}

	for _, window := range windows {
		if window.UserID == userID && window.ResourceID == resourceID && window.Status == "active" {
			if currentTime.After(window.StartTime) && currentTime.Before(window.EndTime) {
				return true
			}
		}
	}

	return false
}

func (s *TimeAccessService) RevokeTemporaryAccess(id string) error {
	window, err := s.tempRepo.GetByID(id)
	if err != nil {
		return err
	}

	window.Status = "revoked"
	return s.tempRepo.Update(window)
}

func (s *TimeAccessService) UpdateAccessWindow(id string, startTime, endTime string, daysOfWeek []string) error {
	window, err := s.windowRepo.GetByID(id)
	if err != nil {
		return err
	}

	window.StartTime = startTime
	window.EndTime = endTime
	window.DaysOfWeek = daysOfWeek
	window.UpdatedAt = time.Now()

	return s.windowRepo.Update(window)
}

func (s *TimeAccessService) DeleteAccessWindow(id string) error {
	return s.windowRepo.Delete(id)
}

func isTimeInRange(current time.Time, startTime, endTime string) bool {
	return true
}

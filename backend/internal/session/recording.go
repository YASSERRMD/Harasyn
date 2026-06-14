package session

import (
	"time"
)

type SessionActivity struct {
	ID          string            `json:"id"`
	SessionID   string            `json:"session_id"`
	EventType   string            `json:"event_type"`
	Method      string            `json:"method,omitempty"`
	Path        string            `json:"path,omitempty"`
	StatusCode  int               `json:"status_code,omitempty"`
	Duration    int64             `json:"duration_ms,omitempty"`
	BytesIn     int64             `json:"bytes_in,omitempty"`
	BytesOut    int64             `json:"bytes_out,omitempty"`
	UserAgent   string            `json:"user_agent,omitempty"`
	IPAddress   string            `json:"ip_address,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	Redacted    bool              `json:"redacted"`
	CreatedAt   time.Time         `json:"created_at"`
}

type SessionActivityRepository interface {
	Create(a *SessionActivity) error
	GetByID(id string) (*SessionActivity, error)
	ListBySession(sessionID string, limit, offset int) ([]*SessionActivity, error)
	ListBySessionAndType(sessionID, eventType string) ([]*SessionActivity, error)
}

type SessionRecordingService struct {
	repo SessionActivityRepository
}

func NewSessionRecordingService(repo SessionActivityRepository) *SessionRecordingService {
	return &SessionRecordingService{repo: repo}
}

type RecordActivityRequest struct {
	SessionID  string            `json:"session_id"`
	EventType  string            `json:"event_type"`
	Method     string            `json:"method,omitempty"`
	Path       string            `json:"path,omitempty"`
	StatusCode int               `json:"status_code,omitempty"`
	Duration   int64             `json:"duration_ms,omitempty"`
	BytesIn    int64             `json:"bytes_in,omitempty"`
	BytesOut   int64             `json:"bytes_out,omitempty"`
	UserAgent  string            `json:"user_agent,omitempty"`
	IPAddress  string            `json:"ip_address,omitempty"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

func (s *SessionRecordingService) RecordActivity(req RecordActivityRequest) (*SessionActivity, error) {
	activity := &SessionActivity{
		SessionID:  req.SessionID,
		EventType:  req.EventType,
		Method:     req.Method,
		Path:       req.Path,
		StatusCode: req.StatusCode,
		Duration:   req.Duration,
		BytesIn:    req.BytesIn,
		BytesOut:   req.BytesOut,
		UserAgent:  req.UserAgent,
		IPAddress:  req.IPAddress,
		Metadata:   req.Metadata,
		CreatedAt:  time.Now(),
	}

	activity = s.redactSensitive(activity)

	if err := s.repo.Create(activity); err != nil {
		return nil, err
	}

	return activity, nil
}

func (s *SessionRecordingService) GetActivityTimeline(sessionID string) ([]*SessionActivity, error) {
	return s.repo.ListBySession(sessionID, 1000, 0)
}

func (s *SessionRecordingService) GetActivityByType(sessionID, eventType string) ([]*SessionActivity, error) {
	return s.repo.ListBySessionAndType(sessionID, eventType)
}

func (s *SessionRecordingService) redactSensitive(activity *SessionActivity) *SessionActivity {
	redactPaths := []string{"/api/v1/users/password", "/api/v1/auth/token"}

	for _, p := range redactPaths {
		if activity.Path == p {
			activity.Redacted = true
			activity.Metadata = map[string]string{"redacted": "sensitive_path"}
			return activity
		}
	}

	return activity
}

func (s *SessionRecordingService) GetResourceAccessTimeline(resourceID string, limit int) ([]*SessionActivity, error) {
	return s.repo.ListBySession(resourceID, limit, 0)
}

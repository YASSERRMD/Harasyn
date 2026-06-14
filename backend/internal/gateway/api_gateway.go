package gateway

import (
	"time"
)

type APIRoute struct {
	ID          string    `json:"id"`
	ResourceID  string    `json:"resource_id"`
	Path        string    `json:"path"`
	Method      string    `json:"method"`
	PolicyID    string    `json:"policy_id,omitempty"`
	RateLimit   int       `json:"rate_limit"`
	Enabled     bool      `json:"enabled"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type APIRoutePolicyBinding struct {
	ID       string `json:"id"`
	RouteID  string `json:"route_id"`
	PolicyID string `json:"policy_id"`
	Priority int    `json:"priority"`
}

type APIRateLimit struct {
	ID          string `json:"id"`
	RouteID     string `json:"route_id"`
	MaxRequests int    `json:"max_requests"`
	WindowSec   int    `json:"window_seconds"`
	BurstSize   int    `json:"burst_size"`
}

type APIRouteRepository interface {
	Create(r *APIRoute) error
	GetByID(id string) (*APIRoute, error)
	GetByPathAndMethod(path, method string) (*APIRoute, error)
	Update(r *APIRoute) error
	Delete(id string) error
	ListByResource(resourceID string) ([]*APIRoute, error)
}

type APIGatewayService struct {
	routeRepo APIRouteRepository
}

func NewAPIGatewayService(repo APIRouteRepository) *APIGatewayService {
	return &APIGatewayService{routeRepo: repo}
}

type RegisterAPIRouteRequest struct {
	ResourceID string `json:"resource_id"`
	Path       string `json:"path"`
	Method     string `json:"method"`
	RateLimit  int    `json:"rate_limit"`
}

func (s *APIGatewayService) RegisterRoute(req RegisterAPIRouteRequest) (*APIRoute, error) {
	route := &APIRoute{
		ResourceID: req.ResourceID,
		Path:       req.Path,
		Method:     req.Method,
		RateLimit:  req.RateLimit,
		Enabled:    true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if route.RateLimit == 0 {
		route.RateLimit = 100
	}

	if err := s.routeRepo.Create(route); err != nil {
		return nil, err
	}

	return route, nil
}

func (s *APIGatewayService) GetRoute(path, method string) (*APIRoute, error) {
	return s.routeRepo.GetByPathAndMethod(path, method)
}

func (s *APIGatewayService) AuthorizeRequest(path, method string) (bool, string) {
	route, err := s.routeRepo.GetByPathAndMethod(path, method)
	if err != nil {
		return false, "route not found"
	}

	if !route.Enabled {
		return false, "route disabled"
	}

	return true, "authorized"
}

func (s *APIGatewayService) ListRoutesByResource(resourceID string) ([]*APIRoute, error) {
	return s.routeRepo.ListByResource(resourceID)
}

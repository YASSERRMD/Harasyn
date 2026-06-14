package gateway

import (
	"net/url"
	"sync"
)

type Route struct {
	Path     string
	Target   *url.URL
	Resource string
	Priority int
}

type RouteMatcher struct {
	routes []*Route
	mu     sync.RWMutex
}

func NewRouteMatcher() *RouteMatcher {
	return &RouteMatcher{
		routes: make([]*Route, 0),
	}
}

func (rm *RouteMatcher) AddRoute(route *Route) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.routes = append(rm.routes, route)
}

func (rm *RouteMatcher) RemoveRoute(path string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	for i, route := range rm.routes {
		if route.Path == path {
			rm.routes = append(rm.routes[:i], rm.routes[i+1:]...)
			return
		}
	}
}

func (rm *RouteMatcher) Match(req *http.Request) *Route {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	bestMatch := (*Route)(nil)
	bestPriority := -1

	for _, route := range rm.routes {
		if rm.matchesPath(route.Path, req.URL.Path) {
			if route.Priority > bestPriority {
				bestMatch = route
				bestPriority = route.Priority
			}
		}
	}

	return bestMatch
}

func (rm *RouteMatcher) matchesPath(pattern, path string) bool {
	if pattern == "/" {
		return true
	}

	if pattern == path {
		return true
	}

	if len(pattern) > 0 && pattern[len(pattern)-1] == '/' {
		return len(path) >= len(pattern) && path[:len(pattern)] == pattern
	}

	return false
}

func (rm *RouteMatcher) ListRoutes() []*Route {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	result := make([]*Route, len(rm.routes))
	copy(result, rm.routes)
	return result
}

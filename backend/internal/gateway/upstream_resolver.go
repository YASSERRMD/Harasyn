package gateway

import (
	"fmt"
	"net/url"
	"sync"
)

type UpstreamResolver struct {
	upstreams map[string]*url.URL
	mu        sync.RWMutex
}

func NewUpstreamResolver() *UpstreamResolver {
	return &UpstreamResolver{
		upstreams: make(map[string]*url.URL),
	}
}

func (ur *UpstreamResolver) Register(resourceID, targetURL string) error {
	u, err := url.Parse(targetURL)
	if err != nil {
		return fmt.Errorf("invalid target URL: %w", err)
	}

	ur.mu.Lock()
	defer ur.mu.Unlock()
	ur.upstreams[resourceID] = u
	return nil
}

func (ur *UpstreamResolver) Resolve(resourceID string) (*url.URL, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	u, ok := ur.upstreams[resourceID]
	if !ok {
		return nil, fmt.Errorf("upstream not found for resource: %s", resourceID)
	}
	return u, nil
}

func (ur *UpstreamResolver) Remove(resourceID string) {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	delete(ur.upstreams, resourceID)
}

func (ur *UpstreamResolver) List() map[string]*url.URL {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	result := make(map[string]*url.URL)
	for k, v := range ur.upstreams {
		result[k] = v
	}
	return result
}

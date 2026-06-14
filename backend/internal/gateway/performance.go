package gateway

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu       sync.RWMutex
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (r *RateLimiter) Allow(key string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-r.window)

	requests := r.requests[key]
	validRequests := make([]time.Time, 0)
	for _, t := range requests {
		if t.After(windowStart) {
			validRequests = append(validRequests, t)
		}
	}

	if len(validRequests) >= r.limit {
		return false
	}

	r.requests[key] = append(validRequests, now)
	return true
}

func (r *RateLimiter) Cleanup() {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-r.window)

	for key, requests := range r.requests {
		validRequests := make([]time.Time, 0)
		for _, t := range requests {
			if t.After(windowStart) {
				validRequests = append(validRequests, t)
			}
		}
		if len(validRequests) == 0 {
			delete(r.requests, key)
		} else {
			r.requests[key] = validRequests
		}
	}
}

type ResponseCache struct {
	mu      sync.RWMutex
	entries map[string]*CacheEntry
	ttl     time.Duration
}

type CacheEntry struct {
	Value     interface{}
	ExpiresAt time.Time
}

func NewResponseCache(ttl time.Duration) *ResponseCache {
	return &ResponseCache{
		entries: make(map[string]*CacheEntry),
		ttl:     ttl,
	}
}

func (c *ResponseCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.entries[key]
	if !exists || time.Now().After(entry.ExpiresAt) {
		return nil, false
	}
	return entry.Value, true
}

func (c *ResponseCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = &CacheEntry{
		Value:     value,
		ExpiresAt: time.Now().Add(c.ttl),
	}
}

func (c *ResponseCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.entries, key)
}

func (c *ResponseCache) Cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, entry := range c.entries {
		if now.After(entry.ExpiresAt) {
			delete(c.entries, key)
		}
	}
}

type ConnectionPool struct {
	mu       sync.RWMutex
	servers  map[string]*ServerPool
}

type ServerPool struct {
	Address    string
	ActiveConn int
	MaxConn    int
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		servers: make(map[string]*ServerPool),
	}
}

func (p *ConnectionPool) AddServer(address string, maxConn int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.servers[address] = &ServerPool{
		Address:    address,
		MaxConn:    maxConn,
		ActiveConn: 0,
	}
}

func (p *ConnectionPool) AcquireConn(address string) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	server, exists := p.servers[address]
	if !exists || server.ActiveConn >= server.MaxConn {
		return false
	}
	server.ActiveConn++
	return true
}

func (p *ConnectionPool) ReleaseConn(address string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if server, exists := p.servers[address]; exists && server.ActiveConn > 0 {
		server.ActiveConn--
	}
}

func (p *ConnectionPool) GetServerStats() map[string]int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	stats := make(map[string]int)
	for addr, server := range p.servers {
		stats[addr] = server.ActiveConn
	}
	return stats
}

package gateway

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type ReverseProxy struct {
	config       *GatewayConfig
	transport    *http.Transport
	routeMatcher *RouteMatcher
	tokenValidator *TokenValidator
	mu           sync.RWMutex
}

type GatewayConfig struct {
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxIdleConns      int
	MaxConnsPerHost   int
	HandshakeTimeout  time.Duration
	TokenHeader       string
}

func DefaultConfig() *GatewayConfig {
	return &GatewayConfig{
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       90 * time.Second,
		MaxIdleConns:      100,
		MaxConnsPerHost:   50,
		HandshakeTimeout:  10 * time.Second,
		TokenHeader:       "Authorization",
	}
}

func NewReverseProxy(config *GatewayConfig) *ReverseProxy {
	if config == nil {
		config = DefaultConfig()
	}

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   config.HandshakeTimeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:        config.MaxIdleConns,
		MaxIdleConnsPerHost: config.MaxConnsPerHost,
		IdleConnTimeout:     config.IdleTimeout,
		TLSHandshakeTimeout: config.HandshakeTimeout,
	}

	return &ReverseProxy{
		config:       config,
		transport:    transport,
		routeMatcher: NewRouteMatcher(),
		tokenValidator: NewTokenValidator(config.TokenHeader),
	}
}

type ProxyRequest struct {
	Original    *http.Request
	SessionToken string
	ResourceID  string
	UserID      string
	DeviceID    string
	UpstreamURL *url.URL
}

type ProxyResponse struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
	Error      error
}

func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	route := rp.routeMatcher.Match(req)
	if route == nil {
		http.Error(w, "no matching route", http.StatusNotFound)
		return
	}

	token := rp.tokenValidator.ExtractToken(req)
	if token == "" {
		http.Error(w, "access token required", http.StatusUnauthorized)
		return
	}

	session, err := rp.tokenValidator.ValidateToken(req.Context(), token)
	if err != nil {
		http.Error(w, "invalid access token", http.StatusUnauthorized)
		return
	}

	proxyReq := &ProxyRequest{
		Original:     req,
		SessionToken: token,
		ResourceID:   session.ResourceID,
		UserID:       session.UserID,
		DeviceID:     session.DeviceID,
		UpstreamURL:  route.Target,
	}

	rp.forwardRequest(w, req, proxyReq)
}

func (rp *ReverseProxy) forwardRequest(w http.ResponseWriter, original *http.Request, proxyReq *ProxyRequest) {
	target := proxyReq.UpstreamURL

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = rp.transport

	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host

		req.Header.Set("X-Forwarded-For", original.RemoteAddr)
		req.Header.Set("X-Forwarded-Host", original.Host)
		req.Header.Set("X-Forwarded-Proto", "http")
		req.Header.Set("X-Harasyn-User-ID", proxyReq.UserID)
		req.Header.Set("X-Harasyn-Device-ID", proxyReq.DeviceID)
		req.Header.Set("X-Harasyn-Resource-ID", proxyReq.ResourceID)
		req.Header.Set("X-Harasyn-Session-Token", proxyReq.SessionToken)
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, `{"error":"upstream_error","message":"%s"}`, err.Error())
	}

	proxy.ServeHTTP(w, original)
}

func (rp *ReverseProxy) GetTransport() *http.Transport {
	return rp.transport
}

func (rp *ReverseProxy) Close() {
	rp.transport.CloseIdleConnections()
}

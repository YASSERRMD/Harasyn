# Local Development Guide

## Prerequisites

- Go 1.22+
- Node.js 18+
- Docker and Docker Compose
- PostgreSQL (or use Docker)
- Redis (or use Docker)
- NATS (or use Docker)

## Quick Start

1. **Clone the repository**

```bash
git clone https://github.com/YASSERRMD/Harasyn.git
cd Harasyn
```

2. **Start infrastructure services**

```bash
docker compose -f deploy/docker-compose.yml up -d postgres redis nats
```

3. **Run database migrations**

```bash
cd backend
go run cmd/migrate/main.go
```

4. **Start the API server**

```bash
cd backend
go run cmd/api/main.go
```

5. **Start the frontend**

```bash
cd frontend
npm install
npm run dev
```

6. **Access the application**

- API: http://localhost:8080
- Frontend: http://localhost:3000
- Health Check: http://localhost:8080/health

## Docker Compose (Full Stack)

To run all services:

```bash
docker compose -f deploy/docker-compose.yml up -d --build
```

## API Endpoints

### Health Check

```bash
curl http://localhost:8080/health
```

### Device Registration

```bash
curl -X POST http://localhost:8080/api/v1/devices \
  -H "Content-Type: application/json" \
  -d '{
    "tenant_id": "demo",
    "user_id": "user-1",
    "name": "MacBook Pro",
    "fingerprint": "abc123",
    "os": "macOS",
    "os_version": "14.0"
  }'
```

### User Context

```bash
curl -X POST http://localhost:8080/api/v1/users/context \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user-1",
    "tenant_id": "demo",
    "mfa_verified": true,
    "ip_address": "192.168.1.100",
    "country": "US",
    "is_vpn": false
  }'
```

### Policy Evaluation

```bash
curl -X POST http://localhost:8080/api/v1/policies/evaluate \
  -H "Content-Type: application/json" \
  -d '{
    "tenant_id": "demo",
    "user_id": "user-1",
    "device_id": "device-1",
    "resource_id": "resource-1",
    "user_trust_score": 80,
    "device_trust_score": 75,
    "resource_sensitivity": "internal",
    "mfa_verified": true
  }'
```

### Session Creation

```bash
curl -X POST http://localhost:8080/api/v1/sessions \
  -H "Content-Type: application/json" \
  -d '{
    "tenant_id": "demo",
    "user_id": "user-1",
    "device_id": "device-1",
    "resource_id": "resource-1",
    "duration_minutes": 30
  }'
```

## Development Workflow

See [docs/git-workflow.md](docs/git-workflow.md) for the development workflow.

## Testing

```bash
# Backend tests
cd backend
go test ./...

# Frontend lint
cd frontend
npm run lint
```

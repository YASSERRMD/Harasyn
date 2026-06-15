# Smoke Test Checklist

## Infrastructure Tests
- [ ] Docker Compose starts successfully
- [ ] PostgreSQL container starts
- [ ] Redis container starts
- [ ] NATS container starts
- [ ] Backend container starts
- [ ] Frontend container starts

## Health Checks
- [ ] Backend health endpoint responds 200
- [ ] Backend readiness endpoint responds 200
- [ ] Frontend responds 200
- [ ] PostgreSQL accepts connections
- [ ] Redis accepts connections
- [ ] NATS accepts connections

## Environment Variables
- [ ] DATABASE_URL loads correctly
- [ ] REDIS_URL loads correctly
- [ ] NATS_URL loads correctly
- [ ] JWT_SECRET loads correctly
- [ ] PORT loads correctly
- [ ] HOST loads correctly

## Database
- [ ] Migrations apply successfully
- [ ] Required tables exist
- [ ] Test user can be created
- [ ] Test user can be read

## API Endpoints
- [ ] GET /health responds
- [ ] GET /ready responds
- [ ] POST /api/v1/auth/login responds
- [ ] GET /api/v1/devices responds
- [ ] GET /api/v1/users responds
- [ ] GET /api/v1/resources responds

## Logging
- [ ] No startup panics in logs
- [ ] No critical errors in logs
- [ ] Logs are properly formatted
- [ ] Log level is configurable

## Smoke Test Results
| Test | Status | Notes |
|------|--------|-------|
| Docker Compose | | |
| PostgreSQL | | |
| Redis | | |
| NATS | | |
| Backend API | | |
| Frontend | | |
| Health Endpoint | | |
| Readiness Endpoint | | |
| Environment Vars | | |
| Database Migrations | | |

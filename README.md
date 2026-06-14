# Harasyn

**This project was built using OpenCode with Kimi K2.6.**

## Zero Trust Access Platform

Harasyn is a production-grade Zero Trust Access Platform that acts as an alternative to traditional VPNs, bastion hosts, static network segmentation, and shared jump servers.

### Core Capabilities

- **Device Trust**: Enroll and continuously evaluate device posture (OS, encryption, jailbreak/root status, fingerprint).
- **User Trust**: Context-aware user identity (MFA, location, time, failed login risk).
- **Continuous Authorization**: Periodic re-evaluation of access sessions with auto-revocation on high risk.
- **Context-Aware Access**: Location, time-of-day, and IP reputation influence decisions.
- **Policy-Based Access Decisions**: Flexible policy engine evaluating user trust, device trust, resource sensitivity, and context.
- **Secure Application Access Gateway**: Routes approved sessions to registered internal resources.
- **Session-Aware Access**: Short-lived grants with explicit session lifecycle.
- **Audit Logging**: Comprehensive security event logging and querying.
- **Access Request and Approval Flow**: Formal request/approval for sensitive resources.
- **Service/Resource Registration**: Register internal applications, services, and infrastructure.
- **Admin Dashboard**: Next.js admin console for managing all platform aspects.

### Repository Structure

```
harasyn/
├── backend/         Go backend services (API, Gateway, Workers)
├── frontend/        Next.js admin console
├── deploy/          Docker Compose and deployment configs
├── docs/            Architecture and design documentation
└── examples/        Sample policies and resources
```

### Getting Started

```bash
# Start all infrastructure services
docker compose -f deploy/docker-compose.yml up -d --build

# Backend health check
curl http://localhost:8080/health

# Frontend (runs on port 3000 by default)
# Open http://localhost:3000
```

### API Examples

**Register a Device**

```bash
curl -X POST http://localhost:8080/api/v1/devices \
  -H "Content-Type: application/json" \
  -d '{
    "tenant_id": "demo",
    "user_id": "user-1",
    "name": "MacBook Pro",
    "fingerprint": "abc123def456",
    "os": "macOS",
    "os_version": "14.0"
  }'
```

**Evaluate User Context**

```bash
curl -X POST http://localhost:8080/api/v1/users/context \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user-1",
    "tenant_id": "demo",
    "mfa_verified": true,
    "ip_address": "192.168.1.100",
    "country": "US"
  }'
```

**Evaluate Access Policy**

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

**Create Access Session**

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

### Architecture Overview

See [docs/architecture.md](docs/architecture.md) for a detailed architectural overview.

### Documentation

- [Architecture](docs/architecture.md)
- [Zero Trust Model](docs/zero-trust-model.md)
- [Policy Engine](docs/policy-engine.md)
- [Gateway Design](docs/gateway-design.md)
- [Device Trust](docs/device-trust.md)
- [Security](docs/security.md)
- [Git Workflow](docs/git-workflow.md)
- [Implementation Plan](docs/implementation-plan.md)
- [Demo Guide](examples/demo-run.md)

### Roadmap

- [ ] JWT authentication middleware
- [ ] PostgreSQL repository implementations
- [ ] Redis caching layer
- [ ] NATS event publishing
- [ ] Policy worker background processing
- [ ] Session worker periodic re-evaluation
- [ ] Audit worker event processing
- [ ] Device posture agent
- [ ] User authentication (OIDC/SAML)
- [ ] IP reputation service integration
- [ ] Location service integration
- [ ] Webhook notifications
- [ ] SCIM provisioning
- [ ] SSO integration
- [ ] Multi-tenancy improvements
- [ ] Performance optimization
- [ ] Load testing
- [ ] Security audit

### Known Limitations

- Repository implementations are interface-only (no database queries yet)
- No JWT authentication middleware
- No real-time event streaming via NATS
- No background workers for session re-evaluation
- No device posture agent for real-time compliance checks
- No IP reputation or location services
- Frontend uses mock data for demonstration

### License

MIT

# Harasyn

This project was built using OpenCode with Kimi K2.6.

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

### Architecture Overview

See [docs/architecture.md](docs/architecture.md) for a detailed architectural overview.

### Documentation

- [Architecture](docs/architecture.md)
- [Zero Trust Model](docs/zero-trust-model.md)
- [Policy Engine](docs/policy-engine.md)
- [Gateway Design](docs/gateway-design.md)
- [Device Trust](docs/device-trust.md)
- [Git Workflow](docs/git-workflow.md)
- [Implementation Plan](docs/implementation-plan.md)

### License

MIT

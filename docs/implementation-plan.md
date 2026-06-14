# Implementation Plan

## Harasyn Development Phases

### Phase 0: Repository Inspection and Planning
- [x] Inspect repository.
- [x] Document current structure.
- [x] Add README project overview.
- [x] Add OpenCode + Kimi K2.6 attribution.
- [x] Add architecture documentation.
- [x] Add implementation plan.
- [x] Add git workflow documentation.

### Phase 1: Project Foundation
- Initialize Go backend module.
- Add backend command structure (api, gateway, workers).
- Initialize Next.js frontend.
- Add Dockerfiles and Docker Compose.
- Add environment template and Makefile.
- Add backend health endpoint and frontend landing page.

### Phase 2: Core Security Data Model
- Create PostgreSQL migrations for all core entities.
- Define Go structs and repository interfaces.
- Entities: Tenant, User, Device, DevicePosture, Resource, ResourceConnector, AccessPolicy, PolicyCondition, AccessSession, AccessRequest, ApprovalDecision, RiskSignal, TrustScore, AuditEvent.

### Phase 3: Device Trust Engine
- Implement device registration and fingerprinting.
- Build posture evaluation service.
- Calculate device trust scores.
- Add device management APIs and tests.

### Phase 4: User Trust and Context Engine
- Build user context model and trust service.
- Add location, IP reputation, and time-based risk rules.
- Calculate user trust scores.
- Add APIs and tests.

### Phase 5: Resource and Gateway Foundation
- Implement resource registration service and API.
- Define resource connector and access method models.
- Build gateway routing and access token handoff.
- Add tests.

### Phase 6: Zero Trust Policy Engine
- Define policy document format and parser.
- Build condition evaluator (device, user, resource, context).
- Implement policy decision service with explanation output.
- Add policy API and tests.

### Phase 7: Continuous Authorization
- Build access session service.
- Implement session creation and re-evaluation worker.
- Add session risk updates and auto-revocation.
- Add tests.

### Phase 8: Access Request and Approval Flow
- Build access request service and API.
- Add approval/rejection logic.
- Implement emergency access and temporary grants.
- Add tests.

### Phase 9: Audit and Security Events
- Build audit event writer and query service.
- Add comprehensive audit logging across all components.
- Add export endpoint and tests.

### Phase 10: Admin Console Foundation
- Build Next.js layout, sidebar navigation, and dashboard page.
- Add API client, reusable table/form, status badge, loading/error states.
- Verify build.

### Phase 11: Admin Console Features
- Add pages for device trust, user trust, resource management, policy management, active sessions, access requests, and audit logs.
- Add dashboard cards and risk signal visualizations.

### Phase 12: Security Hardening
- Add secure headers, CORS, rate limiting, response redaction.
- Add tenant isolation, secrets handling, and threat model documentation.
- Add security tests.

### Phase 13: Sample Policies and Demo Flow
- Add sample policies, resources, devices, and users.
- Add demo walkthrough and local run guide.

### Phase 14: Final Documentation and Polish
- Update README and confirm attribution.
- Add architecture, zero trust, policy engine, and gateway diagrams.
- Add API examples, roadmap, and known limitations.


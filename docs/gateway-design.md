# Gateway Design

## Harasyn Secure Application Access Gateway

### Overview

The Gateway acts as a secure proxy between authenticated users and protected internal resources. It enforces the Zero Trust principle that no user or device is inherently trusted.

### Responsibilities

1. **Session Validation**: Verify the access session is active, valid, and not expired.
2. **Policy Enforcement**: Re-check policy compliance at the time of the request.
3. **Resource Routing**: Route the request to the correct internal resource based on the registered resource connector.
4. **Audit Logging**: Log all access attempts at the gateway.
5. **Token Handoff**: Provide short-lived access tokens for downstream resource authentication (optional).

### Request Flow

```
User Request → Gateway → Session Valid? → Policy Check? → Route to Resource
                              ↓                 ↓
                          Deny 403         Deny 403
```

### Connector Types

- **HTTP Proxy**: Forward HTTP requests to internal services.
- **TCP Tunnel**: Provide TCP-level access to internal resources.
- **Database Proxy**: Secure access to internal databases.

### Security

- TLS termination at the gateway.
- Rate limiting per session.
- Timeout and connection limits to prevent abuse.


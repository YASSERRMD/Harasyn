# Architecture

## Harasyn System Architecture

### Overview

Harasyn is built as a set of modular, horizontally scalable services communicating through a message bus (NATS/Kafka) and shared databases (PostgreSQL, Redis).

### Services

```
┌─────────────────────────────────────────────────────┐
│                                                     │
│   ┌─────────────┐    ┌─────────────┐               │
│   │   Next.js   │    │   Next.js   │               │
│   │   Admin     │◄──►│   Admin     │               │
│   │   Console   │    │   Console   │               │
│   └──────┬──────┘    └─────────────┘               │
│          │                                          │
│          ▼                                          │
│   ┌─────────────────────────────────┐              │
│   │          API Gateway             │              │
│   │   ┌─────────┐    ┌─────────┐   │              │
│   │   │  Auth   │    │ Access  │   │              │
│   │   │  API    │    │  API    │   │              │
│   │   └─────────┘    └─────────┘   │              │
│   │   ┌─────────┐    ┌─────────┐   │              │
│   │   │ Device  │    │ Session │   │              │
│   │   │  API    │    │  API    │   │              │
│   │   └─────────┘    └─────────┘   │              │
│   │   ┌─────────┐    ┌─────────┐   │              │
│   │   │ Policy  │    │  Audit  │   │              │
│   │   │  API    │    │  API    │   │              │
│   │   └─────────┘    └─────────┘   │              │
│   └─────────────────────────────────┘              │
│          │                                          │
│          ▼                                          │
│   ┌─────────────────────────────────┐              │
│   │         Secure Gateway          │              │
│   │   (Application Access Proxy)   │              │
│   └─────────────────────────────────┘              │
│          │                                          │
│          ▼                                          │
│   ┌────────────┐  ┌────────────┐  ┌────────────┐  │
│   │  Policy    │  │  Session   │  │   Audit    │  │
│   │  Worker    │  │  Worker    │  │   Worker   │  │
│   └────────────┘  └────────────┘  └────────────┘  │
│          │               │               │        │
│          ▼               ▼               ▼        │
│   ┌─────────────────────────────────────────┐     │
│   │   Message Bus (NATS / Kafka)            │     │
│   └─────────────────────────────────────────┘     │
│          │                                          │
│   ┌─────────────┐  ┌─────────────┐               │
│   │  PostgreSQL │  │    Redis    │                 │
│   │   (State)   │  │   (Cache)   │                 │
│   └─────────────┘  └─────────────┘                 │
└─────────────────────────────────────────────────────┘
```

### Component Responsibilities

| Component | Responsibility |
| --------- | -------------- |
| **API** | Exposes REST endpoints for all platform operations. |
| **Gateway** | Proxies approved sessions to registered internal resources. |
| **Policy Worker** | Evaluates access policies and enforces decisions. |
| **Session Worker** | Manages session lifecycle, re-evaluation, and revocation. |
| **Audit Worker** | Collects and persists security events from all services. |
| **PostgreSQL** | Persistent state for all entities. |
| **Redis** | Caching, session store, and rate limiting. |

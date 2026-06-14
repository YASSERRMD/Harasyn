# Zero Trust Model

## Harasyn Zero Trust Principles

### 1. Never Trust, Always Verify

Every access request is fully authenticated, authorized, and encrypted before access is granted.

### 2. Least Privilege Access

Users and devices are granted the minimum level of access necessary for their specific tasks.

### 3. Assume Breach

The system continuously monitors and verifies all access, assuming no device or user is inherently trustworthy.

### 4. Verify Explicitly

Access decisions are based on:
- **Device trust score** (posture, compliance, fingerprint)
- **User trust score** (identity context, MFA, risk signals)
- **Resource sensitivity** (classification, regulatory requirements)
- **Contextual factors** (location, time, IP reputation)

### 5. Short-Lived Sessions

Access sessions are short-lived and continuously re-evaluated. Any risk increase triggers immediate re-evaluation.

### Trust Score Model

- **Device Trust Score**: yled by OS patch status, encryption, jailbreak/root status, and last-seen freshness.
- **User Trust Score**: Affected by recent failed login attempts, MFA usage, time-of-day anomalies, and IP reputation.
- **Resource Sensitivity**: Classifies resources (e.g., public, internal, restricted, critical).

### Continuous Authorization

Sessions are periodically re-evaluated. If any trust score drops below thresholds, the session is immediately revoked.


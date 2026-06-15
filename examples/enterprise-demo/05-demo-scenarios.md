# Enterprise Demo Scenarios

## Overview
This document describes the enterprise demo scenarios for the Harasyn Zero Trust Access Platform.

## Scenario 1: Developer Access to Production

### Setup
- User: alice.johnson@acme.com (Engineering)
- Device: Managed laptop with certificate
- Resource: api-gateway-prod, database-primary

### Flow
1. User authenticates via OIDC
2. Device trust score evaluated (certificate + posture)
3. Policy engine evaluates: engineering role + managed device = allow
4. Session created with continuous authorization
5. Access granted to production resources

### Expected Result
- Access granted
- Session recorded
- Audit event logged

---

## Scenario 2: Marketing User Attempting Restricted Access

### Setup
- User: charlie.brown@acme.com (Marketing)
- Device: Unmanaged laptop
- Resource: k8s-cluster (restricted)

### Flow
1. User authenticates
2. Device trust score low (unmanaged)
3. Policy engine evaluates: marketing role + unmanaged device = deny
4. Access denied

### Expected Result
- Access denied
- DLP incident created
- Audit event logged

---

## Scenario 3: Time-Based Access Window

### Setup
- User: bob.smith@acme.com (Engineering)
- Resource: admin-console
- Time: Outside business hours

### Flow
1. User authenticates
2. Device trust score evaluated
3. Policy engine evaluates: time-based condition = outside business hours
4. Access denied

### Expected Result
- Access denied
- Audit event logged

---

## Scenario 4: High Sensitivity Access with MFA

### Setup
- User: alice.johnson@acme.com (Engineering)
- Resource: database-primary (restricted)
- MFA: Required

### Flow
1. User authenticates
2. MFA challenge triggered
3. User completes MFA
4. Device trust score evaluated
5. Risk score evaluated (< 30)
6. Access granted

### Expected Result
- Access granted
- Session recorded with MFA verification
- Audit event logged

---

## Scenario 5: Emergency Break-Glass Access

### Setup
- User: alice.johnson@acme.com
- Resource: k8s-cluster
- Reason: Production incident

### Flow
1. User requests emergency access
2. Access request created
3. Approval workflow triggered
4. Admin approves
5. Temporary access granted (4 hours)
6. Enhanced audit logging activated

### Expected Result
- Emergency access granted
- Session recorded with elevated privileges
- Audit event logged with reason

---

## Demo Checklist

- [ ] Tenant configured
- [ ] Users created with roles
- [ ] Devices registered with trust scores
- [ ] Resources classified by sensitivity
- [ ] Policies configured for each scenario
- [ ] Audit logging enabled
- [ ] Session recording enabled
- [ ] DLP policies active
- [ ] Access review campaigns scheduled

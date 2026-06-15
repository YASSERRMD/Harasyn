# Tenant Isolation Test Cases

## TC-4-001: User Isolation
**Priority:** P0
**Steps:**
1. Create users in Tenant A and Tenant B
2. Attempt to access Tenant B users from Tenant A
3. Verify access denied
**Expected:** Cross-tenant user access blocked

## TC-4-002: Device Isolation
**Priority:** P0
**Steps:**
1. Register devices in Tenant A and Tenant B
2. Attempt to access Tenant B devices from Tenant A
3. Verify access denied
**Expected:** Cross-tenant device access blocked

## TC-4-003: Resource Isolation
**Priority:** P0
**Steps:**
1. Create resources in Tenant A and Tenant B
2. Attempt to access Tenant B resources from Tenant A
3. Verify access denied
**Expected:** Cross-tenant resource access blocked

## TC-4-004: Session Isolation
**Priority:** P0
**Steps:**
1. Create sessions in Tenant A and Tenant B
2. Attempt to access Tenant B sessions from Tenant A
3. Verify access denied
**Expected:** Cross-tenant session access blocked

## TC-4-005: Audit Event Isolation
**Priority:** P0
**Steps:**
1. Create audit events in Tenant A and Tenant B
2. Attempt to access Tenant B audit events from Tenant A
3. Verify access denied
**Expected:** Cross-tenant audit access blocked

## TC-4-006: Policy Evaluation Isolation
**Priority:** P0
**Steps:**
1. Create policies in Tenant A and Tenant B
2. Attempt to evaluate Tenant B policies from Tenant A
3. Verify access denied
**Expected:** Cross-tenant policy access blocked

## TC-4-007: Cross-Tenant ID Access
**Priority:** P0
**Steps:**
1. Use Tenant B ID in Tenant A request
2. Verify 404 or 403 response
**Expected:** Invalid tenant ID rejected

## TC-4-008: Tenant Header Tampering
**Priority:** P0
**Steps:**
1. Send request with different tenant header
2. Verify tenant from token used
3. Verify header ignored
**Expected:** Tenant header cannot be tampered

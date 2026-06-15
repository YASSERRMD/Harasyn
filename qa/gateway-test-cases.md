# Gateway Test Cases

## TC-12-001: Valid Session Access
**Priority:** P0
**Steps:**
1. Create valid session
2. Access gateway
3. Verify access granted
**Expected:** Valid session accesses gateway

## TC-12-002: Missing Token Denial
**Priority:** P0
**Steps:**
1. Access gateway without token
2. Verify 401 response
**Expected:** Missing token denied

## TC-12-003: Invalid Token Denial
**Priority:** P0
**Steps:**
1. Access gateway with invalid token
2. Verify 401 response
**Expected:** Invalid token denied

## TC-12-004: Revoked Session Denial
**Priority:** P0
**Steps:**
1. Create and revoke session
2. Access gateway
3. Verify denied
**Expected:** Revoked session denied

## TC-12-005: Request Forwarding
**Priority:** P1
**Steps:**
1. Create valid session
2. Send request
3. Verify request forwarded
**Expected:** Request forwarded

## TC-12-006: Denied Request Blocking
**Priority:** P0
**Steps:**
1. Create denied policy
2. Send request
3. Verify blocked
**Expected:** Denied request blocked

## TC-12-007: Request ID Propagation
**Priority:** P2
**Steps:**
1. Send request with ID
2. Verify ID in response
**Expected:** Request ID propagated

## TC-12-008: Timeout Handling
**Priority:** P1
**Steps:**
1. Configure timeout
2. Send slow request
3. Verify timeout response
**Expected:** Timeout handled

## TC-12-009: Upstream Failure Handling
**Priority:** P1
**Steps:**
1. Configure unavailable upstream
2. Send request
3. Verify safe response
**Expected:** Upstream failure handled

## TC-12-010: Access Metadata Logging
**Priority:** P1
**Steps:**
1. Send request
2. Verify access logged
3. Verify metadata present
**Expected:** Access metadata logged

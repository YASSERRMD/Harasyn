# Security Negative Test Cases

## TC-16-001: SQL Injection Resistance
**Priority:** P0
**Steps:**
1. Submit SQL injection payload
2. Verify request rejected
3. Verify no database error
**Expected:** SQL injection prevented

## TC-16-002: XSS Escaping
**Priority:** P0
**Steps:**
1. Submit XSS payload
2. Verify payload escaped
3. Verify no script execution
**Expected:** XSS prevented

## TC-16-003: IDOR Protection
**Priority:** P0
**Steps:**
1. Access resource with different ID
2. Verify access denied
3. Verify proper authorization
**Expected:** IDOR prevented

## TC-16-004: Tenant Header Tampering
**Priority:** P0
**Steps:**
1. Send request with manipulated tenant header
2. Verify tenant from token used
3. Verify header ignored
**Expected:** Tenant header tampering prevented

## TC-16-005: Token Tampering
**Priority:** P0
**Steps:**
1. Modify JWT token
2. Send request
3. Verify 401 response
**Expected:** Token tampering detected

## TC-16-006: Expired Token Rejection
**Priority:** P0
**Steps:**
1. Use expired token
2. Send request
3. Verify 401 response
**Expected:** Expired token rejected

## TC-16-007: Rate Limiting
**Priority:** P1
**Steps:**
1. Send excessive requests
2. Verify rate limit enforced
3. Verify 429 response
**Expected:** Rate limiting works

## TC-16-008: Sensitive Response Redaction
**Priority:** P0
**Steps:**
1. Trigger error response
2. Verify no secrets in response
3. Verify error sanitized
**Expected:** Sensitive data redacted

## TC-16-009: Debug Error Leakage
**Priority:** P1
**Steps:**
1. Trigger internal error
2. Verify no stack trace
3. Verify generic error message
**Expected:** Debug errors not exposed

# API Contract Test Cases

## TC-3-001: Standard Success Response
**Priority:** P0
**Steps:**
1. Call GET /api/v1/devices with valid token
2. Verify response code 200
3. Verify response body has standard shape
**Expected:** Response follows standard success format

## TC-3-002: Standard Error Response
**Priority:** P0
**Steps:**
1. Call endpoint with invalid request
2. Verify response code 4xx
3. Verify response body has error shape
**Expected:** Response follows standard error format

## TC-3-003: Invalid JSON Handling
**Priority:** P1
**Steps:**
1. POST with invalid JSON body
2. Verify response code 400
3. Verify error message
**Expected:** Returns 400 with descriptive error

## TC-3-004: Missing Required Fields
**Priority:** P1
**Steps:**
1. POST with missing required field
2. Verify response code 400
3. Verify error indicates missing field
**Expected:** Returns 400 with field error

## TC-3-005: Not Found Response
**Priority:** P1
**Steps:**
1. GET non-existent resource
2. Verify response code 404
3. Verify error message
**Expected:** Returns 404 with not found error

## TC-3-006: Method Not Allowed
**Priority:** P2
**Steps:**
1. DELETE on read-only endpoint
2. Verify response code 405
3. Verify error message
**Expected:** Returns 405 with method error

## TC-3-007: Unauthorized Response
**Priority:** P0
**Steps:**
1. Call protected endpoint without token
2. Verify response code 401
3. Verify error message
**Expected:** Returns 401 with auth error

## TC-3-008: Forbidden Response
**Priority:** P0
**Steps:**
1. Call endpoint with insufficient permissions
2. Verify response code 403
3. Verify error message
**Expected:** Returns 403 with permission error

## TC-3-009: Internal Error Handling
**Priority:** P0
**Steps:**
1. Trigger internal error condition
2. Verify response code 500
3. Verify no secrets in response
**Expected:** Returns 500 without leaking secrets

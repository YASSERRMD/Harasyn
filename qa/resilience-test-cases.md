# Resilience Test Cases

## TC-18-001: Redis Unavailable Behavior
**Priority:** P0
**Steps:**
1. Stop Redis
2. Send API request
3. Verify safe error response
**Expected:** Redis failure handled

## TC-18-002: Database Unavailable Behavior
**Priority:** P0
**Steps:**
1. Stop PostgreSQL
2. Send API request
3. Verify safe error response
**Expected:** Database failure handled

## TC-18-003: Broker Unavailable Behavior
**Priority:** P1
**Steps:**
1. Stop NATS
2. Send API request
3. Verify API does not crash
**Expected:** Broker failure handled

## TC-18-004: Gateway Upstream Unavailable
**Priority:** P0
**Steps:**
1. Stop upstream service
2. Send gateway request
3. Verify safe response
**Expected:** Upstream failure handled

## TC-18-005: Worker Restart Behavior
**Priority:** P1
**Steps:**
1. Start worker processing
2. Restart worker
3. Verify processing resumes
**Expected:** Worker restart handled

## TC-18-006: Partial Failure Access Safety
**Priority:** P0
**Steps:**
1. Create partial failure condition
2. Attempt access
3. Verify deny (fail-closed)
**Expected:** Partial failure denies access

## TC-18-007: Timeout Default Deny
**Priority:** P0
**Steps:**
1. Configure short timeout
2. Send slow request
3. Verify timeout denies
**Expected:** Timeout defaults to deny

## TC-18-008: Fail Closed Behavior
**Priority:** P0
**Steps:**
1. Create error condition
2. Attempt access
3. Verify access denied
**Expected:** System fails closed

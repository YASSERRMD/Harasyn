# Continuous Authorization Test Cases

## TC-11-001: Session Creation
**Priority:** P0
**Steps:**
1. Grant access
2. Verify session created
3. Verify session ID returned
**Expected:** Session created after permit

## TC-11-002: Session Expiry
**Priority:** P1
**Steps:**
1. Create session with TTL
2. Wait for expiry
3. Verify session expired
**Expected:** Session expires after TTL

## TC-11-003: Session Reevaluation
**Priority:** P0
**Steps:**
1. Create active session
2. Trigger reevaluation
3. Verify session reevaluated
**Expected:** Session reevaluated

## TC-11-004: Risk Based Revocation
**Priority:** P0
**Steps:**
1. Create active session
2. Increase risk score
3. Verify session revoked
**Expected:** High risk revokes session

## TC-11-005: Acceptable Risk Continuity
**Priority:** P1
**Steps:**
1. Create active session
2. Maintain acceptable risk
3. Verify session remains active
**Expected:** Acceptable risk continues session

## TC-11-006: Revoked Session Denial
**Priority:** P0
**Steps:**
1. Create session
2. Revoke session
3. Attempt access
4. Verify denied
**Expected:** Revoked session denied

## TC-11-007: Expired Session Denial
**Priority:** P0
**Steps:**
1. Create session
2. Wait for expiry
3. Attempt access
4. Verify denied
**Expected:** Expired session denied

## TC-11-008: Session Audit Events
**Priority:** P1
**Steps:**
1. Create session
2. Update session
3. Verify audit events created
**Expected:** Session events audited

## TC-11-009: Session Timeline
**Priority:** P2
**Steps:**
1. Create session
2. Perform actions
3. Verify timeline accurate
**Expected:** Session timeline accurate

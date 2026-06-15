# User Trust Test Cases

## TC-6-001: User Context Creation
**Priority:** P0
**Steps:**
1. Create user context with valid data
2. Verify context created
3. Verify context ID returned
**Expected:** User context created successfully

## TC-6-002: MFA Trust Factor
**Priority:** P1
**Steps:**
1. Create user with MFA enabled
2. Evaluate trust score
3. Verify higher score
**Expected:** MFA increases trust

## TC-6-003: Failed Login Risk Signal
**Priority:** P1
**Steps:**
1. Record failed login attempt
2. Evaluate risk score
3. Verify risk increased
**Expected:** Failed login increases risk

## TC-6-004: New Location Risk Signal
**Priority:** P1
**Steps:**
1. Record login from new location
2. Evaluate risk score
3. Verify risk increased
**Expected:** New location increases risk

## TC-6-005: Impossible Travel Risk Signal
**Priority:** P1
**Steps:**
1. Record login from distant location quickly
2. Evaluate risk score
3. Verify impossible travel detected
**Expected:** Impossible travel detected

## TC-6-006: Time-Based Risk Rule
**Priority:** P2
**Steps:**
1. Record login outside business hours
2. Evaluate risk score
3. Verify risk increased
**Expected:** Off-hours login increases risk

## TC-6-007: Risk Decay Logic
**Priority:** P1
**Steps:**
1. Record risk event
2. Wait for decay period
3. Verify risk decreased
**Expected:** Risk decays over time

## TC-6-008: User Trust Explanation
**Priority:** P2
**Steps:**
1. Create user with trust score
2. Request trust explanation
3. Verify explanation provided
**Expected:** Trust explanation generated

## TC-6-009: User Trust Audit Events
**Priority:** P1
**Steps:**
1. Create user context
2. Update trust factors
3. Verify audit events created
**Expected:** Audit events logged

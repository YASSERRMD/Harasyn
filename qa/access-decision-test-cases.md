# Access Decision Test Cases

## TC-10-001: Trusted Access Permit
**Priority:** P0
**Steps:**
1. Create trusted user
2. Create trusted device
3. Create allowed resource
4. Evaluate access
5. Verify permit decision
**Expected:** Trusted access permitted

## TC-10-002: Low Trust Device Denial
**Priority:** P0
**Steps:**
1. Create trusted user
2. Create low-trust device
3. Evaluate access
4. Verify deny decision
**Expected:** Low-trust device denied

## TC-10-003: High Risk User Denial
**Priority:** P0
**Steps:**
1. Create high-risk user
2. Create trusted device
3. Evaluate access
4. Verify deny or MFA required
**Expected:** High-risk user denied

## TC-10-004: Sensitive Resource Condition
**Priority:** P1
**Steps:**
1. Create sensitive resource
2. Evaluate access
3. Verify stronger trust required
**Expected:** Sensitive resource requires trust

## TC-10-005: Out of Hours Denial
**Priority:** P1
**Steps:**
1. Create time-based policy
2. Evaluate outside hours
3. Verify deny decision
**Expected:** Out-of-hours denied

## TC-10-006: Blocked Country Denial
**Priority:** P1
**Steps:**
1. Create geo policy
2. Evaluate from blocked country
3. Verify deny decision
**Expected:** Blocked country denied

## TC-10-007: Unknown Resource Denial
**Priority:** P0
**Steps:**
1. Attempt access to unknown resource
2. Verify deny decision
**Expected:** Unknown resource denied

## TC-10-008: Unknown User Denial
**Priority:** P0
**Steps:**
1. Attempt access with unknown user
2. Verify deny decision
**Expected:** Unknown user denied

## TC-10-009: Unknown Device Denial
**Priority:** P0
**Steps:**
1. Attempt access with unknown device
2. Verify deny decision
**Expected:** Unknown device denied

## TC-10-010: Access Decision Audit
**Priority:** P1
**Steps:**
1. Evaluate access
2. Verify audit event created
3. Verify decision recorded
**Expected:** Decision audited

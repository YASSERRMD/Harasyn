# Policy Engine Test Cases

## TC-8-001: Valid Policy Parsing
**Priority:** P0
**Steps:**
1. Submit valid policy document
2. Verify policy parsed
3. Verify policy stored
**Expected:** Valid policy accepted

## TC-8-002: Invalid Policy Rejection
**Priority:** P0
**Steps:**
1. Submit invalid policy
2. Verify validation error
3. Verify error indicates issue
**Expected:** Invalid policy rejected

## TC-8-003: Device Trust Condition
**Priority:** P1
**Steps:**
1. Create policy with device trust condition
2. Evaluate with trusted device
3. Verify condition matches
**Expected:** Device trust condition works

## TC-8-004: User Trust Condition
**Priority:** P1
**Steps:**
1. Create policy with user trust condition
2. Evaluate with trusted user
3. Verify condition matches
**Expected:** User trust condition works

## TC-8-005: Resource Sensitivity Condition
**Priority:** P1
**Steps:**
1. Create policy with sensitivity condition
2. Evaluate with sensitive resource
3. Verify condition matches
**Expected:** Sensitivity condition works

## TC-8-006: Location Condition
**Priority:** P1
**Steps:**
1. Create policy with location condition
2. Evaluate with matching location
3. Verify condition matches
**Expected:** Location condition works

## TC-8-007: Time Window Condition
**Priority:** P1
**Steps:**
1. Create policy with time condition
2. Evaluate within time window
3. Verify condition matches
**Expected:** Time condition works

## TC-8-008: MFA Condition
**Priority:** P1
**Steps:**
1. Create policy with MFA condition
2. Evaluate with MFA verified
3. Verify condition matches
**Expected:** MFA condition works

## TC-8-009: Default Deny Behavior
**Priority:** P0
**Steps:**
1. Evaluate with no matching policy
2. Verify decision is deny
3. Verify audit event created
**Expected:** Default deny works

## TC-8-010: Policy Explanation
**Priority:** P2
**Steps:**
1. Evaluate policy
2. Request explanation
3. Verify explanation accurate
**Expected:** Policy explanation provided

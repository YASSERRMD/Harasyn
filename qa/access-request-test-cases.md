# Access Request Test Cases

## TC-13-001: Access Request Creation
**Priority:** P0
**Steps:**
1. Submit access request
2. Verify request created
3. Verify request ID returned
**Expected:** Access request created

## TC-13-002: Duplicate Pending Request
**Priority:** P1
**Steps:**
1. Submit request
2. Submit duplicate
3. Verify handling
**Expected:** Duplicate handled correctly

## TC-13-003: Access Approval
**Priority:** P0
**Steps:**
1. Submit request
2. Approve request
3. Verify approved status
**Expected:** Access approved

## TC-13-004: Access Rejection
**Priority:** P0
**Steps:**
1. Submit request
2. Reject request
3. Verify rejected status
**Expected:** Access rejected

## TC-13-005: Temporary Grant Creation
**Priority:** P1
**Steps:**
1. Approve request with expiry
2. Verify grant created
3. Verify expiry set
**Expected:** Temporary grant created

## TC-13-006: Expired Grant Denial
**Priority:** P0
**Steps:**
1. Create grant with expiry
2. Wait for expiry
3. Attempt access
4. Verify denied
**Expected:** Expired grant denied

## TC-13-007: Emergency Access Reason
**Priority:** P0
**Steps:**
1. Request emergency access
2. Verify reason required
3. Verify audit logged
**Expected:** Emergency access audited

## TC-13-008: Request Status Transitions
**Priority:** P1
**Steps:**
1. Create request
2. Transition through states
3. Verify valid transitions
**Expected:** Status transitions valid

## TC-13-009: Approval Audit Events
**Priority:** P1
**Steps:**
1. Approve request
2. Verify audit event
3. Verify event details
**Expected:** Approval audited

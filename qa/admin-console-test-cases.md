# Admin Console Test Cases

## TC-15-001: Dashboard Page
**Priority:** P1
**Steps:**
1. Login to admin console
2. Navigate to dashboard
3. Verify dashboard loads
**Expected:** Dashboard loads successfully

## TC-15-002: Users Page
**Priority:** P1
**Steps:**
1. Navigate to users page
2. Verify users list loads
3. Verify pagination works
**Expected:** Users page loads

## TC-15-003: Devices Page
**Priority:** P1
**Steps:**
1. Navigate to devices page
2. Verify devices list loads
3. Verify filters work
**Expected:** Devices page loads

## TC-15-004: Resources Page
**Priority:** P1
**Steps:**
1. Navigate to resources page
2. Verify resources list loads
3. Verify search works
**Expected:** Resources page loads

## TC-15-005: Policies Page
**Priority:** P1
**Steps:**
1. Navigate to policies page
2. Verify policies list loads
3. Verify create button works
**Expected:** Policies page loads

## TC-15-006: Sessions Page
**Priority:** P1
**Steps:**
1. Navigate to sessions page
2. Verify sessions list loads
3. Verify status filters work
**Expected:** Sessions page loads

## TC-15-007: Access Requests Page
**Priority:** P1
**Steps:**
1. Navigate to access requests page
2. Verify requests list loads
3. Verify approval actions work
**Expected:** Access requests page loads

## TC-15-008: Audit Page
**Priority:** P1
**Steps:**
1. Navigate to audit page
2. Verify audit events load
3. Verify filters work
**Expected:** Audit page loads

## TC-15-009: Form Validation
**Priority:** P1
**Steps:**
1. Submit form with missing fields
2. Verify validation errors
3. Verify error messages clear
**Expected:** Forms validate correctly

## TC-15-010: Empty States
**Priority:** P2
**Steps:**
1. View page with no data
2. Verify empty state shown
3. Verify helpful message
**Expected:** Empty states handled

## TC-15-011: API Error States
**Priority:** P1
**Steps:**
1. Trigger API error
2. Verify error shown
3. Verify retry option
**Expected:** API errors shown properly

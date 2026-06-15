# Device Trust Test Cases

## TC-5-001: Device Registration
**Priority:** P0
**Steps:**
1. Register new device with valid data
2. Verify device created
3. Verify device ID returned
**Expected:** Device registered successfully

## TC-5-002: Duplicate Device Fingerprint
**Priority:** P1
**Steps:**
1. Register device with fingerprint
2. Attempt duplicate registration
3. Verify appropriate handling
**Expected:** Duplicate handled correctly

## TC-5-003: Posture Submission
**Priority:** P0
**Steps:**
1. Submit valid posture data
2. Verify posture recorded
3. Verify trust score updated
**Expected:** Posture submitted successfully

## TC-5-004: Missing Posture Validation
**Priority:** P1
**Steps:**
1. Submit posture with missing fields
2. Verify validation error
3. Verify error indicates missing fields
**Expected:** Missing fields rejected

## TC-5-005: Stale Posture Scoring
**Priority:** P1
**Steps:**
1. Submit posture
2. Wait for staleness period
3. Verify trust score decreases
**Expected:** Stale posture lowers score

## TC-5-006: Unmanaged Device Scoring
**Priority:** P1
**Steps:**
1. Register unmanaged device
2. Evaluate trust score
3. Verify lower score
**Expected:** Unmanaged gets lower score

## TC-5-007: Compliant Device Scoring
**Priority:** P1
**Steps:**
1. Register compliant device
2. Submit compliant posture
3. Verify higher score
**Expected:** Compliant gets higher score

## TC-5-008: Device Trust Explanation
**Priority:** P2
**Steps:**
1. Create device with trust score
2. Request trust explanation
3. Verify explanation provided
**Expected:** Trust explanation generated

## TC-5-009: Device Audit Events
**Priority:** P1
**Steps:**
1. Register device
2. Update device posture
3. Verify audit events created
**Expected:** Audit events logged

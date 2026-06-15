# Audit Integrity Test Cases

## TC-14-001: Access Decision Audit
**Priority:** P0
**Steps:**
1. Evaluate access decision
2. Verify audit event created
3. Verify decision recorded
**Expected:** Access decision audited

## TC-14-002: Policy Change Audit
**Priority:** P1
**Steps:**
1. Create policy
2. Update policy
3. Verify audit events created
**Expected:** Policy changes audited

## TC-14-003: Device Posture Audit
**Priority:** P1
**Steps:**
1. Register device
2. Update posture
3. Verify audit events created
**Expected:** Device posture audited

## TC-14-004: Session Revocation Audit
**Priority:** P0
**Steps:**
1. Create session
2. Revoke session
3. Verify audit event created
**Expected:** Session revocation audited

## TC-14-005: Tenant Scoped Access
**Priority:** P0
**Steps:**
1. Query audit for Tenant A
2. Verify only Tenant A events
**Expected:** Audit tenant-scoped

## TC-14-006: Audit Export
**Priority:** P1
**Steps:**
1. Request audit export
2. Verify export created
3. Verify data complete
**Expected:** Audit export works

## TC-14-007: Audit Timestamps
**Priority:** P1
**Steps:**
1. Create audit event
2. Verify timestamp accurate
**Expected:** Timestamps correct

## TC-14-008: Audit Immutability
**Priority:** P0
**Steps:**
1. Create audit event
2. Attempt modification
3. Verify modification blocked
**Expected:** Audit events immutable

## TC-14-009: Sensitive Redaction
**Priority:** P1
**Steps:**
1. Create audit with sensitive data
2. Query audit
3. Verify sensitive fields redacted
**Expected:** Sensitive fields redacted

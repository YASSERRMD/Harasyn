# Compliance Report Test Cases

## TC-19-001: Access Review Report
**Priority:** P1
**Steps:**
1. Generate access review report
2. Verify report created
3. Verify data complete
**Expected:** Access review report works

## TC-19-002: Device Compliance Report
**Priority:** P1
**Steps:**
1. Generate device compliance report
2. Verify report created
3. Verify data accurate
**Expected:** Device compliance report works

## TC-19-003: Policy Change Report
**Priority:** P2
**Steps:**
1. Generate policy change report
2. Verify report created
3. Verify changes listed
**Expected:** Policy change report works

## TC-19-004: Privileged Access Report
**Priority:** P1
**Steps:**
1. Generate privileged access report
2. Verify report created
3. Verify privileged sessions listed
**Expected:** Privileged access report works

## TC-19-005: Audit Export Download
**Priority:** P1
**Steps:**
1. Request audit export
2. Download export
3. Verify file valid
**Expected:** Audit export downloadable

## TC-19-006: Report Filters
**Priority:** P2
**Steps:**
1. Apply date filter
2. Verify filtered results
3. Apply user filter
4. Verify filtered results
**Expected:** Report filters work

## TC-19-007: Report Tenant Scope
**Priority:** P0
**Steps:**
1. Generate report for Tenant A
2. Verify only Tenant A data
**Expected:** Reports tenant-scoped

## TC-19-008: Report Data Leakage
**Priority:** P0
**Steps:**
1. Generate report
2. Verify no other tenant data
3. Verify sensitive data handled
**Expected:** No data leakage

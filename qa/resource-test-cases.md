# Resource and Connector Test Cases

## TC-7-001: Resource Registration
**Priority:** P0
**Steps:**
1. Register new resource with valid data
2. Verify resource created
3. Verify resource ID returned
**Expected:** Resource registered successfully

## TC-7-002: Resource Sensitivity Assignment
**Priority:** P1
**Steps:**
1. Create resource
2. Assign sensitivity level
3. Verify sensitivity recorded
**Expected:** Sensitivity assigned successfully

## TC-7-003: HTTP Connector Configuration
**Priority:** P1
**Steps:**
1. Configure HTTP connector
2. Verify connector created
3. Verify connector active
**Expected:** HTTP connector configured

## TC-7-004: Connector Health Check
**Priority:** P1
**Steps:**
1. Configure connector
2. Run health check
3. Verify health status
**Expected:** Health check works

## TC-7-005: Invalid Connector Validation
**Priority:** P1
**Steps:**
1. Submit invalid connector config
2. Verify validation error
3. Verify error indicates issue
**Expected:** Invalid config rejected

## TC-7-006: Disabled Resource Access
**Priority:** P0
**Steps:**
1. Disable resource
2. Attempt access
3. Verify access denied
**Expected:** Disabled resource blocked

## TC-7-007: Deleted Resource Routing
**Priority:** P0
**Steps:**
1. Delete resource
2. Attempt routing
3. Verify routing denied
**Expected:** Deleted resource not routable

## TC-7-008: Resource Audit Events
**Priority:** P1
**Steps:**
1. Create resource
2. Update resource
3. Verify audit events created
**Expected:** Audit events logged

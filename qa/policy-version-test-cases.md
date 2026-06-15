# Policy Version Test Cases

## TC-9-001: Policy Version Creation
**Priority:** P1
**Steps:**
1. Create policy
2. Update policy
3. Verify new version created
**Expected:** Version created on update

## TC-9-002: Draft Policy Isolation
**Priority:** P1
**Steps:**
1. Create draft policy
2. Evaluate access
3. Verify draft not used
**Expected:** Draft policy isolated

## TC-9-003: Published Policy Behavior
**Priority:** P0
**Steps:**
1. Publish policy
2. Evaluate access
3. Verify published policy used
**Expected:** Published policy affects decisions

## TC-9-004: Policy Rollback
**Priority:** P1
**Steps:**
1. Create policy version 1
2. Update to version 2
3. Rollback to version 1
4. Verify version 1 active
**Expected:** Rollback restores previous

## TC-9-005: Policy Diff Output
**Priority:** P2
**Steps:**
1. Create policy
2. Update policy
3. Request diff
4. Verify diff accurate
**Expected:** Policy diff accurate

## TC-9-006: Policy Conflict Detection
**Priority:** P1
**Steps:**
1. Create conflicting policies
2. Detect conflicts
3. Verify conflicts identified
**Expected:** Conflicts detected

## TC-9-007: Policy Dry Run
**Priority:** P2
**Steps:**
1. Create policy
2. Dry run evaluation
3. Verify no live changes
**Expected:** Dry run does not change live

## TC-9-008: GitOps Placeholder
**Priority:** P2
**Steps:**
1. Configure GitOps sync
2. Verify manual policy works
3. Verify no breakage
**Expected:** GitOps does not break manual

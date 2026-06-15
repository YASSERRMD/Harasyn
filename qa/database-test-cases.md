# Database Migration Test Cases

## TC-2-001: Migration Apply Flow
**Priority:** P0
**Preconditions:** Clean database
**Steps:**
1. Run all migrations from 001 to 014
2. Verify each migration completes
3. Check database state after each migration
**Expected:** All migrations apply successfully

## TC-2-002: Migration Idempotency
**Priority:** P0
**Preconditions:** Migrations already applied
**Steps:**
1. Run all migrations again
2. Verify no errors occur
3. Check database state unchanged
**Expected:** Migrations are idempotent

## TC-2-003: Required Tables Exist
**Priority:** P0
**Preconditions:** Migrations applied
**Steps:**
1. Query information_schema.tables
2. Check for required tables
3. Verify table structure
**Expected:** All required tables exist

## TC-2-004: Foreign Key Constraints
**Priority:** P1
**Preconditions:** Tables exist
**Steps:**
1. Attempt to insert child without parent
2. Verify constraint violation
3. Insert parent then child
4. Verify success
**Expected:** Foreign keys enforced

## TC-2-005: Required Field Constraints
**Priority:** P1
**Preconditions:** Tables exist
**Steps:**
1. Attempt insert with null required field
2. Verify constraint violation
3. Insert with all required fields
4. Verify success
**Expected:** Required fields enforced

## TC-2-006: Seed Data Load
**Priority:** P2
**Preconditions:** Tables exist
**Steps:**
1. Load seed data
2. Verify data exists
3. Check data integrity
**Expected:** Seed data loads successfully

## TC-2-007: Index Availability
**Priority:** P1
**Preconditions:** Tables exist
**Steps:**
1. Query pg_indexes
2. Check for required indexes
3. Verify index usage in queries
**Expected:** All required indexes exist

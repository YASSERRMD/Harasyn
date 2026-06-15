# Performance Test Cases

## TC-17-001: Access Decision Latency
**Priority:** P0
**Steps:**
1. Measure access decision time
2. Run 1000 iterations
3. Calculate p95 latency
**Expected:** Latency < 200ms

## TC-17-002: Policy Evaluation Concurrency
**Priority:** P1
**Steps:**
1. Run 100 concurrent evaluations
2. Measure throughput
3. Verify consistency
**Expected:** Handles concurrent requests

## TC-17-003: Gateway Concurrent Sessions
**Priority:** P1
**Steps:**
1. Create 1000 concurrent sessions
2. Send requests through gateway
3. Measure throughput
**Expected:** Gateway handles load

## TC-17-004: Audit Batch Writer
**Priority:** P1
**Steps:**
1. Generate 5000 audit events
2. Measure write throughput
3. Verify no events lost
**Expected:** Audit writer handles batch

## TC-17-005: Redis Cache Behavior
**Priority:** P2
**Steps:**
1. Make repeated decisions
2. Measure cache hit rate
3. Verify improvement
**Expected:** Cache improves performance

## TC-17-006: Database Index Usage
**Priority:** P1
**Steps:**
1. Run common queries
2. Verify index usage
3. Measure query time
**Expected:** Indexes used properly

## TC-17-007: Large Data Pagination
**Priority:** P2
**Steps:**
1. Query large dataset
2. Verify pagination works
3. Measure response time
**Expected:** Pagination handles large data

## TC-17-008: Worker Concurrency
**Priority:** P2
**Steps:**
1. Run multiple workers
2. Process concurrent jobs
3. Verify no duplicates
**Expected:** Worker concurrency safe

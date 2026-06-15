# Performance Testing Checklist

## Response Time Targets
- [ ] API response time < 200ms (p95)
- [ ] Policy evaluation < 50ms
- [ ] Trust score calculation < 100ms
- [ ] Gateway forwarding < 50ms
- [ ] Database queries < 100ms

## Throughput Targets
- [ ] API handles 1000 req/s
- [ ] Gateway handles 5000 req/s
- [ ] Policy engine handles 10000 eval/s
- [ ] Audit writer handles 5000 events/s

## Concurrency Targets
- [ ] 100 concurrent users
- [ ] 1000 concurrent sessions
- [ ] 100 concurrent policy evaluations
- [ ] 50 concurrent database connections

## Resource Usage
- [ ] Memory usage stable
- [ ] CPU usage within limits
- [ ] Database connection pool not exhausted
- [ ] Redis memory within limits

## Scalability
- [ ] Horizontal scaling works
- [ ] Database read replicas work
- [ ] Redis clustering works
- [ ] NATS clustering works

## Endurance
- [ ] 24-hour soak test passes
- [ ] No memory leaks detected
- [ ] No connection leaks detected
- [ ] No file handle leaks detected

## Load Test Scenarios
1. Normal load (expected traffic)
2. Peak load (2x expected traffic)
3. Stress load (5x expected traffic)
4. Spike load (sudden traffic burst)

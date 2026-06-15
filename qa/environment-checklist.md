# QA Environment Checklist

## Infrastructure Requirements

### Database
- [ ] PostgreSQL 15+ installed
- [ ] Test database created
- [ ] Migrations applied
- [ ] Test data loaded
- [ ] Connection pooling configured

### Cache
- [ ] Redis 7+ installed
- [ ] Test Redis instance running
- [ ] Connection verified
- [ ] Memory limits configured

### Message Broker
- [ ] NATS installed
- [ ] Test NATS instance running
- [ ] Connection verified
- [ ] Queue groups configured

### Application
- [ ] Backend API starts successfully
- [ ] Frontend starts successfully
- [ ] Health endpoints respond
- [ ] Environment variables loaded

## Test Data

### Users
- [ ] Test admin user created
- [ ] Test regular user created
- [ ] Test readonly user created
- [ ] Test external user created

### Devices
- [ ] Managed device registered
- [ ] Unmanaged device registered
- [ ] Compliant device registered
- [ ] Non-compliant device registered

### Resources
- [ ] HTTP resource registered
- [ ] Database resource registered
- [ ] SSH resource registered
- [ ] High sensitivity resource registered

### Policies
- [ ] Allow policy created
- [ ] Deny policy created
- [ ] Conditional policy created
- [ ] Time-based policy created

## Verification Steps
1. Run `make test`
2. Run `make lint`
3. Verify all endpoints respond
4. Verify database connectivity
5. Verify Redis connectivity

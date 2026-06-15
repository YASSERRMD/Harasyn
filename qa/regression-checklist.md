# Regression Testing Checklist

## Core Functionality

### Authentication
- [ ] User login works
- [ ] Token refresh works
- [ ] Session management works
- [ ] Logout works

### Authorization
- [ ] Role-based access works
- [ ] Policy-based access works
- [ ] Resource-level access works
- [ ] Tenant isolation works

### Device Trust
- [ ] Device registration works
- [ ] Posture assessment works
- [ ] Trust score calculation works
- [ ] Certificate validation works

### User Trust
- [ ] Context evaluation works
- [ ] Risk signal processing works
- [ ] Trust score calculation works
- [ ] Risk decay works

### Policy Engine
- [ ] Policy evaluation works
- [ ] Condition matching works
- [ ] Policy versioning works
- [ ] Policy simulation works

### Gateway
- [ ] Request forwarding works
- [ ] Token validation works
- [ ] Session enforcement works
- [ ] Timeout handling works

### Audit
- [ ] Event logging works
- [ ] Event querying works
- [ ] Export functionality works
- [ ] Integrity verification works

## Regression Test Execution
- Run full regression before each release
- Run critical path regression for hotfixes
- Document any regressions found
- Verify fixes do not introduce new issues

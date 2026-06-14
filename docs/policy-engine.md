# Policy Engine

## Harasyn Zero Trust Policy Engine

### Policy Question

```
Can user U using device D access resource R under context C right now?
```

### Policy Factors

The policy engine evaluates the following factors:

- **User Trust Score**: Is the user trustworthy given recent behavior?
- **Device Trust Score**: Is the device compliant and uncompromised?
- **Resource Sensitivity**: Is the resource publicly accessible or restricted?
- **Location Context**: Is the user accessing from an expected location?
- **Time Context**: Is access attempted during approved hours?
- **Risk Score**: Are there any elevated risk signals?
- **MFA Status**: Has the user completed multi-factor authentication?
- **Session Status**: Is the current session valid and active?

### Evaluation Flow

```
1. Parse policy document
2. Evaluate each condition against current context
3. Aggregate condition results into an overall decision
4. Output decision: ALLOW / DENY / NEEDS_APPROVAL
5. Log decision with explanation for audit
```

### Condition Types

- `device_trust >= threshold`
- `user_trust >= threshold`
- `resource_sensitivity in [public, internal]`
- `location in allowed_locations`
- `time in business_hours`
- `risk_score < max_risk`
- `mfa_status == verified`


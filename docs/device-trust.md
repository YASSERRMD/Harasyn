# Device Trust

## Harasyn Device Trust Engine

### Overview

The Device Trust Engine continuously evaluates the trustworthiness of devices attempting to access protected resources. Every device is distrusted by default.

### Device Registration

1. User or admin registers a device.
2. Device provides a unique fingerprint (hardware + software hash).
3. Device posture is evaluated:
   - OS and version
   - Encryption status
   - Jailbreak/Root status
   - Patch level
   - Last seen timestamp

### Continuous Trust Score

- **High Trust**: Compliant device, recent last-seen, no anomalies.
- **Medium Trust**: Minor deviations (e.g., slightly outdated OS).
- **Low Trust**: Significant risk (e.g., unencrypted, jailbroken).
- **Blocked**: Automatically denied access.

### Posture Evaluation

```go
type Posture struct {
    OS              string
    OSVersion       string
    Encrypted       bool
    Jailbroken      bool
    Patched         bool
    LastSeen        time.Time
}
```

### Trust Score Calculation

- Base score + compliance bonus - risk penalty = Final trust score.
- Scores range from 0 to 100.
- Access denied if score < threshold (default: 60).


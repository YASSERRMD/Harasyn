package device

import (
	"testing"
	"time"
)

func TestCalculateTrustScore(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		posture  *Posture
		device   *Device
		expected int
	}{
		{
			name: "compliant device",
			posture: &Posture{
				Encrypted:        true,
				Jailbroken:       false,
				Rooted:           false,
				Patched:          true,
				AntivirusEnabled: true,
				FirewallEnabled:  true,
			},
			device: &Device{
				LastSeenAt: &now,
			},
			expected: 100,
		},
		{
			name: "non-compliant device",
			posture: &Posture{
				Encrypted:        false,
				Jailbroken:       true,
				Rooted:           false,
				Patched:          false,
				AntivirusEnabled: false,
				FirewallEnabled:  false,
			},
			device: &Device{
				LastSeenAt: &now,
			},
			expected: 0,
		},
		{
			name: "partial compliance",
			posture: &Posture{
				Encrypted:        true,
				Jailbroken:       false,
				Rooted:           false,
				Patched:          false,
				AntivirusEnabled: false,
				FirewallEnabled:  false,
			},
			device: &Device{
				LastSeenAt: &now,
			},
			expected: 70,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := CalculateTrustScore(tt.posture, tt.device)
			if score != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, score)
			}
		})
	}
}

func TestGenerateFingerprint(t *testing.T) {
	fp1 := GenerateFingerprint("device-data-1")
	fp2 := GenerateFingerprint("device-data-1")
	fp3 := GenerateFingerprint("device-data-2")

	if fp1 != fp2 {
		t.Error("same input should produce same fingerprint")
	}
	if fp1 == fp3 {
		t.Error("different input should produce different fingerprint")
	}
}

package identity

import (
	"testing"
)

func TestCalculateUserTrustScore(t *testing.T) {
	tests := []struct {
		name        string
		mfaVerified bool
		signals     []*RiskSignal
		expected    int
	}{
		{
			name:        "mfa verified no signals",
			mfaVerified: true,
			signals:     nil,
			expected:    80,
		},
		{
			name:        "no mfa no signals",
			mfaVerified: false,
			signals:     nil,
			expected:    60,
		},
		{
			name:        "mfa verified with high risk signal",
			mfaVerified: true,
			signals: []*RiskSignal{
				{SignalType: "tor_network", Severity: "high", Score: 30},
			},
			expected: 50,
		},
		{
			name:        "no mfa multiple risk signals",
			mfaVerified: false,
			signals: []*RiskSignal{
				{SignalType: "tor_network", Severity: "high", Score: 30},
				{SignalType: "failed_logins", Severity: "high", Score: 25},
			},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := CalculateUserTrustScore(tt.mfaVerified, tt.signals)
			if score != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, score)
			}
		})
	}
}

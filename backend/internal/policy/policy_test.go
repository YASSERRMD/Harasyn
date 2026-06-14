package policy

import (
	"testing"
)

func TestEvaluatorDeviceTrust(t *testing.T) {
	evaluator := NewEvaluator()

	cond := &PolicyCondition{
		ConditionType: "device_trust",
		Operator:      ">=",
		Value:         "70",
	}

	tests := []struct {
		name   string
		score  int
		passed bool
	}{
		{"high trust", 80, true},
		{"exact match", 70, true},
		{"low trust", 60, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &EvaluationContext{DeviceTrustScore: tt.score}
			result := evaluator.evaluateCondition(cond, ctx)
			if result.Passed != tt.passed {
				t.Errorf("expected passed=%v, got %v", tt.passed, result.Passed)
			}
		})
	}
}

func TestEvaluatorResourceSensitivity(t *testing.T) {
	evaluator := NewEvaluator()

	cond := &PolicyCondition{
		ConditionType: "resource_sensitivity",
		Operator:      "in",
		Value:         "public,internal",
	}

	tests := []struct {
		name     string
		sens     string
		expected bool
	}{
		{"public", "public", true},
		{"internal", "internal", true},
		{"restricted", "restricted", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &EvaluationContext{ResourceSensitivity: tt.sens}
			result := evaluator.evaluateCondition(cond, ctx)
			if result.Passed != tt.expected {
				t.Errorf("expected passed=%v, got %v", tt.expected, result.Passed)
			}
		})
	}
}

func TestParser(t *testing.T) {
	parser := NewParser()

	validDoc := []byte(`{
		"name": "test policy",
		"effect": "allow",
		"conditions": [
			{"type": "device_trust", "operator": ">=", "value": "70"}
		]
	}`)

	doc, err := parser.Parse(validDoc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if doc.Name != "test policy" {
		t.Errorf("expected name 'test policy', got '%s'", doc.Name)
	}

	if err := parser.Validate(doc); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}
}

func TestParserInvalidEffect(t *testing.T) {
	parser := NewParser()

	invalidDoc := []byte(`{
		"name": "test",
		"effect": "invalid"
	}`)

	doc, err := parser.Parse(invalidDoc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = parser.Validate(doc)
	if err == nil {
		t.Error("expected validation error for invalid effect")
	}
}

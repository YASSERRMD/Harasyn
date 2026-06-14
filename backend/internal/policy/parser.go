package policy

import (
	"encoding/json"
	"fmt"
)

type PolicyDocument struct {
	Name        string               `json:"name"`
	Description string               `json:"description,omitempty"`
	PolicyType  string               `json:"policy_type"`
	Effect      string               `json:"effect"`
	Priority    int                  `json:"priority"`
	Conditions  []ConditionDocument  `json:"conditions"`
}

type ConditionDocument struct {
	Type     string `json:"type"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(data []byte) (*PolicyDocument, error) {
	var doc PolicyDocument
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("failed to parse policy document: %w", err)
	}

	if doc.Name == "" {
		return nil, fmt.Errorf("policy name is required")
	}
	if doc.Effect == "" {
		doc.Effect = "allow"
	}
	if doc.PolicyType == "" {
		doc.PolicyType = "access"
	}
	if doc.Priority == 0 {
		doc.Priority = 100
	}

	return &doc, nil
}

func (p *Parser) Validate(doc *PolicyDocument) error {
	if doc.Name == "" {
		return fmt.Errorf("policy name is required")
	}

	if doc.Effect != "allow" && doc.Effect != "deny" {
		return fmt.Errorf("effect must be 'allow' or 'deny'")
	}

	validConditionTypes := map[string]bool{
		"device_trust":         true,
		"user_trust":           true,
		"resource_sensitivity": true,
		"location":             true,
		"time":                 true,
		"risk_score":           true,
		"mfa_status":           true,
		"session_status":       true,
	}

	validOperators := map[string]bool{
		">=": true, "<=": true, "==": true, "!=": true,
		">": true, "<": true, "in": true, "not_in": true,
	}

	for _, cond := range doc.Conditions {
		if !validConditionTypes[cond.Type] {
			return fmt.Errorf("invalid condition type: %s", cond.Type)
		}
		if !validOperators[cond.Operator] {
			return fmt.Errorf("invalid operator: %s", cond.Operator)
		}
		if cond.Value == "" {
			return fmt.Errorf("condition value is required for %s", cond.Type)
		}
	}

	return nil
}

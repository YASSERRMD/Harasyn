package policy

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type EvaluationContext struct {
	UserTrustScore    int    `json:"user_trust_score"`
	DeviceTrustScore  int    `json:"device_trust_score"`
	ResourceSensitivity string `json:"resource_sensitivity"`
	Location          string `json:"location"`
	IsBusinessHours   bool   `json:"is_business_hours"`
	RiskScore         int    `json:"risk_score"`
	MFAVerified       bool   `json:"mfa_verified"`
	SessionActive     bool   `json:"session_active"`
}

type EvaluationResult struct {
	Allowed    bool     `json:"allowed"`
	PolicyID   string   `json:"policy_id,omitempty"`
	Reasons    []string `json:"reasons"`
	Conditions []ConditionResult `json:"conditions"`
}

type ConditionResult struct {
	ConditionType string `json:"condition_type"`
	Operator      string `json:"operator"`
	Expected      string `json:"expected"`
	Actual        string `json:"actual"`
	Passed        bool   `json:"passed"`
}

type Evaluator struct{}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func (e *Evaluator) Evaluate(policy *AccessPolicy, conditions []*PolicyCondition, ctx *EvaluationContext) (*EvaluationResult, error) {
	result := &EvaluationResult{
		Allowed:  true,
		PolicyID: policy.ID,
		Reasons:  []string{},
		Conditions: []ConditionResult{},
	}

	for _, condition := range conditions {
		condResult := e.evaluateCondition(condition, ctx)
		result.Conditions = append(result.Conditions, condResult)

		if !condResult.Passed {
			result.Allowed = false
			result.Reasons = append(result.Reasons, fmt.Sprintf(
				"condition %s %s %s failed: expected %s, got %s",
				condition.ConditionType, condition.Operator, condition.Value,
				condition.Value, condResult.Actual,
			))
		}
	}

	if policy.Effect == "deny" {
		result.Allowed = !result.Allowed
	}

	return result, nil
}

func (e *Evaluator) evaluateCondition(condition *PolicyCondition, ctx *EvaluationContext) ConditionResult {
	result := ConditionResult{
		ConditionType: condition.ConditionType,
		Operator:      condition.Operator,
		Expected:      condition.Value,
	}

	switch condition.ConditionType {
	case "device_trust":
		result.Actual = strconv.Itoa(ctx.DeviceTrustScore)
		result.Passed = e.compareInt(ctx.DeviceTrustScore, condition.Operator, condition.Value)
	case "user_trust":
		result.Actual = strconv.Itoa(ctx.UserTrustScore)
		result.Passed = e.compareInt(ctx.UserTrustScore, condition.Operator, condition.Value)
	case "resource_sensitivity":
		result.Actual = ctx.ResourceSensitivity
		result.Passed = e.compareString(ctx.ResourceSensitivity, condition.Operator, condition.Value)
	case "location":
		result.Actual = ctx.Location
		result.Passed = e.compareString(ctx.Location, condition.Operator, condition.Value)
	case "time":
		result.Actual = strconv.FormatBool(ctx.IsBusinessHours)
		result.Passed = e.compareBool(ctx.IsBusinessHours, condition.Operator, condition.Value)
	case "risk_score":
		result.Actual = strconv.Itoa(ctx.RiskScore)
		result.Passed = e.compareInt(ctx.RiskScore, condition.Operator, condition.Value)
	case "mfa_status":
		result.Actual = strconv.FormatBool(ctx.MFAVerified)
		result.Passed = e.compareBool(ctx.MFAVerified, condition.Operator, condition.Value)
	case "session_status":
		result.Actual = strconv.FormatBool(ctx.SessionActive)
		result.Passed = e.compareBool(ctx.SessionActive, condition.Operator, condition.Value)
	default:
		result.Passed = false
		result.Actual = "unknown condition type"
	}

	return result
}

func (e *Evaluator) compareInt(actual int, operator, expected string) bool {
	expVal, err := strconv.Atoi(expected)
	if err != nil {
		return false
	}

	switch operator {
	case ">=":
		return actual >= expVal
	case "<=":
		return actual <= expVal
	case "==":
		return actual == expVal
	case "!=":
		return actual != expVal
	case ">":
		return actual > expVal
	case "<":
		return actual < expVal
	default:
		return false
	}
}

func (e *Evaluator) compareString(actual, operator, expected string) bool {
	actual = strings.ToLower(actual)
	expected = strings.ToLower(expected)

	switch operator {
	case "==":
		return actual == expected
	case "!=":
		return actual != expected
	case "in":
		values := strings.Split(expected, ",")
		for _, v := range values {
			if strings.TrimSpace(v) == actual {
				return true
			}
		}
		return false
	case "not_in":
		values := strings.Split(expected, ",")
		for _, v := range values {
			if strings.TrimSpace(v) == actual {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func (e *Evaluator) compareBool(actual bool, operator, expected string) bool {
	expVal := expected == "true" || expected == "1"

	switch operator {
	case "==":
		return actual == expVal
	case "!=":
		return actual != expVal
	default:
		return false
	}
}

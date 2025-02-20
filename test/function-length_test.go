package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestFuncLengthLimitsStatements(t *testing.T) {
	testRule(t, "function-length1", &rule.FunctionLength{}, &lint.RuleConfig{
		Arguments: []any{int64(2), int64(100)},
	})
}

func TestFuncLengthLimitsLines(t *testing.T) {
	testRule(t, "function-length2", &rule.FunctionLength{}, &lint.RuleConfig{
		Arguments: []any{int64(100), int64(5)},
	})
}

func TestFuncLengthLimitsDeactivated(t *testing.T) {
	testRule(t, "function-length3", &rule.FunctionLength{}, &lint.RuleConfig{
		Arguments: []any{int64(0), int64(0)},
	})
}

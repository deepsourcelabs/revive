package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestCommentsDensity(t *testing.T) {
	testRule(t, "comments-density-1", &rule.CommentsDensityRule{}, &lint.RuleConfig{
		Arguments: []any{int64(60)},
	})

	testRule(t, "comments-density-2", &rule.CommentsDensityRule{}, &lint.RuleConfig{
		Arguments: []any{int64(90)},
	})

	testRule(t, "comments-density-3", &rule.CommentsDensityRule{}, &lint.RuleConfig{})
}

package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

func TestDatarace(t *testing.T) {
	testRule(t, "datarace", &rule.DataRaceRule{})
}

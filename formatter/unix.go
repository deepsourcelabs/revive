package formatter

import (
	"bytes"
	"fmt"

	"github.com/DeepSourceCorp/revive/lint"
)

// Unix is an implementation of the Formatter interface
// which formats the errors to a simple line based error format
//
//	main.go:24:9: [errorf] should replace errors.New(fmt.Sprintf(...)) with fmt.Errorf(...)
type Unix struct {
	Metadata lint.FormatterMetadata
}

// Name returns the name of the formatter
func (*Unix) Name() string {
	return "unix"
}

// Format formats the failures gotten from the lint.
func (*Unix) Format(failures <-chan lint.Failure, _ lint.Config) (string, error) {
	var buf bytes.Buffer
	for failure := range failures {
		fmt.Fprintf(&buf, "%v: [%s] %s\n", failure.Position.Start, failure.RuleName, failure.Failure)
	}
	return buf.String(), nil
}

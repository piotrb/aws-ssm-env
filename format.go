package main

import (
	"fmt"

	shellwords "github.com/buildkite/shellwords"
)

func formatParam(format string, name string, value string) string {
	switch format {
	case "plain":
		return fmt.Sprintf("%s=%s\n", name, value)
	case "bash":
		return fmt.Sprintf("export %s=%s\n", shellwords.Quote(name), shellwords.Quote(value))
	default:
		// fall back to plain format
		return fmt.Sprintf("%s=%s\n", name, value)
	}
}

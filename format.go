package main

import (
	"fmt"

	"gopkg.in/alessio/shellescape.v1"
)

func formatParam(format string, name string, value string) string {
	switch format {
	case "plain":
		return fmt.Sprintf("%s=%s\n", name, value)
	case "bash":
		return fmt.Sprintf("export %s=%s\n", shellescape.Quote(name), shellescape.Quote(value))
	default:
		// fall back to plain format
		return fmt.Sprintf("%s=%s\n", name, value)
	}
}

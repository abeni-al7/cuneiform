package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stderr))
}

func run(args []string, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintf(stderr, "usage: cuneiform <path-to-json-file>\n")
		return 1
	}

	if len(args) > 1 {
		fmt.Fprintf(stderr, "expected a single file path, received %d\n", len(args))
		fmt.Fprintf(stderr, "usage: cuneiform <path-to-json-file>\n")
		return 1
	}

	data, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Fprintf(stderr, "failed to read file %q: %v\n", args[0], err)
		return 1
	}

	if err := validatePlaceholder(data); err != nil {
		fmt.Fprintf(stderr, "%v\n", err)
		return 1
	}

	return 0
}

func validatePlaceholder(_ []byte) error {
	return fmt.Errorf("validation is not implemented yet")
}

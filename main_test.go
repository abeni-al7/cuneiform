package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_NoArgs_ExitsOne(t *testing.T) {
	t.Parallel()

	var stderr bytes.Buffer
	exitCode := run([]string{}, &stderr)

	if exitCode != 1 {
		t.Fatalf("expected exit code 1, got %d", exitCode)
	}

	if !strings.Contains(stderr.String(), "usage") {
		t.Fatalf("expected usage in stderr, got %q", stderr.String())
	}
}

func TestRun_TooManyArgs_ExitsOne(t *testing.T) {
	t.Parallel()

	var stderr bytes.Buffer
	exitCode := run([]string{"a.json", "b.json"}, &stderr)

	if exitCode != 1 {
		t.Fatalf("expected exit code 1, got %d", exitCode)
	}

	if !strings.Contains(stderr.String(), "single file path") {
		t.Fatalf("expected single-path guidance in stderr, got %q", stderr.String())
	}
}

func TestRun_MissingFile_ExitsOne(t *testing.T) {
	t.Parallel()

	var stderr bytes.Buffer
	exitCode := run([]string{"does-not-exist.json"}, &stderr)

	if exitCode != 1 {
		t.Fatalf("expected exit code 1, got %d", exitCode)
	}

	if !strings.Contains(stderr.String(), "failed to read file") {
		t.Fatalf("expected read failure in stderr, got %q", stderr.String())
	}
}

func TestRun_ReadableValidFile_ExitsZero(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "input.json")
	if err := os.WriteFile(path, []byte("{}"), 0o644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	var stderr bytes.Buffer
	exitCode := run([]string{path}, &stderr)

	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d with stderr %q", exitCode, stderr.String())
	}
}

func TestRun_Step1Fixtures_ValidSucceedsInvalidFails(t *testing.T) {
	t.Parallel()

	validPath := filepath.Join("tests", "step1", "valid.json")
	invalidPath := filepath.Join("tests", "step1", "invalid.json")

	t.Run("valid fixture succeeds", func(t *testing.T) {
		var stderr bytes.Buffer

		exitCode := run([]string{validPath}, &stderr)

		if exitCode != 0 {
			t.Fatalf("expected exit code 0 for %s, got %d with stderr %q", validPath, exitCode, stderr.String())
		}
	})

	t.Run("invalid fixture fails", func(t *testing.T) {
		var stderr bytes.Buffer

		exitCode := run([]string{invalidPath}, &stderr)

		if exitCode != 1 {
			t.Fatalf("expected exit code 1 for %s, got %d with stderr %q", invalidPath, exitCode, stderr.String())
		}
	})
}

func TestRun_Step2Fixtures_ValidPassesInvalidErrors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		path      string
		wantCode  int
		wantError bool
	}{
		{name: "valid", path: filepath.Join("tests", "step2", "valid.json"), wantCode: 0, wantError: false},
		{name: "valid2", path: filepath.Join("tests", "step2", "valid2.json"), wantCode: 0, wantError: false},
		{name: "invalid", path: filepath.Join("tests", "step2", "invalid.json"), wantCode: 1, wantError: true},
		{name: "invalid2", path: filepath.Join("tests", "step2", "invalid2.json"), wantCode: 1, wantError: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var stderr bytes.Buffer

			exitCode := run([]string{tc.path}, &stderr)

			if exitCode != tc.wantCode {
				t.Fatalf("expected exit code %d for %s, got %d with stderr %q", tc.wantCode, tc.path, exitCode, stderr.String())
			}

			hasErrorOutput := strings.TrimSpace(stderr.String()) != ""
			if hasErrorOutput != tc.wantError {
				t.Fatalf("expected error output=%t for %s, got %t with stderr %q", tc.wantError, tc.path, hasErrorOutput, stderr.String())
			}
		})
	}
}

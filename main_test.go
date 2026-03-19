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

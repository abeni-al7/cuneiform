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

func TestRun_Step3Fixtures_ValidPassesInvalidErrors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		path      string
		wantCode  int
		wantError bool
	}{
		{name: "valid", path: filepath.Join("tests", "step3", "valid.json"), wantCode: 0, wantError: false},
		{name: "invalid", path: filepath.Join("tests", "step3", "invalid.json"), wantCode: 1, wantError: true},
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

func TestRun_Step4Fixtures_ValidPassesInvalidErrors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		path      string
		wantCode  int
		wantError bool
	}{
		{name: "valid", path: filepath.Join("tests", "step4", "valid.json"), wantCode: 0, wantError: false},
		{name: "valid2", path: filepath.Join("tests", "step4", "valid2.json"), wantCode: 0, wantError: false},
		{name: "invalid", path: filepath.Join("tests", "step4", "invalid.json"), wantCode: 1, wantError: true},
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

func TestRun_TestFixtures_ValidPassesInvalidErrors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		path      string
		wantCode  int
		wantError bool
	}{
		{name: "pass1", path: filepath.Join("tests", "test", "pass1.json"), wantCode: 0, wantError: false},
		{name: "pass2", path: filepath.Join("tests", "test", "pass2.json"), wantCode: 0, wantError: false},
		{name: "pass3", path: filepath.Join("tests", "test", "pass3.json"), wantCode: 0, wantError: false},
		{name: "fail1", path: filepath.Join("tests", "test", "fail1.json"), wantCode: 1, wantError: true},
		{name: "fail2", path: filepath.Join("tests", "test", "fail2.json"), wantCode: 1, wantError: true},
		{name: "fail3", path: filepath.Join("tests", "test", "fail3.json"), wantCode: 1, wantError: true},
		{name: "fail4", path: filepath.Join("tests", "test", "fail4.json"), wantCode: 1, wantError: true},
		{name: "fail5", path: filepath.Join("tests", "test", "fail5.json"), wantCode: 1, wantError: true},
		{name: "fail6", path: filepath.Join("tests", "test", "fail6.json"), wantCode: 1, wantError: true},
		{name: "fail7", path: filepath.Join("tests", "test", "fail7.json"), wantCode: 1, wantError: true},
		{name: "fail8", path: filepath.Join("tests", "test", "fail8.json"), wantCode: 1, wantError: true},
		{name: "fail9", path: filepath.Join("tests", "test", "fail9.json"), wantCode: 1, wantError: true},
		{name: "fail10", path: filepath.Join("tests", "test", "fail10.json"), wantCode: 1, wantError: true},
		{name: "fail11", path: filepath.Join("tests", "test", "fail11.json"), wantCode: 1, wantError: true},
		{name: "fail12", path: filepath.Join("tests", "test", "fail12.json"), wantCode: 1, wantError: true},
		{name: "fail13", path: filepath.Join("tests", "test", "fail13.json"), wantCode: 1, wantError: true},
		{name: "fail14", path: filepath.Join("tests", "test", "fail14.json"), wantCode: 1, wantError: true},
		{name: "fail15", path: filepath.Join("tests", "test", "fail15.json"), wantCode: 1, wantError: true},
		{name: "fail16", path: filepath.Join("tests", "test", "fail16.json"), wantCode: 1, wantError: true},
		{name: "fail17", path: filepath.Join("tests", "test", "fail17.json"), wantCode: 1, wantError: true},
		{name: "fail18", path: filepath.Join("tests", "test", "fail18.json"), wantCode: 1, wantError: true},
		{name: "fail19", path: filepath.Join("tests", "test", "fail19.json"), wantCode: 1, wantError: true},
		{name: "fail20", path: filepath.Join("tests", "test", "fail20.json"), wantCode: 1, wantError: true},
		{name: "fail21", path: filepath.Join("tests", "test", "fail21.json"), wantCode: 1, wantError: true},
		{name: "fail22", path: filepath.Join("tests", "test", "fail22.json"), wantCode: 1, wantError: true},
		{name: "fail23", path: filepath.Join("tests", "test", "fail23.json"), wantCode: 1, wantError: true},
		{name: "fail24", path: filepath.Join("tests", "test", "fail24.json"), wantCode: 1, wantError: true},
		{name: "fail25", path: filepath.Join("tests", "test", "fail25.json"), wantCode: 1, wantError: true},
		{name: "fail26", path: filepath.Join("tests", "test", "fail26.json"), wantCode: 1, wantError: true},
		{name: "fail27", path: filepath.Join("tests", "test", "fail27.json"), wantCode: 1, wantError: true},
		{name: "fail28", path: filepath.Join("tests", "test", "fail28.json"), wantCode: 1, wantError: true},
		{name: "fail29", path: filepath.Join("tests", "test", "fail29.json"), wantCode: 1, wantError: true},
		{name: "fail30", path: filepath.Join("tests", "test", "fail30.json"), wantCode: 1, wantError: true},
		{name: "fail31", path: filepath.Join("tests", "test", "fail31.json"), wantCode: 1, wantError: true},
		{name: "fail32", path: filepath.Join("tests", "test", "fail32.json"), wantCode: 1, wantError: true},
		{name: "fail33", path: filepath.Join("tests", "test", "fail33.json"), wantCode: 1, wantError: true},
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

//go:build windows

package main

import (
	"flag"
	"os"
	"testing"
	"time"
)

func TestInitFlags(t *testing.T) {
	// Save original command line and reset flags
	originalArgs := os.Args
	originalCommandLine := flag.CommandLine

	defer func() {
		os.Args = originalArgs
		flag.CommandLine = originalCommandLine
	}()

	// Create a new flag set for this test
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Test initFlags() function
	cfg := initFlags()

	// Test default values
	if cfg.help != false {
		t.Errorf("Expected help default to be false, got %v", cfg.help)
	}
	if cfg.version != false {
		t.Errorf("Expected version default to be false, got %v", cfg.version)
	}

	// Test that flags can be parsed
	testArgs := []string{
		"progname",
		"-?",
		"-v",
	}

	// Reset flag set and reinitialize
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	cfg = initFlags()

	// Parse test arguments
	err := flag.CommandLine.Parse(testArgs[1:])
	if err != nil {
		t.Fatalf("Failed to parse flags: %v", err)
	}

	// Verify flags were set correctly
	if !cfg.version {
		t.Error("Expected version flag to be true")
	}
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		in   string
		want time.Duration
	}{
		{"0", 0 * time.Second},
		{"5", 5 * time.Second}, // auto "s"
		{"2.5", 2500 * time.Millisecond},
		{"1m", time.Minute},
		{"750ms", 750 * time.Millisecond},
	}

	for _, tc := range tests {
		got := parseDuration(tc.in)
		if got != tc.want {
			t.Errorf("parseDuration(%q) = %v; want %v", tc.in, got, tc.want)
		}
	}
}

package main

import (
	"os"
	"strconv"
	"testing"
)

func TestDetectOpencodeVersion(t *testing.T) {
	expected := os.Getenv("EXPECTED_OPENCODE_VERSION")
	if expected == "" {
		t.Skip("EXPECTED_OPENCODE_VERSION not set (only set in CI compat workflow)")
	}
	ver, err := detectOpencodeVersion()
	if err != nil {
		t.Fatalf("detectOpencodeVersion() failed: %v", err)
	}
	if strconv.Itoa(ver) != expected {
		t.Fatalf("expected opencode version %s, got %d", expected, ver)
	}
}

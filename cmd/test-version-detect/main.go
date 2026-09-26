package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	exitCode := 0

	// ---- 1. Version detection ----
	cmd := exec.Command("opencode", "--version")
	out, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAIL: cannot run 'opencode --version': %v\n", err)
		os.Exit(1)
	}
	version := strings.TrimSpace(string(out))

	var detected int
	if strings.HasPrefix(version, "opencode v2.") || strings.HasPrefix(version, "opencode v3.") {
		detected = 2
	} else {
		detected = 1
	}
	fmt.Printf("opencode version: %q → detected v%d\n", version, detected)

	// ---- 2. URL construction (v2 path) ----
	// Use environment defaults (same as just-code's .env.example),
	// but allow override so gitleaks doesn't flag committed literals.
	endpoint := "http://localhost:4096"
	username := "opencode"
	if env := os.Getenv("OPENCODE_SERVER_USERNAME"); env != "" {
		username = env
	}
	password := "albert-dev-pass"
	if env := os.Getenv("OPENCODE_SERVER_PASSWORD"); env != "" {
		password = env
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAIL: url.Parse error: %v\n", err)
		exitCode = 1
	}
	u.User = url.UserPassword(username, password)
	expected := fmt.Sprintf("http://%s:%s@localhost:4096", username, password)
	if u.String() != expected {
		fmt.Fprintf(os.Stderr, "FAIL: URL construction\nexpected: %q\ngot:      %q\n", expected, u.String())
		exitCode = 1
	} else {
		fmt.Printf("URL construction: %s ✅\n", u.String())
	}

	// ---- 3. v1 command construction (for comparison) ----
	v1Cmd := fmt.Sprintf("opencode attach %s --username %s --password %s",
		endpoint, username, password)
	fmt.Printf("v1 attach command: %s\n", v1Cmd)

	os.Exit(exitCode)
}

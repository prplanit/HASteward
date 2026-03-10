package style

import (
	"os"
	"strings"
)

// colorOverride is set by the --no-color flag. When non-nil, it overrides
// automatic detection.
var colorOverride *bool

// SetColorEnabled explicitly enables or disables color output.
// This is called from cmd/root.go when --no-color is specified.
func SetColorEnabled(enabled bool) {
	colorOverride = &enabled
}

// IsTTY returns true if the given file is a terminal.
func IsTTY(f *os.File) bool {
	info, err := f.Stat()
	if err != nil {
		return false
	}
	return (info.Mode() & os.ModeCharDevice) != 0
}

// ColorEnabled returns true if color output should be used.
// Priority: explicit override (--no-color) > NO_COLOR env > TERM=dumb > TTY detection.
func ColorEnabled() bool {
	if colorOverride != nil {
		return *colorOverride
	}
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		return false
	}
	if strings.ToLower(os.Getenv("TERM")) == "dumb" {
		return false
	}
	return IsTTY(os.Stdout)
}

// IsCI returns true if running in a CI environment.
func IsCI() bool {
	for _, key := range []string{"CI", "GITHUB_ACTIONS", "GITLAB_CI", "JENKINS_URL", "BUILDKITE"} {
		if os.Getenv(key) != "" {
			return true
		}
	}
	return false
}

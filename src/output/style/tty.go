package style

import (
	"os"
	"strings"
)

// IsTTY returns true if the given file is a terminal.
func IsTTY(f *os.File) bool {
	info, err := f.Stat()
	if err != nil {
		return false
	}
	return (info.Mode() & os.ModeCharDevice) != 0
}

// ColorEnabled returns true if color output should be used.
// Respects NO_COLOR (https://no-color.org/), TERM=dumb, and CI environments.
func ColorEnabled() bool {
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

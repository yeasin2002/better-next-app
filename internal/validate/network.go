package validate

import (
	"net"
	"os"
	"strings"
	"time"
)

// IsCI detects if running in a CI environment
func IsCI() bool {
	ciEnvVars := []string{
		"CI",
		"CONTINUOUS_INTEGRATION",
		"GITHUB_ACTIONS",
		"GITLAB_CI",
		"CIRCLECI",
		"TRAVIS",
		"JENKINS_URL",
		"BUILDKITE",
		"DRONE",
		"TEAMCITY_VERSION",
	}

	for _, envVar := range ciEnvVars {
		if val := os.Getenv(envVar); val != "" && strings.ToLower(val) != "false" {
			return true
		}
	}

	return false
}

// IsOnline checks network connectivity
func IsOnline() bool {
	// Try to resolve a well-known domain
	conn, err := net.DialTimeout("tcp", "registry.npmjs.org:443", 3*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

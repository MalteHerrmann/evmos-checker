package checker

// Config is the configuration for the Checker.
type Config struct {
	// checks is a slice of all checks to run.
	checks []Check
}

// NewConfig creates a new Config instance.
func NewConfig(checks []Check) Config {
	return Config{
		checks: checks,
	}
}

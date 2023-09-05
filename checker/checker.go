package checker

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Checker is the main type, which contains all the necessary information to
// check a work-in-progress branch during Evmos development.
//
// It is used to coordinate the terminal UI,
// as well as the logic, which runs the necessary checks.
type Checker struct {
	// checks is a slice of all ran checks.
	checks []Check
	// logger is the logger instance.
	logger *zerolog.Logger
}

// NewChecker creates a new Checker instance.
func NewChecker() *Checker {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	return &Checker{
		logger: &logger,
	}
}

// Run starts the checker.
func (c *Checker) Run() error {
	checkNames := make([]string, 0, len(c.checks))
	for _, check := range c.checks {
		checkNames = append(checkNames, check.name)
	}

	c.logger.Info().Msg("Starting checker...")
	c.logger.Info().Msg(
		fmt.Sprintf("Running %d checks: %s", len(checkNames), strings.Join(checkNames, ", ")),
	)

	return nil
}

// GetChecks returns all checks, which are currently registered.
func (c *Checker) GetChecks() []Check {
	return c.checks
}

// Check is a single check, which can be run by the checker.
type Check struct {
	// name is the name of the check.
	name string
}

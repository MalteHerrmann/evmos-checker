package checker

import "github.com/rs/zerolog/log"

// Checker is the main type, which contains all the necessary information to
// check a work-in-progress branch during Evmos development.
//
// It is used to coordinate the terminal UI,
// as well as the logic, which runs the necessary checks.
type Checker struct {
	// checks is a slice of all ran checks.
	checks []Check
}

// NewChecker creates a new Checker instance.
func NewChecker() *Checker {
	return &Checker{}
}

// Start starts the checker.
func (c *Checker) Start() {
	log.Print("Starting the checker...")
}

// Check is a single check, which can be run by the checker.
type Check struct {
	// name is the name of the check.
	name string
}

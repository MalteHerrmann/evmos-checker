package checker

// Check is a single check, which can be run by the checker.
type Check struct {
	// name is the name of the check.
	name string
}

// NewCheck creates a new Check instance.
func NewCheck(name string) Check {
	return Check{
		name: name,
	}
}

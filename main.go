package main

import (
	"github.com/MalteHerrmann/evmos-checker/checker"
	"github.com/pkg/errors"
)

func main() {
	// Run the checker
	evmosChecker := checker.NewChecker()

	if err := evmosChecker.Run(); err != nil {
		panic(errors.Wrap(err, "failed to run checker"))
	}
}

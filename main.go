/*
Evmos Checker | Malte Herrmann | Evmos Core Team

This tool is a collection of checks that are run during development
of the Evmos core protocol.
*/
package main

import (
	"github.com/MalteHerrmann/evmos-checker/cmd"
	"github.com/MalteHerrmann/evmos-checker/checker"
	"github.com/pkg/errors"
)

func main() {
	cmd.Execute()

  // Run the checker
	evmosChecker := checker.NewChecker()

	if err := evmosChecker.Run(); err != nil {
		panic(errors.Wrap(err, "failed to run checker"))
	}
}

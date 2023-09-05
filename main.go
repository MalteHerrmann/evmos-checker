/*
Evmos Checker | Malte Herrmann | Evmos Core Team

This tool is a collection of checks that are run during development
of the Evmos core protocol.
*/
package main

import (
	"github.com/MalteHerrmann/evmos-checker/cmd"
)

func main() {
	// Execute the cobra CLI binary.
	cmd.Execute()
}

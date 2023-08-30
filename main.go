package main

import (
	"github.com/MalteHerrmann/evmos-checker/checker"
)

func main() {
	// Start the checker
	evmosChecker := checker.NewChecker()
	evmosChecker.Start()
}

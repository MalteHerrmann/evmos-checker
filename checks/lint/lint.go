/*
Package lint contains the various linter checks.
*/
package lint

import (
	"github.com/MalteHerrmann/evmos-checker/checker"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// lintShortDesc is the long description of the linting sub-checks.
const lintLongDesc = `Runs the linter checks.

Included linters are:
 - gofumpt
 - golangci-lint
 - markdownlint
 - markdown-link-check`

// NewLintCmd returns the root command of the evmos-checker linting subchecks.
func NewLintCmd() *cobra.Command {
	lintCmd := &cobra.Command{
		Use:   "lint",
		Short: "Runs the linter checks",
		Long:  lintLongDesc,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: pass lint checks into constructor
			lintChecks := NewLintChecks()
			evmosChecker := checker.NewChecker(
				checker.NewConfig(lintChecks),
			)

			err := evmosChecker.Run()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to run evmos-checker")
			}
		},
	}

	return lintCmd
}

// NewLintChecks returns a slice of all linting checks.
func NewLintChecks() []checker.Check {
	return []checker.Check{
		checker.NewCheck("make lint"),
		checker.NewCheck("make proto-lint"),
	}
}

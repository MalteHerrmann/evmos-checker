/*
Package lint contains the various linter checks.
*/
package lint

import (
	"fmt"

	"github.com/spf13/cobra"
)

// LintCmd represents the CLI command for linting.
var LintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Runs the linter checks",
	Long: `Runs the linter checks.

Included linters are:
 - gofumpt
 - golangci-lint
 - markdownlint
 - markdown-link-check`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lint called")
	},
}

// NewLintCmd returns the root command of the evmos-checker linting subchecks.
func NewLintCmd() *cobra.Command {
	// Define flags and configuration settings here.

	return LintCmd
}

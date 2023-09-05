/*
Package cmd contains the root command of the evmos-checker tool.
*/
package cmd

import (
	"os"

	"github.com/MalteHerrmann/evmos-checker/checker"
	lintchecks "github.com/MalteHerrmann/evmos-checker/checks/lint"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "evmos-checker",
	Short: "A collection of Evmos related checks",
	Long: `Evmos Checker | Malte Herrmann | Evmos Core Team

This tool is a collection of checks that are run during development of the Evmos core protocol.`,
	Run: func(cmd *cobra.Command, args []string) {
		evmosChecker := checker.NewChecker()

		err := evmosChecker.Run()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to run evmos-checker")
		}
	},
}

// NewRootCmd returns the root command of the evmos-checker tool.
func NewRootCmd() *cobra.Command {
	rootCmd.AddCommand(lintchecks.NewLintCmd())

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	root := NewRootCmd()
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

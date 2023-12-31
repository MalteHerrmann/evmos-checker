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

// rootLongDesc is the long description of the root command.
const rootLongDesc = `Evmos Checker | Malte Herrmann | Evmos Core Team

This tool is a collection of checks that are run during development of the Evmos core protocol.`

// NewRootCmd returns the root command of the evmos-checker tool.
func NewRootCmd() *cobra.Command {
	rootCmd := newRootCmdRaw()

	lintCmd := lintchecks.NewLintCmd()
	rootCmd.AddCommand(lintCmd)

	return rootCmd
}

// newRootCmdRaw returns the root command of the evmos-checker tool without any subcommands or flags configured.
func newRootCmdRaw() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "evmos-checker",
		Short: "A collection of Evmos related checks",
		Long:  rootLongDesc,
		Run: func(cmd *cobra.Command, args []string) {
			evmosChecker := checker.NewChecker(checker.Config{})

			err := evmosChecker.Run()
			if err != nil {
				// TODO: replace with context logger
				log.Fatal().Err(err).Msg("failed to run evmos-checker")
			}
		},
	}

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

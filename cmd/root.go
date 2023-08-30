/*
Package cmd contains the root command of the evmos-checker tool.
*/
package cmd

import (
	"os"

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
		log.Info().Msg("started evmos-checker")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.evmos-checks.yaml)")
}

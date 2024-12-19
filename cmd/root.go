/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"jcrose-copilot/pkg/constants"
	"jcrose-copilot/pkg/option"

	"github.com/spf13/cobra"
)

var copilotOpt option.CopilotOption
var verbose string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jcrose-copilot",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	copilotCmd.Flags().StringVarP(&verbose, "verbose", "v", "", "")
	copilotCmd.Flags().StringVarP(&copilotOpt.Endpoint, "endpoint", "e", "", "e.g. https://api.openai.com/v1")
	copilotCmd.Flags().StringVarP(&copilotOpt.Model, "model", "m", "", "e.g. gpt-3.5-turbo")
	copilotCmd.Flags().StringVarP(&copilotOpt.Key, "key", "k", "", "e.g. sk-xxx")
	copilotCmd.Flags().IntVarP(&copilotOpt.History, "history", "", 5, "")
	copilotCmd.Flags().BoolVarP(&copilotOpt.Silence, "silence", "s", false, "")
	copilotCmd.Flags().StringVarP(&copilotOpt.OpsServer, "opsserver", "", "", "")
	copilotCmd.Flags().StringVarP(&copilotOpt.OpsToken, "opstoken", "", "", "")
	copilotCmd.Flags().StringVarP(&copilotOpt.RuntimeImage, "runtimeimage", "", constants.DefaultRuntimeImage, "e.g. ubuntu:22.04")
}

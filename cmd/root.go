package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// Remove the 'completion' command
	var completionCmd *cobra.Command
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "completion" {
			completionCmd = cmd
			break
		}
	}
	if completionCmd != nil {
		rootCmd.RemoveCommand(completionCmd)
	}
}

var rootCmd = &cobra.Command{
	Use:   "solidgo",
	Short: "CLI Tool for SolidGO projects",
	Long: `Tool to provide a CLI interface for SolidGO projects, 
including initialing routes, models, and controllers for the server 
models, components, pages, and services for the client.`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.solidgo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

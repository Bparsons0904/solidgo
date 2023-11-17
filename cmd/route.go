package cmd

import (
	"github.com/spf13/cobra"
)

var routeCmd = &cobra.Command{
	Use:     "route",
	Aliases: []string{"r"},
	Short:   "Manage routes",
	Long:    `Subcommands for managing routes.`,
}

func init() {
	routeCmd.PersistentFlags().StringP("method", "m", "", "HTTP method for the route (GET, POST, etc.)")
	routeCmd.PersistentFlags().StringP("function", "f", "", "Function name to handle the route")
	routeCmd.PersistentFlags().StringP("path", "p", "", "Path to the route file after '/'")

	rootCmd.AddCommand(routeCmd)
}

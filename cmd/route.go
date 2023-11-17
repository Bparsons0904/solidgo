/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// routeCmd represents the route command
var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "Adds a new route to the server.",
	Long: `Creates a new route in the server's routes directory.
Includes a new controller to match the route.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("route called")
	},
}

func init() {
	rootCmd.AddCommand(routeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

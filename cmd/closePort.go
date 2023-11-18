package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var closePortCmd = &cobra.Command{
	Use:     "closeport",
	Aliases: []string{"cp"},
	Short:   "Closes the specified open ports",
	Long:    `Closes one or more ports by terminating the processes that are listening on these ports.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, port := range args {
			err := closePort(port)
			if err != nil {
				fmt.Printf("Failed to close port %s: %v\n", port, err)
			} else {
				fmt.Printf("Successfully closed port %s\n", port)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(closePortCmd)
}

func closePort(port string) error {
	out, err := exec.Command("lsof", "-i", fmt.Sprintf(":%s", port)).Output()
	if err != nil {
		return err
	}

	lines := strings.Split(string(out), "\n")
	if len(lines) > 1 {
		columns := strings.Fields(lines[1])
		if len(columns) > 1 {
			pid := columns[1]
			return exec.Command("kill", pid).Run()
		}
	}
	return fmt.Errorf("no process found listening on port %s", port)
}

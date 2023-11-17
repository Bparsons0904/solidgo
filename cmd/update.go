package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Builds and installs the latest version of SolidGO.",
	Long:  `Runs the build command, then installs the latest version of SolidGO.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildCmd := exec.Command("go", "build", "-o", "solidgo")
		buildOutput, err := buildCmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Failed to execute 'go build': %s\nError: %s", buildOutput, err)
		}
		fmt.Println("Go Build Ran Successfully")

		installCmd := exec.Command("go", "install")
		installOutput, err := installCmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Failed to execute 'go install': %s\nError: %s", installOutput, err)
		}
		fmt.Println("Go Install Ran Successfully")
	},
	DisableFlagParsing: true,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

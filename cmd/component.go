package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var componentCmd = &cobra.Command{
	Use:     "component [path] [name]",
	Short:   "Create a SolidJS component",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path, name := args[0], args[1]
		noModule, _ := cmd.Flags().GetBool("noModule")

		err := createSolidJSComponent(path, name, noModule)
		if err != nil {
			fmt.Println("Error creating component:", err)
		} else {
			fmt.Println("Component created successfully")
		}
	},
}

func init() {
	componentCmd.Flags().Bool("noModule", false, "Set this flag to skip creating the module.scss file")
	rootCmd.AddCommand(componentCmd)
}

func createSolidJSComponent(componentPath, componentName string, noModule bool) error {
	dirPath := filepath.Join("client/src", componentPath, componentName)
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return err
	}

	tsxFilePath := filepath.Join(dirPath, componentName+".tsx")
	tsxContent := fmt.Sprintf(`import styles from "./%s.module.scss";

import { Component } from "solid-js";

interface %sProps {}

export const %s: Component<%sProps> = (props) => {
  return <div class={styles.root}>%s Works!</div>;
};
`, componentName, componentName, componentName, componentName, componentName)
	err = os.WriteFile(tsxFilePath, []byte(tsxContent), 0644)
	if err != nil {
		return err
	}

	if !noModule {
		scssFilePath := filepath.Join(dirPath, componentName+".module.scss")
		scssContent := `.root {}`
		err = os.WriteFile(scssFilePath, []byte(scssContent), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

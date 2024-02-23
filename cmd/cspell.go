package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

// addwordCmd represents the command to add a word to cspell.json
var addwordCmd = &cobra.Command{
	Use:   "addword [word]",
	Short: "Add a word to the cspell.json configuration",
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is passed
	Run: func(cmd *cobra.Command, args []string) {
		word := args[0]

		// Dynamically resolve the path to the user's home directory
		usr, err := user.Current()
		if err != nil {
			fmt.Printf("Error obtaining user home directory: %v\n", err)
			return
		}
		homeDir := usr.HomeDir

		// Specify the path to your cspell.json file within the LunarVim configuration directory
		cspellPath := filepath.Join(homeDir, ".config", "lvim", "cspell.json")

		err = addWordToCSpell(cspellPath, word)
		if err != nil {
			fmt.Printf("Error adding word to cspell.json: %v\n", err)
			return
		}

		fmt.Println("Word added successfully to cspell.json")
	},
}

func init() {
	rootCmd.AddCommand(addwordCmd) // Make sure to replace rootCmd with your actual root command variable if it's named differently
}

// addWordToCSpell adds a word to the cspell.json configuration
func addWordToCSpell(cspellPath, word string) error {
	var config struct {
		Words []string `json:"words"`
	}

	// Read the existing cspell.json file
	content, err := os.ReadFile(cspellPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("cspell.json does not exist, creating a new one.")
		} else {
			return err
		}
	} else {
		// Parse the JSON content
		err = json.Unmarshal(content, &config)
		if err != nil {
			return err
		}
	}

	// Check if the word is already in the list
	for _, w := range config.Words {
		if w == word {
			fmt.Println("Word already exists in cspell.json")
			return nil
		}
	}

	// Add the word to the list
	config.Words = append(config.Words, word)

	// Marshal the updated configuration back to JSON
	updatedContent, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// Write the updated JSON back to cspell.json
	err = os.WriteFile(cspellPath, updatedContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

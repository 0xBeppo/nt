/*
Copyright © 2024 Markel Elorza 0xBeppo<beppo.dev.io@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// todoCmd represents the todo command
var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Create quick note for TODOs",
	Long: `Create quick notes for TODOs
You can create a new TODO note by name, or just use todays date as default
if the note exists, it will open it with neovim`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo called")
	},
}

func init() {
	homeDir, _ = os.UserHomeDir()
	fileName = GetTodaysDate()
	rootCmd.AddCommand(todoCmd)
	// TODO: Add tags as flag for every child
}

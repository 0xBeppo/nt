/*
Copyright © 2024 Markel Elorza 0xBeppo<beppo.dev.io@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/charmbracelet/log"
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
		noteType = "todo"
		log.Debugf("todo called")
		EnableVerbose(isVerbose)
		if len(args) == 0 {
			OpenTeaUi()
		} else {
			fileName = args[0]
		}
		EnsureExtension("-todolist.md")
		log.Debugf("%s", fileName)
		noteName := GetNoteName()
		// check that is new
		exists := CheckIfNoteExists(noteName)
		if !exists {
			t := CreateTemplate(TodoTemplate)
			WriteNote(t, noteName, fileName)
		}
		OpenNewNote(noteName)
	},
}

func init() {
	homeDir, _ = os.UserHomeDir()
	fileName = GetTodaysDate()
	rootCmd.AddCommand(todoCmd)
}

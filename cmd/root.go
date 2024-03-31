/*
Copyright © 2024 Markel Elorza 0xBeppo<beppo.dev.io@gmail.com>
*/
package cmd

import (
	"os"
	"strings"
	"text/template"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	fileName, homeDir string
	isVerbose         bool
	tags              []string
)

const filePath = "notebook/"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nt",
	Short: "Note taking app for personal use",
	Long: `Note taking app for personal use with different use cases:

You can create quick notes, todo notes, meeting notes, weekly meeting notes,
parse and organize lastly created notes by their tags, etc.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		EnableVerbose(isVerbose)
		if len(args) == 0 {
			OpenTeaUi()
		} else {
			fileName = args[0]
		}
		EnsureExtension()
		log.Debugf("%s", fileName)
		noteName := GetNoteName()
		// check that is new
		exists := CheckIfNoteExists(noteName)
		if !exists {
			t := CreateTemplate("note.tmpl")
			writeTemplate(t, noteName, fileName)
		}
		OpenNewNote(noteName)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	homeDir, _ = os.UserHomeDir()
	fileName = GetTodaysDate()
	rootCmd.PersistentFlags().BoolVarP(&isVerbose, "verbose", "v", false, "Enable verbose mode")
	rootCmd.PersistentFlags().StringArrayVarP(&tags, "tags", "t", []string{}, "Tags for the new note")
}

// TODO: Modify to create only main notes
func writeTemplate(t *template.Template, noteName string, fileName string) {
	title, _ := strings.CutSuffix(fileName, ".md")
	file, err := os.Create(noteName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	notes := MyNote{
		Title: title,
		Date:  GetTodaysDate(),
		Tags:  tags,
	}
	err = t.Execute(file, notes)
	if err != nil {
		panic(err)
	}
}

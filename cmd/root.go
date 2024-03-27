/*
Copyright © 2024 Markel Elorza 0xBeppo<beppo.dev.io@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var fileName, homeDir string

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
		if len(args) == 0 {
			p := tea.NewProgram(initialModel())
			if _, err := p.Run(); err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}
		} else {
			fileName = args[0]
		}
		ensureExtension()
		log.Infof("%s", fileName)
		noteName := getNoteName()
		createNewNote(noteName)
	},
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
	homeDir, _ = os.UserHomeDir()
	todaysdate := time.Now()
	rootCmd.Flags().StringVarP(&fileName, "name", "n", todaysdate.Format("2006-01-02"), "Name for the new note, without extension")
}

func getNoteName() string {
	var sb strings.Builder
	sb.WriteString(homeDir + "/")
	sb.WriteString(filePath)

	err := os.Mkdir(sb.String(), 07500)
	if err != nil && !os.IsExist(err) {
		log.Errorf("Prueba: %s", err)
	}
	sb.WriteString(fileName)
	log.Infof("Created file: %s", sb.String())

	return sb.String()
}

func createNewNote(note string) {
	log.Infof("Opening: %s", note)
	command := exec.Command("nvim", note)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	err := command.Run()
	fmt.Println(err)
}

func ensureExtension() {
	if !strings.HasSuffix(fileName, ".md") {
		fileName += ".md"
	}
}

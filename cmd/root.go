/*
Copyright © 2024 Markel Elorza 0xBeppo<beppo.dev.io@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
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
		t := createTemplate("note.tmpl")
		writeTemplate(t, noteName)
		openNewNote(noteName)
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
	date := getTodaysDate()
	rootCmd.Flags().StringVarP(&fileName, "name", "n", date, "Name for the new note, without extension will prompt for it")
}

func getTodaysDate() string {
	todaysdate := time.Now()
	return todaysdate.Format("2006-01-02")
}

func createTemplate(tmpl string) *template.Template {
	t, err := template.New(tmpl).ParseFiles(tmpl)
	if err != nil {
		panic(err)
	}

	return t
}

func writeTemplate(t *template.Template, noteName string) {
	file, err := os.OpenFile(noteName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	notes := MyNote{
		Title: noteName,
		Date:  getTodaysDate(),
		Tags:  []string{},
	}
	err = t.Execute(file, notes)
	if err != nil {
		panic(err)
	}
}

func getNoteName() string {
	var sb strings.Builder
	sb.WriteString(homeDir + "/")
	sb.WriteString(filePath)

	err := os.Mkdir(sb.String(), 0750)
	if err != nil && !os.IsExist(err) {
		log.Errorf("Prueba: %s", err)
	}
	sb.WriteString(fileName)
	log.Infof("Created file: %s", sb.String())

	return sb.String()
}

func openNewNote(note string) {
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

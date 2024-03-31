package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

func GetTodaysDate() string {
	return time.Now().Format("2006-01-02")
}

func CheckIfNoteExists(noteName string) bool {
	if _, err := os.Stat(noteName); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func CreateTemplate(tmpl string) *template.Template {
	t, err := template.New(tmpl).ParseFiles(tmpl)
	if err != nil {
		panic(err)
	}

	return t
}

func GetNoteName() string {
	var sb strings.Builder
	sb.WriteString(homeDir + "/")
	sb.WriteString(filePath)

	err := os.Mkdir(sb.String(), 0750)
	if err != nil && !os.IsExist(err) {
		log.Errorf("Prueba: %s", err)
	}
	sb.WriteString(fileName)
	log.Debugf("Created file: %s", sb.String())

	return sb.String()
}

func OpenNewNote(note string) {
	log.Debugf("Opening: %s", note)
	command := exec.Command("nvim", note)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	err := command.Run()
	fmt.Println(err)
}

func EnsureExtension() {
	if !strings.HasSuffix(fileName, ".md") {
		fileName += ".md"
	}
}

func EnableVerbose(isVerbose bool) {
	if isVerbose {
		log.SetLevel(log.DebugLevel)
	}
}

func OpenTeaUi() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

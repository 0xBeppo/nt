package cmd

import (
	"bufio"
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
	t, err := template.New(tmpl).Parse(tmpl)
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

func EnsureExtension(extension string) {
	if !strings.HasSuffix(fileName, extension) {
		fileName += extension
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

func WriteNote(t *template.Template, noteName string, fileName string) {
	var oldTasks []string
	notesPath := homeDir + "/" + filePath
	title, _ := strings.CutSuffix(fileName, ".md")
	if strings.Contains(fileName, "-todolist.md") {
		lastTodoNote, err := getLastModifiedFileWithSuffix(notesPath, "-todolist.md")
		if err != nil {
			panic(err)
		}
		if lastTodoNote != "" {
			oldTasks, err = getUncompletedTasksFromNote(notesPath + lastTodoNote)
			if err != nil {
				panic(err)
			}
		}
	}
	file, err := os.Create(noteName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	notes := MyNote{
		Title:    title,
		Date:     GetTodaysDate(),
		Tags:     tags,
		OldTasks: oldTasks,
	}
	err = t.Execute(file, notes)
	if err != nil {
		panic(err)
	}
}

func getLastModifiedFileWithSuffix(directory string, suffix string) (string, error) {
	var lastModifiedFile string
	var lastModifiedTime time.Time

	files, err := os.ReadDir(directory)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), suffix) {
			info, err := file.Info()
			if err != nil {
				return "", err
			}

			if info.ModTime().After(lastModifiedTime) {
				lastModifiedFile = file.Name()
				lastModifiedTime = info.ModTime()
			}
		}
	}

	if lastModifiedFile == "" {
		return "", nil
	}

	return lastModifiedFile, nil
}

func getUncompletedTasksFromNote(note string) ([]string, error) {
	file, err := os.Open(note)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "- [ ]") {
			tasks = append(tasks, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

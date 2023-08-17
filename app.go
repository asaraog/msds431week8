package main

import (
	"context"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Lint(pref string) []string {
	//Read in User preference for datum (1) or data (2) and checks if incorrect. Program will not ask for text file unless correct.
	if pref != "1" && pref != "2" {
		return []string{"Incorrect user input", "Please enter 1 or 2"}
	}

	//Read in User Text input
	selection, _ := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	data, _ := os.ReadFile(selection)
	text := string(data)

	// Write User input to current directory
	tmpFile, _ := os.CreateTemp("", "vale-*.txt")
	tmpFile.Write([]byte(text)) //puts text from input into the temporary file
	tmpFile.Close()

	// Interacting with Vale CLI
	cmdArgs := []string{tmpFile.Name()}
	prefint, _ := strconv.Atoi(pref)
	switch prefint {
	case 1:
		cmdArgs = append(cmdArgs, "--config=.vale.ini") //singular
	case 2:
		cmdArgs = append(cmdArgs, "--config=_vale.ini") //plural
	}
	cmd := exec.Command("vale", cmdArgs...) //runs Vale
	outvale, err := cmd.CombinedOutput()    //returns output from Vale CLI
	if err != nil {                         //Error for issues with running Vale
		return []string{"Something is wrong with your input file. Check that it is a .txt or .md.", "Alternately, Vale could not have loaded correctly"}
	}

	//Processing Output from Vale to frontend
	rawout := string(outvale)
	processedout := strings.Split(rawout, "\n")
	processedout = processedout[2 : len(processedout)-3] //removes some unnecesary lines but could be improved further
	return processedout
}

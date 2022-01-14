package main

import (
	"os"

	"github.com/nszilard/slack/cmd"

	"github.com/nszilard/log"
)

//----------------------------------------
// Variables to mock in tests
//----------------------------------------
var (
	exit               = os.Exit
	executeRootCommand = cmd.Execute
)

//----------------------------------------
// Main
//----------------------------------------
func main() {
	if err := executeRootCommand(); err != nil {
		log.Errorf("slack: %v", err)
		exit(1)
	}
}

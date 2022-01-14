package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/nszilard/slack/cmd"
)

func captureOutput(t *testing.T, f func(), name string) string {
	t.Helper()

	// Create and remove capture file
	captured, err := os.Create(name + ".log")
	if err != nil {
		t.Fatalf("error creating capture file: %v", err)
	}
	defer os.Remove(captured.Name())

	// Redirecting output
	stdout := os.Stdout
	os.Stdout = captured

	// Call function
	f()

	// Reset output
	os.Stdout = stdout

	// Read captured data
	data, err := os.ReadFile(captured.Name())
	if err != nil {
		t.Fatalf("failed to read capture file: %v", err)
	}

	// Return as string
	return strings.TrimSpace(string(data))
}

//-------------------------------------------
// Mocks
//-------------------------------------------
func fakeExit(int) {
	fmt.Println("os.Exit(1)")
}

func fakeExecute() error {
	return fmt.Errorf("some error")
}

//-------------------------------------------
// Tests
//-------------------------------------------
func TestMain(t *testing.T) {
	expected := "A simple CLI tool to send Slack messages programmatically using pre-defined templates."
	actual := captureOutput(t, main, "main_test.log")

	if !strings.HasPrefix(actual, expected) {
		t.Errorf("main: unexpected output for root command")
	}
}

func TestMainError(t *testing.T) {
	exit = fakeExit
	executeRootCommand = fakeExecute

	defer func() {
		exit = os.Exit
		executeRootCommand = cmd.Execute
	}()

	expected := "os.Exit(1)"
	actual := captureOutput(t, main, "main_test_error.log")
	if !strings.HasPrefix(actual, expected) {
		t.Errorf("main: unexpected error output for root command")
	}
}

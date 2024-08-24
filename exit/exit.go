// Package exit provides functionality for handling program termination
package exit

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// StatusAborted is the exit status code used when the program is aborted
const StatusAborted = 130

// Abort terminates the program with a status code indicating it was aborted
// It returns a tea.Model and tea.Cmd as required by the Bubble Tea framework,
// but these are not used since the program exits immediately
func Abort() (tea.Model, tea.Cmd) {
	// Exit the program with the StatusAborted code
	os.Exit(StatusAborted)

	// These return values are never reached due to os.Exit,
	// but are included to satisfy the function signature
	return nil, nil
}

package input

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Options struct defines the configuration for the input prompt
type Options struct {
	Title       string // The title or prompt text
	Placeholder string // Placeholder text for the input field
	Help        bool   // Flag to indicate if help should be displayed
}

// Run function initializes and runs the input prompt
func Run(o Options) (string, error) {
	var output string
	model := new(o)            // Create a new model with the given options
	model.output = &output     // Set the output pointer
	p := tea.NewProgram(model) // Create a new Bubble Tea program
	_, err := p.Run()          // Run the program
	return output, err         // Return the user input and any error
}

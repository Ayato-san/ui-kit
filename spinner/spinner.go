package spinner

import (
	tea "github.com/charmbracelet/bubbletea"
)

// program struct holds a pointer to a tea.Program
type program struct {
	program *tea.Program
}

// Options struct defines the configuration for the spinner
type Options struct {
	Text    string // Text to display while the spinner is active
	EndText string // Text to display when the spinner stops
}

// Init creates and returns a new program instance with the given options
func Init(o Options) program {
	var d program
	p := tea.NewProgram(new(o)) // Create a new tea.Program with the options
	d.program = p
	return d
}

// Start runs the spinner program and returns any error encountered
func (d *program) Start() error {
	_, err := d.program.Run()
	return err
}

// Stop sends an endMsg to the program, signaling it to stop
func (d *program) Stop() {
	d.program.Send(endMsg{})
}

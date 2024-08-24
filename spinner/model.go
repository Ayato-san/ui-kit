package spinner

import (
	"github.com/ayato-san/ui-kit/exit"
	"github.com/ayato-san/ui-kit/styles"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// endMsg is an empty struct used to signal the end of the spinner
type endMsg struct{}

// model represents the state of the spinner
type model struct {
	text     string        // Text to display alongside the spinner
	endText  string        // Text to display when the spinner ends
	spinner  spinner.Model // The actual spinner model
	quitting bool          // Flag to indicate if the spinner is quitting
}

// new creates a new spinner model with the given options
func new(o Options) model {
	// Create a new spinner instance
	s := spinner.New()
	// Set the spinner type to MiniDot
	s.Spinner = spinner.MiniDot
	// Apply the accent style to the spinner
	s.Style = styles.AccentStyle

	// Return a new model with the provided options and configured spinner
	return model{text: o.Text, endText: o.EndText, spinner: s}
}

// Init initializes the model and returns the initial command
func (m model) Init() tea.Cmd {
	// Return the spinner's Tick command to start the animation
	return m.spinner.Tick
}

// Update handles incoming messages and updates the model accordingly
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(): // If Ctrl+C is pressed, abort the program
			return exit.Abort()
		default: // For any other key press, do nothing
			return m, nil
		}
	case endMsg:
		// If an endMsg is received, set quitting to true and quit the program
		m.quitting = true
		return m, tea.Quit
	case spinner.TickMsg:
		// Update the spinner's state and get the next command
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	default:
		// For any other message type, do nothing
		return m, nil
	}
}

// View renders the current state of the model as a string
func (m model) View() string {
	var s string

	if m.quitting {
		// If quitting, display the end text
		s = m.endText
	} else {
		// Otherwise, display the spinner and the text
		s = m.spinner.View() + " " + m.text
	}

	return s
}

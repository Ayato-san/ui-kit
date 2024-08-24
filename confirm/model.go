package confirm

import (
	"github.com/ayato-san/ui-kit/exit"
	"github.com/ayato-san/ui-kit/styles"
	tea "github.com/charmbracelet/bubbletea"
)

// errMsg is a custom error type for error messages
type errMsg error

// model represents the state and configuration of the confirmation prompt
type model struct {
	title       string // The title of the confirmation prompt
	affirmative string // The text for the affirmative option (e.g., "Yes")
	negative    string // The text for the negative option (e.g., "No")
	help        bool   // Whether to display help text
	state       *bool  // Pointer to the current selection state (true for affirmative, false for negative)
	quitting    bool   // Whether the user is exiting the prompt
	err         error  // Any error that occurred during the prompt
}

// new creates a new model instance with the given options
func new(o Options) model {
	// Set default affirmative text, override if provided in options
	affirmative := "Yes"
	if o.Affirmative != "" {
		affirmative = o.Affirmative
	}

	// Set default negative text, override if provided in options
	negative := "No"
	if o.Negative != "" {
		negative = o.Negative
	}

	// Return a new model with the configured options
	return model{
		title:       o.Title,
		affirmative: affirmative,
		negative:    negative,
		help:        o.Help,
	}
}

// Init initializes the model
func (m model) Init() tea.Cmd {
	// No initialization needed, return nil command
	return nil
}

// Update handles incoming messages and updates the model accordingly
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyLeft.String(), "h": // Left arrow or 'h' key: set state to true (affirmative)
			*m.state = true
			return m, cmd
		case tea.KeyRight.String(), "l": // Right arrow or 'l' key: set state to false (negative)
			*m.state = false
			return m, cmd
		case "y": // 'y' key: set state to true and quit
			*m.state = true
			m.quitting = true
			return m, tea.Quit
		case "n": // 'n' key: set state to false and quit
			*m.state = false
			m.quitting = true
			return m, tea.Quit
		case tea.KeyEnter.String(): // Enter key: keep current state and quit
			m.quitting = true
			return m, tea.Quit
		case tea.KeyCtrlC.String(), tea.KeyEsc.String():
			// Ctrl+C or Esc: abort the program
			return exit.Abort()
		}

	case errMsg:
		// Handle error messages
		m.err = msg
		return m, nil
	}

	// Return the model and command if no specific action was taken
	return m, cmd
}

// View renders the current state of the model
func (m model) View() string {
	// If quitting, return an empty string (no view)
	if m.quitting {
		return ""
	}

	var view = ""
	// Add title to the view if it's set
	if m.title != "" {
		view = m.title + "\n\n"
	}

	// Render buttons based on the current state
	if *m.state {
		// Affirmative is active
		view += styles.ButtonActiveRender(m.affirmative)
		view += styles.ButtonRender(m.negative)
	} else {
		// Negative is active
		view += styles.ButtonRender(m.affirmative)
		view += styles.ButtonActiveRender(m.negative)
	}

	// Add help text if enabled
	if m.help {
		view += styles.HelpKeys([]styles.Keys{
			{Key: "←/→", Action: "toggle"},
			{Key: "enter", Action: "confirm"},
			{Key: "esc", Action: "abort"},
		}, true)
	}

	return view
}

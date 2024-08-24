package input

import (
	"github.com/ayato-san/ui-kit/exit"
	"github.com/ayato-san/ui-kit/styles"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type model struct {
	textInput textinput.Model // Handles the text input functionality
	output    *string         // Pointer to store the final input value
	title     string          // Title to display above the input field
	help      bool            // Flag to determine if help should be shown
	quitting  bool            // Flag to indicate if the program is exiting
	err       error           // Stores any error that occurs during operation
}

func new(o Options) model {
	ti := textinput.New()
	ti.Placeholder = o.Placeholder
	ti.Focus()
	ti.PromptStyle = styles.AccentStyle
	ti.Cursor.Style = styles.AccentStyle

	// Create and return a new model instance with initialized values
	return model{
		textInput: ti,
		err:       nil,
		title:     o.Title,
		help:      o.Help,
	}
}

func (m model) Init() tea.Cmd {
	// Initialize the model by returning the Blink command for the text input
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter: // When Enter is pressed, store the input value and quit
			*m.output = m.textInput.Value()
			m.quitting = true
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc: // When Ctrl+C or Esc is pressed, abort the program
			return exit.Abort()
		}

	case errMsg:
		// Handle any errors by storing them in the model
		m.err = msg
		return m, nil
	}

	// Update the text input component
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	// If quitting, return an empty string
	if m.quitting {
		return ""
	}

	var view = ""
	// Add the title to the view if it's not empty
	if m.title != "" {
		view = m.title + "\n\n"
	}

	// Add the text input view
	view += m.textInput.View()

	// Add help keys if help is enabled
	if m.help {
		view += styles.HelpKeys([]styles.Keys{{Key: "esc", Action: "abort"}}, true)
	}

	return view
}

package confirm

import tea "github.com/charmbracelet/bubbletea"

// Options represents the configuration for the confirmation prompt
type Options struct {
	Title       string // The title or question to display
	Affirmative string // The text for the affirmative option (e.g., "Yes")
	Negative    string // The text for the negative option (e.g., "No")
	Help        bool   // Whether to display help information
}

// Run displays the confirmation prompt and returns the user's choice
func Run(o Options) (bool, error) {
	var state bool = true
	model := new(o)
	model.state = &state
	p := tea.NewProgram(model)
	_, err := p.Run()
	return state, err
}

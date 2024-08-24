package choice

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

// Options represents the configuration for the choice selection
type Options struct {
	Title string // The title of the selection menu
	Items []Item // The list of items to choose from
	Help  bool   // Whether to display help information
}

// Item represents a single selectable item
type Item struct {
	Display string // The text to display for this item
	Return  string // The value to return if this item is selected
}

// Run executes the choice selection process
// It returns the selected item's Return value and any error encountered
func Run(o Options) (string, error) {
	var selected int = 0         // Index of the selected item in the original list
	var selectedFiltered int = 0 // Index of the selected item in the filtered list

	model := new(o)
	model.selected = &selected
	model.selectedFiltered = &selectedFiltered

	// Create and run a new Bubble Tea program
	p := tea.NewProgram(model)
	_, err := p.Run()

	// Check if no option was selected
	if selected == -1 {
		return "", errors.New("no option selected")
	}

	// Return the selected item's Return value
	return o.Items[selected].Return, err
}

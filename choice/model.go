package choice

import (
	"strings"

	"github.com/ayato-san/ui-kit/exit"
	"github.com/ayato-san/ui-kit/styles"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// filteredItem represents an item in the filtered list
type filteredItem struct {
	Item      // Embedded Item type, containing the actual item data
	index int // Index of the item in the original unfiltered list
}

type errMsg error

// model represents the main structure for the choice UI
type model struct {
	title  string          // The title of the choice UI
	filter textinput.Model // Text input model for filtering options

	options  []Item         // Original list of selectable items
	filtered []filteredItem // Filtered list of items based on user input

	selectedFiltered *int // Pointer to the index of the selected item in the filtered list
	selected         *int // Pointer to the index of the selected item in the original list

	help     bool  // Flag to determine if help text should be displayed
	quitting bool  // Flag to indicate if the UI is in the process of quitting
	err      error // Stores any error that occurs during operation
}

// new creates and initializes a new model
func new(o Options) model {
	// Initialize a new text input for filtering
	filter := textinput.New()
	filter.Focus()
	filter.Prompt = ""
	filter.Cursor.Style = styles.AccentStyle

	// Create a slice of filteredItems with the same length as the input items
	filtered := make([]filteredItem, len(o.Items))
	for i, item := range o.Items {
		filtered[i] = filteredItem{
			Item:  item,
			index: i,
		}
	}

	// Initialize selected indices
	selectedFiltered := 0
	selected := 0

	// Return a new model instance with initialized fields
	return model{
		title:            o.Title,
		help:             o.Help,
		options:          o.Items,
		filter:           filter,
		filtered:         filtered,
		selectedFiltered: &selectedFiltered,
		selected:         &selected,
	}
}

// Init initializes the model
func (m model) Init() tea.Cmd {
	return nil
}

// filterOptions filters the options based on the given query
func (m *model) filterOptions(query string) {
	// Clear the current filtered list
	m.filtered = []filteredItem{}

	// Iterate through all options
	for i, item := range m.options {
		// Check if the item's display text contains the query (case-insensitive)
		if strings.Contains(strings.ToLower(item.Display), strings.ToLower(query)) {
			// Add matching item to the filtered list
			m.filtered = append(m.filtered, filteredItem{
				Item:  item,
				index: i,
			})
		}
	}

	// Adjust the selected index to stay within bounds
	if *m.selectedFiltered <= 0 {
		*m.selectedFiltered = 0
	}
	if *m.selectedFiltered >= len(m.filtered)-1 {
		*m.selectedFiltered = len(m.filtered) - 1
	}

	// Update the selected option based on the new filtered list
	m.selectOption()
}

// selectOption updates the selected option based on the filtered list
func (m *model) selectOption() {
	if len(m.filtered) == 0 {
		// If no items are filtered, set selected to -1 (no selection)
		*m.selected = -1
	} else {
		// Set the selected index to the original index of the selected filtered item
		*m.selected = m.filtered[*m.selectedFiltered].index
	}
}

// Update handles the model updates based on incoming messages
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp: // Move selection up if not at the top
			if *m.selectedFiltered != 0 {
				*m.selectedFiltered--
				m.selectOption()
			}
			return m, cmd
		case tea.KeyDown: // Move selection down if not at the bottom
			if *m.selectedFiltered != len(m.filtered)-1 {
				*m.selectedFiltered++
				m.selectOption()
			}
			return m, cmd
		case tea.KeyEnter: // Confirm selection and quit
			m.quitting = true
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			// Abort and exit
			return exit.Abort()
		}

	case errMsg:
		// Handle error messages
		m.err = msg
		return m, nil
	}

	// Update the filter input and re-filter options
	m.filter, cmd = m.filter.Update(msg)
	m.filterOptions(m.filter.Value())
	return m, cmd
}

// View renders the current state of the model
func (m model) View() string {
	if m.quitting {
		return ""
	}

	var view = ""
	// Add title and filter input to the view
	if m.title != "" {
		view = m.title + " " + m.filter.View() + "\n"
	} else {
		view = m.filter.View() + "\n"
	}

	// Render filtered items
	for index, item := range m.filtered {
		view += "\n"
		// Show selection indicator
		if *m.selectedFiltered == index {
			view += styles.AccentStyle.Render(" ✓ ")
		} else {
			view += " • "
		}

		// Highlight the matched text in the item display
		lowercaseDisplay := strings.ToLower(item.Display)
		lowercaseFilter := strings.ToLower(m.filter.Value())
		if lowercaseFilter != "" {
			startIndex := strings.Index(lowercaseDisplay, lowercaseFilter)
			if startIndex != -1 {
				endIndex := startIndex + len(lowercaseFilter)
				view += item.Display[:startIndex] + styles.AccentStyle.Render(item.Display[startIndex:endIndex]) + item.Display[endIndex:]
			} else {
				view += item.Display
			}
		} else {
			view += item.Display
		}
	}

	// Add help text if enabled
	if m.help {
		view += styles.HelpKeys([]styles.Keys{
			{Key: "↑/↓", Action: "switch"},
			{Key: "enter", Action: "confirm"},
			{Key: "esc", Action: "abort"},
		}, true)
	}

	return view
}

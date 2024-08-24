package styles

import "github.com/charmbracelet/lipgloss"

const activeColor = lipgloss.Color("#2196F3")

var (
	helpRender         = lipgloss.NewStyle().Foreground(lipgloss.Color("238")).Render
	helpKeyRender      = lipgloss.NewStyle().Foreground(lipgloss.Color("251")).Render
	ButtonRender       = buttonStyle.Render
	ButtonActiveRender = buttonStyle.Background(activeColor).Render
)

var (
	AccentStyle = lipgloss.NewStyle().Foreground(activeColor)
	buttonStyle = lipgloss.NewStyle().Margin(0, 2).Padding(0, 1).Bold(true)
)

func helpKey(key string, action string) string {
	return helpKeyRender(key) + " " + helpRender(action)
}

type Keys struct {
	Key    string
	Action string
}

func HelpKeys(keys []Keys, lineBreak bool) string {
	var help string

	for _, key := range keys {
		if help != "" {
			help += helpRender("  •  ")
		}
		help += helpKey(key.Key, key.Action)
	}

	if lineBreak {
		help = "\n\n" + help
	}
	return help
}

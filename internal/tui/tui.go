package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/taylow/awaik-backend/services"
)

// SelectServices returns a slice of services to run
func SelectServices(registry services.ServiceRegistry, selected services.ServiceRegistry) (services.ServiceRegistry, error) {
	p := tea.NewProgram(NewModel(registry, selected))
	if m, err := p.Run(); err != nil {
		return nil, err
	} else {
		registry = registry.Filter(m.(model).servicesToRun, false)
	}

	return registry, nil
}

// model is the model for the TUI
type model struct {
	choices       []string         // available services
	cursor        int              // which service our cursor is pointing at
	selected      map[int]struct{} // which services are selected
	servicesToRun []string         // which services to run
}

// NewModel returns a new model
func NewModel(registry services.ServiceRegistry, selected services.ServiceRegistry) model {
	choices := append([]string{"All"}, registry.NamesWithEmojis()...)
	selectedChoices := make(map[int]struct{})
	for _, service := range selected {
		name := removeEmoji(service.Name())
		for i, choice := range choices {
			if strings.Contains(choice, name) {
				selectedChoices[i] = struct{}{}
			}
		}
	}

	if len(selectedChoices) == len(choices)-1 {
		selectedChoices[0] = struct{}{}
	}

	return model{
		choices:  choices,
		selected: selectedChoices,
	}
}

// Init initialises the model
func (m model) Init() tea.Cmd {
	return nil
}

// Update updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ":
			_, ok := m.selected[m.cursor]
			if m.cursor == 0 && !ok {
				// Select all
				for i := 0; i < len(m.choices); i++ {
					m.selected[i] = struct{}{}
				}
			} else if m.cursor == 0 && ok {
				// Deselect all
				for i := 0; i < len(m.choices); i++ {
					delete(m.selected, i)
				}
			} else {
				if ok {
					delete(m.selected, m.cursor)
					if m.cursor != 0 {
						delete(m.selected, 0)
					}
				} else {
					m.selected[m.cursor] = struct{}{}
					if len(m.selected) == len(m.choices)-1 {
						m.selected[0] = struct{}{}
					}
				}
			}
		case "enter":
			fmt.Println("enter")
			for i := range m.selected {
				if i != 0 {
					m.servicesToRun = append(m.servicesToRun, removeEmoji(m.choices[i]))
				}
			}
			return m, tea.Quit
		}
	}

	return m, nil
}

// View returns the view for the model
func (m model) View() string {
	s := drawLogo() + "\n\n"
	s += "Which services do you want to run?\n\n"

	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// footer
	s += "\nPress q to quit.\n"
	s += "\nPress space to run configuration.\n"

	return s
}

// drawLogo draws the Awaik logo in ascii
func drawLogo() string {
	return `
    /\               (_) |   
   /  \__      ____ _ _| | __
  / /\ \ \ /\ / / _  | | |/ /
 / ____ \ V  V / (_| | |   < 
/_/    \_\_/\_/ \__,_|_|_|\_\`
}

// removeEmoji removes the emoji from a service name
func removeEmoji(s string) string {
	i := strings.LastIndex(s, " ")
	if i == -1 {
		return s
	}
	return s[i+1:]
}

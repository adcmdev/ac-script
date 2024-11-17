package shared

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
)

type Color string

const (
	Red      Color = "red"
	Green    Color = "green"
	Yellow   Color = "yellow"
	Blue     Color = "blue"
	DarkBlue Color = "darkblue"
	Magenta  Color = "magenta"
	Cyan     Color = "cyan"
	White    Color = "white"
)

type MultipleChoiceItem struct {
	Text  string
	Color Color
}

type model struct {
	options          []MultipleChoiceItem
	selected         int
	inputText        string
	cursorPos        int
	cursorInitialPos int
}

func (m *model) Init() tea.Cmd {
	return nil
}
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if len(m.options) > 0 {
			return m.onChangeList(msg)
		} else {
			return m.onInput(msg)
		}
	}
	return m, nil
}

func (m *model) onChangeList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch strings.ToLower(msg.String()) {
	case "up":
		if m.selected > 0 {
			m.selected--
		}
	case "down":
		if m.selected < len(m.options)-1 {
			m.selected++
		}
	case "enter":
		return m, tea.Quit
	default:
	}

	return m, nil
}

func (m *model) onInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		return m, tea.Quit
	case "backspace":
		totalLength := len(msg.String()) + m.cursorPos - 10
		initialPos := m.cursorInitialPos

		if totalLength >= initialPos {
			m.inputText = m.inputText[:m.cursorPos-1] + m.inputText[m.cursorPos:]
			m.cursorPos--
		}
	case "up", "down", "left", "right":
		// Do nothing for navigation keys

	default:
		if len(msg.String()) == 1 {
			message := msg.String()
			char, _ := utf8.DecodeRuneInString(message)

			if unicode.IsLetter(char) || unicode.IsDigit(char) || char == '-' || char == '_' {
				m.inputText = m.inputText[:m.cursorPos] + string(char) + m.inputText[m.cursorPos:]
				m.cursorPos++
			}
		}
	}

	return m, nil
}

func (m *model) View() string {
	var b strings.Builder

	if len(m.options) > 0 {
		b.WriteString("Choose an option:\n")
		for i, option := range m.options {
			if i == m.selected {
				b.WriteString(fmt.Sprintf("> %s\n", option.Text))
			} else {
				b.WriteString(fmt.Sprintf("  %s\n", option.Text))
			}
		}
	}

	if len(m.options) == 0 {
		b.WriteString(m.inputText)
		b.WriteString("|")
	}

	return b.String()
}

func ShowMultipleChoiceSelector(title string, options []MultipleChoiceItem) (result string, err error) {
	m := &model{
		options:          options,
		inputText:        title,
		selected:         0,
		cursorPos:        len(title),
		cursorInitialPos: len(title),
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return "", err
	}

	return options[m.selected].Text, nil
}

func AskForInput(question string) (result string, err error) {
	m := &model{
		options:          nil,
		inputText:        question,
		selected:         0,
		cursorPos:        len(question),
		cursorInitialPos: len(question),
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return "", err
	}

	return strings.Split(strings.TrimSpace(m.inputText), ":")[1], nil
}

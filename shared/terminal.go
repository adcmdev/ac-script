package shared

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type Color string

func (c Color) toTermboxColor() termbox.Attribute {
	var fgColor termbox.Attribute
	switch c {
	case Red:
		fgColor = termbox.ColorRed
	case Green:
		fgColor = termbox.ColorGreen
	case Yellow:
		fgColor = termbox.ColorYellow
	case Blue:
		fgColor = termbox.ColorLightBlue
	case DarkBlue:
		fgColor = termbox.ColorBlue
	case Magenta:
		fgColor = termbox.ColorMagenta
	case Cyan:
		fgColor = termbox.ColorCyan
	case White:
		fgColor = termbox.ColorWhite
	default:
		fgColor = termbox.ColorDefault
	}

	return fgColor
}

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
	Text string
	Color
}

func ShowMultipleChoiceSelector(title string, options []MultipleChoiceItem) (result string, err error) {
	err = termbox.Init()
	if err != nil {
		return
	}
	defer termbox.Close()

	selected := 0

	for {
		err = termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		if err != nil {
			return
		}

		printTitle(title)
		printOptions(options, selected)
		err = termbox.Flush()
		if err != nil {
			return
		}

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyArrowUp {
				if selected > 0 {
					selected--
				}
			} else if ev.Key == termbox.KeyArrowDown {
				if selected < len(options)-1 {
					selected++
				}
			} else if ev.Key == termbox.KeyEnter {
				fmt.Printf("Has seleccionado: %s\n", options[selected])
				result = options[selected].Text

				return
			} else if ev.Key == termbox.KeyEsc {

				return
			}
		default:

		}
	}
}

func printTitle(title string) {
	for i, char := range title {
		termbox.SetCell(i+5, 1, char, termbox.ColorYellow, termbox.ColorDefault)
	}
}

func printOptions(options []MultipleChoiceItem, selected int) {
	for i, option := range options {
		fgColor := option.Color.toTermboxColor()
		if i == selected {
			fgColor = termbox.ColorYellow
		}

		for j, char := range option.Text {
			termbox.SetCell(j+5, i+3, char, fgColor, termbox.ColorDefault)
		}
	}
}

func AskForInput(question string) (result string, err error) {
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	result = getAnswer(question)

	fmt.Printf("Entered answer: %s\n", result)

	return
}

func getAnswer(question string) string {
	answer := ""
	posX := 5
	posY := 3
	cursorPos := len(answer)

	for {
		printQuestion(question, posX, posY)
		printAnswer(answer, posX, posY+2, cursorPos)
		err := termbox.Flush()
		if err != nil {
			return ""
		}

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEnter {
				return answer
			} else if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				if len(answer) > 0 {
					answer = answer[:len(answer)-1]
					cursorPos--
				}
			} else if ev.Key == termbox.KeyEsc {
				return ""
			} else if ev.Ch != 0 {
				answer += string(ev.Ch)
				cursorPos++
			}
		default:
		}
	}
}

func printQuestion(question string, x, y int) {
	for i, char := range question {
		termbox.SetCell(x+i, y, char, termbox.ColorWhite, termbox.ColorDefault)
	}
}

func printAnswer(answer string, x, y, cursorPos int) {
	for i, char := range answer {
		termbox.SetCell(x+i, y, char, termbox.ColorWhite, termbox.ColorDefault)
	}

	termbox.SetCell(x+cursorPos, y, '_', termbox.ColorWhite, termbox.ColorDefault)
}

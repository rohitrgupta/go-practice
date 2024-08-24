package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type digitValue int

type Digit struct {
	*tview.Box
	value []digitValue
}

const (
	Zero digitValue = iota
	One
)

var text = map[digitValue][][]rune{
	Zero: [][]rune{[]rune(`███`), []rune(`█ █`), []rune(`█ █`), []rune(`█ █`), []rune(`███`)},
	One:  [][]rune{[]rune(`██ `), []rune(` █ `), []rune(` █ `), []rune(` █ `), []rune(`███`)},
}

// NewDigit returns a new digit.
func NewDigit() *Digit {
	box := tview.NewBox().SetBorder(true)
	return &Digit{
		Box:   box,
		value: []digitValue{Zero, One},
	}
}

func (d *Digit) SetValue(value []digitValue) {
	d.value = value
}

// Draw draws this primitive onto the screen.
func (d *Digit) Draw(screen tcell.Screen) {
	d.Box.DrawForSubclass(screen, d)
	x, y, _, _ := d.Box.GetInnerRect()
	for digit := range d.value {
		for i, s := range text[d.value[digit]] {
			fmt.Println(i)
			screen.SetContent(x+digit*4, y+i, ' ', s, tcell.StyleDefault)
		}
	}
}

func main() {
	box := NewDigit()
	app := tview.NewApplication()
	app.SetRoot(box, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

package main

import (
	"time"

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
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Dot
	Colon
)

var text = map[digitValue][]string{
	Zero:  {"███", "█ █", "█ █", "█ █", "███"},
	One:   {"██ ", " █ ", " █ ", " █ ", "███"},
	Two:   {"███", "  █", "███", "█  ", "███"},
	Three: {"███", "  █", "███", "  █", "███"},
	Four:  {"█ █", "█ █", "███", "  █", "  █"},
	Five:  {"███", "█  ", "███", "  █", "███"},
	Six:   {"███", "█  ", "███", "█ █", "███"},
	Seven: {"███", "  █", "  █", "  █", "  █"},
	Eight: {"███", "█ █", "███", "█ █", "███"},
	Nine:  {"███", "█ █", "███", "  █", "███"},
	Dot:   {"   ", "   ", "   ", "   ", "  █"},
	Colon: {"   ", " █ ", "   ", " █ ", "   "},
}

// NewDigit returns a new digit.
func NewDigit() *Digit {
	box := tview.NewBox().SetBorder(true)
	return &Digit{
		Box:   box,
		value: []digitValue{Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Dot, Colon},
	}
}

func (d *Digit) SetValue(value []digitValue) {
	d.value = value
}

// Draw draws this primitive onto the screen.
func (d *Digit) Draw(screen tcell.Screen) {
	d.Box.DrawForSubclass(screen, d)
	x, y, w, _ := d.Box.GetInnerRect()

	for digit := range d.value {
		for i, s := range text[d.value[digit]] {
			if x+(digit+1)*4 <= w {
				tview.Print(screen, s, x+(digit)*4, i+y, w-x+(digit+1)*4-4, tview.AlignLeft, tcell.ColorDefault)
			}
		}
	}
}

type Clock struct {
	*Digit
}

func NewClock() *Clock {
	digit := NewDigit()
	digit.SetValue([]digitValue{Zero, Zero, Colon, Zero, Zero, Colon, Zero, Zero})
	return &Clock{
		Digit: digit,
	}
}

func (m *Clock) Update() {
	t := int(time.Now().Unix())
	sec := (t + 5.5*3600) % 86400
	min := sec / 60
	sec %= 60
	hour := min / 60
	min %= 60

	m.SetValue([]digitValue{
		digitValue(hour / 10),
		digitValue(hour % 10),
		Colon,
		digitValue(min / 10),
		digitValue(min % 10),
		Colon,
		digitValue(sec / 10),
		digitValue(sec % 10),
	})
}

var (
	app   *tview.Application
	clock *Clock
)

func refresh() {
	tick := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-tick.C:
			clock.Update()
			app.Draw()
		}
	}
}

func main() {
	clock = NewClock()
	flex := tview.NewFlex().
		AddItem(clock, 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Blink").SetBorderAttributes(tcell.AttrBlink), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("x"), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("x"), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("x"), 0, 1, false)

	// box := NewDigit()
	app = tview.NewApplication()
	app.SetRoot(flex, true)
	go refresh()
	if err := app.Run(); err != nil {
		panic(err)
	}
}

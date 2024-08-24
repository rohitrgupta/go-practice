package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	checkbox := tview.NewCheckbox().SetLabel("Hit Enter to check box: ")
	checkbox.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyTab {
			checkbox.SetChecked(!checkbox.IsChecked())
		}
	})

	if err := app.SetRoot(checkbox, true).SetFocus(checkbox).Run(); err != nil {
		panic(err)
	}
}

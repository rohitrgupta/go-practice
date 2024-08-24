package main

import "github.com/rivo/tview"

func main() {
	app := tview.NewApplication()
	dropdown := tview.NewDropDown().
		SetLabel("Select an option (hit Enter): ").
		SetOptions([]string{"First", "Second", "Third", "Fourth", "Fifth"}, nil)

	dropdown.SetCurrentOption(2)

	if err := app.SetRoot(dropdown, true).SetFocus(dropdown).Run(); err != nil {
		panic(err)
	}
}

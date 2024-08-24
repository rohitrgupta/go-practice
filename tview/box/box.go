package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func main() {
	box := tview.NewBox().
		SetBorder(true).
		SetTitle("Box Demo")

	box.SetFocusFunc(func() {
		fmt.Println("Box focused")
	})
	app := tview.NewApplication()
	app.SetRoot(box, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

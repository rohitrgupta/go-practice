package main

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/sysread/textsel"
)

const text = `[green]This is an [::b]example[::-] of the [red]tview-textsel[-] package.

Use [::b]arrow keys[::-] to move the cursor.
Press [blue]space[-] to start/stop selecting text.`

func main() {
	app := tview.NewApplication()

	textSel := textsel.
		NewTextSel().
		SetText(text)

	textSel.SetSelectFunc(func(selectedText string) {
		app.Stop()
		fmt.Println("Selected text:\n\n", selectedText)
	})

	if err := app.SetRoot(textSel, true).Run(); err != nil {
		panic(err)
	}
}

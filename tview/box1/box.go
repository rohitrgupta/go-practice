package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Bold").SetBorderAttributes(tcell.AttrBold), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Blink").SetBorderAttributes(tcell.AttrBlink), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Reverse").SetBorderAttributes(tcell.AttrReverse), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Underline").SetBorderAttributes(tcell.AttrUnderline), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Dim").SetBorderAttributes(tcell.AttrDim), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Itelic").SetBorderAttributes(tcell.AttrItalic), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Strike").SetBorderAttributes(tcell.AttrStrikeThrough), 0, 1, false)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}

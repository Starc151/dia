package fyne

import (
	bl "github.com/Starc151/dia/pkg/bolus"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Show() {
	a := app.New()
	w := a.NewWindow("BL")
	// w.Resize(fyne.NewSize(200, 0))
	// w.SetFixedSize(true)
	icon, _ := fyne.LoadResourceFromPath("pkg/fyne/icon.png")
	w.SetIcon(icon)

	glucoseText := widget.NewLabel("GL: ")
	glucose := onlyNums()
	xeText := widget.NewLabel("XE: ")
	xe := onlyNums()
	bolus := widget.NewLabel("")
	history := widget.NewLabel(bl.GetHistory())
	getBolusBtnVis := func ()  {}
	getBolusBtn := widget.NewButton("Рассчитать болюс",
		func () {
			bolus.SetText(
				bl.GetBolus(
					glucose.Text,
					xe.Text,
				),
			)
			getBolusBtnVis()
			history.SetText(bl.GetHistory())
		},
	)
	getBolusBtnVis = func ()  {
		getBolusBtn.Disable()
	}

	tabs := container.NewAppTabs(
		container.NewTabItem(
			"Bolus",
			container.NewGridWithRows(3,
				container.NewGridWithColumns(4,
					glucoseText, glucose,
					xeText, xe,
				),
				getBolusBtn,
				bolus,
			),
		),
		container.NewTabItem("History", history),
	)
	w.SetContent(
		tabs,
	)
	w.ShowAndRun()
}
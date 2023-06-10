package fyne

import (
	"github.com/Starc151/dia/bolusPack"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Show() {
	a := app.New()
	w := a.NewWindow("Рассчёт болюса")
	// w.Resize(fyne.NewSize(200, 200))
	icon, _ := fyne.LoadResourceFromPath("icon.png")
	w.SetIcon(icon)

	glucoseText := widget.NewLabel("GL: ")
	glucose := onlyNumsn()
	
	xeText := widget.NewLabel("XE: ")
	xe := onlyNumsn()
	bolus := widget.NewLabel("")

	getBolus := widget.NewButton("Рассчитать болюс",
		func () {
			bolus.SetText(
				bolusPack.GetBolus(
					glucose.Text,
					xe.Text,
				),
			)
		},
	)
	w.SetContent(
		container.NewGridWithRows(3,
			container.NewGridWithColumns(4,
				glucoseText, glucose,
				xeText, xe,
			),
			getBolus,
			bolus,
		),
	)
	w.ShowAndRun()
}
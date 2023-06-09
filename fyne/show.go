package fyne

import (
	"github.com/Starc151/dia/bolusPack"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Show() {
	a := app.New()
	w := a.NewWindow("Рассчёт болюса")

	glucoseText := widget.NewLabel("Уровень глюкозы: ")
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
		})
			
	w.SetContent(
		container.NewVBox(
			glucoseText, glucose,
			xeText, xe,
			getBolus,
			bolus,
		),
	)
	
	w.ShowAndRun()
}
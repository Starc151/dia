package fyne

import (
	"fmt"

	bl "github.com/Starc151/dia/pkg/bolus"
	"github.com/Starc151/dia/pkg/ydb"

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
	glucose := onlyNumsn()
	xeText := widget.NewLabel("XE: ")
	xe := onlyNumsn()
	bolus := widget.NewLabel("")
	resultYdb := ydb.Select()
	history := ""
	history += fmt.Sprint(resultYdb[0].Date.Format("02/Jan"), "\n")
	for i, day := 0, 1; true; i++ {
		if resultYdb[i].Date.Format("02/Jan") == resultYdb[i+1].Date.Format("02/Jan"){
			history += fmt.Sprint(resultYdb[i].Date.Format("15:04"), resultYdb[i].Glucose, resultYdb[i].Xe, resultYdb[i].Bolus, "\n")
		} else{
			history += fmt.Sprint(resultYdb[i].Date.Format("15:04"), resultYdb[i].Glucose, resultYdb[i].Xe, resultYdb[i].Bolus,  "\n")
			day ++
			if day > 2{
				break
			}
			history += fmt.Sprint(resultYdb[i+1].Date.Format("02/Jan"), "\n")
		}
	}

	getBolus := widget.NewButton("Рассчитать болюс",
		func () {
			bolus.SetText(
				bl.GetBolus(
					glucose.Text,
					xe.Text,
				),
			)
		},
	)

	tabs := container.NewAppTabs(
		container.NewTabItem(
			"Bolus",
			container.NewGridWithRows(3,
				container.NewGridWithColumns(4,
					glucoseText, glucose,
					xeText, xe,
				),
				getBolus,
				bolus,
			),
		),
		container.NewTabItem("History", widget.NewLabel(history)),
	)
	w.SetContent(
		tabs,
	)
	w.ShowAndRun()
}
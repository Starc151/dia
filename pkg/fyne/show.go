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
	// w.Resize(fyne.NewSize(200, 230))
	// w.SetFixedSize(true)
	icon, _ := fyne.LoadResourceFromPath("pkg/fyne/icon.png")
	w.SetIcon(icon)

	glucoseText := widget.NewLabel("GL: ")
	glucose := onlyNums()
	xeText := widget.NewLabel("XE: ")
	xe := onlyNums()
	bolus := widget.NewLabel("")

	fullHistory := bl.GetFullHistory()
	t := func () *fyne.Container {
		fullHistory = bl.GetFullHistory()
		history := container.NewVBox()
			if len(fullHistory) == 0 {
				history.Add(widget.NewLabel("NO RESULT"))
			} else {
				for k, v := range fullHistory {
					if k > 1 {
						break
					}
					day := bl.GetDayHistory(v)
					dayW := getHistoryWidget(day)
					history.Add(dayW)
				}
			}
		return history
	}
	history := t()
	bolusBtnVisParam := true
	bolusBtnVis := func() {}
	getBolusBtn := widget.NewButton("Calculate bolus",
		func() {
			bolus.SetText(
				bl.GetBolus(
					glucose.Text,
					xe.Text,
				),
			)
			bolusBtnVis()
			history.RemoveAll()
			history.Add(t())
		},
	)
	updBtn := widget.NewButton("upd",
		func() {
			bolusBtnVis()
		},
	)
	updBtn.Disable()
	bolusBtnVis = func() {
		bolusBtnVisParam = !bolusBtnVisParam
		if !bolusBtnVisParam {
			getBolusBtn.Disable()
			updBtn.Enable()
		} else {
			getBolusBtn.Enable()
			updBtn.Disable()
		}
	}

	tabs := container.NewAppTabs(
		container.NewTabItem(
			"Bolus",
			container.NewGridWithRows(3,
				container.NewGridWithColumns(4,
					glucoseText, glucose,
					xeText, xe,
				),
				container.NewGridWithColumns(2, getBolusBtn, updBtn),
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

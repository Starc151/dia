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
	w.Resize(fyne.NewSize(205, 325))
	w.SetFixedSize(true)
	icon, _ := fyne.LoadResourceFromPath("pkg/fyne/icon.png")
	w.SetIcon(icon)
	iconUpdBtn, _ := fyne.LoadResourceFromPath("pkg/fyne/iconUpdBtn.png")

	glucoseText := widget.NewLabel("GL: ")
	glucose := onlyNums()
	xeText := widget.NewLabel("XE: ")
	xe := onlyNums()
	bolus := widget.NewLabel("")

	fullHistory := bl.GetFullHistory()
	getHistory := func() *fyne.Container {
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
	history := getHistory()
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
			history.Add(getHistory())
		},
	)
	updBtn := widget.NewButton("",
		func() {
			bolusBtnVis()
		},
	)
	getBolusBtn.Resize(fyne.NewSize(150, 30))
	getBolusBtn.Move(fyne.NewPos(5, 5))
	updBtn.Disable()
	updBtn.Resize(fyne.NewSize(30, 30))
	updBtn.Move(fyne.NewPos(160, 5))
	bolusBtnVis = func() {
		bolusBtnVisParam = !bolusBtnVisParam
		if !bolusBtnVisParam {
			getBolusBtn.Disable()
			updBtn.Enable()
			updBtn.SetIcon(iconUpdBtn)
		} else {
			getBolusBtn.Enable()
			updBtn.Disable()
			updBtn.SetIcon(nil)
		}
	}
	btns := container.NewWithoutLayout(getBolusBtn, updBtn)
	tabs := container.NewAppTabs(
		container.NewTabItem(
			"Bolus",
			container.NewGridWithRows(3,
				container.NewGridWithColumns(4,
					glucoseText, glucose,
					xeText, xe,
				),
				btns,
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

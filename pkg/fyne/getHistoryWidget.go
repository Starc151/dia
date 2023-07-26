package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getHistoryWidget(historyDay []string) *fyne.Container {
	cont := container.NewVBox()
	visPart := widget.NewLabel(historyDay[0])
	vis := false
	fVis := func(){}
	btn := widget.NewButton("Show", func() {
		vis = !vis
		fVis()
	})

	if len(historyDay) == 1 {
		cont.Add(visPart)
	} else {
		scrV := container.NewVScroll(visPart)
		scrV.SetMinSize(fyne.NewSize(0, 90))
		cont.Add(scrV)
		cont.Add(btn)
	}
	fVis = func() {
		if !vis {
			visPart.SetText(historyDay[0])
			btn.SetText("Show")
		} else {
			visPart.SetText(historyDay[1])
			btn.SetText("Hide")
		}
	}
	return cont
}
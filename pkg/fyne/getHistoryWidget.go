package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"	
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
	fVis = func() {
		if !vis {
			visPart.SetText(historyDay[0])
			btn.SetText("Show")
		} else {
			visPart.SetText(historyDay[1])
			btn.SetText("Hide")
		}
	}

	if len(historyDay) == 1 {
		cont.Add(visPart)
	} else {
		srcv := container.NewVScroll(
			container.NewVBox(visPart, btn),
		)
		srcv.SetMinSize(fyne.NewSize(0, 120))
		cont.Add(srcv)
	}
	return cont
}
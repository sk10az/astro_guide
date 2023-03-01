package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// TODO:
func calcScreen(_ fyne.Window) fyne.CanvasObject {

	return container.NewBorder(widget.NewLabel("calc"), nil, nil, nil)
}

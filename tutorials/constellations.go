package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// TODO:
func constellationsScreen(_ fyne.Window) fyne.CanvasObject {

	return container.NewBorder(widget.NewLabel("constellations"), nil, nil, nil)
}

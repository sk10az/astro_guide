package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ContentData struct {
	Title, Intro string
	Img          string
}

// iconScreen loads a panel that shows the various icons available in Fyne
func sliderScreen(_ fyne.Window) fyne.CanvasObject {
	img := canvas.NewImageFromFile("resources/img/planets/neptune.png")
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(300, 300))

	return container.NewCenter(container.NewVBox(
		container.NewHBox(
			widget.NewLabel("btn_left"),
			widget.NewLabel("btn_right"),
			widget.NewLabel("btn_dialog"),
		),
		container.NewHBox(
			//widget.NewLabel("img"),
			img,
			widget.NewLabel("info"),
		),
	),
		widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	)
}

func loadContentImg(_ fyne.Window) fyne.CanvasObject {

}

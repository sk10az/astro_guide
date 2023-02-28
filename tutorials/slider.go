package tutorials

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"reflect"
)

// картинка в контенте
type ContentData struct {
	Title, Intro string
	Img          *canvas.Image
}

// Картинка
//type ContentImgInfo struct {
//	icon *canvas.Image
//}

var (
	ContentSlider = map[string]ContentData{
		"earth": {
			Title: "Earth",
			Intro: "Earth is our planet",
			Img:   canvas.NewImageFromFile("resources/img/planets/earth.png"),
		},
		"moon": {},
	}
)

// TODO:
func sliderScreen(_ fyne.Window) fyne.CanvasObject {
	img := canvas.NewImageFromFile("resources/img/planets/neptune.png")
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(300, 300))

	fmt.Println(reflect.TypeOf(img))
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

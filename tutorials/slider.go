package tutorials

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"reflect"
)

// картинка в контенте
type ContentData struct {
	Title, Intro string
	Img          *canvas.Image
}

type ContentBrowser struct {
	current int
	icons   []ContentData

	name *widget.Select
	icon *canvas.Image
}

// Картинка
//type ContentImgInfo struct {
//	icon *canvas.Image
//}

var (
	ContentSlider = map[string]ContentData{
		"mercury": {
			Title: "Mercury",
			Intro: "Mercury is 1st planet's from sun",
			Img:   canvas.NewImageFromFile("resources/img/planets/mercury.png"),
		},
		"earth": {
			Title: "Earth",
			Intro: "Earth is our planet",
			Img:   canvas.NewImageFromFile("resources/img/planets/earth.png"),
		},
		"moon": {
			Title: "Moon",
			Intro: "Moon shart",
			Img:   canvas.NewImageFromFile("resources/img/planets/moon.png"),
		},
	}
)

func (b *ContentBrowser) setIcon(index int) {
	if index < 0 || index > len(b.icons)-1 {
		return
	}
	b.current = index

	b.name.SetSelected(b.icons[index].Title)
	//b.icon.(b.icons[index].Img)

}

// TODO:
func sliderScreen(_ fyne.Window) fyne.CanvasObject {
	b := &ContentBrowser{}
	//b.name.SetSelected(b.icons[b.current].Title)
	img := canvas.NewImageFromFile("resources/img/planets/neptune.png")
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(300, 300))

	img1, _ := fyne.LoadResourceFromPath("resources/img/planets/jupiter.png")
	fmt.Println("IMGGG: %s", reflect.TypeOf(img1))
	prev := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		b.setIcon(b.current - 1)
		fmt.Println("prev")
	})
	next := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
		b.setIcon(b.current + 1)
		fmt.Println("next")
	})

	buttons := container.NewHBox(prev, next)

	fmt.Println(reflect.TypeOf(img))
	return container.NewCenter(container.NewVBox(
		container.NewHBox(
			buttons,
			widget.NewLabel("btn_dialog"),
		),
		container.NewHBox(
			//widget.NewLabel("img"),
			img,
			//b.icon,
			widget.NewLabel("info"),
		),
	),
		widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	)
}

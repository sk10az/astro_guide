package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sk10az/astro_guide/resources"
	"net/url"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func welcomeScreen(_ fyne.Window) fyne.CanvasObject {

	var res = fyne.NewStaticResource("earth_logo.png", resources.EarthLogo)
	logo := canvas.NewImageFromResource(res)
	logo.FillMode = canvas.ImageFillContain

	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(192, 192))
	} else {
		logo.SetMinSize(fyne.NewSize(350, 350))
	}

	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Электронный справочник звёздного неба", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logo,
		container.NewHBox(
			widget.NewHyperlink("fyne.io", parseURL("https://fyne.io/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("documentation", parseURL("https://developer.fyne.io/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("author", parseURL("https://github.com/sk10az/")),
		),
		widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	))
}

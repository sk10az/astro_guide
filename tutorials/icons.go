package tutorials

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"reflect"
)

type iconInfo struct {
	name        string
	icon        fyne.Resource
	description string
}

type browser struct {
	current int
	icons   []iconInfo
	desks   []iconInfo

	desk *widget.Label
	name *widget.Select
	icon *widget.Icon
}

func (b *browser) setIcon(index int) {
	if index < 0 || index > len(b.icons)-1 {
		return
	}
	b.current = index

	//b.desk.SetText(b.icons[index].description)
	b.name.SetSelected(b.icons[index].name)
	b.icon.SetResource(b.icons[index].icon)
	//b.desk.SetText(b.desks[index].description)
}

func (b *browser) setContentText(index int) {

}

// TODO:
// iconScreen loads a panel that shows the various icons available in Fyne
func iconScreen(_ fyne.Window) fyne.CanvasObject {
	b := &browser{}
	b.icons = loadIcons()

	prev := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		b.setIcon(b.current - 1)
	})
	next := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
		b.setIcon(b.current + 1)
	})
	b.name = widget.NewSelect(iconList(b.icons), func(name string) {
		for i, icon := range b.icons {
			if icon.name == name {
				if b.current != i {
					b.setIcon(i)
				}
				break
			}
		}
	})
	b.name.SetSelected(b.icons[b.current].name)
	//b.desk.SetText(b.icons[b.current].description)
	buttons := container.NewHBox(prev, next)
	bar := container.NewBorder(nil, nil, buttons, nil, b.name)

	fmt.Println(reflect.TypeOf(bar))

	background := canvas.NewRasterWithPixels(checkerPattern)
	background.SetMinSize(fyne.NewSize(280, 280))
	b.icon = widget.NewIcon(b.icons[b.current].icon)

	return container.NewBorder(bar, nil, nil, widget.NewLabel("PIZDECPIZDECPIZDECPIZDEC"), background, b.icon)
}

func checkerPattern(x, y, _, _ int) color.Color {
	x /= 20
	y /= 20

	if x%2 == y%2 {
		return theme.BackgroundColor()
	}

	return theme.ShadowColor()
}

func iconList(icons []iconInfo) []string {
	ret := make([]string, len(icons))
	for i, icon := range icons {
		ret[i] = icon.name
	}

	return ret
}

func loadIcons() []iconInfo {
	mercury, _ := fyne.LoadResourceFromPath("resources/img/planets/mercury.png")
	venus, _ := fyne.LoadResourceFromPath("resources/img/planets/venus.png")
	earth, _ := fyne.LoadResourceFromPath("resources/img/planets/earth.png")
	mars, _ := fyne.LoadResourceFromPath("resources/img/planets/mars.png")
	jupiter, _ := fyne.LoadResourceFromPath("resources/img/planets/jupiter.png")
	saturn, _ := fyne.LoadResourceFromPath("resources/img/planets/jupiter.png")
	uranium, _ := fyne.LoadResourceFromPath("resources/img/planets/jupiter.png")
	neptune, _ := fyne.LoadResourceFromPath("resources/img/planets/jupiter.png")
	return []iconInfo{
		{"Mercury", mercury, "1"},
		{"Venus", venus, "2"},
		{"Earth", earth, "3"},
		{"Mars", mars, "4"},
		{"Jupiter", jupiter, "5"},
		{"Saturn", saturn, "6"},
		{"Uranium", uranium, "7"},
		{"Neptune", neptune, "8"},
	}
}

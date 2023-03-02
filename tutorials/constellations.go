package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type constellationsIconInfo struct {
	name string
	icon fyne.Resource
}

type constellationsBrowser struct {
	current int
	icons   []constellationsIconInfo

	name *widget.Select
	icon *widget.Icon
}

func (b *constellationsBrowser) setConstellationsIcon(index int) {
	if index < 0 || index > len(b.icons)-1 {
		return
	}
	b.current = index
	b.name.SetSelected(b.icons[index].name)
	b.icon.SetResource(b.icons[index].icon)
}

// TODO:
func constellationsScreen(_ fyne.Window) fyne.CanvasObject {
	b := &constellationsBrowser{}
	b.icons = loadConstellationsIcons()

	prev := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		b.setConstellationsIcon(b.current - 1)
	})

	next := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
		b.setConstellationsIcon(b.current + 1)
	})

	b.name = widget.NewSelect(iconConstellationsList(b.icons), func(name string) {
		for i, icon := range b.icons {
			if icon.name == name {
				if b.current != i {
					b.setConstellationsIcon(i)
				}
				break
			}
		}
	})

	b.name.SetSelected(b.icons[b.current].name)
	buttons := container.NewHBox(prev, next)
	bar := container.NewBorder(nil, nil, buttons, nil, b.name)
	background := canvas.NewRasterWithPixels(checkerPattern)
	background.Resize(fyne.NewSize(400, 400))
	b.icon = widget.NewIcon(b.icons[b.current].icon)

	return container.NewBorder(bar, nil, nil, nil, b.icon) // background перед b.icon

}

func iconConstellationsList(icons []constellationsIconInfo) []string {
	ret := make([]string, len(icons))
	for i, icon := range icons {
		ret[i] = icon.name
	}
	return ret
}

func loadConstellationsIcons() []constellationsIconInfo {
	bigbear, _ := fyne.LoadResourceFromPath("resources/img/constellations/bigbear.png")
	venus, _ := fyne.LoadResourceFromPath("resources/img/constellations/cefey.png")
	earth, _ := fyne.LoadResourceFromPath("resources/img/constellations/centaurus.png")
	mars, _ := fyne.LoadResourceFromPath("resources/img/constellations/deva.png")
	jupiter, _ := fyne.LoadResourceFromPath("resources/img/constellations/hercules.png")
	saturn, _ := fyne.LoadResourceFromPath("resources/img/constellations/hydra.png")
	uranium, _ := fyne.LoadResourceFromPath("resources/img/constellations/orion.png")
	neptune, _ := fyne.LoadResourceFromPath("resources/img/planets/neptune.png")
	return []constellationsIconInfo{
		{"Большая медведица", bigbear},
		{"Venus", venus},
		{"Earth", earth},
		{"Mars", mars},
		{"Jupiter", jupiter},
		{"Saturn", saturn},
		{"Uranium", uranium},
		{"Neptune", neptune},
	}
}

func checkerConstellationsPattern(x, y, _, _ int) color.Color {
	x /= 20
	y /= 20

	if x%2 == y%2 {
		return theme.BackgroundColor()
	}

	return theme.ShadowColor()
}

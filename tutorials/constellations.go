package tutorials

import (
	"fyne.io/fyne/v2"
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
	icons   []iconInfo

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

	return container.NewBorder(widget.NewLabel("constellations"), nil, nil, nil)
}

func iconConstellationsList(icons []constellationsIconInfo) []string {
	ret := make([]string, len(icons))
	for i, icon := range icons {
		ret[i] = icon.name
	}
	return ret
}

func loadConstellationsIcons() []constellationsIconInfo {
	mercury, _ := fyne.LoadResourceFromPath("resources/img/planets/mercury.png")
	venus, _ := fyne.LoadResourceFromPath("resources/img/planets/venus.png")
	earth, _ := fyne.LoadResourceFromPath("resources/img/planets/earth.png")
	mars, _ := fyne.LoadResourceFromPath("resources/img/planets/mars.png")
	jupiter, _ := fyne.LoadResourceFromPath("resources/img/planets/jupiter.png")
	saturn, _ := fyne.LoadResourceFromPath("resources/img/planets/saturn.png")
	uranium, _ := fyne.LoadResourceFromPath("resources/img/planets/uranium.png")
	neptune, _ := fyne.LoadResourceFromPath("resources/img/planets/neptune.png")
	return []constellationsIconInfo{
		{"Mercury", mercury},
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

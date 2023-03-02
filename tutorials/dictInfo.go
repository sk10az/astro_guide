package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
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

	desk []string
	name *widget.Select
	icon *widget.Icon
}

func (b *browser) setIcon(index int) {
	if index < 0 || index > len(b.icons)-1 {
		return
	}
	b.current = index
	b.name.SetSelected(b.icons[index].name)
	b.icon.SetResource(b.icons[index].icon)
}

func (b *browser) setLabel(index int) string {
	if index < 0 || index > len(b.desks)-1 {
		return ""
	}
	b.current = index
	description := b.desks[index].description
	//fmt.Println("description: ", description)
	return description
}

// TODO: какая же хрень
func iconScreen(_ fyne.Window) fyne.CanvasObject {
	b := &browser{}
	b.icons = loadIcons()
	b.desks = loadIcons()

	prev := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		b.setIcon(b.current - 1)
		b.setLabel(b.current)
	})

	next := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
		b.setIcon(b.current + 1)
		b.setLabel(b.current)
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
	buttons := container.NewHBox(prev, next)
	bar := container.NewBorder(nil, nil, buttons, nil, b.name)

	background := canvas.NewRasterWithPixels(checkerPattern)
	background.Resize(fyne.NewSize(400, 400))
	b.icon = widget.NewIcon(b.icons[b.current].icon)

	return container.NewBorder(bar, nil, nil, nil, b.icon) // background перед b.icon // container.NewGridWithColumns(2, widget.NewLabel(b.desks[b.current].description),
}

func iconList(icons []iconInfo) []string {
	ret := make([]string, len(icons))
	for i, icon := range icons {
		ret[i] = icon.name

	}
	return ret
}

func textList(desks []iconInfo) []string {
	retDesk := make([]string, len(desks))
	for i, desk := range desks {
		retDesk[i] = desk.description
	}
	return retDesk
}

func loadIcons() []iconInfo {
	mercury, _ := fyne.LoadResourceFromPath("resources/img/planets/mercury.png")
	venus, _ := fyne.LoadResourceFromPath("resources/img/planets/venus.png")
	earth, _ := fyne.LoadResourceFromPath("resources/img/planets/earth.png")
	mars, _ := fyne.LoadResourceFromPath("resources/img/planets/mars.png")
	jupiter, _ := fyne.LoadResourceFromPath("resources/img/planets/jupiter.png")
	saturn, _ := fyne.LoadResourceFromPath("resources/img/planets/saturn.png")
	uranium, _ := fyne.LoadResourceFromPath("resources/img/planets/uranium.png")
	neptune, _ := fyne.LoadResourceFromPath("resources/img/planets/neptune.png")
	return []iconInfo{
		{"Mercury", mercury, "Наименьшая планета Солнечной системы и самая близкая к Солнцу.\n " +
			"Названа в честь древнеримского бога торговли - быстрого Меркурия,\n " +
			"поскольку она движется по небу быстрее других планет.\n " +
			"Её период обращения вокруг Солнца составляет всего 87,97 земных\n " +
			"суток - самый короткий среди всех планет Солнечной системы.\n"},
		{"Venus", venus, "venus"},
		{"Earth", earth, "earth"},
		{"Mars", mars, "mars"},
		{"Jupiter", jupiter, "jupiter"},
		{"Saturn", saturn, "saturn"},
		{"Uranium", uranium, "uranium"},
		{"Neptune", neptune, "neptune"},
	}
}

func checkerPattern(x, y, _, _ int) color.Color {
	x /= 20
	y /= 20

	if x%2 == y%2 {
		return theme.BackgroundColor()
	}

	return theme.ShadowColor()
}

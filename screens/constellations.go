package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type constellationsIconInfo struct {
	name        string
	icon        fyne.Resource
	description string // добавлено поле для описания
}

type constellationsBrowser struct {
	current     int
	icons       []constellationsIconInfo
	name        *widget.Select
	icon        *widget.Icon
	description *widget.Label // добавили поле для описания
}

func (b *constellationsBrowser) setConstellationsIcon(index int) {
	if index < 0 || index > len(b.icons)-1 {
		return
	}
	b.current = index
	b.name.SetSelected(b.icons[index].name)
	b.icon.SetResource(b.icons[index].icon)
	b.setDescription(b.icons[index].description) // обновляем описание
}

func (b *constellationsBrowser) setDescription(description string) {
	b.description.SetText(description)
	b.description.Refresh()
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

	b.description = widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}) // создаем Label для описания
	b.icon = widget.NewIcon(b.icons[b.current].icon)

	content := container.NewGridWithColumns(2,
		b.description, // добавляем Label вместо названия
		b.icon,
	)

	b.setDescription(b.icons[b.current].description) // устанавливаем начальное описание

	return container.NewBorder(bar, nil, nil, nil, content)
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
	return []constellationsIconInfo{
		{name: "Большая медведица", icon: bigbear, description: "Большая медведица - созвездие северного полушария зимнего неба."},
		{name: "Венера", icon: venus, description: "Венера - вторая планета от Солнца."},
		{name: "Центавр", icon: earth, description: "Центавр - созвездие южного полушария неба."},
		{name: "Марс", icon: mars, description: "Марс - четвёртая планета от Солнца."},
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

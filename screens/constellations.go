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
	андромеда, _ := fyne.LoadResourceFromPath("resources/img/constellations/1.jpg")
	большаяМедведица, _ := fyne.LoadResourceFromPath("resources/img/constellations/2.jpg")
	большойПёс, _ := fyne.LoadResourceFromPath("resources/img/constellations/3.jpg")
	близнецы, _ := fyne.LoadResourceFromPath("resources/img/constellations/4.jpg")
	чаша, _ := fyne.LoadResourceFromPath("resources/img/constellations/5.jpg")
	часы, _ := fyne.LoadResourceFromPath("resources/img/constellations/6.jpg")
	дельфин, _ := fyne.LoadResourceFromPath("resources/img/constellations/7.jpg")
	волосыВероники, _ := fyne.LoadResourceFromPath("resources/img/constellations/8.jpg")
	дракон, _ := fyne.LoadResourceFromPath("resources/img/constellations/9.jpg")
	единорог, _ := fyne.LoadResourceFromPath("resources/img/constellations/10.jpg")
	эридан, _ := fyne.LoadResourceFromPath("resources/img/constellations/11.jpg")
	геркулес, _ := fyne.LoadResourceFromPath("resources/img/constellations/12.jpg")
	жертвенник, _ := fyne.LoadResourceFromPath("resources/img/constellations/13.jpg")
	жираф, _ := fyne.LoadResourceFromPath("resources/img/constellations/14.jpg")
	return []constellationsIconInfo{
		{name: "Андромеда", icon: андромеда, description: "Андромеда"},
		{name: "Большая медведица", icon: большаяМедведица, description: "Большая медведица"},
		{name: "Большой пёс", icon: большойПёс, description: "Большой пёс"},
		{name: "Близнецы", icon: близнецы, description: "Близнецы"},
		{name: "Чаша", icon: чаша, description: "Близнецы"},
		{name: "Часы", icon: часы, description: "Часы"},
		{name: "Дельфин", icon: дельфин, description: "Дельфин"},
		{name: "Волосы Вероники", icon: волосыВероники, description: "Волосы Вероники"},
		{name: "Дракон", icon: дракон, description: "Дракон"},
		{name: "Единорог", icon: единорог, description: "Единорог"},
		{name: "Эридан", icon: эридан, description: "Эридан"},
		{name: "Геркулес", icon: геркулес, description: "Геркулес"},
		{name: "Жертвенник", icon: жертвенник, description: "Жертвенник"},
		{name: "Жираф", icon: жираф, description: "Жираф"},
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

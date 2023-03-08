package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sk10az/astro_guide/resources"
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

	desk   []string
	name   *widget.Select
	icon   *widget.Icon
	search *widget.Entry // добавлено поле search
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

func (b *browser) searchIcon(name string) int {
	for i, icon := range allIcons {
		if icon.name == name {
			return i
		}
	}
	return -1
}

var allIcons []iconInfo

func loadIcons() []iconInfo {
	var icons []iconInfo
	icons = append(icons, iconInfo{name: "меркурий", icon: fyne.NewStaticResource("mercury.png", resources.MercuryPlanets)})
	icons = append(icons, iconInfo{name: "венера", icon: fyne.NewStaticResource("venus.png", resources.VenusPlanets)})
	icons = append(icons, iconInfo{name: "земля", icon: fyne.NewStaticResource("earth.png", resources.EarthPlanets)})
	icons = append(icons, iconInfo{name: "марс", icon: fyne.NewStaticResource("mars.png", resources.MarsPlanets)})
	icons = append(icons, iconInfo{name: "юпитер", icon: fyne.NewStaticResource("jupiter.png", resources.JupiterPlanets)})
	icons = append(icons, iconInfo{name: "сатурн", icon: fyne.NewStaticResource("saturn.png", resources.SaturnPlanets)})
	icons = append(icons, iconInfo{name: "уран", icon: fyne.NewStaticResource("uranium.png", resources.UraniumPlanets)})
	icons = append(icons, iconInfo{name: "нептун", icon: fyne.NewStaticResource("neptune.png", resources.NeptunePlanets)})

	return icons
}

func iconScreen(_ fyne.Window) fyne.CanvasObject {
	b := &browser{}
	b.icons = loadIcons()
	b.desks = loadIcons()

	for _, icon := range b.icons {
		allIcons = append(allIcons, icon)
	}

	b.search = widget.NewEntry()
	b.search.SetPlaceHolder("Search by name")
	b.search.OnChanged = func(s string) {
		if i := b.searchIcon(s); i >= 0 {
			b.current = i
			b.setIcon(i)
			b.setLabel(i)
		}
	}

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
	searchBox := container.NewBorder(nil, nil, nil, nil, b.search)
	nameBox := container.NewBorder(searchBox, nil, nil, nil, b.name)
	buttons := container.NewHBox(prev, next)
	bar := container.NewBorder(nil, nil, buttons, nil, nameBox)

	background := canvas.NewRasterWithPixels(checkerPattern)
	background.Resize(fyne.NewSize(400, 400))
	b.icon = widget.NewIcon(b.icons[b.current].icon)

	return container.NewBorder(bar, nil, nil, nil, b.icon)
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

func checkerPattern(x, y, _, _ int) color.Color {
	x /= 20
	y /= 20

	if x%2 == y%2 {
		return theme.BackgroundColor()
	}

	return theme.ShadowColor()
}

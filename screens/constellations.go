package screens

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sk10az/astro_guide/resources"
	"image/color"
	"io/ioutil"
	"log"
)

type constellationsIconInfo struct {
	name        string
	icon        fyne.Resource
	description string
}

type constellationsBrowser struct {
	current     int
	icons       []constellationsIconInfo
	name        *widget.Select
	icon        *widget.Icon
	description *widget.Label
	search      *widget.Entry
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

func (b *constellationsBrowser) searchIcon(name string) int {
	for i, icon := range allConstellations {
		if icon.name == name {
			return i
		}
	}
	return -1
}

var allConstellations []constellationsIconInfo

func searchConstellationsByName(name string, icons []constellationsIconInfo) int {
	for i, icon := range icons {
		if icon.name == name {
			return i
		}
	}
	return -1
}

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
	search := widget.NewEntry()

	search.Resize(fyne.NewSize(search.MinSize().Width, search.MinSize().Height))
	search.SetPlaceHolder("Search by name")
	searchSubmit := widget.NewButtonWithIcon("", theme.SearchIcon(), func() {
		index := searchConstellationsByName(search.Text, b.icons)
		if index != -1 {
			b.name.SetSelected(b.icons[index].name)
			b.setConstellationsIcon(index)
			b.setDescription(b.icons[index].description)
		}
	})

	buttons := container.NewHBox(prev, next, search, searchSubmit)
	bar := container.NewBorder(nil, nil, buttons, nil, container.NewScroll(b.name))

	b.description = widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: false, Monospace: false}) // создаем Label для описания
	b.icon = widget.NewIcon(b.icons[b.current].icon)

	b.setDescription(b.icons[b.current].description) // устанавливаем начальное описание
	b.description.Wrapping = fyne.TextWrapWord       // устанавливаем перенос по словам
	b.description.Resize(fyne.NewSize(300, 500))     // устанавливаем максимальную ширину
	content := container.NewGridWithColumns(2,
		b.description, // добавляем Label вместо названия
		b.icon,
	)
	content.MinSize()

	return container.NewBorder(bar, nil, nil, nil, content)
}

func iconConstellationsList(icons []constellationsIconInfo) []string {
	ret := make([]string, len(icons))
	for i, icon := range icons {
		ret[i] = icon.name
	}
	return ret
}

func loadDescriptions() (map[string]string, error) {
	data, err := ioutil.ReadFile("screens/desc.json")
	if err != nil {
		return nil, err
	}
	var descriptions map[string]string
	err = json.Unmarshal(data, &descriptions)
	if err != nil {
		return nil, err
	}
	return descriptions, nil
}

func loadConstellationsIcons() []constellationsIconInfo {
	desc, err := loadDescriptions()
	if err != nil {
		log.Fatal(err)
	}
	var icons []constellationsIconInfo
	icons = append(icons, constellationsIconInfo{name: "созвездия", icon: fyne.NewStaticResource("Empty.png", resources.ConstellationsEmpty), description: ""})
	icons = append(icons, constellationsIconInfo{name: "андромеда", icon: fyne.NewStaticResource("1.jpg", resources.Constellations1), description: desc["андромеда"]})
	icons = append(icons, constellationsIconInfo{name: "большая медведица", icon: fyne.NewStaticResource("2.jpg", resources.Constellations2), description: desc["большая медведица"]})
	icons = append(icons, constellationsIconInfo{name: "большой пёс", icon: fyne.NewStaticResource("3.jpg", resources.Constellations3), description: desc["большой пёс"]})
	icons = append(icons, constellationsIconInfo{name: "близнецы", icon: fyne.NewStaticResource("4.jpg", resources.Constellations4), description: desc["близнецы"]})
	icons = append(icons, constellationsIconInfo{name: "чаша", icon: fyne.NewStaticResource("5.jpg", resources.Constellations5), description: desc["чаша"]})
	icons = append(icons, constellationsIconInfo{name: "часы", icon: fyne.NewStaticResource("6.jpg", resources.Constellations6), description: desc["часы"]})
	icons = append(icons, constellationsIconInfo{name: "дельфин", icon: fyne.NewStaticResource("7.jpg", resources.Constellations7), description: desc["дельфин"]})
	icons = append(icons, constellationsIconInfo{name: "дева", icon: fyne.NewStaticResource("8.jpg", resources.Constellations8), description: desc["дева"]})
	icons = append(icons, constellationsIconInfo{name: "дракон", icon: fyne.NewStaticResource("9.jpg", resources.Constellations9), description: desc["дракон"]})
	icons = append(icons, constellationsIconInfo{name: "единорог", icon: fyne.NewStaticResource("10.jpg", resources.Constellations10), description: desc["единорог"]})
	icons = append(icons, constellationsIconInfo{name: "эридан", icon: fyne.NewStaticResource("11.jpg", resources.Constellations11), description: desc["эридан"]})
	icons = append(icons, constellationsIconInfo{name: "геркулес", icon: fyne.NewStaticResource("12.jpg", resources.Constellations12), description: desc["геркулес"]})
	icons = append(icons, constellationsIconInfo{name: "жертвенник", icon: fyne.NewStaticResource("13.jpg", resources.Constellations13), description: desc["жертвенник"]})
	icons = append(icons, constellationsIconInfo{name: "жираф", icon: fyne.NewStaticResource("14.jpg", resources.Constellations14), description: desc["жираф"]})
	icons = append(icons, constellationsIconInfo{name: "живописец", icon: fyne.NewStaticResource("15.jpg", resources.Constellations15), description: desc["живописец"]})
	icons = append(icons, constellationsIconInfo{name: "журавль", icon: fyne.NewStaticResource("16.jpg", resources.Constellations16), description: desc["журавль"]})
	icons = append(icons, constellationsIconInfo{name: "гидра", icon: fyne.NewStaticResource("17.jpg", resources.Constellations17), description: desc["гидра"]})
	icons = append(icons, constellationsIconInfo{name: "голубь", icon: fyne.NewStaticResource("18.jpg", resources.Constellations18), description: desc["голубь"]})
	icons = append(icons, constellationsIconInfo{name: "гончие псы", icon: fyne.NewStaticResource("19.jpg", resources.Constellations19), description: desc["гончие псы"]})
	icons = append(icons, constellationsIconInfo{name: "хамелеон", icon: fyne.NewStaticResource("20.jpg", resources.Constellations20), description: desc["хамелеон"]})
	icons = append(icons, constellationsIconInfo{name: "индеец", icon: fyne.NewStaticResource("21.jpg", resources.Constellations21), description: desc["индеец"]})
	icons = append(icons, constellationsIconInfo{name: "кассиопея", icon: fyne.NewStaticResource("22.jpg", resources.Constellations22), description: desc["кассиопея"]})
	icons = append(icons, constellationsIconInfo{name: "киль", icon: fyne.NewStaticResource("23.jpg", resources.Constellations23), description: desc["киль"]})
	icons = append(icons, constellationsIconInfo{name: "кит", icon: fyne.NewStaticResource("24.jpg", resources.Constellations24), description: desc["кит"]})
	icons = append(icons, constellationsIconInfo{name: "компас", icon: fyne.NewStaticResource("25.jpg", resources.Constellations25), description: desc["компас"]})
	icons = append(icons, constellationsIconInfo{name: "корма", icon: fyne.NewStaticResource("26.jpg", resources.Constellations26), description: desc["корма"]})
	icons = append(icons, constellationsIconInfo{name: "козерог", icon: fyne.NewStaticResource("27.jpg", resources.Constellations27), description: desc["козерог"]})
	icons = append(icons, constellationsIconInfo{name: "лебедь", icon: fyne.NewStaticResource("28.jpg", resources.Constellations28), description: desc["лебедь"]})
	icons = append(icons, constellationsIconInfo{name: "летучая рыба", icon: fyne.NewStaticResource("29.jpg", resources.Constellations29), description: desc["летучая рыба"]})
	icons = append(icons, constellationsIconInfo{name: "лев", icon: fyne.NewStaticResource("30.jpg", resources.Constellations30), description: desc["лев"]})
	icons = append(icons, constellationsIconInfo{name: "лира", icon: fyne.NewStaticResource("31.jpg", resources.Constellations31), description: desc["лира"]})
	icons = append(icons, constellationsIconInfo{name: "лисичка", icon: fyne.NewStaticResource("32.jpg", resources.Constellations32), description: desc["лисичка"]})
	icons = append(icons, constellationsIconInfo{name: "малая медведица", icon: fyne.NewStaticResource("33.jpg", resources.Constellations33), description: desc["малая медведица"]})
	icons = append(icons, constellationsIconInfo{name: "малый конь", icon: fyne.NewStaticResource("34.jpg", resources.Constellations34), description: desc["малый конь"]})
	icons = append(icons, constellationsIconInfo{name: "малый лев", icon: fyne.NewStaticResource("35.jpg", resources.Constellations35), description: desc["малый лев"]})
	icons = append(icons, constellationsIconInfo{name: "малый пёс", icon: fyne.NewStaticResource("36.jpg", resources.Constellations36), description: desc["малый пёс"]})
	icons = append(icons, constellationsIconInfo{name: "миксроскоп", icon: fyne.NewStaticResource("37.jpg", resources.Constellations37), description: desc["миксроскоп"]})
	icons = append(icons, constellationsIconInfo{name: "муха", icon: fyne.NewStaticResource("38.jpg", resources.Constellations38), description: desc["муха"]})
	icons = append(icons, constellationsIconInfo{name: "насос", icon: fyne.NewStaticResource("39.jpg", resources.Constellations39), description: desc["насос"]})
	icons = append(icons, constellationsIconInfo{name: "наугольник", icon: fyne.NewStaticResource("40.jpg", resources.Constellations40), description: desc["наугольник"]})
	icons = append(icons, constellationsIconInfo{name: "октант", icon: fyne.NewStaticResource("41.jpg", resources.Constellations41), description: desc["октант"]})
	icons = append(icons, constellationsIconInfo{name: "орёл", icon: fyne.NewStaticResource("42.jpg", resources.Constellations42), description: desc["орёл"]})
	icons = append(icons, constellationsIconInfo{name: "орион", icon: fyne.NewStaticResource("43.jpg", resources.Constellations43), description: desc["орион"]})
	icons = append(icons, constellationsIconInfo{name: "овен", icon: fyne.NewStaticResource("44.jpg", resources.Constellations44), description: desc["овен"]})
	icons = append(icons, constellationsIconInfo{name: "паруса", icon: fyne.NewStaticResource("45.jpg", resources.Constellations45), description: desc["паруса"]})
	icons = append(icons, constellationsIconInfo{name: "павлин", icon: fyne.NewStaticResource("46.jpg", resources.Constellations46), description: desc["павлин"]})
	icons = append(icons, constellationsIconInfo{name: "печь", icon: fyne.NewStaticResource("47.jpg", resources.Constellations47), description: desc["печь"]})
	icons = append(icons, constellationsIconInfo{name: "пегас", icon: fyne.NewStaticResource("48.jpg", resources.Constellations48), description: desc["пегас"]})
	icons = append(icons, constellationsIconInfo{name: "персей", icon: fyne.NewStaticResource("49.jpg", resources.Constellations49), description: desc["персей"]})
	icons = append(icons, constellationsIconInfo{name: "феникс", icon: fyne.NewStaticResource("50.jpg", resources.Constellations50), description: desc["феникс"]})
	icons = append(icons, constellationsIconInfo{name: "рак", icon: fyne.NewStaticResource("51.jpg", resources.Constellations51), description: desc["рак"]})
	icons = append(icons, constellationsIconInfo{name: "райская птица", icon: fyne.NewStaticResource("52.jpg", resources.Constellations52), description: desc[""]})
	icons = append(icons, constellationsIconInfo{name: "резец", icon: fyne.NewStaticResource("53.jpg", resources.Constellations53), description: desc["райская птица"]})
	icons = append(icons, constellationsIconInfo{name: "рыбы", icon: fyne.NewStaticResource("54.jpg", resources.Constellations54), description: desc["резец"]})
	icons = append(icons, constellationsIconInfo{name: "рысь", icon: fyne.NewStaticResource("55.jpg", resources.Constellations55), description: desc["рыбы"]})
	icons = append(icons, constellationsIconInfo{name: "секстант", icon: fyne.NewStaticResource("56.jpg", resources.Constellations56), description: desc["секстант"]})
	icons = append(icons, constellationsIconInfo{name: "сетка", icon: fyne.NewStaticResource("57.jpg", resources.Constellations57), description: desc["сетка"]})
	icons = append(icons, constellationsIconInfo{name: "северная корона", icon: fyne.NewStaticResource("58.jpg", resources.Constellations58), description: desc["северная корона"]})
	icons = append(icons, constellationsIconInfo{name: "щит", icon: fyne.NewStaticResource("59.jpg", resources.Constellations59), description: desc["щит"]})
	icons = append(icons, constellationsIconInfo{name: "скорпион", icon: fyne.NewStaticResource("60.jpg", resources.Constellations60), description: desc["скорпион"]})
	icons = append(icons, constellationsIconInfo{name: "скульптор", icon: fyne.NewStaticResource("61.jpg", resources.Constellations61), description: desc["скульптор"]})
	icons = append(icons, constellationsIconInfo{name: "столовая гора", icon: fyne.NewStaticResource("62.jpg", resources.Constellations62), description: desc["столовая гора"]})
	icons = append(icons, constellationsIconInfo{name: "стрела", icon: fyne.NewStaticResource("63.jpg", resources.Constellations63), description: desc["стрела"]})
	icons = append(icons, constellationsIconInfo{name: "стрелец", icon: fyne.NewStaticResource("64.jpg", resources.Constellations64), description: desc["стрелец"]})
	icons = append(icons, constellationsIconInfo{name: "цефей", icon: fyne.NewStaticResource("65.jpg", resources.Constellations65), description: desc["цефей"]})
	icons = append(icons, constellationsIconInfo{name: "центавр", icon: fyne.NewStaticResource("66.jpg", resources.Constellations66), description: desc["центавр"]})
	icons = append(icons, constellationsIconInfo{name: "циркуль", icon: fyne.NewStaticResource("67.jpg", resources.Constellations67), description: desc["циркуль"]})
	icons = append(icons, constellationsIconInfo{name: "телескоп", icon: fyne.NewStaticResource("68.jpg", resources.Constellations68), description: desc["телескоп"]})
	icons = append(icons, constellationsIconInfo{name: "телец", icon: fyne.NewStaticResource("69.jpg", resources.Constellations69), description: desc["телец"]})
	icons = append(icons, constellationsIconInfo{name: "треугольник", icon: fyne.NewStaticResource("70.jpg", resources.Constellations70), description: desc["треугольник"]})
	icons = append(icons, constellationsIconInfo{name: "тука", icon: fyne.NewStaticResource("71.jpg", resources.Constellations71), description: desc["тука"]})
	icons = append(icons, constellationsIconInfo{name: "южная гидра", icon: fyne.NewStaticResource("72.jpg", resources.Constellations72), description: desc["южная гидра"]})
	icons = append(icons, constellationsIconInfo{name: "южная корона", icon: fyne.NewStaticResource("73.jpg", resources.Constellations73), description: desc["южная корона"]})
	icons = append(icons, constellationsIconInfo{name: "южная рыба", icon: fyne.NewStaticResource("74.jpg", resources.Constellations74), description: desc["южная рыба"]})
	icons = append(icons, constellationsIconInfo{name: "южный треугольник", icon: fyne.NewStaticResource("75.jpg", resources.Constellations75), description: desc["южный треугольник"]})
	icons = append(icons, constellationsIconInfo{name: "южный крест", icon: fyne.NewStaticResource("75.jpg", resources.Constellations76), description: desc["южный крест"]})
	icons = append(icons, constellationsIconInfo{name: "весы", icon: fyne.NewStaticResource("77.jpg", resources.Constellations77), description: desc["весы"]})
	icons = append(icons, constellationsIconInfo{name: "водолей", icon: fyne.NewStaticResource("78.jpg", resources.Constellations78), description: desc["водолей"]})
	icons = append(icons, constellationsIconInfo{name: "волк", icon: fyne.NewStaticResource("79.jpg", resources.Constellations79), description: desc["волк"]})
	icons = append(icons, constellationsIconInfo{name: "волопас", icon: fyne.NewStaticResource("80.jpg", resources.Constellations80), description: desc["волопас"]})
	icons = append(icons, constellationsIconInfo{name: "волосы Вероники", icon: fyne.NewStaticResource("81.jpg", resources.Constellations81), description: desc["волосы Вероники"]})
	icons = append(icons, constellationsIconInfo{name: "ворон", icon: fyne.NewStaticResource("82.jpg", resources.Constellations82), description: desc["ворон"]})
	icons = append(icons, constellationsIconInfo{name: "возничий", icon: fyne.NewStaticResource("83.jpg", resources.Constellations83), description: desc["возничий"]})
	icons = append(icons, constellationsIconInfo{name: "ящерица", icon: fyne.NewStaticResource("84.jpg", resources.Constellations84), description: desc["ящерица"]})
	icons = append(icons, constellationsIconInfo{name: "заяц", icon: fyne.NewStaticResource("85.jpg", resources.Constellations85), description: desc["заяц"]})
	icons = append(icons, constellationsIconInfo{name: "змееносец", icon: fyne.NewStaticResource("86.jpg", resources.Constellations86), description: desc["змееносец"]})
	icons = append(icons, constellationsIconInfo{name: "змея", icon: fyne.NewStaticResource("87.jpg", resources.Constellations87), description: desc["змея"]})
	icons = append(icons, constellationsIconInfo{name: "золотая рыбка", icon: fyne.NewStaticResource("88.jpg", resources.Constellations88), description: desc["золотая рыбка"]})

	return icons
}

func checkerConstellationsPattern(x, y, _, _ int) color.Color {
	x /= 20
	y /= 20

	if x%2 == y%2 {
		return theme.BackgroundColor()
	}

	return theme.ShadowColor()
}

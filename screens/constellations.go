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
	bar := container.NewBorder(nil, nil, buttons, nil, container.NewScroll(b.name))

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
	дева, _ := fyne.LoadResourceFromPath("resources/img/constellations/8.jpg")
	дракон, _ := fyne.LoadResourceFromPath("resources/img/constellations/9.jpg")
	единорог, _ := fyne.LoadResourceFromPath("resources/img/constellations/10.jpg")
	эридан, _ := fyne.LoadResourceFromPath("resources/img/constellations/11.jpg")
	геркулес, _ := fyne.LoadResourceFromPath("resources/img/constellations/12.jpg")
	жертвенник, _ := fyne.LoadResourceFromPath("resources/img/constellations/13.jpg")
	жираф, _ := fyne.LoadResourceFromPath("resources/img/constellations/14.jpg")
	живописец, _ := fyne.LoadResourceFromPath("resources/img/constellations/15.jpg")
	журавль, _ := fyne.LoadResourceFromPath("resources/img/constellations/16.jpg")
	гидра, _ := fyne.LoadResourceFromPath("resources/img/constellations/17.jpg")
	голубь, _ := fyne.LoadResourceFromPath("resources/img/constellations/18.jpg")
	гончиеПсы, _ := fyne.LoadResourceFromPath("resources/img/constellations/19.jpg")
	хамелеон, _ := fyne.LoadResourceFromPath("resources/img/constellations/20.jpg")
	индеец, _ := fyne.LoadResourceFromPath("resources/img/constellations/21.jpg")
	кассиопея, _ := fyne.LoadResourceFromPath("resources/img/constellations/22.jpg")
	киль, _ := fyne.LoadResourceFromPath("resources/img/constellations/23.jpg")
	кит, _ := fyne.LoadResourceFromPath("resources/img/constellations/24.jpg")
	компас, _ := fyne.LoadResourceFromPath("resources/img/constellations/25.jpg")
	корма, _ := fyne.LoadResourceFromPath("resources/img/constellations/26.jpg")
	козерог, _ := fyne.LoadResourceFromPath("resources/img/constellations/27.jpg")
	лебедь, _ := fyne.LoadResourceFromPath("resources/img/constellations/28.jpg")
	летучаяРыба, _ := fyne.LoadResourceFromPath("resources/img/constellations/29.jpg")
	лев, _ := fyne.LoadResourceFromPath("resources/img/constellations/30.jpg")
	лира, _ := fyne.LoadResourceFromPath("resources/img/constellations/31.jpg")
	лисичка, _ := fyne.LoadResourceFromPath("resources/img/constellations/32.jpg")
	малаяМедведица, _ := fyne.LoadResourceFromPath("resources/img/constellations/33.jpg")
	малыйКонь, _ := fyne.LoadResourceFromPath("resources/img/constellations/34.jpg")
	малыйЛев, _ := fyne.LoadResourceFromPath("resources/img/constellations/35.jpg")
	малыйПёс, _ := fyne.LoadResourceFromPath("resources/img/constellations/36.jpg")
	микроскоп, _ := fyne.LoadResourceFromPath("resources/img/constellations/37.jpg")
	муха, _ := fyne.LoadResourceFromPath("resources/img/constellations/38.jpg")
	насос, _ := fyne.LoadResourceFromPath("resources/img/constellations/39.jpg")
	наугольник, _ := fyne.LoadResourceFromPath("resources/img/constellations/40.jpg")
	октант, _ := fyne.LoadResourceFromPath("resources/img/constellations/41.jpg")
	орёл, _ := fyne.LoadResourceFromPath("resources/img/constellations/42.jpg")
	орион, _ := fyne.LoadResourceFromPath("resources/img/constellations/43.jpg")
	овен, _ := fyne.LoadResourceFromPath("resources/img/constellations/44.jpg")
	паруса, _ := fyne.LoadResourceFromPath("resources/img/constellations/45.jpg")
	павлин, _ := fyne.LoadResourceFromPath("resources/img/constellations/46.jpg")
	печь, _ := fyne.LoadResourceFromPath("resources/img/constellations/47.jpg")
	пегас, _ := fyne.LoadResourceFromPath("resources/img/constellations/48.jpg")
	персей, _ := fyne.LoadResourceFromPath("resources/img/constellations/49.jpg")
	феникс, _ := fyne.LoadResourceFromPath("resources/img/constellations/50.jpg")
	рак, _ := fyne.LoadResourceFromPath("resources/img/constellations/51.jpg")
	райскаяПтица, _ := fyne.LoadResourceFromPath("resources/img/constellations/52.jpg")
	резец, _ := fyne.LoadResourceFromPath("resources/img/constellations/53.jpg")
	рыбы, _ := fyne.LoadResourceFromPath("resources/img/constellations/54.jpg")
	рысь, _ := fyne.LoadResourceFromPath("resources/img/constellations/55.jpg")
	секстант, _ := fyne.LoadResourceFromPath("resources/img/constellations/56.jpg")
	сетка, _ := fyne.LoadResourceFromPath("resources/img/constellations/57.jpg")
	севернаяКорона, _ := fyne.LoadResourceFromPath("resources/img/constellations/58.jpg")
	щит, _ := fyne.LoadResourceFromPath("resources/img/constellations/59.jpg")
	скорпион, _ := fyne.LoadResourceFromPath("resources/img/constellations/60.jpg")
	скульптор, _ := fyne.LoadResourceFromPath("resources/img/constellations/61.jpg")
	столоваяГора, _ := fyne.LoadResourceFromPath("resources/img/constellations/62.jpg")
	стрела, _ := fyne.LoadResourceFromPath("resources/img/constellations/63.jpg")
	стрелец, _ := fyne.LoadResourceFromPath("resources/img/constellations/64.jpg")
	цефей, _ := fyne.LoadResourceFromPath("resources/img/constellations/65.jpg")
	центавр, _ := fyne.LoadResourceFromPath("resources/img/constellations/66.jpg")
	циркуль, _ := fyne.LoadResourceFromPath("resources/img/constellations/67.jpg")
	телескоп, _ := fyne.LoadResourceFromPath("resources/img/constellations/68.jpg")
	телец, _ := fyne.LoadResourceFromPath("resources/img/constellations/69.jpg")
	треугольник, _ := fyne.LoadResourceFromPath("resources/img/constellations/70.jpg")
	тукан, _ := fyne.LoadResourceFromPath("resources/img/constellations/71.jpg")
	южнаяГидра, _ := fyne.LoadResourceFromPath("resources/img/constellations/72.jpg")
	южнаяКорона, _ := fyne.LoadResourceFromPath("resources/img/constellations/73.jpg")
	южнаяРыба, _ := fyne.LoadResourceFromPath("resources/img/constellations/74.jpg")
	южныйТреугольник, _ := fyne.LoadResourceFromPath("resources/img/constellations/75.jpg")
	южныйКрест, _ := fyne.LoadResourceFromPath("resources/img/constellations/76.jpg")
	весы, _ := fyne.LoadResourceFromPath("resources/img/constellations/77.jpg")
	водолей, _ := fyne.LoadResourceFromPath("resources/img/constellations/78.jpg")
	волк, _ := fyne.LoadResourceFromPath("resources/img/constellations/79.jpg")
	волопас, _ := fyne.LoadResourceFromPath("resources/img/constellations/80.jpg")
	волосыВероники, _ := fyne.LoadResourceFromPath("resources/img/constellations/81.jpg")
	ворон, _ := fyne.LoadResourceFromPath("resources/img/constellations/82.jpg")
	возничий, _ := fyne.LoadResourceFromPath("resources/img/constellations/83.jpg")
	ящерица, _ := fyne.LoadResourceFromPath("resources/img/constellations/84.jpg")
	заяц, _ := fyne.LoadResourceFromPath("resources/img/constellations/85.jpg")
	змееносец, _ := fyne.LoadResourceFromPath("resources/img/constellations/86.jpg")
	змея, _ := fyne.LoadResourceFromPath("resources/img/constellations/87.jpg")
	золотаяРыбка, _ := fyne.LoadResourceFromPath("resources/img/constellations/88.jpg")
	return []constellationsIconInfo{
		{name: "Андромеда", icon: андромеда, description: "Андромеда"},
		{name: "Большая медведица", icon: большаяМедведица, description: "Большая медведица"},
		{name: "Большой пёс", icon: большойПёс, description: "Большой пёс"},
		{name: "Близнецы", icon: близнецы, description: "Близнецы"},
		{name: "Чаша", icon: чаша, description: "Близнецы"},
		{name: "Часы", icon: часы, description: "Часы"},
		{name: "Дельфин", icon: дельфин, description: "Дельфин"},
		{name: "Волосы Вероники", icon: дева, description: "Волосы Вероники"},
		{name: "Дракон", icon: дракон, description: "Дракон"},
		{name: "Единорог", icon: единорог, description: "Единорог"},
		{name: "Эридан", icon: эридан, description: "Эридан"},
		{name: "Геркулес", icon: геркулес, description: "Геркулес"},
		{name: "Жертвенник", icon: жертвенник, description: "Жертвенник"},
		{name: "Жираф", icon: жираф, description: "Жираф"},
		{name: "Живописец", icon: живописец, description: "Живописец"},
		{name: "Журавль", icon: журавль, description: "Журавль"},
		{name: "Гидра", icon: гидра, description: "Гидра"},
		{name: "Голубь", icon: голубь, description: "Голубь"},
		{name: "Гончие псы", icon: гончиеПсы, description: "Гончие псы"},
		{name: "Хамелеон", icon: хамелеон, description: "Хамелеон"},
		{name: "Индеец", icon: индеец, description: "Индеец"},
		{name: "Кассиопея", icon: кассиопея, description: "Кассиопея"},
		{name: "Киль", icon: киль, description: "Киль"},
		{name: "Кит", icon: кит, description: "Кит"},
		{name: "Компас", icon: компас, description: "Компас"},
		{name: "Корма", icon: корма, description: "Корма"},
		{name: "Козерог", icon: козерог, description: "Козерог"},
		{name: "Лебедь", icon: лебедь, description: "Лебедь"},
		{name: "Летучая рыба", icon: летучаяРыба, description: "Летучая рыба"},
		{name: "Лев", icon: лев, description: "Лев"},
		{name: "Лира", icon: лира, description: "Лира"},
		{name: "Лисичка", icon: лисичка, description: "Лисичка"},
		{name: "Малая медведица", icon: малаяМедведица, description: "Малая медведица"},
		{name: "Малый конь", icon: малыйКонь, description: "Малый конь"},
		{name: "Малый лев", icon: малыйЛев, description: "Малый лев"},
		{name: "Малый пёс", icon: малыйПёс, description: "Малый пёс"},
		{name: "Микроскоп", icon: микроскоп, description: "Микроскоп"},
		{name: "Муха", icon: муха, description: "Муха"},
		{name: "Насос", icon: насос, description: "Насос"},
		{name: "Наугольник", icon: наугольник, description: "Наугольник"},
		{name: "Октант", icon: октант, description: "Октант"},
		{name: "Орёл", icon: орёл, description: "Орёл"},
		{name: "Орион", icon: орион, description: "Орион"},
		{name: "Овен", icon: овен, description: "Овен"},
		{name: "Паруса", icon: паруса, description: "Паруса"},
		{name: "Павлин", icon: павлин, description: "Павлин"},
		{name: "Печь", icon: печь, description: "Печь"},
		{name: "Пегас", icon: пегас, description: "Пегас"},
		{name: "Персей", icon: персей, description: "Персей"},
		{name: "Феникс", icon: феникс, description: "Феникс"},
		{name: "Рак", icon: рак, description: "Рак"},
		{name: "Райская птица", icon: райскаяПтица, description: "Райская птица"},
		{name: "Резец", icon: резец, description: "Резец"},
		{name: "Рыбы", icon: рыбы, description: "Рыбы"},
		{name: "Рысь", icon: рысь, description: "Рысь"},
		{name: "Секстант", icon: секстант, description: "Секстант"},
		{name: "Сетка", icon: сетка, description: "Сетка"},
		{name: "Северная сетка", icon: севернаяКорона, description: "Северная сетка"},
		{name: "Щит", icon: щит, description: "Щит"},
		{name: "Скорпион", icon: скорпион, description: "Скорпион"},
		{name: "Скульптор", icon: скульптор, description: "Скульптор"},
		{name: "Столовая гора", icon: столоваяГора, description: "Столовая гора"},
		{name: "Стрела", icon: стрела, description: "Стрела"},
		{name: "Стрелец", icon: стрелец, description: "Стрелец"},
		{name: "Цефей", icon: цефей, description: "Цефей"},
		{name: "Центавр", icon: центавр, description: "Центавр"},
		{name: "Циркуль", icon: циркуль, description: "Циркуль"},
		{name: "Телескоп", icon: телескоп, description: "Телескоп"},
		{name: "Телец", icon: телец, description: "Телец"},
		{name: "Треугольник", icon: треугольник, description: "Треугольник"},
		{name: "Тукан", icon: тукан, description: "Тукан"},
		{name: "Южная гидра", icon: южнаяГидра, description: "Южная гидра"},
		{name: "Южная корона", icon: южнаяКорона, description: "Южная корона"},
		{name: "Южная рыба", icon: южнаяРыба, description: "Южная рыба"},
		{name: "Южный треугольник", icon: южныйТреугольник, description: "Южный треугольник"},
		{name: "Южный крест", icon: южныйКрест, description: "Южный крест"},
		{name: "Весы", icon: весы, description: "Весы"},
		{name: "Водолей", icon: водолей, description: "Водолей"},
		{name: "Волк", icon: волк, description: "Волк"},
		{name: "Волопас", icon: волопас, description: "Волопас"},
		{name: "Волосы Вероники", icon: волосыВероники, description: "Волосы Вероники"},
		{name: "Ворон", icon: ворон, description: "Ворон"},
		{name: "Возничий", icon: возничий, description: "Возничий"},
		{name: "Ящерица", icon: ящерица, description: "Ящерица"},
		{name: "Заяц", icon: заяц, description: "Заяц"},
		{name: "Змееносец", icon: змееносец, description: "Змееносец"},
		{name: "Змея", icon: змея, description: "Змея"},
		{name: "Золотая рыбка", icon: золотаяРыбка, description: "Золотая рыбка"},
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

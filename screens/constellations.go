package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sk10az/astro_guide/resources"
	"image/color"
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

func loadConstellationsIcons() []constellationsIconInfo {
	var icons []constellationsIconInfo
	//icons = append(icons, constellationsIconInfo{name: "empty", icon: fyne.LoadResourceFromPath("resources/img/Empty.png"), description: "df")
	icons = append(icons, constellationsIconInfo{name: "андромеда", icon: fyne.NewStaticResource("1.jpg", resources.Constellations1), description: "андромеда это из песни гостнейма"})
	icons = append(icons, constellationsIconInfo{name: "большая медведица", icon: fyne.NewStaticResource("2.jpg", resources.Constellations2), description: "Большой мишка - хороший мишка"})
	icons = append(icons, constellationsIconInfo{name: "большой пёс", icon: fyne.NewStaticResource("3.jpg", resources.Constellations3), description: "Большой мишка - хороший мишка"})
	icons = append(icons, constellationsIconInfo{name: "близнецы", icon: fyne.NewStaticResource("4.jpg", resources.Constellations4), description: "Близнецы - знак двойственности"})
	icons = append(icons, constellationsIconInfo{name: "чаша", icon: fyne.NewStaticResource("5.jpg", resources.Constellations5), description: "Чаша - знак воды"})
	icons = append(icons, constellationsIconInfo{name: "часы", icon: fyne.NewStaticResource("6.jpg", resources.Constellations6), description: "Часы - знак времени"})
	icons = append(icons, constellationsIconInfo{name: "дельфин", icon: fyne.NewStaticResource("7.jpg", resources.Constellations7), description: "Дельфин - символ свободы"})
	icons = append(icons, constellationsIconInfo{name: "дева", icon: fyne.NewStaticResource("8.jpg", resources.Constellations8), description: "Дева - знак чистоты"})
	icons = append(icons, constellationsIconInfo{name: "дракон", icon: fyne.NewStaticResource("9.jpg", resources.Constellations9), description: "Дракон - символ духовной силы"})
	icons = append(icons, constellationsIconInfo{name: "единорог", icon: fyne.NewStaticResource("10.jpg", resources.Constellations10), description: "Единорог - символ чистоты и невинности"})
	icons = append(icons, constellationsIconInfo{name: "эридан", icon: fyne.NewStaticResource("11.jpg", resources.Constellations11), description: "Эридан - знак реки"})
	icons = append(icons, constellationsIconInfo{name: "геркулес", icon: fyne.NewStaticResource("12.jpg", resources.Constellations12), description: "Геркулес - знак силы"})
	icons = append(icons, constellationsIconInfo{name: "жертвенник", icon: fyne.NewStaticResource("13.jpg", resources.Constellations13), description: "Жертвенник - символ жертвоприношения"})
	icons = append(icons, constellationsIconInfo{name: "жираф", icon: fyne.NewStaticResource("14.jpg", resources.Constellations14), description: "Жираф - знак долголетия"})
	icons = append(icons, constellationsIconInfo{name: "живописец", icon: fyne.NewStaticResource("15.jpg", resources.Constellations15), description: "Живописец - знак творчества"})
	icons = append(icons, constellationsIconInfo{name: "журавль", icon: fyne.NewStaticResource("16.jpg", resources.Constellations16), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "гидра", icon: fyne.NewStaticResource("17.jpg", resources.Constellations17), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "голубь", icon: fyne.NewStaticResource("18.jpg", resources.Constellations18), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "гончие псы", icon: fyne.NewStaticResource("19.jpg", resources.Constellations19), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "хамелеон", icon: fyne.NewStaticResource("20.jpg", resources.Constellations20), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "индеец", icon: fyne.NewStaticResource("21.jpg", resources.Constellations21), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "кассиопея", icon: fyne.NewStaticResource("22.jpg", resources.Constellations22), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "киль", icon: fyne.NewStaticResource("23.jpg", resources.Constellations23), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "кит", icon: fyne.NewStaticResource("24.jpg", resources.Constellations24), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "компас", icon: fyne.NewStaticResource("25.jpg", resources.Constellations25), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "корма", icon: fyne.NewStaticResource("26.jpg", resources.Constellations26), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "козерог", icon: fyne.NewStaticResource("27.jpg", resources.Constellations27), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "лебедь", icon: fyne.NewStaticResource("28.jpg", resources.Constellations28), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "летучая рыба", icon: fyne.NewStaticResource("29.jpg", resources.Constellations29), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "лев", icon: fyne.NewStaticResource("30.jpg", resources.Constellations30), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "лира", icon: fyne.NewStaticResource("31.jpg", resources.Constellations31), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "лисичка", icon: fyne.NewStaticResource("32.jpg", resources.Constellations32), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "малая медведица", icon: fyne.NewStaticResource("33.jpg", resources.Constellations33), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "малый конь", icon: fyne.NewStaticResource("34.jpg", resources.Constellations34), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "малый лев", icon: fyne.NewStaticResource("35.jpg", resources.Constellations35), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "малый пёс", icon: fyne.NewStaticResource("36.jpg", resources.Constellations36), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "миксроскоп", icon: fyne.NewStaticResource("37.jpg", resources.Constellations37), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "муха", icon: fyne.NewStaticResource("38.jpg", resources.Constellations38), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "насос", icon: fyne.NewStaticResource("39.jpg", resources.Constellations39), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "наугольник", icon: fyne.NewStaticResource("40.jpg", resources.Constellations40), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "октант", icon: fyne.NewStaticResource("41.jpg", resources.Constellations41), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "орёл", icon: fyne.NewStaticResource("42.jpg", resources.Constellations42), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "орион", icon: fyne.NewStaticResource("43.jpg", resources.Constellations43), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "овен", icon: fyne.NewStaticResource("44.jpg", resources.Constellations44), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "паруса", icon: fyne.NewStaticResource("45.jpg", resources.Constellations45), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "павлин", icon: fyne.NewStaticResource("46.jpg", resources.Constellations46), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "печь", icon: fyne.NewStaticResource("47.jpg", resources.Constellations47), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "пегас", icon: fyne.NewStaticResource("48.jpg", resources.Constellations48), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "персей", icon: fyne.NewStaticResource("49.jpg", resources.Constellations49), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "феникс", icon: fyne.NewStaticResource("50.jpg", resources.Constellations50), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "рак", icon: fyne.NewStaticResource("51.jpg", resources.Constellations51), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "райская птица", icon: fyne.NewStaticResource("52.jpg", resources.Constellations52), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "резец", icon: fyne.NewStaticResource("53.jpg", resources.Constellations53), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "рыбы", icon: fyne.NewStaticResource("54.jpg", resources.Constellations54), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "рысь", icon: fyne.NewStaticResource("55.jpg", resources.Constellations55), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "секстант", icon: fyne.NewStaticResource("56.jpg", resources.Constellations56), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "сетка", icon: fyne.NewStaticResource("57.jpg", resources.Constellations57), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "северная корона", icon: fyne.NewStaticResource("58.jpg", resources.Constellations58), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "щит", icon: fyne.NewStaticResource("59.jpg", resources.Constellations59), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "скорпион", icon: fyne.NewStaticResource("60.jpg", resources.Constellations60), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "скульптор", icon: fyne.NewStaticResource("61.jpg", resources.Constellations61), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "столовая гора", icon: fyne.NewStaticResource("62.jpg", resources.Constellations62), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "стрела", icon: fyne.NewStaticResource("63.jpg", resources.Constellations63), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "стрелец", icon: fyne.NewStaticResource("64.jpg", resources.Constellations64), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "цефей", icon: fyne.NewStaticResource("65.jpg", resources.Constellations65), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "центавр", icon: fyne.NewStaticResource("66.jpg", resources.Constellations66), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "циркуль", icon: fyne.NewStaticResource("67.jpg", resources.Constellations67), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "телескоп", icon: fyne.NewStaticResource("68.jpg", resources.Constellations68), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "телец", icon: fyne.NewStaticResource("69.jpg", resources.Constellations69), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "треугольник", icon: fyne.NewStaticResource("70.jpg", resources.Constellations70), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "тука", icon: fyne.NewStaticResource("71.jpg", resources.Constellations71), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "южная гидра", icon: fyne.NewStaticResource("72.jpg", resources.Constellations72), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "южная корона", icon: fyne.NewStaticResource("73.jpg", resources.Constellations73), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "южный треугольник", icon: fyne.NewStaticResource("74.jpg", resources.Constellations74), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "южный крест", icon: fyne.NewStaticResource("75.jpg", resources.Constellations75), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "весы", icon: fyne.NewStaticResource("76.jpg", resources.Constellations76), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "весы", icon: fyne.NewStaticResource("77.jpg", resources.Constellations77), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "водолей", icon: fyne.NewStaticResource("78.jpg", resources.Constellations78), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "волк", icon: fyne.NewStaticResource("79.jpg", resources.Constellations79), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "волопас", icon: fyne.NewStaticResource("80.jpg", resources.Constellations80), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "волосы Вероники", icon: fyne.NewStaticResource("81.jpg", resources.Constellations81), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "ворон", icon: fyne.NewStaticResource("82.jpg", resources.Constellations82), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "возничий", icon: fyne.NewStaticResource("83.jpg", resources.Constellations83), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "ящерица", icon: fyne.NewStaticResource("84.jpg", resources.Constellations84), description: "Журавль - символ мудрости"})

	icons = append(icons, constellationsIconInfo{name: "заяц", icon: fyne.NewStaticResource("85.jpg", resources.Constellations85), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "змееносец", icon: fyne.NewStaticResource("86.jpg", resources.Constellations86), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "змея", icon: fyne.NewStaticResource("87.jpg", resources.Constellations87), description: "Журавль - символ мудрости"})
	icons = append(icons, constellationsIconInfo{name: "золотая рыбка", icon: fyne.NewStaticResource("88.jpg", resources.Constellations88), description: "Журавль - символ мудрости"})

	//return []constellationsIconInfo{
	//	{name: "Созвездия", icon: Empty, description: ""},
	//	{name: "Андромеда", icon: андромеда, description: "Андромеда – созвездие, которое можно увидеть в северном полушарии нашей планеты. Имеет в своем арсенале три звезды со второй звездной величиной. Созвездие имеет характерный рисунок, который создают звезды входящие в него. Тянется цепочка этих светил с северо–востока в сторону юго-запада."},
	//	{name: "Большая медведица", icon: большаяМедведица, description: "Созвездие Большая Медведица – одно из наибольших созвездий по площади, третье после Гидры и Девы. Данному участку неба принадлежит более 200 звезд, и до 125 звезд из них можно различить невооруженным глазом, в безлунную ночь далеко за городом.\n\nЗвёзды в Созвездии Большой медведицы образуют так называемый Большой ковш, благодаря которому созвездие стало наиболее узнаваемым. Подобные легкоразличимые группы звезд называют «астеризм».\n Находится на северном полушарии."},
	//	{name: "Большой пёс", icon: большойПёс, description: "Большой пёс - это созвездие южного неба, находящееся в зодиакальной зоне. Включает в себя одну из ярчайших звезд на небосклоне - Сириус, который является самой яркой звездой на ночном небе. Созвездие Большого Пса было известно ещё в древности и было упомянуто в традиционной астрономии многих культур."},
	//	{name: "Близнецы", icon: близнецы, description: "Близнецы - это зодиакальное созвездие северного неба. Включает в себя две яркие звезды - Кастор и Поллукс, которые символизируют двух братьев-близнецов в греческой мифологии. В созвездии Близнецов также находится знаменитый кластер звезд М44, который иногда называют Пчелиным роем."},
	//	{name: "Чаша", icon: чаша, description: "Чаша - это небольшое созвездие южного неба, расположенное между звездами Гидра и Центавр. Включает в себя несколько ярких звезд, образующих форму чаши. В древности созвездие Чаши было ассоциировано с греческой мифологией и считалось символом празднества и пиршества."},
	//	{name: "Часы", icon: часы, description: "Часы - это небольшое созвездие южного неба, расположенное между созвездиями Летучая Рыба и Южный Крест. Включает в себя четыре яркие звезды, расположенные в форме циферблата. Название созвездия было дано в честь изобретения часов, которое помогло мореплавателям ориентироваться на море."},
	//	{name: "Дельфин", icon: дельфин, description: "Дельфин - это небольшое созвездие южного неба, представляющее собой изображение дельфина. Оно было введено в астрономический каталог в XVIII веке. В созвездии находится несколько интересных объектов, включая двойную звезду Дельфин-63 и затменную переменную звезду RX Йединого вида."},
	//	{name: "Волосы Вероники", icon: дева, description: "Волосы Вероники - это созвездие зодиакального пояса северного неба, которое наблюдается с июля по декабрь. Оно представляет собой изображение волос Вероники, которые были отрезаны Ромулом после смерти его жены. В созвездии находятся множество галактик, в том числе знаменитая галактика Вирго."},
	//	{name: "Дракон", icon: дракон, description: "Дракон - это большое созвездие северного неба, расположенное между созвездиями Большая Медведица и Геркулес. Оно было введено в астрономический каталог еще в древности и считается одним из старейших созвездий. В созвездии находится множество интересных объектов, в том числе затменная переменная звезда Альбирео."},
	//	{name: "Единорог", icon: единорог, description: "Единорог - это созвездие зодиакального пояса южного неба, наблюдаемое с ноября по май. Оно представляет собой изображение мифического единорога. В созвездии находятся различные объекты, в том числе галактики и звездные скопления."},
	//	{name: "Эридан", icon: эридан, description: "Эридан - это большое созвездие южного неба, расположенное между созвездиями Орион и Гидра. Оно было введено в астрономический каталог еще в древности и считается одним из старейших созвездий. В созвездии находится знаменитая переменная звезда Betelgeuse, а также несколько интересных галактик."},
	//	{name: "Геркулес", icon: геркулес, description: "Геркулес - это крупное звездное созвездие северного полушария, которое можно увидеть в летние месяцы. Включает в себя яркие звезды, образующие фигуру героя Геркулеса из греческой мифологии."},
	//	{name: "Жертвенник", icon: жертвенник, description: "Жертвенник - это небольшое созвездие южного полушария, расположенное между созвездиями Стрелец и Геркулес. Включает в себя яркую звезду Альфа Жертвенника и несколько интересных глубоких небесных объектов."},
	//	{name: "Жираф", icon: жираф, description: "Жираф - это небольшое созвездие южного неба, расположенное между созвездиями Киль и Южная Гидра. Включает в себя несколько ярких звезд и интересных галактик, которые можно наблюдать через телескоп."},
	//	{name: "Живописец", icon: живописец, description: "Живописец - это небольшое созвездие южного неба, расположенное между созвездиями Летучая Рыба и Гидра. Включает в себя яркую звезду Альфа Живописца и несколько интересных галактик."},
	//	{name: "Журавль", icon: журавль, description: "Журавль - это крупное созвездие южного неба, расположенное между созвездиями Щит и Микроскоп. Включает в себя яркие звезды и интересные глубокие небесные объекты, такие как галактики и туманности."},
	//	{name: "Гидра", icon: гидра, description: "Гидра - это крупное созвездие южного полушария, которое можно увидеть весной и летом. Включает в себя несколько ярких звезд и интересные глубокие небесные объекты, такие как галактики и туманности."},
	//	{name: "Голубь", icon: голубь, description: "Голубь - это маленькое созвездие на южном полушарии, которое можно увидеть с мая по август. Оно включает в себя несколько ярких звезд, но не имеет интересных глубоких небесных объектов."},
	//	{name: "Гончие псы", icon: гончиеПсы, description: "Гончие псы - это созвездие на южном полушарии, которое можно увидеть с мая по август. Оно включает в себя несколько ярких звезд и интересных глубоких небесных объектов, таких как галактики и туманности."},
	//	{name: "Хамелеон", icon: хамелеон, description: "Хамелеон - это небольшое созвездие на южном полушарии, которое можно увидеть с мая по август. Оно включает в себя несколько ярких звезд, но не имеет интересных глубоких небесных объектов."},
	//	{name: "Индеец", icon: индеец, description: "Индеец - это маленькое созвездие на южном полушарии, которое можно увидеть с мая по август. Оно включает в себя несколько ярких звезд, но не имеет интересных глубоких небесных объектов."},
	//	{name: "Кассиопея", icon: кассиопея, description: "Кассиопея - это созвездие на северном полушарии, которое можно увидеть круглый год. Оно включает в себя несколько ярких звезд и интересные глубокие небесные объекты, такие как открытые звездные скопления и туманности."},
	//	{name: "Киль", icon: киль, description: "Киль - это созвездие на южном полушарии, которое можно увидеть с мая по октябрь. Оно включает в себя несколько ярких звезд и интересные глубокие небесные объекты, такие как открытые звездные скопления и туманности."},
	//	{name: "Кит", icon: кит, description: "Кит - это крупное созвездие на северном полушарии, которое можно увидеть в зимнее время года. Включает в себя яркие звезды и интересные глубокие небесные объекты, такие как галактики и туманности."},
	//	{name: "Компас", icon: компас, description: "Компас - это небольшое созвездие на южном полушарии, которое можно увидеть в течение всего года. Включает в себя яркие звезды и интересные двойные звезды."},
	//	{name: "Корма", icon: корма, description: "Корма - это крупное созвездие на южном полушарии, которое можно увидеть в зимнее время года. Включает в себя яркие звезды и интересные глубокие небесные объекты, такие как галактики и туманности."},
	//	{name: "Козерог", icon: козерог, description: "Козерог - это крупное созвездие на южном полушарии, которое можно увидеть в летнее время года. Включает в себя яркие звезды и интересные глубокие небесные объекты, такие как галактики и туманности."},
	//	{name: "Лебедь", icon: лебедь, description: "Лебедь - это крупное созвездие на северном полушарии, изображающее лебедя. Включает в себя несколько ярких звезд и интересные глубокие небесные объекты, такие как галактики и звездные скопления."},
	//	{name: "Летучая рыба", icon: летучаяРыба, description: "Летучая рыба - это маленькое созвездие на южном полушарии, названное в честь рыбы, которая, согласно легендам, могла летать. Включает в себя несколько ярких звезд и интересные глубокие небесные объекты, такие как галактики и туманности."},
	//	{name: "Лев", icon: лев, description: "Лев - это созвездие зодиака, одно из самых знаменитых и крупных на небосводе. Включает в себя несколько ярких звезд, таких как Регулус, и множество интересных объектов, включая галактики и туманности."},
	//	{name: "Лира", icon: лира, description: "Лира - это созвездие, которое можно увидеть в летние месяцы на северном полушарии. Известно благодаря яркой звезде Вега, которая является одной из самых ярких звезд на небосводе."},
	//	{name: "Лисичка", icon: лисичка, description: "Лисичка - это маленькое созвездие на северном полушарии, которое можно увидеть в зимнее время. Содержит несколько интересных двойных звезд и галактик."},
	//	{name: "Малая медведица", icon: малаяМедведица, description: "Малая медведица - это созвездие, которое можно увидеть круглый год на северном полушарии. Состоит из семи ярких звезд, которые образуют форму маленькой медведицы."},
	//	{name: "Малый конь", icon: малыйКонь, description: "Малый конь - это созвездие южного полушария, которое можно увидеть весной и летом. Известно благодаря яркой звезде Альфа-Эридан, которая является одной из наиболее заметных звезд на небосводе."},
	//	{name: "Малый лев", icon: малыйЛев, description: "Малый лев - это созвездие южного полушария, которое можно увидеть в ночное время в летние месяцы. Состоит из нескольких ярких звезд, расположенных в форме льва."},
	//	{name: "Малый пёс", icon: малыйПёс, description: "Малый пёс - это созвездие северного полушария, которое можно увидеть в зимнее время. Состоит из нескольких ярких звезд, расположенных в форме маленькой собаки."},
	//	{name: "Микроскоп", icon: микроскоп, description: "Микроскоп - это созвездие южного полушария, которое можно увидеть в течение всего года. Несмотря на то, что оно не очень заметно, оно содержит несколько интересных объектов, таких как галактики и звездные скопления."},
	//	{name: "Муха", icon: муха, description: "Муха - это созвездие южного полушария, которое можно увидеть в зимние месяцы. Состоит из нескольких ярких звезд, расположенных в форме мухи."},
	//	{name: "Насос", icon: насос, description: "Насос - это созвездие южного полушария, которое можно увидеть весной и летом. Содержит несколько интересных объектов, таких как галактики и звездные скопления."},
	//	{name: "Наугольник", icon: наугольник, description: "Наугольник - это небольшое созвездие южного полушария, которое можно увидеть с местности, расположенной южнее 30 градусов северной широты. Название созвездия происходит от инструмента, который используется для определения положения звезд на небе."},
	//	{name: "Октант", icon: октант, description: "Октант - это созвездие южного полушария, которое можно увидеть с местности, расположенной южнее 30 градусов северной широты. Включает в себя несколько ярких звезд и интересные глубокие небесные объекты."},
	//	{name: "Орёл", icon: орёл, description: "Орёл - это созвездие северного полушария, которое можно увидеть в течение всего года. Включает в себя яркую звезду Альтаир и несколько интересных глубоких небесных объектов."},
	//	{name: "Орион", icon: орион, description: "Орион - это яркое созвездие зимнего неба на северном полушарии, которое можно увидеть с ноября по март. Включает в себя несколько ярких звезд и интересные глубокие небесные объекты, такие как туманность Ориона."},
	//	{name: "Овен", icon: овен, description: "Овен - это созвездие зодиака, которое можно увидеть с северного полушария зимой и весной. Включает в себя яркие звезды, такие как Хамаль и Шератан, и интересные глубокие небесные объекты."},
	//	{name: "Паруса", icon: паруса, description: "Паруса — одно из наиболее ярких созвездий южного полушария неба, представляющее собой паруса судна. Наблюдается на юге России в течение всей года, на севере — в летние месяцы. В Парусах находятся такие звезды, как Альфа и Бета Парусов, которые являются яркими звездами первой величины."},
	//	{name: "Павлин", icon: павлин, description: "Павлин — созвездие южного полушария, которое представляет собой перо павлина. Наблюдается на юге России в течение всей года, на севере — в летние месяцы. В созвездии Павлин находится звезда Альфа Павлина, являющаяся яркой звездой нулевой величины."},
	//	{name: "Печь", icon: печь, description: "Печь — созвездие южного полушария, представляющее собой печь. Наблюдается на юге России в течение всей года, на севере — в летние месяцы. В созвездии Печь находятся такие звезды, как Альфа, Бета и Гамма Печей, которые являются яркими звездами третьей величины."},
	//	{name: "Пегас", icon: пегас, description: "Пегас — созвездие северного полушария, представляющее собой крылатую лошадь. Наблюдается на территории России в течение всей года. В созвездии Пегас находится звезда Альфа Пегаса, являющаяся яркой звездой первой величины."},
	//	{name: "Пегас", icon: пегас, description: "Пегас - созвездие на северном небосклоне, изображающее крылатого коня Пегаса из греческой мифологии. Включает несколько ярких звезд, в том числе Альфа Пегаса, которая является ярчайшей звездой в этом созвездии и одной из ярчайших на ночном небе."},
	//	{name: "Персей", icon: персей, description: "Персей - созвездие на северном небосклоне, названное в честь героя греческой мифологии Персея. Включает несколько звезд, среди которых наиболее яркая Альфа Персея. В этом созвездии находится знаменитая Персеева двойная звезда, которую можно наблюдать даже в небольшом телескопе."},
	//	{name: "Феникс", icon: феникс, description: "Феникс - созвездие на южном небосклоне, названное в честь Феникса - мифической птицы, способной воскресать из пепла. Включает несколько слабых звезд, среди которых наиболее яркая Альфа Феникса."},
	//	{name: "Рак", icon: рак, description: "Рак - созвездие на северном небосклоне, которое представляет собой изображение рака из греческой мифологии. Включает несколько звезд, среди которых наиболее яркая - Альдебаран. В этом созвездии находится знаменитый Раковидный туман, который можно наблюдать даже в небольшом телескопе."},
	//	{name: "Райская птица", icon: райскаяПтица, description: "Райская птица - созвездие, которое можно увидеть в южном полушарии. Оно было впервые описано в античные времена и представляет собой группу ярких звезд, которые образуют изображение птицы с раскрытыми крыльями. Согласно мифологии, эта птица является символом вечной жизни и бессмертия."},
	//	{name: "Резец", icon: резец, description: "Созвездие Резец можно увидеть в северном полушарии, и оно получило свое название благодаря форме, напоминающей инструмент для резки. Созвездие состоит из нескольких ярких звезд, которые создают изображение линии, направленной от юга к северу. Согласно мифологии, это созвездие представляет собой орудие греческого бога Зевса, которым он обрубал голову своему отцу Кроносу."},
	//	{name: "Рыбы", icon: рыбы, description: "Рыбы - это созвездие, которое можно увидеть в зимнее время года. Созвездие состоит из двух ярких звезд, которые образуют изображение двух рыб, плавающих вместе в противоположных направлениях. Согласно мифологии, эти рыбы представляют собой богов воды, которые помогали людям пересекать реки и океаны.\n\n"},
	//	{name: "Рысь", icon: рысь, description: "Созвездие Рысь можно увидеть в северном полушарии, и оно получило свое название благодаря сходству с животным, которое обитает в лесах. Созвездие состоит из нескольких ярких звезд, которые создают изображение головы и передних лап рыси. Согласно мифологии, это созвездие представляет собой животное, которое сопровождало бога во время охоты."},
	//	{name: "Секстант", icon: секстант, description: "Секстант - это созвездие, которое можно увидеть в южном полушарии. Оно было создано в XVIII веке и представляет собой изображение инструмента для измерения углов, который использовался для определения расстояния между звездами. Созвездие состоит из нескольких ярких звезд, которые создают изображение угломера."},
	//	{name: "Сетка", icon: сетка, description: "Сетка - это созвездие, которое можно увидеть в южном полушарии. Оно было создано в XVIII веке и представляет собой изображение сетки, используемой для навигации по звездам. Созвездие состоит из нескольких ярких звезд, которые создают изображение сетки, состоящей из прямых линий, направленных в разные стороны."},
	//	{name: "Северная сетка", icon: севернаяКорона, description: "Северная сетка - это созвездие, которое можно увидеть в северном полушарии. Оно было создано в XVII веке и представляет собой изображение короны, которую носят монархи. Созвездие состоит из нескольких ярких звезд, которые создают изображение короны, состоящей из шести зубцов."},
	//	{name: "Щит", icon: щит, description: "Щит - это созвездие, которое можно увидеть в северном полушарии. Оно было впервые описано в античные времена и представляет собой изображение щита, используемого в бою. Созвездие состоит из нескольких ярких звезд, которые создают изображение щита, имеющего форму полукруга."}, {name: "Скорпион", icon: скорпион, description: "Скорпион - зодиакальное созвездие, находящееся на южном небосклоне. Созвездие получило своё название благодаря сильной похожести на насекомое-скорпиона. В астрологии Скорпион относится к элементу воды и считается одним из самых загадочных знаков зодиака. Яркой звездой созвездия является Антарес, красный сверхгигант, который является одним из ярчайших объектов на ночном небе."},
	//	{name: "Скульптор", icon: скульптор, description: "Скульптор - созвездие южного полушария неба, находящееся между созвездиями Феникса и Павлина. Получило своё название в честь древнеримского бога искусства и ремесла - Сципиона. Яркой звездой созвездия является Альфа Скульптора, яркость которой достигает 4,3 звездной величины."},
	//	{name: "Столовая гора", icon: столоваяГора, description: "Столовая гора - созвездие южного небосклон, названное в честь горы Таволга, на которой, по легенде, древние скандинавские боги устраивали пир. Яркой звездой созвездия является Альфа Столовой Горы, яркость которой составляет 2,8 звездной величины."},
	//	{name: "Стрела", icon: стрела, description: "Стрела - зодиакальное созвездие, находящееся на северном небосклоне, между созвездиями Скорпиона и Кассиопеи. Созвездие символизирует легендарного героя Геракла, который использовал стрелу, чтобы убить орла Этния во время одного из своих трудов. Яркой звездой созвездия является Альфа Стрелы, яркость которой достигает 2,6 звездной величины."},
	//	{name: "Стрелец", icon: стрелец, description: "Стрелец - зодиакальное созвездие, находящееся на южном небосклоне. Символизирует легендарного героя-стрелка, который умел метко стрелять из лука. В астрологии Стрелец относится к элементу огня и считается одним из самых оптимистических знаков зодиака. Яркой звездой созвездия является Саджиттариус, яркость которой достигает 1,8 звездной величины."},
	//	{name: "Цефей", icon: цефей, description: "Цефей - созвездие северного полушария, расположенное между созвездиями Кассиопеи и Лебедя. Получило своё название в честь царя из древнегреческой мифологии, который спас свою дочь от морского чудовища. Яркой звездой созвездия является Альфа Цефея, яркость которой достигает 2,5 звездной величины."},
	//	{name: "Центавр", icon: центавр, description: "Центавр - созвездие южного небосклона, символизирующее легендарных существ, которые имели тело человека и тело коня. Яркой звездой созвездия является Альфа Центавра, яркость которой достигает -0,3 звездной величины. Эта звезда является третьей по яркости на ночном небе и является ближайшей к Земле звездой после Солнца."},
	//	{name: "Циркуль", icon: циркуль, description: "Циркуль - созвездие северного небосклона, расположенное между созвездиями Большой Медведицы и Геркулеса. Получило своё название благодаря похожести на инструмент, который используется для измерения углов и длин. Яркой звездой созвездия является Альфа Циркуля, яркость которой достигает 3,5 звездной величины."},
	//	{name: "Телескоп", icon: телескоп, description: "Телескоп - небольшое южное созвездие, которое было введено в астрономический каталог в XVIII веке. Его яркие звезды были известны индейцам Анд, а также европейским мореплавателям. Главной звездой созвездия является Альфа Телескопа, которая является одной из самых ярких звезд на юге неба."},
	//	{name: "Телец", icon: телец, description: "Телец - одно из знаков зодиака и созвездие, находящееся на северной полусфере земного шара. В астрономии Телец представляет собой созвездие из нескольких ярких звезд, которые образуют форму рогатого скота. Одна из звезд этого созвездия - Альдебаран - является одной из наиболее ярких звезд на ночном небе."},
	//	{name: "Треугольник", icon: треугольник, description: "Треугольник - небольшое звездное созвездие на северной полусфере земного шара. Его звезды были известны еще с древности и были использованы для ориентации в космосе. Главной звездой созвездия является Бета Треугольника, которая является яркой звездой с зеленоватым оттенком."},
	//	{name: "Тукан", icon: тукан, description: "Тукан - большое южное созвездие, которое было введено в астрономический каталог в XVIII веке. Его яркие звезды были известны индейцам Анд, а также европейским мореплавателям. Главной звездой созвездия является Альфа Тукана, которая является одной из самых ярких звезд на юге неба."},
	//	{name: "Южная гидра", icon: южнаяГидра, description: "Южная Гидра — одно из крупнейших созвездий южного полушария. Созвездие не имеет ярких звезд, но содержит множество интересных объектов, включая галактики, квазары и группы звезд."},
	//	{name: "Южная корона", icon: южнаяКорона, description: "Южная Корона — небольшое созвездие на южном небосклоне, которое легко узнать благодаря яркой звезде Альфа Короны. В созвездии можно найти интересные объекты, такие как двойную звезду Альфа Короны и галактику NGC 6720."},
	//	{name: "Южная рыба", icon: южнаяРыба, description: "Южная Рыба — небольшое созвездие на южном небосклоне. Созвездие не содержит ярких звезд, но в нем можно найти интересные объекты, такие как галактики NGC 55 и NGC 300."},
	//	{name: "Южный треугольник", icon: южныйТреугольник, description: "Южный Треугольник — маленькое созвездие южного небосклоне, которое можно увидеть только с южных широт. Созвездие не содержит ярких звезд, но в нем можно найти интересные объекты, такие как галактику NGC 602."},
	//	{name: "Южный крест", icon: южныйКрест, description: "Южный Крест — наиболее яркое и знакомое созвездие южного полушария. Оно названо так из-за своей формы, напоминающей крест. В созвездии можно найти множество интересных объектов, включая галактики, звездные кластеры и туманности."},
	//	{name: "Весы", icon: весы, description: "Весы - созвездие, расположенное на эклиптике, названное в честь весов, символизирующих баланс и справедливость. В древней мифологии астрономический объект был связан с богиней Темидой, которая держала в руках весы и решала споры между богами и людьми."},
	//	{name: "Водолей", icon: водолей, description: "Водолей - знак зодиака и созвездие, расположенное на эклиптике. В древней мифологии Водолей был связан с богом Ганимедом, который был увлечен Юпитером и стал чашей для нектара богов."},
	//	{name: "Волк", icon: волк, description: "Волк - созвездие, расположенное на северном небосклоне. В древних мифах волк был связан с богом Марсом и являлся символом войны и силы."},
	//	{name: "Волопас", icon: волопас, description: "Волопас - созвездие на северном небосклоне, названное в честь охотника, который держит в руках длинную палку с крюком, используемую для ловли зверей."},
	//	{name: "Волосы Вероники", icon: волосыВероники, description: "Волосы Вероники - созвездие на северном небосклоне, расположенное между созвездиями Волка и Волопаса. В древней мифологии созвездие было связано с историей девушки Вероники, которая согласилась стать служанкой богини Афродиты и была вознаграждена за свою преданность небесным венком волос."},
	//	{name: "Ворон", icon: ворон, description: "Ворон (лат. Corvus) - созвездие южного полушария неба, изображающее ворона. Ворон находится между созвездиями Райский птице, Центавр и Гидра. Ворон был известен еще древними греками, которые связывали его с мифами о Аполлоне, Афродите, Аресе и других богах."},
	//	{name: "Возничий", icon: возничий, description: "Возничий (лат. Auriga) - созвездие северного полушария неба, изображающее возницу на карете. Возничий находится между созвездиями Большая Медведица и Персей. Возничий был известен еще древними греками и римлянами, которые связывали его с мифами о Эразме, Икаре, Меламподе и других героях."},
	//	{name: "Ящерица", icon: ящерица, description: "Ящерица (лат. Lacerta) - небольшое созвездие северного полушария неба, которое было введено в список 88 современных созвездий в XVIII веке. Ящерица находится между созвездиями Кассиопея, Лебедь и Стрела. Несмотря на то, что ящерица не имеет высокой яркости и не выделяется на небе, она представляет интерес для любителей астрономии."},
	//	{name: "Заяц", icon: заяц, description: "Заяц (лат. Lepus) - созвездие южного полушария неба, изображающее зайца. Заяц находится между созвездиями Орион, Гидра и Лев. Заяц был известен еще древними греками, которые связывали его с мифами о Артемиде и других богинях."},
	//	{name: "Змееносец", icon: змееносец, description: "Созвездие Змееносец - одно из крупных созвездий южного полушария. Оно было введено в античную астрономию ещё в древности, и в современной астрономии относится к 88 созвездиям. Созвездие получило своё название благодаря тому, что в нём находится яркая звезда Альфа Офиуча, которую астрономы античности считали головой змеи."},
	//	{name: "Змея", icon: змея, description: "Созвездие Змея - одно из крупных созвездий северной полусферы. Оно было введено в античную астрономию ещё в древности, и в современной астрономии относится к 88 созвездиям. Созвездие получило своё название благодаря тому, что в нём находится цепочка звёзд, напоминающая изогнутую змею."},
	//	{name: "Золотая рыбка", icon: золотаяРыбка, description: "Созвездие Золотая Рыбка - одно из самых маленьких и слабых созвездий южного полушария. Оно было введено в античную астрономию ещё в древности, и в современной астрономии относится к 88 созвездиям. Созвездие получило своё название благодаря тому, что в нём находится яркая звезда Фомальгаут, которую астрономы античности считали глазом рыбы."},
	//}
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

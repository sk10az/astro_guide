package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sk10az/astro_guide/resources"
)

func greetingScreen(_ fyne.Window) fyne.CanvasObject {

	министерство := widget.NewLabel("МИНИСТЕРСТВО ОБРАЗОВАНИЯ И НАУКИ ДОНЕЦКОЙ НАРОДНОЙ РЕСПУБЛИКИ\nГОСУДАРСТВЕННОЕ БЮДЖЕТНОЕ ПРОФЕССИОНАЛЬНОЕ ОБРАЗОВАТЕЛЬНОЕ УЧРЕЖДЕНИЕ \n\"ЕНАКИЕВСКИЙ МЕТАЛЛУРГИЧЕСКИЙ ТЕХНИКУМ\"")
	министерство.Alignment = fyne.TextAlignCenter // выравнивание по центру

	специальность := widget.NewLabel("09.02.03 Программирование в компьютерных системах\nЦикловой комиссии физико-математических дисциплин и программирования")
	специальность.Alignment = fyne.TextAlignCenter

	курсоваяРабота := widget.NewLabel("КУРСОВАЯ РАБОТА")
	курсоваяРабота.Alignment = fyne.TextAlignCenter
	курсоваяРабота.TextStyle = fyne.TextStyle{Bold: true}

	мдк := widget.NewLabel("по МДК.03.01 \"Технология разработки программного обеспечения\"\nСтудента группы ПКС-19 1/9")
	мдк.Alignment = fyne.TextAlignCenter
	мдк.TextStyle = fyne.TextStyle{Bold: true}

	фиоИтема := widget.NewLabel("Закса Данила Денисовича\n\"Электронный справочник звёздного неба\"")
	фиоИтема.Alignment = fyne.TextAlignCenter
	фиоИтема.TextStyle = fyne.TextStyle{Bold: true}

	верх := container.NewVBox(министерство, специальность, курсоваяРабота, мдк, фиоИтема)

	енакиево := widget.NewLabel("Енакиево - 2023")
	енакиево.Alignment = fyne.TextAlignCenter

	var res = fyne.NewStaticResource("earth_logo.png", resources.EarthLogo)
	logo := canvas.NewImageFromResource(res)
	logo.FillMode = canvas.ImageFillContain

	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(192, 192))
	} else {
		logo.SetMinSize(fyne.NewSize(200, 200))
	}

	return container.NewBorder(верх, енакиево, nil, nil, logo)
}

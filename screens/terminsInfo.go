package screens

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// TODO:
func terminsScreen(_ fyne.Window) fyne.CanvasObject {

	data := binding.BindStringList(
		&[]string{
			"АЛЬБЕДО  - Доля световой энергии, отраженная поверхностью.",
			"АПЕКС - Точка на небесной сфере, в направлении которой движется в пространстве астрономический объект.",
			"АПОГЕЙ - Наиболее удаленная от Земли точка орбиты Луны или ИСЗ.",
			"АФЕЛИЙ - Наиболее удаленная от Солнца точка орбиты планеты или иного тела Солнечной системы.",
			"БОДЕ ЗАКОН - Эмпирическое правило, указывающее приблизительное расстояние планет от Солнца.",
		},
	)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})

	return container.NewBorder(nil, add, nil, nil, list)
}

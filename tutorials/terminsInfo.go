package tutorials

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
			"Item 1 - это какой-то интересный термин, который тяжело понять.",
			"Item 2 - это какой-то интересный термин, который тяжело понять.",
			"Item 3 - это какой-то интересный термин, который тяжело понять.",
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

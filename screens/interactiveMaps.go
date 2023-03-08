package screens

//import (
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/canvas"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/layout"
//	"fyne.io/fyne/v2/theme"
//	"fyne.io/fyne/v2/widget"
//)
//
//func interactiveMapScreen(_ fyne.Window) fyne.CanvasObject {
//
//	img := canvas.NewImageFromFile("resources/img/constellations.jpg")
//
//	// Создаем канвас для отображения карты
//	canvas := fyne.NewContainerWithLayout(layout.NewMaxLayout(), img)
//	canvas.Resize(fyne.NewSize(800, 600))
//
//	// Добавляем кнопки для управления масштабом и скроллингом
//	zoomInButton := widget.NewButtonWithIcon("", theme.ZoomOutIcon(), func() {
//		img.Scale *= 1.1
//		canvas.Refresh()
//	})
//	zoomOutButton := widget.NewButtonWithIcon("", theme.ZoomOutIcon(), func() {
//		img.Scale /= 1.1
//		canvas.Refresh()
//	})
//	scrollUpButton := widget.NewButtonWithIcon("", theme.ZoomOutIcon(), func() {
//		canvas.Offset.Y -= 10
//		canvas.Refresh()
//	})
//	scrollDownButton := widget.NewButtonWithIcon("", theme.ZoomOutIcon(), func() {
//		canvas.Offset.Y += 10
//		canvas.Refresh()
//	})
//	scrollLeftButton := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
//		canvas.Offset.X -= 10
//		canvas.Refresh()
//	})
//	scrollRightButton := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
//		canvas.Offset.X += 10
//		canvas.Refresh()
//	})
//
//	// Создаем контейнер для кнопок
//	buttons := container.NewGridWithColumns(3,
//		zoomInButton, scrollUpButton, zoomOutButton,
//		scrollLeftButton, nil, scrollRightButton,
//		nil, scrollDownButton, nil,
//	)
//
//	// Создаем контейнер для отображения карты и кнопок
//	content := container.New(layout.NewBorderLayout(nil, nil, nil, buttons),
//		canvas,
//		buttons,
//	)
//
//	return content
//}

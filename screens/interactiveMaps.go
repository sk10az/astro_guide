package screens

//
//import (
//	"image"
//	"image/color"
//
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/canvas"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/widget"
//)
//
//func constellationMap() map[string]image.Point {
//	constellations := map[string]image.Point{
//		"Andromeda":        {150, 200},
//		"Antlia":           {900, 500},
//		"Apus":             {1100, 200},
//		"Aquarius":         {400, 200},
//		"Aquila":           {800, 300},
//		"Ara":              {1000, 400},
//		"Aries":            {200, 400},
//		"Auriga":           {300, 300},
//		"Bootes":           {400, 400},
//		"Caelum":           {1000, 500},
//		"Camelopardalis":   {250, 150},
//		"Cancer":           {250, 500},
//		"Canes Venatici":   {600, 400},
//		"Canis Major":      {750, 500},
//		"Canis Minor":      {600, 500},
//		"Capricornus":      {500, 300},
//		"Carina":           {950, 200},
//		"Cassiopeia":       {200, 200},
//		"Centaurus":        {800, 200},
//		"Cepheus":          {200, 100},
//		"Cetus":            {400, 500},
//		"Chamaeleon":       {1200, 300},
//		"Circinus":         {1100, 400},
//		"Columba":          {900, 400},
//		"Coma Berenices":   {500, 400},
//		"Corona Australis": {1000, 300},
//		"Corona Borealis":  {350, 250},
//		"Corvus":           {600, 450},
//		"Crater":           {550, 450},
//		"Crux":             {950, 300},
//		"Cygnus":           {300, 200}, // Добавлено недостающее созвездие
//	}
//	return constellations
//}
//
//func NewConstellationScreen() fyne.CanvasObject {
//	mapSize := fyne.NewSize(1200, 700)
//	mapImage := canvas.NewRasterWithPixels(func(x, y, w, h int) []byte {
//		img := image.NewRGBA(image.Rect(0, 0, w, h))
//		for name, pos := range constellationMap() {
//			canvas.Circle{
//				StrokeColor: color.RGBA{0, 0, 255, 255},
//				StrokeWidth: 2,
//				Position:    pos,
//				Radius:      10,
//			}.Draw(img)
//			label := canvas.NewText(name, color.White)
//			label.TextSize = 10
//			label.Move(pos)
//			label.Draw(img)
//		}
//		return img.Pix
//	}, mapSize)
//
//	return container.NewCenter(
//		container.NewVBox(
//			mapImage,
//			widget.NewButton("Back", func() {
//				mainScreen()
//			}),
//		),
//	)
//}

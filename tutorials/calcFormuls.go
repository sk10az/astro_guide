package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

// TODO:
func calcScreen(_ fyne.Window) fyne.CanvasObject {

	var_M := widget.NewEntry()
	var_m := widget.NewEntry()
	var_G := widget.NewEntry()
	var_R := widget.NewEntry()

	var answer string

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "M = ", Widget: var_M},
			{Text: "m = ", Widget: var_m},
			{Text: "G = ", Widget: var_G},
			{Text: "r = ", Widget: var_R},
		},
		OnSubmit: func() { // optional, handle form submission
			log.Println("M:", var_M.Text)
			log.Println("m:", var_m.Text)
			log.Println("G:", var_G.Text)
			log.Println("r:", var_R.Text)
		},
	}

	titleOfUniversalGravitation := widget.NewLabel("Формула свободного падения: F=(MmG)/r^2")
	deskOfUniversalGravitation := widget.NewLabel("   Для подсчета используется формула: F = G * (m1 * m2) / R^2,\n" +
		"где m — масса, R — расстояние между телами, G — гравитационная постоянная, \n" +
		"значение которой было определено экспериментально. Эта постоянная G очень \n" +
		"мала (6,67 * 10–11 м^3 / (кг * с^2)) — именно поэтому сила, с которой притягиваются \n" +
		"тела небольшой массы, нами совершенно не ощущается.\n")
	return container.NewVBox(
		container.NewVBox(titleOfUniversalGravitation, deskOfUniversalGravitation),
		form,
		widget.NewLabel("Ответ: "+answer),
	)
}

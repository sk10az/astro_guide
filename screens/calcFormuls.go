package screens

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math"
	"strconv"
)

func calcScreen(_ fyne.Window) fyne.CanvasObject {
	var_M := widget.NewEntry()
	var_m := widget.NewEntry()
	var_G := 6.67 * math.Pow(10, -11)
	var_R := widget.NewEntry()

	var answer string
	newlabel := widget.NewLabel("Ответ: ")

	form := &widget.Form{
		BaseWidget: widget.BaseWidget{},
		Items: []*widget.FormItem{ // we can specify items in the constructor
			//{Text: "G = 6.67^(-11)", Widget: var_G},
			{Text: "M = ", Widget: var_M},
			{Text: "m = ", Widget: var_m},
			{Text: "r = ", Widget: var_R},
		},
		OnSubmit: func() { // optional, handle form submission
			M, err := strconv.ParseFloat(var_M.Text, 64)
			if err != nil {
				fmt.Println(err)
				return
			}

			m, err := strconv.ParseFloat(var_m.Text, 64)
			if err != nil {
				fmt.Println(err)
				return
			}

			R, err := strconv.ParseFloat(var_R.Text, 64)
			if err != nil {
				fmt.Println(err)
				return
			}

			a11 := var_G * ((M * m) / math.Pow(R, 2))
			answer = strconv.FormatFloat(a11, 'f', -1, 64)
			newlabel.SetText("Ответ: " + answer)
			fmt.Println(a11)
		},
	}

	titleOfUniversalGravitation := widget.NewLabel("Формула всемирного тяготения: F=(MmG)/r^2")
	deskOfUniversalGravitation := widget.NewLabel("   Для подсчета используется формула: F = G * (m1 * m2) / R^2,\n" +
		"где m — масса, R — расстояние между телами, G — гравитационная постоянная, \n" +
		"значение которой было определено экспериментально. Эта постоянная G очень \n" +
		"мала (6,67 * 10–11 м^3 / (кг * с^2)) — именно поэтому сила, с которой притягиваются \n" +
		"тела небольшой массы, нами совершенно не ощущается.\n")
	return container.NewVBox(
		container.NewVBox(titleOfUniversalGravitation, deskOfUniversalGravitation),
		widget.NewLabel("G = 6.67*10^(-11)"),
		form,
		newlabel,
	)
}

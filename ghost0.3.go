package main

import (
	"fmt"
	"log"
	"sort"

	// "fyne.io/fyne"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ghost struct {
	name string
	ev   []int
}

func check(keyedValues []int, ghosts []ghost) (ghostName string) {
	// select ghost from evidence keyed in (very verbose)
	var wrong int = 0
	var found = false
	ghostName = ""
	for i, _ := range ghosts { // all ghosts
		sort.Ints(ghosts[i].ev)
		if (found == false) && ((ghosts[i].ev[0] == keyedValues[0]) && (ghosts[i].ev[1] == keyedValues[1]) && (ghosts[i].ev[2] == keyedValues[2])) {
			ghostName = ghosts[i].name
			fmt.Println(ghosts[i].name)
			found = true
		} else if (found == false) && ((ghosts[i].ev[0] == keyedValues[1] && ghosts[i].ev[1] == keyedValues[2]) || (ghosts[i].ev[1] == keyedValues[1] && ghosts[i].ev[2] == keyedValues[2]) && keyedValues[0] == 0) || ((ghosts[i].ev[0] == keyedValues[2] || ghosts[i].ev[1] == keyedValues[2] || ghosts[i].ev[2] == keyedValues[2]) && ((keyedValues[0] == 0) && (keyedValues[1] == 0))) {
			ghostName = ghosts[i].name
			fmt.Println(ghosts[i].name)
			// ghostsPos = append(ghostsPos, ghosts[i])
		} else if (found == false) && !((ghosts[i].ev[0] == keyedValues[0]) && (ghosts[i].ev[1] == keyedValues[1]) && (ghosts[i].ev[2] == keyedValues[2])) {
			wrong++
		}
	}
	// quick way to display "no ghost found"
	if wrong == 13 {
		ghostName = ghosts[0].name
		fmt.Println(ghosts[0].name)
		fmt.Println(keyedValues)
	}
	return ghostName
}

func main() {
	a := app.New()
	w := a.NewWindow("InSpectral")

	str := binding.NewString()

	// Inspectral logic
	evidence := []string{"noEvidence", "emf5", "spiritBox", "fingerprints", "ghostOrb", "ghostWriting", "freezingTemps"}

	ghosts := []ghost{
		{name: "noGhost", ev: []int{0, 0, 0}},
		{name: "Spirit", ev: []int{2, 3, 5}},
		{name: "Wraith", ev: []int{3, 6, 2}},
		{name: "Phantom", ev: []int{1, 4, 6}},
		{name: "Poltergeist", ev: []int{2, 3, 4}},
		{name: "Banshee", ev: []int{1, 3, 6}},
		{name: "Jinn", ev: []int{2, 4, 1}},
		{name: "Mare", ev: []int{2, 4, 6}},
		{name: "Revenant", ev: []int{1, 3, 5}},
		{name: "Shade", ev: []int{1, 4, 5}},
		{name: "Demon", ev: []int{2, 5, 6}},
		{name: "Yurei", ev: []int{4, 5, 6}},
		{name: "Oni", ev: []int{1, 2, 5}},
	}

	// ghostsPos := make([]ghost, 0, 12)

	// input from user - not evidence
	// var keyed4 int = 0
	// var keyed5 int = 0
	// var keyed6 int = 0
	// var keyed7 int = 0

	// keyedValuesSus := []int{keyed4, keyed5, keyed6, keyed7}
	// sort.Ints(keyedValuesSus)

	// evidence that has been found
	title1 := widget.NewLabel("Select The Evidence Has Been Found:")

	selectEvidence1 := widget.NewSelect(evidence, func(value string) {
		log.Println("Select set to", value)
	})
	selectEvidence2 := widget.NewSelect(evidence, func(value string) {
		log.Println("Select set to", value)

	})
	selectEvidence3 := widget.NewSelect(evidence, func(value string) {
		log.Println("Select set to", value)

	})

	// Evidence that has not been found
	space := widget.NewSeparator()
	title2 := widget.NewLabel("Select The Evidence That Has Not Been Found:")

	selectEvidence4 := widget.NewSelect(evidence, func(value string) {
		log.Println("Select set to", value)
	})
	selectEvidence5 := widget.NewSelect(evidence, func(value string) {
		log.Println("Select set to", value)
	})
	selectEvidence6 := widget.NewSelect(evidence, func(value string) {
		log.Println("Select set to", value)
	})
	selectEvidence7 := widget.NewSelect(evidence, func(value string) {
		log.Println("Select set to", value)
	})

	selectEvidence1.SetSelectedIndex(0)
	selectEvidence2.SetSelectedIndex(0)
	selectEvidence3.SetSelectedIndex(0)

	selectEvidence4.SetSelectedIndex(0)
	selectEvidence5.SetSelectedIndex(0)
	selectEvidence6.SetSelectedIndex(0)
	selectEvidence7.SetSelectedIndex(0)

	// clearButtonEv := widget.NewButton("Clear", func() {
	// 	selectEvidence1.ClearSelected()
	// 	selectEvidence2.ClearSelected()
	// 	selectEvidence3.ClearSelected()
	// })

	clearButtonNoEv := widget.NewButton("Clear", func() {
		selectEvidence4.SetSelectedIndex(0)
		selectEvidence5.SetSelectedIndex(0)
		selectEvidence6.SetSelectedIndex(0)
		selectEvidence7.SetSelectedIndex(0)
		var keyed1 = selectEvidence1.SelectedIndex()
		var keyed2 = selectEvidence2.SelectedIndex()
		var keyed3 = selectEvidence3.SelectedIndex()
		var keyedValues = []int{keyed1, keyed2, keyed3}
		sort.Ints(keyedValues)
		str.Set(check(keyedValues, ghosts))
	})

	clearButtonAll := widget.NewButton("Clear All", func() {
		selectEvidence1.SetSelectedIndex(0)
		selectEvidence2.SetSelectedIndex(0)
		selectEvidence3.SetSelectedIndex(0)
		selectEvidence4.SetSelectedIndex(0)
		selectEvidence5.SetSelectedIndex(0)
		selectEvidence6.SetSelectedIndex(0)
		selectEvidence7.SetSelectedIndex(0)
		str.Set(ghosts[0].name)
	})

	checkButton := widget.NewButton("Check", func() {
		var keyed1 = selectEvidence1.SelectedIndex()
		var keyed2 = selectEvidence2.SelectedIndex()
		var keyed3 = selectEvidence3.SelectedIndex()
		var keyedValues = []int{keyed1, keyed2, keyed3}
		sort.Ints(keyedValues)
		str.Set(check(keyedValues, ghosts))
	})

	ghostLabel := widget.NewLabelWithData(str)

	// keyedValues := []int{keyed1, keyed2, keyed3}

	// Render view
	tabs := container.NewAppTabs(
		container.NewTabItem("Inspect", container.NewVBox(
			title1,
			selectEvidence1,
			selectEvidence2,
			selectEvidence3,
			space,
			title2,
			selectEvidence4,
			selectEvidence5,
			selectEvidence6,
			selectEvidence7,
			space,
			ghostLabel,
			space,
			container.NewHBox(
				clearButtonNoEv,
				clearButtonAll,
				layout.NewSpacer(),
				checkButton),
		)),
		container.NewTabItem("Ghosts", widget.NewLabel("World!")),
		container.NewTabItem("Evidence", widget.NewLabel("World!")),
		container.NewTabItem("Tips", widget.NewLabel("World!")),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(tabs)
	w.CenterOnScreen()
	// w.Resize(fyne.NewSize(700, 500))
	w.ShowAndRun()
}

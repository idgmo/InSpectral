package main

import (
	"fmt"
	"log"
	"sort"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ghost struct {
	name string
	ev   []int
}

func check(keyedValues []int, ghosts []ghost) (ghostsNames []string) {
	// select ghost from evidence keyed in (very verbose)
	var wrong int = 0
	var found = false
	ghostName := ""
	ghostsPos := make([]ghost, 0, 12)
	ghostsStill := make([]string, 0, 12)
	for i, _ := range ghosts { // all ghosts
		sort.Ints(ghosts[i].ev)
		if (found == false) && ((ghosts[i].ev[0] == keyedValues[0]) && (ghosts[i].ev[1] == keyedValues[1]) && (ghosts[i].ev[2] == keyedValues[2])) {
			ghostName = ghosts[i].name
			fmt.Println(ghosts[i].name)
			found = true
		} else if (found == false) && ((ghosts[i].ev[0] == keyedValues[1] && ghosts[i].ev[1] == keyedValues[2]) || (ghosts[i].ev[1] == keyedValues[1] && ghosts[i].ev[2] == keyedValues[2]) && keyedValues[0] == 0) || ((ghosts[i].ev[0] == keyedValues[2] || ghosts[i].ev[1] == keyedValues[2] || ghosts[i].ev[2] == keyedValues[2]) && ((keyedValues[0] == 0) && (keyedValues[1] == 0))) {
			ghostName = ghosts[i].name
			fmt.Println(ghostName)
			ghostsPos = append(ghostsPos, ghosts[i])
			ghostsStill = append(ghostsStill, ghosts[i].name)
		} else if (found == false) && !((ghosts[i].ev[0] == keyedValues[0]) && (ghosts[i].ev[1] == keyedValues[1]) && (ghosts[i].ev[2] == keyedValues[2])) {
			wrong++
		}
	}
	// quick way to display "no ghost found"
	if wrong == 15 {
		ghostName = ghosts[0].name
		ghostsStill[0] = ghosts[0].name
		fmt.Println(ghosts[0].name)
		fmt.Println(keyedValues)
	}
	fmt.Println(ghostsPos)
	return ghostsStill
}

func main() {

	// TODO: print out all ghosts that match entered evidence. Currently only the last ghost is printed to a lable - use a slice
	// TODO: add static content to other tabs
	// TODO: add No-Evidence selectors w/ edited slice from evidence section
	// TODO: place No-Evidence group in a section with toggled visibility
	// TODO: adjust UI from basic template

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
		{name: "Hantu", ev: []int{3, 4, 5}},
		{name: "Yokai", ev: []int{2, 4, 5}},
	}

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

	// title2 := widget.NewLabel("Select The Evidence That Has Not Been Found:")

	// selectEvidence4 := widget.NewSelect(evidence, func(value string) {
	// 	log.Println("Select set to", value)
	// })
	// selectEvidence5 := widget.NewSelect(evidence, func(value string) {
	// 	log.Println("Select set to", value)
	// })
	// selectEvidence6 := widget.NewSelect(evidence, func(value string) {
	// 	log.Println("Select set to", value)
	// })
	// selectEvidence7 := widget.NewSelect(evidence, func(value string) {
	// 	log.Println("Select set to", value)
	// })

	// GHOSTS tab
	green := color.NRGBA{R: 0, G: 180, B: 180, A: 255}
	white := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	// black := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	// blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

	// Ghost descriptions

	// Banshee
	ghostBan1 := canvas.NewText("A Banshee is a natural hunter and will attack anything.", white)
	ghostBan2 := canvas.NewText("It has been known to stalk its prey one at a time until making its kill.", white)
	ghostBan3 := canvas.NewText("A Banshee will only target one person at a time.", white)
	ghostBan4 := canvas.NewText("Banshees fear the Crucifix and will be less aggressive when near one.", white)
	ghostBan5 := canvas.NewText("EMF Level 5, Fingerprints, Freezing Temperatures", white)

	// Demon
	ghostDem1 := canvas.NewText("A Demon is one of the worst Ghosts you can encounter.", white)
	ghostDem2 := canvas.NewText("It has been known to attack without a reason.", white)
	ghostDem3 := canvas.NewText("Demons will attack more often than any other Ghost.", white)
	ghostDem4 := canvas.NewText("Asking a Demon successful questions on the Ouija Board won't lower the user's sanity.", white)
	ghostDem5 := canvas.NewText("Freezing Temperatures, Ghost Writing, Spirit Box", white)

	// Hantu
	ghostHan1 := canvas.NewText("A rare ghost that can be found in hot climates.", white)
	ghostHan2 := canvas.NewText("They are known to attack more often in cold weather.", white)
	ghostHan3 := canvas.NewText("Moves faster in colder areas.", white)
	ghostHan4 := canvas.NewText("Moves slower in warmer areas.", white)
	ghostHan5 := canvas.NewText("Fingerprints, Ghost Orb, Ghost Writing", white)

	// Jinn
	ghostJin1 := canvas.NewText("A Jinn is a territorial ghost that will attack when threatened.", white)
	ghostJin2 := canvas.NewText("It has also been known to be able to travel at significant speed.", white)
	ghostJin3 := canvas.NewText("A Jinn will travel at a faster speed if its victim is far away.", white)
	ghostJin4 := canvas.NewText("Turning off the locations power source will prevent the Jinn from using its ability.", white)
	ghostJin5 := canvas.NewText("EMF Level 5, Ghost Orb, Spirit Box", white)

	// Demon
	// ghost1 := canvas.NewText("", white)
	// ghost2 := canvas.NewText("", white)
	// ghost3 := canvas.NewText("", white)
	// ghost4 := canvas.NewText("", white)
	// ghost5 := canvas.NewText("", white)

	ghostTextUni := canvas.NewText("Unique Strengths: ", green)
	ghostTextWeak := canvas.NewText("Weaknesses: ", green)
	ghostTextEvi := canvas.NewText("Evidence: ", green)

	selectEvidence1.SetSelectedIndex(0)
	selectEvidence2.SetSelectedIndex(0)
	selectEvidence3.SetSelectedIndex(0)

	// selectEvidence4.SetSelectedIndex(0)
	// selectEvidence5.SetSelectedIndex(0)
	// selectEvidence6.SetSelectedIndex(0)
	// selectEvidence7.SetSelectedIndex(0)

	clearButtonEv := widget.NewButton("Clear", func() {
		selectEvidence1.SetSelectedIndex(0)
		selectEvidence2.SetSelectedIndex(0)
		selectEvidence3.SetSelectedIndex(0)
		str.Set(ghosts[0].name)
	})

	// clearButtonNoEv := widget.NewButton("Clear", func() {
	// 	selectEvidence4.SetSelectedIndex(0)
	// 	selectEvidence5.SetSelectedIndex(0)
	// 	selectEvidence6.SetSelectedIndex(0)
	// 	selectEvidence7.SetSelectedIndex(0)
	// 	var keyed1 = selectEvidence1.SelectedIndex()
	// 	var keyed2 = selectEvidence2.SelectedIndex()
	// 	var keyed3 = selectEvidence3.SelectedIndex()
	// 	var keyedValues = []int{keyed1, keyed2, keyed3}
	// 	sort.Ints(keyedValues)
	// 	str.Set(check(keyedValues, ghosts))
	// })

	// clearButtonAll := widget.NewButton("Clear All", func() {
	// 	selectEvidence1.SetSelectedIndex(0)
	// 	selectEvidence2.SetSelectedIndex(0)
	// 	selectEvidence3.SetSelectedIndex(0)
	// 	selectEvidence4.SetSelectedIndex(0)
	// 	selectEvidence5.SetSelectedIndex(0)
	// 	selectEvidence6.SetSelectedIndex(0)
	// 	selectEvidence7.SetSelectedIndex(0)
	// 	str.Set(ghosts[0].name)
	// })

	ghostsLeft := make([]string, 0, 12)
	ghostsLeft = append(ghostsLeft, ghosts[0].name)

	checkButton := widget.NewButton("Check", func() {
		var keyed1 = selectEvidence1.SelectedIndex()
		var keyed2 = selectEvidence2.SelectedIndex()
		var keyed3 = selectEvidence3.SelectedIndex()
		var keyedValues = []int{keyed1, keyed2, keyed3}
		sort.Ints(keyedValues)
		ghostsLeft = check(keyedValues, ghosts)

	})

	ghostList := widget.NewList(
		func() int {
			return len(ghostsLeft)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(ghostsLeft[i])
		})

	// keyedValues := []int{keyed1, keyed2, keyed3}

	// Render view
	tabs := container.NewAppTabs(
		container.NewTabItem("Inspect",
			container.NewGridWithRows(3,

				container.NewVBox(
					title1,
					selectEvidence1,
					selectEvidence2,
					selectEvidence3,
					// space,
					// title2,
					// selectEvidence4,
					// selectEvidence5,
					// selectEvidence6,
					// selectEvidence7,
					space,
				),
				container.NewMax(
					ghostList,
					space,
				),

				container.NewVBox(
					container.NewHBox(
						clearButtonEv,
						// clearButtonAll,
						layout.NewSpacer(),
						checkButton,
					),
				),
			),
		),
		container.NewTabItem("Ghosts",
			container.NewVScroll(
				container.NewPadded(
					container.NewVBox(
						widget.NewLabel("Banshee"),
						ghostBan1,
						ghostBan2,
						container.NewHBox(
							ghostTextUni,
							ghostBan3,
						),
						container.NewHBox(
							ghostTextWeak,
							ghostBan4,
						),
						container.NewHBox(
							ghostTextEvi,
							ghostBan5,
						),
						space,
						widget.NewLabel("Demon"),
						ghostDem1,
						ghostDem2,
						container.NewHBox(
							ghostTextUni,
							ghostDem3,
						),
						container.NewHBox(
							ghostTextWeak,
							ghostDem4,
						),
						container.NewHBox(
							ghostTextEvi,
							ghostDem5,
						),
						space,
						widget.NewLabel("Hantu"),
						ghostHan1,
						ghostHan2,
						container.NewHBox(
							ghostTextUni,
							ghostHan3,
						),
						container.NewHBox(
							ghostTextWeak,
							ghostHan4,
						),
						container.NewHBox(
							ghostTextEvi,
							ghostHan5,
						),
						space,
						widget.NewLabel("Jinn"),
						ghostJin1,
						ghostJin2,
						container.NewHBox(
							ghostTextUni,
							ghostJin3,
						),
						container.NewHBox(
							ghostTextWeak,
							ghostJin4,
						),
						container.NewHBox(
							ghostTextEvi,
							ghostJin5,
						),
						space,
					)))),
		container.NewTabItem("Evidence", widget.NewLabel("wowie!")),
		container.NewTabItem("Tips", widget.NewLabel("World!")),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(tabs)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 500))
	w.ShowAndRun()
}

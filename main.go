package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ui struct {
	current *note
	notes   []*note

	content *widget.Entry
	list    *fyne.Container
}

func (u *ui) setNote(n *note) {
	u.content.SetText(n.content)
	u.current = n
	u.refreshList()
}

func (u *ui) refreshList() {
	u.list = container.NewVBox()
	for _, n := range u.notes {
		thisNote := n
		button := widget.NewButton(title(thisNote), func() {
			u.setNote(thisNote)
		})

		u.list.Add(button)
	}
}

func loadUI(notes []*note) fyne.CanvasObject {

	u := &ui{notes: notes, content: widget.NewMultiLineEntry(), list: container.NewVBox()}

	u.refreshList()

	if len(notes) > 0 {
		u.setNote(notes[0])
	}

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
	)

	side := container.New(layout.NewBorderLayout(bar, nil, nil, nil),
		bar, u.list)

	split := container.NewHSplit(side, u.content)
	split.Offset = 0.25
	return split
}

func main() {
	a := app.New()
	w := a.NewWindow("Notes")

	list := []*note{
		{content: "Note1\n Content 1"},
		{content: "Another\nWith\nLines"},
	}

	w.SetContent(loadUI(list))

	w.ShowAndRun()
}

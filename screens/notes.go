package screens

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"runtime"
	"strings"
)

const (
	countKey       = "notecount"
	noteKey        = "note%d"
	noteDeletedKey = "note%ddeleted"
)

type ui struct {
	current *note
	notes   *notelist

	content *widget.Entry
	list    *widget.List
}

type note struct {
	content binding.String
	deleted binding.Bool
}

func (n *note) title() binding.String {
	return newTitleString(n.content)
}

func (l *notelist) load() {
	l.all = nil
	count := l.pref.Int(countKey)
	if count == 0 {
		return
	}

	for i := count - 1; i >= 0; i-- {
		key := fmt.Sprintf(noteKey, i)
		deleteKey := fmt.Sprintf(noteDeletedKey, i)
		content := binding.BindPreferenceString(key, l.pref)
		deleted := binding.BindPreferenceBool(deleteKey, l.pref)
		l.all = append(l.all, &note{content, deleted})
	}
}

type titleString struct {
	binding.String
}

func (t *titleString) Get() (string, error) {
	content, err := t.String.Get()
	if err != nil {
		return "Error", err
	}

	if content == "" {
		return "Untitled", nil
	}

	return strings.SplitN(content, "\n", 2)[0], nil
}

func (t *titleString) Set(string) error {
	return errors.New("cannot set content from title")
}

func newTitleString(in binding.String) binding.String {
	return &titleString{in}
}

func (l *notelist) notes() []*note {
	var visible []*note
	for _, n := range l.all {
		if del, _ := n.deleted.Get(); del {
			continue
		}
		visible = append(visible, n)
	}
	return visible
}

type notelist struct {
	all  []*note
	pref fyne.Preferences
}

func (u *ui) setNote(n *note) {
	u.content.Unbind()
	if n == nil {
		u.content.SetText(u.placeholderContent())
		return
	}
	u.current = n
	u.content.Bind(n.content)
	u.content.Validator = nil
	u.list.Refresh()
}

func (u *ui) placeholderContent() string {
	text := "Welcome!\nTap '+' in the toolbar to add a note."
	if fyne.CurrentDevice().HasKeyboard() {
		modifier := "ctrl"
		if runtime.GOOS == "darwin" {
			modifier = "cmd"
		}
		text += fmt.Sprintf("\n\nOr use the keyboard shortcut %s+N.", modifier)
	}
	return text
}

func (u *ui) loadUI() fyne.CanvasObject {
	u.content = widget.NewMultiLineEntry()
	u.content.SetText(u.placeholderContent())

	u.list = u.buildList()

	visible := u.notes.notes()
	if len(visible) > 0 {
		u.setNote(visible[0])
		u.list.Select(0)
	}

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			//u.addNote()

		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			//u.removeCurrentNote()
		}),
	)

	side := fyne.NewContainerWithLayout(layout.NewBorderLayout(bar, nil, nil, nil),
		bar, container.NewVScroll(u.list))

	return newAdaptiveSplit(side, u.content)
}

func notesScreen(_ fyne.Window) fyne.CanvasObject {

	return container.NewVBox(
		widget.NewLabel("Notes"),
	)
}

func (u *ui) buildList() *widget.List {
	l := widget.NewList(
		func() int {
			return len(u.notes.notes())
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Title")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			l := obj.(*widget.Label)
			n := u.notes.notes()[id]
			l.Bind(n.title())
		})

	l.OnSelected = func(id widget.ListItemID) {
		n := u.notes.notes()[id]
		u.setNote(n)
	}

	return l
}

func newAdaptiveSplit(left, right fyne.CanvasObject) *fyne.Container {
	split := container.NewHSplit(left, right)
	split.Offset = 0.33
	return container.New(&adaptiveLayout{split: split}, split)
}

type adaptiveLayout struct {
	split *container.Split
}

func (a *adaptiveLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	dev := fyne.CurrentDevice()

	a.split.Horizontal = !dev.IsMobile() || fyne.IsHorizontal(dev.Orientation())
	objects[0].Resize(size)
}

func (a *adaptiveLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return a.split.MinSize()
}

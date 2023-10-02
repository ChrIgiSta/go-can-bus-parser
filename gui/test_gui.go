package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TestGui struct {
	fyneApp    fyne.App
	mainWindow fyne.Window
	table      *widget.Table
	items      []Item
}

type Item struct {
	Descriptor string
	Value      any
	Unit       string
}

func NewTestGui() *TestGui {
	return &TestGui{
		fyneApp: app.New(),
		items:   make([]Item, 0),
	}
}

func (g *TestGui) Draw() {
	myWindow := g.fyneApp.NewWindow("Test GUI")

	g.createNewTable()

	vScroll := container.NewVScroll(
		g.table,
	)

	myWindow.SetContent(vScroll)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}

func (g *TestGui) createNewTable() {
	g.table = widget.NewTable(
		func() (int, int) {
			return len(g.items), 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template XYZXYZXYZ")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			item := g.items[i.Row]

			switch i.Col {
			case 0:
				o.(*widget.Label).SetText(item.Descriptor)
			case 1:
				o.(*widget.Label).SetText(fmt.Sprintf("%v %s", item.Value, item.Unit))
			}
		})
}

func (g *TestGui) UpdateValue(description string, value any, unit string) {
	defer g.refreshTable()

	for i, item := range g.items {
		if item.Descriptor == description {
			g.items[i].Value = value
			return
		}
	}
	// not found
	g.items = append(g.items, Item{
		Descriptor: description,
		Value:      value,
		Unit:       unit,
	})
}

func (g *TestGui) refreshTable() {
	g.table.Refresh()
}

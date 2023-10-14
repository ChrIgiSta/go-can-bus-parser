/**
 * Copyright © 2023, Staufi Tech - Switzerland
 * All rights reserved.
 *
 *   ________________________   ___ _     ________________  _  ____
 *  / _____  _  ____________/  / __|_|   /_______________  | | ___/
 * ( (____ _| |_ _____ _   _ _| |__ _      | |_____  ____| |_|_
 *  \____ (_   _|____ | | | (_   __) |     | | ___ |/ ___)  _  \
 *  _____) )| |_/ ___ | |_| | | |  | |     | | ____( (___| | | |
 * (______/  \__)_____|____/  |_|  |_|     |_|_____)\____)_| |_|
 *
 *
 *  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 *  AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 *  IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 *  ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
 *  LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 *  CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 *  SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 *  INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 *  CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 *  ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 *  POSSIBILITY OF SUCH DAMAGE.
 */

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

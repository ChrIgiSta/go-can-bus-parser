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
	"image/color"
	"os"
	"time"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	log "github.com/ChrIgiSta/go-utils/logger"
)

const (
	GUI_MAIN_VIEW_NAME = "Main View"

	GUI_APP_NAME      = "ChrIgiSta Opel BoardComputer"
	GUI_DEFAULT_TITLE = "\nOpel Astra H OPC\n"

	GUI_LABLE_CONSUMPTION_EN      = "Consumtion"
	GUI_LABLE_RANGE_EN            = "Range"
	GUI_LABLE_CURRENT_CONSUMPTION = "Current Consumtion"

	GUI_OVERLAY_SHOW_DURATION = 2
)

type AcMode int

const (
	AcModeTop    = 0
	AcModeMid    = 1
	AcModeBot    = 2
	AcModeTopMid = 3
	AcModeTopBod = 4
	AcModeMidBod = 5
	AcModeAll    = 6
	AcModeAuto   = 7
)

type carData struct {
	Icon  string
	Lable string
	Value string
	Unit  string
}

var carDataSpacer carData = carData{}

type MainViewData struct {
	consumtion         carData
	currentConsumtion  carData
	rangeVal           carData
	coolantTemperature carData
}

type ClimaSection struct {
	outdoorTemperature *widget.Label
	time               *widget.Label
	fanSpeed           *widget.Label
	temperature        *widget.Label
	ecoMode            *widget.Icon
	ecoModePlaceholder *widget.Icon
	acModeBox          *fyne.Container

	overlay *widget.Label
	popup   *widget.PopUp

	climaContent *fyne.Container
	acContainer  *fyne.Container

	fanIcon *widget.Icon
}

type OriginalGui struct {
	fyneApp    fyne.App
	mainWindow fyne.Window

	title *widget.Label

	mainData      MainViewData
	mainDataTable *widget.Table
	mainContent   *fyne.Container

	menu       *fyne.Container
	viewLayout fyne.Layout
	menuView   *fyne.Container

	clima ClimaSection

	placehoderIcon *widget.Icon
}

var (
	acModeMap map[AcMode]*widget.Icon

	ardenBlue = color.NRGBA{R: 36, G: 105, B: 182, A: 150}
)

func NewOriginalGui() *OriginalGui {
	acTop, err := os.ReadFile("gui/icons/ac_top.png")
	if err != nil {
		log.Error("gui", "load ac icon: %v", err)
	}
	acMid, err := os.ReadFile("gui/icons/ac_mid.png")
	if err != nil {
		log.Error("gui", "load ac icon: %v", err)
	}
	acBot, err := os.ReadFile("gui/icons/ac_bot.png")
	if err != nil {
		log.Error("gui", "load ac icon: %v", err)
	}
	acTopMid, err := os.ReadFile("gui/icons/ac_top_mid.png")
	if err != nil {
		log.Error("gui", "load ac icon: %v", err)
	}
	acTopBot, err := os.ReadFile("gui/icons/ac_top_bot.png")
	if err != nil {
		log.Error("gui", "load ac icon: %v", err)
	}
	acMidBot, err := os.ReadFile("gui/icons/ac_mid_bot.png")
	if err != nil {
		log.Error("gui", "load ac icon: %v", err)
	}
	acAll, err := os.ReadFile("gui/icons/ac_all.png")
	if err != nil {
		log.Error("gui", "load ac icon: %v", err)
	}
	acAuto, err := os.ReadFile("gui/icons/ac_auto.png")
	if err != nil {
		log.Error("gui", "load ac icon: %v", err)
	}

	fanB, err := os.ReadFile("gui/icons/fan.png")
	if err != nil {
		log.Error("gui", "read fan icon: %v", err)
	}
	ecoB, err := os.ReadFile("gui/icons/eco.png")
	if err != nil {
		log.Error("gui", "read eco icon: %v", err)
	}

	placeholderB, err := os.ReadFile("gui/icons/placeholder.png")
	if err != nil {
		log.Error("gui", "read placehoder icon: %v", err)
	}

	acModeMap = map[AcMode]*widget.Icon{
		AcModeTop:    widget.NewIcon(fyne.NewStaticResource("ac-top", acTop)),
		AcModeMid:    widget.NewIcon(fyne.NewStaticResource("ac-mid", acMid)),
		AcModeBot:    widget.NewIcon(fyne.NewStaticResource("ac-bot", acBot)),
		AcModeTopMid: widget.NewIcon(fyne.NewStaticResource("ac-top-mid", acTopMid)),
		AcModeTopBod: widget.NewIcon(fyne.NewStaticResource("ac-top-bot", acTopBot)),
		AcModeMidBod: widget.NewIcon(fyne.NewStaticResource("ac-mid-bot", acMidBot)),
		AcModeAll:    widget.NewIcon(fyne.NewStaticResource("ac-all", acAll)),
		AcModeAuto:   widget.NewIcon(fyne.NewStaticResource("ac-auto", acAuto)),
	}

	acModeContainer := container.NewHBox(
		acModeMap[AcModeTop],
		acModeMap[AcModeMid],
		acModeMap[AcModeBot],
		acModeMap[AcModeTopMid],
		acModeMap[AcModeTopBod],
		acModeMap[AcModeMidBod],
		acModeMap[AcModeAll],
		acModeMap[AcModeAuto],
	)

	return &OriginalGui{
		placehoderIcon: widget.NewIcon(fyne.NewStaticResource("placehoder", placeholderB)),
		fyneApp:        app.New(),
		title: widget.NewLabelWithStyle(GUI_DEFAULT_TITLE,
			fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		mainData: MainViewData{
			consumtion: carData{
				Icon:  "gui/icons/fullconsumption.png",
				Lable: GUI_LABLE_CONSUMPTION_EN,
				Value: "0",
				Unit:  "l/100km",
			},
			currentConsumtion: carData{
				Icon:  "gui/icons/fullconsumption.png",
				Lable: GUI_LABLE_CURRENT_CONSUMPTION,
				Value: "0",
				Unit:  "l/h",
			},
			rangeVal: carData{
				Icon:  "gui/icons/range.png",
				Lable: GUI_LABLE_RANGE_EN,
				Value: "0",
				Unit:  "km",
			},
			coolantTemperature: carData{
				Icon:  "gui/icons/coolant_temp.png",
				Lable: GUI_LABLE_RANGE_EN,
				Value: "0",
				Unit:  "°C",
			},
		},
		clima: ClimaSection{
			outdoorTemperature: widget.NewLabel("25°C"),
			time: widget.NewLabel("Current Time: " +
				time.Now().Format("15:04")),
			fanSpeed:           widget.NewLabel("0"),
			temperature:        widget.NewLabel("UN"),
			ecoMode:            widget.NewIcon(fyne.NewStaticResource("eco", ecoB)),
			ecoModePlaceholder: widget.NewIcon(fyne.NewStaticResource("eco-placehoder", placeholderB)),
			fanIcon:            widget.NewIcon(fyne.NewStaticResource("fan", fanB)),
			acModeBox:          acModeContainer,

			overlay: widget.NewLabel("No Info"),
		},
	}
}

func (g *OriginalGui) ShowDefaultView() {

	g.mainWindow = g.fyneApp.NewWindow(GUI_APP_NAME)

	// g.mainWindow.SetFullScreen(true)
	g.mainWindow.Resize(fyne.NewSize(400, 200))

	g.title.Move(fyne.NewPos(0, 10))

	dividerLine := canvas.NewRectangle(ardenBlue)
	dividerLine.SetMinSize(fyne.NewSize(400, 1))

	g.clima.outdoorTemperature.Alignment = fyne.TextAlignLeading
	g.clima.time.Alignment = fyne.TextAlignTrailing

	g.menu = container.NewVBox(
		widget.NewButtonWithIcon("\n\t Home                     \n", theme.HomeIcon(), g.MenuMain),
		widget.NewButtonWithIcon("\n\tBC                       \n", theme.ComputerIcon(), g.MenuBC),
		widget.NewButtonWithIcon("\n\tDashboard                \n", theme.FileApplicationIcon(), g.MenuDashboard),
	)

	g.clima.ecoModePlaceholder.Hide()
	g.clima.acContainer = container.NewHBox(
		g.clima.temperature,
		g.clima.acModeBox,
		g.clima.ecoMode,
		g.clima.ecoModePlaceholder,
		g.clima.fanIcon,
		g.clima.fanSpeed)
	g.clima.climaContent = container.NewHBox(
		container.NewVBox(g.clima.outdoorTemperature),
		layout.NewSpacer(),
		g.clima.acContainer,
		layout.NewSpacer(),
		container.NewVBox(g.clima.time),
	)

	mainData := []*carData{
		&carDataSpacer,
		&g.mainData.consumtion,
		&carDataSpacer,
		&g.mainData.currentConsumtion,
		&carDataSpacer,
		&g.mainData.rangeVal,
		&carDataSpacer,
		&g.mainData.coolantTemperature,
	}

	g.mainDataTable = widget.NewTable(

		func() (int, int) {
			return len(mainData), 3
		},

		func() fyne.CanvasObject {
			return container.NewMax(
				widget.NewLabel("Full consumption is ...."),
				widget.NewIcon(nil),
			)
		},

		func(id widget.TableCellID, cell fyne.CanvasObject) {
			item := mainData[id.Row]

			lable := cell.(*fyne.Container).Objects[0].(*widget.Label)
			icon := cell.(*fyne.Container).Objects[1].(*widget.Icon)

			switch id.Col {
			case 0:
				if item.Icon != "" {
					iconB, err := os.ReadFile(item.Icon)
					if err != nil {
						log.Error("gui", "read icon: %v", err)
					}
					icon.SetResource(fyne.NewStaticResource(item.Icon, iconB))
					icon.Move(fyne.NewPos(40, 0))
					lable.Hide()
					icon.Show()
				} else {
					lable.SetText(item.Lable)
					lable.Show()
					icon.Hide()
				}

			case 1:
				lable.TextStyle.Bold = true
				lable.SetText(item.Value)
				lable.Alignment = fyne.TextAlignTrailing
				lable.Show()
				icon.Hide()
			case 2:
				lable.SetText(item.Unit)
				lable.Show()
				icon.Hide()
			}
		},
	)

	g.clima.overlay.Alignment = fyne.TextAlignCenter
	g.clima.overlay.TextStyle.Bold = true

	titleBox := container.NewVBox(g.title, dividerLine)

	g.viewLayout = layout.NewBorderLayout(titleBox, g.clima.climaContent, nil, nil)

	g.mainContent = container.New(g.viewLayout, titleBox, g.mainDataTable, g.clima.climaContent)

	g.menuView = container.NewBorder(nil, nil, g.menu, nil, g.mainContent)

	fyne.NewContainerWithLayout(layout.NewMaxLayout(),
		canvas.NewImageFromResource(theme.FyneLogo()), g.menuView)

	g.mainWindow.SetContent(g.menuView)
	g.mainWindow.ShowAndRun()
}

func (g *OriginalGui) MenuMain() {
	log.Info("gui", "overlay")
	g.ShowOverlay(23.5, "°C")
}

func (g *OriginalGui) MenuDashboard() {
	log.Info("gui", "overlay")
	g.ShowOverlay(25, "°C")
}

func (g *OriginalGui) MenuBC() {

}

func (g *OriginalGui) UpdateTime(t time.Time) {
	g.clima.time.SetText(t.Format("15:04"))
	g.clima.time.Refresh()
}

func (g *OriginalGui) UpdateOutdoorTemperature(value float32, unit string) {
	g.clima.outdoorTemperature.SetText(fmt.Sprintf("%.1f %s", value, unit))
	g.clima.outdoorTemperature.Refresh()
}

func (g *OriginalGui) UpdateConsumtion(consumption float32, unit string) {
	g.mainData.consumtion.Value = fmt.Sprintf("%.2f", consumption)
	g.mainData.consumtion.Unit = unit
	g.refreshLiveData()
}

func (g *OriginalGui) UpdateCurrentConsumtion(consumption float32, unit string) {
	g.mainData.currentConsumtion.Value = fmt.Sprintf("%.2f", consumption)
	g.mainData.currentConsumtion.Unit = unit
	g.refreshLiveData()
}

func (g *OriginalGui) UpdateRange(rangeV int, unit string) {
	g.mainData.rangeVal.Value = fmt.Sprintf("%d", rangeV)
	g.mainData.rangeVal.Unit = unit
	g.refreshLiveData()
}

func (g *OriginalGui) UpdateCoolantTemp(temp int, unit string) {
	g.mainData.coolantTemperature.Value = fmt.Sprintf("%d", temp)
	g.mainData.coolantTemperature.Unit = unit
	g.refreshLiveData()
}

func (g *OriginalGui) UpdateClimaFanSpeed(speed int) {
	if speed == 100 {
		g.clima.fanSpeed.SetText(fmt.Sprintf("%s", "A"))
		return
	}
	g.clima.fanSpeed.SetText(fmt.Sprintf("%d", speed))
	g.clima.fanSpeed.Refresh()
}

func (g *OriginalGui) UpdateClimaTemperature(temp int, unit string) {
	if temp == 100 {
		g.clima.temperature.SetText(fmt.Sprintf("%s", "HI"))
		return
	} else if temp == -100 {
		g.clima.temperature.SetText(fmt.Sprintf("%s", "LOW"))
		return
	}
	g.clima.temperature.SetText(fmt.Sprintf("%d°", temp))
	g.clima.temperature.Refresh()
}

func (g *OriginalGui) UpdateClimaMode(mode AcMode) {
	for i := 0; i <= AcModeAuto; i++ {
		if AcMode(i) != mode {
			acModeMap[AcMode(i)].Hide()
		} else {
			acModeMap[AcMode(i)].Show()
		}
	}
}

func (g *OriginalGui) UpdateEcoMode(eco bool) {
	if eco {
		g.clima.ecoModePlaceholder.Hide()
		g.clima.ecoMode.Show()
	} else {
		g.clima.ecoMode.Hide()
		g.clima.ecoModePlaceholder.Show()
	}
}

func (g *OriginalGui) UpdateTitle(content string) {
	g.title.SetText("\n" + content + "\n")
}

func (g *OriginalGui) refreshLiveData() {
	if g.mainDataTable != nil {
		g.mainDataTable.Refresh()
	}
}

func (g *OriginalGui) ShowOverlay(value float32, unit string) {
	g.clima.overlay.SetText(fmt.Sprintf("%.1f %s", value, unit))

	if g.clima.popup == nil {
		g.clima.popup = widget.NewPopUp(g.clima.overlay, g.mainWindow.Canvas())
		g.clima.popup.Move(fyne.NewPos(10, g.mainWindow.Canvas().Size().Height-50))

		g.clima.popup.Show()
		time.AfterFunc(GUI_OVERLAY_SHOW_DURATION*time.Second, func() { // ToDo: Restart timer, if not nil
			g.clima.popup.Hide()
			g.clima.popup = nil
		})
	}
}

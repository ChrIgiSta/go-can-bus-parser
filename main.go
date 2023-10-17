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

package main

import (
	"GMCanDecoder/can"
	"GMCanDecoder/connection/tcp"
	"GMCanDecoder/gui"
	"GMCanDecoder/utils"
	"strings"
	"sync"
	"time"

	log "github.com/ChrIgiSta/go-utils/logger"
)

type GuiType string

const (
	GuiTypeNone     GuiType = "None"
	GuiTypeTest     GuiType = "Test"
	GuiTypeOriginal GuiType = "Original"
)

const (
	// TcpHost = "192.168.43.20"
	TcpHost = "192.168.1.234"
)

var (
	testGui     *gui.TestGui
	originalGui *gui.OriginalGui

	guiType GuiType = GuiTypeOriginal
)

func main() {
	var wg sync.WaitGroup = sync.WaitGroup{}
	defer wg.Done()

	isUpdatable := utils.IsUpdatable()
	if isUpdatable {
		log.Info("main", "update application")
		err := utils.Update()
		if err != nil {
			log.Error("main", "update: %v", err)
		}
	}

	wg.Add(1)

	connLow := tcp.NewTcpClientCustomParser(TcpHost, 9000, utils.NewCanDriveParser())
	rxChannelLow, err := connLow.Connect(&wg)
	if err != nil {
		log.Error("main", "lospeed can init: %v", err)
	}
	defer connLow.Disconnect()

	connMid := tcp.NewTcpClientCustomParser(TcpHost, 9001, utils.NewCanDriveParser())
	rxChannelMid, err := connMid.Connect(&wg)
	if err != nil {
		log.Error("main", "get tcp client: %v", err)
		panic(err)
	}
	defer connMid.Disconnect()

	decoder := can.NewCanDecoder()
	evtChannel := decoder.GetEventChannel()
	go func() {
		log.Info("main", "start event catcher")
		eventDecoder(evtChannel)
	}()

	switch guiType {
	case GuiTypeTest:
		testGui = gui.NewTestGui()
	case GuiTypeOriginal:
		originalGui = gui.NewOriginalGui()
	}

	go func() {
		log.Debug("main", "start rx low")
		for true {
			canFrame, ok := <-rxChannelLow
			if !ok {
				log.Error("main", "rx channel")
				panic(ok)
			}
			err = decoder.GMLanDecoder(canFrame)
			if err != nil {
				log.Error("main", "decoding frame: %v", err)
			}
		}
	}()

	go func() {
		log.Debug("main", "start rx")
		for true {
			canFrame, ok := <-rxChannelMid
			if !ok {
				log.Error("main", "error rx channel")
				panic(ok)
			}
			if canFrame.ArbitrationID == uint32(can.EntertainmentCANAirConditioner) { // && canFrame.Data[0] == 0x21
				log.Info("main", "RX > ArbID: [0x%X] %d, \tData: %v \t[0x%X]\t%s", canFrame.ArbitrationID, canFrame.ArbitrationID, canFrame.Data, canFrame.Data, canFrame.Data)
			}
			err = decoder.EntertainmentCANDecoder(canFrame)
			if err != nil {
				log.Error("main", "decoding frame: %v", err)
			}
			// err = decoder.GMLanDecoder(canFrame)
			// if err != nil {
			// 	log.Error("main", "decoding frame: %v", err)
			// }
		}
	}()

	if testGui != nil {
		testGui.Draw()
	}
	if originalGui != nil {
		originalGui.ShowDefaultView()
	}

}

func eventDecoder(rxChannel <-chan can.CanValueMap) {
	time.Sleep(2 * time.Second) // wait init

	var displayR1C1, displayR1C2, displayR1C3, displayR1C4 string // display row 1
	var oldIgnitionVal, currentIgnition float64
	var speed float64

	for true {
		canMsg, ok := <-rxChannel
		if !ok {
			log.Error("main", "read can events")
			return
		}
		switch canMsg.CanValueDef.Name {
		case can.DisplayR1C1:
			displayR1C1 = canMsg.CanValueDef.Value.(string)
		case can.DisplayR1C2:
			displayR1C2 = canMsg.CanValueDef.Value.(string)
		case can.DisplayR1C3:
			displayR1C3 = canMsg.CanValueDef.Value.(string)
		case can.DisplayR1C4:
			displayR1C4 = canMsg.CanValueDef.Value.(string)

		// calc l/km
		case can.FullInjection:
			newIgnitionVal := canMsg.CanValueDef.Value.(float64)
			currentIgnition = newIgnitionVal - oldIgnitionVal
			oldIgnitionVal = newIgnitionVal

		case can.VehicleSpeed:
			speed = canMsg.CanValueDef.Value.(float64)

		}
		if testGui != nil {
			if strings.Contains(string(canMsg.CanValueDef.Name), "Display Row 1") {
				row1 := displayR1C1 + "," + displayR1C2 + "," + displayR1C3 + "," + displayR1C4
				testGui.UpdateValue("Display Row 1", utils.ComaSeperatedDecimalsToAscii(row1), "")
			} else if strings.Contains(string(canMsg.CanValueDef.Name), string(can.FullInjection)) {
				if speed < 5 {
					testGui.UpdateValue("Current Consumtion", currentIgnition, "l/h")
				} else {
					testGui.UpdateValue("Current Consumtion", currentIgnition/speed, "l/100km")
				}
			} else {
				testGui.UpdateValue(string(canMsg.CanValueDef.Name), canMsg.CanValueDef.Value, canMsg.CanValueDef.Unit)
			}
		}
		if originalGui != nil {
			switch canMsg.CanValueDef.Name {
			case can.ACFanSpeed:
				originalGui.UpdateClimaFanSpeed(int(canMsg.CanValueDef.Value.(float64)))
			case can.ACMode:
				switch int(canMsg.CanValueDef.Value.(float64)) {
				case can.AC_MODE_FOOD:
					originalGui.UpdateClimaMode(gui.AcModeBot)
				case can.AC_MODE_AUTO:
					originalGui.UpdateClimaMode(gui.AcModeAuto)
				case can.AC_MODE_BODY:
					originalGui.UpdateClimaMode(gui.AcModeMid)
				case can.AC_MODE_BODY_FOOD:
					originalGui.UpdateClimaMode(gui.AcModeMidBod)
				case can.AC_MODE_HEAD:
					originalGui.UpdateClimaMode(gui.AcModeTop)
				case can.AC_MODE_HEAD_BODY:
					originalGui.UpdateClimaMode(gui.AcModeTopMid)
				case can.AC_MODE_HEAD_BODY_FOOD:
					originalGui.UpdateClimaMode(gui.AcModeAll)
				case can.AC_MODE_HEAD_FOOD:
					originalGui.UpdateClimaMode(gui.AcModeTopBod)
				}

			case can.ACTemperature:
				originalGui.UpdateClimaTemperature(int(canMsg.CanValueDef.Value.(float64)), canMsg.CanValueDef.Unit)
			case can.CoolantTemperature:
				originalGui.UpdateCoolantTemp(int(canMsg.CanValueDef.Value.(float64)), canMsg.CanValueDef.Unit)
			case can.DateTime:
				t, err := utils.CanTimeStringToTime(canMsg.CanValueDef.Value.(string))
				if err != nil {
					log.Error("main", "parsing time: %v", err)
					break
				}
				originalGui.UpdateTime(t)
			case can.DisplayR1C1, can.DisplayR1C2, can.DisplayR1C3, can.DisplayR1C4:
				row1 := displayR1C1 + "," + displayR1C2 + "," + displayR1C3 + "," + displayR1C4
				originalGui.UpdateTitle(utils.ComaSeperatedDecimalsToAscii(row1))
			case can.OutdoorTemperature:
				originalGui.UpdateOutdoorTemperature(float32(canMsg.CanValueDef.Value.(float64)), canMsg.CanValueDef.Unit)
			case can.FullInjection:
				if speed < 5 {
					originalGui.UpdateCurrentConsumtion(float32(currentIgnition), "l/h")
				} else {
					originalGui.UpdateCurrentConsumtion(float32(currentIgnition), "l/100km")
				}
			case can.LeftTravelRange:
				originalGui.UpdateRange(int(canMsg.CanValueDef.Value.(float64)), canMsg.CanValueDef.Unit)
			}
		}
	}
}

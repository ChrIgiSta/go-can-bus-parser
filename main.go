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

var testGui *gui.TestGui

func main() {
	var wg sync.WaitGroup = sync.WaitGroup{}
	defer wg.Done()

	wg.Add(1)

	// conn := serial.NewSerial(serial.SERIAL_CAN_DEFAULT_PORT, serial.SERLAL_CAN_DEFAULT_BAUDRATE)
	// conn := tcp.NewTcpClient("192.168.43.20", 9001)

	// connLow := tcp.NewTcpClientCustomParser("192.168.43.20", 9000, utils.NewCanDriveParser())
	// rxChannelLow, err := connLow.Connect(&wg)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer connLow.Disconnect()

	conn := tcp.NewTcpClientCustomParser("192.168.43.20", 9001, utils.NewCanDriveParser())
	rxChannel, err := conn.Connect(&wg)
	if err != nil {
		log.Error("main", "get tcp client: %v", err)
		panic(err)
	}
	defer conn.Disconnect()

	decoder := can.NewCanDecoder()
	evtChannel := decoder.GetEventChannel()
	go func() {
		log.Info("main", "start event catcher")
		eventDecoder(evtChannel)
	}()

	testGui = gui.NewTestGui()

	// go func() {
	// 	log.Println("start rx low")
	// 	for true {
	// 		canFrame, ok := <-rxChannelLow
	// 		if !ok {
	// 			log.Panic("error rx channel")
	// 		}
	// 		err = decoder.GMLanDecoder(canFrame)
	// 		if err != nil {
	// 			log.Println("error decoding frame, ", err)
	// 		}
	// 	}
	// }()

	go func() {
		log.Debug("main", "start rx")
		for true {
			canFrame, ok := <-rxChannel
			if !ok {
				log.Error("main", "error rx channel")
				panic(ok)
			}
			if canFrame.ArbitrationID == uint32(can.EntertainmentCANAirConditioner) && canFrame.Data[0] == 0x22 {
				log.Debug("main", "RX > ArbID: [0x%X] %d, \tData: %v \t[0x%X]\t%s", canFrame.ArbitrationID, canFrame.ArbitrationID, canFrame.Data, canFrame.Data, canFrame.Data)
			}
			err = decoder.EntertainmentCANDecoder(canFrame)
			if err != nil {
				log.Error("main", "decoding frame: %v", err)
			}
			// err = decoder.GMLanDecoder(canFrame)
			// if err != nil {
			// 	log.Println("error decoding frame, ", err)
			// }
		}
	}()

	testGui.Draw()
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
	}
}

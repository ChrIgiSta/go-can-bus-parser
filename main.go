package main

import (
	"GMCanDecoder/can"
	"GMCanDecoder/connection/tcp"
	"GMCanDecoder/gui"
	"GMCanDecoder/utils"
	"log"
	"strings"
	"sync"
	"time"
)

var testGui *gui.TestGui

func main() {
	var wg sync.WaitGroup = sync.WaitGroup{}
	defer wg.Done()

	wg.Add(1)

	// conn := serial.NewSerial(serial.SERIAL_CAN_DEFAULT_PORT, serial.SERLAL_CAN_DEFAULT_BAUDRATE)
	// conn := tcp.NewTcpClient("192.168.43.20", 9001)
	conn := tcp.NewTcpClientCustomParser("192.168.43.20", 9001, utils.NewCanDriveParser())
	rxChannel, err := conn.Connect(&wg)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Disconnect()

	decoder := can.NewCanDecoder()
	evtChannel := decoder.GetEventChannel()
	go func() {
		log.Println("start event catcher")
		eventDecoder(evtChannel)
	}()

	testGui = gui.NewTestGui()

	go func() {
		log.Println("start rx")
		for true {
			canFrame, ok := <-rxChannel
			if !ok {
				log.Panic("error rx channel")
			}
			if canFrame.ArbitrationID == uint32(can.EntertainmentCANAirConditioner) && canFrame.Data[0] == 0x22 {
				log.Printf("debug: RX > ArbID: [0x%X] %d, \tData: %v \t[0x%X]\t%s", canFrame.ArbitrationID, canFrame.ArbitrationID, canFrame.Data, canFrame.Data, canFrame.Data)
			}
			err = decoder.EntertainmentCANDecoder(canFrame)
			if err != nil {
				log.Println("error decoding frame, ", err)
			}
		}
	}()

	testGui.Draw()
}

func eventDecoder(rxChannel <-chan can.CanValueMap) {
	time.Sleep(2 * time.Second) // wait init

	var displayR1C1, displayR1C2, displayR1C3, displayR1C4 string // display row 1

	for true {
		canMsg, ok := <-rxChannel
		if !ok {
			log.Println("error read can events")
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
		}
		if testGui != nil {
			if strings.Contains(string(canMsg.CanValueDef.Name), "Display Row 1") {
				row1 := displayR1C1 + "," + displayR1C2 + "," + displayR1C3 + "," + displayR1C4
				testGui.UpdateValue("Display Row 1", utils.ComaSeperatedDecimalsToAscii(row1), "")
			} else {
				testGui.UpdateValue(string(canMsg.CanValueDef.Name), canMsg.CanValueDef.Value, canMsg.CanValueDef.Unit)
			}
		}
	}
}

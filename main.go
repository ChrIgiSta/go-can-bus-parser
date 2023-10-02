package main

import (
	"GMCanDecoder/can"
	"GMCanDecoder/connection/serial"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup = sync.WaitGroup{}
	defer wg.Done()

	wg.Add(1)

	conn := serial.NewSerial(serial.SERIAL_CAN_DEFAULT_PORT, serial.SERLAL_CAN_DEFAULT_BAUDRATE)
	rxChannel, err := conn.Connect(&wg)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Disconnect()

	decoder := can.NewCanDecoder()
	evtChannel := decoder.GetEventChannel()
	go func() {
		eventDecoder(evtChannel)
	}()

	for true {
		canFrame, ok := <-rxChannel
		if !ok {
			log.Panic("error rx channel")
		}
		err = decoder.EntertainmentCANDecoder(canFrame)
		if err != nil {
			log.Println("error decoding frame, ", err)
		}
	}
}

func eventDecoder(rxChannel <-chan can.CanValueMap) {
	for true {
		canMsg, ok := <-rxChannel
		if !ok {
			log.Println("error read can events")
			return
		}
		fmt.Println(canMsg)

	}
}

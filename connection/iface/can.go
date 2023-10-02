package iface

import (
	"log"
	"sync"

	"github.com/angelodlfrtr/go-can"
	"github.com/angelodlfrtr/go-can/transports"
)

const (
	IFACE_CAN_DEFAULT_NAME = "can0"
	IFACE_CAN_BUFFER_SIZE  = 2048
)

type Iface struct {
	bus *can.Bus
}

func NewIface(iface string) *Iface {
	return &Iface{
		bus: can.NewBus(&transports.SocketCan{
			Interface: iface,
		}),
	}
}

func (i *Iface) Connect(wg *sync.WaitGroup) (<-chan *can.Frame, error) {

	var err error

	err = i.bus.Open()
	if err != nil {
		return nil, err
	}

	rxCh := make(chan *can.Frame, IFACE_CAN_BUFFER_SIZE)

	go func() {
		defer wg.Done()
		defer close(rxCh)

		for true {
			canFrame, ok := <-i.bus.ReadChan()
			if !ok {
				log.Println("error read can iface: ", err)
				return
			}
			rxCh <- canFrame
		}
	}()

	return rxCh, err
}

func (i *Iface) Disconnect() error {
	return i.bus.Close()
}

func (i *Iface) Send(message *can.Frame) error {
	return i.bus.Write(&can.Frame{})
}

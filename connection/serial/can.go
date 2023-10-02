package serial

import (
	"log"
	"sync"

	"github.com/angelodlfrtr/go-can"
	"github.com/angelodlfrtr/go-can/transports"
)

const (
	SERIAL_CAN_BUFFER_SIZE = 2048
)

const (
	SERIAL_CAN_DEFAULT_PORT     = "/dev/ttyS1234"
	SERLAL_CAN_DEFAULT_BAUDRATE = 25000
)

type Serial struct {
	bus *can.Bus
}

func NewSerial(port string, baudRate int) *Serial {
	return &Serial{
		bus: can.NewBus(&transports.USBCanAnalyzer{
			Port:     port,
			BaudRate: baudRate,
		}),
	}
}

func (s *Serial) Connect(wg *sync.WaitGroup) (<-chan *can.Frame, error) {

	var err error

	err = s.bus.Open()
	if err != nil {
		return nil, err
	}

	rxCh := make(chan *can.Frame, SERIAL_CAN_BUFFER_SIZE)

	go func() {
		defer wg.Done()
		defer close(rxCh)

		for true {
			canFrame, ok := <-s.bus.ReadChan()
			if !ok {
				log.Println("error read can iface: ", err)
				return
			}
			rxCh <- canFrame
		}
	}()

	return rxCh, err
}

func (s *Serial) Disconnect() error {
	return s.bus.Close()
}

func (s *Serial) Send(message *can.Frame) error {
	return s.bus.Write(message)
}

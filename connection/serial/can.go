package serial

import (
	"GMCanDecoder/connection"
	"log"
	"sync"

	"github.com/angelodlfrtr/go-can"
	"github.com/angelodlfrtr/go-can/transports"
	"go.bug.st/serial"
)

const (
	SERIAL_CAN_BUFFER_SIZE = 2048
)

const (
	SERIAL_CAN_DEFAULT_PORT     = "/dev/ttyS1234"
	SERLAL_CAN_DEFAULT_BAUDRATE = 25000

	SERIAL_CAN_PARITY    = serial.NoParity
	SERIAL_CAN_DATA_BITS = 8
	SERIAL_CAN_STOP_BITS = serial.OneStopBit
)

type Serial struct {
	bus             *can.Bus
	useCustomParser bool
	customParser    connection.CanFrameParser
	port            string
	baudrate        int
	com             serial.Port
}

func NewSerial(port string, baudrate int) *Serial {
	return &Serial{
		bus: can.NewBus(&transports.USBCanAnalyzer{
			Port:     port,
			BaudRate: baudrate,
		}),
		useCustomParser: false,
		customParser:    nil,
		port:            port,
		baudrate:        baudrate,
	}
}

func NewSerialCustomParser(port string, baudrate int, parser connection.CanFrameParser) *Serial {
	return &Serial{
		bus:             nil,
		useCustomParser: true,
		customParser:    parser,
		port:            port,
		baudrate:        baudrate,
	}
}

func (s *Serial) Connect(wg *sync.WaitGroup) (<-chan *can.Frame, error) {

	var err error

	if s.useCustomParser {
		return s.connectNative(wg)
	}

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

func (s *Serial) connectNative(wg *sync.WaitGroup) (<-chan *can.Frame, error) {
	var err error

	s.com, err = serial.Open(s.port, &serial.Mode{
		BaudRate: s.baudrate,
		DataBits: SERIAL_CAN_DATA_BITS,
		Parity:   SERIAL_CAN_PARITY,
		StopBits: SERIAL_CAN_STOP_BITS,
	})

	if err != nil {
		return nil, err
	}

	rxCh := make(chan *can.Frame, SERIAL_CAN_BUFFER_SIZE)

	go func() {
		defer wg.Done()
		defer close(rxCh)

		buffer := make([]byte, SERIAL_CAN_BUFFER_SIZE)
		for true {
			n, err := s.com.Read(buffer)
			if err != nil {
				log.Println("error reading serial, ", err)
				return
			}
			rxCh <- s.customParser.Unmarshal(buffer[:n])
		}
	}()

	return rxCh, err
}

func (s *Serial) Disconnect() error {
	if s.useCustomParser {
		return s.com.Close()
	}
	return s.bus.Close()
}

func (s *Serial) Send(message *can.Frame) error {
	if s.useCustomParser {
		_, err := s.com.Write(s.customParser.Marshal(message))
		return err
	}
	return s.bus.Write(message)
}

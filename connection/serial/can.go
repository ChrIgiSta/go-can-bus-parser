package serial

import (
	"errors"
	"log"
	"sync"

	"github.com/tarm/serial"
)

const (
	SERIAL_CAN_BUFFER_SIZE = 2048
	SERIAL_CAN_PARITY_BIT  = serial.ParityNone
	SERIAL_CAN_STOP_BIT    = serial.Stop1
	SERIAL_CAN_DATA_BITS   = 8
)

const (
	SERIAL_CAN_DEFAULT_PATH     = "/dev/ttyS1234"
	SERLAL_CAN_DEFAULT_BAUDRATE = 25000
)

type Serial struct {
	path     string
	baudRate int
	buffer   []byte
	port     *serial.Port
}

func NewSerial(path string, baudRate int) *Serial {
	return &Serial{
		path:     path,
		baudRate: baudRate,
		buffer:   make([]byte, SERIAL_CAN_BUFFER_SIZE),
	}
}

func (s *Serial) Connect(wg *sync.WaitGroup) (chan<- []byte, error) {

	var err error

	config := &serial.Config{
		Name:     s.path,
		Baud:     s.baudRate,
		Size:     SERIAL_CAN_DATA_BITS,
		Parity:   SERIAL_CAN_PARITY_BIT,
		StopBits: SERIAL_CAN_STOP_BIT,
	}

	s.port, err = serial.OpenPort(config)
	if err != nil {
		return nil, err
	}

	rxCh := make(chan []byte, SERIAL_CAN_BUFFER_SIZE)

	go func() {
		defer wg.Done()
		defer close(rxCh)

		for true {
			n, err := s.port.Read(s.buffer)
			if err != nil {
				log.Println("error reading from serial port: ", err)
				return
			}
			if n > 0 {
				rxCh <- s.buffer[:n]
			}
		}

	}()

	return nil, errors.New("not implemented")
}

func (s *Serial) Disconnect() error {
	return s.port.Close()
}

func (s *Serial) Send(message []byte) error {
	_, err := s.port.Write(message)
	return err
}

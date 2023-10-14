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

package serial

import (
	"GMCanDecoder/connection"
	"sync"

	log "github.com/ChrIgiSta/go-utils/logger"

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
				log.Error("serial can", "read can iface: %v", err)
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
				log.Error("serial can", "reading serial, %v", err)
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

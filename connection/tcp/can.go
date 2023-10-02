package tcp

import (
	"log"
	"sync"

	"github.com/angelodlfrtr/go-can"
	"github.com/angelodlfrtr/go-can/transports"
)

const (
	TCP_CAN_BUFFERSIZE   = 2048
	TCP_CAN_NETWORK_TYPE = "tcp"
)

const (
	TCP_CAN_DEFAULT_PORT = 9001
)

type TcpClient struct {
	bus can.Bus
}

func NewTcpClient(address string, port uint16) *TcpClient {
	return &TcpClient{
		bus: *can.NewBus(&transports.TCPCan{
			Host: address,
			Port: int(port),
		}),
	}
}

func (c *TcpClient) Connect(wg *sync.WaitGroup) (<-chan *can.Frame, error) {

	var err error

	err = c.bus.Open()
	if err != nil {
		return nil, err
	}

	rxCh := make(chan *can.Frame, TCP_CAN_BUFFERSIZE)

	go func() {
		defer wg.Done()
		defer close(rxCh)

		for true {
			canFrame, ok := <-c.bus.ReadChan()
			if !ok {
				log.Println("error read can iface: ", err)
				return
			}
			rxCh <- canFrame
		}
	}()

	return rxCh, err
}

func (c *TcpClient) Disconnect() error {
	return c.bus.Close()
}

func (c *TcpClient) Send(message *can.Frame) error {
	return c.bus.Write(message)
}

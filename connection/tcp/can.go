package tcp

import (
	"log"
	"net"
	"strconv"
	"sync"

	"GMCanDecoder/connection"

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
	address         string
	port            uint16
	tcp             net.Conn
	useCustomParser bool
	customParser    connection.CanFrameParser
	bus             can.Bus
}

func NewTcpClient(address string, port uint16) *TcpClient {
	return &TcpClient{
		bus: *can.NewBus(&transports.TCPCan{
			Host: address,
			Port: int(port),
		}),
		useCustomParser: false,
	}
}

func NewTcpClientCustomParser(address string, port uint16, parser connection.CanFrameParser) *TcpClient {
	return &TcpClient{
		address:         address,
		port:            port,
		useCustomParser: true,
		customParser:    parser,
	}
}

func (c *TcpClient) Connect(wg *sync.WaitGroup) (<-chan *can.Frame, error) {

	var err error

	if c.useCustomParser {
		return c.connectTcpNative(wg)
	}

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
				log.Println("error read can tcp: ", err)
				return
			}
			rxCh <- canFrame
		}
	}()

	return rxCh, err
}

func (c *TcpClient) connectTcpNative(wg *sync.WaitGroup) (<-chan *can.Frame, error) {
	var err error

	c.tcp, err = net.Dial(TCP_CAN_NETWORK_TYPE, c.address+":"+strconv.Itoa(int(c.port)))
	if err != nil {
		return nil, err
	}

	rxCh := make(chan *can.Frame, TCP_CAN_BUFFERSIZE)

	go func() {
		defer wg.Done()
		defer close(rxCh)

		buffer := make([]byte, TCP_CAN_BUFFERSIZE)

		for true {
			b, err := c.tcp.Read(buffer)
			if err != nil {
				log.Println("error read tcp driveMode: ", err)
			}
			canFrame := c.customParser.Unmarshal(buffer[:b-1])
			if canFrame != nil {
				rxCh <- canFrame
			}
		}
	}()

	return rxCh, err
}

func (c *TcpClient) Disconnect() error {

	if c.useCustomParser {
		return c.tcp.Close()
	}
	return c.bus.Close()
}

func (c *TcpClient) Send(message *can.Frame) error {

	if c.useCustomParser {
		_, err := c.tcp.Write(c.customParser.Marshal(message))
		return err
	}
	return c.bus.Write(message)
}

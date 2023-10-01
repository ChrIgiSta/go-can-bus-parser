package tcp

import (
	"log"
	"net"
	"strconv"
	"sync"
)

const (
	TCP_CAN_BUFFERSIZE   = 2048
	TCP_CAN_NETWORK_TYPE = "tcp"
)

const (
	TCP_CAN_DEFAULT_PORT = 9001
)

type TcpClient struct {
	port    uint16
	address string
	buffer  []byte
	conn    net.Conn
}

func NewTcpClient(address string, port uint16) *TcpClient {
	return &TcpClient{
		port:    port,
		address: address,
		buffer:  make([]byte, TCP_CAN_BUFFERSIZE),
	}
}

func (c *TcpClient) Connect(wg *sync.WaitGroup) (chan<- []byte, error) {
	var err error

	c.conn, err = net.Dial(TCP_CAN_NETWORK_TYPE, c.address+":"+strconv.Itoa(int(c.port)))
	if err != nil {
		return nil, err
	}

	rxCh := make(chan []byte, TCP_CAN_BUFFERSIZE)

	// handle connection
	go func() {
		defer wg.Done()
		defer close(rxCh)

		for true {
			n, err := c.conn.Read(c.buffer)
			if err != nil {
				log.Println("error read tcp: ", err)
				return
			}
			if n > 0 {
				rxCh <- c.buffer[:n]
			}
		}
	}()

	return rxCh, err
}

func (c *TcpClient) Disconnect() error {
	return c.Disconnect()
}

func (c *TcpClient) Send(message []byte) error {
	_, err := c.conn.Write(message)
	return err
}

package connection

import (
	"sync"

	"github.com/angelodlfrtr/go-can"
)

type Connection interface {
	Connect(wg *sync.WaitGroup) (<-chan *can.Frame, error)
	Disconnect() error
	Send(message *can.Frame) error
}

type CanFrameParser interface {
	Unmarshal(in []byte) *can.Frame
	Marshal(*can.Frame) []byte
}

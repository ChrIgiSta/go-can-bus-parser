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

package iface

import (
	"sync"

	log "github.com/ChrIgiSta/go-utils/logger"

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
				log.Error("can", "read can iface: %v", err)
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

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

package utils

import (
	"encoding/hex"
	"strconv"
	"strings"

	log "github.com/ChrIgiSta/go-utils/logger"
	"github.com/angelodlfrtr/go-can"
)

type CanDriveParser struct {
}

func NewCanDriveParser() *CanDriveParser {
	return &CanDriveParser{}
}

func (p *CanDriveParser) Unmarshal(in []byte) *can.Frame {
	split := strings.Split(string(in), ",")
	if len(split) != 4 {
		log.Error("parser", "canDrive Pck don't seem to formated properly. rx: %s", string(in))
		return nil
	}
	arbitrationID, err := strconv.ParseUint(split[0], 16, 32)
	if err != nil {
		log.Error("parser", "cannot parse canDrive Formated arbitration id: %v", err)
		return nil
	}
	data, err := hex.DecodeString(split[3])
	if err != nil {
		log.Error("parser", "cannot parse canDrive Formated data: %v", err)
		return nil
	}
	if len(data) > 8 {
		log.Error("parser", "data length higher than 8 bytes. %d", len(data))
		return nil
	}

	var dataFinal [8]byte
	copy(dataFinal[:], data)

	return &can.Frame{
		ArbitrationID: uint32(arbitrationID),
		DLC:           uint8(len(data)),
		Data:          dataFinal,
	}
}

func (p *CanDriveParser) Marshal(in *can.Frame) []byte {
	log.Error("parser", "marshal not implemented!")

	return nil
}

package utils

import (
	"encoding/hex"
	"log"
	"strconv"
	"strings"

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
		// log.Println("canDrive Pck don't seem to formated properly. rx: ", string(in))
		return nil
	}
	arbitrationID, err := strconv.ParseUint(split[0], 16, 32)
	if err != nil {
		log.Println("cannot parse canDrive Formated arbitration id: ", err)
		return nil
	}
	data, err := hex.DecodeString(split[3])
	if err != nil {
		log.Println("cannot parse canDrive Formated data: ", err)
		return nil
	}
	if len(data) > 8 {
		log.Println("data length higher than 8 bytes. ", len(data))
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
	log.Println("can drive parser: marshal not implemented!")

	return nil
}

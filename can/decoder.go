package can

import (
	"errors"
	"fmt"
	"go/token"
	"go/types"
	"regexp"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/angelodlfrtr/go-can"
)

type Decoder struct {
	lowSpeedBuffer  []can.Frame
	lowSpeedValues  []CanValueMap
	midSpeedBuffer  []can.Frame
	midSpeedValues  []CanValueMap
	highSpeedBuffer []can.Frame
	highSpeedValues []CanValueMap
}

func NewCanDecoder() *Decoder {
	return &Decoder{
		lowSpeedBuffer:  []can.Frame{},
		midSpeedBuffer:  []can.Frame{},
		highSpeedBuffer: []can.Frame{},

		lowSpeedValues: GMLanValueMapps(),
	}
}

func (d *Decoder) GetGMLanValue(name string) *CanValueMap {
	for _, val := range d.lowSpeedValues {
		if val.CanValueDef.Name == name {
			return &val
		}
	}
	return nil
}

func (d *Decoder) GetEntertainmentCANValue(name string) *CanValueMap {
	for _, val := range d.midSpeedValues {
		if val.CanValueDef.Name == name {
			return &val
		}
	}
	return nil
}

func (d *Decoder) GetHighSpeedCANValue(name string) *CanValueMap {
	for _, val := range d.highSpeedValues {
		if val.CanValueDef.Name == name {
			return &val
		}
	}
	return nil
}

func (d *Decoder) GMLanDecoder(frame *can.Frame) error {

	for i, mapping := range d.lowSpeedValues {
		if mapping.ArbitrationID == GMLanArbitrationIDs(frame.ArbitrationID) {
			condition, err := d.substituteVars(mapping.CanValueDef.Condition, frame)
			if err != nil {
				return err
			}
			fmt.Println(condition)

			fileSet := token.NewFileSet()
			tav, err := types.Eval(fileSet, nil, token.NoPos, condition)
			if err != nil {
				return err
			}
			if tav.Value.String() == "true" {
				equation, err := d.substituteVars(mapping.CanValueDef.Calculation, frame)
				if err != nil {
					return err
				}
				fmt.Println(equation)
				expression, err := govaluate.NewEvaluableExpression(equation)
				if err != nil {
					return err
				}
				result, err := expression.Evaluate(nil)
				if err != nil {
					return err
				}
				d.lowSpeedValues[i].CanValueDef.Value = result
				return nil
			}

		}
	}
	return errors.New("no mapping for incoming frame found")
}

func (d *Decoder) EntertainmentCANDecoder(frame *can.Frame) error {
	return errors.New("not implemented")
}

func (d *Decoder) HighSpeedCANDecoder(frame *can.Frame) error {
	return errors.New("not implemented")
}

func (d *Decoder) GMLanPushFrame(frame *can.Frame) error {
	if frame == nil {
		return errors.New("frame <nil>")
	}

	found := d.gmLanFindFrameByArbitrationId(frame.ArbitrationID)
	if found != nil {
		found.ArbitrationID = frame.ArbitrationID
		found.DLC = frame.DLC
		found.Data = frame.Data
	} else {
		d.lowSpeedBuffer = append(d.lowSpeedBuffer, *frame)
	}

	return d.GMLanDecoder(frame)
}

func (d *Decoder) gmLanFindFrameByArbitrationId(arbitrationID uint32) *can.Frame {
	for i, f := range d.lowSpeedBuffer {
		if f.ArbitrationID == arbitrationID {
			return &d.lowSpeedBuffer[i]
		}
	}
	return nil
}

func (d *Decoder) EntertainmentCANPushFrame(frame *can.Frame) error {
	if frame == nil {
		return errors.New("frame <nil>")
	}

	found := d.entertainmentCANFindFrameByArbitrationId(frame.ArbitrationID)
	if found != nil {
		found.ArbitrationID = frame.ArbitrationID
		found.DLC = frame.DLC
		found.Data = frame.Data
	} else {
		d.lowSpeedBuffer = append(d.lowSpeedBuffer, *frame)
	}

	return d.EntertainmentCANDecoder(frame)
}

func (d *Decoder) entertainmentCANFindFrameByArbitrationId(arbitrationID uint32) *can.Frame {
	for i, f := range d.midSpeedBuffer {
		if f.ArbitrationID == arbitrationID {
			return &d.midSpeedBuffer[i]
		}
	}
	return nil
}

func (d *Decoder) HighSpeedCANPushFrame(frame *can.Frame) error {
	if frame == nil {
		return errors.New("frame <nil>")
	}

	found := d.highSpeedCANFindFrameByArbitrationId(frame.ArbitrationID)
	if found != nil {
		found.ArbitrationID = frame.ArbitrationID
		found.DLC = frame.DLC
		found.Data = frame.Data
	} else {
		d.lowSpeedBuffer = append(d.lowSpeedBuffer, *frame)
	}

	return d.HighSpeedCANDecoder(frame)
}

func (d *Decoder) highSpeedCANFindFrameByArbitrationId(arbitrationID uint32) *can.Frame {
	for i, f := range d.highSpeedBuffer {
		if f.ArbitrationID == arbitrationID {
			return &d.highSpeedBuffer[i]
		}
	}
	return nil
}

func (d *Decoder) substituteVars(query string, frame *can.Frame) (string, error) {
	subst := query
	for i := 0; i < 8; i++ {
		subst = strings.ReplaceAll(subst, "${"+strconv.Itoa(i)+"}", strconv.Itoa(int(frame.Data[i])))
	}
	return replaceHexWithDecimal(subst), nil
}

func replaceHexWithDecimal(text string) string {
	hexPattern := `0x[0-9a-fA-F]+`

	re := regexp.MustCompile(hexPattern)

	hexMatches := re.FindAllString(text, -1)

	for _, hexMatch := range hexMatches {
		decimalValue, err := strconv.ParseInt(hexMatch[2:], 16, 64)
		if err == nil {
			text = strings.Replace(text, hexMatch, strconv.FormatInt(decimalValue, 10), -1)
		}
	}

	return text
}

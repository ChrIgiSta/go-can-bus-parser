package can

import (
	"errors"
	"go/token"
	"go/types"
	"log"
	"regexp"
	"strconv"
	"strings"

	"GMCanDecoder/utils"

	"github.com/Knetic/govaluate"
	"github.com/angelodlfrtr/go-can"
)

const EVENT_CHANNEL_BUFFER_SIZE = 1000

type Decoder struct {
	lowSpeedBuffer  []can.Frame
	lowSpeedValues  []CanValueMap
	midSpeedBuffer  []can.Frame
	midSpeedValues  []CanValueMap
	highSpeedBuffer []can.Frame
	highSpeedValues []CanValueMap

	eventChannels []chan<- CanValueMap
}

func NewCanDecoder() *Decoder {
	return &Decoder{
		lowSpeedBuffer:  []can.Frame{},
		midSpeedBuffer:  []can.Frame{},
		highSpeedBuffer: []can.Frame{},

		lowSpeedValues:  GMLanValueMapps(),
		midSpeedValues:  EntertainmentCANValueMapps(),
		highSpeedValues: HighSpeedValueMapps(),
	}
}

func (d *Decoder) GetEventChannel() <-chan CanValueMap {
	event := make(chan CanValueMap, EVENT_CHANNEL_BUFFER_SIZE)
	d.eventChannels = append(d.eventChannels, event)

	return event
}

func (d *Decoder) GetGMLanValue(name CanVars) *CanValueMap {
	for _, val := range d.lowSpeedValues {
		if val.CanValueDef.Name == name {
			return &val
		}
	}
	return nil
}

func (d *Decoder) GetEntertainmentCANValue(name CanVars) *CanValueMap {
	for _, val := range d.midSpeedValues {
		if val.CanValueDef.Name == name {
			return &val
		}
	}
	return nil
}

func (d *Decoder) GetHighSpeedCANValue(name CanVars) *CanValueMap {
	for _, val := range d.highSpeedValues {
		if val.CanValueDef.Name == name {
			return &val
		}
	}
	return nil
}

func (d *Decoder) GMLanDecoder(frame *can.Frame) error {

	for i, mapping := range d.lowSpeedValues {
		if mapping.ArbitrationID == frame.ArbitrationID {
			err := d.processFrame(&d.lowSpeedValues[i], frame)
			if err != nil {
				return err
			} else {
				continue
			}
		}
	}

	return nil
}

func (d *Decoder) EntertainmentCANDecoder(frame *can.Frame) error {
	for i, mapping := range d.midSpeedValues {
		if mapping.ArbitrationID == frame.ArbitrationID {
			err := d.processFrame(&d.midSpeedValues[i], frame)
			if err != nil {
				return err
			} else {
				continue
			}
		}
	}

	return nil
}

func (d *Decoder) HighSpeedCANDecoder(frame *can.Frame) error {
	for i, mapping := range d.highSpeedValues {
		if mapping.ArbitrationID == frame.ArbitrationID {
			err := d.processFrame(&d.highSpeedValues[i], frame)
			if err != nil {
				return err
			} else {
				continue
			}
		}
	}

	return nil
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

func (d *Decoder) processFrame(mapping *CanValueMap, frame *can.Frame) error {
	condition, err := d.substituteVars(mapping.CanValueDef.Condition, frame)
	if err != nil {
		return err
	}

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
		formatedString := false
		splittedEquation := strings.Split(equation, ";")
		if len(splittedEquation) != 1 {
			formatedString = true
		}
		if !formatedString {
			expression, err := govaluate.NewEvaluableExpression(equation)
			if err != nil {
				return err
			}
			result, err := expression.Evaluate(nil)
			if err != nil {
				return err
			}
			mapping.CanValueDef.Value = result
			if mapping.TriggerEvent {
				d.processEvent(mapping)
			}
			return nil
		} else {
			output := ""
			for sIndx, split := range splittedEquation {
				expression, err := govaluate.NewEvaluableExpression(split)
				if err != nil {
					return err
				}
				result, err := expression.Evaluate(nil)
				if err != nil {
					return err
				}
				output += utils.InterfaceToString(result)
				if len(mapping.CanValueDef.FormatSeperators) > sIndx {
					output += mapping.CanValueDef.FormatSeperators[sIndx]
				}
			}
			mapping.CanValueDef.Value = output
			if mapping.TriggerEvent {
				d.processEvent(mapping)
			}
			return nil
		}
	}

	return nil
}

func (d *Decoder) processEvent(canVal *CanValueMap) {
	if canVal != nil {
		for _, evtCh := range d.eventChannels {
			evtCh <- *canVal
		}
	} else {
		log.Println("error event: value <nil>")
	}
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

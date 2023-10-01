package can

import (
	"testing"

	"github.com/angelodlfrtr/go-can"
)

func TestGMLanDecoder(t *testing.T) {

	d := NewCanDecoder()

	fr := can.Frame{}
	fr.ArbitrationID = uint32(GMLanOutdoorTemperatureSensor)
	fr.Data = [8]uint8{0x46, 0x01, 0x83}

	err := d.GMLanDecoder(&fr)
	if err != nil {
		t.Error(err)
	}

	outTemp := d.GetGMLanValue(OutdoorTemperature)
	if outTemp.CanValueDef.Value.(float64) != 25.5 {
		t.Error("decode temp failed: ", outTemp.CanValueDef.Value)
	}
}

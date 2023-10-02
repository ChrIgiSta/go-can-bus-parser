package can

import (
	"testing"

	"github.com/angelodlfrtr/go-can"
)

func TestGMLanDecoder(t *testing.T) {

	d := NewCanDecoder()

	fr := can.Frame{}

	// Outdoor Temp
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

	t.Logf("%s is %v%s", outTemp.CanValueDef.Name, outTemp.CanValueDef.Value, outTemp.CanValueDef.Unit)

	// DateTime
	fr.ArbitrationID = uint32(GMLanDate)
	fr.Data = [8]uint8{0xff, 0x17, 0x05, 0x0b, 0x0c, 0x03, 0x00}
	err = d.GMLanDecoder(&fr)
	if err != nil {
		t.Error(err)
	}
	dateTime := d.GetGMLanValue(DateTime)
	if dateTime.CanValueDef.Value.(string) != "23-5-11T16:3:0" {
		t.Error("decode date failed: ", dateTime.CanValueDef.Value)
	}

	t.Logf("%s is %v%s", dateTime.CanValueDef.Name, dateTime.CanValueDef.Value, dateTime.CanValueDef.Unit)
}

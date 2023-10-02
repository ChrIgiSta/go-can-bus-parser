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

	printCanValueInfos(t, outTemp)

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
	printCanValueInfos(t, dateTime)

	// Speeds (RPM / km/h)
	fr.ArbitrationID = uint32(GMLanSpeeds)
	fr.Data = [8]uint8{0x23, 0x20, 0x98, 0x00, 0x04, 0xe5, 0x00, 0x00}
	err = d.GMLanDecoder(&fr)
	if err != nil {
		t.Error(err)
	}
	engineSpeed := d.GetGMLanValue(EngineSpeedRPM)
	vehicleSpeed := d.GetGMLanValue(VehicleSpeed)
	// if dateTime.CanValueDef.Value.(string) != "23-5-11T16:3:0" {
	// 	t.Error("decode date failed: ", dateTime.CanValueDef.Value)
	// }

	printCanValueInfos(t, engineSpeed)
	printCanValueInfos(t, vehicleSpeed)
}

func printCanValueInfos(t *testing.T, canValue *CanValueMap) {
	t.Logf("%s is %v%s", canValue.CanValueDef.Name, canValue.CanValueDef.Value, canValue.CanValueDef.Unit)
}

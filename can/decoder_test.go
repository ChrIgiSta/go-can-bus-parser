package can

import (
	"GMCanDecoder/utils"
	"fmt"
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
	if dateTime.CanValueDef.Value.(string) != "23-5-11T16:3:0" {
		t.Error("decode date failed: ", dateTime.CanValueDef.Value)
	}

	printCanValueInfos(t, engineSpeed)
	printCanValueInfos(t, vehicleSpeed)
}

func TestMidspeedDecoder(t *testing.T) {
	displayData := [][8]byte{
		{0x10, 0x42, 0x40, 0x00, 0x3F, 0x03, 0x10, 0x13},
		{0x21, 0x00, 0x1B, 0x00, 0x5B, 0x00, 0x66, 0x00},
		{0x22, 0x53, 0x00, 0x5F, 0x00, 0x67, 0x00, 0x6D},
		{0x23, 0x00, 0x4E, 0x00, 0x6F, 0x00, 0x20, 0x00},
		{0x24, 0x53, 0x00, 0x6F, 0x00, 0x75, 0x00, 0x72},
		{0x25, 0x00, 0x63, 0x00, 0x65, 0x00, 0x20, 0x00},
		{0x26, 0x20, 0x00, 0x20, 0x00, 0x21, 0x00, 0x21},
		{0x27, 0x00, 0x20, 0x00, 0x20, 0x00, 0x20, 0x00},
		{0x28, 0x64, 0x00, 0x6D, 0x00, 0x20, 0x00, 0x4D},
		{0x29, 0x00, 0x41, 0x00, 0x4E, 0x00, 0x21, 0x00},
	}

	d := NewCanDecoder()

	// Display
	fr := can.Frame{}
	fr.ArbitrationID = uint32(EntertainmentCANDisplayData)

	for _, data := range displayData {

		fr.Data = data
		err := d.EntertainmentCANDecoder(&fr)
		if err != nil {
			t.Error(err)
		}
	}
	r1c1 := d.GetEntertainmentCANValue(DisplayR1C1)
	r1c2 := d.GetEntertainmentCANValue(DisplayR1C2)
	r1c3 := d.GetEntertainmentCANValue(DisplayR1C3)
	r1c4 := d.GetEntertainmentCANValue(DisplayR1C4)

	row1Display := utils.ComaSeperatedDecimalsToAscii(
		r1c1.CanValueDef.Value.(string) + "," +
			r1c2.CanValueDef.Value.(string) + "," +
			r1c3.CanValueDef.Value.(string) + "," +
			r1c4.CanValueDef.Value.(string))

	if row1Display != "No Source   !!" {
		t.Error("row 1 display")
	}
	fmt.Println(row1Display)

}

func printCanValueInfos(t *testing.T, canValue *CanValueMap) {
	t.Logf("%s is %v%s", canValue.CanValueDef.Name, canValue.CanValueDef.Value, canValue.CanValueDef.Unit)
}

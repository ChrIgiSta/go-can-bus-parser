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

package can

import (
	"GMCanDecoder/utils"
	"fmt"
	"testing"

	"github.com/angelodlfrtr/go-can"
)

func TestGMLanDecoder(t *testing.T) {

	var err error

	d := NewCanDecoder()

	fr := can.Frame{}

	// // Outdoor Temp
	// fr.ArbitrationID = uint32(GMLanOutdoorTemperatureSensor)
	// fr.Data = [8]uint8{0x46, 0x01, 0x83}

	// err = d.GMLanDecoder(&fr)
	// if err != nil {
	// 	t.Error(err)
	// }

	// outTemp := d.GetGMLanValue(OutdoorTemperature)
	// if outTemp.CanValueDef.Value.(float64) != 25.5 {
	// 	t.Error("decode temp failed: ", outTemp.CanValueDef.Value)
	// }

	// printCanValueInfos(t, outTemp)

	// // DateTime
	// fr.ArbitrationID = uint32(GMLanDate)
	// fr.Data = [8]uint8{0xff, 0x17, 0x05, 0x0b, 0x0c, 0x03, 0x00}
	// err = d.GMLanDecoder(&fr)
	// if err != nil {
	// 	t.Error(err)
	// }
	// dateTime := d.GetGMLanValue(DateTime)
	// if dateTime.CanValueDef.Value.(string) != "23-5-11T16:3:0" {
	// 	t.Error("decode date failed: ", dateTime.CanValueDef.Value)
	// }
	// printCanValueInfos(t, dateTime)

	// Speeds (RPM / km/h)
	fr.ArbitrationID = uint32(GMLanEngineSpeedRPM)
	fr.Data = [8]uint8{0x13, 0x0c, 0xf3, 0x00, 0x04, 0xe5, 0x00, 0x00}
	err = d.GMLanDecoder(&fr)
	if err != nil {
		t.Error(err)
	}
	engineSpeed := d.GetGMLanValue(EngineSpeedRPM)
	vehicleSpeed := d.GetGMLanValue(VehicleSpeed)
	engineState := d.GetGMLanValue(EngineRunningState)

	printCanValueInfos(t, engineSpeed)
	printCanValueInfos(t, vehicleSpeed)
	printCanValueInfos(t, engineState)
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

	// DateTime
	dateData := [8]byte{0x46, 0x01, 0x17, 0x0a, 0x5d, 0x12, 0x27, 0xff}
	err := d.EntertainmentCANDecoder(&can.Frame{
		ArbitrationID: uint32(EntertainmentCANDate),
		DLC:           8,
		Data:          dateData,
	})
	if err != nil {
		t.Error(err)
	}
	dt := d.GetEntertainmentCANValue(DateTime)
	if dt.CanValueDef.Value != "23-10-11T20:18:39" {
		t.Error("Midspeed DateTime")
	}
	fmt.Println(dt.CanValueDef.Value)

	// // ClimaData
	// acData := []byte{0x23, 0xe0, 0x50, 0x00, 0x37, 0x20, 0x26, 0x02}

}

func printCanValueInfos(t *testing.T, canValue *CanValueMap) {
	t.Logf("%s is %v%s", canValue.CanValueDef.Name, canValue.CanValueDef.Value, canValue.CanValueDef.Unit)
}

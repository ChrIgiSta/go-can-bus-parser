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

import "fmt"

type CanVars string

const (
	ACTemperature      CanVars = "AC Temperature"
	ACMode             CanVars = "AC Mode"
	ACFanSpeed         CanVars = "AC Fan Speed"
	BatteryVoltage     CanVars = "Battery Voltage"        // tested
	BusWakeup          CanVars = "CAN-Bus Wakeup"         // tested
	BreakState         CanVars = "Break State"            // tested
	DateTime           CanVars = "Date"                   // tested
	EngineSpeedRPM     CanVars = "Engine RPM"             // tested
	FullInjection      CanVars = "Full Injection"         // partialy tested
	FullLevel          CanVars = "Full Level"             // tested
	LedBrightness      CanVars = "Led Brightness"         // tested
	Milage             CanVars = "Milage"                 // tested
	OutdoorTemperature CanVars = "Output Temperature"     // tested
	VehicleSpeed       CanVars = "Speed"                  // tested
	WeelKey            CanVars = "Weel Remote Key"        // tested
	DisplayR1C1        CanVars = "Display Row 1 Column 1" // tested
	DisplayR1C2        CanVars = "Display Row 1 Column 2" // tested
	DisplayR1C3        CanVars = "Display Row 1 Column 3" // tested
	DisplayR1C4        CanVars = "Display Row 1 Column 4" // tested
	LightSwitch        CanVars = "Light Switch"           // tested
	LightLeveler       CanVars = "Light Leveler"          // tested
	LightBack          CanVars = "Light Back"             // tested
	DoorState          CanVars = "Door State"             // tested
	EngineRunningState CanVars = "Engine State"           // tested
	TurnLights         CanVars = "Turn Lights"            // dont work
	CoolantTemperature CanVars = "Coolant Temperature"    // tested
	TPMS               CanVars = "Tire Pressure Monitoring System"
	CruseControl       CanVars = "Cruse Control" // tested
	SystemTime         CanVars = "System Time"
	DisplayTemperature CanVars = "Display Temperature"
	SensorTemperature  CanVars = "Outdoor Sensor Temperature"
	LeftTravelRange    CanVars = "Range" // tested
	RangeWarning       CanVars = "Range Warning"
)

// low speed
const (
	BREAK_PRESSED = 0x40
	BREAK_OPEN    = 0x00

	WEEL_KEY_SEEK_UP      = 0x10
	WEEL_KEY_SEEK_DOWN    = 0x20
	WEEL_KEY_SEEK_PRESSED = 0x30
	WEEL_KEY_MUTE_PRESSED = 0x40
	WEEL_KEY_MODE_PRESSED = 0x50
	WEEL_KEY_UP_PRESSED   = 0x04
	WEEL_KEY_DOWN_PRESSED = 0x05
	WEEL_KEY_VOLUME_UP    = 0x01
	WEEL_KEY_VOLUME_DOWN  = 0x02

	DRIVING_LIGHT_OFF      = 0x00
	DRIVING_LIGHT_PARKING  = 0x40
	DRIVING_LIGHT_LOW_BEAM = 0xc0
	DRIVING_LIGHT_REVERSE  = 0x01

	LIGHT_LEVELER_HIGH_BEAM_BIT = 0x40
	LIGHT_LEVELER_FOG_FRONT_BIT = 0x20

	LIGHT_BACK_FOG       = 0x80
	LIGHT_BACK_HANDBRAKE = 0x01
	LIGHT_BACK_WISHWATER = 0x03 // maybe also bulbs?

	DOOR_STATE_FRONT_LEFT_OPEN   = 0x40
	DOOR_STATE_FRONT_RIGHT_OPPEN = 0x10
	DOOR_STATE_TRUNK_OPEN        = 0x04
	// guess 5 door car
	DOOR_STATE_BACK_RIGHT_OPEN = 0x01
	DOOR_STATE_BACK_LEFT_OPEN  = 0x08

	ENGINE_OFF             = 0x00
	ENGINE_IGNITION_ON     = 0x03
	ENGINE_STARTER_RUNNING = 0x43 // -> 4 starter                    -> 3 ignition
	ENGINE_RUNNING         = 0x13 // -> 1 engine running             -> 3 ignition
	ENGINE_RUNNING_DRIVING = 0x23 // -> 2 engine running and driving -> 3 ignition

	CRUSE_CONTROLL_ON = 0x06 // 0x04 off

	RANGE_WARNING_OFF = 0x03
	RANGE_WARNING_ON  = 0x00

	AC_MODE_AUTO           = 89
	AC_MODE_HEAD           = 83
	AC_MODE_BODY           = 85
	AC_MODE_FOOD           = 87
	AC_MODE_HEAD_BODY      = 84
	AC_MODE_HEAD_FOOD      = 88
	AC_MODE_BODY_FOOD      = 86
	AC_MODE_HEAD_BODY_FOOD = 82
)

// mid speed
const (
	SOME_MID_SPEED = 0x00
)

// high speed
const (
	SOME_HIGH_SPEED = 0x00
)

const (
	FULL_CAPACITY_L = 52
)

// Opel's LowSpeed SW-CAN
type GMLanArbitrationIDs uint32

const (
	GMLanBusWakeup          GMLanArbitrationIDs = 0x100 // tested
	GMLanEngineSpeedRPM     GMLanArbitrationIDs = 0x108 // tested
	GMLanFullInjection      GMLanArbitrationIDs = 0x130 // not valid as it is, seems to be a counter (ml/s or something like that)
	GMLanCoolant            GMLanArbitrationIDs = 0x145 // tested
	GMLanCruseControl       GMLanArbitrationIDs = 0x145 // tested
	GMLanWeelRemoteControll GMLanArbitrationIDs = 0x175 // tested
	GMLanMilage             GMLanArbitrationIDs = 0x190 // tested
	GMLanDoorState          GMLanArbitrationIDs = 0x230 // tested
	GMLanLedBrightness      GMLanArbitrationIDs = 0x235 // tested
	GMLanLightSwitch        GMLanArbitrationIDs = 0x305 // tested
	GMLanLightLevler        GMLanArbitrationIDs = 0x350 // tested
	GMLanClutchBreak        GMLanArbitrationIDs = 0x360 // tested
	GMLanLightBack          GMLanArbitrationIDs = 0x370 // tested
	GMLanFullLevel          GMLanArbitrationIDs = 0x375 // tested
	GMLanSysTime            GMLanArbitrationIDs = 0x440
	GMLanOutputTemperature  GMLanArbitrationIDs = 0x445 // tested
	GMLanBatteryVoltage     GMLanArbitrationIDs = 0x500 // tested
	GMLanTPMS               GMLanArbitrationIDs = 0x530 // unknown
)

// Opel's MidSpeed CAN
type EntertainmentCANArbitrationIDs uint32

const (
	EntertainmentCANDate                 GMLanArbitrationIDs            = 0x180 // tested
	EntertainmentCANDistance             GMLanArbitrationIDs            = 0x188
	EntertainmentCANRadioButtons         GMLanArbitrationIDs            = 0x201
	EntertainmentCANSteeringWheelButtons GMLanArbitrationIDs            = 0x206
	EntertainmentCANACKnobs              GMLanArbitrationIDs            = 0x208
	EntertainmentCANEngineMotion         GMLanArbitrationIDs            = 0x4e8
	EntertainmentCANEngineTemperature    GMLanArbitrationIDs            = 0x4ec
	EntertainmentCANFullInjection        GMLanArbitrationIDs            = 0x4ed
	EntertainmentCANRange                GMLanArbitrationIDs            = 0x4ee
	EntertainmentCANDisplayTemperature   GMLanArbitrationIDs            = 0x682 // tested
	EntertainmentCANSensorTemperature    GMLanArbitrationIDs            = 0x683 // tested
	EntertainmentCANTPMSPressure         GMLanArbitrationIDs            = 0x684
	EntertainmentCANTPMSBattery          GMLanArbitrationIDs            = 0x685
	EntertainmentCANDisplayData          EntertainmentCANArbitrationIDs = 0x6c1 // partialy tested
	EntertainmentCANAirConditioner       EntertainmentCANArbitrationIDs = 0x6c8 // -> knöpfe?
	// EntertainmentCANAirConditioner GMLanArbitrationIDs = 0x206 //
)

type HighSpeedCANArbitrationIDs uint32

const ()

type CanValueDef struct {
	Unit             string
	Calculation      string
	FormatSeperators []string // if using ; in calc
	Condition        string
	Name             CanVars
	Value            interface{}
}

type CanValueMap struct {
	CanValueDef   CanValueDef
	ArbitrationID uint32
	TriggerEvent  bool
}

// Opel Astra H OPC 2006
func GMLanValueMapps() []CanValueMap {
	return []CanValueMap{
		{
			ArbitrationID: uint32(GMLanWeelRemoteControll),
			CanValueDef: CanValueDef{
				Unit:        "Key Action",
				Calculation: "${5}",
				Condition:   "${2} == 0x00 && ${3} == 0x00",
				Name:        WeelKey,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanWeelRemoteControll),
			CanValueDef: CanValueDef{
				Unit:        "Turn Lights",
				Calculation: "${4}",
				Condition:   "${2} == 0x00 && ${3} == 0x00",
				Name:        TurnLights,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanBusWakeup),
			CanValueDef: CanValueDef{
				Unit:        "Bus Wakeup",
				Calculation: "1",
				Condition:   "1 == 1",
				Name:        BusWakeup,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanEngineSpeedRPM),
			CanValueDef: CanValueDef{
				Unit:        "RPM",
				Calculation: "(${1}*256 + ${2})/4",
				Condition:   "1 == 1", // engine is running -> ${0} == 0x13
				Name:        EngineSpeedRPM,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanEngineSpeedRPM),
			CanValueDef: CanValueDef{
				Unit:        "km/h",
				Calculation: "(${4}*256 + ${5}) / 128",
				Condition:   "1 == 1", // engine is running -> ${0} == 0x13
				Name:        VehicleSpeed,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanEngineSpeedRPM),
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${0}",
				Condition:   "1 == 1",
				Name:        EngineRunningState,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanMilage),
			CanValueDef: CanValueDef{
				Unit:        "km",
				Calculation: "(${2}*65536 + ${3}*256 +${4}) / 64", // 00 00 98 92 c0 00 21
				Condition:   "1 == 1",                             // ${0} == 0x23"
				Name:        Milage,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanClutchBreak),
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${2}",
				Condition:   "${0} == 0x00 && ${1} == 0x00",
				Name:        BreakState,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanBatteryVoltage),
			CanValueDef: CanValueDef{
				Unit:        "V",
				Calculation: "${1} / 8",
				Condition:   "1 == 1",
				Name:        BatteryVoltage,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanLedBrightness),
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${1}",
				Condition:   "${0} == 0x00",
				Name:        LedBrightness,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanFullLevel),
			CanValueDef: CanValueDef{
				Unit:        "l",
				Calculation: fmt.Sprintf("%d - (${1} * %d / 0xff)", FULL_CAPACITY_L, FULL_CAPACITY_L), // 256 / 2.56 -> 100% -> OPC = 52l
				Condition:   "${0} == 0x00",
				Name:        FullLevel,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanFullInjection),
			CanValueDef: CanValueDef{
				Unit:        "1/ l/h",
				Calculation: "(${1} * 256 + ${2}) / 4.5", // dT= 0.5s, 00 23 f1 00 e5 01 0e -> 00 24 11 00 e5 01 0e | 32 * 0.03054 = 0.97 -> 1.94 ml/s -> 6.984 l/h ?
				Condition:   "${0} == 0x00",
				Name:        FullInjection,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanLightSwitch),
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${2}",
				Condition:   "${0} == 0x00 && ${1} == 0x00",
				Name:        LightSwitch,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanLightLevler),
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${0}",
				Condition:   "1 == 1",
				Name:        LightLeveler,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanLightBack),
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${1}",
				Condition:   "${0} == 0",
				Name:        LightBack,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanDoorState),
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${2}",
				Condition:   "${0} == 0 && ${1} == 0",
				Name:        DoorState,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanCoolant),
			CanValueDef: CanValueDef{
				Unit:        "°C",
				Calculation: "${3} - 40",
				Condition:   "${5} == 0x04 && ${6} == 0",
				Name:        CoolantTemperature,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanOutputTemperature),
			CanValueDef: CanValueDef{
				Unit:        "°C",
				Calculation: "${1} / 2 - 40",
				Condition:   "${0} == 0x00",
				Name:        OutdoorTemperature,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanTPMS),
			CanValueDef: CanValueDef{
				Unit:        "bar",
				Calculation: "${2}/25;${3}/25;${4}/25;${5}/25",
				Condition:   "1 == 1",
				Name:        TPMS,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanCruseControl),
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${5}",
				Condition:   "1 == 1",
				Name:        CruseControl,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanSysTime),
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${0};${1};${2}",
				Condition:   "1 == 1",
				Name:        SystemTime,
			},
			TriggerEvent: true,
		},
	}
}

func EntertainmentCANValueMapps() []CanValueMap {
	return []CanValueMap{
		{
			ArbitrationID: uint32(EntertainmentCANDisplayTemperature),
			CanValueDef: CanValueDef{
				Unit:        "°C",
				Calculation: "${2} / 2 - 40", // as normal calculation
				Condition:   "${0} == 0x46 && ${1} == 0x01",
				Name:        DisplayTemperature,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANSensorTemperature),
			CanValueDef: CanValueDef{
				Unit:        "°C",
				Calculation: "${2} / 2 - 40", // as normal calculation
				Condition:   "${0} == 0x46 && ${1} == 0x01",
				Name:        SensorTemperature,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANDate),
			CanValueDef: CanValueDef{
				Unit:             "",
				Calculation:      "${2};${3};${4}>>3;((${4}&0x07)<<2)+(${5}>>6);${5}&0x3f;${6}",
				FormatSeperators: []string{"-", "-", "T", ":", ":"}, // 23-01-24T18:21:3
				Condition:        "1 == 1",                          // ever
				Name:             DateTime,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANAirConditioner), // works
			CanValueDef: CanValueDef{
				Unit:        "°C",
				Calculation: "(((${3} & 0x03) * 10) + (${5} & 0x3f))-48", // oberstes bit (0x80) -> low or high
				Condition:   "${0} == 0x22 && ${1} == 0x03",              // ${0} == 0x22
				Name:        ACTemperature,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANAirConditioner), // don't prove
			CanValueDef: CanValueDef{
				Unit:        "°C",
				Calculation: "100", // Hi
				Condition:   "${0} == 0x22 && ${1} == 0x48",
				Name:        ACTemperature,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANAirConditioner), // don't prove
			CanValueDef: CanValueDef{
				Unit:        "°C",
				Calculation: "-100", // Low
				Condition:   "${0} == 0x22 && ${1} == 0x4c",
				Name:        ACTemperature,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANAirConditioner), // works
			CanValueDef: CanValueDef{
				Unit:        "rpm",
				Calculation: "${3} & 0x0f",
				Condition:   "${0} == 0x22 && ${1} == 0x50",
				Name:        ACFanSpeed,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANAirConditioner), // don't prove
			CanValueDef: CanValueDef{
				Unit:        "rpm",
				Calculation: "100",
				Condition:   "${0} == 0x23 && ${1} == 0x26",
				Name:        ACFanSpeed,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANAirConditioner), // works
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${2}",
				Condition:   "${0} == 0x21 && ${1} == 0xe0",
				Name:        ACMode,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANAirConditioner),
			CanValueDef: CanValueDef{
				Unit:        "l",
				Calculation: "52 - ${2} * 2",
				Condition:   "0 == 0",              // ? 0x46
				Name:        "Full Level Midspeed", // ToDO
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANDistance),
			CanValueDef: CanValueDef{
				Unit:        "km",
				Calculation: "(${2} * 256 + ${3}) * 1.5748",
				Condition:   "${0} == 0x46",
				Name:        "Distance??", // ToDO
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANEngineMotion),
			CanValueDef: CanValueDef{
				Unit:        "rpm",
				Calculation: "(${2} * 256 + ${3}) / 4",
				Condition:   "${0} == 0x46",
				Name:        EngineSpeedRPM,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANEngineMotion),
			CanValueDef: CanValueDef{
				Unit:        "km/h",
				Calculation: "${4}",
				Condition:   "${0} == 0x46",
				Name:        VehicleSpeed,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANEngineTemperature),
			CanValueDef: CanValueDef{
				Unit:        "°C",
				Calculation: "${2} - 40",
				Condition:   "${0} == 0x46",
				Name:        CoolantTemperature,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANFullInjection),
			CanValueDef: CanValueDef{
				Unit:        "num ignitions 0xffff is max",
				Calculation: "${2} * 256 + ${3}",
				Condition:   "${0} == 0x46",
				Name:        "Ign mid", // ToDO (upcounting value -> find out, how much is the value)  highly propable means l since started -> posibilities to calc l/h
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANRange),
			CanValueDef: CanValueDef{
				Unit:        "km",
				Calculation: "(${2} * 256 + ${3}) * 0.5",
				Condition:   "${0} == 0x46",
				Name:        LeftTravelRange,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANRange),
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${1}",
				Condition:   "${0} == 0x46",
				Name:        RangeWarning,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANDisplayData),
			CanValueDef: CanValueDef{
				Unit:             "",
				Calculation:      "${2};${4};${6}",
				FormatSeperators: []string{",", ","},
				Condition:        "${0} == 0x23",
				Name:             DisplayR1C1,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANDisplayData),
			CanValueDef: CanValueDef{
				Unit:             "",
				Calculation:      "${1};${3};${5};${7}",
				Condition:        "${0} == 0x24",
				FormatSeperators: []string{",", ",", ","},
				Name:             DisplayR1C2,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANDisplayData),
			CanValueDef: CanValueDef{
				Unit:             "",
				Calculation:      "${2};${4};${6}",
				FormatSeperators: []string{",", ","},
				Condition:        "${0} == 0x25",
				Name:             DisplayR1C3,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANDisplayData),
			CanValueDef: CanValueDef{
				Unit:             "",
				Calculation:      "${1};${3};${5};${7}",
				FormatSeperators: []string{",", ",", ","},
				Condition:        "${0} == 0x26",
				Name:             DisplayR1C4,
			},
			TriggerEvent: true,
		},
	}
}

func HighSpeedValueMapps() []CanValueMap {
	return []CanValueMap{
		{},
	}
}

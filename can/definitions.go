package can

type CanVars string

const (
	ACTemperature      CanVars = "AC Temperature"
	ACMode             CanVars = "AC Mode"
	ACFanSpeed         CanVars = "AC Fan Speed"
	BatteryVoltage     CanVars = "Battery Voltage" // tested
	BusWakeup          CanVars = "CAN-Bus Wakeup"  // tested
	BreakState         CanVars = "Break State"     // tested
	DateTime           CanVars = "Date"
	EngineSpeedRPM     CanVars = "Engine RPM" // tested
	FullInjection      CanVars = "Full Injection"
	FullLevel          CanVars = "Full Level"     // it is highly possible the level, but 0 - 255
	LedBrightness      CanVars = "Led Brightness" // tested
	Milage             CanVars = "Milage"
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
)

// mid speed
const (
	SOME_MID_SPEED = 0x00
)

// high speed
const (
	SOME_HIGH_SPEED = 0x00
)

// Opel's LowSpeed SW-CAN
type GMLanArbitrationIDs uint32

const (
	GMLanBusWakeup          GMLanArbitrationIDs = 0x100 // tested
	GMLanEngineSpeedRPM     GMLanArbitrationIDs = 0x108 // tested // -->> maybe also cool water temperature ..?
	GMLanFullInjection      GMLanArbitrationIDs = 0x130 // not valid
	GMLanCoolant            GMLanArbitrationIDs = 0x145 // tested
	GMLanCruseControl       GMLanArbitrationIDs = 0x145
	GMLanWeelRemoteControll GMLanArbitrationIDs = 0x175 // tested
	GMLanMilage             GMLanArbitrationIDs = 0x190 // not valid
	GMLanDoorState          GMLanArbitrationIDs = 0x230 // tested
	GMLanLedBrightness      GMLanArbitrationIDs = 0x235 // tested
	GMLanLightSwitch        GMLanArbitrationIDs = 0x305 // tested
	GMLanLightLevler        GMLanArbitrationIDs = 0x350 // tested
	GMLanClutchBreak        GMLanArbitrationIDs = 0x360 // tested
	GMLanLightBack          GMLanArbitrationIDs = 0x370 // tested
	GMLanFullLevel          GMLanArbitrationIDs = 0x375 // not valid
	GMLanSysTime            GMLanArbitrationIDs = 0x440
	GMLanOutputTemperature  GMLanArbitrationIDs = 0x445
	GMLanBatteryVoltage     GMLanArbitrationIDs = 0x500 // tested)
	GMLanTPMS               GMLanArbitrationIDs = 0x530
)

// Opel's MidSpeed CAN
type EntertainmentCANArbitrationIDs uint32

const (
	EntertainmentCANDisplayData    EntertainmentCANArbitrationIDs = 0x6c1 // partialy tested
	EntertainmentCANAirConditioner EntertainmentCANArbitrationIDs = 0x6c8 // tested
	// EntertainmentCANAirConditioner GMLanArbitrationIDs = 0x206 //
	EntertainmentCANDate GMLanArbitrationIDs = 0x180 // don't work after python -> possibly, this was mid speed can -> not possibly, it is
	// EntertainmentCANAirConditioner GMLanArbitrationIDs = 0x683 // don't work after python -> possibly, this was mid speed can -> not possibly, it is
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
				Calculation: "${2}*65536 + ${3}*256 +${4}",
				Condition:   "1 == 1", // ${0} == 0x23"
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
				Calculation: "${1} * 0.21", // 256 / 2.56 -> 100% -> OPC = 52l
				Condition:   "${0} == 0x00",
				Name:        FullLevel,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanFullInjection),
			CanValueDef: CanValueDef{
				Unit:        "ml",
				Calculation: "(${1} * 256 + ${2}) * 0.03054",
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
				Unit:        "km/h",
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
		// {
		// 	ArbitrationID: uint32(EntertainmentCANOutdoorTemperatureSensor),
		// 	CanValueDef: CanValueDef{
		// 		Unit:        "°C",
		// 		Calculation: "${2} / 2 - 40", // as normal calculation
		// 		Condition:   "${0} == 0x46 && ${1} == 0x01",
		// 		Name:        OutdoorTemperature,
		// 	},
		// 	TriggerEvent: true,
		// },
		{
			ArbitrationID: uint32(EntertainmentCANDate),
			CanValueDef: CanValueDef{
				Unit:             "",
				Calculation:      "${1};${2};${3};${4}+4;${5};${6}", // as formated string (;) // Year(2) 	Month(3) 	5B:Day, 3b:Hour 	2b:Hour, 6b:Minute 	Second // see https://github.com/JJToB/Car-CAN-Message-DB/blob/master/Opel/Astra/H/MS-CAN/Body.md
				FormatSeperators: []string{"-", "-", "T", ":", ":"}, // 2007-12-24T18:21
				Condition:        "1 == 1",                          // ever
				Name:             DateTime,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANAirConditioner),
			CanValueDef: CanValueDef{
				Unit:        "°C",
				Calculation: "${5} / 2",
				Condition:   "${0} == 0x22",
				Name:        ACTemperature,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANAirConditioner),
			CanValueDef: CanValueDef{
				Unit:        "rpm",
				Calculation: "${3}",
				Condition:   "${0} == 0x22",
				Name:        ACFanSpeed,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(EntertainmentCANAirConditioner),
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${3}",
				Condition:   "${0} == 0x21",
				Name:        ACMode,
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

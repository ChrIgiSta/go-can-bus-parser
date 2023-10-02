package can

type CanVars string

const (
	ACTemperature      CanVars = "AC Temperature"
	ACMode             CanVars = "AC Mode"
	ACFanSpeed         CanVars = "AC Fan Speed"
	BatteryVoltage     CanVars = "Battery Voltage"
	BusWakeup          CanVars = "CAN-Bus Wakeup"
	ClutchState        CanVars = "Clutch State"
	DateTime           CanVars = "Date"
	EngineSpeedRPM     CanVars = "Engine RPM"
	FullInjection      CanVars = "Full Injection"
	FullLevel          CanVars = "Full Level"
	LedBrightness      CanVars = "Led Brightness"
	Milage             CanVars = "Milage"
	OutdoorTemperature CanVars = "Output Temperature"
	VehicleSpeed       CanVars = "Speed"
	VolumeDown         CanVars = "Volume-"
	VolumeUp           CanVars = "Volume+"
	DisplayR1C1        CanVars = "Display Row 1 Column 1"
	DisplayR1C2        CanVars = "Display Row 1 Column 2"
	DisplayR1C3        CanVars = "Display Row 1 Column 3"
	DisplayR1C4        CanVars = "Display Row 1 Column 4"
)

const (
	CLUTCH_PRESSED = 0x40
	CLUTCH_CLOSED  = 0x00
)

// Opel's LowSpeed SW-CAN
type GMLanArbitrationIDs uint32

const (
	GMLanBusWakeup                GMLanArbitrationIDs = 0x100
	GMLanSpeeds                   GMLanArbitrationIDs = 0x108
	GMLanFullInjection            GMLanArbitrationIDs = 0x130
	GMLanDate                     GMLanArbitrationIDs = 0x180 // tested
	GMLanMilage                   GMLanArbitrationIDs = 0x190
	GMLanWeelRemoteControll       GMLanArbitrationIDs = 0x206 // tested
	GMLanLedBrightness            GMLanArbitrationIDs = 0x235
	GMLanClutch                   GMLanArbitrationIDs = 0x360
	GMLanFullLevel                GMLanArbitrationIDs = 0x375
	GMLanBatteryVoltage           GMLanArbitrationIDs = 0x500
	GMLanOutdoorTemperatureSensor GMLanArbitrationIDs = 0x683 // tested
)

// Opel's MidSpeed CAN
type EntertainmentCANArbitrationIDs uint32

const (
	EntertainmentCANDisplayData    EntertainmentCANArbitrationIDs = 0x6c1 // partialy tested
	EntertainmentCANAirConditioner EntertainmentCANArbitrationIDs = 0x6c8 // tested
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
			ArbitrationID: uint32(GMLanOutdoorTemperatureSensor),
			CanValueDef: CanValueDef{
				Unit:        "°C",
				Calculation: "${2} / 2 - 40", // as normal calculation
				Condition:   "${0} == 0x46 && ${1} == 0x01",
				Name:        OutdoorTemperature,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanDate),
			CanValueDef: CanValueDef{
				Unit:             "",
				Calculation:      "${1};${2};${3};${4}+4;${5};${6}", // as formated string (;)
				FormatSeperators: []string{"-", "-", "T", ":", ":"}, // 2007-12-24T18:21
				Condition:        "1 == 1",                          // ever
				Name:             DateTime,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanWeelRemoteControll),
			CanValueDef: CanValueDef{
				Unit:        "Key Action",
				Calculation: "1",
				Condition:   "${0} == 0x08 && ${1} == 0x93 && ${2} == 0xff",
				Name:        VolumeUp,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanWeelRemoteControll),
			CanValueDef: CanValueDef{
				Unit:        "Key Action",
				Calculation: "1",
				Condition:   "${0} == 0x08 && ${1} == 0x93 && ${2} == 0x01",
				Name:        VolumeDown,
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
			ArbitrationID: uint32(GMLanSpeeds),
			CanValueDef: CanValueDef{
				Unit:        "RPM",
				Calculation: "(${4}*256 + ${5})",
				Condition:   "${0} == 0x23",
				Name:        EngineSpeedRPM,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanSpeeds),
			CanValueDef: CanValueDef{
				Unit:        "km/h",
				Calculation: "(${1}*256 + ${2}) / 128",
				Condition:   "${0} == 0x23",
				Name:        VehicleSpeed,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanMilage),
			CanValueDef: CanValueDef{
				Unit:        "km",
				Calculation: "(${2}*256 + ${4}) / 64",
				Condition:   "${0} == 0x23",
				Name:        Milage,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanClutch),
			CanValueDef: CanValueDef{
				Unit:        "",
				Calculation: "${2}",
				Condition:   "${0} == 0x00 && ${1} == 0x00",
				Name:        ClutchState,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanBatteryVoltage),
			CanValueDef: CanValueDef{
				Unit:        "",
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
				Calculation: "${1}",
				Condition:   "${0} == 0x00",
				Name:        FullLevel,
			},
			TriggerEvent: true,
		},
		{
			ArbitrationID: uint32(GMLanFullLevel),
			CanValueDef: CanValueDef{
				Unit:        "ml",
				Calculation: "(${1} * 256 + ${2}) * 0.03054",
				Condition:   "${0} == 0x00",
				Name:        FullInjection,
			},
			TriggerEvent: true,
		},
	}
}

func EntertainmentCANValueMapps() []CanValueMap {
	return []CanValueMap{
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

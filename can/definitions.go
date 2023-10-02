package can

type CanVars string

const (
	ACTemperature      CanVars = "AC Temperature"
	ACMode             CanVars = "AC Mode"
	ACFanSpeed         CanVars = "AC Fan Speed"
	DateTime           CanVars = "Date"
	OutdoorTemperature CanVars = "Output Temperature"
	VolumeDown         CanVars = "Volume-"
	VolumeUp           CanVars = "Volume+"
)

// Opel's LowSpeed SW-CAN
type GMLanArbitrationIDs uint32

const (
	GMLanDate                     GMLanArbitrationIDs = 0x180
	GMLanWeelRemoteControll       GMLanArbitrationIDs = 0x206
	GMLanOutdoorTemperatureSensor GMLanArbitrationIDs = 0x683
)

// Opel's MidSpeed CAN
type EntertainmentCANArbitrationIDs uint32

const (
	EntertainmentCANDisplayData    EntertainmentCANArbitrationIDs = 0x6c1
	EntertainmentCANAirConditioner EntertainmentCANArbitrationIDs = 0x6c8
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
	}
}

func HighSpeedValueMapps() []CanValueMap {
	return []CanValueMap{
		{},
	}
}

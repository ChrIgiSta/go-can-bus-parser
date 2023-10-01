package can

const (
	OutdoorTemperature = "Output Temperature"
)

// Opel's LowSpeed SW-CAN
type GMLanArbitrationIDs uint32

const (
	GMLanOutdoorTemperatureSensor GMLanArbitrationIDs = 0x683
	GMLanWeelRemoteControll       GMLanArbitrationIDs = 0x206
	GMLanDate                     GMLanArbitrationIDs = 0x180
)

// Opel's MidSpeed CAN
type EntertainmentCANArbitrationIDs uint32

const (
	EntertainmentCANDisplayData EntertainmentCANArbitrationIDs = 0x6c1
)

type HighSpeedCANArbitrationIDs uint32

const ()

type CanValueDef struct {
	Unit        string
	Calculation string
	Condition   string
	Name        string
	Value       interface{}
}

type CanValueMap struct {
	CanValueDef   CanValueDef
	ArbitrationID GMLanArbitrationIDs
}

// Opel Astra H OPC 2006
func GMLanValueMapps() []CanValueMap {
	return []CanValueMap{
		{
			ArbitrationID: GMLanOutdoorTemperatureSensor,
			CanValueDef: CanValueDef{
				Unit:        "°C",
				Calculation: "${2} / 2 - 40",
				Condition:   "${0} == 0x46 && ${1} == 0x01",
				Name:        OutdoorTemperature,
			},
		},
	}
}

func EntertainmentCANValueMapps() []CanValueMap {
	return []CanValueMap{
		{},
	}
}

func HighSpeedValueMapps() []CanValueMap {
	return []CanValueMap{
		{},
	}
}




package streamanalytics

type ClusterSkuName string

const (
	// The default SKU.
	ClusterSkuNameDefault = ClusterSkuName("Default")
)

type CompatibilityLevel string

const (
	CompatibilityLevel_1_0 = CompatibilityLevel("1.0")
)

type Encoding string

const (
	EncodingUTF8 = Encoding("UTF8")
)

type EventSerializationType string

const (
	EventSerializationTypeCsv  = EventSerializationType("Csv")
	EventSerializationTypeAvro = EventSerializationType("Avro")
	EventSerializationTypeJson = EventSerializationType("Json")
)

type EventsOutOfOrderPolicy string

const (
	EventsOutOfOrderPolicyAdjust = EventsOutOfOrderPolicy("Adjust")
	EventsOutOfOrderPolicyDrop   = EventsOutOfOrderPolicy("Drop")
)

type JsonOutputSerializationFormat string

const (
	JsonOutputSerializationFormatLineSeparated = JsonOutputSerializationFormat("LineSeparated")
	JsonOutputSerializationFormatArray         = JsonOutputSerializationFormat("Array")
)

type OutputErrorPolicy string

const (
	OutputErrorPolicyStop = OutputErrorPolicy("Stop")
	OutputErrorPolicyDrop = OutputErrorPolicy("Drop")
)

type OutputStartMode string

const (
	OutputStartModeJobStartTime        = OutputStartMode("JobStartTime")
	OutputStartModeCustomTime          = OutputStartMode("CustomTime")
	OutputStartModeLastOutputEventTime = OutputStartMode("LastOutputEventTime")
)

type SkuName string

const (
	SkuNameStandard = SkuName("Standard")
)

func init() {
}

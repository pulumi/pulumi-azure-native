


package v20170401preview

type AuthenticationMode string

const (
	AuthenticationModeMsi              = AuthenticationMode("Msi")
	AuthenticationModeUserToken        = AuthenticationMode("UserToken")
	AuthenticationModeConnectionString = AuthenticationMode("ConnectionString")
)

type CompatibilityLevel string

const (
	CompatibilityLevel_1_0 = CompatibilityLevel("1.0")
)

type ContentStoragePolicy string

const (
	ContentStoragePolicySystemAccount     = ContentStoragePolicy("SystemAccount")
	ContentStoragePolicyJobStorageAccount = ContentStoragePolicy("JobStorageAccount")
)

type Encoding string

const (
	EncodingUTF8 = Encoding("UTF8")
)

type EventSerializationType string

const (
	EventSerializationTypeCsv       = EventSerializationType("Csv")
	EventSerializationTypeAvro      = EventSerializationType("Avro")
	EventSerializationTypeJson      = EventSerializationType("Json")
	EventSerializationTypeCustomClr = EventSerializationType("CustomClr")
	EventSerializationTypeParquet   = EventSerializationType("Parquet")
	EventSerializationTypeDelta     = EventSerializationType("Delta")
)

type EventsOutOfOrderPolicy string

const (
	EventsOutOfOrderPolicyAdjust = EventsOutOfOrderPolicy("Adjust")
	EventsOutOfOrderPolicyDrop   = EventsOutOfOrderPolicy("Drop")
)

type JobType string

const (
	JobTypeCloud = JobType("Cloud")
	JobTypeEdge  = JobType("Edge")
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

type StreamingJobSkuName string

const (
	StreamingJobSkuNameStandard = StreamingJobSkuName("Standard")
)

func init() {
}

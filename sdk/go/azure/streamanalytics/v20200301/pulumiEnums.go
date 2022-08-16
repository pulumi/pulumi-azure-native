


package v20200301

type AuthenticationMode string

const (
	AuthenticationModeMsi              = AuthenticationMode("Msi")
	AuthenticationModeUserToken        = AuthenticationMode("UserToken")
	AuthenticationModeConnectionString = AuthenticationMode("ConnectionString")
)

type ClusterSkuName string

const (
	// The default SKU.
	ClusterSkuNameDefault = ClusterSkuName("Default")
)

type CompatibilityLevel string

const (
	CompatibilityLevel_1_0 = CompatibilityLevel("1.0")
	CompatibilityLevel_1_2 = CompatibilityLevel("1.2")
)

type CompressionType string

const (
	CompressionTypeNone    = CompressionType("None")
	CompressionTypeGZip    = CompressionType("GZip")
	CompressionTypeDeflate = CompressionType("Deflate")
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
	EventSerializationTypeCsv     = EventSerializationType("Csv")
	EventSerializationTypeAvro    = EventSerializationType("Avro")
	EventSerializationTypeJson    = EventSerializationType("Json")
	EventSerializationTypeParquet = EventSerializationType("Parquet")
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

type RefreshType string

const (
	RefreshTypeStatic                       = RefreshType("Static")
	RefreshTypeRefreshPeriodicallyWithFull  = RefreshType("RefreshPeriodicallyWithFull")
	RefreshTypeRefreshPeriodicallyWithDelta = RefreshType("RefreshPeriodicallyWithDelta")
)

type SkuName string

const (
	SkuNameStandard = SkuName("Standard")
)

func init() {
}

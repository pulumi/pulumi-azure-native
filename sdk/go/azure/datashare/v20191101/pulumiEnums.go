


package v20191101

type DataSetKind string

const (
	DataSetKindBlob               = DataSetKind("Blob")
	DataSetKindContainer          = DataSetKind("Container")
	DataSetKindBlobFolder         = DataSetKind("BlobFolder")
	DataSetKindAdlsGen2FileSystem = DataSetKind("AdlsGen2FileSystem")
	DataSetKindAdlsGen2Folder     = DataSetKind("AdlsGen2Folder")
	DataSetKindAdlsGen2File       = DataSetKind("AdlsGen2File")
	DataSetKindAdlsGen1Folder     = DataSetKind("AdlsGen1Folder")
	DataSetKindAdlsGen1File       = DataSetKind("AdlsGen1File")
	DataSetKindKustoCluster       = DataSetKind("KustoCluster")
	DataSetKindKustoDatabase      = DataSetKind("KustoDatabase")
	DataSetKindSqlDBTable         = DataSetKind("SqlDBTable")
	DataSetKindSqlDWTable         = DataSetKind("SqlDWTable")
)

type DataSetMappingKind string

const (
	DataSetMappingKindBlob               = DataSetMappingKind("Blob")
	DataSetMappingKindContainer          = DataSetMappingKind("Container")
	DataSetMappingKindBlobFolder         = DataSetMappingKind("BlobFolder")
	DataSetMappingKindAdlsGen2FileSystem = DataSetMappingKind("AdlsGen2FileSystem")
	DataSetMappingKindAdlsGen2Folder     = DataSetMappingKind("AdlsGen2Folder")
	DataSetMappingKindAdlsGen2File       = DataSetMappingKind("AdlsGen2File")
	DataSetMappingKindKustoCluster       = DataSetMappingKind("KustoCluster")
	DataSetMappingKindKustoDatabase      = DataSetMappingKind("KustoDatabase")
	DataSetMappingKindSqlDBTable         = DataSetMappingKind("SqlDBTable")
	DataSetMappingKindSqlDWTable         = DataSetMappingKind("SqlDWTable")
)

type OutputType string

const (
	OutputTypeCsv     = OutputType("Csv")
	OutputTypeParquet = OutputType("Parquet")
)

type RecurrenceInterval string

const (
	RecurrenceIntervalHour = RecurrenceInterval("Hour")
	RecurrenceIntervalDay  = RecurrenceInterval("Day")
)

type ShareKind string

const (
	ShareKindCopyBased = ShareKind("CopyBased")
	ShareKindInPlace   = ShareKind("InPlace")
)

type SynchronizationMode string

const (
	SynchronizationModeIncremental = SynchronizationMode("Incremental")
	SynchronizationModeFullSync    = SynchronizationMode("FullSync")
)

type SynchronizationSettingKind string

const (
	SynchronizationSettingKindScheduleBased = SynchronizationSettingKind("ScheduleBased")
)

type TriggerKind string

const (
	TriggerKindScheduleBased = TriggerKind("ScheduleBased")
)

type Type string

const (
	TypeSystemAssigned = Type("SystemAssigned")
)

func init() {
}

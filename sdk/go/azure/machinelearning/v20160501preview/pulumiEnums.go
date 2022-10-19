


package v20160501preview

type AssetType string

const (
	AssetTypeModule   = AssetType("Module")
	AssetTypeResource = AssetType("Resource")
)

type ColumnFormat string

const (
	ColumnFormatByte             = ColumnFormat("Byte")
	ColumnFormatChar             = ColumnFormat("Char")
	ColumnFormatComplex64        = ColumnFormat("Complex64")
	ColumnFormatComplex128       = ColumnFormat("Complex128")
	ColumnFormat_Date_time       = ColumnFormat("Date-time")
	ColumnFormat_Date_timeOffset = ColumnFormat("Date-timeOffset")
	ColumnFormatDouble           = ColumnFormat("Double")
	ColumnFormatDuration         = ColumnFormat("Duration")
	ColumnFormatFloat            = ColumnFormat("Float")
	ColumnFormatInt8             = ColumnFormat("Int8")
	ColumnFormatInt16            = ColumnFormat("Int16")
	ColumnFormatInt32            = ColumnFormat("Int32")
	ColumnFormatInt64            = ColumnFormat("Int64")
	ColumnFormatUint8            = ColumnFormat("Uint8")
	ColumnFormatUint16           = ColumnFormat("Uint16")
	ColumnFormatUint32           = ColumnFormat("Uint32")
	ColumnFormatUint64           = ColumnFormat("Uint64")
)

type ColumnType string

const (
	ColumnTypeBoolean = ColumnType("Boolean")
	ColumnTypeInteger = ColumnType("Integer")
	ColumnTypeNumber  = ColumnType("Number")
	ColumnTypeString  = ColumnType("String")
)

type DiagnosticsLevel string

const (
	DiagnosticsLevelNone  = DiagnosticsLevel("None")
	DiagnosticsLevelError = DiagnosticsLevel("Error")
	DiagnosticsLevelAll   = DiagnosticsLevel("All")
)

type InputPortType string

const (
	InputPortTypeDataset = InputPortType("Dataset")
)

type OutputPortType string

const (
	OutputPortTypeDataset = OutputPortType("Dataset")
)

type ParameterType string

const (
	ParameterTypeString          = ParameterType("String")
	ParameterTypeInt             = ParameterType("Int")
	ParameterTypeFloat           = ParameterType("Float")
	ParameterTypeEnumerated      = ParameterType("Enumerated")
	ParameterTypeScript          = ParameterType("Script")
	ParameterTypeMode            = ParameterType("Mode")
	ParameterTypeCredential      = ParameterType("Credential")
	ParameterTypeBoolean         = ParameterType("Boolean")
	ParameterTypeDouble          = ParameterType("Double")
	ParameterTypeColumnPicker    = ParameterType("ColumnPicker")
	ParameterTypeParameterRange  = ParameterType("ParameterRange")
	ParameterTypeDataGatewayName = ParameterType("DataGatewayName")
)

func init() {
}

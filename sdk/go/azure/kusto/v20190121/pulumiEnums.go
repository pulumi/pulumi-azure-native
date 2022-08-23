


package v20190121

type AzureSkuName string

const (
	AzureSkuName_Standard_DS13_v2_1TB_PS    = AzureSkuName("Standard_DS13_v2+1TB_PS")
	AzureSkuName_Standard_DS13_v2_2TB_PS    = AzureSkuName("Standard_DS13_v2+2TB_PS")
	AzureSkuName_Standard_DS14_v2_3TB_PS    = AzureSkuName("Standard_DS14_v2+3TB_PS")
	AzureSkuName_Standard_DS14_v2_4TB_PS    = AzureSkuName("Standard_DS14_v2+4TB_PS")
	AzureSkuName_Standard_D13_v2            = AzureSkuName("Standard_D13_v2")
	AzureSkuName_Standard_D14_v2            = AzureSkuName("Standard_D14_v2")
	AzureSkuName_Standard_L8s               = AzureSkuName("Standard_L8s")
	AzureSkuName_Standard_L16s              = AzureSkuName("Standard_L16s")
	AzureSkuName_Standard_D11_v2            = AzureSkuName("Standard_D11_v2")
	AzureSkuName_Standard_D12_v2            = AzureSkuName("Standard_D12_v2")
	AzureSkuName_Standard_L4s               = AzureSkuName("Standard_L4s")
	AzureSkuName_Dev_No_SLA_Standard_D11_v2 = AzureSkuName("Dev(No SLA)_Standard_D11_v2")
)

type AzureSkuTier string

const (
	AzureSkuTierBasic    = AzureSkuTier("Basic")
	AzureSkuTierStandard = AzureSkuTier("Standard")
)

type DataFormat string

const (
	DataFormatMULTIJSON  = DataFormat("MULTIJSON")
	DataFormatJSON       = DataFormat("JSON")
	DataFormatCSV        = DataFormat("CSV")
	DataFormatTSV        = DataFormat("TSV")
	DataFormatSCSV       = DataFormat("SCSV")
	DataFormatSOHSV      = DataFormat("SOHSV")
	DataFormatPSV        = DataFormat("PSV")
	DataFormatTXT        = DataFormat("TXT")
	DataFormatRAW        = DataFormat("RAW")
	DataFormatSINGLEJSON = DataFormat("SINGLEJSON")
	DataFormatAVRO       = DataFormat("AVRO")
)

type Kind string

const (
	KindEventHub  = Kind("EventHub")
	KindEventGrid = Kind("EventGrid")
)

func init() {
}

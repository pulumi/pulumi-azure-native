


package v20190701

type AccountType string

const (
	AccountTypeGeneralPurposeStorage = AccountType("GeneralPurposeStorage")
	AccountTypeBlobStorage           = AccountType("BlobStorage")
)

type AzureContainerDataFormat string

const (
	AzureContainerDataFormatBlockBlob = AzureContainerDataFormat("BlockBlob")
	AzureContainerDataFormatPageBlob  = AzureContainerDataFormat("PageBlob")
	AzureContainerDataFormatAzureFile = AzureContainerDataFormat("AzureFile")
)

type ClientPermissionType string

const (
	ClientPermissionTypeNoAccess  = ClientPermissionType("NoAccess")
	ClientPermissionTypeReadOnly  = ClientPermissionType("ReadOnly")
	ClientPermissionTypeReadWrite = ClientPermissionType("ReadWrite")
)

type DataBoxEdgeDeviceStatus string

const (
	DataBoxEdgeDeviceStatusReadyToSetup          = DataBoxEdgeDeviceStatus("ReadyToSetup")
	DataBoxEdgeDeviceStatusOnline                = DataBoxEdgeDeviceStatus("Online")
	DataBoxEdgeDeviceStatusOffline               = DataBoxEdgeDeviceStatus("Offline")
	DataBoxEdgeDeviceStatusNeedsAttention        = DataBoxEdgeDeviceStatus("NeedsAttention")
	DataBoxEdgeDeviceStatusDisconnected          = DataBoxEdgeDeviceStatus("Disconnected")
	DataBoxEdgeDeviceStatusPartiallyDisconnected = DataBoxEdgeDeviceStatus("PartiallyDisconnected")
	DataBoxEdgeDeviceStatusMaintenance           = DataBoxEdgeDeviceStatus("Maintenance")
)

type DataPolicy string

const (
	DataPolicyCloud = DataPolicy("Cloud")
	DataPolicyLocal = DataPolicy("Local")
)

type DayOfWeek string

const (
	DayOfWeekSunday    = DayOfWeek("Sunday")
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
)

type EncryptionAlgorithm string

const (
	EncryptionAlgorithmNone               = EncryptionAlgorithm("None")
	EncryptionAlgorithmAES256             = EncryptionAlgorithm("AES256")
	EncryptionAlgorithm_RSAES_PKCS1_v_1_5 = EncryptionAlgorithm("RSAES_PKCS1_v_1_5")
)

type MonitoringStatus string

const (
	MonitoringStatusEnabled  = MonitoringStatus("Enabled")
	MonitoringStatusDisabled = MonitoringStatus("Disabled")
)

type OrderStateEnum string

const (
	OrderStateEnumUntracked              = OrderStateEnum("Untracked")
	OrderStateEnumAwaitingFulfilment     = OrderStateEnum("AwaitingFulfilment")
	OrderStateEnumAwaitingPreparation    = OrderStateEnum("AwaitingPreparation")
	OrderStateEnumAwaitingShipment       = OrderStateEnum("AwaitingShipment")
	OrderStateEnumShipped                = OrderStateEnum("Shipped")
	OrderStateEnumArriving               = OrderStateEnum("Arriving")
	OrderStateEnumDelivered              = OrderStateEnum("Delivered")
	OrderStateEnumReplacementRequested   = OrderStateEnum("ReplacementRequested")
	OrderStateEnumLostDevice             = OrderStateEnum("LostDevice")
	OrderStateEnumDeclined               = OrderStateEnum("Declined")
	OrderStateEnumReturnInitiated        = OrderStateEnum("ReturnInitiated")
	OrderStateEnumAwaitingReturnShipment = OrderStateEnum("AwaitingReturnShipment")
	OrderStateEnumShippedBack            = OrderStateEnum("ShippedBack")
	OrderStateEnumCollectedAtMicrosoft   = OrderStateEnum("CollectedAtMicrosoft")
)

type PlatformType string

const (
	PlatformTypeWindows = PlatformType("Windows")
	PlatformTypeLinux   = PlatformType("Linux")
)

type RoleStatus string

const (
	RoleStatusEnabled  = RoleStatus("Enabled")
	RoleStatusDisabled = RoleStatus("Disabled")
)

type RoleTypes string

const (
	RoleTypesIOT       = RoleTypes("IOT")
	RoleTypesASA       = RoleTypes("ASA")
	RoleTypesFunctions = RoleTypes("Functions")
	RoleTypesCognitive = RoleTypes("Cognitive")
)

type SSLStatus string

const (
	SSLStatusEnabled  = SSLStatus("Enabled")
	SSLStatusDisabled = SSLStatus("Disabled")
)

type ShareAccessProtocol string

const (
	ShareAccessProtocolSMB = ShareAccessProtocol("SMB")
	ShareAccessProtocolNFS = ShareAccessProtocol("NFS")
)

type ShareAccessType string

const (
	ShareAccessTypeChange = ShareAccessType("Change")
	ShareAccessTypeRead   = ShareAccessType("Read")
	ShareAccessTypeCustom = ShareAccessType("Custom")
)

type ShareStatus string

const (
	ShareStatusOffline        = ShareStatus("Offline")
	ShareStatusUnknown        = ShareStatus("Unknown")
	ShareStatusOK             = ShareStatus("OK")
	ShareStatusUpdating       = ShareStatus("Updating")
	ShareStatusNeedsAttention = ShareStatus("NeedsAttention")
)

type SkuName string

const (
	SkuNameGateway = SkuName("Gateway")
	SkuNameEdge    = SkuName("Edge")
)

type SkuTier string

const (
	SkuTierStandard = SkuTier("Standard")
)

type TriggerEventType string

const (
	TriggerEventTypeFileEvent          = TriggerEventType("FileEvent")
	TriggerEventTypePeriodicTimerEvent = TriggerEventType("PeriodicTimerEvent")
)

func init() {
}

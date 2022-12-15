


package v20210201preview

type AccountType string

const (
	AccountTypeGeneralPurposeStorage = AccountType("GeneralPurposeStorage")
	AccountTypeBlobStorage           = AccountType("BlobStorage")
)

type AddonType string

const (
	AddonTypeIotEdge          = AddonType("IotEdge")
	AddonTypeArcForKubernetes = AddonType("ArcForKubernetes")
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

type MsiIdentityType string

const (
	MsiIdentityTypeNone           = MsiIdentityType("None")
	MsiIdentityTypeSystemAssigned = MsiIdentityType("SystemAssigned")
	MsiIdentityTypeUserAssigned   = MsiIdentityType("UserAssigned")
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
	RoleTypesIOT                 = RoleTypes("IOT")
	RoleTypesASA                 = RoleTypes("ASA")
	RoleTypesFunctions           = RoleTypes("Functions")
	RoleTypesCognitive           = RoleTypes("Cognitive")
	RoleTypesMEC                 = RoleTypes("MEC")
	RoleTypesCloudEdgeManagement = RoleTypes("CloudEdgeManagement")
	RoleTypesKubernetes          = RoleTypes("Kubernetes")
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

type ShipmentType string

const (
	ShipmentTypeNotApplicable     = ShipmentType("NotApplicable")
	ShipmentTypeShippedToCustomer = ShipmentType("ShippedToCustomer")
	ShipmentTypeSelfPickup        = ShipmentType("SelfPickup")
)

type SkuName string

const (
	SkuNameGateway               = SkuName("Gateway")
	SkuNameEdge                  = SkuName("Edge")
	SkuName_TEA_1Node            = SkuName("TEA_1Node")
	SkuName_TEA_1Node_UPS        = SkuName("TEA_1Node_UPS")
	SkuName_TEA_1Node_Heater     = SkuName("TEA_1Node_Heater")
	SkuName_TEA_1Node_UPS_Heater = SkuName("TEA_1Node_UPS_Heater")
	SkuName_TEA_4Node_Heater     = SkuName("TEA_4Node_Heater")
	SkuName_TEA_4Node_UPS_Heater = SkuName("TEA_4Node_UPS_Heater")
	SkuNameTMA                   = SkuName("TMA")
	SkuNameTDC                   = SkuName("TDC")
	SkuName_TCA_Small            = SkuName("TCA_Small")
	SkuNameGPU                   = SkuName("GPU")
	SkuName_TCA_Large            = SkuName("TCA_Large")
	SkuName_EdgeP_Base           = SkuName("EdgeP_Base")
	SkuName_EdgeP_High           = SkuName("EdgeP_High")
	SkuName_EdgePR_Base          = SkuName("EdgePR_Base")
	SkuName_EdgePR_Base_UPS      = SkuName("EdgePR_Base_UPS")
	SkuName_EP2_64_1VPU_W        = SkuName("EP2_64_1VPU_W")
	SkuName_EP2_128_1T4_Mx1_W    = SkuName("EP2_128_1T4_Mx1_W")
	SkuName_EP2_256_2T4_W        = SkuName("EP2_256_2T4_W")
	SkuName_EdgeMR_Mini          = SkuName("EdgeMR_Mini")
	SkuName_RCA_Small            = SkuName("RCA_Small")
	SkuName_RCA_Large            = SkuName("RCA_Large")
	SkuNameRDC                   = SkuName("RDC")
	SkuNameManagement            = SkuName("Management")
)

type SkuTier string

const (
	SkuTierStandard = SkuTier("Standard")
)

type StorageAccountStatus string

const (
	StorageAccountStatusOK             = StorageAccountStatus("OK")
	StorageAccountStatusOffline        = StorageAccountStatus("Offline")
	StorageAccountStatusUnknown        = StorageAccountStatus("Unknown")
	StorageAccountStatusUpdating       = StorageAccountStatus("Updating")
	StorageAccountStatusNeedsAttention = StorageAccountStatus("NeedsAttention")
)

type TriggerEventType string

const (
	TriggerEventTypeFileEvent          = TriggerEventType("FileEvent")
	TriggerEventTypePeriodicTimerEvent = TriggerEventType("PeriodicTimerEvent")
)

type UserType string

const (
	UserTypeShare           = UserType("Share")
	UserTypeLocalManagement = UserType("LocalManagement")
	UserTypeARM             = UserType("ARM")
)

func init() {
}




package v20151101preview

type DataSourceKind string

const (
	DataSourceKindAzureActivityLog              = DataSourceKind("AzureActivityLog")
	DataSourceKindChangeTrackingPath            = DataSourceKind("ChangeTrackingPath")
	DataSourceKindChangeTrackingDefaultPath     = DataSourceKind("ChangeTrackingDefaultPath")
	DataSourceKindChangeTrackingDefaultRegistry = DataSourceKind("ChangeTrackingDefaultRegistry")
	DataSourceKindChangeTrackingCustomRegistry  = DataSourceKind("ChangeTrackingCustomRegistry")
	DataSourceKindCustomLog                     = DataSourceKind("CustomLog")
	DataSourceKindCustomLogCollection           = DataSourceKind("CustomLogCollection")
	DataSourceKindGenericDataSource             = DataSourceKind("GenericDataSource")
	DataSourceKindIISLogs                       = DataSourceKind("IISLogs")
	DataSourceKindLinuxPerformanceObject        = DataSourceKind("LinuxPerformanceObject")
	DataSourceKindLinuxPerformanceCollection    = DataSourceKind("LinuxPerformanceCollection")
	DataSourceKindLinuxSyslog                   = DataSourceKind("LinuxSyslog")
	DataSourceKindLinuxSyslogCollection         = DataSourceKind("LinuxSyslogCollection")
	DataSourceKindWindowsEvent                  = DataSourceKind("WindowsEvent")
	DataSourceKindWindowsPerformanceCounter     = DataSourceKind("WindowsPerformanceCounter")
)

type EntityStatus string

const (
	EntityStatusCreating            = EntityStatus("Creating")
	EntityStatusSucceeded           = EntityStatus("Succeeded")
	EntityStatusFailed              = EntityStatus("Failed")
	EntityStatusCanceled            = EntityStatus("Canceled")
	EntityStatusDeleting            = EntityStatus("Deleting")
	EntityStatusProvisioningAccount = EntityStatus("ProvisioningAccount")
)

type MachineGroupType string

const (
	MachineGroupTypeUnknown      = MachineGroupType("unknown")
	MachineGroupType_Azure_cs    = MachineGroupType("azure-cs")
	MachineGroupType_Azure_sf    = MachineGroupType("azure-sf")
	MachineGroupType_Azure_vmss  = MachineGroupType("azure-vmss")
	MachineGroupType_User_static = MachineGroupType("user-static")
)

type SkuNameEnum string

const (
	SkuNameEnumFree                = SkuNameEnum("Free")
	SkuNameEnumStandard            = SkuNameEnum("Standard")
	SkuNameEnumPremium             = SkuNameEnum("Premium")
	SkuNameEnumPerNode             = SkuNameEnum("PerNode")
	SkuNameEnumPerGB2018           = SkuNameEnum("PerGB2018")
	SkuNameEnumStandalone          = SkuNameEnum("Standalone")
	SkuNameEnumCapacityReservation = SkuNameEnum("CapacityReservation")
)

func init() {
}

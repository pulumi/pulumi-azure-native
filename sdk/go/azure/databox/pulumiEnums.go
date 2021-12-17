


package databox

type AddressType string

const (
	// Address type not known.
	AddressTypeNone = AddressType("None")
	// Residential Address.
	AddressTypeResidential = AddressType("Residential")
	// Commercial Address.
	AddressTypeCommercial = AddressType("Commercial")
)

type ClassDiscriminator string

const (
	// Data Box orders.
	ClassDiscriminatorDataBox = ClassDiscriminator("DataBox")
	// Data Box Disk orders.
	ClassDiscriminatorDataBoxDisk = ClassDiscriminator("DataBoxDisk")
	// Data Box Heavy orders.
	ClassDiscriminatorDataBoxHeavy = ClassDiscriminator("DataBoxHeavy")
)

type DataAccountType string

const (
	// Storage Accounts .
	DataAccountTypeStorageAccount = DataAccountType("StorageAccount")
	// Azure Managed disk storage.
	DataAccountTypeManagedDisk = DataAccountType("ManagedDisk")
)

type DoubleEncryption string

const (
	// Software-based encryption is enabled.
	DoubleEncryptionEnabled = DoubleEncryption("Enabled")
	// Software-based encryption is disabled.
	DoubleEncryptionDisabled = DoubleEncryption("Disabled")
)

type FilterFileType string

const (
	// Filter file is of the type AzureBlob.
	FilterFileTypeAzureBlob = FilterFileType("AzureBlob")
	// Filter file is of the type AzureFiles.
	FilterFileTypeAzureFile = FilterFileType("AzureFile")
)

type JobDeliveryType string

const (
	// Non Scheduled job.
	JobDeliveryTypeNonScheduled = JobDeliveryType("NonScheduled")
	// Scheduled job.
	JobDeliveryTypeScheduled = JobDeliveryType("Scheduled")
)

type KekType string

const (
	// Key encryption key is managed by Microsoft.
	KekTypeMicrosoftManaged = KekType("MicrosoftManaged")
	// Key encryption key is managed by the Customer.
	KekTypeCustomerManaged = KekType("CustomerManaged")
)

type LogCollectionLevel string

const (
	// Only Errors will be collected in the logs.
	LogCollectionLevelError = LogCollectionLevel("Error")
	// Verbose logging (includes Errors, CRC, size information and others).
	LogCollectionLevelVerbose = LogCollectionLevel("Verbose")
)

type NotificationStageName string

const (
	// Notification at device prepared stage.
	NotificationStageNameDevicePrepared = NotificationStageName("DevicePrepared")
	// Notification at device dispatched stage.
	NotificationStageNameDispatched = NotificationStageName("Dispatched")
	// Notification at device delivered stage.
	NotificationStageNameDelivered = NotificationStageName("Delivered")
	// Notification at device picked up from user stage.
	NotificationStageNamePickedUp = NotificationStageName("PickedUp")
	// Notification at device received at Azure datacenter stage.
	NotificationStageNameAtAzureDC = NotificationStageName("AtAzureDC")
	// Notification at data copy started stage.
	NotificationStageNameDataCopy = NotificationStageName("DataCopy")
)

type SkuName string

const (
	// Data Box.
	SkuNameDataBox = SkuName("DataBox")
	// Data Box Disk.
	SkuNameDataBoxDisk = SkuName("DataBoxDisk")
	// Data Box Heavy.
	SkuNameDataBoxHeavy = SkuName("DataBoxHeavy")
)

type TransferConfigurationType string

const (
	// Transfer all the data.
	TransferConfigurationTypeTransferAll = TransferConfigurationType("TransferAll")
	// Transfer using filter.
	TransferConfigurationTypeTransferUsingFilter = TransferConfigurationType("TransferUsingFilter")
)

type TransferType string

const (
	// Import data to azure.
	TransferTypeImportToAzure = TransferType("ImportToAzure")
	// Export data from azure.
	TransferTypeExportFromAzure = TransferType("ExportFromAzure")
)

type TransportShipmentTypes string

const (
	// Shipment Logistics is handled by the customer.
	TransportShipmentTypesCustomerManaged = TransportShipmentTypes("CustomerManaged")
	// Shipment Logistics is handled by Microsoft.
	TransportShipmentTypesMicrosoftManaged = TransportShipmentTypes("MicrosoftManaged")
)

func init() {
}

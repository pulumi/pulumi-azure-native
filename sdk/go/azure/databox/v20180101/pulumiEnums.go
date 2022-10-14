


package v20180101

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
	// DataBox orders.
	ClassDiscriminatorDataBox = ClassDiscriminator("DataBox")
	// DataBoxDisk orders.
	ClassDiscriminatorDataBoxDisk = ClassDiscriminator("DataBoxDisk")
	// DataBoxHeavy orders.
	ClassDiscriminatorDataBoxHeavy = ClassDiscriminator("DataBoxHeavy")
)

type DataDestinationType string

const (
	// Unknown type.
	DataDestinationTypeUnknownType = DataDestinationType("UnknownType")
	// Storage Accounts .
	DataDestinationTypeStorageAccount = DataDestinationType("StorageAccount")
	// Azure Managed disk storage.
	DataDestinationTypeManagedDisk = DataDestinationType("ManagedDisk")
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
	// Notification at device received at azure datacenter stage.
	NotificationStageNameAtAzureDC = NotificationStageName("AtAzureDC")
	// Notification at data copy started stage.
	NotificationStageNameDataCopy = NotificationStageName("DataCopy")
)

type SkuName string

const (
	// DataBox.
	SkuNameDataBox = SkuName("DataBox")
	// DataBoxDisk.
	SkuNameDataBoxDisk = SkuName("DataBoxDisk")
	// DataBoxHeavy.
	SkuNameDataBoxHeavy = SkuName("DataBoxHeavy")
)

func init() {
}

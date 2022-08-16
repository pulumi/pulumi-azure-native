


package v20190901

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
	// Databox orders.
	ClassDiscriminatorDataBox = ClassDiscriminator("DataBox")
	// DataboxDisk orders.
	ClassDiscriminatorDataBoxDisk = ClassDiscriminator("DataBoxDisk")
	// DataboxHeavy orders.
	ClassDiscriminatorDataBoxHeavy = ClassDiscriminator("DataBoxHeavy")
)

type DataDestinationType string

const (
	// Storage Accounts .
	DataDestinationTypeStorageAccount = DataDestinationType("StorageAccount")
	// Azure Managed disk storage.
	DataDestinationTypeManagedDisk = DataDestinationType("ManagedDisk")
)

type JobDeliveryType string

const (
	// Non Scheduled job.
	JobDeliveryTypeNonScheduled = JobDeliveryType("NonScheduled")
	// Scheduled job.
	JobDeliveryTypeScheduled = JobDeliveryType("Scheduled")
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
	// Databox.
	SkuNameDataBox = SkuName("DataBox")
	// DataboxDisk.
	SkuNameDataBoxDisk = SkuName("DataBoxDisk")
	// DataboxHeavy.
	SkuNameDataBoxHeavy = SkuName("DataBoxHeavy")
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

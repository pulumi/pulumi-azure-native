


package edgeorder

type AddressType string

const (
	// Address type not known.
	AddressTypeNone = AddressType("None")
	// Residential Address.
	AddressTypeResidential = AddressType("Residential")
	// Commercial Address.
	AddressTypeCommercial = AddressType("Commercial")
)

type DoubleEncryptionStatus string

const (
	// Double encryption is disabled
	DoubleEncryptionStatusDisabled = DoubleEncryptionStatus("Disabled")
	// Double encryption is enabled
	DoubleEncryptionStatusEnabled = DoubleEncryptionStatus("Enabled")
)

type NotificationStageName string

const (
	// Notification at order item shipped from microsoft datacenter.
	NotificationStageNameShipped = NotificationStageName("Shipped")
	// Notification at order item delivered to customer.
	NotificationStageNameDelivered = NotificationStageName("Delivered")
)

type OrderItemType string

const (
	// Purchase OrderItem.
	OrderItemTypePurchase = OrderItemType("Purchase")
	// Rental OrderItem.
	OrderItemTypeRental = OrderItemType("Rental")
)

type SupportedFilterTypes string

const (
	// Ship to country
	SupportedFilterTypesShipToCountries = SupportedFilterTypes("ShipToCountries")
	// Double encryption status
	SupportedFilterTypesDoubleEncryptionStatus = SupportedFilterTypes("DoubleEncryptionStatus")
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

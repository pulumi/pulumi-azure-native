


package v20201015preview

type AdvancedFilterOperatorType string

const (
	AdvancedFilterOperatorTypeNumberIn                  = AdvancedFilterOperatorType("NumberIn")
	AdvancedFilterOperatorTypeNumberNotIn               = AdvancedFilterOperatorType("NumberNotIn")
	AdvancedFilterOperatorTypeNumberLessThan            = AdvancedFilterOperatorType("NumberLessThan")
	AdvancedFilterOperatorTypeNumberGreaterThan         = AdvancedFilterOperatorType("NumberGreaterThan")
	AdvancedFilterOperatorTypeNumberLessThanOrEquals    = AdvancedFilterOperatorType("NumberLessThanOrEquals")
	AdvancedFilterOperatorTypeNumberGreaterThanOrEquals = AdvancedFilterOperatorType("NumberGreaterThanOrEquals")
	AdvancedFilterOperatorTypeBoolEquals                = AdvancedFilterOperatorType("BoolEquals")
	AdvancedFilterOperatorTypeStringIn                  = AdvancedFilterOperatorType("StringIn")
	AdvancedFilterOperatorTypeStringNotIn               = AdvancedFilterOperatorType("StringNotIn")
	AdvancedFilterOperatorTypeStringBeginsWith          = AdvancedFilterOperatorType("StringBeginsWith")
	AdvancedFilterOperatorTypeStringEndsWith            = AdvancedFilterOperatorType("StringEndsWith")
	AdvancedFilterOperatorTypeStringContains            = AdvancedFilterOperatorType("StringContains")
	AdvancedFilterOperatorTypeNumberInRange             = AdvancedFilterOperatorType("NumberInRange")
	AdvancedFilterOperatorTypeNumberNotInRange          = AdvancedFilterOperatorType("NumberNotInRange")
	AdvancedFilterOperatorTypeStringNotBeginsWith       = AdvancedFilterOperatorType("StringNotBeginsWith")
	AdvancedFilterOperatorTypeStringNotEndsWith         = AdvancedFilterOperatorType("StringNotEndsWith")
	AdvancedFilterOperatorTypeStringNotContains         = AdvancedFilterOperatorType("StringNotContains")
	AdvancedFilterOperatorTypeIsNullOrUndefined         = AdvancedFilterOperatorType("IsNullOrUndefined")
	AdvancedFilterOperatorTypeIsNotNull                 = AdvancedFilterOperatorType("IsNotNull")
)

type DeadLetterEndPointType string

const (
	DeadLetterEndPointTypeStorageBlob = DeadLetterEndPointType("StorageBlob")
)

type DeliveryAttributeMappingType string

const (
	DeliveryAttributeMappingTypeStatic  = DeliveryAttributeMappingType("Static")
	DeliveryAttributeMappingTypeDynamic = DeliveryAttributeMappingType("Dynamic")
)

type EndpointType string

const (
	EndpointTypeWebHook          = EndpointType("WebHook")
	EndpointTypeEventHub         = EndpointType("EventHub")
	EndpointTypeStorageQueue     = EndpointType("StorageQueue")
	EndpointTypeHybridConnection = EndpointType("HybridConnection")
	EndpointTypeServiceBusQueue  = EndpointType("ServiceBusQueue")
	EndpointTypeServiceBusTopic  = EndpointType("ServiceBusTopic")
	EndpointTypeAzureFunction    = EndpointType("AzureFunction")
)

type EventDeliverySchema string

const (
	EventDeliverySchemaEventGridSchema       = EventDeliverySchema("EventGridSchema")
	EventDeliverySchemaCustomInputSchema     = EventDeliverySchema("CustomInputSchema")
	EventDeliverySchema_CloudEventSchemaV1_0 = EventDeliverySchema("CloudEventSchemaV1_0")
)

type EventSubscriptionIdentityType string

const (
	EventSubscriptionIdentityTypeSystemAssigned               = EventSubscriptionIdentityType("SystemAssigned")
	EventSubscriptionIdentityType_SystemAssigned_UserAssigned = EventSubscriptionIdentityType("SystemAssigned, UserAssigned")
	EventSubscriptionIdentityTypeUserAssigned                 = EventSubscriptionIdentityType("UserAssigned")
)

type IdentityType string

const (
	IdentityTypeNone                         = IdentityType("None")
	IdentityTypeSystemAssigned               = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned                 = IdentityType("UserAssigned")
	IdentityType_SystemAssigned_UserAssigned = IdentityType("SystemAssigned, UserAssigned")
)

type InputSchema string

const (
	InputSchemaEventGridSchema       = InputSchema("EventGridSchema")
	InputSchemaCustomEventSchema     = InputSchema("CustomEventSchema")
	InputSchema_CloudEventSchemaV1_0 = InputSchema("CloudEventSchemaV1_0")
)

type InputSchemaMappingType string

const (
	InputSchemaMappingTypeJson = InputSchemaMappingType("Json")
)

type IpActionType string

const (
	IpActionTypeAllow = IpActionType("Allow")
)

type PartnerRegistrationVisibilityState string

const (
	PartnerRegistrationVisibilityStateHidden             = PartnerRegistrationVisibilityState("Hidden")
	PartnerRegistrationVisibilityStatePublicPreview      = PartnerRegistrationVisibilityState("PublicPreview")
	PartnerRegistrationVisibilityStateGenerallyAvailable = PartnerRegistrationVisibilityState("GenerallyAvailable")
)

type PersistedConnectionStatus string

const (
	PersistedConnectionStatusPending      = PersistedConnectionStatus("Pending")
	PersistedConnectionStatusApproved     = PersistedConnectionStatus("Approved")
	PersistedConnectionStatusRejected     = PersistedConnectionStatus("Rejected")
	PersistedConnectionStatusDisconnected = PersistedConnectionStatus("Disconnected")
)

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

type ResourceKind string

const (
	ResourceKindAzure    = ResourceKind("Azure")
	ResourceKindAzureArc = ResourceKind("AzureArc")
)

type ResourceProvisioningState string

const (
	ResourceProvisioningStateCreating  = ResourceProvisioningState("Creating")
	ResourceProvisioningStateUpdating  = ResourceProvisioningState("Updating")
	ResourceProvisioningStateDeleting  = ResourceProvisioningState("Deleting")
	ResourceProvisioningStateSucceeded = ResourceProvisioningState("Succeeded")
	ResourceProvisioningStateCanceled  = ResourceProvisioningState("Canceled")
	ResourceProvisioningStateFailed    = ResourceProvisioningState("Failed")
)

type Sku string

const (
	SkuBasic   = Sku("Basic")
	SkuPremium = Sku("Premium")
)

func init() {
}




package v20220615

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

type ChannelProvisioningState string

const (
	ChannelProvisioningStateCreating                              = ChannelProvisioningState("Creating")
	ChannelProvisioningStateUpdating                              = ChannelProvisioningState("Updating")
	ChannelProvisioningStateDeleting                              = ChannelProvisioningState("Deleting")
	ChannelProvisioningStateSucceeded                             = ChannelProvisioningState("Succeeded")
	ChannelProvisioningStateCanceled                              = ChannelProvisioningState("Canceled")
	ChannelProvisioningStateFailed                                = ChannelProvisioningState("Failed")
	ChannelProvisioningStateIdleDueToMirroredPartnerTopicDeletion = ChannelProvisioningState("IdleDueToMirroredPartnerTopicDeletion")
)

type ChannelType string

const (
	ChannelTypePartnerTopic = ChannelType("PartnerTopic")
)

type DataResidencyBoundary string

const (
	DataResidencyBoundaryWithinGeopair = DataResidencyBoundary("WithinGeopair")
	DataResidencyBoundaryWithinRegion  = DataResidencyBoundary("WithinRegion")
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

type EventDefinitionKind string

const (
	EventDefinitionKindInline = EventDefinitionKind("Inline")
)

type EventDeliverySchema string

const (
	EventDeliverySchemaEventGridSchema       = EventDeliverySchema("EventGridSchema")
	EventDeliverySchemaCustomInputSchema     = EventDeliverySchema("CustomInputSchema")
	EventDeliverySchema_CloudEventSchemaV1_0 = EventDeliverySchema("CloudEventSchemaV1_0")
)

type EventSubscriptionIdentityType string

const (
	EventSubscriptionIdentityTypeSystemAssigned = EventSubscriptionIdentityType("SystemAssigned")
	EventSubscriptionIdentityTypeUserAssigned   = EventSubscriptionIdentityType("UserAssigned")
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

type PartnerConfigurationProvisioningState string

const (
	PartnerConfigurationProvisioningStateCreating  = PartnerConfigurationProvisioningState("Creating")
	PartnerConfigurationProvisioningStateUpdating  = PartnerConfigurationProvisioningState("Updating")
	PartnerConfigurationProvisioningStateDeleting  = PartnerConfigurationProvisioningState("Deleting")
	PartnerConfigurationProvisioningStateSucceeded = PartnerConfigurationProvisioningState("Succeeded")
	PartnerConfigurationProvisioningStateCanceled  = PartnerConfigurationProvisioningState("Canceled")
	PartnerConfigurationProvisioningStateFailed    = PartnerConfigurationProvisioningState("Failed")
)

type PartnerTopicActivationState string

const (
	PartnerTopicActivationStateNeverActivated = PartnerTopicActivationState("NeverActivated")
	PartnerTopicActivationStateActivated      = PartnerTopicActivationState("Activated")
	PartnerTopicActivationStateDeactivated    = PartnerTopicActivationState("Deactivated")
)

type PartnerTopicRoutingMode string

const (
	PartnerTopicRoutingModeSourceEventAttribute = PartnerTopicRoutingMode("SourceEventAttribute")
	PartnerTopicRoutingModeChannelNameHeader    = PartnerTopicRoutingMode("ChannelNameHeader")
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

type ReadinessState string

const (
	ReadinessStateNeverActivated = ReadinessState("NeverActivated")
	ReadinessStateActivated      = ReadinessState("Activated")
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

func init() {
}

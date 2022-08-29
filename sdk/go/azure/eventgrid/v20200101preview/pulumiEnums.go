


package v20200101preview

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
)

type DeadLetterEndPointType string

const (
	DeadLetterEndPointTypeStorageBlob = DeadLetterEndPointType("StorageBlob")
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

func init() {
}

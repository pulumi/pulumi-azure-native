


package v20190201preview

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
)

type EventDeliverySchema string

const (
	EventDeliverySchemaEventGridSchema     = EventDeliverySchema("EventGridSchema")
	EventDeliverySchemaCloudEventV01Schema = EventDeliverySchema("CloudEventV01Schema")
	EventDeliverySchemaCustomInputSchema   = EventDeliverySchema("CustomInputSchema")
)

type InputSchema string

const (
	InputSchemaEventGridSchema     = InputSchema("EventGridSchema")
	InputSchemaCustomEventSchema   = InputSchema("CustomEventSchema")
	InputSchemaCloudEventV01Schema = InputSchema("CloudEventV01Schema")
)

type InputSchemaMappingType string

const (
	InputSchemaMappingTypeJson = InputSchemaMappingType("Json")
)

func init() {
}

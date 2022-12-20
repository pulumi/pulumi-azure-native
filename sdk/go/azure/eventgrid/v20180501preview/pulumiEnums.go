


package v20180501preview

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
)

type EventDeliverySchema string

const (
	EventDeliverySchemaEventGridSchema     = EventDeliverySchema("EventGridSchema")
	EventDeliverySchemaInputEventSchema    = EventDeliverySchema("InputEventSchema")
	EventDeliverySchemaCloudEventV01Schema = EventDeliverySchema("CloudEventV01Schema")
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

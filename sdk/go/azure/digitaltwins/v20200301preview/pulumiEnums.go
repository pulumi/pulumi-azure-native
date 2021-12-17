


package v20200301preview

type DigitalTwinsSku string

const (
	DigitalTwinsSkuF1 = DigitalTwinsSku("F1")
)

type EndpointType string

const (
	EndpointTypeEventHub   = EndpointType("EventHub")
	EndpointTypeEventGrid  = EndpointType("EventGrid")
	EndpointTypeServiceBus = EndpointType("ServiceBus")
)

func init() {
}

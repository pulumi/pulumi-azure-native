


package v20201031

type EndpointType string

const (
	EndpointTypeEventHub   = EndpointType("EventHub")
	EndpointTypeEventGrid  = EndpointType("EventGrid")
	EndpointTypeServiceBus = EndpointType("ServiceBus")
)

func init() {
}

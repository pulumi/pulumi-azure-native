


package v20180101

type EndpointType string

const (
	EndpointTypeWebHook  = EndpointType("WebHook")
	EndpointTypeEventHub = EndpointType("EventHub")
)

func init() {
}

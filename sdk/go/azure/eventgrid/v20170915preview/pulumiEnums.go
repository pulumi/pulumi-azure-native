


package v20170915preview

type EndpointType string

const (
	EndpointTypeWebHook  = EndpointType("WebHook")
	EndpointTypeEventHub = EndpointType("EventHub")
)

func init() {
}




package v20210630preview

type AuthenticationType string

const (
	AuthenticationTypeKeyBased      = AuthenticationType("KeyBased")
	AuthenticationTypeIdentityBased = AuthenticationType("IdentityBased")
)

type ConnectionType string

const (
	ConnectionTypeAzureDataExplorer = ConnectionType("AzureDataExplorer")
)

type DigitalTwinsIdentityType string

const (
	DigitalTwinsIdentityTypeNone           = DigitalTwinsIdentityType("None")
	DigitalTwinsIdentityTypeSystemAssigned = DigitalTwinsIdentityType("SystemAssigned")
)

type EndpointType string

const (
	EndpointTypeEventHub   = EndpointType("EventHub")
	EndpointTypeEventGrid  = EndpointType("EventGrid")
	EndpointTypeServiceBus = EndpointType("ServiceBus")
)

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusPending      = PrivateLinkServiceConnectionStatus("Pending")
	PrivateLinkServiceConnectionStatusApproved     = PrivateLinkServiceConnectionStatus("Approved")
	PrivateLinkServiceConnectionStatusRejected     = PrivateLinkServiceConnectionStatus("Rejected")
	PrivateLinkServiceConnectionStatusDisconnected = PrivateLinkServiceConnectionStatus("Disconnected")
)

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

func init() {
}




package v20200501

type ACLAction string

const (
	ACLActionAllow = ACLAction("Allow")
	ACLActionDeny  = ACLAction("Deny")
)

type FeatureFlags string

const (
	FeatureFlagsServiceMode            = FeatureFlags("ServiceMode")
	FeatureFlagsEnableConnectivityLogs = FeatureFlags("EnableConnectivityLogs")
	FeatureFlagsEnableMessagingLogs    = FeatureFlags("EnableMessagingLogs")
)

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusPending      = PrivateLinkServiceConnectionStatus("Pending")
	PrivateLinkServiceConnectionStatusApproved     = PrivateLinkServiceConnectionStatus("Approved")
	PrivateLinkServiceConnectionStatusRejected     = PrivateLinkServiceConnectionStatus("Rejected")
	PrivateLinkServiceConnectionStatusDisconnected = PrivateLinkServiceConnectionStatus("Disconnected")
)

type ServiceKind string

const (
	ServiceKindSignalR       = ServiceKind("SignalR")
	ServiceKindRawWebSockets = ServiceKind("RawWebSockets")
)

type SignalRRequestType string

const (
	SignalRRequestTypeClientConnection = SignalRRequestType("ClientConnection")
	SignalRRequestTypeServerConnection = SignalRRequestType("ServerConnection")
	SignalRRequestTypeRESTAPI          = SignalRRequestType("RESTAPI")
	SignalRRequestTypeTrace            = SignalRRequestType("Trace")
)

type SignalRSkuTier string

const (
	SignalRSkuTierFree     = SignalRSkuTier("Free")
	SignalRSkuTierBasic    = SignalRSkuTier("Basic")
	SignalRSkuTierStandard = SignalRSkuTier("Standard")
	SignalRSkuTierPremium  = SignalRSkuTier("Premium")
)

func init() {
}

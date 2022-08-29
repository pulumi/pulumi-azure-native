


package webpubsub

type ACLAction string

const (
	ACLActionAllow = ACLAction("Allow")
	ACLActionDeny  = ACLAction("Deny")
)

type FeatureFlags string

const (
	FeatureFlagsEnableConnectivityLogs = FeatureFlags("EnableConnectivityLogs")
	FeatureFlagsEnableMessagingLogs    = FeatureFlags("EnableMessagingLogs")
	FeatureFlagsEnableLiveTrace        = FeatureFlags("EnableLiveTrace")
)

type ManagedIdentityType string

const (
	ManagedIdentityTypeNone           = ManagedIdentityType("None")
	ManagedIdentityTypeSystemAssigned = ManagedIdentityType("SystemAssigned")
	ManagedIdentityTypeUserAssigned   = ManagedIdentityType("UserAssigned")
)

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusPending      = PrivateLinkServiceConnectionStatus("Pending")
	PrivateLinkServiceConnectionStatusApproved     = PrivateLinkServiceConnectionStatus("Approved")
	PrivateLinkServiceConnectionStatusRejected     = PrivateLinkServiceConnectionStatus("Rejected")
	PrivateLinkServiceConnectionStatusDisconnected = PrivateLinkServiceConnectionStatus("Disconnected")
)

type UpstreamAuthType string

const (
	UpstreamAuthTypeNone            = UpstreamAuthType("None")
	UpstreamAuthTypeManagedIdentity = UpstreamAuthType("ManagedIdentity")
)

type WebPubSubRequestType string

const (
	WebPubSubRequestTypeClientConnection = WebPubSubRequestType("ClientConnection")
	WebPubSubRequestTypeServerConnection = WebPubSubRequestType("ServerConnection")
	WebPubSubRequestTypeRESTAPI          = WebPubSubRequestType("RESTAPI")
	WebPubSubRequestTypeTrace            = WebPubSubRequestType("Trace")
)

type WebPubSubSkuTier string

const (
	WebPubSubSkuTierFree     = WebPubSubSkuTier("Free")
	WebPubSubSkuTierBasic    = WebPubSubSkuTier("Basic")
	WebPubSubSkuTierStandard = WebPubSubSkuTier("Standard")
	WebPubSubSkuTierPremium  = WebPubSubSkuTier("Premium")
)

func init() {
}

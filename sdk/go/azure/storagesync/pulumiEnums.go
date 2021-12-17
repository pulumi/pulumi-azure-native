


package storagesync

type FeatureStatus string

const (
	FeatureStatusOn  = FeatureStatus("on")
	FeatureStatusOff = FeatureStatus("off")
)

type IncomingTrafficPolicy string

const (
	IncomingTrafficPolicyAllowAllTraffic          = IncomingTrafficPolicy("AllowAllTraffic")
	IncomingTrafficPolicyAllowVirtualNetworksOnly = IncomingTrafficPolicy("AllowVirtualNetworksOnly")
)

type InitialDownloadPolicy string

const (
	InitialDownloadPolicyNamespaceOnly              = InitialDownloadPolicy("NamespaceOnly")
	InitialDownloadPolicyNamespaceThenModifiedFiles = InitialDownloadPolicy("NamespaceThenModifiedFiles")
	InitialDownloadPolicyAvoidTieredFiles           = InitialDownloadPolicy("AvoidTieredFiles")
)

type LocalCacheMode string

const (
	LocalCacheModeDownloadNewAndModifiedFiles = LocalCacheMode("DownloadNewAndModifiedFiles")
	LocalCacheModeUpdateLocallyCachedFiles    = LocalCacheMode("UpdateLocallyCachedFiles")
)

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

func init() {
}

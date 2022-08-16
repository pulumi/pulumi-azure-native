


package v20180101preview

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending      = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved     = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected     = PrivateEndpointServiceConnectionStatus("Rejected")
	PrivateEndpointServiceConnectionStatusDisconnected = PrivateEndpointServiceConnectionStatus("Disconnected")
)

type SkuName string

const (
	SkuNameStandard = SkuName("Standard")
)

type SkuTier string

const (
	SkuTierStandard = SkuTier("Standard")
)

func init() {
}

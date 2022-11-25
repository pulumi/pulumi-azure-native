


package v20210601preview

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

type PublicNetworkAccessType string

const (
	// Enables public network connectivity to the Attestation Provider REST APIs.
	PublicNetworkAccessTypeEnabled = PublicNetworkAccessType("Enabled")
	// Disables public network connectivity to the Attestation Provider REST APIs.
	PublicNetworkAccessTypeDisabled = PublicNetworkAccessType("Disabled")
)

func init() {
}

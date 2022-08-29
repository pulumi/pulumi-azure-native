


package v20200202

type InfrastructureEncryptionState string

const (
	InfrastructureEncryptionStateEnabled  = InfrastructureEncryptionState("Enabled")
	InfrastructureEncryptionStateDisabled = InfrastructureEncryptionState("Disabled")
)

type PrivateEndpointConnectionStatus string

const (
	PrivateEndpointConnectionStatusPending      = PrivateEndpointConnectionStatus("Pending")
	PrivateEndpointConnectionStatusApproved     = PrivateEndpointConnectionStatus("Approved")
	PrivateEndpointConnectionStatusRejected     = PrivateEndpointConnectionStatus("Rejected")
	PrivateEndpointConnectionStatusDisconnected = PrivateEndpointConnectionStatus("Disconnected")
)

type ProvisioningState string

const (
	ProvisioningStateSucceeded = ProvisioningState("Succeeded")
	ProvisioningStateDeleting  = ProvisioningState("Deleting")
	ProvisioningStateFailed    = ProvisioningState("Failed")
	ProvisioningStatePending   = ProvisioningState("Pending")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
)

type SkuName string

const (
	SkuNameStandard = SkuName("Standard")
	SkuNameRS0      = SkuName("RS0")
)

func init() {
}

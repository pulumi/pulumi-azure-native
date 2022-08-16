


package iotcentral

type AppSku string

const (
	AppSkuST0 = AppSku("ST0")
	AppSkuST1 = AppSku("ST1")
	AppSkuST2 = AppSku("ST2")
)

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

type SystemAssignedServiceIdentityType string

const (
	SystemAssignedServiceIdentityTypeNone           = SystemAssignedServiceIdentityType("None")
	SystemAssignedServiceIdentityTypeSystemAssigned = SystemAssignedServiceIdentityType("SystemAssigned")
)

func init() {
}

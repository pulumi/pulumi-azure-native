


package v20191101

type ProvisioningStateType string

const (
	ProvisioningStateTypeSucceeded = ProvisioningStateType("Succeeded")
	ProvisioningStateTypeFailed    = ProvisioningStateType("Failed")
	ProvisioningStateTypeCancelled = ProvisioningStateType("Cancelled")
	ProvisioningStateTypeCreating  = ProvisioningStateType("Creating")
	ProvisioningStateTypeDeleting  = ProvisioningStateType("Deleting")
	ProvisioningStateTypeUpdating  = ProvisioningStateType("Updating")
)

type StorageTargetType string

const (
	StorageTargetTypeNfs3    = StorageTargetType("nfs3")
	StorageTargetTypeClfs    = StorageTargetType("clfs")
	StorageTargetTypeUnknown = StorageTargetType("unknown")
)

func init() {
}

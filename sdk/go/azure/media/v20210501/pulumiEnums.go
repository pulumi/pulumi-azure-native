


package v20210501

type AccountEncryptionKeyType string

const (
	// The Account Key is encrypted with a System Key.
	AccountEncryptionKeyTypeSystemKey = AccountEncryptionKeyType("SystemKey")
	// The Account Key is encrypted with a Customer Key.
	AccountEncryptionKeyTypeCustomerKey = AccountEncryptionKeyType("CustomerKey")
)

type DefaultAction string

const (
	// All public IP addresses are allowed.
	DefaultActionAllow = DefaultAction("Allow")
	// Public IP addresses are blocked.
	DefaultActionDeny = DefaultAction("Deny")
)

type ManagedIdentityType string

const (
	// A system-assigned managed identity.
	ManagedIdentityTypeSystemAssigned = ManagedIdentityType("SystemAssigned")
	// No managed identity.
	ManagedIdentityTypeNone = ManagedIdentityType("None")
)

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

type StorageAccountType string

const (
	// The primary storage account for the Media Services account.
	StorageAccountTypePrimary = StorageAccountType("Primary")
	// A secondary storage account for the Media Services account.
	StorageAccountTypeSecondary = StorageAccountType("Secondary")
)

type StorageAuthentication string

const (
	// System authentication.
	StorageAuthenticationSystem = StorageAuthentication("System")
	// Managed Identity authentication.
	StorageAuthenticationManagedIdentity = StorageAuthentication("ManagedIdentity")
)

func init() {
}

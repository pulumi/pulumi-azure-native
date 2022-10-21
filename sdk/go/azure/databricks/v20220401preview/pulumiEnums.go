


package v20220401preview

type EncryptionKeySource string

const (
	EncryptionKeySource_Microsoft_Keyvault = EncryptionKeySource("Microsoft.Keyvault")
)

type IdentityType string

const (
	IdentityTypeNone           = IdentityType("None")
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
)

type KeySource string

const (
	KeySourceDefault             = KeySource("Default")
	KeySource_Microsoft_Keyvault = KeySource("Microsoft.Keyvault")
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

type RequiredNsgRules string

const (
	RequiredNsgRulesAllRules               = RequiredNsgRules("AllRules")
	RequiredNsgRulesNoAzureDatabricksRules = RequiredNsgRules("NoAzureDatabricksRules")
	RequiredNsgRulesNoAzureServiceRules    = RequiredNsgRules("NoAzureServiceRules")
)

func init() {
}

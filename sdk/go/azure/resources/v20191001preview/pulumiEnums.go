


package v20191001preview

type CleanupOptions string

const (
	CleanupOptionsAlways       = CleanupOptions("Always")
	CleanupOptionsOnSuccess    = CleanupOptions("OnSuccess")
	CleanupOptionsOnExpiration = CleanupOptions("OnExpiration")
)

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeUserAssigned = ManagedServiceIdentityType("UserAssigned")
)

type ScriptType string

const (
	ScriptTypeAzurePowerShell = ScriptType("AzurePowerShell")
	ScriptTypeAzureCLI        = ScriptType("AzureCLI")
)

func init() {
}

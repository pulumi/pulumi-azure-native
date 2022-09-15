


package v20220501

type CreatedByType string

const (
	CreatedByTypeUser            = CreatedByType("User")
	CreatedByTypeApplication     = CreatedByType("Application")
	CreatedByTypeManagedIdentity = CreatedByType("ManagedIdentity")
	CreatedByTypeKey             = CreatedByType("Key")
)

type DiagnosticLevel string

const (
	DiagnosticLevelOff      = DiagnosticLevel("Off")
	DiagnosticLevelBasic    = DiagnosticLevel("Basic")
	DiagnosticLevelEnhanced = DiagnosticLevel("Enhanced")
)

type WindowsServerSubscription string

const (
	WindowsServerSubscriptionDisabled = WindowsServerSubscription("Disabled")
	WindowsServerSubscriptionEnabled  = WindowsServerSubscription("Enabled")
)

func init() {
}

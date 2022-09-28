


package v20220901

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

type SoftwareAssuranceIntent string

const (
	SoftwareAssuranceIntentEnable  = SoftwareAssuranceIntent("Enable")
	SoftwareAssuranceIntentDisable = SoftwareAssuranceIntent("Disable")
)

type SoftwareAssuranceStatus string

const (
	SoftwareAssuranceStatusEnabled  = SoftwareAssuranceStatus("Enabled")
	SoftwareAssuranceStatusDisabled = SoftwareAssuranceStatus("Disabled")
)

type WindowsServerSubscription string

const (
	WindowsServerSubscriptionDisabled = WindowsServerSubscription("Disabled")
	WindowsServerSubscriptionEnabled  = WindowsServerSubscription("Enabled")
)

func init() {
}

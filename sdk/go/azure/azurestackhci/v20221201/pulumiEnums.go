


package v20221201

type AvailabilityType string

const (
	AvailabilityTypeLocal  = AvailabilityType("Local")
	AvailabilityTypeOnline = AvailabilityType("Online")
	AvailabilityTypeNotify = AvailabilityType("Notify")
)

type DiagnosticLevel string

const (
	DiagnosticLevelOff      = DiagnosticLevel("Off")
	DiagnosticLevelBasic    = DiagnosticLevel("Basic")
	DiagnosticLevelEnhanced = DiagnosticLevel("Enhanced")
)

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned, UserAssigned")
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

type State string

const (
	StateHasPrerequisite                               = State("HasPrerequisite")
	StateObsolete                                      = State("Obsolete")
	StateReady                                         = State("Ready")
	StateNotApplicableBecauseAnotherUpdateIsInProgress = State("NotApplicableBecauseAnotherUpdateIsInProgress")
	StatePreparing                                     = State("Preparing")
	StateInstalling                                    = State("Installing")
	StateInstalled                                     = State("Installed")
	StatePreparationFailed                             = State("PreparationFailed")
	StateInstallationFailed                            = State("InstallationFailed")
	StateInvalid                                       = State("Invalid")
	StateRecalled                                      = State("Recalled")
	StateDownloading                                   = State("Downloading")
	StateDownloadFailed                                = State("DownloadFailed")
	StateHealthChecking                                = State("HealthChecking")
	StateHealthCheckFailed                             = State("HealthCheckFailed")
	StateReadyToInstall                                = State("ReadyToInstall")
	StateScanInProgress                                = State("ScanInProgress")
	StateScanFailed                                    = State("ScanFailed")
)

type UpdateRunPropertiesState string

const (
	UpdateRunPropertiesStateUnknown    = UpdateRunPropertiesState("Unknown")
	UpdateRunPropertiesStateSucceeded  = UpdateRunPropertiesState("Succeeded")
	UpdateRunPropertiesStateInProgress = UpdateRunPropertiesState("InProgress")
	UpdateRunPropertiesStateFailed     = UpdateRunPropertiesState("Failed")
)

type UpdateSummariesPropertiesState string

const (
	UpdateSummariesPropertiesStateUnknown               = UpdateSummariesPropertiesState("Unknown")
	UpdateSummariesPropertiesStateAppliedSuccessfully   = UpdateSummariesPropertiesState("AppliedSuccessfully")
	UpdateSummariesPropertiesStateUpdateAvailable       = UpdateSummariesPropertiesState("UpdateAvailable")
	UpdateSummariesPropertiesStateUpdateInProgress      = UpdateSummariesPropertiesState("UpdateInProgress")
	UpdateSummariesPropertiesStateUpdateFailed          = UpdateSummariesPropertiesState("UpdateFailed")
	UpdateSummariesPropertiesStateNeedsAttention        = UpdateSummariesPropertiesState("NeedsAttention")
	UpdateSummariesPropertiesStatePreparationInProgress = UpdateSummariesPropertiesState("PreparationInProgress")
	UpdateSummariesPropertiesStatePreparationFailed     = UpdateSummariesPropertiesState("PreparationFailed")
)

type WindowsServerSubscription string

const (
	WindowsServerSubscriptionDisabled = WindowsServerSubscription("Disabled")
	WindowsServerSubscriptionEnabled  = WindowsServerSubscription("Enabled")
)

func init() {
}

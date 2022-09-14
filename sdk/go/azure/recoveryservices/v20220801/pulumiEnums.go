


package v20220801

type AgentAutoUpdateStatus string

const (
	AgentAutoUpdateStatusDisabled = AgentAutoUpdateStatus("Disabled")
	AgentAutoUpdateStatusEnabled  = AgentAutoUpdateStatus("Enabled")
)

type AutomationAccountAuthenticationType string

const (
	AutomationAccountAuthenticationTypeRunAsAccount           = AutomationAccountAuthenticationType("RunAsAccount")
	AutomationAccountAuthenticationTypeSystemAssignedIdentity = AutomationAccountAuthenticationType("SystemAssignedIdentity")
)

type DiskAccountType string

const (
	DiskAccountType_Standard_LRS    = DiskAccountType("Standard_LRS")
	DiskAccountType_Premium_LRS     = DiskAccountType("Premium_LRS")
	DiskAccountType_StandardSSD_LRS = DiskAccountType("StandardSSD_LRS")
)

type ExtendedLocationType string

const (
	ExtendedLocationTypeEdgeZone = ExtendedLocationType("EdgeZone")
)

type FailoverDeploymentModel string

const (
	FailoverDeploymentModelNotApplicable   = FailoverDeploymentModel("NotApplicable")
	FailoverDeploymentModelClassic         = FailoverDeploymentModel("Classic")
	FailoverDeploymentModelResourceManager = FailoverDeploymentModel("ResourceManager")
)

type LicenseType string

const (
	LicenseTypeNotSpecified  = LicenseType("NotSpecified")
	LicenseTypeNoLicenseType = LicenseType("NoLicenseType")
	LicenseTypeWindowsServer = LicenseType("WindowsServer")
)

type PossibleOperationsDirections string

const (
	PossibleOperationsDirectionsPrimaryToRecovery = PossibleOperationsDirections("PrimaryToRecovery")
	PossibleOperationsDirectionsRecoveryToPrimary = PossibleOperationsDirections("RecoveryToPrimary")
)

type RecoveryPlanActionLocation string

const (
	RecoveryPlanActionLocationPrimary  = RecoveryPlanActionLocation("Primary")
	RecoveryPlanActionLocationRecovery = RecoveryPlanActionLocation("Recovery")
)

type RecoveryPlanGroupType string

const (
	RecoveryPlanGroupTypeShutdown = RecoveryPlanGroupType("Shutdown")
	RecoveryPlanGroupTypeBoot     = RecoveryPlanGroupType("Boot")
	RecoveryPlanGroupTypeFailover = RecoveryPlanGroupType("Failover")
)

type ReplicationProtectedItemOperation string

const (
	ReplicationProtectedItemOperationReverseReplicate    = ReplicationProtectedItemOperation("ReverseReplicate")
	ReplicationProtectedItemOperationCommit              = ReplicationProtectedItemOperation("Commit")
	ReplicationProtectedItemOperationPlannedFailover     = ReplicationProtectedItemOperation("PlannedFailover")
	ReplicationProtectedItemOperationUnplannedFailover   = ReplicationProtectedItemOperation("UnplannedFailover")
	ReplicationProtectedItemOperationDisableProtection   = ReplicationProtectedItemOperation("DisableProtection")
	ReplicationProtectedItemOperationTestFailover        = ReplicationProtectedItemOperation("TestFailover")
	ReplicationProtectedItemOperationTestFailoverCleanup = ReplicationProtectedItemOperation("TestFailoverCleanup")
	ReplicationProtectedItemOperationFailback            = ReplicationProtectedItemOperation("Failback")
	ReplicationProtectedItemOperationFinalizeFailback    = ReplicationProtectedItemOperation("FinalizeFailback")
	ReplicationProtectedItemOperationCancelFailover      = ReplicationProtectedItemOperation("CancelFailover")
	ReplicationProtectedItemOperationChangePit           = ReplicationProtectedItemOperation("ChangePit")
	ReplicationProtectedItemOperationRepairReplication   = ReplicationProtectedItemOperation("RepairReplication")
	ReplicationProtectedItemOperationSwitchProtection    = ReplicationProtectedItemOperation("SwitchProtection")
	ReplicationProtectedItemOperationCompleteMigration   = ReplicationProtectedItemOperation("CompleteMigration")
)

type SetMultiVmSyncStatus string

const (
	SetMultiVmSyncStatusEnable  = SetMultiVmSyncStatus("Enable")
	SetMultiVmSyncStatusDisable = SetMultiVmSyncStatus("Disable")
)

type SqlServerLicenseType string

const (
	SqlServerLicenseTypeNotSpecified  = SqlServerLicenseType("NotSpecified")
	SqlServerLicenseTypeNoLicenseType = SqlServerLicenseType("NoLicenseType")
	SqlServerLicenseTypePAYG          = SqlServerLicenseType("PAYG")
	SqlServerLicenseTypeAHUB          = SqlServerLicenseType("AHUB")
)

func init() {
}

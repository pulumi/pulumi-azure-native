


package v20180110

type AgentAutoUpdateStatus string

const (
	AgentAutoUpdateStatusDisabled = AgentAutoUpdateStatus("Disabled")
	AgentAutoUpdateStatusEnabled  = AgentAutoUpdateStatus("Enabled")
)

type DiskAccountType string

const (
	DiskAccountType_Standard_LRS    = DiskAccountType("Standard_LRS")
	DiskAccountType_Premium_LRS     = DiskAccountType("Premium_LRS")
	DiskAccountType_StandardSSD_LRS = DiskAccountType("StandardSSD_LRS")
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

func init() {
}




package v20200625

type ActionAfterReboot string

const (
	ActionAfterRebootContinueConfiguration = ActionAfterReboot("ContinueConfiguration")
	ActionAfterRebootStopConfiguration     = ActionAfterReboot("StopConfiguration")
)

type AssignmentType string

const (
	AssignmentTypeAudit                = AssignmentType("Audit")
	AssignmentTypeDeployAndAutoCorrect = AssignmentType("DeployAndAutoCorrect")
	AssignmentTypeApplyAndAutoCorrect  = AssignmentType("ApplyAndAutoCorrect")
	AssignmentTypeApplyAndMonitor      = AssignmentType("ApplyAndMonitor")
)

type ConfigurationMode string

const (
	ConfigurationModeApplyOnly           = ConfigurationMode("ApplyOnly")
	ConfigurationModeApplyAndMonitor     = ConfigurationMode("ApplyAndMonitor")
	ConfigurationModeApplyAndAutoCorrect = ConfigurationMode("ApplyAndAutoCorrect")
)

type Kind string

const (
	KindDSC = Kind("DSC")
)

func init() {
}

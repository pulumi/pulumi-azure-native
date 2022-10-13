


package v20220125

type AssignmentType string

const (
	AssignmentTypeAudit                = AssignmentType("Audit")
	AssignmentTypeDeployAndAutoCorrect = AssignmentType("DeployAndAutoCorrect")
	AssignmentTypeApplyAndAutoCorrect  = AssignmentType("ApplyAndAutoCorrect")
	AssignmentTypeApplyAndMonitor      = AssignmentType("ApplyAndMonitor")
)

type Kind string

const (
	KindDSC = Kind("DSC")
)

func init() {
}

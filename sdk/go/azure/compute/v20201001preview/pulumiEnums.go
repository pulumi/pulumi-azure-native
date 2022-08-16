


package v20201001preview

type CloudServiceUpgradeMode string

const (
	CloudServiceUpgradeModeAuto         = CloudServiceUpgradeMode("Auto")
	CloudServiceUpgradeModeManual       = CloudServiceUpgradeMode("Manual")
	CloudServiceUpgradeModeSimultaneous = CloudServiceUpgradeMode("Simultaneous")
)

func init() {
}

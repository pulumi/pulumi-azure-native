


package v20220904

type CloudServiceSlotType string

const (
	CloudServiceSlotTypeProduction = CloudServiceSlotType("Production")
	CloudServiceSlotTypeStaging    = CloudServiceSlotType("Staging")
)

type CloudServiceUpgradeMode string

const (
	CloudServiceUpgradeModeAuto         = CloudServiceUpgradeMode("Auto")
	CloudServiceUpgradeModeManual       = CloudServiceUpgradeMode("Manual")
	CloudServiceUpgradeModeSimultaneous = CloudServiceUpgradeMode("Simultaneous")
)

func init() {
}

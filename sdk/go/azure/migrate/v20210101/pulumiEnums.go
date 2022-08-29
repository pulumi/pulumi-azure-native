


package v20210101

type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned   = ResourceIdentityType("UserAssigned")
)

type TargetAvailabilityZone string

const (
	TargetAvailabilityZoneOne   = TargetAvailabilityZone("1")
	TargetAvailabilityZoneTwo   = TargetAvailabilityZone("2")
	TargetAvailabilityZoneThree = TargetAvailabilityZone("3")
	TargetAvailabilityZoneNA    = TargetAvailabilityZone("NA")
)

type ZoneRedundant string

const (
	ZoneRedundantEnable  = ZoneRedundant("Enable")
	ZoneRedundantDisable = ZoneRedundant("Disable")
)

func init() {
}

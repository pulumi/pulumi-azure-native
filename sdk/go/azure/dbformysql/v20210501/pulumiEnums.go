


package v20210501

type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
	CreateModeReplica            = CreateMode("Replica")
	CreateModeGeoRestore         = CreateMode("GeoRestore")
)

type EnableStatusEnum string

const (
	EnableStatusEnumEnabled  = EnableStatusEnum("Enabled")
	EnableStatusEnumDisabled = EnableStatusEnum("Disabled")
)

type HighAvailabilityMode string

const (
	HighAvailabilityModeDisabled      = HighAvailabilityMode("Disabled")
	HighAvailabilityModeZoneRedundant = HighAvailabilityMode("ZoneRedundant")
	HighAvailabilityModeSameZone      = HighAvailabilityMode("SameZone")
)

type ReplicationRole string

const (
	ReplicationRoleNone    = ReplicationRole("None")
	ReplicationRoleSource  = ReplicationRole("Source")
	ReplicationRoleReplica = ReplicationRole("Replica")
)

type ServerVersion string

const (
	ServerVersion_5_7    = ServerVersion("5.7")
	ServerVersion_8_0_21 = ServerVersion("8.0.21")
)

type SkuTier string

const (
	SkuTierBurstable       = SkuTier("Burstable")
	SkuTierGeneralPurpose  = SkuTier("GeneralPurpose")
	SkuTierMemoryOptimized = SkuTier("MemoryOptimized")
)

func init() {
}

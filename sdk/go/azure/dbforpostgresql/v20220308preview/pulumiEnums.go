


package v20220308preview

type ArmServerKeyType string

const (
	ArmServerKeyTypeSystemAssigned = ArmServerKeyType("SystemAssigned")
	ArmServerKeyTypeAzureKeyVault  = ArmServerKeyType("AzureKeyVault")
)

type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModeCreate             = CreateMode("Create")
	CreateModeUpdate             = CreateMode("Update")
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
	CreateModeGeoRestore         = CreateMode("GeoRestore")
	CreateModeReplica            = CreateMode("Replica")
)

type GeoRedundantBackupEnum string

const (
	GeoRedundantBackupEnumEnabled  = GeoRedundantBackupEnum("Enabled")
	GeoRedundantBackupEnumDisabled = GeoRedundantBackupEnum("Disabled")
)

type HighAvailabilityMode string

const (
	HighAvailabilityModeDisabled      = HighAvailabilityMode("Disabled")
	HighAvailabilityModeZoneRedundant = HighAvailabilityMode("ZoneRedundant")
	HighAvailabilityModeSameZone      = HighAvailabilityMode("SameZone")
)

type IdentityType string

const (
	IdentityTypeNone           = IdentityType("None")
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned   = IdentityType("UserAssigned")
)

type PrincipalType string

const (
	PrincipalTypeUnknown          = PrincipalType("Unknown")
	PrincipalTypeUser             = PrincipalType("User")
	PrincipalTypeGroup            = PrincipalType("Group")
	PrincipalTypeServicePrincipal = PrincipalType("ServicePrincipal")
)

type ReplicationRole string

const (
	ReplicationRoleNone            = ReplicationRole("None")
	ReplicationRolePrimary         = ReplicationRole("Primary")
	ReplicationRoleSecondary       = ReplicationRole("Secondary")
	ReplicationRoleWalReplica      = ReplicationRole("WalReplica")
	ReplicationRoleSyncReplica     = ReplicationRole("SyncReplica")
	ReplicationRoleAsyncReplica    = ReplicationRole("AsyncReplica")
	ReplicationRoleGeoSyncReplica  = ReplicationRole("GeoSyncReplica")
	ReplicationRoleGeoAsyncReplica = ReplicationRole("GeoAsyncReplica")
)

type ServerVersion string

const (
	ServerVersion_14 = ServerVersion("14")
	ServerVersion_13 = ServerVersion("13")
	ServerVersion_12 = ServerVersion("12")
	ServerVersion_11 = ServerVersion("11")
)

type SkuTier string

const (
	SkuTierBurstable       = SkuTier("Burstable")
	SkuTierGeneralPurpose  = SkuTier("GeneralPurpose")
	SkuTierMemoryOptimized = SkuTier("MemoryOptimized")
)

func init() {
}

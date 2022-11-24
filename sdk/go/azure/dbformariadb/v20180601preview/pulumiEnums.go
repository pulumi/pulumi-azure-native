


package v20180601preview

type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
	CreateModeGeoRestore         = CreateMode("GeoRestore")
	CreateModeReplica            = CreateMode("Replica")
)

type GeoRedundantBackup string

const (
	GeoRedundantBackupEnabled  = GeoRedundantBackup("Enabled")
	GeoRedundantBackupDisabled = GeoRedundantBackup("Disabled")
)

type MinimalTlsVersionEnum string

const (
	MinimalTlsVersionEnum_TLS1_0                = MinimalTlsVersionEnum("TLS1_0")
	MinimalTlsVersionEnum_TLS1_1                = MinimalTlsVersionEnum("TLS1_1")
	MinimalTlsVersionEnum_TLS1_2                = MinimalTlsVersionEnum("TLS1_2")
	MinimalTlsVersionEnumTLSEnforcementDisabled = MinimalTlsVersionEnum("TLSEnforcementDisabled")
)

type ServerVersion string

const (
	ServerVersion_10_2 = ServerVersion("10.2")
	ServerVersion_10_3 = ServerVersion("10.3")
)

type SkuTier string

const (
	SkuTierBasic           = SkuTier("Basic")
	SkuTierGeneralPurpose  = SkuTier("GeneralPurpose")
	SkuTierMemoryOptimized = SkuTier("MemoryOptimized")
)

type SslEnforcementEnum string

const (
	SslEnforcementEnumEnabled  = SslEnforcementEnum("Enabled")
	SslEnforcementEnumDisabled = SslEnforcementEnum("Disabled")
)

type StorageAutogrow string

const (
	StorageAutogrowEnabled  = StorageAutogrow("Enabled")
	StorageAutogrowDisabled = StorageAutogrow("Disabled")
)

func init() {
}

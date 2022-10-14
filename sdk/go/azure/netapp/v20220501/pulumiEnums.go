


package v20220501

type ApplicationType string

const (
	ApplicationType_SAP_HANA = ApplicationType("SAP-HANA")
)

type AvsDataStore string

const (
	// avsDataStore is enabled
	AvsDataStoreEnabled = AvsDataStore("Enabled")
	// avsDataStore is disabled
	AvsDataStoreDisabled = AvsDataStore("Disabled")
)

type ChownMode string

const (
	ChownModeRestricted   = ChownMode("Restricted")
	ChownModeUnrestricted = ChownMode("Unrestricted")
)

type EnableSubvolumes string

const (
	// subvolumes are enabled
	EnableSubvolumesEnabled = EnableSubvolumes("Enabled")
	// subvolumes are not enabled
	EnableSubvolumesDisabled = EnableSubvolumes("Disabled")
)

type EncryptionKeySource string

const (
	// Microsoft-managed key encryption
	EncryptionKeySource_Microsoft_NetApp = EncryptionKeySource("Microsoft.NetApp")
	// Customer-managed key encryption
	EncryptionKeySource_Microsoft_KeyVault = EncryptionKeySource("Microsoft.KeyVault")
)

type EncryptionType string

const (
	// EncryptionType Single, volumes will use single encryption at rest
	EncryptionTypeSingle = EncryptionType("Single")
	// EncryptionType Double, volumes will use double encryption at rest
	EncryptionTypeDouble = EncryptionType("Double")
)

type EndpointType string

const (
	EndpointTypeSrc = EndpointType("src")
	EndpointTypeDst = EndpointType("dst")
)

type IdentityType string

const (
	IdentityTypeNone                         = IdentityType("None")
	IdentityTypeSystemAssigned               = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned                 = IdentityType("UserAssigned")
	IdentityType_SystemAssigned_UserAssigned = IdentityType("SystemAssigned,UserAssigned")
)

type KeySource string

const (
	// Microsoft-managed key encryption
	KeySource_Microsoft_NetApp = KeySource("Microsoft.NetApp")
	// Customer-managed key encryption
	KeySource_Microsoft_KeyVault = KeySource("Microsoft.KeyVault")
)

type NetworkFeatures string

const (
	// Basic network feature.
	NetworkFeaturesBasic = NetworkFeatures("Basic")
	// Standard network feature.
	NetworkFeaturesStandard = NetworkFeatures("Standard")
)

type QosType string

const (
	// qos type Auto
	QosTypeAuto = QosType("Auto")
	// qos type Manual
	QosTypeManual = QosType("Manual")
)

type ReplicationSchedule string

const (
	ReplicationSchedule_10minutely = ReplicationSchedule("_10minutely")
	ReplicationScheduleHourly      = ReplicationSchedule("hourly")
	ReplicationScheduleDaily       = ReplicationSchedule("daily")
)

type SecurityStyle string

const (
	SecurityStyleNtfs = SecurityStyle("ntfs")
	SecurityStyleUnix = SecurityStyle("unix")
)

type ServiceLevel string

const (
	// Standard service level
	ServiceLevelStandard = ServiceLevel("Standard")
	// Premium service level
	ServiceLevelPremium = ServiceLevel("Premium")
	// Ultra service level
	ServiceLevelUltra = ServiceLevel("Ultra")
	// Zone redundant storage service level
	ServiceLevelStandardZRS = ServiceLevel("StandardZRS")
)

type SmbAccessBasedEnumeration string

const (
	// smbAccessBasedEnumeration share setting is disabled
	SmbAccessBasedEnumerationDisabled = SmbAccessBasedEnumeration("Disabled")
	// smbAccessBasedEnumeration share setting is enabled
	SmbAccessBasedEnumerationEnabled = SmbAccessBasedEnumeration("Enabled")
)

type SmbNonBrowsable string

const (
	// smbNonBrowsable share setting is disabled
	SmbNonBrowsableDisabled = SmbNonBrowsable("Disabled")
	// smbNonBrowsable share setting is enabled
	SmbNonBrowsableEnabled = SmbNonBrowsable("Enabled")
)

type Type string

const (
	// Default user quota
	TypeDefaultUserQuota = Type("DefaultUserQuota")
	// Default group quota
	TypeDefaultGroupQuota = Type("DefaultGroupQuota")
	// Individual user quota
	TypeIndividualUserQuota = Type("IndividualUserQuota")
	// Individual group quota
	TypeIndividualGroupQuota = Type("IndividualGroupQuota")
)

func init() {
}

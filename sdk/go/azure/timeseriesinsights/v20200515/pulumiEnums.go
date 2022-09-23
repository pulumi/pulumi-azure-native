


package v20200515

type AccessPolicyRole string

const (
	AccessPolicyRoleReader      = AccessPolicyRole("Reader")
	AccessPolicyRoleContributor = AccessPolicyRole("Contributor")
)

type DataStringComparisonBehavior string

const (
	DataStringComparisonBehaviorOrdinal           = DataStringComparisonBehavior("Ordinal")
	DataStringComparisonBehaviorOrdinalIgnoreCase = DataStringComparisonBehavior("OrdinalIgnoreCase")
)

type EnvironmentKind string

const (
	EnvironmentKindGen1 = EnvironmentKind("Gen1")
	EnvironmentKindGen2 = EnvironmentKind("Gen2")
)

type EventSourceKind string

const (
	EventSourceKind_Microsoft_EventHub = EventSourceKind("Microsoft.EventHub")
	EventSourceKind_Microsoft_IoTHub   = EventSourceKind("Microsoft.IoTHub")
)

type IngressStartAtType string

const (
	IngressStartAtTypeEarliestAvailable       = IngressStartAtType("EarliestAvailable")
	IngressStartAtTypeEventSourceCreationTime = IngressStartAtType("EventSourceCreationTime")
	IngressStartAtTypeCustomEnqueuedTime      = IngressStartAtType("CustomEnqueuedTime")
)

type LocalTimestampFormat string

const (
	LocalTimestampFormatEmbedded = LocalTimestampFormat("Embedded")
)

type PropertyType string

const (
	PropertyTypeString = PropertyType("String")
)

type ReferenceDataKeyPropertyType string

const (
	ReferenceDataKeyPropertyTypeString   = ReferenceDataKeyPropertyType("String")
	ReferenceDataKeyPropertyTypeDouble   = ReferenceDataKeyPropertyType("Double")
	ReferenceDataKeyPropertyTypeBool     = ReferenceDataKeyPropertyType("Bool")
	ReferenceDataKeyPropertyTypeDateTime = ReferenceDataKeyPropertyType("DateTime")
)

type SkuName string

const (
	SkuNameS1 = SkuName("S1")
	SkuNameS2 = SkuName("S2")
	SkuNameP1 = SkuName("P1")
	SkuNameL1 = SkuName("L1")
)

type StorageLimitExceededBehavior string

const (
	StorageLimitExceededBehaviorPurgeOldData = StorageLimitExceededBehavior("PurgeOldData")
	StorageLimitExceededBehaviorPauseIngress = StorageLimitExceededBehavior("PauseIngress")
)

func init() {
}

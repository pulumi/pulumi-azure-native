


package appplatform

type ManagedIdentityType string

const (
	ManagedIdentityTypeNone                         = ManagedIdentityType("None")
	ManagedIdentityTypeSystemAssigned               = ManagedIdentityType("SystemAssigned")
	ManagedIdentityTypeUserAssigned                 = ManagedIdentityType("UserAssigned")
	ManagedIdentityType_SystemAssigned_UserAssigned = ManagedIdentityType("SystemAssigned,UserAssigned")
)

type RuntimeVersion string

const (
	RuntimeVersion_Java_8     = RuntimeVersion("Java_8")
	RuntimeVersion_Java_11    = RuntimeVersion("Java_11")
	RuntimeVersion_NetCore_31 = RuntimeVersion("NetCore_31")
)

type UserSourceType string

const (
	UserSourceTypeJar        = UserSourceType("Jar")
	UserSourceTypeNetCoreZip = UserSourceType("NetCoreZip")
	UserSourceTypeSource     = UserSourceType("Source")
)

func init() {
}

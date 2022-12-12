


package v20210901preview

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

type StorageType string

const (
	StorageTypeStorageAccount = StorageType("StorageAccount")
)

type Type string

const (
	TypeAzureFileVolume = Type("AzureFileVolume")
)

type UserSourceType string

const (
	UserSourceTypeJar        = UserSourceType("Jar")
	UserSourceTypeNetCoreZip = UserSourceType("NetCoreZip")
	UserSourceTypeSource     = UserSourceType("Source")
	UserSourceTypeContainer  = UserSourceType("Container")
)

func init() {
}

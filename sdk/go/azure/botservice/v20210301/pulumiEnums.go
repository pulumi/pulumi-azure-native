


package v20210301

type Kind string

const (
	KindSdk      = Kind("sdk")
	KindDesigner = Kind("designer")
	KindBot      = Kind("bot")
	KindFunction = Kind("function")
	KindAzurebot = Kind("azurebot")
)

type MsaAppType string

const (
	MsaAppTypeUserAssignedMSI = MsaAppType("UserAssignedMSI")
	MsaAppTypeSingleTenant    = MsaAppType("SingleTenant")
	MsaAppTypeMultiTenant     = MsaAppType("MultiTenant")
)

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

type SkuName string

const (
	SkuNameF0 = SkuName("F0")
	SkuNameS1 = SkuName("S1")
)

func init() {
}

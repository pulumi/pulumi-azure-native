


package v20220415preview

type Distro string

const (
	DistroAKSEdge = Distro("AKSEdge")
)

type Provider string

const (
	ProviderVMWare    = Provider("VMWare")
	ProviderHCI       = Provider("HCI")
	ProviderSCVMM     = Provider("SCVMM")
	ProviderKubeVirt  = Provider("KubeVirt")
	ProviderOpenStack = Provider("OpenStack")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
)

func init() {
}

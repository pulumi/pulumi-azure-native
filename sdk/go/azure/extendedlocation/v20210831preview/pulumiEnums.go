


package v20210831preview

type HostType string

const (
	HostTypeKubernetes = HostType("Kubernetes")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
)

func init() {
}




package loadtestservice

type SystemAssignedServiceIdentityType string

const (
	SystemAssignedServiceIdentityTypeNone           = SystemAssignedServiceIdentityType("None")
	SystemAssignedServiceIdentityTypeSystemAssigned = SystemAssignedServiceIdentityType("SystemAssigned")
)

func init() {
}

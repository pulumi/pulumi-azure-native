


package v20200804preview

type HealthAlertsNamespace string

const (
	HealthAlertsNamespaceGuestVmHealth = HealthAlertsNamespace("GuestVmHealth")
)

type HealthStateName string

const (
	HealthStateNameWarning  = HealthStateName("Warning")
	HealthStateNameCritical = HealthStateName("Critical")
)

func init() {
}

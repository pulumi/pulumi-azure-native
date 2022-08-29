


package v20211001preview

type DomainManagement string

const (
	DomainManagementAzureManaged                    = DomainManagement("AzureManaged")
	DomainManagementCustomerManaged                 = DomainManagement("CustomerManaged")
	DomainManagementCustomerManagedInExchangeOnline = DomainManagement("CustomerManagedInExchangeOnline")
)

type UserEngagementTracking string

const (
	UserEngagementTrackingDisabled = UserEngagementTracking("Disabled")
	UserEngagementTrackingEnabled  = UserEngagementTracking("Enabled")
)

func init() {
}

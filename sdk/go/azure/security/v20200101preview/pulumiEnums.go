


package v20200101preview

type AuthenticationType string

const (
	// AWS cloud account connector user credentials authentication
	AuthenticationTypeAwsCreds = AuthenticationType("awsCreds")
	// AWS account connector assume role authentication
	AuthenticationTypeAwsAssumeRole = AuthenticationType("awsAssumeRole")
	// GCP account connector service to service authentication
	AuthenticationTypeGcpCredentials = AuthenticationType("gcpCredentials")
)

type AutoProvision string

const (
	// Install missing Azure Arc agents on machines automatically
	AutoProvisionOn = AutoProvision("On")
	// Do not install Azure Arc agent on the machines automatically
	AutoProvisionOff = AutoProvision("Off")
)

type MinimalSeverity string

const (
	// Get notifications on new alerts with High severity
	MinimalSeverityHigh = MinimalSeverity("High")
	// Get notifications on new alerts with medium or high severity
	MinimalSeverityMedium = MinimalSeverity("Medium")
	// Don't get notifications on new alerts with low, medium or high severity
	MinimalSeverityLow = MinimalSeverity("Low")
)

type Roles string

const (
	// If enabled, send notification on new alerts to the account admins
	RolesAccountAdmin = Roles("AccountAdmin")
	// If enabled, send notification on new alerts to the service admins
	RolesServiceAdmin = Roles("ServiceAdmin")
	// If enabled, send notification on new alerts to the subscription owners
	RolesOwner = Roles("Owner")
	// If enabled, send notification on new alerts to the subscription contributors
	RolesContributor = Roles("Contributor")
)

type State string

const (
	// Send notification on new alerts to the subscription's admins
	StateOn = State("On")
	// Don't send notification on new alerts to the subscription's admins
	StateOff = State("Off")
)

func init() {
}




package v20220801preview

type CloudName string

const (
	CloudNameAzure       = CloudName("Azure")
	CloudNameAWS         = CloudName("AWS")
	CloudNameGCP         = CloudName("GCP")
	CloudNameGithub      = CloudName("Github")
	CloudNameAzureDevOps = CloudName("AzureDevOps")
)

type EnvironmentType string

const (
	EnvironmentTypeAwsAccount       = EnvironmentType("AwsAccount")
	EnvironmentTypeGcpProject       = EnvironmentType("GcpProject")
	EnvironmentTypeGithubScope      = EnvironmentType("GithubScope")
	EnvironmentTypeAzureDevOpsScope = EnvironmentType("AzureDevOpsScope")
)

type OfferingType string

const (
	OfferingTypeCspmMonitorAws               = OfferingType("CspmMonitorAws")
	OfferingTypeDefenderForContainersAws     = OfferingType("DefenderForContainersAws")
	OfferingTypeDefenderForServersAws        = OfferingType("DefenderForServersAws")
	OfferingTypeDefenderForDatabasesAws      = OfferingType("DefenderForDatabasesAws")
	OfferingTypeInformationProtectionAws     = OfferingType("InformationProtectionAws")
	OfferingTypeCspmMonitorGcp               = OfferingType("CspmMonitorGcp")
	OfferingTypeCspmMonitorGithub            = OfferingType("CspmMonitorGithub")
	OfferingTypeCspmMonitorAzureDevOps       = OfferingType("CspmMonitorAzureDevOps")
	OfferingTypeDefenderForServersGcp        = OfferingType("DefenderForServersGcp")
	OfferingTypeDefenderForContainersGcp     = OfferingType("DefenderForContainersGcp")
	OfferingTypeDefenderForDatabasesGcp      = OfferingType("DefenderForDatabasesGcp")
	OfferingTypeDefenderCspmAws              = OfferingType("DefenderCspmAws")
	OfferingTypeDefenderCspmGcp              = OfferingType("DefenderCspmGcp")
	OfferingTypeDefenderForDevOpsGithub      = OfferingType("DefenderForDevOpsGithub")
	OfferingTypeDefenderForDevOpsAzureDevOps = OfferingType("DefenderForDevOpsAzureDevOps")
)

type OrganizationMembershipType string

const (
	OrganizationMembershipTypeMember       = OrganizationMembershipType("Member")
	OrganizationMembershipTypeOrganization = OrganizationMembershipType("Organization")
)

type ScanningMode string

const (
	ScanningModeDefault = ScanningMode("Default")
)

type SubPlan string

const (
	SubPlanP1 = SubPlan("P1")
	SubPlanP2 = SubPlan("P2")
)

type Type string

const (
	TypeQualys = Type("Qualys")
	TypeTVM    = Type("TVM")
)

func init() {
}

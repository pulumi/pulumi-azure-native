


package v20211201preview

type CloudName string

const (
	CloudNameAzure  = CloudName("Azure")
	CloudNameAWS    = CloudName("AWS")
	CloudNameGCP    = CloudName("GCP")
	CloudNameGithub = CloudName("Github")
)

type EnvironmentType string

const (
	EnvironmentTypeAwsAccount  = EnvironmentType("AwsAccount")
	EnvironmentTypeGcpProject  = EnvironmentType("GcpProject")
	EnvironmentTypeGithubScope = EnvironmentType("GithubScope")
)

type OfferingType string

const (
	OfferingTypeCspmMonitorAws           = OfferingType("CspmMonitorAws")
	OfferingTypeDefenderForContainersAws = OfferingType("DefenderForContainersAws")
	OfferingTypeDefenderForServersAws    = OfferingType("DefenderForServersAws")
	OfferingTypeInformationProtectionAws = OfferingType("InformationProtectionAws")
	OfferingTypeCspmMonitorGcp           = OfferingType("CspmMonitorGcp")
	OfferingTypeCspmMonitorGithub        = OfferingType("CspmMonitorGithub")
	OfferingTypeDefenderForServersGcp    = OfferingType("DefenderForServersGcp")
	OfferingTypeDefenderForContainersGcp = OfferingType("DefenderForContainersGcp")
)

type OrganizationMembershipType string

const (
	OrganizationMembershipTypeMember       = OrganizationMembershipType("Member")
	OrganizationMembershipTypeOrganization = OrganizationMembershipType("Organization")
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

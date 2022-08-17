


package v20210701preview

type CloudName string

const (
	CloudNameAzure = CloudName("Azure")
	CloudNameAWS   = CloudName("AWS")
	CloudNameGCP   = CloudName("GCP")
)

type OfferingType string

const (
	OfferingTypeCspmMonitorAws           = OfferingType("CspmMonitorAws")
	OfferingTypeDefenderForContainersAws = OfferingType("DefenderForContainersAws")
	OfferingTypeDefenderForServersAws    = OfferingType("DefenderForServersAws")
	OfferingTypeInformationProtectionAws = OfferingType("InformationProtectionAws")
)

type OrganizationMembershipType string

const (
	OrganizationMembershipTypeMember       = OrganizationMembershipType("Member")
	OrganizationMembershipTypeOrganization = OrganizationMembershipType("Organization")
)

type SeverityEnum string

const (
	SeverityEnumHigh   = SeverityEnum("High")
	SeverityEnumMedium = SeverityEnum("Medium")
	SeverityEnumLow    = SeverityEnum("Low")
)

type SupportedCloudEnum string

const (
	SupportedCloudEnumAWS = SupportedCloudEnum("AWS")
	SupportedCloudEnumGCP = SupportedCloudEnum("GCP")
)

func init() {
}




package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionType string

const (
	ActionTypeLogicApp  = ActionType("LogicApp")
	ActionTypeEventHub  = ActionType("EventHub")
	ActionTypeWorkspace = ActionType("Workspace")
)

type AdditionalWorkspaceDataType string

const (
	AdditionalWorkspaceDataTypeAlerts    = AdditionalWorkspaceDataType("Alerts")
	AdditionalWorkspaceDataTypeRawEvents = AdditionalWorkspaceDataType("RawEvents")
)

type AdditionalWorkspaceType string

const (
	AdditionalWorkspaceTypeSentinel = AdditionalWorkspaceType("Sentinel")
)

type ApplicationSourceResourceType string

const (
	// The source of the application is assessments
	ApplicationSourceResourceTypeAssessments = ApplicationSourceResourceType("Assessments")
)

type AssessmentStatusCode string

const (
	// The resource is healthy
	AssessmentStatusCodeHealthy = AssessmentStatusCode("Healthy")
	// The resource has a security issue that needs to be addressed
	AssessmentStatusCodeUnhealthy = AssessmentStatusCode("Unhealthy")
	// Assessment for this resource did not happen
	AssessmentStatusCodeNotApplicable = AssessmentStatusCode("NotApplicable")
)

type AssessmentType string

const (
	// Microsoft Defender for Cloud managed assessments
	AssessmentTypeBuiltIn = AssessmentType("BuiltIn")
	// User defined policies that are automatically ingested from Azure Policy to Microsoft Defender for Cloud
	AssessmentTypeCustomPolicy = AssessmentType("CustomPolicy")
	// User assessments pushed directly by the user or other third party to Microsoft Defender for Cloud
	AssessmentTypeCustomerManaged = AssessmentType("CustomerManaged")
	// An assessment that was created by a verified 3rd party if the user connected it to ASC
	AssessmentTypeVerifiedPartner = AssessmentType("VerifiedPartner")
)

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

type Categories string

const (
	CategoriesCompute           = Categories("Compute")
	CategoriesNetworking        = Categories("Networking")
	CategoriesData              = Categories("Data")
	CategoriesIdentityAndAccess = Categories("IdentityAndAccess")
	CategoriesIoT               = Categories("IoT")
)

type CloudName string

const (
	CloudNameAzure = CloudName("Azure")
	CloudNameAWS   = CloudName("AWS")
	CloudNameGCP   = CloudName("GCP")
)

type DataSource string

const (
	// Devices twin data
	DataSourceTwinData = DataSource("TwinData")
)

type EventSource string

const (
	EventSourceAssessments                            = EventSource("Assessments")
	EventSourceAssessmentsSnapshot                    = EventSource("AssessmentsSnapshot")
	EventSourceSubAssessments                         = EventSource("SubAssessments")
	EventSourceSubAssessmentsSnapshot                 = EventSource("SubAssessmentsSnapshot")
	EventSourceAlerts                                 = EventSource("Alerts")
	EventSourceSecureScores                           = EventSource("SecureScores")
	EventSourceSecureScoresSnapshot                   = EventSource("SecureScoresSnapshot")
	EventSourceSecureScoreControls                    = EventSource("SecureScoreControls")
	EventSourceSecureScoreControlsSnapshot            = EventSource("SecureScoreControlsSnapshot")
	EventSourceRegulatoryComplianceAssessment         = EventSource("RegulatoryComplianceAssessment")
	EventSourceRegulatoryComplianceAssessmentSnapshot = EventSource("RegulatoryComplianceAssessmentSnapshot")
)

type ExportData string

const (
	// Agent raw events
	ExportDataRawEvents = ExportData("RawEvents")
)

type ImplementationEffort string

const (
	ImplementationEffortLow      = ImplementationEffort("Low")
	ImplementationEffortModerate = ImplementationEffort("Moderate")
	ImplementationEffortHigh     = ImplementationEffort("High")
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

type OfferingType string

const (
	OfferingTypeCspmMonitorAws           = OfferingType("CspmMonitorAws")
	OfferingTypeDefenderForContainersAws = OfferingType("DefenderForContainersAws")
	OfferingTypeDefenderForServersAws    = OfferingType("DefenderForServersAws")
	OfferingTypeInformationProtectionAws = OfferingType("InformationProtectionAws")
)

type Operator string

const (
	// Applies for decimal and non-decimal operands
	OperatorEquals = Operator("Equals")
	// Applies only for decimal operands
	OperatorGreaterThan = Operator("GreaterThan")
	// Applies only for decimal operands
	OperatorGreaterThanOrEqualTo = Operator("GreaterThanOrEqualTo")
	// Applies only for decimal operands
	OperatorLesserThan = Operator("LesserThan")
	// Applies only for decimal operands
	OperatorLesserThanOrEqualTo = Operator("LesserThanOrEqualTo")
	// Applies  for decimal and non-decimal operands
	OperatorNotEquals = Operator("NotEquals")
	// Applies only for non-decimal operands
	OperatorContains = Operator("Contains")
	// Applies only for non-decimal operands
	OperatorStartsWith = Operator("StartsWith")
	// Applies only for non-decimal operands
	OperatorEndsWith = Operator("EndsWith")
)

type OrganizationMembershipType string

const (
	OrganizationMembershipTypeMember       = OrganizationMembershipType("Member")
	OrganizationMembershipTypeOrganization = OrganizationMembershipType("Organization")
)

type PropertyType string

const (
	PropertyTypeString  = PropertyType("String")
	PropertyTypeInteger = PropertyType("Integer")
	PropertyTypeNumber  = PropertyType("Number")
	PropertyTypeBoolean = PropertyType("Boolean")
)

type Protocol string

const (
	ProtocolTCP = Protocol("TCP")
	ProtocolUDP = Protocol("UDP")
	ProtocolAll = Protocol("*")
)

type RecommendationConfigStatus string

const (
	RecommendationConfigStatusDisabled = RecommendationConfigStatus("Disabled")
	RecommendationConfigStatusEnabled  = RecommendationConfigStatus("Enabled")
)

type RecommendationType string

const (
	// Authentication schema used for pull an edge module from an ACR repository does not use Service Principal Authentication.
	RecommendationType_IoT_ACRAuthentication = RecommendationType("IoT_ACRAuthentication")
	// IoT agent message size capacity is currently underutilized, causing an increase in the number of sent messages. Adjust message intervals for better utilization.
	RecommendationType_IoT_AgentSendsUnutilizedMessages = RecommendationType("IoT_AgentSendsUnutilizedMessages")
	// Identified security related system configuration issues.
	RecommendationType_IoT_Baseline = RecommendationType("IoT_Baseline")
	// You can optimize Edge Hub memory usage by turning off protocol heads for any protocols not used by Edge modules in your solution.
	RecommendationType_IoT_EdgeHubMemOptimize = RecommendationType("IoT_EdgeHubMemOptimize")
	// Logging is disabled for this edge module.
	RecommendationType_IoT_EdgeLoggingOptions = RecommendationType("IoT_EdgeLoggingOptions")
	// A minority within a device security group has inconsistent Edge Module settings with the rest of their group.
	RecommendationType_IoT_InconsistentModuleSettings = RecommendationType("IoT_InconsistentModuleSettings")
	// Install the Azure Security of Things Agent.
	RecommendationType_IoT_InstallAgent = RecommendationType("IoT_InstallAgent")
	// IP Filter Configuration should have rules defined for allowed traffic and should deny all other traffic by default.
	RecommendationType_IoT_IPFilter_DenyAll = RecommendationType("IoT_IPFilter_DenyAll")
	// An Allow IP Filter rules source IP range is too large. Overly permissive rules might expose your IoT hub to malicious intenders.
	RecommendationType_IoT_IPFilter_PermissiveRule = RecommendationType("IoT_IPFilter_PermissiveRule")
	// A listening endpoint was found on the device.
	RecommendationType_IoT_OpenPorts = RecommendationType("IoT_OpenPorts")
	// An Allowed firewall policy was found (INPUT/OUTPUT). The policy should Deny all traffic by default and define rules to allow necessary communication to/from the device.
	RecommendationType_IoT_PermissiveFirewallPolicy = RecommendationType("IoT_PermissiveFirewallPolicy")
	// A rule in the firewall has been found that contains a permissive pattern for a wide range of IP addresses or Ports.
	RecommendationType_IoT_PermissiveInputFirewallRules = RecommendationType("IoT_PermissiveInputFirewallRules")
	// A rule in the firewall has been found that contains a permissive pattern for a wide range of IP addresses or Ports.
	RecommendationType_IoT_PermissiveOutputFirewallRules = RecommendationType("IoT_PermissiveOutputFirewallRules")
	// Edge module is configured to run in privileged mode, with extensive Linux capabilities or with host-level network access (send/receive data to host machine).
	RecommendationType_IoT_PrivilegedDockerOptions = RecommendationType("IoT_PrivilegedDockerOptions")
	// Same authentication credentials to the IoT Hub used by multiple devices. This could indicate an illegitimate device impersonating a legitimate device. It also exposes the risk of device impersonation by an attacker.
	RecommendationType_IoT_SharedCredentials = RecommendationType("IoT_SharedCredentials")
	// Insecure TLS configurations detected. Immediate upgrade recommended.
	RecommendationType_IoT_VulnerableTLSCipherSuite = RecommendationType("IoT_VulnerableTLSCipherSuite")
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

type RuleState string

const (
	RuleStateEnabled  = RuleState("Enabled")
	RuleStateDisabled = RuleState("Disabled")
	RuleStateExpired  = RuleState("Expired")
)

type SecuritySolutionStatus string

const (
	SecuritySolutionStatusEnabled  = SecuritySolutionStatus("Enabled")
	SecuritySolutionStatusDisabled = SecuritySolutionStatus("Disabled")
)

type Severity string

const (
	SeverityLow    = Severity("Low")
	SeverityMedium = Severity("Medium")
	SeverityHigh   = Severity("High")
)

type SeverityEnum string

const (
	SeverityEnumHigh   = SeverityEnum("High")
	SeverityEnumMedium = SeverityEnum("Medium")
	SeverityEnumLow    = SeverityEnum("Low")
)

type Source string

const (
	// Resource is in Azure
	SourceAzure = Source("Azure")
	// Resource in an on premise machine connected to Azure cloud
	SourceOnPremise = Source("OnPremise")
	// SQL Resource in an on premise machine connected to Azure cloud
	SourceOnPremiseSql = Source("OnPremiseSql")
)

type StandardSupportedClouds string

const (
	StandardSupportedCloudsAWS = StandardSupportedClouds("AWS")
	StandardSupportedCloudsGCP = StandardSupportedClouds("GCP")
)

func (StandardSupportedClouds) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardSupportedClouds)(nil)).Elem()
}

func (e StandardSupportedClouds) ToStandardSupportedCloudsOutput() StandardSupportedCloudsOutput {
	return pulumi.ToOutput(e).(StandardSupportedCloudsOutput)
}

func (e StandardSupportedClouds) ToStandardSupportedCloudsOutputWithContext(ctx context.Context) StandardSupportedCloudsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StandardSupportedCloudsOutput)
}

func (e StandardSupportedClouds) ToStandardSupportedCloudsPtrOutput() StandardSupportedCloudsPtrOutput {
	return e.ToStandardSupportedCloudsPtrOutputWithContext(context.Background())
}

func (e StandardSupportedClouds) ToStandardSupportedCloudsPtrOutputWithContext(ctx context.Context) StandardSupportedCloudsPtrOutput {
	return StandardSupportedClouds(e).ToStandardSupportedCloudsOutputWithContext(ctx).ToStandardSupportedCloudsPtrOutputWithContext(ctx)
}

func (e StandardSupportedClouds) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StandardSupportedClouds) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StandardSupportedClouds) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StandardSupportedClouds) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StandardSupportedCloudsOutput struct{ *pulumi.OutputState }

func (StandardSupportedCloudsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardSupportedClouds)(nil)).Elem()
}

func (o StandardSupportedCloudsOutput) ToStandardSupportedCloudsOutput() StandardSupportedCloudsOutput {
	return o
}

func (o StandardSupportedCloudsOutput) ToStandardSupportedCloudsOutputWithContext(ctx context.Context) StandardSupportedCloudsOutput {
	return o
}

func (o StandardSupportedCloudsOutput) ToStandardSupportedCloudsPtrOutput() StandardSupportedCloudsPtrOutput {
	return o.ToStandardSupportedCloudsPtrOutputWithContext(context.Background())
}

func (o StandardSupportedCloudsOutput) ToStandardSupportedCloudsPtrOutputWithContext(ctx context.Context) StandardSupportedCloudsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StandardSupportedClouds) *StandardSupportedClouds {
		return &v
	}).(StandardSupportedCloudsPtrOutput)
}

func (o StandardSupportedCloudsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StandardSupportedCloudsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StandardSupportedClouds) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StandardSupportedCloudsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StandardSupportedCloudsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StandardSupportedClouds) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StandardSupportedCloudsPtrOutput struct{ *pulumi.OutputState }

func (StandardSupportedCloudsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StandardSupportedClouds)(nil)).Elem()
}

func (o StandardSupportedCloudsPtrOutput) ToStandardSupportedCloudsPtrOutput() StandardSupportedCloudsPtrOutput {
	return o
}

func (o StandardSupportedCloudsPtrOutput) ToStandardSupportedCloudsPtrOutputWithContext(ctx context.Context) StandardSupportedCloudsPtrOutput {
	return o
}

func (o StandardSupportedCloudsPtrOutput) Elem() StandardSupportedCloudsOutput {
	return o.ApplyT(func(v *StandardSupportedClouds) StandardSupportedClouds {
		if v != nil {
			return *v
		}
		var ret StandardSupportedClouds
		return ret
	}).(StandardSupportedCloudsOutput)
}

func (o StandardSupportedCloudsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StandardSupportedCloudsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StandardSupportedClouds) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StandardSupportedCloudsInput interface {
	pulumi.Input

	ToStandardSupportedCloudsOutput() StandardSupportedCloudsOutput
	ToStandardSupportedCloudsOutputWithContext(context.Context) StandardSupportedCloudsOutput
}

var standardSupportedCloudsPtrType = reflect.TypeOf((**StandardSupportedClouds)(nil)).Elem()

type StandardSupportedCloudsPtrInput interface {
	pulumi.Input

	ToStandardSupportedCloudsPtrOutput() StandardSupportedCloudsPtrOutput
	ToStandardSupportedCloudsPtrOutputWithContext(context.Context) StandardSupportedCloudsPtrOutput
}

type standardSupportedCloudsPtr string

func StandardSupportedCloudsPtr(v string) StandardSupportedCloudsPtrInput {
	return (*standardSupportedCloudsPtr)(&v)
}

func (*standardSupportedCloudsPtr) ElementType() reflect.Type {
	return standardSupportedCloudsPtrType
}

func (in *standardSupportedCloudsPtr) ToStandardSupportedCloudsPtrOutput() StandardSupportedCloudsPtrOutput {
	return pulumi.ToOutput(in).(StandardSupportedCloudsPtrOutput)
}

func (in *standardSupportedCloudsPtr) ToStandardSupportedCloudsPtrOutputWithContext(ctx context.Context) StandardSupportedCloudsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StandardSupportedCloudsPtrOutput)
}





type StandardSupportedCloudsArrayInput interface {
	pulumi.Input

	ToStandardSupportedCloudsArrayOutput() StandardSupportedCloudsArrayOutput
	ToStandardSupportedCloudsArrayOutputWithContext(context.Context) StandardSupportedCloudsArrayOutput
}

type StandardSupportedCloudsArray []StandardSupportedClouds

func (StandardSupportedCloudsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardSupportedClouds)(nil)).Elem()
}

func (i StandardSupportedCloudsArray) ToStandardSupportedCloudsArrayOutput() StandardSupportedCloudsArrayOutput {
	return i.ToStandardSupportedCloudsArrayOutputWithContext(context.Background())
}

func (i StandardSupportedCloudsArray) ToStandardSupportedCloudsArrayOutputWithContext(ctx context.Context) StandardSupportedCloudsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardSupportedCloudsArrayOutput)
}

type StandardSupportedCloudsArrayOutput struct{ *pulumi.OutputState }

func (StandardSupportedCloudsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardSupportedClouds)(nil)).Elem()
}

func (o StandardSupportedCloudsArrayOutput) ToStandardSupportedCloudsArrayOutput() StandardSupportedCloudsArrayOutput {
	return o
}

func (o StandardSupportedCloudsArrayOutput) ToStandardSupportedCloudsArrayOutputWithContext(ctx context.Context) StandardSupportedCloudsArrayOutput {
	return o
}

func (o StandardSupportedCloudsArrayOutput) Index(i pulumi.IntInput) StandardSupportedCloudsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StandardSupportedClouds {
		return vs[0].([]StandardSupportedClouds)[vs[1].(int)]
	}).(StandardSupportedCloudsOutput)
}

type State string

const (
	// Send notification on new alerts to the subscription's admins
	StateOn = State("On")
	// Don't send notification on new alerts to the subscription's admins
	StateOff = State("Off")
)

type Status string

const (
	StatusRevoked   = Status("Revoked")
	StatusInitiated = Status("Initiated")
)

type StatusReason string

const (
	StatusReasonExpired               = StatusReason("Expired")
	StatusReasonUserRequested         = StatusReason("UserRequested")
	StatusReasonNewerRequestInitiated = StatusReason("NewerRequestInitiated")
)

type SupportedCloudEnum string

const (
	SupportedCloudEnumAWS = SupportedCloudEnum("AWS")
	SupportedCloudEnumGCP = SupportedCloudEnum("GCP")
)

type Threats string

const (
	ThreatsAccountBreach        = Threats("accountBreach")
	ThreatsDataExfiltration     = Threats("dataExfiltration")
	ThreatsDataSpillage         = Threats("dataSpillage")
	ThreatsMaliciousInsider     = Threats("maliciousInsider")
	ThreatsElevationOfPrivilege = Threats("elevationOfPrivilege")
	ThreatsThreatResistance     = Threats("threatResistance")
	ThreatsMissingCoverage      = Threats("missingCoverage")
	ThreatsDenialOfService      = Threats("denialOfService")
)

type UnmaskedIpLoggingStatus string

const (
	// Unmasked IP logging is disabled
	UnmaskedIpLoggingStatusDisabled = UnmaskedIpLoggingStatus("Disabled")
	// Unmasked IP logging is enabled
	UnmaskedIpLoggingStatusEnabled = UnmaskedIpLoggingStatus("Enabled")
)

type UserImpact string

const (
	UserImpactLow      = UserImpact("Low")
	UserImpactModerate = UserImpact("Moderate")
	UserImpactHigh     = UserImpact("High")
)

func init() {
	pulumi.RegisterOutputType(StandardSupportedCloudsOutput{})
	pulumi.RegisterOutputType(StandardSupportedCloudsPtrOutput{})
	pulumi.RegisterOutputType(StandardSupportedCloudsArrayOutput{})
}

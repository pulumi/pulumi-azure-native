


package v20190101preview

type ActionType string

const (
	ActionTypeLogicApp  = ActionType("LogicApp")
	ActionTypeEventHub  = ActionType("EventHub")
	ActionTypeWorkspace = ActionType("Workspace")
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
)

type Categories string

const (
	CategoriesCompute           = Categories("Compute")
	CategoriesNetworking        = Categories("Networking")
	CategoriesData              = Categories("Data")
	CategoriesIdentityAndAccess = Categories("IdentityAndAccess")
	CategoriesIoT               = Categories("IoT")
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

type ImplementationEffort string

const (
	ImplementationEffortLow      = ImplementationEffort("Low")
	ImplementationEffortModerate = ImplementationEffort("Moderate")
	ImplementationEffortHigh     = ImplementationEffort("High")
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

type PropertyType string

const (
	PropertyTypeString  = PropertyType("String")
	PropertyTypeInteger = PropertyType("Integer")
	PropertyTypeNumber  = PropertyType("Number")
	PropertyTypeBoolean = PropertyType("Boolean")
)

type RuleState string

const (
	RuleStateEnabled  = RuleState("Enabled")
	RuleStateDisabled = RuleState("Disabled")
	RuleStateExpired  = RuleState("Expired")
)

type Severity string

const (
	SeverityLow    = Severity("Low")
	SeverityMedium = Severity("Medium")
	SeverityHigh   = Severity("High")
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

type UserImpact string

const (
	UserImpactLow      = UserImpact("Low")
	UserImpactModerate = UserImpact("Moderate")
	UserImpactHigh     = UserImpact("High")
)

func init() {
}

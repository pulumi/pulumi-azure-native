


package v20220101preview

type GovernanceRuleOwnerSourceType string

const (
	// The rule source type defined using resource tag
	GovernanceRuleOwnerSourceTypeByTag = GovernanceRuleOwnerSourceType("ByTag")
	// The rule source type defined manually
	GovernanceRuleOwnerSourceTypeManually = GovernanceRuleOwnerSourceType("Manually")
)

type GovernanceRuleSourceResourceType string

const (
	// The source of the governance rule is assessments
	GovernanceRuleSourceResourceTypeAssessments = GovernanceRuleSourceResourceType("Assessments")
)

type GovernanceRuleType string

const (
	// The source of the rule type definition is integrated
	GovernanceRuleTypeIntegrated = GovernanceRuleType("Integrated")
	// The source of the rule type definition is ServiceNow
	GovernanceRuleTypeServiceNow = GovernanceRuleType("ServiceNow")
)

func init() {
}

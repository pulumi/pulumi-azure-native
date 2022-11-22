


package v20220701preview

type AssignmentScopeValidation string

const (
	// This option will validate the exemption is at or under the assignment scope.
	AssignmentScopeValidationDefault = AssignmentScopeValidation("Default")
	// This option will bypass the validation the exemption scope is at or under the policy assignment scope.
	AssignmentScopeValidationDoNotValidate = AssignmentScopeValidation("DoNotValidate")
)

type ExemptionCategory string

const (
	// This category of exemptions usually means the scope is not applicable for the policy.
	ExemptionCategoryWaiver = ExemptionCategory("Waiver")
	// This category of exemptions usually means the mitigation actions have been applied to the scope.
	ExemptionCategoryMitigated = ExemptionCategory("Mitigated")
)

type SelectorKind string

const (
	// The selector kind to filter policies by the resource location.
	SelectorKindResourceLocation = SelectorKind("resourceLocation")
	// The selector kind to filter policies by the resource type.
	SelectorKindResourceType = SelectorKind("resourceType")
	// The selector kind to filter policies by the resource without location.
	SelectorKindResourceWithoutLocation = SelectorKind("resourceWithoutLocation")
	// The selector kind to filter policies by the policy definition reference ID.
	SelectorKindPolicyDefinitionReferenceId = SelectorKind("policyDefinitionReferenceId")
)

func init() {
}

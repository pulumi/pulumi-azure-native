


package v20200701preview

type ExemptionCategory string

const (
	// This category of exemptions usually means the scope is not applicable for the policy.
	ExemptionCategoryWaiver = ExemptionCategory("Waiver")
	// This category of exemptions usually means the mitigation actions have been applied to the scope.
	ExemptionCategoryMitigated = ExemptionCategory("Mitigated")
)

func init() {
}

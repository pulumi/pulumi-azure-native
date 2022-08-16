


package v20200301preview

type CostAllocationPolicyType string

const (
	CostAllocationPolicyTypeFixedProportion = CostAllocationPolicyType("FixedProportion")
)

type CostAllocationResourceType string

const (
	// Indicates an Azure dimension such as a subscription id or resource group name is being used for allocation.
	CostAllocationResourceTypeDimension = CostAllocationResourceType("Dimension")
	// Allocates cost based on Azure Tag key value pairs.
	CostAllocationResourceTypeTag = CostAllocationResourceType("Tag")
)

type RuleStatus string

const (
	// Rule is saved but not used to allocate costs.
	RuleStatusNotActive = RuleStatus("NotActive")
	// Rule is saved and impacting cost allocation.
	RuleStatusActive = RuleStatus("Active")
	// Rule is saved and cost allocation is being updated. Readonly value that cannot be submitted in a put request.
	RuleStatusProcessing = RuleStatus("Processing")
)

func init() {
}

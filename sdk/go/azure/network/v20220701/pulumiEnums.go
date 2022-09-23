


package v20220701

type ForwardingRuleStateEnum string

const (
	ForwardingRuleStateEnumEnabled  = ForwardingRuleStateEnum("Enabled")
	ForwardingRuleStateEnumDisabled = ForwardingRuleStateEnum("Disabled")
)

type IpAllocationMethod string

const (
	IpAllocationMethodStatic  = IpAllocationMethod("Static")
	IpAllocationMethodDynamic = IpAllocationMethod("Dynamic")
)

func init() {
}

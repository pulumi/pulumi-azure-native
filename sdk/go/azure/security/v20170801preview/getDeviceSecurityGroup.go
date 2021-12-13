


package v20170801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeviceSecurityGroup(ctx *pulumi.Context, args *LookupDeviceSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupDeviceSecurityGroupResult, error) {
	var rv LookupDeviceSecurityGroupResult
	err := ctx.Invoke("azure-native:security/v20170801preview:getDeviceSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceSecurityGroupArgs struct {
	DeviceSecurityGroupName string `pulumi:"deviceSecurityGroupName"`
	ResourceId              string `pulumi:"resourceId"`
}


type LookupDeviceSecurityGroupResult struct {
	AllowlistRules  []AllowlistCustomAlertRuleResponse  `pulumi:"allowlistRules"`
	DenylistRules   []DenylistCustomAlertRuleResponse   `pulumi:"denylistRules"`
	Id              string                              `pulumi:"id"`
	Name            string                              `pulumi:"name"`
	ThresholdRules  []ThresholdCustomAlertRuleResponse  `pulumi:"thresholdRules"`
	TimeWindowRules []TimeWindowCustomAlertRuleResponse `pulumi:"timeWindowRules"`
	Type            string                              `pulumi:"type"`
}

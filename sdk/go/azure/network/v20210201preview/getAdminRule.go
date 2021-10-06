


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupAdminRule(ctx *pulumi.Context, args *LookupAdminRuleArgs, opts ...pulumi.InvokeOption) (*LookupAdminRuleResult, error) {
	var rv LookupAdminRuleResult
	err := ctx.Invoke("azure-native:network/v20210201preview:getAdminRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdminRuleArgs struct {
	ConfigurationName  string `pulumi:"configurationName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	RuleCollectionName string `pulumi:"ruleCollectionName"`
	RuleName           string `pulumi:"ruleName"`
}


type LookupAdminRuleResult struct {
	Etag       string             `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

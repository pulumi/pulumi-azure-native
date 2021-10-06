


package v20181102privatepreview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupActionRuleByName(ctx *pulumi.Context, args *LookupActionRuleByNameArgs, opts ...pulumi.InvokeOption) (*LookupActionRuleByNameResult, error) {
	var rv LookupActionRuleByNameResult
	err := ctx.Invoke("azure-native:alertsmanagement/v20181102privatepreview:getActionRuleByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupActionRuleByNameArgs struct {
	ActionRuleName string `pulumi:"actionRuleName"`
	ResourceGroup  string `pulumi:"resourceGroup"`
}


type LookupActionRuleByNameResult struct {
	Id         string                       `pulumi:"id"`
	Location   string                       `pulumi:"location"`
	Name       string                       `pulumi:"name"`
	Properties ActionRulePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string            `pulumi:"tags"`
	Type       string                       `pulumi:"type"`
}

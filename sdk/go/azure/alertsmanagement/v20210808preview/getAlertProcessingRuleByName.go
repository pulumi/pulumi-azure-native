


package v20210808preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAlertProcessingRuleByName(ctx *pulumi.Context, args *LookupAlertProcessingRuleByNameArgs, opts ...pulumi.InvokeOption) (*LookupAlertProcessingRuleByNameResult, error) {
	var rv LookupAlertProcessingRuleByNameResult
	err := ctx.Invoke("azure-native:alertsmanagement/v20210808preview:getAlertProcessingRuleByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAlertProcessingRuleByNameArgs struct {
	AlertProcessingRuleName string `pulumi:"alertProcessingRuleName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupAlertProcessingRuleByNameResult struct {
	Id         string                                `pulumi:"id"`
	Location   string                                `pulumi:"location"`
	Name       string                                `pulumi:"name"`
	Properties AlertProcessingRulePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                    `pulumi:"systemData"`
	Tags       map[string]string                     `pulumi:"tags"`
	Type       string                                `pulumi:"type"`
}

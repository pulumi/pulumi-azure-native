


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAction(ctx *pulumi.Context, args *LookupActionArgs, opts ...pulumi.InvokeOption) (*LookupActionResult, error) {
	var rv LookupActionResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupActionArgs struct {
	ActionId          string `pulumi:"actionId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupActionResult struct {
	Etag               *string            `pulumi:"etag"`
	Id                 string             `pulumi:"id"`
	LogicAppResourceId string             `pulumi:"logicAppResourceId"`
	Name               string             `pulumi:"name"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
	WorkflowId         *string            `pulumi:"workflowId"`
}




package logic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkflowAccessKey(ctx *pulumi.Context, args *LookupWorkflowAccessKeyArgs, opts ...pulumi.InvokeOption) (*LookupWorkflowAccessKeyResult, error) {
	var rv LookupWorkflowAccessKeyResult
	err := ctx.Invoke("azure-native:logic:getWorkflowAccessKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkflowAccessKeyArgs struct {
	AccessKeyName     string `pulumi:"accessKeyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkflowName      string `pulumi:"workflowName"`
}

type LookupWorkflowAccessKeyResult struct {
	Id        *string `pulumi:"id"`
	Name      string  `pulumi:"name"`
	NotAfter  *string `pulumi:"notAfter"`
	NotBefore *string `pulumi:"notBefore"`
	Type      string  `pulumi:"type"`
}

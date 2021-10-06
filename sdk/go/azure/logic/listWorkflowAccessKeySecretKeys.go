


package logic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkflowAccessKeySecretKeys(ctx *pulumi.Context, args *ListWorkflowAccessKeySecretKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkflowAccessKeySecretKeysResult, error) {
	var rv ListWorkflowAccessKeySecretKeysResult
	err := ctx.Invoke("azure-native:logic:listWorkflowAccessKeySecretKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowAccessKeySecretKeysArgs struct {
	AccessKeyName     string `pulumi:"accessKeyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkflowName      string `pulumi:"workflowName"`
}

type ListWorkflowAccessKeySecretKeysResult struct {
	PrimarySecretKey   string `pulumi:"primarySecretKey"`
	SecondarySecretKey string `pulumi:"secondarySecretKey"`
}

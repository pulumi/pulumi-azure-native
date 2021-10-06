


package logic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkflowVersionTriggerCallbackUrl(ctx *pulumi.Context, args *ListWorkflowVersionTriggerCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListWorkflowVersionTriggerCallbackUrlResult, error) {
	var rv ListWorkflowVersionTriggerCallbackUrlResult
	err := ctx.Invoke("azure-native:logic:listWorkflowVersionTriggerCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowVersionTriggerCallbackUrlArgs struct {
	KeyType           *string `pulumi:"keyType"`
	NotAfter          *string `pulumi:"notAfter"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	TriggerName       string  `pulumi:"triggerName"`
	VersionId         string  `pulumi:"versionId"`
	WorkflowName      string  `pulumi:"workflowName"`
}


type ListWorkflowVersionTriggerCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

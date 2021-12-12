


package v20190501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkflowTriggerCallbackUrl(ctx *pulumi.Context, args *ListWorkflowTriggerCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListWorkflowTriggerCallbackUrlResult, error) {
	var rv ListWorkflowTriggerCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20190501:listWorkflowTriggerCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowTriggerCallbackUrlArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TriggerName       string `pulumi:"triggerName"`
	WorkflowName      string `pulumi:"workflowName"`
}


type ListWorkflowTriggerCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

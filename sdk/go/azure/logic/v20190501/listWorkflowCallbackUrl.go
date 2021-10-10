


package v20190501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkflowCallbackUrl(ctx *pulumi.Context, args *ListWorkflowCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListWorkflowCallbackUrlResult, error) {
	var rv ListWorkflowCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20190501:listWorkflowCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowCallbackUrlArgs struct {
	KeyType           *string `pulumi:"keyType"`
	NotAfter          *string `pulumi:"notAfter"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WorkflowName      string  `pulumi:"workflowName"`
}


type ListWorkflowCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}




package v20160601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkflowVersionCallbackUrl(ctx *pulumi.Context, args *ListWorkflowVersionCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListWorkflowVersionCallbackUrlResult, error) {
	var rv ListWorkflowVersionCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20160601:listWorkflowVersionCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowVersionCallbackUrlArgs struct {
	KeyType           *KeyType `pulumi:"keyType"`
	NotAfter          *string  `pulumi:"notAfter"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	TriggerName       string   `pulumi:"triggerName"`
	VersionId         string   `pulumi:"versionId"`
	WorkflowName      string   `pulumi:"workflowName"`
}


type ListWorkflowVersionCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

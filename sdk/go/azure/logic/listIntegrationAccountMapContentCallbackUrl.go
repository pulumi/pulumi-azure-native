


package logic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountMapContentCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountMapContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountMapContentCallbackUrlResult, error) {
	var rv ListIntegrationAccountMapContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic:listIntegrationAccountMapContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountMapContentCallbackUrlArgs struct {
	IntegrationAccountName string  `pulumi:"integrationAccountName"`
	KeyType                *string `pulumi:"keyType"`
	MapName                string  `pulumi:"mapName"`
	NotAfter               *string `pulumi:"notAfter"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}


type ListIntegrationAccountMapContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

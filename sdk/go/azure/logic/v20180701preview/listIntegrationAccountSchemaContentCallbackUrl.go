


package v20180701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountSchemaContentCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountSchemaContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountSchemaContentCallbackUrlResult, error) {
	var rv ListIntegrationAccountSchemaContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20180701preview:listIntegrationAccountSchemaContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountSchemaContentCallbackUrlArgs struct {
	IntegrationAccountName string  `pulumi:"integrationAccountName"`
	KeyType                *string `pulumi:"keyType"`
	NotAfter               *string `pulumi:"notAfter"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	SchemaName             string  `pulumi:"schemaName"`
}


type ListIntegrationAccountSchemaContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

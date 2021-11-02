


package v20160601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSchemaContentCallbackUrl(ctx *pulumi.Context, args *ListSchemaContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListSchemaContentCallbackUrlResult, error) {
	var rv ListSchemaContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20160601:listSchemaContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSchemaContentCallbackUrlArgs struct {
	IntegrationAccountName string   `pulumi:"integrationAccountName"`
	KeyType                *KeyType `pulumi:"keyType"`
	NotAfter               *string  `pulumi:"notAfter"`
	ResourceGroupName      string   `pulumi:"resourceGroupName"`
	SchemaName             string   `pulumi:"schemaName"`
}


type ListSchemaContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

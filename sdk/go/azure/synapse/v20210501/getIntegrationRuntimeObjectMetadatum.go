


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIntegrationRuntimeObjectMetadatum(ctx *pulumi.Context, args *GetIntegrationRuntimeObjectMetadatumArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeObjectMetadatumResult, error) {
	var rv GetIntegrationRuntimeObjectMetadatumResult
	err := ctx.Invoke("azure-native:synapse/v20210501:getIntegrationRuntimeObjectMetadatum", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeObjectMetadatumArgs struct {
	IntegrationRuntimeName string  `pulumi:"integrationRuntimeName"`
	MetadataPath           *string `pulumi:"metadataPath"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	WorkspaceName          string  `pulumi:"workspaceName"`
}


type GetIntegrationRuntimeObjectMetadatumResult struct {
	NextLink *string       `pulumi:"nextLink"`
	Value    []interface{} `pulumi:"value"`
}

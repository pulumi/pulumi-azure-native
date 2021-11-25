


package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIntegrationRuntimeObjectMetadatum(ctx *pulumi.Context, args *GetIntegrationRuntimeObjectMetadatumArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeObjectMetadatumResult, error) {
	var rv GetIntegrationRuntimeObjectMetadatumResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getIntegrationRuntimeObjectMetadatum", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeObjectMetadatumArgs struct {
	FactoryName            string  `pulumi:"factoryName"`
	IntegrationRuntimeName string  `pulumi:"integrationRuntimeName"`
	MetadataPath           *string `pulumi:"metadataPath"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}


type GetIntegrationRuntimeObjectMetadatumResult struct {
	NextLink *string       `pulumi:"nextLink"`
	Value    []interface{} `pulumi:"value"`
}

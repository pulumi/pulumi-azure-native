


package v20180601

import (
	"context"
	"reflect"

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

func GetIntegrationRuntimeObjectMetadatumOutput(ctx *pulumi.Context, args GetIntegrationRuntimeObjectMetadatumOutputArgs, opts ...pulumi.InvokeOption) GetIntegrationRuntimeObjectMetadatumResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIntegrationRuntimeObjectMetadatumResult, error) {
			args := v.(GetIntegrationRuntimeObjectMetadatumArgs)
			r, err := GetIntegrationRuntimeObjectMetadatum(ctx, &args, opts...)
			var s GetIntegrationRuntimeObjectMetadatumResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIntegrationRuntimeObjectMetadatumResultOutput)
}

type GetIntegrationRuntimeObjectMetadatumOutputArgs struct {
	FactoryName            pulumi.StringInput    `pulumi:"factoryName"`
	IntegrationRuntimeName pulumi.StringInput    `pulumi:"integrationRuntimeName"`
	MetadataPath           pulumi.StringPtrInput `pulumi:"metadataPath"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (GetIntegrationRuntimeObjectMetadatumOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeObjectMetadatumArgs)(nil)).Elem()
}


type GetIntegrationRuntimeObjectMetadatumResultOutput struct{ *pulumi.OutputState }

func (GetIntegrationRuntimeObjectMetadatumResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeObjectMetadatumResult)(nil)).Elem()
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) ToGetIntegrationRuntimeObjectMetadatumResultOutput() GetIntegrationRuntimeObjectMetadatumResultOutput {
	return o
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) ToGetIntegrationRuntimeObjectMetadatumResultOutputWithContext(ctx context.Context) GetIntegrationRuntimeObjectMetadatumResultOutput {
	return o
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeObjectMetadatumResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeObjectMetadatumResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIntegrationRuntimeObjectMetadatumResultOutput{})
}

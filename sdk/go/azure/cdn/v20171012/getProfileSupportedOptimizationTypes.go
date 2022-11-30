


package v20171012

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetProfileSupportedOptimizationTypes(ctx *pulumi.Context, args *GetProfileSupportedOptimizationTypesArgs, opts ...pulumi.InvokeOption) (*GetProfileSupportedOptimizationTypesResult, error) {
	var rv GetProfileSupportedOptimizationTypesResult
	err := ctx.Invoke("azure-native:cdn/v20171012:getProfileSupportedOptimizationTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetProfileSupportedOptimizationTypesArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetProfileSupportedOptimizationTypesResult struct {
	SupportedOptimizationTypes []string `pulumi:"supportedOptimizationTypes"`
}

func GetProfileSupportedOptimizationTypesOutput(ctx *pulumi.Context, args GetProfileSupportedOptimizationTypesOutputArgs, opts ...pulumi.InvokeOption) GetProfileSupportedOptimizationTypesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProfileSupportedOptimizationTypesResult, error) {
			args := v.(GetProfileSupportedOptimizationTypesArgs)
			r, err := GetProfileSupportedOptimizationTypes(ctx, &args, opts...)
			var s GetProfileSupportedOptimizationTypesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProfileSupportedOptimizationTypesResultOutput)
}

type GetProfileSupportedOptimizationTypesOutputArgs struct {
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetProfileSupportedOptimizationTypesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProfileSupportedOptimizationTypesArgs)(nil)).Elem()
}


type GetProfileSupportedOptimizationTypesResultOutput struct{ *pulumi.OutputState }

func (GetProfileSupportedOptimizationTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProfileSupportedOptimizationTypesResult)(nil)).Elem()
}

func (o GetProfileSupportedOptimizationTypesResultOutput) ToGetProfileSupportedOptimizationTypesResultOutput() GetProfileSupportedOptimizationTypesResultOutput {
	return o
}

func (o GetProfileSupportedOptimizationTypesResultOutput) ToGetProfileSupportedOptimizationTypesResultOutputWithContext(ctx context.Context) GetProfileSupportedOptimizationTypesResultOutput {
	return o
}

func (o GetProfileSupportedOptimizationTypesResultOutput) SupportedOptimizationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProfileSupportedOptimizationTypesResult) []string { return v.SupportedOptimizationTypes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProfileSupportedOptimizationTypesResultOutput{})
}

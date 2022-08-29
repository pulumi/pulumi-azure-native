


package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRegistryBuildSourceUploadUrl(ctx *pulumi.Context, args *GetRegistryBuildSourceUploadUrlArgs, opts ...pulumi.InvokeOption) (*GetRegistryBuildSourceUploadUrlResult, error) {
	var rv GetRegistryBuildSourceUploadUrlResult
	err := ctx.Invoke("azure-native:containerregistry:getRegistryBuildSourceUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetRegistryBuildSourceUploadUrlArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetRegistryBuildSourceUploadUrlResult struct {
	RelativePath *string `pulumi:"relativePath"`
	UploadUrl    *string `pulumi:"uploadUrl"`
}

func GetRegistryBuildSourceUploadUrlOutput(ctx *pulumi.Context, args GetRegistryBuildSourceUploadUrlOutputArgs, opts ...pulumi.InvokeOption) GetRegistryBuildSourceUploadUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRegistryBuildSourceUploadUrlResult, error) {
			args := v.(GetRegistryBuildSourceUploadUrlArgs)
			r, err := GetRegistryBuildSourceUploadUrl(ctx, &args, opts...)
			var s GetRegistryBuildSourceUploadUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRegistryBuildSourceUploadUrlResultOutput)
}

type GetRegistryBuildSourceUploadUrlOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetRegistryBuildSourceUploadUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistryBuildSourceUploadUrlArgs)(nil)).Elem()
}


type GetRegistryBuildSourceUploadUrlResultOutput struct{ *pulumi.OutputState }

func (GetRegistryBuildSourceUploadUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistryBuildSourceUploadUrlResult)(nil)).Elem()
}

func (o GetRegistryBuildSourceUploadUrlResultOutput) ToGetRegistryBuildSourceUploadUrlResultOutput() GetRegistryBuildSourceUploadUrlResultOutput {
	return o
}

func (o GetRegistryBuildSourceUploadUrlResultOutput) ToGetRegistryBuildSourceUploadUrlResultOutputWithContext(ctx context.Context) GetRegistryBuildSourceUploadUrlResultOutput {
	return o
}

func (o GetRegistryBuildSourceUploadUrlResultOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryBuildSourceUploadUrlResult) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func (o GetRegistryBuildSourceUploadUrlResultOutput) UploadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryBuildSourceUploadUrlResult) *string { return v.UploadUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRegistryBuildSourceUploadUrlResultOutput{})
}

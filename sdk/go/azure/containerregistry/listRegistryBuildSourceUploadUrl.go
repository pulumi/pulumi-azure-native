


package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRegistryBuildSourceUploadUrl(ctx *pulumi.Context, args *ListRegistryBuildSourceUploadUrlArgs, opts ...pulumi.InvokeOption) (*ListRegistryBuildSourceUploadUrlResult, error) {
	var rv ListRegistryBuildSourceUploadUrlResult
	err := ctx.Invoke("azure-native:containerregistry:listRegistryBuildSourceUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRegistryBuildSourceUploadUrlArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListRegistryBuildSourceUploadUrlResult struct {
	RelativePath *string `pulumi:"relativePath"`
	UploadUrl    *string `pulumi:"uploadUrl"`
}

func ListRegistryBuildSourceUploadUrlOutput(ctx *pulumi.Context, args ListRegistryBuildSourceUploadUrlOutputArgs, opts ...pulumi.InvokeOption) ListRegistryBuildSourceUploadUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRegistryBuildSourceUploadUrlResult, error) {
			args := v.(ListRegistryBuildSourceUploadUrlArgs)
			r, err := ListRegistryBuildSourceUploadUrl(ctx, &args, opts...)
			var s ListRegistryBuildSourceUploadUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRegistryBuildSourceUploadUrlResultOutput)
}

type ListRegistryBuildSourceUploadUrlOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListRegistryBuildSourceUploadUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryBuildSourceUploadUrlArgs)(nil)).Elem()
}


type ListRegistryBuildSourceUploadUrlResultOutput struct{ *pulumi.OutputState }

func (ListRegistryBuildSourceUploadUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryBuildSourceUploadUrlResult)(nil)).Elem()
}

func (o ListRegistryBuildSourceUploadUrlResultOutput) ToListRegistryBuildSourceUploadUrlResultOutput() ListRegistryBuildSourceUploadUrlResultOutput {
	return o
}

func (o ListRegistryBuildSourceUploadUrlResultOutput) ToListRegistryBuildSourceUploadUrlResultOutputWithContext(ctx context.Context) ListRegistryBuildSourceUploadUrlResultOutput {
	return o
}

func (o ListRegistryBuildSourceUploadUrlResultOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListRegistryBuildSourceUploadUrlResult) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func (o ListRegistryBuildSourceUploadUrlResultOutput) UploadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListRegistryBuildSourceUploadUrlResult) *string { return v.UploadUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRegistryBuildSourceUploadUrlResultOutput{})
}

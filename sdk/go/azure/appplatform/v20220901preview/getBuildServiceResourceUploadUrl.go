


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBuildServiceResourceUploadUrl(ctx *pulumi.Context, args *GetBuildServiceResourceUploadUrlArgs, opts ...pulumi.InvokeOption) (*GetBuildServiceResourceUploadUrlResult, error) {
	var rv GetBuildServiceResourceUploadUrlResult
	err := ctx.Invoke("azure-native:appplatform/v20220901preview:getBuildServiceResourceUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBuildServiceResourceUploadUrlArgs struct {
	BuildServiceName  string `pulumi:"buildServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type GetBuildServiceResourceUploadUrlResult struct {
	RelativePath *string `pulumi:"relativePath"`
	UploadUrl    *string `pulumi:"uploadUrl"`
}

func GetBuildServiceResourceUploadUrlOutput(ctx *pulumi.Context, args GetBuildServiceResourceUploadUrlOutputArgs, opts ...pulumi.InvokeOption) GetBuildServiceResourceUploadUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBuildServiceResourceUploadUrlResult, error) {
			args := v.(GetBuildServiceResourceUploadUrlArgs)
			r, err := GetBuildServiceResourceUploadUrl(ctx, &args, opts...)
			var s GetBuildServiceResourceUploadUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBuildServiceResourceUploadUrlResultOutput)
}

type GetBuildServiceResourceUploadUrlOutputArgs struct {
	BuildServiceName  pulumi.StringInput `pulumi:"buildServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (GetBuildServiceResourceUploadUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBuildServiceResourceUploadUrlArgs)(nil)).Elem()
}


type GetBuildServiceResourceUploadUrlResultOutput struct{ *pulumi.OutputState }

func (GetBuildServiceResourceUploadUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBuildServiceResourceUploadUrlResult)(nil)).Elem()
}

func (o GetBuildServiceResourceUploadUrlResultOutput) ToGetBuildServiceResourceUploadUrlResultOutput() GetBuildServiceResourceUploadUrlResultOutput {
	return o
}

func (o GetBuildServiceResourceUploadUrlResultOutput) ToGetBuildServiceResourceUploadUrlResultOutputWithContext(ctx context.Context) GetBuildServiceResourceUploadUrlResultOutput {
	return o
}

func (o GetBuildServiceResourceUploadUrlResultOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBuildServiceResourceUploadUrlResult) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func (o GetBuildServiceResourceUploadUrlResultOutput) UploadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBuildServiceResourceUploadUrlResult) *string { return v.UploadUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBuildServiceResourceUploadUrlResultOutput{})
}

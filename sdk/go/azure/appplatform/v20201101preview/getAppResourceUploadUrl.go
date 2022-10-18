


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAppResourceUploadUrl(ctx *pulumi.Context, args *GetAppResourceUploadUrlArgs, opts ...pulumi.InvokeOption) (*GetAppResourceUploadUrlResult, error) {
	var rv GetAppResourceUploadUrlResult
	err := ctx.Invoke("azure-native:appplatform/v20201101preview:getAppResourceUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAppResourceUploadUrlArgs struct {
	AppName           string `pulumi:"appName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type GetAppResourceUploadUrlResult struct {
	RelativePath *string `pulumi:"relativePath"`
	UploadUrl    *string `pulumi:"uploadUrl"`
}

func GetAppResourceUploadUrlOutput(ctx *pulumi.Context, args GetAppResourceUploadUrlOutputArgs, opts ...pulumi.InvokeOption) GetAppResourceUploadUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppResourceUploadUrlResult, error) {
			args := v.(GetAppResourceUploadUrlArgs)
			r, err := GetAppResourceUploadUrl(ctx, &args, opts...)
			var s GetAppResourceUploadUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppResourceUploadUrlResultOutput)
}

type GetAppResourceUploadUrlOutputArgs struct {
	AppName           pulumi.StringInput `pulumi:"appName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (GetAppResourceUploadUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppResourceUploadUrlArgs)(nil)).Elem()
}


type GetAppResourceUploadUrlResultOutput struct{ *pulumi.OutputState }

func (GetAppResourceUploadUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppResourceUploadUrlResult)(nil)).Elem()
}

func (o GetAppResourceUploadUrlResultOutput) ToGetAppResourceUploadUrlResultOutput() GetAppResourceUploadUrlResultOutput {
	return o
}

func (o GetAppResourceUploadUrlResultOutput) ToGetAppResourceUploadUrlResultOutputWithContext(ctx context.Context) GetAppResourceUploadUrlResultOutput {
	return o
}

func (o GetAppResourceUploadUrlResultOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppResourceUploadUrlResult) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func (o GetAppResourceUploadUrlResultOutput) UploadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppResourceUploadUrlResult) *string { return v.UploadUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppResourceUploadUrlResultOutput{})
}

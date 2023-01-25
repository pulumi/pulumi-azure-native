


package testbase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTestBaseAccountFileUploadUrl(ctx *pulumi.Context, args *GetTestBaseAccountFileUploadUrlArgs, opts ...pulumi.InvokeOption) (*GetTestBaseAccountFileUploadUrlResult, error) {
	var rv GetTestBaseAccountFileUploadUrlResult
	err := ctx.Invoke("azure-native:testbase:getTestBaseAccountFileUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTestBaseAccountFileUploadUrlArgs struct {
	BlobName            *string `pulumi:"blobName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	TestBaseAccountName string  `pulumi:"testBaseAccountName"`
}


type GetTestBaseAccountFileUploadUrlResult struct {
	BlobPath  string `pulumi:"blobPath"`
	UploadUrl string `pulumi:"uploadUrl"`
}

func GetTestBaseAccountFileUploadUrlOutput(ctx *pulumi.Context, args GetTestBaseAccountFileUploadUrlOutputArgs, opts ...pulumi.InvokeOption) GetTestBaseAccountFileUploadUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTestBaseAccountFileUploadUrlResult, error) {
			args := v.(GetTestBaseAccountFileUploadUrlArgs)
			r, err := GetTestBaseAccountFileUploadUrl(ctx, &args, opts...)
			var s GetTestBaseAccountFileUploadUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTestBaseAccountFileUploadUrlResultOutput)
}

type GetTestBaseAccountFileUploadUrlOutputArgs struct {
	BlobName            pulumi.StringPtrInput `pulumi:"blobName"`
	ResourceGroupName   pulumi.StringInput    `pulumi:"resourceGroupName"`
	TestBaseAccountName pulumi.StringInput    `pulumi:"testBaseAccountName"`
}

func (GetTestBaseAccountFileUploadUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestBaseAccountFileUploadUrlArgs)(nil)).Elem()
}


type GetTestBaseAccountFileUploadUrlResultOutput struct{ *pulumi.OutputState }

func (GetTestBaseAccountFileUploadUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestBaseAccountFileUploadUrlResult)(nil)).Elem()
}

func (o GetTestBaseAccountFileUploadUrlResultOutput) ToGetTestBaseAccountFileUploadUrlResultOutput() GetTestBaseAccountFileUploadUrlResultOutput {
	return o
}

func (o GetTestBaseAccountFileUploadUrlResultOutput) ToGetTestBaseAccountFileUploadUrlResultOutputWithContext(ctx context.Context) GetTestBaseAccountFileUploadUrlResultOutput {
	return o
}

func (o GetTestBaseAccountFileUploadUrlResultOutput) BlobPath() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestBaseAccountFileUploadUrlResult) string { return v.BlobPath }).(pulumi.StringOutput)
}

func (o GetTestBaseAccountFileUploadUrlResultOutput) UploadUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestBaseAccountFileUploadUrlResult) string { return v.UploadUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTestBaseAccountFileUploadUrlResultOutput{})
}

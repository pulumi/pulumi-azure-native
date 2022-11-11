


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTestResultDownloadURL(ctx *pulumi.Context, args *GetTestResultDownloadURLArgs, opts ...pulumi.InvokeOption) (*GetTestResultDownloadURLResult, error) {
	var rv GetTestResultDownloadURLResult
	err := ctx.Invoke("azure-native:testbase/v20220401preview:getTestResultDownloadURL", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTestResultDownloadURLArgs struct {
	PackageName         string `pulumi:"packageName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
	TestResultName      string `pulumi:"testResultName"`
}


type GetTestResultDownloadURLResult struct {
	DownloadUrl    string `pulumi:"downloadUrl"`
	ExpirationTime string `pulumi:"expirationTime"`
}

func GetTestResultDownloadURLOutput(ctx *pulumi.Context, args GetTestResultDownloadURLOutputArgs, opts ...pulumi.InvokeOption) GetTestResultDownloadURLResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTestResultDownloadURLResult, error) {
			args := v.(GetTestResultDownloadURLArgs)
			r, err := GetTestResultDownloadURL(ctx, &args, opts...)
			var s GetTestResultDownloadURLResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTestResultDownloadURLResultOutput)
}

type GetTestResultDownloadURLOutputArgs struct {
	PackageName         pulumi.StringInput `pulumi:"packageName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
	TestResultName      pulumi.StringInput `pulumi:"testResultName"`
}

func (GetTestResultDownloadURLOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestResultDownloadURLArgs)(nil)).Elem()
}


type GetTestResultDownloadURLResultOutput struct{ *pulumi.OutputState }

func (GetTestResultDownloadURLResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestResultDownloadURLResult)(nil)).Elem()
}

func (o GetTestResultDownloadURLResultOutput) ToGetTestResultDownloadURLResultOutput() GetTestResultDownloadURLResultOutput {
	return o
}

func (o GetTestResultDownloadURLResultOutput) ToGetTestResultDownloadURLResultOutputWithContext(ctx context.Context) GetTestResultDownloadURLResultOutput {
	return o
}

func (o GetTestResultDownloadURLResultOutput) DownloadUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestResultDownloadURLResult) string { return v.DownloadUrl }).(pulumi.StringOutput)
}

func (o GetTestResultDownloadURLResultOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestResultDownloadURLResult) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTestResultDownloadURLResultOutput{})
}

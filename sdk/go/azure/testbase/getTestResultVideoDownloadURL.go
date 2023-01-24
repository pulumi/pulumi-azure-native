


package testbase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTestResultVideoDownloadURL(ctx *pulumi.Context, args *GetTestResultVideoDownloadURLArgs, opts ...pulumi.InvokeOption) (*GetTestResultVideoDownloadURLResult, error) {
	var rv GetTestResultVideoDownloadURLResult
	err := ctx.Invoke("azure-native:testbase:getTestResultVideoDownloadURL", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTestResultVideoDownloadURLArgs struct {
	PackageName         string `pulumi:"packageName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
	TestResultName      string `pulumi:"testResultName"`
}


type GetTestResultVideoDownloadURLResult struct {
	DownloadUrl    string `pulumi:"downloadUrl"`
	ExpirationTime string `pulumi:"expirationTime"`
}

func GetTestResultVideoDownloadURLOutput(ctx *pulumi.Context, args GetTestResultVideoDownloadURLOutputArgs, opts ...pulumi.InvokeOption) GetTestResultVideoDownloadURLResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTestResultVideoDownloadURLResult, error) {
			args := v.(GetTestResultVideoDownloadURLArgs)
			r, err := GetTestResultVideoDownloadURL(ctx, &args, opts...)
			var s GetTestResultVideoDownloadURLResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTestResultVideoDownloadURLResultOutput)
}

type GetTestResultVideoDownloadURLOutputArgs struct {
	PackageName         pulumi.StringInput `pulumi:"packageName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
	TestResultName      pulumi.StringInput `pulumi:"testResultName"`
}

func (GetTestResultVideoDownloadURLOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestResultVideoDownloadURLArgs)(nil)).Elem()
}


type GetTestResultVideoDownloadURLResultOutput struct{ *pulumi.OutputState }

func (GetTestResultVideoDownloadURLResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestResultVideoDownloadURLResult)(nil)).Elem()
}

func (o GetTestResultVideoDownloadURLResultOutput) ToGetTestResultVideoDownloadURLResultOutput() GetTestResultVideoDownloadURLResultOutput {
	return o
}

func (o GetTestResultVideoDownloadURLResultOutput) ToGetTestResultVideoDownloadURLResultOutputWithContext(ctx context.Context) GetTestResultVideoDownloadURLResultOutput {
	return o
}

func (o GetTestResultVideoDownloadURLResultOutput) DownloadUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestResultVideoDownloadURLResult) string { return v.DownloadUrl }).(pulumi.StringOutput)
}

func (o GetTestResultVideoDownloadURLResultOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestResultVideoDownloadURLResult) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTestResultVideoDownloadURLResultOutput{})
}

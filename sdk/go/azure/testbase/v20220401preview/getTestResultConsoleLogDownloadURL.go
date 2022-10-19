


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTestResultConsoleLogDownloadURL(ctx *pulumi.Context, args *GetTestResultConsoleLogDownloadURLArgs, opts ...pulumi.InvokeOption) (*GetTestResultConsoleLogDownloadURLResult, error) {
	var rv GetTestResultConsoleLogDownloadURLResult
	err := ctx.Invoke("azure-native:testbase/v20220401preview:getTestResultConsoleLogDownloadURL", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTestResultConsoleLogDownloadURLArgs struct {
	LogFileName         string `pulumi:"logFileName"`
	PackageName         string `pulumi:"packageName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
	TestResultName      string `pulumi:"testResultName"`
}


type GetTestResultConsoleLogDownloadURLResult struct {
	DownloadUrl    string `pulumi:"downloadUrl"`
	ExpirationTime string `pulumi:"expirationTime"`
}

func GetTestResultConsoleLogDownloadURLOutput(ctx *pulumi.Context, args GetTestResultConsoleLogDownloadURLOutputArgs, opts ...pulumi.InvokeOption) GetTestResultConsoleLogDownloadURLResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTestResultConsoleLogDownloadURLResult, error) {
			args := v.(GetTestResultConsoleLogDownloadURLArgs)
			r, err := GetTestResultConsoleLogDownloadURL(ctx, &args, opts...)
			var s GetTestResultConsoleLogDownloadURLResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTestResultConsoleLogDownloadURLResultOutput)
}

type GetTestResultConsoleLogDownloadURLOutputArgs struct {
	LogFileName         pulumi.StringInput `pulumi:"logFileName"`
	PackageName         pulumi.StringInput `pulumi:"packageName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
	TestResultName      pulumi.StringInput `pulumi:"testResultName"`
}

func (GetTestResultConsoleLogDownloadURLOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestResultConsoleLogDownloadURLArgs)(nil)).Elem()
}


type GetTestResultConsoleLogDownloadURLResultOutput struct{ *pulumi.OutputState }

func (GetTestResultConsoleLogDownloadURLResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestResultConsoleLogDownloadURLResult)(nil)).Elem()
}

func (o GetTestResultConsoleLogDownloadURLResultOutput) ToGetTestResultConsoleLogDownloadURLResultOutput() GetTestResultConsoleLogDownloadURLResultOutput {
	return o
}

func (o GetTestResultConsoleLogDownloadURLResultOutput) ToGetTestResultConsoleLogDownloadURLResultOutputWithContext(ctx context.Context) GetTestResultConsoleLogDownloadURLResultOutput {
	return o
}

func (o GetTestResultConsoleLogDownloadURLResultOutput) DownloadUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestResultConsoleLogDownloadURLResult) string { return v.DownloadUrl }).(pulumi.StringOutput)
}

func (o GetTestResultConsoleLogDownloadURLResultOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestResultConsoleLogDownloadURLResult) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTestResultConsoleLogDownloadURLResultOutput{})
}

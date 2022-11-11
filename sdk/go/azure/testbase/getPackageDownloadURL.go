


package testbase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPackageDownloadURL(ctx *pulumi.Context, args *GetPackageDownloadURLArgs, opts ...pulumi.InvokeOption) (*GetPackageDownloadURLResult, error) {
	var rv GetPackageDownloadURLResult
	err := ctx.Invoke("azure-native:testbase:getPackageDownloadURL", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetPackageDownloadURLArgs struct {
	PackageName         string `pulumi:"packageName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}


type GetPackageDownloadURLResult struct {
	DownloadUrl    string `pulumi:"downloadUrl"`
	ExpirationTime string `pulumi:"expirationTime"`
}

func GetPackageDownloadURLOutput(ctx *pulumi.Context, args GetPackageDownloadURLOutputArgs, opts ...pulumi.InvokeOption) GetPackageDownloadURLResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPackageDownloadURLResult, error) {
			args := v.(GetPackageDownloadURLArgs)
			r, err := GetPackageDownloadURL(ctx, &args, opts...)
			var s GetPackageDownloadURLResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPackageDownloadURLResultOutput)
}

type GetPackageDownloadURLOutputArgs struct {
	PackageName         pulumi.StringInput `pulumi:"packageName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (GetPackageDownloadURLOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPackageDownloadURLArgs)(nil)).Elem()
}


type GetPackageDownloadURLResultOutput struct{ *pulumi.OutputState }

func (GetPackageDownloadURLResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPackageDownloadURLResult)(nil)).Elem()
}

func (o GetPackageDownloadURLResultOutput) ToGetPackageDownloadURLResultOutput() GetPackageDownloadURLResultOutput {
	return o
}

func (o GetPackageDownloadURLResultOutput) ToGetPackageDownloadURLResultOutputWithContext(ctx context.Context) GetPackageDownloadURLResultOutput {
	return o
}

func (o GetPackageDownloadURLResultOutput) DownloadUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackageDownloadURLResult) string { return v.DownloadUrl }).(pulumi.StringOutput)
}

func (o GetPackageDownloadURLResultOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetPackageDownloadURLResult) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPackageDownloadURLResultOutput{})
}

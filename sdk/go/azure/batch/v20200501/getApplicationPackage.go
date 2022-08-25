


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupApplicationPackage(ctx *pulumi.Context, args *LookupApplicationPackageArgs, opts ...pulumi.InvokeOption) (*LookupApplicationPackageResult, error) {
	var rv LookupApplicationPackageResult
	err := ctx.Invoke("azure-native:batch/v20200501:getApplicationPackage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationPackageArgs struct {
	AccountName       string `pulumi:"accountName"`
	ApplicationName   string `pulumi:"applicationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VersionName       string `pulumi:"versionName"`
}


type LookupApplicationPackageResult struct {
	Etag               string `pulumi:"etag"`
	Format             string `pulumi:"format"`
	Id                 string `pulumi:"id"`
	LastActivationTime string `pulumi:"lastActivationTime"`
	Name               string `pulumi:"name"`
	State              string `pulumi:"state"`
	StorageUrl         string `pulumi:"storageUrl"`
	StorageUrlExpiry   string `pulumi:"storageUrlExpiry"`
	Type               string `pulumi:"type"`
}

func LookupApplicationPackageOutput(ctx *pulumi.Context, args LookupApplicationPackageOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationPackageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationPackageResult, error) {
			args := v.(LookupApplicationPackageArgs)
			r, err := LookupApplicationPackage(ctx, &args, opts...)
			var s LookupApplicationPackageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationPackageResultOutput)
}

type LookupApplicationPackageOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ApplicationName   pulumi.StringInput `pulumi:"applicationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VersionName       pulumi.StringInput `pulumi:"versionName"`
}

func (LookupApplicationPackageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationPackageArgs)(nil)).Elem()
}


type LookupApplicationPackageResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationPackageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationPackageResult)(nil)).Elem()
}

func (o LookupApplicationPackageResultOutput) ToLookupApplicationPackageResultOutput() LookupApplicationPackageResultOutput {
	return o
}

func (o LookupApplicationPackageResultOutput) ToLookupApplicationPackageResultOutputWithContext(ctx context.Context) LookupApplicationPackageResultOutput {
	return o
}

func (o LookupApplicationPackageResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApplicationPackageResultOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.Format }).(pulumi.StringOutput)
}

func (o LookupApplicationPackageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationPackageResultOutput) LastActivationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.LastActivationTime }).(pulumi.StringOutput)
}

func (o LookupApplicationPackageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationPackageResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupApplicationPackageResultOutput) StorageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.StorageUrl }).(pulumi.StringOutput)
}

func (o LookupApplicationPackageResultOutput) StorageUrlExpiry() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.StorageUrlExpiry }).(pulumi.StringOutput)
}

func (o LookupApplicationPackageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationPackageResultOutput{})
}




package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupApplicationPackage(ctx *pulumi.Context, args *LookupApplicationPackageArgs, opts ...pulumi.InvokeOption) (*LookupApplicationPackageResult, error) {
	var rv LookupApplicationPackageResult
	err := ctx.Invoke("azure-native:batch/v20170101:getApplicationPackage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationPackageArgs struct {
	AccountName       string `pulumi:"accountName"`
	ApplicationId     string `pulumi:"applicationId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
}


type LookupApplicationPackageResult struct {
	Format             string `pulumi:"format"`
	Id                 string `pulumi:"id"`
	LastActivationTime string `pulumi:"lastActivationTime"`
	State              string `pulumi:"state"`
	StorageUrl         string `pulumi:"storageUrl"`
	StorageUrlExpiry   string `pulumi:"storageUrlExpiry"`
	Version            string `pulumi:"version"`
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
	ApplicationId     pulumi.StringInput `pulumi:"applicationId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Version           pulumi.StringInput `pulumi:"version"`
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

func (o LookupApplicationPackageResultOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.Format }).(pulumi.StringOutput)
}

func (o LookupApplicationPackageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationPackageResultOutput) LastActivationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.LastActivationTime }).(pulumi.StringOutput)
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

func (o LookupApplicationPackageResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationPackageResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationPackageResultOutput{})
}

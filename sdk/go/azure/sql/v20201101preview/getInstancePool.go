


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInstancePool(ctx *pulumi.Context, args *LookupInstancePoolArgs, opts ...pulumi.InvokeOption) (*LookupInstancePoolResult, error) {
	var rv LookupInstancePoolResult
	err := ctx.Invoke("azure-native:sql/v20201101preview:getInstancePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstancePoolArgs struct {
	InstancePoolName  string `pulumi:"instancePoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupInstancePoolResult struct {
	Id          string            `pulumi:"id"`
	LicenseType string            `pulumi:"licenseType"`
	Location    string            `pulumi:"location"`
	Name        string            `pulumi:"name"`
	Sku         *SkuResponse      `pulumi:"sku"`
	SubnetId    string            `pulumi:"subnetId"`
	Tags        map[string]string `pulumi:"tags"`
	Type        string            `pulumi:"type"`
	VCores      int               `pulumi:"vCores"`
}

func LookupInstancePoolOutput(ctx *pulumi.Context, args LookupInstancePoolOutputArgs, opts ...pulumi.InvokeOption) LookupInstancePoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstancePoolResult, error) {
			args := v.(LookupInstancePoolArgs)
			r, err := LookupInstancePool(ctx, &args, opts...)
			var s LookupInstancePoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstancePoolResultOutput)
}

type LookupInstancePoolOutputArgs struct {
	InstancePoolName  pulumi.StringInput `pulumi:"instancePoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInstancePoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstancePoolArgs)(nil)).Elem()
}


type LookupInstancePoolResultOutput struct{ *pulumi.OutputState }

func (LookupInstancePoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstancePoolResult)(nil)).Elem()
}

func (o LookupInstancePoolResultOutput) ToLookupInstancePoolResultOutput() LookupInstancePoolResultOutput {
	return o
}

func (o LookupInstancePoolResultOutput) ToLookupInstancePoolResultOutputWithContext(ctx context.Context) LookupInstancePoolResultOutput {
	return o
}

func (o LookupInstancePoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInstancePoolResultOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.LicenseType }).(pulumi.StringOutput)
}

func (o LookupInstancePoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupInstancePoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInstancePoolResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupInstancePoolResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupInstancePoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupInstancePoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupInstancePoolResultOutput) VCores() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstancePoolResult) int { return v.VCores }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstancePoolResultOutput{})
}

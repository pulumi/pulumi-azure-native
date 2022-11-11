


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupVirtualNetworkResource(ctx *pulumi.Context, args *LookupVirtualNetworkResourceArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResourceResult, error) {
	var rv LookupVirtualNetworkResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getVirtualNetworkResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkResourceArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupVirtualNetworkResourceResult struct {
	AllowedSubnets             []SubnetResponse         `pulumi:"allowedSubnets"`
	Description                *string                  `pulumi:"description"`
	ExternalProviderResourceId *string                  `pulumi:"externalProviderResourceId"`
	Id                         *string                  `pulumi:"id"`
	Location                   *string                  `pulumi:"location"`
	Name                       *string                  `pulumi:"name"`
	ProvisioningState          *string                  `pulumi:"provisioningState"`
	SubnetOverrides            []SubnetOverrideResponse `pulumi:"subnetOverrides"`
	Tags                       map[string]string        `pulumi:"tags"`
	Type                       *string                  `pulumi:"type"`
}

func LookupVirtualNetworkResourceOutput(ctx *pulumi.Context, args LookupVirtualNetworkResourceOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkResourceResult, error) {
			args := v.(LookupVirtualNetworkResourceArgs)
			r, err := LookupVirtualNetworkResource(ctx, &args, opts...)
			var s LookupVirtualNetworkResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkResourceResultOutput)
}

type LookupVirtualNetworkResourceOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVirtualNetworkResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkResourceArgs)(nil)).Elem()
}


type LookupVirtualNetworkResourceResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkResourceResult)(nil)).Elem()
}

func (o LookupVirtualNetworkResourceResultOutput) ToLookupVirtualNetworkResourceResultOutput() LookupVirtualNetworkResourceResultOutput {
	return o
}

func (o LookupVirtualNetworkResourceResultOutput) ToLookupVirtualNetworkResourceResultOutputWithContext(ctx context.Context) LookupVirtualNetworkResourceResultOutput {
	return o
}

func (o LookupVirtualNetworkResourceResultOutput) AllowedSubnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResourceResult) []SubnetResponse { return v.AllowedSubnets }).(SubnetResponseArrayOutput)
}

func (o LookupVirtualNetworkResourceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResourceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResourceResultOutput) ExternalProviderResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResourceResult) *string { return v.ExternalProviderResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResourceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResourceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResourceResultOutput) SubnetOverrides() SubnetOverrideResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResourceResult) []SubnetOverrideResponse { return v.SubnetOverrides }).(SubnetOverrideResponseArrayOutput)
}

func (o LookupVirtualNetworkResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualNetworkResourceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResourceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkResourceResultOutput{})
}

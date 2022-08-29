


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetwork(ctx *pulumi.Context, args *LookupVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResult, error) {
	var rv LookupVirtualNetworkResult
	err := ctx.Invoke("azure-native:devtestlab:getVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupVirtualNetworkResult struct {
	AllowedSubnets             []SubnetResponse         `pulumi:"allowedSubnets"`
	CreatedDate                string                   `pulumi:"createdDate"`
	Description                *string                  `pulumi:"description"`
	ExternalProviderResourceId *string                  `pulumi:"externalProviderResourceId"`
	ExternalSubnets            []ExternalSubnetResponse `pulumi:"externalSubnets"`
	Id                         string                   `pulumi:"id"`
	Location                   *string                  `pulumi:"location"`
	Name                       string                   `pulumi:"name"`
	ProvisioningState          string                   `pulumi:"provisioningState"`
	SubnetOverrides            []SubnetOverrideResponse `pulumi:"subnetOverrides"`
	Tags                       map[string]string        `pulumi:"tags"`
	Type                       string                   `pulumi:"type"`
	UniqueIdentifier           string                   `pulumi:"uniqueIdentifier"`
}

func LookupVirtualNetworkOutput(ctx *pulumi.Context, args LookupVirtualNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkResult, error) {
			args := v.(LookupVirtualNetworkArgs)
			r, err := LookupVirtualNetwork(ctx, &args, opts...)
			var s LookupVirtualNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkResultOutput)
}

type LookupVirtualNetworkOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupVirtualNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkArgs)(nil)).Elem()
}


type LookupVirtualNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkResult)(nil)).Elem()
}

func (o LookupVirtualNetworkResultOutput) ToLookupVirtualNetworkResultOutput() LookupVirtualNetworkResultOutput {
	return o
}

func (o LookupVirtualNetworkResultOutput) ToLookupVirtualNetworkResultOutputWithContext(ctx context.Context) LookupVirtualNetworkResultOutput {
	return o
}

func (o LookupVirtualNetworkResultOutput) AllowedSubnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) []SubnetResponse { return v.AllowedSubnets }).(SubnetResponseArrayOutput)
}

func (o LookupVirtualNetworkResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResultOutput) ExternalProviderResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.ExternalProviderResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResultOutput) ExternalSubnets() ExternalSubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) []ExternalSubnetResponse { return v.ExternalSubnets }).(ExternalSubnetResponseArrayOutput)
}

func (o LookupVirtualNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) SubnetOverrides() SubnetOverrideResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) []SubnetOverrideResponse { return v.SubnetOverrides }).(SubnetOverrideResponseArrayOutput)
}

func (o LookupVirtualNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkResultOutput{})
}

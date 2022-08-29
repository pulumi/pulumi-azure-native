


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNatGateway(ctx *pulumi.Context, args *LookupNatGatewayArgs, opts ...pulumi.InvokeOption) (*LookupNatGatewayResult, error) {
	var rv LookupNatGatewayResult
	err := ctx.Invoke("azure-native:network/v20210801:getNatGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNatGatewayArgs struct {
	Expand            *string `pulumi:"expand"`
	NatGatewayName    string  `pulumi:"natGatewayName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupNatGatewayResult struct {
	Etag                 string                 `pulumi:"etag"`
	Id                   *string                `pulumi:"id"`
	IdleTimeoutInMinutes *int                   `pulumi:"idleTimeoutInMinutes"`
	Location             *string                `pulumi:"location"`
	Name                 string                 `pulumi:"name"`
	ProvisioningState    string                 `pulumi:"provisioningState"`
	PublicIpAddresses    []SubResourceResponse  `pulumi:"publicIpAddresses"`
	PublicIpPrefixes     []SubResourceResponse  `pulumi:"publicIpPrefixes"`
	ResourceGuid         string                 `pulumi:"resourceGuid"`
	Sku                  *NatGatewaySkuResponse `pulumi:"sku"`
	Subnets              []SubResourceResponse  `pulumi:"subnets"`
	Tags                 map[string]string      `pulumi:"tags"`
	Type                 string                 `pulumi:"type"`
	Zones                []string               `pulumi:"zones"`
}

func LookupNatGatewayOutput(ctx *pulumi.Context, args LookupNatGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupNatGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNatGatewayResult, error) {
			args := v.(LookupNatGatewayArgs)
			r, err := LookupNatGateway(ctx, &args, opts...)
			var s LookupNatGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNatGatewayResultOutput)
}

type LookupNatGatewayOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	NatGatewayName    pulumi.StringInput    `pulumi:"natGatewayName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupNatGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNatGatewayArgs)(nil)).Elem()
}


type LookupNatGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupNatGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNatGatewayResult)(nil)).Elem()
}

func (o LookupNatGatewayResultOutput) ToLookupNatGatewayResultOutput() LookupNatGatewayResultOutput {
	return o
}

func (o LookupNatGatewayResultOutput) ToLookupNatGatewayResultOutputWithContext(ctx context.Context) LookupNatGatewayResultOutput {
	return o
}

func (o LookupNatGatewayResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupNatGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNatGatewayResultOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o LookupNatGatewayResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNatGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNatGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNatGatewayResultOutput) PublicIpAddresses() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) []SubResourceResponse { return v.PublicIpAddresses }).(SubResourceResponseArrayOutput)
}

func (o LookupNatGatewayResultOutput) PublicIpPrefixes() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) []SubResourceResponse { return v.PublicIpPrefixes }).(SubResourceResponseArrayOutput)
}

func (o LookupNatGatewayResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupNatGatewayResultOutput) Sku() NatGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) *NatGatewaySkuResponse { return v.Sku }).(NatGatewaySkuResponsePtrOutput)
}

func (o LookupNatGatewayResultOutput) Subnets() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) []SubResourceResponse { return v.Subnets }).(SubResourceResponseArrayOutput)
}

func (o LookupNatGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNatGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNatGatewayResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNatGatewayResultOutput{})
}

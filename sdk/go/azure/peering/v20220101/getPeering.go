


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPeering(ctx *pulumi.Context, args *LookupPeeringArgs, opts ...pulumi.InvokeOption) (*LookupPeeringResult, error) {
	var rv LookupPeeringResult
	err := ctx.Invoke("azure-native:peering/v20220101:getPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPeeringArgs struct {
	PeeringName       string `pulumi:"peeringName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPeeringResult struct {
	Direct            *PeeringPropertiesDirectResponse   `pulumi:"direct"`
	Exchange          *PeeringPropertiesExchangeResponse `pulumi:"exchange"`
	Id                string                             `pulumi:"id"`
	Kind              string                             `pulumi:"kind"`
	Location          string                             `pulumi:"location"`
	Name              string                             `pulumi:"name"`
	PeeringLocation   *string                            `pulumi:"peeringLocation"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	Sku               PeeringSkuResponse                 `pulumi:"sku"`
	Tags              map[string]string                  `pulumi:"tags"`
	Type              string                             `pulumi:"type"`
}

func LookupPeeringOutput(ctx *pulumi.Context, args LookupPeeringOutputArgs, opts ...pulumi.InvokeOption) LookupPeeringResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPeeringResult, error) {
			args := v.(LookupPeeringArgs)
			r, err := LookupPeering(ctx, &args, opts...)
			var s LookupPeeringResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPeeringResultOutput)
}

type LookupPeeringOutputArgs struct {
	PeeringName       pulumi.StringInput `pulumi:"peeringName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPeeringOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPeeringArgs)(nil)).Elem()
}


type LookupPeeringResultOutput struct{ *pulumi.OutputState }

func (LookupPeeringResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPeeringResult)(nil)).Elem()
}

func (o LookupPeeringResultOutput) ToLookupPeeringResultOutput() LookupPeeringResultOutput {
	return o
}

func (o LookupPeeringResultOutput) ToLookupPeeringResultOutputWithContext(ctx context.Context) LookupPeeringResultOutput {
	return o
}

func (o LookupPeeringResultOutput) Direct() PeeringPropertiesDirectResponsePtrOutput {
	return o.ApplyT(func(v LookupPeeringResult) *PeeringPropertiesDirectResponse { return v.Direct }).(PeeringPropertiesDirectResponsePtrOutput)
}

func (o LookupPeeringResultOutput) Exchange() PeeringPropertiesExchangeResponsePtrOutput {
	return o.ApplyT(func(v LookupPeeringResult) *PeeringPropertiesExchangeResponse { return v.Exchange }).(PeeringPropertiesExchangeResponsePtrOutput)
}

func (o LookupPeeringResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPeeringResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupPeeringResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPeeringResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPeeringResultOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPeeringResult) *string { return v.PeeringLocation }).(pulumi.StringPtrOutput)
}

func (o LookupPeeringResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPeeringResultOutput) Sku() PeeringSkuResponseOutput {
	return o.ApplyT(func(v LookupPeeringResult) PeeringSkuResponse { return v.Sku }).(PeeringSkuResponseOutput)
}

func (o LookupPeeringResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPeeringResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPeeringResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPeeringResultOutput{})
}

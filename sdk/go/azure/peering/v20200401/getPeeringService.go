


package v20200401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupPeeringService(ctx *pulumi.Context, args *LookupPeeringServiceArgs, opts ...pulumi.InvokeOption) (*LookupPeeringServiceResult, error) {
	var rv LookupPeeringServiceResult
	err := ctx.Invoke("azure-native:peering/v20200401:getPeeringService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPeeringServiceArgs struct {
	PeeringServiceName string `pulumi:"peeringServiceName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupPeeringServiceResult struct {
	Id                     string                     `pulumi:"id"`
	Location               string                     `pulumi:"location"`
	Name                   string                     `pulumi:"name"`
	PeeringServiceLocation *string                    `pulumi:"peeringServiceLocation"`
	PeeringServiceProvider *string                    `pulumi:"peeringServiceProvider"`
	ProvisioningState      string                     `pulumi:"provisioningState"`
	Sku                    *PeeringServiceSkuResponse `pulumi:"sku"`
	Tags                   map[string]string          `pulumi:"tags"`
	Type                   string                     `pulumi:"type"`
}

func LookupPeeringServiceOutput(ctx *pulumi.Context, args LookupPeeringServiceOutputArgs, opts ...pulumi.InvokeOption) LookupPeeringServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPeeringServiceResult, error) {
			args := v.(LookupPeeringServiceArgs)
			r, err := LookupPeeringService(ctx, &args, opts...)
			var s LookupPeeringServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPeeringServiceResultOutput)
}

type LookupPeeringServiceOutputArgs struct {
	PeeringServiceName pulumi.StringInput `pulumi:"peeringServiceName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPeeringServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPeeringServiceArgs)(nil)).Elem()
}


type LookupPeeringServiceResultOutput struct{ *pulumi.OutputState }

func (LookupPeeringServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPeeringServiceResult)(nil)).Elem()
}

func (o LookupPeeringServiceResultOutput) ToLookupPeeringServiceResultOutput() LookupPeeringServiceResultOutput {
	return o
}

func (o LookupPeeringServiceResultOutput) ToLookupPeeringServiceResultOutputWithContext(ctx context.Context) LookupPeeringServiceResultOutput {
	return o
}

func (o LookupPeeringServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPeeringServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPeeringServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPeeringServiceResultOutput) PeeringServiceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPeeringServiceResult) *string { return v.PeeringServiceLocation }).(pulumi.StringPtrOutput)
}

func (o LookupPeeringServiceResultOutput) PeeringServiceProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPeeringServiceResult) *string { return v.PeeringServiceProvider }).(pulumi.StringPtrOutput)
}

func (o LookupPeeringServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPeeringServiceResultOutput) Sku() PeeringServiceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupPeeringServiceResult) *PeeringServiceSkuResponse { return v.Sku }).(PeeringServiceSkuResponsePtrOutput)
}

func (o LookupPeeringServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPeeringServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPeeringServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPeeringServiceResultOutput{})
}

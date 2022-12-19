


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegisteredAsn(ctx *pulumi.Context, args *LookupRegisteredAsnArgs, opts ...pulumi.InvokeOption) (*LookupRegisteredAsnResult, error) {
	var rv LookupRegisteredAsnResult
	err := ctx.Invoke("azure-native:peering/v20221001:getRegisteredAsn", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegisteredAsnArgs struct {
	PeeringName       string `pulumi:"peeringName"`
	RegisteredAsnName string `pulumi:"registeredAsnName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRegisteredAsnResult struct {
	Asn                     *int   `pulumi:"asn"`
	Id                      string `pulumi:"id"`
	Name                    string `pulumi:"name"`
	PeeringServicePrefixKey string `pulumi:"peeringServicePrefixKey"`
	ProvisioningState       string `pulumi:"provisioningState"`
	Type                    string `pulumi:"type"`
}

func LookupRegisteredAsnOutput(ctx *pulumi.Context, args LookupRegisteredAsnOutputArgs, opts ...pulumi.InvokeOption) LookupRegisteredAsnResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegisteredAsnResult, error) {
			args := v.(LookupRegisteredAsnArgs)
			r, err := LookupRegisteredAsn(ctx, &args, opts...)
			var s LookupRegisteredAsnResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegisteredAsnResultOutput)
}

type LookupRegisteredAsnOutputArgs struct {
	PeeringName       pulumi.StringInput `pulumi:"peeringName"`
	RegisteredAsnName pulumi.StringInput `pulumi:"registeredAsnName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegisteredAsnOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegisteredAsnArgs)(nil)).Elem()
}


type LookupRegisteredAsnResultOutput struct{ *pulumi.OutputState }

func (LookupRegisteredAsnResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegisteredAsnResult)(nil)).Elem()
}

func (o LookupRegisteredAsnResultOutput) ToLookupRegisteredAsnResultOutput() LookupRegisteredAsnResultOutput {
	return o
}

func (o LookupRegisteredAsnResultOutput) ToLookupRegisteredAsnResultOutputWithContext(ctx context.Context) LookupRegisteredAsnResultOutput {
	return o
}

func (o LookupRegisteredAsnResultOutput) Asn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRegisteredAsnResult) *int { return v.Asn }).(pulumi.IntPtrOutput)
}

func (o LookupRegisteredAsnResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredAsnResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegisteredAsnResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredAsnResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegisteredAsnResultOutput) PeeringServicePrefixKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredAsnResult) string { return v.PeeringServicePrefixKey }).(pulumi.StringOutput)
}

func (o LookupRegisteredAsnResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredAsnResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRegisteredAsnResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredAsnResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegisteredAsnResultOutput{})
}

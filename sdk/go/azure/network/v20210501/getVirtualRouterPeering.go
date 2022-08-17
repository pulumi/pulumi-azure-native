


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualRouterPeering(ctx *pulumi.Context, args *LookupVirtualRouterPeeringArgs, opts ...pulumi.InvokeOption) (*LookupVirtualRouterPeeringResult, error) {
	var rv LookupVirtualRouterPeeringResult
	err := ctx.Invoke("azure-native:network/v20210501:getVirtualRouterPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualRouterPeeringArgs struct {
	PeeringName       string `pulumi:"peeringName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualRouterName string `pulumi:"virtualRouterName"`
}


type LookupVirtualRouterPeeringResult struct {
	Etag              string   `pulumi:"etag"`
	Id                *string  `pulumi:"id"`
	Name              *string  `pulumi:"name"`
	PeerAsn           *float64 `pulumi:"peerAsn"`
	PeerIp            *string  `pulumi:"peerIp"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Type              string   `pulumi:"type"`
}

func LookupVirtualRouterPeeringOutput(ctx *pulumi.Context, args LookupVirtualRouterPeeringOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualRouterPeeringResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualRouterPeeringResult, error) {
			args := v.(LookupVirtualRouterPeeringArgs)
			r, err := LookupVirtualRouterPeering(ctx, &args, opts...)
			var s LookupVirtualRouterPeeringResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualRouterPeeringResultOutput)
}

type LookupVirtualRouterPeeringOutputArgs struct {
	PeeringName       pulumi.StringInput `pulumi:"peeringName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualRouterName pulumi.StringInput `pulumi:"virtualRouterName"`
}

func (LookupVirtualRouterPeeringOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualRouterPeeringArgs)(nil)).Elem()
}


type LookupVirtualRouterPeeringResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualRouterPeeringResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualRouterPeeringResult)(nil)).Elem()
}

func (o LookupVirtualRouterPeeringResultOutput) ToLookupVirtualRouterPeeringResultOutput() LookupVirtualRouterPeeringResultOutput {
	return o
}

func (o LookupVirtualRouterPeeringResultOutput) ToLookupVirtualRouterPeeringResultOutputWithContext(ctx context.Context) LookupVirtualRouterPeeringResultOutput {
	return o
}

func (o LookupVirtualRouterPeeringResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualRouterPeeringResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualRouterPeeringResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualRouterPeeringResultOutput) PeerAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) *float64 { return v.PeerAsn }).(pulumi.Float64PtrOutput)
}

func (o LookupVirtualRouterPeeringResultOutput) PeerIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) *string { return v.PeerIp }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualRouterPeeringResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualRouterPeeringResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterPeeringResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualRouterPeeringResultOutput{})
}

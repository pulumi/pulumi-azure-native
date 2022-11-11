


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkLink(ctx *pulumi.Context, args *LookupVirtualNetworkLinkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkLinkResult, error) {
	var rv LookupVirtualNetworkLinkResult
	err := ctx.Invoke("azure-native:network/v20200401preview:getVirtualNetworkLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkLinkArgs struct {
	DnsForwardingRulesetName string `pulumi:"dnsForwardingRulesetName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
	VirtualNetworkLinkName   string `pulumi:"virtualNetworkLinkName"`
}


type LookupVirtualNetworkLinkResult struct {
	Etag              string               `pulumi:"etag"`
	Id                string               `pulumi:"id"`
	Metadata          map[string]string    `pulumi:"metadata"`
	Name              string               `pulumi:"name"`
	ProvisioningState string               `pulumi:"provisioningState"`
	SystemData        SystemDataResponse   `pulumi:"systemData"`
	Type              string               `pulumi:"type"`
	VirtualNetwork    *SubResourceResponse `pulumi:"virtualNetwork"`
}

func LookupVirtualNetworkLinkOutput(ctx *pulumi.Context, args LookupVirtualNetworkLinkOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkLinkResult, error) {
			args := v.(LookupVirtualNetworkLinkArgs)
			r, err := LookupVirtualNetworkLink(ctx, &args, opts...)
			var s LookupVirtualNetworkLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkLinkResultOutput)
}

type LookupVirtualNetworkLinkOutputArgs struct {
	DnsForwardingRulesetName pulumi.StringInput `pulumi:"dnsForwardingRulesetName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkLinkName   pulumi.StringInput `pulumi:"virtualNetworkLinkName"`
}

func (LookupVirtualNetworkLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkLinkArgs)(nil)).Elem()
}


type LookupVirtualNetworkLinkResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkLinkResult)(nil)).Elem()
}

func (o LookupVirtualNetworkLinkResultOutput) ToLookupVirtualNetworkLinkResultOutput() LookupVirtualNetworkLinkResultOutput {
	return o
}

func (o LookupVirtualNetworkLinkResultOutput) ToLookupVirtualNetworkLinkResultOutputWithContext(ctx context.Context) LookupVirtualNetworkLinkResultOutput {
	return o
}

func (o LookupVirtualNetworkLinkResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) VirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) *SubResourceResponse { return v.VirtualNetwork }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkLinkResultOutput{})
}

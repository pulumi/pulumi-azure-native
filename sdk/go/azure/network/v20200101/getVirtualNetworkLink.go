


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkLink(ctx *pulumi.Context, args *LookupVirtualNetworkLinkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkLinkResult, error) {
	var rv LookupVirtualNetworkLinkResult
	err := ctx.Invoke("azure-native:network/v20200101:getVirtualNetworkLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkLinkArgs struct {
	PrivateZoneName        string `pulumi:"privateZoneName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	VirtualNetworkLinkName string `pulumi:"virtualNetworkLinkName"`
}


type LookupVirtualNetworkLinkResult struct {
	Etag                    *string              `pulumi:"etag"`
	Id                      string               `pulumi:"id"`
	Location                *string              `pulumi:"location"`
	Name                    string               `pulumi:"name"`
	ProvisioningState       string               `pulumi:"provisioningState"`
	RegistrationEnabled     *bool                `pulumi:"registrationEnabled"`
	Tags                    map[string]string    `pulumi:"tags"`
	Type                    string               `pulumi:"type"`
	VirtualNetwork          *SubResourceResponse `pulumi:"virtualNetwork"`
	VirtualNetworkLinkState string               `pulumi:"virtualNetworkLinkState"`
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
	PrivateZoneName        pulumi.StringInput `pulumi:"privateZoneName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkLinkName pulumi.StringInput `pulumi:"virtualNetworkLinkName"`
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

func (o LookupVirtualNetworkLinkResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) RegistrationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) *bool { return v.RegistrationEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) VirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) *SubResourceResponse { return v.VirtualNetwork }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualNetworkLinkResultOutput) VirtualNetworkLinkState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkLinkResult) string { return v.VirtualNetworkLinkState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkLinkResultOutput{})
}

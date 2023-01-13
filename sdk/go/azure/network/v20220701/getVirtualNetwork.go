


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetwork(ctx *pulumi.Context, args *LookupVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResult, error) {
	var rv LookupVirtualNetworkResult
	err := ctx.Invoke("azure-native:network/v20220701:getVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualNetworkArgs struct {
	Expand             *string `pulumi:"expand"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	VirtualNetworkName string  `pulumi:"virtualNetworkName"`
}


type LookupVirtualNetworkResult struct {
	AddressSpace           *AddressSpaceResponse                 `pulumi:"addressSpace"`
	BgpCommunities         *VirtualNetworkBgpCommunitiesResponse `pulumi:"bgpCommunities"`
	DdosProtectionPlan     *SubResourceResponse                  `pulumi:"ddosProtectionPlan"`
	DhcpOptions            *DhcpOptionsResponse                  `pulumi:"dhcpOptions"`
	EnableDdosProtection   *bool                                 `pulumi:"enableDdosProtection"`
	EnableVmProtection     *bool                                 `pulumi:"enableVmProtection"`
	Encryption             *VirtualNetworkEncryptionResponse     `pulumi:"encryption"`
	Etag                   string                                `pulumi:"etag"`
	ExtendedLocation       *ExtendedLocationResponse             `pulumi:"extendedLocation"`
	FlowTimeoutInMinutes   *int                                  `pulumi:"flowTimeoutInMinutes"`
	Id                     *string                               `pulumi:"id"`
	IpAllocations          []SubResourceResponse                 `pulumi:"ipAllocations"`
	Location               *string                               `pulumi:"location"`
	Name                   string                                `pulumi:"name"`
	ProvisioningState      string                                `pulumi:"provisioningState"`
	ResourceGuid           string                                `pulumi:"resourceGuid"`
	Subnets                []SubnetResponse                      `pulumi:"subnets"`
	Tags                   map[string]string                     `pulumi:"tags"`
	Type                   string                                `pulumi:"type"`
	VirtualNetworkPeerings []VirtualNetworkPeeringResponse       `pulumi:"virtualNetworkPeerings"`
}


func (val *LookupVirtualNetworkResult) Defaults() *LookupVirtualNetworkResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableDdosProtection) {
		enableDdosProtection_ := false
		tmp.EnableDdosProtection = &enableDdosProtection_
	}
	if isZero(tmp.EnableVmProtection) {
		enableVmProtection_ := false
		tmp.EnableVmProtection = &enableVmProtection_
	}
	return &tmp
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
	Expand             pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	VirtualNetworkName pulumi.StringInput    `pulumi:"virtualNetworkName"`
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

func (o LookupVirtualNetworkResultOutput) AddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *AddressSpaceResponse { return v.AddressSpace }).(AddressSpaceResponsePtrOutput)
}

func (o LookupVirtualNetworkResultOutput) BgpCommunities() VirtualNetworkBgpCommunitiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *VirtualNetworkBgpCommunitiesResponse { return v.BgpCommunities }).(VirtualNetworkBgpCommunitiesResponsePtrOutput)
}

func (o LookupVirtualNetworkResultOutput) DdosProtectionPlan() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *SubResourceResponse { return v.DdosProtectionPlan }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualNetworkResultOutput) DhcpOptions() DhcpOptionsResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *DhcpOptionsResponse { return v.DhcpOptions }).(DhcpOptionsResponsePtrOutput)
}

func (o LookupVirtualNetworkResultOutput) EnableDdosProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *bool { return v.EnableDdosProtection }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkResultOutput) EnableVmProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *bool { return v.EnableVmProtection }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkResultOutput) Encryption() VirtualNetworkEncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *VirtualNetworkEncryptionResponse { return v.Encryption }).(VirtualNetworkEncryptionResponsePtrOutput)
}

func (o LookupVirtualNetworkResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupVirtualNetworkResultOutput) FlowTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *int { return v.FlowTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualNetworkResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkResultOutput) IpAllocations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) []SubResourceResponse { return v.IpAllocations }).(SubResourceResponseArrayOutput)
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

func (o LookupVirtualNetworkResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

func (o LookupVirtualNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkResultOutput) VirtualNetworkPeerings() VirtualNetworkPeeringResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) []VirtualNetworkPeeringResponse { return v.VirtualNetworkPeerings }).(VirtualNetworkPeeringResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkResultOutput{})
}

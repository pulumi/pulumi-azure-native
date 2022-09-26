


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIpAllocation(ctx *pulumi.Context, args *LookupIpAllocationArgs, opts ...pulumi.InvokeOption) (*LookupIpAllocationResult, error) {
	var rv LookupIpAllocationResult
	err := ctx.Invoke("azure-native:network/v20220501:getIpAllocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupIpAllocationArgs struct {
	Expand            *string `pulumi:"expand"`
	IpAllocationName  string  `pulumi:"ipAllocationName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupIpAllocationResult struct {
	AllocationTags   map[string]string   `pulumi:"allocationTags"`
	Etag             string              `pulumi:"etag"`
	Id               *string             `pulumi:"id"`
	IpamAllocationId *string             `pulumi:"ipamAllocationId"`
	Location         *string             `pulumi:"location"`
	Name             string              `pulumi:"name"`
	Prefix           *string             `pulumi:"prefix"`
	PrefixLength     *int                `pulumi:"prefixLength"`
	PrefixType       *string             `pulumi:"prefixType"`
	Subnet           SubResourceResponse `pulumi:"subnet"`
	Tags             map[string]string   `pulumi:"tags"`
	Type             string              `pulumi:"type"`
	VirtualNetwork   SubResourceResponse `pulumi:"virtualNetwork"`
}


func (val *LookupIpAllocationResult) Defaults() *LookupIpAllocationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PrefixLength) {
		prefixLength_ := 0
		tmp.PrefixLength = &prefixLength_
	}
	return &tmp
}

func LookupIpAllocationOutput(ctx *pulumi.Context, args LookupIpAllocationOutputArgs, opts ...pulumi.InvokeOption) LookupIpAllocationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpAllocationResult, error) {
			args := v.(LookupIpAllocationArgs)
			r, err := LookupIpAllocation(ctx, &args, opts...)
			var s LookupIpAllocationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpAllocationResultOutput)
}

type LookupIpAllocationOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	IpAllocationName  pulumi.StringInput    `pulumi:"ipAllocationName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupIpAllocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpAllocationArgs)(nil)).Elem()
}


type LookupIpAllocationResultOutput struct{ *pulumi.OutputState }

func (LookupIpAllocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpAllocationResult)(nil)).Elem()
}

func (o LookupIpAllocationResultOutput) ToLookupIpAllocationResultOutput() LookupIpAllocationResultOutput {
	return o
}

func (o LookupIpAllocationResultOutput) ToLookupIpAllocationResultOutputWithContext(ctx context.Context) LookupIpAllocationResultOutput {
	return o
}

func (o LookupIpAllocationResultOutput) AllocationTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) map[string]string { return v.AllocationTags }).(pulumi.StringMapOutput)
}

func (o LookupIpAllocationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupIpAllocationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupIpAllocationResultOutput) IpamAllocationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) *string { return v.IpamAllocationId }).(pulumi.StringPtrOutput)
}

func (o LookupIpAllocationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIpAllocationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIpAllocationResultOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

func (o LookupIpAllocationResultOutput) PrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) *int { return v.PrefixLength }).(pulumi.IntPtrOutput)
}

func (o LookupIpAllocationResultOutput) PrefixType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) *string { return v.PrefixType }).(pulumi.StringPtrOutput)
}

func (o LookupIpAllocationResultOutput) Subnet() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) SubResourceResponse { return v.Subnet }).(SubResourceResponseOutput)
}

func (o LookupIpAllocationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIpAllocationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupIpAllocationResultOutput) VirtualNetwork() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) SubResourceResponse { return v.VirtualNetwork }).(SubResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpAllocationResultOutput{})
}

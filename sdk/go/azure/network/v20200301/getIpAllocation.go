


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIpAllocation(ctx *pulumi.Context, args *LookupIpAllocationArgs, opts ...pulumi.InvokeOption) (*LookupIpAllocationResult, error) {
	var rv LookupIpAllocationResult
	err := ctx.Invoke("azure-native:network/v20200301:getIpAllocation", args, &rv, opts...)
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

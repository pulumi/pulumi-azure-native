


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIpGroup(ctx *pulumi.Context, args *LookupIpGroupArgs, opts ...pulumi.InvokeOption) (*LookupIpGroupResult, error) {
	var rv LookupIpGroupResult
	err := ctx.Invoke("azure-native:network/v20210501:getIpGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpGroupArgs struct {
	Expand            *string `pulumi:"expand"`
	IpGroupsName      string  `pulumi:"ipGroupsName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupIpGroupResult struct {
	Etag              string                `pulumi:"etag"`
	FirewallPolicies  []SubResourceResponse `pulumi:"firewallPolicies"`
	Firewalls         []SubResourceResponse `pulumi:"firewalls"`
	Id                *string               `pulumi:"id"`
	IpAddresses       []string              `pulumi:"ipAddresses"`
	Location          *string               `pulumi:"location"`
	Name              string                `pulumi:"name"`
	ProvisioningState string                `pulumi:"provisioningState"`
	Tags              map[string]string     `pulumi:"tags"`
	Type              string                `pulumi:"type"`
}




package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDedicatedHostGroup(ctx *pulumi.Context, args *LookupDedicatedHostGroupArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHostGroupResult, error) {
	var rv LookupDedicatedHostGroupResult
	err := ctx.Invoke("azure-native:compute/v20190701:getDedicatedHostGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedHostGroupArgs struct {
	HostGroupName     string `pulumi:"hostGroupName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDedicatedHostGroupResult struct {
	Hosts                    []SubResourceReadOnlyResponse `pulumi:"hosts"`
	Id                       string                        `pulumi:"id"`
	Location                 string                        `pulumi:"location"`
	Name                     string                        `pulumi:"name"`
	PlatformFaultDomainCount int                           `pulumi:"platformFaultDomainCount"`
	Tags                     map[string]string             `pulumi:"tags"`
	Type                     string                        `pulumi:"type"`
	Zones                    []string                      `pulumi:"zones"`
}

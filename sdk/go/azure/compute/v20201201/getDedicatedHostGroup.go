


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDedicatedHostGroup(ctx *pulumi.Context, args *LookupDedicatedHostGroupArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHostGroupResult, error) {
	var rv LookupDedicatedHostGroupResult
	err := ctx.Invoke("azure-native:compute/v20201201:getDedicatedHostGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedHostGroupArgs struct {
	Expand            *string `pulumi:"expand"`
	HostGroupName     string  `pulumi:"hostGroupName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupDedicatedHostGroupResult struct {
	Hosts                     []SubResourceReadOnlyResponse          `pulumi:"hosts"`
	Id                        string                                 `pulumi:"id"`
	InstanceView              DedicatedHostGroupInstanceViewResponse `pulumi:"instanceView"`
	Location                  string                                 `pulumi:"location"`
	Name                      string                                 `pulumi:"name"`
	PlatformFaultDomainCount  int                                    `pulumi:"platformFaultDomainCount"`
	SupportAutomaticPlacement *bool                                  `pulumi:"supportAutomaticPlacement"`
	Tags                      map[string]string                      `pulumi:"tags"`
	Type                      string                                 `pulumi:"type"`
	Zones                     []string                               `pulumi:"zones"`
}

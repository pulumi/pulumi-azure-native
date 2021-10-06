


package web

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServiceEnvironment(ctx *pulumi.Context, args *LookupAppServiceEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceEnvironmentResult, error) {
	var rv LookupAppServiceEnvironmentResult
	err := ctx.Invoke("azure-native:web:getAppServiceEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServiceEnvironmentArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAppServiceEnvironmentResult struct {
	ClusterSettings           []NameValuePairResponse       `pulumi:"clusterSettings"`
	DedicatedHostCount        int                           `pulumi:"dedicatedHostCount"`
	DnsSuffix                 *string                       `pulumi:"dnsSuffix"`
	FrontEndScaleFactor       *int                          `pulumi:"frontEndScaleFactor"`
	HasLinuxWorkers           bool                          `pulumi:"hasLinuxWorkers"`
	Id                        string                        `pulumi:"id"`
	InternalLoadBalancingMode *string                       `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount         *int                          `pulumi:"ipsslAddressCount"`
	Kind                      *string                       `pulumi:"kind"`
	Location                  string                        `pulumi:"location"`
	MaximumNumberOfMachines   int                           `pulumi:"maximumNumberOfMachines"`
	MultiRoleCount            int                           `pulumi:"multiRoleCount"`
	MultiSize                 *string                       `pulumi:"multiSize"`
	Name                      string                        `pulumi:"name"`
	ProvisioningState         string                        `pulumi:"provisioningState"`
	Status                    string                        `pulumi:"status"`
	Suspended                 bool                          `pulumi:"suspended"`
	Tags                      map[string]string             `pulumi:"tags"`
	Type                      string                        `pulumi:"type"`
	UserWhitelistedIpRanges   []string                      `pulumi:"userWhitelistedIpRanges"`
	VirtualNetwork            VirtualNetworkProfileResponse `pulumi:"virtualNetwork"`
}

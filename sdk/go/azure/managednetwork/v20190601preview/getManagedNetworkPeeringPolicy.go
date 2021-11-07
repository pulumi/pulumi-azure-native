


package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedNetworkPeeringPolicy(ctx *pulumi.Context, args *LookupManagedNetworkPeeringPolicyArgs, opts ...pulumi.InvokeOption) (*LookupManagedNetworkPeeringPolicyResult, error) {
	var rv LookupManagedNetworkPeeringPolicyResult
	err := ctx.Invoke("azure-native:managednetwork/v20190601preview:getManagedNetworkPeeringPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedNetworkPeeringPolicyArgs struct {
	ManagedNetworkName              string `pulumi:"managedNetworkName"`
	ManagedNetworkPeeringPolicyName string `pulumi:"managedNetworkPeeringPolicyName"`
	ResourceGroupName               string `pulumi:"resourceGroupName"`
}


type LookupManagedNetworkPeeringPolicyResult struct {
	Id         string                                        `pulumi:"id"`
	Location   *string                                       `pulumi:"location"`
	Name       string                                        `pulumi:"name"`
	Properties ManagedNetworkPeeringPolicyPropertiesResponse `pulumi:"properties"`
	Type       string                                        `pulumi:"type"`
}

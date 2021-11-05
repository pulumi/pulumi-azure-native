


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoutingIntent(ctx *pulumi.Context, args *LookupRoutingIntentArgs, opts ...pulumi.InvokeOption) (*LookupRoutingIntentResult, error) {
	var rv LookupRoutingIntentResult
	err := ctx.Invoke("azure-native:network:getRoutingIntent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoutingIntentArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RoutingIntentName string `pulumi:"routingIntentName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupRoutingIntentResult struct {
	Etag              string                  `pulumi:"etag"`
	Id                *string                 `pulumi:"id"`
	Name              *string                 `pulumi:"name"`
	ProvisioningState string                  `pulumi:"provisioningState"`
	RoutingPolicies   []RoutingPolicyResponse `pulumi:"routingPolicies"`
	Type              string                  `pulumi:"type"`
}

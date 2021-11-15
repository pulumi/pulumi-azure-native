


package v20210115

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppPrivateEndpointConnectionSlot(ctx *pulumi.Context, args *LookupWebAppPrivateEndpointConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPrivateEndpointConnectionSlotResult, error) {
	var rv LookupWebAppPrivateEndpointConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20210115:getWebAppPrivateEndpointConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPrivateEndpointConnectionSlotArgs struct {
	Name                          string `pulumi:"name"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	Slot                          string `pulumi:"slot"`
}


type LookupWebAppPrivateEndpointConnectionSlotResult struct {
	Id                                string                              `pulumi:"id"`
	IpAddresses                       []string                            `pulumi:"ipAddresses"`
	Kind                              *string                             `pulumi:"kind"`
	Name                              string                              `pulumi:"name"`
	PrivateEndpoint                   *ArmIdWrapperResponse               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
	Type                              string                              `pulumi:"type"`
}

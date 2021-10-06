


package cdn

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrigin(ctx *pulumi.Context, args *LookupOriginArgs, opts ...pulumi.InvokeOption) (*LookupOriginResult, error) {
	var rv LookupOriginResult
	err := ctx.Invoke("azure-native:cdn:getOrigin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOriginArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	OriginName        string `pulumi:"originName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOriginResult struct {
	Enabled                    *bool              `pulumi:"enabled"`
	HostName                   string             `pulumi:"hostName"`
	HttpPort                   *int               `pulumi:"httpPort"`
	HttpsPort                  *int               `pulumi:"httpsPort"`
	Id                         string             `pulumi:"id"`
	Name                       string             `pulumi:"name"`
	OriginHostHeader           *string            `pulumi:"originHostHeader"`
	Priority                   *int               `pulumi:"priority"`
	PrivateEndpointStatus      string             `pulumi:"privateEndpointStatus"`
	PrivateLinkAlias           *string            `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage *string            `pulumi:"privateLinkApprovalMessage"`
	PrivateLinkLocation        *string            `pulumi:"privateLinkLocation"`
	PrivateLinkResourceId      *string            `pulumi:"privateLinkResourceId"`
	ProvisioningState          string             `pulumi:"provisioningState"`
	ResourceState              string             `pulumi:"resourceState"`
	SystemData                 SystemDataResponse `pulumi:"systemData"`
	Type                       string             `pulumi:"type"`
	Weight                     *int               `pulumi:"weight"`
}

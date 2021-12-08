


package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupResourceManagementPrivateLink(ctx *pulumi.Context, args *LookupResourceManagementPrivateLinkArgs, opts ...pulumi.InvokeOption) (*LookupResourceManagementPrivateLinkResult, error) {
	var rv LookupResourceManagementPrivateLinkResult
	err := ctx.Invoke("azure-native:authorization/v20200501:getResourceManagementPrivateLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceManagementPrivateLinkArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RmplName          string `pulumi:"rmplName"`
}

type LookupResourceManagementPrivateLinkResult struct {
	Id         string                                                   `pulumi:"id"`
	Location   *string                                                  `pulumi:"location"`
	Name       string                                                   `pulumi:"name"`
	Properties ResourceManagementPrivateLinkEndpointConnectionsResponse `pulumi:"properties"`
	Type       string                                                   `pulumi:"type"`
}

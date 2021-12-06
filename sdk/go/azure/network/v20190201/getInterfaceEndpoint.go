


package v20190201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInterfaceEndpoint(ctx *pulumi.Context, args *LookupInterfaceEndpointArgs, opts ...pulumi.InvokeOption) (*LookupInterfaceEndpointResult, error) {
	var rv LookupInterfaceEndpointResult
	err := ctx.Invoke("azure-native:network/v20190201:getInterfaceEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInterfaceEndpointArgs struct {
	Expand                *string `pulumi:"expand"`
	InterfaceEndpointName string  `pulumi:"interfaceEndpointName"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
}


type LookupInterfaceEndpointResult struct {
	EndpointService   *EndpointServiceResponse   `pulumi:"endpointService"`
	Etag              *string                    `pulumi:"etag"`
	Fqdn              *string                    `pulumi:"fqdn"`
	Id                *string                    `pulumi:"id"`
	Location          *string                    `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	NetworkInterfaces []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	Owner             string                     `pulumi:"owner"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	Subnet            *SubnetResponse            `pulumi:"subnet"`
	Tags              map[string]string          `pulumi:"tags"`
	Type              string                     `pulumi:"type"`
}

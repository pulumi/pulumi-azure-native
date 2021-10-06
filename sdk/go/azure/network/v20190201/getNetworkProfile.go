


package v20190201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkProfile(ctx *pulumi.Context, args *LookupNetworkProfileArgs, opts ...pulumi.InvokeOption) (*LookupNetworkProfileResult, error) {
	var rv LookupNetworkProfileResult
	err := ctx.Invoke("azure-native:network/v20190201:getNetworkProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkProfileArgs struct {
	Expand             *string `pulumi:"expand"`
	NetworkProfileName string  `pulumi:"networkProfileName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type LookupNetworkProfileResult struct {
	ContainerNetworkInterfaceConfigurations []ContainerNetworkInterfaceConfigurationResponse `pulumi:"containerNetworkInterfaceConfigurations"`
	ContainerNetworkInterfaces              []ContainerNetworkInterfaceResponse              `pulumi:"containerNetworkInterfaces"`
	Etag                                    *string                                          `pulumi:"etag"`
	Id                                      *string                                          `pulumi:"id"`
	Location                                *string                                          `pulumi:"location"`
	Name                                    string                                           `pulumi:"name"`
	ProvisioningState                       string                                           `pulumi:"provisioningState"`
	ResourceGuid                            string                                           `pulumi:"resourceGuid"`
	Tags                                    map[string]string                                `pulumi:"tags"`
	Type                                    string                                           `pulumi:"type"`
}




package v20211015

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotDpsResourcePrivateEndpointConnection(ctx *pulumi.Context, args *LookupIotDpsResourcePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupIotDpsResourcePrivateEndpointConnectionResult, error) {
	var rv LookupIotDpsResourcePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:devices/v20211015:getIotDpsResourcePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotDpsResourcePrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupIotDpsResourcePrivateEndpointConnectionResult struct {
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                          `pulumi:"systemData"`
	Type       string                                      `pulumi:"type"`
}

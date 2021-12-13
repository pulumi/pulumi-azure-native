


package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionControllerPrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult, error) {
	var rv LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:migrate/v20200501:getPrivateEndpointConnectionControllerPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionArgs struct {
	MigrateProjectName string `pulumi:"migrateProjectName"`
	PeConnectionName   string `pulumi:"peConnectionName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult struct {
	ETag       string                                      `pulumi:"eTag"`
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                          `pulumi:"systemData"`
	Type       string                                      `pulumi:"type"`
}

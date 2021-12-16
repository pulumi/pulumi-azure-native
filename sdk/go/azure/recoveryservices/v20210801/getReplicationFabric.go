


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationFabric(ctx *pulumi.Context, args *LookupReplicationFabricArgs, opts ...pulumi.InvokeOption) (*LookupReplicationFabricResult, error) {
	var rv LookupReplicationFabricResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210801:getReplicationFabric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationFabricArgs struct {
	FabricName        string  `pulumi:"fabricName"`
	Filter            *string `pulumi:"filter"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceName      string  `pulumi:"resourceName"`
}


type LookupReplicationFabricResult struct {
	Id         string                   `pulumi:"id"`
	Location   *string                  `pulumi:"location"`
	Name       string                   `pulumi:"name"`
	Properties FabricPropertiesResponse `pulumi:"properties"`
	Type       string                   `pulumi:"type"`
}

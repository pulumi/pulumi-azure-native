


package recoveryservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationProtectionContainerMapping(ctx *pulumi.Context, args *LookupReplicationProtectionContainerMappingArgs, opts ...pulumi.InvokeOption) (*LookupReplicationProtectionContainerMappingResult, error) {
	var rv LookupReplicationProtectionContainerMappingResult
	err := ctx.Invoke("azure-native:recoveryservices:getReplicationProtectionContainerMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationProtectionContainerMappingArgs struct {
	FabricName              string `pulumi:"fabricName"`
	MappingName             string `pulumi:"mappingName"`
	ProtectionContainerName string `pulumi:"protectionContainerName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	ResourceName            string `pulumi:"resourceName"`
}


type LookupReplicationProtectionContainerMappingResult struct {
	Id         string                                       `pulumi:"id"`
	Location   *string                                      `pulumi:"location"`
	Name       string                                       `pulumi:"name"`
	Properties ProtectionContainerMappingPropertiesResponse `pulumi:"properties"`
	Type       string                                       `pulumi:"type"`
}

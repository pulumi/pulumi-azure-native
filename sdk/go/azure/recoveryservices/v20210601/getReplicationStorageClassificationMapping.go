


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationStorageClassificationMapping(ctx *pulumi.Context, args *LookupReplicationStorageClassificationMappingArgs, opts ...pulumi.InvokeOption) (*LookupReplicationStorageClassificationMappingResult, error) {
	var rv LookupReplicationStorageClassificationMappingResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210601:getReplicationStorageClassificationMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationStorageClassificationMappingArgs struct {
	FabricName                       string `pulumi:"fabricName"`
	ResourceGroupName                string `pulumi:"resourceGroupName"`
	ResourceName                     string `pulumi:"resourceName"`
	StorageClassificationMappingName string `pulumi:"storageClassificationMappingName"`
	StorageClassificationName        string `pulumi:"storageClassificationName"`
}


type LookupReplicationStorageClassificationMappingResult struct {
	Id         string                                         `pulumi:"id"`
	Location   *string                                        `pulumi:"location"`
	Name       string                                         `pulumi:"name"`
	Properties StorageClassificationMappingPropertiesResponse `pulumi:"properties"`
	Type       string                                         `pulumi:"type"`
}

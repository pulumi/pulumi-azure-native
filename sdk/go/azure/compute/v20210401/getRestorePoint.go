


package v20210401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRestorePoint(ctx *pulumi.Context, args *LookupRestorePointArgs, opts ...pulumi.InvokeOption) (*LookupRestorePointResult, error) {
	var rv LookupRestorePointResult
	err := ctx.Invoke("azure-native:compute/v20210401:getRestorePoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRestorePointArgs struct {
	ResourceGroupName          string `pulumi:"resourceGroupName"`
	RestorePointCollectionName string `pulumi:"restorePointCollectionName"`
	RestorePointName           string `pulumi:"restorePointName"`
}


type LookupRestorePointResult struct {
	ConsistencyMode   string                             `pulumi:"consistencyMode"`
	ExcludeDisks      []ApiEntityReferenceResponse       `pulumi:"excludeDisks"`
	Id                string                             `pulumi:"id"`
	Name              string                             `pulumi:"name"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	SourceMetadata    RestorePointSourceMetadataResponse `pulumi:"sourceMetadata"`
	TimeCreated       *string                            `pulumi:"timeCreated"`
	Type              string                             `pulumi:"type"`
}

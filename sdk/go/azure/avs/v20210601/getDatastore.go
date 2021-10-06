


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatastore(ctx *pulumi.Context, args *LookupDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupDatastoreResult, error) {
	var rv LookupDatastoreResult
	err := ctx.Invoke("azure-native:avs/v20210601:getDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatastoreArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatastoreName     string `pulumi:"datastoreName"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatastoreResult struct {
	DiskPoolVolume    *DiskPoolVolumeResponse `pulumi:"diskPoolVolume"`
	Id                string                  `pulumi:"id"`
	Name              string                  `pulumi:"name"`
	NetAppVolume      *NetAppVolumeResponse   `pulumi:"netAppVolume"`
	ProvisioningState string                  `pulumi:"provisioningState"`
	Type              string                  `pulumi:"type"`
}

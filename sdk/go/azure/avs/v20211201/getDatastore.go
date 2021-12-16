


package v20211201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatastore(ctx *pulumi.Context, args *LookupDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupDatastoreResult, error) {
	var rv LookupDatastoreResult
	err := ctx.Invoke("azure-native:avs/v20211201:getDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
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
	Status            string                  `pulumi:"status"`
	Type              string                  `pulumi:"type"`
}


func (val *LookupDatastoreResult) Defaults() *LookupDatastoreResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DiskPoolVolume = tmp.DiskPoolVolume.Defaults()

	return &tmp
}

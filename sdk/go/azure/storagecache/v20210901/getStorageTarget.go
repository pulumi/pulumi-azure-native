


package v20210901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageTarget(ctx *pulumi.Context, args *LookupStorageTargetArgs, opts ...pulumi.InvokeOption) (*LookupStorageTargetResult, error) {
	var rv LookupStorageTargetResult
	err := ctx.Invoke("azure-native:storagecache/v20210901:getStorageTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageTargetArgs struct {
	CacheName         string `pulumi:"cacheName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StorageTargetName string `pulumi:"storageTargetName"`
}


type LookupStorageTargetResult struct {
	BlobNfs           *BlobNfsTargetResponse      `pulumi:"blobNfs"`
	Clfs              *ClfsTargetResponse         `pulumi:"clfs"`
	Id                string                      `pulumi:"id"`
	Junctions         []NamespaceJunctionResponse `pulumi:"junctions"`
	Location          string                      `pulumi:"location"`
	Name              string                      `pulumi:"name"`
	Nfs3              *Nfs3TargetResponse         `pulumi:"nfs3"`
	ProvisioningState string                      `pulumi:"provisioningState"`
	State             *string                     `pulumi:"state"`
	SystemData        SystemDataResponse          `pulumi:"systemData"`
	TargetType        string                      `pulumi:"targetType"`
	Type              string                      `pulumi:"type"`
	Unknown           *UnknownTargetResponse      `pulumi:"unknown"`
}

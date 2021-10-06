


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageTarget(ctx *pulumi.Context, args *LookupStorageTargetArgs, opts ...pulumi.InvokeOption) (*LookupStorageTargetResult, error) {
	var rv LookupStorageTargetResult
	err := ctx.Invoke("azure-native:storagecache/v20191101:getStorageTarget", args, &rv, opts...)
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
	Clfs              *ClfsTargetResponse         `pulumi:"clfs"`
	Id                string                      `pulumi:"id"`
	Junctions         []NamespaceJunctionResponse `pulumi:"junctions"`
	Name              string                      `pulumi:"name"`
	Nfs3              *Nfs3TargetResponse         `pulumi:"nfs3"`
	ProvisioningState *string                     `pulumi:"provisioningState"`
	TargetType        *string                     `pulumi:"targetType"`
	Type              string                      `pulumi:"type"`
	Unknown           *UnknownTargetResponse      `pulumi:"unknown"`
}

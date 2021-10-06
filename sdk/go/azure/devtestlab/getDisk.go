


package devtestlab

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDisk(ctx *pulumi.Context, args *LookupDiskArgs, opts ...pulumi.InvokeOption) (*LookupDiskResult, error) {
	var rv LookupDiskResult
	err := ctx.Invoke("azure-native:devtestlab:getDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	UserName          string  `pulumi:"userName"`
}


type LookupDiskResult struct {
	CreatedDate       string            `pulumi:"createdDate"`
	DiskBlobName      *string           `pulumi:"diskBlobName"`
	DiskSizeGiB       *int              `pulumi:"diskSizeGiB"`
	DiskType          *string           `pulumi:"diskType"`
	DiskUri           *string           `pulumi:"diskUri"`
	HostCaching       *string           `pulumi:"hostCaching"`
	Id                string            `pulumi:"id"`
	LeasedByLabVmId   *string           `pulumi:"leasedByLabVmId"`
	Location          *string           `pulumi:"location"`
	ManagedDiskId     *string           `pulumi:"managedDiskId"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	StorageAccountId  *string           `pulumi:"storageAccountId"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
	UniqueIdentifier  string            `pulumi:"uniqueIdentifier"`
}

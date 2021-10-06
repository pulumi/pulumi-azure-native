


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiskPool(ctx *pulumi.Context, args *LookupDiskPoolArgs, opts ...pulumi.InvokeOption) (*LookupDiskPoolResult, error) {
	var rv LookupDiskPoolResult
	err := ctx.Invoke("azure-native:storagepool/v20210801:getDiskPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskPoolArgs struct {
	DiskPoolName      string `pulumi:"diskPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDiskPoolResult struct {
	AdditionalCapabilities []string               `pulumi:"additionalCapabilities"`
	AvailabilityZones      []string               `pulumi:"availabilityZones"`
	Disks                  []DiskResponse         `pulumi:"disks"`
	Id                     string                 `pulumi:"id"`
	Location               string                 `pulumi:"location"`
	ManagedBy              string                 `pulumi:"managedBy"`
	ManagedByExtended      []string               `pulumi:"managedByExtended"`
	Name                   string                 `pulumi:"name"`
	ProvisioningState      string                 `pulumi:"provisioningState"`
	Status                 string                 `pulumi:"status"`
	SubnetId               string                 `pulumi:"subnetId"`
	SystemData             SystemMetadataResponse `pulumi:"systemData"`
	Tags                   map[string]string      `pulumi:"tags"`
	Tier                   *string                `pulumi:"tier"`
	Type                   string                 `pulumi:"type"`
}

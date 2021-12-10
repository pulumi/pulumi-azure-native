


package v20161001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIscsiDisk(ctx *pulumi.Context, args *LookupIscsiDiskArgs, opts ...pulumi.InvokeOption) (*LookupIscsiDiskResult, error) {
	var rv LookupIscsiDiskResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getIscsiDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIscsiDiskArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	DiskName          string `pulumi:"diskName"`
	IscsiServerName   string `pulumi:"iscsiServerName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupIscsiDiskResult struct {
	AccessControlRecords       []string `pulumi:"accessControlRecords"`
	DataPolicy                 string   `pulumi:"dataPolicy"`
	Description                *string  `pulumi:"description"`
	DiskStatus                 string   `pulumi:"diskStatus"`
	Id                         string   `pulumi:"id"`
	LocalUsedCapacityInBytes   float64  `pulumi:"localUsedCapacityInBytes"`
	MonitoringStatus           string   `pulumi:"monitoringStatus"`
	Name                       string   `pulumi:"name"`
	ProvisionedCapacityInBytes float64  `pulumi:"provisionedCapacityInBytes"`
	Type                       string   `pulumi:"type"`
	UsedCapacityInBytes        float64  `pulumi:"usedCapacityInBytes"`
}

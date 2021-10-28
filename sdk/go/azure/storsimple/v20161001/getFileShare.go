


package v20161001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFileShare(ctx *pulumi.Context, args *LookupFileShareArgs, opts ...pulumi.InvokeOption) (*LookupFileShareResult, error) {
	var rv LookupFileShareResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getFileShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileShareArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	FileServerName    string `pulumi:"fileServerName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupFileShareResult struct {
	AdminUser                  string  `pulumi:"adminUser"`
	DataPolicy                 string  `pulumi:"dataPolicy"`
	Description                *string `pulumi:"description"`
	Id                         string  `pulumi:"id"`
	LocalUsedCapacityInBytes   float64 `pulumi:"localUsedCapacityInBytes"`
	MonitoringStatus           string  `pulumi:"monitoringStatus"`
	Name                       string  `pulumi:"name"`
	ProvisionedCapacityInBytes float64 `pulumi:"provisionedCapacityInBytes"`
	ShareStatus                string  `pulumi:"shareStatus"`
	Type                       string  `pulumi:"type"`
	UsedCapacityInBytes        float64 `pulumi:"usedCapacityInBytes"`
}

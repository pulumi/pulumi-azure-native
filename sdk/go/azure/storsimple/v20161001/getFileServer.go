


package v20161001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFileServer(ctx *pulumi.Context, args *LookupFileServerArgs, opts ...pulumi.InvokeOption) (*LookupFileServerResult, error) {
	var rv LookupFileServerResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getFileServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileServerArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	FileServerName    string `pulumi:"fileServerName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFileServerResult struct {
	BackupScheduleGroupId string  `pulumi:"backupScheduleGroupId"`
	Description           *string `pulumi:"description"`
	DomainName            string  `pulumi:"domainName"`
	Id                    string  `pulumi:"id"`
	Name                  string  `pulumi:"name"`
	StorageDomainId       string  `pulumi:"storageDomainId"`
	Type                  string  `pulumi:"type"`
}

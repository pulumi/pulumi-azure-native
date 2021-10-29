


package v20161001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIscsiServer(ctx *pulumi.Context, args *LookupIscsiServerArgs, opts ...pulumi.InvokeOption) (*LookupIscsiServerResult, error) {
	var rv LookupIscsiServerResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getIscsiServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIscsiServerArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	IscsiServerName   string `pulumi:"iscsiServerName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupIscsiServerResult struct {
	BackupScheduleGroupId string  `pulumi:"backupScheduleGroupId"`
	ChapId                *string `pulumi:"chapId"`
	Description           *string `pulumi:"description"`
	Id                    string  `pulumi:"id"`
	Name                  string  `pulumi:"name"`
	ReverseChapId         *string `pulumi:"reverseChapId"`
	StorageDomainId       string  `pulumi:"storageDomainId"`
	Type                  string  `pulumi:"type"`
}

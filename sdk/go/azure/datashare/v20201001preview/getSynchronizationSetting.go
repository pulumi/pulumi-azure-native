


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSynchronizationSetting(ctx *pulumi.Context, args *LookupSynchronizationSettingArgs, opts ...pulumi.InvokeOption) (*LookupSynchronizationSettingResult, error) {
	var rv LookupSynchronizationSettingResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getSynchronizationSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSynchronizationSettingArgs struct {
	AccountName                string `pulumi:"accountName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
	ShareName                  string `pulumi:"shareName"`
	SynchronizationSettingName string `pulumi:"synchronizationSettingName"`
}


type LookupSynchronizationSettingResult struct {
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

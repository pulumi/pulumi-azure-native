


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScheduledSynchronizationSetting(ctx *pulumi.Context, args *LookupScheduledSynchronizationSettingArgs, opts ...pulumi.InvokeOption) (*LookupScheduledSynchronizationSettingResult, error) {
	var rv LookupScheduledSynchronizationSettingResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getScheduledSynchronizationSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledSynchronizationSettingArgs struct {
	AccountName                string `pulumi:"accountName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
	ShareName                  string `pulumi:"shareName"`
	SynchronizationSettingName string `pulumi:"synchronizationSettingName"`
}


type LookupScheduledSynchronizationSettingResult struct {
	CreatedAt           string             `pulumi:"createdAt"`
	Id                  string             `pulumi:"id"`
	Kind                string             `pulumi:"kind"`
	Name                string             `pulumi:"name"`
	ProvisioningState   string             `pulumi:"provisioningState"`
	RecurrenceInterval  string             `pulumi:"recurrenceInterval"`
	SynchronizationTime string             `pulumi:"synchronizationTime"`
	SystemData          SystemDataResponse `pulumi:"systemData"`
	Type                string             `pulumi:"type"`
	UserName            string             `pulumi:"userName"`
}

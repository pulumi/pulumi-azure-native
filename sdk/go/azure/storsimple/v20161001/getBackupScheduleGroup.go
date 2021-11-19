


package v20161001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupScheduleGroup(ctx *pulumi.Context, args *LookupBackupScheduleGroupArgs, opts ...pulumi.InvokeOption) (*LookupBackupScheduleGroupResult, error) {
	var rv LookupBackupScheduleGroupResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getBackupScheduleGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupScheduleGroupArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScheduleGroupName string `pulumi:"scheduleGroupName"`
}


type LookupBackupScheduleGroupResult struct {
	Id        string       `pulumi:"id"`
	Name      string       `pulumi:"name"`
	StartTime TimeResponse `pulumi:"startTime"`
	Type      string       `pulumi:"type"`
}

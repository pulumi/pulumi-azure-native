


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSnapshotPolicy(ctx *pulumi.Context, args *LookupSnapshotPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotPolicyResult, error) {
	var rv LookupSnapshotPolicyResult
	err := ctx.Invoke("azure-native:netapp/v20210401preview:getSnapshotPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSnapshotPolicyArgs struct {
	AccountName        string `pulumi:"accountName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	SnapshotPolicyName string `pulumi:"snapshotPolicyName"`
}


type LookupSnapshotPolicyResult struct {
	DailySchedule     *DailyScheduleResponse   `pulumi:"dailySchedule"`
	Enabled           *bool                    `pulumi:"enabled"`
	HourlySchedule    *HourlyScheduleResponse  `pulumi:"hourlySchedule"`
	Id                string                   `pulumi:"id"`
	Location          string                   `pulumi:"location"`
	MonthlySchedule   *MonthlyScheduleResponse `pulumi:"monthlySchedule"`
	Name              string                   `pulumi:"name"`
	ProvisioningState string                   `pulumi:"provisioningState"`
	Tags              map[string]string        `pulumi:"tags"`
	Type              string                   `pulumi:"type"`
	WeeklySchedule    *WeeklyScheduleResponse  `pulumi:"weeklySchedule"`
}

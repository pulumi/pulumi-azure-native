


package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSchedule(ctx *pulumi.Context, args *LookupScheduleArgs, opts ...pulumi.InvokeOption) (*LookupScheduleResult, error) {
	var rv LookupScheduleResult
	err := ctx.Invoke("azure-native:labservices:getSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduleArgs struct {
	LabName           string `pulumi:"labName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScheduleName      string `pulumi:"scheduleName"`
}


type LookupScheduleResult struct {
	Id                string                     `pulumi:"id"`
	Name              string                     `pulumi:"name"`
	Notes             *string                    `pulumi:"notes"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	RecurrencePattern *RecurrencePatternResponse `pulumi:"recurrencePattern"`
	StartAt           *string                    `pulumi:"startAt"`
	StopAt            string                     `pulumi:"stopAt"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	TimeZoneId        string                     `pulumi:"timeZoneId"`
	Type              string                     `pulumi:"type"`
}

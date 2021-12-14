


package v20151031

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSchedule(ctx *pulumi.Context, args *LookupScheduleArgs, opts ...pulumi.InvokeOption) (*LookupScheduleResult, error) {
	var rv LookupScheduleResult
	err := ctx.Invoke("azure-native:automation/v20151031:getSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupScheduleArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ScheduleName          string `pulumi:"scheduleName"`
}


type LookupScheduleResult struct {
	AdvancedSchedule        *AdvancedScheduleResponse `pulumi:"advancedSchedule"`
	CreationTime            *string                   `pulumi:"creationTime"`
	Description             *string                   `pulumi:"description"`
	ExpiryTime              *string                   `pulumi:"expiryTime"`
	ExpiryTimeOffsetMinutes *float64                  `pulumi:"expiryTimeOffsetMinutes"`
	Frequency               *string                   `pulumi:"frequency"`
	Id                      string                    `pulumi:"id"`
	Interval                interface{}               `pulumi:"interval"`
	IsEnabled               *bool                     `pulumi:"isEnabled"`
	LastModifiedTime        *string                   `pulumi:"lastModifiedTime"`
	Name                    string                    `pulumi:"name"`
	NextRun                 *string                   `pulumi:"nextRun"`
	NextRunOffsetMinutes    *float64                  `pulumi:"nextRunOffsetMinutes"`
	StartTime               *string                   `pulumi:"startTime"`
	StartTimeOffsetMinutes  float64                   `pulumi:"startTimeOffsetMinutes"`
	TimeZone                *string                   `pulumi:"timeZone"`
	Type                    string                    `pulumi:"type"`
}


func (val *LookupScheduleResult) Defaults() *LookupScheduleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsEnabled) {
		isEnabled_ := false
		tmp.IsEnabled = &isEnabled_
	}
	return &tmp
}

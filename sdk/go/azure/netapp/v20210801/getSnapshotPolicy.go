


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSnapshotPolicy(ctx *pulumi.Context, args *LookupSnapshotPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotPolicyResult, error) {
	var rv LookupSnapshotPolicyResult
	err := ctx.Invoke("azure-native:netapp/v20210801:getSnapshotPolicy", args, &rv, opts...)
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
	Etag              string                   `pulumi:"etag"`
	HourlySchedule    *HourlyScheduleResponse  `pulumi:"hourlySchedule"`
	Id                string                   `pulumi:"id"`
	Location          string                   `pulumi:"location"`
	MonthlySchedule   *MonthlyScheduleResponse `pulumi:"monthlySchedule"`
	Name              string                   `pulumi:"name"`
	ProvisioningState string                   `pulumi:"provisioningState"`
	SystemData        SystemDataResponse       `pulumi:"systemData"`
	Tags              map[string]string        `pulumi:"tags"`
	Type              string                   `pulumi:"type"`
	WeeklySchedule    *WeeklyScheduleResponse  `pulumi:"weeklySchedule"`
}

func LookupSnapshotPolicyOutput(ctx *pulumi.Context, args LookupSnapshotPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupSnapshotPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSnapshotPolicyResult, error) {
			args := v.(LookupSnapshotPolicyArgs)
			r, err := LookupSnapshotPolicy(ctx, &args, opts...)
			var s LookupSnapshotPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSnapshotPolicyResultOutput)
}

type LookupSnapshotPolicyOutputArgs struct {
	AccountName        pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	SnapshotPolicyName pulumi.StringInput `pulumi:"snapshotPolicyName"`
}

func (LookupSnapshotPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotPolicyArgs)(nil)).Elem()
}


type LookupSnapshotPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupSnapshotPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotPolicyResult)(nil)).Elem()
}

func (o LookupSnapshotPolicyResultOutput) ToLookupSnapshotPolicyResultOutput() LookupSnapshotPolicyResultOutput {
	return o
}

func (o LookupSnapshotPolicyResultOutput) ToLookupSnapshotPolicyResultOutputWithContext(ctx context.Context) LookupSnapshotPolicyResultOutput {
	return o
}

func (o LookupSnapshotPolicyResultOutput) DailySchedule() DailyScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) *DailyScheduleResponse { return v.DailySchedule }).(DailyScheduleResponsePtrOutput)
}

func (o LookupSnapshotPolicyResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSnapshotPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupSnapshotPolicyResultOutput) HourlySchedule() HourlyScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) *HourlyScheduleResponse { return v.HourlySchedule }).(HourlyScheduleResponsePtrOutput)
}

func (o LookupSnapshotPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSnapshotPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSnapshotPolicyResultOutput) MonthlySchedule() MonthlyScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) *MonthlyScheduleResponse { return v.MonthlySchedule }).(MonthlyScheduleResponsePtrOutput)
}

func (o LookupSnapshotPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSnapshotPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSnapshotPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSnapshotPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSnapshotPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSnapshotPolicyResultOutput) WeeklySchedule() WeeklyScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) *WeeklyScheduleResponse { return v.WeeklySchedule }).(WeeklyScheduleResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnapshotPolicyResultOutput{})
}

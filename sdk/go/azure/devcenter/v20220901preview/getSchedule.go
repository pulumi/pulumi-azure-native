


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSchedule(ctx *pulumi.Context, args *LookupScheduleArgs, opts ...pulumi.InvokeOption) (*LookupScheduleResult, error) {
	var rv LookupScheduleResult
	err := ctx.Invoke("azure-native:devcenter/v20220901preview:getSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduleArgs struct {
	PoolName          string `pulumi:"poolName"`
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScheduleName      string `pulumi:"scheduleName"`
	Top               *int   `pulumi:"top"`
}


type LookupScheduleResult struct {
	Frequency         string             `pulumi:"frequency"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	State             *string            `pulumi:"state"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Time              string             `pulumi:"time"`
	TimeZone          string             `pulumi:"timeZone"`
	Type              string             `pulumi:"type"`
}

func LookupScheduleOutput(ctx *pulumi.Context, args LookupScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduleResult, error) {
			args := v.(LookupScheduleArgs)
			r, err := LookupSchedule(ctx, &args, opts...)
			var s LookupScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduleResultOutput)
}

type LookupScheduleOutputArgs struct {
	PoolName          pulumi.StringInput `pulumi:"poolName"`
	ProjectName       pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ScheduleName      pulumi.StringInput `pulumi:"scheduleName"`
	Top               pulumi.IntPtrInput `pulumi:"top"`
}

func (LookupScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleArgs)(nil)).Elem()
}


type LookupScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleResult)(nil)).Elem()
}

func (o LookupScheduleResultOutput) ToLookupScheduleResultOutput() LookupScheduleResultOutput {
	return o
}

func (o LookupScheduleResultOutput) ToLookupScheduleResultOutputWithContext(ctx context.Context) LookupScheduleResultOutput {
	return o
}

func (o LookupScheduleResultOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Frequency }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScheduleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupScheduleResultOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Time }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduleResultOutput{})
}

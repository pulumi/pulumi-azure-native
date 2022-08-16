


package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobSchedule(ctx *pulumi.Context, args *LookupJobScheduleArgs, opts ...pulumi.InvokeOption) (*LookupJobScheduleResult, error) {
	var rv LookupJobScheduleResult
	err := ctx.Invoke("azure-native:automation:getJobSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobScheduleArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	JobScheduleId         string `pulumi:"jobScheduleId"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupJobScheduleResult struct {
	Id            string                               `pulumi:"id"`
	JobScheduleId *string                              `pulumi:"jobScheduleId"`
	Name          string                               `pulumi:"name"`
	Parameters    map[string]string                    `pulumi:"parameters"`
	RunOn         *string                              `pulumi:"runOn"`
	Runbook       *RunbookAssociationPropertyResponse  `pulumi:"runbook"`
	Schedule      *ScheduleAssociationPropertyResponse `pulumi:"schedule"`
	Type          string                               `pulumi:"type"`
}

func LookupJobScheduleOutput(ctx *pulumi.Context, args LookupJobScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupJobScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobScheduleResult, error) {
			args := v.(LookupJobScheduleArgs)
			r, err := LookupJobSchedule(ctx, &args, opts...)
			var s LookupJobScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobScheduleResultOutput)
}

type LookupJobScheduleOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	JobScheduleId         pulumi.StringInput `pulumi:"jobScheduleId"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupJobScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobScheduleArgs)(nil)).Elem()
}


type LookupJobScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupJobScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobScheduleResult)(nil)).Elem()
}

func (o LookupJobScheduleResultOutput) ToLookupJobScheduleResultOutput() LookupJobScheduleResultOutput {
	return o
}

func (o LookupJobScheduleResultOutput) ToLookupJobScheduleResultOutputWithContext(ctx context.Context) LookupJobScheduleResultOutput {
	return o
}

func (o LookupJobScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobScheduleResultOutput) JobScheduleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) *string { return v.JobScheduleId }).(pulumi.StringPtrOutput)
}

func (o LookupJobScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJobScheduleResultOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o LookupJobScheduleResultOutput) RunOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) *string { return v.RunOn }).(pulumi.StringPtrOutput)
}

func (o LookupJobScheduleResultOutput) Runbook() RunbookAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) *RunbookAssociationPropertyResponse { return v.Runbook }).(RunbookAssociationPropertyResponsePtrOutput)
}

func (o LookupJobScheduleResultOutput) Schedule() ScheduleAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) *ScheduleAssociationPropertyResponse { return v.Schedule }).(ScheduleAssociationPropertyResponsePtrOutput)
}

func (o LookupJobScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobScheduleResultOutput{})
}

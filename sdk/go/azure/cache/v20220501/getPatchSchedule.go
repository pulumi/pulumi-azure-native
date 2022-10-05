


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPatchSchedule(ctx *pulumi.Context, args *LookupPatchScheduleArgs, opts ...pulumi.InvokeOption) (*LookupPatchScheduleResult, error) {
	var rv LookupPatchScheduleResult
	err := ctx.Invoke("azure-native:cache/v20220501:getPatchSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPatchScheduleArgs struct {
	Default           string `pulumi:"default"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPatchScheduleResult struct {
	Id              string                  `pulumi:"id"`
	Location        string                  `pulumi:"location"`
	Name            string                  `pulumi:"name"`
	ScheduleEntries []ScheduleEntryResponse `pulumi:"scheduleEntries"`
	Type            string                  `pulumi:"type"`
}

func LookupPatchScheduleOutput(ctx *pulumi.Context, args LookupPatchScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupPatchScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPatchScheduleResult, error) {
			args := v.(LookupPatchScheduleArgs)
			r, err := LookupPatchSchedule(ctx, &args, opts...)
			var s LookupPatchScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPatchScheduleResultOutput)
}

type LookupPatchScheduleOutputArgs struct {
	Default           pulumi.StringInput `pulumi:"default"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPatchScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPatchScheduleArgs)(nil)).Elem()
}


type LookupPatchScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupPatchScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPatchScheduleResult)(nil)).Elem()
}

func (o LookupPatchScheduleResultOutput) ToLookupPatchScheduleResultOutput() LookupPatchScheduleResultOutput {
	return o
}

func (o LookupPatchScheduleResultOutput) ToLookupPatchScheduleResultOutputWithContext(ctx context.Context) LookupPatchScheduleResultOutput {
	return o
}

func (o LookupPatchScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPatchScheduleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchScheduleResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPatchScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPatchScheduleResultOutput) ScheduleEntries() ScheduleEntryResponseArrayOutput {
	return o.ApplyT(func(v LookupPatchScheduleResult) []ScheduleEntryResponse { return v.ScheduleEntries }).(ScheduleEntryResponseArrayOutput)
}

func (o LookupPatchScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPatchScheduleResultOutput{})
}

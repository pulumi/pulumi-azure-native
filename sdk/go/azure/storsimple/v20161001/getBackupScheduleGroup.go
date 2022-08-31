


package v20161001

import (
	"context"
	"reflect"

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

func LookupBackupScheduleGroupOutput(ctx *pulumi.Context, args LookupBackupScheduleGroupOutputArgs, opts ...pulumi.InvokeOption) LookupBackupScheduleGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupScheduleGroupResult, error) {
			args := v.(LookupBackupScheduleGroupArgs)
			r, err := LookupBackupScheduleGroup(ctx, &args, opts...)
			var s LookupBackupScheduleGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupScheduleGroupResultOutput)
}

type LookupBackupScheduleGroupOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ScheduleGroupName pulumi.StringInput `pulumi:"scheduleGroupName"`
}

func (LookupBackupScheduleGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupScheduleGroupArgs)(nil)).Elem()
}


type LookupBackupScheduleGroupResultOutput struct{ *pulumi.OutputState }

func (LookupBackupScheduleGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupScheduleGroupResult)(nil)).Elem()
}

func (o LookupBackupScheduleGroupResultOutput) ToLookupBackupScheduleGroupResultOutput() LookupBackupScheduleGroupResultOutput {
	return o
}

func (o LookupBackupScheduleGroupResultOutput) ToLookupBackupScheduleGroupResultOutputWithContext(ctx context.Context) LookupBackupScheduleGroupResultOutput {
	return o
}

func (o LookupBackupScheduleGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupScheduleGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBackupScheduleGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupScheduleGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBackupScheduleGroupResultOutput) StartTime() TimeResponseOutput {
	return o.ApplyT(func(v LookupBackupScheduleGroupResult) TimeResponse { return v.StartTime }).(TimeResponseOutput)
}

func (o LookupBackupScheduleGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupScheduleGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupScheduleGroupResultOutput{})
}

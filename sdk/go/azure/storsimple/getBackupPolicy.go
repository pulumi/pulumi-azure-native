


package storsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupPolicy(ctx *pulumi.Context, args *LookupBackupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupPolicyResult, error) {
	var rv LookupBackupPolicyResult
	err := ctx.Invoke("azure-native:storsimple:getBackupPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupPolicyArgs struct {
	BackupPolicyName  string `pulumi:"backupPolicyName"`
	DeviceName        string `pulumi:"deviceName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBackupPolicyResult struct {
	BackupPolicyCreationType string   `pulumi:"backupPolicyCreationType"`
	Id                       string   `pulumi:"id"`
	Kind                     *string  `pulumi:"kind"`
	LastBackupTime           string   `pulumi:"lastBackupTime"`
	Name                     string   `pulumi:"name"`
	NextBackupTime           string   `pulumi:"nextBackupTime"`
	ScheduledBackupStatus    string   `pulumi:"scheduledBackupStatus"`
	SchedulesCount           float64  `pulumi:"schedulesCount"`
	SsmHostName              string   `pulumi:"ssmHostName"`
	Type                     string   `pulumi:"type"`
	VolumeIds                []string `pulumi:"volumeIds"`
}

func LookupBackupPolicyOutput(ctx *pulumi.Context, args LookupBackupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBackupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupPolicyResult, error) {
			args := v.(LookupBackupPolicyArgs)
			r, err := LookupBackupPolicy(ctx, &args, opts...)
			var s LookupBackupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupPolicyResultOutput)
}

type LookupBackupPolicyOutputArgs struct {
	BackupPolicyName  pulumi.StringInput `pulumi:"backupPolicyName"`
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBackupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupPolicyArgs)(nil)).Elem()
}


type LookupBackupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBackupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupPolicyResult)(nil)).Elem()
}

func (o LookupBackupPolicyResultOutput) ToLookupBackupPolicyResultOutput() LookupBackupPolicyResultOutput {
	return o
}

func (o LookupBackupPolicyResultOutput) ToLookupBackupPolicyResultOutputWithContext(ctx context.Context) LookupBackupPolicyResultOutput {
	return o
}

func (o LookupBackupPolicyResultOutput) BackupPolicyCreationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.BackupPolicyCreationType }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupBackupPolicyResultOutput) LastBackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.LastBackupTime }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) NextBackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.NextBackupTime }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) ScheduledBackupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.ScheduledBackupStatus }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) SchedulesCount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupBackupPolicyResult) float64 { return v.SchedulesCount }).(pulumi.Float64Output)
}

func (o LookupBackupPolicyResultOutput) SsmHostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.SsmHostName }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) VolumeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) []string { return v.VolumeIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupPolicyResultOutput{})
}

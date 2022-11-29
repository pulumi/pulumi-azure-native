


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupPolicy(ctx *pulumi.Context, args *LookupBackupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupPolicyResult, error) {
	var rv LookupBackupPolicyResult
	err := ctx.Invoke("azure-native:netapp/v20210601:getBackupPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupPolicyArgs struct {
	AccountName       string `pulumi:"accountName"`
	BackupPolicyName  string `pulumi:"backupPolicyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBackupPolicyResult struct {
	BackupPolicyId       string                  `pulumi:"backupPolicyId"`
	DailyBackupsToKeep   *int                    `pulumi:"dailyBackupsToKeep"`
	Enabled              *bool                   `pulumi:"enabled"`
	Etag                 string                  `pulumi:"etag"`
	Id                   string                  `pulumi:"id"`
	Location             string                  `pulumi:"location"`
	MonthlyBackupsToKeep *int                    `pulumi:"monthlyBackupsToKeep"`
	Name                 string                  `pulumi:"name"`
	ProvisioningState    string                  `pulumi:"provisioningState"`
	Tags                 map[string]string       `pulumi:"tags"`
	Type                 string                  `pulumi:"type"`
	VolumeBackups        []VolumeBackupsResponse `pulumi:"volumeBackups"`
	VolumesAssigned      int                     `pulumi:"volumesAssigned"`
	WeeklyBackupsToKeep  *int                    `pulumi:"weeklyBackupsToKeep"`
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
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	BackupPolicyName  pulumi.StringInput `pulumi:"backupPolicyName"`
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

func (o LookupBackupPolicyResultOutput) BackupPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.BackupPolicyId }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) DailyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) *int { return v.DailyBackupsToKeep }).(pulumi.IntPtrOutput)
}

func (o LookupBackupPolicyResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupBackupPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) MonthlyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) *int { return v.MonthlyBackupsToKeep }).(pulumi.IntPtrOutput)
}

func (o LookupBackupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBackupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBackupPolicyResultOutput) VolumeBackups() VolumeBackupsResponseArrayOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) []VolumeBackupsResponse { return v.VolumeBackups }).(VolumeBackupsResponseArrayOutput)
}

func (o LookupBackupPolicyResultOutput) VolumesAssigned() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) int { return v.VolumesAssigned }).(pulumi.IntOutput)
}

func (o LookupBackupPolicyResultOutput) WeeklyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) *int { return v.WeeklyBackupsToKeep }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupPolicyResultOutput{})
}

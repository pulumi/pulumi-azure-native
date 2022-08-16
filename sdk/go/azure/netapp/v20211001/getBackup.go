


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackup(ctx *pulumi.Context, args *LookupBackupArgs, opts ...pulumi.InvokeOption) (*LookupBackupResult, error) {
	var rv LookupBackupResult
	err := ctx.Invoke("azure-native:netapp/v20211001:getBackup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBackupArgs struct {
	AccountName       string `pulumi:"accountName"`
	BackupName        string `pulumi:"backupName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VolumeName        string `pulumi:"volumeName"`
}


type LookupBackupResult struct {
	BackupId            string  `pulumi:"backupId"`
	BackupType          string  `pulumi:"backupType"`
	CreationDate        string  `pulumi:"creationDate"`
	FailureReason       string  `pulumi:"failureReason"`
	Id                  string  `pulumi:"id"`
	Label               *string `pulumi:"label"`
	Location            string  `pulumi:"location"`
	Name                string  `pulumi:"name"`
	ProvisioningState   string  `pulumi:"provisioningState"`
	Size                float64 `pulumi:"size"`
	Type                string  `pulumi:"type"`
	UseExistingSnapshot *bool   `pulumi:"useExistingSnapshot"`
	VolumeName          string  `pulumi:"volumeName"`
}


func (val *LookupBackupResult) Defaults() *LookupBackupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseExistingSnapshot) {
		useExistingSnapshot_ := false
		tmp.UseExistingSnapshot = &useExistingSnapshot_
	}
	return &tmp
}

func LookupBackupOutput(ctx *pulumi.Context, args LookupBackupOutputArgs, opts ...pulumi.InvokeOption) LookupBackupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupResult, error) {
			args := v.(LookupBackupArgs)
			r, err := LookupBackup(ctx, &args, opts...)
			var s LookupBackupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupResultOutput)
}

type LookupBackupOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	BackupName        pulumi.StringInput `pulumi:"backupName"`
	PoolName          pulumi.StringInput `pulumi:"poolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VolumeName        pulumi.StringInput `pulumi:"volumeName"`
}

func (LookupBackupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupArgs)(nil)).Elem()
}


type LookupBackupResultOutput struct{ *pulumi.OutputState }

func (LookupBackupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupResult)(nil)).Elem()
}

func (o LookupBackupResultOutput) ToLookupBackupResultOutput() LookupBackupResultOutput {
	return o
}

func (o LookupBackupResultOutput) ToLookupBackupResultOutputWithContext(ctx context.Context) LookupBackupResultOutput {
	return o
}

func (o LookupBackupResultOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.BackupId }).(pulumi.StringOutput)
}

func (o LookupBackupResultOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.BackupType }).(pulumi.StringOutput)
}

func (o LookupBackupResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupBackupResultOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.FailureReason }).(pulumi.StringOutput)
}

func (o LookupBackupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBackupResultOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupResult) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o LookupBackupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBackupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBackupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupBackupResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v LookupBackupResult) float64 { return v.Size }).(pulumi.Float64Output)
}

func (o LookupBackupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBackupResultOutput) UseExistingSnapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBackupResult) *bool { return v.UseExistingSnapshot }).(pulumi.BoolPtrOutput)
}

func (o LookupBackupResultOutput) VolumeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.VolumeName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupResultOutput{})
}

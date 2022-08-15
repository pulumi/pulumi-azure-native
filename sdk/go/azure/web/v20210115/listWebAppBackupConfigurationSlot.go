


package v20210115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppBackupConfigurationSlot(ctx *pulumi.Context, args *ListWebAppBackupConfigurationSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppBackupConfigurationSlotResult, error) {
	var rv ListWebAppBackupConfigurationSlotResult
	err := ctx.Invoke("azure-native:web/v20210115:listWebAppBackupConfigurationSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListWebAppBackupConfigurationSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppBackupConfigurationSlotResult struct {
	BackupName        *string                         `pulumi:"backupName"`
	BackupSchedule    *BackupScheduleResponse         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSettingResponse `pulumi:"databases"`
	Enabled           *bool                           `pulumi:"enabled"`
	Id                string                          `pulumi:"id"`
	Kind              *string                         `pulumi:"kind"`
	Name              string                          `pulumi:"name"`
	StorageAccountUrl string                          `pulumi:"storageAccountUrl"`
	Type              string                          `pulumi:"type"`
}


func (val *ListWebAppBackupConfigurationSlotResult) Defaults() *ListWebAppBackupConfigurationSlotResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BackupSchedule = tmp.BackupSchedule.Defaults()

	return &tmp
}

func ListWebAppBackupConfigurationSlotOutput(ctx *pulumi.Context, args ListWebAppBackupConfigurationSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppBackupConfigurationSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppBackupConfigurationSlotResult, error) {
			args := v.(ListWebAppBackupConfigurationSlotArgs)
			r, err := ListWebAppBackupConfigurationSlot(ctx, &args, opts...)
			var s ListWebAppBackupConfigurationSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppBackupConfigurationSlotResultOutput)
}

type ListWebAppBackupConfigurationSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppBackupConfigurationSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupConfigurationSlotArgs)(nil)).Elem()
}


type ListWebAppBackupConfigurationSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppBackupConfigurationSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupConfigurationSlotResult)(nil)).Elem()
}

func (o ListWebAppBackupConfigurationSlotResultOutput) ToListWebAppBackupConfigurationSlotResultOutput() ListWebAppBackupConfigurationSlotResultOutput {
	return o
}

func (o ListWebAppBackupConfigurationSlotResultOutput) ToListWebAppBackupConfigurationSlotResultOutputWithContext(ctx context.Context) ListWebAppBackupConfigurationSlotResultOutput {
	return o
}

func (o ListWebAppBackupConfigurationSlotResultOutput) BackupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) *string { return v.BackupName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppBackupConfigurationSlotResultOutput) BackupSchedule() BackupScheduleResponsePtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) *BackupScheduleResponse { return v.BackupSchedule }).(BackupScheduleResponsePtrOutput)
}

func (o ListWebAppBackupConfigurationSlotResultOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o ListWebAppBackupConfigurationSlotResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ListWebAppBackupConfigurationSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppBackupConfigurationSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppBackupConfigurationSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppBackupConfigurationSlotResultOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) string { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

func (o ListWebAppBackupConfigurationSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppBackupConfigurationSlotResultOutput{})
}

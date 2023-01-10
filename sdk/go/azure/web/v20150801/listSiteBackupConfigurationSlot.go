


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteBackupConfigurationSlot(ctx *pulumi.Context, args *ListSiteBackupConfigurationSlotArgs, opts ...pulumi.InvokeOption) (*ListSiteBackupConfigurationSlotResult, error) {
	var rv ListSiteBackupConfigurationSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteBackupConfigurationSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteBackupConfigurationSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListSiteBackupConfigurationSlotResult struct {
	BackupSchedule    *BackupScheduleResponse         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSettingResponse `pulumi:"databases"`
	Enabled           *bool                           `pulumi:"enabled"`
	Id                *string                         `pulumi:"id"`
	Kind              *string                         `pulumi:"kind"`
	Location          string                          `pulumi:"location"`
	Name              *string                         `pulumi:"name"`
	StorageAccountUrl *string                         `pulumi:"storageAccountUrl"`
	Tags              map[string]string               `pulumi:"tags"`
	Type              string                          `pulumi:"type"`
}

func ListSiteBackupConfigurationSlotOutput(ctx *pulumi.Context, args ListSiteBackupConfigurationSlotOutputArgs, opts ...pulumi.InvokeOption) ListSiteBackupConfigurationSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteBackupConfigurationSlotResult, error) {
			args := v.(ListSiteBackupConfigurationSlotArgs)
			r, err := ListSiteBackupConfigurationSlot(ctx, &args, opts...)
			var s ListSiteBackupConfigurationSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteBackupConfigurationSlotResultOutput)
}

type ListSiteBackupConfigurationSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListSiteBackupConfigurationSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteBackupConfigurationSlotArgs)(nil)).Elem()
}


type ListSiteBackupConfigurationSlotResultOutput struct{ *pulumi.OutputState }

func (ListSiteBackupConfigurationSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteBackupConfigurationSlotResult)(nil)).Elem()
}

func (o ListSiteBackupConfigurationSlotResultOutput) ToListSiteBackupConfigurationSlotResultOutput() ListSiteBackupConfigurationSlotResultOutput {
	return o
}

func (o ListSiteBackupConfigurationSlotResultOutput) ToListSiteBackupConfigurationSlotResultOutputWithContext(ctx context.Context) ListSiteBackupConfigurationSlotResultOutput {
	return o
}

func (o ListSiteBackupConfigurationSlotResultOutput) BackupSchedule() BackupScheduleResponsePtrOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationSlotResult) *BackupScheduleResponse { return v.BackupSchedule }).(BackupScheduleResponsePtrOutput)
}

func (o ListSiteBackupConfigurationSlotResultOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationSlotResult) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o ListSiteBackupConfigurationSlotResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationSlotResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ListSiteBackupConfigurationSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupConfigurationSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupConfigurationSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListSiteBackupConfigurationSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupConfigurationSlotResultOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationSlotResult) *string { return v.StorageAccountUrl }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupConfigurationSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListSiteBackupConfigurationSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteBackupConfigurationSlotResultOutput{})
}

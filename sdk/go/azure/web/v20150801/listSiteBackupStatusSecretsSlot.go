


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteBackupStatusSecretsSlot(ctx *pulumi.Context, args *ListSiteBackupStatusSecretsSlotArgs, opts ...pulumi.InvokeOption) (*ListSiteBackupStatusSecretsSlotResult, error) {
	var rv ListSiteBackupStatusSecretsSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteBackupStatusSecretsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteBackupStatusSecretsSlotArgs struct {
	BackupId          string                  `pulumi:"backupId"`
	BackupSchedule    *BackupSchedule         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSetting `pulumi:"databases"`
	Enabled           *bool                   `pulumi:"enabled"`
	Id                *string                 `pulumi:"id"`
	Kind              *string                 `pulumi:"kind"`
	Location          *string                 `pulumi:"location"`
	Name              string                  `pulumi:"name"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Slot              string                  `pulumi:"slot"`
	StorageAccountUrl *string                 `pulumi:"storageAccountUrl"`
	Tags              map[string]string       `pulumi:"tags"`
	Type              string                  `pulumi:"type"`
}


type ListSiteBackupStatusSecretsSlotResult struct {
	BlobName             *string                         `pulumi:"blobName"`
	CorrelationId        *string                         `pulumi:"correlationId"`
	Created              *string                         `pulumi:"created"`
	Databases            []DatabaseBackupSettingResponse `pulumi:"databases"`
	FinishedTimeStamp    *string                         `pulumi:"finishedTimeStamp"`
	Id                   *string                         `pulumi:"id"`
	Kind                 *string                         `pulumi:"kind"`
	LastRestoreTimeStamp *string                         `pulumi:"lastRestoreTimeStamp"`
	Location             string                          `pulumi:"location"`
	Log                  *string                         `pulumi:"log"`
	Name                 *string                         `pulumi:"name"`
	Scheduled            *bool                           `pulumi:"scheduled"`
	SizeInBytes          *float64                        `pulumi:"sizeInBytes"`
	Status               string                          `pulumi:"status"`
	StorageAccountUrl    *string                         `pulumi:"storageAccountUrl"`
	Tags                 map[string]string               `pulumi:"tags"`
	Type                 *string                         `pulumi:"type"`
	WebsiteSizeInBytes   *float64                        `pulumi:"websiteSizeInBytes"`
}

func ListSiteBackupStatusSecretsSlotOutput(ctx *pulumi.Context, args ListSiteBackupStatusSecretsSlotOutputArgs, opts ...pulumi.InvokeOption) ListSiteBackupStatusSecretsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteBackupStatusSecretsSlotResult, error) {
			args := v.(ListSiteBackupStatusSecretsSlotArgs)
			r, err := ListSiteBackupStatusSecretsSlot(ctx, &args, opts...)
			var s ListSiteBackupStatusSecretsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteBackupStatusSecretsSlotResultOutput)
}

type ListSiteBackupStatusSecretsSlotOutputArgs struct {
	BackupId          pulumi.StringInput              `pulumi:"backupId"`
	BackupSchedule    BackupSchedulePtrInput          `pulumi:"backupSchedule"`
	Databases         DatabaseBackupSettingArrayInput `pulumi:"databases"`
	Enabled           pulumi.BoolPtrInput             `pulumi:"enabled"`
	Id                pulumi.StringPtrInput           `pulumi:"id"`
	Kind              pulumi.StringPtrInput           `pulumi:"kind"`
	Location          pulumi.StringPtrInput           `pulumi:"location"`
	Name              pulumi.StringInput              `pulumi:"name"`
	ResourceGroupName pulumi.StringInput              `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput              `pulumi:"slot"`
	StorageAccountUrl pulumi.StringPtrInput           `pulumi:"storageAccountUrl"`
	Tags              pulumi.StringMapInput           `pulumi:"tags"`
	Type              pulumi.StringInput              `pulumi:"type"`
}

func (ListSiteBackupStatusSecretsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteBackupStatusSecretsSlotArgs)(nil)).Elem()
}


type ListSiteBackupStatusSecretsSlotResultOutput struct{ *pulumi.OutputState }

func (ListSiteBackupStatusSecretsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteBackupStatusSecretsSlotResult)(nil)).Elem()
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) ToListSiteBackupStatusSecretsSlotResultOutput() ListSiteBackupStatusSecretsSlotResultOutput {
	return o
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) ToListSiteBackupStatusSecretsSlotResultOutputWithContext(ctx context.Context) ListSiteBackupStatusSecretsSlotResultOutput {
	return o
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) BlobName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *string { return v.BlobName }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *string { return v.Created }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) FinishedTimeStamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *string { return v.FinishedTimeStamp }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) LastRestoreTimeStamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *string { return v.LastRestoreTimeStamp }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) Log() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *string { return v.Log }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) Scheduled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *bool { return v.Scheduled }).(pulumi.BoolPtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *float64 { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *string { return v.StorageAccountUrl }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsSlotResultOutput) WebsiteSizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsSlotResult) *float64 { return v.WebsiteSizeInBytes }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteBackupStatusSecretsSlotResultOutput{})
}

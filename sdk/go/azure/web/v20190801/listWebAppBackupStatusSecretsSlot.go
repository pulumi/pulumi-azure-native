


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppBackupStatusSecretsSlot(ctx *pulumi.Context, args *ListWebAppBackupStatusSecretsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppBackupStatusSecretsSlotResult, error) {
	var rv ListWebAppBackupStatusSecretsSlotResult
	err := ctx.Invoke("azure-native:web/v20190801:listWebAppBackupStatusSecretsSlot", args.Defaults(), &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppBackupStatusSecretsSlotArgs struct {
	BackupId          string                  `pulumi:"backupId"`
	BackupName        *string                 `pulumi:"backupName"`
	BackupSchedule    *BackupSchedule         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSetting `pulumi:"databases"`
	Enabled           *bool                   `pulumi:"enabled"`
	Kind              *string                 `pulumi:"kind"`
	Name              string                  `pulumi:"name"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Slot              string                  `pulumi:"slot"`
	StorageAccountUrl string                  `pulumi:"storageAccountUrl"`
}


func (val *ListWebAppBackupStatusSecretsSlotArgs) Defaults() *ListWebAppBackupStatusSecretsSlotArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BackupSchedule = tmp.BackupSchedule.Defaults()

	return &tmp
}


type ListWebAppBackupStatusSecretsSlotResult struct {
	BackupId             int                             `pulumi:"backupId"`
	BlobName             string                          `pulumi:"blobName"`
	CorrelationId        string                          `pulumi:"correlationId"`
	Created              string                          `pulumi:"created"`
	Databases            []DatabaseBackupSettingResponse `pulumi:"databases"`
	FinishedTimeStamp    string                          `pulumi:"finishedTimeStamp"`
	Id                   string                          `pulumi:"id"`
	Kind                 *string                         `pulumi:"kind"`
	LastRestoreTimeStamp string                          `pulumi:"lastRestoreTimeStamp"`
	Log                  string                          `pulumi:"log"`
	Name                 string                          `pulumi:"name"`
	Scheduled            bool                            `pulumi:"scheduled"`
	SizeInBytes          float64                         `pulumi:"sizeInBytes"`
	Status               string                          `pulumi:"status"`
	StorageAccountUrl    string                          `pulumi:"storageAccountUrl"`
	Type                 string                          `pulumi:"type"`
	WebsiteSizeInBytes   float64                         `pulumi:"websiteSizeInBytes"`
}

func ListWebAppBackupStatusSecretsSlotOutput(ctx *pulumi.Context, args ListWebAppBackupStatusSecretsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppBackupStatusSecretsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppBackupStatusSecretsSlotResult, error) {
			args := v.(ListWebAppBackupStatusSecretsSlotArgs)
			r, err := ListWebAppBackupStatusSecretsSlot(ctx, &args, opts...)
			var s ListWebAppBackupStatusSecretsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppBackupStatusSecretsSlotResultOutput)
}

type ListWebAppBackupStatusSecretsSlotOutputArgs struct {
	BackupId          pulumi.StringInput              `pulumi:"backupId"`
	BackupName        pulumi.StringPtrInput           `pulumi:"backupName"`
	BackupSchedule    BackupSchedulePtrInput          `pulumi:"backupSchedule"`
	Databases         DatabaseBackupSettingArrayInput `pulumi:"databases"`
	Enabled           pulumi.BoolPtrInput             `pulumi:"enabled"`
	Kind              pulumi.StringPtrInput           `pulumi:"kind"`
	Name              pulumi.StringInput              `pulumi:"name"`
	ResourceGroupName pulumi.StringInput              `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput              `pulumi:"slot"`
	StorageAccountUrl pulumi.StringInput              `pulumi:"storageAccountUrl"`
}

func (ListWebAppBackupStatusSecretsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupStatusSecretsSlotArgs)(nil)).Elem()
}


type ListWebAppBackupStatusSecretsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppBackupStatusSecretsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupStatusSecretsSlotResult)(nil)).Elem()
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) ToListWebAppBackupStatusSecretsSlotResultOutput() ListWebAppBackupStatusSecretsSlotResultOutput {
	return o
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) ToListWebAppBackupStatusSecretsSlotResultOutputWithContext(ctx context.Context) ListWebAppBackupStatusSecretsSlotResultOutput {
	return o
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) BackupId() pulumi.IntOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) int { return v.BackupId }).(pulumi.IntOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) BlobName() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) string { return v.BlobName }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) string { return v.CorrelationId }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) FinishedTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) string { return v.FinishedTimeStamp }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) LastRestoreTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) string { return v.LastRestoreTimeStamp }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) Log() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) string { return v.Log }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) Scheduled() pulumi.BoolOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) bool { return v.Scheduled }).(pulumi.BoolOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) float64 { return v.SizeInBytes }).(pulumi.Float64Output)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) string { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsSlotResultOutput) WebsiteSizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsSlotResult) float64 { return v.WebsiteSizeInBytes }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppBackupStatusSecretsSlotResultOutput{})
}

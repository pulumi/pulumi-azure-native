


package v20160801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppBackupStatusSecrets(ctx *pulumi.Context, args *ListWebAppBackupStatusSecretsArgs, opts ...pulumi.InvokeOption) (*ListWebAppBackupStatusSecretsResult, error) {
	var rv ListWebAppBackupStatusSecretsResult
	err := ctx.Invoke("azure-native:web/v20160801:listWebAppBackupStatusSecrets", args.Defaults(), &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppBackupStatusSecretsArgs struct {
	BackupId          string                      `pulumi:"backupId"`
	BackupRequestName string                      `pulumi:"backupRequestName"`
	BackupSchedule    *BackupSchedule             `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSetting     `pulumi:"databases"`
	Enabled           *bool                       `pulumi:"enabled"`
	Kind              *string                     `pulumi:"kind"`
	Name              string                      `pulumi:"name"`
	ResourceGroupName string                      `pulumi:"resourceGroupName"`
	StorageAccountUrl string                      `pulumi:"storageAccountUrl"`
	Type              *BackupRestoreOperationType `pulumi:"type"`
}


func (val *ListWebAppBackupStatusSecretsArgs) Defaults() *ListWebAppBackupStatusSecretsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BackupSchedule = tmp.BackupSchedule.Defaults()

	return &tmp
}


type ListWebAppBackupStatusSecretsResult struct {
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

func ListWebAppBackupStatusSecretsOutput(ctx *pulumi.Context, args ListWebAppBackupStatusSecretsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppBackupStatusSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppBackupStatusSecretsResult, error) {
			args := v.(ListWebAppBackupStatusSecretsArgs)
			r, err := ListWebAppBackupStatusSecrets(ctx, &args, opts...)
			var s ListWebAppBackupStatusSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppBackupStatusSecretsResultOutput)
}

type ListWebAppBackupStatusSecretsOutputArgs struct {
	BackupId          pulumi.StringInput                 `pulumi:"backupId"`
	BackupRequestName pulumi.StringInput                 `pulumi:"backupRequestName"`
	BackupSchedule    BackupSchedulePtrInput             `pulumi:"backupSchedule"`
	Databases         DatabaseBackupSettingArrayInput    `pulumi:"databases"`
	Enabled           pulumi.BoolPtrInput                `pulumi:"enabled"`
	Kind              pulumi.StringPtrInput              `pulumi:"kind"`
	Name              pulumi.StringInput                 `pulumi:"name"`
	ResourceGroupName pulumi.StringInput                 `pulumi:"resourceGroupName"`
	StorageAccountUrl pulumi.StringInput                 `pulumi:"storageAccountUrl"`
	Type              BackupRestoreOperationTypePtrInput `pulumi:"type"`
}

func (ListWebAppBackupStatusSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupStatusSecretsArgs)(nil)).Elem()
}


type ListWebAppBackupStatusSecretsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppBackupStatusSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupStatusSecretsResult)(nil)).Elem()
}

func (o ListWebAppBackupStatusSecretsResultOutput) ToListWebAppBackupStatusSecretsResultOutput() ListWebAppBackupStatusSecretsResultOutput {
	return o
}

func (o ListWebAppBackupStatusSecretsResultOutput) ToListWebAppBackupStatusSecretsResultOutputWithContext(ctx context.Context) ListWebAppBackupStatusSecretsResultOutput {
	return o
}

func (o ListWebAppBackupStatusSecretsResultOutput) BackupId() pulumi.IntOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) int { return v.BackupId }).(pulumi.IntOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) BlobName() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.BlobName }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.CorrelationId }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) FinishedTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.FinishedTimeStamp }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) LastRestoreTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.LastRestoreTimeStamp }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) Log() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.Log }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) Scheduled() pulumi.BoolOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) bool { return v.Scheduled }).(pulumi.BoolOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) float64 { return v.SizeInBytes }).(pulumi.Float64Output)
}

func (o ListWebAppBackupStatusSecretsResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o ListWebAppBackupStatusSecretsResultOutput) WebsiteSizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) float64 { return v.WebsiteSizeInBytes }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppBackupStatusSecretsResultOutput{})
}

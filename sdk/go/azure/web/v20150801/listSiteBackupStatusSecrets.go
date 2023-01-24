


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteBackupStatusSecrets(ctx *pulumi.Context, args *ListSiteBackupStatusSecretsArgs, opts ...pulumi.InvokeOption) (*ListSiteBackupStatusSecretsResult, error) {
	var rv ListSiteBackupStatusSecretsResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteBackupStatusSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteBackupStatusSecretsArgs struct {
	BackupId          string                  `pulumi:"backupId"`
	BackupSchedule    *BackupSchedule         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSetting `pulumi:"databases"`
	Enabled           *bool                   `pulumi:"enabled"`
	Id                *string                 `pulumi:"id"`
	Kind              *string                 `pulumi:"kind"`
	Location          *string                 `pulumi:"location"`
	Name              string                  `pulumi:"name"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	StorageAccountUrl *string                 `pulumi:"storageAccountUrl"`
	Tags              map[string]string       `pulumi:"tags"`
	Type              string                  `pulumi:"type"`
}


type ListSiteBackupStatusSecretsResult struct {
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

func ListSiteBackupStatusSecretsOutput(ctx *pulumi.Context, args ListSiteBackupStatusSecretsOutputArgs, opts ...pulumi.InvokeOption) ListSiteBackupStatusSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteBackupStatusSecretsResult, error) {
			args := v.(ListSiteBackupStatusSecretsArgs)
			r, err := ListSiteBackupStatusSecrets(ctx, &args, opts...)
			var s ListSiteBackupStatusSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteBackupStatusSecretsResultOutput)
}

type ListSiteBackupStatusSecretsOutputArgs struct {
	BackupId          pulumi.StringInput              `pulumi:"backupId"`
	BackupSchedule    BackupSchedulePtrInput          `pulumi:"backupSchedule"`
	Databases         DatabaseBackupSettingArrayInput `pulumi:"databases"`
	Enabled           pulumi.BoolPtrInput             `pulumi:"enabled"`
	Id                pulumi.StringPtrInput           `pulumi:"id"`
	Kind              pulumi.StringPtrInput           `pulumi:"kind"`
	Location          pulumi.StringPtrInput           `pulumi:"location"`
	Name              pulumi.StringInput              `pulumi:"name"`
	ResourceGroupName pulumi.StringInput              `pulumi:"resourceGroupName"`
	StorageAccountUrl pulumi.StringPtrInput           `pulumi:"storageAccountUrl"`
	Tags              pulumi.StringMapInput           `pulumi:"tags"`
	Type              pulumi.StringInput              `pulumi:"type"`
}

func (ListSiteBackupStatusSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteBackupStatusSecretsArgs)(nil)).Elem()
}


type ListSiteBackupStatusSecretsResultOutput struct{ *pulumi.OutputState }

func (ListSiteBackupStatusSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteBackupStatusSecretsResult)(nil)).Elem()
}

func (o ListSiteBackupStatusSecretsResultOutput) ToListSiteBackupStatusSecretsResultOutput() ListSiteBackupStatusSecretsResultOutput {
	return o
}

func (o ListSiteBackupStatusSecretsResultOutput) ToListSiteBackupStatusSecretsResultOutputWithContext(ctx context.Context) ListSiteBackupStatusSecretsResultOutput {
	return o
}

func (o ListSiteBackupStatusSecretsResultOutput) BlobName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *string { return v.BlobName }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *string { return v.Created }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) FinishedTimeStamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *string { return v.FinishedTimeStamp }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) LastRestoreTimeStamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *string { return v.LastRestoreTimeStamp }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) Log() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *string { return v.Log }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) Scheduled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *bool { return v.Scheduled }).(pulumi.BoolPtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *float64 { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *string { return v.StorageAccountUrl }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupStatusSecretsResultOutput) WebsiteSizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListSiteBackupStatusSecretsResult) *float64 { return v.WebsiteSizeInBytes }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteBackupStatusSecretsResultOutput{})
}

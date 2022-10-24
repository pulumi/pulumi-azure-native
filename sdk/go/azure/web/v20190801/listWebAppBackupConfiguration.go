


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppBackupConfiguration(ctx *pulumi.Context, args *ListWebAppBackupConfigurationArgs, opts ...pulumi.InvokeOption) (*ListWebAppBackupConfigurationResult, error) {
	var rv ListWebAppBackupConfigurationResult
	err := ctx.Invoke("azure-native:web/v20190801:listWebAppBackupConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListWebAppBackupConfigurationArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppBackupConfigurationResult struct {
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


func (val *ListWebAppBackupConfigurationResult) Defaults() *ListWebAppBackupConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BackupSchedule = tmp.BackupSchedule.Defaults()

	return &tmp
}

func ListWebAppBackupConfigurationOutput(ctx *pulumi.Context, args ListWebAppBackupConfigurationOutputArgs, opts ...pulumi.InvokeOption) ListWebAppBackupConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppBackupConfigurationResult, error) {
			args := v.(ListWebAppBackupConfigurationArgs)
			r, err := ListWebAppBackupConfiguration(ctx, &args, opts...)
			var s ListWebAppBackupConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppBackupConfigurationResultOutput)
}

type ListWebAppBackupConfigurationOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppBackupConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupConfigurationArgs)(nil)).Elem()
}


type ListWebAppBackupConfigurationResultOutput struct{ *pulumi.OutputState }

func (ListWebAppBackupConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupConfigurationResult)(nil)).Elem()
}

func (o ListWebAppBackupConfigurationResultOutput) ToListWebAppBackupConfigurationResultOutput() ListWebAppBackupConfigurationResultOutput {
	return o
}

func (o ListWebAppBackupConfigurationResultOutput) ToListWebAppBackupConfigurationResultOutputWithContext(ctx context.Context) ListWebAppBackupConfigurationResultOutput {
	return o
}

func (o ListWebAppBackupConfigurationResultOutput) BackupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) *string { return v.BackupName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppBackupConfigurationResultOutput) BackupSchedule() BackupScheduleResponsePtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) *BackupScheduleResponse { return v.BackupSchedule }).(BackupScheduleResponsePtrOutput)
}

func (o ListWebAppBackupConfigurationResultOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o ListWebAppBackupConfigurationResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ListWebAppBackupConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppBackupConfigurationResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppBackupConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppBackupConfigurationResultOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) string { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

func (o ListWebAppBackupConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppBackupConfigurationResultOutput{})
}

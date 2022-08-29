


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteBackupConfiguration(ctx *pulumi.Context, args *ListSiteBackupConfigurationArgs, opts ...pulumi.InvokeOption) (*ListSiteBackupConfigurationResult, error) {
	var rv ListSiteBackupConfigurationResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteBackupConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteBackupConfigurationArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSiteBackupConfigurationResult struct {
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

func ListSiteBackupConfigurationOutput(ctx *pulumi.Context, args ListSiteBackupConfigurationOutputArgs, opts ...pulumi.InvokeOption) ListSiteBackupConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteBackupConfigurationResult, error) {
			args := v.(ListSiteBackupConfigurationArgs)
			r, err := ListSiteBackupConfiguration(ctx, &args, opts...)
			var s ListSiteBackupConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteBackupConfigurationResultOutput)
}

type ListSiteBackupConfigurationOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListSiteBackupConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteBackupConfigurationArgs)(nil)).Elem()
}


type ListSiteBackupConfigurationResultOutput struct{ *pulumi.OutputState }

func (ListSiteBackupConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteBackupConfigurationResult)(nil)).Elem()
}

func (o ListSiteBackupConfigurationResultOutput) ToListSiteBackupConfigurationResultOutput() ListSiteBackupConfigurationResultOutput {
	return o
}

func (o ListSiteBackupConfigurationResultOutput) ToListSiteBackupConfigurationResultOutputWithContext(ctx context.Context) ListSiteBackupConfigurationResultOutput {
	return o
}

func (o ListSiteBackupConfigurationResultOutput) BackupSchedule() BackupScheduleResponsePtrOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationResult) *BackupScheduleResponse { return v.BackupSchedule }).(BackupScheduleResponsePtrOutput)
}

func (o ListSiteBackupConfigurationResultOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationResult) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o ListSiteBackupConfigurationResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ListSiteBackupConfigurationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupConfigurationResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupConfigurationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListSiteBackupConfigurationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupConfigurationResultOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationResult) *string { return v.StorageAccountUrl }).(pulumi.StringPtrOutput)
}

func (o ListSiteBackupConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListSiteBackupConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteBackupConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteBackupConfigurationResultOutput{})
}

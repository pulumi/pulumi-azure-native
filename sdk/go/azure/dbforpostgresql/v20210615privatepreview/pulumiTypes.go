


package v20210615privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AADApp struct {
	AadSecret string `pulumi:"aadSecret"`
	ClientId  string `pulumi:"clientId"`
	TenantId  string `pulumi:"tenantId"`
}





type AADAppInput interface {
	pulumi.Input

	ToAADAppOutput() AADAppOutput
	ToAADAppOutputWithContext(context.Context) AADAppOutput
}

type AADAppArgs struct {
	AadSecret pulumi.StringInput `pulumi:"aadSecret"`
	ClientId  pulumi.StringInput `pulumi:"clientId"`
	TenantId  pulumi.StringInput `pulumi:"tenantId"`
}

func (AADAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AADApp)(nil)).Elem()
}

func (i AADAppArgs) ToAADAppOutput() AADAppOutput {
	return i.ToAADAppOutputWithContext(context.Background())
}

func (i AADAppArgs) ToAADAppOutputWithContext(ctx context.Context) AADAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADAppOutput)
}

func (i AADAppArgs) ToAADAppPtrOutput() AADAppPtrOutput {
	return i.ToAADAppPtrOutputWithContext(context.Background())
}

func (i AADAppArgs) ToAADAppPtrOutputWithContext(ctx context.Context) AADAppPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADAppOutput).ToAADAppPtrOutputWithContext(ctx)
}









type AADAppPtrInput interface {
	pulumi.Input

	ToAADAppPtrOutput() AADAppPtrOutput
	ToAADAppPtrOutputWithContext(context.Context) AADAppPtrOutput
}

type aadappPtrType AADAppArgs

func AADAppPtr(v *AADAppArgs) AADAppPtrInput {
	return (*aadappPtrType)(v)
}

func (*aadappPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AADApp)(nil)).Elem()
}

func (i *aadappPtrType) ToAADAppPtrOutput() AADAppPtrOutput {
	return i.ToAADAppPtrOutputWithContext(context.Background())
}

func (i *aadappPtrType) ToAADAppPtrOutputWithContext(ctx context.Context) AADAppPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADAppPtrOutput)
}

type AADAppOutput struct{ *pulumi.OutputState }

func (AADAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AADApp)(nil)).Elem()
}

func (o AADAppOutput) ToAADAppOutput() AADAppOutput {
	return o
}

func (o AADAppOutput) ToAADAppOutputWithContext(ctx context.Context) AADAppOutput {
	return o
}

func (o AADAppOutput) ToAADAppPtrOutput() AADAppPtrOutput {
	return o.ToAADAppPtrOutputWithContext(context.Background())
}

func (o AADAppOutput) ToAADAppPtrOutputWithContext(ctx context.Context) AADAppPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AADApp) *AADApp {
		return &v
	}).(AADAppPtrOutput)
}

func (o AADAppOutput) AadSecret() pulumi.StringOutput {
	return o.ApplyT(func(v AADApp) string { return v.AadSecret }).(pulumi.StringOutput)
}

func (o AADAppOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v AADApp) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o AADAppOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v AADApp) string { return v.TenantId }).(pulumi.StringOutput)
}

type AADAppPtrOutput struct{ *pulumi.OutputState }

func (AADAppPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AADApp)(nil)).Elem()
}

func (o AADAppPtrOutput) ToAADAppPtrOutput() AADAppPtrOutput {
	return o
}

func (o AADAppPtrOutput) ToAADAppPtrOutputWithContext(ctx context.Context) AADAppPtrOutput {
	return o
}

func (o AADAppPtrOutput) Elem() AADAppOutput {
	return o.ApplyT(func(v *AADApp) AADApp {
		if v != nil {
			return *v
		}
		var ret AADApp
		return ret
	}).(AADAppOutput)
}

func (o AADAppPtrOutput) AadSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADApp) *string {
		if v == nil {
			return nil
		}
		return &v.AadSecret
	}).(pulumi.StringPtrOutput)
}

func (o AADAppPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADApp) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AADAppPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADApp) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

type AADAppResponse struct {
	ClientId string `pulumi:"clientId"`
	TenantId string `pulumi:"tenantId"`
}





type AADAppResponseInput interface {
	pulumi.Input

	ToAADAppResponseOutput() AADAppResponseOutput
	ToAADAppResponseOutputWithContext(context.Context) AADAppResponseOutput
}

type AADAppResponseArgs struct {
	ClientId pulumi.StringInput `pulumi:"clientId"`
	TenantId pulumi.StringInput `pulumi:"tenantId"`
}

func (AADAppResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AADAppResponse)(nil)).Elem()
}

func (i AADAppResponseArgs) ToAADAppResponseOutput() AADAppResponseOutput {
	return i.ToAADAppResponseOutputWithContext(context.Background())
}

func (i AADAppResponseArgs) ToAADAppResponseOutputWithContext(ctx context.Context) AADAppResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADAppResponseOutput)
}

func (i AADAppResponseArgs) ToAADAppResponsePtrOutput() AADAppResponsePtrOutput {
	return i.ToAADAppResponsePtrOutputWithContext(context.Background())
}

func (i AADAppResponseArgs) ToAADAppResponsePtrOutputWithContext(ctx context.Context) AADAppResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADAppResponseOutput).ToAADAppResponsePtrOutputWithContext(ctx)
}









type AADAppResponsePtrInput interface {
	pulumi.Input

	ToAADAppResponsePtrOutput() AADAppResponsePtrOutput
	ToAADAppResponsePtrOutputWithContext(context.Context) AADAppResponsePtrOutput
}

type aadappResponsePtrType AADAppResponseArgs

func AADAppResponsePtr(v *AADAppResponseArgs) AADAppResponsePtrInput {
	return (*aadappResponsePtrType)(v)
}

func (*aadappResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AADAppResponse)(nil)).Elem()
}

func (i *aadappResponsePtrType) ToAADAppResponsePtrOutput() AADAppResponsePtrOutput {
	return i.ToAADAppResponsePtrOutputWithContext(context.Background())
}

func (i *aadappResponsePtrType) ToAADAppResponsePtrOutputWithContext(ctx context.Context) AADAppResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADAppResponsePtrOutput)
}

type AADAppResponseOutput struct{ *pulumi.OutputState }

func (AADAppResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AADAppResponse)(nil)).Elem()
}

func (o AADAppResponseOutput) ToAADAppResponseOutput() AADAppResponseOutput {
	return o
}

func (o AADAppResponseOutput) ToAADAppResponseOutputWithContext(ctx context.Context) AADAppResponseOutput {
	return o
}

func (o AADAppResponseOutput) ToAADAppResponsePtrOutput() AADAppResponsePtrOutput {
	return o.ToAADAppResponsePtrOutputWithContext(context.Background())
}

func (o AADAppResponseOutput) ToAADAppResponsePtrOutputWithContext(ctx context.Context) AADAppResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AADAppResponse) *AADAppResponse {
		return &v
	}).(AADAppResponsePtrOutput)
}

func (o AADAppResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v AADAppResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o AADAppResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v AADAppResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type AADAppResponsePtrOutput struct{ *pulumi.OutputState }

func (AADAppResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AADAppResponse)(nil)).Elem()
}

func (o AADAppResponsePtrOutput) ToAADAppResponsePtrOutput() AADAppResponsePtrOutput {
	return o
}

func (o AADAppResponsePtrOutput) ToAADAppResponsePtrOutputWithContext(ctx context.Context) AADAppResponsePtrOutput {
	return o
}

func (o AADAppResponsePtrOutput) Elem() AADAppResponseOutput {
	return o.ApplyT(func(v *AADAppResponse) AADAppResponse {
		if v != nil {
			return *v
		}
		var ret AADAppResponse
		return ret
	}).(AADAppResponseOutput)
}

func (o AADAppResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADAppResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AADAppResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADAppResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

type AdminCredentials struct {
	SourceServerPassword string `pulumi:"sourceServerPassword"`
	TargetServerPassword string `pulumi:"targetServerPassword"`
}





type AdminCredentialsInput interface {
	pulumi.Input

	ToAdminCredentialsOutput() AdminCredentialsOutput
	ToAdminCredentialsOutputWithContext(context.Context) AdminCredentialsOutput
}

type AdminCredentialsArgs struct {
	SourceServerPassword pulumi.StringInput `pulumi:"sourceServerPassword"`
	TargetServerPassword pulumi.StringInput `pulumi:"targetServerPassword"`
}

func (AdminCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdminCredentials)(nil)).Elem()
}

func (i AdminCredentialsArgs) ToAdminCredentialsOutput() AdminCredentialsOutput {
	return i.ToAdminCredentialsOutputWithContext(context.Background())
}

func (i AdminCredentialsArgs) ToAdminCredentialsOutputWithContext(ctx context.Context) AdminCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdminCredentialsOutput)
}

func (i AdminCredentialsArgs) ToAdminCredentialsPtrOutput() AdminCredentialsPtrOutput {
	return i.ToAdminCredentialsPtrOutputWithContext(context.Background())
}

func (i AdminCredentialsArgs) ToAdminCredentialsPtrOutputWithContext(ctx context.Context) AdminCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdminCredentialsOutput).ToAdminCredentialsPtrOutputWithContext(ctx)
}









type AdminCredentialsPtrInput interface {
	pulumi.Input

	ToAdminCredentialsPtrOutput() AdminCredentialsPtrOutput
	ToAdminCredentialsPtrOutputWithContext(context.Context) AdminCredentialsPtrOutput
}

type adminCredentialsPtrType AdminCredentialsArgs

func AdminCredentialsPtr(v *AdminCredentialsArgs) AdminCredentialsPtrInput {
	return (*adminCredentialsPtrType)(v)
}

func (*adminCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AdminCredentials)(nil)).Elem()
}

func (i *adminCredentialsPtrType) ToAdminCredentialsPtrOutput() AdminCredentialsPtrOutput {
	return i.ToAdminCredentialsPtrOutputWithContext(context.Background())
}

func (i *adminCredentialsPtrType) ToAdminCredentialsPtrOutputWithContext(ctx context.Context) AdminCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdminCredentialsPtrOutput)
}

type AdminCredentialsOutput struct{ *pulumi.OutputState }

func (AdminCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdminCredentials)(nil)).Elem()
}

func (o AdminCredentialsOutput) ToAdminCredentialsOutput() AdminCredentialsOutput {
	return o
}

func (o AdminCredentialsOutput) ToAdminCredentialsOutputWithContext(ctx context.Context) AdminCredentialsOutput {
	return o
}

func (o AdminCredentialsOutput) ToAdminCredentialsPtrOutput() AdminCredentialsPtrOutput {
	return o.ToAdminCredentialsPtrOutputWithContext(context.Background())
}

func (o AdminCredentialsOutput) ToAdminCredentialsPtrOutputWithContext(ctx context.Context) AdminCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdminCredentials) *AdminCredentials {
		return &v
	}).(AdminCredentialsPtrOutput)
}

func (o AdminCredentialsOutput) SourceServerPassword() pulumi.StringOutput {
	return o.ApplyT(func(v AdminCredentials) string { return v.SourceServerPassword }).(pulumi.StringOutput)
}

func (o AdminCredentialsOutput) TargetServerPassword() pulumi.StringOutput {
	return o.ApplyT(func(v AdminCredentials) string { return v.TargetServerPassword }).(pulumi.StringOutput)
}

type AdminCredentialsPtrOutput struct{ *pulumi.OutputState }

func (AdminCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdminCredentials)(nil)).Elem()
}

func (o AdminCredentialsPtrOutput) ToAdminCredentialsPtrOutput() AdminCredentialsPtrOutput {
	return o
}

func (o AdminCredentialsPtrOutput) ToAdminCredentialsPtrOutputWithContext(ctx context.Context) AdminCredentialsPtrOutput {
	return o
}

func (o AdminCredentialsPtrOutput) Elem() AdminCredentialsOutput {
	return o.ApplyT(func(v *AdminCredentials) AdminCredentials {
		if v != nil {
			return *v
		}
		var ret AdminCredentials
		return ret
	}).(AdminCredentialsOutput)
}

func (o AdminCredentialsPtrOutput) SourceServerPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdminCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.SourceServerPassword
	}).(pulumi.StringPtrOutput)
}

func (o AdminCredentialsPtrOutput) TargetServerPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdminCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.TargetServerPassword
	}).(pulumi.StringPtrOutput)
}

type Backup struct {
	BackupRetentionDays *int    `pulumi:"backupRetentionDays"`
	GeoRedundantBackup  *string `pulumi:"geoRedundantBackup"`
}





type BackupInput interface {
	pulumi.Input

	ToBackupOutput() BackupOutput
	ToBackupOutputWithContext(context.Context) BackupOutput
}

type BackupArgs struct {
	BackupRetentionDays pulumi.IntPtrInput    `pulumi:"backupRetentionDays"`
	GeoRedundantBackup  pulumi.StringPtrInput `pulumi:"geoRedundantBackup"`
}

func (BackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Backup)(nil)).Elem()
}

func (i BackupArgs) ToBackupOutput() BackupOutput {
	return i.ToBackupOutputWithContext(context.Background())
}

func (i BackupArgs) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput)
}

func (i BackupArgs) ToBackupPtrOutput() BackupPtrOutput {
	return i.ToBackupPtrOutputWithContext(context.Background())
}

func (i BackupArgs) ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput).ToBackupPtrOutputWithContext(ctx)
}









type BackupPtrInput interface {
	pulumi.Input

	ToBackupPtrOutput() BackupPtrOutput
	ToBackupPtrOutputWithContext(context.Context) BackupPtrOutput
}

type backupPtrType BackupArgs

func BackupPtr(v *BackupArgs) BackupPtrInput {
	return (*backupPtrType)(v)
}

func (*backupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (i *backupPtrType) ToBackupPtrOutput() BackupPtrOutput {
	return i.ToBackupPtrOutputWithContext(context.Background())
}

func (i *backupPtrType) ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPtrOutput)
}

type BackupOutput struct{ *pulumi.OutputState }

func (BackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Backup)(nil)).Elem()
}

func (o BackupOutput) ToBackupOutput() BackupOutput {
	return o
}

func (o BackupOutput) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return o
}

func (o BackupOutput) ToBackupPtrOutput() BackupPtrOutput {
	return o.ToBackupPtrOutputWithContext(context.Background())
}

func (o BackupOutput) ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Backup) *Backup {
		return &v
	}).(BackupPtrOutput)
}

func (o BackupOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Backup) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o BackupOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backup) *string { return v.GeoRedundantBackup }).(pulumi.StringPtrOutput)
}

type BackupPtrOutput struct{ *pulumi.OutputState }

func (BackupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (o BackupPtrOutput) ToBackupPtrOutput() BackupPtrOutput {
	return o
}

func (o BackupPtrOutput) ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput {
	return o
}

func (o BackupPtrOutput) Elem() BackupOutput {
	return o.ApplyT(func(v *Backup) Backup {
		if v != nil {
			return *v
		}
		var ret Backup
		return ret
	}).(BackupOutput)
}

func (o BackupPtrOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Backup) *int {
		if v == nil {
			return nil
		}
		return v.BackupRetentionDays
	}).(pulumi.IntPtrOutput)
}

func (o BackupPtrOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backup) *string {
		if v == nil {
			return nil
		}
		return v.GeoRedundantBackup
	}).(pulumi.StringPtrOutput)
}

type BackupResponse struct {
	BackupRetentionDays *int    `pulumi:"backupRetentionDays"`
	EarliestRestoreDate string  `pulumi:"earliestRestoreDate"`
	GeoRedundantBackup  *string `pulumi:"geoRedundantBackup"`
}





type BackupResponseInput interface {
	pulumi.Input

	ToBackupResponseOutput() BackupResponseOutput
	ToBackupResponseOutputWithContext(context.Context) BackupResponseOutput
}

type BackupResponseArgs struct {
	BackupRetentionDays pulumi.IntPtrInput    `pulumi:"backupRetentionDays"`
	EarliestRestoreDate pulumi.StringInput    `pulumi:"earliestRestoreDate"`
	GeoRedundantBackup  pulumi.StringPtrInput `pulumi:"geoRedundantBackup"`
}

func (BackupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupResponse)(nil)).Elem()
}

func (i BackupResponseArgs) ToBackupResponseOutput() BackupResponseOutput {
	return i.ToBackupResponseOutputWithContext(context.Background())
}

func (i BackupResponseArgs) ToBackupResponseOutputWithContext(ctx context.Context) BackupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupResponseOutput)
}

func (i BackupResponseArgs) ToBackupResponsePtrOutput() BackupResponsePtrOutput {
	return i.ToBackupResponsePtrOutputWithContext(context.Background())
}

func (i BackupResponseArgs) ToBackupResponsePtrOutputWithContext(ctx context.Context) BackupResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupResponseOutput).ToBackupResponsePtrOutputWithContext(ctx)
}









type BackupResponsePtrInput interface {
	pulumi.Input

	ToBackupResponsePtrOutput() BackupResponsePtrOutput
	ToBackupResponsePtrOutputWithContext(context.Context) BackupResponsePtrOutput
}

type backupResponsePtrType BackupResponseArgs

func BackupResponsePtr(v *BackupResponseArgs) BackupResponsePtrInput {
	return (*backupResponsePtrType)(v)
}

func (*backupResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupResponse)(nil)).Elem()
}

func (i *backupResponsePtrType) ToBackupResponsePtrOutput() BackupResponsePtrOutput {
	return i.ToBackupResponsePtrOutputWithContext(context.Background())
}

func (i *backupResponsePtrType) ToBackupResponsePtrOutputWithContext(ctx context.Context) BackupResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupResponsePtrOutput)
}

type BackupResponseOutput struct{ *pulumi.OutputState }

func (BackupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupResponse)(nil)).Elem()
}

func (o BackupResponseOutput) ToBackupResponseOutput() BackupResponseOutput {
	return o
}

func (o BackupResponseOutput) ToBackupResponseOutputWithContext(ctx context.Context) BackupResponseOutput {
	return o
}

func (o BackupResponseOutput) ToBackupResponsePtrOutput() BackupResponsePtrOutput {
	return o.ToBackupResponsePtrOutputWithContext(context.Background())
}

func (o BackupResponseOutput) ToBackupResponsePtrOutputWithContext(ctx context.Context) BackupResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupResponse) *BackupResponse {
		return &v
	}).(BackupResponsePtrOutput)
}

func (o BackupResponseOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackupResponse) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o BackupResponseOutput) EarliestRestoreDate() pulumi.StringOutput {
	return o.ApplyT(func(v BackupResponse) string { return v.EarliestRestoreDate }).(pulumi.StringOutput)
}

func (o BackupResponseOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupResponse) *string { return v.GeoRedundantBackup }).(pulumi.StringPtrOutput)
}

type BackupResponsePtrOutput struct{ *pulumi.OutputState }

func (BackupResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupResponse)(nil)).Elem()
}

func (o BackupResponsePtrOutput) ToBackupResponsePtrOutput() BackupResponsePtrOutput {
	return o
}

func (o BackupResponsePtrOutput) ToBackupResponsePtrOutputWithContext(ctx context.Context) BackupResponsePtrOutput {
	return o
}

func (o BackupResponsePtrOutput) Elem() BackupResponseOutput {
	return o.ApplyT(func(v *BackupResponse) BackupResponse {
		if v != nil {
			return *v
		}
		var ret BackupResponse
		return ret
	}).(BackupResponseOutput)
}

func (o BackupResponsePtrOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupResponse) *int {
		if v == nil {
			return nil
		}
		return v.BackupRetentionDays
	}).(pulumi.IntPtrOutput)
}

func (o BackupResponsePtrOutput) EarliestRestoreDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EarliestRestoreDate
	}).(pulumi.StringPtrOutput)
}

func (o BackupResponsePtrOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupResponse) *string {
		if v == nil {
			return nil
		}
		return v.GeoRedundantBackup
	}).(pulumi.StringPtrOutput)
}

type DBServerMetadataResponse struct {
	Location  *string            `pulumi:"location"`
	Sku       *ServerSkuResponse `pulumi:"sku"`
	StorageMB *int               `pulumi:"storageMB"`
	Version   *string            `pulumi:"version"`
}





type DBServerMetadataResponseInput interface {
	pulumi.Input

	ToDBServerMetadataResponseOutput() DBServerMetadataResponseOutput
	ToDBServerMetadataResponseOutputWithContext(context.Context) DBServerMetadataResponseOutput
}

type DBServerMetadataResponseArgs struct {
	Location  pulumi.StringPtrInput     `pulumi:"location"`
	Sku       ServerSkuResponsePtrInput `pulumi:"sku"`
	StorageMB pulumi.IntPtrInput        `pulumi:"storageMB"`
	Version   pulumi.StringPtrInput     `pulumi:"version"`
}

func (DBServerMetadataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DBServerMetadataResponse)(nil)).Elem()
}

func (i DBServerMetadataResponseArgs) ToDBServerMetadataResponseOutput() DBServerMetadataResponseOutput {
	return i.ToDBServerMetadataResponseOutputWithContext(context.Background())
}

func (i DBServerMetadataResponseArgs) ToDBServerMetadataResponseOutputWithContext(ctx context.Context) DBServerMetadataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBServerMetadataResponseOutput)
}

func (i DBServerMetadataResponseArgs) ToDBServerMetadataResponsePtrOutput() DBServerMetadataResponsePtrOutput {
	return i.ToDBServerMetadataResponsePtrOutputWithContext(context.Background())
}

func (i DBServerMetadataResponseArgs) ToDBServerMetadataResponsePtrOutputWithContext(ctx context.Context) DBServerMetadataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBServerMetadataResponseOutput).ToDBServerMetadataResponsePtrOutputWithContext(ctx)
}









type DBServerMetadataResponsePtrInput interface {
	pulumi.Input

	ToDBServerMetadataResponsePtrOutput() DBServerMetadataResponsePtrOutput
	ToDBServerMetadataResponsePtrOutputWithContext(context.Context) DBServerMetadataResponsePtrOutput
}

type dbserverMetadataResponsePtrType DBServerMetadataResponseArgs

func DBServerMetadataResponsePtr(v *DBServerMetadataResponseArgs) DBServerMetadataResponsePtrInput {
	return (*dbserverMetadataResponsePtrType)(v)
}

func (*dbserverMetadataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DBServerMetadataResponse)(nil)).Elem()
}

func (i *dbserverMetadataResponsePtrType) ToDBServerMetadataResponsePtrOutput() DBServerMetadataResponsePtrOutput {
	return i.ToDBServerMetadataResponsePtrOutputWithContext(context.Background())
}

func (i *dbserverMetadataResponsePtrType) ToDBServerMetadataResponsePtrOutputWithContext(ctx context.Context) DBServerMetadataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBServerMetadataResponsePtrOutput)
}

type DBServerMetadataResponseOutput struct{ *pulumi.OutputState }

func (DBServerMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBServerMetadataResponse)(nil)).Elem()
}

func (o DBServerMetadataResponseOutput) ToDBServerMetadataResponseOutput() DBServerMetadataResponseOutput {
	return o
}

func (o DBServerMetadataResponseOutput) ToDBServerMetadataResponseOutputWithContext(ctx context.Context) DBServerMetadataResponseOutput {
	return o
}

func (o DBServerMetadataResponseOutput) ToDBServerMetadataResponsePtrOutput() DBServerMetadataResponsePtrOutput {
	return o.ToDBServerMetadataResponsePtrOutputWithContext(context.Background())
}

func (o DBServerMetadataResponseOutput) ToDBServerMetadataResponsePtrOutputWithContext(ctx context.Context) DBServerMetadataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DBServerMetadataResponse) *DBServerMetadataResponse {
		return &v
	}).(DBServerMetadataResponsePtrOutput)
}

func (o DBServerMetadataResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DBServerMetadataResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DBServerMetadataResponseOutput) Sku() ServerSkuResponsePtrOutput {
	return o.ApplyT(func(v DBServerMetadataResponse) *ServerSkuResponse { return v.Sku }).(ServerSkuResponsePtrOutput)
}

func (o DBServerMetadataResponseOutput) StorageMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DBServerMetadataResponse) *int { return v.StorageMB }).(pulumi.IntPtrOutput)
}

func (o DBServerMetadataResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DBServerMetadataResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type DBServerMetadataResponsePtrOutput struct{ *pulumi.OutputState }

func (DBServerMetadataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DBServerMetadataResponse)(nil)).Elem()
}

func (o DBServerMetadataResponsePtrOutput) ToDBServerMetadataResponsePtrOutput() DBServerMetadataResponsePtrOutput {
	return o
}

func (o DBServerMetadataResponsePtrOutput) ToDBServerMetadataResponsePtrOutputWithContext(ctx context.Context) DBServerMetadataResponsePtrOutput {
	return o
}

func (o DBServerMetadataResponsePtrOutput) Elem() DBServerMetadataResponseOutput {
	return o.ApplyT(func(v *DBServerMetadataResponse) DBServerMetadataResponse {
		if v != nil {
			return *v
		}
		var ret DBServerMetadataResponse
		return ret
	}).(DBServerMetadataResponseOutput)
}

func (o DBServerMetadataResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBServerMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o DBServerMetadataResponsePtrOutput) Sku() ServerSkuResponsePtrOutput {
	return o.ApplyT(func(v *DBServerMetadataResponse) *ServerSkuResponse {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(ServerSkuResponsePtrOutput)
}

func (o DBServerMetadataResponsePtrOutput) StorageMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DBServerMetadataResponse) *int {
		if v == nil {
			return nil
		}
		return v.StorageMB
	}).(pulumi.IntPtrOutput)
}

func (o DBServerMetadataResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBServerMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type HighAvailability struct {
	Mode                    *string `pulumi:"mode"`
	StandbyAvailabilityZone *string `pulumi:"standbyAvailabilityZone"`
}





type HighAvailabilityInput interface {
	pulumi.Input

	ToHighAvailabilityOutput() HighAvailabilityOutput
	ToHighAvailabilityOutputWithContext(context.Context) HighAvailabilityOutput
}

type HighAvailabilityArgs struct {
	Mode                    pulumi.StringPtrInput `pulumi:"mode"`
	StandbyAvailabilityZone pulumi.StringPtrInput `pulumi:"standbyAvailabilityZone"`
}

func (HighAvailabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HighAvailability)(nil)).Elem()
}

func (i HighAvailabilityArgs) ToHighAvailabilityOutput() HighAvailabilityOutput {
	return i.ToHighAvailabilityOutputWithContext(context.Background())
}

func (i HighAvailabilityArgs) ToHighAvailabilityOutputWithContext(ctx context.Context) HighAvailabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HighAvailabilityOutput)
}

func (i HighAvailabilityArgs) ToHighAvailabilityPtrOutput() HighAvailabilityPtrOutput {
	return i.ToHighAvailabilityPtrOutputWithContext(context.Background())
}

func (i HighAvailabilityArgs) ToHighAvailabilityPtrOutputWithContext(ctx context.Context) HighAvailabilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HighAvailabilityOutput).ToHighAvailabilityPtrOutputWithContext(ctx)
}









type HighAvailabilityPtrInput interface {
	pulumi.Input

	ToHighAvailabilityPtrOutput() HighAvailabilityPtrOutput
	ToHighAvailabilityPtrOutputWithContext(context.Context) HighAvailabilityPtrOutput
}

type highAvailabilityPtrType HighAvailabilityArgs

func HighAvailabilityPtr(v *HighAvailabilityArgs) HighAvailabilityPtrInput {
	return (*highAvailabilityPtrType)(v)
}

func (*highAvailabilityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HighAvailability)(nil)).Elem()
}

func (i *highAvailabilityPtrType) ToHighAvailabilityPtrOutput() HighAvailabilityPtrOutput {
	return i.ToHighAvailabilityPtrOutputWithContext(context.Background())
}

func (i *highAvailabilityPtrType) ToHighAvailabilityPtrOutputWithContext(ctx context.Context) HighAvailabilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HighAvailabilityPtrOutput)
}

type HighAvailabilityOutput struct{ *pulumi.OutputState }

func (HighAvailabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HighAvailability)(nil)).Elem()
}

func (o HighAvailabilityOutput) ToHighAvailabilityOutput() HighAvailabilityOutput {
	return o
}

func (o HighAvailabilityOutput) ToHighAvailabilityOutputWithContext(ctx context.Context) HighAvailabilityOutput {
	return o
}

func (o HighAvailabilityOutput) ToHighAvailabilityPtrOutput() HighAvailabilityPtrOutput {
	return o.ToHighAvailabilityPtrOutputWithContext(context.Background())
}

func (o HighAvailabilityOutput) ToHighAvailabilityPtrOutputWithContext(ctx context.Context) HighAvailabilityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HighAvailability) *HighAvailability {
		return &v
	}).(HighAvailabilityPtrOutput)
}

func (o HighAvailabilityOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HighAvailability) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o HighAvailabilityOutput) StandbyAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HighAvailability) *string { return v.StandbyAvailabilityZone }).(pulumi.StringPtrOutput)
}

type HighAvailabilityPtrOutput struct{ *pulumi.OutputState }

func (HighAvailabilityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HighAvailability)(nil)).Elem()
}

func (o HighAvailabilityPtrOutput) ToHighAvailabilityPtrOutput() HighAvailabilityPtrOutput {
	return o
}

func (o HighAvailabilityPtrOutput) ToHighAvailabilityPtrOutputWithContext(ctx context.Context) HighAvailabilityPtrOutput {
	return o
}

func (o HighAvailabilityPtrOutput) Elem() HighAvailabilityOutput {
	return o.ApplyT(func(v *HighAvailability) HighAvailability {
		if v != nil {
			return *v
		}
		var ret HighAvailability
		return ret
	}).(HighAvailabilityOutput)
}

func (o HighAvailabilityPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HighAvailability) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o HighAvailabilityPtrOutput) StandbyAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HighAvailability) *string {
		if v == nil {
			return nil
		}
		return v.StandbyAvailabilityZone
	}).(pulumi.StringPtrOutput)
}

type HighAvailabilityResponse struct {
	Mode                    *string `pulumi:"mode"`
	StandbyAvailabilityZone *string `pulumi:"standbyAvailabilityZone"`
	State                   string  `pulumi:"state"`
}





type HighAvailabilityResponseInput interface {
	pulumi.Input

	ToHighAvailabilityResponseOutput() HighAvailabilityResponseOutput
	ToHighAvailabilityResponseOutputWithContext(context.Context) HighAvailabilityResponseOutput
}

type HighAvailabilityResponseArgs struct {
	Mode                    pulumi.StringPtrInput `pulumi:"mode"`
	StandbyAvailabilityZone pulumi.StringPtrInput `pulumi:"standbyAvailabilityZone"`
	State                   pulumi.StringInput    `pulumi:"state"`
}

func (HighAvailabilityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HighAvailabilityResponse)(nil)).Elem()
}

func (i HighAvailabilityResponseArgs) ToHighAvailabilityResponseOutput() HighAvailabilityResponseOutput {
	return i.ToHighAvailabilityResponseOutputWithContext(context.Background())
}

func (i HighAvailabilityResponseArgs) ToHighAvailabilityResponseOutputWithContext(ctx context.Context) HighAvailabilityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HighAvailabilityResponseOutput)
}

func (i HighAvailabilityResponseArgs) ToHighAvailabilityResponsePtrOutput() HighAvailabilityResponsePtrOutput {
	return i.ToHighAvailabilityResponsePtrOutputWithContext(context.Background())
}

func (i HighAvailabilityResponseArgs) ToHighAvailabilityResponsePtrOutputWithContext(ctx context.Context) HighAvailabilityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HighAvailabilityResponseOutput).ToHighAvailabilityResponsePtrOutputWithContext(ctx)
}









type HighAvailabilityResponsePtrInput interface {
	pulumi.Input

	ToHighAvailabilityResponsePtrOutput() HighAvailabilityResponsePtrOutput
	ToHighAvailabilityResponsePtrOutputWithContext(context.Context) HighAvailabilityResponsePtrOutput
}

type highAvailabilityResponsePtrType HighAvailabilityResponseArgs

func HighAvailabilityResponsePtr(v *HighAvailabilityResponseArgs) HighAvailabilityResponsePtrInput {
	return (*highAvailabilityResponsePtrType)(v)
}

func (*highAvailabilityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HighAvailabilityResponse)(nil)).Elem()
}

func (i *highAvailabilityResponsePtrType) ToHighAvailabilityResponsePtrOutput() HighAvailabilityResponsePtrOutput {
	return i.ToHighAvailabilityResponsePtrOutputWithContext(context.Background())
}

func (i *highAvailabilityResponsePtrType) ToHighAvailabilityResponsePtrOutputWithContext(ctx context.Context) HighAvailabilityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HighAvailabilityResponsePtrOutput)
}

type HighAvailabilityResponseOutput struct{ *pulumi.OutputState }

func (HighAvailabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HighAvailabilityResponse)(nil)).Elem()
}

func (o HighAvailabilityResponseOutput) ToHighAvailabilityResponseOutput() HighAvailabilityResponseOutput {
	return o
}

func (o HighAvailabilityResponseOutput) ToHighAvailabilityResponseOutputWithContext(ctx context.Context) HighAvailabilityResponseOutput {
	return o
}

func (o HighAvailabilityResponseOutput) ToHighAvailabilityResponsePtrOutput() HighAvailabilityResponsePtrOutput {
	return o.ToHighAvailabilityResponsePtrOutputWithContext(context.Background())
}

func (o HighAvailabilityResponseOutput) ToHighAvailabilityResponsePtrOutputWithContext(ctx context.Context) HighAvailabilityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HighAvailabilityResponse) *HighAvailabilityResponse {
		return &v
	}).(HighAvailabilityResponsePtrOutput)
}

func (o HighAvailabilityResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HighAvailabilityResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o HighAvailabilityResponseOutput) StandbyAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HighAvailabilityResponse) *string { return v.StandbyAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o HighAvailabilityResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v HighAvailabilityResponse) string { return v.State }).(pulumi.StringOutput)
}

type HighAvailabilityResponsePtrOutput struct{ *pulumi.OutputState }

func (HighAvailabilityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HighAvailabilityResponse)(nil)).Elem()
}

func (o HighAvailabilityResponsePtrOutput) ToHighAvailabilityResponsePtrOutput() HighAvailabilityResponsePtrOutput {
	return o
}

func (o HighAvailabilityResponsePtrOutput) ToHighAvailabilityResponsePtrOutputWithContext(ctx context.Context) HighAvailabilityResponsePtrOutput {
	return o
}

func (o HighAvailabilityResponsePtrOutput) Elem() HighAvailabilityResponseOutput {
	return o.ApplyT(func(v *HighAvailabilityResponse) HighAvailabilityResponse {
		if v != nil {
			return *v
		}
		var ret HighAvailabilityResponse
		return ret
	}).(HighAvailabilityResponseOutput)
}

func (o HighAvailabilityResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HighAvailabilityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o HighAvailabilityResponsePtrOutput) StandbyAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HighAvailabilityResponse) *string {
		if v == nil {
			return nil
		}
		return v.StandbyAvailabilityZone
	}).(pulumi.StringPtrOutput)
}

func (o HighAvailabilityResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HighAvailabilityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type Identity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type MaintenanceWindow struct {
	CustomWindow *string `pulumi:"customWindow"`
	DayOfWeek    *int    `pulumi:"dayOfWeek"`
	StartHour    *int    `pulumi:"startHour"`
	StartMinute  *int    `pulumi:"startMinute"`
}





type MaintenanceWindowInput interface {
	pulumi.Input

	ToMaintenanceWindowOutput() MaintenanceWindowOutput
	ToMaintenanceWindowOutputWithContext(context.Context) MaintenanceWindowOutput
}

type MaintenanceWindowArgs struct {
	CustomWindow pulumi.StringPtrInput `pulumi:"customWindow"`
	DayOfWeek    pulumi.IntPtrInput    `pulumi:"dayOfWeek"`
	StartHour    pulumi.IntPtrInput    `pulumi:"startHour"`
	StartMinute  pulumi.IntPtrInput    `pulumi:"startMinute"`
}

func (MaintenanceWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindowArgs) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return i.ToMaintenanceWindowOutputWithContext(context.Background())
}

func (i MaintenanceWindowArgs) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput)
}

func (i MaintenanceWindowArgs) ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput {
	return i.ToMaintenanceWindowPtrOutputWithContext(context.Background())
}

func (i MaintenanceWindowArgs) ToMaintenanceWindowPtrOutputWithContext(ctx context.Context) MaintenanceWindowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput).ToMaintenanceWindowPtrOutputWithContext(ctx)
}









type MaintenanceWindowPtrInput interface {
	pulumi.Input

	ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput
	ToMaintenanceWindowPtrOutputWithContext(context.Context) MaintenanceWindowPtrOutput
}

type maintenanceWindowPtrType MaintenanceWindowArgs

func MaintenanceWindowPtr(v *MaintenanceWindowArgs) MaintenanceWindowPtrInput {
	return (*maintenanceWindowPtrType)(v)
}

func (*maintenanceWindowPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (i *maintenanceWindowPtrType) ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput {
	return i.ToMaintenanceWindowPtrOutputWithContext(context.Background())
}

func (i *maintenanceWindowPtrType) ToMaintenanceWindowPtrOutputWithContext(ctx context.Context) MaintenanceWindowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowPtrOutput)
}

type MaintenanceWindowOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput {
	return o.ToMaintenanceWindowPtrOutputWithContext(context.Background())
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowPtrOutputWithContext(ctx context.Context) MaintenanceWindowPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MaintenanceWindow) *MaintenanceWindow {
		return &v
	}).(MaintenanceWindowPtrOutput)
}

func (o MaintenanceWindowOutput) CustomWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceWindow) *string { return v.CustomWindow }).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowOutput) DayOfWeek() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindow) *int { return v.DayOfWeek }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowOutput) StartHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindow) *int { return v.StartHour }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindow) *int { return v.StartMinute }).(pulumi.IntPtrOutput)
}

type MaintenanceWindowPtrOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowPtrOutput) ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput {
	return o
}

func (o MaintenanceWindowPtrOutput) ToMaintenanceWindowPtrOutputWithContext(ctx context.Context) MaintenanceWindowPtrOutput {
	return o
}

func (o MaintenanceWindowPtrOutput) Elem() MaintenanceWindowOutput {
	return o.ApplyT(func(v *MaintenanceWindow) MaintenanceWindow {
		if v != nil {
			return *v
		}
		var ret MaintenanceWindow
		return ret
	}).(MaintenanceWindowOutput)
}

func (o MaintenanceWindowPtrOutput) CustomWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) *string {
		if v == nil {
			return nil
		}
		return v.CustomWindow
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowPtrOutput) DayOfWeek() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) *int {
		if v == nil {
			return nil
		}
		return v.DayOfWeek
	}).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowPtrOutput) StartHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) *int {
		if v == nil {
			return nil
		}
		return v.StartHour
	}).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowPtrOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) *int {
		if v == nil {
			return nil
		}
		return v.StartMinute
	}).(pulumi.IntPtrOutput)
}

type MaintenanceWindowResponse struct {
	CustomWindow *string `pulumi:"customWindow"`
	DayOfWeek    *int    `pulumi:"dayOfWeek"`
	StartHour    *int    `pulumi:"startHour"`
	StartMinute  *int    `pulumi:"startMinute"`
}





type MaintenanceWindowResponseInput interface {
	pulumi.Input

	ToMaintenanceWindowResponseOutput() MaintenanceWindowResponseOutput
	ToMaintenanceWindowResponseOutputWithContext(context.Context) MaintenanceWindowResponseOutput
}

type MaintenanceWindowResponseArgs struct {
	CustomWindow pulumi.StringPtrInput `pulumi:"customWindow"`
	DayOfWeek    pulumi.IntPtrInput    `pulumi:"dayOfWeek"`
	StartHour    pulumi.IntPtrInput    `pulumi:"startHour"`
	StartMinute  pulumi.IntPtrInput    `pulumi:"startMinute"`
}

func (MaintenanceWindowResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowResponse)(nil)).Elem()
}

func (i MaintenanceWindowResponseArgs) ToMaintenanceWindowResponseOutput() MaintenanceWindowResponseOutput {
	return i.ToMaintenanceWindowResponseOutputWithContext(context.Background())
}

func (i MaintenanceWindowResponseArgs) ToMaintenanceWindowResponseOutputWithContext(ctx context.Context) MaintenanceWindowResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowResponseOutput)
}

func (i MaintenanceWindowResponseArgs) ToMaintenanceWindowResponsePtrOutput() MaintenanceWindowResponsePtrOutput {
	return i.ToMaintenanceWindowResponsePtrOutputWithContext(context.Background())
}

func (i MaintenanceWindowResponseArgs) ToMaintenanceWindowResponsePtrOutputWithContext(ctx context.Context) MaintenanceWindowResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowResponseOutput).ToMaintenanceWindowResponsePtrOutputWithContext(ctx)
}









type MaintenanceWindowResponsePtrInput interface {
	pulumi.Input

	ToMaintenanceWindowResponsePtrOutput() MaintenanceWindowResponsePtrOutput
	ToMaintenanceWindowResponsePtrOutputWithContext(context.Context) MaintenanceWindowResponsePtrOutput
}

type maintenanceWindowResponsePtrType MaintenanceWindowResponseArgs

func MaintenanceWindowResponsePtr(v *MaintenanceWindowResponseArgs) MaintenanceWindowResponsePtrInput {
	return (*maintenanceWindowResponsePtrType)(v)
}

func (*maintenanceWindowResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindowResponse)(nil)).Elem()
}

func (i *maintenanceWindowResponsePtrType) ToMaintenanceWindowResponsePtrOutput() MaintenanceWindowResponsePtrOutput {
	return i.ToMaintenanceWindowResponsePtrOutputWithContext(context.Background())
}

func (i *maintenanceWindowResponsePtrType) ToMaintenanceWindowResponsePtrOutputWithContext(ctx context.Context) MaintenanceWindowResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowResponsePtrOutput)
}

type MaintenanceWindowResponseOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowResponse)(nil)).Elem()
}

func (o MaintenanceWindowResponseOutput) ToMaintenanceWindowResponseOutput() MaintenanceWindowResponseOutput {
	return o
}

func (o MaintenanceWindowResponseOutput) ToMaintenanceWindowResponseOutputWithContext(ctx context.Context) MaintenanceWindowResponseOutput {
	return o
}

func (o MaintenanceWindowResponseOutput) ToMaintenanceWindowResponsePtrOutput() MaintenanceWindowResponsePtrOutput {
	return o.ToMaintenanceWindowResponsePtrOutputWithContext(context.Background())
}

func (o MaintenanceWindowResponseOutput) ToMaintenanceWindowResponsePtrOutputWithContext(ctx context.Context) MaintenanceWindowResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MaintenanceWindowResponse) *MaintenanceWindowResponse {
		return &v
	}).(MaintenanceWindowResponsePtrOutput)
}

func (o MaintenanceWindowResponseOutput) CustomWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowResponse) *string { return v.CustomWindow }).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowResponseOutput) DayOfWeek() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowResponse) *int { return v.DayOfWeek }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowResponseOutput) StartHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowResponse) *int { return v.StartHour }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowResponseOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowResponse) *int { return v.StartMinute }).(pulumi.IntPtrOutput)
}

type MaintenanceWindowResponsePtrOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindowResponse)(nil)).Elem()
}

func (o MaintenanceWindowResponsePtrOutput) ToMaintenanceWindowResponsePtrOutput() MaintenanceWindowResponsePtrOutput {
	return o
}

func (o MaintenanceWindowResponsePtrOutput) ToMaintenanceWindowResponsePtrOutputWithContext(ctx context.Context) MaintenanceWindowResponsePtrOutput {
	return o
}

func (o MaintenanceWindowResponsePtrOutput) Elem() MaintenanceWindowResponseOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) MaintenanceWindowResponse {
		if v != nil {
			return *v
		}
		var ret MaintenanceWindowResponse
		return ret
	}).(MaintenanceWindowResponseOutput)
}

func (o MaintenanceWindowResponsePtrOutput) CustomWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomWindow
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowResponsePtrOutput) DayOfWeek() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) *int {
		if v == nil {
			return nil
		}
		return v.DayOfWeek
	}).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowResponsePtrOutput) StartHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) *int {
		if v == nil {
			return nil
		}
		return v.StartHour
	}).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowResponsePtrOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) *int {
		if v == nil {
			return nil
		}
		return v.StartMinute
	}).(pulumi.IntPtrOutput)
}

type MigrationResourceGroup struct {
	ResourceId       *string `pulumi:"resourceId"`
	SubnetResourceId *string `pulumi:"subnetResourceId"`
}





type MigrationResourceGroupInput interface {
	pulumi.Input

	ToMigrationResourceGroupOutput() MigrationResourceGroupOutput
	ToMigrationResourceGroupOutputWithContext(context.Context) MigrationResourceGroupOutput
}

type MigrationResourceGroupArgs struct {
	ResourceId       pulumi.StringPtrInput `pulumi:"resourceId"`
	SubnetResourceId pulumi.StringPtrInput `pulumi:"subnetResourceId"`
}

func (MigrationResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationResourceGroup)(nil)).Elem()
}

func (i MigrationResourceGroupArgs) ToMigrationResourceGroupOutput() MigrationResourceGroupOutput {
	return i.ToMigrationResourceGroupOutputWithContext(context.Background())
}

func (i MigrationResourceGroupArgs) ToMigrationResourceGroupOutputWithContext(ctx context.Context) MigrationResourceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationResourceGroupOutput)
}

func (i MigrationResourceGroupArgs) ToMigrationResourceGroupPtrOutput() MigrationResourceGroupPtrOutput {
	return i.ToMigrationResourceGroupPtrOutputWithContext(context.Background())
}

func (i MigrationResourceGroupArgs) ToMigrationResourceGroupPtrOutputWithContext(ctx context.Context) MigrationResourceGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationResourceGroupOutput).ToMigrationResourceGroupPtrOutputWithContext(ctx)
}









type MigrationResourceGroupPtrInput interface {
	pulumi.Input

	ToMigrationResourceGroupPtrOutput() MigrationResourceGroupPtrOutput
	ToMigrationResourceGroupPtrOutputWithContext(context.Context) MigrationResourceGroupPtrOutput
}

type migrationResourceGroupPtrType MigrationResourceGroupArgs

func MigrationResourceGroupPtr(v *MigrationResourceGroupArgs) MigrationResourceGroupPtrInput {
	return (*migrationResourceGroupPtrType)(v)
}

func (*migrationResourceGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationResourceGroup)(nil)).Elem()
}

func (i *migrationResourceGroupPtrType) ToMigrationResourceGroupPtrOutput() MigrationResourceGroupPtrOutput {
	return i.ToMigrationResourceGroupPtrOutputWithContext(context.Background())
}

func (i *migrationResourceGroupPtrType) ToMigrationResourceGroupPtrOutputWithContext(ctx context.Context) MigrationResourceGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationResourceGroupPtrOutput)
}

type MigrationResourceGroupOutput struct{ *pulumi.OutputState }

func (MigrationResourceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationResourceGroup)(nil)).Elem()
}

func (o MigrationResourceGroupOutput) ToMigrationResourceGroupOutput() MigrationResourceGroupOutput {
	return o
}

func (o MigrationResourceGroupOutput) ToMigrationResourceGroupOutputWithContext(ctx context.Context) MigrationResourceGroupOutput {
	return o
}

func (o MigrationResourceGroupOutput) ToMigrationResourceGroupPtrOutput() MigrationResourceGroupPtrOutput {
	return o.ToMigrationResourceGroupPtrOutputWithContext(context.Background())
}

func (o MigrationResourceGroupOutput) ToMigrationResourceGroupPtrOutputWithContext(ctx context.Context) MigrationResourceGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrationResourceGroup) *MigrationResourceGroup {
		return &v
	}).(MigrationResourceGroupPtrOutput)
}

func (o MigrationResourceGroupOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrationResourceGroup) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o MigrationResourceGroupOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrationResourceGroup) *string { return v.SubnetResourceId }).(pulumi.StringPtrOutput)
}

type MigrationResourceGroupPtrOutput struct{ *pulumi.OutputState }

func (MigrationResourceGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationResourceGroup)(nil)).Elem()
}

func (o MigrationResourceGroupPtrOutput) ToMigrationResourceGroupPtrOutput() MigrationResourceGroupPtrOutput {
	return o
}

func (o MigrationResourceGroupPtrOutput) ToMigrationResourceGroupPtrOutputWithContext(ctx context.Context) MigrationResourceGroupPtrOutput {
	return o
}

func (o MigrationResourceGroupPtrOutput) Elem() MigrationResourceGroupOutput {
	return o.ApplyT(func(v *MigrationResourceGroup) MigrationResourceGroup {
		if v != nil {
			return *v
		}
		var ret MigrationResourceGroup
		return ret
	}).(MigrationResourceGroupOutput)
}

func (o MigrationResourceGroupPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationResourceGroup) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

func (o MigrationResourceGroupPtrOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationResourceGroup) *string {
		if v == nil {
			return nil
		}
		return v.SubnetResourceId
	}).(pulumi.StringPtrOutput)
}

type MigrationResourceGroupResponse struct {
	ResourceId       *string `pulumi:"resourceId"`
	SubnetResourceId *string `pulumi:"subnetResourceId"`
}





type MigrationResourceGroupResponseInput interface {
	pulumi.Input

	ToMigrationResourceGroupResponseOutput() MigrationResourceGroupResponseOutput
	ToMigrationResourceGroupResponseOutputWithContext(context.Context) MigrationResourceGroupResponseOutput
}

type MigrationResourceGroupResponseArgs struct {
	ResourceId       pulumi.StringPtrInput `pulumi:"resourceId"`
	SubnetResourceId pulumi.StringPtrInput `pulumi:"subnetResourceId"`
}

func (MigrationResourceGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationResourceGroupResponse)(nil)).Elem()
}

func (i MigrationResourceGroupResponseArgs) ToMigrationResourceGroupResponseOutput() MigrationResourceGroupResponseOutput {
	return i.ToMigrationResourceGroupResponseOutputWithContext(context.Background())
}

func (i MigrationResourceGroupResponseArgs) ToMigrationResourceGroupResponseOutputWithContext(ctx context.Context) MigrationResourceGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationResourceGroupResponseOutput)
}

func (i MigrationResourceGroupResponseArgs) ToMigrationResourceGroupResponsePtrOutput() MigrationResourceGroupResponsePtrOutput {
	return i.ToMigrationResourceGroupResponsePtrOutputWithContext(context.Background())
}

func (i MigrationResourceGroupResponseArgs) ToMigrationResourceGroupResponsePtrOutputWithContext(ctx context.Context) MigrationResourceGroupResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationResourceGroupResponseOutput).ToMigrationResourceGroupResponsePtrOutputWithContext(ctx)
}









type MigrationResourceGroupResponsePtrInput interface {
	pulumi.Input

	ToMigrationResourceGroupResponsePtrOutput() MigrationResourceGroupResponsePtrOutput
	ToMigrationResourceGroupResponsePtrOutputWithContext(context.Context) MigrationResourceGroupResponsePtrOutput
}

type migrationResourceGroupResponsePtrType MigrationResourceGroupResponseArgs

func MigrationResourceGroupResponsePtr(v *MigrationResourceGroupResponseArgs) MigrationResourceGroupResponsePtrInput {
	return (*migrationResourceGroupResponsePtrType)(v)
}

func (*migrationResourceGroupResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationResourceGroupResponse)(nil)).Elem()
}

func (i *migrationResourceGroupResponsePtrType) ToMigrationResourceGroupResponsePtrOutput() MigrationResourceGroupResponsePtrOutput {
	return i.ToMigrationResourceGroupResponsePtrOutputWithContext(context.Background())
}

func (i *migrationResourceGroupResponsePtrType) ToMigrationResourceGroupResponsePtrOutputWithContext(ctx context.Context) MigrationResourceGroupResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationResourceGroupResponsePtrOutput)
}

type MigrationResourceGroupResponseOutput struct{ *pulumi.OutputState }

func (MigrationResourceGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationResourceGroupResponse)(nil)).Elem()
}

func (o MigrationResourceGroupResponseOutput) ToMigrationResourceGroupResponseOutput() MigrationResourceGroupResponseOutput {
	return o
}

func (o MigrationResourceGroupResponseOutput) ToMigrationResourceGroupResponseOutputWithContext(ctx context.Context) MigrationResourceGroupResponseOutput {
	return o
}

func (o MigrationResourceGroupResponseOutput) ToMigrationResourceGroupResponsePtrOutput() MigrationResourceGroupResponsePtrOutput {
	return o.ToMigrationResourceGroupResponsePtrOutputWithContext(context.Background())
}

func (o MigrationResourceGroupResponseOutput) ToMigrationResourceGroupResponsePtrOutputWithContext(ctx context.Context) MigrationResourceGroupResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrationResourceGroupResponse) *MigrationResourceGroupResponse {
		return &v
	}).(MigrationResourceGroupResponsePtrOutput)
}

func (o MigrationResourceGroupResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrationResourceGroupResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o MigrationResourceGroupResponseOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrationResourceGroupResponse) *string { return v.SubnetResourceId }).(pulumi.StringPtrOutput)
}

type MigrationResourceGroupResponsePtrOutput struct{ *pulumi.OutputState }

func (MigrationResourceGroupResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationResourceGroupResponse)(nil)).Elem()
}

func (o MigrationResourceGroupResponsePtrOutput) ToMigrationResourceGroupResponsePtrOutput() MigrationResourceGroupResponsePtrOutput {
	return o
}

func (o MigrationResourceGroupResponsePtrOutput) ToMigrationResourceGroupResponsePtrOutputWithContext(ctx context.Context) MigrationResourceGroupResponsePtrOutput {
	return o
}

func (o MigrationResourceGroupResponsePtrOutput) Elem() MigrationResourceGroupResponseOutput {
	return o.ApplyT(func(v *MigrationResourceGroupResponse) MigrationResourceGroupResponse {
		if v != nil {
			return *v
		}
		var ret MigrationResourceGroupResponse
		return ret
	}).(MigrationResourceGroupResponseOutput)
}

func (o MigrationResourceGroupResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationResourceGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

func (o MigrationResourceGroupResponsePtrOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationResourceGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetResourceId
	}).(pulumi.StringPtrOutput)
}

type MigrationSecretParameters struct {
	AadApp           AADApp           `pulumi:"aadApp"`
	AdminCredentials AdminCredentials `pulumi:"adminCredentials"`
}





type MigrationSecretParametersInput interface {
	pulumi.Input

	ToMigrationSecretParametersOutput() MigrationSecretParametersOutput
	ToMigrationSecretParametersOutputWithContext(context.Context) MigrationSecretParametersOutput
}

type MigrationSecretParametersArgs struct {
	AadApp           AADAppInput           `pulumi:"aadApp"`
	AdminCredentials AdminCredentialsInput `pulumi:"adminCredentials"`
}

func (MigrationSecretParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationSecretParameters)(nil)).Elem()
}

func (i MigrationSecretParametersArgs) ToMigrationSecretParametersOutput() MigrationSecretParametersOutput {
	return i.ToMigrationSecretParametersOutputWithContext(context.Background())
}

func (i MigrationSecretParametersArgs) ToMigrationSecretParametersOutputWithContext(ctx context.Context) MigrationSecretParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationSecretParametersOutput)
}

func (i MigrationSecretParametersArgs) ToMigrationSecretParametersPtrOutput() MigrationSecretParametersPtrOutput {
	return i.ToMigrationSecretParametersPtrOutputWithContext(context.Background())
}

func (i MigrationSecretParametersArgs) ToMigrationSecretParametersPtrOutputWithContext(ctx context.Context) MigrationSecretParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationSecretParametersOutput).ToMigrationSecretParametersPtrOutputWithContext(ctx)
}









type MigrationSecretParametersPtrInput interface {
	pulumi.Input

	ToMigrationSecretParametersPtrOutput() MigrationSecretParametersPtrOutput
	ToMigrationSecretParametersPtrOutputWithContext(context.Context) MigrationSecretParametersPtrOutput
}

type migrationSecretParametersPtrType MigrationSecretParametersArgs

func MigrationSecretParametersPtr(v *MigrationSecretParametersArgs) MigrationSecretParametersPtrInput {
	return (*migrationSecretParametersPtrType)(v)
}

func (*migrationSecretParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationSecretParameters)(nil)).Elem()
}

func (i *migrationSecretParametersPtrType) ToMigrationSecretParametersPtrOutput() MigrationSecretParametersPtrOutput {
	return i.ToMigrationSecretParametersPtrOutputWithContext(context.Background())
}

func (i *migrationSecretParametersPtrType) ToMigrationSecretParametersPtrOutputWithContext(ctx context.Context) MigrationSecretParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationSecretParametersPtrOutput)
}

type MigrationSecretParametersOutput struct{ *pulumi.OutputState }

func (MigrationSecretParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationSecretParameters)(nil)).Elem()
}

func (o MigrationSecretParametersOutput) ToMigrationSecretParametersOutput() MigrationSecretParametersOutput {
	return o
}

func (o MigrationSecretParametersOutput) ToMigrationSecretParametersOutputWithContext(ctx context.Context) MigrationSecretParametersOutput {
	return o
}

func (o MigrationSecretParametersOutput) ToMigrationSecretParametersPtrOutput() MigrationSecretParametersPtrOutput {
	return o.ToMigrationSecretParametersPtrOutputWithContext(context.Background())
}

func (o MigrationSecretParametersOutput) ToMigrationSecretParametersPtrOutputWithContext(ctx context.Context) MigrationSecretParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrationSecretParameters) *MigrationSecretParameters {
		return &v
	}).(MigrationSecretParametersPtrOutput)
}

func (o MigrationSecretParametersOutput) AadApp() AADAppOutput {
	return o.ApplyT(func(v MigrationSecretParameters) AADApp { return v.AadApp }).(AADAppOutput)
}

func (o MigrationSecretParametersOutput) AdminCredentials() AdminCredentialsOutput {
	return o.ApplyT(func(v MigrationSecretParameters) AdminCredentials { return v.AdminCredentials }).(AdminCredentialsOutput)
}

type MigrationSecretParametersPtrOutput struct{ *pulumi.OutputState }

func (MigrationSecretParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationSecretParameters)(nil)).Elem()
}

func (o MigrationSecretParametersPtrOutput) ToMigrationSecretParametersPtrOutput() MigrationSecretParametersPtrOutput {
	return o
}

func (o MigrationSecretParametersPtrOutput) ToMigrationSecretParametersPtrOutputWithContext(ctx context.Context) MigrationSecretParametersPtrOutput {
	return o
}

func (o MigrationSecretParametersPtrOutput) Elem() MigrationSecretParametersOutput {
	return o.ApplyT(func(v *MigrationSecretParameters) MigrationSecretParameters {
		if v != nil {
			return *v
		}
		var ret MigrationSecretParameters
		return ret
	}).(MigrationSecretParametersOutput)
}

func (o MigrationSecretParametersPtrOutput) AadApp() AADAppPtrOutput {
	return o.ApplyT(func(v *MigrationSecretParameters) *AADApp {
		if v == nil {
			return nil
		}
		return &v.AadApp
	}).(AADAppPtrOutput)
}

func (o MigrationSecretParametersPtrOutput) AdminCredentials() AdminCredentialsPtrOutput {
	return o.ApplyT(func(v *MigrationSecretParameters) *AdminCredentials {
		if v == nil {
			return nil
		}
		return &v.AdminCredentials
	}).(AdminCredentialsPtrOutput)
}

type MigrationSecretParametersResponse struct {
	AadApp AADAppResponse `pulumi:"aadApp"`
}





type MigrationSecretParametersResponseInput interface {
	pulumi.Input

	ToMigrationSecretParametersResponseOutput() MigrationSecretParametersResponseOutput
	ToMigrationSecretParametersResponseOutputWithContext(context.Context) MigrationSecretParametersResponseOutput
}

type MigrationSecretParametersResponseArgs struct {
	AadApp AADAppResponseInput `pulumi:"aadApp"`
}

func (MigrationSecretParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationSecretParametersResponse)(nil)).Elem()
}

func (i MigrationSecretParametersResponseArgs) ToMigrationSecretParametersResponseOutput() MigrationSecretParametersResponseOutput {
	return i.ToMigrationSecretParametersResponseOutputWithContext(context.Background())
}

func (i MigrationSecretParametersResponseArgs) ToMigrationSecretParametersResponseOutputWithContext(ctx context.Context) MigrationSecretParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationSecretParametersResponseOutput)
}

func (i MigrationSecretParametersResponseArgs) ToMigrationSecretParametersResponsePtrOutput() MigrationSecretParametersResponsePtrOutput {
	return i.ToMigrationSecretParametersResponsePtrOutputWithContext(context.Background())
}

func (i MigrationSecretParametersResponseArgs) ToMigrationSecretParametersResponsePtrOutputWithContext(ctx context.Context) MigrationSecretParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationSecretParametersResponseOutput).ToMigrationSecretParametersResponsePtrOutputWithContext(ctx)
}









type MigrationSecretParametersResponsePtrInput interface {
	pulumi.Input

	ToMigrationSecretParametersResponsePtrOutput() MigrationSecretParametersResponsePtrOutput
	ToMigrationSecretParametersResponsePtrOutputWithContext(context.Context) MigrationSecretParametersResponsePtrOutput
}

type migrationSecretParametersResponsePtrType MigrationSecretParametersResponseArgs

func MigrationSecretParametersResponsePtr(v *MigrationSecretParametersResponseArgs) MigrationSecretParametersResponsePtrInput {
	return (*migrationSecretParametersResponsePtrType)(v)
}

func (*migrationSecretParametersResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationSecretParametersResponse)(nil)).Elem()
}

func (i *migrationSecretParametersResponsePtrType) ToMigrationSecretParametersResponsePtrOutput() MigrationSecretParametersResponsePtrOutput {
	return i.ToMigrationSecretParametersResponsePtrOutputWithContext(context.Background())
}

func (i *migrationSecretParametersResponsePtrType) ToMigrationSecretParametersResponsePtrOutputWithContext(ctx context.Context) MigrationSecretParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationSecretParametersResponsePtrOutput)
}

type MigrationSecretParametersResponseOutput struct{ *pulumi.OutputState }

func (MigrationSecretParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationSecretParametersResponse)(nil)).Elem()
}

func (o MigrationSecretParametersResponseOutput) ToMigrationSecretParametersResponseOutput() MigrationSecretParametersResponseOutput {
	return o
}

func (o MigrationSecretParametersResponseOutput) ToMigrationSecretParametersResponseOutputWithContext(ctx context.Context) MigrationSecretParametersResponseOutput {
	return o
}

func (o MigrationSecretParametersResponseOutput) ToMigrationSecretParametersResponsePtrOutput() MigrationSecretParametersResponsePtrOutput {
	return o.ToMigrationSecretParametersResponsePtrOutputWithContext(context.Background())
}

func (o MigrationSecretParametersResponseOutput) ToMigrationSecretParametersResponsePtrOutputWithContext(ctx context.Context) MigrationSecretParametersResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrationSecretParametersResponse) *MigrationSecretParametersResponse {
		return &v
	}).(MigrationSecretParametersResponsePtrOutput)
}

func (o MigrationSecretParametersResponseOutput) AadApp() AADAppResponseOutput {
	return o.ApplyT(func(v MigrationSecretParametersResponse) AADAppResponse { return v.AadApp }).(AADAppResponseOutput)
}

type MigrationSecretParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (MigrationSecretParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationSecretParametersResponse)(nil)).Elem()
}

func (o MigrationSecretParametersResponsePtrOutput) ToMigrationSecretParametersResponsePtrOutput() MigrationSecretParametersResponsePtrOutput {
	return o
}

func (o MigrationSecretParametersResponsePtrOutput) ToMigrationSecretParametersResponsePtrOutputWithContext(ctx context.Context) MigrationSecretParametersResponsePtrOutput {
	return o
}

func (o MigrationSecretParametersResponsePtrOutput) Elem() MigrationSecretParametersResponseOutput {
	return o.ApplyT(func(v *MigrationSecretParametersResponse) MigrationSecretParametersResponse {
		if v != nil {
			return *v
		}
		var ret MigrationSecretParametersResponse
		return ret
	}).(MigrationSecretParametersResponseOutput)
}

func (o MigrationSecretParametersResponsePtrOutput) AadApp() AADAppResponsePtrOutput {
	return o.ApplyT(func(v *MigrationSecretParametersResponse) *AADAppResponse {
		if v == nil {
			return nil
		}
		return &v.AadApp
	}).(AADAppResponsePtrOutput)
}

type MigrationStatusResponse struct {
	CurrentSubStateDetails MigrationSubStateDetailsResponse `pulumi:"currentSubStateDetails"`
	Error                  string                           `pulumi:"error"`
	State                  string                           `pulumi:"state"`
}





type MigrationStatusResponseInput interface {
	pulumi.Input

	ToMigrationStatusResponseOutput() MigrationStatusResponseOutput
	ToMigrationStatusResponseOutputWithContext(context.Context) MigrationStatusResponseOutput
}

type MigrationStatusResponseArgs struct {
	CurrentSubStateDetails MigrationSubStateDetailsResponseInput `pulumi:"currentSubStateDetails"`
	Error                  pulumi.StringInput                    `pulumi:"error"`
	State                  pulumi.StringInput                    `pulumi:"state"`
}

func (MigrationStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationStatusResponse)(nil)).Elem()
}

func (i MigrationStatusResponseArgs) ToMigrationStatusResponseOutput() MigrationStatusResponseOutput {
	return i.ToMigrationStatusResponseOutputWithContext(context.Background())
}

func (i MigrationStatusResponseArgs) ToMigrationStatusResponseOutputWithContext(ctx context.Context) MigrationStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationStatusResponseOutput)
}

func (i MigrationStatusResponseArgs) ToMigrationStatusResponsePtrOutput() MigrationStatusResponsePtrOutput {
	return i.ToMigrationStatusResponsePtrOutputWithContext(context.Background())
}

func (i MigrationStatusResponseArgs) ToMigrationStatusResponsePtrOutputWithContext(ctx context.Context) MigrationStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationStatusResponseOutput).ToMigrationStatusResponsePtrOutputWithContext(ctx)
}









type MigrationStatusResponsePtrInput interface {
	pulumi.Input

	ToMigrationStatusResponsePtrOutput() MigrationStatusResponsePtrOutput
	ToMigrationStatusResponsePtrOutputWithContext(context.Context) MigrationStatusResponsePtrOutput
}

type migrationStatusResponsePtrType MigrationStatusResponseArgs

func MigrationStatusResponsePtr(v *MigrationStatusResponseArgs) MigrationStatusResponsePtrInput {
	return (*migrationStatusResponsePtrType)(v)
}

func (*migrationStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationStatusResponse)(nil)).Elem()
}

func (i *migrationStatusResponsePtrType) ToMigrationStatusResponsePtrOutput() MigrationStatusResponsePtrOutput {
	return i.ToMigrationStatusResponsePtrOutputWithContext(context.Background())
}

func (i *migrationStatusResponsePtrType) ToMigrationStatusResponsePtrOutputWithContext(ctx context.Context) MigrationStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationStatusResponsePtrOutput)
}

type MigrationStatusResponseOutput struct{ *pulumi.OutputState }

func (MigrationStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationStatusResponse)(nil)).Elem()
}

func (o MigrationStatusResponseOutput) ToMigrationStatusResponseOutput() MigrationStatusResponseOutput {
	return o
}

func (o MigrationStatusResponseOutput) ToMigrationStatusResponseOutputWithContext(ctx context.Context) MigrationStatusResponseOutput {
	return o
}

func (o MigrationStatusResponseOutput) ToMigrationStatusResponsePtrOutput() MigrationStatusResponsePtrOutput {
	return o.ToMigrationStatusResponsePtrOutputWithContext(context.Background())
}

func (o MigrationStatusResponseOutput) ToMigrationStatusResponsePtrOutputWithContext(ctx context.Context) MigrationStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrationStatusResponse) *MigrationStatusResponse {
		return &v
	}).(MigrationStatusResponsePtrOutput)
}

func (o MigrationStatusResponseOutput) CurrentSubStateDetails() MigrationSubStateDetailsResponseOutput {
	return o.ApplyT(func(v MigrationStatusResponse) MigrationSubStateDetailsResponse { return v.CurrentSubStateDetails }).(MigrationSubStateDetailsResponseOutput)
}

func (o MigrationStatusResponseOutput) Error() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationStatusResponse) string { return v.Error }).(pulumi.StringOutput)
}

func (o MigrationStatusResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationStatusResponse) string { return v.State }).(pulumi.StringOutput)
}

type MigrationStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (MigrationStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationStatusResponse)(nil)).Elem()
}

func (o MigrationStatusResponsePtrOutput) ToMigrationStatusResponsePtrOutput() MigrationStatusResponsePtrOutput {
	return o
}

func (o MigrationStatusResponsePtrOutput) ToMigrationStatusResponsePtrOutputWithContext(ctx context.Context) MigrationStatusResponsePtrOutput {
	return o
}

func (o MigrationStatusResponsePtrOutput) Elem() MigrationStatusResponseOutput {
	return o.ApplyT(func(v *MigrationStatusResponse) MigrationStatusResponse {
		if v != nil {
			return *v
		}
		var ret MigrationStatusResponse
		return ret
	}).(MigrationStatusResponseOutput)
}

func (o MigrationStatusResponsePtrOutput) CurrentSubStateDetails() MigrationSubStateDetailsResponsePtrOutput {
	return o.ApplyT(func(v *MigrationStatusResponse) *MigrationSubStateDetailsResponse {
		if v == nil {
			return nil
		}
		return &v.CurrentSubStateDetails
	}).(MigrationSubStateDetailsResponsePtrOutput)
}

func (o MigrationStatusResponsePtrOutput) Error() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Error
	}).(pulumi.StringPtrOutput)
}

func (o MigrationStatusResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type MigrationSubStateDetailsResponse struct {
	CurrentSubState string `pulumi:"currentSubState"`
}





type MigrationSubStateDetailsResponseInput interface {
	pulumi.Input

	ToMigrationSubStateDetailsResponseOutput() MigrationSubStateDetailsResponseOutput
	ToMigrationSubStateDetailsResponseOutputWithContext(context.Context) MigrationSubStateDetailsResponseOutput
}

type MigrationSubStateDetailsResponseArgs struct {
	CurrentSubState pulumi.StringInput `pulumi:"currentSubState"`
}

func (MigrationSubStateDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationSubStateDetailsResponse)(nil)).Elem()
}

func (i MigrationSubStateDetailsResponseArgs) ToMigrationSubStateDetailsResponseOutput() MigrationSubStateDetailsResponseOutput {
	return i.ToMigrationSubStateDetailsResponseOutputWithContext(context.Background())
}

func (i MigrationSubStateDetailsResponseArgs) ToMigrationSubStateDetailsResponseOutputWithContext(ctx context.Context) MigrationSubStateDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationSubStateDetailsResponseOutput)
}

func (i MigrationSubStateDetailsResponseArgs) ToMigrationSubStateDetailsResponsePtrOutput() MigrationSubStateDetailsResponsePtrOutput {
	return i.ToMigrationSubStateDetailsResponsePtrOutputWithContext(context.Background())
}

func (i MigrationSubStateDetailsResponseArgs) ToMigrationSubStateDetailsResponsePtrOutputWithContext(ctx context.Context) MigrationSubStateDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationSubStateDetailsResponseOutput).ToMigrationSubStateDetailsResponsePtrOutputWithContext(ctx)
}









type MigrationSubStateDetailsResponsePtrInput interface {
	pulumi.Input

	ToMigrationSubStateDetailsResponsePtrOutput() MigrationSubStateDetailsResponsePtrOutput
	ToMigrationSubStateDetailsResponsePtrOutputWithContext(context.Context) MigrationSubStateDetailsResponsePtrOutput
}

type migrationSubStateDetailsResponsePtrType MigrationSubStateDetailsResponseArgs

func MigrationSubStateDetailsResponsePtr(v *MigrationSubStateDetailsResponseArgs) MigrationSubStateDetailsResponsePtrInput {
	return (*migrationSubStateDetailsResponsePtrType)(v)
}

func (*migrationSubStateDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationSubStateDetailsResponse)(nil)).Elem()
}

func (i *migrationSubStateDetailsResponsePtrType) ToMigrationSubStateDetailsResponsePtrOutput() MigrationSubStateDetailsResponsePtrOutput {
	return i.ToMigrationSubStateDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *migrationSubStateDetailsResponsePtrType) ToMigrationSubStateDetailsResponsePtrOutputWithContext(ctx context.Context) MigrationSubStateDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationSubStateDetailsResponsePtrOutput)
}

type MigrationSubStateDetailsResponseOutput struct{ *pulumi.OutputState }

func (MigrationSubStateDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationSubStateDetailsResponse)(nil)).Elem()
}

func (o MigrationSubStateDetailsResponseOutput) ToMigrationSubStateDetailsResponseOutput() MigrationSubStateDetailsResponseOutput {
	return o
}

func (o MigrationSubStateDetailsResponseOutput) ToMigrationSubStateDetailsResponseOutputWithContext(ctx context.Context) MigrationSubStateDetailsResponseOutput {
	return o
}

func (o MigrationSubStateDetailsResponseOutput) ToMigrationSubStateDetailsResponsePtrOutput() MigrationSubStateDetailsResponsePtrOutput {
	return o.ToMigrationSubStateDetailsResponsePtrOutputWithContext(context.Background())
}

func (o MigrationSubStateDetailsResponseOutput) ToMigrationSubStateDetailsResponsePtrOutputWithContext(ctx context.Context) MigrationSubStateDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrationSubStateDetailsResponse) *MigrationSubStateDetailsResponse {
		return &v
	}).(MigrationSubStateDetailsResponsePtrOutput)
}

func (o MigrationSubStateDetailsResponseOutput) CurrentSubState() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationSubStateDetailsResponse) string { return v.CurrentSubState }).(pulumi.StringOutput)
}

type MigrationSubStateDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (MigrationSubStateDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationSubStateDetailsResponse)(nil)).Elem()
}

func (o MigrationSubStateDetailsResponsePtrOutput) ToMigrationSubStateDetailsResponsePtrOutput() MigrationSubStateDetailsResponsePtrOutput {
	return o
}

func (o MigrationSubStateDetailsResponsePtrOutput) ToMigrationSubStateDetailsResponsePtrOutputWithContext(ctx context.Context) MigrationSubStateDetailsResponsePtrOutput {
	return o
}

func (o MigrationSubStateDetailsResponsePtrOutput) Elem() MigrationSubStateDetailsResponseOutput {
	return o.ApplyT(func(v *MigrationSubStateDetailsResponse) MigrationSubStateDetailsResponse {
		if v != nil {
			return *v
		}
		var ret MigrationSubStateDetailsResponse
		return ret
	}).(MigrationSubStateDetailsResponseOutput)
}

func (o MigrationSubStateDetailsResponsePtrOutput) CurrentSubState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationSubStateDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CurrentSubState
	}).(pulumi.StringPtrOutput)
}

type Network struct {
	DelegatedSubnetResourceId   *string `pulumi:"delegatedSubnetResourceId"`
	PrivateDnsZoneArmResourceId *string `pulumi:"privateDnsZoneArmResourceId"`
}





type NetworkInput interface {
	pulumi.Input

	ToNetworkOutput() NetworkOutput
	ToNetworkOutputWithContext(context.Context) NetworkOutput
}

type NetworkArgs struct {
	DelegatedSubnetResourceId   pulumi.StringPtrInput `pulumi:"delegatedSubnetResourceId"`
	PrivateDnsZoneArmResourceId pulumi.StringPtrInput `pulumi:"privateDnsZoneArmResourceId"`
}

func (NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil)).Elem()
}

func (i NetworkArgs) ToNetworkOutput() NetworkOutput {
	return i.ToNetworkOutputWithContext(context.Background())
}

func (i NetworkArgs) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput)
}

func (i NetworkArgs) ToNetworkPtrOutput() NetworkPtrOutput {
	return i.ToNetworkPtrOutputWithContext(context.Background())
}

func (i NetworkArgs) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput).ToNetworkPtrOutputWithContext(ctx)
}









type NetworkPtrInput interface {
	pulumi.Input

	ToNetworkPtrOutput() NetworkPtrOutput
	ToNetworkPtrOutputWithContext(context.Context) NetworkPtrOutput
}

type networkPtrType NetworkArgs

func NetworkPtr(v *NetworkArgs) NetworkPtrInput {
	return (*networkPtrType)(v)
}

func (*networkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil)).Elem()
}

func (i *networkPtrType) ToNetworkPtrOutput() NetworkPtrOutput {
	return i.ToNetworkPtrOutputWithContext(context.Background())
}

func (i *networkPtrType) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPtrOutput)
}

type NetworkOutput struct{ *pulumi.OutputState }

func (NetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil)).Elem()
}

func (o NetworkOutput) ToNetworkOutput() NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkPtrOutput() NetworkPtrOutput {
	return o.ToNetworkPtrOutputWithContext(context.Background())
}

func (o NetworkOutput) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Network) *Network {
		return &v
	}).(NetworkPtrOutput)
}

func (o NetworkOutput) DelegatedSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Network) *string { return v.DelegatedSubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o NetworkOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Network) *string { return v.PrivateDnsZoneArmResourceId }).(pulumi.StringPtrOutput)
}

type NetworkPtrOutput struct{ *pulumi.OutputState }

func (NetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil)).Elem()
}

func (o NetworkPtrOutput) ToNetworkPtrOutput() NetworkPtrOutput {
	return o
}

func (o NetworkPtrOutput) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return o
}

func (o NetworkPtrOutput) Elem() NetworkOutput {
	return o.ApplyT(func(v *Network) Network {
		if v != nil {
			return *v
		}
		var ret Network
		return ret
	}).(NetworkOutput)
}

func (o NetworkPtrOutput) DelegatedSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) *string {
		if v == nil {
			return nil
		}
		return v.DelegatedSubnetResourceId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkPtrOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) *string {
		if v == nil {
			return nil
		}
		return v.PrivateDnsZoneArmResourceId
	}).(pulumi.StringPtrOutput)
}

type NetworkResponse struct {
	DelegatedSubnetResourceId   *string `pulumi:"delegatedSubnetResourceId"`
	PrivateDnsZoneArmResourceId *string `pulumi:"privateDnsZoneArmResourceId"`
	PublicNetworkAccess         string  `pulumi:"publicNetworkAccess"`
}





type NetworkResponseInput interface {
	pulumi.Input

	ToNetworkResponseOutput() NetworkResponseOutput
	ToNetworkResponseOutputWithContext(context.Context) NetworkResponseOutput
}

type NetworkResponseArgs struct {
	DelegatedSubnetResourceId   pulumi.StringPtrInput `pulumi:"delegatedSubnetResourceId"`
	PrivateDnsZoneArmResourceId pulumi.StringPtrInput `pulumi:"privateDnsZoneArmResourceId"`
	PublicNetworkAccess         pulumi.StringInput    `pulumi:"publicNetworkAccess"`
}

func (NetworkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkResponse)(nil)).Elem()
}

func (i NetworkResponseArgs) ToNetworkResponseOutput() NetworkResponseOutput {
	return i.ToNetworkResponseOutputWithContext(context.Background())
}

func (i NetworkResponseArgs) ToNetworkResponseOutputWithContext(ctx context.Context) NetworkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkResponseOutput)
}

func (i NetworkResponseArgs) ToNetworkResponsePtrOutput() NetworkResponsePtrOutput {
	return i.ToNetworkResponsePtrOutputWithContext(context.Background())
}

func (i NetworkResponseArgs) ToNetworkResponsePtrOutputWithContext(ctx context.Context) NetworkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkResponseOutput).ToNetworkResponsePtrOutputWithContext(ctx)
}









type NetworkResponsePtrInput interface {
	pulumi.Input

	ToNetworkResponsePtrOutput() NetworkResponsePtrOutput
	ToNetworkResponsePtrOutputWithContext(context.Context) NetworkResponsePtrOutput
}

type networkResponsePtrType NetworkResponseArgs

func NetworkResponsePtr(v *NetworkResponseArgs) NetworkResponsePtrInput {
	return (*networkResponsePtrType)(v)
}

func (*networkResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkResponse)(nil)).Elem()
}

func (i *networkResponsePtrType) ToNetworkResponsePtrOutput() NetworkResponsePtrOutput {
	return i.ToNetworkResponsePtrOutputWithContext(context.Background())
}

func (i *networkResponsePtrType) ToNetworkResponsePtrOutputWithContext(ctx context.Context) NetworkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkResponsePtrOutput)
}

type NetworkResponseOutput struct{ *pulumi.OutputState }

func (NetworkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkResponse)(nil)).Elem()
}

func (o NetworkResponseOutput) ToNetworkResponseOutput() NetworkResponseOutput {
	return o
}

func (o NetworkResponseOutput) ToNetworkResponseOutputWithContext(ctx context.Context) NetworkResponseOutput {
	return o
}

func (o NetworkResponseOutput) ToNetworkResponsePtrOutput() NetworkResponsePtrOutput {
	return o.ToNetworkResponsePtrOutputWithContext(context.Background())
}

func (o NetworkResponseOutput) ToNetworkResponsePtrOutputWithContext(ctx context.Context) NetworkResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkResponse) *NetworkResponse {
		return &v
	}).(NetworkResponsePtrOutput)
}

func (o NetworkResponseOutput) DelegatedSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkResponse) *string { return v.DelegatedSubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o NetworkResponseOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkResponse) *string { return v.PrivateDnsZoneArmResourceId }).(pulumi.StringPtrOutput)
}

func (o NetworkResponseOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkResponse) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

type NetworkResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkResponse)(nil)).Elem()
}

func (o NetworkResponsePtrOutput) ToNetworkResponsePtrOutput() NetworkResponsePtrOutput {
	return o
}

func (o NetworkResponsePtrOutput) ToNetworkResponsePtrOutputWithContext(ctx context.Context) NetworkResponsePtrOutput {
	return o
}

func (o NetworkResponsePtrOutput) Elem() NetworkResponseOutput {
	return o.ApplyT(func(v *NetworkResponse) NetworkResponse {
		if v != nil {
			return *v
		}
		var ret NetworkResponse
		return ret
	}).(NetworkResponseOutput)
}

func (o NetworkResponsePtrOutput) DelegatedSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkResponse) *string {
		if v == nil {
			return nil
		}
		return v.DelegatedSubnetResourceId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkResponsePtrOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateDnsZoneArmResourceId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type ServerSkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}





type ServerSkuResponseInput interface {
	pulumi.Input

	ToServerSkuResponseOutput() ServerSkuResponseOutput
	ToServerSkuResponseOutputWithContext(context.Context) ServerSkuResponseOutput
}

type ServerSkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (ServerSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerSkuResponse)(nil)).Elem()
}

func (i ServerSkuResponseArgs) ToServerSkuResponseOutput() ServerSkuResponseOutput {
	return i.ToServerSkuResponseOutputWithContext(context.Background())
}

func (i ServerSkuResponseArgs) ToServerSkuResponseOutputWithContext(ctx context.Context) ServerSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerSkuResponseOutput)
}

func (i ServerSkuResponseArgs) ToServerSkuResponsePtrOutput() ServerSkuResponsePtrOutput {
	return i.ToServerSkuResponsePtrOutputWithContext(context.Background())
}

func (i ServerSkuResponseArgs) ToServerSkuResponsePtrOutputWithContext(ctx context.Context) ServerSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerSkuResponseOutput).ToServerSkuResponsePtrOutputWithContext(ctx)
}









type ServerSkuResponsePtrInput interface {
	pulumi.Input

	ToServerSkuResponsePtrOutput() ServerSkuResponsePtrOutput
	ToServerSkuResponsePtrOutputWithContext(context.Context) ServerSkuResponsePtrOutput
}

type serverSkuResponsePtrType ServerSkuResponseArgs

func ServerSkuResponsePtr(v *ServerSkuResponseArgs) ServerSkuResponsePtrInput {
	return (*serverSkuResponsePtrType)(v)
}

func (*serverSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerSkuResponse)(nil)).Elem()
}

func (i *serverSkuResponsePtrType) ToServerSkuResponsePtrOutput() ServerSkuResponsePtrOutput {
	return i.ToServerSkuResponsePtrOutputWithContext(context.Background())
}

func (i *serverSkuResponsePtrType) ToServerSkuResponsePtrOutputWithContext(ctx context.Context) ServerSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerSkuResponsePtrOutput)
}

type ServerSkuResponseOutput struct{ *pulumi.OutputState }

func (ServerSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerSkuResponse)(nil)).Elem()
}

func (o ServerSkuResponseOutput) ToServerSkuResponseOutput() ServerSkuResponseOutput {
	return o
}

func (o ServerSkuResponseOutput) ToServerSkuResponseOutputWithContext(ctx context.Context) ServerSkuResponseOutput {
	return o
}

func (o ServerSkuResponseOutput) ToServerSkuResponsePtrOutput() ServerSkuResponsePtrOutput {
	return o.ToServerSkuResponsePtrOutputWithContext(context.Background())
}

func (o ServerSkuResponseOutput) ToServerSkuResponsePtrOutputWithContext(ctx context.Context) ServerSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerSkuResponse) *ServerSkuResponse {
		return &v
	}).(ServerSkuResponsePtrOutput)
}

func (o ServerSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServerSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ServerSkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v ServerSkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type ServerSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerSkuResponse)(nil)).Elem()
}

func (o ServerSkuResponsePtrOutput) ToServerSkuResponsePtrOutput() ServerSkuResponsePtrOutput {
	return o
}

func (o ServerSkuResponsePtrOutput) ToServerSkuResponsePtrOutputWithContext(ctx context.Context) ServerSkuResponsePtrOutput {
	return o
}

func (o ServerSkuResponsePtrOutput) Elem() ServerSkuResponseOutput {
	return o.ApplyT(func(v *ServerSkuResponse) ServerSkuResponse {
		if v != nil {
			return *v
		}
		var ret ServerSkuResponse
		return ret
	}).(ServerSkuResponseOutput)
}

func (o ServerSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ServerSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type Storage struct {
	StorageSizeGB *int `pulumi:"storageSizeGB"`
}





type StorageInput interface {
	pulumi.Input

	ToStorageOutput() StorageOutput
	ToStorageOutputWithContext(context.Context) StorageOutput
}

type StorageArgs struct {
	StorageSizeGB pulumi.IntPtrInput `pulumi:"storageSizeGB"`
}

func (StorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Storage)(nil)).Elem()
}

func (i StorageArgs) ToStorageOutput() StorageOutput {
	return i.ToStorageOutputWithContext(context.Background())
}

func (i StorageArgs) ToStorageOutputWithContext(ctx context.Context) StorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageOutput)
}

func (i StorageArgs) ToStoragePtrOutput() StoragePtrOutput {
	return i.ToStoragePtrOutputWithContext(context.Background())
}

func (i StorageArgs) ToStoragePtrOutputWithContext(ctx context.Context) StoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageOutput).ToStoragePtrOutputWithContext(ctx)
}









type StoragePtrInput interface {
	pulumi.Input

	ToStoragePtrOutput() StoragePtrOutput
	ToStoragePtrOutputWithContext(context.Context) StoragePtrOutput
}

type storagePtrType StorageArgs

func StoragePtr(v *StorageArgs) StoragePtrInput {
	return (*storagePtrType)(v)
}

func (*storagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Storage)(nil)).Elem()
}

func (i *storagePtrType) ToStoragePtrOutput() StoragePtrOutput {
	return i.ToStoragePtrOutputWithContext(context.Background())
}

func (i *storagePtrType) ToStoragePtrOutputWithContext(ctx context.Context) StoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoragePtrOutput)
}

type StorageOutput struct{ *pulumi.OutputState }

func (StorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Storage)(nil)).Elem()
}

func (o StorageOutput) ToStorageOutput() StorageOutput {
	return o
}

func (o StorageOutput) ToStorageOutputWithContext(ctx context.Context) StorageOutput {
	return o
}

func (o StorageOutput) ToStoragePtrOutput() StoragePtrOutput {
	return o.ToStoragePtrOutputWithContext(context.Background())
}

func (o StorageOutput) ToStoragePtrOutputWithContext(ctx context.Context) StoragePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Storage) *Storage {
		return &v
	}).(StoragePtrOutput)
}

func (o StorageOutput) StorageSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Storage) *int { return v.StorageSizeGB }).(pulumi.IntPtrOutput)
}

type StoragePtrOutput struct{ *pulumi.OutputState }

func (StoragePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Storage)(nil)).Elem()
}

func (o StoragePtrOutput) ToStoragePtrOutput() StoragePtrOutput {
	return o
}

func (o StoragePtrOutput) ToStoragePtrOutputWithContext(ctx context.Context) StoragePtrOutput {
	return o
}

func (o StoragePtrOutput) Elem() StorageOutput {
	return o.ApplyT(func(v *Storage) Storage {
		if v != nil {
			return *v
		}
		var ret Storage
		return ret
	}).(StorageOutput)
}

func (o StoragePtrOutput) StorageSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Storage) *int {
		if v == nil {
			return nil
		}
		return v.StorageSizeGB
	}).(pulumi.IntPtrOutput)
}

type StorageResponse struct {
	StorageSizeGB *int `pulumi:"storageSizeGB"`
}





type StorageResponseInput interface {
	pulumi.Input

	ToStorageResponseOutput() StorageResponseOutput
	ToStorageResponseOutputWithContext(context.Context) StorageResponseOutput
}

type StorageResponseArgs struct {
	StorageSizeGB pulumi.IntPtrInput `pulumi:"storageSizeGB"`
}

func (StorageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageResponse)(nil)).Elem()
}

func (i StorageResponseArgs) ToStorageResponseOutput() StorageResponseOutput {
	return i.ToStorageResponseOutputWithContext(context.Background())
}

func (i StorageResponseArgs) ToStorageResponseOutputWithContext(ctx context.Context) StorageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageResponseOutput)
}

func (i StorageResponseArgs) ToStorageResponsePtrOutput() StorageResponsePtrOutput {
	return i.ToStorageResponsePtrOutputWithContext(context.Background())
}

func (i StorageResponseArgs) ToStorageResponsePtrOutputWithContext(ctx context.Context) StorageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageResponseOutput).ToStorageResponsePtrOutputWithContext(ctx)
}









type StorageResponsePtrInput interface {
	pulumi.Input

	ToStorageResponsePtrOutput() StorageResponsePtrOutput
	ToStorageResponsePtrOutputWithContext(context.Context) StorageResponsePtrOutput
}

type storageResponsePtrType StorageResponseArgs

func StorageResponsePtr(v *StorageResponseArgs) StorageResponsePtrInput {
	return (*storageResponsePtrType)(v)
}

func (*storageResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageResponse)(nil)).Elem()
}

func (i *storageResponsePtrType) ToStorageResponsePtrOutput() StorageResponsePtrOutput {
	return i.ToStorageResponsePtrOutputWithContext(context.Background())
}

func (i *storageResponsePtrType) ToStorageResponsePtrOutputWithContext(ctx context.Context) StorageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageResponsePtrOutput)
}

type StorageResponseOutput struct{ *pulumi.OutputState }

func (StorageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageResponse)(nil)).Elem()
}

func (o StorageResponseOutput) ToStorageResponseOutput() StorageResponseOutput {
	return o
}

func (o StorageResponseOutput) ToStorageResponseOutputWithContext(ctx context.Context) StorageResponseOutput {
	return o
}

func (o StorageResponseOutput) ToStorageResponsePtrOutput() StorageResponsePtrOutput {
	return o.ToStorageResponsePtrOutputWithContext(context.Background())
}

func (o StorageResponseOutput) ToStorageResponsePtrOutputWithContext(ctx context.Context) StorageResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageResponse) *StorageResponse {
		return &v
	}).(StorageResponsePtrOutput)
}

func (o StorageResponseOutput) StorageSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageResponse) *int { return v.StorageSizeGB }).(pulumi.IntPtrOutput)
}

type StorageResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageResponse)(nil)).Elem()
}

func (o StorageResponsePtrOutput) ToStorageResponsePtrOutput() StorageResponsePtrOutput {
	return o
}

func (o StorageResponsePtrOutput) ToStorageResponsePtrOutputWithContext(ctx context.Context) StorageResponsePtrOutput {
	return o
}

func (o StorageResponsePtrOutput) Elem() StorageResponseOutput {
	return o.ApplyT(func(v *StorageResponse) StorageResponse {
		if v != nil {
			return *v
		}
		var ret StorageResponse
		return ret
	}).(StorageResponseOutput)
}

func (o StorageResponsePtrOutput) StorageSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageResponse) *int {
		if v == nil {
			return nil
		}
		return v.StorageSizeGB
	}).(pulumi.IntPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AADAppOutput{})
	pulumi.RegisterOutputType(AADAppPtrOutput{})
	pulumi.RegisterOutputType(AADAppResponseOutput{})
	pulumi.RegisterOutputType(AADAppResponsePtrOutput{})
	pulumi.RegisterOutputType(AdminCredentialsOutput{})
	pulumi.RegisterOutputType(AdminCredentialsPtrOutput{})
	pulumi.RegisterOutputType(BackupOutput{})
	pulumi.RegisterOutputType(BackupPtrOutput{})
	pulumi.RegisterOutputType(BackupResponseOutput{})
	pulumi.RegisterOutputType(BackupResponsePtrOutput{})
	pulumi.RegisterOutputType(DBServerMetadataResponseOutput{})
	pulumi.RegisterOutputType(DBServerMetadataResponsePtrOutput{})
	pulumi.RegisterOutputType(HighAvailabilityOutput{})
	pulumi.RegisterOutputType(HighAvailabilityPtrOutput{})
	pulumi.RegisterOutputType(HighAvailabilityResponseOutput{})
	pulumi.RegisterOutputType(HighAvailabilityResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowPtrOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowResponseOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowResponsePtrOutput{})
	pulumi.RegisterOutputType(MigrationResourceGroupOutput{})
	pulumi.RegisterOutputType(MigrationResourceGroupPtrOutput{})
	pulumi.RegisterOutputType(MigrationResourceGroupResponseOutput{})
	pulumi.RegisterOutputType(MigrationResourceGroupResponsePtrOutput{})
	pulumi.RegisterOutputType(MigrationSecretParametersOutput{})
	pulumi.RegisterOutputType(MigrationSecretParametersPtrOutput{})
	pulumi.RegisterOutputType(MigrationSecretParametersResponseOutput{})
	pulumi.RegisterOutputType(MigrationSecretParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(MigrationStatusResponseOutput{})
	pulumi.RegisterOutputType(MigrationStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(MigrationSubStateDetailsResponseOutput{})
	pulumi.RegisterOutputType(MigrationSubStateDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkOutput{})
	pulumi.RegisterOutputType(NetworkPtrOutput{})
	pulumi.RegisterOutputType(NetworkResponseOutput{})
	pulumi.RegisterOutputType(NetworkResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerSkuResponseOutput{})
	pulumi.RegisterOutputType(ServerSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageOutput{})
	pulumi.RegisterOutputType(StoragePtrOutput{})
	pulumi.RegisterOutputType(StorageResponseOutput{})
	pulumi.RegisterOutputType(StorageResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}

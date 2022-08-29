


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationServerConfiguration struct {
	InstanceCount               float64                     `pulumi:"instanceCount"`
	SubnetId                    string                      `pulumi:"subnetId"`
	VirtualMachineConfiguration VirtualMachineConfiguration `pulumi:"virtualMachineConfiguration"`
}

type ApplicationServerConfigurationResponse struct {
	InstanceCount               float64                             `pulumi:"instanceCount"`
	SubnetId                    string                              `pulumi:"subnetId"`
	VirtualMachineConfiguration VirtualMachineConfigurationResponse `pulumi:"virtualMachineConfiguration"`
}

type BackupProfile struct {
	BackupEnabled string `pulumi:"backupEnabled"`
}





type BackupProfileInput interface {
	pulumi.Input

	ToBackupProfileOutput() BackupProfileOutput
	ToBackupProfileOutputWithContext(context.Context) BackupProfileOutput
}

type BackupProfileArgs struct {
	BackupEnabled pulumi.StringInput `pulumi:"backupEnabled"`
}

func (BackupProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupProfile)(nil)).Elem()
}

func (i BackupProfileArgs) ToBackupProfileOutput() BackupProfileOutput {
	return i.ToBackupProfileOutputWithContext(context.Background())
}

func (i BackupProfileArgs) ToBackupProfileOutputWithContext(ctx context.Context) BackupProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupProfileOutput)
}

func (i BackupProfileArgs) ToBackupProfilePtrOutput() BackupProfilePtrOutput {
	return i.ToBackupProfilePtrOutputWithContext(context.Background())
}

func (i BackupProfileArgs) ToBackupProfilePtrOutputWithContext(ctx context.Context) BackupProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupProfileOutput).ToBackupProfilePtrOutputWithContext(ctx)
}









type BackupProfilePtrInput interface {
	pulumi.Input

	ToBackupProfilePtrOutput() BackupProfilePtrOutput
	ToBackupProfilePtrOutputWithContext(context.Context) BackupProfilePtrOutput
}

type backupProfilePtrType BackupProfileArgs

func BackupProfilePtr(v *BackupProfileArgs) BackupProfilePtrInput {
	return (*backupProfilePtrType)(v)
}

func (*backupProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupProfile)(nil)).Elem()
}

func (i *backupProfilePtrType) ToBackupProfilePtrOutput() BackupProfilePtrOutput {
	return i.ToBackupProfilePtrOutputWithContext(context.Background())
}

func (i *backupProfilePtrType) ToBackupProfilePtrOutputWithContext(ctx context.Context) BackupProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupProfilePtrOutput)
}

type BackupProfileOutput struct{ *pulumi.OutputState }

func (BackupProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupProfile)(nil)).Elem()
}

func (o BackupProfileOutput) ToBackupProfileOutput() BackupProfileOutput {
	return o
}

func (o BackupProfileOutput) ToBackupProfileOutputWithContext(ctx context.Context) BackupProfileOutput {
	return o
}

func (o BackupProfileOutput) ToBackupProfilePtrOutput() BackupProfilePtrOutput {
	return o.ToBackupProfilePtrOutputWithContext(context.Background())
}

func (o BackupProfileOutput) ToBackupProfilePtrOutputWithContext(ctx context.Context) BackupProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupProfile) *BackupProfile {
		return &v
	}).(BackupProfilePtrOutput)
}

func (o BackupProfileOutput) BackupEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v BackupProfile) string { return v.BackupEnabled }).(pulumi.StringOutput)
}

type BackupProfilePtrOutput struct{ *pulumi.OutputState }

func (BackupProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupProfile)(nil)).Elem()
}

func (o BackupProfilePtrOutput) ToBackupProfilePtrOutput() BackupProfilePtrOutput {
	return o
}

func (o BackupProfilePtrOutput) ToBackupProfilePtrOutputWithContext(ctx context.Context) BackupProfilePtrOutput {
	return o
}

func (o BackupProfilePtrOutput) Elem() BackupProfileOutput {
	return o.ApplyT(func(v *BackupProfile) BackupProfile {
		if v != nil {
			return *v
		}
		var ret BackupProfile
		return ret
	}).(BackupProfileOutput)
}

func (o BackupProfilePtrOutput) BackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupProfile) *string {
		if v == nil {
			return nil
		}
		return &v.BackupEnabled
	}).(pulumi.StringPtrOutput)
}

type BackupProfileResponse struct {
	BackupEnabled   string `pulumi:"backupEnabled"`
	VaultResourceId string `pulumi:"vaultResourceId"`
}

type BackupProfileResponseOutput struct{ *pulumi.OutputState }

func (BackupProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupProfileResponse)(nil)).Elem()
}

func (o BackupProfileResponseOutput) ToBackupProfileResponseOutput() BackupProfileResponseOutput {
	return o
}

func (o BackupProfileResponseOutput) ToBackupProfileResponseOutputWithContext(ctx context.Context) BackupProfileResponseOutput {
	return o
}

func (o BackupProfileResponseOutput) BackupEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v BackupProfileResponse) string { return v.BackupEnabled }).(pulumi.StringOutput)
}

func (o BackupProfileResponseOutput) VaultResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v BackupProfileResponse) string { return v.VaultResourceId }).(pulumi.StringOutput)
}

type BackupProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (BackupProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupProfileResponse)(nil)).Elem()
}

func (o BackupProfileResponsePtrOutput) ToBackupProfileResponsePtrOutput() BackupProfileResponsePtrOutput {
	return o
}

func (o BackupProfileResponsePtrOutput) ToBackupProfileResponsePtrOutputWithContext(ctx context.Context) BackupProfileResponsePtrOutput {
	return o
}

func (o BackupProfileResponsePtrOutput) Elem() BackupProfileResponseOutput {
	return o.ApplyT(func(v *BackupProfileResponse) BackupProfileResponse {
		if v != nil {
			return *v
		}
		var ret BackupProfileResponse
		return ret
	}).(BackupProfileResponseOutput)
}

func (o BackupProfileResponsePtrOutput) BackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.BackupEnabled
	}).(pulumi.StringPtrOutput)
}

func (o BackupProfileResponsePtrOutput) VaultResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VaultResourceId
	}).(pulumi.StringPtrOutput)
}

type CacheProfile struct {
	Capacity float64 `pulumi:"capacity"`
	Family   string  `pulumi:"family"`
	Name     *string `pulumi:"name"`
	SkuName  string  `pulumi:"skuName"`
}





type CacheProfileInput interface {
	pulumi.Input

	ToCacheProfileOutput() CacheProfileOutput
	ToCacheProfileOutputWithContext(context.Context) CacheProfileOutput
}

type CacheProfileArgs struct {
	Capacity pulumi.Float64Input   `pulumi:"capacity"`
	Family   pulumi.StringInput    `pulumi:"family"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	SkuName  pulumi.StringInput    `pulumi:"skuName"`
}

func (CacheProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheProfile)(nil)).Elem()
}

func (i CacheProfileArgs) ToCacheProfileOutput() CacheProfileOutput {
	return i.ToCacheProfileOutputWithContext(context.Background())
}

func (i CacheProfileArgs) ToCacheProfileOutputWithContext(ctx context.Context) CacheProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheProfileOutput)
}

func (i CacheProfileArgs) ToCacheProfilePtrOutput() CacheProfilePtrOutput {
	return i.ToCacheProfilePtrOutputWithContext(context.Background())
}

func (i CacheProfileArgs) ToCacheProfilePtrOutputWithContext(ctx context.Context) CacheProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheProfileOutput).ToCacheProfilePtrOutputWithContext(ctx)
}









type CacheProfilePtrInput interface {
	pulumi.Input

	ToCacheProfilePtrOutput() CacheProfilePtrOutput
	ToCacheProfilePtrOutputWithContext(context.Context) CacheProfilePtrOutput
}

type cacheProfilePtrType CacheProfileArgs

func CacheProfilePtr(v *CacheProfileArgs) CacheProfilePtrInput {
	return (*cacheProfilePtrType)(v)
}

func (*cacheProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheProfile)(nil)).Elem()
}

func (i *cacheProfilePtrType) ToCacheProfilePtrOutput() CacheProfilePtrOutput {
	return i.ToCacheProfilePtrOutputWithContext(context.Background())
}

func (i *cacheProfilePtrType) ToCacheProfilePtrOutputWithContext(ctx context.Context) CacheProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheProfilePtrOutput)
}

type CacheProfileOutput struct{ *pulumi.OutputState }

func (CacheProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheProfile)(nil)).Elem()
}

func (o CacheProfileOutput) ToCacheProfileOutput() CacheProfileOutput {
	return o
}

func (o CacheProfileOutput) ToCacheProfileOutputWithContext(ctx context.Context) CacheProfileOutput {
	return o
}

func (o CacheProfileOutput) ToCacheProfilePtrOutput() CacheProfilePtrOutput {
	return o.ToCacheProfilePtrOutputWithContext(context.Background())
}

func (o CacheProfileOutput) ToCacheProfilePtrOutputWithContext(ctx context.Context) CacheProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheProfile) *CacheProfile {
		return &v
	}).(CacheProfilePtrOutput)
}

func (o CacheProfileOutput) Capacity() pulumi.Float64Output {
	return o.ApplyT(func(v CacheProfile) float64 { return v.Capacity }).(pulumi.Float64Output)
}

func (o CacheProfileOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v CacheProfile) string { return v.Family }).(pulumi.StringOutput)
}

func (o CacheProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CacheProfileOutput) SkuName() pulumi.StringOutput {
	return o.ApplyT(func(v CacheProfile) string { return v.SkuName }).(pulumi.StringOutput)
}

type CacheProfilePtrOutput struct{ *pulumi.OutputState }

func (CacheProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheProfile)(nil)).Elem()
}

func (o CacheProfilePtrOutput) ToCacheProfilePtrOutput() CacheProfilePtrOutput {
	return o
}

func (o CacheProfilePtrOutput) ToCacheProfilePtrOutputWithContext(ctx context.Context) CacheProfilePtrOutput {
	return o
}

func (o CacheProfilePtrOutput) Elem() CacheProfileOutput {
	return o.ApplyT(func(v *CacheProfile) CacheProfile {
		if v != nil {
			return *v
		}
		var ret CacheProfile
		return ret
	}).(CacheProfileOutput)
}

func (o CacheProfilePtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CacheProfile) *float64 {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o CacheProfilePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheProfile) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

func (o CacheProfilePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheProfile) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CacheProfilePtrOutput) SkuName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheProfile) *string {
		if v == nil {
			return nil
		}
		return &v.SkuName
	}).(pulumi.StringPtrOutput)
}

type CacheProfileResponse struct {
	CacheResourceId string  `pulumi:"cacheResourceId"`
	Capacity        float64 `pulumi:"capacity"`
	Family          string  `pulumi:"family"`
	Name            *string `pulumi:"name"`
	SkuName         string  `pulumi:"skuName"`
}

type CacheProfileResponseOutput struct{ *pulumi.OutputState }

func (CacheProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheProfileResponse)(nil)).Elem()
}

func (o CacheProfileResponseOutput) ToCacheProfileResponseOutput() CacheProfileResponseOutput {
	return o
}

func (o CacheProfileResponseOutput) ToCacheProfileResponseOutputWithContext(ctx context.Context) CacheProfileResponseOutput {
	return o
}

func (o CacheProfileResponseOutput) CacheResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v CacheProfileResponse) string { return v.CacheResourceId }).(pulumi.StringOutput)
}

func (o CacheProfileResponseOutput) Capacity() pulumi.Float64Output {
	return o.ApplyT(func(v CacheProfileResponse) float64 { return v.Capacity }).(pulumi.Float64Output)
}

func (o CacheProfileResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v CacheProfileResponse) string { return v.Family }).(pulumi.StringOutput)
}

func (o CacheProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CacheProfileResponseOutput) SkuName() pulumi.StringOutput {
	return o.ApplyT(func(v CacheProfileResponse) string { return v.SkuName }).(pulumi.StringOutput)
}

type CacheProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (CacheProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheProfileResponse)(nil)).Elem()
}

func (o CacheProfileResponsePtrOutput) ToCacheProfileResponsePtrOutput() CacheProfileResponsePtrOutput {
	return o
}

func (o CacheProfileResponsePtrOutput) ToCacheProfileResponsePtrOutputWithContext(ctx context.Context) CacheProfileResponsePtrOutput {
	return o
}

func (o CacheProfileResponsePtrOutput) Elem() CacheProfileResponseOutput {
	return o.ApplyT(func(v *CacheProfileResponse) CacheProfileResponse {
		if v != nil {
			return *v
		}
		var ret CacheProfileResponse
		return ret
	}).(CacheProfileResponseOutput)
}

func (o CacheProfileResponsePtrOutput) CacheResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CacheResourceId
	}).(pulumi.StringPtrOutput)
}

func (o CacheProfileResponsePtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CacheProfileResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o CacheProfileResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

func (o CacheProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CacheProfileResponsePtrOutput) SkuName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SkuName
	}).(pulumi.StringPtrOutput)
}

type CentralServerConfiguration struct {
	InstanceCount               float64                     `pulumi:"instanceCount"`
	SubnetId                    string                      `pulumi:"subnetId"`
	VirtualMachineConfiguration VirtualMachineConfiguration `pulumi:"virtualMachineConfiguration"`
}

type CentralServerConfigurationResponse struct {
	InstanceCount               float64                             `pulumi:"instanceCount"`
	SubnetId                    string                              `pulumi:"subnetId"`
	VirtualMachineConfiguration VirtualMachineConfigurationResponse `pulumi:"virtualMachineConfiguration"`
}

type CentralServerVmDetailsResponse struct {
	Type             string `pulumi:"type"`
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

type CentralServerVmDetailsResponseOutput struct{ *pulumi.OutputState }

func (CentralServerVmDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CentralServerVmDetailsResponse)(nil)).Elem()
}

func (o CentralServerVmDetailsResponseOutput) ToCentralServerVmDetailsResponseOutput() CentralServerVmDetailsResponseOutput {
	return o
}

func (o CentralServerVmDetailsResponseOutput) ToCentralServerVmDetailsResponseOutputWithContext(ctx context.Context) CentralServerVmDetailsResponseOutput {
	return o
}

func (o CentralServerVmDetailsResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v CentralServerVmDetailsResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o CentralServerVmDetailsResponseOutput) VirtualMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v CentralServerVmDetailsResponse) string { return v.VirtualMachineId }).(pulumi.StringOutput)
}

type CentralServerVmDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (CentralServerVmDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CentralServerVmDetailsResponse)(nil)).Elem()
}

func (o CentralServerVmDetailsResponseArrayOutput) ToCentralServerVmDetailsResponseArrayOutput() CentralServerVmDetailsResponseArrayOutput {
	return o
}

func (o CentralServerVmDetailsResponseArrayOutput) ToCentralServerVmDetailsResponseArrayOutputWithContext(ctx context.Context) CentralServerVmDetailsResponseArrayOutput {
	return o
}

func (o CentralServerVmDetailsResponseArrayOutput) Index(i pulumi.IntInput) CentralServerVmDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CentralServerVmDetailsResponse {
		return vs[0].([]CentralServerVmDetailsResponse)[vs[1].(int)]
	}).(CentralServerVmDetailsResponseOutput)
}

type DB2ProviderInstanceProperties struct {
	DbName        *string `pulumi:"dbName"`
	DbPassword    *string `pulumi:"dbPassword"`
	DbPasswordUri *string `pulumi:"dbPasswordUri"`
	DbPort        *string `pulumi:"dbPort"`
	DbUsername    *string `pulumi:"dbUsername"`
	Hostname      *string `pulumi:"hostname"`
	ProviderType  string  `pulumi:"providerType"`
	SapSid        *string `pulumi:"sapSid"`
}

type DB2ProviderInstancePropertiesResponse struct {
	DbName        *string `pulumi:"dbName"`
	DbPassword    *string `pulumi:"dbPassword"`
	DbPasswordUri *string `pulumi:"dbPasswordUri"`
	DbPort        *string `pulumi:"dbPort"`
	DbUsername    *string `pulumi:"dbUsername"`
	Hostname      *string `pulumi:"hostname"`
	ProviderType  string  `pulumi:"providerType"`
	SapSid        *string `pulumi:"sapSid"`
}

type DatabaseConfiguration struct {
	DatabaseType                *string                     `pulumi:"databaseType"`
	InstanceCount               float64                     `pulumi:"instanceCount"`
	SubnetId                    string                      `pulumi:"subnetId"`
	VirtualMachineConfiguration VirtualMachineConfiguration `pulumi:"virtualMachineConfiguration"`
}

type DatabaseConfigurationResponse struct {
	DatabaseType                *string                             `pulumi:"databaseType"`
	InstanceCount               float64                             `pulumi:"instanceCount"`
	SubnetId                    string                              `pulumi:"subnetId"`
	VirtualMachineConfiguration VirtualMachineConfigurationResponse `pulumi:"virtualMachineConfiguration"`
}

type DatabaseProfile struct {
	BackupRetentionDays   *int         `pulumi:"backupRetentionDays"`
	HaEnabled             *string      `pulumi:"haEnabled"`
	ServerName            *string      `pulumi:"serverName"`
	Sku                   string       `pulumi:"sku"`
	SslEnforcementEnabled *string      `pulumi:"sslEnforcementEnabled"`
	StorageInGB           *float64     `pulumi:"storageInGB"`
	StorageIops           *float64     `pulumi:"storageIops"`
	StorageSku            *string      `pulumi:"storageSku"`
	Tier                  DatabaseTier `pulumi:"tier"`
	Type                  string       `pulumi:"type"`
	Version               *string      `pulumi:"version"`
}





type DatabaseProfileInput interface {
	pulumi.Input

	ToDatabaseProfileOutput() DatabaseProfileOutput
	ToDatabaseProfileOutputWithContext(context.Context) DatabaseProfileOutput
}

type DatabaseProfileArgs struct {
	BackupRetentionDays   pulumi.IntPtrInput     `pulumi:"backupRetentionDays"`
	HaEnabled             pulumi.StringPtrInput  `pulumi:"haEnabled"`
	ServerName            pulumi.StringPtrInput  `pulumi:"serverName"`
	Sku                   pulumi.StringInput     `pulumi:"sku"`
	SslEnforcementEnabled pulumi.StringPtrInput  `pulumi:"sslEnforcementEnabled"`
	StorageInGB           pulumi.Float64PtrInput `pulumi:"storageInGB"`
	StorageIops           pulumi.Float64PtrInput `pulumi:"storageIops"`
	StorageSku            pulumi.StringPtrInput  `pulumi:"storageSku"`
	Tier                  DatabaseTierInput      `pulumi:"tier"`
	Type                  pulumi.StringInput     `pulumi:"type"`
	Version               pulumi.StringPtrInput  `pulumi:"version"`
}

func (DatabaseProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseProfile)(nil)).Elem()
}

func (i DatabaseProfileArgs) ToDatabaseProfileOutput() DatabaseProfileOutput {
	return i.ToDatabaseProfileOutputWithContext(context.Background())
}

func (i DatabaseProfileArgs) ToDatabaseProfileOutputWithContext(ctx context.Context) DatabaseProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseProfileOutput)
}

type DatabaseProfileOutput struct{ *pulumi.OutputState }

func (DatabaseProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseProfile)(nil)).Elem()
}

func (o DatabaseProfileOutput) ToDatabaseProfileOutput() DatabaseProfileOutput {
	return o
}

func (o DatabaseProfileOutput) ToDatabaseProfileOutputWithContext(ctx context.Context) DatabaseProfileOutput {
	return o
}

func (o DatabaseProfileOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DatabaseProfile) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o DatabaseProfileOutput) HaEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseProfile) *string { return v.HaEnabled }).(pulumi.StringPtrOutput)
}

func (o DatabaseProfileOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseProfile) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o DatabaseProfileOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseProfile) string { return v.Sku }).(pulumi.StringOutput)
}

func (o DatabaseProfileOutput) SslEnforcementEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseProfile) *string { return v.SslEnforcementEnabled }).(pulumi.StringPtrOutput)
}

func (o DatabaseProfileOutput) StorageInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DatabaseProfile) *float64 { return v.StorageInGB }).(pulumi.Float64PtrOutput)
}

func (o DatabaseProfileOutput) StorageIops() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DatabaseProfile) *float64 { return v.StorageIops }).(pulumi.Float64PtrOutput)
}

func (o DatabaseProfileOutput) StorageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseProfile) *string { return v.StorageSku }).(pulumi.StringPtrOutput)
}

func (o DatabaseProfileOutput) Tier() DatabaseTierOutput {
	return o.ApplyT(func(v DatabaseProfile) DatabaseTier { return v.Tier }).(DatabaseTierOutput)
}

func (o DatabaseProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseProfile) string { return v.Type }).(pulumi.StringOutput)
}

func (o DatabaseProfileOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseProfile) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type DatabaseProfileResponse struct {
	BackupRetentionDays   *int     `pulumi:"backupRetentionDays"`
	HaEnabled             *string  `pulumi:"haEnabled"`
	ServerName            *string  `pulumi:"serverName"`
	ServerResourceId      string   `pulumi:"serverResourceId"`
	Sku                   string   `pulumi:"sku"`
	SslEnforcementEnabled *string  `pulumi:"sslEnforcementEnabled"`
	StorageInGB           *float64 `pulumi:"storageInGB"`
	StorageIops           *float64 `pulumi:"storageIops"`
	StorageSku            *string  `pulumi:"storageSku"`
	Tier                  string   `pulumi:"tier"`
	Type                  string   `pulumi:"type"`
	Version               *string  `pulumi:"version"`
}

type DatabaseProfileResponseOutput struct{ *pulumi.OutputState }

func (DatabaseProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseProfileResponse)(nil)).Elem()
}

func (o DatabaseProfileResponseOutput) ToDatabaseProfileResponseOutput() DatabaseProfileResponseOutput {
	return o
}

func (o DatabaseProfileResponseOutput) ToDatabaseProfileResponseOutputWithContext(ctx context.Context) DatabaseProfileResponseOutput {
	return o
}

func (o DatabaseProfileResponseOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DatabaseProfileResponse) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o DatabaseProfileResponseOutput) HaEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseProfileResponse) *string { return v.HaEnabled }).(pulumi.StringPtrOutput)
}

func (o DatabaseProfileResponseOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseProfileResponse) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o DatabaseProfileResponseOutput) ServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseProfileResponse) string { return v.ServerResourceId }).(pulumi.StringOutput)
}

func (o DatabaseProfileResponseOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseProfileResponse) string { return v.Sku }).(pulumi.StringOutput)
}

func (o DatabaseProfileResponseOutput) SslEnforcementEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseProfileResponse) *string { return v.SslEnforcementEnabled }).(pulumi.StringPtrOutput)
}

func (o DatabaseProfileResponseOutput) StorageInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DatabaseProfileResponse) *float64 { return v.StorageInGB }).(pulumi.Float64PtrOutput)
}

func (o DatabaseProfileResponseOutput) StorageIops() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DatabaseProfileResponse) *float64 { return v.StorageIops }).(pulumi.Float64PtrOutput)
}

func (o DatabaseProfileResponseOutput) StorageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseProfileResponse) *string { return v.StorageSku }).(pulumi.StringPtrOutput)
}

func (o DatabaseProfileResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseProfileResponse) string { return v.Tier }).(pulumi.StringOutput)
}

func (o DatabaseProfileResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseProfileResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o DatabaseProfileResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseProfileResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type DatabaseVmDetailsResponse struct {
	Status           string `pulumi:"status"`
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

type DatabaseVmDetailsResponseOutput struct{ *pulumi.OutputState }

func (DatabaseVmDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseVmDetailsResponse)(nil)).Elem()
}

func (o DatabaseVmDetailsResponseOutput) ToDatabaseVmDetailsResponseOutput() DatabaseVmDetailsResponseOutput {
	return o
}

func (o DatabaseVmDetailsResponseOutput) ToDatabaseVmDetailsResponseOutputWithContext(ctx context.Context) DatabaseVmDetailsResponseOutput {
	return o
}

func (o DatabaseVmDetailsResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseVmDetailsResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o DatabaseVmDetailsResponseOutput) VirtualMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseVmDetailsResponse) string { return v.VirtualMachineId }).(pulumi.StringOutput)
}

type DatabaseVmDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseVmDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseVmDetailsResponse)(nil)).Elem()
}

func (o DatabaseVmDetailsResponseArrayOutput) ToDatabaseVmDetailsResponseArrayOutput() DatabaseVmDetailsResponseArrayOutput {
	return o
}

func (o DatabaseVmDetailsResponseArrayOutput) ToDatabaseVmDetailsResponseArrayOutputWithContext(ctx context.Context) DatabaseVmDetailsResponseArrayOutput {
	return o
}

func (o DatabaseVmDetailsResponseArrayOutput) Index(i pulumi.IntInput) DatabaseVmDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseVmDetailsResponse {
		return vs[0].([]DatabaseVmDetailsResponse)[vs[1].(int)]
	}).(DatabaseVmDetailsResponseOutput)
}

type DeployerVmPackages struct {
	StorageAccountId *string `pulumi:"storageAccountId"`
	Url              *string `pulumi:"url"`
}

type DeployerVmPackagesResponse struct {
	StorageAccountId *string `pulumi:"storageAccountId"`
	Url              *string `pulumi:"url"`
}

type DeploymentConfiguration struct {
	AppLocation                 *string     `pulumi:"appLocation"`
	ConfigurationType           string      `pulumi:"configurationType"`
	InfrastructureConfiguration interface{} `pulumi:"infrastructureConfiguration"`
	SoftwareConfiguration       interface{} `pulumi:"softwareConfiguration"`
}

type DeploymentConfigurationResponse struct {
	AppLocation                 *string     `pulumi:"appLocation"`
	ConfigurationType           string      `pulumi:"configurationType"`
	InfrastructureConfiguration interface{} `pulumi:"infrastructureConfiguration"`
	SoftwareConfiguration       interface{} `pulumi:"softwareConfiguration"`
}

type DeploymentWithOSConfiguration struct {
	AppLocation                 *string             `pulumi:"appLocation"`
	ConfigurationType           string              `pulumi:"configurationType"`
	InfrastructureConfiguration interface{}         `pulumi:"infrastructureConfiguration"`
	OsSapConfiguration          *OsSapConfiguration `pulumi:"osSapConfiguration"`
	SoftwareConfiguration       interface{}         `pulumi:"softwareConfiguration"`
}

type DeploymentWithOSConfigurationResponse struct {
	AppLocation                 *string                     `pulumi:"appLocation"`
	ConfigurationType           string                      `pulumi:"configurationType"`
	InfrastructureConfiguration interface{}                 `pulumi:"infrastructureConfiguration"`
	OsSapConfiguration          *OsSapConfigurationResponse `pulumi:"osSapConfiguration"`
	SoftwareConfiguration       interface{}                 `pulumi:"softwareConfiguration"`
}

type DiscoveryConfiguration struct {
	CentralServerVmId *string `pulumi:"centralServerVmId"`
	ConfigurationType string  `pulumi:"configurationType"`
}

type DiscoveryConfigurationResponse struct {
	AppLocation       string  `pulumi:"appLocation"`
	CentralServerVmId *string `pulumi:"centralServerVmId"`
	ConfigurationType string  `pulumi:"configurationType"`
}

type DiskInfo struct {
	SizeInGB    *float64        `pulumi:"sizeInGB"`
	StorageType DiskStorageType `pulumi:"storageType"`
}





type DiskInfoInput interface {
	pulumi.Input

	ToDiskInfoOutput() DiskInfoOutput
	ToDiskInfoOutputWithContext(context.Context) DiskInfoOutput
}

type DiskInfoArgs struct {
	SizeInGB    pulumi.Float64PtrInput `pulumi:"sizeInGB"`
	StorageType DiskStorageTypeInput   `pulumi:"storageType"`
}

func (DiskInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskInfo)(nil)).Elem()
}

func (i DiskInfoArgs) ToDiskInfoOutput() DiskInfoOutput {
	return i.ToDiskInfoOutputWithContext(context.Background())
}

func (i DiskInfoArgs) ToDiskInfoOutputWithContext(ctx context.Context) DiskInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskInfoOutput)
}

func (i DiskInfoArgs) ToDiskInfoPtrOutput() DiskInfoPtrOutput {
	return i.ToDiskInfoPtrOutputWithContext(context.Background())
}

func (i DiskInfoArgs) ToDiskInfoPtrOutputWithContext(ctx context.Context) DiskInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskInfoOutput).ToDiskInfoPtrOutputWithContext(ctx)
}









type DiskInfoPtrInput interface {
	pulumi.Input

	ToDiskInfoPtrOutput() DiskInfoPtrOutput
	ToDiskInfoPtrOutputWithContext(context.Context) DiskInfoPtrOutput
}

type diskInfoPtrType DiskInfoArgs

func DiskInfoPtr(v *DiskInfoArgs) DiskInfoPtrInput {
	return (*diskInfoPtrType)(v)
}

func (*diskInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskInfo)(nil)).Elem()
}

func (i *diskInfoPtrType) ToDiskInfoPtrOutput() DiskInfoPtrOutput {
	return i.ToDiskInfoPtrOutputWithContext(context.Background())
}

func (i *diskInfoPtrType) ToDiskInfoPtrOutputWithContext(ctx context.Context) DiskInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskInfoPtrOutput)
}





type DiskInfoArrayInput interface {
	pulumi.Input

	ToDiskInfoArrayOutput() DiskInfoArrayOutput
	ToDiskInfoArrayOutputWithContext(context.Context) DiskInfoArrayOutput
}

type DiskInfoArray []DiskInfoInput

func (DiskInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskInfo)(nil)).Elem()
}

func (i DiskInfoArray) ToDiskInfoArrayOutput() DiskInfoArrayOutput {
	return i.ToDiskInfoArrayOutputWithContext(context.Background())
}

func (i DiskInfoArray) ToDiskInfoArrayOutputWithContext(ctx context.Context) DiskInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskInfoArrayOutput)
}

type DiskInfoOutput struct{ *pulumi.OutputState }

func (DiskInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskInfo)(nil)).Elem()
}

func (o DiskInfoOutput) ToDiskInfoOutput() DiskInfoOutput {
	return o
}

func (o DiskInfoOutput) ToDiskInfoOutputWithContext(ctx context.Context) DiskInfoOutput {
	return o
}

func (o DiskInfoOutput) ToDiskInfoPtrOutput() DiskInfoPtrOutput {
	return o.ToDiskInfoPtrOutputWithContext(context.Background())
}

func (o DiskInfoOutput) ToDiskInfoPtrOutputWithContext(ctx context.Context) DiskInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskInfo) *DiskInfo {
		return &v
	}).(DiskInfoPtrOutput)
}

func (o DiskInfoOutput) SizeInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DiskInfo) *float64 { return v.SizeInGB }).(pulumi.Float64PtrOutput)
}

func (o DiskInfoOutput) StorageType() DiskStorageTypeOutput {
	return o.ApplyT(func(v DiskInfo) DiskStorageType { return v.StorageType }).(DiskStorageTypeOutput)
}

type DiskInfoPtrOutput struct{ *pulumi.OutputState }

func (DiskInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskInfo)(nil)).Elem()
}

func (o DiskInfoPtrOutput) ToDiskInfoPtrOutput() DiskInfoPtrOutput {
	return o
}

func (o DiskInfoPtrOutput) ToDiskInfoPtrOutputWithContext(ctx context.Context) DiskInfoPtrOutput {
	return o
}

func (o DiskInfoPtrOutput) Elem() DiskInfoOutput {
	return o.ApplyT(func(v *DiskInfo) DiskInfo {
		if v != nil {
			return *v
		}
		var ret DiskInfo
		return ret
	}).(DiskInfoOutput)
}

func (o DiskInfoPtrOutput) SizeInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DiskInfo) *float64 {
		if v == nil {
			return nil
		}
		return v.SizeInGB
	}).(pulumi.Float64PtrOutput)
}

func (o DiskInfoPtrOutput) StorageType() DiskStorageTypePtrOutput {
	return o.ApplyT(func(v *DiskInfo) *DiskStorageType {
		if v == nil {
			return nil
		}
		return &v.StorageType
	}).(DiskStorageTypePtrOutput)
}

type DiskInfoArrayOutput struct{ *pulumi.OutputState }

func (DiskInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskInfo)(nil)).Elem()
}

func (o DiskInfoArrayOutput) ToDiskInfoArrayOutput() DiskInfoArrayOutput {
	return o
}

func (o DiskInfoArrayOutput) ToDiskInfoArrayOutputWithContext(ctx context.Context) DiskInfoArrayOutput {
	return o
}

func (o DiskInfoArrayOutput) Index(i pulumi.IntInput) DiskInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskInfo {
		return vs[0].([]DiskInfo)[vs[1].(int)]
	}).(DiskInfoOutput)
}

type DiskInfoResponse struct {
	SizeInGB    *float64 `pulumi:"sizeInGB"`
	StorageType string   `pulumi:"storageType"`
}

type DiskInfoResponseOutput struct{ *pulumi.OutputState }

func (DiskInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskInfoResponse)(nil)).Elem()
}

func (o DiskInfoResponseOutput) ToDiskInfoResponseOutput() DiskInfoResponseOutput {
	return o
}

func (o DiskInfoResponseOutput) ToDiskInfoResponseOutputWithContext(ctx context.Context) DiskInfoResponseOutput {
	return o
}

func (o DiskInfoResponseOutput) SizeInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DiskInfoResponse) *float64 { return v.SizeInGB }).(pulumi.Float64PtrOutput)
}

func (o DiskInfoResponseOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v DiskInfoResponse) string { return v.StorageType }).(pulumi.StringOutput)
}

type DiskInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskInfoResponse)(nil)).Elem()
}

func (o DiskInfoResponsePtrOutput) ToDiskInfoResponsePtrOutput() DiskInfoResponsePtrOutput {
	return o
}

func (o DiskInfoResponsePtrOutput) ToDiskInfoResponsePtrOutputWithContext(ctx context.Context) DiskInfoResponsePtrOutput {
	return o
}

func (o DiskInfoResponsePtrOutput) Elem() DiskInfoResponseOutput {
	return o.ApplyT(func(v *DiskInfoResponse) DiskInfoResponse {
		if v != nil {
			return *v
		}
		var ret DiskInfoResponse
		return ret
	}).(DiskInfoResponseOutput)
}

func (o DiskInfoResponsePtrOutput) SizeInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DiskInfoResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SizeInGB
	}).(pulumi.Float64PtrOutput)
}

func (o DiskInfoResponsePtrOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageType
	}).(pulumi.StringPtrOutput)
}

type DiskInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskInfoResponse)(nil)).Elem()
}

func (o DiskInfoResponseArrayOutput) ToDiskInfoResponseArrayOutput() DiskInfoResponseArrayOutput {
	return o
}

func (o DiskInfoResponseArrayOutput) ToDiskInfoResponseArrayOutputWithContext(ctx context.Context) DiskInfoResponseArrayOutput {
	return o
}

func (o DiskInfoResponseArrayOutput) Index(i pulumi.IntInput) DiskInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskInfoResponse {
		return vs[0].([]DiskInfoResponse)[vs[1].(int)]
	}).(DiskInfoResponseOutput)
}

type EnqueueReplicationServerPropertiesResponse struct {
	ErsVersion    string `pulumi:"ersVersion"`
	Health        string `pulumi:"health"`
	Hostname      string `pulumi:"hostname"`
	InstanceNo    string `pulumi:"instanceNo"`
	IpAddress     string `pulumi:"ipAddress"`
	KernelPatch   string `pulumi:"kernelPatch"`
	KernelVersion string `pulumi:"kernelVersion"`
}

type EnqueueReplicationServerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EnqueueReplicationServerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnqueueReplicationServerPropertiesResponse)(nil)).Elem()
}

func (o EnqueueReplicationServerPropertiesResponseOutput) ToEnqueueReplicationServerPropertiesResponseOutput() EnqueueReplicationServerPropertiesResponseOutput {
	return o
}

func (o EnqueueReplicationServerPropertiesResponseOutput) ToEnqueueReplicationServerPropertiesResponseOutputWithContext(ctx context.Context) EnqueueReplicationServerPropertiesResponseOutput {
	return o
}

func (o EnqueueReplicationServerPropertiesResponseOutput) ErsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v EnqueueReplicationServerPropertiesResponse) string { return v.ErsVersion }).(pulumi.StringOutput)
}

func (o EnqueueReplicationServerPropertiesResponseOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v EnqueueReplicationServerPropertiesResponse) string { return v.Health }).(pulumi.StringOutput)
}

func (o EnqueueReplicationServerPropertiesResponseOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v EnqueueReplicationServerPropertiesResponse) string { return v.Hostname }).(pulumi.StringOutput)
}

func (o EnqueueReplicationServerPropertiesResponseOutput) InstanceNo() pulumi.StringOutput {
	return o.ApplyT(func(v EnqueueReplicationServerPropertiesResponse) string { return v.InstanceNo }).(pulumi.StringOutput)
}

func (o EnqueueReplicationServerPropertiesResponseOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v EnqueueReplicationServerPropertiesResponse) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o EnqueueReplicationServerPropertiesResponseOutput) KernelPatch() pulumi.StringOutput {
	return o.ApplyT(func(v EnqueueReplicationServerPropertiesResponse) string { return v.KernelPatch }).(pulumi.StringOutput)
}

func (o EnqueueReplicationServerPropertiesResponseOutput) KernelVersion() pulumi.StringOutput {
	return o.ApplyT(func(v EnqueueReplicationServerPropertiesResponse) string { return v.KernelVersion }).(pulumi.StringOutput)
}

type EnqueueReplicationServerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EnqueueReplicationServerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnqueueReplicationServerPropertiesResponse)(nil)).Elem()
}

func (o EnqueueReplicationServerPropertiesResponsePtrOutput) ToEnqueueReplicationServerPropertiesResponsePtrOutput() EnqueueReplicationServerPropertiesResponsePtrOutput {
	return o
}

func (o EnqueueReplicationServerPropertiesResponsePtrOutput) ToEnqueueReplicationServerPropertiesResponsePtrOutputWithContext(ctx context.Context) EnqueueReplicationServerPropertiesResponsePtrOutput {
	return o
}

func (o EnqueueReplicationServerPropertiesResponsePtrOutput) Elem() EnqueueReplicationServerPropertiesResponseOutput {
	return o.ApplyT(func(v *EnqueueReplicationServerPropertiesResponse) EnqueueReplicationServerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EnqueueReplicationServerPropertiesResponse
		return ret
	}).(EnqueueReplicationServerPropertiesResponseOutput)
}

func (o EnqueueReplicationServerPropertiesResponsePtrOutput) ErsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnqueueReplicationServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErsVersion
	}).(pulumi.StringPtrOutput)
}

func (o EnqueueReplicationServerPropertiesResponsePtrOutput) Health() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnqueueReplicationServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Health
	}).(pulumi.StringPtrOutput)
}

func (o EnqueueReplicationServerPropertiesResponsePtrOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnqueueReplicationServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Hostname
	}).(pulumi.StringPtrOutput)
}

func (o EnqueueReplicationServerPropertiesResponsePtrOutput) InstanceNo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnqueueReplicationServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.InstanceNo
	}).(pulumi.StringPtrOutput)
}

func (o EnqueueReplicationServerPropertiesResponsePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnqueueReplicationServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o EnqueueReplicationServerPropertiesResponsePtrOutput) KernelPatch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnqueueReplicationServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KernelPatch
	}).(pulumi.StringPtrOutput)
}

func (o EnqueueReplicationServerPropertiesResponsePtrOutput) KernelVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnqueueReplicationServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KernelVersion
	}).(pulumi.StringPtrOutput)
}

type EnqueueServerPropertiesResponse struct {
	Health    string  `pulumi:"health"`
	Hostname  string  `pulumi:"hostname"`
	IpAddress string  `pulumi:"ipAddress"`
	Port      float64 `pulumi:"port"`
}

type EnqueueServerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EnqueueServerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnqueueServerPropertiesResponse)(nil)).Elem()
}

func (o EnqueueServerPropertiesResponseOutput) ToEnqueueServerPropertiesResponseOutput() EnqueueServerPropertiesResponseOutput {
	return o
}

func (o EnqueueServerPropertiesResponseOutput) ToEnqueueServerPropertiesResponseOutputWithContext(ctx context.Context) EnqueueServerPropertiesResponseOutput {
	return o
}

func (o EnqueueServerPropertiesResponseOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v EnqueueServerPropertiesResponse) string { return v.Health }).(pulumi.StringOutput)
}

func (o EnqueueServerPropertiesResponseOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v EnqueueServerPropertiesResponse) string { return v.Hostname }).(pulumi.StringOutput)
}

func (o EnqueueServerPropertiesResponseOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v EnqueueServerPropertiesResponse) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o EnqueueServerPropertiesResponseOutput) Port() pulumi.Float64Output {
	return o.ApplyT(func(v EnqueueServerPropertiesResponse) float64 { return v.Port }).(pulumi.Float64Output)
}

type EnqueueServerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EnqueueServerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnqueueServerPropertiesResponse)(nil)).Elem()
}

func (o EnqueueServerPropertiesResponsePtrOutput) ToEnqueueServerPropertiesResponsePtrOutput() EnqueueServerPropertiesResponsePtrOutput {
	return o
}

func (o EnqueueServerPropertiesResponsePtrOutput) ToEnqueueServerPropertiesResponsePtrOutputWithContext(ctx context.Context) EnqueueServerPropertiesResponsePtrOutput {
	return o
}

func (o EnqueueServerPropertiesResponsePtrOutput) Elem() EnqueueServerPropertiesResponseOutput {
	return o.ApplyT(func(v *EnqueueServerPropertiesResponse) EnqueueServerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EnqueueServerPropertiesResponse
		return ret
	}).(EnqueueServerPropertiesResponseOutput)
}

func (o EnqueueServerPropertiesResponsePtrOutput) Health() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnqueueServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Health
	}).(pulumi.StringPtrOutput)
}

func (o EnqueueServerPropertiesResponsePtrOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnqueueServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Hostname
	}).(pulumi.StringPtrOutput)
}

func (o EnqueueServerPropertiesResponsePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnqueueServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o EnqueueServerPropertiesResponsePtrOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EnqueueServerPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.Port
	}).(pulumi.Float64PtrOutput)
}

type ErrorDefinitionResponse struct {
	Code    string                    `pulumi:"code"`
	Details []ErrorDefinitionResponse `pulumi:"details"`
	Message string                    `pulumi:"message"`
}

type ErrorDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ErrorDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDefinitionResponse)(nil)).Elem()
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponseOutput() ErrorDefinitionResponseOutput {
	return o
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponseOutputWithContext(ctx context.Context) ErrorDefinitionResponseOutput {
	return o
}

func (o ErrorDefinitionResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDefinitionResponseOutput) Details() ErrorDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) []ErrorDefinitionResponse { return v.Details }).(ErrorDefinitionResponseArrayOutput)
}

func (o ErrorDefinitionResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Message }).(pulumi.StringOutput)
}

type ErrorDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDefinitionResponse)(nil)).Elem()
}

func (o ErrorDefinitionResponsePtrOutput) ToErrorDefinitionResponsePtrOutput() ErrorDefinitionResponsePtrOutput {
	return o
}

func (o ErrorDefinitionResponsePtrOutput) ToErrorDefinitionResponsePtrOutputWithContext(ctx context.Context) ErrorDefinitionResponsePtrOutput {
	return o
}

func (o ErrorDefinitionResponsePtrOutput) Elem() ErrorDefinitionResponseOutput {
	return o.ApplyT(func(v *ErrorDefinitionResponse) ErrorDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret ErrorDefinitionResponse
		return ret
	}).(ErrorDefinitionResponseOutput)
}

func (o ErrorDefinitionResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDefinitionResponsePtrOutput) Details() ErrorDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *ErrorDefinitionResponse) []ErrorDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(ErrorDefinitionResponseArrayOutput)
}

func (o ErrorDefinitionResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

type ErrorDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDefinitionResponse)(nil)).Elem()
}

func (o ErrorDefinitionResponseArrayOutput) ToErrorDefinitionResponseArrayOutput() ErrorDefinitionResponseArrayOutput {
	return o
}

func (o ErrorDefinitionResponseArrayOutput) ToErrorDefinitionResponseArrayOutputWithContext(ctx context.Context) ErrorDefinitionResponseArrayOutput {
	return o
}

func (o ErrorDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ErrorDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDefinitionResponse {
		return vs[0].([]ErrorDefinitionResponse)[vs[1].(int)]
	}).(ErrorDefinitionResponseOutput)
}

type ErrorResponse struct {
	Code       string                  `pulumi:"code"`
	Details    []ErrorResponse         `pulumi:"details"`
	InnerError ErrorResponseInnerError `pulumi:"innerError"`
	Message    string                  `pulumi:"message"`
	Target     string                  `pulumi:"target"`
}

type ErrorResponseOutput struct{ *pulumi.OutputState }

func (ErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponse)(nil)).Elem()
}

func (o ErrorResponseOutput) ToErrorResponseOutput() ErrorResponseOutput {
	return o
}

func (o ErrorResponseOutput) ToErrorResponseOutputWithContext(ctx context.Context) ErrorResponseOutput {
	return o
}

func (o ErrorResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorResponseOutput) Details() ErrorResponseArrayOutput {
	return o.ApplyT(func(v ErrorResponse) []ErrorResponse { return v.Details }).(ErrorResponseArrayOutput)
}

func (o ErrorResponseOutput) InnerError() ErrorResponseInnerErrorOutput {
	return o.ApplyT(func(v ErrorResponse) ErrorResponseInnerError { return v.InnerError }).(ErrorResponseInnerErrorOutput)
}

func (o ErrorResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponse)(nil)).Elem()
}

func (o ErrorResponsePtrOutput) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return o
}

func (o ErrorResponsePtrOutput) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return o
}

func (o ErrorResponsePtrOutput) Elem() ErrorResponseOutput {
	return o.ApplyT(func(v *ErrorResponse) ErrorResponse {
		if v != nil {
			return *v
		}
		var ret ErrorResponse
		return ret
	}).(ErrorResponseOutput)
}

func (o ErrorResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorResponsePtrOutput) Details() ErrorResponseArrayOutput {
	return o.ApplyT(func(v *ErrorResponse) []ErrorResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(ErrorResponseArrayOutput)
}

func (o ErrorResponsePtrOutput) InnerError() ErrorResponseInnerErrorPtrOutput {
	return o.ApplyT(func(v *ErrorResponse) *ErrorResponseInnerError {
		if v == nil {
			return nil
		}
		return &v.InnerError
	}).(ErrorResponseInnerErrorPtrOutput)
}

func (o ErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ErrorResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Target
	}).(pulumi.StringPtrOutput)
}

type ErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorResponse)(nil)).Elem()
}

func (o ErrorResponseArrayOutput) ToErrorResponseArrayOutput() ErrorResponseArrayOutput {
	return o
}

func (o ErrorResponseArrayOutput) ToErrorResponseArrayOutputWithContext(ctx context.Context) ErrorResponseArrayOutput {
	return o
}

func (o ErrorResponseArrayOutput) Index(i pulumi.IntInput) ErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorResponse {
		return vs[0].([]ErrorResponse)[vs[1].(int)]
	}).(ErrorResponseOutput)
}

type ErrorResponseInnerError struct {
	InnerError *ErrorResponse `pulumi:"innerError"`
}

type ErrorResponseInnerErrorOutput struct{ *pulumi.OutputState }

func (ErrorResponseInnerErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponseInnerError)(nil)).Elem()
}

func (o ErrorResponseInnerErrorOutput) ToErrorResponseInnerErrorOutput() ErrorResponseInnerErrorOutput {
	return o
}

func (o ErrorResponseInnerErrorOutput) ToErrorResponseInnerErrorOutputWithContext(ctx context.Context) ErrorResponseInnerErrorOutput {
	return o
}

func (o ErrorResponseInnerErrorOutput) InnerError() ErrorResponsePtrOutput {
	return o.ApplyT(func(v ErrorResponseInnerError) *ErrorResponse { return v.InnerError }).(ErrorResponsePtrOutput)
}

type ErrorResponseInnerErrorPtrOutput struct{ *pulumi.OutputState }

func (ErrorResponseInnerErrorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponseInnerError)(nil)).Elem()
}

func (o ErrorResponseInnerErrorPtrOutput) ToErrorResponseInnerErrorPtrOutput() ErrorResponseInnerErrorPtrOutput {
	return o
}

func (o ErrorResponseInnerErrorPtrOutput) ToErrorResponseInnerErrorPtrOutputWithContext(ctx context.Context) ErrorResponseInnerErrorPtrOutput {
	return o
}

func (o ErrorResponseInnerErrorPtrOutput) Elem() ErrorResponseInnerErrorOutput {
	return o.ApplyT(func(v *ErrorResponseInnerError) ErrorResponseInnerError {
		if v != nil {
			return *v
		}
		var ret ErrorResponseInnerError
		return ret
	}).(ErrorResponseInnerErrorOutput)
}

func (o ErrorResponseInnerErrorPtrOutput) InnerError() ErrorResponsePtrOutput {
	return o.ApplyT(func(v *ErrorResponseInnerError) *ErrorResponse {
		if v == nil {
			return nil
		}
		return v.InnerError
	}).(ErrorResponsePtrOutput)
}

type FileshareProfile struct {
	ShareSizeInGB *float64 `pulumi:"shareSizeInGB"`
	ShareType     string   `pulumi:"shareType"`
	StorageType   string   `pulumi:"storageType"`
}





type FileshareProfileInput interface {
	pulumi.Input

	ToFileshareProfileOutput() FileshareProfileOutput
	ToFileshareProfileOutputWithContext(context.Context) FileshareProfileOutput
}

type FileshareProfileArgs struct {
	ShareSizeInGB pulumi.Float64PtrInput `pulumi:"shareSizeInGB"`
	ShareType     pulumi.StringInput     `pulumi:"shareType"`
	StorageType   pulumi.StringInput     `pulumi:"storageType"`
}

func (FileshareProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileshareProfile)(nil)).Elem()
}

func (i FileshareProfileArgs) ToFileshareProfileOutput() FileshareProfileOutput {
	return i.ToFileshareProfileOutputWithContext(context.Background())
}

func (i FileshareProfileArgs) ToFileshareProfileOutputWithContext(ctx context.Context) FileshareProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileshareProfileOutput)
}

func (i FileshareProfileArgs) ToFileshareProfilePtrOutput() FileshareProfilePtrOutput {
	return i.ToFileshareProfilePtrOutputWithContext(context.Background())
}

func (i FileshareProfileArgs) ToFileshareProfilePtrOutputWithContext(ctx context.Context) FileshareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileshareProfileOutput).ToFileshareProfilePtrOutputWithContext(ctx)
}









type FileshareProfilePtrInput interface {
	pulumi.Input

	ToFileshareProfilePtrOutput() FileshareProfilePtrOutput
	ToFileshareProfilePtrOutputWithContext(context.Context) FileshareProfilePtrOutput
}

type fileshareProfilePtrType FileshareProfileArgs

func FileshareProfilePtr(v *FileshareProfileArgs) FileshareProfilePtrInput {
	return (*fileshareProfilePtrType)(v)
}

func (*fileshareProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileshareProfile)(nil)).Elem()
}

func (i *fileshareProfilePtrType) ToFileshareProfilePtrOutput() FileshareProfilePtrOutput {
	return i.ToFileshareProfilePtrOutputWithContext(context.Background())
}

func (i *fileshareProfilePtrType) ToFileshareProfilePtrOutputWithContext(ctx context.Context) FileshareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileshareProfilePtrOutput)
}

type FileshareProfileOutput struct{ *pulumi.OutputState }

func (FileshareProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileshareProfile)(nil)).Elem()
}

func (o FileshareProfileOutput) ToFileshareProfileOutput() FileshareProfileOutput {
	return o
}

func (o FileshareProfileOutput) ToFileshareProfileOutputWithContext(ctx context.Context) FileshareProfileOutput {
	return o
}

func (o FileshareProfileOutput) ToFileshareProfilePtrOutput() FileshareProfilePtrOutput {
	return o.ToFileshareProfilePtrOutputWithContext(context.Background())
}

func (o FileshareProfileOutput) ToFileshareProfilePtrOutputWithContext(ctx context.Context) FileshareProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileshareProfile) *FileshareProfile {
		return &v
	}).(FileshareProfilePtrOutput)
}

func (o FileshareProfileOutput) ShareSizeInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v FileshareProfile) *float64 { return v.ShareSizeInGB }).(pulumi.Float64PtrOutput)
}

func (o FileshareProfileOutput) ShareType() pulumi.StringOutput {
	return o.ApplyT(func(v FileshareProfile) string { return v.ShareType }).(pulumi.StringOutput)
}

func (o FileshareProfileOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v FileshareProfile) string { return v.StorageType }).(pulumi.StringOutput)
}

type FileshareProfilePtrOutput struct{ *pulumi.OutputState }

func (FileshareProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileshareProfile)(nil)).Elem()
}

func (o FileshareProfilePtrOutput) ToFileshareProfilePtrOutput() FileshareProfilePtrOutput {
	return o
}

func (o FileshareProfilePtrOutput) ToFileshareProfilePtrOutputWithContext(ctx context.Context) FileshareProfilePtrOutput {
	return o
}

func (o FileshareProfilePtrOutput) Elem() FileshareProfileOutput {
	return o.ApplyT(func(v *FileshareProfile) FileshareProfile {
		if v != nil {
			return *v
		}
		var ret FileshareProfile
		return ret
	}).(FileshareProfileOutput)
}

func (o FileshareProfilePtrOutput) ShareSizeInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *FileshareProfile) *float64 {
		if v == nil {
			return nil
		}
		return v.ShareSizeInGB
	}).(pulumi.Float64PtrOutput)
}

func (o FileshareProfilePtrOutput) ShareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileshareProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ShareType
	}).(pulumi.StringPtrOutput)
}

func (o FileshareProfilePtrOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileshareProfile) *string {
		if v == nil {
			return nil
		}
		return &v.StorageType
	}).(pulumi.StringPtrOutput)
}

type FileshareProfileResponse struct {
	ShareName         string   `pulumi:"shareName"`
	ShareSizeInGB     *float64 `pulumi:"shareSizeInGB"`
	ShareType         string   `pulumi:"shareType"`
	StorageResourceId string   `pulumi:"storageResourceId"`
	StorageType       string   `pulumi:"storageType"`
}

type FileshareProfileResponseOutput struct{ *pulumi.OutputState }

func (FileshareProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileshareProfileResponse)(nil)).Elem()
}

func (o FileshareProfileResponseOutput) ToFileshareProfileResponseOutput() FileshareProfileResponseOutput {
	return o
}

func (o FileshareProfileResponseOutput) ToFileshareProfileResponseOutputWithContext(ctx context.Context) FileshareProfileResponseOutput {
	return o
}

func (o FileshareProfileResponseOutput) ShareName() pulumi.StringOutput {
	return o.ApplyT(func(v FileshareProfileResponse) string { return v.ShareName }).(pulumi.StringOutput)
}

func (o FileshareProfileResponseOutput) ShareSizeInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v FileshareProfileResponse) *float64 { return v.ShareSizeInGB }).(pulumi.Float64PtrOutput)
}

func (o FileshareProfileResponseOutput) ShareType() pulumi.StringOutput {
	return o.ApplyT(func(v FileshareProfileResponse) string { return v.ShareType }).(pulumi.StringOutput)
}

func (o FileshareProfileResponseOutput) StorageResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v FileshareProfileResponse) string { return v.StorageResourceId }).(pulumi.StringOutput)
}

func (o FileshareProfileResponseOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v FileshareProfileResponse) string { return v.StorageType }).(pulumi.StringOutput)
}

type FileshareProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (FileshareProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileshareProfileResponse)(nil)).Elem()
}

func (o FileshareProfileResponsePtrOutput) ToFileshareProfileResponsePtrOutput() FileshareProfileResponsePtrOutput {
	return o
}

func (o FileshareProfileResponsePtrOutput) ToFileshareProfileResponsePtrOutputWithContext(ctx context.Context) FileshareProfileResponsePtrOutput {
	return o
}

func (o FileshareProfileResponsePtrOutput) Elem() FileshareProfileResponseOutput {
	return o.ApplyT(func(v *FileshareProfileResponse) FileshareProfileResponse {
		if v != nil {
			return *v
		}
		var ret FileshareProfileResponse
		return ret
	}).(FileshareProfileResponseOutput)
}

func (o FileshareProfileResponsePtrOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileshareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ShareName
	}).(pulumi.StringPtrOutput)
}

func (o FileshareProfileResponsePtrOutput) ShareSizeInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *FileshareProfileResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.ShareSizeInGB
	}).(pulumi.Float64PtrOutput)
}

func (o FileshareProfileResponsePtrOutput) ShareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileshareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ShareType
	}).(pulumi.StringPtrOutput)
}

func (o FileshareProfileResponsePtrOutput) StorageResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileshareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageResourceId
	}).(pulumi.StringPtrOutput)
}

func (o FileshareProfileResponsePtrOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileshareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageType
	}).(pulumi.StringPtrOutput)
}

type GatewayServerPropertiesResponse struct {
	Health string  `pulumi:"health"`
	Port   float64 `pulumi:"port"`
}

type GatewayServerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GatewayServerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayServerPropertiesResponse)(nil)).Elem()
}

func (o GatewayServerPropertiesResponseOutput) ToGatewayServerPropertiesResponseOutput() GatewayServerPropertiesResponseOutput {
	return o
}

func (o GatewayServerPropertiesResponseOutput) ToGatewayServerPropertiesResponseOutputWithContext(ctx context.Context) GatewayServerPropertiesResponseOutput {
	return o
}

func (o GatewayServerPropertiesResponseOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayServerPropertiesResponse) string { return v.Health }).(pulumi.StringOutput)
}

func (o GatewayServerPropertiesResponseOutput) Port() pulumi.Float64Output {
	return o.ApplyT(func(v GatewayServerPropertiesResponse) float64 { return v.Port }).(pulumi.Float64Output)
}

type GatewayServerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (GatewayServerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayServerPropertiesResponse)(nil)).Elem()
}

func (o GatewayServerPropertiesResponsePtrOutput) ToGatewayServerPropertiesResponsePtrOutput() GatewayServerPropertiesResponsePtrOutput {
	return o
}

func (o GatewayServerPropertiesResponsePtrOutput) ToGatewayServerPropertiesResponsePtrOutputWithContext(ctx context.Context) GatewayServerPropertiesResponsePtrOutput {
	return o
}

func (o GatewayServerPropertiesResponsePtrOutput) Elem() GatewayServerPropertiesResponseOutput {
	return o.ApplyT(func(v *GatewayServerPropertiesResponse) GatewayServerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret GatewayServerPropertiesResponse
		return ret
	}).(GatewayServerPropertiesResponseOutput)
}

func (o GatewayServerPropertiesResponsePtrOutput) Health() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Health
	}).(pulumi.StringPtrOutput)
}

func (o GatewayServerPropertiesResponsePtrOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GatewayServerPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.Port
	}).(pulumi.Float64PtrOutput)
}

type HanaDbProviderInstanceProperties struct {
	DbName                   *string `pulumi:"dbName"`
	DbPassword               *string `pulumi:"dbPassword"`
	DbPasswordUri            *string `pulumi:"dbPasswordUri"`
	DbSslCertificateUri      *string `pulumi:"dbSslCertificateUri"`
	DbUsername               *string `pulumi:"dbUsername"`
	Hostname                 *string `pulumi:"hostname"`
	InstanceNumber           *string `pulumi:"instanceNumber"`
	ProviderType             string  `pulumi:"providerType"`
	SqlPort                  *string `pulumi:"sqlPort"`
	SslHostNameInCertificate *string `pulumi:"sslHostNameInCertificate"`
}

type HanaDbProviderInstancePropertiesResponse struct {
	DbName                   *string `pulumi:"dbName"`
	DbPassword               *string `pulumi:"dbPassword"`
	DbPasswordUri            *string `pulumi:"dbPasswordUri"`
	DbSslCertificateUri      *string `pulumi:"dbSslCertificateUri"`
	DbUsername               *string `pulumi:"dbUsername"`
	Hostname                 *string `pulumi:"hostname"`
	InstanceNumber           *string `pulumi:"instanceNumber"`
	ProviderType             string  `pulumi:"providerType"`
	SqlPort                  *string `pulumi:"sqlPort"`
	SslHostNameInCertificate *string `pulumi:"sslHostNameInCertificate"`
}

type HighAvailabilityConfiguration struct {
	HighAvailabilityType string `pulumi:"highAvailabilityType"`
}

type HighAvailabilityConfigurationResponse struct {
	HighAvailabilityType string `pulumi:"highAvailabilityType"`
}

type HighAvailabilitySoftwareConfiguration struct {
	FencingClientId       string `pulumi:"fencingClientId"`
	FencingClientPassword string `pulumi:"fencingClientPassword"`
}

type HighAvailabilitySoftwareConfigurationResponse struct {
	FencingClientId       string `pulumi:"fencingClientId"`
	FencingClientPassword string `pulumi:"fencingClientPassword"`
}

type ImageReference struct {
	Offer                *string `pulumi:"offer"`
	Publisher            *string `pulumi:"publisher"`
	SharedGalleryImageId *string `pulumi:"sharedGalleryImageId"`
	Sku                  *string `pulumi:"sku"`
	Version              *string `pulumi:"version"`
}

type ImageReferenceResponse struct {
	ExactVersion         string  `pulumi:"exactVersion"`
	Offer                *string `pulumi:"offer"`
	Publisher            *string `pulumi:"publisher"`
	SharedGalleryImageId *string `pulumi:"sharedGalleryImageId"`
	Sku                  *string `pulumi:"sku"`
	Version              *string `pulumi:"version"`
}

type LinuxConfiguration struct {
	DisablePasswordAuthentication *bool             `pulumi:"disablePasswordAuthentication"`
	OsType                        string            `pulumi:"osType"`
	Ssh                           *SshConfiguration `pulumi:"ssh"`
	SshKeyPair                    *SshKeyPair       `pulumi:"sshKeyPair"`
}

type LinuxConfigurationResponse struct {
	DisablePasswordAuthentication *bool                     `pulumi:"disablePasswordAuthentication"`
	OsType                        string                    `pulumi:"osType"`
	Ssh                           *SshConfigurationResponse `pulumi:"ssh"`
	SshKeyPair                    *SshKeyPairResponse       `pulumi:"sshKeyPair"`
}

type ManagedRGConfiguration struct {
	Name *string `pulumi:"name"`
}





type ManagedRGConfigurationInput interface {
	pulumi.Input

	ToManagedRGConfigurationOutput() ManagedRGConfigurationOutput
	ToManagedRGConfigurationOutputWithContext(context.Context) ManagedRGConfigurationOutput
}

type ManagedRGConfigurationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ManagedRGConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRGConfiguration)(nil)).Elem()
}

func (i ManagedRGConfigurationArgs) ToManagedRGConfigurationOutput() ManagedRGConfigurationOutput {
	return i.ToManagedRGConfigurationOutputWithContext(context.Background())
}

func (i ManagedRGConfigurationArgs) ToManagedRGConfigurationOutputWithContext(ctx context.Context) ManagedRGConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRGConfigurationOutput)
}

func (i ManagedRGConfigurationArgs) ToManagedRGConfigurationPtrOutput() ManagedRGConfigurationPtrOutput {
	return i.ToManagedRGConfigurationPtrOutputWithContext(context.Background())
}

func (i ManagedRGConfigurationArgs) ToManagedRGConfigurationPtrOutputWithContext(ctx context.Context) ManagedRGConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRGConfigurationOutput).ToManagedRGConfigurationPtrOutputWithContext(ctx)
}









type ManagedRGConfigurationPtrInput interface {
	pulumi.Input

	ToManagedRGConfigurationPtrOutput() ManagedRGConfigurationPtrOutput
	ToManagedRGConfigurationPtrOutputWithContext(context.Context) ManagedRGConfigurationPtrOutput
}

type managedRGConfigurationPtrType ManagedRGConfigurationArgs

func ManagedRGConfigurationPtr(v *ManagedRGConfigurationArgs) ManagedRGConfigurationPtrInput {
	return (*managedRGConfigurationPtrType)(v)
}

func (*managedRGConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRGConfiguration)(nil)).Elem()
}

func (i *managedRGConfigurationPtrType) ToManagedRGConfigurationPtrOutput() ManagedRGConfigurationPtrOutput {
	return i.ToManagedRGConfigurationPtrOutputWithContext(context.Background())
}

func (i *managedRGConfigurationPtrType) ToManagedRGConfigurationPtrOutputWithContext(ctx context.Context) ManagedRGConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRGConfigurationPtrOutput)
}

type ManagedRGConfigurationOutput struct{ *pulumi.OutputState }

func (ManagedRGConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRGConfiguration)(nil)).Elem()
}

func (o ManagedRGConfigurationOutput) ToManagedRGConfigurationOutput() ManagedRGConfigurationOutput {
	return o
}

func (o ManagedRGConfigurationOutput) ToManagedRGConfigurationOutputWithContext(ctx context.Context) ManagedRGConfigurationOutput {
	return o
}

func (o ManagedRGConfigurationOutput) ToManagedRGConfigurationPtrOutput() ManagedRGConfigurationPtrOutput {
	return o.ToManagedRGConfigurationPtrOutputWithContext(context.Background())
}

func (o ManagedRGConfigurationOutput) ToManagedRGConfigurationPtrOutputWithContext(ctx context.Context) ManagedRGConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedRGConfiguration) *ManagedRGConfiguration {
		return &v
	}).(ManagedRGConfigurationPtrOutput)
}

func (o ManagedRGConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedRGConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ManagedRGConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ManagedRGConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRGConfiguration)(nil)).Elem()
}

func (o ManagedRGConfigurationPtrOutput) ToManagedRGConfigurationPtrOutput() ManagedRGConfigurationPtrOutput {
	return o
}

func (o ManagedRGConfigurationPtrOutput) ToManagedRGConfigurationPtrOutputWithContext(ctx context.Context) ManagedRGConfigurationPtrOutput {
	return o
}

func (o ManagedRGConfigurationPtrOutput) Elem() ManagedRGConfigurationOutput {
	return o.ApplyT(func(v *ManagedRGConfiguration) ManagedRGConfiguration {
		if v != nil {
			return *v
		}
		var ret ManagedRGConfiguration
		return ret
	}).(ManagedRGConfigurationOutput)
}

func (o ManagedRGConfigurationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedRGConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ManagedRGConfigurationResponse struct {
	Name *string `pulumi:"name"`
}

type ManagedRGConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ManagedRGConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRGConfigurationResponse)(nil)).Elem()
}

func (o ManagedRGConfigurationResponseOutput) ToManagedRGConfigurationResponseOutput() ManagedRGConfigurationResponseOutput {
	return o
}

func (o ManagedRGConfigurationResponseOutput) ToManagedRGConfigurationResponseOutputWithContext(ctx context.Context) ManagedRGConfigurationResponseOutput {
	return o
}

func (o ManagedRGConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedRGConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ManagedRGConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedRGConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRGConfigurationResponse)(nil)).Elem()
}

func (o ManagedRGConfigurationResponsePtrOutput) ToManagedRGConfigurationResponsePtrOutput() ManagedRGConfigurationResponsePtrOutput {
	return o
}

func (o ManagedRGConfigurationResponsePtrOutput) ToManagedRGConfigurationResponsePtrOutputWithContext(ctx context.Context) ManagedRGConfigurationResponsePtrOutput {
	return o
}

func (o ManagedRGConfigurationResponsePtrOutput) Elem() ManagedRGConfigurationResponseOutput {
	return o.ApplyT(func(v *ManagedRGConfigurationResponse) ManagedRGConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ManagedRGConfigurationResponse
		return ret
	}).(ManagedRGConfigurationResponseOutput)
}

func (o ManagedRGConfigurationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedRGConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type MessageServerPropertiesResponse struct {
	Health         string  `pulumi:"health"`
	Hostname       string  `pulumi:"hostname"`
	HttpPort       float64 `pulumi:"httpPort"`
	HttpsPort      float64 `pulumi:"httpsPort"`
	InternalMsPort float64 `pulumi:"internalMsPort"`
	IpAddress      string  `pulumi:"ipAddress"`
	MsPort         float64 `pulumi:"msPort"`
}

type MessageServerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MessageServerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageServerPropertiesResponse)(nil)).Elem()
}

func (o MessageServerPropertiesResponseOutput) ToMessageServerPropertiesResponseOutput() MessageServerPropertiesResponseOutput {
	return o
}

func (o MessageServerPropertiesResponseOutput) ToMessageServerPropertiesResponseOutputWithContext(ctx context.Context) MessageServerPropertiesResponseOutput {
	return o
}

func (o MessageServerPropertiesResponseOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v MessageServerPropertiesResponse) string { return v.Health }).(pulumi.StringOutput)
}

func (o MessageServerPropertiesResponseOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v MessageServerPropertiesResponse) string { return v.Hostname }).(pulumi.StringOutput)
}

func (o MessageServerPropertiesResponseOutput) HttpPort() pulumi.Float64Output {
	return o.ApplyT(func(v MessageServerPropertiesResponse) float64 { return v.HttpPort }).(pulumi.Float64Output)
}

func (o MessageServerPropertiesResponseOutput) HttpsPort() pulumi.Float64Output {
	return o.ApplyT(func(v MessageServerPropertiesResponse) float64 { return v.HttpsPort }).(pulumi.Float64Output)
}

func (o MessageServerPropertiesResponseOutput) InternalMsPort() pulumi.Float64Output {
	return o.ApplyT(func(v MessageServerPropertiesResponse) float64 { return v.InternalMsPort }).(pulumi.Float64Output)
}

func (o MessageServerPropertiesResponseOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v MessageServerPropertiesResponse) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o MessageServerPropertiesResponseOutput) MsPort() pulumi.Float64Output {
	return o.ApplyT(func(v MessageServerPropertiesResponse) float64 { return v.MsPort }).(pulumi.Float64Output)
}

type MessageServerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MessageServerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MessageServerPropertiesResponse)(nil)).Elem()
}

func (o MessageServerPropertiesResponsePtrOutput) ToMessageServerPropertiesResponsePtrOutput() MessageServerPropertiesResponsePtrOutput {
	return o
}

func (o MessageServerPropertiesResponsePtrOutput) ToMessageServerPropertiesResponsePtrOutputWithContext(ctx context.Context) MessageServerPropertiesResponsePtrOutput {
	return o
}

func (o MessageServerPropertiesResponsePtrOutput) Elem() MessageServerPropertiesResponseOutput {
	return o.ApplyT(func(v *MessageServerPropertiesResponse) MessageServerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MessageServerPropertiesResponse
		return ret
	}).(MessageServerPropertiesResponseOutput)
}

func (o MessageServerPropertiesResponsePtrOutput) Health() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MessageServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Health
	}).(pulumi.StringPtrOutput)
}

func (o MessageServerPropertiesResponsePtrOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MessageServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Hostname
	}).(pulumi.StringPtrOutput)
}

func (o MessageServerPropertiesResponsePtrOutput) HttpPort() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageServerPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.HttpPort
	}).(pulumi.Float64PtrOutput)
}

func (o MessageServerPropertiesResponsePtrOutput) HttpsPort() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageServerPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.HttpsPort
	}).(pulumi.Float64PtrOutput)
}

func (o MessageServerPropertiesResponsePtrOutput) InternalMsPort() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageServerPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.InternalMsPort
	}).(pulumi.Float64PtrOutput)
}

func (o MessageServerPropertiesResponsePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MessageServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o MessageServerPropertiesResponsePtrOutput) MsPort() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageServerPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.MsPort
	}).(pulumi.Float64PtrOutput)
}

type MonitorPropertiesResponseErrors struct {
	Code       string                  `pulumi:"code"`
	Details    []ErrorResponse         `pulumi:"details"`
	InnerError ErrorResponseInnerError `pulumi:"innerError"`
	Message    string                  `pulumi:"message"`
	Target     string                  `pulumi:"target"`
}

type MonitorPropertiesResponseErrorsOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesResponseErrorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorPropertiesResponseErrors)(nil)).Elem()
}

func (o MonitorPropertiesResponseErrorsOutput) ToMonitorPropertiesResponseErrorsOutput() MonitorPropertiesResponseErrorsOutput {
	return o
}

func (o MonitorPropertiesResponseErrorsOutput) ToMonitorPropertiesResponseErrorsOutputWithContext(ctx context.Context) MonitorPropertiesResponseErrorsOutput {
	return o
}

func (o MonitorPropertiesResponseErrorsOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v MonitorPropertiesResponseErrors) string { return v.Code }).(pulumi.StringOutput)
}

func (o MonitorPropertiesResponseErrorsOutput) Details() ErrorResponseArrayOutput {
	return o.ApplyT(func(v MonitorPropertiesResponseErrors) []ErrorResponse { return v.Details }).(ErrorResponseArrayOutput)
}

func (o MonitorPropertiesResponseErrorsOutput) InnerError() ErrorResponseInnerErrorOutput {
	return o.ApplyT(func(v MonitorPropertiesResponseErrors) ErrorResponseInnerError { return v.InnerError }).(ErrorResponseInnerErrorOutput)
}

func (o MonitorPropertiesResponseErrorsOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v MonitorPropertiesResponseErrors) string { return v.Message }).(pulumi.StringOutput)
}

func (o MonitorPropertiesResponseErrorsOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v MonitorPropertiesResponseErrors) string { return v.Target }).(pulumi.StringOutput)
}

type MsSqlServerProviderInstanceProperties struct {
	DbPassword    *string `pulumi:"dbPassword"`
	DbPasswordUri *string `pulumi:"dbPasswordUri"`
	DbPort        *string `pulumi:"dbPort"`
	DbUsername    *string `pulumi:"dbUsername"`
	Hostname      *string `pulumi:"hostname"`
	ProviderType  string  `pulumi:"providerType"`
	SapSid        *string `pulumi:"sapSid"`
}

type MsSqlServerProviderInstancePropertiesResponse struct {
	DbPassword    *string `pulumi:"dbPassword"`
	DbPasswordUri *string `pulumi:"dbPasswordUri"`
	DbPort        *string `pulumi:"dbPort"`
	DbUsername    *string `pulumi:"dbUsername"`
	Hostname      *string `pulumi:"hostname"`
	ProviderType  string  `pulumi:"providerType"`
	SapSid        *string `pulumi:"sapSid"`
}

type NetworkConfiguration struct {
	IsSecondaryIpEnabled *bool `pulumi:"isSecondaryIpEnabled"`
}


func (val *NetworkConfiguration) Defaults() *NetworkConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSecondaryIpEnabled) {
		isSecondaryIpEnabled_ := false
		tmp.IsSecondaryIpEnabled = &isSecondaryIpEnabled_
	}
	return &tmp
}

type NetworkConfigurationResponse struct {
	IsSecondaryIpEnabled *bool `pulumi:"isSecondaryIpEnabled"`
}


func (val *NetworkConfigurationResponse) Defaults() *NetworkConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSecondaryIpEnabled) {
		isSecondaryIpEnabled_ := false
		tmp.IsSecondaryIpEnabled = &isSecondaryIpEnabled_
	}
	return &tmp
}

type NetworkProfile struct {
	AzureFrontDoorEnabled *string `pulumi:"azureFrontDoorEnabled"`
	Capacity              *int    `pulumi:"capacity"`
	LoadBalancerSku       *string `pulumi:"loadBalancerSku"`
	LoadBalancerTier      *string `pulumi:"loadBalancerTier"`
	LoadBalancerType      string  `pulumi:"loadBalancerType"`
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	AzureFrontDoorEnabled pulumi.StringPtrInput `pulumi:"azureFrontDoorEnabled"`
	Capacity              pulumi.IntPtrInput    `pulumi:"capacity"`
	LoadBalancerSku       pulumi.StringPtrInput `pulumi:"loadBalancerSku"`
	LoadBalancerTier      pulumi.StringPtrInput `pulumi:"loadBalancerTier"`
	LoadBalancerType      pulumi.StringInput    `pulumi:"loadBalancerType"`
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArgs) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput).ToNetworkProfilePtrOutputWithContext(ctx)
}









type NetworkProfilePtrInput interface {
	pulumi.Input

	ToNetworkProfilePtrOutput() NetworkProfilePtrOutput
	ToNetworkProfilePtrOutputWithContext(context.Context) NetworkProfilePtrOutput
}

type networkProfilePtrType NetworkProfileArgs

func NetworkProfilePtr(v *NetworkProfileArgs) NetworkProfilePtrInput {
	return (*networkProfilePtrType)(v)
}

func (*networkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfilePtrOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfile) *NetworkProfile {
		return &v
	}).(NetworkProfilePtrOutput)
}

func (o NetworkProfileOutput) AzureFrontDoorEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.AzureFrontDoorEnabled }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o NetworkProfileOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.LoadBalancerSku }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) LoadBalancerTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.LoadBalancerTier }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) LoadBalancerType() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkProfile) string { return v.LoadBalancerType }).(pulumi.StringOutput)
}

type NetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) Elem() NetworkProfileOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfile {
		if v != nil {
			return *v
		}
		var ret NetworkProfile
		return ret
	}).(NetworkProfileOutput)
}

func (o NetworkProfilePtrOutput) AzureFrontDoorEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.AzureFrontDoorEnabled
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o NetworkProfilePtrOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerSku
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) LoadBalancerTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerTier
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) LoadBalancerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return &v.LoadBalancerType
	}).(pulumi.StringPtrOutput)
}

type NetworkProfileResponse struct {
	AzureFrontDoorEnabled       *string  `pulumi:"azureFrontDoorEnabled"`
	AzureFrontDoorResourceId    string   `pulumi:"azureFrontDoorResourceId"`
	Capacity                    *int     `pulumi:"capacity"`
	FrontEndPublicIpResourceId  string   `pulumi:"frontEndPublicIpResourceId"`
	LoadBalancerResourceId      string   `pulumi:"loadBalancerResourceId"`
	LoadBalancerSku             *string  `pulumi:"loadBalancerSku"`
	LoadBalancerTier            *string  `pulumi:"loadBalancerTier"`
	LoadBalancerType            string   `pulumi:"loadBalancerType"`
	OutboundPublicIpResourceIds []string `pulumi:"outboundPublicIpResourceIds"`
	VNetResourceId              string   `pulumi:"vNetResourceId"`
}

type NetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) AzureFrontDoorEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.AzureFrontDoorEnabled }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) AzureFrontDoorResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkProfileResponse) string { return v.AzureFrontDoorResourceId }).(pulumi.StringOutput)
}

func (o NetworkProfileResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o NetworkProfileResponseOutput) FrontEndPublicIpResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkProfileResponse) string { return v.FrontEndPublicIpResourceId }).(pulumi.StringOutput)
}

func (o NetworkProfileResponseOutput) LoadBalancerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkProfileResponse) string { return v.LoadBalancerResourceId }).(pulumi.StringOutput)
}

func (o NetworkProfileResponseOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.LoadBalancerSku }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) LoadBalancerTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.LoadBalancerTier }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) LoadBalancerType() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkProfileResponse) string { return v.LoadBalancerType }).(pulumi.StringOutput)
}

func (o NetworkProfileResponseOutput) OutboundPublicIpResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []string { return v.OutboundPublicIpResourceIds }).(pulumi.StringArrayOutput)
}

func (o NetworkProfileResponseOutput) VNetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkProfileResponse) string { return v.VNetResourceId }).(pulumi.StringOutput)
}

type NetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) Elem() NetworkProfileResponseOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) NetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponse
		return ret
	}).(NetworkProfileResponseOutput)
}

func (o NetworkProfileResponsePtrOutput) AzureFrontDoorEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AzureFrontDoorEnabled
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) AzureFrontDoorResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AzureFrontDoorResourceId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) FrontEndPublicIpResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FrontEndPublicIpResourceId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) LoadBalancerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LoadBalancerResourceId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerSku
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) LoadBalancerTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerTier
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) LoadBalancerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LoadBalancerType
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) OutboundPublicIpResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.OutboundPublicIpResourceIds
	}).(pulumi.StringArrayOutput)
}

func (o NetworkProfileResponsePtrOutput) VNetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VNetResourceId
	}).(pulumi.StringPtrOutput)
}

type NodeProfile struct {
	DataDisks []DiskInfo     `pulumi:"dataDisks"`
	Name      *string        `pulumi:"name"`
	NodeSku   string         `pulumi:"nodeSku"`
	OsDisk    DiskInfo       `pulumi:"osDisk"`
	OsImage   OsImageProfile `pulumi:"osImage"`
}





type NodeProfileInput interface {
	pulumi.Input

	ToNodeProfileOutput() NodeProfileOutput
	ToNodeProfileOutputWithContext(context.Context) NodeProfileOutput
}

type NodeProfileArgs struct {
	DataDisks DiskInfoArrayInput    `pulumi:"dataDisks"`
	Name      pulumi.StringPtrInput `pulumi:"name"`
	NodeSku   pulumi.StringInput    `pulumi:"nodeSku"`
	OsDisk    DiskInfoInput         `pulumi:"osDisk"`
	OsImage   OsImageProfileInput   `pulumi:"osImage"`
}

func (NodeProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeProfile)(nil)).Elem()
}

func (i NodeProfileArgs) ToNodeProfileOutput() NodeProfileOutput {
	return i.ToNodeProfileOutputWithContext(context.Background())
}

func (i NodeProfileArgs) ToNodeProfileOutputWithContext(ctx context.Context) NodeProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeProfileOutput)
}

type NodeProfileOutput struct{ *pulumi.OutputState }

func (NodeProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeProfile)(nil)).Elem()
}

func (o NodeProfileOutput) ToNodeProfileOutput() NodeProfileOutput {
	return o
}

func (o NodeProfileOutput) ToNodeProfileOutputWithContext(ctx context.Context) NodeProfileOutput {
	return o
}

func (o NodeProfileOutput) DataDisks() DiskInfoArrayOutput {
	return o.ApplyT(func(v NodeProfile) []DiskInfo { return v.DataDisks }).(DiskInfoArrayOutput)
}

func (o NodeProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NodeProfileOutput) NodeSku() pulumi.StringOutput {
	return o.ApplyT(func(v NodeProfile) string { return v.NodeSku }).(pulumi.StringOutput)
}

func (o NodeProfileOutput) OsDisk() DiskInfoOutput {
	return o.ApplyT(func(v NodeProfile) DiskInfo { return v.OsDisk }).(DiskInfoOutput)
}

func (o NodeProfileOutput) OsImage() OsImageProfileOutput {
	return o.ApplyT(func(v NodeProfile) OsImageProfile { return v.OsImage }).(OsImageProfileOutput)
}

type NodeProfileResponse struct {
	DataDisks       []DiskInfoResponse     `pulumi:"dataDisks"`
	Name            *string                `pulumi:"name"`
	NodeResourceIds []string               `pulumi:"nodeResourceIds"`
	NodeSku         string                 `pulumi:"nodeSku"`
	OsDisk          DiskInfoResponse       `pulumi:"osDisk"`
	OsImage         OsImageProfileResponse `pulumi:"osImage"`
}

type NodeProfileResponseOutput struct{ *pulumi.OutputState }

func (NodeProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeProfileResponse)(nil)).Elem()
}

func (o NodeProfileResponseOutput) ToNodeProfileResponseOutput() NodeProfileResponseOutput {
	return o
}

func (o NodeProfileResponseOutput) ToNodeProfileResponseOutputWithContext(ctx context.Context) NodeProfileResponseOutput {
	return o
}

func (o NodeProfileResponseOutput) DataDisks() DiskInfoResponseArrayOutput {
	return o.ApplyT(func(v NodeProfileResponse) []DiskInfoResponse { return v.DataDisks }).(DiskInfoResponseArrayOutput)
}

func (o NodeProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NodeProfileResponseOutput) NodeResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodeProfileResponse) []string { return v.NodeResourceIds }).(pulumi.StringArrayOutput)
}

func (o NodeProfileResponseOutput) NodeSku() pulumi.StringOutput {
	return o.ApplyT(func(v NodeProfileResponse) string { return v.NodeSku }).(pulumi.StringOutput)
}

func (o NodeProfileResponseOutput) OsDisk() DiskInfoResponseOutput {
	return o.ApplyT(func(v NodeProfileResponse) DiskInfoResponse { return v.OsDisk }).(DiskInfoResponseOutput)
}

func (o NodeProfileResponseOutput) OsImage() OsImageProfileResponseOutput {
	return o.ApplyT(func(v NodeProfileResponse) OsImageProfileResponse { return v.OsImage }).(OsImageProfileResponseOutput)
}

type OSProfile struct {
	AdminPassword   *string     `pulumi:"adminPassword"`
	AdminUsername   *string     `pulumi:"adminUsername"`
	OsConfiguration interface{} `pulumi:"osConfiguration"`
}

type OSProfileResponse struct {
	AdminPassword   *string     `pulumi:"adminPassword"`
	AdminUsername   *string     `pulumi:"adminUsername"`
	OsConfiguration interface{} `pulumi:"osConfiguration"`
}

type OsImageProfile struct {
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}





type OsImageProfileInput interface {
	pulumi.Input

	ToOsImageProfileOutput() OsImageProfileOutput
	ToOsImageProfileOutputWithContext(context.Context) OsImageProfileOutput
}

type OsImageProfileArgs struct {
	Offer     pulumi.StringPtrInput `pulumi:"offer"`
	Publisher pulumi.StringPtrInput `pulumi:"publisher"`
	Sku       pulumi.StringPtrInput `pulumi:"sku"`
	Version   pulumi.StringPtrInput `pulumi:"version"`
}

func (OsImageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OsImageProfile)(nil)).Elem()
}

func (i OsImageProfileArgs) ToOsImageProfileOutput() OsImageProfileOutput {
	return i.ToOsImageProfileOutputWithContext(context.Background())
}

func (i OsImageProfileArgs) ToOsImageProfileOutputWithContext(ctx context.Context) OsImageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsImageProfileOutput)
}

func (i OsImageProfileArgs) ToOsImageProfilePtrOutput() OsImageProfilePtrOutput {
	return i.ToOsImageProfilePtrOutputWithContext(context.Background())
}

func (i OsImageProfileArgs) ToOsImageProfilePtrOutputWithContext(ctx context.Context) OsImageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsImageProfileOutput).ToOsImageProfilePtrOutputWithContext(ctx)
}









type OsImageProfilePtrInput interface {
	pulumi.Input

	ToOsImageProfilePtrOutput() OsImageProfilePtrOutput
	ToOsImageProfilePtrOutputWithContext(context.Context) OsImageProfilePtrOutput
}

type osImageProfilePtrType OsImageProfileArgs

func OsImageProfilePtr(v *OsImageProfileArgs) OsImageProfilePtrInput {
	return (*osImageProfilePtrType)(v)
}

func (*osImageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OsImageProfile)(nil)).Elem()
}

func (i *osImageProfilePtrType) ToOsImageProfilePtrOutput() OsImageProfilePtrOutput {
	return i.ToOsImageProfilePtrOutputWithContext(context.Background())
}

func (i *osImageProfilePtrType) ToOsImageProfilePtrOutputWithContext(ctx context.Context) OsImageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsImageProfilePtrOutput)
}

type OsImageProfileOutput struct{ *pulumi.OutputState }

func (OsImageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsImageProfile)(nil)).Elem()
}

func (o OsImageProfileOutput) ToOsImageProfileOutput() OsImageProfileOutput {
	return o
}

func (o OsImageProfileOutput) ToOsImageProfileOutputWithContext(ctx context.Context) OsImageProfileOutput {
	return o
}

func (o OsImageProfileOutput) ToOsImageProfilePtrOutput() OsImageProfilePtrOutput {
	return o.ToOsImageProfilePtrOutputWithContext(context.Background())
}

func (o OsImageProfileOutput) ToOsImageProfilePtrOutputWithContext(ctx context.Context) OsImageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OsImageProfile) *OsImageProfile {
		return &v
	}).(OsImageProfilePtrOutput)
}

func (o OsImageProfileOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsImageProfile) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o OsImageProfileOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsImageProfile) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o OsImageProfileOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsImageProfile) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o OsImageProfileOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsImageProfile) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type OsImageProfilePtrOutput struct{ *pulumi.OutputState }

func (OsImageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsImageProfile)(nil)).Elem()
}

func (o OsImageProfilePtrOutput) ToOsImageProfilePtrOutput() OsImageProfilePtrOutput {
	return o
}

func (o OsImageProfilePtrOutput) ToOsImageProfilePtrOutputWithContext(ctx context.Context) OsImageProfilePtrOutput {
	return o
}

func (o OsImageProfilePtrOutput) Elem() OsImageProfileOutput {
	return o.ApplyT(func(v *OsImageProfile) OsImageProfile {
		if v != nil {
			return *v
		}
		var ret OsImageProfile
		return ret
	}).(OsImageProfileOutput)
}

func (o OsImageProfilePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsImageProfile) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o OsImageProfilePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsImageProfile) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o OsImageProfilePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsImageProfile) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o OsImageProfilePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsImageProfile) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type OsImageProfileResponse struct {
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}

type OsImageProfileResponseOutput struct{ *pulumi.OutputState }

func (OsImageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsImageProfileResponse)(nil)).Elem()
}

func (o OsImageProfileResponseOutput) ToOsImageProfileResponseOutput() OsImageProfileResponseOutput {
	return o
}

func (o OsImageProfileResponseOutput) ToOsImageProfileResponseOutputWithContext(ctx context.Context) OsImageProfileResponseOutput {
	return o
}

func (o OsImageProfileResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsImageProfileResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o OsImageProfileResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsImageProfileResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o OsImageProfileResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsImageProfileResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o OsImageProfileResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsImageProfileResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type OsImageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OsImageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsImageProfileResponse)(nil)).Elem()
}

func (o OsImageProfileResponsePtrOutput) ToOsImageProfileResponsePtrOutput() OsImageProfileResponsePtrOutput {
	return o
}

func (o OsImageProfileResponsePtrOutput) ToOsImageProfileResponsePtrOutputWithContext(ctx context.Context) OsImageProfileResponsePtrOutput {
	return o
}

func (o OsImageProfileResponsePtrOutput) Elem() OsImageProfileResponseOutput {
	return o.ApplyT(func(v *OsImageProfileResponse) OsImageProfileResponse {
		if v != nil {
			return *v
		}
		var ret OsImageProfileResponse
		return ret
	}).(OsImageProfileResponseOutput)
}

func (o OsImageProfileResponsePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsImageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o OsImageProfileResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsImageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o OsImageProfileResponsePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsImageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o OsImageProfileResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsImageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type OsSapConfiguration struct {
	DeployerVmPackages *DeployerVmPackages `pulumi:"deployerVmPackages"`
	SapFqdn            *string             `pulumi:"sapFqdn"`
}

type OsSapConfigurationResponse struct {
	DeployerVmPackages *DeployerVmPackagesResponse `pulumi:"deployerVmPackages"`
	SapFqdn            *string                     `pulumi:"sapFqdn"`
}

type PhpProfile struct {
	Version string `pulumi:"version"`
}





type PhpProfileInput interface {
	pulumi.Input

	ToPhpProfileOutput() PhpProfileOutput
	ToPhpProfileOutputWithContext(context.Context) PhpProfileOutput
}

type PhpProfileArgs struct {
	Version pulumi.StringInput `pulumi:"version"`
}

func (PhpProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PhpProfile)(nil)).Elem()
}

func (i PhpProfileArgs) ToPhpProfileOutput() PhpProfileOutput {
	return i.ToPhpProfileOutputWithContext(context.Background())
}

func (i PhpProfileArgs) ToPhpProfileOutputWithContext(ctx context.Context) PhpProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhpProfileOutput)
}

func (i PhpProfileArgs) ToPhpProfilePtrOutput() PhpProfilePtrOutput {
	return i.ToPhpProfilePtrOutputWithContext(context.Background())
}

func (i PhpProfileArgs) ToPhpProfilePtrOutputWithContext(ctx context.Context) PhpProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhpProfileOutput).ToPhpProfilePtrOutputWithContext(ctx)
}









type PhpProfilePtrInput interface {
	pulumi.Input

	ToPhpProfilePtrOutput() PhpProfilePtrOutput
	ToPhpProfilePtrOutputWithContext(context.Context) PhpProfilePtrOutput
}

type phpProfilePtrType PhpProfileArgs

func PhpProfilePtr(v *PhpProfileArgs) PhpProfilePtrInput {
	return (*phpProfilePtrType)(v)
}

func (*phpProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PhpProfile)(nil)).Elem()
}

func (i *phpProfilePtrType) ToPhpProfilePtrOutput() PhpProfilePtrOutput {
	return i.ToPhpProfilePtrOutputWithContext(context.Background())
}

func (i *phpProfilePtrType) ToPhpProfilePtrOutputWithContext(ctx context.Context) PhpProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhpProfilePtrOutput)
}

type PhpProfileOutput struct{ *pulumi.OutputState }

func (PhpProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PhpProfile)(nil)).Elem()
}

func (o PhpProfileOutput) ToPhpProfileOutput() PhpProfileOutput {
	return o
}

func (o PhpProfileOutput) ToPhpProfileOutputWithContext(ctx context.Context) PhpProfileOutput {
	return o
}

func (o PhpProfileOutput) ToPhpProfilePtrOutput() PhpProfilePtrOutput {
	return o.ToPhpProfilePtrOutputWithContext(context.Background())
}

func (o PhpProfileOutput) ToPhpProfilePtrOutputWithContext(ctx context.Context) PhpProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PhpProfile) *PhpProfile {
		return &v
	}).(PhpProfilePtrOutput)
}

func (o PhpProfileOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v PhpProfile) string { return v.Version }).(pulumi.StringOutput)
}

type PhpProfilePtrOutput struct{ *pulumi.OutputState }

func (PhpProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PhpProfile)(nil)).Elem()
}

func (o PhpProfilePtrOutput) ToPhpProfilePtrOutput() PhpProfilePtrOutput {
	return o
}

func (o PhpProfilePtrOutput) ToPhpProfilePtrOutputWithContext(ctx context.Context) PhpProfilePtrOutput {
	return o
}

func (o PhpProfilePtrOutput) Elem() PhpProfileOutput {
	return o.ApplyT(func(v *PhpProfile) PhpProfile {
		if v != nil {
			return *v
		}
		var ret PhpProfile
		return ret
	}).(PhpProfileOutput)
}

func (o PhpProfilePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PhpProfile) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type PhpProfileResponse struct {
	Version string `pulumi:"version"`
}

type PhpProfileResponseOutput struct{ *pulumi.OutputState }

func (PhpProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PhpProfileResponse)(nil)).Elem()
}

func (o PhpProfileResponseOutput) ToPhpProfileResponseOutput() PhpProfileResponseOutput {
	return o
}

func (o PhpProfileResponseOutput) ToPhpProfileResponseOutputWithContext(ctx context.Context) PhpProfileResponseOutput {
	return o
}

func (o PhpProfileResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v PhpProfileResponse) string { return v.Version }).(pulumi.StringOutput)
}

type PhpProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (PhpProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PhpProfileResponse)(nil)).Elem()
}

func (o PhpProfileResponsePtrOutput) ToPhpProfileResponsePtrOutput() PhpProfileResponsePtrOutput {
	return o
}

func (o PhpProfileResponsePtrOutput) ToPhpProfileResponsePtrOutputWithContext(ctx context.Context) PhpProfileResponsePtrOutput {
	return o
}

func (o PhpProfileResponsePtrOutput) Elem() PhpProfileResponseOutput {
	return o.ApplyT(func(v *PhpProfileResponse) PhpProfileResponse {
		if v != nil {
			return *v
		}
		var ret PhpProfileResponse
		return ret
	}).(PhpProfileResponseOutput)
}

func (o PhpProfileResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PhpProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type PhpWorkloadResourceIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type PhpWorkloadResourceIdentityInput interface {
	pulumi.Input

	ToPhpWorkloadResourceIdentityOutput() PhpWorkloadResourceIdentityOutput
	ToPhpWorkloadResourceIdentityOutputWithContext(context.Context) PhpWorkloadResourceIdentityOutput
}

type PhpWorkloadResourceIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (PhpWorkloadResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PhpWorkloadResourceIdentity)(nil)).Elem()
}

func (i PhpWorkloadResourceIdentityArgs) ToPhpWorkloadResourceIdentityOutput() PhpWorkloadResourceIdentityOutput {
	return i.ToPhpWorkloadResourceIdentityOutputWithContext(context.Background())
}

func (i PhpWorkloadResourceIdentityArgs) ToPhpWorkloadResourceIdentityOutputWithContext(ctx context.Context) PhpWorkloadResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhpWorkloadResourceIdentityOutput)
}

func (i PhpWorkloadResourceIdentityArgs) ToPhpWorkloadResourceIdentityPtrOutput() PhpWorkloadResourceIdentityPtrOutput {
	return i.ToPhpWorkloadResourceIdentityPtrOutputWithContext(context.Background())
}

func (i PhpWorkloadResourceIdentityArgs) ToPhpWorkloadResourceIdentityPtrOutputWithContext(ctx context.Context) PhpWorkloadResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhpWorkloadResourceIdentityOutput).ToPhpWorkloadResourceIdentityPtrOutputWithContext(ctx)
}









type PhpWorkloadResourceIdentityPtrInput interface {
	pulumi.Input

	ToPhpWorkloadResourceIdentityPtrOutput() PhpWorkloadResourceIdentityPtrOutput
	ToPhpWorkloadResourceIdentityPtrOutputWithContext(context.Context) PhpWorkloadResourceIdentityPtrOutput
}

type phpWorkloadResourceIdentityPtrType PhpWorkloadResourceIdentityArgs

func PhpWorkloadResourceIdentityPtr(v *PhpWorkloadResourceIdentityArgs) PhpWorkloadResourceIdentityPtrInput {
	return (*phpWorkloadResourceIdentityPtrType)(v)
}

func (*phpWorkloadResourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PhpWorkloadResourceIdentity)(nil)).Elem()
}

func (i *phpWorkloadResourceIdentityPtrType) ToPhpWorkloadResourceIdentityPtrOutput() PhpWorkloadResourceIdentityPtrOutput {
	return i.ToPhpWorkloadResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *phpWorkloadResourceIdentityPtrType) ToPhpWorkloadResourceIdentityPtrOutputWithContext(ctx context.Context) PhpWorkloadResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhpWorkloadResourceIdentityPtrOutput)
}

type PhpWorkloadResourceIdentityOutput struct{ *pulumi.OutputState }

func (PhpWorkloadResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PhpWorkloadResourceIdentity)(nil)).Elem()
}

func (o PhpWorkloadResourceIdentityOutput) ToPhpWorkloadResourceIdentityOutput() PhpWorkloadResourceIdentityOutput {
	return o
}

func (o PhpWorkloadResourceIdentityOutput) ToPhpWorkloadResourceIdentityOutputWithContext(ctx context.Context) PhpWorkloadResourceIdentityOutput {
	return o
}

func (o PhpWorkloadResourceIdentityOutput) ToPhpWorkloadResourceIdentityPtrOutput() PhpWorkloadResourceIdentityPtrOutput {
	return o.ToPhpWorkloadResourceIdentityPtrOutputWithContext(context.Background())
}

func (o PhpWorkloadResourceIdentityOutput) ToPhpWorkloadResourceIdentityPtrOutputWithContext(ctx context.Context) PhpWorkloadResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PhpWorkloadResourceIdentity) *PhpWorkloadResourceIdentity {
		return &v
	}).(PhpWorkloadResourceIdentityPtrOutput)
}

func (o PhpWorkloadResourceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PhpWorkloadResourceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o PhpWorkloadResourceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v PhpWorkloadResourceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type PhpWorkloadResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (PhpWorkloadResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PhpWorkloadResourceIdentity)(nil)).Elem()
}

func (o PhpWorkloadResourceIdentityPtrOutput) ToPhpWorkloadResourceIdentityPtrOutput() PhpWorkloadResourceIdentityPtrOutput {
	return o
}

func (o PhpWorkloadResourceIdentityPtrOutput) ToPhpWorkloadResourceIdentityPtrOutputWithContext(ctx context.Context) PhpWorkloadResourceIdentityPtrOutput {
	return o
}

func (o PhpWorkloadResourceIdentityPtrOutput) Elem() PhpWorkloadResourceIdentityOutput {
	return o.ApplyT(func(v *PhpWorkloadResourceIdentity) PhpWorkloadResourceIdentity {
		if v != nil {
			return *v
		}
		var ret PhpWorkloadResourceIdentity
		return ret
	}).(PhpWorkloadResourceIdentityOutput)
}

func (o PhpWorkloadResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PhpWorkloadResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o PhpWorkloadResourceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *PhpWorkloadResourceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type PhpWorkloadResourceResponseIdentity struct {
	Type                   string                                  `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type PhpWorkloadResourceResponseIdentityOutput struct{ *pulumi.OutputState }

func (PhpWorkloadResourceResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PhpWorkloadResourceResponseIdentity)(nil)).Elem()
}

func (o PhpWorkloadResourceResponseIdentityOutput) ToPhpWorkloadResourceResponseIdentityOutput() PhpWorkloadResourceResponseIdentityOutput {
	return o
}

func (o PhpWorkloadResourceResponseIdentityOutput) ToPhpWorkloadResourceResponseIdentityOutputWithContext(ctx context.Context) PhpWorkloadResourceResponseIdentityOutput {
	return o
}

func (o PhpWorkloadResourceResponseIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PhpWorkloadResourceResponseIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o PhpWorkloadResourceResponseIdentityOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v PhpWorkloadResourceResponseIdentity) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type PhpWorkloadResourceResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (PhpWorkloadResourceResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PhpWorkloadResourceResponseIdentity)(nil)).Elem()
}

func (o PhpWorkloadResourceResponseIdentityPtrOutput) ToPhpWorkloadResourceResponseIdentityPtrOutput() PhpWorkloadResourceResponseIdentityPtrOutput {
	return o
}

func (o PhpWorkloadResourceResponseIdentityPtrOutput) ToPhpWorkloadResourceResponseIdentityPtrOutputWithContext(ctx context.Context) PhpWorkloadResourceResponseIdentityPtrOutput {
	return o
}

func (o PhpWorkloadResourceResponseIdentityPtrOutput) Elem() PhpWorkloadResourceResponseIdentityOutput {
	return o.ApplyT(func(v *PhpWorkloadResourceResponseIdentity) PhpWorkloadResourceResponseIdentity {
		if v != nil {
			return *v
		}
		var ret PhpWorkloadResourceResponseIdentity
		return ret
	}).(PhpWorkloadResourceResponseIdentityOutput)
}

func (o PhpWorkloadResourceResponseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PhpWorkloadResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o PhpWorkloadResourceResponseIdentityPtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *PhpWorkloadResourceResponseIdentity) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type PrometheusHaClusterProviderInstanceProperties struct {
	ClusterName   *string `pulumi:"clusterName"`
	Hostname      *string `pulumi:"hostname"`
	PrometheusUrl *string `pulumi:"prometheusUrl"`
	ProviderType  string  `pulumi:"providerType"`
	Sid           *string `pulumi:"sid"`
}

type PrometheusHaClusterProviderInstancePropertiesResponse struct {
	ClusterName   *string `pulumi:"clusterName"`
	Hostname      *string `pulumi:"hostname"`
	PrometheusUrl *string `pulumi:"prometheusUrl"`
	ProviderType  string  `pulumi:"providerType"`
	Sid           *string `pulumi:"sid"`
}

type PrometheusOSProviderInstanceProperties struct {
	PrometheusUrl *string `pulumi:"prometheusUrl"`
	ProviderType  string  `pulumi:"providerType"`
}

type PrometheusOSProviderInstancePropertiesResponse struct {
	PrometheusUrl *string `pulumi:"prometheusUrl"`
	ProviderType  string  `pulumi:"providerType"`
}

type ProviderInstancePropertiesResponseErrors struct {
	Code       string                  `pulumi:"code"`
	Details    []ErrorResponse         `pulumi:"details"`
	InnerError ErrorResponseInnerError `pulumi:"innerError"`
	Message    string                  `pulumi:"message"`
	Target     string                  `pulumi:"target"`
}

type ProviderInstancePropertiesResponseErrorsOutput struct{ *pulumi.OutputState }

func (ProviderInstancePropertiesResponseErrorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderInstancePropertiesResponseErrors)(nil)).Elem()
}

func (o ProviderInstancePropertiesResponseErrorsOutput) ToProviderInstancePropertiesResponseErrorsOutput() ProviderInstancePropertiesResponseErrorsOutput {
	return o
}

func (o ProviderInstancePropertiesResponseErrorsOutput) ToProviderInstancePropertiesResponseErrorsOutputWithContext(ctx context.Context) ProviderInstancePropertiesResponseErrorsOutput {
	return o
}

func (o ProviderInstancePropertiesResponseErrorsOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderInstancePropertiesResponseErrors) string { return v.Code }).(pulumi.StringOutput)
}

func (o ProviderInstancePropertiesResponseErrorsOutput) Details() ErrorResponseArrayOutput {
	return o.ApplyT(func(v ProviderInstancePropertiesResponseErrors) []ErrorResponse { return v.Details }).(ErrorResponseArrayOutput)
}

func (o ProviderInstancePropertiesResponseErrorsOutput) InnerError() ErrorResponseInnerErrorOutput {
	return o.ApplyT(func(v ProviderInstancePropertiesResponseErrors) ErrorResponseInnerError { return v.InnerError }).(ErrorResponseInnerErrorOutput)
}

func (o ProviderInstancePropertiesResponseErrorsOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderInstancePropertiesResponseErrors) string { return v.Message }).(pulumi.StringOutput)
}

func (o ProviderInstancePropertiesResponseErrorsOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderInstancePropertiesResponseErrors) string { return v.Target }).(pulumi.StringOutput)
}

type SAPAvailabilityZonePairResponse struct {
	ZoneA *float64 `pulumi:"zoneA"`
	ZoneB *float64 `pulumi:"zoneB"`
}

type SAPAvailabilityZonePairResponseOutput struct{ *pulumi.OutputState }

func (SAPAvailabilityZonePairResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SAPAvailabilityZonePairResponse)(nil)).Elem()
}

func (o SAPAvailabilityZonePairResponseOutput) ToSAPAvailabilityZonePairResponseOutput() SAPAvailabilityZonePairResponseOutput {
	return o
}

func (o SAPAvailabilityZonePairResponseOutput) ToSAPAvailabilityZonePairResponseOutputWithContext(ctx context.Context) SAPAvailabilityZonePairResponseOutput {
	return o
}

func (o SAPAvailabilityZonePairResponseOutput) ZoneA() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SAPAvailabilityZonePairResponse) *float64 { return v.ZoneA }).(pulumi.Float64PtrOutput)
}

func (o SAPAvailabilityZonePairResponseOutput) ZoneB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SAPAvailabilityZonePairResponse) *float64 { return v.ZoneB }).(pulumi.Float64PtrOutput)
}

type SAPAvailabilityZonePairResponseArrayOutput struct{ *pulumi.OutputState }

func (SAPAvailabilityZonePairResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SAPAvailabilityZonePairResponse)(nil)).Elem()
}

func (o SAPAvailabilityZonePairResponseArrayOutput) ToSAPAvailabilityZonePairResponseArrayOutput() SAPAvailabilityZonePairResponseArrayOutput {
	return o
}

func (o SAPAvailabilityZonePairResponseArrayOutput) ToSAPAvailabilityZonePairResponseArrayOutputWithContext(ctx context.Context) SAPAvailabilityZonePairResponseArrayOutput {
	return o
}

func (o SAPAvailabilityZonePairResponseArrayOutput) Index(i pulumi.IntInput) SAPAvailabilityZonePairResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SAPAvailabilityZonePairResponse {
		return vs[0].([]SAPAvailabilityZonePairResponse)[vs[1].(int)]
	}).(SAPAvailabilityZonePairResponseOutput)
}

type SAPDiskConfigurationResponse struct {
	DiskCount         *float64 `pulumi:"diskCount"`
	DiskIopsReadWrite *float64 `pulumi:"diskIopsReadWrite"`
	DiskMBpsReadWrite *float64 `pulumi:"diskMBpsReadWrite"`
	DiskSizeGB        *float64 `pulumi:"diskSizeGB"`
	DiskStorageType   *string  `pulumi:"diskStorageType"`
	DiskType          *string  `pulumi:"diskType"`
	Volume            *string  `pulumi:"volume"`
}

type SAPDiskConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SAPDiskConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SAPDiskConfigurationResponse)(nil)).Elem()
}

func (o SAPDiskConfigurationResponseOutput) ToSAPDiskConfigurationResponseOutput() SAPDiskConfigurationResponseOutput {
	return o
}

func (o SAPDiskConfigurationResponseOutput) ToSAPDiskConfigurationResponseOutputWithContext(ctx context.Context) SAPDiskConfigurationResponseOutput {
	return o
}

func (o SAPDiskConfigurationResponseOutput) DiskCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SAPDiskConfigurationResponse) *float64 { return v.DiskCount }).(pulumi.Float64PtrOutput)
}

func (o SAPDiskConfigurationResponseOutput) DiskIopsReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SAPDiskConfigurationResponse) *float64 { return v.DiskIopsReadWrite }).(pulumi.Float64PtrOutput)
}

func (o SAPDiskConfigurationResponseOutput) DiskMBpsReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SAPDiskConfigurationResponse) *float64 { return v.DiskMBpsReadWrite }).(pulumi.Float64PtrOutput)
}

func (o SAPDiskConfigurationResponseOutput) DiskSizeGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SAPDiskConfigurationResponse) *float64 { return v.DiskSizeGB }).(pulumi.Float64PtrOutput)
}

func (o SAPDiskConfigurationResponseOutput) DiskStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPDiskConfigurationResponse) *string { return v.DiskStorageType }).(pulumi.StringPtrOutput)
}

func (o SAPDiskConfigurationResponseOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPDiskConfigurationResponse) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

func (o SAPDiskConfigurationResponseOutput) Volume() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPDiskConfigurationResponse) *string { return v.Volume }).(pulumi.StringPtrOutput)
}

type SAPDiskConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (SAPDiskConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SAPDiskConfigurationResponse)(nil)).Elem()
}

func (o SAPDiskConfigurationResponseArrayOutput) ToSAPDiskConfigurationResponseArrayOutput() SAPDiskConfigurationResponseArrayOutput {
	return o
}

func (o SAPDiskConfigurationResponseArrayOutput) ToSAPDiskConfigurationResponseArrayOutputWithContext(ctx context.Context) SAPDiskConfigurationResponseArrayOutput {
	return o
}

func (o SAPDiskConfigurationResponseArrayOutput) Index(i pulumi.IntInput) SAPDiskConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SAPDiskConfigurationResponse {
		return vs[0].([]SAPDiskConfigurationResponse)[vs[1].(int)]
	}).(SAPDiskConfigurationResponseOutput)
}

type SAPInstallWithoutOSConfigSoftwareConfiguration struct {
	BomUrl                                string                                 `pulumi:"bomUrl"`
	HighAvailabilitySoftwareConfiguration *HighAvailabilitySoftwareConfiguration `pulumi:"highAvailabilitySoftwareConfiguration"`
	SapBitsStorageAccountId               string                                 `pulumi:"sapBitsStorageAccountId"`
	SoftwareInstallationType              string                                 `pulumi:"softwareInstallationType"`
	SoftwareVersion                       string                                 `pulumi:"softwareVersion"`
}

type SAPInstallWithoutOSConfigSoftwareConfigurationResponse struct {
	BomUrl                                string                                         `pulumi:"bomUrl"`
	HighAvailabilitySoftwareConfiguration *HighAvailabilitySoftwareConfigurationResponse `pulumi:"highAvailabilitySoftwareConfiguration"`
	SapBitsStorageAccountId               string                                         `pulumi:"sapBitsStorageAccountId"`
	SoftwareInstallationType              string                                         `pulumi:"softwareInstallationType"`
	SoftwareVersion                       string                                         `pulumi:"softwareVersion"`
}

type SAPSupportedSkuResponse struct {
	IsAppServerCertified *bool   `pulumi:"isAppServerCertified"`
	IsDatabaseCertified  *bool   `pulumi:"isDatabaseCertified"`
	VmSku                *string `pulumi:"vmSku"`
}

type SAPSupportedSkuResponseOutput struct{ *pulumi.OutputState }

func (SAPSupportedSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SAPSupportedSkuResponse)(nil)).Elem()
}

func (o SAPSupportedSkuResponseOutput) ToSAPSupportedSkuResponseOutput() SAPSupportedSkuResponseOutput {
	return o
}

func (o SAPSupportedSkuResponseOutput) ToSAPSupportedSkuResponseOutputWithContext(ctx context.Context) SAPSupportedSkuResponseOutput {
	return o
}

func (o SAPSupportedSkuResponseOutput) IsAppServerCertified() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SAPSupportedSkuResponse) *bool { return v.IsAppServerCertified }).(pulumi.BoolPtrOutput)
}

func (o SAPSupportedSkuResponseOutput) IsDatabaseCertified() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SAPSupportedSkuResponse) *bool { return v.IsDatabaseCertified }).(pulumi.BoolPtrOutput)
}

func (o SAPSupportedSkuResponseOutput) VmSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SAPSupportedSkuResponse) *string { return v.VmSku }).(pulumi.StringPtrOutput)
}

type SAPSupportedSkuResponseArrayOutput struct{ *pulumi.OutputState }

func (SAPSupportedSkuResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SAPSupportedSkuResponse)(nil)).Elem()
}

func (o SAPSupportedSkuResponseArrayOutput) ToSAPSupportedSkuResponseArrayOutput() SAPSupportedSkuResponseArrayOutput {
	return o
}

func (o SAPSupportedSkuResponseArrayOutput) ToSAPSupportedSkuResponseArrayOutputWithContext(ctx context.Context) SAPSupportedSkuResponseArrayOutput {
	return o
}

func (o SAPSupportedSkuResponseArrayOutput) Index(i pulumi.IntInput) SAPSupportedSkuResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SAPSupportedSkuResponse {
		return vs[0].([]SAPSupportedSkuResponse)[vs[1].(int)]
	}).(SAPSupportedSkuResponseOutput)
}

type SAPVirtualInstanceErrorResponse struct {
	Properties *ErrorDefinitionResponse `pulumi:"properties"`
}

type SAPVirtualInstanceErrorResponseOutput struct{ *pulumi.OutputState }

func (SAPVirtualInstanceErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SAPVirtualInstanceErrorResponse)(nil)).Elem()
}

func (o SAPVirtualInstanceErrorResponseOutput) ToSAPVirtualInstanceErrorResponseOutput() SAPVirtualInstanceErrorResponseOutput {
	return o
}

func (o SAPVirtualInstanceErrorResponseOutput) ToSAPVirtualInstanceErrorResponseOutputWithContext(ctx context.Context) SAPVirtualInstanceErrorResponseOutput {
	return o
}

func (o SAPVirtualInstanceErrorResponseOutput) Properties() ErrorDefinitionResponsePtrOutput {
	return o.ApplyT(func(v SAPVirtualInstanceErrorResponse) *ErrorDefinitionResponse { return v.Properties }).(ErrorDefinitionResponsePtrOutput)
}

type SapNetWeaverProviderInstanceProperties struct {
	ProviderType         string   `pulumi:"providerType"`
	SapClientId          *string  `pulumi:"sapClientId"`
	SapHostFileEntries   []string `pulumi:"sapHostFileEntries"`
	SapHostname          *string  `pulumi:"sapHostname"`
	SapInstanceNr        *string  `pulumi:"sapInstanceNr"`
	SapPassword          *string  `pulumi:"sapPassword"`
	SapPasswordUri       *string  `pulumi:"sapPasswordUri"`
	SapPortNumber        *string  `pulumi:"sapPortNumber"`
	SapSid               *string  `pulumi:"sapSid"`
	SapSslCertificateUri *string  `pulumi:"sapSslCertificateUri"`
	SapUsername          *string  `pulumi:"sapUsername"`
}

type SapNetWeaverProviderInstancePropertiesResponse struct {
	ProviderType         string   `pulumi:"providerType"`
	SapClientId          *string  `pulumi:"sapClientId"`
	SapHostFileEntries   []string `pulumi:"sapHostFileEntries"`
	SapHostname          *string  `pulumi:"sapHostname"`
	SapInstanceNr        *string  `pulumi:"sapInstanceNr"`
	SapPassword          *string  `pulumi:"sapPassword"`
	SapPasswordUri       *string  `pulumi:"sapPasswordUri"`
	SapPortNumber        *string  `pulumi:"sapPortNumber"`
	SapSid               *string  `pulumi:"sapSid"`
	SapSslCertificateUri *string  `pulumi:"sapSslCertificateUri"`
	SapUsername          *string  `pulumi:"sapUsername"`
}

type SearchProfile struct {
	DataDisks  []DiskInfo     `pulumi:"dataDisks"`
	Name       *string        `pulumi:"name"`
	NodeSku    string         `pulumi:"nodeSku"`
	OsDisk     DiskInfo       `pulumi:"osDisk"`
	OsImage    OsImageProfile `pulumi:"osImage"`
	SearchType string         `pulumi:"searchType"`
}





type SearchProfileInput interface {
	pulumi.Input

	ToSearchProfileOutput() SearchProfileOutput
	ToSearchProfileOutputWithContext(context.Context) SearchProfileOutput
}

type SearchProfileArgs struct {
	DataDisks  DiskInfoArrayInput    `pulumi:"dataDisks"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	NodeSku    pulumi.StringInput    `pulumi:"nodeSku"`
	OsDisk     DiskInfoInput         `pulumi:"osDisk"`
	OsImage    OsImageProfileInput   `pulumi:"osImage"`
	SearchType pulumi.StringInput    `pulumi:"searchType"`
}

func (SearchProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchProfile)(nil)).Elem()
}

func (i SearchProfileArgs) ToSearchProfileOutput() SearchProfileOutput {
	return i.ToSearchProfileOutputWithContext(context.Background())
}

func (i SearchProfileArgs) ToSearchProfileOutputWithContext(ctx context.Context) SearchProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SearchProfileOutput)
}

func (i SearchProfileArgs) ToSearchProfilePtrOutput() SearchProfilePtrOutput {
	return i.ToSearchProfilePtrOutputWithContext(context.Background())
}

func (i SearchProfileArgs) ToSearchProfilePtrOutputWithContext(ctx context.Context) SearchProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SearchProfileOutput).ToSearchProfilePtrOutputWithContext(ctx)
}









type SearchProfilePtrInput interface {
	pulumi.Input

	ToSearchProfilePtrOutput() SearchProfilePtrOutput
	ToSearchProfilePtrOutputWithContext(context.Context) SearchProfilePtrOutput
}

type searchProfilePtrType SearchProfileArgs

func SearchProfilePtr(v *SearchProfileArgs) SearchProfilePtrInput {
	return (*searchProfilePtrType)(v)
}

func (*searchProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SearchProfile)(nil)).Elem()
}

func (i *searchProfilePtrType) ToSearchProfilePtrOutput() SearchProfilePtrOutput {
	return i.ToSearchProfilePtrOutputWithContext(context.Background())
}

func (i *searchProfilePtrType) ToSearchProfilePtrOutputWithContext(ctx context.Context) SearchProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SearchProfilePtrOutput)
}

type SearchProfileOutput struct{ *pulumi.OutputState }

func (SearchProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchProfile)(nil)).Elem()
}

func (o SearchProfileOutput) ToSearchProfileOutput() SearchProfileOutput {
	return o
}

func (o SearchProfileOutput) ToSearchProfileOutputWithContext(ctx context.Context) SearchProfileOutput {
	return o
}

func (o SearchProfileOutput) ToSearchProfilePtrOutput() SearchProfilePtrOutput {
	return o.ToSearchProfilePtrOutputWithContext(context.Background())
}

func (o SearchProfileOutput) ToSearchProfilePtrOutputWithContext(ctx context.Context) SearchProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SearchProfile) *SearchProfile {
		return &v
	}).(SearchProfilePtrOutput)
}

func (o SearchProfileOutput) DataDisks() DiskInfoArrayOutput {
	return o.ApplyT(func(v SearchProfile) []DiskInfo { return v.DataDisks }).(DiskInfoArrayOutput)
}

func (o SearchProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SearchProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SearchProfileOutput) NodeSku() pulumi.StringOutput {
	return o.ApplyT(func(v SearchProfile) string { return v.NodeSku }).(pulumi.StringOutput)
}

func (o SearchProfileOutput) OsDisk() DiskInfoOutput {
	return o.ApplyT(func(v SearchProfile) DiskInfo { return v.OsDisk }).(DiskInfoOutput)
}

func (o SearchProfileOutput) OsImage() OsImageProfileOutput {
	return o.ApplyT(func(v SearchProfile) OsImageProfile { return v.OsImage }).(OsImageProfileOutput)
}

func (o SearchProfileOutput) SearchType() pulumi.StringOutput {
	return o.ApplyT(func(v SearchProfile) string { return v.SearchType }).(pulumi.StringOutput)
}

type SearchProfilePtrOutput struct{ *pulumi.OutputState }

func (SearchProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SearchProfile)(nil)).Elem()
}

func (o SearchProfilePtrOutput) ToSearchProfilePtrOutput() SearchProfilePtrOutput {
	return o
}

func (o SearchProfilePtrOutput) ToSearchProfilePtrOutputWithContext(ctx context.Context) SearchProfilePtrOutput {
	return o
}

func (o SearchProfilePtrOutput) Elem() SearchProfileOutput {
	return o.ApplyT(func(v *SearchProfile) SearchProfile {
		if v != nil {
			return *v
		}
		var ret SearchProfile
		return ret
	}).(SearchProfileOutput)
}

func (o SearchProfilePtrOutput) DataDisks() DiskInfoArrayOutput {
	return o.ApplyT(func(v *SearchProfile) []DiskInfo {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(DiskInfoArrayOutput)
}

func (o SearchProfilePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchProfile) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SearchProfilePtrOutput) NodeSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchProfile) *string {
		if v == nil {
			return nil
		}
		return &v.NodeSku
	}).(pulumi.StringPtrOutput)
}

func (o SearchProfilePtrOutput) OsDisk() DiskInfoPtrOutput {
	return o.ApplyT(func(v *SearchProfile) *DiskInfo {
		if v == nil {
			return nil
		}
		return &v.OsDisk
	}).(DiskInfoPtrOutput)
}

func (o SearchProfilePtrOutput) OsImage() OsImageProfilePtrOutput {
	return o.ApplyT(func(v *SearchProfile) *OsImageProfile {
		if v == nil {
			return nil
		}
		return &v.OsImage
	}).(OsImageProfilePtrOutput)
}

func (o SearchProfilePtrOutput) SearchType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchProfile) *string {
		if v == nil {
			return nil
		}
		return &v.SearchType
	}).(pulumi.StringPtrOutput)
}

type SearchProfileResponse struct {
	DataDisks       []DiskInfoResponse     `pulumi:"dataDisks"`
	Name            *string                `pulumi:"name"`
	NodeResourceIds []string               `pulumi:"nodeResourceIds"`
	NodeSku         string                 `pulumi:"nodeSku"`
	OsDisk          DiskInfoResponse       `pulumi:"osDisk"`
	OsImage         OsImageProfileResponse `pulumi:"osImage"`
	SearchType      string                 `pulumi:"searchType"`
}

type SearchProfileResponseOutput struct{ *pulumi.OutputState }

func (SearchProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchProfileResponse)(nil)).Elem()
}

func (o SearchProfileResponseOutput) ToSearchProfileResponseOutput() SearchProfileResponseOutput {
	return o
}

func (o SearchProfileResponseOutput) ToSearchProfileResponseOutputWithContext(ctx context.Context) SearchProfileResponseOutput {
	return o
}

func (o SearchProfileResponseOutput) DataDisks() DiskInfoResponseArrayOutput {
	return o.ApplyT(func(v SearchProfileResponse) []DiskInfoResponse { return v.DataDisks }).(DiskInfoResponseArrayOutput)
}

func (o SearchProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SearchProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SearchProfileResponseOutput) NodeResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SearchProfileResponse) []string { return v.NodeResourceIds }).(pulumi.StringArrayOutput)
}

func (o SearchProfileResponseOutput) NodeSku() pulumi.StringOutput {
	return o.ApplyT(func(v SearchProfileResponse) string { return v.NodeSku }).(pulumi.StringOutput)
}

func (o SearchProfileResponseOutput) OsDisk() DiskInfoResponseOutput {
	return o.ApplyT(func(v SearchProfileResponse) DiskInfoResponse { return v.OsDisk }).(DiskInfoResponseOutput)
}

func (o SearchProfileResponseOutput) OsImage() OsImageProfileResponseOutput {
	return o.ApplyT(func(v SearchProfileResponse) OsImageProfileResponse { return v.OsImage }).(OsImageProfileResponseOutput)
}

func (o SearchProfileResponseOutput) SearchType() pulumi.StringOutput {
	return o.ApplyT(func(v SearchProfileResponse) string { return v.SearchType }).(pulumi.StringOutput)
}

type SearchProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (SearchProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SearchProfileResponse)(nil)).Elem()
}

func (o SearchProfileResponsePtrOutput) ToSearchProfileResponsePtrOutput() SearchProfileResponsePtrOutput {
	return o
}

func (o SearchProfileResponsePtrOutput) ToSearchProfileResponsePtrOutputWithContext(ctx context.Context) SearchProfileResponsePtrOutput {
	return o
}

func (o SearchProfileResponsePtrOutput) Elem() SearchProfileResponseOutput {
	return o.ApplyT(func(v *SearchProfileResponse) SearchProfileResponse {
		if v != nil {
			return *v
		}
		var ret SearchProfileResponse
		return ret
	}).(SearchProfileResponseOutput)
}

func (o SearchProfileResponsePtrOutput) DataDisks() DiskInfoResponseArrayOutput {
	return o.ApplyT(func(v *SearchProfileResponse) []DiskInfoResponse {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(DiskInfoResponseArrayOutput)
}

func (o SearchProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SearchProfileResponsePtrOutput) NodeResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SearchProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.NodeResourceIds
	}).(pulumi.StringArrayOutput)
}

func (o SearchProfileResponsePtrOutput) NodeSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NodeSku
	}).(pulumi.StringPtrOutput)
}

func (o SearchProfileResponsePtrOutput) OsDisk() DiskInfoResponsePtrOutput {
	return o.ApplyT(func(v *SearchProfileResponse) *DiskInfoResponse {
		if v == nil {
			return nil
		}
		return &v.OsDisk
	}).(DiskInfoResponsePtrOutput)
}

func (o SearchProfileResponsePtrOutput) OsImage() OsImageProfileResponsePtrOutput {
	return o.ApplyT(func(v *SearchProfileResponse) *OsImageProfileResponse {
		if v == nil {
			return nil
		}
		return &v.OsImage
	}).(OsImageProfileResponsePtrOutput)
}

func (o SearchProfileResponsePtrOutput) SearchType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SearchType
	}).(pulumi.StringPtrOutput)
}

type ServiceInitiatedSoftwareConfiguration struct {
	BomUrl                                string                                 `pulumi:"bomUrl"`
	HighAvailabilitySoftwareConfiguration *HighAvailabilitySoftwareConfiguration `pulumi:"highAvailabilitySoftwareConfiguration"`
	SapBitsStorageAccountId               string                                 `pulumi:"sapBitsStorageAccountId"`
	SapFqdn                               string                                 `pulumi:"sapFqdn"`
	SoftwareInstallationType              string                                 `pulumi:"softwareInstallationType"`
	SoftwareVersion                       string                                 `pulumi:"softwareVersion"`
	SshPrivateKey                         string                                 `pulumi:"sshPrivateKey"`
}

type ServiceInitiatedSoftwareConfigurationResponse struct {
	BomUrl                                string                                         `pulumi:"bomUrl"`
	HighAvailabilitySoftwareConfiguration *HighAvailabilitySoftwareConfigurationResponse `pulumi:"highAvailabilitySoftwareConfiguration"`
	SapBitsStorageAccountId               string                                         `pulumi:"sapBitsStorageAccountId"`
	SapFqdn                               string                                         `pulumi:"sapFqdn"`
	SoftwareInstallationType              string                                         `pulumi:"softwareInstallationType"`
	SoftwareVersion                       string                                         `pulumi:"softwareVersion"`
	SshPrivateKey                         string                                         `pulumi:"sshPrivateKey"`
}

type SingleServerConfiguration struct {
	AppResourceGroup            string                      `pulumi:"appResourceGroup"`
	DatabaseType                *string                     `pulumi:"databaseType"`
	DeploymentType              string                      `pulumi:"deploymentType"`
	NetworkConfiguration        *NetworkConfiguration       `pulumi:"networkConfiguration"`
	SubnetId                    string                      `pulumi:"subnetId"`
	VirtualMachineConfiguration VirtualMachineConfiguration `pulumi:"virtualMachineConfiguration"`
}


func (val *SingleServerConfiguration) Defaults() *SingleServerConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkConfiguration = tmp.NetworkConfiguration.Defaults()

	return &tmp
}

type SingleServerConfigurationResponse struct {
	AppResourceGroup            string                              `pulumi:"appResourceGroup"`
	DatabaseType                *string                             `pulumi:"databaseType"`
	DeploymentType              string                              `pulumi:"deploymentType"`
	NetworkConfiguration        *NetworkConfigurationResponse       `pulumi:"networkConfiguration"`
	SubnetId                    string                              `pulumi:"subnetId"`
	VirtualMachineConfiguration VirtualMachineConfigurationResponse `pulumi:"virtualMachineConfiguration"`
}


func (val *SingleServerConfigurationResponse) Defaults() *SingleServerConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkConfiguration = tmp.NetworkConfiguration.Defaults()

	return &tmp
}

type SiteProfile struct {
	DomainName *string `pulumi:"domainName"`
}





type SiteProfileInput interface {
	pulumi.Input

	ToSiteProfileOutput() SiteProfileOutput
	ToSiteProfileOutputWithContext(context.Context) SiteProfileOutput
}

type SiteProfileArgs struct {
	DomainName pulumi.StringPtrInput `pulumi:"domainName"`
}

func (SiteProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteProfile)(nil)).Elem()
}

func (i SiteProfileArgs) ToSiteProfileOutput() SiteProfileOutput {
	return i.ToSiteProfileOutputWithContext(context.Background())
}

func (i SiteProfileArgs) ToSiteProfileOutputWithContext(ctx context.Context) SiteProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteProfileOutput)
}

func (i SiteProfileArgs) ToSiteProfilePtrOutput() SiteProfilePtrOutput {
	return i.ToSiteProfilePtrOutputWithContext(context.Background())
}

func (i SiteProfileArgs) ToSiteProfilePtrOutputWithContext(ctx context.Context) SiteProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteProfileOutput).ToSiteProfilePtrOutputWithContext(ctx)
}









type SiteProfilePtrInput interface {
	pulumi.Input

	ToSiteProfilePtrOutput() SiteProfilePtrOutput
	ToSiteProfilePtrOutputWithContext(context.Context) SiteProfilePtrOutput
}

type siteProfilePtrType SiteProfileArgs

func SiteProfilePtr(v *SiteProfileArgs) SiteProfilePtrInput {
	return (*siteProfilePtrType)(v)
}

func (*siteProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteProfile)(nil)).Elem()
}

func (i *siteProfilePtrType) ToSiteProfilePtrOutput() SiteProfilePtrOutput {
	return i.ToSiteProfilePtrOutputWithContext(context.Background())
}

func (i *siteProfilePtrType) ToSiteProfilePtrOutputWithContext(ctx context.Context) SiteProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteProfilePtrOutput)
}

type SiteProfileOutput struct{ *pulumi.OutputState }

func (SiteProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteProfile)(nil)).Elem()
}

func (o SiteProfileOutput) ToSiteProfileOutput() SiteProfileOutput {
	return o
}

func (o SiteProfileOutput) ToSiteProfileOutputWithContext(ctx context.Context) SiteProfileOutput {
	return o
}

func (o SiteProfileOutput) ToSiteProfilePtrOutput() SiteProfilePtrOutput {
	return o.ToSiteProfilePtrOutputWithContext(context.Background())
}

func (o SiteProfileOutput) ToSiteProfilePtrOutputWithContext(ctx context.Context) SiteProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteProfile) *SiteProfile {
		return &v
	}).(SiteProfilePtrOutput)
}

func (o SiteProfileOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteProfile) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

type SiteProfilePtrOutput struct{ *pulumi.OutputState }

func (SiteProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteProfile)(nil)).Elem()
}

func (o SiteProfilePtrOutput) ToSiteProfilePtrOutput() SiteProfilePtrOutput {
	return o
}

func (o SiteProfilePtrOutput) ToSiteProfilePtrOutputWithContext(ctx context.Context) SiteProfilePtrOutput {
	return o
}

func (o SiteProfilePtrOutput) Elem() SiteProfileOutput {
	return o.ApplyT(func(v *SiteProfile) SiteProfile {
		if v != nil {
			return *v
		}
		var ret SiteProfile
		return ret
	}).(SiteProfileOutput)
}

func (o SiteProfilePtrOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteProfile) *string {
		if v == nil {
			return nil
		}
		return v.DomainName
	}).(pulumi.StringPtrOutput)
}

type SiteProfileResponse struct {
	DomainName *string `pulumi:"domainName"`
}

type SiteProfileResponseOutput struct{ *pulumi.OutputState }

func (SiteProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteProfileResponse)(nil)).Elem()
}

func (o SiteProfileResponseOutput) ToSiteProfileResponseOutput() SiteProfileResponseOutput {
	return o
}

func (o SiteProfileResponseOutput) ToSiteProfileResponseOutputWithContext(ctx context.Context) SiteProfileResponseOutput {
	return o
}

func (o SiteProfileResponseOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteProfileResponse) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

type SiteProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (SiteProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteProfileResponse)(nil)).Elem()
}

func (o SiteProfileResponsePtrOutput) ToSiteProfileResponsePtrOutput() SiteProfileResponsePtrOutput {
	return o
}

func (o SiteProfileResponsePtrOutput) ToSiteProfileResponsePtrOutputWithContext(ctx context.Context) SiteProfileResponsePtrOutput {
	return o
}

func (o SiteProfileResponsePtrOutput) Elem() SiteProfileResponseOutput {
	return o.ApplyT(func(v *SiteProfileResponse) SiteProfileResponse {
		if v != nil {
			return *v
		}
		var ret SiteProfileResponse
		return ret
	}).(SiteProfileResponseOutput)
}

func (o SiteProfileResponsePtrOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DomainName
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *int     `pulumi:"capacity"`
	Family   *string  `pulumi:"family"`
	Name     string   `pulumi:"name"`
	Size     *string  `pulumi:"size"`
	Tier     *SkuTier `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     SkuTierPtrInput       `pulumi:"tier"`
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

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v Sku) *SkuTier { return v.Tier }).(SkuTierPtrOutput)
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

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v *Sku) *SkuTier {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(SkuTierPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SshConfiguration struct {
	PublicKeys []SshPublicKey `pulumi:"publicKeys"`
}

type SshConfigurationResponse struct {
	PublicKeys []SshPublicKeyResponse `pulumi:"publicKeys"`
}

type SshKeyPair struct {
	PrivateKey *string `pulumi:"privateKey"`
	PublicKey  *string `pulumi:"publicKey"`
}

type SshKeyPairResponse struct {
	PrivateKey *string `pulumi:"privateKey"`
	PublicKey  *string `pulumi:"publicKey"`
}

type SshPublicKey struct {
	KeyData *string `pulumi:"keyData"`
}

type SshPublicKeyResponse struct {
	KeyData *string `pulumi:"keyData"`
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
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

type ThreeTierConfiguration struct {
	AppResourceGroup       string                         `pulumi:"appResourceGroup"`
	ApplicationServer      ApplicationServerConfiguration `pulumi:"applicationServer"`
	CentralServer          CentralServerConfiguration     `pulumi:"centralServer"`
	DatabaseServer         DatabaseConfiguration          `pulumi:"databaseServer"`
	DeploymentType         string                         `pulumi:"deploymentType"`
	HighAvailabilityConfig *HighAvailabilityConfiguration `pulumi:"highAvailabilityConfig"`
	NetworkConfiguration   *NetworkConfiguration          `pulumi:"networkConfiguration"`
}


func (val *ThreeTierConfiguration) Defaults() *ThreeTierConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkConfiguration = tmp.NetworkConfiguration.Defaults()

	return &tmp
}

type ThreeTierConfigurationResponse struct {
	AppResourceGroup       string                                 `pulumi:"appResourceGroup"`
	ApplicationServer      ApplicationServerConfigurationResponse `pulumi:"applicationServer"`
	CentralServer          CentralServerConfigurationResponse     `pulumi:"centralServer"`
	DatabaseServer         DatabaseConfigurationResponse          `pulumi:"databaseServer"`
	DeploymentType         string                                 `pulumi:"deploymentType"`
	HighAvailabilityConfig *HighAvailabilityConfigurationResponse `pulumi:"highAvailabilityConfig"`
	NetworkConfiguration   *NetworkConfigurationResponse          `pulumi:"networkConfiguration"`
}


func (val *ThreeTierConfigurationResponse) Defaults() *ThreeTierConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkConfiguration = tmp.NetworkConfiguration.Defaults()

	return &tmp
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type UserAssignedServiceIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type UserAssignedServiceIdentityInput interface {
	pulumi.Input

	ToUserAssignedServiceIdentityOutput() UserAssignedServiceIdentityOutput
	ToUserAssignedServiceIdentityOutputWithContext(context.Context) UserAssignedServiceIdentityOutput
}

type UserAssignedServiceIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (UserAssignedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedServiceIdentity)(nil)).Elem()
}

func (i UserAssignedServiceIdentityArgs) ToUserAssignedServiceIdentityOutput() UserAssignedServiceIdentityOutput {
	return i.ToUserAssignedServiceIdentityOutputWithContext(context.Background())
}

func (i UserAssignedServiceIdentityArgs) ToUserAssignedServiceIdentityOutputWithContext(ctx context.Context) UserAssignedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedServiceIdentityOutput)
}

func (i UserAssignedServiceIdentityArgs) ToUserAssignedServiceIdentityPtrOutput() UserAssignedServiceIdentityPtrOutput {
	return i.ToUserAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i UserAssignedServiceIdentityArgs) ToUserAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) UserAssignedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedServiceIdentityOutput).ToUserAssignedServiceIdentityPtrOutputWithContext(ctx)
}









type UserAssignedServiceIdentityPtrInput interface {
	pulumi.Input

	ToUserAssignedServiceIdentityPtrOutput() UserAssignedServiceIdentityPtrOutput
	ToUserAssignedServiceIdentityPtrOutputWithContext(context.Context) UserAssignedServiceIdentityPtrOutput
}

type userAssignedServiceIdentityPtrType UserAssignedServiceIdentityArgs

func UserAssignedServiceIdentityPtr(v *UserAssignedServiceIdentityArgs) UserAssignedServiceIdentityPtrInput {
	return (*userAssignedServiceIdentityPtrType)(v)
}

func (*userAssignedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedServiceIdentity)(nil)).Elem()
}

func (i *userAssignedServiceIdentityPtrType) ToUserAssignedServiceIdentityPtrOutput() UserAssignedServiceIdentityPtrOutput {
	return i.ToUserAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *userAssignedServiceIdentityPtrType) ToUserAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) UserAssignedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedServiceIdentityPtrOutput)
}

type UserAssignedServiceIdentityOutput struct{ *pulumi.OutputState }

func (UserAssignedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedServiceIdentity)(nil)).Elem()
}

func (o UserAssignedServiceIdentityOutput) ToUserAssignedServiceIdentityOutput() UserAssignedServiceIdentityOutput {
	return o
}

func (o UserAssignedServiceIdentityOutput) ToUserAssignedServiceIdentityOutputWithContext(ctx context.Context) UserAssignedServiceIdentityOutput {
	return o
}

func (o UserAssignedServiceIdentityOutput) ToUserAssignedServiceIdentityPtrOutput() UserAssignedServiceIdentityPtrOutput {
	return o.ToUserAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o UserAssignedServiceIdentityOutput) ToUserAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) UserAssignedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserAssignedServiceIdentity) *UserAssignedServiceIdentity {
		return &v
	}).(UserAssignedServiceIdentityPtrOutput)
}

func (o UserAssignedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o UserAssignedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v UserAssignedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type UserAssignedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (UserAssignedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedServiceIdentity)(nil)).Elem()
}

func (o UserAssignedServiceIdentityPtrOutput) ToUserAssignedServiceIdentityPtrOutput() UserAssignedServiceIdentityPtrOutput {
	return o
}

func (o UserAssignedServiceIdentityPtrOutput) ToUserAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) UserAssignedServiceIdentityPtrOutput {
	return o
}

func (o UserAssignedServiceIdentityPtrOutput) Elem() UserAssignedServiceIdentityOutput {
	return o.ApplyT(func(v *UserAssignedServiceIdentity) UserAssignedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret UserAssignedServiceIdentity
		return ret
	}).(UserAssignedServiceIdentityOutput)
}

func (o UserAssignedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAssignedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o UserAssignedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *UserAssignedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type UserAssignedServiceIdentityResponse struct {
	Type                   string                                  `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type UserAssignedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedServiceIdentityResponse)(nil)).Elem()
}

func (o UserAssignedServiceIdentityResponseOutput) ToUserAssignedServiceIdentityResponseOutput() UserAssignedServiceIdentityResponseOutput {
	return o
}

func (o UserAssignedServiceIdentityResponseOutput) ToUserAssignedServiceIdentityResponseOutputWithContext(ctx context.Context) UserAssignedServiceIdentityResponseOutput {
	return o
}

func (o UserAssignedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o UserAssignedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v UserAssignedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type UserAssignedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (UserAssignedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedServiceIdentityResponse)(nil)).Elem()
}

func (o UserAssignedServiceIdentityResponsePtrOutput) ToUserAssignedServiceIdentityResponsePtrOutput() UserAssignedServiceIdentityResponsePtrOutput {
	return o
}

func (o UserAssignedServiceIdentityResponsePtrOutput) ToUserAssignedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) UserAssignedServiceIdentityResponsePtrOutput {
	return o
}

func (o UserAssignedServiceIdentityResponsePtrOutput) Elem() UserAssignedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *UserAssignedServiceIdentityResponse) UserAssignedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret UserAssignedServiceIdentityResponse
		return ret
	}).(UserAssignedServiceIdentityResponseOutput)
}

func (o UserAssignedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o UserAssignedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *UserAssignedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type UserProfile struct {
	SshPublicKey string `pulumi:"sshPublicKey"`
	UserName     string `pulumi:"userName"`
}





type UserProfileInput interface {
	pulumi.Input

	ToUserProfileOutput() UserProfileOutput
	ToUserProfileOutputWithContext(context.Context) UserProfileOutput
}

type UserProfileArgs struct {
	SshPublicKey pulumi.StringInput `pulumi:"sshPublicKey"`
	UserName     pulumi.StringInput `pulumi:"userName"`
}

func (UserProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProfile)(nil)).Elem()
}

func (i UserProfileArgs) ToUserProfileOutput() UserProfileOutput {
	return i.ToUserProfileOutputWithContext(context.Background())
}

func (i UserProfileArgs) ToUserProfileOutputWithContext(ctx context.Context) UserProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserProfileOutput)
}

type UserProfileOutput struct{ *pulumi.OutputState }

func (UserProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProfile)(nil)).Elem()
}

func (o UserProfileOutput) ToUserProfileOutput() UserProfileOutput {
	return o
}

func (o UserProfileOutput) ToUserProfileOutputWithContext(ctx context.Context) UserProfileOutput {
	return o
}

func (o UserProfileOutput) SshPublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v UserProfile) string { return v.SshPublicKey }).(pulumi.StringOutput)
}

func (o UserProfileOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v UserProfile) string { return v.UserName }).(pulumi.StringOutput)
}

type UserProfileResponse struct {
	SshPublicKey string `pulumi:"sshPublicKey"`
	UserName     string `pulumi:"userName"`
}

type UserProfileResponseOutput struct{ *pulumi.OutputState }

func (UserProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProfileResponse)(nil)).Elem()
}

func (o UserProfileResponseOutput) ToUserProfileResponseOutput() UserProfileResponseOutput {
	return o
}

func (o UserProfileResponseOutput) ToUserProfileResponseOutputWithContext(ctx context.Context) UserProfileResponseOutput {
	return o
}

func (o UserProfileResponseOutput) SshPublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v UserProfileResponse) string { return v.SshPublicKey }).(pulumi.StringOutput)
}

func (o UserProfileResponseOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v UserProfileResponse) string { return v.UserName }).(pulumi.StringOutput)
}

type VirtualMachineConfiguration struct {
	ImageReference ImageReference `pulumi:"imageReference"`
	OsProfile      OSProfile      `pulumi:"osProfile"`
	VmSize         string         `pulumi:"vmSize"`
}

type VirtualMachineConfigurationResponse struct {
	ImageReference ImageReferenceResponse `pulumi:"imageReference"`
	OsProfile      OSProfileResponse      `pulumi:"osProfile"`
	VmSize         string                 `pulumi:"vmSize"`
}

type VmssNodesProfile struct {
	AutoScaleMaxCount *int           `pulumi:"autoScaleMaxCount"`
	AutoScaleMinCount *int           `pulumi:"autoScaleMinCount"`
	DataDisks         []DiskInfo     `pulumi:"dataDisks"`
	Name              *string        `pulumi:"name"`
	NodeSku           string         `pulumi:"nodeSku"`
	OsDisk            DiskInfo       `pulumi:"osDisk"`
	OsImage           OsImageProfile `pulumi:"osImage"`
}





type VmssNodesProfileInput interface {
	pulumi.Input

	ToVmssNodesProfileOutput() VmssNodesProfileOutput
	ToVmssNodesProfileOutputWithContext(context.Context) VmssNodesProfileOutput
}

type VmssNodesProfileArgs struct {
	AutoScaleMaxCount pulumi.IntPtrInput    `pulumi:"autoScaleMaxCount"`
	AutoScaleMinCount pulumi.IntPtrInput    `pulumi:"autoScaleMinCount"`
	DataDisks         DiskInfoArrayInput    `pulumi:"dataDisks"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	NodeSku           pulumi.StringInput    `pulumi:"nodeSku"`
	OsDisk            DiskInfoInput         `pulumi:"osDisk"`
	OsImage           OsImageProfileInput   `pulumi:"osImage"`
}

func (VmssNodesProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmssNodesProfile)(nil)).Elem()
}

func (i VmssNodesProfileArgs) ToVmssNodesProfileOutput() VmssNodesProfileOutput {
	return i.ToVmssNodesProfileOutputWithContext(context.Background())
}

func (i VmssNodesProfileArgs) ToVmssNodesProfileOutputWithContext(ctx context.Context) VmssNodesProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmssNodesProfileOutput)
}

type VmssNodesProfileOutput struct{ *pulumi.OutputState }

func (VmssNodesProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmssNodesProfile)(nil)).Elem()
}

func (o VmssNodesProfileOutput) ToVmssNodesProfileOutput() VmssNodesProfileOutput {
	return o
}

func (o VmssNodesProfileOutput) ToVmssNodesProfileOutputWithContext(ctx context.Context) VmssNodesProfileOutput {
	return o
}

func (o VmssNodesProfileOutput) AutoScaleMaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VmssNodesProfile) *int { return v.AutoScaleMaxCount }).(pulumi.IntPtrOutput)
}

func (o VmssNodesProfileOutput) AutoScaleMinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VmssNodesProfile) *int { return v.AutoScaleMinCount }).(pulumi.IntPtrOutput)
}

func (o VmssNodesProfileOutput) DataDisks() DiskInfoArrayOutput {
	return o.ApplyT(func(v VmssNodesProfile) []DiskInfo { return v.DataDisks }).(DiskInfoArrayOutput)
}

func (o VmssNodesProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmssNodesProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VmssNodesProfileOutput) NodeSku() pulumi.StringOutput {
	return o.ApplyT(func(v VmssNodesProfile) string { return v.NodeSku }).(pulumi.StringOutput)
}

func (o VmssNodesProfileOutput) OsDisk() DiskInfoOutput {
	return o.ApplyT(func(v VmssNodesProfile) DiskInfo { return v.OsDisk }).(DiskInfoOutput)
}

func (o VmssNodesProfileOutput) OsImage() OsImageProfileOutput {
	return o.ApplyT(func(v VmssNodesProfile) OsImageProfile { return v.OsImage }).(OsImageProfileOutput)
}

type VmssNodesProfileResponse struct {
	AutoScaleMaxCount *int                   `pulumi:"autoScaleMaxCount"`
	AutoScaleMinCount *int                   `pulumi:"autoScaleMinCount"`
	DataDisks         []DiskInfoResponse     `pulumi:"dataDisks"`
	Name              *string                `pulumi:"name"`
	NodeResourceIds   []string               `pulumi:"nodeResourceIds"`
	NodeSku           string                 `pulumi:"nodeSku"`
	OsDisk            DiskInfoResponse       `pulumi:"osDisk"`
	OsImage           OsImageProfileResponse `pulumi:"osImage"`
}

type VmssNodesProfileResponseOutput struct{ *pulumi.OutputState }

func (VmssNodesProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmssNodesProfileResponse)(nil)).Elem()
}

func (o VmssNodesProfileResponseOutput) ToVmssNodesProfileResponseOutput() VmssNodesProfileResponseOutput {
	return o
}

func (o VmssNodesProfileResponseOutput) ToVmssNodesProfileResponseOutputWithContext(ctx context.Context) VmssNodesProfileResponseOutput {
	return o
}

func (o VmssNodesProfileResponseOutput) AutoScaleMaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VmssNodesProfileResponse) *int { return v.AutoScaleMaxCount }).(pulumi.IntPtrOutput)
}

func (o VmssNodesProfileResponseOutput) AutoScaleMinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VmssNodesProfileResponse) *int { return v.AutoScaleMinCount }).(pulumi.IntPtrOutput)
}

func (o VmssNodesProfileResponseOutput) DataDisks() DiskInfoResponseArrayOutput {
	return o.ApplyT(func(v VmssNodesProfileResponse) []DiskInfoResponse { return v.DataDisks }).(DiskInfoResponseArrayOutput)
}

func (o VmssNodesProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmssNodesProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VmssNodesProfileResponseOutput) NodeResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmssNodesProfileResponse) []string { return v.NodeResourceIds }).(pulumi.StringArrayOutput)
}

func (o VmssNodesProfileResponseOutput) NodeSku() pulumi.StringOutput {
	return o.ApplyT(func(v VmssNodesProfileResponse) string { return v.NodeSku }).(pulumi.StringOutput)
}

func (o VmssNodesProfileResponseOutput) OsDisk() DiskInfoResponseOutput {
	return o.ApplyT(func(v VmssNodesProfileResponse) DiskInfoResponse { return v.OsDisk }).(DiskInfoResponseOutput)
}

func (o VmssNodesProfileResponseOutput) OsImage() OsImageProfileResponseOutput {
	return o.ApplyT(func(v VmssNodesProfileResponse) OsImageProfileResponse { return v.OsImage }).(OsImageProfileResponseOutput)
}

type WindowsConfiguration struct {
	OsType string `pulumi:"osType"`
}

type WindowsConfigurationResponse struct {
	OsType string `pulumi:"osType"`
}

func init() {
	pulumi.RegisterOutputType(BackupProfileOutput{})
	pulumi.RegisterOutputType(BackupProfilePtrOutput{})
	pulumi.RegisterOutputType(BackupProfileResponseOutput{})
	pulumi.RegisterOutputType(BackupProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(CacheProfileOutput{})
	pulumi.RegisterOutputType(CacheProfilePtrOutput{})
	pulumi.RegisterOutputType(CacheProfileResponseOutput{})
	pulumi.RegisterOutputType(CacheProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(CentralServerVmDetailsResponseOutput{})
	pulumi.RegisterOutputType(CentralServerVmDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseProfileOutput{})
	pulumi.RegisterOutputType(DatabaseProfileResponseOutput{})
	pulumi.RegisterOutputType(DatabaseVmDetailsResponseOutput{})
	pulumi.RegisterOutputType(DatabaseVmDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskInfoOutput{})
	pulumi.RegisterOutputType(DiskInfoPtrOutput{})
	pulumi.RegisterOutputType(DiskInfoArrayOutput{})
	pulumi.RegisterOutputType(DiskInfoResponseOutput{})
	pulumi.RegisterOutputType(DiskInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(EnqueueReplicationServerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EnqueueReplicationServerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EnqueueServerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EnqueueServerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ErrorDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorResponseInnerErrorOutput{})
	pulumi.RegisterOutputType(ErrorResponseInnerErrorPtrOutput{})
	pulumi.RegisterOutputType(FileshareProfileOutput{})
	pulumi.RegisterOutputType(FileshareProfilePtrOutput{})
	pulumi.RegisterOutputType(FileshareProfileResponseOutput{})
	pulumi.RegisterOutputType(FileshareProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(GatewayServerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GatewayServerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedRGConfigurationOutput{})
	pulumi.RegisterOutputType(ManagedRGConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ManagedRGConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ManagedRGConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(MessageServerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MessageServerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesResponseErrorsOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(NodeProfileOutput{})
	pulumi.RegisterOutputType(NodeProfileResponseOutput{})
	pulumi.RegisterOutputType(OsImageProfileOutput{})
	pulumi.RegisterOutputType(OsImageProfilePtrOutput{})
	pulumi.RegisterOutputType(OsImageProfileResponseOutput{})
	pulumi.RegisterOutputType(OsImageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(PhpProfileOutput{})
	pulumi.RegisterOutputType(PhpProfilePtrOutput{})
	pulumi.RegisterOutputType(PhpProfileResponseOutput{})
	pulumi.RegisterOutputType(PhpProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(PhpWorkloadResourceIdentityOutput{})
	pulumi.RegisterOutputType(PhpWorkloadResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(PhpWorkloadResourceResponseIdentityOutput{})
	pulumi.RegisterOutputType(PhpWorkloadResourceResponseIdentityPtrOutput{})
	pulumi.RegisterOutputType(ProviderInstancePropertiesResponseErrorsOutput{})
	pulumi.RegisterOutputType(SAPAvailabilityZonePairResponseOutput{})
	pulumi.RegisterOutputType(SAPAvailabilityZonePairResponseArrayOutput{})
	pulumi.RegisterOutputType(SAPDiskConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SAPDiskConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(SAPSupportedSkuResponseOutput{})
	pulumi.RegisterOutputType(SAPSupportedSkuResponseArrayOutput{})
	pulumi.RegisterOutputType(SAPVirtualInstanceErrorResponseOutput{})
	pulumi.RegisterOutputType(SearchProfileOutput{})
	pulumi.RegisterOutputType(SearchProfilePtrOutput{})
	pulumi.RegisterOutputType(SearchProfileResponseOutput{})
	pulumi.RegisterOutputType(SearchProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SiteProfileOutput{})
	pulumi.RegisterOutputType(SiteProfilePtrOutput{})
	pulumi.RegisterOutputType(SiteProfileResponseOutput{})
	pulumi.RegisterOutputType(SiteProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(UserAssignedServiceIdentityOutput{})
	pulumi.RegisterOutputType(UserAssignedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(UserAssignedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(UserProfileOutput{})
	pulumi.RegisterOutputType(UserProfileResponseOutput{})
	pulumi.RegisterOutputType(VmssNodesProfileOutput{})
	pulumi.RegisterOutputType(VmssNodesProfileResponseOutput{})
}

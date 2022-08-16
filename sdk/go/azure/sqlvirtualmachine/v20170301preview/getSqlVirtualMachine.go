


package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlVirtualMachine(ctx *pulumi.Context, args *LookupSqlVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupSqlVirtualMachineResult, error) {
	var rv LookupSqlVirtualMachineResult
	err := ctx.Invoke("azure-native:sqlvirtualmachine/v20170301preview:getSqlVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlVirtualMachineArgs struct {
	Expand                *string `pulumi:"expand"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	SqlVirtualMachineName string  `pulumi:"sqlVirtualMachineName"`
}


type LookupSqlVirtualMachineResult struct {
	AutoBackupSettings                     *AutoBackupSettingsResponse                     `pulumi:"autoBackupSettings"`
	AutoPatchingSettings                   *AutoPatchingSettingsResponse                   `pulumi:"autoPatchingSettings"`
	Id                                     string                                          `pulumi:"id"`
	Identity                               *ResourceIdentityResponse                       `pulumi:"identity"`
	KeyVaultCredentialSettings             *KeyVaultCredentialSettingsResponse             `pulumi:"keyVaultCredentialSettings"`
	Location                               string                                          `pulumi:"location"`
	Name                                   string                                          `pulumi:"name"`
	ProvisioningState                      string                                          `pulumi:"provisioningState"`
	ServerConfigurationsManagementSettings *ServerConfigurationsManagementSettingsResponse `pulumi:"serverConfigurationsManagementSettings"`
	SqlImageOffer                          *string                                         `pulumi:"sqlImageOffer"`
	SqlImageSku                            *string                                         `pulumi:"sqlImageSku"`
	SqlManagement                          *string                                         `pulumi:"sqlManagement"`
	SqlServerLicenseType                   *string                                         `pulumi:"sqlServerLicenseType"`
	SqlVirtualMachineGroupResourceId       *string                                         `pulumi:"sqlVirtualMachineGroupResourceId"`
	StorageConfigurationSettings           *StorageConfigurationSettingsResponse           `pulumi:"storageConfigurationSettings"`
	Tags                                   map[string]string                               `pulumi:"tags"`
	Type                                   string                                          `pulumi:"type"`
	VirtualMachineResourceId               *string                                         `pulumi:"virtualMachineResourceId"`
	WsfcDomainCredentials                  *WsfcDomainCredentialsResponse                  `pulumi:"wsfcDomainCredentials"`
}

func LookupSqlVirtualMachineOutput(ctx *pulumi.Context, args LookupSqlVirtualMachineOutputArgs, opts ...pulumi.InvokeOption) LookupSqlVirtualMachineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlVirtualMachineResult, error) {
			args := v.(LookupSqlVirtualMachineArgs)
			r, err := LookupSqlVirtualMachine(ctx, &args, opts...)
			var s LookupSqlVirtualMachineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlVirtualMachineResultOutput)
}

type LookupSqlVirtualMachineOutputArgs struct {
	Expand                pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName     pulumi.StringInput    `pulumi:"resourceGroupName"`
	SqlVirtualMachineName pulumi.StringInput    `pulumi:"sqlVirtualMachineName"`
}

func (LookupSqlVirtualMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlVirtualMachineArgs)(nil)).Elem()
}


type LookupSqlVirtualMachineResultOutput struct{ *pulumi.OutputState }

func (LookupSqlVirtualMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlVirtualMachineResult)(nil)).Elem()
}

func (o LookupSqlVirtualMachineResultOutput) ToLookupSqlVirtualMachineResultOutput() LookupSqlVirtualMachineResultOutput {
	return o
}

func (o LookupSqlVirtualMachineResultOutput) ToLookupSqlVirtualMachineResultOutputWithContext(ctx context.Context) LookupSqlVirtualMachineResultOutput {
	return o
}

func (o LookupSqlVirtualMachineResultOutput) AutoBackupSettings() AutoBackupSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *AutoBackupSettingsResponse { return v.AutoBackupSettings }).(AutoBackupSettingsResponsePtrOutput)
}

func (o LookupSqlVirtualMachineResultOutput) AutoPatchingSettings() AutoPatchingSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *AutoPatchingSettingsResponse { return v.AutoPatchingSettings }).(AutoPatchingSettingsResponsePtrOutput)
}

func (o LookupSqlVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o LookupSqlVirtualMachineResultOutput) KeyVaultCredentialSettings() KeyVaultCredentialSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *KeyVaultCredentialSettingsResponse {
		return v.KeyVaultCredentialSettings
	}).(KeyVaultCredentialSettingsResponsePtrOutput)
}

func (o LookupSqlVirtualMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineResultOutput) ServerConfigurationsManagementSettings() ServerConfigurationsManagementSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *ServerConfigurationsManagementSettingsResponse {
		return v.ServerConfigurationsManagementSettings
	}).(ServerConfigurationsManagementSettingsResponsePtrOutput)
}

func (o LookupSqlVirtualMachineResultOutput) SqlImageOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.SqlImageOffer }).(pulumi.StringPtrOutput)
}

func (o LookupSqlVirtualMachineResultOutput) SqlImageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.SqlImageSku }).(pulumi.StringPtrOutput)
}

func (o LookupSqlVirtualMachineResultOutput) SqlManagement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.SqlManagement }).(pulumi.StringPtrOutput)
}

func (o LookupSqlVirtualMachineResultOutput) SqlServerLicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.SqlServerLicenseType }).(pulumi.StringPtrOutput)
}

func (o LookupSqlVirtualMachineResultOutput) SqlVirtualMachineGroupResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.SqlVirtualMachineGroupResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupSqlVirtualMachineResultOutput) StorageConfigurationSettings() StorageConfigurationSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *StorageConfigurationSettingsResponse {
		return v.StorageConfigurationSettings
	}).(StorageConfigurationSettingsResponsePtrOutput)
}

func (o LookupSqlVirtualMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlVirtualMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineResultOutput) VirtualMachineResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.VirtualMachineResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupSqlVirtualMachineResultOutput) WsfcDomainCredentials() WsfcDomainCredentialsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *WsfcDomainCredentialsResponse { return v.WsfcDomainCredentials }).(WsfcDomainCredentialsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlVirtualMachineResultOutput{})
}

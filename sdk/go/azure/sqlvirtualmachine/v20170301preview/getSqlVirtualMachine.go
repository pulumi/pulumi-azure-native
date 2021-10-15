


package v20170301preview

import (
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

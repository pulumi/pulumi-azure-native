


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlVirtualMachine struct {
	pulumi.CustomResourceState

	AssessmentSettings                     AssessmentSettingsResponsePtrOutput                     `pulumi:"assessmentSettings"`
	AutoBackupSettings                     AutoBackupSettingsResponsePtrOutput                     `pulumi:"autoBackupSettings"`
	AutoPatchingSettings                   AutoPatchingSettingsResponsePtrOutput                   `pulumi:"autoPatchingSettings"`
	EnableAutomaticUpgrade                 pulumi.BoolPtrOutput                                    `pulumi:"enableAutomaticUpgrade"`
	Identity                               ResourceIdentityResponsePtrOutput                       `pulumi:"identity"`
	KeyVaultCredentialSettings             KeyVaultCredentialSettingsResponsePtrOutput             `pulumi:"keyVaultCredentialSettings"`
	LeastPrivilegeMode                     pulumi.StringPtrOutput                                  `pulumi:"leastPrivilegeMode"`
	Location                               pulumi.StringOutput                                     `pulumi:"location"`
	Name                                   pulumi.StringOutput                                     `pulumi:"name"`
	ProvisioningState                      pulumi.StringOutput                                     `pulumi:"provisioningState"`
	ServerConfigurationsManagementSettings ServerConfigurationsManagementSettingsResponsePtrOutput `pulumi:"serverConfigurationsManagementSettings"`
	SqlImageOffer                          pulumi.StringPtrOutput                                  `pulumi:"sqlImageOffer"`
	SqlImageSku                            pulumi.StringPtrOutput                                  `pulumi:"sqlImageSku"`
	SqlManagement                          pulumi.StringPtrOutput                                  `pulumi:"sqlManagement"`
	SqlServerLicenseType                   pulumi.StringPtrOutput                                  `pulumi:"sqlServerLicenseType"`
	SqlVirtualMachineGroupResourceId       pulumi.StringPtrOutput                                  `pulumi:"sqlVirtualMachineGroupResourceId"`
	StorageConfigurationSettings           StorageConfigurationSettingsResponsePtrOutput           `pulumi:"storageConfigurationSettings"`
	SystemData                             SystemDataResponseOutput                                `pulumi:"systemData"`
	Tags                                   pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type                                   pulumi.StringOutput                                     `pulumi:"type"`
	VirtualMachineResourceId               pulumi.StringPtrOutput                                  `pulumi:"virtualMachineResourceId"`
	WsfcDomainCredentials                  WsfcDomainCredentialsResponsePtrOutput                  `pulumi:"wsfcDomainCredentials"`
	WsfcStaticIp                           pulumi.StringPtrOutput                                  `pulumi:"wsfcStaticIp"`
}


func NewSqlVirtualMachine(ctx *pulumi.Context,
	name string, args *SqlVirtualMachineArgs, opts ...pulumi.ResourceOption) (*SqlVirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.EnableAutomaticUpgrade) {
		args.EnableAutomaticUpgrade = pulumi.BoolPtr(false)
	}
	if isZero(args.LeastPrivilegeMode) {
		args.LeastPrivilegeMode = pulumi.StringPtr("NotSet")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine:SqlVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20170301preview:SqlVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20211101preview:SqlVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220201:SqlVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220201preview:SqlVirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlVirtualMachine
	err := ctx.RegisterResource("azure-native:sqlvirtualmachine/v20220701preview:SqlVirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlVirtualMachineState, opts ...pulumi.ResourceOption) (*SqlVirtualMachine, error) {
	var resource SqlVirtualMachine
	err := ctx.ReadResource("azure-native:sqlvirtualmachine/v20220701preview:SqlVirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlVirtualMachineState struct {
}

type SqlVirtualMachineState struct {
}

func (SqlVirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlVirtualMachineState)(nil)).Elem()
}

type sqlVirtualMachineArgs struct {
	AssessmentSettings                     *AssessmentSettings                     `pulumi:"assessmentSettings"`
	AutoBackupSettings                     *AutoBackupSettings                     `pulumi:"autoBackupSettings"`
	AutoPatchingSettings                   *AutoPatchingSettings                   `pulumi:"autoPatchingSettings"`
	EnableAutomaticUpgrade                 *bool                                   `pulumi:"enableAutomaticUpgrade"`
	Identity                               *ResourceIdentity                       `pulumi:"identity"`
	KeyVaultCredentialSettings             *KeyVaultCredentialSettings             `pulumi:"keyVaultCredentialSettings"`
	LeastPrivilegeMode                     *string                                 `pulumi:"leastPrivilegeMode"`
	Location                               *string                                 `pulumi:"location"`
	ResourceGroupName                      string                                  `pulumi:"resourceGroupName"`
	ServerConfigurationsManagementSettings *ServerConfigurationsManagementSettings `pulumi:"serverConfigurationsManagementSettings"`
	SqlImageOffer                          *string                                 `pulumi:"sqlImageOffer"`
	SqlImageSku                            *string                                 `pulumi:"sqlImageSku"`
	SqlManagement                          *string                                 `pulumi:"sqlManagement"`
	SqlServerLicenseType                   *string                                 `pulumi:"sqlServerLicenseType"`
	SqlVirtualMachineGroupResourceId       *string                                 `pulumi:"sqlVirtualMachineGroupResourceId"`
	SqlVirtualMachineName                  *string                                 `pulumi:"sqlVirtualMachineName"`
	StorageConfigurationSettings           *StorageConfigurationSettings           `pulumi:"storageConfigurationSettings"`
	Tags                                   map[string]string                       `pulumi:"tags"`
	VirtualMachineResourceId               *string                                 `pulumi:"virtualMachineResourceId"`
	WsfcDomainCredentials                  *WsfcDomainCredentials                  `pulumi:"wsfcDomainCredentials"`
	WsfcStaticIp                           *string                                 `pulumi:"wsfcStaticIp"`
}


type SqlVirtualMachineArgs struct {
	AssessmentSettings                     AssessmentSettingsPtrInput
	AutoBackupSettings                     AutoBackupSettingsPtrInput
	AutoPatchingSettings                   AutoPatchingSettingsPtrInput
	EnableAutomaticUpgrade                 pulumi.BoolPtrInput
	Identity                               ResourceIdentityPtrInput
	KeyVaultCredentialSettings             KeyVaultCredentialSettingsPtrInput
	LeastPrivilegeMode                     pulumi.StringPtrInput
	Location                               pulumi.StringPtrInput
	ResourceGroupName                      pulumi.StringInput
	ServerConfigurationsManagementSettings ServerConfigurationsManagementSettingsPtrInput
	SqlImageOffer                          pulumi.StringPtrInput
	SqlImageSku                            pulumi.StringPtrInput
	SqlManagement                          pulumi.StringPtrInput
	SqlServerLicenseType                   pulumi.StringPtrInput
	SqlVirtualMachineGroupResourceId       pulumi.StringPtrInput
	SqlVirtualMachineName                  pulumi.StringPtrInput
	StorageConfigurationSettings           StorageConfigurationSettingsPtrInput
	Tags                                   pulumi.StringMapInput
	VirtualMachineResourceId               pulumi.StringPtrInput
	WsfcDomainCredentials                  WsfcDomainCredentialsPtrInput
	WsfcStaticIp                           pulumi.StringPtrInput
}

func (SqlVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlVirtualMachineArgs)(nil)).Elem()
}

type SqlVirtualMachineInput interface {
	pulumi.Input

	ToSqlVirtualMachineOutput() SqlVirtualMachineOutput
	ToSqlVirtualMachineOutputWithContext(ctx context.Context) SqlVirtualMachineOutput
}

func (*SqlVirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlVirtualMachine)(nil)).Elem()
}

func (i *SqlVirtualMachine) ToSqlVirtualMachineOutput() SqlVirtualMachineOutput {
	return i.ToSqlVirtualMachineOutputWithContext(context.Background())
}

func (i *SqlVirtualMachine) ToSqlVirtualMachineOutputWithContext(ctx context.Context) SqlVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlVirtualMachineOutput)
}

type SqlVirtualMachineOutput struct{ *pulumi.OutputState }

func (SqlVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlVirtualMachine)(nil)).Elem()
}

func (o SqlVirtualMachineOutput) ToSqlVirtualMachineOutput() SqlVirtualMachineOutput {
	return o
}

func (o SqlVirtualMachineOutput) ToSqlVirtualMachineOutputWithContext(ctx context.Context) SqlVirtualMachineOutput {
	return o
}

func (o SqlVirtualMachineOutput) AssessmentSettings() AssessmentSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) AssessmentSettingsResponsePtrOutput { return v.AssessmentSettings }).(AssessmentSettingsResponsePtrOutput)
}

func (o SqlVirtualMachineOutput) AutoBackupSettings() AutoBackupSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) AutoBackupSettingsResponsePtrOutput { return v.AutoBackupSettings }).(AutoBackupSettingsResponsePtrOutput)
}

func (o SqlVirtualMachineOutput) AutoPatchingSettings() AutoPatchingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) AutoPatchingSettingsResponsePtrOutput { return v.AutoPatchingSettings }).(AutoPatchingSettingsResponsePtrOutput)
}

func (o SqlVirtualMachineOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.BoolPtrOutput { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

func (o SqlVirtualMachineOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o SqlVirtualMachineOutput) KeyVaultCredentialSettings() KeyVaultCredentialSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) KeyVaultCredentialSettingsResponsePtrOutput {
		return v.KeyVaultCredentialSettings
	}).(KeyVaultCredentialSettingsResponsePtrOutput)
}

func (o SqlVirtualMachineOutput) LeastPrivilegeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.LeastPrivilegeMode }).(pulumi.StringPtrOutput)
}

func (o SqlVirtualMachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SqlVirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlVirtualMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SqlVirtualMachineOutput) ServerConfigurationsManagementSettings() ServerConfigurationsManagementSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) ServerConfigurationsManagementSettingsResponsePtrOutput {
		return v.ServerConfigurationsManagementSettings
	}).(ServerConfigurationsManagementSettingsResponsePtrOutput)
}

func (o SqlVirtualMachineOutput) SqlImageOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.SqlImageOffer }).(pulumi.StringPtrOutput)
}

func (o SqlVirtualMachineOutput) SqlImageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.SqlImageSku }).(pulumi.StringPtrOutput)
}

func (o SqlVirtualMachineOutput) SqlManagement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.SqlManagement }).(pulumi.StringPtrOutput)
}

func (o SqlVirtualMachineOutput) SqlServerLicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.SqlServerLicenseType }).(pulumi.StringPtrOutput)
}

func (o SqlVirtualMachineOutput) SqlVirtualMachineGroupResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.SqlVirtualMachineGroupResourceId }).(pulumi.StringPtrOutput)
}

func (o SqlVirtualMachineOutput) StorageConfigurationSettings() StorageConfigurationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) StorageConfigurationSettingsResponsePtrOutput {
		return v.StorageConfigurationSettings
	}).(StorageConfigurationSettingsResponsePtrOutput)
}

func (o SqlVirtualMachineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SqlVirtualMachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlVirtualMachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SqlVirtualMachineOutput) VirtualMachineResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.VirtualMachineResourceId }).(pulumi.StringPtrOutput)
}

func (o SqlVirtualMachineOutput) WsfcDomainCredentials() WsfcDomainCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) WsfcDomainCredentialsResponsePtrOutput { return v.WsfcDomainCredentials }).(WsfcDomainCredentialsResponsePtrOutput)
}

func (o SqlVirtualMachineOutput) WsfcStaticIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachine) pulumi.StringPtrOutput { return v.WsfcStaticIp }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlVirtualMachineOutput{})
}

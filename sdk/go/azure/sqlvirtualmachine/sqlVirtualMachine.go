


package sqlvirtualmachine

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlVirtualMachine struct {
	pulumi.CustomResourceState

	AutoBackupSettings                     AutoBackupSettingsResponsePtrOutput                     `pulumi:"autoBackupSettings"`
	AutoPatchingSettings                   AutoPatchingSettingsResponsePtrOutput                   `pulumi:"autoPatchingSettings"`
	Identity                               ResourceIdentityResponsePtrOutput                       `pulumi:"identity"`
	KeyVaultCredentialSettings             KeyVaultCredentialSettingsResponsePtrOutput             `pulumi:"keyVaultCredentialSettings"`
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
	Tags                                   pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type                                   pulumi.StringOutput                                     `pulumi:"type"`
	VirtualMachineResourceId               pulumi.StringPtrOutput                                  `pulumi:"virtualMachineResourceId"`
	WsfcDomainCredentials                  WsfcDomainCredentialsResponsePtrOutput                  `pulumi:"wsfcDomainCredentials"`
}


func NewSqlVirtualMachine(ctx *pulumi.Context,
	name string, args *SqlVirtualMachineArgs, opts ...pulumi.ResourceOption) (*SqlVirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sqlvirtualmachine:SqlVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20170301preview:SqlVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:sqlvirtualmachine/v20170301preview:SqlVirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlVirtualMachine
	err := ctx.RegisterResource("azure-native:sqlvirtualmachine:SqlVirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlVirtualMachineState, opts ...pulumi.ResourceOption) (*SqlVirtualMachine, error) {
	var resource SqlVirtualMachine
	err := ctx.ReadResource("azure-native:sqlvirtualmachine:SqlVirtualMachine", name, id, state, &resource, opts...)
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
	AutoBackupSettings                     *AutoBackupSettings                     `pulumi:"autoBackupSettings"`
	AutoPatchingSettings                   *AutoPatchingSettings                   `pulumi:"autoPatchingSettings"`
	Identity                               *ResourceIdentity                       `pulumi:"identity"`
	KeyVaultCredentialSettings             *KeyVaultCredentialSettings             `pulumi:"keyVaultCredentialSettings"`
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
}


type SqlVirtualMachineArgs struct {
	AutoBackupSettings                     AutoBackupSettingsPtrInput
	AutoPatchingSettings                   AutoPatchingSettingsPtrInput
	Identity                               ResourceIdentityPtrInput
	KeyVaultCredentialSettings             KeyVaultCredentialSettingsPtrInput
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
	return reflect.TypeOf((*SqlVirtualMachine)(nil))
}

func (i *SqlVirtualMachine) ToSqlVirtualMachineOutput() SqlVirtualMachineOutput {
	return i.ToSqlVirtualMachineOutputWithContext(context.Background())
}

func (i *SqlVirtualMachine) ToSqlVirtualMachineOutputWithContext(ctx context.Context) SqlVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlVirtualMachineOutput)
}

type SqlVirtualMachineOutput struct{ *pulumi.OutputState }

func (SqlVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlVirtualMachine)(nil))
}

func (o SqlVirtualMachineOutput) ToSqlVirtualMachineOutput() SqlVirtualMachineOutput {
	return o
}

func (o SqlVirtualMachineOutput) ToSqlVirtualMachineOutputWithContext(ctx context.Context) SqlVirtualMachineOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SqlVirtualMachineOutput{})
}

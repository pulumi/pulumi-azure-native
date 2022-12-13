


package resources

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzurePowerShellScript struct {
	pulumi.CustomResourceState

	Arguments              pulumi.StringPtrOutput                       `pulumi:"arguments"`
	AzPowerShellVersion    pulumi.StringOutput                          `pulumi:"azPowerShellVersion"`
	CleanupPreference      pulumi.StringPtrOutput                       `pulumi:"cleanupPreference"`
	ContainerSettings      ContainerConfigurationResponsePtrOutput      `pulumi:"containerSettings"`
	EnvironmentVariables   EnvironmentVariableResponseArrayOutput       `pulumi:"environmentVariables"`
	ForceUpdateTag         pulumi.StringPtrOutput                       `pulumi:"forceUpdateTag"`
	Identity               ManagedServiceIdentityResponsePtrOutput      `pulumi:"identity"`
	Kind                   pulumi.StringOutput                          `pulumi:"kind"`
	Location               pulumi.StringOutput                          `pulumi:"location"`
	Name                   pulumi.StringOutput                          `pulumi:"name"`
	Outputs                pulumi.MapOutput                             `pulumi:"outputs"`
	PrimaryScriptUri       pulumi.StringPtrOutput                       `pulumi:"primaryScriptUri"`
	ProvisioningState      pulumi.StringOutput                          `pulumi:"provisioningState"`
	RetentionInterval      pulumi.StringOutput                          `pulumi:"retentionInterval"`
	ScriptContent          pulumi.StringPtrOutput                       `pulumi:"scriptContent"`
	Status                 ScriptStatusResponseOutput                   `pulumi:"status"`
	StorageAccountSettings StorageAccountConfigurationResponsePtrOutput `pulumi:"storageAccountSettings"`
	SupportingScriptUris   pulumi.StringArrayOutput                     `pulumi:"supportingScriptUris"`
	SystemData             SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                   pulumi.StringMapOutput                       `pulumi:"tags"`
	Timeout                pulumi.StringPtrOutput                       `pulumi:"timeout"`
	Type                   pulumi.StringOutput                          `pulumi:"type"`
}


func NewAzurePowerShellScript(ctx *pulumi.Context,
	name string, args *AzurePowerShellScriptArgs, opts ...pulumi.ResourceOption) (*AzurePowerShellScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzPowerShellVersion == nil {
		return nil, errors.New("invalid value for required argument 'AzPowerShellVersion'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RetentionInterval == nil {
		return nil, errors.New("invalid value for required argument 'RetentionInterval'")
	}
	if isZero(args.CleanupPreference) {
		args.CleanupPreference = pulumi.StringPtr("Always")
	}
	args.Kind = pulumi.String("AzurePowerShell")
	if isZero(args.Timeout) {
		args.Timeout = pulumi.StringPtr("P1D")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources/v20191001preview:AzurePowerShellScript"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:AzurePowerShellScript"),
		},
	})
	opts = append(opts, aliases)
	var resource AzurePowerShellScript
	err := ctx.RegisterResource("azure-native:resources:AzurePowerShellScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAzurePowerShellScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzurePowerShellScriptState, opts ...pulumi.ResourceOption) (*AzurePowerShellScript, error) {
	var resource AzurePowerShellScript
	err := ctx.ReadResource("azure-native:resources:AzurePowerShellScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type azurePowerShellScriptState struct {
}

type AzurePowerShellScriptState struct {
}

func (AzurePowerShellScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*azurePowerShellScriptState)(nil)).Elem()
}

type azurePowerShellScriptArgs struct {
	Arguments              *string                      `pulumi:"arguments"`
	AzPowerShellVersion    string                       `pulumi:"azPowerShellVersion"`
	CleanupPreference      *string                      `pulumi:"cleanupPreference"`
	ContainerSettings      *ContainerConfiguration      `pulumi:"containerSettings"`
	EnvironmentVariables   []EnvironmentVariable        `pulumi:"environmentVariables"`
	ForceUpdateTag         *string                      `pulumi:"forceUpdateTag"`
	Identity               *ManagedServiceIdentity      `pulumi:"identity"`
	Kind                   string                       `pulumi:"kind"`
	Location               *string                      `pulumi:"location"`
	PrimaryScriptUri       *string                      `pulumi:"primaryScriptUri"`
	ResourceGroupName      string                       `pulumi:"resourceGroupName"`
	RetentionInterval      string                       `pulumi:"retentionInterval"`
	ScriptContent          *string                      `pulumi:"scriptContent"`
	ScriptName             *string                      `pulumi:"scriptName"`
	StorageAccountSettings *StorageAccountConfiguration `pulumi:"storageAccountSettings"`
	SupportingScriptUris   []string                     `pulumi:"supportingScriptUris"`
	Tags                   map[string]string            `pulumi:"tags"`
	Timeout                *string                      `pulumi:"timeout"`
}


type AzurePowerShellScriptArgs struct {
	Arguments              pulumi.StringPtrInput
	AzPowerShellVersion    pulumi.StringInput
	CleanupPreference      pulumi.StringPtrInput
	ContainerSettings      ContainerConfigurationPtrInput
	EnvironmentVariables   EnvironmentVariableArrayInput
	ForceUpdateTag         pulumi.StringPtrInput
	Identity               ManagedServiceIdentityPtrInput
	Kind                   pulumi.StringInput
	Location               pulumi.StringPtrInput
	PrimaryScriptUri       pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	RetentionInterval      pulumi.StringInput
	ScriptContent          pulumi.StringPtrInput
	ScriptName             pulumi.StringPtrInput
	StorageAccountSettings StorageAccountConfigurationPtrInput
	SupportingScriptUris   pulumi.StringArrayInput
	Tags                   pulumi.StringMapInput
	Timeout                pulumi.StringPtrInput
}

func (AzurePowerShellScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azurePowerShellScriptArgs)(nil)).Elem()
}

type AzurePowerShellScriptInput interface {
	pulumi.Input

	ToAzurePowerShellScriptOutput() AzurePowerShellScriptOutput
	ToAzurePowerShellScriptOutputWithContext(ctx context.Context) AzurePowerShellScriptOutput
}

func (*AzurePowerShellScript) ElementType() reflect.Type {
	return reflect.TypeOf((**AzurePowerShellScript)(nil)).Elem()
}

func (i *AzurePowerShellScript) ToAzurePowerShellScriptOutput() AzurePowerShellScriptOutput {
	return i.ToAzurePowerShellScriptOutputWithContext(context.Background())
}

func (i *AzurePowerShellScript) ToAzurePowerShellScriptOutputWithContext(ctx context.Context) AzurePowerShellScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzurePowerShellScriptOutput)
}

type AzurePowerShellScriptOutput struct{ *pulumi.OutputState }

func (AzurePowerShellScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzurePowerShellScript)(nil)).Elem()
}

func (o AzurePowerShellScriptOutput) ToAzurePowerShellScriptOutput() AzurePowerShellScriptOutput {
	return o
}

func (o AzurePowerShellScriptOutput) ToAzurePowerShellScriptOutputWithContext(ctx context.Context) AzurePowerShellScriptOutput {
	return o
}

func (o AzurePowerShellScriptOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringPtrOutput { return v.Arguments }).(pulumi.StringPtrOutput)
}

func (o AzurePowerShellScriptOutput) AzPowerShellVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.AzPowerShellVersion }).(pulumi.StringOutput)
}

func (o AzurePowerShellScriptOutput) CleanupPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringPtrOutput { return v.CleanupPreference }).(pulumi.StringPtrOutput)
}

func (o AzurePowerShellScriptOutput) ContainerSettings() ContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) ContainerConfigurationResponsePtrOutput { return v.ContainerSettings }).(ContainerConfigurationResponsePtrOutput)
}

func (o AzurePowerShellScriptOutput) EnvironmentVariables() EnvironmentVariableResponseArrayOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) EnvironmentVariableResponseArrayOutput { return v.EnvironmentVariables }).(EnvironmentVariableResponseArrayOutput)
}

func (o AzurePowerShellScriptOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o AzurePowerShellScriptOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o AzurePowerShellScriptOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o AzurePowerShellScriptOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AzurePowerShellScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AzurePowerShellScriptOutput) Outputs() pulumi.MapOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.MapOutput { return v.Outputs }).(pulumi.MapOutput)
}

func (o AzurePowerShellScriptOutput) PrimaryScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringPtrOutput { return v.PrimaryScriptUri }).(pulumi.StringPtrOutput)
}

func (o AzurePowerShellScriptOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AzurePowerShellScriptOutput) RetentionInterval() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.RetentionInterval }).(pulumi.StringOutput)
}

func (o AzurePowerShellScriptOutput) ScriptContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringPtrOutput { return v.ScriptContent }).(pulumi.StringPtrOutput)
}

func (o AzurePowerShellScriptOutput) Status() ScriptStatusResponseOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) ScriptStatusResponseOutput { return v.Status }).(ScriptStatusResponseOutput)
}

func (o AzurePowerShellScriptOutput) StorageAccountSettings() StorageAccountConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) StorageAccountConfigurationResponsePtrOutput {
		return v.StorageAccountSettings
	}).(StorageAccountConfigurationResponsePtrOutput)
}

func (o AzurePowerShellScriptOutput) SupportingScriptUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringArrayOutput { return v.SupportingScriptUris }).(pulumi.StringArrayOutput)
}

func (o AzurePowerShellScriptOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AzurePowerShellScriptOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AzurePowerShellScriptOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringPtrOutput { return v.Timeout }).(pulumi.StringPtrOutput)
}

func (o AzurePowerShellScriptOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AzurePowerShellScriptOutput{})
}

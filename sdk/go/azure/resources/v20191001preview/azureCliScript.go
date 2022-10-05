


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureCliScript struct {
	pulumi.CustomResourceState

	Arguments              pulumi.StringPtrOutput                       `pulumi:"arguments"`
	AzCliVersion           pulumi.StringOutput                          `pulumi:"azCliVersion"`
	CleanupPreference      pulumi.StringPtrOutput                       `pulumi:"cleanupPreference"`
	ContainerSettings      ContainerConfigurationResponsePtrOutput      `pulumi:"containerSettings"`
	EnvironmentVariables   EnvironmentVariableResponseArrayOutput       `pulumi:"environmentVariables"`
	ForceUpdateTag         pulumi.StringPtrOutput                       `pulumi:"forceUpdateTag"`
	Identity               ManagedServiceIdentityResponseOutput         `pulumi:"identity"`
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


func NewAzureCliScript(ctx *pulumi.Context,
	name string, args *AzureCliScriptArgs, opts ...pulumi.ResourceOption) (*AzureCliScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzCliVersion == nil {
		return nil, errors.New("invalid value for required argument 'AzCliVersion'")
	}
	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
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
	args.Kind = pulumi.String("AzureCLI")
	if isZero(args.Timeout) {
		args.Timeout = pulumi.StringPtr("P1D")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:AzureCliScript"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:AzureCliScript"),
		},
	})
	opts = append(opts, aliases)
	var resource AzureCliScript
	err := ctx.RegisterResource("azure-native:resources/v20191001preview:AzureCliScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAzureCliScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureCliScriptState, opts ...pulumi.ResourceOption) (*AzureCliScript, error) {
	var resource AzureCliScript
	err := ctx.ReadResource("azure-native:resources/v20191001preview:AzureCliScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type azureCliScriptState struct {
}

type AzureCliScriptState struct {
}

func (AzureCliScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureCliScriptState)(nil)).Elem()
}

type azureCliScriptArgs struct {
	Arguments              *string                      `pulumi:"arguments"`
	AzCliVersion           string                       `pulumi:"azCliVersion"`
	CleanupPreference      *string                      `pulumi:"cleanupPreference"`
	ContainerSettings      *ContainerConfiguration      `pulumi:"containerSettings"`
	EnvironmentVariables   []EnvironmentVariable        `pulumi:"environmentVariables"`
	ForceUpdateTag         *string                      `pulumi:"forceUpdateTag"`
	Identity               ManagedServiceIdentity       `pulumi:"identity"`
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


type AzureCliScriptArgs struct {
	Arguments              pulumi.StringPtrInput
	AzCliVersion           pulumi.StringInput
	CleanupPreference      pulumi.StringPtrInput
	ContainerSettings      ContainerConfigurationPtrInput
	EnvironmentVariables   EnvironmentVariableArrayInput
	ForceUpdateTag         pulumi.StringPtrInput
	Identity               ManagedServiceIdentityInput
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

func (AzureCliScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureCliScriptArgs)(nil)).Elem()
}

type AzureCliScriptInput interface {
	pulumi.Input

	ToAzureCliScriptOutput() AzureCliScriptOutput
	ToAzureCliScriptOutputWithContext(ctx context.Context) AzureCliScriptOutput
}

func (*AzureCliScript) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCliScript)(nil)).Elem()
}

func (i *AzureCliScript) ToAzureCliScriptOutput() AzureCliScriptOutput {
	return i.ToAzureCliScriptOutputWithContext(context.Background())
}

func (i *AzureCliScript) ToAzureCliScriptOutputWithContext(ctx context.Context) AzureCliScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureCliScriptOutput)
}

type AzureCliScriptOutput struct{ *pulumi.OutputState }

func (AzureCliScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCliScript)(nil)).Elem()
}

func (o AzureCliScriptOutput) ToAzureCliScriptOutput() AzureCliScriptOutput {
	return o
}

func (o AzureCliScriptOutput) ToAzureCliScriptOutputWithContext(ctx context.Context) AzureCliScriptOutput {
	return o
}

func (o AzureCliScriptOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringPtrOutput { return v.Arguments }).(pulumi.StringPtrOutput)
}

func (o AzureCliScriptOutput) AzCliVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.AzCliVersion }).(pulumi.StringOutput)
}

func (o AzureCliScriptOutput) CleanupPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringPtrOutput { return v.CleanupPreference }).(pulumi.StringPtrOutput)
}

func (o AzureCliScriptOutput) ContainerSettings() ContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AzureCliScript) ContainerConfigurationResponsePtrOutput { return v.ContainerSettings }).(ContainerConfigurationResponsePtrOutput)
}

func (o AzureCliScriptOutput) EnvironmentVariables() EnvironmentVariableResponseArrayOutput {
	return o.ApplyT(func(v *AzureCliScript) EnvironmentVariableResponseArrayOutput { return v.EnvironmentVariables }).(EnvironmentVariableResponseArrayOutput)
}

func (o AzureCliScriptOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o AzureCliScriptOutput) Identity() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *AzureCliScript) ManagedServiceIdentityResponseOutput { return v.Identity }).(ManagedServiceIdentityResponseOutput)
}

func (o AzureCliScriptOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o AzureCliScriptOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AzureCliScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AzureCliScriptOutput) Outputs() pulumi.MapOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.MapOutput { return v.Outputs }).(pulumi.MapOutput)
}

func (o AzureCliScriptOutput) PrimaryScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringPtrOutput { return v.PrimaryScriptUri }).(pulumi.StringPtrOutput)
}

func (o AzureCliScriptOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AzureCliScriptOutput) RetentionInterval() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.RetentionInterval }).(pulumi.StringOutput)
}

func (o AzureCliScriptOutput) ScriptContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringPtrOutput { return v.ScriptContent }).(pulumi.StringPtrOutput)
}

func (o AzureCliScriptOutput) Status() ScriptStatusResponseOutput {
	return o.ApplyT(func(v *AzureCliScript) ScriptStatusResponseOutput { return v.Status }).(ScriptStatusResponseOutput)
}

func (o AzureCliScriptOutput) StorageAccountSettings() StorageAccountConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AzureCliScript) StorageAccountConfigurationResponsePtrOutput { return v.StorageAccountSettings }).(StorageAccountConfigurationResponsePtrOutput)
}

func (o AzureCliScriptOutput) SupportingScriptUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringArrayOutput { return v.SupportingScriptUris }).(pulumi.StringArrayOutput)
}

func (o AzureCliScriptOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AzureCliScript) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AzureCliScriptOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AzureCliScriptOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringPtrOutput { return v.Timeout }).(pulumi.StringPtrOutput)
}

func (o AzureCliScriptOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCliScript) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureCliScriptOutput{})
}




package v20201001

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
			Type: pulumi.String("azure-native:resources:AzurePowerShellScript"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001preview:AzurePowerShellScript"),
		},
	})
	opts = append(opts, aliases)
	var resource AzurePowerShellScript
	err := ctx.RegisterResource("azure-native:resources/v20201001:AzurePowerShellScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAzurePowerShellScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzurePowerShellScriptState, opts ...pulumi.ResourceOption) (*AzurePowerShellScript, error) {
	var resource AzurePowerShellScript
	err := ctx.ReadResource("azure-native:resources/v20201001:AzurePowerShellScript", name, id, state, &resource, opts...)
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

func init() {
	pulumi.RegisterOutputType(AzurePowerShellScriptOutput{})
}

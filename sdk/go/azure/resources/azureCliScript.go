


package resources

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


func NewAzureCliScript(ctx *pulumi.Context,
	name string, args *AzureCliScriptArgs, opts ...pulumi.ResourceOption) (*AzureCliScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzCliVersion == nil {
		return nil, errors.New("invalid value for required argument 'AzCliVersion'")
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
	if args.CleanupPreference == nil {
		args.CleanupPreference = pulumi.StringPtr("Always")
	}
	args.Kind = pulumi.String("AzureCLI")
	if args.Timeout == nil {
		args.Timeout = pulumi.StringPtr("P1D")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:resources:AzureCliScript"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001preview:AzureCliScript"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20191001preview:AzureCliScript"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:AzureCliScript"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20201001:AzureCliScript"),
		},
	})
	opts = append(opts, aliases)
	var resource AzureCliScript
	err := ctx.RegisterResource("azure-native:resources:AzureCliScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAzureCliScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureCliScriptState, opts ...pulumi.ResourceOption) (*AzureCliScript, error) {
	var resource AzureCliScript
	err := ctx.ReadResource("azure-native:resources:AzureCliScript", name, id, state, &resource, opts...)
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


type AzureCliScriptArgs struct {
	Arguments              pulumi.StringPtrInput
	AzCliVersion           pulumi.StringInput
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

func (AzureCliScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureCliScriptArgs)(nil)).Elem()
}

type AzureCliScriptInput interface {
	pulumi.Input

	ToAzureCliScriptOutput() AzureCliScriptOutput
	ToAzureCliScriptOutputWithContext(ctx context.Context) AzureCliScriptOutput
}

func (*AzureCliScript) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureCliScript)(nil))
}

func (i *AzureCliScript) ToAzureCliScriptOutput() AzureCliScriptOutput {
	return i.ToAzureCliScriptOutputWithContext(context.Background())
}

func (i *AzureCliScript) ToAzureCliScriptOutputWithContext(ctx context.Context) AzureCliScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureCliScriptOutput)
}

type AzureCliScriptOutput struct{ *pulumi.OutputState }

func (AzureCliScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureCliScript)(nil))
}

func (o AzureCliScriptOutput) ToAzureCliScriptOutput() AzureCliScriptOutput {
	return o
}

func (o AzureCliScriptOutput) ToAzureCliScriptOutputWithContext(ctx context.Context) AzureCliScriptOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AzureCliScriptOutput{})
}

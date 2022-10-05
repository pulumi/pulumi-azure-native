


package app

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedEnvironment struct {
	pulumi.CustomResourceState

	AppLogsConfiguration     AppLogsConfigurationResponsePtrOutput `pulumi:"appLogsConfiguration"`
	DaprAIConnectionString   pulumi.StringPtrOutput                `pulumi:"daprAIConnectionString"`
	DaprAIInstrumentationKey pulumi.StringPtrOutput                `pulumi:"daprAIInstrumentationKey"`
	DefaultDomain            pulumi.StringOutput                   `pulumi:"defaultDomain"`
	DeploymentErrors         pulumi.StringOutput                   `pulumi:"deploymentErrors"`
	Location                 pulumi.StringOutput                   `pulumi:"location"`
	Name                     pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState        pulumi.StringOutput                   `pulumi:"provisioningState"`
	StaticIp                 pulumi.StringOutput                   `pulumi:"staticIp"`
	SystemData               SystemDataResponseOutput              `pulumi:"systemData"`
	Tags                     pulumi.StringMapOutput                `pulumi:"tags"`
	Type                     pulumi.StringOutput                   `pulumi:"type"`
	VnetConfiguration        VnetConfigurationResponsePtrOutput    `pulumi:"vnetConfiguration"`
	ZoneRedundant            pulumi.BoolPtrOutput                  `pulumi:"zoneRedundant"`
}


func NewManagedEnvironment(ctx *pulumi.Context,
	name string, args *ManagedEnvironmentArgs, opts ...pulumi.ResourceOption) (*ManagedEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app/v20220101preview:ManagedEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220301:ManagedEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220601preview:ManagedEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedEnvironment
	err := ctx.RegisterResource("azure-native:app:ManagedEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedEnvironmentState, opts ...pulumi.ResourceOption) (*ManagedEnvironment, error) {
	var resource ManagedEnvironment
	err := ctx.ReadResource("azure-native:app:ManagedEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedEnvironmentState struct {
}

type ManagedEnvironmentState struct {
}

func (ManagedEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedEnvironmentState)(nil)).Elem()
}

type managedEnvironmentArgs struct {
	AppLogsConfiguration     *AppLogsConfiguration `pulumi:"appLogsConfiguration"`
	DaprAIConnectionString   *string               `pulumi:"daprAIConnectionString"`
	DaprAIInstrumentationKey *string               `pulumi:"daprAIInstrumentationKey"`
	EnvironmentName          *string               `pulumi:"environmentName"`
	Location                 *string               `pulumi:"location"`
	ResourceGroupName        string                `pulumi:"resourceGroupName"`
	Tags                     map[string]string     `pulumi:"tags"`
	VnetConfiguration        *VnetConfiguration    `pulumi:"vnetConfiguration"`
	ZoneRedundant            *bool                 `pulumi:"zoneRedundant"`
}


type ManagedEnvironmentArgs struct {
	AppLogsConfiguration     AppLogsConfigurationPtrInput
	DaprAIConnectionString   pulumi.StringPtrInput
	DaprAIInstrumentationKey pulumi.StringPtrInput
	EnvironmentName          pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	Tags                     pulumi.StringMapInput
	VnetConfiguration        VnetConfigurationPtrInput
	ZoneRedundant            pulumi.BoolPtrInput
}

func (ManagedEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedEnvironmentArgs)(nil)).Elem()
}

type ManagedEnvironmentInput interface {
	pulumi.Input

	ToManagedEnvironmentOutput() ManagedEnvironmentOutput
	ToManagedEnvironmentOutputWithContext(ctx context.Context) ManagedEnvironmentOutput
}

func (*ManagedEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedEnvironment)(nil)).Elem()
}

func (i *ManagedEnvironment) ToManagedEnvironmentOutput() ManagedEnvironmentOutput {
	return i.ToManagedEnvironmentOutputWithContext(context.Background())
}

func (i *ManagedEnvironment) ToManagedEnvironmentOutputWithContext(ctx context.Context) ManagedEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedEnvironmentOutput)
}

type ManagedEnvironmentOutput struct{ *pulumi.OutputState }

func (ManagedEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedEnvironment)(nil)).Elem()
}

func (o ManagedEnvironmentOutput) ToManagedEnvironmentOutput() ManagedEnvironmentOutput {
	return o
}

func (o ManagedEnvironmentOutput) ToManagedEnvironmentOutputWithContext(ctx context.Context) ManagedEnvironmentOutput {
	return o
}

func (o ManagedEnvironmentOutput) AppLogsConfiguration() AppLogsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ManagedEnvironment) AppLogsConfigurationResponsePtrOutput { return v.AppLogsConfiguration }).(AppLogsConfigurationResponsePtrOutput)
}

func (o ManagedEnvironmentOutput) DaprAIConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringPtrOutput { return v.DaprAIConnectionString }).(pulumi.StringPtrOutput)
}

func (o ManagedEnvironmentOutput) DaprAIInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringPtrOutput { return v.DaprAIInstrumentationKey }).(pulumi.StringPtrOutput)
}

func (o ManagedEnvironmentOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.DefaultDomain }).(pulumi.StringOutput)
}

func (o ManagedEnvironmentOutput) DeploymentErrors() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.DeploymentErrors }).(pulumi.StringOutput)
}

func (o ManagedEnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ManagedEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedEnvironmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedEnvironmentOutput) StaticIp() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.StaticIp }).(pulumi.StringOutput)
}

func (o ManagedEnvironmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedEnvironment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ManagedEnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedEnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedEnvironmentOutput) VnetConfiguration() VnetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ManagedEnvironment) VnetConfigurationResponsePtrOutput { return v.VnetConfiguration }).(VnetConfigurationResponsePtrOutput)
}

func (o ManagedEnvironmentOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedEnvironment) pulumi.BoolPtrOutput { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedEnvironmentOutput{})
}




package v20221001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectedEnvironment struct {
	pulumi.CustomResourceState

	CustomDomainConfiguration CustomDomainConfigurationResponsePtrOutput `pulumi:"customDomainConfiguration"`
	DaprAIConnectionString    pulumi.StringPtrOutput                     `pulumi:"daprAIConnectionString"`
	DefaultDomain             pulumi.StringOutput                        `pulumi:"defaultDomain"`
	DeploymentErrors          pulumi.StringOutput                        `pulumi:"deploymentErrors"`
	ExtendedLocation          ExtendedLocationResponsePtrOutput          `pulumi:"extendedLocation"`
	Location                  pulumi.StringOutput                        `pulumi:"location"`
	Name                      pulumi.StringOutput                        `pulumi:"name"`
	ProvisioningState         pulumi.StringOutput                        `pulumi:"provisioningState"`
	StaticIp                  pulumi.StringPtrOutput                     `pulumi:"staticIp"`
	SystemData                SystemDataResponseOutput                   `pulumi:"systemData"`
	Tags                      pulumi.StringMapOutput                     `pulumi:"tags"`
	Type                      pulumi.StringOutput                        `pulumi:"type"`
}


func NewConnectedEnvironment(ctx *pulumi.Context,
	name string, args *ConnectedEnvironmentArgs, opts ...pulumi.ResourceOption) (*ConnectedEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app/v20220601preview:ConnectedEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectedEnvironment
	err := ctx.RegisterResource("azure-native:app/v20221001:ConnectedEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectedEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedEnvironmentState, opts ...pulumi.ResourceOption) (*ConnectedEnvironment, error) {
	var resource ConnectedEnvironment
	err := ctx.ReadResource("azure-native:app/v20221001:ConnectedEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectedEnvironmentState struct {
}

type ConnectedEnvironmentState struct {
}

func (ConnectedEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentState)(nil)).Elem()
}

type connectedEnvironmentArgs struct {
	ConnectedEnvironmentName  *string                    `pulumi:"connectedEnvironmentName"`
	CustomDomainConfiguration *CustomDomainConfiguration `pulumi:"customDomainConfiguration"`
	DaprAIConnectionString    *string                    `pulumi:"daprAIConnectionString"`
	ExtendedLocation          *ExtendedLocation          `pulumi:"extendedLocation"`
	Location                  *string                    `pulumi:"location"`
	ResourceGroupName         string                     `pulumi:"resourceGroupName"`
	StaticIp                  *string                    `pulumi:"staticIp"`
	Tags                      map[string]string          `pulumi:"tags"`
}


type ConnectedEnvironmentArgs struct {
	ConnectedEnvironmentName  pulumi.StringPtrInput
	CustomDomainConfiguration CustomDomainConfigurationPtrInput
	DaprAIConnectionString    pulumi.StringPtrInput
	ExtendedLocation          ExtendedLocationPtrInput
	Location                  pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	StaticIp                  pulumi.StringPtrInput
	Tags                      pulumi.StringMapInput
}

func (ConnectedEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentArgs)(nil)).Elem()
}

type ConnectedEnvironmentInput interface {
	pulumi.Input

	ToConnectedEnvironmentOutput() ConnectedEnvironmentOutput
	ToConnectedEnvironmentOutputWithContext(ctx context.Context) ConnectedEnvironmentOutput
}

func (*ConnectedEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironment)(nil)).Elem()
}

func (i *ConnectedEnvironment) ToConnectedEnvironmentOutput() ConnectedEnvironmentOutput {
	return i.ToConnectedEnvironmentOutputWithContext(context.Background())
}

func (i *ConnectedEnvironment) ToConnectedEnvironmentOutputWithContext(ctx context.Context) ConnectedEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedEnvironmentOutput)
}

type ConnectedEnvironmentOutput struct{ *pulumi.OutputState }

func (ConnectedEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironment)(nil)).Elem()
}

func (o ConnectedEnvironmentOutput) ToConnectedEnvironmentOutput() ConnectedEnvironmentOutput {
	return o
}

func (o ConnectedEnvironmentOutput) ToConnectedEnvironmentOutputWithContext(ctx context.Context) ConnectedEnvironmentOutput {
	return o
}

func (o ConnectedEnvironmentOutput) CustomDomainConfiguration() CustomDomainConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) CustomDomainConfigurationResponsePtrOutput {
		return v.CustomDomainConfiguration
	}).(CustomDomainConfigurationResponsePtrOutput)
}

func (o ConnectedEnvironmentOutput) DaprAIConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringPtrOutput { return v.DaprAIConnectionString }).(pulumi.StringPtrOutput)
}

func (o ConnectedEnvironmentOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringOutput { return v.DefaultDomain }).(pulumi.StringOutput)
}

func (o ConnectedEnvironmentOutput) DeploymentErrors() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringOutput { return v.DeploymentErrors }).(pulumi.StringOutput)
}

func (o ConnectedEnvironmentOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o ConnectedEnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ConnectedEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectedEnvironmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConnectedEnvironmentOutput) StaticIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringPtrOutput { return v.StaticIp }).(pulumi.StringPtrOutput)
}

func (o ConnectedEnvironmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConnectedEnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConnectedEnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedEnvironmentOutput{})
}

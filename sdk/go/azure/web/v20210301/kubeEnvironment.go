


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KubeEnvironment struct {
	pulumi.CustomResourceState

	AksResourceID               pulumi.StringPtrOutput                      `pulumi:"aksResourceID"`
	AppLogsConfiguration        AppLogsConfigurationResponsePtrOutput       `pulumi:"appLogsConfiguration"`
	ArcConfiguration            ArcConfigurationResponsePtrOutput           `pulumi:"arcConfiguration"`
	ContainerAppsConfiguration  ContainerAppsConfigurationResponsePtrOutput `pulumi:"containerAppsConfiguration"`
	DefaultDomain               pulumi.StringOutput                         `pulumi:"defaultDomain"`
	DeploymentErrors            pulumi.StringOutput                         `pulumi:"deploymentErrors"`
	EnvironmentType             pulumi.StringPtrOutput                      `pulumi:"environmentType"`
	ExtendedLocation            ExtendedLocationResponsePtrOutput           `pulumi:"extendedLocation"`
	InternalLoadBalancerEnabled pulumi.BoolPtrOutput                        `pulumi:"internalLoadBalancerEnabled"`
	Kind                        pulumi.StringPtrOutput                      `pulumi:"kind"`
	Location                    pulumi.StringOutput                         `pulumi:"location"`
	Name                        pulumi.StringOutput                         `pulumi:"name"`
	ProvisioningState           pulumi.StringOutput                         `pulumi:"provisioningState"`
	StaticIp                    pulumi.StringPtrOutput                      `pulumi:"staticIp"`
	Tags                        pulumi.StringMapOutput                      `pulumi:"tags"`
	Type                        pulumi.StringOutput                         `pulumi:"type"`
}


func NewKubeEnvironment(ctx *pulumi.Context,
	name string, args *KubeEnvironmentArgs, opts ...pulumi.ResourceOption) (*KubeEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:KubeEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource KubeEnvironment
	err := ctx.RegisterResource("azure-native:web/v20210301:KubeEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKubeEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubeEnvironmentState, opts ...pulumi.ResourceOption) (*KubeEnvironment, error) {
	var resource KubeEnvironment
	err := ctx.ReadResource("azure-native:web/v20210301:KubeEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kubeEnvironmentState struct {
}

type KubeEnvironmentState struct {
}

func (KubeEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeEnvironmentState)(nil)).Elem()
}

type kubeEnvironmentArgs struct {
	AksResourceID               *string                     `pulumi:"aksResourceID"`
	AppLogsConfiguration        *AppLogsConfiguration       `pulumi:"appLogsConfiguration"`
	ArcConfiguration            *ArcConfiguration           `pulumi:"arcConfiguration"`
	ContainerAppsConfiguration  *ContainerAppsConfiguration `pulumi:"containerAppsConfiguration"`
	EnvironmentType             *string                     `pulumi:"environmentType"`
	ExtendedLocation            *ExtendedLocation           `pulumi:"extendedLocation"`
	InternalLoadBalancerEnabled *bool                       `pulumi:"internalLoadBalancerEnabled"`
	Kind                        *string                     `pulumi:"kind"`
	Location                    *string                     `pulumi:"location"`
	Name                        *string                     `pulumi:"name"`
	ResourceGroupName           string                      `pulumi:"resourceGroupName"`
	StaticIp                    *string                     `pulumi:"staticIp"`
	Tags                        map[string]string           `pulumi:"tags"`
}


type KubeEnvironmentArgs struct {
	AksResourceID               pulumi.StringPtrInput
	AppLogsConfiguration        AppLogsConfigurationPtrInput
	ArcConfiguration            ArcConfigurationPtrInput
	ContainerAppsConfiguration  ContainerAppsConfigurationPtrInput
	EnvironmentType             pulumi.StringPtrInput
	ExtendedLocation            ExtendedLocationPtrInput
	InternalLoadBalancerEnabled pulumi.BoolPtrInput
	Kind                        pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	Name                        pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	StaticIp                    pulumi.StringPtrInput
	Tags                        pulumi.StringMapInput
}

func (KubeEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeEnvironmentArgs)(nil)).Elem()
}

type KubeEnvironmentInput interface {
	pulumi.Input

	ToKubeEnvironmentOutput() KubeEnvironmentOutput
	ToKubeEnvironmentOutputWithContext(ctx context.Context) KubeEnvironmentOutput
}

func (*KubeEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeEnvironment)(nil)).Elem()
}

func (i *KubeEnvironment) ToKubeEnvironmentOutput() KubeEnvironmentOutput {
	return i.ToKubeEnvironmentOutputWithContext(context.Background())
}

func (i *KubeEnvironment) ToKubeEnvironmentOutputWithContext(ctx context.Context) KubeEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeEnvironmentOutput)
}

type KubeEnvironmentOutput struct{ *pulumi.OutputState }

func (KubeEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeEnvironment)(nil)).Elem()
}

func (o KubeEnvironmentOutput) ToKubeEnvironmentOutput() KubeEnvironmentOutput {
	return o
}

func (o KubeEnvironmentOutput) ToKubeEnvironmentOutputWithContext(ctx context.Context) KubeEnvironmentOutput {
	return o
}

func (o KubeEnvironmentOutput) AksResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeEnvironment) pulumi.StringPtrOutput { return v.AksResourceID }).(pulumi.StringPtrOutput)
}

func (o KubeEnvironmentOutput) AppLogsConfiguration() AppLogsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *KubeEnvironment) AppLogsConfigurationResponsePtrOutput { return v.AppLogsConfiguration }).(AppLogsConfigurationResponsePtrOutput)
}

func (o KubeEnvironmentOutput) ArcConfiguration() ArcConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *KubeEnvironment) ArcConfigurationResponsePtrOutput { return v.ArcConfiguration }).(ArcConfigurationResponsePtrOutput)
}

func (o KubeEnvironmentOutput) ContainerAppsConfiguration() ContainerAppsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *KubeEnvironment) ContainerAppsConfigurationResponsePtrOutput {
		return v.ContainerAppsConfiguration
	}).(ContainerAppsConfigurationResponsePtrOutput)
}

func (o KubeEnvironmentOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeEnvironment) pulumi.StringOutput { return v.DefaultDomain }).(pulumi.StringOutput)
}

func (o KubeEnvironmentOutput) DeploymentErrors() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeEnvironment) pulumi.StringOutput { return v.DeploymentErrors }).(pulumi.StringOutput)
}

func (o KubeEnvironmentOutput) EnvironmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeEnvironment) pulumi.StringPtrOutput { return v.EnvironmentType }).(pulumi.StringPtrOutput)
}

func (o KubeEnvironmentOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *KubeEnvironment) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o KubeEnvironmentOutput) InternalLoadBalancerEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubeEnvironment) pulumi.BoolPtrOutput { return v.InternalLoadBalancerEnabled }).(pulumi.BoolPtrOutput)
}

func (o KubeEnvironmentOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeEnvironment) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o KubeEnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeEnvironment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o KubeEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KubeEnvironmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeEnvironment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o KubeEnvironmentOutput) StaticIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeEnvironment) pulumi.StringPtrOutput { return v.StaticIp }).(pulumi.StringPtrOutput)
}

func (o KubeEnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *KubeEnvironment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o KubeEnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeEnvironment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KubeEnvironmentOutput{})
}

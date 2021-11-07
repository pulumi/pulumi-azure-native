


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KubeEnvironment struct {
	pulumi.CustomResourceState

	AksResourceID               pulumi.StringPtrOutput                `pulumi:"aksResourceID"`
	AppLogsConfiguration        AppLogsConfigurationResponsePtrOutput `pulumi:"appLogsConfiguration"`
	ArcConfiguration            ArcConfigurationResponsePtrOutput     `pulumi:"arcConfiguration"`
	DefaultDomain               pulumi.StringOutput                   `pulumi:"defaultDomain"`
	DeploymentErrors            pulumi.StringOutput                   `pulumi:"deploymentErrors"`
	ExtendedLocation            ExtendedLocationResponsePtrOutput     `pulumi:"extendedLocation"`
	InternalLoadBalancerEnabled pulumi.BoolPtrOutput                  `pulumi:"internalLoadBalancerEnabled"`
	Kind                        pulumi.StringPtrOutput                `pulumi:"kind"`
	Location                    pulumi.StringOutput                   `pulumi:"location"`
	Name                        pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState           pulumi.StringOutput                   `pulumi:"provisioningState"`
	StaticIp                    pulumi.StringPtrOutput                `pulumi:"staticIp"`
	Tags                        pulumi.StringMapOutput                `pulumi:"tags"`
	Type                        pulumi.StringOutput                   `pulumi:"type"`
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
			Type: pulumi.String("azure-nextgen:web/v20210201:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:KubeEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210301:KubeEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource KubeEnvironment
	err := ctx.RegisterResource("azure-native:web/v20210201:KubeEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKubeEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubeEnvironmentState, opts ...pulumi.ResourceOption) (*KubeEnvironment, error) {
	var resource KubeEnvironment
	err := ctx.ReadResource("azure-native:web/v20210201:KubeEnvironment", name, id, state, &resource, opts...)
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
	AksResourceID               *string               `pulumi:"aksResourceID"`
	AppLogsConfiguration        *AppLogsConfiguration `pulumi:"appLogsConfiguration"`
	ArcConfiguration            *ArcConfiguration     `pulumi:"arcConfiguration"`
	ExtendedLocation            *ExtendedLocation     `pulumi:"extendedLocation"`
	InternalLoadBalancerEnabled *bool                 `pulumi:"internalLoadBalancerEnabled"`
	Kind                        *string               `pulumi:"kind"`
	Location                    *string               `pulumi:"location"`
	Name                        *string               `pulumi:"name"`
	ResourceGroupName           string                `pulumi:"resourceGroupName"`
	StaticIp                    *string               `pulumi:"staticIp"`
	Tags                        map[string]string     `pulumi:"tags"`
}


type KubeEnvironmentArgs struct {
	AksResourceID               pulumi.StringPtrInput
	AppLogsConfiguration        AppLogsConfigurationPtrInput
	ArcConfiguration            ArcConfigurationPtrInput
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
	return reflect.TypeOf((*KubeEnvironment)(nil))
}

func (i *KubeEnvironment) ToKubeEnvironmentOutput() KubeEnvironmentOutput {
	return i.ToKubeEnvironmentOutputWithContext(context.Background())
}

func (i *KubeEnvironment) ToKubeEnvironmentOutputWithContext(ctx context.Context) KubeEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeEnvironmentOutput)
}

type KubeEnvironmentOutput struct{ *pulumi.OutputState }

func (KubeEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeEnvironment)(nil))
}

func (o KubeEnvironmentOutput) ToKubeEnvironmentOutput() KubeEnvironmentOutput {
	return o
}

func (o KubeEnvironmentOutput) ToKubeEnvironmentOutputWithContext(ctx context.Context) KubeEnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KubeEnvironmentOutput{})
}

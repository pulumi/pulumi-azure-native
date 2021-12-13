


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerApp struct {
	pulumi.CustomResourceState

	Configuration      ConfigurationResponsePtrOutput `pulumi:"configuration"`
	Kind               pulumi.StringPtrOutput         `pulumi:"kind"`
	KubeEnvironmentId  pulumi.StringPtrOutput         `pulumi:"kubeEnvironmentId"`
	LatestRevisionFqdn pulumi.StringOutput            `pulumi:"latestRevisionFqdn"`
	LatestRevisionName pulumi.StringOutput            `pulumi:"latestRevisionName"`
	Location           pulumi.StringOutput            `pulumi:"location"`
	Name               pulumi.StringOutput            `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput            `pulumi:"provisioningState"`
	Tags               pulumi.StringMapOutput         `pulumi:"tags"`
	Template           TemplateResponsePtrOutput      `pulumi:"template"`
	Type               pulumi.StringOutput            `pulumi:"type"`
}


func NewContainerApp(ctx *pulumi.Context,
	name string, args *ContainerAppArgs, opts ...pulumi.ResourceOption) (*ContainerApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Configuration != nil {
		args.Configuration = args.Configuration.ToConfigurationPtrOutput().ApplyT(func(v *Configuration) *Configuration { return v.Defaults() }).(ConfigurationPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:ContainerApp"),
		},
	})
	opts = append(opts, aliases)
	var resource ContainerApp
	err := ctx.RegisterResource("azure-native:web/v20210301:ContainerApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContainerApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerAppState, opts ...pulumi.ResourceOption) (*ContainerApp, error) {
	var resource ContainerApp
	err := ctx.ReadResource("azure-native:web/v20210301:ContainerApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type containerAppState struct {
}

type ContainerAppState struct {
}

func (ContainerAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppState)(nil)).Elem()
}

type containerAppArgs struct {
	Configuration     *Configuration    `pulumi:"configuration"`
	Kind              *string           `pulumi:"kind"`
	KubeEnvironmentId *string           `pulumi:"kubeEnvironmentId"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	Template          *Template         `pulumi:"template"`
}


type ContainerAppArgs struct {
	Configuration     ConfigurationPtrInput
	Kind              pulumi.StringPtrInput
	KubeEnvironmentId pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	Template          TemplatePtrInput
}

func (ContainerAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppArgs)(nil)).Elem()
}

type ContainerAppInput interface {
	pulumi.Input

	ToContainerAppOutput() ContainerAppOutput
	ToContainerAppOutputWithContext(ctx context.Context) ContainerAppOutput
}

func (*ContainerApp) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerApp)(nil)).Elem()
}

func (i *ContainerApp) ToContainerAppOutput() ContainerAppOutput {
	return i.ToContainerAppOutputWithContext(context.Background())
}

func (i *ContainerApp) ToContainerAppOutputWithContext(ctx context.Context) ContainerAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppOutput)
}

type ContainerAppOutput struct{ *pulumi.OutputState }

func (ContainerAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerApp)(nil)).Elem()
}

func (o ContainerAppOutput) ToContainerAppOutput() ContainerAppOutput {
	return o
}

func (o ContainerAppOutput) ToContainerAppOutputWithContext(ctx context.Context) ContainerAppOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ContainerAppOutput{})
}




package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BuildServiceAgentPool struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                           `pulumi:"name"`
	Properties BuildServiceAgentPoolPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                      `pulumi:"systemData"`
	Type       pulumi.StringOutput                           `pulumi:"type"`
}


func NewBuildServiceAgentPool(ctx *pulumi.Context,
	name string, args *BuildServiceAgentPoolArgs, opts ...pulumi.ResourceOption) (*BuildServiceAgentPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BuildServiceName == nil {
		return nil, errors.New("invalid value for required argument 'BuildServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:BuildServiceAgentPool"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:BuildServiceAgentPool"),
		},
	})
	opts = append(opts, aliases)
	var resource BuildServiceAgentPool
	err := ctx.RegisterResource("azure-native:appplatform/v20221101preview:BuildServiceAgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBuildServiceAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildServiceAgentPoolState, opts ...pulumi.ResourceOption) (*BuildServiceAgentPool, error) {
	var resource BuildServiceAgentPool
	err := ctx.ReadResource("azure-native:appplatform/v20221101preview:BuildServiceAgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type buildServiceAgentPoolState struct {
}

type BuildServiceAgentPoolState struct {
}

func (BuildServiceAgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildServiceAgentPoolState)(nil)).Elem()
}

type buildServiceAgentPoolArgs struct {
	AgentPoolName     *string                          `pulumi:"agentPoolName"`
	BuildServiceName  string                           `pulumi:"buildServiceName"`
	Properties        *BuildServiceAgentPoolProperties `pulumi:"properties"`
	ResourceGroupName string                           `pulumi:"resourceGroupName"`
	ServiceName       string                           `pulumi:"serviceName"`
}


type BuildServiceAgentPoolArgs struct {
	AgentPoolName     pulumi.StringPtrInput
	BuildServiceName  pulumi.StringInput
	Properties        BuildServiceAgentPoolPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (BuildServiceAgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildServiceAgentPoolArgs)(nil)).Elem()
}

type BuildServiceAgentPoolInput interface {
	pulumi.Input

	ToBuildServiceAgentPoolOutput() BuildServiceAgentPoolOutput
	ToBuildServiceAgentPoolOutputWithContext(ctx context.Context) BuildServiceAgentPoolOutput
}

func (*BuildServiceAgentPool) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildServiceAgentPool)(nil)).Elem()
}

func (i *BuildServiceAgentPool) ToBuildServiceAgentPoolOutput() BuildServiceAgentPoolOutput {
	return i.ToBuildServiceAgentPoolOutputWithContext(context.Background())
}

func (i *BuildServiceAgentPool) ToBuildServiceAgentPoolOutputWithContext(ctx context.Context) BuildServiceAgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildServiceAgentPoolOutput)
}

type BuildServiceAgentPoolOutput struct{ *pulumi.OutputState }

func (BuildServiceAgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildServiceAgentPool)(nil)).Elem()
}

func (o BuildServiceAgentPoolOutput) ToBuildServiceAgentPoolOutput() BuildServiceAgentPoolOutput {
	return o
}

func (o BuildServiceAgentPoolOutput) ToBuildServiceAgentPoolOutputWithContext(ctx context.Context) BuildServiceAgentPoolOutput {
	return o
}

func (o BuildServiceAgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildServiceAgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BuildServiceAgentPoolOutput) Properties() BuildServiceAgentPoolPropertiesResponseOutput {
	return o.ApplyT(func(v *BuildServiceAgentPool) BuildServiceAgentPoolPropertiesResponseOutput { return v.Properties }).(BuildServiceAgentPoolPropertiesResponseOutput)
}

func (o BuildServiceAgentPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BuildServiceAgentPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BuildServiceAgentPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildServiceAgentPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildServiceAgentPoolOutput{})
}

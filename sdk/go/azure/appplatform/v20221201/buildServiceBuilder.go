


package v20221201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BuildServiceBuilder struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput             `pulumi:"name"`
	Properties BuilderPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput        `pulumi:"systemData"`
	Type       pulumi.StringOutput             `pulumi:"type"`
}


func NewBuildServiceBuilder(ctx *pulumi.Context,
	name string, args *BuildServiceBuilderArgs, opts ...pulumi.ResourceOption) (*BuildServiceBuilder, error) {
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
			Type: pulumi.String("azure-native:appplatform:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:BuildServiceBuilder"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:BuildServiceBuilder"),
		},
	})
	opts = append(opts, aliases)
	var resource BuildServiceBuilder
	err := ctx.RegisterResource("azure-native:appplatform/v20221201:BuildServiceBuilder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBuildServiceBuilder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildServiceBuilderState, opts ...pulumi.ResourceOption) (*BuildServiceBuilder, error) {
	var resource BuildServiceBuilder
	err := ctx.ReadResource("azure-native:appplatform/v20221201:BuildServiceBuilder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type buildServiceBuilderState struct {
}

type BuildServiceBuilderState struct {
}

func (BuildServiceBuilderState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildServiceBuilderState)(nil)).Elem()
}

type buildServiceBuilderArgs struct {
	BuildServiceName  string             `pulumi:"buildServiceName"`
	BuilderName       *string            `pulumi:"builderName"`
	Properties        *BuilderProperties `pulumi:"properties"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	ServiceName       string             `pulumi:"serviceName"`
}


type BuildServiceBuilderArgs struct {
	BuildServiceName  pulumi.StringInput
	BuilderName       pulumi.StringPtrInput
	Properties        BuilderPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (BuildServiceBuilderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildServiceBuilderArgs)(nil)).Elem()
}

type BuildServiceBuilderInput interface {
	pulumi.Input

	ToBuildServiceBuilderOutput() BuildServiceBuilderOutput
	ToBuildServiceBuilderOutputWithContext(ctx context.Context) BuildServiceBuilderOutput
}

func (*BuildServiceBuilder) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildServiceBuilder)(nil)).Elem()
}

func (i *BuildServiceBuilder) ToBuildServiceBuilderOutput() BuildServiceBuilderOutput {
	return i.ToBuildServiceBuilderOutputWithContext(context.Background())
}

func (i *BuildServiceBuilder) ToBuildServiceBuilderOutputWithContext(ctx context.Context) BuildServiceBuilderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildServiceBuilderOutput)
}

type BuildServiceBuilderOutput struct{ *pulumi.OutputState }

func (BuildServiceBuilderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildServiceBuilder)(nil)).Elem()
}

func (o BuildServiceBuilderOutput) ToBuildServiceBuilderOutput() BuildServiceBuilderOutput {
	return o
}

func (o BuildServiceBuilderOutput) ToBuildServiceBuilderOutputWithContext(ctx context.Context) BuildServiceBuilderOutput {
	return o
}

func (o BuildServiceBuilderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildServiceBuilder) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BuildServiceBuilderOutput) Properties() BuilderPropertiesResponseOutput {
	return o.ApplyT(func(v *BuildServiceBuilder) BuilderPropertiesResponseOutput { return v.Properties }).(BuilderPropertiesResponseOutput)
}

func (o BuildServiceBuilderOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BuildServiceBuilder) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BuildServiceBuilderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildServiceBuilder) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildServiceBuilderOutput{})
}

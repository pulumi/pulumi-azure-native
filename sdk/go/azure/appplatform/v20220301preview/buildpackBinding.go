


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BuildpackBinding struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                      `pulumi:"name"`
	Properties BuildpackBindingPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                 `pulumi:"systemData"`
	Type       pulumi.StringOutput                      `pulumi:"type"`
}


func NewBuildpackBinding(ctx *pulumi.Context,
	name string, args *BuildpackBindingArgs, opts ...pulumi.ResourceOption) (*BuildpackBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BuildServiceName == nil {
		return nil, errors.New("invalid value for required argument 'BuildServiceName'")
	}
	if args.BuilderName == nil {
		return nil, errors.New("invalid value for required argument 'BuilderName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:BuildpackBinding"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:BuildpackBinding"),
		},
	})
	opts = append(opts, aliases)
	var resource BuildpackBinding
	err := ctx.RegisterResource("azure-native:appplatform/v20220301preview:BuildpackBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBuildpackBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildpackBindingState, opts ...pulumi.ResourceOption) (*BuildpackBinding, error) {
	var resource BuildpackBinding
	err := ctx.ReadResource("azure-native:appplatform/v20220301preview:BuildpackBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type buildpackBindingState struct {
}

type BuildpackBindingState struct {
}

func (BuildpackBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildpackBindingState)(nil)).Elem()
}

type buildpackBindingArgs struct {
	BuildServiceName     string                      `pulumi:"buildServiceName"`
	BuilderName          string                      `pulumi:"builderName"`
	BuildpackBindingName *string                     `pulumi:"buildpackBindingName"`
	Properties           *BuildpackBindingProperties `pulumi:"properties"`
	ResourceGroupName    string                      `pulumi:"resourceGroupName"`
	ServiceName          string                      `pulumi:"serviceName"`
}


type BuildpackBindingArgs struct {
	BuildServiceName     pulumi.StringInput
	BuilderName          pulumi.StringInput
	BuildpackBindingName pulumi.StringPtrInput
	Properties           BuildpackBindingPropertiesPtrInput
	ResourceGroupName    pulumi.StringInput
	ServiceName          pulumi.StringInput
}

func (BuildpackBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildpackBindingArgs)(nil)).Elem()
}

type BuildpackBindingInput interface {
	pulumi.Input

	ToBuildpackBindingOutput() BuildpackBindingOutput
	ToBuildpackBindingOutputWithContext(ctx context.Context) BuildpackBindingOutput
}

func (*BuildpackBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildpackBinding)(nil)).Elem()
}

func (i *BuildpackBinding) ToBuildpackBindingOutput() BuildpackBindingOutput {
	return i.ToBuildpackBindingOutputWithContext(context.Background())
}

func (i *BuildpackBinding) ToBuildpackBindingOutputWithContext(ctx context.Context) BuildpackBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildpackBindingOutput)
}

type BuildpackBindingOutput struct{ *pulumi.OutputState }

func (BuildpackBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildpackBinding)(nil)).Elem()
}

func (o BuildpackBindingOutput) ToBuildpackBindingOutput() BuildpackBindingOutput {
	return o
}

func (o BuildpackBindingOutput) ToBuildpackBindingOutputWithContext(ctx context.Context) BuildpackBindingOutput {
	return o
}

func (o BuildpackBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildpackBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BuildpackBindingOutput) Properties() BuildpackBindingPropertiesResponseOutput {
	return o.ApplyT(func(v *BuildpackBinding) BuildpackBindingPropertiesResponseOutput { return v.Properties }).(BuildpackBindingPropertiesResponseOutput)
}

func (o BuildpackBindingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BuildpackBinding) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BuildpackBindingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildpackBinding) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildpackBindingOutput{})
}

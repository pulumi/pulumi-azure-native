


package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GlobalParameter struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput                           `pulumi:"etag"`
	Name       pulumi.StringOutput                           `pulumi:"name"`
	Properties GlobalParameterSpecificationResponseMapOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                           `pulumi:"type"`
}


func NewGlobalParameter(ctx *pulumi.Context,
	name string, args *GlobalParameterArgs, opts ...pulumi.ResourceOption) (*GlobalParameter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FactoryName == nil {
		return nil, errors.New("invalid value for required argument 'FactoryName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datafactory/v20180601:GlobalParameter"),
		},
	})
	opts = append(opts, aliases)
	var resource GlobalParameter
	err := ctx.RegisterResource("azure-native:datafactory:GlobalParameter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGlobalParameter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalParameterState, opts ...pulumi.ResourceOption) (*GlobalParameter, error) {
	var resource GlobalParameter
	err := ctx.ReadResource("azure-native:datafactory:GlobalParameter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type globalParameterState struct {
}

type GlobalParameterState struct {
}

func (GlobalParameterState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalParameterState)(nil)).Elem()
}

type globalParameterArgs struct {
	FactoryName         string                                  `pulumi:"factoryName"`
	GlobalParameterName *string                                 `pulumi:"globalParameterName"`
	Properties          map[string]GlobalParameterSpecification `pulumi:"properties"`
	ResourceGroupName   string                                  `pulumi:"resourceGroupName"`
}


type GlobalParameterArgs struct {
	FactoryName         pulumi.StringInput
	GlobalParameterName pulumi.StringPtrInput
	Properties          GlobalParameterSpecificationMapInput
	ResourceGroupName   pulumi.StringInput
}

func (GlobalParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalParameterArgs)(nil)).Elem()
}

type GlobalParameterInput interface {
	pulumi.Input

	ToGlobalParameterOutput() GlobalParameterOutput
	ToGlobalParameterOutputWithContext(ctx context.Context) GlobalParameterOutput
}

func (*GlobalParameter) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalParameter)(nil)).Elem()
}

func (i *GlobalParameter) ToGlobalParameterOutput() GlobalParameterOutput {
	return i.ToGlobalParameterOutputWithContext(context.Background())
}

func (i *GlobalParameter) ToGlobalParameterOutputWithContext(ctx context.Context) GlobalParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalParameterOutput)
}

type GlobalParameterOutput struct{ *pulumi.OutputState }

func (GlobalParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalParameter)(nil)).Elem()
}

func (o GlobalParameterOutput) ToGlobalParameterOutput() GlobalParameterOutput {
	return o
}

func (o GlobalParameterOutput) ToGlobalParameterOutputWithContext(ctx context.Context) GlobalParameterOutput {
	return o
}

func (o GlobalParameterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalParameter) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o GlobalParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalParameter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GlobalParameterOutput) Properties() GlobalParameterSpecificationResponseMapOutput {
	return o.ApplyT(func(v *GlobalParameter) GlobalParameterSpecificationResponseMapOutput { return v.Properties }).(GlobalParameterSpecificationResponseMapOutput)
}

func (o GlobalParameterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalParameter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GlobalParameterOutput{})
}

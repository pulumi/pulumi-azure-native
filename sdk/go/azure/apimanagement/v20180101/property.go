


package v20180101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Property struct {
	pulumi.CustomResourceState

	DisplayName pulumi.StringOutput      `pulumi:"displayName"`
	Name        pulumi.StringOutput      `pulumi:"name"`
	Secret      pulumi.BoolPtrOutput     `pulumi:"secret"`
	Tags        pulumi.StringArrayOutput `pulumi:"tags"`
	Type        pulumi.StringOutput      `pulumi:"type"`
	Value       pulumi.StringOutput      `pulumi:"value"`
}


func NewProperty(ctx *pulumi.Context,
	name string, args *PropertyArgs, opts ...pulumi.ResourceOption) (*Property, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Property"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Property"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Property"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Property"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Property"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Property"),
		},
	})
	opts = append(opts, aliases)
	var resource Property
	err := ctx.RegisterResource("azure-native:apimanagement/v20180101:Property", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProperty(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PropertyState, opts ...pulumi.ResourceOption) (*Property, error) {
	var resource Property
	err := ctx.ReadResource("azure-native:apimanagement/v20180101:Property", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type propertyState struct {
}

type PropertyState struct {
}

func (PropertyState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyState)(nil)).Elem()
}

type propertyArgs struct {
	DisplayName       string   `pulumi:"displayName"`
	PropId            *string  `pulumi:"propId"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Secret            *bool    `pulumi:"secret"`
	ServiceName       string   `pulumi:"serviceName"`
	Tags              []string `pulumi:"tags"`
	Value             string   `pulumi:"value"`
}


type PropertyArgs struct {
	DisplayName       pulumi.StringInput
	PropId            pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Secret            pulumi.BoolPtrInput
	ServiceName       pulumi.StringInput
	Tags              pulumi.StringArrayInput
	Value             pulumi.StringInput
}

func (PropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyArgs)(nil)).Elem()
}

type PropertyInput interface {
	pulumi.Input

	ToPropertyOutput() PropertyOutput
	ToPropertyOutputWithContext(ctx context.Context) PropertyOutput
}

func (*Property) ElementType() reflect.Type {
	return reflect.TypeOf((**Property)(nil)).Elem()
}

func (i *Property) ToPropertyOutput() PropertyOutput {
	return i.ToPropertyOutputWithContext(context.Background())
}

func (i *Property) ToPropertyOutputWithContext(ctx context.Context) PropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyOutput)
}

type PropertyOutput struct{ *pulumi.OutputState }

func (PropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Property)(nil)).Elem()
}

func (o PropertyOutput) ToPropertyOutput() PropertyOutput {
	return o
}

func (o PropertyOutput) ToPropertyOutputWithContext(ctx context.Context) PropertyOutput {
	return o
}

func (o PropertyOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Property) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o PropertyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Property) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PropertyOutput) Secret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Property) pulumi.BoolPtrOutput { return v.Secret }).(pulumi.BoolPtrOutput)
}

func (o PropertyOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Property) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o PropertyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Property) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PropertyOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *Property) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PropertyOutput{})
}

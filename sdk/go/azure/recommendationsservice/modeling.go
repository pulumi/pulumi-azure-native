


package recommendationsservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Modeling struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                      `pulumi:"location"`
	Name       pulumi.StringOutput                      `pulumi:"name"`
	Properties ModelingResourceResponsePropertiesOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                 `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                   `pulumi:"tags"`
	Type       pulumi.StringOutput                      `pulumi:"type"`
}


func NewModeling(ctx *pulumi.Context,
	name string, args *ModelingArgs, opts ...pulumi.ResourceOption) (*Modeling, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recommendationsservice/v20220201:Modeling"),
		},
		{
			Type: pulumi.String("azure-native:recommendationsservice/v20220301preview:Modeling"),
		},
	})
	opts = append(opts, aliases)
	var resource Modeling
	err := ctx.RegisterResource("azure-native:recommendationsservice:Modeling", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetModeling(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelingState, opts ...pulumi.ResourceOption) (*Modeling, error) {
	var resource Modeling
	err := ctx.ReadResource("azure-native:recommendationsservice:Modeling", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type modelingState struct {
}

type ModelingState struct {
}

func (ModelingState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelingState)(nil)).Elem()
}

type modelingArgs struct {
	AccountName       string                      `pulumi:"accountName"`
	Location          *string                     `pulumi:"location"`
	ModelingName      *string                     `pulumi:"modelingName"`
	Properties        *ModelingResourceProperties `pulumi:"properties"`
	ResourceGroupName string                      `pulumi:"resourceGroupName"`
	Tags              map[string]string           `pulumi:"tags"`
}


type ModelingArgs struct {
	AccountName       pulumi.StringInput
	Location          pulumi.StringPtrInput
	ModelingName      pulumi.StringPtrInput
	Properties        ModelingResourcePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (ModelingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelingArgs)(nil)).Elem()
}

type ModelingInput interface {
	pulumi.Input

	ToModelingOutput() ModelingOutput
	ToModelingOutputWithContext(ctx context.Context) ModelingOutput
}

func (*Modeling) ElementType() reflect.Type {
	return reflect.TypeOf((**Modeling)(nil)).Elem()
}

func (i *Modeling) ToModelingOutput() ModelingOutput {
	return i.ToModelingOutputWithContext(context.Background())
}

func (i *Modeling) ToModelingOutputWithContext(ctx context.Context) ModelingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelingOutput)
}

type ModelingOutput struct{ *pulumi.OutputState }

func (ModelingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Modeling)(nil)).Elem()
}

func (o ModelingOutput) ToModelingOutput() ModelingOutput {
	return o
}

func (o ModelingOutput) ToModelingOutputWithContext(ctx context.Context) ModelingOutput {
	return o
}

func (o ModelingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Modeling) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ModelingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Modeling) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ModelingOutput) Properties() ModelingResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *Modeling) ModelingResourceResponsePropertiesOutput { return v.Properties }).(ModelingResourceResponsePropertiesOutput)
}

func (o ModelingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Modeling) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ModelingOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Modeling) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ModelingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Modeling) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ModelingOutput{})
}

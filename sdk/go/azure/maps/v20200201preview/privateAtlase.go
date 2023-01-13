


package v20200201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateAtlase struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                  `pulumi:"location"`
	Name       pulumi.StringOutput                  `pulumi:"name"`
	Properties PrivateAtlasPropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput               `pulumi:"tags"`
	Type       pulumi.StringOutput                  `pulumi:"type"`
}


func NewPrivateAtlase(ctx *pulumi.Context,
	name string, args *PrivateAtlaseArgs, opts ...pulumi.ResourceOption) (*PrivateAtlase, error) {
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
			Type: pulumi.String("azure-native:maps:PrivateAtlase"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateAtlase
	err := ctx.RegisterResource("azure-native:maps/v20200201preview:PrivateAtlase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateAtlase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateAtlaseState, opts ...pulumi.ResourceOption) (*PrivateAtlase, error) {
	var resource PrivateAtlase
	err := ctx.ReadResource("azure-native:maps/v20200201preview:PrivateAtlase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateAtlaseState struct {
}

type PrivateAtlaseState struct {
}

func (PrivateAtlaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateAtlaseState)(nil)).Elem()
}

type privateAtlaseArgs struct {
	AccountName       string            `pulumi:"accountName"`
	Location          *string           `pulumi:"location"`
	PrivateAtlasName  *string           `pulumi:"privateAtlasName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type PrivateAtlaseArgs struct {
	AccountName       pulumi.StringInput
	Location          pulumi.StringPtrInput
	PrivateAtlasName  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (PrivateAtlaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateAtlaseArgs)(nil)).Elem()
}

type PrivateAtlaseInput interface {
	pulumi.Input

	ToPrivateAtlaseOutput() PrivateAtlaseOutput
	ToPrivateAtlaseOutputWithContext(ctx context.Context) PrivateAtlaseOutput
}

func (*PrivateAtlase) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateAtlase)(nil)).Elem()
}

func (i *PrivateAtlase) ToPrivateAtlaseOutput() PrivateAtlaseOutput {
	return i.ToPrivateAtlaseOutputWithContext(context.Background())
}

func (i *PrivateAtlase) ToPrivateAtlaseOutputWithContext(ctx context.Context) PrivateAtlaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateAtlaseOutput)
}

type PrivateAtlaseOutput struct{ *pulumi.OutputState }

func (PrivateAtlaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateAtlase)(nil)).Elem()
}

func (o PrivateAtlaseOutput) ToPrivateAtlaseOutput() PrivateAtlaseOutput {
	return o
}

func (o PrivateAtlaseOutput) ToPrivateAtlaseOutputWithContext(ctx context.Context) PrivateAtlaseOutput {
	return o
}

func (o PrivateAtlaseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateAtlase) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateAtlaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateAtlase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateAtlaseOutput) Properties() PrivateAtlasPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateAtlase) PrivateAtlasPropertiesResponseOutput { return v.Properties }).(PrivateAtlasPropertiesResponseOutput)
}

func (o PrivateAtlaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateAtlase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateAtlaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateAtlase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateAtlaseOutput{})
}

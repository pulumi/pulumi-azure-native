


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Target struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput   `pulumi:"location"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	Properties pulumi.AnyOutput         `pulumi:"properties"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewTarget(ctx *pulumi.Context,
	name string, args *TargetArgs, opts ...pulumi.ResourceOption) (*Target, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ParentProviderNamespace'")
	}
	if args.ParentResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourceName'")
	}
	if args.ParentResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourceType'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:chaos:Target"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20210915preview:Target"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20221001preview:Target"),
		},
	})
	opts = append(opts, aliases)
	var resource Target
	err := ctx.RegisterResource("azure-native:chaos/v20220701preview:Target", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetState, opts ...pulumi.ResourceOption) (*Target, error) {
	var resource Target
	err := ctx.ReadResource("azure-native:chaos/v20220701preview:Target", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type targetState struct {
}

type TargetState struct {
}

func (TargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetState)(nil)).Elem()
}

type targetArgs struct {
	Location                *string     `pulumi:"location"`
	ParentProviderNamespace string      `pulumi:"parentProviderNamespace"`
	ParentResourceName      string      `pulumi:"parentResourceName"`
	ParentResourceType      string      `pulumi:"parentResourceType"`
	Properties              interface{} `pulumi:"properties"`
	ResourceGroupName       string      `pulumi:"resourceGroupName"`
	TargetName              *string     `pulumi:"targetName"`
}


type TargetArgs struct {
	Location                pulumi.StringPtrInput
	ParentProviderNamespace pulumi.StringInput
	ParentResourceName      pulumi.StringInput
	ParentResourceType      pulumi.StringInput
	Properties              pulumi.Input
	ResourceGroupName       pulumi.StringInput
	TargetName              pulumi.StringPtrInput
}

func (TargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetArgs)(nil)).Elem()
}

type TargetInput interface {
	pulumi.Input

	ToTargetOutput() TargetOutput
	ToTargetOutputWithContext(ctx context.Context) TargetOutput
}

func (*Target) ElementType() reflect.Type {
	return reflect.TypeOf((**Target)(nil)).Elem()
}

func (i *Target) ToTargetOutput() TargetOutput {
	return i.ToTargetOutputWithContext(context.Background())
}

func (i *Target) ToTargetOutputWithContext(ctx context.Context) TargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetOutput)
}

type TargetOutput struct{ *pulumi.OutputState }

func (TargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Target)(nil)).Elem()
}

func (o TargetOutput) ToTargetOutput() TargetOutput {
	return o
}

func (o TargetOutput) ToTargetOutputWithContext(ctx context.Context) TargetOutput {
	return o
}

func (o TargetOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Target) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o TargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TargetOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Target) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o TargetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Target) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o TargetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Target) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TargetOutput{})
}

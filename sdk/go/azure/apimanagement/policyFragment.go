


package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyFragment struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	Format      pulumi.StringPtrOutput `pulumi:"format"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Type        pulumi.StringOutput    `pulumi:"type"`
	Value       pulumi.StringOutput    `pulumi:"value"`
}


func NewPolicyFragment(ctx *pulumi.Context,
	name string, args *PolicyFragmentArgs, opts ...pulumi.ResourceOption) (*PolicyFragment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	if isZero(args.Format) {
		args.Format = pulumi.StringPtr("xml")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:PolicyFragment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:PolicyFragment"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicyFragment
	err := ctx.RegisterResource("azure-native:apimanagement:PolicyFragment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicyFragment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyFragmentState, opts ...pulumi.ResourceOption) (*PolicyFragment, error) {
	var resource PolicyFragment
	err := ctx.ReadResource("azure-native:apimanagement:PolicyFragment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policyFragmentState struct {
}

type PolicyFragmentState struct {
}

func (PolicyFragmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyFragmentState)(nil)).Elem()
}

type policyFragmentArgs struct {
	Description       *string `pulumi:"description"`
	Format            *string `pulumi:"format"`
	Id                *string `pulumi:"id"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	Value             string  `pulumi:"value"`
}


type PolicyFragmentArgs struct {
	Description       pulumi.StringPtrInput
	Format            pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	Value             pulumi.StringInput
}

func (PolicyFragmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyFragmentArgs)(nil)).Elem()
}

type PolicyFragmentInput interface {
	pulumi.Input

	ToPolicyFragmentOutput() PolicyFragmentOutput
	ToPolicyFragmentOutputWithContext(ctx context.Context) PolicyFragmentOutput
}

func (*PolicyFragment) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyFragment)(nil)).Elem()
}

func (i *PolicyFragment) ToPolicyFragmentOutput() PolicyFragmentOutput {
	return i.ToPolicyFragmentOutputWithContext(context.Background())
}

func (i *PolicyFragment) ToPolicyFragmentOutputWithContext(ctx context.Context) PolicyFragmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyFragmentOutput)
}

type PolicyFragmentOutput struct{ *pulumi.OutputState }

func (PolicyFragmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyFragment)(nil)).Elem()
}

func (o PolicyFragmentOutput) ToPolicyFragmentOutput() PolicyFragmentOutput {
	return o
}

func (o PolicyFragmentOutput) ToPolicyFragmentOutputWithContext(ctx context.Context) PolicyFragmentOutput {
	return o
}

func (o PolicyFragmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyFragment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyFragmentOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyFragment) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

func (o PolicyFragmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyFragment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicyFragmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyFragment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PolicyFragmentOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyFragment) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyFragmentOutput{})
}

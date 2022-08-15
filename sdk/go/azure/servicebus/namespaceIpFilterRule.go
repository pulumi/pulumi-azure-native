


package servicebus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NamespaceIpFilterRule struct {
	pulumi.CustomResourceState

	Action     pulumi.StringPtrOutput `pulumi:"action"`
	FilterName pulumi.StringPtrOutput `pulumi:"filterName"`
	IpMask     pulumi.StringPtrOutput `pulumi:"ipMask"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewNamespaceIpFilterRule(ctx *pulumi.Context,
	name string, args *NamespaceIpFilterRuleArgs, opts ...pulumi.ResourceOption) (*NamespaceIpFilterRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:NamespaceIpFilterRule"),
		},
	})
	opts = append(opts, aliases)
	var resource NamespaceIpFilterRule
	err := ctx.RegisterResource("azure-native:servicebus:NamespaceIpFilterRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespaceIpFilterRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceIpFilterRuleState, opts ...pulumi.ResourceOption) (*NamespaceIpFilterRule, error) {
	var resource NamespaceIpFilterRule
	err := ctx.ReadResource("azure-native:servicebus:NamespaceIpFilterRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type namespaceIpFilterRuleState struct {
}

type NamespaceIpFilterRuleState struct {
}

func (NamespaceIpFilterRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceIpFilterRuleState)(nil)).Elem()
}

type namespaceIpFilterRuleArgs struct {
	Action            *string `pulumi:"action"`
	FilterName        *string `pulumi:"filterName"`
	IpFilterRuleName  *string `pulumi:"ipFilterRuleName"`
	IpMask            *string `pulumi:"ipMask"`
	NamespaceName     string  `pulumi:"namespaceName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type NamespaceIpFilterRuleArgs struct {
	Action            pulumi.StringPtrInput
	FilterName        pulumi.StringPtrInput
	IpFilterRuleName  pulumi.StringPtrInput
	IpMask            pulumi.StringPtrInput
	NamespaceName     pulumi.StringInput
	ResourceGroupName pulumi.StringInput
}

func (NamespaceIpFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceIpFilterRuleArgs)(nil)).Elem()
}

type NamespaceIpFilterRuleInput interface {
	pulumi.Input

	ToNamespaceIpFilterRuleOutput() NamespaceIpFilterRuleOutput
	ToNamespaceIpFilterRuleOutputWithContext(ctx context.Context) NamespaceIpFilterRuleOutput
}

func (*NamespaceIpFilterRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceIpFilterRule)(nil)).Elem()
}

func (i *NamespaceIpFilterRule) ToNamespaceIpFilterRuleOutput() NamespaceIpFilterRuleOutput {
	return i.ToNamespaceIpFilterRuleOutputWithContext(context.Background())
}

func (i *NamespaceIpFilterRule) ToNamespaceIpFilterRuleOutputWithContext(ctx context.Context) NamespaceIpFilterRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIpFilterRuleOutput)
}

type NamespaceIpFilterRuleOutput struct{ *pulumi.OutputState }

func (NamespaceIpFilterRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceIpFilterRule)(nil)).Elem()
}

func (o NamespaceIpFilterRuleOutput) ToNamespaceIpFilterRuleOutput() NamespaceIpFilterRuleOutput {
	return o
}

func (o NamespaceIpFilterRuleOutput) ToNamespaceIpFilterRuleOutputWithContext(ctx context.Context) NamespaceIpFilterRuleOutput {
	return o
}

func (o NamespaceIpFilterRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceIpFilterRule) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NamespaceIpFilterRuleOutput) FilterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceIpFilterRule) pulumi.StringPtrOutput { return v.FilterName }).(pulumi.StringPtrOutput)
}

func (o NamespaceIpFilterRuleOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceIpFilterRule) pulumi.StringPtrOutput { return v.IpMask }).(pulumi.StringPtrOutput)
}

func (o NamespaceIpFilterRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceIpFilterRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NamespaceIpFilterRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceIpFilterRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceIpFilterRuleOutput{})
}

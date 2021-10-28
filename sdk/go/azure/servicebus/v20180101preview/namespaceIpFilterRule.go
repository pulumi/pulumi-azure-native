


package v20180101preview

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
			Type: pulumi.String("azure-nextgen:servicebus/v20180101preview:NamespaceIpFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus:NamespaceIpFilterRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus:NamespaceIpFilterRule"),
		},
	})
	opts = append(opts, aliases)
	var resource NamespaceIpFilterRule
	err := ctx.RegisterResource("azure-native:servicebus/v20180101preview:NamespaceIpFilterRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespaceIpFilterRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceIpFilterRuleState, opts ...pulumi.ResourceOption) (*NamespaceIpFilterRule, error) {
	var resource NamespaceIpFilterRule
	err := ctx.ReadResource("azure-native:servicebus/v20180101preview:NamespaceIpFilterRule", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*NamespaceIpFilterRule)(nil))
}

func (i *NamespaceIpFilterRule) ToNamespaceIpFilterRuleOutput() NamespaceIpFilterRuleOutput {
	return i.ToNamespaceIpFilterRuleOutputWithContext(context.Background())
}

func (i *NamespaceIpFilterRule) ToNamespaceIpFilterRuleOutputWithContext(ctx context.Context) NamespaceIpFilterRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIpFilterRuleOutput)
}

type NamespaceIpFilterRuleOutput struct{ *pulumi.OutputState }

func (NamespaceIpFilterRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceIpFilterRule)(nil))
}

func (o NamespaceIpFilterRuleOutput) ToNamespaceIpFilterRuleOutput() NamespaceIpFilterRuleOutput {
	return o
}

func (o NamespaceIpFilterRuleOutput) ToNamespaceIpFilterRuleOutputWithContext(ctx context.Context) NamespaceIpFilterRuleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceIpFilterRuleInput)(nil)).Elem(), &NamespaceIpFilterRule{})
	pulumi.RegisterOutputType(NamespaceIpFilterRuleOutput{})
}

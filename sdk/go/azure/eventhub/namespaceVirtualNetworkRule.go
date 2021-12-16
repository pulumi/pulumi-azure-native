


package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NamespaceVirtualNetworkRule struct {
	pulumi.CustomResourceState

	Name                   pulumi.StringOutput    `pulumi:"name"`
	Type                   pulumi.StringOutput    `pulumi:"type"`
	VirtualNetworkSubnetId pulumi.StringPtrOutput `pulumi:"virtualNetworkSubnetId"`
}


func NewNamespaceVirtualNetworkRule(ctx *pulumi.Context,
	name string, args *NamespaceVirtualNetworkRuleArgs, opts ...pulumi.ResourceOption) (*NamespaceVirtualNetworkRule, error) {
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
			Type: pulumi.String("azure-native:eventhub/v20180101preview:NamespaceVirtualNetworkRule"),
		},
	})
	opts = append(opts, aliases)
	var resource NamespaceVirtualNetworkRule
	err := ctx.RegisterResource("azure-native:eventhub:NamespaceVirtualNetworkRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespaceVirtualNetworkRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceVirtualNetworkRuleState, opts ...pulumi.ResourceOption) (*NamespaceVirtualNetworkRule, error) {
	var resource NamespaceVirtualNetworkRule
	err := ctx.ReadResource("azure-native:eventhub:NamespaceVirtualNetworkRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type namespaceVirtualNetworkRuleState struct {
}

type NamespaceVirtualNetworkRuleState struct {
}

func (NamespaceVirtualNetworkRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceVirtualNetworkRuleState)(nil)).Elem()
}

type namespaceVirtualNetworkRuleArgs struct {
	NamespaceName          string  `pulumi:"namespaceName"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	VirtualNetworkRuleName *string `pulumi:"virtualNetworkRuleName"`
	VirtualNetworkSubnetId *string `pulumi:"virtualNetworkSubnetId"`
}


type NamespaceVirtualNetworkRuleArgs struct {
	NamespaceName          pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	VirtualNetworkRuleName pulumi.StringPtrInput
	VirtualNetworkSubnetId pulumi.StringPtrInput
}

func (NamespaceVirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceVirtualNetworkRuleArgs)(nil)).Elem()
}

type NamespaceVirtualNetworkRuleInput interface {
	pulumi.Input

	ToNamespaceVirtualNetworkRuleOutput() NamespaceVirtualNetworkRuleOutput
	ToNamespaceVirtualNetworkRuleOutputWithContext(ctx context.Context) NamespaceVirtualNetworkRuleOutput
}

func (*NamespaceVirtualNetworkRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceVirtualNetworkRule)(nil)).Elem()
}

func (i *NamespaceVirtualNetworkRule) ToNamespaceVirtualNetworkRuleOutput() NamespaceVirtualNetworkRuleOutput {
	return i.ToNamespaceVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i *NamespaceVirtualNetworkRule) ToNamespaceVirtualNetworkRuleOutputWithContext(ctx context.Context) NamespaceVirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceVirtualNetworkRuleOutput)
}

type NamespaceVirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (NamespaceVirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceVirtualNetworkRule)(nil)).Elem()
}

func (o NamespaceVirtualNetworkRuleOutput) ToNamespaceVirtualNetworkRuleOutput() NamespaceVirtualNetworkRuleOutput {
	return o
}

func (o NamespaceVirtualNetworkRuleOutput) ToNamespaceVirtualNetworkRuleOutputWithContext(ctx context.Context) NamespaceVirtualNetworkRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamespaceVirtualNetworkRuleOutput{})
}

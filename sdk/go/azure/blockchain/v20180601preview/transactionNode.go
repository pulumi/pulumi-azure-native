


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TransactionNode struct {
	pulumi.CustomResourceState

	Dns               pulumi.StringOutput             `pulumi:"dns"`
	FirewallRules     FirewallRuleResponseArrayOutput `pulumi:"firewallRules"`
	Location          pulumi.StringPtrOutput          `pulumi:"location"`
	Name              pulumi.StringOutput             `pulumi:"name"`
	Password          pulumi.StringPtrOutput          `pulumi:"password"`
	ProvisioningState pulumi.StringOutput             `pulumi:"provisioningState"`
	PublicKey         pulumi.StringOutput             `pulumi:"publicKey"`
	Type              pulumi.StringOutput             `pulumi:"type"`
	UserName          pulumi.StringOutput             `pulumi:"userName"`
}


func NewTransactionNode(ctx *pulumi.Context,
	name string, args *TransactionNodeArgs, opts ...pulumi.ResourceOption) (*TransactionNode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlockchainMemberName == nil {
		return nil, errors.New("invalid value for required argument 'BlockchainMemberName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:blockchain:TransactionNode"),
		},
	})
	opts = append(opts, aliases)
	var resource TransactionNode
	err := ctx.RegisterResource("azure-native:blockchain/v20180601preview:TransactionNode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTransactionNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransactionNodeState, opts ...pulumi.ResourceOption) (*TransactionNode, error) {
	var resource TransactionNode
	err := ctx.ReadResource("azure-native:blockchain/v20180601preview:TransactionNode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type transactionNodeState struct {
}

type TransactionNodeState struct {
}

func (TransactionNodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*transactionNodeState)(nil)).Elem()
}

type transactionNodeArgs struct {
	BlockchainMemberName string         `pulumi:"blockchainMemberName"`
	FirewallRules        []FirewallRule `pulumi:"firewallRules"`
	Location             *string        `pulumi:"location"`
	Password             *string        `pulumi:"password"`
	ResourceGroupName    string         `pulumi:"resourceGroupName"`
	TransactionNodeName  *string        `pulumi:"transactionNodeName"`
}


type TransactionNodeArgs struct {
	BlockchainMemberName pulumi.StringInput
	FirewallRules        FirewallRuleArrayInput
	Location             pulumi.StringPtrInput
	Password             pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	TransactionNodeName  pulumi.StringPtrInput
}

func (TransactionNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transactionNodeArgs)(nil)).Elem()
}

type TransactionNodeInput interface {
	pulumi.Input

	ToTransactionNodeOutput() TransactionNodeOutput
	ToTransactionNodeOutputWithContext(ctx context.Context) TransactionNodeOutput
}

func (*TransactionNode) ElementType() reflect.Type {
	return reflect.TypeOf((**TransactionNode)(nil)).Elem()
}

func (i *TransactionNode) ToTransactionNodeOutput() TransactionNodeOutput {
	return i.ToTransactionNodeOutputWithContext(context.Background())
}

func (i *TransactionNode) ToTransactionNodeOutputWithContext(ctx context.Context) TransactionNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransactionNodeOutput)
}

type TransactionNodeOutput struct{ *pulumi.OutputState }

func (TransactionNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransactionNode)(nil)).Elem()
}

func (o TransactionNodeOutput) ToTransactionNodeOutput() TransactionNodeOutput {
	return o
}

func (o TransactionNodeOutput) ToTransactionNodeOutputWithContext(ctx context.Context) TransactionNodeOutput {
	return o
}

func (o TransactionNodeOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *TransactionNode) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

func (o TransactionNodeOutput) FirewallRules() FirewallRuleResponseArrayOutput {
	return o.ApplyT(func(v *TransactionNode) FirewallRuleResponseArrayOutput { return v.FirewallRules }).(FirewallRuleResponseArrayOutput)
}

func (o TransactionNodeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransactionNode) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o TransactionNodeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TransactionNode) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TransactionNodeOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransactionNode) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o TransactionNodeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *TransactionNode) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o TransactionNodeOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *TransactionNode) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

func (o TransactionNodeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TransactionNode) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o TransactionNodeOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *TransactionNode) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TransactionNodeOutput{})
}

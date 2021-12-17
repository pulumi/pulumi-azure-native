


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlockchainMember struct {
	pulumi.CustomResourceState

	Consortium                          pulumi.StringPtrOutput                    `pulumi:"consortium"`
	ConsortiumManagementAccountAddress  pulumi.StringOutput                       `pulumi:"consortiumManagementAccountAddress"`
	ConsortiumManagementAccountPassword pulumi.StringPtrOutput                    `pulumi:"consortiumManagementAccountPassword"`
	ConsortiumMemberDisplayName         pulumi.StringPtrOutput                    `pulumi:"consortiumMemberDisplayName"`
	ConsortiumRole                      pulumi.StringPtrOutput                    `pulumi:"consortiumRole"`
	Dns                                 pulumi.StringOutput                       `pulumi:"dns"`
	FirewallRules                       FirewallRuleResponseArrayOutput           `pulumi:"firewallRules"`
	Location                            pulumi.StringPtrOutput                    `pulumi:"location"`
	Name                                pulumi.StringOutput                       `pulumi:"name"`
	Password                            pulumi.StringPtrOutput                    `pulumi:"password"`
	Protocol                            pulumi.StringPtrOutput                    `pulumi:"protocol"`
	ProvisioningState                   pulumi.StringOutput                       `pulumi:"provisioningState"`
	PublicKey                           pulumi.StringOutput                       `pulumi:"publicKey"`
	RootContractAddress                 pulumi.StringOutput                       `pulumi:"rootContractAddress"`
	Sku                                 SkuResponsePtrOutput                      `pulumi:"sku"`
	Tags                                pulumi.StringMapOutput                    `pulumi:"tags"`
	Type                                pulumi.StringOutput                       `pulumi:"type"`
	UserName                            pulumi.StringOutput                       `pulumi:"userName"`
	ValidatorNodesSku                   BlockchainMemberNodesSkuResponsePtrOutput `pulumi:"validatorNodesSku"`
}


func NewBlockchainMember(ctx *pulumi.Context,
	name string, args *BlockchainMemberArgs, opts ...pulumi.ResourceOption) (*BlockchainMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:blockchain:BlockchainMember"),
		},
	})
	opts = append(opts, aliases)
	var resource BlockchainMember
	err := ctx.RegisterResource("azure-native:blockchain/v20180601preview:BlockchainMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlockchainMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlockchainMemberState, opts ...pulumi.ResourceOption) (*BlockchainMember, error) {
	var resource BlockchainMember
	err := ctx.ReadResource("azure-native:blockchain/v20180601preview:BlockchainMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blockchainMemberState struct {
}

type BlockchainMemberState struct {
}

func (BlockchainMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*blockchainMemberState)(nil)).Elem()
}

type blockchainMemberArgs struct {
	BlockchainMemberName                *string                   `pulumi:"blockchainMemberName"`
	Consortium                          *string                   `pulumi:"consortium"`
	ConsortiumManagementAccountPassword *string                   `pulumi:"consortiumManagementAccountPassword"`
	ConsortiumMemberDisplayName         *string                   `pulumi:"consortiumMemberDisplayName"`
	ConsortiumRole                      *string                   `pulumi:"consortiumRole"`
	FirewallRules                       []FirewallRule            `pulumi:"firewallRules"`
	Location                            *string                   `pulumi:"location"`
	Password                            *string                   `pulumi:"password"`
	Protocol                            *string                   `pulumi:"protocol"`
	ResourceGroupName                   string                    `pulumi:"resourceGroupName"`
	Sku                                 *Sku                      `pulumi:"sku"`
	Tags                                map[string]string         `pulumi:"tags"`
	ValidatorNodesSku                   *BlockchainMemberNodesSku `pulumi:"validatorNodesSku"`
}


type BlockchainMemberArgs struct {
	BlockchainMemberName                pulumi.StringPtrInput
	Consortium                          pulumi.StringPtrInput
	ConsortiumManagementAccountPassword pulumi.StringPtrInput
	ConsortiumMemberDisplayName         pulumi.StringPtrInput
	ConsortiumRole                      pulumi.StringPtrInput
	FirewallRules                       FirewallRuleArrayInput
	Location                            pulumi.StringPtrInput
	Password                            pulumi.StringPtrInput
	Protocol                            pulumi.StringPtrInput
	ResourceGroupName                   pulumi.StringInput
	Sku                                 SkuPtrInput
	Tags                                pulumi.StringMapInput
	ValidatorNodesSku                   BlockchainMemberNodesSkuPtrInput
}

func (BlockchainMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blockchainMemberArgs)(nil)).Elem()
}

type BlockchainMemberInput interface {
	pulumi.Input

	ToBlockchainMemberOutput() BlockchainMemberOutput
	ToBlockchainMemberOutputWithContext(ctx context.Context) BlockchainMemberOutput
}

func (*BlockchainMember) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockchainMember)(nil)).Elem()
}

func (i *BlockchainMember) ToBlockchainMemberOutput() BlockchainMemberOutput {
	return i.ToBlockchainMemberOutputWithContext(context.Background())
}

func (i *BlockchainMember) ToBlockchainMemberOutputWithContext(ctx context.Context) BlockchainMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockchainMemberOutput)
}

type BlockchainMemberOutput struct{ *pulumi.OutputState }

func (BlockchainMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockchainMember)(nil)).Elem()
}

func (o BlockchainMemberOutput) ToBlockchainMemberOutput() BlockchainMemberOutput {
	return o
}

func (o BlockchainMemberOutput) ToBlockchainMemberOutputWithContext(ctx context.Context) BlockchainMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BlockchainMemberOutput{})
}

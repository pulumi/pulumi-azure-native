


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IPv6FirewallRule struct {
	pulumi.CustomResourceState

	EndIPv6Address   pulumi.StringPtrOutput `pulumi:"endIPv6Address"`
	Name             pulumi.StringPtrOutput `pulumi:"name"`
	StartIPv6Address pulumi.StringPtrOutput `pulumi:"startIPv6Address"`
	Type             pulumi.StringOutput    `pulumi:"type"`
}


func NewIPv6FirewallRule(ctx *pulumi.Context,
	name string, args *IPv6FirewallRuleArgs, opts ...pulumi.ResourceOption) (*IPv6FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:IPv6FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:IPv6FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:IPv6FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:IPv6FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:IPv6FirewallRule"),
		},
	})
	opts = append(opts, aliases)
	var resource IPv6FirewallRule
	err := ctx.RegisterResource("azure-native:sql/v20210801preview:IPv6FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIPv6FirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPv6FirewallRuleState, opts ...pulumi.ResourceOption) (*IPv6FirewallRule, error) {
	var resource IPv6FirewallRule
	err := ctx.ReadResource("azure-native:sql/v20210801preview:IPv6FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ipv6FirewallRuleState struct {
}

type IPv6FirewallRuleState struct {
}

func (IPv6FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6FirewallRuleState)(nil)).Elem()
}

type ipv6FirewallRuleArgs struct {
	EndIPv6Address    *string `pulumi:"endIPv6Address"`
	FirewallRuleName  *string `pulumi:"firewallRuleName"`
	Name              *string `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
	StartIPv6Address  *string `pulumi:"startIPv6Address"`
}


type IPv6FirewallRuleArgs struct {
	EndIPv6Address    pulumi.StringPtrInput
	FirewallRuleName  pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	StartIPv6Address  pulumi.StringPtrInput
}

func (IPv6FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6FirewallRuleArgs)(nil)).Elem()
}

type IPv6FirewallRuleInput interface {
	pulumi.Input

	ToIPv6FirewallRuleOutput() IPv6FirewallRuleOutput
	ToIPv6FirewallRuleOutputWithContext(ctx context.Context) IPv6FirewallRuleOutput
}

func (*IPv6FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**IPv6FirewallRule)(nil)).Elem()
}

func (i *IPv6FirewallRule) ToIPv6FirewallRuleOutput() IPv6FirewallRuleOutput {
	return i.ToIPv6FirewallRuleOutputWithContext(context.Background())
}

func (i *IPv6FirewallRule) ToIPv6FirewallRuleOutputWithContext(ctx context.Context) IPv6FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv6FirewallRuleOutput)
}

type IPv6FirewallRuleOutput struct{ *pulumi.OutputState }

func (IPv6FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPv6FirewallRule)(nil)).Elem()
}

func (o IPv6FirewallRuleOutput) ToIPv6FirewallRuleOutput() IPv6FirewallRuleOutput {
	return o
}

func (o IPv6FirewallRuleOutput) ToIPv6FirewallRuleOutputWithContext(ctx context.Context) IPv6FirewallRuleOutput {
	return o
}

func (o IPv6FirewallRuleOutput) EndIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPv6FirewallRule) pulumi.StringPtrOutput { return v.EndIPv6Address }).(pulumi.StringPtrOutput)
}

func (o IPv6FirewallRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPv6FirewallRule) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IPv6FirewallRuleOutput) StartIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPv6FirewallRule) pulumi.StringPtrOutput { return v.StartIPv6Address }).(pulumi.StringPtrOutput)
}

func (o IPv6FirewallRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IPv6FirewallRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IPv6FirewallRuleOutput{})
}

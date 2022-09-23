


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OutboundFirewallRule struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput `pulumi:"name"`
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	Type              pulumi.StringOutput `pulumi:"type"`
}


func NewOutboundFirewallRule(ctx *pulumi.Context,
	name string, args *OutboundFirewallRuleArgs, opts ...pulumi.ResourceOption) (*OutboundFirewallRule, error) {
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
			Type: pulumi.String("azure-native:sql:OutboundFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:OutboundFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:OutboundFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:OutboundFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:OutboundFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:OutboundFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:OutboundFirewallRule"),
		},
	})
	opts = append(opts, aliases)
	var resource OutboundFirewallRule
	err := ctx.RegisterResource("azure-native:sql/v20220201preview:OutboundFirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOutboundFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutboundFirewallRuleState, opts ...pulumi.ResourceOption) (*OutboundFirewallRule, error) {
	var resource OutboundFirewallRule
	err := ctx.ReadResource("azure-native:sql/v20220201preview:OutboundFirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type outboundFirewallRuleState struct {
}

type OutboundFirewallRuleState struct {
}

func (OutboundFirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*outboundFirewallRuleState)(nil)).Elem()
}

type outboundFirewallRuleArgs struct {
	OutboundRuleFqdn  *string `pulumi:"outboundRuleFqdn"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
}


type OutboundFirewallRuleArgs struct {
	OutboundRuleFqdn  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
}

func (OutboundFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outboundFirewallRuleArgs)(nil)).Elem()
}

type OutboundFirewallRuleInput interface {
	pulumi.Input

	ToOutboundFirewallRuleOutput() OutboundFirewallRuleOutput
	ToOutboundFirewallRuleOutputWithContext(ctx context.Context) OutboundFirewallRuleOutput
}

func (*OutboundFirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**OutboundFirewallRule)(nil)).Elem()
}

func (i *OutboundFirewallRule) ToOutboundFirewallRuleOutput() OutboundFirewallRuleOutput {
	return i.ToOutboundFirewallRuleOutputWithContext(context.Background())
}

func (i *OutboundFirewallRule) ToOutboundFirewallRuleOutputWithContext(ctx context.Context) OutboundFirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundFirewallRuleOutput)
}

type OutboundFirewallRuleOutput struct{ *pulumi.OutputState }

func (OutboundFirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutboundFirewallRule)(nil)).Elem()
}

func (o OutboundFirewallRuleOutput) ToOutboundFirewallRuleOutput() OutboundFirewallRuleOutput {
	return o
}

func (o OutboundFirewallRuleOutput) ToOutboundFirewallRuleOutputWithContext(ctx context.Context) OutboundFirewallRuleOutput {
	return o
}

func (o OutboundFirewallRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundFirewallRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OutboundFirewallRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundFirewallRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o OutboundFirewallRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundFirewallRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OutboundFirewallRuleOutput{})
}




package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallRule struct {
	pulumi.CustomResourceState

	EndIpAddress   pulumi.StringPtrOutput `pulumi:"endIpAddress"`
	Name           pulumi.StringPtrOutput `pulumi:"name"`
	StartIpAddress pulumi.StringPtrOutput `pulumi:"startIpAddress"`
	Type           pulumi.StringOutput    `pulumi:"type"`
}


func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
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
			Type: pulumi.String("azure-native:sql:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:FirewallRule"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallRule
	err := ctx.RegisterResource("azure-native:sql/v20211101preview:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("azure-native:sql/v20211101preview:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type firewallRuleState struct {
}

type FirewallRuleState struct {
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	EndIpAddress      *string `pulumi:"endIpAddress"`
	FirewallRuleName  *string `pulumi:"firewallRuleName"`
	Name              *string `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
	StartIpAddress    *string `pulumi:"startIpAddress"`
}


type FirewallRuleArgs struct {
	EndIpAddress      pulumi.StringPtrInput
	FirewallRuleName  pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	StartIpAddress    pulumi.StringPtrInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}

type FirewallRuleInput interface {
	pulumi.Input

	ToFirewallRuleOutput() FirewallRuleOutput
	ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput
}

func (*FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (i *FirewallRule) ToFirewallRuleOutput() FirewallRuleOutput {
	return i.ToFirewallRuleOutputWithContext(context.Background())
}

func (i *FirewallRule) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleOutput)
}

type FirewallRuleOutput struct{ *pulumi.OutputState }

func (FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (o FirewallRuleOutput) ToFirewallRuleOutput() FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) EndIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.EndIpAddress }).(pulumi.StringPtrOutput)
}

func (o FirewallRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FirewallRuleOutput) StartIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.StartIpAddress }).(pulumi.StringPtrOutput)
}

func (o FirewallRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallRuleOutput{})
}

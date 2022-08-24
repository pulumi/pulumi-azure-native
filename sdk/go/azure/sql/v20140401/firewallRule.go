


package v20140401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallRule struct {
	pulumi.CustomResourceState

	EndIpAddress   pulumi.StringOutput `pulumi:"endIpAddress"`
	Kind           pulumi.StringOutput `pulumi:"kind"`
	Location       pulumi.StringOutput `pulumi:"location"`
	Name           pulumi.StringOutput `pulumi:"name"`
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
	Type           pulumi.StringOutput `pulumi:"type"`
}


func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndIpAddress'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.StartIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartIpAddress'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:FirewallRule"),
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
			Type: pulumi.String("azure-native:sql/v20211101preview:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:FirewallRule"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallRule
	err := ctx.RegisterResource("azure-native:sql/v20140401:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("azure-native:sql/v20140401:FirewallRule", name, id, state, &resource, opts...)
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
	EndIpAddress      string  `pulumi:"endIpAddress"`
	FirewallRuleName  *string `pulumi:"firewallRuleName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
	StartIpAddress    string  `pulumi:"startIpAddress"`
}


type FirewallRuleArgs struct {
	EndIpAddress      pulumi.StringInput
	FirewallRuleName  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	StartIpAddress    pulumi.StringInput
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

func (o FirewallRuleOutput) EndIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.EndIpAddress }).(pulumi.StringOutput)
}

func (o FirewallRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o FirewallRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o FirewallRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallRuleOutput) StartIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.StartIpAddress }).(pulumi.StringOutput)
}

func (o FirewallRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallRuleOutput{})
}

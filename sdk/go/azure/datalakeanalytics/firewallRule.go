


package datalakeanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallRule struct {
	pulumi.CustomResourceState

	EndIpAddress   pulumi.StringOutput `pulumi:"endIpAddress"`
	Name           pulumi.StringOutput `pulumi:"name"`
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
	Type           pulumi.StringOutput `pulumi:"type"`
}


func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.EndIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndIpAddress'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StartIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartIpAddress'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20151001preview:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20161101:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20191101preview:FirewallRule"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallRule
	err := ctx.RegisterResource("azure-native:datalakeanalytics:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("azure-native:datalakeanalytics:FirewallRule", name, id, state, &resource, opts...)
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
	AccountName       string  `pulumi:"accountName"`
	EndIpAddress      string  `pulumi:"endIpAddress"`
	FirewallRuleName  *string `pulumi:"firewallRuleName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	StartIpAddress    string  `pulumi:"startIpAddress"`
}


type FirewallRuleArgs struct {
	AccountName       pulumi.StringInput
	EndIpAddress      pulumi.StringInput
	FirewallRuleName  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
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

func init() {
	pulumi.RegisterOutputType(FirewallRuleOutput{})
}




package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallRule struct {
	pulumi.CustomResourceState

	EndIP   pulumi.StringOutput `pulumi:"endIP"`
	Name    pulumi.StringOutput `pulumi:"name"`
	StartIP pulumi.StringOutput `pulumi:"startIP"`
	Type    pulumi.StringOutput `pulumi:"type"`
}


func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CacheName == nil {
		return nil, errors.New("invalid value for required argument 'CacheName'")
	}
	if args.EndIP == nil {
		return nil, errors.New("invalid value for required argument 'EndIP'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StartIP == nil {
		return nil, errors.New("invalid value for required argument 'StartIP'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cache:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20160401:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20170201:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20171001:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20180301:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20190701:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20200601:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201201:FirewallRule"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallRule
	err := ctx.RegisterResource("azure-native:cache/v20210601:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("azure-native:cache/v20210601:FirewallRule", name, id, state, &resource, opts...)
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
	CacheName         string  `pulumi:"cacheName"`
	EndIP             string  `pulumi:"endIP"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RuleName          *string `pulumi:"ruleName"`
	StartIP           string  `pulumi:"startIP"`
}


type FirewallRuleArgs struct {
	CacheName         pulumi.StringInput
	EndIP             pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	RuleName          pulumi.StringPtrInput
	StartIP           pulumi.StringInput
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

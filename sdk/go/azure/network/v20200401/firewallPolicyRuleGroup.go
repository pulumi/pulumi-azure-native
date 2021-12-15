


package v20200401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallPolicyRuleGroup struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput    `pulumi:"etag"`
	Name              pulumi.StringPtrOutput `pulumi:"name"`
	Priority          pulumi.IntPtrOutput    `pulumi:"priority"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Rules             pulumi.ArrayOutput     `pulumi:"rules"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewFirewallPolicyRuleGroup(ctx *pulumi.Context,
	name string, args *FirewallPolicyRuleGroupArgs, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicyName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:FirewallPolicyRuleGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallPolicyRuleGroup
	err := ctx.RegisterResource("azure-native:network/v20200401:FirewallPolicyRuleGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFirewallPolicyRuleGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyRuleGroupState, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleGroup, error) {
	var resource FirewallPolicyRuleGroup
	err := ctx.ReadResource("azure-native:network/v20200401:FirewallPolicyRuleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type firewallPolicyRuleGroupState struct {
}

type FirewallPolicyRuleGroupState struct {
}

func (FirewallPolicyRuleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleGroupState)(nil)).Elem()
}

type firewallPolicyRuleGroupArgs struct {
	FirewallPolicyName string        `pulumi:"firewallPolicyName"`
	Id                 *string       `pulumi:"id"`
	Name               *string       `pulumi:"name"`
	Priority           *int          `pulumi:"priority"`
	ResourceGroupName  string        `pulumi:"resourceGroupName"`
	RuleGroupName      *string       `pulumi:"ruleGroupName"`
	Rules              []interface{} `pulumi:"rules"`
}


type FirewallPolicyRuleGroupArgs struct {
	FirewallPolicyName pulumi.StringInput
	Id                 pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	Priority           pulumi.IntPtrInput
	ResourceGroupName  pulumi.StringInput
	RuleGroupName      pulumi.StringPtrInput
	Rules              pulumi.ArrayInput
}

func (FirewallPolicyRuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleGroupArgs)(nil)).Elem()
}

type FirewallPolicyRuleGroupInput interface {
	pulumi.Input

	ToFirewallPolicyRuleGroupOutput() FirewallPolicyRuleGroupOutput
	ToFirewallPolicyRuleGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleGroupOutput
}

func (*FirewallPolicyRuleGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleGroup)(nil)).Elem()
}

func (i *FirewallPolicyRuleGroup) ToFirewallPolicyRuleGroupOutput() FirewallPolicyRuleGroupOutput {
	return i.ToFirewallPolicyRuleGroupOutputWithContext(context.Background())
}

func (i *FirewallPolicyRuleGroup) ToFirewallPolicyRuleGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleGroupOutput)
}

type FirewallPolicyRuleGroupOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleGroup)(nil)).Elem()
}

func (o FirewallPolicyRuleGroupOutput) ToFirewallPolicyRuleGroupOutput() FirewallPolicyRuleGroupOutput {
	return o
}

func (o FirewallPolicyRuleGroupOutput) ToFirewallPolicyRuleGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyRuleGroupOutput{})
}

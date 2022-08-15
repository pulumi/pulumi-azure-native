


package v20200801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallPolicyRuleCollectionGroup struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput    `pulumi:"etag"`
	Name              pulumi.StringPtrOutput `pulumi:"name"`
	Priority          pulumi.IntPtrOutput    `pulumi:"priority"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	RuleCollections   pulumi.ArrayOutput     `pulumi:"ruleCollections"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewFirewallPolicyRuleCollectionGroup(ctx *pulumi.Context,
	name string, args *FirewallPolicyRuleCollectionGroupArgs, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleCollectionGroup, error) {
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
			Type: pulumi.String("azure-native:network:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:FirewallPolicyRuleCollectionGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallPolicyRuleCollectionGroup
	err := ctx.RegisterResource("azure-native:network/v20200801:FirewallPolicyRuleCollectionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFirewallPolicyRuleCollectionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyRuleCollectionGroupState, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleCollectionGroup, error) {
	var resource FirewallPolicyRuleCollectionGroup
	err := ctx.ReadResource("azure-native:network/v20200801:FirewallPolicyRuleCollectionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type firewallPolicyRuleCollectionGroupState struct {
}

type FirewallPolicyRuleCollectionGroupState struct {
}

func (FirewallPolicyRuleCollectionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleCollectionGroupState)(nil)).Elem()
}

type firewallPolicyRuleCollectionGroupArgs struct {
	FirewallPolicyName      string        `pulumi:"firewallPolicyName"`
	Id                      *string       `pulumi:"id"`
	Name                    *string       `pulumi:"name"`
	Priority                *int          `pulumi:"priority"`
	ResourceGroupName       string        `pulumi:"resourceGroupName"`
	RuleCollectionGroupName *string       `pulumi:"ruleCollectionGroupName"`
	RuleCollections         []interface{} `pulumi:"ruleCollections"`
}


type FirewallPolicyRuleCollectionGroupArgs struct {
	FirewallPolicyName      pulumi.StringInput
	Id                      pulumi.StringPtrInput
	Name                    pulumi.StringPtrInput
	Priority                pulumi.IntPtrInput
	ResourceGroupName       pulumi.StringInput
	RuleCollectionGroupName pulumi.StringPtrInput
	RuleCollections         pulumi.ArrayInput
}

func (FirewallPolicyRuleCollectionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleCollectionGroupArgs)(nil)).Elem()
}

type FirewallPolicyRuleCollectionGroupInput interface {
	pulumi.Input

	ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput
	ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput
}

func (*FirewallPolicyRuleCollectionGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleCollectionGroup)(nil)).Elem()
}

func (i *FirewallPolicyRuleCollectionGroup) ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput {
	return i.ToFirewallPolicyRuleCollectionGroupOutputWithContext(context.Background())
}

func (i *FirewallPolicyRuleCollectionGroup) ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleCollectionGroupOutput)
}

type FirewallPolicyRuleCollectionGroupOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleCollectionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleCollectionGroup)(nil)).Elem()
}

func (o FirewallPolicyRuleCollectionGroupOutput) ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput {
	return o
}

func (o FirewallPolicyRuleCollectionGroupOutput) ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput {
	return o
}

func (o FirewallPolicyRuleCollectionGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o FirewallPolicyRuleCollectionGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FirewallPolicyRuleCollectionGroupOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o FirewallPolicyRuleCollectionGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o FirewallPolicyRuleCollectionGroupOutput) RuleCollections() pulumi.ArrayOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.ArrayOutput { return v.RuleCollections }).(pulumi.ArrayOutput)
}

func (o FirewallPolicyRuleCollectionGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyRuleCollectionGroupOutput{})
}

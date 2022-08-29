


package v20200401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallPolicy struct {
	pulumi.CustomResourceState

	BasePolicy           SubResourceResponsePtrOutput                        `pulumi:"basePolicy"`
	ChildPolicies        SubResourceResponseArrayOutput                      `pulumi:"childPolicies"`
	Etag                 pulumi.StringOutput                                 `pulumi:"etag"`
	Firewalls            SubResourceResponseArrayOutput                      `pulumi:"firewalls"`
	Location             pulumi.StringPtrOutput                              `pulumi:"location"`
	Name                 pulumi.StringOutput                                 `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput                                 `pulumi:"provisioningState"`
	RuleGroups           SubResourceResponseArrayOutput                      `pulumi:"ruleGroups"`
	Tags                 pulumi.StringMapOutput                              `pulumi:"tags"`
	ThreatIntelMode      pulumi.StringPtrOutput                              `pulumi:"threatIntelMode"`
	ThreatIntelWhitelist FirewallPolicyThreatIntelWhitelistResponsePtrOutput `pulumi:"threatIntelWhitelist"`
	Type                 pulumi.StringOutput                                 `pulumi:"type"`
}


func NewFirewallPolicy(ctx *pulumi.Context,
	name string, args *FirewallPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:FirewallPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallPolicy
	err := ctx.RegisterResource("azure-native:network/v20200401:FirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyState, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	var resource FirewallPolicy
	err := ctx.ReadResource("azure-native:network/v20200401:FirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type firewallPolicyState struct {
}

type FirewallPolicyState struct {
}

func (FirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyState)(nil)).Elem()
}

type firewallPolicyArgs struct {
	BasePolicy           *SubResource                        `pulumi:"basePolicy"`
	FirewallPolicyName   *string                             `pulumi:"firewallPolicyName"`
	Id                   *string                             `pulumi:"id"`
	Location             *string                             `pulumi:"location"`
	ResourceGroupName    string                              `pulumi:"resourceGroupName"`
	Tags                 map[string]string                   `pulumi:"tags"`
	ThreatIntelMode      *string                             `pulumi:"threatIntelMode"`
	ThreatIntelWhitelist *FirewallPolicyThreatIntelWhitelist `pulumi:"threatIntelWhitelist"`
}


type FirewallPolicyArgs struct {
	BasePolicy           SubResourcePtrInput
	FirewallPolicyName   pulumi.StringPtrInput
	Id                   pulumi.StringPtrInput
	Location             pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
	ThreatIntelMode      pulumi.StringPtrInput
	ThreatIntelWhitelist FirewallPolicyThreatIntelWhitelistPtrInput
}

func (FirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyArgs)(nil)).Elem()
}

type FirewallPolicyInput interface {
	pulumi.Input

	ToFirewallPolicyOutput() FirewallPolicyOutput
	ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput
}

func (*FirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil)).Elem()
}

func (i *FirewallPolicy) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return i.ToFirewallPolicyOutputWithContext(context.Background())
}

func (i *FirewallPolicy) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyOutput)
}

type FirewallPolicyOutput struct{ *pulumi.OutputState }

func (FirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil)).Elem()
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) BasePolicy() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) SubResourceResponsePtrOutput { return v.BasePolicy }).(SubResourceResponsePtrOutput)
}

func (o FirewallPolicyOutput) ChildPolicies() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *FirewallPolicy) SubResourceResponseArrayOutput { return v.ChildPolicies }).(SubResourceResponseArrayOutput)
}

func (o FirewallPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o FirewallPolicyOutput) Firewalls() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *FirewallPolicy) SubResourceResponseArrayOutput { return v.Firewalls }).(SubResourceResponseArrayOutput)
}

func (o FirewallPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o FirewallPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o FirewallPolicyOutput) RuleGroups() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *FirewallPolicy) SubResourceResponseArrayOutput { return v.RuleGroups }).(SubResourceResponseArrayOutput)
}

func (o FirewallPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o FirewallPolicyOutput) ThreatIntelMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringPtrOutput { return v.ThreatIntelMode }).(pulumi.StringPtrOutput)
}

func (o FirewallPolicyOutput) ThreatIntelWhitelist() FirewallPolicyThreatIntelWhitelistResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) FirewallPolicyThreatIntelWhitelistResponsePtrOutput {
		return v.ThreatIntelWhitelist
	}).(FirewallPolicyThreatIntelWhitelistResponsePtrOutput)
}

func (o FirewallPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyOutput{})
}

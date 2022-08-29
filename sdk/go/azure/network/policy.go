


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Policy struct {
	pulumi.CustomResourceState

	CustomRules           CustomRuleListResponsePtrOutput          `pulumi:"customRules"`
	Etag                  pulumi.StringPtrOutput                   `pulumi:"etag"`
	FrontendEndpointLinks FrontendEndpointLinkResponseArrayOutput  `pulumi:"frontendEndpointLinks"`
	Location              pulumi.StringPtrOutput                   `pulumi:"location"`
	ManagedRules          ManagedRuleSetListResponsePtrOutput      `pulumi:"managedRules"`
	Name                  pulumi.StringOutput                      `pulumi:"name"`
	PolicySettings        FrontDoorPolicySettingsResponsePtrOutput `pulumi:"policySettings"`
	ProvisioningState     pulumi.StringOutput                      `pulumi:"provisioningState"`
	ResourceState         pulumi.StringOutput                      `pulumi:"resourceState"`
	RoutingRuleLinks      RoutingRuleLinkResponseArrayOutput       `pulumi:"routingRuleLinks"`
	SecurityPolicyLinks   SecurityPolicyLinkResponseArrayOutput    `pulumi:"securityPolicyLinks"`
	Sku                   SkuResponsePtrOutput                     `pulumi:"sku"`
	Tags                  pulumi.StringMapOutput                   `pulumi:"tags"`
	Type                  pulumi.StringOutput                      `pulumi:"type"`
}


func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20190301:Policy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191001:Policy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:Policy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:Policy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210601:Policy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:Policy"),
		},
	})
	opts = append(opts, aliases)
	var resource Policy
	err := ctx.RegisterResource("azure-native:network:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("azure-native:network:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policyState struct {
}

type PolicyState struct {
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	CustomRules       *CustomRuleList          `pulumi:"customRules"`
	Location          *string                  `pulumi:"location"`
	ManagedRules      *ManagedRuleSetList      `pulumi:"managedRules"`
	PolicyName        *string                  `pulumi:"policyName"`
	PolicySettings    *FrontDoorPolicySettings `pulumi:"policySettings"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	Sku               *Sku                     `pulumi:"sku"`
	Tags              map[string]string        `pulumi:"tags"`
}


type PolicyArgs struct {
	CustomRules       CustomRuleListPtrInput
	Location          pulumi.StringPtrInput
	ManagedRules      ManagedRuleSetListPtrInput
	PolicyName        pulumi.StringPtrInput
	PolicySettings    FrontDoorPolicySettingsPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func (o PolicyOutput) CustomRules() CustomRuleListResponsePtrOutput {
	return o.ApplyT(func(v *Policy) CustomRuleListResponsePtrOutput { return v.CustomRules }).(CustomRuleListResponsePtrOutput)
}

func (o PolicyOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PolicyOutput) FrontendEndpointLinks() FrontendEndpointLinkResponseArrayOutput {
	return o.ApplyT(func(v *Policy) FrontendEndpointLinkResponseArrayOutput { return v.FrontendEndpointLinks }).(FrontendEndpointLinkResponseArrayOutput)
}

func (o PolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PolicyOutput) ManagedRules() ManagedRuleSetListResponsePtrOutput {
	return o.ApplyT(func(v *Policy) ManagedRuleSetListResponsePtrOutput { return v.ManagedRules }).(ManagedRuleSetListResponsePtrOutput)
}

func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicyOutput) PolicySettings() FrontDoorPolicySettingsResponsePtrOutput {
	return o.ApplyT(func(v *Policy) FrontDoorPolicySettingsResponsePtrOutput { return v.PolicySettings }).(FrontDoorPolicySettingsResponsePtrOutput)
}

func (o PolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PolicyOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o PolicyOutput) RoutingRuleLinks() RoutingRuleLinkResponseArrayOutput {
	return o.ApplyT(func(v *Policy) RoutingRuleLinkResponseArrayOutput { return v.RoutingRuleLinks }).(RoutingRuleLinkResponseArrayOutput)
}

func (o PolicyOutput) SecurityPolicyLinks() SecurityPolicyLinkResponseArrayOutput {
	return o.ApplyT(func(v *Policy) SecurityPolicyLinkResponseArrayOutput { return v.SecurityPolicyLinks }).(SecurityPolicyLinkResponseArrayOutput)
}

func (o PolicyOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Policy) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o PolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
}

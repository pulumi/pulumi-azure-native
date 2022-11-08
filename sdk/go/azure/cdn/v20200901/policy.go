


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Policy struct {
	pulumi.CustomResourceState

	CustomRules       CustomRuleListResponsePtrOutput     `pulumi:"customRules"`
	EndpointLinks     CdnEndpointResponseArrayOutput      `pulumi:"endpointLinks"`
	Etag              pulumi.StringPtrOutput              `pulumi:"etag"`
	Location          pulumi.StringOutput                 `pulumi:"location"`
	ManagedRules      ManagedRuleSetListResponsePtrOutput `pulumi:"managedRules"`
	Name              pulumi.StringOutput                 `pulumi:"name"`
	PolicySettings    PolicySettingsResponsePtrOutput     `pulumi:"policySettings"`
	ProvisioningState pulumi.StringOutput                 `pulumi:"provisioningState"`
	RateLimitRules    RateLimitRuleListResponsePtrOutput  `pulumi:"rateLimitRules"`
	ResourceState     pulumi.StringOutput                 `pulumi:"resourceState"`
	Sku               SkuResponseOutput                   `pulumi:"sku"`
	SystemData        SystemDataResponseOutput            `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput              `pulumi:"tags"`
	Type              pulumi.StringOutput                 `pulumi:"type"`
}


func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:Policy"),
		},
	})
	opts = append(opts, aliases)
	var resource Policy
	err := ctx.RegisterResource("azure-native:cdn/v20200901:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("azure-native:cdn/v20200901:Policy", name, id, state, &resource, opts...)
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
	CustomRules       *CustomRuleList     `pulumi:"customRules"`
	Location          *string             `pulumi:"location"`
	ManagedRules      *ManagedRuleSetList `pulumi:"managedRules"`
	PolicyName        *string             `pulumi:"policyName"`
	PolicySettings    *PolicySettings     `pulumi:"policySettings"`
	RateLimitRules    *RateLimitRuleList  `pulumi:"rateLimitRules"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	Sku               Sku                 `pulumi:"sku"`
	Tags              map[string]string   `pulumi:"tags"`
}


type PolicyArgs struct {
	CustomRules       CustomRuleListPtrInput
	Location          pulumi.StringPtrInput
	ManagedRules      ManagedRuleSetListPtrInput
	PolicyName        pulumi.StringPtrInput
	PolicySettings    PolicySettingsPtrInput
	RateLimitRules    RateLimitRuleListPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuInput
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

func (o PolicyOutput) EndpointLinks() CdnEndpointResponseArrayOutput {
	return o.ApplyT(func(v *Policy) CdnEndpointResponseArrayOutput { return v.EndpointLinks }).(CdnEndpointResponseArrayOutput)
}

func (o PolicyOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PolicyOutput) ManagedRules() ManagedRuleSetListResponsePtrOutput {
	return o.ApplyT(func(v *Policy) ManagedRuleSetListResponsePtrOutput { return v.ManagedRules }).(ManagedRuleSetListResponsePtrOutput)
}

func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicyOutput) PolicySettings() PolicySettingsResponsePtrOutput {
	return o.ApplyT(func(v *Policy) PolicySettingsResponsePtrOutput { return v.PolicySettings }).(PolicySettingsResponsePtrOutput)
}

func (o PolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PolicyOutput) RateLimitRules() RateLimitRuleListResponsePtrOutput {
	return o.ApplyT(func(v *Policy) RateLimitRuleListResponsePtrOutput { return v.RateLimitRules }).(RateLimitRuleListResponsePtrOutput)
}

func (o PolicyOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o PolicyOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Policy) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o PolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Policy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
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

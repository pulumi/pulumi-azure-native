


package v20190201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebApplicationFirewallPolicy struct {
	pulumi.CustomResourceState

	ApplicationGateways ApplicationGatewayResponseArrayOutput               `pulumi:"applicationGateways"`
	CustomRules         WebApplicationFirewallCustomRuleResponseArrayOutput `pulumi:"customRules"`
	Etag                pulumi.StringPtrOutput                              `pulumi:"etag"`
	Location            pulumi.StringPtrOutput                              `pulumi:"location"`
	Name                pulumi.StringOutput                                 `pulumi:"name"`
	PolicySettings      PolicySettingsResponsePtrOutput                     `pulumi:"policySettings"`
	ProvisioningState   pulumi.StringOutput                                 `pulumi:"provisioningState"`
	ResourceState       pulumi.StringOutput                                 `pulumi:"resourceState"`
	Tags                pulumi.StringMapOutput                              `pulumi:"tags"`
	Type                pulumi.StringOutput                                 `pulumi:"type"`
}


func NewWebApplicationFirewallPolicy(ctx *pulumi.Context,
	name string, args *WebApplicationFirewallPolicyArgs, opts ...pulumi.ResourceOption) (*WebApplicationFirewallPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:WebApplicationFirewallPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource WebApplicationFirewallPolicy
	err := ctx.RegisterResource("azure-native:network/v20190201:WebApplicationFirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebApplicationFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebApplicationFirewallPolicyState, opts ...pulumi.ResourceOption) (*WebApplicationFirewallPolicy, error) {
	var resource WebApplicationFirewallPolicy
	err := ctx.ReadResource("azure-native:network/v20190201:WebApplicationFirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webApplicationFirewallPolicyState struct {
}

type WebApplicationFirewallPolicyState struct {
}

func (WebApplicationFirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*webApplicationFirewallPolicyState)(nil)).Elem()
}

type webApplicationFirewallPolicyArgs struct {
	CustomRules       []WebApplicationFirewallCustomRule `pulumi:"customRules"`
	Id                *string                            `pulumi:"id"`
	Location          *string                            `pulumi:"location"`
	PolicyName        *string                            `pulumi:"policyName"`
	PolicySettings    *PolicySettings                    `pulumi:"policySettings"`
	ResourceGroupName string                             `pulumi:"resourceGroupName"`
	Tags              map[string]string                  `pulumi:"tags"`
}


type WebApplicationFirewallPolicyArgs struct {
	CustomRules       WebApplicationFirewallCustomRuleArrayInput
	Id                pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	PolicyName        pulumi.StringPtrInput
	PolicySettings    PolicySettingsPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (WebApplicationFirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webApplicationFirewallPolicyArgs)(nil)).Elem()
}

type WebApplicationFirewallPolicyInput interface {
	pulumi.Input

	ToWebApplicationFirewallPolicyOutput() WebApplicationFirewallPolicyOutput
	ToWebApplicationFirewallPolicyOutputWithContext(ctx context.Context) WebApplicationFirewallPolicyOutput
}

func (*WebApplicationFirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApplicationFirewallPolicy)(nil)).Elem()
}

func (i *WebApplicationFirewallPolicy) ToWebApplicationFirewallPolicyOutput() WebApplicationFirewallPolicyOutput {
	return i.ToWebApplicationFirewallPolicyOutputWithContext(context.Background())
}

func (i *WebApplicationFirewallPolicy) ToWebApplicationFirewallPolicyOutputWithContext(ctx context.Context) WebApplicationFirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebApplicationFirewallPolicyOutput)
}

type WebApplicationFirewallPolicyOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApplicationFirewallPolicy)(nil)).Elem()
}

func (o WebApplicationFirewallPolicyOutput) ToWebApplicationFirewallPolicyOutput() WebApplicationFirewallPolicyOutput {
	return o
}

func (o WebApplicationFirewallPolicyOutput) ToWebApplicationFirewallPolicyOutputWithContext(ctx context.Context) WebApplicationFirewallPolicyOutput {
	return o
}

func (o WebApplicationFirewallPolicyOutput) ApplicationGateways() ApplicationGatewayResponseArrayOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) ApplicationGatewayResponseArrayOutput {
		return v.ApplicationGateways
	}).(ApplicationGatewayResponseArrayOutput)
}

func (o WebApplicationFirewallPolicyOutput) CustomRules() WebApplicationFirewallCustomRuleResponseArrayOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) WebApplicationFirewallCustomRuleResponseArrayOutput {
		return v.CustomRules
	}).(WebApplicationFirewallCustomRuleResponseArrayOutput)
}

func (o WebApplicationFirewallPolicyOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o WebApplicationFirewallPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o WebApplicationFirewallPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebApplicationFirewallPolicyOutput) PolicySettings() PolicySettingsResponsePtrOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) PolicySettingsResponsePtrOutput { return v.PolicySettings }).(PolicySettingsResponsePtrOutput)
}

func (o WebApplicationFirewallPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WebApplicationFirewallPolicyOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o WebApplicationFirewallPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WebApplicationFirewallPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebApplicationFirewallPolicyOutput{})
}

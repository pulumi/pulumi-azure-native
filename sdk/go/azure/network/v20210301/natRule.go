


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NatRule struct {
	pulumi.CustomResourceState

	EgressVpnSiteLinkConnections  SubResourceResponseArrayOutput       `pulumi:"egressVpnSiteLinkConnections"`
	Etag                          pulumi.StringOutput                  `pulumi:"etag"`
	ExternalMappings              VpnNatRuleMappingResponseArrayOutput `pulumi:"externalMappings"`
	IngressVpnSiteLinkConnections SubResourceResponseArrayOutput       `pulumi:"ingressVpnSiteLinkConnections"`
	InternalMappings              VpnNatRuleMappingResponseArrayOutput `pulumi:"internalMappings"`
	IpConfigurationId             pulumi.StringPtrOutput               `pulumi:"ipConfigurationId"`
	Mode                          pulumi.StringPtrOutput               `pulumi:"mode"`
	Name                          pulumi.StringPtrOutput               `pulumi:"name"`
	ProvisioningState             pulumi.StringOutput                  `pulumi:"provisioningState"`
	Type                          pulumi.StringOutput                  `pulumi:"type"`
}


func NewNatRule(ctx *pulumi.Context,
	name string, args *NatRuleArgs, opts ...pulumi.ResourceOption) (*NatRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayName == nil {
		return nil, errors.New("invalid value for required argument 'GatewayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:NatRule"),
		},
	})
	opts = append(opts, aliases)
	var resource NatRule
	err := ctx.RegisterResource("azure-native:network/v20210301:NatRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNatRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatRuleState, opts ...pulumi.ResourceOption) (*NatRule, error) {
	var resource NatRule
	err := ctx.ReadResource("azure-native:network/v20210301:NatRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type natRuleState struct {
}

type NatRuleState struct {
}

func (NatRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*natRuleState)(nil)).Elem()
}

type natRuleArgs struct {
	ExternalMappings  []VpnNatRuleMapping `pulumi:"externalMappings"`
	GatewayName       string              `pulumi:"gatewayName"`
	Id                *string             `pulumi:"id"`
	InternalMappings  []VpnNatRuleMapping `pulumi:"internalMappings"`
	IpConfigurationId *string             `pulumi:"ipConfigurationId"`
	Mode              *string             `pulumi:"mode"`
	Name              *string             `pulumi:"name"`
	NatRuleName       *string             `pulumi:"natRuleName"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	Type              *string             `pulumi:"type"`
}


type NatRuleArgs struct {
	ExternalMappings  VpnNatRuleMappingArrayInput
	GatewayName       pulumi.StringInput
	Id                pulumi.StringPtrInput
	InternalMappings  VpnNatRuleMappingArrayInput
	IpConfigurationId pulumi.StringPtrInput
	Mode              pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	NatRuleName       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Type              pulumi.StringPtrInput
}

func (NatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natRuleArgs)(nil)).Elem()
}

type NatRuleInput interface {
	pulumi.Input

	ToNatRuleOutput() NatRuleOutput
	ToNatRuleOutputWithContext(ctx context.Context) NatRuleOutput
}

func (*NatRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NatRule)(nil)).Elem()
}

func (i *NatRule) ToNatRuleOutput() NatRuleOutput {
	return i.ToNatRuleOutputWithContext(context.Background())
}

func (i *NatRule) ToNatRuleOutputWithContext(ctx context.Context) NatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatRuleOutput)
}

type NatRuleOutput struct{ *pulumi.OutputState }

func (NatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatRule)(nil)).Elem()
}

func (o NatRuleOutput) ToNatRuleOutput() NatRuleOutput {
	return o
}

func (o NatRuleOutput) ToNatRuleOutputWithContext(ctx context.Context) NatRuleOutput {
	return o
}

func (o NatRuleOutput) EgressVpnSiteLinkConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NatRule) SubResourceResponseArrayOutput { return v.EgressVpnSiteLinkConnections }).(SubResourceResponseArrayOutput)
}

func (o NatRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NatRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o NatRuleOutput) ExternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v *NatRule) VpnNatRuleMappingResponseArrayOutput { return v.ExternalMappings }).(VpnNatRuleMappingResponseArrayOutput)
}

func (o NatRuleOutput) IngressVpnSiteLinkConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NatRule) SubResourceResponseArrayOutput { return v.IngressVpnSiteLinkConnections }).(SubResourceResponseArrayOutput)
}

func (o NatRuleOutput) InternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v *NatRule) VpnNatRuleMappingResponseArrayOutput { return v.InternalMappings }).(VpnNatRuleMappingResponseArrayOutput)
}

func (o NatRuleOutput) IpConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatRule) pulumi.StringPtrOutput { return v.IpConfigurationId }).(pulumi.StringPtrOutput)
}

func (o NatRuleOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatRule) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o NatRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatRule) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NatRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NatRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NatRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NatRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NatRuleOutput{})
}

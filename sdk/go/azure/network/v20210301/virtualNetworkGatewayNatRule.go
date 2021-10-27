


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkGatewayNatRule struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput                  `pulumi:"etag"`
	ExternalMappings  VpnNatRuleMappingResponseArrayOutput `pulumi:"externalMappings"`
	InternalMappings  VpnNatRuleMappingResponseArrayOutput `pulumi:"internalMappings"`
	IpConfigurationId pulumi.StringPtrOutput               `pulumi:"ipConfigurationId"`
	Mode              pulumi.StringPtrOutput               `pulumi:"mode"`
	Name              pulumi.StringPtrOutput               `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                  `pulumi:"provisioningState"`
	Type              pulumi.StringOutput                  `pulumi:"type"`
}


func NewVirtualNetworkGatewayNatRule(ctx *pulumi.Context,
	name string, args *VirtualNetworkGatewayNatRuleArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayNatRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualNetworkGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetworkGatewayName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210501:VirtualNetworkGatewayNatRule"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkGatewayNatRule
	err := ctx.RegisterResource("azure-native:network/v20210301:VirtualNetworkGatewayNatRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkGatewayNatRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkGatewayNatRuleState, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayNatRule, error) {
	var resource VirtualNetworkGatewayNatRule
	err := ctx.ReadResource("azure-native:network/v20210301:VirtualNetworkGatewayNatRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualNetworkGatewayNatRuleState struct {
}

type VirtualNetworkGatewayNatRuleState struct {
}

func (VirtualNetworkGatewayNatRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayNatRuleState)(nil)).Elem()
}

type virtualNetworkGatewayNatRuleArgs struct {
	ExternalMappings          []VpnNatRuleMapping `pulumi:"externalMappings"`
	Id                        *string             `pulumi:"id"`
	InternalMappings          []VpnNatRuleMapping `pulumi:"internalMappings"`
	IpConfigurationId         *string             `pulumi:"ipConfigurationId"`
	Mode                      *string             `pulumi:"mode"`
	Name                      *string             `pulumi:"name"`
	NatRuleName               *string             `pulumi:"natRuleName"`
	ResourceGroupName         string              `pulumi:"resourceGroupName"`
	Type                      *string             `pulumi:"type"`
	VirtualNetworkGatewayName string              `pulumi:"virtualNetworkGatewayName"`
}


type VirtualNetworkGatewayNatRuleArgs struct {
	ExternalMappings          VpnNatRuleMappingArrayInput
	Id                        pulumi.StringPtrInput
	InternalMappings          VpnNatRuleMappingArrayInput
	IpConfigurationId         pulumi.StringPtrInput
	Mode                      pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	NatRuleName               pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Type                      pulumi.StringPtrInput
	VirtualNetworkGatewayName pulumi.StringInput
}

func (VirtualNetworkGatewayNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayNatRuleArgs)(nil)).Elem()
}

type VirtualNetworkGatewayNatRuleInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayNatRuleOutput() VirtualNetworkGatewayNatRuleOutput
	ToVirtualNetworkGatewayNatRuleOutputWithContext(ctx context.Context) VirtualNetworkGatewayNatRuleOutput
}

func (*VirtualNetworkGatewayNatRule) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayNatRule)(nil))
}

func (i *VirtualNetworkGatewayNatRule) ToVirtualNetworkGatewayNatRuleOutput() VirtualNetworkGatewayNatRuleOutput {
	return i.ToVirtualNetworkGatewayNatRuleOutputWithContext(context.Background())
}

func (i *VirtualNetworkGatewayNatRule) ToVirtualNetworkGatewayNatRuleOutputWithContext(ctx context.Context) VirtualNetworkGatewayNatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayNatRuleOutput)
}

type VirtualNetworkGatewayNatRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayNatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayNatRule)(nil))
}

func (o VirtualNetworkGatewayNatRuleOutput) ToVirtualNetworkGatewayNatRuleOutput() VirtualNetworkGatewayNatRuleOutput {
	return o
}

func (o VirtualNetworkGatewayNatRuleOutput) ToVirtualNetworkGatewayNatRuleOutputWithContext(ctx context.Context) VirtualNetworkGatewayNatRuleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkGatewayNatRuleInput)(nil)).Elem(), &VirtualNetworkGatewayNatRule{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayNatRuleOutput{})
}

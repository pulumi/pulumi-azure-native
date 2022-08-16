


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureFirewall struct {
	pulumi.CustomResourceState

	AdditionalProperties       pulumi.StringMapOutput                                    `pulumi:"additionalProperties"`
	ApplicationRuleCollections AzureFirewallApplicationRuleCollectionResponseArrayOutput `pulumi:"applicationRuleCollections"`
	Etag                       pulumi.StringOutput                                       `pulumi:"etag"`
	FirewallPolicy             SubResourceResponsePtrOutput                              `pulumi:"firewallPolicy"`
	HubIPAddresses             HubIPAddressesResponsePtrOutput                           `pulumi:"hubIPAddresses"`
	IpConfigurations           AzureFirewallIPConfigurationResponseArrayOutput           `pulumi:"ipConfigurations"`
	IpGroups                   AzureFirewallIpGroupsResponseArrayOutput                  `pulumi:"ipGroups"`
	Location                   pulumi.StringPtrOutput                                    `pulumi:"location"`
	ManagementIpConfiguration  AzureFirewallIPConfigurationResponsePtrOutput             `pulumi:"managementIpConfiguration"`
	Name                       pulumi.StringOutput                                       `pulumi:"name"`
	NatRuleCollections         AzureFirewallNatRuleCollectionResponseArrayOutput         `pulumi:"natRuleCollections"`
	NetworkRuleCollections     AzureFirewallNetworkRuleCollectionResponseArrayOutput     `pulumi:"networkRuleCollections"`
	ProvisioningState          pulumi.StringOutput                                       `pulumi:"provisioningState"`
	Sku                        AzureFirewallSkuResponsePtrOutput                         `pulumi:"sku"`
	Tags                       pulumi.StringMapOutput                                    `pulumi:"tags"`
	ThreatIntelMode            pulumi.StringPtrOutput                                    `pulumi:"threatIntelMode"`
	Type                       pulumi.StringOutput                                       `pulumi:"type"`
	VirtualHub                 SubResourceResponsePtrOutput                              `pulumi:"virtualHub"`
	Zones                      pulumi.StringArrayOutput                                  `pulumi:"zones"`
}


func NewAzureFirewall(ctx *pulumi.Context,
	name string, args *AzureFirewallArgs, opts ...pulumi.ResourceOption) (*AzureFirewall, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:AzureFirewall"),
		},
	})
	opts = append(opts, aliases)
	var resource AzureFirewall
	err := ctx.RegisterResource("azure-native:network/v20210201:AzureFirewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAzureFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureFirewallState, opts ...pulumi.ResourceOption) (*AzureFirewall, error) {
	var resource AzureFirewall
	err := ctx.ReadResource("azure-native:network/v20210201:AzureFirewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type azureFirewallState struct {
}

type AzureFirewallState struct {
}

func (AzureFirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureFirewallState)(nil)).Elem()
}

type azureFirewallArgs struct {
	AdditionalProperties       map[string]string                        `pulumi:"additionalProperties"`
	ApplicationRuleCollections []AzureFirewallApplicationRuleCollection `pulumi:"applicationRuleCollections"`
	AzureFirewallName          *string                                  `pulumi:"azureFirewallName"`
	FirewallPolicy             *SubResource                             `pulumi:"firewallPolicy"`
	HubIPAddresses             *HubIPAddresses                          `pulumi:"hubIPAddresses"`
	Id                         *string                                  `pulumi:"id"`
	IpConfigurations           []AzureFirewallIPConfiguration           `pulumi:"ipConfigurations"`
	Location                   *string                                  `pulumi:"location"`
	ManagementIpConfiguration  *AzureFirewallIPConfiguration            `pulumi:"managementIpConfiguration"`
	NatRuleCollections         []AzureFirewallNatRuleCollection         `pulumi:"natRuleCollections"`
	NetworkRuleCollections     []AzureFirewallNetworkRuleCollection     `pulumi:"networkRuleCollections"`
	ResourceGroupName          string                                   `pulumi:"resourceGroupName"`
	Sku                        *AzureFirewallSku                        `pulumi:"sku"`
	Tags                       map[string]string                        `pulumi:"tags"`
	ThreatIntelMode            *string                                  `pulumi:"threatIntelMode"`
	VirtualHub                 *SubResource                             `pulumi:"virtualHub"`
	Zones                      []string                                 `pulumi:"zones"`
}


type AzureFirewallArgs struct {
	AdditionalProperties       pulumi.StringMapInput
	ApplicationRuleCollections AzureFirewallApplicationRuleCollectionArrayInput
	AzureFirewallName          pulumi.StringPtrInput
	FirewallPolicy             SubResourcePtrInput
	HubIPAddresses             HubIPAddressesPtrInput
	Id                         pulumi.StringPtrInput
	IpConfigurations           AzureFirewallIPConfigurationArrayInput
	Location                   pulumi.StringPtrInput
	ManagementIpConfiguration  AzureFirewallIPConfigurationPtrInput
	NatRuleCollections         AzureFirewallNatRuleCollectionArrayInput
	NetworkRuleCollections     AzureFirewallNetworkRuleCollectionArrayInput
	ResourceGroupName          pulumi.StringInput
	Sku                        AzureFirewallSkuPtrInput
	Tags                       pulumi.StringMapInput
	ThreatIntelMode            pulumi.StringPtrInput
	VirtualHub                 SubResourcePtrInput
	Zones                      pulumi.StringArrayInput
}

func (AzureFirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureFirewallArgs)(nil)).Elem()
}

type AzureFirewallInput interface {
	pulumi.Input

	ToAzureFirewallOutput() AzureFirewallOutput
	ToAzureFirewallOutputWithContext(ctx context.Context) AzureFirewallOutput
}

func (*AzureFirewall) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFirewall)(nil)).Elem()
}

func (i *AzureFirewall) ToAzureFirewallOutput() AzureFirewallOutput {
	return i.ToAzureFirewallOutputWithContext(context.Background())
}

func (i *AzureFirewall) ToAzureFirewallOutputWithContext(ctx context.Context) AzureFirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFirewallOutput)
}

type AzureFirewallOutput struct{ *pulumi.OutputState }

func (AzureFirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFirewall)(nil)).Elem()
}

func (o AzureFirewallOutput) ToAzureFirewallOutput() AzureFirewallOutput {
	return o
}

func (o AzureFirewallOutput) ToAzureFirewallOutputWithContext(ctx context.Context) AzureFirewallOutput {
	return o
}

func (o AzureFirewallOutput) AdditionalProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringMapOutput { return v.AdditionalProperties }).(pulumi.StringMapOutput)
}

func (o AzureFirewallOutput) ApplicationRuleCollections() AzureFirewallApplicationRuleCollectionResponseArrayOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallApplicationRuleCollectionResponseArrayOutput {
		return v.ApplicationRuleCollections
	}).(AzureFirewallApplicationRuleCollectionResponseArrayOutput)
}

func (o AzureFirewallOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o AzureFirewallOutput) FirewallPolicy() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *AzureFirewall) SubResourceResponsePtrOutput { return v.FirewallPolicy }).(SubResourceResponsePtrOutput)
}

func (o AzureFirewallOutput) HubIPAddresses() HubIPAddressesResponsePtrOutput {
	return o.ApplyT(func(v *AzureFirewall) HubIPAddressesResponsePtrOutput { return v.HubIPAddresses }).(HubIPAddressesResponsePtrOutput)
}

func (o AzureFirewallOutput) IpConfigurations() AzureFirewallIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallIPConfigurationResponseArrayOutput { return v.IpConfigurations }).(AzureFirewallIPConfigurationResponseArrayOutput)
}

func (o AzureFirewallOutput) IpGroups() AzureFirewallIpGroupsResponseArrayOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallIpGroupsResponseArrayOutput { return v.IpGroups }).(AzureFirewallIpGroupsResponseArrayOutput)
}

func (o AzureFirewallOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AzureFirewallOutput) ManagementIpConfiguration() AzureFirewallIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallIPConfigurationResponsePtrOutput {
		return v.ManagementIpConfiguration
	}).(AzureFirewallIPConfigurationResponsePtrOutput)
}

func (o AzureFirewallOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AzureFirewallOutput) NatRuleCollections() AzureFirewallNatRuleCollectionResponseArrayOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallNatRuleCollectionResponseArrayOutput { return v.NatRuleCollections }).(AzureFirewallNatRuleCollectionResponseArrayOutput)
}

func (o AzureFirewallOutput) NetworkRuleCollections() AzureFirewallNetworkRuleCollectionResponseArrayOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallNetworkRuleCollectionResponseArrayOutput {
		return v.NetworkRuleCollections
	}).(AzureFirewallNetworkRuleCollectionResponseArrayOutput)
}

func (o AzureFirewallOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AzureFirewallOutput) Sku() AzureFirewallSkuResponsePtrOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallSkuResponsePtrOutput { return v.Sku }).(AzureFirewallSkuResponsePtrOutput)
}

func (o AzureFirewallOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AzureFirewallOutput) ThreatIntelMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringPtrOutput { return v.ThreatIntelMode }).(pulumi.StringPtrOutput)
}

func (o AzureFirewallOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AzureFirewallOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *AzureFirewall) SubResourceResponsePtrOutput { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func (o AzureFirewallOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureFirewallOutput{})
}

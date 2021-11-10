


package v20200301

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
	HubIpAddresses             HubIPAddressesResponseOutput                              `pulumi:"hubIpAddresses"`
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
			Type: pulumi.String("azure-native:network/v20210201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:AzureFirewall"),
		},
	})
	opts = append(opts, aliases)
	var resource AzureFirewall
	err := ctx.RegisterResource("azure-native:network/v20200301:AzureFirewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAzureFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureFirewallState, opts ...pulumi.ResourceOption) (*AzureFirewall, error) {
	var resource AzureFirewall
	err := ctx.ReadResource("azure-native:network/v20200301:AzureFirewall", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*AzureFirewall)(nil))
}

func (i *AzureFirewall) ToAzureFirewallOutput() AzureFirewallOutput {
	return i.ToAzureFirewallOutputWithContext(context.Background())
}

func (i *AzureFirewall) ToAzureFirewallOutputWithContext(ctx context.Context) AzureFirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFirewallOutput)
}

type AzureFirewallOutput struct{ *pulumi.OutputState }

func (AzureFirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewall)(nil))
}

func (o AzureFirewallOutput) ToAzureFirewallOutput() AzureFirewallOutput {
	return o
}

func (o AzureFirewallOutput) ToAzureFirewallOutputWithContext(ctx context.Context) AzureFirewallOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AzureFirewallOutput{})
}

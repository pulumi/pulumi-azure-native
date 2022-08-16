


package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAzureFirewall(ctx *pulumi.Context, args *LookupAzureFirewallArgs, opts ...pulumi.InvokeOption) (*LookupAzureFirewallResult, error) {
	var rv LookupAzureFirewallResult
	err := ctx.Invoke("azure-native:network/v20200601:getAzureFirewall", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAzureFirewallArgs struct {
	AzureFirewallName string `pulumi:"azureFirewallName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAzureFirewallResult struct {
	AdditionalProperties       map[string]string                                `pulumi:"additionalProperties"`
	ApplicationRuleCollections []AzureFirewallApplicationRuleCollectionResponse `pulumi:"applicationRuleCollections"`
	Etag                       string                                           `pulumi:"etag"`
	FirewallPolicy             *SubResourceResponse                             `pulumi:"firewallPolicy"`
	HubIPAddresses             *HubIPAddressesResponse                          `pulumi:"hubIPAddresses"`
	Id                         *string                                          `pulumi:"id"`
	IpConfigurations           []AzureFirewallIPConfigurationResponse           `pulumi:"ipConfigurations"`
	IpGroups                   []AzureFirewallIpGroupsResponse                  `pulumi:"ipGroups"`
	Location                   *string                                          `pulumi:"location"`
	ManagementIpConfiguration  *AzureFirewallIPConfigurationResponse            `pulumi:"managementIpConfiguration"`
	Name                       string                                           `pulumi:"name"`
	NatRuleCollections         []AzureFirewallNatRuleCollectionResponse         `pulumi:"natRuleCollections"`
	NetworkRuleCollections     []AzureFirewallNetworkRuleCollectionResponse     `pulumi:"networkRuleCollections"`
	ProvisioningState          string                                           `pulumi:"provisioningState"`
	Sku                        *AzureFirewallSkuResponse                        `pulumi:"sku"`
	Tags                       map[string]string                                `pulumi:"tags"`
	ThreatIntelMode            *string                                          `pulumi:"threatIntelMode"`
	Type                       string                                           `pulumi:"type"`
	VirtualHub                 *SubResourceResponse                             `pulumi:"virtualHub"`
	Zones                      []string                                         `pulumi:"zones"`
}

func LookupAzureFirewallOutput(ctx *pulumi.Context, args LookupAzureFirewallOutputArgs, opts ...pulumi.InvokeOption) LookupAzureFirewallResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzureFirewallResult, error) {
			args := v.(LookupAzureFirewallArgs)
			r, err := LookupAzureFirewall(ctx, &args, opts...)
			var s LookupAzureFirewallResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzureFirewallResultOutput)
}

type LookupAzureFirewallOutputArgs struct {
	AzureFirewallName pulumi.StringInput `pulumi:"azureFirewallName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAzureFirewallOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureFirewallArgs)(nil)).Elem()
}


type LookupAzureFirewallResultOutput struct{ *pulumi.OutputState }

func (LookupAzureFirewallResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureFirewallResult)(nil)).Elem()
}

func (o LookupAzureFirewallResultOutput) ToLookupAzureFirewallResultOutput() LookupAzureFirewallResultOutput {
	return o
}

func (o LookupAzureFirewallResultOutput) ToLookupAzureFirewallResultOutputWithContext(ctx context.Context) LookupAzureFirewallResultOutput {
	return o
}

func (o LookupAzureFirewallResultOutput) AdditionalProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) map[string]string { return v.AdditionalProperties }).(pulumi.StringMapOutput)
}

func (o LookupAzureFirewallResultOutput) ApplicationRuleCollections() AzureFirewallApplicationRuleCollectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) []AzureFirewallApplicationRuleCollectionResponse {
		return v.ApplicationRuleCollections
	}).(AzureFirewallApplicationRuleCollectionResponseArrayOutput)
}

func (o LookupAzureFirewallResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupAzureFirewallResultOutput) FirewallPolicy() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *SubResourceResponse { return v.FirewallPolicy }).(SubResourceResponsePtrOutput)
}

func (o LookupAzureFirewallResultOutput) HubIPAddresses() HubIPAddressesResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *HubIPAddressesResponse { return v.HubIPAddresses }).(HubIPAddressesResponsePtrOutput)
}

func (o LookupAzureFirewallResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAzureFirewallResultOutput) IpConfigurations() AzureFirewallIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) []AzureFirewallIPConfigurationResponse { return v.IpConfigurations }).(AzureFirewallIPConfigurationResponseArrayOutput)
}

func (o LookupAzureFirewallResultOutput) IpGroups() AzureFirewallIpGroupsResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) []AzureFirewallIpGroupsResponse { return v.IpGroups }).(AzureFirewallIpGroupsResponseArrayOutput)
}

func (o LookupAzureFirewallResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAzureFirewallResultOutput) ManagementIpConfiguration() AzureFirewallIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *AzureFirewallIPConfigurationResponse {
		return v.ManagementIpConfiguration
	}).(AzureFirewallIPConfigurationResponsePtrOutput)
}

func (o LookupAzureFirewallResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAzureFirewallResultOutput) NatRuleCollections() AzureFirewallNatRuleCollectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) []AzureFirewallNatRuleCollectionResponse {
		return v.NatRuleCollections
	}).(AzureFirewallNatRuleCollectionResponseArrayOutput)
}

func (o LookupAzureFirewallResultOutput) NetworkRuleCollections() AzureFirewallNetworkRuleCollectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) []AzureFirewallNetworkRuleCollectionResponse {
		return v.NetworkRuleCollections
	}).(AzureFirewallNetworkRuleCollectionResponseArrayOutput)
}

func (o LookupAzureFirewallResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAzureFirewallResultOutput) Sku() AzureFirewallSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *AzureFirewallSkuResponse { return v.Sku }).(AzureFirewallSkuResponsePtrOutput)
}

func (o LookupAzureFirewallResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAzureFirewallResultOutput) ThreatIntelMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *string { return v.ThreatIntelMode }).(pulumi.StringPtrOutput)
}

func (o LookupAzureFirewallResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAzureFirewallResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func (o LookupAzureFirewallResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureFirewallResultOutput{})
}




package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActiveConnectivityConfigurationResponse struct {
	AppliesToGroups       []ConnectivityGroupItemResponse `pulumi:"appliesToGroups"`
	CommitTime            *string                         `pulumi:"commitTime"`
	ConfigurationGroups   []ConfigurationGroupResponse    `pulumi:"configurationGroups"`
	ConnectivityTopology  string                          `pulumi:"connectivityTopology"`
	DeleteExistingPeering *string                         `pulumi:"deleteExistingPeering"`
	Description           *string                         `pulumi:"description"`
	Hubs                  []HubResponse                   `pulumi:"hubs"`
	Id                    *string                         `pulumi:"id"`
	IsGlobal              *string                         `pulumi:"isGlobal"`
	ProvisioningState     string                          `pulumi:"provisioningState"`
	Region                *string                         `pulumi:"region"`
}

type ActiveConnectivityConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ActiveConnectivityConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveConnectivityConfigurationResponse)(nil)).Elem()
}

func (o ActiveConnectivityConfigurationResponseOutput) ToActiveConnectivityConfigurationResponseOutput() ActiveConnectivityConfigurationResponseOutput {
	return o
}

func (o ActiveConnectivityConfigurationResponseOutput) ToActiveConnectivityConfigurationResponseOutputWithContext(ctx context.Context) ActiveConnectivityConfigurationResponseOutput {
	return o
}

func (o ActiveConnectivityConfigurationResponseOutput) AppliesToGroups() ConnectivityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveConnectivityConfigurationResponse) []ConnectivityGroupItemResponse {
		return v.AppliesToGroups
	}).(ConnectivityGroupItemResponseArrayOutput)
}

func (o ActiveConnectivityConfigurationResponseOutput) CommitTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveConnectivityConfigurationResponse) *string { return v.CommitTime }).(pulumi.StringPtrOutput)
}

func (o ActiveConnectivityConfigurationResponseOutput) ConfigurationGroups() ConfigurationGroupResponseArrayOutput {
	return o.ApplyT(func(v ActiveConnectivityConfigurationResponse) []ConfigurationGroupResponse {
		return v.ConfigurationGroups
	}).(ConfigurationGroupResponseArrayOutput)
}

func (o ActiveConnectivityConfigurationResponseOutput) ConnectivityTopology() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveConnectivityConfigurationResponse) string { return v.ConnectivityTopology }).(pulumi.StringOutput)
}

func (o ActiveConnectivityConfigurationResponseOutput) DeleteExistingPeering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveConnectivityConfigurationResponse) *string { return v.DeleteExistingPeering }).(pulumi.StringPtrOutput)
}

func (o ActiveConnectivityConfigurationResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveConnectivityConfigurationResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ActiveConnectivityConfigurationResponseOutput) Hubs() HubResponseArrayOutput {
	return o.ApplyT(func(v ActiveConnectivityConfigurationResponse) []HubResponse { return v.Hubs }).(HubResponseArrayOutput)
}

func (o ActiveConnectivityConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveConnectivityConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ActiveConnectivityConfigurationResponseOutput) IsGlobal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveConnectivityConfigurationResponse) *string { return v.IsGlobal }).(pulumi.StringPtrOutput)
}

func (o ActiveConnectivityConfigurationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveConnectivityConfigurationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ActiveConnectivityConfigurationResponseOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveConnectivityConfigurationResponse) *string { return v.Region }).(pulumi.StringPtrOutput)
}

type ActiveConnectivityConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (ActiveConnectivityConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveConnectivityConfigurationResponse)(nil)).Elem()
}

func (o ActiveConnectivityConfigurationResponseArrayOutput) ToActiveConnectivityConfigurationResponseArrayOutput() ActiveConnectivityConfigurationResponseArrayOutput {
	return o
}

func (o ActiveConnectivityConfigurationResponseArrayOutput) ToActiveConnectivityConfigurationResponseArrayOutputWithContext(ctx context.Context) ActiveConnectivityConfigurationResponseArrayOutput {
	return o
}

func (o ActiveConnectivityConfigurationResponseArrayOutput) Index(i pulumi.IntInput) ActiveConnectivityConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActiveConnectivityConfigurationResponse {
		return vs[0].([]ActiveConnectivityConfigurationResponse)[vs[1].(int)]
	}).(ActiveConnectivityConfigurationResponseOutput)
}

type ActiveDefaultSecurityAdminRuleResponse struct {
	Access                        string                                    `pulumi:"access"`
	CommitTime                    *string                                   `pulumi:"commitTime"`
	ConfigurationDescription      *string                                   `pulumi:"configurationDescription"`
	Description                   string                                    `pulumi:"description"`
	DestinationPortRanges         []string                                  `pulumi:"destinationPortRanges"`
	Destinations                  []AddressPrefixItemResponse               `pulumi:"destinations"`
	Direction                     string                                    `pulumi:"direction"`
	Flag                          *string                                   `pulumi:"flag"`
	Id                            *string                                   `pulumi:"id"`
	Kind                          string                                    `pulumi:"kind"`
	Priority                      int                                       `pulumi:"priority"`
	Protocol                      string                                    `pulumi:"protocol"`
	ProvisioningState             string                                    `pulumi:"provisioningState"`
	Region                        *string                                   `pulumi:"region"`
	RuleCollectionAppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     *string                                   `pulumi:"ruleCollectionDescription"`
	RuleGroups                    []ConfigurationGroupResponse              `pulumi:"ruleGroups"`
	SourcePortRanges              []string                                  `pulumi:"sourcePortRanges"`
	Sources                       []AddressPrefixItemResponse               `pulumi:"sources"`
}

type ActiveDefaultSecurityUserRuleResponse struct {
	CommitTime                    *string                                   `pulumi:"commitTime"`
	ConfigurationDescription      *string                                   `pulumi:"configurationDescription"`
	Description                   string                                    `pulumi:"description"`
	DestinationPortRanges         []string                                  `pulumi:"destinationPortRanges"`
	Destinations                  []AddressPrefixItemResponse               `pulumi:"destinations"`
	Direction                     string                                    `pulumi:"direction"`
	Flag                          *string                                   `pulumi:"flag"`
	Id                            *string                                   `pulumi:"id"`
	Kind                          string                                    `pulumi:"kind"`
	Protocol                      string                                    `pulumi:"protocol"`
	ProvisioningState             string                                    `pulumi:"provisioningState"`
	Region                        *string                                   `pulumi:"region"`
	RuleCollectionAppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     *string                                   `pulumi:"ruleCollectionDescription"`
	RuleGroups                    []ConfigurationGroupResponse              `pulumi:"ruleGroups"`
	SourcePortRanges              []string                                  `pulumi:"sourcePortRanges"`
	Sources                       []AddressPrefixItemResponse               `pulumi:"sources"`
}

type ActiveSecurityAdminRuleResponse struct {
	Access                        string                                    `pulumi:"access"`
	CommitTime                    *string                                   `pulumi:"commitTime"`
	ConfigurationDescription      *string                                   `pulumi:"configurationDescription"`
	Description                   *string                                   `pulumi:"description"`
	DestinationPortRanges         []string                                  `pulumi:"destinationPortRanges"`
	Destinations                  []AddressPrefixItemResponse               `pulumi:"destinations"`
	Direction                     string                                    `pulumi:"direction"`
	Id                            *string                                   `pulumi:"id"`
	Kind                          string                                    `pulumi:"kind"`
	Priority                      int                                       `pulumi:"priority"`
	Protocol                      string                                    `pulumi:"protocol"`
	ProvisioningState             string                                    `pulumi:"provisioningState"`
	Region                        *string                                   `pulumi:"region"`
	RuleCollectionAppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     *string                                   `pulumi:"ruleCollectionDescription"`
	RuleGroups                    []ConfigurationGroupResponse              `pulumi:"ruleGroups"`
	SourcePortRanges              []string                                  `pulumi:"sourcePortRanges"`
	Sources                       []AddressPrefixItemResponse               `pulumi:"sources"`
}

type ActiveSecurityUserRuleResponse struct {
	CommitTime                    *string                                   `pulumi:"commitTime"`
	ConfigurationDescription      *string                                   `pulumi:"configurationDescription"`
	Description                   *string                                   `pulumi:"description"`
	DestinationPortRanges         []string                                  `pulumi:"destinationPortRanges"`
	Destinations                  []AddressPrefixItemResponse               `pulumi:"destinations"`
	Direction                     string                                    `pulumi:"direction"`
	Id                            *string                                   `pulumi:"id"`
	Kind                          string                                    `pulumi:"kind"`
	Protocol                      string                                    `pulumi:"protocol"`
	ProvisioningState             string                                    `pulumi:"provisioningState"`
	Region                        *string                                   `pulumi:"region"`
	RuleCollectionAppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     *string                                   `pulumi:"ruleCollectionDescription"`
	RuleGroups                    []ConfigurationGroupResponse              `pulumi:"ruleGroups"`
	SourcePortRanges              []string                                  `pulumi:"sourcePortRanges"`
	Sources                       []AddressPrefixItemResponse               `pulumi:"sources"`
}

type AddressPrefixItem struct {
	AddressPrefix     *string `pulumi:"addressPrefix"`
	AddressPrefixType *string `pulumi:"addressPrefixType"`
}

type AddressPrefixItemResponse struct {
	AddressPrefix     *string `pulumi:"addressPrefix"`
	AddressPrefixType *string `pulumi:"addressPrefixType"`
}

type AddressPrefixItemResponseOutput struct{ *pulumi.OutputState }

func (AddressPrefixItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressPrefixItemResponse)(nil)).Elem()
}

func (o AddressPrefixItemResponseOutput) ToAddressPrefixItemResponseOutput() AddressPrefixItemResponseOutput {
	return o
}

func (o AddressPrefixItemResponseOutput) ToAddressPrefixItemResponseOutputWithContext(ctx context.Context) AddressPrefixItemResponseOutput {
	return o
}

func (o AddressPrefixItemResponseOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressPrefixItemResponse) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o AddressPrefixItemResponseOutput) AddressPrefixType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressPrefixItemResponse) *string { return v.AddressPrefixType }).(pulumi.StringPtrOutput)
}

type AddressPrefixItemResponseArrayOutput struct{ *pulumi.OutputState }

func (AddressPrefixItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AddressPrefixItemResponse)(nil)).Elem()
}

func (o AddressPrefixItemResponseArrayOutput) ToAddressPrefixItemResponseArrayOutput() AddressPrefixItemResponseArrayOutput {
	return o
}

func (o AddressPrefixItemResponseArrayOutput) ToAddressPrefixItemResponseArrayOutputWithContext(ctx context.Context) AddressPrefixItemResponseArrayOutput {
	return o
}

func (o AddressPrefixItemResponseArrayOutput) Index(i pulumi.IntInput) AddressPrefixItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AddressPrefixItemResponse {
		return vs[0].([]AddressPrefixItemResponse)[vs[1].(int)]
	}).(AddressPrefixItemResponseOutput)
}

type ConfigurationGroupResponse struct {
	Description       *string `pulumi:"description"`
	Id                *string `pulumi:"id"`
	MemberType        string  `pulumi:"memberType"`
	ProvisioningState string  `pulumi:"provisioningState"`
}

type ConfigurationGroupResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationGroupResponse)(nil)).Elem()
}

func (o ConfigurationGroupResponseOutput) ToConfigurationGroupResponseOutput() ConfigurationGroupResponseOutput {
	return o
}

func (o ConfigurationGroupResponseOutput) ToConfigurationGroupResponseOutputWithContext(ctx context.Context) ConfigurationGroupResponseOutput {
	return o
}

func (o ConfigurationGroupResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationGroupResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConfigurationGroupResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationGroupResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ConfigurationGroupResponseOutput) MemberType() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationGroupResponse) string { return v.MemberType }).(pulumi.StringOutput)
}

func (o ConfigurationGroupResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationGroupResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ConfigurationGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationGroupResponse)(nil)).Elem()
}

func (o ConfigurationGroupResponseArrayOutput) ToConfigurationGroupResponseArrayOutput() ConfigurationGroupResponseArrayOutput {
	return o
}

func (o ConfigurationGroupResponseArrayOutput) ToConfigurationGroupResponseArrayOutputWithContext(ctx context.Context) ConfigurationGroupResponseArrayOutput {
	return o
}

func (o ConfigurationGroupResponseArrayOutput) Index(i pulumi.IntInput) ConfigurationGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationGroupResponse {
		return vs[0].([]ConfigurationGroupResponse)[vs[1].(int)]
	}).(ConfigurationGroupResponseOutput)
}

type ConnectivityGroupItem struct {
	GroupConnectivity string  `pulumi:"groupConnectivity"`
	IsGlobal          *string `pulumi:"isGlobal"`
	NetworkGroupId    string  `pulumi:"networkGroupId"`
	UseHubGateway     *string `pulumi:"useHubGateway"`
}





type ConnectivityGroupItemInput interface {
	pulumi.Input

	ToConnectivityGroupItemOutput() ConnectivityGroupItemOutput
	ToConnectivityGroupItemOutputWithContext(context.Context) ConnectivityGroupItemOutput
}

type ConnectivityGroupItemArgs struct {
	GroupConnectivity pulumi.StringInput    `pulumi:"groupConnectivity"`
	IsGlobal          pulumi.StringPtrInput `pulumi:"isGlobal"`
	NetworkGroupId    pulumi.StringInput    `pulumi:"networkGroupId"`
	UseHubGateway     pulumi.StringPtrInput `pulumi:"useHubGateway"`
}

func (ConnectivityGroupItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityGroupItem)(nil)).Elem()
}

func (i ConnectivityGroupItemArgs) ToConnectivityGroupItemOutput() ConnectivityGroupItemOutput {
	return i.ToConnectivityGroupItemOutputWithContext(context.Background())
}

func (i ConnectivityGroupItemArgs) ToConnectivityGroupItemOutputWithContext(ctx context.Context) ConnectivityGroupItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectivityGroupItemOutput)
}





type ConnectivityGroupItemArrayInput interface {
	pulumi.Input

	ToConnectivityGroupItemArrayOutput() ConnectivityGroupItemArrayOutput
	ToConnectivityGroupItemArrayOutputWithContext(context.Context) ConnectivityGroupItemArrayOutput
}

type ConnectivityGroupItemArray []ConnectivityGroupItemInput

func (ConnectivityGroupItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectivityGroupItem)(nil)).Elem()
}

func (i ConnectivityGroupItemArray) ToConnectivityGroupItemArrayOutput() ConnectivityGroupItemArrayOutput {
	return i.ToConnectivityGroupItemArrayOutputWithContext(context.Background())
}

func (i ConnectivityGroupItemArray) ToConnectivityGroupItemArrayOutputWithContext(ctx context.Context) ConnectivityGroupItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectivityGroupItemArrayOutput)
}

type ConnectivityGroupItemOutput struct{ *pulumi.OutputState }

func (ConnectivityGroupItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityGroupItem)(nil)).Elem()
}

func (o ConnectivityGroupItemOutput) ToConnectivityGroupItemOutput() ConnectivityGroupItemOutput {
	return o
}

func (o ConnectivityGroupItemOutput) ToConnectivityGroupItemOutputWithContext(ctx context.Context) ConnectivityGroupItemOutput {
	return o
}

func (o ConnectivityGroupItemOutput) GroupConnectivity() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectivityGroupItem) string { return v.GroupConnectivity }).(pulumi.StringOutput)
}

func (o ConnectivityGroupItemOutput) IsGlobal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityGroupItem) *string { return v.IsGlobal }).(pulumi.StringPtrOutput)
}

func (o ConnectivityGroupItemOutput) NetworkGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectivityGroupItem) string { return v.NetworkGroupId }).(pulumi.StringOutput)
}

func (o ConnectivityGroupItemOutput) UseHubGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityGroupItem) *string { return v.UseHubGateway }).(pulumi.StringPtrOutput)
}

type ConnectivityGroupItemArrayOutput struct{ *pulumi.OutputState }

func (ConnectivityGroupItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectivityGroupItem)(nil)).Elem()
}

func (o ConnectivityGroupItemArrayOutput) ToConnectivityGroupItemArrayOutput() ConnectivityGroupItemArrayOutput {
	return o
}

func (o ConnectivityGroupItemArrayOutput) ToConnectivityGroupItemArrayOutputWithContext(ctx context.Context) ConnectivityGroupItemArrayOutput {
	return o
}

func (o ConnectivityGroupItemArrayOutput) Index(i pulumi.IntInput) ConnectivityGroupItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectivityGroupItem {
		return vs[0].([]ConnectivityGroupItem)[vs[1].(int)]
	}).(ConnectivityGroupItemOutput)
}

type ConnectivityGroupItemResponse struct {
	GroupConnectivity string  `pulumi:"groupConnectivity"`
	IsGlobal          *string `pulumi:"isGlobal"`
	NetworkGroupId    string  `pulumi:"networkGroupId"`
	UseHubGateway     *string `pulumi:"useHubGateway"`
}

type ConnectivityGroupItemResponseOutput struct{ *pulumi.OutputState }

func (ConnectivityGroupItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityGroupItemResponse)(nil)).Elem()
}

func (o ConnectivityGroupItemResponseOutput) ToConnectivityGroupItemResponseOutput() ConnectivityGroupItemResponseOutput {
	return o
}

func (o ConnectivityGroupItemResponseOutput) ToConnectivityGroupItemResponseOutputWithContext(ctx context.Context) ConnectivityGroupItemResponseOutput {
	return o
}

func (o ConnectivityGroupItemResponseOutput) GroupConnectivity() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectivityGroupItemResponse) string { return v.GroupConnectivity }).(pulumi.StringOutput)
}

func (o ConnectivityGroupItemResponseOutput) IsGlobal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityGroupItemResponse) *string { return v.IsGlobal }).(pulumi.StringPtrOutput)
}

func (o ConnectivityGroupItemResponseOutput) NetworkGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectivityGroupItemResponse) string { return v.NetworkGroupId }).(pulumi.StringOutput)
}

func (o ConnectivityGroupItemResponseOutput) UseHubGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityGroupItemResponse) *string { return v.UseHubGateway }).(pulumi.StringPtrOutput)
}

type ConnectivityGroupItemResponseArrayOutput struct{ *pulumi.OutputState }

func (ConnectivityGroupItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectivityGroupItemResponse)(nil)).Elem()
}

func (o ConnectivityGroupItemResponseArrayOutput) ToConnectivityGroupItemResponseArrayOutput() ConnectivityGroupItemResponseArrayOutput {
	return o
}

func (o ConnectivityGroupItemResponseArrayOutput) ToConnectivityGroupItemResponseArrayOutputWithContext(ctx context.Context) ConnectivityGroupItemResponseArrayOutput {
	return o
}

func (o ConnectivityGroupItemResponseArrayOutput) Index(i pulumi.IntInput) ConnectivityGroupItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectivityGroupItemResponse {
		return vs[0].([]ConnectivityGroupItemResponse)[vs[1].(int)]
	}).(ConnectivityGroupItemResponseOutput)
}

type CrossTenantScopesResponse struct {
	ManagementGroups []string `pulumi:"managementGroups"`
	Subscriptions    []string `pulumi:"subscriptions"`
	TenantId         string   `pulumi:"tenantId"`
}

type CrossTenantScopesResponseOutput struct{ *pulumi.OutputState }

func (CrossTenantScopesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossTenantScopesResponse)(nil)).Elem()
}

func (o CrossTenantScopesResponseOutput) ToCrossTenantScopesResponseOutput() CrossTenantScopesResponseOutput {
	return o
}

func (o CrossTenantScopesResponseOutput) ToCrossTenantScopesResponseOutputWithContext(ctx context.Context) CrossTenantScopesResponseOutput {
	return o
}

func (o CrossTenantScopesResponseOutput) ManagementGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CrossTenantScopesResponse) []string { return v.ManagementGroups }).(pulumi.StringArrayOutput)
}

func (o CrossTenantScopesResponseOutput) Subscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CrossTenantScopesResponse) []string { return v.Subscriptions }).(pulumi.StringArrayOutput)
}

func (o CrossTenantScopesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v CrossTenantScopesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type CrossTenantScopesResponseArrayOutput struct{ *pulumi.OutputState }

func (CrossTenantScopesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CrossTenantScopesResponse)(nil)).Elem()
}

func (o CrossTenantScopesResponseArrayOutput) ToCrossTenantScopesResponseArrayOutput() CrossTenantScopesResponseArrayOutput {
	return o
}

func (o CrossTenantScopesResponseArrayOutput) ToCrossTenantScopesResponseArrayOutputWithContext(ctx context.Context) CrossTenantScopesResponseArrayOutput {
	return o
}

func (o CrossTenantScopesResponseArrayOutput) Index(i pulumi.IntInput) CrossTenantScopesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CrossTenantScopesResponse {
		return vs[0].([]CrossTenantScopesResponse)[vs[1].(int)]
	}).(CrossTenantScopesResponseOutput)
}

type DnsConfig struct {
	RelativeName *string  `pulumi:"relativeName"`
	Ttl          *float64 `pulumi:"ttl"`
}





type DnsConfigInput interface {
	pulumi.Input

	ToDnsConfigOutput() DnsConfigOutput
	ToDnsConfigOutputWithContext(context.Context) DnsConfigOutput
}

type DnsConfigArgs struct {
	RelativeName pulumi.StringPtrInput  `pulumi:"relativeName"`
	Ttl          pulumi.Float64PtrInput `pulumi:"ttl"`
}

func (DnsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfig)(nil)).Elem()
}

func (i DnsConfigArgs) ToDnsConfigOutput() DnsConfigOutput {
	return i.ToDnsConfigOutputWithContext(context.Background())
}

func (i DnsConfigArgs) ToDnsConfigOutputWithContext(ctx context.Context) DnsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigOutput)
}

func (i DnsConfigArgs) ToDnsConfigPtrOutput() DnsConfigPtrOutput {
	return i.ToDnsConfigPtrOutputWithContext(context.Background())
}

func (i DnsConfigArgs) ToDnsConfigPtrOutputWithContext(ctx context.Context) DnsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigOutput).ToDnsConfigPtrOutputWithContext(ctx)
}









type DnsConfigPtrInput interface {
	pulumi.Input

	ToDnsConfigPtrOutput() DnsConfigPtrOutput
	ToDnsConfigPtrOutputWithContext(context.Context) DnsConfigPtrOutput
}

type dnsConfigPtrType DnsConfigArgs

func DnsConfigPtr(v *DnsConfigArgs) DnsConfigPtrInput {
	return (*dnsConfigPtrType)(v)
}

func (*dnsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfig)(nil)).Elem()
}

func (i *dnsConfigPtrType) ToDnsConfigPtrOutput() DnsConfigPtrOutput {
	return i.ToDnsConfigPtrOutputWithContext(context.Background())
}

func (i *dnsConfigPtrType) ToDnsConfigPtrOutputWithContext(ctx context.Context) DnsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigPtrOutput)
}

type DnsConfigOutput struct{ *pulumi.OutputState }

func (DnsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfig)(nil)).Elem()
}

func (o DnsConfigOutput) ToDnsConfigOutput() DnsConfigOutput {
	return o
}

func (o DnsConfigOutput) ToDnsConfigOutputWithContext(ctx context.Context) DnsConfigOutput {
	return o
}

func (o DnsConfigOutput) ToDnsConfigPtrOutput() DnsConfigPtrOutput {
	return o.ToDnsConfigPtrOutputWithContext(context.Background())
}

func (o DnsConfigOutput) ToDnsConfigPtrOutputWithContext(ctx context.Context) DnsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DnsConfig) *DnsConfig {
		return &v
	}).(DnsConfigPtrOutput)
}

func (o DnsConfigOutput) RelativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfig) *string { return v.RelativeName }).(pulumi.StringPtrOutput)
}

func (o DnsConfigOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DnsConfig) *float64 { return v.Ttl }).(pulumi.Float64PtrOutput)
}

type DnsConfigPtrOutput struct{ *pulumi.OutputState }

func (DnsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfig)(nil)).Elem()
}

func (o DnsConfigPtrOutput) ToDnsConfigPtrOutput() DnsConfigPtrOutput {
	return o
}

func (o DnsConfigPtrOutput) ToDnsConfigPtrOutputWithContext(ctx context.Context) DnsConfigPtrOutput {
	return o
}

func (o DnsConfigPtrOutput) Elem() DnsConfigOutput {
	return o.ApplyT(func(v *DnsConfig) DnsConfig {
		if v != nil {
			return *v
		}
		var ret DnsConfig
		return ret
	}).(DnsConfigOutput)
}

func (o DnsConfigPtrOutput) RelativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfig) *string {
		if v == nil {
			return nil
		}
		return v.RelativeName
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigPtrOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DnsConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.Ttl
	}).(pulumi.Float64PtrOutput)
}

type DnsConfigResponse struct {
	Fqdn         string   `pulumi:"fqdn"`
	RelativeName *string  `pulumi:"relativeName"`
	Ttl          *float64 `pulumi:"ttl"`
}

type DnsConfigResponseOutput struct{ *pulumi.OutputState }

func (DnsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfigResponse)(nil)).Elem()
}

func (o DnsConfigResponseOutput) ToDnsConfigResponseOutput() DnsConfigResponseOutput {
	return o
}

func (o DnsConfigResponseOutput) ToDnsConfigResponseOutputWithContext(ctx context.Context) DnsConfigResponseOutput {
	return o
}

func (o DnsConfigResponseOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v DnsConfigResponse) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o DnsConfigResponseOutput) RelativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfigResponse) *string { return v.RelativeName }).(pulumi.StringPtrOutput)
}

func (o DnsConfigResponseOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DnsConfigResponse) *float64 { return v.Ttl }).(pulumi.Float64PtrOutput)
}

type DnsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (DnsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfigResponse)(nil)).Elem()
}

func (o DnsConfigResponsePtrOutput) ToDnsConfigResponsePtrOutput() DnsConfigResponsePtrOutput {
	return o
}

func (o DnsConfigResponsePtrOutput) ToDnsConfigResponsePtrOutputWithContext(ctx context.Context) DnsConfigResponsePtrOutput {
	return o
}

func (o DnsConfigResponsePtrOutput) Elem() DnsConfigResponseOutput {
	return o.ApplyT(func(v *DnsConfigResponse) DnsConfigResponse {
		if v != nil {
			return *v
		}
		var ret DnsConfigResponse
		return ret
	}).(DnsConfigResponseOutput)
}

func (o DnsConfigResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigResponsePtrOutput) RelativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.RelativeName
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigResponsePtrOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DnsConfigResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Ttl
	}).(pulumi.Float64PtrOutput)
}

type EffectiveConnectivityConfigurationResponse struct {
	AppliesToGroups       []ConnectivityGroupItemResponse `pulumi:"appliesToGroups"`
	ConfigurationGroups   []ConfigurationGroupResponse    `pulumi:"configurationGroups"`
	ConnectivityTopology  string                          `pulumi:"connectivityTopology"`
	DeleteExistingPeering *string                         `pulumi:"deleteExistingPeering"`
	Description           *string                         `pulumi:"description"`
	Hubs                  []HubResponse                   `pulumi:"hubs"`
	Id                    *string                         `pulumi:"id"`
	IsGlobal              *string                         `pulumi:"isGlobal"`
	ProvisioningState     string                          `pulumi:"provisioningState"`
}

type EffectiveConnectivityConfigurationResponseOutput struct{ *pulumi.OutputState }

func (EffectiveConnectivityConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EffectiveConnectivityConfigurationResponse)(nil)).Elem()
}

func (o EffectiveConnectivityConfigurationResponseOutput) ToEffectiveConnectivityConfigurationResponseOutput() EffectiveConnectivityConfigurationResponseOutput {
	return o
}

func (o EffectiveConnectivityConfigurationResponseOutput) ToEffectiveConnectivityConfigurationResponseOutputWithContext(ctx context.Context) EffectiveConnectivityConfigurationResponseOutput {
	return o
}

func (o EffectiveConnectivityConfigurationResponseOutput) AppliesToGroups() ConnectivityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v EffectiveConnectivityConfigurationResponse) []ConnectivityGroupItemResponse {
		return v.AppliesToGroups
	}).(ConnectivityGroupItemResponseArrayOutput)
}

func (o EffectiveConnectivityConfigurationResponseOutput) ConfigurationGroups() ConfigurationGroupResponseArrayOutput {
	return o.ApplyT(func(v EffectiveConnectivityConfigurationResponse) []ConfigurationGroupResponse {
		return v.ConfigurationGroups
	}).(ConfigurationGroupResponseArrayOutput)
}

func (o EffectiveConnectivityConfigurationResponseOutput) ConnectivityTopology() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveConnectivityConfigurationResponse) string { return v.ConnectivityTopology }).(pulumi.StringOutput)
}

func (o EffectiveConnectivityConfigurationResponseOutput) DeleteExistingPeering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveConnectivityConfigurationResponse) *string { return v.DeleteExistingPeering }).(pulumi.StringPtrOutput)
}

func (o EffectiveConnectivityConfigurationResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveConnectivityConfigurationResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EffectiveConnectivityConfigurationResponseOutput) Hubs() HubResponseArrayOutput {
	return o.ApplyT(func(v EffectiveConnectivityConfigurationResponse) []HubResponse { return v.Hubs }).(HubResponseArrayOutput)
}

func (o EffectiveConnectivityConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveConnectivityConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EffectiveConnectivityConfigurationResponseOutput) IsGlobal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveConnectivityConfigurationResponse) *string { return v.IsGlobal }).(pulumi.StringPtrOutput)
}

func (o EffectiveConnectivityConfigurationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveConnectivityConfigurationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type EffectiveConnectivityConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (EffectiveConnectivityConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EffectiveConnectivityConfigurationResponse)(nil)).Elem()
}

func (o EffectiveConnectivityConfigurationResponseArrayOutput) ToEffectiveConnectivityConfigurationResponseArrayOutput() EffectiveConnectivityConfigurationResponseArrayOutput {
	return o
}

func (o EffectiveConnectivityConfigurationResponseArrayOutput) ToEffectiveConnectivityConfigurationResponseArrayOutputWithContext(ctx context.Context) EffectiveConnectivityConfigurationResponseArrayOutput {
	return o
}

func (o EffectiveConnectivityConfigurationResponseArrayOutput) Index(i pulumi.IntInput) EffectiveConnectivityConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EffectiveConnectivityConfigurationResponse {
		return vs[0].([]EffectiveConnectivityConfigurationResponse)[vs[1].(int)]
	}).(EffectiveConnectivityConfigurationResponseOutput)
}

type EffectiveDefaultSecurityAdminRuleResponse struct {
	Access                        string                                    `pulumi:"access"`
	ConfigurationDescription      *string                                   `pulumi:"configurationDescription"`
	Description                   string                                    `pulumi:"description"`
	DestinationPortRanges         []string                                  `pulumi:"destinationPortRanges"`
	Destinations                  []AddressPrefixItemResponse               `pulumi:"destinations"`
	Direction                     string                                    `pulumi:"direction"`
	Flag                          *string                                   `pulumi:"flag"`
	Id                            *string                                   `pulumi:"id"`
	Kind                          string                                    `pulumi:"kind"`
	Priority                      int                                       `pulumi:"priority"`
	Protocol                      string                                    `pulumi:"protocol"`
	ProvisioningState             string                                    `pulumi:"provisioningState"`
	RuleCollectionAppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     *string                                   `pulumi:"ruleCollectionDescription"`
	RuleGroups                    []ConfigurationGroupResponse              `pulumi:"ruleGroups"`
	SourcePortRanges              []string                                  `pulumi:"sourcePortRanges"`
	Sources                       []AddressPrefixItemResponse               `pulumi:"sources"`
}

type EffectiveSecurityAdminRuleResponse struct {
	Access                        string                                    `pulumi:"access"`
	ConfigurationDescription      *string                                   `pulumi:"configurationDescription"`
	Description                   *string                                   `pulumi:"description"`
	DestinationPortRanges         []string                                  `pulumi:"destinationPortRanges"`
	Destinations                  []AddressPrefixItemResponse               `pulumi:"destinations"`
	Direction                     string                                    `pulumi:"direction"`
	Id                            *string                                   `pulumi:"id"`
	Kind                          string                                    `pulumi:"kind"`
	Priority                      int                                       `pulumi:"priority"`
	Protocol                      string                                    `pulumi:"protocol"`
	ProvisioningState             string                                    `pulumi:"provisioningState"`
	RuleCollectionAppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     *string                                   `pulumi:"ruleCollectionDescription"`
	RuleGroups                    []ConfigurationGroupResponse              `pulumi:"ruleGroups"`
	SourcePortRanges              []string                                  `pulumi:"sourcePortRanges"`
	Sources                       []AddressPrefixItemResponse               `pulumi:"sources"`
}

type EffectiveVirtualNetworkResponse struct {
	Id             *string `pulumi:"id"`
	Location       *string `pulumi:"location"`
	MembershipType *string `pulumi:"membershipType"`
}

type EffectiveVirtualNetworkResponseOutput struct{ *pulumi.OutputState }

func (EffectiveVirtualNetworkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EffectiveVirtualNetworkResponse)(nil)).Elem()
}

func (o EffectiveVirtualNetworkResponseOutput) ToEffectiveVirtualNetworkResponseOutput() EffectiveVirtualNetworkResponseOutput {
	return o
}

func (o EffectiveVirtualNetworkResponseOutput) ToEffectiveVirtualNetworkResponseOutputWithContext(ctx context.Context) EffectiveVirtualNetworkResponseOutput {
	return o
}

func (o EffectiveVirtualNetworkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveVirtualNetworkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EffectiveVirtualNetworkResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveVirtualNetworkResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o EffectiveVirtualNetworkResponseOutput) MembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveVirtualNetworkResponse) *string { return v.MembershipType }).(pulumi.StringPtrOutput)
}

type EffectiveVirtualNetworkResponseArrayOutput struct{ *pulumi.OutputState }

func (EffectiveVirtualNetworkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EffectiveVirtualNetworkResponse)(nil)).Elem()
}

func (o EffectiveVirtualNetworkResponseArrayOutput) ToEffectiveVirtualNetworkResponseArrayOutput() EffectiveVirtualNetworkResponseArrayOutput {
	return o
}

func (o EffectiveVirtualNetworkResponseArrayOutput) ToEffectiveVirtualNetworkResponseArrayOutputWithContext(ctx context.Context) EffectiveVirtualNetworkResponseArrayOutput {
	return o
}

func (o EffectiveVirtualNetworkResponseArrayOutput) Index(i pulumi.IntInput) EffectiveVirtualNetworkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EffectiveVirtualNetworkResponse {
		return vs[0].([]EffectiveVirtualNetworkResponse)[vs[1].(int)]
	}).(EffectiveVirtualNetworkResponseOutput)
}

type EndpointType struct {
	AlwaysServe           *string                           `pulumi:"alwaysServe"`
	CustomHeaders         []EndpointPropertiesCustomHeaders `pulumi:"customHeaders"`
	EndpointLocation      *string                           `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string                           `pulumi:"endpointMonitorStatus"`
	EndpointStatus        *string                           `pulumi:"endpointStatus"`
	GeoMapping            []string                          `pulumi:"geoMapping"`
	Id                    *string                           `pulumi:"id"`
	MinChildEndpoints     *float64                          `pulumi:"minChildEndpoints"`
	MinChildEndpointsIPv4 *float64                          `pulumi:"minChildEndpointsIPv4"`
	MinChildEndpointsIPv6 *float64                          `pulumi:"minChildEndpointsIPv6"`
	Name                  *string                           `pulumi:"name"`
	Priority              *float64                          `pulumi:"priority"`
	Subnets               []EndpointPropertiesSubnets       `pulumi:"subnets"`
	Target                *string                           `pulumi:"target"`
	TargetResourceId      *string                           `pulumi:"targetResourceId"`
	Type                  *string                           `pulumi:"type"`
	Weight                *float64                          `pulumi:"weight"`
}





type EndpointTypeInput interface {
	pulumi.Input

	ToEndpointTypeOutput() EndpointTypeOutput
	ToEndpointTypeOutputWithContext(context.Context) EndpointTypeOutput
}

type EndpointTypeArgs struct {
	AlwaysServe           pulumi.StringPtrInput                     `pulumi:"alwaysServe"`
	CustomHeaders         EndpointPropertiesCustomHeadersArrayInput `pulumi:"customHeaders"`
	EndpointLocation      pulumi.StringPtrInput                     `pulumi:"endpointLocation"`
	EndpointMonitorStatus pulumi.StringPtrInput                     `pulumi:"endpointMonitorStatus"`
	EndpointStatus        pulumi.StringPtrInput                     `pulumi:"endpointStatus"`
	GeoMapping            pulumi.StringArrayInput                   `pulumi:"geoMapping"`
	Id                    pulumi.StringPtrInput                     `pulumi:"id"`
	MinChildEndpoints     pulumi.Float64PtrInput                    `pulumi:"minChildEndpoints"`
	MinChildEndpointsIPv4 pulumi.Float64PtrInput                    `pulumi:"minChildEndpointsIPv4"`
	MinChildEndpointsIPv6 pulumi.Float64PtrInput                    `pulumi:"minChildEndpointsIPv6"`
	Name                  pulumi.StringPtrInput                     `pulumi:"name"`
	Priority              pulumi.Float64PtrInput                    `pulumi:"priority"`
	Subnets               EndpointPropertiesSubnetsArrayInput       `pulumi:"subnets"`
	Target                pulumi.StringPtrInput                     `pulumi:"target"`
	TargetResourceId      pulumi.StringPtrInput                     `pulumi:"targetResourceId"`
	Type                  pulumi.StringPtrInput                     `pulumi:"type"`
	Weight                pulumi.Float64PtrInput                    `pulumi:"weight"`
}

func (EndpointTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointType)(nil)).Elem()
}

func (i EndpointTypeArgs) ToEndpointTypeOutput() EndpointTypeOutput {
	return i.ToEndpointTypeOutputWithContext(context.Background())
}

func (i EndpointTypeArgs) ToEndpointTypeOutputWithContext(ctx context.Context) EndpointTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointTypeOutput)
}





type EndpointTypeArrayInput interface {
	pulumi.Input

	ToEndpointTypeArrayOutput() EndpointTypeArrayOutput
	ToEndpointTypeArrayOutputWithContext(context.Context) EndpointTypeArrayOutput
}

type EndpointTypeArray []EndpointTypeInput

func (EndpointTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointType)(nil)).Elem()
}

func (i EndpointTypeArray) ToEndpointTypeArrayOutput() EndpointTypeArrayOutput {
	return i.ToEndpointTypeArrayOutputWithContext(context.Background())
}

func (i EndpointTypeArray) ToEndpointTypeArrayOutputWithContext(ctx context.Context) EndpointTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointTypeArrayOutput)
}

type EndpointTypeOutput struct{ *pulumi.OutputState }

func (EndpointTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointType)(nil)).Elem()
}

func (o EndpointTypeOutput) ToEndpointTypeOutput() EndpointTypeOutput {
	return o
}

func (o EndpointTypeOutput) ToEndpointTypeOutputWithContext(ctx context.Context) EndpointTypeOutput {
	return o
}

func (o EndpointTypeOutput) AlwaysServe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.AlwaysServe }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) CustomHeaders() EndpointPropertiesCustomHeadersArrayOutput {
	return o.ApplyT(func(v EndpointType) []EndpointPropertiesCustomHeaders { return v.CustomHeaders }).(EndpointPropertiesCustomHeadersArrayOutput)
}

func (o EndpointTypeOutput) EndpointLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.EndpointLocation }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) EndpointMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.EndpointMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) EndpointStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.EndpointStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) GeoMapping() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EndpointType) []string { return v.GeoMapping }).(pulumi.StringArrayOutput)
}

func (o EndpointTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) MinChildEndpoints() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointType) *float64 { return v.MinChildEndpoints }).(pulumi.Float64PtrOutput)
}

func (o EndpointTypeOutput) MinChildEndpointsIPv4() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointType) *float64 { return v.MinChildEndpointsIPv4 }).(pulumi.Float64PtrOutput)
}

func (o EndpointTypeOutput) MinChildEndpointsIPv6() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointType) *float64 { return v.MinChildEndpointsIPv6 }).(pulumi.Float64PtrOutput)
}

func (o EndpointTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) Priority() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointType) *float64 { return v.Priority }).(pulumi.Float64PtrOutput)
}

func (o EndpointTypeOutput) Subnets() EndpointPropertiesSubnetsArrayOutput {
	return o.ApplyT(func(v EndpointType) []EndpointPropertiesSubnets { return v.Subnets }).(EndpointPropertiesSubnetsArrayOutput)
}

func (o EndpointTypeOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) Weight() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointType) *float64 { return v.Weight }).(pulumi.Float64PtrOutput)
}

type EndpointTypeArrayOutput struct{ *pulumi.OutputState }

func (EndpointTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointType)(nil)).Elem()
}

func (o EndpointTypeArrayOutput) ToEndpointTypeArrayOutput() EndpointTypeArrayOutput {
	return o
}

func (o EndpointTypeArrayOutput) ToEndpointTypeArrayOutputWithContext(ctx context.Context) EndpointTypeArrayOutput {
	return o
}

func (o EndpointTypeArrayOutput) Index(i pulumi.IntInput) EndpointTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointType {
		return vs[0].([]EndpointType)[vs[1].(int)]
	}).(EndpointTypeOutput)
}

type EndpointPropertiesCustomHeaders struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type EndpointPropertiesCustomHeadersInput interface {
	pulumi.Input

	ToEndpointPropertiesCustomHeadersOutput() EndpointPropertiesCustomHeadersOutput
	ToEndpointPropertiesCustomHeadersOutputWithContext(context.Context) EndpointPropertiesCustomHeadersOutput
}

type EndpointPropertiesCustomHeadersArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (EndpointPropertiesCustomHeadersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesCustomHeaders)(nil)).Elem()
}

func (i EndpointPropertiesCustomHeadersArgs) ToEndpointPropertiesCustomHeadersOutput() EndpointPropertiesCustomHeadersOutput {
	return i.ToEndpointPropertiesCustomHeadersOutputWithContext(context.Background())
}

func (i EndpointPropertiesCustomHeadersArgs) ToEndpointPropertiesCustomHeadersOutputWithContext(ctx context.Context) EndpointPropertiesCustomHeadersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesCustomHeadersOutput)
}





type EndpointPropertiesCustomHeadersArrayInput interface {
	pulumi.Input

	ToEndpointPropertiesCustomHeadersArrayOutput() EndpointPropertiesCustomHeadersArrayOutput
	ToEndpointPropertiesCustomHeadersArrayOutputWithContext(context.Context) EndpointPropertiesCustomHeadersArrayOutput
}

type EndpointPropertiesCustomHeadersArray []EndpointPropertiesCustomHeadersInput

func (EndpointPropertiesCustomHeadersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointPropertiesCustomHeaders)(nil)).Elem()
}

func (i EndpointPropertiesCustomHeadersArray) ToEndpointPropertiesCustomHeadersArrayOutput() EndpointPropertiesCustomHeadersArrayOutput {
	return i.ToEndpointPropertiesCustomHeadersArrayOutputWithContext(context.Background())
}

func (i EndpointPropertiesCustomHeadersArray) ToEndpointPropertiesCustomHeadersArrayOutputWithContext(ctx context.Context) EndpointPropertiesCustomHeadersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesCustomHeadersArrayOutput)
}

type EndpointPropertiesCustomHeadersOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesCustomHeadersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesCustomHeaders)(nil)).Elem()
}

func (o EndpointPropertiesCustomHeadersOutput) ToEndpointPropertiesCustomHeadersOutput() EndpointPropertiesCustomHeadersOutput {
	return o
}

func (o EndpointPropertiesCustomHeadersOutput) ToEndpointPropertiesCustomHeadersOutputWithContext(ctx context.Context) EndpointPropertiesCustomHeadersOutput {
	return o
}

func (o EndpointPropertiesCustomHeadersOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesCustomHeaders) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EndpointPropertiesCustomHeadersOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesCustomHeaders) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EndpointPropertiesCustomHeadersArrayOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesCustomHeadersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointPropertiesCustomHeaders)(nil)).Elem()
}

func (o EndpointPropertiesCustomHeadersArrayOutput) ToEndpointPropertiesCustomHeadersArrayOutput() EndpointPropertiesCustomHeadersArrayOutput {
	return o
}

func (o EndpointPropertiesCustomHeadersArrayOutput) ToEndpointPropertiesCustomHeadersArrayOutputWithContext(ctx context.Context) EndpointPropertiesCustomHeadersArrayOutput {
	return o
}

func (o EndpointPropertiesCustomHeadersArrayOutput) Index(i pulumi.IntInput) EndpointPropertiesCustomHeadersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointPropertiesCustomHeaders {
		return vs[0].([]EndpointPropertiesCustomHeaders)[vs[1].(int)]
	}).(EndpointPropertiesCustomHeadersOutput)
}

type EndpointPropertiesResponseCustomHeaders struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type EndpointPropertiesResponseCustomHeadersOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesResponseCustomHeadersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesResponseCustomHeaders)(nil)).Elem()
}

func (o EndpointPropertiesResponseCustomHeadersOutput) ToEndpointPropertiesResponseCustomHeadersOutput() EndpointPropertiesResponseCustomHeadersOutput {
	return o
}

func (o EndpointPropertiesResponseCustomHeadersOutput) ToEndpointPropertiesResponseCustomHeadersOutputWithContext(ctx context.Context) EndpointPropertiesResponseCustomHeadersOutput {
	return o
}

func (o EndpointPropertiesResponseCustomHeadersOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesResponseCustomHeaders) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EndpointPropertiesResponseCustomHeadersOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesResponseCustomHeaders) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EndpointPropertiesResponseCustomHeadersArrayOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesResponseCustomHeadersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointPropertiesResponseCustomHeaders)(nil)).Elem()
}

func (o EndpointPropertiesResponseCustomHeadersArrayOutput) ToEndpointPropertiesResponseCustomHeadersArrayOutput() EndpointPropertiesResponseCustomHeadersArrayOutput {
	return o
}

func (o EndpointPropertiesResponseCustomHeadersArrayOutput) ToEndpointPropertiesResponseCustomHeadersArrayOutputWithContext(ctx context.Context) EndpointPropertiesResponseCustomHeadersArrayOutput {
	return o
}

func (o EndpointPropertiesResponseCustomHeadersArrayOutput) Index(i pulumi.IntInput) EndpointPropertiesResponseCustomHeadersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointPropertiesResponseCustomHeaders {
		return vs[0].([]EndpointPropertiesResponseCustomHeaders)[vs[1].(int)]
	}).(EndpointPropertiesResponseCustomHeadersOutput)
}

type EndpointPropertiesResponseSubnets struct {
	First *string `pulumi:"first"`
	Last  *string `pulumi:"last"`
	Scope *int    `pulumi:"scope"`
}

type EndpointPropertiesResponseSubnetsOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesResponseSubnetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesResponseSubnets)(nil)).Elem()
}

func (o EndpointPropertiesResponseSubnetsOutput) ToEndpointPropertiesResponseSubnetsOutput() EndpointPropertiesResponseSubnetsOutput {
	return o
}

func (o EndpointPropertiesResponseSubnetsOutput) ToEndpointPropertiesResponseSubnetsOutputWithContext(ctx context.Context) EndpointPropertiesResponseSubnetsOutput {
	return o
}

func (o EndpointPropertiesResponseSubnetsOutput) First() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesResponseSubnets) *string { return v.First }).(pulumi.StringPtrOutput)
}

func (o EndpointPropertiesResponseSubnetsOutput) Last() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesResponseSubnets) *string { return v.Last }).(pulumi.StringPtrOutput)
}

func (o EndpointPropertiesResponseSubnetsOutput) Scope() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesResponseSubnets) *int { return v.Scope }).(pulumi.IntPtrOutput)
}

type EndpointPropertiesResponseSubnetsArrayOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesResponseSubnetsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointPropertiesResponseSubnets)(nil)).Elem()
}

func (o EndpointPropertiesResponseSubnetsArrayOutput) ToEndpointPropertiesResponseSubnetsArrayOutput() EndpointPropertiesResponseSubnetsArrayOutput {
	return o
}

func (o EndpointPropertiesResponseSubnetsArrayOutput) ToEndpointPropertiesResponseSubnetsArrayOutputWithContext(ctx context.Context) EndpointPropertiesResponseSubnetsArrayOutput {
	return o
}

func (o EndpointPropertiesResponseSubnetsArrayOutput) Index(i pulumi.IntInput) EndpointPropertiesResponseSubnetsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointPropertiesResponseSubnets {
		return vs[0].([]EndpointPropertiesResponseSubnets)[vs[1].(int)]
	}).(EndpointPropertiesResponseSubnetsOutput)
}

type EndpointPropertiesSubnets struct {
	First *string `pulumi:"first"`
	Last  *string `pulumi:"last"`
	Scope *int    `pulumi:"scope"`
}





type EndpointPropertiesSubnetsInput interface {
	pulumi.Input

	ToEndpointPropertiesSubnetsOutput() EndpointPropertiesSubnetsOutput
	ToEndpointPropertiesSubnetsOutputWithContext(context.Context) EndpointPropertiesSubnetsOutput
}

type EndpointPropertiesSubnetsArgs struct {
	First pulumi.StringPtrInput `pulumi:"first"`
	Last  pulumi.StringPtrInput `pulumi:"last"`
	Scope pulumi.IntPtrInput    `pulumi:"scope"`
}

func (EndpointPropertiesSubnetsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesSubnets)(nil)).Elem()
}

func (i EndpointPropertiesSubnetsArgs) ToEndpointPropertiesSubnetsOutput() EndpointPropertiesSubnetsOutput {
	return i.ToEndpointPropertiesSubnetsOutputWithContext(context.Background())
}

func (i EndpointPropertiesSubnetsArgs) ToEndpointPropertiesSubnetsOutputWithContext(ctx context.Context) EndpointPropertiesSubnetsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesSubnetsOutput)
}





type EndpointPropertiesSubnetsArrayInput interface {
	pulumi.Input

	ToEndpointPropertiesSubnetsArrayOutput() EndpointPropertiesSubnetsArrayOutput
	ToEndpointPropertiesSubnetsArrayOutputWithContext(context.Context) EndpointPropertiesSubnetsArrayOutput
}

type EndpointPropertiesSubnetsArray []EndpointPropertiesSubnetsInput

func (EndpointPropertiesSubnetsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointPropertiesSubnets)(nil)).Elem()
}

func (i EndpointPropertiesSubnetsArray) ToEndpointPropertiesSubnetsArrayOutput() EndpointPropertiesSubnetsArrayOutput {
	return i.ToEndpointPropertiesSubnetsArrayOutputWithContext(context.Background())
}

func (i EndpointPropertiesSubnetsArray) ToEndpointPropertiesSubnetsArrayOutputWithContext(ctx context.Context) EndpointPropertiesSubnetsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesSubnetsArrayOutput)
}

type EndpointPropertiesSubnetsOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesSubnetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesSubnets)(nil)).Elem()
}

func (o EndpointPropertiesSubnetsOutput) ToEndpointPropertiesSubnetsOutput() EndpointPropertiesSubnetsOutput {
	return o
}

func (o EndpointPropertiesSubnetsOutput) ToEndpointPropertiesSubnetsOutputWithContext(ctx context.Context) EndpointPropertiesSubnetsOutput {
	return o
}

func (o EndpointPropertiesSubnetsOutput) First() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesSubnets) *string { return v.First }).(pulumi.StringPtrOutput)
}

func (o EndpointPropertiesSubnetsOutput) Last() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesSubnets) *string { return v.Last }).(pulumi.StringPtrOutput)
}

func (o EndpointPropertiesSubnetsOutput) Scope() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesSubnets) *int { return v.Scope }).(pulumi.IntPtrOutput)
}

type EndpointPropertiesSubnetsArrayOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesSubnetsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointPropertiesSubnets)(nil)).Elem()
}

func (o EndpointPropertiesSubnetsArrayOutput) ToEndpointPropertiesSubnetsArrayOutput() EndpointPropertiesSubnetsArrayOutput {
	return o
}

func (o EndpointPropertiesSubnetsArrayOutput) ToEndpointPropertiesSubnetsArrayOutputWithContext(ctx context.Context) EndpointPropertiesSubnetsArrayOutput {
	return o
}

func (o EndpointPropertiesSubnetsArrayOutput) Index(i pulumi.IntInput) EndpointPropertiesSubnetsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointPropertiesSubnets {
		return vs[0].([]EndpointPropertiesSubnets)[vs[1].(int)]
	}).(EndpointPropertiesSubnetsOutput)
}

type EndpointResponse struct {
	AlwaysServe           *string                                   `pulumi:"alwaysServe"`
	CustomHeaders         []EndpointPropertiesResponseCustomHeaders `pulumi:"customHeaders"`
	EndpointLocation      *string                                   `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string                                   `pulumi:"endpointMonitorStatus"`
	EndpointStatus        *string                                   `pulumi:"endpointStatus"`
	GeoMapping            []string                                  `pulumi:"geoMapping"`
	Id                    *string                                   `pulumi:"id"`
	MinChildEndpoints     *float64                                  `pulumi:"minChildEndpoints"`
	MinChildEndpointsIPv4 *float64                                  `pulumi:"minChildEndpointsIPv4"`
	MinChildEndpointsIPv6 *float64                                  `pulumi:"minChildEndpointsIPv6"`
	Name                  *string                                   `pulumi:"name"`
	Priority              *float64                                  `pulumi:"priority"`
	Subnets               []EndpointPropertiesResponseSubnets       `pulumi:"subnets"`
	Target                *string                                   `pulumi:"target"`
	TargetResourceId      *string                                   `pulumi:"targetResourceId"`
	Type                  *string                                   `pulumi:"type"`
	Weight                *float64                                  `pulumi:"weight"`
}

type EndpointResponseOutput struct{ *pulumi.OutputState }

func (EndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointResponse)(nil)).Elem()
}

func (o EndpointResponseOutput) ToEndpointResponseOutput() EndpointResponseOutput {
	return o
}

func (o EndpointResponseOutput) ToEndpointResponseOutputWithContext(ctx context.Context) EndpointResponseOutput {
	return o
}

func (o EndpointResponseOutput) AlwaysServe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.AlwaysServe }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) CustomHeaders() EndpointPropertiesResponseCustomHeadersArrayOutput {
	return o.ApplyT(func(v EndpointResponse) []EndpointPropertiesResponseCustomHeaders { return v.CustomHeaders }).(EndpointPropertiesResponseCustomHeadersArrayOutput)
}

func (o EndpointResponseOutput) EndpointLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.EndpointLocation }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) EndpointMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.EndpointMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) EndpointStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.EndpointStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) GeoMapping() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EndpointResponse) []string { return v.GeoMapping }).(pulumi.StringArrayOutput)
}

func (o EndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) MinChildEndpoints() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointResponse) *float64 { return v.MinChildEndpoints }).(pulumi.Float64PtrOutput)
}

func (o EndpointResponseOutput) MinChildEndpointsIPv4() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointResponse) *float64 { return v.MinChildEndpointsIPv4 }).(pulumi.Float64PtrOutput)
}

func (o EndpointResponseOutput) MinChildEndpointsIPv6() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointResponse) *float64 { return v.MinChildEndpointsIPv6 }).(pulumi.Float64PtrOutput)
}

func (o EndpointResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) Priority() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointResponse) *float64 { return v.Priority }).(pulumi.Float64PtrOutput)
}

func (o EndpointResponseOutput) Subnets() EndpointPropertiesResponseSubnetsArrayOutput {
	return o.ApplyT(func(v EndpointResponse) []EndpointPropertiesResponseSubnets { return v.Subnets }).(EndpointPropertiesResponseSubnetsArrayOutput)
}

func (o EndpointResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) Weight() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointResponse) *float64 { return v.Weight }).(pulumi.Float64PtrOutput)
}

type EndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (EndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointResponse)(nil)).Elem()
}

func (o EndpointResponseArrayOutput) ToEndpointResponseArrayOutput() EndpointResponseArrayOutput {
	return o
}

func (o EndpointResponseArrayOutput) ToEndpointResponseArrayOutputWithContext(ctx context.Context) EndpointResponseArrayOutput {
	return o
}

func (o EndpointResponseArrayOutput) Index(i pulumi.IntInput) EndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointResponse {
		return vs[0].([]EndpointResponse)[vs[1].(int)]
	}).(EndpointResponseOutput)
}

type Hub struct {
	ResourceId   *string `pulumi:"resourceId"`
	ResourceType *string `pulumi:"resourceType"`
}





type HubInput interface {
	pulumi.Input

	ToHubOutput() HubOutput
	ToHubOutputWithContext(context.Context) HubOutput
}

type HubArgs struct {
	ResourceId   pulumi.StringPtrInput `pulumi:"resourceId"`
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
}

func (HubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Hub)(nil)).Elem()
}

func (i HubArgs) ToHubOutput() HubOutput {
	return i.ToHubOutputWithContext(context.Background())
}

func (i HubArgs) ToHubOutputWithContext(ctx context.Context) HubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubOutput)
}





type HubArrayInput interface {
	pulumi.Input

	ToHubArrayOutput() HubArrayOutput
	ToHubArrayOutputWithContext(context.Context) HubArrayOutput
}

type HubArray []HubInput

func (HubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Hub)(nil)).Elem()
}

func (i HubArray) ToHubArrayOutput() HubArrayOutput {
	return i.ToHubArrayOutputWithContext(context.Background())
}

func (i HubArray) ToHubArrayOutputWithContext(ctx context.Context) HubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubArrayOutput)
}

type HubOutput struct{ *pulumi.OutputState }

func (HubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Hub)(nil)).Elem()
}

func (o HubOutput) ToHubOutput() HubOutput {
	return o
}

func (o HubOutput) ToHubOutputWithContext(ctx context.Context) HubOutput {
	return o
}

func (o HubOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Hub) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o HubOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Hub) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

type HubArrayOutput struct{ *pulumi.OutputState }

func (HubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Hub)(nil)).Elem()
}

func (o HubArrayOutput) ToHubArrayOutput() HubArrayOutput {
	return o
}

func (o HubArrayOutput) ToHubArrayOutputWithContext(ctx context.Context) HubArrayOutput {
	return o
}

func (o HubArrayOutput) Index(i pulumi.IntInput) HubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Hub {
		return vs[0].([]Hub)[vs[1].(int)]
	}).(HubOutput)
}

type HubResponse struct {
	ResourceId   *string `pulumi:"resourceId"`
	ResourceType *string `pulumi:"resourceType"`
}

type HubResponseOutput struct{ *pulumi.OutputState }

func (HubResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HubResponse)(nil)).Elem()
}

func (o HubResponseOutput) ToHubResponseOutput() HubResponseOutput {
	return o
}

func (o HubResponseOutput) ToHubResponseOutputWithContext(ctx context.Context) HubResponseOutput {
	return o
}

func (o HubResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HubResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o HubResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HubResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

type HubResponseArrayOutput struct{ *pulumi.OutputState }

func (HubResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HubResponse)(nil)).Elem()
}

func (o HubResponseArrayOutput) ToHubResponseArrayOutput() HubResponseArrayOutput {
	return o
}

func (o HubResponseArrayOutput) ToHubResponseArrayOutputWithContext(ctx context.Context) HubResponseArrayOutput {
	return o
}

func (o HubResponseArrayOutput) Index(i pulumi.IntInput) HubResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HubResponse {
		return vs[0].([]HubResponse)[vs[1].(int)]
	}).(HubResponseOutput)
}

type MonitorConfig struct {
	CustomHeaders             []MonitorConfigCustomHeaders            `pulumi:"customHeaders"`
	ExpectedStatusCodeRanges  []MonitorConfigExpectedStatusCodeRanges `pulumi:"expectedStatusCodeRanges"`
	IntervalInSeconds         *float64                                `pulumi:"intervalInSeconds"`
	Path                      *string                                 `pulumi:"path"`
	Port                      *float64                                `pulumi:"port"`
	ProfileMonitorStatus      *string                                 `pulumi:"profileMonitorStatus"`
	Protocol                  *string                                 `pulumi:"protocol"`
	TimeoutInSeconds          *float64                                `pulumi:"timeoutInSeconds"`
	ToleratedNumberOfFailures *float64                                `pulumi:"toleratedNumberOfFailures"`
}





type MonitorConfigInput interface {
	pulumi.Input

	ToMonitorConfigOutput() MonitorConfigOutput
	ToMonitorConfigOutputWithContext(context.Context) MonitorConfigOutput
}

type MonitorConfigArgs struct {
	CustomHeaders             MonitorConfigCustomHeadersArrayInput            `pulumi:"customHeaders"`
	ExpectedStatusCodeRanges  MonitorConfigExpectedStatusCodeRangesArrayInput `pulumi:"expectedStatusCodeRanges"`
	IntervalInSeconds         pulumi.Float64PtrInput                          `pulumi:"intervalInSeconds"`
	Path                      pulumi.StringPtrInput                           `pulumi:"path"`
	Port                      pulumi.Float64PtrInput                          `pulumi:"port"`
	ProfileMonitorStatus      pulumi.StringPtrInput                           `pulumi:"profileMonitorStatus"`
	Protocol                  pulumi.StringPtrInput                           `pulumi:"protocol"`
	TimeoutInSeconds          pulumi.Float64PtrInput                          `pulumi:"timeoutInSeconds"`
	ToleratedNumberOfFailures pulumi.Float64PtrInput                          `pulumi:"toleratedNumberOfFailures"`
}

func (MonitorConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfig)(nil)).Elem()
}

func (i MonitorConfigArgs) ToMonitorConfigOutput() MonitorConfigOutput {
	return i.ToMonitorConfigOutputWithContext(context.Background())
}

func (i MonitorConfigArgs) ToMonitorConfigOutputWithContext(ctx context.Context) MonitorConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigOutput)
}

func (i MonitorConfigArgs) ToMonitorConfigPtrOutput() MonitorConfigPtrOutput {
	return i.ToMonitorConfigPtrOutputWithContext(context.Background())
}

func (i MonitorConfigArgs) ToMonitorConfigPtrOutputWithContext(ctx context.Context) MonitorConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigOutput).ToMonitorConfigPtrOutputWithContext(ctx)
}









type MonitorConfigPtrInput interface {
	pulumi.Input

	ToMonitorConfigPtrOutput() MonitorConfigPtrOutput
	ToMonitorConfigPtrOutputWithContext(context.Context) MonitorConfigPtrOutput
}

type monitorConfigPtrType MonitorConfigArgs

func MonitorConfigPtr(v *MonitorConfigArgs) MonitorConfigPtrInput {
	return (*monitorConfigPtrType)(v)
}

func (*monitorConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorConfig)(nil)).Elem()
}

func (i *monitorConfigPtrType) ToMonitorConfigPtrOutput() MonitorConfigPtrOutput {
	return i.ToMonitorConfigPtrOutputWithContext(context.Background())
}

func (i *monitorConfigPtrType) ToMonitorConfigPtrOutputWithContext(ctx context.Context) MonitorConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigPtrOutput)
}

type MonitorConfigOutput struct{ *pulumi.OutputState }

func (MonitorConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfig)(nil)).Elem()
}

func (o MonitorConfigOutput) ToMonitorConfigOutput() MonitorConfigOutput {
	return o
}

func (o MonitorConfigOutput) ToMonitorConfigOutputWithContext(ctx context.Context) MonitorConfigOutput {
	return o
}

func (o MonitorConfigOutput) ToMonitorConfigPtrOutput() MonitorConfigPtrOutput {
	return o.ToMonitorConfigPtrOutputWithContext(context.Background())
}

func (o MonitorConfigOutput) ToMonitorConfigPtrOutputWithContext(ctx context.Context) MonitorConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorConfig) *MonitorConfig {
		return &v
	}).(MonitorConfigPtrOutput)
}

func (o MonitorConfigOutput) CustomHeaders() MonitorConfigCustomHeadersArrayOutput {
	return o.ApplyT(func(v MonitorConfig) []MonitorConfigCustomHeaders { return v.CustomHeaders }).(MonitorConfigCustomHeadersArrayOutput)
}

func (o MonitorConfigOutput) ExpectedStatusCodeRanges() MonitorConfigExpectedStatusCodeRangesArrayOutput {
	return o.ApplyT(func(v MonitorConfig) []MonitorConfigExpectedStatusCodeRanges { return v.ExpectedStatusCodeRanges }).(MonitorConfigExpectedStatusCodeRangesArrayOutput)
}

func (o MonitorConfigOutput) IntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfig) *float64 { return v.IntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfig) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfig) *float64 { return v.Port }).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigOutput) ProfileMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfig) *string { return v.ProfileMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfig) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfig) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigOutput) ToleratedNumberOfFailures() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfig) *float64 { return v.ToleratedNumberOfFailures }).(pulumi.Float64PtrOutput)
}

type MonitorConfigPtrOutput struct{ *pulumi.OutputState }

func (MonitorConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorConfig)(nil)).Elem()
}

func (o MonitorConfigPtrOutput) ToMonitorConfigPtrOutput() MonitorConfigPtrOutput {
	return o
}

func (o MonitorConfigPtrOutput) ToMonitorConfigPtrOutputWithContext(ctx context.Context) MonitorConfigPtrOutput {
	return o
}

func (o MonitorConfigPtrOutput) Elem() MonitorConfigOutput {
	return o.ApplyT(func(v *MonitorConfig) MonitorConfig {
		if v != nil {
			return *v
		}
		var ret MonitorConfig
		return ret
	}).(MonitorConfigOutput)
}

func (o MonitorConfigPtrOutput) CustomHeaders() MonitorConfigCustomHeadersArrayOutput {
	return o.ApplyT(func(v *MonitorConfig) []MonitorConfigCustomHeaders {
		if v == nil {
			return nil
		}
		return v.CustomHeaders
	}).(MonitorConfigCustomHeadersArrayOutput)
}

func (o MonitorConfigPtrOutput) ExpectedStatusCodeRanges() MonitorConfigExpectedStatusCodeRangesArrayOutput {
	return o.ApplyT(func(v *MonitorConfig) []MonitorConfigExpectedStatusCodeRanges {
		if v == nil {
			return nil
		}
		return v.ExpectedStatusCodeRanges
	}).(MonitorConfigExpectedStatusCodeRangesArrayOutput)
}

func (o MonitorConfigPtrOutput) IntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.IntervalInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigPtrOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigPtrOutput) ProfileMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *string {
		if v == nil {
			return nil
		}
		return v.ProfileMonitorStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigPtrOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *string {
		if v == nil {
			return nil
		}
		return v.Protocol
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigPtrOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigPtrOutput) ToleratedNumberOfFailures() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.ToleratedNumberOfFailures
	}).(pulumi.Float64PtrOutput)
}

type MonitorConfigCustomHeaders struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type MonitorConfigCustomHeadersInput interface {
	pulumi.Input

	ToMonitorConfigCustomHeadersOutput() MonitorConfigCustomHeadersOutput
	ToMonitorConfigCustomHeadersOutputWithContext(context.Context) MonitorConfigCustomHeadersOutput
}

type MonitorConfigCustomHeadersArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (MonitorConfigCustomHeadersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigCustomHeaders)(nil)).Elem()
}

func (i MonitorConfigCustomHeadersArgs) ToMonitorConfigCustomHeadersOutput() MonitorConfigCustomHeadersOutput {
	return i.ToMonitorConfigCustomHeadersOutputWithContext(context.Background())
}

func (i MonitorConfigCustomHeadersArgs) ToMonitorConfigCustomHeadersOutputWithContext(ctx context.Context) MonitorConfigCustomHeadersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigCustomHeadersOutput)
}





type MonitorConfigCustomHeadersArrayInput interface {
	pulumi.Input

	ToMonitorConfigCustomHeadersArrayOutput() MonitorConfigCustomHeadersArrayOutput
	ToMonitorConfigCustomHeadersArrayOutputWithContext(context.Context) MonitorConfigCustomHeadersArrayOutput
}

type MonitorConfigCustomHeadersArray []MonitorConfigCustomHeadersInput

func (MonitorConfigCustomHeadersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigCustomHeaders)(nil)).Elem()
}

func (i MonitorConfigCustomHeadersArray) ToMonitorConfigCustomHeadersArrayOutput() MonitorConfigCustomHeadersArrayOutput {
	return i.ToMonitorConfigCustomHeadersArrayOutputWithContext(context.Background())
}

func (i MonitorConfigCustomHeadersArray) ToMonitorConfigCustomHeadersArrayOutputWithContext(ctx context.Context) MonitorConfigCustomHeadersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigCustomHeadersArrayOutput)
}

type MonitorConfigCustomHeadersOutput struct{ *pulumi.OutputState }

func (MonitorConfigCustomHeadersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigCustomHeaders)(nil)).Elem()
}

func (o MonitorConfigCustomHeadersOutput) ToMonitorConfigCustomHeadersOutput() MonitorConfigCustomHeadersOutput {
	return o
}

func (o MonitorConfigCustomHeadersOutput) ToMonitorConfigCustomHeadersOutputWithContext(ctx context.Context) MonitorConfigCustomHeadersOutput {
	return o
}

func (o MonitorConfigCustomHeadersOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigCustomHeaders) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigCustomHeadersOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigCustomHeaders) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type MonitorConfigCustomHeadersArrayOutput struct{ *pulumi.OutputState }

func (MonitorConfigCustomHeadersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigCustomHeaders)(nil)).Elem()
}

func (o MonitorConfigCustomHeadersArrayOutput) ToMonitorConfigCustomHeadersArrayOutput() MonitorConfigCustomHeadersArrayOutput {
	return o
}

func (o MonitorConfigCustomHeadersArrayOutput) ToMonitorConfigCustomHeadersArrayOutputWithContext(ctx context.Context) MonitorConfigCustomHeadersArrayOutput {
	return o
}

func (o MonitorConfigCustomHeadersArrayOutput) Index(i pulumi.IntInput) MonitorConfigCustomHeadersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonitorConfigCustomHeaders {
		return vs[0].([]MonitorConfigCustomHeaders)[vs[1].(int)]
	}).(MonitorConfigCustomHeadersOutput)
}

type MonitorConfigExpectedStatusCodeRanges struct {
	Max *int `pulumi:"max"`
	Min *int `pulumi:"min"`
}





type MonitorConfigExpectedStatusCodeRangesInput interface {
	pulumi.Input

	ToMonitorConfigExpectedStatusCodeRangesOutput() MonitorConfigExpectedStatusCodeRangesOutput
	ToMonitorConfigExpectedStatusCodeRangesOutputWithContext(context.Context) MonitorConfigExpectedStatusCodeRangesOutput
}

type MonitorConfigExpectedStatusCodeRangesArgs struct {
	Max pulumi.IntPtrInput `pulumi:"max"`
	Min pulumi.IntPtrInput `pulumi:"min"`
}

func (MonitorConfigExpectedStatusCodeRangesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigExpectedStatusCodeRanges)(nil)).Elem()
}

func (i MonitorConfigExpectedStatusCodeRangesArgs) ToMonitorConfigExpectedStatusCodeRangesOutput() MonitorConfigExpectedStatusCodeRangesOutput {
	return i.ToMonitorConfigExpectedStatusCodeRangesOutputWithContext(context.Background())
}

func (i MonitorConfigExpectedStatusCodeRangesArgs) ToMonitorConfigExpectedStatusCodeRangesOutputWithContext(ctx context.Context) MonitorConfigExpectedStatusCodeRangesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigExpectedStatusCodeRangesOutput)
}





type MonitorConfigExpectedStatusCodeRangesArrayInput interface {
	pulumi.Input

	ToMonitorConfigExpectedStatusCodeRangesArrayOutput() MonitorConfigExpectedStatusCodeRangesArrayOutput
	ToMonitorConfigExpectedStatusCodeRangesArrayOutputWithContext(context.Context) MonitorConfigExpectedStatusCodeRangesArrayOutput
}

type MonitorConfigExpectedStatusCodeRangesArray []MonitorConfigExpectedStatusCodeRangesInput

func (MonitorConfigExpectedStatusCodeRangesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigExpectedStatusCodeRanges)(nil)).Elem()
}

func (i MonitorConfigExpectedStatusCodeRangesArray) ToMonitorConfigExpectedStatusCodeRangesArrayOutput() MonitorConfigExpectedStatusCodeRangesArrayOutput {
	return i.ToMonitorConfigExpectedStatusCodeRangesArrayOutputWithContext(context.Background())
}

func (i MonitorConfigExpectedStatusCodeRangesArray) ToMonitorConfigExpectedStatusCodeRangesArrayOutputWithContext(ctx context.Context) MonitorConfigExpectedStatusCodeRangesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigExpectedStatusCodeRangesArrayOutput)
}

type MonitorConfigExpectedStatusCodeRangesOutput struct{ *pulumi.OutputState }

func (MonitorConfigExpectedStatusCodeRangesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigExpectedStatusCodeRanges)(nil)).Elem()
}

func (o MonitorConfigExpectedStatusCodeRangesOutput) ToMonitorConfigExpectedStatusCodeRangesOutput() MonitorConfigExpectedStatusCodeRangesOutput {
	return o
}

func (o MonitorConfigExpectedStatusCodeRangesOutput) ToMonitorConfigExpectedStatusCodeRangesOutputWithContext(ctx context.Context) MonitorConfigExpectedStatusCodeRangesOutput {
	return o
}

func (o MonitorConfigExpectedStatusCodeRangesOutput) Max() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonitorConfigExpectedStatusCodeRanges) *int { return v.Max }).(pulumi.IntPtrOutput)
}

func (o MonitorConfigExpectedStatusCodeRangesOutput) Min() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonitorConfigExpectedStatusCodeRanges) *int { return v.Min }).(pulumi.IntPtrOutput)
}

type MonitorConfigExpectedStatusCodeRangesArrayOutput struct{ *pulumi.OutputState }

func (MonitorConfigExpectedStatusCodeRangesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigExpectedStatusCodeRanges)(nil)).Elem()
}

func (o MonitorConfigExpectedStatusCodeRangesArrayOutput) ToMonitorConfigExpectedStatusCodeRangesArrayOutput() MonitorConfigExpectedStatusCodeRangesArrayOutput {
	return o
}

func (o MonitorConfigExpectedStatusCodeRangesArrayOutput) ToMonitorConfigExpectedStatusCodeRangesArrayOutputWithContext(ctx context.Context) MonitorConfigExpectedStatusCodeRangesArrayOutput {
	return o
}

func (o MonitorConfigExpectedStatusCodeRangesArrayOutput) Index(i pulumi.IntInput) MonitorConfigExpectedStatusCodeRangesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonitorConfigExpectedStatusCodeRanges {
		return vs[0].([]MonitorConfigExpectedStatusCodeRanges)[vs[1].(int)]
	}).(MonitorConfigExpectedStatusCodeRangesOutput)
}

type MonitorConfigResponse struct {
	CustomHeaders             []MonitorConfigResponseCustomHeaders            `pulumi:"customHeaders"`
	ExpectedStatusCodeRanges  []MonitorConfigResponseExpectedStatusCodeRanges `pulumi:"expectedStatusCodeRanges"`
	IntervalInSeconds         *float64                                        `pulumi:"intervalInSeconds"`
	Path                      *string                                         `pulumi:"path"`
	Port                      *float64                                        `pulumi:"port"`
	ProfileMonitorStatus      *string                                         `pulumi:"profileMonitorStatus"`
	Protocol                  *string                                         `pulumi:"protocol"`
	TimeoutInSeconds          *float64                                        `pulumi:"timeoutInSeconds"`
	ToleratedNumberOfFailures *float64                                        `pulumi:"toleratedNumberOfFailures"`
}

type MonitorConfigResponseOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigResponse)(nil)).Elem()
}

func (o MonitorConfigResponseOutput) ToMonitorConfigResponseOutput() MonitorConfigResponseOutput {
	return o
}

func (o MonitorConfigResponseOutput) ToMonitorConfigResponseOutputWithContext(ctx context.Context) MonitorConfigResponseOutput {
	return o
}

func (o MonitorConfigResponseOutput) CustomHeaders() MonitorConfigResponseCustomHeadersArrayOutput {
	return o.ApplyT(func(v MonitorConfigResponse) []MonitorConfigResponseCustomHeaders { return v.CustomHeaders }).(MonitorConfigResponseCustomHeadersArrayOutput)
}

func (o MonitorConfigResponseOutput) ExpectedStatusCodeRanges() MonitorConfigResponseExpectedStatusCodeRangesArrayOutput {
	return o.ApplyT(func(v MonitorConfigResponse) []MonitorConfigResponseExpectedStatusCodeRanges {
		return v.ExpectedStatusCodeRanges
	}).(MonitorConfigResponseExpectedStatusCodeRangesArrayOutput)
}

func (o MonitorConfigResponseOutput) IntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *float64 { return v.IntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponseOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *float64 { return v.Port }).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigResponseOutput) ProfileMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *string { return v.ProfileMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponseOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigResponseOutput) ToleratedNumberOfFailures() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *float64 { return v.ToleratedNumberOfFailures }).(pulumi.Float64PtrOutput)
}

type MonitorConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorConfigResponse)(nil)).Elem()
}

func (o MonitorConfigResponsePtrOutput) ToMonitorConfigResponsePtrOutput() MonitorConfigResponsePtrOutput {
	return o
}

func (o MonitorConfigResponsePtrOutput) ToMonitorConfigResponsePtrOutputWithContext(ctx context.Context) MonitorConfigResponsePtrOutput {
	return o
}

func (o MonitorConfigResponsePtrOutput) Elem() MonitorConfigResponseOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) MonitorConfigResponse {
		if v != nil {
			return *v
		}
		var ret MonitorConfigResponse
		return ret
	}).(MonitorConfigResponseOutput)
}

func (o MonitorConfigResponsePtrOutput) CustomHeaders() MonitorConfigResponseCustomHeadersArrayOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) []MonitorConfigResponseCustomHeaders {
		if v == nil {
			return nil
		}
		return v.CustomHeaders
	}).(MonitorConfigResponseCustomHeadersArrayOutput)
}

func (o MonitorConfigResponsePtrOutput) ExpectedStatusCodeRanges() MonitorConfigResponseExpectedStatusCodeRangesArrayOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) []MonitorConfigResponseExpectedStatusCodeRanges {
		if v == nil {
			return nil
		}
		return v.ExpectedStatusCodeRanges
	}).(MonitorConfigResponseExpectedStatusCodeRangesArrayOutput)
}

func (o MonitorConfigResponsePtrOutput) IntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.IntervalInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponsePtrOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigResponsePtrOutput) ProfileMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProfileMonitorStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponsePtrOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Protocol
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponsePtrOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigResponsePtrOutput) ToleratedNumberOfFailures() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.ToleratedNumberOfFailures
	}).(pulumi.Float64PtrOutput)
}

type MonitorConfigResponseCustomHeaders struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type MonitorConfigResponseCustomHeadersOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponseCustomHeadersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigResponseCustomHeaders)(nil)).Elem()
}

func (o MonitorConfigResponseCustomHeadersOutput) ToMonitorConfigResponseCustomHeadersOutput() MonitorConfigResponseCustomHeadersOutput {
	return o
}

func (o MonitorConfigResponseCustomHeadersOutput) ToMonitorConfigResponseCustomHeadersOutputWithContext(ctx context.Context) MonitorConfigResponseCustomHeadersOutput {
	return o
}

func (o MonitorConfigResponseCustomHeadersOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponseCustomHeaders) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponseCustomHeadersOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponseCustomHeaders) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type MonitorConfigResponseCustomHeadersArrayOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponseCustomHeadersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigResponseCustomHeaders)(nil)).Elem()
}

func (o MonitorConfigResponseCustomHeadersArrayOutput) ToMonitorConfigResponseCustomHeadersArrayOutput() MonitorConfigResponseCustomHeadersArrayOutput {
	return o
}

func (o MonitorConfigResponseCustomHeadersArrayOutput) ToMonitorConfigResponseCustomHeadersArrayOutputWithContext(ctx context.Context) MonitorConfigResponseCustomHeadersArrayOutput {
	return o
}

func (o MonitorConfigResponseCustomHeadersArrayOutput) Index(i pulumi.IntInput) MonitorConfigResponseCustomHeadersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonitorConfigResponseCustomHeaders {
		return vs[0].([]MonitorConfigResponseCustomHeaders)[vs[1].(int)]
	}).(MonitorConfigResponseCustomHeadersOutput)
}

type MonitorConfigResponseExpectedStatusCodeRanges struct {
	Max *int `pulumi:"max"`
	Min *int `pulumi:"min"`
}

type MonitorConfigResponseExpectedStatusCodeRangesOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponseExpectedStatusCodeRangesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigResponseExpectedStatusCodeRanges)(nil)).Elem()
}

func (o MonitorConfigResponseExpectedStatusCodeRangesOutput) ToMonitorConfigResponseExpectedStatusCodeRangesOutput() MonitorConfigResponseExpectedStatusCodeRangesOutput {
	return o
}

func (o MonitorConfigResponseExpectedStatusCodeRangesOutput) ToMonitorConfigResponseExpectedStatusCodeRangesOutputWithContext(ctx context.Context) MonitorConfigResponseExpectedStatusCodeRangesOutput {
	return o
}

func (o MonitorConfigResponseExpectedStatusCodeRangesOutput) Max() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponseExpectedStatusCodeRanges) *int { return v.Max }).(pulumi.IntPtrOutput)
}

func (o MonitorConfigResponseExpectedStatusCodeRangesOutput) Min() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponseExpectedStatusCodeRanges) *int { return v.Min }).(pulumi.IntPtrOutput)
}

type MonitorConfigResponseExpectedStatusCodeRangesArrayOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponseExpectedStatusCodeRangesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigResponseExpectedStatusCodeRanges)(nil)).Elem()
}

func (o MonitorConfigResponseExpectedStatusCodeRangesArrayOutput) ToMonitorConfigResponseExpectedStatusCodeRangesArrayOutput() MonitorConfigResponseExpectedStatusCodeRangesArrayOutput {
	return o
}

func (o MonitorConfigResponseExpectedStatusCodeRangesArrayOutput) ToMonitorConfigResponseExpectedStatusCodeRangesArrayOutputWithContext(ctx context.Context) MonitorConfigResponseExpectedStatusCodeRangesArrayOutput {
	return o
}

func (o MonitorConfigResponseExpectedStatusCodeRangesArrayOutput) Index(i pulumi.IntInput) MonitorConfigResponseExpectedStatusCodeRangesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonitorConfigResponseExpectedStatusCodeRanges {
		return vs[0].([]MonitorConfigResponseExpectedStatusCodeRanges)[vs[1].(int)]
	}).(MonitorConfigResponseExpectedStatusCodeRangesOutput)
}

type NetworkManagerDeploymentStatusResponse struct {
	CommitTime       *string  `pulumi:"commitTime"`
	ConfigurationIds []string `pulumi:"configurationIds"`
	DeploymentStatus *string  `pulumi:"deploymentStatus"`
	DeploymentType   *string  `pulumi:"deploymentType"`
	ErrorMessage     *string  `pulumi:"errorMessage"`
	Region           *string  `pulumi:"region"`
}

type NetworkManagerDeploymentStatusResponseOutput struct{ *pulumi.OutputState }

func (NetworkManagerDeploymentStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkManagerDeploymentStatusResponse)(nil)).Elem()
}

func (o NetworkManagerDeploymentStatusResponseOutput) ToNetworkManagerDeploymentStatusResponseOutput() NetworkManagerDeploymentStatusResponseOutput {
	return o
}

func (o NetworkManagerDeploymentStatusResponseOutput) ToNetworkManagerDeploymentStatusResponseOutputWithContext(ctx context.Context) NetworkManagerDeploymentStatusResponseOutput {
	return o
}

func (o NetworkManagerDeploymentStatusResponseOutput) CommitTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkManagerDeploymentStatusResponse) *string { return v.CommitTime }).(pulumi.StringPtrOutput)
}

func (o NetworkManagerDeploymentStatusResponseOutput) ConfigurationIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkManagerDeploymentStatusResponse) []string { return v.ConfigurationIds }).(pulumi.StringArrayOutput)
}

func (o NetworkManagerDeploymentStatusResponseOutput) DeploymentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkManagerDeploymentStatusResponse) *string { return v.DeploymentStatus }).(pulumi.StringPtrOutput)
}

func (o NetworkManagerDeploymentStatusResponseOutput) DeploymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkManagerDeploymentStatusResponse) *string { return v.DeploymentType }).(pulumi.StringPtrOutput)
}

func (o NetworkManagerDeploymentStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkManagerDeploymentStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o NetworkManagerDeploymentStatusResponseOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkManagerDeploymentStatusResponse) *string { return v.Region }).(pulumi.StringPtrOutput)
}

type NetworkManagerDeploymentStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkManagerDeploymentStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkManagerDeploymentStatusResponse)(nil)).Elem()
}

func (o NetworkManagerDeploymentStatusResponseArrayOutput) ToNetworkManagerDeploymentStatusResponseArrayOutput() NetworkManagerDeploymentStatusResponseArrayOutput {
	return o
}

func (o NetworkManagerDeploymentStatusResponseArrayOutput) ToNetworkManagerDeploymentStatusResponseArrayOutputWithContext(ctx context.Context) NetworkManagerDeploymentStatusResponseArrayOutput {
	return o
}

func (o NetworkManagerDeploymentStatusResponseArrayOutput) Index(i pulumi.IntInput) NetworkManagerDeploymentStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkManagerDeploymentStatusResponse {
		return vs[0].([]NetworkManagerDeploymentStatusResponse)[vs[1].(int)]
	}).(NetworkManagerDeploymentStatusResponseOutput)
}

type NetworkManagerPropertiesNetworkManagerScopes struct {
	ManagementGroups []string `pulumi:"managementGroups"`
	Subscriptions    []string `pulumi:"subscriptions"`
}





type NetworkManagerPropertiesNetworkManagerScopesInput interface {
	pulumi.Input

	ToNetworkManagerPropertiesNetworkManagerScopesOutput() NetworkManagerPropertiesNetworkManagerScopesOutput
	ToNetworkManagerPropertiesNetworkManagerScopesOutputWithContext(context.Context) NetworkManagerPropertiesNetworkManagerScopesOutput
}

type NetworkManagerPropertiesNetworkManagerScopesArgs struct {
	ManagementGroups pulumi.StringArrayInput `pulumi:"managementGroups"`
	Subscriptions    pulumi.StringArrayInput `pulumi:"subscriptions"`
}

func (NetworkManagerPropertiesNetworkManagerScopesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkManagerPropertiesNetworkManagerScopes)(nil)).Elem()
}

func (i NetworkManagerPropertiesNetworkManagerScopesArgs) ToNetworkManagerPropertiesNetworkManagerScopesOutput() NetworkManagerPropertiesNetworkManagerScopesOutput {
	return i.ToNetworkManagerPropertiesNetworkManagerScopesOutputWithContext(context.Background())
}

func (i NetworkManagerPropertiesNetworkManagerScopesArgs) ToNetworkManagerPropertiesNetworkManagerScopesOutputWithContext(ctx context.Context) NetworkManagerPropertiesNetworkManagerScopesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerPropertiesNetworkManagerScopesOutput)
}

type NetworkManagerPropertiesNetworkManagerScopesOutput struct{ *pulumi.OutputState }

func (NetworkManagerPropertiesNetworkManagerScopesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkManagerPropertiesNetworkManagerScopes)(nil)).Elem()
}

func (o NetworkManagerPropertiesNetworkManagerScopesOutput) ToNetworkManagerPropertiesNetworkManagerScopesOutput() NetworkManagerPropertiesNetworkManagerScopesOutput {
	return o
}

func (o NetworkManagerPropertiesNetworkManagerScopesOutput) ToNetworkManagerPropertiesNetworkManagerScopesOutputWithContext(ctx context.Context) NetworkManagerPropertiesNetworkManagerScopesOutput {
	return o
}

func (o NetworkManagerPropertiesNetworkManagerScopesOutput) ManagementGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkManagerPropertiesNetworkManagerScopes) []string { return v.ManagementGroups }).(pulumi.StringArrayOutput)
}

func (o NetworkManagerPropertiesNetworkManagerScopesOutput) Subscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkManagerPropertiesNetworkManagerScopes) []string { return v.Subscriptions }).(pulumi.StringArrayOutput)
}

type NetworkManagerPropertiesResponseNetworkManagerScopes struct {
	CrossTenantScopes []CrossTenantScopesResponse `pulumi:"crossTenantScopes"`
	ManagementGroups  []string                    `pulumi:"managementGroups"`
	Subscriptions     []string                    `pulumi:"subscriptions"`
}

type NetworkManagerPropertiesResponseNetworkManagerScopesOutput struct{ *pulumi.OutputState }

func (NetworkManagerPropertiesResponseNetworkManagerScopesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkManagerPropertiesResponseNetworkManagerScopes)(nil)).Elem()
}

func (o NetworkManagerPropertiesResponseNetworkManagerScopesOutput) ToNetworkManagerPropertiesResponseNetworkManagerScopesOutput() NetworkManagerPropertiesResponseNetworkManagerScopesOutput {
	return o
}

func (o NetworkManagerPropertiesResponseNetworkManagerScopesOutput) ToNetworkManagerPropertiesResponseNetworkManagerScopesOutputWithContext(ctx context.Context) NetworkManagerPropertiesResponseNetworkManagerScopesOutput {
	return o
}

func (o NetworkManagerPropertiesResponseNetworkManagerScopesOutput) CrossTenantScopes() CrossTenantScopesResponseArrayOutput {
	return o.ApplyT(func(v NetworkManagerPropertiesResponseNetworkManagerScopes) []CrossTenantScopesResponse {
		return v.CrossTenantScopes
	}).(CrossTenantScopesResponseArrayOutput)
}

func (o NetworkManagerPropertiesResponseNetworkManagerScopesOutput) ManagementGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkManagerPropertiesResponseNetworkManagerScopes) []string { return v.ManagementGroups }).(pulumi.StringArrayOutput)
}

func (o NetworkManagerPropertiesResponseNetworkManagerScopesOutput) Subscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkManagerPropertiesResponseNetworkManagerScopes) []string { return v.Subscriptions }).(pulumi.StringArrayOutput)
}

type NetworkManagerSecurityGroupItem struct {
	NetworkGroupId string `pulumi:"networkGroupId"`
}





type NetworkManagerSecurityGroupItemInput interface {
	pulumi.Input

	ToNetworkManagerSecurityGroupItemOutput() NetworkManagerSecurityGroupItemOutput
	ToNetworkManagerSecurityGroupItemOutputWithContext(context.Context) NetworkManagerSecurityGroupItemOutput
}

type NetworkManagerSecurityGroupItemArgs struct {
	NetworkGroupId pulumi.StringInput `pulumi:"networkGroupId"`
}

func (NetworkManagerSecurityGroupItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkManagerSecurityGroupItem)(nil)).Elem()
}

func (i NetworkManagerSecurityGroupItemArgs) ToNetworkManagerSecurityGroupItemOutput() NetworkManagerSecurityGroupItemOutput {
	return i.ToNetworkManagerSecurityGroupItemOutputWithContext(context.Background())
}

func (i NetworkManagerSecurityGroupItemArgs) ToNetworkManagerSecurityGroupItemOutputWithContext(ctx context.Context) NetworkManagerSecurityGroupItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerSecurityGroupItemOutput)
}





type NetworkManagerSecurityGroupItemArrayInput interface {
	pulumi.Input

	ToNetworkManagerSecurityGroupItemArrayOutput() NetworkManagerSecurityGroupItemArrayOutput
	ToNetworkManagerSecurityGroupItemArrayOutputWithContext(context.Context) NetworkManagerSecurityGroupItemArrayOutput
}

type NetworkManagerSecurityGroupItemArray []NetworkManagerSecurityGroupItemInput

func (NetworkManagerSecurityGroupItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkManagerSecurityGroupItem)(nil)).Elem()
}

func (i NetworkManagerSecurityGroupItemArray) ToNetworkManagerSecurityGroupItemArrayOutput() NetworkManagerSecurityGroupItemArrayOutput {
	return i.ToNetworkManagerSecurityGroupItemArrayOutputWithContext(context.Background())
}

func (i NetworkManagerSecurityGroupItemArray) ToNetworkManagerSecurityGroupItemArrayOutputWithContext(ctx context.Context) NetworkManagerSecurityGroupItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerSecurityGroupItemArrayOutput)
}

type NetworkManagerSecurityGroupItemOutput struct{ *pulumi.OutputState }

func (NetworkManagerSecurityGroupItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkManagerSecurityGroupItem)(nil)).Elem()
}

func (o NetworkManagerSecurityGroupItemOutput) ToNetworkManagerSecurityGroupItemOutput() NetworkManagerSecurityGroupItemOutput {
	return o
}

func (o NetworkManagerSecurityGroupItemOutput) ToNetworkManagerSecurityGroupItemOutputWithContext(ctx context.Context) NetworkManagerSecurityGroupItemOutput {
	return o
}

func (o NetworkManagerSecurityGroupItemOutput) NetworkGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkManagerSecurityGroupItem) string { return v.NetworkGroupId }).(pulumi.StringOutput)
}

type NetworkManagerSecurityGroupItemArrayOutput struct{ *pulumi.OutputState }

func (NetworkManagerSecurityGroupItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkManagerSecurityGroupItem)(nil)).Elem()
}

func (o NetworkManagerSecurityGroupItemArrayOutput) ToNetworkManagerSecurityGroupItemArrayOutput() NetworkManagerSecurityGroupItemArrayOutput {
	return o
}

func (o NetworkManagerSecurityGroupItemArrayOutput) ToNetworkManagerSecurityGroupItemArrayOutputWithContext(ctx context.Context) NetworkManagerSecurityGroupItemArrayOutput {
	return o
}

func (o NetworkManagerSecurityGroupItemArrayOutput) Index(i pulumi.IntInput) NetworkManagerSecurityGroupItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkManagerSecurityGroupItem {
		return vs[0].([]NetworkManagerSecurityGroupItem)[vs[1].(int)]
	}).(NetworkManagerSecurityGroupItemOutput)
}

type NetworkManagerSecurityGroupItemResponse struct {
	NetworkGroupId string `pulumi:"networkGroupId"`
}

type NetworkManagerSecurityGroupItemResponseOutput struct{ *pulumi.OutputState }

func (NetworkManagerSecurityGroupItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkManagerSecurityGroupItemResponse)(nil)).Elem()
}

func (o NetworkManagerSecurityGroupItemResponseOutput) ToNetworkManagerSecurityGroupItemResponseOutput() NetworkManagerSecurityGroupItemResponseOutput {
	return o
}

func (o NetworkManagerSecurityGroupItemResponseOutput) ToNetworkManagerSecurityGroupItemResponseOutputWithContext(ctx context.Context) NetworkManagerSecurityGroupItemResponseOutput {
	return o
}

func (o NetworkManagerSecurityGroupItemResponseOutput) NetworkGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkManagerSecurityGroupItemResponse) string { return v.NetworkGroupId }).(pulumi.StringOutput)
}

type NetworkManagerSecurityGroupItemResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkManagerSecurityGroupItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkManagerSecurityGroupItemResponse)(nil)).Elem()
}

func (o NetworkManagerSecurityGroupItemResponseArrayOutput) ToNetworkManagerSecurityGroupItemResponseArrayOutput() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o
}

func (o NetworkManagerSecurityGroupItemResponseArrayOutput) ToNetworkManagerSecurityGroupItemResponseArrayOutputWithContext(ctx context.Context) NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o
}

func (o NetworkManagerSecurityGroupItemResponseArrayOutput) Index(i pulumi.IntInput) NetworkManagerSecurityGroupItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkManagerSecurityGroupItemResponse {
		return vs[0].([]NetworkManagerSecurityGroupItemResponse)[vs[1].(int)]
	}).(NetworkManagerSecurityGroupItemResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActiveConnectivityConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ActiveConnectivityConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(AddressPrefixItemResponseOutput{})
	pulumi.RegisterOutputType(AddressPrefixItemResponseArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationGroupResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(ConnectivityGroupItemOutput{})
	pulumi.RegisterOutputType(ConnectivityGroupItemArrayOutput{})
	pulumi.RegisterOutputType(ConnectivityGroupItemResponseOutput{})
	pulumi.RegisterOutputType(ConnectivityGroupItemResponseArrayOutput{})
	pulumi.RegisterOutputType(CrossTenantScopesResponseOutput{})
	pulumi.RegisterOutputType(CrossTenantScopesResponseArrayOutput{})
	pulumi.RegisterOutputType(DnsConfigOutput{})
	pulumi.RegisterOutputType(DnsConfigPtrOutput{})
	pulumi.RegisterOutputType(DnsConfigResponseOutput{})
	pulumi.RegisterOutputType(DnsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(EffectiveConnectivityConfigurationResponseOutput{})
	pulumi.RegisterOutputType(EffectiveConnectivityConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(EffectiveVirtualNetworkResponseOutput{})
	pulumi.RegisterOutputType(EffectiveVirtualNetworkResponseArrayOutput{})
	pulumi.RegisterOutputType(EndpointTypeOutput{})
	pulumi.RegisterOutputType(EndpointTypeArrayOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesCustomHeadersOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesCustomHeadersArrayOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesResponseCustomHeadersOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesResponseCustomHeadersArrayOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesResponseSubnetsOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesResponseSubnetsArrayOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesSubnetsOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesSubnetsArrayOutput{})
	pulumi.RegisterOutputType(EndpointResponseOutput{})
	pulumi.RegisterOutputType(EndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(HubOutput{})
	pulumi.RegisterOutputType(HubArrayOutput{})
	pulumi.RegisterOutputType(HubResponseOutput{})
	pulumi.RegisterOutputType(HubResponseArrayOutput{})
	pulumi.RegisterOutputType(MonitorConfigOutput{})
	pulumi.RegisterOutputType(MonitorConfigPtrOutput{})
	pulumi.RegisterOutputType(MonitorConfigCustomHeadersOutput{})
	pulumi.RegisterOutputType(MonitorConfigCustomHeadersArrayOutput{})
	pulumi.RegisterOutputType(MonitorConfigExpectedStatusCodeRangesOutput{})
	pulumi.RegisterOutputType(MonitorConfigExpectedStatusCodeRangesArrayOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponseOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponseCustomHeadersOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponseCustomHeadersArrayOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponseExpectedStatusCodeRangesOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponseExpectedStatusCodeRangesArrayOutput{})
	pulumi.RegisterOutputType(NetworkManagerDeploymentStatusResponseOutput{})
	pulumi.RegisterOutputType(NetworkManagerDeploymentStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkManagerPropertiesNetworkManagerScopesOutput{})
	pulumi.RegisterOutputType(NetworkManagerPropertiesResponseNetworkManagerScopesOutput{})
	pulumi.RegisterOutputType(NetworkManagerSecurityGroupItemOutput{})
	pulumi.RegisterOutputType(NetworkManagerSecurityGroupItemArrayOutput{})
	pulumi.RegisterOutputType(NetworkManagerSecurityGroupItemResponseOutput{})
	pulumi.RegisterOutputType(NetworkManagerSecurityGroupItemResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

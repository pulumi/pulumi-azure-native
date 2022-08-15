


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
	pulumi.RegisterOutputType(EffectiveConnectivityConfigurationResponseOutput{})
	pulumi.RegisterOutputType(EffectiveConnectivityConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(EffectiveVirtualNetworkResponseOutput{})
	pulumi.RegisterOutputType(EffectiveVirtualNetworkResponseArrayOutput{})
	pulumi.RegisterOutputType(HubOutput{})
	pulumi.RegisterOutputType(HubArrayOutput{})
	pulumi.RegisterOutputType(HubResponseOutput{})
	pulumi.RegisterOutputType(HubResponseArrayOutput{})
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

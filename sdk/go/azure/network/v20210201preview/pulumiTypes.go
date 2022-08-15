


package v20210201preview

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
	DisplayName           *string                         `pulumi:"displayName"`
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

func (o ActiveConnectivityConfigurationResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveConnectivityConfigurationResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
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
	ConfigurationDisplayName      *string                                   `pulumi:"configurationDisplayName"`
	Description                   string                                    `pulumi:"description"`
	DestinationPortRanges         []string                                  `pulumi:"destinationPortRanges"`
	Destinations                  []AddressPrefixItemResponse               `pulumi:"destinations"`
	Direction                     string                                    `pulumi:"direction"`
	DisplayName                   string                                    `pulumi:"displayName"`
	Flag                          *string                                   `pulumi:"flag"`
	Id                            *string                                   `pulumi:"id"`
	Kind                          string                                    `pulumi:"kind"`
	Priority                      int                                       `pulumi:"priority"`
	Protocol                      string                                    `pulumi:"protocol"`
	ProvisioningState             string                                    `pulumi:"provisioningState"`
	Region                        *string                                   `pulumi:"region"`
	RuleCollectionAppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     *string                                   `pulumi:"ruleCollectionDescription"`
	RuleCollectionDisplayName     *string                                   `pulumi:"ruleCollectionDisplayName"`
	RuleGroups                    []ConfigurationGroupResponse              `pulumi:"ruleGroups"`
	SourcePortRanges              []string                                  `pulumi:"sourcePortRanges"`
	Sources                       []AddressPrefixItemResponse               `pulumi:"sources"`
}

type ActiveDefaultSecurityUserRuleResponse struct {
	CommitTime                    *string                                   `pulumi:"commitTime"`
	ConfigurationDescription      *string                                   `pulumi:"configurationDescription"`
	ConfigurationDisplayName      *string                                   `pulumi:"configurationDisplayName"`
	Description                   string                                    `pulumi:"description"`
	DestinationPortRanges         []string                                  `pulumi:"destinationPortRanges"`
	Destinations                  []AddressPrefixItemResponse               `pulumi:"destinations"`
	Direction                     string                                    `pulumi:"direction"`
	DisplayName                   string                                    `pulumi:"displayName"`
	Flag                          *string                                   `pulumi:"flag"`
	Id                            *string                                   `pulumi:"id"`
	Kind                          string                                    `pulumi:"kind"`
	Protocol                      string                                    `pulumi:"protocol"`
	ProvisioningState             string                                    `pulumi:"provisioningState"`
	Region                        *string                                   `pulumi:"region"`
	RuleCollectionAppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     *string                                   `pulumi:"ruleCollectionDescription"`
	RuleCollectionDisplayName     *string                                   `pulumi:"ruleCollectionDisplayName"`
	RuleGroups                    []ConfigurationGroupResponse              `pulumi:"ruleGroups"`
	SourcePortRanges              []string                                  `pulumi:"sourcePortRanges"`
	Sources                       []AddressPrefixItemResponse               `pulumi:"sources"`
}

type ActiveSecurityAdminRuleResponse struct {
	Access                        string                                    `pulumi:"access"`
	CommitTime                    *string                                   `pulumi:"commitTime"`
	ConfigurationDescription      *string                                   `pulumi:"configurationDescription"`
	ConfigurationDisplayName      *string                                   `pulumi:"configurationDisplayName"`
	Description                   *string                                   `pulumi:"description"`
	DestinationPortRanges         []string                                  `pulumi:"destinationPortRanges"`
	Destinations                  []AddressPrefixItemResponse               `pulumi:"destinations"`
	Direction                     string                                    `pulumi:"direction"`
	DisplayName                   *string                                   `pulumi:"displayName"`
	Id                            *string                                   `pulumi:"id"`
	Kind                          string                                    `pulumi:"kind"`
	Priority                      *int                                      `pulumi:"priority"`
	Protocol                      string                                    `pulumi:"protocol"`
	ProvisioningState             string                                    `pulumi:"provisioningState"`
	Region                        *string                                   `pulumi:"region"`
	RuleCollectionAppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     *string                                   `pulumi:"ruleCollectionDescription"`
	RuleCollectionDisplayName     *string                                   `pulumi:"ruleCollectionDisplayName"`
	RuleGroups                    []ConfigurationGroupResponse              `pulumi:"ruleGroups"`
	SourcePortRanges              []string                                  `pulumi:"sourcePortRanges"`
	Sources                       []AddressPrefixItemResponse               `pulumi:"sources"`
}

type ActiveSecurityUserRuleResponse struct {
	CommitTime                    *string                                   `pulumi:"commitTime"`
	ConfigurationDescription      *string                                   `pulumi:"configurationDescription"`
	ConfigurationDisplayName      *string                                   `pulumi:"configurationDisplayName"`
	Description                   *string                                   `pulumi:"description"`
	DestinationPortRanges         []string                                  `pulumi:"destinationPortRanges"`
	Destinations                  []AddressPrefixItemResponse               `pulumi:"destinations"`
	Direction                     string                                    `pulumi:"direction"`
	DisplayName                   *string                                   `pulumi:"displayName"`
	Id                            *string                                   `pulumi:"id"`
	Kind                          string                                    `pulumi:"kind"`
	Protocol                      string                                    `pulumi:"protocol"`
	ProvisioningState             string                                    `pulumi:"provisioningState"`
	Region                        *string                                   `pulumi:"region"`
	RuleCollectionAppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     *string                                   `pulumi:"ruleCollectionDescription"`
	RuleCollectionDisplayName     *string                                   `pulumi:"ruleCollectionDisplayName"`
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
	ConditionalMembership *string                    `pulumi:"conditionalMembership"`
	Description           *string                    `pulumi:"description"`
	DisplayName           *string                    `pulumi:"displayName"`
	GroupMembers          []GroupMembersItemResponse `pulumi:"groupMembers"`
	Id                    *string                    `pulumi:"id"`
	MemberType            *string                    `pulumi:"memberType"`
	ProvisioningState     string                     `pulumi:"provisioningState"`
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

func (o ConfigurationGroupResponseOutput) ConditionalMembership() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationGroupResponse) *string { return v.ConditionalMembership }).(pulumi.StringPtrOutput)
}

func (o ConfigurationGroupResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationGroupResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConfigurationGroupResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationGroupResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ConfigurationGroupResponseOutput) GroupMembers() GroupMembersItemResponseArrayOutput {
	return o.ApplyT(func(v ConfigurationGroupResponse) []GroupMembersItemResponse { return v.GroupMembers }).(GroupMembersItemResponseArrayOutput)
}

func (o ConfigurationGroupResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationGroupResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ConfigurationGroupResponseOutput) MemberType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationGroupResponse) *string { return v.MemberType }).(pulumi.StringPtrOutput)
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
	GroupConnectivity *string `pulumi:"groupConnectivity"`
	IsGlobal          *string `pulumi:"isGlobal"`
	NetworkGroupId    *string `pulumi:"networkGroupId"`
	UseHubGateway     *string `pulumi:"useHubGateway"`
}





type ConnectivityGroupItemInput interface {
	pulumi.Input

	ToConnectivityGroupItemOutput() ConnectivityGroupItemOutput
	ToConnectivityGroupItemOutputWithContext(context.Context) ConnectivityGroupItemOutput
}

type ConnectivityGroupItemArgs struct {
	GroupConnectivity pulumi.StringPtrInput `pulumi:"groupConnectivity"`
	IsGlobal          pulumi.StringPtrInput `pulumi:"isGlobal"`
	NetworkGroupId    pulumi.StringPtrInput `pulumi:"networkGroupId"`
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

func (o ConnectivityGroupItemOutput) GroupConnectivity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityGroupItem) *string { return v.GroupConnectivity }).(pulumi.StringPtrOutput)
}

func (o ConnectivityGroupItemOutput) IsGlobal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityGroupItem) *string { return v.IsGlobal }).(pulumi.StringPtrOutput)
}

func (o ConnectivityGroupItemOutput) NetworkGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityGroupItem) *string { return v.NetworkGroupId }).(pulumi.StringPtrOutput)
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
	GroupConnectivity *string `pulumi:"groupConnectivity"`
	IsGlobal          *string `pulumi:"isGlobal"`
	NetworkGroupId    *string `pulumi:"networkGroupId"`
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

func (o ConnectivityGroupItemResponseOutput) GroupConnectivity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityGroupItemResponse) *string { return v.GroupConnectivity }).(pulumi.StringPtrOutput)
}

func (o ConnectivityGroupItemResponseOutput) IsGlobal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityGroupItemResponse) *string { return v.IsGlobal }).(pulumi.StringPtrOutput)
}

func (o ConnectivityGroupItemResponseOutput) NetworkGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityGroupItemResponse) *string { return v.NetworkGroupId }).(pulumi.StringPtrOutput)
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

type EffectiveConnectivityConfigurationResponse struct {
	AppliesToGroups       []ConnectivityGroupItemResponse `pulumi:"appliesToGroups"`
	ConfigurationGroups   []ConfigurationGroupResponse    `pulumi:"configurationGroups"`
	ConnectivityTopology  string                          `pulumi:"connectivityTopology"`
	DeleteExistingPeering *string                         `pulumi:"deleteExistingPeering"`
	Description           *string                         `pulumi:"description"`
	DisplayName           *string                         `pulumi:"displayName"`
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

func (o EffectiveConnectivityConfigurationResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveConnectivityConfigurationResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
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
	ConfigurationDisplayName      *string                                   `pulumi:"configurationDisplayName"`
	Description                   string                                    `pulumi:"description"`
	DestinationPortRanges         []string                                  `pulumi:"destinationPortRanges"`
	Destinations                  []AddressPrefixItemResponse               `pulumi:"destinations"`
	Direction                     string                                    `pulumi:"direction"`
	DisplayName                   string                                    `pulumi:"displayName"`
	Flag                          *string                                   `pulumi:"flag"`
	Id                            *string                                   `pulumi:"id"`
	Kind                          string                                    `pulumi:"kind"`
	Priority                      int                                       `pulumi:"priority"`
	Protocol                      string                                    `pulumi:"protocol"`
	ProvisioningState             string                                    `pulumi:"provisioningState"`
	RuleCollectionAppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     *string                                   `pulumi:"ruleCollectionDescription"`
	RuleCollectionDisplayName     *string                                   `pulumi:"ruleCollectionDisplayName"`
	RuleGroups                    []ConfigurationGroupResponse              `pulumi:"ruleGroups"`
	SourcePortRanges              []string                                  `pulumi:"sourcePortRanges"`
	Sources                       []AddressPrefixItemResponse               `pulumi:"sources"`
}

type EffectiveSecurityAdminRuleResponse struct {
	Access                        string                                    `pulumi:"access"`
	ConfigurationDescription      *string                                   `pulumi:"configurationDescription"`
	ConfigurationDisplayName      *string                                   `pulumi:"configurationDisplayName"`
	Description                   *string                                   `pulumi:"description"`
	DestinationPortRanges         []string                                  `pulumi:"destinationPortRanges"`
	Destinations                  []AddressPrefixItemResponse               `pulumi:"destinations"`
	Direction                     string                                    `pulumi:"direction"`
	DisplayName                   *string                                   `pulumi:"displayName"`
	Id                            *string                                   `pulumi:"id"`
	Kind                          string                                    `pulumi:"kind"`
	Priority                      *int                                      `pulumi:"priority"`
	Protocol                      string                                    `pulumi:"protocol"`
	ProvisioningState             string                                    `pulumi:"provisioningState"`
	RuleCollectionAppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     *string                                   `pulumi:"ruleCollectionDescription"`
	RuleCollectionDisplayName     *string                                   `pulumi:"ruleCollectionDisplayName"`
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

type GroupMembersItem struct {
	ResourceId *string `pulumi:"resourceId"`
}





type GroupMembersItemInput interface {
	pulumi.Input

	ToGroupMembersItemOutput() GroupMembersItemOutput
	ToGroupMembersItemOutputWithContext(context.Context) GroupMembersItemOutput
}

type GroupMembersItemArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (GroupMembersItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembersItem)(nil)).Elem()
}

func (i GroupMembersItemArgs) ToGroupMembersItemOutput() GroupMembersItemOutput {
	return i.ToGroupMembersItemOutputWithContext(context.Background())
}

func (i GroupMembersItemArgs) ToGroupMembersItemOutputWithContext(ctx context.Context) GroupMembersItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembersItemOutput)
}





type GroupMembersItemArrayInput interface {
	pulumi.Input

	ToGroupMembersItemArrayOutput() GroupMembersItemArrayOutput
	ToGroupMembersItemArrayOutputWithContext(context.Context) GroupMembersItemArrayOutput
}

type GroupMembersItemArray []GroupMembersItemInput

func (GroupMembersItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupMembersItem)(nil)).Elem()
}

func (i GroupMembersItemArray) ToGroupMembersItemArrayOutput() GroupMembersItemArrayOutput {
	return i.ToGroupMembersItemArrayOutputWithContext(context.Background())
}

func (i GroupMembersItemArray) ToGroupMembersItemArrayOutputWithContext(ctx context.Context) GroupMembersItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembersItemArrayOutput)
}

type GroupMembersItemOutput struct{ *pulumi.OutputState }

func (GroupMembersItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembersItem)(nil)).Elem()
}

func (o GroupMembersItemOutput) ToGroupMembersItemOutput() GroupMembersItemOutput {
	return o
}

func (o GroupMembersItemOutput) ToGroupMembersItemOutputWithContext(ctx context.Context) GroupMembersItemOutput {
	return o
}

func (o GroupMembersItemOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupMembersItem) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type GroupMembersItemArrayOutput struct{ *pulumi.OutputState }

func (GroupMembersItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupMembersItem)(nil)).Elem()
}

func (o GroupMembersItemArrayOutput) ToGroupMembersItemArrayOutput() GroupMembersItemArrayOutput {
	return o
}

func (o GroupMembersItemArrayOutput) ToGroupMembersItemArrayOutputWithContext(ctx context.Context) GroupMembersItemArrayOutput {
	return o
}

func (o GroupMembersItemArrayOutput) Index(i pulumi.IntInput) GroupMembersItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupMembersItem {
		return vs[0].([]GroupMembersItem)[vs[1].(int)]
	}).(GroupMembersItemOutput)
}

type GroupMembersItemResponse struct {
	ResourceId *string `pulumi:"resourceId"`
}

type GroupMembersItemResponseOutput struct{ *pulumi.OutputState }

func (GroupMembersItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembersItemResponse)(nil)).Elem()
}

func (o GroupMembersItemResponseOutput) ToGroupMembersItemResponseOutput() GroupMembersItemResponseOutput {
	return o
}

func (o GroupMembersItemResponseOutput) ToGroupMembersItemResponseOutputWithContext(ctx context.Context) GroupMembersItemResponseOutput {
	return o
}

func (o GroupMembersItemResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupMembersItemResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type GroupMembersItemResponseArrayOutput struct{ *pulumi.OutputState }

func (GroupMembersItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupMembersItemResponse)(nil)).Elem()
}

func (o GroupMembersItemResponseArrayOutput) ToGroupMembersItemResponseArrayOutput() GroupMembersItemResponseArrayOutput {
	return o
}

func (o GroupMembersItemResponseArrayOutput) ToGroupMembersItemResponseArrayOutputWithContext(ctx context.Context) GroupMembersItemResponseArrayOutput {
	return o
}

func (o GroupMembersItemResponseArrayOutput) Index(i pulumi.IntInput) GroupMembersItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupMembersItemResponse {
		return vs[0].([]GroupMembersItemResponse)[vs[1].(int)]
	}).(GroupMembersItemResponseOutput)
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

func (i NetworkManagerPropertiesNetworkManagerScopesArgs) ToNetworkManagerPropertiesNetworkManagerScopesPtrOutput() NetworkManagerPropertiesNetworkManagerScopesPtrOutput {
	return i.ToNetworkManagerPropertiesNetworkManagerScopesPtrOutputWithContext(context.Background())
}

func (i NetworkManagerPropertiesNetworkManagerScopesArgs) ToNetworkManagerPropertiesNetworkManagerScopesPtrOutputWithContext(ctx context.Context) NetworkManagerPropertiesNetworkManagerScopesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerPropertiesNetworkManagerScopesOutput).ToNetworkManagerPropertiesNetworkManagerScopesPtrOutputWithContext(ctx)
}









type NetworkManagerPropertiesNetworkManagerScopesPtrInput interface {
	pulumi.Input

	ToNetworkManagerPropertiesNetworkManagerScopesPtrOutput() NetworkManagerPropertiesNetworkManagerScopesPtrOutput
	ToNetworkManagerPropertiesNetworkManagerScopesPtrOutputWithContext(context.Context) NetworkManagerPropertiesNetworkManagerScopesPtrOutput
}

type networkManagerPropertiesNetworkManagerScopesPtrType NetworkManagerPropertiesNetworkManagerScopesArgs

func NetworkManagerPropertiesNetworkManagerScopesPtr(v *NetworkManagerPropertiesNetworkManagerScopesArgs) NetworkManagerPropertiesNetworkManagerScopesPtrInput {
	return (*networkManagerPropertiesNetworkManagerScopesPtrType)(v)
}

func (*networkManagerPropertiesNetworkManagerScopesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkManagerPropertiesNetworkManagerScopes)(nil)).Elem()
}

func (i *networkManagerPropertiesNetworkManagerScopesPtrType) ToNetworkManagerPropertiesNetworkManagerScopesPtrOutput() NetworkManagerPropertiesNetworkManagerScopesPtrOutput {
	return i.ToNetworkManagerPropertiesNetworkManagerScopesPtrOutputWithContext(context.Background())
}

func (i *networkManagerPropertiesNetworkManagerScopesPtrType) ToNetworkManagerPropertiesNetworkManagerScopesPtrOutputWithContext(ctx context.Context) NetworkManagerPropertiesNetworkManagerScopesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerPropertiesNetworkManagerScopesPtrOutput)
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

func (o NetworkManagerPropertiesNetworkManagerScopesOutput) ToNetworkManagerPropertiesNetworkManagerScopesPtrOutput() NetworkManagerPropertiesNetworkManagerScopesPtrOutput {
	return o.ToNetworkManagerPropertiesNetworkManagerScopesPtrOutputWithContext(context.Background())
}

func (o NetworkManagerPropertiesNetworkManagerScopesOutput) ToNetworkManagerPropertiesNetworkManagerScopesPtrOutputWithContext(ctx context.Context) NetworkManagerPropertiesNetworkManagerScopesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkManagerPropertiesNetworkManagerScopes) *NetworkManagerPropertiesNetworkManagerScopes {
		return &v
	}).(NetworkManagerPropertiesNetworkManagerScopesPtrOutput)
}

func (o NetworkManagerPropertiesNetworkManagerScopesOutput) ManagementGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkManagerPropertiesNetworkManagerScopes) []string { return v.ManagementGroups }).(pulumi.StringArrayOutput)
}

func (o NetworkManagerPropertiesNetworkManagerScopesOutput) Subscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkManagerPropertiesNetworkManagerScopes) []string { return v.Subscriptions }).(pulumi.StringArrayOutput)
}

type NetworkManagerPropertiesNetworkManagerScopesPtrOutput struct{ *pulumi.OutputState }

func (NetworkManagerPropertiesNetworkManagerScopesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkManagerPropertiesNetworkManagerScopes)(nil)).Elem()
}

func (o NetworkManagerPropertiesNetworkManagerScopesPtrOutput) ToNetworkManagerPropertiesNetworkManagerScopesPtrOutput() NetworkManagerPropertiesNetworkManagerScopesPtrOutput {
	return o
}

func (o NetworkManagerPropertiesNetworkManagerScopesPtrOutput) ToNetworkManagerPropertiesNetworkManagerScopesPtrOutputWithContext(ctx context.Context) NetworkManagerPropertiesNetworkManagerScopesPtrOutput {
	return o
}

func (o NetworkManagerPropertiesNetworkManagerScopesPtrOutput) Elem() NetworkManagerPropertiesNetworkManagerScopesOutput {
	return o.ApplyT(func(v *NetworkManagerPropertiesNetworkManagerScopes) NetworkManagerPropertiesNetworkManagerScopes {
		if v != nil {
			return *v
		}
		var ret NetworkManagerPropertiesNetworkManagerScopes
		return ret
	}).(NetworkManagerPropertiesNetworkManagerScopesOutput)
}

func (o NetworkManagerPropertiesNetworkManagerScopesPtrOutput) ManagementGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkManagerPropertiesNetworkManagerScopes) []string {
		if v == nil {
			return nil
		}
		return v.ManagementGroups
	}).(pulumi.StringArrayOutput)
}

func (o NetworkManagerPropertiesNetworkManagerScopesPtrOutput) Subscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkManagerPropertiesNetworkManagerScopes) []string {
		if v == nil {
			return nil
		}
		return v.Subscriptions
	}).(pulumi.StringArrayOutput)
}

type NetworkManagerPropertiesResponseNetworkManagerScopes struct {
	ManagementGroups []string `pulumi:"managementGroups"`
	Subscriptions    []string `pulumi:"subscriptions"`
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

func (o NetworkManagerPropertiesResponseNetworkManagerScopesOutput) ManagementGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkManagerPropertiesResponseNetworkManagerScopes) []string { return v.ManagementGroups }).(pulumi.StringArrayOutput)
}

func (o NetworkManagerPropertiesResponseNetworkManagerScopesOutput) Subscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkManagerPropertiesResponseNetworkManagerScopes) []string { return v.Subscriptions }).(pulumi.StringArrayOutput)
}

type NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput struct{ *pulumi.OutputState }

func (NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkManagerPropertiesResponseNetworkManagerScopes)(nil)).Elem()
}

func (o NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput) ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput() NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput {
	return o
}

func (o NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput) ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutputWithContext(ctx context.Context) NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput {
	return o
}

func (o NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput) Elem() NetworkManagerPropertiesResponseNetworkManagerScopesOutput {
	return o.ApplyT(func(v *NetworkManagerPropertiesResponseNetworkManagerScopes) NetworkManagerPropertiesResponseNetworkManagerScopes {
		if v != nil {
			return *v
		}
		var ret NetworkManagerPropertiesResponseNetworkManagerScopes
		return ret
	}).(NetworkManagerPropertiesResponseNetworkManagerScopesOutput)
}

func (o NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput) ManagementGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkManagerPropertiesResponseNetworkManagerScopes) []string {
		if v == nil {
			return nil
		}
		return v.ManagementGroups
	}).(pulumi.StringArrayOutput)
}

func (o NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput) Subscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkManagerPropertiesResponseNetworkManagerScopes) []string {
		if v == nil {
			return nil
		}
		return v.Subscriptions
	}).(pulumi.StringArrayOutput)
}

type NetworkManagerSecurityGroupItem struct {
	NetworkGroupId *string `pulumi:"networkGroupId"`
}





type NetworkManagerSecurityGroupItemInput interface {
	pulumi.Input

	ToNetworkManagerSecurityGroupItemOutput() NetworkManagerSecurityGroupItemOutput
	ToNetworkManagerSecurityGroupItemOutputWithContext(context.Context) NetworkManagerSecurityGroupItemOutput
}

type NetworkManagerSecurityGroupItemArgs struct {
	NetworkGroupId pulumi.StringPtrInput `pulumi:"networkGroupId"`
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

func (o NetworkManagerSecurityGroupItemOutput) NetworkGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkManagerSecurityGroupItem) *string { return v.NetworkGroupId }).(pulumi.StringPtrOutput)
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
	NetworkGroupId *string `pulumi:"networkGroupId"`
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

func (o NetworkManagerSecurityGroupItemResponseOutput) NetworkGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkManagerSecurityGroupItemResponse) *string { return v.NetworkGroupId }).(pulumi.StringPtrOutput)
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

type PerimeterBasedAccessRule struct {
	Id *string `pulumi:"id"`
}





type PerimeterBasedAccessRuleInput interface {
	pulumi.Input

	ToPerimeterBasedAccessRuleOutput() PerimeterBasedAccessRuleOutput
	ToPerimeterBasedAccessRuleOutputWithContext(context.Context) PerimeterBasedAccessRuleOutput
}

type PerimeterBasedAccessRuleArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PerimeterBasedAccessRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PerimeterBasedAccessRule)(nil)).Elem()
}

func (i PerimeterBasedAccessRuleArgs) ToPerimeterBasedAccessRuleOutput() PerimeterBasedAccessRuleOutput {
	return i.ToPerimeterBasedAccessRuleOutputWithContext(context.Background())
}

func (i PerimeterBasedAccessRuleArgs) ToPerimeterBasedAccessRuleOutputWithContext(ctx context.Context) PerimeterBasedAccessRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerimeterBasedAccessRuleOutput)
}





type PerimeterBasedAccessRuleArrayInput interface {
	pulumi.Input

	ToPerimeterBasedAccessRuleArrayOutput() PerimeterBasedAccessRuleArrayOutput
	ToPerimeterBasedAccessRuleArrayOutputWithContext(context.Context) PerimeterBasedAccessRuleArrayOutput
}

type PerimeterBasedAccessRuleArray []PerimeterBasedAccessRuleInput

func (PerimeterBasedAccessRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerimeterBasedAccessRule)(nil)).Elem()
}

func (i PerimeterBasedAccessRuleArray) ToPerimeterBasedAccessRuleArrayOutput() PerimeterBasedAccessRuleArrayOutput {
	return i.ToPerimeterBasedAccessRuleArrayOutputWithContext(context.Background())
}

func (i PerimeterBasedAccessRuleArray) ToPerimeterBasedAccessRuleArrayOutputWithContext(ctx context.Context) PerimeterBasedAccessRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerimeterBasedAccessRuleArrayOutput)
}

type PerimeterBasedAccessRuleOutput struct{ *pulumi.OutputState }

func (PerimeterBasedAccessRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerimeterBasedAccessRule)(nil)).Elem()
}

func (o PerimeterBasedAccessRuleOutput) ToPerimeterBasedAccessRuleOutput() PerimeterBasedAccessRuleOutput {
	return o
}

func (o PerimeterBasedAccessRuleOutput) ToPerimeterBasedAccessRuleOutputWithContext(ctx context.Context) PerimeterBasedAccessRuleOutput {
	return o
}

func (o PerimeterBasedAccessRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PerimeterBasedAccessRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PerimeterBasedAccessRuleArrayOutput struct{ *pulumi.OutputState }

func (PerimeterBasedAccessRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerimeterBasedAccessRule)(nil)).Elem()
}

func (o PerimeterBasedAccessRuleArrayOutput) ToPerimeterBasedAccessRuleArrayOutput() PerimeterBasedAccessRuleArrayOutput {
	return o
}

func (o PerimeterBasedAccessRuleArrayOutput) ToPerimeterBasedAccessRuleArrayOutputWithContext(ctx context.Context) PerimeterBasedAccessRuleArrayOutput {
	return o
}

func (o PerimeterBasedAccessRuleArrayOutput) Index(i pulumi.IntInput) PerimeterBasedAccessRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerimeterBasedAccessRule {
		return vs[0].([]PerimeterBasedAccessRule)[vs[1].(int)]
	}).(PerimeterBasedAccessRuleOutput)
}

type PerimeterBasedAccessRuleResponse struct {
	Id            *string `pulumi:"id"`
	Location      string  `pulumi:"location"`
	PerimeterGuid string  `pulumi:"perimeterGuid"`
}

type PerimeterBasedAccessRuleResponseOutput struct{ *pulumi.OutputState }

func (PerimeterBasedAccessRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerimeterBasedAccessRuleResponse)(nil)).Elem()
}

func (o PerimeterBasedAccessRuleResponseOutput) ToPerimeterBasedAccessRuleResponseOutput() PerimeterBasedAccessRuleResponseOutput {
	return o
}

func (o PerimeterBasedAccessRuleResponseOutput) ToPerimeterBasedAccessRuleResponseOutputWithContext(ctx context.Context) PerimeterBasedAccessRuleResponseOutput {
	return o
}

func (o PerimeterBasedAccessRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PerimeterBasedAccessRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PerimeterBasedAccessRuleResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v PerimeterBasedAccessRuleResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o PerimeterBasedAccessRuleResponseOutput) PerimeterGuid() pulumi.StringOutput {
	return o.ApplyT(func(v PerimeterBasedAccessRuleResponse) string { return v.PerimeterGuid }).(pulumi.StringOutput)
}

type PerimeterBasedAccessRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (PerimeterBasedAccessRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerimeterBasedAccessRuleResponse)(nil)).Elem()
}

func (o PerimeterBasedAccessRuleResponseArrayOutput) ToPerimeterBasedAccessRuleResponseArrayOutput() PerimeterBasedAccessRuleResponseArrayOutput {
	return o
}

func (o PerimeterBasedAccessRuleResponseArrayOutput) ToPerimeterBasedAccessRuleResponseArrayOutputWithContext(ctx context.Context) PerimeterBasedAccessRuleResponseArrayOutput {
	return o
}

func (o PerimeterBasedAccessRuleResponseArrayOutput) Index(i pulumi.IntInput) PerimeterBasedAccessRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerimeterBasedAccessRuleResponse {
		return vs[0].([]PerimeterBasedAccessRuleResponse)[vs[1].(int)]
	}).(PerimeterBasedAccessRuleResponseOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

func (i SubResourceArgs) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput).ToSubResourcePtrOutputWithContext(ctx)
}









type SubResourcePtrInput interface {
	pulumi.Input

	ToSubResourcePtrOutput() SubResourcePtrOutput
	ToSubResourcePtrOutputWithContext(context.Context) SubResourcePtrOutput
}

type subResourcePtrType SubResourceArgs

func SubResourcePtr(v *SubResourceArgs) SubResourcePtrInput {
	return (*subResourcePtrType)(v)
}

func (*subResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (i *subResourcePtrType) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i *subResourcePtrType) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourcePtrOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o.ToSubResourcePtrOutputWithContext(context.Background())
}

func (o SubResourceOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResource) *SubResource {
		return &v
	}).(SubResourcePtrOutput)
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourcePtrOutput struct{ *pulumi.OutputState }

func (SubResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) Elem() SubResourceOutput {
	return o.ApplyT(func(v *SubResource) SubResource {
		if v != nil {
			return *v
		}
		var ret SubResource
		return ret
	}).(SubResourceOutput)
}

func (o SubResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(EffectiveConnectivityConfigurationResponseOutput{})
	pulumi.RegisterOutputType(EffectiveConnectivityConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(EffectiveVirtualNetworkResponseOutput{})
	pulumi.RegisterOutputType(EffectiveVirtualNetworkResponseArrayOutput{})
	pulumi.RegisterOutputType(GroupMembersItemOutput{})
	pulumi.RegisterOutputType(GroupMembersItemArrayOutput{})
	pulumi.RegisterOutputType(GroupMembersItemResponseOutput{})
	pulumi.RegisterOutputType(GroupMembersItemResponseArrayOutput{})
	pulumi.RegisterOutputType(HubOutput{})
	pulumi.RegisterOutputType(HubArrayOutput{})
	pulumi.RegisterOutputType(HubResponseOutput{})
	pulumi.RegisterOutputType(HubResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkManagerDeploymentStatusResponseOutput{})
	pulumi.RegisterOutputType(NetworkManagerDeploymentStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkManagerPropertiesNetworkManagerScopesOutput{})
	pulumi.RegisterOutputType(NetworkManagerPropertiesNetworkManagerScopesPtrOutput{})
	pulumi.RegisterOutputType(NetworkManagerPropertiesResponseNetworkManagerScopesOutput{})
	pulumi.RegisterOutputType(NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput{})
	pulumi.RegisterOutputType(NetworkManagerSecurityGroupItemOutput{})
	pulumi.RegisterOutputType(NetworkManagerSecurityGroupItemArrayOutput{})
	pulumi.RegisterOutputType(NetworkManagerSecurityGroupItemResponseOutput{})
	pulumi.RegisterOutputType(NetworkManagerSecurityGroupItemResponseArrayOutput{})
	pulumi.RegisterOutputType(PerimeterBasedAccessRuleOutput{})
	pulumi.RegisterOutputType(PerimeterBasedAccessRuleArrayOutput{})
	pulumi.RegisterOutputType(PerimeterBasedAccessRuleResponseOutput{})
	pulumi.RegisterOutputType(PerimeterBasedAccessRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

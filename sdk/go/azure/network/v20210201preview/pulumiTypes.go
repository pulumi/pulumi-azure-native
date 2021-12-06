


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





type ActiveConnectivityConfigurationResponseInput interface {
	pulumi.Input

	ToActiveConnectivityConfigurationResponseOutput() ActiveConnectivityConfigurationResponseOutput
	ToActiveConnectivityConfigurationResponseOutputWithContext(context.Context) ActiveConnectivityConfigurationResponseOutput
}

type ActiveConnectivityConfigurationResponseArgs struct {
	AppliesToGroups       ConnectivityGroupItemResponseArrayInput `pulumi:"appliesToGroups"`
	CommitTime            pulumi.StringPtrInput                   `pulumi:"commitTime"`
	ConfigurationGroups   ConfigurationGroupResponseArrayInput    `pulumi:"configurationGroups"`
	ConnectivityTopology  pulumi.StringInput                      `pulumi:"connectivityTopology"`
	DeleteExistingPeering pulumi.StringPtrInput                   `pulumi:"deleteExistingPeering"`
	Description           pulumi.StringPtrInput                   `pulumi:"description"`
	DisplayName           pulumi.StringPtrInput                   `pulumi:"displayName"`
	Hubs                  HubResponseArrayInput                   `pulumi:"hubs"`
	Id                    pulumi.StringPtrInput                   `pulumi:"id"`
	IsGlobal              pulumi.StringPtrInput                   `pulumi:"isGlobal"`
	ProvisioningState     pulumi.StringInput                      `pulumi:"provisioningState"`
	Region                pulumi.StringPtrInput                   `pulumi:"region"`
}

func (ActiveConnectivityConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveConnectivityConfigurationResponse)(nil)).Elem()
}

func (i ActiveConnectivityConfigurationResponseArgs) ToActiveConnectivityConfigurationResponseOutput() ActiveConnectivityConfigurationResponseOutput {
	return i.ToActiveConnectivityConfigurationResponseOutputWithContext(context.Background())
}

func (i ActiveConnectivityConfigurationResponseArgs) ToActiveConnectivityConfigurationResponseOutputWithContext(ctx context.Context) ActiveConnectivityConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveConnectivityConfigurationResponseOutput)
}





type ActiveConnectivityConfigurationResponseArrayInput interface {
	pulumi.Input

	ToActiveConnectivityConfigurationResponseArrayOutput() ActiveConnectivityConfigurationResponseArrayOutput
	ToActiveConnectivityConfigurationResponseArrayOutputWithContext(context.Context) ActiveConnectivityConfigurationResponseArrayOutput
}

type ActiveConnectivityConfigurationResponseArray []ActiveConnectivityConfigurationResponseInput

func (ActiveConnectivityConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveConnectivityConfigurationResponse)(nil)).Elem()
}

func (i ActiveConnectivityConfigurationResponseArray) ToActiveConnectivityConfigurationResponseArrayOutput() ActiveConnectivityConfigurationResponseArrayOutput {
	return i.ToActiveConnectivityConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i ActiveConnectivityConfigurationResponseArray) ToActiveConnectivityConfigurationResponseArrayOutputWithContext(ctx context.Context) ActiveConnectivityConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveConnectivityConfigurationResponseArrayOutput)
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





type ActiveDefaultSecurityAdminRuleResponseInput interface {
	pulumi.Input

	ToActiveDefaultSecurityAdminRuleResponseOutput() ActiveDefaultSecurityAdminRuleResponseOutput
	ToActiveDefaultSecurityAdminRuleResponseOutputWithContext(context.Context) ActiveDefaultSecurityAdminRuleResponseOutput
}

type ActiveDefaultSecurityAdminRuleResponseArgs struct {
	Access                        pulumi.StringInput                                `pulumi:"access"`
	CommitTime                    pulumi.StringPtrInput                             `pulumi:"commitTime"`
	ConfigurationDescription      pulumi.StringPtrInput                             `pulumi:"configurationDescription"`
	ConfigurationDisplayName      pulumi.StringPtrInput                             `pulumi:"configurationDisplayName"`
	Description                   pulumi.StringInput                                `pulumi:"description"`
	DestinationPortRanges         pulumi.StringArrayInput                           `pulumi:"destinationPortRanges"`
	Destinations                  AddressPrefixItemResponseArrayInput               `pulumi:"destinations"`
	Direction                     pulumi.StringInput                                `pulumi:"direction"`
	DisplayName                   pulumi.StringInput                                `pulumi:"displayName"`
	Flag                          pulumi.StringPtrInput                             `pulumi:"flag"`
	Id                            pulumi.StringPtrInput                             `pulumi:"id"`
	Kind                          pulumi.StringInput                                `pulumi:"kind"`
	Priority                      pulumi.IntInput                                   `pulumi:"priority"`
	Protocol                      pulumi.StringInput                                `pulumi:"protocol"`
	ProvisioningState             pulumi.StringInput                                `pulumi:"provisioningState"`
	Region                        pulumi.StringPtrInput                             `pulumi:"region"`
	RuleCollectionAppliesToGroups NetworkManagerSecurityGroupItemResponseArrayInput `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     pulumi.StringPtrInput                             `pulumi:"ruleCollectionDescription"`
	RuleCollectionDisplayName     pulumi.StringPtrInput                             `pulumi:"ruleCollectionDisplayName"`
	RuleGroups                    ConfigurationGroupResponseArrayInput              `pulumi:"ruleGroups"`
	SourcePortRanges              pulumi.StringArrayInput                           `pulumi:"sourcePortRanges"`
	Sources                       AddressPrefixItemResponseArrayInput               `pulumi:"sources"`
}

func (ActiveDefaultSecurityAdminRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDefaultSecurityAdminRuleResponse)(nil)).Elem()
}

func (i ActiveDefaultSecurityAdminRuleResponseArgs) ToActiveDefaultSecurityAdminRuleResponseOutput() ActiveDefaultSecurityAdminRuleResponseOutput {
	return i.ToActiveDefaultSecurityAdminRuleResponseOutputWithContext(context.Background())
}

func (i ActiveDefaultSecurityAdminRuleResponseArgs) ToActiveDefaultSecurityAdminRuleResponseOutputWithContext(ctx context.Context) ActiveDefaultSecurityAdminRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDefaultSecurityAdminRuleResponseOutput)
}

type ActiveDefaultSecurityAdminRuleResponseOutput struct{ *pulumi.OutputState }

func (ActiveDefaultSecurityAdminRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDefaultSecurityAdminRuleResponse)(nil)).Elem()
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) ToActiveDefaultSecurityAdminRuleResponseOutput() ActiveDefaultSecurityAdminRuleResponseOutput {
	return o
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) ToActiveDefaultSecurityAdminRuleResponseOutputWithContext(ctx context.Context) ActiveDefaultSecurityAdminRuleResponseOutput {
	return o
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) string { return v.Access }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) CommitTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) *string { return v.CommitTime }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) ConfigurationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) *string { return v.ConfigurationDescription }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) ConfigurationDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) *string { return v.ConfigurationDisplayName }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) []AddressPrefixItemResponse { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) Flag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) *string { return v.Flag }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) RuleCollectionAppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) []NetworkManagerSecurityGroupItemResponse {
		return v.RuleCollectionAppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) RuleCollectionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) *string { return v.RuleCollectionDescription }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) RuleCollectionDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) *string { return v.RuleCollectionDisplayName }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) RuleGroups() ConfigurationGroupResponseArrayOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) []ConfigurationGroupResponse { return v.RuleGroups }).(ConfigurationGroupResponseArrayOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func (o ActiveDefaultSecurityAdminRuleResponseOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityAdminRuleResponse) []AddressPrefixItemResponse { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
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





type ActiveDefaultSecurityUserRuleResponseInput interface {
	pulumi.Input

	ToActiveDefaultSecurityUserRuleResponseOutput() ActiveDefaultSecurityUserRuleResponseOutput
	ToActiveDefaultSecurityUserRuleResponseOutputWithContext(context.Context) ActiveDefaultSecurityUserRuleResponseOutput
}

type ActiveDefaultSecurityUserRuleResponseArgs struct {
	CommitTime                    pulumi.StringPtrInput                             `pulumi:"commitTime"`
	ConfigurationDescription      pulumi.StringPtrInput                             `pulumi:"configurationDescription"`
	ConfigurationDisplayName      pulumi.StringPtrInput                             `pulumi:"configurationDisplayName"`
	Description                   pulumi.StringInput                                `pulumi:"description"`
	DestinationPortRanges         pulumi.StringArrayInput                           `pulumi:"destinationPortRanges"`
	Destinations                  AddressPrefixItemResponseArrayInput               `pulumi:"destinations"`
	Direction                     pulumi.StringInput                                `pulumi:"direction"`
	DisplayName                   pulumi.StringInput                                `pulumi:"displayName"`
	Flag                          pulumi.StringPtrInput                             `pulumi:"flag"`
	Id                            pulumi.StringPtrInput                             `pulumi:"id"`
	Kind                          pulumi.StringInput                                `pulumi:"kind"`
	Protocol                      pulumi.StringInput                                `pulumi:"protocol"`
	ProvisioningState             pulumi.StringInput                                `pulumi:"provisioningState"`
	Region                        pulumi.StringPtrInput                             `pulumi:"region"`
	RuleCollectionAppliesToGroups NetworkManagerSecurityGroupItemResponseArrayInput `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     pulumi.StringPtrInput                             `pulumi:"ruleCollectionDescription"`
	RuleCollectionDisplayName     pulumi.StringPtrInput                             `pulumi:"ruleCollectionDisplayName"`
	RuleGroups                    ConfigurationGroupResponseArrayInput              `pulumi:"ruleGroups"`
	SourcePortRanges              pulumi.StringArrayInput                           `pulumi:"sourcePortRanges"`
	Sources                       AddressPrefixItemResponseArrayInput               `pulumi:"sources"`
}

func (ActiveDefaultSecurityUserRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDefaultSecurityUserRuleResponse)(nil)).Elem()
}

func (i ActiveDefaultSecurityUserRuleResponseArgs) ToActiveDefaultSecurityUserRuleResponseOutput() ActiveDefaultSecurityUserRuleResponseOutput {
	return i.ToActiveDefaultSecurityUserRuleResponseOutputWithContext(context.Background())
}

func (i ActiveDefaultSecurityUserRuleResponseArgs) ToActiveDefaultSecurityUserRuleResponseOutputWithContext(ctx context.Context) ActiveDefaultSecurityUserRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDefaultSecurityUserRuleResponseOutput)
}

type ActiveDefaultSecurityUserRuleResponseOutput struct{ *pulumi.OutputState }

func (ActiveDefaultSecurityUserRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDefaultSecurityUserRuleResponse)(nil)).Elem()
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) ToActiveDefaultSecurityUserRuleResponseOutput() ActiveDefaultSecurityUserRuleResponseOutput {
	return o
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) ToActiveDefaultSecurityUserRuleResponseOutputWithContext(ctx context.Context) ActiveDefaultSecurityUserRuleResponseOutput {
	return o
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) CommitTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) *string { return v.CommitTime }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) ConfigurationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) *string { return v.ConfigurationDescription }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) ConfigurationDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) *string { return v.ConfigurationDisplayName }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) []AddressPrefixItemResponse { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) Flag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) *string { return v.Flag }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) RuleCollectionAppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) []NetworkManagerSecurityGroupItemResponse {
		return v.RuleCollectionAppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) RuleCollectionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) *string { return v.RuleCollectionDescription }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) RuleCollectionDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) *string { return v.RuleCollectionDisplayName }).(pulumi.StringPtrOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) RuleGroups() ConfigurationGroupResponseArrayOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) []ConfigurationGroupResponse { return v.RuleGroups }).(ConfigurationGroupResponseArrayOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func (o ActiveDefaultSecurityUserRuleResponseOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveDefaultSecurityUserRuleResponse) []AddressPrefixItemResponse { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
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





type ActiveSecurityAdminRuleResponseInput interface {
	pulumi.Input

	ToActiveSecurityAdminRuleResponseOutput() ActiveSecurityAdminRuleResponseOutput
	ToActiveSecurityAdminRuleResponseOutputWithContext(context.Context) ActiveSecurityAdminRuleResponseOutput
}

type ActiveSecurityAdminRuleResponseArgs struct {
	Access                        pulumi.StringInput                                `pulumi:"access"`
	CommitTime                    pulumi.StringPtrInput                             `pulumi:"commitTime"`
	ConfigurationDescription      pulumi.StringPtrInput                             `pulumi:"configurationDescription"`
	ConfigurationDisplayName      pulumi.StringPtrInput                             `pulumi:"configurationDisplayName"`
	Description                   pulumi.StringPtrInput                             `pulumi:"description"`
	DestinationPortRanges         pulumi.StringArrayInput                           `pulumi:"destinationPortRanges"`
	Destinations                  AddressPrefixItemResponseArrayInput               `pulumi:"destinations"`
	Direction                     pulumi.StringInput                                `pulumi:"direction"`
	DisplayName                   pulumi.StringPtrInput                             `pulumi:"displayName"`
	Id                            pulumi.StringPtrInput                             `pulumi:"id"`
	Kind                          pulumi.StringInput                                `pulumi:"kind"`
	Priority                      pulumi.IntPtrInput                                `pulumi:"priority"`
	Protocol                      pulumi.StringInput                                `pulumi:"protocol"`
	ProvisioningState             pulumi.StringInput                                `pulumi:"provisioningState"`
	Region                        pulumi.StringPtrInput                             `pulumi:"region"`
	RuleCollectionAppliesToGroups NetworkManagerSecurityGroupItemResponseArrayInput `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     pulumi.StringPtrInput                             `pulumi:"ruleCollectionDescription"`
	RuleCollectionDisplayName     pulumi.StringPtrInput                             `pulumi:"ruleCollectionDisplayName"`
	RuleGroups                    ConfigurationGroupResponseArrayInput              `pulumi:"ruleGroups"`
	SourcePortRanges              pulumi.StringArrayInput                           `pulumi:"sourcePortRanges"`
	Sources                       AddressPrefixItemResponseArrayInput               `pulumi:"sources"`
}

func (ActiveSecurityAdminRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveSecurityAdminRuleResponse)(nil)).Elem()
}

func (i ActiveSecurityAdminRuleResponseArgs) ToActiveSecurityAdminRuleResponseOutput() ActiveSecurityAdminRuleResponseOutput {
	return i.ToActiveSecurityAdminRuleResponseOutputWithContext(context.Background())
}

func (i ActiveSecurityAdminRuleResponseArgs) ToActiveSecurityAdminRuleResponseOutputWithContext(ctx context.Context) ActiveSecurityAdminRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveSecurityAdminRuleResponseOutput)
}

type ActiveSecurityAdminRuleResponseOutput struct{ *pulumi.OutputState }

func (ActiveSecurityAdminRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveSecurityAdminRuleResponse)(nil)).Elem()
}

func (o ActiveSecurityAdminRuleResponseOutput) ToActiveSecurityAdminRuleResponseOutput() ActiveSecurityAdminRuleResponseOutput {
	return o
}

func (o ActiveSecurityAdminRuleResponseOutput) ToActiveSecurityAdminRuleResponseOutputWithContext(ctx context.Context) ActiveSecurityAdminRuleResponseOutput {
	return o
}

func (o ActiveSecurityAdminRuleResponseOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) string { return v.Access }).(pulumi.StringOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) CommitTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) *string { return v.CommitTime }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) ConfigurationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) *string { return v.ConfigurationDescription }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) ConfigurationDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) *string { return v.ConfigurationDisplayName }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) []AddressPrefixItemResponse { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) RuleCollectionAppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) []NetworkManagerSecurityGroupItemResponse {
		return v.RuleCollectionAppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) RuleCollectionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) *string { return v.RuleCollectionDescription }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) RuleCollectionDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) *string { return v.RuleCollectionDisplayName }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) RuleGroups() ConfigurationGroupResponseArrayOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) []ConfigurationGroupResponse { return v.RuleGroups }).(ConfigurationGroupResponseArrayOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func (o ActiveSecurityAdminRuleResponseOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveSecurityAdminRuleResponse) []AddressPrefixItemResponse { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
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





type ActiveSecurityUserRuleResponseInput interface {
	pulumi.Input

	ToActiveSecurityUserRuleResponseOutput() ActiveSecurityUserRuleResponseOutput
	ToActiveSecurityUserRuleResponseOutputWithContext(context.Context) ActiveSecurityUserRuleResponseOutput
}

type ActiveSecurityUserRuleResponseArgs struct {
	CommitTime                    pulumi.StringPtrInput                             `pulumi:"commitTime"`
	ConfigurationDescription      pulumi.StringPtrInput                             `pulumi:"configurationDescription"`
	ConfigurationDisplayName      pulumi.StringPtrInput                             `pulumi:"configurationDisplayName"`
	Description                   pulumi.StringPtrInput                             `pulumi:"description"`
	DestinationPortRanges         pulumi.StringArrayInput                           `pulumi:"destinationPortRanges"`
	Destinations                  AddressPrefixItemResponseArrayInput               `pulumi:"destinations"`
	Direction                     pulumi.StringInput                                `pulumi:"direction"`
	DisplayName                   pulumi.StringPtrInput                             `pulumi:"displayName"`
	Id                            pulumi.StringPtrInput                             `pulumi:"id"`
	Kind                          pulumi.StringInput                                `pulumi:"kind"`
	Protocol                      pulumi.StringInput                                `pulumi:"protocol"`
	ProvisioningState             pulumi.StringInput                                `pulumi:"provisioningState"`
	Region                        pulumi.StringPtrInput                             `pulumi:"region"`
	RuleCollectionAppliesToGroups NetworkManagerSecurityGroupItemResponseArrayInput `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     pulumi.StringPtrInput                             `pulumi:"ruleCollectionDescription"`
	RuleCollectionDisplayName     pulumi.StringPtrInput                             `pulumi:"ruleCollectionDisplayName"`
	RuleGroups                    ConfigurationGroupResponseArrayInput              `pulumi:"ruleGroups"`
	SourcePortRanges              pulumi.StringArrayInput                           `pulumi:"sourcePortRanges"`
	Sources                       AddressPrefixItemResponseArrayInput               `pulumi:"sources"`
}

func (ActiveSecurityUserRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveSecurityUserRuleResponse)(nil)).Elem()
}

func (i ActiveSecurityUserRuleResponseArgs) ToActiveSecurityUserRuleResponseOutput() ActiveSecurityUserRuleResponseOutput {
	return i.ToActiveSecurityUserRuleResponseOutputWithContext(context.Background())
}

func (i ActiveSecurityUserRuleResponseArgs) ToActiveSecurityUserRuleResponseOutputWithContext(ctx context.Context) ActiveSecurityUserRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveSecurityUserRuleResponseOutput)
}

type ActiveSecurityUserRuleResponseOutput struct{ *pulumi.OutputState }

func (ActiveSecurityUserRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveSecurityUserRuleResponse)(nil)).Elem()
}

func (o ActiveSecurityUserRuleResponseOutput) ToActiveSecurityUserRuleResponseOutput() ActiveSecurityUserRuleResponseOutput {
	return o
}

func (o ActiveSecurityUserRuleResponseOutput) ToActiveSecurityUserRuleResponseOutputWithContext(ctx context.Context) ActiveSecurityUserRuleResponseOutput {
	return o
}

func (o ActiveSecurityUserRuleResponseOutput) CommitTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) *string { return v.CommitTime }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) ConfigurationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) *string { return v.ConfigurationDescription }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) ConfigurationDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) *string { return v.ConfigurationDisplayName }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) []AddressPrefixItemResponse { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) RuleCollectionAppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) []NetworkManagerSecurityGroupItemResponse {
		return v.RuleCollectionAppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) RuleCollectionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) *string { return v.RuleCollectionDescription }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) RuleCollectionDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) *string { return v.RuleCollectionDisplayName }).(pulumi.StringPtrOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) RuleGroups() ConfigurationGroupResponseArrayOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) []ConfigurationGroupResponse { return v.RuleGroups }).(ConfigurationGroupResponseArrayOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func (o ActiveSecurityUserRuleResponseOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v ActiveSecurityUserRuleResponse) []AddressPrefixItemResponse { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
}

type AddressPrefixItem struct {
	AddressPrefix     *string `pulumi:"addressPrefix"`
	AddressPrefixType *string `pulumi:"addressPrefixType"`
}





type AddressPrefixItemInput interface {
	pulumi.Input

	ToAddressPrefixItemOutput() AddressPrefixItemOutput
	ToAddressPrefixItemOutputWithContext(context.Context) AddressPrefixItemOutput
}

type AddressPrefixItemArgs struct {
	AddressPrefix     pulumi.StringPtrInput `pulumi:"addressPrefix"`
	AddressPrefixType pulumi.StringPtrInput `pulumi:"addressPrefixType"`
}

func (AddressPrefixItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressPrefixItem)(nil)).Elem()
}

func (i AddressPrefixItemArgs) ToAddressPrefixItemOutput() AddressPrefixItemOutput {
	return i.ToAddressPrefixItemOutputWithContext(context.Background())
}

func (i AddressPrefixItemArgs) ToAddressPrefixItemOutputWithContext(ctx context.Context) AddressPrefixItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPrefixItemOutput)
}

type AddressPrefixItemOutput struct{ *pulumi.OutputState }

func (AddressPrefixItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressPrefixItem)(nil)).Elem()
}

func (o AddressPrefixItemOutput) ToAddressPrefixItemOutput() AddressPrefixItemOutput {
	return o
}

func (o AddressPrefixItemOutput) ToAddressPrefixItemOutputWithContext(ctx context.Context) AddressPrefixItemOutput {
	return o
}

func (o AddressPrefixItemOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressPrefixItem) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o AddressPrefixItemOutput) AddressPrefixType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressPrefixItem) *string { return v.AddressPrefixType }).(pulumi.StringPtrOutput)
}

type AddressPrefixItemResponse struct {
	AddressPrefix     *string `pulumi:"addressPrefix"`
	AddressPrefixType *string `pulumi:"addressPrefixType"`
}





type AddressPrefixItemResponseInput interface {
	pulumi.Input

	ToAddressPrefixItemResponseOutput() AddressPrefixItemResponseOutput
	ToAddressPrefixItemResponseOutputWithContext(context.Context) AddressPrefixItemResponseOutput
}

type AddressPrefixItemResponseArgs struct {
	AddressPrefix     pulumi.StringPtrInput `pulumi:"addressPrefix"`
	AddressPrefixType pulumi.StringPtrInput `pulumi:"addressPrefixType"`
}

func (AddressPrefixItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressPrefixItemResponse)(nil)).Elem()
}

func (i AddressPrefixItemResponseArgs) ToAddressPrefixItemResponseOutput() AddressPrefixItemResponseOutput {
	return i.ToAddressPrefixItemResponseOutputWithContext(context.Background())
}

func (i AddressPrefixItemResponseArgs) ToAddressPrefixItemResponseOutputWithContext(ctx context.Context) AddressPrefixItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPrefixItemResponseOutput)
}





type AddressPrefixItemResponseArrayInput interface {
	pulumi.Input

	ToAddressPrefixItemResponseArrayOutput() AddressPrefixItemResponseArrayOutput
	ToAddressPrefixItemResponseArrayOutputWithContext(context.Context) AddressPrefixItemResponseArrayOutput
}

type AddressPrefixItemResponseArray []AddressPrefixItemResponseInput

func (AddressPrefixItemResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AddressPrefixItemResponse)(nil)).Elem()
}

func (i AddressPrefixItemResponseArray) ToAddressPrefixItemResponseArrayOutput() AddressPrefixItemResponseArrayOutput {
	return i.ToAddressPrefixItemResponseArrayOutputWithContext(context.Background())
}

func (i AddressPrefixItemResponseArray) ToAddressPrefixItemResponseArrayOutputWithContext(ctx context.Context) AddressPrefixItemResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPrefixItemResponseArrayOutput)
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





type ConfigurationGroupResponseInput interface {
	pulumi.Input

	ToConfigurationGroupResponseOutput() ConfigurationGroupResponseOutput
	ToConfigurationGroupResponseOutputWithContext(context.Context) ConfigurationGroupResponseOutput
}

type ConfigurationGroupResponseArgs struct {
	ConditionalMembership pulumi.StringPtrInput              `pulumi:"conditionalMembership"`
	Description           pulumi.StringPtrInput              `pulumi:"description"`
	DisplayName           pulumi.StringPtrInput              `pulumi:"displayName"`
	GroupMembers          GroupMembersItemResponseArrayInput `pulumi:"groupMembers"`
	Id                    pulumi.StringPtrInput              `pulumi:"id"`
	MemberType            pulumi.StringPtrInput              `pulumi:"memberType"`
	ProvisioningState     pulumi.StringInput                 `pulumi:"provisioningState"`
}

func (ConfigurationGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationGroupResponse)(nil)).Elem()
}

func (i ConfigurationGroupResponseArgs) ToConfigurationGroupResponseOutput() ConfigurationGroupResponseOutput {
	return i.ToConfigurationGroupResponseOutputWithContext(context.Background())
}

func (i ConfigurationGroupResponseArgs) ToConfigurationGroupResponseOutputWithContext(ctx context.Context) ConfigurationGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationGroupResponseOutput)
}





type ConfigurationGroupResponseArrayInput interface {
	pulumi.Input

	ToConfigurationGroupResponseArrayOutput() ConfigurationGroupResponseArrayOutput
	ToConfigurationGroupResponseArrayOutputWithContext(context.Context) ConfigurationGroupResponseArrayOutput
}

type ConfigurationGroupResponseArray []ConfigurationGroupResponseInput

func (ConfigurationGroupResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationGroupResponse)(nil)).Elem()
}

func (i ConfigurationGroupResponseArray) ToConfigurationGroupResponseArrayOutput() ConfigurationGroupResponseArrayOutput {
	return i.ToConfigurationGroupResponseArrayOutputWithContext(context.Background())
}

func (i ConfigurationGroupResponseArray) ToConfigurationGroupResponseArrayOutputWithContext(ctx context.Context) ConfigurationGroupResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationGroupResponseArrayOutput)
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





type ConnectivityGroupItemResponseInput interface {
	pulumi.Input

	ToConnectivityGroupItemResponseOutput() ConnectivityGroupItemResponseOutput
	ToConnectivityGroupItemResponseOutputWithContext(context.Context) ConnectivityGroupItemResponseOutput
}

type ConnectivityGroupItemResponseArgs struct {
	GroupConnectivity pulumi.StringPtrInput `pulumi:"groupConnectivity"`
	IsGlobal          pulumi.StringPtrInput `pulumi:"isGlobal"`
	NetworkGroupId    pulumi.StringPtrInput `pulumi:"networkGroupId"`
	UseHubGateway     pulumi.StringPtrInput `pulumi:"useHubGateway"`
}

func (ConnectivityGroupItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityGroupItemResponse)(nil)).Elem()
}

func (i ConnectivityGroupItemResponseArgs) ToConnectivityGroupItemResponseOutput() ConnectivityGroupItemResponseOutput {
	return i.ToConnectivityGroupItemResponseOutputWithContext(context.Background())
}

func (i ConnectivityGroupItemResponseArgs) ToConnectivityGroupItemResponseOutputWithContext(ctx context.Context) ConnectivityGroupItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectivityGroupItemResponseOutput)
}





type ConnectivityGroupItemResponseArrayInput interface {
	pulumi.Input

	ToConnectivityGroupItemResponseArrayOutput() ConnectivityGroupItemResponseArrayOutput
	ToConnectivityGroupItemResponseArrayOutputWithContext(context.Context) ConnectivityGroupItemResponseArrayOutput
}

type ConnectivityGroupItemResponseArray []ConnectivityGroupItemResponseInput

func (ConnectivityGroupItemResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectivityGroupItemResponse)(nil)).Elem()
}

func (i ConnectivityGroupItemResponseArray) ToConnectivityGroupItemResponseArrayOutput() ConnectivityGroupItemResponseArrayOutput {
	return i.ToConnectivityGroupItemResponseArrayOutputWithContext(context.Background())
}

func (i ConnectivityGroupItemResponseArray) ToConnectivityGroupItemResponseArrayOutputWithContext(ctx context.Context) ConnectivityGroupItemResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectivityGroupItemResponseArrayOutput)
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





type EffectiveConnectivityConfigurationResponseInput interface {
	pulumi.Input

	ToEffectiveConnectivityConfigurationResponseOutput() EffectiveConnectivityConfigurationResponseOutput
	ToEffectiveConnectivityConfigurationResponseOutputWithContext(context.Context) EffectiveConnectivityConfigurationResponseOutput
}

type EffectiveConnectivityConfigurationResponseArgs struct {
	AppliesToGroups       ConnectivityGroupItemResponseArrayInput `pulumi:"appliesToGroups"`
	ConfigurationGroups   ConfigurationGroupResponseArrayInput    `pulumi:"configurationGroups"`
	ConnectivityTopology  pulumi.StringInput                      `pulumi:"connectivityTopology"`
	DeleteExistingPeering pulumi.StringPtrInput                   `pulumi:"deleteExistingPeering"`
	Description           pulumi.StringPtrInput                   `pulumi:"description"`
	DisplayName           pulumi.StringPtrInput                   `pulumi:"displayName"`
	Hubs                  HubResponseArrayInput                   `pulumi:"hubs"`
	Id                    pulumi.StringPtrInput                   `pulumi:"id"`
	IsGlobal              pulumi.StringPtrInput                   `pulumi:"isGlobal"`
	ProvisioningState     pulumi.StringInput                      `pulumi:"provisioningState"`
}

func (EffectiveConnectivityConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EffectiveConnectivityConfigurationResponse)(nil)).Elem()
}

func (i EffectiveConnectivityConfigurationResponseArgs) ToEffectiveConnectivityConfigurationResponseOutput() EffectiveConnectivityConfigurationResponseOutput {
	return i.ToEffectiveConnectivityConfigurationResponseOutputWithContext(context.Background())
}

func (i EffectiveConnectivityConfigurationResponseArgs) ToEffectiveConnectivityConfigurationResponseOutputWithContext(ctx context.Context) EffectiveConnectivityConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EffectiveConnectivityConfigurationResponseOutput)
}





type EffectiveConnectivityConfigurationResponseArrayInput interface {
	pulumi.Input

	ToEffectiveConnectivityConfigurationResponseArrayOutput() EffectiveConnectivityConfigurationResponseArrayOutput
	ToEffectiveConnectivityConfigurationResponseArrayOutputWithContext(context.Context) EffectiveConnectivityConfigurationResponseArrayOutput
}

type EffectiveConnectivityConfigurationResponseArray []EffectiveConnectivityConfigurationResponseInput

func (EffectiveConnectivityConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EffectiveConnectivityConfigurationResponse)(nil)).Elem()
}

func (i EffectiveConnectivityConfigurationResponseArray) ToEffectiveConnectivityConfigurationResponseArrayOutput() EffectiveConnectivityConfigurationResponseArrayOutput {
	return i.ToEffectiveConnectivityConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i EffectiveConnectivityConfigurationResponseArray) ToEffectiveConnectivityConfigurationResponseArrayOutputWithContext(ctx context.Context) EffectiveConnectivityConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EffectiveConnectivityConfigurationResponseArrayOutput)
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





type EffectiveDefaultSecurityAdminRuleResponseInput interface {
	pulumi.Input

	ToEffectiveDefaultSecurityAdminRuleResponseOutput() EffectiveDefaultSecurityAdminRuleResponseOutput
	ToEffectiveDefaultSecurityAdminRuleResponseOutputWithContext(context.Context) EffectiveDefaultSecurityAdminRuleResponseOutput
}

type EffectiveDefaultSecurityAdminRuleResponseArgs struct {
	Access                        pulumi.StringInput                                `pulumi:"access"`
	ConfigurationDescription      pulumi.StringPtrInput                             `pulumi:"configurationDescription"`
	ConfigurationDisplayName      pulumi.StringPtrInput                             `pulumi:"configurationDisplayName"`
	Description                   pulumi.StringInput                                `pulumi:"description"`
	DestinationPortRanges         pulumi.StringArrayInput                           `pulumi:"destinationPortRanges"`
	Destinations                  AddressPrefixItemResponseArrayInput               `pulumi:"destinations"`
	Direction                     pulumi.StringInput                                `pulumi:"direction"`
	DisplayName                   pulumi.StringInput                                `pulumi:"displayName"`
	Flag                          pulumi.StringPtrInput                             `pulumi:"flag"`
	Id                            pulumi.StringPtrInput                             `pulumi:"id"`
	Kind                          pulumi.StringInput                                `pulumi:"kind"`
	Priority                      pulumi.IntInput                                   `pulumi:"priority"`
	Protocol                      pulumi.StringInput                                `pulumi:"protocol"`
	ProvisioningState             pulumi.StringInput                                `pulumi:"provisioningState"`
	RuleCollectionAppliesToGroups NetworkManagerSecurityGroupItemResponseArrayInput `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     pulumi.StringPtrInput                             `pulumi:"ruleCollectionDescription"`
	RuleCollectionDisplayName     pulumi.StringPtrInput                             `pulumi:"ruleCollectionDisplayName"`
	RuleGroups                    ConfigurationGroupResponseArrayInput              `pulumi:"ruleGroups"`
	SourcePortRanges              pulumi.StringArrayInput                           `pulumi:"sourcePortRanges"`
	Sources                       AddressPrefixItemResponseArrayInput               `pulumi:"sources"`
}

func (EffectiveDefaultSecurityAdminRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EffectiveDefaultSecurityAdminRuleResponse)(nil)).Elem()
}

func (i EffectiveDefaultSecurityAdminRuleResponseArgs) ToEffectiveDefaultSecurityAdminRuleResponseOutput() EffectiveDefaultSecurityAdminRuleResponseOutput {
	return i.ToEffectiveDefaultSecurityAdminRuleResponseOutputWithContext(context.Background())
}

func (i EffectiveDefaultSecurityAdminRuleResponseArgs) ToEffectiveDefaultSecurityAdminRuleResponseOutputWithContext(ctx context.Context) EffectiveDefaultSecurityAdminRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EffectiveDefaultSecurityAdminRuleResponseOutput)
}

type EffectiveDefaultSecurityAdminRuleResponseOutput struct{ *pulumi.OutputState }

func (EffectiveDefaultSecurityAdminRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EffectiveDefaultSecurityAdminRuleResponse)(nil)).Elem()
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) ToEffectiveDefaultSecurityAdminRuleResponseOutput() EffectiveDefaultSecurityAdminRuleResponseOutput {
	return o
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) ToEffectiveDefaultSecurityAdminRuleResponseOutputWithContext(ctx context.Context) EffectiveDefaultSecurityAdminRuleResponseOutput {
	return o
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) string { return v.Access }).(pulumi.StringOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) ConfigurationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) *string { return v.ConfigurationDescription }).(pulumi.StringPtrOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) ConfigurationDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) *string { return v.ConfigurationDisplayName }).(pulumi.StringPtrOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) []AddressPrefixItemResponse { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) Flag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) *string { return v.Flag }).(pulumi.StringPtrOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) RuleCollectionAppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) []NetworkManagerSecurityGroupItemResponse {
		return v.RuleCollectionAppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) RuleCollectionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) *string { return v.RuleCollectionDescription }).(pulumi.StringPtrOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) RuleCollectionDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) *string { return v.RuleCollectionDisplayName }).(pulumi.StringPtrOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) RuleGroups() ConfigurationGroupResponseArrayOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) []ConfigurationGroupResponse { return v.RuleGroups }).(ConfigurationGroupResponseArrayOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func (o EffectiveDefaultSecurityAdminRuleResponseOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v EffectiveDefaultSecurityAdminRuleResponse) []AddressPrefixItemResponse { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
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





type EffectiveSecurityAdminRuleResponseInput interface {
	pulumi.Input

	ToEffectiveSecurityAdminRuleResponseOutput() EffectiveSecurityAdminRuleResponseOutput
	ToEffectiveSecurityAdminRuleResponseOutputWithContext(context.Context) EffectiveSecurityAdminRuleResponseOutput
}

type EffectiveSecurityAdminRuleResponseArgs struct {
	Access                        pulumi.StringInput                                `pulumi:"access"`
	ConfigurationDescription      pulumi.StringPtrInput                             `pulumi:"configurationDescription"`
	ConfigurationDisplayName      pulumi.StringPtrInput                             `pulumi:"configurationDisplayName"`
	Description                   pulumi.StringPtrInput                             `pulumi:"description"`
	DestinationPortRanges         pulumi.StringArrayInput                           `pulumi:"destinationPortRanges"`
	Destinations                  AddressPrefixItemResponseArrayInput               `pulumi:"destinations"`
	Direction                     pulumi.StringInput                                `pulumi:"direction"`
	DisplayName                   pulumi.StringPtrInput                             `pulumi:"displayName"`
	Id                            pulumi.StringPtrInput                             `pulumi:"id"`
	Kind                          pulumi.StringInput                                `pulumi:"kind"`
	Priority                      pulumi.IntPtrInput                                `pulumi:"priority"`
	Protocol                      pulumi.StringInput                                `pulumi:"protocol"`
	ProvisioningState             pulumi.StringInput                                `pulumi:"provisioningState"`
	RuleCollectionAppliesToGroups NetworkManagerSecurityGroupItemResponseArrayInput `pulumi:"ruleCollectionAppliesToGroups"`
	RuleCollectionDescription     pulumi.StringPtrInput                             `pulumi:"ruleCollectionDescription"`
	RuleCollectionDisplayName     pulumi.StringPtrInput                             `pulumi:"ruleCollectionDisplayName"`
	RuleGroups                    ConfigurationGroupResponseArrayInput              `pulumi:"ruleGroups"`
	SourcePortRanges              pulumi.StringArrayInput                           `pulumi:"sourcePortRanges"`
	Sources                       AddressPrefixItemResponseArrayInput               `pulumi:"sources"`
}

func (EffectiveSecurityAdminRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EffectiveSecurityAdminRuleResponse)(nil)).Elem()
}

func (i EffectiveSecurityAdminRuleResponseArgs) ToEffectiveSecurityAdminRuleResponseOutput() EffectiveSecurityAdminRuleResponseOutput {
	return i.ToEffectiveSecurityAdminRuleResponseOutputWithContext(context.Background())
}

func (i EffectiveSecurityAdminRuleResponseArgs) ToEffectiveSecurityAdminRuleResponseOutputWithContext(ctx context.Context) EffectiveSecurityAdminRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EffectiveSecurityAdminRuleResponseOutput)
}

type EffectiveSecurityAdminRuleResponseOutput struct{ *pulumi.OutputState }

func (EffectiveSecurityAdminRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EffectiveSecurityAdminRuleResponse)(nil)).Elem()
}

func (o EffectiveSecurityAdminRuleResponseOutput) ToEffectiveSecurityAdminRuleResponseOutput() EffectiveSecurityAdminRuleResponseOutput {
	return o
}

func (o EffectiveSecurityAdminRuleResponseOutput) ToEffectiveSecurityAdminRuleResponseOutputWithContext(ctx context.Context) EffectiveSecurityAdminRuleResponseOutput {
	return o
}

func (o EffectiveSecurityAdminRuleResponseOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) string { return v.Access }).(pulumi.StringOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) ConfigurationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) *string { return v.ConfigurationDescription }).(pulumi.StringPtrOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) ConfigurationDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) *string { return v.ConfigurationDisplayName }).(pulumi.StringPtrOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) []AddressPrefixItemResponse { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) RuleCollectionAppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) []NetworkManagerSecurityGroupItemResponse {
		return v.RuleCollectionAppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) RuleCollectionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) *string { return v.RuleCollectionDescription }).(pulumi.StringPtrOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) RuleCollectionDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) *string { return v.RuleCollectionDisplayName }).(pulumi.StringPtrOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) RuleGroups() ConfigurationGroupResponseArrayOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) []ConfigurationGroupResponse { return v.RuleGroups }).(ConfigurationGroupResponseArrayOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func (o EffectiveSecurityAdminRuleResponseOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v EffectiveSecurityAdminRuleResponse) []AddressPrefixItemResponse { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
}

type EffectiveVirtualNetworkResponse struct {
	Id             *string `pulumi:"id"`
	Location       *string `pulumi:"location"`
	MembershipType *string `pulumi:"membershipType"`
}





type EffectiveVirtualNetworkResponseInput interface {
	pulumi.Input

	ToEffectiveVirtualNetworkResponseOutput() EffectiveVirtualNetworkResponseOutput
	ToEffectiveVirtualNetworkResponseOutputWithContext(context.Context) EffectiveVirtualNetworkResponseOutput
}

type EffectiveVirtualNetworkResponseArgs struct {
	Id             pulumi.StringPtrInput `pulumi:"id"`
	Location       pulumi.StringPtrInput `pulumi:"location"`
	MembershipType pulumi.StringPtrInput `pulumi:"membershipType"`
}

func (EffectiveVirtualNetworkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EffectiveVirtualNetworkResponse)(nil)).Elem()
}

func (i EffectiveVirtualNetworkResponseArgs) ToEffectiveVirtualNetworkResponseOutput() EffectiveVirtualNetworkResponseOutput {
	return i.ToEffectiveVirtualNetworkResponseOutputWithContext(context.Background())
}

func (i EffectiveVirtualNetworkResponseArgs) ToEffectiveVirtualNetworkResponseOutputWithContext(ctx context.Context) EffectiveVirtualNetworkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EffectiveVirtualNetworkResponseOutput)
}





type EffectiveVirtualNetworkResponseArrayInput interface {
	pulumi.Input

	ToEffectiveVirtualNetworkResponseArrayOutput() EffectiveVirtualNetworkResponseArrayOutput
	ToEffectiveVirtualNetworkResponseArrayOutputWithContext(context.Context) EffectiveVirtualNetworkResponseArrayOutput
}

type EffectiveVirtualNetworkResponseArray []EffectiveVirtualNetworkResponseInput

func (EffectiveVirtualNetworkResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EffectiveVirtualNetworkResponse)(nil)).Elem()
}

func (i EffectiveVirtualNetworkResponseArray) ToEffectiveVirtualNetworkResponseArrayOutput() EffectiveVirtualNetworkResponseArrayOutput {
	return i.ToEffectiveVirtualNetworkResponseArrayOutputWithContext(context.Background())
}

func (i EffectiveVirtualNetworkResponseArray) ToEffectiveVirtualNetworkResponseArrayOutputWithContext(ctx context.Context) EffectiveVirtualNetworkResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EffectiveVirtualNetworkResponseArrayOutput)
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





type GroupMembersItemResponseInput interface {
	pulumi.Input

	ToGroupMembersItemResponseOutput() GroupMembersItemResponseOutput
	ToGroupMembersItemResponseOutputWithContext(context.Context) GroupMembersItemResponseOutput
}

type GroupMembersItemResponseArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (GroupMembersItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembersItemResponse)(nil)).Elem()
}

func (i GroupMembersItemResponseArgs) ToGroupMembersItemResponseOutput() GroupMembersItemResponseOutput {
	return i.ToGroupMembersItemResponseOutputWithContext(context.Background())
}

func (i GroupMembersItemResponseArgs) ToGroupMembersItemResponseOutputWithContext(ctx context.Context) GroupMembersItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembersItemResponseOutput)
}





type GroupMembersItemResponseArrayInput interface {
	pulumi.Input

	ToGroupMembersItemResponseArrayOutput() GroupMembersItemResponseArrayOutput
	ToGroupMembersItemResponseArrayOutputWithContext(context.Context) GroupMembersItemResponseArrayOutput
}

type GroupMembersItemResponseArray []GroupMembersItemResponseInput

func (GroupMembersItemResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupMembersItemResponse)(nil)).Elem()
}

func (i GroupMembersItemResponseArray) ToGroupMembersItemResponseArrayOutput() GroupMembersItemResponseArrayOutput {
	return i.ToGroupMembersItemResponseArrayOutputWithContext(context.Background())
}

func (i GroupMembersItemResponseArray) ToGroupMembersItemResponseArrayOutputWithContext(ctx context.Context) GroupMembersItemResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembersItemResponseArrayOutput)
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





type HubResponseInput interface {
	pulumi.Input

	ToHubResponseOutput() HubResponseOutput
	ToHubResponseOutputWithContext(context.Context) HubResponseOutput
}

type HubResponseArgs struct {
	ResourceId   pulumi.StringPtrInput `pulumi:"resourceId"`
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
}

func (HubResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HubResponse)(nil)).Elem()
}

func (i HubResponseArgs) ToHubResponseOutput() HubResponseOutput {
	return i.ToHubResponseOutputWithContext(context.Background())
}

func (i HubResponseArgs) ToHubResponseOutputWithContext(ctx context.Context) HubResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubResponseOutput)
}





type HubResponseArrayInput interface {
	pulumi.Input

	ToHubResponseArrayOutput() HubResponseArrayOutput
	ToHubResponseArrayOutputWithContext(context.Context) HubResponseArrayOutput
}

type HubResponseArray []HubResponseInput

func (HubResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HubResponse)(nil)).Elem()
}

func (i HubResponseArray) ToHubResponseArrayOutput() HubResponseArrayOutput {
	return i.ToHubResponseArrayOutputWithContext(context.Background())
}

func (i HubResponseArray) ToHubResponseArrayOutputWithContext(ctx context.Context) HubResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubResponseArrayOutput)
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





type NetworkManagerDeploymentStatusResponseInput interface {
	pulumi.Input

	ToNetworkManagerDeploymentStatusResponseOutput() NetworkManagerDeploymentStatusResponseOutput
	ToNetworkManagerDeploymentStatusResponseOutputWithContext(context.Context) NetworkManagerDeploymentStatusResponseOutput
}

type NetworkManagerDeploymentStatusResponseArgs struct {
	CommitTime       pulumi.StringPtrInput   `pulumi:"commitTime"`
	ConfigurationIds pulumi.StringArrayInput `pulumi:"configurationIds"`
	DeploymentStatus pulumi.StringPtrInput   `pulumi:"deploymentStatus"`
	DeploymentType   pulumi.StringPtrInput   `pulumi:"deploymentType"`
	ErrorMessage     pulumi.StringPtrInput   `pulumi:"errorMessage"`
	Region           pulumi.StringPtrInput   `pulumi:"region"`
}

func (NetworkManagerDeploymentStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkManagerDeploymentStatusResponse)(nil)).Elem()
}

func (i NetworkManagerDeploymentStatusResponseArgs) ToNetworkManagerDeploymentStatusResponseOutput() NetworkManagerDeploymentStatusResponseOutput {
	return i.ToNetworkManagerDeploymentStatusResponseOutputWithContext(context.Background())
}

func (i NetworkManagerDeploymentStatusResponseArgs) ToNetworkManagerDeploymentStatusResponseOutputWithContext(ctx context.Context) NetworkManagerDeploymentStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerDeploymentStatusResponseOutput)
}





type NetworkManagerDeploymentStatusResponseArrayInput interface {
	pulumi.Input

	ToNetworkManagerDeploymentStatusResponseArrayOutput() NetworkManagerDeploymentStatusResponseArrayOutput
	ToNetworkManagerDeploymentStatusResponseArrayOutputWithContext(context.Context) NetworkManagerDeploymentStatusResponseArrayOutput
}

type NetworkManagerDeploymentStatusResponseArray []NetworkManagerDeploymentStatusResponseInput

func (NetworkManagerDeploymentStatusResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkManagerDeploymentStatusResponse)(nil)).Elem()
}

func (i NetworkManagerDeploymentStatusResponseArray) ToNetworkManagerDeploymentStatusResponseArrayOutput() NetworkManagerDeploymentStatusResponseArrayOutput {
	return i.ToNetworkManagerDeploymentStatusResponseArrayOutputWithContext(context.Background())
}

func (i NetworkManagerDeploymentStatusResponseArray) ToNetworkManagerDeploymentStatusResponseArrayOutputWithContext(ctx context.Context) NetworkManagerDeploymentStatusResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerDeploymentStatusResponseArrayOutput)
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





type NetworkManagerPropertiesResponseNetworkManagerScopesInput interface {
	pulumi.Input

	ToNetworkManagerPropertiesResponseNetworkManagerScopesOutput() NetworkManagerPropertiesResponseNetworkManagerScopesOutput
	ToNetworkManagerPropertiesResponseNetworkManagerScopesOutputWithContext(context.Context) NetworkManagerPropertiesResponseNetworkManagerScopesOutput
}

type NetworkManagerPropertiesResponseNetworkManagerScopesArgs struct {
	ManagementGroups pulumi.StringArrayInput `pulumi:"managementGroups"`
	Subscriptions    pulumi.StringArrayInput `pulumi:"subscriptions"`
}

func (NetworkManagerPropertiesResponseNetworkManagerScopesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkManagerPropertiesResponseNetworkManagerScopes)(nil)).Elem()
}

func (i NetworkManagerPropertiesResponseNetworkManagerScopesArgs) ToNetworkManagerPropertiesResponseNetworkManagerScopesOutput() NetworkManagerPropertiesResponseNetworkManagerScopesOutput {
	return i.ToNetworkManagerPropertiesResponseNetworkManagerScopesOutputWithContext(context.Background())
}

func (i NetworkManagerPropertiesResponseNetworkManagerScopesArgs) ToNetworkManagerPropertiesResponseNetworkManagerScopesOutputWithContext(ctx context.Context) NetworkManagerPropertiesResponseNetworkManagerScopesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerPropertiesResponseNetworkManagerScopesOutput)
}

func (i NetworkManagerPropertiesResponseNetworkManagerScopesArgs) ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput() NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput {
	return i.ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutputWithContext(context.Background())
}

func (i NetworkManagerPropertiesResponseNetworkManagerScopesArgs) ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutputWithContext(ctx context.Context) NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerPropertiesResponseNetworkManagerScopesOutput).ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutputWithContext(ctx)
}









type NetworkManagerPropertiesResponseNetworkManagerScopesPtrInput interface {
	pulumi.Input

	ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput() NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput
	ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutputWithContext(context.Context) NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput
}

type networkManagerPropertiesResponseNetworkManagerScopesPtrType NetworkManagerPropertiesResponseNetworkManagerScopesArgs

func NetworkManagerPropertiesResponseNetworkManagerScopesPtr(v *NetworkManagerPropertiesResponseNetworkManagerScopesArgs) NetworkManagerPropertiesResponseNetworkManagerScopesPtrInput {
	return (*networkManagerPropertiesResponseNetworkManagerScopesPtrType)(v)
}

func (*networkManagerPropertiesResponseNetworkManagerScopesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkManagerPropertiesResponseNetworkManagerScopes)(nil)).Elem()
}

func (i *networkManagerPropertiesResponseNetworkManagerScopesPtrType) ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput() NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput {
	return i.ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutputWithContext(context.Background())
}

func (i *networkManagerPropertiesResponseNetworkManagerScopesPtrType) ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutputWithContext(ctx context.Context) NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput)
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

func (o NetworkManagerPropertiesResponseNetworkManagerScopesOutput) ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput() NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput {
	return o.ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutputWithContext(context.Background())
}

func (o NetworkManagerPropertiesResponseNetworkManagerScopesOutput) ToNetworkManagerPropertiesResponseNetworkManagerScopesPtrOutputWithContext(ctx context.Context) NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkManagerPropertiesResponseNetworkManagerScopes) *NetworkManagerPropertiesResponseNetworkManagerScopes {
		return &v
	}).(NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput)
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





type NetworkManagerSecurityGroupItemResponseInput interface {
	pulumi.Input

	ToNetworkManagerSecurityGroupItemResponseOutput() NetworkManagerSecurityGroupItemResponseOutput
	ToNetworkManagerSecurityGroupItemResponseOutputWithContext(context.Context) NetworkManagerSecurityGroupItemResponseOutput
}

type NetworkManagerSecurityGroupItemResponseArgs struct {
	NetworkGroupId pulumi.StringPtrInput `pulumi:"networkGroupId"`
}

func (NetworkManagerSecurityGroupItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkManagerSecurityGroupItemResponse)(nil)).Elem()
}

func (i NetworkManagerSecurityGroupItemResponseArgs) ToNetworkManagerSecurityGroupItemResponseOutput() NetworkManagerSecurityGroupItemResponseOutput {
	return i.ToNetworkManagerSecurityGroupItemResponseOutputWithContext(context.Background())
}

func (i NetworkManagerSecurityGroupItemResponseArgs) ToNetworkManagerSecurityGroupItemResponseOutputWithContext(ctx context.Context) NetworkManagerSecurityGroupItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerSecurityGroupItemResponseOutput)
}





type NetworkManagerSecurityGroupItemResponseArrayInput interface {
	pulumi.Input

	ToNetworkManagerSecurityGroupItemResponseArrayOutput() NetworkManagerSecurityGroupItemResponseArrayOutput
	ToNetworkManagerSecurityGroupItemResponseArrayOutputWithContext(context.Context) NetworkManagerSecurityGroupItemResponseArrayOutput
}

type NetworkManagerSecurityGroupItemResponseArray []NetworkManagerSecurityGroupItemResponseInput

func (NetworkManagerSecurityGroupItemResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkManagerSecurityGroupItemResponse)(nil)).Elem()
}

func (i NetworkManagerSecurityGroupItemResponseArray) ToNetworkManagerSecurityGroupItemResponseArrayOutput() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return i.ToNetworkManagerSecurityGroupItemResponseArrayOutputWithContext(context.Background())
}

func (i NetworkManagerSecurityGroupItemResponseArray) ToNetworkManagerSecurityGroupItemResponseArrayOutputWithContext(ctx context.Context) NetworkManagerSecurityGroupItemResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerSecurityGroupItemResponseArrayOutput)
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

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActiveConnectivityConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ActiveConnectivityConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(ActiveDefaultSecurityAdminRuleResponseOutput{})
	pulumi.RegisterOutputType(ActiveDefaultSecurityUserRuleResponseOutput{})
	pulumi.RegisterOutputType(ActiveSecurityAdminRuleResponseOutput{})
	pulumi.RegisterOutputType(ActiveSecurityUserRuleResponseOutput{})
	pulumi.RegisterOutputType(AddressPrefixItemOutput{})
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
	pulumi.RegisterOutputType(EffectiveDefaultSecurityAdminRuleResponseOutput{})
	pulumi.RegisterOutputType(EffectiveSecurityAdminRuleResponseOutput{})
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
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}

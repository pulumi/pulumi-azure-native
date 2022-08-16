


package deviceupdate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectionDetailsResponse struct {
	GroupId          string `pulumi:"groupId"`
	Id               string `pulumi:"id"`
	LinkIdentifier   string `pulumi:"linkIdentifier"`
	MemberName       string `pulumi:"memberName"`
	PrivateIpAddress string `pulumi:"privateIpAddress"`
}

type ConnectionDetailsResponseOutput struct{ *pulumi.OutputState }

func (ConnectionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionDetailsResponse)(nil)).Elem()
}

func (o ConnectionDetailsResponseOutput) ToConnectionDetailsResponseOutput() ConnectionDetailsResponseOutput {
	return o
}

func (o ConnectionDetailsResponseOutput) ToConnectionDetailsResponseOutputWithContext(ctx context.Context) ConnectionDetailsResponseOutput {
	return o
}

func (o ConnectionDetailsResponseOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionDetailsResponse) string { return v.GroupId }).(pulumi.StringOutput)
}

func (o ConnectionDetailsResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionDetailsResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ConnectionDetailsResponseOutput) LinkIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionDetailsResponse) string { return v.LinkIdentifier }).(pulumi.StringOutput)
}

func (o ConnectionDetailsResponseOutput) MemberName() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionDetailsResponse) string { return v.MemberName }).(pulumi.StringOutput)
}

func (o ConnectionDetailsResponseOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionDetailsResponse) string { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

type ConnectionDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (ConnectionDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionDetailsResponse)(nil)).Elem()
}

func (o ConnectionDetailsResponseArrayOutput) ToConnectionDetailsResponseArrayOutput() ConnectionDetailsResponseArrayOutput {
	return o
}

func (o ConnectionDetailsResponseArrayOutput) ToConnectionDetailsResponseArrayOutputWithContext(ctx context.Context) ConnectionDetailsResponseArrayOutput {
	return o
}

func (o ConnectionDetailsResponseArrayOutput) Index(i pulumi.IntInput) ConnectionDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectionDetailsResponse {
		return vs[0].([]ConnectionDetailsResponse)[vs[1].(int)]
	}).(ConnectionDetailsResponseOutput)
}

type DiagnosticStorageProperties struct {
	AuthenticationType string  `pulumi:"authenticationType"`
	ConnectionString   *string `pulumi:"connectionString"`
	ResourceId         string  `pulumi:"resourceId"`
}





type DiagnosticStoragePropertiesInput interface {
	pulumi.Input

	ToDiagnosticStoragePropertiesOutput() DiagnosticStoragePropertiesOutput
	ToDiagnosticStoragePropertiesOutputWithContext(context.Context) DiagnosticStoragePropertiesOutput
}

type DiagnosticStoragePropertiesArgs struct {
	AuthenticationType pulumi.StringInput    `pulumi:"authenticationType"`
	ConnectionString   pulumi.StringPtrInput `pulumi:"connectionString"`
	ResourceId         pulumi.StringInput    `pulumi:"resourceId"`
}

func (DiagnosticStoragePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticStorageProperties)(nil)).Elem()
}

func (i DiagnosticStoragePropertiesArgs) ToDiagnosticStoragePropertiesOutput() DiagnosticStoragePropertiesOutput {
	return i.ToDiagnosticStoragePropertiesOutputWithContext(context.Background())
}

func (i DiagnosticStoragePropertiesArgs) ToDiagnosticStoragePropertiesOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticStoragePropertiesOutput)
}

func (i DiagnosticStoragePropertiesArgs) ToDiagnosticStoragePropertiesPtrOutput() DiagnosticStoragePropertiesPtrOutput {
	return i.ToDiagnosticStoragePropertiesPtrOutputWithContext(context.Background())
}

func (i DiagnosticStoragePropertiesArgs) ToDiagnosticStoragePropertiesPtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticStoragePropertiesOutput).ToDiagnosticStoragePropertiesPtrOutputWithContext(ctx)
}









type DiagnosticStoragePropertiesPtrInput interface {
	pulumi.Input

	ToDiagnosticStoragePropertiesPtrOutput() DiagnosticStoragePropertiesPtrOutput
	ToDiagnosticStoragePropertiesPtrOutputWithContext(context.Context) DiagnosticStoragePropertiesPtrOutput
}

type diagnosticStoragePropertiesPtrType DiagnosticStoragePropertiesArgs

func DiagnosticStoragePropertiesPtr(v *DiagnosticStoragePropertiesArgs) DiagnosticStoragePropertiesPtrInput {
	return (*diagnosticStoragePropertiesPtrType)(v)
}

func (*diagnosticStoragePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticStorageProperties)(nil)).Elem()
}

func (i *diagnosticStoragePropertiesPtrType) ToDiagnosticStoragePropertiesPtrOutput() DiagnosticStoragePropertiesPtrOutput {
	return i.ToDiagnosticStoragePropertiesPtrOutputWithContext(context.Background())
}

func (i *diagnosticStoragePropertiesPtrType) ToDiagnosticStoragePropertiesPtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticStoragePropertiesPtrOutput)
}

type DiagnosticStoragePropertiesOutput struct{ *pulumi.OutputState }

func (DiagnosticStoragePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticStorageProperties)(nil)).Elem()
}

func (o DiagnosticStoragePropertiesOutput) ToDiagnosticStoragePropertiesOutput() DiagnosticStoragePropertiesOutput {
	return o
}

func (o DiagnosticStoragePropertiesOutput) ToDiagnosticStoragePropertiesOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesOutput {
	return o
}

func (o DiagnosticStoragePropertiesOutput) ToDiagnosticStoragePropertiesPtrOutput() DiagnosticStoragePropertiesPtrOutput {
	return o.ToDiagnosticStoragePropertiesPtrOutputWithContext(context.Background())
}

func (o DiagnosticStoragePropertiesOutput) ToDiagnosticStoragePropertiesPtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticStorageProperties) *DiagnosticStorageProperties {
		return &v
	}).(DiagnosticStoragePropertiesPtrOutput)
}

func (o DiagnosticStoragePropertiesOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticStorageProperties) string { return v.AuthenticationType }).(pulumi.StringOutput)
}

func (o DiagnosticStoragePropertiesOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiagnosticStorageProperties) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o DiagnosticStoragePropertiesOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticStorageProperties) string { return v.ResourceId }).(pulumi.StringOutput)
}

type DiagnosticStoragePropertiesPtrOutput struct{ *pulumi.OutputState }

func (DiagnosticStoragePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticStorageProperties)(nil)).Elem()
}

func (o DiagnosticStoragePropertiesPtrOutput) ToDiagnosticStoragePropertiesPtrOutput() DiagnosticStoragePropertiesPtrOutput {
	return o
}

func (o DiagnosticStoragePropertiesPtrOutput) ToDiagnosticStoragePropertiesPtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesPtrOutput {
	return o
}

func (o DiagnosticStoragePropertiesPtrOutput) Elem() DiagnosticStoragePropertiesOutput {
	return o.ApplyT(func(v *DiagnosticStorageProperties) DiagnosticStorageProperties {
		if v != nil {
			return *v
		}
		var ret DiagnosticStorageProperties
		return ret
	}).(DiagnosticStoragePropertiesOutput)
}

func (o DiagnosticStoragePropertiesPtrOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticStorageProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AuthenticationType
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticStoragePropertiesPtrOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticStorageProperties) *string {
		if v == nil {
			return nil
		}
		return v.ConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticStoragePropertiesPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticStorageProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type DiagnosticStoragePropertiesResponse struct {
	AuthenticationType string  `pulumi:"authenticationType"`
	ConnectionString   *string `pulumi:"connectionString"`
	ResourceId         string  `pulumi:"resourceId"`
}

type DiagnosticStoragePropertiesResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticStoragePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticStoragePropertiesResponse)(nil)).Elem()
}

func (o DiagnosticStoragePropertiesResponseOutput) ToDiagnosticStoragePropertiesResponseOutput() DiagnosticStoragePropertiesResponseOutput {
	return o
}

func (o DiagnosticStoragePropertiesResponseOutput) ToDiagnosticStoragePropertiesResponseOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesResponseOutput {
	return o
}

func (o DiagnosticStoragePropertiesResponseOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticStoragePropertiesResponse) string { return v.AuthenticationType }).(pulumi.StringOutput)
}

func (o DiagnosticStoragePropertiesResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiagnosticStoragePropertiesResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o DiagnosticStoragePropertiesResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticStoragePropertiesResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

type DiagnosticStoragePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticStoragePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticStoragePropertiesResponse)(nil)).Elem()
}

func (o DiagnosticStoragePropertiesResponsePtrOutput) ToDiagnosticStoragePropertiesResponsePtrOutput() DiagnosticStoragePropertiesResponsePtrOutput {
	return o
}

func (o DiagnosticStoragePropertiesResponsePtrOutput) ToDiagnosticStoragePropertiesResponsePtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesResponsePtrOutput {
	return o
}

func (o DiagnosticStoragePropertiesResponsePtrOutput) Elem() DiagnosticStoragePropertiesResponseOutput {
	return o.ApplyT(func(v *DiagnosticStoragePropertiesResponse) DiagnosticStoragePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticStoragePropertiesResponse
		return ret
	}).(DiagnosticStoragePropertiesResponseOutput)
}

func (o DiagnosticStoragePropertiesResponsePtrOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticStoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AuthenticationType
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticStoragePropertiesResponsePtrOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticStoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticStoragePropertiesResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticStoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type GroupConnectivityInformation struct {
	CustomerVisibleFqdns        []string `pulumi:"customerVisibleFqdns"`
	PrivateLinkServiceArmRegion *string  `pulumi:"privateLinkServiceArmRegion"`
	RedirectMapId               *string  `pulumi:"redirectMapId"`
}





type GroupConnectivityInformationInput interface {
	pulumi.Input

	ToGroupConnectivityInformationOutput() GroupConnectivityInformationOutput
	ToGroupConnectivityInformationOutputWithContext(context.Context) GroupConnectivityInformationOutput
}

type GroupConnectivityInformationArgs struct {
	CustomerVisibleFqdns        pulumi.StringArrayInput `pulumi:"customerVisibleFqdns"`
	PrivateLinkServiceArmRegion pulumi.StringPtrInput   `pulumi:"privateLinkServiceArmRegion"`
	RedirectMapId               pulumi.StringPtrInput   `pulumi:"redirectMapId"`
}

func (GroupConnectivityInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupConnectivityInformation)(nil)).Elem()
}

func (i GroupConnectivityInformationArgs) ToGroupConnectivityInformationOutput() GroupConnectivityInformationOutput {
	return i.ToGroupConnectivityInformationOutputWithContext(context.Background())
}

func (i GroupConnectivityInformationArgs) ToGroupConnectivityInformationOutputWithContext(ctx context.Context) GroupConnectivityInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupConnectivityInformationOutput)
}





type GroupConnectivityInformationArrayInput interface {
	pulumi.Input

	ToGroupConnectivityInformationArrayOutput() GroupConnectivityInformationArrayOutput
	ToGroupConnectivityInformationArrayOutputWithContext(context.Context) GroupConnectivityInformationArrayOutput
}

type GroupConnectivityInformationArray []GroupConnectivityInformationInput

func (GroupConnectivityInformationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupConnectivityInformation)(nil)).Elem()
}

func (i GroupConnectivityInformationArray) ToGroupConnectivityInformationArrayOutput() GroupConnectivityInformationArrayOutput {
	return i.ToGroupConnectivityInformationArrayOutputWithContext(context.Background())
}

func (i GroupConnectivityInformationArray) ToGroupConnectivityInformationArrayOutputWithContext(ctx context.Context) GroupConnectivityInformationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupConnectivityInformationArrayOutput)
}

type GroupConnectivityInformationOutput struct{ *pulumi.OutputState }

func (GroupConnectivityInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupConnectivityInformation)(nil)).Elem()
}

func (o GroupConnectivityInformationOutput) ToGroupConnectivityInformationOutput() GroupConnectivityInformationOutput {
	return o
}

func (o GroupConnectivityInformationOutput) ToGroupConnectivityInformationOutputWithContext(ctx context.Context) GroupConnectivityInformationOutput {
	return o
}

func (o GroupConnectivityInformationOutput) CustomerVisibleFqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupConnectivityInformation) []string { return v.CustomerVisibleFqdns }).(pulumi.StringArrayOutput)
}

func (o GroupConnectivityInformationOutput) PrivateLinkServiceArmRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupConnectivityInformation) *string { return v.PrivateLinkServiceArmRegion }).(pulumi.StringPtrOutput)
}

func (o GroupConnectivityInformationOutput) RedirectMapId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupConnectivityInformation) *string { return v.RedirectMapId }).(pulumi.StringPtrOutput)
}

type GroupConnectivityInformationArrayOutput struct{ *pulumi.OutputState }

func (GroupConnectivityInformationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupConnectivityInformation)(nil)).Elem()
}

func (o GroupConnectivityInformationArrayOutput) ToGroupConnectivityInformationArrayOutput() GroupConnectivityInformationArrayOutput {
	return o
}

func (o GroupConnectivityInformationArrayOutput) ToGroupConnectivityInformationArrayOutputWithContext(ctx context.Context) GroupConnectivityInformationArrayOutput {
	return o
}

func (o GroupConnectivityInformationArrayOutput) Index(i pulumi.IntInput) GroupConnectivityInformationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupConnectivityInformation {
		return vs[0].([]GroupConnectivityInformation)[vs[1].(int)]
	}).(GroupConnectivityInformationOutput)
}

type GroupConnectivityInformationResponse struct {
	CustomerVisibleFqdns        []string `pulumi:"customerVisibleFqdns"`
	GroupId                     string   `pulumi:"groupId"`
	InternalFqdn                string   `pulumi:"internalFqdn"`
	MemberName                  string   `pulumi:"memberName"`
	PrivateLinkServiceArmRegion *string  `pulumi:"privateLinkServiceArmRegion"`
	RedirectMapId               *string  `pulumi:"redirectMapId"`
}

type GroupConnectivityInformationResponseOutput struct{ *pulumi.OutputState }

func (GroupConnectivityInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupConnectivityInformationResponse)(nil)).Elem()
}

func (o GroupConnectivityInformationResponseOutput) ToGroupConnectivityInformationResponseOutput() GroupConnectivityInformationResponseOutput {
	return o
}

func (o GroupConnectivityInformationResponseOutput) ToGroupConnectivityInformationResponseOutputWithContext(ctx context.Context) GroupConnectivityInformationResponseOutput {
	return o
}

func (o GroupConnectivityInformationResponseOutput) CustomerVisibleFqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupConnectivityInformationResponse) []string { return v.CustomerVisibleFqdns }).(pulumi.StringArrayOutput)
}

func (o GroupConnectivityInformationResponseOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GroupConnectivityInformationResponse) string { return v.GroupId }).(pulumi.StringOutput)
}

func (o GroupConnectivityInformationResponseOutput) InternalFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v GroupConnectivityInformationResponse) string { return v.InternalFqdn }).(pulumi.StringOutput)
}

func (o GroupConnectivityInformationResponseOutput) MemberName() pulumi.StringOutput {
	return o.ApplyT(func(v GroupConnectivityInformationResponse) string { return v.MemberName }).(pulumi.StringOutput)
}

func (o GroupConnectivityInformationResponseOutput) PrivateLinkServiceArmRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupConnectivityInformationResponse) *string { return v.PrivateLinkServiceArmRegion }).(pulumi.StringPtrOutput)
}

func (o GroupConnectivityInformationResponseOutput) RedirectMapId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupConnectivityInformationResponse) *string { return v.RedirectMapId }).(pulumi.StringPtrOutput)
}

type GroupConnectivityInformationResponseArrayOutput struct{ *pulumi.OutputState }

func (GroupConnectivityInformationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupConnectivityInformationResponse)(nil)).Elem()
}

func (o GroupConnectivityInformationResponseArrayOutput) ToGroupConnectivityInformationResponseArrayOutput() GroupConnectivityInformationResponseArrayOutput {
	return o
}

func (o GroupConnectivityInformationResponseArrayOutput) ToGroupConnectivityInformationResponseArrayOutputWithContext(ctx context.Context) GroupConnectivityInformationResponseArrayOutput {
	return o
}

func (o GroupConnectivityInformationResponseArrayOutput) Index(i pulumi.IntInput) GroupConnectivityInformationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupConnectivityInformationResponse {
		return vs[0].([]GroupConnectivityInformationResponse)[vs[1].(int)]
	}).(GroupConnectivityInformationResponseOutput)
}

type IotHubSettings struct {
	EventHubConnectionString *string `pulumi:"eventHubConnectionString"`
	IoTHubConnectionString   *string `pulumi:"ioTHubConnectionString"`
	ResourceId               string  `pulumi:"resourceId"`
}





type IotHubSettingsInput interface {
	pulumi.Input

	ToIotHubSettingsOutput() IotHubSettingsOutput
	ToIotHubSettingsOutputWithContext(context.Context) IotHubSettingsOutput
}

type IotHubSettingsArgs struct {
	EventHubConnectionString pulumi.StringPtrInput `pulumi:"eventHubConnectionString"`
	IoTHubConnectionString   pulumi.StringPtrInput `pulumi:"ioTHubConnectionString"`
	ResourceId               pulumi.StringInput    `pulumi:"resourceId"`
}

func (IotHubSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSettings)(nil)).Elem()
}

func (i IotHubSettingsArgs) ToIotHubSettingsOutput() IotHubSettingsOutput {
	return i.ToIotHubSettingsOutputWithContext(context.Background())
}

func (i IotHubSettingsArgs) ToIotHubSettingsOutputWithContext(ctx context.Context) IotHubSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSettingsOutput)
}





type IotHubSettingsArrayInput interface {
	pulumi.Input

	ToIotHubSettingsArrayOutput() IotHubSettingsArrayOutput
	ToIotHubSettingsArrayOutputWithContext(context.Context) IotHubSettingsArrayOutput
}

type IotHubSettingsArray []IotHubSettingsInput

func (IotHubSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubSettings)(nil)).Elem()
}

func (i IotHubSettingsArray) ToIotHubSettingsArrayOutput() IotHubSettingsArrayOutput {
	return i.ToIotHubSettingsArrayOutputWithContext(context.Background())
}

func (i IotHubSettingsArray) ToIotHubSettingsArrayOutputWithContext(ctx context.Context) IotHubSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSettingsArrayOutput)
}

type IotHubSettingsOutput struct{ *pulumi.OutputState }

func (IotHubSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSettings)(nil)).Elem()
}

func (o IotHubSettingsOutput) ToIotHubSettingsOutput() IotHubSettingsOutput {
	return o
}

func (o IotHubSettingsOutput) ToIotHubSettingsOutputWithContext(ctx context.Context) IotHubSettingsOutput {
	return o
}

func (o IotHubSettingsOutput) EventHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubSettings) *string { return v.EventHubConnectionString }).(pulumi.StringPtrOutput)
}

func (o IotHubSettingsOutput) IoTHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubSettings) *string { return v.IoTHubConnectionString }).(pulumi.StringPtrOutput)
}

func (o IotHubSettingsOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubSettings) string { return v.ResourceId }).(pulumi.StringOutput)
}

type IotHubSettingsArrayOutput struct{ *pulumi.OutputState }

func (IotHubSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubSettings)(nil)).Elem()
}

func (o IotHubSettingsArrayOutput) ToIotHubSettingsArrayOutput() IotHubSettingsArrayOutput {
	return o
}

func (o IotHubSettingsArrayOutput) ToIotHubSettingsArrayOutputWithContext(ctx context.Context) IotHubSettingsArrayOutput {
	return o
}

func (o IotHubSettingsArrayOutput) Index(i pulumi.IntInput) IotHubSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IotHubSettings {
		return vs[0].([]IotHubSettings)[vs[1].(int)]
	}).(IotHubSettingsOutput)
}

type IotHubSettingsResponse struct {
	EventHubConnectionString *string `pulumi:"eventHubConnectionString"`
	IoTHubConnectionString   *string `pulumi:"ioTHubConnectionString"`
	ResourceId               string  `pulumi:"resourceId"`
}

type IotHubSettingsResponseOutput struct{ *pulumi.OutputState }

func (IotHubSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSettingsResponse)(nil)).Elem()
}

func (o IotHubSettingsResponseOutput) ToIotHubSettingsResponseOutput() IotHubSettingsResponseOutput {
	return o
}

func (o IotHubSettingsResponseOutput) ToIotHubSettingsResponseOutputWithContext(ctx context.Context) IotHubSettingsResponseOutput {
	return o
}

func (o IotHubSettingsResponseOutput) EventHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubSettingsResponse) *string { return v.EventHubConnectionString }).(pulumi.StringPtrOutput)
}

func (o IotHubSettingsResponseOutput) IoTHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubSettingsResponse) *string { return v.IoTHubConnectionString }).(pulumi.StringPtrOutput)
}

func (o IotHubSettingsResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubSettingsResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

type IotHubSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (IotHubSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubSettingsResponse)(nil)).Elem()
}

func (o IotHubSettingsResponseArrayOutput) ToIotHubSettingsResponseArrayOutput() IotHubSettingsResponseArrayOutput {
	return o
}

func (o IotHubSettingsResponseArrayOutput) ToIotHubSettingsResponseArrayOutputWithContext(ctx context.Context) IotHubSettingsResponseArrayOutput {
	return o
}

func (o IotHubSettingsResponseArrayOutput) Index(i pulumi.IntInput) IotHubSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IotHubSettingsResponse {
		return vs[0].([]IotHubSettingsResponse)[vs[1].(int)]
	}).(IotHubSettingsResponseOutput)
}

type ManagedServiceIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

func (o ManagedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   string                                  `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type PrivateEndpointConnectionType struct {
	GroupIds                          []string                          `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	GroupIds                          pulumi.StringArrayInput                `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}





type PrivateEndpointConnectionTypeArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput
	ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Context) PrivateEndpointConnectionTypeArrayOutput
}

type PrivateEndpointConnectionTypeArray []PrivateEndpointConnectionTypeInput

func (PrivateEndpointConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return i.ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeArrayOutput)
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionTypeOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateEndpointConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionType {
		return vs[0].([]PrivateEndpointConnectionType)[vs[1].(int)]
	}).(PrivateEndpointConnectionTypeOutput)
}

type PrivateEndpointConnectionResponse struct {
	GroupIds                          []string                                  `pulumi:"groupIds"`
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnection struct {
	GroupIds       []string `pulumi:"groupIds"`
	Name           *string  `pulumi:"name"`
	RequestMessage *string  `pulumi:"requestMessage"`
}





type PrivateLinkServiceConnectionInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionOutput() PrivateLinkServiceConnectionOutput
	ToPrivateLinkServiceConnectionOutputWithContext(context.Context) PrivateLinkServiceConnectionOutput
}

type PrivateLinkServiceConnectionArgs struct {
	GroupIds       pulumi.StringArrayInput `pulumi:"groupIds"`
	Name           pulumi.StringPtrInput   `pulumi:"name"`
	RequestMessage pulumi.StringPtrInput   `pulumi:"requestMessage"`
}

func (PrivateLinkServiceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnection)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionArgs) ToPrivateLinkServiceConnectionOutput() PrivateLinkServiceConnectionOutput {
	return i.ToPrivateLinkServiceConnectionOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionArgs) ToPrivateLinkServiceConnectionOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionOutput)
}





type PrivateLinkServiceConnectionArrayInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionArrayOutput() PrivateLinkServiceConnectionArrayOutput
	ToPrivateLinkServiceConnectionArrayOutputWithContext(context.Context) PrivateLinkServiceConnectionArrayOutput
}

type PrivateLinkServiceConnectionArray []PrivateLinkServiceConnectionInput

func (PrivateLinkServiceConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkServiceConnection)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionArray) ToPrivateLinkServiceConnectionArrayOutput() PrivateLinkServiceConnectionArrayOutput {
	return i.ToPrivateLinkServiceConnectionArrayOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionArray) ToPrivateLinkServiceConnectionArrayOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionArrayOutput)
}

type PrivateLinkServiceConnectionOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnection)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionOutput) ToPrivateLinkServiceConnectionOutput() PrivateLinkServiceConnectionOutput {
	return o
}

func (o PrivateLinkServiceConnectionOutput) ToPrivateLinkServiceConnectionOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionOutput {
	return o
}

func (o PrivateLinkServiceConnectionOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnection) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateLinkServiceConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnection) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnection) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkServiceConnection)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionArrayOutput) ToPrivateLinkServiceConnectionArrayOutput() PrivateLinkServiceConnectionArrayOutput {
	return o
}

func (o PrivateLinkServiceConnectionArrayOutput) ToPrivateLinkServiceConnectionArrayOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionArrayOutput {
	return o
}

func (o PrivateLinkServiceConnectionArrayOutput) Index(i pulumi.IntInput) PrivateLinkServiceConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkServiceConnection {
		return vs[0].([]PrivateLinkServiceConnection)[vs[1].(int)]
	}).(PrivateLinkServiceConnectionOutput)
}

type PrivateLinkServiceConnectionResponse struct {
	GroupIds       []string `pulumi:"groupIds"`
	Name           *string  `pulumi:"name"`
	RequestMessage *string  `pulumi:"requestMessage"`
}

type PrivateLinkServiceConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionResponseOutput) ToPrivateLinkServiceConnectionResponseOutput() PrivateLinkServiceConnectionResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionResponseOutput) ToPrivateLinkServiceConnectionResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateLinkServiceConnectionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionResponseOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionResponse) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkServiceConnectionResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionResponseArrayOutput) ToPrivateLinkServiceConnectionResponseArrayOutput() PrivateLinkServiceConnectionResponseArrayOutput {
	return o
}

func (o PrivateLinkServiceConnectionResponseArrayOutput) ToPrivateLinkServiceConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionResponseArrayOutput {
	return o
}

func (o PrivateLinkServiceConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateLinkServiceConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkServiceConnectionResponse {
		return vs[0].([]PrivateLinkServiceConnectionResponse)[vs[1].(int)]
	}).(PrivateLinkServiceConnectionResponseOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceProxy struct {
	GroupConnectivityInformation            []GroupConnectivityInformation     `pulumi:"groupConnectivityInformation"`
	Id                                      *string                            `pulumi:"id"`
	RemotePrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"remotePrivateLinkServiceConnectionState"`
}





type PrivateLinkServiceProxyInput interface {
	pulumi.Input

	ToPrivateLinkServiceProxyOutput() PrivateLinkServiceProxyOutput
	ToPrivateLinkServiceProxyOutputWithContext(context.Context) PrivateLinkServiceProxyOutput
}

type PrivateLinkServiceProxyArgs struct {
	GroupConnectivityInformation            GroupConnectivityInformationArrayInput    `pulumi:"groupConnectivityInformation"`
	Id                                      pulumi.StringPtrInput                     `pulumi:"id"`
	RemotePrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput `pulumi:"remotePrivateLinkServiceConnectionState"`
}

func (PrivateLinkServiceProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceProxy)(nil)).Elem()
}

func (i PrivateLinkServiceProxyArgs) ToPrivateLinkServiceProxyOutput() PrivateLinkServiceProxyOutput {
	return i.ToPrivateLinkServiceProxyOutputWithContext(context.Background())
}

func (i PrivateLinkServiceProxyArgs) ToPrivateLinkServiceProxyOutputWithContext(ctx context.Context) PrivateLinkServiceProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceProxyOutput)
}





type PrivateLinkServiceProxyArrayInput interface {
	pulumi.Input

	ToPrivateLinkServiceProxyArrayOutput() PrivateLinkServiceProxyArrayOutput
	ToPrivateLinkServiceProxyArrayOutputWithContext(context.Context) PrivateLinkServiceProxyArrayOutput
}

type PrivateLinkServiceProxyArray []PrivateLinkServiceProxyInput

func (PrivateLinkServiceProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkServiceProxy)(nil)).Elem()
}

func (i PrivateLinkServiceProxyArray) ToPrivateLinkServiceProxyArrayOutput() PrivateLinkServiceProxyArrayOutput {
	return i.ToPrivateLinkServiceProxyArrayOutputWithContext(context.Background())
}

func (i PrivateLinkServiceProxyArray) ToPrivateLinkServiceProxyArrayOutputWithContext(ctx context.Context) PrivateLinkServiceProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceProxyArrayOutput)
}

type PrivateLinkServiceProxyOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceProxy)(nil)).Elem()
}

func (o PrivateLinkServiceProxyOutput) ToPrivateLinkServiceProxyOutput() PrivateLinkServiceProxyOutput {
	return o
}

func (o PrivateLinkServiceProxyOutput) ToPrivateLinkServiceProxyOutputWithContext(ctx context.Context) PrivateLinkServiceProxyOutput {
	return o
}

func (o PrivateLinkServiceProxyOutput) GroupConnectivityInformation() GroupConnectivityInformationArrayOutput {
	return o.ApplyT(func(v PrivateLinkServiceProxy) []GroupConnectivityInformation { return v.GroupConnectivityInformation }).(GroupConnectivityInformationArrayOutput)
}

func (o PrivateLinkServiceProxyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceProxy) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceProxyOutput) RemotePrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceProxy) *PrivateLinkServiceConnectionState {
		return v.RemotePrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceProxyArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkServiceProxy)(nil)).Elem()
}

func (o PrivateLinkServiceProxyArrayOutput) ToPrivateLinkServiceProxyArrayOutput() PrivateLinkServiceProxyArrayOutput {
	return o
}

func (o PrivateLinkServiceProxyArrayOutput) ToPrivateLinkServiceProxyArrayOutputWithContext(ctx context.Context) PrivateLinkServiceProxyArrayOutput {
	return o
}

func (o PrivateLinkServiceProxyArrayOutput) Index(i pulumi.IntInput) PrivateLinkServiceProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkServiceProxy {
		return vs[0].([]PrivateLinkServiceProxy)[vs[1].(int)]
	}).(PrivateLinkServiceProxyOutput)
}

type PrivateLinkServiceProxyResponse struct {
	GroupConnectivityInformation            []GroupConnectivityInformationResponse                          `pulumi:"groupConnectivityInformation"`
	Id                                      *string                                                         `pulumi:"id"`
	RemotePrivateEndpointConnection         *PrivateLinkServiceProxyResponseRemotePrivateEndpointConnection `pulumi:"remotePrivateEndpointConnection"`
	RemotePrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse                      `pulumi:"remotePrivateLinkServiceConnectionState"`
}

type PrivateLinkServiceProxyResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceProxyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceProxyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceProxyResponseOutput) ToPrivateLinkServiceProxyResponseOutput() PrivateLinkServiceProxyResponseOutput {
	return o
}

func (o PrivateLinkServiceProxyResponseOutput) ToPrivateLinkServiceProxyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceProxyResponseOutput {
	return o
}

func (o PrivateLinkServiceProxyResponseOutput) GroupConnectivityInformation() GroupConnectivityInformationResponseArrayOutput {
	return o.ApplyT(func(v PrivateLinkServiceProxyResponse) []GroupConnectivityInformationResponse {
		return v.GroupConnectivityInformation
	}).(GroupConnectivityInformationResponseArrayOutput)
}

func (o PrivateLinkServiceProxyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceProxyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceProxyResponseOutput) RemotePrivateEndpointConnection() PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceProxyResponse) *PrivateLinkServiceProxyResponseRemotePrivateEndpointConnection {
		return v.RemotePrivateEndpointConnection
	}).(PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutput)
}

func (o PrivateLinkServiceProxyResponseOutput) RemotePrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceProxyResponse) *PrivateLinkServiceConnectionStateResponse {
		return v.RemotePrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceProxyResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceProxyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkServiceProxyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceProxyResponseArrayOutput) ToPrivateLinkServiceProxyResponseArrayOutput() PrivateLinkServiceProxyResponseArrayOutput {
	return o
}

func (o PrivateLinkServiceProxyResponseArrayOutput) ToPrivateLinkServiceProxyResponseArrayOutputWithContext(ctx context.Context) PrivateLinkServiceProxyResponseArrayOutput {
	return o
}

func (o PrivateLinkServiceProxyResponseArrayOutput) Index(i pulumi.IntInput) PrivateLinkServiceProxyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkServiceProxyResponse {
		return vs[0].([]PrivateLinkServiceProxyResponse)[vs[1].(int)]
	}).(PrivateLinkServiceProxyResponseOutput)
}

type PrivateLinkServiceProxyResponseRemotePrivateEndpointConnection struct {
	Id string `pulumi:"id"`
}

type PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceProxyResponseRemotePrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionOutput) ToPrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionOutput() PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionOutput {
	return o
}

func (o PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionOutput) ToPrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionOutput {
	return o
}

func (o PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceProxyResponseRemotePrivateEndpointConnection) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceProxyResponseRemotePrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutput) ToPrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutput() PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutput {
	return o
}

func (o PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutput) ToPrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutputWithContext(ctx context.Context) PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutput {
	return o
}

func (o PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutput) Elem() PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionOutput {
	return o.ApplyT(func(v *PrivateLinkServiceProxyResponseRemotePrivateEndpointConnection) PrivateLinkServiceProxyResponseRemotePrivateEndpointConnection {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceProxyResponseRemotePrivateEndpointConnection
		return ret
	}).(PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionOutput)
}

func (o PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceProxyResponseRemotePrivateEndpointConnection) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type RemotePrivateEndpoint struct {
	Id                                  *string                        `pulumi:"id"`
	ImmutableResourceId                 *string                        `pulumi:"immutableResourceId"`
	ImmutableSubscriptionId             *string                        `pulumi:"immutableSubscriptionId"`
	Location                            *string                        `pulumi:"location"`
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnection `pulumi:"manualPrivateLinkServiceConnections"`
	PrivateLinkServiceConnections       []PrivateLinkServiceConnection `pulumi:"privateLinkServiceConnections"`
	PrivateLinkServiceProxies           []PrivateLinkServiceProxy      `pulumi:"privateLinkServiceProxies"`
	VnetTrafficTag                      *string                        `pulumi:"vnetTrafficTag"`
}





type RemotePrivateEndpointInput interface {
	pulumi.Input

	ToRemotePrivateEndpointOutput() RemotePrivateEndpointOutput
	ToRemotePrivateEndpointOutputWithContext(context.Context) RemotePrivateEndpointOutput
}

type RemotePrivateEndpointArgs struct {
	Id                                  pulumi.StringPtrInput                  `pulumi:"id"`
	ImmutableResourceId                 pulumi.StringPtrInput                  `pulumi:"immutableResourceId"`
	ImmutableSubscriptionId             pulumi.StringPtrInput                  `pulumi:"immutableSubscriptionId"`
	Location                            pulumi.StringPtrInput                  `pulumi:"location"`
	ManualPrivateLinkServiceConnections PrivateLinkServiceConnectionArrayInput `pulumi:"manualPrivateLinkServiceConnections"`
	PrivateLinkServiceConnections       PrivateLinkServiceConnectionArrayInput `pulumi:"privateLinkServiceConnections"`
	PrivateLinkServiceProxies           PrivateLinkServiceProxyArrayInput      `pulumi:"privateLinkServiceProxies"`
	VnetTrafficTag                      pulumi.StringPtrInput                  `pulumi:"vnetTrafficTag"`
}

func (RemotePrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemotePrivateEndpoint)(nil)).Elem()
}

func (i RemotePrivateEndpointArgs) ToRemotePrivateEndpointOutput() RemotePrivateEndpointOutput {
	return i.ToRemotePrivateEndpointOutputWithContext(context.Background())
}

func (i RemotePrivateEndpointArgs) ToRemotePrivateEndpointOutputWithContext(ctx context.Context) RemotePrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemotePrivateEndpointOutput)
}

func (i RemotePrivateEndpointArgs) ToRemotePrivateEndpointPtrOutput() RemotePrivateEndpointPtrOutput {
	return i.ToRemotePrivateEndpointPtrOutputWithContext(context.Background())
}

func (i RemotePrivateEndpointArgs) ToRemotePrivateEndpointPtrOutputWithContext(ctx context.Context) RemotePrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemotePrivateEndpointOutput).ToRemotePrivateEndpointPtrOutputWithContext(ctx)
}









type RemotePrivateEndpointPtrInput interface {
	pulumi.Input

	ToRemotePrivateEndpointPtrOutput() RemotePrivateEndpointPtrOutput
	ToRemotePrivateEndpointPtrOutputWithContext(context.Context) RemotePrivateEndpointPtrOutput
}

type remotePrivateEndpointPtrType RemotePrivateEndpointArgs

func RemotePrivateEndpointPtr(v *RemotePrivateEndpointArgs) RemotePrivateEndpointPtrInput {
	return (*remotePrivateEndpointPtrType)(v)
}

func (*remotePrivateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemotePrivateEndpoint)(nil)).Elem()
}

func (i *remotePrivateEndpointPtrType) ToRemotePrivateEndpointPtrOutput() RemotePrivateEndpointPtrOutput {
	return i.ToRemotePrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *remotePrivateEndpointPtrType) ToRemotePrivateEndpointPtrOutputWithContext(ctx context.Context) RemotePrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemotePrivateEndpointPtrOutput)
}

type RemotePrivateEndpointOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemotePrivateEndpoint)(nil)).Elem()
}

func (o RemotePrivateEndpointOutput) ToRemotePrivateEndpointOutput() RemotePrivateEndpointOutput {
	return o
}

func (o RemotePrivateEndpointOutput) ToRemotePrivateEndpointOutputWithContext(ctx context.Context) RemotePrivateEndpointOutput {
	return o
}

func (o RemotePrivateEndpointOutput) ToRemotePrivateEndpointPtrOutput() RemotePrivateEndpointPtrOutput {
	return o.ToRemotePrivateEndpointPtrOutputWithContext(context.Background())
}

func (o RemotePrivateEndpointOutput) ToRemotePrivateEndpointPtrOutputWithContext(ctx context.Context) RemotePrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemotePrivateEndpoint) *RemotePrivateEndpoint {
		return &v
	}).(RemotePrivateEndpointPtrOutput)
}

func (o RemotePrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointOutput) ImmutableResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpoint) *string { return v.ImmutableResourceId }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointOutput) ImmutableSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpoint) *string { return v.ImmutableSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpoint) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointOutput) ManualPrivateLinkServiceConnections() PrivateLinkServiceConnectionArrayOutput {
	return o.ApplyT(func(v RemotePrivateEndpoint) []PrivateLinkServiceConnection {
		return v.ManualPrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionArrayOutput)
}

func (o RemotePrivateEndpointOutput) PrivateLinkServiceConnections() PrivateLinkServiceConnectionArrayOutput {
	return o.ApplyT(func(v RemotePrivateEndpoint) []PrivateLinkServiceConnection { return v.PrivateLinkServiceConnections }).(PrivateLinkServiceConnectionArrayOutput)
}

func (o RemotePrivateEndpointOutput) PrivateLinkServiceProxies() PrivateLinkServiceProxyArrayOutput {
	return o.ApplyT(func(v RemotePrivateEndpoint) []PrivateLinkServiceProxy { return v.PrivateLinkServiceProxies }).(PrivateLinkServiceProxyArrayOutput)
}

func (o RemotePrivateEndpointOutput) VnetTrafficTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpoint) *string { return v.VnetTrafficTag }).(pulumi.StringPtrOutput)
}

type RemotePrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemotePrivateEndpoint)(nil)).Elem()
}

func (o RemotePrivateEndpointPtrOutput) ToRemotePrivateEndpointPtrOutput() RemotePrivateEndpointPtrOutput {
	return o
}

func (o RemotePrivateEndpointPtrOutput) ToRemotePrivateEndpointPtrOutputWithContext(ctx context.Context) RemotePrivateEndpointPtrOutput {
	return o
}

func (o RemotePrivateEndpointPtrOutput) Elem() RemotePrivateEndpointOutput {
	return o.ApplyT(func(v *RemotePrivateEndpoint) RemotePrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret RemotePrivateEndpoint
		return ret
	}).(RemotePrivateEndpointOutput)
}

func (o RemotePrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointPtrOutput) ImmutableResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.ImmutableResourceId
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointPtrOutput) ImmutableSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.ImmutableSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointPtrOutput) ManualPrivateLinkServiceConnections() PrivateLinkServiceConnectionArrayOutput {
	return o.ApplyT(func(v *RemotePrivateEndpoint) []PrivateLinkServiceConnection {
		if v == nil {
			return nil
		}
		return v.ManualPrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionArrayOutput)
}

func (o RemotePrivateEndpointPtrOutput) PrivateLinkServiceConnections() PrivateLinkServiceConnectionArrayOutput {
	return o.ApplyT(func(v *RemotePrivateEndpoint) []PrivateLinkServiceConnection {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionArrayOutput)
}

func (o RemotePrivateEndpointPtrOutput) PrivateLinkServiceProxies() PrivateLinkServiceProxyArrayOutput {
	return o.ApplyT(func(v *RemotePrivateEndpoint) []PrivateLinkServiceProxy {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceProxies
	}).(PrivateLinkServiceProxyArrayOutput)
}

func (o RemotePrivateEndpointPtrOutput) VnetTrafficTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.VnetTrafficTag
	}).(pulumi.StringPtrOutput)
}

type RemotePrivateEndpointResponse struct {
	ConnectionDetails                   []ConnectionDetailsResponse            `pulumi:"connectionDetails"`
	Id                                  *string                                `pulumi:"id"`
	ImmutableResourceId                 *string                                `pulumi:"immutableResourceId"`
	ImmutableSubscriptionId             *string                                `pulumi:"immutableSubscriptionId"`
	Location                            *string                                `pulumi:"location"`
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnectionResponse `pulumi:"manualPrivateLinkServiceConnections"`
	PrivateLinkServiceConnections       []PrivateLinkServiceConnectionResponse `pulumi:"privateLinkServiceConnections"`
	PrivateLinkServiceProxies           []PrivateLinkServiceProxyResponse      `pulumi:"privateLinkServiceProxies"`
	VnetTrafficTag                      *string                                `pulumi:"vnetTrafficTag"`
}

type RemotePrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemotePrivateEndpointResponse)(nil)).Elem()
}

func (o RemotePrivateEndpointResponseOutput) ToRemotePrivateEndpointResponseOutput() RemotePrivateEndpointResponseOutput {
	return o
}

func (o RemotePrivateEndpointResponseOutput) ToRemotePrivateEndpointResponseOutputWithContext(ctx context.Context) RemotePrivateEndpointResponseOutput {
	return o
}

func (o RemotePrivateEndpointResponseOutput) ConnectionDetails() ConnectionDetailsResponseArrayOutput {
	return o.ApplyT(func(v RemotePrivateEndpointResponse) []ConnectionDetailsResponse { return v.ConnectionDetails }).(ConnectionDetailsResponseArrayOutput)
}

func (o RemotePrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointResponseOutput) ImmutableResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointResponse) *string { return v.ImmutableResourceId }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointResponseOutput) ImmutableSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointResponse) *string { return v.ImmutableSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointResponseOutput) ManualPrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v RemotePrivateEndpointResponse) []PrivateLinkServiceConnectionResponse {
		return v.ManualPrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

func (o RemotePrivateEndpointResponseOutput) PrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v RemotePrivateEndpointResponse) []PrivateLinkServiceConnectionResponse {
		return v.PrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

func (o RemotePrivateEndpointResponseOutput) PrivateLinkServiceProxies() PrivateLinkServiceProxyResponseArrayOutput {
	return o.ApplyT(func(v RemotePrivateEndpointResponse) []PrivateLinkServiceProxyResponse {
		return v.PrivateLinkServiceProxies
	}).(PrivateLinkServiceProxyResponseArrayOutput)
}

func (o RemotePrivateEndpointResponseOutput) VnetTrafficTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointResponse) *string { return v.VnetTrafficTag }).(pulumi.StringPtrOutput)
}

type RemotePrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemotePrivateEndpointResponse)(nil)).Elem()
}

func (o RemotePrivateEndpointResponsePtrOutput) ToRemotePrivateEndpointResponsePtrOutput() RemotePrivateEndpointResponsePtrOutput {
	return o
}

func (o RemotePrivateEndpointResponsePtrOutput) ToRemotePrivateEndpointResponsePtrOutputWithContext(ctx context.Context) RemotePrivateEndpointResponsePtrOutput {
	return o
}

func (o RemotePrivateEndpointResponsePtrOutput) Elem() RemotePrivateEndpointResponseOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointResponse) RemotePrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret RemotePrivateEndpointResponse
		return ret
	}).(RemotePrivateEndpointResponseOutput)
}

func (o RemotePrivateEndpointResponsePtrOutput) ConnectionDetails() ConnectionDetailsResponseArrayOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointResponse) []ConnectionDetailsResponse {
		if v == nil {
			return nil
		}
		return v.ConnectionDetails
	}).(ConnectionDetailsResponseArrayOutput)
}

func (o RemotePrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointResponsePtrOutput) ImmutableResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.ImmutableResourceId
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointResponsePtrOutput) ImmutableSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.ImmutableSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointResponsePtrOutput) ManualPrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointResponse) []PrivateLinkServiceConnectionResponse {
		if v == nil {
			return nil
		}
		return v.ManualPrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

func (o RemotePrivateEndpointResponsePtrOutput) PrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointResponse) []PrivateLinkServiceConnectionResponse {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

func (o RemotePrivateEndpointResponsePtrOutput) PrivateLinkServiceProxies() PrivateLinkServiceProxyResponseArrayOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointResponse) []PrivateLinkServiceProxyResponse {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceProxies
	}).(PrivateLinkServiceProxyResponseArrayOutput)
}

func (o RemotePrivateEndpointResponsePtrOutput) VnetTrafficTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.VnetTrafficTag
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

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionDetailsResponseOutput{})
	pulumi.RegisterOutputType(ConnectionDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticStoragePropertiesOutput{})
	pulumi.RegisterOutputType(DiagnosticStoragePropertiesPtrOutput{})
	pulumi.RegisterOutputType(DiagnosticStoragePropertiesResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticStoragePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(GroupConnectivityInformationOutput{})
	pulumi.RegisterOutputType(GroupConnectivityInformationArrayOutput{})
	pulumi.RegisterOutputType(GroupConnectivityInformationResponseOutput{})
	pulumi.RegisterOutputType(GroupConnectivityInformationResponseArrayOutput{})
	pulumi.RegisterOutputType(IotHubSettingsOutput{})
	pulumi.RegisterOutputType(IotHubSettingsArrayOutput{})
	pulumi.RegisterOutputType(IotHubSettingsResponseOutput{})
	pulumi.RegisterOutputType(IotHubSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceProxyOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceProxyArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceProxyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceProxyResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceProxyResponseRemotePrivateEndpointConnectionPtrOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}

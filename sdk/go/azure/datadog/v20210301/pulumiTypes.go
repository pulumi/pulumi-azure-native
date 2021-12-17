


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatadogApiKeyResponse struct {
	Created   *string `pulumi:"created"`
	CreatedBy *string `pulumi:"createdBy"`
	Key       string  `pulumi:"key"`
	Name      *string `pulumi:"name"`
}

type DatadogHostMetadataResponse struct {
	AgentVersion  *string                       `pulumi:"agentVersion"`
	InstallMethod *DatadogInstallMethodResponse `pulumi:"installMethod"`
	LogsAgent     *DatadogLogsAgentResponse     `pulumi:"logsAgent"`
}

type DatadogHostResponse struct {
	Aliases []string                     `pulumi:"aliases"`
	Apps    []string                     `pulumi:"apps"`
	Meta    *DatadogHostMetadataResponse `pulumi:"meta"`
	Name    *string                      `pulumi:"name"`
}

type DatadogInstallMethodResponse struct {
	InstallerVersion *string `pulumi:"installerVersion"`
	Tool             *string `pulumi:"tool"`
	ToolVersion      *string `pulumi:"toolVersion"`
}

type DatadogLogsAgentResponse struct {
	Transport *string `pulumi:"transport"`
}

type DatadogOrganizationProperties struct {
	ApiKey          *string `pulumi:"apiKey"`
	ApplicationKey  *string `pulumi:"applicationKey"`
	EnterpriseAppId *string `pulumi:"enterpriseAppId"`
	LinkingAuthCode *string `pulumi:"linkingAuthCode"`
	LinkingClientId *string `pulumi:"linkingClientId"`
	RedirectUri     *string `pulumi:"redirectUri"`
}





type DatadogOrganizationPropertiesInput interface {
	pulumi.Input

	ToDatadogOrganizationPropertiesOutput() DatadogOrganizationPropertiesOutput
	ToDatadogOrganizationPropertiesOutputWithContext(context.Context) DatadogOrganizationPropertiesOutput
}

type DatadogOrganizationPropertiesArgs struct {
	ApiKey          pulumi.StringPtrInput `pulumi:"apiKey"`
	ApplicationKey  pulumi.StringPtrInput `pulumi:"applicationKey"`
	EnterpriseAppId pulumi.StringPtrInput `pulumi:"enterpriseAppId"`
	LinkingAuthCode pulumi.StringPtrInput `pulumi:"linkingAuthCode"`
	LinkingClientId pulumi.StringPtrInput `pulumi:"linkingClientId"`
	RedirectUri     pulumi.StringPtrInput `pulumi:"redirectUri"`
}

func (DatadogOrganizationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogOrganizationProperties)(nil)).Elem()
}

func (i DatadogOrganizationPropertiesArgs) ToDatadogOrganizationPropertiesOutput() DatadogOrganizationPropertiesOutput {
	return i.ToDatadogOrganizationPropertiesOutputWithContext(context.Background())
}

func (i DatadogOrganizationPropertiesArgs) ToDatadogOrganizationPropertiesOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogOrganizationPropertiesOutput)
}

func (i DatadogOrganizationPropertiesArgs) ToDatadogOrganizationPropertiesPtrOutput() DatadogOrganizationPropertiesPtrOutput {
	return i.ToDatadogOrganizationPropertiesPtrOutputWithContext(context.Background())
}

func (i DatadogOrganizationPropertiesArgs) ToDatadogOrganizationPropertiesPtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogOrganizationPropertiesOutput).ToDatadogOrganizationPropertiesPtrOutputWithContext(ctx)
}









type DatadogOrganizationPropertiesPtrInput interface {
	pulumi.Input

	ToDatadogOrganizationPropertiesPtrOutput() DatadogOrganizationPropertiesPtrOutput
	ToDatadogOrganizationPropertiesPtrOutputWithContext(context.Context) DatadogOrganizationPropertiesPtrOutput
}

type datadogOrganizationPropertiesPtrType DatadogOrganizationPropertiesArgs

func DatadogOrganizationPropertiesPtr(v *DatadogOrganizationPropertiesArgs) DatadogOrganizationPropertiesPtrInput {
	return (*datadogOrganizationPropertiesPtrType)(v)
}

func (*datadogOrganizationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogOrganizationProperties)(nil)).Elem()
}

func (i *datadogOrganizationPropertiesPtrType) ToDatadogOrganizationPropertiesPtrOutput() DatadogOrganizationPropertiesPtrOutput {
	return i.ToDatadogOrganizationPropertiesPtrOutputWithContext(context.Background())
}

func (i *datadogOrganizationPropertiesPtrType) ToDatadogOrganizationPropertiesPtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogOrganizationPropertiesPtrOutput)
}

type DatadogOrganizationPropertiesOutput struct{ *pulumi.OutputState }

func (DatadogOrganizationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogOrganizationProperties)(nil)).Elem()
}

func (o DatadogOrganizationPropertiesOutput) ToDatadogOrganizationPropertiesOutput() DatadogOrganizationPropertiesOutput {
	return o
}

func (o DatadogOrganizationPropertiesOutput) ToDatadogOrganizationPropertiesOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesOutput {
	return o
}

func (o DatadogOrganizationPropertiesOutput) ToDatadogOrganizationPropertiesPtrOutput() DatadogOrganizationPropertiesPtrOutput {
	return o.ToDatadogOrganizationPropertiesPtrOutputWithContext(context.Background())
}

func (o DatadogOrganizationPropertiesOutput) ToDatadogOrganizationPropertiesPtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatadogOrganizationProperties) *DatadogOrganizationProperties {
		return &v
	}).(DatadogOrganizationPropertiesPtrOutput)
}

func (o DatadogOrganizationPropertiesOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogOrganizationProperties) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesOutput) ApplicationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogOrganizationProperties) *string { return v.ApplicationKey }).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesOutput) EnterpriseAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogOrganizationProperties) *string { return v.EnterpriseAppId }).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesOutput) LinkingAuthCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogOrganizationProperties) *string { return v.LinkingAuthCode }).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesOutput) LinkingClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogOrganizationProperties) *string { return v.LinkingClientId }).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesOutput) RedirectUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogOrganizationProperties) *string { return v.RedirectUri }).(pulumi.StringPtrOutput)
}

type DatadogOrganizationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DatadogOrganizationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogOrganizationProperties)(nil)).Elem()
}

func (o DatadogOrganizationPropertiesPtrOutput) ToDatadogOrganizationPropertiesPtrOutput() DatadogOrganizationPropertiesPtrOutput {
	return o
}

func (o DatadogOrganizationPropertiesPtrOutput) ToDatadogOrganizationPropertiesPtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesPtrOutput {
	return o
}

func (o DatadogOrganizationPropertiesPtrOutput) Elem() DatadogOrganizationPropertiesOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) DatadogOrganizationProperties {
		if v != nil {
			return *v
		}
		var ret DatadogOrganizationProperties
		return ret
	}).(DatadogOrganizationPropertiesOutput)
}

func (o DatadogOrganizationPropertiesPtrOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApiKey
	}).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesPtrOutput) ApplicationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationKey
	}).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesPtrOutput) EnterpriseAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.EnterpriseAppId
	}).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesPtrOutput) LinkingAuthCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.LinkingAuthCode
	}).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesPtrOutput) LinkingClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.LinkingClientId
	}).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesPtrOutput) RedirectUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.RedirectUri
	}).(pulumi.StringPtrOutput)
}

type DatadogOrganizationPropertiesResponse struct {
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

type DatadogOrganizationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DatadogOrganizationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogOrganizationPropertiesResponse)(nil)).Elem()
}

func (o DatadogOrganizationPropertiesResponseOutput) ToDatadogOrganizationPropertiesResponseOutput() DatadogOrganizationPropertiesResponseOutput {
	return o
}

func (o DatadogOrganizationPropertiesResponseOutput) ToDatadogOrganizationPropertiesResponseOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesResponseOutput {
	return o
}

func (o DatadogOrganizationPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DatadogOrganizationPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o DatadogOrganizationPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DatadogOrganizationPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

type DatadogOrganizationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DatadogOrganizationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogOrganizationPropertiesResponse)(nil)).Elem()
}

func (o DatadogOrganizationPropertiesResponsePtrOutput) ToDatadogOrganizationPropertiesResponsePtrOutput() DatadogOrganizationPropertiesResponsePtrOutput {
	return o
}

func (o DatadogOrganizationPropertiesResponsePtrOutput) ToDatadogOrganizationPropertiesResponsePtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesResponsePtrOutput {
	return o
}

func (o DatadogOrganizationPropertiesResponsePtrOutput) Elem() DatadogOrganizationPropertiesResponseOutput {
	return o.ApplyT(func(v *DatadogOrganizationPropertiesResponse) DatadogOrganizationPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DatadogOrganizationPropertiesResponse
		return ret
	}).(DatadogOrganizationPropertiesResponseOutput)
}

func (o DatadogOrganizationPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type IdentityProperties struct {
	Type *string `pulumi:"type"`
}





type IdentityPropertiesInput interface {
	pulumi.Input

	ToIdentityPropertiesOutput() IdentityPropertiesOutput
	ToIdentityPropertiesOutputWithContext(context.Context) IdentityPropertiesOutput
}

type IdentityPropertiesArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return i.ToIdentityPropertiesOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput)
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput).ToIdentityPropertiesPtrOutputWithContext(ctx)
}









type IdentityPropertiesPtrInput interface {
	pulumi.Input

	ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput
	ToIdentityPropertiesPtrOutputWithContext(context.Context) IdentityPropertiesPtrOutput
}

type identityPropertiesPtrType IdentityPropertiesArgs

func IdentityPropertiesPtr(v *IdentityPropertiesArgs) IdentityPropertiesPtrInput {
	return (*identityPropertiesPtrType)(v)
}

func (*identityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesPtrOutput)
}

type IdentityPropertiesOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProperties) *IdentityProperties {
		return &v
	}).(IdentityPropertiesPtrOutput)
}

func (o IdentityPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) Elem() IdentityPropertiesOutput {
	return o.ApplyT(func(v *IdentityProperties) IdentityProperties {
		if v != nil {
			return *v
		}
		var ret IdentityProperties
		return ret
	}).(IdentityPropertiesOutput)
}

func (o IdentityPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentityPropertiesResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type IdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutputWithContext(ctx context.Context) IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityPropertiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) Elem() IdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) IdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret IdentityPropertiesResponse
		return ret
	}).(IdentityPropertiesResponseOutput)
}

func (o IdentityPropertiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type LinkedResourceResponse struct {
	Id *string `pulumi:"id"`
}

type MonitorProperties struct {
	DatadogOrganizationProperties *DatadogOrganizationProperties `pulumi:"datadogOrganizationProperties"`
	MonitoringStatus              *string                        `pulumi:"monitoringStatus"`
	UserInfo                      *UserInfo                      `pulumi:"userInfo"`
}





type MonitorPropertiesInput interface {
	pulumi.Input

	ToMonitorPropertiesOutput() MonitorPropertiesOutput
	ToMonitorPropertiesOutputWithContext(context.Context) MonitorPropertiesOutput
}

type MonitorPropertiesArgs struct {
	DatadogOrganizationProperties DatadogOrganizationPropertiesPtrInput `pulumi:"datadogOrganizationProperties"`
	MonitoringStatus              pulumi.StringPtrInput                 `pulumi:"monitoringStatus"`
	UserInfo                      UserInfoPtrInput                      `pulumi:"userInfo"`
}

func (MonitorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorProperties)(nil)).Elem()
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesOutput() MonitorPropertiesOutput {
	return i.ToMonitorPropertiesOutputWithContext(context.Background())
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesOutputWithContext(ctx context.Context) MonitorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesOutput)
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return i.ToMonitorPropertiesPtrOutputWithContext(context.Background())
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesOutput).ToMonitorPropertiesPtrOutputWithContext(ctx)
}









type MonitorPropertiesPtrInput interface {
	pulumi.Input

	ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput
	ToMonitorPropertiesPtrOutputWithContext(context.Context) MonitorPropertiesPtrOutput
}

type monitorPropertiesPtrType MonitorPropertiesArgs

func MonitorPropertiesPtr(v *MonitorPropertiesArgs) MonitorPropertiesPtrInput {
	return (*monitorPropertiesPtrType)(v)
}

func (*monitorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorProperties)(nil)).Elem()
}

func (i *monitorPropertiesPtrType) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return i.ToMonitorPropertiesPtrOutputWithContext(context.Background())
}

func (i *monitorPropertiesPtrType) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesPtrOutput)
}

type MonitorPropertiesOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorProperties)(nil)).Elem()
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesOutput() MonitorPropertiesOutput {
	return o
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesOutputWithContext(ctx context.Context) MonitorPropertiesOutput {
	return o
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return o.ToMonitorPropertiesPtrOutputWithContext(context.Background())
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorProperties) *MonitorProperties {
		return &v
	}).(MonitorPropertiesPtrOutput)
}

func (o MonitorPropertiesOutput) DatadogOrganizationProperties() DatadogOrganizationPropertiesPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *DatadogOrganizationProperties { return v.DatadogOrganizationProperties }).(DatadogOrganizationPropertiesPtrOutput)
}

func (o MonitorPropertiesOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *string { return v.MonitoringStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesOutput) UserInfo() UserInfoPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *UserInfo { return v.UserInfo }).(UserInfoPtrOutput)
}

type MonitorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorProperties)(nil)).Elem()
}

func (o MonitorPropertiesPtrOutput) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return o
}

func (o MonitorPropertiesPtrOutput) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return o
}

func (o MonitorPropertiesPtrOutput) Elem() MonitorPropertiesOutput {
	return o.ApplyT(func(v *MonitorProperties) MonitorProperties {
		if v != nil {
			return *v
		}
		var ret MonitorProperties
		return ret
	}).(MonitorPropertiesOutput)
}

func (o MonitorPropertiesPtrOutput) DatadogOrganizationProperties() DatadogOrganizationPropertiesPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *DatadogOrganizationProperties {
		if v == nil {
			return nil
		}
		return v.DatadogOrganizationProperties
	}).(DatadogOrganizationPropertiesPtrOutput)
}

func (o MonitorPropertiesPtrOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *string {
		if v == nil {
			return nil
		}
		return v.MonitoringStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesPtrOutput) UserInfo() UserInfoPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *UserInfo {
		if v == nil {
			return nil
		}
		return v.UserInfo
	}).(UserInfoPtrOutput)
}

type MonitorPropertiesResponse struct {
	DatadogOrganizationProperties *DatadogOrganizationPropertiesResponse `pulumi:"datadogOrganizationProperties"`
	LiftrResourceCategory         string                                 `pulumi:"liftrResourceCategory"`
	LiftrResourcePreference       int                                    `pulumi:"liftrResourcePreference"`
	MarketplaceSubscriptionStatus string                                 `pulumi:"marketplaceSubscriptionStatus"`
	MonitoringStatus              *string                                `pulumi:"monitoringStatus"`
	ProvisioningState             string                                 `pulumi:"provisioningState"`
	UserInfo                      *UserInfoResponse                      `pulumi:"userInfo"`
}

type MonitorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorPropertiesResponse)(nil)).Elem()
}

func (o MonitorPropertiesResponseOutput) ToMonitorPropertiesResponseOutput() MonitorPropertiesResponseOutput {
	return o
}

func (o MonitorPropertiesResponseOutput) ToMonitorPropertiesResponseOutputWithContext(ctx context.Context) MonitorPropertiesResponseOutput {
	return o
}

func (o MonitorPropertiesResponseOutput) DatadogOrganizationProperties() DatadogOrganizationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *DatadogOrganizationPropertiesResponse {
		return v.DatadogOrganizationProperties
	}).(DatadogOrganizationPropertiesResponsePtrOutput)
}

func (o MonitorPropertiesResponseOutput) LiftrResourceCategory() pulumi.StringOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) string { return v.LiftrResourceCategory }).(pulumi.StringOutput)
}

func (o MonitorPropertiesResponseOutput) LiftrResourcePreference() pulumi.IntOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) int { return v.LiftrResourcePreference }).(pulumi.IntOutput)
}

func (o MonitorPropertiesResponseOutput) MarketplaceSubscriptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) string { return v.MarketplaceSubscriptionStatus }).(pulumi.StringOutput)
}

func (o MonitorPropertiesResponseOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *string { return v.MonitoringStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MonitorPropertiesResponseOutput) UserInfo() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *UserInfoResponse { return v.UserInfo }).(UserInfoResponsePtrOutput)
}

type MonitoredResourceResponse struct {
	Id                     *string `pulumi:"id"`
	ReasonForLogsStatus    *string `pulumi:"reasonForLogsStatus"`
	ReasonForMetricsStatus *string `pulumi:"reasonForMetricsStatus"`
	SendingLogs            *bool   `pulumi:"sendingLogs"`
	SendingMetrics         *bool   `pulumi:"sendingMetrics"`
}

type ResourceSku struct {
	Name string `pulumi:"name"`
}





type ResourceSkuInput interface {
	pulumi.Input

	ToResourceSkuOutput() ResourceSkuOutput
	ToResourceSkuOutputWithContext(context.Context) ResourceSkuOutput
}

type ResourceSkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (ResourceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (i ResourceSkuArgs) ToResourceSkuOutput() ResourceSkuOutput {
	return i.ToResourceSkuOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput)
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput).ToResourceSkuPtrOutputWithContext(ctx)
}









type ResourceSkuPtrInput interface {
	pulumi.Input

	ToResourceSkuPtrOutput() ResourceSkuPtrOutput
	ToResourceSkuPtrOutputWithContext(context.Context) ResourceSkuPtrOutput
}

type resourceSkuPtrType ResourceSkuArgs

func ResourceSkuPtr(v *ResourceSkuArgs) ResourceSkuPtrInput {
	return (*resourceSkuPtrType)(v)
}

func (*resourceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuPtrOutput)
}

type ResourceSkuOutput struct{ *pulumi.OutputState }

func (ResourceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (o ResourceSkuOutput) ToResourceSkuOutput() ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSku) *ResourceSku {
		return &v
	}).(ResourceSkuPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSku) string { return v.Name }).(pulumi.StringOutput)
}

type ResourceSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) Elem() ResourceSkuOutput {
	return o.ApplyT(func(v *ResourceSku) ResourceSku {
		if v != nil {
			return *v
		}
		var ret ResourceSku
		return ret
	}).(ResourceSkuOutput)
}

func (o ResourceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Name string `pulumi:"name"`
}

type ResourceSkuResponseOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ResourceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) Elem() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) ResourceSkuResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSkuResponse
		return ret
	}).(ResourceSkuResponseOutput)
}

func (o ResourceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
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

type UserInfo struct {
	EmailAddress *string `pulumi:"emailAddress"`
	Name         *string `pulumi:"name"`
	PhoneNumber  *string `pulumi:"phoneNumber"`
}





type UserInfoInput interface {
	pulumi.Input

	ToUserInfoOutput() UserInfoOutput
	ToUserInfoOutputWithContext(context.Context) UserInfoOutput
}

type UserInfoArgs struct {
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	PhoneNumber  pulumi.StringPtrInput `pulumi:"phoneNumber"`
}

func (UserInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfo)(nil)).Elem()
}

func (i UserInfoArgs) ToUserInfoOutput() UserInfoOutput {
	return i.ToUserInfoOutputWithContext(context.Background())
}

func (i UserInfoArgs) ToUserInfoOutputWithContext(ctx context.Context) UserInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoOutput)
}

func (i UserInfoArgs) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return i.ToUserInfoPtrOutputWithContext(context.Background())
}

func (i UserInfoArgs) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoOutput).ToUserInfoPtrOutputWithContext(ctx)
}









type UserInfoPtrInput interface {
	pulumi.Input

	ToUserInfoPtrOutput() UserInfoPtrOutput
	ToUserInfoPtrOutputWithContext(context.Context) UserInfoPtrOutput
}

type userInfoPtrType UserInfoArgs

func UserInfoPtr(v *UserInfoArgs) UserInfoPtrInput {
	return (*userInfoPtrType)(v)
}

func (*userInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfo)(nil)).Elem()
}

func (i *userInfoPtrType) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return i.ToUserInfoPtrOutputWithContext(context.Background())
}

func (i *userInfoPtrType) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoPtrOutput)
}

type UserInfoOutput struct{ *pulumi.OutputState }

func (UserInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfo)(nil)).Elem()
}

func (o UserInfoOutput) ToUserInfoOutput() UserInfoOutput {
	return o
}

func (o UserInfoOutput) ToUserInfoOutputWithContext(ctx context.Context) UserInfoOutput {
	return o
}

func (o UserInfoOutput) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return o.ToUserInfoPtrOutputWithContext(context.Background())
}

func (o UserInfoOutput) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserInfo) *UserInfo {
		return &v
	}).(UserInfoPtrOutput)
}

func (o UserInfoOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

func (o UserInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o UserInfoOutput) PhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.PhoneNumber }).(pulumi.StringPtrOutput)
}

type UserInfoPtrOutput struct{ *pulumi.OutputState }

func (UserInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfo)(nil)).Elem()
}

func (o UserInfoPtrOutput) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return o
}

func (o UserInfoPtrOutput) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return o
}

func (o UserInfoPtrOutput) Elem() UserInfoOutput {
	return o.ApplyT(func(v *UserInfo) UserInfo {
		if v != nil {
			return *v
		}
		var ret UserInfo
		return ret
	}).(UserInfoOutput)
}

func (o UserInfoPtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoPtrOutput) PhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.PhoneNumber
	}).(pulumi.StringPtrOutput)
}

type UserInfoResponse struct {
	EmailAddress *string `pulumi:"emailAddress"`
	Name         *string `pulumi:"name"`
	PhoneNumber  *string `pulumi:"phoneNumber"`
}

type UserInfoResponseOutput struct{ *pulumi.OutputState }

func (UserInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfoResponse)(nil)).Elem()
}

func (o UserInfoResponseOutput) ToUserInfoResponseOutput() UserInfoResponseOutput {
	return o
}

func (o UserInfoResponseOutput) ToUserInfoResponseOutputWithContext(ctx context.Context) UserInfoResponseOutput {
	return o
}

func (o UserInfoResponseOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) PhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.PhoneNumber }).(pulumi.StringPtrOutput)
}

type UserInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (UserInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfoResponse)(nil)).Elem()
}

func (o UserInfoResponsePtrOutput) ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput {
	return o
}

func (o UserInfoResponsePtrOutput) ToUserInfoResponsePtrOutputWithContext(ctx context.Context) UserInfoResponsePtrOutput {
	return o
}

func (o UserInfoResponsePtrOutput) Elem() UserInfoResponseOutput {
	return o.ApplyT(func(v *UserInfoResponse) UserInfoResponse {
		if v != nil {
			return *v
		}
		var ret UserInfoResponse
		return ret
	}).(UserInfoResponseOutput)
}

func (o UserInfoResponsePtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) PhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PhoneNumber
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DatadogOrganizationPropertiesOutput{})
	pulumi.RegisterOutputType(DatadogOrganizationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DatadogOrganizationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DatadogOrganizationPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserInfoOutput{})
	pulumi.RegisterOutputType(UserInfoPtrOutput{})
	pulumi.RegisterOutputType(UserInfoResponseOutput{})
	pulumi.RegisterOutputType(UserInfoResponsePtrOutput{})
}

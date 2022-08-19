


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Backend struct {
	Address                    *string `pulumi:"address"`
	BackendHostHeader          *string `pulumi:"backendHostHeader"`
	EnabledState               *string `pulumi:"enabledState"`
	HttpPort                   *int    `pulumi:"httpPort"`
	HttpsPort                  *int    `pulumi:"httpsPort"`
	Priority                   *int    `pulumi:"priority"`
	PrivateLinkAlias           *string `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage *string `pulumi:"privateLinkApprovalMessage"`
	PrivateLinkLocation        *string `pulumi:"privateLinkLocation"`
	PrivateLinkResourceId      *string `pulumi:"privateLinkResourceId"`
	Weight                     *int    `pulumi:"weight"`
}





type BackendInput interface {
	pulumi.Input

	ToBackendOutput() BackendOutput
	ToBackendOutputWithContext(context.Context) BackendOutput
}

type BackendArgs struct {
	Address                    pulumi.StringPtrInput `pulumi:"address"`
	BackendHostHeader          pulumi.StringPtrInput `pulumi:"backendHostHeader"`
	EnabledState               pulumi.StringPtrInput `pulumi:"enabledState"`
	HttpPort                   pulumi.IntPtrInput    `pulumi:"httpPort"`
	HttpsPort                  pulumi.IntPtrInput    `pulumi:"httpsPort"`
	Priority                   pulumi.IntPtrInput    `pulumi:"priority"`
	PrivateLinkAlias           pulumi.StringPtrInput `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage pulumi.StringPtrInput `pulumi:"privateLinkApprovalMessage"`
	PrivateLinkLocation        pulumi.StringPtrInput `pulumi:"privateLinkLocation"`
	PrivateLinkResourceId      pulumi.StringPtrInput `pulumi:"privateLinkResourceId"`
	Weight                     pulumi.IntPtrInput    `pulumi:"weight"`
}

func (BackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Backend)(nil)).Elem()
}

func (i BackendArgs) ToBackendOutput() BackendOutput {
	return i.ToBackendOutputWithContext(context.Background())
}

func (i BackendArgs) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendOutput)
}





type BackendArrayInput interface {
	pulumi.Input

	ToBackendArrayOutput() BackendArrayOutput
	ToBackendArrayOutputWithContext(context.Context) BackendArrayOutput
}

type BackendArray []BackendInput

func (BackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Backend)(nil)).Elem()
}

func (i BackendArray) ToBackendArrayOutput() BackendArrayOutput {
	return i.ToBackendArrayOutputWithContext(context.Background())
}

func (i BackendArray) ToBackendArrayOutputWithContext(ctx context.Context) BackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendArrayOutput)
}

type BackendOutput struct{ *pulumi.OutputState }

func (BackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Backend)(nil)).Elem()
}

func (o BackendOutput) ToBackendOutput() BackendOutput {
	return o
}

func (o BackendOutput) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return o
}

func (o BackendOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backend) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) BackendHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backend) *string { return v.BackendHostHeader }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backend) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Backend) *int { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o BackendOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Backend) *int { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o BackendOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Backend) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o BackendOutput) PrivateLinkAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backend) *string { return v.PrivateLinkAlias }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) PrivateLinkApprovalMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backend) *string { return v.PrivateLinkApprovalMessage }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) PrivateLinkLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backend) *string { return v.PrivateLinkLocation }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backend) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Backend) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type BackendArrayOutput struct{ *pulumi.OutputState }

func (BackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Backend)(nil)).Elem()
}

func (o BackendArrayOutput) ToBackendArrayOutput() BackendArrayOutput {
	return o
}

func (o BackendArrayOutput) ToBackendArrayOutputWithContext(ctx context.Context) BackendArrayOutput {
	return o
}

func (o BackendArrayOutput) Index(i pulumi.IntInput) BackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Backend {
		return vs[0].([]Backend)[vs[1].(int)]
	}).(BackendOutput)
}

type BackendPool struct {
	Backends              []Backend    `pulumi:"backends"`
	HealthProbeSettings   *SubResource `pulumi:"healthProbeSettings"`
	Id                    *string      `pulumi:"id"`
	LoadBalancingSettings *SubResource `pulumi:"loadBalancingSettings"`
	Name                  *string      `pulumi:"name"`
}





type BackendPoolInput interface {
	pulumi.Input

	ToBackendPoolOutput() BackendPoolOutput
	ToBackendPoolOutputWithContext(context.Context) BackendPoolOutput
}

type BackendPoolArgs struct {
	Backends              BackendArrayInput     `pulumi:"backends"`
	HealthProbeSettings   SubResourcePtrInput   `pulumi:"healthProbeSettings"`
	Id                    pulumi.StringPtrInput `pulumi:"id"`
	LoadBalancingSettings SubResourcePtrInput   `pulumi:"loadBalancingSettings"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
}

func (BackendPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPool)(nil)).Elem()
}

func (i BackendPoolArgs) ToBackendPoolOutput() BackendPoolOutput {
	return i.ToBackendPoolOutputWithContext(context.Background())
}

func (i BackendPoolArgs) ToBackendPoolOutputWithContext(ctx context.Context) BackendPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPoolOutput)
}





type BackendPoolArrayInput interface {
	pulumi.Input

	ToBackendPoolArrayOutput() BackendPoolArrayOutput
	ToBackendPoolArrayOutputWithContext(context.Context) BackendPoolArrayOutput
}

type BackendPoolArray []BackendPoolInput

func (BackendPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendPool)(nil)).Elem()
}

func (i BackendPoolArray) ToBackendPoolArrayOutput() BackendPoolArrayOutput {
	return i.ToBackendPoolArrayOutputWithContext(context.Background())
}

func (i BackendPoolArray) ToBackendPoolArrayOutputWithContext(ctx context.Context) BackendPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPoolArrayOutput)
}

type BackendPoolOutput struct{ *pulumi.OutputState }

func (BackendPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPool)(nil)).Elem()
}

func (o BackendPoolOutput) ToBackendPoolOutput() BackendPoolOutput {
	return o
}

func (o BackendPoolOutput) ToBackendPoolOutputWithContext(ctx context.Context) BackendPoolOutput {
	return o
}

func (o BackendPoolOutput) Backends() BackendArrayOutput {
	return o.ApplyT(func(v BackendPool) []Backend { return v.Backends }).(BackendArrayOutput)
}

func (o BackendPoolOutput) HealthProbeSettings() SubResourcePtrOutput {
	return o.ApplyT(func(v BackendPool) *SubResource { return v.HealthProbeSettings }).(SubResourcePtrOutput)
}

func (o BackendPoolOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendPool) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BackendPoolOutput) LoadBalancingSettings() SubResourcePtrOutput {
	return o.ApplyT(func(v BackendPool) *SubResource { return v.LoadBalancingSettings }).(SubResourcePtrOutput)
}

func (o BackendPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendPool) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type BackendPoolArrayOutput struct{ *pulumi.OutputState }

func (BackendPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendPool)(nil)).Elem()
}

func (o BackendPoolArrayOutput) ToBackendPoolArrayOutput() BackendPoolArrayOutput {
	return o
}

func (o BackendPoolArrayOutput) ToBackendPoolArrayOutputWithContext(ctx context.Context) BackendPoolArrayOutput {
	return o
}

func (o BackendPoolArrayOutput) Index(i pulumi.IntInput) BackendPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendPool {
		return vs[0].([]BackendPool)[vs[1].(int)]
	}).(BackendPoolOutput)
}

type BackendPoolResponse struct {
	Backends              []BackendResponse    `pulumi:"backends"`
	HealthProbeSettings   *SubResourceResponse `pulumi:"healthProbeSettings"`
	Id                    *string              `pulumi:"id"`
	LoadBalancingSettings *SubResourceResponse `pulumi:"loadBalancingSettings"`
	Name                  *string              `pulumi:"name"`
	ResourceState         string               `pulumi:"resourceState"`
	Type                  string               `pulumi:"type"`
}

type BackendPoolResponseOutput struct{ *pulumi.OutputState }

func (BackendPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPoolResponse)(nil)).Elem()
}

func (o BackendPoolResponseOutput) ToBackendPoolResponseOutput() BackendPoolResponseOutput {
	return o
}

func (o BackendPoolResponseOutput) ToBackendPoolResponseOutputWithContext(ctx context.Context) BackendPoolResponseOutput {
	return o
}

func (o BackendPoolResponseOutput) Backends() BackendResponseArrayOutput {
	return o.ApplyT(func(v BackendPoolResponse) []BackendResponse { return v.Backends }).(BackendResponseArrayOutput)
}

func (o BackendPoolResponseOutput) HealthProbeSettings() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v BackendPoolResponse) *SubResourceResponse { return v.HealthProbeSettings }).(SubResourceResponsePtrOutput)
}

func (o BackendPoolResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendPoolResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BackendPoolResponseOutput) LoadBalancingSettings() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v BackendPoolResponse) *SubResourceResponse { return v.LoadBalancingSettings }).(SubResourceResponsePtrOutput)
}

func (o BackendPoolResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendPoolResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o BackendPoolResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v BackendPoolResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o BackendPoolResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BackendPoolResponse) string { return v.Type }).(pulumi.StringOutput)
}

type BackendPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (BackendPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendPoolResponse)(nil)).Elem()
}

func (o BackendPoolResponseArrayOutput) ToBackendPoolResponseArrayOutput() BackendPoolResponseArrayOutput {
	return o
}

func (o BackendPoolResponseArrayOutput) ToBackendPoolResponseArrayOutputWithContext(ctx context.Context) BackendPoolResponseArrayOutput {
	return o
}

func (o BackendPoolResponseArrayOutput) Index(i pulumi.IntInput) BackendPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendPoolResponse {
		return vs[0].([]BackendPoolResponse)[vs[1].(int)]
	}).(BackendPoolResponseOutput)
}

type BackendPoolsSettings struct {
	EnforceCertificateNameCheck *string `pulumi:"enforceCertificateNameCheck"`
	SendRecvTimeoutSeconds      *int    `pulumi:"sendRecvTimeoutSeconds"`
}


func (val *BackendPoolsSettings) Defaults() *BackendPoolsSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnforceCertificateNameCheck) {
		enforceCertificateNameCheck_ := "Enabled"
		tmp.EnforceCertificateNameCheck = &enforceCertificateNameCheck_
	}
	return &tmp
}





type BackendPoolsSettingsInput interface {
	pulumi.Input

	ToBackendPoolsSettingsOutput() BackendPoolsSettingsOutput
	ToBackendPoolsSettingsOutputWithContext(context.Context) BackendPoolsSettingsOutput
}

type BackendPoolsSettingsArgs struct {
	EnforceCertificateNameCheck pulumi.StringPtrInput `pulumi:"enforceCertificateNameCheck"`
	SendRecvTimeoutSeconds      pulumi.IntPtrInput    `pulumi:"sendRecvTimeoutSeconds"`
}


func (val *BackendPoolsSettingsArgs) Defaults() *BackendPoolsSettingsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnforceCertificateNameCheck) {
		tmp.EnforceCertificateNameCheck = pulumi.StringPtr("Enabled")
	}
	return &tmp
}
func (BackendPoolsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPoolsSettings)(nil)).Elem()
}

func (i BackendPoolsSettingsArgs) ToBackendPoolsSettingsOutput() BackendPoolsSettingsOutput {
	return i.ToBackendPoolsSettingsOutputWithContext(context.Background())
}

func (i BackendPoolsSettingsArgs) ToBackendPoolsSettingsOutputWithContext(ctx context.Context) BackendPoolsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPoolsSettingsOutput)
}

func (i BackendPoolsSettingsArgs) ToBackendPoolsSettingsPtrOutput() BackendPoolsSettingsPtrOutput {
	return i.ToBackendPoolsSettingsPtrOutputWithContext(context.Background())
}

func (i BackendPoolsSettingsArgs) ToBackendPoolsSettingsPtrOutputWithContext(ctx context.Context) BackendPoolsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPoolsSettingsOutput).ToBackendPoolsSettingsPtrOutputWithContext(ctx)
}









type BackendPoolsSettingsPtrInput interface {
	pulumi.Input

	ToBackendPoolsSettingsPtrOutput() BackendPoolsSettingsPtrOutput
	ToBackendPoolsSettingsPtrOutputWithContext(context.Context) BackendPoolsSettingsPtrOutput
}

type backendPoolsSettingsPtrType BackendPoolsSettingsArgs

func BackendPoolsSettingsPtr(v *BackendPoolsSettingsArgs) BackendPoolsSettingsPtrInput {
	return (*backendPoolsSettingsPtrType)(v)
}

func (*backendPoolsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendPoolsSettings)(nil)).Elem()
}

func (i *backendPoolsSettingsPtrType) ToBackendPoolsSettingsPtrOutput() BackendPoolsSettingsPtrOutput {
	return i.ToBackendPoolsSettingsPtrOutputWithContext(context.Background())
}

func (i *backendPoolsSettingsPtrType) ToBackendPoolsSettingsPtrOutputWithContext(ctx context.Context) BackendPoolsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPoolsSettingsPtrOutput)
}

type BackendPoolsSettingsOutput struct{ *pulumi.OutputState }

func (BackendPoolsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPoolsSettings)(nil)).Elem()
}

func (o BackendPoolsSettingsOutput) ToBackendPoolsSettingsOutput() BackendPoolsSettingsOutput {
	return o
}

func (o BackendPoolsSettingsOutput) ToBackendPoolsSettingsOutputWithContext(ctx context.Context) BackendPoolsSettingsOutput {
	return o
}

func (o BackendPoolsSettingsOutput) ToBackendPoolsSettingsPtrOutput() BackendPoolsSettingsPtrOutput {
	return o.ToBackendPoolsSettingsPtrOutputWithContext(context.Background())
}

func (o BackendPoolsSettingsOutput) ToBackendPoolsSettingsPtrOutputWithContext(ctx context.Context) BackendPoolsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendPoolsSettings) *BackendPoolsSettings {
		return &v
	}).(BackendPoolsSettingsPtrOutput)
}

func (o BackendPoolsSettingsOutput) EnforceCertificateNameCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendPoolsSettings) *string { return v.EnforceCertificateNameCheck }).(pulumi.StringPtrOutput)
}

func (o BackendPoolsSettingsOutput) SendRecvTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendPoolsSettings) *int { return v.SendRecvTimeoutSeconds }).(pulumi.IntPtrOutput)
}

type BackendPoolsSettingsPtrOutput struct{ *pulumi.OutputState }

func (BackendPoolsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendPoolsSettings)(nil)).Elem()
}

func (o BackendPoolsSettingsPtrOutput) ToBackendPoolsSettingsPtrOutput() BackendPoolsSettingsPtrOutput {
	return o
}

func (o BackendPoolsSettingsPtrOutput) ToBackendPoolsSettingsPtrOutputWithContext(ctx context.Context) BackendPoolsSettingsPtrOutput {
	return o
}

func (o BackendPoolsSettingsPtrOutput) Elem() BackendPoolsSettingsOutput {
	return o.ApplyT(func(v *BackendPoolsSettings) BackendPoolsSettings {
		if v != nil {
			return *v
		}
		var ret BackendPoolsSettings
		return ret
	}).(BackendPoolsSettingsOutput)
}

func (o BackendPoolsSettingsPtrOutput) EnforceCertificateNameCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendPoolsSettings) *string {
		if v == nil {
			return nil
		}
		return v.EnforceCertificateNameCheck
	}).(pulumi.StringPtrOutput)
}

func (o BackendPoolsSettingsPtrOutput) SendRecvTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackendPoolsSettings) *int {
		if v == nil {
			return nil
		}
		return v.SendRecvTimeoutSeconds
	}).(pulumi.IntPtrOutput)
}

type BackendPoolsSettingsResponse struct {
	EnforceCertificateNameCheck *string `pulumi:"enforceCertificateNameCheck"`
	SendRecvTimeoutSeconds      *int    `pulumi:"sendRecvTimeoutSeconds"`
}


func (val *BackendPoolsSettingsResponse) Defaults() *BackendPoolsSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnforceCertificateNameCheck) {
		enforceCertificateNameCheck_ := "Enabled"
		tmp.EnforceCertificateNameCheck = &enforceCertificateNameCheck_
	}
	return &tmp
}

type BackendPoolsSettingsResponseOutput struct{ *pulumi.OutputState }

func (BackendPoolsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPoolsSettingsResponse)(nil)).Elem()
}

func (o BackendPoolsSettingsResponseOutput) ToBackendPoolsSettingsResponseOutput() BackendPoolsSettingsResponseOutput {
	return o
}

func (o BackendPoolsSettingsResponseOutput) ToBackendPoolsSettingsResponseOutputWithContext(ctx context.Context) BackendPoolsSettingsResponseOutput {
	return o
}

func (o BackendPoolsSettingsResponseOutput) EnforceCertificateNameCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendPoolsSettingsResponse) *string { return v.EnforceCertificateNameCheck }).(pulumi.StringPtrOutput)
}

func (o BackendPoolsSettingsResponseOutput) SendRecvTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendPoolsSettingsResponse) *int { return v.SendRecvTimeoutSeconds }).(pulumi.IntPtrOutput)
}

type BackendPoolsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (BackendPoolsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendPoolsSettingsResponse)(nil)).Elem()
}

func (o BackendPoolsSettingsResponsePtrOutput) ToBackendPoolsSettingsResponsePtrOutput() BackendPoolsSettingsResponsePtrOutput {
	return o
}

func (o BackendPoolsSettingsResponsePtrOutput) ToBackendPoolsSettingsResponsePtrOutputWithContext(ctx context.Context) BackendPoolsSettingsResponsePtrOutput {
	return o
}

func (o BackendPoolsSettingsResponsePtrOutput) Elem() BackendPoolsSettingsResponseOutput {
	return o.ApplyT(func(v *BackendPoolsSettingsResponse) BackendPoolsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret BackendPoolsSettingsResponse
		return ret
	}).(BackendPoolsSettingsResponseOutput)
}

func (o BackendPoolsSettingsResponsePtrOutput) EnforceCertificateNameCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendPoolsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.EnforceCertificateNameCheck
	}).(pulumi.StringPtrOutput)
}

func (o BackendPoolsSettingsResponsePtrOutput) SendRecvTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackendPoolsSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.SendRecvTimeoutSeconds
	}).(pulumi.IntPtrOutput)
}

type BackendResponse struct {
	Address                    *string `pulumi:"address"`
	BackendHostHeader          *string `pulumi:"backendHostHeader"`
	EnabledState               *string `pulumi:"enabledState"`
	HttpPort                   *int    `pulumi:"httpPort"`
	HttpsPort                  *int    `pulumi:"httpsPort"`
	Priority                   *int    `pulumi:"priority"`
	PrivateEndpointStatus      string  `pulumi:"privateEndpointStatus"`
	PrivateLinkAlias           *string `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage *string `pulumi:"privateLinkApprovalMessage"`
	PrivateLinkLocation        *string `pulumi:"privateLinkLocation"`
	PrivateLinkResourceId      *string `pulumi:"privateLinkResourceId"`
	Weight                     *int    `pulumi:"weight"`
}

type BackendResponseOutput struct{ *pulumi.OutputState }

func (BackendResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendResponse)(nil)).Elem()
}

func (o BackendResponseOutput) ToBackendResponseOutput() BackendResponseOutput {
	return o
}

func (o BackendResponseOutput) ToBackendResponseOutputWithContext(ctx context.Context) BackendResponseOutput {
	return o
}

func (o BackendResponseOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendResponse) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o BackendResponseOutput) BackendHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendResponse) *string { return v.BackendHostHeader }).(pulumi.StringPtrOutput)
}

func (o BackendResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o BackendResponseOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendResponse) *int { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o BackendResponseOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendResponse) *int { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o BackendResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o BackendResponseOutput) PrivateEndpointStatus() pulumi.StringOutput {
	return o.ApplyT(func(v BackendResponse) string { return v.PrivateEndpointStatus }).(pulumi.StringOutput)
}

func (o BackendResponseOutput) PrivateLinkAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendResponse) *string { return v.PrivateLinkAlias }).(pulumi.StringPtrOutput)
}

func (o BackendResponseOutput) PrivateLinkApprovalMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendResponse) *string { return v.PrivateLinkApprovalMessage }).(pulumi.StringPtrOutput)
}

func (o BackendResponseOutput) PrivateLinkLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendResponse) *string { return v.PrivateLinkLocation }).(pulumi.StringPtrOutput)
}

func (o BackendResponseOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendResponse) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

func (o BackendResponseOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendResponse) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type BackendResponseArrayOutput struct{ *pulumi.OutputState }

func (BackendResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendResponse)(nil)).Elem()
}

func (o BackendResponseArrayOutput) ToBackendResponseArrayOutput() BackendResponseArrayOutput {
	return o
}

func (o BackendResponseArrayOutput) ToBackendResponseArrayOutputWithContext(ctx context.Context) BackendResponseArrayOutput {
	return o
}

func (o BackendResponseArrayOutput) Index(i pulumi.IntInput) BackendResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendResponse {
		return vs[0].([]BackendResponse)[vs[1].(int)]
	}).(BackendResponseOutput)
}

type CacheConfiguration struct {
	CacheDuration                *string `pulumi:"cacheDuration"`
	DynamicCompression           *string `pulumi:"dynamicCompression"`
	QueryParameterStripDirective *string `pulumi:"queryParameterStripDirective"`
	QueryParameters              *string `pulumi:"queryParameters"`
}

type CacheConfigurationResponse struct {
	CacheDuration                *string `pulumi:"cacheDuration"`
	DynamicCompression           *string `pulumi:"dynamicCompression"`
	QueryParameterStripDirective *string `pulumi:"queryParameterStripDirective"`
	QueryParameters              *string `pulumi:"queryParameters"`
}

type CustomHttpsConfigurationResponse struct {
	CertificateSource string                                            `pulumi:"certificateSource"`
	CertificateType   *string                                           `pulumi:"certificateType"`
	MinimumTlsVersion string                                            `pulumi:"minimumTlsVersion"`
	ProtocolType      string                                            `pulumi:"protocolType"`
	SecretName        *string                                           `pulumi:"secretName"`
	SecretVersion     *string                                           `pulumi:"secretVersion"`
	Vault             *KeyVaultCertificateSourceParametersResponseVault `pulumi:"vault"`
}

type CustomHttpsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (CustomHttpsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHttpsConfigurationResponse)(nil)).Elem()
}

func (o CustomHttpsConfigurationResponseOutput) ToCustomHttpsConfigurationResponseOutput() CustomHttpsConfigurationResponseOutput {
	return o
}

func (o CustomHttpsConfigurationResponseOutput) ToCustomHttpsConfigurationResponseOutputWithContext(ctx context.Context) CustomHttpsConfigurationResponseOutput {
	return o
}

func (o CustomHttpsConfigurationResponseOutput) CertificateSource() pulumi.StringOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) string { return v.CertificateSource }).(pulumi.StringOutput)
}

func (o CustomHttpsConfigurationResponseOutput) CertificateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) *string { return v.CertificateType }).(pulumi.StringPtrOutput)
}

func (o CustomHttpsConfigurationResponseOutput) MinimumTlsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) string { return v.MinimumTlsVersion }).(pulumi.StringOutput)
}

func (o CustomHttpsConfigurationResponseOutput) ProtocolType() pulumi.StringOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) string { return v.ProtocolType }).(pulumi.StringOutput)
}

func (o CustomHttpsConfigurationResponseOutput) SecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) *string { return v.SecretName }).(pulumi.StringPtrOutput)
}

func (o CustomHttpsConfigurationResponseOutput) SecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) *string { return v.SecretVersion }).(pulumi.StringPtrOutput)
}

func (o CustomHttpsConfigurationResponseOutput) Vault() KeyVaultCertificateSourceParametersResponseVaultPtrOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) *KeyVaultCertificateSourceParametersResponseVault {
		return v.Vault
	}).(KeyVaultCertificateSourceParametersResponseVaultPtrOutput)
}

type CustomRule struct {
	Action                     string                    `pulumi:"action"`
	EnabledState               *string                   `pulumi:"enabledState"`
	MatchConditions            []FrontDoorMatchCondition `pulumi:"matchConditions"`
	Name                       *string                   `pulumi:"name"`
	Priority                   int                       `pulumi:"priority"`
	RateLimitDurationInMinutes *int                      `pulumi:"rateLimitDurationInMinutes"`
	RateLimitThreshold         *int                      `pulumi:"rateLimitThreshold"`
	RuleType                   string                    `pulumi:"ruleType"`
}





type CustomRuleInput interface {
	pulumi.Input

	ToCustomRuleOutput() CustomRuleOutput
	ToCustomRuleOutputWithContext(context.Context) CustomRuleOutput
}

type CustomRuleArgs struct {
	Action                     pulumi.StringInput                `pulumi:"action"`
	EnabledState               pulumi.StringPtrInput             `pulumi:"enabledState"`
	MatchConditions            FrontDoorMatchConditionArrayInput `pulumi:"matchConditions"`
	Name                       pulumi.StringPtrInput             `pulumi:"name"`
	Priority                   pulumi.IntInput                   `pulumi:"priority"`
	RateLimitDurationInMinutes pulumi.IntPtrInput                `pulumi:"rateLimitDurationInMinutes"`
	RateLimitThreshold         pulumi.IntPtrInput                `pulumi:"rateLimitThreshold"`
	RuleType                   pulumi.StringInput                `pulumi:"ruleType"`
}

func (CustomRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRule)(nil)).Elem()
}

func (i CustomRuleArgs) ToCustomRuleOutput() CustomRuleOutput {
	return i.ToCustomRuleOutputWithContext(context.Background())
}

func (i CustomRuleArgs) ToCustomRuleOutputWithContext(ctx context.Context) CustomRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleOutput)
}





type CustomRuleArrayInput interface {
	pulumi.Input

	ToCustomRuleArrayOutput() CustomRuleArrayOutput
	ToCustomRuleArrayOutputWithContext(context.Context) CustomRuleArrayOutput
}

type CustomRuleArray []CustomRuleInput

func (CustomRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRule)(nil)).Elem()
}

func (i CustomRuleArray) ToCustomRuleArrayOutput() CustomRuleArrayOutput {
	return i.ToCustomRuleArrayOutputWithContext(context.Background())
}

func (i CustomRuleArray) ToCustomRuleArrayOutputWithContext(ctx context.Context) CustomRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleArrayOutput)
}

type CustomRuleOutput struct{ *pulumi.OutputState }

func (CustomRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRule)(nil)).Elem()
}

func (o CustomRuleOutput) ToCustomRuleOutput() CustomRuleOutput {
	return o
}

func (o CustomRuleOutput) ToCustomRuleOutputWithContext(ctx context.Context) CustomRuleOutput {
	return o
}

func (o CustomRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRule) string { return v.Action }).(pulumi.StringOutput)
}

func (o CustomRuleOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRule) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o CustomRuleOutput) MatchConditions() FrontDoorMatchConditionArrayOutput {
	return o.ApplyT(func(v CustomRule) []FrontDoorMatchCondition { return v.MatchConditions }).(FrontDoorMatchConditionArrayOutput)
}

func (o CustomRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CustomRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v CustomRule) int { return v.Priority }).(pulumi.IntOutput)
}

func (o CustomRuleOutput) RateLimitDurationInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CustomRule) *int { return v.RateLimitDurationInMinutes }).(pulumi.IntPtrOutput)
}

func (o CustomRuleOutput) RateLimitThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CustomRule) *int { return v.RateLimitThreshold }).(pulumi.IntPtrOutput)
}

func (o CustomRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRule) string { return v.RuleType }).(pulumi.StringOutput)
}

type CustomRuleArrayOutput struct{ *pulumi.OutputState }

func (CustomRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRule)(nil)).Elem()
}

func (o CustomRuleArrayOutput) ToCustomRuleArrayOutput() CustomRuleArrayOutput {
	return o
}

func (o CustomRuleArrayOutput) ToCustomRuleArrayOutputWithContext(ctx context.Context) CustomRuleArrayOutput {
	return o
}

func (o CustomRuleArrayOutput) Index(i pulumi.IntInput) CustomRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomRule {
		return vs[0].([]CustomRule)[vs[1].(int)]
	}).(CustomRuleOutput)
}

type CustomRuleList struct {
	Rules []CustomRule `pulumi:"rules"`
}





type CustomRuleListInput interface {
	pulumi.Input

	ToCustomRuleListOutput() CustomRuleListOutput
	ToCustomRuleListOutputWithContext(context.Context) CustomRuleListOutput
}

type CustomRuleListArgs struct {
	Rules CustomRuleArrayInput `pulumi:"rules"`
}

func (CustomRuleListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleList)(nil)).Elem()
}

func (i CustomRuleListArgs) ToCustomRuleListOutput() CustomRuleListOutput {
	return i.ToCustomRuleListOutputWithContext(context.Background())
}

func (i CustomRuleListArgs) ToCustomRuleListOutputWithContext(ctx context.Context) CustomRuleListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleListOutput)
}

func (i CustomRuleListArgs) ToCustomRuleListPtrOutput() CustomRuleListPtrOutput {
	return i.ToCustomRuleListPtrOutputWithContext(context.Background())
}

func (i CustomRuleListArgs) ToCustomRuleListPtrOutputWithContext(ctx context.Context) CustomRuleListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleListOutput).ToCustomRuleListPtrOutputWithContext(ctx)
}









type CustomRuleListPtrInput interface {
	pulumi.Input

	ToCustomRuleListPtrOutput() CustomRuleListPtrOutput
	ToCustomRuleListPtrOutputWithContext(context.Context) CustomRuleListPtrOutput
}

type customRuleListPtrType CustomRuleListArgs

func CustomRuleListPtr(v *CustomRuleListArgs) CustomRuleListPtrInput {
	return (*customRuleListPtrType)(v)
}

func (*customRuleListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRuleList)(nil)).Elem()
}

func (i *customRuleListPtrType) ToCustomRuleListPtrOutput() CustomRuleListPtrOutput {
	return i.ToCustomRuleListPtrOutputWithContext(context.Background())
}

func (i *customRuleListPtrType) ToCustomRuleListPtrOutputWithContext(ctx context.Context) CustomRuleListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleListPtrOutput)
}

type CustomRuleListOutput struct{ *pulumi.OutputState }

func (CustomRuleListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleList)(nil)).Elem()
}

func (o CustomRuleListOutput) ToCustomRuleListOutput() CustomRuleListOutput {
	return o
}

func (o CustomRuleListOutput) ToCustomRuleListOutputWithContext(ctx context.Context) CustomRuleListOutput {
	return o
}

func (o CustomRuleListOutput) ToCustomRuleListPtrOutput() CustomRuleListPtrOutput {
	return o.ToCustomRuleListPtrOutputWithContext(context.Background())
}

func (o CustomRuleListOutput) ToCustomRuleListPtrOutputWithContext(ctx context.Context) CustomRuleListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomRuleList) *CustomRuleList {
		return &v
	}).(CustomRuleListPtrOutput)
}

func (o CustomRuleListOutput) Rules() CustomRuleArrayOutput {
	return o.ApplyT(func(v CustomRuleList) []CustomRule { return v.Rules }).(CustomRuleArrayOutput)
}

type CustomRuleListPtrOutput struct{ *pulumi.OutputState }

func (CustomRuleListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRuleList)(nil)).Elem()
}

func (o CustomRuleListPtrOutput) ToCustomRuleListPtrOutput() CustomRuleListPtrOutput {
	return o
}

func (o CustomRuleListPtrOutput) ToCustomRuleListPtrOutputWithContext(ctx context.Context) CustomRuleListPtrOutput {
	return o
}

func (o CustomRuleListPtrOutput) Elem() CustomRuleListOutput {
	return o.ApplyT(func(v *CustomRuleList) CustomRuleList {
		if v != nil {
			return *v
		}
		var ret CustomRuleList
		return ret
	}).(CustomRuleListOutput)
}

func (o CustomRuleListPtrOutput) Rules() CustomRuleArrayOutput {
	return o.ApplyT(func(v *CustomRuleList) []CustomRule {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(CustomRuleArrayOutput)
}

type CustomRuleListResponse struct {
	Rules []CustomRuleResponse `pulumi:"rules"`
}

type CustomRuleListResponseOutput struct{ *pulumi.OutputState }

func (CustomRuleListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleListResponse)(nil)).Elem()
}

func (o CustomRuleListResponseOutput) ToCustomRuleListResponseOutput() CustomRuleListResponseOutput {
	return o
}

func (o CustomRuleListResponseOutput) ToCustomRuleListResponseOutputWithContext(ctx context.Context) CustomRuleListResponseOutput {
	return o
}

func (o CustomRuleListResponseOutput) Rules() CustomRuleResponseArrayOutput {
	return o.ApplyT(func(v CustomRuleListResponse) []CustomRuleResponse { return v.Rules }).(CustomRuleResponseArrayOutput)
}

type CustomRuleListResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomRuleListResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRuleListResponse)(nil)).Elem()
}

func (o CustomRuleListResponsePtrOutput) ToCustomRuleListResponsePtrOutput() CustomRuleListResponsePtrOutput {
	return o
}

func (o CustomRuleListResponsePtrOutput) ToCustomRuleListResponsePtrOutputWithContext(ctx context.Context) CustomRuleListResponsePtrOutput {
	return o
}

func (o CustomRuleListResponsePtrOutput) Elem() CustomRuleListResponseOutput {
	return o.ApplyT(func(v *CustomRuleListResponse) CustomRuleListResponse {
		if v != nil {
			return *v
		}
		var ret CustomRuleListResponse
		return ret
	}).(CustomRuleListResponseOutput)
}

func (o CustomRuleListResponsePtrOutput) Rules() CustomRuleResponseArrayOutput {
	return o.ApplyT(func(v *CustomRuleListResponse) []CustomRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(CustomRuleResponseArrayOutput)
}

type CustomRuleResponse struct {
	Action                     string                            `pulumi:"action"`
	EnabledState               *string                           `pulumi:"enabledState"`
	MatchConditions            []FrontDoorMatchConditionResponse `pulumi:"matchConditions"`
	Name                       *string                           `pulumi:"name"`
	Priority                   int                               `pulumi:"priority"`
	RateLimitDurationInMinutes *int                              `pulumi:"rateLimitDurationInMinutes"`
	RateLimitThreshold         *int                              `pulumi:"rateLimitThreshold"`
	RuleType                   string                            `pulumi:"ruleType"`
}

type CustomRuleResponseOutput struct{ *pulumi.OutputState }

func (CustomRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleResponse)(nil)).Elem()
}

func (o CustomRuleResponseOutput) ToCustomRuleResponseOutput() CustomRuleResponseOutput {
	return o
}

func (o CustomRuleResponseOutput) ToCustomRuleResponseOutputWithContext(ctx context.Context) CustomRuleResponseOutput {
	return o
}

func (o CustomRuleResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRuleResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o CustomRuleResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRuleResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o CustomRuleResponseOutput) MatchConditions() FrontDoorMatchConditionResponseArrayOutput {
	return o.ApplyT(func(v CustomRuleResponse) []FrontDoorMatchConditionResponse { return v.MatchConditions }).(FrontDoorMatchConditionResponseArrayOutput)
}

func (o CustomRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CustomRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v CustomRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

func (o CustomRuleResponseOutput) RateLimitDurationInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CustomRuleResponse) *int { return v.RateLimitDurationInMinutes }).(pulumi.IntPtrOutput)
}

func (o CustomRuleResponseOutput) RateLimitThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CustomRuleResponse) *int { return v.RateLimitThreshold }).(pulumi.IntPtrOutput)
}

func (o CustomRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

type CustomRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (CustomRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRuleResponse)(nil)).Elem()
}

func (o CustomRuleResponseArrayOutput) ToCustomRuleResponseArrayOutput() CustomRuleResponseArrayOutput {
	return o
}

func (o CustomRuleResponseArrayOutput) ToCustomRuleResponseArrayOutputWithContext(ctx context.Context) CustomRuleResponseArrayOutput {
	return o
}

func (o CustomRuleResponseArrayOutput) Index(i pulumi.IntInput) CustomRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomRuleResponse {
		return vs[0].([]CustomRuleResponse)[vs[1].(int)]
	}).(CustomRuleResponseOutput)
}

type ForwardingConfiguration struct {
	BackendPool          *SubResource        `pulumi:"backendPool"`
	CacheConfiguration   *CacheConfiguration `pulumi:"cacheConfiguration"`
	CustomForwardingPath *string             `pulumi:"customForwardingPath"`
	ForwardingProtocol   *string             `pulumi:"forwardingProtocol"`
	OdataType            string              `pulumi:"odataType"`
}

type ForwardingConfigurationResponse struct {
	BackendPool          *SubResourceResponse        `pulumi:"backendPool"`
	CacheConfiguration   *CacheConfigurationResponse `pulumi:"cacheConfiguration"`
	CustomForwardingPath *string                     `pulumi:"customForwardingPath"`
	ForwardingProtocol   *string                     `pulumi:"forwardingProtocol"`
	OdataType            string                      `pulumi:"odataType"`
}

type FrontDoorManagedRuleGroupOverride struct {
	Exclusions    []ManagedRuleExclusion         `pulumi:"exclusions"`
	RuleGroupName string                         `pulumi:"ruleGroupName"`
	Rules         []FrontDoorManagedRuleOverride `pulumi:"rules"`
}





type FrontDoorManagedRuleGroupOverrideInput interface {
	pulumi.Input

	ToFrontDoorManagedRuleGroupOverrideOutput() FrontDoorManagedRuleGroupOverrideOutput
	ToFrontDoorManagedRuleGroupOverrideOutputWithContext(context.Context) FrontDoorManagedRuleGroupOverrideOutput
}

type FrontDoorManagedRuleGroupOverrideArgs struct {
	Exclusions    ManagedRuleExclusionArrayInput         `pulumi:"exclusions"`
	RuleGroupName pulumi.StringInput                     `pulumi:"ruleGroupName"`
	Rules         FrontDoorManagedRuleOverrideArrayInput `pulumi:"rules"`
}

func (FrontDoorManagedRuleGroupOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleGroupOverride)(nil)).Elem()
}

func (i FrontDoorManagedRuleGroupOverrideArgs) ToFrontDoorManagedRuleGroupOverrideOutput() FrontDoorManagedRuleGroupOverrideOutput {
	return i.ToFrontDoorManagedRuleGroupOverrideOutputWithContext(context.Background())
}

func (i FrontDoorManagedRuleGroupOverrideArgs) ToFrontDoorManagedRuleGroupOverrideOutputWithContext(ctx context.Context) FrontDoorManagedRuleGroupOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorManagedRuleGroupOverrideOutput)
}





type FrontDoorManagedRuleGroupOverrideArrayInput interface {
	pulumi.Input

	ToFrontDoorManagedRuleGroupOverrideArrayOutput() FrontDoorManagedRuleGroupOverrideArrayOutput
	ToFrontDoorManagedRuleGroupOverrideArrayOutputWithContext(context.Context) FrontDoorManagedRuleGroupOverrideArrayOutput
}

type FrontDoorManagedRuleGroupOverrideArray []FrontDoorManagedRuleGroupOverrideInput

func (FrontDoorManagedRuleGroupOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleGroupOverride)(nil)).Elem()
}

func (i FrontDoorManagedRuleGroupOverrideArray) ToFrontDoorManagedRuleGroupOverrideArrayOutput() FrontDoorManagedRuleGroupOverrideArrayOutput {
	return i.ToFrontDoorManagedRuleGroupOverrideArrayOutputWithContext(context.Background())
}

func (i FrontDoorManagedRuleGroupOverrideArray) ToFrontDoorManagedRuleGroupOverrideArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleGroupOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorManagedRuleGroupOverrideArrayOutput)
}

type FrontDoorManagedRuleGroupOverrideOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleGroupOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleGroupOverride)(nil)).Elem()
}

func (o FrontDoorManagedRuleGroupOverrideOutput) ToFrontDoorManagedRuleGroupOverrideOutput() FrontDoorManagedRuleGroupOverrideOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideOutput) ToFrontDoorManagedRuleGroupOverrideOutputWithContext(ctx context.Context) FrontDoorManagedRuleGroupOverrideOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideOutput) Exclusions() ManagedRuleExclusionArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleGroupOverride) []ManagedRuleExclusion { return v.Exclusions }).(ManagedRuleExclusionArrayOutput)
}

func (o FrontDoorManagedRuleGroupOverrideOutput) RuleGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleGroupOverride) string { return v.RuleGroupName }).(pulumi.StringOutput)
}

func (o FrontDoorManagedRuleGroupOverrideOutput) Rules() FrontDoorManagedRuleOverrideArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleGroupOverride) []FrontDoorManagedRuleOverride { return v.Rules }).(FrontDoorManagedRuleOverrideArrayOutput)
}

type FrontDoorManagedRuleGroupOverrideArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleGroupOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleGroupOverride)(nil)).Elem()
}

func (o FrontDoorManagedRuleGroupOverrideArrayOutput) ToFrontDoorManagedRuleGroupOverrideArrayOutput() FrontDoorManagedRuleGroupOverrideArrayOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideArrayOutput) ToFrontDoorManagedRuleGroupOverrideArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleGroupOverrideArrayOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideArrayOutput) Index(i pulumi.IntInput) FrontDoorManagedRuleGroupOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorManagedRuleGroupOverride {
		return vs[0].([]FrontDoorManagedRuleGroupOverride)[vs[1].(int)]
	}).(FrontDoorManagedRuleGroupOverrideOutput)
}

type FrontDoorManagedRuleGroupOverrideResponse struct {
	Exclusions    []ManagedRuleExclusionResponse         `pulumi:"exclusions"`
	RuleGroupName string                                 `pulumi:"ruleGroupName"`
	Rules         []FrontDoorManagedRuleOverrideResponse `pulumi:"rules"`
}

type FrontDoorManagedRuleGroupOverrideResponseOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleGroupOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleGroupOverrideResponse)(nil)).Elem()
}

func (o FrontDoorManagedRuleGroupOverrideResponseOutput) ToFrontDoorManagedRuleGroupOverrideResponseOutput() FrontDoorManagedRuleGroupOverrideResponseOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideResponseOutput) ToFrontDoorManagedRuleGroupOverrideResponseOutputWithContext(ctx context.Context) FrontDoorManagedRuleGroupOverrideResponseOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideResponseOutput) Exclusions() ManagedRuleExclusionResponseArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleGroupOverrideResponse) []ManagedRuleExclusionResponse { return v.Exclusions }).(ManagedRuleExclusionResponseArrayOutput)
}

func (o FrontDoorManagedRuleGroupOverrideResponseOutput) RuleGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleGroupOverrideResponse) string { return v.RuleGroupName }).(pulumi.StringOutput)
}

func (o FrontDoorManagedRuleGroupOverrideResponseOutput) Rules() FrontDoorManagedRuleOverrideResponseArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleGroupOverrideResponse) []FrontDoorManagedRuleOverrideResponse {
		return v.Rules
	}).(FrontDoorManagedRuleOverrideResponseArrayOutput)
}

type FrontDoorManagedRuleGroupOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleGroupOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleGroupOverrideResponse)(nil)).Elem()
}

func (o FrontDoorManagedRuleGroupOverrideResponseArrayOutput) ToFrontDoorManagedRuleGroupOverrideResponseArrayOutput() FrontDoorManagedRuleGroupOverrideResponseArrayOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideResponseArrayOutput) ToFrontDoorManagedRuleGroupOverrideResponseArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleGroupOverrideResponseArrayOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideResponseArrayOutput) Index(i pulumi.IntInput) FrontDoorManagedRuleGroupOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorManagedRuleGroupOverrideResponse {
		return vs[0].([]FrontDoorManagedRuleGroupOverrideResponse)[vs[1].(int)]
	}).(FrontDoorManagedRuleGroupOverrideResponseOutput)
}

type FrontDoorManagedRuleOverride struct {
	Action       *string                `pulumi:"action"`
	EnabledState *string                `pulumi:"enabledState"`
	Exclusions   []ManagedRuleExclusion `pulumi:"exclusions"`
	RuleId       string                 `pulumi:"ruleId"`
}





type FrontDoorManagedRuleOverrideInput interface {
	pulumi.Input

	ToFrontDoorManagedRuleOverrideOutput() FrontDoorManagedRuleOverrideOutput
	ToFrontDoorManagedRuleOverrideOutputWithContext(context.Context) FrontDoorManagedRuleOverrideOutput
}

type FrontDoorManagedRuleOverrideArgs struct {
	Action       pulumi.StringPtrInput          `pulumi:"action"`
	EnabledState pulumi.StringPtrInput          `pulumi:"enabledState"`
	Exclusions   ManagedRuleExclusionArrayInput `pulumi:"exclusions"`
	RuleId       pulumi.StringInput             `pulumi:"ruleId"`
}

func (FrontDoorManagedRuleOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleOverride)(nil)).Elem()
}

func (i FrontDoorManagedRuleOverrideArgs) ToFrontDoorManagedRuleOverrideOutput() FrontDoorManagedRuleOverrideOutput {
	return i.ToFrontDoorManagedRuleOverrideOutputWithContext(context.Background())
}

func (i FrontDoorManagedRuleOverrideArgs) ToFrontDoorManagedRuleOverrideOutputWithContext(ctx context.Context) FrontDoorManagedRuleOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorManagedRuleOverrideOutput)
}





type FrontDoorManagedRuleOverrideArrayInput interface {
	pulumi.Input

	ToFrontDoorManagedRuleOverrideArrayOutput() FrontDoorManagedRuleOverrideArrayOutput
	ToFrontDoorManagedRuleOverrideArrayOutputWithContext(context.Context) FrontDoorManagedRuleOverrideArrayOutput
}

type FrontDoorManagedRuleOverrideArray []FrontDoorManagedRuleOverrideInput

func (FrontDoorManagedRuleOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleOverride)(nil)).Elem()
}

func (i FrontDoorManagedRuleOverrideArray) ToFrontDoorManagedRuleOverrideArrayOutput() FrontDoorManagedRuleOverrideArrayOutput {
	return i.ToFrontDoorManagedRuleOverrideArrayOutputWithContext(context.Background())
}

func (i FrontDoorManagedRuleOverrideArray) ToFrontDoorManagedRuleOverrideArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorManagedRuleOverrideArrayOutput)
}

type FrontDoorManagedRuleOverrideOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleOverride)(nil)).Elem()
}

func (o FrontDoorManagedRuleOverrideOutput) ToFrontDoorManagedRuleOverrideOutput() FrontDoorManagedRuleOverrideOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideOutput) ToFrontDoorManagedRuleOverrideOutputWithContext(ctx context.Context) FrontDoorManagedRuleOverrideOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverride) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o FrontDoorManagedRuleOverrideOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverride) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontDoorManagedRuleOverrideOutput) Exclusions() ManagedRuleExclusionArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverride) []ManagedRuleExclusion { return v.Exclusions }).(ManagedRuleExclusionArrayOutput)
}

func (o FrontDoorManagedRuleOverrideOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverride) string { return v.RuleId }).(pulumi.StringOutput)
}

type FrontDoorManagedRuleOverrideArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleOverride)(nil)).Elem()
}

func (o FrontDoorManagedRuleOverrideArrayOutput) ToFrontDoorManagedRuleOverrideArrayOutput() FrontDoorManagedRuleOverrideArrayOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideArrayOutput) ToFrontDoorManagedRuleOverrideArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleOverrideArrayOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideArrayOutput) Index(i pulumi.IntInput) FrontDoorManagedRuleOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorManagedRuleOverride {
		return vs[0].([]FrontDoorManagedRuleOverride)[vs[1].(int)]
	}).(FrontDoorManagedRuleOverrideOutput)
}

type FrontDoorManagedRuleOverrideResponse struct {
	Action       *string                        `pulumi:"action"`
	EnabledState *string                        `pulumi:"enabledState"`
	Exclusions   []ManagedRuleExclusionResponse `pulumi:"exclusions"`
	RuleId       string                         `pulumi:"ruleId"`
}

type FrontDoorManagedRuleOverrideResponseOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleOverrideResponse)(nil)).Elem()
}

func (o FrontDoorManagedRuleOverrideResponseOutput) ToFrontDoorManagedRuleOverrideResponseOutput() FrontDoorManagedRuleOverrideResponseOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideResponseOutput) ToFrontDoorManagedRuleOverrideResponseOutputWithContext(ctx context.Context) FrontDoorManagedRuleOverrideResponseOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverrideResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o FrontDoorManagedRuleOverrideResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverrideResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontDoorManagedRuleOverrideResponseOutput) Exclusions() ManagedRuleExclusionResponseArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverrideResponse) []ManagedRuleExclusionResponse { return v.Exclusions }).(ManagedRuleExclusionResponseArrayOutput)
}

func (o FrontDoorManagedRuleOverrideResponseOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverrideResponse) string { return v.RuleId }).(pulumi.StringOutput)
}

type FrontDoorManagedRuleOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleOverrideResponse)(nil)).Elem()
}

func (o FrontDoorManagedRuleOverrideResponseArrayOutput) ToFrontDoorManagedRuleOverrideResponseArrayOutput() FrontDoorManagedRuleOverrideResponseArrayOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideResponseArrayOutput) ToFrontDoorManagedRuleOverrideResponseArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleOverrideResponseArrayOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideResponseArrayOutput) Index(i pulumi.IntInput) FrontDoorManagedRuleOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorManagedRuleOverrideResponse {
		return vs[0].([]FrontDoorManagedRuleOverrideResponse)[vs[1].(int)]
	}).(FrontDoorManagedRuleOverrideResponseOutput)
}

type FrontDoorManagedRuleSet struct {
	Exclusions         []ManagedRuleExclusion              `pulumi:"exclusions"`
	RuleGroupOverrides []FrontDoorManagedRuleGroupOverride `pulumi:"ruleGroupOverrides"`
	RuleSetAction      *string                             `pulumi:"ruleSetAction"`
	RuleSetType        string                              `pulumi:"ruleSetType"`
	RuleSetVersion     string                              `pulumi:"ruleSetVersion"`
}





type FrontDoorManagedRuleSetInput interface {
	pulumi.Input

	ToFrontDoorManagedRuleSetOutput() FrontDoorManagedRuleSetOutput
	ToFrontDoorManagedRuleSetOutputWithContext(context.Context) FrontDoorManagedRuleSetOutput
}

type FrontDoorManagedRuleSetArgs struct {
	Exclusions         ManagedRuleExclusionArrayInput              `pulumi:"exclusions"`
	RuleGroupOverrides FrontDoorManagedRuleGroupOverrideArrayInput `pulumi:"ruleGroupOverrides"`
	RuleSetAction      pulumi.StringPtrInput                       `pulumi:"ruleSetAction"`
	RuleSetType        pulumi.StringInput                          `pulumi:"ruleSetType"`
	RuleSetVersion     pulumi.StringInput                          `pulumi:"ruleSetVersion"`
}

func (FrontDoorManagedRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleSet)(nil)).Elem()
}

func (i FrontDoorManagedRuleSetArgs) ToFrontDoorManagedRuleSetOutput() FrontDoorManagedRuleSetOutput {
	return i.ToFrontDoorManagedRuleSetOutputWithContext(context.Background())
}

func (i FrontDoorManagedRuleSetArgs) ToFrontDoorManagedRuleSetOutputWithContext(ctx context.Context) FrontDoorManagedRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorManagedRuleSetOutput)
}





type FrontDoorManagedRuleSetArrayInput interface {
	pulumi.Input

	ToFrontDoorManagedRuleSetArrayOutput() FrontDoorManagedRuleSetArrayOutput
	ToFrontDoorManagedRuleSetArrayOutputWithContext(context.Context) FrontDoorManagedRuleSetArrayOutput
}

type FrontDoorManagedRuleSetArray []FrontDoorManagedRuleSetInput

func (FrontDoorManagedRuleSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleSet)(nil)).Elem()
}

func (i FrontDoorManagedRuleSetArray) ToFrontDoorManagedRuleSetArrayOutput() FrontDoorManagedRuleSetArrayOutput {
	return i.ToFrontDoorManagedRuleSetArrayOutputWithContext(context.Background())
}

func (i FrontDoorManagedRuleSetArray) ToFrontDoorManagedRuleSetArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorManagedRuleSetArrayOutput)
}

type FrontDoorManagedRuleSetOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleSet)(nil)).Elem()
}

func (o FrontDoorManagedRuleSetOutput) ToFrontDoorManagedRuleSetOutput() FrontDoorManagedRuleSetOutput {
	return o
}

func (o FrontDoorManagedRuleSetOutput) ToFrontDoorManagedRuleSetOutputWithContext(ctx context.Context) FrontDoorManagedRuleSetOutput {
	return o
}

func (o FrontDoorManagedRuleSetOutput) Exclusions() ManagedRuleExclusionArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSet) []ManagedRuleExclusion { return v.Exclusions }).(ManagedRuleExclusionArrayOutput)
}

func (o FrontDoorManagedRuleSetOutput) RuleGroupOverrides() FrontDoorManagedRuleGroupOverrideArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSet) []FrontDoorManagedRuleGroupOverride { return v.RuleGroupOverrides }).(FrontDoorManagedRuleGroupOverrideArrayOutput)
}

func (o FrontDoorManagedRuleSetOutput) RuleSetAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSet) *string { return v.RuleSetAction }).(pulumi.StringPtrOutput)
}

func (o FrontDoorManagedRuleSetOutput) RuleSetType() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSet) string { return v.RuleSetType }).(pulumi.StringOutput)
}

func (o FrontDoorManagedRuleSetOutput) RuleSetVersion() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSet) string { return v.RuleSetVersion }).(pulumi.StringOutput)
}

type FrontDoorManagedRuleSetArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleSet)(nil)).Elem()
}

func (o FrontDoorManagedRuleSetArrayOutput) ToFrontDoorManagedRuleSetArrayOutput() FrontDoorManagedRuleSetArrayOutput {
	return o
}

func (o FrontDoorManagedRuleSetArrayOutput) ToFrontDoorManagedRuleSetArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleSetArrayOutput {
	return o
}

func (o FrontDoorManagedRuleSetArrayOutput) Index(i pulumi.IntInput) FrontDoorManagedRuleSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorManagedRuleSet {
		return vs[0].([]FrontDoorManagedRuleSet)[vs[1].(int)]
	}).(FrontDoorManagedRuleSetOutput)
}

type FrontDoorManagedRuleSetResponse struct {
	Exclusions         []ManagedRuleExclusionResponse              `pulumi:"exclusions"`
	RuleGroupOverrides []FrontDoorManagedRuleGroupOverrideResponse `pulumi:"ruleGroupOverrides"`
	RuleSetAction      *string                                     `pulumi:"ruleSetAction"`
	RuleSetType        string                                      `pulumi:"ruleSetType"`
	RuleSetVersion     string                                      `pulumi:"ruleSetVersion"`
}

type FrontDoorManagedRuleSetResponseOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleSetResponse)(nil)).Elem()
}

func (o FrontDoorManagedRuleSetResponseOutput) ToFrontDoorManagedRuleSetResponseOutput() FrontDoorManagedRuleSetResponseOutput {
	return o
}

func (o FrontDoorManagedRuleSetResponseOutput) ToFrontDoorManagedRuleSetResponseOutputWithContext(ctx context.Context) FrontDoorManagedRuleSetResponseOutput {
	return o
}

func (o FrontDoorManagedRuleSetResponseOutput) Exclusions() ManagedRuleExclusionResponseArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSetResponse) []ManagedRuleExclusionResponse { return v.Exclusions }).(ManagedRuleExclusionResponseArrayOutput)
}

func (o FrontDoorManagedRuleSetResponseOutput) RuleGroupOverrides() FrontDoorManagedRuleGroupOverrideResponseArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSetResponse) []FrontDoorManagedRuleGroupOverrideResponse {
		return v.RuleGroupOverrides
	}).(FrontDoorManagedRuleGroupOverrideResponseArrayOutput)
}

func (o FrontDoorManagedRuleSetResponseOutput) RuleSetAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSetResponse) *string { return v.RuleSetAction }).(pulumi.StringPtrOutput)
}

func (o FrontDoorManagedRuleSetResponseOutput) RuleSetType() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSetResponse) string { return v.RuleSetType }).(pulumi.StringOutput)
}

func (o FrontDoorManagedRuleSetResponseOutput) RuleSetVersion() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSetResponse) string { return v.RuleSetVersion }).(pulumi.StringOutput)
}

type FrontDoorManagedRuleSetResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleSetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleSetResponse)(nil)).Elem()
}

func (o FrontDoorManagedRuleSetResponseArrayOutput) ToFrontDoorManagedRuleSetResponseArrayOutput() FrontDoorManagedRuleSetResponseArrayOutput {
	return o
}

func (o FrontDoorManagedRuleSetResponseArrayOutput) ToFrontDoorManagedRuleSetResponseArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleSetResponseArrayOutput {
	return o
}

func (o FrontDoorManagedRuleSetResponseArrayOutput) Index(i pulumi.IntInput) FrontDoorManagedRuleSetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorManagedRuleSetResponse {
		return vs[0].([]FrontDoorManagedRuleSetResponse)[vs[1].(int)]
	}).(FrontDoorManagedRuleSetResponseOutput)
}

type FrontDoorMatchCondition struct {
	MatchValue      []string `pulumi:"matchValue"`
	MatchVariable   string   `pulumi:"matchVariable"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}





type FrontDoorMatchConditionInput interface {
	pulumi.Input

	ToFrontDoorMatchConditionOutput() FrontDoorMatchConditionOutput
	ToFrontDoorMatchConditionOutputWithContext(context.Context) FrontDoorMatchConditionOutput
}

type FrontDoorMatchConditionArgs struct {
	MatchValue      pulumi.StringArrayInput `pulumi:"matchValue"`
	MatchVariable   pulumi.StringInput      `pulumi:"matchVariable"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Selector        pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (FrontDoorMatchConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorMatchCondition)(nil)).Elem()
}

func (i FrontDoorMatchConditionArgs) ToFrontDoorMatchConditionOutput() FrontDoorMatchConditionOutput {
	return i.ToFrontDoorMatchConditionOutputWithContext(context.Background())
}

func (i FrontDoorMatchConditionArgs) ToFrontDoorMatchConditionOutputWithContext(ctx context.Context) FrontDoorMatchConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorMatchConditionOutput)
}





type FrontDoorMatchConditionArrayInput interface {
	pulumi.Input

	ToFrontDoorMatchConditionArrayOutput() FrontDoorMatchConditionArrayOutput
	ToFrontDoorMatchConditionArrayOutputWithContext(context.Context) FrontDoorMatchConditionArrayOutput
}

type FrontDoorMatchConditionArray []FrontDoorMatchConditionInput

func (FrontDoorMatchConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorMatchCondition)(nil)).Elem()
}

func (i FrontDoorMatchConditionArray) ToFrontDoorMatchConditionArrayOutput() FrontDoorMatchConditionArrayOutput {
	return i.ToFrontDoorMatchConditionArrayOutputWithContext(context.Background())
}

func (i FrontDoorMatchConditionArray) ToFrontDoorMatchConditionArrayOutputWithContext(ctx context.Context) FrontDoorMatchConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorMatchConditionArrayOutput)
}

type FrontDoorMatchConditionOutput struct{ *pulumi.OutputState }

func (FrontDoorMatchConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorMatchCondition)(nil)).Elem()
}

func (o FrontDoorMatchConditionOutput) ToFrontDoorMatchConditionOutput() FrontDoorMatchConditionOutput {
	return o
}

func (o FrontDoorMatchConditionOutput) ToFrontDoorMatchConditionOutputWithContext(ctx context.Context) FrontDoorMatchConditionOutput {
	return o
}

func (o FrontDoorMatchConditionOutput) MatchValue() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FrontDoorMatchCondition) []string { return v.MatchValue }).(pulumi.StringArrayOutput)
}

func (o FrontDoorMatchConditionOutput) MatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorMatchCondition) string { return v.MatchVariable }).(pulumi.StringOutput)
}

func (o FrontDoorMatchConditionOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FrontDoorMatchCondition) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o FrontDoorMatchConditionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorMatchCondition) string { return v.Operator }).(pulumi.StringOutput)
}

func (o FrontDoorMatchConditionOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorMatchCondition) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o FrontDoorMatchConditionOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FrontDoorMatchCondition) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type FrontDoorMatchConditionArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorMatchConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorMatchCondition)(nil)).Elem()
}

func (o FrontDoorMatchConditionArrayOutput) ToFrontDoorMatchConditionArrayOutput() FrontDoorMatchConditionArrayOutput {
	return o
}

func (o FrontDoorMatchConditionArrayOutput) ToFrontDoorMatchConditionArrayOutputWithContext(ctx context.Context) FrontDoorMatchConditionArrayOutput {
	return o
}

func (o FrontDoorMatchConditionArrayOutput) Index(i pulumi.IntInput) FrontDoorMatchConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorMatchCondition {
		return vs[0].([]FrontDoorMatchCondition)[vs[1].(int)]
	}).(FrontDoorMatchConditionOutput)
}

type FrontDoorMatchConditionResponse struct {
	MatchValue      []string `pulumi:"matchValue"`
	MatchVariable   string   `pulumi:"matchVariable"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}

type FrontDoorMatchConditionResponseOutput struct{ *pulumi.OutputState }

func (FrontDoorMatchConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorMatchConditionResponse)(nil)).Elem()
}

func (o FrontDoorMatchConditionResponseOutput) ToFrontDoorMatchConditionResponseOutput() FrontDoorMatchConditionResponseOutput {
	return o
}

func (o FrontDoorMatchConditionResponseOutput) ToFrontDoorMatchConditionResponseOutputWithContext(ctx context.Context) FrontDoorMatchConditionResponseOutput {
	return o
}

func (o FrontDoorMatchConditionResponseOutput) MatchValue() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FrontDoorMatchConditionResponse) []string { return v.MatchValue }).(pulumi.StringArrayOutput)
}

func (o FrontDoorMatchConditionResponseOutput) MatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorMatchConditionResponse) string { return v.MatchVariable }).(pulumi.StringOutput)
}

func (o FrontDoorMatchConditionResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FrontDoorMatchConditionResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o FrontDoorMatchConditionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorMatchConditionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o FrontDoorMatchConditionResponseOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorMatchConditionResponse) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o FrontDoorMatchConditionResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FrontDoorMatchConditionResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type FrontDoorMatchConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorMatchConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorMatchConditionResponse)(nil)).Elem()
}

func (o FrontDoorMatchConditionResponseArrayOutput) ToFrontDoorMatchConditionResponseArrayOutput() FrontDoorMatchConditionResponseArrayOutput {
	return o
}

func (o FrontDoorMatchConditionResponseArrayOutput) ToFrontDoorMatchConditionResponseArrayOutputWithContext(ctx context.Context) FrontDoorMatchConditionResponseArrayOutput {
	return o
}

func (o FrontDoorMatchConditionResponseArrayOutput) Index(i pulumi.IntInput) FrontDoorMatchConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorMatchConditionResponse {
		return vs[0].([]FrontDoorMatchConditionResponse)[vs[1].(int)]
	}).(FrontDoorMatchConditionResponseOutput)
}

type FrontDoorPolicySettings struct {
	CustomBlockResponseBody       *string `pulumi:"customBlockResponseBody"`
	CustomBlockResponseStatusCode *int    `pulumi:"customBlockResponseStatusCode"`
	EnabledState                  *string `pulumi:"enabledState"`
	Mode                          *string `pulumi:"mode"`
	RedirectUrl                   *string `pulumi:"redirectUrl"`
	RequestBodyCheck              *string `pulumi:"requestBodyCheck"`
}





type FrontDoorPolicySettingsInput interface {
	pulumi.Input

	ToFrontDoorPolicySettingsOutput() FrontDoorPolicySettingsOutput
	ToFrontDoorPolicySettingsOutputWithContext(context.Context) FrontDoorPolicySettingsOutput
}

type FrontDoorPolicySettingsArgs struct {
	CustomBlockResponseBody       pulumi.StringPtrInput `pulumi:"customBlockResponseBody"`
	CustomBlockResponseStatusCode pulumi.IntPtrInput    `pulumi:"customBlockResponseStatusCode"`
	EnabledState                  pulumi.StringPtrInput `pulumi:"enabledState"`
	Mode                          pulumi.StringPtrInput `pulumi:"mode"`
	RedirectUrl                   pulumi.StringPtrInput `pulumi:"redirectUrl"`
	RequestBodyCheck              pulumi.StringPtrInput `pulumi:"requestBodyCheck"`
}

func (FrontDoorPolicySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorPolicySettings)(nil)).Elem()
}

func (i FrontDoorPolicySettingsArgs) ToFrontDoorPolicySettingsOutput() FrontDoorPolicySettingsOutput {
	return i.ToFrontDoorPolicySettingsOutputWithContext(context.Background())
}

func (i FrontDoorPolicySettingsArgs) ToFrontDoorPolicySettingsOutputWithContext(ctx context.Context) FrontDoorPolicySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorPolicySettingsOutput)
}

func (i FrontDoorPolicySettingsArgs) ToFrontDoorPolicySettingsPtrOutput() FrontDoorPolicySettingsPtrOutput {
	return i.ToFrontDoorPolicySettingsPtrOutputWithContext(context.Background())
}

func (i FrontDoorPolicySettingsArgs) ToFrontDoorPolicySettingsPtrOutputWithContext(ctx context.Context) FrontDoorPolicySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorPolicySettingsOutput).ToFrontDoorPolicySettingsPtrOutputWithContext(ctx)
}









type FrontDoorPolicySettingsPtrInput interface {
	pulumi.Input

	ToFrontDoorPolicySettingsPtrOutput() FrontDoorPolicySettingsPtrOutput
	ToFrontDoorPolicySettingsPtrOutputWithContext(context.Context) FrontDoorPolicySettingsPtrOutput
}

type frontDoorPolicySettingsPtrType FrontDoorPolicySettingsArgs

func FrontDoorPolicySettingsPtr(v *FrontDoorPolicySettingsArgs) FrontDoorPolicySettingsPtrInput {
	return (*frontDoorPolicySettingsPtrType)(v)
}

func (*frontDoorPolicySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontDoorPolicySettings)(nil)).Elem()
}

func (i *frontDoorPolicySettingsPtrType) ToFrontDoorPolicySettingsPtrOutput() FrontDoorPolicySettingsPtrOutput {
	return i.ToFrontDoorPolicySettingsPtrOutputWithContext(context.Background())
}

func (i *frontDoorPolicySettingsPtrType) ToFrontDoorPolicySettingsPtrOutputWithContext(ctx context.Context) FrontDoorPolicySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorPolicySettingsPtrOutput)
}

type FrontDoorPolicySettingsOutput struct{ *pulumi.OutputState }

func (FrontDoorPolicySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorPolicySettings)(nil)).Elem()
}

func (o FrontDoorPolicySettingsOutput) ToFrontDoorPolicySettingsOutput() FrontDoorPolicySettingsOutput {
	return o
}

func (o FrontDoorPolicySettingsOutput) ToFrontDoorPolicySettingsOutputWithContext(ctx context.Context) FrontDoorPolicySettingsOutput {
	return o
}

func (o FrontDoorPolicySettingsOutput) ToFrontDoorPolicySettingsPtrOutput() FrontDoorPolicySettingsPtrOutput {
	return o.ToFrontDoorPolicySettingsPtrOutputWithContext(context.Background())
}

func (o FrontDoorPolicySettingsOutput) ToFrontDoorPolicySettingsPtrOutputWithContext(ctx context.Context) FrontDoorPolicySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrontDoorPolicySettings) *FrontDoorPolicySettings {
		return &v
	}).(FrontDoorPolicySettingsPtrOutput)
}

func (o FrontDoorPolicySettingsOutput) CustomBlockResponseBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettings) *string { return v.CustomBlockResponseBody }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsOutput) CustomBlockResponseStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettings) *int { return v.CustomBlockResponseStatusCode }).(pulumi.IntPtrOutput)
}

func (o FrontDoorPolicySettingsOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettings) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettings) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettings) *string { return v.RedirectUrl }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsOutput) RequestBodyCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettings) *string { return v.RequestBodyCheck }).(pulumi.StringPtrOutput)
}

type FrontDoorPolicySettingsPtrOutput struct{ *pulumi.OutputState }

func (FrontDoorPolicySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontDoorPolicySettings)(nil)).Elem()
}

func (o FrontDoorPolicySettingsPtrOutput) ToFrontDoorPolicySettingsPtrOutput() FrontDoorPolicySettingsPtrOutput {
	return o
}

func (o FrontDoorPolicySettingsPtrOutput) ToFrontDoorPolicySettingsPtrOutputWithContext(ctx context.Context) FrontDoorPolicySettingsPtrOutput {
	return o
}

func (o FrontDoorPolicySettingsPtrOutput) Elem() FrontDoorPolicySettingsOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) FrontDoorPolicySettings {
		if v != nil {
			return *v
		}
		var ret FrontDoorPolicySettings
		return ret
	}).(FrontDoorPolicySettingsOutput)
}

func (o FrontDoorPolicySettingsPtrOutput) CustomBlockResponseBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.CustomBlockResponseBody
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsPtrOutput) CustomBlockResponseStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) *int {
		if v == nil {
			return nil
		}
		return v.CustomBlockResponseStatusCode
	}).(pulumi.IntPtrOutput)
}

func (o FrontDoorPolicySettingsPtrOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.EnabledState
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsPtrOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.RedirectUrl
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsPtrOutput) RequestBodyCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.RequestBodyCheck
	}).(pulumi.StringPtrOutput)
}

type FrontDoorPolicySettingsResponse struct {
	CustomBlockResponseBody       *string `pulumi:"customBlockResponseBody"`
	CustomBlockResponseStatusCode *int    `pulumi:"customBlockResponseStatusCode"`
	EnabledState                  *string `pulumi:"enabledState"`
	Mode                          *string `pulumi:"mode"`
	RedirectUrl                   *string `pulumi:"redirectUrl"`
	RequestBodyCheck              *string `pulumi:"requestBodyCheck"`
}

type FrontDoorPolicySettingsResponseOutput struct{ *pulumi.OutputState }

func (FrontDoorPolicySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorPolicySettingsResponse)(nil)).Elem()
}

func (o FrontDoorPolicySettingsResponseOutput) ToFrontDoorPolicySettingsResponseOutput() FrontDoorPolicySettingsResponseOutput {
	return o
}

func (o FrontDoorPolicySettingsResponseOutput) ToFrontDoorPolicySettingsResponseOutputWithContext(ctx context.Context) FrontDoorPolicySettingsResponseOutput {
	return o
}

func (o FrontDoorPolicySettingsResponseOutput) CustomBlockResponseBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettingsResponse) *string { return v.CustomBlockResponseBody }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponseOutput) CustomBlockResponseStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettingsResponse) *int { return v.CustomBlockResponseStatusCode }).(pulumi.IntPtrOutput)
}

func (o FrontDoorPolicySettingsResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettingsResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettingsResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponseOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettingsResponse) *string { return v.RedirectUrl }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponseOutput) RequestBodyCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettingsResponse) *string { return v.RequestBodyCheck }).(pulumi.StringPtrOutput)
}

type FrontDoorPolicySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (FrontDoorPolicySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontDoorPolicySettingsResponse)(nil)).Elem()
}

func (o FrontDoorPolicySettingsResponsePtrOutput) ToFrontDoorPolicySettingsResponsePtrOutput() FrontDoorPolicySettingsResponsePtrOutput {
	return o
}

func (o FrontDoorPolicySettingsResponsePtrOutput) ToFrontDoorPolicySettingsResponsePtrOutputWithContext(ctx context.Context) FrontDoorPolicySettingsResponsePtrOutput {
	return o
}

func (o FrontDoorPolicySettingsResponsePtrOutput) Elem() FrontDoorPolicySettingsResponseOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) FrontDoorPolicySettingsResponse {
		if v != nil {
			return *v
		}
		var ret FrontDoorPolicySettingsResponse
		return ret
	}).(FrontDoorPolicySettingsResponseOutput)
}

func (o FrontDoorPolicySettingsResponsePtrOutput) CustomBlockResponseBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomBlockResponseBody
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponsePtrOutput) CustomBlockResponseStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.CustomBlockResponseStatusCode
	}).(pulumi.IntPtrOutput)
}

func (o FrontDoorPolicySettingsResponsePtrOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.EnabledState
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponsePtrOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RedirectUrl
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponsePtrOutput) RequestBodyCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RequestBodyCheck
	}).(pulumi.StringPtrOutput)
}

type FrontendEndpoint struct {
	HostName                         *string                                                           `pulumi:"hostName"`
	Id                               *string                                                           `pulumi:"id"`
	Name                             *string                                                           `pulumi:"name"`
	SessionAffinityEnabledState      *string                                                           `pulumi:"sessionAffinityEnabledState"`
	SessionAffinityTtlSeconds        *int                                                              `pulumi:"sessionAffinityTtlSeconds"`
	WebApplicationFirewallPolicyLink *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink `pulumi:"webApplicationFirewallPolicyLink"`
}





type FrontendEndpointInput interface {
	pulumi.Input

	ToFrontendEndpointOutput() FrontendEndpointOutput
	ToFrontendEndpointOutputWithContext(context.Context) FrontendEndpointOutput
}

type FrontendEndpointArgs struct {
	HostName                         pulumi.StringPtrInput                                                    `pulumi:"hostName"`
	Id                               pulumi.StringPtrInput                                                    `pulumi:"id"`
	Name                             pulumi.StringPtrInput                                                    `pulumi:"name"`
	SessionAffinityEnabledState      pulumi.StringPtrInput                                                    `pulumi:"sessionAffinityEnabledState"`
	SessionAffinityTtlSeconds        pulumi.IntPtrInput                                                       `pulumi:"sessionAffinityTtlSeconds"`
	WebApplicationFirewallPolicyLink FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrInput `pulumi:"webApplicationFirewallPolicyLink"`
}

func (FrontendEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpoint)(nil)).Elem()
}

func (i FrontendEndpointArgs) ToFrontendEndpointOutput() FrontendEndpointOutput {
	return i.ToFrontendEndpointOutputWithContext(context.Background())
}

func (i FrontendEndpointArgs) ToFrontendEndpointOutputWithContext(ctx context.Context) FrontendEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendEndpointOutput)
}





type FrontendEndpointArrayInput interface {
	pulumi.Input

	ToFrontendEndpointArrayOutput() FrontendEndpointArrayOutput
	ToFrontendEndpointArrayOutputWithContext(context.Context) FrontendEndpointArrayOutput
}

type FrontendEndpointArray []FrontendEndpointInput

func (FrontendEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendEndpoint)(nil)).Elem()
}

func (i FrontendEndpointArray) ToFrontendEndpointArrayOutput() FrontendEndpointArrayOutput {
	return i.ToFrontendEndpointArrayOutputWithContext(context.Background())
}

func (i FrontendEndpointArray) ToFrontendEndpointArrayOutputWithContext(ctx context.Context) FrontendEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendEndpointArrayOutput)
}

type FrontendEndpointOutput struct{ *pulumi.OutputState }

func (FrontendEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpoint)(nil)).Elem()
}

func (o FrontendEndpointOutput) ToFrontendEndpointOutput() FrontendEndpointOutput {
	return o
}

func (o FrontendEndpointOutput) ToFrontendEndpointOutputWithContext(ctx context.Context) FrontendEndpointOutput {
	return o
}

func (o FrontendEndpointOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpoint) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpoint) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointOutput) SessionAffinityEnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpoint) *string { return v.SessionAffinityEnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointOutput) SessionAffinityTtlSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FrontendEndpoint) *int { return v.SessionAffinityTtlSeconds }).(pulumi.IntPtrOutput)
}

func (o FrontendEndpointOutput) WebApplicationFirewallPolicyLink() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyT(func(v FrontendEndpoint) *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink {
		return v.WebApplicationFirewallPolicyLink
	}).(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

type FrontendEndpointArrayOutput struct{ *pulumi.OutputState }

func (FrontendEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendEndpoint)(nil)).Elem()
}

func (o FrontendEndpointArrayOutput) ToFrontendEndpointArrayOutput() FrontendEndpointArrayOutput {
	return o
}

func (o FrontendEndpointArrayOutput) ToFrontendEndpointArrayOutputWithContext(ctx context.Context) FrontendEndpointArrayOutput {
	return o
}

func (o FrontendEndpointArrayOutput) Index(i pulumi.IntInput) FrontendEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendEndpoint {
		return vs[0].([]FrontendEndpoint)[vs[1].(int)]
	}).(FrontendEndpointOutput)
}

type FrontendEndpointLinkResponse struct {
	Id *string `pulumi:"id"`
}

type FrontendEndpointLinkResponseOutput struct{ *pulumi.OutputState }

func (FrontendEndpointLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpointLinkResponse)(nil)).Elem()
}

func (o FrontendEndpointLinkResponseOutput) ToFrontendEndpointLinkResponseOutput() FrontendEndpointLinkResponseOutput {
	return o
}

func (o FrontendEndpointLinkResponseOutput) ToFrontendEndpointLinkResponseOutputWithContext(ctx context.Context) FrontendEndpointLinkResponseOutput {
	return o
}

func (o FrontendEndpointLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type FrontendEndpointLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontendEndpointLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendEndpointLinkResponse)(nil)).Elem()
}

func (o FrontendEndpointLinkResponseArrayOutput) ToFrontendEndpointLinkResponseArrayOutput() FrontendEndpointLinkResponseArrayOutput {
	return o
}

func (o FrontendEndpointLinkResponseArrayOutput) ToFrontendEndpointLinkResponseArrayOutputWithContext(ctx context.Context) FrontendEndpointLinkResponseArrayOutput {
	return o
}

func (o FrontendEndpointLinkResponseArrayOutput) Index(i pulumi.IntInput) FrontendEndpointLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendEndpointLinkResponse {
		return vs[0].([]FrontendEndpointLinkResponse)[vs[1].(int)]
	}).(FrontendEndpointLinkResponseOutput)
}

type FrontendEndpointResponse struct {
	CustomHttpsConfiguration         CustomHttpsConfigurationResponse                                          `pulumi:"customHttpsConfiguration"`
	CustomHttpsProvisioningState     string                                                                    `pulumi:"customHttpsProvisioningState"`
	CustomHttpsProvisioningSubstate  string                                                                    `pulumi:"customHttpsProvisioningSubstate"`
	HostName                         *string                                                                   `pulumi:"hostName"`
	Id                               *string                                                                   `pulumi:"id"`
	Name                             *string                                                                   `pulumi:"name"`
	ResourceState                    string                                                                    `pulumi:"resourceState"`
	SessionAffinityEnabledState      *string                                                                   `pulumi:"sessionAffinityEnabledState"`
	SessionAffinityTtlSeconds        *int                                                                      `pulumi:"sessionAffinityTtlSeconds"`
	Type                             string                                                                    `pulumi:"type"`
	WebApplicationFirewallPolicyLink *FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink `pulumi:"webApplicationFirewallPolicyLink"`
}

type FrontendEndpointResponseOutput struct{ *pulumi.OutputState }

func (FrontendEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpointResponse)(nil)).Elem()
}

func (o FrontendEndpointResponseOutput) ToFrontendEndpointResponseOutput() FrontendEndpointResponseOutput {
	return o
}

func (o FrontendEndpointResponseOutput) ToFrontendEndpointResponseOutputWithContext(ctx context.Context) FrontendEndpointResponseOutput {
	return o
}

func (o FrontendEndpointResponseOutput) CustomHttpsConfiguration() CustomHttpsConfigurationResponseOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) CustomHttpsConfigurationResponse { return v.CustomHttpsConfiguration }).(CustomHttpsConfigurationResponseOutput)
}

func (o FrontendEndpointResponseOutput) CustomHttpsProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) string { return v.CustomHttpsProvisioningState }).(pulumi.StringOutput)
}

func (o FrontendEndpointResponseOutput) CustomHttpsProvisioningSubstate() pulumi.StringOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) string { return v.CustomHttpsProvisioningSubstate }).(pulumi.StringOutput)
}

func (o FrontendEndpointResponseOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o FrontendEndpointResponseOutput) SessionAffinityEnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) *string { return v.SessionAffinityEnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointResponseOutput) SessionAffinityTtlSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) *int { return v.SessionAffinityTtlSeconds }).(pulumi.IntPtrOutput)
}

func (o FrontendEndpointResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o FrontendEndpointResponseOutput) WebApplicationFirewallPolicyLink() FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) *FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink {
		return v.WebApplicationFirewallPolicyLink
	}).(FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput)
}

type FrontendEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontendEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendEndpointResponse)(nil)).Elem()
}

func (o FrontendEndpointResponseArrayOutput) ToFrontendEndpointResponseArrayOutput() FrontendEndpointResponseArrayOutput {
	return o
}

func (o FrontendEndpointResponseArrayOutput) ToFrontendEndpointResponseArrayOutputWithContext(ctx context.Context) FrontendEndpointResponseArrayOutput {
	return o
}

func (o FrontendEndpointResponseArrayOutput) Index(i pulumi.IntInput) FrontendEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendEndpointResponse {
		return vs[0].([]FrontendEndpointResponse)[vs[1].(int)]
	}).(FrontendEndpointResponseOutput)
}

type FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink struct {
	Id *string `pulumi:"id"`
}

type FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput struct{ *pulumi.OutputState }

func (FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ToFrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput() FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ToFrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput struct{ *pulumi.OutputState }

func (FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ToFrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput() FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ToFrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) Elem() FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o.ApplyT(func(v *FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink) FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink {
		if v != nil {
			return *v
		}
		var ret FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink
		return ret
	}).(FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput)
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink struct {
	Id *string `pulumi:"id"`
}





type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkInput interface {
	pulumi.Input

	ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput
	ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput
}

type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (i FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return i.ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(context.Background())
}

func (i FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput)
}

func (i FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return i.ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (i FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput).ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx)
}









type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrInput interface {
	pulumi.Input

	ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput
	ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput
}

type frontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrType FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs

func FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtr(v *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrInput {
	return (*frontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrType)(v)
}

func (*frontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (i *frontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return i.ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (i *frontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput struct{ *pulumi.OutputState }

func (FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink) *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink {
		return &v
	}).(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput struct{ *pulumi.OutputState }

func (FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) Elem() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o.ApplyT(func(v *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink {
		if v != nil {
			return *v
		}
		var ret FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink
		return ret
	}).(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput)
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type HeaderAction struct {
	HeaderActionType string  `pulumi:"headerActionType"`
	HeaderName       string  `pulumi:"headerName"`
	Value            *string `pulumi:"value"`
}





type HeaderActionInput interface {
	pulumi.Input

	ToHeaderActionOutput() HeaderActionOutput
	ToHeaderActionOutputWithContext(context.Context) HeaderActionOutput
}

type HeaderActionArgs struct {
	HeaderActionType pulumi.StringInput    `pulumi:"headerActionType"`
	HeaderName       pulumi.StringInput    `pulumi:"headerName"`
	Value            pulumi.StringPtrInput `pulumi:"value"`
}

func (HeaderActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderAction)(nil)).Elem()
}

func (i HeaderActionArgs) ToHeaderActionOutput() HeaderActionOutput {
	return i.ToHeaderActionOutputWithContext(context.Background())
}

func (i HeaderActionArgs) ToHeaderActionOutputWithContext(ctx context.Context) HeaderActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeaderActionOutput)
}





type HeaderActionArrayInput interface {
	pulumi.Input

	ToHeaderActionArrayOutput() HeaderActionArrayOutput
	ToHeaderActionArrayOutputWithContext(context.Context) HeaderActionArrayOutput
}

type HeaderActionArray []HeaderActionInput

func (HeaderActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HeaderAction)(nil)).Elem()
}

func (i HeaderActionArray) ToHeaderActionArrayOutput() HeaderActionArrayOutput {
	return i.ToHeaderActionArrayOutputWithContext(context.Background())
}

func (i HeaderActionArray) ToHeaderActionArrayOutputWithContext(ctx context.Context) HeaderActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeaderActionArrayOutput)
}

type HeaderActionOutput struct{ *pulumi.OutputState }

func (HeaderActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderAction)(nil)).Elem()
}

func (o HeaderActionOutput) ToHeaderActionOutput() HeaderActionOutput {
	return o
}

func (o HeaderActionOutput) ToHeaderActionOutputWithContext(ctx context.Context) HeaderActionOutput {
	return o
}

func (o HeaderActionOutput) HeaderActionType() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderAction) string { return v.HeaderActionType }).(pulumi.StringOutput)
}

func (o HeaderActionOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderAction) string { return v.HeaderName }).(pulumi.StringOutput)
}

func (o HeaderActionOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HeaderAction) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type HeaderActionArrayOutput struct{ *pulumi.OutputState }

func (HeaderActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HeaderAction)(nil)).Elem()
}

func (o HeaderActionArrayOutput) ToHeaderActionArrayOutput() HeaderActionArrayOutput {
	return o
}

func (o HeaderActionArrayOutput) ToHeaderActionArrayOutputWithContext(ctx context.Context) HeaderActionArrayOutput {
	return o
}

func (o HeaderActionArrayOutput) Index(i pulumi.IntInput) HeaderActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HeaderAction {
		return vs[0].([]HeaderAction)[vs[1].(int)]
	}).(HeaderActionOutput)
}

type HeaderActionResponse struct {
	HeaderActionType string  `pulumi:"headerActionType"`
	HeaderName       string  `pulumi:"headerName"`
	Value            *string `pulumi:"value"`
}

type HeaderActionResponseOutput struct{ *pulumi.OutputState }

func (HeaderActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderActionResponse)(nil)).Elem()
}

func (o HeaderActionResponseOutput) ToHeaderActionResponseOutput() HeaderActionResponseOutput {
	return o
}

func (o HeaderActionResponseOutput) ToHeaderActionResponseOutputWithContext(ctx context.Context) HeaderActionResponseOutput {
	return o
}

func (o HeaderActionResponseOutput) HeaderActionType() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderActionResponse) string { return v.HeaderActionType }).(pulumi.StringOutput)
}

func (o HeaderActionResponseOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderActionResponse) string { return v.HeaderName }).(pulumi.StringOutput)
}

func (o HeaderActionResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HeaderActionResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type HeaderActionResponseArrayOutput struct{ *pulumi.OutputState }

func (HeaderActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HeaderActionResponse)(nil)).Elem()
}

func (o HeaderActionResponseArrayOutput) ToHeaderActionResponseArrayOutput() HeaderActionResponseArrayOutput {
	return o
}

func (o HeaderActionResponseArrayOutput) ToHeaderActionResponseArrayOutputWithContext(ctx context.Context) HeaderActionResponseArrayOutput {
	return o
}

func (o HeaderActionResponseArrayOutput) Index(i pulumi.IntInput) HeaderActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HeaderActionResponse {
		return vs[0].([]HeaderActionResponse)[vs[1].(int)]
	}).(HeaderActionResponseOutput)
}

type HealthProbeSettingsModel struct {
	EnabledState      *string `pulumi:"enabledState"`
	HealthProbeMethod *string `pulumi:"healthProbeMethod"`
	Id                *string `pulumi:"id"`
	IntervalInSeconds *int    `pulumi:"intervalInSeconds"`
	Name              *string `pulumi:"name"`
	Path              *string `pulumi:"path"`
	Protocol          *string `pulumi:"protocol"`
}


func (val *HealthProbeSettingsModel) Defaults() *HealthProbeSettingsModel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HealthProbeMethod) {
		healthProbeMethod_ := "HEAD"
		tmp.HealthProbeMethod = &healthProbeMethod_
	}
	return &tmp
}





type HealthProbeSettingsModelInput interface {
	pulumi.Input

	ToHealthProbeSettingsModelOutput() HealthProbeSettingsModelOutput
	ToHealthProbeSettingsModelOutputWithContext(context.Context) HealthProbeSettingsModelOutput
}

type HealthProbeSettingsModelArgs struct {
	EnabledState      pulumi.StringPtrInput `pulumi:"enabledState"`
	HealthProbeMethod pulumi.StringPtrInput `pulumi:"healthProbeMethod"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	IntervalInSeconds pulumi.IntPtrInput    `pulumi:"intervalInSeconds"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	Path              pulumi.StringPtrInput `pulumi:"path"`
	Protocol          pulumi.StringPtrInput `pulumi:"protocol"`
}


func (val *HealthProbeSettingsModelArgs) Defaults() *HealthProbeSettingsModelArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HealthProbeMethod) {
		tmp.HealthProbeMethod = pulumi.StringPtr("HEAD")
	}
	return &tmp
}
func (HealthProbeSettingsModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthProbeSettingsModel)(nil)).Elem()
}

func (i HealthProbeSettingsModelArgs) ToHealthProbeSettingsModelOutput() HealthProbeSettingsModelOutput {
	return i.ToHealthProbeSettingsModelOutputWithContext(context.Background())
}

func (i HealthProbeSettingsModelArgs) ToHealthProbeSettingsModelOutputWithContext(ctx context.Context) HealthProbeSettingsModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthProbeSettingsModelOutput)
}





type HealthProbeSettingsModelArrayInput interface {
	pulumi.Input

	ToHealthProbeSettingsModelArrayOutput() HealthProbeSettingsModelArrayOutput
	ToHealthProbeSettingsModelArrayOutputWithContext(context.Context) HealthProbeSettingsModelArrayOutput
}

type HealthProbeSettingsModelArray []HealthProbeSettingsModelInput

func (HealthProbeSettingsModelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthProbeSettingsModel)(nil)).Elem()
}

func (i HealthProbeSettingsModelArray) ToHealthProbeSettingsModelArrayOutput() HealthProbeSettingsModelArrayOutput {
	return i.ToHealthProbeSettingsModelArrayOutputWithContext(context.Background())
}

func (i HealthProbeSettingsModelArray) ToHealthProbeSettingsModelArrayOutputWithContext(ctx context.Context) HealthProbeSettingsModelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthProbeSettingsModelArrayOutput)
}

type HealthProbeSettingsModelOutput struct{ *pulumi.OutputState }

func (HealthProbeSettingsModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthProbeSettingsModel)(nil)).Elem()
}

func (o HealthProbeSettingsModelOutput) ToHealthProbeSettingsModelOutput() HealthProbeSettingsModelOutput {
	return o
}

func (o HealthProbeSettingsModelOutput) ToHealthProbeSettingsModelOutputWithContext(ctx context.Context) HealthProbeSettingsModelOutput {
	return o
}

func (o HealthProbeSettingsModelOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelOutput) HealthProbeMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *string { return v.HealthProbeMethod }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *int { return v.IntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o HealthProbeSettingsModelOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type HealthProbeSettingsModelArrayOutput struct{ *pulumi.OutputState }

func (HealthProbeSettingsModelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthProbeSettingsModel)(nil)).Elem()
}

func (o HealthProbeSettingsModelArrayOutput) ToHealthProbeSettingsModelArrayOutput() HealthProbeSettingsModelArrayOutput {
	return o
}

func (o HealthProbeSettingsModelArrayOutput) ToHealthProbeSettingsModelArrayOutputWithContext(ctx context.Context) HealthProbeSettingsModelArrayOutput {
	return o
}

func (o HealthProbeSettingsModelArrayOutput) Index(i pulumi.IntInput) HealthProbeSettingsModelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthProbeSettingsModel {
		return vs[0].([]HealthProbeSettingsModel)[vs[1].(int)]
	}).(HealthProbeSettingsModelOutput)
}

type HealthProbeSettingsModelResponse struct {
	EnabledState      *string `pulumi:"enabledState"`
	HealthProbeMethod *string `pulumi:"healthProbeMethod"`
	Id                *string `pulumi:"id"`
	IntervalInSeconds *int    `pulumi:"intervalInSeconds"`
	Name              *string `pulumi:"name"`
	Path              *string `pulumi:"path"`
	Protocol          *string `pulumi:"protocol"`
	ResourceState     string  `pulumi:"resourceState"`
	Type              string  `pulumi:"type"`
}


func (val *HealthProbeSettingsModelResponse) Defaults() *HealthProbeSettingsModelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HealthProbeMethod) {
		healthProbeMethod_ := "HEAD"
		tmp.HealthProbeMethod = &healthProbeMethod_
	}
	return &tmp
}

type HealthProbeSettingsModelResponseOutput struct{ *pulumi.OutputState }

func (HealthProbeSettingsModelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthProbeSettingsModelResponse)(nil)).Elem()
}

func (o HealthProbeSettingsModelResponseOutput) ToHealthProbeSettingsModelResponseOutput() HealthProbeSettingsModelResponseOutput {
	return o
}

func (o HealthProbeSettingsModelResponseOutput) ToHealthProbeSettingsModelResponseOutputWithContext(ctx context.Context) HealthProbeSettingsModelResponseOutput {
	return o
}

func (o HealthProbeSettingsModelResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) HealthProbeMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *string { return v.HealthProbeMethod }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *int { return v.IntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o HealthProbeSettingsModelResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) string { return v.Type }).(pulumi.StringOutput)
}

type HealthProbeSettingsModelResponseArrayOutput struct{ *pulumi.OutputState }

func (HealthProbeSettingsModelResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthProbeSettingsModelResponse)(nil)).Elem()
}

func (o HealthProbeSettingsModelResponseArrayOutput) ToHealthProbeSettingsModelResponseArrayOutput() HealthProbeSettingsModelResponseArrayOutput {
	return o
}

func (o HealthProbeSettingsModelResponseArrayOutput) ToHealthProbeSettingsModelResponseArrayOutputWithContext(ctx context.Context) HealthProbeSettingsModelResponseArrayOutput {
	return o
}

func (o HealthProbeSettingsModelResponseArrayOutput) Index(i pulumi.IntInput) HealthProbeSettingsModelResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthProbeSettingsModelResponse {
		return vs[0].([]HealthProbeSettingsModelResponse)[vs[1].(int)]
	}).(HealthProbeSettingsModelResponseOutput)
}

type KeyVaultCertificateSourceParametersResponseVault struct {
	Id *string `pulumi:"id"`
}

type KeyVaultCertificateSourceParametersResponseVaultOutput struct{ *pulumi.OutputState }

func (KeyVaultCertificateSourceParametersResponseVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultCertificateSourceParametersResponseVault)(nil)).Elem()
}

func (o KeyVaultCertificateSourceParametersResponseVaultOutput) ToKeyVaultCertificateSourceParametersResponseVaultOutput() KeyVaultCertificateSourceParametersResponseVaultOutput {
	return o
}

func (o KeyVaultCertificateSourceParametersResponseVaultOutput) ToKeyVaultCertificateSourceParametersResponseVaultOutputWithContext(ctx context.Context) KeyVaultCertificateSourceParametersResponseVaultOutput {
	return o
}

func (o KeyVaultCertificateSourceParametersResponseVaultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultCertificateSourceParametersResponseVault) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type KeyVaultCertificateSourceParametersResponseVaultPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultCertificateSourceParametersResponseVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultCertificateSourceParametersResponseVault)(nil)).Elem()
}

func (o KeyVaultCertificateSourceParametersResponseVaultPtrOutput) ToKeyVaultCertificateSourceParametersResponseVaultPtrOutput() KeyVaultCertificateSourceParametersResponseVaultPtrOutput {
	return o
}

func (o KeyVaultCertificateSourceParametersResponseVaultPtrOutput) ToKeyVaultCertificateSourceParametersResponseVaultPtrOutputWithContext(ctx context.Context) KeyVaultCertificateSourceParametersResponseVaultPtrOutput {
	return o
}

func (o KeyVaultCertificateSourceParametersResponseVaultPtrOutput) Elem() KeyVaultCertificateSourceParametersResponseVaultOutput {
	return o.ApplyT(func(v *KeyVaultCertificateSourceParametersResponseVault) KeyVaultCertificateSourceParametersResponseVault {
		if v != nil {
			return *v
		}
		var ret KeyVaultCertificateSourceParametersResponseVault
		return ret
	}).(KeyVaultCertificateSourceParametersResponseVaultOutput)
}

func (o KeyVaultCertificateSourceParametersResponseVaultPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultCertificateSourceParametersResponseVault) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type LoadBalancingSettingsModel struct {
	AdditionalLatencyMilliseconds *int    `pulumi:"additionalLatencyMilliseconds"`
	Id                            *string `pulumi:"id"`
	Name                          *string `pulumi:"name"`
	SampleSize                    *int    `pulumi:"sampleSize"`
	SuccessfulSamplesRequired     *int    `pulumi:"successfulSamplesRequired"`
}





type LoadBalancingSettingsModelInput interface {
	pulumi.Input

	ToLoadBalancingSettingsModelOutput() LoadBalancingSettingsModelOutput
	ToLoadBalancingSettingsModelOutputWithContext(context.Context) LoadBalancingSettingsModelOutput
}

type LoadBalancingSettingsModelArgs struct {
	AdditionalLatencyMilliseconds pulumi.IntPtrInput    `pulumi:"additionalLatencyMilliseconds"`
	Id                            pulumi.StringPtrInput `pulumi:"id"`
	Name                          pulumi.StringPtrInput `pulumi:"name"`
	SampleSize                    pulumi.IntPtrInput    `pulumi:"sampleSize"`
	SuccessfulSamplesRequired     pulumi.IntPtrInput    `pulumi:"successfulSamplesRequired"`
}

func (LoadBalancingSettingsModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingSettingsModel)(nil)).Elem()
}

func (i LoadBalancingSettingsModelArgs) ToLoadBalancingSettingsModelOutput() LoadBalancingSettingsModelOutput {
	return i.ToLoadBalancingSettingsModelOutputWithContext(context.Background())
}

func (i LoadBalancingSettingsModelArgs) ToLoadBalancingSettingsModelOutputWithContext(ctx context.Context) LoadBalancingSettingsModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingSettingsModelOutput)
}





type LoadBalancingSettingsModelArrayInput interface {
	pulumi.Input

	ToLoadBalancingSettingsModelArrayOutput() LoadBalancingSettingsModelArrayOutput
	ToLoadBalancingSettingsModelArrayOutputWithContext(context.Context) LoadBalancingSettingsModelArrayOutput
}

type LoadBalancingSettingsModelArray []LoadBalancingSettingsModelInput

func (LoadBalancingSettingsModelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingSettingsModel)(nil)).Elem()
}

func (i LoadBalancingSettingsModelArray) ToLoadBalancingSettingsModelArrayOutput() LoadBalancingSettingsModelArrayOutput {
	return i.ToLoadBalancingSettingsModelArrayOutputWithContext(context.Background())
}

func (i LoadBalancingSettingsModelArray) ToLoadBalancingSettingsModelArrayOutputWithContext(ctx context.Context) LoadBalancingSettingsModelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingSettingsModelArrayOutput)
}

type LoadBalancingSettingsModelOutput struct{ *pulumi.OutputState }

func (LoadBalancingSettingsModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingSettingsModel)(nil)).Elem()
}

func (o LoadBalancingSettingsModelOutput) ToLoadBalancingSettingsModelOutput() LoadBalancingSettingsModelOutput {
	return o
}

func (o LoadBalancingSettingsModelOutput) ToLoadBalancingSettingsModelOutputWithContext(ctx context.Context) LoadBalancingSettingsModelOutput {
	return o
}

func (o LoadBalancingSettingsModelOutput) AdditionalLatencyMilliseconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModel) *int { return v.AdditionalLatencyMilliseconds }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsModelOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModel) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingSettingsModelOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModel) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingSettingsModelOutput) SampleSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModel) *int { return v.SampleSize }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsModelOutput) SuccessfulSamplesRequired() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModel) *int { return v.SuccessfulSamplesRequired }).(pulumi.IntPtrOutput)
}

type LoadBalancingSettingsModelArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingSettingsModelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingSettingsModel)(nil)).Elem()
}

func (o LoadBalancingSettingsModelArrayOutput) ToLoadBalancingSettingsModelArrayOutput() LoadBalancingSettingsModelArrayOutput {
	return o
}

func (o LoadBalancingSettingsModelArrayOutput) ToLoadBalancingSettingsModelArrayOutputWithContext(ctx context.Context) LoadBalancingSettingsModelArrayOutput {
	return o
}

func (o LoadBalancingSettingsModelArrayOutput) Index(i pulumi.IntInput) LoadBalancingSettingsModelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingSettingsModel {
		return vs[0].([]LoadBalancingSettingsModel)[vs[1].(int)]
	}).(LoadBalancingSettingsModelOutput)
}

type LoadBalancingSettingsModelResponse struct {
	AdditionalLatencyMilliseconds *int    `pulumi:"additionalLatencyMilliseconds"`
	Id                            *string `pulumi:"id"`
	Name                          *string `pulumi:"name"`
	ResourceState                 string  `pulumi:"resourceState"`
	SampleSize                    *int    `pulumi:"sampleSize"`
	SuccessfulSamplesRequired     *int    `pulumi:"successfulSamplesRequired"`
	Type                          string  `pulumi:"type"`
}

type LoadBalancingSettingsModelResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancingSettingsModelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingSettingsModelResponse)(nil)).Elem()
}

func (o LoadBalancingSettingsModelResponseOutput) ToLoadBalancingSettingsModelResponseOutput() LoadBalancingSettingsModelResponseOutput {
	return o
}

func (o LoadBalancingSettingsModelResponseOutput) ToLoadBalancingSettingsModelResponseOutputWithContext(ctx context.Context) LoadBalancingSettingsModelResponseOutput {
	return o
}

func (o LoadBalancingSettingsModelResponseOutput) AdditionalLatencyMilliseconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) *int { return v.AdditionalLatencyMilliseconds }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsModelResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingSettingsModelResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingSettingsModelResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LoadBalancingSettingsModelResponseOutput) SampleSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) *int { return v.SampleSize }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsModelResponseOutput) SuccessfulSamplesRequired() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) *int { return v.SuccessfulSamplesRequired }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsModelResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) string { return v.Type }).(pulumi.StringOutput)
}

type LoadBalancingSettingsModelResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingSettingsModelResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingSettingsModelResponse)(nil)).Elem()
}

func (o LoadBalancingSettingsModelResponseArrayOutput) ToLoadBalancingSettingsModelResponseArrayOutput() LoadBalancingSettingsModelResponseArrayOutput {
	return o
}

func (o LoadBalancingSettingsModelResponseArrayOutput) ToLoadBalancingSettingsModelResponseArrayOutputWithContext(ctx context.Context) LoadBalancingSettingsModelResponseArrayOutput {
	return o
}

func (o LoadBalancingSettingsModelResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancingSettingsModelResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingSettingsModelResponse {
		return vs[0].([]LoadBalancingSettingsModelResponse)[vs[1].(int)]
	}).(LoadBalancingSettingsModelResponseOutput)
}

type ManagedRuleExclusion struct {
	MatchVariable         string `pulumi:"matchVariable"`
	Selector              string `pulumi:"selector"`
	SelectorMatchOperator string `pulumi:"selectorMatchOperator"`
}





type ManagedRuleExclusionInput interface {
	pulumi.Input

	ToManagedRuleExclusionOutput() ManagedRuleExclusionOutput
	ToManagedRuleExclusionOutputWithContext(context.Context) ManagedRuleExclusionOutput
}

type ManagedRuleExclusionArgs struct {
	MatchVariable         pulumi.StringInput `pulumi:"matchVariable"`
	Selector              pulumi.StringInput `pulumi:"selector"`
	SelectorMatchOperator pulumi.StringInput `pulumi:"selectorMatchOperator"`
}

func (ManagedRuleExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleExclusion)(nil)).Elem()
}

func (i ManagedRuleExclusionArgs) ToManagedRuleExclusionOutput() ManagedRuleExclusionOutput {
	return i.ToManagedRuleExclusionOutputWithContext(context.Background())
}

func (i ManagedRuleExclusionArgs) ToManagedRuleExclusionOutputWithContext(ctx context.Context) ManagedRuleExclusionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleExclusionOutput)
}





type ManagedRuleExclusionArrayInput interface {
	pulumi.Input

	ToManagedRuleExclusionArrayOutput() ManagedRuleExclusionArrayOutput
	ToManagedRuleExclusionArrayOutputWithContext(context.Context) ManagedRuleExclusionArrayOutput
}

type ManagedRuleExclusionArray []ManagedRuleExclusionInput

func (ManagedRuleExclusionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleExclusion)(nil)).Elem()
}

func (i ManagedRuleExclusionArray) ToManagedRuleExclusionArrayOutput() ManagedRuleExclusionArrayOutput {
	return i.ToManagedRuleExclusionArrayOutputWithContext(context.Background())
}

func (i ManagedRuleExclusionArray) ToManagedRuleExclusionArrayOutputWithContext(ctx context.Context) ManagedRuleExclusionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleExclusionArrayOutput)
}

type ManagedRuleExclusionOutput struct{ *pulumi.OutputState }

func (ManagedRuleExclusionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleExclusion)(nil)).Elem()
}

func (o ManagedRuleExclusionOutput) ToManagedRuleExclusionOutput() ManagedRuleExclusionOutput {
	return o
}

func (o ManagedRuleExclusionOutput) ToManagedRuleExclusionOutputWithContext(ctx context.Context) ManagedRuleExclusionOutput {
	return o
}

func (o ManagedRuleExclusionOutput) MatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleExclusion) string { return v.MatchVariable }).(pulumi.StringOutput)
}

func (o ManagedRuleExclusionOutput) Selector() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleExclusion) string { return v.Selector }).(pulumi.StringOutput)
}

func (o ManagedRuleExclusionOutput) SelectorMatchOperator() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleExclusion) string { return v.SelectorMatchOperator }).(pulumi.StringOutput)
}

type ManagedRuleExclusionArrayOutput struct{ *pulumi.OutputState }

func (ManagedRuleExclusionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleExclusion)(nil)).Elem()
}

func (o ManagedRuleExclusionArrayOutput) ToManagedRuleExclusionArrayOutput() ManagedRuleExclusionArrayOutput {
	return o
}

func (o ManagedRuleExclusionArrayOutput) ToManagedRuleExclusionArrayOutputWithContext(ctx context.Context) ManagedRuleExclusionArrayOutput {
	return o
}

func (o ManagedRuleExclusionArrayOutput) Index(i pulumi.IntInput) ManagedRuleExclusionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedRuleExclusion {
		return vs[0].([]ManagedRuleExclusion)[vs[1].(int)]
	}).(ManagedRuleExclusionOutput)
}

type ManagedRuleExclusionResponse struct {
	MatchVariable         string `pulumi:"matchVariable"`
	Selector              string `pulumi:"selector"`
	SelectorMatchOperator string `pulumi:"selectorMatchOperator"`
}

type ManagedRuleExclusionResponseOutput struct{ *pulumi.OutputState }

func (ManagedRuleExclusionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleExclusionResponse)(nil)).Elem()
}

func (o ManagedRuleExclusionResponseOutput) ToManagedRuleExclusionResponseOutput() ManagedRuleExclusionResponseOutput {
	return o
}

func (o ManagedRuleExclusionResponseOutput) ToManagedRuleExclusionResponseOutputWithContext(ctx context.Context) ManagedRuleExclusionResponseOutput {
	return o
}

func (o ManagedRuleExclusionResponseOutput) MatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleExclusionResponse) string { return v.MatchVariable }).(pulumi.StringOutput)
}

func (o ManagedRuleExclusionResponseOutput) Selector() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleExclusionResponse) string { return v.Selector }).(pulumi.StringOutput)
}

func (o ManagedRuleExclusionResponseOutput) SelectorMatchOperator() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleExclusionResponse) string { return v.SelectorMatchOperator }).(pulumi.StringOutput)
}

type ManagedRuleExclusionResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedRuleExclusionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleExclusionResponse)(nil)).Elem()
}

func (o ManagedRuleExclusionResponseArrayOutput) ToManagedRuleExclusionResponseArrayOutput() ManagedRuleExclusionResponseArrayOutput {
	return o
}

func (o ManagedRuleExclusionResponseArrayOutput) ToManagedRuleExclusionResponseArrayOutputWithContext(ctx context.Context) ManagedRuleExclusionResponseArrayOutput {
	return o
}

func (o ManagedRuleExclusionResponseArrayOutput) Index(i pulumi.IntInput) ManagedRuleExclusionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedRuleExclusionResponse {
		return vs[0].([]ManagedRuleExclusionResponse)[vs[1].(int)]
	}).(ManagedRuleExclusionResponseOutput)
}

type ManagedRuleSetList struct {
	ManagedRuleSets []FrontDoorManagedRuleSet `pulumi:"managedRuleSets"`
}





type ManagedRuleSetListInput interface {
	pulumi.Input

	ToManagedRuleSetListOutput() ManagedRuleSetListOutput
	ToManagedRuleSetListOutputWithContext(context.Context) ManagedRuleSetListOutput
}

type ManagedRuleSetListArgs struct {
	ManagedRuleSets FrontDoorManagedRuleSetArrayInput `pulumi:"managedRuleSets"`
}

func (ManagedRuleSetListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSetList)(nil)).Elem()
}

func (i ManagedRuleSetListArgs) ToManagedRuleSetListOutput() ManagedRuleSetListOutput {
	return i.ToManagedRuleSetListOutputWithContext(context.Background())
}

func (i ManagedRuleSetListArgs) ToManagedRuleSetListOutputWithContext(ctx context.Context) ManagedRuleSetListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetListOutput)
}

func (i ManagedRuleSetListArgs) ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput {
	return i.ToManagedRuleSetListPtrOutputWithContext(context.Background())
}

func (i ManagedRuleSetListArgs) ToManagedRuleSetListPtrOutputWithContext(ctx context.Context) ManagedRuleSetListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetListOutput).ToManagedRuleSetListPtrOutputWithContext(ctx)
}









type ManagedRuleSetListPtrInput interface {
	pulumi.Input

	ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput
	ToManagedRuleSetListPtrOutputWithContext(context.Context) ManagedRuleSetListPtrOutput
}

type managedRuleSetListPtrType ManagedRuleSetListArgs

func ManagedRuleSetListPtr(v *ManagedRuleSetListArgs) ManagedRuleSetListPtrInput {
	return (*managedRuleSetListPtrType)(v)
}

func (*managedRuleSetListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRuleSetList)(nil)).Elem()
}

func (i *managedRuleSetListPtrType) ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput {
	return i.ToManagedRuleSetListPtrOutputWithContext(context.Background())
}

func (i *managedRuleSetListPtrType) ToManagedRuleSetListPtrOutputWithContext(ctx context.Context) ManagedRuleSetListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetListPtrOutput)
}

type ManagedRuleSetListOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSetList)(nil)).Elem()
}

func (o ManagedRuleSetListOutput) ToManagedRuleSetListOutput() ManagedRuleSetListOutput {
	return o
}

func (o ManagedRuleSetListOutput) ToManagedRuleSetListOutputWithContext(ctx context.Context) ManagedRuleSetListOutput {
	return o
}

func (o ManagedRuleSetListOutput) ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput {
	return o.ToManagedRuleSetListPtrOutputWithContext(context.Background())
}

func (o ManagedRuleSetListOutput) ToManagedRuleSetListPtrOutputWithContext(ctx context.Context) ManagedRuleSetListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedRuleSetList) *ManagedRuleSetList {
		return &v
	}).(ManagedRuleSetListPtrOutput)
}

func (o ManagedRuleSetListOutput) ManagedRuleSets() FrontDoorManagedRuleSetArrayOutput {
	return o.ApplyT(func(v ManagedRuleSetList) []FrontDoorManagedRuleSet { return v.ManagedRuleSets }).(FrontDoorManagedRuleSetArrayOutput)
}

type ManagedRuleSetListPtrOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRuleSetList)(nil)).Elem()
}

func (o ManagedRuleSetListPtrOutput) ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput {
	return o
}

func (o ManagedRuleSetListPtrOutput) ToManagedRuleSetListPtrOutputWithContext(ctx context.Context) ManagedRuleSetListPtrOutput {
	return o
}

func (o ManagedRuleSetListPtrOutput) Elem() ManagedRuleSetListOutput {
	return o.ApplyT(func(v *ManagedRuleSetList) ManagedRuleSetList {
		if v != nil {
			return *v
		}
		var ret ManagedRuleSetList
		return ret
	}).(ManagedRuleSetListOutput)
}

func (o ManagedRuleSetListPtrOutput) ManagedRuleSets() FrontDoorManagedRuleSetArrayOutput {
	return o.ApplyT(func(v *ManagedRuleSetList) []FrontDoorManagedRuleSet {
		if v == nil {
			return nil
		}
		return v.ManagedRuleSets
	}).(FrontDoorManagedRuleSetArrayOutput)
}

type ManagedRuleSetListResponse struct {
	ManagedRuleSets []FrontDoorManagedRuleSetResponse `pulumi:"managedRuleSets"`
}

type ManagedRuleSetListResponseOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSetListResponse)(nil)).Elem()
}

func (o ManagedRuleSetListResponseOutput) ToManagedRuleSetListResponseOutput() ManagedRuleSetListResponseOutput {
	return o
}

func (o ManagedRuleSetListResponseOutput) ToManagedRuleSetListResponseOutputWithContext(ctx context.Context) ManagedRuleSetListResponseOutput {
	return o
}

func (o ManagedRuleSetListResponseOutput) ManagedRuleSets() FrontDoorManagedRuleSetResponseArrayOutput {
	return o.ApplyT(func(v ManagedRuleSetListResponse) []FrontDoorManagedRuleSetResponse { return v.ManagedRuleSets }).(FrontDoorManagedRuleSetResponseArrayOutput)
}

type ManagedRuleSetListResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetListResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRuleSetListResponse)(nil)).Elem()
}

func (o ManagedRuleSetListResponsePtrOutput) ToManagedRuleSetListResponsePtrOutput() ManagedRuleSetListResponsePtrOutput {
	return o
}

func (o ManagedRuleSetListResponsePtrOutput) ToManagedRuleSetListResponsePtrOutputWithContext(ctx context.Context) ManagedRuleSetListResponsePtrOutput {
	return o
}

func (o ManagedRuleSetListResponsePtrOutput) Elem() ManagedRuleSetListResponseOutput {
	return o.ApplyT(func(v *ManagedRuleSetListResponse) ManagedRuleSetListResponse {
		if v != nil {
			return *v
		}
		var ret ManagedRuleSetListResponse
		return ret
	}).(ManagedRuleSetListResponseOutput)
}

func (o ManagedRuleSetListResponsePtrOutput) ManagedRuleSets() FrontDoorManagedRuleSetResponseArrayOutput {
	return o.ApplyT(func(v *ManagedRuleSetListResponse) []FrontDoorManagedRuleSetResponse {
		if v == nil {
			return nil
		}
		return v.ManagedRuleSets
	}).(FrontDoorManagedRuleSetResponseArrayOutput)
}

type RedirectConfiguration struct {
	CustomFragment    *string `pulumi:"customFragment"`
	CustomHost        *string `pulumi:"customHost"`
	CustomPath        *string `pulumi:"customPath"`
	CustomQueryString *string `pulumi:"customQueryString"`
	OdataType         string  `pulumi:"odataType"`
	RedirectProtocol  *string `pulumi:"redirectProtocol"`
	RedirectType      *string `pulumi:"redirectType"`
}

type RedirectConfigurationResponse struct {
	CustomFragment    *string `pulumi:"customFragment"`
	CustomHost        *string `pulumi:"customHost"`
	CustomPath        *string `pulumi:"customPath"`
	CustomQueryString *string `pulumi:"customQueryString"`
	OdataType         string  `pulumi:"odataType"`
	RedirectProtocol  *string `pulumi:"redirectProtocol"`
	RedirectType      *string `pulumi:"redirectType"`
}

type RoutingRule struct {
	AcceptedProtocols                []string                                                     `pulumi:"acceptedProtocols"`
	EnabledState                     *string                                                      `pulumi:"enabledState"`
	FrontendEndpoints                []SubResource                                                `pulumi:"frontendEndpoints"`
	Id                               *string                                                      `pulumi:"id"`
	Name                             *string                                                      `pulumi:"name"`
	PatternsToMatch                  []string                                                     `pulumi:"patternsToMatch"`
	RouteConfiguration               interface{}                                                  `pulumi:"routeConfiguration"`
	RulesEngine                      *SubResource                                                 `pulumi:"rulesEngine"`
	WebApplicationFirewallPolicyLink *RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink `pulumi:"webApplicationFirewallPolicyLink"`
}





type RoutingRuleInput interface {
	pulumi.Input

	ToRoutingRuleOutput() RoutingRuleOutput
	ToRoutingRuleOutputWithContext(context.Context) RoutingRuleOutput
}

type RoutingRuleArgs struct {
	AcceptedProtocols                pulumi.StringArrayInput                                             `pulumi:"acceptedProtocols"`
	EnabledState                     pulumi.StringPtrInput                                               `pulumi:"enabledState"`
	FrontendEndpoints                SubResourceArrayInput                                               `pulumi:"frontendEndpoints"`
	Id                               pulumi.StringPtrInput                                               `pulumi:"id"`
	Name                             pulumi.StringPtrInput                                               `pulumi:"name"`
	PatternsToMatch                  pulumi.StringArrayInput                                             `pulumi:"patternsToMatch"`
	RouteConfiguration               pulumi.Input                                                        `pulumi:"routeConfiguration"`
	RulesEngine                      SubResourcePtrInput                                                 `pulumi:"rulesEngine"`
	WebApplicationFirewallPolicyLink RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrInput `pulumi:"webApplicationFirewallPolicyLink"`
}

func (RoutingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRule)(nil)).Elem()
}

func (i RoutingRuleArgs) ToRoutingRuleOutput() RoutingRuleOutput {
	return i.ToRoutingRuleOutputWithContext(context.Background())
}

func (i RoutingRuleArgs) ToRoutingRuleOutputWithContext(ctx context.Context) RoutingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingRuleOutput)
}





type RoutingRuleArrayInput interface {
	pulumi.Input

	ToRoutingRuleArrayOutput() RoutingRuleArrayOutput
	ToRoutingRuleArrayOutputWithContext(context.Context) RoutingRuleArrayOutput
}

type RoutingRuleArray []RoutingRuleInput

func (RoutingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingRule)(nil)).Elem()
}

func (i RoutingRuleArray) ToRoutingRuleArrayOutput() RoutingRuleArrayOutput {
	return i.ToRoutingRuleArrayOutputWithContext(context.Background())
}

func (i RoutingRuleArray) ToRoutingRuleArrayOutputWithContext(ctx context.Context) RoutingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingRuleArrayOutput)
}

type RoutingRuleOutput struct{ *pulumi.OutputState }

func (RoutingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRule)(nil)).Elem()
}

func (o RoutingRuleOutput) ToRoutingRuleOutput() RoutingRuleOutput {
	return o
}

func (o RoutingRuleOutput) ToRoutingRuleOutputWithContext(ctx context.Context) RoutingRuleOutput {
	return o
}

func (o RoutingRuleOutput) AcceptedProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoutingRule) []string { return v.AcceptedProtocols }).(pulumi.StringArrayOutput)
}

func (o RoutingRuleOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRule) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o RoutingRuleOutput) FrontendEndpoints() SubResourceArrayOutput {
	return o.ApplyT(func(v RoutingRule) []SubResource { return v.FrontendEndpoints }).(SubResourceArrayOutput)
}

func (o RoutingRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoutingRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RoutingRuleOutput) PatternsToMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoutingRule) []string { return v.PatternsToMatch }).(pulumi.StringArrayOutput)
}

func (o RoutingRuleOutput) RouteConfiguration() pulumi.AnyOutput {
	return o.ApplyT(func(v RoutingRule) interface{} { return v.RouteConfiguration }).(pulumi.AnyOutput)
}

func (o RoutingRuleOutput) RulesEngine() SubResourcePtrOutput {
	return o.ApplyT(func(v RoutingRule) *SubResource { return v.RulesEngine }).(SubResourcePtrOutput)
}

func (o RoutingRuleOutput) WebApplicationFirewallPolicyLink() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyT(func(v RoutingRule) *RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink {
		return v.WebApplicationFirewallPolicyLink
	}).(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

type RoutingRuleArrayOutput struct{ *pulumi.OutputState }

func (RoutingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingRule)(nil)).Elem()
}

func (o RoutingRuleArrayOutput) ToRoutingRuleArrayOutput() RoutingRuleArrayOutput {
	return o
}

func (o RoutingRuleArrayOutput) ToRoutingRuleArrayOutputWithContext(ctx context.Context) RoutingRuleArrayOutput {
	return o
}

func (o RoutingRuleArrayOutput) Index(i pulumi.IntInput) RoutingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingRule {
		return vs[0].([]RoutingRule)[vs[1].(int)]
	}).(RoutingRuleOutput)
}

type RoutingRuleLinkResponse struct {
	Id *string `pulumi:"id"`
}

type RoutingRuleLinkResponseOutput struct{ *pulumi.OutputState }

func (RoutingRuleLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRuleLinkResponse)(nil)).Elem()
}

func (o RoutingRuleLinkResponseOutput) ToRoutingRuleLinkResponseOutput() RoutingRuleLinkResponseOutput {
	return o
}

func (o RoutingRuleLinkResponseOutput) ToRoutingRuleLinkResponseOutputWithContext(ctx context.Context) RoutingRuleLinkResponseOutput {
	return o
}

func (o RoutingRuleLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRuleLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type RoutingRuleLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutingRuleLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingRuleLinkResponse)(nil)).Elem()
}

func (o RoutingRuleLinkResponseArrayOutput) ToRoutingRuleLinkResponseArrayOutput() RoutingRuleLinkResponseArrayOutput {
	return o
}

func (o RoutingRuleLinkResponseArrayOutput) ToRoutingRuleLinkResponseArrayOutputWithContext(ctx context.Context) RoutingRuleLinkResponseArrayOutput {
	return o
}

func (o RoutingRuleLinkResponseArrayOutput) Index(i pulumi.IntInput) RoutingRuleLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingRuleLinkResponse {
		return vs[0].([]RoutingRuleLinkResponse)[vs[1].(int)]
	}).(RoutingRuleLinkResponseOutput)
}

type RoutingRuleResponse struct {
	AcceptedProtocols                []string                                                             `pulumi:"acceptedProtocols"`
	EnabledState                     *string                                                              `pulumi:"enabledState"`
	FrontendEndpoints                []SubResourceResponse                                                `pulumi:"frontendEndpoints"`
	Id                               *string                                                              `pulumi:"id"`
	Name                             *string                                                              `pulumi:"name"`
	PatternsToMatch                  []string                                                             `pulumi:"patternsToMatch"`
	ResourceState                    string                                                               `pulumi:"resourceState"`
	RouteConfiguration               interface{}                                                          `pulumi:"routeConfiguration"`
	RulesEngine                      *SubResourceResponse                                                 `pulumi:"rulesEngine"`
	Type                             string                                                               `pulumi:"type"`
	WebApplicationFirewallPolicyLink *RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink `pulumi:"webApplicationFirewallPolicyLink"`
}

type RoutingRuleResponseOutput struct{ *pulumi.OutputState }

func (RoutingRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRuleResponse)(nil)).Elem()
}

func (o RoutingRuleResponseOutput) ToRoutingRuleResponseOutput() RoutingRuleResponseOutput {
	return o
}

func (o RoutingRuleResponseOutput) ToRoutingRuleResponseOutputWithContext(ctx context.Context) RoutingRuleResponseOutput {
	return o
}

func (o RoutingRuleResponseOutput) AcceptedProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoutingRuleResponse) []string { return v.AcceptedProtocols }).(pulumi.StringArrayOutput)
}

func (o RoutingRuleResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRuleResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o RoutingRuleResponseOutput) FrontendEndpoints() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v RoutingRuleResponse) []SubResourceResponse { return v.FrontendEndpoints }).(SubResourceResponseArrayOutput)
}

func (o RoutingRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoutingRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RoutingRuleResponseOutput) PatternsToMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoutingRuleResponse) []string { return v.PatternsToMatch }).(pulumi.StringArrayOutput)
}

func (o RoutingRuleResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingRuleResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o RoutingRuleResponseOutput) RouteConfiguration() pulumi.AnyOutput {
	return o.ApplyT(func(v RoutingRuleResponse) interface{} { return v.RouteConfiguration }).(pulumi.AnyOutput)
}

func (o RoutingRuleResponseOutput) RulesEngine() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v RoutingRuleResponse) *SubResourceResponse { return v.RulesEngine }).(SubResourceResponsePtrOutput)
}

func (o RoutingRuleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingRuleResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o RoutingRuleResponseOutput) WebApplicationFirewallPolicyLink() RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyT(func(v RoutingRuleResponse) *RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink {
		return v.WebApplicationFirewallPolicyLink
	}).(RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput)
}

type RoutingRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutingRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingRuleResponse)(nil)).Elem()
}

func (o RoutingRuleResponseArrayOutput) ToRoutingRuleResponseArrayOutput() RoutingRuleResponseArrayOutput {
	return o
}

func (o RoutingRuleResponseArrayOutput) ToRoutingRuleResponseArrayOutputWithContext(ctx context.Context) RoutingRuleResponseArrayOutput {
	return o
}

func (o RoutingRuleResponseArrayOutput) Index(i pulumi.IntInput) RoutingRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingRuleResponse {
		return vs[0].([]RoutingRuleResponse)[vs[1].(int)]
	}).(RoutingRuleResponseOutput)
}

type RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink struct {
	Id *string `pulumi:"id"`
}

type RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput struct{ *pulumi.OutputState }

func (RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ToRoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput() RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ToRoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput struct{ *pulumi.OutputState }

func (RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ToRoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput() RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ToRoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) Elem() RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o.ApplyT(func(v *RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink) RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink {
		if v != nil {
			return *v
		}
		var ret RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink
		return ret
	}).(RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput)
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink struct {
	Id *string `pulumi:"id"`
}





type RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkInput interface {
	pulumi.Input

	ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput
	ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput
}

type RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (i RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return i.ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(context.Background())
}

func (i RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput)
}

func (i RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return i.ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (i RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput).ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx)
}









type RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrInput interface {
	pulumi.Input

	ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput
	ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput
}

type routingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrType RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs

func RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtr(v *RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrInput {
	return (*routingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrType)(v)
}

func (*routingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (i *routingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return i.ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (i *routingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

type RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput struct{ *pulumi.OutputState }

func (RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink) *RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink {
		return &v
	}).(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput struct{ *pulumi.OutputState }

func (RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) Elem() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o.ApplyT(func(v *RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink {
		if v != nil {
			return *v
		}
		var ret RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink
		return ret
	}).(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput)
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type RulesEngineAction struct {
	RequestHeaderActions       []HeaderAction `pulumi:"requestHeaderActions"`
	ResponseHeaderActions      []HeaderAction `pulumi:"responseHeaderActions"`
	RouteConfigurationOverride interface{}    `pulumi:"routeConfigurationOverride"`
}





type RulesEngineActionInput interface {
	pulumi.Input

	ToRulesEngineActionOutput() RulesEngineActionOutput
	ToRulesEngineActionOutputWithContext(context.Context) RulesEngineActionOutput
}

type RulesEngineActionArgs struct {
	RequestHeaderActions       HeaderActionArrayInput `pulumi:"requestHeaderActions"`
	ResponseHeaderActions      HeaderActionArrayInput `pulumi:"responseHeaderActions"`
	RouteConfigurationOverride pulumi.Input           `pulumi:"routeConfigurationOverride"`
}

func (RulesEngineActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineAction)(nil)).Elem()
}

func (i RulesEngineActionArgs) ToRulesEngineActionOutput() RulesEngineActionOutput {
	return i.ToRulesEngineActionOutputWithContext(context.Background())
}

func (i RulesEngineActionArgs) ToRulesEngineActionOutputWithContext(ctx context.Context) RulesEngineActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineActionOutput)
}

type RulesEngineActionOutput struct{ *pulumi.OutputState }

func (RulesEngineActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineAction)(nil)).Elem()
}

func (o RulesEngineActionOutput) ToRulesEngineActionOutput() RulesEngineActionOutput {
	return o
}

func (o RulesEngineActionOutput) ToRulesEngineActionOutputWithContext(ctx context.Context) RulesEngineActionOutput {
	return o
}

func (o RulesEngineActionOutput) RequestHeaderActions() HeaderActionArrayOutput {
	return o.ApplyT(func(v RulesEngineAction) []HeaderAction { return v.RequestHeaderActions }).(HeaderActionArrayOutput)
}

func (o RulesEngineActionOutput) ResponseHeaderActions() HeaderActionArrayOutput {
	return o.ApplyT(func(v RulesEngineAction) []HeaderAction { return v.ResponseHeaderActions }).(HeaderActionArrayOutput)
}

func (o RulesEngineActionOutput) RouteConfigurationOverride() pulumi.AnyOutput {
	return o.ApplyT(func(v RulesEngineAction) interface{} { return v.RouteConfigurationOverride }).(pulumi.AnyOutput)
}

type RulesEngineActionResponse struct {
	RequestHeaderActions       []HeaderActionResponse `pulumi:"requestHeaderActions"`
	ResponseHeaderActions      []HeaderActionResponse `pulumi:"responseHeaderActions"`
	RouteConfigurationOverride interface{}            `pulumi:"routeConfigurationOverride"`
}

type RulesEngineActionResponseOutput struct{ *pulumi.OutputState }

func (RulesEngineActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineActionResponse)(nil)).Elem()
}

func (o RulesEngineActionResponseOutput) ToRulesEngineActionResponseOutput() RulesEngineActionResponseOutput {
	return o
}

func (o RulesEngineActionResponseOutput) ToRulesEngineActionResponseOutputWithContext(ctx context.Context) RulesEngineActionResponseOutput {
	return o
}

func (o RulesEngineActionResponseOutput) RequestHeaderActions() HeaderActionResponseArrayOutput {
	return o.ApplyT(func(v RulesEngineActionResponse) []HeaderActionResponse { return v.RequestHeaderActions }).(HeaderActionResponseArrayOutput)
}

func (o RulesEngineActionResponseOutput) ResponseHeaderActions() HeaderActionResponseArrayOutput {
	return o.ApplyT(func(v RulesEngineActionResponse) []HeaderActionResponse { return v.ResponseHeaderActions }).(HeaderActionResponseArrayOutput)
}

func (o RulesEngineActionResponseOutput) RouteConfigurationOverride() pulumi.AnyOutput {
	return o.ApplyT(func(v RulesEngineActionResponse) interface{} { return v.RouteConfigurationOverride }).(pulumi.AnyOutput)
}

type RulesEngineMatchCondition struct {
	NegateCondition          *bool    `pulumi:"negateCondition"`
	RulesEngineMatchValue    []string `pulumi:"rulesEngineMatchValue"`
	RulesEngineMatchVariable string   `pulumi:"rulesEngineMatchVariable"`
	RulesEngineOperator      string   `pulumi:"rulesEngineOperator"`
	Selector                 *string  `pulumi:"selector"`
	Transforms               []string `pulumi:"transforms"`
}





type RulesEngineMatchConditionInput interface {
	pulumi.Input

	ToRulesEngineMatchConditionOutput() RulesEngineMatchConditionOutput
	ToRulesEngineMatchConditionOutputWithContext(context.Context) RulesEngineMatchConditionOutput
}

type RulesEngineMatchConditionArgs struct {
	NegateCondition          pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	RulesEngineMatchValue    pulumi.StringArrayInput `pulumi:"rulesEngineMatchValue"`
	RulesEngineMatchVariable pulumi.StringInput      `pulumi:"rulesEngineMatchVariable"`
	RulesEngineOperator      pulumi.StringInput      `pulumi:"rulesEngineOperator"`
	Selector                 pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms               pulumi.StringArrayInput `pulumi:"transforms"`
}

func (RulesEngineMatchConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineMatchCondition)(nil)).Elem()
}

func (i RulesEngineMatchConditionArgs) ToRulesEngineMatchConditionOutput() RulesEngineMatchConditionOutput {
	return i.ToRulesEngineMatchConditionOutputWithContext(context.Background())
}

func (i RulesEngineMatchConditionArgs) ToRulesEngineMatchConditionOutputWithContext(ctx context.Context) RulesEngineMatchConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineMatchConditionOutput)
}





type RulesEngineMatchConditionArrayInput interface {
	pulumi.Input

	ToRulesEngineMatchConditionArrayOutput() RulesEngineMatchConditionArrayOutput
	ToRulesEngineMatchConditionArrayOutputWithContext(context.Context) RulesEngineMatchConditionArrayOutput
}

type RulesEngineMatchConditionArray []RulesEngineMatchConditionInput

func (RulesEngineMatchConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineMatchCondition)(nil)).Elem()
}

func (i RulesEngineMatchConditionArray) ToRulesEngineMatchConditionArrayOutput() RulesEngineMatchConditionArrayOutput {
	return i.ToRulesEngineMatchConditionArrayOutputWithContext(context.Background())
}

func (i RulesEngineMatchConditionArray) ToRulesEngineMatchConditionArrayOutputWithContext(ctx context.Context) RulesEngineMatchConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineMatchConditionArrayOutput)
}

type RulesEngineMatchConditionOutput struct{ *pulumi.OutputState }

func (RulesEngineMatchConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineMatchCondition)(nil)).Elem()
}

func (o RulesEngineMatchConditionOutput) ToRulesEngineMatchConditionOutput() RulesEngineMatchConditionOutput {
	return o
}

func (o RulesEngineMatchConditionOutput) ToRulesEngineMatchConditionOutputWithContext(ctx context.Context) RulesEngineMatchConditionOutput {
	return o
}

func (o RulesEngineMatchConditionOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RulesEngineMatchConditionOutput) RulesEngineMatchValue() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) []string { return v.RulesEngineMatchValue }).(pulumi.StringArrayOutput)
}

func (o RulesEngineMatchConditionOutput) RulesEngineMatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) string { return v.RulesEngineMatchVariable }).(pulumi.StringOutput)
}

func (o RulesEngineMatchConditionOutput) RulesEngineOperator() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) string { return v.RulesEngineOperator }).(pulumi.StringOutput)
}

func (o RulesEngineMatchConditionOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o RulesEngineMatchConditionOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RulesEngineMatchConditionArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineMatchConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineMatchCondition)(nil)).Elem()
}

func (o RulesEngineMatchConditionArrayOutput) ToRulesEngineMatchConditionArrayOutput() RulesEngineMatchConditionArrayOutput {
	return o
}

func (o RulesEngineMatchConditionArrayOutput) ToRulesEngineMatchConditionArrayOutputWithContext(ctx context.Context) RulesEngineMatchConditionArrayOutput {
	return o
}

func (o RulesEngineMatchConditionArrayOutput) Index(i pulumi.IntInput) RulesEngineMatchConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineMatchCondition {
		return vs[0].([]RulesEngineMatchCondition)[vs[1].(int)]
	}).(RulesEngineMatchConditionOutput)
}

type RulesEngineMatchConditionResponse struct {
	NegateCondition          *bool    `pulumi:"negateCondition"`
	RulesEngineMatchValue    []string `pulumi:"rulesEngineMatchValue"`
	RulesEngineMatchVariable string   `pulumi:"rulesEngineMatchVariable"`
	RulesEngineOperator      string   `pulumi:"rulesEngineOperator"`
	Selector                 *string  `pulumi:"selector"`
	Transforms               []string `pulumi:"transforms"`
}

type RulesEngineMatchConditionResponseOutput struct{ *pulumi.OutputState }

func (RulesEngineMatchConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineMatchConditionResponse)(nil)).Elem()
}

func (o RulesEngineMatchConditionResponseOutput) ToRulesEngineMatchConditionResponseOutput() RulesEngineMatchConditionResponseOutput {
	return o
}

func (o RulesEngineMatchConditionResponseOutput) ToRulesEngineMatchConditionResponseOutputWithContext(ctx context.Context) RulesEngineMatchConditionResponseOutput {
	return o
}

func (o RulesEngineMatchConditionResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RulesEngineMatchConditionResponseOutput) RulesEngineMatchValue() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) []string { return v.RulesEngineMatchValue }).(pulumi.StringArrayOutput)
}

func (o RulesEngineMatchConditionResponseOutput) RulesEngineMatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) string { return v.RulesEngineMatchVariable }).(pulumi.StringOutput)
}

func (o RulesEngineMatchConditionResponseOutput) RulesEngineOperator() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) string { return v.RulesEngineOperator }).(pulumi.StringOutput)
}

func (o RulesEngineMatchConditionResponseOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o RulesEngineMatchConditionResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RulesEngineMatchConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineMatchConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineMatchConditionResponse)(nil)).Elem()
}

func (o RulesEngineMatchConditionResponseArrayOutput) ToRulesEngineMatchConditionResponseArrayOutput() RulesEngineMatchConditionResponseArrayOutput {
	return o
}

func (o RulesEngineMatchConditionResponseArrayOutput) ToRulesEngineMatchConditionResponseArrayOutputWithContext(ctx context.Context) RulesEngineMatchConditionResponseArrayOutput {
	return o
}

func (o RulesEngineMatchConditionResponseArrayOutput) Index(i pulumi.IntInput) RulesEngineMatchConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineMatchConditionResponse {
		return vs[0].([]RulesEngineMatchConditionResponse)[vs[1].(int)]
	}).(RulesEngineMatchConditionResponseOutput)
}

type RulesEngineResponse struct {
	Id            string                    `pulumi:"id"`
	Name          string                    `pulumi:"name"`
	ResourceState string                    `pulumi:"resourceState"`
	Rules         []RulesEngineRuleResponse `pulumi:"rules"`
	Type          string                    `pulumi:"type"`
}

type RulesEngineResponseOutput struct{ *pulumi.OutputState }

func (RulesEngineResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineResponse)(nil)).Elem()
}

func (o RulesEngineResponseOutput) ToRulesEngineResponseOutput() RulesEngineResponseOutput {
	return o
}

func (o RulesEngineResponseOutput) ToRulesEngineResponseOutputWithContext(ctx context.Context) RulesEngineResponseOutput {
	return o
}

func (o RulesEngineResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o RulesEngineResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RulesEngineResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o RulesEngineResponseOutput) Rules() RulesEngineRuleResponseArrayOutput {
	return o.ApplyT(func(v RulesEngineResponse) []RulesEngineRuleResponse { return v.Rules }).(RulesEngineRuleResponseArrayOutput)
}

func (o RulesEngineResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RulesEngineResponseArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineResponse)(nil)).Elem()
}

func (o RulesEngineResponseArrayOutput) ToRulesEngineResponseArrayOutput() RulesEngineResponseArrayOutput {
	return o
}

func (o RulesEngineResponseArrayOutput) ToRulesEngineResponseArrayOutputWithContext(ctx context.Context) RulesEngineResponseArrayOutput {
	return o
}

func (o RulesEngineResponseArrayOutput) Index(i pulumi.IntInput) RulesEngineResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineResponse {
		return vs[0].([]RulesEngineResponse)[vs[1].(int)]
	}).(RulesEngineResponseOutput)
}

type RulesEngineRule struct {
	Action                  RulesEngineAction           `pulumi:"action"`
	MatchConditions         []RulesEngineMatchCondition `pulumi:"matchConditions"`
	MatchProcessingBehavior *string                     `pulumi:"matchProcessingBehavior"`
	Name                    string                      `pulumi:"name"`
	Priority                int                         `pulumi:"priority"`
}





type RulesEngineRuleInput interface {
	pulumi.Input

	ToRulesEngineRuleOutput() RulesEngineRuleOutput
	ToRulesEngineRuleOutputWithContext(context.Context) RulesEngineRuleOutput
}

type RulesEngineRuleArgs struct {
	Action                  RulesEngineActionInput              `pulumi:"action"`
	MatchConditions         RulesEngineMatchConditionArrayInput `pulumi:"matchConditions"`
	MatchProcessingBehavior pulumi.StringPtrInput               `pulumi:"matchProcessingBehavior"`
	Name                    pulumi.StringInput                  `pulumi:"name"`
	Priority                pulumi.IntInput                     `pulumi:"priority"`
}

func (RulesEngineRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineRule)(nil)).Elem()
}

func (i RulesEngineRuleArgs) ToRulesEngineRuleOutput() RulesEngineRuleOutput {
	return i.ToRulesEngineRuleOutputWithContext(context.Background())
}

func (i RulesEngineRuleArgs) ToRulesEngineRuleOutputWithContext(ctx context.Context) RulesEngineRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineRuleOutput)
}





type RulesEngineRuleArrayInput interface {
	pulumi.Input

	ToRulesEngineRuleArrayOutput() RulesEngineRuleArrayOutput
	ToRulesEngineRuleArrayOutputWithContext(context.Context) RulesEngineRuleArrayOutput
}

type RulesEngineRuleArray []RulesEngineRuleInput

func (RulesEngineRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineRule)(nil)).Elem()
}

func (i RulesEngineRuleArray) ToRulesEngineRuleArrayOutput() RulesEngineRuleArrayOutput {
	return i.ToRulesEngineRuleArrayOutputWithContext(context.Background())
}

func (i RulesEngineRuleArray) ToRulesEngineRuleArrayOutputWithContext(ctx context.Context) RulesEngineRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineRuleArrayOutput)
}

type RulesEngineRuleOutput struct{ *pulumi.OutputState }

func (RulesEngineRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineRule)(nil)).Elem()
}

func (o RulesEngineRuleOutput) ToRulesEngineRuleOutput() RulesEngineRuleOutput {
	return o
}

func (o RulesEngineRuleOutput) ToRulesEngineRuleOutputWithContext(ctx context.Context) RulesEngineRuleOutput {
	return o
}

func (o RulesEngineRuleOutput) Action() RulesEngineActionOutput {
	return o.ApplyT(func(v RulesEngineRule) RulesEngineAction { return v.Action }).(RulesEngineActionOutput)
}

func (o RulesEngineRuleOutput) MatchConditions() RulesEngineMatchConditionArrayOutput {
	return o.ApplyT(func(v RulesEngineRule) []RulesEngineMatchCondition { return v.MatchConditions }).(RulesEngineMatchConditionArrayOutput)
}

func (o RulesEngineRuleOutput) MatchProcessingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesEngineRule) *string { return v.MatchProcessingBehavior }).(pulumi.StringPtrOutput)
}

func (o RulesEngineRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineRule) string { return v.Name }).(pulumi.StringOutput)
}

func (o RulesEngineRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v RulesEngineRule) int { return v.Priority }).(pulumi.IntOutput)
}

type RulesEngineRuleArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineRule)(nil)).Elem()
}

func (o RulesEngineRuleArrayOutput) ToRulesEngineRuleArrayOutput() RulesEngineRuleArrayOutput {
	return o
}

func (o RulesEngineRuleArrayOutput) ToRulesEngineRuleArrayOutputWithContext(ctx context.Context) RulesEngineRuleArrayOutput {
	return o
}

func (o RulesEngineRuleArrayOutput) Index(i pulumi.IntInput) RulesEngineRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineRule {
		return vs[0].([]RulesEngineRule)[vs[1].(int)]
	}).(RulesEngineRuleOutput)
}

type RulesEngineRuleResponse struct {
	Action                  RulesEngineActionResponse           `pulumi:"action"`
	MatchConditions         []RulesEngineMatchConditionResponse `pulumi:"matchConditions"`
	MatchProcessingBehavior *string                             `pulumi:"matchProcessingBehavior"`
	Name                    string                              `pulumi:"name"`
	Priority                int                                 `pulumi:"priority"`
}

type RulesEngineRuleResponseOutput struct{ *pulumi.OutputState }

func (RulesEngineRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineRuleResponse)(nil)).Elem()
}

func (o RulesEngineRuleResponseOutput) ToRulesEngineRuleResponseOutput() RulesEngineRuleResponseOutput {
	return o
}

func (o RulesEngineRuleResponseOutput) ToRulesEngineRuleResponseOutputWithContext(ctx context.Context) RulesEngineRuleResponseOutput {
	return o
}

func (o RulesEngineRuleResponseOutput) Action() RulesEngineActionResponseOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) RulesEngineActionResponse { return v.Action }).(RulesEngineActionResponseOutput)
}

func (o RulesEngineRuleResponseOutput) MatchConditions() RulesEngineMatchConditionResponseArrayOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) []RulesEngineMatchConditionResponse { return v.MatchConditions }).(RulesEngineMatchConditionResponseArrayOutput)
}

func (o RulesEngineRuleResponseOutput) MatchProcessingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) *string { return v.MatchProcessingBehavior }).(pulumi.StringPtrOutput)
}

func (o RulesEngineRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RulesEngineRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

type RulesEngineRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineRuleResponse)(nil)).Elem()
}

func (o RulesEngineRuleResponseArrayOutput) ToRulesEngineRuleResponseArrayOutput() RulesEngineRuleResponseArrayOutput {
	return o
}

func (o RulesEngineRuleResponseArrayOutput) ToRulesEngineRuleResponseArrayOutputWithContext(ctx context.Context) RulesEngineRuleResponseArrayOutput {
	return o
}

func (o RulesEngineRuleResponseArrayOutput) Index(i pulumi.IntInput) RulesEngineRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineRuleResponse {
		return vs[0].([]RulesEngineRuleResponse)[vs[1].(int)]
	}).(RulesEngineRuleResponseOutput)
}

type SecurityPolicyLinkResponse struct {
	Id *string `pulumi:"id"`
}

type SecurityPolicyLinkResponseOutput struct{ *pulumi.OutputState }

func (SecurityPolicyLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPolicyLinkResponse)(nil)).Elem()
}

func (o SecurityPolicyLinkResponseOutput) ToSecurityPolicyLinkResponseOutput() SecurityPolicyLinkResponseOutput {
	return o
}

func (o SecurityPolicyLinkResponseOutput) ToSecurityPolicyLinkResponseOutputWithContext(ctx context.Context) SecurityPolicyLinkResponseOutput {
	return o
}

func (o SecurityPolicyLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityPolicyLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SecurityPolicyLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (SecurityPolicyLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityPolicyLinkResponse)(nil)).Elem()
}

func (o SecurityPolicyLinkResponseArrayOutput) ToSecurityPolicyLinkResponseArrayOutput() SecurityPolicyLinkResponseArrayOutput {
	return o
}

func (o SecurityPolicyLinkResponseArrayOutput) ToSecurityPolicyLinkResponseArrayOutputWithContext(ctx context.Context) SecurityPolicyLinkResponseArrayOutput {
	return o
}

func (o SecurityPolicyLinkResponseArrayOutput) Index(i pulumi.IntInput) SecurityPolicyLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityPolicyLinkResponse {
		return vs[0].([]SecurityPolicyLinkResponse)[vs[1].(int)]
	}).(SecurityPolicyLinkResponseOutput)
}

type Sku struct {
	Name *string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
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





type SubResourceArrayInput interface {
	pulumi.Input

	ToSubResourceArrayOutput() SubResourceArrayOutput
	ToSubResourceArrayOutputWithContext(context.Context) SubResourceArrayOutput
}

type SubResourceArray []SubResourceInput

func (SubResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (i SubResourceArray) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return i.ToSubResourceArrayOutputWithContext(context.Background())
}

func (i SubResourceArray) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceArrayOutput)
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

type SubResourceArrayOutput struct{ *pulumi.OutputState }

func (SubResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) Index(i pulumi.IntInput) SubResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResource {
		return vs[0].([]SubResource)[vs[1].(int)]
	}).(SubResourceOutput)
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

type SubResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) Index(i pulumi.IntInput) SubResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceResponse {
		return vs[0].([]SubResourceResponse)[vs[1].(int)]
	}).(SubResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(BackendOutput{})
	pulumi.RegisterOutputType(BackendArrayOutput{})
	pulumi.RegisterOutputType(BackendPoolOutput{})
	pulumi.RegisterOutputType(BackendPoolArrayOutput{})
	pulumi.RegisterOutputType(BackendPoolResponseOutput{})
	pulumi.RegisterOutputType(BackendPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(BackendPoolsSettingsOutput{})
	pulumi.RegisterOutputType(BackendPoolsSettingsPtrOutput{})
	pulumi.RegisterOutputType(BackendPoolsSettingsResponseOutput{})
	pulumi.RegisterOutputType(BackendPoolsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(BackendResponseOutput{})
	pulumi.RegisterOutputType(BackendResponseArrayOutput{})
	pulumi.RegisterOutputType(CustomHttpsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CustomRuleOutput{})
	pulumi.RegisterOutputType(CustomRuleArrayOutput{})
	pulumi.RegisterOutputType(CustomRuleListOutput{})
	pulumi.RegisterOutputType(CustomRuleListPtrOutput{})
	pulumi.RegisterOutputType(CustomRuleListResponseOutput{})
	pulumi.RegisterOutputType(CustomRuleListResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomRuleResponseOutput{})
	pulumi.RegisterOutputType(CustomRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleGroupOverrideOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleGroupOverrideArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleGroupOverrideResponseOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleGroupOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleOverrideOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleOverrideArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleOverrideResponseOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleSetOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleSetArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleSetResponseOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleSetResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorMatchConditionOutput{})
	pulumi.RegisterOutputType(FrontDoorMatchConditionArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorMatchConditionResponseOutput{})
	pulumi.RegisterOutputType(FrontDoorMatchConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorPolicySettingsOutput{})
	pulumi.RegisterOutputType(FrontDoorPolicySettingsPtrOutput{})
	pulumi.RegisterOutputType(FrontDoorPolicySettingsResponseOutput{})
	pulumi.RegisterOutputType(FrontDoorPolicySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(FrontendEndpointOutput{})
	pulumi.RegisterOutputType(FrontendEndpointArrayOutput{})
	pulumi.RegisterOutputType(FrontendEndpointLinkResponseOutput{})
	pulumi.RegisterOutputType(FrontendEndpointLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontendEndpointResponseOutput{})
	pulumi.RegisterOutputType(FrontendEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput{})
	pulumi.RegisterOutputType(FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput{})
	pulumi.RegisterOutputType(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput{})
	pulumi.RegisterOutputType(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput{})
	pulumi.RegisterOutputType(HeaderActionOutput{})
	pulumi.RegisterOutputType(HeaderActionArrayOutput{})
	pulumi.RegisterOutputType(HeaderActionResponseOutput{})
	pulumi.RegisterOutputType(HeaderActionResponseArrayOutput{})
	pulumi.RegisterOutputType(HealthProbeSettingsModelOutput{})
	pulumi.RegisterOutputType(HealthProbeSettingsModelArrayOutput{})
	pulumi.RegisterOutputType(HealthProbeSettingsModelResponseOutput{})
	pulumi.RegisterOutputType(HealthProbeSettingsModelResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultCertificateSourceParametersResponseVaultOutput{})
	pulumi.RegisterOutputType(KeyVaultCertificateSourceParametersResponseVaultPtrOutput{})
	pulumi.RegisterOutputType(LoadBalancingSettingsModelOutput{})
	pulumi.RegisterOutputType(LoadBalancingSettingsModelArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancingSettingsModelResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancingSettingsModelResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedRuleExclusionOutput{})
	pulumi.RegisterOutputType(ManagedRuleExclusionArrayOutput{})
	pulumi.RegisterOutputType(ManagedRuleExclusionResponseOutput{})
	pulumi.RegisterOutputType(ManagedRuleExclusionResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetListOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetListPtrOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetListResponseOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetListResponsePtrOutput{})
	pulumi.RegisterOutputType(RoutingRuleOutput{})
	pulumi.RegisterOutputType(RoutingRuleArrayOutput{})
	pulumi.RegisterOutputType(RoutingRuleLinkResponseOutput{})
	pulumi.RegisterOutputType(RoutingRuleLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(RoutingRuleResponseOutput{})
	pulumi.RegisterOutputType(RoutingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput{})
	pulumi.RegisterOutputType(RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput{})
	pulumi.RegisterOutputType(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput{})
	pulumi.RegisterOutputType(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput{})
	pulumi.RegisterOutputType(RulesEngineActionOutput{})
	pulumi.RegisterOutputType(RulesEngineActionResponseOutput{})
	pulumi.RegisterOutputType(RulesEngineMatchConditionOutput{})
	pulumi.RegisterOutputType(RulesEngineMatchConditionArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineMatchConditionResponseOutput{})
	pulumi.RegisterOutputType(RulesEngineMatchConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineResponseOutput{})
	pulumi.RegisterOutputType(RulesEngineResponseArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineRuleOutput{})
	pulumi.RegisterOutputType(RulesEngineRuleArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineRuleResponseOutput{})
	pulumi.RegisterOutputType(RulesEngineRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SecurityPolicyLinkResponseOutput{})
	pulumi.RegisterOutputType(SecurityPolicyLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
}

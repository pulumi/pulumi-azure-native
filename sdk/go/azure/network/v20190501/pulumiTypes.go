


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Backend struct {
	Address           *string `pulumi:"address"`
	BackendHostHeader *string `pulumi:"backendHostHeader"`
	EnabledState      *string `pulumi:"enabledState"`
	HttpPort          *int    `pulumi:"httpPort"`
	HttpsPort         *int    `pulumi:"httpsPort"`
	Priority          *int    `pulumi:"priority"`
	Weight            *int    `pulumi:"weight"`
}





type BackendInput interface {
	pulumi.Input

	ToBackendOutput() BackendOutput
	ToBackendOutputWithContext(context.Context) BackendOutput
}

type BackendArgs struct {
	Address           pulumi.StringPtrInput `pulumi:"address"`
	BackendHostHeader pulumi.StringPtrInput `pulumi:"backendHostHeader"`
	EnabledState      pulumi.StringPtrInput `pulumi:"enabledState"`
	HttpPort          pulumi.IntPtrInput    `pulumi:"httpPort"`
	HttpsPort         pulumi.IntPtrInput    `pulumi:"httpsPort"`
	Priority          pulumi.IntPtrInput    `pulumi:"priority"`
	Weight            pulumi.IntPtrInput    `pulumi:"weight"`
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
	Address           *string `pulumi:"address"`
	BackendHostHeader *string `pulumi:"backendHostHeader"`
	EnabledState      *string `pulumi:"enabledState"`
	HttpPort          *int    `pulumi:"httpPort"`
	HttpsPort         *int    `pulumi:"httpsPort"`
	Priority          *int    `pulumi:"priority"`
	Weight            *int    `pulumi:"weight"`
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
	DynamicCompression           *string `pulumi:"dynamicCompression"`
	QueryParameterStripDirective *string `pulumi:"queryParameterStripDirective"`
}

type CacheConfigurationResponse struct {
	DynamicCompression           *string `pulumi:"dynamicCompression"`
	QueryParameterStripDirective *string `pulumi:"queryParameterStripDirective"`
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
	AcceptedProtocols  []string      `pulumi:"acceptedProtocols"`
	EnabledState       *string       `pulumi:"enabledState"`
	FrontendEndpoints  []SubResource `pulumi:"frontendEndpoints"`
	Id                 *string       `pulumi:"id"`
	Name               *string       `pulumi:"name"`
	PatternsToMatch    []string      `pulumi:"patternsToMatch"`
	RouteConfiguration interface{}   `pulumi:"routeConfiguration"`
}





type RoutingRuleInput interface {
	pulumi.Input

	ToRoutingRuleOutput() RoutingRuleOutput
	ToRoutingRuleOutputWithContext(context.Context) RoutingRuleOutput
}

type RoutingRuleArgs struct {
	AcceptedProtocols  pulumi.StringArrayInput `pulumi:"acceptedProtocols"`
	EnabledState       pulumi.StringPtrInput   `pulumi:"enabledState"`
	FrontendEndpoints  SubResourceArrayInput   `pulumi:"frontendEndpoints"`
	Id                 pulumi.StringPtrInput   `pulumi:"id"`
	Name               pulumi.StringPtrInput   `pulumi:"name"`
	PatternsToMatch    pulumi.StringArrayInput `pulumi:"patternsToMatch"`
	RouteConfiguration pulumi.Input            `pulumi:"routeConfiguration"`
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

type RoutingRuleResponse struct {
	AcceptedProtocols  []string              `pulumi:"acceptedProtocols"`
	EnabledState       *string               `pulumi:"enabledState"`
	FrontendEndpoints  []SubResourceResponse `pulumi:"frontendEndpoints"`
	Id                 *string               `pulumi:"id"`
	Name               *string               `pulumi:"name"`
	PatternsToMatch    []string              `pulumi:"patternsToMatch"`
	ResourceState      string                `pulumi:"resourceState"`
	RouteConfiguration interface{}           `pulumi:"routeConfiguration"`
	Type               string                `pulumi:"type"`
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

func (o RoutingRuleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingRuleResponse) string { return v.Type }).(pulumi.StringOutput)
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
	pulumi.RegisterOutputType(FrontendEndpointOutput{})
	pulumi.RegisterOutputType(FrontendEndpointArrayOutput{})
	pulumi.RegisterOutputType(FrontendEndpointResponseOutput{})
	pulumi.RegisterOutputType(FrontendEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput{})
	pulumi.RegisterOutputType(FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput{})
	pulumi.RegisterOutputType(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput{})
	pulumi.RegisterOutputType(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput{})
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
	pulumi.RegisterOutputType(RoutingRuleOutput{})
	pulumi.RegisterOutputType(RoutingRuleArrayOutput{})
	pulumi.RegisterOutputType(RoutingRuleResponseOutput{})
	pulumi.RegisterOutputType(RoutingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
}

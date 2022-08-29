


package v20220201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSignalR(ctx *pulumi.Context, args *LookupSignalRArgs, opts ...pulumi.InvokeOption) (*LookupSignalRResult, error) {
	var rv LookupSignalRResult
	err := ctx.Invoke("azure-native:signalrservice/v20220201:getSignalR", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSignalRArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupSignalRResult struct {
	Cors                       *SignalRCorsSettingsResponse        `pulumi:"cors"`
	DisableAadAuth             *bool                               `pulumi:"disableAadAuth"`
	DisableLocalAuth           *bool                               `pulumi:"disableLocalAuth"`
	ExternalIP                 string                              `pulumi:"externalIP"`
	Features                   []SignalRFeatureResponse            `pulumi:"features"`
	HostName                   string                              `pulumi:"hostName"`
	HostNamePrefix             string                              `pulumi:"hostNamePrefix"`
	Id                         string                              `pulumi:"id"`
	Identity                   *ManagedIdentityResponse            `pulumi:"identity"`
	Kind                       *string                             `pulumi:"kind"`
	LiveTraceConfiguration     *LiveTraceConfigurationResponse     `pulumi:"liveTraceConfiguration"`
	Location                   *string                             `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	NetworkACLs                *SignalRNetworkACLsResponse         `pulumi:"networkACLs"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                             `pulumi:"publicNetworkAccess"`
	PublicPort                 int                                 `pulumi:"publicPort"`
	ResourceLogConfiguration   *ResourceLogConfigurationResponse   `pulumi:"resourceLogConfiguration"`
	ServerPort                 int                                 `pulumi:"serverPort"`
	SharedPrivateLinkResources []SharedPrivateLinkResourceResponse `pulumi:"sharedPrivateLinkResources"`
	Sku                        *ResourceSkuResponse                `pulumi:"sku"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Tls                        *SignalRTlsSettingsResponse         `pulumi:"tls"`
	Type                       string                              `pulumi:"type"`
	Upstream                   *ServerlessUpstreamSettingsResponse `pulumi:"upstream"`
	Version                    string                              `pulumi:"version"`
}


func (val *LookupSignalRResult) Defaults() *LookupSignalRResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DisableAadAuth) {
		disableAadAuth_ := false
		tmp.DisableAadAuth = &disableAadAuth_
	}
	if isZero(tmp.DisableLocalAuth) {
		disableLocalAuth_ := false
		tmp.DisableLocalAuth = &disableLocalAuth_
	}
	tmp.LiveTraceConfiguration = tmp.LiveTraceConfiguration.Defaults()

	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	tmp.Tls = tmp.Tls.Defaults()

	return &tmp
}

func LookupSignalROutput(ctx *pulumi.Context, args LookupSignalROutputArgs, opts ...pulumi.InvokeOption) LookupSignalRResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSignalRResult, error) {
			args := v.(LookupSignalRArgs)
			r, err := LookupSignalR(ctx, &args, opts...)
			var s LookupSignalRResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSignalRResultOutput)
}

type LookupSignalROutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSignalROutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRArgs)(nil)).Elem()
}


type LookupSignalRResultOutput struct{ *pulumi.OutputState }

func (LookupSignalRResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRResult)(nil)).Elem()
}

func (o LookupSignalRResultOutput) ToLookupSignalRResultOutput() LookupSignalRResultOutput {
	return o
}

func (o LookupSignalRResultOutput) ToLookupSignalRResultOutputWithContext(ctx context.Context) LookupSignalRResultOutput {
	return o
}

func (o LookupSignalRResultOutput) Cors() SignalRCorsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *SignalRCorsSettingsResponse { return v.Cors }).(SignalRCorsSettingsResponsePtrOutput)
}

func (o LookupSignalRResultOutput) DisableAadAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *bool { return v.DisableAadAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupSignalRResultOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupSignalRResultOutput) ExternalIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.ExternalIP }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) Features() SignalRFeatureResponseArrayOutput {
	return o.ApplyT(func(v LookupSignalRResult) []SignalRFeatureResponse { return v.Features }).(SignalRFeatureResponseArrayOutput)
}

func (o LookupSignalRResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) HostNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.HostNamePrefix }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *ManagedIdentityResponse { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

func (o LookupSignalRResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSignalRResultOutput) LiveTraceConfiguration() LiveTraceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *LiveTraceConfigurationResponse { return v.LiveTraceConfiguration }).(LiveTraceConfigurationResponsePtrOutput)
}

func (o LookupSignalRResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSignalRResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) NetworkACLs() SignalRNetworkACLsResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *SignalRNetworkACLsResponse { return v.NetworkACLs }).(SignalRNetworkACLsResponsePtrOutput)
}

func (o LookupSignalRResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupSignalRResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupSignalRResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupSignalRResultOutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSignalRResult) int { return v.PublicPort }).(pulumi.IntOutput)
}

func (o LookupSignalRResultOutput) ResourceLogConfiguration() ResourceLogConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *ResourceLogConfigurationResponse { return v.ResourceLogConfiguration }).(ResourceLogConfigurationResponsePtrOutput)
}

func (o LookupSignalRResultOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSignalRResult) int { return v.ServerPort }).(pulumi.IntOutput)
}

func (o LookupSignalRResultOutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupSignalRResult) []SharedPrivateLinkResourceResponse { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

func (o LookupSignalRResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

func (o LookupSignalRResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSignalRResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSignalRResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSignalRResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSignalRResultOutput) Tls() SignalRTlsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *SignalRTlsSettingsResponse { return v.Tls }).(SignalRTlsSettingsResponsePtrOutput)
}

func (o LookupSignalRResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) Upstream() ServerlessUpstreamSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *ServerlessUpstreamSettingsResponse { return v.Upstream }).(ServerlessUpstreamSettingsResponsePtrOutput)
}

func (o LookupSignalRResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalRResultOutput{})
}

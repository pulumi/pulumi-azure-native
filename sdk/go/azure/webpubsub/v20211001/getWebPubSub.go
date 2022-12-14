


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebPubSub(ctx *pulumi.Context, args *LookupWebPubSubArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubResult, error) {
	var rv LookupWebPubSubResult
	err := ctx.Invoke("azure-native:webpubsub/v20211001:getWebPubSub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebPubSubArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWebPubSubResult struct {
	DisableAadAuth             *bool                               `pulumi:"disableAadAuth"`
	DisableLocalAuth           *bool                               `pulumi:"disableLocalAuth"`
	ExternalIP                 string                              `pulumi:"externalIP"`
	HostName                   string                              `pulumi:"hostName"`
	HostNamePrefix             string                              `pulumi:"hostNamePrefix"`
	Id                         string                              `pulumi:"id"`
	Identity                   *ManagedIdentityResponse            `pulumi:"identity"`
	LiveTraceConfiguration     *LiveTraceConfigurationResponse     `pulumi:"liveTraceConfiguration"`
	Location                   *string                             `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	NetworkACLs                *WebPubSubNetworkACLsResponse       `pulumi:"networkACLs"`
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
	Tls                        *WebPubSubTlsSettingsResponse       `pulumi:"tls"`
	Type                       string                              `pulumi:"type"`
	Version                    string                              `pulumi:"version"`
}


func (val *LookupWebPubSubResult) Defaults() *LookupWebPubSubResult {
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

func LookupWebPubSubOutput(ctx *pulumi.Context, args LookupWebPubSubOutputArgs, opts ...pulumi.InvokeOption) LookupWebPubSubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebPubSubResult, error) {
			args := v.(LookupWebPubSubArgs)
			r, err := LookupWebPubSub(ctx, &args, opts...)
			var s LookupWebPubSubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebPubSubResultOutput)
}

type LookupWebPubSubOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWebPubSubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubArgs)(nil)).Elem()
}


type LookupWebPubSubResultOutput struct{ *pulumi.OutputState }

func (LookupWebPubSubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubResult)(nil)).Elem()
}

func (o LookupWebPubSubResultOutput) ToLookupWebPubSubResultOutput() LookupWebPubSubResultOutput {
	return o
}

func (o LookupWebPubSubResultOutput) ToLookupWebPubSubResultOutputWithContext(ctx context.Context) LookupWebPubSubResultOutput {
	return o
}

func (o LookupWebPubSubResultOutput) DisableAadAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *bool { return v.DisableAadAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupWebPubSubResultOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupWebPubSubResultOutput) ExternalIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.ExternalIP }).(pulumi.StringOutput)
}

func (o LookupWebPubSubResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupWebPubSubResultOutput) HostNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.HostNamePrefix }).(pulumi.StringOutput)
}

func (o LookupWebPubSubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebPubSubResultOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *ManagedIdentityResponse { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

func (o LookupWebPubSubResultOutput) LiveTraceConfiguration() LiveTraceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *LiveTraceConfigurationResponse { return v.LiveTraceConfiguration }).(LiveTraceConfigurationResponsePtrOutput)
}

func (o LookupWebPubSubResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupWebPubSubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebPubSubResultOutput) NetworkACLs() WebPubSubNetworkACLsResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *WebPubSubNetworkACLsResponse { return v.NetworkACLs }).(WebPubSubNetworkACLsResponsePtrOutput)
}

func (o LookupWebPubSubResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupWebPubSubResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWebPubSubResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupWebPubSubResultOutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) int { return v.PublicPort }).(pulumi.IntOutput)
}

func (o LookupWebPubSubResultOutput) ResourceLogConfiguration() ResourceLogConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *ResourceLogConfigurationResponse { return v.ResourceLogConfiguration }).(ResourceLogConfigurationResponsePtrOutput)
}

func (o LookupWebPubSubResultOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) int { return v.ServerPort }).(pulumi.IntOutput)
}

func (o LookupWebPubSubResultOutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) []SharedPrivateLinkResourceResponse { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

func (o LookupWebPubSubResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

func (o LookupWebPubSubResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWebPubSubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWebPubSubResultOutput) Tls() WebPubSubTlsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *WebPubSubTlsSettingsResponse { return v.Tls }).(WebPubSubTlsSettingsResponsePtrOutput)
}

func (o LookupWebPubSubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebPubSubResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebPubSubResultOutput{})
}

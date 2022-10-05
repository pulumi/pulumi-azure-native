


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationGetEndpoint struct {
	DestinationPort  *int    `pulumi:"destinationPort"`
	Location         *string `pulumi:"location"`
	PrivateIPAddress *string `pulumi:"privateIPAddress"`
	PublicPort       *int    `pulumi:"publicPort"`
}





type ApplicationGetEndpointInput interface {
	pulumi.Input

	ToApplicationGetEndpointOutput() ApplicationGetEndpointOutput
	ToApplicationGetEndpointOutputWithContext(context.Context) ApplicationGetEndpointOutput
}

type ApplicationGetEndpointArgs struct {
	DestinationPort  pulumi.IntPtrInput    `pulumi:"destinationPort"`
	Location         pulumi.StringPtrInput `pulumi:"location"`
	PrivateIPAddress pulumi.StringPtrInput `pulumi:"privateIPAddress"`
	PublicPort       pulumi.IntPtrInput    `pulumi:"publicPort"`
}

func (ApplicationGetEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGetEndpoint)(nil)).Elem()
}

func (i ApplicationGetEndpointArgs) ToApplicationGetEndpointOutput() ApplicationGetEndpointOutput {
	return i.ToApplicationGetEndpointOutputWithContext(context.Background())
}

func (i ApplicationGetEndpointArgs) ToApplicationGetEndpointOutputWithContext(ctx context.Context) ApplicationGetEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGetEndpointOutput)
}





type ApplicationGetEndpointArrayInput interface {
	pulumi.Input

	ToApplicationGetEndpointArrayOutput() ApplicationGetEndpointArrayOutput
	ToApplicationGetEndpointArrayOutputWithContext(context.Context) ApplicationGetEndpointArrayOutput
}

type ApplicationGetEndpointArray []ApplicationGetEndpointInput

func (ApplicationGetEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGetEndpoint)(nil)).Elem()
}

func (i ApplicationGetEndpointArray) ToApplicationGetEndpointArrayOutput() ApplicationGetEndpointArrayOutput {
	return i.ToApplicationGetEndpointArrayOutputWithContext(context.Background())
}

func (i ApplicationGetEndpointArray) ToApplicationGetEndpointArrayOutputWithContext(ctx context.Context) ApplicationGetEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGetEndpointArrayOutput)
}

type ApplicationGetEndpointOutput struct{ *pulumi.OutputState }

func (ApplicationGetEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGetEndpoint)(nil)).Elem()
}

func (o ApplicationGetEndpointOutput) ToApplicationGetEndpointOutput() ApplicationGetEndpointOutput {
	return o
}

func (o ApplicationGetEndpointOutput) ToApplicationGetEndpointOutputWithContext(ctx context.Context) ApplicationGetEndpointOutput {
	return o
}

func (o ApplicationGetEndpointOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGetEndpoint) *int { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o ApplicationGetEndpointOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGetEndpoint) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicationGetEndpointOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGetEndpoint) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o ApplicationGetEndpointOutput) PublicPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGetEndpoint) *int { return v.PublicPort }).(pulumi.IntPtrOutput)
}

type ApplicationGetEndpointArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGetEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGetEndpoint)(nil)).Elem()
}

func (o ApplicationGetEndpointArrayOutput) ToApplicationGetEndpointArrayOutput() ApplicationGetEndpointArrayOutput {
	return o
}

func (o ApplicationGetEndpointArrayOutput) ToApplicationGetEndpointArrayOutputWithContext(ctx context.Context) ApplicationGetEndpointArrayOutput {
	return o
}

func (o ApplicationGetEndpointArrayOutput) Index(i pulumi.IntInput) ApplicationGetEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGetEndpoint {
		return vs[0].([]ApplicationGetEndpoint)[vs[1].(int)]
	}).(ApplicationGetEndpointOutput)
}

type ApplicationGetEndpointResponse struct {
	DestinationPort  *int    `pulumi:"destinationPort"`
	Location         *string `pulumi:"location"`
	PrivateIPAddress *string `pulumi:"privateIPAddress"`
	PublicPort       *int    `pulumi:"publicPort"`
}

type ApplicationGetEndpointResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGetEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGetEndpointResponse)(nil)).Elem()
}

func (o ApplicationGetEndpointResponseOutput) ToApplicationGetEndpointResponseOutput() ApplicationGetEndpointResponseOutput {
	return o
}

func (o ApplicationGetEndpointResponseOutput) ToApplicationGetEndpointResponseOutputWithContext(ctx context.Context) ApplicationGetEndpointResponseOutput {
	return o
}

func (o ApplicationGetEndpointResponseOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGetEndpointResponse) *int { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o ApplicationGetEndpointResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGetEndpointResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicationGetEndpointResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGetEndpointResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o ApplicationGetEndpointResponseOutput) PublicPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGetEndpointResponse) *int { return v.PublicPort }).(pulumi.IntPtrOutput)
}

type ApplicationGetEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGetEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGetEndpointResponse)(nil)).Elem()
}

func (o ApplicationGetEndpointResponseArrayOutput) ToApplicationGetEndpointResponseArrayOutput() ApplicationGetEndpointResponseArrayOutput {
	return o
}

func (o ApplicationGetEndpointResponseArrayOutput) ToApplicationGetEndpointResponseArrayOutputWithContext(ctx context.Context) ApplicationGetEndpointResponseArrayOutput {
	return o
}

func (o ApplicationGetEndpointResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGetEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGetEndpointResponse {
		return vs[0].([]ApplicationGetEndpointResponse)[vs[1].(int)]
	}).(ApplicationGetEndpointResponseOutput)
}

type ApplicationGetHttpsEndpoint struct {
	AccessModes        []string `pulumi:"accessModes"`
	DestinationPort    *int     `pulumi:"destinationPort"`
	DisableGatewayAuth *bool    `pulumi:"disableGatewayAuth"`
	PrivateIPAddress   *string  `pulumi:"privateIPAddress"`
	SubDomainSuffix    *string  `pulumi:"subDomainSuffix"`
}





type ApplicationGetHttpsEndpointInput interface {
	pulumi.Input

	ToApplicationGetHttpsEndpointOutput() ApplicationGetHttpsEndpointOutput
	ToApplicationGetHttpsEndpointOutputWithContext(context.Context) ApplicationGetHttpsEndpointOutput
}

type ApplicationGetHttpsEndpointArgs struct {
	AccessModes        pulumi.StringArrayInput `pulumi:"accessModes"`
	DestinationPort    pulumi.IntPtrInput      `pulumi:"destinationPort"`
	DisableGatewayAuth pulumi.BoolPtrInput     `pulumi:"disableGatewayAuth"`
	PrivateIPAddress   pulumi.StringPtrInput   `pulumi:"privateIPAddress"`
	SubDomainSuffix    pulumi.StringPtrInput   `pulumi:"subDomainSuffix"`
}

func (ApplicationGetHttpsEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGetHttpsEndpoint)(nil)).Elem()
}

func (i ApplicationGetHttpsEndpointArgs) ToApplicationGetHttpsEndpointOutput() ApplicationGetHttpsEndpointOutput {
	return i.ToApplicationGetHttpsEndpointOutputWithContext(context.Background())
}

func (i ApplicationGetHttpsEndpointArgs) ToApplicationGetHttpsEndpointOutputWithContext(ctx context.Context) ApplicationGetHttpsEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGetHttpsEndpointOutput)
}





type ApplicationGetHttpsEndpointArrayInput interface {
	pulumi.Input

	ToApplicationGetHttpsEndpointArrayOutput() ApplicationGetHttpsEndpointArrayOutput
	ToApplicationGetHttpsEndpointArrayOutputWithContext(context.Context) ApplicationGetHttpsEndpointArrayOutput
}

type ApplicationGetHttpsEndpointArray []ApplicationGetHttpsEndpointInput

func (ApplicationGetHttpsEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGetHttpsEndpoint)(nil)).Elem()
}

func (i ApplicationGetHttpsEndpointArray) ToApplicationGetHttpsEndpointArrayOutput() ApplicationGetHttpsEndpointArrayOutput {
	return i.ToApplicationGetHttpsEndpointArrayOutputWithContext(context.Background())
}

func (i ApplicationGetHttpsEndpointArray) ToApplicationGetHttpsEndpointArrayOutputWithContext(ctx context.Context) ApplicationGetHttpsEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGetHttpsEndpointArrayOutput)
}

type ApplicationGetHttpsEndpointOutput struct{ *pulumi.OutputState }

func (ApplicationGetHttpsEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGetHttpsEndpoint)(nil)).Elem()
}

func (o ApplicationGetHttpsEndpointOutput) ToApplicationGetHttpsEndpointOutput() ApplicationGetHttpsEndpointOutput {
	return o
}

func (o ApplicationGetHttpsEndpointOutput) ToApplicationGetHttpsEndpointOutputWithContext(ctx context.Context) ApplicationGetHttpsEndpointOutput {
	return o
}

func (o ApplicationGetHttpsEndpointOutput) AccessModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGetHttpsEndpoint) []string { return v.AccessModes }).(pulumi.StringArrayOutput)
}

func (o ApplicationGetHttpsEndpointOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGetHttpsEndpoint) *int { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o ApplicationGetHttpsEndpointOutput) DisableGatewayAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGetHttpsEndpoint) *bool { return v.DisableGatewayAuth }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGetHttpsEndpointOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGetHttpsEndpoint) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o ApplicationGetHttpsEndpointOutput) SubDomainSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGetHttpsEndpoint) *string { return v.SubDomainSuffix }).(pulumi.StringPtrOutput)
}

type ApplicationGetHttpsEndpointArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGetHttpsEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGetHttpsEndpoint)(nil)).Elem()
}

func (o ApplicationGetHttpsEndpointArrayOutput) ToApplicationGetHttpsEndpointArrayOutput() ApplicationGetHttpsEndpointArrayOutput {
	return o
}

func (o ApplicationGetHttpsEndpointArrayOutput) ToApplicationGetHttpsEndpointArrayOutputWithContext(ctx context.Context) ApplicationGetHttpsEndpointArrayOutput {
	return o
}

func (o ApplicationGetHttpsEndpointArrayOutput) Index(i pulumi.IntInput) ApplicationGetHttpsEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGetHttpsEndpoint {
		return vs[0].([]ApplicationGetHttpsEndpoint)[vs[1].(int)]
	}).(ApplicationGetHttpsEndpointOutput)
}

type ApplicationGetHttpsEndpointResponse struct {
	AccessModes        []string `pulumi:"accessModes"`
	DestinationPort    *int     `pulumi:"destinationPort"`
	DisableGatewayAuth *bool    `pulumi:"disableGatewayAuth"`
	Location           string   `pulumi:"location"`
	PrivateIPAddress   *string  `pulumi:"privateIPAddress"`
	PublicPort         int      `pulumi:"publicPort"`
}

type ApplicationGetHttpsEndpointResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGetHttpsEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGetHttpsEndpointResponse)(nil)).Elem()
}

func (o ApplicationGetHttpsEndpointResponseOutput) ToApplicationGetHttpsEndpointResponseOutput() ApplicationGetHttpsEndpointResponseOutput {
	return o
}

func (o ApplicationGetHttpsEndpointResponseOutput) ToApplicationGetHttpsEndpointResponseOutputWithContext(ctx context.Context) ApplicationGetHttpsEndpointResponseOutput {
	return o
}

func (o ApplicationGetHttpsEndpointResponseOutput) AccessModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGetHttpsEndpointResponse) []string { return v.AccessModes }).(pulumi.StringArrayOutput)
}

func (o ApplicationGetHttpsEndpointResponseOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGetHttpsEndpointResponse) *int { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o ApplicationGetHttpsEndpointResponseOutput) DisableGatewayAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGetHttpsEndpointResponse) *bool { return v.DisableGatewayAuth }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGetHttpsEndpointResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGetHttpsEndpointResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o ApplicationGetHttpsEndpointResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGetHttpsEndpointResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o ApplicationGetHttpsEndpointResponseOutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v ApplicationGetHttpsEndpointResponse) int { return v.PublicPort }).(pulumi.IntOutput)
}

type ApplicationGetHttpsEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGetHttpsEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGetHttpsEndpointResponse)(nil)).Elem()
}

func (o ApplicationGetHttpsEndpointResponseArrayOutput) ToApplicationGetHttpsEndpointResponseArrayOutput() ApplicationGetHttpsEndpointResponseArrayOutput {
	return o
}

func (o ApplicationGetHttpsEndpointResponseArrayOutput) ToApplicationGetHttpsEndpointResponseArrayOutputWithContext(ctx context.Context) ApplicationGetHttpsEndpointResponseArrayOutput {
	return o
}

func (o ApplicationGetHttpsEndpointResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGetHttpsEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGetHttpsEndpointResponse {
		return vs[0].([]ApplicationGetHttpsEndpointResponse)[vs[1].(int)]
	}).(ApplicationGetHttpsEndpointResponseOutput)
}

type ApplicationProperties struct {
	ApplicationType           *string                       `pulumi:"applicationType"`
	ComputeProfile            *ComputeProfile               `pulumi:"computeProfile"`
	Errors                    []Errors                      `pulumi:"errors"`
	HttpsEndpoints            []ApplicationGetHttpsEndpoint `pulumi:"httpsEndpoints"`
	InstallScriptActions      []RuntimeScriptAction         `pulumi:"installScriptActions"`
	PrivateLinkConfigurations []PrivateLinkConfiguration    `pulumi:"privateLinkConfigurations"`
	SshEndpoints              []ApplicationGetEndpoint      `pulumi:"sshEndpoints"`
	UninstallScriptActions    []RuntimeScriptAction         `pulumi:"uninstallScriptActions"`
}





type ApplicationPropertiesInput interface {
	pulumi.Input

	ToApplicationPropertiesOutput() ApplicationPropertiesOutput
	ToApplicationPropertiesOutputWithContext(context.Context) ApplicationPropertiesOutput
}

type ApplicationPropertiesArgs struct {
	ApplicationType           pulumi.StringPtrInput                 `pulumi:"applicationType"`
	ComputeProfile            ComputeProfilePtrInput                `pulumi:"computeProfile"`
	Errors                    ErrorsArrayInput                      `pulumi:"errors"`
	HttpsEndpoints            ApplicationGetHttpsEndpointArrayInput `pulumi:"httpsEndpoints"`
	InstallScriptActions      RuntimeScriptActionArrayInput         `pulumi:"installScriptActions"`
	PrivateLinkConfigurations PrivateLinkConfigurationArrayInput    `pulumi:"privateLinkConfigurations"`
	SshEndpoints              ApplicationGetEndpointArrayInput      `pulumi:"sshEndpoints"`
	UninstallScriptActions    RuntimeScriptActionArrayInput         `pulumi:"uninstallScriptActions"`
}

func (ApplicationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationProperties)(nil)).Elem()
}

func (i ApplicationPropertiesArgs) ToApplicationPropertiesOutput() ApplicationPropertiesOutput {
	return i.ToApplicationPropertiesOutputWithContext(context.Background())
}

func (i ApplicationPropertiesArgs) ToApplicationPropertiesOutputWithContext(ctx context.Context) ApplicationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPropertiesOutput)
}

func (i ApplicationPropertiesArgs) ToApplicationPropertiesPtrOutput() ApplicationPropertiesPtrOutput {
	return i.ToApplicationPropertiesPtrOutputWithContext(context.Background())
}

func (i ApplicationPropertiesArgs) ToApplicationPropertiesPtrOutputWithContext(ctx context.Context) ApplicationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPropertiesOutput).ToApplicationPropertiesPtrOutputWithContext(ctx)
}









type ApplicationPropertiesPtrInput interface {
	pulumi.Input

	ToApplicationPropertiesPtrOutput() ApplicationPropertiesPtrOutput
	ToApplicationPropertiesPtrOutputWithContext(context.Context) ApplicationPropertiesPtrOutput
}

type applicationPropertiesPtrType ApplicationPropertiesArgs

func ApplicationPropertiesPtr(v *ApplicationPropertiesArgs) ApplicationPropertiesPtrInput {
	return (*applicationPropertiesPtrType)(v)
}

func (*applicationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationProperties)(nil)).Elem()
}

func (i *applicationPropertiesPtrType) ToApplicationPropertiesPtrOutput() ApplicationPropertiesPtrOutput {
	return i.ToApplicationPropertiesPtrOutputWithContext(context.Background())
}

func (i *applicationPropertiesPtrType) ToApplicationPropertiesPtrOutputWithContext(ctx context.Context) ApplicationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPropertiesPtrOutput)
}

type ApplicationPropertiesOutput struct{ *pulumi.OutputState }

func (ApplicationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationProperties)(nil)).Elem()
}

func (o ApplicationPropertiesOutput) ToApplicationPropertiesOutput() ApplicationPropertiesOutput {
	return o
}

func (o ApplicationPropertiesOutput) ToApplicationPropertiesOutputWithContext(ctx context.Context) ApplicationPropertiesOutput {
	return o
}

func (o ApplicationPropertiesOutput) ToApplicationPropertiesPtrOutput() ApplicationPropertiesPtrOutput {
	return o.ToApplicationPropertiesPtrOutputWithContext(context.Background())
}

func (o ApplicationPropertiesOutput) ToApplicationPropertiesPtrOutputWithContext(ctx context.Context) ApplicationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationProperties) *ApplicationProperties {
		return &v
	}).(ApplicationPropertiesPtrOutput)
}

func (o ApplicationPropertiesOutput) ApplicationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationProperties) *string { return v.ApplicationType }).(pulumi.StringPtrOutput)
}

func (o ApplicationPropertiesOutput) ComputeProfile() ComputeProfilePtrOutput {
	return o.ApplyT(func(v ApplicationProperties) *ComputeProfile { return v.ComputeProfile }).(ComputeProfilePtrOutput)
}

func (o ApplicationPropertiesOutput) Errors() ErrorsArrayOutput {
	return o.ApplyT(func(v ApplicationProperties) []Errors { return v.Errors }).(ErrorsArrayOutput)
}

func (o ApplicationPropertiesOutput) HttpsEndpoints() ApplicationGetHttpsEndpointArrayOutput {
	return o.ApplyT(func(v ApplicationProperties) []ApplicationGetHttpsEndpoint { return v.HttpsEndpoints }).(ApplicationGetHttpsEndpointArrayOutput)
}

func (o ApplicationPropertiesOutput) InstallScriptActions() RuntimeScriptActionArrayOutput {
	return o.ApplyT(func(v ApplicationProperties) []RuntimeScriptAction { return v.InstallScriptActions }).(RuntimeScriptActionArrayOutput)
}

func (o ApplicationPropertiesOutput) PrivateLinkConfigurations() PrivateLinkConfigurationArrayOutput {
	return o.ApplyT(func(v ApplicationProperties) []PrivateLinkConfiguration { return v.PrivateLinkConfigurations }).(PrivateLinkConfigurationArrayOutput)
}

func (o ApplicationPropertiesOutput) SshEndpoints() ApplicationGetEndpointArrayOutput {
	return o.ApplyT(func(v ApplicationProperties) []ApplicationGetEndpoint { return v.SshEndpoints }).(ApplicationGetEndpointArrayOutput)
}

func (o ApplicationPropertiesOutput) UninstallScriptActions() RuntimeScriptActionArrayOutput {
	return o.ApplyT(func(v ApplicationProperties) []RuntimeScriptAction { return v.UninstallScriptActions }).(RuntimeScriptActionArrayOutput)
}

type ApplicationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ApplicationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationProperties)(nil)).Elem()
}

func (o ApplicationPropertiesPtrOutput) ToApplicationPropertiesPtrOutput() ApplicationPropertiesPtrOutput {
	return o
}

func (o ApplicationPropertiesPtrOutput) ToApplicationPropertiesPtrOutputWithContext(ctx context.Context) ApplicationPropertiesPtrOutput {
	return o
}

func (o ApplicationPropertiesPtrOutput) Elem() ApplicationPropertiesOutput {
	return o.ApplyT(func(v *ApplicationProperties) ApplicationProperties {
		if v != nil {
			return *v
		}
		var ret ApplicationProperties
		return ret
	}).(ApplicationPropertiesOutput)
}

func (o ApplicationPropertiesPtrOutput) ApplicationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationType
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationPropertiesPtrOutput) ComputeProfile() ComputeProfilePtrOutput {
	return o.ApplyT(func(v *ApplicationProperties) *ComputeProfile {
		if v == nil {
			return nil
		}
		return v.ComputeProfile
	}).(ComputeProfilePtrOutput)
}

func (o ApplicationPropertiesPtrOutput) Errors() ErrorsArrayOutput {
	return o.ApplyT(func(v *ApplicationProperties) []Errors {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(ErrorsArrayOutput)
}

func (o ApplicationPropertiesPtrOutput) HttpsEndpoints() ApplicationGetHttpsEndpointArrayOutput {
	return o.ApplyT(func(v *ApplicationProperties) []ApplicationGetHttpsEndpoint {
		if v == nil {
			return nil
		}
		return v.HttpsEndpoints
	}).(ApplicationGetHttpsEndpointArrayOutput)
}

func (o ApplicationPropertiesPtrOutput) InstallScriptActions() RuntimeScriptActionArrayOutput {
	return o.ApplyT(func(v *ApplicationProperties) []RuntimeScriptAction {
		if v == nil {
			return nil
		}
		return v.InstallScriptActions
	}).(RuntimeScriptActionArrayOutput)
}

func (o ApplicationPropertiesPtrOutput) PrivateLinkConfigurations() PrivateLinkConfigurationArrayOutput {
	return o.ApplyT(func(v *ApplicationProperties) []PrivateLinkConfiguration {
		if v == nil {
			return nil
		}
		return v.PrivateLinkConfigurations
	}).(PrivateLinkConfigurationArrayOutput)
}

func (o ApplicationPropertiesPtrOutput) SshEndpoints() ApplicationGetEndpointArrayOutput {
	return o.ApplyT(func(v *ApplicationProperties) []ApplicationGetEndpoint {
		if v == nil {
			return nil
		}
		return v.SshEndpoints
	}).(ApplicationGetEndpointArrayOutput)
}

func (o ApplicationPropertiesPtrOutput) UninstallScriptActions() RuntimeScriptActionArrayOutput {
	return o.ApplyT(func(v *ApplicationProperties) []RuntimeScriptAction {
		if v == nil {
			return nil
		}
		return v.UninstallScriptActions
	}).(RuntimeScriptActionArrayOutput)
}

type ApplicationPropertiesResponse struct {
	ApplicationState          string                                `pulumi:"applicationState"`
	ApplicationType           *string                               `pulumi:"applicationType"`
	ComputeProfile            *ComputeProfileResponse               `pulumi:"computeProfile"`
	CreatedDate               string                                `pulumi:"createdDate"`
	Errors                    []ErrorsResponse                      `pulumi:"errors"`
	HttpsEndpoints            []ApplicationGetHttpsEndpointResponse `pulumi:"httpsEndpoints"`
	InstallScriptActions      []RuntimeScriptActionResponse         `pulumi:"installScriptActions"`
	MarketplaceIdentifier     string                                `pulumi:"marketplaceIdentifier"`
	PrivateLinkConfigurations []PrivateLinkConfigurationResponse    `pulumi:"privateLinkConfigurations"`
	ProvisioningState         string                                `pulumi:"provisioningState"`
	SshEndpoints              []ApplicationGetEndpointResponse      `pulumi:"sshEndpoints"`
	UninstallScriptActions    []RuntimeScriptActionResponse         `pulumi:"uninstallScriptActions"`
}

type ApplicationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPropertiesResponse)(nil)).Elem()
}

func (o ApplicationPropertiesResponseOutput) ToApplicationPropertiesResponseOutput() ApplicationPropertiesResponseOutput {
	return o
}

func (o ApplicationPropertiesResponseOutput) ToApplicationPropertiesResponseOutputWithContext(ctx context.Context) ApplicationPropertiesResponseOutput {
	return o
}

func (o ApplicationPropertiesResponseOutput) ApplicationState() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPropertiesResponse) string { return v.ApplicationState }).(pulumi.StringOutput)
}

func (o ApplicationPropertiesResponseOutput) ApplicationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPropertiesResponse) *string { return v.ApplicationType }).(pulumi.StringPtrOutput)
}

func (o ApplicationPropertiesResponseOutput) ComputeProfile() ComputeProfileResponsePtrOutput {
	return o.ApplyT(func(v ApplicationPropertiesResponse) *ComputeProfileResponse { return v.ComputeProfile }).(ComputeProfileResponsePtrOutput)
}

func (o ApplicationPropertiesResponseOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPropertiesResponse) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o ApplicationPropertiesResponseOutput) Errors() ErrorsResponseArrayOutput {
	return o.ApplyT(func(v ApplicationPropertiesResponse) []ErrorsResponse { return v.Errors }).(ErrorsResponseArrayOutput)
}

func (o ApplicationPropertiesResponseOutput) HttpsEndpoints() ApplicationGetHttpsEndpointResponseArrayOutput {
	return o.ApplyT(func(v ApplicationPropertiesResponse) []ApplicationGetHttpsEndpointResponse { return v.HttpsEndpoints }).(ApplicationGetHttpsEndpointResponseArrayOutput)
}

func (o ApplicationPropertiesResponseOutput) InstallScriptActions() RuntimeScriptActionResponseArrayOutput {
	return o.ApplyT(func(v ApplicationPropertiesResponse) []RuntimeScriptActionResponse { return v.InstallScriptActions }).(RuntimeScriptActionResponseArrayOutput)
}

func (o ApplicationPropertiesResponseOutput) MarketplaceIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPropertiesResponse) string { return v.MarketplaceIdentifier }).(pulumi.StringOutput)
}

func (o ApplicationPropertiesResponseOutput) PrivateLinkConfigurations() PrivateLinkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ApplicationPropertiesResponse) []PrivateLinkConfigurationResponse {
		return v.PrivateLinkConfigurations
	}).(PrivateLinkConfigurationResponseArrayOutput)
}

func (o ApplicationPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApplicationPropertiesResponseOutput) SshEndpoints() ApplicationGetEndpointResponseArrayOutput {
	return o.ApplyT(func(v ApplicationPropertiesResponse) []ApplicationGetEndpointResponse { return v.SshEndpoints }).(ApplicationGetEndpointResponseArrayOutput)
}

func (o ApplicationPropertiesResponseOutput) UninstallScriptActions() RuntimeScriptActionResponseArrayOutput {
	return o.ApplyT(func(v ApplicationPropertiesResponse) []RuntimeScriptActionResponse { return v.UninstallScriptActions }).(RuntimeScriptActionResponseArrayOutput)
}

type Autoscale struct {
	Capacity   *AutoscaleCapacity   `pulumi:"capacity"`
	Recurrence *AutoscaleRecurrence `pulumi:"recurrence"`
}





type AutoscaleInput interface {
	pulumi.Input

	ToAutoscaleOutput() AutoscaleOutput
	ToAutoscaleOutputWithContext(context.Context) AutoscaleOutput
}

type AutoscaleArgs struct {
	Capacity   AutoscaleCapacityPtrInput   `pulumi:"capacity"`
	Recurrence AutoscaleRecurrencePtrInput `pulumi:"recurrence"`
}

func (AutoscaleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Autoscale)(nil)).Elem()
}

func (i AutoscaleArgs) ToAutoscaleOutput() AutoscaleOutput {
	return i.ToAutoscaleOutputWithContext(context.Background())
}

func (i AutoscaleArgs) ToAutoscaleOutputWithContext(ctx context.Context) AutoscaleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleOutput)
}

func (i AutoscaleArgs) ToAutoscalePtrOutput() AutoscalePtrOutput {
	return i.ToAutoscalePtrOutputWithContext(context.Background())
}

func (i AutoscaleArgs) ToAutoscalePtrOutputWithContext(ctx context.Context) AutoscalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleOutput).ToAutoscalePtrOutputWithContext(ctx)
}









type AutoscalePtrInput interface {
	pulumi.Input

	ToAutoscalePtrOutput() AutoscalePtrOutput
	ToAutoscalePtrOutputWithContext(context.Context) AutoscalePtrOutput
}

type autoscalePtrType AutoscaleArgs

func AutoscalePtr(v *AutoscaleArgs) AutoscalePtrInput {
	return (*autoscalePtrType)(v)
}

func (*autoscalePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Autoscale)(nil)).Elem()
}

func (i *autoscalePtrType) ToAutoscalePtrOutput() AutoscalePtrOutput {
	return i.ToAutoscalePtrOutputWithContext(context.Background())
}

func (i *autoscalePtrType) ToAutoscalePtrOutputWithContext(ctx context.Context) AutoscalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalePtrOutput)
}

type AutoscaleOutput struct{ *pulumi.OutputState }

func (AutoscaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Autoscale)(nil)).Elem()
}

func (o AutoscaleOutput) ToAutoscaleOutput() AutoscaleOutput {
	return o
}

func (o AutoscaleOutput) ToAutoscaleOutputWithContext(ctx context.Context) AutoscaleOutput {
	return o
}

func (o AutoscaleOutput) ToAutoscalePtrOutput() AutoscalePtrOutput {
	return o.ToAutoscalePtrOutputWithContext(context.Background())
}

func (o AutoscaleOutput) ToAutoscalePtrOutputWithContext(ctx context.Context) AutoscalePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Autoscale) *Autoscale {
		return &v
	}).(AutoscalePtrOutput)
}

func (o AutoscaleOutput) Capacity() AutoscaleCapacityPtrOutput {
	return o.ApplyT(func(v Autoscale) *AutoscaleCapacity { return v.Capacity }).(AutoscaleCapacityPtrOutput)
}

func (o AutoscaleOutput) Recurrence() AutoscaleRecurrencePtrOutput {
	return o.ApplyT(func(v Autoscale) *AutoscaleRecurrence { return v.Recurrence }).(AutoscaleRecurrencePtrOutput)
}

type AutoscalePtrOutput struct{ *pulumi.OutputState }

func (AutoscalePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Autoscale)(nil)).Elem()
}

func (o AutoscalePtrOutput) ToAutoscalePtrOutput() AutoscalePtrOutput {
	return o
}

func (o AutoscalePtrOutput) ToAutoscalePtrOutputWithContext(ctx context.Context) AutoscalePtrOutput {
	return o
}

func (o AutoscalePtrOutput) Elem() AutoscaleOutput {
	return o.ApplyT(func(v *Autoscale) Autoscale {
		if v != nil {
			return *v
		}
		var ret Autoscale
		return ret
	}).(AutoscaleOutput)
}

func (o AutoscalePtrOutput) Capacity() AutoscaleCapacityPtrOutput {
	return o.ApplyT(func(v *Autoscale) *AutoscaleCapacity {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(AutoscaleCapacityPtrOutput)
}

func (o AutoscalePtrOutput) Recurrence() AutoscaleRecurrencePtrOutput {
	return o.ApplyT(func(v *Autoscale) *AutoscaleRecurrence {
		if v == nil {
			return nil
		}
		return v.Recurrence
	}).(AutoscaleRecurrencePtrOutput)
}

type AutoscaleCapacity struct {
	MaxInstanceCount *int `pulumi:"maxInstanceCount"`
	MinInstanceCount *int `pulumi:"minInstanceCount"`
}





type AutoscaleCapacityInput interface {
	pulumi.Input

	ToAutoscaleCapacityOutput() AutoscaleCapacityOutput
	ToAutoscaleCapacityOutputWithContext(context.Context) AutoscaleCapacityOutput
}

type AutoscaleCapacityArgs struct {
	MaxInstanceCount pulumi.IntPtrInput `pulumi:"maxInstanceCount"`
	MinInstanceCount pulumi.IntPtrInput `pulumi:"minInstanceCount"`
}

func (AutoscaleCapacityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleCapacity)(nil)).Elem()
}

func (i AutoscaleCapacityArgs) ToAutoscaleCapacityOutput() AutoscaleCapacityOutput {
	return i.ToAutoscaleCapacityOutputWithContext(context.Background())
}

func (i AutoscaleCapacityArgs) ToAutoscaleCapacityOutputWithContext(ctx context.Context) AutoscaleCapacityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleCapacityOutput)
}

func (i AutoscaleCapacityArgs) ToAutoscaleCapacityPtrOutput() AutoscaleCapacityPtrOutput {
	return i.ToAutoscaleCapacityPtrOutputWithContext(context.Background())
}

func (i AutoscaleCapacityArgs) ToAutoscaleCapacityPtrOutputWithContext(ctx context.Context) AutoscaleCapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleCapacityOutput).ToAutoscaleCapacityPtrOutputWithContext(ctx)
}









type AutoscaleCapacityPtrInput interface {
	pulumi.Input

	ToAutoscaleCapacityPtrOutput() AutoscaleCapacityPtrOutput
	ToAutoscaleCapacityPtrOutputWithContext(context.Context) AutoscaleCapacityPtrOutput
}

type autoscaleCapacityPtrType AutoscaleCapacityArgs

func AutoscaleCapacityPtr(v *AutoscaleCapacityArgs) AutoscaleCapacityPtrInput {
	return (*autoscaleCapacityPtrType)(v)
}

func (*autoscaleCapacityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleCapacity)(nil)).Elem()
}

func (i *autoscaleCapacityPtrType) ToAutoscaleCapacityPtrOutput() AutoscaleCapacityPtrOutput {
	return i.ToAutoscaleCapacityPtrOutputWithContext(context.Background())
}

func (i *autoscaleCapacityPtrType) ToAutoscaleCapacityPtrOutputWithContext(ctx context.Context) AutoscaleCapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleCapacityPtrOutput)
}

type AutoscaleCapacityOutput struct{ *pulumi.OutputState }

func (AutoscaleCapacityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleCapacity)(nil)).Elem()
}

func (o AutoscaleCapacityOutput) ToAutoscaleCapacityOutput() AutoscaleCapacityOutput {
	return o
}

func (o AutoscaleCapacityOutput) ToAutoscaleCapacityOutputWithContext(ctx context.Context) AutoscaleCapacityOutput {
	return o
}

func (o AutoscaleCapacityOutput) ToAutoscaleCapacityPtrOutput() AutoscaleCapacityPtrOutput {
	return o.ToAutoscaleCapacityPtrOutputWithContext(context.Background())
}

func (o AutoscaleCapacityOutput) ToAutoscaleCapacityPtrOutputWithContext(ctx context.Context) AutoscaleCapacityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoscaleCapacity) *AutoscaleCapacity {
		return &v
	}).(AutoscaleCapacityPtrOutput)
}

func (o AutoscaleCapacityOutput) MaxInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoscaleCapacity) *int { return v.MaxInstanceCount }).(pulumi.IntPtrOutput)
}

func (o AutoscaleCapacityOutput) MinInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoscaleCapacity) *int { return v.MinInstanceCount }).(pulumi.IntPtrOutput)
}

type AutoscaleCapacityPtrOutput struct{ *pulumi.OutputState }

func (AutoscaleCapacityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleCapacity)(nil)).Elem()
}

func (o AutoscaleCapacityPtrOutput) ToAutoscaleCapacityPtrOutput() AutoscaleCapacityPtrOutput {
	return o
}

func (o AutoscaleCapacityPtrOutput) ToAutoscaleCapacityPtrOutputWithContext(ctx context.Context) AutoscaleCapacityPtrOutput {
	return o
}

func (o AutoscaleCapacityPtrOutput) Elem() AutoscaleCapacityOutput {
	return o.ApplyT(func(v *AutoscaleCapacity) AutoscaleCapacity {
		if v != nil {
			return *v
		}
		var ret AutoscaleCapacity
		return ret
	}).(AutoscaleCapacityOutput)
}

func (o AutoscaleCapacityPtrOutput) MaxInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoscaleCapacity) *int {
		if v == nil {
			return nil
		}
		return v.MaxInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o AutoscaleCapacityPtrOutput) MinInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoscaleCapacity) *int {
		if v == nil {
			return nil
		}
		return v.MinInstanceCount
	}).(pulumi.IntPtrOutput)
}

type AutoscaleCapacityResponse struct {
	MaxInstanceCount *int `pulumi:"maxInstanceCount"`
	MinInstanceCount *int `pulumi:"minInstanceCount"`
}

type AutoscaleCapacityResponseOutput struct{ *pulumi.OutputState }

func (AutoscaleCapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleCapacityResponse)(nil)).Elem()
}

func (o AutoscaleCapacityResponseOutput) ToAutoscaleCapacityResponseOutput() AutoscaleCapacityResponseOutput {
	return o
}

func (o AutoscaleCapacityResponseOutput) ToAutoscaleCapacityResponseOutputWithContext(ctx context.Context) AutoscaleCapacityResponseOutput {
	return o
}

func (o AutoscaleCapacityResponseOutput) MaxInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoscaleCapacityResponse) *int { return v.MaxInstanceCount }).(pulumi.IntPtrOutput)
}

func (o AutoscaleCapacityResponseOutput) MinInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoscaleCapacityResponse) *int { return v.MinInstanceCount }).(pulumi.IntPtrOutput)
}

type AutoscaleCapacityResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoscaleCapacityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleCapacityResponse)(nil)).Elem()
}

func (o AutoscaleCapacityResponsePtrOutput) ToAutoscaleCapacityResponsePtrOutput() AutoscaleCapacityResponsePtrOutput {
	return o
}

func (o AutoscaleCapacityResponsePtrOutput) ToAutoscaleCapacityResponsePtrOutputWithContext(ctx context.Context) AutoscaleCapacityResponsePtrOutput {
	return o
}

func (o AutoscaleCapacityResponsePtrOutput) Elem() AutoscaleCapacityResponseOutput {
	return o.ApplyT(func(v *AutoscaleCapacityResponse) AutoscaleCapacityResponse {
		if v != nil {
			return *v
		}
		var ret AutoscaleCapacityResponse
		return ret
	}).(AutoscaleCapacityResponseOutput)
}

func (o AutoscaleCapacityResponsePtrOutput) MaxInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoscaleCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o AutoscaleCapacityResponsePtrOutput) MinInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoscaleCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinInstanceCount
	}).(pulumi.IntPtrOutput)
}

type AutoscaleRecurrence struct {
	Schedule []AutoscaleSchedule `pulumi:"schedule"`
	TimeZone *string             `pulumi:"timeZone"`
}





type AutoscaleRecurrenceInput interface {
	pulumi.Input

	ToAutoscaleRecurrenceOutput() AutoscaleRecurrenceOutput
	ToAutoscaleRecurrenceOutputWithContext(context.Context) AutoscaleRecurrenceOutput
}

type AutoscaleRecurrenceArgs struct {
	Schedule AutoscaleScheduleArrayInput `pulumi:"schedule"`
	TimeZone pulumi.StringPtrInput       `pulumi:"timeZone"`
}

func (AutoscaleRecurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleRecurrence)(nil)).Elem()
}

func (i AutoscaleRecurrenceArgs) ToAutoscaleRecurrenceOutput() AutoscaleRecurrenceOutput {
	return i.ToAutoscaleRecurrenceOutputWithContext(context.Background())
}

func (i AutoscaleRecurrenceArgs) ToAutoscaleRecurrenceOutputWithContext(ctx context.Context) AutoscaleRecurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleRecurrenceOutput)
}

func (i AutoscaleRecurrenceArgs) ToAutoscaleRecurrencePtrOutput() AutoscaleRecurrencePtrOutput {
	return i.ToAutoscaleRecurrencePtrOutputWithContext(context.Background())
}

func (i AutoscaleRecurrenceArgs) ToAutoscaleRecurrencePtrOutputWithContext(ctx context.Context) AutoscaleRecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleRecurrenceOutput).ToAutoscaleRecurrencePtrOutputWithContext(ctx)
}









type AutoscaleRecurrencePtrInput interface {
	pulumi.Input

	ToAutoscaleRecurrencePtrOutput() AutoscaleRecurrencePtrOutput
	ToAutoscaleRecurrencePtrOutputWithContext(context.Context) AutoscaleRecurrencePtrOutput
}

type autoscaleRecurrencePtrType AutoscaleRecurrenceArgs

func AutoscaleRecurrencePtr(v *AutoscaleRecurrenceArgs) AutoscaleRecurrencePtrInput {
	return (*autoscaleRecurrencePtrType)(v)
}

func (*autoscaleRecurrencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleRecurrence)(nil)).Elem()
}

func (i *autoscaleRecurrencePtrType) ToAutoscaleRecurrencePtrOutput() AutoscaleRecurrencePtrOutput {
	return i.ToAutoscaleRecurrencePtrOutputWithContext(context.Background())
}

func (i *autoscaleRecurrencePtrType) ToAutoscaleRecurrencePtrOutputWithContext(ctx context.Context) AutoscaleRecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleRecurrencePtrOutput)
}

type AutoscaleRecurrenceOutput struct{ *pulumi.OutputState }

func (AutoscaleRecurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleRecurrence)(nil)).Elem()
}

func (o AutoscaleRecurrenceOutput) ToAutoscaleRecurrenceOutput() AutoscaleRecurrenceOutput {
	return o
}

func (o AutoscaleRecurrenceOutput) ToAutoscaleRecurrenceOutputWithContext(ctx context.Context) AutoscaleRecurrenceOutput {
	return o
}

func (o AutoscaleRecurrenceOutput) ToAutoscaleRecurrencePtrOutput() AutoscaleRecurrencePtrOutput {
	return o.ToAutoscaleRecurrencePtrOutputWithContext(context.Background())
}

func (o AutoscaleRecurrenceOutput) ToAutoscaleRecurrencePtrOutputWithContext(ctx context.Context) AutoscaleRecurrencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoscaleRecurrence) *AutoscaleRecurrence {
		return &v
	}).(AutoscaleRecurrencePtrOutput)
}

func (o AutoscaleRecurrenceOutput) Schedule() AutoscaleScheduleArrayOutput {
	return o.ApplyT(func(v AutoscaleRecurrence) []AutoscaleSchedule { return v.Schedule }).(AutoscaleScheduleArrayOutput)
}

func (o AutoscaleRecurrenceOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoscaleRecurrence) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type AutoscaleRecurrencePtrOutput struct{ *pulumi.OutputState }

func (AutoscaleRecurrencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleRecurrence)(nil)).Elem()
}

func (o AutoscaleRecurrencePtrOutput) ToAutoscaleRecurrencePtrOutput() AutoscaleRecurrencePtrOutput {
	return o
}

func (o AutoscaleRecurrencePtrOutput) ToAutoscaleRecurrencePtrOutputWithContext(ctx context.Context) AutoscaleRecurrencePtrOutput {
	return o
}

func (o AutoscaleRecurrencePtrOutput) Elem() AutoscaleRecurrenceOutput {
	return o.ApplyT(func(v *AutoscaleRecurrence) AutoscaleRecurrence {
		if v != nil {
			return *v
		}
		var ret AutoscaleRecurrence
		return ret
	}).(AutoscaleRecurrenceOutput)
}

func (o AutoscaleRecurrencePtrOutput) Schedule() AutoscaleScheduleArrayOutput {
	return o.ApplyT(func(v *AutoscaleRecurrence) []AutoscaleSchedule {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(AutoscaleScheduleArrayOutput)
}

func (o AutoscaleRecurrencePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscaleRecurrence) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type AutoscaleRecurrenceResponse struct {
	Schedule []AutoscaleScheduleResponse `pulumi:"schedule"`
	TimeZone *string                     `pulumi:"timeZone"`
}

type AutoscaleRecurrenceResponseOutput struct{ *pulumi.OutputState }

func (AutoscaleRecurrenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleRecurrenceResponse)(nil)).Elem()
}

func (o AutoscaleRecurrenceResponseOutput) ToAutoscaleRecurrenceResponseOutput() AutoscaleRecurrenceResponseOutput {
	return o
}

func (o AutoscaleRecurrenceResponseOutput) ToAutoscaleRecurrenceResponseOutputWithContext(ctx context.Context) AutoscaleRecurrenceResponseOutput {
	return o
}

func (o AutoscaleRecurrenceResponseOutput) Schedule() AutoscaleScheduleResponseArrayOutput {
	return o.ApplyT(func(v AutoscaleRecurrenceResponse) []AutoscaleScheduleResponse { return v.Schedule }).(AutoscaleScheduleResponseArrayOutput)
}

func (o AutoscaleRecurrenceResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoscaleRecurrenceResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type AutoscaleRecurrenceResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoscaleRecurrenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleRecurrenceResponse)(nil)).Elem()
}

func (o AutoscaleRecurrenceResponsePtrOutput) ToAutoscaleRecurrenceResponsePtrOutput() AutoscaleRecurrenceResponsePtrOutput {
	return o
}

func (o AutoscaleRecurrenceResponsePtrOutput) ToAutoscaleRecurrenceResponsePtrOutputWithContext(ctx context.Context) AutoscaleRecurrenceResponsePtrOutput {
	return o
}

func (o AutoscaleRecurrenceResponsePtrOutput) Elem() AutoscaleRecurrenceResponseOutput {
	return o.ApplyT(func(v *AutoscaleRecurrenceResponse) AutoscaleRecurrenceResponse {
		if v != nil {
			return *v
		}
		var ret AutoscaleRecurrenceResponse
		return ret
	}).(AutoscaleRecurrenceResponseOutput)
}

func (o AutoscaleRecurrenceResponsePtrOutput) Schedule() AutoscaleScheduleResponseArrayOutput {
	return o.ApplyT(func(v *AutoscaleRecurrenceResponse) []AutoscaleScheduleResponse {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(AutoscaleScheduleResponseArrayOutput)
}

func (o AutoscaleRecurrenceResponsePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscaleRecurrenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type AutoscaleResponse struct {
	Capacity   *AutoscaleCapacityResponse   `pulumi:"capacity"`
	Recurrence *AutoscaleRecurrenceResponse `pulumi:"recurrence"`
}

type AutoscaleResponseOutput struct{ *pulumi.OutputState }

func (AutoscaleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleResponse)(nil)).Elem()
}

func (o AutoscaleResponseOutput) ToAutoscaleResponseOutput() AutoscaleResponseOutput {
	return o
}

func (o AutoscaleResponseOutput) ToAutoscaleResponseOutputWithContext(ctx context.Context) AutoscaleResponseOutput {
	return o
}

func (o AutoscaleResponseOutput) Capacity() AutoscaleCapacityResponsePtrOutput {
	return o.ApplyT(func(v AutoscaleResponse) *AutoscaleCapacityResponse { return v.Capacity }).(AutoscaleCapacityResponsePtrOutput)
}

func (o AutoscaleResponseOutput) Recurrence() AutoscaleRecurrenceResponsePtrOutput {
	return o.ApplyT(func(v AutoscaleResponse) *AutoscaleRecurrenceResponse { return v.Recurrence }).(AutoscaleRecurrenceResponsePtrOutput)
}

type AutoscaleResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoscaleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleResponse)(nil)).Elem()
}

func (o AutoscaleResponsePtrOutput) ToAutoscaleResponsePtrOutput() AutoscaleResponsePtrOutput {
	return o
}

func (o AutoscaleResponsePtrOutput) ToAutoscaleResponsePtrOutputWithContext(ctx context.Context) AutoscaleResponsePtrOutput {
	return o
}

func (o AutoscaleResponsePtrOutput) Elem() AutoscaleResponseOutput {
	return o.ApplyT(func(v *AutoscaleResponse) AutoscaleResponse {
		if v != nil {
			return *v
		}
		var ret AutoscaleResponse
		return ret
	}).(AutoscaleResponseOutput)
}

func (o AutoscaleResponsePtrOutput) Capacity() AutoscaleCapacityResponsePtrOutput {
	return o.ApplyT(func(v *AutoscaleResponse) *AutoscaleCapacityResponse {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(AutoscaleCapacityResponsePtrOutput)
}

func (o AutoscaleResponsePtrOutput) Recurrence() AutoscaleRecurrenceResponsePtrOutput {
	return o.ApplyT(func(v *AutoscaleResponse) *AutoscaleRecurrenceResponse {
		if v == nil {
			return nil
		}
		return v.Recurrence
	}).(AutoscaleRecurrenceResponsePtrOutput)
}

type AutoscaleSchedule struct {
	Days            []string                  `pulumi:"days"`
	TimeAndCapacity *AutoscaleTimeAndCapacity `pulumi:"timeAndCapacity"`
}





type AutoscaleScheduleInput interface {
	pulumi.Input

	ToAutoscaleScheduleOutput() AutoscaleScheduleOutput
	ToAutoscaleScheduleOutputWithContext(context.Context) AutoscaleScheduleOutput
}

type AutoscaleScheduleArgs struct {
	Days            pulumi.StringArrayInput          `pulumi:"days"`
	TimeAndCapacity AutoscaleTimeAndCapacityPtrInput `pulumi:"timeAndCapacity"`
}

func (AutoscaleScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleSchedule)(nil)).Elem()
}

func (i AutoscaleScheduleArgs) ToAutoscaleScheduleOutput() AutoscaleScheduleOutput {
	return i.ToAutoscaleScheduleOutputWithContext(context.Background())
}

func (i AutoscaleScheduleArgs) ToAutoscaleScheduleOutputWithContext(ctx context.Context) AutoscaleScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleScheduleOutput)
}





type AutoscaleScheduleArrayInput interface {
	pulumi.Input

	ToAutoscaleScheduleArrayOutput() AutoscaleScheduleArrayOutput
	ToAutoscaleScheduleArrayOutputWithContext(context.Context) AutoscaleScheduleArrayOutput
}

type AutoscaleScheduleArray []AutoscaleScheduleInput

func (AutoscaleScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoscaleSchedule)(nil)).Elem()
}

func (i AutoscaleScheduleArray) ToAutoscaleScheduleArrayOutput() AutoscaleScheduleArrayOutput {
	return i.ToAutoscaleScheduleArrayOutputWithContext(context.Background())
}

func (i AutoscaleScheduleArray) ToAutoscaleScheduleArrayOutputWithContext(ctx context.Context) AutoscaleScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleScheduleArrayOutput)
}

type AutoscaleScheduleOutput struct{ *pulumi.OutputState }

func (AutoscaleScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleSchedule)(nil)).Elem()
}

func (o AutoscaleScheduleOutput) ToAutoscaleScheduleOutput() AutoscaleScheduleOutput {
	return o
}

func (o AutoscaleScheduleOutput) ToAutoscaleScheduleOutputWithContext(ctx context.Context) AutoscaleScheduleOutput {
	return o
}

func (o AutoscaleScheduleOutput) Days() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AutoscaleSchedule) []string { return v.Days }).(pulumi.StringArrayOutput)
}

func (o AutoscaleScheduleOutput) TimeAndCapacity() AutoscaleTimeAndCapacityPtrOutput {
	return o.ApplyT(func(v AutoscaleSchedule) *AutoscaleTimeAndCapacity { return v.TimeAndCapacity }).(AutoscaleTimeAndCapacityPtrOutput)
}

type AutoscaleScheduleArrayOutput struct{ *pulumi.OutputState }

func (AutoscaleScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoscaleSchedule)(nil)).Elem()
}

func (o AutoscaleScheduleArrayOutput) ToAutoscaleScheduleArrayOutput() AutoscaleScheduleArrayOutput {
	return o
}

func (o AutoscaleScheduleArrayOutput) ToAutoscaleScheduleArrayOutputWithContext(ctx context.Context) AutoscaleScheduleArrayOutput {
	return o
}

func (o AutoscaleScheduleArrayOutput) Index(i pulumi.IntInput) AutoscaleScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutoscaleSchedule {
		return vs[0].([]AutoscaleSchedule)[vs[1].(int)]
	}).(AutoscaleScheduleOutput)
}

type AutoscaleScheduleResponse struct {
	Days            []string                          `pulumi:"days"`
	TimeAndCapacity *AutoscaleTimeAndCapacityResponse `pulumi:"timeAndCapacity"`
}

type AutoscaleScheduleResponseOutput struct{ *pulumi.OutputState }

func (AutoscaleScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleScheduleResponse)(nil)).Elem()
}

func (o AutoscaleScheduleResponseOutput) ToAutoscaleScheduleResponseOutput() AutoscaleScheduleResponseOutput {
	return o
}

func (o AutoscaleScheduleResponseOutput) ToAutoscaleScheduleResponseOutputWithContext(ctx context.Context) AutoscaleScheduleResponseOutput {
	return o
}

func (o AutoscaleScheduleResponseOutput) Days() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AutoscaleScheduleResponse) []string { return v.Days }).(pulumi.StringArrayOutput)
}

func (o AutoscaleScheduleResponseOutput) TimeAndCapacity() AutoscaleTimeAndCapacityResponsePtrOutput {
	return o.ApplyT(func(v AutoscaleScheduleResponse) *AutoscaleTimeAndCapacityResponse { return v.TimeAndCapacity }).(AutoscaleTimeAndCapacityResponsePtrOutput)
}

type AutoscaleScheduleResponseArrayOutput struct{ *pulumi.OutputState }

func (AutoscaleScheduleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoscaleScheduleResponse)(nil)).Elem()
}

func (o AutoscaleScheduleResponseArrayOutput) ToAutoscaleScheduleResponseArrayOutput() AutoscaleScheduleResponseArrayOutput {
	return o
}

func (o AutoscaleScheduleResponseArrayOutput) ToAutoscaleScheduleResponseArrayOutputWithContext(ctx context.Context) AutoscaleScheduleResponseArrayOutput {
	return o
}

func (o AutoscaleScheduleResponseArrayOutput) Index(i pulumi.IntInput) AutoscaleScheduleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutoscaleScheduleResponse {
		return vs[0].([]AutoscaleScheduleResponse)[vs[1].(int)]
	}).(AutoscaleScheduleResponseOutput)
}

type AutoscaleTimeAndCapacity struct {
	MaxInstanceCount *int    `pulumi:"maxInstanceCount"`
	MinInstanceCount *int    `pulumi:"minInstanceCount"`
	Time             *string `pulumi:"time"`
}





type AutoscaleTimeAndCapacityInput interface {
	pulumi.Input

	ToAutoscaleTimeAndCapacityOutput() AutoscaleTimeAndCapacityOutput
	ToAutoscaleTimeAndCapacityOutputWithContext(context.Context) AutoscaleTimeAndCapacityOutput
}

type AutoscaleTimeAndCapacityArgs struct {
	MaxInstanceCount pulumi.IntPtrInput    `pulumi:"maxInstanceCount"`
	MinInstanceCount pulumi.IntPtrInput    `pulumi:"minInstanceCount"`
	Time             pulumi.StringPtrInput `pulumi:"time"`
}

func (AutoscaleTimeAndCapacityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleTimeAndCapacity)(nil)).Elem()
}

func (i AutoscaleTimeAndCapacityArgs) ToAutoscaleTimeAndCapacityOutput() AutoscaleTimeAndCapacityOutput {
	return i.ToAutoscaleTimeAndCapacityOutputWithContext(context.Background())
}

func (i AutoscaleTimeAndCapacityArgs) ToAutoscaleTimeAndCapacityOutputWithContext(ctx context.Context) AutoscaleTimeAndCapacityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleTimeAndCapacityOutput)
}

func (i AutoscaleTimeAndCapacityArgs) ToAutoscaleTimeAndCapacityPtrOutput() AutoscaleTimeAndCapacityPtrOutput {
	return i.ToAutoscaleTimeAndCapacityPtrOutputWithContext(context.Background())
}

func (i AutoscaleTimeAndCapacityArgs) ToAutoscaleTimeAndCapacityPtrOutputWithContext(ctx context.Context) AutoscaleTimeAndCapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleTimeAndCapacityOutput).ToAutoscaleTimeAndCapacityPtrOutputWithContext(ctx)
}









type AutoscaleTimeAndCapacityPtrInput interface {
	pulumi.Input

	ToAutoscaleTimeAndCapacityPtrOutput() AutoscaleTimeAndCapacityPtrOutput
	ToAutoscaleTimeAndCapacityPtrOutputWithContext(context.Context) AutoscaleTimeAndCapacityPtrOutput
}

type autoscaleTimeAndCapacityPtrType AutoscaleTimeAndCapacityArgs

func AutoscaleTimeAndCapacityPtr(v *AutoscaleTimeAndCapacityArgs) AutoscaleTimeAndCapacityPtrInput {
	return (*autoscaleTimeAndCapacityPtrType)(v)
}

func (*autoscaleTimeAndCapacityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleTimeAndCapacity)(nil)).Elem()
}

func (i *autoscaleTimeAndCapacityPtrType) ToAutoscaleTimeAndCapacityPtrOutput() AutoscaleTimeAndCapacityPtrOutput {
	return i.ToAutoscaleTimeAndCapacityPtrOutputWithContext(context.Background())
}

func (i *autoscaleTimeAndCapacityPtrType) ToAutoscaleTimeAndCapacityPtrOutputWithContext(ctx context.Context) AutoscaleTimeAndCapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleTimeAndCapacityPtrOutput)
}

type AutoscaleTimeAndCapacityOutput struct{ *pulumi.OutputState }

func (AutoscaleTimeAndCapacityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleTimeAndCapacity)(nil)).Elem()
}

func (o AutoscaleTimeAndCapacityOutput) ToAutoscaleTimeAndCapacityOutput() AutoscaleTimeAndCapacityOutput {
	return o
}

func (o AutoscaleTimeAndCapacityOutput) ToAutoscaleTimeAndCapacityOutputWithContext(ctx context.Context) AutoscaleTimeAndCapacityOutput {
	return o
}

func (o AutoscaleTimeAndCapacityOutput) ToAutoscaleTimeAndCapacityPtrOutput() AutoscaleTimeAndCapacityPtrOutput {
	return o.ToAutoscaleTimeAndCapacityPtrOutputWithContext(context.Background())
}

func (o AutoscaleTimeAndCapacityOutput) ToAutoscaleTimeAndCapacityPtrOutputWithContext(ctx context.Context) AutoscaleTimeAndCapacityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoscaleTimeAndCapacity) *AutoscaleTimeAndCapacity {
		return &v
	}).(AutoscaleTimeAndCapacityPtrOutput)
}

func (o AutoscaleTimeAndCapacityOutput) MaxInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoscaleTimeAndCapacity) *int { return v.MaxInstanceCount }).(pulumi.IntPtrOutput)
}

func (o AutoscaleTimeAndCapacityOutput) MinInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoscaleTimeAndCapacity) *int { return v.MinInstanceCount }).(pulumi.IntPtrOutput)
}

func (o AutoscaleTimeAndCapacityOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoscaleTimeAndCapacity) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type AutoscaleTimeAndCapacityPtrOutput struct{ *pulumi.OutputState }

func (AutoscaleTimeAndCapacityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleTimeAndCapacity)(nil)).Elem()
}

func (o AutoscaleTimeAndCapacityPtrOutput) ToAutoscaleTimeAndCapacityPtrOutput() AutoscaleTimeAndCapacityPtrOutput {
	return o
}

func (o AutoscaleTimeAndCapacityPtrOutput) ToAutoscaleTimeAndCapacityPtrOutputWithContext(ctx context.Context) AutoscaleTimeAndCapacityPtrOutput {
	return o
}

func (o AutoscaleTimeAndCapacityPtrOutput) Elem() AutoscaleTimeAndCapacityOutput {
	return o.ApplyT(func(v *AutoscaleTimeAndCapacity) AutoscaleTimeAndCapacity {
		if v != nil {
			return *v
		}
		var ret AutoscaleTimeAndCapacity
		return ret
	}).(AutoscaleTimeAndCapacityOutput)
}

func (o AutoscaleTimeAndCapacityPtrOutput) MaxInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoscaleTimeAndCapacity) *int {
		if v == nil {
			return nil
		}
		return v.MaxInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o AutoscaleTimeAndCapacityPtrOutput) MinInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoscaleTimeAndCapacity) *int {
		if v == nil {
			return nil
		}
		return v.MinInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o AutoscaleTimeAndCapacityPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscaleTimeAndCapacity) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type AutoscaleTimeAndCapacityResponse struct {
	MaxInstanceCount *int    `pulumi:"maxInstanceCount"`
	MinInstanceCount *int    `pulumi:"minInstanceCount"`
	Time             *string `pulumi:"time"`
}

type AutoscaleTimeAndCapacityResponseOutput struct{ *pulumi.OutputState }

func (AutoscaleTimeAndCapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleTimeAndCapacityResponse)(nil)).Elem()
}

func (o AutoscaleTimeAndCapacityResponseOutput) ToAutoscaleTimeAndCapacityResponseOutput() AutoscaleTimeAndCapacityResponseOutput {
	return o
}

func (o AutoscaleTimeAndCapacityResponseOutput) ToAutoscaleTimeAndCapacityResponseOutputWithContext(ctx context.Context) AutoscaleTimeAndCapacityResponseOutput {
	return o
}

func (o AutoscaleTimeAndCapacityResponseOutput) MaxInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoscaleTimeAndCapacityResponse) *int { return v.MaxInstanceCount }).(pulumi.IntPtrOutput)
}

func (o AutoscaleTimeAndCapacityResponseOutput) MinInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoscaleTimeAndCapacityResponse) *int { return v.MinInstanceCount }).(pulumi.IntPtrOutput)
}

func (o AutoscaleTimeAndCapacityResponseOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoscaleTimeAndCapacityResponse) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type AutoscaleTimeAndCapacityResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoscaleTimeAndCapacityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleTimeAndCapacityResponse)(nil)).Elem()
}

func (o AutoscaleTimeAndCapacityResponsePtrOutput) ToAutoscaleTimeAndCapacityResponsePtrOutput() AutoscaleTimeAndCapacityResponsePtrOutput {
	return o
}

func (o AutoscaleTimeAndCapacityResponsePtrOutput) ToAutoscaleTimeAndCapacityResponsePtrOutputWithContext(ctx context.Context) AutoscaleTimeAndCapacityResponsePtrOutput {
	return o
}

func (o AutoscaleTimeAndCapacityResponsePtrOutput) Elem() AutoscaleTimeAndCapacityResponseOutput {
	return o.ApplyT(func(v *AutoscaleTimeAndCapacityResponse) AutoscaleTimeAndCapacityResponse {
		if v != nil {
			return *v
		}
		var ret AutoscaleTimeAndCapacityResponse
		return ret
	}).(AutoscaleTimeAndCapacityResponseOutput)
}

func (o AutoscaleTimeAndCapacityResponsePtrOutput) MaxInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoscaleTimeAndCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o AutoscaleTimeAndCapacityResponsePtrOutput) MinInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoscaleTimeAndCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o AutoscaleTimeAndCapacityResponsePtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoscaleTimeAndCapacityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type ClientGroupInfo struct {
	GroupId   *string `pulumi:"groupId"`
	GroupName *string `pulumi:"groupName"`
}





type ClientGroupInfoInput interface {
	pulumi.Input

	ToClientGroupInfoOutput() ClientGroupInfoOutput
	ToClientGroupInfoOutputWithContext(context.Context) ClientGroupInfoOutput
}

type ClientGroupInfoArgs struct {
	GroupId   pulumi.StringPtrInput `pulumi:"groupId"`
	GroupName pulumi.StringPtrInput `pulumi:"groupName"`
}

func (ClientGroupInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientGroupInfo)(nil)).Elem()
}

func (i ClientGroupInfoArgs) ToClientGroupInfoOutput() ClientGroupInfoOutput {
	return i.ToClientGroupInfoOutputWithContext(context.Background())
}

func (i ClientGroupInfoArgs) ToClientGroupInfoOutputWithContext(ctx context.Context) ClientGroupInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGroupInfoOutput)
}

func (i ClientGroupInfoArgs) ToClientGroupInfoPtrOutput() ClientGroupInfoPtrOutput {
	return i.ToClientGroupInfoPtrOutputWithContext(context.Background())
}

func (i ClientGroupInfoArgs) ToClientGroupInfoPtrOutputWithContext(ctx context.Context) ClientGroupInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGroupInfoOutput).ToClientGroupInfoPtrOutputWithContext(ctx)
}









type ClientGroupInfoPtrInput interface {
	pulumi.Input

	ToClientGroupInfoPtrOutput() ClientGroupInfoPtrOutput
	ToClientGroupInfoPtrOutputWithContext(context.Context) ClientGroupInfoPtrOutput
}

type clientGroupInfoPtrType ClientGroupInfoArgs

func ClientGroupInfoPtr(v *ClientGroupInfoArgs) ClientGroupInfoPtrInput {
	return (*clientGroupInfoPtrType)(v)
}

func (*clientGroupInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientGroupInfo)(nil)).Elem()
}

func (i *clientGroupInfoPtrType) ToClientGroupInfoPtrOutput() ClientGroupInfoPtrOutput {
	return i.ToClientGroupInfoPtrOutputWithContext(context.Background())
}

func (i *clientGroupInfoPtrType) ToClientGroupInfoPtrOutputWithContext(ctx context.Context) ClientGroupInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGroupInfoPtrOutput)
}

type ClientGroupInfoOutput struct{ *pulumi.OutputState }

func (ClientGroupInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientGroupInfo)(nil)).Elem()
}

func (o ClientGroupInfoOutput) ToClientGroupInfoOutput() ClientGroupInfoOutput {
	return o
}

func (o ClientGroupInfoOutput) ToClientGroupInfoOutputWithContext(ctx context.Context) ClientGroupInfoOutput {
	return o
}

func (o ClientGroupInfoOutput) ToClientGroupInfoPtrOutput() ClientGroupInfoPtrOutput {
	return o.ToClientGroupInfoPtrOutputWithContext(context.Background())
}

func (o ClientGroupInfoOutput) ToClientGroupInfoPtrOutputWithContext(ctx context.Context) ClientGroupInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClientGroupInfo) *ClientGroupInfo {
		return &v
	}).(ClientGroupInfoPtrOutput)
}

func (o ClientGroupInfoOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientGroupInfo) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o ClientGroupInfoOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientGroupInfo) *string { return v.GroupName }).(pulumi.StringPtrOutput)
}

type ClientGroupInfoPtrOutput struct{ *pulumi.OutputState }

func (ClientGroupInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientGroupInfo)(nil)).Elem()
}

func (o ClientGroupInfoPtrOutput) ToClientGroupInfoPtrOutput() ClientGroupInfoPtrOutput {
	return o
}

func (o ClientGroupInfoPtrOutput) ToClientGroupInfoPtrOutputWithContext(ctx context.Context) ClientGroupInfoPtrOutput {
	return o
}

func (o ClientGroupInfoPtrOutput) Elem() ClientGroupInfoOutput {
	return o.ApplyT(func(v *ClientGroupInfo) ClientGroupInfo {
		if v != nil {
			return *v
		}
		var ret ClientGroupInfo
		return ret
	}).(ClientGroupInfoOutput)
}

func (o ClientGroupInfoPtrOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientGroupInfo) *string {
		if v == nil {
			return nil
		}
		return v.GroupId
	}).(pulumi.StringPtrOutput)
}

func (o ClientGroupInfoPtrOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientGroupInfo) *string {
		if v == nil {
			return nil
		}
		return v.GroupName
	}).(pulumi.StringPtrOutput)
}

type ClientGroupInfoResponse struct {
	GroupId   *string `pulumi:"groupId"`
	GroupName *string `pulumi:"groupName"`
}

type ClientGroupInfoResponseOutput struct{ *pulumi.OutputState }

func (ClientGroupInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientGroupInfoResponse)(nil)).Elem()
}

func (o ClientGroupInfoResponseOutput) ToClientGroupInfoResponseOutput() ClientGroupInfoResponseOutput {
	return o
}

func (o ClientGroupInfoResponseOutput) ToClientGroupInfoResponseOutputWithContext(ctx context.Context) ClientGroupInfoResponseOutput {
	return o
}

func (o ClientGroupInfoResponseOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientGroupInfoResponse) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o ClientGroupInfoResponseOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientGroupInfoResponse) *string { return v.GroupName }).(pulumi.StringPtrOutput)
}

type ClientGroupInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ClientGroupInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientGroupInfoResponse)(nil)).Elem()
}

func (o ClientGroupInfoResponsePtrOutput) ToClientGroupInfoResponsePtrOutput() ClientGroupInfoResponsePtrOutput {
	return o
}

func (o ClientGroupInfoResponsePtrOutput) ToClientGroupInfoResponsePtrOutputWithContext(ctx context.Context) ClientGroupInfoResponsePtrOutput {
	return o
}

func (o ClientGroupInfoResponsePtrOutput) Elem() ClientGroupInfoResponseOutput {
	return o.ApplyT(func(v *ClientGroupInfoResponse) ClientGroupInfoResponse {
		if v != nil {
			return *v
		}
		var ret ClientGroupInfoResponse
		return ret
	}).(ClientGroupInfoResponseOutput)
}

func (o ClientGroupInfoResponsePtrOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientGroupInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupId
	}).(pulumi.StringPtrOutput)
}

func (o ClientGroupInfoResponsePtrOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientGroupInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupName
	}).(pulumi.StringPtrOutput)
}

type ClusterCreateProperties struct {
	ClusterDefinition             *ClusterDefinition             `pulumi:"clusterDefinition"`
	ClusterVersion                *string                        `pulumi:"clusterVersion"`
	ComputeIsolationProperties    *ComputeIsolationProperties    `pulumi:"computeIsolationProperties"`
	ComputeProfile                *ComputeProfile                `pulumi:"computeProfile"`
	DiskEncryptionProperties      *DiskEncryptionProperties      `pulumi:"diskEncryptionProperties"`
	EncryptionInTransitProperties *EncryptionInTransitProperties `pulumi:"encryptionInTransitProperties"`
	KafkaRestProperties           *KafkaRestProperties           `pulumi:"kafkaRestProperties"`
	MinSupportedTlsVersion        *string                        `pulumi:"minSupportedTlsVersion"`
	NetworkProperties             *NetworkProperties             `pulumi:"networkProperties"`
	OsType                        *string                        `pulumi:"osType"`
	PrivateLinkConfigurations     []PrivateLinkConfiguration     `pulumi:"privateLinkConfigurations"`
	SecurityProfile               *SecurityProfile               `pulumi:"securityProfile"`
	StorageProfile                *StorageProfile                `pulumi:"storageProfile"`
	Tier                          *string                        `pulumi:"tier"`
}


func (val *ClusterCreateProperties) Defaults() *ClusterCreateProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ComputeIsolationProperties = tmp.ComputeIsolationProperties.Defaults()

	tmp.DiskEncryptionProperties = tmp.DiskEncryptionProperties.Defaults()

	tmp.EncryptionInTransitProperties = tmp.EncryptionInTransitProperties.Defaults()

	if isZero(tmp.Tier) {
		tier_ := "Standard"
		tmp.Tier = &tier_
	}
	return &tmp
}





type ClusterCreatePropertiesInput interface {
	pulumi.Input

	ToClusterCreatePropertiesOutput() ClusterCreatePropertiesOutput
	ToClusterCreatePropertiesOutputWithContext(context.Context) ClusterCreatePropertiesOutput
}

type ClusterCreatePropertiesArgs struct {
	ClusterDefinition             ClusterDefinitionPtrInput             `pulumi:"clusterDefinition"`
	ClusterVersion                pulumi.StringPtrInput                 `pulumi:"clusterVersion"`
	ComputeIsolationProperties    ComputeIsolationPropertiesPtrInput    `pulumi:"computeIsolationProperties"`
	ComputeProfile                ComputeProfilePtrInput                `pulumi:"computeProfile"`
	DiskEncryptionProperties      DiskEncryptionPropertiesPtrInput      `pulumi:"diskEncryptionProperties"`
	EncryptionInTransitProperties EncryptionInTransitPropertiesPtrInput `pulumi:"encryptionInTransitProperties"`
	KafkaRestProperties           KafkaRestPropertiesPtrInput           `pulumi:"kafkaRestProperties"`
	MinSupportedTlsVersion        pulumi.StringPtrInput                 `pulumi:"minSupportedTlsVersion"`
	NetworkProperties             NetworkPropertiesPtrInput             `pulumi:"networkProperties"`
	OsType                        pulumi.StringPtrInput                 `pulumi:"osType"`
	PrivateLinkConfigurations     PrivateLinkConfigurationArrayInput    `pulumi:"privateLinkConfigurations"`
	SecurityProfile               SecurityProfilePtrInput               `pulumi:"securityProfile"`
	StorageProfile                StorageProfilePtrInput                `pulumi:"storageProfile"`
	Tier                          pulumi.StringPtrInput                 `pulumi:"tier"`
}


func (val *ClusterCreatePropertiesArgs) Defaults() *ClusterCreatePropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	if isZero(tmp.Tier) {
		tmp.Tier = pulumi.StringPtr("Standard")
	}
	return &tmp
}
func (ClusterCreatePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterCreateProperties)(nil)).Elem()
}

func (i ClusterCreatePropertiesArgs) ToClusterCreatePropertiesOutput() ClusterCreatePropertiesOutput {
	return i.ToClusterCreatePropertiesOutputWithContext(context.Background())
}

func (i ClusterCreatePropertiesArgs) ToClusterCreatePropertiesOutputWithContext(ctx context.Context) ClusterCreatePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCreatePropertiesOutput)
}

func (i ClusterCreatePropertiesArgs) ToClusterCreatePropertiesPtrOutput() ClusterCreatePropertiesPtrOutput {
	return i.ToClusterCreatePropertiesPtrOutputWithContext(context.Background())
}

func (i ClusterCreatePropertiesArgs) ToClusterCreatePropertiesPtrOutputWithContext(ctx context.Context) ClusterCreatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCreatePropertiesOutput).ToClusterCreatePropertiesPtrOutputWithContext(ctx)
}









type ClusterCreatePropertiesPtrInput interface {
	pulumi.Input

	ToClusterCreatePropertiesPtrOutput() ClusterCreatePropertiesPtrOutput
	ToClusterCreatePropertiesPtrOutputWithContext(context.Context) ClusterCreatePropertiesPtrOutput
}

type clusterCreatePropertiesPtrType ClusterCreatePropertiesArgs

func ClusterCreatePropertiesPtr(v *ClusterCreatePropertiesArgs) ClusterCreatePropertiesPtrInput {
	return (*clusterCreatePropertiesPtrType)(v)
}

func (*clusterCreatePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterCreateProperties)(nil)).Elem()
}

func (i *clusterCreatePropertiesPtrType) ToClusterCreatePropertiesPtrOutput() ClusterCreatePropertiesPtrOutput {
	return i.ToClusterCreatePropertiesPtrOutputWithContext(context.Background())
}

func (i *clusterCreatePropertiesPtrType) ToClusterCreatePropertiesPtrOutputWithContext(ctx context.Context) ClusterCreatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCreatePropertiesPtrOutput)
}

type ClusterCreatePropertiesOutput struct{ *pulumi.OutputState }

func (ClusterCreatePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterCreateProperties)(nil)).Elem()
}

func (o ClusterCreatePropertiesOutput) ToClusterCreatePropertiesOutput() ClusterCreatePropertiesOutput {
	return o
}

func (o ClusterCreatePropertiesOutput) ToClusterCreatePropertiesOutputWithContext(ctx context.Context) ClusterCreatePropertiesOutput {
	return o
}

func (o ClusterCreatePropertiesOutput) ToClusterCreatePropertiesPtrOutput() ClusterCreatePropertiesPtrOutput {
	return o.ToClusterCreatePropertiesPtrOutputWithContext(context.Background())
}

func (o ClusterCreatePropertiesOutput) ToClusterCreatePropertiesPtrOutputWithContext(ctx context.Context) ClusterCreatePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterCreateProperties) *ClusterCreateProperties {
		return &v
	}).(ClusterCreatePropertiesPtrOutput)
}

func (o ClusterCreatePropertiesOutput) ClusterDefinition() ClusterDefinitionPtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *ClusterDefinition { return v.ClusterDefinition }).(ClusterDefinitionPtrOutput)
}

func (o ClusterCreatePropertiesOutput) ClusterVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *string { return v.ClusterVersion }).(pulumi.StringPtrOutput)
}

func (o ClusterCreatePropertiesOutput) ComputeIsolationProperties() ComputeIsolationPropertiesPtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *ComputeIsolationProperties { return v.ComputeIsolationProperties }).(ComputeIsolationPropertiesPtrOutput)
}

func (o ClusterCreatePropertiesOutput) ComputeProfile() ComputeProfilePtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *ComputeProfile { return v.ComputeProfile }).(ComputeProfilePtrOutput)
}

func (o ClusterCreatePropertiesOutput) DiskEncryptionProperties() DiskEncryptionPropertiesPtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *DiskEncryptionProperties { return v.DiskEncryptionProperties }).(DiskEncryptionPropertiesPtrOutput)
}

func (o ClusterCreatePropertiesOutput) EncryptionInTransitProperties() EncryptionInTransitPropertiesPtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *EncryptionInTransitProperties { return v.EncryptionInTransitProperties }).(EncryptionInTransitPropertiesPtrOutput)
}

func (o ClusterCreatePropertiesOutput) KafkaRestProperties() KafkaRestPropertiesPtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *KafkaRestProperties { return v.KafkaRestProperties }).(KafkaRestPropertiesPtrOutput)
}

func (o ClusterCreatePropertiesOutput) MinSupportedTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *string { return v.MinSupportedTlsVersion }).(pulumi.StringPtrOutput)
}

func (o ClusterCreatePropertiesOutput) NetworkProperties() NetworkPropertiesPtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *NetworkProperties { return v.NetworkProperties }).(NetworkPropertiesPtrOutput)
}

func (o ClusterCreatePropertiesOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ClusterCreatePropertiesOutput) PrivateLinkConfigurations() PrivateLinkConfigurationArrayOutput {
	return o.ApplyT(func(v ClusterCreateProperties) []PrivateLinkConfiguration { return v.PrivateLinkConfigurations }).(PrivateLinkConfigurationArrayOutput)
}

func (o ClusterCreatePropertiesOutput) SecurityProfile() SecurityProfilePtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *SecurityProfile { return v.SecurityProfile }).(SecurityProfilePtrOutput)
}

func (o ClusterCreatePropertiesOutput) StorageProfile() StorageProfilePtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *StorageProfile { return v.StorageProfile }).(StorageProfilePtrOutput)
}

func (o ClusterCreatePropertiesOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterCreateProperties) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ClusterCreatePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ClusterCreatePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterCreateProperties)(nil)).Elem()
}

func (o ClusterCreatePropertiesPtrOutput) ToClusterCreatePropertiesPtrOutput() ClusterCreatePropertiesPtrOutput {
	return o
}

func (o ClusterCreatePropertiesPtrOutput) ToClusterCreatePropertiesPtrOutputWithContext(ctx context.Context) ClusterCreatePropertiesPtrOutput {
	return o
}

func (o ClusterCreatePropertiesPtrOutput) Elem() ClusterCreatePropertiesOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) ClusterCreateProperties {
		if v != nil {
			return *v
		}
		var ret ClusterCreateProperties
		return ret
	}).(ClusterCreatePropertiesOutput)
}

func (o ClusterCreatePropertiesPtrOutput) ClusterDefinition() ClusterDefinitionPtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *ClusterDefinition {
		if v == nil {
			return nil
		}
		return v.ClusterDefinition
	}).(ClusterDefinitionPtrOutput)
}

func (o ClusterCreatePropertiesPtrOutput) ClusterVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClusterVersion
	}).(pulumi.StringPtrOutput)
}

func (o ClusterCreatePropertiesPtrOutput) ComputeIsolationProperties() ComputeIsolationPropertiesPtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *ComputeIsolationProperties {
		if v == nil {
			return nil
		}
		return v.ComputeIsolationProperties
	}).(ComputeIsolationPropertiesPtrOutput)
}

func (o ClusterCreatePropertiesPtrOutput) ComputeProfile() ComputeProfilePtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *ComputeProfile {
		if v == nil {
			return nil
		}
		return v.ComputeProfile
	}).(ComputeProfilePtrOutput)
}

func (o ClusterCreatePropertiesPtrOutput) DiskEncryptionProperties() DiskEncryptionPropertiesPtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *DiskEncryptionProperties {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionProperties
	}).(DiskEncryptionPropertiesPtrOutput)
}

func (o ClusterCreatePropertiesPtrOutput) EncryptionInTransitProperties() EncryptionInTransitPropertiesPtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *EncryptionInTransitProperties {
		if v == nil {
			return nil
		}
		return v.EncryptionInTransitProperties
	}).(EncryptionInTransitPropertiesPtrOutput)
}

func (o ClusterCreatePropertiesPtrOutput) KafkaRestProperties() KafkaRestPropertiesPtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *KafkaRestProperties {
		if v == nil {
			return nil
		}
		return v.KafkaRestProperties
	}).(KafkaRestPropertiesPtrOutput)
}

func (o ClusterCreatePropertiesPtrOutput) MinSupportedTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *string {
		if v == nil {
			return nil
		}
		return v.MinSupportedTlsVersion
	}).(pulumi.StringPtrOutput)
}

func (o ClusterCreatePropertiesPtrOutput) NetworkProperties() NetworkPropertiesPtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *NetworkProperties {
		if v == nil {
			return nil
		}
		return v.NetworkProperties
	}).(NetworkPropertiesPtrOutput)
}

func (o ClusterCreatePropertiesPtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o ClusterCreatePropertiesPtrOutput) PrivateLinkConfigurations() PrivateLinkConfigurationArrayOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) []PrivateLinkConfiguration {
		if v == nil {
			return nil
		}
		return v.PrivateLinkConfigurations
	}).(PrivateLinkConfigurationArrayOutput)
}

func (o ClusterCreatePropertiesPtrOutput) SecurityProfile() SecurityProfilePtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *SecurityProfile {
		if v == nil {
			return nil
		}
		return v.SecurityProfile
	}).(SecurityProfilePtrOutput)
}

func (o ClusterCreatePropertiesPtrOutput) StorageProfile() StorageProfilePtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *StorageProfile {
		if v == nil {
			return nil
		}
		return v.StorageProfile
	}).(StorageProfilePtrOutput)
}

func (o ClusterCreatePropertiesPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterCreateProperties) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ClusterDefinition struct {
	Blueprint        *string           `pulumi:"blueprint"`
	ComponentVersion map[string]string `pulumi:"componentVersion"`
	Configurations   interface{}       `pulumi:"configurations"`
	Kind             *string           `pulumi:"kind"`
}





type ClusterDefinitionInput interface {
	pulumi.Input

	ToClusterDefinitionOutput() ClusterDefinitionOutput
	ToClusterDefinitionOutputWithContext(context.Context) ClusterDefinitionOutput
}

type ClusterDefinitionArgs struct {
	Blueprint        pulumi.StringPtrInput `pulumi:"blueprint"`
	ComponentVersion pulumi.StringMapInput `pulumi:"componentVersion"`
	Configurations   pulumi.Input          `pulumi:"configurations"`
	Kind             pulumi.StringPtrInput `pulumi:"kind"`
}

func (ClusterDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterDefinition)(nil)).Elem()
}

func (i ClusterDefinitionArgs) ToClusterDefinitionOutput() ClusterDefinitionOutput {
	return i.ToClusterDefinitionOutputWithContext(context.Background())
}

func (i ClusterDefinitionArgs) ToClusterDefinitionOutputWithContext(ctx context.Context) ClusterDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterDefinitionOutput)
}

func (i ClusterDefinitionArgs) ToClusterDefinitionPtrOutput() ClusterDefinitionPtrOutput {
	return i.ToClusterDefinitionPtrOutputWithContext(context.Background())
}

func (i ClusterDefinitionArgs) ToClusterDefinitionPtrOutputWithContext(ctx context.Context) ClusterDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterDefinitionOutput).ToClusterDefinitionPtrOutputWithContext(ctx)
}









type ClusterDefinitionPtrInput interface {
	pulumi.Input

	ToClusterDefinitionPtrOutput() ClusterDefinitionPtrOutput
	ToClusterDefinitionPtrOutputWithContext(context.Context) ClusterDefinitionPtrOutput
}

type clusterDefinitionPtrType ClusterDefinitionArgs

func ClusterDefinitionPtr(v *ClusterDefinitionArgs) ClusterDefinitionPtrInput {
	return (*clusterDefinitionPtrType)(v)
}

func (*clusterDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterDefinition)(nil)).Elem()
}

func (i *clusterDefinitionPtrType) ToClusterDefinitionPtrOutput() ClusterDefinitionPtrOutput {
	return i.ToClusterDefinitionPtrOutputWithContext(context.Background())
}

func (i *clusterDefinitionPtrType) ToClusterDefinitionPtrOutputWithContext(ctx context.Context) ClusterDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterDefinitionPtrOutput)
}

type ClusterDefinitionOutput struct{ *pulumi.OutputState }

func (ClusterDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterDefinition)(nil)).Elem()
}

func (o ClusterDefinitionOutput) ToClusterDefinitionOutput() ClusterDefinitionOutput {
	return o
}

func (o ClusterDefinitionOutput) ToClusterDefinitionOutputWithContext(ctx context.Context) ClusterDefinitionOutput {
	return o
}

func (o ClusterDefinitionOutput) ToClusterDefinitionPtrOutput() ClusterDefinitionPtrOutput {
	return o.ToClusterDefinitionPtrOutputWithContext(context.Background())
}

func (o ClusterDefinitionOutput) ToClusterDefinitionPtrOutputWithContext(ctx context.Context) ClusterDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterDefinition) *ClusterDefinition {
		return &v
	}).(ClusterDefinitionPtrOutput)
}

func (o ClusterDefinitionOutput) Blueprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterDefinition) *string { return v.Blueprint }).(pulumi.StringPtrOutput)
}

func (o ClusterDefinitionOutput) ComponentVersion() pulumi.StringMapOutput {
	return o.ApplyT(func(v ClusterDefinition) map[string]string { return v.ComponentVersion }).(pulumi.StringMapOutput)
}

func (o ClusterDefinitionOutput) Configurations() pulumi.AnyOutput {
	return o.ApplyT(func(v ClusterDefinition) interface{} { return v.Configurations }).(pulumi.AnyOutput)
}

func (o ClusterDefinitionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterDefinition) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

type ClusterDefinitionPtrOutput struct{ *pulumi.OutputState }

func (ClusterDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterDefinition)(nil)).Elem()
}

func (o ClusterDefinitionPtrOutput) ToClusterDefinitionPtrOutput() ClusterDefinitionPtrOutput {
	return o
}

func (o ClusterDefinitionPtrOutput) ToClusterDefinitionPtrOutputWithContext(ctx context.Context) ClusterDefinitionPtrOutput {
	return o
}

func (o ClusterDefinitionPtrOutput) Elem() ClusterDefinitionOutput {
	return o.ApplyT(func(v *ClusterDefinition) ClusterDefinition {
		if v != nil {
			return *v
		}
		var ret ClusterDefinition
		return ret
	}).(ClusterDefinitionOutput)
}

func (o ClusterDefinitionPtrOutput) Blueprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Blueprint
	}).(pulumi.StringPtrOutput)
}

func (o ClusterDefinitionPtrOutput) ComponentVersion() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClusterDefinition) map[string]string {
		if v == nil {
			return nil
		}
		return v.ComponentVersion
	}).(pulumi.StringMapOutput)
}

func (o ClusterDefinitionPtrOutput) Configurations() pulumi.AnyOutput {
	return o.ApplyT(func(v *ClusterDefinition) interface{} {
		if v == nil {
			return nil
		}
		return v.Configurations
	}).(pulumi.AnyOutput)
}

func (o ClusterDefinitionPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

type ClusterDefinitionResponse struct {
	Blueprint        *string           `pulumi:"blueprint"`
	ComponentVersion map[string]string `pulumi:"componentVersion"`
	Configurations   interface{}       `pulumi:"configurations"`
	Kind             *string           `pulumi:"kind"`
}

type ClusterDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ClusterDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterDefinitionResponse)(nil)).Elem()
}

func (o ClusterDefinitionResponseOutput) ToClusterDefinitionResponseOutput() ClusterDefinitionResponseOutput {
	return o
}

func (o ClusterDefinitionResponseOutput) ToClusterDefinitionResponseOutputWithContext(ctx context.Context) ClusterDefinitionResponseOutput {
	return o
}

func (o ClusterDefinitionResponseOutput) Blueprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterDefinitionResponse) *string { return v.Blueprint }).(pulumi.StringPtrOutput)
}

func (o ClusterDefinitionResponseOutput) ComponentVersion() pulumi.StringMapOutput {
	return o.ApplyT(func(v ClusterDefinitionResponse) map[string]string { return v.ComponentVersion }).(pulumi.StringMapOutput)
}

func (o ClusterDefinitionResponseOutput) Configurations() pulumi.AnyOutput {
	return o.ApplyT(func(v ClusterDefinitionResponse) interface{} { return v.Configurations }).(pulumi.AnyOutput)
}

func (o ClusterDefinitionResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterDefinitionResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

type ClusterGetPropertiesResponse struct {
	ClusterDefinition             ClusterDefinitionResponse              `pulumi:"clusterDefinition"`
	ClusterHdpVersion             *string                                `pulumi:"clusterHdpVersion"`
	ClusterId                     *string                                `pulumi:"clusterId"`
	ClusterState                  *string                                `pulumi:"clusterState"`
	ClusterVersion                *string                                `pulumi:"clusterVersion"`
	ComputeIsolationProperties    *ComputeIsolationPropertiesResponse    `pulumi:"computeIsolationProperties"`
	ComputeProfile                *ComputeProfileResponse                `pulumi:"computeProfile"`
	ConnectivityEndpoints         []ConnectivityEndpointResponse         `pulumi:"connectivityEndpoints"`
	CreatedDate                   *string                                `pulumi:"createdDate"`
	DiskEncryptionProperties      *DiskEncryptionPropertiesResponse      `pulumi:"diskEncryptionProperties"`
	EncryptionInTransitProperties *EncryptionInTransitPropertiesResponse `pulumi:"encryptionInTransitProperties"`
	Errors                        []ErrorsResponse                       `pulumi:"errors"`
	ExcludedServicesConfig        *ExcludedServicesConfigResponse        `pulumi:"excludedServicesConfig"`
	KafkaRestProperties           *KafkaRestPropertiesResponse           `pulumi:"kafkaRestProperties"`
	MinSupportedTlsVersion        *string                                `pulumi:"minSupportedTlsVersion"`
	NetworkProperties             *NetworkPropertiesResponse             `pulumi:"networkProperties"`
	OsType                        *string                                `pulumi:"osType"`
	PrivateEndpointConnections    []PrivateEndpointConnectionResponse    `pulumi:"privateEndpointConnections"`
	PrivateLinkConfigurations     []PrivateLinkConfigurationResponse     `pulumi:"privateLinkConfigurations"`
	ProvisioningState             *string                                `pulumi:"provisioningState"`
	QuotaInfo                     *QuotaInfoResponse                     `pulumi:"quotaInfo"`
	SecurityProfile               *SecurityProfileResponse               `pulumi:"securityProfile"`
	StorageProfile                *StorageProfileResponse                `pulumi:"storageProfile"`
	Tier                          *string                                `pulumi:"tier"`
}


func (val *ClusterGetPropertiesResponse) Defaults() *ClusterGetPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ComputeIsolationProperties = tmp.ComputeIsolationProperties.Defaults()

	tmp.DiskEncryptionProperties = tmp.DiskEncryptionProperties.Defaults()

	tmp.EncryptionInTransitProperties = tmp.EncryptionInTransitProperties.Defaults()

	return &tmp
}

type ClusterGetPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ClusterGetPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterGetPropertiesResponse)(nil)).Elem()
}

func (o ClusterGetPropertiesResponseOutput) ToClusterGetPropertiesResponseOutput() ClusterGetPropertiesResponseOutput {
	return o
}

func (o ClusterGetPropertiesResponseOutput) ToClusterGetPropertiesResponseOutputWithContext(ctx context.Context) ClusterGetPropertiesResponseOutput {
	return o
}

func (o ClusterGetPropertiesResponseOutput) ClusterDefinition() ClusterDefinitionResponseOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) ClusterDefinitionResponse { return v.ClusterDefinition }).(ClusterDefinitionResponseOutput)
}

func (o ClusterGetPropertiesResponseOutput) ClusterHdpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *string { return v.ClusterHdpVersion }).(pulumi.StringPtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) ClusterState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *string { return v.ClusterState }).(pulumi.StringPtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) ClusterVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *string { return v.ClusterVersion }).(pulumi.StringPtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) ComputeIsolationProperties() ComputeIsolationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *ComputeIsolationPropertiesResponse {
		return v.ComputeIsolationProperties
	}).(ComputeIsolationPropertiesResponsePtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) ComputeProfile() ComputeProfileResponsePtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *ComputeProfileResponse { return v.ComputeProfile }).(ComputeProfileResponsePtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) ConnectivityEndpoints() ConnectivityEndpointResponseArrayOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) []ConnectivityEndpointResponse { return v.ConnectivityEndpoints }).(ConnectivityEndpointResponseArrayOutput)
}

func (o ClusterGetPropertiesResponseOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *string { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) DiskEncryptionProperties() DiskEncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *DiskEncryptionPropertiesResponse {
		return v.DiskEncryptionProperties
	}).(DiskEncryptionPropertiesResponsePtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) EncryptionInTransitProperties() EncryptionInTransitPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *EncryptionInTransitPropertiesResponse {
		return v.EncryptionInTransitProperties
	}).(EncryptionInTransitPropertiesResponsePtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) Errors() ErrorsResponseArrayOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) []ErrorsResponse { return v.Errors }).(ErrorsResponseArrayOutput)
}

func (o ClusterGetPropertiesResponseOutput) ExcludedServicesConfig() ExcludedServicesConfigResponsePtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *ExcludedServicesConfigResponse { return v.ExcludedServicesConfig }).(ExcludedServicesConfigResponsePtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) KafkaRestProperties() KafkaRestPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *KafkaRestPropertiesResponse { return v.KafkaRestProperties }).(KafkaRestPropertiesResponsePtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) MinSupportedTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *string { return v.MinSupportedTlsVersion }).(pulumi.StringPtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) NetworkProperties() NetworkPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *NetworkPropertiesResponse { return v.NetworkProperties }).(NetworkPropertiesResponsePtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o ClusterGetPropertiesResponseOutput) PrivateLinkConfigurations() PrivateLinkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) []PrivateLinkConfigurationResponse {
		return v.PrivateLinkConfigurations
	}).(PrivateLinkConfigurationResponseArrayOutput)
}

func (o ClusterGetPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) QuotaInfo() QuotaInfoResponsePtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *QuotaInfoResponse { return v.QuotaInfo }).(QuotaInfoResponsePtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) SecurityProfile() SecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *SecurityProfileResponse { return v.SecurityProfile }).(SecurityProfileResponsePtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o ClusterGetPropertiesResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterGetPropertiesResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ClusterIdentity struct {
	Type                   *string                         `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentity `pulumi:"userAssignedIdentities"`
}





type ClusterIdentityInput interface {
	pulumi.Input

	ToClusterIdentityOutput() ClusterIdentityOutput
	ToClusterIdentityOutputWithContext(context.Context) ClusterIdentityOutput
}

type ClusterIdentityArgs struct {
	Type                   pulumi.StringPtrInput        `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityMapInput `pulumi:"userAssignedIdentities"`
}

func (ClusterIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterIdentity)(nil)).Elem()
}

func (i ClusterIdentityArgs) ToClusterIdentityOutput() ClusterIdentityOutput {
	return i.ToClusterIdentityOutputWithContext(context.Background())
}

func (i ClusterIdentityArgs) ToClusterIdentityOutputWithContext(ctx context.Context) ClusterIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIdentityOutput)
}

func (i ClusterIdentityArgs) ToClusterIdentityPtrOutput() ClusterIdentityPtrOutput {
	return i.ToClusterIdentityPtrOutputWithContext(context.Background())
}

func (i ClusterIdentityArgs) ToClusterIdentityPtrOutputWithContext(ctx context.Context) ClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIdentityOutput).ToClusterIdentityPtrOutputWithContext(ctx)
}









type ClusterIdentityPtrInput interface {
	pulumi.Input

	ToClusterIdentityPtrOutput() ClusterIdentityPtrOutput
	ToClusterIdentityPtrOutputWithContext(context.Context) ClusterIdentityPtrOutput
}

type clusterIdentityPtrType ClusterIdentityArgs

func ClusterIdentityPtr(v *ClusterIdentityArgs) ClusterIdentityPtrInput {
	return (*clusterIdentityPtrType)(v)
}

func (*clusterIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIdentity)(nil)).Elem()
}

func (i *clusterIdentityPtrType) ToClusterIdentityPtrOutput() ClusterIdentityPtrOutput {
	return i.ToClusterIdentityPtrOutputWithContext(context.Background())
}

func (i *clusterIdentityPtrType) ToClusterIdentityPtrOutputWithContext(ctx context.Context) ClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIdentityPtrOutput)
}

type ClusterIdentityOutput struct{ *pulumi.OutputState }

func (ClusterIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterIdentity)(nil)).Elem()
}

func (o ClusterIdentityOutput) ToClusterIdentityOutput() ClusterIdentityOutput {
	return o
}

func (o ClusterIdentityOutput) ToClusterIdentityOutputWithContext(ctx context.Context) ClusterIdentityOutput {
	return o
}

func (o ClusterIdentityOutput) ToClusterIdentityPtrOutput() ClusterIdentityPtrOutput {
	return o.ToClusterIdentityPtrOutputWithContext(context.Background())
}

func (o ClusterIdentityOutput) ToClusterIdentityPtrOutputWithContext(ctx context.Context) ClusterIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterIdentity) *ClusterIdentity {
		return &v
	}).(ClusterIdentityPtrOutput)
}

func (o ClusterIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ClusterIdentityOutput) UserAssignedIdentities() UserAssignedIdentityMapOutput {
	return o.ApplyT(func(v ClusterIdentity) map[string]UserAssignedIdentity { return v.UserAssignedIdentities }).(UserAssignedIdentityMapOutput)
}

type ClusterIdentityPtrOutput struct{ *pulumi.OutputState }

func (ClusterIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIdentity)(nil)).Elem()
}

func (o ClusterIdentityPtrOutput) ToClusterIdentityPtrOutput() ClusterIdentityPtrOutput {
	return o
}

func (o ClusterIdentityPtrOutput) ToClusterIdentityPtrOutputWithContext(ctx context.Context) ClusterIdentityPtrOutput {
	return o
}

func (o ClusterIdentityPtrOutput) Elem() ClusterIdentityOutput {
	return o.ApplyT(func(v *ClusterIdentity) ClusterIdentity {
		if v != nil {
			return *v
		}
		var ret ClusterIdentity
		return ret
	}).(ClusterIdentityOutput)
}

func (o ClusterIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ClusterIdentityPtrOutput) UserAssignedIdentities() UserAssignedIdentityMapOutput {
	return o.ApplyT(func(v *ClusterIdentity) map[string]UserAssignedIdentity {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityMapOutput)
}

type ClusterIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ClusterIdentityResponseOutput struct{ *pulumi.OutputState }

func (ClusterIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterIdentityResponse)(nil)).Elem()
}

func (o ClusterIdentityResponseOutput) ToClusterIdentityResponseOutput() ClusterIdentityResponseOutput {
	return o
}

func (o ClusterIdentityResponseOutput) ToClusterIdentityResponseOutputWithContext(ctx context.Context) ClusterIdentityResponseOutput {
	return o
}

func (o ClusterIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ClusterIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ClusterIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ClusterIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ClusterIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ClusterIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIdentityResponse)(nil)).Elem()
}

func (o ClusterIdentityResponsePtrOutput) ToClusterIdentityResponsePtrOutput() ClusterIdentityResponsePtrOutput {
	return o
}

func (o ClusterIdentityResponsePtrOutput) ToClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ClusterIdentityResponsePtrOutput {
	return o
}

func (o ClusterIdentityResponsePtrOutput) Elem() ClusterIdentityResponseOutput {
	return o.ApplyT(func(v *ClusterIdentityResponse) ClusterIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ClusterIdentityResponse
		return ret
	}).(ClusterIdentityResponseOutput)
}

func (o ClusterIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ClusterIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ClusterIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ClusterIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ClusterIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ComputeIsolationProperties struct {
	EnableComputeIsolation *bool   `pulumi:"enableComputeIsolation"`
	HostSku                *string `pulumi:"hostSku"`
}


func (val *ComputeIsolationProperties) Defaults() *ComputeIsolationProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableComputeIsolation) {
		enableComputeIsolation_ := false
		tmp.EnableComputeIsolation = &enableComputeIsolation_
	}
	return &tmp
}





type ComputeIsolationPropertiesInput interface {
	pulumi.Input

	ToComputeIsolationPropertiesOutput() ComputeIsolationPropertiesOutput
	ToComputeIsolationPropertiesOutputWithContext(context.Context) ComputeIsolationPropertiesOutput
}

type ComputeIsolationPropertiesArgs struct {
	EnableComputeIsolation pulumi.BoolPtrInput   `pulumi:"enableComputeIsolation"`
	HostSku                pulumi.StringPtrInput `pulumi:"hostSku"`
}


func (val *ComputeIsolationPropertiesArgs) Defaults() *ComputeIsolationPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableComputeIsolation) {
		tmp.EnableComputeIsolation = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (ComputeIsolationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeIsolationProperties)(nil)).Elem()
}

func (i ComputeIsolationPropertiesArgs) ToComputeIsolationPropertiesOutput() ComputeIsolationPropertiesOutput {
	return i.ToComputeIsolationPropertiesOutputWithContext(context.Background())
}

func (i ComputeIsolationPropertiesArgs) ToComputeIsolationPropertiesOutputWithContext(ctx context.Context) ComputeIsolationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeIsolationPropertiesOutput)
}

func (i ComputeIsolationPropertiesArgs) ToComputeIsolationPropertiesPtrOutput() ComputeIsolationPropertiesPtrOutput {
	return i.ToComputeIsolationPropertiesPtrOutputWithContext(context.Background())
}

func (i ComputeIsolationPropertiesArgs) ToComputeIsolationPropertiesPtrOutputWithContext(ctx context.Context) ComputeIsolationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeIsolationPropertiesOutput).ToComputeIsolationPropertiesPtrOutputWithContext(ctx)
}









type ComputeIsolationPropertiesPtrInput interface {
	pulumi.Input

	ToComputeIsolationPropertiesPtrOutput() ComputeIsolationPropertiesPtrOutput
	ToComputeIsolationPropertiesPtrOutputWithContext(context.Context) ComputeIsolationPropertiesPtrOutput
}

type computeIsolationPropertiesPtrType ComputeIsolationPropertiesArgs

func ComputeIsolationPropertiesPtr(v *ComputeIsolationPropertiesArgs) ComputeIsolationPropertiesPtrInput {
	return (*computeIsolationPropertiesPtrType)(v)
}

func (*computeIsolationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeIsolationProperties)(nil)).Elem()
}

func (i *computeIsolationPropertiesPtrType) ToComputeIsolationPropertiesPtrOutput() ComputeIsolationPropertiesPtrOutput {
	return i.ToComputeIsolationPropertiesPtrOutputWithContext(context.Background())
}

func (i *computeIsolationPropertiesPtrType) ToComputeIsolationPropertiesPtrOutputWithContext(ctx context.Context) ComputeIsolationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeIsolationPropertiesPtrOutput)
}

type ComputeIsolationPropertiesOutput struct{ *pulumi.OutputState }

func (ComputeIsolationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeIsolationProperties)(nil)).Elem()
}

func (o ComputeIsolationPropertiesOutput) ToComputeIsolationPropertiesOutput() ComputeIsolationPropertiesOutput {
	return o
}

func (o ComputeIsolationPropertiesOutput) ToComputeIsolationPropertiesOutputWithContext(ctx context.Context) ComputeIsolationPropertiesOutput {
	return o
}

func (o ComputeIsolationPropertiesOutput) ToComputeIsolationPropertiesPtrOutput() ComputeIsolationPropertiesPtrOutput {
	return o.ToComputeIsolationPropertiesPtrOutputWithContext(context.Background())
}

func (o ComputeIsolationPropertiesOutput) ToComputeIsolationPropertiesPtrOutputWithContext(ctx context.Context) ComputeIsolationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeIsolationProperties) *ComputeIsolationProperties {
		return &v
	}).(ComputeIsolationPropertiesPtrOutput)
}

func (o ComputeIsolationPropertiesOutput) EnableComputeIsolation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ComputeIsolationProperties) *bool { return v.EnableComputeIsolation }).(pulumi.BoolPtrOutput)
}

func (o ComputeIsolationPropertiesOutput) HostSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeIsolationProperties) *string { return v.HostSku }).(pulumi.StringPtrOutput)
}

type ComputeIsolationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ComputeIsolationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeIsolationProperties)(nil)).Elem()
}

func (o ComputeIsolationPropertiesPtrOutput) ToComputeIsolationPropertiesPtrOutput() ComputeIsolationPropertiesPtrOutput {
	return o
}

func (o ComputeIsolationPropertiesPtrOutput) ToComputeIsolationPropertiesPtrOutputWithContext(ctx context.Context) ComputeIsolationPropertiesPtrOutput {
	return o
}

func (o ComputeIsolationPropertiesPtrOutput) Elem() ComputeIsolationPropertiesOutput {
	return o.ApplyT(func(v *ComputeIsolationProperties) ComputeIsolationProperties {
		if v != nil {
			return *v
		}
		var ret ComputeIsolationProperties
		return ret
	}).(ComputeIsolationPropertiesOutput)
}

func (o ComputeIsolationPropertiesPtrOutput) EnableComputeIsolation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeIsolationProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableComputeIsolation
	}).(pulumi.BoolPtrOutput)
}

func (o ComputeIsolationPropertiesPtrOutput) HostSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeIsolationProperties) *string {
		if v == nil {
			return nil
		}
		return v.HostSku
	}).(pulumi.StringPtrOutput)
}

type ComputeIsolationPropertiesResponse struct {
	EnableComputeIsolation *bool   `pulumi:"enableComputeIsolation"`
	HostSku                *string `pulumi:"hostSku"`
}


func (val *ComputeIsolationPropertiesResponse) Defaults() *ComputeIsolationPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableComputeIsolation) {
		enableComputeIsolation_ := false
		tmp.EnableComputeIsolation = &enableComputeIsolation_
	}
	return &tmp
}

type ComputeIsolationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ComputeIsolationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeIsolationPropertiesResponse)(nil)).Elem()
}

func (o ComputeIsolationPropertiesResponseOutput) ToComputeIsolationPropertiesResponseOutput() ComputeIsolationPropertiesResponseOutput {
	return o
}

func (o ComputeIsolationPropertiesResponseOutput) ToComputeIsolationPropertiesResponseOutputWithContext(ctx context.Context) ComputeIsolationPropertiesResponseOutput {
	return o
}

func (o ComputeIsolationPropertiesResponseOutput) EnableComputeIsolation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ComputeIsolationPropertiesResponse) *bool { return v.EnableComputeIsolation }).(pulumi.BoolPtrOutput)
}

func (o ComputeIsolationPropertiesResponseOutput) HostSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeIsolationPropertiesResponse) *string { return v.HostSku }).(pulumi.StringPtrOutput)
}

type ComputeIsolationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ComputeIsolationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeIsolationPropertiesResponse)(nil)).Elem()
}

func (o ComputeIsolationPropertiesResponsePtrOutput) ToComputeIsolationPropertiesResponsePtrOutput() ComputeIsolationPropertiesResponsePtrOutput {
	return o
}

func (o ComputeIsolationPropertiesResponsePtrOutput) ToComputeIsolationPropertiesResponsePtrOutputWithContext(ctx context.Context) ComputeIsolationPropertiesResponsePtrOutput {
	return o
}

func (o ComputeIsolationPropertiesResponsePtrOutput) Elem() ComputeIsolationPropertiesResponseOutput {
	return o.ApplyT(func(v *ComputeIsolationPropertiesResponse) ComputeIsolationPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ComputeIsolationPropertiesResponse
		return ret
	}).(ComputeIsolationPropertiesResponseOutput)
}

func (o ComputeIsolationPropertiesResponsePtrOutput) EnableComputeIsolation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeIsolationPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableComputeIsolation
	}).(pulumi.BoolPtrOutput)
}

func (o ComputeIsolationPropertiesResponsePtrOutput) HostSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeIsolationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.HostSku
	}).(pulumi.StringPtrOutput)
}

type ComputeProfile struct {
	Roles []Role `pulumi:"roles"`
}





type ComputeProfileInput interface {
	pulumi.Input

	ToComputeProfileOutput() ComputeProfileOutput
	ToComputeProfileOutputWithContext(context.Context) ComputeProfileOutput
}

type ComputeProfileArgs struct {
	Roles RoleArrayInput `pulumi:"roles"`
}

func (ComputeProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeProfile)(nil)).Elem()
}

func (i ComputeProfileArgs) ToComputeProfileOutput() ComputeProfileOutput {
	return i.ToComputeProfileOutputWithContext(context.Background())
}

func (i ComputeProfileArgs) ToComputeProfileOutputWithContext(ctx context.Context) ComputeProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeProfileOutput)
}

func (i ComputeProfileArgs) ToComputeProfilePtrOutput() ComputeProfilePtrOutput {
	return i.ToComputeProfilePtrOutputWithContext(context.Background())
}

func (i ComputeProfileArgs) ToComputeProfilePtrOutputWithContext(ctx context.Context) ComputeProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeProfileOutput).ToComputeProfilePtrOutputWithContext(ctx)
}









type ComputeProfilePtrInput interface {
	pulumi.Input

	ToComputeProfilePtrOutput() ComputeProfilePtrOutput
	ToComputeProfilePtrOutputWithContext(context.Context) ComputeProfilePtrOutput
}

type computeProfilePtrType ComputeProfileArgs

func ComputeProfilePtr(v *ComputeProfileArgs) ComputeProfilePtrInput {
	return (*computeProfilePtrType)(v)
}

func (*computeProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeProfile)(nil)).Elem()
}

func (i *computeProfilePtrType) ToComputeProfilePtrOutput() ComputeProfilePtrOutput {
	return i.ToComputeProfilePtrOutputWithContext(context.Background())
}

func (i *computeProfilePtrType) ToComputeProfilePtrOutputWithContext(ctx context.Context) ComputeProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeProfilePtrOutput)
}

type ComputeProfileOutput struct{ *pulumi.OutputState }

func (ComputeProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeProfile)(nil)).Elem()
}

func (o ComputeProfileOutput) ToComputeProfileOutput() ComputeProfileOutput {
	return o
}

func (o ComputeProfileOutput) ToComputeProfileOutputWithContext(ctx context.Context) ComputeProfileOutput {
	return o
}

func (o ComputeProfileOutput) ToComputeProfilePtrOutput() ComputeProfilePtrOutput {
	return o.ToComputeProfilePtrOutputWithContext(context.Background())
}

func (o ComputeProfileOutput) ToComputeProfilePtrOutputWithContext(ctx context.Context) ComputeProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeProfile) *ComputeProfile {
		return &v
	}).(ComputeProfilePtrOutput)
}

func (o ComputeProfileOutput) Roles() RoleArrayOutput {
	return o.ApplyT(func(v ComputeProfile) []Role { return v.Roles }).(RoleArrayOutput)
}

type ComputeProfilePtrOutput struct{ *pulumi.OutputState }

func (ComputeProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeProfile)(nil)).Elem()
}

func (o ComputeProfilePtrOutput) ToComputeProfilePtrOutput() ComputeProfilePtrOutput {
	return o
}

func (o ComputeProfilePtrOutput) ToComputeProfilePtrOutputWithContext(ctx context.Context) ComputeProfilePtrOutput {
	return o
}

func (o ComputeProfilePtrOutput) Elem() ComputeProfileOutput {
	return o.ApplyT(func(v *ComputeProfile) ComputeProfile {
		if v != nil {
			return *v
		}
		var ret ComputeProfile
		return ret
	}).(ComputeProfileOutput)
}

func (o ComputeProfilePtrOutput) Roles() RoleArrayOutput {
	return o.ApplyT(func(v *ComputeProfile) []Role {
		if v == nil {
			return nil
		}
		return v.Roles
	}).(RoleArrayOutput)
}

type ComputeProfileResponse struct {
	Roles []RoleResponse `pulumi:"roles"`
}

type ComputeProfileResponseOutput struct{ *pulumi.OutputState }

func (ComputeProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeProfileResponse)(nil)).Elem()
}

func (o ComputeProfileResponseOutput) ToComputeProfileResponseOutput() ComputeProfileResponseOutput {
	return o
}

func (o ComputeProfileResponseOutput) ToComputeProfileResponseOutputWithContext(ctx context.Context) ComputeProfileResponseOutput {
	return o
}

func (o ComputeProfileResponseOutput) Roles() RoleResponseArrayOutput {
	return o.ApplyT(func(v ComputeProfileResponse) []RoleResponse { return v.Roles }).(RoleResponseArrayOutput)
}

type ComputeProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ComputeProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeProfileResponse)(nil)).Elem()
}

func (o ComputeProfileResponsePtrOutput) ToComputeProfileResponsePtrOutput() ComputeProfileResponsePtrOutput {
	return o
}

func (o ComputeProfileResponsePtrOutput) ToComputeProfileResponsePtrOutputWithContext(ctx context.Context) ComputeProfileResponsePtrOutput {
	return o
}

func (o ComputeProfileResponsePtrOutput) Elem() ComputeProfileResponseOutput {
	return o.ApplyT(func(v *ComputeProfileResponse) ComputeProfileResponse {
		if v != nil {
			return *v
		}
		var ret ComputeProfileResponse
		return ret
	}).(ComputeProfileResponseOutput)
}

func (o ComputeProfileResponsePtrOutput) Roles() RoleResponseArrayOutput {
	return o.ApplyT(func(v *ComputeProfileResponse) []RoleResponse {
		if v == nil {
			return nil
		}
		return v.Roles
	}).(RoleResponseArrayOutput)
}

type ConnectivityEndpointResponse struct {
	Location         *string `pulumi:"location"`
	Name             *string `pulumi:"name"`
	Port             *int    `pulumi:"port"`
	PrivateIPAddress *string `pulumi:"privateIPAddress"`
	Protocol         *string `pulumi:"protocol"`
}

type ConnectivityEndpointResponseOutput struct{ *pulumi.OutputState }

func (ConnectivityEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityEndpointResponse)(nil)).Elem()
}

func (o ConnectivityEndpointResponseOutput) ToConnectivityEndpointResponseOutput() ConnectivityEndpointResponseOutput {
	return o
}

func (o ConnectivityEndpointResponseOutput) ToConnectivityEndpointResponseOutputWithContext(ctx context.Context) ConnectivityEndpointResponseOutput {
	return o
}

func (o ConnectivityEndpointResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityEndpointResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ConnectivityEndpointResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityEndpointResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnectivityEndpointResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConnectivityEndpointResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ConnectivityEndpointResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityEndpointResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o ConnectivityEndpointResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectivityEndpointResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type ConnectivityEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (ConnectivityEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectivityEndpointResponse)(nil)).Elem()
}

func (o ConnectivityEndpointResponseArrayOutput) ToConnectivityEndpointResponseArrayOutput() ConnectivityEndpointResponseArrayOutput {
	return o
}

func (o ConnectivityEndpointResponseArrayOutput) ToConnectivityEndpointResponseArrayOutputWithContext(ctx context.Context) ConnectivityEndpointResponseArrayOutput {
	return o
}

func (o ConnectivityEndpointResponseArrayOutput) Index(i pulumi.IntInput) ConnectivityEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectivityEndpointResponse {
		return vs[0].([]ConnectivityEndpointResponse)[vs[1].(int)]
	}).(ConnectivityEndpointResponseOutput)
}

type DataDisksGroups struct {
	DisksPerNode *int `pulumi:"disksPerNode"`
}





type DataDisksGroupsInput interface {
	pulumi.Input

	ToDataDisksGroupsOutput() DataDisksGroupsOutput
	ToDataDisksGroupsOutputWithContext(context.Context) DataDisksGroupsOutput
}

type DataDisksGroupsArgs struct {
	DisksPerNode pulumi.IntPtrInput `pulumi:"disksPerNode"`
}

func (DataDisksGroupsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisksGroups)(nil)).Elem()
}

func (i DataDisksGroupsArgs) ToDataDisksGroupsOutput() DataDisksGroupsOutput {
	return i.ToDataDisksGroupsOutputWithContext(context.Background())
}

func (i DataDisksGroupsArgs) ToDataDisksGroupsOutputWithContext(ctx context.Context) DataDisksGroupsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDisksGroupsOutput)
}





type DataDisksGroupsArrayInput interface {
	pulumi.Input

	ToDataDisksGroupsArrayOutput() DataDisksGroupsArrayOutput
	ToDataDisksGroupsArrayOutputWithContext(context.Context) DataDisksGroupsArrayOutput
}

type DataDisksGroupsArray []DataDisksGroupsInput

func (DataDisksGroupsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisksGroups)(nil)).Elem()
}

func (i DataDisksGroupsArray) ToDataDisksGroupsArrayOutput() DataDisksGroupsArrayOutput {
	return i.ToDataDisksGroupsArrayOutputWithContext(context.Background())
}

func (i DataDisksGroupsArray) ToDataDisksGroupsArrayOutputWithContext(ctx context.Context) DataDisksGroupsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDisksGroupsArrayOutput)
}

type DataDisksGroupsOutput struct{ *pulumi.OutputState }

func (DataDisksGroupsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisksGroups)(nil)).Elem()
}

func (o DataDisksGroupsOutput) ToDataDisksGroupsOutput() DataDisksGroupsOutput {
	return o
}

func (o DataDisksGroupsOutput) ToDataDisksGroupsOutputWithContext(ctx context.Context) DataDisksGroupsOutput {
	return o
}

func (o DataDisksGroupsOutput) DisksPerNode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataDisksGroups) *int { return v.DisksPerNode }).(pulumi.IntPtrOutput)
}

type DataDisksGroupsArrayOutput struct{ *pulumi.OutputState }

func (DataDisksGroupsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisksGroups)(nil)).Elem()
}

func (o DataDisksGroupsArrayOutput) ToDataDisksGroupsArrayOutput() DataDisksGroupsArrayOutput {
	return o
}

func (o DataDisksGroupsArrayOutput) ToDataDisksGroupsArrayOutputWithContext(ctx context.Context) DataDisksGroupsArrayOutput {
	return o
}

func (o DataDisksGroupsArrayOutput) Index(i pulumi.IntInput) DataDisksGroupsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDisksGroups {
		return vs[0].([]DataDisksGroups)[vs[1].(int)]
	}).(DataDisksGroupsOutput)
}

type DataDisksGroupsResponse struct {
	DiskSizeGB         int    `pulumi:"diskSizeGB"`
	DisksPerNode       *int   `pulumi:"disksPerNode"`
	StorageAccountType string `pulumi:"storageAccountType"`
}

type DataDisksGroupsResponseOutput struct{ *pulumi.OutputState }

func (DataDisksGroupsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisksGroupsResponse)(nil)).Elem()
}

func (o DataDisksGroupsResponseOutput) ToDataDisksGroupsResponseOutput() DataDisksGroupsResponseOutput {
	return o
}

func (o DataDisksGroupsResponseOutput) ToDataDisksGroupsResponseOutputWithContext(ctx context.Context) DataDisksGroupsResponseOutput {
	return o
}

func (o DataDisksGroupsResponseOutput) DiskSizeGB() pulumi.IntOutput {
	return o.ApplyT(func(v DataDisksGroupsResponse) int { return v.DiskSizeGB }).(pulumi.IntOutput)
}

func (o DataDisksGroupsResponseOutput) DisksPerNode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataDisksGroupsResponse) *int { return v.DisksPerNode }).(pulumi.IntPtrOutput)
}

func (o DataDisksGroupsResponseOutput) StorageAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v DataDisksGroupsResponse) string { return v.StorageAccountType }).(pulumi.StringOutput)
}

type DataDisksGroupsResponseArrayOutput struct{ *pulumi.OutputState }

func (DataDisksGroupsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisksGroupsResponse)(nil)).Elem()
}

func (o DataDisksGroupsResponseArrayOutput) ToDataDisksGroupsResponseArrayOutput() DataDisksGroupsResponseArrayOutput {
	return o
}

func (o DataDisksGroupsResponseArrayOutput) ToDataDisksGroupsResponseArrayOutputWithContext(ctx context.Context) DataDisksGroupsResponseArrayOutput {
	return o
}

func (o DataDisksGroupsResponseArrayOutput) Index(i pulumi.IntInput) DataDisksGroupsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDisksGroupsResponse {
		return vs[0].([]DataDisksGroupsResponse)[vs[1].(int)]
	}).(DataDisksGroupsResponseOutput)
}

type DiskEncryptionProperties struct {
	EncryptionAlgorithm *string `pulumi:"encryptionAlgorithm"`
	EncryptionAtHost    *bool   `pulumi:"encryptionAtHost"`
	KeyName             *string `pulumi:"keyName"`
	KeyVersion          *string `pulumi:"keyVersion"`
	MsiResourceId       *string `pulumi:"msiResourceId"`
	VaultUri            *string `pulumi:"vaultUri"`
}


func (val *DiskEncryptionProperties) Defaults() *DiskEncryptionProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EncryptionAtHost) {
		encryptionAtHost_ := false
		tmp.EncryptionAtHost = &encryptionAtHost_
	}
	return &tmp
}





type DiskEncryptionPropertiesInput interface {
	pulumi.Input

	ToDiskEncryptionPropertiesOutput() DiskEncryptionPropertiesOutput
	ToDiskEncryptionPropertiesOutputWithContext(context.Context) DiskEncryptionPropertiesOutput
}

type DiskEncryptionPropertiesArgs struct {
	EncryptionAlgorithm pulumi.StringPtrInput `pulumi:"encryptionAlgorithm"`
	EncryptionAtHost    pulumi.BoolPtrInput   `pulumi:"encryptionAtHost"`
	KeyName             pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVersion          pulumi.StringPtrInput `pulumi:"keyVersion"`
	MsiResourceId       pulumi.StringPtrInput `pulumi:"msiResourceId"`
	VaultUri            pulumi.StringPtrInput `pulumi:"vaultUri"`
}


func (val *DiskEncryptionPropertiesArgs) Defaults() *DiskEncryptionPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EncryptionAtHost) {
		tmp.EncryptionAtHost = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (DiskEncryptionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionProperties)(nil)).Elem()
}

func (i DiskEncryptionPropertiesArgs) ToDiskEncryptionPropertiesOutput() DiskEncryptionPropertiesOutput {
	return i.ToDiskEncryptionPropertiesOutputWithContext(context.Background())
}

func (i DiskEncryptionPropertiesArgs) ToDiskEncryptionPropertiesOutputWithContext(ctx context.Context) DiskEncryptionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionPropertiesOutput)
}

func (i DiskEncryptionPropertiesArgs) ToDiskEncryptionPropertiesPtrOutput() DiskEncryptionPropertiesPtrOutput {
	return i.ToDiskEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i DiskEncryptionPropertiesArgs) ToDiskEncryptionPropertiesPtrOutputWithContext(ctx context.Context) DiskEncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionPropertiesOutput).ToDiskEncryptionPropertiesPtrOutputWithContext(ctx)
}









type DiskEncryptionPropertiesPtrInput interface {
	pulumi.Input

	ToDiskEncryptionPropertiesPtrOutput() DiskEncryptionPropertiesPtrOutput
	ToDiskEncryptionPropertiesPtrOutputWithContext(context.Context) DiskEncryptionPropertiesPtrOutput
}

type diskEncryptionPropertiesPtrType DiskEncryptionPropertiesArgs

func DiskEncryptionPropertiesPtr(v *DiskEncryptionPropertiesArgs) DiskEncryptionPropertiesPtrInput {
	return (*diskEncryptionPropertiesPtrType)(v)
}

func (*diskEncryptionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionProperties)(nil)).Elem()
}

func (i *diskEncryptionPropertiesPtrType) ToDiskEncryptionPropertiesPtrOutput() DiskEncryptionPropertiesPtrOutput {
	return i.ToDiskEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i *diskEncryptionPropertiesPtrType) ToDiskEncryptionPropertiesPtrOutputWithContext(ctx context.Context) DiskEncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionPropertiesPtrOutput)
}

type DiskEncryptionPropertiesOutput struct{ *pulumi.OutputState }

func (DiskEncryptionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionProperties)(nil)).Elem()
}

func (o DiskEncryptionPropertiesOutput) ToDiskEncryptionPropertiesOutput() DiskEncryptionPropertiesOutput {
	return o
}

func (o DiskEncryptionPropertiesOutput) ToDiskEncryptionPropertiesOutputWithContext(ctx context.Context) DiskEncryptionPropertiesOutput {
	return o
}

func (o DiskEncryptionPropertiesOutput) ToDiskEncryptionPropertiesPtrOutput() DiskEncryptionPropertiesPtrOutput {
	return o.ToDiskEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (o DiskEncryptionPropertiesOutput) ToDiskEncryptionPropertiesPtrOutputWithContext(ctx context.Context) DiskEncryptionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskEncryptionProperties) *DiskEncryptionProperties {
		return &v
	}).(DiskEncryptionPropertiesPtrOutput)
}

func (o DiskEncryptionPropertiesOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionProperties) *string { return v.EncryptionAlgorithm }).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesOutput) EncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiskEncryptionProperties) *bool { return v.EncryptionAtHost }).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesOutput) MsiResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionProperties) *string { return v.MsiResourceId }).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionProperties) *string { return v.VaultUri }).(pulumi.StringPtrOutput)
}

type DiskEncryptionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionProperties)(nil)).Elem()
}

func (o DiskEncryptionPropertiesPtrOutput) ToDiskEncryptionPropertiesPtrOutput() DiskEncryptionPropertiesPtrOutput {
	return o
}

func (o DiskEncryptionPropertiesPtrOutput) ToDiskEncryptionPropertiesPtrOutputWithContext(ctx context.Context) DiskEncryptionPropertiesPtrOutput {
	return o
}

func (o DiskEncryptionPropertiesPtrOutput) Elem() DiskEncryptionPropertiesOutput {
	return o.ApplyT(func(v *DiskEncryptionProperties) DiskEncryptionProperties {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionProperties
		return ret
	}).(DiskEncryptionPropertiesOutput)
}

func (o DiskEncryptionPropertiesPtrOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesPtrOutput) EncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EncryptionAtHost
	}).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesPtrOutput) MsiResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return v.MsiResourceId
	}).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesPtrOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return v.VaultUri
	}).(pulumi.StringPtrOutput)
}

type DiskEncryptionPropertiesResponse struct {
	EncryptionAlgorithm *string `pulumi:"encryptionAlgorithm"`
	EncryptionAtHost    *bool   `pulumi:"encryptionAtHost"`
	KeyName             *string `pulumi:"keyName"`
	KeyVersion          *string `pulumi:"keyVersion"`
	MsiResourceId       *string `pulumi:"msiResourceId"`
	VaultUri            *string `pulumi:"vaultUri"`
}


func (val *DiskEncryptionPropertiesResponse) Defaults() *DiskEncryptionPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EncryptionAtHost) {
		encryptionAtHost_ := false
		tmp.EncryptionAtHost = &encryptionAtHost_
	}
	return &tmp
}

type DiskEncryptionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DiskEncryptionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionPropertiesResponse)(nil)).Elem()
}

func (o DiskEncryptionPropertiesResponseOutput) ToDiskEncryptionPropertiesResponseOutput() DiskEncryptionPropertiesResponseOutput {
	return o
}

func (o DiskEncryptionPropertiesResponseOutput) ToDiskEncryptionPropertiesResponseOutputWithContext(ctx context.Context) DiskEncryptionPropertiesResponseOutput {
	return o
}

func (o DiskEncryptionPropertiesResponseOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionPropertiesResponse) *string { return v.EncryptionAlgorithm }).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesResponseOutput) EncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiskEncryptionPropertiesResponse) *bool { return v.EncryptionAtHost }).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesResponseOutput) MsiResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionPropertiesResponse) *string { return v.MsiResourceId }).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesResponseOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionPropertiesResponse) *string { return v.VaultUri }).(pulumi.StringPtrOutput)
}

type DiskEncryptionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionPropertiesResponse)(nil)).Elem()
}

func (o DiskEncryptionPropertiesResponsePtrOutput) ToDiskEncryptionPropertiesResponsePtrOutput() DiskEncryptionPropertiesResponsePtrOutput {
	return o
}

func (o DiskEncryptionPropertiesResponsePtrOutput) ToDiskEncryptionPropertiesResponsePtrOutputWithContext(ctx context.Context) DiskEncryptionPropertiesResponsePtrOutput {
	return o
}

func (o DiskEncryptionPropertiesResponsePtrOutput) Elem() DiskEncryptionPropertiesResponseOutput {
	return o.ApplyT(func(v *DiskEncryptionPropertiesResponse) DiskEncryptionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionPropertiesResponse
		return ret
	}).(DiskEncryptionPropertiesResponseOutput)
}

func (o DiskEncryptionPropertiesResponsePtrOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesResponsePtrOutput) EncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EncryptionAtHost
	}).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesResponsePtrOutput) MsiResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.MsiResourceId
	}).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionPropertiesResponsePtrOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.VaultUri
	}).(pulumi.StringPtrOutput)
}

type EncryptionInTransitProperties struct {
	IsEncryptionInTransitEnabled *bool `pulumi:"isEncryptionInTransitEnabled"`
}


func (val *EncryptionInTransitProperties) Defaults() *EncryptionInTransitProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsEncryptionInTransitEnabled) {
		isEncryptionInTransitEnabled_ := false
		tmp.IsEncryptionInTransitEnabled = &isEncryptionInTransitEnabled_
	}
	return &tmp
}





type EncryptionInTransitPropertiesInput interface {
	pulumi.Input

	ToEncryptionInTransitPropertiesOutput() EncryptionInTransitPropertiesOutput
	ToEncryptionInTransitPropertiesOutputWithContext(context.Context) EncryptionInTransitPropertiesOutput
}

type EncryptionInTransitPropertiesArgs struct {
	IsEncryptionInTransitEnabled pulumi.BoolPtrInput `pulumi:"isEncryptionInTransitEnabled"`
}


func (val *EncryptionInTransitPropertiesArgs) Defaults() *EncryptionInTransitPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsEncryptionInTransitEnabled) {
		tmp.IsEncryptionInTransitEnabled = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (EncryptionInTransitPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionInTransitProperties)(nil)).Elem()
}

func (i EncryptionInTransitPropertiesArgs) ToEncryptionInTransitPropertiesOutput() EncryptionInTransitPropertiesOutput {
	return i.ToEncryptionInTransitPropertiesOutputWithContext(context.Background())
}

func (i EncryptionInTransitPropertiesArgs) ToEncryptionInTransitPropertiesOutputWithContext(ctx context.Context) EncryptionInTransitPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionInTransitPropertiesOutput)
}

func (i EncryptionInTransitPropertiesArgs) ToEncryptionInTransitPropertiesPtrOutput() EncryptionInTransitPropertiesPtrOutput {
	return i.ToEncryptionInTransitPropertiesPtrOutputWithContext(context.Background())
}

func (i EncryptionInTransitPropertiesArgs) ToEncryptionInTransitPropertiesPtrOutputWithContext(ctx context.Context) EncryptionInTransitPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionInTransitPropertiesOutput).ToEncryptionInTransitPropertiesPtrOutputWithContext(ctx)
}









type EncryptionInTransitPropertiesPtrInput interface {
	pulumi.Input

	ToEncryptionInTransitPropertiesPtrOutput() EncryptionInTransitPropertiesPtrOutput
	ToEncryptionInTransitPropertiesPtrOutputWithContext(context.Context) EncryptionInTransitPropertiesPtrOutput
}

type encryptionInTransitPropertiesPtrType EncryptionInTransitPropertiesArgs

func EncryptionInTransitPropertiesPtr(v *EncryptionInTransitPropertiesArgs) EncryptionInTransitPropertiesPtrInput {
	return (*encryptionInTransitPropertiesPtrType)(v)
}

func (*encryptionInTransitPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionInTransitProperties)(nil)).Elem()
}

func (i *encryptionInTransitPropertiesPtrType) ToEncryptionInTransitPropertiesPtrOutput() EncryptionInTransitPropertiesPtrOutput {
	return i.ToEncryptionInTransitPropertiesPtrOutputWithContext(context.Background())
}

func (i *encryptionInTransitPropertiesPtrType) ToEncryptionInTransitPropertiesPtrOutputWithContext(ctx context.Context) EncryptionInTransitPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionInTransitPropertiesPtrOutput)
}

type EncryptionInTransitPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionInTransitPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionInTransitProperties)(nil)).Elem()
}

func (o EncryptionInTransitPropertiesOutput) ToEncryptionInTransitPropertiesOutput() EncryptionInTransitPropertiesOutput {
	return o
}

func (o EncryptionInTransitPropertiesOutput) ToEncryptionInTransitPropertiesOutputWithContext(ctx context.Context) EncryptionInTransitPropertiesOutput {
	return o
}

func (o EncryptionInTransitPropertiesOutput) ToEncryptionInTransitPropertiesPtrOutput() EncryptionInTransitPropertiesPtrOutput {
	return o.ToEncryptionInTransitPropertiesPtrOutputWithContext(context.Background())
}

func (o EncryptionInTransitPropertiesOutput) ToEncryptionInTransitPropertiesPtrOutputWithContext(ctx context.Context) EncryptionInTransitPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionInTransitProperties) *EncryptionInTransitProperties {
		return &v
	}).(EncryptionInTransitPropertiesPtrOutput)
}

func (o EncryptionInTransitPropertiesOutput) IsEncryptionInTransitEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionInTransitProperties) *bool { return v.IsEncryptionInTransitEnabled }).(pulumi.BoolPtrOutput)
}

type EncryptionInTransitPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionInTransitPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionInTransitProperties)(nil)).Elem()
}

func (o EncryptionInTransitPropertiesPtrOutput) ToEncryptionInTransitPropertiesPtrOutput() EncryptionInTransitPropertiesPtrOutput {
	return o
}

func (o EncryptionInTransitPropertiesPtrOutput) ToEncryptionInTransitPropertiesPtrOutputWithContext(ctx context.Context) EncryptionInTransitPropertiesPtrOutput {
	return o
}

func (o EncryptionInTransitPropertiesPtrOutput) Elem() EncryptionInTransitPropertiesOutput {
	return o.ApplyT(func(v *EncryptionInTransitProperties) EncryptionInTransitProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionInTransitProperties
		return ret
	}).(EncryptionInTransitPropertiesOutput)
}

func (o EncryptionInTransitPropertiesPtrOutput) IsEncryptionInTransitEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionInTransitProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsEncryptionInTransitEnabled
	}).(pulumi.BoolPtrOutput)
}

type EncryptionInTransitPropertiesResponse struct {
	IsEncryptionInTransitEnabled *bool `pulumi:"isEncryptionInTransitEnabled"`
}


func (val *EncryptionInTransitPropertiesResponse) Defaults() *EncryptionInTransitPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsEncryptionInTransitEnabled) {
		isEncryptionInTransitEnabled_ := false
		tmp.IsEncryptionInTransitEnabled = &isEncryptionInTransitEnabled_
	}
	return &tmp
}

type EncryptionInTransitPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionInTransitPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionInTransitPropertiesResponse)(nil)).Elem()
}

func (o EncryptionInTransitPropertiesResponseOutput) ToEncryptionInTransitPropertiesResponseOutput() EncryptionInTransitPropertiesResponseOutput {
	return o
}

func (o EncryptionInTransitPropertiesResponseOutput) ToEncryptionInTransitPropertiesResponseOutputWithContext(ctx context.Context) EncryptionInTransitPropertiesResponseOutput {
	return o
}

func (o EncryptionInTransitPropertiesResponseOutput) IsEncryptionInTransitEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionInTransitPropertiesResponse) *bool { return v.IsEncryptionInTransitEnabled }).(pulumi.BoolPtrOutput)
}

type EncryptionInTransitPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionInTransitPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionInTransitPropertiesResponse)(nil)).Elem()
}

func (o EncryptionInTransitPropertiesResponsePtrOutput) ToEncryptionInTransitPropertiesResponsePtrOutput() EncryptionInTransitPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionInTransitPropertiesResponsePtrOutput) ToEncryptionInTransitPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionInTransitPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionInTransitPropertiesResponsePtrOutput) Elem() EncryptionInTransitPropertiesResponseOutput {
	return o.ApplyT(func(v *EncryptionInTransitPropertiesResponse) EncryptionInTransitPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionInTransitPropertiesResponse
		return ret
	}).(EncryptionInTransitPropertiesResponseOutput)
}

func (o EncryptionInTransitPropertiesResponsePtrOutput) IsEncryptionInTransitEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionInTransitPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsEncryptionInTransitEnabled
	}).(pulumi.BoolPtrOutput)
}

type Errors struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}





type ErrorsInput interface {
	pulumi.Input

	ToErrorsOutput() ErrorsOutput
	ToErrorsOutputWithContext(context.Context) ErrorsOutput
}

type ErrorsArgs struct {
	Code    pulumi.StringPtrInput `pulumi:"code"`
	Message pulumi.StringPtrInput `pulumi:"message"`
}

func (ErrorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Errors)(nil)).Elem()
}

func (i ErrorsArgs) ToErrorsOutput() ErrorsOutput {
	return i.ToErrorsOutputWithContext(context.Background())
}

func (i ErrorsArgs) ToErrorsOutputWithContext(ctx context.Context) ErrorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorsOutput)
}





type ErrorsArrayInput interface {
	pulumi.Input

	ToErrorsArrayOutput() ErrorsArrayOutput
	ToErrorsArrayOutputWithContext(context.Context) ErrorsArrayOutput
}

type ErrorsArray []ErrorsInput

func (ErrorsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Errors)(nil)).Elem()
}

func (i ErrorsArray) ToErrorsArrayOutput() ErrorsArrayOutput {
	return i.ToErrorsArrayOutputWithContext(context.Background())
}

func (i ErrorsArray) ToErrorsArrayOutputWithContext(ctx context.Context) ErrorsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorsArrayOutput)
}

type ErrorsOutput struct{ *pulumi.OutputState }

func (ErrorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Errors)(nil)).Elem()
}

func (o ErrorsOutput) ToErrorsOutput() ErrorsOutput {
	return o
}

func (o ErrorsOutput) ToErrorsOutputWithContext(ctx context.Context) ErrorsOutput {
	return o
}

func (o ErrorsOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Errors) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ErrorsOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Errors) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type ErrorsArrayOutput struct{ *pulumi.OutputState }

func (ErrorsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Errors)(nil)).Elem()
}

func (o ErrorsArrayOutput) ToErrorsArrayOutput() ErrorsArrayOutput {
	return o
}

func (o ErrorsArrayOutput) ToErrorsArrayOutputWithContext(ctx context.Context) ErrorsArrayOutput {
	return o
}

func (o ErrorsArrayOutput) Index(i pulumi.IntInput) ErrorsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Errors {
		return vs[0].([]Errors)[vs[1].(int)]
	}).(ErrorsOutput)
}

type ErrorsResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}

type ErrorsResponseOutput struct{ *pulumi.OutputState }

func (ErrorsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorsResponse)(nil)).Elem()
}

func (o ErrorsResponseOutput) ToErrorsResponseOutput() ErrorsResponseOutput {
	return o
}

func (o ErrorsResponseOutput) ToErrorsResponseOutputWithContext(ctx context.Context) ErrorsResponseOutput {
	return o
}

func (o ErrorsResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorsResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ErrorsResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorsResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type ErrorsResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorsResponse)(nil)).Elem()
}

func (o ErrorsResponseArrayOutput) ToErrorsResponseArrayOutput() ErrorsResponseArrayOutput {
	return o
}

func (o ErrorsResponseArrayOutput) ToErrorsResponseArrayOutputWithContext(ctx context.Context) ErrorsResponseArrayOutput {
	return o
}

func (o ErrorsResponseArrayOutput) Index(i pulumi.IntInput) ErrorsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorsResponse {
		return vs[0].([]ErrorsResponse)[vs[1].(int)]
	}).(ErrorsResponseOutput)
}

type ExcludedServicesConfigResponse struct {
	ExcludedServicesConfigId *string `pulumi:"excludedServicesConfigId"`
	ExcludedServicesList     *string `pulumi:"excludedServicesList"`
}

type ExcludedServicesConfigResponseOutput struct{ *pulumi.OutputState }

func (ExcludedServicesConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExcludedServicesConfigResponse)(nil)).Elem()
}

func (o ExcludedServicesConfigResponseOutput) ToExcludedServicesConfigResponseOutput() ExcludedServicesConfigResponseOutput {
	return o
}

func (o ExcludedServicesConfigResponseOutput) ToExcludedServicesConfigResponseOutputWithContext(ctx context.Context) ExcludedServicesConfigResponseOutput {
	return o
}

func (o ExcludedServicesConfigResponseOutput) ExcludedServicesConfigId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExcludedServicesConfigResponse) *string { return v.ExcludedServicesConfigId }).(pulumi.StringPtrOutput)
}

func (o ExcludedServicesConfigResponseOutput) ExcludedServicesList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExcludedServicesConfigResponse) *string { return v.ExcludedServicesList }).(pulumi.StringPtrOutput)
}

type ExcludedServicesConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (ExcludedServicesConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExcludedServicesConfigResponse)(nil)).Elem()
}

func (o ExcludedServicesConfigResponsePtrOutput) ToExcludedServicesConfigResponsePtrOutput() ExcludedServicesConfigResponsePtrOutput {
	return o
}

func (o ExcludedServicesConfigResponsePtrOutput) ToExcludedServicesConfigResponsePtrOutputWithContext(ctx context.Context) ExcludedServicesConfigResponsePtrOutput {
	return o
}

func (o ExcludedServicesConfigResponsePtrOutput) Elem() ExcludedServicesConfigResponseOutput {
	return o.ApplyT(func(v *ExcludedServicesConfigResponse) ExcludedServicesConfigResponse {
		if v != nil {
			return *v
		}
		var ret ExcludedServicesConfigResponse
		return ret
	}).(ExcludedServicesConfigResponseOutput)
}

func (o ExcludedServicesConfigResponsePtrOutput) ExcludedServicesConfigId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExcludedServicesConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExcludedServicesConfigId
	}).(pulumi.StringPtrOutput)
}

func (o ExcludedServicesConfigResponsePtrOutput) ExcludedServicesList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExcludedServicesConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExcludedServicesList
	}).(pulumi.StringPtrOutput)
}

type HardwareProfile struct {
	VmSize *string `pulumi:"vmSize"`
}





type HardwareProfileInput interface {
	pulumi.Input

	ToHardwareProfileOutput() HardwareProfileOutput
	ToHardwareProfileOutputWithContext(context.Context) HardwareProfileOutput
}

type HardwareProfileArgs struct {
	VmSize pulumi.StringPtrInput `pulumi:"vmSize"`
}

func (HardwareProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfile)(nil)).Elem()
}

func (i HardwareProfileArgs) ToHardwareProfileOutput() HardwareProfileOutput {
	return i.ToHardwareProfileOutputWithContext(context.Background())
}

func (i HardwareProfileArgs) ToHardwareProfileOutputWithContext(ctx context.Context) HardwareProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfileOutput)
}

func (i HardwareProfileArgs) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return i.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (i HardwareProfileArgs) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfileOutput).ToHardwareProfilePtrOutputWithContext(ctx)
}









type HardwareProfilePtrInput interface {
	pulumi.Input

	ToHardwareProfilePtrOutput() HardwareProfilePtrOutput
	ToHardwareProfilePtrOutputWithContext(context.Context) HardwareProfilePtrOutput
}

type hardwareProfilePtrType HardwareProfileArgs

func HardwareProfilePtr(v *HardwareProfileArgs) HardwareProfilePtrInput {
	return (*hardwareProfilePtrType)(v)
}

func (*hardwareProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfile)(nil)).Elem()
}

func (i *hardwareProfilePtrType) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return i.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (i *hardwareProfilePtrType) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfilePtrOutput)
}

type HardwareProfileOutput struct{ *pulumi.OutputState }

func (HardwareProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfile)(nil)).Elem()
}

func (o HardwareProfileOutput) ToHardwareProfileOutput() HardwareProfileOutput {
	return o
}

func (o HardwareProfileOutput) ToHardwareProfileOutputWithContext(ctx context.Context) HardwareProfileOutput {
	return o
}

func (o HardwareProfileOutput) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return o.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (o HardwareProfileOutput) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HardwareProfile) *HardwareProfile {
		return &v
	}).(HardwareProfilePtrOutput)
}

func (o HardwareProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type HardwareProfilePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfile)(nil)).Elem()
}

func (o HardwareProfilePtrOutput) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return o
}

func (o HardwareProfilePtrOutput) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return o
}

func (o HardwareProfilePtrOutput) Elem() HardwareProfileOutput {
	return o.ApplyT(func(v *HardwareProfile) HardwareProfile {
		if v != nil {
			return *v
		}
		var ret HardwareProfile
		return ret
	}).(HardwareProfileOutput)
}

func (o HardwareProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type HardwareProfileResponse struct {
	VmSize *string `pulumi:"vmSize"`
}

type HardwareProfileResponseOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutput() HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutputWithContext(ctx context.Context) HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type HardwareProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutput() HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutputWithContext(ctx context.Context) HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) Elem() HardwareProfileResponseOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) HardwareProfileResponse {
		if v != nil {
			return *v
		}
		var ret HardwareProfileResponse
		return ret
	}).(HardwareProfileResponseOutput)
}

func (o HardwareProfileResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type IPConfiguration struct {
	Name                      string      `pulumi:"name"`
	Primary                   *bool       `pulumi:"primary"`
	PrivateIPAddress          *string     `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string     `pulumi:"privateIPAllocationMethod"`
	Subnet                    *ResourceId `pulumi:"subnet"`
}





type IPConfigurationInput interface {
	pulumi.Input

	ToIPConfigurationOutput() IPConfigurationOutput
	ToIPConfigurationOutputWithContext(context.Context) IPConfigurationOutput
}

type IPConfigurationArgs struct {
	Name                      pulumi.StringInput    `pulumi:"name"`
	Primary                   pulumi.BoolPtrInput   `pulumi:"primary"`
	PrivateIPAddress          pulumi.StringPtrInput `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrInput `pulumi:"privateIPAllocationMethod"`
	Subnet                    ResourceIdPtrInput    `pulumi:"subnet"`
}

func (IPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPConfiguration)(nil)).Elem()
}

func (i IPConfigurationArgs) ToIPConfigurationOutput() IPConfigurationOutput {
	return i.ToIPConfigurationOutputWithContext(context.Background())
}

func (i IPConfigurationArgs) ToIPConfigurationOutputWithContext(ctx context.Context) IPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPConfigurationOutput)
}





type IPConfigurationArrayInput interface {
	pulumi.Input

	ToIPConfigurationArrayOutput() IPConfigurationArrayOutput
	ToIPConfigurationArrayOutputWithContext(context.Context) IPConfigurationArrayOutput
}

type IPConfigurationArray []IPConfigurationInput

func (IPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPConfiguration)(nil)).Elem()
}

func (i IPConfigurationArray) ToIPConfigurationArrayOutput() IPConfigurationArrayOutput {
	return i.ToIPConfigurationArrayOutputWithContext(context.Background())
}

func (i IPConfigurationArray) ToIPConfigurationArrayOutputWithContext(ctx context.Context) IPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPConfigurationArrayOutput)
}

type IPConfigurationOutput struct{ *pulumi.OutputState }

func (IPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPConfiguration)(nil)).Elem()
}

func (o IPConfigurationOutput) ToIPConfigurationOutput() IPConfigurationOutput {
	return o
}

func (o IPConfigurationOutput) ToIPConfigurationOutputWithContext(ctx context.Context) IPConfigurationOutput {
	return o
}

func (o IPConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IPConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o IPConfigurationOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IPConfiguration) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o IPConfigurationOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfiguration) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfiguration) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationOutput) Subnet() ResourceIdPtrOutput {
	return o.ApplyT(func(v IPConfiguration) *ResourceId { return v.Subnet }).(ResourceIdPtrOutput)
}

type IPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (IPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPConfiguration)(nil)).Elem()
}

func (o IPConfigurationArrayOutput) ToIPConfigurationArrayOutput() IPConfigurationArrayOutput {
	return o
}

func (o IPConfigurationArrayOutput) ToIPConfigurationArrayOutputWithContext(ctx context.Context) IPConfigurationArrayOutput {
	return o
}

func (o IPConfigurationArrayOutput) Index(i pulumi.IntInput) IPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPConfiguration {
		return vs[0].([]IPConfiguration)[vs[1].(int)]
	}).(IPConfigurationOutput)
}

type IPConfigurationResponse struct {
	Id                        string              `pulumi:"id"`
	Name                      string              `pulumi:"name"`
	Primary                   *bool               `pulumi:"primary"`
	PrivateIPAddress          *string             `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string             `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         string              `pulumi:"provisioningState"`
	Subnet                    *ResourceIdResponse `pulumi:"subnet"`
	Type                      string              `pulumi:"type"`
}

type IPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (IPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPConfigurationResponse)(nil)).Elem()
}

func (o IPConfigurationResponseOutput) ToIPConfigurationResponseOutput() IPConfigurationResponseOutput {
	return o
}

func (o IPConfigurationResponseOutput) ToIPConfigurationResponseOutputWithContext(ctx context.Context) IPConfigurationResponseOutput {
	return o
}

func (o IPConfigurationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IPConfigurationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o IPConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IPConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o IPConfigurationResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o IPConfigurationResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v IPConfigurationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o IPConfigurationResponseOutput) Subnet() ResourceIdResponsePtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *ResourceIdResponse { return v.Subnet }).(ResourceIdResponsePtrOutput)
}

func (o IPConfigurationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IPConfigurationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type IPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (IPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPConfigurationResponse)(nil)).Elem()
}

func (o IPConfigurationResponseArrayOutput) ToIPConfigurationResponseArrayOutput() IPConfigurationResponseArrayOutput {
	return o
}

func (o IPConfigurationResponseArrayOutput) ToIPConfigurationResponseArrayOutputWithContext(ctx context.Context) IPConfigurationResponseArrayOutput {
	return o
}

func (o IPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) IPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPConfigurationResponse {
		return vs[0].([]IPConfigurationResponse)[vs[1].(int)]
	}).(IPConfigurationResponseOutput)
}

type KafkaRestProperties struct {
	ClientGroupInfo       *ClientGroupInfo  `pulumi:"clientGroupInfo"`
	ConfigurationOverride map[string]string `pulumi:"configurationOverride"`
}





type KafkaRestPropertiesInput interface {
	pulumi.Input

	ToKafkaRestPropertiesOutput() KafkaRestPropertiesOutput
	ToKafkaRestPropertiesOutputWithContext(context.Context) KafkaRestPropertiesOutput
}

type KafkaRestPropertiesArgs struct {
	ClientGroupInfo       ClientGroupInfoPtrInput `pulumi:"clientGroupInfo"`
	ConfigurationOverride pulumi.StringMapInput   `pulumi:"configurationOverride"`
}

func (KafkaRestPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KafkaRestProperties)(nil)).Elem()
}

func (i KafkaRestPropertiesArgs) ToKafkaRestPropertiesOutput() KafkaRestPropertiesOutput {
	return i.ToKafkaRestPropertiesOutputWithContext(context.Background())
}

func (i KafkaRestPropertiesArgs) ToKafkaRestPropertiesOutputWithContext(ctx context.Context) KafkaRestPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaRestPropertiesOutput)
}

func (i KafkaRestPropertiesArgs) ToKafkaRestPropertiesPtrOutput() KafkaRestPropertiesPtrOutput {
	return i.ToKafkaRestPropertiesPtrOutputWithContext(context.Background())
}

func (i KafkaRestPropertiesArgs) ToKafkaRestPropertiesPtrOutputWithContext(ctx context.Context) KafkaRestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaRestPropertiesOutput).ToKafkaRestPropertiesPtrOutputWithContext(ctx)
}









type KafkaRestPropertiesPtrInput interface {
	pulumi.Input

	ToKafkaRestPropertiesPtrOutput() KafkaRestPropertiesPtrOutput
	ToKafkaRestPropertiesPtrOutputWithContext(context.Context) KafkaRestPropertiesPtrOutput
}

type kafkaRestPropertiesPtrType KafkaRestPropertiesArgs

func KafkaRestPropertiesPtr(v *KafkaRestPropertiesArgs) KafkaRestPropertiesPtrInput {
	return (*kafkaRestPropertiesPtrType)(v)
}

func (*kafkaRestPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaRestProperties)(nil)).Elem()
}

func (i *kafkaRestPropertiesPtrType) ToKafkaRestPropertiesPtrOutput() KafkaRestPropertiesPtrOutput {
	return i.ToKafkaRestPropertiesPtrOutputWithContext(context.Background())
}

func (i *kafkaRestPropertiesPtrType) ToKafkaRestPropertiesPtrOutputWithContext(ctx context.Context) KafkaRestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaRestPropertiesPtrOutput)
}

type KafkaRestPropertiesOutput struct{ *pulumi.OutputState }

func (KafkaRestPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KafkaRestProperties)(nil)).Elem()
}

func (o KafkaRestPropertiesOutput) ToKafkaRestPropertiesOutput() KafkaRestPropertiesOutput {
	return o
}

func (o KafkaRestPropertiesOutput) ToKafkaRestPropertiesOutputWithContext(ctx context.Context) KafkaRestPropertiesOutput {
	return o
}

func (o KafkaRestPropertiesOutput) ToKafkaRestPropertiesPtrOutput() KafkaRestPropertiesPtrOutput {
	return o.ToKafkaRestPropertiesPtrOutputWithContext(context.Background())
}

func (o KafkaRestPropertiesOutput) ToKafkaRestPropertiesPtrOutputWithContext(ctx context.Context) KafkaRestPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KafkaRestProperties) *KafkaRestProperties {
		return &v
	}).(KafkaRestPropertiesPtrOutput)
}

func (o KafkaRestPropertiesOutput) ClientGroupInfo() ClientGroupInfoPtrOutput {
	return o.ApplyT(func(v KafkaRestProperties) *ClientGroupInfo { return v.ClientGroupInfo }).(ClientGroupInfoPtrOutput)
}

func (o KafkaRestPropertiesOutput) ConfigurationOverride() pulumi.StringMapOutput {
	return o.ApplyT(func(v KafkaRestProperties) map[string]string { return v.ConfigurationOverride }).(pulumi.StringMapOutput)
}

type KafkaRestPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KafkaRestPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaRestProperties)(nil)).Elem()
}

func (o KafkaRestPropertiesPtrOutput) ToKafkaRestPropertiesPtrOutput() KafkaRestPropertiesPtrOutput {
	return o
}

func (o KafkaRestPropertiesPtrOutput) ToKafkaRestPropertiesPtrOutputWithContext(ctx context.Context) KafkaRestPropertiesPtrOutput {
	return o
}

func (o KafkaRestPropertiesPtrOutput) Elem() KafkaRestPropertiesOutput {
	return o.ApplyT(func(v *KafkaRestProperties) KafkaRestProperties {
		if v != nil {
			return *v
		}
		var ret KafkaRestProperties
		return ret
	}).(KafkaRestPropertiesOutput)
}

func (o KafkaRestPropertiesPtrOutput) ClientGroupInfo() ClientGroupInfoPtrOutput {
	return o.ApplyT(func(v *KafkaRestProperties) *ClientGroupInfo {
		if v == nil {
			return nil
		}
		return v.ClientGroupInfo
	}).(ClientGroupInfoPtrOutput)
}

func (o KafkaRestPropertiesPtrOutput) ConfigurationOverride() pulumi.StringMapOutput {
	return o.ApplyT(func(v *KafkaRestProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.ConfigurationOverride
	}).(pulumi.StringMapOutput)
}

type KafkaRestPropertiesResponse struct {
	ClientGroupInfo       *ClientGroupInfoResponse `pulumi:"clientGroupInfo"`
	ConfigurationOverride map[string]string        `pulumi:"configurationOverride"`
}

type KafkaRestPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KafkaRestPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KafkaRestPropertiesResponse)(nil)).Elem()
}

func (o KafkaRestPropertiesResponseOutput) ToKafkaRestPropertiesResponseOutput() KafkaRestPropertiesResponseOutput {
	return o
}

func (o KafkaRestPropertiesResponseOutput) ToKafkaRestPropertiesResponseOutputWithContext(ctx context.Context) KafkaRestPropertiesResponseOutput {
	return o
}

func (o KafkaRestPropertiesResponseOutput) ClientGroupInfo() ClientGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v KafkaRestPropertiesResponse) *ClientGroupInfoResponse { return v.ClientGroupInfo }).(ClientGroupInfoResponsePtrOutput)
}

func (o KafkaRestPropertiesResponseOutput) ConfigurationOverride() pulumi.StringMapOutput {
	return o.ApplyT(func(v KafkaRestPropertiesResponse) map[string]string { return v.ConfigurationOverride }).(pulumi.StringMapOutput)
}

type KafkaRestPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KafkaRestPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaRestPropertiesResponse)(nil)).Elem()
}

func (o KafkaRestPropertiesResponsePtrOutput) ToKafkaRestPropertiesResponsePtrOutput() KafkaRestPropertiesResponsePtrOutput {
	return o
}

func (o KafkaRestPropertiesResponsePtrOutput) ToKafkaRestPropertiesResponsePtrOutputWithContext(ctx context.Context) KafkaRestPropertiesResponsePtrOutput {
	return o
}

func (o KafkaRestPropertiesResponsePtrOutput) Elem() KafkaRestPropertiesResponseOutput {
	return o.ApplyT(func(v *KafkaRestPropertiesResponse) KafkaRestPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KafkaRestPropertiesResponse
		return ret
	}).(KafkaRestPropertiesResponseOutput)
}

func (o KafkaRestPropertiesResponsePtrOutput) ClientGroupInfo() ClientGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v *KafkaRestPropertiesResponse) *ClientGroupInfoResponse {
		if v == nil {
			return nil
		}
		return v.ClientGroupInfo
	}).(ClientGroupInfoResponsePtrOutput)
}

func (o KafkaRestPropertiesResponsePtrOutput) ConfigurationOverride() pulumi.StringMapOutput {
	return o.ApplyT(func(v *KafkaRestPropertiesResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.ConfigurationOverride
	}).(pulumi.StringMapOutput)
}

type LinuxOperatingSystemProfile struct {
	Password   *string     `pulumi:"password"`
	SshProfile *SshProfile `pulumi:"sshProfile"`
	Username   *string     `pulumi:"username"`
}





type LinuxOperatingSystemProfileInput interface {
	pulumi.Input

	ToLinuxOperatingSystemProfileOutput() LinuxOperatingSystemProfileOutput
	ToLinuxOperatingSystemProfileOutputWithContext(context.Context) LinuxOperatingSystemProfileOutput
}

type LinuxOperatingSystemProfileArgs struct {
	Password   pulumi.StringPtrInput `pulumi:"password"`
	SshProfile SshProfilePtrInput    `pulumi:"sshProfile"`
	Username   pulumi.StringPtrInput `pulumi:"username"`
}

func (LinuxOperatingSystemProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOperatingSystemProfile)(nil)).Elem()
}

func (i LinuxOperatingSystemProfileArgs) ToLinuxOperatingSystemProfileOutput() LinuxOperatingSystemProfileOutput {
	return i.ToLinuxOperatingSystemProfileOutputWithContext(context.Background())
}

func (i LinuxOperatingSystemProfileArgs) ToLinuxOperatingSystemProfileOutputWithContext(ctx context.Context) LinuxOperatingSystemProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOperatingSystemProfileOutput)
}

func (i LinuxOperatingSystemProfileArgs) ToLinuxOperatingSystemProfilePtrOutput() LinuxOperatingSystemProfilePtrOutput {
	return i.ToLinuxOperatingSystemProfilePtrOutputWithContext(context.Background())
}

func (i LinuxOperatingSystemProfileArgs) ToLinuxOperatingSystemProfilePtrOutputWithContext(ctx context.Context) LinuxOperatingSystemProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOperatingSystemProfileOutput).ToLinuxOperatingSystemProfilePtrOutputWithContext(ctx)
}









type LinuxOperatingSystemProfilePtrInput interface {
	pulumi.Input

	ToLinuxOperatingSystemProfilePtrOutput() LinuxOperatingSystemProfilePtrOutput
	ToLinuxOperatingSystemProfilePtrOutputWithContext(context.Context) LinuxOperatingSystemProfilePtrOutput
}

type linuxOperatingSystemProfilePtrType LinuxOperatingSystemProfileArgs

func LinuxOperatingSystemProfilePtr(v *LinuxOperatingSystemProfileArgs) LinuxOperatingSystemProfilePtrInput {
	return (*linuxOperatingSystemProfilePtrType)(v)
}

func (*linuxOperatingSystemProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOperatingSystemProfile)(nil)).Elem()
}

func (i *linuxOperatingSystemProfilePtrType) ToLinuxOperatingSystemProfilePtrOutput() LinuxOperatingSystemProfilePtrOutput {
	return i.ToLinuxOperatingSystemProfilePtrOutputWithContext(context.Background())
}

func (i *linuxOperatingSystemProfilePtrType) ToLinuxOperatingSystemProfilePtrOutputWithContext(ctx context.Context) LinuxOperatingSystemProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOperatingSystemProfilePtrOutput)
}

type LinuxOperatingSystemProfileOutput struct{ *pulumi.OutputState }

func (LinuxOperatingSystemProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOperatingSystemProfile)(nil)).Elem()
}

func (o LinuxOperatingSystemProfileOutput) ToLinuxOperatingSystemProfileOutput() LinuxOperatingSystemProfileOutput {
	return o
}

func (o LinuxOperatingSystemProfileOutput) ToLinuxOperatingSystemProfileOutputWithContext(ctx context.Context) LinuxOperatingSystemProfileOutput {
	return o
}

func (o LinuxOperatingSystemProfileOutput) ToLinuxOperatingSystemProfilePtrOutput() LinuxOperatingSystemProfilePtrOutput {
	return o.ToLinuxOperatingSystemProfilePtrOutputWithContext(context.Background())
}

func (o LinuxOperatingSystemProfileOutput) ToLinuxOperatingSystemProfilePtrOutputWithContext(ctx context.Context) LinuxOperatingSystemProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxOperatingSystemProfile) *LinuxOperatingSystemProfile {
		return &v
	}).(LinuxOperatingSystemProfilePtrOutput)
}

func (o LinuxOperatingSystemProfileOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxOperatingSystemProfile) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LinuxOperatingSystemProfileOutput) SshProfile() SshProfilePtrOutput {
	return o.ApplyT(func(v LinuxOperatingSystemProfile) *SshProfile { return v.SshProfile }).(SshProfilePtrOutput)
}

func (o LinuxOperatingSystemProfileOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxOperatingSystemProfile) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type LinuxOperatingSystemProfilePtrOutput struct{ *pulumi.OutputState }

func (LinuxOperatingSystemProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOperatingSystemProfile)(nil)).Elem()
}

func (o LinuxOperatingSystemProfilePtrOutput) ToLinuxOperatingSystemProfilePtrOutput() LinuxOperatingSystemProfilePtrOutput {
	return o
}

func (o LinuxOperatingSystemProfilePtrOutput) ToLinuxOperatingSystemProfilePtrOutputWithContext(ctx context.Context) LinuxOperatingSystemProfilePtrOutput {
	return o
}

func (o LinuxOperatingSystemProfilePtrOutput) Elem() LinuxOperatingSystemProfileOutput {
	return o.ApplyT(func(v *LinuxOperatingSystemProfile) LinuxOperatingSystemProfile {
		if v != nil {
			return *v
		}
		var ret LinuxOperatingSystemProfile
		return ret
	}).(LinuxOperatingSystemProfileOutput)
}

func (o LinuxOperatingSystemProfilePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxOperatingSystemProfile) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o LinuxOperatingSystemProfilePtrOutput) SshProfile() SshProfilePtrOutput {
	return o.ApplyT(func(v *LinuxOperatingSystemProfile) *SshProfile {
		if v == nil {
			return nil
		}
		return v.SshProfile
	}).(SshProfilePtrOutput)
}

func (o LinuxOperatingSystemProfilePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxOperatingSystemProfile) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type LinuxOperatingSystemProfileResponse struct {
	Password   *string             `pulumi:"password"`
	SshProfile *SshProfileResponse `pulumi:"sshProfile"`
	Username   *string             `pulumi:"username"`
}

type LinuxOperatingSystemProfileResponseOutput struct{ *pulumi.OutputState }

func (LinuxOperatingSystemProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOperatingSystemProfileResponse)(nil)).Elem()
}

func (o LinuxOperatingSystemProfileResponseOutput) ToLinuxOperatingSystemProfileResponseOutput() LinuxOperatingSystemProfileResponseOutput {
	return o
}

func (o LinuxOperatingSystemProfileResponseOutput) ToLinuxOperatingSystemProfileResponseOutputWithContext(ctx context.Context) LinuxOperatingSystemProfileResponseOutput {
	return o
}

func (o LinuxOperatingSystemProfileResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxOperatingSystemProfileResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LinuxOperatingSystemProfileResponseOutput) SshProfile() SshProfileResponsePtrOutput {
	return o.ApplyT(func(v LinuxOperatingSystemProfileResponse) *SshProfileResponse { return v.SshProfile }).(SshProfileResponsePtrOutput)
}

func (o LinuxOperatingSystemProfileResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxOperatingSystemProfileResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type LinuxOperatingSystemProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxOperatingSystemProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOperatingSystemProfileResponse)(nil)).Elem()
}

func (o LinuxOperatingSystemProfileResponsePtrOutput) ToLinuxOperatingSystemProfileResponsePtrOutput() LinuxOperatingSystemProfileResponsePtrOutput {
	return o
}

func (o LinuxOperatingSystemProfileResponsePtrOutput) ToLinuxOperatingSystemProfileResponsePtrOutputWithContext(ctx context.Context) LinuxOperatingSystemProfileResponsePtrOutput {
	return o
}

func (o LinuxOperatingSystemProfileResponsePtrOutput) Elem() LinuxOperatingSystemProfileResponseOutput {
	return o.ApplyT(func(v *LinuxOperatingSystemProfileResponse) LinuxOperatingSystemProfileResponse {
		if v != nil {
			return *v
		}
		var ret LinuxOperatingSystemProfileResponse
		return ret
	}).(LinuxOperatingSystemProfileResponseOutput)
}

func (o LinuxOperatingSystemProfileResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxOperatingSystemProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o LinuxOperatingSystemProfileResponsePtrOutput) SshProfile() SshProfileResponsePtrOutput {
	return o.ApplyT(func(v *LinuxOperatingSystemProfileResponse) *SshProfileResponse {
		if v == nil {
			return nil
		}
		return v.SshProfile
	}).(SshProfileResponsePtrOutput)
}

func (o LinuxOperatingSystemProfileResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxOperatingSystemProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type NetworkProperties struct {
	PrivateLink                *string `pulumi:"privateLink"`
	ResourceProviderConnection *string `pulumi:"resourceProviderConnection"`
}





type NetworkPropertiesInput interface {
	pulumi.Input

	ToNetworkPropertiesOutput() NetworkPropertiesOutput
	ToNetworkPropertiesOutputWithContext(context.Context) NetworkPropertiesOutput
}

type NetworkPropertiesArgs struct {
	PrivateLink                pulumi.StringPtrInput `pulumi:"privateLink"`
	ResourceProviderConnection pulumi.StringPtrInput `pulumi:"resourceProviderConnection"`
}

func (NetworkPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProperties)(nil)).Elem()
}

func (i NetworkPropertiesArgs) ToNetworkPropertiesOutput() NetworkPropertiesOutput {
	return i.ToNetworkPropertiesOutputWithContext(context.Background())
}

func (i NetworkPropertiesArgs) ToNetworkPropertiesOutputWithContext(ctx context.Context) NetworkPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPropertiesOutput)
}

func (i NetworkPropertiesArgs) ToNetworkPropertiesPtrOutput() NetworkPropertiesPtrOutput {
	return i.ToNetworkPropertiesPtrOutputWithContext(context.Background())
}

func (i NetworkPropertiesArgs) ToNetworkPropertiesPtrOutputWithContext(ctx context.Context) NetworkPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPropertiesOutput).ToNetworkPropertiesPtrOutputWithContext(ctx)
}









type NetworkPropertiesPtrInput interface {
	pulumi.Input

	ToNetworkPropertiesPtrOutput() NetworkPropertiesPtrOutput
	ToNetworkPropertiesPtrOutputWithContext(context.Context) NetworkPropertiesPtrOutput
}

type networkPropertiesPtrType NetworkPropertiesArgs

func NetworkPropertiesPtr(v *NetworkPropertiesArgs) NetworkPropertiesPtrInput {
	return (*networkPropertiesPtrType)(v)
}

func (*networkPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProperties)(nil)).Elem()
}

func (i *networkPropertiesPtrType) ToNetworkPropertiesPtrOutput() NetworkPropertiesPtrOutput {
	return i.ToNetworkPropertiesPtrOutputWithContext(context.Background())
}

func (i *networkPropertiesPtrType) ToNetworkPropertiesPtrOutputWithContext(ctx context.Context) NetworkPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPropertiesPtrOutput)
}

type NetworkPropertiesOutput struct{ *pulumi.OutputState }

func (NetworkPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProperties)(nil)).Elem()
}

func (o NetworkPropertiesOutput) ToNetworkPropertiesOutput() NetworkPropertiesOutput {
	return o
}

func (o NetworkPropertiesOutput) ToNetworkPropertiesOutputWithContext(ctx context.Context) NetworkPropertiesOutput {
	return o
}

func (o NetworkPropertiesOutput) ToNetworkPropertiesPtrOutput() NetworkPropertiesPtrOutput {
	return o.ToNetworkPropertiesPtrOutputWithContext(context.Background())
}

func (o NetworkPropertiesOutput) ToNetworkPropertiesPtrOutputWithContext(ctx context.Context) NetworkPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProperties) *NetworkProperties {
		return &v
	}).(NetworkPropertiesPtrOutput)
}

func (o NetworkPropertiesOutput) PrivateLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProperties) *string { return v.PrivateLink }).(pulumi.StringPtrOutput)
}

func (o NetworkPropertiesOutput) ResourceProviderConnection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProperties) *string { return v.ResourceProviderConnection }).(pulumi.StringPtrOutput)
}

type NetworkPropertiesPtrOutput struct{ *pulumi.OutputState }

func (NetworkPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProperties)(nil)).Elem()
}

func (o NetworkPropertiesPtrOutput) ToNetworkPropertiesPtrOutput() NetworkPropertiesPtrOutput {
	return o
}

func (o NetworkPropertiesPtrOutput) ToNetworkPropertiesPtrOutputWithContext(ctx context.Context) NetworkPropertiesPtrOutput {
	return o
}

func (o NetworkPropertiesPtrOutput) Elem() NetworkPropertiesOutput {
	return o.ApplyT(func(v *NetworkProperties) NetworkProperties {
		if v != nil {
			return *v
		}
		var ret NetworkProperties
		return ret
	}).(NetworkPropertiesOutput)
}

func (o NetworkPropertiesPtrOutput) PrivateLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateLink
	}).(pulumi.StringPtrOutput)
}

func (o NetworkPropertiesPtrOutput) ResourceProviderConnection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResourceProviderConnection
	}).(pulumi.StringPtrOutput)
}

type NetworkPropertiesResponse struct {
	PrivateLink                *string `pulumi:"privateLink"`
	ResourceProviderConnection *string `pulumi:"resourceProviderConnection"`
}

type NetworkPropertiesResponseOutput struct{ *pulumi.OutputState }

func (NetworkPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPropertiesResponse)(nil)).Elem()
}

func (o NetworkPropertiesResponseOutput) ToNetworkPropertiesResponseOutput() NetworkPropertiesResponseOutput {
	return o
}

func (o NetworkPropertiesResponseOutput) ToNetworkPropertiesResponseOutputWithContext(ctx context.Context) NetworkPropertiesResponseOutput {
	return o
}

func (o NetworkPropertiesResponseOutput) PrivateLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkPropertiesResponse) *string { return v.PrivateLink }).(pulumi.StringPtrOutput)
}

func (o NetworkPropertiesResponseOutput) ResourceProviderConnection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkPropertiesResponse) *string { return v.ResourceProviderConnection }).(pulumi.StringPtrOutput)
}

type NetworkPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPropertiesResponse)(nil)).Elem()
}

func (o NetworkPropertiesResponsePtrOutput) ToNetworkPropertiesResponsePtrOutput() NetworkPropertiesResponsePtrOutput {
	return o
}

func (o NetworkPropertiesResponsePtrOutput) ToNetworkPropertiesResponsePtrOutputWithContext(ctx context.Context) NetworkPropertiesResponsePtrOutput {
	return o
}

func (o NetworkPropertiesResponsePtrOutput) Elem() NetworkPropertiesResponseOutput {
	return o.ApplyT(func(v *NetworkPropertiesResponse) NetworkPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret NetworkPropertiesResponse
		return ret
	}).(NetworkPropertiesResponseOutput)
}

func (o NetworkPropertiesResponsePtrOutput) PrivateLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateLink
	}).(pulumi.StringPtrOutput)
}

func (o NetworkPropertiesResponsePtrOutput) ResourceProviderConnection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceProviderConnection
	}).(pulumi.StringPtrOutput)
}

type OsProfile struct {
	LinuxOperatingSystemProfile *LinuxOperatingSystemProfile `pulumi:"linuxOperatingSystemProfile"`
}





type OsProfileInput interface {
	pulumi.Input

	ToOsProfileOutput() OsProfileOutput
	ToOsProfileOutputWithContext(context.Context) OsProfileOutput
}

type OsProfileArgs struct {
	LinuxOperatingSystemProfile LinuxOperatingSystemProfilePtrInput `pulumi:"linuxOperatingSystemProfile"`
}

func (OsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfile)(nil)).Elem()
}

func (i OsProfileArgs) ToOsProfileOutput() OsProfileOutput {
	return i.ToOsProfileOutputWithContext(context.Background())
}

func (i OsProfileArgs) ToOsProfileOutputWithContext(ctx context.Context) OsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfileOutput)
}

func (i OsProfileArgs) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return i.ToOsProfilePtrOutputWithContext(context.Background())
}

func (i OsProfileArgs) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfileOutput).ToOsProfilePtrOutputWithContext(ctx)
}









type OsProfilePtrInput interface {
	pulumi.Input

	ToOsProfilePtrOutput() OsProfilePtrOutput
	ToOsProfilePtrOutputWithContext(context.Context) OsProfilePtrOutput
}

type osProfilePtrType OsProfileArgs

func OsProfilePtr(v *OsProfileArgs) OsProfilePtrInput {
	return (*osProfilePtrType)(v)
}

func (*osProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfile)(nil)).Elem()
}

func (i *osProfilePtrType) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return i.ToOsProfilePtrOutputWithContext(context.Background())
}

func (i *osProfilePtrType) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfilePtrOutput)
}

type OsProfileOutput struct{ *pulumi.OutputState }

func (OsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfile)(nil)).Elem()
}

func (o OsProfileOutput) ToOsProfileOutput() OsProfileOutput {
	return o
}

func (o OsProfileOutput) ToOsProfileOutputWithContext(ctx context.Context) OsProfileOutput {
	return o
}

func (o OsProfileOutput) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return o.ToOsProfilePtrOutputWithContext(context.Background())
}

func (o OsProfileOutput) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OsProfile) *OsProfile {
		return &v
	}).(OsProfilePtrOutput)
}

func (o OsProfileOutput) LinuxOperatingSystemProfile() LinuxOperatingSystemProfilePtrOutput {
	return o.ApplyT(func(v OsProfile) *LinuxOperatingSystemProfile { return v.LinuxOperatingSystemProfile }).(LinuxOperatingSystemProfilePtrOutput)
}

type OsProfilePtrOutput struct{ *pulumi.OutputState }

func (OsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfile)(nil)).Elem()
}

func (o OsProfilePtrOutput) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return o
}

func (o OsProfilePtrOutput) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return o
}

func (o OsProfilePtrOutput) Elem() OsProfileOutput {
	return o.ApplyT(func(v *OsProfile) OsProfile {
		if v != nil {
			return *v
		}
		var ret OsProfile
		return ret
	}).(OsProfileOutput)
}

func (o OsProfilePtrOutput) LinuxOperatingSystemProfile() LinuxOperatingSystemProfilePtrOutput {
	return o.ApplyT(func(v *OsProfile) *LinuxOperatingSystemProfile {
		if v == nil {
			return nil
		}
		return v.LinuxOperatingSystemProfile
	}).(LinuxOperatingSystemProfilePtrOutput)
}

type OsProfileResponse struct {
	LinuxOperatingSystemProfile *LinuxOperatingSystemProfileResponse `pulumi:"linuxOperatingSystemProfile"`
}

type OsProfileResponseOutput struct{ *pulumi.OutputState }

func (OsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfileResponse)(nil)).Elem()
}

func (o OsProfileResponseOutput) ToOsProfileResponseOutput() OsProfileResponseOutput {
	return o
}

func (o OsProfileResponseOutput) ToOsProfileResponseOutputWithContext(ctx context.Context) OsProfileResponseOutput {
	return o
}

func (o OsProfileResponseOutput) LinuxOperatingSystemProfile() LinuxOperatingSystemProfileResponsePtrOutput {
	return o.ApplyT(func(v OsProfileResponse) *LinuxOperatingSystemProfileResponse { return v.LinuxOperatingSystemProfile }).(LinuxOperatingSystemProfileResponsePtrOutput)
}

type OsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfileResponse)(nil)).Elem()
}

func (o OsProfileResponsePtrOutput) ToOsProfileResponsePtrOutput() OsProfileResponsePtrOutput {
	return o
}

func (o OsProfileResponsePtrOutput) ToOsProfileResponsePtrOutputWithContext(ctx context.Context) OsProfileResponsePtrOutput {
	return o
}

func (o OsProfileResponsePtrOutput) Elem() OsProfileResponseOutput {
	return o.ApplyT(func(v *OsProfileResponse) OsProfileResponse {
		if v != nil {
			return *v
		}
		var ret OsProfileResponse
		return ret
	}).(OsProfileResponseOutput)
}

func (o OsProfileResponsePtrOutput) LinuxOperatingSystemProfile() LinuxOperatingSystemProfileResponsePtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *LinuxOperatingSystemProfileResponse {
		if v == nil {
			return nil
		}
		return v.LinuxOperatingSystemProfile
	}).(LinuxOperatingSystemProfileResponsePtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	LinkIdentifier                    string                                    `pulumi:"linkIdentifier"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
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

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) LinkIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.LinkIdentifier }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponseOutput)
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
	Id *string `pulumi:"id"`
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

func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateLinkConfiguration struct {
	GroupId          string            `pulumi:"groupId"`
	IpConfigurations []IPConfiguration `pulumi:"ipConfigurations"`
	Name             string            `pulumi:"name"`
}





type PrivateLinkConfigurationInput interface {
	pulumi.Input

	ToPrivateLinkConfigurationOutput() PrivateLinkConfigurationOutput
	ToPrivateLinkConfigurationOutputWithContext(context.Context) PrivateLinkConfigurationOutput
}

type PrivateLinkConfigurationArgs struct {
	GroupId          pulumi.StringInput        `pulumi:"groupId"`
	IpConfigurations IPConfigurationArrayInput `pulumi:"ipConfigurations"`
	Name             pulumi.StringInput        `pulumi:"name"`
}

func (PrivateLinkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConfiguration)(nil)).Elem()
}

func (i PrivateLinkConfigurationArgs) ToPrivateLinkConfigurationOutput() PrivateLinkConfigurationOutput {
	return i.ToPrivateLinkConfigurationOutputWithContext(context.Background())
}

func (i PrivateLinkConfigurationArgs) ToPrivateLinkConfigurationOutputWithContext(ctx context.Context) PrivateLinkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConfigurationOutput)
}





type PrivateLinkConfigurationArrayInput interface {
	pulumi.Input

	ToPrivateLinkConfigurationArrayOutput() PrivateLinkConfigurationArrayOutput
	ToPrivateLinkConfigurationArrayOutputWithContext(context.Context) PrivateLinkConfigurationArrayOutput
}

type PrivateLinkConfigurationArray []PrivateLinkConfigurationInput

func (PrivateLinkConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkConfiguration)(nil)).Elem()
}

func (i PrivateLinkConfigurationArray) ToPrivateLinkConfigurationArrayOutput() PrivateLinkConfigurationArrayOutput {
	return i.ToPrivateLinkConfigurationArrayOutputWithContext(context.Background())
}

func (i PrivateLinkConfigurationArray) ToPrivateLinkConfigurationArrayOutputWithContext(ctx context.Context) PrivateLinkConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConfigurationArrayOutput)
}

type PrivateLinkConfigurationOutput struct{ *pulumi.OutputState }

func (PrivateLinkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConfiguration)(nil)).Elem()
}

func (o PrivateLinkConfigurationOutput) ToPrivateLinkConfigurationOutput() PrivateLinkConfigurationOutput {
	return o
}

func (o PrivateLinkConfigurationOutput) ToPrivateLinkConfigurationOutputWithContext(ctx context.Context) PrivateLinkConfigurationOutput {
	return o
}

func (o PrivateLinkConfigurationOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConfiguration) string { return v.GroupId }).(pulumi.StringOutput)
}

func (o PrivateLinkConfigurationOutput) IpConfigurations() IPConfigurationArrayOutput {
	return o.ApplyT(func(v PrivateLinkConfiguration) []IPConfiguration { return v.IpConfigurations }).(IPConfigurationArrayOutput)
}

func (o PrivateLinkConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

type PrivateLinkConfigurationArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkConfiguration)(nil)).Elem()
}

func (o PrivateLinkConfigurationArrayOutput) ToPrivateLinkConfigurationArrayOutput() PrivateLinkConfigurationArrayOutput {
	return o
}

func (o PrivateLinkConfigurationArrayOutput) ToPrivateLinkConfigurationArrayOutputWithContext(ctx context.Context) PrivateLinkConfigurationArrayOutput {
	return o
}

func (o PrivateLinkConfigurationArrayOutput) Index(i pulumi.IntInput) PrivateLinkConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkConfiguration {
		return vs[0].([]PrivateLinkConfiguration)[vs[1].(int)]
	}).(PrivateLinkConfigurationOutput)
}

type PrivateLinkConfigurationResponse struct {
	GroupId           string                    `pulumi:"groupId"`
	Id                string                    `pulumi:"id"`
	IpConfigurations  []IPConfigurationResponse `pulumi:"ipConfigurations"`
	Name              string                    `pulumi:"name"`
	ProvisioningState string                    `pulumi:"provisioningState"`
	Type              string                    `pulumi:"type"`
}

type PrivateLinkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConfigurationResponse)(nil)).Elem()
}

func (o PrivateLinkConfigurationResponseOutput) ToPrivateLinkConfigurationResponseOutput() PrivateLinkConfigurationResponseOutput {
	return o
}

func (o PrivateLinkConfigurationResponseOutput) ToPrivateLinkConfigurationResponseOutputWithContext(ctx context.Context) PrivateLinkConfigurationResponseOutput {
	return o
}

func (o PrivateLinkConfigurationResponseOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConfigurationResponse) string { return v.GroupId }).(pulumi.StringOutput)
}

func (o PrivateLinkConfigurationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConfigurationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateLinkConfigurationResponseOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v PrivateLinkConfigurationResponse) []IPConfigurationResponse { return v.IpConfigurations }).(IPConfigurationResponseArrayOutput)
}

func (o PrivateLinkConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateLinkConfigurationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConfigurationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateLinkConfigurationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConfigurationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateLinkConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkConfigurationResponse)(nil)).Elem()
}

func (o PrivateLinkConfigurationResponseArrayOutput) ToPrivateLinkConfigurationResponseArrayOutput() PrivateLinkConfigurationResponseArrayOutput {
	return o
}

func (o PrivateLinkConfigurationResponseArrayOutput) ToPrivateLinkConfigurationResponseArrayOutputWithContext(ctx context.Context) PrivateLinkConfigurationResponseArrayOutput {
	return o
}

func (o PrivateLinkConfigurationResponseArrayOutput) Index(i pulumi.IntInput) PrivateLinkConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkConfigurationResponse {
		return vs[0].([]PrivateLinkConfigurationResponse)[vs[1].(int)]
	}).(PrivateLinkConfigurationResponseOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          string  `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringInput    `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          string  `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) string { return v.Status }).(pulumi.StringOutput)
}

type QuotaInfoResponse struct {
	CoresUsed *int `pulumi:"coresUsed"`
}

type QuotaInfoResponseOutput struct{ *pulumi.OutputState }

func (QuotaInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuotaInfoResponse)(nil)).Elem()
}

func (o QuotaInfoResponseOutput) ToQuotaInfoResponseOutput() QuotaInfoResponseOutput {
	return o
}

func (o QuotaInfoResponseOutput) ToQuotaInfoResponseOutputWithContext(ctx context.Context) QuotaInfoResponseOutput {
	return o
}

func (o QuotaInfoResponseOutput) CoresUsed() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QuotaInfoResponse) *int { return v.CoresUsed }).(pulumi.IntPtrOutput)
}

type QuotaInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (QuotaInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaInfoResponse)(nil)).Elem()
}

func (o QuotaInfoResponsePtrOutput) ToQuotaInfoResponsePtrOutput() QuotaInfoResponsePtrOutput {
	return o
}

func (o QuotaInfoResponsePtrOutput) ToQuotaInfoResponsePtrOutputWithContext(ctx context.Context) QuotaInfoResponsePtrOutput {
	return o
}

func (o QuotaInfoResponsePtrOutput) Elem() QuotaInfoResponseOutput {
	return o.ApplyT(func(v *QuotaInfoResponse) QuotaInfoResponse {
		if v != nil {
			return *v
		}
		var ret QuotaInfoResponse
		return ret
	}).(QuotaInfoResponseOutput)
}

func (o QuotaInfoResponsePtrOutput) CoresUsed() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QuotaInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.CoresUsed
	}).(pulumi.IntPtrOutput)
}

type ResourceId struct {
	Id *string `pulumi:"id"`
}





type ResourceIdInput interface {
	pulumi.Input

	ToResourceIdOutput() ResourceIdOutput
	ToResourceIdOutputWithContext(context.Context) ResourceIdOutput
}

type ResourceIdArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceId)(nil)).Elem()
}

func (i ResourceIdArgs) ToResourceIdOutput() ResourceIdOutput {
	return i.ToResourceIdOutputWithContext(context.Background())
}

func (i ResourceIdArgs) ToResourceIdOutputWithContext(ctx context.Context) ResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdOutput)
}

func (i ResourceIdArgs) ToResourceIdPtrOutput() ResourceIdPtrOutput {
	return i.ToResourceIdPtrOutputWithContext(context.Background())
}

func (i ResourceIdArgs) ToResourceIdPtrOutputWithContext(ctx context.Context) ResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdOutput).ToResourceIdPtrOutputWithContext(ctx)
}









type ResourceIdPtrInput interface {
	pulumi.Input

	ToResourceIdPtrOutput() ResourceIdPtrOutput
	ToResourceIdPtrOutputWithContext(context.Context) ResourceIdPtrOutput
}

type resourceIdPtrType ResourceIdArgs

func ResourceIdPtr(v *ResourceIdArgs) ResourceIdPtrInput {
	return (*resourceIdPtrType)(v)
}

func (*resourceIdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceId)(nil)).Elem()
}

func (i *resourceIdPtrType) ToResourceIdPtrOutput() ResourceIdPtrOutput {
	return i.ToResourceIdPtrOutputWithContext(context.Background())
}

func (i *resourceIdPtrType) ToResourceIdPtrOutputWithContext(ctx context.Context) ResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdPtrOutput)
}

type ResourceIdOutput struct{ *pulumi.OutputState }

func (ResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceId)(nil)).Elem()
}

func (o ResourceIdOutput) ToResourceIdOutput() ResourceIdOutput {
	return o
}

func (o ResourceIdOutput) ToResourceIdOutputWithContext(ctx context.Context) ResourceIdOutput {
	return o
}

func (o ResourceIdOutput) ToResourceIdPtrOutput() ResourceIdPtrOutput {
	return o.ToResourceIdPtrOutputWithContext(context.Background())
}

func (o ResourceIdOutput) ToResourceIdPtrOutputWithContext(ctx context.Context) ResourceIdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceId) *ResourceId {
		return &v
	}).(ResourceIdPtrOutput)
}

func (o ResourceIdOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceId) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceIdPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceId)(nil)).Elem()
}

func (o ResourceIdPtrOutput) ToResourceIdPtrOutput() ResourceIdPtrOutput {
	return o
}

func (o ResourceIdPtrOutput) ToResourceIdPtrOutputWithContext(ctx context.Context) ResourceIdPtrOutput {
	return o
}

func (o ResourceIdPtrOutput) Elem() ResourceIdOutput {
	return o.ApplyT(func(v *ResourceId) ResourceId {
		if v != nil {
			return *v
		}
		var ret ResourceId
		return ret
	}).(ResourceIdOutput)
}

func (o ResourceIdPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceId) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ResourceIdResponse struct {
	Id *string `pulumi:"id"`
}

type ResourceIdResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdResponse)(nil)).Elem()
}

func (o ResourceIdResponseOutput) ToResourceIdResponseOutput() ResourceIdResponseOutput {
	return o
}

func (o ResourceIdResponseOutput) ToResourceIdResponseOutputWithContext(ctx context.Context) ResourceIdResponseOutput {
	return o
}

func (o ResourceIdResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceIdResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdResponse)(nil)).Elem()
}

func (o ResourceIdResponsePtrOutput) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return o
}

func (o ResourceIdResponsePtrOutput) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return o
}

func (o ResourceIdResponsePtrOutput) Elem() ResourceIdResponseOutput {
	return o.ApplyT(func(v *ResourceIdResponse) ResourceIdResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdResponse
		return ret
	}).(ResourceIdResponseOutput)
}

func (o ResourceIdResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type Role struct {
	AutoscaleConfiguration *Autoscale             `pulumi:"autoscaleConfiguration"`
	DataDisksGroups        []DataDisksGroups      `pulumi:"dataDisksGroups"`
	EncryptDataDisks       *bool                  `pulumi:"encryptDataDisks"`
	HardwareProfile        *HardwareProfile       `pulumi:"hardwareProfile"`
	MinInstanceCount       *int                   `pulumi:"minInstanceCount"`
	Name                   *string                `pulumi:"name"`
	OsProfile              *OsProfile             `pulumi:"osProfile"`
	ScriptActions          []ScriptAction         `pulumi:"scriptActions"`
	TargetInstanceCount    *int                   `pulumi:"targetInstanceCount"`
	VMGroupName            *string                `pulumi:"vMGroupName"`
	VirtualNetworkProfile  *VirtualNetworkProfile `pulumi:"virtualNetworkProfile"`
}


func (val *Role) Defaults() *Role {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EncryptDataDisks) {
		encryptDataDisks_ := false
		tmp.EncryptDataDisks = &encryptDataDisks_
	}
	return &tmp
}





type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(context.Context) RoleOutput
}

type RoleArgs struct {
	AutoscaleConfiguration AutoscalePtrInput             `pulumi:"autoscaleConfiguration"`
	DataDisksGroups        DataDisksGroupsArrayInput     `pulumi:"dataDisksGroups"`
	EncryptDataDisks       pulumi.BoolPtrInput           `pulumi:"encryptDataDisks"`
	HardwareProfile        HardwareProfilePtrInput       `pulumi:"hardwareProfile"`
	MinInstanceCount       pulumi.IntPtrInput            `pulumi:"minInstanceCount"`
	Name                   pulumi.StringPtrInput         `pulumi:"name"`
	OsProfile              OsProfilePtrInput             `pulumi:"osProfile"`
	ScriptActions          ScriptActionArrayInput        `pulumi:"scriptActions"`
	TargetInstanceCount    pulumi.IntPtrInput            `pulumi:"targetInstanceCount"`
	VMGroupName            pulumi.StringPtrInput         `pulumi:"vMGroupName"`
	VirtualNetworkProfile  VirtualNetworkProfilePtrInput `pulumi:"virtualNetworkProfile"`
}


func (val *RoleArgs) Defaults() *RoleArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EncryptDataDisks) {
		tmp.EncryptDataDisks = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Role)(nil)).Elem()
}

func (i RoleArgs) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i RoleArgs) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}





type RoleArrayInput interface {
	pulumi.Input

	ToRoleArrayOutput() RoleArrayOutput
	ToRoleArrayOutputWithContext(context.Context) RoleArrayOutput
}

type RoleArray []RoleInput

func (RoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Role)(nil)).Elem()
}

func (i RoleArray) ToRoleArrayOutput() RoleArrayOutput {
	return i.ToRoleArrayOutputWithContext(context.Background())
}

func (i RoleArray) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleArrayOutput)
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Role)(nil)).Elem()
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

func (o RoleOutput) AutoscaleConfiguration() AutoscalePtrOutput {
	return o.ApplyT(func(v Role) *Autoscale { return v.AutoscaleConfiguration }).(AutoscalePtrOutput)
}

func (o RoleOutput) DataDisksGroups() DataDisksGroupsArrayOutput {
	return o.ApplyT(func(v Role) []DataDisksGroups { return v.DataDisksGroups }).(DataDisksGroupsArrayOutput)
}

func (o RoleOutput) EncryptDataDisks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Role) *bool { return v.EncryptDataDisks }).(pulumi.BoolPtrOutput)
}

func (o RoleOutput) HardwareProfile() HardwareProfilePtrOutput {
	return o.ApplyT(func(v Role) *HardwareProfile { return v.HardwareProfile }).(HardwareProfilePtrOutput)
}

func (o RoleOutput) MinInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Role) *int { return v.MinInstanceCount }).(pulumi.IntPtrOutput)
}

func (o RoleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Role) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RoleOutput) OsProfile() OsProfilePtrOutput {
	return o.ApplyT(func(v Role) *OsProfile { return v.OsProfile }).(OsProfilePtrOutput)
}

func (o RoleOutput) ScriptActions() ScriptActionArrayOutput {
	return o.ApplyT(func(v Role) []ScriptAction { return v.ScriptActions }).(ScriptActionArrayOutput)
}

func (o RoleOutput) TargetInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Role) *int { return v.TargetInstanceCount }).(pulumi.IntPtrOutput)
}

func (o RoleOutput) VMGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Role) *string { return v.VMGroupName }).(pulumi.StringPtrOutput)
}

func (o RoleOutput) VirtualNetworkProfile() VirtualNetworkProfilePtrOutput {
	return o.ApplyT(func(v Role) *VirtualNetworkProfile { return v.VirtualNetworkProfile }).(VirtualNetworkProfilePtrOutput)
}

type RoleArrayOutput struct{ *pulumi.OutputState }

func (RoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Role)(nil)).Elem()
}

func (o RoleArrayOutput) ToRoleArrayOutput() RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) Index(i pulumi.IntInput) RoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Role {
		return vs[0].([]Role)[vs[1].(int)]
	}).(RoleOutput)
}

type RoleResponse struct {
	AutoscaleConfiguration *AutoscaleResponse             `pulumi:"autoscaleConfiguration"`
	DataDisksGroups        []DataDisksGroupsResponse      `pulumi:"dataDisksGroups"`
	EncryptDataDisks       *bool                          `pulumi:"encryptDataDisks"`
	HardwareProfile        *HardwareProfileResponse       `pulumi:"hardwareProfile"`
	MinInstanceCount       *int                           `pulumi:"minInstanceCount"`
	Name                   *string                        `pulumi:"name"`
	OsProfile              *OsProfileResponse             `pulumi:"osProfile"`
	ScriptActions          []ScriptActionResponse         `pulumi:"scriptActions"`
	TargetInstanceCount    *int                           `pulumi:"targetInstanceCount"`
	VMGroupName            *string                        `pulumi:"vMGroupName"`
	VirtualNetworkProfile  *VirtualNetworkProfileResponse `pulumi:"virtualNetworkProfile"`
}


func (val *RoleResponse) Defaults() *RoleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EncryptDataDisks) {
		encryptDataDisks_ := false
		tmp.EncryptDataDisks = &encryptDataDisks_
	}
	return &tmp
}

type RoleResponseOutput struct{ *pulumi.OutputState }

func (RoleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleResponse)(nil)).Elem()
}

func (o RoleResponseOutput) ToRoleResponseOutput() RoleResponseOutput {
	return o
}

func (o RoleResponseOutput) ToRoleResponseOutputWithContext(ctx context.Context) RoleResponseOutput {
	return o
}

func (o RoleResponseOutput) AutoscaleConfiguration() AutoscaleResponsePtrOutput {
	return o.ApplyT(func(v RoleResponse) *AutoscaleResponse { return v.AutoscaleConfiguration }).(AutoscaleResponsePtrOutput)
}

func (o RoleResponseOutput) DataDisksGroups() DataDisksGroupsResponseArrayOutput {
	return o.ApplyT(func(v RoleResponse) []DataDisksGroupsResponse { return v.DataDisksGroups }).(DataDisksGroupsResponseArrayOutput)
}

func (o RoleResponseOutput) EncryptDataDisks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RoleResponse) *bool { return v.EncryptDataDisks }).(pulumi.BoolPtrOutput)
}

func (o RoleResponseOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v RoleResponse) *HardwareProfileResponse { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o RoleResponseOutput) MinInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RoleResponse) *int { return v.MinInstanceCount }).(pulumi.IntPtrOutput)
}

func (o RoleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RoleResponseOutput) OsProfile() OsProfileResponsePtrOutput {
	return o.ApplyT(func(v RoleResponse) *OsProfileResponse { return v.OsProfile }).(OsProfileResponsePtrOutput)
}

func (o RoleResponseOutput) ScriptActions() ScriptActionResponseArrayOutput {
	return o.ApplyT(func(v RoleResponse) []ScriptActionResponse { return v.ScriptActions }).(ScriptActionResponseArrayOutput)
}

func (o RoleResponseOutput) TargetInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RoleResponse) *int { return v.TargetInstanceCount }).(pulumi.IntPtrOutput)
}

func (o RoleResponseOutput) VMGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleResponse) *string { return v.VMGroupName }).(pulumi.StringPtrOutput)
}

func (o RoleResponseOutput) VirtualNetworkProfile() VirtualNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v RoleResponse) *VirtualNetworkProfileResponse { return v.VirtualNetworkProfile }).(VirtualNetworkProfileResponsePtrOutput)
}

type RoleResponseArrayOutput struct{ *pulumi.OutputState }

func (RoleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoleResponse)(nil)).Elem()
}

func (o RoleResponseArrayOutput) ToRoleResponseArrayOutput() RoleResponseArrayOutput {
	return o
}

func (o RoleResponseArrayOutput) ToRoleResponseArrayOutputWithContext(ctx context.Context) RoleResponseArrayOutput {
	return o
}

func (o RoleResponseArrayOutput) Index(i pulumi.IntInput) RoleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoleResponse {
		return vs[0].([]RoleResponse)[vs[1].(int)]
	}).(RoleResponseOutput)
}

type RuntimeScriptAction struct {
	Name       string   `pulumi:"name"`
	Parameters *string  `pulumi:"parameters"`
	Roles      []string `pulumi:"roles"`
	Uri        string   `pulumi:"uri"`
}





type RuntimeScriptActionInput interface {
	pulumi.Input

	ToRuntimeScriptActionOutput() RuntimeScriptActionOutput
	ToRuntimeScriptActionOutputWithContext(context.Context) RuntimeScriptActionOutput
}

type RuntimeScriptActionArgs struct {
	Name       pulumi.StringInput      `pulumi:"name"`
	Parameters pulumi.StringPtrInput   `pulumi:"parameters"`
	Roles      pulumi.StringArrayInput `pulumi:"roles"`
	Uri        pulumi.StringInput      `pulumi:"uri"`
}

func (RuntimeScriptActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeScriptAction)(nil)).Elem()
}

func (i RuntimeScriptActionArgs) ToRuntimeScriptActionOutput() RuntimeScriptActionOutput {
	return i.ToRuntimeScriptActionOutputWithContext(context.Background())
}

func (i RuntimeScriptActionArgs) ToRuntimeScriptActionOutputWithContext(ctx context.Context) RuntimeScriptActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeScriptActionOutput)
}





type RuntimeScriptActionArrayInput interface {
	pulumi.Input

	ToRuntimeScriptActionArrayOutput() RuntimeScriptActionArrayOutput
	ToRuntimeScriptActionArrayOutputWithContext(context.Context) RuntimeScriptActionArrayOutput
}

type RuntimeScriptActionArray []RuntimeScriptActionInput

func (RuntimeScriptActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuntimeScriptAction)(nil)).Elem()
}

func (i RuntimeScriptActionArray) ToRuntimeScriptActionArrayOutput() RuntimeScriptActionArrayOutput {
	return i.ToRuntimeScriptActionArrayOutputWithContext(context.Background())
}

func (i RuntimeScriptActionArray) ToRuntimeScriptActionArrayOutputWithContext(ctx context.Context) RuntimeScriptActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeScriptActionArrayOutput)
}

type RuntimeScriptActionOutput struct{ *pulumi.OutputState }

func (RuntimeScriptActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeScriptAction)(nil)).Elem()
}

func (o RuntimeScriptActionOutput) ToRuntimeScriptActionOutput() RuntimeScriptActionOutput {
	return o
}

func (o RuntimeScriptActionOutput) ToRuntimeScriptActionOutputWithContext(ctx context.Context) RuntimeScriptActionOutput {
	return o
}

func (o RuntimeScriptActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RuntimeScriptAction) string { return v.Name }).(pulumi.StringOutput)
}

func (o RuntimeScriptActionOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuntimeScriptAction) *string { return v.Parameters }).(pulumi.StringPtrOutput)
}

func (o RuntimeScriptActionOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RuntimeScriptAction) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o RuntimeScriptActionOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v RuntimeScriptAction) string { return v.Uri }).(pulumi.StringOutput)
}

type RuntimeScriptActionArrayOutput struct{ *pulumi.OutputState }

func (RuntimeScriptActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuntimeScriptAction)(nil)).Elem()
}

func (o RuntimeScriptActionArrayOutput) ToRuntimeScriptActionArrayOutput() RuntimeScriptActionArrayOutput {
	return o
}

func (o RuntimeScriptActionArrayOutput) ToRuntimeScriptActionArrayOutputWithContext(ctx context.Context) RuntimeScriptActionArrayOutput {
	return o
}

func (o RuntimeScriptActionArrayOutput) Index(i pulumi.IntInput) RuntimeScriptActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuntimeScriptAction {
		return vs[0].([]RuntimeScriptAction)[vs[1].(int)]
	}).(RuntimeScriptActionOutput)
}

type RuntimeScriptActionResponse struct {
	ApplicationName string   `pulumi:"applicationName"`
	Name            string   `pulumi:"name"`
	Roles           []string `pulumi:"roles"`
	Uri             string   `pulumi:"uri"`
}

type RuntimeScriptActionResponseOutput struct{ *pulumi.OutputState }

func (RuntimeScriptActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeScriptActionResponse)(nil)).Elem()
}

func (o RuntimeScriptActionResponseOutput) ToRuntimeScriptActionResponseOutput() RuntimeScriptActionResponseOutput {
	return o
}

func (o RuntimeScriptActionResponseOutput) ToRuntimeScriptActionResponseOutputWithContext(ctx context.Context) RuntimeScriptActionResponseOutput {
	return o
}

func (o RuntimeScriptActionResponseOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v RuntimeScriptActionResponse) string { return v.ApplicationName }).(pulumi.StringOutput)
}

func (o RuntimeScriptActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RuntimeScriptActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RuntimeScriptActionResponseOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RuntimeScriptActionResponse) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o RuntimeScriptActionResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v RuntimeScriptActionResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type RuntimeScriptActionResponseArrayOutput struct{ *pulumi.OutputState }

func (RuntimeScriptActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuntimeScriptActionResponse)(nil)).Elem()
}

func (o RuntimeScriptActionResponseArrayOutput) ToRuntimeScriptActionResponseArrayOutput() RuntimeScriptActionResponseArrayOutput {
	return o
}

func (o RuntimeScriptActionResponseArrayOutput) ToRuntimeScriptActionResponseArrayOutputWithContext(ctx context.Context) RuntimeScriptActionResponseArrayOutput {
	return o
}

func (o RuntimeScriptActionResponseArrayOutput) Index(i pulumi.IntInput) RuntimeScriptActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuntimeScriptActionResponse {
		return vs[0].([]RuntimeScriptActionResponse)[vs[1].(int)]
	}).(RuntimeScriptActionResponseOutput)
}

type ScriptAction struct {
	Name       string `pulumi:"name"`
	Parameters string `pulumi:"parameters"`
	Uri        string `pulumi:"uri"`
}





type ScriptActionInput interface {
	pulumi.Input

	ToScriptActionOutput() ScriptActionOutput
	ToScriptActionOutputWithContext(context.Context) ScriptActionOutput
}

type ScriptActionArgs struct {
	Name       pulumi.StringInput `pulumi:"name"`
	Parameters pulumi.StringInput `pulumi:"parameters"`
	Uri        pulumi.StringInput `pulumi:"uri"`
}

func (ScriptActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptAction)(nil)).Elem()
}

func (i ScriptActionArgs) ToScriptActionOutput() ScriptActionOutput {
	return i.ToScriptActionOutputWithContext(context.Background())
}

func (i ScriptActionArgs) ToScriptActionOutputWithContext(ctx context.Context) ScriptActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptActionOutput)
}





type ScriptActionArrayInput interface {
	pulumi.Input

	ToScriptActionArrayOutput() ScriptActionArrayOutput
	ToScriptActionArrayOutputWithContext(context.Context) ScriptActionArrayOutput
}

type ScriptActionArray []ScriptActionInput

func (ScriptActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScriptAction)(nil)).Elem()
}

func (i ScriptActionArray) ToScriptActionArrayOutput() ScriptActionArrayOutput {
	return i.ToScriptActionArrayOutputWithContext(context.Background())
}

func (i ScriptActionArray) ToScriptActionArrayOutputWithContext(ctx context.Context) ScriptActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptActionArrayOutput)
}

type ScriptActionOutput struct{ *pulumi.OutputState }

func (ScriptActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptAction)(nil)).Elem()
}

func (o ScriptActionOutput) ToScriptActionOutput() ScriptActionOutput {
	return o
}

func (o ScriptActionOutput) ToScriptActionOutputWithContext(ctx context.Context) ScriptActionOutput {
	return o
}

func (o ScriptActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptAction) string { return v.Name }).(pulumi.StringOutput)
}

func (o ScriptActionOutput) Parameters() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptAction) string { return v.Parameters }).(pulumi.StringOutput)
}

func (o ScriptActionOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptAction) string { return v.Uri }).(pulumi.StringOutput)
}

type ScriptActionArrayOutput struct{ *pulumi.OutputState }

func (ScriptActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScriptAction)(nil)).Elem()
}

func (o ScriptActionArrayOutput) ToScriptActionArrayOutput() ScriptActionArrayOutput {
	return o
}

func (o ScriptActionArrayOutput) ToScriptActionArrayOutputWithContext(ctx context.Context) ScriptActionArrayOutput {
	return o
}

func (o ScriptActionArrayOutput) Index(i pulumi.IntInput) ScriptActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScriptAction {
		return vs[0].([]ScriptAction)[vs[1].(int)]
	}).(ScriptActionOutput)
}

type ScriptActionResponse struct {
	Name       string `pulumi:"name"`
	Parameters string `pulumi:"parameters"`
	Uri        string `pulumi:"uri"`
}

type ScriptActionResponseOutput struct{ *pulumi.OutputState }

func (ScriptActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptActionResponse)(nil)).Elem()
}

func (o ScriptActionResponseOutput) ToScriptActionResponseOutput() ScriptActionResponseOutput {
	return o
}

func (o ScriptActionResponseOutput) ToScriptActionResponseOutputWithContext(ctx context.Context) ScriptActionResponseOutput {
	return o
}

func (o ScriptActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ScriptActionResponseOutput) Parameters() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptActionResponse) string { return v.Parameters }).(pulumi.StringOutput)
}

func (o ScriptActionResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptActionResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type ScriptActionResponseArrayOutput struct{ *pulumi.OutputState }

func (ScriptActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScriptActionResponse)(nil)).Elem()
}

func (o ScriptActionResponseArrayOutput) ToScriptActionResponseArrayOutput() ScriptActionResponseArrayOutput {
	return o
}

func (o ScriptActionResponseArrayOutput) ToScriptActionResponseArrayOutputWithContext(ctx context.Context) ScriptActionResponseArrayOutput {
	return o
}

func (o ScriptActionResponseArrayOutput) Index(i pulumi.IntInput) ScriptActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScriptActionResponse {
		return vs[0].([]ScriptActionResponse)[vs[1].(int)]
	}).(ScriptActionResponseOutput)
}

type SecurityProfile struct {
	AaddsResourceId      *string  `pulumi:"aaddsResourceId"`
	ClusterUsersGroupDNs []string `pulumi:"clusterUsersGroupDNs"`
	DirectoryType        *string  `pulumi:"directoryType"`
	Domain               *string  `pulumi:"domain"`
	DomainUserPassword   *string  `pulumi:"domainUserPassword"`
	DomainUsername       *string  `pulumi:"domainUsername"`
	LdapsUrls            []string `pulumi:"ldapsUrls"`
	MsiResourceId        *string  `pulumi:"msiResourceId"`
	OrganizationalUnitDN *string  `pulumi:"organizationalUnitDN"`
}





type SecurityProfileInput interface {
	pulumi.Input

	ToSecurityProfileOutput() SecurityProfileOutput
	ToSecurityProfileOutputWithContext(context.Context) SecurityProfileOutput
}

type SecurityProfileArgs struct {
	AaddsResourceId      pulumi.StringPtrInput   `pulumi:"aaddsResourceId"`
	ClusterUsersGroupDNs pulumi.StringArrayInput `pulumi:"clusterUsersGroupDNs"`
	DirectoryType        pulumi.StringPtrInput   `pulumi:"directoryType"`
	Domain               pulumi.StringPtrInput   `pulumi:"domain"`
	DomainUserPassword   pulumi.StringPtrInput   `pulumi:"domainUserPassword"`
	DomainUsername       pulumi.StringPtrInput   `pulumi:"domainUsername"`
	LdapsUrls            pulumi.StringArrayInput `pulumi:"ldapsUrls"`
	MsiResourceId        pulumi.StringPtrInput   `pulumi:"msiResourceId"`
	OrganizationalUnitDN pulumi.StringPtrInput   `pulumi:"organizationalUnitDN"`
}

func (SecurityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityProfile)(nil)).Elem()
}

func (i SecurityProfileArgs) ToSecurityProfileOutput() SecurityProfileOutput {
	return i.ToSecurityProfileOutputWithContext(context.Background())
}

func (i SecurityProfileArgs) ToSecurityProfileOutputWithContext(ctx context.Context) SecurityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityProfileOutput)
}

func (i SecurityProfileArgs) ToSecurityProfilePtrOutput() SecurityProfilePtrOutput {
	return i.ToSecurityProfilePtrOutputWithContext(context.Background())
}

func (i SecurityProfileArgs) ToSecurityProfilePtrOutputWithContext(ctx context.Context) SecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityProfileOutput).ToSecurityProfilePtrOutputWithContext(ctx)
}









type SecurityProfilePtrInput interface {
	pulumi.Input

	ToSecurityProfilePtrOutput() SecurityProfilePtrOutput
	ToSecurityProfilePtrOutputWithContext(context.Context) SecurityProfilePtrOutput
}

type securityProfilePtrType SecurityProfileArgs

func SecurityProfilePtr(v *SecurityProfileArgs) SecurityProfilePtrInput {
	return (*securityProfilePtrType)(v)
}

func (*securityProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityProfile)(nil)).Elem()
}

func (i *securityProfilePtrType) ToSecurityProfilePtrOutput() SecurityProfilePtrOutput {
	return i.ToSecurityProfilePtrOutputWithContext(context.Background())
}

func (i *securityProfilePtrType) ToSecurityProfilePtrOutputWithContext(ctx context.Context) SecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityProfilePtrOutput)
}

type SecurityProfileOutput struct{ *pulumi.OutputState }

func (SecurityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityProfile)(nil)).Elem()
}

func (o SecurityProfileOutput) ToSecurityProfileOutput() SecurityProfileOutput {
	return o
}

func (o SecurityProfileOutput) ToSecurityProfileOutputWithContext(ctx context.Context) SecurityProfileOutput {
	return o
}

func (o SecurityProfileOutput) ToSecurityProfilePtrOutput() SecurityProfilePtrOutput {
	return o.ToSecurityProfilePtrOutputWithContext(context.Background())
}

func (o SecurityProfileOutput) ToSecurityProfilePtrOutputWithContext(ctx context.Context) SecurityProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityProfile) *SecurityProfile {
		return &v
	}).(SecurityProfilePtrOutput)
}

func (o SecurityProfileOutput) AaddsResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfile) *string { return v.AaddsResourceId }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileOutput) ClusterUsersGroupDNs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityProfile) []string { return v.ClusterUsersGroupDNs }).(pulumi.StringArrayOutput)
}

func (o SecurityProfileOutput) DirectoryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfile) *string { return v.DirectoryType }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfile) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileOutput) DomainUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfile) *string { return v.DomainUserPassword }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileOutput) DomainUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfile) *string { return v.DomainUsername }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileOutput) LdapsUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityProfile) []string { return v.LdapsUrls }).(pulumi.StringArrayOutput)
}

func (o SecurityProfileOutput) MsiResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfile) *string { return v.MsiResourceId }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileOutput) OrganizationalUnitDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfile) *string { return v.OrganizationalUnitDN }).(pulumi.StringPtrOutput)
}

type SecurityProfilePtrOutput struct{ *pulumi.OutputState }

func (SecurityProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityProfile)(nil)).Elem()
}

func (o SecurityProfilePtrOutput) ToSecurityProfilePtrOutput() SecurityProfilePtrOutput {
	return o
}

func (o SecurityProfilePtrOutput) ToSecurityProfilePtrOutputWithContext(ctx context.Context) SecurityProfilePtrOutput {
	return o
}

func (o SecurityProfilePtrOutput) Elem() SecurityProfileOutput {
	return o.ApplyT(func(v *SecurityProfile) SecurityProfile {
		if v != nil {
			return *v
		}
		var ret SecurityProfile
		return ret
	}).(SecurityProfileOutput)
}

func (o SecurityProfilePtrOutput) AaddsResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.AaddsResourceId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfilePtrOutput) ClusterUsersGroupDNs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityProfile) []string {
		if v == nil {
			return nil
		}
		return v.ClusterUsersGroupDNs
	}).(pulumi.StringArrayOutput)
}

func (o SecurityProfilePtrOutput) DirectoryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.DirectoryType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfilePtrOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.Domain
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfilePtrOutput) DomainUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.DomainUserPassword
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfilePtrOutput) DomainUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.DomainUsername
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfilePtrOutput) LdapsUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityProfile) []string {
		if v == nil {
			return nil
		}
		return v.LdapsUrls
	}).(pulumi.StringArrayOutput)
}

func (o SecurityProfilePtrOutput) MsiResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.MsiResourceId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfilePtrOutput) OrganizationalUnitDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.OrganizationalUnitDN
	}).(pulumi.StringPtrOutput)
}

type SecurityProfileResponse struct {
	AaddsResourceId      *string  `pulumi:"aaddsResourceId"`
	ClusterUsersGroupDNs []string `pulumi:"clusterUsersGroupDNs"`
	DirectoryType        *string  `pulumi:"directoryType"`
	Domain               *string  `pulumi:"domain"`
	DomainUserPassword   *string  `pulumi:"domainUserPassword"`
	DomainUsername       *string  `pulumi:"domainUsername"`
	LdapsUrls            []string `pulumi:"ldapsUrls"`
	MsiResourceId        *string  `pulumi:"msiResourceId"`
	OrganizationalUnitDN *string  `pulumi:"organizationalUnitDN"`
}

type SecurityProfileResponseOutput struct{ *pulumi.OutputState }

func (SecurityProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityProfileResponse)(nil)).Elem()
}

func (o SecurityProfileResponseOutput) ToSecurityProfileResponseOutput() SecurityProfileResponseOutput {
	return o
}

func (o SecurityProfileResponseOutput) ToSecurityProfileResponseOutputWithContext(ctx context.Context) SecurityProfileResponseOutput {
	return o
}

func (o SecurityProfileResponseOutput) AaddsResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfileResponse) *string { return v.AaddsResourceId }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponseOutput) ClusterUsersGroupDNs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityProfileResponse) []string { return v.ClusterUsersGroupDNs }).(pulumi.StringArrayOutput)
}

func (o SecurityProfileResponseOutput) DirectoryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfileResponse) *string { return v.DirectoryType }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponseOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfileResponse) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponseOutput) DomainUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfileResponse) *string { return v.DomainUserPassword }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponseOutput) DomainUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfileResponse) *string { return v.DomainUsername }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponseOutput) LdapsUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityProfileResponse) []string { return v.LdapsUrls }).(pulumi.StringArrayOutput)
}

func (o SecurityProfileResponseOutput) MsiResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfileResponse) *string { return v.MsiResourceId }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponseOutput) OrganizationalUnitDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfileResponse) *string { return v.OrganizationalUnitDN }).(pulumi.StringPtrOutput)
}

type SecurityProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityProfileResponse)(nil)).Elem()
}

func (o SecurityProfileResponsePtrOutput) ToSecurityProfileResponsePtrOutput() SecurityProfileResponsePtrOutput {
	return o
}

func (o SecurityProfileResponsePtrOutput) ToSecurityProfileResponsePtrOutputWithContext(ctx context.Context) SecurityProfileResponsePtrOutput {
	return o
}

func (o SecurityProfileResponsePtrOutput) Elem() SecurityProfileResponseOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) SecurityProfileResponse {
		if v != nil {
			return *v
		}
		var ret SecurityProfileResponse
		return ret
	}).(SecurityProfileResponseOutput)
}

func (o SecurityProfileResponsePtrOutput) AaddsResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AaddsResourceId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponsePtrOutput) ClusterUsersGroupDNs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.ClusterUsersGroupDNs
	}).(pulumi.StringArrayOutput)
}

func (o SecurityProfileResponsePtrOutput) DirectoryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DirectoryType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponsePtrOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Domain
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponsePtrOutput) DomainUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DomainUserPassword
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponsePtrOutput) DomainUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DomainUsername
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponsePtrOutput) LdapsUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.LdapsUrls
	}).(pulumi.StringArrayOutput)
}

func (o SecurityProfileResponsePtrOutput) MsiResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.MsiResourceId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponsePtrOutput) OrganizationalUnitDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.OrganizationalUnitDN
	}).(pulumi.StringPtrOutput)
}

type SshProfile struct {
	PublicKeys []SshPublicKey `pulumi:"publicKeys"`
}





type SshProfileInput interface {
	pulumi.Input

	ToSshProfileOutput() SshProfileOutput
	ToSshProfileOutputWithContext(context.Context) SshProfileOutput
}

type SshProfileArgs struct {
	PublicKeys SshPublicKeyArrayInput `pulumi:"publicKeys"`
}

func (SshProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshProfile)(nil)).Elem()
}

func (i SshProfileArgs) ToSshProfileOutput() SshProfileOutput {
	return i.ToSshProfileOutputWithContext(context.Background())
}

func (i SshProfileArgs) ToSshProfileOutputWithContext(ctx context.Context) SshProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshProfileOutput)
}

func (i SshProfileArgs) ToSshProfilePtrOutput() SshProfilePtrOutput {
	return i.ToSshProfilePtrOutputWithContext(context.Background())
}

func (i SshProfileArgs) ToSshProfilePtrOutputWithContext(ctx context.Context) SshProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshProfileOutput).ToSshProfilePtrOutputWithContext(ctx)
}









type SshProfilePtrInput interface {
	pulumi.Input

	ToSshProfilePtrOutput() SshProfilePtrOutput
	ToSshProfilePtrOutputWithContext(context.Context) SshProfilePtrOutput
}

type sshProfilePtrType SshProfileArgs

func SshProfilePtr(v *SshProfileArgs) SshProfilePtrInput {
	return (*sshProfilePtrType)(v)
}

func (*sshProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SshProfile)(nil)).Elem()
}

func (i *sshProfilePtrType) ToSshProfilePtrOutput() SshProfilePtrOutput {
	return i.ToSshProfilePtrOutputWithContext(context.Background())
}

func (i *sshProfilePtrType) ToSshProfilePtrOutputWithContext(ctx context.Context) SshProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshProfilePtrOutput)
}

type SshProfileOutput struct{ *pulumi.OutputState }

func (SshProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshProfile)(nil)).Elem()
}

func (o SshProfileOutput) ToSshProfileOutput() SshProfileOutput {
	return o
}

func (o SshProfileOutput) ToSshProfileOutputWithContext(ctx context.Context) SshProfileOutput {
	return o
}

func (o SshProfileOutput) ToSshProfilePtrOutput() SshProfilePtrOutput {
	return o.ToSshProfilePtrOutputWithContext(context.Background())
}

func (o SshProfileOutput) ToSshProfilePtrOutputWithContext(ctx context.Context) SshProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SshProfile) *SshProfile {
		return &v
	}).(SshProfilePtrOutput)
}

func (o SshProfileOutput) PublicKeys() SshPublicKeyArrayOutput {
	return o.ApplyT(func(v SshProfile) []SshPublicKey { return v.PublicKeys }).(SshPublicKeyArrayOutput)
}

type SshProfilePtrOutput struct{ *pulumi.OutputState }

func (SshProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshProfile)(nil)).Elem()
}

func (o SshProfilePtrOutput) ToSshProfilePtrOutput() SshProfilePtrOutput {
	return o
}

func (o SshProfilePtrOutput) ToSshProfilePtrOutputWithContext(ctx context.Context) SshProfilePtrOutput {
	return o
}

func (o SshProfilePtrOutput) Elem() SshProfileOutput {
	return o.ApplyT(func(v *SshProfile) SshProfile {
		if v != nil {
			return *v
		}
		var ret SshProfile
		return ret
	}).(SshProfileOutput)
}

func (o SshProfilePtrOutput) PublicKeys() SshPublicKeyArrayOutput {
	return o.ApplyT(func(v *SshProfile) []SshPublicKey {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(SshPublicKeyArrayOutput)
}

type SshProfileResponse struct {
	PublicKeys []SshPublicKeyResponse `pulumi:"publicKeys"`
}

type SshProfileResponseOutput struct{ *pulumi.OutputState }

func (SshProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshProfileResponse)(nil)).Elem()
}

func (o SshProfileResponseOutput) ToSshProfileResponseOutput() SshProfileResponseOutput {
	return o
}

func (o SshProfileResponseOutput) ToSshProfileResponseOutputWithContext(ctx context.Context) SshProfileResponseOutput {
	return o
}

func (o SshProfileResponseOutput) PublicKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v SshProfileResponse) []SshPublicKeyResponse { return v.PublicKeys }).(SshPublicKeyResponseArrayOutput)
}

type SshProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (SshProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshProfileResponse)(nil)).Elem()
}

func (o SshProfileResponsePtrOutput) ToSshProfileResponsePtrOutput() SshProfileResponsePtrOutput {
	return o
}

func (o SshProfileResponsePtrOutput) ToSshProfileResponsePtrOutputWithContext(ctx context.Context) SshProfileResponsePtrOutput {
	return o
}

func (o SshProfileResponsePtrOutput) Elem() SshProfileResponseOutput {
	return o.ApplyT(func(v *SshProfileResponse) SshProfileResponse {
		if v != nil {
			return *v
		}
		var ret SshProfileResponse
		return ret
	}).(SshProfileResponseOutput)
}

func (o SshProfileResponsePtrOutput) PublicKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v *SshProfileResponse) []SshPublicKeyResponse {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(SshPublicKeyResponseArrayOutput)
}

type SshPublicKey struct {
	CertificateData *string `pulumi:"certificateData"`
}





type SshPublicKeyInput interface {
	pulumi.Input

	ToSshPublicKeyOutput() SshPublicKeyOutput
	ToSshPublicKeyOutputWithContext(context.Context) SshPublicKeyOutput
}

type SshPublicKeyArgs struct {
	CertificateData pulumi.StringPtrInput `pulumi:"certificateData"`
}

func (SshPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKey)(nil)).Elem()
}

func (i SshPublicKeyArgs) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return i.ToSshPublicKeyOutputWithContext(context.Background())
}

func (i SshPublicKeyArgs) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyOutput)
}





type SshPublicKeyArrayInput interface {
	pulumi.Input

	ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput
	ToSshPublicKeyArrayOutputWithContext(context.Context) SshPublicKeyArrayOutput
}

type SshPublicKeyArray []SshPublicKeyInput

func (SshPublicKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKey)(nil)).Elem()
}

func (i SshPublicKeyArray) ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput {
	return i.ToSshPublicKeyArrayOutputWithContext(context.Background())
}

func (i SshPublicKeyArray) ToSshPublicKeyArrayOutputWithContext(ctx context.Context) SshPublicKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyArrayOutput)
}

type SshPublicKeyOutput struct{ *pulumi.OutputState }

func (SshPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) CertificateData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKey) *string { return v.CertificateData }).(pulumi.StringPtrOutput)
}

type SshPublicKeyArrayOutput struct{ *pulumi.OutputState }

func (SshPublicKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyArrayOutput) ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput {
	return o
}

func (o SshPublicKeyArrayOutput) ToSshPublicKeyArrayOutputWithContext(ctx context.Context) SshPublicKeyArrayOutput {
	return o
}

func (o SshPublicKeyArrayOutput) Index(i pulumi.IntInput) SshPublicKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SshPublicKey {
		return vs[0].([]SshPublicKey)[vs[1].(int)]
	}).(SshPublicKeyOutput)
}

type SshPublicKeyResponse struct {
	CertificateData *string `pulumi:"certificateData"`
}

type SshPublicKeyResponseOutput struct{ *pulumi.OutputState }

func (SshPublicKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKeyResponse)(nil)).Elem()
}

func (o SshPublicKeyResponseOutput) ToSshPublicKeyResponseOutput() SshPublicKeyResponseOutput {
	return o
}

func (o SshPublicKeyResponseOutput) ToSshPublicKeyResponseOutputWithContext(ctx context.Context) SshPublicKeyResponseOutput {
	return o
}

func (o SshPublicKeyResponseOutput) CertificateData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKeyResponse) *string { return v.CertificateData }).(pulumi.StringPtrOutput)
}

type SshPublicKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (SshPublicKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKeyResponse)(nil)).Elem()
}

func (o SshPublicKeyResponseArrayOutput) ToSshPublicKeyResponseArrayOutput() SshPublicKeyResponseArrayOutput {
	return o
}

func (o SshPublicKeyResponseArrayOutput) ToSshPublicKeyResponseArrayOutputWithContext(ctx context.Context) SshPublicKeyResponseArrayOutput {
	return o
}

func (o SshPublicKeyResponseArrayOutput) Index(i pulumi.IntInput) SshPublicKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SshPublicKeyResponse {
		return vs[0].([]SshPublicKeyResponse)[vs[1].(int)]
	}).(SshPublicKeyResponseOutput)
}

type StorageAccount struct {
	Container     *string `pulumi:"container"`
	FileSystem    *string `pulumi:"fileSystem"`
	Fileshare     *string `pulumi:"fileshare"`
	IsDefault     *bool   `pulumi:"isDefault"`
	Key           *string `pulumi:"key"`
	MsiResourceId *string `pulumi:"msiResourceId"`
	Name          *string `pulumi:"name"`
	ResourceId    *string `pulumi:"resourceId"`
	Saskey        *string `pulumi:"saskey"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	Container     pulumi.StringPtrInput `pulumi:"container"`
	FileSystem    pulumi.StringPtrInput `pulumi:"fileSystem"`
	Fileshare     pulumi.StringPtrInput `pulumi:"fileshare"`
	IsDefault     pulumi.BoolPtrInput   `pulumi:"isDefault"`
	Key           pulumi.StringPtrInput `pulumi:"key"`
	MsiResourceId pulumi.StringPtrInput `pulumi:"msiResourceId"`
	Name          pulumi.StringPtrInput `pulumi:"name"`
	ResourceId    pulumi.StringPtrInput `pulumi:"resourceId"`
	Saskey        pulumi.StringPtrInput `pulumi:"saskey"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}





type StorageAccountArrayInput interface {
	pulumi.Input

	ToStorageAccountArrayOutput() StorageAccountArrayOutput
	ToStorageAccountArrayOutputWithContext(context.Context) StorageAccountArrayOutput
}

type StorageAccountArray []StorageAccountInput

func (StorageAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (i StorageAccountArray) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return i.ToStorageAccountArrayOutputWithContext(context.Background())
}

func (i StorageAccountArray) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountArrayOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) FileSystem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.FileSystem }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) Fileshare() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.Fileshare }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StorageAccount) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) MsiResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.MsiResourceId }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) Saskey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.Saskey }).(pulumi.StringPtrOutput)
}

type StorageAccountArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) Index(i pulumi.IntInput) StorageAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccount {
		return vs[0].([]StorageAccount)[vs[1].(int)]
	}).(StorageAccountOutput)
}

type StorageAccountResponse struct {
	Container     *string `pulumi:"container"`
	FileSystem    *string `pulumi:"fileSystem"`
	Fileshare     *string `pulumi:"fileshare"`
	IsDefault     *bool   `pulumi:"isDefault"`
	Key           *string `pulumi:"key"`
	MsiResourceId *string `pulumi:"msiResourceId"`
	Name          *string `pulumi:"name"`
	ResourceId    *string `pulumi:"resourceId"`
	Saskey        *string `pulumi:"saskey"`
}

type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) FileSystem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.FileSystem }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) Fileshare() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.Fileshare }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) MsiResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.MsiResourceId }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) Saskey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.Saskey }).(pulumi.StringPtrOutput)
}

type StorageAccountResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutputWithContext(ctx context.Context) StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) Index(i pulumi.IntInput) StorageAccountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccountResponse {
		return vs[0].([]StorageAccountResponse)[vs[1].(int)]
	}).(StorageAccountResponseOutput)
}

type StorageProfile struct {
	Storageaccounts []StorageAccount `pulumi:"storageaccounts"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	Storageaccounts StorageAccountArrayInput `pulumi:"storageaccounts"`
}

func (StorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (i StorageProfileArgs) ToStorageProfileOutput() StorageProfileOutput {
	return i.ToStorageProfileOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput)
}

func (i StorageProfileArgs) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput).ToStorageProfilePtrOutputWithContext(ctx)
}









type StorageProfilePtrInput interface {
	pulumi.Input

	ToStorageProfilePtrOutput() StorageProfilePtrOutput
	ToStorageProfilePtrOutputWithContext(context.Context) StorageProfilePtrOutput
}

type storageProfilePtrType StorageProfileArgs

func StorageProfilePtr(v *StorageProfileArgs) StorageProfilePtrInput {
	return (*storageProfilePtrType)(v)
}

func (*storageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfilePtrOutput)
}

type StorageProfileOutput struct{ *pulumi.OutputState }

func (StorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (o StorageProfileOutput) ToStorageProfileOutput() StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (o StorageProfileOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfile) *StorageProfile {
		return &v
	}).(StorageProfilePtrOutput)
}

func (o StorageProfileOutput) Storageaccounts() StorageAccountArrayOutput {
	return o.ApplyT(func(v StorageProfile) []StorageAccount { return v.Storageaccounts }).(StorageAccountArrayOutput)
}

type StorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) Elem() StorageProfileOutput {
	return o.ApplyT(func(v *StorageProfile) StorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageProfile
		return ret
	}).(StorageProfileOutput)
}

func (o StorageProfilePtrOutput) Storageaccounts() StorageAccountArrayOutput {
	return o.ApplyT(func(v *StorageProfile) []StorageAccount {
		if v == nil {
			return nil
		}
		return v.Storageaccounts
	}).(StorageAccountArrayOutput)
}

type StorageProfileResponse struct {
	Storageaccounts []StorageAccountResponse `pulumi:"storageaccounts"`
}

type StorageProfileResponseOutput struct{ *pulumi.OutputState }

func (StorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) Storageaccounts() StorageAccountResponseArrayOutput {
	return o.ApplyT(func(v StorageProfileResponse) []StorageAccountResponse { return v.Storageaccounts }).(StorageAccountResponseArrayOutput)
}

type StorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) Elem() StorageProfileResponseOutput {
	return o.ApplyT(func(v *StorageProfileResponse) StorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret StorageProfileResponse
		return ret
	}).(StorageProfileResponseOutput)
}

func (o StorageProfileResponsePtrOutput) Storageaccounts() StorageAccountResponseArrayOutput {
	return o.ApplyT(func(v *StorageProfileResponse) []StorageAccountResponse {
		if v == nil {
			return nil
		}
		return v.Storageaccounts
	}).(StorageAccountResponseArrayOutput)
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

type UserAssignedIdentity struct {
	TenantId *string `pulumi:"tenantId"`
}





type UserAssignedIdentityInput interface {
	pulumi.Input

	ToUserAssignedIdentityOutput() UserAssignedIdentityOutput
	ToUserAssignedIdentityOutputWithContext(context.Context) UserAssignedIdentityOutput
}

type UserAssignedIdentityArgs struct {
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (UserAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentity)(nil)).Elem()
}

func (i UserAssignedIdentityArgs) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return i.ToUserAssignedIdentityOutputWithContext(context.Background())
}

func (i UserAssignedIdentityArgs) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityOutput)
}





type UserAssignedIdentityMapInput interface {
	pulumi.Input

	ToUserAssignedIdentityMapOutput() UserAssignedIdentityMapOutput
	ToUserAssignedIdentityMapOutputWithContext(context.Context) UserAssignedIdentityMapOutput
}

type UserAssignedIdentityMap map[string]UserAssignedIdentityInput

func (UserAssignedIdentityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentity)(nil)).Elem()
}

func (i UserAssignedIdentityMap) ToUserAssignedIdentityMapOutput() UserAssignedIdentityMapOutput {
	return i.ToUserAssignedIdentityMapOutputWithContext(context.Background())
}

func (i UserAssignedIdentityMap) ToUserAssignedIdentityMapOutputWithContext(ctx context.Context) UserAssignedIdentityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityMapOutput)
}

type UserAssignedIdentityOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentity)(nil)).Elem()
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentity)(nil)).Elem()
}

func (o UserAssignedIdentityMapOutput) ToUserAssignedIdentityMapOutput() UserAssignedIdentityMapOutput {
	return o
}

func (o UserAssignedIdentityMapOutput) ToUserAssignedIdentityMapOutputWithContext(ctx context.Context) UserAssignedIdentityMapOutput {
	return o
}

func (o UserAssignedIdentityMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentity {
		return vs[0].(map[string]UserAssignedIdentity)[vs[1].(string)]
	}).(UserAssignedIdentityOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string  `pulumi:"clientId"`
	PrincipalId string  `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
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

func (o UserAssignedIdentityResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
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

type VirtualNetworkProfile struct {
	Id     *string `pulumi:"id"`
	Subnet *string `pulumi:"subnet"`
}





type VirtualNetworkProfileInput interface {
	pulumi.Input

	ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput
	ToVirtualNetworkProfileOutputWithContext(context.Context) VirtualNetworkProfileOutput
}

type VirtualNetworkProfileArgs struct {
	Id     pulumi.StringPtrInput `pulumi:"id"`
	Subnet pulumi.StringPtrInput `pulumi:"subnet"`
}

func (VirtualNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfile)(nil)).Elem()
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput {
	return i.ToVirtualNetworkProfileOutputWithContext(context.Background())
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfileOutputWithContext(ctx context.Context) VirtualNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileOutput)
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return i.ToVirtualNetworkProfilePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileOutput).ToVirtualNetworkProfilePtrOutputWithContext(ctx)
}









type VirtualNetworkProfilePtrInput interface {
	pulumi.Input

	ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput
	ToVirtualNetworkProfilePtrOutputWithContext(context.Context) VirtualNetworkProfilePtrOutput
}

type virtualNetworkProfilePtrType VirtualNetworkProfileArgs

func VirtualNetworkProfilePtr(v *VirtualNetworkProfileArgs) VirtualNetworkProfilePtrInput {
	return (*virtualNetworkProfilePtrType)(v)
}

func (*virtualNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfile)(nil)).Elem()
}

func (i *virtualNetworkProfilePtrType) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return i.ToVirtualNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkProfilePtrType) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfilePtrOutput)
}

type VirtualNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfile)(nil)).Elem()
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput {
	return o
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfileOutputWithContext(ctx context.Context) VirtualNetworkProfileOutput {
	return o
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return o.ToVirtualNetworkProfilePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkProfile) *VirtualNetworkProfile {
		return &v
	}).(VirtualNetworkProfilePtrOutput)
}

func (o VirtualNetworkProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfile)(nil)).Elem()
}

func (o VirtualNetworkProfilePtrOutput) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return o
}

func (o VirtualNetworkProfilePtrOutput) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return o
}

func (o VirtualNetworkProfilePtrOutput) Elem() VirtualNetworkProfileOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) VirtualNetworkProfile {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkProfile
		return ret
	}).(VirtualNetworkProfileOutput)
}

func (o VirtualNetworkProfilePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfilePtrOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfileResponse struct {
	Id     *string `pulumi:"id"`
	Subnet *string `pulumi:"subnet"`
}

type VirtualNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponseOutput() VirtualNetworkProfileResponseOutput {
	return o
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponseOutputWithContext(ctx context.Context) VirtualNetworkProfileResponseOutput {
	return o
}

func (o VirtualNetworkProfileResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponseOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualNetworkProfileResponsePtrOutput) ToVirtualNetworkProfileResponsePtrOutput() VirtualNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualNetworkProfileResponsePtrOutput) ToVirtualNetworkProfileResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualNetworkProfileResponsePtrOutput) Elem() VirtualNetworkProfileResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) VirtualNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkProfileResponse
		return ret
	}).(VirtualNetworkProfileResponseOutput)
}

func (o VirtualNetworkProfileResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponsePtrOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationGetEndpointOutput{})
	pulumi.RegisterOutputType(ApplicationGetEndpointArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGetEndpointResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGetEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGetHttpsEndpointOutput{})
	pulumi.RegisterOutputType(ApplicationGetHttpsEndpointArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGetHttpsEndpointResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGetHttpsEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationPropertiesOutput{})
	pulumi.RegisterOutputType(ApplicationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApplicationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AutoscaleOutput{})
	pulumi.RegisterOutputType(AutoscalePtrOutput{})
	pulumi.RegisterOutputType(AutoscaleCapacityOutput{})
	pulumi.RegisterOutputType(AutoscaleCapacityPtrOutput{})
	pulumi.RegisterOutputType(AutoscaleCapacityResponseOutput{})
	pulumi.RegisterOutputType(AutoscaleCapacityResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoscaleRecurrenceOutput{})
	pulumi.RegisterOutputType(AutoscaleRecurrencePtrOutput{})
	pulumi.RegisterOutputType(AutoscaleRecurrenceResponseOutput{})
	pulumi.RegisterOutputType(AutoscaleRecurrenceResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoscaleResponseOutput{})
	pulumi.RegisterOutputType(AutoscaleResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoscaleScheduleOutput{})
	pulumi.RegisterOutputType(AutoscaleScheduleArrayOutput{})
	pulumi.RegisterOutputType(AutoscaleScheduleResponseOutput{})
	pulumi.RegisterOutputType(AutoscaleScheduleResponseArrayOutput{})
	pulumi.RegisterOutputType(AutoscaleTimeAndCapacityOutput{})
	pulumi.RegisterOutputType(AutoscaleTimeAndCapacityPtrOutput{})
	pulumi.RegisterOutputType(AutoscaleTimeAndCapacityResponseOutput{})
	pulumi.RegisterOutputType(AutoscaleTimeAndCapacityResponsePtrOutput{})
	pulumi.RegisterOutputType(ClientGroupInfoOutput{})
	pulumi.RegisterOutputType(ClientGroupInfoPtrOutput{})
	pulumi.RegisterOutputType(ClientGroupInfoResponseOutput{})
	pulumi.RegisterOutputType(ClientGroupInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterCreatePropertiesOutput{})
	pulumi.RegisterOutputType(ClusterCreatePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ClusterDefinitionOutput{})
	pulumi.RegisterOutputType(ClusterDefinitionPtrOutput{})
	pulumi.RegisterOutputType(ClusterDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ClusterGetPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ClusterIdentityOutput{})
	pulumi.RegisterOutputType(ClusterIdentityPtrOutput{})
	pulumi.RegisterOutputType(ClusterIdentityResponseOutput{})
	pulumi.RegisterOutputType(ClusterIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ComputeIsolationPropertiesOutput{})
	pulumi.RegisterOutputType(ComputeIsolationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ComputeIsolationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ComputeIsolationPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ComputeProfileOutput{})
	pulumi.RegisterOutputType(ComputeProfilePtrOutput{})
	pulumi.RegisterOutputType(ComputeProfileResponseOutput{})
	pulumi.RegisterOutputType(ComputeProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectivityEndpointResponseOutput{})
	pulumi.RegisterOutputType(ConnectivityEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(DataDisksGroupsOutput{})
	pulumi.RegisterOutputType(DataDisksGroupsArrayOutput{})
	pulumi.RegisterOutputType(DataDisksGroupsResponseOutput{})
	pulumi.RegisterOutputType(DataDisksGroupsResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskEncryptionPropertiesOutput{})
	pulumi.RegisterOutputType(DiskEncryptionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DiskEncryptionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionInTransitPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionInTransitPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionInTransitPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionInTransitPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorsOutput{})
	pulumi.RegisterOutputType(ErrorsArrayOutput{})
	pulumi.RegisterOutputType(ErrorsResponseOutput{})
	pulumi.RegisterOutputType(ErrorsResponseArrayOutput{})
	pulumi.RegisterOutputType(ExcludedServicesConfigResponseOutput{})
	pulumi.RegisterOutputType(ExcludedServicesConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(HardwareProfileOutput{})
	pulumi.RegisterOutputType(HardwareProfilePtrOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponseOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(IPConfigurationOutput{})
	pulumi.RegisterOutputType(IPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(IPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(IPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(KafkaRestPropertiesOutput{})
	pulumi.RegisterOutputType(KafkaRestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KafkaRestPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KafkaRestPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxOperatingSystemProfileOutput{})
	pulumi.RegisterOutputType(LinuxOperatingSystemProfilePtrOutput{})
	pulumi.RegisterOutputType(LinuxOperatingSystemProfileResponseOutput{})
	pulumi.RegisterOutputType(LinuxOperatingSystemProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkPropertiesOutput{})
	pulumi.RegisterOutputType(NetworkPropertiesPtrOutput{})
	pulumi.RegisterOutputType(NetworkPropertiesResponseOutput{})
	pulumi.RegisterOutputType(NetworkPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(OsProfileOutput{})
	pulumi.RegisterOutputType(OsProfilePtrOutput{})
	pulumi.RegisterOutputType(OsProfileResponseOutput{})
	pulumi.RegisterOutputType(OsProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkConfigurationOutput{})
	pulumi.RegisterOutputType(PrivateLinkConfigurationArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(QuotaInfoResponseOutput{})
	pulumi.RegisterOutputType(QuotaInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdOutput{})
	pulumi.RegisterOutputType(ResourceIdPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RoleArrayOutput{})
	pulumi.RegisterOutputType(RoleResponseOutput{})
	pulumi.RegisterOutputType(RoleResponseArrayOutput{})
	pulumi.RegisterOutputType(RuntimeScriptActionOutput{})
	pulumi.RegisterOutputType(RuntimeScriptActionArrayOutput{})
	pulumi.RegisterOutputType(RuntimeScriptActionResponseOutput{})
	pulumi.RegisterOutputType(RuntimeScriptActionResponseArrayOutput{})
	pulumi.RegisterOutputType(ScriptActionOutput{})
	pulumi.RegisterOutputType(ScriptActionArrayOutput{})
	pulumi.RegisterOutputType(ScriptActionResponseOutput{})
	pulumi.RegisterOutputType(ScriptActionResponseArrayOutput{})
	pulumi.RegisterOutputType(SecurityProfileOutput{})
	pulumi.RegisterOutputType(SecurityProfilePtrOutput{})
	pulumi.RegisterOutputType(SecurityProfileResponseOutput{})
	pulumi.RegisterOutputType(SecurityProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SshProfileOutput{})
	pulumi.RegisterOutputType(SshProfilePtrOutput{})
	pulumi.RegisterOutputType(SshProfileResponseOutput{})
	pulumi.RegisterOutputType(SshProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SshPublicKeyOutput{})
	pulumi.RegisterOutputType(SshPublicKeyArrayOutput{})
	pulumi.RegisterOutputType(SshPublicKeyResponseOutput{})
	pulumi.RegisterOutputType(SshPublicKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityMapOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileResponsePtrOutput{})
}

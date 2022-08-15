


package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureMonitorWorkspaceIntegration struct {
	AzureMonitorWorkspaceResourceId *string `pulumi:"azureMonitorWorkspaceResourceId"`
}





type AzureMonitorWorkspaceIntegrationInput interface {
	pulumi.Input

	ToAzureMonitorWorkspaceIntegrationOutput() AzureMonitorWorkspaceIntegrationOutput
	ToAzureMonitorWorkspaceIntegrationOutputWithContext(context.Context) AzureMonitorWorkspaceIntegrationOutput
}

type AzureMonitorWorkspaceIntegrationArgs struct {
	AzureMonitorWorkspaceResourceId pulumi.StringPtrInput `pulumi:"azureMonitorWorkspaceResourceId"`
}

func (AzureMonitorWorkspaceIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorWorkspaceIntegration)(nil)).Elem()
}

func (i AzureMonitorWorkspaceIntegrationArgs) ToAzureMonitorWorkspaceIntegrationOutput() AzureMonitorWorkspaceIntegrationOutput {
	return i.ToAzureMonitorWorkspaceIntegrationOutputWithContext(context.Background())
}

func (i AzureMonitorWorkspaceIntegrationArgs) ToAzureMonitorWorkspaceIntegrationOutputWithContext(ctx context.Context) AzureMonitorWorkspaceIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorWorkspaceIntegrationOutput)
}





type AzureMonitorWorkspaceIntegrationArrayInput interface {
	pulumi.Input

	ToAzureMonitorWorkspaceIntegrationArrayOutput() AzureMonitorWorkspaceIntegrationArrayOutput
	ToAzureMonitorWorkspaceIntegrationArrayOutputWithContext(context.Context) AzureMonitorWorkspaceIntegrationArrayOutput
}

type AzureMonitorWorkspaceIntegrationArray []AzureMonitorWorkspaceIntegrationInput

func (AzureMonitorWorkspaceIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMonitorWorkspaceIntegration)(nil)).Elem()
}

func (i AzureMonitorWorkspaceIntegrationArray) ToAzureMonitorWorkspaceIntegrationArrayOutput() AzureMonitorWorkspaceIntegrationArrayOutput {
	return i.ToAzureMonitorWorkspaceIntegrationArrayOutputWithContext(context.Background())
}

func (i AzureMonitorWorkspaceIntegrationArray) ToAzureMonitorWorkspaceIntegrationArrayOutputWithContext(ctx context.Context) AzureMonitorWorkspaceIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorWorkspaceIntegrationArrayOutput)
}

type AzureMonitorWorkspaceIntegrationOutput struct{ *pulumi.OutputState }

func (AzureMonitorWorkspaceIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorWorkspaceIntegration)(nil)).Elem()
}

func (o AzureMonitorWorkspaceIntegrationOutput) ToAzureMonitorWorkspaceIntegrationOutput() AzureMonitorWorkspaceIntegrationOutput {
	return o
}

func (o AzureMonitorWorkspaceIntegrationOutput) ToAzureMonitorWorkspaceIntegrationOutputWithContext(ctx context.Context) AzureMonitorWorkspaceIntegrationOutput {
	return o
}

func (o AzureMonitorWorkspaceIntegrationOutput) AzureMonitorWorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMonitorWorkspaceIntegration) *string { return v.AzureMonitorWorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type AzureMonitorWorkspaceIntegrationArrayOutput struct{ *pulumi.OutputState }

func (AzureMonitorWorkspaceIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMonitorWorkspaceIntegration)(nil)).Elem()
}

func (o AzureMonitorWorkspaceIntegrationArrayOutput) ToAzureMonitorWorkspaceIntegrationArrayOutput() AzureMonitorWorkspaceIntegrationArrayOutput {
	return o
}

func (o AzureMonitorWorkspaceIntegrationArrayOutput) ToAzureMonitorWorkspaceIntegrationArrayOutputWithContext(ctx context.Context) AzureMonitorWorkspaceIntegrationArrayOutput {
	return o
}

func (o AzureMonitorWorkspaceIntegrationArrayOutput) Index(i pulumi.IntInput) AzureMonitorWorkspaceIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureMonitorWorkspaceIntegration {
		return vs[0].([]AzureMonitorWorkspaceIntegration)[vs[1].(int)]
	}).(AzureMonitorWorkspaceIntegrationOutput)
}

type AzureMonitorWorkspaceIntegrationResponse struct {
	AzureMonitorWorkspaceResourceId *string `pulumi:"azureMonitorWorkspaceResourceId"`
}

type AzureMonitorWorkspaceIntegrationResponseOutput struct{ *pulumi.OutputState }

func (AzureMonitorWorkspaceIntegrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorWorkspaceIntegrationResponse)(nil)).Elem()
}

func (o AzureMonitorWorkspaceIntegrationResponseOutput) ToAzureMonitorWorkspaceIntegrationResponseOutput() AzureMonitorWorkspaceIntegrationResponseOutput {
	return o
}

func (o AzureMonitorWorkspaceIntegrationResponseOutput) ToAzureMonitorWorkspaceIntegrationResponseOutputWithContext(ctx context.Context) AzureMonitorWorkspaceIntegrationResponseOutput {
	return o
}

func (o AzureMonitorWorkspaceIntegrationResponseOutput) AzureMonitorWorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMonitorWorkspaceIntegrationResponse) *string { return v.AzureMonitorWorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type AzureMonitorWorkspaceIntegrationResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureMonitorWorkspaceIntegrationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMonitorWorkspaceIntegrationResponse)(nil)).Elem()
}

func (o AzureMonitorWorkspaceIntegrationResponseArrayOutput) ToAzureMonitorWorkspaceIntegrationResponseArrayOutput() AzureMonitorWorkspaceIntegrationResponseArrayOutput {
	return o
}

func (o AzureMonitorWorkspaceIntegrationResponseArrayOutput) ToAzureMonitorWorkspaceIntegrationResponseArrayOutputWithContext(ctx context.Context) AzureMonitorWorkspaceIntegrationResponseArrayOutput {
	return o
}

func (o AzureMonitorWorkspaceIntegrationResponseArrayOutput) Index(i pulumi.IntInput) AzureMonitorWorkspaceIntegrationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureMonitorWorkspaceIntegrationResponse {
		return vs[0].([]AzureMonitorWorkspaceIntegrationResponse)[vs[1].(int)]
	}).(AzureMonitorWorkspaceIntegrationResponseOutput)
}

type GrafanaIntegrations struct {
	AzureMonitorWorkspaceIntegrations []AzureMonitorWorkspaceIntegration `pulumi:"azureMonitorWorkspaceIntegrations"`
}





type GrafanaIntegrationsInput interface {
	pulumi.Input

	ToGrafanaIntegrationsOutput() GrafanaIntegrationsOutput
	ToGrafanaIntegrationsOutputWithContext(context.Context) GrafanaIntegrationsOutput
}

type GrafanaIntegrationsArgs struct {
	AzureMonitorWorkspaceIntegrations AzureMonitorWorkspaceIntegrationArrayInput `pulumi:"azureMonitorWorkspaceIntegrations"`
}

func (GrafanaIntegrationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GrafanaIntegrations)(nil)).Elem()
}

func (i GrafanaIntegrationsArgs) ToGrafanaIntegrationsOutput() GrafanaIntegrationsOutput {
	return i.ToGrafanaIntegrationsOutputWithContext(context.Background())
}

func (i GrafanaIntegrationsArgs) ToGrafanaIntegrationsOutputWithContext(ctx context.Context) GrafanaIntegrationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrafanaIntegrationsOutput)
}

func (i GrafanaIntegrationsArgs) ToGrafanaIntegrationsPtrOutput() GrafanaIntegrationsPtrOutput {
	return i.ToGrafanaIntegrationsPtrOutputWithContext(context.Background())
}

func (i GrafanaIntegrationsArgs) ToGrafanaIntegrationsPtrOutputWithContext(ctx context.Context) GrafanaIntegrationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrafanaIntegrationsOutput).ToGrafanaIntegrationsPtrOutputWithContext(ctx)
}









type GrafanaIntegrationsPtrInput interface {
	pulumi.Input

	ToGrafanaIntegrationsPtrOutput() GrafanaIntegrationsPtrOutput
	ToGrafanaIntegrationsPtrOutputWithContext(context.Context) GrafanaIntegrationsPtrOutput
}

type grafanaIntegrationsPtrType GrafanaIntegrationsArgs

func GrafanaIntegrationsPtr(v *GrafanaIntegrationsArgs) GrafanaIntegrationsPtrInput {
	return (*grafanaIntegrationsPtrType)(v)
}

func (*grafanaIntegrationsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GrafanaIntegrations)(nil)).Elem()
}

func (i *grafanaIntegrationsPtrType) ToGrafanaIntegrationsPtrOutput() GrafanaIntegrationsPtrOutput {
	return i.ToGrafanaIntegrationsPtrOutputWithContext(context.Background())
}

func (i *grafanaIntegrationsPtrType) ToGrafanaIntegrationsPtrOutputWithContext(ctx context.Context) GrafanaIntegrationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrafanaIntegrationsPtrOutput)
}

type GrafanaIntegrationsOutput struct{ *pulumi.OutputState }

func (GrafanaIntegrationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GrafanaIntegrations)(nil)).Elem()
}

func (o GrafanaIntegrationsOutput) ToGrafanaIntegrationsOutput() GrafanaIntegrationsOutput {
	return o
}

func (o GrafanaIntegrationsOutput) ToGrafanaIntegrationsOutputWithContext(ctx context.Context) GrafanaIntegrationsOutput {
	return o
}

func (o GrafanaIntegrationsOutput) ToGrafanaIntegrationsPtrOutput() GrafanaIntegrationsPtrOutput {
	return o.ToGrafanaIntegrationsPtrOutputWithContext(context.Background())
}

func (o GrafanaIntegrationsOutput) ToGrafanaIntegrationsPtrOutputWithContext(ctx context.Context) GrafanaIntegrationsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GrafanaIntegrations) *GrafanaIntegrations {
		return &v
	}).(GrafanaIntegrationsPtrOutput)
}

func (o GrafanaIntegrationsOutput) AzureMonitorWorkspaceIntegrations() AzureMonitorWorkspaceIntegrationArrayOutput {
	return o.ApplyT(func(v GrafanaIntegrations) []AzureMonitorWorkspaceIntegration {
		return v.AzureMonitorWorkspaceIntegrations
	}).(AzureMonitorWorkspaceIntegrationArrayOutput)
}

type GrafanaIntegrationsPtrOutput struct{ *pulumi.OutputState }

func (GrafanaIntegrationsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GrafanaIntegrations)(nil)).Elem()
}

func (o GrafanaIntegrationsPtrOutput) ToGrafanaIntegrationsPtrOutput() GrafanaIntegrationsPtrOutput {
	return o
}

func (o GrafanaIntegrationsPtrOutput) ToGrafanaIntegrationsPtrOutputWithContext(ctx context.Context) GrafanaIntegrationsPtrOutput {
	return o
}

func (o GrafanaIntegrationsPtrOutput) Elem() GrafanaIntegrationsOutput {
	return o.ApplyT(func(v *GrafanaIntegrations) GrafanaIntegrations {
		if v != nil {
			return *v
		}
		var ret GrafanaIntegrations
		return ret
	}).(GrafanaIntegrationsOutput)
}

func (o GrafanaIntegrationsPtrOutput) AzureMonitorWorkspaceIntegrations() AzureMonitorWorkspaceIntegrationArrayOutput {
	return o.ApplyT(func(v *GrafanaIntegrations) []AzureMonitorWorkspaceIntegration {
		if v == nil {
			return nil
		}
		return v.AzureMonitorWorkspaceIntegrations
	}).(AzureMonitorWorkspaceIntegrationArrayOutput)
}

type GrafanaIntegrationsResponse struct {
	AzureMonitorWorkspaceIntegrations []AzureMonitorWorkspaceIntegrationResponse `pulumi:"azureMonitorWorkspaceIntegrations"`
}

type GrafanaIntegrationsResponseOutput struct{ *pulumi.OutputState }

func (GrafanaIntegrationsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GrafanaIntegrationsResponse)(nil)).Elem()
}

func (o GrafanaIntegrationsResponseOutput) ToGrafanaIntegrationsResponseOutput() GrafanaIntegrationsResponseOutput {
	return o
}

func (o GrafanaIntegrationsResponseOutput) ToGrafanaIntegrationsResponseOutputWithContext(ctx context.Context) GrafanaIntegrationsResponseOutput {
	return o
}

func (o GrafanaIntegrationsResponseOutput) AzureMonitorWorkspaceIntegrations() AzureMonitorWorkspaceIntegrationResponseArrayOutput {
	return o.ApplyT(func(v GrafanaIntegrationsResponse) []AzureMonitorWorkspaceIntegrationResponse {
		return v.AzureMonitorWorkspaceIntegrations
	}).(AzureMonitorWorkspaceIntegrationResponseArrayOutput)
}

type GrafanaIntegrationsResponsePtrOutput struct{ *pulumi.OutputState }

func (GrafanaIntegrationsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GrafanaIntegrationsResponse)(nil)).Elem()
}

func (o GrafanaIntegrationsResponsePtrOutput) ToGrafanaIntegrationsResponsePtrOutput() GrafanaIntegrationsResponsePtrOutput {
	return o
}

func (o GrafanaIntegrationsResponsePtrOutput) ToGrafanaIntegrationsResponsePtrOutputWithContext(ctx context.Context) GrafanaIntegrationsResponsePtrOutput {
	return o
}

func (o GrafanaIntegrationsResponsePtrOutput) Elem() GrafanaIntegrationsResponseOutput {
	return o.ApplyT(func(v *GrafanaIntegrationsResponse) GrafanaIntegrationsResponse {
		if v != nil {
			return *v
		}
		var ret GrafanaIntegrationsResponse
		return ret
	}).(GrafanaIntegrationsResponseOutput)
}

func (o GrafanaIntegrationsResponsePtrOutput) AzureMonitorWorkspaceIntegrations() AzureMonitorWorkspaceIntegrationResponseArrayOutput {
	return o.ApplyT(func(v *GrafanaIntegrationsResponse) []AzureMonitorWorkspaceIntegrationResponse {
		if v == nil {
			return nil
		}
		return v.AzureMonitorWorkspaceIntegrations
	}).(AzureMonitorWorkspaceIntegrationResponseArrayOutput)
}

type ManagedGrafanaProperties struct {
	ApiKey                            *string              `pulumi:"apiKey"`
	AutoGeneratedDomainNameLabelScope *string              `pulumi:"autoGeneratedDomainNameLabelScope"`
	DeterministicOutboundIP           *string              `pulumi:"deterministicOutboundIP"`
	GrafanaIntegrations               *GrafanaIntegrations `pulumi:"grafanaIntegrations"`
	PublicNetworkAccess               *string              `pulumi:"publicNetworkAccess"`
	ZoneRedundancy                    *string              `pulumi:"zoneRedundancy"`
}





type ManagedGrafanaPropertiesInput interface {
	pulumi.Input

	ToManagedGrafanaPropertiesOutput() ManagedGrafanaPropertiesOutput
	ToManagedGrafanaPropertiesOutputWithContext(context.Context) ManagedGrafanaPropertiesOutput
}

type ManagedGrafanaPropertiesArgs struct {
	ApiKey                            pulumi.StringPtrInput       `pulumi:"apiKey"`
	AutoGeneratedDomainNameLabelScope pulumi.StringPtrInput       `pulumi:"autoGeneratedDomainNameLabelScope"`
	DeterministicOutboundIP           pulumi.StringPtrInput       `pulumi:"deterministicOutboundIP"`
	GrafanaIntegrations               GrafanaIntegrationsPtrInput `pulumi:"grafanaIntegrations"`
	PublicNetworkAccess               pulumi.StringPtrInput       `pulumi:"publicNetworkAccess"`
	ZoneRedundancy                    pulumi.StringPtrInput       `pulumi:"zoneRedundancy"`
}

func (ManagedGrafanaPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedGrafanaProperties)(nil)).Elem()
}

func (i ManagedGrafanaPropertiesArgs) ToManagedGrafanaPropertiesOutput() ManagedGrafanaPropertiesOutput {
	return i.ToManagedGrafanaPropertiesOutputWithContext(context.Background())
}

func (i ManagedGrafanaPropertiesArgs) ToManagedGrafanaPropertiesOutputWithContext(ctx context.Context) ManagedGrafanaPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedGrafanaPropertiesOutput)
}

func (i ManagedGrafanaPropertiesArgs) ToManagedGrafanaPropertiesPtrOutput() ManagedGrafanaPropertiesPtrOutput {
	return i.ToManagedGrafanaPropertiesPtrOutputWithContext(context.Background())
}

func (i ManagedGrafanaPropertiesArgs) ToManagedGrafanaPropertiesPtrOutputWithContext(ctx context.Context) ManagedGrafanaPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedGrafanaPropertiesOutput).ToManagedGrafanaPropertiesPtrOutputWithContext(ctx)
}









type ManagedGrafanaPropertiesPtrInput interface {
	pulumi.Input

	ToManagedGrafanaPropertiesPtrOutput() ManagedGrafanaPropertiesPtrOutput
	ToManagedGrafanaPropertiesPtrOutputWithContext(context.Context) ManagedGrafanaPropertiesPtrOutput
}

type managedGrafanaPropertiesPtrType ManagedGrafanaPropertiesArgs

func ManagedGrafanaPropertiesPtr(v *ManagedGrafanaPropertiesArgs) ManagedGrafanaPropertiesPtrInput {
	return (*managedGrafanaPropertiesPtrType)(v)
}

func (*managedGrafanaPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedGrafanaProperties)(nil)).Elem()
}

func (i *managedGrafanaPropertiesPtrType) ToManagedGrafanaPropertiesPtrOutput() ManagedGrafanaPropertiesPtrOutput {
	return i.ToManagedGrafanaPropertiesPtrOutputWithContext(context.Background())
}

func (i *managedGrafanaPropertiesPtrType) ToManagedGrafanaPropertiesPtrOutputWithContext(ctx context.Context) ManagedGrafanaPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedGrafanaPropertiesPtrOutput)
}

type ManagedGrafanaPropertiesOutput struct{ *pulumi.OutputState }

func (ManagedGrafanaPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedGrafanaProperties)(nil)).Elem()
}

func (o ManagedGrafanaPropertiesOutput) ToManagedGrafanaPropertiesOutput() ManagedGrafanaPropertiesOutput {
	return o
}

func (o ManagedGrafanaPropertiesOutput) ToManagedGrafanaPropertiesOutputWithContext(ctx context.Context) ManagedGrafanaPropertiesOutput {
	return o
}

func (o ManagedGrafanaPropertiesOutput) ToManagedGrafanaPropertiesPtrOutput() ManagedGrafanaPropertiesPtrOutput {
	return o.ToManagedGrafanaPropertiesPtrOutputWithContext(context.Background())
}

func (o ManagedGrafanaPropertiesOutput) ToManagedGrafanaPropertiesPtrOutputWithContext(ctx context.Context) ManagedGrafanaPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedGrafanaProperties) *ManagedGrafanaProperties {
		return &v
	}).(ManagedGrafanaPropertiesPtrOutput)
}

func (o ManagedGrafanaPropertiesOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedGrafanaProperties) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o ManagedGrafanaPropertiesOutput) AutoGeneratedDomainNameLabelScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedGrafanaProperties) *string { return v.AutoGeneratedDomainNameLabelScope }).(pulumi.StringPtrOutput)
}

func (o ManagedGrafanaPropertiesOutput) DeterministicOutboundIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedGrafanaProperties) *string { return v.DeterministicOutboundIP }).(pulumi.StringPtrOutput)
}

func (o ManagedGrafanaPropertiesOutput) GrafanaIntegrations() GrafanaIntegrationsPtrOutput {
	return o.ApplyT(func(v ManagedGrafanaProperties) *GrafanaIntegrations { return v.GrafanaIntegrations }).(GrafanaIntegrationsPtrOutput)
}

func (o ManagedGrafanaPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedGrafanaProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o ManagedGrafanaPropertiesOutput) ZoneRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedGrafanaProperties) *string { return v.ZoneRedundancy }).(pulumi.StringPtrOutput)
}

type ManagedGrafanaPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ManagedGrafanaPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedGrafanaProperties)(nil)).Elem()
}

func (o ManagedGrafanaPropertiesPtrOutput) ToManagedGrafanaPropertiesPtrOutput() ManagedGrafanaPropertiesPtrOutput {
	return o
}

func (o ManagedGrafanaPropertiesPtrOutput) ToManagedGrafanaPropertiesPtrOutputWithContext(ctx context.Context) ManagedGrafanaPropertiesPtrOutput {
	return o
}

func (o ManagedGrafanaPropertiesPtrOutput) Elem() ManagedGrafanaPropertiesOutput {
	return o.ApplyT(func(v *ManagedGrafanaProperties) ManagedGrafanaProperties {
		if v != nil {
			return *v
		}
		var ret ManagedGrafanaProperties
		return ret
	}).(ManagedGrafanaPropertiesOutput)
}

func (o ManagedGrafanaPropertiesPtrOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedGrafanaProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApiKey
	}).(pulumi.StringPtrOutput)
}

func (o ManagedGrafanaPropertiesPtrOutput) AutoGeneratedDomainNameLabelScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedGrafanaProperties) *string {
		if v == nil {
			return nil
		}
		return v.AutoGeneratedDomainNameLabelScope
	}).(pulumi.StringPtrOutput)
}

func (o ManagedGrafanaPropertiesPtrOutput) DeterministicOutboundIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedGrafanaProperties) *string {
		if v == nil {
			return nil
		}
		return v.DeterministicOutboundIP
	}).(pulumi.StringPtrOutput)
}

func (o ManagedGrafanaPropertiesPtrOutput) GrafanaIntegrations() GrafanaIntegrationsPtrOutput {
	return o.ApplyT(func(v *ManagedGrafanaProperties) *GrafanaIntegrations {
		if v == nil {
			return nil
		}
		return v.GrafanaIntegrations
	}).(GrafanaIntegrationsPtrOutput)
}

func (o ManagedGrafanaPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedGrafanaProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o ManagedGrafanaPropertiesPtrOutput) ZoneRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedGrafanaProperties) *string {
		if v == nil {
			return nil
		}
		return v.ZoneRedundancy
	}).(pulumi.StringPtrOutput)
}

type ManagedGrafanaPropertiesResponse struct {
	ApiKey                            *string                             `pulumi:"apiKey"`
	AutoGeneratedDomainNameLabelScope *string                             `pulumi:"autoGeneratedDomainNameLabelScope"`
	DeterministicOutboundIP           *string                             `pulumi:"deterministicOutboundIP"`
	Endpoint                          string                              `pulumi:"endpoint"`
	GrafanaIntegrations               *GrafanaIntegrationsResponse        `pulumi:"grafanaIntegrations"`
	GrafanaVersion                    string                              `pulumi:"grafanaVersion"`
	OutboundIPs                       []string                            `pulumi:"outboundIPs"`
	PrivateEndpointConnections        []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
	PublicNetworkAccess               *string                             `pulumi:"publicNetworkAccess"`
	ZoneRedundancy                    *string                             `pulumi:"zoneRedundancy"`
}

type ManagedGrafanaPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManagedGrafanaPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedGrafanaPropertiesResponse)(nil)).Elem()
}

func (o ManagedGrafanaPropertiesResponseOutput) ToManagedGrafanaPropertiesResponseOutput() ManagedGrafanaPropertiesResponseOutput {
	return o
}

func (o ManagedGrafanaPropertiesResponseOutput) ToManagedGrafanaPropertiesResponseOutputWithContext(ctx context.Context) ManagedGrafanaPropertiesResponseOutput {
	return o
}

func (o ManagedGrafanaPropertiesResponseOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedGrafanaPropertiesResponse) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o ManagedGrafanaPropertiesResponseOutput) AutoGeneratedDomainNameLabelScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedGrafanaPropertiesResponse) *string { return v.AutoGeneratedDomainNameLabelScope }).(pulumi.StringPtrOutput)
}

func (o ManagedGrafanaPropertiesResponseOutput) DeterministicOutboundIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedGrafanaPropertiesResponse) *string { return v.DeterministicOutboundIP }).(pulumi.StringPtrOutput)
}

func (o ManagedGrafanaPropertiesResponseOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedGrafanaPropertiesResponse) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o ManagedGrafanaPropertiesResponseOutput) GrafanaIntegrations() GrafanaIntegrationsResponsePtrOutput {
	return o.ApplyT(func(v ManagedGrafanaPropertiesResponse) *GrafanaIntegrationsResponse { return v.GrafanaIntegrations }).(GrafanaIntegrationsResponsePtrOutput)
}

func (o ManagedGrafanaPropertiesResponseOutput) GrafanaVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedGrafanaPropertiesResponse) string { return v.GrafanaVersion }).(pulumi.StringOutput)
}

func (o ManagedGrafanaPropertiesResponseOutput) OutboundIPs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedGrafanaPropertiesResponse) []string { return v.OutboundIPs }).(pulumi.StringArrayOutput)
}

func (o ManagedGrafanaPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v ManagedGrafanaPropertiesResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o ManagedGrafanaPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedGrafanaPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedGrafanaPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedGrafanaPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o ManagedGrafanaPropertiesResponseOutput) ZoneRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedGrafanaPropertiesResponse) *string { return v.ZoneRedundancy }).(pulumi.StringPtrOutput)
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

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(AzureMonitorWorkspaceIntegrationOutput{})
	pulumi.RegisterOutputType(AzureMonitorWorkspaceIntegrationArrayOutput{})
	pulumi.RegisterOutputType(AzureMonitorWorkspaceIntegrationResponseOutput{})
	pulumi.RegisterOutputType(AzureMonitorWorkspaceIntegrationResponseArrayOutput{})
	pulumi.RegisterOutputType(GrafanaIntegrationsOutput{})
	pulumi.RegisterOutputType(GrafanaIntegrationsPtrOutput{})
	pulumi.RegisterOutputType(GrafanaIntegrationsResponseOutput{})
	pulumi.RegisterOutputType(GrafanaIntegrationsResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedGrafanaPropertiesOutput{})
	pulumi.RegisterOutputType(ManagedGrafanaPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagedGrafanaPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}

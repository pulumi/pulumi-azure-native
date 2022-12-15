


package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ArcConnectivityProperties struct {
	Enabled *bool `pulumi:"enabled"`
}





type ArcConnectivityPropertiesInput interface {
	pulumi.Input

	ToArcConnectivityPropertiesOutput() ArcConnectivityPropertiesOutput
	ToArcConnectivityPropertiesOutputWithContext(context.Context) ArcConnectivityPropertiesOutput
}

type ArcConnectivityPropertiesArgs struct {
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (ArcConnectivityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcConnectivityProperties)(nil)).Elem()
}

func (i ArcConnectivityPropertiesArgs) ToArcConnectivityPropertiesOutput() ArcConnectivityPropertiesOutput {
	return i.ToArcConnectivityPropertiesOutputWithContext(context.Background())
}

func (i ArcConnectivityPropertiesArgs) ToArcConnectivityPropertiesOutputWithContext(ctx context.Context) ArcConnectivityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcConnectivityPropertiesOutput)
}





type ArcConnectivityPropertiesArrayInput interface {
	pulumi.Input

	ToArcConnectivityPropertiesArrayOutput() ArcConnectivityPropertiesArrayOutput
	ToArcConnectivityPropertiesArrayOutputWithContext(context.Context) ArcConnectivityPropertiesArrayOutput
}

type ArcConnectivityPropertiesArray []ArcConnectivityPropertiesInput

func (ArcConnectivityPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArcConnectivityProperties)(nil)).Elem()
}

func (i ArcConnectivityPropertiesArray) ToArcConnectivityPropertiesArrayOutput() ArcConnectivityPropertiesArrayOutput {
	return i.ToArcConnectivityPropertiesArrayOutputWithContext(context.Background())
}

func (i ArcConnectivityPropertiesArray) ToArcConnectivityPropertiesArrayOutputWithContext(ctx context.Context) ArcConnectivityPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcConnectivityPropertiesArrayOutput)
}

type ArcConnectivityPropertiesOutput struct{ *pulumi.OutputState }

func (ArcConnectivityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcConnectivityProperties)(nil)).Elem()
}

func (o ArcConnectivityPropertiesOutput) ToArcConnectivityPropertiesOutput() ArcConnectivityPropertiesOutput {
	return o
}

func (o ArcConnectivityPropertiesOutput) ToArcConnectivityPropertiesOutputWithContext(ctx context.Context) ArcConnectivityPropertiesOutput {
	return o
}

func (o ArcConnectivityPropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ArcConnectivityProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type ArcConnectivityPropertiesArrayOutput struct{ *pulumi.OutputState }

func (ArcConnectivityPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArcConnectivityProperties)(nil)).Elem()
}

func (o ArcConnectivityPropertiesArrayOutput) ToArcConnectivityPropertiesArrayOutput() ArcConnectivityPropertiesArrayOutput {
	return o
}

func (o ArcConnectivityPropertiesArrayOutput) ToArcConnectivityPropertiesArrayOutputWithContext(ctx context.Context) ArcConnectivityPropertiesArrayOutput {
	return o
}

func (o ArcConnectivityPropertiesArrayOutput) Index(i pulumi.IntInput) ArcConnectivityPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArcConnectivityProperties {
		return vs[0].([]ArcConnectivityProperties)[vs[1].(int)]
	}).(ArcConnectivityPropertiesOutput)
}

type ArcConnectivityPropertiesResponse struct {
	Enabled *bool `pulumi:"enabled"`
}

type ArcConnectivityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ArcConnectivityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcConnectivityPropertiesResponse)(nil)).Elem()
}

func (o ArcConnectivityPropertiesResponseOutput) ToArcConnectivityPropertiesResponseOutput() ArcConnectivityPropertiesResponseOutput {
	return o
}

func (o ArcConnectivityPropertiesResponseOutput) ToArcConnectivityPropertiesResponseOutputWithContext(ctx context.Context) ArcConnectivityPropertiesResponseOutput {
	return o
}

func (o ArcConnectivityPropertiesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ArcConnectivityPropertiesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type ArcConnectivityPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (ArcConnectivityPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArcConnectivityPropertiesResponse)(nil)).Elem()
}

func (o ArcConnectivityPropertiesResponseArrayOutput) ToArcConnectivityPropertiesResponseArrayOutput() ArcConnectivityPropertiesResponseArrayOutput {
	return o
}

func (o ArcConnectivityPropertiesResponseArrayOutput) ToArcConnectivityPropertiesResponseArrayOutputWithContext(ctx context.Context) ArcConnectivityPropertiesResponseArrayOutput {
	return o
}

func (o ArcConnectivityPropertiesResponseArrayOutput) Index(i pulumi.IntInput) ArcConnectivityPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArcConnectivityPropertiesResponse {
		return vs[0].([]ArcConnectivityPropertiesResponse)[vs[1].(int)]
	}).(ArcConnectivityPropertiesResponseOutput)
}

type ClusterDesiredProperties struct {
	DiagnosticLevel           *string `pulumi:"diagnosticLevel"`
	WindowsServerSubscription *string `pulumi:"windowsServerSubscription"`
}





type ClusterDesiredPropertiesInput interface {
	pulumi.Input

	ToClusterDesiredPropertiesOutput() ClusterDesiredPropertiesOutput
	ToClusterDesiredPropertiesOutputWithContext(context.Context) ClusterDesiredPropertiesOutput
}

type ClusterDesiredPropertiesArgs struct {
	DiagnosticLevel           pulumi.StringPtrInput `pulumi:"diagnosticLevel"`
	WindowsServerSubscription pulumi.StringPtrInput `pulumi:"windowsServerSubscription"`
}

func (ClusterDesiredPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterDesiredProperties)(nil)).Elem()
}

func (i ClusterDesiredPropertiesArgs) ToClusterDesiredPropertiesOutput() ClusterDesiredPropertiesOutput {
	return i.ToClusterDesiredPropertiesOutputWithContext(context.Background())
}

func (i ClusterDesiredPropertiesArgs) ToClusterDesiredPropertiesOutputWithContext(ctx context.Context) ClusterDesiredPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterDesiredPropertiesOutput)
}

func (i ClusterDesiredPropertiesArgs) ToClusterDesiredPropertiesPtrOutput() ClusterDesiredPropertiesPtrOutput {
	return i.ToClusterDesiredPropertiesPtrOutputWithContext(context.Background())
}

func (i ClusterDesiredPropertiesArgs) ToClusterDesiredPropertiesPtrOutputWithContext(ctx context.Context) ClusterDesiredPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterDesiredPropertiesOutput).ToClusterDesiredPropertiesPtrOutputWithContext(ctx)
}









type ClusterDesiredPropertiesPtrInput interface {
	pulumi.Input

	ToClusterDesiredPropertiesPtrOutput() ClusterDesiredPropertiesPtrOutput
	ToClusterDesiredPropertiesPtrOutputWithContext(context.Context) ClusterDesiredPropertiesPtrOutput
}

type clusterDesiredPropertiesPtrType ClusterDesiredPropertiesArgs

func ClusterDesiredPropertiesPtr(v *ClusterDesiredPropertiesArgs) ClusterDesiredPropertiesPtrInput {
	return (*clusterDesiredPropertiesPtrType)(v)
}

func (*clusterDesiredPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterDesiredProperties)(nil)).Elem()
}

func (i *clusterDesiredPropertiesPtrType) ToClusterDesiredPropertiesPtrOutput() ClusterDesiredPropertiesPtrOutput {
	return i.ToClusterDesiredPropertiesPtrOutputWithContext(context.Background())
}

func (i *clusterDesiredPropertiesPtrType) ToClusterDesiredPropertiesPtrOutputWithContext(ctx context.Context) ClusterDesiredPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterDesiredPropertiesPtrOutput)
}

type ClusterDesiredPropertiesOutput struct{ *pulumi.OutputState }

func (ClusterDesiredPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterDesiredProperties)(nil)).Elem()
}

func (o ClusterDesiredPropertiesOutput) ToClusterDesiredPropertiesOutput() ClusterDesiredPropertiesOutput {
	return o
}

func (o ClusterDesiredPropertiesOutput) ToClusterDesiredPropertiesOutputWithContext(ctx context.Context) ClusterDesiredPropertiesOutput {
	return o
}

func (o ClusterDesiredPropertiesOutput) ToClusterDesiredPropertiesPtrOutput() ClusterDesiredPropertiesPtrOutput {
	return o.ToClusterDesiredPropertiesPtrOutputWithContext(context.Background())
}

func (o ClusterDesiredPropertiesOutput) ToClusterDesiredPropertiesPtrOutputWithContext(ctx context.Context) ClusterDesiredPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterDesiredProperties) *ClusterDesiredProperties {
		return &v
	}).(ClusterDesiredPropertiesPtrOutput)
}

func (o ClusterDesiredPropertiesOutput) DiagnosticLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterDesiredProperties) *string { return v.DiagnosticLevel }).(pulumi.StringPtrOutput)
}

func (o ClusterDesiredPropertiesOutput) WindowsServerSubscription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterDesiredProperties) *string { return v.WindowsServerSubscription }).(pulumi.StringPtrOutput)
}

type ClusterDesiredPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ClusterDesiredPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterDesiredProperties)(nil)).Elem()
}

func (o ClusterDesiredPropertiesPtrOutput) ToClusterDesiredPropertiesPtrOutput() ClusterDesiredPropertiesPtrOutput {
	return o
}

func (o ClusterDesiredPropertiesPtrOutput) ToClusterDesiredPropertiesPtrOutputWithContext(ctx context.Context) ClusterDesiredPropertiesPtrOutput {
	return o
}

func (o ClusterDesiredPropertiesPtrOutput) Elem() ClusterDesiredPropertiesOutput {
	return o.ApplyT(func(v *ClusterDesiredProperties) ClusterDesiredProperties {
		if v != nil {
			return *v
		}
		var ret ClusterDesiredProperties
		return ret
	}).(ClusterDesiredPropertiesOutput)
}

func (o ClusterDesiredPropertiesPtrOutput) DiagnosticLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterDesiredProperties) *string {
		if v == nil {
			return nil
		}
		return v.DiagnosticLevel
	}).(pulumi.StringPtrOutput)
}

func (o ClusterDesiredPropertiesPtrOutput) WindowsServerSubscription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterDesiredProperties) *string {
		if v == nil {
			return nil
		}
		return v.WindowsServerSubscription
	}).(pulumi.StringPtrOutput)
}

type ClusterDesiredPropertiesResponse struct {
	DiagnosticLevel           *string `pulumi:"diagnosticLevel"`
	WindowsServerSubscription *string `pulumi:"windowsServerSubscription"`
}

type ClusterDesiredPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ClusterDesiredPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterDesiredPropertiesResponse)(nil)).Elem()
}

func (o ClusterDesiredPropertiesResponseOutput) ToClusterDesiredPropertiesResponseOutput() ClusterDesiredPropertiesResponseOutput {
	return o
}

func (o ClusterDesiredPropertiesResponseOutput) ToClusterDesiredPropertiesResponseOutputWithContext(ctx context.Context) ClusterDesiredPropertiesResponseOutput {
	return o
}

func (o ClusterDesiredPropertiesResponseOutput) DiagnosticLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterDesiredPropertiesResponse) *string { return v.DiagnosticLevel }).(pulumi.StringPtrOutput)
}

func (o ClusterDesiredPropertiesResponseOutput) WindowsServerSubscription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterDesiredPropertiesResponse) *string { return v.WindowsServerSubscription }).(pulumi.StringPtrOutput)
}

type ClusterDesiredPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterDesiredPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterDesiredPropertiesResponse)(nil)).Elem()
}

func (o ClusterDesiredPropertiesResponsePtrOutput) ToClusterDesiredPropertiesResponsePtrOutput() ClusterDesiredPropertiesResponsePtrOutput {
	return o
}

func (o ClusterDesiredPropertiesResponsePtrOutput) ToClusterDesiredPropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterDesiredPropertiesResponsePtrOutput {
	return o
}

func (o ClusterDesiredPropertiesResponsePtrOutput) Elem() ClusterDesiredPropertiesResponseOutput {
	return o.ApplyT(func(v *ClusterDesiredPropertiesResponse) ClusterDesiredPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ClusterDesiredPropertiesResponse
		return ret
	}).(ClusterDesiredPropertiesResponseOutput)
}

func (o ClusterDesiredPropertiesResponsePtrOutput) DiagnosticLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterDesiredPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiagnosticLevel
	}).(pulumi.StringPtrOutput)
}

func (o ClusterDesiredPropertiesResponsePtrOutput) WindowsServerSubscription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterDesiredPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.WindowsServerSubscription
	}).(pulumi.StringPtrOutput)
}

type ClusterNodeResponse struct {
	CoreCount                 float64 `pulumi:"coreCount"`
	EhcResourceId             string  `pulumi:"ehcResourceId"`
	Id                        float64 `pulumi:"id"`
	Manufacturer              string  `pulumi:"manufacturer"`
	MemoryInGiB               float64 `pulumi:"memoryInGiB"`
	Model                     string  `pulumi:"model"`
	Name                      string  `pulumi:"name"`
	NodeType                  string  `pulumi:"nodeType"`
	OsDisplayVersion          string  `pulumi:"osDisplayVersion"`
	OsName                    string  `pulumi:"osName"`
	OsVersion                 string  `pulumi:"osVersion"`
	SerialNumber              string  `pulumi:"serialNumber"`
	WindowsServerSubscription string  `pulumi:"windowsServerSubscription"`
}

type ClusterNodeResponseOutput struct{ *pulumi.OutputState }

func (ClusterNodeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterNodeResponse)(nil)).Elem()
}

func (o ClusterNodeResponseOutput) ToClusterNodeResponseOutput() ClusterNodeResponseOutput {
	return o
}

func (o ClusterNodeResponseOutput) ToClusterNodeResponseOutputWithContext(ctx context.Context) ClusterNodeResponseOutput {
	return o
}

func (o ClusterNodeResponseOutput) CoreCount() pulumi.Float64Output {
	return o.ApplyT(func(v ClusterNodeResponse) float64 { return v.CoreCount }).(pulumi.Float64Output)
}

func (o ClusterNodeResponseOutput) EhcResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.EhcResourceId }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) Id() pulumi.Float64Output {
	return o.ApplyT(func(v ClusterNodeResponse) float64 { return v.Id }).(pulumi.Float64Output)
}

func (o ClusterNodeResponseOutput) Manufacturer() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.Manufacturer }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) MemoryInGiB() pulumi.Float64Output {
	return o.ApplyT(func(v ClusterNodeResponse) float64 { return v.MemoryInGiB }).(pulumi.Float64Output)
}

func (o ClusterNodeResponseOutput) Model() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.Model }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.NodeType }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) OsDisplayVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.OsDisplayVersion }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.OsName }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.OsVersion }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) WindowsServerSubscription() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.WindowsServerSubscription }).(pulumi.StringOutput)
}

type ClusterNodeResponseArrayOutput struct{ *pulumi.OutputState }

func (ClusterNodeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterNodeResponse)(nil)).Elem()
}

func (o ClusterNodeResponseArrayOutput) ToClusterNodeResponseArrayOutput() ClusterNodeResponseArrayOutput {
	return o
}

func (o ClusterNodeResponseArrayOutput) ToClusterNodeResponseArrayOutputWithContext(ctx context.Context) ClusterNodeResponseArrayOutput {
	return o
}

func (o ClusterNodeResponseArrayOutput) Index(i pulumi.IntInput) ClusterNodeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterNodeResponse {
		return vs[0].([]ClusterNodeResponse)[vs[1].(int)]
	}).(ClusterNodeResponseOutput)
}

type ClusterReportedPropertiesResponse struct {
	ClusterId             string                `pulumi:"clusterId"`
	ClusterName           string                `pulumi:"clusterName"`
	ClusterVersion        string                `pulumi:"clusterVersion"`
	DiagnosticLevel       *string               `pulumi:"diagnosticLevel"`
	ImdsAttestation       string                `pulumi:"imdsAttestation"`
	LastUpdated           string                `pulumi:"lastUpdated"`
	Nodes                 []ClusterNodeResponse `pulumi:"nodes"`
	SupportedCapabilities []string              `pulumi:"supportedCapabilities"`
}

type ClusterReportedPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ClusterReportedPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterReportedPropertiesResponse)(nil)).Elem()
}

func (o ClusterReportedPropertiesResponseOutput) ToClusterReportedPropertiesResponseOutput() ClusterReportedPropertiesResponseOutput {
	return o
}

func (o ClusterReportedPropertiesResponseOutput) ToClusterReportedPropertiesResponseOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponseOutput {
	return o
}

func (o ClusterReportedPropertiesResponseOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o ClusterReportedPropertiesResponseOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.ClusterName }).(pulumi.StringOutput)
}

func (o ClusterReportedPropertiesResponseOutput) ClusterVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.ClusterVersion }).(pulumi.StringOutput)
}

func (o ClusterReportedPropertiesResponseOutput) DiagnosticLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) *string { return v.DiagnosticLevel }).(pulumi.StringPtrOutput)
}

func (o ClusterReportedPropertiesResponseOutput) ImdsAttestation() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.ImdsAttestation }).(pulumi.StringOutput)
}

func (o ClusterReportedPropertiesResponseOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.LastUpdated }).(pulumi.StringOutput)
}

func (o ClusterReportedPropertiesResponseOutput) Nodes() ClusterNodeResponseArrayOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) []ClusterNodeResponse { return v.Nodes }).(ClusterNodeResponseArrayOutput)
}

func (o ClusterReportedPropertiesResponseOutput) SupportedCapabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) []string { return v.SupportedCapabilities }).(pulumi.StringArrayOutput)
}

type PerNodeExtensionStateResponse struct {
	Extension string `pulumi:"extension"`
	Name      string `pulumi:"name"`
	State     string `pulumi:"state"`
}

type PerNodeExtensionStateResponseOutput struct{ *pulumi.OutputState }

func (PerNodeExtensionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerNodeExtensionStateResponse)(nil)).Elem()
}

func (o PerNodeExtensionStateResponseOutput) ToPerNodeExtensionStateResponseOutput() PerNodeExtensionStateResponseOutput {
	return o
}

func (o PerNodeExtensionStateResponseOutput) ToPerNodeExtensionStateResponseOutputWithContext(ctx context.Context) PerNodeExtensionStateResponseOutput {
	return o
}

func (o PerNodeExtensionStateResponseOutput) Extension() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeExtensionStateResponse) string { return v.Extension }).(pulumi.StringOutput)
}

func (o PerNodeExtensionStateResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeExtensionStateResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PerNodeExtensionStateResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeExtensionStateResponse) string { return v.State }).(pulumi.StringOutput)
}

type PerNodeExtensionStateResponseArrayOutput struct{ *pulumi.OutputState }

func (PerNodeExtensionStateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerNodeExtensionStateResponse)(nil)).Elem()
}

func (o PerNodeExtensionStateResponseArrayOutput) ToPerNodeExtensionStateResponseArrayOutput() PerNodeExtensionStateResponseArrayOutput {
	return o
}

func (o PerNodeExtensionStateResponseArrayOutput) ToPerNodeExtensionStateResponseArrayOutputWithContext(ctx context.Context) PerNodeExtensionStateResponseArrayOutput {
	return o
}

func (o PerNodeExtensionStateResponseArrayOutput) Index(i pulumi.IntInput) PerNodeExtensionStateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerNodeExtensionStateResponse {
		return vs[0].([]PerNodeExtensionStateResponse)[vs[1].(int)]
	}).(PerNodeExtensionStateResponseOutput)
}

type PerNodeStateResponse struct {
	ArcInstance string `pulumi:"arcInstance"`
	Name        string `pulumi:"name"`
	State       string `pulumi:"state"`
}

type PerNodeStateResponseOutput struct{ *pulumi.OutputState }

func (PerNodeStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerNodeStateResponse)(nil)).Elem()
}

func (o PerNodeStateResponseOutput) ToPerNodeStateResponseOutput() PerNodeStateResponseOutput {
	return o
}

func (o PerNodeStateResponseOutput) ToPerNodeStateResponseOutputWithContext(ctx context.Context) PerNodeStateResponseOutput {
	return o
}

func (o PerNodeStateResponseOutput) ArcInstance() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeStateResponse) string { return v.ArcInstance }).(pulumi.StringOutput)
}

func (o PerNodeStateResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeStateResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PerNodeStateResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeStateResponse) string { return v.State }).(pulumi.StringOutput)
}

type PerNodeStateResponseArrayOutput struct{ *pulumi.OutputState }

func (PerNodeStateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerNodeStateResponse)(nil)).Elem()
}

func (o PerNodeStateResponseArrayOutput) ToPerNodeStateResponseArrayOutput() PerNodeStateResponseArrayOutput {
	return o
}

func (o PerNodeStateResponseArrayOutput) ToPerNodeStateResponseArrayOutputWithContext(ctx context.Context) PerNodeStateResponseArrayOutput {
	return o
}

func (o PerNodeStateResponseArrayOutput) Index(i pulumi.IntInput) PerNodeStateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerNodeStateResponse {
		return vs[0].([]PerNodeStateResponse)[vs[1].(int)]
	}).(PerNodeStateResponseOutput)
}

type SoftwareAssuranceProperties struct {
	SoftwareAssuranceIntent *string `pulumi:"softwareAssuranceIntent"`
	SoftwareAssuranceStatus *string `pulumi:"softwareAssuranceStatus"`
}





type SoftwareAssurancePropertiesInput interface {
	pulumi.Input

	ToSoftwareAssurancePropertiesOutput() SoftwareAssurancePropertiesOutput
	ToSoftwareAssurancePropertiesOutputWithContext(context.Context) SoftwareAssurancePropertiesOutput
}

type SoftwareAssurancePropertiesArgs struct {
	SoftwareAssuranceIntent pulumi.StringPtrInput `pulumi:"softwareAssuranceIntent"`
	SoftwareAssuranceStatus pulumi.StringPtrInput `pulumi:"softwareAssuranceStatus"`
}

func (SoftwareAssurancePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftwareAssuranceProperties)(nil)).Elem()
}

func (i SoftwareAssurancePropertiesArgs) ToSoftwareAssurancePropertiesOutput() SoftwareAssurancePropertiesOutput {
	return i.ToSoftwareAssurancePropertiesOutputWithContext(context.Background())
}

func (i SoftwareAssurancePropertiesArgs) ToSoftwareAssurancePropertiesOutputWithContext(ctx context.Context) SoftwareAssurancePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwareAssurancePropertiesOutput)
}

func (i SoftwareAssurancePropertiesArgs) ToSoftwareAssurancePropertiesPtrOutput() SoftwareAssurancePropertiesPtrOutput {
	return i.ToSoftwareAssurancePropertiesPtrOutputWithContext(context.Background())
}

func (i SoftwareAssurancePropertiesArgs) ToSoftwareAssurancePropertiesPtrOutputWithContext(ctx context.Context) SoftwareAssurancePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwareAssurancePropertiesOutput).ToSoftwareAssurancePropertiesPtrOutputWithContext(ctx)
}









type SoftwareAssurancePropertiesPtrInput interface {
	pulumi.Input

	ToSoftwareAssurancePropertiesPtrOutput() SoftwareAssurancePropertiesPtrOutput
	ToSoftwareAssurancePropertiesPtrOutputWithContext(context.Context) SoftwareAssurancePropertiesPtrOutput
}

type softwareAssurancePropertiesPtrType SoftwareAssurancePropertiesArgs

func SoftwareAssurancePropertiesPtr(v *SoftwareAssurancePropertiesArgs) SoftwareAssurancePropertiesPtrInput {
	return (*softwareAssurancePropertiesPtrType)(v)
}

func (*softwareAssurancePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareAssuranceProperties)(nil)).Elem()
}

func (i *softwareAssurancePropertiesPtrType) ToSoftwareAssurancePropertiesPtrOutput() SoftwareAssurancePropertiesPtrOutput {
	return i.ToSoftwareAssurancePropertiesPtrOutputWithContext(context.Background())
}

func (i *softwareAssurancePropertiesPtrType) ToSoftwareAssurancePropertiesPtrOutputWithContext(ctx context.Context) SoftwareAssurancePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwareAssurancePropertiesPtrOutput)
}

type SoftwareAssurancePropertiesOutput struct{ *pulumi.OutputState }

func (SoftwareAssurancePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftwareAssuranceProperties)(nil)).Elem()
}

func (o SoftwareAssurancePropertiesOutput) ToSoftwareAssurancePropertiesOutput() SoftwareAssurancePropertiesOutput {
	return o
}

func (o SoftwareAssurancePropertiesOutput) ToSoftwareAssurancePropertiesOutputWithContext(ctx context.Context) SoftwareAssurancePropertiesOutput {
	return o
}

func (o SoftwareAssurancePropertiesOutput) ToSoftwareAssurancePropertiesPtrOutput() SoftwareAssurancePropertiesPtrOutput {
	return o.ToSoftwareAssurancePropertiesPtrOutputWithContext(context.Background())
}

func (o SoftwareAssurancePropertiesOutput) ToSoftwareAssurancePropertiesPtrOutputWithContext(ctx context.Context) SoftwareAssurancePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SoftwareAssuranceProperties) *SoftwareAssuranceProperties {
		return &v
	}).(SoftwareAssurancePropertiesPtrOutput)
}

func (o SoftwareAssurancePropertiesOutput) SoftwareAssuranceIntent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoftwareAssuranceProperties) *string { return v.SoftwareAssuranceIntent }).(pulumi.StringPtrOutput)
}

func (o SoftwareAssurancePropertiesOutput) SoftwareAssuranceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoftwareAssuranceProperties) *string { return v.SoftwareAssuranceStatus }).(pulumi.StringPtrOutput)
}

type SoftwareAssurancePropertiesPtrOutput struct{ *pulumi.OutputState }

func (SoftwareAssurancePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareAssuranceProperties)(nil)).Elem()
}

func (o SoftwareAssurancePropertiesPtrOutput) ToSoftwareAssurancePropertiesPtrOutput() SoftwareAssurancePropertiesPtrOutput {
	return o
}

func (o SoftwareAssurancePropertiesPtrOutput) ToSoftwareAssurancePropertiesPtrOutputWithContext(ctx context.Context) SoftwareAssurancePropertiesPtrOutput {
	return o
}

func (o SoftwareAssurancePropertiesPtrOutput) Elem() SoftwareAssurancePropertiesOutput {
	return o.ApplyT(func(v *SoftwareAssuranceProperties) SoftwareAssuranceProperties {
		if v != nil {
			return *v
		}
		var ret SoftwareAssuranceProperties
		return ret
	}).(SoftwareAssurancePropertiesOutput)
}

func (o SoftwareAssurancePropertiesPtrOutput) SoftwareAssuranceIntent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoftwareAssuranceProperties) *string {
		if v == nil {
			return nil
		}
		return v.SoftwareAssuranceIntent
	}).(pulumi.StringPtrOutput)
}

func (o SoftwareAssurancePropertiesPtrOutput) SoftwareAssuranceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoftwareAssuranceProperties) *string {
		if v == nil {
			return nil
		}
		return v.SoftwareAssuranceStatus
	}).(pulumi.StringPtrOutput)
}

type SoftwareAssurancePropertiesResponse struct {
	LastUpdated             string  `pulumi:"lastUpdated"`
	SoftwareAssuranceIntent *string `pulumi:"softwareAssuranceIntent"`
	SoftwareAssuranceStatus *string `pulumi:"softwareAssuranceStatus"`
}

type SoftwareAssurancePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SoftwareAssurancePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftwareAssurancePropertiesResponse)(nil)).Elem()
}

func (o SoftwareAssurancePropertiesResponseOutput) ToSoftwareAssurancePropertiesResponseOutput() SoftwareAssurancePropertiesResponseOutput {
	return o
}

func (o SoftwareAssurancePropertiesResponseOutput) ToSoftwareAssurancePropertiesResponseOutputWithContext(ctx context.Context) SoftwareAssurancePropertiesResponseOutput {
	return o
}

func (o SoftwareAssurancePropertiesResponseOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v SoftwareAssurancePropertiesResponse) string { return v.LastUpdated }).(pulumi.StringOutput)
}

func (o SoftwareAssurancePropertiesResponseOutput) SoftwareAssuranceIntent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoftwareAssurancePropertiesResponse) *string { return v.SoftwareAssuranceIntent }).(pulumi.StringPtrOutput)
}

func (o SoftwareAssurancePropertiesResponseOutput) SoftwareAssuranceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoftwareAssurancePropertiesResponse) *string { return v.SoftwareAssuranceStatus }).(pulumi.StringPtrOutput)
}

type SoftwareAssurancePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SoftwareAssurancePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareAssurancePropertiesResponse)(nil)).Elem()
}

func (o SoftwareAssurancePropertiesResponsePtrOutput) ToSoftwareAssurancePropertiesResponsePtrOutput() SoftwareAssurancePropertiesResponsePtrOutput {
	return o
}

func (o SoftwareAssurancePropertiesResponsePtrOutput) ToSoftwareAssurancePropertiesResponsePtrOutputWithContext(ctx context.Context) SoftwareAssurancePropertiesResponsePtrOutput {
	return o
}

func (o SoftwareAssurancePropertiesResponsePtrOutput) Elem() SoftwareAssurancePropertiesResponseOutput {
	return o.ApplyT(func(v *SoftwareAssurancePropertiesResponse) SoftwareAssurancePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SoftwareAssurancePropertiesResponse
		return ret
	}).(SoftwareAssurancePropertiesResponseOutput)
}

func (o SoftwareAssurancePropertiesResponsePtrOutput) LastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoftwareAssurancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdated
	}).(pulumi.StringPtrOutput)
}

func (o SoftwareAssurancePropertiesResponsePtrOutput) SoftwareAssuranceIntent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoftwareAssurancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SoftwareAssuranceIntent
	}).(pulumi.StringPtrOutput)
}

func (o SoftwareAssurancePropertiesResponsePtrOutput) SoftwareAssuranceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoftwareAssurancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SoftwareAssuranceStatus
	}).(pulumi.StringPtrOutput)
}

type Step struct {
	Description        *string `pulumi:"description"`
	EndTimeUtc         *string `pulumi:"endTimeUtc"`
	ErrorMessage       *string `pulumi:"errorMessage"`
	LastUpdatedTimeUtc *string `pulumi:"lastUpdatedTimeUtc"`
	Name               *string `pulumi:"name"`
	StartTimeUtc       *string `pulumi:"startTimeUtc"`
	Status             *string `pulumi:"status"`
	Steps              []Step  `pulumi:"steps"`
}





type StepInput interface {
	pulumi.Input

	ToStepOutput() StepOutput
	ToStepOutputWithContext(context.Context) StepOutput
}

type StepArgs struct {
	Description        pulumi.StringPtrInput `pulumi:"description"`
	EndTimeUtc         pulumi.StringPtrInput `pulumi:"endTimeUtc"`
	ErrorMessage       pulumi.StringPtrInput `pulumi:"errorMessage"`
	LastUpdatedTimeUtc pulumi.StringPtrInput `pulumi:"lastUpdatedTimeUtc"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	StartTimeUtc       pulumi.StringPtrInput `pulumi:"startTimeUtc"`
	Status             pulumi.StringPtrInput `pulumi:"status"`
	Steps              StepArrayInput        `pulumi:"steps"`
}

func (StepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Step)(nil)).Elem()
}

func (i StepArgs) ToStepOutput() StepOutput {
	return i.ToStepOutputWithContext(context.Background())
}

func (i StepArgs) ToStepOutputWithContext(ctx context.Context) StepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepOutput)
}





type StepArrayInput interface {
	pulumi.Input

	ToStepArrayOutput() StepArrayOutput
	ToStepArrayOutputWithContext(context.Context) StepArrayOutput
}

type StepArray []StepInput

func (StepArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Step)(nil)).Elem()
}

func (i StepArray) ToStepArrayOutput() StepArrayOutput {
	return i.ToStepArrayOutputWithContext(context.Background())
}

func (i StepArray) ToStepArrayOutputWithContext(ctx context.Context) StepArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepArrayOutput)
}

type StepOutput struct{ *pulumi.OutputState }

func (StepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Step)(nil)).Elem()
}

func (o StepOutput) ToStepOutput() StepOutput {
	return o
}

func (o StepOutput) ToStepOutputWithContext(ctx context.Context) StepOutput {
	return o
}

func (o StepOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Step) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StepOutput) EndTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Step) *string { return v.EndTimeUtc }).(pulumi.StringPtrOutput)
}

func (o StepOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Step) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o StepOutput) LastUpdatedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Step) *string { return v.LastUpdatedTimeUtc }).(pulumi.StringPtrOutput)
}

func (o StepOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Step) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StepOutput) StartTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Step) *string { return v.StartTimeUtc }).(pulumi.StringPtrOutput)
}

func (o StepOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Step) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o StepOutput) Steps() StepArrayOutput {
	return o.ApplyT(func(v Step) []Step { return v.Steps }).(StepArrayOutput)
}

type StepArrayOutput struct{ *pulumi.OutputState }

func (StepArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Step)(nil)).Elem()
}

func (o StepArrayOutput) ToStepArrayOutput() StepArrayOutput {
	return o
}

func (o StepArrayOutput) ToStepArrayOutputWithContext(ctx context.Context) StepArrayOutput {
	return o
}

func (o StepArrayOutput) Index(i pulumi.IntInput) StepOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Step {
		return vs[0].([]Step)[vs[1].(int)]
	}).(StepOutput)
}

type StepResponse struct {
	Description        *string        `pulumi:"description"`
	EndTimeUtc         *string        `pulumi:"endTimeUtc"`
	ErrorMessage       *string        `pulumi:"errorMessage"`
	LastUpdatedTimeUtc *string        `pulumi:"lastUpdatedTimeUtc"`
	Name               *string        `pulumi:"name"`
	StartTimeUtc       *string        `pulumi:"startTimeUtc"`
	Status             *string        `pulumi:"status"`
	Steps              []StepResponse `pulumi:"steps"`
}

type StepResponseOutput struct{ *pulumi.OutputState }

func (StepResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StepResponse)(nil)).Elem()
}

func (o StepResponseOutput) ToStepResponseOutput() StepResponseOutput {
	return o
}

func (o StepResponseOutput) ToStepResponseOutputWithContext(ctx context.Context) StepResponseOutput {
	return o
}

func (o StepResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StepResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StepResponseOutput) EndTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StepResponse) *string { return v.EndTimeUtc }).(pulumi.StringPtrOutput)
}

func (o StepResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StepResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o StepResponseOutput) LastUpdatedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StepResponse) *string { return v.LastUpdatedTimeUtc }).(pulumi.StringPtrOutput)
}

func (o StepResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StepResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StepResponseOutput) StartTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StepResponse) *string { return v.StartTimeUtc }).(pulumi.StringPtrOutput)
}

func (o StepResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StepResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o StepResponseOutput) Steps() StepResponseArrayOutput {
	return o.ApplyT(func(v StepResponse) []StepResponse { return v.Steps }).(StepResponseArrayOutput)
}

type StepResponseArrayOutput struct{ *pulumi.OutputState }

func (StepResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepResponse)(nil)).Elem()
}

func (o StepResponseArrayOutput) ToStepResponseArrayOutput() StepResponseArrayOutput {
	return o
}

func (o StepResponseArrayOutput) ToStepResponseArrayOutputWithContext(ctx context.Context) StepResponseArrayOutput {
	return o
}

func (o StepResponseArrayOutput) Index(i pulumi.IntInput) StepResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StepResponse {
		return vs[0].([]StepResponse)[vs[1].(int)]
	}).(StepResponseOutput)
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

type UpdatePrerequisite struct {
	PackageName *string `pulumi:"packageName"`
	UpdateType  *string `pulumi:"updateType"`
	Version     *string `pulumi:"version"`
}





type UpdatePrerequisiteInput interface {
	pulumi.Input

	ToUpdatePrerequisiteOutput() UpdatePrerequisiteOutput
	ToUpdatePrerequisiteOutputWithContext(context.Context) UpdatePrerequisiteOutput
}

type UpdatePrerequisiteArgs struct {
	PackageName pulumi.StringPtrInput `pulumi:"packageName"`
	UpdateType  pulumi.StringPtrInput `pulumi:"updateType"`
	Version     pulumi.StringPtrInput `pulumi:"version"`
}

func (UpdatePrerequisiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdatePrerequisite)(nil)).Elem()
}

func (i UpdatePrerequisiteArgs) ToUpdatePrerequisiteOutput() UpdatePrerequisiteOutput {
	return i.ToUpdatePrerequisiteOutputWithContext(context.Background())
}

func (i UpdatePrerequisiteArgs) ToUpdatePrerequisiteOutputWithContext(ctx context.Context) UpdatePrerequisiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdatePrerequisiteOutput)
}





type UpdatePrerequisiteArrayInput interface {
	pulumi.Input

	ToUpdatePrerequisiteArrayOutput() UpdatePrerequisiteArrayOutput
	ToUpdatePrerequisiteArrayOutputWithContext(context.Context) UpdatePrerequisiteArrayOutput
}

type UpdatePrerequisiteArray []UpdatePrerequisiteInput

func (UpdatePrerequisiteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UpdatePrerequisite)(nil)).Elem()
}

func (i UpdatePrerequisiteArray) ToUpdatePrerequisiteArrayOutput() UpdatePrerequisiteArrayOutput {
	return i.ToUpdatePrerequisiteArrayOutputWithContext(context.Background())
}

func (i UpdatePrerequisiteArray) ToUpdatePrerequisiteArrayOutputWithContext(ctx context.Context) UpdatePrerequisiteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdatePrerequisiteArrayOutput)
}

type UpdatePrerequisiteOutput struct{ *pulumi.OutputState }

func (UpdatePrerequisiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdatePrerequisite)(nil)).Elem()
}

func (o UpdatePrerequisiteOutput) ToUpdatePrerequisiteOutput() UpdatePrerequisiteOutput {
	return o
}

func (o UpdatePrerequisiteOutput) ToUpdatePrerequisiteOutputWithContext(ctx context.Context) UpdatePrerequisiteOutput {
	return o
}

func (o UpdatePrerequisiteOutput) PackageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpdatePrerequisite) *string { return v.PackageName }).(pulumi.StringPtrOutput)
}

func (o UpdatePrerequisiteOutput) UpdateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpdatePrerequisite) *string { return v.UpdateType }).(pulumi.StringPtrOutput)
}

func (o UpdatePrerequisiteOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpdatePrerequisite) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type UpdatePrerequisiteArrayOutput struct{ *pulumi.OutputState }

func (UpdatePrerequisiteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UpdatePrerequisite)(nil)).Elem()
}

func (o UpdatePrerequisiteArrayOutput) ToUpdatePrerequisiteArrayOutput() UpdatePrerequisiteArrayOutput {
	return o
}

func (o UpdatePrerequisiteArrayOutput) ToUpdatePrerequisiteArrayOutputWithContext(ctx context.Context) UpdatePrerequisiteArrayOutput {
	return o
}

func (o UpdatePrerequisiteArrayOutput) Index(i pulumi.IntInput) UpdatePrerequisiteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UpdatePrerequisite {
		return vs[0].([]UpdatePrerequisite)[vs[1].(int)]
	}).(UpdatePrerequisiteOutput)
}

type UpdatePrerequisiteResponse struct {
	PackageName *string `pulumi:"packageName"`
	UpdateType  *string `pulumi:"updateType"`
	Version     *string `pulumi:"version"`
}

type UpdatePrerequisiteResponseOutput struct{ *pulumi.OutputState }

func (UpdatePrerequisiteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdatePrerequisiteResponse)(nil)).Elem()
}

func (o UpdatePrerequisiteResponseOutput) ToUpdatePrerequisiteResponseOutput() UpdatePrerequisiteResponseOutput {
	return o
}

func (o UpdatePrerequisiteResponseOutput) ToUpdatePrerequisiteResponseOutputWithContext(ctx context.Context) UpdatePrerequisiteResponseOutput {
	return o
}

func (o UpdatePrerequisiteResponseOutput) PackageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpdatePrerequisiteResponse) *string { return v.PackageName }).(pulumi.StringPtrOutput)
}

func (o UpdatePrerequisiteResponseOutput) UpdateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpdatePrerequisiteResponse) *string { return v.UpdateType }).(pulumi.StringPtrOutput)
}

func (o UpdatePrerequisiteResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpdatePrerequisiteResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type UpdatePrerequisiteResponseArrayOutput struct{ *pulumi.OutputState }

func (UpdatePrerequisiteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UpdatePrerequisiteResponse)(nil)).Elem()
}

func (o UpdatePrerequisiteResponseArrayOutput) ToUpdatePrerequisiteResponseArrayOutput() UpdatePrerequisiteResponseArrayOutput {
	return o
}

func (o UpdatePrerequisiteResponseArrayOutput) ToUpdatePrerequisiteResponseArrayOutputWithContext(ctx context.Context) UpdatePrerequisiteResponseArrayOutput {
	return o
}

func (o UpdatePrerequisiteResponseArrayOutput) Index(i pulumi.IntInput) UpdatePrerequisiteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UpdatePrerequisiteResponse {
		return vs[0].([]UpdatePrerequisiteResponse)[vs[1].(int)]
	}).(UpdatePrerequisiteResponseOutput)
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
	pulumi.RegisterOutputType(ArcConnectivityPropertiesOutput{})
	pulumi.RegisterOutputType(ArcConnectivityPropertiesArrayOutput{})
	pulumi.RegisterOutputType(ArcConnectivityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ArcConnectivityPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(ClusterDesiredPropertiesOutput{})
	pulumi.RegisterOutputType(ClusterDesiredPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ClusterDesiredPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ClusterDesiredPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterNodeResponseOutput{})
	pulumi.RegisterOutputType(ClusterNodeResponseArrayOutput{})
	pulumi.RegisterOutputType(ClusterReportedPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PerNodeExtensionStateResponseOutput{})
	pulumi.RegisterOutputType(PerNodeExtensionStateResponseArrayOutput{})
	pulumi.RegisterOutputType(PerNodeStateResponseOutput{})
	pulumi.RegisterOutputType(PerNodeStateResponseArrayOutput{})
	pulumi.RegisterOutputType(SoftwareAssurancePropertiesOutput{})
	pulumi.RegisterOutputType(SoftwareAssurancePropertiesPtrOutput{})
	pulumi.RegisterOutputType(SoftwareAssurancePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SoftwareAssurancePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(StepOutput{})
	pulumi.RegisterOutputType(StepArrayOutput{})
	pulumi.RegisterOutputType(StepResponseOutput{})
	pulumi.RegisterOutputType(StepResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UpdatePrerequisiteOutput{})
	pulumi.RegisterOutputType(UpdatePrerequisiteArrayOutput{})
	pulumi.RegisterOutputType(UpdatePrerequisiteResponseOutput{})
	pulumi.RegisterOutputType(UpdatePrerequisiteResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}

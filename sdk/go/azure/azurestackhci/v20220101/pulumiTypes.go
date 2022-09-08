


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	Id                        float64 `pulumi:"id"`
	Manufacturer              string  `pulumi:"manufacturer"`
	MemoryInGiB               float64 `pulumi:"memoryInGiB"`
	Model                     string  `pulumi:"model"`
	Name                      string  `pulumi:"name"`
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
	ClusterId       string                `pulumi:"clusterId"`
	ClusterName     string                `pulumi:"clusterName"`
	ClusterVersion  string                `pulumi:"clusterVersion"`
	DiagnosticLevel *string               `pulumi:"diagnosticLevel"`
	ImdsAttestation string                `pulumi:"imdsAttestation"`
	LastUpdated     string                `pulumi:"lastUpdated"`
	Nodes           []ClusterNodeResponse `pulumi:"nodes"`
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

func init() {
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
}

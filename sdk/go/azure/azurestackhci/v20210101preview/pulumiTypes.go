


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterNodeResponse struct {
	CoreCount    float64 `pulumi:"coreCount"`
	Id           float64 `pulumi:"id"`
	Manufacturer string  `pulumi:"manufacturer"`
	MemoryInGiB  float64 `pulumi:"memoryInGiB"`
	Model        string  `pulumi:"model"`
	Name         string  `pulumi:"name"`
	OsName       string  `pulumi:"osName"`
	OsVersion    string  `pulumi:"osVersion"`
	SerialNumber string  `pulumi:"serialNumber"`
}





type ClusterNodeResponseInput interface {
	pulumi.Input

	ToClusterNodeResponseOutput() ClusterNodeResponseOutput
	ToClusterNodeResponseOutputWithContext(context.Context) ClusterNodeResponseOutput
}

type ClusterNodeResponseArgs struct {
	CoreCount    pulumi.Float64Input `pulumi:"coreCount"`
	Id           pulumi.Float64Input `pulumi:"id"`
	Manufacturer pulumi.StringInput  `pulumi:"manufacturer"`
	MemoryInGiB  pulumi.Float64Input `pulumi:"memoryInGiB"`
	Model        pulumi.StringInput  `pulumi:"model"`
	Name         pulumi.StringInput  `pulumi:"name"`
	OsName       pulumi.StringInput  `pulumi:"osName"`
	OsVersion    pulumi.StringInput  `pulumi:"osVersion"`
	SerialNumber pulumi.StringInput  `pulumi:"serialNumber"`
}

func (ClusterNodeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterNodeResponse)(nil)).Elem()
}

func (i ClusterNodeResponseArgs) ToClusterNodeResponseOutput() ClusterNodeResponseOutput {
	return i.ToClusterNodeResponseOutputWithContext(context.Background())
}

func (i ClusterNodeResponseArgs) ToClusterNodeResponseOutputWithContext(ctx context.Context) ClusterNodeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterNodeResponseOutput)
}





type ClusterNodeResponseArrayInput interface {
	pulumi.Input

	ToClusterNodeResponseArrayOutput() ClusterNodeResponseArrayOutput
	ToClusterNodeResponseArrayOutputWithContext(context.Context) ClusterNodeResponseArrayOutput
}

type ClusterNodeResponseArray []ClusterNodeResponseInput

func (ClusterNodeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterNodeResponse)(nil)).Elem()
}

func (i ClusterNodeResponseArray) ToClusterNodeResponseArrayOutput() ClusterNodeResponseArrayOutput {
	return i.ToClusterNodeResponseArrayOutputWithContext(context.Background())
}

func (i ClusterNodeResponseArray) ToClusterNodeResponseArrayOutputWithContext(ctx context.Context) ClusterNodeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterNodeResponseArrayOutput)
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
	ClusterId      string                `pulumi:"clusterId"`
	ClusterName    string                `pulumi:"clusterName"`
	ClusterVersion string                `pulumi:"clusterVersion"`
	LastUpdated    string                `pulumi:"lastUpdated"`
	Nodes          []ClusterNodeResponse `pulumi:"nodes"`
}





type ClusterReportedPropertiesResponseInput interface {
	pulumi.Input

	ToClusterReportedPropertiesResponseOutput() ClusterReportedPropertiesResponseOutput
	ToClusterReportedPropertiesResponseOutputWithContext(context.Context) ClusterReportedPropertiesResponseOutput
}

type ClusterReportedPropertiesResponseArgs struct {
	ClusterId      pulumi.StringInput            `pulumi:"clusterId"`
	ClusterName    pulumi.StringInput            `pulumi:"clusterName"`
	ClusterVersion pulumi.StringInput            `pulumi:"clusterVersion"`
	LastUpdated    pulumi.StringInput            `pulumi:"lastUpdated"`
	Nodes          ClusterNodeResponseArrayInput `pulumi:"nodes"`
}

func (ClusterReportedPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterReportedPropertiesResponse)(nil)).Elem()
}

func (i ClusterReportedPropertiesResponseArgs) ToClusterReportedPropertiesResponseOutput() ClusterReportedPropertiesResponseOutput {
	return i.ToClusterReportedPropertiesResponseOutputWithContext(context.Background())
}

func (i ClusterReportedPropertiesResponseArgs) ToClusterReportedPropertiesResponseOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterReportedPropertiesResponseOutput)
}

func (i ClusterReportedPropertiesResponseArgs) ToClusterReportedPropertiesResponsePtrOutput() ClusterReportedPropertiesResponsePtrOutput {
	return i.ToClusterReportedPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ClusterReportedPropertiesResponseArgs) ToClusterReportedPropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterReportedPropertiesResponseOutput).ToClusterReportedPropertiesResponsePtrOutputWithContext(ctx)
}









type ClusterReportedPropertiesResponsePtrInput interface {
	pulumi.Input

	ToClusterReportedPropertiesResponsePtrOutput() ClusterReportedPropertiesResponsePtrOutput
	ToClusterReportedPropertiesResponsePtrOutputWithContext(context.Context) ClusterReportedPropertiesResponsePtrOutput
}

type clusterReportedPropertiesResponsePtrType ClusterReportedPropertiesResponseArgs

func ClusterReportedPropertiesResponsePtr(v *ClusterReportedPropertiesResponseArgs) ClusterReportedPropertiesResponsePtrInput {
	return (*clusterReportedPropertiesResponsePtrType)(v)
}

func (*clusterReportedPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterReportedPropertiesResponse)(nil)).Elem()
}

func (i *clusterReportedPropertiesResponsePtrType) ToClusterReportedPropertiesResponsePtrOutput() ClusterReportedPropertiesResponsePtrOutput {
	return i.ToClusterReportedPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *clusterReportedPropertiesResponsePtrType) ToClusterReportedPropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterReportedPropertiesResponsePtrOutput)
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

func (o ClusterReportedPropertiesResponseOutput) ToClusterReportedPropertiesResponsePtrOutput() ClusterReportedPropertiesResponsePtrOutput {
	return o.ToClusterReportedPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ClusterReportedPropertiesResponseOutput) ToClusterReportedPropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterReportedPropertiesResponse) *ClusterReportedPropertiesResponse {
		return &v
	}).(ClusterReportedPropertiesResponsePtrOutput)
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

func (o ClusterReportedPropertiesResponseOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.LastUpdated }).(pulumi.StringOutput)
}

func (o ClusterReportedPropertiesResponseOutput) Nodes() ClusterNodeResponseArrayOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) []ClusterNodeResponse { return v.Nodes }).(ClusterNodeResponseArrayOutput)
}

type ClusterReportedPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterReportedPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterReportedPropertiesResponse)(nil)).Elem()
}

func (o ClusterReportedPropertiesResponsePtrOutput) ToClusterReportedPropertiesResponsePtrOutput() ClusterReportedPropertiesResponsePtrOutput {
	return o
}

func (o ClusterReportedPropertiesResponsePtrOutput) ToClusterReportedPropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponsePtrOutput {
	return o
}

func (o ClusterReportedPropertiesResponsePtrOutput) Elem() ClusterReportedPropertiesResponseOutput {
	return o.ApplyT(func(v *ClusterReportedPropertiesResponse) ClusterReportedPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ClusterReportedPropertiesResponse
		return ret
	}).(ClusterReportedPropertiesResponseOutput)
}

func (o ClusterReportedPropertiesResponsePtrOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterReportedPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClusterId
	}).(pulumi.StringPtrOutput)
}

func (o ClusterReportedPropertiesResponsePtrOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterReportedPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClusterName
	}).(pulumi.StringPtrOutput)
}

func (o ClusterReportedPropertiesResponsePtrOutput) ClusterVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterReportedPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClusterVersion
	}).(pulumi.StringPtrOutput)
}

func (o ClusterReportedPropertiesResponsePtrOutput) LastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterReportedPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdated
	}).(pulumi.StringPtrOutput)
}

func (o ClusterReportedPropertiesResponsePtrOutput) Nodes() ClusterNodeResponseArrayOutput {
	return o.ApplyT(func(v *ClusterReportedPropertiesResponse) []ClusterNodeResponse {
		if v == nil {
			return nil
		}
		return v.Nodes
	}).(ClusterNodeResponseArrayOutput)
}

type PerNodeExtensionStateResponse struct {
	Extension string `pulumi:"extension"`
	Name      string `pulumi:"name"`
	State     string `pulumi:"state"`
}





type PerNodeExtensionStateResponseInput interface {
	pulumi.Input

	ToPerNodeExtensionStateResponseOutput() PerNodeExtensionStateResponseOutput
	ToPerNodeExtensionStateResponseOutputWithContext(context.Context) PerNodeExtensionStateResponseOutput
}

type PerNodeExtensionStateResponseArgs struct {
	Extension pulumi.StringInput `pulumi:"extension"`
	Name      pulumi.StringInput `pulumi:"name"`
	State     pulumi.StringInput `pulumi:"state"`
}

func (PerNodeExtensionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PerNodeExtensionStateResponse)(nil)).Elem()
}

func (i PerNodeExtensionStateResponseArgs) ToPerNodeExtensionStateResponseOutput() PerNodeExtensionStateResponseOutput {
	return i.ToPerNodeExtensionStateResponseOutputWithContext(context.Background())
}

func (i PerNodeExtensionStateResponseArgs) ToPerNodeExtensionStateResponseOutputWithContext(ctx context.Context) PerNodeExtensionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerNodeExtensionStateResponseOutput)
}





type PerNodeExtensionStateResponseArrayInput interface {
	pulumi.Input

	ToPerNodeExtensionStateResponseArrayOutput() PerNodeExtensionStateResponseArrayOutput
	ToPerNodeExtensionStateResponseArrayOutputWithContext(context.Context) PerNodeExtensionStateResponseArrayOutput
}

type PerNodeExtensionStateResponseArray []PerNodeExtensionStateResponseInput

func (PerNodeExtensionStateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerNodeExtensionStateResponse)(nil)).Elem()
}

func (i PerNodeExtensionStateResponseArray) ToPerNodeExtensionStateResponseArrayOutput() PerNodeExtensionStateResponseArrayOutput {
	return i.ToPerNodeExtensionStateResponseArrayOutputWithContext(context.Background())
}

func (i PerNodeExtensionStateResponseArray) ToPerNodeExtensionStateResponseArrayOutputWithContext(ctx context.Context) PerNodeExtensionStateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerNodeExtensionStateResponseArrayOutput)
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





type PerNodeStateResponseInput interface {
	pulumi.Input

	ToPerNodeStateResponseOutput() PerNodeStateResponseOutput
	ToPerNodeStateResponseOutputWithContext(context.Context) PerNodeStateResponseOutput
}

type PerNodeStateResponseArgs struct {
	ArcInstance pulumi.StringInput `pulumi:"arcInstance"`
	Name        pulumi.StringInput `pulumi:"name"`
	State       pulumi.StringInput `pulumi:"state"`
}

func (PerNodeStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PerNodeStateResponse)(nil)).Elem()
}

func (i PerNodeStateResponseArgs) ToPerNodeStateResponseOutput() PerNodeStateResponseOutput {
	return i.ToPerNodeStateResponseOutputWithContext(context.Background())
}

func (i PerNodeStateResponseArgs) ToPerNodeStateResponseOutputWithContext(ctx context.Context) PerNodeStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerNodeStateResponseOutput)
}





type PerNodeStateResponseArrayInput interface {
	pulumi.Input

	ToPerNodeStateResponseArrayOutput() PerNodeStateResponseArrayOutput
	ToPerNodeStateResponseArrayOutputWithContext(context.Context) PerNodeStateResponseArrayOutput
}

type PerNodeStateResponseArray []PerNodeStateResponseInput

func (PerNodeStateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerNodeStateResponse)(nil)).Elem()
}

func (i PerNodeStateResponseArray) ToPerNodeStateResponseArrayOutput() PerNodeStateResponseArrayOutput {
	return i.ToPerNodeStateResponseArrayOutputWithContext(context.Background())
}

func (i PerNodeStateResponseArray) ToPerNodeStateResponseArrayOutputWithContext(ctx context.Context) PerNodeStateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerNodeStateResponseArrayOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterNodeResponseInput)(nil)).Elem(), ClusterNodeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterNodeResponseArrayInput)(nil)).Elem(), ClusterNodeResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterReportedPropertiesResponseInput)(nil)).Elem(), ClusterReportedPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterReportedPropertiesResponsePtrInput)(nil)).Elem(), ClusterReportedPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PerNodeExtensionStateResponseInput)(nil)).Elem(), PerNodeExtensionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PerNodeExtensionStateResponseArrayInput)(nil)).Elem(), PerNodeExtensionStateResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PerNodeStateResponseInput)(nil)).Elem(), PerNodeStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PerNodeStateResponseArrayInput)(nil)).Elem(), PerNodeStateResponseArray{})
	pulumi.RegisterOutputType(ClusterNodeResponseOutput{})
	pulumi.RegisterOutputType(ClusterNodeResponseArrayOutput{})
	pulumi.RegisterOutputType(ClusterReportedPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ClusterReportedPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PerNodeExtensionStateResponseOutput{})
	pulumi.RegisterOutputType(PerNodeExtensionStateResponseArrayOutput{})
	pulumi.RegisterOutputType(PerNodeStateResponseOutput{})
	pulumi.RegisterOutputType(PerNodeStateResponseArrayOutput{})
}

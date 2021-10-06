


package v20201001

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

func init() {
	pulumi.RegisterOutputType(ClusterNodeResponseOutput{})
	pulumi.RegisterOutputType(ClusterNodeResponseArrayOutput{})
	pulumi.RegisterOutputType(ClusterReportedPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ClusterReportedPropertiesResponsePtrOutput{})
}

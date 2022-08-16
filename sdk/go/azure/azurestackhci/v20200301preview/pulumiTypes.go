


package v20200301preview

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

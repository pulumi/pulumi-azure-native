


package v20200320

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CircuitResponse struct {
	ExpressRouteID               string `pulumi:"expressRouteID"`
	ExpressRoutePrivatePeeringID string `pulumi:"expressRoutePrivatePeeringID"`
	PrimarySubnet                string `pulumi:"primarySubnet"`
	SecondarySubnet              string `pulumi:"secondarySubnet"`
}

type CircuitResponseOutput struct{ *pulumi.OutputState }

func (CircuitResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CircuitResponse)(nil)).Elem()
}

func (o CircuitResponseOutput) ToCircuitResponseOutput() CircuitResponseOutput {
	return o
}

func (o CircuitResponseOutput) ToCircuitResponseOutputWithContext(ctx context.Context) CircuitResponseOutput {
	return o
}

func (o CircuitResponseOutput) ExpressRouteID() pulumi.StringOutput {
	return o.ApplyT(func(v CircuitResponse) string { return v.ExpressRouteID }).(pulumi.StringOutput)
}

func (o CircuitResponseOutput) ExpressRoutePrivatePeeringID() pulumi.StringOutput {
	return o.ApplyT(func(v CircuitResponse) string { return v.ExpressRoutePrivatePeeringID }).(pulumi.StringOutput)
}

func (o CircuitResponseOutput) PrimarySubnet() pulumi.StringOutput {
	return o.ApplyT(func(v CircuitResponse) string { return v.PrimarySubnet }).(pulumi.StringOutput)
}

func (o CircuitResponseOutput) SecondarySubnet() pulumi.StringOutput {
	return o.ApplyT(func(v CircuitResponse) string { return v.SecondarySubnet }).(pulumi.StringOutput)
}

type CircuitResponsePtrOutput struct{ *pulumi.OutputState }

func (CircuitResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CircuitResponse)(nil)).Elem()
}

func (o CircuitResponsePtrOutput) ToCircuitResponsePtrOutput() CircuitResponsePtrOutput {
	return o
}

func (o CircuitResponsePtrOutput) ToCircuitResponsePtrOutputWithContext(ctx context.Context) CircuitResponsePtrOutput {
	return o
}

func (o CircuitResponsePtrOutput) Elem() CircuitResponseOutput {
	return o.ApplyT(func(v *CircuitResponse) CircuitResponse {
		if v != nil {
			return *v
		}
		var ret CircuitResponse
		return ret
	}).(CircuitResponseOutput)
}

func (o CircuitResponsePtrOutput) ExpressRouteID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CircuitResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ExpressRouteID
	}).(pulumi.StringPtrOutput)
}

func (o CircuitResponsePtrOutput) ExpressRoutePrivatePeeringID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CircuitResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ExpressRoutePrivatePeeringID
	}).(pulumi.StringPtrOutput)
}

func (o CircuitResponsePtrOutput) PrimarySubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CircuitResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrimarySubnet
	}).(pulumi.StringPtrOutput)
}

func (o CircuitResponsePtrOutput) SecondarySubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CircuitResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecondarySubnet
	}).(pulumi.StringPtrOutput)
}

type EndpointsResponse struct {
	HcxCloudManager string `pulumi:"hcxCloudManager"`
	NsxtManager     string `pulumi:"nsxtManager"`
	Vcsa            string `pulumi:"vcsa"`
}

type EndpointsResponseOutput struct{ *pulumi.OutputState }

func (EndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) HcxCloudManager() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.HcxCloudManager }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) NsxtManager() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.NsxtManager }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Vcsa() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Vcsa }).(pulumi.StringOutput)
}

type IdentitySource struct {
	Alias           *string `pulumi:"alias"`
	BaseGroupDN     *string `pulumi:"baseGroupDN"`
	BaseUserDN      *string `pulumi:"baseUserDN"`
	Domain          *string `pulumi:"domain"`
	Name            *string `pulumi:"name"`
	Password        *string `pulumi:"password"`
	PrimaryServer   *string `pulumi:"primaryServer"`
	SecondaryServer *string `pulumi:"secondaryServer"`
	Ssl             *string `pulumi:"ssl"`
	Username        *string `pulumi:"username"`
}





type IdentitySourceInput interface {
	pulumi.Input

	ToIdentitySourceOutput() IdentitySourceOutput
	ToIdentitySourceOutputWithContext(context.Context) IdentitySourceOutput
}

type IdentitySourceArgs struct {
	Alias           pulumi.StringPtrInput `pulumi:"alias"`
	BaseGroupDN     pulumi.StringPtrInput `pulumi:"baseGroupDN"`
	BaseUserDN      pulumi.StringPtrInput `pulumi:"baseUserDN"`
	Domain          pulumi.StringPtrInput `pulumi:"domain"`
	Name            pulumi.StringPtrInput `pulumi:"name"`
	Password        pulumi.StringPtrInput `pulumi:"password"`
	PrimaryServer   pulumi.StringPtrInput `pulumi:"primaryServer"`
	SecondaryServer pulumi.StringPtrInput `pulumi:"secondaryServer"`
	Ssl             pulumi.StringPtrInput `pulumi:"ssl"`
	Username        pulumi.StringPtrInput `pulumi:"username"`
}

func (IdentitySourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySource)(nil)).Elem()
}

func (i IdentitySourceArgs) ToIdentitySourceOutput() IdentitySourceOutput {
	return i.ToIdentitySourceOutputWithContext(context.Background())
}

func (i IdentitySourceArgs) ToIdentitySourceOutputWithContext(ctx context.Context) IdentitySourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitySourceOutput)
}





type IdentitySourceArrayInput interface {
	pulumi.Input

	ToIdentitySourceArrayOutput() IdentitySourceArrayOutput
	ToIdentitySourceArrayOutputWithContext(context.Context) IdentitySourceArrayOutput
}

type IdentitySourceArray []IdentitySourceInput

func (IdentitySourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentitySource)(nil)).Elem()
}

func (i IdentitySourceArray) ToIdentitySourceArrayOutput() IdentitySourceArrayOutput {
	return i.ToIdentitySourceArrayOutputWithContext(context.Background())
}

func (i IdentitySourceArray) ToIdentitySourceArrayOutputWithContext(ctx context.Context) IdentitySourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitySourceArrayOutput)
}

type IdentitySourceOutput struct{ *pulumi.OutputState }

func (IdentitySourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySource)(nil)).Elem()
}

func (o IdentitySourceOutput) ToIdentitySourceOutput() IdentitySourceOutput {
	return o
}

func (o IdentitySourceOutput) ToIdentitySourceOutputWithContext(ctx context.Context) IdentitySourceOutput {
	return o
}

func (o IdentitySourceOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) BaseGroupDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.BaseGroupDN }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) BaseUserDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.BaseUserDN }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) PrimaryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.PrimaryServer }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) SecondaryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.SecondaryServer }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Ssl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Ssl }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type IdentitySourceArrayOutput struct{ *pulumi.OutputState }

func (IdentitySourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentitySource)(nil)).Elem()
}

func (o IdentitySourceArrayOutput) ToIdentitySourceArrayOutput() IdentitySourceArrayOutput {
	return o
}

func (o IdentitySourceArrayOutput) ToIdentitySourceArrayOutputWithContext(ctx context.Context) IdentitySourceArrayOutput {
	return o
}

func (o IdentitySourceArrayOutput) Index(i pulumi.IntInput) IdentitySourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentitySource {
		return vs[0].([]IdentitySource)[vs[1].(int)]
	}).(IdentitySourceOutput)
}

type IdentitySourceResponse struct {
	Alias           *string `pulumi:"alias"`
	BaseGroupDN     *string `pulumi:"baseGroupDN"`
	BaseUserDN      *string `pulumi:"baseUserDN"`
	Domain          *string `pulumi:"domain"`
	Name            *string `pulumi:"name"`
	Password        *string `pulumi:"password"`
	PrimaryServer   *string `pulumi:"primaryServer"`
	SecondaryServer *string `pulumi:"secondaryServer"`
	Ssl             *string `pulumi:"ssl"`
	Username        *string `pulumi:"username"`
}

type IdentitySourceResponseOutput struct{ *pulumi.OutputState }

func (IdentitySourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySourceResponse)(nil)).Elem()
}

func (o IdentitySourceResponseOutput) ToIdentitySourceResponseOutput() IdentitySourceResponseOutput {
	return o
}

func (o IdentitySourceResponseOutput) ToIdentitySourceResponseOutputWithContext(ctx context.Context) IdentitySourceResponseOutput {
	return o
}

func (o IdentitySourceResponseOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) BaseGroupDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.BaseGroupDN }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) BaseUserDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.BaseUserDN }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) PrimaryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.PrimaryServer }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) SecondaryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.SecondaryServer }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Ssl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Ssl }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type IdentitySourceResponseArrayOutput struct{ *pulumi.OutputState }

func (IdentitySourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentitySourceResponse)(nil)).Elem()
}

func (o IdentitySourceResponseArrayOutput) ToIdentitySourceResponseArrayOutput() IdentitySourceResponseArrayOutput {
	return o
}

func (o IdentitySourceResponseArrayOutput) ToIdentitySourceResponseArrayOutputWithContext(ctx context.Context) IdentitySourceResponseArrayOutput {
	return o
}

func (o IdentitySourceResponseArrayOutput) Index(i pulumi.IntInput) IdentitySourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentitySourceResponse {
		return vs[0].([]IdentitySourceResponse)[vs[1].(int)]
	}).(IdentitySourceResponseOutput)
}

type ManagementCluster struct {
	ClusterSize int `pulumi:"clusterSize"`
}





type ManagementClusterInput interface {
	pulumi.Input

	ToManagementClusterOutput() ManagementClusterOutput
	ToManagementClusterOutputWithContext(context.Context) ManagementClusterOutput
}

type ManagementClusterArgs struct {
	ClusterSize pulumi.IntInput `pulumi:"clusterSize"`
}

func (ManagementClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementCluster)(nil)).Elem()
}

func (i ManagementClusterArgs) ToManagementClusterOutput() ManagementClusterOutput {
	return i.ToManagementClusterOutputWithContext(context.Background())
}

func (i ManagementClusterArgs) ToManagementClusterOutputWithContext(ctx context.Context) ManagementClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementClusterOutput)
}

type ManagementClusterOutput struct{ *pulumi.OutputState }

func (ManagementClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementCluster)(nil)).Elem()
}

func (o ManagementClusterOutput) ToManagementClusterOutput() ManagementClusterOutput {
	return o
}

func (o ManagementClusterOutput) ToManagementClusterOutputWithContext(ctx context.Context) ManagementClusterOutput {
	return o
}

func (o ManagementClusterOutput) ClusterSize() pulumi.IntOutput {
	return o.ApplyT(func(v ManagementCluster) int { return v.ClusterSize }).(pulumi.IntOutput)
}

type ManagementClusterResponse struct {
	ClusterId         int      `pulumi:"clusterId"`
	ClusterSize       int      `pulumi:"clusterSize"`
	Hosts             []string `pulumi:"hosts"`
	ProvisioningState string   `pulumi:"provisioningState"`
}

type ManagementClusterResponseOutput struct{ *pulumi.OutputState }

func (ManagementClusterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementClusterResponse)(nil)).Elem()
}

func (o ManagementClusterResponseOutput) ToManagementClusterResponseOutput() ManagementClusterResponseOutput {
	return o
}

func (o ManagementClusterResponseOutput) ToManagementClusterResponseOutputWithContext(ctx context.Context) ManagementClusterResponseOutput {
	return o
}

func (o ManagementClusterResponseOutput) ClusterId() pulumi.IntOutput {
	return o.ApplyT(func(v ManagementClusterResponse) int { return v.ClusterId }).(pulumi.IntOutput)
}

func (o ManagementClusterResponseOutput) ClusterSize() pulumi.IntOutput {
	return o.ApplyT(func(v ManagementClusterResponse) int { return v.ClusterSize }).(pulumi.IntOutput)
}

func (o ManagementClusterResponseOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementClusterResponse) []string { return v.Hosts }).(pulumi.StringArrayOutput)
}

func (o ManagementClusterResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementClusterResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CircuitResponseOutput{})
	pulumi.RegisterOutputType(CircuitResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointsResponseOutput{})
	pulumi.RegisterOutputType(IdentitySourceOutput{})
	pulumi.RegisterOutputType(IdentitySourceArrayOutput{})
	pulumi.RegisterOutputType(IdentitySourceResponseOutput{})
	pulumi.RegisterOutputType(IdentitySourceResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementClusterOutput{})
	pulumi.RegisterOutputType(ManagementClusterResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
}

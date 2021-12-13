


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterJobResponse struct {
	Id             string `pulumi:"id"`
	JobState       string `pulumi:"jobState"`
	StreamingUnits int    `pulumi:"streamingUnits"`
}

type ClusterPropertiesResponse struct {
	CapacityAllocated int    `pulumi:"capacityAllocated"`
	CapacityAssigned  int    `pulumi:"capacityAssigned"`
	ClusterId         string `pulumi:"clusterId"`
	CreatedDate       string `pulumi:"createdDate"`
	ProvisioningState string `pulumi:"provisioningState"`
}

type ClusterPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ClusterPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterPropertiesResponse)(nil)).Elem()
}

func (o ClusterPropertiesResponseOutput) ToClusterPropertiesResponseOutput() ClusterPropertiesResponseOutput {
	return o
}

func (o ClusterPropertiesResponseOutput) ToClusterPropertiesResponseOutputWithContext(ctx context.Context) ClusterPropertiesResponseOutput {
	return o
}

func (o ClusterPropertiesResponseOutput) CapacityAllocated() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterPropertiesResponse) int { return v.CapacityAllocated }).(pulumi.IntOutput)
}

func (o ClusterPropertiesResponseOutput) CapacityAssigned() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterPropertiesResponse) int { return v.CapacityAssigned }).(pulumi.IntOutput)
}

func (o ClusterPropertiesResponseOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterPropertiesResponse) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o ClusterPropertiesResponseOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterPropertiesResponse) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o ClusterPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ClusterSku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
}





type ClusterSkuInput interface {
	pulumi.Input

	ToClusterSkuOutput() ClusterSkuOutput
	ToClusterSkuOutputWithContext(context.Context) ClusterSkuOutput
}

type ClusterSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (ClusterSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSku)(nil)).Elem()
}

func (i ClusterSkuArgs) ToClusterSkuOutput() ClusterSkuOutput {
	return i.ToClusterSkuOutputWithContext(context.Background())
}

func (i ClusterSkuArgs) ToClusterSkuOutputWithContext(ctx context.Context) ClusterSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuOutput)
}

func (i ClusterSkuArgs) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return i.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (i ClusterSkuArgs) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuOutput).ToClusterSkuPtrOutputWithContext(ctx)
}









type ClusterSkuPtrInput interface {
	pulumi.Input

	ToClusterSkuPtrOutput() ClusterSkuPtrOutput
	ToClusterSkuPtrOutputWithContext(context.Context) ClusterSkuPtrOutput
}

type clusterSkuPtrType ClusterSkuArgs

func ClusterSkuPtr(v *ClusterSkuArgs) ClusterSkuPtrInput {
	return (*clusterSkuPtrType)(v)
}

func (*clusterSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSku)(nil)).Elem()
}

func (i *clusterSkuPtrType) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return i.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (i *clusterSkuPtrType) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuPtrOutput)
}

type ClusterSkuOutput struct{ *pulumi.OutputState }

func (ClusterSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSku)(nil)).Elem()
}

func (o ClusterSkuOutput) ToClusterSkuOutput() ClusterSkuOutput {
	return o
}

func (o ClusterSkuOutput) ToClusterSkuOutputWithContext(ctx context.Context) ClusterSkuOutput {
	return o
}

func (o ClusterSkuOutput) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return o.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (o ClusterSkuOutput) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterSku) *ClusterSku {
		return &v
	}).(ClusterSkuPtrOutput)
}

func (o ClusterSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ClusterSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ClusterSkuPtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSku)(nil)).Elem()
}

func (o ClusterSkuPtrOutput) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return o
}

func (o ClusterSkuPtrOutput) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return o
}

func (o ClusterSkuPtrOutput) Elem() ClusterSkuOutput {
	return o.ApplyT(func(v *ClusterSku) ClusterSku {
		if v != nil {
			return *v
		}
		var ret ClusterSku
		return ret
	}).(ClusterSkuOutput)
}

func (o ClusterSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ClusterSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ClusterSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
}

type ClusterSkuResponseOutput struct{ *pulumi.OutputState }

func (ClusterSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSkuResponse)(nil)).Elem()
}

func (o ClusterSkuResponseOutput) ToClusterSkuResponseOutput() ClusterSkuResponseOutput {
	return o
}

func (o ClusterSkuResponseOutput) ToClusterSkuResponseOutputWithContext(ctx context.Context) ClusterSkuResponseOutput {
	return o
}

func (o ClusterSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ClusterSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ClusterSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSkuResponse)(nil)).Elem()
}

func (o ClusterSkuResponsePtrOutput) ToClusterSkuResponsePtrOutput() ClusterSkuResponsePtrOutput {
	return o
}

func (o ClusterSkuResponsePtrOutput) ToClusterSkuResponsePtrOutputWithContext(ctx context.Context) ClusterSkuResponsePtrOutput {
	return o
}

func (o ClusterSkuResponsePtrOutput) Elem() ClusterSkuResponseOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) ClusterSkuResponse {
		if v != nil {
			return *v
		}
		var ret ClusterSkuResponse
		return ret
	}).(ClusterSkuResponseOutput)
}

func (o ClusterSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ClusterSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointProperties struct {
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnection `pulumi:"manualPrivateLinkServiceConnections"`
}





type PrivateEndpointPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointPropertiesOutput() PrivateEndpointPropertiesOutput
	ToPrivateEndpointPropertiesOutputWithContext(context.Context) PrivateEndpointPropertiesOutput
}

type PrivateEndpointPropertiesArgs struct {
	ManualPrivateLinkServiceConnections PrivateLinkServiceConnectionArrayInput `pulumi:"manualPrivateLinkServiceConnections"`
}

func (PrivateEndpointPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperties)(nil)).Elem()
}

func (i PrivateEndpointPropertiesArgs) ToPrivateEndpointPropertiesOutput() PrivateEndpointPropertiesOutput {
	return i.ToPrivateEndpointPropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertiesArgs) ToPrivateEndpointPropertiesOutputWithContext(ctx context.Context) PrivateEndpointPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertiesOutput)
}

func (i PrivateEndpointPropertiesArgs) ToPrivateEndpointPropertiesPtrOutput() PrivateEndpointPropertiesPtrOutput {
	return i.ToPrivateEndpointPropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertiesArgs) ToPrivateEndpointPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertiesOutput).ToPrivateEndpointPropertiesPtrOutputWithContext(ctx)
}









type PrivateEndpointPropertiesPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertiesPtrOutput() PrivateEndpointPropertiesPtrOutput
	ToPrivateEndpointPropertiesPtrOutputWithContext(context.Context) PrivateEndpointPropertiesPtrOutput
}

type privateEndpointPropertiesPtrType PrivateEndpointPropertiesArgs

func PrivateEndpointPropertiesPtr(v *PrivateEndpointPropertiesArgs) PrivateEndpointPropertiesPtrInput {
	return (*privateEndpointPropertiesPtrType)(v)
}

func (*privateEndpointPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperties)(nil)).Elem()
}

func (i *privateEndpointPropertiesPtrType) ToPrivateEndpointPropertiesPtrOutput() PrivateEndpointPropertiesPtrOutput {
	return i.ToPrivateEndpointPropertiesPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertiesPtrType) ToPrivateEndpointPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertiesPtrOutput)
}

type PrivateEndpointPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperties)(nil)).Elem()
}

func (o PrivateEndpointPropertiesOutput) ToPrivateEndpointPropertiesOutput() PrivateEndpointPropertiesOutput {
	return o
}

func (o PrivateEndpointPropertiesOutput) ToPrivateEndpointPropertiesOutputWithContext(ctx context.Context) PrivateEndpointPropertiesOutput {
	return o
}

func (o PrivateEndpointPropertiesOutput) ToPrivateEndpointPropertiesPtrOutput() PrivateEndpointPropertiesPtrOutput {
	return o.ToPrivateEndpointPropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertiesOutput) ToPrivateEndpointPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointProperties) *PrivateEndpointProperties {
		return &v
	}).(PrivateEndpointPropertiesPtrOutput)
}

func (o PrivateEndpointPropertiesOutput) ManualPrivateLinkServiceConnections() PrivateLinkServiceConnectionArrayOutput {
	return o.ApplyT(func(v PrivateEndpointProperties) []PrivateLinkServiceConnection {
		return v.ManualPrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionArrayOutput)
}

type PrivateEndpointPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperties)(nil)).Elem()
}

func (o PrivateEndpointPropertiesPtrOutput) ToPrivateEndpointPropertiesPtrOutput() PrivateEndpointPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointPropertiesPtrOutput) ToPrivateEndpointPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointPropertiesPtrOutput) Elem() PrivateEndpointPropertiesOutput {
	return o.ApplyT(func(v *PrivateEndpointProperties) PrivateEndpointProperties {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointProperties
		return ret
	}).(PrivateEndpointPropertiesOutput)
}

func (o PrivateEndpointPropertiesPtrOutput) ManualPrivateLinkServiceConnections() PrivateLinkServiceConnectionArrayOutput {
	return o.ApplyT(func(v *PrivateEndpointProperties) []PrivateLinkServiceConnection {
		if v == nil {
			return nil
		}
		return v.ManualPrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionArrayOutput)
}

type PrivateEndpointPropertiesResponse struct {
	CreatedDate                         string                                 `pulumi:"createdDate"`
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnectionResponse `pulumi:"manualPrivateLinkServiceConnections"`
}

type PrivateEndpointPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertiesResponseOutput) ToPrivateEndpointPropertiesResponseOutput() PrivateEndpointPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointPropertiesResponseOutput) ToPrivateEndpointPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointPropertiesResponseOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointPropertiesResponse) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o PrivateEndpointPropertiesResponseOutput) ManualPrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v PrivateEndpointPropertiesResponse) []PrivateLinkServiceConnectionResponse {
		return v.ManualPrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

type PrivateLinkConnectionStateResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}

type PrivateLinkConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkConnectionStateResponseOutput) ToPrivateLinkConnectionStateResponseOutput() PrivateLinkConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkConnectionStateResponseOutput) ToPrivateLinkConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkConnectionStateResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkConnectionStateResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkConnectionStateResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ToPrivateLinkConnectionStateResponsePtrOutput() PrivateLinkConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ToPrivateLinkConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Elem() PrivateLinkConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) PrivateLinkConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkConnectionStateResponse
		return ret
	}).(PrivateLinkConnectionStateResponseOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnection struct {
	GroupIds             []string `pulumi:"groupIds"`
	PrivateLinkServiceId *string  `pulumi:"privateLinkServiceId"`
	RequestMessage       *string  `pulumi:"requestMessage"`
}





type PrivateLinkServiceConnectionInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionOutput() PrivateLinkServiceConnectionOutput
	ToPrivateLinkServiceConnectionOutputWithContext(context.Context) PrivateLinkServiceConnectionOutput
}

type PrivateLinkServiceConnectionArgs struct {
	GroupIds             pulumi.StringArrayInput `pulumi:"groupIds"`
	PrivateLinkServiceId pulumi.StringPtrInput   `pulumi:"privateLinkServiceId"`
	RequestMessage       pulumi.StringPtrInput   `pulumi:"requestMessage"`
}

func (PrivateLinkServiceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnection)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionArgs) ToPrivateLinkServiceConnectionOutput() PrivateLinkServiceConnectionOutput {
	return i.ToPrivateLinkServiceConnectionOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionArgs) ToPrivateLinkServiceConnectionOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionOutput)
}





type PrivateLinkServiceConnectionArrayInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionArrayOutput() PrivateLinkServiceConnectionArrayOutput
	ToPrivateLinkServiceConnectionArrayOutputWithContext(context.Context) PrivateLinkServiceConnectionArrayOutput
}

type PrivateLinkServiceConnectionArray []PrivateLinkServiceConnectionInput

func (PrivateLinkServiceConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkServiceConnection)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionArray) ToPrivateLinkServiceConnectionArrayOutput() PrivateLinkServiceConnectionArrayOutput {
	return i.ToPrivateLinkServiceConnectionArrayOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionArray) ToPrivateLinkServiceConnectionArrayOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionArrayOutput)
}

type PrivateLinkServiceConnectionOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnection)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionOutput) ToPrivateLinkServiceConnectionOutput() PrivateLinkServiceConnectionOutput {
	return o
}

func (o PrivateLinkServiceConnectionOutput) ToPrivateLinkServiceConnectionOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionOutput {
	return o
}

func (o PrivateLinkServiceConnectionOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnection) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateLinkServiceConnectionOutput) PrivateLinkServiceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnection) *string { return v.PrivateLinkServiceId }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnection) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkServiceConnection)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionArrayOutput) ToPrivateLinkServiceConnectionArrayOutput() PrivateLinkServiceConnectionArrayOutput {
	return o
}

func (o PrivateLinkServiceConnectionArrayOutput) ToPrivateLinkServiceConnectionArrayOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionArrayOutput {
	return o
}

func (o PrivateLinkServiceConnectionArrayOutput) Index(i pulumi.IntInput) PrivateLinkServiceConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkServiceConnection {
		return vs[0].([]PrivateLinkServiceConnection)[vs[1].(int)]
	}).(PrivateLinkServiceConnectionOutput)
}

type PrivateLinkServiceConnectionResponse struct {
	GroupIds                          []string                            `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	PrivateLinkServiceId              *string                             `pulumi:"privateLinkServiceId"`
	RequestMessage                    *string                             `pulumi:"requestMessage"`
}

type PrivateLinkServiceConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionResponseOutput) ToPrivateLinkServiceConnectionResponseOutput() PrivateLinkServiceConnectionResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionResponseOutput) ToPrivateLinkServiceConnectionResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateLinkServiceConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionResponse) *PrivateLinkConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionResponseOutput) PrivateLinkServiceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionResponse) *string { return v.PrivateLinkServiceId }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionResponseOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionResponse) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkServiceConnectionResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionResponseArrayOutput) ToPrivateLinkServiceConnectionResponseArrayOutput() PrivateLinkServiceConnectionResponseArrayOutput {
	return o
}

func (o PrivateLinkServiceConnectionResponseArrayOutput) ToPrivateLinkServiceConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionResponseArrayOutput {
	return o
}

func (o PrivateLinkServiceConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateLinkServiceConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkServiceConnectionResponse {
		return vs[0].([]PrivateLinkServiceConnectionResponse)[vs[1].(int)]
	}).(PrivateLinkServiceConnectionResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ClusterSkuOutput{})
	pulumi.RegisterOutputType(ClusterSkuPtrOutput{})
	pulumi.RegisterOutputType(ClusterSkuResponseOutput{})
	pulumi.RegisterOutputType(ClusterSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionResponseArrayOutput{})
}

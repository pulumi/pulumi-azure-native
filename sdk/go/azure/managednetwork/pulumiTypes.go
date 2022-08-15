


package managednetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectivityCollectionResponse struct {
	Groups   []ManagedNetworkGroupResponse         `pulumi:"groups"`
	Peerings []ManagedNetworkPeeringPolicyResponse `pulumi:"peerings"`
}

type ConnectivityCollectionResponseOutput struct{ *pulumi.OutputState }

func (ConnectivityCollectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityCollectionResponse)(nil)).Elem()
}

func (o ConnectivityCollectionResponseOutput) ToConnectivityCollectionResponseOutput() ConnectivityCollectionResponseOutput {
	return o
}

func (o ConnectivityCollectionResponseOutput) ToConnectivityCollectionResponseOutputWithContext(ctx context.Context) ConnectivityCollectionResponseOutput {
	return o
}

func (o ConnectivityCollectionResponseOutput) Groups() ManagedNetworkGroupResponseArrayOutput {
	return o.ApplyT(func(v ConnectivityCollectionResponse) []ManagedNetworkGroupResponse { return v.Groups }).(ManagedNetworkGroupResponseArrayOutput)
}

func (o ConnectivityCollectionResponseOutput) Peerings() ManagedNetworkPeeringPolicyResponseArrayOutput {
	return o.ApplyT(func(v ConnectivityCollectionResponse) []ManagedNetworkPeeringPolicyResponse { return v.Peerings }).(ManagedNetworkPeeringPolicyResponseArrayOutput)
}

type ManagedNetworkGroupResponse struct {
	Etag              string               `pulumi:"etag"`
	Id                string               `pulumi:"id"`
	Kind              *string              `pulumi:"kind"`
	Location          *string              `pulumi:"location"`
	ManagementGroups  []ResourceIdResponse `pulumi:"managementGroups"`
	Name              string               `pulumi:"name"`
	ProvisioningState string               `pulumi:"provisioningState"`
	Subnets           []ResourceIdResponse `pulumi:"subnets"`
	Subscriptions     []ResourceIdResponse `pulumi:"subscriptions"`
	Type              string               `pulumi:"type"`
	VirtualNetworks   []ResourceIdResponse `pulumi:"virtualNetworks"`
}

type ManagedNetworkGroupResponseOutput struct{ *pulumi.OutputState }

func (ManagedNetworkGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedNetworkGroupResponse)(nil)).Elem()
}

func (o ManagedNetworkGroupResponseOutput) ToManagedNetworkGroupResponseOutput() ManagedNetworkGroupResponseOutput {
	return o
}

func (o ManagedNetworkGroupResponseOutput) ToManagedNetworkGroupResponseOutputWithContext(ctx context.Context) ManagedNetworkGroupResponseOutput {
	return o
}

func (o ManagedNetworkGroupResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedNetworkGroupResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ManagedNetworkGroupResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedNetworkGroupResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ManagedNetworkGroupResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedNetworkGroupResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ManagedNetworkGroupResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedNetworkGroupResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ManagedNetworkGroupResponseOutput) ManagementGroups() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v ManagedNetworkGroupResponse) []ResourceIdResponse { return v.ManagementGroups }).(ResourceIdResponseArrayOutput)
}

func (o ManagedNetworkGroupResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedNetworkGroupResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedNetworkGroupResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedNetworkGroupResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedNetworkGroupResponseOutput) Subnets() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v ManagedNetworkGroupResponse) []ResourceIdResponse { return v.Subnets }).(ResourceIdResponseArrayOutput)
}

func (o ManagedNetworkGroupResponseOutput) Subscriptions() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v ManagedNetworkGroupResponse) []ResourceIdResponse { return v.Subscriptions }).(ResourceIdResponseArrayOutput)
}

func (o ManagedNetworkGroupResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedNetworkGroupResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedNetworkGroupResponseOutput) VirtualNetworks() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v ManagedNetworkGroupResponse) []ResourceIdResponse { return v.VirtualNetworks }).(ResourceIdResponseArrayOutput)
}

type ManagedNetworkGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedNetworkGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedNetworkGroupResponse)(nil)).Elem()
}

func (o ManagedNetworkGroupResponseArrayOutput) ToManagedNetworkGroupResponseArrayOutput() ManagedNetworkGroupResponseArrayOutput {
	return o
}

func (o ManagedNetworkGroupResponseArrayOutput) ToManagedNetworkGroupResponseArrayOutputWithContext(ctx context.Context) ManagedNetworkGroupResponseArrayOutput {
	return o
}

func (o ManagedNetworkGroupResponseArrayOutput) Index(i pulumi.IntInput) ManagedNetworkGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedNetworkGroupResponse {
		return vs[0].([]ManagedNetworkGroupResponse)[vs[1].(int)]
	}).(ManagedNetworkGroupResponseOutput)
}

type ManagedNetworkPeeringPolicyProperties struct {
	Hub    *ResourceId  `pulumi:"hub"`
	Mesh   []ResourceId `pulumi:"mesh"`
	Spokes []ResourceId `pulumi:"spokes"`
	Type   string       `pulumi:"type"`
}





type ManagedNetworkPeeringPolicyPropertiesInput interface {
	pulumi.Input

	ToManagedNetworkPeeringPolicyPropertiesOutput() ManagedNetworkPeeringPolicyPropertiesOutput
	ToManagedNetworkPeeringPolicyPropertiesOutputWithContext(context.Context) ManagedNetworkPeeringPolicyPropertiesOutput
}

type ManagedNetworkPeeringPolicyPropertiesArgs struct {
	Hub    ResourceIdPtrInput   `pulumi:"hub"`
	Mesh   ResourceIdArrayInput `pulumi:"mesh"`
	Spokes ResourceIdArrayInput `pulumi:"spokes"`
	Type   pulumi.StringInput   `pulumi:"type"`
}

func (ManagedNetworkPeeringPolicyPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedNetworkPeeringPolicyProperties)(nil)).Elem()
}

func (i ManagedNetworkPeeringPolicyPropertiesArgs) ToManagedNetworkPeeringPolicyPropertiesOutput() ManagedNetworkPeeringPolicyPropertiesOutput {
	return i.ToManagedNetworkPeeringPolicyPropertiesOutputWithContext(context.Background())
}

func (i ManagedNetworkPeeringPolicyPropertiesArgs) ToManagedNetworkPeeringPolicyPropertiesOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedNetworkPeeringPolicyPropertiesOutput)
}

func (i ManagedNetworkPeeringPolicyPropertiesArgs) ToManagedNetworkPeeringPolicyPropertiesPtrOutput() ManagedNetworkPeeringPolicyPropertiesPtrOutput {
	return i.ToManagedNetworkPeeringPolicyPropertiesPtrOutputWithContext(context.Background())
}

func (i ManagedNetworkPeeringPolicyPropertiesArgs) ToManagedNetworkPeeringPolicyPropertiesPtrOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedNetworkPeeringPolicyPropertiesOutput).ToManagedNetworkPeeringPolicyPropertiesPtrOutputWithContext(ctx)
}









type ManagedNetworkPeeringPolicyPropertiesPtrInput interface {
	pulumi.Input

	ToManagedNetworkPeeringPolicyPropertiesPtrOutput() ManagedNetworkPeeringPolicyPropertiesPtrOutput
	ToManagedNetworkPeeringPolicyPropertiesPtrOutputWithContext(context.Context) ManagedNetworkPeeringPolicyPropertiesPtrOutput
}

type managedNetworkPeeringPolicyPropertiesPtrType ManagedNetworkPeeringPolicyPropertiesArgs

func ManagedNetworkPeeringPolicyPropertiesPtr(v *ManagedNetworkPeeringPolicyPropertiesArgs) ManagedNetworkPeeringPolicyPropertiesPtrInput {
	return (*managedNetworkPeeringPolicyPropertiesPtrType)(v)
}

func (*managedNetworkPeeringPolicyPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetworkPeeringPolicyProperties)(nil)).Elem()
}

func (i *managedNetworkPeeringPolicyPropertiesPtrType) ToManagedNetworkPeeringPolicyPropertiesPtrOutput() ManagedNetworkPeeringPolicyPropertiesPtrOutput {
	return i.ToManagedNetworkPeeringPolicyPropertiesPtrOutputWithContext(context.Background())
}

func (i *managedNetworkPeeringPolicyPropertiesPtrType) ToManagedNetworkPeeringPolicyPropertiesPtrOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedNetworkPeeringPolicyPropertiesPtrOutput)
}

type ManagedNetworkPeeringPolicyPropertiesOutput struct{ *pulumi.OutputState }

func (ManagedNetworkPeeringPolicyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedNetworkPeeringPolicyProperties)(nil)).Elem()
}

func (o ManagedNetworkPeeringPolicyPropertiesOutput) ToManagedNetworkPeeringPolicyPropertiesOutput() ManagedNetworkPeeringPolicyPropertiesOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyPropertiesOutput) ToManagedNetworkPeeringPolicyPropertiesOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyPropertiesOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyPropertiesOutput) ToManagedNetworkPeeringPolicyPropertiesPtrOutput() ManagedNetworkPeeringPolicyPropertiesPtrOutput {
	return o.ToManagedNetworkPeeringPolicyPropertiesPtrOutputWithContext(context.Background())
}

func (o ManagedNetworkPeeringPolicyPropertiesOutput) ToManagedNetworkPeeringPolicyPropertiesPtrOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedNetworkPeeringPolicyProperties) *ManagedNetworkPeeringPolicyProperties {
		return &v
	}).(ManagedNetworkPeeringPolicyPropertiesPtrOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesOutput) Hub() ResourceIdPtrOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyProperties) *ResourceId { return v.Hub }).(ResourceIdPtrOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesOutput) Mesh() ResourceIdArrayOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyProperties) []ResourceId { return v.Mesh }).(ResourceIdArrayOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesOutput) Spokes() ResourceIdArrayOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyProperties) []ResourceId { return v.Spokes }).(ResourceIdArrayOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyProperties) string { return v.Type }).(pulumi.StringOutput)
}

type ManagedNetworkPeeringPolicyPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ManagedNetworkPeeringPolicyPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetworkPeeringPolicyProperties)(nil)).Elem()
}

func (o ManagedNetworkPeeringPolicyPropertiesPtrOutput) ToManagedNetworkPeeringPolicyPropertiesPtrOutput() ManagedNetworkPeeringPolicyPropertiesPtrOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyPropertiesPtrOutput) ToManagedNetworkPeeringPolicyPropertiesPtrOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyPropertiesPtrOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyPropertiesPtrOutput) Elem() ManagedNetworkPeeringPolicyPropertiesOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicyProperties) ManagedNetworkPeeringPolicyProperties {
		if v != nil {
			return *v
		}
		var ret ManagedNetworkPeeringPolicyProperties
		return ret
	}).(ManagedNetworkPeeringPolicyPropertiesOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesPtrOutput) Hub() ResourceIdPtrOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicyProperties) *ResourceId {
		if v == nil {
			return nil
		}
		return v.Hub
	}).(ResourceIdPtrOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesPtrOutput) Mesh() ResourceIdArrayOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicyProperties) []ResourceId {
		if v == nil {
			return nil
		}
		return v.Mesh
	}).(ResourceIdArrayOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesPtrOutput) Spokes() ResourceIdArrayOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicyProperties) []ResourceId {
		if v == nil {
			return nil
		}
		return v.Spokes
	}).(ResourceIdArrayOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicyProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ManagedNetworkPeeringPolicyPropertiesResponse struct {
	Etag              string               `pulumi:"etag"`
	Hub               *ResourceIdResponse  `pulumi:"hub"`
	Mesh              []ResourceIdResponse `pulumi:"mesh"`
	ProvisioningState string               `pulumi:"provisioningState"`
	Spokes            []ResourceIdResponse `pulumi:"spokes"`
	Type              string               `pulumi:"type"`
}

type ManagedNetworkPeeringPolicyPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManagedNetworkPeeringPolicyPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedNetworkPeeringPolicyPropertiesResponse)(nil)).Elem()
}

func (o ManagedNetworkPeeringPolicyPropertiesResponseOutput) ToManagedNetworkPeeringPolicyPropertiesResponseOutput() ManagedNetworkPeeringPolicyPropertiesResponseOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyPropertiesResponseOutput) ToManagedNetworkPeeringPolicyPropertiesResponseOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyPropertiesResponseOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyPropertiesResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyPropertiesResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesResponseOutput) Hub() ResourceIdResponsePtrOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyPropertiesResponse) *ResourceIdResponse { return v.Hub }).(ResourceIdResponsePtrOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesResponseOutput) Mesh() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyPropertiesResponse) []ResourceIdResponse { return v.Mesh }).(ResourceIdResponseArrayOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesResponseOutput) Spokes() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyPropertiesResponse) []ResourceIdResponse { return v.Spokes }).(ResourceIdResponseArrayOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetworkPeeringPolicyPropertiesResponse)(nil)).Elem()
}

func (o ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput) ToManagedNetworkPeeringPolicyPropertiesResponsePtrOutput() ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput) ToManagedNetworkPeeringPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput) Elem() ManagedNetworkPeeringPolicyPropertiesResponseOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicyPropertiesResponse) ManagedNetworkPeeringPolicyPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ManagedNetworkPeeringPolicyPropertiesResponse
		return ret
	}).(ManagedNetworkPeeringPolicyPropertiesResponseOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicyPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput) Hub() ResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicyPropertiesResponse) *ResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.Hub
	}).(ResourceIdResponsePtrOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput) Mesh() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicyPropertiesResponse) []ResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.Mesh
	}).(ResourceIdResponseArrayOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicyPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput) Spokes() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicyPropertiesResponse) []ResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.Spokes
	}).(ResourceIdResponseArrayOutput)
}

func (o ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicyPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ManagedNetworkPeeringPolicyResponse struct {
	Id         string                                         `pulumi:"id"`
	Location   *string                                        `pulumi:"location"`
	Name       string                                         `pulumi:"name"`
	Properties *ManagedNetworkPeeringPolicyPropertiesResponse `pulumi:"properties"`
	Type       string                                         `pulumi:"type"`
}

type ManagedNetworkPeeringPolicyResponseOutput struct{ *pulumi.OutputState }

func (ManagedNetworkPeeringPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedNetworkPeeringPolicyResponse)(nil)).Elem()
}

func (o ManagedNetworkPeeringPolicyResponseOutput) ToManagedNetworkPeeringPolicyResponseOutput() ManagedNetworkPeeringPolicyResponseOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyResponseOutput) ToManagedNetworkPeeringPolicyResponseOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyResponseOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ManagedNetworkPeeringPolicyResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ManagedNetworkPeeringPolicyResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedNetworkPeeringPolicyResponseOutput) Properties() ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyResponse) *ManagedNetworkPeeringPolicyPropertiesResponse {
		return v.Properties
	}).(ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput)
}

func (o ManagedNetworkPeeringPolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedNetworkPeeringPolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ManagedNetworkPeeringPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedNetworkPeeringPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedNetworkPeeringPolicyResponse)(nil)).Elem()
}

func (o ManagedNetworkPeeringPolicyResponseArrayOutput) ToManagedNetworkPeeringPolicyResponseArrayOutput() ManagedNetworkPeeringPolicyResponseArrayOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyResponseArrayOutput) ToManagedNetworkPeeringPolicyResponseArrayOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyResponseArrayOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyResponseArrayOutput) Index(i pulumi.IntInput) ManagedNetworkPeeringPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedNetworkPeeringPolicyResponse {
		return vs[0].([]ManagedNetworkPeeringPolicyResponse)[vs[1].(int)]
	}).(ManagedNetworkPeeringPolicyResponseOutput)
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





type ResourceIdArrayInput interface {
	pulumi.Input

	ToResourceIdArrayOutput() ResourceIdArrayOutput
	ToResourceIdArrayOutputWithContext(context.Context) ResourceIdArrayOutput
}

type ResourceIdArray []ResourceIdInput

func (ResourceIdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceId)(nil)).Elem()
}

func (i ResourceIdArray) ToResourceIdArrayOutput() ResourceIdArrayOutput {
	return i.ToResourceIdArrayOutputWithContext(context.Background())
}

func (i ResourceIdArray) ToResourceIdArrayOutputWithContext(ctx context.Context) ResourceIdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdArrayOutput)
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

type ResourceIdArrayOutput struct{ *pulumi.OutputState }

func (ResourceIdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceId)(nil)).Elem()
}

func (o ResourceIdArrayOutput) ToResourceIdArrayOutput() ResourceIdArrayOutput {
	return o
}

func (o ResourceIdArrayOutput) ToResourceIdArrayOutputWithContext(ctx context.Context) ResourceIdArrayOutput {
	return o
}

func (o ResourceIdArrayOutput) Index(i pulumi.IntInput) ResourceIdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceId {
		return vs[0].([]ResourceId)[vs[1].(int)]
	}).(ResourceIdOutput)
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

type ResourceIdResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceIdResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceIdResponse)(nil)).Elem()
}

func (o ResourceIdResponseArrayOutput) ToResourceIdResponseArrayOutput() ResourceIdResponseArrayOutput {
	return o
}

func (o ResourceIdResponseArrayOutput) ToResourceIdResponseArrayOutputWithContext(ctx context.Context) ResourceIdResponseArrayOutput {
	return o
}

func (o ResourceIdResponseArrayOutput) Index(i pulumi.IntInput) ResourceIdResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceIdResponse {
		return vs[0].([]ResourceIdResponse)[vs[1].(int)]
	}).(ResourceIdResponseOutput)
}

type Scope struct {
	ManagementGroups []ResourceId `pulumi:"managementGroups"`
	Subnets          []ResourceId `pulumi:"subnets"`
	Subscriptions    []ResourceId `pulumi:"subscriptions"`
	VirtualNetworks  []ResourceId `pulumi:"virtualNetworks"`
}





type ScopeInput interface {
	pulumi.Input

	ToScopeOutput() ScopeOutput
	ToScopeOutputWithContext(context.Context) ScopeOutput
}

type ScopeArgs struct {
	ManagementGroups ResourceIdArrayInput `pulumi:"managementGroups"`
	Subnets          ResourceIdArrayInput `pulumi:"subnets"`
	Subscriptions    ResourceIdArrayInput `pulumi:"subscriptions"`
	VirtualNetworks  ResourceIdArrayInput `pulumi:"virtualNetworks"`
}

func (ScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Scope)(nil)).Elem()
}

func (i ScopeArgs) ToScopeOutput() ScopeOutput {
	return i.ToScopeOutputWithContext(context.Background())
}

func (i ScopeArgs) ToScopeOutputWithContext(ctx context.Context) ScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeOutput)
}

func (i ScopeArgs) ToScopePtrOutput() ScopePtrOutput {
	return i.ToScopePtrOutputWithContext(context.Background())
}

func (i ScopeArgs) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeOutput).ToScopePtrOutputWithContext(ctx)
}









type ScopePtrInput interface {
	pulumi.Input

	ToScopePtrOutput() ScopePtrOutput
	ToScopePtrOutputWithContext(context.Context) ScopePtrOutput
}

type scopePtrType ScopeArgs

func ScopePtr(v *ScopeArgs) ScopePtrInput {
	return (*scopePtrType)(v)
}

func (*scopePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Scope)(nil)).Elem()
}

func (i *scopePtrType) ToScopePtrOutput() ScopePtrOutput {
	return i.ToScopePtrOutputWithContext(context.Background())
}

func (i *scopePtrType) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopePtrOutput)
}

type ScopeOutput struct{ *pulumi.OutputState }

func (ScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Scope)(nil)).Elem()
}

func (o ScopeOutput) ToScopeOutput() ScopeOutput {
	return o
}

func (o ScopeOutput) ToScopeOutputWithContext(ctx context.Context) ScopeOutput {
	return o
}

func (o ScopeOutput) ToScopePtrOutput() ScopePtrOutput {
	return o.ToScopePtrOutputWithContext(context.Background())
}

func (o ScopeOutput) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Scope) *Scope {
		return &v
	}).(ScopePtrOutput)
}

func (o ScopeOutput) ManagementGroups() ResourceIdArrayOutput {
	return o.ApplyT(func(v Scope) []ResourceId { return v.ManagementGroups }).(ResourceIdArrayOutput)
}

func (o ScopeOutput) Subnets() ResourceIdArrayOutput {
	return o.ApplyT(func(v Scope) []ResourceId { return v.Subnets }).(ResourceIdArrayOutput)
}

func (o ScopeOutput) Subscriptions() ResourceIdArrayOutput {
	return o.ApplyT(func(v Scope) []ResourceId { return v.Subscriptions }).(ResourceIdArrayOutput)
}

func (o ScopeOutput) VirtualNetworks() ResourceIdArrayOutput {
	return o.ApplyT(func(v Scope) []ResourceId { return v.VirtualNetworks }).(ResourceIdArrayOutput)
}

type ScopePtrOutput struct{ *pulumi.OutputState }

func (ScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scope)(nil)).Elem()
}

func (o ScopePtrOutput) ToScopePtrOutput() ScopePtrOutput {
	return o
}

func (o ScopePtrOutput) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return o
}

func (o ScopePtrOutput) Elem() ScopeOutput {
	return o.ApplyT(func(v *Scope) Scope {
		if v != nil {
			return *v
		}
		var ret Scope
		return ret
	}).(ScopeOutput)
}

func (o ScopePtrOutput) ManagementGroups() ResourceIdArrayOutput {
	return o.ApplyT(func(v *Scope) []ResourceId {
		if v == nil {
			return nil
		}
		return v.ManagementGroups
	}).(ResourceIdArrayOutput)
}

func (o ScopePtrOutput) Subnets() ResourceIdArrayOutput {
	return o.ApplyT(func(v *Scope) []ResourceId {
		if v == nil {
			return nil
		}
		return v.Subnets
	}).(ResourceIdArrayOutput)
}

func (o ScopePtrOutput) Subscriptions() ResourceIdArrayOutput {
	return o.ApplyT(func(v *Scope) []ResourceId {
		if v == nil {
			return nil
		}
		return v.Subscriptions
	}).(ResourceIdArrayOutput)
}

func (o ScopePtrOutput) VirtualNetworks() ResourceIdArrayOutput {
	return o.ApplyT(func(v *Scope) []ResourceId {
		if v == nil {
			return nil
		}
		return v.VirtualNetworks
	}).(ResourceIdArrayOutput)
}

type ScopeResponse struct {
	ManagementGroups []ResourceIdResponse `pulumi:"managementGroups"`
	Subnets          []ResourceIdResponse `pulumi:"subnets"`
	Subscriptions    []ResourceIdResponse `pulumi:"subscriptions"`
	VirtualNetworks  []ResourceIdResponse `pulumi:"virtualNetworks"`
}

type ScopeResponseOutput struct{ *pulumi.OutputState }

func (ScopeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeResponse)(nil)).Elem()
}

func (o ScopeResponseOutput) ToScopeResponseOutput() ScopeResponseOutput {
	return o
}

func (o ScopeResponseOutput) ToScopeResponseOutputWithContext(ctx context.Context) ScopeResponseOutput {
	return o
}

func (o ScopeResponseOutput) ManagementGroups() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v ScopeResponse) []ResourceIdResponse { return v.ManagementGroups }).(ResourceIdResponseArrayOutput)
}

func (o ScopeResponseOutput) Subnets() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v ScopeResponse) []ResourceIdResponse { return v.Subnets }).(ResourceIdResponseArrayOutput)
}

func (o ScopeResponseOutput) Subscriptions() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v ScopeResponse) []ResourceIdResponse { return v.Subscriptions }).(ResourceIdResponseArrayOutput)
}

func (o ScopeResponseOutput) VirtualNetworks() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v ScopeResponse) []ResourceIdResponse { return v.VirtualNetworks }).(ResourceIdResponseArrayOutput)
}

type ScopeResponsePtrOutput struct{ *pulumi.OutputState }

func (ScopeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeResponse)(nil)).Elem()
}

func (o ScopeResponsePtrOutput) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return o
}

func (o ScopeResponsePtrOutput) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return o
}

func (o ScopeResponsePtrOutput) Elem() ScopeResponseOutput {
	return o.ApplyT(func(v *ScopeResponse) ScopeResponse {
		if v != nil {
			return *v
		}
		var ret ScopeResponse
		return ret
	}).(ScopeResponseOutput)
}

func (o ScopeResponsePtrOutput) ManagementGroups() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v *ScopeResponse) []ResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.ManagementGroups
	}).(ResourceIdResponseArrayOutput)
}

func (o ScopeResponsePtrOutput) Subnets() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v *ScopeResponse) []ResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.Subnets
	}).(ResourceIdResponseArrayOutput)
}

func (o ScopeResponsePtrOutput) Subscriptions() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v *ScopeResponse) []ResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.Subscriptions
	}).(ResourceIdResponseArrayOutput)
}

func (o ScopeResponsePtrOutput) VirtualNetworks() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v *ScopeResponse) []ResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.VirtualNetworks
	}).(ResourceIdResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectivityCollectionResponseOutput{})
	pulumi.RegisterOutputType(ManagedNetworkGroupResponseOutput{})
	pulumi.RegisterOutputType(ManagedNetworkGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedNetworkPeeringPolicyPropertiesOutput{})
	pulumi.RegisterOutputType(ManagedNetworkPeeringPolicyPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagedNetworkPeeringPolicyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManagedNetworkPeeringPolicyPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedNetworkPeeringPolicyResponseOutput{})
	pulumi.RegisterOutputType(ManagedNetworkPeeringPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceIdOutput{})
	pulumi.RegisterOutputType(ResourceIdPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdArrayOutput{})
	pulumi.RegisterOutputType(ResourceIdResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdResponseArrayOutput{})
	pulumi.RegisterOutputType(ScopeOutput{})
	pulumi.RegisterOutputType(ScopePtrOutput{})
	pulumi.RegisterOutputType(ScopeResponseOutput{})
	pulumi.RegisterOutputType(ScopeResponsePtrOutput{})
}

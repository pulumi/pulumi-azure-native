


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectionStateRequestBodyProperties struct {
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type ConnectionStateRequestBodyPropertiesInput interface {
	pulumi.Input

	ToConnectionStateRequestBodyPropertiesOutput() ConnectionStateRequestBodyPropertiesOutput
	ToConnectionStateRequestBodyPropertiesOutputWithContext(context.Context) ConnectionStateRequestBodyPropertiesOutput
}

type ConnectionStateRequestBodyPropertiesArgs struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (ConnectionStateRequestBodyPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStateRequestBodyProperties)(nil)).Elem()
}

func (i ConnectionStateRequestBodyPropertiesArgs) ToConnectionStateRequestBodyPropertiesOutput() ConnectionStateRequestBodyPropertiesOutput {
	return i.ToConnectionStateRequestBodyPropertiesOutputWithContext(context.Background())
}

func (i ConnectionStateRequestBodyPropertiesArgs) ToConnectionStateRequestBodyPropertiesOutputWithContext(ctx context.Context) ConnectionStateRequestBodyPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateRequestBodyPropertiesOutput)
}

func (i ConnectionStateRequestBodyPropertiesArgs) ToConnectionStateRequestBodyPropertiesPtrOutput() ConnectionStateRequestBodyPropertiesPtrOutput {
	return i.ToConnectionStateRequestBodyPropertiesPtrOutputWithContext(context.Background())
}

func (i ConnectionStateRequestBodyPropertiesArgs) ToConnectionStateRequestBodyPropertiesPtrOutputWithContext(ctx context.Context) ConnectionStateRequestBodyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateRequestBodyPropertiesOutput).ToConnectionStateRequestBodyPropertiesPtrOutputWithContext(ctx)
}









type ConnectionStateRequestBodyPropertiesPtrInput interface {
	pulumi.Input

	ToConnectionStateRequestBodyPropertiesPtrOutput() ConnectionStateRequestBodyPropertiesPtrOutput
	ToConnectionStateRequestBodyPropertiesPtrOutputWithContext(context.Context) ConnectionStateRequestBodyPropertiesPtrOutput
}

type connectionStateRequestBodyPropertiesPtrType ConnectionStateRequestBodyPropertiesArgs

func ConnectionStateRequestBodyPropertiesPtr(v *ConnectionStateRequestBodyPropertiesArgs) ConnectionStateRequestBodyPropertiesPtrInput {
	return (*connectionStateRequestBodyPropertiesPtrType)(v)
}

func (*connectionStateRequestBodyPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStateRequestBodyProperties)(nil)).Elem()
}

func (i *connectionStateRequestBodyPropertiesPtrType) ToConnectionStateRequestBodyPropertiesPtrOutput() ConnectionStateRequestBodyPropertiesPtrOutput {
	return i.ToConnectionStateRequestBodyPropertiesPtrOutputWithContext(context.Background())
}

func (i *connectionStateRequestBodyPropertiesPtrType) ToConnectionStateRequestBodyPropertiesPtrOutputWithContext(ctx context.Context) ConnectionStateRequestBodyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateRequestBodyPropertiesPtrOutput)
}

type ConnectionStateRequestBodyPropertiesOutput struct{ *pulumi.OutputState }

func (ConnectionStateRequestBodyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStateRequestBodyProperties)(nil)).Elem()
}

func (o ConnectionStateRequestBodyPropertiesOutput) ToConnectionStateRequestBodyPropertiesOutput() ConnectionStateRequestBodyPropertiesOutput {
	return o
}

func (o ConnectionStateRequestBodyPropertiesOutput) ToConnectionStateRequestBodyPropertiesOutputWithContext(ctx context.Context) ConnectionStateRequestBodyPropertiesOutput {
	return o
}

func (o ConnectionStateRequestBodyPropertiesOutput) ToConnectionStateRequestBodyPropertiesPtrOutput() ConnectionStateRequestBodyPropertiesPtrOutput {
	return o.ToConnectionStateRequestBodyPropertiesPtrOutputWithContext(context.Background())
}

func (o ConnectionStateRequestBodyPropertiesOutput) ToConnectionStateRequestBodyPropertiesPtrOutputWithContext(ctx context.Context) ConnectionStateRequestBodyPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionStateRequestBodyProperties) *ConnectionStateRequestBodyProperties {
		return &v
	}).(ConnectionStateRequestBodyPropertiesPtrOutput)
}

func (o ConnectionStateRequestBodyPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v ConnectionStateRequestBodyProperties) *PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type ConnectionStateRequestBodyPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConnectionStateRequestBodyPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStateRequestBodyProperties)(nil)).Elem()
}

func (o ConnectionStateRequestBodyPropertiesPtrOutput) ToConnectionStateRequestBodyPropertiesPtrOutput() ConnectionStateRequestBodyPropertiesPtrOutput {
	return o
}

func (o ConnectionStateRequestBodyPropertiesPtrOutput) ToConnectionStateRequestBodyPropertiesPtrOutputWithContext(ctx context.Context) ConnectionStateRequestBodyPropertiesPtrOutput {
	return o
}

func (o ConnectionStateRequestBodyPropertiesPtrOutput) Elem() ConnectionStateRequestBodyPropertiesOutput {
	return o.ApplyT(func(v *ConnectionStateRequestBodyProperties) ConnectionStateRequestBodyProperties {
		if v != nil {
			return *v
		}
		var ret ConnectionStateRequestBodyProperties
		return ret
	}).(ConnectionStateRequestBodyPropertiesOutput)
}

func (o ConnectionStateRequestBodyPropertiesPtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *ConnectionStateRequestBodyProperties) *PrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type MigrateProjectProperties struct {
	PublicNetworkAccess     *string `pulumi:"publicNetworkAccess"`
	ServiceEndpoint         *string `pulumi:"serviceEndpoint"`
	UtilityStorageAccountId *string `pulumi:"utilityStorageAccountId"`
}





type MigrateProjectPropertiesInput interface {
	pulumi.Input

	ToMigrateProjectPropertiesOutput() MigrateProjectPropertiesOutput
	ToMigrateProjectPropertiesOutputWithContext(context.Context) MigrateProjectPropertiesOutput
}

type MigrateProjectPropertiesArgs struct {
	PublicNetworkAccess     pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
	ServiceEndpoint         pulumi.StringPtrInput `pulumi:"serviceEndpoint"`
	UtilityStorageAccountId pulumi.StringPtrInput `pulumi:"utilityStorageAccountId"`
}

func (MigrateProjectPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectProperties)(nil)).Elem()
}

func (i MigrateProjectPropertiesArgs) ToMigrateProjectPropertiesOutput() MigrateProjectPropertiesOutput {
	return i.ToMigrateProjectPropertiesOutputWithContext(context.Background())
}

func (i MigrateProjectPropertiesArgs) ToMigrateProjectPropertiesOutputWithContext(ctx context.Context) MigrateProjectPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectPropertiesOutput)
}

func (i MigrateProjectPropertiesArgs) ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput {
	return i.ToMigrateProjectPropertiesPtrOutputWithContext(context.Background())
}

func (i MigrateProjectPropertiesArgs) ToMigrateProjectPropertiesPtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectPropertiesOutput).ToMigrateProjectPropertiesPtrOutputWithContext(ctx)
}









type MigrateProjectPropertiesPtrInput interface {
	pulumi.Input

	ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput
	ToMigrateProjectPropertiesPtrOutputWithContext(context.Context) MigrateProjectPropertiesPtrOutput
}

type migrateProjectPropertiesPtrType MigrateProjectPropertiesArgs

func MigrateProjectPropertiesPtr(v *MigrateProjectPropertiesArgs) MigrateProjectPropertiesPtrInput {
	return (*migrateProjectPropertiesPtrType)(v)
}

func (*migrateProjectPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectProperties)(nil)).Elem()
}

func (i *migrateProjectPropertiesPtrType) ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput {
	return i.ToMigrateProjectPropertiesPtrOutputWithContext(context.Background())
}

func (i *migrateProjectPropertiesPtrType) ToMigrateProjectPropertiesPtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectPropertiesPtrOutput)
}

type MigrateProjectPropertiesOutput struct{ *pulumi.OutputState }

func (MigrateProjectPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectProperties)(nil)).Elem()
}

func (o MigrateProjectPropertiesOutput) ToMigrateProjectPropertiesOutput() MigrateProjectPropertiesOutput {
	return o
}

func (o MigrateProjectPropertiesOutput) ToMigrateProjectPropertiesOutputWithContext(ctx context.Context) MigrateProjectPropertiesOutput {
	return o
}

func (o MigrateProjectPropertiesOutput) ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput {
	return o.ToMigrateProjectPropertiesPtrOutputWithContext(context.Background())
}

func (o MigrateProjectPropertiesOutput) ToMigrateProjectPropertiesPtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrateProjectProperties) *MigrateProjectProperties {
		return &v
	}).(MigrateProjectPropertiesPtrOutput)
}

func (o MigrateProjectPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesOutput) ServiceEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectProperties) *string { return v.ServiceEndpoint }).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesOutput) UtilityStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectProperties) *string { return v.UtilityStorageAccountId }).(pulumi.StringPtrOutput)
}

type MigrateProjectPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MigrateProjectPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectProperties)(nil)).Elem()
}

func (o MigrateProjectPropertiesPtrOutput) ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput {
	return o
}

func (o MigrateProjectPropertiesPtrOutput) ToMigrateProjectPropertiesPtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesPtrOutput {
	return o
}

func (o MigrateProjectPropertiesPtrOutput) Elem() MigrateProjectPropertiesOutput {
	return o.ApplyT(func(v *MigrateProjectProperties) MigrateProjectProperties {
		if v != nil {
			return *v
		}
		var ret MigrateProjectProperties
		return ret
	}).(MigrateProjectPropertiesOutput)
}

func (o MigrateProjectPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesPtrOutput) ServiceEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectProperties) *string {
		if v == nil {
			return nil
		}
		return v.ServiceEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesPtrOutput) UtilityStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectProperties) *string {
		if v == nil {
			return nil
		}
		return v.UtilityStorageAccountId
	}).(pulumi.StringPtrOutput)
}

type MigrateProjectPropertiesResponse struct {
	LastSummaryRefreshedTime   string                              `pulumi:"lastSummaryRefreshedTime"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess        *string                             `pulumi:"publicNetworkAccess"`
	RefreshSummaryState        string                              `pulumi:"refreshSummaryState"`
	RegisteredTools            []string                            `pulumi:"registeredTools"`
	ServiceEndpoint            *string                             `pulumi:"serviceEndpoint"`
	Summary                    map[string]ProjectSummaryResponse   `pulumi:"summary"`
	UtilityStorageAccountId    *string                             `pulumi:"utilityStorageAccountId"`
}

type MigrateProjectPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MigrateProjectPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectPropertiesResponse)(nil)).Elem()
}

func (o MigrateProjectPropertiesResponseOutput) ToMigrateProjectPropertiesResponseOutput() MigrateProjectPropertiesResponseOutput {
	return o
}

func (o MigrateProjectPropertiesResponseOutput) ToMigrateProjectPropertiesResponseOutputWithContext(ctx context.Context) MigrateProjectPropertiesResponseOutput {
	return o
}

func (o MigrateProjectPropertiesResponseOutput) LastSummaryRefreshedTime() pulumi.StringOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) string { return v.LastSummaryRefreshedTime }).(pulumi.StringOutput)
}

func (o MigrateProjectPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o MigrateProjectPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesResponseOutput) RefreshSummaryState() pulumi.StringOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) string { return v.RefreshSummaryState }).(pulumi.StringOutput)
}

func (o MigrateProjectPropertiesResponseOutput) RegisteredTools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) []string { return v.RegisteredTools }).(pulumi.StringArrayOutput)
}

func (o MigrateProjectPropertiesResponseOutput) ServiceEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) *string { return v.ServiceEndpoint }).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesResponseOutput) Summary() ProjectSummaryResponseMapOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) map[string]ProjectSummaryResponse { return v.Summary }).(ProjectSummaryResponseMapOutput)
}

func (o MigrateProjectPropertiesResponseOutput) UtilityStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) *string { return v.UtilityStorageAccountId }).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   ResourceIdResponse                         `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() ResourceIdResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) ResourceIdResponse { return v.PrivateEndpoint }).(ResourceIdResponseOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponse struct {
	ETag       string                                      `pulumi:"eTag"`
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                          `pulumi:"systemData"`
	Type       string                                      `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ETag }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
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

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
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

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
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

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
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

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ProjectSummaryResponse struct {
	ExtendedSummary          map[string]string `pulumi:"extendedSummary"`
	InstanceType             string            `pulumi:"instanceType"`
	LastSummaryRefreshedTime *string           `pulumi:"lastSummaryRefreshedTime"`
	RefreshSummaryState      *string           `pulumi:"refreshSummaryState"`
}

type ProjectSummaryResponseOutput struct{ *pulumi.OutputState }

func (ProjectSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectSummaryResponse)(nil)).Elem()
}

func (o ProjectSummaryResponseOutput) ToProjectSummaryResponseOutput() ProjectSummaryResponseOutput {
	return o
}

func (o ProjectSummaryResponseOutput) ToProjectSummaryResponseOutputWithContext(ctx context.Context) ProjectSummaryResponseOutput {
	return o
}

func (o ProjectSummaryResponseOutput) ExtendedSummary() pulumi.StringMapOutput {
	return o.ApplyT(func(v ProjectSummaryResponse) map[string]string { return v.ExtendedSummary }).(pulumi.StringMapOutput)
}

func (o ProjectSummaryResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectSummaryResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o ProjectSummaryResponseOutput) LastSummaryRefreshedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectSummaryResponse) *string { return v.LastSummaryRefreshedTime }).(pulumi.StringPtrOutput)
}

func (o ProjectSummaryResponseOutput) RefreshSummaryState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectSummaryResponse) *string { return v.RefreshSummaryState }).(pulumi.StringPtrOutput)
}

type ProjectSummaryResponseMapOutput struct{ *pulumi.OutputState }

func (ProjectSummaryResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ProjectSummaryResponse)(nil)).Elem()
}

func (o ProjectSummaryResponseMapOutput) ToProjectSummaryResponseMapOutput() ProjectSummaryResponseMapOutput {
	return o
}

func (o ProjectSummaryResponseMapOutput) ToProjectSummaryResponseMapOutputWithContext(ctx context.Context) ProjectSummaryResponseMapOutput {
	return o
}

func (o ProjectSummaryResponseMapOutput) MapIndex(k pulumi.StringInput) ProjectSummaryResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ProjectSummaryResponse {
		return vs[0].(map[string]ProjectSummaryResponse)[vs[1].(string)]
	}).(ProjectSummaryResponseOutput)
}

type ResourceIdResponse struct {
	Id string `pulumi:"id"`
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

func (o ResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
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

func init() {
	pulumi.RegisterOutputType(ConnectionStateRequestBodyPropertiesOutput{})
	pulumi.RegisterOutputType(ConnectionStateRequestBodyPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MigrateProjectPropertiesOutput{})
	pulumi.RegisterOutputType(MigrateProjectPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MigrateProjectPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ProjectSummaryResponseOutput{})
	pulumi.RegisterOutputType(ProjectSummaryResponseMapOutput{})
	pulumi.RegisterOutputType(ResourceIdResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

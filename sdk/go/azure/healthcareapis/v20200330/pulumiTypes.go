


package v20200330

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionType struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}





type PrivateEndpointConnectionTypeArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput
	ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Context) PrivateEndpointConnectionTypeArrayOutput
}

type PrivateEndpointConnectionTypeArray []PrivateEndpointConnectionTypeInput

func (PrivateEndpointConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return i.ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeArrayOutput)
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateEndpointConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionType {
		return vs[0].([]PrivateEndpointConnectionType)[vs[1].(int)]
	}).(PrivateEndpointConnectionTypeOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
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

type ServiceAccessPolicyEntry struct {
	ObjectId string `pulumi:"objectId"`
}





type ServiceAccessPolicyEntryInput interface {
	pulumi.Input

	ToServiceAccessPolicyEntryOutput() ServiceAccessPolicyEntryOutput
	ToServiceAccessPolicyEntryOutputWithContext(context.Context) ServiceAccessPolicyEntryOutput
}

type ServiceAccessPolicyEntryArgs struct {
	ObjectId pulumi.StringInput `pulumi:"objectId"`
}

func (ServiceAccessPolicyEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAccessPolicyEntry)(nil)).Elem()
}

func (i ServiceAccessPolicyEntryArgs) ToServiceAccessPolicyEntryOutput() ServiceAccessPolicyEntryOutput {
	return i.ToServiceAccessPolicyEntryOutputWithContext(context.Background())
}

func (i ServiceAccessPolicyEntryArgs) ToServiceAccessPolicyEntryOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccessPolicyEntryOutput)
}





type ServiceAccessPolicyEntryArrayInput interface {
	pulumi.Input

	ToServiceAccessPolicyEntryArrayOutput() ServiceAccessPolicyEntryArrayOutput
	ToServiceAccessPolicyEntryArrayOutputWithContext(context.Context) ServiceAccessPolicyEntryArrayOutput
}

type ServiceAccessPolicyEntryArray []ServiceAccessPolicyEntryInput

func (ServiceAccessPolicyEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceAccessPolicyEntry)(nil)).Elem()
}

func (i ServiceAccessPolicyEntryArray) ToServiceAccessPolicyEntryArrayOutput() ServiceAccessPolicyEntryArrayOutput {
	return i.ToServiceAccessPolicyEntryArrayOutputWithContext(context.Background())
}

func (i ServiceAccessPolicyEntryArray) ToServiceAccessPolicyEntryArrayOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccessPolicyEntryArrayOutput)
}

type ServiceAccessPolicyEntryOutput struct{ *pulumi.OutputState }

func (ServiceAccessPolicyEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAccessPolicyEntry)(nil)).Elem()
}

func (o ServiceAccessPolicyEntryOutput) ToServiceAccessPolicyEntryOutput() ServiceAccessPolicyEntryOutput {
	return o
}

func (o ServiceAccessPolicyEntryOutput) ToServiceAccessPolicyEntryOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryOutput {
	return o
}

func (o ServiceAccessPolicyEntryOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAccessPolicyEntry) string { return v.ObjectId }).(pulumi.StringOutput)
}

type ServiceAccessPolicyEntryArrayOutput struct{ *pulumi.OutputState }

func (ServiceAccessPolicyEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceAccessPolicyEntry)(nil)).Elem()
}

func (o ServiceAccessPolicyEntryArrayOutput) ToServiceAccessPolicyEntryArrayOutput() ServiceAccessPolicyEntryArrayOutput {
	return o
}

func (o ServiceAccessPolicyEntryArrayOutput) ToServiceAccessPolicyEntryArrayOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryArrayOutput {
	return o
}

func (o ServiceAccessPolicyEntryArrayOutput) Index(i pulumi.IntInput) ServiceAccessPolicyEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceAccessPolicyEntry {
		return vs[0].([]ServiceAccessPolicyEntry)[vs[1].(int)]
	}).(ServiceAccessPolicyEntryOutput)
}

type ServiceAccessPolicyEntryResponse struct {
	ObjectId string `pulumi:"objectId"`
}

type ServiceAccessPolicyEntryResponseOutput struct{ *pulumi.OutputState }

func (ServiceAccessPolicyEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAccessPolicyEntryResponse)(nil)).Elem()
}

func (o ServiceAccessPolicyEntryResponseOutput) ToServiceAccessPolicyEntryResponseOutput() ServiceAccessPolicyEntryResponseOutput {
	return o
}

func (o ServiceAccessPolicyEntryResponseOutput) ToServiceAccessPolicyEntryResponseOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryResponseOutput {
	return o
}

func (o ServiceAccessPolicyEntryResponseOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAccessPolicyEntryResponse) string { return v.ObjectId }).(pulumi.StringOutput)
}

type ServiceAccessPolicyEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceAccessPolicyEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceAccessPolicyEntryResponse)(nil)).Elem()
}

func (o ServiceAccessPolicyEntryResponseArrayOutput) ToServiceAccessPolicyEntryResponseArrayOutput() ServiceAccessPolicyEntryResponseArrayOutput {
	return o
}

func (o ServiceAccessPolicyEntryResponseArrayOutput) ToServiceAccessPolicyEntryResponseArrayOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryResponseArrayOutput {
	return o
}

func (o ServiceAccessPolicyEntryResponseArrayOutput) Index(i pulumi.IntInput) ServiceAccessPolicyEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceAccessPolicyEntryResponse {
		return vs[0].([]ServiceAccessPolicyEntryResponse)[vs[1].(int)]
	}).(ServiceAccessPolicyEntryResponseOutput)
}

type ServiceAuthenticationConfigurationInfo struct {
	Audience          *string `pulumi:"audience"`
	Authority         *string `pulumi:"authority"`
	SmartProxyEnabled *bool   `pulumi:"smartProxyEnabled"`
}





type ServiceAuthenticationConfigurationInfoInput interface {
	pulumi.Input

	ToServiceAuthenticationConfigurationInfoOutput() ServiceAuthenticationConfigurationInfoOutput
	ToServiceAuthenticationConfigurationInfoOutputWithContext(context.Context) ServiceAuthenticationConfigurationInfoOutput
}

type ServiceAuthenticationConfigurationInfoArgs struct {
	Audience          pulumi.StringPtrInput `pulumi:"audience"`
	Authority         pulumi.StringPtrInput `pulumi:"authority"`
	SmartProxyEnabled pulumi.BoolPtrInput   `pulumi:"smartProxyEnabled"`
}

func (ServiceAuthenticationConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthenticationConfigurationInfo)(nil)).Elem()
}

func (i ServiceAuthenticationConfigurationInfoArgs) ToServiceAuthenticationConfigurationInfoOutput() ServiceAuthenticationConfigurationInfoOutput {
	return i.ToServiceAuthenticationConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceAuthenticationConfigurationInfoArgs) ToServiceAuthenticationConfigurationInfoOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationInfoOutput)
}

func (i ServiceAuthenticationConfigurationInfoArgs) ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput {
	return i.ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceAuthenticationConfigurationInfoArgs) ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationInfoOutput).ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceAuthenticationConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput
	ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(context.Context) ServiceAuthenticationConfigurationInfoPtrOutput
}

type serviceAuthenticationConfigurationInfoPtrType ServiceAuthenticationConfigurationInfoArgs

func ServiceAuthenticationConfigurationInfoPtr(v *ServiceAuthenticationConfigurationInfoArgs) ServiceAuthenticationConfigurationInfoPtrInput {
	return (*serviceAuthenticationConfigurationInfoPtrType)(v)
}

func (*serviceAuthenticationConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthenticationConfigurationInfo)(nil)).Elem()
}

func (i *serviceAuthenticationConfigurationInfoPtrType) ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput {
	return i.ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceAuthenticationConfigurationInfoPtrType) ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationInfoPtrOutput)
}

type ServiceAuthenticationConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthenticationConfigurationInfo)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationInfoOutput) ToServiceAuthenticationConfigurationInfoOutput() ServiceAuthenticationConfigurationInfoOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoOutput) ToServiceAuthenticationConfigurationInfoOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoOutput) ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput {
	return o.ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceAuthenticationConfigurationInfoOutput) ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceAuthenticationConfigurationInfo) *ServiceAuthenticationConfigurationInfo {
		return &v
	}).(ServiceAuthenticationConfigurationInfoPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfo) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfo) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfo) *bool { return v.SmartProxyEnabled }).(pulumi.BoolPtrOutput)
}

type ServiceAuthenticationConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthenticationConfigurationInfo)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoPtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) Elem() ServiceAuthenticationConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfo) ServiceAuthenticationConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceAuthenticationConfigurationInfo
		return ret
	}).(ServiceAuthenticationConfigurationInfoOutput)
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfo) *bool {
		if v == nil {
			return nil
		}
		return v.SmartProxyEnabled
	}).(pulumi.BoolPtrOutput)
}

type ServiceAuthenticationConfigurationInfoResponse struct {
	Audience          *string `pulumi:"audience"`
	Authority         *string `pulumi:"authority"`
	SmartProxyEnabled *bool   `pulumi:"smartProxyEnabled"`
}

type ServiceAuthenticationConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthenticationConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) ToServiceAuthenticationConfigurationInfoResponseOutput() ServiceAuthenticationConfigurationInfoResponseOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) ToServiceAuthenticationConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoResponseOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfoResponse) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfoResponse) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfoResponse) *bool { return v.SmartProxyEnabled }).(pulumi.BoolPtrOutput)
}

type ServiceAuthenticationConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthenticationConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) ToServiceAuthenticationConfigurationInfoResponsePtrOutput() ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) ToServiceAuthenticationConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) Elem() ServiceAuthenticationConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfoResponse) ServiceAuthenticationConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceAuthenticationConfigurationInfoResponse
		return ret
	}).(ServiceAuthenticationConfigurationInfoResponseOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SmartProxyEnabled
	}).(pulumi.BoolPtrOutput)
}

type ServiceCorsConfigurationInfo struct {
	AllowCredentials *bool    `pulumi:"allowCredentials"`
	Headers          []string `pulumi:"headers"`
	MaxAge           *int     `pulumi:"maxAge"`
	Methods          []string `pulumi:"methods"`
	Origins          []string `pulumi:"origins"`
}





type ServiceCorsConfigurationInfoInput interface {
	pulumi.Input

	ToServiceCorsConfigurationInfoOutput() ServiceCorsConfigurationInfoOutput
	ToServiceCorsConfigurationInfoOutputWithContext(context.Context) ServiceCorsConfigurationInfoOutput
}

type ServiceCorsConfigurationInfoArgs struct {
	AllowCredentials pulumi.BoolPtrInput     `pulumi:"allowCredentials"`
	Headers          pulumi.StringArrayInput `pulumi:"headers"`
	MaxAge           pulumi.IntPtrInput      `pulumi:"maxAge"`
	Methods          pulumi.StringArrayInput `pulumi:"methods"`
	Origins          pulumi.StringArrayInput `pulumi:"origins"`
}

func (ServiceCorsConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorsConfigurationInfo)(nil)).Elem()
}

func (i ServiceCorsConfigurationInfoArgs) ToServiceCorsConfigurationInfoOutput() ServiceCorsConfigurationInfoOutput {
	return i.ToServiceCorsConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceCorsConfigurationInfoArgs) ToServiceCorsConfigurationInfoOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationInfoOutput)
}

func (i ServiceCorsConfigurationInfoArgs) ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput {
	return i.ToServiceCorsConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceCorsConfigurationInfoArgs) ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationInfoOutput).ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceCorsConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput
	ToServiceCorsConfigurationInfoPtrOutputWithContext(context.Context) ServiceCorsConfigurationInfoPtrOutput
}

type serviceCorsConfigurationInfoPtrType ServiceCorsConfigurationInfoArgs

func ServiceCorsConfigurationInfoPtr(v *ServiceCorsConfigurationInfoArgs) ServiceCorsConfigurationInfoPtrInput {
	return (*serviceCorsConfigurationInfoPtrType)(v)
}

func (*serviceCorsConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorsConfigurationInfo)(nil)).Elem()
}

func (i *serviceCorsConfigurationInfoPtrType) ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput {
	return i.ToServiceCorsConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceCorsConfigurationInfoPtrType) ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationInfoPtrOutput)
}

type ServiceCorsConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorsConfigurationInfo)(nil)).Elem()
}

func (o ServiceCorsConfigurationInfoOutput) ToServiceCorsConfigurationInfoOutput() ServiceCorsConfigurationInfoOutput {
	return o
}

func (o ServiceCorsConfigurationInfoOutput) ToServiceCorsConfigurationInfoOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoOutput {
	return o
}

func (o ServiceCorsConfigurationInfoOutput) ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput {
	return o.ToServiceCorsConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceCorsConfigurationInfoOutput) ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceCorsConfigurationInfo) *ServiceCorsConfigurationInfo {
		return &v
	}).(ServiceCorsConfigurationInfoPtrOutput)
}

func (o ServiceCorsConfigurationInfoOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

func (o ServiceCorsConfigurationInfoOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) []string { return v.Headers }).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) *int { return v.MaxAge }).(pulumi.IntPtrOutput)
}

func (o ServiceCorsConfigurationInfoOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) []string { return v.Methods }).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) []string { return v.Origins }).(pulumi.StringArrayOutput)
}

type ServiceCorsConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorsConfigurationInfo)(nil)).Elem()
}

func (o ServiceCorsConfigurationInfoPtrOutput) ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput {
	return o
}

func (o ServiceCorsConfigurationInfoPtrOutput) ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoPtrOutput {
	return o
}

func (o ServiceCorsConfigurationInfoPtrOutput) Elem() ServiceCorsConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) ServiceCorsConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceCorsConfigurationInfo
		return ret
	}).(ServiceCorsConfigurationInfoOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) *bool {
		if v == nil {
			return nil
		}
		return v.AllowCredentials
	}).(pulumi.BoolPtrOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) []string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) *int {
		if v == nil {
			return nil
		}
		return v.MaxAge
	}).(pulumi.IntPtrOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) []string {
		if v == nil {
			return nil
		}
		return v.Methods
	}).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) []string {
		if v == nil {
			return nil
		}
		return v.Origins
	}).(pulumi.StringArrayOutput)
}

type ServiceCorsConfigurationInfoResponse struct {
	AllowCredentials *bool    `pulumi:"allowCredentials"`
	Headers          []string `pulumi:"headers"`
	MaxAge           *int     `pulumi:"maxAge"`
	Methods          []string `pulumi:"methods"`
	Origins          []string `pulumi:"origins"`
}

type ServiceCorsConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorsConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceCorsConfigurationInfoResponseOutput) ToServiceCorsConfigurationInfoResponseOutput() ServiceCorsConfigurationInfoResponseOutput {
	return o
}

func (o ServiceCorsConfigurationInfoResponseOutput) ToServiceCorsConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoResponseOutput {
	return o
}

func (o ServiceCorsConfigurationInfoResponseOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) []string { return v.Headers }).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) *int { return v.MaxAge }).(pulumi.IntPtrOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) []string { return v.Methods }).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) []string { return v.Origins }).(pulumi.StringArrayOutput)
}

type ServiceCorsConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorsConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) ToServiceCorsConfigurationInfoResponsePtrOutput() ServiceCorsConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) ToServiceCorsConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) Elem() ServiceCorsConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) ServiceCorsConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceCorsConfigurationInfoResponse
		return ret
	}).(ServiceCorsConfigurationInfoResponseOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowCredentials
	}).(pulumi.BoolPtrOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxAge
	}).(pulumi.IntPtrOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.Methods
	}).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.Origins
	}).(pulumi.StringArrayOutput)
}

type ServiceCosmosDbConfigurationInfo struct {
	KeyVaultKeyUri  *string `pulumi:"keyVaultKeyUri"`
	OfferThroughput *int    `pulumi:"offerThroughput"`
}





type ServiceCosmosDbConfigurationInfoInput interface {
	pulumi.Input

	ToServiceCosmosDbConfigurationInfoOutput() ServiceCosmosDbConfigurationInfoOutput
	ToServiceCosmosDbConfigurationInfoOutputWithContext(context.Context) ServiceCosmosDbConfigurationInfoOutput
}

type ServiceCosmosDbConfigurationInfoArgs struct {
	KeyVaultKeyUri  pulumi.StringPtrInput `pulumi:"keyVaultKeyUri"`
	OfferThroughput pulumi.IntPtrInput    `pulumi:"offerThroughput"`
}

func (ServiceCosmosDbConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCosmosDbConfigurationInfo)(nil)).Elem()
}

func (i ServiceCosmosDbConfigurationInfoArgs) ToServiceCosmosDbConfigurationInfoOutput() ServiceCosmosDbConfigurationInfoOutput {
	return i.ToServiceCosmosDbConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceCosmosDbConfigurationInfoArgs) ToServiceCosmosDbConfigurationInfoOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCosmosDbConfigurationInfoOutput)
}

func (i ServiceCosmosDbConfigurationInfoArgs) ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput {
	return i.ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceCosmosDbConfigurationInfoArgs) ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCosmosDbConfigurationInfoOutput).ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceCosmosDbConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput
	ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(context.Context) ServiceCosmosDbConfigurationInfoPtrOutput
}

type serviceCosmosDbConfigurationInfoPtrType ServiceCosmosDbConfigurationInfoArgs

func ServiceCosmosDbConfigurationInfoPtr(v *ServiceCosmosDbConfigurationInfoArgs) ServiceCosmosDbConfigurationInfoPtrInput {
	return (*serviceCosmosDbConfigurationInfoPtrType)(v)
}

func (*serviceCosmosDbConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCosmosDbConfigurationInfo)(nil)).Elem()
}

func (i *serviceCosmosDbConfigurationInfoPtrType) ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput {
	return i.ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceCosmosDbConfigurationInfoPtrType) ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCosmosDbConfigurationInfoPtrOutput)
}

type ServiceCosmosDbConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceCosmosDbConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCosmosDbConfigurationInfo)(nil)).Elem()
}

func (o ServiceCosmosDbConfigurationInfoOutput) ToServiceCosmosDbConfigurationInfoOutput() ServiceCosmosDbConfigurationInfoOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoOutput) ToServiceCosmosDbConfigurationInfoOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoOutput) ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput {
	return o.ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceCosmosDbConfigurationInfoOutput) ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceCosmosDbConfigurationInfo) *ServiceCosmosDbConfigurationInfo {
		return &v
	}).(ServiceCosmosDbConfigurationInfoPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceCosmosDbConfigurationInfo) *string { return v.KeyVaultKeyUri }).(pulumi.StringPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoOutput) OfferThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceCosmosDbConfigurationInfo) *int { return v.OfferThroughput }).(pulumi.IntPtrOutput)
}

type ServiceCosmosDbConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceCosmosDbConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCosmosDbConfigurationInfo)(nil)).Elem()
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoPtrOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) Elem() ServiceCosmosDbConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfo) ServiceCosmosDbConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceCosmosDbConfigurationInfo
		return ret
	}).(ServiceCosmosDbConfigurationInfoOutput)
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultKeyUri
	}).(pulumi.StringPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) OfferThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfo) *int {
		if v == nil {
			return nil
		}
		return v.OfferThroughput
	}).(pulumi.IntPtrOutput)
}

type ServiceCosmosDbConfigurationInfoResponse struct {
	KeyVaultKeyUri  *string `pulumi:"keyVaultKeyUri"`
	OfferThroughput *int    `pulumi:"offerThroughput"`
}

type ServiceCosmosDbConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceCosmosDbConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCosmosDbConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) ToServiceCosmosDbConfigurationInfoResponseOutput() ServiceCosmosDbConfigurationInfoResponseOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) ToServiceCosmosDbConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoResponseOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceCosmosDbConfigurationInfoResponse) *string { return v.KeyVaultKeyUri }).(pulumi.StringPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) OfferThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceCosmosDbConfigurationInfoResponse) *int { return v.OfferThroughput }).(pulumi.IntPtrOutput)
}

type ServiceCosmosDbConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceCosmosDbConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCosmosDbConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) ToServiceCosmosDbConfigurationInfoResponsePtrOutput() ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) ToServiceCosmosDbConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) Elem() ServiceCosmosDbConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfoResponse) ServiceCosmosDbConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceCosmosDbConfigurationInfoResponse
		return ret
	}).(ServiceCosmosDbConfigurationInfoResponseOutput)
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultKeyUri
	}).(pulumi.StringPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) OfferThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.OfferThroughput
	}).(pulumi.IntPtrOutput)
}

type ServiceExportConfigurationInfo struct {
	StorageAccountName *string `pulumi:"storageAccountName"`
}





type ServiceExportConfigurationInfoInput interface {
	pulumi.Input

	ToServiceExportConfigurationInfoOutput() ServiceExportConfigurationInfoOutput
	ToServiceExportConfigurationInfoOutputWithContext(context.Context) ServiceExportConfigurationInfoOutput
}

type ServiceExportConfigurationInfoArgs struct {
	StorageAccountName pulumi.StringPtrInput `pulumi:"storageAccountName"`
}

func (ServiceExportConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceExportConfigurationInfo)(nil)).Elem()
}

func (i ServiceExportConfigurationInfoArgs) ToServiceExportConfigurationInfoOutput() ServiceExportConfigurationInfoOutput {
	return i.ToServiceExportConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceExportConfigurationInfoArgs) ToServiceExportConfigurationInfoOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceExportConfigurationInfoOutput)
}

func (i ServiceExportConfigurationInfoArgs) ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput {
	return i.ToServiceExportConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceExportConfigurationInfoArgs) ToServiceExportConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceExportConfigurationInfoOutput).ToServiceExportConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceExportConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput
	ToServiceExportConfigurationInfoPtrOutputWithContext(context.Context) ServiceExportConfigurationInfoPtrOutput
}

type serviceExportConfigurationInfoPtrType ServiceExportConfigurationInfoArgs

func ServiceExportConfigurationInfoPtr(v *ServiceExportConfigurationInfoArgs) ServiceExportConfigurationInfoPtrInput {
	return (*serviceExportConfigurationInfoPtrType)(v)
}

func (*serviceExportConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceExportConfigurationInfo)(nil)).Elem()
}

func (i *serviceExportConfigurationInfoPtrType) ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput {
	return i.ToServiceExportConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceExportConfigurationInfoPtrType) ToServiceExportConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceExportConfigurationInfoPtrOutput)
}

type ServiceExportConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceExportConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceExportConfigurationInfo)(nil)).Elem()
}

func (o ServiceExportConfigurationInfoOutput) ToServiceExportConfigurationInfoOutput() ServiceExportConfigurationInfoOutput {
	return o
}

func (o ServiceExportConfigurationInfoOutput) ToServiceExportConfigurationInfoOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoOutput {
	return o
}

func (o ServiceExportConfigurationInfoOutput) ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput {
	return o.ToServiceExportConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceExportConfigurationInfoOutput) ToServiceExportConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceExportConfigurationInfo) *ServiceExportConfigurationInfo {
		return &v
	}).(ServiceExportConfigurationInfoPtrOutput)
}

func (o ServiceExportConfigurationInfoOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceExportConfigurationInfo) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

type ServiceExportConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceExportConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceExportConfigurationInfo)(nil)).Elem()
}

func (o ServiceExportConfigurationInfoPtrOutput) ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput {
	return o
}

func (o ServiceExportConfigurationInfoPtrOutput) ToServiceExportConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoPtrOutput {
	return o
}

func (o ServiceExportConfigurationInfoPtrOutput) Elem() ServiceExportConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceExportConfigurationInfo) ServiceExportConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceExportConfigurationInfo
		return ret
	}).(ServiceExportConfigurationInfoOutput)
}

func (o ServiceExportConfigurationInfoPtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceExportConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

type ServiceExportConfigurationInfoResponse struct {
	StorageAccountName *string `pulumi:"storageAccountName"`
}

type ServiceExportConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceExportConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceExportConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceExportConfigurationInfoResponseOutput) ToServiceExportConfigurationInfoResponseOutput() ServiceExportConfigurationInfoResponseOutput {
	return o
}

func (o ServiceExportConfigurationInfoResponseOutput) ToServiceExportConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoResponseOutput {
	return o
}

func (o ServiceExportConfigurationInfoResponseOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceExportConfigurationInfoResponse) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

type ServiceExportConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceExportConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceExportConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceExportConfigurationInfoResponsePtrOutput) ToServiceExportConfigurationInfoResponsePtrOutput() ServiceExportConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceExportConfigurationInfoResponsePtrOutput) ToServiceExportConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceExportConfigurationInfoResponsePtrOutput) Elem() ServiceExportConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceExportConfigurationInfoResponse) ServiceExportConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceExportConfigurationInfoResponse
		return ret
	}).(ServiceExportConfigurationInfoResponseOutput)
}

func (o ServiceExportConfigurationInfoResponsePtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceExportConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

type ServicesProperties struct {
	AccessPolicies              []ServiceAccessPolicyEntry              `pulumi:"accessPolicies"`
	AuthenticationConfiguration *ServiceAuthenticationConfigurationInfo `pulumi:"authenticationConfiguration"`
	CorsConfiguration           *ServiceCorsConfigurationInfo           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       *ServiceCosmosDbConfigurationInfo       `pulumi:"cosmosDbConfiguration"`
	ExportConfiguration         *ServiceExportConfigurationInfo         `pulumi:"exportConfiguration"`
	PrivateEndpointConnections  []PrivateEndpointConnectionType         `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess         *string                                 `pulumi:"publicNetworkAccess"`
}





type ServicesPropertiesInput interface {
	pulumi.Input

	ToServicesPropertiesOutput() ServicesPropertiesOutput
	ToServicesPropertiesOutputWithContext(context.Context) ServicesPropertiesOutput
}

type ServicesPropertiesArgs struct {
	AccessPolicies              ServiceAccessPolicyEntryArrayInput             `pulumi:"accessPolicies"`
	AuthenticationConfiguration ServiceAuthenticationConfigurationInfoPtrInput `pulumi:"authenticationConfiguration"`
	CorsConfiguration           ServiceCorsConfigurationInfoPtrInput           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       ServiceCosmosDbConfigurationInfoPtrInput       `pulumi:"cosmosDbConfiguration"`
	ExportConfiguration         ServiceExportConfigurationInfoPtrInput         `pulumi:"exportConfiguration"`
	PrivateEndpointConnections  PrivateEndpointConnectionTypeArrayInput        `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess         pulumi.StringPtrInput                          `pulumi:"publicNetworkAccess"`
}

func (ServicesPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesProperties)(nil)).Elem()
}

func (i ServicesPropertiesArgs) ToServicesPropertiesOutput() ServicesPropertiesOutput {
	return i.ToServicesPropertiesOutputWithContext(context.Background())
}

func (i ServicesPropertiesArgs) ToServicesPropertiesOutputWithContext(ctx context.Context) ServicesPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesPropertiesOutput)
}

func (i ServicesPropertiesArgs) ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput {
	return i.ToServicesPropertiesPtrOutputWithContext(context.Background())
}

func (i ServicesPropertiesArgs) ToServicesPropertiesPtrOutputWithContext(ctx context.Context) ServicesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesPropertiesOutput).ToServicesPropertiesPtrOutputWithContext(ctx)
}









type ServicesPropertiesPtrInput interface {
	pulumi.Input

	ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput
	ToServicesPropertiesPtrOutputWithContext(context.Context) ServicesPropertiesPtrOutput
}

type servicesPropertiesPtrType ServicesPropertiesArgs

func ServicesPropertiesPtr(v *ServicesPropertiesArgs) ServicesPropertiesPtrInput {
	return (*servicesPropertiesPtrType)(v)
}

func (*servicesPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesProperties)(nil)).Elem()
}

func (i *servicesPropertiesPtrType) ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput {
	return i.ToServicesPropertiesPtrOutputWithContext(context.Background())
}

func (i *servicesPropertiesPtrType) ToServicesPropertiesPtrOutputWithContext(ctx context.Context) ServicesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesPropertiesPtrOutput)
}

type ServicesPropertiesOutput struct{ *pulumi.OutputState }

func (ServicesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesProperties)(nil)).Elem()
}

func (o ServicesPropertiesOutput) ToServicesPropertiesOutput() ServicesPropertiesOutput {
	return o
}

func (o ServicesPropertiesOutput) ToServicesPropertiesOutputWithContext(ctx context.Context) ServicesPropertiesOutput {
	return o
}

func (o ServicesPropertiesOutput) ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput {
	return o.ToServicesPropertiesPtrOutputWithContext(context.Background())
}

func (o ServicesPropertiesOutput) ToServicesPropertiesPtrOutputWithContext(ctx context.Context) ServicesPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicesProperties) *ServicesProperties {
		return &v
	}).(ServicesPropertiesPtrOutput)
}

func (o ServicesPropertiesOutput) AccessPolicies() ServiceAccessPolicyEntryArrayOutput {
	return o.ApplyT(func(v ServicesProperties) []ServiceAccessPolicyEntry { return v.AccessPolicies }).(ServiceAccessPolicyEntryArrayOutput)
}

func (o ServicesPropertiesOutput) AuthenticationConfiguration() ServiceAuthenticationConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceAuthenticationConfigurationInfo {
		return v.AuthenticationConfiguration
	}).(ServiceAuthenticationConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) CorsConfiguration() ServiceCorsConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceCorsConfigurationInfo { return v.CorsConfiguration }).(ServiceCorsConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) CosmosDbConfiguration() ServiceCosmosDbConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceCosmosDbConfigurationInfo { return v.CosmosDbConfiguration }).(ServiceCosmosDbConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) ExportConfiguration() ServiceExportConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceExportConfigurationInfo { return v.ExportConfiguration }).(ServiceExportConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) PrivateEndpointConnections() PrivateEndpointConnectionTypeArrayOutput {
	return o.ApplyT(func(v ServicesProperties) []PrivateEndpointConnectionType { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionTypeArrayOutput)
}

func (o ServicesPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type ServicesPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ServicesPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesProperties)(nil)).Elem()
}

func (o ServicesPropertiesPtrOutput) ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput {
	return o
}

func (o ServicesPropertiesPtrOutput) ToServicesPropertiesPtrOutputWithContext(ctx context.Context) ServicesPropertiesPtrOutput {
	return o
}

func (o ServicesPropertiesPtrOutput) Elem() ServicesPropertiesOutput {
	return o.ApplyT(func(v *ServicesProperties) ServicesProperties {
		if v != nil {
			return *v
		}
		var ret ServicesProperties
		return ret
	}).(ServicesPropertiesOutput)
}

func (o ServicesPropertiesPtrOutput) AccessPolicies() ServiceAccessPolicyEntryArrayOutput {
	return o.ApplyT(func(v *ServicesProperties) []ServiceAccessPolicyEntry {
		if v == nil {
			return nil
		}
		return v.AccessPolicies
	}).(ServiceAccessPolicyEntryArrayOutput)
}

func (o ServicesPropertiesPtrOutput) AuthenticationConfiguration() ServiceAuthenticationConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceAuthenticationConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.AuthenticationConfiguration
	}).(ServiceAuthenticationConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) CorsConfiguration() ServiceCorsConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceCorsConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.CorsConfiguration
	}).(ServiceCorsConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) CosmosDbConfiguration() ServiceCosmosDbConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceCosmosDbConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.CosmosDbConfiguration
	}).(ServiceCosmosDbConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) ExportConfiguration() ServiceExportConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceExportConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.ExportConfiguration
	}).(ServiceExportConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionTypeArrayOutput {
	return o.ApplyT(func(v *ServicesProperties) []PrivateEndpointConnectionType {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionTypeArrayOutput)
}

func (o ServicesPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type ServicesPropertiesResponse struct {
	AccessPolicies              []ServiceAccessPolicyEntryResponse              `pulumi:"accessPolicies"`
	AuthenticationConfiguration *ServiceAuthenticationConfigurationInfoResponse `pulumi:"authenticationConfiguration"`
	CorsConfiguration           *ServiceCorsConfigurationInfoResponse           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       *ServiceCosmosDbConfigurationInfoResponse       `pulumi:"cosmosDbConfiguration"`
	ExportConfiguration         *ServiceExportConfigurationInfoResponse         `pulumi:"exportConfiguration"`
	PrivateEndpointConnections  []PrivateEndpointConnectionResponse             `pulumi:"privateEndpointConnections"`
	ProvisioningState           string                                          `pulumi:"provisioningState"`
	PublicNetworkAccess         *string                                         `pulumi:"publicNetworkAccess"`
}

type ServicesPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServicesPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesPropertiesResponse)(nil)).Elem()
}

func (o ServicesPropertiesResponseOutput) ToServicesPropertiesResponseOutput() ServicesPropertiesResponseOutput {
	return o
}

func (o ServicesPropertiesResponseOutput) ToServicesPropertiesResponseOutputWithContext(ctx context.Context) ServicesPropertiesResponseOutput {
	return o
}

func (o ServicesPropertiesResponseOutput) AccessPolicies() ServiceAccessPolicyEntryResponseArrayOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) []ServiceAccessPolicyEntryResponse { return v.AccessPolicies }).(ServiceAccessPolicyEntryResponseArrayOutput)
}

func (o ServicesPropertiesResponseOutput) AuthenticationConfiguration() ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceAuthenticationConfigurationInfoResponse {
		return v.AuthenticationConfiguration
	}).(ServiceAuthenticationConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) CorsConfiguration() ServiceCorsConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceCorsConfigurationInfoResponse { return v.CorsConfiguration }).(ServiceCorsConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) CosmosDbConfiguration() ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceCosmosDbConfigurationInfoResponse {
		return v.CosmosDbConfiguration
	}).(ServiceCosmosDbConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) ExportConfiguration() ServiceExportConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceExportConfigurationInfoResponse {
		return v.ExportConfiguration
	}).(ServiceExportConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o ServicesPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServicesPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type ServicesResourceIdentity struct {
	Type *string `pulumi:"type"`
}





type ServicesResourceIdentityInput interface {
	pulumi.Input

	ToServicesResourceIdentityOutput() ServicesResourceIdentityOutput
	ToServicesResourceIdentityOutputWithContext(context.Context) ServicesResourceIdentityOutput
}

type ServicesResourceIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ServicesResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesResourceIdentity)(nil)).Elem()
}

func (i ServicesResourceIdentityArgs) ToServicesResourceIdentityOutput() ServicesResourceIdentityOutput {
	return i.ToServicesResourceIdentityOutputWithContext(context.Background())
}

func (i ServicesResourceIdentityArgs) ToServicesResourceIdentityOutputWithContext(ctx context.Context) ServicesResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesResourceIdentityOutput)
}

func (i ServicesResourceIdentityArgs) ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput {
	return i.ToServicesResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ServicesResourceIdentityArgs) ToServicesResourceIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesResourceIdentityOutput).ToServicesResourceIdentityPtrOutputWithContext(ctx)
}









type ServicesResourceIdentityPtrInput interface {
	pulumi.Input

	ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput
	ToServicesResourceIdentityPtrOutputWithContext(context.Context) ServicesResourceIdentityPtrOutput
}

type servicesResourceIdentityPtrType ServicesResourceIdentityArgs

func ServicesResourceIdentityPtr(v *ServicesResourceIdentityArgs) ServicesResourceIdentityPtrInput {
	return (*servicesResourceIdentityPtrType)(v)
}

func (*servicesResourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesResourceIdentity)(nil)).Elem()
}

func (i *servicesResourceIdentityPtrType) ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput {
	return i.ToServicesResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *servicesResourceIdentityPtrType) ToServicesResourceIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesResourceIdentityPtrOutput)
}

type ServicesResourceIdentityOutput struct{ *pulumi.OutputState }

func (ServicesResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesResourceIdentity)(nil)).Elem()
}

func (o ServicesResourceIdentityOutput) ToServicesResourceIdentityOutput() ServicesResourceIdentityOutput {
	return o
}

func (o ServicesResourceIdentityOutput) ToServicesResourceIdentityOutputWithContext(ctx context.Context) ServicesResourceIdentityOutput {
	return o
}

func (o ServicesResourceIdentityOutput) ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput {
	return o.ToServicesResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ServicesResourceIdentityOutput) ToServicesResourceIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicesResourceIdentity) *ServicesResourceIdentity {
		return &v
	}).(ServicesResourceIdentityPtrOutput)
}

func (o ServicesResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicesResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ServicesResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ServicesResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesResourceIdentity)(nil)).Elem()
}

func (o ServicesResourceIdentityPtrOutput) ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput {
	return o
}

func (o ServicesResourceIdentityPtrOutput) ToServicesResourceIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceIdentityPtrOutput {
	return o
}

func (o ServicesResourceIdentityPtrOutput) Elem() ServicesResourceIdentityOutput {
	return o.ApplyT(func(v *ServicesResourceIdentity) ServicesResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ServicesResourceIdentity
		return ret
	}).(ServicesResourceIdentityOutput)
}

func (o ServicesResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ServicesResourceResponseIdentity struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type ServicesResourceResponseIdentityOutput struct{ *pulumi.OutputState }

func (ServicesResourceResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesResourceResponseIdentity)(nil)).Elem()
}

func (o ServicesResourceResponseIdentityOutput) ToServicesResourceResponseIdentityOutput() ServicesResourceResponseIdentityOutput {
	return o
}

func (o ServicesResourceResponseIdentityOutput) ToServicesResourceResponseIdentityOutputWithContext(ctx context.Context) ServicesResourceResponseIdentityOutput {
	return o
}

func (o ServicesResourceResponseIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicesResourceResponseIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ServicesResourceResponseIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicesResourceResponseIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ServicesResourceResponseIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicesResourceResponseIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ServicesResourceResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (ServicesResourceResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesResourceResponseIdentity)(nil)).Elem()
}

func (o ServicesResourceResponseIdentityPtrOutput) ToServicesResourceResponseIdentityPtrOutput() ServicesResourceResponseIdentityPtrOutput {
	return o
}

func (o ServicesResourceResponseIdentityPtrOutput) ToServicesResourceResponseIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceResponseIdentityPtrOutput {
	return o
}

func (o ServicesResourceResponseIdentityPtrOutput) Elem() ServicesResourceResponseIdentityOutput {
	return o.ApplyT(func(v *ServicesResourceResponseIdentity) ServicesResourceResponseIdentity {
		if v != nil {
			return *v
		}
		var ret ServicesResourceResponseIdentity
		return ret
	}).(ServicesResourceResponseIdentityOutput)
}

func (o ServicesResourceResponseIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ServicesResourceResponseIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ServicesResourceResponseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryArrayOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryResponseOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceCosmosDbConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceCosmosDbConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceCosmosDbConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceCosmosDbConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceExportConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceExportConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceExportConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceExportConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServicesPropertiesOutput{})
	pulumi.RegisterOutputType(ServicesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ServicesPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServicesResourceIdentityOutput{})
	pulumi.RegisterOutputType(ServicesResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ServicesResourceResponseIdentityOutput{})
	pulumi.RegisterOutputType(ServicesResourceResponseIdentityPtrOutput{})
}




package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiProperties struct {
	ApiFreshnessWindowInMinutes *int `pulumi:"apiFreshnessWindowInMinutes"`
}





type ApiPropertiesInput interface {
	pulumi.Input

	ToApiPropertiesOutput() ApiPropertiesOutput
	ToApiPropertiesOutputWithContext(context.Context) ApiPropertiesOutput
}

type ApiPropertiesArgs struct {
	ApiFreshnessWindowInMinutes pulumi.IntPtrInput `pulumi:"apiFreshnessWindowInMinutes"`
}

func (ApiPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiProperties)(nil)).Elem()
}

func (i ApiPropertiesArgs) ToApiPropertiesOutput() ApiPropertiesOutput {
	return i.ToApiPropertiesOutputWithContext(context.Background())
}

func (i ApiPropertiesArgs) ToApiPropertiesOutputWithContext(ctx context.Context) ApiPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesOutput)
}





type ApiPropertiesMapInput interface {
	pulumi.Input

	ToApiPropertiesMapOutput() ApiPropertiesMapOutput
	ToApiPropertiesMapOutputWithContext(context.Context) ApiPropertiesMapOutput
}

type ApiPropertiesMap map[string]ApiPropertiesInput

func (ApiPropertiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiProperties)(nil)).Elem()
}

func (i ApiPropertiesMap) ToApiPropertiesMapOutput() ApiPropertiesMapOutput {
	return i.ToApiPropertiesMapOutputWithContext(context.Background())
}

func (i ApiPropertiesMap) ToApiPropertiesMapOutputWithContext(ctx context.Context) ApiPropertiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesMapOutput)
}

type ApiPropertiesOutput struct{ *pulumi.OutputState }

func (ApiPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiProperties)(nil)).Elem()
}

func (o ApiPropertiesOutput) ToApiPropertiesOutput() ApiPropertiesOutput {
	return o
}

func (o ApiPropertiesOutput) ToApiPropertiesOutputWithContext(ctx context.Context) ApiPropertiesOutput {
	return o
}

func (o ApiPropertiesOutput) ApiFreshnessWindowInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApiProperties) *int { return v.ApiFreshnessWindowInMinutes }).(pulumi.IntPtrOutput)
}

type ApiPropertiesMapOutput struct{ *pulumi.OutputState }

func (ApiPropertiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiProperties)(nil)).Elem()
}

func (o ApiPropertiesMapOutput) ToApiPropertiesMapOutput() ApiPropertiesMapOutput {
	return o
}

func (o ApiPropertiesMapOutput) ToApiPropertiesMapOutputWithContext(ctx context.Context) ApiPropertiesMapOutput {
	return o
}

func (o ApiPropertiesMapOutput) MapIndex(k pulumi.StringInput) ApiPropertiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiProperties {
		return vs[0].(map[string]ApiProperties)[vs[1].(string)]
	}).(ApiPropertiesOutput)
}

type ApiPropertiesResponse struct {
	ApiFreshnessWindowInMinutes *int `pulumi:"apiFreshnessWindowInMinutes"`
}

type ApiPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ApiPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPropertiesResponse)(nil)).Elem()
}

func (o ApiPropertiesResponseOutput) ToApiPropertiesResponseOutput() ApiPropertiesResponseOutput {
	return o
}

func (o ApiPropertiesResponseOutput) ToApiPropertiesResponseOutputWithContext(ctx context.Context) ApiPropertiesResponseOutput {
	return o
}

func (o ApiPropertiesResponseOutput) ApiFreshnessWindowInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *int { return v.ApiFreshnessWindowInMinutes }).(pulumi.IntPtrOutput)
}

type ApiPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (ApiPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiPropertiesResponse)(nil)).Elem()
}

func (o ApiPropertiesResponseMapOutput) ToApiPropertiesResponseMapOutput() ApiPropertiesResponseMapOutput {
	return o
}

func (o ApiPropertiesResponseMapOutput) ToApiPropertiesResponseMapOutputWithContext(ctx context.Context) ApiPropertiesResponseMapOutput {
	return o
}

func (o ApiPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) ApiPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiPropertiesResponse {
		return vs[0].(map[string]ApiPropertiesResponse)[vs[1].(string)]
	}).(ApiPropertiesResponseOutput)
}

type ErrorAdditionalInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}

type ErrorAdditionalInfoResponseOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

func (o ErrorAdditionalInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ErrorAdditionalInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) Index(i pulumi.IntInput) ErrorAdditionalInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorAdditionalInfoResponse {
		return vs[0].([]ErrorAdditionalInfoResponse)[vs[1].(int)]
	}).(ErrorAdditionalInfoResponseOutput)
}

type ErrorDetailResponse struct {
	AdditionalInfo []ErrorAdditionalInfoResponse `pulumi:"additionalInfo"`
	Code           string                        `pulumi:"code"`
	Details        []ErrorDetailResponse         `pulumi:"details"`
	Message        string                        `pulumi:"message"`
	Target         string                        `pulumi:"target"`
}

type ErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorAdditionalInfoResponse { return v.AdditionalInfo }).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorDetailResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponsePtrOutput) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return o
}

func (o ErrorDetailResponsePtrOutput) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return o
}

func (o ErrorDetailResponsePtrOutput) Elem() ErrorDetailResponseOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) ErrorDetailResponse {
		if v != nil {
			return *v
		}
		var ret ErrorDetailResponse
		return ret
	}).(ErrorDetailResponseOutput)
}

func (o ErrorDetailResponsePtrOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) []ErrorAdditionalInfoResponse {
		if v == nil {
			return nil
		}
		return v.AdditionalInfo
	}).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorDetailResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDetailResponsePtrOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) []ErrorDetailResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDetailResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Target
	}).(pulumi.StringPtrOutput)
}

type ErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) ErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDetailResponse {
		return vs[0].([]ErrorDetailResponse)[vs[1].(int)]
	}).(ErrorDetailResponseOutput)
}

type ErrorResponseResponse struct {
	Error *ErrorDetailResponse `pulumi:"error"`
}

type ErrorResponseResponseOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponseOutput() ErrorResponseResponseOutput {
	return o
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponseOutputWithContext(ctx context.Context) ErrorResponseResponseOutput {
	return o
}

func (o ErrorResponseResponseOutput) Error() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v ErrorResponseResponse) *ErrorDetailResponse { return v.Error }).(ErrorDetailResponsePtrOutput)
}

type ErrorResponseResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponsePtrOutput) ToErrorResponseResponsePtrOutput() ErrorResponseResponsePtrOutput {
	return o
}

func (o ErrorResponseResponsePtrOutput) ToErrorResponseResponsePtrOutputWithContext(ctx context.Context) ErrorResponseResponsePtrOutput {
	return o
}

func (o ErrorResponseResponsePtrOutput) Elem() ErrorResponseResponseOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) ErrorResponseResponse {
		if v != nil {
			return *v
		}
		var ret ErrorResponseResponse
		return ret
	}).(ErrorResponseResponseOutput)
}

func (o ErrorResponseResponsePtrOutput) Error() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) *ErrorDetailResponse {
		if v == nil {
			return nil
		}
		return v.Error
	}).(ErrorDetailResponsePtrOutput)
}

type Identity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	GroupIds                          []string                                  `pulumi:"groupIds"`
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
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

func (o PrivateEndpointConnectionResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
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

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
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

type SensorIntegration struct {
	Enabled *string `pulumi:"enabled"`
}





type SensorIntegrationInput interface {
	pulumi.Input

	ToSensorIntegrationOutput() SensorIntegrationOutput
	ToSensorIntegrationOutputWithContext(context.Context) SensorIntegrationOutput
}

type SensorIntegrationArgs struct {
	Enabled pulumi.StringPtrInput `pulumi:"enabled"`
}

func (SensorIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SensorIntegration)(nil)).Elem()
}

func (i SensorIntegrationArgs) ToSensorIntegrationOutput() SensorIntegrationOutput {
	return i.ToSensorIntegrationOutputWithContext(context.Background())
}

func (i SensorIntegrationArgs) ToSensorIntegrationOutputWithContext(ctx context.Context) SensorIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensorIntegrationOutput)
}

func (i SensorIntegrationArgs) ToSensorIntegrationPtrOutput() SensorIntegrationPtrOutput {
	return i.ToSensorIntegrationPtrOutputWithContext(context.Background())
}

func (i SensorIntegrationArgs) ToSensorIntegrationPtrOutputWithContext(ctx context.Context) SensorIntegrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensorIntegrationOutput).ToSensorIntegrationPtrOutputWithContext(ctx)
}









type SensorIntegrationPtrInput interface {
	pulumi.Input

	ToSensorIntegrationPtrOutput() SensorIntegrationPtrOutput
	ToSensorIntegrationPtrOutputWithContext(context.Context) SensorIntegrationPtrOutput
}

type sensorIntegrationPtrType SensorIntegrationArgs

func SensorIntegrationPtr(v *SensorIntegrationArgs) SensorIntegrationPtrInput {
	return (*sensorIntegrationPtrType)(v)
}

func (*sensorIntegrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SensorIntegration)(nil)).Elem()
}

func (i *sensorIntegrationPtrType) ToSensorIntegrationPtrOutput() SensorIntegrationPtrOutput {
	return i.ToSensorIntegrationPtrOutputWithContext(context.Background())
}

func (i *sensorIntegrationPtrType) ToSensorIntegrationPtrOutputWithContext(ctx context.Context) SensorIntegrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensorIntegrationPtrOutput)
}

type SensorIntegrationOutput struct{ *pulumi.OutputState }

func (SensorIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SensorIntegration)(nil)).Elem()
}

func (o SensorIntegrationOutput) ToSensorIntegrationOutput() SensorIntegrationOutput {
	return o
}

func (o SensorIntegrationOutput) ToSensorIntegrationOutputWithContext(ctx context.Context) SensorIntegrationOutput {
	return o
}

func (o SensorIntegrationOutput) ToSensorIntegrationPtrOutput() SensorIntegrationPtrOutput {
	return o.ToSensorIntegrationPtrOutputWithContext(context.Background())
}

func (o SensorIntegrationOutput) ToSensorIntegrationPtrOutputWithContext(ctx context.Context) SensorIntegrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SensorIntegration) *SensorIntegration {
		return &v
	}).(SensorIntegrationPtrOutput)
}

func (o SensorIntegrationOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SensorIntegration) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

type SensorIntegrationPtrOutput struct{ *pulumi.OutputState }

func (SensorIntegrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SensorIntegration)(nil)).Elem()
}

func (o SensorIntegrationPtrOutput) ToSensorIntegrationPtrOutput() SensorIntegrationPtrOutput {
	return o
}

func (o SensorIntegrationPtrOutput) ToSensorIntegrationPtrOutputWithContext(ctx context.Context) SensorIntegrationPtrOutput {
	return o
}

func (o SensorIntegrationPtrOutput) Elem() SensorIntegrationOutput {
	return o.ApplyT(func(v *SensorIntegration) SensorIntegration {
		if v != nil {
			return *v
		}
		var ret SensorIntegration
		return ret
	}).(SensorIntegrationOutput)
}

func (o SensorIntegrationPtrOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensorIntegration) *string {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.StringPtrOutput)
}

type SensorIntegrationResponse struct {
	Enabled           *string                `pulumi:"enabled"`
	ProvisioningInfo  *ErrorResponseResponse `pulumi:"provisioningInfo"`
	ProvisioningState string                 `pulumi:"provisioningState"`
}

type SensorIntegrationResponseOutput struct{ *pulumi.OutputState }

func (SensorIntegrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SensorIntegrationResponse)(nil)).Elem()
}

func (o SensorIntegrationResponseOutput) ToSensorIntegrationResponseOutput() SensorIntegrationResponseOutput {
	return o
}

func (o SensorIntegrationResponseOutput) ToSensorIntegrationResponseOutputWithContext(ctx context.Context) SensorIntegrationResponseOutput {
	return o
}

func (o SensorIntegrationResponseOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SensorIntegrationResponse) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o SensorIntegrationResponseOutput) ProvisioningInfo() ErrorResponseResponsePtrOutput {
	return o.ApplyT(func(v SensorIntegrationResponse) *ErrorResponseResponse { return v.ProvisioningInfo }).(ErrorResponseResponsePtrOutput)
}

func (o SensorIntegrationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SensorIntegrationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type SensorIntegrationResponsePtrOutput struct{ *pulumi.OutputState }

func (SensorIntegrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SensorIntegrationResponse)(nil)).Elem()
}

func (o SensorIntegrationResponsePtrOutput) ToSensorIntegrationResponsePtrOutput() SensorIntegrationResponsePtrOutput {
	return o
}

func (o SensorIntegrationResponsePtrOutput) ToSensorIntegrationResponsePtrOutputWithContext(ctx context.Context) SensorIntegrationResponsePtrOutput {
	return o
}

func (o SensorIntegrationResponsePtrOutput) Elem() SensorIntegrationResponseOutput {
	return o.ApplyT(func(v *SensorIntegrationResponse) SensorIntegrationResponse {
		if v != nil {
			return *v
		}
		var ret SensorIntegrationResponse
		return ret
	}).(SensorIntegrationResponseOutput)
}

func (o SensorIntegrationResponsePtrOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensorIntegrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.StringPtrOutput)
}

func (o SensorIntegrationResponsePtrOutput) ProvisioningInfo() ErrorResponseResponsePtrOutput {
	return o.ApplyT(func(v *SensorIntegrationResponse) *ErrorResponseResponse {
		if v == nil {
			return nil
		}
		return v.ProvisioningInfo
	}).(ErrorResponseResponsePtrOutput)
}

func (o SensorIntegrationResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensorIntegrationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(ApiPropertiesOutput{})
	pulumi.RegisterOutputType(ApiPropertiesMapOutput{})
	pulumi.RegisterOutputType(ApiPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApiPropertiesResponseMapOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(SensorIntegrationOutput{})
	pulumi.RegisterOutputType(SensorIntegrationPtrOutput{})
	pulumi.RegisterOutputType(SensorIntegrationResponseOutput{})
	pulumi.RegisterOutputType(SensorIntegrationResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

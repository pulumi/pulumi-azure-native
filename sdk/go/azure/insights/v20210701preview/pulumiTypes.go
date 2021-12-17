


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessModeSettings struct {
	Exclusions          []AccessModeSettingsExclusion `pulumi:"exclusions"`
	IngestionAccessMode string                        `pulumi:"ingestionAccessMode"`
	QueryAccessMode     string                        `pulumi:"queryAccessMode"`
}





type AccessModeSettingsInput interface {
	pulumi.Input

	ToAccessModeSettingsOutput() AccessModeSettingsOutput
	ToAccessModeSettingsOutputWithContext(context.Context) AccessModeSettingsOutput
}

type AccessModeSettingsArgs struct {
	Exclusions          AccessModeSettingsExclusionArrayInput `pulumi:"exclusions"`
	IngestionAccessMode pulumi.StringInput                    `pulumi:"ingestionAccessMode"`
	QueryAccessMode     pulumi.StringInput                    `pulumi:"queryAccessMode"`
}

func (AccessModeSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettings)(nil)).Elem()
}

func (i AccessModeSettingsArgs) ToAccessModeSettingsOutput() AccessModeSettingsOutput {
	return i.ToAccessModeSettingsOutputWithContext(context.Background())
}

func (i AccessModeSettingsArgs) ToAccessModeSettingsOutputWithContext(ctx context.Context) AccessModeSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsOutput)
}

type AccessModeSettingsOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettings)(nil)).Elem()
}

func (o AccessModeSettingsOutput) ToAccessModeSettingsOutput() AccessModeSettingsOutput {
	return o
}

func (o AccessModeSettingsOutput) ToAccessModeSettingsOutputWithContext(ctx context.Context) AccessModeSettingsOutput {
	return o
}

func (o AccessModeSettingsOutput) Exclusions() AccessModeSettingsExclusionArrayOutput {
	return o.ApplyT(func(v AccessModeSettings) []AccessModeSettingsExclusion { return v.Exclusions }).(AccessModeSettingsExclusionArrayOutput)
}

func (o AccessModeSettingsOutput) IngestionAccessMode() pulumi.StringOutput {
	return o.ApplyT(func(v AccessModeSettings) string { return v.IngestionAccessMode }).(pulumi.StringOutput)
}

func (o AccessModeSettingsOutput) QueryAccessMode() pulumi.StringOutput {
	return o.ApplyT(func(v AccessModeSettings) string { return v.QueryAccessMode }).(pulumi.StringOutput)
}

type AccessModeSettingsExclusion struct {
	IngestionAccessMode           *string `pulumi:"ingestionAccessMode"`
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	QueryAccessMode               *string `pulumi:"queryAccessMode"`
}





type AccessModeSettingsExclusionInput interface {
	pulumi.Input

	ToAccessModeSettingsExclusionOutput() AccessModeSettingsExclusionOutput
	ToAccessModeSettingsExclusionOutputWithContext(context.Context) AccessModeSettingsExclusionOutput
}

type AccessModeSettingsExclusionArgs struct {
	IngestionAccessMode           pulumi.StringPtrInput `pulumi:"ingestionAccessMode"`
	PrivateEndpointConnectionName pulumi.StringPtrInput `pulumi:"privateEndpointConnectionName"`
	QueryAccessMode               pulumi.StringPtrInput `pulumi:"queryAccessMode"`
}

func (AccessModeSettingsExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettingsExclusion)(nil)).Elem()
}

func (i AccessModeSettingsExclusionArgs) ToAccessModeSettingsExclusionOutput() AccessModeSettingsExclusionOutput {
	return i.ToAccessModeSettingsExclusionOutputWithContext(context.Background())
}

func (i AccessModeSettingsExclusionArgs) ToAccessModeSettingsExclusionOutputWithContext(ctx context.Context) AccessModeSettingsExclusionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsExclusionOutput)
}





type AccessModeSettingsExclusionArrayInput interface {
	pulumi.Input

	ToAccessModeSettingsExclusionArrayOutput() AccessModeSettingsExclusionArrayOutput
	ToAccessModeSettingsExclusionArrayOutputWithContext(context.Context) AccessModeSettingsExclusionArrayOutput
}

type AccessModeSettingsExclusionArray []AccessModeSettingsExclusionInput

func (AccessModeSettingsExclusionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessModeSettingsExclusion)(nil)).Elem()
}

func (i AccessModeSettingsExclusionArray) ToAccessModeSettingsExclusionArrayOutput() AccessModeSettingsExclusionArrayOutput {
	return i.ToAccessModeSettingsExclusionArrayOutputWithContext(context.Background())
}

func (i AccessModeSettingsExclusionArray) ToAccessModeSettingsExclusionArrayOutputWithContext(ctx context.Context) AccessModeSettingsExclusionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsExclusionArrayOutput)
}

type AccessModeSettingsExclusionOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsExclusionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettingsExclusion)(nil)).Elem()
}

func (o AccessModeSettingsExclusionOutput) ToAccessModeSettingsExclusionOutput() AccessModeSettingsExclusionOutput {
	return o
}

func (o AccessModeSettingsExclusionOutput) ToAccessModeSettingsExclusionOutputWithContext(ctx context.Context) AccessModeSettingsExclusionOutput {
	return o
}

func (o AccessModeSettingsExclusionOutput) IngestionAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessModeSettingsExclusion) *string { return v.IngestionAccessMode }).(pulumi.StringPtrOutput)
}

func (o AccessModeSettingsExclusionOutput) PrivateEndpointConnectionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessModeSettingsExclusion) *string { return v.PrivateEndpointConnectionName }).(pulumi.StringPtrOutput)
}

func (o AccessModeSettingsExclusionOutput) QueryAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessModeSettingsExclusion) *string { return v.QueryAccessMode }).(pulumi.StringPtrOutput)
}

type AccessModeSettingsExclusionArrayOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsExclusionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessModeSettingsExclusion)(nil)).Elem()
}

func (o AccessModeSettingsExclusionArrayOutput) ToAccessModeSettingsExclusionArrayOutput() AccessModeSettingsExclusionArrayOutput {
	return o
}

func (o AccessModeSettingsExclusionArrayOutput) ToAccessModeSettingsExclusionArrayOutputWithContext(ctx context.Context) AccessModeSettingsExclusionArrayOutput {
	return o
}

func (o AccessModeSettingsExclusionArrayOutput) Index(i pulumi.IntInput) AccessModeSettingsExclusionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessModeSettingsExclusion {
		return vs[0].([]AccessModeSettingsExclusion)[vs[1].(int)]
	}).(AccessModeSettingsExclusionOutput)
}

type AccessModeSettingsExclusionResponse struct {
	IngestionAccessMode           *string `pulumi:"ingestionAccessMode"`
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	QueryAccessMode               *string `pulumi:"queryAccessMode"`
}

type AccessModeSettingsExclusionResponseOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsExclusionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettingsExclusionResponse)(nil)).Elem()
}

func (o AccessModeSettingsExclusionResponseOutput) ToAccessModeSettingsExclusionResponseOutput() AccessModeSettingsExclusionResponseOutput {
	return o
}

func (o AccessModeSettingsExclusionResponseOutput) ToAccessModeSettingsExclusionResponseOutputWithContext(ctx context.Context) AccessModeSettingsExclusionResponseOutput {
	return o
}

func (o AccessModeSettingsExclusionResponseOutput) IngestionAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessModeSettingsExclusionResponse) *string { return v.IngestionAccessMode }).(pulumi.StringPtrOutput)
}

func (o AccessModeSettingsExclusionResponseOutput) PrivateEndpointConnectionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessModeSettingsExclusionResponse) *string { return v.PrivateEndpointConnectionName }).(pulumi.StringPtrOutput)
}

func (o AccessModeSettingsExclusionResponseOutput) QueryAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessModeSettingsExclusionResponse) *string { return v.QueryAccessMode }).(pulumi.StringPtrOutput)
}

type AccessModeSettingsExclusionResponseArrayOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsExclusionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessModeSettingsExclusionResponse)(nil)).Elem()
}

func (o AccessModeSettingsExclusionResponseArrayOutput) ToAccessModeSettingsExclusionResponseArrayOutput() AccessModeSettingsExclusionResponseArrayOutput {
	return o
}

func (o AccessModeSettingsExclusionResponseArrayOutput) ToAccessModeSettingsExclusionResponseArrayOutputWithContext(ctx context.Context) AccessModeSettingsExclusionResponseArrayOutput {
	return o
}

func (o AccessModeSettingsExclusionResponseArrayOutput) Index(i pulumi.IntInput) AccessModeSettingsExclusionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessModeSettingsExclusionResponse {
		return vs[0].([]AccessModeSettingsExclusionResponse)[vs[1].(int)]
	}).(AccessModeSettingsExclusionResponseOutput)
}

type AccessModeSettingsResponse struct {
	Exclusions          []AccessModeSettingsExclusionResponse `pulumi:"exclusions"`
	IngestionAccessMode string                                `pulumi:"ingestionAccessMode"`
	QueryAccessMode     string                                `pulumi:"queryAccessMode"`
}

type AccessModeSettingsResponseOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettingsResponse)(nil)).Elem()
}

func (o AccessModeSettingsResponseOutput) ToAccessModeSettingsResponseOutput() AccessModeSettingsResponseOutput {
	return o
}

func (o AccessModeSettingsResponseOutput) ToAccessModeSettingsResponseOutputWithContext(ctx context.Context) AccessModeSettingsResponseOutput {
	return o
}

func (o AccessModeSettingsResponseOutput) Exclusions() AccessModeSettingsExclusionResponseArrayOutput {
	return o.ApplyT(func(v AccessModeSettingsResponse) []AccessModeSettingsExclusionResponse { return v.Exclusions }).(AccessModeSettingsExclusionResponseArrayOutput)
}

func (o AccessModeSettingsResponseOutput) IngestionAccessMode() pulumi.StringOutput {
	return o.ApplyT(func(v AccessModeSettingsResponse) string { return v.IngestionAccessMode }).(pulumi.StringOutput)
}

func (o AccessModeSettingsResponseOutput) QueryAccessMode() pulumi.StringOutput {
	return o.ApplyT(func(v AccessModeSettingsResponse) string { return v.QueryAccessMode }).(pulumi.StringOutput)
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
	pulumi.RegisterOutputType(AccessModeSettingsOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsExclusionOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsExclusionArrayOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsExclusionResponseOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsExclusionResponseArrayOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

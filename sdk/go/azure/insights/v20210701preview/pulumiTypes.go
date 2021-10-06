


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

func (i AccessModeSettingsArgs) ToAccessModeSettingsPtrOutput() AccessModeSettingsPtrOutput {
	return i.ToAccessModeSettingsPtrOutputWithContext(context.Background())
}

func (i AccessModeSettingsArgs) ToAccessModeSettingsPtrOutputWithContext(ctx context.Context) AccessModeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsOutput).ToAccessModeSettingsPtrOutputWithContext(ctx)
}









type AccessModeSettingsPtrInput interface {
	pulumi.Input

	ToAccessModeSettingsPtrOutput() AccessModeSettingsPtrOutput
	ToAccessModeSettingsPtrOutputWithContext(context.Context) AccessModeSettingsPtrOutput
}

type accessModeSettingsPtrType AccessModeSettingsArgs

func AccessModeSettingsPtr(v *AccessModeSettingsArgs) AccessModeSettingsPtrInput {
	return (*accessModeSettingsPtrType)(v)
}

func (*accessModeSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessModeSettings)(nil)).Elem()
}

func (i *accessModeSettingsPtrType) ToAccessModeSettingsPtrOutput() AccessModeSettingsPtrOutput {
	return i.ToAccessModeSettingsPtrOutputWithContext(context.Background())
}

func (i *accessModeSettingsPtrType) ToAccessModeSettingsPtrOutputWithContext(ctx context.Context) AccessModeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsPtrOutput)
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

func (o AccessModeSettingsOutput) ToAccessModeSettingsPtrOutput() AccessModeSettingsPtrOutput {
	return o.ToAccessModeSettingsPtrOutputWithContext(context.Background())
}

func (o AccessModeSettingsOutput) ToAccessModeSettingsPtrOutputWithContext(ctx context.Context) AccessModeSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessModeSettings) *AccessModeSettings {
		return &v
	}).(AccessModeSettingsPtrOutput)
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

type AccessModeSettingsPtrOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessModeSettings)(nil)).Elem()
}

func (o AccessModeSettingsPtrOutput) ToAccessModeSettingsPtrOutput() AccessModeSettingsPtrOutput {
	return o
}

func (o AccessModeSettingsPtrOutput) ToAccessModeSettingsPtrOutputWithContext(ctx context.Context) AccessModeSettingsPtrOutput {
	return o
}

func (o AccessModeSettingsPtrOutput) Elem() AccessModeSettingsOutput {
	return o.ApplyT(func(v *AccessModeSettings) AccessModeSettings {
		if v != nil {
			return *v
		}
		var ret AccessModeSettings
		return ret
	}).(AccessModeSettingsOutput)
}

func (o AccessModeSettingsPtrOutput) Exclusions() AccessModeSettingsExclusionArrayOutput {
	return o.ApplyT(func(v *AccessModeSettings) []AccessModeSettingsExclusion {
		if v == nil {
			return nil
		}
		return v.Exclusions
	}).(AccessModeSettingsExclusionArrayOutput)
}

func (o AccessModeSettingsPtrOutput) IngestionAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessModeSettings) *string {
		if v == nil {
			return nil
		}
		return &v.IngestionAccessMode
	}).(pulumi.StringPtrOutput)
}

func (o AccessModeSettingsPtrOutput) QueryAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessModeSettings) *string {
		if v == nil {
			return nil
		}
		return &v.QueryAccessMode
	}).(pulumi.StringPtrOutput)
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





type AccessModeSettingsExclusionResponseInput interface {
	pulumi.Input

	ToAccessModeSettingsExclusionResponseOutput() AccessModeSettingsExclusionResponseOutput
	ToAccessModeSettingsExclusionResponseOutputWithContext(context.Context) AccessModeSettingsExclusionResponseOutput
}

type AccessModeSettingsExclusionResponseArgs struct {
	IngestionAccessMode           pulumi.StringPtrInput `pulumi:"ingestionAccessMode"`
	PrivateEndpointConnectionName pulumi.StringPtrInput `pulumi:"privateEndpointConnectionName"`
	QueryAccessMode               pulumi.StringPtrInput `pulumi:"queryAccessMode"`
}

func (AccessModeSettingsExclusionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettingsExclusionResponse)(nil)).Elem()
}

func (i AccessModeSettingsExclusionResponseArgs) ToAccessModeSettingsExclusionResponseOutput() AccessModeSettingsExclusionResponseOutput {
	return i.ToAccessModeSettingsExclusionResponseOutputWithContext(context.Background())
}

func (i AccessModeSettingsExclusionResponseArgs) ToAccessModeSettingsExclusionResponseOutputWithContext(ctx context.Context) AccessModeSettingsExclusionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsExclusionResponseOutput)
}





type AccessModeSettingsExclusionResponseArrayInput interface {
	pulumi.Input

	ToAccessModeSettingsExclusionResponseArrayOutput() AccessModeSettingsExclusionResponseArrayOutput
	ToAccessModeSettingsExclusionResponseArrayOutputWithContext(context.Context) AccessModeSettingsExclusionResponseArrayOutput
}

type AccessModeSettingsExclusionResponseArray []AccessModeSettingsExclusionResponseInput

func (AccessModeSettingsExclusionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessModeSettingsExclusionResponse)(nil)).Elem()
}

func (i AccessModeSettingsExclusionResponseArray) ToAccessModeSettingsExclusionResponseArrayOutput() AccessModeSettingsExclusionResponseArrayOutput {
	return i.ToAccessModeSettingsExclusionResponseArrayOutputWithContext(context.Background())
}

func (i AccessModeSettingsExclusionResponseArray) ToAccessModeSettingsExclusionResponseArrayOutputWithContext(ctx context.Context) AccessModeSettingsExclusionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsExclusionResponseArrayOutput)
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





type AccessModeSettingsResponseInput interface {
	pulumi.Input

	ToAccessModeSettingsResponseOutput() AccessModeSettingsResponseOutput
	ToAccessModeSettingsResponseOutputWithContext(context.Context) AccessModeSettingsResponseOutput
}

type AccessModeSettingsResponseArgs struct {
	Exclusions          AccessModeSettingsExclusionResponseArrayInput `pulumi:"exclusions"`
	IngestionAccessMode pulumi.StringInput                            `pulumi:"ingestionAccessMode"`
	QueryAccessMode     pulumi.StringInput                            `pulumi:"queryAccessMode"`
}

func (AccessModeSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettingsResponse)(nil)).Elem()
}

func (i AccessModeSettingsResponseArgs) ToAccessModeSettingsResponseOutput() AccessModeSettingsResponseOutput {
	return i.ToAccessModeSettingsResponseOutputWithContext(context.Background())
}

func (i AccessModeSettingsResponseArgs) ToAccessModeSettingsResponseOutputWithContext(ctx context.Context) AccessModeSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsResponseOutput)
}

func (i AccessModeSettingsResponseArgs) ToAccessModeSettingsResponsePtrOutput() AccessModeSettingsResponsePtrOutput {
	return i.ToAccessModeSettingsResponsePtrOutputWithContext(context.Background())
}

func (i AccessModeSettingsResponseArgs) ToAccessModeSettingsResponsePtrOutputWithContext(ctx context.Context) AccessModeSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsResponseOutput).ToAccessModeSettingsResponsePtrOutputWithContext(ctx)
}









type AccessModeSettingsResponsePtrInput interface {
	pulumi.Input

	ToAccessModeSettingsResponsePtrOutput() AccessModeSettingsResponsePtrOutput
	ToAccessModeSettingsResponsePtrOutputWithContext(context.Context) AccessModeSettingsResponsePtrOutput
}

type accessModeSettingsResponsePtrType AccessModeSettingsResponseArgs

func AccessModeSettingsResponsePtr(v *AccessModeSettingsResponseArgs) AccessModeSettingsResponsePtrInput {
	return (*accessModeSettingsResponsePtrType)(v)
}

func (*accessModeSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessModeSettingsResponse)(nil)).Elem()
}

func (i *accessModeSettingsResponsePtrType) ToAccessModeSettingsResponsePtrOutput() AccessModeSettingsResponsePtrOutput {
	return i.ToAccessModeSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *accessModeSettingsResponsePtrType) ToAccessModeSettingsResponsePtrOutputWithContext(ctx context.Context) AccessModeSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsResponsePtrOutput)
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

func (o AccessModeSettingsResponseOutput) ToAccessModeSettingsResponsePtrOutput() AccessModeSettingsResponsePtrOutput {
	return o.ToAccessModeSettingsResponsePtrOutputWithContext(context.Background())
}

func (o AccessModeSettingsResponseOutput) ToAccessModeSettingsResponsePtrOutputWithContext(ctx context.Context) AccessModeSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessModeSettingsResponse) *AccessModeSettingsResponse {
		return &v
	}).(AccessModeSettingsResponsePtrOutput)
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

type AccessModeSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessModeSettingsResponse)(nil)).Elem()
}

func (o AccessModeSettingsResponsePtrOutput) ToAccessModeSettingsResponsePtrOutput() AccessModeSettingsResponsePtrOutput {
	return o
}

func (o AccessModeSettingsResponsePtrOutput) ToAccessModeSettingsResponsePtrOutputWithContext(ctx context.Context) AccessModeSettingsResponsePtrOutput {
	return o
}

func (o AccessModeSettingsResponsePtrOutput) Elem() AccessModeSettingsResponseOutput {
	return o.ApplyT(func(v *AccessModeSettingsResponse) AccessModeSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AccessModeSettingsResponse
		return ret
	}).(AccessModeSettingsResponseOutput)
}

func (o AccessModeSettingsResponsePtrOutput) Exclusions() AccessModeSettingsExclusionResponseArrayOutput {
	return o.ApplyT(func(v *AccessModeSettingsResponse) []AccessModeSettingsExclusionResponse {
		if v == nil {
			return nil
		}
		return v.Exclusions
	}).(AccessModeSettingsExclusionResponseArrayOutput)
}

func (o AccessModeSettingsResponsePtrOutput) IngestionAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessModeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.IngestionAccessMode
	}).(pulumi.StringPtrOutput)
}

func (o AccessModeSettingsResponsePtrOutput) QueryAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessModeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.QueryAccessMode
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id                                pulumi.StringInput                             `pulumi:"id"`
	Name                              pulumi.StringInput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                             `pulumi:"provisioningState"`
	Type                              pulumi.StringInput                             `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
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





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
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

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
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





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
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

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
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

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessModeSettingsOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsPtrOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsExclusionOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsExclusionArrayOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsExclusionResponseOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsExclusionResponseArrayOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsResponseOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}

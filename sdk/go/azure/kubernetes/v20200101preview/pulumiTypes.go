


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthenticationDetailsValue struct {
	Token *string `pulumi:"token"`
}





type AuthenticationDetailsValueInput interface {
	pulumi.Input

	ToAuthenticationDetailsValueOutput() AuthenticationDetailsValueOutput
	ToAuthenticationDetailsValueOutputWithContext(context.Context) AuthenticationDetailsValueOutput
}

type AuthenticationDetailsValueArgs struct {
	Token pulumi.StringPtrInput `pulumi:"token"`
}

func (AuthenticationDetailsValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationDetailsValue)(nil)).Elem()
}

func (i AuthenticationDetailsValueArgs) ToAuthenticationDetailsValueOutput() AuthenticationDetailsValueOutput {
	return i.ToAuthenticationDetailsValueOutputWithContext(context.Background())
}

func (i AuthenticationDetailsValueArgs) ToAuthenticationDetailsValueOutputWithContext(ctx context.Context) AuthenticationDetailsValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationDetailsValueOutput)
}

type AuthenticationDetailsValueOutput struct{ *pulumi.OutputState }

func (AuthenticationDetailsValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationDetailsValue)(nil)).Elem()
}

func (o AuthenticationDetailsValueOutput) ToAuthenticationDetailsValueOutput() AuthenticationDetailsValueOutput {
	return o
}

func (o AuthenticationDetailsValueOutput) ToAuthenticationDetailsValueOutputWithContext(ctx context.Context) AuthenticationDetailsValueOutput {
	return o
}

func (o AuthenticationDetailsValueOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthenticationDetailsValue) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type ConnectedClusterAADProfile struct {
	ClientAppId string `pulumi:"clientAppId"`
	ServerAppId string `pulumi:"serverAppId"`
	TenantId    string `pulumi:"tenantId"`
}





type ConnectedClusterAADProfileInput interface {
	pulumi.Input

	ToConnectedClusterAADProfileOutput() ConnectedClusterAADProfileOutput
	ToConnectedClusterAADProfileOutputWithContext(context.Context) ConnectedClusterAADProfileOutput
}

type ConnectedClusterAADProfileArgs struct {
	ClientAppId pulumi.StringInput `pulumi:"clientAppId"`
	ServerAppId pulumi.StringInput `pulumi:"serverAppId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
}

func (ConnectedClusterAADProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterAADProfile)(nil)).Elem()
}

func (i ConnectedClusterAADProfileArgs) ToConnectedClusterAADProfileOutput() ConnectedClusterAADProfileOutput {
	return i.ToConnectedClusterAADProfileOutputWithContext(context.Background())
}

func (i ConnectedClusterAADProfileArgs) ToConnectedClusterAADProfileOutputWithContext(ctx context.Context) ConnectedClusterAADProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterAADProfileOutput)
}

func (i ConnectedClusterAADProfileArgs) ToConnectedClusterAADProfilePtrOutput() ConnectedClusterAADProfilePtrOutput {
	return i.ToConnectedClusterAADProfilePtrOutputWithContext(context.Background())
}

func (i ConnectedClusterAADProfileArgs) ToConnectedClusterAADProfilePtrOutputWithContext(ctx context.Context) ConnectedClusterAADProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterAADProfileOutput).ToConnectedClusterAADProfilePtrOutputWithContext(ctx)
}









type ConnectedClusterAADProfilePtrInput interface {
	pulumi.Input

	ToConnectedClusterAADProfilePtrOutput() ConnectedClusterAADProfilePtrOutput
	ToConnectedClusterAADProfilePtrOutputWithContext(context.Context) ConnectedClusterAADProfilePtrOutput
}

type connectedClusterAADProfilePtrType ConnectedClusterAADProfileArgs

func ConnectedClusterAADProfilePtr(v *ConnectedClusterAADProfileArgs) ConnectedClusterAADProfilePtrInput {
	return (*connectedClusterAADProfilePtrType)(v)
}

func (*connectedClusterAADProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterAADProfile)(nil)).Elem()
}

func (i *connectedClusterAADProfilePtrType) ToConnectedClusterAADProfilePtrOutput() ConnectedClusterAADProfilePtrOutput {
	return i.ToConnectedClusterAADProfilePtrOutputWithContext(context.Background())
}

func (i *connectedClusterAADProfilePtrType) ToConnectedClusterAADProfilePtrOutputWithContext(ctx context.Context) ConnectedClusterAADProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterAADProfilePtrOutput)
}

type ConnectedClusterAADProfileOutput struct{ *pulumi.OutputState }

func (ConnectedClusterAADProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterAADProfile)(nil)).Elem()
}

func (o ConnectedClusterAADProfileOutput) ToConnectedClusterAADProfileOutput() ConnectedClusterAADProfileOutput {
	return o
}

func (o ConnectedClusterAADProfileOutput) ToConnectedClusterAADProfileOutputWithContext(ctx context.Context) ConnectedClusterAADProfileOutput {
	return o
}

func (o ConnectedClusterAADProfileOutput) ToConnectedClusterAADProfilePtrOutput() ConnectedClusterAADProfilePtrOutput {
	return o.ToConnectedClusterAADProfilePtrOutputWithContext(context.Background())
}

func (o ConnectedClusterAADProfileOutput) ToConnectedClusterAADProfilePtrOutputWithContext(ctx context.Context) ConnectedClusterAADProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectedClusterAADProfile) *ConnectedClusterAADProfile {
		return &v
	}).(ConnectedClusterAADProfilePtrOutput)
}

func (o ConnectedClusterAADProfileOutput) ClientAppId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterAADProfile) string { return v.ClientAppId }).(pulumi.StringOutput)
}

func (o ConnectedClusterAADProfileOutput) ServerAppId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterAADProfile) string { return v.ServerAppId }).(pulumi.StringOutput)
}

func (o ConnectedClusterAADProfileOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterAADProfile) string { return v.TenantId }).(pulumi.StringOutput)
}

type ConnectedClusterAADProfilePtrOutput struct{ *pulumi.OutputState }

func (ConnectedClusterAADProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterAADProfile)(nil)).Elem()
}

func (o ConnectedClusterAADProfilePtrOutput) ToConnectedClusterAADProfilePtrOutput() ConnectedClusterAADProfilePtrOutput {
	return o
}

func (o ConnectedClusterAADProfilePtrOutput) ToConnectedClusterAADProfilePtrOutputWithContext(ctx context.Context) ConnectedClusterAADProfilePtrOutput {
	return o
}

func (o ConnectedClusterAADProfilePtrOutput) Elem() ConnectedClusterAADProfileOutput {
	return o.ApplyT(func(v *ConnectedClusterAADProfile) ConnectedClusterAADProfile {
		if v != nil {
			return *v
		}
		var ret ConnectedClusterAADProfile
		return ret
	}).(ConnectedClusterAADProfileOutput)
}

func (o ConnectedClusterAADProfilePtrOutput) ClientAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ClientAppId
	}).(pulumi.StringPtrOutput)
}

func (o ConnectedClusterAADProfilePtrOutput) ServerAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ServerAppId
	}).(pulumi.StringPtrOutput)
}

func (o ConnectedClusterAADProfilePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

type ConnectedClusterAADProfileResponse struct {
	ClientAppId string `pulumi:"clientAppId"`
	ServerAppId string `pulumi:"serverAppId"`
	TenantId    string `pulumi:"tenantId"`
}





type ConnectedClusterAADProfileResponseInput interface {
	pulumi.Input

	ToConnectedClusterAADProfileResponseOutput() ConnectedClusterAADProfileResponseOutput
	ToConnectedClusterAADProfileResponseOutputWithContext(context.Context) ConnectedClusterAADProfileResponseOutput
}

type ConnectedClusterAADProfileResponseArgs struct {
	ClientAppId pulumi.StringInput `pulumi:"clientAppId"`
	ServerAppId pulumi.StringInput `pulumi:"serverAppId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
}

func (ConnectedClusterAADProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterAADProfileResponse)(nil)).Elem()
}

func (i ConnectedClusterAADProfileResponseArgs) ToConnectedClusterAADProfileResponseOutput() ConnectedClusterAADProfileResponseOutput {
	return i.ToConnectedClusterAADProfileResponseOutputWithContext(context.Background())
}

func (i ConnectedClusterAADProfileResponseArgs) ToConnectedClusterAADProfileResponseOutputWithContext(ctx context.Context) ConnectedClusterAADProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterAADProfileResponseOutput)
}

func (i ConnectedClusterAADProfileResponseArgs) ToConnectedClusterAADProfileResponsePtrOutput() ConnectedClusterAADProfileResponsePtrOutput {
	return i.ToConnectedClusterAADProfileResponsePtrOutputWithContext(context.Background())
}

func (i ConnectedClusterAADProfileResponseArgs) ToConnectedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterAADProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterAADProfileResponseOutput).ToConnectedClusterAADProfileResponsePtrOutputWithContext(ctx)
}









type ConnectedClusterAADProfileResponsePtrInput interface {
	pulumi.Input

	ToConnectedClusterAADProfileResponsePtrOutput() ConnectedClusterAADProfileResponsePtrOutput
	ToConnectedClusterAADProfileResponsePtrOutputWithContext(context.Context) ConnectedClusterAADProfileResponsePtrOutput
}

type connectedClusterAADProfileResponsePtrType ConnectedClusterAADProfileResponseArgs

func ConnectedClusterAADProfileResponsePtr(v *ConnectedClusterAADProfileResponseArgs) ConnectedClusterAADProfileResponsePtrInput {
	return (*connectedClusterAADProfileResponsePtrType)(v)
}

func (*connectedClusterAADProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterAADProfileResponse)(nil)).Elem()
}

func (i *connectedClusterAADProfileResponsePtrType) ToConnectedClusterAADProfileResponsePtrOutput() ConnectedClusterAADProfileResponsePtrOutput {
	return i.ToConnectedClusterAADProfileResponsePtrOutputWithContext(context.Background())
}

func (i *connectedClusterAADProfileResponsePtrType) ToConnectedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterAADProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterAADProfileResponsePtrOutput)
}

type ConnectedClusterAADProfileResponseOutput struct{ *pulumi.OutputState }

func (ConnectedClusterAADProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterAADProfileResponse)(nil)).Elem()
}

func (o ConnectedClusterAADProfileResponseOutput) ToConnectedClusterAADProfileResponseOutput() ConnectedClusterAADProfileResponseOutput {
	return o
}

func (o ConnectedClusterAADProfileResponseOutput) ToConnectedClusterAADProfileResponseOutputWithContext(ctx context.Context) ConnectedClusterAADProfileResponseOutput {
	return o
}

func (o ConnectedClusterAADProfileResponseOutput) ToConnectedClusterAADProfileResponsePtrOutput() ConnectedClusterAADProfileResponsePtrOutput {
	return o.ToConnectedClusterAADProfileResponsePtrOutputWithContext(context.Background())
}

func (o ConnectedClusterAADProfileResponseOutput) ToConnectedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterAADProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectedClusterAADProfileResponse) *ConnectedClusterAADProfileResponse {
		return &v
	}).(ConnectedClusterAADProfileResponsePtrOutput)
}

func (o ConnectedClusterAADProfileResponseOutput) ClientAppId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterAADProfileResponse) string { return v.ClientAppId }).(pulumi.StringOutput)
}

func (o ConnectedClusterAADProfileResponseOutput) ServerAppId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterAADProfileResponse) string { return v.ServerAppId }).(pulumi.StringOutput)
}

func (o ConnectedClusterAADProfileResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterAADProfileResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type ConnectedClusterAADProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectedClusterAADProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterAADProfileResponse)(nil)).Elem()
}

func (o ConnectedClusterAADProfileResponsePtrOutput) ToConnectedClusterAADProfileResponsePtrOutput() ConnectedClusterAADProfileResponsePtrOutput {
	return o
}

func (o ConnectedClusterAADProfileResponsePtrOutput) ToConnectedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterAADProfileResponsePtrOutput {
	return o
}

func (o ConnectedClusterAADProfileResponsePtrOutput) Elem() ConnectedClusterAADProfileResponseOutput {
	return o.ApplyT(func(v *ConnectedClusterAADProfileResponse) ConnectedClusterAADProfileResponse {
		if v != nil {
			return *v
		}
		var ret ConnectedClusterAADProfileResponse
		return ret
	}).(ConnectedClusterAADProfileResponseOutput)
}

func (o ConnectedClusterAADProfileResponsePtrOutput) ClientAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientAppId
	}).(pulumi.StringPtrOutput)
}

func (o ConnectedClusterAADProfileResponsePtrOutput) ServerAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServerAppId
	}).(pulumi.StringPtrOutput)
}

func (o ConnectedClusterAADProfileResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

type ConnectedClusterIdentity struct {
	Type ResourceIdentityType `pulumi:"type"`
}





type ConnectedClusterIdentityInput interface {
	pulumi.Input

	ToConnectedClusterIdentityOutput() ConnectedClusterIdentityOutput
	ToConnectedClusterIdentityOutputWithContext(context.Context) ConnectedClusterIdentityOutput
}

type ConnectedClusterIdentityArgs struct {
	Type ResourceIdentityTypeInput `pulumi:"type"`
}

func (ConnectedClusterIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterIdentity)(nil)).Elem()
}

func (i ConnectedClusterIdentityArgs) ToConnectedClusterIdentityOutput() ConnectedClusterIdentityOutput {
	return i.ToConnectedClusterIdentityOutputWithContext(context.Background())
}

func (i ConnectedClusterIdentityArgs) ToConnectedClusterIdentityOutputWithContext(ctx context.Context) ConnectedClusterIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityOutput)
}

func (i ConnectedClusterIdentityArgs) ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput {
	return i.ToConnectedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i ConnectedClusterIdentityArgs) ToConnectedClusterIdentityPtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityOutput).ToConnectedClusterIdentityPtrOutputWithContext(ctx)
}









type ConnectedClusterIdentityPtrInput interface {
	pulumi.Input

	ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput
	ToConnectedClusterIdentityPtrOutputWithContext(context.Context) ConnectedClusterIdentityPtrOutput
}

type connectedClusterIdentityPtrType ConnectedClusterIdentityArgs

func ConnectedClusterIdentityPtr(v *ConnectedClusterIdentityArgs) ConnectedClusterIdentityPtrInput {
	return (*connectedClusterIdentityPtrType)(v)
}

func (*connectedClusterIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterIdentity)(nil)).Elem()
}

func (i *connectedClusterIdentityPtrType) ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput {
	return i.ToConnectedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i *connectedClusterIdentityPtrType) ToConnectedClusterIdentityPtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityPtrOutput)
}

type ConnectedClusterIdentityOutput struct{ *pulumi.OutputState }

func (ConnectedClusterIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterIdentity)(nil)).Elem()
}

func (o ConnectedClusterIdentityOutput) ToConnectedClusterIdentityOutput() ConnectedClusterIdentityOutput {
	return o
}

func (o ConnectedClusterIdentityOutput) ToConnectedClusterIdentityOutputWithContext(ctx context.Context) ConnectedClusterIdentityOutput {
	return o
}

func (o ConnectedClusterIdentityOutput) ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput {
	return o.ToConnectedClusterIdentityPtrOutputWithContext(context.Background())
}

func (o ConnectedClusterIdentityOutput) ToConnectedClusterIdentityPtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectedClusterIdentity) *ConnectedClusterIdentity {
		return &v
	}).(ConnectedClusterIdentityPtrOutput)
}

func (o ConnectedClusterIdentityOutput) Type() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v ConnectedClusterIdentity) ResourceIdentityType { return v.Type }).(ResourceIdentityTypeOutput)
}

type ConnectedClusterIdentityPtrOutput struct{ *pulumi.OutputState }

func (ConnectedClusterIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterIdentity)(nil)).Elem()
}

func (o ConnectedClusterIdentityPtrOutput) ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput {
	return o
}

func (o ConnectedClusterIdentityPtrOutput) ToConnectedClusterIdentityPtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityPtrOutput {
	return o
}

func (o ConnectedClusterIdentityPtrOutput) Elem() ConnectedClusterIdentityOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentity) ConnectedClusterIdentity {
		if v != nil {
			return *v
		}
		var ret ConnectedClusterIdentity
		return ret
	}).(ConnectedClusterIdentityOutput)
}

func (o ConnectedClusterIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type ConnectedClusterIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}





type ConnectedClusterIdentityResponseInput interface {
	pulumi.Input

	ToConnectedClusterIdentityResponseOutput() ConnectedClusterIdentityResponseOutput
	ToConnectedClusterIdentityResponseOutputWithContext(context.Context) ConnectedClusterIdentityResponseOutput
}

type ConnectedClusterIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
	Type        pulumi.StringInput `pulumi:"type"`
}

func (ConnectedClusterIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterIdentityResponse)(nil)).Elem()
}

func (i ConnectedClusterIdentityResponseArgs) ToConnectedClusterIdentityResponseOutput() ConnectedClusterIdentityResponseOutput {
	return i.ToConnectedClusterIdentityResponseOutputWithContext(context.Background())
}

func (i ConnectedClusterIdentityResponseArgs) ToConnectedClusterIdentityResponseOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityResponseOutput)
}

func (i ConnectedClusterIdentityResponseArgs) ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput {
	return i.ToConnectedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ConnectedClusterIdentityResponseArgs) ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityResponseOutput).ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx)
}









type ConnectedClusterIdentityResponsePtrInput interface {
	pulumi.Input

	ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput
	ToConnectedClusterIdentityResponsePtrOutputWithContext(context.Context) ConnectedClusterIdentityResponsePtrOutput
}

type connectedClusterIdentityResponsePtrType ConnectedClusterIdentityResponseArgs

func ConnectedClusterIdentityResponsePtr(v *ConnectedClusterIdentityResponseArgs) ConnectedClusterIdentityResponsePtrInput {
	return (*connectedClusterIdentityResponsePtrType)(v)
}

func (*connectedClusterIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterIdentityResponse)(nil)).Elem()
}

func (i *connectedClusterIdentityResponsePtrType) ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput {
	return i.ToConnectedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *connectedClusterIdentityResponsePtrType) ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityResponsePtrOutput)
}

type ConnectedClusterIdentityResponseOutput struct{ *pulumi.OutputState }

func (ConnectedClusterIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterIdentityResponse)(nil)).Elem()
}

func (o ConnectedClusterIdentityResponseOutput) ToConnectedClusterIdentityResponseOutput() ConnectedClusterIdentityResponseOutput {
	return o
}

func (o ConnectedClusterIdentityResponseOutput) ToConnectedClusterIdentityResponseOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponseOutput {
	return o
}

func (o ConnectedClusterIdentityResponseOutput) ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput {
	return o.ToConnectedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ConnectedClusterIdentityResponseOutput) ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectedClusterIdentityResponse) *ConnectedClusterIdentityResponse {
		return &v
	}).(ConnectedClusterIdentityResponsePtrOutput)
}

func (o ConnectedClusterIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ConnectedClusterIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ConnectedClusterIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ConnectedClusterIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectedClusterIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterIdentityResponse)(nil)).Elem()
}

func (o ConnectedClusterIdentityResponsePtrOutput) ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput {
	return o
}

func (o ConnectedClusterIdentityResponsePtrOutput) ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponsePtrOutput {
	return o
}

func (o ConnectedClusterIdentityResponsePtrOutput) Elem() ConnectedClusterIdentityResponseOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentityResponse) ConnectedClusterIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ConnectedClusterIdentityResponse
		return ret
	}).(ConnectedClusterIdentityResponseOutput)
}

func (o ConnectedClusterIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ConnectedClusterIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ConnectedClusterIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type CredentialResultResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type CredentialResultResponseInput interface {
	pulumi.Input

	ToCredentialResultResponseOutput() CredentialResultResponseOutput
	ToCredentialResultResponseOutputWithContext(context.Context) CredentialResultResponseOutput
}

type CredentialResultResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (CredentialResultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialResultResponse)(nil)).Elem()
}

func (i CredentialResultResponseArgs) ToCredentialResultResponseOutput() CredentialResultResponseOutput {
	return i.ToCredentialResultResponseOutputWithContext(context.Background())
}

func (i CredentialResultResponseArgs) ToCredentialResultResponseOutputWithContext(ctx context.Context) CredentialResultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialResultResponseOutput)
}





type CredentialResultResponseArrayInput interface {
	pulumi.Input

	ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput
	ToCredentialResultResponseArrayOutputWithContext(context.Context) CredentialResultResponseArrayOutput
}

type CredentialResultResponseArray []CredentialResultResponseInput

func (CredentialResultResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CredentialResultResponse)(nil)).Elem()
}

func (i CredentialResultResponseArray) ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput {
	return i.ToCredentialResultResponseArrayOutputWithContext(context.Background())
}

func (i CredentialResultResponseArray) ToCredentialResultResponseArrayOutputWithContext(ctx context.Context) CredentialResultResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialResultResponseArrayOutput)
}

type CredentialResultResponseOutput struct{ *pulumi.OutputState }

func (CredentialResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialResultResponse)(nil)).Elem()
}

func (o CredentialResultResponseOutput) ToCredentialResultResponseOutput() CredentialResultResponseOutput {
	return o
}

func (o CredentialResultResponseOutput) ToCredentialResultResponseOutputWithContext(ctx context.Context) CredentialResultResponseOutput {
	return o
}

func (o CredentialResultResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CredentialResultResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CredentialResultResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v CredentialResultResponse) string { return v.Value }).(pulumi.StringOutput)
}

type CredentialResultResponseArrayOutput struct{ *pulumi.OutputState }

func (CredentialResultResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CredentialResultResponse)(nil)).Elem()
}

func (o CredentialResultResponseArrayOutput) ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput {
	return o
}

func (o CredentialResultResponseArrayOutput) ToCredentialResultResponseArrayOutputWithContext(ctx context.Context) CredentialResultResponseArrayOutput {
	return o
}

func (o CredentialResultResponseArrayOutput) Index(i pulumi.IntInput) CredentialResultResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CredentialResultResponse {
		return vs[0].([]CredentialResultResponse)[vs[1].(int)]
	}).(CredentialResultResponseOutput)
}

type HybridConnectionConfigResponse struct {
	ExpirationTime       float64 `pulumi:"expirationTime"`
	HybridConnectionName string  `pulumi:"hybridConnectionName"`
	Relay                string  `pulumi:"relay"`
	Token                string  `pulumi:"token"`
}





type HybridConnectionConfigResponseInput interface {
	pulumi.Input

	ToHybridConnectionConfigResponseOutput() HybridConnectionConfigResponseOutput
	ToHybridConnectionConfigResponseOutputWithContext(context.Context) HybridConnectionConfigResponseOutput
}

type HybridConnectionConfigResponseArgs struct {
	ExpirationTime       pulumi.Float64Input `pulumi:"expirationTime"`
	HybridConnectionName pulumi.StringInput  `pulumi:"hybridConnectionName"`
	Relay                pulumi.StringInput  `pulumi:"relay"`
	Token                pulumi.StringInput  `pulumi:"token"`
}

func (HybridConnectionConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionConfigResponse)(nil)).Elem()
}

func (i HybridConnectionConfigResponseArgs) ToHybridConnectionConfigResponseOutput() HybridConnectionConfigResponseOutput {
	return i.ToHybridConnectionConfigResponseOutputWithContext(context.Background())
}

func (i HybridConnectionConfigResponseArgs) ToHybridConnectionConfigResponseOutputWithContext(ctx context.Context) HybridConnectionConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridConnectionConfigResponseOutput)
}

type HybridConnectionConfigResponseOutput struct{ *pulumi.OutputState }

func (HybridConnectionConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionConfigResponse)(nil)).Elem()
}

func (o HybridConnectionConfigResponseOutput) ToHybridConnectionConfigResponseOutput() HybridConnectionConfigResponseOutput {
	return o
}

func (o HybridConnectionConfigResponseOutput) ToHybridConnectionConfigResponseOutputWithContext(ctx context.Context) HybridConnectionConfigResponseOutput {
	return o
}

func (o HybridConnectionConfigResponseOutput) ExpirationTime() pulumi.Float64Output {
	return o.ApplyT(func(v HybridConnectionConfigResponse) float64 { return v.ExpirationTime }).(pulumi.Float64Output)
}

func (o HybridConnectionConfigResponseOutput) HybridConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionConfigResponse) string { return v.HybridConnectionName }).(pulumi.StringOutput)
}

func (o HybridConnectionConfigResponseOutput) Relay() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionConfigResponse) string { return v.Relay }).(pulumi.StringOutput)
}

func (o HybridConnectionConfigResponseOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionConfigResponse) string { return v.Token }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthenticationDetailsValueOutput{})
	pulumi.RegisterOutputType(ConnectedClusterAADProfileOutput{})
	pulumi.RegisterOutputType(ConnectedClusterAADProfilePtrOutput{})
	pulumi.RegisterOutputType(ConnectedClusterAADProfileResponseOutput{})
	pulumi.RegisterOutputType(ConnectedClusterAADProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityPtrOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityResponseOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseArrayOutput{})
	pulumi.RegisterOutputType(HybridConnectionConfigResponseOutput{})
}

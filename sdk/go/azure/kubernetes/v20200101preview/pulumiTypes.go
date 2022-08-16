


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

func (o ConnectedClusterAADProfileOutput) ClientAppId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterAADProfile) string { return v.ClientAppId }).(pulumi.StringOutput)
}

func (o ConnectedClusterAADProfileOutput) ServerAppId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterAADProfile) string { return v.ServerAppId }).(pulumi.StringOutput)
}

func (o ConnectedClusterAADProfileOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterAADProfile) string { return v.TenantId }).(pulumi.StringOutput)
}

type ConnectedClusterAADProfileResponse struct {
	ClientAppId string `pulumi:"clientAppId"`
	ServerAppId string `pulumi:"serverAppId"`
	TenantId    string `pulumi:"tenantId"`
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

func (o ConnectedClusterAADProfileResponseOutput) ClientAppId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterAADProfileResponse) string { return v.ClientAppId }).(pulumi.StringOutput)
}

func (o ConnectedClusterAADProfileResponseOutput) ServerAppId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterAADProfileResponse) string { return v.ServerAppId }).(pulumi.StringOutput)
}

func (o ConnectedClusterAADProfileResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterAADProfileResponse) string { return v.TenantId }).(pulumi.StringOutput)
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

func (o ConnectedClusterIdentityOutput) Type() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v ConnectedClusterIdentity) ResourceIdentityType { return v.Type }).(ResourceIdentityTypeOutput)
}

type ConnectedClusterIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
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

func (o ConnectedClusterIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ConnectedClusterIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ConnectedClusterIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type CredentialResultResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
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
	pulumi.RegisterOutputType(ConnectedClusterAADProfileResponseOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityResponseOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseArrayOutput{})
	pulumi.RegisterOutputType(HybridConnectionConfigResponseOutput{})
}

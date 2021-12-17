


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthenticationDetailsValue struct {
	Token *string `pulumi:"token"`
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

type HybridConnectionConfigResponse struct {
	ExpirationTime       float64 `pulumi:"expirationTime"`
	HybridConnectionName string  `pulumi:"hybridConnectionName"`
	Relay                string  `pulumi:"relay"`
	Token                string  `pulumi:"token"`
}

func init() {
	pulumi.RegisterOutputType(ConnectedClusterAADProfileOutput{})
	pulumi.RegisterOutputType(ConnectedClusterAADProfileResponseOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityResponseOutput{})
}

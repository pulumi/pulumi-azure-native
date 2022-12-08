


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectedClusterIdentity struct {
	Type ResourceIdentityType `pulumi:"type"`
}


func (val *ConnectedClusterIdentity) Defaults() *ConnectedClusterIdentity {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		tmp.Type = ResourceIdentityType("SystemAssigned")
	}
	return &tmp
}





type ConnectedClusterIdentityInput interface {
	pulumi.Input

	ToConnectedClusterIdentityOutput() ConnectedClusterIdentityOutput
	ToConnectedClusterIdentityOutputWithContext(context.Context) ConnectedClusterIdentityOutput
}

type ConnectedClusterIdentityArgs struct {
	Type ResourceIdentityTypeInput `pulumi:"type"`
}


func (val *ConnectedClusterIdentityArgs) Defaults() *ConnectedClusterIdentityArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		tmp.Type = ResourceIdentityType("SystemAssigned")
	}
	return &tmp
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


func (val *ConnectedClusterIdentityResponse) Defaults() *ConnectedClusterIdentityResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		tmp.Type = "SystemAssigned"
	}
	return &tmp
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
	pulumi.RegisterOutputType(ConnectedClusterIdentityOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityResponseOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseArrayOutput{})
	pulumi.RegisterOutputType(HybridConnectionConfigResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

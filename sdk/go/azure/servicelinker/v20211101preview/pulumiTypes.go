


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecretAuthInfo struct {
	AuthType string  `pulumi:"authType"`
	Name     *string `pulumi:"name"`
	Secret   *string `pulumi:"secret"`
}

type SecretAuthInfoResponse struct {
	AuthType string  `pulumi:"authType"`
	Name     *string `pulumi:"name"`
	Secret   *string `pulumi:"secret"`
}

type SecretStore struct {
	KeyVaultId *string `pulumi:"keyVaultId"`
}





type SecretStoreInput interface {
	pulumi.Input

	ToSecretStoreOutput() SecretStoreOutput
	ToSecretStoreOutputWithContext(context.Context) SecretStoreOutput
}

type SecretStoreArgs struct {
	KeyVaultId pulumi.StringPtrInput `pulumi:"keyVaultId"`
}

func (SecretStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretStore)(nil)).Elem()
}

func (i SecretStoreArgs) ToSecretStoreOutput() SecretStoreOutput {
	return i.ToSecretStoreOutputWithContext(context.Background())
}

func (i SecretStoreArgs) ToSecretStoreOutputWithContext(ctx context.Context) SecretStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStoreOutput)
}

func (i SecretStoreArgs) ToSecretStorePtrOutput() SecretStorePtrOutput {
	return i.ToSecretStorePtrOutputWithContext(context.Background())
}

func (i SecretStoreArgs) ToSecretStorePtrOutputWithContext(ctx context.Context) SecretStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStoreOutput).ToSecretStorePtrOutputWithContext(ctx)
}









type SecretStorePtrInput interface {
	pulumi.Input

	ToSecretStorePtrOutput() SecretStorePtrOutput
	ToSecretStorePtrOutputWithContext(context.Context) SecretStorePtrOutput
}

type secretStorePtrType SecretStoreArgs

func SecretStorePtr(v *SecretStoreArgs) SecretStorePtrInput {
	return (*secretStorePtrType)(v)
}

func (*secretStorePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStore)(nil)).Elem()
}

func (i *secretStorePtrType) ToSecretStorePtrOutput() SecretStorePtrOutput {
	return i.ToSecretStorePtrOutputWithContext(context.Background())
}

func (i *secretStorePtrType) ToSecretStorePtrOutputWithContext(ctx context.Context) SecretStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStorePtrOutput)
}

type SecretStoreOutput struct{ *pulumi.OutputState }

func (SecretStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretStore)(nil)).Elem()
}

func (o SecretStoreOutput) ToSecretStoreOutput() SecretStoreOutput {
	return o
}

func (o SecretStoreOutput) ToSecretStoreOutputWithContext(ctx context.Context) SecretStoreOutput {
	return o
}

func (o SecretStoreOutput) ToSecretStorePtrOutput() SecretStorePtrOutput {
	return o.ToSecretStorePtrOutputWithContext(context.Background())
}

func (o SecretStoreOutput) ToSecretStorePtrOutputWithContext(ctx context.Context) SecretStorePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretStore) *SecretStore {
		return &v
	}).(SecretStorePtrOutput)
}

func (o SecretStoreOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretStore) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

type SecretStorePtrOutput struct{ *pulumi.OutputState }

func (SecretStorePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStore)(nil)).Elem()
}

func (o SecretStorePtrOutput) ToSecretStorePtrOutput() SecretStorePtrOutput {
	return o
}

func (o SecretStorePtrOutput) ToSecretStorePtrOutputWithContext(ctx context.Context) SecretStorePtrOutput {
	return o
}

func (o SecretStorePtrOutput) Elem() SecretStoreOutput {
	return o.ApplyT(func(v *SecretStore) SecretStore {
		if v != nil {
			return *v
		}
		var ret SecretStore
		return ret
	}).(SecretStoreOutput)
}

func (o SecretStorePtrOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretStore) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultId
	}).(pulumi.StringPtrOutput)
}

type SecretStoreResponse struct {
	KeyVaultId *string `pulumi:"keyVaultId"`
}

type SecretStoreResponseOutput struct{ *pulumi.OutputState }

func (SecretStoreResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretStoreResponse)(nil)).Elem()
}

func (o SecretStoreResponseOutput) ToSecretStoreResponseOutput() SecretStoreResponseOutput {
	return o
}

func (o SecretStoreResponseOutput) ToSecretStoreResponseOutputWithContext(ctx context.Context) SecretStoreResponseOutput {
	return o
}

func (o SecretStoreResponseOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretStoreResponse) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

type SecretStoreResponsePtrOutput struct{ *pulumi.OutputState }

func (SecretStoreResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStoreResponse)(nil)).Elem()
}

func (o SecretStoreResponsePtrOutput) ToSecretStoreResponsePtrOutput() SecretStoreResponsePtrOutput {
	return o
}

func (o SecretStoreResponsePtrOutput) ToSecretStoreResponsePtrOutputWithContext(ctx context.Context) SecretStoreResponsePtrOutput {
	return o
}

func (o SecretStoreResponsePtrOutput) Elem() SecretStoreResponseOutput {
	return o.ApplyT(func(v *SecretStoreResponse) SecretStoreResponse {
		if v != nil {
			return *v
		}
		var ret SecretStoreResponse
		return ret
	}).(SecretStoreResponseOutput)
}

func (o SecretStoreResponsePtrOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretStoreResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultId
	}).(pulumi.StringPtrOutput)
}

type ServicePrincipalCertificateAuthInfo struct {
	AuthType    string `pulumi:"authType"`
	Certificate string `pulumi:"certificate"`
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type ServicePrincipalCertificateAuthInfoResponse struct {
	AuthType    string `pulumi:"authType"`
	Certificate string `pulumi:"certificate"`
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type ServicePrincipalSecretAuthInfo struct {
	AuthType    string `pulumi:"authType"`
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
	Secret      string `pulumi:"secret"`
}

type ServicePrincipalSecretAuthInfoResponse struct {
	AuthType    string `pulumi:"authType"`
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
	Secret      string `pulumi:"secret"`
}

type SourceConfigurationResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type SourceConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SourceConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceConfigurationResponse)(nil)).Elem()
}

func (o SourceConfigurationResponseOutput) ToSourceConfigurationResponseOutput() SourceConfigurationResponseOutput {
	return o
}

func (o SourceConfigurationResponseOutput) ToSourceConfigurationResponseOutputWithContext(ctx context.Context) SourceConfigurationResponseOutput {
	return o
}

func (o SourceConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SourceConfigurationResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceConfigurationResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SourceConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (SourceConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceConfigurationResponse)(nil)).Elem()
}

func (o SourceConfigurationResponseArrayOutput) ToSourceConfigurationResponseArrayOutput() SourceConfigurationResponseArrayOutput {
	return o
}

func (o SourceConfigurationResponseArrayOutput) ToSourceConfigurationResponseArrayOutputWithContext(ctx context.Context) SourceConfigurationResponseArrayOutput {
	return o
}

func (o SourceConfigurationResponseArrayOutput) Index(i pulumi.IntInput) SourceConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SourceConfigurationResponse {
		return vs[0].([]SourceConfigurationResponse)[vs[1].(int)]
	}).(SourceConfigurationResponseOutput)
}

type SystemAssignedIdentityAuthInfo struct {
	AuthType string `pulumi:"authType"`
}

type SystemAssignedIdentityAuthInfoResponse struct {
	AuthType string `pulumi:"authType"`
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

type UserAssignedIdentityAuthInfo struct {
	AuthType       string `pulumi:"authType"`
	ClientId       string `pulumi:"clientId"`
	SubscriptionId string `pulumi:"subscriptionId"`
}

type UserAssignedIdentityAuthInfoResponse struct {
	AuthType       string `pulumi:"authType"`
	ClientId       string `pulumi:"clientId"`
	SubscriptionId string `pulumi:"subscriptionId"`
}

type VNetSolution struct {
	Type *string `pulumi:"type"`
}





type VNetSolutionInput interface {
	pulumi.Input

	ToVNetSolutionOutput() VNetSolutionOutput
	ToVNetSolutionOutputWithContext(context.Context) VNetSolutionOutput
}

type VNetSolutionArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (VNetSolutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolution)(nil)).Elem()
}

func (i VNetSolutionArgs) ToVNetSolutionOutput() VNetSolutionOutput {
	return i.ToVNetSolutionOutputWithContext(context.Background())
}

func (i VNetSolutionArgs) ToVNetSolutionOutputWithContext(ctx context.Context) VNetSolutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VNetSolutionOutput)
}

func (i VNetSolutionArgs) ToVNetSolutionPtrOutput() VNetSolutionPtrOutput {
	return i.ToVNetSolutionPtrOutputWithContext(context.Background())
}

func (i VNetSolutionArgs) ToVNetSolutionPtrOutputWithContext(ctx context.Context) VNetSolutionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VNetSolutionOutput).ToVNetSolutionPtrOutputWithContext(ctx)
}









type VNetSolutionPtrInput interface {
	pulumi.Input

	ToVNetSolutionPtrOutput() VNetSolutionPtrOutput
	ToVNetSolutionPtrOutputWithContext(context.Context) VNetSolutionPtrOutput
}

type vnetSolutionPtrType VNetSolutionArgs

func VNetSolutionPtr(v *VNetSolutionArgs) VNetSolutionPtrInput {
	return (*vnetSolutionPtrType)(v)
}

func (*vnetSolutionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VNetSolution)(nil)).Elem()
}

func (i *vnetSolutionPtrType) ToVNetSolutionPtrOutput() VNetSolutionPtrOutput {
	return i.ToVNetSolutionPtrOutputWithContext(context.Background())
}

func (i *vnetSolutionPtrType) ToVNetSolutionPtrOutputWithContext(ctx context.Context) VNetSolutionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VNetSolutionPtrOutput)
}

type VNetSolutionOutput struct{ *pulumi.OutputState }

func (VNetSolutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolution)(nil)).Elem()
}

func (o VNetSolutionOutput) ToVNetSolutionOutput() VNetSolutionOutput {
	return o
}

func (o VNetSolutionOutput) ToVNetSolutionOutputWithContext(ctx context.Context) VNetSolutionOutput {
	return o
}

func (o VNetSolutionOutput) ToVNetSolutionPtrOutput() VNetSolutionPtrOutput {
	return o.ToVNetSolutionPtrOutputWithContext(context.Background())
}

func (o VNetSolutionOutput) ToVNetSolutionPtrOutputWithContext(ctx context.Context) VNetSolutionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VNetSolution) *VNetSolution {
		return &v
	}).(VNetSolutionPtrOutput)
}

func (o VNetSolutionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VNetSolution) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VNetSolutionPtrOutput struct{ *pulumi.OutputState }

func (VNetSolutionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VNetSolution)(nil)).Elem()
}

func (o VNetSolutionPtrOutput) ToVNetSolutionPtrOutput() VNetSolutionPtrOutput {
	return o
}

func (o VNetSolutionPtrOutput) ToVNetSolutionPtrOutputWithContext(ctx context.Context) VNetSolutionPtrOutput {
	return o
}

func (o VNetSolutionPtrOutput) Elem() VNetSolutionOutput {
	return o.ApplyT(func(v *VNetSolution) VNetSolution {
		if v != nil {
			return *v
		}
		var ret VNetSolution
		return ret
	}).(VNetSolutionOutput)
}

func (o VNetSolutionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VNetSolution) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type VNetSolutionResponse struct {
	Type *string `pulumi:"type"`
}

type VNetSolutionResponseOutput struct{ *pulumi.OutputState }

func (VNetSolutionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolutionResponse)(nil)).Elem()
}

func (o VNetSolutionResponseOutput) ToVNetSolutionResponseOutput() VNetSolutionResponseOutput {
	return o
}

func (o VNetSolutionResponseOutput) ToVNetSolutionResponseOutputWithContext(ctx context.Context) VNetSolutionResponseOutput {
	return o
}

func (o VNetSolutionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VNetSolutionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VNetSolutionResponsePtrOutput struct{ *pulumi.OutputState }

func (VNetSolutionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VNetSolutionResponse)(nil)).Elem()
}

func (o VNetSolutionResponsePtrOutput) ToVNetSolutionResponsePtrOutput() VNetSolutionResponsePtrOutput {
	return o
}

func (o VNetSolutionResponsePtrOutput) ToVNetSolutionResponsePtrOutputWithContext(ctx context.Context) VNetSolutionResponsePtrOutput {
	return o
}

func (o VNetSolutionResponsePtrOutput) Elem() VNetSolutionResponseOutput {
	return o.ApplyT(func(v *VNetSolutionResponse) VNetSolutionResponse {
		if v != nil {
			return *v
		}
		var ret VNetSolutionResponse
		return ret
	}).(VNetSolutionResponseOutput)
}

func (o VNetSolutionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VNetSolutionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretStoreOutput{})
	pulumi.RegisterOutputType(SecretStorePtrOutput{})
	pulumi.RegisterOutputType(SecretStoreResponseOutput{})
	pulumi.RegisterOutputType(SecretStoreResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SourceConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VNetSolutionOutput{})
	pulumi.RegisterOutputType(VNetSolutionPtrOutput{})
	pulumi.RegisterOutputType(VNetSolutionResponseOutput{})
	pulumi.RegisterOutputType(VNetSolutionResponsePtrOutput{})
}

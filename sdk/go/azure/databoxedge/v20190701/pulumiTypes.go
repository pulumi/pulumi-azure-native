


package v20190701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Address struct {
	AddressLine1 string  `pulumi:"addressLine1"`
	AddressLine2 *string `pulumi:"addressLine2"`
	AddressLine3 *string `pulumi:"addressLine3"`
	City         string  `pulumi:"city"`
	Country      string  `pulumi:"country"`
	PostalCode   string  `pulumi:"postalCode"`
	State        string  `pulumi:"state"`
}





type AddressInput interface {
	pulumi.Input

	ToAddressOutput() AddressOutput
	ToAddressOutputWithContext(context.Context) AddressOutput
}

type AddressArgs struct {
	AddressLine1 pulumi.StringInput    `pulumi:"addressLine1"`
	AddressLine2 pulumi.StringPtrInput `pulumi:"addressLine2"`
	AddressLine3 pulumi.StringPtrInput `pulumi:"addressLine3"`
	City         pulumi.StringInput    `pulumi:"city"`
	Country      pulumi.StringInput    `pulumi:"country"`
	PostalCode   pulumi.StringInput    `pulumi:"postalCode"`
	State        pulumi.StringInput    `pulumi:"state"`
}

func (AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Address)(nil)).Elem()
}

func (i AddressArgs) ToAddressOutput() AddressOutput {
	return i.ToAddressOutputWithContext(context.Background())
}

func (i AddressArgs) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput)
}

type AddressOutput struct{ *pulumi.OutputState }

func (AddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Address)(nil)).Elem()
}

func (o AddressOutput) ToAddressOutput() AddressOutput {
	return o
}

func (o AddressOutput) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return o
}

func (o AddressOutput) AddressLine1() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.AddressLine1 }).(pulumi.StringOutput)
}

func (o AddressOutput) AddressLine2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Address) *string { return v.AddressLine2 }).(pulumi.StringPtrOutput)
}

func (o AddressOutput) AddressLine3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Address) *string { return v.AddressLine3 }).(pulumi.StringPtrOutput)
}

func (o AddressOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.City }).(pulumi.StringOutput)
}

func (o AddressOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.Country }).(pulumi.StringOutput)
}

func (o AddressOutput) PostalCode() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.PostalCode }).(pulumi.StringOutput)
}

func (o AddressOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.State }).(pulumi.StringOutput)
}

type AddressResponse struct {
	AddressLine1 string  `pulumi:"addressLine1"`
	AddressLine2 *string `pulumi:"addressLine2"`
	AddressLine3 *string `pulumi:"addressLine3"`
	City         string  `pulumi:"city"`
	Country      string  `pulumi:"country"`
	PostalCode   string  `pulumi:"postalCode"`
	State        string  `pulumi:"state"`
}

type AddressResponseOutput struct{ *pulumi.OutputState }

func (AddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressResponse)(nil)).Elem()
}

func (o AddressResponseOutput) ToAddressResponseOutput() AddressResponseOutput {
	return o
}

func (o AddressResponseOutput) ToAddressResponseOutputWithContext(ctx context.Context) AddressResponseOutput {
	return o
}

func (o AddressResponseOutput) AddressLine1() pulumi.StringOutput {
	return o.ApplyT(func(v AddressResponse) string { return v.AddressLine1 }).(pulumi.StringOutput)
}

func (o AddressResponseOutput) AddressLine2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressResponse) *string { return v.AddressLine2 }).(pulumi.StringPtrOutput)
}

func (o AddressResponseOutput) AddressLine3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressResponse) *string { return v.AddressLine3 }).(pulumi.StringPtrOutput)
}

func (o AddressResponseOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v AddressResponse) string { return v.City }).(pulumi.StringOutput)
}

func (o AddressResponseOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v AddressResponse) string { return v.Country }).(pulumi.StringOutput)
}

func (o AddressResponseOutput) PostalCode() pulumi.StringOutput {
	return o.ApplyT(func(v AddressResponse) string { return v.PostalCode }).(pulumi.StringOutput)
}

func (o AddressResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AddressResponse) string { return v.State }).(pulumi.StringOutput)
}

type AsymmetricEncryptedSecret struct {
	EncryptionAlgorithm      string  `pulumi:"encryptionAlgorithm"`
	EncryptionCertThumbprint *string `pulumi:"encryptionCertThumbprint"`
	Value                    string  `pulumi:"value"`
}





type AsymmetricEncryptedSecretInput interface {
	pulumi.Input

	ToAsymmetricEncryptedSecretOutput() AsymmetricEncryptedSecretOutput
	ToAsymmetricEncryptedSecretOutputWithContext(context.Context) AsymmetricEncryptedSecretOutput
}

type AsymmetricEncryptedSecretArgs struct {
	EncryptionAlgorithm      pulumi.StringInput    `pulumi:"encryptionAlgorithm"`
	EncryptionCertThumbprint pulumi.StringPtrInput `pulumi:"encryptionCertThumbprint"`
	Value                    pulumi.StringInput    `pulumi:"value"`
}

func (AsymmetricEncryptedSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AsymmetricEncryptedSecret)(nil)).Elem()
}

func (i AsymmetricEncryptedSecretArgs) ToAsymmetricEncryptedSecretOutput() AsymmetricEncryptedSecretOutput {
	return i.ToAsymmetricEncryptedSecretOutputWithContext(context.Background())
}

func (i AsymmetricEncryptedSecretArgs) ToAsymmetricEncryptedSecretOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretOutput)
}

func (i AsymmetricEncryptedSecretArgs) ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput {
	return i.ToAsymmetricEncryptedSecretPtrOutputWithContext(context.Background())
}

func (i AsymmetricEncryptedSecretArgs) ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretOutput).ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx)
}









type AsymmetricEncryptedSecretPtrInput interface {
	pulumi.Input

	ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput
	ToAsymmetricEncryptedSecretPtrOutputWithContext(context.Context) AsymmetricEncryptedSecretPtrOutput
}

type asymmetricEncryptedSecretPtrType AsymmetricEncryptedSecretArgs

func AsymmetricEncryptedSecretPtr(v *AsymmetricEncryptedSecretArgs) AsymmetricEncryptedSecretPtrInput {
	return (*asymmetricEncryptedSecretPtrType)(v)
}

func (*asymmetricEncryptedSecretPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AsymmetricEncryptedSecret)(nil)).Elem()
}

func (i *asymmetricEncryptedSecretPtrType) ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput {
	return i.ToAsymmetricEncryptedSecretPtrOutputWithContext(context.Background())
}

func (i *asymmetricEncryptedSecretPtrType) ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretPtrOutput)
}

type AsymmetricEncryptedSecretOutput struct{ *pulumi.OutputState }

func (AsymmetricEncryptedSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AsymmetricEncryptedSecret)(nil)).Elem()
}

func (o AsymmetricEncryptedSecretOutput) ToAsymmetricEncryptedSecretOutput() AsymmetricEncryptedSecretOutput {
	return o
}

func (o AsymmetricEncryptedSecretOutput) ToAsymmetricEncryptedSecretOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretOutput {
	return o
}

func (o AsymmetricEncryptedSecretOutput) ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput {
	return o.ToAsymmetricEncryptedSecretPtrOutputWithContext(context.Background())
}

func (o AsymmetricEncryptedSecretOutput) ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AsymmetricEncryptedSecret) *AsymmetricEncryptedSecret {
		return &v
	}).(AsymmetricEncryptedSecretPtrOutput)
}

func (o AsymmetricEncryptedSecretOutput) EncryptionAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecret) string { return v.EncryptionAlgorithm }).(pulumi.StringOutput)
}

func (o AsymmetricEncryptedSecretOutput) EncryptionCertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecret) *string { return v.EncryptionCertThumbprint }).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecret) string { return v.Value }).(pulumi.StringOutput)
}

type AsymmetricEncryptedSecretPtrOutput struct{ *pulumi.OutputState }

func (AsymmetricEncryptedSecretPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AsymmetricEncryptedSecret)(nil)).Elem()
}

func (o AsymmetricEncryptedSecretPtrOutput) ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput {
	return o
}

func (o AsymmetricEncryptedSecretPtrOutput) ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretPtrOutput {
	return o
}

func (o AsymmetricEncryptedSecretPtrOutput) Elem() AsymmetricEncryptedSecretOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecret) AsymmetricEncryptedSecret {
		if v != nil {
			return *v
		}
		var ret AsymmetricEncryptedSecret
		return ret
	}).(AsymmetricEncryptedSecretOutput)
}

func (o AsymmetricEncryptedSecretPtrOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecret) *string {
		if v == nil {
			return nil
		}
		return &v.EncryptionAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretPtrOutput) EncryptionCertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecret) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionCertThumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecret) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type AsymmetricEncryptedSecretResponse struct {
	EncryptionAlgorithm      string  `pulumi:"encryptionAlgorithm"`
	EncryptionCertThumbprint *string `pulumi:"encryptionCertThumbprint"`
	Value                    string  `pulumi:"value"`
}

type AsymmetricEncryptedSecretResponseOutput struct{ *pulumi.OutputState }

func (AsymmetricEncryptedSecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AsymmetricEncryptedSecretResponse)(nil)).Elem()
}

func (o AsymmetricEncryptedSecretResponseOutput) ToAsymmetricEncryptedSecretResponseOutput() AsymmetricEncryptedSecretResponseOutput {
	return o
}

func (o AsymmetricEncryptedSecretResponseOutput) ToAsymmetricEncryptedSecretResponseOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponseOutput {
	return o
}

func (o AsymmetricEncryptedSecretResponseOutput) EncryptionAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecretResponse) string { return v.EncryptionAlgorithm }).(pulumi.StringOutput)
}

func (o AsymmetricEncryptedSecretResponseOutput) EncryptionCertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecretResponse) *string { return v.EncryptionCertThumbprint }).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecretResponse) string { return v.Value }).(pulumi.StringOutput)
}

type AsymmetricEncryptedSecretResponsePtrOutput struct{ *pulumi.OutputState }

func (AsymmetricEncryptedSecretResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AsymmetricEncryptedSecretResponse)(nil)).Elem()
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput {
	return o
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponsePtrOutput {
	return o
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) Elem() AsymmetricEncryptedSecretResponseOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecretResponse) AsymmetricEncryptedSecretResponse {
		if v != nil {
			return *v
		}
		var ret AsymmetricEncryptedSecretResponse
		return ret
	}).(AsymmetricEncryptedSecretResponseOutput)
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecretResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EncryptionAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) EncryptionCertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecretResponse) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionCertThumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecretResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type Authentication struct {
	SymmetricKey *SymmetricKey `pulumi:"symmetricKey"`
}





type AuthenticationInput interface {
	pulumi.Input

	ToAuthenticationOutput() AuthenticationOutput
	ToAuthenticationOutputWithContext(context.Context) AuthenticationOutput
}

type AuthenticationArgs struct {
	SymmetricKey SymmetricKeyPtrInput `pulumi:"symmetricKey"`
}

func (AuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Authentication)(nil)).Elem()
}

func (i AuthenticationArgs) ToAuthenticationOutput() AuthenticationOutput {
	return i.ToAuthenticationOutputWithContext(context.Background())
}

func (i AuthenticationArgs) ToAuthenticationOutputWithContext(ctx context.Context) AuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationOutput)
}

func (i AuthenticationArgs) ToAuthenticationPtrOutput() AuthenticationPtrOutput {
	return i.ToAuthenticationPtrOutputWithContext(context.Background())
}

func (i AuthenticationArgs) ToAuthenticationPtrOutputWithContext(ctx context.Context) AuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationOutput).ToAuthenticationPtrOutputWithContext(ctx)
}









type AuthenticationPtrInput interface {
	pulumi.Input

	ToAuthenticationPtrOutput() AuthenticationPtrOutput
	ToAuthenticationPtrOutputWithContext(context.Context) AuthenticationPtrOutput
}

type authenticationPtrType AuthenticationArgs

func AuthenticationPtr(v *AuthenticationArgs) AuthenticationPtrInput {
	return (*authenticationPtrType)(v)
}

func (*authenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Authentication)(nil)).Elem()
}

func (i *authenticationPtrType) ToAuthenticationPtrOutput() AuthenticationPtrOutput {
	return i.ToAuthenticationPtrOutputWithContext(context.Background())
}

func (i *authenticationPtrType) ToAuthenticationPtrOutputWithContext(ctx context.Context) AuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationPtrOutput)
}

type AuthenticationOutput struct{ *pulumi.OutputState }

func (AuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Authentication)(nil)).Elem()
}

func (o AuthenticationOutput) ToAuthenticationOutput() AuthenticationOutput {
	return o
}

func (o AuthenticationOutput) ToAuthenticationOutputWithContext(ctx context.Context) AuthenticationOutput {
	return o
}

func (o AuthenticationOutput) ToAuthenticationPtrOutput() AuthenticationPtrOutput {
	return o.ToAuthenticationPtrOutputWithContext(context.Background())
}

func (o AuthenticationOutput) ToAuthenticationPtrOutputWithContext(ctx context.Context) AuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Authentication) *Authentication {
		return &v
	}).(AuthenticationPtrOutput)
}

func (o AuthenticationOutput) SymmetricKey() SymmetricKeyPtrOutput {
	return o.ApplyT(func(v Authentication) *SymmetricKey { return v.SymmetricKey }).(SymmetricKeyPtrOutput)
}

type AuthenticationPtrOutput struct{ *pulumi.OutputState }

func (AuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Authentication)(nil)).Elem()
}

func (o AuthenticationPtrOutput) ToAuthenticationPtrOutput() AuthenticationPtrOutput {
	return o
}

func (o AuthenticationPtrOutput) ToAuthenticationPtrOutputWithContext(ctx context.Context) AuthenticationPtrOutput {
	return o
}

func (o AuthenticationPtrOutput) Elem() AuthenticationOutput {
	return o.ApplyT(func(v *Authentication) Authentication {
		if v != nil {
			return *v
		}
		var ret Authentication
		return ret
	}).(AuthenticationOutput)
}

func (o AuthenticationPtrOutput) SymmetricKey() SymmetricKeyPtrOutput {
	return o.ApplyT(func(v *Authentication) *SymmetricKey {
		if v == nil {
			return nil
		}
		return v.SymmetricKey
	}).(SymmetricKeyPtrOutput)
}

type AuthenticationResponse struct {
	SymmetricKey *SymmetricKeyResponse `pulumi:"symmetricKey"`
}

type AuthenticationResponseOutput struct{ *pulumi.OutputState }

func (AuthenticationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationResponse)(nil)).Elem()
}

func (o AuthenticationResponseOutput) ToAuthenticationResponseOutput() AuthenticationResponseOutput {
	return o
}

func (o AuthenticationResponseOutput) ToAuthenticationResponseOutputWithContext(ctx context.Context) AuthenticationResponseOutput {
	return o
}

func (o AuthenticationResponseOutput) SymmetricKey() SymmetricKeyResponsePtrOutput {
	return o.ApplyT(func(v AuthenticationResponse) *SymmetricKeyResponse { return v.SymmetricKey }).(SymmetricKeyResponsePtrOutput)
}

type AuthenticationResponsePtrOutput struct{ *pulumi.OutputState }

func (AuthenticationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationResponse)(nil)).Elem()
}

func (o AuthenticationResponsePtrOutput) ToAuthenticationResponsePtrOutput() AuthenticationResponsePtrOutput {
	return o
}

func (o AuthenticationResponsePtrOutput) ToAuthenticationResponsePtrOutputWithContext(ctx context.Context) AuthenticationResponsePtrOutput {
	return o
}

func (o AuthenticationResponsePtrOutput) Elem() AuthenticationResponseOutput {
	return o.ApplyT(func(v *AuthenticationResponse) AuthenticationResponse {
		if v != nil {
			return *v
		}
		var ret AuthenticationResponse
		return ret
	}).(AuthenticationResponseOutput)
}

func (o AuthenticationResponsePtrOutput) SymmetricKey() SymmetricKeyResponsePtrOutput {
	return o.ApplyT(func(v *AuthenticationResponse) *SymmetricKeyResponse {
		if v == nil {
			return nil
		}
		return v.SymmetricKey
	}).(SymmetricKeyResponsePtrOutput)
}

type AzureContainerInfo struct {
	ContainerName              string `pulumi:"containerName"`
	DataFormat                 string `pulumi:"dataFormat"`
	StorageAccountCredentialId string `pulumi:"storageAccountCredentialId"`
}





type AzureContainerInfoInput interface {
	pulumi.Input

	ToAzureContainerInfoOutput() AzureContainerInfoOutput
	ToAzureContainerInfoOutputWithContext(context.Context) AzureContainerInfoOutput
}

type AzureContainerInfoArgs struct {
	ContainerName              pulumi.StringInput `pulumi:"containerName"`
	DataFormat                 pulumi.StringInput `pulumi:"dataFormat"`
	StorageAccountCredentialId pulumi.StringInput `pulumi:"storageAccountCredentialId"`
}

func (AzureContainerInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureContainerInfo)(nil)).Elem()
}

func (i AzureContainerInfoArgs) ToAzureContainerInfoOutput() AzureContainerInfoOutput {
	return i.ToAzureContainerInfoOutputWithContext(context.Background())
}

func (i AzureContainerInfoArgs) ToAzureContainerInfoOutputWithContext(ctx context.Context) AzureContainerInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureContainerInfoOutput)
}

func (i AzureContainerInfoArgs) ToAzureContainerInfoPtrOutput() AzureContainerInfoPtrOutput {
	return i.ToAzureContainerInfoPtrOutputWithContext(context.Background())
}

func (i AzureContainerInfoArgs) ToAzureContainerInfoPtrOutputWithContext(ctx context.Context) AzureContainerInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureContainerInfoOutput).ToAzureContainerInfoPtrOutputWithContext(ctx)
}









type AzureContainerInfoPtrInput interface {
	pulumi.Input

	ToAzureContainerInfoPtrOutput() AzureContainerInfoPtrOutput
	ToAzureContainerInfoPtrOutputWithContext(context.Context) AzureContainerInfoPtrOutput
}

type azureContainerInfoPtrType AzureContainerInfoArgs

func AzureContainerInfoPtr(v *AzureContainerInfoArgs) AzureContainerInfoPtrInput {
	return (*azureContainerInfoPtrType)(v)
}

func (*azureContainerInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureContainerInfo)(nil)).Elem()
}

func (i *azureContainerInfoPtrType) ToAzureContainerInfoPtrOutput() AzureContainerInfoPtrOutput {
	return i.ToAzureContainerInfoPtrOutputWithContext(context.Background())
}

func (i *azureContainerInfoPtrType) ToAzureContainerInfoPtrOutputWithContext(ctx context.Context) AzureContainerInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureContainerInfoPtrOutput)
}

type AzureContainerInfoOutput struct{ *pulumi.OutputState }

func (AzureContainerInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureContainerInfo)(nil)).Elem()
}

func (o AzureContainerInfoOutput) ToAzureContainerInfoOutput() AzureContainerInfoOutput {
	return o
}

func (o AzureContainerInfoOutput) ToAzureContainerInfoOutputWithContext(ctx context.Context) AzureContainerInfoOutput {
	return o
}

func (o AzureContainerInfoOutput) ToAzureContainerInfoPtrOutput() AzureContainerInfoPtrOutput {
	return o.ToAzureContainerInfoPtrOutputWithContext(context.Background())
}

func (o AzureContainerInfoOutput) ToAzureContainerInfoPtrOutputWithContext(ctx context.Context) AzureContainerInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureContainerInfo) *AzureContainerInfo {
		return &v
	}).(AzureContainerInfoPtrOutput)
}

func (o AzureContainerInfoOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v AzureContainerInfo) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o AzureContainerInfoOutput) DataFormat() pulumi.StringOutput {
	return o.ApplyT(func(v AzureContainerInfo) string { return v.DataFormat }).(pulumi.StringOutput)
}

func (o AzureContainerInfoOutput) StorageAccountCredentialId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureContainerInfo) string { return v.StorageAccountCredentialId }).(pulumi.StringOutput)
}

type AzureContainerInfoPtrOutput struct{ *pulumi.OutputState }

func (AzureContainerInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureContainerInfo)(nil)).Elem()
}

func (o AzureContainerInfoPtrOutput) ToAzureContainerInfoPtrOutput() AzureContainerInfoPtrOutput {
	return o
}

func (o AzureContainerInfoPtrOutput) ToAzureContainerInfoPtrOutputWithContext(ctx context.Context) AzureContainerInfoPtrOutput {
	return o
}

func (o AzureContainerInfoPtrOutput) Elem() AzureContainerInfoOutput {
	return o.ApplyT(func(v *AzureContainerInfo) AzureContainerInfo {
		if v != nil {
			return *v
		}
		var ret AzureContainerInfo
		return ret
	}).(AzureContainerInfoOutput)
}

func (o AzureContainerInfoPtrOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureContainerInfo) *string {
		if v == nil {
			return nil
		}
		return &v.ContainerName
	}).(pulumi.StringPtrOutput)
}

func (o AzureContainerInfoPtrOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureContainerInfo) *string {
		if v == nil {
			return nil
		}
		return &v.DataFormat
	}).(pulumi.StringPtrOutput)
}

func (o AzureContainerInfoPtrOutput) StorageAccountCredentialId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureContainerInfo) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountCredentialId
	}).(pulumi.StringPtrOutput)
}

type AzureContainerInfoResponse struct {
	ContainerName              string `pulumi:"containerName"`
	DataFormat                 string `pulumi:"dataFormat"`
	StorageAccountCredentialId string `pulumi:"storageAccountCredentialId"`
}

type AzureContainerInfoResponseOutput struct{ *pulumi.OutputState }

func (AzureContainerInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureContainerInfoResponse)(nil)).Elem()
}

func (o AzureContainerInfoResponseOutput) ToAzureContainerInfoResponseOutput() AzureContainerInfoResponseOutput {
	return o
}

func (o AzureContainerInfoResponseOutput) ToAzureContainerInfoResponseOutputWithContext(ctx context.Context) AzureContainerInfoResponseOutput {
	return o
}

func (o AzureContainerInfoResponseOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v AzureContainerInfoResponse) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o AzureContainerInfoResponseOutput) DataFormat() pulumi.StringOutput {
	return o.ApplyT(func(v AzureContainerInfoResponse) string { return v.DataFormat }).(pulumi.StringOutput)
}

func (o AzureContainerInfoResponseOutput) StorageAccountCredentialId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureContainerInfoResponse) string { return v.StorageAccountCredentialId }).(pulumi.StringOutput)
}

type AzureContainerInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureContainerInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureContainerInfoResponse)(nil)).Elem()
}

func (o AzureContainerInfoResponsePtrOutput) ToAzureContainerInfoResponsePtrOutput() AzureContainerInfoResponsePtrOutput {
	return o
}

func (o AzureContainerInfoResponsePtrOutput) ToAzureContainerInfoResponsePtrOutputWithContext(ctx context.Context) AzureContainerInfoResponsePtrOutput {
	return o
}

func (o AzureContainerInfoResponsePtrOutput) Elem() AzureContainerInfoResponseOutput {
	return o.ApplyT(func(v *AzureContainerInfoResponse) AzureContainerInfoResponse {
		if v != nil {
			return *v
		}
		var ret AzureContainerInfoResponse
		return ret
	}).(AzureContainerInfoResponseOutput)
}

func (o AzureContainerInfoResponsePtrOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureContainerInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ContainerName
	}).(pulumi.StringPtrOutput)
}

func (o AzureContainerInfoResponsePtrOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureContainerInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DataFormat
	}).(pulumi.StringPtrOutput)
}

func (o AzureContainerInfoResponsePtrOutput) StorageAccountCredentialId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureContainerInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountCredentialId
	}).(pulumi.StringPtrOutput)
}

type ClientAccessRight struct {
	AccessPermission string `pulumi:"accessPermission"`
	Client           string `pulumi:"client"`
}





type ClientAccessRightInput interface {
	pulumi.Input

	ToClientAccessRightOutput() ClientAccessRightOutput
	ToClientAccessRightOutputWithContext(context.Context) ClientAccessRightOutput
}

type ClientAccessRightArgs struct {
	AccessPermission pulumi.StringInput `pulumi:"accessPermission"`
	Client           pulumi.StringInput `pulumi:"client"`
}

func (ClientAccessRightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAccessRight)(nil)).Elem()
}

func (i ClientAccessRightArgs) ToClientAccessRightOutput() ClientAccessRightOutput {
	return i.ToClientAccessRightOutputWithContext(context.Background())
}

func (i ClientAccessRightArgs) ToClientAccessRightOutputWithContext(ctx context.Context) ClientAccessRightOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAccessRightOutput)
}





type ClientAccessRightArrayInput interface {
	pulumi.Input

	ToClientAccessRightArrayOutput() ClientAccessRightArrayOutput
	ToClientAccessRightArrayOutputWithContext(context.Context) ClientAccessRightArrayOutput
}

type ClientAccessRightArray []ClientAccessRightInput

func (ClientAccessRightArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientAccessRight)(nil)).Elem()
}

func (i ClientAccessRightArray) ToClientAccessRightArrayOutput() ClientAccessRightArrayOutput {
	return i.ToClientAccessRightArrayOutputWithContext(context.Background())
}

func (i ClientAccessRightArray) ToClientAccessRightArrayOutputWithContext(ctx context.Context) ClientAccessRightArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAccessRightArrayOutput)
}

type ClientAccessRightOutput struct{ *pulumi.OutputState }

func (ClientAccessRightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAccessRight)(nil)).Elem()
}

func (o ClientAccessRightOutput) ToClientAccessRightOutput() ClientAccessRightOutput {
	return o
}

func (o ClientAccessRightOutput) ToClientAccessRightOutputWithContext(ctx context.Context) ClientAccessRightOutput {
	return o
}

func (o ClientAccessRightOutput) AccessPermission() pulumi.StringOutput {
	return o.ApplyT(func(v ClientAccessRight) string { return v.AccessPermission }).(pulumi.StringOutput)
}

func (o ClientAccessRightOutput) Client() pulumi.StringOutput {
	return o.ApplyT(func(v ClientAccessRight) string { return v.Client }).(pulumi.StringOutput)
}

type ClientAccessRightArrayOutput struct{ *pulumi.OutputState }

func (ClientAccessRightArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientAccessRight)(nil)).Elem()
}

func (o ClientAccessRightArrayOutput) ToClientAccessRightArrayOutput() ClientAccessRightArrayOutput {
	return o
}

func (o ClientAccessRightArrayOutput) ToClientAccessRightArrayOutputWithContext(ctx context.Context) ClientAccessRightArrayOutput {
	return o
}

func (o ClientAccessRightArrayOutput) Index(i pulumi.IntInput) ClientAccessRightOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientAccessRight {
		return vs[0].([]ClientAccessRight)[vs[1].(int)]
	}).(ClientAccessRightOutput)
}

type ClientAccessRightResponse struct {
	AccessPermission string `pulumi:"accessPermission"`
	Client           string `pulumi:"client"`
}

type ClientAccessRightResponseOutput struct{ *pulumi.OutputState }

func (ClientAccessRightResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAccessRightResponse)(nil)).Elem()
}

func (o ClientAccessRightResponseOutput) ToClientAccessRightResponseOutput() ClientAccessRightResponseOutput {
	return o
}

func (o ClientAccessRightResponseOutput) ToClientAccessRightResponseOutputWithContext(ctx context.Context) ClientAccessRightResponseOutput {
	return o
}

func (o ClientAccessRightResponseOutput) AccessPermission() pulumi.StringOutput {
	return o.ApplyT(func(v ClientAccessRightResponse) string { return v.AccessPermission }).(pulumi.StringOutput)
}

func (o ClientAccessRightResponseOutput) Client() pulumi.StringOutput {
	return o.ApplyT(func(v ClientAccessRightResponse) string { return v.Client }).(pulumi.StringOutput)
}

type ClientAccessRightResponseArrayOutput struct{ *pulumi.OutputState }

func (ClientAccessRightResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientAccessRightResponse)(nil)).Elem()
}

func (o ClientAccessRightResponseArrayOutput) ToClientAccessRightResponseArrayOutput() ClientAccessRightResponseArrayOutput {
	return o
}

func (o ClientAccessRightResponseArrayOutput) ToClientAccessRightResponseArrayOutputWithContext(ctx context.Context) ClientAccessRightResponseArrayOutput {
	return o
}

func (o ClientAccessRightResponseArrayOutput) Index(i pulumi.IntInput) ClientAccessRightResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientAccessRightResponse {
		return vs[0].([]ClientAccessRightResponse)[vs[1].(int)]
	}).(ClientAccessRightResponseOutput)
}

type ContactDetails struct {
	CompanyName   string   `pulumi:"companyName"`
	ContactPerson string   `pulumi:"contactPerson"`
	EmailList     []string `pulumi:"emailList"`
	Phone         string   `pulumi:"phone"`
}





type ContactDetailsInput interface {
	pulumi.Input

	ToContactDetailsOutput() ContactDetailsOutput
	ToContactDetailsOutputWithContext(context.Context) ContactDetailsOutput
}

type ContactDetailsArgs struct {
	CompanyName   pulumi.StringInput      `pulumi:"companyName"`
	ContactPerson pulumi.StringInput      `pulumi:"contactPerson"`
	EmailList     pulumi.StringArrayInput `pulumi:"emailList"`
	Phone         pulumi.StringInput      `pulumi:"phone"`
}

func (ContactDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetails)(nil)).Elem()
}

func (i ContactDetailsArgs) ToContactDetailsOutput() ContactDetailsOutput {
	return i.ToContactDetailsOutputWithContext(context.Background())
}

func (i ContactDetailsArgs) ToContactDetailsOutputWithContext(ctx context.Context) ContactDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsOutput)
}

type ContactDetailsOutput struct{ *pulumi.OutputState }

func (ContactDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetails)(nil)).Elem()
}

func (o ContactDetailsOutput) ToContactDetailsOutput() ContactDetailsOutput {
	return o
}

func (o ContactDetailsOutput) ToContactDetailsOutputWithContext(ctx context.Context) ContactDetailsOutput {
	return o
}

func (o ContactDetailsOutput) CompanyName() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetails) string { return v.CompanyName }).(pulumi.StringOutput)
}

func (o ContactDetailsOutput) ContactPerson() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetails) string { return v.ContactPerson }).(pulumi.StringOutput)
}

func (o ContactDetailsOutput) EmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContactDetails) []string { return v.EmailList }).(pulumi.StringArrayOutput)
}

func (o ContactDetailsOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetails) string { return v.Phone }).(pulumi.StringOutput)
}

type ContactDetailsResponse struct {
	CompanyName   string   `pulumi:"companyName"`
	ContactPerson string   `pulumi:"contactPerson"`
	EmailList     []string `pulumi:"emailList"`
	Phone         string   `pulumi:"phone"`
}

type ContactDetailsResponseOutput struct{ *pulumi.OutputState }

func (ContactDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetailsResponse)(nil)).Elem()
}

func (o ContactDetailsResponseOutput) ToContactDetailsResponseOutput() ContactDetailsResponseOutput {
	return o
}

func (o ContactDetailsResponseOutput) ToContactDetailsResponseOutputWithContext(ctx context.Context) ContactDetailsResponseOutput {
	return o
}

func (o ContactDetailsResponseOutput) CompanyName() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetailsResponse) string { return v.CompanyName }).(pulumi.StringOutput)
}

func (o ContactDetailsResponseOutput) ContactPerson() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetailsResponse) string { return v.ContactPerson }).(pulumi.StringOutput)
}

func (o ContactDetailsResponseOutput) EmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContactDetailsResponse) []string { return v.EmailList }).(pulumi.StringArrayOutput)
}

func (o ContactDetailsResponseOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetailsResponse) string { return v.Phone }).(pulumi.StringOutput)
}

type FileSourceInfo struct {
	ShareId string `pulumi:"shareId"`
}





type FileSourceInfoInput interface {
	pulumi.Input

	ToFileSourceInfoOutput() FileSourceInfoOutput
	ToFileSourceInfoOutputWithContext(context.Context) FileSourceInfoOutput
}

type FileSourceInfoArgs struct {
	ShareId pulumi.StringInput `pulumi:"shareId"`
}

func (FileSourceInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSourceInfo)(nil)).Elem()
}

func (i FileSourceInfoArgs) ToFileSourceInfoOutput() FileSourceInfoOutput {
	return i.ToFileSourceInfoOutputWithContext(context.Background())
}

func (i FileSourceInfoArgs) ToFileSourceInfoOutputWithContext(ctx context.Context) FileSourceInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSourceInfoOutput)
}

type FileSourceInfoOutput struct{ *pulumi.OutputState }

func (FileSourceInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSourceInfo)(nil)).Elem()
}

func (o FileSourceInfoOutput) ToFileSourceInfoOutput() FileSourceInfoOutput {
	return o
}

func (o FileSourceInfoOutput) ToFileSourceInfoOutputWithContext(ctx context.Context) FileSourceInfoOutput {
	return o
}

func (o FileSourceInfoOutput) ShareId() pulumi.StringOutput {
	return o.ApplyT(func(v FileSourceInfo) string { return v.ShareId }).(pulumi.StringOutput)
}

type FileSourceInfoResponse struct {
	ShareId string `pulumi:"shareId"`
}

type FileSourceInfoResponseOutput struct{ *pulumi.OutputState }

func (FileSourceInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSourceInfoResponse)(nil)).Elem()
}

func (o FileSourceInfoResponseOutput) ToFileSourceInfoResponseOutput() FileSourceInfoResponseOutput {
	return o
}

func (o FileSourceInfoResponseOutput) ToFileSourceInfoResponseOutputWithContext(ctx context.Context) FileSourceInfoResponseOutput {
	return o
}

func (o FileSourceInfoResponseOutput) ShareId() pulumi.StringOutput {
	return o.ApplyT(func(v FileSourceInfoResponse) string { return v.ShareId }).(pulumi.StringOutput)
}

type IoTDeviceInfo struct {
	Authentication *Authentication `pulumi:"authentication"`
	DeviceId       string          `pulumi:"deviceId"`
	IoTHostHub     string          `pulumi:"ioTHostHub"`
	IoTHostHubId   *string         `pulumi:"ioTHostHubId"`
}





type IoTDeviceInfoInput interface {
	pulumi.Input

	ToIoTDeviceInfoOutput() IoTDeviceInfoOutput
	ToIoTDeviceInfoOutputWithContext(context.Context) IoTDeviceInfoOutput
}

type IoTDeviceInfoArgs struct {
	Authentication AuthenticationPtrInput `pulumi:"authentication"`
	DeviceId       pulumi.StringInput     `pulumi:"deviceId"`
	IoTHostHub     pulumi.StringInput     `pulumi:"ioTHostHub"`
	IoTHostHubId   pulumi.StringPtrInput  `pulumi:"ioTHostHubId"`
}

func (IoTDeviceInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTDeviceInfo)(nil)).Elem()
}

func (i IoTDeviceInfoArgs) ToIoTDeviceInfoOutput() IoTDeviceInfoOutput {
	return i.ToIoTDeviceInfoOutputWithContext(context.Background())
}

func (i IoTDeviceInfoArgs) ToIoTDeviceInfoOutputWithContext(ctx context.Context) IoTDeviceInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTDeviceInfoOutput)
}

type IoTDeviceInfoOutput struct{ *pulumi.OutputState }

func (IoTDeviceInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTDeviceInfo)(nil)).Elem()
}

func (o IoTDeviceInfoOutput) ToIoTDeviceInfoOutput() IoTDeviceInfoOutput {
	return o
}

func (o IoTDeviceInfoOutput) ToIoTDeviceInfoOutputWithContext(ctx context.Context) IoTDeviceInfoOutput {
	return o
}

func (o IoTDeviceInfoOutput) Authentication() AuthenticationPtrOutput {
	return o.ApplyT(func(v IoTDeviceInfo) *Authentication { return v.Authentication }).(AuthenticationPtrOutput)
}

func (o IoTDeviceInfoOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v IoTDeviceInfo) string { return v.DeviceId }).(pulumi.StringOutput)
}

func (o IoTDeviceInfoOutput) IoTHostHub() pulumi.StringOutput {
	return o.ApplyT(func(v IoTDeviceInfo) string { return v.IoTHostHub }).(pulumi.StringOutput)
}

func (o IoTDeviceInfoOutput) IoTHostHubId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IoTDeviceInfo) *string { return v.IoTHostHubId }).(pulumi.StringPtrOutput)
}

type IoTDeviceInfoResponse struct {
	Authentication *AuthenticationResponse `pulumi:"authentication"`
	DeviceId       string                  `pulumi:"deviceId"`
	IoTHostHub     string                  `pulumi:"ioTHostHub"`
	IoTHostHubId   *string                 `pulumi:"ioTHostHubId"`
}

type IoTDeviceInfoResponseOutput struct{ *pulumi.OutputState }

func (IoTDeviceInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTDeviceInfoResponse)(nil)).Elem()
}

func (o IoTDeviceInfoResponseOutput) ToIoTDeviceInfoResponseOutput() IoTDeviceInfoResponseOutput {
	return o
}

func (o IoTDeviceInfoResponseOutput) ToIoTDeviceInfoResponseOutputWithContext(ctx context.Context) IoTDeviceInfoResponseOutput {
	return o
}

func (o IoTDeviceInfoResponseOutput) Authentication() AuthenticationResponsePtrOutput {
	return o.ApplyT(func(v IoTDeviceInfoResponse) *AuthenticationResponse { return v.Authentication }).(AuthenticationResponsePtrOutput)
}

func (o IoTDeviceInfoResponseOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v IoTDeviceInfoResponse) string { return v.DeviceId }).(pulumi.StringOutput)
}

func (o IoTDeviceInfoResponseOutput) IoTHostHub() pulumi.StringOutput {
	return o.ApplyT(func(v IoTDeviceInfoResponse) string { return v.IoTHostHub }).(pulumi.StringOutput)
}

func (o IoTDeviceInfoResponseOutput) IoTHostHubId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IoTDeviceInfoResponse) *string { return v.IoTHostHubId }).(pulumi.StringPtrOutput)
}

type MountPointMap struct {
	ShareId string `pulumi:"shareId"`
}





type MountPointMapInput interface {
	pulumi.Input

	ToMountPointMapOutput() MountPointMapOutput
	ToMountPointMapOutputWithContext(context.Context) MountPointMapOutput
}

type MountPointMapArgs struct {
	ShareId pulumi.StringInput `pulumi:"shareId"`
}

func (MountPointMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MountPointMap)(nil)).Elem()
}

func (i MountPointMapArgs) ToMountPointMapOutput() MountPointMapOutput {
	return i.ToMountPointMapOutputWithContext(context.Background())
}

func (i MountPointMapArgs) ToMountPointMapOutputWithContext(ctx context.Context) MountPointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointMapOutput)
}





type MountPointMapArrayInput interface {
	pulumi.Input

	ToMountPointMapArrayOutput() MountPointMapArrayOutput
	ToMountPointMapArrayOutputWithContext(context.Context) MountPointMapArrayOutput
}

type MountPointMapArray []MountPointMapInput

func (MountPointMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MountPointMap)(nil)).Elem()
}

func (i MountPointMapArray) ToMountPointMapArrayOutput() MountPointMapArrayOutput {
	return i.ToMountPointMapArrayOutputWithContext(context.Background())
}

func (i MountPointMapArray) ToMountPointMapArrayOutputWithContext(ctx context.Context) MountPointMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointMapArrayOutput)
}

type MountPointMapOutput struct{ *pulumi.OutputState }

func (MountPointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MountPointMap)(nil)).Elem()
}

func (o MountPointMapOutput) ToMountPointMapOutput() MountPointMapOutput {
	return o
}

func (o MountPointMapOutput) ToMountPointMapOutputWithContext(ctx context.Context) MountPointMapOutput {
	return o
}

func (o MountPointMapOutput) ShareId() pulumi.StringOutput {
	return o.ApplyT(func(v MountPointMap) string { return v.ShareId }).(pulumi.StringOutput)
}

type MountPointMapArrayOutput struct{ *pulumi.OutputState }

func (MountPointMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MountPointMap)(nil)).Elem()
}

func (o MountPointMapArrayOutput) ToMountPointMapArrayOutput() MountPointMapArrayOutput {
	return o
}

func (o MountPointMapArrayOutput) ToMountPointMapArrayOutputWithContext(ctx context.Context) MountPointMapArrayOutput {
	return o
}

func (o MountPointMapArrayOutput) Index(i pulumi.IntInput) MountPointMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MountPointMap {
		return vs[0].([]MountPointMap)[vs[1].(int)]
	}).(MountPointMapOutput)
}

type MountPointMapResponse struct {
	MountPoint string `pulumi:"mountPoint"`
	RoleId     string `pulumi:"roleId"`
	RoleType   string `pulumi:"roleType"`
	ShareId    string `pulumi:"shareId"`
}

type MountPointMapResponseOutput struct{ *pulumi.OutputState }

func (MountPointMapResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MountPointMapResponse)(nil)).Elem()
}

func (o MountPointMapResponseOutput) ToMountPointMapResponseOutput() MountPointMapResponseOutput {
	return o
}

func (o MountPointMapResponseOutput) ToMountPointMapResponseOutputWithContext(ctx context.Context) MountPointMapResponseOutput {
	return o
}

func (o MountPointMapResponseOutput) MountPoint() pulumi.StringOutput {
	return o.ApplyT(func(v MountPointMapResponse) string { return v.MountPoint }).(pulumi.StringOutput)
}

func (o MountPointMapResponseOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v MountPointMapResponse) string { return v.RoleId }).(pulumi.StringOutput)
}

func (o MountPointMapResponseOutput) RoleType() pulumi.StringOutput {
	return o.ApplyT(func(v MountPointMapResponse) string { return v.RoleType }).(pulumi.StringOutput)
}

func (o MountPointMapResponseOutput) ShareId() pulumi.StringOutput {
	return o.ApplyT(func(v MountPointMapResponse) string { return v.ShareId }).(pulumi.StringOutput)
}

type MountPointMapResponseArrayOutput struct{ *pulumi.OutputState }

func (MountPointMapResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MountPointMapResponse)(nil)).Elem()
}

func (o MountPointMapResponseArrayOutput) ToMountPointMapResponseArrayOutput() MountPointMapResponseArrayOutput {
	return o
}

func (o MountPointMapResponseArrayOutput) ToMountPointMapResponseArrayOutputWithContext(ctx context.Context) MountPointMapResponseArrayOutput {
	return o
}

func (o MountPointMapResponseArrayOutput) Index(i pulumi.IntInput) MountPointMapResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MountPointMapResponse {
		return vs[0].([]MountPointMapResponse)[vs[1].(int)]
	}).(MountPointMapResponseOutput)
}

type OrderStatus struct {
	Comments *string `pulumi:"comments"`
	Status   string  `pulumi:"status"`
}





type OrderStatusInput interface {
	pulumi.Input

	ToOrderStatusOutput() OrderStatusOutput
	ToOrderStatusOutputWithContext(context.Context) OrderStatusOutput
}

type OrderStatusArgs struct {
	Comments pulumi.StringPtrInput `pulumi:"comments"`
	Status   pulumi.StringInput    `pulumi:"status"`
}

func (OrderStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderStatus)(nil)).Elem()
}

func (i OrderStatusArgs) ToOrderStatusOutput() OrderStatusOutput {
	return i.ToOrderStatusOutputWithContext(context.Background())
}

func (i OrderStatusArgs) ToOrderStatusOutputWithContext(ctx context.Context) OrderStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderStatusOutput)
}

func (i OrderStatusArgs) ToOrderStatusPtrOutput() OrderStatusPtrOutput {
	return i.ToOrderStatusPtrOutputWithContext(context.Background())
}

func (i OrderStatusArgs) ToOrderStatusPtrOutputWithContext(ctx context.Context) OrderStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderStatusOutput).ToOrderStatusPtrOutputWithContext(ctx)
}









type OrderStatusPtrInput interface {
	pulumi.Input

	ToOrderStatusPtrOutput() OrderStatusPtrOutput
	ToOrderStatusPtrOutputWithContext(context.Context) OrderStatusPtrOutput
}

type orderStatusPtrType OrderStatusArgs

func OrderStatusPtr(v *OrderStatusArgs) OrderStatusPtrInput {
	return (*orderStatusPtrType)(v)
}

func (*orderStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderStatus)(nil)).Elem()
}

func (i *orderStatusPtrType) ToOrderStatusPtrOutput() OrderStatusPtrOutput {
	return i.ToOrderStatusPtrOutputWithContext(context.Background())
}

func (i *orderStatusPtrType) ToOrderStatusPtrOutputWithContext(ctx context.Context) OrderStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderStatusPtrOutput)
}

type OrderStatusOutput struct{ *pulumi.OutputState }

func (OrderStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderStatus)(nil)).Elem()
}

func (o OrderStatusOutput) ToOrderStatusOutput() OrderStatusOutput {
	return o
}

func (o OrderStatusOutput) ToOrderStatusOutputWithContext(ctx context.Context) OrderStatusOutput {
	return o
}

func (o OrderStatusOutput) ToOrderStatusPtrOutput() OrderStatusPtrOutput {
	return o.ToOrderStatusPtrOutputWithContext(context.Background())
}

func (o OrderStatusOutput) ToOrderStatusPtrOutputWithContext(ctx context.Context) OrderStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrderStatus) *OrderStatus {
		return &v
	}).(OrderStatusPtrOutput)
}

func (o OrderStatusOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrderStatus) *string { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o OrderStatusOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v OrderStatus) string { return v.Status }).(pulumi.StringOutput)
}

type OrderStatusPtrOutput struct{ *pulumi.OutputState }

func (OrderStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderStatus)(nil)).Elem()
}

func (o OrderStatusPtrOutput) ToOrderStatusPtrOutput() OrderStatusPtrOutput {
	return o
}

func (o OrderStatusPtrOutput) ToOrderStatusPtrOutputWithContext(ctx context.Context) OrderStatusPtrOutput {
	return o
}

func (o OrderStatusPtrOutput) Elem() OrderStatusOutput {
	return o.ApplyT(func(v *OrderStatus) OrderStatus {
		if v != nil {
			return *v
		}
		var ret OrderStatus
		return ret
	}).(OrderStatusOutput)
}

func (o OrderStatusPtrOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderStatus) *string {
		if v == nil {
			return nil
		}
		return v.Comments
	}).(pulumi.StringPtrOutput)
}

func (o OrderStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type OrderStatusResponse struct {
	Comments       *string `pulumi:"comments"`
	Status         string  `pulumi:"status"`
	UpdateDateTime string  `pulumi:"updateDateTime"`
}

type OrderStatusResponseOutput struct{ *pulumi.OutputState }

func (OrderStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderStatusResponse)(nil)).Elem()
}

func (o OrderStatusResponseOutput) ToOrderStatusResponseOutput() OrderStatusResponseOutput {
	return o
}

func (o OrderStatusResponseOutput) ToOrderStatusResponseOutputWithContext(ctx context.Context) OrderStatusResponseOutput {
	return o
}

func (o OrderStatusResponseOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrderStatusResponse) *string { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o OrderStatusResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v OrderStatusResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o OrderStatusResponseOutput) UpdateDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v OrderStatusResponse) string { return v.UpdateDateTime }).(pulumi.StringOutput)
}

type OrderStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (OrderStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderStatusResponse)(nil)).Elem()
}

func (o OrderStatusResponsePtrOutput) ToOrderStatusResponsePtrOutput() OrderStatusResponsePtrOutput {
	return o
}

func (o OrderStatusResponsePtrOutput) ToOrderStatusResponsePtrOutputWithContext(ctx context.Context) OrderStatusResponsePtrOutput {
	return o
}

func (o OrderStatusResponsePtrOutput) Elem() OrderStatusResponseOutput {
	return o.ApplyT(func(v *OrderStatusResponse) OrderStatusResponse {
		if v != nil {
			return *v
		}
		var ret OrderStatusResponse
		return ret
	}).(OrderStatusResponseOutput)
}

func (o OrderStatusResponsePtrOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Comments
	}).(pulumi.StringPtrOutput)
}

func (o OrderStatusResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o OrderStatusResponsePtrOutput) UpdateDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrderStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpdateDateTime
	}).(pulumi.StringPtrOutput)
}

type OrderStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (OrderStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrderStatusResponse)(nil)).Elem()
}

func (o OrderStatusResponseArrayOutput) ToOrderStatusResponseArrayOutput() OrderStatusResponseArrayOutput {
	return o
}

func (o OrderStatusResponseArrayOutput) ToOrderStatusResponseArrayOutputWithContext(ctx context.Context) OrderStatusResponseArrayOutput {
	return o
}

func (o OrderStatusResponseArrayOutput) Index(i pulumi.IntInput) OrderStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OrderStatusResponse {
		return vs[0].([]OrderStatusResponse)[vs[1].(int)]
	}).(OrderStatusResponseOutput)
}

type PeriodicTimerSourceInfo struct {
	Schedule  string  `pulumi:"schedule"`
	StartTime string  `pulumi:"startTime"`
	Topic     *string `pulumi:"topic"`
}





type PeriodicTimerSourceInfoInput interface {
	pulumi.Input

	ToPeriodicTimerSourceInfoOutput() PeriodicTimerSourceInfoOutput
	ToPeriodicTimerSourceInfoOutputWithContext(context.Context) PeriodicTimerSourceInfoOutput
}

type PeriodicTimerSourceInfoArgs struct {
	Schedule  pulumi.StringInput    `pulumi:"schedule"`
	StartTime pulumi.StringInput    `pulumi:"startTime"`
	Topic     pulumi.StringPtrInput `pulumi:"topic"`
}

func (PeriodicTimerSourceInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeriodicTimerSourceInfo)(nil)).Elem()
}

func (i PeriodicTimerSourceInfoArgs) ToPeriodicTimerSourceInfoOutput() PeriodicTimerSourceInfoOutput {
	return i.ToPeriodicTimerSourceInfoOutputWithContext(context.Background())
}

func (i PeriodicTimerSourceInfoArgs) ToPeriodicTimerSourceInfoOutputWithContext(ctx context.Context) PeriodicTimerSourceInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicTimerSourceInfoOutput)
}

type PeriodicTimerSourceInfoOutput struct{ *pulumi.OutputState }

func (PeriodicTimerSourceInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeriodicTimerSourceInfo)(nil)).Elem()
}

func (o PeriodicTimerSourceInfoOutput) ToPeriodicTimerSourceInfoOutput() PeriodicTimerSourceInfoOutput {
	return o
}

func (o PeriodicTimerSourceInfoOutput) ToPeriodicTimerSourceInfoOutputWithContext(ctx context.Context) PeriodicTimerSourceInfoOutput {
	return o
}

func (o PeriodicTimerSourceInfoOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v PeriodicTimerSourceInfo) string { return v.Schedule }).(pulumi.StringOutput)
}

func (o PeriodicTimerSourceInfoOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v PeriodicTimerSourceInfo) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o PeriodicTimerSourceInfoOutput) Topic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeriodicTimerSourceInfo) *string { return v.Topic }).(pulumi.StringPtrOutput)
}

type PeriodicTimerSourceInfoResponse struct {
	Schedule  string  `pulumi:"schedule"`
	StartTime string  `pulumi:"startTime"`
	Topic     *string `pulumi:"topic"`
}

type PeriodicTimerSourceInfoResponseOutput struct{ *pulumi.OutputState }

func (PeriodicTimerSourceInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeriodicTimerSourceInfoResponse)(nil)).Elem()
}

func (o PeriodicTimerSourceInfoResponseOutput) ToPeriodicTimerSourceInfoResponseOutput() PeriodicTimerSourceInfoResponseOutput {
	return o
}

func (o PeriodicTimerSourceInfoResponseOutput) ToPeriodicTimerSourceInfoResponseOutputWithContext(ctx context.Context) PeriodicTimerSourceInfoResponseOutput {
	return o
}

func (o PeriodicTimerSourceInfoResponseOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v PeriodicTimerSourceInfoResponse) string { return v.Schedule }).(pulumi.StringOutput)
}

func (o PeriodicTimerSourceInfoResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v PeriodicTimerSourceInfoResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o PeriodicTimerSourceInfoResponseOutput) Topic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeriodicTimerSourceInfoResponse) *string { return v.Topic }).(pulumi.StringPtrOutput)
}

type RefreshDetails struct {
	ErrorManifestFile                *string `pulumi:"errorManifestFile"`
	InProgressRefreshJobId           *string `pulumi:"inProgressRefreshJobId"`
	LastCompletedRefreshJobTimeInUTC *string `pulumi:"lastCompletedRefreshJobTimeInUTC"`
	LastJob                          *string `pulumi:"lastJob"`
}





type RefreshDetailsInput interface {
	pulumi.Input

	ToRefreshDetailsOutput() RefreshDetailsOutput
	ToRefreshDetailsOutputWithContext(context.Context) RefreshDetailsOutput
}

type RefreshDetailsArgs struct {
	ErrorManifestFile                pulumi.StringPtrInput `pulumi:"errorManifestFile"`
	InProgressRefreshJobId           pulumi.StringPtrInput `pulumi:"inProgressRefreshJobId"`
	LastCompletedRefreshJobTimeInUTC pulumi.StringPtrInput `pulumi:"lastCompletedRefreshJobTimeInUTC"`
	LastJob                          pulumi.StringPtrInput `pulumi:"lastJob"`
}

func (RefreshDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RefreshDetails)(nil)).Elem()
}

func (i RefreshDetailsArgs) ToRefreshDetailsOutput() RefreshDetailsOutput {
	return i.ToRefreshDetailsOutputWithContext(context.Background())
}

func (i RefreshDetailsArgs) ToRefreshDetailsOutputWithContext(ctx context.Context) RefreshDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RefreshDetailsOutput)
}

func (i RefreshDetailsArgs) ToRefreshDetailsPtrOutput() RefreshDetailsPtrOutput {
	return i.ToRefreshDetailsPtrOutputWithContext(context.Background())
}

func (i RefreshDetailsArgs) ToRefreshDetailsPtrOutputWithContext(ctx context.Context) RefreshDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RefreshDetailsOutput).ToRefreshDetailsPtrOutputWithContext(ctx)
}









type RefreshDetailsPtrInput interface {
	pulumi.Input

	ToRefreshDetailsPtrOutput() RefreshDetailsPtrOutput
	ToRefreshDetailsPtrOutputWithContext(context.Context) RefreshDetailsPtrOutput
}

type refreshDetailsPtrType RefreshDetailsArgs

func RefreshDetailsPtr(v *RefreshDetailsArgs) RefreshDetailsPtrInput {
	return (*refreshDetailsPtrType)(v)
}

func (*refreshDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RefreshDetails)(nil)).Elem()
}

func (i *refreshDetailsPtrType) ToRefreshDetailsPtrOutput() RefreshDetailsPtrOutput {
	return i.ToRefreshDetailsPtrOutputWithContext(context.Background())
}

func (i *refreshDetailsPtrType) ToRefreshDetailsPtrOutputWithContext(ctx context.Context) RefreshDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RefreshDetailsPtrOutput)
}

type RefreshDetailsOutput struct{ *pulumi.OutputState }

func (RefreshDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RefreshDetails)(nil)).Elem()
}

func (o RefreshDetailsOutput) ToRefreshDetailsOutput() RefreshDetailsOutput {
	return o
}

func (o RefreshDetailsOutput) ToRefreshDetailsOutputWithContext(ctx context.Context) RefreshDetailsOutput {
	return o
}

func (o RefreshDetailsOutput) ToRefreshDetailsPtrOutput() RefreshDetailsPtrOutput {
	return o.ToRefreshDetailsPtrOutputWithContext(context.Background())
}

func (o RefreshDetailsOutput) ToRefreshDetailsPtrOutputWithContext(ctx context.Context) RefreshDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RefreshDetails) *RefreshDetails {
		return &v
	}).(RefreshDetailsPtrOutput)
}

func (o RefreshDetailsOutput) ErrorManifestFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RefreshDetails) *string { return v.ErrorManifestFile }).(pulumi.StringPtrOutput)
}

func (o RefreshDetailsOutput) InProgressRefreshJobId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RefreshDetails) *string { return v.InProgressRefreshJobId }).(pulumi.StringPtrOutput)
}

func (o RefreshDetailsOutput) LastCompletedRefreshJobTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RefreshDetails) *string { return v.LastCompletedRefreshJobTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o RefreshDetailsOutput) LastJob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RefreshDetails) *string { return v.LastJob }).(pulumi.StringPtrOutput)
}

type RefreshDetailsPtrOutput struct{ *pulumi.OutputState }

func (RefreshDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RefreshDetails)(nil)).Elem()
}

func (o RefreshDetailsPtrOutput) ToRefreshDetailsPtrOutput() RefreshDetailsPtrOutput {
	return o
}

func (o RefreshDetailsPtrOutput) ToRefreshDetailsPtrOutputWithContext(ctx context.Context) RefreshDetailsPtrOutput {
	return o
}

func (o RefreshDetailsPtrOutput) Elem() RefreshDetailsOutput {
	return o.ApplyT(func(v *RefreshDetails) RefreshDetails {
		if v != nil {
			return *v
		}
		var ret RefreshDetails
		return ret
	}).(RefreshDetailsOutput)
}

func (o RefreshDetailsPtrOutput) ErrorManifestFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RefreshDetails) *string {
		if v == nil {
			return nil
		}
		return v.ErrorManifestFile
	}).(pulumi.StringPtrOutput)
}

func (o RefreshDetailsPtrOutput) InProgressRefreshJobId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RefreshDetails) *string {
		if v == nil {
			return nil
		}
		return v.InProgressRefreshJobId
	}).(pulumi.StringPtrOutput)
}

func (o RefreshDetailsPtrOutput) LastCompletedRefreshJobTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RefreshDetails) *string {
		if v == nil {
			return nil
		}
		return v.LastCompletedRefreshJobTimeInUTC
	}).(pulumi.StringPtrOutput)
}

func (o RefreshDetailsPtrOutput) LastJob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RefreshDetails) *string {
		if v == nil {
			return nil
		}
		return v.LastJob
	}).(pulumi.StringPtrOutput)
}

type RefreshDetailsResponse struct {
	ErrorManifestFile                *string `pulumi:"errorManifestFile"`
	InProgressRefreshJobId           *string `pulumi:"inProgressRefreshJobId"`
	LastCompletedRefreshJobTimeInUTC *string `pulumi:"lastCompletedRefreshJobTimeInUTC"`
	LastJob                          *string `pulumi:"lastJob"`
}

type RefreshDetailsResponseOutput struct{ *pulumi.OutputState }

func (RefreshDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RefreshDetailsResponse)(nil)).Elem()
}

func (o RefreshDetailsResponseOutput) ToRefreshDetailsResponseOutput() RefreshDetailsResponseOutput {
	return o
}

func (o RefreshDetailsResponseOutput) ToRefreshDetailsResponseOutputWithContext(ctx context.Context) RefreshDetailsResponseOutput {
	return o
}

func (o RefreshDetailsResponseOutput) ErrorManifestFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RefreshDetailsResponse) *string { return v.ErrorManifestFile }).(pulumi.StringPtrOutput)
}

func (o RefreshDetailsResponseOutput) InProgressRefreshJobId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RefreshDetailsResponse) *string { return v.InProgressRefreshJobId }).(pulumi.StringPtrOutput)
}

func (o RefreshDetailsResponseOutput) LastCompletedRefreshJobTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RefreshDetailsResponse) *string { return v.LastCompletedRefreshJobTimeInUTC }).(pulumi.StringPtrOutput)
}

func (o RefreshDetailsResponseOutput) LastJob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RefreshDetailsResponse) *string { return v.LastJob }).(pulumi.StringPtrOutput)
}

type RefreshDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (RefreshDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RefreshDetailsResponse)(nil)).Elem()
}

func (o RefreshDetailsResponsePtrOutput) ToRefreshDetailsResponsePtrOutput() RefreshDetailsResponsePtrOutput {
	return o
}

func (o RefreshDetailsResponsePtrOutput) ToRefreshDetailsResponsePtrOutputWithContext(ctx context.Context) RefreshDetailsResponsePtrOutput {
	return o
}

func (o RefreshDetailsResponsePtrOutput) Elem() RefreshDetailsResponseOutput {
	return o.ApplyT(func(v *RefreshDetailsResponse) RefreshDetailsResponse {
		if v != nil {
			return *v
		}
		var ret RefreshDetailsResponse
		return ret
	}).(RefreshDetailsResponseOutput)
}

func (o RefreshDetailsResponsePtrOutput) ErrorManifestFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RefreshDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ErrorManifestFile
	}).(pulumi.StringPtrOutput)
}

func (o RefreshDetailsResponsePtrOutput) InProgressRefreshJobId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RefreshDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InProgressRefreshJobId
	}).(pulumi.StringPtrOutput)
}

func (o RefreshDetailsResponsePtrOutput) LastCompletedRefreshJobTimeInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RefreshDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastCompletedRefreshJobTimeInUTC
	}).(pulumi.StringPtrOutput)
}

func (o RefreshDetailsResponsePtrOutput) LastJob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RefreshDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastJob
	}).(pulumi.StringPtrOutput)
}

type RoleSinkInfo struct {
	RoleId string `pulumi:"roleId"`
}





type RoleSinkInfoInput interface {
	pulumi.Input

	ToRoleSinkInfoOutput() RoleSinkInfoOutput
	ToRoleSinkInfoOutputWithContext(context.Context) RoleSinkInfoOutput
}

type RoleSinkInfoArgs struct {
	RoleId pulumi.StringInput `pulumi:"roleId"`
}

func (RoleSinkInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleSinkInfo)(nil)).Elem()
}

func (i RoleSinkInfoArgs) ToRoleSinkInfoOutput() RoleSinkInfoOutput {
	return i.ToRoleSinkInfoOutputWithContext(context.Background())
}

func (i RoleSinkInfoArgs) ToRoleSinkInfoOutputWithContext(ctx context.Context) RoleSinkInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleSinkInfoOutput)
}

type RoleSinkInfoOutput struct{ *pulumi.OutputState }

func (RoleSinkInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleSinkInfo)(nil)).Elem()
}

func (o RoleSinkInfoOutput) ToRoleSinkInfoOutput() RoleSinkInfoOutput {
	return o
}

func (o RoleSinkInfoOutput) ToRoleSinkInfoOutputWithContext(ctx context.Context) RoleSinkInfoOutput {
	return o
}

func (o RoleSinkInfoOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v RoleSinkInfo) string { return v.RoleId }).(pulumi.StringOutput)
}

type RoleSinkInfoResponse struct {
	RoleId string `pulumi:"roleId"`
}

type RoleSinkInfoResponseOutput struct{ *pulumi.OutputState }

func (RoleSinkInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleSinkInfoResponse)(nil)).Elem()
}

func (o RoleSinkInfoResponseOutput) ToRoleSinkInfoResponseOutput() RoleSinkInfoResponseOutput {
	return o
}

func (o RoleSinkInfoResponseOutput) ToRoleSinkInfoResponseOutputWithContext(ctx context.Context) RoleSinkInfoResponseOutput {
	return o
}

func (o RoleSinkInfoResponseOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v RoleSinkInfoResponse) string { return v.RoleId }).(pulumi.StringOutput)
}

type ShareAccessRight struct {
	AccessType string `pulumi:"accessType"`
	ShareId    string `pulumi:"shareId"`
}





type ShareAccessRightInput interface {
	pulumi.Input

	ToShareAccessRightOutput() ShareAccessRightOutput
	ToShareAccessRightOutputWithContext(context.Context) ShareAccessRightOutput
}

type ShareAccessRightArgs struct {
	AccessType pulumi.StringInput `pulumi:"accessType"`
	ShareId    pulumi.StringInput `pulumi:"shareId"`
}

func (ShareAccessRightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareAccessRight)(nil)).Elem()
}

func (i ShareAccessRightArgs) ToShareAccessRightOutput() ShareAccessRightOutput {
	return i.ToShareAccessRightOutputWithContext(context.Background())
}

func (i ShareAccessRightArgs) ToShareAccessRightOutputWithContext(ctx context.Context) ShareAccessRightOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareAccessRightOutput)
}





type ShareAccessRightArrayInput interface {
	pulumi.Input

	ToShareAccessRightArrayOutput() ShareAccessRightArrayOutput
	ToShareAccessRightArrayOutputWithContext(context.Context) ShareAccessRightArrayOutput
}

type ShareAccessRightArray []ShareAccessRightInput

func (ShareAccessRightArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareAccessRight)(nil)).Elem()
}

func (i ShareAccessRightArray) ToShareAccessRightArrayOutput() ShareAccessRightArrayOutput {
	return i.ToShareAccessRightArrayOutputWithContext(context.Background())
}

func (i ShareAccessRightArray) ToShareAccessRightArrayOutputWithContext(ctx context.Context) ShareAccessRightArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareAccessRightArrayOutput)
}

type ShareAccessRightOutput struct{ *pulumi.OutputState }

func (ShareAccessRightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareAccessRight)(nil)).Elem()
}

func (o ShareAccessRightOutput) ToShareAccessRightOutput() ShareAccessRightOutput {
	return o
}

func (o ShareAccessRightOutput) ToShareAccessRightOutputWithContext(ctx context.Context) ShareAccessRightOutput {
	return o
}

func (o ShareAccessRightOutput) AccessType() pulumi.StringOutput {
	return o.ApplyT(func(v ShareAccessRight) string { return v.AccessType }).(pulumi.StringOutput)
}

func (o ShareAccessRightOutput) ShareId() pulumi.StringOutput {
	return o.ApplyT(func(v ShareAccessRight) string { return v.ShareId }).(pulumi.StringOutput)
}

type ShareAccessRightArrayOutput struct{ *pulumi.OutputState }

func (ShareAccessRightArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareAccessRight)(nil)).Elem()
}

func (o ShareAccessRightArrayOutput) ToShareAccessRightArrayOutput() ShareAccessRightArrayOutput {
	return o
}

func (o ShareAccessRightArrayOutput) ToShareAccessRightArrayOutputWithContext(ctx context.Context) ShareAccessRightArrayOutput {
	return o
}

func (o ShareAccessRightArrayOutput) Index(i pulumi.IntInput) ShareAccessRightOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ShareAccessRight {
		return vs[0].([]ShareAccessRight)[vs[1].(int)]
	}).(ShareAccessRightOutput)
}

type ShareAccessRightResponse struct {
	AccessType string `pulumi:"accessType"`
	ShareId    string `pulumi:"shareId"`
}

type ShareAccessRightResponseOutput struct{ *pulumi.OutputState }

func (ShareAccessRightResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareAccessRightResponse)(nil)).Elem()
}

func (o ShareAccessRightResponseOutput) ToShareAccessRightResponseOutput() ShareAccessRightResponseOutput {
	return o
}

func (o ShareAccessRightResponseOutput) ToShareAccessRightResponseOutputWithContext(ctx context.Context) ShareAccessRightResponseOutput {
	return o
}

func (o ShareAccessRightResponseOutput) AccessType() pulumi.StringOutput {
	return o.ApplyT(func(v ShareAccessRightResponse) string { return v.AccessType }).(pulumi.StringOutput)
}

func (o ShareAccessRightResponseOutput) ShareId() pulumi.StringOutput {
	return o.ApplyT(func(v ShareAccessRightResponse) string { return v.ShareId }).(pulumi.StringOutput)
}

type ShareAccessRightResponseArrayOutput struct{ *pulumi.OutputState }

func (ShareAccessRightResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareAccessRightResponse)(nil)).Elem()
}

func (o ShareAccessRightResponseArrayOutput) ToShareAccessRightResponseArrayOutput() ShareAccessRightResponseArrayOutput {
	return o
}

func (o ShareAccessRightResponseArrayOutput) ToShareAccessRightResponseArrayOutputWithContext(ctx context.Context) ShareAccessRightResponseArrayOutput {
	return o
}

func (o ShareAccessRightResponseArrayOutput) Index(i pulumi.IntInput) ShareAccessRightResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ShareAccessRightResponse {
		return vs[0].([]ShareAccessRightResponse)[vs[1].(int)]
	}).(ShareAccessRightResponseOutput)
}

type Sku struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SymmetricKey struct {
	ConnectionString *AsymmetricEncryptedSecret `pulumi:"connectionString"`
}





type SymmetricKeyInput interface {
	pulumi.Input

	ToSymmetricKeyOutput() SymmetricKeyOutput
	ToSymmetricKeyOutputWithContext(context.Context) SymmetricKeyOutput
}

type SymmetricKeyArgs struct {
	ConnectionString AsymmetricEncryptedSecretPtrInput `pulumi:"connectionString"`
}

func (SymmetricKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SymmetricKey)(nil)).Elem()
}

func (i SymmetricKeyArgs) ToSymmetricKeyOutput() SymmetricKeyOutput {
	return i.ToSymmetricKeyOutputWithContext(context.Background())
}

func (i SymmetricKeyArgs) ToSymmetricKeyOutputWithContext(ctx context.Context) SymmetricKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SymmetricKeyOutput)
}

func (i SymmetricKeyArgs) ToSymmetricKeyPtrOutput() SymmetricKeyPtrOutput {
	return i.ToSymmetricKeyPtrOutputWithContext(context.Background())
}

func (i SymmetricKeyArgs) ToSymmetricKeyPtrOutputWithContext(ctx context.Context) SymmetricKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SymmetricKeyOutput).ToSymmetricKeyPtrOutputWithContext(ctx)
}









type SymmetricKeyPtrInput interface {
	pulumi.Input

	ToSymmetricKeyPtrOutput() SymmetricKeyPtrOutput
	ToSymmetricKeyPtrOutputWithContext(context.Context) SymmetricKeyPtrOutput
}

type symmetricKeyPtrType SymmetricKeyArgs

func SymmetricKeyPtr(v *SymmetricKeyArgs) SymmetricKeyPtrInput {
	return (*symmetricKeyPtrType)(v)
}

func (*symmetricKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SymmetricKey)(nil)).Elem()
}

func (i *symmetricKeyPtrType) ToSymmetricKeyPtrOutput() SymmetricKeyPtrOutput {
	return i.ToSymmetricKeyPtrOutputWithContext(context.Background())
}

func (i *symmetricKeyPtrType) ToSymmetricKeyPtrOutputWithContext(ctx context.Context) SymmetricKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SymmetricKeyPtrOutput)
}

type SymmetricKeyOutput struct{ *pulumi.OutputState }

func (SymmetricKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SymmetricKey)(nil)).Elem()
}

func (o SymmetricKeyOutput) ToSymmetricKeyOutput() SymmetricKeyOutput {
	return o
}

func (o SymmetricKeyOutput) ToSymmetricKeyOutputWithContext(ctx context.Context) SymmetricKeyOutput {
	return o
}

func (o SymmetricKeyOutput) ToSymmetricKeyPtrOutput() SymmetricKeyPtrOutput {
	return o.ToSymmetricKeyPtrOutputWithContext(context.Background())
}

func (o SymmetricKeyOutput) ToSymmetricKeyPtrOutputWithContext(ctx context.Context) SymmetricKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SymmetricKey) *SymmetricKey {
		return &v
	}).(SymmetricKeyPtrOutput)
}

func (o SymmetricKeyOutput) ConnectionString() AsymmetricEncryptedSecretPtrOutput {
	return o.ApplyT(func(v SymmetricKey) *AsymmetricEncryptedSecret { return v.ConnectionString }).(AsymmetricEncryptedSecretPtrOutput)
}

type SymmetricKeyPtrOutput struct{ *pulumi.OutputState }

func (SymmetricKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SymmetricKey)(nil)).Elem()
}

func (o SymmetricKeyPtrOutput) ToSymmetricKeyPtrOutput() SymmetricKeyPtrOutput {
	return o
}

func (o SymmetricKeyPtrOutput) ToSymmetricKeyPtrOutputWithContext(ctx context.Context) SymmetricKeyPtrOutput {
	return o
}

func (o SymmetricKeyPtrOutput) Elem() SymmetricKeyOutput {
	return o.ApplyT(func(v *SymmetricKey) SymmetricKey {
		if v != nil {
			return *v
		}
		var ret SymmetricKey
		return ret
	}).(SymmetricKeyOutput)
}

func (o SymmetricKeyPtrOutput) ConnectionString() AsymmetricEncryptedSecretPtrOutput {
	return o.ApplyT(func(v *SymmetricKey) *AsymmetricEncryptedSecret {
		if v == nil {
			return nil
		}
		return v.ConnectionString
	}).(AsymmetricEncryptedSecretPtrOutput)
}

type SymmetricKeyResponse struct {
	ConnectionString *AsymmetricEncryptedSecretResponse `pulumi:"connectionString"`
}

type SymmetricKeyResponseOutput struct{ *pulumi.OutputState }

func (SymmetricKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SymmetricKeyResponse)(nil)).Elem()
}

func (o SymmetricKeyResponseOutput) ToSymmetricKeyResponseOutput() SymmetricKeyResponseOutput {
	return o
}

func (o SymmetricKeyResponseOutput) ToSymmetricKeyResponseOutputWithContext(ctx context.Context) SymmetricKeyResponseOutput {
	return o
}

func (o SymmetricKeyResponseOutput) ConnectionString() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v SymmetricKeyResponse) *AsymmetricEncryptedSecretResponse { return v.ConnectionString }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

type SymmetricKeyResponsePtrOutput struct{ *pulumi.OutputState }

func (SymmetricKeyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SymmetricKeyResponse)(nil)).Elem()
}

func (o SymmetricKeyResponsePtrOutput) ToSymmetricKeyResponsePtrOutput() SymmetricKeyResponsePtrOutput {
	return o
}

func (o SymmetricKeyResponsePtrOutput) ToSymmetricKeyResponsePtrOutputWithContext(ctx context.Context) SymmetricKeyResponsePtrOutput {
	return o
}

func (o SymmetricKeyResponsePtrOutput) Elem() SymmetricKeyResponseOutput {
	return o.ApplyT(func(v *SymmetricKeyResponse) SymmetricKeyResponse {
		if v != nil {
			return *v
		}
		var ret SymmetricKeyResponse
		return ret
	}).(SymmetricKeyResponseOutput)
}

func (o SymmetricKeyResponsePtrOutput) ConnectionString() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v *SymmetricKeyResponse) *AsymmetricEncryptedSecretResponse {
		if v == nil {
			return nil
		}
		return v.ConnectionString
	}).(AsymmetricEncryptedSecretResponsePtrOutput)
}

type TrackingInfoResponse struct {
	CarrierName  *string `pulumi:"carrierName"`
	SerialNumber *string `pulumi:"serialNumber"`
	TrackingId   *string `pulumi:"trackingId"`
	TrackingUrl  *string `pulumi:"trackingUrl"`
}

type TrackingInfoResponseOutput struct{ *pulumi.OutputState }

func (TrackingInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackingInfoResponse)(nil)).Elem()
}

func (o TrackingInfoResponseOutput) ToTrackingInfoResponseOutput() TrackingInfoResponseOutput {
	return o
}

func (o TrackingInfoResponseOutput) ToTrackingInfoResponseOutputWithContext(ctx context.Context) TrackingInfoResponseOutput {
	return o
}

func (o TrackingInfoResponseOutput) CarrierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackingInfoResponse) *string { return v.CarrierName }).(pulumi.StringPtrOutput)
}

func (o TrackingInfoResponseOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackingInfoResponse) *string { return v.SerialNumber }).(pulumi.StringPtrOutput)
}

func (o TrackingInfoResponseOutput) TrackingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackingInfoResponse) *string { return v.TrackingId }).(pulumi.StringPtrOutput)
}

func (o TrackingInfoResponseOutput) TrackingUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackingInfoResponse) *string { return v.TrackingUrl }).(pulumi.StringPtrOutput)
}

type TrackingInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (TrackingInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackingInfoResponse)(nil)).Elem()
}

func (o TrackingInfoResponseArrayOutput) ToTrackingInfoResponseArrayOutput() TrackingInfoResponseArrayOutput {
	return o
}

func (o TrackingInfoResponseArrayOutput) ToTrackingInfoResponseArrayOutputWithContext(ctx context.Context) TrackingInfoResponseArrayOutput {
	return o
}

func (o TrackingInfoResponseArrayOutput) Index(i pulumi.IntInput) TrackingInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrackingInfoResponse {
		return vs[0].([]TrackingInfoResponse)[vs[1].(int)]
	}).(TrackingInfoResponseOutput)
}

type UserAccessRight struct {
	AccessType string `pulumi:"accessType"`
	UserId     string `pulumi:"userId"`
}





type UserAccessRightInput interface {
	pulumi.Input

	ToUserAccessRightOutput() UserAccessRightOutput
	ToUserAccessRightOutputWithContext(context.Context) UserAccessRightOutput
}

type UserAccessRightArgs struct {
	AccessType pulumi.StringInput `pulumi:"accessType"`
	UserId     pulumi.StringInput `pulumi:"userId"`
}

func (UserAccessRightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccessRight)(nil)).Elem()
}

func (i UserAccessRightArgs) ToUserAccessRightOutput() UserAccessRightOutput {
	return i.ToUserAccessRightOutputWithContext(context.Background())
}

func (i UserAccessRightArgs) ToUserAccessRightOutputWithContext(ctx context.Context) UserAccessRightOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccessRightOutput)
}





type UserAccessRightArrayInput interface {
	pulumi.Input

	ToUserAccessRightArrayOutput() UserAccessRightArrayOutput
	ToUserAccessRightArrayOutputWithContext(context.Context) UserAccessRightArrayOutput
}

type UserAccessRightArray []UserAccessRightInput

func (UserAccessRightArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserAccessRight)(nil)).Elem()
}

func (i UserAccessRightArray) ToUserAccessRightArrayOutput() UserAccessRightArrayOutput {
	return i.ToUserAccessRightArrayOutputWithContext(context.Background())
}

func (i UserAccessRightArray) ToUserAccessRightArrayOutputWithContext(ctx context.Context) UserAccessRightArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccessRightArrayOutput)
}

type UserAccessRightOutput struct{ *pulumi.OutputState }

func (UserAccessRightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccessRight)(nil)).Elem()
}

func (o UserAccessRightOutput) ToUserAccessRightOutput() UserAccessRightOutput {
	return o
}

func (o UserAccessRightOutput) ToUserAccessRightOutputWithContext(ctx context.Context) UserAccessRightOutput {
	return o
}

func (o UserAccessRightOutput) AccessType() pulumi.StringOutput {
	return o.ApplyT(func(v UserAccessRight) string { return v.AccessType }).(pulumi.StringOutput)
}

func (o UserAccessRightOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAccessRight) string { return v.UserId }).(pulumi.StringOutput)
}

type UserAccessRightArrayOutput struct{ *pulumi.OutputState }

func (UserAccessRightArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserAccessRight)(nil)).Elem()
}

func (o UserAccessRightArrayOutput) ToUserAccessRightArrayOutput() UserAccessRightArrayOutput {
	return o
}

func (o UserAccessRightArrayOutput) ToUserAccessRightArrayOutputWithContext(ctx context.Context) UserAccessRightArrayOutput {
	return o
}

func (o UserAccessRightArrayOutput) Index(i pulumi.IntInput) UserAccessRightOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserAccessRight {
		return vs[0].([]UserAccessRight)[vs[1].(int)]
	}).(UserAccessRightOutput)
}

type UserAccessRightResponse struct {
	AccessType string `pulumi:"accessType"`
	UserId     string `pulumi:"userId"`
}

type UserAccessRightResponseOutput struct{ *pulumi.OutputState }

func (UserAccessRightResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccessRightResponse)(nil)).Elem()
}

func (o UserAccessRightResponseOutput) ToUserAccessRightResponseOutput() UserAccessRightResponseOutput {
	return o
}

func (o UserAccessRightResponseOutput) ToUserAccessRightResponseOutputWithContext(ctx context.Context) UserAccessRightResponseOutput {
	return o
}

func (o UserAccessRightResponseOutput) AccessType() pulumi.StringOutput {
	return o.ApplyT(func(v UserAccessRightResponse) string { return v.AccessType }).(pulumi.StringOutput)
}

func (o UserAccessRightResponseOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAccessRightResponse) string { return v.UserId }).(pulumi.StringOutput)
}

type UserAccessRightResponseArrayOutput struct{ *pulumi.OutputState }

func (UserAccessRightResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserAccessRightResponse)(nil)).Elem()
}

func (o UserAccessRightResponseArrayOutput) ToUserAccessRightResponseArrayOutput() UserAccessRightResponseArrayOutput {
	return o
}

func (o UserAccessRightResponseArrayOutput) ToUserAccessRightResponseArrayOutputWithContext(ctx context.Context) UserAccessRightResponseArrayOutput {
	return o
}

func (o UserAccessRightResponseArrayOutput) Index(i pulumi.IntInput) UserAccessRightResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserAccessRightResponse {
		return vs[0].([]UserAccessRightResponse)[vs[1].(int)]
	}).(UserAccessRightResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AddressOutput{})
	pulumi.RegisterOutputType(AddressResponseOutput{})
	pulumi.RegisterOutputType(AsymmetricEncryptedSecretOutput{})
	pulumi.RegisterOutputType(AsymmetricEncryptedSecretPtrOutput{})
	pulumi.RegisterOutputType(AsymmetricEncryptedSecretResponseOutput{})
	pulumi.RegisterOutputType(AsymmetricEncryptedSecretResponsePtrOutput{})
	pulumi.RegisterOutputType(AuthenticationOutput{})
	pulumi.RegisterOutputType(AuthenticationPtrOutput{})
	pulumi.RegisterOutputType(AuthenticationResponseOutput{})
	pulumi.RegisterOutputType(AuthenticationResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureContainerInfoOutput{})
	pulumi.RegisterOutputType(AzureContainerInfoPtrOutput{})
	pulumi.RegisterOutputType(AzureContainerInfoResponseOutput{})
	pulumi.RegisterOutputType(AzureContainerInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ClientAccessRightOutput{})
	pulumi.RegisterOutputType(ClientAccessRightArrayOutput{})
	pulumi.RegisterOutputType(ClientAccessRightResponseOutput{})
	pulumi.RegisterOutputType(ClientAccessRightResponseArrayOutput{})
	pulumi.RegisterOutputType(ContactDetailsOutput{})
	pulumi.RegisterOutputType(ContactDetailsResponseOutput{})
	pulumi.RegisterOutputType(FileSourceInfoOutput{})
	pulumi.RegisterOutputType(FileSourceInfoResponseOutput{})
	pulumi.RegisterOutputType(IoTDeviceInfoOutput{})
	pulumi.RegisterOutputType(IoTDeviceInfoResponseOutput{})
	pulumi.RegisterOutputType(MountPointMapOutput{})
	pulumi.RegisterOutputType(MountPointMapArrayOutput{})
	pulumi.RegisterOutputType(MountPointMapResponseOutput{})
	pulumi.RegisterOutputType(MountPointMapResponseArrayOutput{})
	pulumi.RegisterOutputType(OrderStatusOutput{})
	pulumi.RegisterOutputType(OrderStatusPtrOutput{})
	pulumi.RegisterOutputType(OrderStatusResponseOutput{})
	pulumi.RegisterOutputType(OrderStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(OrderStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(PeriodicTimerSourceInfoOutput{})
	pulumi.RegisterOutputType(PeriodicTimerSourceInfoResponseOutput{})
	pulumi.RegisterOutputType(RefreshDetailsOutput{})
	pulumi.RegisterOutputType(RefreshDetailsPtrOutput{})
	pulumi.RegisterOutputType(RefreshDetailsResponseOutput{})
	pulumi.RegisterOutputType(RefreshDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(RoleSinkInfoOutput{})
	pulumi.RegisterOutputType(RoleSinkInfoResponseOutput{})
	pulumi.RegisterOutputType(ShareAccessRightOutput{})
	pulumi.RegisterOutputType(ShareAccessRightArrayOutput{})
	pulumi.RegisterOutputType(ShareAccessRightResponseOutput{})
	pulumi.RegisterOutputType(ShareAccessRightResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SymmetricKeyOutput{})
	pulumi.RegisterOutputType(SymmetricKeyPtrOutput{})
	pulumi.RegisterOutputType(SymmetricKeyResponseOutput{})
	pulumi.RegisterOutputType(SymmetricKeyResponsePtrOutput{})
	pulumi.RegisterOutputType(TrackingInfoResponseOutput{})
	pulumi.RegisterOutputType(TrackingInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAccessRightOutput{})
	pulumi.RegisterOutputType(UserAccessRightArrayOutput{})
	pulumi.RegisterOutputType(UserAccessRightResponseOutput{})
	pulumi.RegisterOutputType(UserAccessRightResponseArrayOutput{})
}

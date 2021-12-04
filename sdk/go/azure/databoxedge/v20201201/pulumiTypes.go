


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Address struct {
	AddressLine1 *string `pulumi:"addressLine1"`
	AddressLine2 *string `pulumi:"addressLine2"`
	AddressLine3 *string `pulumi:"addressLine3"`
	City         *string `pulumi:"city"`
	Country      string  `pulumi:"country"`
	PostalCode   *string `pulumi:"postalCode"`
	State        *string `pulumi:"state"`
}





type AddressInput interface {
	pulumi.Input

	ToAddressOutput() AddressOutput
	ToAddressOutputWithContext(context.Context) AddressOutput
}

type AddressArgs struct {
	AddressLine1 pulumi.StringPtrInput `pulumi:"addressLine1"`
	AddressLine2 pulumi.StringPtrInput `pulumi:"addressLine2"`
	AddressLine3 pulumi.StringPtrInput `pulumi:"addressLine3"`
	City         pulumi.StringPtrInput `pulumi:"city"`
	Country      pulumi.StringInput    `pulumi:"country"`
	PostalCode   pulumi.StringPtrInput `pulumi:"postalCode"`
	State        pulumi.StringPtrInput `pulumi:"state"`
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

func (i AddressArgs) ToAddressPtrOutput() AddressPtrOutput {
	return i.ToAddressPtrOutputWithContext(context.Background())
}

func (i AddressArgs) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput).ToAddressPtrOutputWithContext(ctx)
}









type AddressPtrInput interface {
	pulumi.Input

	ToAddressPtrOutput() AddressPtrOutput
	ToAddressPtrOutputWithContext(context.Context) AddressPtrOutput
}

type addressPtrType AddressArgs

func AddressPtr(v *AddressArgs) AddressPtrInput {
	return (*addressPtrType)(v)
}

func (*addressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (i *addressPtrType) ToAddressPtrOutput() AddressPtrOutput {
	return i.ToAddressPtrOutputWithContext(context.Background())
}

func (i *addressPtrType) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPtrOutput)
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

func (o AddressOutput) ToAddressPtrOutput() AddressPtrOutput {
	return o.ToAddressPtrOutputWithContext(context.Background())
}

func (o AddressOutput) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Address) *Address {
		return &v
	}).(AddressPtrOutput)
}

func (o AddressOutput) AddressLine1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Address) *string { return v.AddressLine1 }).(pulumi.StringPtrOutput)
}

func (o AddressOutput) AddressLine2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Address) *string { return v.AddressLine2 }).(pulumi.StringPtrOutput)
}

func (o AddressOutput) AddressLine3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Address) *string { return v.AddressLine3 }).(pulumi.StringPtrOutput)
}

func (o AddressOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Address) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o AddressOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.Country }).(pulumi.StringOutput)
}

func (o AddressOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Address) *string { return v.PostalCode }).(pulumi.StringPtrOutput)
}

func (o AddressOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Address) *string { return v.State }).(pulumi.StringPtrOutput)
}

type AddressPtrOutput struct{ *pulumi.OutputState }

func (AddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (o AddressPtrOutput) ToAddressPtrOutput() AddressPtrOutput {
	return o
}

func (o AddressPtrOutput) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return o
}

func (o AddressPtrOutput) Elem() AddressOutput {
	return o.ApplyT(func(v *Address) Address {
		if v != nil {
			return *v
		}
		var ret Address
		return ret
	}).(AddressOutput)
}

func (o AddressPtrOutput) AddressLine1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return v.AddressLine1
	}).(pulumi.StringPtrOutput)
}

func (o AddressPtrOutput) AddressLine2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return v.AddressLine2
	}).(pulumi.StringPtrOutput)
}

func (o AddressPtrOutput) AddressLine3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return v.AddressLine3
	}).(pulumi.StringPtrOutput)
}

func (o AddressPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o AddressPtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return &v.Country
	}).(pulumi.StringPtrOutput)
}

func (o AddressPtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return v.PostalCode
	}).(pulumi.StringPtrOutput)
}

func (o AddressPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type AddressResponse struct {
	AddressLine1 *string `pulumi:"addressLine1"`
	AddressLine2 *string `pulumi:"addressLine2"`
	AddressLine3 *string `pulumi:"addressLine3"`
	City         *string `pulumi:"city"`
	Country      string  `pulumi:"country"`
	PostalCode   *string `pulumi:"postalCode"`
	State        *string `pulumi:"state"`
}





type AddressResponseInput interface {
	pulumi.Input

	ToAddressResponseOutput() AddressResponseOutput
	ToAddressResponseOutputWithContext(context.Context) AddressResponseOutput
}

type AddressResponseArgs struct {
	AddressLine1 pulumi.StringPtrInput `pulumi:"addressLine1"`
	AddressLine2 pulumi.StringPtrInput `pulumi:"addressLine2"`
	AddressLine3 pulumi.StringPtrInput `pulumi:"addressLine3"`
	City         pulumi.StringPtrInput `pulumi:"city"`
	Country      pulumi.StringInput    `pulumi:"country"`
	PostalCode   pulumi.StringPtrInput `pulumi:"postalCode"`
	State        pulumi.StringPtrInput `pulumi:"state"`
}

func (AddressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressResponse)(nil)).Elem()
}

func (i AddressResponseArgs) ToAddressResponseOutput() AddressResponseOutput {
	return i.ToAddressResponseOutputWithContext(context.Background())
}

func (i AddressResponseArgs) ToAddressResponseOutputWithContext(ctx context.Context) AddressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressResponseOutput)
}

func (i AddressResponseArgs) ToAddressResponsePtrOutput() AddressResponsePtrOutput {
	return i.ToAddressResponsePtrOutputWithContext(context.Background())
}

func (i AddressResponseArgs) ToAddressResponsePtrOutputWithContext(ctx context.Context) AddressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressResponseOutput).ToAddressResponsePtrOutputWithContext(ctx)
}









type AddressResponsePtrInput interface {
	pulumi.Input

	ToAddressResponsePtrOutput() AddressResponsePtrOutput
	ToAddressResponsePtrOutputWithContext(context.Context) AddressResponsePtrOutput
}

type addressResponsePtrType AddressResponseArgs

func AddressResponsePtr(v *AddressResponseArgs) AddressResponsePtrInput {
	return (*addressResponsePtrType)(v)
}

func (*addressResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressResponse)(nil)).Elem()
}

func (i *addressResponsePtrType) ToAddressResponsePtrOutput() AddressResponsePtrOutput {
	return i.ToAddressResponsePtrOutputWithContext(context.Background())
}

func (i *addressResponsePtrType) ToAddressResponsePtrOutputWithContext(ctx context.Context) AddressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressResponsePtrOutput)
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

func (o AddressResponseOutput) ToAddressResponsePtrOutput() AddressResponsePtrOutput {
	return o.ToAddressResponsePtrOutputWithContext(context.Background())
}

func (o AddressResponseOutput) ToAddressResponsePtrOutputWithContext(ctx context.Context) AddressResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddressResponse) *AddressResponse {
		return &v
	}).(AddressResponsePtrOutput)
}

func (o AddressResponseOutput) AddressLine1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressResponse) *string { return v.AddressLine1 }).(pulumi.StringPtrOutput)
}

func (o AddressResponseOutput) AddressLine2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressResponse) *string { return v.AddressLine2 }).(pulumi.StringPtrOutput)
}

func (o AddressResponseOutput) AddressLine3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressResponse) *string { return v.AddressLine3 }).(pulumi.StringPtrOutput)
}

func (o AddressResponseOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressResponse) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o AddressResponseOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v AddressResponse) string { return v.Country }).(pulumi.StringOutput)
}

func (o AddressResponseOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressResponse) *string { return v.PostalCode }).(pulumi.StringPtrOutput)
}

func (o AddressResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type AddressResponsePtrOutput struct{ *pulumi.OutputState }

func (AddressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressResponse)(nil)).Elem()
}

func (o AddressResponsePtrOutput) ToAddressResponsePtrOutput() AddressResponsePtrOutput {
	return o
}

func (o AddressResponsePtrOutput) ToAddressResponsePtrOutputWithContext(ctx context.Context) AddressResponsePtrOutput {
	return o
}

func (o AddressResponsePtrOutput) Elem() AddressResponseOutput {
	return o.ApplyT(func(v *AddressResponse) AddressResponse {
		if v != nil {
			return *v
		}
		var ret AddressResponse
		return ret
	}).(AddressResponseOutput)
}

func (o AddressResponsePtrOutput) AddressLine1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.AddressLine1
	}).(pulumi.StringPtrOutput)
}

func (o AddressResponsePtrOutput) AddressLine2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.AddressLine2
	}).(pulumi.StringPtrOutput)
}

func (o AddressResponsePtrOutput) AddressLine3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.AddressLine3
	}).(pulumi.StringPtrOutput)
}

func (o AddressResponsePtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o AddressResponsePtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Country
	}).(pulumi.StringPtrOutput)
}

func (o AddressResponsePtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.PostalCode
	}).(pulumi.StringPtrOutput)
}

func (o AddressResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
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





type AsymmetricEncryptedSecretResponseInput interface {
	pulumi.Input

	ToAsymmetricEncryptedSecretResponseOutput() AsymmetricEncryptedSecretResponseOutput
	ToAsymmetricEncryptedSecretResponseOutputWithContext(context.Context) AsymmetricEncryptedSecretResponseOutput
}

type AsymmetricEncryptedSecretResponseArgs struct {
	EncryptionAlgorithm      pulumi.StringInput    `pulumi:"encryptionAlgorithm"`
	EncryptionCertThumbprint pulumi.StringPtrInput `pulumi:"encryptionCertThumbprint"`
	Value                    pulumi.StringInput    `pulumi:"value"`
}

func (AsymmetricEncryptedSecretResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AsymmetricEncryptedSecretResponse)(nil)).Elem()
}

func (i AsymmetricEncryptedSecretResponseArgs) ToAsymmetricEncryptedSecretResponseOutput() AsymmetricEncryptedSecretResponseOutput {
	return i.ToAsymmetricEncryptedSecretResponseOutputWithContext(context.Background())
}

func (i AsymmetricEncryptedSecretResponseArgs) ToAsymmetricEncryptedSecretResponseOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretResponseOutput)
}

func (i AsymmetricEncryptedSecretResponseArgs) ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput {
	return i.ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(context.Background())
}

func (i AsymmetricEncryptedSecretResponseArgs) ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretResponseOutput).ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx)
}









type AsymmetricEncryptedSecretResponsePtrInput interface {
	pulumi.Input

	ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput
	ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(context.Context) AsymmetricEncryptedSecretResponsePtrOutput
}

type asymmetricEncryptedSecretResponsePtrType AsymmetricEncryptedSecretResponseArgs

func AsymmetricEncryptedSecretResponsePtr(v *AsymmetricEncryptedSecretResponseArgs) AsymmetricEncryptedSecretResponsePtrInput {
	return (*asymmetricEncryptedSecretResponsePtrType)(v)
}

func (*asymmetricEncryptedSecretResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AsymmetricEncryptedSecretResponse)(nil)).Elem()
}

func (i *asymmetricEncryptedSecretResponsePtrType) ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput {
	return i.ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(context.Background())
}

func (i *asymmetricEncryptedSecretResponsePtrType) ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretResponsePtrOutput)
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

func (o AsymmetricEncryptedSecretResponseOutput) ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(context.Background())
}

func (o AsymmetricEncryptedSecretResponseOutput) ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AsymmetricEncryptedSecretResponse) *AsymmetricEncryptedSecretResponse {
		return &v
	}).(AsymmetricEncryptedSecretResponsePtrOutput)
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





type AuthenticationResponseInput interface {
	pulumi.Input

	ToAuthenticationResponseOutput() AuthenticationResponseOutput
	ToAuthenticationResponseOutputWithContext(context.Context) AuthenticationResponseOutput
}

type AuthenticationResponseArgs struct {
	SymmetricKey SymmetricKeyResponsePtrInput `pulumi:"symmetricKey"`
}

func (AuthenticationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationResponse)(nil)).Elem()
}

func (i AuthenticationResponseArgs) ToAuthenticationResponseOutput() AuthenticationResponseOutput {
	return i.ToAuthenticationResponseOutputWithContext(context.Background())
}

func (i AuthenticationResponseArgs) ToAuthenticationResponseOutputWithContext(ctx context.Context) AuthenticationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationResponseOutput)
}

func (i AuthenticationResponseArgs) ToAuthenticationResponsePtrOutput() AuthenticationResponsePtrOutput {
	return i.ToAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (i AuthenticationResponseArgs) ToAuthenticationResponsePtrOutputWithContext(ctx context.Context) AuthenticationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationResponseOutput).ToAuthenticationResponsePtrOutputWithContext(ctx)
}









type AuthenticationResponsePtrInput interface {
	pulumi.Input

	ToAuthenticationResponsePtrOutput() AuthenticationResponsePtrOutput
	ToAuthenticationResponsePtrOutputWithContext(context.Context) AuthenticationResponsePtrOutput
}

type authenticationResponsePtrType AuthenticationResponseArgs

func AuthenticationResponsePtr(v *AuthenticationResponseArgs) AuthenticationResponsePtrInput {
	return (*authenticationResponsePtrType)(v)
}

func (*authenticationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationResponse)(nil)).Elem()
}

func (i *authenticationResponsePtrType) ToAuthenticationResponsePtrOutput() AuthenticationResponsePtrOutput {
	return i.ToAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (i *authenticationResponsePtrType) ToAuthenticationResponsePtrOutputWithContext(ctx context.Context) AuthenticationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationResponsePtrOutput)
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

func (o AuthenticationResponseOutput) ToAuthenticationResponsePtrOutput() AuthenticationResponsePtrOutput {
	return o.ToAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (o AuthenticationResponseOutput) ToAuthenticationResponsePtrOutputWithContext(ctx context.Context) AuthenticationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthenticationResponse) *AuthenticationResponse {
		return &v
	}).(AuthenticationResponsePtrOutput)
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





type AzureContainerInfoResponseInput interface {
	pulumi.Input

	ToAzureContainerInfoResponseOutput() AzureContainerInfoResponseOutput
	ToAzureContainerInfoResponseOutputWithContext(context.Context) AzureContainerInfoResponseOutput
}

type AzureContainerInfoResponseArgs struct {
	ContainerName              pulumi.StringInput `pulumi:"containerName"`
	DataFormat                 pulumi.StringInput `pulumi:"dataFormat"`
	StorageAccountCredentialId pulumi.StringInput `pulumi:"storageAccountCredentialId"`
}

func (AzureContainerInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureContainerInfoResponse)(nil)).Elem()
}

func (i AzureContainerInfoResponseArgs) ToAzureContainerInfoResponseOutput() AzureContainerInfoResponseOutput {
	return i.ToAzureContainerInfoResponseOutputWithContext(context.Background())
}

func (i AzureContainerInfoResponseArgs) ToAzureContainerInfoResponseOutputWithContext(ctx context.Context) AzureContainerInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureContainerInfoResponseOutput)
}

func (i AzureContainerInfoResponseArgs) ToAzureContainerInfoResponsePtrOutput() AzureContainerInfoResponsePtrOutput {
	return i.ToAzureContainerInfoResponsePtrOutputWithContext(context.Background())
}

func (i AzureContainerInfoResponseArgs) ToAzureContainerInfoResponsePtrOutputWithContext(ctx context.Context) AzureContainerInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureContainerInfoResponseOutput).ToAzureContainerInfoResponsePtrOutputWithContext(ctx)
}









type AzureContainerInfoResponsePtrInput interface {
	pulumi.Input

	ToAzureContainerInfoResponsePtrOutput() AzureContainerInfoResponsePtrOutput
	ToAzureContainerInfoResponsePtrOutputWithContext(context.Context) AzureContainerInfoResponsePtrOutput
}

type azureContainerInfoResponsePtrType AzureContainerInfoResponseArgs

func AzureContainerInfoResponsePtr(v *AzureContainerInfoResponseArgs) AzureContainerInfoResponsePtrInput {
	return (*azureContainerInfoResponsePtrType)(v)
}

func (*azureContainerInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureContainerInfoResponse)(nil)).Elem()
}

func (i *azureContainerInfoResponsePtrType) ToAzureContainerInfoResponsePtrOutput() AzureContainerInfoResponsePtrOutput {
	return i.ToAzureContainerInfoResponsePtrOutputWithContext(context.Background())
}

func (i *azureContainerInfoResponsePtrType) ToAzureContainerInfoResponsePtrOutputWithContext(ctx context.Context) AzureContainerInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureContainerInfoResponsePtrOutput)
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

func (o AzureContainerInfoResponseOutput) ToAzureContainerInfoResponsePtrOutput() AzureContainerInfoResponsePtrOutput {
	return o.ToAzureContainerInfoResponsePtrOutputWithContext(context.Background())
}

func (o AzureContainerInfoResponseOutput) ToAzureContainerInfoResponsePtrOutputWithContext(ctx context.Context) AzureContainerInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureContainerInfoResponse) *AzureContainerInfoResponse {
		return &v
	}).(AzureContainerInfoResponsePtrOutput)
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





type ClientAccessRightResponseInput interface {
	pulumi.Input

	ToClientAccessRightResponseOutput() ClientAccessRightResponseOutput
	ToClientAccessRightResponseOutputWithContext(context.Context) ClientAccessRightResponseOutput
}

type ClientAccessRightResponseArgs struct {
	AccessPermission pulumi.StringInput `pulumi:"accessPermission"`
	Client           pulumi.StringInput `pulumi:"client"`
}

func (ClientAccessRightResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAccessRightResponse)(nil)).Elem()
}

func (i ClientAccessRightResponseArgs) ToClientAccessRightResponseOutput() ClientAccessRightResponseOutput {
	return i.ToClientAccessRightResponseOutputWithContext(context.Background())
}

func (i ClientAccessRightResponseArgs) ToClientAccessRightResponseOutputWithContext(ctx context.Context) ClientAccessRightResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAccessRightResponseOutput)
}





type ClientAccessRightResponseArrayInput interface {
	pulumi.Input

	ToClientAccessRightResponseArrayOutput() ClientAccessRightResponseArrayOutput
	ToClientAccessRightResponseArrayOutputWithContext(context.Context) ClientAccessRightResponseArrayOutput
}

type ClientAccessRightResponseArray []ClientAccessRightResponseInput

func (ClientAccessRightResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientAccessRightResponse)(nil)).Elem()
}

func (i ClientAccessRightResponseArray) ToClientAccessRightResponseArrayOutput() ClientAccessRightResponseArrayOutput {
	return i.ToClientAccessRightResponseArrayOutputWithContext(context.Background())
}

func (i ClientAccessRightResponseArray) ToClientAccessRightResponseArrayOutputWithContext(ctx context.Context) ClientAccessRightResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAccessRightResponseArrayOutput)
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

type CniConfigResponse struct {
	PodSubnet     string `pulumi:"podSubnet"`
	ServiceSubnet string `pulumi:"serviceSubnet"`
	Type          string `pulumi:"type"`
	Version       string `pulumi:"version"`
}





type CniConfigResponseInput interface {
	pulumi.Input

	ToCniConfigResponseOutput() CniConfigResponseOutput
	ToCniConfigResponseOutputWithContext(context.Context) CniConfigResponseOutput
}

type CniConfigResponseArgs struct {
	PodSubnet     pulumi.StringInput `pulumi:"podSubnet"`
	ServiceSubnet pulumi.StringInput `pulumi:"serviceSubnet"`
	Type          pulumi.StringInput `pulumi:"type"`
	Version       pulumi.StringInput `pulumi:"version"`
}

func (CniConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CniConfigResponse)(nil)).Elem()
}

func (i CniConfigResponseArgs) ToCniConfigResponseOutput() CniConfigResponseOutput {
	return i.ToCniConfigResponseOutputWithContext(context.Background())
}

func (i CniConfigResponseArgs) ToCniConfigResponseOutputWithContext(ctx context.Context) CniConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CniConfigResponseOutput)
}

func (i CniConfigResponseArgs) ToCniConfigResponsePtrOutput() CniConfigResponsePtrOutput {
	return i.ToCniConfigResponsePtrOutputWithContext(context.Background())
}

func (i CniConfigResponseArgs) ToCniConfigResponsePtrOutputWithContext(ctx context.Context) CniConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CniConfigResponseOutput).ToCniConfigResponsePtrOutputWithContext(ctx)
}









type CniConfigResponsePtrInput interface {
	pulumi.Input

	ToCniConfigResponsePtrOutput() CniConfigResponsePtrOutput
	ToCniConfigResponsePtrOutputWithContext(context.Context) CniConfigResponsePtrOutput
}

type cniConfigResponsePtrType CniConfigResponseArgs

func CniConfigResponsePtr(v *CniConfigResponseArgs) CniConfigResponsePtrInput {
	return (*cniConfigResponsePtrType)(v)
}

func (*cniConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CniConfigResponse)(nil)).Elem()
}

func (i *cniConfigResponsePtrType) ToCniConfigResponsePtrOutput() CniConfigResponsePtrOutput {
	return i.ToCniConfigResponsePtrOutputWithContext(context.Background())
}

func (i *cniConfigResponsePtrType) ToCniConfigResponsePtrOutputWithContext(ctx context.Context) CniConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CniConfigResponsePtrOutput)
}

type CniConfigResponseOutput struct{ *pulumi.OutputState }

func (CniConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CniConfigResponse)(nil)).Elem()
}

func (o CniConfigResponseOutput) ToCniConfigResponseOutput() CniConfigResponseOutput {
	return o
}

func (o CniConfigResponseOutput) ToCniConfigResponseOutputWithContext(ctx context.Context) CniConfigResponseOutput {
	return o
}

func (o CniConfigResponseOutput) ToCniConfigResponsePtrOutput() CniConfigResponsePtrOutput {
	return o.ToCniConfigResponsePtrOutputWithContext(context.Background())
}

func (o CniConfigResponseOutput) ToCniConfigResponsePtrOutputWithContext(ctx context.Context) CniConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CniConfigResponse) *CniConfigResponse {
		return &v
	}).(CniConfigResponsePtrOutput)
}

func (o CniConfigResponseOutput) PodSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v CniConfigResponse) string { return v.PodSubnet }).(pulumi.StringOutput)
}

func (o CniConfigResponseOutput) ServiceSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v CniConfigResponse) string { return v.ServiceSubnet }).(pulumi.StringOutput)
}

func (o CniConfigResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v CniConfigResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o CniConfigResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v CniConfigResponse) string { return v.Version }).(pulumi.StringOutput)
}

type CniConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (CniConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CniConfigResponse)(nil)).Elem()
}

func (o CniConfigResponsePtrOutput) ToCniConfigResponsePtrOutput() CniConfigResponsePtrOutput {
	return o
}

func (o CniConfigResponsePtrOutput) ToCniConfigResponsePtrOutputWithContext(ctx context.Context) CniConfigResponsePtrOutput {
	return o
}

func (o CniConfigResponsePtrOutput) Elem() CniConfigResponseOutput {
	return o.ApplyT(func(v *CniConfigResponse) CniConfigResponse {
		if v != nil {
			return *v
		}
		var ret CniConfigResponse
		return ret
	}).(CniConfigResponseOutput)
}

func (o CniConfigResponsePtrOutput) PodSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CniConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PodSubnet
	}).(pulumi.StringPtrOutput)
}

func (o CniConfigResponsePtrOutput) ServiceSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CniConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServiceSubnet
	}).(pulumi.StringPtrOutput)
}

func (o CniConfigResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CniConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o CniConfigResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CniConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type ComputeResource struct {
	MemoryInGB     float64 `pulumi:"memoryInGB"`
	ProcessorCount int     `pulumi:"processorCount"`
}





type ComputeResourceInput interface {
	pulumi.Input

	ToComputeResourceOutput() ComputeResourceOutput
	ToComputeResourceOutputWithContext(context.Context) ComputeResourceOutput
}

type ComputeResourceArgs struct {
	MemoryInGB     pulumi.Float64Input `pulumi:"memoryInGB"`
	ProcessorCount pulumi.IntInput     `pulumi:"processorCount"`
}

func (ComputeResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeResource)(nil)).Elem()
}

func (i ComputeResourceArgs) ToComputeResourceOutput() ComputeResourceOutput {
	return i.ToComputeResourceOutputWithContext(context.Background())
}

func (i ComputeResourceArgs) ToComputeResourceOutputWithContext(ctx context.Context) ComputeResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeResourceOutput)
}

func (i ComputeResourceArgs) ToComputeResourcePtrOutput() ComputeResourcePtrOutput {
	return i.ToComputeResourcePtrOutputWithContext(context.Background())
}

func (i ComputeResourceArgs) ToComputeResourcePtrOutputWithContext(ctx context.Context) ComputeResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeResourceOutput).ToComputeResourcePtrOutputWithContext(ctx)
}









type ComputeResourcePtrInput interface {
	pulumi.Input

	ToComputeResourcePtrOutput() ComputeResourcePtrOutput
	ToComputeResourcePtrOutputWithContext(context.Context) ComputeResourcePtrOutput
}

type computeResourcePtrType ComputeResourceArgs

func ComputeResourcePtr(v *ComputeResourceArgs) ComputeResourcePtrInput {
	return (*computeResourcePtrType)(v)
}

func (*computeResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeResource)(nil)).Elem()
}

func (i *computeResourcePtrType) ToComputeResourcePtrOutput() ComputeResourcePtrOutput {
	return i.ToComputeResourcePtrOutputWithContext(context.Background())
}

func (i *computeResourcePtrType) ToComputeResourcePtrOutputWithContext(ctx context.Context) ComputeResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeResourcePtrOutput)
}

type ComputeResourceOutput struct{ *pulumi.OutputState }

func (ComputeResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeResource)(nil)).Elem()
}

func (o ComputeResourceOutput) ToComputeResourceOutput() ComputeResourceOutput {
	return o
}

func (o ComputeResourceOutput) ToComputeResourceOutputWithContext(ctx context.Context) ComputeResourceOutput {
	return o
}

func (o ComputeResourceOutput) ToComputeResourcePtrOutput() ComputeResourcePtrOutput {
	return o.ToComputeResourcePtrOutputWithContext(context.Background())
}

func (o ComputeResourceOutput) ToComputeResourcePtrOutputWithContext(ctx context.Context) ComputeResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeResource) *ComputeResource {
		return &v
	}).(ComputeResourcePtrOutput)
}

func (o ComputeResourceOutput) MemoryInGB() pulumi.Float64Output {
	return o.ApplyT(func(v ComputeResource) float64 { return v.MemoryInGB }).(pulumi.Float64Output)
}

func (o ComputeResourceOutput) ProcessorCount() pulumi.IntOutput {
	return o.ApplyT(func(v ComputeResource) int { return v.ProcessorCount }).(pulumi.IntOutput)
}

type ComputeResourcePtrOutput struct{ *pulumi.OutputState }

func (ComputeResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeResource)(nil)).Elem()
}

func (o ComputeResourcePtrOutput) ToComputeResourcePtrOutput() ComputeResourcePtrOutput {
	return o
}

func (o ComputeResourcePtrOutput) ToComputeResourcePtrOutputWithContext(ctx context.Context) ComputeResourcePtrOutput {
	return o
}

func (o ComputeResourcePtrOutput) Elem() ComputeResourceOutput {
	return o.ApplyT(func(v *ComputeResource) ComputeResource {
		if v != nil {
			return *v
		}
		var ret ComputeResource
		return ret
	}).(ComputeResourceOutput)
}

func (o ComputeResourcePtrOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ComputeResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.MemoryInGB
	}).(pulumi.Float64PtrOutput)
}

func (o ComputeResourcePtrOutput) ProcessorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ComputeResource) *int {
		if v == nil {
			return nil
		}
		return &v.ProcessorCount
	}).(pulumi.IntPtrOutput)
}

type ComputeResourceResponse struct {
	MemoryInGB     float64 `pulumi:"memoryInGB"`
	ProcessorCount int     `pulumi:"processorCount"`
}





type ComputeResourceResponseInput interface {
	pulumi.Input

	ToComputeResourceResponseOutput() ComputeResourceResponseOutput
	ToComputeResourceResponseOutputWithContext(context.Context) ComputeResourceResponseOutput
}

type ComputeResourceResponseArgs struct {
	MemoryInGB     pulumi.Float64Input `pulumi:"memoryInGB"`
	ProcessorCount pulumi.IntInput     `pulumi:"processorCount"`
}

func (ComputeResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeResourceResponse)(nil)).Elem()
}

func (i ComputeResourceResponseArgs) ToComputeResourceResponseOutput() ComputeResourceResponseOutput {
	return i.ToComputeResourceResponseOutputWithContext(context.Background())
}

func (i ComputeResourceResponseArgs) ToComputeResourceResponseOutputWithContext(ctx context.Context) ComputeResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeResourceResponseOutput)
}

func (i ComputeResourceResponseArgs) ToComputeResourceResponsePtrOutput() ComputeResourceResponsePtrOutput {
	return i.ToComputeResourceResponsePtrOutputWithContext(context.Background())
}

func (i ComputeResourceResponseArgs) ToComputeResourceResponsePtrOutputWithContext(ctx context.Context) ComputeResourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeResourceResponseOutput).ToComputeResourceResponsePtrOutputWithContext(ctx)
}









type ComputeResourceResponsePtrInput interface {
	pulumi.Input

	ToComputeResourceResponsePtrOutput() ComputeResourceResponsePtrOutput
	ToComputeResourceResponsePtrOutputWithContext(context.Context) ComputeResourceResponsePtrOutput
}

type computeResourceResponsePtrType ComputeResourceResponseArgs

func ComputeResourceResponsePtr(v *ComputeResourceResponseArgs) ComputeResourceResponsePtrInput {
	return (*computeResourceResponsePtrType)(v)
}

func (*computeResourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeResourceResponse)(nil)).Elem()
}

func (i *computeResourceResponsePtrType) ToComputeResourceResponsePtrOutput() ComputeResourceResponsePtrOutput {
	return i.ToComputeResourceResponsePtrOutputWithContext(context.Background())
}

func (i *computeResourceResponsePtrType) ToComputeResourceResponsePtrOutputWithContext(ctx context.Context) ComputeResourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeResourceResponsePtrOutput)
}

type ComputeResourceResponseOutput struct{ *pulumi.OutputState }

func (ComputeResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeResourceResponse)(nil)).Elem()
}

func (o ComputeResourceResponseOutput) ToComputeResourceResponseOutput() ComputeResourceResponseOutput {
	return o
}

func (o ComputeResourceResponseOutput) ToComputeResourceResponseOutputWithContext(ctx context.Context) ComputeResourceResponseOutput {
	return o
}

func (o ComputeResourceResponseOutput) ToComputeResourceResponsePtrOutput() ComputeResourceResponsePtrOutput {
	return o.ToComputeResourceResponsePtrOutputWithContext(context.Background())
}

func (o ComputeResourceResponseOutput) ToComputeResourceResponsePtrOutputWithContext(ctx context.Context) ComputeResourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeResourceResponse) *ComputeResourceResponse {
		return &v
	}).(ComputeResourceResponsePtrOutput)
}

func (o ComputeResourceResponseOutput) MemoryInGB() pulumi.Float64Output {
	return o.ApplyT(func(v ComputeResourceResponse) float64 { return v.MemoryInGB }).(pulumi.Float64Output)
}

func (o ComputeResourceResponseOutput) ProcessorCount() pulumi.IntOutput {
	return o.ApplyT(func(v ComputeResourceResponse) int { return v.ProcessorCount }).(pulumi.IntOutput)
}

type ComputeResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (ComputeResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeResourceResponse)(nil)).Elem()
}

func (o ComputeResourceResponsePtrOutput) ToComputeResourceResponsePtrOutput() ComputeResourceResponsePtrOutput {
	return o
}

func (o ComputeResourceResponsePtrOutput) ToComputeResourceResponsePtrOutputWithContext(ctx context.Context) ComputeResourceResponsePtrOutput {
	return o
}

func (o ComputeResourceResponsePtrOutput) Elem() ComputeResourceResponseOutput {
	return o.ApplyT(func(v *ComputeResourceResponse) ComputeResourceResponse {
		if v != nil {
			return *v
		}
		var ret ComputeResourceResponse
		return ret
	}).(ComputeResourceResponseOutput)
}

func (o ComputeResourceResponsePtrOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ComputeResourceResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.MemoryInGB
	}).(pulumi.Float64PtrOutput)
}

func (o ComputeResourceResponsePtrOutput) ProcessorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ComputeResourceResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ProcessorCount
	}).(pulumi.IntPtrOutput)
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

func (i ContactDetailsArgs) ToContactDetailsPtrOutput() ContactDetailsPtrOutput {
	return i.ToContactDetailsPtrOutputWithContext(context.Background())
}

func (i ContactDetailsArgs) ToContactDetailsPtrOutputWithContext(ctx context.Context) ContactDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsOutput).ToContactDetailsPtrOutputWithContext(ctx)
}









type ContactDetailsPtrInput interface {
	pulumi.Input

	ToContactDetailsPtrOutput() ContactDetailsPtrOutput
	ToContactDetailsPtrOutputWithContext(context.Context) ContactDetailsPtrOutput
}

type contactDetailsPtrType ContactDetailsArgs

func ContactDetailsPtr(v *ContactDetailsArgs) ContactDetailsPtrInput {
	return (*contactDetailsPtrType)(v)
}

func (*contactDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactDetails)(nil)).Elem()
}

func (i *contactDetailsPtrType) ToContactDetailsPtrOutput() ContactDetailsPtrOutput {
	return i.ToContactDetailsPtrOutputWithContext(context.Background())
}

func (i *contactDetailsPtrType) ToContactDetailsPtrOutputWithContext(ctx context.Context) ContactDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsPtrOutput)
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

func (o ContactDetailsOutput) ToContactDetailsPtrOutput() ContactDetailsPtrOutput {
	return o.ToContactDetailsPtrOutputWithContext(context.Background())
}

func (o ContactDetailsOutput) ToContactDetailsPtrOutputWithContext(ctx context.Context) ContactDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContactDetails) *ContactDetails {
		return &v
	}).(ContactDetailsPtrOutput)
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

type ContactDetailsPtrOutput struct{ *pulumi.OutputState }

func (ContactDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactDetails)(nil)).Elem()
}

func (o ContactDetailsPtrOutput) ToContactDetailsPtrOutput() ContactDetailsPtrOutput {
	return o
}

func (o ContactDetailsPtrOutput) ToContactDetailsPtrOutputWithContext(ctx context.Context) ContactDetailsPtrOutput {
	return o
}

func (o ContactDetailsPtrOutput) Elem() ContactDetailsOutput {
	return o.ApplyT(func(v *ContactDetails) ContactDetails {
		if v != nil {
			return *v
		}
		var ret ContactDetails
		return ret
	}).(ContactDetailsOutput)
}

func (o ContactDetailsPtrOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetails) *string {
		if v == nil {
			return nil
		}
		return &v.CompanyName
	}).(pulumi.StringPtrOutput)
}

func (o ContactDetailsPtrOutput) ContactPerson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetails) *string {
		if v == nil {
			return nil
		}
		return &v.ContactPerson
	}).(pulumi.StringPtrOutput)
}

func (o ContactDetailsPtrOutput) EmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContactDetails) []string {
		if v == nil {
			return nil
		}
		return v.EmailList
	}).(pulumi.StringArrayOutput)
}

func (o ContactDetailsPtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetails) *string {
		if v == nil {
			return nil
		}
		return &v.Phone
	}).(pulumi.StringPtrOutput)
}

type ContactDetailsResponse struct {
	CompanyName   string   `pulumi:"companyName"`
	ContactPerson string   `pulumi:"contactPerson"`
	EmailList     []string `pulumi:"emailList"`
	Phone         string   `pulumi:"phone"`
}





type ContactDetailsResponseInput interface {
	pulumi.Input

	ToContactDetailsResponseOutput() ContactDetailsResponseOutput
	ToContactDetailsResponseOutputWithContext(context.Context) ContactDetailsResponseOutput
}

type ContactDetailsResponseArgs struct {
	CompanyName   pulumi.StringInput      `pulumi:"companyName"`
	ContactPerson pulumi.StringInput      `pulumi:"contactPerson"`
	EmailList     pulumi.StringArrayInput `pulumi:"emailList"`
	Phone         pulumi.StringInput      `pulumi:"phone"`
}

func (ContactDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetailsResponse)(nil)).Elem()
}

func (i ContactDetailsResponseArgs) ToContactDetailsResponseOutput() ContactDetailsResponseOutput {
	return i.ToContactDetailsResponseOutputWithContext(context.Background())
}

func (i ContactDetailsResponseArgs) ToContactDetailsResponseOutputWithContext(ctx context.Context) ContactDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsResponseOutput)
}

func (i ContactDetailsResponseArgs) ToContactDetailsResponsePtrOutput() ContactDetailsResponsePtrOutput {
	return i.ToContactDetailsResponsePtrOutputWithContext(context.Background())
}

func (i ContactDetailsResponseArgs) ToContactDetailsResponsePtrOutputWithContext(ctx context.Context) ContactDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsResponseOutput).ToContactDetailsResponsePtrOutputWithContext(ctx)
}









type ContactDetailsResponsePtrInput interface {
	pulumi.Input

	ToContactDetailsResponsePtrOutput() ContactDetailsResponsePtrOutput
	ToContactDetailsResponsePtrOutputWithContext(context.Context) ContactDetailsResponsePtrOutput
}

type contactDetailsResponsePtrType ContactDetailsResponseArgs

func ContactDetailsResponsePtr(v *ContactDetailsResponseArgs) ContactDetailsResponsePtrInput {
	return (*contactDetailsResponsePtrType)(v)
}

func (*contactDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactDetailsResponse)(nil)).Elem()
}

func (i *contactDetailsResponsePtrType) ToContactDetailsResponsePtrOutput() ContactDetailsResponsePtrOutput {
	return i.ToContactDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *contactDetailsResponsePtrType) ToContactDetailsResponsePtrOutputWithContext(ctx context.Context) ContactDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsResponsePtrOutput)
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

func (o ContactDetailsResponseOutput) ToContactDetailsResponsePtrOutput() ContactDetailsResponsePtrOutput {
	return o.ToContactDetailsResponsePtrOutputWithContext(context.Background())
}

func (o ContactDetailsResponseOutput) ToContactDetailsResponsePtrOutputWithContext(ctx context.Context) ContactDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContactDetailsResponse) *ContactDetailsResponse {
		return &v
	}).(ContactDetailsResponsePtrOutput)
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

type ContactDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (ContactDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactDetailsResponse)(nil)).Elem()
}

func (o ContactDetailsResponsePtrOutput) ToContactDetailsResponsePtrOutput() ContactDetailsResponsePtrOutput {
	return o
}

func (o ContactDetailsResponsePtrOutput) ToContactDetailsResponsePtrOutputWithContext(ctx context.Context) ContactDetailsResponsePtrOutput {
	return o
}

func (o ContactDetailsResponsePtrOutput) Elem() ContactDetailsResponseOutput {
	return o.ApplyT(func(v *ContactDetailsResponse) ContactDetailsResponse {
		if v != nil {
			return *v
		}
		var ret ContactDetailsResponse
		return ret
	}).(ContactDetailsResponseOutput)
}

func (o ContactDetailsResponsePtrOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CompanyName
	}).(pulumi.StringPtrOutput)
}

func (o ContactDetailsResponsePtrOutput) ContactPerson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ContactPerson
	}).(pulumi.StringPtrOutput)
}

func (o ContactDetailsResponsePtrOutput) EmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContactDetailsResponse) []string {
		if v == nil {
			return nil
		}
		return v.EmailList
	}).(pulumi.StringArrayOutput)
}

func (o ContactDetailsResponsePtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Phone
	}).(pulumi.StringPtrOutput)
}

type DeviceSecretsResponse struct {
	BmcDefaultUserPassword                *SecretResponse `pulumi:"bmcDefaultUserPassword"`
	HcsDataVolumeBitLockerExternalKey     *SecretResponse `pulumi:"hcsDataVolumeBitLockerExternalKey"`
	HcsInternalVolumeBitLockerExternalKey *SecretResponse `pulumi:"hcsInternalVolumeBitLockerExternalKey"`
	RotateKeyForDataVolumeBitlocker       *SecretResponse `pulumi:"rotateKeyForDataVolumeBitlocker"`
	RotateKeysForSedDrivesSerialized      *SecretResponse `pulumi:"rotateKeysForSedDrivesSerialized"`
	SedEncryptionExternalKey              *SecretResponse `pulumi:"sedEncryptionExternalKey"`
	SedEncryptionExternalKeyId            *SecretResponse `pulumi:"sedEncryptionExternalKeyId"`
	SystemVolumeBitLockerRecoveryKey      *SecretResponse `pulumi:"systemVolumeBitLockerRecoveryKey"`
}





type DeviceSecretsResponseInput interface {
	pulumi.Input

	ToDeviceSecretsResponseOutput() DeviceSecretsResponseOutput
	ToDeviceSecretsResponseOutputWithContext(context.Context) DeviceSecretsResponseOutput
}

type DeviceSecretsResponseArgs struct {
	BmcDefaultUserPassword                SecretResponsePtrInput `pulumi:"bmcDefaultUserPassword"`
	HcsDataVolumeBitLockerExternalKey     SecretResponsePtrInput `pulumi:"hcsDataVolumeBitLockerExternalKey"`
	HcsInternalVolumeBitLockerExternalKey SecretResponsePtrInput `pulumi:"hcsInternalVolumeBitLockerExternalKey"`
	RotateKeyForDataVolumeBitlocker       SecretResponsePtrInput `pulumi:"rotateKeyForDataVolumeBitlocker"`
	RotateKeysForSedDrivesSerialized      SecretResponsePtrInput `pulumi:"rotateKeysForSedDrivesSerialized"`
	SedEncryptionExternalKey              SecretResponsePtrInput `pulumi:"sedEncryptionExternalKey"`
	SedEncryptionExternalKeyId            SecretResponsePtrInput `pulumi:"sedEncryptionExternalKeyId"`
	SystemVolumeBitLockerRecoveryKey      SecretResponsePtrInput `pulumi:"systemVolumeBitLockerRecoveryKey"`
}

func (DeviceSecretsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceSecretsResponse)(nil)).Elem()
}

func (i DeviceSecretsResponseArgs) ToDeviceSecretsResponseOutput() DeviceSecretsResponseOutput {
	return i.ToDeviceSecretsResponseOutputWithContext(context.Background())
}

func (i DeviceSecretsResponseArgs) ToDeviceSecretsResponseOutputWithContext(ctx context.Context) DeviceSecretsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceSecretsResponseOutput)
}

type DeviceSecretsResponseOutput struct{ *pulumi.OutputState }

func (DeviceSecretsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceSecretsResponse)(nil)).Elem()
}

func (o DeviceSecretsResponseOutput) ToDeviceSecretsResponseOutput() DeviceSecretsResponseOutput {
	return o
}

func (o DeviceSecretsResponseOutput) ToDeviceSecretsResponseOutputWithContext(ctx context.Context) DeviceSecretsResponseOutput {
	return o
}

func (o DeviceSecretsResponseOutput) BmcDefaultUserPassword() SecretResponsePtrOutput {
	return o.ApplyT(func(v DeviceSecretsResponse) *SecretResponse { return v.BmcDefaultUserPassword }).(SecretResponsePtrOutput)
}

func (o DeviceSecretsResponseOutput) HcsDataVolumeBitLockerExternalKey() SecretResponsePtrOutput {
	return o.ApplyT(func(v DeviceSecretsResponse) *SecretResponse { return v.HcsDataVolumeBitLockerExternalKey }).(SecretResponsePtrOutput)
}

func (o DeviceSecretsResponseOutput) HcsInternalVolumeBitLockerExternalKey() SecretResponsePtrOutput {
	return o.ApplyT(func(v DeviceSecretsResponse) *SecretResponse { return v.HcsInternalVolumeBitLockerExternalKey }).(SecretResponsePtrOutput)
}

func (o DeviceSecretsResponseOutput) RotateKeyForDataVolumeBitlocker() SecretResponsePtrOutput {
	return o.ApplyT(func(v DeviceSecretsResponse) *SecretResponse { return v.RotateKeyForDataVolumeBitlocker }).(SecretResponsePtrOutput)
}

func (o DeviceSecretsResponseOutput) RotateKeysForSedDrivesSerialized() SecretResponsePtrOutput {
	return o.ApplyT(func(v DeviceSecretsResponse) *SecretResponse { return v.RotateKeysForSedDrivesSerialized }).(SecretResponsePtrOutput)
}

func (o DeviceSecretsResponseOutput) SedEncryptionExternalKey() SecretResponsePtrOutput {
	return o.ApplyT(func(v DeviceSecretsResponse) *SecretResponse { return v.SedEncryptionExternalKey }).(SecretResponsePtrOutput)
}

func (o DeviceSecretsResponseOutput) SedEncryptionExternalKeyId() SecretResponsePtrOutput {
	return o.ApplyT(func(v DeviceSecretsResponse) *SecretResponse { return v.SedEncryptionExternalKeyId }).(SecretResponsePtrOutput)
}

func (o DeviceSecretsResponseOutput) SystemVolumeBitLockerRecoveryKey() SecretResponsePtrOutput {
	return o.ApplyT(func(v DeviceSecretsResponse) *SecretResponse { return v.SystemVolumeBitLockerRecoveryKey }).(SecretResponsePtrOutput)
}

type EdgeProfileResponse struct {
	Subscription *EdgeProfileSubscriptionResponse `pulumi:"subscription"`
}





type EdgeProfileResponseInput interface {
	pulumi.Input

	ToEdgeProfileResponseOutput() EdgeProfileResponseOutput
	ToEdgeProfileResponseOutputWithContext(context.Context) EdgeProfileResponseOutput
}

type EdgeProfileResponseArgs struct {
	Subscription EdgeProfileSubscriptionResponsePtrInput `pulumi:"subscription"`
}

func (EdgeProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeProfileResponse)(nil)).Elem()
}

func (i EdgeProfileResponseArgs) ToEdgeProfileResponseOutput() EdgeProfileResponseOutput {
	return i.ToEdgeProfileResponseOutputWithContext(context.Background())
}

func (i EdgeProfileResponseArgs) ToEdgeProfileResponseOutputWithContext(ctx context.Context) EdgeProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeProfileResponseOutput)
}

func (i EdgeProfileResponseArgs) ToEdgeProfileResponsePtrOutput() EdgeProfileResponsePtrOutput {
	return i.ToEdgeProfileResponsePtrOutputWithContext(context.Background())
}

func (i EdgeProfileResponseArgs) ToEdgeProfileResponsePtrOutputWithContext(ctx context.Context) EdgeProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeProfileResponseOutput).ToEdgeProfileResponsePtrOutputWithContext(ctx)
}









type EdgeProfileResponsePtrInput interface {
	pulumi.Input

	ToEdgeProfileResponsePtrOutput() EdgeProfileResponsePtrOutput
	ToEdgeProfileResponsePtrOutputWithContext(context.Context) EdgeProfileResponsePtrOutput
}

type edgeProfileResponsePtrType EdgeProfileResponseArgs

func EdgeProfileResponsePtr(v *EdgeProfileResponseArgs) EdgeProfileResponsePtrInput {
	return (*edgeProfileResponsePtrType)(v)
}

func (*edgeProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeProfileResponse)(nil)).Elem()
}

func (i *edgeProfileResponsePtrType) ToEdgeProfileResponsePtrOutput() EdgeProfileResponsePtrOutput {
	return i.ToEdgeProfileResponsePtrOutputWithContext(context.Background())
}

func (i *edgeProfileResponsePtrType) ToEdgeProfileResponsePtrOutputWithContext(ctx context.Context) EdgeProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeProfileResponsePtrOutput)
}

type EdgeProfileResponseOutput struct{ *pulumi.OutputState }

func (EdgeProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeProfileResponse)(nil)).Elem()
}

func (o EdgeProfileResponseOutput) ToEdgeProfileResponseOutput() EdgeProfileResponseOutput {
	return o
}

func (o EdgeProfileResponseOutput) ToEdgeProfileResponseOutputWithContext(ctx context.Context) EdgeProfileResponseOutput {
	return o
}

func (o EdgeProfileResponseOutput) ToEdgeProfileResponsePtrOutput() EdgeProfileResponsePtrOutput {
	return o.ToEdgeProfileResponsePtrOutputWithContext(context.Background())
}

func (o EdgeProfileResponseOutput) ToEdgeProfileResponsePtrOutputWithContext(ctx context.Context) EdgeProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdgeProfileResponse) *EdgeProfileResponse {
		return &v
	}).(EdgeProfileResponsePtrOutput)
}

func (o EdgeProfileResponseOutput) Subscription() EdgeProfileSubscriptionResponsePtrOutput {
	return o.ApplyT(func(v EdgeProfileResponse) *EdgeProfileSubscriptionResponse { return v.Subscription }).(EdgeProfileSubscriptionResponsePtrOutput)
}

type EdgeProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (EdgeProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeProfileResponse)(nil)).Elem()
}

func (o EdgeProfileResponsePtrOutput) ToEdgeProfileResponsePtrOutput() EdgeProfileResponsePtrOutput {
	return o
}

func (o EdgeProfileResponsePtrOutput) ToEdgeProfileResponsePtrOutputWithContext(ctx context.Context) EdgeProfileResponsePtrOutput {
	return o
}

func (o EdgeProfileResponsePtrOutput) Elem() EdgeProfileResponseOutput {
	return o.ApplyT(func(v *EdgeProfileResponse) EdgeProfileResponse {
		if v != nil {
			return *v
		}
		var ret EdgeProfileResponse
		return ret
	}).(EdgeProfileResponseOutput)
}

func (o EdgeProfileResponsePtrOutput) Subscription() EdgeProfileSubscriptionResponsePtrOutput {
	return o.ApplyT(func(v *EdgeProfileResponse) *EdgeProfileSubscriptionResponse {
		if v == nil {
			return nil
		}
		return v.Subscription
	}).(EdgeProfileSubscriptionResponsePtrOutput)
}

type EdgeProfileSubscriptionResponse struct {
	Id                  *string                                  `pulumi:"id"`
	LocationPlacementId *string                                  `pulumi:"locationPlacementId"`
	QuotaId             *string                                  `pulumi:"quotaId"`
	RegisteredFeatures  []SubscriptionRegisteredFeaturesResponse `pulumi:"registeredFeatures"`
	RegistrationDate    *string                                  `pulumi:"registrationDate"`
	RegistrationId      *string                                  `pulumi:"registrationId"`
	SerializedDetails   *string                                  `pulumi:"serializedDetails"`
	State               *string                                  `pulumi:"state"`
	SubscriptionId      *string                                  `pulumi:"subscriptionId"`
	TenantId            *string                                  `pulumi:"tenantId"`
}





type EdgeProfileSubscriptionResponseInput interface {
	pulumi.Input

	ToEdgeProfileSubscriptionResponseOutput() EdgeProfileSubscriptionResponseOutput
	ToEdgeProfileSubscriptionResponseOutputWithContext(context.Context) EdgeProfileSubscriptionResponseOutput
}

type EdgeProfileSubscriptionResponseArgs struct {
	Id                  pulumi.StringPtrInput                            `pulumi:"id"`
	LocationPlacementId pulumi.StringPtrInput                            `pulumi:"locationPlacementId"`
	QuotaId             pulumi.StringPtrInput                            `pulumi:"quotaId"`
	RegisteredFeatures  SubscriptionRegisteredFeaturesResponseArrayInput `pulumi:"registeredFeatures"`
	RegistrationDate    pulumi.StringPtrInput                            `pulumi:"registrationDate"`
	RegistrationId      pulumi.StringPtrInput                            `pulumi:"registrationId"`
	SerializedDetails   pulumi.StringPtrInput                            `pulumi:"serializedDetails"`
	State               pulumi.StringPtrInput                            `pulumi:"state"`
	SubscriptionId      pulumi.StringPtrInput                            `pulumi:"subscriptionId"`
	TenantId            pulumi.StringPtrInput                            `pulumi:"tenantId"`
}

func (EdgeProfileSubscriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeProfileSubscriptionResponse)(nil)).Elem()
}

func (i EdgeProfileSubscriptionResponseArgs) ToEdgeProfileSubscriptionResponseOutput() EdgeProfileSubscriptionResponseOutput {
	return i.ToEdgeProfileSubscriptionResponseOutputWithContext(context.Background())
}

func (i EdgeProfileSubscriptionResponseArgs) ToEdgeProfileSubscriptionResponseOutputWithContext(ctx context.Context) EdgeProfileSubscriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeProfileSubscriptionResponseOutput)
}

func (i EdgeProfileSubscriptionResponseArgs) ToEdgeProfileSubscriptionResponsePtrOutput() EdgeProfileSubscriptionResponsePtrOutput {
	return i.ToEdgeProfileSubscriptionResponsePtrOutputWithContext(context.Background())
}

func (i EdgeProfileSubscriptionResponseArgs) ToEdgeProfileSubscriptionResponsePtrOutputWithContext(ctx context.Context) EdgeProfileSubscriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeProfileSubscriptionResponseOutput).ToEdgeProfileSubscriptionResponsePtrOutputWithContext(ctx)
}









type EdgeProfileSubscriptionResponsePtrInput interface {
	pulumi.Input

	ToEdgeProfileSubscriptionResponsePtrOutput() EdgeProfileSubscriptionResponsePtrOutput
	ToEdgeProfileSubscriptionResponsePtrOutputWithContext(context.Context) EdgeProfileSubscriptionResponsePtrOutput
}

type edgeProfileSubscriptionResponsePtrType EdgeProfileSubscriptionResponseArgs

func EdgeProfileSubscriptionResponsePtr(v *EdgeProfileSubscriptionResponseArgs) EdgeProfileSubscriptionResponsePtrInput {
	return (*edgeProfileSubscriptionResponsePtrType)(v)
}

func (*edgeProfileSubscriptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeProfileSubscriptionResponse)(nil)).Elem()
}

func (i *edgeProfileSubscriptionResponsePtrType) ToEdgeProfileSubscriptionResponsePtrOutput() EdgeProfileSubscriptionResponsePtrOutput {
	return i.ToEdgeProfileSubscriptionResponsePtrOutputWithContext(context.Background())
}

func (i *edgeProfileSubscriptionResponsePtrType) ToEdgeProfileSubscriptionResponsePtrOutputWithContext(ctx context.Context) EdgeProfileSubscriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeProfileSubscriptionResponsePtrOutput)
}

type EdgeProfileSubscriptionResponseOutput struct{ *pulumi.OutputState }

func (EdgeProfileSubscriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeProfileSubscriptionResponse)(nil)).Elem()
}

func (o EdgeProfileSubscriptionResponseOutput) ToEdgeProfileSubscriptionResponseOutput() EdgeProfileSubscriptionResponseOutput {
	return o
}

func (o EdgeProfileSubscriptionResponseOutput) ToEdgeProfileSubscriptionResponseOutputWithContext(ctx context.Context) EdgeProfileSubscriptionResponseOutput {
	return o
}

func (o EdgeProfileSubscriptionResponseOutput) ToEdgeProfileSubscriptionResponsePtrOutput() EdgeProfileSubscriptionResponsePtrOutput {
	return o.ToEdgeProfileSubscriptionResponsePtrOutputWithContext(context.Background())
}

func (o EdgeProfileSubscriptionResponseOutput) ToEdgeProfileSubscriptionResponsePtrOutputWithContext(ctx context.Context) EdgeProfileSubscriptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdgeProfileSubscriptionResponse) *EdgeProfileSubscriptionResponse {
		return &v
	}).(EdgeProfileSubscriptionResponsePtrOutput)
}

func (o EdgeProfileSubscriptionResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeProfileSubscriptionResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponseOutput) LocationPlacementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeProfileSubscriptionResponse) *string { return v.LocationPlacementId }).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponseOutput) QuotaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeProfileSubscriptionResponse) *string { return v.QuotaId }).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponseOutput) RegisteredFeatures() SubscriptionRegisteredFeaturesResponseArrayOutput {
	return o.ApplyT(func(v EdgeProfileSubscriptionResponse) []SubscriptionRegisteredFeaturesResponse {
		return v.RegisteredFeatures
	}).(SubscriptionRegisteredFeaturesResponseArrayOutput)
}

func (o EdgeProfileSubscriptionResponseOutput) RegistrationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeProfileSubscriptionResponse) *string { return v.RegistrationDate }).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponseOutput) RegistrationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeProfileSubscriptionResponse) *string { return v.RegistrationId }).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponseOutput) SerializedDetails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeProfileSubscriptionResponse) *string { return v.SerializedDetails }).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeProfileSubscriptionResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeProfileSubscriptionResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeProfileSubscriptionResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type EdgeProfileSubscriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EdgeProfileSubscriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeProfileSubscriptionResponse)(nil)).Elem()
}

func (o EdgeProfileSubscriptionResponsePtrOutput) ToEdgeProfileSubscriptionResponsePtrOutput() EdgeProfileSubscriptionResponsePtrOutput {
	return o
}

func (o EdgeProfileSubscriptionResponsePtrOutput) ToEdgeProfileSubscriptionResponsePtrOutputWithContext(ctx context.Context) EdgeProfileSubscriptionResponsePtrOutput {
	return o
}

func (o EdgeProfileSubscriptionResponsePtrOutput) Elem() EdgeProfileSubscriptionResponseOutput {
	return o.ApplyT(func(v *EdgeProfileSubscriptionResponse) EdgeProfileSubscriptionResponse {
		if v != nil {
			return *v
		}
		var ret EdgeProfileSubscriptionResponse
		return ret
	}).(EdgeProfileSubscriptionResponseOutput)
}

func (o EdgeProfileSubscriptionResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeProfileSubscriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponsePtrOutput) LocationPlacementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeProfileSubscriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.LocationPlacementId
	}).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponsePtrOutput) QuotaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeProfileSubscriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.QuotaId
	}).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponsePtrOutput) RegisteredFeatures() SubscriptionRegisteredFeaturesResponseArrayOutput {
	return o.ApplyT(func(v *EdgeProfileSubscriptionResponse) []SubscriptionRegisteredFeaturesResponse {
		if v == nil {
			return nil
		}
		return v.RegisteredFeatures
	}).(SubscriptionRegisteredFeaturesResponseArrayOutput)
}

func (o EdgeProfileSubscriptionResponsePtrOutput) RegistrationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeProfileSubscriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationDate
	}).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponsePtrOutput) RegistrationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeProfileSubscriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationId
	}).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponsePtrOutput) SerializedDetails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeProfileSubscriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SerializedDetails
	}).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeProfileSubscriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeProfileSubscriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o EdgeProfileSubscriptionResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeProfileSubscriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type EtcdInfoResponse struct {
	Type    string `pulumi:"type"`
	Version string `pulumi:"version"`
}





type EtcdInfoResponseInput interface {
	pulumi.Input

	ToEtcdInfoResponseOutput() EtcdInfoResponseOutput
	ToEtcdInfoResponseOutputWithContext(context.Context) EtcdInfoResponseOutput
}

type EtcdInfoResponseArgs struct {
	Type    pulumi.StringInput `pulumi:"type"`
	Version pulumi.StringInput `pulumi:"version"`
}

func (EtcdInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EtcdInfoResponse)(nil)).Elem()
}

func (i EtcdInfoResponseArgs) ToEtcdInfoResponseOutput() EtcdInfoResponseOutput {
	return i.ToEtcdInfoResponseOutputWithContext(context.Background())
}

func (i EtcdInfoResponseArgs) ToEtcdInfoResponseOutputWithContext(ctx context.Context) EtcdInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdInfoResponseOutput)
}

func (i EtcdInfoResponseArgs) ToEtcdInfoResponsePtrOutput() EtcdInfoResponsePtrOutput {
	return i.ToEtcdInfoResponsePtrOutputWithContext(context.Background())
}

func (i EtcdInfoResponseArgs) ToEtcdInfoResponsePtrOutputWithContext(ctx context.Context) EtcdInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdInfoResponseOutput).ToEtcdInfoResponsePtrOutputWithContext(ctx)
}









type EtcdInfoResponsePtrInput interface {
	pulumi.Input

	ToEtcdInfoResponsePtrOutput() EtcdInfoResponsePtrOutput
	ToEtcdInfoResponsePtrOutputWithContext(context.Context) EtcdInfoResponsePtrOutput
}

type etcdInfoResponsePtrType EtcdInfoResponseArgs

func EtcdInfoResponsePtr(v *EtcdInfoResponseArgs) EtcdInfoResponsePtrInput {
	return (*etcdInfoResponsePtrType)(v)
}

func (*etcdInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EtcdInfoResponse)(nil)).Elem()
}

func (i *etcdInfoResponsePtrType) ToEtcdInfoResponsePtrOutput() EtcdInfoResponsePtrOutput {
	return i.ToEtcdInfoResponsePtrOutputWithContext(context.Background())
}

func (i *etcdInfoResponsePtrType) ToEtcdInfoResponsePtrOutputWithContext(ctx context.Context) EtcdInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdInfoResponsePtrOutput)
}

type EtcdInfoResponseOutput struct{ *pulumi.OutputState }

func (EtcdInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EtcdInfoResponse)(nil)).Elem()
}

func (o EtcdInfoResponseOutput) ToEtcdInfoResponseOutput() EtcdInfoResponseOutput {
	return o
}

func (o EtcdInfoResponseOutput) ToEtcdInfoResponseOutputWithContext(ctx context.Context) EtcdInfoResponseOutput {
	return o
}

func (o EtcdInfoResponseOutput) ToEtcdInfoResponsePtrOutput() EtcdInfoResponsePtrOutput {
	return o.ToEtcdInfoResponsePtrOutputWithContext(context.Background())
}

func (o EtcdInfoResponseOutput) ToEtcdInfoResponsePtrOutputWithContext(ctx context.Context) EtcdInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EtcdInfoResponse) *EtcdInfoResponse {
		return &v
	}).(EtcdInfoResponsePtrOutput)
}

func (o EtcdInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EtcdInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o EtcdInfoResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v EtcdInfoResponse) string { return v.Version }).(pulumi.StringOutput)
}

type EtcdInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (EtcdInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EtcdInfoResponse)(nil)).Elem()
}

func (o EtcdInfoResponsePtrOutput) ToEtcdInfoResponsePtrOutput() EtcdInfoResponsePtrOutput {
	return o
}

func (o EtcdInfoResponsePtrOutput) ToEtcdInfoResponsePtrOutputWithContext(ctx context.Context) EtcdInfoResponsePtrOutput {
	return o
}

func (o EtcdInfoResponsePtrOutput) Elem() EtcdInfoResponseOutput {
	return o.ApplyT(func(v *EtcdInfoResponse) EtcdInfoResponse {
		if v != nil {
			return *v
		}
		var ret EtcdInfoResponse
		return ret
	}).(EtcdInfoResponseOutput)
}

func (o EtcdInfoResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EtcdInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o EtcdInfoResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EtcdInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
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

func (i FileSourceInfoArgs) ToFileSourceInfoPtrOutput() FileSourceInfoPtrOutput {
	return i.ToFileSourceInfoPtrOutputWithContext(context.Background())
}

func (i FileSourceInfoArgs) ToFileSourceInfoPtrOutputWithContext(ctx context.Context) FileSourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSourceInfoOutput).ToFileSourceInfoPtrOutputWithContext(ctx)
}









type FileSourceInfoPtrInput interface {
	pulumi.Input

	ToFileSourceInfoPtrOutput() FileSourceInfoPtrOutput
	ToFileSourceInfoPtrOutputWithContext(context.Context) FileSourceInfoPtrOutput
}

type fileSourceInfoPtrType FileSourceInfoArgs

func FileSourceInfoPtr(v *FileSourceInfoArgs) FileSourceInfoPtrInput {
	return (*fileSourceInfoPtrType)(v)
}

func (*fileSourceInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSourceInfo)(nil)).Elem()
}

func (i *fileSourceInfoPtrType) ToFileSourceInfoPtrOutput() FileSourceInfoPtrOutput {
	return i.ToFileSourceInfoPtrOutputWithContext(context.Background())
}

func (i *fileSourceInfoPtrType) ToFileSourceInfoPtrOutputWithContext(ctx context.Context) FileSourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSourceInfoPtrOutput)
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

func (o FileSourceInfoOutput) ToFileSourceInfoPtrOutput() FileSourceInfoPtrOutput {
	return o.ToFileSourceInfoPtrOutputWithContext(context.Background())
}

func (o FileSourceInfoOutput) ToFileSourceInfoPtrOutputWithContext(ctx context.Context) FileSourceInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileSourceInfo) *FileSourceInfo {
		return &v
	}).(FileSourceInfoPtrOutput)
}

func (o FileSourceInfoOutput) ShareId() pulumi.StringOutput {
	return o.ApplyT(func(v FileSourceInfo) string { return v.ShareId }).(pulumi.StringOutput)
}

type FileSourceInfoPtrOutput struct{ *pulumi.OutputState }

func (FileSourceInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSourceInfo)(nil)).Elem()
}

func (o FileSourceInfoPtrOutput) ToFileSourceInfoPtrOutput() FileSourceInfoPtrOutput {
	return o
}

func (o FileSourceInfoPtrOutput) ToFileSourceInfoPtrOutputWithContext(ctx context.Context) FileSourceInfoPtrOutput {
	return o
}

func (o FileSourceInfoPtrOutput) Elem() FileSourceInfoOutput {
	return o.ApplyT(func(v *FileSourceInfo) FileSourceInfo {
		if v != nil {
			return *v
		}
		var ret FileSourceInfo
		return ret
	}).(FileSourceInfoOutput)
}

func (o FileSourceInfoPtrOutput) ShareId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSourceInfo) *string {
		if v == nil {
			return nil
		}
		return &v.ShareId
	}).(pulumi.StringPtrOutput)
}

type FileSourceInfoResponse struct {
	ShareId string `pulumi:"shareId"`
}





type FileSourceInfoResponseInput interface {
	pulumi.Input

	ToFileSourceInfoResponseOutput() FileSourceInfoResponseOutput
	ToFileSourceInfoResponseOutputWithContext(context.Context) FileSourceInfoResponseOutput
}

type FileSourceInfoResponseArgs struct {
	ShareId pulumi.StringInput `pulumi:"shareId"`
}

func (FileSourceInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSourceInfoResponse)(nil)).Elem()
}

func (i FileSourceInfoResponseArgs) ToFileSourceInfoResponseOutput() FileSourceInfoResponseOutput {
	return i.ToFileSourceInfoResponseOutputWithContext(context.Background())
}

func (i FileSourceInfoResponseArgs) ToFileSourceInfoResponseOutputWithContext(ctx context.Context) FileSourceInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSourceInfoResponseOutput)
}

func (i FileSourceInfoResponseArgs) ToFileSourceInfoResponsePtrOutput() FileSourceInfoResponsePtrOutput {
	return i.ToFileSourceInfoResponsePtrOutputWithContext(context.Background())
}

func (i FileSourceInfoResponseArgs) ToFileSourceInfoResponsePtrOutputWithContext(ctx context.Context) FileSourceInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSourceInfoResponseOutput).ToFileSourceInfoResponsePtrOutputWithContext(ctx)
}









type FileSourceInfoResponsePtrInput interface {
	pulumi.Input

	ToFileSourceInfoResponsePtrOutput() FileSourceInfoResponsePtrOutput
	ToFileSourceInfoResponsePtrOutputWithContext(context.Context) FileSourceInfoResponsePtrOutput
}

type fileSourceInfoResponsePtrType FileSourceInfoResponseArgs

func FileSourceInfoResponsePtr(v *FileSourceInfoResponseArgs) FileSourceInfoResponsePtrInput {
	return (*fileSourceInfoResponsePtrType)(v)
}

func (*fileSourceInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSourceInfoResponse)(nil)).Elem()
}

func (i *fileSourceInfoResponsePtrType) ToFileSourceInfoResponsePtrOutput() FileSourceInfoResponsePtrOutput {
	return i.ToFileSourceInfoResponsePtrOutputWithContext(context.Background())
}

func (i *fileSourceInfoResponsePtrType) ToFileSourceInfoResponsePtrOutputWithContext(ctx context.Context) FileSourceInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSourceInfoResponsePtrOutput)
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

func (o FileSourceInfoResponseOutput) ToFileSourceInfoResponsePtrOutput() FileSourceInfoResponsePtrOutput {
	return o.ToFileSourceInfoResponsePtrOutputWithContext(context.Background())
}

func (o FileSourceInfoResponseOutput) ToFileSourceInfoResponsePtrOutputWithContext(ctx context.Context) FileSourceInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileSourceInfoResponse) *FileSourceInfoResponse {
		return &v
	}).(FileSourceInfoResponsePtrOutput)
}

func (o FileSourceInfoResponseOutput) ShareId() pulumi.StringOutput {
	return o.ApplyT(func(v FileSourceInfoResponse) string { return v.ShareId }).(pulumi.StringOutput)
}

type FileSourceInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (FileSourceInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSourceInfoResponse)(nil)).Elem()
}

func (o FileSourceInfoResponsePtrOutput) ToFileSourceInfoResponsePtrOutput() FileSourceInfoResponsePtrOutput {
	return o
}

func (o FileSourceInfoResponsePtrOutput) ToFileSourceInfoResponsePtrOutputWithContext(ctx context.Context) FileSourceInfoResponsePtrOutput {
	return o
}

func (o FileSourceInfoResponsePtrOutput) Elem() FileSourceInfoResponseOutput {
	return o.ApplyT(func(v *FileSourceInfoResponse) FileSourceInfoResponse {
		if v != nil {
			return *v
		}
		var ret FileSourceInfoResponse
		return ret
	}).(FileSourceInfoResponseOutput)
}

func (o FileSourceInfoResponsePtrOutput) ShareId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ShareId
	}).(pulumi.StringPtrOutput)
}

type ImageRepositoryCredential struct {
	ImageRepositoryUrl string                     `pulumi:"imageRepositoryUrl"`
	Password           *AsymmetricEncryptedSecret `pulumi:"password"`
	UserName           string                     `pulumi:"userName"`
}





type ImageRepositoryCredentialInput interface {
	pulumi.Input

	ToImageRepositoryCredentialOutput() ImageRepositoryCredentialOutput
	ToImageRepositoryCredentialOutputWithContext(context.Context) ImageRepositoryCredentialOutput
}

type ImageRepositoryCredentialArgs struct {
	ImageRepositoryUrl pulumi.StringInput                `pulumi:"imageRepositoryUrl"`
	Password           AsymmetricEncryptedSecretPtrInput `pulumi:"password"`
	UserName           pulumi.StringInput                `pulumi:"userName"`
}

func (ImageRepositoryCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRepositoryCredential)(nil)).Elem()
}

func (i ImageRepositoryCredentialArgs) ToImageRepositoryCredentialOutput() ImageRepositoryCredentialOutput {
	return i.ToImageRepositoryCredentialOutputWithContext(context.Background())
}

func (i ImageRepositoryCredentialArgs) ToImageRepositoryCredentialOutputWithContext(ctx context.Context) ImageRepositoryCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRepositoryCredentialOutput)
}

func (i ImageRepositoryCredentialArgs) ToImageRepositoryCredentialPtrOutput() ImageRepositoryCredentialPtrOutput {
	return i.ToImageRepositoryCredentialPtrOutputWithContext(context.Background())
}

func (i ImageRepositoryCredentialArgs) ToImageRepositoryCredentialPtrOutputWithContext(ctx context.Context) ImageRepositoryCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRepositoryCredentialOutput).ToImageRepositoryCredentialPtrOutputWithContext(ctx)
}









type ImageRepositoryCredentialPtrInput interface {
	pulumi.Input

	ToImageRepositoryCredentialPtrOutput() ImageRepositoryCredentialPtrOutput
	ToImageRepositoryCredentialPtrOutputWithContext(context.Context) ImageRepositoryCredentialPtrOutput
}

type imageRepositoryCredentialPtrType ImageRepositoryCredentialArgs

func ImageRepositoryCredentialPtr(v *ImageRepositoryCredentialArgs) ImageRepositoryCredentialPtrInput {
	return (*imageRepositoryCredentialPtrType)(v)
}

func (*imageRepositoryCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRepositoryCredential)(nil)).Elem()
}

func (i *imageRepositoryCredentialPtrType) ToImageRepositoryCredentialPtrOutput() ImageRepositoryCredentialPtrOutput {
	return i.ToImageRepositoryCredentialPtrOutputWithContext(context.Background())
}

func (i *imageRepositoryCredentialPtrType) ToImageRepositoryCredentialPtrOutputWithContext(ctx context.Context) ImageRepositoryCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRepositoryCredentialPtrOutput)
}

type ImageRepositoryCredentialOutput struct{ *pulumi.OutputState }

func (ImageRepositoryCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRepositoryCredential)(nil)).Elem()
}

func (o ImageRepositoryCredentialOutput) ToImageRepositoryCredentialOutput() ImageRepositoryCredentialOutput {
	return o
}

func (o ImageRepositoryCredentialOutput) ToImageRepositoryCredentialOutputWithContext(ctx context.Context) ImageRepositoryCredentialOutput {
	return o
}

func (o ImageRepositoryCredentialOutput) ToImageRepositoryCredentialPtrOutput() ImageRepositoryCredentialPtrOutput {
	return o.ToImageRepositoryCredentialPtrOutputWithContext(context.Background())
}

func (o ImageRepositoryCredentialOutput) ToImageRepositoryCredentialPtrOutputWithContext(ctx context.Context) ImageRepositoryCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageRepositoryCredential) *ImageRepositoryCredential {
		return &v
	}).(ImageRepositoryCredentialPtrOutput)
}

func (o ImageRepositoryCredentialOutput) ImageRepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRepositoryCredential) string { return v.ImageRepositoryUrl }).(pulumi.StringOutput)
}

func (o ImageRepositoryCredentialOutput) Password() AsymmetricEncryptedSecretPtrOutput {
	return o.ApplyT(func(v ImageRepositoryCredential) *AsymmetricEncryptedSecret { return v.Password }).(AsymmetricEncryptedSecretPtrOutput)
}

func (o ImageRepositoryCredentialOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRepositoryCredential) string { return v.UserName }).(pulumi.StringOutput)
}

type ImageRepositoryCredentialPtrOutput struct{ *pulumi.OutputState }

func (ImageRepositoryCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRepositoryCredential)(nil)).Elem()
}

func (o ImageRepositoryCredentialPtrOutput) ToImageRepositoryCredentialPtrOutput() ImageRepositoryCredentialPtrOutput {
	return o
}

func (o ImageRepositoryCredentialPtrOutput) ToImageRepositoryCredentialPtrOutputWithContext(ctx context.Context) ImageRepositoryCredentialPtrOutput {
	return o
}

func (o ImageRepositoryCredentialPtrOutput) Elem() ImageRepositoryCredentialOutput {
	return o.ApplyT(func(v *ImageRepositoryCredential) ImageRepositoryCredential {
		if v != nil {
			return *v
		}
		var ret ImageRepositoryCredential
		return ret
	}).(ImageRepositoryCredentialOutput)
}

func (o ImageRepositoryCredentialPtrOutput) ImageRepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRepositoryCredential) *string {
		if v == nil {
			return nil
		}
		return &v.ImageRepositoryUrl
	}).(pulumi.StringPtrOutput)
}

func (o ImageRepositoryCredentialPtrOutput) Password() AsymmetricEncryptedSecretPtrOutput {
	return o.ApplyT(func(v *ImageRepositoryCredential) *AsymmetricEncryptedSecret {
		if v == nil {
			return nil
		}
		return v.Password
	}).(AsymmetricEncryptedSecretPtrOutput)
}

func (o ImageRepositoryCredentialPtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRepositoryCredential) *string {
		if v == nil {
			return nil
		}
		return &v.UserName
	}).(pulumi.StringPtrOutput)
}

type ImageRepositoryCredentialResponse struct {
	ImageRepositoryUrl string                             `pulumi:"imageRepositoryUrl"`
	Password           *AsymmetricEncryptedSecretResponse `pulumi:"password"`
	UserName           string                             `pulumi:"userName"`
}





type ImageRepositoryCredentialResponseInput interface {
	pulumi.Input

	ToImageRepositoryCredentialResponseOutput() ImageRepositoryCredentialResponseOutput
	ToImageRepositoryCredentialResponseOutputWithContext(context.Context) ImageRepositoryCredentialResponseOutput
}

type ImageRepositoryCredentialResponseArgs struct {
	ImageRepositoryUrl pulumi.StringInput                        `pulumi:"imageRepositoryUrl"`
	Password           AsymmetricEncryptedSecretResponsePtrInput `pulumi:"password"`
	UserName           pulumi.StringInput                        `pulumi:"userName"`
}

func (ImageRepositoryCredentialResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRepositoryCredentialResponse)(nil)).Elem()
}

func (i ImageRepositoryCredentialResponseArgs) ToImageRepositoryCredentialResponseOutput() ImageRepositoryCredentialResponseOutput {
	return i.ToImageRepositoryCredentialResponseOutputWithContext(context.Background())
}

func (i ImageRepositoryCredentialResponseArgs) ToImageRepositoryCredentialResponseOutputWithContext(ctx context.Context) ImageRepositoryCredentialResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRepositoryCredentialResponseOutput)
}

func (i ImageRepositoryCredentialResponseArgs) ToImageRepositoryCredentialResponsePtrOutput() ImageRepositoryCredentialResponsePtrOutput {
	return i.ToImageRepositoryCredentialResponsePtrOutputWithContext(context.Background())
}

func (i ImageRepositoryCredentialResponseArgs) ToImageRepositoryCredentialResponsePtrOutputWithContext(ctx context.Context) ImageRepositoryCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRepositoryCredentialResponseOutput).ToImageRepositoryCredentialResponsePtrOutputWithContext(ctx)
}









type ImageRepositoryCredentialResponsePtrInput interface {
	pulumi.Input

	ToImageRepositoryCredentialResponsePtrOutput() ImageRepositoryCredentialResponsePtrOutput
	ToImageRepositoryCredentialResponsePtrOutputWithContext(context.Context) ImageRepositoryCredentialResponsePtrOutput
}

type imageRepositoryCredentialResponsePtrType ImageRepositoryCredentialResponseArgs

func ImageRepositoryCredentialResponsePtr(v *ImageRepositoryCredentialResponseArgs) ImageRepositoryCredentialResponsePtrInput {
	return (*imageRepositoryCredentialResponsePtrType)(v)
}

func (*imageRepositoryCredentialResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRepositoryCredentialResponse)(nil)).Elem()
}

func (i *imageRepositoryCredentialResponsePtrType) ToImageRepositoryCredentialResponsePtrOutput() ImageRepositoryCredentialResponsePtrOutput {
	return i.ToImageRepositoryCredentialResponsePtrOutputWithContext(context.Background())
}

func (i *imageRepositoryCredentialResponsePtrType) ToImageRepositoryCredentialResponsePtrOutputWithContext(ctx context.Context) ImageRepositoryCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRepositoryCredentialResponsePtrOutput)
}

type ImageRepositoryCredentialResponseOutput struct{ *pulumi.OutputState }

func (ImageRepositoryCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRepositoryCredentialResponse)(nil)).Elem()
}

func (o ImageRepositoryCredentialResponseOutput) ToImageRepositoryCredentialResponseOutput() ImageRepositoryCredentialResponseOutput {
	return o
}

func (o ImageRepositoryCredentialResponseOutput) ToImageRepositoryCredentialResponseOutputWithContext(ctx context.Context) ImageRepositoryCredentialResponseOutput {
	return o
}

func (o ImageRepositoryCredentialResponseOutput) ToImageRepositoryCredentialResponsePtrOutput() ImageRepositoryCredentialResponsePtrOutput {
	return o.ToImageRepositoryCredentialResponsePtrOutputWithContext(context.Background())
}

func (o ImageRepositoryCredentialResponseOutput) ToImageRepositoryCredentialResponsePtrOutputWithContext(ctx context.Context) ImageRepositoryCredentialResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageRepositoryCredentialResponse) *ImageRepositoryCredentialResponse {
		return &v
	}).(ImageRepositoryCredentialResponsePtrOutput)
}

func (o ImageRepositoryCredentialResponseOutput) ImageRepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRepositoryCredentialResponse) string { return v.ImageRepositoryUrl }).(pulumi.StringOutput)
}

func (o ImageRepositoryCredentialResponseOutput) Password() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v ImageRepositoryCredentialResponse) *AsymmetricEncryptedSecretResponse { return v.Password }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o ImageRepositoryCredentialResponseOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRepositoryCredentialResponse) string { return v.UserName }).(pulumi.StringOutput)
}

type ImageRepositoryCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageRepositoryCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRepositoryCredentialResponse)(nil)).Elem()
}

func (o ImageRepositoryCredentialResponsePtrOutput) ToImageRepositoryCredentialResponsePtrOutput() ImageRepositoryCredentialResponsePtrOutput {
	return o
}

func (o ImageRepositoryCredentialResponsePtrOutput) ToImageRepositoryCredentialResponsePtrOutputWithContext(ctx context.Context) ImageRepositoryCredentialResponsePtrOutput {
	return o
}

func (o ImageRepositoryCredentialResponsePtrOutput) Elem() ImageRepositoryCredentialResponseOutput {
	return o.ApplyT(func(v *ImageRepositoryCredentialResponse) ImageRepositoryCredentialResponse {
		if v != nil {
			return *v
		}
		var ret ImageRepositoryCredentialResponse
		return ret
	}).(ImageRepositoryCredentialResponseOutput)
}

func (o ImageRepositoryCredentialResponsePtrOutput) ImageRepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRepositoryCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ImageRepositoryUrl
	}).(pulumi.StringPtrOutput)
}

func (o ImageRepositoryCredentialResponsePtrOutput) Password() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v *ImageRepositoryCredentialResponse) *AsymmetricEncryptedSecretResponse {
		if v == nil {
			return nil
		}
		return v.Password
	}).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o ImageRepositoryCredentialResponsePtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRepositoryCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UserName
	}).(pulumi.StringPtrOutput)
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

func (i IoTDeviceInfoArgs) ToIoTDeviceInfoPtrOutput() IoTDeviceInfoPtrOutput {
	return i.ToIoTDeviceInfoPtrOutputWithContext(context.Background())
}

func (i IoTDeviceInfoArgs) ToIoTDeviceInfoPtrOutputWithContext(ctx context.Context) IoTDeviceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTDeviceInfoOutput).ToIoTDeviceInfoPtrOutputWithContext(ctx)
}









type IoTDeviceInfoPtrInput interface {
	pulumi.Input

	ToIoTDeviceInfoPtrOutput() IoTDeviceInfoPtrOutput
	ToIoTDeviceInfoPtrOutputWithContext(context.Context) IoTDeviceInfoPtrOutput
}

type ioTDeviceInfoPtrType IoTDeviceInfoArgs

func IoTDeviceInfoPtr(v *IoTDeviceInfoArgs) IoTDeviceInfoPtrInput {
	return (*ioTDeviceInfoPtrType)(v)
}

func (*ioTDeviceInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTDeviceInfo)(nil)).Elem()
}

func (i *ioTDeviceInfoPtrType) ToIoTDeviceInfoPtrOutput() IoTDeviceInfoPtrOutput {
	return i.ToIoTDeviceInfoPtrOutputWithContext(context.Background())
}

func (i *ioTDeviceInfoPtrType) ToIoTDeviceInfoPtrOutputWithContext(ctx context.Context) IoTDeviceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTDeviceInfoPtrOutput)
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

func (o IoTDeviceInfoOutput) ToIoTDeviceInfoPtrOutput() IoTDeviceInfoPtrOutput {
	return o.ToIoTDeviceInfoPtrOutputWithContext(context.Background())
}

func (o IoTDeviceInfoOutput) ToIoTDeviceInfoPtrOutputWithContext(ctx context.Context) IoTDeviceInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IoTDeviceInfo) *IoTDeviceInfo {
		return &v
	}).(IoTDeviceInfoPtrOutput)
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

type IoTDeviceInfoPtrOutput struct{ *pulumi.OutputState }

func (IoTDeviceInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTDeviceInfo)(nil)).Elem()
}

func (o IoTDeviceInfoPtrOutput) ToIoTDeviceInfoPtrOutput() IoTDeviceInfoPtrOutput {
	return o
}

func (o IoTDeviceInfoPtrOutput) ToIoTDeviceInfoPtrOutputWithContext(ctx context.Context) IoTDeviceInfoPtrOutput {
	return o
}

func (o IoTDeviceInfoPtrOutput) Elem() IoTDeviceInfoOutput {
	return o.ApplyT(func(v *IoTDeviceInfo) IoTDeviceInfo {
		if v != nil {
			return *v
		}
		var ret IoTDeviceInfo
		return ret
	}).(IoTDeviceInfoOutput)
}

func (o IoTDeviceInfoPtrOutput) Authentication() AuthenticationPtrOutput {
	return o.ApplyT(func(v *IoTDeviceInfo) *Authentication {
		if v == nil {
			return nil
		}
		return v.Authentication
	}).(AuthenticationPtrOutput)
}

func (o IoTDeviceInfoPtrOutput) DeviceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTDeviceInfo) *string {
		if v == nil {
			return nil
		}
		return &v.DeviceId
	}).(pulumi.StringPtrOutput)
}

func (o IoTDeviceInfoPtrOutput) IoTHostHub() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTDeviceInfo) *string {
		if v == nil {
			return nil
		}
		return &v.IoTHostHub
	}).(pulumi.StringPtrOutput)
}

func (o IoTDeviceInfoPtrOutput) IoTHostHubId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTDeviceInfo) *string {
		if v == nil {
			return nil
		}
		return v.IoTHostHubId
	}).(pulumi.StringPtrOutput)
}

type IoTDeviceInfoResponse struct {
	Authentication *AuthenticationResponse `pulumi:"authentication"`
	DeviceId       string                  `pulumi:"deviceId"`
	IoTHostHub     string                  `pulumi:"ioTHostHub"`
	IoTHostHubId   *string                 `pulumi:"ioTHostHubId"`
}





type IoTDeviceInfoResponseInput interface {
	pulumi.Input

	ToIoTDeviceInfoResponseOutput() IoTDeviceInfoResponseOutput
	ToIoTDeviceInfoResponseOutputWithContext(context.Context) IoTDeviceInfoResponseOutput
}

type IoTDeviceInfoResponseArgs struct {
	Authentication AuthenticationResponsePtrInput `pulumi:"authentication"`
	DeviceId       pulumi.StringInput             `pulumi:"deviceId"`
	IoTHostHub     pulumi.StringInput             `pulumi:"ioTHostHub"`
	IoTHostHubId   pulumi.StringPtrInput          `pulumi:"ioTHostHubId"`
}

func (IoTDeviceInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTDeviceInfoResponse)(nil)).Elem()
}

func (i IoTDeviceInfoResponseArgs) ToIoTDeviceInfoResponseOutput() IoTDeviceInfoResponseOutput {
	return i.ToIoTDeviceInfoResponseOutputWithContext(context.Background())
}

func (i IoTDeviceInfoResponseArgs) ToIoTDeviceInfoResponseOutputWithContext(ctx context.Context) IoTDeviceInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTDeviceInfoResponseOutput)
}

func (i IoTDeviceInfoResponseArgs) ToIoTDeviceInfoResponsePtrOutput() IoTDeviceInfoResponsePtrOutput {
	return i.ToIoTDeviceInfoResponsePtrOutputWithContext(context.Background())
}

func (i IoTDeviceInfoResponseArgs) ToIoTDeviceInfoResponsePtrOutputWithContext(ctx context.Context) IoTDeviceInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTDeviceInfoResponseOutput).ToIoTDeviceInfoResponsePtrOutputWithContext(ctx)
}









type IoTDeviceInfoResponsePtrInput interface {
	pulumi.Input

	ToIoTDeviceInfoResponsePtrOutput() IoTDeviceInfoResponsePtrOutput
	ToIoTDeviceInfoResponsePtrOutputWithContext(context.Context) IoTDeviceInfoResponsePtrOutput
}

type ioTDeviceInfoResponsePtrType IoTDeviceInfoResponseArgs

func IoTDeviceInfoResponsePtr(v *IoTDeviceInfoResponseArgs) IoTDeviceInfoResponsePtrInput {
	return (*ioTDeviceInfoResponsePtrType)(v)
}

func (*ioTDeviceInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTDeviceInfoResponse)(nil)).Elem()
}

func (i *ioTDeviceInfoResponsePtrType) ToIoTDeviceInfoResponsePtrOutput() IoTDeviceInfoResponsePtrOutput {
	return i.ToIoTDeviceInfoResponsePtrOutputWithContext(context.Background())
}

func (i *ioTDeviceInfoResponsePtrType) ToIoTDeviceInfoResponsePtrOutputWithContext(ctx context.Context) IoTDeviceInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTDeviceInfoResponsePtrOutput)
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

func (o IoTDeviceInfoResponseOutput) ToIoTDeviceInfoResponsePtrOutput() IoTDeviceInfoResponsePtrOutput {
	return o.ToIoTDeviceInfoResponsePtrOutputWithContext(context.Background())
}

func (o IoTDeviceInfoResponseOutput) ToIoTDeviceInfoResponsePtrOutputWithContext(ctx context.Context) IoTDeviceInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IoTDeviceInfoResponse) *IoTDeviceInfoResponse {
		return &v
	}).(IoTDeviceInfoResponsePtrOutput)
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

type IoTDeviceInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (IoTDeviceInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTDeviceInfoResponse)(nil)).Elem()
}

func (o IoTDeviceInfoResponsePtrOutput) ToIoTDeviceInfoResponsePtrOutput() IoTDeviceInfoResponsePtrOutput {
	return o
}

func (o IoTDeviceInfoResponsePtrOutput) ToIoTDeviceInfoResponsePtrOutputWithContext(ctx context.Context) IoTDeviceInfoResponsePtrOutput {
	return o
}

func (o IoTDeviceInfoResponsePtrOutput) Elem() IoTDeviceInfoResponseOutput {
	return o.ApplyT(func(v *IoTDeviceInfoResponse) IoTDeviceInfoResponse {
		if v != nil {
			return *v
		}
		var ret IoTDeviceInfoResponse
		return ret
	}).(IoTDeviceInfoResponseOutput)
}

func (o IoTDeviceInfoResponsePtrOutput) Authentication() AuthenticationResponsePtrOutput {
	return o.ApplyT(func(v *IoTDeviceInfoResponse) *AuthenticationResponse {
		if v == nil {
			return nil
		}
		return v.Authentication
	}).(AuthenticationResponsePtrOutput)
}

func (o IoTDeviceInfoResponsePtrOutput) DeviceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTDeviceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DeviceId
	}).(pulumi.StringPtrOutput)
}

func (o IoTDeviceInfoResponsePtrOutput) IoTHostHub() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTDeviceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.IoTHostHub
	}).(pulumi.StringPtrOutput)
}

func (o IoTDeviceInfoResponsePtrOutput) IoTHostHubId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTDeviceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.IoTHostHubId
	}).(pulumi.StringPtrOutput)
}

type IoTEdgeAgentInfo struct {
	ImageName       string                     `pulumi:"imageName"`
	ImageRepository *ImageRepositoryCredential `pulumi:"imageRepository"`
	Tag             string                     `pulumi:"tag"`
}





type IoTEdgeAgentInfoInput interface {
	pulumi.Input

	ToIoTEdgeAgentInfoOutput() IoTEdgeAgentInfoOutput
	ToIoTEdgeAgentInfoOutputWithContext(context.Context) IoTEdgeAgentInfoOutput
}

type IoTEdgeAgentInfoArgs struct {
	ImageName       pulumi.StringInput                `pulumi:"imageName"`
	ImageRepository ImageRepositoryCredentialPtrInput `pulumi:"imageRepository"`
	Tag             pulumi.StringInput                `pulumi:"tag"`
}

func (IoTEdgeAgentInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTEdgeAgentInfo)(nil)).Elem()
}

func (i IoTEdgeAgentInfoArgs) ToIoTEdgeAgentInfoOutput() IoTEdgeAgentInfoOutput {
	return i.ToIoTEdgeAgentInfoOutputWithContext(context.Background())
}

func (i IoTEdgeAgentInfoArgs) ToIoTEdgeAgentInfoOutputWithContext(ctx context.Context) IoTEdgeAgentInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTEdgeAgentInfoOutput)
}

func (i IoTEdgeAgentInfoArgs) ToIoTEdgeAgentInfoPtrOutput() IoTEdgeAgentInfoPtrOutput {
	return i.ToIoTEdgeAgentInfoPtrOutputWithContext(context.Background())
}

func (i IoTEdgeAgentInfoArgs) ToIoTEdgeAgentInfoPtrOutputWithContext(ctx context.Context) IoTEdgeAgentInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTEdgeAgentInfoOutput).ToIoTEdgeAgentInfoPtrOutputWithContext(ctx)
}









type IoTEdgeAgentInfoPtrInput interface {
	pulumi.Input

	ToIoTEdgeAgentInfoPtrOutput() IoTEdgeAgentInfoPtrOutput
	ToIoTEdgeAgentInfoPtrOutputWithContext(context.Context) IoTEdgeAgentInfoPtrOutput
}

type ioTEdgeAgentInfoPtrType IoTEdgeAgentInfoArgs

func IoTEdgeAgentInfoPtr(v *IoTEdgeAgentInfoArgs) IoTEdgeAgentInfoPtrInput {
	return (*ioTEdgeAgentInfoPtrType)(v)
}

func (*ioTEdgeAgentInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTEdgeAgentInfo)(nil)).Elem()
}

func (i *ioTEdgeAgentInfoPtrType) ToIoTEdgeAgentInfoPtrOutput() IoTEdgeAgentInfoPtrOutput {
	return i.ToIoTEdgeAgentInfoPtrOutputWithContext(context.Background())
}

func (i *ioTEdgeAgentInfoPtrType) ToIoTEdgeAgentInfoPtrOutputWithContext(ctx context.Context) IoTEdgeAgentInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTEdgeAgentInfoPtrOutput)
}

type IoTEdgeAgentInfoOutput struct{ *pulumi.OutputState }

func (IoTEdgeAgentInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTEdgeAgentInfo)(nil)).Elem()
}

func (o IoTEdgeAgentInfoOutput) ToIoTEdgeAgentInfoOutput() IoTEdgeAgentInfoOutput {
	return o
}

func (o IoTEdgeAgentInfoOutput) ToIoTEdgeAgentInfoOutputWithContext(ctx context.Context) IoTEdgeAgentInfoOutput {
	return o
}

func (o IoTEdgeAgentInfoOutput) ToIoTEdgeAgentInfoPtrOutput() IoTEdgeAgentInfoPtrOutput {
	return o.ToIoTEdgeAgentInfoPtrOutputWithContext(context.Background())
}

func (o IoTEdgeAgentInfoOutput) ToIoTEdgeAgentInfoPtrOutputWithContext(ctx context.Context) IoTEdgeAgentInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IoTEdgeAgentInfo) *IoTEdgeAgentInfo {
		return &v
	}).(IoTEdgeAgentInfoPtrOutput)
}

func (o IoTEdgeAgentInfoOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v IoTEdgeAgentInfo) string { return v.ImageName }).(pulumi.StringOutput)
}

func (o IoTEdgeAgentInfoOutput) ImageRepository() ImageRepositoryCredentialPtrOutput {
	return o.ApplyT(func(v IoTEdgeAgentInfo) *ImageRepositoryCredential { return v.ImageRepository }).(ImageRepositoryCredentialPtrOutput)
}

func (o IoTEdgeAgentInfoOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v IoTEdgeAgentInfo) string { return v.Tag }).(pulumi.StringOutput)
}

type IoTEdgeAgentInfoPtrOutput struct{ *pulumi.OutputState }

func (IoTEdgeAgentInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTEdgeAgentInfo)(nil)).Elem()
}

func (o IoTEdgeAgentInfoPtrOutput) ToIoTEdgeAgentInfoPtrOutput() IoTEdgeAgentInfoPtrOutput {
	return o
}

func (o IoTEdgeAgentInfoPtrOutput) ToIoTEdgeAgentInfoPtrOutputWithContext(ctx context.Context) IoTEdgeAgentInfoPtrOutput {
	return o
}

func (o IoTEdgeAgentInfoPtrOutput) Elem() IoTEdgeAgentInfoOutput {
	return o.ApplyT(func(v *IoTEdgeAgentInfo) IoTEdgeAgentInfo {
		if v != nil {
			return *v
		}
		var ret IoTEdgeAgentInfo
		return ret
	}).(IoTEdgeAgentInfoOutput)
}

func (o IoTEdgeAgentInfoPtrOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTEdgeAgentInfo) *string {
		if v == nil {
			return nil
		}
		return &v.ImageName
	}).(pulumi.StringPtrOutput)
}

func (o IoTEdgeAgentInfoPtrOutput) ImageRepository() ImageRepositoryCredentialPtrOutput {
	return o.ApplyT(func(v *IoTEdgeAgentInfo) *ImageRepositoryCredential {
		if v == nil {
			return nil
		}
		return v.ImageRepository
	}).(ImageRepositoryCredentialPtrOutput)
}

func (o IoTEdgeAgentInfoPtrOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTEdgeAgentInfo) *string {
		if v == nil {
			return nil
		}
		return &v.Tag
	}).(pulumi.StringPtrOutput)
}

type IoTEdgeAgentInfoResponse struct {
	ImageName       string                             `pulumi:"imageName"`
	ImageRepository *ImageRepositoryCredentialResponse `pulumi:"imageRepository"`
	Tag             string                             `pulumi:"tag"`
}





type IoTEdgeAgentInfoResponseInput interface {
	pulumi.Input

	ToIoTEdgeAgentInfoResponseOutput() IoTEdgeAgentInfoResponseOutput
	ToIoTEdgeAgentInfoResponseOutputWithContext(context.Context) IoTEdgeAgentInfoResponseOutput
}

type IoTEdgeAgentInfoResponseArgs struct {
	ImageName       pulumi.StringInput                        `pulumi:"imageName"`
	ImageRepository ImageRepositoryCredentialResponsePtrInput `pulumi:"imageRepository"`
	Tag             pulumi.StringInput                        `pulumi:"tag"`
}

func (IoTEdgeAgentInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTEdgeAgentInfoResponse)(nil)).Elem()
}

func (i IoTEdgeAgentInfoResponseArgs) ToIoTEdgeAgentInfoResponseOutput() IoTEdgeAgentInfoResponseOutput {
	return i.ToIoTEdgeAgentInfoResponseOutputWithContext(context.Background())
}

func (i IoTEdgeAgentInfoResponseArgs) ToIoTEdgeAgentInfoResponseOutputWithContext(ctx context.Context) IoTEdgeAgentInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTEdgeAgentInfoResponseOutput)
}

func (i IoTEdgeAgentInfoResponseArgs) ToIoTEdgeAgentInfoResponsePtrOutput() IoTEdgeAgentInfoResponsePtrOutput {
	return i.ToIoTEdgeAgentInfoResponsePtrOutputWithContext(context.Background())
}

func (i IoTEdgeAgentInfoResponseArgs) ToIoTEdgeAgentInfoResponsePtrOutputWithContext(ctx context.Context) IoTEdgeAgentInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTEdgeAgentInfoResponseOutput).ToIoTEdgeAgentInfoResponsePtrOutputWithContext(ctx)
}









type IoTEdgeAgentInfoResponsePtrInput interface {
	pulumi.Input

	ToIoTEdgeAgentInfoResponsePtrOutput() IoTEdgeAgentInfoResponsePtrOutput
	ToIoTEdgeAgentInfoResponsePtrOutputWithContext(context.Context) IoTEdgeAgentInfoResponsePtrOutput
}

type ioTEdgeAgentInfoResponsePtrType IoTEdgeAgentInfoResponseArgs

func IoTEdgeAgentInfoResponsePtr(v *IoTEdgeAgentInfoResponseArgs) IoTEdgeAgentInfoResponsePtrInput {
	return (*ioTEdgeAgentInfoResponsePtrType)(v)
}

func (*ioTEdgeAgentInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTEdgeAgentInfoResponse)(nil)).Elem()
}

func (i *ioTEdgeAgentInfoResponsePtrType) ToIoTEdgeAgentInfoResponsePtrOutput() IoTEdgeAgentInfoResponsePtrOutput {
	return i.ToIoTEdgeAgentInfoResponsePtrOutputWithContext(context.Background())
}

func (i *ioTEdgeAgentInfoResponsePtrType) ToIoTEdgeAgentInfoResponsePtrOutputWithContext(ctx context.Context) IoTEdgeAgentInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTEdgeAgentInfoResponsePtrOutput)
}

type IoTEdgeAgentInfoResponseOutput struct{ *pulumi.OutputState }

func (IoTEdgeAgentInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTEdgeAgentInfoResponse)(nil)).Elem()
}

func (o IoTEdgeAgentInfoResponseOutput) ToIoTEdgeAgentInfoResponseOutput() IoTEdgeAgentInfoResponseOutput {
	return o
}

func (o IoTEdgeAgentInfoResponseOutput) ToIoTEdgeAgentInfoResponseOutputWithContext(ctx context.Context) IoTEdgeAgentInfoResponseOutput {
	return o
}

func (o IoTEdgeAgentInfoResponseOutput) ToIoTEdgeAgentInfoResponsePtrOutput() IoTEdgeAgentInfoResponsePtrOutput {
	return o.ToIoTEdgeAgentInfoResponsePtrOutputWithContext(context.Background())
}

func (o IoTEdgeAgentInfoResponseOutput) ToIoTEdgeAgentInfoResponsePtrOutputWithContext(ctx context.Context) IoTEdgeAgentInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IoTEdgeAgentInfoResponse) *IoTEdgeAgentInfoResponse {
		return &v
	}).(IoTEdgeAgentInfoResponsePtrOutput)
}

func (o IoTEdgeAgentInfoResponseOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v IoTEdgeAgentInfoResponse) string { return v.ImageName }).(pulumi.StringOutput)
}

func (o IoTEdgeAgentInfoResponseOutput) ImageRepository() ImageRepositoryCredentialResponsePtrOutput {
	return o.ApplyT(func(v IoTEdgeAgentInfoResponse) *ImageRepositoryCredentialResponse { return v.ImageRepository }).(ImageRepositoryCredentialResponsePtrOutput)
}

func (o IoTEdgeAgentInfoResponseOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v IoTEdgeAgentInfoResponse) string { return v.Tag }).(pulumi.StringOutput)
}

type IoTEdgeAgentInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (IoTEdgeAgentInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTEdgeAgentInfoResponse)(nil)).Elem()
}

func (o IoTEdgeAgentInfoResponsePtrOutput) ToIoTEdgeAgentInfoResponsePtrOutput() IoTEdgeAgentInfoResponsePtrOutput {
	return o
}

func (o IoTEdgeAgentInfoResponsePtrOutput) ToIoTEdgeAgentInfoResponsePtrOutputWithContext(ctx context.Context) IoTEdgeAgentInfoResponsePtrOutput {
	return o
}

func (o IoTEdgeAgentInfoResponsePtrOutput) Elem() IoTEdgeAgentInfoResponseOutput {
	return o.ApplyT(func(v *IoTEdgeAgentInfoResponse) IoTEdgeAgentInfoResponse {
		if v != nil {
			return *v
		}
		var ret IoTEdgeAgentInfoResponse
		return ret
	}).(IoTEdgeAgentInfoResponseOutput)
}

func (o IoTEdgeAgentInfoResponsePtrOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTEdgeAgentInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ImageName
	}).(pulumi.StringPtrOutput)
}

func (o IoTEdgeAgentInfoResponsePtrOutput) ImageRepository() ImageRepositoryCredentialResponsePtrOutput {
	return o.ApplyT(func(v *IoTEdgeAgentInfoResponse) *ImageRepositoryCredentialResponse {
		if v == nil {
			return nil
		}
		return v.ImageRepository
	}).(ImageRepositoryCredentialResponsePtrOutput)
}

func (o IoTEdgeAgentInfoResponsePtrOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTEdgeAgentInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tag
	}).(pulumi.StringPtrOutput)
}

type KubernetesClusterInfo struct {
	Version string `pulumi:"version"`
}





type KubernetesClusterInfoInput interface {
	pulumi.Input

	ToKubernetesClusterInfoOutput() KubernetesClusterInfoOutput
	ToKubernetesClusterInfoOutputWithContext(context.Context) KubernetesClusterInfoOutput
}

type KubernetesClusterInfoArgs struct {
	Version pulumi.StringInput `pulumi:"version"`
}

func (KubernetesClusterInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesClusterInfo)(nil)).Elem()
}

func (i KubernetesClusterInfoArgs) ToKubernetesClusterInfoOutput() KubernetesClusterInfoOutput {
	return i.ToKubernetesClusterInfoOutputWithContext(context.Background())
}

func (i KubernetesClusterInfoArgs) ToKubernetesClusterInfoOutputWithContext(ctx context.Context) KubernetesClusterInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterInfoOutput)
}

func (i KubernetesClusterInfoArgs) ToKubernetesClusterInfoPtrOutput() KubernetesClusterInfoPtrOutput {
	return i.ToKubernetesClusterInfoPtrOutputWithContext(context.Background())
}

func (i KubernetesClusterInfoArgs) ToKubernetesClusterInfoPtrOutputWithContext(ctx context.Context) KubernetesClusterInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterInfoOutput).ToKubernetesClusterInfoPtrOutputWithContext(ctx)
}









type KubernetesClusterInfoPtrInput interface {
	pulumi.Input

	ToKubernetesClusterInfoPtrOutput() KubernetesClusterInfoPtrOutput
	ToKubernetesClusterInfoPtrOutputWithContext(context.Context) KubernetesClusterInfoPtrOutput
}

type kubernetesClusterInfoPtrType KubernetesClusterInfoArgs

func KubernetesClusterInfoPtr(v *KubernetesClusterInfoArgs) KubernetesClusterInfoPtrInput {
	return (*kubernetesClusterInfoPtrType)(v)
}

func (*kubernetesClusterInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesClusterInfo)(nil)).Elem()
}

func (i *kubernetesClusterInfoPtrType) ToKubernetesClusterInfoPtrOutput() KubernetesClusterInfoPtrOutput {
	return i.ToKubernetesClusterInfoPtrOutputWithContext(context.Background())
}

func (i *kubernetesClusterInfoPtrType) ToKubernetesClusterInfoPtrOutputWithContext(ctx context.Context) KubernetesClusterInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterInfoPtrOutput)
}

type KubernetesClusterInfoOutput struct{ *pulumi.OutputState }

func (KubernetesClusterInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesClusterInfo)(nil)).Elem()
}

func (o KubernetesClusterInfoOutput) ToKubernetesClusterInfoOutput() KubernetesClusterInfoOutput {
	return o
}

func (o KubernetesClusterInfoOutput) ToKubernetesClusterInfoOutputWithContext(ctx context.Context) KubernetesClusterInfoOutput {
	return o
}

func (o KubernetesClusterInfoOutput) ToKubernetesClusterInfoPtrOutput() KubernetesClusterInfoPtrOutput {
	return o.ToKubernetesClusterInfoPtrOutputWithContext(context.Background())
}

func (o KubernetesClusterInfoOutput) ToKubernetesClusterInfoPtrOutputWithContext(ctx context.Context) KubernetesClusterInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesClusterInfo) *KubernetesClusterInfo {
		return &v
	}).(KubernetesClusterInfoPtrOutput)
}

func (o KubernetesClusterInfoOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesClusterInfo) string { return v.Version }).(pulumi.StringOutput)
}

type KubernetesClusterInfoPtrOutput struct{ *pulumi.OutputState }

func (KubernetesClusterInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesClusterInfo)(nil)).Elem()
}

func (o KubernetesClusterInfoPtrOutput) ToKubernetesClusterInfoPtrOutput() KubernetesClusterInfoPtrOutput {
	return o
}

func (o KubernetesClusterInfoPtrOutput) ToKubernetesClusterInfoPtrOutputWithContext(ctx context.Context) KubernetesClusterInfoPtrOutput {
	return o
}

func (o KubernetesClusterInfoPtrOutput) Elem() KubernetesClusterInfoOutput {
	return o.ApplyT(func(v *KubernetesClusterInfo) KubernetesClusterInfo {
		if v != nil {
			return *v
		}
		var ret KubernetesClusterInfo
		return ret
	}).(KubernetesClusterInfoOutput)
}

func (o KubernetesClusterInfoPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesClusterInfo) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type KubernetesClusterInfoResponse struct {
	EtcdInfo EtcdInfoResponse   `pulumi:"etcdInfo"`
	Nodes    []NodeInfoResponse `pulumi:"nodes"`
	Version  string             `pulumi:"version"`
}





type KubernetesClusterInfoResponseInput interface {
	pulumi.Input

	ToKubernetesClusterInfoResponseOutput() KubernetesClusterInfoResponseOutput
	ToKubernetesClusterInfoResponseOutputWithContext(context.Context) KubernetesClusterInfoResponseOutput
}

type KubernetesClusterInfoResponseArgs struct {
	EtcdInfo EtcdInfoResponseInput      `pulumi:"etcdInfo"`
	Nodes    NodeInfoResponseArrayInput `pulumi:"nodes"`
	Version  pulumi.StringInput         `pulumi:"version"`
}

func (KubernetesClusterInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesClusterInfoResponse)(nil)).Elem()
}

func (i KubernetesClusterInfoResponseArgs) ToKubernetesClusterInfoResponseOutput() KubernetesClusterInfoResponseOutput {
	return i.ToKubernetesClusterInfoResponseOutputWithContext(context.Background())
}

func (i KubernetesClusterInfoResponseArgs) ToKubernetesClusterInfoResponseOutputWithContext(ctx context.Context) KubernetesClusterInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterInfoResponseOutput)
}

func (i KubernetesClusterInfoResponseArgs) ToKubernetesClusterInfoResponsePtrOutput() KubernetesClusterInfoResponsePtrOutput {
	return i.ToKubernetesClusterInfoResponsePtrOutputWithContext(context.Background())
}

func (i KubernetesClusterInfoResponseArgs) ToKubernetesClusterInfoResponsePtrOutputWithContext(ctx context.Context) KubernetesClusterInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterInfoResponseOutput).ToKubernetesClusterInfoResponsePtrOutputWithContext(ctx)
}









type KubernetesClusterInfoResponsePtrInput interface {
	pulumi.Input

	ToKubernetesClusterInfoResponsePtrOutput() KubernetesClusterInfoResponsePtrOutput
	ToKubernetesClusterInfoResponsePtrOutputWithContext(context.Context) KubernetesClusterInfoResponsePtrOutput
}

type kubernetesClusterInfoResponsePtrType KubernetesClusterInfoResponseArgs

func KubernetesClusterInfoResponsePtr(v *KubernetesClusterInfoResponseArgs) KubernetesClusterInfoResponsePtrInput {
	return (*kubernetesClusterInfoResponsePtrType)(v)
}

func (*kubernetesClusterInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesClusterInfoResponse)(nil)).Elem()
}

func (i *kubernetesClusterInfoResponsePtrType) ToKubernetesClusterInfoResponsePtrOutput() KubernetesClusterInfoResponsePtrOutput {
	return i.ToKubernetesClusterInfoResponsePtrOutputWithContext(context.Background())
}

func (i *kubernetesClusterInfoResponsePtrType) ToKubernetesClusterInfoResponsePtrOutputWithContext(ctx context.Context) KubernetesClusterInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterInfoResponsePtrOutput)
}

type KubernetesClusterInfoResponseOutput struct{ *pulumi.OutputState }

func (KubernetesClusterInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesClusterInfoResponse)(nil)).Elem()
}

func (o KubernetesClusterInfoResponseOutput) ToKubernetesClusterInfoResponseOutput() KubernetesClusterInfoResponseOutput {
	return o
}

func (o KubernetesClusterInfoResponseOutput) ToKubernetesClusterInfoResponseOutputWithContext(ctx context.Context) KubernetesClusterInfoResponseOutput {
	return o
}

func (o KubernetesClusterInfoResponseOutput) ToKubernetesClusterInfoResponsePtrOutput() KubernetesClusterInfoResponsePtrOutput {
	return o.ToKubernetesClusterInfoResponsePtrOutputWithContext(context.Background())
}

func (o KubernetesClusterInfoResponseOutput) ToKubernetesClusterInfoResponsePtrOutputWithContext(ctx context.Context) KubernetesClusterInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesClusterInfoResponse) *KubernetesClusterInfoResponse {
		return &v
	}).(KubernetesClusterInfoResponsePtrOutput)
}

func (o KubernetesClusterInfoResponseOutput) EtcdInfo() EtcdInfoResponseOutput {
	return o.ApplyT(func(v KubernetesClusterInfoResponse) EtcdInfoResponse { return v.EtcdInfo }).(EtcdInfoResponseOutput)
}

func (o KubernetesClusterInfoResponseOutput) Nodes() NodeInfoResponseArrayOutput {
	return o.ApplyT(func(v KubernetesClusterInfoResponse) []NodeInfoResponse { return v.Nodes }).(NodeInfoResponseArrayOutput)
}

func (o KubernetesClusterInfoResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesClusterInfoResponse) string { return v.Version }).(pulumi.StringOutput)
}

type KubernetesClusterInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (KubernetesClusterInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesClusterInfoResponse)(nil)).Elem()
}

func (o KubernetesClusterInfoResponsePtrOutput) ToKubernetesClusterInfoResponsePtrOutput() KubernetesClusterInfoResponsePtrOutput {
	return o
}

func (o KubernetesClusterInfoResponsePtrOutput) ToKubernetesClusterInfoResponsePtrOutputWithContext(ctx context.Context) KubernetesClusterInfoResponsePtrOutput {
	return o
}

func (o KubernetesClusterInfoResponsePtrOutput) Elem() KubernetesClusterInfoResponseOutput {
	return o.ApplyT(func(v *KubernetesClusterInfoResponse) KubernetesClusterInfoResponse {
		if v != nil {
			return *v
		}
		var ret KubernetesClusterInfoResponse
		return ret
	}).(KubernetesClusterInfoResponseOutput)
}

func (o KubernetesClusterInfoResponsePtrOutput) EtcdInfo() EtcdInfoResponsePtrOutput {
	return o.ApplyT(func(v *KubernetesClusterInfoResponse) *EtcdInfoResponse {
		if v == nil {
			return nil
		}
		return &v.EtcdInfo
	}).(EtcdInfoResponsePtrOutput)
}

func (o KubernetesClusterInfoResponsePtrOutput) Nodes() NodeInfoResponseArrayOutput {
	return o.ApplyT(func(v *KubernetesClusterInfoResponse) []NodeInfoResponse {
		if v == nil {
			return nil
		}
		return v.Nodes
	}).(NodeInfoResponseArrayOutput)
}

func (o KubernetesClusterInfoResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesClusterInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type KubernetesIPConfigurationResponse struct {
	IpAddress *string `pulumi:"ipAddress"`
	Port      string  `pulumi:"port"`
}





type KubernetesIPConfigurationResponseInput interface {
	pulumi.Input

	ToKubernetesIPConfigurationResponseOutput() KubernetesIPConfigurationResponseOutput
	ToKubernetesIPConfigurationResponseOutputWithContext(context.Context) KubernetesIPConfigurationResponseOutput
}

type KubernetesIPConfigurationResponseArgs struct {
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	Port      pulumi.StringInput    `pulumi:"port"`
}

func (KubernetesIPConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesIPConfigurationResponse)(nil)).Elem()
}

func (i KubernetesIPConfigurationResponseArgs) ToKubernetesIPConfigurationResponseOutput() KubernetesIPConfigurationResponseOutput {
	return i.ToKubernetesIPConfigurationResponseOutputWithContext(context.Background())
}

func (i KubernetesIPConfigurationResponseArgs) ToKubernetesIPConfigurationResponseOutputWithContext(ctx context.Context) KubernetesIPConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesIPConfigurationResponseOutput)
}





type KubernetesIPConfigurationResponseArrayInput interface {
	pulumi.Input

	ToKubernetesIPConfigurationResponseArrayOutput() KubernetesIPConfigurationResponseArrayOutput
	ToKubernetesIPConfigurationResponseArrayOutputWithContext(context.Context) KubernetesIPConfigurationResponseArrayOutput
}

type KubernetesIPConfigurationResponseArray []KubernetesIPConfigurationResponseInput

func (KubernetesIPConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesIPConfigurationResponse)(nil)).Elem()
}

func (i KubernetesIPConfigurationResponseArray) ToKubernetesIPConfigurationResponseArrayOutput() KubernetesIPConfigurationResponseArrayOutput {
	return i.ToKubernetesIPConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i KubernetesIPConfigurationResponseArray) ToKubernetesIPConfigurationResponseArrayOutputWithContext(ctx context.Context) KubernetesIPConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesIPConfigurationResponseArrayOutput)
}

type KubernetesIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (KubernetesIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesIPConfigurationResponse)(nil)).Elem()
}

func (o KubernetesIPConfigurationResponseOutput) ToKubernetesIPConfigurationResponseOutput() KubernetesIPConfigurationResponseOutput {
	return o
}

func (o KubernetesIPConfigurationResponseOutput) ToKubernetesIPConfigurationResponseOutputWithContext(ctx context.Context) KubernetesIPConfigurationResponseOutput {
	return o
}

func (o KubernetesIPConfigurationResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesIPConfigurationResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o KubernetesIPConfigurationResponseOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesIPConfigurationResponse) string { return v.Port }).(pulumi.StringOutput)
}

type KubernetesIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (KubernetesIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesIPConfigurationResponse)(nil)).Elem()
}

func (o KubernetesIPConfigurationResponseArrayOutput) ToKubernetesIPConfigurationResponseArrayOutput() KubernetesIPConfigurationResponseArrayOutput {
	return o
}

func (o KubernetesIPConfigurationResponseArrayOutput) ToKubernetesIPConfigurationResponseArrayOutputWithContext(ctx context.Context) KubernetesIPConfigurationResponseArrayOutput {
	return o
}

func (o KubernetesIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) KubernetesIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KubernetesIPConfigurationResponse {
		return vs[0].([]KubernetesIPConfigurationResponse)[vs[1].(int)]
	}).(KubernetesIPConfigurationResponseOutput)
}

type KubernetesRoleCompute struct {
	VmProfile string `pulumi:"vmProfile"`
}





type KubernetesRoleComputeInput interface {
	pulumi.Input

	ToKubernetesRoleComputeOutput() KubernetesRoleComputeOutput
	ToKubernetesRoleComputeOutputWithContext(context.Context) KubernetesRoleComputeOutput
}

type KubernetesRoleComputeArgs struct {
	VmProfile pulumi.StringInput `pulumi:"vmProfile"`
}

func (KubernetesRoleComputeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleCompute)(nil)).Elem()
}

func (i KubernetesRoleComputeArgs) ToKubernetesRoleComputeOutput() KubernetesRoleComputeOutput {
	return i.ToKubernetesRoleComputeOutputWithContext(context.Background())
}

func (i KubernetesRoleComputeArgs) ToKubernetesRoleComputeOutputWithContext(ctx context.Context) KubernetesRoleComputeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleComputeOutput)
}

func (i KubernetesRoleComputeArgs) ToKubernetesRoleComputePtrOutput() KubernetesRoleComputePtrOutput {
	return i.ToKubernetesRoleComputePtrOutputWithContext(context.Background())
}

func (i KubernetesRoleComputeArgs) ToKubernetesRoleComputePtrOutputWithContext(ctx context.Context) KubernetesRoleComputePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleComputeOutput).ToKubernetesRoleComputePtrOutputWithContext(ctx)
}









type KubernetesRoleComputePtrInput interface {
	pulumi.Input

	ToKubernetesRoleComputePtrOutput() KubernetesRoleComputePtrOutput
	ToKubernetesRoleComputePtrOutputWithContext(context.Context) KubernetesRoleComputePtrOutput
}

type kubernetesRoleComputePtrType KubernetesRoleComputeArgs

func KubernetesRoleComputePtr(v *KubernetesRoleComputeArgs) KubernetesRoleComputePtrInput {
	return (*kubernetesRoleComputePtrType)(v)
}

func (*kubernetesRoleComputePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleCompute)(nil)).Elem()
}

func (i *kubernetesRoleComputePtrType) ToKubernetesRoleComputePtrOutput() KubernetesRoleComputePtrOutput {
	return i.ToKubernetesRoleComputePtrOutputWithContext(context.Background())
}

func (i *kubernetesRoleComputePtrType) ToKubernetesRoleComputePtrOutputWithContext(ctx context.Context) KubernetesRoleComputePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleComputePtrOutput)
}

type KubernetesRoleComputeOutput struct{ *pulumi.OutputState }

func (KubernetesRoleComputeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleCompute)(nil)).Elem()
}

func (o KubernetesRoleComputeOutput) ToKubernetesRoleComputeOutput() KubernetesRoleComputeOutput {
	return o
}

func (o KubernetesRoleComputeOutput) ToKubernetesRoleComputeOutputWithContext(ctx context.Context) KubernetesRoleComputeOutput {
	return o
}

func (o KubernetesRoleComputeOutput) ToKubernetesRoleComputePtrOutput() KubernetesRoleComputePtrOutput {
	return o.ToKubernetesRoleComputePtrOutputWithContext(context.Background())
}

func (o KubernetesRoleComputeOutput) ToKubernetesRoleComputePtrOutputWithContext(ctx context.Context) KubernetesRoleComputePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesRoleCompute) *KubernetesRoleCompute {
		return &v
	}).(KubernetesRoleComputePtrOutput)
}

func (o KubernetesRoleComputeOutput) VmProfile() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesRoleCompute) string { return v.VmProfile }).(pulumi.StringOutput)
}

type KubernetesRoleComputePtrOutput struct{ *pulumi.OutputState }

func (KubernetesRoleComputePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleCompute)(nil)).Elem()
}

func (o KubernetesRoleComputePtrOutput) ToKubernetesRoleComputePtrOutput() KubernetesRoleComputePtrOutput {
	return o
}

func (o KubernetesRoleComputePtrOutput) ToKubernetesRoleComputePtrOutputWithContext(ctx context.Context) KubernetesRoleComputePtrOutput {
	return o
}

func (o KubernetesRoleComputePtrOutput) Elem() KubernetesRoleComputeOutput {
	return o.ApplyT(func(v *KubernetesRoleCompute) KubernetesRoleCompute {
		if v != nil {
			return *v
		}
		var ret KubernetesRoleCompute
		return ret
	}).(KubernetesRoleComputeOutput)
}

func (o KubernetesRoleComputePtrOutput) VmProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesRoleCompute) *string {
		if v == nil {
			return nil
		}
		return &v.VmProfile
	}).(pulumi.StringPtrOutput)
}

type KubernetesRoleComputeResponse struct {
	MemoryInBytes  float64 `pulumi:"memoryInBytes"`
	ProcessorCount int     `pulumi:"processorCount"`
	VmProfile      string  `pulumi:"vmProfile"`
}





type KubernetesRoleComputeResponseInput interface {
	pulumi.Input

	ToKubernetesRoleComputeResponseOutput() KubernetesRoleComputeResponseOutput
	ToKubernetesRoleComputeResponseOutputWithContext(context.Context) KubernetesRoleComputeResponseOutput
}

type KubernetesRoleComputeResponseArgs struct {
	MemoryInBytes  pulumi.Float64Input `pulumi:"memoryInBytes"`
	ProcessorCount pulumi.IntInput     `pulumi:"processorCount"`
	VmProfile      pulumi.StringInput  `pulumi:"vmProfile"`
}

func (KubernetesRoleComputeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleComputeResponse)(nil)).Elem()
}

func (i KubernetesRoleComputeResponseArgs) ToKubernetesRoleComputeResponseOutput() KubernetesRoleComputeResponseOutput {
	return i.ToKubernetesRoleComputeResponseOutputWithContext(context.Background())
}

func (i KubernetesRoleComputeResponseArgs) ToKubernetesRoleComputeResponseOutputWithContext(ctx context.Context) KubernetesRoleComputeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleComputeResponseOutput)
}

func (i KubernetesRoleComputeResponseArgs) ToKubernetesRoleComputeResponsePtrOutput() KubernetesRoleComputeResponsePtrOutput {
	return i.ToKubernetesRoleComputeResponsePtrOutputWithContext(context.Background())
}

func (i KubernetesRoleComputeResponseArgs) ToKubernetesRoleComputeResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleComputeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleComputeResponseOutput).ToKubernetesRoleComputeResponsePtrOutputWithContext(ctx)
}









type KubernetesRoleComputeResponsePtrInput interface {
	pulumi.Input

	ToKubernetesRoleComputeResponsePtrOutput() KubernetesRoleComputeResponsePtrOutput
	ToKubernetesRoleComputeResponsePtrOutputWithContext(context.Context) KubernetesRoleComputeResponsePtrOutput
}

type kubernetesRoleComputeResponsePtrType KubernetesRoleComputeResponseArgs

func KubernetesRoleComputeResponsePtr(v *KubernetesRoleComputeResponseArgs) KubernetesRoleComputeResponsePtrInput {
	return (*kubernetesRoleComputeResponsePtrType)(v)
}

func (*kubernetesRoleComputeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleComputeResponse)(nil)).Elem()
}

func (i *kubernetesRoleComputeResponsePtrType) ToKubernetesRoleComputeResponsePtrOutput() KubernetesRoleComputeResponsePtrOutput {
	return i.ToKubernetesRoleComputeResponsePtrOutputWithContext(context.Background())
}

func (i *kubernetesRoleComputeResponsePtrType) ToKubernetesRoleComputeResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleComputeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleComputeResponsePtrOutput)
}

type KubernetesRoleComputeResponseOutput struct{ *pulumi.OutputState }

func (KubernetesRoleComputeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleComputeResponse)(nil)).Elem()
}

func (o KubernetesRoleComputeResponseOutput) ToKubernetesRoleComputeResponseOutput() KubernetesRoleComputeResponseOutput {
	return o
}

func (o KubernetesRoleComputeResponseOutput) ToKubernetesRoleComputeResponseOutputWithContext(ctx context.Context) KubernetesRoleComputeResponseOutput {
	return o
}

func (o KubernetesRoleComputeResponseOutput) ToKubernetesRoleComputeResponsePtrOutput() KubernetesRoleComputeResponsePtrOutput {
	return o.ToKubernetesRoleComputeResponsePtrOutputWithContext(context.Background())
}

func (o KubernetesRoleComputeResponseOutput) ToKubernetesRoleComputeResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleComputeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesRoleComputeResponse) *KubernetesRoleComputeResponse {
		return &v
	}).(KubernetesRoleComputeResponsePtrOutput)
}

func (o KubernetesRoleComputeResponseOutput) MemoryInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v KubernetesRoleComputeResponse) float64 { return v.MemoryInBytes }).(pulumi.Float64Output)
}

func (o KubernetesRoleComputeResponseOutput) ProcessorCount() pulumi.IntOutput {
	return o.ApplyT(func(v KubernetesRoleComputeResponse) int { return v.ProcessorCount }).(pulumi.IntOutput)
}

func (o KubernetesRoleComputeResponseOutput) VmProfile() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesRoleComputeResponse) string { return v.VmProfile }).(pulumi.StringOutput)
}

type KubernetesRoleComputeResponsePtrOutput struct{ *pulumi.OutputState }

func (KubernetesRoleComputeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleComputeResponse)(nil)).Elem()
}

func (o KubernetesRoleComputeResponsePtrOutput) ToKubernetesRoleComputeResponsePtrOutput() KubernetesRoleComputeResponsePtrOutput {
	return o
}

func (o KubernetesRoleComputeResponsePtrOutput) ToKubernetesRoleComputeResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleComputeResponsePtrOutput {
	return o
}

func (o KubernetesRoleComputeResponsePtrOutput) Elem() KubernetesRoleComputeResponseOutput {
	return o.ApplyT(func(v *KubernetesRoleComputeResponse) KubernetesRoleComputeResponse {
		if v != nil {
			return *v
		}
		var ret KubernetesRoleComputeResponse
		return ret
	}).(KubernetesRoleComputeResponseOutput)
}

func (o KubernetesRoleComputeResponsePtrOutput) MemoryInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KubernetesRoleComputeResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.MemoryInBytes
	}).(pulumi.Float64PtrOutput)
}

func (o KubernetesRoleComputeResponsePtrOutput) ProcessorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubernetesRoleComputeResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ProcessorCount
	}).(pulumi.IntPtrOutput)
}

func (o KubernetesRoleComputeResponsePtrOutput) VmProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesRoleComputeResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VmProfile
	}).(pulumi.StringPtrOutput)
}

type KubernetesRoleNetworkResponse struct {
	CniConfig          CniConfigResponse          `pulumi:"cniConfig"`
	LoadBalancerConfig LoadBalancerConfigResponse `pulumi:"loadBalancerConfig"`
}





type KubernetesRoleNetworkResponseInput interface {
	pulumi.Input

	ToKubernetesRoleNetworkResponseOutput() KubernetesRoleNetworkResponseOutput
	ToKubernetesRoleNetworkResponseOutputWithContext(context.Context) KubernetesRoleNetworkResponseOutput
}

type KubernetesRoleNetworkResponseArgs struct {
	CniConfig          CniConfigResponseInput          `pulumi:"cniConfig"`
	LoadBalancerConfig LoadBalancerConfigResponseInput `pulumi:"loadBalancerConfig"`
}

func (KubernetesRoleNetworkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleNetworkResponse)(nil)).Elem()
}

func (i KubernetesRoleNetworkResponseArgs) ToKubernetesRoleNetworkResponseOutput() KubernetesRoleNetworkResponseOutput {
	return i.ToKubernetesRoleNetworkResponseOutputWithContext(context.Background())
}

func (i KubernetesRoleNetworkResponseArgs) ToKubernetesRoleNetworkResponseOutputWithContext(ctx context.Context) KubernetesRoleNetworkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleNetworkResponseOutput)
}

func (i KubernetesRoleNetworkResponseArgs) ToKubernetesRoleNetworkResponsePtrOutput() KubernetesRoleNetworkResponsePtrOutput {
	return i.ToKubernetesRoleNetworkResponsePtrOutputWithContext(context.Background())
}

func (i KubernetesRoleNetworkResponseArgs) ToKubernetesRoleNetworkResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleNetworkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleNetworkResponseOutput).ToKubernetesRoleNetworkResponsePtrOutputWithContext(ctx)
}









type KubernetesRoleNetworkResponsePtrInput interface {
	pulumi.Input

	ToKubernetesRoleNetworkResponsePtrOutput() KubernetesRoleNetworkResponsePtrOutput
	ToKubernetesRoleNetworkResponsePtrOutputWithContext(context.Context) KubernetesRoleNetworkResponsePtrOutput
}

type kubernetesRoleNetworkResponsePtrType KubernetesRoleNetworkResponseArgs

func KubernetesRoleNetworkResponsePtr(v *KubernetesRoleNetworkResponseArgs) KubernetesRoleNetworkResponsePtrInput {
	return (*kubernetesRoleNetworkResponsePtrType)(v)
}

func (*kubernetesRoleNetworkResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleNetworkResponse)(nil)).Elem()
}

func (i *kubernetesRoleNetworkResponsePtrType) ToKubernetesRoleNetworkResponsePtrOutput() KubernetesRoleNetworkResponsePtrOutput {
	return i.ToKubernetesRoleNetworkResponsePtrOutputWithContext(context.Background())
}

func (i *kubernetesRoleNetworkResponsePtrType) ToKubernetesRoleNetworkResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleNetworkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleNetworkResponsePtrOutput)
}

type KubernetesRoleNetworkResponseOutput struct{ *pulumi.OutputState }

func (KubernetesRoleNetworkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleNetworkResponse)(nil)).Elem()
}

func (o KubernetesRoleNetworkResponseOutput) ToKubernetesRoleNetworkResponseOutput() KubernetesRoleNetworkResponseOutput {
	return o
}

func (o KubernetesRoleNetworkResponseOutput) ToKubernetesRoleNetworkResponseOutputWithContext(ctx context.Context) KubernetesRoleNetworkResponseOutput {
	return o
}

func (o KubernetesRoleNetworkResponseOutput) ToKubernetesRoleNetworkResponsePtrOutput() KubernetesRoleNetworkResponsePtrOutput {
	return o.ToKubernetesRoleNetworkResponsePtrOutputWithContext(context.Background())
}

func (o KubernetesRoleNetworkResponseOutput) ToKubernetesRoleNetworkResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleNetworkResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesRoleNetworkResponse) *KubernetesRoleNetworkResponse {
		return &v
	}).(KubernetesRoleNetworkResponsePtrOutput)
}

func (o KubernetesRoleNetworkResponseOutput) CniConfig() CniConfigResponseOutput {
	return o.ApplyT(func(v KubernetesRoleNetworkResponse) CniConfigResponse { return v.CniConfig }).(CniConfigResponseOutput)
}

func (o KubernetesRoleNetworkResponseOutput) LoadBalancerConfig() LoadBalancerConfigResponseOutput {
	return o.ApplyT(func(v KubernetesRoleNetworkResponse) LoadBalancerConfigResponse { return v.LoadBalancerConfig }).(LoadBalancerConfigResponseOutput)
}

type KubernetesRoleNetworkResponsePtrOutput struct{ *pulumi.OutputState }

func (KubernetesRoleNetworkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleNetworkResponse)(nil)).Elem()
}

func (o KubernetesRoleNetworkResponsePtrOutput) ToKubernetesRoleNetworkResponsePtrOutput() KubernetesRoleNetworkResponsePtrOutput {
	return o
}

func (o KubernetesRoleNetworkResponsePtrOutput) ToKubernetesRoleNetworkResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleNetworkResponsePtrOutput {
	return o
}

func (o KubernetesRoleNetworkResponsePtrOutput) Elem() KubernetesRoleNetworkResponseOutput {
	return o.ApplyT(func(v *KubernetesRoleNetworkResponse) KubernetesRoleNetworkResponse {
		if v != nil {
			return *v
		}
		var ret KubernetesRoleNetworkResponse
		return ret
	}).(KubernetesRoleNetworkResponseOutput)
}

func (o KubernetesRoleNetworkResponsePtrOutput) CniConfig() CniConfigResponsePtrOutput {
	return o.ApplyT(func(v *KubernetesRoleNetworkResponse) *CniConfigResponse {
		if v == nil {
			return nil
		}
		return &v.CniConfig
	}).(CniConfigResponsePtrOutput)
}

func (o KubernetesRoleNetworkResponsePtrOutput) LoadBalancerConfig() LoadBalancerConfigResponsePtrOutput {
	return o.ApplyT(func(v *KubernetesRoleNetworkResponse) *LoadBalancerConfigResponse {
		if v == nil {
			return nil
		}
		return &v.LoadBalancerConfig
	}).(LoadBalancerConfigResponsePtrOutput)
}

type KubernetesRoleResources struct {
	Compute KubernetesRoleCompute  `pulumi:"compute"`
	Storage *KubernetesRoleStorage `pulumi:"storage"`
}





type KubernetesRoleResourcesInput interface {
	pulumi.Input

	ToKubernetesRoleResourcesOutput() KubernetesRoleResourcesOutput
	ToKubernetesRoleResourcesOutputWithContext(context.Context) KubernetesRoleResourcesOutput
}

type KubernetesRoleResourcesArgs struct {
	Compute KubernetesRoleComputeInput    `pulumi:"compute"`
	Storage KubernetesRoleStoragePtrInput `pulumi:"storage"`
}

func (KubernetesRoleResourcesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleResources)(nil)).Elem()
}

func (i KubernetesRoleResourcesArgs) ToKubernetesRoleResourcesOutput() KubernetesRoleResourcesOutput {
	return i.ToKubernetesRoleResourcesOutputWithContext(context.Background())
}

func (i KubernetesRoleResourcesArgs) ToKubernetesRoleResourcesOutputWithContext(ctx context.Context) KubernetesRoleResourcesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleResourcesOutput)
}

func (i KubernetesRoleResourcesArgs) ToKubernetesRoleResourcesPtrOutput() KubernetesRoleResourcesPtrOutput {
	return i.ToKubernetesRoleResourcesPtrOutputWithContext(context.Background())
}

func (i KubernetesRoleResourcesArgs) ToKubernetesRoleResourcesPtrOutputWithContext(ctx context.Context) KubernetesRoleResourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleResourcesOutput).ToKubernetesRoleResourcesPtrOutputWithContext(ctx)
}









type KubernetesRoleResourcesPtrInput interface {
	pulumi.Input

	ToKubernetesRoleResourcesPtrOutput() KubernetesRoleResourcesPtrOutput
	ToKubernetesRoleResourcesPtrOutputWithContext(context.Context) KubernetesRoleResourcesPtrOutput
}

type kubernetesRoleResourcesPtrType KubernetesRoleResourcesArgs

func KubernetesRoleResourcesPtr(v *KubernetesRoleResourcesArgs) KubernetesRoleResourcesPtrInput {
	return (*kubernetesRoleResourcesPtrType)(v)
}

func (*kubernetesRoleResourcesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleResources)(nil)).Elem()
}

func (i *kubernetesRoleResourcesPtrType) ToKubernetesRoleResourcesPtrOutput() KubernetesRoleResourcesPtrOutput {
	return i.ToKubernetesRoleResourcesPtrOutputWithContext(context.Background())
}

func (i *kubernetesRoleResourcesPtrType) ToKubernetesRoleResourcesPtrOutputWithContext(ctx context.Context) KubernetesRoleResourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleResourcesPtrOutput)
}

type KubernetesRoleResourcesOutput struct{ *pulumi.OutputState }

func (KubernetesRoleResourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleResources)(nil)).Elem()
}

func (o KubernetesRoleResourcesOutput) ToKubernetesRoleResourcesOutput() KubernetesRoleResourcesOutput {
	return o
}

func (o KubernetesRoleResourcesOutput) ToKubernetesRoleResourcesOutputWithContext(ctx context.Context) KubernetesRoleResourcesOutput {
	return o
}

func (o KubernetesRoleResourcesOutput) ToKubernetesRoleResourcesPtrOutput() KubernetesRoleResourcesPtrOutput {
	return o.ToKubernetesRoleResourcesPtrOutputWithContext(context.Background())
}

func (o KubernetesRoleResourcesOutput) ToKubernetesRoleResourcesPtrOutputWithContext(ctx context.Context) KubernetesRoleResourcesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesRoleResources) *KubernetesRoleResources {
		return &v
	}).(KubernetesRoleResourcesPtrOutput)
}

func (o KubernetesRoleResourcesOutput) Compute() KubernetesRoleComputeOutput {
	return o.ApplyT(func(v KubernetesRoleResources) KubernetesRoleCompute { return v.Compute }).(KubernetesRoleComputeOutput)
}

func (o KubernetesRoleResourcesOutput) Storage() KubernetesRoleStoragePtrOutput {
	return o.ApplyT(func(v KubernetesRoleResources) *KubernetesRoleStorage { return v.Storage }).(KubernetesRoleStoragePtrOutput)
}

type KubernetesRoleResourcesPtrOutput struct{ *pulumi.OutputState }

func (KubernetesRoleResourcesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleResources)(nil)).Elem()
}

func (o KubernetesRoleResourcesPtrOutput) ToKubernetesRoleResourcesPtrOutput() KubernetesRoleResourcesPtrOutput {
	return o
}

func (o KubernetesRoleResourcesPtrOutput) ToKubernetesRoleResourcesPtrOutputWithContext(ctx context.Context) KubernetesRoleResourcesPtrOutput {
	return o
}

func (o KubernetesRoleResourcesPtrOutput) Elem() KubernetesRoleResourcesOutput {
	return o.ApplyT(func(v *KubernetesRoleResources) KubernetesRoleResources {
		if v != nil {
			return *v
		}
		var ret KubernetesRoleResources
		return ret
	}).(KubernetesRoleResourcesOutput)
}

func (o KubernetesRoleResourcesPtrOutput) Compute() KubernetesRoleComputePtrOutput {
	return o.ApplyT(func(v *KubernetesRoleResources) *KubernetesRoleCompute {
		if v == nil {
			return nil
		}
		return &v.Compute
	}).(KubernetesRoleComputePtrOutput)
}

func (o KubernetesRoleResourcesPtrOutput) Storage() KubernetesRoleStoragePtrOutput {
	return o.ApplyT(func(v *KubernetesRoleResources) *KubernetesRoleStorage {
		if v == nil {
			return nil
		}
		return v.Storage
	}).(KubernetesRoleStoragePtrOutput)
}

type KubernetesRoleResourcesResponse struct {
	Compute KubernetesRoleComputeResponse  `pulumi:"compute"`
	Network KubernetesRoleNetworkResponse  `pulumi:"network"`
	Storage *KubernetesRoleStorageResponse `pulumi:"storage"`
}





type KubernetesRoleResourcesResponseInput interface {
	pulumi.Input

	ToKubernetesRoleResourcesResponseOutput() KubernetesRoleResourcesResponseOutput
	ToKubernetesRoleResourcesResponseOutputWithContext(context.Context) KubernetesRoleResourcesResponseOutput
}

type KubernetesRoleResourcesResponseArgs struct {
	Compute KubernetesRoleComputeResponseInput    `pulumi:"compute"`
	Network KubernetesRoleNetworkResponseInput    `pulumi:"network"`
	Storage KubernetesRoleStorageResponsePtrInput `pulumi:"storage"`
}

func (KubernetesRoleResourcesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleResourcesResponse)(nil)).Elem()
}

func (i KubernetesRoleResourcesResponseArgs) ToKubernetesRoleResourcesResponseOutput() KubernetesRoleResourcesResponseOutput {
	return i.ToKubernetesRoleResourcesResponseOutputWithContext(context.Background())
}

func (i KubernetesRoleResourcesResponseArgs) ToKubernetesRoleResourcesResponseOutputWithContext(ctx context.Context) KubernetesRoleResourcesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleResourcesResponseOutput)
}

func (i KubernetesRoleResourcesResponseArgs) ToKubernetesRoleResourcesResponsePtrOutput() KubernetesRoleResourcesResponsePtrOutput {
	return i.ToKubernetesRoleResourcesResponsePtrOutputWithContext(context.Background())
}

func (i KubernetesRoleResourcesResponseArgs) ToKubernetesRoleResourcesResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleResourcesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleResourcesResponseOutput).ToKubernetesRoleResourcesResponsePtrOutputWithContext(ctx)
}









type KubernetesRoleResourcesResponsePtrInput interface {
	pulumi.Input

	ToKubernetesRoleResourcesResponsePtrOutput() KubernetesRoleResourcesResponsePtrOutput
	ToKubernetesRoleResourcesResponsePtrOutputWithContext(context.Context) KubernetesRoleResourcesResponsePtrOutput
}

type kubernetesRoleResourcesResponsePtrType KubernetesRoleResourcesResponseArgs

func KubernetesRoleResourcesResponsePtr(v *KubernetesRoleResourcesResponseArgs) KubernetesRoleResourcesResponsePtrInput {
	return (*kubernetesRoleResourcesResponsePtrType)(v)
}

func (*kubernetesRoleResourcesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleResourcesResponse)(nil)).Elem()
}

func (i *kubernetesRoleResourcesResponsePtrType) ToKubernetesRoleResourcesResponsePtrOutput() KubernetesRoleResourcesResponsePtrOutput {
	return i.ToKubernetesRoleResourcesResponsePtrOutputWithContext(context.Background())
}

func (i *kubernetesRoleResourcesResponsePtrType) ToKubernetesRoleResourcesResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleResourcesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleResourcesResponsePtrOutput)
}

type KubernetesRoleResourcesResponseOutput struct{ *pulumi.OutputState }

func (KubernetesRoleResourcesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleResourcesResponse)(nil)).Elem()
}

func (o KubernetesRoleResourcesResponseOutput) ToKubernetesRoleResourcesResponseOutput() KubernetesRoleResourcesResponseOutput {
	return o
}

func (o KubernetesRoleResourcesResponseOutput) ToKubernetesRoleResourcesResponseOutputWithContext(ctx context.Context) KubernetesRoleResourcesResponseOutput {
	return o
}

func (o KubernetesRoleResourcesResponseOutput) ToKubernetesRoleResourcesResponsePtrOutput() KubernetesRoleResourcesResponsePtrOutput {
	return o.ToKubernetesRoleResourcesResponsePtrOutputWithContext(context.Background())
}

func (o KubernetesRoleResourcesResponseOutput) ToKubernetesRoleResourcesResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleResourcesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesRoleResourcesResponse) *KubernetesRoleResourcesResponse {
		return &v
	}).(KubernetesRoleResourcesResponsePtrOutput)
}

func (o KubernetesRoleResourcesResponseOutput) Compute() KubernetesRoleComputeResponseOutput {
	return o.ApplyT(func(v KubernetesRoleResourcesResponse) KubernetesRoleComputeResponse { return v.Compute }).(KubernetesRoleComputeResponseOutput)
}

func (o KubernetesRoleResourcesResponseOutput) Network() KubernetesRoleNetworkResponseOutput {
	return o.ApplyT(func(v KubernetesRoleResourcesResponse) KubernetesRoleNetworkResponse { return v.Network }).(KubernetesRoleNetworkResponseOutput)
}

func (o KubernetesRoleResourcesResponseOutput) Storage() KubernetesRoleStorageResponsePtrOutput {
	return o.ApplyT(func(v KubernetesRoleResourcesResponse) *KubernetesRoleStorageResponse { return v.Storage }).(KubernetesRoleStorageResponsePtrOutput)
}

type KubernetesRoleResourcesResponsePtrOutput struct{ *pulumi.OutputState }

func (KubernetesRoleResourcesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleResourcesResponse)(nil)).Elem()
}

func (o KubernetesRoleResourcesResponsePtrOutput) ToKubernetesRoleResourcesResponsePtrOutput() KubernetesRoleResourcesResponsePtrOutput {
	return o
}

func (o KubernetesRoleResourcesResponsePtrOutput) ToKubernetesRoleResourcesResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleResourcesResponsePtrOutput {
	return o
}

func (o KubernetesRoleResourcesResponsePtrOutput) Elem() KubernetesRoleResourcesResponseOutput {
	return o.ApplyT(func(v *KubernetesRoleResourcesResponse) KubernetesRoleResourcesResponse {
		if v != nil {
			return *v
		}
		var ret KubernetesRoleResourcesResponse
		return ret
	}).(KubernetesRoleResourcesResponseOutput)
}

func (o KubernetesRoleResourcesResponsePtrOutput) Compute() KubernetesRoleComputeResponsePtrOutput {
	return o.ApplyT(func(v *KubernetesRoleResourcesResponse) *KubernetesRoleComputeResponse {
		if v == nil {
			return nil
		}
		return &v.Compute
	}).(KubernetesRoleComputeResponsePtrOutput)
}

func (o KubernetesRoleResourcesResponsePtrOutput) Network() KubernetesRoleNetworkResponsePtrOutput {
	return o.ApplyT(func(v *KubernetesRoleResourcesResponse) *KubernetesRoleNetworkResponse {
		if v == nil {
			return nil
		}
		return &v.Network
	}).(KubernetesRoleNetworkResponsePtrOutput)
}

func (o KubernetesRoleResourcesResponsePtrOutput) Storage() KubernetesRoleStorageResponsePtrOutput {
	return o.ApplyT(func(v *KubernetesRoleResourcesResponse) *KubernetesRoleStorageResponse {
		if v == nil {
			return nil
		}
		return v.Storage
	}).(KubernetesRoleStorageResponsePtrOutput)
}

type KubernetesRoleStorage struct {
	Endpoints []MountPointMap `pulumi:"endpoints"`
}





type KubernetesRoleStorageInput interface {
	pulumi.Input

	ToKubernetesRoleStorageOutput() KubernetesRoleStorageOutput
	ToKubernetesRoleStorageOutputWithContext(context.Context) KubernetesRoleStorageOutput
}

type KubernetesRoleStorageArgs struct {
	Endpoints MountPointMapArrayInput `pulumi:"endpoints"`
}

func (KubernetesRoleStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleStorage)(nil)).Elem()
}

func (i KubernetesRoleStorageArgs) ToKubernetesRoleStorageOutput() KubernetesRoleStorageOutput {
	return i.ToKubernetesRoleStorageOutputWithContext(context.Background())
}

func (i KubernetesRoleStorageArgs) ToKubernetesRoleStorageOutputWithContext(ctx context.Context) KubernetesRoleStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleStorageOutput)
}

func (i KubernetesRoleStorageArgs) ToKubernetesRoleStoragePtrOutput() KubernetesRoleStoragePtrOutput {
	return i.ToKubernetesRoleStoragePtrOutputWithContext(context.Background())
}

func (i KubernetesRoleStorageArgs) ToKubernetesRoleStoragePtrOutputWithContext(ctx context.Context) KubernetesRoleStoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleStorageOutput).ToKubernetesRoleStoragePtrOutputWithContext(ctx)
}









type KubernetesRoleStoragePtrInput interface {
	pulumi.Input

	ToKubernetesRoleStoragePtrOutput() KubernetesRoleStoragePtrOutput
	ToKubernetesRoleStoragePtrOutputWithContext(context.Context) KubernetesRoleStoragePtrOutput
}

type kubernetesRoleStoragePtrType KubernetesRoleStorageArgs

func KubernetesRoleStoragePtr(v *KubernetesRoleStorageArgs) KubernetesRoleStoragePtrInput {
	return (*kubernetesRoleStoragePtrType)(v)
}

func (*kubernetesRoleStoragePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleStorage)(nil)).Elem()
}

func (i *kubernetesRoleStoragePtrType) ToKubernetesRoleStoragePtrOutput() KubernetesRoleStoragePtrOutput {
	return i.ToKubernetesRoleStoragePtrOutputWithContext(context.Background())
}

func (i *kubernetesRoleStoragePtrType) ToKubernetesRoleStoragePtrOutputWithContext(ctx context.Context) KubernetesRoleStoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleStoragePtrOutput)
}

type KubernetesRoleStorageOutput struct{ *pulumi.OutputState }

func (KubernetesRoleStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleStorage)(nil)).Elem()
}

func (o KubernetesRoleStorageOutput) ToKubernetesRoleStorageOutput() KubernetesRoleStorageOutput {
	return o
}

func (o KubernetesRoleStorageOutput) ToKubernetesRoleStorageOutputWithContext(ctx context.Context) KubernetesRoleStorageOutput {
	return o
}

func (o KubernetesRoleStorageOutput) ToKubernetesRoleStoragePtrOutput() KubernetesRoleStoragePtrOutput {
	return o.ToKubernetesRoleStoragePtrOutputWithContext(context.Background())
}

func (o KubernetesRoleStorageOutput) ToKubernetesRoleStoragePtrOutputWithContext(ctx context.Context) KubernetesRoleStoragePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesRoleStorage) *KubernetesRoleStorage {
		return &v
	}).(KubernetesRoleStoragePtrOutput)
}

func (o KubernetesRoleStorageOutput) Endpoints() MountPointMapArrayOutput {
	return o.ApplyT(func(v KubernetesRoleStorage) []MountPointMap { return v.Endpoints }).(MountPointMapArrayOutput)
}

type KubernetesRoleStoragePtrOutput struct{ *pulumi.OutputState }

func (KubernetesRoleStoragePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleStorage)(nil)).Elem()
}

func (o KubernetesRoleStoragePtrOutput) ToKubernetesRoleStoragePtrOutput() KubernetesRoleStoragePtrOutput {
	return o
}

func (o KubernetesRoleStoragePtrOutput) ToKubernetesRoleStoragePtrOutputWithContext(ctx context.Context) KubernetesRoleStoragePtrOutput {
	return o
}

func (o KubernetesRoleStoragePtrOutput) Elem() KubernetesRoleStorageOutput {
	return o.ApplyT(func(v *KubernetesRoleStorage) KubernetesRoleStorage {
		if v != nil {
			return *v
		}
		var ret KubernetesRoleStorage
		return ret
	}).(KubernetesRoleStorageOutput)
}

func (o KubernetesRoleStoragePtrOutput) Endpoints() MountPointMapArrayOutput {
	return o.ApplyT(func(v *KubernetesRoleStorage) []MountPointMap {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(MountPointMapArrayOutput)
}

type KubernetesRoleStorageClassInfoResponse struct {
	Name           string `pulumi:"name"`
	PosixCompliant string `pulumi:"posixCompliant"`
	Type           string `pulumi:"type"`
}





type KubernetesRoleStorageClassInfoResponseInput interface {
	pulumi.Input

	ToKubernetesRoleStorageClassInfoResponseOutput() KubernetesRoleStorageClassInfoResponseOutput
	ToKubernetesRoleStorageClassInfoResponseOutputWithContext(context.Context) KubernetesRoleStorageClassInfoResponseOutput
}

type KubernetesRoleStorageClassInfoResponseArgs struct {
	Name           pulumi.StringInput `pulumi:"name"`
	PosixCompliant pulumi.StringInput `pulumi:"posixCompliant"`
	Type           pulumi.StringInput `pulumi:"type"`
}

func (KubernetesRoleStorageClassInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleStorageClassInfoResponse)(nil)).Elem()
}

func (i KubernetesRoleStorageClassInfoResponseArgs) ToKubernetesRoleStorageClassInfoResponseOutput() KubernetesRoleStorageClassInfoResponseOutput {
	return i.ToKubernetesRoleStorageClassInfoResponseOutputWithContext(context.Background())
}

func (i KubernetesRoleStorageClassInfoResponseArgs) ToKubernetesRoleStorageClassInfoResponseOutputWithContext(ctx context.Context) KubernetesRoleStorageClassInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleStorageClassInfoResponseOutput)
}





type KubernetesRoleStorageClassInfoResponseArrayInput interface {
	pulumi.Input

	ToKubernetesRoleStorageClassInfoResponseArrayOutput() KubernetesRoleStorageClassInfoResponseArrayOutput
	ToKubernetesRoleStorageClassInfoResponseArrayOutputWithContext(context.Context) KubernetesRoleStorageClassInfoResponseArrayOutput
}

type KubernetesRoleStorageClassInfoResponseArray []KubernetesRoleStorageClassInfoResponseInput

func (KubernetesRoleStorageClassInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesRoleStorageClassInfoResponse)(nil)).Elem()
}

func (i KubernetesRoleStorageClassInfoResponseArray) ToKubernetesRoleStorageClassInfoResponseArrayOutput() KubernetesRoleStorageClassInfoResponseArrayOutput {
	return i.ToKubernetesRoleStorageClassInfoResponseArrayOutputWithContext(context.Background())
}

func (i KubernetesRoleStorageClassInfoResponseArray) ToKubernetesRoleStorageClassInfoResponseArrayOutputWithContext(ctx context.Context) KubernetesRoleStorageClassInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleStorageClassInfoResponseArrayOutput)
}

type KubernetesRoleStorageClassInfoResponseOutput struct{ *pulumi.OutputState }

func (KubernetesRoleStorageClassInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleStorageClassInfoResponse)(nil)).Elem()
}

func (o KubernetesRoleStorageClassInfoResponseOutput) ToKubernetesRoleStorageClassInfoResponseOutput() KubernetesRoleStorageClassInfoResponseOutput {
	return o
}

func (o KubernetesRoleStorageClassInfoResponseOutput) ToKubernetesRoleStorageClassInfoResponseOutputWithContext(ctx context.Context) KubernetesRoleStorageClassInfoResponseOutput {
	return o
}

func (o KubernetesRoleStorageClassInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesRoleStorageClassInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o KubernetesRoleStorageClassInfoResponseOutput) PosixCompliant() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesRoleStorageClassInfoResponse) string { return v.PosixCompliant }).(pulumi.StringOutput)
}

func (o KubernetesRoleStorageClassInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesRoleStorageClassInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type KubernetesRoleStorageClassInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (KubernetesRoleStorageClassInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesRoleStorageClassInfoResponse)(nil)).Elem()
}

func (o KubernetesRoleStorageClassInfoResponseArrayOutput) ToKubernetesRoleStorageClassInfoResponseArrayOutput() KubernetesRoleStorageClassInfoResponseArrayOutput {
	return o
}

func (o KubernetesRoleStorageClassInfoResponseArrayOutput) ToKubernetesRoleStorageClassInfoResponseArrayOutputWithContext(ctx context.Context) KubernetesRoleStorageClassInfoResponseArrayOutput {
	return o
}

func (o KubernetesRoleStorageClassInfoResponseArrayOutput) Index(i pulumi.IntInput) KubernetesRoleStorageClassInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KubernetesRoleStorageClassInfoResponse {
		return vs[0].([]KubernetesRoleStorageClassInfoResponse)[vs[1].(int)]
	}).(KubernetesRoleStorageClassInfoResponseOutput)
}

type KubernetesRoleStorageResponse struct {
	Endpoints      []MountPointMapResponse                  `pulumi:"endpoints"`
	StorageClasses []KubernetesRoleStorageClassInfoResponse `pulumi:"storageClasses"`
}





type KubernetesRoleStorageResponseInput interface {
	pulumi.Input

	ToKubernetesRoleStorageResponseOutput() KubernetesRoleStorageResponseOutput
	ToKubernetesRoleStorageResponseOutputWithContext(context.Context) KubernetesRoleStorageResponseOutput
}

type KubernetesRoleStorageResponseArgs struct {
	Endpoints      MountPointMapResponseArrayInput                  `pulumi:"endpoints"`
	StorageClasses KubernetesRoleStorageClassInfoResponseArrayInput `pulumi:"storageClasses"`
}

func (KubernetesRoleStorageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleStorageResponse)(nil)).Elem()
}

func (i KubernetesRoleStorageResponseArgs) ToKubernetesRoleStorageResponseOutput() KubernetesRoleStorageResponseOutput {
	return i.ToKubernetesRoleStorageResponseOutputWithContext(context.Background())
}

func (i KubernetesRoleStorageResponseArgs) ToKubernetesRoleStorageResponseOutputWithContext(ctx context.Context) KubernetesRoleStorageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleStorageResponseOutput)
}

func (i KubernetesRoleStorageResponseArgs) ToKubernetesRoleStorageResponsePtrOutput() KubernetesRoleStorageResponsePtrOutput {
	return i.ToKubernetesRoleStorageResponsePtrOutputWithContext(context.Background())
}

func (i KubernetesRoleStorageResponseArgs) ToKubernetesRoleStorageResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleStorageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleStorageResponseOutput).ToKubernetesRoleStorageResponsePtrOutputWithContext(ctx)
}









type KubernetesRoleStorageResponsePtrInput interface {
	pulumi.Input

	ToKubernetesRoleStorageResponsePtrOutput() KubernetesRoleStorageResponsePtrOutput
	ToKubernetesRoleStorageResponsePtrOutputWithContext(context.Context) KubernetesRoleStorageResponsePtrOutput
}

type kubernetesRoleStorageResponsePtrType KubernetesRoleStorageResponseArgs

func KubernetesRoleStorageResponsePtr(v *KubernetesRoleStorageResponseArgs) KubernetesRoleStorageResponsePtrInput {
	return (*kubernetesRoleStorageResponsePtrType)(v)
}

func (*kubernetesRoleStorageResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleStorageResponse)(nil)).Elem()
}

func (i *kubernetesRoleStorageResponsePtrType) ToKubernetesRoleStorageResponsePtrOutput() KubernetesRoleStorageResponsePtrOutput {
	return i.ToKubernetesRoleStorageResponsePtrOutputWithContext(context.Background())
}

func (i *kubernetesRoleStorageResponsePtrType) ToKubernetesRoleStorageResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleStorageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleStorageResponsePtrOutput)
}

type KubernetesRoleStorageResponseOutput struct{ *pulumi.OutputState }

func (KubernetesRoleStorageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRoleStorageResponse)(nil)).Elem()
}

func (o KubernetesRoleStorageResponseOutput) ToKubernetesRoleStorageResponseOutput() KubernetesRoleStorageResponseOutput {
	return o
}

func (o KubernetesRoleStorageResponseOutput) ToKubernetesRoleStorageResponseOutputWithContext(ctx context.Context) KubernetesRoleStorageResponseOutput {
	return o
}

func (o KubernetesRoleStorageResponseOutput) ToKubernetesRoleStorageResponsePtrOutput() KubernetesRoleStorageResponsePtrOutput {
	return o.ToKubernetesRoleStorageResponsePtrOutputWithContext(context.Background())
}

func (o KubernetesRoleStorageResponseOutput) ToKubernetesRoleStorageResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleStorageResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesRoleStorageResponse) *KubernetesRoleStorageResponse {
		return &v
	}).(KubernetesRoleStorageResponsePtrOutput)
}

func (o KubernetesRoleStorageResponseOutput) Endpoints() MountPointMapResponseArrayOutput {
	return o.ApplyT(func(v KubernetesRoleStorageResponse) []MountPointMapResponse { return v.Endpoints }).(MountPointMapResponseArrayOutput)
}

func (o KubernetesRoleStorageResponseOutput) StorageClasses() KubernetesRoleStorageClassInfoResponseArrayOutput {
	return o.ApplyT(func(v KubernetesRoleStorageResponse) []KubernetesRoleStorageClassInfoResponse {
		return v.StorageClasses
	}).(KubernetesRoleStorageClassInfoResponseArrayOutput)
}

type KubernetesRoleStorageResponsePtrOutput struct{ *pulumi.OutputState }

func (KubernetesRoleStorageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesRoleStorageResponse)(nil)).Elem()
}

func (o KubernetesRoleStorageResponsePtrOutput) ToKubernetesRoleStorageResponsePtrOutput() KubernetesRoleStorageResponsePtrOutput {
	return o
}

func (o KubernetesRoleStorageResponsePtrOutput) ToKubernetesRoleStorageResponsePtrOutputWithContext(ctx context.Context) KubernetesRoleStorageResponsePtrOutput {
	return o
}

func (o KubernetesRoleStorageResponsePtrOutput) Elem() KubernetesRoleStorageResponseOutput {
	return o.ApplyT(func(v *KubernetesRoleStorageResponse) KubernetesRoleStorageResponse {
		if v != nil {
			return *v
		}
		var ret KubernetesRoleStorageResponse
		return ret
	}).(KubernetesRoleStorageResponseOutput)
}

func (o KubernetesRoleStorageResponsePtrOutput) Endpoints() MountPointMapResponseArrayOutput {
	return o.ApplyT(func(v *KubernetesRoleStorageResponse) []MountPointMapResponse {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(MountPointMapResponseArrayOutput)
}

func (o KubernetesRoleStorageResponsePtrOutput) StorageClasses() KubernetesRoleStorageClassInfoResponseArrayOutput {
	return o.ApplyT(func(v *KubernetesRoleStorageResponse) []KubernetesRoleStorageClassInfoResponse {
		if v == nil {
			return nil
		}
		return v.StorageClasses
	}).(KubernetesRoleStorageClassInfoResponseArrayOutput)
}

type LoadBalancerConfigResponse struct {
	Type    string `pulumi:"type"`
	Version string `pulumi:"version"`
}





type LoadBalancerConfigResponseInput interface {
	pulumi.Input

	ToLoadBalancerConfigResponseOutput() LoadBalancerConfigResponseOutput
	ToLoadBalancerConfigResponseOutputWithContext(context.Context) LoadBalancerConfigResponseOutput
}

type LoadBalancerConfigResponseArgs struct {
	Type    pulumi.StringInput `pulumi:"type"`
	Version pulumi.StringInput `pulumi:"version"`
}

func (LoadBalancerConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerConfigResponse)(nil)).Elem()
}

func (i LoadBalancerConfigResponseArgs) ToLoadBalancerConfigResponseOutput() LoadBalancerConfigResponseOutput {
	return i.ToLoadBalancerConfigResponseOutputWithContext(context.Background())
}

func (i LoadBalancerConfigResponseArgs) ToLoadBalancerConfigResponseOutputWithContext(ctx context.Context) LoadBalancerConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerConfigResponseOutput)
}

func (i LoadBalancerConfigResponseArgs) ToLoadBalancerConfigResponsePtrOutput() LoadBalancerConfigResponsePtrOutput {
	return i.ToLoadBalancerConfigResponsePtrOutputWithContext(context.Background())
}

func (i LoadBalancerConfigResponseArgs) ToLoadBalancerConfigResponsePtrOutputWithContext(ctx context.Context) LoadBalancerConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerConfigResponseOutput).ToLoadBalancerConfigResponsePtrOutputWithContext(ctx)
}









type LoadBalancerConfigResponsePtrInput interface {
	pulumi.Input

	ToLoadBalancerConfigResponsePtrOutput() LoadBalancerConfigResponsePtrOutput
	ToLoadBalancerConfigResponsePtrOutputWithContext(context.Context) LoadBalancerConfigResponsePtrOutput
}

type loadBalancerConfigResponsePtrType LoadBalancerConfigResponseArgs

func LoadBalancerConfigResponsePtr(v *LoadBalancerConfigResponseArgs) LoadBalancerConfigResponsePtrInput {
	return (*loadBalancerConfigResponsePtrType)(v)
}

func (*loadBalancerConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerConfigResponse)(nil)).Elem()
}

func (i *loadBalancerConfigResponsePtrType) ToLoadBalancerConfigResponsePtrOutput() LoadBalancerConfigResponsePtrOutput {
	return i.ToLoadBalancerConfigResponsePtrOutputWithContext(context.Background())
}

func (i *loadBalancerConfigResponsePtrType) ToLoadBalancerConfigResponsePtrOutputWithContext(ctx context.Context) LoadBalancerConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerConfigResponsePtrOutput)
}

type LoadBalancerConfigResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerConfigResponse)(nil)).Elem()
}

func (o LoadBalancerConfigResponseOutput) ToLoadBalancerConfigResponseOutput() LoadBalancerConfigResponseOutput {
	return o
}

func (o LoadBalancerConfigResponseOutput) ToLoadBalancerConfigResponseOutputWithContext(ctx context.Context) LoadBalancerConfigResponseOutput {
	return o
}

func (o LoadBalancerConfigResponseOutput) ToLoadBalancerConfigResponsePtrOutput() LoadBalancerConfigResponsePtrOutput {
	return o.ToLoadBalancerConfigResponsePtrOutputWithContext(context.Background())
}

func (o LoadBalancerConfigResponseOutput) ToLoadBalancerConfigResponsePtrOutputWithContext(ctx context.Context) LoadBalancerConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancerConfigResponse) *LoadBalancerConfigResponse {
		return &v
	}).(LoadBalancerConfigResponsePtrOutput)
}

func (o LoadBalancerConfigResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerConfigResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o LoadBalancerConfigResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerConfigResponse) string { return v.Version }).(pulumi.StringOutput)
}

type LoadBalancerConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerConfigResponse)(nil)).Elem()
}

func (o LoadBalancerConfigResponsePtrOutput) ToLoadBalancerConfigResponsePtrOutput() LoadBalancerConfigResponsePtrOutput {
	return o
}

func (o LoadBalancerConfigResponsePtrOutput) ToLoadBalancerConfigResponsePtrOutputWithContext(ctx context.Context) LoadBalancerConfigResponsePtrOutput {
	return o
}

func (o LoadBalancerConfigResponsePtrOutput) Elem() LoadBalancerConfigResponseOutput {
	return o.ApplyT(func(v *LoadBalancerConfigResponse) LoadBalancerConfigResponse {
		if v != nil {
			return *v
		}
		var ret LoadBalancerConfigResponse
		return ret
	}).(LoadBalancerConfigResponseOutput)
}

func (o LoadBalancerConfigResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o LoadBalancerConfigResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type MetricConfiguration struct {
	CounterSets     []MetricCounterSet `pulumi:"counterSets"`
	MdmAccount      *string            `pulumi:"mdmAccount"`
	MetricNameSpace *string            `pulumi:"metricNameSpace"`
	ResourceId      string             `pulumi:"resourceId"`
}





type MetricConfigurationInput interface {
	pulumi.Input

	ToMetricConfigurationOutput() MetricConfigurationOutput
	ToMetricConfigurationOutputWithContext(context.Context) MetricConfigurationOutput
}

type MetricConfigurationArgs struct {
	CounterSets     MetricCounterSetArrayInput `pulumi:"counterSets"`
	MdmAccount      pulumi.StringPtrInput      `pulumi:"mdmAccount"`
	MetricNameSpace pulumi.StringPtrInput      `pulumi:"metricNameSpace"`
	ResourceId      pulumi.StringInput         `pulumi:"resourceId"`
}

func (MetricConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricConfiguration)(nil)).Elem()
}

func (i MetricConfigurationArgs) ToMetricConfigurationOutput() MetricConfigurationOutput {
	return i.ToMetricConfigurationOutputWithContext(context.Background())
}

func (i MetricConfigurationArgs) ToMetricConfigurationOutputWithContext(ctx context.Context) MetricConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricConfigurationOutput)
}





type MetricConfigurationArrayInput interface {
	pulumi.Input

	ToMetricConfigurationArrayOutput() MetricConfigurationArrayOutput
	ToMetricConfigurationArrayOutputWithContext(context.Context) MetricConfigurationArrayOutput
}

type MetricConfigurationArray []MetricConfigurationInput

func (MetricConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricConfiguration)(nil)).Elem()
}

func (i MetricConfigurationArray) ToMetricConfigurationArrayOutput() MetricConfigurationArrayOutput {
	return i.ToMetricConfigurationArrayOutputWithContext(context.Background())
}

func (i MetricConfigurationArray) ToMetricConfigurationArrayOutputWithContext(ctx context.Context) MetricConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricConfigurationArrayOutput)
}

type MetricConfigurationOutput struct{ *pulumi.OutputState }

func (MetricConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricConfiguration)(nil)).Elem()
}

func (o MetricConfigurationOutput) ToMetricConfigurationOutput() MetricConfigurationOutput {
	return o
}

func (o MetricConfigurationOutput) ToMetricConfigurationOutputWithContext(ctx context.Context) MetricConfigurationOutput {
	return o
}

func (o MetricConfigurationOutput) CounterSets() MetricCounterSetArrayOutput {
	return o.ApplyT(func(v MetricConfiguration) []MetricCounterSet { return v.CounterSets }).(MetricCounterSetArrayOutput)
}

func (o MetricConfigurationOutput) MdmAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricConfiguration) *string { return v.MdmAccount }).(pulumi.StringPtrOutput)
}

func (o MetricConfigurationOutput) MetricNameSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricConfiguration) *string { return v.MetricNameSpace }).(pulumi.StringPtrOutput)
}

func (o MetricConfigurationOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v MetricConfiguration) string { return v.ResourceId }).(pulumi.StringOutput)
}

type MetricConfigurationArrayOutput struct{ *pulumi.OutputState }

func (MetricConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricConfiguration)(nil)).Elem()
}

func (o MetricConfigurationArrayOutput) ToMetricConfigurationArrayOutput() MetricConfigurationArrayOutput {
	return o
}

func (o MetricConfigurationArrayOutput) ToMetricConfigurationArrayOutputWithContext(ctx context.Context) MetricConfigurationArrayOutput {
	return o
}

func (o MetricConfigurationArrayOutput) Index(i pulumi.IntInput) MetricConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricConfiguration {
		return vs[0].([]MetricConfiguration)[vs[1].(int)]
	}).(MetricConfigurationOutput)
}

type MetricConfigurationResponse struct {
	CounterSets     []MetricCounterSetResponse `pulumi:"counterSets"`
	MdmAccount      *string                    `pulumi:"mdmAccount"`
	MetricNameSpace *string                    `pulumi:"metricNameSpace"`
	ResourceId      string                     `pulumi:"resourceId"`
}





type MetricConfigurationResponseInput interface {
	pulumi.Input

	ToMetricConfigurationResponseOutput() MetricConfigurationResponseOutput
	ToMetricConfigurationResponseOutputWithContext(context.Context) MetricConfigurationResponseOutput
}

type MetricConfigurationResponseArgs struct {
	CounterSets     MetricCounterSetResponseArrayInput `pulumi:"counterSets"`
	MdmAccount      pulumi.StringPtrInput              `pulumi:"mdmAccount"`
	MetricNameSpace pulumi.StringPtrInput              `pulumi:"metricNameSpace"`
	ResourceId      pulumi.StringInput                 `pulumi:"resourceId"`
}

func (MetricConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricConfigurationResponse)(nil)).Elem()
}

func (i MetricConfigurationResponseArgs) ToMetricConfigurationResponseOutput() MetricConfigurationResponseOutput {
	return i.ToMetricConfigurationResponseOutputWithContext(context.Background())
}

func (i MetricConfigurationResponseArgs) ToMetricConfigurationResponseOutputWithContext(ctx context.Context) MetricConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricConfigurationResponseOutput)
}





type MetricConfigurationResponseArrayInput interface {
	pulumi.Input

	ToMetricConfigurationResponseArrayOutput() MetricConfigurationResponseArrayOutput
	ToMetricConfigurationResponseArrayOutputWithContext(context.Context) MetricConfigurationResponseArrayOutput
}

type MetricConfigurationResponseArray []MetricConfigurationResponseInput

func (MetricConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricConfigurationResponse)(nil)).Elem()
}

func (i MetricConfigurationResponseArray) ToMetricConfigurationResponseArrayOutput() MetricConfigurationResponseArrayOutput {
	return i.ToMetricConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i MetricConfigurationResponseArray) ToMetricConfigurationResponseArrayOutputWithContext(ctx context.Context) MetricConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricConfigurationResponseArrayOutput)
}

type MetricConfigurationResponseOutput struct{ *pulumi.OutputState }

func (MetricConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricConfigurationResponse)(nil)).Elem()
}

func (o MetricConfigurationResponseOutput) ToMetricConfigurationResponseOutput() MetricConfigurationResponseOutput {
	return o
}

func (o MetricConfigurationResponseOutput) ToMetricConfigurationResponseOutputWithContext(ctx context.Context) MetricConfigurationResponseOutput {
	return o
}

func (o MetricConfigurationResponseOutput) CounterSets() MetricCounterSetResponseArrayOutput {
	return o.ApplyT(func(v MetricConfigurationResponse) []MetricCounterSetResponse { return v.CounterSets }).(MetricCounterSetResponseArrayOutput)
}

func (o MetricConfigurationResponseOutput) MdmAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricConfigurationResponse) *string { return v.MdmAccount }).(pulumi.StringPtrOutput)
}

func (o MetricConfigurationResponseOutput) MetricNameSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricConfigurationResponse) *string { return v.MetricNameSpace }).(pulumi.StringPtrOutput)
}

func (o MetricConfigurationResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v MetricConfigurationResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

type MetricConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (MetricConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricConfigurationResponse)(nil)).Elem()
}

func (o MetricConfigurationResponseArrayOutput) ToMetricConfigurationResponseArrayOutput() MetricConfigurationResponseArrayOutput {
	return o
}

func (o MetricConfigurationResponseArrayOutput) ToMetricConfigurationResponseArrayOutputWithContext(ctx context.Context) MetricConfigurationResponseArrayOutput {
	return o
}

func (o MetricConfigurationResponseArrayOutput) Index(i pulumi.IntInput) MetricConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricConfigurationResponse {
		return vs[0].([]MetricConfigurationResponse)[vs[1].(int)]
	}).(MetricConfigurationResponseOutput)
}

type MetricCounter struct {
	AdditionalDimensions []MetricDimension `pulumi:"additionalDimensions"`
	DimensionFilter      []MetricDimension `pulumi:"dimensionFilter"`
	Instance             *string           `pulumi:"instance"`
	Name                 string            `pulumi:"name"`
}





type MetricCounterInput interface {
	pulumi.Input

	ToMetricCounterOutput() MetricCounterOutput
	ToMetricCounterOutputWithContext(context.Context) MetricCounterOutput
}

type MetricCounterArgs struct {
	AdditionalDimensions MetricDimensionArrayInput `pulumi:"additionalDimensions"`
	DimensionFilter      MetricDimensionArrayInput `pulumi:"dimensionFilter"`
	Instance             pulumi.StringPtrInput     `pulumi:"instance"`
	Name                 pulumi.StringInput        `pulumi:"name"`
}

func (MetricCounterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricCounter)(nil)).Elem()
}

func (i MetricCounterArgs) ToMetricCounterOutput() MetricCounterOutput {
	return i.ToMetricCounterOutputWithContext(context.Background())
}

func (i MetricCounterArgs) ToMetricCounterOutputWithContext(ctx context.Context) MetricCounterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricCounterOutput)
}





type MetricCounterArrayInput interface {
	pulumi.Input

	ToMetricCounterArrayOutput() MetricCounterArrayOutput
	ToMetricCounterArrayOutputWithContext(context.Context) MetricCounterArrayOutput
}

type MetricCounterArray []MetricCounterInput

func (MetricCounterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricCounter)(nil)).Elem()
}

func (i MetricCounterArray) ToMetricCounterArrayOutput() MetricCounterArrayOutput {
	return i.ToMetricCounterArrayOutputWithContext(context.Background())
}

func (i MetricCounterArray) ToMetricCounterArrayOutputWithContext(ctx context.Context) MetricCounterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricCounterArrayOutput)
}

type MetricCounterOutput struct{ *pulumi.OutputState }

func (MetricCounterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricCounter)(nil)).Elem()
}

func (o MetricCounterOutput) ToMetricCounterOutput() MetricCounterOutput {
	return o
}

func (o MetricCounterOutput) ToMetricCounterOutputWithContext(ctx context.Context) MetricCounterOutput {
	return o
}

func (o MetricCounterOutput) AdditionalDimensions() MetricDimensionArrayOutput {
	return o.ApplyT(func(v MetricCounter) []MetricDimension { return v.AdditionalDimensions }).(MetricDimensionArrayOutput)
}

func (o MetricCounterOutput) DimensionFilter() MetricDimensionArrayOutput {
	return o.ApplyT(func(v MetricCounter) []MetricDimension { return v.DimensionFilter }).(MetricDimensionArrayOutput)
}

func (o MetricCounterOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricCounter) *string { return v.Instance }).(pulumi.StringPtrOutput)
}

func (o MetricCounterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MetricCounter) string { return v.Name }).(pulumi.StringOutput)
}

type MetricCounterArrayOutput struct{ *pulumi.OutputState }

func (MetricCounterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricCounter)(nil)).Elem()
}

func (o MetricCounterArrayOutput) ToMetricCounterArrayOutput() MetricCounterArrayOutput {
	return o
}

func (o MetricCounterArrayOutput) ToMetricCounterArrayOutputWithContext(ctx context.Context) MetricCounterArrayOutput {
	return o
}

func (o MetricCounterArrayOutput) Index(i pulumi.IntInput) MetricCounterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricCounter {
		return vs[0].([]MetricCounter)[vs[1].(int)]
	}).(MetricCounterOutput)
}

type MetricCounterResponse struct {
	AdditionalDimensions []MetricDimensionResponse `pulumi:"additionalDimensions"`
	DimensionFilter      []MetricDimensionResponse `pulumi:"dimensionFilter"`
	Instance             *string                   `pulumi:"instance"`
	Name                 string                    `pulumi:"name"`
}





type MetricCounterResponseInput interface {
	pulumi.Input

	ToMetricCounterResponseOutput() MetricCounterResponseOutput
	ToMetricCounterResponseOutputWithContext(context.Context) MetricCounterResponseOutput
}

type MetricCounterResponseArgs struct {
	AdditionalDimensions MetricDimensionResponseArrayInput `pulumi:"additionalDimensions"`
	DimensionFilter      MetricDimensionResponseArrayInput `pulumi:"dimensionFilter"`
	Instance             pulumi.StringPtrInput             `pulumi:"instance"`
	Name                 pulumi.StringInput                `pulumi:"name"`
}

func (MetricCounterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricCounterResponse)(nil)).Elem()
}

func (i MetricCounterResponseArgs) ToMetricCounterResponseOutput() MetricCounterResponseOutput {
	return i.ToMetricCounterResponseOutputWithContext(context.Background())
}

func (i MetricCounterResponseArgs) ToMetricCounterResponseOutputWithContext(ctx context.Context) MetricCounterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricCounterResponseOutput)
}





type MetricCounterResponseArrayInput interface {
	pulumi.Input

	ToMetricCounterResponseArrayOutput() MetricCounterResponseArrayOutput
	ToMetricCounterResponseArrayOutputWithContext(context.Context) MetricCounterResponseArrayOutput
}

type MetricCounterResponseArray []MetricCounterResponseInput

func (MetricCounterResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricCounterResponse)(nil)).Elem()
}

func (i MetricCounterResponseArray) ToMetricCounterResponseArrayOutput() MetricCounterResponseArrayOutput {
	return i.ToMetricCounterResponseArrayOutputWithContext(context.Background())
}

func (i MetricCounterResponseArray) ToMetricCounterResponseArrayOutputWithContext(ctx context.Context) MetricCounterResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricCounterResponseArrayOutput)
}

type MetricCounterResponseOutput struct{ *pulumi.OutputState }

func (MetricCounterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricCounterResponse)(nil)).Elem()
}

func (o MetricCounterResponseOutput) ToMetricCounterResponseOutput() MetricCounterResponseOutput {
	return o
}

func (o MetricCounterResponseOutput) ToMetricCounterResponseOutputWithContext(ctx context.Context) MetricCounterResponseOutput {
	return o
}

func (o MetricCounterResponseOutput) AdditionalDimensions() MetricDimensionResponseArrayOutput {
	return o.ApplyT(func(v MetricCounterResponse) []MetricDimensionResponse { return v.AdditionalDimensions }).(MetricDimensionResponseArrayOutput)
}

func (o MetricCounterResponseOutput) DimensionFilter() MetricDimensionResponseArrayOutput {
	return o.ApplyT(func(v MetricCounterResponse) []MetricDimensionResponse { return v.DimensionFilter }).(MetricDimensionResponseArrayOutput)
}

func (o MetricCounterResponseOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricCounterResponse) *string { return v.Instance }).(pulumi.StringPtrOutput)
}

func (o MetricCounterResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MetricCounterResponse) string { return v.Name }).(pulumi.StringOutput)
}

type MetricCounterResponseArrayOutput struct{ *pulumi.OutputState }

func (MetricCounterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricCounterResponse)(nil)).Elem()
}

func (o MetricCounterResponseArrayOutput) ToMetricCounterResponseArrayOutput() MetricCounterResponseArrayOutput {
	return o
}

func (o MetricCounterResponseArrayOutput) ToMetricCounterResponseArrayOutputWithContext(ctx context.Context) MetricCounterResponseArrayOutput {
	return o
}

func (o MetricCounterResponseArrayOutput) Index(i pulumi.IntInput) MetricCounterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricCounterResponse {
		return vs[0].([]MetricCounterResponse)[vs[1].(int)]
	}).(MetricCounterResponseOutput)
}

type MetricCounterSet struct {
	Counters []MetricCounter `pulumi:"counters"`
}





type MetricCounterSetInput interface {
	pulumi.Input

	ToMetricCounterSetOutput() MetricCounterSetOutput
	ToMetricCounterSetOutputWithContext(context.Context) MetricCounterSetOutput
}

type MetricCounterSetArgs struct {
	Counters MetricCounterArrayInput `pulumi:"counters"`
}

func (MetricCounterSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricCounterSet)(nil)).Elem()
}

func (i MetricCounterSetArgs) ToMetricCounterSetOutput() MetricCounterSetOutput {
	return i.ToMetricCounterSetOutputWithContext(context.Background())
}

func (i MetricCounterSetArgs) ToMetricCounterSetOutputWithContext(ctx context.Context) MetricCounterSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricCounterSetOutput)
}





type MetricCounterSetArrayInput interface {
	pulumi.Input

	ToMetricCounterSetArrayOutput() MetricCounterSetArrayOutput
	ToMetricCounterSetArrayOutputWithContext(context.Context) MetricCounterSetArrayOutput
}

type MetricCounterSetArray []MetricCounterSetInput

func (MetricCounterSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricCounterSet)(nil)).Elem()
}

func (i MetricCounterSetArray) ToMetricCounterSetArrayOutput() MetricCounterSetArrayOutput {
	return i.ToMetricCounterSetArrayOutputWithContext(context.Background())
}

func (i MetricCounterSetArray) ToMetricCounterSetArrayOutputWithContext(ctx context.Context) MetricCounterSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricCounterSetArrayOutput)
}

type MetricCounterSetOutput struct{ *pulumi.OutputState }

func (MetricCounterSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricCounterSet)(nil)).Elem()
}

func (o MetricCounterSetOutput) ToMetricCounterSetOutput() MetricCounterSetOutput {
	return o
}

func (o MetricCounterSetOutput) ToMetricCounterSetOutputWithContext(ctx context.Context) MetricCounterSetOutput {
	return o
}

func (o MetricCounterSetOutput) Counters() MetricCounterArrayOutput {
	return o.ApplyT(func(v MetricCounterSet) []MetricCounter { return v.Counters }).(MetricCounterArrayOutput)
}

type MetricCounterSetArrayOutput struct{ *pulumi.OutputState }

func (MetricCounterSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricCounterSet)(nil)).Elem()
}

func (o MetricCounterSetArrayOutput) ToMetricCounterSetArrayOutput() MetricCounterSetArrayOutput {
	return o
}

func (o MetricCounterSetArrayOutput) ToMetricCounterSetArrayOutputWithContext(ctx context.Context) MetricCounterSetArrayOutput {
	return o
}

func (o MetricCounterSetArrayOutput) Index(i pulumi.IntInput) MetricCounterSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricCounterSet {
		return vs[0].([]MetricCounterSet)[vs[1].(int)]
	}).(MetricCounterSetOutput)
}

type MetricCounterSetResponse struct {
	Counters []MetricCounterResponse `pulumi:"counters"`
}





type MetricCounterSetResponseInput interface {
	pulumi.Input

	ToMetricCounterSetResponseOutput() MetricCounterSetResponseOutput
	ToMetricCounterSetResponseOutputWithContext(context.Context) MetricCounterSetResponseOutput
}

type MetricCounterSetResponseArgs struct {
	Counters MetricCounterResponseArrayInput `pulumi:"counters"`
}

func (MetricCounterSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricCounterSetResponse)(nil)).Elem()
}

func (i MetricCounterSetResponseArgs) ToMetricCounterSetResponseOutput() MetricCounterSetResponseOutput {
	return i.ToMetricCounterSetResponseOutputWithContext(context.Background())
}

func (i MetricCounterSetResponseArgs) ToMetricCounterSetResponseOutputWithContext(ctx context.Context) MetricCounterSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricCounterSetResponseOutput)
}





type MetricCounterSetResponseArrayInput interface {
	pulumi.Input

	ToMetricCounterSetResponseArrayOutput() MetricCounterSetResponseArrayOutput
	ToMetricCounterSetResponseArrayOutputWithContext(context.Context) MetricCounterSetResponseArrayOutput
}

type MetricCounterSetResponseArray []MetricCounterSetResponseInput

func (MetricCounterSetResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricCounterSetResponse)(nil)).Elem()
}

func (i MetricCounterSetResponseArray) ToMetricCounterSetResponseArrayOutput() MetricCounterSetResponseArrayOutput {
	return i.ToMetricCounterSetResponseArrayOutputWithContext(context.Background())
}

func (i MetricCounterSetResponseArray) ToMetricCounterSetResponseArrayOutputWithContext(ctx context.Context) MetricCounterSetResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricCounterSetResponseArrayOutput)
}

type MetricCounterSetResponseOutput struct{ *pulumi.OutputState }

func (MetricCounterSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricCounterSetResponse)(nil)).Elem()
}

func (o MetricCounterSetResponseOutput) ToMetricCounterSetResponseOutput() MetricCounterSetResponseOutput {
	return o
}

func (o MetricCounterSetResponseOutput) ToMetricCounterSetResponseOutputWithContext(ctx context.Context) MetricCounterSetResponseOutput {
	return o
}

func (o MetricCounterSetResponseOutput) Counters() MetricCounterResponseArrayOutput {
	return o.ApplyT(func(v MetricCounterSetResponse) []MetricCounterResponse { return v.Counters }).(MetricCounterResponseArrayOutput)
}

type MetricCounterSetResponseArrayOutput struct{ *pulumi.OutputState }

func (MetricCounterSetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricCounterSetResponse)(nil)).Elem()
}

func (o MetricCounterSetResponseArrayOutput) ToMetricCounterSetResponseArrayOutput() MetricCounterSetResponseArrayOutput {
	return o
}

func (o MetricCounterSetResponseArrayOutput) ToMetricCounterSetResponseArrayOutputWithContext(ctx context.Context) MetricCounterSetResponseArrayOutput {
	return o
}

func (o MetricCounterSetResponseArrayOutput) Index(i pulumi.IntInput) MetricCounterSetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricCounterSetResponse {
		return vs[0].([]MetricCounterSetResponse)[vs[1].(int)]
	}).(MetricCounterSetResponseOutput)
}

type MetricDimension struct {
	SourceName string `pulumi:"sourceName"`
	SourceType string `pulumi:"sourceType"`
}





type MetricDimensionInput interface {
	pulumi.Input

	ToMetricDimensionOutput() MetricDimensionOutput
	ToMetricDimensionOutputWithContext(context.Context) MetricDimensionOutput
}

type MetricDimensionArgs struct {
	SourceName pulumi.StringInput `pulumi:"sourceName"`
	SourceType pulumi.StringInput `pulumi:"sourceType"`
}

func (MetricDimensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricDimension)(nil)).Elem()
}

func (i MetricDimensionArgs) ToMetricDimensionOutput() MetricDimensionOutput {
	return i.ToMetricDimensionOutputWithContext(context.Background())
}

func (i MetricDimensionArgs) ToMetricDimensionOutputWithContext(ctx context.Context) MetricDimensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricDimensionOutput)
}





type MetricDimensionArrayInput interface {
	pulumi.Input

	ToMetricDimensionArrayOutput() MetricDimensionArrayOutput
	ToMetricDimensionArrayOutputWithContext(context.Context) MetricDimensionArrayOutput
}

type MetricDimensionArray []MetricDimensionInput

func (MetricDimensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricDimension)(nil)).Elem()
}

func (i MetricDimensionArray) ToMetricDimensionArrayOutput() MetricDimensionArrayOutput {
	return i.ToMetricDimensionArrayOutputWithContext(context.Background())
}

func (i MetricDimensionArray) ToMetricDimensionArrayOutputWithContext(ctx context.Context) MetricDimensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricDimensionArrayOutput)
}

type MetricDimensionOutput struct{ *pulumi.OutputState }

func (MetricDimensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricDimension)(nil)).Elem()
}

func (o MetricDimensionOutput) ToMetricDimensionOutput() MetricDimensionOutput {
	return o
}

func (o MetricDimensionOutput) ToMetricDimensionOutputWithContext(ctx context.Context) MetricDimensionOutput {
	return o
}

func (o MetricDimensionOutput) SourceName() pulumi.StringOutput {
	return o.ApplyT(func(v MetricDimension) string { return v.SourceName }).(pulumi.StringOutput)
}

func (o MetricDimensionOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v MetricDimension) string { return v.SourceType }).(pulumi.StringOutput)
}

type MetricDimensionArrayOutput struct{ *pulumi.OutputState }

func (MetricDimensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricDimension)(nil)).Elem()
}

func (o MetricDimensionArrayOutput) ToMetricDimensionArrayOutput() MetricDimensionArrayOutput {
	return o
}

func (o MetricDimensionArrayOutput) ToMetricDimensionArrayOutputWithContext(ctx context.Context) MetricDimensionArrayOutput {
	return o
}

func (o MetricDimensionArrayOutput) Index(i pulumi.IntInput) MetricDimensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricDimension {
		return vs[0].([]MetricDimension)[vs[1].(int)]
	}).(MetricDimensionOutput)
}

type MetricDimensionResponse struct {
	SourceName string `pulumi:"sourceName"`
	SourceType string `pulumi:"sourceType"`
}





type MetricDimensionResponseInput interface {
	pulumi.Input

	ToMetricDimensionResponseOutput() MetricDimensionResponseOutput
	ToMetricDimensionResponseOutputWithContext(context.Context) MetricDimensionResponseOutput
}

type MetricDimensionResponseArgs struct {
	SourceName pulumi.StringInput `pulumi:"sourceName"`
	SourceType pulumi.StringInput `pulumi:"sourceType"`
}

func (MetricDimensionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricDimensionResponse)(nil)).Elem()
}

func (i MetricDimensionResponseArgs) ToMetricDimensionResponseOutput() MetricDimensionResponseOutput {
	return i.ToMetricDimensionResponseOutputWithContext(context.Background())
}

func (i MetricDimensionResponseArgs) ToMetricDimensionResponseOutputWithContext(ctx context.Context) MetricDimensionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricDimensionResponseOutput)
}





type MetricDimensionResponseArrayInput interface {
	pulumi.Input

	ToMetricDimensionResponseArrayOutput() MetricDimensionResponseArrayOutput
	ToMetricDimensionResponseArrayOutputWithContext(context.Context) MetricDimensionResponseArrayOutput
}

type MetricDimensionResponseArray []MetricDimensionResponseInput

func (MetricDimensionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricDimensionResponse)(nil)).Elem()
}

func (i MetricDimensionResponseArray) ToMetricDimensionResponseArrayOutput() MetricDimensionResponseArrayOutput {
	return i.ToMetricDimensionResponseArrayOutputWithContext(context.Background())
}

func (i MetricDimensionResponseArray) ToMetricDimensionResponseArrayOutputWithContext(ctx context.Context) MetricDimensionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricDimensionResponseArrayOutput)
}

type MetricDimensionResponseOutput struct{ *pulumi.OutputState }

func (MetricDimensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricDimensionResponse)(nil)).Elem()
}

func (o MetricDimensionResponseOutput) ToMetricDimensionResponseOutput() MetricDimensionResponseOutput {
	return o
}

func (o MetricDimensionResponseOutput) ToMetricDimensionResponseOutputWithContext(ctx context.Context) MetricDimensionResponseOutput {
	return o
}

func (o MetricDimensionResponseOutput) SourceName() pulumi.StringOutput {
	return o.ApplyT(func(v MetricDimensionResponse) string { return v.SourceName }).(pulumi.StringOutput)
}

func (o MetricDimensionResponseOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v MetricDimensionResponse) string { return v.SourceType }).(pulumi.StringOutput)
}

type MetricDimensionResponseArrayOutput struct{ *pulumi.OutputState }

func (MetricDimensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricDimensionResponse)(nil)).Elem()
}

func (o MetricDimensionResponseArrayOutput) ToMetricDimensionResponseArrayOutput() MetricDimensionResponseArrayOutput {
	return o
}

func (o MetricDimensionResponseArrayOutput) ToMetricDimensionResponseArrayOutputWithContext(ctx context.Context) MetricDimensionResponseArrayOutput {
	return o
}

func (o MetricDimensionResponseArrayOutput) Index(i pulumi.IntInput) MetricDimensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricDimensionResponse {
		return vs[0].([]MetricDimensionResponse)[vs[1].(int)]
	}).(MetricDimensionResponseOutput)
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
	MountType  string `pulumi:"mountType"`
	RoleId     string `pulumi:"roleId"`
	RoleType   string `pulumi:"roleType"`
	ShareId    string `pulumi:"shareId"`
}





type MountPointMapResponseInput interface {
	pulumi.Input

	ToMountPointMapResponseOutput() MountPointMapResponseOutput
	ToMountPointMapResponseOutputWithContext(context.Context) MountPointMapResponseOutput
}

type MountPointMapResponseArgs struct {
	MountPoint pulumi.StringInput `pulumi:"mountPoint"`
	MountType  pulumi.StringInput `pulumi:"mountType"`
	RoleId     pulumi.StringInput `pulumi:"roleId"`
	RoleType   pulumi.StringInput `pulumi:"roleType"`
	ShareId    pulumi.StringInput `pulumi:"shareId"`
}

func (MountPointMapResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MountPointMapResponse)(nil)).Elem()
}

func (i MountPointMapResponseArgs) ToMountPointMapResponseOutput() MountPointMapResponseOutput {
	return i.ToMountPointMapResponseOutputWithContext(context.Background())
}

func (i MountPointMapResponseArgs) ToMountPointMapResponseOutputWithContext(ctx context.Context) MountPointMapResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointMapResponseOutput)
}





type MountPointMapResponseArrayInput interface {
	pulumi.Input

	ToMountPointMapResponseArrayOutput() MountPointMapResponseArrayOutput
	ToMountPointMapResponseArrayOutputWithContext(context.Context) MountPointMapResponseArrayOutput
}

type MountPointMapResponseArray []MountPointMapResponseInput

func (MountPointMapResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MountPointMapResponse)(nil)).Elem()
}

func (i MountPointMapResponseArray) ToMountPointMapResponseArrayOutput() MountPointMapResponseArrayOutput {
	return i.ToMountPointMapResponseArrayOutputWithContext(context.Background())
}

func (i MountPointMapResponseArray) ToMountPointMapResponseArrayOutputWithContext(ctx context.Context) MountPointMapResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointMapResponseArrayOutput)
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

func (o MountPointMapResponseOutput) MountType() pulumi.StringOutput {
	return o.ApplyT(func(v MountPointMapResponse) string { return v.MountType }).(pulumi.StringOutput)
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

type NodeInfoResponse struct {
	IpConfiguration []KubernetesIPConfigurationResponse `pulumi:"ipConfiguration"`
	Name            string                              `pulumi:"name"`
	Type            string                              `pulumi:"type"`
}





type NodeInfoResponseInput interface {
	pulumi.Input

	ToNodeInfoResponseOutput() NodeInfoResponseOutput
	ToNodeInfoResponseOutputWithContext(context.Context) NodeInfoResponseOutput
}

type NodeInfoResponseArgs struct {
	IpConfiguration KubernetesIPConfigurationResponseArrayInput `pulumi:"ipConfiguration"`
	Name            pulumi.StringInput                          `pulumi:"name"`
	Type            pulumi.StringInput                          `pulumi:"type"`
}

func (NodeInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeInfoResponse)(nil)).Elem()
}

func (i NodeInfoResponseArgs) ToNodeInfoResponseOutput() NodeInfoResponseOutput {
	return i.ToNodeInfoResponseOutputWithContext(context.Background())
}

func (i NodeInfoResponseArgs) ToNodeInfoResponseOutputWithContext(ctx context.Context) NodeInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeInfoResponseOutput)
}





type NodeInfoResponseArrayInput interface {
	pulumi.Input

	ToNodeInfoResponseArrayOutput() NodeInfoResponseArrayOutput
	ToNodeInfoResponseArrayOutputWithContext(context.Context) NodeInfoResponseArrayOutput
}

type NodeInfoResponseArray []NodeInfoResponseInput

func (NodeInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeInfoResponse)(nil)).Elem()
}

func (i NodeInfoResponseArray) ToNodeInfoResponseArrayOutput() NodeInfoResponseArrayOutput {
	return i.ToNodeInfoResponseArrayOutputWithContext(context.Background())
}

func (i NodeInfoResponseArray) ToNodeInfoResponseArrayOutputWithContext(ctx context.Context) NodeInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeInfoResponseArrayOutput)
}

type NodeInfoResponseOutput struct{ *pulumi.OutputState }

func (NodeInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeInfoResponse)(nil)).Elem()
}

func (o NodeInfoResponseOutput) ToNodeInfoResponseOutput() NodeInfoResponseOutput {
	return o
}

func (o NodeInfoResponseOutput) ToNodeInfoResponseOutputWithContext(ctx context.Context) NodeInfoResponseOutput {
	return o
}

func (o NodeInfoResponseOutput) IpConfiguration() KubernetesIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v NodeInfoResponse) []KubernetesIPConfigurationResponse { return v.IpConfiguration }).(KubernetesIPConfigurationResponseArrayOutput)
}

func (o NodeInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NodeInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o NodeInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v NodeInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type NodeInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (NodeInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeInfoResponse)(nil)).Elem()
}

func (o NodeInfoResponseArrayOutput) ToNodeInfoResponseArrayOutput() NodeInfoResponseArrayOutput {
	return o
}

func (o NodeInfoResponseArrayOutput) ToNodeInfoResponseArrayOutputWithContext(ctx context.Context) NodeInfoResponseArrayOutput {
	return o
}

func (o NodeInfoResponseArrayOutput) Index(i pulumi.IntInput) NodeInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeInfoResponse {
		return vs[0].([]NodeInfoResponse)[vs[1].(int)]
	}).(NodeInfoResponseOutput)
}

type OrderStatusResponse struct {
	AdditionalOrderDetails map[string]string    `pulumi:"additionalOrderDetails"`
	Comments               *string              `pulumi:"comments"`
	Status                 string               `pulumi:"status"`
	TrackingInformation    TrackingInfoResponse `pulumi:"trackingInformation"`
	UpdateDateTime         string               `pulumi:"updateDateTime"`
}





type OrderStatusResponseInput interface {
	pulumi.Input

	ToOrderStatusResponseOutput() OrderStatusResponseOutput
	ToOrderStatusResponseOutputWithContext(context.Context) OrderStatusResponseOutput
}

type OrderStatusResponseArgs struct {
	AdditionalOrderDetails pulumi.StringMapInput     `pulumi:"additionalOrderDetails"`
	Comments               pulumi.StringPtrInput     `pulumi:"comments"`
	Status                 pulumi.StringInput        `pulumi:"status"`
	TrackingInformation    TrackingInfoResponseInput `pulumi:"trackingInformation"`
	UpdateDateTime         pulumi.StringInput        `pulumi:"updateDateTime"`
}

func (OrderStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderStatusResponse)(nil)).Elem()
}

func (i OrderStatusResponseArgs) ToOrderStatusResponseOutput() OrderStatusResponseOutput {
	return i.ToOrderStatusResponseOutputWithContext(context.Background())
}

func (i OrderStatusResponseArgs) ToOrderStatusResponseOutputWithContext(ctx context.Context) OrderStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderStatusResponseOutput)
}

func (i OrderStatusResponseArgs) ToOrderStatusResponsePtrOutput() OrderStatusResponsePtrOutput {
	return i.ToOrderStatusResponsePtrOutputWithContext(context.Background())
}

func (i OrderStatusResponseArgs) ToOrderStatusResponsePtrOutputWithContext(ctx context.Context) OrderStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderStatusResponseOutput).ToOrderStatusResponsePtrOutputWithContext(ctx)
}









type OrderStatusResponsePtrInput interface {
	pulumi.Input

	ToOrderStatusResponsePtrOutput() OrderStatusResponsePtrOutput
	ToOrderStatusResponsePtrOutputWithContext(context.Context) OrderStatusResponsePtrOutput
}

type orderStatusResponsePtrType OrderStatusResponseArgs

func OrderStatusResponsePtr(v *OrderStatusResponseArgs) OrderStatusResponsePtrInput {
	return (*orderStatusResponsePtrType)(v)
}

func (*orderStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderStatusResponse)(nil)).Elem()
}

func (i *orderStatusResponsePtrType) ToOrderStatusResponsePtrOutput() OrderStatusResponsePtrOutput {
	return i.ToOrderStatusResponsePtrOutputWithContext(context.Background())
}

func (i *orderStatusResponsePtrType) ToOrderStatusResponsePtrOutputWithContext(ctx context.Context) OrderStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderStatusResponsePtrOutput)
}





type OrderStatusResponseArrayInput interface {
	pulumi.Input

	ToOrderStatusResponseArrayOutput() OrderStatusResponseArrayOutput
	ToOrderStatusResponseArrayOutputWithContext(context.Context) OrderStatusResponseArrayOutput
}

type OrderStatusResponseArray []OrderStatusResponseInput

func (OrderStatusResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrderStatusResponse)(nil)).Elem()
}

func (i OrderStatusResponseArray) ToOrderStatusResponseArrayOutput() OrderStatusResponseArrayOutput {
	return i.ToOrderStatusResponseArrayOutputWithContext(context.Background())
}

func (i OrderStatusResponseArray) ToOrderStatusResponseArrayOutputWithContext(ctx context.Context) OrderStatusResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderStatusResponseArrayOutput)
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

func (o OrderStatusResponseOutput) ToOrderStatusResponsePtrOutput() OrderStatusResponsePtrOutput {
	return o.ToOrderStatusResponsePtrOutputWithContext(context.Background())
}

func (o OrderStatusResponseOutput) ToOrderStatusResponsePtrOutputWithContext(ctx context.Context) OrderStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrderStatusResponse) *OrderStatusResponse {
		return &v
	}).(OrderStatusResponsePtrOutput)
}

func (o OrderStatusResponseOutput) AdditionalOrderDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v OrderStatusResponse) map[string]string { return v.AdditionalOrderDetails }).(pulumi.StringMapOutput)
}

func (o OrderStatusResponseOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrderStatusResponse) *string { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o OrderStatusResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v OrderStatusResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o OrderStatusResponseOutput) TrackingInformation() TrackingInfoResponseOutput {
	return o.ApplyT(func(v OrderStatusResponse) TrackingInfoResponse { return v.TrackingInformation }).(TrackingInfoResponseOutput)
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

func (o OrderStatusResponsePtrOutput) AdditionalOrderDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OrderStatusResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.AdditionalOrderDetails
	}).(pulumi.StringMapOutput)
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

func (o OrderStatusResponsePtrOutput) TrackingInformation() TrackingInfoResponsePtrOutput {
	return o.ApplyT(func(v *OrderStatusResponse) *TrackingInfoResponse {
		if v == nil {
			return nil
		}
		return &v.TrackingInformation
	}).(TrackingInfoResponsePtrOutput)
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

func (i PeriodicTimerSourceInfoArgs) ToPeriodicTimerSourceInfoPtrOutput() PeriodicTimerSourceInfoPtrOutput {
	return i.ToPeriodicTimerSourceInfoPtrOutputWithContext(context.Background())
}

func (i PeriodicTimerSourceInfoArgs) ToPeriodicTimerSourceInfoPtrOutputWithContext(ctx context.Context) PeriodicTimerSourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicTimerSourceInfoOutput).ToPeriodicTimerSourceInfoPtrOutputWithContext(ctx)
}









type PeriodicTimerSourceInfoPtrInput interface {
	pulumi.Input

	ToPeriodicTimerSourceInfoPtrOutput() PeriodicTimerSourceInfoPtrOutput
	ToPeriodicTimerSourceInfoPtrOutputWithContext(context.Context) PeriodicTimerSourceInfoPtrOutput
}

type periodicTimerSourceInfoPtrType PeriodicTimerSourceInfoArgs

func PeriodicTimerSourceInfoPtr(v *PeriodicTimerSourceInfoArgs) PeriodicTimerSourceInfoPtrInput {
	return (*periodicTimerSourceInfoPtrType)(v)
}

func (*periodicTimerSourceInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PeriodicTimerSourceInfo)(nil)).Elem()
}

func (i *periodicTimerSourceInfoPtrType) ToPeriodicTimerSourceInfoPtrOutput() PeriodicTimerSourceInfoPtrOutput {
	return i.ToPeriodicTimerSourceInfoPtrOutputWithContext(context.Background())
}

func (i *periodicTimerSourceInfoPtrType) ToPeriodicTimerSourceInfoPtrOutputWithContext(ctx context.Context) PeriodicTimerSourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicTimerSourceInfoPtrOutput)
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

func (o PeriodicTimerSourceInfoOutput) ToPeriodicTimerSourceInfoPtrOutput() PeriodicTimerSourceInfoPtrOutput {
	return o.ToPeriodicTimerSourceInfoPtrOutputWithContext(context.Background())
}

func (o PeriodicTimerSourceInfoOutput) ToPeriodicTimerSourceInfoPtrOutputWithContext(ctx context.Context) PeriodicTimerSourceInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PeriodicTimerSourceInfo) *PeriodicTimerSourceInfo {
		return &v
	}).(PeriodicTimerSourceInfoPtrOutput)
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

type PeriodicTimerSourceInfoPtrOutput struct{ *pulumi.OutputState }

func (PeriodicTimerSourceInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeriodicTimerSourceInfo)(nil)).Elem()
}

func (o PeriodicTimerSourceInfoPtrOutput) ToPeriodicTimerSourceInfoPtrOutput() PeriodicTimerSourceInfoPtrOutput {
	return o
}

func (o PeriodicTimerSourceInfoPtrOutput) ToPeriodicTimerSourceInfoPtrOutputWithContext(ctx context.Context) PeriodicTimerSourceInfoPtrOutput {
	return o
}

func (o PeriodicTimerSourceInfoPtrOutput) Elem() PeriodicTimerSourceInfoOutput {
	return o.ApplyT(func(v *PeriodicTimerSourceInfo) PeriodicTimerSourceInfo {
		if v != nil {
			return *v
		}
		var ret PeriodicTimerSourceInfo
		return ret
	}).(PeriodicTimerSourceInfoOutput)
}

func (o PeriodicTimerSourceInfoPtrOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeriodicTimerSourceInfo) *string {
		if v == nil {
			return nil
		}
		return &v.Schedule
	}).(pulumi.StringPtrOutput)
}

func (o PeriodicTimerSourceInfoPtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeriodicTimerSourceInfo) *string {
		if v == nil {
			return nil
		}
		return &v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o PeriodicTimerSourceInfoPtrOutput) Topic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeriodicTimerSourceInfo) *string {
		if v == nil {
			return nil
		}
		return v.Topic
	}).(pulumi.StringPtrOutput)
}

type PeriodicTimerSourceInfoResponse struct {
	Schedule  string  `pulumi:"schedule"`
	StartTime string  `pulumi:"startTime"`
	Topic     *string `pulumi:"topic"`
}





type PeriodicTimerSourceInfoResponseInput interface {
	pulumi.Input

	ToPeriodicTimerSourceInfoResponseOutput() PeriodicTimerSourceInfoResponseOutput
	ToPeriodicTimerSourceInfoResponseOutputWithContext(context.Context) PeriodicTimerSourceInfoResponseOutput
}

type PeriodicTimerSourceInfoResponseArgs struct {
	Schedule  pulumi.StringInput    `pulumi:"schedule"`
	StartTime pulumi.StringInput    `pulumi:"startTime"`
	Topic     pulumi.StringPtrInput `pulumi:"topic"`
}

func (PeriodicTimerSourceInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeriodicTimerSourceInfoResponse)(nil)).Elem()
}

func (i PeriodicTimerSourceInfoResponseArgs) ToPeriodicTimerSourceInfoResponseOutput() PeriodicTimerSourceInfoResponseOutput {
	return i.ToPeriodicTimerSourceInfoResponseOutputWithContext(context.Background())
}

func (i PeriodicTimerSourceInfoResponseArgs) ToPeriodicTimerSourceInfoResponseOutputWithContext(ctx context.Context) PeriodicTimerSourceInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicTimerSourceInfoResponseOutput)
}

func (i PeriodicTimerSourceInfoResponseArgs) ToPeriodicTimerSourceInfoResponsePtrOutput() PeriodicTimerSourceInfoResponsePtrOutput {
	return i.ToPeriodicTimerSourceInfoResponsePtrOutputWithContext(context.Background())
}

func (i PeriodicTimerSourceInfoResponseArgs) ToPeriodicTimerSourceInfoResponsePtrOutputWithContext(ctx context.Context) PeriodicTimerSourceInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicTimerSourceInfoResponseOutput).ToPeriodicTimerSourceInfoResponsePtrOutputWithContext(ctx)
}









type PeriodicTimerSourceInfoResponsePtrInput interface {
	pulumi.Input

	ToPeriodicTimerSourceInfoResponsePtrOutput() PeriodicTimerSourceInfoResponsePtrOutput
	ToPeriodicTimerSourceInfoResponsePtrOutputWithContext(context.Context) PeriodicTimerSourceInfoResponsePtrOutput
}

type periodicTimerSourceInfoResponsePtrType PeriodicTimerSourceInfoResponseArgs

func PeriodicTimerSourceInfoResponsePtr(v *PeriodicTimerSourceInfoResponseArgs) PeriodicTimerSourceInfoResponsePtrInput {
	return (*periodicTimerSourceInfoResponsePtrType)(v)
}

func (*periodicTimerSourceInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PeriodicTimerSourceInfoResponse)(nil)).Elem()
}

func (i *periodicTimerSourceInfoResponsePtrType) ToPeriodicTimerSourceInfoResponsePtrOutput() PeriodicTimerSourceInfoResponsePtrOutput {
	return i.ToPeriodicTimerSourceInfoResponsePtrOutputWithContext(context.Background())
}

func (i *periodicTimerSourceInfoResponsePtrType) ToPeriodicTimerSourceInfoResponsePtrOutputWithContext(ctx context.Context) PeriodicTimerSourceInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicTimerSourceInfoResponsePtrOutput)
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

func (o PeriodicTimerSourceInfoResponseOutput) ToPeriodicTimerSourceInfoResponsePtrOutput() PeriodicTimerSourceInfoResponsePtrOutput {
	return o.ToPeriodicTimerSourceInfoResponsePtrOutputWithContext(context.Background())
}

func (o PeriodicTimerSourceInfoResponseOutput) ToPeriodicTimerSourceInfoResponsePtrOutputWithContext(ctx context.Context) PeriodicTimerSourceInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PeriodicTimerSourceInfoResponse) *PeriodicTimerSourceInfoResponse {
		return &v
	}).(PeriodicTimerSourceInfoResponsePtrOutput)
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

type PeriodicTimerSourceInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (PeriodicTimerSourceInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeriodicTimerSourceInfoResponse)(nil)).Elem()
}

func (o PeriodicTimerSourceInfoResponsePtrOutput) ToPeriodicTimerSourceInfoResponsePtrOutput() PeriodicTimerSourceInfoResponsePtrOutput {
	return o
}

func (o PeriodicTimerSourceInfoResponsePtrOutput) ToPeriodicTimerSourceInfoResponsePtrOutputWithContext(ctx context.Context) PeriodicTimerSourceInfoResponsePtrOutput {
	return o
}

func (o PeriodicTimerSourceInfoResponsePtrOutput) Elem() PeriodicTimerSourceInfoResponseOutput {
	return o.ApplyT(func(v *PeriodicTimerSourceInfoResponse) PeriodicTimerSourceInfoResponse {
		if v != nil {
			return *v
		}
		var ret PeriodicTimerSourceInfoResponse
		return ret
	}).(PeriodicTimerSourceInfoResponseOutput)
}

func (o PeriodicTimerSourceInfoResponsePtrOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeriodicTimerSourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Schedule
	}).(pulumi.StringPtrOutput)
}

func (o PeriodicTimerSourceInfoResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeriodicTimerSourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o PeriodicTimerSourceInfoResponsePtrOutput) Topic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeriodicTimerSourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Topic
	}).(pulumi.StringPtrOutput)
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





type RefreshDetailsResponseInput interface {
	pulumi.Input

	ToRefreshDetailsResponseOutput() RefreshDetailsResponseOutput
	ToRefreshDetailsResponseOutputWithContext(context.Context) RefreshDetailsResponseOutput
}

type RefreshDetailsResponseArgs struct {
	ErrorManifestFile                pulumi.StringPtrInput `pulumi:"errorManifestFile"`
	InProgressRefreshJobId           pulumi.StringPtrInput `pulumi:"inProgressRefreshJobId"`
	LastCompletedRefreshJobTimeInUTC pulumi.StringPtrInput `pulumi:"lastCompletedRefreshJobTimeInUTC"`
	LastJob                          pulumi.StringPtrInput `pulumi:"lastJob"`
}

func (RefreshDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RefreshDetailsResponse)(nil)).Elem()
}

func (i RefreshDetailsResponseArgs) ToRefreshDetailsResponseOutput() RefreshDetailsResponseOutput {
	return i.ToRefreshDetailsResponseOutputWithContext(context.Background())
}

func (i RefreshDetailsResponseArgs) ToRefreshDetailsResponseOutputWithContext(ctx context.Context) RefreshDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RefreshDetailsResponseOutput)
}

func (i RefreshDetailsResponseArgs) ToRefreshDetailsResponsePtrOutput() RefreshDetailsResponsePtrOutput {
	return i.ToRefreshDetailsResponsePtrOutputWithContext(context.Background())
}

func (i RefreshDetailsResponseArgs) ToRefreshDetailsResponsePtrOutputWithContext(ctx context.Context) RefreshDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RefreshDetailsResponseOutput).ToRefreshDetailsResponsePtrOutputWithContext(ctx)
}









type RefreshDetailsResponsePtrInput interface {
	pulumi.Input

	ToRefreshDetailsResponsePtrOutput() RefreshDetailsResponsePtrOutput
	ToRefreshDetailsResponsePtrOutputWithContext(context.Context) RefreshDetailsResponsePtrOutput
}

type refreshDetailsResponsePtrType RefreshDetailsResponseArgs

func RefreshDetailsResponsePtr(v *RefreshDetailsResponseArgs) RefreshDetailsResponsePtrInput {
	return (*refreshDetailsResponsePtrType)(v)
}

func (*refreshDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RefreshDetailsResponse)(nil)).Elem()
}

func (i *refreshDetailsResponsePtrType) ToRefreshDetailsResponsePtrOutput() RefreshDetailsResponsePtrOutput {
	return i.ToRefreshDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *refreshDetailsResponsePtrType) ToRefreshDetailsResponsePtrOutputWithContext(ctx context.Context) RefreshDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RefreshDetailsResponsePtrOutput)
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

func (o RefreshDetailsResponseOutput) ToRefreshDetailsResponsePtrOutput() RefreshDetailsResponsePtrOutput {
	return o.ToRefreshDetailsResponsePtrOutputWithContext(context.Background())
}

func (o RefreshDetailsResponseOutput) ToRefreshDetailsResponsePtrOutputWithContext(ctx context.Context) RefreshDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RefreshDetailsResponse) *RefreshDetailsResponse {
		return &v
	}).(RefreshDetailsResponsePtrOutput)
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

type ResourceIdentity struct {
	Type *string `pulumi:"type"`
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}









type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

func (o ResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

func (o ResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type ResourceIdentityResponseInput interface {
	pulumi.Input

	ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput
	ToResourceIdentityResponseOutputWithContext(context.Context) ResourceIdentityResponseOutput
}

type ResourceIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ResourceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return i.ToResourceIdentityResponseOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput)
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput).ToResourceIdentityResponsePtrOutputWithContext(ctx)
}









type ResourceIdentityResponsePtrInput interface {
	pulumi.Input

	ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput
	ToResourceIdentityResponsePtrOutputWithContext(context.Context) ResourceIdentityResponsePtrOutput
}

type resourceIdentityResponsePtrType ResourceIdentityResponseArgs

func ResourceIdentityResponsePtr(v *ResourceIdentityResponseArgs) ResourceIdentityResponsePtrInput {
	return (*resourceIdentityResponsePtrType)(v)
}

func (*resourceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponsePtrOutput)
}

type ResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityResponse) *ResourceIdentityResponse {
		return &v
	}).(ResourceIdentityResponsePtrOutput)
}

func (o ResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) Elem() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) ResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityResponse
		return ret
	}).(ResourceIdentityResponseOutput)
}

func (o ResourceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ResourceMoveDetailsResponse struct {
	OperationInProgress                 *string `pulumi:"operationInProgress"`
	OperationInProgressLockTimeoutInUTC *string `pulumi:"operationInProgressLockTimeoutInUTC"`
}





type ResourceMoveDetailsResponseInput interface {
	pulumi.Input

	ToResourceMoveDetailsResponseOutput() ResourceMoveDetailsResponseOutput
	ToResourceMoveDetailsResponseOutputWithContext(context.Context) ResourceMoveDetailsResponseOutput
}

type ResourceMoveDetailsResponseArgs struct {
	OperationInProgress                 pulumi.StringPtrInput `pulumi:"operationInProgress"`
	OperationInProgressLockTimeoutInUTC pulumi.StringPtrInput `pulumi:"operationInProgressLockTimeoutInUTC"`
}

func (ResourceMoveDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMoveDetailsResponse)(nil)).Elem()
}

func (i ResourceMoveDetailsResponseArgs) ToResourceMoveDetailsResponseOutput() ResourceMoveDetailsResponseOutput {
	return i.ToResourceMoveDetailsResponseOutputWithContext(context.Background())
}

func (i ResourceMoveDetailsResponseArgs) ToResourceMoveDetailsResponseOutputWithContext(ctx context.Context) ResourceMoveDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMoveDetailsResponseOutput)
}

func (i ResourceMoveDetailsResponseArgs) ToResourceMoveDetailsResponsePtrOutput() ResourceMoveDetailsResponsePtrOutput {
	return i.ToResourceMoveDetailsResponsePtrOutputWithContext(context.Background())
}

func (i ResourceMoveDetailsResponseArgs) ToResourceMoveDetailsResponsePtrOutputWithContext(ctx context.Context) ResourceMoveDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMoveDetailsResponseOutput).ToResourceMoveDetailsResponsePtrOutputWithContext(ctx)
}









type ResourceMoveDetailsResponsePtrInput interface {
	pulumi.Input

	ToResourceMoveDetailsResponsePtrOutput() ResourceMoveDetailsResponsePtrOutput
	ToResourceMoveDetailsResponsePtrOutputWithContext(context.Context) ResourceMoveDetailsResponsePtrOutput
}

type resourceMoveDetailsResponsePtrType ResourceMoveDetailsResponseArgs

func ResourceMoveDetailsResponsePtr(v *ResourceMoveDetailsResponseArgs) ResourceMoveDetailsResponsePtrInput {
	return (*resourceMoveDetailsResponsePtrType)(v)
}

func (*resourceMoveDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceMoveDetailsResponse)(nil)).Elem()
}

func (i *resourceMoveDetailsResponsePtrType) ToResourceMoveDetailsResponsePtrOutput() ResourceMoveDetailsResponsePtrOutput {
	return i.ToResourceMoveDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *resourceMoveDetailsResponsePtrType) ToResourceMoveDetailsResponsePtrOutputWithContext(ctx context.Context) ResourceMoveDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMoveDetailsResponsePtrOutput)
}

type ResourceMoveDetailsResponseOutput struct{ *pulumi.OutputState }

func (ResourceMoveDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMoveDetailsResponse)(nil)).Elem()
}

func (o ResourceMoveDetailsResponseOutput) ToResourceMoveDetailsResponseOutput() ResourceMoveDetailsResponseOutput {
	return o
}

func (o ResourceMoveDetailsResponseOutput) ToResourceMoveDetailsResponseOutputWithContext(ctx context.Context) ResourceMoveDetailsResponseOutput {
	return o
}

func (o ResourceMoveDetailsResponseOutput) ToResourceMoveDetailsResponsePtrOutput() ResourceMoveDetailsResponsePtrOutput {
	return o.ToResourceMoveDetailsResponsePtrOutputWithContext(context.Background())
}

func (o ResourceMoveDetailsResponseOutput) ToResourceMoveDetailsResponsePtrOutputWithContext(ctx context.Context) ResourceMoveDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceMoveDetailsResponse) *ResourceMoveDetailsResponse {
		return &v
	}).(ResourceMoveDetailsResponsePtrOutput)
}

func (o ResourceMoveDetailsResponseOutput) OperationInProgress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMoveDetailsResponse) *string { return v.OperationInProgress }).(pulumi.StringPtrOutput)
}

func (o ResourceMoveDetailsResponseOutput) OperationInProgressLockTimeoutInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceMoveDetailsResponse) *string { return v.OperationInProgressLockTimeoutInUTC }).(pulumi.StringPtrOutput)
}

type ResourceMoveDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceMoveDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceMoveDetailsResponse)(nil)).Elem()
}

func (o ResourceMoveDetailsResponsePtrOutput) ToResourceMoveDetailsResponsePtrOutput() ResourceMoveDetailsResponsePtrOutput {
	return o
}

func (o ResourceMoveDetailsResponsePtrOutput) ToResourceMoveDetailsResponsePtrOutputWithContext(ctx context.Context) ResourceMoveDetailsResponsePtrOutput {
	return o
}

func (o ResourceMoveDetailsResponsePtrOutput) Elem() ResourceMoveDetailsResponseOutput {
	return o.ApplyT(func(v *ResourceMoveDetailsResponse) ResourceMoveDetailsResponse {
		if v != nil {
			return *v
		}
		var ret ResourceMoveDetailsResponse
		return ret
	}).(ResourceMoveDetailsResponseOutput)
}

func (o ResourceMoveDetailsResponsePtrOutput) OperationInProgress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceMoveDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.OperationInProgress
	}).(pulumi.StringPtrOutput)
}

func (o ResourceMoveDetailsResponsePtrOutput) OperationInProgressLockTimeoutInUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceMoveDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.OperationInProgressLockTimeoutInUTC
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

func (i RoleSinkInfoArgs) ToRoleSinkInfoPtrOutput() RoleSinkInfoPtrOutput {
	return i.ToRoleSinkInfoPtrOutputWithContext(context.Background())
}

func (i RoleSinkInfoArgs) ToRoleSinkInfoPtrOutputWithContext(ctx context.Context) RoleSinkInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleSinkInfoOutput).ToRoleSinkInfoPtrOutputWithContext(ctx)
}









type RoleSinkInfoPtrInput interface {
	pulumi.Input

	ToRoleSinkInfoPtrOutput() RoleSinkInfoPtrOutput
	ToRoleSinkInfoPtrOutputWithContext(context.Context) RoleSinkInfoPtrOutput
}

type roleSinkInfoPtrType RoleSinkInfoArgs

func RoleSinkInfoPtr(v *RoleSinkInfoArgs) RoleSinkInfoPtrInput {
	return (*roleSinkInfoPtrType)(v)
}

func (*roleSinkInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleSinkInfo)(nil)).Elem()
}

func (i *roleSinkInfoPtrType) ToRoleSinkInfoPtrOutput() RoleSinkInfoPtrOutput {
	return i.ToRoleSinkInfoPtrOutputWithContext(context.Background())
}

func (i *roleSinkInfoPtrType) ToRoleSinkInfoPtrOutputWithContext(ctx context.Context) RoleSinkInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleSinkInfoPtrOutput)
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

func (o RoleSinkInfoOutput) ToRoleSinkInfoPtrOutput() RoleSinkInfoPtrOutput {
	return o.ToRoleSinkInfoPtrOutputWithContext(context.Background())
}

func (o RoleSinkInfoOutput) ToRoleSinkInfoPtrOutputWithContext(ctx context.Context) RoleSinkInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoleSinkInfo) *RoleSinkInfo {
		return &v
	}).(RoleSinkInfoPtrOutput)
}

func (o RoleSinkInfoOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v RoleSinkInfo) string { return v.RoleId }).(pulumi.StringOutput)
}

type RoleSinkInfoPtrOutput struct{ *pulumi.OutputState }

func (RoleSinkInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleSinkInfo)(nil)).Elem()
}

func (o RoleSinkInfoPtrOutput) ToRoleSinkInfoPtrOutput() RoleSinkInfoPtrOutput {
	return o
}

func (o RoleSinkInfoPtrOutput) ToRoleSinkInfoPtrOutputWithContext(ctx context.Context) RoleSinkInfoPtrOutput {
	return o
}

func (o RoleSinkInfoPtrOutput) Elem() RoleSinkInfoOutput {
	return o.ApplyT(func(v *RoleSinkInfo) RoleSinkInfo {
		if v != nil {
			return *v
		}
		var ret RoleSinkInfo
		return ret
	}).(RoleSinkInfoOutput)
}

func (o RoleSinkInfoPtrOutput) RoleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleSinkInfo) *string {
		if v == nil {
			return nil
		}
		return &v.RoleId
	}).(pulumi.StringPtrOutput)
}

type RoleSinkInfoResponse struct {
	RoleId string `pulumi:"roleId"`
}





type RoleSinkInfoResponseInput interface {
	pulumi.Input

	ToRoleSinkInfoResponseOutput() RoleSinkInfoResponseOutput
	ToRoleSinkInfoResponseOutputWithContext(context.Context) RoleSinkInfoResponseOutput
}

type RoleSinkInfoResponseArgs struct {
	RoleId pulumi.StringInput `pulumi:"roleId"`
}

func (RoleSinkInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleSinkInfoResponse)(nil)).Elem()
}

func (i RoleSinkInfoResponseArgs) ToRoleSinkInfoResponseOutput() RoleSinkInfoResponseOutput {
	return i.ToRoleSinkInfoResponseOutputWithContext(context.Background())
}

func (i RoleSinkInfoResponseArgs) ToRoleSinkInfoResponseOutputWithContext(ctx context.Context) RoleSinkInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleSinkInfoResponseOutput)
}

func (i RoleSinkInfoResponseArgs) ToRoleSinkInfoResponsePtrOutput() RoleSinkInfoResponsePtrOutput {
	return i.ToRoleSinkInfoResponsePtrOutputWithContext(context.Background())
}

func (i RoleSinkInfoResponseArgs) ToRoleSinkInfoResponsePtrOutputWithContext(ctx context.Context) RoleSinkInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleSinkInfoResponseOutput).ToRoleSinkInfoResponsePtrOutputWithContext(ctx)
}









type RoleSinkInfoResponsePtrInput interface {
	pulumi.Input

	ToRoleSinkInfoResponsePtrOutput() RoleSinkInfoResponsePtrOutput
	ToRoleSinkInfoResponsePtrOutputWithContext(context.Context) RoleSinkInfoResponsePtrOutput
}

type roleSinkInfoResponsePtrType RoleSinkInfoResponseArgs

func RoleSinkInfoResponsePtr(v *RoleSinkInfoResponseArgs) RoleSinkInfoResponsePtrInput {
	return (*roleSinkInfoResponsePtrType)(v)
}

func (*roleSinkInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleSinkInfoResponse)(nil)).Elem()
}

func (i *roleSinkInfoResponsePtrType) ToRoleSinkInfoResponsePtrOutput() RoleSinkInfoResponsePtrOutput {
	return i.ToRoleSinkInfoResponsePtrOutputWithContext(context.Background())
}

func (i *roleSinkInfoResponsePtrType) ToRoleSinkInfoResponsePtrOutputWithContext(ctx context.Context) RoleSinkInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleSinkInfoResponsePtrOutput)
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

func (o RoleSinkInfoResponseOutput) ToRoleSinkInfoResponsePtrOutput() RoleSinkInfoResponsePtrOutput {
	return o.ToRoleSinkInfoResponsePtrOutputWithContext(context.Background())
}

func (o RoleSinkInfoResponseOutput) ToRoleSinkInfoResponsePtrOutputWithContext(ctx context.Context) RoleSinkInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoleSinkInfoResponse) *RoleSinkInfoResponse {
		return &v
	}).(RoleSinkInfoResponsePtrOutput)
}

func (o RoleSinkInfoResponseOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v RoleSinkInfoResponse) string { return v.RoleId }).(pulumi.StringOutput)
}

type RoleSinkInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (RoleSinkInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleSinkInfoResponse)(nil)).Elem()
}

func (o RoleSinkInfoResponsePtrOutput) ToRoleSinkInfoResponsePtrOutput() RoleSinkInfoResponsePtrOutput {
	return o
}

func (o RoleSinkInfoResponsePtrOutput) ToRoleSinkInfoResponsePtrOutputWithContext(ctx context.Context) RoleSinkInfoResponsePtrOutput {
	return o
}

func (o RoleSinkInfoResponsePtrOutput) Elem() RoleSinkInfoResponseOutput {
	return o.ApplyT(func(v *RoleSinkInfoResponse) RoleSinkInfoResponse {
		if v != nil {
			return *v
		}
		var ret RoleSinkInfoResponse
		return ret
	}).(RoleSinkInfoResponseOutput)
}

func (o RoleSinkInfoResponsePtrOutput) RoleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleSinkInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RoleId
	}).(pulumi.StringPtrOutput)
}

type SecretResponse struct {
	EncryptedSecret *AsymmetricEncryptedSecretResponse `pulumi:"encryptedSecret"`
	KeyVaultId      *string                            `pulumi:"keyVaultId"`
}





type SecretResponseInput interface {
	pulumi.Input

	ToSecretResponseOutput() SecretResponseOutput
	ToSecretResponseOutputWithContext(context.Context) SecretResponseOutput
}

type SecretResponseArgs struct {
	EncryptedSecret AsymmetricEncryptedSecretResponsePtrInput `pulumi:"encryptedSecret"`
	KeyVaultId      pulumi.StringPtrInput                     `pulumi:"keyVaultId"`
}

func (SecretResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretResponse)(nil)).Elem()
}

func (i SecretResponseArgs) ToSecretResponseOutput() SecretResponseOutput {
	return i.ToSecretResponseOutputWithContext(context.Background())
}

func (i SecretResponseArgs) ToSecretResponseOutputWithContext(ctx context.Context) SecretResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretResponseOutput)
}

func (i SecretResponseArgs) ToSecretResponsePtrOutput() SecretResponsePtrOutput {
	return i.ToSecretResponsePtrOutputWithContext(context.Background())
}

func (i SecretResponseArgs) ToSecretResponsePtrOutputWithContext(ctx context.Context) SecretResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretResponseOutput).ToSecretResponsePtrOutputWithContext(ctx)
}









type SecretResponsePtrInput interface {
	pulumi.Input

	ToSecretResponsePtrOutput() SecretResponsePtrOutput
	ToSecretResponsePtrOutputWithContext(context.Context) SecretResponsePtrOutput
}

type secretResponsePtrType SecretResponseArgs

func SecretResponsePtr(v *SecretResponseArgs) SecretResponsePtrInput {
	return (*secretResponsePtrType)(v)
}

func (*secretResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretResponse)(nil)).Elem()
}

func (i *secretResponsePtrType) ToSecretResponsePtrOutput() SecretResponsePtrOutput {
	return i.ToSecretResponsePtrOutputWithContext(context.Background())
}

func (i *secretResponsePtrType) ToSecretResponsePtrOutputWithContext(ctx context.Context) SecretResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretResponsePtrOutput)
}

type SecretResponseOutput struct{ *pulumi.OutputState }

func (SecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretResponse)(nil)).Elem()
}

func (o SecretResponseOutput) ToSecretResponseOutput() SecretResponseOutput {
	return o
}

func (o SecretResponseOutput) ToSecretResponseOutputWithContext(ctx context.Context) SecretResponseOutput {
	return o
}

func (o SecretResponseOutput) ToSecretResponsePtrOutput() SecretResponsePtrOutput {
	return o.ToSecretResponsePtrOutputWithContext(context.Background())
}

func (o SecretResponseOutput) ToSecretResponsePtrOutputWithContext(ctx context.Context) SecretResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretResponse) *SecretResponse {
		return &v
	}).(SecretResponsePtrOutput)
}

func (o SecretResponseOutput) EncryptedSecret() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v SecretResponse) *AsymmetricEncryptedSecretResponse { return v.EncryptedSecret }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o SecretResponseOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretResponse) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

type SecretResponsePtrOutput struct{ *pulumi.OutputState }

func (SecretResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretResponse)(nil)).Elem()
}

func (o SecretResponsePtrOutput) ToSecretResponsePtrOutput() SecretResponsePtrOutput {
	return o
}

func (o SecretResponsePtrOutput) ToSecretResponsePtrOutputWithContext(ctx context.Context) SecretResponsePtrOutput {
	return o
}

func (o SecretResponsePtrOutput) Elem() SecretResponseOutput {
	return o.ApplyT(func(v *SecretResponse) SecretResponse {
		if v != nil {
			return *v
		}
		var ret SecretResponse
		return ret
	}).(SecretResponseOutput)
}

func (o SecretResponsePtrOutput) EncryptedSecret() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v *SecretResponse) *AsymmetricEncryptedSecretResponse {
		if v == nil {
			return nil
		}
		return v.EncryptedSecret
	}).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o SecretResponsePtrOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultId
	}).(pulumi.StringPtrOutput)
}

type ShareAccessRightResponse struct {
	AccessType string `pulumi:"accessType"`
	ShareId    string `pulumi:"shareId"`
}





type ShareAccessRightResponseInput interface {
	pulumi.Input

	ToShareAccessRightResponseOutput() ShareAccessRightResponseOutput
	ToShareAccessRightResponseOutputWithContext(context.Context) ShareAccessRightResponseOutput
}

type ShareAccessRightResponseArgs struct {
	AccessType pulumi.StringInput `pulumi:"accessType"`
	ShareId    pulumi.StringInput `pulumi:"shareId"`
}

func (ShareAccessRightResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareAccessRightResponse)(nil)).Elem()
}

func (i ShareAccessRightResponseArgs) ToShareAccessRightResponseOutput() ShareAccessRightResponseOutput {
	return i.ToShareAccessRightResponseOutputWithContext(context.Background())
}

func (i ShareAccessRightResponseArgs) ToShareAccessRightResponseOutputWithContext(ctx context.Context) ShareAccessRightResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareAccessRightResponseOutput)
}





type ShareAccessRightResponseArrayInput interface {
	pulumi.Input

	ToShareAccessRightResponseArrayOutput() ShareAccessRightResponseArrayOutput
	ToShareAccessRightResponseArrayOutputWithContext(context.Context) ShareAccessRightResponseArrayOutput
}

type ShareAccessRightResponseArray []ShareAccessRightResponseInput

func (ShareAccessRightResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareAccessRightResponse)(nil)).Elem()
}

func (i ShareAccessRightResponseArray) ToShareAccessRightResponseArrayOutput() ShareAccessRightResponseArrayOutput {
	return i.ToShareAccessRightResponseArrayOutputWithContext(context.Background())
}

func (i ShareAccessRightResponseArray) ToShareAccessRightResponseArrayOutputWithContext(ctx context.Context) ShareAccessRightResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareAccessRightResponseArrayOutput)
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





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
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

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
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

type SubscriptionRegisteredFeaturesResponse struct {
	Name  *string `pulumi:"name"`
	State *string `pulumi:"state"`
}





type SubscriptionRegisteredFeaturesResponseInput interface {
	pulumi.Input

	ToSubscriptionRegisteredFeaturesResponseOutput() SubscriptionRegisteredFeaturesResponseOutput
	ToSubscriptionRegisteredFeaturesResponseOutputWithContext(context.Context) SubscriptionRegisteredFeaturesResponseOutput
}

type SubscriptionRegisteredFeaturesResponseArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (SubscriptionRegisteredFeaturesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionRegisteredFeaturesResponse)(nil)).Elem()
}

func (i SubscriptionRegisteredFeaturesResponseArgs) ToSubscriptionRegisteredFeaturesResponseOutput() SubscriptionRegisteredFeaturesResponseOutput {
	return i.ToSubscriptionRegisteredFeaturesResponseOutputWithContext(context.Background())
}

func (i SubscriptionRegisteredFeaturesResponseArgs) ToSubscriptionRegisteredFeaturesResponseOutputWithContext(ctx context.Context) SubscriptionRegisteredFeaturesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionRegisteredFeaturesResponseOutput)
}





type SubscriptionRegisteredFeaturesResponseArrayInput interface {
	pulumi.Input

	ToSubscriptionRegisteredFeaturesResponseArrayOutput() SubscriptionRegisteredFeaturesResponseArrayOutput
	ToSubscriptionRegisteredFeaturesResponseArrayOutputWithContext(context.Context) SubscriptionRegisteredFeaturesResponseArrayOutput
}

type SubscriptionRegisteredFeaturesResponseArray []SubscriptionRegisteredFeaturesResponseInput

func (SubscriptionRegisteredFeaturesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionRegisteredFeaturesResponse)(nil)).Elem()
}

func (i SubscriptionRegisteredFeaturesResponseArray) ToSubscriptionRegisteredFeaturesResponseArrayOutput() SubscriptionRegisteredFeaturesResponseArrayOutput {
	return i.ToSubscriptionRegisteredFeaturesResponseArrayOutputWithContext(context.Background())
}

func (i SubscriptionRegisteredFeaturesResponseArray) ToSubscriptionRegisteredFeaturesResponseArrayOutputWithContext(ctx context.Context) SubscriptionRegisteredFeaturesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionRegisteredFeaturesResponseArrayOutput)
}

type SubscriptionRegisteredFeaturesResponseOutput struct{ *pulumi.OutputState }

func (SubscriptionRegisteredFeaturesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionRegisteredFeaturesResponse)(nil)).Elem()
}

func (o SubscriptionRegisteredFeaturesResponseOutput) ToSubscriptionRegisteredFeaturesResponseOutput() SubscriptionRegisteredFeaturesResponseOutput {
	return o
}

func (o SubscriptionRegisteredFeaturesResponseOutput) ToSubscriptionRegisteredFeaturesResponseOutputWithContext(ctx context.Context) SubscriptionRegisteredFeaturesResponseOutput {
	return o
}

func (o SubscriptionRegisteredFeaturesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRegisteredFeaturesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubscriptionRegisteredFeaturesResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRegisteredFeaturesResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type SubscriptionRegisteredFeaturesResponseArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionRegisteredFeaturesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionRegisteredFeaturesResponse)(nil)).Elem()
}

func (o SubscriptionRegisteredFeaturesResponseArrayOutput) ToSubscriptionRegisteredFeaturesResponseArrayOutput() SubscriptionRegisteredFeaturesResponseArrayOutput {
	return o
}

func (o SubscriptionRegisteredFeaturesResponseArrayOutput) ToSubscriptionRegisteredFeaturesResponseArrayOutputWithContext(ctx context.Context) SubscriptionRegisteredFeaturesResponseArrayOutput {
	return o
}

func (o SubscriptionRegisteredFeaturesResponseArrayOutput) Index(i pulumi.IntInput) SubscriptionRegisteredFeaturesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubscriptionRegisteredFeaturesResponse {
		return vs[0].([]SubscriptionRegisteredFeaturesResponse)[vs[1].(int)]
	}).(SubscriptionRegisteredFeaturesResponseOutput)
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





type SymmetricKeyResponseInput interface {
	pulumi.Input

	ToSymmetricKeyResponseOutput() SymmetricKeyResponseOutput
	ToSymmetricKeyResponseOutputWithContext(context.Context) SymmetricKeyResponseOutput
}

type SymmetricKeyResponseArgs struct {
	ConnectionString AsymmetricEncryptedSecretResponsePtrInput `pulumi:"connectionString"`
}

func (SymmetricKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SymmetricKeyResponse)(nil)).Elem()
}

func (i SymmetricKeyResponseArgs) ToSymmetricKeyResponseOutput() SymmetricKeyResponseOutput {
	return i.ToSymmetricKeyResponseOutputWithContext(context.Background())
}

func (i SymmetricKeyResponseArgs) ToSymmetricKeyResponseOutputWithContext(ctx context.Context) SymmetricKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SymmetricKeyResponseOutput)
}

func (i SymmetricKeyResponseArgs) ToSymmetricKeyResponsePtrOutput() SymmetricKeyResponsePtrOutput {
	return i.ToSymmetricKeyResponsePtrOutputWithContext(context.Background())
}

func (i SymmetricKeyResponseArgs) ToSymmetricKeyResponsePtrOutputWithContext(ctx context.Context) SymmetricKeyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SymmetricKeyResponseOutput).ToSymmetricKeyResponsePtrOutputWithContext(ctx)
}









type SymmetricKeyResponsePtrInput interface {
	pulumi.Input

	ToSymmetricKeyResponsePtrOutput() SymmetricKeyResponsePtrOutput
	ToSymmetricKeyResponsePtrOutputWithContext(context.Context) SymmetricKeyResponsePtrOutput
}

type symmetricKeyResponsePtrType SymmetricKeyResponseArgs

func SymmetricKeyResponsePtr(v *SymmetricKeyResponseArgs) SymmetricKeyResponsePtrInput {
	return (*symmetricKeyResponsePtrType)(v)
}

func (*symmetricKeyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SymmetricKeyResponse)(nil)).Elem()
}

func (i *symmetricKeyResponsePtrType) ToSymmetricKeyResponsePtrOutput() SymmetricKeyResponsePtrOutput {
	return i.ToSymmetricKeyResponsePtrOutputWithContext(context.Background())
}

func (i *symmetricKeyResponsePtrType) ToSymmetricKeyResponsePtrOutputWithContext(ctx context.Context) SymmetricKeyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SymmetricKeyResponsePtrOutput)
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

func (o SymmetricKeyResponseOutput) ToSymmetricKeyResponsePtrOutput() SymmetricKeyResponsePtrOutput {
	return o.ToSymmetricKeyResponsePtrOutputWithContext(context.Background())
}

func (o SymmetricKeyResponseOutput) ToSymmetricKeyResponsePtrOutputWithContext(ctx context.Context) SymmetricKeyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SymmetricKeyResponse) *SymmetricKeyResponse {
		return &v
	}).(SymmetricKeyResponsePtrOutput)
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

type TrackingInfoResponse struct {
	CarrierName  *string `pulumi:"carrierName"`
	SerialNumber *string `pulumi:"serialNumber"`
	TrackingId   *string `pulumi:"trackingId"`
	TrackingUrl  *string `pulumi:"trackingUrl"`
}





type TrackingInfoResponseInput interface {
	pulumi.Input

	ToTrackingInfoResponseOutput() TrackingInfoResponseOutput
	ToTrackingInfoResponseOutputWithContext(context.Context) TrackingInfoResponseOutput
}

type TrackingInfoResponseArgs struct {
	CarrierName  pulumi.StringPtrInput `pulumi:"carrierName"`
	SerialNumber pulumi.StringPtrInput `pulumi:"serialNumber"`
	TrackingId   pulumi.StringPtrInput `pulumi:"trackingId"`
	TrackingUrl  pulumi.StringPtrInput `pulumi:"trackingUrl"`
}

func (TrackingInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackingInfoResponse)(nil)).Elem()
}

func (i TrackingInfoResponseArgs) ToTrackingInfoResponseOutput() TrackingInfoResponseOutput {
	return i.ToTrackingInfoResponseOutputWithContext(context.Background())
}

func (i TrackingInfoResponseArgs) ToTrackingInfoResponseOutputWithContext(ctx context.Context) TrackingInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackingInfoResponseOutput)
}

func (i TrackingInfoResponseArgs) ToTrackingInfoResponsePtrOutput() TrackingInfoResponsePtrOutput {
	return i.ToTrackingInfoResponsePtrOutputWithContext(context.Background())
}

func (i TrackingInfoResponseArgs) ToTrackingInfoResponsePtrOutputWithContext(ctx context.Context) TrackingInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackingInfoResponseOutput).ToTrackingInfoResponsePtrOutputWithContext(ctx)
}









type TrackingInfoResponsePtrInput interface {
	pulumi.Input

	ToTrackingInfoResponsePtrOutput() TrackingInfoResponsePtrOutput
	ToTrackingInfoResponsePtrOutputWithContext(context.Context) TrackingInfoResponsePtrOutput
}

type trackingInfoResponsePtrType TrackingInfoResponseArgs

func TrackingInfoResponsePtr(v *TrackingInfoResponseArgs) TrackingInfoResponsePtrInput {
	return (*trackingInfoResponsePtrType)(v)
}

func (*trackingInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TrackingInfoResponse)(nil)).Elem()
}

func (i *trackingInfoResponsePtrType) ToTrackingInfoResponsePtrOutput() TrackingInfoResponsePtrOutput {
	return i.ToTrackingInfoResponsePtrOutputWithContext(context.Background())
}

func (i *trackingInfoResponsePtrType) ToTrackingInfoResponsePtrOutputWithContext(ctx context.Context) TrackingInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackingInfoResponsePtrOutput)
}





type TrackingInfoResponseArrayInput interface {
	pulumi.Input

	ToTrackingInfoResponseArrayOutput() TrackingInfoResponseArrayOutput
	ToTrackingInfoResponseArrayOutputWithContext(context.Context) TrackingInfoResponseArrayOutput
}

type TrackingInfoResponseArray []TrackingInfoResponseInput

func (TrackingInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackingInfoResponse)(nil)).Elem()
}

func (i TrackingInfoResponseArray) ToTrackingInfoResponseArrayOutput() TrackingInfoResponseArrayOutput {
	return i.ToTrackingInfoResponseArrayOutputWithContext(context.Background())
}

func (i TrackingInfoResponseArray) ToTrackingInfoResponseArrayOutputWithContext(ctx context.Context) TrackingInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackingInfoResponseArrayOutput)
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

func (o TrackingInfoResponseOutput) ToTrackingInfoResponsePtrOutput() TrackingInfoResponsePtrOutput {
	return o.ToTrackingInfoResponsePtrOutputWithContext(context.Background())
}

func (o TrackingInfoResponseOutput) ToTrackingInfoResponsePtrOutputWithContext(ctx context.Context) TrackingInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrackingInfoResponse) *TrackingInfoResponse {
		return &v
	}).(TrackingInfoResponsePtrOutput)
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

type TrackingInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (TrackingInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrackingInfoResponse)(nil)).Elem()
}

func (o TrackingInfoResponsePtrOutput) ToTrackingInfoResponsePtrOutput() TrackingInfoResponsePtrOutput {
	return o
}

func (o TrackingInfoResponsePtrOutput) ToTrackingInfoResponsePtrOutputWithContext(ctx context.Context) TrackingInfoResponsePtrOutput {
	return o
}

func (o TrackingInfoResponsePtrOutput) Elem() TrackingInfoResponseOutput {
	return o.ApplyT(func(v *TrackingInfoResponse) TrackingInfoResponse {
		if v != nil {
			return *v
		}
		var ret TrackingInfoResponse
		return ret
	}).(TrackingInfoResponseOutput)
}

func (o TrackingInfoResponsePtrOutput) CarrierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrackingInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.CarrierName
	}).(pulumi.StringPtrOutput)
}

func (o TrackingInfoResponsePtrOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrackingInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.StringPtrOutput)
}

func (o TrackingInfoResponsePtrOutput) TrackingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrackingInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.TrackingId
	}).(pulumi.StringPtrOutput)
}

func (o TrackingInfoResponsePtrOutput) TrackingUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrackingInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.TrackingUrl
	}).(pulumi.StringPtrOutput)
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





type UserAccessRightResponseInput interface {
	pulumi.Input

	ToUserAccessRightResponseOutput() UserAccessRightResponseOutput
	ToUserAccessRightResponseOutputWithContext(context.Context) UserAccessRightResponseOutput
}

type UserAccessRightResponseArgs struct {
	AccessType pulumi.StringInput `pulumi:"accessType"`
	UserId     pulumi.StringInput `pulumi:"userId"`
}

func (UserAccessRightResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccessRightResponse)(nil)).Elem()
}

func (i UserAccessRightResponseArgs) ToUserAccessRightResponseOutput() UserAccessRightResponseOutput {
	return i.ToUserAccessRightResponseOutputWithContext(context.Background())
}

func (i UserAccessRightResponseArgs) ToUserAccessRightResponseOutputWithContext(ctx context.Context) UserAccessRightResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccessRightResponseOutput)
}





type UserAccessRightResponseArrayInput interface {
	pulumi.Input

	ToUserAccessRightResponseArrayOutput() UserAccessRightResponseArrayOutput
	ToUserAccessRightResponseArrayOutputWithContext(context.Context) UserAccessRightResponseArrayOutput
}

type UserAccessRightResponseArray []UserAccessRightResponseInput

func (UserAccessRightResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserAccessRightResponse)(nil)).Elem()
}

func (i UserAccessRightResponseArray) ToUserAccessRightResponseArrayOutput() UserAccessRightResponseArrayOutput {
	return i.ToUserAccessRightResponseArrayOutputWithContext(context.Background())
}

func (i UserAccessRightResponseArray) ToUserAccessRightResponseArrayOutputWithContext(ctx context.Context) UserAccessRightResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccessRightResponseArrayOutput)
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
	pulumi.RegisterOutputType(AddressPtrOutput{})
	pulumi.RegisterOutputType(AddressResponseOutput{})
	pulumi.RegisterOutputType(AddressResponsePtrOutput{})
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
	pulumi.RegisterOutputType(CniConfigResponseOutput{})
	pulumi.RegisterOutputType(CniConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(ComputeResourceOutput{})
	pulumi.RegisterOutputType(ComputeResourcePtrOutput{})
	pulumi.RegisterOutputType(ComputeResourceResponseOutput{})
	pulumi.RegisterOutputType(ComputeResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(ContactDetailsOutput{})
	pulumi.RegisterOutputType(ContactDetailsPtrOutput{})
	pulumi.RegisterOutputType(ContactDetailsResponseOutput{})
	pulumi.RegisterOutputType(ContactDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(DeviceSecretsResponseOutput{})
	pulumi.RegisterOutputType(EdgeProfileResponseOutput{})
	pulumi.RegisterOutputType(EdgeProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(EdgeProfileSubscriptionResponseOutput{})
	pulumi.RegisterOutputType(EdgeProfileSubscriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(EtcdInfoResponseOutput{})
	pulumi.RegisterOutputType(EtcdInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(FileSourceInfoOutput{})
	pulumi.RegisterOutputType(FileSourceInfoPtrOutput{})
	pulumi.RegisterOutputType(FileSourceInfoResponseOutput{})
	pulumi.RegisterOutputType(FileSourceInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageRepositoryCredentialOutput{})
	pulumi.RegisterOutputType(ImageRepositoryCredentialPtrOutput{})
	pulumi.RegisterOutputType(ImageRepositoryCredentialResponseOutput{})
	pulumi.RegisterOutputType(ImageRepositoryCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(IoTDeviceInfoOutput{})
	pulumi.RegisterOutputType(IoTDeviceInfoPtrOutput{})
	pulumi.RegisterOutputType(IoTDeviceInfoResponseOutput{})
	pulumi.RegisterOutputType(IoTDeviceInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(IoTEdgeAgentInfoOutput{})
	pulumi.RegisterOutputType(IoTEdgeAgentInfoPtrOutput{})
	pulumi.RegisterOutputType(IoTEdgeAgentInfoResponseOutput{})
	pulumi.RegisterOutputType(IoTEdgeAgentInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(KubernetesClusterInfoOutput{})
	pulumi.RegisterOutputType(KubernetesClusterInfoPtrOutput{})
	pulumi.RegisterOutputType(KubernetesClusterInfoResponseOutput{})
	pulumi.RegisterOutputType(KubernetesClusterInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(KubernetesIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(KubernetesIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(KubernetesRoleComputeOutput{})
	pulumi.RegisterOutputType(KubernetesRoleComputePtrOutput{})
	pulumi.RegisterOutputType(KubernetesRoleComputeResponseOutput{})
	pulumi.RegisterOutputType(KubernetesRoleComputeResponsePtrOutput{})
	pulumi.RegisterOutputType(KubernetesRoleNetworkResponseOutput{})
	pulumi.RegisterOutputType(KubernetesRoleNetworkResponsePtrOutput{})
	pulumi.RegisterOutputType(KubernetesRoleResourcesOutput{})
	pulumi.RegisterOutputType(KubernetesRoleResourcesPtrOutput{})
	pulumi.RegisterOutputType(KubernetesRoleResourcesResponseOutput{})
	pulumi.RegisterOutputType(KubernetesRoleResourcesResponsePtrOutput{})
	pulumi.RegisterOutputType(KubernetesRoleStorageOutput{})
	pulumi.RegisterOutputType(KubernetesRoleStoragePtrOutput{})
	pulumi.RegisterOutputType(KubernetesRoleStorageClassInfoResponseOutput{})
	pulumi.RegisterOutputType(KubernetesRoleStorageClassInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(KubernetesRoleStorageResponseOutput{})
	pulumi.RegisterOutputType(KubernetesRoleStorageResponsePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerConfigResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancerConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(MetricConfigurationOutput{})
	pulumi.RegisterOutputType(MetricConfigurationArrayOutput{})
	pulumi.RegisterOutputType(MetricConfigurationResponseOutput{})
	pulumi.RegisterOutputType(MetricConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(MetricCounterOutput{})
	pulumi.RegisterOutputType(MetricCounterArrayOutput{})
	pulumi.RegisterOutputType(MetricCounterResponseOutput{})
	pulumi.RegisterOutputType(MetricCounterResponseArrayOutput{})
	pulumi.RegisterOutputType(MetricCounterSetOutput{})
	pulumi.RegisterOutputType(MetricCounterSetArrayOutput{})
	pulumi.RegisterOutputType(MetricCounterSetResponseOutput{})
	pulumi.RegisterOutputType(MetricCounterSetResponseArrayOutput{})
	pulumi.RegisterOutputType(MetricDimensionOutput{})
	pulumi.RegisterOutputType(MetricDimensionArrayOutput{})
	pulumi.RegisterOutputType(MetricDimensionResponseOutput{})
	pulumi.RegisterOutputType(MetricDimensionResponseArrayOutput{})
	pulumi.RegisterOutputType(MountPointMapOutput{})
	pulumi.RegisterOutputType(MountPointMapArrayOutput{})
	pulumi.RegisterOutputType(MountPointMapResponseOutput{})
	pulumi.RegisterOutputType(MountPointMapResponseArrayOutput{})
	pulumi.RegisterOutputType(NodeInfoResponseOutput{})
	pulumi.RegisterOutputType(NodeInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(OrderStatusResponseOutput{})
	pulumi.RegisterOutputType(OrderStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(OrderStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(PeriodicTimerSourceInfoOutput{})
	pulumi.RegisterOutputType(PeriodicTimerSourceInfoPtrOutput{})
	pulumi.RegisterOutputType(PeriodicTimerSourceInfoResponseOutput{})
	pulumi.RegisterOutputType(PeriodicTimerSourceInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(RefreshDetailsOutput{})
	pulumi.RegisterOutputType(RefreshDetailsPtrOutput{})
	pulumi.RegisterOutputType(RefreshDetailsResponseOutput{})
	pulumi.RegisterOutputType(RefreshDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceMoveDetailsResponseOutput{})
	pulumi.RegisterOutputType(ResourceMoveDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(RoleSinkInfoOutput{})
	pulumi.RegisterOutputType(RoleSinkInfoPtrOutput{})
	pulumi.RegisterOutputType(RoleSinkInfoResponseOutput{})
	pulumi.RegisterOutputType(RoleSinkInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(SecretResponseOutput{})
	pulumi.RegisterOutputType(SecretResponsePtrOutput{})
	pulumi.RegisterOutputType(ShareAccessRightResponseOutput{})
	pulumi.RegisterOutputType(ShareAccessRightResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SubscriptionRegisteredFeaturesResponseOutput{})
	pulumi.RegisterOutputType(SubscriptionRegisteredFeaturesResponseArrayOutput{})
	pulumi.RegisterOutputType(SymmetricKeyOutput{})
	pulumi.RegisterOutputType(SymmetricKeyPtrOutput{})
	pulumi.RegisterOutputType(SymmetricKeyResponseOutput{})
	pulumi.RegisterOutputType(SymmetricKeyResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TrackingInfoResponseOutput{})
	pulumi.RegisterOutputType(TrackingInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(TrackingInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAccessRightOutput{})
	pulumi.RegisterOutputType(UserAccessRightArrayOutput{})
	pulumi.RegisterOutputType(UserAccessRightResponseOutput{})
	pulumi.RegisterOutputType(UserAccessRightResponseArrayOutput{})
}

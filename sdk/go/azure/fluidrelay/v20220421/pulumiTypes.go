


package v20220421

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomerManagedKeyEncryptionProperties struct {
	KeyEncryptionKeyIdentity *CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity `pulumi:"keyEncryptionKeyIdentity"`
	KeyEncryptionKeyUrl      *string                                                         `pulumi:"keyEncryptionKeyUrl"`
}





type CustomerManagedKeyEncryptionPropertiesInput interface {
	pulumi.Input

	ToCustomerManagedKeyEncryptionPropertiesOutput() CustomerManagedKeyEncryptionPropertiesOutput
	ToCustomerManagedKeyEncryptionPropertiesOutputWithContext(context.Context) CustomerManagedKeyEncryptionPropertiesOutput
}

type CustomerManagedKeyEncryptionPropertiesArgs struct {
	KeyEncryptionKeyIdentity CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrInput `pulumi:"keyEncryptionKeyIdentity"`
	KeyEncryptionKeyUrl      pulumi.StringPtrInput                                                  `pulumi:"keyEncryptionKeyUrl"`
}

func (CustomerManagedKeyEncryptionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerManagedKeyEncryptionProperties)(nil)).Elem()
}

func (i CustomerManagedKeyEncryptionPropertiesArgs) ToCustomerManagedKeyEncryptionPropertiesOutput() CustomerManagedKeyEncryptionPropertiesOutput {
	return i.ToCustomerManagedKeyEncryptionPropertiesOutputWithContext(context.Background())
}

func (i CustomerManagedKeyEncryptionPropertiesArgs) ToCustomerManagedKeyEncryptionPropertiesOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerManagedKeyEncryptionPropertiesOutput)
}

func (i CustomerManagedKeyEncryptionPropertiesArgs) ToCustomerManagedKeyEncryptionPropertiesPtrOutput() CustomerManagedKeyEncryptionPropertiesPtrOutput {
	return i.ToCustomerManagedKeyEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i CustomerManagedKeyEncryptionPropertiesArgs) ToCustomerManagedKeyEncryptionPropertiesPtrOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerManagedKeyEncryptionPropertiesOutput).ToCustomerManagedKeyEncryptionPropertiesPtrOutputWithContext(ctx)
}









type CustomerManagedKeyEncryptionPropertiesPtrInput interface {
	pulumi.Input

	ToCustomerManagedKeyEncryptionPropertiesPtrOutput() CustomerManagedKeyEncryptionPropertiesPtrOutput
	ToCustomerManagedKeyEncryptionPropertiesPtrOutputWithContext(context.Context) CustomerManagedKeyEncryptionPropertiesPtrOutput
}

type customerManagedKeyEncryptionPropertiesPtrType CustomerManagedKeyEncryptionPropertiesArgs

func CustomerManagedKeyEncryptionPropertiesPtr(v *CustomerManagedKeyEncryptionPropertiesArgs) CustomerManagedKeyEncryptionPropertiesPtrInput {
	return (*customerManagedKeyEncryptionPropertiesPtrType)(v)
}

func (*customerManagedKeyEncryptionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerManagedKeyEncryptionProperties)(nil)).Elem()
}

func (i *customerManagedKeyEncryptionPropertiesPtrType) ToCustomerManagedKeyEncryptionPropertiesPtrOutput() CustomerManagedKeyEncryptionPropertiesPtrOutput {
	return i.ToCustomerManagedKeyEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i *customerManagedKeyEncryptionPropertiesPtrType) ToCustomerManagedKeyEncryptionPropertiesPtrOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerManagedKeyEncryptionPropertiesPtrOutput)
}

type CustomerManagedKeyEncryptionPropertiesOutput struct{ *pulumi.OutputState }

func (CustomerManagedKeyEncryptionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerManagedKeyEncryptionProperties)(nil)).Elem()
}

func (o CustomerManagedKeyEncryptionPropertiesOutput) ToCustomerManagedKeyEncryptionPropertiesOutput() CustomerManagedKeyEncryptionPropertiesOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesOutput) ToCustomerManagedKeyEncryptionPropertiesOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesOutput) ToCustomerManagedKeyEncryptionPropertiesPtrOutput() CustomerManagedKeyEncryptionPropertiesPtrOutput {
	return o.ToCustomerManagedKeyEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (o CustomerManagedKeyEncryptionPropertiesOutput) ToCustomerManagedKeyEncryptionPropertiesPtrOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomerManagedKeyEncryptionProperties) *CustomerManagedKeyEncryptionProperties {
		return &v
	}).(CustomerManagedKeyEncryptionPropertiesPtrOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesOutput) KeyEncryptionKeyIdentity() CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput {
	return o.ApplyT(func(v CustomerManagedKeyEncryptionProperties) *CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity {
		return v.KeyEncryptionKeyIdentity
	}).(CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesOutput) KeyEncryptionKeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomerManagedKeyEncryptionProperties) *string { return v.KeyEncryptionKeyUrl }).(pulumi.StringPtrOutput)
}

type CustomerManagedKeyEncryptionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CustomerManagedKeyEncryptionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerManagedKeyEncryptionProperties)(nil)).Elem()
}

func (o CustomerManagedKeyEncryptionPropertiesPtrOutput) ToCustomerManagedKeyEncryptionPropertiesPtrOutput() CustomerManagedKeyEncryptionPropertiesPtrOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesPtrOutput) ToCustomerManagedKeyEncryptionPropertiesPtrOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesPtrOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesPtrOutput) Elem() CustomerManagedKeyEncryptionPropertiesOutput {
	return o.ApplyT(func(v *CustomerManagedKeyEncryptionProperties) CustomerManagedKeyEncryptionProperties {
		if v != nil {
			return *v
		}
		var ret CustomerManagedKeyEncryptionProperties
		return ret
	}).(CustomerManagedKeyEncryptionPropertiesOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesPtrOutput) KeyEncryptionKeyIdentity() CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyEncryptionProperties) *CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKeyIdentity
	}).(CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesPtrOutput) KeyEncryptionKeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyEncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKeyUrl
	}).(pulumi.StringPtrOutput)
}

type CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity struct {
	IdentityType                   *CmkIdentityType `pulumi:"identityType"`
	UserAssignedIdentityResourceId *string          `pulumi:"userAssignedIdentityResourceId"`
}





type CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityInput interface {
	pulumi.Input

	ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput() CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput
	ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutputWithContext(context.Context) CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput
}

type CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityArgs struct {
	IdentityType                   CmkIdentityTypePtrInput `pulumi:"identityType"`
	UserAssignedIdentityResourceId pulumi.StringPtrInput   `pulumi:"userAssignedIdentityResourceId"`
}

func (CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity)(nil)).Elem()
}

func (i CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityArgs) ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput() CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput {
	return i.ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutputWithContext(context.Background())
}

func (i CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityArgs) ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput)
}

func (i CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityArgs) ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput() CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput {
	return i.ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutputWithContext(context.Background())
}

func (i CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityArgs) ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput).ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutputWithContext(ctx)
}









type CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrInput interface {
	pulumi.Input

	ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput() CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput
	ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutputWithContext(context.Context) CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput
}

type customerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrType CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityArgs

func CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtr(v *CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityArgs) CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrInput {
	return (*customerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrType)(v)
}

func (*customerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity)(nil)).Elem()
}

func (i *customerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrType) ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput() CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput {
	return i.ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutputWithContext(context.Background())
}

func (i *customerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrType) ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput)
}

type CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput struct{ *pulumi.OutputState }

func (CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity)(nil)).Elem()
}

func (o CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput) ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput() CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput) ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput) ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput() CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput {
	return o.ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutputWithContext(context.Background())
}

func (o CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput) ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity) *CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity {
		return &v
	}).(CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput) IdentityType() CmkIdentityTypePtrOutput {
	return o.ApplyT(func(v CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity) *CmkIdentityType {
		return v.IdentityType
	}).(CmkIdentityTypePtrOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput) UserAssignedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity) *string {
		return v.UserAssignedIdentityResourceId
	}).(pulumi.StringPtrOutput)
}

type CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput struct{ *pulumi.OutputState }

func (CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity)(nil)).Elem()
}

func (o CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput) ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput() CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput) ToCustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput) Elem() CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput {
	return o.ApplyT(func(v *CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity) CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity {
		if v != nil {
			return *v
		}
		var ret CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity
		return ret
	}).(CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput) IdentityType() CmkIdentityTypePtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity) *CmkIdentityType {
		if v == nil {
			return nil
		}
		return v.IdentityType
	}).(CmkIdentityTypePtrOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput) UserAssignedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentity) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentityResourceId
	}).(pulumi.StringPtrOutput)
}

type CustomerManagedKeyEncryptionPropertiesResponse struct {
	KeyEncryptionKeyIdentity *CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity `pulumi:"keyEncryptionKeyIdentity"`
	KeyEncryptionKeyUrl      *string                                                                 `pulumi:"keyEncryptionKeyUrl"`
}

type CustomerManagedKeyEncryptionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CustomerManagedKeyEncryptionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerManagedKeyEncryptionPropertiesResponse)(nil)).Elem()
}

func (o CustomerManagedKeyEncryptionPropertiesResponseOutput) ToCustomerManagedKeyEncryptionPropertiesResponseOutput() CustomerManagedKeyEncryptionPropertiesResponseOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesResponseOutput) ToCustomerManagedKeyEncryptionPropertiesResponseOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesResponseOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesResponseOutput) KeyEncryptionKeyIdentity() CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput {
	return o.ApplyT(func(v CustomerManagedKeyEncryptionPropertiesResponse) *CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity {
		return v.KeyEncryptionKeyIdentity
	}).(CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesResponseOutput) KeyEncryptionKeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomerManagedKeyEncryptionPropertiesResponse) *string { return v.KeyEncryptionKeyUrl }).(pulumi.StringPtrOutput)
}

type CustomerManagedKeyEncryptionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomerManagedKeyEncryptionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerManagedKeyEncryptionPropertiesResponse)(nil)).Elem()
}

func (o CustomerManagedKeyEncryptionPropertiesResponsePtrOutput) ToCustomerManagedKeyEncryptionPropertiesResponsePtrOutput() CustomerManagedKeyEncryptionPropertiesResponsePtrOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesResponsePtrOutput) ToCustomerManagedKeyEncryptionPropertiesResponsePtrOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesResponsePtrOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesResponsePtrOutput) Elem() CustomerManagedKeyEncryptionPropertiesResponseOutput {
	return o.ApplyT(func(v *CustomerManagedKeyEncryptionPropertiesResponse) CustomerManagedKeyEncryptionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CustomerManagedKeyEncryptionPropertiesResponse
		return ret
	}).(CustomerManagedKeyEncryptionPropertiesResponseOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesResponsePtrOutput) KeyEncryptionKeyIdentity() CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyEncryptionPropertiesResponse) *CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKeyIdentity
	}).(CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesResponsePtrOutput) KeyEncryptionKeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyEncryptionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKeyUrl
	}).(pulumi.StringPtrOutput)
}

type CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity struct {
	IdentityType                   *string `pulumi:"identityType"`
	UserAssignedIdentityResourceId *string `pulumi:"userAssignedIdentityResourceId"`
}

type CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutput struct{ *pulumi.OutputState }

func (CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity)(nil)).Elem()
}

func (o CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutput) ToCustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutput() CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutput) ToCustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutput) IdentityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity) *string {
		return v.IdentityType
	}).(pulumi.StringPtrOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutput) UserAssignedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity) *string {
		return v.UserAssignedIdentityResourceId
	}).(pulumi.StringPtrOutput)
}

type CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput struct{ *pulumi.OutputState }

func (CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity)(nil)).Elem()
}

func (o CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput) ToCustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput() CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput) ToCustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutputWithContext(ctx context.Context) CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput {
	return o
}

func (o CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput) Elem() CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutput {
	return o.ApplyT(func(v *CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity) CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity {
		if v != nil {
			return *v
		}
		var ret CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity
		return ret
	}).(CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput) IdentityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity) *string {
		if v == nil {
			return nil
		}
		return v.IdentityType
	}).(pulumi.StringPtrOutput)
}

func (o CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput) UserAssignedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentity) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentityResourceId
	}).(pulumi.StringPtrOutput)
}

type EncryptionProperties struct {
	CustomerManagedKeyEncryption *CustomerManagedKeyEncryptionProperties `pulumi:"customerManagedKeyEncryption"`
}





type EncryptionPropertiesInput interface {
	pulumi.Input

	ToEncryptionPropertiesOutput() EncryptionPropertiesOutput
	ToEncryptionPropertiesOutputWithContext(context.Context) EncryptionPropertiesOutput
}

type EncryptionPropertiesArgs struct {
	CustomerManagedKeyEncryption CustomerManagedKeyEncryptionPropertiesPtrInput `pulumi:"customerManagedKeyEncryption"`
}

func (EncryptionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperties)(nil)).Elem()
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesOutput() EncryptionPropertiesOutput {
	return i.ToEncryptionPropertiesOutputWithContext(context.Background())
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesOutputWithContext(ctx context.Context) EncryptionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesOutput)
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return i.ToEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesOutput).ToEncryptionPropertiesPtrOutputWithContext(ctx)
}









type EncryptionPropertiesPtrInput interface {
	pulumi.Input

	ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput
	ToEncryptionPropertiesPtrOutputWithContext(context.Context) EncryptionPropertiesPtrOutput
}

type encryptionPropertiesPtrType EncryptionPropertiesArgs

func EncryptionPropertiesPtr(v *EncryptionPropertiesArgs) EncryptionPropertiesPtrInput {
	return (*encryptionPropertiesPtrType)(v)
}

func (*encryptionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperties)(nil)).Elem()
}

func (i *encryptionPropertiesPtrType) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return i.ToEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i *encryptionPropertiesPtrType) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesPtrOutput)
}

type EncryptionPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperties)(nil)).Elem()
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesOutput() EncryptionPropertiesOutput {
	return o
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesOutputWithContext(ctx context.Context) EncryptionPropertiesOutput {
	return o
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return o.ToEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionProperties) *EncryptionProperties {
		return &v
	}).(EncryptionPropertiesPtrOutput)
}

func (o EncryptionPropertiesOutput) CustomerManagedKeyEncryption() CustomerManagedKeyEncryptionPropertiesPtrOutput {
	return o.ApplyT(func(v EncryptionProperties) *CustomerManagedKeyEncryptionProperties {
		return v.CustomerManagedKeyEncryption
	}).(CustomerManagedKeyEncryptionPropertiesPtrOutput)
}

type EncryptionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperties)(nil)).Elem()
}

func (o EncryptionPropertiesPtrOutput) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return o
}

func (o EncryptionPropertiesPtrOutput) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return o
}

func (o EncryptionPropertiesPtrOutput) Elem() EncryptionPropertiesOutput {
	return o.ApplyT(func(v *EncryptionProperties) EncryptionProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionProperties
		return ret
	}).(EncryptionPropertiesOutput)
}

func (o EncryptionPropertiesPtrOutput) CustomerManagedKeyEncryption() CustomerManagedKeyEncryptionPropertiesPtrOutput {
	return o.ApplyT(func(v *EncryptionProperties) *CustomerManagedKeyEncryptionProperties {
		if v == nil {
			return nil
		}
		return v.CustomerManagedKeyEncryption
	}).(CustomerManagedKeyEncryptionPropertiesPtrOutput)
}

type EncryptionPropertiesResponse struct {
	CustomerManagedKeyEncryption *CustomerManagedKeyEncryptionPropertiesResponse `pulumi:"customerManagedKeyEncryption"`
}

type EncryptionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesResponse)(nil)).Elem()
}

func (o EncryptionPropertiesResponseOutput) ToEncryptionPropertiesResponseOutput() EncryptionPropertiesResponseOutput {
	return o
}

func (o EncryptionPropertiesResponseOutput) ToEncryptionPropertiesResponseOutputWithContext(ctx context.Context) EncryptionPropertiesResponseOutput {
	return o
}

func (o EncryptionPropertiesResponseOutput) CustomerManagedKeyEncryption() CustomerManagedKeyEncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesResponse) *CustomerManagedKeyEncryptionPropertiesResponse {
		return v.CustomerManagedKeyEncryption
	}).(CustomerManagedKeyEncryptionPropertiesResponsePtrOutput)
}

type EncryptionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesResponse)(nil)).Elem()
}

func (o EncryptionPropertiesResponsePtrOutput) ToEncryptionPropertiesResponsePtrOutput() EncryptionPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionPropertiesResponsePtrOutput) ToEncryptionPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionPropertiesResponsePtrOutput) Elem() EncryptionPropertiesResponseOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) EncryptionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertiesResponse
		return ret
	}).(EncryptionPropertiesResponseOutput)
}

func (o EncryptionPropertiesResponsePtrOutput) CustomerManagedKeyEncryption() CustomerManagedKeyEncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) *CustomerManagedKeyEncryptionPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.CustomerManagedKeyEncryption
	}).(CustomerManagedKeyEncryptionPropertiesResponsePtrOutput)
}

type FluidRelayEndpointsResponse struct {
	OrdererEndpoints []string `pulumi:"ordererEndpoints"`
	StorageEndpoints []string `pulumi:"storageEndpoints"`
}

type FluidRelayEndpointsResponseOutput struct{ *pulumi.OutputState }

func (FluidRelayEndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FluidRelayEndpointsResponse)(nil)).Elem()
}

func (o FluidRelayEndpointsResponseOutput) ToFluidRelayEndpointsResponseOutput() FluidRelayEndpointsResponseOutput {
	return o
}

func (o FluidRelayEndpointsResponseOutput) ToFluidRelayEndpointsResponseOutputWithContext(ctx context.Context) FluidRelayEndpointsResponseOutput {
	return o
}

func (o FluidRelayEndpointsResponseOutput) OrdererEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FluidRelayEndpointsResponse) []string { return v.OrdererEndpoints }).(pulumi.StringArrayOutput)
}

func (o FluidRelayEndpointsResponseOutput) StorageEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FluidRelayEndpointsResponse) []string { return v.StorageEndpoints }).(pulumi.StringArrayOutput)
}

type Identity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
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

func (o IdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v Identity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
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

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *Identity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityResponse struct {
	PrincipalId            string                                            `pulumi:"principalId"`
	TenantId               string                                            `pulumi:"tenantId"`
	Type                   *string                                           `pulumi:"type"`
	UserAssignedIdentities map[string]IdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
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

func (o IdentityResponseOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
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

func (o IdentityResponsePtrOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
}

type IdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type IdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutput() IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o IdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type IdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutput() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) IdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]IdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(IdentityResponseUserAssignedIdentitiesOutput)
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
	pulumi.RegisterOutputType(CustomerManagedKeyEncryptionPropertiesOutput{})
	pulumi.RegisterOutputType(CustomerManagedKeyEncryptionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityOutput{})
	pulumi.RegisterOutputType(CustomerManagedKeyEncryptionPropertiesKeyEncryptionKeyIdentityPtrOutput{})
	pulumi.RegisterOutputType(CustomerManagedKeyEncryptionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CustomerManagedKeyEncryptionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityOutput{})
	pulumi.RegisterOutputType(CustomerManagedKeyEncryptionPropertiesResponseKeyEncryptionKeyIdentityPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(FluidRelayEndpointsResponseOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

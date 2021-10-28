


package v20190501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActiveDirectoryObject struct {
	ObjectId *string `pulumi:"objectId"`
	TenantId *string `pulumi:"tenantId"`
}





type ActiveDirectoryObjectInput interface {
	pulumi.Input

	ToActiveDirectoryObjectOutput() ActiveDirectoryObjectOutput
	ToActiveDirectoryObjectOutputWithContext(context.Context) ActiveDirectoryObjectOutput
}

type ActiveDirectoryObjectArgs struct {
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (ActiveDirectoryObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryObject)(nil)).Elem()
}

func (i ActiveDirectoryObjectArgs) ToActiveDirectoryObjectOutput() ActiveDirectoryObjectOutput {
	return i.ToActiveDirectoryObjectOutputWithContext(context.Background())
}

func (i ActiveDirectoryObjectArgs) ToActiveDirectoryObjectOutputWithContext(ctx context.Context) ActiveDirectoryObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryObjectOutput)
}

func (i ActiveDirectoryObjectArgs) ToActiveDirectoryObjectPtrOutput() ActiveDirectoryObjectPtrOutput {
	return i.ToActiveDirectoryObjectPtrOutputWithContext(context.Background())
}

func (i ActiveDirectoryObjectArgs) ToActiveDirectoryObjectPtrOutputWithContext(ctx context.Context) ActiveDirectoryObjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryObjectOutput).ToActiveDirectoryObjectPtrOutputWithContext(ctx)
}









type ActiveDirectoryObjectPtrInput interface {
	pulumi.Input

	ToActiveDirectoryObjectPtrOutput() ActiveDirectoryObjectPtrOutput
	ToActiveDirectoryObjectPtrOutputWithContext(context.Context) ActiveDirectoryObjectPtrOutput
}

type activeDirectoryObjectPtrType ActiveDirectoryObjectArgs

func ActiveDirectoryObjectPtr(v *ActiveDirectoryObjectArgs) ActiveDirectoryObjectPtrInput {
	return (*activeDirectoryObjectPtrType)(v)
}

func (*activeDirectoryObjectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryObject)(nil)).Elem()
}

func (i *activeDirectoryObjectPtrType) ToActiveDirectoryObjectPtrOutput() ActiveDirectoryObjectPtrOutput {
	return i.ToActiveDirectoryObjectPtrOutputWithContext(context.Background())
}

func (i *activeDirectoryObjectPtrType) ToActiveDirectoryObjectPtrOutputWithContext(ctx context.Context) ActiveDirectoryObjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryObjectPtrOutput)
}

type ActiveDirectoryObjectOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryObject)(nil)).Elem()
}

func (o ActiveDirectoryObjectOutput) ToActiveDirectoryObjectOutput() ActiveDirectoryObjectOutput {
	return o
}

func (o ActiveDirectoryObjectOutput) ToActiveDirectoryObjectOutputWithContext(ctx context.Context) ActiveDirectoryObjectOutput {
	return o
}

func (o ActiveDirectoryObjectOutput) ToActiveDirectoryObjectPtrOutput() ActiveDirectoryObjectPtrOutput {
	return o.ToActiveDirectoryObjectPtrOutputWithContext(context.Background())
}

func (o ActiveDirectoryObjectOutput) ToActiveDirectoryObjectPtrOutputWithContext(ctx context.Context) ActiveDirectoryObjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActiveDirectoryObject) *ActiveDirectoryObject {
		return &v
	}).(ActiveDirectoryObjectPtrOutput)
}

func (o ActiveDirectoryObjectOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryObject) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryObjectOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryObject) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ActiveDirectoryObjectPtrOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryObjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryObject)(nil)).Elem()
}

func (o ActiveDirectoryObjectPtrOutput) ToActiveDirectoryObjectPtrOutput() ActiveDirectoryObjectPtrOutput {
	return o
}

func (o ActiveDirectoryObjectPtrOutput) ToActiveDirectoryObjectPtrOutputWithContext(ctx context.Context) ActiveDirectoryObjectPtrOutput {
	return o
}

func (o ActiveDirectoryObjectPtrOutput) Elem() ActiveDirectoryObjectOutput {
	return o.ApplyT(func(v *ActiveDirectoryObject) ActiveDirectoryObject {
		if v != nil {
			return *v
		}
		var ret ActiveDirectoryObject
		return ret
	}).(ActiveDirectoryObjectOutput)
}

func (o ActiveDirectoryObjectPtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryObject) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryObjectPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryObject) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type ActiveDirectoryObjectResponse struct {
	ObjectId *string `pulumi:"objectId"`
	TenantId *string `pulumi:"tenantId"`
}





type ActiveDirectoryObjectResponseInput interface {
	pulumi.Input

	ToActiveDirectoryObjectResponseOutput() ActiveDirectoryObjectResponseOutput
	ToActiveDirectoryObjectResponseOutputWithContext(context.Context) ActiveDirectoryObjectResponseOutput
}

type ActiveDirectoryObjectResponseArgs struct {
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (ActiveDirectoryObjectResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryObjectResponse)(nil)).Elem()
}

func (i ActiveDirectoryObjectResponseArgs) ToActiveDirectoryObjectResponseOutput() ActiveDirectoryObjectResponseOutput {
	return i.ToActiveDirectoryObjectResponseOutputWithContext(context.Background())
}

func (i ActiveDirectoryObjectResponseArgs) ToActiveDirectoryObjectResponseOutputWithContext(ctx context.Context) ActiveDirectoryObjectResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryObjectResponseOutput)
}

func (i ActiveDirectoryObjectResponseArgs) ToActiveDirectoryObjectResponsePtrOutput() ActiveDirectoryObjectResponsePtrOutput {
	return i.ToActiveDirectoryObjectResponsePtrOutputWithContext(context.Background())
}

func (i ActiveDirectoryObjectResponseArgs) ToActiveDirectoryObjectResponsePtrOutputWithContext(ctx context.Context) ActiveDirectoryObjectResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryObjectResponseOutput).ToActiveDirectoryObjectResponsePtrOutputWithContext(ctx)
}









type ActiveDirectoryObjectResponsePtrInput interface {
	pulumi.Input

	ToActiveDirectoryObjectResponsePtrOutput() ActiveDirectoryObjectResponsePtrOutput
	ToActiveDirectoryObjectResponsePtrOutputWithContext(context.Context) ActiveDirectoryObjectResponsePtrOutput
}

type activeDirectoryObjectResponsePtrType ActiveDirectoryObjectResponseArgs

func ActiveDirectoryObjectResponsePtr(v *ActiveDirectoryObjectResponseArgs) ActiveDirectoryObjectResponsePtrInput {
	return (*activeDirectoryObjectResponsePtrType)(v)
}

func (*activeDirectoryObjectResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryObjectResponse)(nil)).Elem()
}

func (i *activeDirectoryObjectResponsePtrType) ToActiveDirectoryObjectResponsePtrOutput() ActiveDirectoryObjectResponsePtrOutput {
	return i.ToActiveDirectoryObjectResponsePtrOutputWithContext(context.Background())
}

func (i *activeDirectoryObjectResponsePtrType) ToActiveDirectoryObjectResponsePtrOutputWithContext(ctx context.Context) ActiveDirectoryObjectResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryObjectResponsePtrOutput)
}

type ActiveDirectoryObjectResponseOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryObjectResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryObjectResponse)(nil)).Elem()
}

func (o ActiveDirectoryObjectResponseOutput) ToActiveDirectoryObjectResponseOutput() ActiveDirectoryObjectResponseOutput {
	return o
}

func (o ActiveDirectoryObjectResponseOutput) ToActiveDirectoryObjectResponseOutputWithContext(ctx context.Context) ActiveDirectoryObjectResponseOutput {
	return o
}

func (o ActiveDirectoryObjectResponseOutput) ToActiveDirectoryObjectResponsePtrOutput() ActiveDirectoryObjectResponsePtrOutput {
	return o.ToActiveDirectoryObjectResponsePtrOutputWithContext(context.Background())
}

func (o ActiveDirectoryObjectResponseOutput) ToActiveDirectoryObjectResponsePtrOutputWithContext(ctx context.Context) ActiveDirectoryObjectResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActiveDirectoryObjectResponse) *ActiveDirectoryObjectResponse {
		return &v
	}).(ActiveDirectoryObjectResponsePtrOutput)
}

func (o ActiveDirectoryObjectResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryObjectResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryObjectResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryObjectResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ActiveDirectoryObjectResponsePtrOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryObjectResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryObjectResponse)(nil)).Elem()
}

func (o ActiveDirectoryObjectResponsePtrOutput) ToActiveDirectoryObjectResponsePtrOutput() ActiveDirectoryObjectResponsePtrOutput {
	return o
}

func (o ActiveDirectoryObjectResponsePtrOutput) ToActiveDirectoryObjectResponsePtrOutputWithContext(ctx context.Context) ActiveDirectoryObjectResponsePtrOutput {
	return o
}

func (o ActiveDirectoryObjectResponsePtrOutput) Elem() ActiveDirectoryObjectResponseOutput {
	return o.ApplyT(func(v *ActiveDirectoryObjectResponse) ActiveDirectoryObjectResponse {
		if v != nil {
			return *v
		}
		var ret ActiveDirectoryObjectResponse
		return ret
	}).(ActiveDirectoryObjectResponseOutput)
}

func (o ActiveDirectoryObjectResponsePtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryObjectResponse) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryObjectResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryObjectResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
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

type TokenCertificate struct {
	EncodedPemCertificate *string `pulumi:"encodedPemCertificate"`
	Expiry                *string `pulumi:"expiry"`
	Name                  *string `pulumi:"name"`
	Thumbprint            *string `pulumi:"thumbprint"`
}





type TokenCertificateInput interface {
	pulumi.Input

	ToTokenCertificateOutput() TokenCertificateOutput
	ToTokenCertificateOutputWithContext(context.Context) TokenCertificateOutput
}

type TokenCertificateArgs struct {
	EncodedPemCertificate pulumi.StringPtrInput `pulumi:"encodedPemCertificate"`
	Expiry                pulumi.StringPtrInput `pulumi:"expiry"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint            pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (TokenCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificate)(nil)).Elem()
}

func (i TokenCertificateArgs) ToTokenCertificateOutput() TokenCertificateOutput {
	return i.ToTokenCertificateOutputWithContext(context.Background())
}

func (i TokenCertificateArgs) ToTokenCertificateOutputWithContext(ctx context.Context) TokenCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCertificateOutput)
}





type TokenCertificateArrayInput interface {
	pulumi.Input

	ToTokenCertificateArrayOutput() TokenCertificateArrayOutput
	ToTokenCertificateArrayOutputWithContext(context.Context) TokenCertificateArrayOutput
}

type TokenCertificateArray []TokenCertificateInput

func (TokenCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenCertificate)(nil)).Elem()
}

func (i TokenCertificateArray) ToTokenCertificateArrayOutput() TokenCertificateArrayOutput {
	return i.ToTokenCertificateArrayOutputWithContext(context.Background())
}

func (i TokenCertificateArray) ToTokenCertificateArrayOutputWithContext(ctx context.Context) TokenCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCertificateArrayOutput)
}

type TokenCertificateOutput struct{ *pulumi.OutputState }

func (TokenCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificate)(nil)).Elem()
}

func (o TokenCertificateOutput) ToTokenCertificateOutput() TokenCertificateOutput {
	return o
}

func (o TokenCertificateOutput) ToTokenCertificateOutputWithContext(ctx context.Context) TokenCertificateOutput {
	return o
}

func (o TokenCertificateOutput) EncodedPemCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificate) *string { return v.EncodedPemCertificate }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificate) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type TokenCertificateArrayOutput struct{ *pulumi.OutputState }

func (TokenCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenCertificate)(nil)).Elem()
}

func (o TokenCertificateArrayOutput) ToTokenCertificateArrayOutput() TokenCertificateArrayOutput {
	return o
}

func (o TokenCertificateArrayOutput) ToTokenCertificateArrayOutputWithContext(ctx context.Context) TokenCertificateArrayOutput {
	return o
}

func (o TokenCertificateArrayOutput) Index(i pulumi.IntInput) TokenCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenCertificate {
		return vs[0].([]TokenCertificate)[vs[1].(int)]
	}).(TokenCertificateOutput)
}

type TokenCertificateResponse struct {
	EncodedPemCertificate *string `pulumi:"encodedPemCertificate"`
	Expiry                *string `pulumi:"expiry"`
	Name                  *string `pulumi:"name"`
	Thumbprint            *string `pulumi:"thumbprint"`
}





type TokenCertificateResponseInput interface {
	pulumi.Input

	ToTokenCertificateResponseOutput() TokenCertificateResponseOutput
	ToTokenCertificateResponseOutputWithContext(context.Context) TokenCertificateResponseOutput
}

type TokenCertificateResponseArgs struct {
	EncodedPemCertificate pulumi.StringPtrInput `pulumi:"encodedPemCertificate"`
	Expiry                pulumi.StringPtrInput `pulumi:"expiry"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint            pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (TokenCertificateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificateResponse)(nil)).Elem()
}

func (i TokenCertificateResponseArgs) ToTokenCertificateResponseOutput() TokenCertificateResponseOutput {
	return i.ToTokenCertificateResponseOutputWithContext(context.Background())
}

func (i TokenCertificateResponseArgs) ToTokenCertificateResponseOutputWithContext(ctx context.Context) TokenCertificateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCertificateResponseOutput)
}





type TokenCertificateResponseArrayInput interface {
	pulumi.Input

	ToTokenCertificateResponseArrayOutput() TokenCertificateResponseArrayOutput
	ToTokenCertificateResponseArrayOutputWithContext(context.Context) TokenCertificateResponseArrayOutput
}

type TokenCertificateResponseArray []TokenCertificateResponseInput

func (TokenCertificateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenCertificateResponse)(nil)).Elem()
}

func (i TokenCertificateResponseArray) ToTokenCertificateResponseArrayOutput() TokenCertificateResponseArrayOutput {
	return i.ToTokenCertificateResponseArrayOutputWithContext(context.Background())
}

func (i TokenCertificateResponseArray) ToTokenCertificateResponseArrayOutputWithContext(ctx context.Context) TokenCertificateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCertificateResponseArrayOutput)
}

type TokenCertificateResponseOutput struct{ *pulumi.OutputState }

func (TokenCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificateResponse)(nil)).Elem()
}

func (o TokenCertificateResponseOutput) ToTokenCertificateResponseOutput() TokenCertificateResponseOutput {
	return o
}

func (o TokenCertificateResponseOutput) ToTokenCertificateResponseOutputWithContext(ctx context.Context) TokenCertificateResponseOutput {
	return o
}

func (o TokenCertificateResponseOutput) EncodedPemCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificateResponse) *string { return v.EncodedPemCertificate }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateResponseOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificateResponse) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type TokenCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (TokenCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenCertificateResponse)(nil)).Elem()
}

func (o TokenCertificateResponseArrayOutput) ToTokenCertificateResponseArrayOutput() TokenCertificateResponseArrayOutput {
	return o
}

func (o TokenCertificateResponseArrayOutput) ToTokenCertificateResponseArrayOutputWithContext(ctx context.Context) TokenCertificateResponseArrayOutput {
	return o
}

func (o TokenCertificateResponseArrayOutput) Index(i pulumi.IntInput) TokenCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenCertificateResponse {
		return vs[0].([]TokenCertificateResponse)[vs[1].(int)]
	}).(TokenCertificateResponseOutput)
}

type TokenCredentialsProperties struct {
	ActiveDirectoryObject *ActiveDirectoryObject `pulumi:"activeDirectoryObject"`
	Certificates          []TokenCertificate     `pulumi:"certificates"`
	Passwords             []TokenPassword        `pulumi:"passwords"`
}





type TokenCredentialsPropertiesInput interface {
	pulumi.Input

	ToTokenCredentialsPropertiesOutput() TokenCredentialsPropertiesOutput
	ToTokenCredentialsPropertiesOutputWithContext(context.Context) TokenCredentialsPropertiesOutput
}

type TokenCredentialsPropertiesArgs struct {
	ActiveDirectoryObject ActiveDirectoryObjectPtrInput `pulumi:"activeDirectoryObject"`
	Certificates          TokenCertificateArrayInput    `pulumi:"certificates"`
	Passwords             TokenPasswordArrayInput       `pulumi:"passwords"`
}

func (TokenCredentialsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCredentialsProperties)(nil)).Elem()
}

func (i TokenCredentialsPropertiesArgs) ToTokenCredentialsPropertiesOutput() TokenCredentialsPropertiesOutput {
	return i.ToTokenCredentialsPropertiesOutputWithContext(context.Background())
}

func (i TokenCredentialsPropertiesArgs) ToTokenCredentialsPropertiesOutputWithContext(ctx context.Context) TokenCredentialsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesOutput)
}

func (i TokenCredentialsPropertiesArgs) ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput {
	return i.ToTokenCredentialsPropertiesPtrOutputWithContext(context.Background())
}

func (i TokenCredentialsPropertiesArgs) ToTokenCredentialsPropertiesPtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesOutput).ToTokenCredentialsPropertiesPtrOutputWithContext(ctx)
}









type TokenCredentialsPropertiesPtrInput interface {
	pulumi.Input

	ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput
	ToTokenCredentialsPropertiesPtrOutputWithContext(context.Context) TokenCredentialsPropertiesPtrOutput
}

type tokenCredentialsPropertiesPtrType TokenCredentialsPropertiesArgs

func TokenCredentialsPropertiesPtr(v *TokenCredentialsPropertiesArgs) TokenCredentialsPropertiesPtrInput {
	return (*tokenCredentialsPropertiesPtrType)(v)
}

func (*tokenCredentialsPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCredentialsProperties)(nil)).Elem()
}

func (i *tokenCredentialsPropertiesPtrType) ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput {
	return i.ToTokenCredentialsPropertiesPtrOutputWithContext(context.Background())
}

func (i *tokenCredentialsPropertiesPtrType) ToTokenCredentialsPropertiesPtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesPtrOutput)
}

type TokenCredentialsPropertiesOutput struct{ *pulumi.OutputState }

func (TokenCredentialsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCredentialsProperties)(nil)).Elem()
}

func (o TokenCredentialsPropertiesOutput) ToTokenCredentialsPropertiesOutput() TokenCredentialsPropertiesOutput {
	return o
}

func (o TokenCredentialsPropertiesOutput) ToTokenCredentialsPropertiesOutputWithContext(ctx context.Context) TokenCredentialsPropertiesOutput {
	return o
}

func (o TokenCredentialsPropertiesOutput) ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput {
	return o.ToTokenCredentialsPropertiesPtrOutputWithContext(context.Background())
}

func (o TokenCredentialsPropertiesOutput) ToTokenCredentialsPropertiesPtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenCredentialsProperties) *TokenCredentialsProperties {
		return &v
	}).(TokenCredentialsPropertiesPtrOutput)
}

func (o TokenCredentialsPropertiesOutput) ActiveDirectoryObject() ActiveDirectoryObjectPtrOutput {
	return o.ApplyT(func(v TokenCredentialsProperties) *ActiveDirectoryObject { return v.ActiveDirectoryObject }).(ActiveDirectoryObjectPtrOutput)
}

func (o TokenCredentialsPropertiesOutput) Certificates() TokenCertificateArrayOutput {
	return o.ApplyT(func(v TokenCredentialsProperties) []TokenCertificate { return v.Certificates }).(TokenCertificateArrayOutput)
}

func (o TokenCredentialsPropertiesOutput) Passwords() TokenPasswordArrayOutput {
	return o.ApplyT(func(v TokenCredentialsProperties) []TokenPassword { return v.Passwords }).(TokenPasswordArrayOutput)
}

type TokenCredentialsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TokenCredentialsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCredentialsProperties)(nil)).Elem()
}

func (o TokenCredentialsPropertiesPtrOutput) ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput {
	return o
}

func (o TokenCredentialsPropertiesPtrOutput) ToTokenCredentialsPropertiesPtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesPtrOutput {
	return o
}

func (o TokenCredentialsPropertiesPtrOutput) Elem() TokenCredentialsPropertiesOutput {
	return o.ApplyT(func(v *TokenCredentialsProperties) TokenCredentialsProperties {
		if v != nil {
			return *v
		}
		var ret TokenCredentialsProperties
		return ret
	}).(TokenCredentialsPropertiesOutput)
}

func (o TokenCredentialsPropertiesPtrOutput) ActiveDirectoryObject() ActiveDirectoryObjectPtrOutput {
	return o.ApplyT(func(v *TokenCredentialsProperties) *ActiveDirectoryObject {
		if v == nil {
			return nil
		}
		return v.ActiveDirectoryObject
	}).(ActiveDirectoryObjectPtrOutput)
}

func (o TokenCredentialsPropertiesPtrOutput) Certificates() TokenCertificateArrayOutput {
	return o.ApplyT(func(v *TokenCredentialsProperties) []TokenCertificate {
		if v == nil {
			return nil
		}
		return v.Certificates
	}).(TokenCertificateArrayOutput)
}

func (o TokenCredentialsPropertiesPtrOutput) Passwords() TokenPasswordArrayOutput {
	return o.ApplyT(func(v *TokenCredentialsProperties) []TokenPassword {
		if v == nil {
			return nil
		}
		return v.Passwords
	}).(TokenPasswordArrayOutput)
}

type TokenCredentialsPropertiesResponse struct {
	ActiveDirectoryObject *ActiveDirectoryObjectResponse `pulumi:"activeDirectoryObject"`
	Certificates          []TokenCertificateResponse     `pulumi:"certificates"`
	Passwords             []TokenPasswordResponse        `pulumi:"passwords"`
}





type TokenCredentialsPropertiesResponseInput interface {
	pulumi.Input

	ToTokenCredentialsPropertiesResponseOutput() TokenCredentialsPropertiesResponseOutput
	ToTokenCredentialsPropertiesResponseOutputWithContext(context.Context) TokenCredentialsPropertiesResponseOutput
}

type TokenCredentialsPropertiesResponseArgs struct {
	ActiveDirectoryObject ActiveDirectoryObjectResponsePtrInput `pulumi:"activeDirectoryObject"`
	Certificates          TokenCertificateResponseArrayInput    `pulumi:"certificates"`
	Passwords             TokenPasswordResponseArrayInput       `pulumi:"passwords"`
}

func (TokenCredentialsPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCredentialsPropertiesResponse)(nil)).Elem()
}

func (i TokenCredentialsPropertiesResponseArgs) ToTokenCredentialsPropertiesResponseOutput() TokenCredentialsPropertiesResponseOutput {
	return i.ToTokenCredentialsPropertiesResponseOutputWithContext(context.Background())
}

func (i TokenCredentialsPropertiesResponseArgs) ToTokenCredentialsPropertiesResponseOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesResponseOutput)
}

func (i TokenCredentialsPropertiesResponseArgs) ToTokenCredentialsPropertiesResponsePtrOutput() TokenCredentialsPropertiesResponsePtrOutput {
	return i.ToTokenCredentialsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i TokenCredentialsPropertiesResponseArgs) ToTokenCredentialsPropertiesResponsePtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesResponseOutput).ToTokenCredentialsPropertiesResponsePtrOutputWithContext(ctx)
}









type TokenCredentialsPropertiesResponsePtrInput interface {
	pulumi.Input

	ToTokenCredentialsPropertiesResponsePtrOutput() TokenCredentialsPropertiesResponsePtrOutput
	ToTokenCredentialsPropertiesResponsePtrOutputWithContext(context.Context) TokenCredentialsPropertiesResponsePtrOutput
}

type tokenCredentialsPropertiesResponsePtrType TokenCredentialsPropertiesResponseArgs

func TokenCredentialsPropertiesResponsePtr(v *TokenCredentialsPropertiesResponseArgs) TokenCredentialsPropertiesResponsePtrInput {
	return (*tokenCredentialsPropertiesResponsePtrType)(v)
}

func (*tokenCredentialsPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCredentialsPropertiesResponse)(nil)).Elem()
}

func (i *tokenCredentialsPropertiesResponsePtrType) ToTokenCredentialsPropertiesResponsePtrOutput() TokenCredentialsPropertiesResponsePtrOutput {
	return i.ToTokenCredentialsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *tokenCredentialsPropertiesResponsePtrType) ToTokenCredentialsPropertiesResponsePtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesResponsePtrOutput)
}

type TokenCredentialsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TokenCredentialsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCredentialsPropertiesResponse)(nil)).Elem()
}

func (o TokenCredentialsPropertiesResponseOutput) ToTokenCredentialsPropertiesResponseOutput() TokenCredentialsPropertiesResponseOutput {
	return o
}

func (o TokenCredentialsPropertiesResponseOutput) ToTokenCredentialsPropertiesResponseOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponseOutput {
	return o
}

func (o TokenCredentialsPropertiesResponseOutput) ToTokenCredentialsPropertiesResponsePtrOutput() TokenCredentialsPropertiesResponsePtrOutput {
	return o.ToTokenCredentialsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o TokenCredentialsPropertiesResponseOutput) ToTokenCredentialsPropertiesResponsePtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenCredentialsPropertiesResponse) *TokenCredentialsPropertiesResponse {
		return &v
	}).(TokenCredentialsPropertiesResponsePtrOutput)
}

func (o TokenCredentialsPropertiesResponseOutput) ActiveDirectoryObject() ActiveDirectoryObjectResponsePtrOutput {
	return o.ApplyT(func(v TokenCredentialsPropertiesResponse) *ActiveDirectoryObjectResponse {
		return v.ActiveDirectoryObject
	}).(ActiveDirectoryObjectResponsePtrOutput)
}

func (o TokenCredentialsPropertiesResponseOutput) Certificates() TokenCertificateResponseArrayOutput {
	return o.ApplyT(func(v TokenCredentialsPropertiesResponse) []TokenCertificateResponse { return v.Certificates }).(TokenCertificateResponseArrayOutput)
}

func (o TokenCredentialsPropertiesResponseOutput) Passwords() TokenPasswordResponseArrayOutput {
	return o.ApplyT(func(v TokenCredentialsPropertiesResponse) []TokenPasswordResponse { return v.Passwords }).(TokenPasswordResponseArrayOutput)
}

type TokenCredentialsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TokenCredentialsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCredentialsPropertiesResponse)(nil)).Elem()
}

func (o TokenCredentialsPropertiesResponsePtrOutput) ToTokenCredentialsPropertiesResponsePtrOutput() TokenCredentialsPropertiesResponsePtrOutput {
	return o
}

func (o TokenCredentialsPropertiesResponsePtrOutput) ToTokenCredentialsPropertiesResponsePtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponsePtrOutput {
	return o
}

func (o TokenCredentialsPropertiesResponsePtrOutput) Elem() TokenCredentialsPropertiesResponseOutput {
	return o.ApplyT(func(v *TokenCredentialsPropertiesResponse) TokenCredentialsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TokenCredentialsPropertiesResponse
		return ret
	}).(TokenCredentialsPropertiesResponseOutput)
}

func (o TokenCredentialsPropertiesResponsePtrOutput) ActiveDirectoryObject() ActiveDirectoryObjectResponsePtrOutput {
	return o.ApplyT(func(v *TokenCredentialsPropertiesResponse) *ActiveDirectoryObjectResponse {
		if v == nil {
			return nil
		}
		return v.ActiveDirectoryObject
	}).(ActiveDirectoryObjectResponsePtrOutput)
}

func (o TokenCredentialsPropertiesResponsePtrOutput) Certificates() TokenCertificateResponseArrayOutput {
	return o.ApplyT(func(v *TokenCredentialsPropertiesResponse) []TokenCertificateResponse {
		if v == nil {
			return nil
		}
		return v.Certificates
	}).(TokenCertificateResponseArrayOutput)
}

func (o TokenCredentialsPropertiesResponsePtrOutput) Passwords() TokenPasswordResponseArrayOutput {
	return o.ApplyT(func(v *TokenCredentialsPropertiesResponse) []TokenPasswordResponse {
		if v == nil {
			return nil
		}
		return v.Passwords
	}).(TokenPasswordResponseArrayOutput)
}

type TokenPassword struct {
	CreationTime *string `pulumi:"creationTime"`
	Expiry       *string `pulumi:"expiry"`
	Name         *string `pulumi:"name"`
}





type TokenPasswordInput interface {
	pulumi.Input

	ToTokenPasswordOutput() TokenPasswordOutput
	ToTokenPasswordOutputWithContext(context.Context) TokenPasswordOutput
}

type TokenPasswordArgs struct {
	CreationTime pulumi.StringPtrInput `pulumi:"creationTime"`
	Expiry       pulumi.StringPtrInput `pulumi:"expiry"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
}

func (TokenPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPassword)(nil)).Elem()
}

func (i TokenPasswordArgs) ToTokenPasswordOutput() TokenPasswordOutput {
	return i.ToTokenPasswordOutputWithContext(context.Background())
}

func (i TokenPasswordArgs) ToTokenPasswordOutputWithContext(ctx context.Context) TokenPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenPasswordOutput)
}





type TokenPasswordArrayInput interface {
	pulumi.Input

	ToTokenPasswordArrayOutput() TokenPasswordArrayOutput
	ToTokenPasswordArrayOutputWithContext(context.Context) TokenPasswordArrayOutput
}

type TokenPasswordArray []TokenPasswordInput

func (TokenPasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenPassword)(nil)).Elem()
}

func (i TokenPasswordArray) ToTokenPasswordArrayOutput() TokenPasswordArrayOutput {
	return i.ToTokenPasswordArrayOutputWithContext(context.Background())
}

func (i TokenPasswordArray) ToTokenPasswordArrayOutputWithContext(ctx context.Context) TokenPasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenPasswordArrayOutput)
}

type TokenPasswordOutput struct{ *pulumi.OutputState }

func (TokenPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPassword)(nil)).Elem()
}

func (o TokenPasswordOutput) ToTokenPasswordOutput() TokenPasswordOutput {
	return o
}

func (o TokenPasswordOutput) ToTokenPasswordOutputWithContext(ctx context.Context) TokenPasswordOutput {
	return o
}

func (o TokenPasswordOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPassword) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPassword) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPassword) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type TokenPasswordArrayOutput struct{ *pulumi.OutputState }

func (TokenPasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenPassword)(nil)).Elem()
}

func (o TokenPasswordArrayOutput) ToTokenPasswordArrayOutput() TokenPasswordArrayOutput {
	return o
}

func (o TokenPasswordArrayOutput) ToTokenPasswordArrayOutputWithContext(ctx context.Context) TokenPasswordArrayOutput {
	return o
}

func (o TokenPasswordArrayOutput) Index(i pulumi.IntInput) TokenPasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenPassword {
		return vs[0].([]TokenPassword)[vs[1].(int)]
	}).(TokenPasswordOutput)
}

type TokenPasswordResponse struct {
	CreationTime *string `pulumi:"creationTime"`
	Expiry       *string `pulumi:"expiry"`
	Name         *string `pulumi:"name"`
	Value        string  `pulumi:"value"`
}





type TokenPasswordResponseInput interface {
	pulumi.Input

	ToTokenPasswordResponseOutput() TokenPasswordResponseOutput
	ToTokenPasswordResponseOutputWithContext(context.Context) TokenPasswordResponseOutput
}

type TokenPasswordResponseArgs struct {
	CreationTime pulumi.StringPtrInput `pulumi:"creationTime"`
	Expiry       pulumi.StringPtrInput `pulumi:"expiry"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	Value        pulumi.StringInput    `pulumi:"value"`
}

func (TokenPasswordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPasswordResponse)(nil)).Elem()
}

func (i TokenPasswordResponseArgs) ToTokenPasswordResponseOutput() TokenPasswordResponseOutput {
	return i.ToTokenPasswordResponseOutputWithContext(context.Background())
}

func (i TokenPasswordResponseArgs) ToTokenPasswordResponseOutputWithContext(ctx context.Context) TokenPasswordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenPasswordResponseOutput)
}





type TokenPasswordResponseArrayInput interface {
	pulumi.Input

	ToTokenPasswordResponseArrayOutput() TokenPasswordResponseArrayOutput
	ToTokenPasswordResponseArrayOutputWithContext(context.Context) TokenPasswordResponseArrayOutput
}

type TokenPasswordResponseArray []TokenPasswordResponseInput

func (TokenPasswordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenPasswordResponse)(nil)).Elem()
}

func (i TokenPasswordResponseArray) ToTokenPasswordResponseArrayOutput() TokenPasswordResponseArrayOutput {
	return i.ToTokenPasswordResponseArrayOutputWithContext(context.Background())
}

func (i TokenPasswordResponseArray) ToTokenPasswordResponseArrayOutputWithContext(ctx context.Context) TokenPasswordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenPasswordResponseArrayOutput)
}

type TokenPasswordResponseOutput struct{ *pulumi.OutputState }

func (TokenPasswordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPasswordResponse)(nil)).Elem()
}

func (o TokenPasswordResponseOutput) ToTokenPasswordResponseOutput() TokenPasswordResponseOutput {
	return o
}

func (o TokenPasswordResponseOutput) ToTokenPasswordResponseOutputWithContext(ctx context.Context) TokenPasswordResponseOutput {
	return o
}

func (o TokenPasswordResponseOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPasswordResponse) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordResponseOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPasswordResponse) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPasswordResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TokenPasswordResponse) string { return v.Value }).(pulumi.StringOutput)
}

type TokenPasswordResponseArrayOutput struct{ *pulumi.OutputState }

func (TokenPasswordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenPasswordResponse)(nil)).Elem()
}

func (o TokenPasswordResponseArrayOutput) ToTokenPasswordResponseArrayOutput() TokenPasswordResponseArrayOutput {
	return o
}

func (o TokenPasswordResponseArrayOutput) ToTokenPasswordResponseArrayOutputWithContext(ctx context.Context) TokenPasswordResponseArrayOutput {
	return o
}

func (o TokenPasswordResponseArrayOutput) Index(i pulumi.IntInput) TokenPasswordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenPasswordResponse {
		return vs[0].([]TokenPasswordResponse)[vs[1].(int)]
	}).(TokenPasswordResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ActiveDirectoryObjectOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryObjectPtrOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryObjectResponseOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryObjectResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TokenCertificateOutput{})
	pulumi.RegisterOutputType(TokenCertificateArrayOutput{})
	pulumi.RegisterOutputType(TokenCertificateResponseOutput{})
	pulumi.RegisterOutputType(TokenCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(TokenCredentialsPropertiesOutput{})
	pulumi.RegisterOutputType(TokenCredentialsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TokenCredentialsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TokenCredentialsPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TokenPasswordOutput{})
	pulumi.RegisterOutputType(TokenPasswordArrayOutput{})
	pulumi.RegisterOutputType(TokenPasswordResponseOutput{})
	pulumi.RegisterOutputType(TokenPasswordResponseArrayOutput{})
}

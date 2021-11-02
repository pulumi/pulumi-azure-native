


package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdmCredential struct {
	AuthTokenUrl *string `pulumi:"authTokenUrl"`
	ClientId     *string `pulumi:"clientId"`
	ClientSecret *string `pulumi:"clientSecret"`
}





type AdmCredentialInput interface {
	pulumi.Input

	ToAdmCredentialOutput() AdmCredentialOutput
	ToAdmCredentialOutputWithContext(context.Context) AdmCredentialOutput
}

type AdmCredentialArgs struct {
	AuthTokenUrl pulumi.StringPtrInput `pulumi:"authTokenUrl"`
	ClientId     pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecret pulumi.StringPtrInput `pulumi:"clientSecret"`
}

func (AdmCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdmCredential)(nil)).Elem()
}

func (i AdmCredentialArgs) ToAdmCredentialOutput() AdmCredentialOutput {
	return i.ToAdmCredentialOutputWithContext(context.Background())
}

func (i AdmCredentialArgs) ToAdmCredentialOutputWithContext(ctx context.Context) AdmCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmCredentialOutput)
}

func (i AdmCredentialArgs) ToAdmCredentialPtrOutput() AdmCredentialPtrOutput {
	return i.ToAdmCredentialPtrOutputWithContext(context.Background())
}

func (i AdmCredentialArgs) ToAdmCredentialPtrOutputWithContext(ctx context.Context) AdmCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmCredentialOutput).ToAdmCredentialPtrOutputWithContext(ctx)
}









type AdmCredentialPtrInput interface {
	pulumi.Input

	ToAdmCredentialPtrOutput() AdmCredentialPtrOutput
	ToAdmCredentialPtrOutputWithContext(context.Context) AdmCredentialPtrOutput
}

type admCredentialPtrType AdmCredentialArgs

func AdmCredentialPtr(v *AdmCredentialArgs) AdmCredentialPtrInput {
	return (*admCredentialPtrType)(v)
}

func (*admCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AdmCredential)(nil)).Elem()
}

func (i *admCredentialPtrType) ToAdmCredentialPtrOutput() AdmCredentialPtrOutput {
	return i.ToAdmCredentialPtrOutputWithContext(context.Background())
}

func (i *admCredentialPtrType) ToAdmCredentialPtrOutputWithContext(ctx context.Context) AdmCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmCredentialPtrOutput)
}

type AdmCredentialOutput struct{ *pulumi.OutputState }

func (AdmCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdmCredential)(nil)).Elem()
}

func (o AdmCredentialOutput) ToAdmCredentialOutput() AdmCredentialOutput {
	return o
}

func (o AdmCredentialOutput) ToAdmCredentialOutputWithContext(ctx context.Context) AdmCredentialOutput {
	return o
}

func (o AdmCredentialOutput) ToAdmCredentialPtrOutput() AdmCredentialPtrOutput {
	return o.ToAdmCredentialPtrOutputWithContext(context.Background())
}

func (o AdmCredentialOutput) ToAdmCredentialPtrOutputWithContext(ctx context.Context) AdmCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdmCredential) *AdmCredential {
		return &v
	}).(AdmCredentialPtrOutput)
}

func (o AdmCredentialOutput) AuthTokenUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdmCredential) *string { return v.AuthTokenUrl }).(pulumi.StringPtrOutput)
}

func (o AdmCredentialOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdmCredential) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AdmCredentialOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdmCredential) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

type AdmCredentialPtrOutput struct{ *pulumi.OutputState }

func (AdmCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdmCredential)(nil)).Elem()
}

func (o AdmCredentialPtrOutput) ToAdmCredentialPtrOutput() AdmCredentialPtrOutput {
	return o
}

func (o AdmCredentialPtrOutput) ToAdmCredentialPtrOutputWithContext(ctx context.Context) AdmCredentialPtrOutput {
	return o
}

func (o AdmCredentialPtrOutput) Elem() AdmCredentialOutput {
	return o.ApplyT(func(v *AdmCredential) AdmCredential {
		if v != nil {
			return *v
		}
		var ret AdmCredential
		return ret
	}).(AdmCredentialOutput)
}

func (o AdmCredentialPtrOutput) AuthTokenUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdmCredential) *string {
		if v == nil {
			return nil
		}
		return v.AuthTokenUrl
	}).(pulumi.StringPtrOutput)
}

func (o AdmCredentialPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdmCredential) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AdmCredentialPtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdmCredential) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

type AdmCredentialResponse struct {
	AuthTokenUrl *string `pulumi:"authTokenUrl"`
	ClientId     *string `pulumi:"clientId"`
	ClientSecret *string `pulumi:"clientSecret"`
}





type AdmCredentialResponseInput interface {
	pulumi.Input

	ToAdmCredentialResponseOutput() AdmCredentialResponseOutput
	ToAdmCredentialResponseOutputWithContext(context.Context) AdmCredentialResponseOutput
}

type AdmCredentialResponseArgs struct {
	AuthTokenUrl pulumi.StringPtrInput `pulumi:"authTokenUrl"`
	ClientId     pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecret pulumi.StringPtrInput `pulumi:"clientSecret"`
}

func (AdmCredentialResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdmCredentialResponse)(nil)).Elem()
}

func (i AdmCredentialResponseArgs) ToAdmCredentialResponseOutput() AdmCredentialResponseOutput {
	return i.ToAdmCredentialResponseOutputWithContext(context.Background())
}

func (i AdmCredentialResponseArgs) ToAdmCredentialResponseOutputWithContext(ctx context.Context) AdmCredentialResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmCredentialResponseOutput)
}

func (i AdmCredentialResponseArgs) ToAdmCredentialResponsePtrOutput() AdmCredentialResponsePtrOutput {
	return i.ToAdmCredentialResponsePtrOutputWithContext(context.Background())
}

func (i AdmCredentialResponseArgs) ToAdmCredentialResponsePtrOutputWithContext(ctx context.Context) AdmCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmCredentialResponseOutput).ToAdmCredentialResponsePtrOutputWithContext(ctx)
}









type AdmCredentialResponsePtrInput interface {
	pulumi.Input

	ToAdmCredentialResponsePtrOutput() AdmCredentialResponsePtrOutput
	ToAdmCredentialResponsePtrOutputWithContext(context.Context) AdmCredentialResponsePtrOutput
}

type admCredentialResponsePtrType AdmCredentialResponseArgs

func AdmCredentialResponsePtr(v *AdmCredentialResponseArgs) AdmCredentialResponsePtrInput {
	return (*admCredentialResponsePtrType)(v)
}

func (*admCredentialResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AdmCredentialResponse)(nil)).Elem()
}

func (i *admCredentialResponsePtrType) ToAdmCredentialResponsePtrOutput() AdmCredentialResponsePtrOutput {
	return i.ToAdmCredentialResponsePtrOutputWithContext(context.Background())
}

func (i *admCredentialResponsePtrType) ToAdmCredentialResponsePtrOutputWithContext(ctx context.Context) AdmCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmCredentialResponsePtrOutput)
}

type AdmCredentialResponseOutput struct{ *pulumi.OutputState }

func (AdmCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdmCredentialResponse)(nil)).Elem()
}

func (o AdmCredentialResponseOutput) ToAdmCredentialResponseOutput() AdmCredentialResponseOutput {
	return o
}

func (o AdmCredentialResponseOutput) ToAdmCredentialResponseOutputWithContext(ctx context.Context) AdmCredentialResponseOutput {
	return o
}

func (o AdmCredentialResponseOutput) ToAdmCredentialResponsePtrOutput() AdmCredentialResponsePtrOutput {
	return o.ToAdmCredentialResponsePtrOutputWithContext(context.Background())
}

func (o AdmCredentialResponseOutput) ToAdmCredentialResponsePtrOutputWithContext(ctx context.Context) AdmCredentialResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdmCredentialResponse) *AdmCredentialResponse {
		return &v
	}).(AdmCredentialResponsePtrOutput)
}

func (o AdmCredentialResponseOutput) AuthTokenUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdmCredentialResponse) *string { return v.AuthTokenUrl }).(pulumi.StringPtrOutput)
}

func (o AdmCredentialResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdmCredentialResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AdmCredentialResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdmCredentialResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

type AdmCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (AdmCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdmCredentialResponse)(nil)).Elem()
}

func (o AdmCredentialResponsePtrOutput) ToAdmCredentialResponsePtrOutput() AdmCredentialResponsePtrOutput {
	return o
}

func (o AdmCredentialResponsePtrOutput) ToAdmCredentialResponsePtrOutputWithContext(ctx context.Context) AdmCredentialResponsePtrOutput {
	return o
}

func (o AdmCredentialResponsePtrOutput) Elem() AdmCredentialResponseOutput {
	return o.ApplyT(func(v *AdmCredentialResponse) AdmCredentialResponse {
		if v != nil {
			return *v
		}
		var ret AdmCredentialResponse
		return ret
	}).(AdmCredentialResponseOutput)
}

func (o AdmCredentialResponsePtrOutput) AuthTokenUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdmCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthTokenUrl
	}).(pulumi.StringPtrOutput)
}

func (o AdmCredentialResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdmCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AdmCredentialResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdmCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

type ApnsCredential struct {
	ApnsCertificate *string `pulumi:"apnsCertificate"`
	CertificateKey  *string `pulumi:"certificateKey"`
	Endpoint        *string `pulumi:"endpoint"`
	Thumbprint      *string `pulumi:"thumbprint"`
}





type ApnsCredentialInput interface {
	pulumi.Input

	ToApnsCredentialOutput() ApnsCredentialOutput
	ToApnsCredentialOutputWithContext(context.Context) ApnsCredentialOutput
}

type ApnsCredentialArgs struct {
	ApnsCertificate pulumi.StringPtrInput `pulumi:"apnsCertificate"`
	CertificateKey  pulumi.StringPtrInput `pulumi:"certificateKey"`
	Endpoint        pulumi.StringPtrInput `pulumi:"endpoint"`
	Thumbprint      pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (ApnsCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsCredential)(nil)).Elem()
}

func (i ApnsCredentialArgs) ToApnsCredentialOutput() ApnsCredentialOutput {
	return i.ToApnsCredentialOutputWithContext(context.Background())
}

func (i ApnsCredentialArgs) ToApnsCredentialOutputWithContext(ctx context.Context) ApnsCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsCredentialOutput)
}

func (i ApnsCredentialArgs) ToApnsCredentialPtrOutput() ApnsCredentialPtrOutput {
	return i.ToApnsCredentialPtrOutputWithContext(context.Background())
}

func (i ApnsCredentialArgs) ToApnsCredentialPtrOutputWithContext(ctx context.Context) ApnsCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsCredentialOutput).ToApnsCredentialPtrOutputWithContext(ctx)
}









type ApnsCredentialPtrInput interface {
	pulumi.Input

	ToApnsCredentialPtrOutput() ApnsCredentialPtrOutput
	ToApnsCredentialPtrOutputWithContext(context.Context) ApnsCredentialPtrOutput
}

type apnsCredentialPtrType ApnsCredentialArgs

func ApnsCredentialPtr(v *ApnsCredentialArgs) ApnsCredentialPtrInput {
	return (*apnsCredentialPtrType)(v)
}

func (*apnsCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsCredential)(nil)).Elem()
}

func (i *apnsCredentialPtrType) ToApnsCredentialPtrOutput() ApnsCredentialPtrOutput {
	return i.ToApnsCredentialPtrOutputWithContext(context.Background())
}

func (i *apnsCredentialPtrType) ToApnsCredentialPtrOutputWithContext(ctx context.Context) ApnsCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsCredentialPtrOutput)
}

type ApnsCredentialOutput struct{ *pulumi.OutputState }

func (ApnsCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsCredential)(nil)).Elem()
}

func (o ApnsCredentialOutput) ToApnsCredentialOutput() ApnsCredentialOutput {
	return o
}

func (o ApnsCredentialOutput) ToApnsCredentialOutputWithContext(ctx context.Context) ApnsCredentialOutput {
	return o
}

func (o ApnsCredentialOutput) ToApnsCredentialPtrOutput() ApnsCredentialPtrOutput {
	return o.ToApnsCredentialPtrOutputWithContext(context.Background())
}

func (o ApnsCredentialOutput) ToApnsCredentialPtrOutputWithContext(ctx context.Context) ApnsCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApnsCredential) *ApnsCredential {
		return &v
	}).(ApnsCredentialPtrOutput)
}

func (o ApnsCredentialOutput) ApnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredential) *string { return v.ApnsCertificate }).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredential) *string { return v.CertificateKey }).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredential) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredential) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ApnsCredentialPtrOutput struct{ *pulumi.OutputState }

func (ApnsCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsCredential)(nil)).Elem()
}

func (o ApnsCredentialPtrOutput) ToApnsCredentialPtrOutput() ApnsCredentialPtrOutput {
	return o
}

func (o ApnsCredentialPtrOutput) ToApnsCredentialPtrOutputWithContext(ctx context.Context) ApnsCredentialPtrOutput {
	return o
}

func (o ApnsCredentialPtrOutput) Elem() ApnsCredentialOutput {
	return o.ApplyT(func(v *ApnsCredential) ApnsCredential {
		if v != nil {
			return *v
		}
		var ret ApnsCredential
		return ret
	}).(ApnsCredentialOutput)
}

func (o ApnsCredentialPtrOutput) ApnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredential) *string {
		if v == nil {
			return nil
		}
		return v.ApnsCertificate
	}).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPtrOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredential) *string {
		if v == nil {
			return nil
		}
		return v.CertificateKey
	}).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredential) *string {
		if v == nil {
			return nil
		}
		return v.Endpoint
	}).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredential) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type ApnsCredentialResponse struct {
	ApnsCertificate *string `pulumi:"apnsCertificate"`
	CertificateKey  *string `pulumi:"certificateKey"`
	Endpoint        *string `pulumi:"endpoint"`
	Thumbprint      *string `pulumi:"thumbprint"`
}





type ApnsCredentialResponseInput interface {
	pulumi.Input

	ToApnsCredentialResponseOutput() ApnsCredentialResponseOutput
	ToApnsCredentialResponseOutputWithContext(context.Context) ApnsCredentialResponseOutput
}

type ApnsCredentialResponseArgs struct {
	ApnsCertificate pulumi.StringPtrInput `pulumi:"apnsCertificate"`
	CertificateKey  pulumi.StringPtrInput `pulumi:"certificateKey"`
	Endpoint        pulumi.StringPtrInput `pulumi:"endpoint"`
	Thumbprint      pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (ApnsCredentialResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsCredentialResponse)(nil)).Elem()
}

func (i ApnsCredentialResponseArgs) ToApnsCredentialResponseOutput() ApnsCredentialResponseOutput {
	return i.ToApnsCredentialResponseOutputWithContext(context.Background())
}

func (i ApnsCredentialResponseArgs) ToApnsCredentialResponseOutputWithContext(ctx context.Context) ApnsCredentialResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsCredentialResponseOutput)
}

func (i ApnsCredentialResponseArgs) ToApnsCredentialResponsePtrOutput() ApnsCredentialResponsePtrOutput {
	return i.ToApnsCredentialResponsePtrOutputWithContext(context.Background())
}

func (i ApnsCredentialResponseArgs) ToApnsCredentialResponsePtrOutputWithContext(ctx context.Context) ApnsCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsCredentialResponseOutput).ToApnsCredentialResponsePtrOutputWithContext(ctx)
}









type ApnsCredentialResponsePtrInput interface {
	pulumi.Input

	ToApnsCredentialResponsePtrOutput() ApnsCredentialResponsePtrOutput
	ToApnsCredentialResponsePtrOutputWithContext(context.Context) ApnsCredentialResponsePtrOutput
}

type apnsCredentialResponsePtrType ApnsCredentialResponseArgs

func ApnsCredentialResponsePtr(v *ApnsCredentialResponseArgs) ApnsCredentialResponsePtrInput {
	return (*apnsCredentialResponsePtrType)(v)
}

func (*apnsCredentialResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsCredentialResponse)(nil)).Elem()
}

func (i *apnsCredentialResponsePtrType) ToApnsCredentialResponsePtrOutput() ApnsCredentialResponsePtrOutput {
	return i.ToApnsCredentialResponsePtrOutputWithContext(context.Background())
}

func (i *apnsCredentialResponsePtrType) ToApnsCredentialResponsePtrOutputWithContext(ctx context.Context) ApnsCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsCredentialResponsePtrOutput)
}

type ApnsCredentialResponseOutput struct{ *pulumi.OutputState }

func (ApnsCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsCredentialResponse)(nil)).Elem()
}

func (o ApnsCredentialResponseOutput) ToApnsCredentialResponseOutput() ApnsCredentialResponseOutput {
	return o
}

func (o ApnsCredentialResponseOutput) ToApnsCredentialResponseOutputWithContext(ctx context.Context) ApnsCredentialResponseOutput {
	return o
}

func (o ApnsCredentialResponseOutput) ToApnsCredentialResponsePtrOutput() ApnsCredentialResponsePtrOutput {
	return o.ToApnsCredentialResponsePtrOutputWithContext(context.Background())
}

func (o ApnsCredentialResponseOutput) ToApnsCredentialResponsePtrOutputWithContext(ctx context.Context) ApnsCredentialResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApnsCredentialResponse) *ApnsCredentialResponse {
		return &v
	}).(ApnsCredentialResponsePtrOutput)
}

func (o ApnsCredentialResponseOutput) ApnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredentialResponse) *string { return v.ApnsCertificate }).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialResponseOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredentialResponse) *string { return v.CertificateKey }).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialResponseOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredentialResponse) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredentialResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ApnsCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (ApnsCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsCredentialResponse)(nil)).Elem()
}

func (o ApnsCredentialResponsePtrOutput) ToApnsCredentialResponsePtrOutput() ApnsCredentialResponsePtrOutput {
	return o
}

func (o ApnsCredentialResponsePtrOutput) ToApnsCredentialResponsePtrOutputWithContext(ctx context.Context) ApnsCredentialResponsePtrOutput {
	return o
}

func (o ApnsCredentialResponsePtrOutput) Elem() ApnsCredentialResponseOutput {
	return o.ApplyT(func(v *ApnsCredentialResponse) ApnsCredentialResponse {
		if v != nil {
			return *v
		}
		var ret ApnsCredentialResponse
		return ret
	}).(ApnsCredentialResponseOutput)
}

func (o ApnsCredentialResponsePtrOutput) ApnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApnsCertificate
	}).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialResponsePtrOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.CertificateKey
	}).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialResponsePtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.Endpoint
	}).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type BaiduCredential struct {
	BaiduApiKey    *string `pulumi:"baiduApiKey"`
	BaiduEndPoint  *string `pulumi:"baiduEndPoint"`
	BaiduSecretKey *string `pulumi:"baiduSecretKey"`
}





type BaiduCredentialInput interface {
	pulumi.Input

	ToBaiduCredentialOutput() BaiduCredentialOutput
	ToBaiduCredentialOutputWithContext(context.Context) BaiduCredentialOutput
}

type BaiduCredentialArgs struct {
	BaiduApiKey    pulumi.StringPtrInput `pulumi:"baiduApiKey"`
	BaiduEndPoint  pulumi.StringPtrInput `pulumi:"baiduEndPoint"`
	BaiduSecretKey pulumi.StringPtrInput `pulumi:"baiduSecretKey"`
}

func (BaiduCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BaiduCredential)(nil)).Elem()
}

func (i BaiduCredentialArgs) ToBaiduCredentialOutput() BaiduCredentialOutput {
	return i.ToBaiduCredentialOutputWithContext(context.Background())
}

func (i BaiduCredentialArgs) ToBaiduCredentialOutputWithContext(ctx context.Context) BaiduCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduCredentialOutput)
}

func (i BaiduCredentialArgs) ToBaiduCredentialPtrOutput() BaiduCredentialPtrOutput {
	return i.ToBaiduCredentialPtrOutputWithContext(context.Background())
}

func (i BaiduCredentialArgs) ToBaiduCredentialPtrOutputWithContext(ctx context.Context) BaiduCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduCredentialOutput).ToBaiduCredentialPtrOutputWithContext(ctx)
}









type BaiduCredentialPtrInput interface {
	pulumi.Input

	ToBaiduCredentialPtrOutput() BaiduCredentialPtrOutput
	ToBaiduCredentialPtrOutputWithContext(context.Context) BaiduCredentialPtrOutput
}

type baiduCredentialPtrType BaiduCredentialArgs

func BaiduCredentialPtr(v *BaiduCredentialArgs) BaiduCredentialPtrInput {
	return (*baiduCredentialPtrType)(v)
}

func (*baiduCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BaiduCredential)(nil)).Elem()
}

func (i *baiduCredentialPtrType) ToBaiduCredentialPtrOutput() BaiduCredentialPtrOutput {
	return i.ToBaiduCredentialPtrOutputWithContext(context.Background())
}

func (i *baiduCredentialPtrType) ToBaiduCredentialPtrOutputWithContext(ctx context.Context) BaiduCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduCredentialPtrOutput)
}

type BaiduCredentialOutput struct{ *pulumi.OutputState }

func (BaiduCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BaiduCredential)(nil)).Elem()
}

func (o BaiduCredentialOutput) ToBaiduCredentialOutput() BaiduCredentialOutput {
	return o
}

func (o BaiduCredentialOutput) ToBaiduCredentialOutputWithContext(ctx context.Context) BaiduCredentialOutput {
	return o
}

func (o BaiduCredentialOutput) ToBaiduCredentialPtrOutput() BaiduCredentialPtrOutput {
	return o.ToBaiduCredentialPtrOutputWithContext(context.Background())
}

func (o BaiduCredentialOutput) ToBaiduCredentialPtrOutputWithContext(ctx context.Context) BaiduCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BaiduCredential) *BaiduCredential {
		return &v
	}).(BaiduCredentialPtrOutput)
}

func (o BaiduCredentialOutput) BaiduApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaiduCredential) *string { return v.BaiduApiKey }).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialOutput) BaiduEndPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaiduCredential) *string { return v.BaiduEndPoint }).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialOutput) BaiduSecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaiduCredential) *string { return v.BaiduSecretKey }).(pulumi.StringPtrOutput)
}

type BaiduCredentialPtrOutput struct{ *pulumi.OutputState }

func (BaiduCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BaiduCredential)(nil)).Elem()
}

func (o BaiduCredentialPtrOutput) ToBaiduCredentialPtrOutput() BaiduCredentialPtrOutput {
	return o
}

func (o BaiduCredentialPtrOutput) ToBaiduCredentialPtrOutputWithContext(ctx context.Context) BaiduCredentialPtrOutput {
	return o
}

func (o BaiduCredentialPtrOutput) Elem() BaiduCredentialOutput {
	return o.ApplyT(func(v *BaiduCredential) BaiduCredential {
		if v != nil {
			return *v
		}
		var ret BaiduCredential
		return ret
	}).(BaiduCredentialOutput)
}

func (o BaiduCredentialPtrOutput) BaiduApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaiduCredential) *string {
		if v == nil {
			return nil
		}
		return v.BaiduApiKey
	}).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialPtrOutput) BaiduEndPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaiduCredential) *string {
		if v == nil {
			return nil
		}
		return v.BaiduEndPoint
	}).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialPtrOutput) BaiduSecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaiduCredential) *string {
		if v == nil {
			return nil
		}
		return v.BaiduSecretKey
	}).(pulumi.StringPtrOutput)
}

type BaiduCredentialResponse struct {
	BaiduApiKey    *string `pulumi:"baiduApiKey"`
	BaiduEndPoint  *string `pulumi:"baiduEndPoint"`
	BaiduSecretKey *string `pulumi:"baiduSecretKey"`
}





type BaiduCredentialResponseInput interface {
	pulumi.Input

	ToBaiduCredentialResponseOutput() BaiduCredentialResponseOutput
	ToBaiduCredentialResponseOutputWithContext(context.Context) BaiduCredentialResponseOutput
}

type BaiduCredentialResponseArgs struct {
	BaiduApiKey    pulumi.StringPtrInput `pulumi:"baiduApiKey"`
	BaiduEndPoint  pulumi.StringPtrInput `pulumi:"baiduEndPoint"`
	BaiduSecretKey pulumi.StringPtrInput `pulumi:"baiduSecretKey"`
}

func (BaiduCredentialResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BaiduCredentialResponse)(nil)).Elem()
}

func (i BaiduCredentialResponseArgs) ToBaiduCredentialResponseOutput() BaiduCredentialResponseOutput {
	return i.ToBaiduCredentialResponseOutputWithContext(context.Background())
}

func (i BaiduCredentialResponseArgs) ToBaiduCredentialResponseOutputWithContext(ctx context.Context) BaiduCredentialResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduCredentialResponseOutput)
}

func (i BaiduCredentialResponseArgs) ToBaiduCredentialResponsePtrOutput() BaiduCredentialResponsePtrOutput {
	return i.ToBaiduCredentialResponsePtrOutputWithContext(context.Background())
}

func (i BaiduCredentialResponseArgs) ToBaiduCredentialResponsePtrOutputWithContext(ctx context.Context) BaiduCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduCredentialResponseOutput).ToBaiduCredentialResponsePtrOutputWithContext(ctx)
}









type BaiduCredentialResponsePtrInput interface {
	pulumi.Input

	ToBaiduCredentialResponsePtrOutput() BaiduCredentialResponsePtrOutput
	ToBaiduCredentialResponsePtrOutputWithContext(context.Context) BaiduCredentialResponsePtrOutput
}

type baiduCredentialResponsePtrType BaiduCredentialResponseArgs

func BaiduCredentialResponsePtr(v *BaiduCredentialResponseArgs) BaiduCredentialResponsePtrInput {
	return (*baiduCredentialResponsePtrType)(v)
}

func (*baiduCredentialResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BaiduCredentialResponse)(nil)).Elem()
}

func (i *baiduCredentialResponsePtrType) ToBaiduCredentialResponsePtrOutput() BaiduCredentialResponsePtrOutput {
	return i.ToBaiduCredentialResponsePtrOutputWithContext(context.Background())
}

func (i *baiduCredentialResponsePtrType) ToBaiduCredentialResponsePtrOutputWithContext(ctx context.Context) BaiduCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduCredentialResponsePtrOutput)
}

type BaiduCredentialResponseOutput struct{ *pulumi.OutputState }

func (BaiduCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BaiduCredentialResponse)(nil)).Elem()
}

func (o BaiduCredentialResponseOutput) ToBaiduCredentialResponseOutput() BaiduCredentialResponseOutput {
	return o
}

func (o BaiduCredentialResponseOutput) ToBaiduCredentialResponseOutputWithContext(ctx context.Context) BaiduCredentialResponseOutput {
	return o
}

func (o BaiduCredentialResponseOutput) ToBaiduCredentialResponsePtrOutput() BaiduCredentialResponsePtrOutput {
	return o.ToBaiduCredentialResponsePtrOutputWithContext(context.Background())
}

func (o BaiduCredentialResponseOutput) ToBaiduCredentialResponsePtrOutputWithContext(ctx context.Context) BaiduCredentialResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BaiduCredentialResponse) *BaiduCredentialResponse {
		return &v
	}).(BaiduCredentialResponsePtrOutput)
}

func (o BaiduCredentialResponseOutput) BaiduApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaiduCredentialResponse) *string { return v.BaiduApiKey }).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialResponseOutput) BaiduEndPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaiduCredentialResponse) *string { return v.BaiduEndPoint }).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialResponseOutput) BaiduSecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaiduCredentialResponse) *string { return v.BaiduSecretKey }).(pulumi.StringPtrOutput)
}

type BaiduCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (BaiduCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BaiduCredentialResponse)(nil)).Elem()
}

func (o BaiduCredentialResponsePtrOutput) ToBaiduCredentialResponsePtrOutput() BaiduCredentialResponsePtrOutput {
	return o
}

func (o BaiduCredentialResponsePtrOutput) ToBaiduCredentialResponsePtrOutputWithContext(ctx context.Context) BaiduCredentialResponsePtrOutput {
	return o
}

func (o BaiduCredentialResponsePtrOutput) Elem() BaiduCredentialResponseOutput {
	return o.ApplyT(func(v *BaiduCredentialResponse) BaiduCredentialResponse {
		if v != nil {
			return *v
		}
		var ret BaiduCredentialResponse
		return ret
	}).(BaiduCredentialResponseOutput)
}

func (o BaiduCredentialResponsePtrOutput) BaiduApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaiduCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.BaiduApiKey
	}).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialResponsePtrOutput) BaiduEndPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaiduCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.BaiduEndPoint
	}).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialResponsePtrOutput) BaiduSecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaiduCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.BaiduSecretKey
	}).(pulumi.StringPtrOutput)
}

type GcmCredential struct {
	GcmEndpoint  *string `pulumi:"gcmEndpoint"`
	GoogleApiKey *string `pulumi:"googleApiKey"`
}





type GcmCredentialInput interface {
	pulumi.Input

	ToGcmCredentialOutput() GcmCredentialOutput
	ToGcmCredentialOutputWithContext(context.Context) GcmCredentialOutput
}

type GcmCredentialArgs struct {
	GcmEndpoint  pulumi.StringPtrInput `pulumi:"gcmEndpoint"`
	GoogleApiKey pulumi.StringPtrInput `pulumi:"googleApiKey"`
}

func (GcmCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GcmCredential)(nil)).Elem()
}

func (i GcmCredentialArgs) ToGcmCredentialOutput() GcmCredentialOutput {
	return i.ToGcmCredentialOutputWithContext(context.Background())
}

func (i GcmCredentialArgs) ToGcmCredentialOutputWithContext(ctx context.Context) GcmCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmCredentialOutput)
}

func (i GcmCredentialArgs) ToGcmCredentialPtrOutput() GcmCredentialPtrOutput {
	return i.ToGcmCredentialPtrOutputWithContext(context.Background())
}

func (i GcmCredentialArgs) ToGcmCredentialPtrOutputWithContext(ctx context.Context) GcmCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmCredentialOutput).ToGcmCredentialPtrOutputWithContext(ctx)
}









type GcmCredentialPtrInput interface {
	pulumi.Input

	ToGcmCredentialPtrOutput() GcmCredentialPtrOutput
	ToGcmCredentialPtrOutputWithContext(context.Context) GcmCredentialPtrOutput
}

type gcmCredentialPtrType GcmCredentialArgs

func GcmCredentialPtr(v *GcmCredentialArgs) GcmCredentialPtrInput {
	return (*gcmCredentialPtrType)(v)
}

func (*gcmCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GcmCredential)(nil)).Elem()
}

func (i *gcmCredentialPtrType) ToGcmCredentialPtrOutput() GcmCredentialPtrOutput {
	return i.ToGcmCredentialPtrOutputWithContext(context.Background())
}

func (i *gcmCredentialPtrType) ToGcmCredentialPtrOutputWithContext(ctx context.Context) GcmCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmCredentialPtrOutput)
}

type GcmCredentialOutput struct{ *pulumi.OutputState }

func (GcmCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GcmCredential)(nil)).Elem()
}

func (o GcmCredentialOutput) ToGcmCredentialOutput() GcmCredentialOutput {
	return o
}

func (o GcmCredentialOutput) ToGcmCredentialOutputWithContext(ctx context.Context) GcmCredentialOutput {
	return o
}

func (o GcmCredentialOutput) ToGcmCredentialPtrOutput() GcmCredentialPtrOutput {
	return o.ToGcmCredentialPtrOutputWithContext(context.Background())
}

func (o GcmCredentialOutput) ToGcmCredentialPtrOutputWithContext(ctx context.Context) GcmCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GcmCredential) *GcmCredential {
		return &v
	}).(GcmCredentialPtrOutput)
}

func (o GcmCredentialOutput) GcmEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GcmCredential) *string { return v.GcmEndpoint }).(pulumi.StringPtrOutput)
}

func (o GcmCredentialOutput) GoogleApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GcmCredential) *string { return v.GoogleApiKey }).(pulumi.StringPtrOutput)
}

type GcmCredentialPtrOutput struct{ *pulumi.OutputState }

func (GcmCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GcmCredential)(nil)).Elem()
}

func (o GcmCredentialPtrOutput) ToGcmCredentialPtrOutput() GcmCredentialPtrOutput {
	return o
}

func (o GcmCredentialPtrOutput) ToGcmCredentialPtrOutputWithContext(ctx context.Context) GcmCredentialPtrOutput {
	return o
}

func (o GcmCredentialPtrOutput) Elem() GcmCredentialOutput {
	return o.ApplyT(func(v *GcmCredential) GcmCredential {
		if v != nil {
			return *v
		}
		var ret GcmCredential
		return ret
	}).(GcmCredentialOutput)
}

func (o GcmCredentialPtrOutput) GcmEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GcmCredential) *string {
		if v == nil {
			return nil
		}
		return v.GcmEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o GcmCredentialPtrOutput) GoogleApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GcmCredential) *string {
		if v == nil {
			return nil
		}
		return v.GoogleApiKey
	}).(pulumi.StringPtrOutput)
}

type GcmCredentialResponse struct {
	GcmEndpoint  *string `pulumi:"gcmEndpoint"`
	GoogleApiKey *string `pulumi:"googleApiKey"`
}





type GcmCredentialResponseInput interface {
	pulumi.Input

	ToGcmCredentialResponseOutput() GcmCredentialResponseOutput
	ToGcmCredentialResponseOutputWithContext(context.Context) GcmCredentialResponseOutput
}

type GcmCredentialResponseArgs struct {
	GcmEndpoint  pulumi.StringPtrInput `pulumi:"gcmEndpoint"`
	GoogleApiKey pulumi.StringPtrInput `pulumi:"googleApiKey"`
}

func (GcmCredentialResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GcmCredentialResponse)(nil)).Elem()
}

func (i GcmCredentialResponseArgs) ToGcmCredentialResponseOutput() GcmCredentialResponseOutput {
	return i.ToGcmCredentialResponseOutputWithContext(context.Background())
}

func (i GcmCredentialResponseArgs) ToGcmCredentialResponseOutputWithContext(ctx context.Context) GcmCredentialResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmCredentialResponseOutput)
}

func (i GcmCredentialResponseArgs) ToGcmCredentialResponsePtrOutput() GcmCredentialResponsePtrOutput {
	return i.ToGcmCredentialResponsePtrOutputWithContext(context.Background())
}

func (i GcmCredentialResponseArgs) ToGcmCredentialResponsePtrOutputWithContext(ctx context.Context) GcmCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmCredentialResponseOutput).ToGcmCredentialResponsePtrOutputWithContext(ctx)
}









type GcmCredentialResponsePtrInput interface {
	pulumi.Input

	ToGcmCredentialResponsePtrOutput() GcmCredentialResponsePtrOutput
	ToGcmCredentialResponsePtrOutputWithContext(context.Context) GcmCredentialResponsePtrOutput
}

type gcmCredentialResponsePtrType GcmCredentialResponseArgs

func GcmCredentialResponsePtr(v *GcmCredentialResponseArgs) GcmCredentialResponsePtrInput {
	return (*gcmCredentialResponsePtrType)(v)
}

func (*gcmCredentialResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GcmCredentialResponse)(nil)).Elem()
}

func (i *gcmCredentialResponsePtrType) ToGcmCredentialResponsePtrOutput() GcmCredentialResponsePtrOutput {
	return i.ToGcmCredentialResponsePtrOutputWithContext(context.Background())
}

func (i *gcmCredentialResponsePtrType) ToGcmCredentialResponsePtrOutputWithContext(ctx context.Context) GcmCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmCredentialResponsePtrOutput)
}

type GcmCredentialResponseOutput struct{ *pulumi.OutputState }

func (GcmCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GcmCredentialResponse)(nil)).Elem()
}

func (o GcmCredentialResponseOutput) ToGcmCredentialResponseOutput() GcmCredentialResponseOutput {
	return o
}

func (o GcmCredentialResponseOutput) ToGcmCredentialResponseOutputWithContext(ctx context.Context) GcmCredentialResponseOutput {
	return o
}

func (o GcmCredentialResponseOutput) ToGcmCredentialResponsePtrOutput() GcmCredentialResponsePtrOutput {
	return o.ToGcmCredentialResponsePtrOutputWithContext(context.Background())
}

func (o GcmCredentialResponseOutput) ToGcmCredentialResponsePtrOutputWithContext(ctx context.Context) GcmCredentialResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GcmCredentialResponse) *GcmCredentialResponse {
		return &v
	}).(GcmCredentialResponsePtrOutput)
}

func (o GcmCredentialResponseOutput) GcmEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GcmCredentialResponse) *string { return v.GcmEndpoint }).(pulumi.StringPtrOutput)
}

func (o GcmCredentialResponseOutput) GoogleApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GcmCredentialResponse) *string { return v.GoogleApiKey }).(pulumi.StringPtrOutput)
}

type GcmCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (GcmCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GcmCredentialResponse)(nil)).Elem()
}

func (o GcmCredentialResponsePtrOutput) ToGcmCredentialResponsePtrOutput() GcmCredentialResponsePtrOutput {
	return o
}

func (o GcmCredentialResponsePtrOutput) ToGcmCredentialResponsePtrOutputWithContext(ctx context.Context) GcmCredentialResponsePtrOutput {
	return o
}

func (o GcmCredentialResponsePtrOutput) Elem() GcmCredentialResponseOutput {
	return o.ApplyT(func(v *GcmCredentialResponse) GcmCredentialResponse {
		if v != nil {
			return *v
		}
		var ret GcmCredentialResponse
		return ret
	}).(GcmCredentialResponseOutput)
}

func (o GcmCredentialResponsePtrOutput) GcmEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GcmCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.GcmEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o GcmCredentialResponsePtrOutput) GoogleApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GcmCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.GoogleApiKey
	}).(pulumi.StringPtrOutput)
}

type MpnsCredential struct {
	CertificateKey  *string `pulumi:"certificateKey"`
	MpnsCertificate *string `pulumi:"mpnsCertificate"`
	Thumbprint      *string `pulumi:"thumbprint"`
}





type MpnsCredentialInput interface {
	pulumi.Input

	ToMpnsCredentialOutput() MpnsCredentialOutput
	ToMpnsCredentialOutputWithContext(context.Context) MpnsCredentialOutput
}

type MpnsCredentialArgs struct {
	CertificateKey  pulumi.StringPtrInput `pulumi:"certificateKey"`
	MpnsCertificate pulumi.StringPtrInput `pulumi:"mpnsCertificate"`
	Thumbprint      pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (MpnsCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MpnsCredential)(nil)).Elem()
}

func (i MpnsCredentialArgs) ToMpnsCredentialOutput() MpnsCredentialOutput {
	return i.ToMpnsCredentialOutputWithContext(context.Background())
}

func (i MpnsCredentialArgs) ToMpnsCredentialOutputWithContext(ctx context.Context) MpnsCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpnsCredentialOutput)
}

func (i MpnsCredentialArgs) ToMpnsCredentialPtrOutput() MpnsCredentialPtrOutput {
	return i.ToMpnsCredentialPtrOutputWithContext(context.Background())
}

func (i MpnsCredentialArgs) ToMpnsCredentialPtrOutputWithContext(ctx context.Context) MpnsCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpnsCredentialOutput).ToMpnsCredentialPtrOutputWithContext(ctx)
}









type MpnsCredentialPtrInput interface {
	pulumi.Input

	ToMpnsCredentialPtrOutput() MpnsCredentialPtrOutput
	ToMpnsCredentialPtrOutputWithContext(context.Context) MpnsCredentialPtrOutput
}

type mpnsCredentialPtrType MpnsCredentialArgs

func MpnsCredentialPtr(v *MpnsCredentialArgs) MpnsCredentialPtrInput {
	return (*mpnsCredentialPtrType)(v)
}

func (*mpnsCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MpnsCredential)(nil)).Elem()
}

func (i *mpnsCredentialPtrType) ToMpnsCredentialPtrOutput() MpnsCredentialPtrOutput {
	return i.ToMpnsCredentialPtrOutputWithContext(context.Background())
}

func (i *mpnsCredentialPtrType) ToMpnsCredentialPtrOutputWithContext(ctx context.Context) MpnsCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpnsCredentialPtrOutput)
}

type MpnsCredentialOutput struct{ *pulumi.OutputState }

func (MpnsCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MpnsCredential)(nil)).Elem()
}

func (o MpnsCredentialOutput) ToMpnsCredentialOutput() MpnsCredentialOutput {
	return o
}

func (o MpnsCredentialOutput) ToMpnsCredentialOutputWithContext(ctx context.Context) MpnsCredentialOutput {
	return o
}

func (o MpnsCredentialOutput) ToMpnsCredentialPtrOutput() MpnsCredentialPtrOutput {
	return o.ToMpnsCredentialPtrOutputWithContext(context.Background())
}

func (o MpnsCredentialOutput) ToMpnsCredentialPtrOutputWithContext(ctx context.Context) MpnsCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MpnsCredential) *MpnsCredential {
		return &v
	}).(MpnsCredentialPtrOutput)
}

func (o MpnsCredentialOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MpnsCredential) *string { return v.CertificateKey }).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialOutput) MpnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MpnsCredential) *string { return v.MpnsCertificate }).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MpnsCredential) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type MpnsCredentialPtrOutput struct{ *pulumi.OutputState }

func (MpnsCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MpnsCredential)(nil)).Elem()
}

func (o MpnsCredentialPtrOutput) ToMpnsCredentialPtrOutput() MpnsCredentialPtrOutput {
	return o
}

func (o MpnsCredentialPtrOutput) ToMpnsCredentialPtrOutputWithContext(ctx context.Context) MpnsCredentialPtrOutput {
	return o
}

func (o MpnsCredentialPtrOutput) Elem() MpnsCredentialOutput {
	return o.ApplyT(func(v *MpnsCredential) MpnsCredential {
		if v != nil {
			return *v
		}
		var ret MpnsCredential
		return ret
	}).(MpnsCredentialOutput)
}

func (o MpnsCredentialPtrOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MpnsCredential) *string {
		if v == nil {
			return nil
		}
		return v.CertificateKey
	}).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialPtrOutput) MpnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MpnsCredential) *string {
		if v == nil {
			return nil
		}
		return v.MpnsCertificate
	}).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialPtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MpnsCredential) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type MpnsCredentialResponse struct {
	CertificateKey  *string `pulumi:"certificateKey"`
	MpnsCertificate *string `pulumi:"mpnsCertificate"`
	Thumbprint      *string `pulumi:"thumbprint"`
}





type MpnsCredentialResponseInput interface {
	pulumi.Input

	ToMpnsCredentialResponseOutput() MpnsCredentialResponseOutput
	ToMpnsCredentialResponseOutputWithContext(context.Context) MpnsCredentialResponseOutput
}

type MpnsCredentialResponseArgs struct {
	CertificateKey  pulumi.StringPtrInput `pulumi:"certificateKey"`
	MpnsCertificate pulumi.StringPtrInput `pulumi:"mpnsCertificate"`
	Thumbprint      pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (MpnsCredentialResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MpnsCredentialResponse)(nil)).Elem()
}

func (i MpnsCredentialResponseArgs) ToMpnsCredentialResponseOutput() MpnsCredentialResponseOutput {
	return i.ToMpnsCredentialResponseOutputWithContext(context.Background())
}

func (i MpnsCredentialResponseArgs) ToMpnsCredentialResponseOutputWithContext(ctx context.Context) MpnsCredentialResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpnsCredentialResponseOutput)
}

func (i MpnsCredentialResponseArgs) ToMpnsCredentialResponsePtrOutput() MpnsCredentialResponsePtrOutput {
	return i.ToMpnsCredentialResponsePtrOutputWithContext(context.Background())
}

func (i MpnsCredentialResponseArgs) ToMpnsCredentialResponsePtrOutputWithContext(ctx context.Context) MpnsCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpnsCredentialResponseOutput).ToMpnsCredentialResponsePtrOutputWithContext(ctx)
}









type MpnsCredentialResponsePtrInput interface {
	pulumi.Input

	ToMpnsCredentialResponsePtrOutput() MpnsCredentialResponsePtrOutput
	ToMpnsCredentialResponsePtrOutputWithContext(context.Context) MpnsCredentialResponsePtrOutput
}

type mpnsCredentialResponsePtrType MpnsCredentialResponseArgs

func MpnsCredentialResponsePtr(v *MpnsCredentialResponseArgs) MpnsCredentialResponsePtrInput {
	return (*mpnsCredentialResponsePtrType)(v)
}

func (*mpnsCredentialResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MpnsCredentialResponse)(nil)).Elem()
}

func (i *mpnsCredentialResponsePtrType) ToMpnsCredentialResponsePtrOutput() MpnsCredentialResponsePtrOutput {
	return i.ToMpnsCredentialResponsePtrOutputWithContext(context.Background())
}

func (i *mpnsCredentialResponsePtrType) ToMpnsCredentialResponsePtrOutputWithContext(ctx context.Context) MpnsCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpnsCredentialResponsePtrOutput)
}

type MpnsCredentialResponseOutput struct{ *pulumi.OutputState }

func (MpnsCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MpnsCredentialResponse)(nil)).Elem()
}

func (o MpnsCredentialResponseOutput) ToMpnsCredentialResponseOutput() MpnsCredentialResponseOutput {
	return o
}

func (o MpnsCredentialResponseOutput) ToMpnsCredentialResponseOutputWithContext(ctx context.Context) MpnsCredentialResponseOutput {
	return o
}

func (o MpnsCredentialResponseOutput) ToMpnsCredentialResponsePtrOutput() MpnsCredentialResponsePtrOutput {
	return o.ToMpnsCredentialResponsePtrOutputWithContext(context.Background())
}

func (o MpnsCredentialResponseOutput) ToMpnsCredentialResponsePtrOutputWithContext(ctx context.Context) MpnsCredentialResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MpnsCredentialResponse) *MpnsCredentialResponse {
		return &v
	}).(MpnsCredentialResponsePtrOutput)
}

func (o MpnsCredentialResponseOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MpnsCredentialResponse) *string { return v.CertificateKey }).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialResponseOutput) MpnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MpnsCredentialResponse) *string { return v.MpnsCertificate }).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MpnsCredentialResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type MpnsCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (MpnsCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MpnsCredentialResponse)(nil)).Elem()
}

func (o MpnsCredentialResponsePtrOutput) ToMpnsCredentialResponsePtrOutput() MpnsCredentialResponsePtrOutput {
	return o
}

func (o MpnsCredentialResponsePtrOutput) ToMpnsCredentialResponsePtrOutputWithContext(ctx context.Context) MpnsCredentialResponsePtrOutput {
	return o
}

func (o MpnsCredentialResponsePtrOutput) Elem() MpnsCredentialResponseOutput {
	return o.ApplyT(func(v *MpnsCredentialResponse) MpnsCredentialResponse {
		if v != nil {
			return *v
		}
		var ret MpnsCredentialResponse
		return ret
	}).(MpnsCredentialResponseOutput)
}

func (o MpnsCredentialResponsePtrOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MpnsCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.CertificateKey
	}).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialResponsePtrOutput) MpnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MpnsCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.MpnsCertificate
	}).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MpnsCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type SharedAccessAuthorizationRuleProperties struct {
	Rights []AccessRights `pulumi:"rights"`
}





type SharedAccessAuthorizationRulePropertiesInput interface {
	pulumi.Input

	ToSharedAccessAuthorizationRulePropertiesOutput() SharedAccessAuthorizationRulePropertiesOutput
	ToSharedAccessAuthorizationRulePropertiesOutputWithContext(context.Context) SharedAccessAuthorizationRulePropertiesOutput
}

type SharedAccessAuthorizationRulePropertiesArgs struct {
	Rights AccessRightsArrayInput `pulumi:"rights"`
}

func (SharedAccessAuthorizationRulePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessAuthorizationRuleProperties)(nil)).Elem()
}

func (i SharedAccessAuthorizationRulePropertiesArgs) ToSharedAccessAuthorizationRulePropertiesOutput() SharedAccessAuthorizationRulePropertiesOutput {
	return i.ToSharedAccessAuthorizationRulePropertiesOutputWithContext(context.Background())
}

func (i SharedAccessAuthorizationRulePropertiesArgs) ToSharedAccessAuthorizationRulePropertiesOutputWithContext(ctx context.Context) SharedAccessAuthorizationRulePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessAuthorizationRulePropertiesOutput)
}

func (i SharedAccessAuthorizationRulePropertiesArgs) ToSharedAccessAuthorizationRulePropertiesPtrOutput() SharedAccessAuthorizationRulePropertiesPtrOutput {
	return i.ToSharedAccessAuthorizationRulePropertiesPtrOutputWithContext(context.Background())
}

func (i SharedAccessAuthorizationRulePropertiesArgs) ToSharedAccessAuthorizationRulePropertiesPtrOutputWithContext(ctx context.Context) SharedAccessAuthorizationRulePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessAuthorizationRulePropertiesOutput).ToSharedAccessAuthorizationRulePropertiesPtrOutputWithContext(ctx)
}









type SharedAccessAuthorizationRulePropertiesPtrInput interface {
	pulumi.Input

	ToSharedAccessAuthorizationRulePropertiesPtrOutput() SharedAccessAuthorizationRulePropertiesPtrOutput
	ToSharedAccessAuthorizationRulePropertiesPtrOutputWithContext(context.Context) SharedAccessAuthorizationRulePropertiesPtrOutput
}

type sharedAccessAuthorizationRulePropertiesPtrType SharedAccessAuthorizationRulePropertiesArgs

func SharedAccessAuthorizationRulePropertiesPtr(v *SharedAccessAuthorizationRulePropertiesArgs) SharedAccessAuthorizationRulePropertiesPtrInput {
	return (*sharedAccessAuthorizationRulePropertiesPtrType)(v)
}

func (*sharedAccessAuthorizationRulePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedAccessAuthorizationRuleProperties)(nil)).Elem()
}

func (i *sharedAccessAuthorizationRulePropertiesPtrType) ToSharedAccessAuthorizationRulePropertiesPtrOutput() SharedAccessAuthorizationRulePropertiesPtrOutput {
	return i.ToSharedAccessAuthorizationRulePropertiesPtrOutputWithContext(context.Background())
}

func (i *sharedAccessAuthorizationRulePropertiesPtrType) ToSharedAccessAuthorizationRulePropertiesPtrOutputWithContext(ctx context.Context) SharedAccessAuthorizationRulePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessAuthorizationRulePropertiesPtrOutput)
}





type SharedAccessAuthorizationRulePropertiesArrayInput interface {
	pulumi.Input

	ToSharedAccessAuthorizationRulePropertiesArrayOutput() SharedAccessAuthorizationRulePropertiesArrayOutput
	ToSharedAccessAuthorizationRulePropertiesArrayOutputWithContext(context.Context) SharedAccessAuthorizationRulePropertiesArrayOutput
}

type SharedAccessAuthorizationRulePropertiesArray []SharedAccessAuthorizationRulePropertiesInput

func (SharedAccessAuthorizationRulePropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessAuthorizationRuleProperties)(nil)).Elem()
}

func (i SharedAccessAuthorizationRulePropertiesArray) ToSharedAccessAuthorizationRulePropertiesArrayOutput() SharedAccessAuthorizationRulePropertiesArrayOutput {
	return i.ToSharedAccessAuthorizationRulePropertiesArrayOutputWithContext(context.Background())
}

func (i SharedAccessAuthorizationRulePropertiesArray) ToSharedAccessAuthorizationRulePropertiesArrayOutputWithContext(ctx context.Context) SharedAccessAuthorizationRulePropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessAuthorizationRulePropertiesArrayOutput)
}

type SharedAccessAuthorizationRulePropertiesOutput struct{ *pulumi.OutputState }

func (SharedAccessAuthorizationRulePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessAuthorizationRuleProperties)(nil)).Elem()
}

func (o SharedAccessAuthorizationRulePropertiesOutput) ToSharedAccessAuthorizationRulePropertiesOutput() SharedAccessAuthorizationRulePropertiesOutput {
	return o
}

func (o SharedAccessAuthorizationRulePropertiesOutput) ToSharedAccessAuthorizationRulePropertiesOutputWithContext(ctx context.Context) SharedAccessAuthorizationRulePropertiesOutput {
	return o
}

func (o SharedAccessAuthorizationRulePropertiesOutput) ToSharedAccessAuthorizationRulePropertiesPtrOutput() SharedAccessAuthorizationRulePropertiesPtrOutput {
	return o.ToSharedAccessAuthorizationRulePropertiesPtrOutputWithContext(context.Background())
}

func (o SharedAccessAuthorizationRulePropertiesOutput) ToSharedAccessAuthorizationRulePropertiesPtrOutputWithContext(ctx context.Context) SharedAccessAuthorizationRulePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SharedAccessAuthorizationRuleProperties) *SharedAccessAuthorizationRuleProperties {
		return &v
	}).(SharedAccessAuthorizationRulePropertiesPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesOutput) Rights() AccessRightsArrayOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRuleProperties) []AccessRights { return v.Rights }).(AccessRightsArrayOutput)
}

type SharedAccessAuthorizationRulePropertiesPtrOutput struct{ *pulumi.OutputState }

func (SharedAccessAuthorizationRulePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedAccessAuthorizationRuleProperties)(nil)).Elem()
}

func (o SharedAccessAuthorizationRulePropertiesPtrOutput) ToSharedAccessAuthorizationRulePropertiesPtrOutput() SharedAccessAuthorizationRulePropertiesPtrOutput {
	return o
}

func (o SharedAccessAuthorizationRulePropertiesPtrOutput) ToSharedAccessAuthorizationRulePropertiesPtrOutputWithContext(ctx context.Context) SharedAccessAuthorizationRulePropertiesPtrOutput {
	return o
}

func (o SharedAccessAuthorizationRulePropertiesPtrOutput) Elem() SharedAccessAuthorizationRulePropertiesOutput {
	return o.ApplyT(func(v *SharedAccessAuthorizationRuleProperties) SharedAccessAuthorizationRuleProperties {
		if v != nil {
			return *v
		}
		var ret SharedAccessAuthorizationRuleProperties
		return ret
	}).(SharedAccessAuthorizationRulePropertiesOutput)
}

func (o SharedAccessAuthorizationRulePropertiesPtrOutput) Rights() AccessRightsArrayOutput {
	return o.ApplyT(func(v *SharedAccessAuthorizationRuleProperties) []AccessRights {
		if v == nil {
			return nil
		}
		return v.Rights
	}).(AccessRightsArrayOutput)
}

type SharedAccessAuthorizationRulePropertiesArrayOutput struct{ *pulumi.OutputState }

func (SharedAccessAuthorizationRulePropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessAuthorizationRuleProperties)(nil)).Elem()
}

func (o SharedAccessAuthorizationRulePropertiesArrayOutput) ToSharedAccessAuthorizationRulePropertiesArrayOutput() SharedAccessAuthorizationRulePropertiesArrayOutput {
	return o
}

func (o SharedAccessAuthorizationRulePropertiesArrayOutput) ToSharedAccessAuthorizationRulePropertiesArrayOutputWithContext(ctx context.Context) SharedAccessAuthorizationRulePropertiesArrayOutput {
	return o
}

func (o SharedAccessAuthorizationRulePropertiesArrayOutput) Index(i pulumi.IntInput) SharedAccessAuthorizationRulePropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedAccessAuthorizationRuleProperties {
		return vs[0].([]SharedAccessAuthorizationRuleProperties)[vs[1].(int)]
	}).(SharedAccessAuthorizationRulePropertiesOutput)
}

type SharedAccessAuthorizationRulePropertiesResponse struct {
	Rights []string `pulumi:"rights"`
}





type SharedAccessAuthorizationRulePropertiesResponseInput interface {
	pulumi.Input

	ToSharedAccessAuthorizationRulePropertiesResponseOutput() SharedAccessAuthorizationRulePropertiesResponseOutput
	ToSharedAccessAuthorizationRulePropertiesResponseOutputWithContext(context.Context) SharedAccessAuthorizationRulePropertiesResponseOutput
}

type SharedAccessAuthorizationRulePropertiesResponseArgs struct {
	Rights pulumi.StringArrayInput `pulumi:"rights"`
}

func (SharedAccessAuthorizationRulePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessAuthorizationRulePropertiesResponse)(nil)).Elem()
}

func (i SharedAccessAuthorizationRulePropertiesResponseArgs) ToSharedAccessAuthorizationRulePropertiesResponseOutput() SharedAccessAuthorizationRulePropertiesResponseOutput {
	return i.ToSharedAccessAuthorizationRulePropertiesResponseOutputWithContext(context.Background())
}

func (i SharedAccessAuthorizationRulePropertiesResponseArgs) ToSharedAccessAuthorizationRulePropertiesResponseOutputWithContext(ctx context.Context) SharedAccessAuthorizationRulePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessAuthorizationRulePropertiesResponseOutput)
}





type SharedAccessAuthorizationRulePropertiesResponseArrayInput interface {
	pulumi.Input

	ToSharedAccessAuthorizationRulePropertiesResponseArrayOutput() SharedAccessAuthorizationRulePropertiesResponseArrayOutput
	ToSharedAccessAuthorizationRulePropertiesResponseArrayOutputWithContext(context.Context) SharedAccessAuthorizationRulePropertiesResponseArrayOutput
}

type SharedAccessAuthorizationRulePropertiesResponseArray []SharedAccessAuthorizationRulePropertiesResponseInput

func (SharedAccessAuthorizationRulePropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessAuthorizationRulePropertiesResponse)(nil)).Elem()
}

func (i SharedAccessAuthorizationRulePropertiesResponseArray) ToSharedAccessAuthorizationRulePropertiesResponseArrayOutput() SharedAccessAuthorizationRulePropertiesResponseArrayOutput {
	return i.ToSharedAccessAuthorizationRulePropertiesResponseArrayOutputWithContext(context.Background())
}

func (i SharedAccessAuthorizationRulePropertiesResponseArray) ToSharedAccessAuthorizationRulePropertiesResponseArrayOutputWithContext(ctx context.Context) SharedAccessAuthorizationRulePropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessAuthorizationRulePropertiesResponseArrayOutput)
}

type SharedAccessAuthorizationRulePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SharedAccessAuthorizationRulePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessAuthorizationRulePropertiesResponse)(nil)).Elem()
}

func (o SharedAccessAuthorizationRulePropertiesResponseOutput) ToSharedAccessAuthorizationRulePropertiesResponseOutput() SharedAccessAuthorizationRulePropertiesResponseOutput {
	return o
}

func (o SharedAccessAuthorizationRulePropertiesResponseOutput) ToSharedAccessAuthorizationRulePropertiesResponseOutputWithContext(ctx context.Context) SharedAccessAuthorizationRulePropertiesResponseOutput {
	return o
}

func (o SharedAccessAuthorizationRulePropertiesResponseOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRulePropertiesResponse) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

type SharedAccessAuthorizationRulePropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (SharedAccessAuthorizationRulePropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessAuthorizationRulePropertiesResponse)(nil)).Elem()
}

func (o SharedAccessAuthorizationRulePropertiesResponseArrayOutput) ToSharedAccessAuthorizationRulePropertiesResponseArrayOutput() SharedAccessAuthorizationRulePropertiesResponseArrayOutput {
	return o
}

func (o SharedAccessAuthorizationRulePropertiesResponseArrayOutput) ToSharedAccessAuthorizationRulePropertiesResponseArrayOutputWithContext(ctx context.Context) SharedAccessAuthorizationRulePropertiesResponseArrayOutput {
	return o
}

func (o SharedAccessAuthorizationRulePropertiesResponseArrayOutput) Index(i pulumi.IntInput) SharedAccessAuthorizationRulePropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedAccessAuthorizationRulePropertiesResponse {
		return vs[0].([]SharedAccessAuthorizationRulePropertiesResponse)[vs[1].(int)]
	}).(SharedAccessAuthorizationRulePropertiesResponseOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
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

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
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

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
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
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
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

type WnsCredential struct {
	PackageSid          *string `pulumi:"packageSid"`
	SecretKey           *string `pulumi:"secretKey"`
	WindowsLiveEndpoint *string `pulumi:"windowsLiveEndpoint"`
}





type WnsCredentialInput interface {
	pulumi.Input

	ToWnsCredentialOutput() WnsCredentialOutput
	ToWnsCredentialOutputWithContext(context.Context) WnsCredentialOutput
}

type WnsCredentialArgs struct {
	PackageSid          pulumi.StringPtrInput `pulumi:"packageSid"`
	SecretKey           pulumi.StringPtrInput `pulumi:"secretKey"`
	WindowsLiveEndpoint pulumi.StringPtrInput `pulumi:"windowsLiveEndpoint"`
}

func (WnsCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WnsCredential)(nil)).Elem()
}

func (i WnsCredentialArgs) ToWnsCredentialOutput() WnsCredentialOutput {
	return i.ToWnsCredentialOutputWithContext(context.Background())
}

func (i WnsCredentialArgs) ToWnsCredentialOutputWithContext(ctx context.Context) WnsCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WnsCredentialOutput)
}

func (i WnsCredentialArgs) ToWnsCredentialPtrOutput() WnsCredentialPtrOutput {
	return i.ToWnsCredentialPtrOutputWithContext(context.Background())
}

func (i WnsCredentialArgs) ToWnsCredentialPtrOutputWithContext(ctx context.Context) WnsCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WnsCredentialOutput).ToWnsCredentialPtrOutputWithContext(ctx)
}









type WnsCredentialPtrInput interface {
	pulumi.Input

	ToWnsCredentialPtrOutput() WnsCredentialPtrOutput
	ToWnsCredentialPtrOutputWithContext(context.Context) WnsCredentialPtrOutput
}

type wnsCredentialPtrType WnsCredentialArgs

func WnsCredentialPtr(v *WnsCredentialArgs) WnsCredentialPtrInput {
	return (*wnsCredentialPtrType)(v)
}

func (*wnsCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WnsCredential)(nil)).Elem()
}

func (i *wnsCredentialPtrType) ToWnsCredentialPtrOutput() WnsCredentialPtrOutput {
	return i.ToWnsCredentialPtrOutputWithContext(context.Background())
}

func (i *wnsCredentialPtrType) ToWnsCredentialPtrOutputWithContext(ctx context.Context) WnsCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WnsCredentialPtrOutput)
}

type WnsCredentialOutput struct{ *pulumi.OutputState }

func (WnsCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WnsCredential)(nil)).Elem()
}

func (o WnsCredentialOutput) ToWnsCredentialOutput() WnsCredentialOutput {
	return o
}

func (o WnsCredentialOutput) ToWnsCredentialOutputWithContext(ctx context.Context) WnsCredentialOutput {
	return o
}

func (o WnsCredentialOutput) ToWnsCredentialPtrOutput() WnsCredentialPtrOutput {
	return o.ToWnsCredentialPtrOutputWithContext(context.Background())
}

func (o WnsCredentialOutput) ToWnsCredentialPtrOutputWithContext(ctx context.Context) WnsCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WnsCredential) *WnsCredential {
		return &v
	}).(WnsCredentialPtrOutput)
}

func (o WnsCredentialOutput) PackageSid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WnsCredential) *string { return v.PackageSid }).(pulumi.StringPtrOutput)
}

func (o WnsCredentialOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WnsCredential) *string { return v.SecretKey }).(pulumi.StringPtrOutput)
}

func (o WnsCredentialOutput) WindowsLiveEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WnsCredential) *string { return v.WindowsLiveEndpoint }).(pulumi.StringPtrOutput)
}

type WnsCredentialPtrOutput struct{ *pulumi.OutputState }

func (WnsCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WnsCredential)(nil)).Elem()
}

func (o WnsCredentialPtrOutput) ToWnsCredentialPtrOutput() WnsCredentialPtrOutput {
	return o
}

func (o WnsCredentialPtrOutput) ToWnsCredentialPtrOutputWithContext(ctx context.Context) WnsCredentialPtrOutput {
	return o
}

func (o WnsCredentialPtrOutput) Elem() WnsCredentialOutput {
	return o.ApplyT(func(v *WnsCredential) WnsCredential {
		if v != nil {
			return *v
		}
		var ret WnsCredential
		return ret
	}).(WnsCredentialOutput)
}

func (o WnsCredentialPtrOutput) PackageSid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WnsCredential) *string {
		if v == nil {
			return nil
		}
		return v.PackageSid
	}).(pulumi.StringPtrOutput)
}

func (o WnsCredentialPtrOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WnsCredential) *string {
		if v == nil {
			return nil
		}
		return v.SecretKey
	}).(pulumi.StringPtrOutput)
}

func (o WnsCredentialPtrOutput) WindowsLiveEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WnsCredential) *string {
		if v == nil {
			return nil
		}
		return v.WindowsLiveEndpoint
	}).(pulumi.StringPtrOutput)
}

type WnsCredentialResponse struct {
	PackageSid          *string `pulumi:"packageSid"`
	SecretKey           *string `pulumi:"secretKey"`
	WindowsLiveEndpoint *string `pulumi:"windowsLiveEndpoint"`
}





type WnsCredentialResponseInput interface {
	pulumi.Input

	ToWnsCredentialResponseOutput() WnsCredentialResponseOutput
	ToWnsCredentialResponseOutputWithContext(context.Context) WnsCredentialResponseOutput
}

type WnsCredentialResponseArgs struct {
	PackageSid          pulumi.StringPtrInput `pulumi:"packageSid"`
	SecretKey           pulumi.StringPtrInput `pulumi:"secretKey"`
	WindowsLiveEndpoint pulumi.StringPtrInput `pulumi:"windowsLiveEndpoint"`
}

func (WnsCredentialResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WnsCredentialResponse)(nil)).Elem()
}

func (i WnsCredentialResponseArgs) ToWnsCredentialResponseOutput() WnsCredentialResponseOutput {
	return i.ToWnsCredentialResponseOutputWithContext(context.Background())
}

func (i WnsCredentialResponseArgs) ToWnsCredentialResponseOutputWithContext(ctx context.Context) WnsCredentialResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WnsCredentialResponseOutput)
}

func (i WnsCredentialResponseArgs) ToWnsCredentialResponsePtrOutput() WnsCredentialResponsePtrOutput {
	return i.ToWnsCredentialResponsePtrOutputWithContext(context.Background())
}

func (i WnsCredentialResponseArgs) ToWnsCredentialResponsePtrOutputWithContext(ctx context.Context) WnsCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WnsCredentialResponseOutput).ToWnsCredentialResponsePtrOutputWithContext(ctx)
}









type WnsCredentialResponsePtrInput interface {
	pulumi.Input

	ToWnsCredentialResponsePtrOutput() WnsCredentialResponsePtrOutput
	ToWnsCredentialResponsePtrOutputWithContext(context.Context) WnsCredentialResponsePtrOutput
}

type wnsCredentialResponsePtrType WnsCredentialResponseArgs

func WnsCredentialResponsePtr(v *WnsCredentialResponseArgs) WnsCredentialResponsePtrInput {
	return (*wnsCredentialResponsePtrType)(v)
}

func (*wnsCredentialResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WnsCredentialResponse)(nil)).Elem()
}

func (i *wnsCredentialResponsePtrType) ToWnsCredentialResponsePtrOutput() WnsCredentialResponsePtrOutput {
	return i.ToWnsCredentialResponsePtrOutputWithContext(context.Background())
}

func (i *wnsCredentialResponsePtrType) ToWnsCredentialResponsePtrOutputWithContext(ctx context.Context) WnsCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WnsCredentialResponsePtrOutput)
}

type WnsCredentialResponseOutput struct{ *pulumi.OutputState }

func (WnsCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WnsCredentialResponse)(nil)).Elem()
}

func (o WnsCredentialResponseOutput) ToWnsCredentialResponseOutput() WnsCredentialResponseOutput {
	return o
}

func (o WnsCredentialResponseOutput) ToWnsCredentialResponseOutputWithContext(ctx context.Context) WnsCredentialResponseOutput {
	return o
}

func (o WnsCredentialResponseOutput) ToWnsCredentialResponsePtrOutput() WnsCredentialResponsePtrOutput {
	return o.ToWnsCredentialResponsePtrOutputWithContext(context.Background())
}

func (o WnsCredentialResponseOutput) ToWnsCredentialResponsePtrOutputWithContext(ctx context.Context) WnsCredentialResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WnsCredentialResponse) *WnsCredentialResponse {
		return &v
	}).(WnsCredentialResponsePtrOutput)
}

func (o WnsCredentialResponseOutput) PackageSid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WnsCredentialResponse) *string { return v.PackageSid }).(pulumi.StringPtrOutput)
}

func (o WnsCredentialResponseOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WnsCredentialResponse) *string { return v.SecretKey }).(pulumi.StringPtrOutput)
}

func (o WnsCredentialResponseOutput) WindowsLiveEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WnsCredentialResponse) *string { return v.WindowsLiveEndpoint }).(pulumi.StringPtrOutput)
}

type WnsCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (WnsCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WnsCredentialResponse)(nil)).Elem()
}

func (o WnsCredentialResponsePtrOutput) ToWnsCredentialResponsePtrOutput() WnsCredentialResponsePtrOutput {
	return o
}

func (o WnsCredentialResponsePtrOutput) ToWnsCredentialResponsePtrOutputWithContext(ctx context.Context) WnsCredentialResponsePtrOutput {
	return o
}

func (o WnsCredentialResponsePtrOutput) Elem() WnsCredentialResponseOutput {
	return o.ApplyT(func(v *WnsCredentialResponse) WnsCredentialResponse {
		if v != nil {
			return *v
		}
		var ret WnsCredentialResponse
		return ret
	}).(WnsCredentialResponseOutput)
}

func (o WnsCredentialResponsePtrOutput) PackageSid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WnsCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.PackageSid
	}).(pulumi.StringPtrOutput)
}

func (o WnsCredentialResponsePtrOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WnsCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecretKey
	}).(pulumi.StringPtrOutput)
}

func (o WnsCredentialResponsePtrOutput) WindowsLiveEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WnsCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.WindowsLiveEndpoint
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AdmCredentialOutput{})
	pulumi.RegisterOutputType(AdmCredentialPtrOutput{})
	pulumi.RegisterOutputType(AdmCredentialResponseOutput{})
	pulumi.RegisterOutputType(AdmCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(ApnsCredentialOutput{})
	pulumi.RegisterOutputType(ApnsCredentialPtrOutput{})
	pulumi.RegisterOutputType(ApnsCredentialResponseOutput{})
	pulumi.RegisterOutputType(ApnsCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(BaiduCredentialOutput{})
	pulumi.RegisterOutputType(BaiduCredentialPtrOutput{})
	pulumi.RegisterOutputType(BaiduCredentialResponseOutput{})
	pulumi.RegisterOutputType(BaiduCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(GcmCredentialOutput{})
	pulumi.RegisterOutputType(GcmCredentialPtrOutput{})
	pulumi.RegisterOutputType(GcmCredentialResponseOutput{})
	pulumi.RegisterOutputType(GcmCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(MpnsCredentialOutput{})
	pulumi.RegisterOutputType(MpnsCredentialPtrOutput{})
	pulumi.RegisterOutputType(MpnsCredentialResponseOutput{})
	pulumi.RegisterOutputType(MpnsCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(SharedAccessAuthorizationRulePropertiesOutput{})
	pulumi.RegisterOutputType(SharedAccessAuthorizationRulePropertiesPtrOutput{})
	pulumi.RegisterOutputType(SharedAccessAuthorizationRulePropertiesArrayOutput{})
	pulumi.RegisterOutputType(SharedAccessAuthorizationRulePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SharedAccessAuthorizationRulePropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(WnsCredentialOutput{})
	pulumi.RegisterOutputType(WnsCredentialPtrOutput{})
	pulumi.RegisterOutputType(WnsCredentialResponseOutput{})
	pulumi.RegisterOutputType(WnsCredentialResponsePtrOutput{})
}

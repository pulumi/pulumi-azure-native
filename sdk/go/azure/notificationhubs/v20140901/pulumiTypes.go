


package v20140901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdmCredential struct {
	Properties *AdmCredentialProperties `pulumi:"properties"`
}





type AdmCredentialInput interface {
	pulumi.Input

	ToAdmCredentialOutput() AdmCredentialOutput
	ToAdmCredentialOutputWithContext(context.Context) AdmCredentialOutput
}

type AdmCredentialArgs struct {
	Properties AdmCredentialPropertiesPtrInput `pulumi:"properties"`
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

func (o AdmCredentialOutput) Properties() AdmCredentialPropertiesPtrOutput {
	return o.ApplyT(func(v AdmCredential) *AdmCredentialProperties { return v.Properties }).(AdmCredentialPropertiesPtrOutput)
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

func (o AdmCredentialPtrOutput) Properties() AdmCredentialPropertiesPtrOutput {
	return o.ApplyT(func(v *AdmCredential) *AdmCredentialProperties {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(AdmCredentialPropertiesPtrOutput)
}

type AdmCredentialProperties struct {
	AuthTokenUrl *string `pulumi:"authTokenUrl"`
	ClientId     *string `pulumi:"clientId"`
	ClientSecret *string `pulumi:"clientSecret"`
}





type AdmCredentialPropertiesInput interface {
	pulumi.Input

	ToAdmCredentialPropertiesOutput() AdmCredentialPropertiesOutput
	ToAdmCredentialPropertiesOutputWithContext(context.Context) AdmCredentialPropertiesOutput
}

type AdmCredentialPropertiesArgs struct {
	AuthTokenUrl pulumi.StringPtrInput `pulumi:"authTokenUrl"`
	ClientId     pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecret pulumi.StringPtrInput `pulumi:"clientSecret"`
}

func (AdmCredentialPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdmCredentialProperties)(nil)).Elem()
}

func (i AdmCredentialPropertiesArgs) ToAdmCredentialPropertiesOutput() AdmCredentialPropertiesOutput {
	return i.ToAdmCredentialPropertiesOutputWithContext(context.Background())
}

func (i AdmCredentialPropertiesArgs) ToAdmCredentialPropertiesOutputWithContext(ctx context.Context) AdmCredentialPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmCredentialPropertiesOutput)
}

func (i AdmCredentialPropertiesArgs) ToAdmCredentialPropertiesPtrOutput() AdmCredentialPropertiesPtrOutput {
	return i.ToAdmCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (i AdmCredentialPropertiesArgs) ToAdmCredentialPropertiesPtrOutputWithContext(ctx context.Context) AdmCredentialPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmCredentialPropertiesOutput).ToAdmCredentialPropertiesPtrOutputWithContext(ctx)
}









type AdmCredentialPropertiesPtrInput interface {
	pulumi.Input

	ToAdmCredentialPropertiesPtrOutput() AdmCredentialPropertiesPtrOutput
	ToAdmCredentialPropertiesPtrOutputWithContext(context.Context) AdmCredentialPropertiesPtrOutput
}

type admCredentialPropertiesPtrType AdmCredentialPropertiesArgs

func AdmCredentialPropertiesPtr(v *AdmCredentialPropertiesArgs) AdmCredentialPropertiesPtrInput {
	return (*admCredentialPropertiesPtrType)(v)
}

func (*admCredentialPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AdmCredentialProperties)(nil)).Elem()
}

func (i *admCredentialPropertiesPtrType) ToAdmCredentialPropertiesPtrOutput() AdmCredentialPropertiesPtrOutput {
	return i.ToAdmCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (i *admCredentialPropertiesPtrType) ToAdmCredentialPropertiesPtrOutputWithContext(ctx context.Context) AdmCredentialPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmCredentialPropertiesPtrOutput)
}

type AdmCredentialPropertiesOutput struct{ *pulumi.OutputState }

func (AdmCredentialPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdmCredentialProperties)(nil)).Elem()
}

func (o AdmCredentialPropertiesOutput) ToAdmCredentialPropertiesOutput() AdmCredentialPropertiesOutput {
	return o
}

func (o AdmCredentialPropertiesOutput) ToAdmCredentialPropertiesOutputWithContext(ctx context.Context) AdmCredentialPropertiesOutput {
	return o
}

func (o AdmCredentialPropertiesOutput) ToAdmCredentialPropertiesPtrOutput() AdmCredentialPropertiesPtrOutput {
	return o.ToAdmCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (o AdmCredentialPropertiesOutput) ToAdmCredentialPropertiesPtrOutputWithContext(ctx context.Context) AdmCredentialPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdmCredentialProperties) *AdmCredentialProperties {
		return &v
	}).(AdmCredentialPropertiesPtrOutput)
}

func (o AdmCredentialPropertiesOutput) AuthTokenUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdmCredentialProperties) *string { return v.AuthTokenUrl }).(pulumi.StringPtrOutput)
}

func (o AdmCredentialPropertiesOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdmCredentialProperties) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AdmCredentialPropertiesOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdmCredentialProperties) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

type AdmCredentialPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AdmCredentialPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdmCredentialProperties)(nil)).Elem()
}

func (o AdmCredentialPropertiesPtrOutput) ToAdmCredentialPropertiesPtrOutput() AdmCredentialPropertiesPtrOutput {
	return o
}

func (o AdmCredentialPropertiesPtrOutput) ToAdmCredentialPropertiesPtrOutputWithContext(ctx context.Context) AdmCredentialPropertiesPtrOutput {
	return o
}

func (o AdmCredentialPropertiesPtrOutput) Elem() AdmCredentialPropertiesOutput {
	return o.ApplyT(func(v *AdmCredentialProperties) AdmCredentialProperties {
		if v != nil {
			return *v
		}
		var ret AdmCredentialProperties
		return ret
	}).(AdmCredentialPropertiesOutput)
}

func (o AdmCredentialPropertiesPtrOutput) AuthTokenUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdmCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.AuthTokenUrl
	}).(pulumi.StringPtrOutput)
}

func (o AdmCredentialPropertiesPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdmCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AdmCredentialPropertiesPtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdmCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

type AdmCredentialPropertiesResponse struct {
	AuthTokenUrl *string `pulumi:"authTokenUrl"`
	ClientId     *string `pulumi:"clientId"`
	ClientSecret *string `pulumi:"clientSecret"`
}





type AdmCredentialPropertiesResponseInput interface {
	pulumi.Input

	ToAdmCredentialPropertiesResponseOutput() AdmCredentialPropertiesResponseOutput
	ToAdmCredentialPropertiesResponseOutputWithContext(context.Context) AdmCredentialPropertiesResponseOutput
}

type AdmCredentialPropertiesResponseArgs struct {
	AuthTokenUrl pulumi.StringPtrInput `pulumi:"authTokenUrl"`
	ClientId     pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecret pulumi.StringPtrInput `pulumi:"clientSecret"`
}

func (AdmCredentialPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdmCredentialPropertiesResponse)(nil)).Elem()
}

func (i AdmCredentialPropertiesResponseArgs) ToAdmCredentialPropertiesResponseOutput() AdmCredentialPropertiesResponseOutput {
	return i.ToAdmCredentialPropertiesResponseOutputWithContext(context.Background())
}

func (i AdmCredentialPropertiesResponseArgs) ToAdmCredentialPropertiesResponseOutputWithContext(ctx context.Context) AdmCredentialPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmCredentialPropertiesResponseOutput)
}

func (i AdmCredentialPropertiesResponseArgs) ToAdmCredentialPropertiesResponsePtrOutput() AdmCredentialPropertiesResponsePtrOutput {
	return i.ToAdmCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AdmCredentialPropertiesResponseArgs) ToAdmCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) AdmCredentialPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmCredentialPropertiesResponseOutput).ToAdmCredentialPropertiesResponsePtrOutputWithContext(ctx)
}









type AdmCredentialPropertiesResponsePtrInput interface {
	pulumi.Input

	ToAdmCredentialPropertiesResponsePtrOutput() AdmCredentialPropertiesResponsePtrOutput
	ToAdmCredentialPropertiesResponsePtrOutputWithContext(context.Context) AdmCredentialPropertiesResponsePtrOutput
}

type admCredentialPropertiesResponsePtrType AdmCredentialPropertiesResponseArgs

func AdmCredentialPropertiesResponsePtr(v *AdmCredentialPropertiesResponseArgs) AdmCredentialPropertiesResponsePtrInput {
	return (*admCredentialPropertiesResponsePtrType)(v)
}

func (*admCredentialPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AdmCredentialPropertiesResponse)(nil)).Elem()
}

func (i *admCredentialPropertiesResponsePtrType) ToAdmCredentialPropertiesResponsePtrOutput() AdmCredentialPropertiesResponsePtrOutput {
	return i.ToAdmCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *admCredentialPropertiesResponsePtrType) ToAdmCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) AdmCredentialPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmCredentialPropertiesResponsePtrOutput)
}

type AdmCredentialPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AdmCredentialPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdmCredentialPropertiesResponse)(nil)).Elem()
}

func (o AdmCredentialPropertiesResponseOutput) ToAdmCredentialPropertiesResponseOutput() AdmCredentialPropertiesResponseOutput {
	return o
}

func (o AdmCredentialPropertiesResponseOutput) ToAdmCredentialPropertiesResponseOutputWithContext(ctx context.Context) AdmCredentialPropertiesResponseOutput {
	return o
}

func (o AdmCredentialPropertiesResponseOutput) ToAdmCredentialPropertiesResponsePtrOutput() AdmCredentialPropertiesResponsePtrOutput {
	return o.ToAdmCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AdmCredentialPropertiesResponseOutput) ToAdmCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) AdmCredentialPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdmCredentialPropertiesResponse) *AdmCredentialPropertiesResponse {
		return &v
	}).(AdmCredentialPropertiesResponsePtrOutput)
}

func (o AdmCredentialPropertiesResponseOutput) AuthTokenUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdmCredentialPropertiesResponse) *string { return v.AuthTokenUrl }).(pulumi.StringPtrOutput)
}

func (o AdmCredentialPropertiesResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdmCredentialPropertiesResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AdmCredentialPropertiesResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdmCredentialPropertiesResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

type AdmCredentialPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AdmCredentialPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdmCredentialPropertiesResponse)(nil)).Elem()
}

func (o AdmCredentialPropertiesResponsePtrOutput) ToAdmCredentialPropertiesResponsePtrOutput() AdmCredentialPropertiesResponsePtrOutput {
	return o
}

func (o AdmCredentialPropertiesResponsePtrOutput) ToAdmCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) AdmCredentialPropertiesResponsePtrOutput {
	return o
}

func (o AdmCredentialPropertiesResponsePtrOutput) Elem() AdmCredentialPropertiesResponseOutput {
	return o.ApplyT(func(v *AdmCredentialPropertiesResponse) AdmCredentialPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AdmCredentialPropertiesResponse
		return ret
	}).(AdmCredentialPropertiesResponseOutput)
}

func (o AdmCredentialPropertiesResponsePtrOutput) AuthTokenUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdmCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthTokenUrl
	}).(pulumi.StringPtrOutput)
}

func (o AdmCredentialPropertiesResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdmCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AdmCredentialPropertiesResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdmCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

type AdmCredentialResponse struct {
	Properties *AdmCredentialPropertiesResponse `pulumi:"properties"`
}





type AdmCredentialResponseInput interface {
	pulumi.Input

	ToAdmCredentialResponseOutput() AdmCredentialResponseOutput
	ToAdmCredentialResponseOutputWithContext(context.Context) AdmCredentialResponseOutput
}

type AdmCredentialResponseArgs struct {
	Properties AdmCredentialPropertiesResponsePtrInput `pulumi:"properties"`
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

func (o AdmCredentialResponseOutput) Properties() AdmCredentialPropertiesResponsePtrOutput {
	return o.ApplyT(func(v AdmCredentialResponse) *AdmCredentialPropertiesResponse { return v.Properties }).(AdmCredentialPropertiesResponsePtrOutput)
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

func (o AdmCredentialResponsePtrOutput) Properties() AdmCredentialPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AdmCredentialResponse) *AdmCredentialPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(AdmCredentialPropertiesResponsePtrOutput)
}

type ApnsCredential struct {
	Properties *ApnsCredentialProperties `pulumi:"properties"`
}





type ApnsCredentialInput interface {
	pulumi.Input

	ToApnsCredentialOutput() ApnsCredentialOutput
	ToApnsCredentialOutputWithContext(context.Context) ApnsCredentialOutput
}

type ApnsCredentialArgs struct {
	Properties ApnsCredentialPropertiesPtrInput `pulumi:"properties"`
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

func (o ApnsCredentialOutput) Properties() ApnsCredentialPropertiesPtrOutput {
	return o.ApplyT(func(v ApnsCredential) *ApnsCredentialProperties { return v.Properties }).(ApnsCredentialPropertiesPtrOutput)
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

func (o ApnsCredentialPtrOutput) Properties() ApnsCredentialPropertiesPtrOutput {
	return o.ApplyT(func(v *ApnsCredential) *ApnsCredentialProperties {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(ApnsCredentialPropertiesPtrOutput)
}

type ApnsCredentialProperties struct {
	ApnsCertificate *string `pulumi:"apnsCertificate"`
	CertificateKey  *string `pulumi:"certificateKey"`
	Endpoint        *string `pulumi:"endpoint"`
	Thumbprint      *string `pulumi:"thumbprint"`
}





type ApnsCredentialPropertiesInput interface {
	pulumi.Input

	ToApnsCredentialPropertiesOutput() ApnsCredentialPropertiesOutput
	ToApnsCredentialPropertiesOutputWithContext(context.Context) ApnsCredentialPropertiesOutput
}

type ApnsCredentialPropertiesArgs struct {
	ApnsCertificate pulumi.StringPtrInput `pulumi:"apnsCertificate"`
	CertificateKey  pulumi.StringPtrInput `pulumi:"certificateKey"`
	Endpoint        pulumi.StringPtrInput `pulumi:"endpoint"`
	Thumbprint      pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (ApnsCredentialPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsCredentialProperties)(nil)).Elem()
}

func (i ApnsCredentialPropertiesArgs) ToApnsCredentialPropertiesOutput() ApnsCredentialPropertiesOutput {
	return i.ToApnsCredentialPropertiesOutputWithContext(context.Background())
}

func (i ApnsCredentialPropertiesArgs) ToApnsCredentialPropertiesOutputWithContext(ctx context.Context) ApnsCredentialPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsCredentialPropertiesOutput)
}

func (i ApnsCredentialPropertiesArgs) ToApnsCredentialPropertiesPtrOutput() ApnsCredentialPropertiesPtrOutput {
	return i.ToApnsCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (i ApnsCredentialPropertiesArgs) ToApnsCredentialPropertiesPtrOutputWithContext(ctx context.Context) ApnsCredentialPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsCredentialPropertiesOutput).ToApnsCredentialPropertiesPtrOutputWithContext(ctx)
}









type ApnsCredentialPropertiesPtrInput interface {
	pulumi.Input

	ToApnsCredentialPropertiesPtrOutput() ApnsCredentialPropertiesPtrOutput
	ToApnsCredentialPropertiesPtrOutputWithContext(context.Context) ApnsCredentialPropertiesPtrOutput
}

type apnsCredentialPropertiesPtrType ApnsCredentialPropertiesArgs

func ApnsCredentialPropertiesPtr(v *ApnsCredentialPropertiesArgs) ApnsCredentialPropertiesPtrInput {
	return (*apnsCredentialPropertiesPtrType)(v)
}

func (*apnsCredentialPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsCredentialProperties)(nil)).Elem()
}

func (i *apnsCredentialPropertiesPtrType) ToApnsCredentialPropertiesPtrOutput() ApnsCredentialPropertiesPtrOutput {
	return i.ToApnsCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (i *apnsCredentialPropertiesPtrType) ToApnsCredentialPropertiesPtrOutputWithContext(ctx context.Context) ApnsCredentialPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsCredentialPropertiesPtrOutput)
}

type ApnsCredentialPropertiesOutput struct{ *pulumi.OutputState }

func (ApnsCredentialPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsCredentialProperties)(nil)).Elem()
}

func (o ApnsCredentialPropertiesOutput) ToApnsCredentialPropertiesOutput() ApnsCredentialPropertiesOutput {
	return o
}

func (o ApnsCredentialPropertiesOutput) ToApnsCredentialPropertiesOutputWithContext(ctx context.Context) ApnsCredentialPropertiesOutput {
	return o
}

func (o ApnsCredentialPropertiesOutput) ToApnsCredentialPropertiesPtrOutput() ApnsCredentialPropertiesPtrOutput {
	return o.ToApnsCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (o ApnsCredentialPropertiesOutput) ToApnsCredentialPropertiesPtrOutputWithContext(ctx context.Context) ApnsCredentialPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApnsCredentialProperties) *ApnsCredentialProperties {
		return &v
	}).(ApnsCredentialPropertiesPtrOutput)
}

func (o ApnsCredentialPropertiesOutput) ApnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredentialProperties) *string { return v.ApnsCertificate }).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPropertiesOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredentialProperties) *string { return v.CertificateKey }).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPropertiesOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredentialProperties) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPropertiesOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredentialProperties) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ApnsCredentialPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ApnsCredentialPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsCredentialProperties)(nil)).Elem()
}

func (o ApnsCredentialPropertiesPtrOutput) ToApnsCredentialPropertiesPtrOutput() ApnsCredentialPropertiesPtrOutput {
	return o
}

func (o ApnsCredentialPropertiesPtrOutput) ToApnsCredentialPropertiesPtrOutputWithContext(ctx context.Context) ApnsCredentialPropertiesPtrOutput {
	return o
}

func (o ApnsCredentialPropertiesPtrOutput) Elem() ApnsCredentialPropertiesOutput {
	return o.ApplyT(func(v *ApnsCredentialProperties) ApnsCredentialProperties {
		if v != nil {
			return *v
		}
		var ret ApnsCredentialProperties
		return ret
	}).(ApnsCredentialPropertiesOutput)
}

func (o ApnsCredentialPropertiesPtrOutput) ApnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApnsCertificate
	}).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPropertiesPtrOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.CertificateKey
	}).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPropertiesPtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.Endpoint
	}).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPropertiesPtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type ApnsCredentialPropertiesResponse struct {
	ApnsCertificate *string `pulumi:"apnsCertificate"`
	CertificateKey  *string `pulumi:"certificateKey"`
	Endpoint        *string `pulumi:"endpoint"`
	Thumbprint      *string `pulumi:"thumbprint"`
}





type ApnsCredentialPropertiesResponseInput interface {
	pulumi.Input

	ToApnsCredentialPropertiesResponseOutput() ApnsCredentialPropertiesResponseOutput
	ToApnsCredentialPropertiesResponseOutputWithContext(context.Context) ApnsCredentialPropertiesResponseOutput
}

type ApnsCredentialPropertiesResponseArgs struct {
	ApnsCertificate pulumi.StringPtrInput `pulumi:"apnsCertificate"`
	CertificateKey  pulumi.StringPtrInput `pulumi:"certificateKey"`
	Endpoint        pulumi.StringPtrInput `pulumi:"endpoint"`
	Thumbprint      pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (ApnsCredentialPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsCredentialPropertiesResponse)(nil)).Elem()
}

func (i ApnsCredentialPropertiesResponseArgs) ToApnsCredentialPropertiesResponseOutput() ApnsCredentialPropertiesResponseOutput {
	return i.ToApnsCredentialPropertiesResponseOutputWithContext(context.Background())
}

func (i ApnsCredentialPropertiesResponseArgs) ToApnsCredentialPropertiesResponseOutputWithContext(ctx context.Context) ApnsCredentialPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsCredentialPropertiesResponseOutput)
}

func (i ApnsCredentialPropertiesResponseArgs) ToApnsCredentialPropertiesResponsePtrOutput() ApnsCredentialPropertiesResponsePtrOutput {
	return i.ToApnsCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ApnsCredentialPropertiesResponseArgs) ToApnsCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) ApnsCredentialPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsCredentialPropertiesResponseOutput).ToApnsCredentialPropertiesResponsePtrOutputWithContext(ctx)
}









type ApnsCredentialPropertiesResponsePtrInput interface {
	pulumi.Input

	ToApnsCredentialPropertiesResponsePtrOutput() ApnsCredentialPropertiesResponsePtrOutput
	ToApnsCredentialPropertiesResponsePtrOutputWithContext(context.Context) ApnsCredentialPropertiesResponsePtrOutput
}

type apnsCredentialPropertiesResponsePtrType ApnsCredentialPropertiesResponseArgs

func ApnsCredentialPropertiesResponsePtr(v *ApnsCredentialPropertiesResponseArgs) ApnsCredentialPropertiesResponsePtrInput {
	return (*apnsCredentialPropertiesResponsePtrType)(v)
}

func (*apnsCredentialPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsCredentialPropertiesResponse)(nil)).Elem()
}

func (i *apnsCredentialPropertiesResponsePtrType) ToApnsCredentialPropertiesResponsePtrOutput() ApnsCredentialPropertiesResponsePtrOutput {
	return i.ToApnsCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *apnsCredentialPropertiesResponsePtrType) ToApnsCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) ApnsCredentialPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsCredentialPropertiesResponsePtrOutput)
}

type ApnsCredentialPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ApnsCredentialPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsCredentialPropertiesResponse)(nil)).Elem()
}

func (o ApnsCredentialPropertiesResponseOutput) ToApnsCredentialPropertiesResponseOutput() ApnsCredentialPropertiesResponseOutput {
	return o
}

func (o ApnsCredentialPropertiesResponseOutput) ToApnsCredentialPropertiesResponseOutputWithContext(ctx context.Context) ApnsCredentialPropertiesResponseOutput {
	return o
}

func (o ApnsCredentialPropertiesResponseOutput) ToApnsCredentialPropertiesResponsePtrOutput() ApnsCredentialPropertiesResponsePtrOutput {
	return o.ToApnsCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ApnsCredentialPropertiesResponseOutput) ToApnsCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) ApnsCredentialPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApnsCredentialPropertiesResponse) *ApnsCredentialPropertiesResponse {
		return &v
	}).(ApnsCredentialPropertiesResponsePtrOutput)
}

func (o ApnsCredentialPropertiesResponseOutput) ApnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredentialPropertiesResponse) *string { return v.ApnsCertificate }).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPropertiesResponseOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredentialPropertiesResponse) *string { return v.CertificateKey }).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPropertiesResponseOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredentialPropertiesResponse) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPropertiesResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApnsCredentialPropertiesResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ApnsCredentialPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ApnsCredentialPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsCredentialPropertiesResponse)(nil)).Elem()
}

func (o ApnsCredentialPropertiesResponsePtrOutput) ToApnsCredentialPropertiesResponsePtrOutput() ApnsCredentialPropertiesResponsePtrOutput {
	return o
}

func (o ApnsCredentialPropertiesResponsePtrOutput) ToApnsCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) ApnsCredentialPropertiesResponsePtrOutput {
	return o
}

func (o ApnsCredentialPropertiesResponsePtrOutput) Elem() ApnsCredentialPropertiesResponseOutput {
	return o.ApplyT(func(v *ApnsCredentialPropertiesResponse) ApnsCredentialPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ApnsCredentialPropertiesResponse
		return ret
	}).(ApnsCredentialPropertiesResponseOutput)
}

func (o ApnsCredentialPropertiesResponsePtrOutput) ApnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApnsCertificate
	}).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPropertiesResponsePtrOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CertificateKey
	}).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPropertiesResponsePtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Endpoint
	}).(pulumi.StringPtrOutput)
}

func (o ApnsCredentialPropertiesResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type ApnsCredentialResponse struct {
	Properties *ApnsCredentialPropertiesResponse `pulumi:"properties"`
}





type ApnsCredentialResponseInput interface {
	pulumi.Input

	ToApnsCredentialResponseOutput() ApnsCredentialResponseOutput
	ToApnsCredentialResponseOutputWithContext(context.Context) ApnsCredentialResponseOutput
}

type ApnsCredentialResponseArgs struct {
	Properties ApnsCredentialPropertiesResponsePtrInput `pulumi:"properties"`
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

func (o ApnsCredentialResponseOutput) Properties() ApnsCredentialPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ApnsCredentialResponse) *ApnsCredentialPropertiesResponse { return v.Properties }).(ApnsCredentialPropertiesResponsePtrOutput)
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

func (o ApnsCredentialResponsePtrOutput) Properties() ApnsCredentialPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ApnsCredentialResponse) *ApnsCredentialPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(ApnsCredentialPropertiesResponsePtrOutput)
}

type BaiduCredential struct {
	Properties *BaiduCredentialProperties `pulumi:"properties"`
}





type BaiduCredentialInput interface {
	pulumi.Input

	ToBaiduCredentialOutput() BaiduCredentialOutput
	ToBaiduCredentialOutputWithContext(context.Context) BaiduCredentialOutput
}

type BaiduCredentialArgs struct {
	Properties BaiduCredentialPropertiesPtrInput `pulumi:"properties"`
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

func (o BaiduCredentialOutput) Properties() BaiduCredentialPropertiesPtrOutput {
	return o.ApplyT(func(v BaiduCredential) *BaiduCredentialProperties { return v.Properties }).(BaiduCredentialPropertiesPtrOutput)
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

func (o BaiduCredentialPtrOutput) Properties() BaiduCredentialPropertiesPtrOutput {
	return o.ApplyT(func(v *BaiduCredential) *BaiduCredentialProperties {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(BaiduCredentialPropertiesPtrOutput)
}

type BaiduCredentialProperties struct {
	BaiduApiKey    *string `pulumi:"baiduApiKey"`
	BaiduEndPoint  *string `pulumi:"baiduEndPoint"`
	BaiduSecretKey *string `pulumi:"baiduSecretKey"`
}





type BaiduCredentialPropertiesInput interface {
	pulumi.Input

	ToBaiduCredentialPropertiesOutput() BaiduCredentialPropertiesOutput
	ToBaiduCredentialPropertiesOutputWithContext(context.Context) BaiduCredentialPropertiesOutput
}

type BaiduCredentialPropertiesArgs struct {
	BaiduApiKey    pulumi.StringPtrInput `pulumi:"baiduApiKey"`
	BaiduEndPoint  pulumi.StringPtrInput `pulumi:"baiduEndPoint"`
	BaiduSecretKey pulumi.StringPtrInput `pulumi:"baiduSecretKey"`
}

func (BaiduCredentialPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BaiduCredentialProperties)(nil)).Elem()
}

func (i BaiduCredentialPropertiesArgs) ToBaiduCredentialPropertiesOutput() BaiduCredentialPropertiesOutput {
	return i.ToBaiduCredentialPropertiesOutputWithContext(context.Background())
}

func (i BaiduCredentialPropertiesArgs) ToBaiduCredentialPropertiesOutputWithContext(ctx context.Context) BaiduCredentialPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduCredentialPropertiesOutput)
}

func (i BaiduCredentialPropertiesArgs) ToBaiduCredentialPropertiesPtrOutput() BaiduCredentialPropertiesPtrOutput {
	return i.ToBaiduCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (i BaiduCredentialPropertiesArgs) ToBaiduCredentialPropertiesPtrOutputWithContext(ctx context.Context) BaiduCredentialPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduCredentialPropertiesOutput).ToBaiduCredentialPropertiesPtrOutputWithContext(ctx)
}









type BaiduCredentialPropertiesPtrInput interface {
	pulumi.Input

	ToBaiduCredentialPropertiesPtrOutput() BaiduCredentialPropertiesPtrOutput
	ToBaiduCredentialPropertiesPtrOutputWithContext(context.Context) BaiduCredentialPropertiesPtrOutput
}

type baiduCredentialPropertiesPtrType BaiduCredentialPropertiesArgs

func BaiduCredentialPropertiesPtr(v *BaiduCredentialPropertiesArgs) BaiduCredentialPropertiesPtrInput {
	return (*baiduCredentialPropertiesPtrType)(v)
}

func (*baiduCredentialPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BaiduCredentialProperties)(nil)).Elem()
}

func (i *baiduCredentialPropertiesPtrType) ToBaiduCredentialPropertiesPtrOutput() BaiduCredentialPropertiesPtrOutput {
	return i.ToBaiduCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (i *baiduCredentialPropertiesPtrType) ToBaiduCredentialPropertiesPtrOutputWithContext(ctx context.Context) BaiduCredentialPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduCredentialPropertiesPtrOutput)
}

type BaiduCredentialPropertiesOutput struct{ *pulumi.OutputState }

func (BaiduCredentialPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BaiduCredentialProperties)(nil)).Elem()
}

func (o BaiduCredentialPropertiesOutput) ToBaiduCredentialPropertiesOutput() BaiduCredentialPropertiesOutput {
	return o
}

func (o BaiduCredentialPropertiesOutput) ToBaiduCredentialPropertiesOutputWithContext(ctx context.Context) BaiduCredentialPropertiesOutput {
	return o
}

func (o BaiduCredentialPropertiesOutput) ToBaiduCredentialPropertiesPtrOutput() BaiduCredentialPropertiesPtrOutput {
	return o.ToBaiduCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (o BaiduCredentialPropertiesOutput) ToBaiduCredentialPropertiesPtrOutputWithContext(ctx context.Context) BaiduCredentialPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BaiduCredentialProperties) *BaiduCredentialProperties {
		return &v
	}).(BaiduCredentialPropertiesPtrOutput)
}

func (o BaiduCredentialPropertiesOutput) BaiduApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaiduCredentialProperties) *string { return v.BaiduApiKey }).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialPropertiesOutput) BaiduEndPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaiduCredentialProperties) *string { return v.BaiduEndPoint }).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialPropertiesOutput) BaiduSecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaiduCredentialProperties) *string { return v.BaiduSecretKey }).(pulumi.StringPtrOutput)
}

type BaiduCredentialPropertiesPtrOutput struct{ *pulumi.OutputState }

func (BaiduCredentialPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BaiduCredentialProperties)(nil)).Elem()
}

func (o BaiduCredentialPropertiesPtrOutput) ToBaiduCredentialPropertiesPtrOutput() BaiduCredentialPropertiesPtrOutput {
	return o
}

func (o BaiduCredentialPropertiesPtrOutput) ToBaiduCredentialPropertiesPtrOutputWithContext(ctx context.Context) BaiduCredentialPropertiesPtrOutput {
	return o
}

func (o BaiduCredentialPropertiesPtrOutput) Elem() BaiduCredentialPropertiesOutput {
	return o.ApplyT(func(v *BaiduCredentialProperties) BaiduCredentialProperties {
		if v != nil {
			return *v
		}
		var ret BaiduCredentialProperties
		return ret
	}).(BaiduCredentialPropertiesOutput)
}

func (o BaiduCredentialPropertiesPtrOutput) BaiduApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaiduCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.BaiduApiKey
	}).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialPropertiesPtrOutput) BaiduEndPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaiduCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.BaiduEndPoint
	}).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialPropertiesPtrOutput) BaiduSecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaiduCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.BaiduSecretKey
	}).(pulumi.StringPtrOutput)
}

type BaiduCredentialPropertiesResponse struct {
	BaiduApiKey    *string `pulumi:"baiduApiKey"`
	BaiduEndPoint  *string `pulumi:"baiduEndPoint"`
	BaiduSecretKey *string `pulumi:"baiduSecretKey"`
}





type BaiduCredentialPropertiesResponseInput interface {
	pulumi.Input

	ToBaiduCredentialPropertiesResponseOutput() BaiduCredentialPropertiesResponseOutput
	ToBaiduCredentialPropertiesResponseOutputWithContext(context.Context) BaiduCredentialPropertiesResponseOutput
}

type BaiduCredentialPropertiesResponseArgs struct {
	BaiduApiKey    pulumi.StringPtrInput `pulumi:"baiduApiKey"`
	BaiduEndPoint  pulumi.StringPtrInput `pulumi:"baiduEndPoint"`
	BaiduSecretKey pulumi.StringPtrInput `pulumi:"baiduSecretKey"`
}

func (BaiduCredentialPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BaiduCredentialPropertiesResponse)(nil)).Elem()
}

func (i BaiduCredentialPropertiesResponseArgs) ToBaiduCredentialPropertiesResponseOutput() BaiduCredentialPropertiesResponseOutput {
	return i.ToBaiduCredentialPropertiesResponseOutputWithContext(context.Background())
}

func (i BaiduCredentialPropertiesResponseArgs) ToBaiduCredentialPropertiesResponseOutputWithContext(ctx context.Context) BaiduCredentialPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduCredentialPropertiesResponseOutput)
}

func (i BaiduCredentialPropertiesResponseArgs) ToBaiduCredentialPropertiesResponsePtrOutput() BaiduCredentialPropertiesResponsePtrOutput {
	return i.ToBaiduCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i BaiduCredentialPropertiesResponseArgs) ToBaiduCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) BaiduCredentialPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduCredentialPropertiesResponseOutput).ToBaiduCredentialPropertiesResponsePtrOutputWithContext(ctx)
}









type BaiduCredentialPropertiesResponsePtrInput interface {
	pulumi.Input

	ToBaiduCredentialPropertiesResponsePtrOutput() BaiduCredentialPropertiesResponsePtrOutput
	ToBaiduCredentialPropertiesResponsePtrOutputWithContext(context.Context) BaiduCredentialPropertiesResponsePtrOutput
}

type baiduCredentialPropertiesResponsePtrType BaiduCredentialPropertiesResponseArgs

func BaiduCredentialPropertiesResponsePtr(v *BaiduCredentialPropertiesResponseArgs) BaiduCredentialPropertiesResponsePtrInput {
	return (*baiduCredentialPropertiesResponsePtrType)(v)
}

func (*baiduCredentialPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BaiduCredentialPropertiesResponse)(nil)).Elem()
}

func (i *baiduCredentialPropertiesResponsePtrType) ToBaiduCredentialPropertiesResponsePtrOutput() BaiduCredentialPropertiesResponsePtrOutput {
	return i.ToBaiduCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *baiduCredentialPropertiesResponsePtrType) ToBaiduCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) BaiduCredentialPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduCredentialPropertiesResponsePtrOutput)
}

type BaiduCredentialPropertiesResponseOutput struct{ *pulumi.OutputState }

func (BaiduCredentialPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BaiduCredentialPropertiesResponse)(nil)).Elem()
}

func (o BaiduCredentialPropertiesResponseOutput) ToBaiduCredentialPropertiesResponseOutput() BaiduCredentialPropertiesResponseOutput {
	return o
}

func (o BaiduCredentialPropertiesResponseOutput) ToBaiduCredentialPropertiesResponseOutputWithContext(ctx context.Context) BaiduCredentialPropertiesResponseOutput {
	return o
}

func (o BaiduCredentialPropertiesResponseOutput) ToBaiduCredentialPropertiesResponsePtrOutput() BaiduCredentialPropertiesResponsePtrOutput {
	return o.ToBaiduCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o BaiduCredentialPropertiesResponseOutput) ToBaiduCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) BaiduCredentialPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BaiduCredentialPropertiesResponse) *BaiduCredentialPropertiesResponse {
		return &v
	}).(BaiduCredentialPropertiesResponsePtrOutput)
}

func (o BaiduCredentialPropertiesResponseOutput) BaiduApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaiduCredentialPropertiesResponse) *string { return v.BaiduApiKey }).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialPropertiesResponseOutput) BaiduEndPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaiduCredentialPropertiesResponse) *string { return v.BaiduEndPoint }).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialPropertiesResponseOutput) BaiduSecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaiduCredentialPropertiesResponse) *string { return v.BaiduSecretKey }).(pulumi.StringPtrOutput)
}

type BaiduCredentialPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (BaiduCredentialPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BaiduCredentialPropertiesResponse)(nil)).Elem()
}

func (o BaiduCredentialPropertiesResponsePtrOutput) ToBaiduCredentialPropertiesResponsePtrOutput() BaiduCredentialPropertiesResponsePtrOutput {
	return o
}

func (o BaiduCredentialPropertiesResponsePtrOutput) ToBaiduCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) BaiduCredentialPropertiesResponsePtrOutput {
	return o
}

func (o BaiduCredentialPropertiesResponsePtrOutput) Elem() BaiduCredentialPropertiesResponseOutput {
	return o.ApplyT(func(v *BaiduCredentialPropertiesResponse) BaiduCredentialPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret BaiduCredentialPropertiesResponse
		return ret
	}).(BaiduCredentialPropertiesResponseOutput)
}

func (o BaiduCredentialPropertiesResponsePtrOutput) BaiduApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaiduCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.BaiduApiKey
	}).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialPropertiesResponsePtrOutput) BaiduEndPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaiduCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.BaiduEndPoint
	}).(pulumi.StringPtrOutput)
}

func (o BaiduCredentialPropertiesResponsePtrOutput) BaiduSecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BaiduCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.BaiduSecretKey
	}).(pulumi.StringPtrOutput)
}

type BaiduCredentialResponse struct {
	Properties *BaiduCredentialPropertiesResponse `pulumi:"properties"`
}





type BaiduCredentialResponseInput interface {
	pulumi.Input

	ToBaiduCredentialResponseOutput() BaiduCredentialResponseOutput
	ToBaiduCredentialResponseOutputWithContext(context.Context) BaiduCredentialResponseOutput
}

type BaiduCredentialResponseArgs struct {
	Properties BaiduCredentialPropertiesResponsePtrInput `pulumi:"properties"`
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

func (o BaiduCredentialResponseOutput) Properties() BaiduCredentialPropertiesResponsePtrOutput {
	return o.ApplyT(func(v BaiduCredentialResponse) *BaiduCredentialPropertiesResponse { return v.Properties }).(BaiduCredentialPropertiesResponsePtrOutput)
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

func (o BaiduCredentialResponsePtrOutput) Properties() BaiduCredentialPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *BaiduCredentialResponse) *BaiduCredentialPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(BaiduCredentialPropertiesResponsePtrOutput)
}

type GcmCredential struct {
	Properties *GcmCredentialProperties `pulumi:"properties"`
}





type GcmCredentialInput interface {
	pulumi.Input

	ToGcmCredentialOutput() GcmCredentialOutput
	ToGcmCredentialOutputWithContext(context.Context) GcmCredentialOutput
}

type GcmCredentialArgs struct {
	Properties GcmCredentialPropertiesPtrInput `pulumi:"properties"`
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

func (o GcmCredentialOutput) Properties() GcmCredentialPropertiesPtrOutput {
	return o.ApplyT(func(v GcmCredential) *GcmCredentialProperties { return v.Properties }).(GcmCredentialPropertiesPtrOutput)
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

func (o GcmCredentialPtrOutput) Properties() GcmCredentialPropertiesPtrOutput {
	return o.ApplyT(func(v *GcmCredential) *GcmCredentialProperties {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(GcmCredentialPropertiesPtrOutput)
}

type GcmCredentialProperties struct {
	GcmEndpoint  *string `pulumi:"gcmEndpoint"`
	GoogleApiKey *string `pulumi:"googleApiKey"`
}





type GcmCredentialPropertiesInput interface {
	pulumi.Input

	ToGcmCredentialPropertiesOutput() GcmCredentialPropertiesOutput
	ToGcmCredentialPropertiesOutputWithContext(context.Context) GcmCredentialPropertiesOutput
}

type GcmCredentialPropertiesArgs struct {
	GcmEndpoint  pulumi.StringPtrInput `pulumi:"gcmEndpoint"`
	GoogleApiKey pulumi.StringPtrInput `pulumi:"googleApiKey"`
}

func (GcmCredentialPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GcmCredentialProperties)(nil)).Elem()
}

func (i GcmCredentialPropertiesArgs) ToGcmCredentialPropertiesOutput() GcmCredentialPropertiesOutput {
	return i.ToGcmCredentialPropertiesOutputWithContext(context.Background())
}

func (i GcmCredentialPropertiesArgs) ToGcmCredentialPropertiesOutputWithContext(ctx context.Context) GcmCredentialPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmCredentialPropertiesOutput)
}

func (i GcmCredentialPropertiesArgs) ToGcmCredentialPropertiesPtrOutput() GcmCredentialPropertiesPtrOutput {
	return i.ToGcmCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (i GcmCredentialPropertiesArgs) ToGcmCredentialPropertiesPtrOutputWithContext(ctx context.Context) GcmCredentialPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmCredentialPropertiesOutput).ToGcmCredentialPropertiesPtrOutputWithContext(ctx)
}









type GcmCredentialPropertiesPtrInput interface {
	pulumi.Input

	ToGcmCredentialPropertiesPtrOutput() GcmCredentialPropertiesPtrOutput
	ToGcmCredentialPropertiesPtrOutputWithContext(context.Context) GcmCredentialPropertiesPtrOutput
}

type gcmCredentialPropertiesPtrType GcmCredentialPropertiesArgs

func GcmCredentialPropertiesPtr(v *GcmCredentialPropertiesArgs) GcmCredentialPropertiesPtrInput {
	return (*gcmCredentialPropertiesPtrType)(v)
}

func (*gcmCredentialPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GcmCredentialProperties)(nil)).Elem()
}

func (i *gcmCredentialPropertiesPtrType) ToGcmCredentialPropertiesPtrOutput() GcmCredentialPropertiesPtrOutput {
	return i.ToGcmCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (i *gcmCredentialPropertiesPtrType) ToGcmCredentialPropertiesPtrOutputWithContext(ctx context.Context) GcmCredentialPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmCredentialPropertiesPtrOutput)
}

type GcmCredentialPropertiesOutput struct{ *pulumi.OutputState }

func (GcmCredentialPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GcmCredentialProperties)(nil)).Elem()
}

func (o GcmCredentialPropertiesOutput) ToGcmCredentialPropertiesOutput() GcmCredentialPropertiesOutput {
	return o
}

func (o GcmCredentialPropertiesOutput) ToGcmCredentialPropertiesOutputWithContext(ctx context.Context) GcmCredentialPropertiesOutput {
	return o
}

func (o GcmCredentialPropertiesOutput) ToGcmCredentialPropertiesPtrOutput() GcmCredentialPropertiesPtrOutput {
	return o.ToGcmCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (o GcmCredentialPropertiesOutput) ToGcmCredentialPropertiesPtrOutputWithContext(ctx context.Context) GcmCredentialPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GcmCredentialProperties) *GcmCredentialProperties {
		return &v
	}).(GcmCredentialPropertiesPtrOutput)
}

func (o GcmCredentialPropertiesOutput) GcmEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GcmCredentialProperties) *string { return v.GcmEndpoint }).(pulumi.StringPtrOutput)
}

func (o GcmCredentialPropertiesOutput) GoogleApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GcmCredentialProperties) *string { return v.GoogleApiKey }).(pulumi.StringPtrOutput)
}

type GcmCredentialPropertiesPtrOutput struct{ *pulumi.OutputState }

func (GcmCredentialPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GcmCredentialProperties)(nil)).Elem()
}

func (o GcmCredentialPropertiesPtrOutput) ToGcmCredentialPropertiesPtrOutput() GcmCredentialPropertiesPtrOutput {
	return o
}

func (o GcmCredentialPropertiesPtrOutput) ToGcmCredentialPropertiesPtrOutputWithContext(ctx context.Context) GcmCredentialPropertiesPtrOutput {
	return o
}

func (o GcmCredentialPropertiesPtrOutput) Elem() GcmCredentialPropertiesOutput {
	return o.ApplyT(func(v *GcmCredentialProperties) GcmCredentialProperties {
		if v != nil {
			return *v
		}
		var ret GcmCredentialProperties
		return ret
	}).(GcmCredentialPropertiesOutput)
}

func (o GcmCredentialPropertiesPtrOutput) GcmEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GcmCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.GcmEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o GcmCredentialPropertiesPtrOutput) GoogleApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GcmCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.GoogleApiKey
	}).(pulumi.StringPtrOutput)
}

type GcmCredentialPropertiesResponse struct {
	GcmEndpoint  *string `pulumi:"gcmEndpoint"`
	GoogleApiKey *string `pulumi:"googleApiKey"`
}





type GcmCredentialPropertiesResponseInput interface {
	pulumi.Input

	ToGcmCredentialPropertiesResponseOutput() GcmCredentialPropertiesResponseOutput
	ToGcmCredentialPropertiesResponseOutputWithContext(context.Context) GcmCredentialPropertiesResponseOutput
}

type GcmCredentialPropertiesResponseArgs struct {
	GcmEndpoint  pulumi.StringPtrInput `pulumi:"gcmEndpoint"`
	GoogleApiKey pulumi.StringPtrInput `pulumi:"googleApiKey"`
}

func (GcmCredentialPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GcmCredentialPropertiesResponse)(nil)).Elem()
}

func (i GcmCredentialPropertiesResponseArgs) ToGcmCredentialPropertiesResponseOutput() GcmCredentialPropertiesResponseOutput {
	return i.ToGcmCredentialPropertiesResponseOutputWithContext(context.Background())
}

func (i GcmCredentialPropertiesResponseArgs) ToGcmCredentialPropertiesResponseOutputWithContext(ctx context.Context) GcmCredentialPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmCredentialPropertiesResponseOutput)
}

func (i GcmCredentialPropertiesResponseArgs) ToGcmCredentialPropertiesResponsePtrOutput() GcmCredentialPropertiesResponsePtrOutput {
	return i.ToGcmCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i GcmCredentialPropertiesResponseArgs) ToGcmCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) GcmCredentialPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmCredentialPropertiesResponseOutput).ToGcmCredentialPropertiesResponsePtrOutputWithContext(ctx)
}









type GcmCredentialPropertiesResponsePtrInput interface {
	pulumi.Input

	ToGcmCredentialPropertiesResponsePtrOutput() GcmCredentialPropertiesResponsePtrOutput
	ToGcmCredentialPropertiesResponsePtrOutputWithContext(context.Context) GcmCredentialPropertiesResponsePtrOutput
}

type gcmCredentialPropertiesResponsePtrType GcmCredentialPropertiesResponseArgs

func GcmCredentialPropertiesResponsePtr(v *GcmCredentialPropertiesResponseArgs) GcmCredentialPropertiesResponsePtrInput {
	return (*gcmCredentialPropertiesResponsePtrType)(v)
}

func (*gcmCredentialPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GcmCredentialPropertiesResponse)(nil)).Elem()
}

func (i *gcmCredentialPropertiesResponsePtrType) ToGcmCredentialPropertiesResponsePtrOutput() GcmCredentialPropertiesResponsePtrOutput {
	return i.ToGcmCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *gcmCredentialPropertiesResponsePtrType) ToGcmCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) GcmCredentialPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmCredentialPropertiesResponsePtrOutput)
}

type GcmCredentialPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GcmCredentialPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GcmCredentialPropertiesResponse)(nil)).Elem()
}

func (o GcmCredentialPropertiesResponseOutput) ToGcmCredentialPropertiesResponseOutput() GcmCredentialPropertiesResponseOutput {
	return o
}

func (o GcmCredentialPropertiesResponseOutput) ToGcmCredentialPropertiesResponseOutputWithContext(ctx context.Context) GcmCredentialPropertiesResponseOutput {
	return o
}

func (o GcmCredentialPropertiesResponseOutput) ToGcmCredentialPropertiesResponsePtrOutput() GcmCredentialPropertiesResponsePtrOutput {
	return o.ToGcmCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o GcmCredentialPropertiesResponseOutput) ToGcmCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) GcmCredentialPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GcmCredentialPropertiesResponse) *GcmCredentialPropertiesResponse {
		return &v
	}).(GcmCredentialPropertiesResponsePtrOutput)
}

func (o GcmCredentialPropertiesResponseOutput) GcmEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GcmCredentialPropertiesResponse) *string { return v.GcmEndpoint }).(pulumi.StringPtrOutput)
}

func (o GcmCredentialPropertiesResponseOutput) GoogleApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GcmCredentialPropertiesResponse) *string { return v.GoogleApiKey }).(pulumi.StringPtrOutput)
}

type GcmCredentialPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (GcmCredentialPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GcmCredentialPropertiesResponse)(nil)).Elem()
}

func (o GcmCredentialPropertiesResponsePtrOutput) ToGcmCredentialPropertiesResponsePtrOutput() GcmCredentialPropertiesResponsePtrOutput {
	return o
}

func (o GcmCredentialPropertiesResponsePtrOutput) ToGcmCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) GcmCredentialPropertiesResponsePtrOutput {
	return o
}

func (o GcmCredentialPropertiesResponsePtrOutput) Elem() GcmCredentialPropertiesResponseOutput {
	return o.ApplyT(func(v *GcmCredentialPropertiesResponse) GcmCredentialPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret GcmCredentialPropertiesResponse
		return ret
	}).(GcmCredentialPropertiesResponseOutput)
}

func (o GcmCredentialPropertiesResponsePtrOutput) GcmEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GcmCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.GcmEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o GcmCredentialPropertiesResponsePtrOutput) GoogleApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GcmCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.GoogleApiKey
	}).(pulumi.StringPtrOutput)
}

type GcmCredentialResponse struct {
	Properties *GcmCredentialPropertiesResponse `pulumi:"properties"`
}





type GcmCredentialResponseInput interface {
	pulumi.Input

	ToGcmCredentialResponseOutput() GcmCredentialResponseOutput
	ToGcmCredentialResponseOutputWithContext(context.Context) GcmCredentialResponseOutput
}

type GcmCredentialResponseArgs struct {
	Properties GcmCredentialPropertiesResponsePtrInput `pulumi:"properties"`
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

func (o GcmCredentialResponseOutput) Properties() GcmCredentialPropertiesResponsePtrOutput {
	return o.ApplyT(func(v GcmCredentialResponse) *GcmCredentialPropertiesResponse { return v.Properties }).(GcmCredentialPropertiesResponsePtrOutput)
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

func (o GcmCredentialResponsePtrOutput) Properties() GcmCredentialPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *GcmCredentialResponse) *GcmCredentialPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(GcmCredentialPropertiesResponsePtrOutput)
}

type MpnsCredential struct {
	Properties *MpnsCredentialProperties `pulumi:"properties"`
}





type MpnsCredentialInput interface {
	pulumi.Input

	ToMpnsCredentialOutput() MpnsCredentialOutput
	ToMpnsCredentialOutputWithContext(context.Context) MpnsCredentialOutput
}

type MpnsCredentialArgs struct {
	Properties MpnsCredentialPropertiesPtrInput `pulumi:"properties"`
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

func (o MpnsCredentialOutput) Properties() MpnsCredentialPropertiesPtrOutput {
	return o.ApplyT(func(v MpnsCredential) *MpnsCredentialProperties { return v.Properties }).(MpnsCredentialPropertiesPtrOutput)
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

func (o MpnsCredentialPtrOutput) Properties() MpnsCredentialPropertiesPtrOutput {
	return o.ApplyT(func(v *MpnsCredential) *MpnsCredentialProperties {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(MpnsCredentialPropertiesPtrOutput)
}

type MpnsCredentialProperties struct {
	CertificateKey  *string `pulumi:"certificateKey"`
	MpnsCertificate *string `pulumi:"mpnsCertificate"`
	Thumbprint      *string `pulumi:"thumbprint"`
}





type MpnsCredentialPropertiesInput interface {
	pulumi.Input

	ToMpnsCredentialPropertiesOutput() MpnsCredentialPropertiesOutput
	ToMpnsCredentialPropertiesOutputWithContext(context.Context) MpnsCredentialPropertiesOutput
}

type MpnsCredentialPropertiesArgs struct {
	CertificateKey  pulumi.StringPtrInput `pulumi:"certificateKey"`
	MpnsCertificate pulumi.StringPtrInput `pulumi:"mpnsCertificate"`
	Thumbprint      pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (MpnsCredentialPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MpnsCredentialProperties)(nil)).Elem()
}

func (i MpnsCredentialPropertiesArgs) ToMpnsCredentialPropertiesOutput() MpnsCredentialPropertiesOutput {
	return i.ToMpnsCredentialPropertiesOutputWithContext(context.Background())
}

func (i MpnsCredentialPropertiesArgs) ToMpnsCredentialPropertiesOutputWithContext(ctx context.Context) MpnsCredentialPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpnsCredentialPropertiesOutput)
}

func (i MpnsCredentialPropertiesArgs) ToMpnsCredentialPropertiesPtrOutput() MpnsCredentialPropertiesPtrOutput {
	return i.ToMpnsCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (i MpnsCredentialPropertiesArgs) ToMpnsCredentialPropertiesPtrOutputWithContext(ctx context.Context) MpnsCredentialPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpnsCredentialPropertiesOutput).ToMpnsCredentialPropertiesPtrOutputWithContext(ctx)
}









type MpnsCredentialPropertiesPtrInput interface {
	pulumi.Input

	ToMpnsCredentialPropertiesPtrOutput() MpnsCredentialPropertiesPtrOutput
	ToMpnsCredentialPropertiesPtrOutputWithContext(context.Context) MpnsCredentialPropertiesPtrOutput
}

type mpnsCredentialPropertiesPtrType MpnsCredentialPropertiesArgs

func MpnsCredentialPropertiesPtr(v *MpnsCredentialPropertiesArgs) MpnsCredentialPropertiesPtrInput {
	return (*mpnsCredentialPropertiesPtrType)(v)
}

func (*mpnsCredentialPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MpnsCredentialProperties)(nil)).Elem()
}

func (i *mpnsCredentialPropertiesPtrType) ToMpnsCredentialPropertiesPtrOutput() MpnsCredentialPropertiesPtrOutput {
	return i.ToMpnsCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (i *mpnsCredentialPropertiesPtrType) ToMpnsCredentialPropertiesPtrOutputWithContext(ctx context.Context) MpnsCredentialPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpnsCredentialPropertiesPtrOutput)
}

type MpnsCredentialPropertiesOutput struct{ *pulumi.OutputState }

func (MpnsCredentialPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MpnsCredentialProperties)(nil)).Elem()
}

func (o MpnsCredentialPropertiesOutput) ToMpnsCredentialPropertiesOutput() MpnsCredentialPropertiesOutput {
	return o
}

func (o MpnsCredentialPropertiesOutput) ToMpnsCredentialPropertiesOutputWithContext(ctx context.Context) MpnsCredentialPropertiesOutput {
	return o
}

func (o MpnsCredentialPropertiesOutput) ToMpnsCredentialPropertiesPtrOutput() MpnsCredentialPropertiesPtrOutput {
	return o.ToMpnsCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (o MpnsCredentialPropertiesOutput) ToMpnsCredentialPropertiesPtrOutputWithContext(ctx context.Context) MpnsCredentialPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MpnsCredentialProperties) *MpnsCredentialProperties {
		return &v
	}).(MpnsCredentialPropertiesPtrOutput)
}

func (o MpnsCredentialPropertiesOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MpnsCredentialProperties) *string { return v.CertificateKey }).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialPropertiesOutput) MpnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MpnsCredentialProperties) *string { return v.MpnsCertificate }).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialPropertiesOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MpnsCredentialProperties) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type MpnsCredentialPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MpnsCredentialPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MpnsCredentialProperties)(nil)).Elem()
}

func (o MpnsCredentialPropertiesPtrOutput) ToMpnsCredentialPropertiesPtrOutput() MpnsCredentialPropertiesPtrOutput {
	return o
}

func (o MpnsCredentialPropertiesPtrOutput) ToMpnsCredentialPropertiesPtrOutputWithContext(ctx context.Context) MpnsCredentialPropertiesPtrOutput {
	return o
}

func (o MpnsCredentialPropertiesPtrOutput) Elem() MpnsCredentialPropertiesOutput {
	return o.ApplyT(func(v *MpnsCredentialProperties) MpnsCredentialProperties {
		if v != nil {
			return *v
		}
		var ret MpnsCredentialProperties
		return ret
	}).(MpnsCredentialPropertiesOutput)
}

func (o MpnsCredentialPropertiesPtrOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MpnsCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.CertificateKey
	}).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialPropertiesPtrOutput) MpnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MpnsCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.MpnsCertificate
	}).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialPropertiesPtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MpnsCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type MpnsCredentialPropertiesResponse struct {
	CertificateKey  *string `pulumi:"certificateKey"`
	MpnsCertificate *string `pulumi:"mpnsCertificate"`
	Thumbprint      *string `pulumi:"thumbprint"`
}





type MpnsCredentialPropertiesResponseInput interface {
	pulumi.Input

	ToMpnsCredentialPropertiesResponseOutput() MpnsCredentialPropertiesResponseOutput
	ToMpnsCredentialPropertiesResponseOutputWithContext(context.Context) MpnsCredentialPropertiesResponseOutput
}

type MpnsCredentialPropertiesResponseArgs struct {
	CertificateKey  pulumi.StringPtrInput `pulumi:"certificateKey"`
	MpnsCertificate pulumi.StringPtrInput `pulumi:"mpnsCertificate"`
	Thumbprint      pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (MpnsCredentialPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MpnsCredentialPropertiesResponse)(nil)).Elem()
}

func (i MpnsCredentialPropertiesResponseArgs) ToMpnsCredentialPropertiesResponseOutput() MpnsCredentialPropertiesResponseOutput {
	return i.ToMpnsCredentialPropertiesResponseOutputWithContext(context.Background())
}

func (i MpnsCredentialPropertiesResponseArgs) ToMpnsCredentialPropertiesResponseOutputWithContext(ctx context.Context) MpnsCredentialPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpnsCredentialPropertiesResponseOutput)
}

func (i MpnsCredentialPropertiesResponseArgs) ToMpnsCredentialPropertiesResponsePtrOutput() MpnsCredentialPropertiesResponsePtrOutput {
	return i.ToMpnsCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MpnsCredentialPropertiesResponseArgs) ToMpnsCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) MpnsCredentialPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpnsCredentialPropertiesResponseOutput).ToMpnsCredentialPropertiesResponsePtrOutputWithContext(ctx)
}









type MpnsCredentialPropertiesResponsePtrInput interface {
	pulumi.Input

	ToMpnsCredentialPropertiesResponsePtrOutput() MpnsCredentialPropertiesResponsePtrOutput
	ToMpnsCredentialPropertiesResponsePtrOutputWithContext(context.Context) MpnsCredentialPropertiesResponsePtrOutput
}

type mpnsCredentialPropertiesResponsePtrType MpnsCredentialPropertiesResponseArgs

func MpnsCredentialPropertiesResponsePtr(v *MpnsCredentialPropertiesResponseArgs) MpnsCredentialPropertiesResponsePtrInput {
	return (*mpnsCredentialPropertiesResponsePtrType)(v)
}

func (*mpnsCredentialPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MpnsCredentialPropertiesResponse)(nil)).Elem()
}

func (i *mpnsCredentialPropertiesResponsePtrType) ToMpnsCredentialPropertiesResponsePtrOutput() MpnsCredentialPropertiesResponsePtrOutput {
	return i.ToMpnsCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *mpnsCredentialPropertiesResponsePtrType) ToMpnsCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) MpnsCredentialPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MpnsCredentialPropertiesResponsePtrOutput)
}

type MpnsCredentialPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MpnsCredentialPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MpnsCredentialPropertiesResponse)(nil)).Elem()
}

func (o MpnsCredentialPropertiesResponseOutput) ToMpnsCredentialPropertiesResponseOutput() MpnsCredentialPropertiesResponseOutput {
	return o
}

func (o MpnsCredentialPropertiesResponseOutput) ToMpnsCredentialPropertiesResponseOutputWithContext(ctx context.Context) MpnsCredentialPropertiesResponseOutput {
	return o
}

func (o MpnsCredentialPropertiesResponseOutput) ToMpnsCredentialPropertiesResponsePtrOutput() MpnsCredentialPropertiesResponsePtrOutput {
	return o.ToMpnsCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MpnsCredentialPropertiesResponseOutput) ToMpnsCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) MpnsCredentialPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MpnsCredentialPropertiesResponse) *MpnsCredentialPropertiesResponse {
		return &v
	}).(MpnsCredentialPropertiesResponsePtrOutput)
}

func (o MpnsCredentialPropertiesResponseOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MpnsCredentialPropertiesResponse) *string { return v.CertificateKey }).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialPropertiesResponseOutput) MpnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MpnsCredentialPropertiesResponse) *string { return v.MpnsCertificate }).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialPropertiesResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MpnsCredentialPropertiesResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type MpnsCredentialPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MpnsCredentialPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MpnsCredentialPropertiesResponse)(nil)).Elem()
}

func (o MpnsCredentialPropertiesResponsePtrOutput) ToMpnsCredentialPropertiesResponsePtrOutput() MpnsCredentialPropertiesResponsePtrOutput {
	return o
}

func (o MpnsCredentialPropertiesResponsePtrOutput) ToMpnsCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) MpnsCredentialPropertiesResponsePtrOutput {
	return o
}

func (o MpnsCredentialPropertiesResponsePtrOutput) Elem() MpnsCredentialPropertiesResponseOutput {
	return o.ApplyT(func(v *MpnsCredentialPropertiesResponse) MpnsCredentialPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MpnsCredentialPropertiesResponse
		return ret
	}).(MpnsCredentialPropertiesResponseOutput)
}

func (o MpnsCredentialPropertiesResponsePtrOutput) CertificateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MpnsCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CertificateKey
	}).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialPropertiesResponsePtrOutput) MpnsCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MpnsCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.MpnsCertificate
	}).(pulumi.StringPtrOutput)
}

func (o MpnsCredentialPropertiesResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MpnsCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type MpnsCredentialResponse struct {
	Properties *MpnsCredentialPropertiesResponse `pulumi:"properties"`
}





type MpnsCredentialResponseInput interface {
	pulumi.Input

	ToMpnsCredentialResponseOutput() MpnsCredentialResponseOutput
	ToMpnsCredentialResponseOutputWithContext(context.Context) MpnsCredentialResponseOutput
}

type MpnsCredentialResponseArgs struct {
	Properties MpnsCredentialPropertiesResponsePtrInput `pulumi:"properties"`
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

func (o MpnsCredentialResponseOutput) Properties() MpnsCredentialPropertiesResponsePtrOutput {
	return o.ApplyT(func(v MpnsCredentialResponse) *MpnsCredentialPropertiesResponse { return v.Properties }).(MpnsCredentialPropertiesResponsePtrOutput)
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

func (o MpnsCredentialResponsePtrOutput) Properties() MpnsCredentialPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *MpnsCredentialResponse) *MpnsCredentialPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(MpnsCredentialPropertiesResponsePtrOutput)
}

type NamespaceProperties struct {
	CreatedAt          *string        `pulumi:"createdAt"`
	Critical           *bool          `pulumi:"critical"`
	Enabled            *bool          `pulumi:"enabled"`
	Name               *string        `pulumi:"name"`
	NamespaceType      *NamespaceType `pulumi:"namespaceType"`
	ProvisioningState  *string        `pulumi:"provisioningState"`
	Region             *string        `pulumi:"region"`
	ScaleUnit          *string        `pulumi:"scaleUnit"`
	ServiceBusEndpoint *string        `pulumi:"serviceBusEndpoint"`
	Status             *string        `pulumi:"status"`
	SubscriptionId     *string        `pulumi:"subscriptionId"`
}





type NamespacePropertiesInput interface {
	pulumi.Input

	ToNamespacePropertiesOutput() NamespacePropertiesOutput
	ToNamespacePropertiesOutputWithContext(context.Context) NamespacePropertiesOutput
}

type NamespacePropertiesArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	Critical           pulumi.BoolPtrInput   `pulumi:"critical"`
	Enabled            pulumi.BoolPtrInput   `pulumi:"enabled"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	NamespaceType      NamespaceTypePtrInput `pulumi:"namespaceType"`
	ProvisioningState  pulumi.StringPtrInput `pulumi:"provisioningState"`
	Region             pulumi.StringPtrInput `pulumi:"region"`
	ScaleUnit          pulumi.StringPtrInput `pulumi:"scaleUnit"`
	ServiceBusEndpoint pulumi.StringPtrInput `pulumi:"serviceBusEndpoint"`
	Status             pulumi.StringPtrInput `pulumi:"status"`
	SubscriptionId     pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (NamespacePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceProperties)(nil)).Elem()
}

func (i NamespacePropertiesArgs) ToNamespacePropertiesOutput() NamespacePropertiesOutput {
	return i.ToNamespacePropertiesOutputWithContext(context.Background())
}

func (i NamespacePropertiesArgs) ToNamespacePropertiesOutputWithContext(ctx context.Context) NamespacePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePropertiesOutput)
}

func (i NamespacePropertiesArgs) ToNamespacePropertiesPtrOutput() NamespacePropertiesPtrOutput {
	return i.ToNamespacePropertiesPtrOutputWithContext(context.Background())
}

func (i NamespacePropertiesArgs) ToNamespacePropertiesPtrOutputWithContext(ctx context.Context) NamespacePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePropertiesOutput).ToNamespacePropertiesPtrOutputWithContext(ctx)
}









type NamespacePropertiesPtrInput interface {
	pulumi.Input

	ToNamespacePropertiesPtrOutput() NamespacePropertiesPtrOutput
	ToNamespacePropertiesPtrOutputWithContext(context.Context) NamespacePropertiesPtrOutput
}

type namespacePropertiesPtrType NamespacePropertiesArgs

func NamespacePropertiesPtr(v *NamespacePropertiesArgs) NamespacePropertiesPtrInput {
	return (*namespacePropertiesPtrType)(v)
}

func (*namespacePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceProperties)(nil)).Elem()
}

func (i *namespacePropertiesPtrType) ToNamespacePropertiesPtrOutput() NamespacePropertiesPtrOutput {
	return i.ToNamespacePropertiesPtrOutputWithContext(context.Background())
}

func (i *namespacePropertiesPtrType) ToNamespacePropertiesPtrOutputWithContext(ctx context.Context) NamespacePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePropertiesPtrOutput)
}

type NamespacePropertiesOutput struct{ *pulumi.OutputState }

func (NamespacePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceProperties)(nil)).Elem()
}

func (o NamespacePropertiesOutput) ToNamespacePropertiesOutput() NamespacePropertiesOutput {
	return o
}

func (o NamespacePropertiesOutput) ToNamespacePropertiesOutputWithContext(ctx context.Context) NamespacePropertiesOutput {
	return o
}

func (o NamespacePropertiesOutput) ToNamespacePropertiesPtrOutput() NamespacePropertiesPtrOutput {
	return o.ToNamespacePropertiesPtrOutputWithContext(context.Background())
}

func (o NamespacePropertiesOutput) ToNamespacePropertiesPtrOutputWithContext(ctx context.Context) NamespacePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NamespaceProperties) *NamespaceProperties {
		return &v
	}).(NamespacePropertiesPtrOutput)
}

func (o NamespacePropertiesOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceProperties) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesOutput) Critical() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NamespaceProperties) *bool { return v.Critical }).(pulumi.BoolPtrOutput)
}

func (o NamespacePropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NamespaceProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o NamespacePropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesOutput) NamespaceType() NamespaceTypePtrOutput {
	return o.ApplyT(func(v NamespaceProperties) *NamespaceType { return v.NamespaceType }).(NamespaceTypePtrOutput)
}

func (o NamespacePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceProperties) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesOutput) ScaleUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceProperties) *string { return v.ScaleUnit }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesOutput) ServiceBusEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceProperties) *string { return v.ServiceBusEndpoint }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceProperties) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type NamespacePropertiesPtrOutput struct{ *pulumi.OutputState }

func (NamespacePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceProperties)(nil)).Elem()
}

func (o NamespacePropertiesPtrOutput) ToNamespacePropertiesPtrOutput() NamespacePropertiesPtrOutput {
	return o
}

func (o NamespacePropertiesPtrOutput) ToNamespacePropertiesPtrOutputWithContext(ctx context.Context) NamespacePropertiesPtrOutput {
	return o
}

func (o NamespacePropertiesPtrOutput) Elem() NamespacePropertiesOutput {
	return o.ApplyT(func(v *NamespaceProperties) NamespaceProperties {
		if v != nil {
			return *v
		}
		var ret NamespaceProperties
		return ret
	}).(NamespacePropertiesOutput)
}

func (o NamespacePropertiesPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesPtrOutput) Critical() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NamespaceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Critical
	}).(pulumi.BoolPtrOutput)
}

func (o NamespacePropertiesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NamespaceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o NamespacePropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesPtrOutput) NamespaceType() NamespaceTypePtrOutput {
	return o.ApplyT(func(v *NamespaceProperties) *NamespaceType {
		if v == nil {
			return nil
		}
		return v.NamespaceType
	}).(NamespaceTypePtrOutput)
}

func (o NamespacePropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesPtrOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Region
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesPtrOutput) ScaleUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ScaleUnit
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesPtrOutput) ServiceBusEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ServiceBusEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesPtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

type NamespacePropertiesResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	Critical           *bool   `pulumi:"critical"`
	Enabled            *bool   `pulumi:"enabled"`
	Name               *string `pulumi:"name"`
	NamespaceType      *string `pulumi:"namespaceType"`
	ProvisioningState  *string `pulumi:"provisioningState"`
	Region             *string `pulumi:"region"`
	ScaleUnit          *string `pulumi:"scaleUnit"`
	ServiceBusEndpoint *string `pulumi:"serviceBusEndpoint"`
	Status             *string `pulumi:"status"`
	SubscriptionId     *string `pulumi:"subscriptionId"`
}





type NamespacePropertiesResponseInput interface {
	pulumi.Input

	ToNamespacePropertiesResponseOutput() NamespacePropertiesResponseOutput
	ToNamespacePropertiesResponseOutputWithContext(context.Context) NamespacePropertiesResponseOutput
}

type NamespacePropertiesResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	Critical           pulumi.BoolPtrInput   `pulumi:"critical"`
	Enabled            pulumi.BoolPtrInput   `pulumi:"enabled"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	NamespaceType      pulumi.StringPtrInput `pulumi:"namespaceType"`
	ProvisioningState  pulumi.StringPtrInput `pulumi:"provisioningState"`
	Region             pulumi.StringPtrInput `pulumi:"region"`
	ScaleUnit          pulumi.StringPtrInput `pulumi:"scaleUnit"`
	ServiceBusEndpoint pulumi.StringPtrInput `pulumi:"serviceBusEndpoint"`
	Status             pulumi.StringPtrInput `pulumi:"status"`
	SubscriptionId     pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (NamespacePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespacePropertiesResponse)(nil)).Elem()
}

func (i NamespacePropertiesResponseArgs) ToNamespacePropertiesResponseOutput() NamespacePropertiesResponseOutput {
	return i.ToNamespacePropertiesResponseOutputWithContext(context.Background())
}

func (i NamespacePropertiesResponseArgs) ToNamespacePropertiesResponseOutputWithContext(ctx context.Context) NamespacePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePropertiesResponseOutput)
}

func (i NamespacePropertiesResponseArgs) ToNamespacePropertiesResponsePtrOutput() NamespacePropertiesResponsePtrOutput {
	return i.ToNamespacePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i NamespacePropertiesResponseArgs) ToNamespacePropertiesResponsePtrOutputWithContext(ctx context.Context) NamespacePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePropertiesResponseOutput).ToNamespacePropertiesResponsePtrOutputWithContext(ctx)
}









type NamespacePropertiesResponsePtrInput interface {
	pulumi.Input

	ToNamespacePropertiesResponsePtrOutput() NamespacePropertiesResponsePtrOutput
	ToNamespacePropertiesResponsePtrOutputWithContext(context.Context) NamespacePropertiesResponsePtrOutput
}

type namespacePropertiesResponsePtrType NamespacePropertiesResponseArgs

func NamespacePropertiesResponsePtr(v *NamespacePropertiesResponseArgs) NamespacePropertiesResponsePtrInput {
	return (*namespacePropertiesResponsePtrType)(v)
}

func (*namespacePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespacePropertiesResponse)(nil)).Elem()
}

func (i *namespacePropertiesResponsePtrType) ToNamespacePropertiesResponsePtrOutput() NamespacePropertiesResponsePtrOutput {
	return i.ToNamespacePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *namespacePropertiesResponsePtrType) ToNamespacePropertiesResponsePtrOutputWithContext(ctx context.Context) NamespacePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePropertiesResponsePtrOutput)
}

type NamespacePropertiesResponseOutput struct{ *pulumi.OutputState }

func (NamespacePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespacePropertiesResponse)(nil)).Elem()
}

func (o NamespacePropertiesResponseOutput) ToNamespacePropertiesResponseOutput() NamespacePropertiesResponseOutput {
	return o
}

func (o NamespacePropertiesResponseOutput) ToNamespacePropertiesResponseOutputWithContext(ctx context.Context) NamespacePropertiesResponseOutput {
	return o
}

func (o NamespacePropertiesResponseOutput) ToNamespacePropertiesResponsePtrOutput() NamespacePropertiesResponsePtrOutput {
	return o.ToNamespacePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o NamespacePropertiesResponseOutput) ToNamespacePropertiesResponsePtrOutputWithContext(ctx context.Context) NamespacePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NamespacePropertiesResponse) *NamespacePropertiesResponse {
		return &v
	}).(NamespacePropertiesResponsePtrOutput)
}

func (o NamespacePropertiesResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespacePropertiesResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponseOutput) Critical() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NamespacePropertiesResponse) *bool { return v.Critical }).(pulumi.BoolPtrOutput)
}

func (o NamespacePropertiesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NamespacePropertiesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o NamespacePropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespacePropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponseOutput) NamespaceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespacePropertiesResponse) *string { return v.NamespaceType }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespacePropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponseOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespacePropertiesResponse) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponseOutput) ScaleUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespacePropertiesResponse) *string { return v.ScaleUnit }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponseOutput) ServiceBusEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespacePropertiesResponse) *string { return v.ServiceBusEndpoint }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespacePropertiesResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespacePropertiesResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type NamespacePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (NamespacePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespacePropertiesResponse)(nil)).Elem()
}

func (o NamespacePropertiesResponsePtrOutput) ToNamespacePropertiesResponsePtrOutput() NamespacePropertiesResponsePtrOutput {
	return o
}

func (o NamespacePropertiesResponsePtrOutput) ToNamespacePropertiesResponsePtrOutputWithContext(ctx context.Context) NamespacePropertiesResponsePtrOutput {
	return o
}

func (o NamespacePropertiesResponsePtrOutput) Elem() NamespacePropertiesResponseOutput {
	return o.ApplyT(func(v *NamespacePropertiesResponse) NamespacePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret NamespacePropertiesResponse
		return ret
	}).(NamespacePropertiesResponseOutput)
}

func (o NamespacePropertiesResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponsePtrOutput) Critical() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NamespacePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Critical
	}).(pulumi.BoolPtrOutput)
}

func (o NamespacePropertiesResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NamespacePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o NamespacePropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponsePtrOutput) NamespaceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.NamespaceType
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponsePtrOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Region
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponsePtrOutput) ScaleUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScaleUnit
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponsePtrOutput) ServiceBusEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceBusEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o NamespacePropertiesResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

type NotificationHubProperties struct {
	AdmCredential      *AdmCredential                            `pulumi:"admCredential"`
	ApnsCredential     *ApnsCredential                           `pulumi:"apnsCredential"`
	AuthorizationRules []SharedAccessAuthorizationRuleProperties `pulumi:"authorizationRules"`
	BaiduCredential    *BaiduCredential                          `pulumi:"baiduCredential"`
	GcmCredential      *GcmCredential                            `pulumi:"gcmCredential"`
	MpnsCredential     *MpnsCredential                           `pulumi:"mpnsCredential"`
	Name               *string                                   `pulumi:"name"`
	RegistrationTtl    *string                                   `pulumi:"registrationTtl"`
	WnsCredential      *WnsCredential                            `pulumi:"wnsCredential"`
}





type NotificationHubPropertiesInput interface {
	pulumi.Input

	ToNotificationHubPropertiesOutput() NotificationHubPropertiesOutput
	ToNotificationHubPropertiesOutputWithContext(context.Context) NotificationHubPropertiesOutput
}

type NotificationHubPropertiesArgs struct {
	AdmCredential      AdmCredentialPtrInput                             `pulumi:"admCredential"`
	ApnsCredential     ApnsCredentialPtrInput                            `pulumi:"apnsCredential"`
	AuthorizationRules SharedAccessAuthorizationRulePropertiesArrayInput `pulumi:"authorizationRules"`
	BaiduCredential    BaiduCredentialPtrInput                           `pulumi:"baiduCredential"`
	GcmCredential      GcmCredentialPtrInput                             `pulumi:"gcmCredential"`
	MpnsCredential     MpnsCredentialPtrInput                            `pulumi:"mpnsCredential"`
	Name               pulumi.StringPtrInput                             `pulumi:"name"`
	RegistrationTtl    pulumi.StringPtrInput                             `pulumi:"registrationTtl"`
	WnsCredential      WnsCredentialPtrInput                             `pulumi:"wnsCredential"`
}

func (NotificationHubPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationHubProperties)(nil)).Elem()
}

func (i NotificationHubPropertiesArgs) ToNotificationHubPropertiesOutput() NotificationHubPropertiesOutput {
	return i.ToNotificationHubPropertiesOutputWithContext(context.Background())
}

func (i NotificationHubPropertiesArgs) ToNotificationHubPropertiesOutputWithContext(ctx context.Context) NotificationHubPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationHubPropertiesOutput)
}

func (i NotificationHubPropertiesArgs) ToNotificationHubPropertiesPtrOutput() NotificationHubPropertiesPtrOutput {
	return i.ToNotificationHubPropertiesPtrOutputWithContext(context.Background())
}

func (i NotificationHubPropertiesArgs) ToNotificationHubPropertiesPtrOutputWithContext(ctx context.Context) NotificationHubPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationHubPropertiesOutput).ToNotificationHubPropertiesPtrOutputWithContext(ctx)
}









type NotificationHubPropertiesPtrInput interface {
	pulumi.Input

	ToNotificationHubPropertiesPtrOutput() NotificationHubPropertiesPtrOutput
	ToNotificationHubPropertiesPtrOutputWithContext(context.Context) NotificationHubPropertiesPtrOutput
}

type notificationHubPropertiesPtrType NotificationHubPropertiesArgs

func NotificationHubPropertiesPtr(v *NotificationHubPropertiesArgs) NotificationHubPropertiesPtrInput {
	return (*notificationHubPropertiesPtrType)(v)
}

func (*notificationHubPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationHubProperties)(nil)).Elem()
}

func (i *notificationHubPropertiesPtrType) ToNotificationHubPropertiesPtrOutput() NotificationHubPropertiesPtrOutput {
	return i.ToNotificationHubPropertiesPtrOutputWithContext(context.Background())
}

func (i *notificationHubPropertiesPtrType) ToNotificationHubPropertiesPtrOutputWithContext(ctx context.Context) NotificationHubPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationHubPropertiesPtrOutput)
}

type NotificationHubPropertiesOutput struct{ *pulumi.OutputState }

func (NotificationHubPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationHubProperties)(nil)).Elem()
}

func (o NotificationHubPropertiesOutput) ToNotificationHubPropertiesOutput() NotificationHubPropertiesOutput {
	return o
}

func (o NotificationHubPropertiesOutput) ToNotificationHubPropertiesOutputWithContext(ctx context.Context) NotificationHubPropertiesOutput {
	return o
}

func (o NotificationHubPropertiesOutput) ToNotificationHubPropertiesPtrOutput() NotificationHubPropertiesPtrOutput {
	return o.ToNotificationHubPropertiesPtrOutputWithContext(context.Background())
}

func (o NotificationHubPropertiesOutput) ToNotificationHubPropertiesPtrOutputWithContext(ctx context.Context) NotificationHubPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationHubProperties) *NotificationHubProperties {
		return &v
	}).(NotificationHubPropertiesPtrOutput)
}

func (o NotificationHubPropertiesOutput) AdmCredential() AdmCredentialPtrOutput {
	return o.ApplyT(func(v NotificationHubProperties) *AdmCredential { return v.AdmCredential }).(AdmCredentialPtrOutput)
}

func (o NotificationHubPropertiesOutput) ApnsCredential() ApnsCredentialPtrOutput {
	return o.ApplyT(func(v NotificationHubProperties) *ApnsCredential { return v.ApnsCredential }).(ApnsCredentialPtrOutput)
}

func (o NotificationHubPropertiesOutput) AuthorizationRules() SharedAccessAuthorizationRulePropertiesArrayOutput {
	return o.ApplyT(func(v NotificationHubProperties) []SharedAccessAuthorizationRuleProperties {
		return v.AuthorizationRules
	}).(SharedAccessAuthorizationRulePropertiesArrayOutput)
}

func (o NotificationHubPropertiesOutput) BaiduCredential() BaiduCredentialPtrOutput {
	return o.ApplyT(func(v NotificationHubProperties) *BaiduCredential { return v.BaiduCredential }).(BaiduCredentialPtrOutput)
}

func (o NotificationHubPropertiesOutput) GcmCredential() GcmCredentialPtrOutput {
	return o.ApplyT(func(v NotificationHubProperties) *GcmCredential { return v.GcmCredential }).(GcmCredentialPtrOutput)
}

func (o NotificationHubPropertiesOutput) MpnsCredential() MpnsCredentialPtrOutput {
	return o.ApplyT(func(v NotificationHubProperties) *MpnsCredential { return v.MpnsCredential }).(MpnsCredentialPtrOutput)
}

func (o NotificationHubPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationHubProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NotificationHubPropertiesOutput) RegistrationTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationHubProperties) *string { return v.RegistrationTtl }).(pulumi.StringPtrOutput)
}

func (o NotificationHubPropertiesOutput) WnsCredential() WnsCredentialPtrOutput {
	return o.ApplyT(func(v NotificationHubProperties) *WnsCredential { return v.WnsCredential }).(WnsCredentialPtrOutput)
}

type NotificationHubPropertiesPtrOutput struct{ *pulumi.OutputState }

func (NotificationHubPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationHubProperties)(nil)).Elem()
}

func (o NotificationHubPropertiesPtrOutput) ToNotificationHubPropertiesPtrOutput() NotificationHubPropertiesPtrOutput {
	return o
}

func (o NotificationHubPropertiesPtrOutput) ToNotificationHubPropertiesPtrOutputWithContext(ctx context.Context) NotificationHubPropertiesPtrOutput {
	return o
}

func (o NotificationHubPropertiesPtrOutput) Elem() NotificationHubPropertiesOutput {
	return o.ApplyT(func(v *NotificationHubProperties) NotificationHubProperties {
		if v != nil {
			return *v
		}
		var ret NotificationHubProperties
		return ret
	}).(NotificationHubPropertiesOutput)
}

func (o NotificationHubPropertiesPtrOutput) AdmCredential() AdmCredentialPtrOutput {
	return o.ApplyT(func(v *NotificationHubProperties) *AdmCredential {
		if v == nil {
			return nil
		}
		return v.AdmCredential
	}).(AdmCredentialPtrOutput)
}

func (o NotificationHubPropertiesPtrOutput) ApnsCredential() ApnsCredentialPtrOutput {
	return o.ApplyT(func(v *NotificationHubProperties) *ApnsCredential {
		if v == nil {
			return nil
		}
		return v.ApnsCredential
	}).(ApnsCredentialPtrOutput)
}

func (o NotificationHubPropertiesPtrOutput) AuthorizationRules() SharedAccessAuthorizationRulePropertiesArrayOutput {
	return o.ApplyT(func(v *NotificationHubProperties) []SharedAccessAuthorizationRuleProperties {
		if v == nil {
			return nil
		}
		return v.AuthorizationRules
	}).(SharedAccessAuthorizationRulePropertiesArrayOutput)
}

func (o NotificationHubPropertiesPtrOutput) BaiduCredential() BaiduCredentialPtrOutput {
	return o.ApplyT(func(v *NotificationHubProperties) *BaiduCredential {
		if v == nil {
			return nil
		}
		return v.BaiduCredential
	}).(BaiduCredentialPtrOutput)
}

func (o NotificationHubPropertiesPtrOutput) GcmCredential() GcmCredentialPtrOutput {
	return o.ApplyT(func(v *NotificationHubProperties) *GcmCredential {
		if v == nil {
			return nil
		}
		return v.GcmCredential
	}).(GcmCredentialPtrOutput)
}

func (o NotificationHubPropertiesPtrOutput) MpnsCredential() MpnsCredentialPtrOutput {
	return o.ApplyT(func(v *NotificationHubProperties) *MpnsCredential {
		if v == nil {
			return nil
		}
		return v.MpnsCredential
	}).(MpnsCredentialPtrOutput)
}

func (o NotificationHubPropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationHubProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o NotificationHubPropertiesPtrOutput) RegistrationTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationHubProperties) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationTtl
	}).(pulumi.StringPtrOutput)
}

func (o NotificationHubPropertiesPtrOutput) WnsCredential() WnsCredentialPtrOutput {
	return o.ApplyT(func(v *NotificationHubProperties) *WnsCredential {
		if v == nil {
			return nil
		}
		return v.WnsCredential
	}).(WnsCredentialPtrOutput)
}

type NotificationHubPropertiesResponse struct {
	AdmCredential      *AdmCredentialResponse                            `pulumi:"admCredential"`
	ApnsCredential     *ApnsCredentialResponse                           `pulumi:"apnsCredential"`
	AuthorizationRules []SharedAccessAuthorizationRulePropertiesResponse `pulumi:"authorizationRules"`
	BaiduCredential    *BaiduCredentialResponse                          `pulumi:"baiduCredential"`
	GcmCredential      *GcmCredentialResponse                            `pulumi:"gcmCredential"`
	MpnsCredential     *MpnsCredentialResponse                           `pulumi:"mpnsCredential"`
	Name               *string                                           `pulumi:"name"`
	RegistrationTtl    *string                                           `pulumi:"registrationTtl"`
	WnsCredential      *WnsCredentialResponse                            `pulumi:"wnsCredential"`
}





type NotificationHubPropertiesResponseInput interface {
	pulumi.Input

	ToNotificationHubPropertiesResponseOutput() NotificationHubPropertiesResponseOutput
	ToNotificationHubPropertiesResponseOutputWithContext(context.Context) NotificationHubPropertiesResponseOutput
}

type NotificationHubPropertiesResponseArgs struct {
	AdmCredential      AdmCredentialResponsePtrInput                             `pulumi:"admCredential"`
	ApnsCredential     ApnsCredentialResponsePtrInput                            `pulumi:"apnsCredential"`
	AuthorizationRules SharedAccessAuthorizationRulePropertiesResponseArrayInput `pulumi:"authorizationRules"`
	BaiduCredential    BaiduCredentialResponsePtrInput                           `pulumi:"baiduCredential"`
	GcmCredential      GcmCredentialResponsePtrInput                             `pulumi:"gcmCredential"`
	MpnsCredential     MpnsCredentialResponsePtrInput                            `pulumi:"mpnsCredential"`
	Name               pulumi.StringPtrInput                                     `pulumi:"name"`
	RegistrationTtl    pulumi.StringPtrInput                                     `pulumi:"registrationTtl"`
	WnsCredential      WnsCredentialResponsePtrInput                             `pulumi:"wnsCredential"`
}

func (NotificationHubPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationHubPropertiesResponse)(nil)).Elem()
}

func (i NotificationHubPropertiesResponseArgs) ToNotificationHubPropertiesResponseOutput() NotificationHubPropertiesResponseOutput {
	return i.ToNotificationHubPropertiesResponseOutputWithContext(context.Background())
}

func (i NotificationHubPropertiesResponseArgs) ToNotificationHubPropertiesResponseOutputWithContext(ctx context.Context) NotificationHubPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationHubPropertiesResponseOutput)
}

func (i NotificationHubPropertiesResponseArgs) ToNotificationHubPropertiesResponsePtrOutput() NotificationHubPropertiesResponsePtrOutput {
	return i.ToNotificationHubPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i NotificationHubPropertiesResponseArgs) ToNotificationHubPropertiesResponsePtrOutputWithContext(ctx context.Context) NotificationHubPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationHubPropertiesResponseOutput).ToNotificationHubPropertiesResponsePtrOutputWithContext(ctx)
}









type NotificationHubPropertiesResponsePtrInput interface {
	pulumi.Input

	ToNotificationHubPropertiesResponsePtrOutput() NotificationHubPropertiesResponsePtrOutput
	ToNotificationHubPropertiesResponsePtrOutputWithContext(context.Context) NotificationHubPropertiesResponsePtrOutput
}

type notificationHubPropertiesResponsePtrType NotificationHubPropertiesResponseArgs

func NotificationHubPropertiesResponsePtr(v *NotificationHubPropertiesResponseArgs) NotificationHubPropertiesResponsePtrInput {
	return (*notificationHubPropertiesResponsePtrType)(v)
}

func (*notificationHubPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationHubPropertiesResponse)(nil)).Elem()
}

func (i *notificationHubPropertiesResponsePtrType) ToNotificationHubPropertiesResponsePtrOutput() NotificationHubPropertiesResponsePtrOutput {
	return i.ToNotificationHubPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *notificationHubPropertiesResponsePtrType) ToNotificationHubPropertiesResponsePtrOutputWithContext(ctx context.Context) NotificationHubPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationHubPropertiesResponsePtrOutput)
}

type NotificationHubPropertiesResponseOutput struct{ *pulumi.OutputState }

func (NotificationHubPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationHubPropertiesResponse)(nil)).Elem()
}

func (o NotificationHubPropertiesResponseOutput) ToNotificationHubPropertiesResponseOutput() NotificationHubPropertiesResponseOutput {
	return o
}

func (o NotificationHubPropertiesResponseOutput) ToNotificationHubPropertiesResponseOutputWithContext(ctx context.Context) NotificationHubPropertiesResponseOutput {
	return o
}

func (o NotificationHubPropertiesResponseOutput) ToNotificationHubPropertiesResponsePtrOutput() NotificationHubPropertiesResponsePtrOutput {
	return o.ToNotificationHubPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o NotificationHubPropertiesResponseOutput) ToNotificationHubPropertiesResponsePtrOutputWithContext(ctx context.Context) NotificationHubPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationHubPropertiesResponse) *NotificationHubPropertiesResponse {
		return &v
	}).(NotificationHubPropertiesResponsePtrOutput)
}

func (o NotificationHubPropertiesResponseOutput) AdmCredential() AdmCredentialResponsePtrOutput {
	return o.ApplyT(func(v NotificationHubPropertiesResponse) *AdmCredentialResponse { return v.AdmCredential }).(AdmCredentialResponsePtrOutput)
}

func (o NotificationHubPropertiesResponseOutput) ApnsCredential() ApnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v NotificationHubPropertiesResponse) *ApnsCredentialResponse { return v.ApnsCredential }).(ApnsCredentialResponsePtrOutput)
}

func (o NotificationHubPropertiesResponseOutput) AuthorizationRules() SharedAccessAuthorizationRulePropertiesResponseArrayOutput {
	return o.ApplyT(func(v NotificationHubPropertiesResponse) []SharedAccessAuthorizationRulePropertiesResponse {
		return v.AuthorizationRules
	}).(SharedAccessAuthorizationRulePropertiesResponseArrayOutput)
}

func (o NotificationHubPropertiesResponseOutput) BaiduCredential() BaiduCredentialResponsePtrOutput {
	return o.ApplyT(func(v NotificationHubPropertiesResponse) *BaiduCredentialResponse { return v.BaiduCredential }).(BaiduCredentialResponsePtrOutput)
}

func (o NotificationHubPropertiesResponseOutput) GcmCredential() GcmCredentialResponsePtrOutput {
	return o.ApplyT(func(v NotificationHubPropertiesResponse) *GcmCredentialResponse { return v.GcmCredential }).(GcmCredentialResponsePtrOutput)
}

func (o NotificationHubPropertiesResponseOutput) MpnsCredential() MpnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v NotificationHubPropertiesResponse) *MpnsCredentialResponse { return v.MpnsCredential }).(MpnsCredentialResponsePtrOutput)
}

func (o NotificationHubPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationHubPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NotificationHubPropertiesResponseOutput) RegistrationTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationHubPropertiesResponse) *string { return v.RegistrationTtl }).(pulumi.StringPtrOutput)
}

func (o NotificationHubPropertiesResponseOutput) WnsCredential() WnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v NotificationHubPropertiesResponse) *WnsCredentialResponse { return v.WnsCredential }).(WnsCredentialResponsePtrOutput)
}

type NotificationHubPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (NotificationHubPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationHubPropertiesResponse)(nil)).Elem()
}

func (o NotificationHubPropertiesResponsePtrOutput) ToNotificationHubPropertiesResponsePtrOutput() NotificationHubPropertiesResponsePtrOutput {
	return o
}

func (o NotificationHubPropertiesResponsePtrOutput) ToNotificationHubPropertiesResponsePtrOutputWithContext(ctx context.Context) NotificationHubPropertiesResponsePtrOutput {
	return o
}

func (o NotificationHubPropertiesResponsePtrOutput) Elem() NotificationHubPropertiesResponseOutput {
	return o.ApplyT(func(v *NotificationHubPropertiesResponse) NotificationHubPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret NotificationHubPropertiesResponse
		return ret
	}).(NotificationHubPropertiesResponseOutput)
}

func (o NotificationHubPropertiesResponsePtrOutput) AdmCredential() AdmCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHubPropertiesResponse) *AdmCredentialResponse {
		if v == nil {
			return nil
		}
		return v.AdmCredential
	}).(AdmCredentialResponsePtrOutput)
}

func (o NotificationHubPropertiesResponsePtrOutput) ApnsCredential() ApnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHubPropertiesResponse) *ApnsCredentialResponse {
		if v == nil {
			return nil
		}
		return v.ApnsCredential
	}).(ApnsCredentialResponsePtrOutput)
}

func (o NotificationHubPropertiesResponsePtrOutput) AuthorizationRules() SharedAccessAuthorizationRulePropertiesResponseArrayOutput {
	return o.ApplyT(func(v *NotificationHubPropertiesResponse) []SharedAccessAuthorizationRulePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.AuthorizationRules
	}).(SharedAccessAuthorizationRulePropertiesResponseArrayOutput)
}

func (o NotificationHubPropertiesResponsePtrOutput) BaiduCredential() BaiduCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHubPropertiesResponse) *BaiduCredentialResponse {
		if v == nil {
			return nil
		}
		return v.BaiduCredential
	}).(BaiduCredentialResponsePtrOutput)
}

func (o NotificationHubPropertiesResponsePtrOutput) GcmCredential() GcmCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHubPropertiesResponse) *GcmCredentialResponse {
		if v == nil {
			return nil
		}
		return v.GcmCredential
	}).(GcmCredentialResponsePtrOutput)
}

func (o NotificationHubPropertiesResponsePtrOutput) MpnsCredential() MpnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHubPropertiesResponse) *MpnsCredentialResponse {
		if v == nil {
			return nil
		}
		return v.MpnsCredential
	}).(MpnsCredentialResponsePtrOutput)
}

func (o NotificationHubPropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationHubPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o NotificationHubPropertiesResponsePtrOutput) RegistrationTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationHubPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationTtl
	}).(pulumi.StringPtrOutput)
}

func (o NotificationHubPropertiesResponsePtrOutput) WnsCredential() WnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v *NotificationHubPropertiesResponse) *WnsCredentialResponse {
		if v == nil {
			return nil
		}
		return v.WnsCredential
	}).(WnsCredentialResponsePtrOutput)
}

type SharedAccessAuthorizationRuleProperties struct {
	ClaimType    *string        `pulumi:"claimType"`
	ClaimValue   *string        `pulumi:"claimValue"`
	CreatedTime  *string        `pulumi:"createdTime"`
	KeyName      *string        `pulumi:"keyName"`
	ModifiedTime *string        `pulumi:"modifiedTime"`
	PrimaryKey   *string        `pulumi:"primaryKey"`
	Revision     *int           `pulumi:"revision"`
	Rights       []AccessRights `pulumi:"rights"`
	SecondaryKey *string        `pulumi:"secondaryKey"`
}





type SharedAccessAuthorizationRulePropertiesInput interface {
	pulumi.Input

	ToSharedAccessAuthorizationRulePropertiesOutput() SharedAccessAuthorizationRulePropertiesOutput
	ToSharedAccessAuthorizationRulePropertiesOutputWithContext(context.Context) SharedAccessAuthorizationRulePropertiesOutput
}

type SharedAccessAuthorizationRulePropertiesArgs struct {
	ClaimType    pulumi.StringPtrInput  `pulumi:"claimType"`
	ClaimValue   pulumi.StringPtrInput  `pulumi:"claimValue"`
	CreatedTime  pulumi.StringPtrInput  `pulumi:"createdTime"`
	KeyName      pulumi.StringPtrInput  `pulumi:"keyName"`
	ModifiedTime pulumi.StringPtrInput  `pulumi:"modifiedTime"`
	PrimaryKey   pulumi.StringPtrInput  `pulumi:"primaryKey"`
	Revision     pulumi.IntPtrInput     `pulumi:"revision"`
	Rights       AccessRightsArrayInput `pulumi:"rights"`
	SecondaryKey pulumi.StringPtrInput  `pulumi:"secondaryKey"`
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

func (o SharedAccessAuthorizationRulePropertiesOutput) ClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRuleProperties) *string { return v.ClaimType }).(pulumi.StringPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesOutput) ClaimValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRuleProperties) *string { return v.ClaimValue }).(pulumi.StringPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRuleProperties) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRuleProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesOutput) ModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRuleProperties) *string { return v.ModifiedTime }).(pulumi.StringPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRuleProperties) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesOutput) Revision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRuleProperties) *int { return v.Revision }).(pulumi.IntPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesOutput) Rights() AccessRightsArrayOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRuleProperties) []AccessRights { return v.Rights }).(AccessRightsArrayOutput)
}

func (o SharedAccessAuthorizationRulePropertiesOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRuleProperties) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
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
	ClaimType    *string  `pulumi:"claimType"`
	ClaimValue   *string  `pulumi:"claimValue"`
	CreatedTime  *string  `pulumi:"createdTime"`
	KeyName      *string  `pulumi:"keyName"`
	ModifiedTime *string  `pulumi:"modifiedTime"`
	PrimaryKey   *string  `pulumi:"primaryKey"`
	Revision     *int     `pulumi:"revision"`
	Rights       []string `pulumi:"rights"`
	SecondaryKey *string  `pulumi:"secondaryKey"`
}





type SharedAccessAuthorizationRulePropertiesResponseInput interface {
	pulumi.Input

	ToSharedAccessAuthorizationRulePropertiesResponseOutput() SharedAccessAuthorizationRulePropertiesResponseOutput
	ToSharedAccessAuthorizationRulePropertiesResponseOutputWithContext(context.Context) SharedAccessAuthorizationRulePropertiesResponseOutput
}

type SharedAccessAuthorizationRulePropertiesResponseArgs struct {
	ClaimType    pulumi.StringPtrInput   `pulumi:"claimType"`
	ClaimValue   pulumi.StringPtrInput   `pulumi:"claimValue"`
	CreatedTime  pulumi.StringPtrInput   `pulumi:"createdTime"`
	KeyName      pulumi.StringPtrInput   `pulumi:"keyName"`
	ModifiedTime pulumi.StringPtrInput   `pulumi:"modifiedTime"`
	PrimaryKey   pulumi.StringPtrInput   `pulumi:"primaryKey"`
	Revision     pulumi.IntPtrInput      `pulumi:"revision"`
	Rights       pulumi.StringArrayInput `pulumi:"rights"`
	SecondaryKey pulumi.StringPtrInput   `pulumi:"secondaryKey"`
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

func (o SharedAccessAuthorizationRulePropertiesResponseOutput) ClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRulePropertiesResponse) *string { return v.ClaimType }).(pulumi.StringPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesResponseOutput) ClaimValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRulePropertiesResponse) *string { return v.ClaimValue }).(pulumi.StringPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesResponseOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRulePropertiesResponse) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRulePropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesResponseOutput) ModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRulePropertiesResponse) *string { return v.ModifiedTime }).(pulumi.StringPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesResponseOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRulePropertiesResponse) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesResponseOutput) Revision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRulePropertiesResponse) *int { return v.Revision }).(pulumi.IntPtrOutput)
}

func (o SharedAccessAuthorizationRulePropertiesResponseOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRulePropertiesResponse) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o SharedAccessAuthorizationRulePropertiesResponseOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessAuthorizationRulePropertiesResponse) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
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

type WnsCredential struct {
	Properties *WnsCredentialProperties `pulumi:"properties"`
}





type WnsCredentialInput interface {
	pulumi.Input

	ToWnsCredentialOutput() WnsCredentialOutput
	ToWnsCredentialOutputWithContext(context.Context) WnsCredentialOutput
}

type WnsCredentialArgs struct {
	Properties WnsCredentialPropertiesPtrInput `pulumi:"properties"`
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

func (o WnsCredentialOutput) Properties() WnsCredentialPropertiesPtrOutput {
	return o.ApplyT(func(v WnsCredential) *WnsCredentialProperties { return v.Properties }).(WnsCredentialPropertiesPtrOutput)
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

func (o WnsCredentialPtrOutput) Properties() WnsCredentialPropertiesPtrOutput {
	return o.ApplyT(func(v *WnsCredential) *WnsCredentialProperties {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(WnsCredentialPropertiesPtrOutput)
}

type WnsCredentialProperties struct {
	PackageSid          *string `pulumi:"packageSid"`
	SecretKey           *string `pulumi:"secretKey"`
	WindowsLiveEndpoint *string `pulumi:"windowsLiveEndpoint"`
}





type WnsCredentialPropertiesInput interface {
	pulumi.Input

	ToWnsCredentialPropertiesOutput() WnsCredentialPropertiesOutput
	ToWnsCredentialPropertiesOutputWithContext(context.Context) WnsCredentialPropertiesOutput
}

type WnsCredentialPropertiesArgs struct {
	PackageSid          pulumi.StringPtrInput `pulumi:"packageSid"`
	SecretKey           pulumi.StringPtrInput `pulumi:"secretKey"`
	WindowsLiveEndpoint pulumi.StringPtrInput `pulumi:"windowsLiveEndpoint"`
}

func (WnsCredentialPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WnsCredentialProperties)(nil)).Elem()
}

func (i WnsCredentialPropertiesArgs) ToWnsCredentialPropertiesOutput() WnsCredentialPropertiesOutput {
	return i.ToWnsCredentialPropertiesOutputWithContext(context.Background())
}

func (i WnsCredentialPropertiesArgs) ToWnsCredentialPropertiesOutputWithContext(ctx context.Context) WnsCredentialPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WnsCredentialPropertiesOutput)
}

func (i WnsCredentialPropertiesArgs) ToWnsCredentialPropertiesPtrOutput() WnsCredentialPropertiesPtrOutput {
	return i.ToWnsCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (i WnsCredentialPropertiesArgs) ToWnsCredentialPropertiesPtrOutputWithContext(ctx context.Context) WnsCredentialPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WnsCredentialPropertiesOutput).ToWnsCredentialPropertiesPtrOutputWithContext(ctx)
}









type WnsCredentialPropertiesPtrInput interface {
	pulumi.Input

	ToWnsCredentialPropertiesPtrOutput() WnsCredentialPropertiesPtrOutput
	ToWnsCredentialPropertiesPtrOutputWithContext(context.Context) WnsCredentialPropertiesPtrOutput
}

type wnsCredentialPropertiesPtrType WnsCredentialPropertiesArgs

func WnsCredentialPropertiesPtr(v *WnsCredentialPropertiesArgs) WnsCredentialPropertiesPtrInput {
	return (*wnsCredentialPropertiesPtrType)(v)
}

func (*wnsCredentialPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WnsCredentialProperties)(nil)).Elem()
}

func (i *wnsCredentialPropertiesPtrType) ToWnsCredentialPropertiesPtrOutput() WnsCredentialPropertiesPtrOutput {
	return i.ToWnsCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (i *wnsCredentialPropertiesPtrType) ToWnsCredentialPropertiesPtrOutputWithContext(ctx context.Context) WnsCredentialPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WnsCredentialPropertiesPtrOutput)
}

type WnsCredentialPropertiesOutput struct{ *pulumi.OutputState }

func (WnsCredentialPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WnsCredentialProperties)(nil)).Elem()
}

func (o WnsCredentialPropertiesOutput) ToWnsCredentialPropertiesOutput() WnsCredentialPropertiesOutput {
	return o
}

func (o WnsCredentialPropertiesOutput) ToWnsCredentialPropertiesOutputWithContext(ctx context.Context) WnsCredentialPropertiesOutput {
	return o
}

func (o WnsCredentialPropertiesOutput) ToWnsCredentialPropertiesPtrOutput() WnsCredentialPropertiesPtrOutput {
	return o.ToWnsCredentialPropertiesPtrOutputWithContext(context.Background())
}

func (o WnsCredentialPropertiesOutput) ToWnsCredentialPropertiesPtrOutputWithContext(ctx context.Context) WnsCredentialPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WnsCredentialProperties) *WnsCredentialProperties {
		return &v
	}).(WnsCredentialPropertiesPtrOutput)
}

func (o WnsCredentialPropertiesOutput) PackageSid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WnsCredentialProperties) *string { return v.PackageSid }).(pulumi.StringPtrOutput)
}

func (o WnsCredentialPropertiesOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WnsCredentialProperties) *string { return v.SecretKey }).(pulumi.StringPtrOutput)
}

func (o WnsCredentialPropertiesOutput) WindowsLiveEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WnsCredentialProperties) *string { return v.WindowsLiveEndpoint }).(pulumi.StringPtrOutput)
}

type WnsCredentialPropertiesPtrOutput struct{ *pulumi.OutputState }

func (WnsCredentialPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WnsCredentialProperties)(nil)).Elem()
}

func (o WnsCredentialPropertiesPtrOutput) ToWnsCredentialPropertiesPtrOutput() WnsCredentialPropertiesPtrOutput {
	return o
}

func (o WnsCredentialPropertiesPtrOutput) ToWnsCredentialPropertiesPtrOutputWithContext(ctx context.Context) WnsCredentialPropertiesPtrOutput {
	return o
}

func (o WnsCredentialPropertiesPtrOutput) Elem() WnsCredentialPropertiesOutput {
	return o.ApplyT(func(v *WnsCredentialProperties) WnsCredentialProperties {
		if v != nil {
			return *v
		}
		var ret WnsCredentialProperties
		return ret
	}).(WnsCredentialPropertiesOutput)
}

func (o WnsCredentialPropertiesPtrOutput) PackageSid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WnsCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.PackageSid
	}).(pulumi.StringPtrOutput)
}

func (o WnsCredentialPropertiesPtrOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WnsCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.SecretKey
	}).(pulumi.StringPtrOutput)
}

func (o WnsCredentialPropertiesPtrOutput) WindowsLiveEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WnsCredentialProperties) *string {
		if v == nil {
			return nil
		}
		return v.WindowsLiveEndpoint
	}).(pulumi.StringPtrOutput)
}

type WnsCredentialPropertiesResponse struct {
	PackageSid          *string `pulumi:"packageSid"`
	SecretKey           *string `pulumi:"secretKey"`
	WindowsLiveEndpoint *string `pulumi:"windowsLiveEndpoint"`
}





type WnsCredentialPropertiesResponseInput interface {
	pulumi.Input

	ToWnsCredentialPropertiesResponseOutput() WnsCredentialPropertiesResponseOutput
	ToWnsCredentialPropertiesResponseOutputWithContext(context.Context) WnsCredentialPropertiesResponseOutput
}

type WnsCredentialPropertiesResponseArgs struct {
	PackageSid          pulumi.StringPtrInput `pulumi:"packageSid"`
	SecretKey           pulumi.StringPtrInput `pulumi:"secretKey"`
	WindowsLiveEndpoint pulumi.StringPtrInput `pulumi:"windowsLiveEndpoint"`
}

func (WnsCredentialPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WnsCredentialPropertiesResponse)(nil)).Elem()
}

func (i WnsCredentialPropertiesResponseArgs) ToWnsCredentialPropertiesResponseOutput() WnsCredentialPropertiesResponseOutput {
	return i.ToWnsCredentialPropertiesResponseOutputWithContext(context.Background())
}

func (i WnsCredentialPropertiesResponseArgs) ToWnsCredentialPropertiesResponseOutputWithContext(ctx context.Context) WnsCredentialPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WnsCredentialPropertiesResponseOutput)
}

func (i WnsCredentialPropertiesResponseArgs) ToWnsCredentialPropertiesResponsePtrOutput() WnsCredentialPropertiesResponsePtrOutput {
	return i.ToWnsCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i WnsCredentialPropertiesResponseArgs) ToWnsCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) WnsCredentialPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WnsCredentialPropertiesResponseOutput).ToWnsCredentialPropertiesResponsePtrOutputWithContext(ctx)
}









type WnsCredentialPropertiesResponsePtrInput interface {
	pulumi.Input

	ToWnsCredentialPropertiesResponsePtrOutput() WnsCredentialPropertiesResponsePtrOutput
	ToWnsCredentialPropertiesResponsePtrOutputWithContext(context.Context) WnsCredentialPropertiesResponsePtrOutput
}

type wnsCredentialPropertiesResponsePtrType WnsCredentialPropertiesResponseArgs

func WnsCredentialPropertiesResponsePtr(v *WnsCredentialPropertiesResponseArgs) WnsCredentialPropertiesResponsePtrInput {
	return (*wnsCredentialPropertiesResponsePtrType)(v)
}

func (*wnsCredentialPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WnsCredentialPropertiesResponse)(nil)).Elem()
}

func (i *wnsCredentialPropertiesResponsePtrType) ToWnsCredentialPropertiesResponsePtrOutput() WnsCredentialPropertiesResponsePtrOutput {
	return i.ToWnsCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *wnsCredentialPropertiesResponsePtrType) ToWnsCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) WnsCredentialPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WnsCredentialPropertiesResponsePtrOutput)
}

type WnsCredentialPropertiesResponseOutput struct{ *pulumi.OutputState }

func (WnsCredentialPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WnsCredentialPropertiesResponse)(nil)).Elem()
}

func (o WnsCredentialPropertiesResponseOutput) ToWnsCredentialPropertiesResponseOutput() WnsCredentialPropertiesResponseOutput {
	return o
}

func (o WnsCredentialPropertiesResponseOutput) ToWnsCredentialPropertiesResponseOutputWithContext(ctx context.Context) WnsCredentialPropertiesResponseOutput {
	return o
}

func (o WnsCredentialPropertiesResponseOutput) ToWnsCredentialPropertiesResponsePtrOutput() WnsCredentialPropertiesResponsePtrOutput {
	return o.ToWnsCredentialPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o WnsCredentialPropertiesResponseOutput) ToWnsCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) WnsCredentialPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WnsCredentialPropertiesResponse) *WnsCredentialPropertiesResponse {
		return &v
	}).(WnsCredentialPropertiesResponsePtrOutput)
}

func (o WnsCredentialPropertiesResponseOutput) PackageSid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WnsCredentialPropertiesResponse) *string { return v.PackageSid }).(pulumi.StringPtrOutput)
}

func (o WnsCredentialPropertiesResponseOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WnsCredentialPropertiesResponse) *string { return v.SecretKey }).(pulumi.StringPtrOutput)
}

func (o WnsCredentialPropertiesResponseOutput) WindowsLiveEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WnsCredentialPropertiesResponse) *string { return v.WindowsLiveEndpoint }).(pulumi.StringPtrOutput)
}

type WnsCredentialPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (WnsCredentialPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WnsCredentialPropertiesResponse)(nil)).Elem()
}

func (o WnsCredentialPropertiesResponsePtrOutput) ToWnsCredentialPropertiesResponsePtrOutput() WnsCredentialPropertiesResponsePtrOutput {
	return o
}

func (o WnsCredentialPropertiesResponsePtrOutput) ToWnsCredentialPropertiesResponsePtrOutputWithContext(ctx context.Context) WnsCredentialPropertiesResponsePtrOutput {
	return o
}

func (o WnsCredentialPropertiesResponsePtrOutput) Elem() WnsCredentialPropertiesResponseOutput {
	return o.ApplyT(func(v *WnsCredentialPropertiesResponse) WnsCredentialPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret WnsCredentialPropertiesResponse
		return ret
	}).(WnsCredentialPropertiesResponseOutput)
}

func (o WnsCredentialPropertiesResponsePtrOutput) PackageSid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WnsCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PackageSid
	}).(pulumi.StringPtrOutput)
}

func (o WnsCredentialPropertiesResponsePtrOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WnsCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecretKey
	}).(pulumi.StringPtrOutput)
}

func (o WnsCredentialPropertiesResponsePtrOutput) WindowsLiveEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WnsCredentialPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.WindowsLiveEndpoint
	}).(pulumi.StringPtrOutput)
}

type WnsCredentialResponse struct {
	Properties *WnsCredentialPropertiesResponse `pulumi:"properties"`
}





type WnsCredentialResponseInput interface {
	pulumi.Input

	ToWnsCredentialResponseOutput() WnsCredentialResponseOutput
	ToWnsCredentialResponseOutputWithContext(context.Context) WnsCredentialResponseOutput
}

type WnsCredentialResponseArgs struct {
	Properties WnsCredentialPropertiesResponsePtrInput `pulumi:"properties"`
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

func (o WnsCredentialResponseOutput) Properties() WnsCredentialPropertiesResponsePtrOutput {
	return o.ApplyT(func(v WnsCredentialResponse) *WnsCredentialPropertiesResponse { return v.Properties }).(WnsCredentialPropertiesResponsePtrOutput)
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

func (o WnsCredentialResponsePtrOutput) Properties() WnsCredentialPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *WnsCredentialResponse) *WnsCredentialPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(WnsCredentialPropertiesResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AdmCredentialOutput{})
	pulumi.RegisterOutputType(AdmCredentialPtrOutput{})
	pulumi.RegisterOutputType(AdmCredentialPropertiesOutput{})
	pulumi.RegisterOutputType(AdmCredentialPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AdmCredentialPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AdmCredentialPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AdmCredentialResponseOutput{})
	pulumi.RegisterOutputType(AdmCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(ApnsCredentialOutput{})
	pulumi.RegisterOutputType(ApnsCredentialPtrOutput{})
	pulumi.RegisterOutputType(ApnsCredentialPropertiesOutput{})
	pulumi.RegisterOutputType(ApnsCredentialPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApnsCredentialPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApnsCredentialPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ApnsCredentialResponseOutput{})
	pulumi.RegisterOutputType(ApnsCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(BaiduCredentialOutput{})
	pulumi.RegisterOutputType(BaiduCredentialPtrOutput{})
	pulumi.RegisterOutputType(BaiduCredentialPropertiesOutput{})
	pulumi.RegisterOutputType(BaiduCredentialPropertiesPtrOutput{})
	pulumi.RegisterOutputType(BaiduCredentialPropertiesResponseOutput{})
	pulumi.RegisterOutputType(BaiduCredentialPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(BaiduCredentialResponseOutput{})
	pulumi.RegisterOutputType(BaiduCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(GcmCredentialOutput{})
	pulumi.RegisterOutputType(GcmCredentialPtrOutput{})
	pulumi.RegisterOutputType(GcmCredentialPropertiesOutput{})
	pulumi.RegisterOutputType(GcmCredentialPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GcmCredentialPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GcmCredentialPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(GcmCredentialResponseOutput{})
	pulumi.RegisterOutputType(GcmCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(MpnsCredentialOutput{})
	pulumi.RegisterOutputType(MpnsCredentialPtrOutput{})
	pulumi.RegisterOutputType(MpnsCredentialPropertiesOutput{})
	pulumi.RegisterOutputType(MpnsCredentialPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MpnsCredentialPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MpnsCredentialPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MpnsCredentialResponseOutput{})
	pulumi.RegisterOutputType(MpnsCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(NamespacePropertiesOutput{})
	pulumi.RegisterOutputType(NamespacePropertiesPtrOutput{})
	pulumi.RegisterOutputType(NamespacePropertiesResponseOutput{})
	pulumi.RegisterOutputType(NamespacePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NotificationHubPropertiesOutput{})
	pulumi.RegisterOutputType(NotificationHubPropertiesPtrOutput{})
	pulumi.RegisterOutputType(NotificationHubPropertiesResponseOutput{})
	pulumi.RegisterOutputType(NotificationHubPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SharedAccessAuthorizationRulePropertiesOutput{})
	pulumi.RegisterOutputType(SharedAccessAuthorizationRulePropertiesArrayOutput{})
	pulumi.RegisterOutputType(SharedAccessAuthorizationRulePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SharedAccessAuthorizationRulePropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(WnsCredentialOutput{})
	pulumi.RegisterOutputType(WnsCredentialPtrOutput{})
	pulumi.RegisterOutputType(WnsCredentialPropertiesOutput{})
	pulumi.RegisterOutputType(WnsCredentialPropertiesPtrOutput{})
	pulumi.RegisterOutputType(WnsCredentialPropertiesResponseOutput{})
	pulumi.RegisterOutputType(WnsCredentialPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(WnsCredentialResponseOutput{})
	pulumi.RegisterOutputType(WnsCredentialResponsePtrOutput{})
}

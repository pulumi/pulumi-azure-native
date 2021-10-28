


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppServiceCertificate struct {
	KeyVaultId         *string `pulumi:"keyVaultId"`
	KeyVaultSecretName *string `pulumi:"keyVaultSecretName"`
}





type AppServiceCertificateInput interface {
	pulumi.Input

	ToAppServiceCertificateOutput() AppServiceCertificateOutput
	ToAppServiceCertificateOutputWithContext(context.Context) AppServiceCertificateOutput
}

type AppServiceCertificateArgs struct {
	KeyVaultId         pulumi.StringPtrInput `pulumi:"keyVaultId"`
	KeyVaultSecretName pulumi.StringPtrInput `pulumi:"keyVaultSecretName"`
}

func (AppServiceCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServiceCertificate)(nil)).Elem()
}

func (i AppServiceCertificateArgs) ToAppServiceCertificateOutput() AppServiceCertificateOutput {
	return i.ToAppServiceCertificateOutputWithContext(context.Background())
}

func (i AppServiceCertificateArgs) ToAppServiceCertificateOutputWithContext(ctx context.Context) AppServiceCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceCertificateOutput)
}





type AppServiceCertificateMapInput interface {
	pulumi.Input

	ToAppServiceCertificateMapOutput() AppServiceCertificateMapOutput
	ToAppServiceCertificateMapOutputWithContext(context.Context) AppServiceCertificateMapOutput
}

type AppServiceCertificateMap map[string]AppServiceCertificateInput

func (AppServiceCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppServiceCertificate)(nil)).Elem()
}

func (i AppServiceCertificateMap) ToAppServiceCertificateMapOutput() AppServiceCertificateMapOutput {
	return i.ToAppServiceCertificateMapOutputWithContext(context.Background())
}

func (i AppServiceCertificateMap) ToAppServiceCertificateMapOutputWithContext(ctx context.Context) AppServiceCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceCertificateMapOutput)
}

type AppServiceCertificateOutput struct{ *pulumi.OutputState }

func (AppServiceCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServiceCertificate)(nil)).Elem()
}

func (o AppServiceCertificateOutput) ToAppServiceCertificateOutput() AppServiceCertificateOutput {
	return o
}

func (o AppServiceCertificateOutput) ToAppServiceCertificateOutputWithContext(ctx context.Context) AppServiceCertificateOutput {
	return o
}

func (o AppServiceCertificateOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppServiceCertificate) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o AppServiceCertificateOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppServiceCertificate) *string { return v.KeyVaultSecretName }).(pulumi.StringPtrOutput)
}

type AppServiceCertificateMapOutput struct{ *pulumi.OutputState }

func (AppServiceCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppServiceCertificate)(nil)).Elem()
}

func (o AppServiceCertificateMapOutput) ToAppServiceCertificateMapOutput() AppServiceCertificateMapOutput {
	return o
}

func (o AppServiceCertificateMapOutput) ToAppServiceCertificateMapOutputWithContext(ctx context.Context) AppServiceCertificateMapOutput {
	return o
}

func (o AppServiceCertificateMapOutput) MapIndex(k pulumi.StringInput) AppServiceCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppServiceCertificate {
		return vs[0].(map[string]AppServiceCertificate)[vs[1].(string)]
	}).(AppServiceCertificateOutput)
}

type AppServiceCertificateResponse struct {
	KeyVaultId         *string `pulumi:"keyVaultId"`
	KeyVaultSecretName *string `pulumi:"keyVaultSecretName"`
	ProvisioningState  string  `pulumi:"provisioningState"`
}





type AppServiceCertificateResponseInput interface {
	pulumi.Input

	ToAppServiceCertificateResponseOutput() AppServiceCertificateResponseOutput
	ToAppServiceCertificateResponseOutputWithContext(context.Context) AppServiceCertificateResponseOutput
}

type AppServiceCertificateResponseArgs struct {
	KeyVaultId         pulumi.StringPtrInput `pulumi:"keyVaultId"`
	KeyVaultSecretName pulumi.StringPtrInput `pulumi:"keyVaultSecretName"`
	ProvisioningState  pulumi.StringInput    `pulumi:"provisioningState"`
}

func (AppServiceCertificateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServiceCertificateResponse)(nil)).Elem()
}

func (i AppServiceCertificateResponseArgs) ToAppServiceCertificateResponseOutput() AppServiceCertificateResponseOutput {
	return i.ToAppServiceCertificateResponseOutputWithContext(context.Background())
}

func (i AppServiceCertificateResponseArgs) ToAppServiceCertificateResponseOutputWithContext(ctx context.Context) AppServiceCertificateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceCertificateResponseOutput)
}





type AppServiceCertificateResponseMapInput interface {
	pulumi.Input

	ToAppServiceCertificateResponseMapOutput() AppServiceCertificateResponseMapOutput
	ToAppServiceCertificateResponseMapOutputWithContext(context.Context) AppServiceCertificateResponseMapOutput
}

type AppServiceCertificateResponseMap map[string]AppServiceCertificateResponseInput

func (AppServiceCertificateResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppServiceCertificateResponse)(nil)).Elem()
}

func (i AppServiceCertificateResponseMap) ToAppServiceCertificateResponseMapOutput() AppServiceCertificateResponseMapOutput {
	return i.ToAppServiceCertificateResponseMapOutputWithContext(context.Background())
}

func (i AppServiceCertificateResponseMap) ToAppServiceCertificateResponseMapOutputWithContext(ctx context.Context) AppServiceCertificateResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceCertificateResponseMapOutput)
}

type AppServiceCertificateResponseOutput struct{ *pulumi.OutputState }

func (AppServiceCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServiceCertificateResponse)(nil)).Elem()
}

func (o AppServiceCertificateResponseOutput) ToAppServiceCertificateResponseOutput() AppServiceCertificateResponseOutput {
	return o
}

func (o AppServiceCertificateResponseOutput) ToAppServiceCertificateResponseOutputWithContext(ctx context.Context) AppServiceCertificateResponseOutput {
	return o
}

func (o AppServiceCertificateResponseOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppServiceCertificateResponse) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o AppServiceCertificateResponseOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppServiceCertificateResponse) *string { return v.KeyVaultSecretName }).(pulumi.StringPtrOutput)
}

func (o AppServiceCertificateResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AppServiceCertificateResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type AppServiceCertificateResponseMapOutput struct{ *pulumi.OutputState }

func (AppServiceCertificateResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppServiceCertificateResponse)(nil)).Elem()
}

func (o AppServiceCertificateResponseMapOutput) ToAppServiceCertificateResponseMapOutput() AppServiceCertificateResponseMapOutput {
	return o
}

func (o AppServiceCertificateResponseMapOutput) ToAppServiceCertificateResponseMapOutputWithContext(ctx context.Context) AppServiceCertificateResponseMapOutput {
	return o
}

func (o AppServiceCertificateResponseMapOutput) MapIndex(k pulumi.StringInput) AppServiceCertificateResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppServiceCertificateResponse {
		return vs[0].(map[string]AppServiceCertificateResponse)[vs[1].(string)]
	}).(AppServiceCertificateResponseOutput)
}

type CertificateDetailsResponse struct {
	Issuer             string `pulumi:"issuer"`
	NotAfter           string `pulumi:"notAfter"`
	NotBefore          string `pulumi:"notBefore"`
	RawData            string `pulumi:"rawData"`
	SerialNumber       string `pulumi:"serialNumber"`
	SignatureAlgorithm string `pulumi:"signatureAlgorithm"`
	Subject            string `pulumi:"subject"`
	Thumbprint         string `pulumi:"thumbprint"`
	Version            int    `pulumi:"version"`
}





type CertificateDetailsResponseInput interface {
	pulumi.Input

	ToCertificateDetailsResponseOutput() CertificateDetailsResponseOutput
	ToCertificateDetailsResponseOutputWithContext(context.Context) CertificateDetailsResponseOutput
}

type CertificateDetailsResponseArgs struct {
	Issuer             pulumi.StringInput `pulumi:"issuer"`
	NotAfter           pulumi.StringInput `pulumi:"notAfter"`
	NotBefore          pulumi.StringInput `pulumi:"notBefore"`
	RawData            pulumi.StringInput `pulumi:"rawData"`
	SerialNumber       pulumi.StringInput `pulumi:"serialNumber"`
	SignatureAlgorithm pulumi.StringInput `pulumi:"signatureAlgorithm"`
	Subject            pulumi.StringInput `pulumi:"subject"`
	Thumbprint         pulumi.StringInput `pulumi:"thumbprint"`
	Version            pulumi.IntInput    `pulumi:"version"`
}

func (CertificateDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateDetailsResponse)(nil)).Elem()
}

func (i CertificateDetailsResponseArgs) ToCertificateDetailsResponseOutput() CertificateDetailsResponseOutput {
	return i.ToCertificateDetailsResponseOutputWithContext(context.Background())
}

func (i CertificateDetailsResponseArgs) ToCertificateDetailsResponseOutputWithContext(ctx context.Context) CertificateDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateDetailsResponseOutput)
}

func (i CertificateDetailsResponseArgs) ToCertificateDetailsResponsePtrOutput() CertificateDetailsResponsePtrOutput {
	return i.ToCertificateDetailsResponsePtrOutputWithContext(context.Background())
}

func (i CertificateDetailsResponseArgs) ToCertificateDetailsResponsePtrOutputWithContext(ctx context.Context) CertificateDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateDetailsResponseOutput).ToCertificateDetailsResponsePtrOutputWithContext(ctx)
}









type CertificateDetailsResponsePtrInput interface {
	pulumi.Input

	ToCertificateDetailsResponsePtrOutput() CertificateDetailsResponsePtrOutput
	ToCertificateDetailsResponsePtrOutputWithContext(context.Context) CertificateDetailsResponsePtrOutput
}

type certificateDetailsResponsePtrType CertificateDetailsResponseArgs

func CertificateDetailsResponsePtr(v *CertificateDetailsResponseArgs) CertificateDetailsResponsePtrInput {
	return (*certificateDetailsResponsePtrType)(v)
}

func (*certificateDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateDetailsResponse)(nil)).Elem()
}

func (i *certificateDetailsResponsePtrType) ToCertificateDetailsResponsePtrOutput() CertificateDetailsResponsePtrOutput {
	return i.ToCertificateDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *certificateDetailsResponsePtrType) ToCertificateDetailsResponsePtrOutputWithContext(ctx context.Context) CertificateDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateDetailsResponsePtrOutput)
}

type CertificateDetailsResponseOutput struct{ *pulumi.OutputState }

func (CertificateDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateDetailsResponse)(nil)).Elem()
}

func (o CertificateDetailsResponseOutput) ToCertificateDetailsResponseOutput() CertificateDetailsResponseOutput {
	return o
}

func (o CertificateDetailsResponseOutput) ToCertificateDetailsResponseOutputWithContext(ctx context.Context) CertificateDetailsResponseOutput {
	return o
}

func (o CertificateDetailsResponseOutput) ToCertificateDetailsResponsePtrOutput() CertificateDetailsResponsePtrOutput {
	return o.ToCertificateDetailsResponsePtrOutputWithContext(context.Background())
}

func (o CertificateDetailsResponseOutput) ToCertificateDetailsResponsePtrOutputWithContext(ctx context.Context) CertificateDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateDetailsResponse) *CertificateDetailsResponse {
		return &v
	}).(CertificateDetailsResponsePtrOutput)
}

func (o CertificateDetailsResponseOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.Issuer }).(pulumi.StringOutput)
}

func (o CertificateDetailsResponseOutput) NotAfter() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.NotAfter }).(pulumi.StringOutput)
}

func (o CertificateDetailsResponseOutput) NotBefore() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.NotBefore }).(pulumi.StringOutput)
}

func (o CertificateDetailsResponseOutput) RawData() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.RawData }).(pulumi.StringOutput)
}

func (o CertificateDetailsResponseOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o CertificateDetailsResponseOutput) SignatureAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.SignatureAlgorithm }).(pulumi.StringOutput)
}

func (o CertificateDetailsResponseOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.Subject }).(pulumi.StringOutput)
}

func (o CertificateDetailsResponseOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificateDetailsResponseOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) int { return v.Version }).(pulumi.IntOutput)
}

type CertificateDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (CertificateDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateDetailsResponse)(nil)).Elem()
}

func (o CertificateDetailsResponsePtrOutput) ToCertificateDetailsResponsePtrOutput() CertificateDetailsResponsePtrOutput {
	return o
}

func (o CertificateDetailsResponsePtrOutput) ToCertificateDetailsResponsePtrOutputWithContext(ctx context.Context) CertificateDetailsResponsePtrOutput {
	return o
}

func (o CertificateDetailsResponsePtrOutput) Elem() CertificateDetailsResponseOutput {
	return o.ApplyT(func(v *CertificateDetailsResponse) CertificateDetailsResponse {
		if v != nil {
			return *v
		}
		var ret CertificateDetailsResponse
		return ret
	}).(CertificateDetailsResponseOutput)
}

func (o CertificateDetailsResponsePtrOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Issuer
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDetailsResponsePtrOutput) NotAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NotAfter
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDetailsResponsePtrOutput) NotBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NotBefore
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDetailsResponsePtrOutput) RawData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RawData
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDetailsResponsePtrOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SerialNumber
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDetailsResponsePtrOutput) SignatureAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SignatureAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDetailsResponsePtrOutput) Subject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Subject
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDetailsResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDetailsResponsePtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CertificateDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.IntPtrOutput)
}

type CertificateOrderContactResponse struct {
	Email     *string `pulumi:"email"`
	NameFirst *string `pulumi:"nameFirst"`
	NameLast  *string `pulumi:"nameLast"`
	Phone     *string `pulumi:"phone"`
}





type CertificateOrderContactResponseInput interface {
	pulumi.Input

	ToCertificateOrderContactResponseOutput() CertificateOrderContactResponseOutput
	ToCertificateOrderContactResponseOutputWithContext(context.Context) CertificateOrderContactResponseOutput
}

type CertificateOrderContactResponseArgs struct {
	Email     pulumi.StringPtrInput `pulumi:"email"`
	NameFirst pulumi.StringPtrInput `pulumi:"nameFirst"`
	NameLast  pulumi.StringPtrInput `pulumi:"nameLast"`
	Phone     pulumi.StringPtrInput `pulumi:"phone"`
}

func (CertificateOrderContactResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateOrderContactResponse)(nil)).Elem()
}

func (i CertificateOrderContactResponseArgs) ToCertificateOrderContactResponseOutput() CertificateOrderContactResponseOutput {
	return i.ToCertificateOrderContactResponseOutputWithContext(context.Background())
}

func (i CertificateOrderContactResponseArgs) ToCertificateOrderContactResponseOutputWithContext(ctx context.Context) CertificateOrderContactResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOrderContactResponseOutput)
}

func (i CertificateOrderContactResponseArgs) ToCertificateOrderContactResponsePtrOutput() CertificateOrderContactResponsePtrOutput {
	return i.ToCertificateOrderContactResponsePtrOutputWithContext(context.Background())
}

func (i CertificateOrderContactResponseArgs) ToCertificateOrderContactResponsePtrOutputWithContext(ctx context.Context) CertificateOrderContactResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOrderContactResponseOutput).ToCertificateOrderContactResponsePtrOutputWithContext(ctx)
}









type CertificateOrderContactResponsePtrInput interface {
	pulumi.Input

	ToCertificateOrderContactResponsePtrOutput() CertificateOrderContactResponsePtrOutput
	ToCertificateOrderContactResponsePtrOutputWithContext(context.Context) CertificateOrderContactResponsePtrOutput
}

type certificateOrderContactResponsePtrType CertificateOrderContactResponseArgs

func CertificateOrderContactResponsePtr(v *CertificateOrderContactResponseArgs) CertificateOrderContactResponsePtrInput {
	return (*certificateOrderContactResponsePtrType)(v)
}

func (*certificateOrderContactResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateOrderContactResponse)(nil)).Elem()
}

func (i *certificateOrderContactResponsePtrType) ToCertificateOrderContactResponsePtrOutput() CertificateOrderContactResponsePtrOutput {
	return i.ToCertificateOrderContactResponsePtrOutputWithContext(context.Background())
}

func (i *certificateOrderContactResponsePtrType) ToCertificateOrderContactResponsePtrOutputWithContext(ctx context.Context) CertificateOrderContactResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOrderContactResponsePtrOutput)
}

type CertificateOrderContactResponseOutput struct{ *pulumi.OutputState }

func (CertificateOrderContactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateOrderContactResponse)(nil)).Elem()
}

func (o CertificateOrderContactResponseOutput) ToCertificateOrderContactResponseOutput() CertificateOrderContactResponseOutput {
	return o
}

func (o CertificateOrderContactResponseOutput) ToCertificateOrderContactResponseOutputWithContext(ctx context.Context) CertificateOrderContactResponseOutput {
	return o
}

func (o CertificateOrderContactResponseOutput) ToCertificateOrderContactResponsePtrOutput() CertificateOrderContactResponsePtrOutput {
	return o.ToCertificateOrderContactResponsePtrOutputWithContext(context.Background())
}

func (o CertificateOrderContactResponseOutput) ToCertificateOrderContactResponsePtrOutputWithContext(ctx context.Context) CertificateOrderContactResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateOrderContactResponse) *CertificateOrderContactResponse {
		return &v
	}).(CertificateOrderContactResponsePtrOutput)
}

func (o CertificateOrderContactResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateOrderContactResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o CertificateOrderContactResponseOutput) NameFirst() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateOrderContactResponse) *string { return v.NameFirst }).(pulumi.StringPtrOutput)
}

func (o CertificateOrderContactResponseOutput) NameLast() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateOrderContactResponse) *string { return v.NameLast }).(pulumi.StringPtrOutput)
}

func (o CertificateOrderContactResponseOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateOrderContactResponse) *string { return v.Phone }).(pulumi.StringPtrOutput)
}

type CertificateOrderContactResponsePtrOutput struct{ *pulumi.OutputState }

func (CertificateOrderContactResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateOrderContactResponse)(nil)).Elem()
}

func (o CertificateOrderContactResponsePtrOutput) ToCertificateOrderContactResponsePtrOutput() CertificateOrderContactResponsePtrOutput {
	return o
}

func (o CertificateOrderContactResponsePtrOutput) ToCertificateOrderContactResponsePtrOutputWithContext(ctx context.Context) CertificateOrderContactResponsePtrOutput {
	return o
}

func (o CertificateOrderContactResponsePtrOutput) Elem() CertificateOrderContactResponseOutput {
	return o.ApplyT(func(v *CertificateOrderContactResponse) CertificateOrderContactResponse {
		if v != nil {
			return *v
		}
		var ret CertificateOrderContactResponse
		return ret
	}).(CertificateOrderContactResponseOutput)
}

func (o CertificateOrderContactResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateOrderContactResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o CertificateOrderContactResponsePtrOutput) NameFirst() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateOrderContactResponse) *string {
		if v == nil {
			return nil
		}
		return v.NameFirst
	}).(pulumi.StringPtrOutput)
}

func (o CertificateOrderContactResponsePtrOutput) NameLast() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateOrderContactResponse) *string {
		if v == nil {
			return nil
		}
		return v.NameLast
	}).(pulumi.StringPtrOutput)
}

func (o CertificateOrderContactResponsePtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateOrderContactResponse) *string {
		if v == nil {
			return nil
		}
		return v.Phone
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AppServiceCertificateOutput{})
	pulumi.RegisterOutputType(AppServiceCertificateMapOutput{})
	pulumi.RegisterOutputType(AppServiceCertificateResponseOutput{})
	pulumi.RegisterOutputType(AppServiceCertificateResponseMapOutput{})
	pulumi.RegisterOutputType(CertificateDetailsResponseOutput{})
	pulumi.RegisterOutputType(CertificateDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(CertificateOrderContactResponseOutput{})
	pulumi.RegisterOutputType(CertificateOrderContactResponsePtrOutput{})
}

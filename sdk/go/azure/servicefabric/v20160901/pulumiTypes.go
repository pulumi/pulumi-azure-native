


package v20160901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureActiveDirectory struct {
	ClientApplication  *string `pulumi:"clientApplication"`
	ClusterApplication *string `pulumi:"clusterApplication"`
	TenantId           *string `pulumi:"tenantId"`
}





type AzureActiveDirectoryInput interface {
	pulumi.Input

	ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput
	ToAzureActiveDirectoryOutputWithContext(context.Context) AzureActiveDirectoryOutput
}

type AzureActiveDirectoryArgs struct {
	ClientApplication  pulumi.StringPtrInput `pulumi:"clientApplication"`
	ClusterApplication pulumi.StringPtrInput `pulumi:"clusterApplication"`
	TenantId           pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (AzureActiveDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectory)(nil)).Elem()
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput {
	return i.ToAzureActiveDirectoryOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryOutputWithContext(ctx context.Context) AzureActiveDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryOutput)
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return i.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryOutput).ToAzureActiveDirectoryPtrOutputWithContext(ctx)
}









type AzureActiveDirectoryPtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput
	ToAzureActiveDirectoryPtrOutputWithContext(context.Context) AzureActiveDirectoryPtrOutput
}

type azureActiveDirectoryPtrType AzureActiveDirectoryArgs

func AzureActiveDirectoryPtr(v *AzureActiveDirectoryArgs) AzureActiveDirectoryPtrInput {
	return (*azureActiveDirectoryPtrType)(v)
}

func (*azureActiveDirectoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectory)(nil)).Elem()
}

func (i *azureActiveDirectoryPtrType) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return i.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryPtrType) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryPtrOutput)
}

type AzureActiveDirectoryOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectory)(nil)).Elem()
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput {
	return o
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryOutputWithContext(ctx context.Context) AzureActiveDirectoryOutput {
	return o
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return o.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectory) *AzureActiveDirectory {
		return &v
	}).(AzureActiveDirectoryPtrOutput)
}

func (o AzureActiveDirectoryOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *string { return v.ClientApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *string { return v.ClusterApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryPtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectory)(nil)).Elem()
}

func (o AzureActiveDirectoryPtrOutput) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return o
}

func (o AzureActiveDirectoryPtrOutput) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return o
}

func (o AzureActiveDirectoryPtrOutput) Elem() AzureActiveDirectoryOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) AzureActiveDirectory {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectory
		return ret
	}).(AzureActiveDirectoryOutput)
}

func (o AzureActiveDirectoryPtrOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.ClientApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.ClusterApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryResponse struct {
	ClientApplication  *string `pulumi:"clientApplication"`
	ClusterApplication *string `pulumi:"clusterApplication"`
	TenantId           *string `pulumi:"tenantId"`
}

type AzureActiveDirectoryResponseOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponseOutput() AzureActiveDirectoryResponseOutput {
	return o
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryResponseOutput {
	return o
}

func (o AzureActiveDirectoryResponseOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *string { return v.ClientApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *string { return v.ClusterApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryResponsePtrOutput) ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryResponsePtrOutput) ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryResponsePtrOutput) Elem() AzureActiveDirectoryResponseOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) AzureActiveDirectoryResponse {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryResponse
		return ret
	}).(AzureActiveDirectoryResponseOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClusterApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type CertificateDescription struct {
	Thumbprint          string  `pulumi:"thumbprint"`
	ThumbprintSecondary *string `pulumi:"thumbprintSecondary"`
	X509StoreName       *string `pulumi:"x509StoreName"`
}





type CertificateDescriptionInput interface {
	pulumi.Input

	ToCertificateDescriptionOutput() CertificateDescriptionOutput
	ToCertificateDescriptionOutputWithContext(context.Context) CertificateDescriptionOutput
}

type CertificateDescriptionArgs struct {
	Thumbprint          pulumi.StringInput    `pulumi:"thumbprint"`
	ThumbprintSecondary pulumi.StringPtrInput `pulumi:"thumbprintSecondary"`
	X509StoreName       pulumi.StringPtrInput `pulumi:"x509StoreName"`
}

func (CertificateDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateDescription)(nil)).Elem()
}

func (i CertificateDescriptionArgs) ToCertificateDescriptionOutput() CertificateDescriptionOutput {
	return i.ToCertificateDescriptionOutputWithContext(context.Background())
}

func (i CertificateDescriptionArgs) ToCertificateDescriptionOutputWithContext(ctx context.Context) CertificateDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateDescriptionOutput)
}

func (i CertificateDescriptionArgs) ToCertificateDescriptionPtrOutput() CertificateDescriptionPtrOutput {
	return i.ToCertificateDescriptionPtrOutputWithContext(context.Background())
}

func (i CertificateDescriptionArgs) ToCertificateDescriptionPtrOutputWithContext(ctx context.Context) CertificateDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateDescriptionOutput).ToCertificateDescriptionPtrOutputWithContext(ctx)
}









type CertificateDescriptionPtrInput interface {
	pulumi.Input

	ToCertificateDescriptionPtrOutput() CertificateDescriptionPtrOutput
	ToCertificateDescriptionPtrOutputWithContext(context.Context) CertificateDescriptionPtrOutput
}

type certificateDescriptionPtrType CertificateDescriptionArgs

func CertificateDescriptionPtr(v *CertificateDescriptionArgs) CertificateDescriptionPtrInput {
	return (*certificateDescriptionPtrType)(v)
}

func (*certificateDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateDescription)(nil)).Elem()
}

func (i *certificateDescriptionPtrType) ToCertificateDescriptionPtrOutput() CertificateDescriptionPtrOutput {
	return i.ToCertificateDescriptionPtrOutputWithContext(context.Background())
}

func (i *certificateDescriptionPtrType) ToCertificateDescriptionPtrOutputWithContext(ctx context.Context) CertificateDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateDescriptionPtrOutput)
}

type CertificateDescriptionOutput struct{ *pulumi.OutputState }

func (CertificateDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateDescription)(nil)).Elem()
}

func (o CertificateDescriptionOutput) ToCertificateDescriptionOutput() CertificateDescriptionOutput {
	return o
}

func (o CertificateDescriptionOutput) ToCertificateDescriptionOutputWithContext(ctx context.Context) CertificateDescriptionOutput {
	return o
}

func (o CertificateDescriptionOutput) ToCertificateDescriptionPtrOutput() CertificateDescriptionPtrOutput {
	return o.ToCertificateDescriptionPtrOutputWithContext(context.Background())
}

func (o CertificateDescriptionOutput) ToCertificateDescriptionPtrOutputWithContext(ctx context.Context) CertificateDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateDescription) *CertificateDescription {
		return &v
	}).(CertificateDescriptionPtrOutput)
}

func (o CertificateDescriptionOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDescription) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificateDescriptionOutput) ThumbprintSecondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateDescription) *string { return v.ThumbprintSecondary }).(pulumi.StringPtrOutput)
}

func (o CertificateDescriptionOutput) X509StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateDescription) *string { return v.X509StoreName }).(pulumi.StringPtrOutput)
}

type CertificateDescriptionPtrOutput struct{ *pulumi.OutputState }

func (CertificateDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateDescription)(nil)).Elem()
}

func (o CertificateDescriptionPtrOutput) ToCertificateDescriptionPtrOutput() CertificateDescriptionPtrOutput {
	return o
}

func (o CertificateDescriptionPtrOutput) ToCertificateDescriptionPtrOutputWithContext(ctx context.Context) CertificateDescriptionPtrOutput {
	return o
}

func (o CertificateDescriptionPtrOutput) Elem() CertificateDescriptionOutput {
	return o.ApplyT(func(v *CertificateDescription) CertificateDescription {
		if v != nil {
			return *v
		}
		var ret CertificateDescription
		return ret
	}).(CertificateDescriptionOutput)
}

func (o CertificateDescriptionPtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDescription) *string {
		if v == nil {
			return nil
		}
		return &v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDescriptionPtrOutput) ThumbprintSecondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDescription) *string {
		if v == nil {
			return nil
		}
		return v.ThumbprintSecondary
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDescriptionPtrOutput) X509StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDescription) *string {
		if v == nil {
			return nil
		}
		return v.X509StoreName
	}).(pulumi.StringPtrOutput)
}

type CertificateDescriptionResponse struct {
	Thumbprint          string  `pulumi:"thumbprint"`
	ThumbprintSecondary *string `pulumi:"thumbprintSecondary"`
	X509StoreName       *string `pulumi:"x509StoreName"`
}

type CertificateDescriptionResponseOutput struct{ *pulumi.OutputState }

func (CertificateDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateDescriptionResponse)(nil)).Elem()
}

func (o CertificateDescriptionResponseOutput) ToCertificateDescriptionResponseOutput() CertificateDescriptionResponseOutput {
	return o
}

func (o CertificateDescriptionResponseOutput) ToCertificateDescriptionResponseOutputWithContext(ctx context.Context) CertificateDescriptionResponseOutput {
	return o
}

func (o CertificateDescriptionResponseOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDescriptionResponse) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificateDescriptionResponseOutput) ThumbprintSecondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateDescriptionResponse) *string { return v.ThumbprintSecondary }).(pulumi.StringPtrOutput)
}

func (o CertificateDescriptionResponseOutput) X509StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateDescriptionResponse) *string { return v.X509StoreName }).(pulumi.StringPtrOutput)
}

type CertificateDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (CertificateDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateDescriptionResponse)(nil)).Elem()
}

func (o CertificateDescriptionResponsePtrOutput) ToCertificateDescriptionResponsePtrOutput() CertificateDescriptionResponsePtrOutput {
	return o
}

func (o CertificateDescriptionResponsePtrOutput) ToCertificateDescriptionResponsePtrOutputWithContext(ctx context.Context) CertificateDescriptionResponsePtrOutput {
	return o
}

func (o CertificateDescriptionResponsePtrOutput) Elem() CertificateDescriptionResponseOutput {
	return o.ApplyT(func(v *CertificateDescriptionResponse) CertificateDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret CertificateDescriptionResponse
		return ret
	}).(CertificateDescriptionResponseOutput)
}

func (o CertificateDescriptionResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDescriptionResponsePtrOutput) ThumbprintSecondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ThumbprintSecondary
	}).(pulumi.StringPtrOutput)
}

func (o CertificateDescriptionResponsePtrOutput) X509StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.X509StoreName
	}).(pulumi.StringPtrOutput)
}

type ClientCertificateCommonName struct {
	CertificateCommonName       string `pulumi:"certificateCommonName"`
	CertificateIssuerThumbprint string `pulumi:"certificateIssuerThumbprint"`
	IsAdmin                     bool   `pulumi:"isAdmin"`
}





type ClientCertificateCommonNameInput interface {
	pulumi.Input

	ToClientCertificateCommonNameOutput() ClientCertificateCommonNameOutput
	ToClientCertificateCommonNameOutputWithContext(context.Context) ClientCertificateCommonNameOutput
}

type ClientCertificateCommonNameArgs struct {
	CertificateCommonName       pulumi.StringInput `pulumi:"certificateCommonName"`
	CertificateIssuerThumbprint pulumi.StringInput `pulumi:"certificateIssuerThumbprint"`
	IsAdmin                     pulumi.BoolInput   `pulumi:"isAdmin"`
}

func (ClientCertificateCommonNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateCommonName)(nil)).Elem()
}

func (i ClientCertificateCommonNameArgs) ToClientCertificateCommonNameOutput() ClientCertificateCommonNameOutput {
	return i.ToClientCertificateCommonNameOutputWithContext(context.Background())
}

func (i ClientCertificateCommonNameArgs) ToClientCertificateCommonNameOutputWithContext(ctx context.Context) ClientCertificateCommonNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateCommonNameOutput)
}





type ClientCertificateCommonNameArrayInput interface {
	pulumi.Input

	ToClientCertificateCommonNameArrayOutput() ClientCertificateCommonNameArrayOutput
	ToClientCertificateCommonNameArrayOutputWithContext(context.Context) ClientCertificateCommonNameArrayOutput
}

type ClientCertificateCommonNameArray []ClientCertificateCommonNameInput

func (ClientCertificateCommonNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateCommonName)(nil)).Elem()
}

func (i ClientCertificateCommonNameArray) ToClientCertificateCommonNameArrayOutput() ClientCertificateCommonNameArrayOutput {
	return i.ToClientCertificateCommonNameArrayOutputWithContext(context.Background())
}

func (i ClientCertificateCommonNameArray) ToClientCertificateCommonNameArrayOutputWithContext(ctx context.Context) ClientCertificateCommonNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateCommonNameArrayOutput)
}

type ClientCertificateCommonNameOutput struct{ *pulumi.OutputState }

func (ClientCertificateCommonNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateCommonName)(nil)).Elem()
}

func (o ClientCertificateCommonNameOutput) ToClientCertificateCommonNameOutput() ClientCertificateCommonNameOutput {
	return o
}

func (o ClientCertificateCommonNameOutput) ToClientCertificateCommonNameOutputWithContext(ctx context.Context) ClientCertificateCommonNameOutput {
	return o
}

func (o ClientCertificateCommonNameOutput) CertificateCommonName() pulumi.StringOutput {
	return o.ApplyT(func(v ClientCertificateCommonName) string { return v.CertificateCommonName }).(pulumi.StringOutput)
}

func (o ClientCertificateCommonNameOutput) CertificateIssuerThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v ClientCertificateCommonName) string { return v.CertificateIssuerThumbprint }).(pulumi.StringOutput)
}

func (o ClientCertificateCommonNameOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificateCommonName) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

type ClientCertificateCommonNameArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateCommonNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateCommonName)(nil)).Elem()
}

func (o ClientCertificateCommonNameArrayOutput) ToClientCertificateCommonNameArrayOutput() ClientCertificateCommonNameArrayOutput {
	return o
}

func (o ClientCertificateCommonNameArrayOutput) ToClientCertificateCommonNameArrayOutputWithContext(ctx context.Context) ClientCertificateCommonNameArrayOutput {
	return o
}

func (o ClientCertificateCommonNameArrayOutput) Index(i pulumi.IntInput) ClientCertificateCommonNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificateCommonName {
		return vs[0].([]ClientCertificateCommonName)[vs[1].(int)]
	}).(ClientCertificateCommonNameOutput)
}

type ClientCertificateCommonNameResponse struct {
	CertificateCommonName       string `pulumi:"certificateCommonName"`
	CertificateIssuerThumbprint string `pulumi:"certificateIssuerThumbprint"`
	IsAdmin                     bool   `pulumi:"isAdmin"`
}

type ClientCertificateCommonNameResponseOutput struct{ *pulumi.OutputState }

func (ClientCertificateCommonNameResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateCommonNameResponse)(nil)).Elem()
}

func (o ClientCertificateCommonNameResponseOutput) ToClientCertificateCommonNameResponseOutput() ClientCertificateCommonNameResponseOutput {
	return o
}

func (o ClientCertificateCommonNameResponseOutput) ToClientCertificateCommonNameResponseOutputWithContext(ctx context.Context) ClientCertificateCommonNameResponseOutput {
	return o
}

func (o ClientCertificateCommonNameResponseOutput) CertificateCommonName() pulumi.StringOutput {
	return o.ApplyT(func(v ClientCertificateCommonNameResponse) string { return v.CertificateCommonName }).(pulumi.StringOutput)
}

func (o ClientCertificateCommonNameResponseOutput) CertificateIssuerThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v ClientCertificateCommonNameResponse) string { return v.CertificateIssuerThumbprint }).(pulumi.StringOutput)
}

func (o ClientCertificateCommonNameResponseOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificateCommonNameResponse) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

type ClientCertificateCommonNameResponseArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateCommonNameResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateCommonNameResponse)(nil)).Elem()
}

func (o ClientCertificateCommonNameResponseArrayOutput) ToClientCertificateCommonNameResponseArrayOutput() ClientCertificateCommonNameResponseArrayOutput {
	return o
}

func (o ClientCertificateCommonNameResponseArrayOutput) ToClientCertificateCommonNameResponseArrayOutputWithContext(ctx context.Context) ClientCertificateCommonNameResponseArrayOutput {
	return o
}

func (o ClientCertificateCommonNameResponseArrayOutput) Index(i pulumi.IntInput) ClientCertificateCommonNameResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificateCommonNameResponse {
		return vs[0].([]ClientCertificateCommonNameResponse)[vs[1].(int)]
	}).(ClientCertificateCommonNameResponseOutput)
}

type ClientCertificateThumbprint struct {
	CertificateThumbprint string `pulumi:"certificateThumbprint"`
	IsAdmin               bool   `pulumi:"isAdmin"`
}





type ClientCertificateThumbprintInput interface {
	pulumi.Input

	ToClientCertificateThumbprintOutput() ClientCertificateThumbprintOutput
	ToClientCertificateThumbprintOutputWithContext(context.Context) ClientCertificateThumbprintOutput
}

type ClientCertificateThumbprintArgs struct {
	CertificateThumbprint pulumi.StringInput `pulumi:"certificateThumbprint"`
	IsAdmin               pulumi.BoolInput   `pulumi:"isAdmin"`
}

func (ClientCertificateThumbprintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateThumbprint)(nil)).Elem()
}

func (i ClientCertificateThumbprintArgs) ToClientCertificateThumbprintOutput() ClientCertificateThumbprintOutput {
	return i.ToClientCertificateThumbprintOutputWithContext(context.Background())
}

func (i ClientCertificateThumbprintArgs) ToClientCertificateThumbprintOutputWithContext(ctx context.Context) ClientCertificateThumbprintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateThumbprintOutput)
}





type ClientCertificateThumbprintArrayInput interface {
	pulumi.Input

	ToClientCertificateThumbprintArrayOutput() ClientCertificateThumbprintArrayOutput
	ToClientCertificateThumbprintArrayOutputWithContext(context.Context) ClientCertificateThumbprintArrayOutput
}

type ClientCertificateThumbprintArray []ClientCertificateThumbprintInput

func (ClientCertificateThumbprintArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateThumbprint)(nil)).Elem()
}

func (i ClientCertificateThumbprintArray) ToClientCertificateThumbprintArrayOutput() ClientCertificateThumbprintArrayOutput {
	return i.ToClientCertificateThumbprintArrayOutputWithContext(context.Background())
}

func (i ClientCertificateThumbprintArray) ToClientCertificateThumbprintArrayOutputWithContext(ctx context.Context) ClientCertificateThumbprintArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateThumbprintArrayOutput)
}

type ClientCertificateThumbprintOutput struct{ *pulumi.OutputState }

func (ClientCertificateThumbprintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateThumbprint)(nil)).Elem()
}

func (o ClientCertificateThumbprintOutput) ToClientCertificateThumbprintOutput() ClientCertificateThumbprintOutput {
	return o
}

func (o ClientCertificateThumbprintOutput) ToClientCertificateThumbprintOutputWithContext(ctx context.Context) ClientCertificateThumbprintOutput {
	return o
}

func (o ClientCertificateThumbprintOutput) CertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v ClientCertificateThumbprint) string { return v.CertificateThumbprint }).(pulumi.StringOutput)
}

func (o ClientCertificateThumbprintOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificateThumbprint) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

type ClientCertificateThumbprintArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateThumbprintArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateThumbprint)(nil)).Elem()
}

func (o ClientCertificateThumbprintArrayOutput) ToClientCertificateThumbprintArrayOutput() ClientCertificateThumbprintArrayOutput {
	return o
}

func (o ClientCertificateThumbprintArrayOutput) ToClientCertificateThumbprintArrayOutputWithContext(ctx context.Context) ClientCertificateThumbprintArrayOutput {
	return o
}

func (o ClientCertificateThumbprintArrayOutput) Index(i pulumi.IntInput) ClientCertificateThumbprintOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificateThumbprint {
		return vs[0].([]ClientCertificateThumbprint)[vs[1].(int)]
	}).(ClientCertificateThumbprintOutput)
}

type ClientCertificateThumbprintResponse struct {
	CertificateThumbprint string `pulumi:"certificateThumbprint"`
	IsAdmin               bool   `pulumi:"isAdmin"`
}

type ClientCertificateThumbprintResponseOutput struct{ *pulumi.OutputState }

func (ClientCertificateThumbprintResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateThumbprintResponse)(nil)).Elem()
}

func (o ClientCertificateThumbprintResponseOutput) ToClientCertificateThumbprintResponseOutput() ClientCertificateThumbprintResponseOutput {
	return o
}

func (o ClientCertificateThumbprintResponseOutput) ToClientCertificateThumbprintResponseOutputWithContext(ctx context.Context) ClientCertificateThumbprintResponseOutput {
	return o
}

func (o ClientCertificateThumbprintResponseOutput) CertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v ClientCertificateThumbprintResponse) string { return v.CertificateThumbprint }).(pulumi.StringOutput)
}

func (o ClientCertificateThumbprintResponseOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificateThumbprintResponse) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

type ClientCertificateThumbprintResponseArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateThumbprintResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateThumbprintResponse)(nil)).Elem()
}

func (o ClientCertificateThumbprintResponseArrayOutput) ToClientCertificateThumbprintResponseArrayOutput() ClientCertificateThumbprintResponseArrayOutput {
	return o
}

func (o ClientCertificateThumbprintResponseArrayOutput) ToClientCertificateThumbprintResponseArrayOutputWithContext(ctx context.Context) ClientCertificateThumbprintResponseArrayOutput {
	return o
}

func (o ClientCertificateThumbprintResponseArrayOutput) Index(i pulumi.IntInput) ClientCertificateThumbprintResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificateThumbprintResponse {
		return vs[0].([]ClientCertificateThumbprintResponse)[vs[1].(int)]
	}).(ClientCertificateThumbprintResponseOutput)
}

type ClusterHealthPolicy struct {
	MaxPercentUnhealthyApplications *int `pulumi:"maxPercentUnhealthyApplications"`
	MaxPercentUnhealthyNodes        *int `pulumi:"maxPercentUnhealthyNodes"`
}





type ClusterHealthPolicyInput interface {
	pulumi.Input

	ToClusterHealthPolicyOutput() ClusterHealthPolicyOutput
	ToClusterHealthPolicyOutputWithContext(context.Context) ClusterHealthPolicyOutput
}

type ClusterHealthPolicyArgs struct {
	MaxPercentUnhealthyApplications pulumi.IntPtrInput `pulumi:"maxPercentUnhealthyApplications"`
	MaxPercentUnhealthyNodes        pulumi.IntPtrInput `pulumi:"maxPercentUnhealthyNodes"`
}

func (ClusterHealthPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterHealthPolicy)(nil)).Elem()
}

func (i ClusterHealthPolicyArgs) ToClusterHealthPolicyOutput() ClusterHealthPolicyOutput {
	return i.ToClusterHealthPolicyOutputWithContext(context.Background())
}

func (i ClusterHealthPolicyArgs) ToClusterHealthPolicyOutputWithContext(ctx context.Context) ClusterHealthPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterHealthPolicyOutput)
}

func (i ClusterHealthPolicyArgs) ToClusterHealthPolicyPtrOutput() ClusterHealthPolicyPtrOutput {
	return i.ToClusterHealthPolicyPtrOutputWithContext(context.Background())
}

func (i ClusterHealthPolicyArgs) ToClusterHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterHealthPolicyOutput).ToClusterHealthPolicyPtrOutputWithContext(ctx)
}









type ClusterHealthPolicyPtrInput interface {
	pulumi.Input

	ToClusterHealthPolicyPtrOutput() ClusterHealthPolicyPtrOutput
	ToClusterHealthPolicyPtrOutputWithContext(context.Context) ClusterHealthPolicyPtrOutput
}

type clusterHealthPolicyPtrType ClusterHealthPolicyArgs

func ClusterHealthPolicyPtr(v *ClusterHealthPolicyArgs) ClusterHealthPolicyPtrInput {
	return (*clusterHealthPolicyPtrType)(v)
}

func (*clusterHealthPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterHealthPolicy)(nil)).Elem()
}

func (i *clusterHealthPolicyPtrType) ToClusterHealthPolicyPtrOutput() ClusterHealthPolicyPtrOutput {
	return i.ToClusterHealthPolicyPtrOutputWithContext(context.Background())
}

func (i *clusterHealthPolicyPtrType) ToClusterHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterHealthPolicyPtrOutput)
}

type ClusterHealthPolicyOutput struct{ *pulumi.OutputState }

func (ClusterHealthPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterHealthPolicy)(nil)).Elem()
}

func (o ClusterHealthPolicyOutput) ToClusterHealthPolicyOutput() ClusterHealthPolicyOutput {
	return o
}

func (o ClusterHealthPolicyOutput) ToClusterHealthPolicyOutputWithContext(ctx context.Context) ClusterHealthPolicyOutput {
	return o
}

func (o ClusterHealthPolicyOutput) ToClusterHealthPolicyPtrOutput() ClusterHealthPolicyPtrOutput {
	return o.ToClusterHealthPolicyPtrOutputWithContext(context.Background())
}

func (o ClusterHealthPolicyOutput) ToClusterHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterHealthPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterHealthPolicy) *ClusterHealthPolicy {
		return &v
	}).(ClusterHealthPolicyPtrOutput)
}

func (o ClusterHealthPolicyOutput) MaxPercentUnhealthyApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterHealthPolicy) *int { return v.MaxPercentUnhealthyApplications }).(pulumi.IntPtrOutput)
}

func (o ClusterHealthPolicyOutput) MaxPercentUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterHealthPolicy) *int { return v.MaxPercentUnhealthyNodes }).(pulumi.IntPtrOutput)
}

type ClusterHealthPolicyPtrOutput struct{ *pulumi.OutputState }

func (ClusterHealthPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterHealthPolicy)(nil)).Elem()
}

func (o ClusterHealthPolicyPtrOutput) ToClusterHealthPolicyPtrOutput() ClusterHealthPolicyPtrOutput {
	return o
}

func (o ClusterHealthPolicyPtrOutput) ToClusterHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterHealthPolicyPtrOutput {
	return o
}

func (o ClusterHealthPolicyPtrOutput) Elem() ClusterHealthPolicyOutput {
	return o.ApplyT(func(v *ClusterHealthPolicy) ClusterHealthPolicy {
		if v != nil {
			return *v
		}
		var ret ClusterHealthPolicy
		return ret
	}).(ClusterHealthPolicyOutput)
}

func (o ClusterHealthPolicyPtrOutput) MaxPercentUnhealthyApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyApplications
	}).(pulumi.IntPtrOutput)
}

func (o ClusterHealthPolicyPtrOutput) MaxPercentUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyNodes
	}).(pulumi.IntPtrOutput)
}

type ClusterHealthPolicyResponse struct {
	MaxPercentUnhealthyApplications *int `pulumi:"maxPercentUnhealthyApplications"`
	MaxPercentUnhealthyNodes        *int `pulumi:"maxPercentUnhealthyNodes"`
}

type ClusterHealthPolicyResponseOutput struct{ *pulumi.OutputState }

func (ClusterHealthPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterHealthPolicyResponse)(nil)).Elem()
}

func (o ClusterHealthPolicyResponseOutput) ToClusterHealthPolicyResponseOutput() ClusterHealthPolicyResponseOutput {
	return o
}

func (o ClusterHealthPolicyResponseOutput) ToClusterHealthPolicyResponseOutputWithContext(ctx context.Context) ClusterHealthPolicyResponseOutput {
	return o
}

func (o ClusterHealthPolicyResponseOutput) MaxPercentUnhealthyApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterHealthPolicyResponse) *int { return v.MaxPercentUnhealthyApplications }).(pulumi.IntPtrOutput)
}

func (o ClusterHealthPolicyResponseOutput) MaxPercentUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterHealthPolicyResponse) *int { return v.MaxPercentUnhealthyNodes }).(pulumi.IntPtrOutput)
}

type ClusterHealthPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterHealthPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterHealthPolicyResponse)(nil)).Elem()
}

func (o ClusterHealthPolicyResponsePtrOutput) ToClusterHealthPolicyResponsePtrOutput() ClusterHealthPolicyResponsePtrOutput {
	return o
}

func (o ClusterHealthPolicyResponsePtrOutput) ToClusterHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ClusterHealthPolicyResponsePtrOutput {
	return o
}

func (o ClusterHealthPolicyResponsePtrOutput) Elem() ClusterHealthPolicyResponseOutput {
	return o.ApplyT(func(v *ClusterHealthPolicyResponse) ClusterHealthPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ClusterHealthPolicyResponse
		return ret
	}).(ClusterHealthPolicyResponseOutput)
}

func (o ClusterHealthPolicyResponsePtrOutput) MaxPercentUnhealthyApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyApplications
	}).(pulumi.IntPtrOutput)
}

func (o ClusterHealthPolicyResponsePtrOutput) MaxPercentUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPercentUnhealthyNodes
	}).(pulumi.IntPtrOutput)
}

type ClusterUpgradeDeltaHealthPolicy struct {
	MaxPercentDeltaUnhealthyApplications       int `pulumi:"maxPercentDeltaUnhealthyApplications"`
	MaxPercentDeltaUnhealthyNodes              int `pulumi:"maxPercentDeltaUnhealthyNodes"`
	MaxPercentUpgradeDomainDeltaUnhealthyNodes int `pulumi:"maxPercentUpgradeDomainDeltaUnhealthyNodes"`
}





type ClusterUpgradeDeltaHealthPolicyInput interface {
	pulumi.Input

	ToClusterUpgradeDeltaHealthPolicyOutput() ClusterUpgradeDeltaHealthPolicyOutput
	ToClusterUpgradeDeltaHealthPolicyOutputWithContext(context.Context) ClusterUpgradeDeltaHealthPolicyOutput
}

type ClusterUpgradeDeltaHealthPolicyArgs struct {
	MaxPercentDeltaUnhealthyApplications       pulumi.IntInput `pulumi:"maxPercentDeltaUnhealthyApplications"`
	MaxPercentDeltaUnhealthyNodes              pulumi.IntInput `pulumi:"maxPercentDeltaUnhealthyNodes"`
	MaxPercentUpgradeDomainDeltaUnhealthyNodes pulumi.IntInput `pulumi:"maxPercentUpgradeDomainDeltaUnhealthyNodes"`
}

func (ClusterUpgradeDeltaHealthPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradeDeltaHealthPolicy)(nil)).Elem()
}

func (i ClusterUpgradeDeltaHealthPolicyArgs) ToClusterUpgradeDeltaHealthPolicyOutput() ClusterUpgradeDeltaHealthPolicyOutput {
	return i.ToClusterUpgradeDeltaHealthPolicyOutputWithContext(context.Background())
}

func (i ClusterUpgradeDeltaHealthPolicyArgs) ToClusterUpgradeDeltaHealthPolicyOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradeDeltaHealthPolicyOutput)
}

func (i ClusterUpgradeDeltaHealthPolicyArgs) ToClusterUpgradeDeltaHealthPolicyPtrOutput() ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return i.ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(context.Background())
}

func (i ClusterUpgradeDeltaHealthPolicyArgs) ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradeDeltaHealthPolicyOutput).ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(ctx)
}









type ClusterUpgradeDeltaHealthPolicyPtrInput interface {
	pulumi.Input

	ToClusterUpgradeDeltaHealthPolicyPtrOutput() ClusterUpgradeDeltaHealthPolicyPtrOutput
	ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(context.Context) ClusterUpgradeDeltaHealthPolicyPtrOutput
}

type clusterUpgradeDeltaHealthPolicyPtrType ClusterUpgradeDeltaHealthPolicyArgs

func ClusterUpgradeDeltaHealthPolicyPtr(v *ClusterUpgradeDeltaHealthPolicyArgs) ClusterUpgradeDeltaHealthPolicyPtrInput {
	return (*clusterUpgradeDeltaHealthPolicyPtrType)(v)
}

func (*clusterUpgradeDeltaHealthPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradeDeltaHealthPolicy)(nil)).Elem()
}

func (i *clusterUpgradeDeltaHealthPolicyPtrType) ToClusterUpgradeDeltaHealthPolicyPtrOutput() ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return i.ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(context.Background())
}

func (i *clusterUpgradeDeltaHealthPolicyPtrType) ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradeDeltaHealthPolicyPtrOutput)
}

type ClusterUpgradeDeltaHealthPolicyOutput struct{ *pulumi.OutputState }

func (ClusterUpgradeDeltaHealthPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradeDeltaHealthPolicy)(nil)).Elem()
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) ToClusterUpgradeDeltaHealthPolicyOutput() ClusterUpgradeDeltaHealthPolicyOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) ToClusterUpgradeDeltaHealthPolicyOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) ToClusterUpgradeDeltaHealthPolicyPtrOutput() ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return o.ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(context.Background())
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterUpgradeDeltaHealthPolicy) *ClusterUpgradeDeltaHealthPolicy {
		return &v
	}).(ClusterUpgradeDeltaHealthPolicyPtrOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) MaxPercentDeltaUnhealthyApplications() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicy) int { return v.MaxPercentDeltaUnhealthyApplications }).(pulumi.IntOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) MaxPercentDeltaUnhealthyNodes() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicy) int { return v.MaxPercentDeltaUnhealthyNodes }).(pulumi.IntOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyOutput) MaxPercentUpgradeDomainDeltaUnhealthyNodes() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicy) int { return v.MaxPercentUpgradeDomainDeltaUnhealthyNodes }).(pulumi.IntOutput)
}

type ClusterUpgradeDeltaHealthPolicyPtrOutput struct{ *pulumi.OutputState }

func (ClusterUpgradeDeltaHealthPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradeDeltaHealthPolicy)(nil)).Elem()
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) ToClusterUpgradeDeltaHealthPolicyPtrOutput() ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) ToClusterUpgradeDeltaHealthPolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) Elem() ClusterUpgradeDeltaHealthPolicyOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicy) ClusterUpgradeDeltaHealthPolicy {
		if v != nil {
			return *v
		}
		var ret ClusterUpgradeDeltaHealthPolicy
		return ret
	}).(ClusterUpgradeDeltaHealthPolicyOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) MaxPercentDeltaUnhealthyApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentDeltaUnhealthyApplications
	}).(pulumi.IntPtrOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) MaxPercentDeltaUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentDeltaUnhealthyNodes
	}).(pulumi.IntPtrOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyPtrOutput) MaxPercentUpgradeDomainDeltaUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentUpgradeDomainDeltaUnhealthyNodes
	}).(pulumi.IntPtrOutput)
}

type ClusterUpgradeDeltaHealthPolicyResponse struct {
	MaxPercentDeltaUnhealthyApplications       int `pulumi:"maxPercentDeltaUnhealthyApplications"`
	MaxPercentDeltaUnhealthyNodes              int `pulumi:"maxPercentDeltaUnhealthyNodes"`
	MaxPercentUpgradeDomainDeltaUnhealthyNodes int `pulumi:"maxPercentUpgradeDomainDeltaUnhealthyNodes"`
}

type ClusterUpgradeDeltaHealthPolicyResponseOutput struct{ *pulumi.OutputState }

func (ClusterUpgradeDeltaHealthPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradeDeltaHealthPolicyResponse)(nil)).Elem()
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) ToClusterUpgradeDeltaHealthPolicyResponseOutput() ClusterUpgradeDeltaHealthPolicyResponseOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) ToClusterUpgradeDeltaHealthPolicyResponseOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyResponseOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) MaxPercentDeltaUnhealthyApplications() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicyResponse) int { return v.MaxPercentDeltaUnhealthyApplications }).(pulumi.IntOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) MaxPercentDeltaUnhealthyNodes() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicyResponse) int { return v.MaxPercentDeltaUnhealthyNodes }).(pulumi.IntOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponseOutput) MaxPercentUpgradeDomainDeltaUnhealthyNodes() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterUpgradeDeltaHealthPolicyResponse) int {
		return v.MaxPercentUpgradeDomainDeltaUnhealthyNodes
	}).(pulumi.IntOutput)
}

type ClusterUpgradeDeltaHealthPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradeDeltaHealthPolicyResponse)(nil)).Elem()
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) ToClusterUpgradeDeltaHealthPolicyResponsePtrOutput() ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) ToClusterUpgradeDeltaHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return o
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) Elem() ClusterUpgradeDeltaHealthPolicyResponseOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicyResponse) ClusterUpgradeDeltaHealthPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ClusterUpgradeDeltaHealthPolicyResponse
		return ret
	}).(ClusterUpgradeDeltaHealthPolicyResponseOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) MaxPercentDeltaUnhealthyApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentDeltaUnhealthyApplications
	}).(pulumi.IntPtrOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) MaxPercentDeltaUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentDeltaUnhealthyNodes
	}).(pulumi.IntPtrOutput)
}

func (o ClusterUpgradeDeltaHealthPolicyResponsePtrOutput) MaxPercentUpgradeDomainDeltaUnhealthyNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradeDeltaHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentUpgradeDomainDeltaUnhealthyNodes
	}).(pulumi.IntPtrOutput)
}

type ClusterUpgradePolicy struct {
	DeltaHealthPolicy             *ClusterUpgradeDeltaHealthPolicy `pulumi:"deltaHealthPolicy"`
	ForceRestart                  *bool                            `pulumi:"forceRestart"`
	HealthCheckRetryTimeout       string                           `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration     string                           `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration       string                           `pulumi:"healthCheckWaitDuration"`
	HealthPolicy                  ClusterHealthPolicy              `pulumi:"healthPolicy"`
	OverrideUserUpgradePolicy     *bool                            `pulumi:"overrideUserUpgradePolicy"`
	UpgradeDomainTimeout          string                           `pulumi:"upgradeDomainTimeout"`
	UpgradeReplicaSetCheckTimeout string                           `pulumi:"upgradeReplicaSetCheckTimeout"`
	UpgradeTimeout                string                           `pulumi:"upgradeTimeout"`
}





type ClusterUpgradePolicyInput interface {
	pulumi.Input

	ToClusterUpgradePolicyOutput() ClusterUpgradePolicyOutput
	ToClusterUpgradePolicyOutputWithContext(context.Context) ClusterUpgradePolicyOutput
}

type ClusterUpgradePolicyArgs struct {
	DeltaHealthPolicy             ClusterUpgradeDeltaHealthPolicyPtrInput `pulumi:"deltaHealthPolicy"`
	ForceRestart                  pulumi.BoolPtrInput                     `pulumi:"forceRestart"`
	HealthCheckRetryTimeout       pulumi.StringInput                      `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration     pulumi.StringInput                      `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration       pulumi.StringInput                      `pulumi:"healthCheckWaitDuration"`
	HealthPolicy                  ClusterHealthPolicyInput                `pulumi:"healthPolicy"`
	OverrideUserUpgradePolicy     pulumi.BoolPtrInput                     `pulumi:"overrideUserUpgradePolicy"`
	UpgradeDomainTimeout          pulumi.StringInput                      `pulumi:"upgradeDomainTimeout"`
	UpgradeReplicaSetCheckTimeout pulumi.StringInput                      `pulumi:"upgradeReplicaSetCheckTimeout"`
	UpgradeTimeout                pulumi.StringInput                      `pulumi:"upgradeTimeout"`
}

func (ClusterUpgradePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradePolicy)(nil)).Elem()
}

func (i ClusterUpgradePolicyArgs) ToClusterUpgradePolicyOutput() ClusterUpgradePolicyOutput {
	return i.ToClusterUpgradePolicyOutputWithContext(context.Background())
}

func (i ClusterUpgradePolicyArgs) ToClusterUpgradePolicyOutputWithContext(ctx context.Context) ClusterUpgradePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradePolicyOutput)
}

func (i ClusterUpgradePolicyArgs) ToClusterUpgradePolicyPtrOutput() ClusterUpgradePolicyPtrOutput {
	return i.ToClusterUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i ClusterUpgradePolicyArgs) ToClusterUpgradePolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradePolicyOutput).ToClusterUpgradePolicyPtrOutputWithContext(ctx)
}









type ClusterUpgradePolicyPtrInput interface {
	pulumi.Input

	ToClusterUpgradePolicyPtrOutput() ClusterUpgradePolicyPtrOutput
	ToClusterUpgradePolicyPtrOutputWithContext(context.Context) ClusterUpgradePolicyPtrOutput
}

type clusterUpgradePolicyPtrType ClusterUpgradePolicyArgs

func ClusterUpgradePolicyPtr(v *ClusterUpgradePolicyArgs) ClusterUpgradePolicyPtrInput {
	return (*clusterUpgradePolicyPtrType)(v)
}

func (*clusterUpgradePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradePolicy)(nil)).Elem()
}

func (i *clusterUpgradePolicyPtrType) ToClusterUpgradePolicyPtrOutput() ClusterUpgradePolicyPtrOutput {
	return i.ToClusterUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i *clusterUpgradePolicyPtrType) ToClusterUpgradePolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterUpgradePolicyPtrOutput)
}

type ClusterUpgradePolicyOutput struct{ *pulumi.OutputState }

func (ClusterUpgradePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradePolicy)(nil)).Elem()
}

func (o ClusterUpgradePolicyOutput) ToClusterUpgradePolicyOutput() ClusterUpgradePolicyOutput {
	return o
}

func (o ClusterUpgradePolicyOutput) ToClusterUpgradePolicyOutputWithContext(ctx context.Context) ClusterUpgradePolicyOutput {
	return o
}

func (o ClusterUpgradePolicyOutput) ToClusterUpgradePolicyPtrOutput() ClusterUpgradePolicyPtrOutput {
	return o.ToClusterUpgradePolicyPtrOutputWithContext(context.Background())
}

func (o ClusterUpgradePolicyOutput) ToClusterUpgradePolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterUpgradePolicy) *ClusterUpgradePolicy {
		return &v
	}).(ClusterUpgradePolicyPtrOutput)
}

func (o ClusterUpgradePolicyOutput) DeltaHealthPolicy() ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) *ClusterUpgradeDeltaHealthPolicy { return v.DeltaHealthPolicy }).(ClusterUpgradeDeltaHealthPolicyPtrOutput)
}

func (o ClusterUpgradePolicyOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) *bool { return v.ForceRestart }).(pulumi.BoolPtrOutput)
}

func (o ClusterUpgradePolicyOutput) HealthCheckRetryTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) string { return v.HealthCheckRetryTimeout }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyOutput) HealthCheckStableDuration() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) string { return v.HealthCheckStableDuration }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyOutput) HealthCheckWaitDuration() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) string { return v.HealthCheckWaitDuration }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyOutput) HealthPolicy() ClusterHealthPolicyOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) ClusterHealthPolicy { return v.HealthPolicy }).(ClusterHealthPolicyOutput)
}

func (o ClusterUpgradePolicyOutput) OverrideUserUpgradePolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) *bool { return v.OverrideUserUpgradePolicy }).(pulumi.BoolPtrOutput)
}

func (o ClusterUpgradePolicyOutput) UpgradeDomainTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) string { return v.UpgradeDomainTimeout }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyOutput) UpgradeReplicaSetCheckTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) string { return v.UpgradeReplicaSetCheckTimeout }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyOutput) UpgradeTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicy) string { return v.UpgradeTimeout }).(pulumi.StringOutput)
}

type ClusterUpgradePolicyPtrOutput struct{ *pulumi.OutputState }

func (ClusterUpgradePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradePolicy)(nil)).Elem()
}

func (o ClusterUpgradePolicyPtrOutput) ToClusterUpgradePolicyPtrOutput() ClusterUpgradePolicyPtrOutput {
	return o
}

func (o ClusterUpgradePolicyPtrOutput) ToClusterUpgradePolicyPtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyPtrOutput {
	return o
}

func (o ClusterUpgradePolicyPtrOutput) Elem() ClusterUpgradePolicyOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) ClusterUpgradePolicy {
		if v != nil {
			return *v
		}
		var ret ClusterUpgradePolicy
		return ret
	}).(ClusterUpgradePolicyOutput)
}

func (o ClusterUpgradePolicyPtrOutput) DeltaHealthPolicy() ClusterUpgradeDeltaHealthPolicyPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *ClusterUpgradeDeltaHealthPolicy {
		if v == nil {
			return nil
		}
		return v.DeltaHealthPolicy
	}).(ClusterUpgradeDeltaHealthPolicyPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.ForceRestart
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) HealthCheckRetryTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckRetryTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) HealthCheckStableDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckStableDuration
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) HealthCheckWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckWaitDuration
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) HealthPolicy() ClusterHealthPolicyPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *ClusterHealthPolicy {
		if v == nil {
			return nil
		}
		return &v.HealthPolicy
	}).(ClusterHealthPolicyPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) OverrideUserUpgradePolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.OverrideUserUpgradePolicy
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) UpgradeDomainTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeDomainTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) UpgradeReplicaSetCheckTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeReplicaSetCheckTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyPtrOutput) UpgradeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeTimeout
	}).(pulumi.StringPtrOutput)
}

type ClusterUpgradePolicyResponse struct {
	DeltaHealthPolicy             *ClusterUpgradeDeltaHealthPolicyResponse `pulumi:"deltaHealthPolicy"`
	ForceRestart                  *bool                                    `pulumi:"forceRestart"`
	HealthCheckRetryTimeout       string                                   `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration     string                                   `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration       string                                   `pulumi:"healthCheckWaitDuration"`
	HealthPolicy                  ClusterHealthPolicyResponse              `pulumi:"healthPolicy"`
	OverrideUserUpgradePolicy     *bool                                    `pulumi:"overrideUserUpgradePolicy"`
	UpgradeDomainTimeout          string                                   `pulumi:"upgradeDomainTimeout"`
	UpgradeReplicaSetCheckTimeout string                                   `pulumi:"upgradeReplicaSetCheckTimeout"`
	UpgradeTimeout                string                                   `pulumi:"upgradeTimeout"`
}

type ClusterUpgradePolicyResponseOutput struct{ *pulumi.OutputState }

func (ClusterUpgradePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradePolicyResponse)(nil)).Elem()
}

func (o ClusterUpgradePolicyResponseOutput) ToClusterUpgradePolicyResponseOutput() ClusterUpgradePolicyResponseOutput {
	return o
}

func (o ClusterUpgradePolicyResponseOutput) ToClusterUpgradePolicyResponseOutputWithContext(ctx context.Context) ClusterUpgradePolicyResponseOutput {
	return o
}

func (o ClusterUpgradePolicyResponseOutput) DeltaHealthPolicy() ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) *ClusterUpgradeDeltaHealthPolicyResponse {
		return v.DeltaHealthPolicy
	}).(ClusterUpgradeDeltaHealthPolicyResponsePtrOutput)
}

func (o ClusterUpgradePolicyResponseOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) *bool { return v.ForceRestart }).(pulumi.BoolPtrOutput)
}

func (o ClusterUpgradePolicyResponseOutput) HealthCheckRetryTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) string { return v.HealthCheckRetryTimeout }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyResponseOutput) HealthCheckStableDuration() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) string { return v.HealthCheckStableDuration }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyResponseOutput) HealthCheckWaitDuration() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) string { return v.HealthCheckWaitDuration }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyResponseOutput) HealthPolicy() ClusterHealthPolicyResponseOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) ClusterHealthPolicyResponse { return v.HealthPolicy }).(ClusterHealthPolicyResponseOutput)
}

func (o ClusterUpgradePolicyResponseOutput) OverrideUserUpgradePolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) *bool { return v.OverrideUserUpgradePolicy }).(pulumi.BoolPtrOutput)
}

func (o ClusterUpgradePolicyResponseOutput) UpgradeDomainTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) string { return v.UpgradeDomainTimeout }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyResponseOutput) UpgradeReplicaSetCheckTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) string { return v.UpgradeReplicaSetCheckTimeout }).(pulumi.StringOutput)
}

func (o ClusterUpgradePolicyResponseOutput) UpgradeTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterUpgradePolicyResponse) string { return v.UpgradeTimeout }).(pulumi.StringOutput)
}

type ClusterUpgradePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterUpgradePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradePolicyResponse)(nil)).Elem()
}

func (o ClusterUpgradePolicyResponsePtrOutput) ToClusterUpgradePolicyResponsePtrOutput() ClusterUpgradePolicyResponsePtrOutput {
	return o
}

func (o ClusterUpgradePolicyResponsePtrOutput) ToClusterUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ClusterUpgradePolicyResponsePtrOutput {
	return o
}

func (o ClusterUpgradePolicyResponsePtrOutput) Elem() ClusterUpgradePolicyResponseOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) ClusterUpgradePolicyResponse {
		if v != nil {
			return *v
		}
		var ret ClusterUpgradePolicyResponse
		return ret
	}).(ClusterUpgradePolicyResponseOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) DeltaHealthPolicy() ClusterUpgradeDeltaHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *ClusterUpgradeDeltaHealthPolicyResponse {
		if v == nil {
			return nil
		}
		return v.DeltaHealthPolicy
	}).(ClusterUpgradeDeltaHealthPolicyResponsePtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ForceRestart
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) HealthCheckRetryTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckRetryTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) HealthCheckStableDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckStableDuration
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) HealthCheckWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckWaitDuration
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) HealthPolicy() ClusterHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *ClusterHealthPolicyResponse {
		if v == nil {
			return nil
		}
		return &v.HealthPolicy
	}).(ClusterHealthPolicyResponsePtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) OverrideUserUpgradePolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.OverrideUserUpgradePolicy
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) UpgradeDomainTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeDomainTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) UpgradeReplicaSetCheckTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeReplicaSetCheckTimeout
	}).(pulumi.StringPtrOutput)
}

func (o ClusterUpgradePolicyResponsePtrOutput) UpgradeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeTimeout
	}).(pulumi.StringPtrOutput)
}

type ClusterVersionDetailsResponse struct {
	CodeVersion      *string `pulumi:"codeVersion"`
	Environment      *string `pulumi:"environment"`
	SupportExpiryUtc *string `pulumi:"supportExpiryUtc"`
}

type ClusterVersionDetailsResponseOutput struct{ *pulumi.OutputState }

func (ClusterVersionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterVersionDetailsResponse)(nil)).Elem()
}

func (o ClusterVersionDetailsResponseOutput) ToClusterVersionDetailsResponseOutput() ClusterVersionDetailsResponseOutput {
	return o
}

func (o ClusterVersionDetailsResponseOutput) ToClusterVersionDetailsResponseOutputWithContext(ctx context.Context) ClusterVersionDetailsResponseOutput {
	return o
}

func (o ClusterVersionDetailsResponseOutput) CodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterVersionDetailsResponse) *string { return v.CodeVersion }).(pulumi.StringPtrOutput)
}

func (o ClusterVersionDetailsResponseOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterVersionDetailsResponse) *string { return v.Environment }).(pulumi.StringPtrOutput)
}

func (o ClusterVersionDetailsResponseOutput) SupportExpiryUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterVersionDetailsResponse) *string { return v.SupportExpiryUtc }).(pulumi.StringPtrOutput)
}

type ClusterVersionDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (ClusterVersionDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterVersionDetailsResponse)(nil)).Elem()
}

func (o ClusterVersionDetailsResponseArrayOutput) ToClusterVersionDetailsResponseArrayOutput() ClusterVersionDetailsResponseArrayOutput {
	return o
}

func (o ClusterVersionDetailsResponseArrayOutput) ToClusterVersionDetailsResponseArrayOutputWithContext(ctx context.Context) ClusterVersionDetailsResponseArrayOutput {
	return o
}

func (o ClusterVersionDetailsResponseArrayOutput) Index(i pulumi.IntInput) ClusterVersionDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterVersionDetailsResponse {
		return vs[0].([]ClusterVersionDetailsResponse)[vs[1].(int)]
	}).(ClusterVersionDetailsResponseOutput)
}

type DiagnosticsStorageAccountConfig struct {
	BlobEndpoint            string `pulumi:"blobEndpoint"`
	ProtectedAccountKeyName string `pulumi:"protectedAccountKeyName"`
	QueueEndpoint           string `pulumi:"queueEndpoint"`
	StorageAccountName      string `pulumi:"storageAccountName"`
	TableEndpoint           string `pulumi:"tableEndpoint"`
}





type DiagnosticsStorageAccountConfigInput interface {
	pulumi.Input

	ToDiagnosticsStorageAccountConfigOutput() DiagnosticsStorageAccountConfigOutput
	ToDiagnosticsStorageAccountConfigOutputWithContext(context.Context) DiagnosticsStorageAccountConfigOutput
}

type DiagnosticsStorageAccountConfigArgs struct {
	BlobEndpoint            pulumi.StringInput `pulumi:"blobEndpoint"`
	ProtectedAccountKeyName pulumi.StringInput `pulumi:"protectedAccountKeyName"`
	QueueEndpoint           pulumi.StringInput `pulumi:"queueEndpoint"`
	StorageAccountName      pulumi.StringInput `pulumi:"storageAccountName"`
	TableEndpoint           pulumi.StringInput `pulumi:"tableEndpoint"`
}

func (DiagnosticsStorageAccountConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsStorageAccountConfig)(nil)).Elem()
}

func (i DiagnosticsStorageAccountConfigArgs) ToDiagnosticsStorageAccountConfigOutput() DiagnosticsStorageAccountConfigOutput {
	return i.ToDiagnosticsStorageAccountConfigOutputWithContext(context.Background())
}

func (i DiagnosticsStorageAccountConfigArgs) ToDiagnosticsStorageAccountConfigOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsStorageAccountConfigOutput)
}

func (i DiagnosticsStorageAccountConfigArgs) ToDiagnosticsStorageAccountConfigPtrOutput() DiagnosticsStorageAccountConfigPtrOutput {
	return i.ToDiagnosticsStorageAccountConfigPtrOutputWithContext(context.Background())
}

func (i DiagnosticsStorageAccountConfigArgs) ToDiagnosticsStorageAccountConfigPtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsStorageAccountConfigOutput).ToDiagnosticsStorageAccountConfigPtrOutputWithContext(ctx)
}









type DiagnosticsStorageAccountConfigPtrInput interface {
	pulumi.Input

	ToDiagnosticsStorageAccountConfigPtrOutput() DiagnosticsStorageAccountConfigPtrOutput
	ToDiagnosticsStorageAccountConfigPtrOutputWithContext(context.Context) DiagnosticsStorageAccountConfigPtrOutput
}

type diagnosticsStorageAccountConfigPtrType DiagnosticsStorageAccountConfigArgs

func DiagnosticsStorageAccountConfigPtr(v *DiagnosticsStorageAccountConfigArgs) DiagnosticsStorageAccountConfigPtrInput {
	return (*diagnosticsStorageAccountConfigPtrType)(v)
}

func (*diagnosticsStorageAccountConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsStorageAccountConfig)(nil)).Elem()
}

func (i *diagnosticsStorageAccountConfigPtrType) ToDiagnosticsStorageAccountConfigPtrOutput() DiagnosticsStorageAccountConfigPtrOutput {
	return i.ToDiagnosticsStorageAccountConfigPtrOutputWithContext(context.Background())
}

func (i *diagnosticsStorageAccountConfigPtrType) ToDiagnosticsStorageAccountConfigPtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsStorageAccountConfigPtrOutput)
}

type DiagnosticsStorageAccountConfigOutput struct{ *pulumi.OutputState }

func (DiagnosticsStorageAccountConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsStorageAccountConfig)(nil)).Elem()
}

func (o DiagnosticsStorageAccountConfigOutput) ToDiagnosticsStorageAccountConfigOutput() DiagnosticsStorageAccountConfigOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigOutput) ToDiagnosticsStorageAccountConfigOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigOutput) ToDiagnosticsStorageAccountConfigPtrOutput() DiagnosticsStorageAccountConfigPtrOutput {
	return o.ToDiagnosticsStorageAccountConfigPtrOutputWithContext(context.Background())
}

func (o DiagnosticsStorageAccountConfigOutput) ToDiagnosticsStorageAccountConfigPtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsStorageAccountConfig) *DiagnosticsStorageAccountConfig {
		return &v
	}).(DiagnosticsStorageAccountConfigPtrOutput)
}

func (o DiagnosticsStorageAccountConfigOutput) BlobEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfig) string { return v.BlobEndpoint }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigOutput) ProtectedAccountKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfig) string { return v.ProtectedAccountKeyName }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigOutput) QueueEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfig) string { return v.QueueEndpoint }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfig) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigOutput) TableEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfig) string { return v.TableEndpoint }).(pulumi.StringOutput)
}

type DiagnosticsStorageAccountConfigPtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsStorageAccountConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsStorageAccountConfig)(nil)).Elem()
}

func (o DiagnosticsStorageAccountConfigPtrOutput) ToDiagnosticsStorageAccountConfigPtrOutput() DiagnosticsStorageAccountConfigPtrOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigPtrOutput) ToDiagnosticsStorageAccountConfigPtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigPtrOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigPtrOutput) Elem() DiagnosticsStorageAccountConfigOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) DiagnosticsStorageAccountConfig {
		if v != nil {
			return *v
		}
		var ret DiagnosticsStorageAccountConfig
		return ret
	}).(DiagnosticsStorageAccountConfigOutput)
}

func (o DiagnosticsStorageAccountConfigPtrOutput) BlobEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) *string {
		if v == nil {
			return nil
		}
		return &v.BlobEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigPtrOutput) ProtectedAccountKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) *string {
		if v == nil {
			return nil
		}
		return &v.ProtectedAccountKeyName
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigPtrOutput) QueueEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) *string {
		if v == nil {
			return nil
		}
		return &v.QueueEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigPtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigPtrOutput) TableEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfig) *string {
		if v == nil {
			return nil
		}
		return &v.TableEndpoint
	}).(pulumi.StringPtrOutput)
}

type DiagnosticsStorageAccountConfigResponse struct {
	BlobEndpoint            string `pulumi:"blobEndpoint"`
	ProtectedAccountKeyName string `pulumi:"protectedAccountKeyName"`
	QueueEndpoint           string `pulumi:"queueEndpoint"`
	StorageAccountName      string `pulumi:"storageAccountName"`
	TableEndpoint           string `pulumi:"tableEndpoint"`
}

type DiagnosticsStorageAccountConfigResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsStorageAccountConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsStorageAccountConfigResponse)(nil)).Elem()
}

func (o DiagnosticsStorageAccountConfigResponseOutput) ToDiagnosticsStorageAccountConfigResponseOutput() DiagnosticsStorageAccountConfigResponseOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigResponseOutput) ToDiagnosticsStorageAccountConfigResponseOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigResponseOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigResponseOutput) BlobEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfigResponse) string { return v.BlobEndpoint }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigResponseOutput) ProtectedAccountKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfigResponse) string { return v.ProtectedAccountKeyName }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigResponseOutput) QueueEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfigResponse) string { return v.QueueEndpoint }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigResponseOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfigResponse) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o DiagnosticsStorageAccountConfigResponseOutput) TableEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsStorageAccountConfigResponse) string { return v.TableEndpoint }).(pulumi.StringOutput)
}

type DiagnosticsStorageAccountConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsStorageAccountConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsStorageAccountConfigResponse)(nil)).Elem()
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) ToDiagnosticsStorageAccountConfigResponsePtrOutput() DiagnosticsStorageAccountConfigResponsePtrOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) ToDiagnosticsStorageAccountConfigResponsePtrOutputWithContext(ctx context.Context) DiagnosticsStorageAccountConfigResponsePtrOutput {
	return o
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) Elem() DiagnosticsStorageAccountConfigResponseOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) DiagnosticsStorageAccountConfigResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticsStorageAccountConfigResponse
		return ret
	}).(DiagnosticsStorageAccountConfigResponseOutput)
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) BlobEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.BlobEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) ProtectedAccountKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProtectedAccountKeyName
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) QueueEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.QueueEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsStorageAccountConfigResponsePtrOutput) TableEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsStorageAccountConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TableEndpoint
	}).(pulumi.StringPtrOutput)
}

type EndpointRangeDescription struct {
	EndPort   int `pulumi:"endPort"`
	StartPort int `pulumi:"startPort"`
}





type EndpointRangeDescriptionInput interface {
	pulumi.Input

	ToEndpointRangeDescriptionOutput() EndpointRangeDescriptionOutput
	ToEndpointRangeDescriptionOutputWithContext(context.Context) EndpointRangeDescriptionOutput
}

type EndpointRangeDescriptionArgs struct {
	EndPort   pulumi.IntInput `pulumi:"endPort"`
	StartPort pulumi.IntInput `pulumi:"startPort"`
}

func (EndpointRangeDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescription)(nil)).Elem()
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionOutput() EndpointRangeDescriptionOutput {
	return i.ToEndpointRangeDescriptionOutputWithContext(context.Background())
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionOutputWithContext(ctx context.Context) EndpointRangeDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionOutput)
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return i.ToEndpointRangeDescriptionPtrOutputWithContext(context.Background())
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionOutput).ToEndpointRangeDescriptionPtrOutputWithContext(ctx)
}









type EndpointRangeDescriptionPtrInput interface {
	pulumi.Input

	ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput
	ToEndpointRangeDescriptionPtrOutputWithContext(context.Context) EndpointRangeDescriptionPtrOutput
}

type endpointRangeDescriptionPtrType EndpointRangeDescriptionArgs

func EndpointRangeDescriptionPtr(v *EndpointRangeDescriptionArgs) EndpointRangeDescriptionPtrInput {
	return (*endpointRangeDescriptionPtrType)(v)
}

func (*endpointRangeDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescription)(nil)).Elem()
}

func (i *endpointRangeDescriptionPtrType) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return i.ToEndpointRangeDescriptionPtrOutputWithContext(context.Background())
}

func (i *endpointRangeDescriptionPtrType) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionPtrOutput)
}

type EndpointRangeDescriptionOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescription)(nil)).Elem()
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionOutput() EndpointRangeDescriptionOutput {
	return o
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionOutputWithContext(ctx context.Context) EndpointRangeDescriptionOutput {
	return o
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return o.ToEndpointRangeDescriptionPtrOutputWithContext(context.Background())
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointRangeDescription) *EndpointRangeDescription {
		return &v
	}).(EndpointRangeDescriptionPtrOutput)
}

func (o EndpointRangeDescriptionOutput) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescription) int { return v.EndPort }).(pulumi.IntOutput)
}

func (o EndpointRangeDescriptionOutput) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescription) int { return v.StartPort }).(pulumi.IntOutput)
}

type EndpointRangeDescriptionPtrOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescription)(nil)).Elem()
}

func (o EndpointRangeDescriptionPtrOutput) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return o
}

func (o EndpointRangeDescriptionPtrOutput) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return o
}

func (o EndpointRangeDescriptionPtrOutput) Elem() EndpointRangeDescriptionOutput {
	return o.ApplyT(func(v *EndpointRangeDescription) EndpointRangeDescription {
		if v != nil {
			return *v
		}
		var ret EndpointRangeDescription
		return ret
	}).(EndpointRangeDescriptionOutput)
}

func (o EndpointRangeDescriptionPtrOutput) EndPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescription) *int {
		if v == nil {
			return nil
		}
		return &v.EndPort
	}).(pulumi.IntPtrOutput)
}

func (o EndpointRangeDescriptionPtrOutput) StartPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescription) *int {
		if v == nil {
			return nil
		}
		return &v.StartPort
	}).(pulumi.IntPtrOutput)
}

type EndpointRangeDescriptionResponse struct {
	EndPort   int `pulumi:"endPort"`
	StartPort int `pulumi:"startPort"`
}

type EndpointRangeDescriptionResponseOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescriptionResponse)(nil)).Elem()
}

func (o EndpointRangeDescriptionResponseOutput) ToEndpointRangeDescriptionResponseOutput() EndpointRangeDescriptionResponseOutput {
	return o
}

func (o EndpointRangeDescriptionResponseOutput) ToEndpointRangeDescriptionResponseOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponseOutput {
	return o
}

func (o EndpointRangeDescriptionResponseOutput) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescriptionResponse) int { return v.EndPort }).(pulumi.IntOutput)
}

func (o EndpointRangeDescriptionResponseOutput) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescriptionResponse) int { return v.StartPort }).(pulumi.IntOutput)
}

type EndpointRangeDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescriptionResponse)(nil)).Elem()
}

func (o EndpointRangeDescriptionResponsePtrOutput) ToEndpointRangeDescriptionResponsePtrOutput() EndpointRangeDescriptionResponsePtrOutput {
	return o
}

func (o EndpointRangeDescriptionResponsePtrOutput) ToEndpointRangeDescriptionResponsePtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponsePtrOutput {
	return o
}

func (o EndpointRangeDescriptionResponsePtrOutput) Elem() EndpointRangeDescriptionResponseOutput {
	return o.ApplyT(func(v *EndpointRangeDescriptionResponse) EndpointRangeDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret EndpointRangeDescriptionResponse
		return ret
	}).(EndpointRangeDescriptionResponseOutput)
}

func (o EndpointRangeDescriptionResponsePtrOutput) EndPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return &v.EndPort
	}).(pulumi.IntPtrOutput)
}

func (o EndpointRangeDescriptionResponsePtrOutput) StartPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return &v.StartPort
	}).(pulumi.IntPtrOutput)
}

type NodeTypeDescription struct {
	ApplicationPorts             *EndpointRangeDescription `pulumi:"applicationPorts"`
	Capacities                   map[string]string         `pulumi:"capacities"`
	ClientConnectionEndpointPort int                       `pulumi:"clientConnectionEndpointPort"`
	DurabilityLevel              *string                   `pulumi:"durabilityLevel"`
	EphemeralPorts               *EndpointRangeDescription `pulumi:"ephemeralPorts"`
	HttpGatewayEndpointPort      int                       `pulumi:"httpGatewayEndpointPort"`
	IsPrimary                    bool                      `pulumi:"isPrimary"`
	Name                         string                    `pulumi:"name"`
	PlacementProperties          map[string]string         `pulumi:"placementProperties"`
	ReverseProxyEndpointPort     *int                      `pulumi:"reverseProxyEndpointPort"`
	VmInstanceCount              int                       `pulumi:"vmInstanceCount"`
}





type NodeTypeDescriptionInput interface {
	pulumi.Input

	ToNodeTypeDescriptionOutput() NodeTypeDescriptionOutput
	ToNodeTypeDescriptionOutputWithContext(context.Context) NodeTypeDescriptionOutput
}

type NodeTypeDescriptionArgs struct {
	ApplicationPorts             EndpointRangeDescriptionPtrInput `pulumi:"applicationPorts"`
	Capacities                   pulumi.StringMapInput            `pulumi:"capacities"`
	ClientConnectionEndpointPort pulumi.IntInput                  `pulumi:"clientConnectionEndpointPort"`
	DurabilityLevel              pulumi.StringPtrInput            `pulumi:"durabilityLevel"`
	EphemeralPorts               EndpointRangeDescriptionPtrInput `pulumi:"ephemeralPorts"`
	HttpGatewayEndpointPort      pulumi.IntInput                  `pulumi:"httpGatewayEndpointPort"`
	IsPrimary                    pulumi.BoolInput                 `pulumi:"isPrimary"`
	Name                         pulumi.StringInput               `pulumi:"name"`
	PlacementProperties          pulumi.StringMapInput            `pulumi:"placementProperties"`
	ReverseProxyEndpointPort     pulumi.IntPtrInput               `pulumi:"reverseProxyEndpointPort"`
	VmInstanceCount              pulumi.IntInput                  `pulumi:"vmInstanceCount"`
}

func (NodeTypeDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTypeDescription)(nil)).Elem()
}

func (i NodeTypeDescriptionArgs) ToNodeTypeDescriptionOutput() NodeTypeDescriptionOutput {
	return i.ToNodeTypeDescriptionOutputWithContext(context.Background())
}

func (i NodeTypeDescriptionArgs) ToNodeTypeDescriptionOutputWithContext(ctx context.Context) NodeTypeDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeDescriptionOutput)
}





type NodeTypeDescriptionArrayInput interface {
	pulumi.Input

	ToNodeTypeDescriptionArrayOutput() NodeTypeDescriptionArrayOutput
	ToNodeTypeDescriptionArrayOutputWithContext(context.Context) NodeTypeDescriptionArrayOutput
}

type NodeTypeDescriptionArray []NodeTypeDescriptionInput

func (NodeTypeDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeTypeDescription)(nil)).Elem()
}

func (i NodeTypeDescriptionArray) ToNodeTypeDescriptionArrayOutput() NodeTypeDescriptionArrayOutput {
	return i.ToNodeTypeDescriptionArrayOutputWithContext(context.Background())
}

func (i NodeTypeDescriptionArray) ToNodeTypeDescriptionArrayOutputWithContext(ctx context.Context) NodeTypeDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeDescriptionArrayOutput)
}

type NodeTypeDescriptionOutput struct{ *pulumi.OutputState }

func (NodeTypeDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTypeDescription)(nil)).Elem()
}

func (o NodeTypeDescriptionOutput) ToNodeTypeDescriptionOutput() NodeTypeDescriptionOutput {
	return o
}

func (o NodeTypeDescriptionOutput) ToNodeTypeDescriptionOutputWithContext(ctx context.Context) NodeTypeDescriptionOutput {
	return o
}

func (o NodeTypeDescriptionOutput) ApplicationPorts() EndpointRangeDescriptionPtrOutput {
	return o.ApplyT(func(v NodeTypeDescription) *EndpointRangeDescription { return v.ApplicationPorts }).(EndpointRangeDescriptionPtrOutput)
}

func (o NodeTypeDescriptionOutput) Capacities() pulumi.StringMapOutput {
	return o.ApplyT(func(v NodeTypeDescription) map[string]string { return v.Capacities }).(pulumi.StringMapOutput)
}

func (o NodeTypeDescriptionOutput) ClientConnectionEndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeDescription) int { return v.ClientConnectionEndpointPort }).(pulumi.IntOutput)
}

func (o NodeTypeDescriptionOutput) DurabilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeTypeDescription) *string { return v.DurabilityLevel }).(pulumi.StringPtrOutput)
}

func (o NodeTypeDescriptionOutput) EphemeralPorts() EndpointRangeDescriptionPtrOutput {
	return o.ApplyT(func(v NodeTypeDescription) *EndpointRangeDescription { return v.EphemeralPorts }).(EndpointRangeDescriptionPtrOutput)
}

func (o NodeTypeDescriptionOutput) HttpGatewayEndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeDescription) int { return v.HttpGatewayEndpointPort }).(pulumi.IntOutput)
}

func (o NodeTypeDescriptionOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v NodeTypeDescription) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

func (o NodeTypeDescriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NodeTypeDescription) string { return v.Name }).(pulumi.StringOutput)
}

func (o NodeTypeDescriptionOutput) PlacementProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v NodeTypeDescription) map[string]string { return v.PlacementProperties }).(pulumi.StringMapOutput)
}

func (o NodeTypeDescriptionOutput) ReverseProxyEndpointPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NodeTypeDescription) *int { return v.ReverseProxyEndpointPort }).(pulumi.IntPtrOutput)
}

func (o NodeTypeDescriptionOutput) VmInstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeDescription) int { return v.VmInstanceCount }).(pulumi.IntOutput)
}

type NodeTypeDescriptionArrayOutput struct{ *pulumi.OutputState }

func (NodeTypeDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeTypeDescription)(nil)).Elem()
}

func (o NodeTypeDescriptionArrayOutput) ToNodeTypeDescriptionArrayOutput() NodeTypeDescriptionArrayOutput {
	return o
}

func (o NodeTypeDescriptionArrayOutput) ToNodeTypeDescriptionArrayOutputWithContext(ctx context.Context) NodeTypeDescriptionArrayOutput {
	return o
}

func (o NodeTypeDescriptionArrayOutput) Index(i pulumi.IntInput) NodeTypeDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeTypeDescription {
		return vs[0].([]NodeTypeDescription)[vs[1].(int)]
	}).(NodeTypeDescriptionOutput)
}

type NodeTypeDescriptionResponse struct {
	ApplicationPorts             *EndpointRangeDescriptionResponse `pulumi:"applicationPorts"`
	Capacities                   map[string]string                 `pulumi:"capacities"`
	ClientConnectionEndpointPort int                               `pulumi:"clientConnectionEndpointPort"`
	DurabilityLevel              *string                           `pulumi:"durabilityLevel"`
	EphemeralPorts               *EndpointRangeDescriptionResponse `pulumi:"ephemeralPorts"`
	HttpGatewayEndpointPort      int                               `pulumi:"httpGatewayEndpointPort"`
	IsPrimary                    bool                              `pulumi:"isPrimary"`
	Name                         string                            `pulumi:"name"`
	PlacementProperties          map[string]string                 `pulumi:"placementProperties"`
	ReverseProxyEndpointPort     *int                              `pulumi:"reverseProxyEndpointPort"`
	VmInstanceCount              int                               `pulumi:"vmInstanceCount"`
}

type NodeTypeDescriptionResponseOutput struct{ *pulumi.OutputState }

func (NodeTypeDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTypeDescriptionResponse)(nil)).Elem()
}

func (o NodeTypeDescriptionResponseOutput) ToNodeTypeDescriptionResponseOutput() NodeTypeDescriptionResponseOutput {
	return o
}

func (o NodeTypeDescriptionResponseOutput) ToNodeTypeDescriptionResponseOutputWithContext(ctx context.Context) NodeTypeDescriptionResponseOutput {
	return o
}

func (o NodeTypeDescriptionResponseOutput) ApplicationPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) *EndpointRangeDescriptionResponse { return v.ApplicationPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o NodeTypeDescriptionResponseOutput) Capacities() pulumi.StringMapOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) map[string]string { return v.Capacities }).(pulumi.StringMapOutput)
}

func (o NodeTypeDescriptionResponseOutput) ClientConnectionEndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) int { return v.ClientConnectionEndpointPort }).(pulumi.IntOutput)
}

func (o NodeTypeDescriptionResponseOutput) DurabilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) *string { return v.DurabilityLevel }).(pulumi.StringPtrOutput)
}

func (o NodeTypeDescriptionResponseOutput) EphemeralPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) *EndpointRangeDescriptionResponse { return v.EphemeralPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o NodeTypeDescriptionResponseOutput) HttpGatewayEndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) int { return v.HttpGatewayEndpointPort }).(pulumi.IntOutput)
}

func (o NodeTypeDescriptionResponseOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

func (o NodeTypeDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o NodeTypeDescriptionResponseOutput) PlacementProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) map[string]string { return v.PlacementProperties }).(pulumi.StringMapOutput)
}

func (o NodeTypeDescriptionResponseOutput) ReverseProxyEndpointPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) *int { return v.ReverseProxyEndpointPort }).(pulumi.IntPtrOutput)
}

func (o NodeTypeDescriptionResponseOutput) VmInstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeDescriptionResponse) int { return v.VmInstanceCount }).(pulumi.IntOutput)
}

type NodeTypeDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (NodeTypeDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeTypeDescriptionResponse)(nil)).Elem()
}

func (o NodeTypeDescriptionResponseArrayOutput) ToNodeTypeDescriptionResponseArrayOutput() NodeTypeDescriptionResponseArrayOutput {
	return o
}

func (o NodeTypeDescriptionResponseArrayOutput) ToNodeTypeDescriptionResponseArrayOutputWithContext(ctx context.Context) NodeTypeDescriptionResponseArrayOutput {
	return o
}

func (o NodeTypeDescriptionResponseArrayOutput) Index(i pulumi.IntInput) NodeTypeDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeTypeDescriptionResponse {
		return vs[0].([]NodeTypeDescriptionResponse)[vs[1].(int)]
	}).(NodeTypeDescriptionResponseOutput)
}

type SettingsParameterDescription struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type SettingsParameterDescriptionInput interface {
	pulumi.Input

	ToSettingsParameterDescriptionOutput() SettingsParameterDescriptionOutput
	ToSettingsParameterDescriptionOutputWithContext(context.Context) SettingsParameterDescriptionOutput
}

type SettingsParameterDescriptionArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (SettingsParameterDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescription)(nil)).Elem()
}

func (i SettingsParameterDescriptionArgs) ToSettingsParameterDescriptionOutput() SettingsParameterDescriptionOutput {
	return i.ToSettingsParameterDescriptionOutputWithContext(context.Background())
}

func (i SettingsParameterDescriptionArgs) ToSettingsParameterDescriptionOutputWithContext(ctx context.Context) SettingsParameterDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsParameterDescriptionOutput)
}





type SettingsParameterDescriptionArrayInput interface {
	pulumi.Input

	ToSettingsParameterDescriptionArrayOutput() SettingsParameterDescriptionArrayOutput
	ToSettingsParameterDescriptionArrayOutputWithContext(context.Context) SettingsParameterDescriptionArrayOutput
}

type SettingsParameterDescriptionArray []SettingsParameterDescriptionInput

func (SettingsParameterDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescription)(nil)).Elem()
}

func (i SettingsParameterDescriptionArray) ToSettingsParameterDescriptionArrayOutput() SettingsParameterDescriptionArrayOutput {
	return i.ToSettingsParameterDescriptionArrayOutputWithContext(context.Background())
}

func (i SettingsParameterDescriptionArray) ToSettingsParameterDescriptionArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsParameterDescriptionArrayOutput)
}

type SettingsParameterDescriptionOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescription)(nil)).Elem()
}

func (o SettingsParameterDescriptionOutput) ToSettingsParameterDescriptionOutput() SettingsParameterDescriptionOutput {
	return o
}

func (o SettingsParameterDescriptionOutput) ToSettingsParameterDescriptionOutputWithContext(ctx context.Context) SettingsParameterDescriptionOutput {
	return o
}

func (o SettingsParameterDescriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescription) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsParameterDescriptionOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescription) string { return v.Value }).(pulumi.StringOutput)
}

type SettingsParameterDescriptionArrayOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescription)(nil)).Elem()
}

func (o SettingsParameterDescriptionArrayOutput) ToSettingsParameterDescriptionArrayOutput() SettingsParameterDescriptionArrayOutput {
	return o
}

func (o SettingsParameterDescriptionArrayOutput) ToSettingsParameterDescriptionArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionArrayOutput {
	return o
}

func (o SettingsParameterDescriptionArrayOutput) Index(i pulumi.IntInput) SettingsParameterDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsParameterDescription {
		return vs[0].([]SettingsParameterDescription)[vs[1].(int)]
	}).(SettingsParameterDescriptionOutput)
}

type SettingsParameterDescriptionResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type SettingsParameterDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescriptionResponse)(nil)).Elem()
}

func (o SettingsParameterDescriptionResponseOutput) ToSettingsParameterDescriptionResponseOutput() SettingsParameterDescriptionResponseOutput {
	return o
}

func (o SettingsParameterDescriptionResponseOutput) ToSettingsParameterDescriptionResponseOutputWithContext(ctx context.Context) SettingsParameterDescriptionResponseOutput {
	return o
}

func (o SettingsParameterDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsParameterDescriptionResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescriptionResponse) string { return v.Value }).(pulumi.StringOutput)
}

type SettingsParameterDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescriptionResponse)(nil)).Elem()
}

func (o SettingsParameterDescriptionResponseArrayOutput) ToSettingsParameterDescriptionResponseArrayOutput() SettingsParameterDescriptionResponseArrayOutput {
	return o
}

func (o SettingsParameterDescriptionResponseArrayOutput) ToSettingsParameterDescriptionResponseArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionResponseArrayOutput {
	return o
}

func (o SettingsParameterDescriptionResponseArrayOutput) Index(i pulumi.IntInput) SettingsParameterDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsParameterDescriptionResponse {
		return vs[0].([]SettingsParameterDescriptionResponse)[vs[1].(int)]
	}).(SettingsParameterDescriptionResponseOutput)
}

type SettingsSectionDescription struct {
	Name       string                         `pulumi:"name"`
	Parameters []SettingsParameterDescription `pulumi:"parameters"`
}





type SettingsSectionDescriptionInput interface {
	pulumi.Input

	ToSettingsSectionDescriptionOutput() SettingsSectionDescriptionOutput
	ToSettingsSectionDescriptionOutputWithContext(context.Context) SettingsSectionDescriptionOutput
}

type SettingsSectionDescriptionArgs struct {
	Name       pulumi.StringInput                     `pulumi:"name"`
	Parameters SettingsParameterDescriptionArrayInput `pulumi:"parameters"`
}

func (SettingsSectionDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescription)(nil)).Elem()
}

func (i SettingsSectionDescriptionArgs) ToSettingsSectionDescriptionOutput() SettingsSectionDescriptionOutput {
	return i.ToSettingsSectionDescriptionOutputWithContext(context.Background())
}

func (i SettingsSectionDescriptionArgs) ToSettingsSectionDescriptionOutputWithContext(ctx context.Context) SettingsSectionDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsSectionDescriptionOutput)
}





type SettingsSectionDescriptionArrayInput interface {
	pulumi.Input

	ToSettingsSectionDescriptionArrayOutput() SettingsSectionDescriptionArrayOutput
	ToSettingsSectionDescriptionArrayOutputWithContext(context.Context) SettingsSectionDescriptionArrayOutput
}

type SettingsSectionDescriptionArray []SettingsSectionDescriptionInput

func (SettingsSectionDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescription)(nil)).Elem()
}

func (i SettingsSectionDescriptionArray) ToSettingsSectionDescriptionArrayOutput() SettingsSectionDescriptionArrayOutput {
	return i.ToSettingsSectionDescriptionArrayOutputWithContext(context.Background())
}

func (i SettingsSectionDescriptionArray) ToSettingsSectionDescriptionArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsSectionDescriptionArrayOutput)
}

type SettingsSectionDescriptionOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescription)(nil)).Elem()
}

func (o SettingsSectionDescriptionOutput) ToSettingsSectionDescriptionOutput() SettingsSectionDescriptionOutput {
	return o
}

func (o SettingsSectionDescriptionOutput) ToSettingsSectionDescriptionOutputWithContext(ctx context.Context) SettingsSectionDescriptionOutput {
	return o
}

func (o SettingsSectionDescriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsSectionDescription) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsSectionDescriptionOutput) Parameters() SettingsParameterDescriptionArrayOutput {
	return o.ApplyT(func(v SettingsSectionDescription) []SettingsParameterDescription { return v.Parameters }).(SettingsParameterDescriptionArrayOutput)
}

type SettingsSectionDescriptionArrayOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescription)(nil)).Elem()
}

func (o SettingsSectionDescriptionArrayOutput) ToSettingsSectionDescriptionArrayOutput() SettingsSectionDescriptionArrayOutput {
	return o
}

func (o SettingsSectionDescriptionArrayOutput) ToSettingsSectionDescriptionArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionArrayOutput {
	return o
}

func (o SettingsSectionDescriptionArrayOutput) Index(i pulumi.IntInput) SettingsSectionDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsSectionDescription {
		return vs[0].([]SettingsSectionDescription)[vs[1].(int)]
	}).(SettingsSectionDescriptionOutput)
}

type SettingsSectionDescriptionResponse struct {
	Name       string                                 `pulumi:"name"`
	Parameters []SettingsParameterDescriptionResponse `pulumi:"parameters"`
}

type SettingsSectionDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescriptionResponse)(nil)).Elem()
}

func (o SettingsSectionDescriptionResponseOutput) ToSettingsSectionDescriptionResponseOutput() SettingsSectionDescriptionResponseOutput {
	return o
}

func (o SettingsSectionDescriptionResponseOutput) ToSettingsSectionDescriptionResponseOutputWithContext(ctx context.Context) SettingsSectionDescriptionResponseOutput {
	return o
}

func (o SettingsSectionDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsSectionDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsSectionDescriptionResponseOutput) Parameters() SettingsParameterDescriptionResponseArrayOutput {
	return o.ApplyT(func(v SettingsSectionDescriptionResponse) []SettingsParameterDescriptionResponse { return v.Parameters }).(SettingsParameterDescriptionResponseArrayOutput)
}

type SettingsSectionDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescriptionResponse)(nil)).Elem()
}

func (o SettingsSectionDescriptionResponseArrayOutput) ToSettingsSectionDescriptionResponseArrayOutput() SettingsSectionDescriptionResponseArrayOutput {
	return o
}

func (o SettingsSectionDescriptionResponseArrayOutput) ToSettingsSectionDescriptionResponseArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionResponseArrayOutput {
	return o
}

func (o SettingsSectionDescriptionResponseArrayOutput) Index(i pulumi.IntInput) SettingsSectionDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsSectionDescriptionResponse {
		return vs[0].([]SettingsSectionDescriptionResponse)[vs[1].(int)]
	}).(SettingsSectionDescriptionResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureActiveDirectoryOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryPtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryResponsePtrOutput{})
	pulumi.RegisterOutputType(CertificateDescriptionOutput{})
	pulumi.RegisterOutputType(CertificateDescriptionPtrOutput{})
	pulumi.RegisterOutputType(CertificateDescriptionResponseOutput{})
	pulumi.RegisterOutputType(CertificateDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(ClientCertificateCommonNameOutput{})
	pulumi.RegisterOutputType(ClientCertificateCommonNameArrayOutput{})
	pulumi.RegisterOutputType(ClientCertificateCommonNameResponseOutput{})
	pulumi.RegisterOutputType(ClientCertificateCommonNameResponseArrayOutput{})
	pulumi.RegisterOutputType(ClientCertificateThumbprintOutput{})
	pulumi.RegisterOutputType(ClientCertificateThumbprintArrayOutput{})
	pulumi.RegisterOutputType(ClientCertificateThumbprintResponseOutput{})
	pulumi.RegisterOutputType(ClientCertificateThumbprintResponseArrayOutput{})
	pulumi.RegisterOutputType(ClusterHealthPolicyOutput{})
	pulumi.RegisterOutputType(ClusterHealthPolicyPtrOutput{})
	pulumi.RegisterOutputType(ClusterHealthPolicyResponseOutput{})
	pulumi.RegisterOutputType(ClusterHealthPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterUpgradeDeltaHealthPolicyOutput{})
	pulumi.RegisterOutputType(ClusterUpgradeDeltaHealthPolicyPtrOutput{})
	pulumi.RegisterOutputType(ClusterUpgradeDeltaHealthPolicyResponseOutput{})
	pulumi.RegisterOutputType(ClusterUpgradeDeltaHealthPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterUpgradePolicyOutput{})
	pulumi.RegisterOutputType(ClusterUpgradePolicyPtrOutput{})
	pulumi.RegisterOutputType(ClusterUpgradePolicyResponseOutput{})
	pulumi.RegisterOutputType(ClusterUpgradePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterVersionDetailsResponseOutput{})
	pulumi.RegisterOutputType(ClusterVersionDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticsStorageAccountConfigOutput{})
	pulumi.RegisterOutputType(DiagnosticsStorageAccountConfigPtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsStorageAccountConfigResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticsStorageAccountConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionPtrOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionResponseOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(NodeTypeDescriptionOutput{})
	pulumi.RegisterOutputType(NodeTypeDescriptionArrayOutput{})
	pulumi.RegisterOutputType(NodeTypeDescriptionResponseOutput{})
	pulumi.RegisterOutputType(NodeTypeDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionArrayOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionArrayOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionResponseArrayOutput{})
}

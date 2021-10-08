


package servicelinker

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





type SecretAuthInfoInput interface {
	pulumi.Input

	ToSecretAuthInfoOutput() SecretAuthInfoOutput
	ToSecretAuthInfoOutputWithContext(context.Context) SecretAuthInfoOutput
}

type SecretAuthInfoArgs struct {
	AuthType pulumi.StringInput    `pulumi:"authType"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Secret   pulumi.StringPtrInput `pulumi:"secret"`
}

func (SecretAuthInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretAuthInfo)(nil)).Elem()
}

func (i SecretAuthInfoArgs) ToSecretAuthInfoOutput() SecretAuthInfoOutput {
	return i.ToSecretAuthInfoOutputWithContext(context.Background())
}

func (i SecretAuthInfoArgs) ToSecretAuthInfoOutputWithContext(ctx context.Context) SecretAuthInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretAuthInfoOutput)
}

type SecretAuthInfoOutput struct{ *pulumi.OutputState }

func (SecretAuthInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretAuthInfo)(nil)).Elem()
}

func (o SecretAuthInfoOutput) ToSecretAuthInfoOutput() SecretAuthInfoOutput {
	return o
}

func (o SecretAuthInfoOutput) ToSecretAuthInfoOutputWithContext(ctx context.Context) SecretAuthInfoOutput {
	return o
}

func (o SecretAuthInfoOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v SecretAuthInfo) string { return v.AuthType }).(pulumi.StringOutput)
}

func (o SecretAuthInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretAuthInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SecretAuthInfoOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretAuthInfo) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type SecretAuthInfoResponse struct {
	AuthType string  `pulumi:"authType"`
	Name     *string `pulumi:"name"`
	Secret   *string `pulumi:"secret"`
}





type SecretAuthInfoResponseInput interface {
	pulumi.Input

	ToSecretAuthInfoResponseOutput() SecretAuthInfoResponseOutput
	ToSecretAuthInfoResponseOutputWithContext(context.Context) SecretAuthInfoResponseOutput
}

type SecretAuthInfoResponseArgs struct {
	AuthType pulumi.StringInput    `pulumi:"authType"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Secret   pulumi.StringPtrInput `pulumi:"secret"`
}

func (SecretAuthInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretAuthInfoResponse)(nil)).Elem()
}

func (i SecretAuthInfoResponseArgs) ToSecretAuthInfoResponseOutput() SecretAuthInfoResponseOutput {
	return i.ToSecretAuthInfoResponseOutputWithContext(context.Background())
}

func (i SecretAuthInfoResponseArgs) ToSecretAuthInfoResponseOutputWithContext(ctx context.Context) SecretAuthInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretAuthInfoResponseOutput)
}

type SecretAuthInfoResponseOutput struct{ *pulumi.OutputState }

func (SecretAuthInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretAuthInfoResponse)(nil)).Elem()
}

func (o SecretAuthInfoResponseOutput) ToSecretAuthInfoResponseOutput() SecretAuthInfoResponseOutput {
	return o
}

func (o SecretAuthInfoResponseOutput) ToSecretAuthInfoResponseOutputWithContext(ctx context.Context) SecretAuthInfoResponseOutput {
	return o
}

func (o SecretAuthInfoResponseOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v SecretAuthInfoResponse) string { return v.AuthType }).(pulumi.StringOutput)
}

func (o SecretAuthInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretAuthInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SecretAuthInfoResponseOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretAuthInfoResponse) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ServicePrincipalCertificateAuthInfo struct {
	AuthType    string `pulumi:"authType"`
	Certificate string `pulumi:"certificate"`
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type ServicePrincipalCertificateAuthInfoInput interface {
	pulumi.Input

	ToServicePrincipalCertificateAuthInfoOutput() ServicePrincipalCertificateAuthInfoOutput
	ToServicePrincipalCertificateAuthInfoOutputWithContext(context.Context) ServicePrincipalCertificateAuthInfoOutput
}

type ServicePrincipalCertificateAuthInfoArgs struct {
	AuthType    pulumi.StringInput `pulumi:"authType"`
	Certificate pulumi.StringInput `pulumi:"certificate"`
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (ServicePrincipalCertificateAuthInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalCertificateAuthInfo)(nil)).Elem()
}

func (i ServicePrincipalCertificateAuthInfoArgs) ToServicePrincipalCertificateAuthInfoOutput() ServicePrincipalCertificateAuthInfoOutput {
	return i.ToServicePrincipalCertificateAuthInfoOutputWithContext(context.Background())
}

func (i ServicePrincipalCertificateAuthInfoArgs) ToServicePrincipalCertificateAuthInfoOutputWithContext(ctx context.Context) ServicePrincipalCertificateAuthInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalCertificateAuthInfoOutput)
}

type ServicePrincipalCertificateAuthInfoOutput struct{ *pulumi.OutputState }

func (ServicePrincipalCertificateAuthInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalCertificateAuthInfo)(nil)).Elem()
}

func (o ServicePrincipalCertificateAuthInfoOutput) ToServicePrincipalCertificateAuthInfoOutput() ServicePrincipalCertificateAuthInfoOutput {
	return o
}

func (o ServicePrincipalCertificateAuthInfoOutput) ToServicePrincipalCertificateAuthInfoOutputWithContext(ctx context.Context) ServicePrincipalCertificateAuthInfoOutput {
	return o
}

func (o ServicePrincipalCertificateAuthInfoOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalCertificateAuthInfo) string { return v.AuthType }).(pulumi.StringOutput)
}

func (o ServicePrincipalCertificateAuthInfoOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalCertificateAuthInfo) string { return v.Certificate }).(pulumi.StringOutput)
}

func (o ServicePrincipalCertificateAuthInfoOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalCertificateAuthInfo) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ServicePrincipalCertificateAuthInfoOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalCertificateAuthInfo) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ServicePrincipalCertificateAuthInfoResponse struct {
	AuthType    string `pulumi:"authType"`
	Certificate string `pulumi:"certificate"`
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type ServicePrincipalCertificateAuthInfoResponseInput interface {
	pulumi.Input

	ToServicePrincipalCertificateAuthInfoResponseOutput() ServicePrincipalCertificateAuthInfoResponseOutput
	ToServicePrincipalCertificateAuthInfoResponseOutputWithContext(context.Context) ServicePrincipalCertificateAuthInfoResponseOutput
}

type ServicePrincipalCertificateAuthInfoResponseArgs struct {
	AuthType    pulumi.StringInput `pulumi:"authType"`
	Certificate pulumi.StringInput `pulumi:"certificate"`
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (ServicePrincipalCertificateAuthInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalCertificateAuthInfoResponse)(nil)).Elem()
}

func (i ServicePrincipalCertificateAuthInfoResponseArgs) ToServicePrincipalCertificateAuthInfoResponseOutput() ServicePrincipalCertificateAuthInfoResponseOutput {
	return i.ToServicePrincipalCertificateAuthInfoResponseOutputWithContext(context.Background())
}

func (i ServicePrincipalCertificateAuthInfoResponseArgs) ToServicePrincipalCertificateAuthInfoResponseOutputWithContext(ctx context.Context) ServicePrincipalCertificateAuthInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalCertificateAuthInfoResponseOutput)
}

type ServicePrincipalCertificateAuthInfoResponseOutput struct{ *pulumi.OutputState }

func (ServicePrincipalCertificateAuthInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalCertificateAuthInfoResponse)(nil)).Elem()
}

func (o ServicePrincipalCertificateAuthInfoResponseOutput) ToServicePrincipalCertificateAuthInfoResponseOutput() ServicePrincipalCertificateAuthInfoResponseOutput {
	return o
}

func (o ServicePrincipalCertificateAuthInfoResponseOutput) ToServicePrincipalCertificateAuthInfoResponseOutputWithContext(ctx context.Context) ServicePrincipalCertificateAuthInfoResponseOutput {
	return o
}

func (o ServicePrincipalCertificateAuthInfoResponseOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalCertificateAuthInfoResponse) string { return v.AuthType }).(pulumi.StringOutput)
}

func (o ServicePrincipalCertificateAuthInfoResponseOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalCertificateAuthInfoResponse) string { return v.Certificate }).(pulumi.StringOutput)
}

func (o ServicePrincipalCertificateAuthInfoResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalCertificateAuthInfoResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ServicePrincipalCertificateAuthInfoResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalCertificateAuthInfoResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ServicePrincipalSecretAuthInfo struct {
	AuthType    string `pulumi:"authType"`
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
	Secret      string `pulumi:"secret"`
}





type ServicePrincipalSecretAuthInfoInput interface {
	pulumi.Input

	ToServicePrincipalSecretAuthInfoOutput() ServicePrincipalSecretAuthInfoOutput
	ToServicePrincipalSecretAuthInfoOutputWithContext(context.Context) ServicePrincipalSecretAuthInfoOutput
}

type ServicePrincipalSecretAuthInfoArgs struct {
	AuthType    pulumi.StringInput `pulumi:"authType"`
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	Secret      pulumi.StringInput `pulumi:"secret"`
}

func (ServicePrincipalSecretAuthInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalSecretAuthInfo)(nil)).Elem()
}

func (i ServicePrincipalSecretAuthInfoArgs) ToServicePrincipalSecretAuthInfoOutput() ServicePrincipalSecretAuthInfoOutput {
	return i.ToServicePrincipalSecretAuthInfoOutputWithContext(context.Background())
}

func (i ServicePrincipalSecretAuthInfoArgs) ToServicePrincipalSecretAuthInfoOutputWithContext(ctx context.Context) ServicePrincipalSecretAuthInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalSecretAuthInfoOutput)
}

type ServicePrincipalSecretAuthInfoOutput struct{ *pulumi.OutputState }

func (ServicePrincipalSecretAuthInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalSecretAuthInfo)(nil)).Elem()
}

func (o ServicePrincipalSecretAuthInfoOutput) ToServicePrincipalSecretAuthInfoOutput() ServicePrincipalSecretAuthInfoOutput {
	return o
}

func (o ServicePrincipalSecretAuthInfoOutput) ToServicePrincipalSecretAuthInfoOutputWithContext(ctx context.Context) ServicePrincipalSecretAuthInfoOutput {
	return o
}

func (o ServicePrincipalSecretAuthInfoOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalSecretAuthInfo) string { return v.AuthType }).(pulumi.StringOutput)
}

func (o ServicePrincipalSecretAuthInfoOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalSecretAuthInfo) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ServicePrincipalSecretAuthInfoOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalSecretAuthInfo) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ServicePrincipalSecretAuthInfoOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalSecretAuthInfo) string { return v.Secret }).(pulumi.StringOutput)
}

type ServicePrincipalSecretAuthInfoResponse struct {
	AuthType    string `pulumi:"authType"`
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
	Secret      string `pulumi:"secret"`
}





type ServicePrincipalSecretAuthInfoResponseInput interface {
	pulumi.Input

	ToServicePrincipalSecretAuthInfoResponseOutput() ServicePrincipalSecretAuthInfoResponseOutput
	ToServicePrincipalSecretAuthInfoResponseOutputWithContext(context.Context) ServicePrincipalSecretAuthInfoResponseOutput
}

type ServicePrincipalSecretAuthInfoResponseArgs struct {
	AuthType    pulumi.StringInput `pulumi:"authType"`
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	Secret      pulumi.StringInput `pulumi:"secret"`
}

func (ServicePrincipalSecretAuthInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalSecretAuthInfoResponse)(nil)).Elem()
}

func (i ServicePrincipalSecretAuthInfoResponseArgs) ToServicePrincipalSecretAuthInfoResponseOutput() ServicePrincipalSecretAuthInfoResponseOutput {
	return i.ToServicePrincipalSecretAuthInfoResponseOutputWithContext(context.Background())
}

func (i ServicePrincipalSecretAuthInfoResponseArgs) ToServicePrincipalSecretAuthInfoResponseOutputWithContext(ctx context.Context) ServicePrincipalSecretAuthInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalSecretAuthInfoResponseOutput)
}

type ServicePrincipalSecretAuthInfoResponseOutput struct{ *pulumi.OutputState }

func (ServicePrincipalSecretAuthInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalSecretAuthInfoResponse)(nil)).Elem()
}

func (o ServicePrincipalSecretAuthInfoResponseOutput) ToServicePrincipalSecretAuthInfoResponseOutput() ServicePrincipalSecretAuthInfoResponseOutput {
	return o
}

func (o ServicePrincipalSecretAuthInfoResponseOutput) ToServicePrincipalSecretAuthInfoResponseOutputWithContext(ctx context.Context) ServicePrincipalSecretAuthInfoResponseOutput {
	return o
}

func (o ServicePrincipalSecretAuthInfoResponseOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalSecretAuthInfoResponse) string { return v.AuthType }).(pulumi.StringOutput)
}

func (o ServicePrincipalSecretAuthInfoResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalSecretAuthInfoResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ServicePrincipalSecretAuthInfoResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalSecretAuthInfoResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ServicePrincipalSecretAuthInfoResponseOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalSecretAuthInfoResponse) string { return v.Secret }).(pulumi.StringOutput)
}

type SourceConfigurationResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type SourceConfigurationResponseInput interface {
	pulumi.Input

	ToSourceConfigurationResponseOutput() SourceConfigurationResponseOutput
	ToSourceConfigurationResponseOutputWithContext(context.Context) SourceConfigurationResponseOutput
}

type SourceConfigurationResponseArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (SourceConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceConfigurationResponse)(nil)).Elem()
}

func (i SourceConfigurationResponseArgs) ToSourceConfigurationResponseOutput() SourceConfigurationResponseOutput {
	return i.ToSourceConfigurationResponseOutputWithContext(context.Background())
}

func (i SourceConfigurationResponseArgs) ToSourceConfigurationResponseOutputWithContext(ctx context.Context) SourceConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceConfigurationResponseOutput)
}





type SourceConfigurationResponseArrayInput interface {
	pulumi.Input

	ToSourceConfigurationResponseArrayOutput() SourceConfigurationResponseArrayOutput
	ToSourceConfigurationResponseArrayOutputWithContext(context.Context) SourceConfigurationResponseArrayOutput
}

type SourceConfigurationResponseArray []SourceConfigurationResponseInput

func (SourceConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceConfigurationResponse)(nil)).Elem()
}

func (i SourceConfigurationResponseArray) ToSourceConfigurationResponseArrayOutput() SourceConfigurationResponseArrayOutput {
	return i.ToSourceConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i SourceConfigurationResponseArray) ToSourceConfigurationResponseArrayOutputWithContext(ctx context.Context) SourceConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceConfigurationResponseArrayOutput)
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





type SystemAssignedIdentityAuthInfoInput interface {
	pulumi.Input

	ToSystemAssignedIdentityAuthInfoOutput() SystemAssignedIdentityAuthInfoOutput
	ToSystemAssignedIdentityAuthInfoOutputWithContext(context.Context) SystemAssignedIdentityAuthInfoOutput
}

type SystemAssignedIdentityAuthInfoArgs struct {
	AuthType pulumi.StringInput `pulumi:"authType"`
}

func (SystemAssignedIdentityAuthInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedIdentityAuthInfo)(nil)).Elem()
}

func (i SystemAssignedIdentityAuthInfoArgs) ToSystemAssignedIdentityAuthInfoOutput() SystemAssignedIdentityAuthInfoOutput {
	return i.ToSystemAssignedIdentityAuthInfoOutputWithContext(context.Background())
}

func (i SystemAssignedIdentityAuthInfoArgs) ToSystemAssignedIdentityAuthInfoOutputWithContext(ctx context.Context) SystemAssignedIdentityAuthInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedIdentityAuthInfoOutput)
}

type SystemAssignedIdentityAuthInfoOutput struct{ *pulumi.OutputState }

func (SystemAssignedIdentityAuthInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedIdentityAuthInfo)(nil)).Elem()
}

func (o SystemAssignedIdentityAuthInfoOutput) ToSystemAssignedIdentityAuthInfoOutput() SystemAssignedIdentityAuthInfoOutput {
	return o
}

func (o SystemAssignedIdentityAuthInfoOutput) ToSystemAssignedIdentityAuthInfoOutputWithContext(ctx context.Context) SystemAssignedIdentityAuthInfoOutput {
	return o
}

func (o SystemAssignedIdentityAuthInfoOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedIdentityAuthInfo) string { return v.AuthType }).(pulumi.StringOutput)
}

type SystemAssignedIdentityAuthInfoResponse struct {
	AuthType string `pulumi:"authType"`
}





type SystemAssignedIdentityAuthInfoResponseInput interface {
	pulumi.Input

	ToSystemAssignedIdentityAuthInfoResponseOutput() SystemAssignedIdentityAuthInfoResponseOutput
	ToSystemAssignedIdentityAuthInfoResponseOutputWithContext(context.Context) SystemAssignedIdentityAuthInfoResponseOutput
}

type SystemAssignedIdentityAuthInfoResponseArgs struct {
	AuthType pulumi.StringInput `pulumi:"authType"`
}

func (SystemAssignedIdentityAuthInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedIdentityAuthInfoResponse)(nil)).Elem()
}

func (i SystemAssignedIdentityAuthInfoResponseArgs) ToSystemAssignedIdentityAuthInfoResponseOutput() SystemAssignedIdentityAuthInfoResponseOutput {
	return i.ToSystemAssignedIdentityAuthInfoResponseOutputWithContext(context.Background())
}

func (i SystemAssignedIdentityAuthInfoResponseArgs) ToSystemAssignedIdentityAuthInfoResponseOutputWithContext(ctx context.Context) SystemAssignedIdentityAuthInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedIdentityAuthInfoResponseOutput)
}

type SystemAssignedIdentityAuthInfoResponseOutput struct{ *pulumi.OutputState }

func (SystemAssignedIdentityAuthInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedIdentityAuthInfoResponse)(nil)).Elem()
}

func (o SystemAssignedIdentityAuthInfoResponseOutput) ToSystemAssignedIdentityAuthInfoResponseOutput() SystemAssignedIdentityAuthInfoResponseOutput {
	return o
}

func (o SystemAssignedIdentityAuthInfoResponseOutput) ToSystemAssignedIdentityAuthInfoResponseOutputWithContext(ctx context.Context) SystemAssignedIdentityAuthInfoResponseOutput {
	return o
}

func (o SystemAssignedIdentityAuthInfoResponseOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedIdentityAuthInfoResponse) string { return v.AuthType }).(pulumi.StringOutput)
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

type UserAssignedIdentityAuthInfo struct {
	AuthType       string `pulumi:"authType"`
	ClientId       string `pulumi:"clientId"`
	SubscriptionId string `pulumi:"subscriptionId"`
}





type UserAssignedIdentityAuthInfoInput interface {
	pulumi.Input

	ToUserAssignedIdentityAuthInfoOutput() UserAssignedIdentityAuthInfoOutput
	ToUserAssignedIdentityAuthInfoOutputWithContext(context.Context) UserAssignedIdentityAuthInfoOutput
}

type UserAssignedIdentityAuthInfoArgs struct {
	AuthType       pulumi.StringInput `pulumi:"authType"`
	ClientId       pulumi.StringInput `pulumi:"clientId"`
	SubscriptionId pulumi.StringInput `pulumi:"subscriptionId"`
}

func (UserAssignedIdentityAuthInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityAuthInfo)(nil)).Elem()
}

func (i UserAssignedIdentityAuthInfoArgs) ToUserAssignedIdentityAuthInfoOutput() UserAssignedIdentityAuthInfoOutput {
	return i.ToUserAssignedIdentityAuthInfoOutputWithContext(context.Background())
}

func (i UserAssignedIdentityAuthInfoArgs) ToUserAssignedIdentityAuthInfoOutputWithContext(ctx context.Context) UserAssignedIdentityAuthInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityAuthInfoOutput)
}

type UserAssignedIdentityAuthInfoOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityAuthInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityAuthInfo)(nil)).Elem()
}

func (o UserAssignedIdentityAuthInfoOutput) ToUserAssignedIdentityAuthInfoOutput() UserAssignedIdentityAuthInfoOutput {
	return o
}

func (o UserAssignedIdentityAuthInfoOutput) ToUserAssignedIdentityAuthInfoOutputWithContext(ctx context.Context) UserAssignedIdentityAuthInfoOutput {
	return o
}

func (o UserAssignedIdentityAuthInfoOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityAuthInfo) string { return v.AuthType }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityAuthInfoOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityAuthInfo) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityAuthInfoOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityAuthInfo) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

type UserAssignedIdentityAuthInfoResponse struct {
	AuthType       string `pulumi:"authType"`
	ClientId       string `pulumi:"clientId"`
	SubscriptionId string `pulumi:"subscriptionId"`
}





type UserAssignedIdentityAuthInfoResponseInput interface {
	pulumi.Input

	ToUserAssignedIdentityAuthInfoResponseOutput() UserAssignedIdentityAuthInfoResponseOutput
	ToUserAssignedIdentityAuthInfoResponseOutputWithContext(context.Context) UserAssignedIdentityAuthInfoResponseOutput
}

type UserAssignedIdentityAuthInfoResponseArgs struct {
	AuthType       pulumi.StringInput `pulumi:"authType"`
	ClientId       pulumi.StringInput `pulumi:"clientId"`
	SubscriptionId pulumi.StringInput `pulumi:"subscriptionId"`
}

func (UserAssignedIdentityAuthInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityAuthInfoResponse)(nil)).Elem()
}

func (i UserAssignedIdentityAuthInfoResponseArgs) ToUserAssignedIdentityAuthInfoResponseOutput() UserAssignedIdentityAuthInfoResponseOutput {
	return i.ToUserAssignedIdentityAuthInfoResponseOutputWithContext(context.Background())
}

func (i UserAssignedIdentityAuthInfoResponseArgs) ToUserAssignedIdentityAuthInfoResponseOutputWithContext(ctx context.Context) UserAssignedIdentityAuthInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityAuthInfoResponseOutput)
}

type UserAssignedIdentityAuthInfoResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityAuthInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityAuthInfoResponse)(nil)).Elem()
}

func (o UserAssignedIdentityAuthInfoResponseOutput) ToUserAssignedIdentityAuthInfoResponseOutput() UserAssignedIdentityAuthInfoResponseOutput {
	return o
}

func (o UserAssignedIdentityAuthInfoResponseOutput) ToUserAssignedIdentityAuthInfoResponseOutputWithContext(ctx context.Context) UserAssignedIdentityAuthInfoResponseOutput {
	return o
}

func (o UserAssignedIdentityAuthInfoResponseOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityAuthInfoResponse) string { return v.AuthType }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityAuthInfoResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityAuthInfoResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityAuthInfoResponseOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityAuthInfoResponse) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretAuthInfoInput)(nil)).Elem(), SecretAuthInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretAuthInfoResponseInput)(nil)).Elem(), SecretAuthInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalCertificateAuthInfoInput)(nil)).Elem(), ServicePrincipalCertificateAuthInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalCertificateAuthInfoResponseInput)(nil)).Elem(), ServicePrincipalCertificateAuthInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalSecretAuthInfoInput)(nil)).Elem(), ServicePrincipalSecretAuthInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalSecretAuthInfoResponseInput)(nil)).Elem(), ServicePrincipalSecretAuthInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceConfigurationResponseInput)(nil)).Elem(), SourceConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceConfigurationResponseArrayInput)(nil)).Elem(), SourceConfigurationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAssignedIdentityAuthInfoInput)(nil)).Elem(), SystemAssignedIdentityAuthInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAssignedIdentityAuthInfoResponseInput)(nil)).Elem(), SystemAssignedIdentityAuthInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAssignedIdentityAuthInfoInput)(nil)).Elem(), UserAssignedIdentityAuthInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAssignedIdentityAuthInfoResponseInput)(nil)).Elem(), UserAssignedIdentityAuthInfoResponseArgs{})
	pulumi.RegisterOutputType(SecretAuthInfoOutput{})
	pulumi.RegisterOutputType(SecretAuthInfoResponseOutput{})
	pulumi.RegisterOutputType(ServicePrincipalCertificateAuthInfoOutput{})
	pulumi.RegisterOutputType(ServicePrincipalCertificateAuthInfoResponseOutput{})
	pulumi.RegisterOutputType(ServicePrincipalSecretAuthInfoOutput{})
	pulumi.RegisterOutputType(ServicePrincipalSecretAuthInfoResponseOutput{})
	pulumi.RegisterOutputType(SourceConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SourceConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemAssignedIdentityAuthInfoOutput{})
	pulumi.RegisterOutputType(SystemAssignedIdentityAuthInfoResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityAuthInfoOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityAuthInfoResponseOutput{})
}

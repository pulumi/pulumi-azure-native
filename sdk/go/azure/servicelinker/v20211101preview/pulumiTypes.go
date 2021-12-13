


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

func init() {
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

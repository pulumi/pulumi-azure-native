


package v20220908preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AADBasedSecurityPrincipal struct {
	LedgerRoleName *string `pulumi:"ledgerRoleName"`
	PrincipalId    *string `pulumi:"principalId"`
	TenantId       *string `pulumi:"tenantId"`
}





type AADBasedSecurityPrincipalInput interface {
	pulumi.Input

	ToAADBasedSecurityPrincipalOutput() AADBasedSecurityPrincipalOutput
	ToAADBasedSecurityPrincipalOutputWithContext(context.Context) AADBasedSecurityPrincipalOutput
}

type AADBasedSecurityPrincipalArgs struct {
	LedgerRoleName pulumi.StringPtrInput `pulumi:"ledgerRoleName"`
	PrincipalId    pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId       pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (AADBasedSecurityPrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AADBasedSecurityPrincipal)(nil)).Elem()
}

func (i AADBasedSecurityPrincipalArgs) ToAADBasedSecurityPrincipalOutput() AADBasedSecurityPrincipalOutput {
	return i.ToAADBasedSecurityPrincipalOutputWithContext(context.Background())
}

func (i AADBasedSecurityPrincipalArgs) ToAADBasedSecurityPrincipalOutputWithContext(ctx context.Context) AADBasedSecurityPrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADBasedSecurityPrincipalOutput)
}





type AADBasedSecurityPrincipalArrayInput interface {
	pulumi.Input

	ToAADBasedSecurityPrincipalArrayOutput() AADBasedSecurityPrincipalArrayOutput
	ToAADBasedSecurityPrincipalArrayOutputWithContext(context.Context) AADBasedSecurityPrincipalArrayOutput
}

type AADBasedSecurityPrincipalArray []AADBasedSecurityPrincipalInput

func (AADBasedSecurityPrincipalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AADBasedSecurityPrincipal)(nil)).Elem()
}

func (i AADBasedSecurityPrincipalArray) ToAADBasedSecurityPrincipalArrayOutput() AADBasedSecurityPrincipalArrayOutput {
	return i.ToAADBasedSecurityPrincipalArrayOutputWithContext(context.Background())
}

func (i AADBasedSecurityPrincipalArray) ToAADBasedSecurityPrincipalArrayOutputWithContext(ctx context.Context) AADBasedSecurityPrincipalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADBasedSecurityPrincipalArrayOutput)
}

type AADBasedSecurityPrincipalOutput struct{ *pulumi.OutputState }

func (AADBasedSecurityPrincipalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AADBasedSecurityPrincipal)(nil)).Elem()
}

func (o AADBasedSecurityPrincipalOutput) ToAADBasedSecurityPrincipalOutput() AADBasedSecurityPrincipalOutput {
	return o
}

func (o AADBasedSecurityPrincipalOutput) ToAADBasedSecurityPrincipalOutputWithContext(ctx context.Context) AADBasedSecurityPrincipalOutput {
	return o
}

func (o AADBasedSecurityPrincipalOutput) LedgerRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADBasedSecurityPrincipal) *string { return v.LedgerRoleName }).(pulumi.StringPtrOutput)
}

func (o AADBasedSecurityPrincipalOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADBasedSecurityPrincipal) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o AADBasedSecurityPrincipalOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADBasedSecurityPrincipal) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AADBasedSecurityPrincipalArrayOutput struct{ *pulumi.OutputState }

func (AADBasedSecurityPrincipalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AADBasedSecurityPrincipal)(nil)).Elem()
}

func (o AADBasedSecurityPrincipalArrayOutput) ToAADBasedSecurityPrincipalArrayOutput() AADBasedSecurityPrincipalArrayOutput {
	return o
}

func (o AADBasedSecurityPrincipalArrayOutput) ToAADBasedSecurityPrincipalArrayOutputWithContext(ctx context.Context) AADBasedSecurityPrincipalArrayOutput {
	return o
}

func (o AADBasedSecurityPrincipalArrayOutput) Index(i pulumi.IntInput) AADBasedSecurityPrincipalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AADBasedSecurityPrincipal {
		return vs[0].([]AADBasedSecurityPrincipal)[vs[1].(int)]
	}).(AADBasedSecurityPrincipalOutput)
}

type AADBasedSecurityPrincipalResponse struct {
	LedgerRoleName *string `pulumi:"ledgerRoleName"`
	PrincipalId    *string `pulumi:"principalId"`
	TenantId       *string `pulumi:"tenantId"`
}

type AADBasedSecurityPrincipalResponseOutput struct{ *pulumi.OutputState }

func (AADBasedSecurityPrincipalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AADBasedSecurityPrincipalResponse)(nil)).Elem()
}

func (o AADBasedSecurityPrincipalResponseOutput) ToAADBasedSecurityPrincipalResponseOutput() AADBasedSecurityPrincipalResponseOutput {
	return o
}

func (o AADBasedSecurityPrincipalResponseOutput) ToAADBasedSecurityPrincipalResponseOutputWithContext(ctx context.Context) AADBasedSecurityPrincipalResponseOutput {
	return o
}

func (o AADBasedSecurityPrincipalResponseOutput) LedgerRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADBasedSecurityPrincipalResponse) *string { return v.LedgerRoleName }).(pulumi.StringPtrOutput)
}

func (o AADBasedSecurityPrincipalResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADBasedSecurityPrincipalResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o AADBasedSecurityPrincipalResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADBasedSecurityPrincipalResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AADBasedSecurityPrincipalResponseArrayOutput struct{ *pulumi.OutputState }

func (AADBasedSecurityPrincipalResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AADBasedSecurityPrincipalResponse)(nil)).Elem()
}

func (o AADBasedSecurityPrincipalResponseArrayOutput) ToAADBasedSecurityPrincipalResponseArrayOutput() AADBasedSecurityPrincipalResponseArrayOutput {
	return o
}

func (o AADBasedSecurityPrincipalResponseArrayOutput) ToAADBasedSecurityPrincipalResponseArrayOutputWithContext(ctx context.Context) AADBasedSecurityPrincipalResponseArrayOutput {
	return o
}

func (o AADBasedSecurityPrincipalResponseArrayOutput) Index(i pulumi.IntInput) AADBasedSecurityPrincipalResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AADBasedSecurityPrincipalResponse {
		return vs[0].([]AADBasedSecurityPrincipalResponse)[vs[1].(int)]
	}).(AADBasedSecurityPrincipalResponseOutput)
}

type CertBasedSecurityPrincipal struct {
	Cert           *string `pulumi:"cert"`
	LedgerRoleName *string `pulumi:"ledgerRoleName"`
}





type CertBasedSecurityPrincipalInput interface {
	pulumi.Input

	ToCertBasedSecurityPrincipalOutput() CertBasedSecurityPrincipalOutput
	ToCertBasedSecurityPrincipalOutputWithContext(context.Context) CertBasedSecurityPrincipalOutput
}

type CertBasedSecurityPrincipalArgs struct {
	Cert           pulumi.StringPtrInput `pulumi:"cert"`
	LedgerRoleName pulumi.StringPtrInput `pulumi:"ledgerRoleName"`
}

func (CertBasedSecurityPrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertBasedSecurityPrincipal)(nil)).Elem()
}

func (i CertBasedSecurityPrincipalArgs) ToCertBasedSecurityPrincipalOutput() CertBasedSecurityPrincipalOutput {
	return i.ToCertBasedSecurityPrincipalOutputWithContext(context.Background())
}

func (i CertBasedSecurityPrincipalArgs) ToCertBasedSecurityPrincipalOutputWithContext(ctx context.Context) CertBasedSecurityPrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertBasedSecurityPrincipalOutput)
}





type CertBasedSecurityPrincipalArrayInput interface {
	pulumi.Input

	ToCertBasedSecurityPrincipalArrayOutput() CertBasedSecurityPrincipalArrayOutput
	ToCertBasedSecurityPrincipalArrayOutputWithContext(context.Context) CertBasedSecurityPrincipalArrayOutput
}

type CertBasedSecurityPrincipalArray []CertBasedSecurityPrincipalInput

func (CertBasedSecurityPrincipalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertBasedSecurityPrincipal)(nil)).Elem()
}

func (i CertBasedSecurityPrincipalArray) ToCertBasedSecurityPrincipalArrayOutput() CertBasedSecurityPrincipalArrayOutput {
	return i.ToCertBasedSecurityPrincipalArrayOutputWithContext(context.Background())
}

func (i CertBasedSecurityPrincipalArray) ToCertBasedSecurityPrincipalArrayOutputWithContext(ctx context.Context) CertBasedSecurityPrincipalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertBasedSecurityPrincipalArrayOutput)
}

type CertBasedSecurityPrincipalOutput struct{ *pulumi.OutputState }

func (CertBasedSecurityPrincipalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertBasedSecurityPrincipal)(nil)).Elem()
}

func (o CertBasedSecurityPrincipalOutput) ToCertBasedSecurityPrincipalOutput() CertBasedSecurityPrincipalOutput {
	return o
}

func (o CertBasedSecurityPrincipalOutput) ToCertBasedSecurityPrincipalOutputWithContext(ctx context.Context) CertBasedSecurityPrincipalOutput {
	return o
}

func (o CertBasedSecurityPrincipalOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertBasedSecurityPrincipal) *string { return v.Cert }).(pulumi.StringPtrOutput)
}

func (o CertBasedSecurityPrincipalOutput) LedgerRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertBasedSecurityPrincipal) *string { return v.LedgerRoleName }).(pulumi.StringPtrOutput)
}

type CertBasedSecurityPrincipalArrayOutput struct{ *pulumi.OutputState }

func (CertBasedSecurityPrincipalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertBasedSecurityPrincipal)(nil)).Elem()
}

func (o CertBasedSecurityPrincipalArrayOutput) ToCertBasedSecurityPrincipalArrayOutput() CertBasedSecurityPrincipalArrayOutput {
	return o
}

func (o CertBasedSecurityPrincipalArrayOutput) ToCertBasedSecurityPrincipalArrayOutputWithContext(ctx context.Context) CertBasedSecurityPrincipalArrayOutput {
	return o
}

func (o CertBasedSecurityPrincipalArrayOutput) Index(i pulumi.IntInput) CertBasedSecurityPrincipalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertBasedSecurityPrincipal {
		return vs[0].([]CertBasedSecurityPrincipal)[vs[1].(int)]
	}).(CertBasedSecurityPrincipalOutput)
}

type CertBasedSecurityPrincipalResponse struct {
	Cert           *string `pulumi:"cert"`
	LedgerRoleName *string `pulumi:"ledgerRoleName"`
}

type CertBasedSecurityPrincipalResponseOutput struct{ *pulumi.OutputState }

func (CertBasedSecurityPrincipalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertBasedSecurityPrincipalResponse)(nil)).Elem()
}

func (o CertBasedSecurityPrincipalResponseOutput) ToCertBasedSecurityPrincipalResponseOutput() CertBasedSecurityPrincipalResponseOutput {
	return o
}

func (o CertBasedSecurityPrincipalResponseOutput) ToCertBasedSecurityPrincipalResponseOutputWithContext(ctx context.Context) CertBasedSecurityPrincipalResponseOutput {
	return o
}

func (o CertBasedSecurityPrincipalResponseOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertBasedSecurityPrincipalResponse) *string { return v.Cert }).(pulumi.StringPtrOutput)
}

func (o CertBasedSecurityPrincipalResponseOutput) LedgerRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertBasedSecurityPrincipalResponse) *string { return v.LedgerRoleName }).(pulumi.StringPtrOutput)
}

type CertBasedSecurityPrincipalResponseArrayOutput struct{ *pulumi.OutputState }

func (CertBasedSecurityPrincipalResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertBasedSecurityPrincipalResponse)(nil)).Elem()
}

func (o CertBasedSecurityPrincipalResponseArrayOutput) ToCertBasedSecurityPrincipalResponseArrayOutput() CertBasedSecurityPrincipalResponseArrayOutput {
	return o
}

func (o CertBasedSecurityPrincipalResponseArrayOutput) ToCertBasedSecurityPrincipalResponseArrayOutputWithContext(ctx context.Context) CertBasedSecurityPrincipalResponseArrayOutput {
	return o
}

func (o CertBasedSecurityPrincipalResponseArrayOutput) Index(i pulumi.IntInput) CertBasedSecurityPrincipalResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertBasedSecurityPrincipalResponse {
		return vs[0].([]CertBasedSecurityPrincipalResponse)[vs[1].(int)]
	}).(CertBasedSecurityPrincipalResponseOutput)
}

type CertificateTags struct {
	Tags map[string]string `pulumi:"tags"`
}





type CertificateTagsInput interface {
	pulumi.Input

	ToCertificateTagsOutput() CertificateTagsOutput
	ToCertificateTagsOutputWithContext(context.Context) CertificateTagsOutput
}

type CertificateTagsArgs struct {
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (CertificateTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateTags)(nil)).Elem()
}

func (i CertificateTagsArgs) ToCertificateTagsOutput() CertificateTagsOutput {
	return i.ToCertificateTagsOutputWithContext(context.Background())
}

func (i CertificateTagsArgs) ToCertificateTagsOutputWithContext(ctx context.Context) CertificateTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateTagsOutput)
}





type CertificateTagsArrayInput interface {
	pulumi.Input

	ToCertificateTagsArrayOutput() CertificateTagsArrayOutput
	ToCertificateTagsArrayOutputWithContext(context.Context) CertificateTagsArrayOutput
}

type CertificateTagsArray []CertificateTagsInput

func (CertificateTagsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateTags)(nil)).Elem()
}

func (i CertificateTagsArray) ToCertificateTagsArrayOutput() CertificateTagsArrayOutput {
	return i.ToCertificateTagsArrayOutputWithContext(context.Background())
}

func (i CertificateTagsArray) ToCertificateTagsArrayOutputWithContext(ctx context.Context) CertificateTagsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateTagsArrayOutput)
}

type CertificateTagsOutput struct{ *pulumi.OutputState }

func (CertificateTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateTags)(nil)).Elem()
}

func (o CertificateTagsOutput) ToCertificateTagsOutput() CertificateTagsOutput {
	return o
}

func (o CertificateTagsOutput) ToCertificateTagsOutputWithContext(ctx context.Context) CertificateTagsOutput {
	return o
}

func (o CertificateTagsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v CertificateTags) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type CertificateTagsArrayOutput struct{ *pulumi.OutputState }

func (CertificateTagsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateTags)(nil)).Elem()
}

func (o CertificateTagsArrayOutput) ToCertificateTagsArrayOutput() CertificateTagsArrayOutput {
	return o
}

func (o CertificateTagsArrayOutput) ToCertificateTagsArrayOutputWithContext(ctx context.Context) CertificateTagsArrayOutput {
	return o
}

func (o CertificateTagsArrayOutput) Index(i pulumi.IntInput) CertificateTagsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateTags {
		return vs[0].([]CertificateTags)[vs[1].(int)]
	}).(CertificateTagsOutput)
}

type CertificateTagsResponse struct {
	Tags map[string]string `pulumi:"tags"`
}

type CertificateTagsResponseOutput struct{ *pulumi.OutputState }

func (CertificateTagsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateTagsResponse)(nil)).Elem()
}

func (o CertificateTagsResponseOutput) ToCertificateTagsResponseOutput() CertificateTagsResponseOutput {
	return o
}

func (o CertificateTagsResponseOutput) ToCertificateTagsResponseOutputWithContext(ctx context.Context) CertificateTagsResponseOutput {
	return o
}

func (o CertificateTagsResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v CertificateTagsResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type CertificateTagsResponseArrayOutput struct{ *pulumi.OutputState }

func (CertificateTagsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateTagsResponse)(nil)).Elem()
}

func (o CertificateTagsResponseArrayOutput) ToCertificateTagsResponseArrayOutput() CertificateTagsResponseArrayOutput {
	return o
}

func (o CertificateTagsResponseArrayOutput) ToCertificateTagsResponseArrayOutputWithContext(ctx context.Context) CertificateTagsResponseArrayOutput {
	return o
}

func (o CertificateTagsResponseArrayOutput) Index(i pulumi.IntInput) CertificateTagsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateTagsResponse {
		return vs[0].([]CertificateTagsResponse)[vs[1].(int)]
	}).(CertificateTagsResponseOutput)
}

type DeploymentType struct {
	AppSourceUri    *string `pulumi:"appSourceUri"`
	LanguageRuntime *string `pulumi:"languageRuntime"`
}





type DeploymentTypeInput interface {
	pulumi.Input

	ToDeploymentTypeOutput() DeploymentTypeOutput
	ToDeploymentTypeOutputWithContext(context.Context) DeploymentTypeOutput
}

type DeploymentTypeArgs struct {
	AppSourceUri    pulumi.StringPtrInput `pulumi:"appSourceUri"`
	LanguageRuntime pulumi.StringPtrInput `pulumi:"languageRuntime"`
}

func (DeploymentTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentType)(nil)).Elem()
}

func (i DeploymentTypeArgs) ToDeploymentTypeOutput() DeploymentTypeOutput {
	return i.ToDeploymentTypeOutputWithContext(context.Background())
}

func (i DeploymentTypeArgs) ToDeploymentTypeOutputWithContext(ctx context.Context) DeploymentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTypeOutput)
}

func (i DeploymentTypeArgs) ToDeploymentTypePtrOutput() DeploymentTypePtrOutput {
	return i.ToDeploymentTypePtrOutputWithContext(context.Background())
}

func (i DeploymentTypeArgs) ToDeploymentTypePtrOutputWithContext(ctx context.Context) DeploymentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTypeOutput).ToDeploymentTypePtrOutputWithContext(ctx)
}









type DeploymentTypePtrInput interface {
	pulumi.Input

	ToDeploymentTypePtrOutput() DeploymentTypePtrOutput
	ToDeploymentTypePtrOutputWithContext(context.Context) DeploymentTypePtrOutput
}

type deploymentTypePtrType DeploymentTypeArgs

func DeploymentTypePtr(v *DeploymentTypeArgs) DeploymentTypePtrInput {
	return (*deploymentTypePtrType)(v)
}

func (*deploymentTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentType)(nil)).Elem()
}

func (i *deploymentTypePtrType) ToDeploymentTypePtrOutput() DeploymentTypePtrOutput {
	return i.ToDeploymentTypePtrOutputWithContext(context.Background())
}

func (i *deploymentTypePtrType) ToDeploymentTypePtrOutputWithContext(ctx context.Context) DeploymentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTypePtrOutput)
}

type DeploymentTypeOutput struct{ *pulumi.OutputState }

func (DeploymentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentType)(nil)).Elem()
}

func (o DeploymentTypeOutput) ToDeploymentTypeOutput() DeploymentTypeOutput {
	return o
}

func (o DeploymentTypeOutput) ToDeploymentTypeOutputWithContext(ctx context.Context) DeploymentTypeOutput {
	return o
}

func (o DeploymentTypeOutput) ToDeploymentTypePtrOutput() DeploymentTypePtrOutput {
	return o.ToDeploymentTypePtrOutputWithContext(context.Background())
}

func (o DeploymentTypeOutput) ToDeploymentTypePtrOutputWithContext(ctx context.Context) DeploymentTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentType) *DeploymentType {
		return &v
	}).(DeploymentTypePtrOutput)
}

func (o DeploymentTypeOutput) AppSourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentType) *string { return v.AppSourceUri }).(pulumi.StringPtrOutput)
}

func (o DeploymentTypeOutput) LanguageRuntime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentType) *string { return v.LanguageRuntime }).(pulumi.StringPtrOutput)
}

type DeploymentTypePtrOutput struct{ *pulumi.OutputState }

func (DeploymentTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentType)(nil)).Elem()
}

func (o DeploymentTypePtrOutput) ToDeploymentTypePtrOutput() DeploymentTypePtrOutput {
	return o
}

func (o DeploymentTypePtrOutput) ToDeploymentTypePtrOutputWithContext(ctx context.Context) DeploymentTypePtrOutput {
	return o
}

func (o DeploymentTypePtrOutput) Elem() DeploymentTypeOutput {
	return o.ApplyT(func(v *DeploymentType) DeploymentType {
		if v != nil {
			return *v
		}
		var ret DeploymentType
		return ret
	}).(DeploymentTypeOutput)
}

func (o DeploymentTypePtrOutput) AppSourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentType) *string {
		if v == nil {
			return nil
		}
		return v.AppSourceUri
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentTypePtrOutput) LanguageRuntime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentType) *string {
		if v == nil {
			return nil
		}
		return v.LanguageRuntime
	}).(pulumi.StringPtrOutput)
}

type DeploymentTypeResponse struct {
	AppSourceUri    *string `pulumi:"appSourceUri"`
	LanguageRuntime *string `pulumi:"languageRuntime"`
}

type DeploymentTypeResponseOutput struct{ *pulumi.OutputState }

func (DeploymentTypeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentTypeResponse)(nil)).Elem()
}

func (o DeploymentTypeResponseOutput) ToDeploymentTypeResponseOutput() DeploymentTypeResponseOutput {
	return o
}

func (o DeploymentTypeResponseOutput) ToDeploymentTypeResponseOutputWithContext(ctx context.Context) DeploymentTypeResponseOutput {
	return o
}

func (o DeploymentTypeResponseOutput) AppSourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentTypeResponse) *string { return v.AppSourceUri }).(pulumi.StringPtrOutput)
}

func (o DeploymentTypeResponseOutput) LanguageRuntime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentTypeResponse) *string { return v.LanguageRuntime }).(pulumi.StringPtrOutput)
}

type DeploymentTypeResponsePtrOutput struct{ *pulumi.OutputState }

func (DeploymentTypeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentTypeResponse)(nil)).Elem()
}

func (o DeploymentTypeResponsePtrOutput) ToDeploymentTypeResponsePtrOutput() DeploymentTypeResponsePtrOutput {
	return o
}

func (o DeploymentTypeResponsePtrOutput) ToDeploymentTypeResponsePtrOutputWithContext(ctx context.Context) DeploymentTypeResponsePtrOutput {
	return o
}

func (o DeploymentTypeResponsePtrOutput) Elem() DeploymentTypeResponseOutput {
	return o.ApplyT(func(v *DeploymentTypeResponse) DeploymentTypeResponse {
		if v != nil {
			return *v
		}
		var ret DeploymentTypeResponse
		return ret
	}).(DeploymentTypeResponseOutput)
}

func (o DeploymentTypeResponsePtrOutput) AppSourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentTypeResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppSourceUri
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentTypeResponsePtrOutput) LanguageRuntime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentTypeResponse) *string {
		if v == nil {
			return nil
		}
		return v.LanguageRuntime
	}).(pulumi.StringPtrOutput)
}

type LedgerProperties struct {
	AadBasedSecurityPrincipals  []AADBasedSecurityPrincipal  `pulumi:"aadBasedSecurityPrincipals"`
	CertBasedSecurityPrincipals []CertBasedSecurityPrincipal `pulumi:"certBasedSecurityPrincipals"`
	LedgerType                  *string                      `pulumi:"ledgerType"`
}





type LedgerPropertiesInput interface {
	pulumi.Input

	ToLedgerPropertiesOutput() LedgerPropertiesOutput
	ToLedgerPropertiesOutputWithContext(context.Context) LedgerPropertiesOutput
}

type LedgerPropertiesArgs struct {
	AadBasedSecurityPrincipals  AADBasedSecurityPrincipalArrayInput  `pulumi:"aadBasedSecurityPrincipals"`
	CertBasedSecurityPrincipals CertBasedSecurityPrincipalArrayInput `pulumi:"certBasedSecurityPrincipals"`
	LedgerType                  pulumi.StringPtrInput                `pulumi:"ledgerType"`
}

func (LedgerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LedgerProperties)(nil)).Elem()
}

func (i LedgerPropertiesArgs) ToLedgerPropertiesOutput() LedgerPropertiesOutput {
	return i.ToLedgerPropertiesOutputWithContext(context.Background())
}

func (i LedgerPropertiesArgs) ToLedgerPropertiesOutputWithContext(ctx context.Context) LedgerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LedgerPropertiesOutput)
}

func (i LedgerPropertiesArgs) ToLedgerPropertiesPtrOutput() LedgerPropertiesPtrOutput {
	return i.ToLedgerPropertiesPtrOutputWithContext(context.Background())
}

func (i LedgerPropertiesArgs) ToLedgerPropertiesPtrOutputWithContext(ctx context.Context) LedgerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LedgerPropertiesOutput).ToLedgerPropertiesPtrOutputWithContext(ctx)
}









type LedgerPropertiesPtrInput interface {
	pulumi.Input

	ToLedgerPropertiesPtrOutput() LedgerPropertiesPtrOutput
	ToLedgerPropertiesPtrOutputWithContext(context.Context) LedgerPropertiesPtrOutput
}

type ledgerPropertiesPtrType LedgerPropertiesArgs

func LedgerPropertiesPtr(v *LedgerPropertiesArgs) LedgerPropertiesPtrInput {
	return (*ledgerPropertiesPtrType)(v)
}

func (*ledgerPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LedgerProperties)(nil)).Elem()
}

func (i *ledgerPropertiesPtrType) ToLedgerPropertiesPtrOutput() LedgerPropertiesPtrOutput {
	return i.ToLedgerPropertiesPtrOutputWithContext(context.Background())
}

func (i *ledgerPropertiesPtrType) ToLedgerPropertiesPtrOutputWithContext(ctx context.Context) LedgerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LedgerPropertiesPtrOutput)
}

type LedgerPropertiesOutput struct{ *pulumi.OutputState }

func (LedgerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LedgerProperties)(nil)).Elem()
}

func (o LedgerPropertiesOutput) ToLedgerPropertiesOutput() LedgerPropertiesOutput {
	return o
}

func (o LedgerPropertiesOutput) ToLedgerPropertiesOutputWithContext(ctx context.Context) LedgerPropertiesOutput {
	return o
}

func (o LedgerPropertiesOutput) ToLedgerPropertiesPtrOutput() LedgerPropertiesPtrOutput {
	return o.ToLedgerPropertiesPtrOutputWithContext(context.Background())
}

func (o LedgerPropertiesOutput) ToLedgerPropertiesPtrOutputWithContext(ctx context.Context) LedgerPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LedgerProperties) *LedgerProperties {
		return &v
	}).(LedgerPropertiesPtrOutput)
}

func (o LedgerPropertiesOutput) AadBasedSecurityPrincipals() AADBasedSecurityPrincipalArrayOutput {
	return o.ApplyT(func(v LedgerProperties) []AADBasedSecurityPrincipal { return v.AadBasedSecurityPrincipals }).(AADBasedSecurityPrincipalArrayOutput)
}

func (o LedgerPropertiesOutput) CertBasedSecurityPrincipals() CertBasedSecurityPrincipalArrayOutput {
	return o.ApplyT(func(v LedgerProperties) []CertBasedSecurityPrincipal { return v.CertBasedSecurityPrincipals }).(CertBasedSecurityPrincipalArrayOutput)
}

func (o LedgerPropertiesOutput) LedgerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LedgerProperties) *string { return v.LedgerType }).(pulumi.StringPtrOutput)
}

type LedgerPropertiesPtrOutput struct{ *pulumi.OutputState }

func (LedgerPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LedgerProperties)(nil)).Elem()
}

func (o LedgerPropertiesPtrOutput) ToLedgerPropertiesPtrOutput() LedgerPropertiesPtrOutput {
	return o
}

func (o LedgerPropertiesPtrOutput) ToLedgerPropertiesPtrOutputWithContext(ctx context.Context) LedgerPropertiesPtrOutput {
	return o
}

func (o LedgerPropertiesPtrOutput) Elem() LedgerPropertiesOutput {
	return o.ApplyT(func(v *LedgerProperties) LedgerProperties {
		if v != nil {
			return *v
		}
		var ret LedgerProperties
		return ret
	}).(LedgerPropertiesOutput)
}

func (o LedgerPropertiesPtrOutput) AadBasedSecurityPrincipals() AADBasedSecurityPrincipalArrayOutput {
	return o.ApplyT(func(v *LedgerProperties) []AADBasedSecurityPrincipal {
		if v == nil {
			return nil
		}
		return v.AadBasedSecurityPrincipals
	}).(AADBasedSecurityPrincipalArrayOutput)
}

func (o LedgerPropertiesPtrOutput) CertBasedSecurityPrincipals() CertBasedSecurityPrincipalArrayOutput {
	return o.ApplyT(func(v *LedgerProperties) []CertBasedSecurityPrincipal {
		if v == nil {
			return nil
		}
		return v.CertBasedSecurityPrincipals
	}).(CertBasedSecurityPrincipalArrayOutput)
}

func (o LedgerPropertiesPtrOutput) LedgerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LedgerProperties) *string {
		if v == nil {
			return nil
		}
		return v.LedgerType
	}).(pulumi.StringPtrOutput)
}

type LedgerPropertiesResponse struct {
	AadBasedSecurityPrincipals  []AADBasedSecurityPrincipalResponse  `pulumi:"aadBasedSecurityPrincipals"`
	CertBasedSecurityPrincipals []CertBasedSecurityPrincipalResponse `pulumi:"certBasedSecurityPrincipals"`
	IdentityServiceUri          string                               `pulumi:"identityServiceUri"`
	LedgerInternalNamespace     string                               `pulumi:"ledgerInternalNamespace"`
	LedgerName                  string                               `pulumi:"ledgerName"`
	LedgerType                  *string                              `pulumi:"ledgerType"`
	LedgerUri                   string                               `pulumi:"ledgerUri"`
	ProvisioningState           string                               `pulumi:"provisioningState"`
}

type LedgerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LedgerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LedgerPropertiesResponse)(nil)).Elem()
}

func (o LedgerPropertiesResponseOutput) ToLedgerPropertiesResponseOutput() LedgerPropertiesResponseOutput {
	return o
}

func (o LedgerPropertiesResponseOutput) ToLedgerPropertiesResponseOutputWithContext(ctx context.Context) LedgerPropertiesResponseOutput {
	return o
}

func (o LedgerPropertiesResponseOutput) AadBasedSecurityPrincipals() AADBasedSecurityPrincipalResponseArrayOutput {
	return o.ApplyT(func(v LedgerPropertiesResponse) []AADBasedSecurityPrincipalResponse {
		return v.AadBasedSecurityPrincipals
	}).(AADBasedSecurityPrincipalResponseArrayOutput)
}

func (o LedgerPropertiesResponseOutput) CertBasedSecurityPrincipals() CertBasedSecurityPrincipalResponseArrayOutput {
	return o.ApplyT(func(v LedgerPropertiesResponse) []CertBasedSecurityPrincipalResponse {
		return v.CertBasedSecurityPrincipals
	}).(CertBasedSecurityPrincipalResponseArrayOutput)
}

func (o LedgerPropertiesResponseOutput) IdentityServiceUri() pulumi.StringOutput {
	return o.ApplyT(func(v LedgerPropertiesResponse) string { return v.IdentityServiceUri }).(pulumi.StringOutput)
}

func (o LedgerPropertiesResponseOutput) LedgerInternalNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v LedgerPropertiesResponse) string { return v.LedgerInternalNamespace }).(pulumi.StringOutput)
}

func (o LedgerPropertiesResponseOutput) LedgerName() pulumi.StringOutput {
	return o.ApplyT(func(v LedgerPropertiesResponse) string { return v.LedgerName }).(pulumi.StringOutput)
}

func (o LedgerPropertiesResponseOutput) LedgerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LedgerPropertiesResponse) *string { return v.LedgerType }).(pulumi.StringPtrOutput)
}

func (o LedgerPropertiesResponseOutput) LedgerUri() pulumi.StringOutput {
	return o.ApplyT(func(v LedgerPropertiesResponse) string { return v.LedgerUri }).(pulumi.StringOutput)
}

func (o LedgerPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LedgerPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ManagedCCFProperties struct {
	DeploymentType             *DeploymentType             `pulumi:"deploymentType"`
	MemberIdentityCertificates []MemberIdentityCertificate `pulumi:"memberIdentityCertificates"`
}





type ManagedCCFPropertiesInput interface {
	pulumi.Input

	ToManagedCCFPropertiesOutput() ManagedCCFPropertiesOutput
	ToManagedCCFPropertiesOutputWithContext(context.Context) ManagedCCFPropertiesOutput
}

type ManagedCCFPropertiesArgs struct {
	DeploymentType             DeploymentTypePtrInput              `pulumi:"deploymentType"`
	MemberIdentityCertificates MemberIdentityCertificateArrayInput `pulumi:"memberIdentityCertificates"`
}

func (ManagedCCFPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedCCFProperties)(nil)).Elem()
}

func (i ManagedCCFPropertiesArgs) ToManagedCCFPropertiesOutput() ManagedCCFPropertiesOutput {
	return i.ToManagedCCFPropertiesOutputWithContext(context.Background())
}

func (i ManagedCCFPropertiesArgs) ToManagedCCFPropertiesOutputWithContext(ctx context.Context) ManagedCCFPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedCCFPropertiesOutput)
}

func (i ManagedCCFPropertiesArgs) ToManagedCCFPropertiesPtrOutput() ManagedCCFPropertiesPtrOutput {
	return i.ToManagedCCFPropertiesPtrOutputWithContext(context.Background())
}

func (i ManagedCCFPropertiesArgs) ToManagedCCFPropertiesPtrOutputWithContext(ctx context.Context) ManagedCCFPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedCCFPropertiesOutput).ToManagedCCFPropertiesPtrOutputWithContext(ctx)
}









type ManagedCCFPropertiesPtrInput interface {
	pulumi.Input

	ToManagedCCFPropertiesPtrOutput() ManagedCCFPropertiesPtrOutput
	ToManagedCCFPropertiesPtrOutputWithContext(context.Context) ManagedCCFPropertiesPtrOutput
}

type managedCCFPropertiesPtrType ManagedCCFPropertiesArgs

func ManagedCCFPropertiesPtr(v *ManagedCCFPropertiesArgs) ManagedCCFPropertiesPtrInput {
	return (*managedCCFPropertiesPtrType)(v)
}

func (*managedCCFPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCCFProperties)(nil)).Elem()
}

func (i *managedCCFPropertiesPtrType) ToManagedCCFPropertiesPtrOutput() ManagedCCFPropertiesPtrOutput {
	return i.ToManagedCCFPropertiesPtrOutputWithContext(context.Background())
}

func (i *managedCCFPropertiesPtrType) ToManagedCCFPropertiesPtrOutputWithContext(ctx context.Context) ManagedCCFPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedCCFPropertiesPtrOutput)
}

type ManagedCCFPropertiesOutput struct{ *pulumi.OutputState }

func (ManagedCCFPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedCCFProperties)(nil)).Elem()
}

func (o ManagedCCFPropertiesOutput) ToManagedCCFPropertiesOutput() ManagedCCFPropertiesOutput {
	return o
}

func (o ManagedCCFPropertiesOutput) ToManagedCCFPropertiesOutputWithContext(ctx context.Context) ManagedCCFPropertiesOutput {
	return o
}

func (o ManagedCCFPropertiesOutput) ToManagedCCFPropertiesPtrOutput() ManagedCCFPropertiesPtrOutput {
	return o.ToManagedCCFPropertiesPtrOutputWithContext(context.Background())
}

func (o ManagedCCFPropertiesOutput) ToManagedCCFPropertiesPtrOutputWithContext(ctx context.Context) ManagedCCFPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedCCFProperties) *ManagedCCFProperties {
		return &v
	}).(ManagedCCFPropertiesPtrOutput)
}

func (o ManagedCCFPropertiesOutput) DeploymentType() DeploymentTypePtrOutput {
	return o.ApplyT(func(v ManagedCCFProperties) *DeploymentType { return v.DeploymentType }).(DeploymentTypePtrOutput)
}

func (o ManagedCCFPropertiesOutput) MemberIdentityCertificates() MemberIdentityCertificateArrayOutput {
	return o.ApplyT(func(v ManagedCCFProperties) []MemberIdentityCertificate { return v.MemberIdentityCertificates }).(MemberIdentityCertificateArrayOutput)
}

type ManagedCCFPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ManagedCCFPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCCFProperties)(nil)).Elem()
}

func (o ManagedCCFPropertiesPtrOutput) ToManagedCCFPropertiesPtrOutput() ManagedCCFPropertiesPtrOutput {
	return o
}

func (o ManagedCCFPropertiesPtrOutput) ToManagedCCFPropertiesPtrOutputWithContext(ctx context.Context) ManagedCCFPropertiesPtrOutput {
	return o
}

func (o ManagedCCFPropertiesPtrOutput) Elem() ManagedCCFPropertiesOutput {
	return o.ApplyT(func(v *ManagedCCFProperties) ManagedCCFProperties {
		if v != nil {
			return *v
		}
		var ret ManagedCCFProperties
		return ret
	}).(ManagedCCFPropertiesOutput)
}

func (o ManagedCCFPropertiesPtrOutput) DeploymentType() DeploymentTypePtrOutput {
	return o.ApplyT(func(v *ManagedCCFProperties) *DeploymentType {
		if v == nil {
			return nil
		}
		return v.DeploymentType
	}).(DeploymentTypePtrOutput)
}

func (o ManagedCCFPropertiesPtrOutput) MemberIdentityCertificates() MemberIdentityCertificateArrayOutput {
	return o.ApplyT(func(v *ManagedCCFProperties) []MemberIdentityCertificate {
		if v == nil {
			return nil
		}
		return v.MemberIdentityCertificates
	}).(MemberIdentityCertificateArrayOutput)
}

type ManagedCCFPropertiesResponse struct {
	AppName                    string                              `pulumi:"appName"`
	AppUri                     string                              `pulumi:"appUri"`
	DeploymentType             *DeploymentTypeResponse             `pulumi:"deploymentType"`
	IdentityServiceUri         string                              `pulumi:"identityServiceUri"`
	MemberIdentityCertificates []MemberIdentityCertificateResponse `pulumi:"memberIdentityCertificates"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
}

type ManagedCCFPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManagedCCFPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedCCFPropertiesResponse)(nil)).Elem()
}

func (o ManagedCCFPropertiesResponseOutput) ToManagedCCFPropertiesResponseOutput() ManagedCCFPropertiesResponseOutput {
	return o
}

func (o ManagedCCFPropertiesResponseOutput) ToManagedCCFPropertiesResponseOutputWithContext(ctx context.Context) ManagedCCFPropertiesResponseOutput {
	return o
}

func (o ManagedCCFPropertiesResponseOutput) AppName() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedCCFPropertiesResponse) string { return v.AppName }).(pulumi.StringOutput)
}

func (o ManagedCCFPropertiesResponseOutput) AppUri() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedCCFPropertiesResponse) string { return v.AppUri }).(pulumi.StringOutput)
}

func (o ManagedCCFPropertiesResponseOutput) DeploymentType() DeploymentTypeResponsePtrOutput {
	return o.ApplyT(func(v ManagedCCFPropertiesResponse) *DeploymentTypeResponse { return v.DeploymentType }).(DeploymentTypeResponsePtrOutput)
}

func (o ManagedCCFPropertiesResponseOutput) IdentityServiceUri() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedCCFPropertiesResponse) string { return v.IdentityServiceUri }).(pulumi.StringOutput)
}

func (o ManagedCCFPropertiesResponseOutput) MemberIdentityCertificates() MemberIdentityCertificateResponseArrayOutput {
	return o.ApplyT(func(v ManagedCCFPropertiesResponse) []MemberIdentityCertificateResponse {
		return v.MemberIdentityCertificates
	}).(MemberIdentityCertificateResponseArrayOutput)
}

func (o ManagedCCFPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedCCFPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type MemberIdentityCertificate struct {
	Certificate   *string           `pulumi:"certificate"`
	Encryptionkey *string           `pulumi:"encryptionkey"`
	Tags          []CertificateTags `pulumi:"tags"`
}





type MemberIdentityCertificateInput interface {
	pulumi.Input

	ToMemberIdentityCertificateOutput() MemberIdentityCertificateOutput
	ToMemberIdentityCertificateOutputWithContext(context.Context) MemberIdentityCertificateOutput
}

type MemberIdentityCertificateArgs struct {
	Certificate   pulumi.StringPtrInput     `pulumi:"certificate"`
	Encryptionkey pulumi.StringPtrInput     `pulumi:"encryptionkey"`
	Tags          CertificateTagsArrayInput `pulumi:"tags"`
}

func (MemberIdentityCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MemberIdentityCertificate)(nil)).Elem()
}

func (i MemberIdentityCertificateArgs) ToMemberIdentityCertificateOutput() MemberIdentityCertificateOutput {
	return i.ToMemberIdentityCertificateOutputWithContext(context.Background())
}

func (i MemberIdentityCertificateArgs) ToMemberIdentityCertificateOutputWithContext(ctx context.Context) MemberIdentityCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberIdentityCertificateOutput)
}





type MemberIdentityCertificateArrayInput interface {
	pulumi.Input

	ToMemberIdentityCertificateArrayOutput() MemberIdentityCertificateArrayOutput
	ToMemberIdentityCertificateArrayOutputWithContext(context.Context) MemberIdentityCertificateArrayOutput
}

type MemberIdentityCertificateArray []MemberIdentityCertificateInput

func (MemberIdentityCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MemberIdentityCertificate)(nil)).Elem()
}

func (i MemberIdentityCertificateArray) ToMemberIdentityCertificateArrayOutput() MemberIdentityCertificateArrayOutput {
	return i.ToMemberIdentityCertificateArrayOutputWithContext(context.Background())
}

func (i MemberIdentityCertificateArray) ToMemberIdentityCertificateArrayOutputWithContext(ctx context.Context) MemberIdentityCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberIdentityCertificateArrayOutput)
}

type MemberIdentityCertificateOutput struct{ *pulumi.OutputState }

func (MemberIdentityCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MemberIdentityCertificate)(nil)).Elem()
}

func (o MemberIdentityCertificateOutput) ToMemberIdentityCertificateOutput() MemberIdentityCertificateOutput {
	return o
}

func (o MemberIdentityCertificateOutput) ToMemberIdentityCertificateOutputWithContext(ctx context.Context) MemberIdentityCertificateOutput {
	return o
}

func (o MemberIdentityCertificateOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MemberIdentityCertificate) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o MemberIdentityCertificateOutput) Encryptionkey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MemberIdentityCertificate) *string { return v.Encryptionkey }).(pulumi.StringPtrOutput)
}

func (o MemberIdentityCertificateOutput) Tags() CertificateTagsArrayOutput {
	return o.ApplyT(func(v MemberIdentityCertificate) []CertificateTags { return v.Tags }).(CertificateTagsArrayOutput)
}

type MemberIdentityCertificateArrayOutput struct{ *pulumi.OutputState }

func (MemberIdentityCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MemberIdentityCertificate)(nil)).Elem()
}

func (o MemberIdentityCertificateArrayOutput) ToMemberIdentityCertificateArrayOutput() MemberIdentityCertificateArrayOutput {
	return o
}

func (o MemberIdentityCertificateArrayOutput) ToMemberIdentityCertificateArrayOutputWithContext(ctx context.Context) MemberIdentityCertificateArrayOutput {
	return o
}

func (o MemberIdentityCertificateArrayOutput) Index(i pulumi.IntInput) MemberIdentityCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MemberIdentityCertificate {
		return vs[0].([]MemberIdentityCertificate)[vs[1].(int)]
	}).(MemberIdentityCertificateOutput)
}

type MemberIdentityCertificateResponse struct {
	Certificate   *string                   `pulumi:"certificate"`
	Encryptionkey *string                   `pulumi:"encryptionkey"`
	Tags          []CertificateTagsResponse `pulumi:"tags"`
}

type MemberIdentityCertificateResponseOutput struct{ *pulumi.OutputState }

func (MemberIdentityCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MemberIdentityCertificateResponse)(nil)).Elem()
}

func (o MemberIdentityCertificateResponseOutput) ToMemberIdentityCertificateResponseOutput() MemberIdentityCertificateResponseOutput {
	return o
}

func (o MemberIdentityCertificateResponseOutput) ToMemberIdentityCertificateResponseOutputWithContext(ctx context.Context) MemberIdentityCertificateResponseOutput {
	return o
}

func (o MemberIdentityCertificateResponseOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MemberIdentityCertificateResponse) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o MemberIdentityCertificateResponseOutput) Encryptionkey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MemberIdentityCertificateResponse) *string { return v.Encryptionkey }).(pulumi.StringPtrOutput)
}

func (o MemberIdentityCertificateResponseOutput) Tags() CertificateTagsResponseArrayOutput {
	return o.ApplyT(func(v MemberIdentityCertificateResponse) []CertificateTagsResponse { return v.Tags }).(CertificateTagsResponseArrayOutput)
}

type MemberIdentityCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (MemberIdentityCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MemberIdentityCertificateResponse)(nil)).Elem()
}

func (o MemberIdentityCertificateResponseArrayOutput) ToMemberIdentityCertificateResponseArrayOutput() MemberIdentityCertificateResponseArrayOutput {
	return o
}

func (o MemberIdentityCertificateResponseArrayOutput) ToMemberIdentityCertificateResponseArrayOutputWithContext(ctx context.Context) MemberIdentityCertificateResponseArrayOutput {
	return o
}

func (o MemberIdentityCertificateResponseArrayOutput) Index(i pulumi.IntInput) MemberIdentityCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MemberIdentityCertificateResponse {
		return vs[0].([]MemberIdentityCertificateResponse)[vs[1].(int)]
	}).(MemberIdentityCertificateResponseOutput)
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
	pulumi.RegisterOutputType(AADBasedSecurityPrincipalOutput{})
	pulumi.RegisterOutputType(AADBasedSecurityPrincipalArrayOutput{})
	pulumi.RegisterOutputType(AADBasedSecurityPrincipalResponseOutput{})
	pulumi.RegisterOutputType(AADBasedSecurityPrincipalResponseArrayOutput{})
	pulumi.RegisterOutputType(CertBasedSecurityPrincipalOutput{})
	pulumi.RegisterOutputType(CertBasedSecurityPrincipalArrayOutput{})
	pulumi.RegisterOutputType(CertBasedSecurityPrincipalResponseOutput{})
	pulumi.RegisterOutputType(CertBasedSecurityPrincipalResponseArrayOutput{})
	pulumi.RegisterOutputType(CertificateTagsOutput{})
	pulumi.RegisterOutputType(CertificateTagsArrayOutput{})
	pulumi.RegisterOutputType(CertificateTagsResponseOutput{})
	pulumi.RegisterOutputType(CertificateTagsResponseArrayOutput{})
	pulumi.RegisterOutputType(DeploymentTypeOutput{})
	pulumi.RegisterOutputType(DeploymentTypePtrOutput{})
	pulumi.RegisterOutputType(DeploymentTypeResponseOutput{})
	pulumi.RegisterOutputType(DeploymentTypeResponsePtrOutput{})
	pulumi.RegisterOutputType(LedgerPropertiesOutput{})
	pulumi.RegisterOutputType(LedgerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(LedgerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManagedCCFPropertiesOutput{})
	pulumi.RegisterOutputType(ManagedCCFPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagedCCFPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MemberIdentityCertificateOutput{})
	pulumi.RegisterOutputType(MemberIdentityCertificateArrayOutput{})
	pulumi.RegisterOutputType(MemberIdentityCertificateResponseOutput{})
	pulumi.RegisterOutputType(MemberIdentityCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

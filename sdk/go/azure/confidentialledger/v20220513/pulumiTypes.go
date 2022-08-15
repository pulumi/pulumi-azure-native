


package v20220513

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
	pulumi.RegisterOutputType(LedgerPropertiesOutput{})
	pulumi.RegisterOutputType(LedgerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(LedgerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

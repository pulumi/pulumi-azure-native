


package v20190501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AkamaiAccessControl struct {
	AkamaiSignatureHeaderAuthenticationKeyList []AkamaiSignatureHeaderAuthenticationKey `pulumi:"akamaiSignatureHeaderAuthenticationKeyList"`
}





type AkamaiAccessControlInput interface {
	pulumi.Input

	ToAkamaiAccessControlOutput() AkamaiAccessControlOutput
	ToAkamaiAccessControlOutputWithContext(context.Context) AkamaiAccessControlOutput
}

type AkamaiAccessControlArgs struct {
	AkamaiSignatureHeaderAuthenticationKeyList AkamaiSignatureHeaderAuthenticationKeyArrayInput `pulumi:"akamaiSignatureHeaderAuthenticationKeyList"`
}

func (AkamaiAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiAccessControl)(nil)).Elem()
}

func (i AkamaiAccessControlArgs) ToAkamaiAccessControlOutput() AkamaiAccessControlOutput {
	return i.ToAkamaiAccessControlOutputWithContext(context.Background())
}

func (i AkamaiAccessControlArgs) ToAkamaiAccessControlOutputWithContext(ctx context.Context) AkamaiAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlOutput)
}

func (i AkamaiAccessControlArgs) ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput {
	return i.ToAkamaiAccessControlPtrOutputWithContext(context.Background())
}

func (i AkamaiAccessControlArgs) ToAkamaiAccessControlPtrOutputWithContext(ctx context.Context) AkamaiAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlOutput).ToAkamaiAccessControlPtrOutputWithContext(ctx)
}









type AkamaiAccessControlPtrInput interface {
	pulumi.Input

	ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput
	ToAkamaiAccessControlPtrOutputWithContext(context.Context) AkamaiAccessControlPtrOutput
}

type akamaiAccessControlPtrType AkamaiAccessControlArgs

func AkamaiAccessControlPtr(v *AkamaiAccessControlArgs) AkamaiAccessControlPtrInput {
	return (*akamaiAccessControlPtrType)(v)
}

func (*akamaiAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AkamaiAccessControl)(nil)).Elem()
}

func (i *akamaiAccessControlPtrType) ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput {
	return i.ToAkamaiAccessControlPtrOutputWithContext(context.Background())
}

func (i *akamaiAccessControlPtrType) ToAkamaiAccessControlPtrOutputWithContext(ctx context.Context) AkamaiAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlPtrOutput)
}

type AkamaiAccessControlOutput struct{ *pulumi.OutputState }

func (AkamaiAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiAccessControl)(nil)).Elem()
}

func (o AkamaiAccessControlOutput) ToAkamaiAccessControlOutput() AkamaiAccessControlOutput {
	return o
}

func (o AkamaiAccessControlOutput) ToAkamaiAccessControlOutputWithContext(ctx context.Context) AkamaiAccessControlOutput {
	return o
}

func (o AkamaiAccessControlOutput) ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput {
	return o.ToAkamaiAccessControlPtrOutputWithContext(context.Background())
}

func (o AkamaiAccessControlOutput) ToAkamaiAccessControlPtrOutputWithContext(ctx context.Context) AkamaiAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AkamaiAccessControl) *AkamaiAccessControl {
		return &v
	}).(AkamaiAccessControlPtrOutput)
}

func (o AkamaiAccessControlOutput) AkamaiSignatureHeaderAuthenticationKeyList() AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return o.ApplyT(func(v AkamaiAccessControl) []AkamaiSignatureHeaderAuthenticationKey {
		return v.AkamaiSignatureHeaderAuthenticationKeyList
	}).(AkamaiSignatureHeaderAuthenticationKeyArrayOutput)
}

type AkamaiAccessControlPtrOutput struct{ *pulumi.OutputState }

func (AkamaiAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AkamaiAccessControl)(nil)).Elem()
}

func (o AkamaiAccessControlPtrOutput) ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput {
	return o
}

func (o AkamaiAccessControlPtrOutput) ToAkamaiAccessControlPtrOutputWithContext(ctx context.Context) AkamaiAccessControlPtrOutput {
	return o
}

func (o AkamaiAccessControlPtrOutput) Elem() AkamaiAccessControlOutput {
	return o.ApplyT(func(v *AkamaiAccessControl) AkamaiAccessControl {
		if v != nil {
			return *v
		}
		var ret AkamaiAccessControl
		return ret
	}).(AkamaiAccessControlOutput)
}

func (o AkamaiAccessControlPtrOutput) AkamaiSignatureHeaderAuthenticationKeyList() AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return o.ApplyT(func(v *AkamaiAccessControl) []AkamaiSignatureHeaderAuthenticationKey {
		if v == nil {
			return nil
		}
		return v.AkamaiSignatureHeaderAuthenticationKeyList
	}).(AkamaiSignatureHeaderAuthenticationKeyArrayOutput)
}

type AkamaiAccessControlResponse struct {
	AkamaiSignatureHeaderAuthenticationKeyList []AkamaiSignatureHeaderAuthenticationKeyResponse `pulumi:"akamaiSignatureHeaderAuthenticationKeyList"`
}





type AkamaiAccessControlResponseInput interface {
	pulumi.Input

	ToAkamaiAccessControlResponseOutput() AkamaiAccessControlResponseOutput
	ToAkamaiAccessControlResponseOutputWithContext(context.Context) AkamaiAccessControlResponseOutput
}

type AkamaiAccessControlResponseArgs struct {
	AkamaiSignatureHeaderAuthenticationKeyList AkamaiSignatureHeaderAuthenticationKeyResponseArrayInput `pulumi:"akamaiSignatureHeaderAuthenticationKeyList"`
}

func (AkamaiAccessControlResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiAccessControlResponse)(nil)).Elem()
}

func (i AkamaiAccessControlResponseArgs) ToAkamaiAccessControlResponseOutput() AkamaiAccessControlResponseOutput {
	return i.ToAkamaiAccessControlResponseOutputWithContext(context.Background())
}

func (i AkamaiAccessControlResponseArgs) ToAkamaiAccessControlResponseOutputWithContext(ctx context.Context) AkamaiAccessControlResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlResponseOutput)
}

func (i AkamaiAccessControlResponseArgs) ToAkamaiAccessControlResponsePtrOutput() AkamaiAccessControlResponsePtrOutput {
	return i.ToAkamaiAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i AkamaiAccessControlResponseArgs) ToAkamaiAccessControlResponsePtrOutputWithContext(ctx context.Context) AkamaiAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlResponseOutput).ToAkamaiAccessControlResponsePtrOutputWithContext(ctx)
}









type AkamaiAccessControlResponsePtrInput interface {
	pulumi.Input

	ToAkamaiAccessControlResponsePtrOutput() AkamaiAccessControlResponsePtrOutput
	ToAkamaiAccessControlResponsePtrOutputWithContext(context.Context) AkamaiAccessControlResponsePtrOutput
}

type akamaiAccessControlResponsePtrType AkamaiAccessControlResponseArgs

func AkamaiAccessControlResponsePtr(v *AkamaiAccessControlResponseArgs) AkamaiAccessControlResponsePtrInput {
	return (*akamaiAccessControlResponsePtrType)(v)
}

func (*akamaiAccessControlResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AkamaiAccessControlResponse)(nil)).Elem()
}

func (i *akamaiAccessControlResponsePtrType) ToAkamaiAccessControlResponsePtrOutput() AkamaiAccessControlResponsePtrOutput {
	return i.ToAkamaiAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i *akamaiAccessControlResponsePtrType) ToAkamaiAccessControlResponsePtrOutputWithContext(ctx context.Context) AkamaiAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlResponsePtrOutput)
}

type AkamaiAccessControlResponseOutput struct{ *pulumi.OutputState }

func (AkamaiAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiAccessControlResponse)(nil)).Elem()
}

func (o AkamaiAccessControlResponseOutput) ToAkamaiAccessControlResponseOutput() AkamaiAccessControlResponseOutput {
	return o
}

func (o AkamaiAccessControlResponseOutput) ToAkamaiAccessControlResponseOutputWithContext(ctx context.Context) AkamaiAccessControlResponseOutput {
	return o
}

func (o AkamaiAccessControlResponseOutput) ToAkamaiAccessControlResponsePtrOutput() AkamaiAccessControlResponsePtrOutput {
	return o.ToAkamaiAccessControlResponsePtrOutputWithContext(context.Background())
}

func (o AkamaiAccessControlResponseOutput) ToAkamaiAccessControlResponsePtrOutputWithContext(ctx context.Context) AkamaiAccessControlResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AkamaiAccessControlResponse) *AkamaiAccessControlResponse {
		return &v
	}).(AkamaiAccessControlResponsePtrOutput)
}

func (o AkamaiAccessControlResponseOutput) AkamaiSignatureHeaderAuthenticationKeyList() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return o.ApplyT(func(v AkamaiAccessControlResponse) []AkamaiSignatureHeaderAuthenticationKeyResponse {
		return v.AkamaiSignatureHeaderAuthenticationKeyList
	}).(AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput)
}

type AkamaiAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (AkamaiAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AkamaiAccessControlResponse)(nil)).Elem()
}

func (o AkamaiAccessControlResponsePtrOutput) ToAkamaiAccessControlResponsePtrOutput() AkamaiAccessControlResponsePtrOutput {
	return o
}

func (o AkamaiAccessControlResponsePtrOutput) ToAkamaiAccessControlResponsePtrOutputWithContext(ctx context.Context) AkamaiAccessControlResponsePtrOutput {
	return o
}

func (o AkamaiAccessControlResponsePtrOutput) Elem() AkamaiAccessControlResponseOutput {
	return o.ApplyT(func(v *AkamaiAccessControlResponse) AkamaiAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret AkamaiAccessControlResponse
		return ret
	}).(AkamaiAccessControlResponseOutput)
}

func (o AkamaiAccessControlResponsePtrOutput) AkamaiSignatureHeaderAuthenticationKeyList() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return o.ApplyT(func(v *AkamaiAccessControlResponse) []AkamaiSignatureHeaderAuthenticationKeyResponse {
		if v == nil {
			return nil
		}
		return v.AkamaiSignatureHeaderAuthenticationKeyList
	}).(AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput)
}

type AkamaiSignatureHeaderAuthenticationKey struct {
	Base64Key  *string `pulumi:"base64Key"`
	Expiration *string `pulumi:"expiration"`
	Identifier *string `pulumi:"identifier"`
}





type AkamaiSignatureHeaderAuthenticationKeyInput interface {
	pulumi.Input

	ToAkamaiSignatureHeaderAuthenticationKeyOutput() AkamaiSignatureHeaderAuthenticationKeyOutput
	ToAkamaiSignatureHeaderAuthenticationKeyOutputWithContext(context.Context) AkamaiSignatureHeaderAuthenticationKeyOutput
}

type AkamaiSignatureHeaderAuthenticationKeyArgs struct {
	Base64Key  pulumi.StringPtrInput `pulumi:"base64Key"`
	Expiration pulumi.StringPtrInput `pulumi:"expiration"`
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
}

func (AkamaiSignatureHeaderAuthenticationKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKey)(nil)).Elem()
}

func (i AkamaiSignatureHeaderAuthenticationKeyArgs) ToAkamaiSignatureHeaderAuthenticationKeyOutput() AkamaiSignatureHeaderAuthenticationKeyOutput {
	return i.ToAkamaiSignatureHeaderAuthenticationKeyOutputWithContext(context.Background())
}

func (i AkamaiSignatureHeaderAuthenticationKeyArgs) ToAkamaiSignatureHeaderAuthenticationKeyOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiSignatureHeaderAuthenticationKeyOutput)
}





type AkamaiSignatureHeaderAuthenticationKeyArrayInput interface {
	pulumi.Input

	ToAkamaiSignatureHeaderAuthenticationKeyArrayOutput() AkamaiSignatureHeaderAuthenticationKeyArrayOutput
	ToAkamaiSignatureHeaderAuthenticationKeyArrayOutputWithContext(context.Context) AkamaiSignatureHeaderAuthenticationKeyArrayOutput
}

type AkamaiSignatureHeaderAuthenticationKeyArray []AkamaiSignatureHeaderAuthenticationKeyInput

func (AkamaiSignatureHeaderAuthenticationKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AkamaiSignatureHeaderAuthenticationKey)(nil)).Elem()
}

func (i AkamaiSignatureHeaderAuthenticationKeyArray) ToAkamaiSignatureHeaderAuthenticationKeyArrayOutput() AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return i.ToAkamaiSignatureHeaderAuthenticationKeyArrayOutputWithContext(context.Background())
}

func (i AkamaiSignatureHeaderAuthenticationKeyArray) ToAkamaiSignatureHeaderAuthenticationKeyArrayOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiSignatureHeaderAuthenticationKeyArrayOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyOutput struct{ *pulumi.OutputState }

func (AkamaiSignatureHeaderAuthenticationKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKey)(nil)).Elem()
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) ToAkamaiSignatureHeaderAuthenticationKeyOutput() AkamaiSignatureHeaderAuthenticationKeyOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) ToAkamaiSignatureHeaderAuthenticationKeyOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) Base64Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKey) *string { return v.Base64Key }).(pulumi.StringPtrOutput)
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) Expiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKey) *string { return v.Expiration }).(pulumi.StringPtrOutput)
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKey) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyArrayOutput struct{ *pulumi.OutputState }

func (AkamaiSignatureHeaderAuthenticationKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AkamaiSignatureHeaderAuthenticationKey)(nil)).Elem()
}

func (o AkamaiSignatureHeaderAuthenticationKeyArrayOutput) ToAkamaiSignatureHeaderAuthenticationKeyArrayOutput() AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyArrayOutput) ToAkamaiSignatureHeaderAuthenticationKeyArrayOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyArrayOutput) Index(i pulumi.IntInput) AkamaiSignatureHeaderAuthenticationKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AkamaiSignatureHeaderAuthenticationKey {
		return vs[0].([]AkamaiSignatureHeaderAuthenticationKey)[vs[1].(int)]
	}).(AkamaiSignatureHeaderAuthenticationKeyOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyResponse struct {
	Base64Key  *string `pulumi:"base64Key"`
	Expiration *string `pulumi:"expiration"`
	Identifier *string `pulumi:"identifier"`
}





type AkamaiSignatureHeaderAuthenticationKeyResponseInput interface {
	pulumi.Input

	ToAkamaiSignatureHeaderAuthenticationKeyResponseOutput() AkamaiSignatureHeaderAuthenticationKeyResponseOutput
	ToAkamaiSignatureHeaderAuthenticationKeyResponseOutputWithContext(context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseOutput
}

type AkamaiSignatureHeaderAuthenticationKeyResponseArgs struct {
	Base64Key  pulumi.StringPtrInput `pulumi:"base64Key"`
	Expiration pulumi.StringPtrInput `pulumi:"expiration"`
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
}

func (AkamaiSignatureHeaderAuthenticationKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyResponse)(nil)).Elem()
}

func (i AkamaiSignatureHeaderAuthenticationKeyResponseArgs) ToAkamaiSignatureHeaderAuthenticationKeyResponseOutput() AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return i.ToAkamaiSignatureHeaderAuthenticationKeyResponseOutputWithContext(context.Background())
}

func (i AkamaiSignatureHeaderAuthenticationKeyResponseArgs) ToAkamaiSignatureHeaderAuthenticationKeyResponseOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiSignatureHeaderAuthenticationKeyResponseOutput)
}





type AkamaiSignatureHeaderAuthenticationKeyResponseArrayInput interface {
	pulumi.Input

	ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput
	ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutputWithContext(context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput
}

type AkamaiSignatureHeaderAuthenticationKeyResponseArray []AkamaiSignatureHeaderAuthenticationKeyResponseInput

func (AkamaiSignatureHeaderAuthenticationKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AkamaiSignatureHeaderAuthenticationKeyResponse)(nil)).Elem()
}

func (i AkamaiSignatureHeaderAuthenticationKeyResponseArray) ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return i.ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutputWithContext(context.Background())
}

func (i AkamaiSignatureHeaderAuthenticationKeyResponseArray) ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyResponseOutput struct{ *pulumi.OutputState }

func (AkamaiSignatureHeaderAuthenticationKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyResponse)(nil)).Elem()
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) ToAkamaiSignatureHeaderAuthenticationKeyResponseOutput() AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) ToAkamaiSignatureHeaderAuthenticationKeyResponseOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) Base64Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKeyResponse) *string { return v.Base64Key }).(pulumi.StringPtrOutput)
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) Expiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKeyResponse) *string { return v.Expiration }).(pulumi.StringPtrOutput)
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKeyResponse) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AkamaiSignatureHeaderAuthenticationKeyResponse)(nil)).Elem()
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput) ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput) ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput) Index(i pulumi.IntInput) AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AkamaiSignatureHeaderAuthenticationKeyResponse {
		return vs[0].([]AkamaiSignatureHeaderAuthenticationKeyResponse)[vs[1].(int)]
	}).(AkamaiSignatureHeaderAuthenticationKeyResponseOutput)
}

type CrossSiteAccessPolicies struct {
	ClientAccessPolicy *string `pulumi:"clientAccessPolicy"`
	CrossDomainPolicy  *string `pulumi:"crossDomainPolicy"`
}





type CrossSiteAccessPoliciesInput interface {
	pulumi.Input

	ToCrossSiteAccessPoliciesOutput() CrossSiteAccessPoliciesOutput
	ToCrossSiteAccessPoliciesOutputWithContext(context.Context) CrossSiteAccessPoliciesOutput
}

type CrossSiteAccessPoliciesArgs struct {
	ClientAccessPolicy pulumi.StringPtrInput `pulumi:"clientAccessPolicy"`
	CrossDomainPolicy  pulumi.StringPtrInput `pulumi:"crossDomainPolicy"`
}

func (CrossSiteAccessPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossSiteAccessPolicies)(nil)).Elem()
}

func (i CrossSiteAccessPoliciesArgs) ToCrossSiteAccessPoliciesOutput() CrossSiteAccessPoliciesOutput {
	return i.ToCrossSiteAccessPoliciesOutputWithContext(context.Background())
}

func (i CrossSiteAccessPoliciesArgs) ToCrossSiteAccessPoliciesOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesOutput)
}

func (i CrossSiteAccessPoliciesArgs) ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput {
	return i.ToCrossSiteAccessPoliciesPtrOutputWithContext(context.Background())
}

func (i CrossSiteAccessPoliciesArgs) ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesOutput).ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx)
}









type CrossSiteAccessPoliciesPtrInput interface {
	pulumi.Input

	ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput
	ToCrossSiteAccessPoliciesPtrOutputWithContext(context.Context) CrossSiteAccessPoliciesPtrOutput
}

type crossSiteAccessPoliciesPtrType CrossSiteAccessPoliciesArgs

func CrossSiteAccessPoliciesPtr(v *CrossSiteAccessPoliciesArgs) CrossSiteAccessPoliciesPtrInput {
	return (*crossSiteAccessPoliciesPtrType)(v)
}

func (*crossSiteAccessPoliciesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossSiteAccessPolicies)(nil)).Elem()
}

func (i *crossSiteAccessPoliciesPtrType) ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput {
	return i.ToCrossSiteAccessPoliciesPtrOutputWithContext(context.Background())
}

func (i *crossSiteAccessPoliciesPtrType) ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesPtrOutput)
}

type CrossSiteAccessPoliciesOutput struct{ *pulumi.OutputState }

func (CrossSiteAccessPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossSiteAccessPolicies)(nil)).Elem()
}

func (o CrossSiteAccessPoliciesOutput) ToCrossSiteAccessPoliciesOutput() CrossSiteAccessPoliciesOutput {
	return o
}

func (o CrossSiteAccessPoliciesOutput) ToCrossSiteAccessPoliciesOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesOutput {
	return o
}

func (o CrossSiteAccessPoliciesOutput) ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput {
	return o.ToCrossSiteAccessPoliciesPtrOutputWithContext(context.Background())
}

func (o CrossSiteAccessPoliciesOutput) ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CrossSiteAccessPolicies) *CrossSiteAccessPolicies {
		return &v
	}).(CrossSiteAccessPoliciesPtrOutput)
}

func (o CrossSiteAccessPoliciesOutput) ClientAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CrossSiteAccessPolicies) *string { return v.ClientAccessPolicy }).(pulumi.StringPtrOutput)
}

func (o CrossSiteAccessPoliciesOutput) CrossDomainPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CrossSiteAccessPolicies) *string { return v.CrossDomainPolicy }).(pulumi.StringPtrOutput)
}

type CrossSiteAccessPoliciesPtrOutput struct{ *pulumi.OutputState }

func (CrossSiteAccessPoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossSiteAccessPolicies)(nil)).Elem()
}

func (o CrossSiteAccessPoliciesPtrOutput) ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput {
	return o
}

func (o CrossSiteAccessPoliciesPtrOutput) ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesPtrOutput {
	return o
}

func (o CrossSiteAccessPoliciesPtrOutput) Elem() CrossSiteAccessPoliciesOutput {
	return o.ApplyT(func(v *CrossSiteAccessPolicies) CrossSiteAccessPolicies {
		if v != nil {
			return *v
		}
		var ret CrossSiteAccessPolicies
		return ret
	}).(CrossSiteAccessPoliciesOutput)
}

func (o CrossSiteAccessPoliciesPtrOutput) ClientAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrossSiteAccessPolicies) *string {
		if v == nil {
			return nil
		}
		return v.ClientAccessPolicy
	}).(pulumi.StringPtrOutput)
}

func (o CrossSiteAccessPoliciesPtrOutput) CrossDomainPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrossSiteAccessPolicies) *string {
		if v == nil {
			return nil
		}
		return v.CrossDomainPolicy
	}).(pulumi.StringPtrOutput)
}

type CrossSiteAccessPoliciesResponse struct {
	ClientAccessPolicy *string `pulumi:"clientAccessPolicy"`
	CrossDomainPolicy  *string `pulumi:"crossDomainPolicy"`
}





type CrossSiteAccessPoliciesResponseInput interface {
	pulumi.Input

	ToCrossSiteAccessPoliciesResponseOutput() CrossSiteAccessPoliciesResponseOutput
	ToCrossSiteAccessPoliciesResponseOutputWithContext(context.Context) CrossSiteAccessPoliciesResponseOutput
}

type CrossSiteAccessPoliciesResponseArgs struct {
	ClientAccessPolicy pulumi.StringPtrInput `pulumi:"clientAccessPolicy"`
	CrossDomainPolicy  pulumi.StringPtrInput `pulumi:"crossDomainPolicy"`
}

func (CrossSiteAccessPoliciesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossSiteAccessPoliciesResponse)(nil)).Elem()
}

func (i CrossSiteAccessPoliciesResponseArgs) ToCrossSiteAccessPoliciesResponseOutput() CrossSiteAccessPoliciesResponseOutput {
	return i.ToCrossSiteAccessPoliciesResponseOutputWithContext(context.Background())
}

func (i CrossSiteAccessPoliciesResponseArgs) ToCrossSiteAccessPoliciesResponseOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesResponseOutput)
}

func (i CrossSiteAccessPoliciesResponseArgs) ToCrossSiteAccessPoliciesResponsePtrOutput() CrossSiteAccessPoliciesResponsePtrOutput {
	return i.ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(context.Background())
}

func (i CrossSiteAccessPoliciesResponseArgs) ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesResponseOutput).ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(ctx)
}









type CrossSiteAccessPoliciesResponsePtrInput interface {
	pulumi.Input

	ToCrossSiteAccessPoliciesResponsePtrOutput() CrossSiteAccessPoliciesResponsePtrOutput
	ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(context.Context) CrossSiteAccessPoliciesResponsePtrOutput
}

type crossSiteAccessPoliciesResponsePtrType CrossSiteAccessPoliciesResponseArgs

func CrossSiteAccessPoliciesResponsePtr(v *CrossSiteAccessPoliciesResponseArgs) CrossSiteAccessPoliciesResponsePtrInput {
	return (*crossSiteAccessPoliciesResponsePtrType)(v)
}

func (*crossSiteAccessPoliciesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossSiteAccessPoliciesResponse)(nil)).Elem()
}

func (i *crossSiteAccessPoliciesResponsePtrType) ToCrossSiteAccessPoliciesResponsePtrOutput() CrossSiteAccessPoliciesResponsePtrOutput {
	return i.ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(context.Background())
}

func (i *crossSiteAccessPoliciesResponsePtrType) ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesResponsePtrOutput)
}

type CrossSiteAccessPoliciesResponseOutput struct{ *pulumi.OutputState }

func (CrossSiteAccessPoliciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossSiteAccessPoliciesResponse)(nil)).Elem()
}

func (o CrossSiteAccessPoliciesResponseOutput) ToCrossSiteAccessPoliciesResponseOutput() CrossSiteAccessPoliciesResponseOutput {
	return o
}

func (o CrossSiteAccessPoliciesResponseOutput) ToCrossSiteAccessPoliciesResponseOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponseOutput {
	return o
}

func (o CrossSiteAccessPoliciesResponseOutput) ToCrossSiteAccessPoliciesResponsePtrOutput() CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(context.Background())
}

func (o CrossSiteAccessPoliciesResponseOutput) ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CrossSiteAccessPoliciesResponse) *CrossSiteAccessPoliciesResponse {
		return &v
	}).(CrossSiteAccessPoliciesResponsePtrOutput)
}

func (o CrossSiteAccessPoliciesResponseOutput) ClientAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CrossSiteAccessPoliciesResponse) *string { return v.ClientAccessPolicy }).(pulumi.StringPtrOutput)
}

func (o CrossSiteAccessPoliciesResponseOutput) CrossDomainPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CrossSiteAccessPoliciesResponse) *string { return v.CrossDomainPolicy }).(pulumi.StringPtrOutput)
}

type CrossSiteAccessPoliciesResponsePtrOutput struct{ *pulumi.OutputState }

func (CrossSiteAccessPoliciesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossSiteAccessPoliciesResponse)(nil)).Elem()
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) ToCrossSiteAccessPoliciesResponsePtrOutput() CrossSiteAccessPoliciesResponsePtrOutput {
	return o
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponsePtrOutput {
	return o
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) Elem() CrossSiteAccessPoliciesResponseOutput {
	return o.ApplyT(func(v *CrossSiteAccessPoliciesResponse) CrossSiteAccessPoliciesResponse {
		if v != nil {
			return *v
		}
		var ret CrossSiteAccessPoliciesResponse
		return ret
	}).(CrossSiteAccessPoliciesResponseOutput)
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) ClientAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrossSiteAccessPoliciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientAccessPolicy
	}).(pulumi.StringPtrOutput)
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) CrossDomainPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrossSiteAccessPoliciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CrossDomainPolicy
	}).(pulumi.StringPtrOutput)
}

type Hls struct {
	FragmentsPerTsSegment *int `pulumi:"fragmentsPerTsSegment"`
}





type HlsInput interface {
	pulumi.Input

	ToHlsOutput() HlsOutput
	ToHlsOutputWithContext(context.Context) HlsOutput
}

type HlsArgs struct {
	FragmentsPerTsSegment pulumi.IntPtrInput `pulumi:"fragmentsPerTsSegment"`
}

func (HlsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Hls)(nil)).Elem()
}

func (i HlsArgs) ToHlsOutput() HlsOutput {
	return i.ToHlsOutputWithContext(context.Background())
}

func (i HlsArgs) ToHlsOutputWithContext(ctx context.Context) HlsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsOutput)
}

func (i HlsArgs) ToHlsPtrOutput() HlsPtrOutput {
	return i.ToHlsPtrOutputWithContext(context.Background())
}

func (i HlsArgs) ToHlsPtrOutputWithContext(ctx context.Context) HlsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsOutput).ToHlsPtrOutputWithContext(ctx)
}









type HlsPtrInput interface {
	pulumi.Input

	ToHlsPtrOutput() HlsPtrOutput
	ToHlsPtrOutputWithContext(context.Context) HlsPtrOutput
}

type hlsPtrType HlsArgs

func HlsPtr(v *HlsArgs) HlsPtrInput {
	return (*hlsPtrType)(v)
}

func (*hlsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Hls)(nil)).Elem()
}

func (i *hlsPtrType) ToHlsPtrOutput() HlsPtrOutput {
	return i.ToHlsPtrOutputWithContext(context.Background())
}

func (i *hlsPtrType) ToHlsPtrOutputWithContext(ctx context.Context) HlsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsPtrOutput)
}

type HlsOutput struct{ *pulumi.OutputState }

func (HlsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Hls)(nil)).Elem()
}

func (o HlsOutput) ToHlsOutput() HlsOutput {
	return o
}

func (o HlsOutput) ToHlsOutputWithContext(ctx context.Context) HlsOutput {
	return o
}

func (o HlsOutput) ToHlsPtrOutput() HlsPtrOutput {
	return o.ToHlsPtrOutputWithContext(context.Background())
}

func (o HlsOutput) ToHlsPtrOutputWithContext(ctx context.Context) HlsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Hls) *Hls {
		return &v
	}).(HlsPtrOutput)
}

func (o HlsOutput) FragmentsPerTsSegment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Hls) *int { return v.FragmentsPerTsSegment }).(pulumi.IntPtrOutput)
}

type HlsPtrOutput struct{ *pulumi.OutputState }

func (HlsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hls)(nil)).Elem()
}

func (o HlsPtrOutput) ToHlsPtrOutput() HlsPtrOutput {
	return o
}

func (o HlsPtrOutput) ToHlsPtrOutputWithContext(ctx context.Context) HlsPtrOutput {
	return o
}

func (o HlsPtrOutput) Elem() HlsOutput {
	return o.ApplyT(func(v *Hls) Hls {
		if v != nil {
			return *v
		}
		var ret Hls
		return ret
	}).(HlsOutput)
}

func (o HlsPtrOutput) FragmentsPerTsSegment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Hls) *int {
		if v == nil {
			return nil
		}
		return v.FragmentsPerTsSegment
	}).(pulumi.IntPtrOutput)
}

type HlsResponse struct {
	FragmentsPerTsSegment *int `pulumi:"fragmentsPerTsSegment"`
}





type HlsResponseInput interface {
	pulumi.Input

	ToHlsResponseOutput() HlsResponseOutput
	ToHlsResponseOutputWithContext(context.Context) HlsResponseOutput
}

type HlsResponseArgs struct {
	FragmentsPerTsSegment pulumi.IntPtrInput `pulumi:"fragmentsPerTsSegment"`
}

func (HlsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HlsResponse)(nil)).Elem()
}

func (i HlsResponseArgs) ToHlsResponseOutput() HlsResponseOutput {
	return i.ToHlsResponseOutputWithContext(context.Background())
}

func (i HlsResponseArgs) ToHlsResponseOutputWithContext(ctx context.Context) HlsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsResponseOutput)
}

func (i HlsResponseArgs) ToHlsResponsePtrOutput() HlsResponsePtrOutput {
	return i.ToHlsResponsePtrOutputWithContext(context.Background())
}

func (i HlsResponseArgs) ToHlsResponsePtrOutputWithContext(ctx context.Context) HlsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsResponseOutput).ToHlsResponsePtrOutputWithContext(ctx)
}









type HlsResponsePtrInput interface {
	pulumi.Input

	ToHlsResponsePtrOutput() HlsResponsePtrOutput
	ToHlsResponsePtrOutputWithContext(context.Context) HlsResponsePtrOutput
}

type hlsResponsePtrType HlsResponseArgs

func HlsResponsePtr(v *HlsResponseArgs) HlsResponsePtrInput {
	return (*hlsResponsePtrType)(v)
}

func (*hlsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HlsResponse)(nil)).Elem()
}

func (i *hlsResponsePtrType) ToHlsResponsePtrOutput() HlsResponsePtrOutput {
	return i.ToHlsResponsePtrOutputWithContext(context.Background())
}

func (i *hlsResponsePtrType) ToHlsResponsePtrOutputWithContext(ctx context.Context) HlsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsResponsePtrOutput)
}

type HlsResponseOutput struct{ *pulumi.OutputState }

func (HlsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HlsResponse)(nil)).Elem()
}

func (o HlsResponseOutput) ToHlsResponseOutput() HlsResponseOutput {
	return o
}

func (o HlsResponseOutput) ToHlsResponseOutputWithContext(ctx context.Context) HlsResponseOutput {
	return o
}

func (o HlsResponseOutput) ToHlsResponsePtrOutput() HlsResponsePtrOutput {
	return o.ToHlsResponsePtrOutputWithContext(context.Background())
}

func (o HlsResponseOutput) ToHlsResponsePtrOutputWithContext(ctx context.Context) HlsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HlsResponse) *HlsResponse {
		return &v
	}).(HlsResponsePtrOutput)
}

func (o HlsResponseOutput) FragmentsPerTsSegment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HlsResponse) *int { return v.FragmentsPerTsSegment }).(pulumi.IntPtrOutput)
}

type HlsResponsePtrOutput struct{ *pulumi.OutputState }

func (HlsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HlsResponse)(nil)).Elem()
}

func (o HlsResponsePtrOutput) ToHlsResponsePtrOutput() HlsResponsePtrOutput {
	return o
}

func (o HlsResponsePtrOutput) ToHlsResponsePtrOutputWithContext(ctx context.Context) HlsResponsePtrOutput {
	return o
}

func (o HlsResponsePtrOutput) Elem() HlsResponseOutput {
	return o.ApplyT(func(v *HlsResponse) HlsResponse {
		if v != nil {
			return *v
		}
		var ret HlsResponse
		return ret
	}).(HlsResponseOutput)
}

func (o HlsResponsePtrOutput) FragmentsPerTsSegment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HlsResponse) *int {
		if v == nil {
			return nil
		}
		return v.FragmentsPerTsSegment
	}).(pulumi.IntPtrOutput)
}

type IPAccessControl struct {
	Allow []IPRange `pulumi:"allow"`
}





type IPAccessControlInput interface {
	pulumi.Input

	ToIPAccessControlOutput() IPAccessControlOutput
	ToIPAccessControlOutputWithContext(context.Context) IPAccessControlOutput
}

type IPAccessControlArgs struct {
	Allow IPRangeArrayInput `pulumi:"allow"`
}

func (IPAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAccessControl)(nil)).Elem()
}

func (i IPAccessControlArgs) ToIPAccessControlOutput() IPAccessControlOutput {
	return i.ToIPAccessControlOutputWithContext(context.Background())
}

func (i IPAccessControlArgs) ToIPAccessControlOutputWithContext(ctx context.Context) IPAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlOutput)
}

func (i IPAccessControlArgs) ToIPAccessControlPtrOutput() IPAccessControlPtrOutput {
	return i.ToIPAccessControlPtrOutputWithContext(context.Background())
}

func (i IPAccessControlArgs) ToIPAccessControlPtrOutputWithContext(ctx context.Context) IPAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlOutput).ToIPAccessControlPtrOutputWithContext(ctx)
}









type IPAccessControlPtrInput interface {
	pulumi.Input

	ToIPAccessControlPtrOutput() IPAccessControlPtrOutput
	ToIPAccessControlPtrOutputWithContext(context.Context) IPAccessControlPtrOutput
}

type ipaccessControlPtrType IPAccessControlArgs

func IPAccessControlPtr(v *IPAccessControlArgs) IPAccessControlPtrInput {
	return (*ipaccessControlPtrType)(v)
}

func (*ipaccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAccessControl)(nil)).Elem()
}

func (i *ipaccessControlPtrType) ToIPAccessControlPtrOutput() IPAccessControlPtrOutput {
	return i.ToIPAccessControlPtrOutputWithContext(context.Background())
}

func (i *ipaccessControlPtrType) ToIPAccessControlPtrOutputWithContext(ctx context.Context) IPAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlPtrOutput)
}

type IPAccessControlOutput struct{ *pulumi.OutputState }

func (IPAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAccessControl)(nil)).Elem()
}

func (o IPAccessControlOutput) ToIPAccessControlOutput() IPAccessControlOutput {
	return o
}

func (o IPAccessControlOutput) ToIPAccessControlOutputWithContext(ctx context.Context) IPAccessControlOutput {
	return o
}

func (o IPAccessControlOutput) ToIPAccessControlPtrOutput() IPAccessControlPtrOutput {
	return o.ToIPAccessControlPtrOutputWithContext(context.Background())
}

func (o IPAccessControlOutput) ToIPAccessControlPtrOutputWithContext(ctx context.Context) IPAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPAccessControl) *IPAccessControl {
		return &v
	}).(IPAccessControlPtrOutput)
}

func (o IPAccessControlOutput) Allow() IPRangeArrayOutput {
	return o.ApplyT(func(v IPAccessControl) []IPRange { return v.Allow }).(IPRangeArrayOutput)
}

type IPAccessControlPtrOutput struct{ *pulumi.OutputState }

func (IPAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAccessControl)(nil)).Elem()
}

func (o IPAccessControlPtrOutput) ToIPAccessControlPtrOutput() IPAccessControlPtrOutput {
	return o
}

func (o IPAccessControlPtrOutput) ToIPAccessControlPtrOutputWithContext(ctx context.Context) IPAccessControlPtrOutput {
	return o
}

func (o IPAccessControlPtrOutput) Elem() IPAccessControlOutput {
	return o.ApplyT(func(v *IPAccessControl) IPAccessControl {
		if v != nil {
			return *v
		}
		var ret IPAccessControl
		return ret
	}).(IPAccessControlOutput)
}

func (o IPAccessControlPtrOutput) Allow() IPRangeArrayOutput {
	return o.ApplyT(func(v *IPAccessControl) []IPRange {
		if v == nil {
			return nil
		}
		return v.Allow
	}).(IPRangeArrayOutput)
}

type IPAccessControlResponse struct {
	Allow []IPRangeResponse `pulumi:"allow"`
}





type IPAccessControlResponseInput interface {
	pulumi.Input

	ToIPAccessControlResponseOutput() IPAccessControlResponseOutput
	ToIPAccessControlResponseOutputWithContext(context.Context) IPAccessControlResponseOutput
}

type IPAccessControlResponseArgs struct {
	Allow IPRangeResponseArrayInput `pulumi:"allow"`
}

func (IPAccessControlResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAccessControlResponse)(nil)).Elem()
}

func (i IPAccessControlResponseArgs) ToIPAccessControlResponseOutput() IPAccessControlResponseOutput {
	return i.ToIPAccessControlResponseOutputWithContext(context.Background())
}

func (i IPAccessControlResponseArgs) ToIPAccessControlResponseOutputWithContext(ctx context.Context) IPAccessControlResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlResponseOutput)
}

func (i IPAccessControlResponseArgs) ToIPAccessControlResponsePtrOutput() IPAccessControlResponsePtrOutput {
	return i.ToIPAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i IPAccessControlResponseArgs) ToIPAccessControlResponsePtrOutputWithContext(ctx context.Context) IPAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlResponseOutput).ToIPAccessControlResponsePtrOutputWithContext(ctx)
}









type IPAccessControlResponsePtrInput interface {
	pulumi.Input

	ToIPAccessControlResponsePtrOutput() IPAccessControlResponsePtrOutput
	ToIPAccessControlResponsePtrOutputWithContext(context.Context) IPAccessControlResponsePtrOutput
}

type ipaccessControlResponsePtrType IPAccessControlResponseArgs

func IPAccessControlResponsePtr(v *IPAccessControlResponseArgs) IPAccessControlResponsePtrInput {
	return (*ipaccessControlResponsePtrType)(v)
}

func (*ipaccessControlResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAccessControlResponse)(nil)).Elem()
}

func (i *ipaccessControlResponsePtrType) ToIPAccessControlResponsePtrOutput() IPAccessControlResponsePtrOutput {
	return i.ToIPAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i *ipaccessControlResponsePtrType) ToIPAccessControlResponsePtrOutputWithContext(ctx context.Context) IPAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlResponsePtrOutput)
}

type IPAccessControlResponseOutput struct{ *pulumi.OutputState }

func (IPAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAccessControlResponse)(nil)).Elem()
}

func (o IPAccessControlResponseOutput) ToIPAccessControlResponseOutput() IPAccessControlResponseOutput {
	return o
}

func (o IPAccessControlResponseOutput) ToIPAccessControlResponseOutputWithContext(ctx context.Context) IPAccessControlResponseOutput {
	return o
}

func (o IPAccessControlResponseOutput) ToIPAccessControlResponsePtrOutput() IPAccessControlResponsePtrOutput {
	return o.ToIPAccessControlResponsePtrOutputWithContext(context.Background())
}

func (o IPAccessControlResponseOutput) ToIPAccessControlResponsePtrOutputWithContext(ctx context.Context) IPAccessControlResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPAccessControlResponse) *IPAccessControlResponse {
		return &v
	}).(IPAccessControlResponsePtrOutput)
}

func (o IPAccessControlResponseOutput) Allow() IPRangeResponseArrayOutput {
	return o.ApplyT(func(v IPAccessControlResponse) []IPRangeResponse { return v.Allow }).(IPRangeResponseArrayOutput)
}

type IPAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (IPAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAccessControlResponse)(nil)).Elem()
}

func (o IPAccessControlResponsePtrOutput) ToIPAccessControlResponsePtrOutput() IPAccessControlResponsePtrOutput {
	return o
}

func (o IPAccessControlResponsePtrOutput) ToIPAccessControlResponsePtrOutputWithContext(ctx context.Context) IPAccessControlResponsePtrOutput {
	return o
}

func (o IPAccessControlResponsePtrOutput) Elem() IPAccessControlResponseOutput {
	return o.ApplyT(func(v *IPAccessControlResponse) IPAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret IPAccessControlResponse
		return ret
	}).(IPAccessControlResponseOutput)
}

func (o IPAccessControlResponsePtrOutput) Allow() IPRangeResponseArrayOutput {
	return o.ApplyT(func(v *IPAccessControlResponse) []IPRangeResponse {
		if v == nil {
			return nil
		}
		return v.Allow
	}).(IPRangeResponseArrayOutput)
}

type IPRange struct {
	Address            *string `pulumi:"address"`
	Name               *string `pulumi:"name"`
	SubnetPrefixLength *int    `pulumi:"subnetPrefixLength"`
}





type IPRangeInput interface {
	pulumi.Input

	ToIPRangeOutput() IPRangeOutput
	ToIPRangeOutputWithContext(context.Context) IPRangeOutput
}

type IPRangeArgs struct {
	Address            pulumi.StringPtrInput `pulumi:"address"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	SubnetPrefixLength pulumi.IntPtrInput    `pulumi:"subnetPrefixLength"`
}

func (IPRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRange)(nil)).Elem()
}

func (i IPRangeArgs) ToIPRangeOutput() IPRangeOutput {
	return i.ToIPRangeOutputWithContext(context.Background())
}

func (i IPRangeArgs) ToIPRangeOutputWithContext(ctx context.Context) IPRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRangeOutput)
}





type IPRangeArrayInput interface {
	pulumi.Input

	ToIPRangeArrayOutput() IPRangeArrayOutput
	ToIPRangeArrayOutputWithContext(context.Context) IPRangeArrayOutput
}

type IPRangeArray []IPRangeInput

func (IPRangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRange)(nil)).Elem()
}

func (i IPRangeArray) ToIPRangeArrayOutput() IPRangeArrayOutput {
	return i.ToIPRangeArrayOutputWithContext(context.Background())
}

func (i IPRangeArray) ToIPRangeArrayOutputWithContext(ctx context.Context) IPRangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRangeArrayOutput)
}

type IPRangeOutput struct{ *pulumi.OutputState }

func (IPRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRange)(nil)).Elem()
}

func (o IPRangeOutput) ToIPRangeOutput() IPRangeOutput {
	return o
}

func (o IPRangeOutput) ToIPRangeOutputWithContext(ctx context.Context) IPRangeOutput {
	return o
}

func (o IPRangeOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRange) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o IPRangeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRange) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IPRangeOutput) SubnetPrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IPRange) *int { return v.SubnetPrefixLength }).(pulumi.IntPtrOutput)
}

type IPRangeArrayOutput struct{ *pulumi.OutputState }

func (IPRangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRange)(nil)).Elem()
}

func (o IPRangeArrayOutput) ToIPRangeArrayOutput() IPRangeArrayOutput {
	return o
}

func (o IPRangeArrayOutput) ToIPRangeArrayOutputWithContext(ctx context.Context) IPRangeArrayOutput {
	return o
}

func (o IPRangeArrayOutput) Index(i pulumi.IntInput) IPRangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRange {
		return vs[0].([]IPRange)[vs[1].(int)]
	}).(IPRangeOutput)
}

type IPRangeResponse struct {
	Address            *string `pulumi:"address"`
	Name               *string `pulumi:"name"`
	SubnetPrefixLength *int    `pulumi:"subnetPrefixLength"`
}





type IPRangeResponseInput interface {
	pulumi.Input

	ToIPRangeResponseOutput() IPRangeResponseOutput
	ToIPRangeResponseOutputWithContext(context.Context) IPRangeResponseOutput
}

type IPRangeResponseArgs struct {
	Address            pulumi.StringPtrInput `pulumi:"address"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	SubnetPrefixLength pulumi.IntPtrInput    `pulumi:"subnetPrefixLength"`
}

func (IPRangeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRangeResponse)(nil)).Elem()
}

func (i IPRangeResponseArgs) ToIPRangeResponseOutput() IPRangeResponseOutput {
	return i.ToIPRangeResponseOutputWithContext(context.Background())
}

func (i IPRangeResponseArgs) ToIPRangeResponseOutputWithContext(ctx context.Context) IPRangeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRangeResponseOutput)
}





type IPRangeResponseArrayInput interface {
	pulumi.Input

	ToIPRangeResponseArrayOutput() IPRangeResponseArrayOutput
	ToIPRangeResponseArrayOutputWithContext(context.Context) IPRangeResponseArrayOutput
}

type IPRangeResponseArray []IPRangeResponseInput

func (IPRangeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRangeResponse)(nil)).Elem()
}

func (i IPRangeResponseArray) ToIPRangeResponseArrayOutput() IPRangeResponseArrayOutput {
	return i.ToIPRangeResponseArrayOutputWithContext(context.Background())
}

func (i IPRangeResponseArray) ToIPRangeResponseArrayOutputWithContext(ctx context.Context) IPRangeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRangeResponseArrayOutput)
}

type IPRangeResponseOutput struct{ *pulumi.OutputState }

func (IPRangeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRangeResponse)(nil)).Elem()
}

func (o IPRangeResponseOutput) ToIPRangeResponseOutput() IPRangeResponseOutput {
	return o
}

func (o IPRangeResponseOutput) ToIPRangeResponseOutputWithContext(ctx context.Context) IPRangeResponseOutput {
	return o
}

func (o IPRangeResponseOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRangeResponse) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o IPRangeResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRangeResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IPRangeResponseOutput) SubnetPrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IPRangeResponse) *int { return v.SubnetPrefixLength }).(pulumi.IntPtrOutput)
}

type IPRangeResponseArrayOutput struct{ *pulumi.OutputState }

func (IPRangeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRangeResponse)(nil)).Elem()
}

func (o IPRangeResponseArrayOutput) ToIPRangeResponseArrayOutput() IPRangeResponseArrayOutput {
	return o
}

func (o IPRangeResponseArrayOutput) ToIPRangeResponseArrayOutputWithContext(ctx context.Context) IPRangeResponseArrayOutput {
	return o
}

func (o IPRangeResponseArrayOutput) Index(i pulumi.IntInput) IPRangeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRangeResponse {
		return vs[0].([]IPRangeResponse)[vs[1].(int)]
	}).(IPRangeResponseOutput)
}

type LiveEventEncoding struct {
	EncodingType *string `pulumi:"encodingType"`
	PresetName   *string `pulumi:"presetName"`
}





type LiveEventEncodingInput interface {
	pulumi.Input

	ToLiveEventEncodingOutput() LiveEventEncodingOutput
	ToLiveEventEncodingOutputWithContext(context.Context) LiveEventEncodingOutput
}

type LiveEventEncodingArgs struct {
	EncodingType pulumi.StringPtrInput `pulumi:"encodingType"`
	PresetName   pulumi.StringPtrInput `pulumi:"presetName"`
}

func (LiveEventEncodingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncoding)(nil)).Elem()
}

func (i LiveEventEncodingArgs) ToLiveEventEncodingOutput() LiveEventEncodingOutput {
	return i.ToLiveEventEncodingOutputWithContext(context.Background())
}

func (i LiveEventEncodingArgs) ToLiveEventEncodingOutputWithContext(ctx context.Context) LiveEventEncodingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingOutput)
}

func (i LiveEventEncodingArgs) ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput {
	return i.ToLiveEventEncodingPtrOutputWithContext(context.Background())
}

func (i LiveEventEncodingArgs) ToLiveEventEncodingPtrOutputWithContext(ctx context.Context) LiveEventEncodingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingOutput).ToLiveEventEncodingPtrOutputWithContext(ctx)
}









type LiveEventEncodingPtrInput interface {
	pulumi.Input

	ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput
	ToLiveEventEncodingPtrOutputWithContext(context.Context) LiveEventEncodingPtrOutput
}

type liveEventEncodingPtrType LiveEventEncodingArgs

func LiveEventEncodingPtr(v *LiveEventEncodingArgs) LiveEventEncodingPtrInput {
	return (*liveEventEncodingPtrType)(v)
}

func (*liveEventEncodingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncoding)(nil)).Elem()
}

func (i *liveEventEncodingPtrType) ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput {
	return i.ToLiveEventEncodingPtrOutputWithContext(context.Background())
}

func (i *liveEventEncodingPtrType) ToLiveEventEncodingPtrOutputWithContext(ctx context.Context) LiveEventEncodingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingPtrOutput)
}

type LiveEventEncodingOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncoding)(nil)).Elem()
}

func (o LiveEventEncodingOutput) ToLiveEventEncodingOutput() LiveEventEncodingOutput {
	return o
}

func (o LiveEventEncodingOutput) ToLiveEventEncodingOutputWithContext(ctx context.Context) LiveEventEncodingOutput {
	return o
}

func (o LiveEventEncodingOutput) ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput {
	return o.ToLiveEventEncodingPtrOutputWithContext(context.Background())
}

func (o LiveEventEncodingOutput) ToLiveEventEncodingPtrOutputWithContext(ctx context.Context) LiveEventEncodingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventEncoding) *LiveEventEncoding {
		return &v
	}).(LiveEventEncodingPtrOutput)
}

func (o LiveEventEncodingOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncoding) *string { return v.EncodingType }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncoding) *string { return v.PresetName }).(pulumi.StringPtrOutput)
}

type LiveEventEncodingPtrOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncoding)(nil)).Elem()
}

func (o LiveEventEncodingPtrOutput) ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput {
	return o
}

func (o LiveEventEncodingPtrOutput) ToLiveEventEncodingPtrOutputWithContext(ctx context.Context) LiveEventEncodingPtrOutput {
	return o
}

func (o LiveEventEncodingPtrOutput) Elem() LiveEventEncodingOutput {
	return o.ApplyT(func(v *LiveEventEncoding) LiveEventEncoding {
		if v != nil {
			return *v
		}
		var ret LiveEventEncoding
		return ret
	}).(LiveEventEncodingOutput)
}

func (o LiveEventEncodingPtrOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncoding) *string {
		if v == nil {
			return nil
		}
		return v.EncodingType
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingPtrOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncoding) *string {
		if v == nil {
			return nil
		}
		return v.PresetName
	}).(pulumi.StringPtrOutput)
}

type LiveEventEncodingResponse struct {
	EncodingType *string `pulumi:"encodingType"`
	PresetName   *string `pulumi:"presetName"`
}





type LiveEventEncodingResponseInput interface {
	pulumi.Input

	ToLiveEventEncodingResponseOutput() LiveEventEncodingResponseOutput
	ToLiveEventEncodingResponseOutputWithContext(context.Context) LiveEventEncodingResponseOutput
}

type LiveEventEncodingResponseArgs struct {
	EncodingType pulumi.StringPtrInput `pulumi:"encodingType"`
	PresetName   pulumi.StringPtrInput `pulumi:"presetName"`
}

func (LiveEventEncodingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncodingResponse)(nil)).Elem()
}

func (i LiveEventEncodingResponseArgs) ToLiveEventEncodingResponseOutput() LiveEventEncodingResponseOutput {
	return i.ToLiveEventEncodingResponseOutputWithContext(context.Background())
}

func (i LiveEventEncodingResponseArgs) ToLiveEventEncodingResponseOutputWithContext(ctx context.Context) LiveEventEncodingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingResponseOutput)
}

func (i LiveEventEncodingResponseArgs) ToLiveEventEncodingResponsePtrOutput() LiveEventEncodingResponsePtrOutput {
	return i.ToLiveEventEncodingResponsePtrOutputWithContext(context.Background())
}

func (i LiveEventEncodingResponseArgs) ToLiveEventEncodingResponsePtrOutputWithContext(ctx context.Context) LiveEventEncodingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingResponseOutput).ToLiveEventEncodingResponsePtrOutputWithContext(ctx)
}









type LiveEventEncodingResponsePtrInput interface {
	pulumi.Input

	ToLiveEventEncodingResponsePtrOutput() LiveEventEncodingResponsePtrOutput
	ToLiveEventEncodingResponsePtrOutputWithContext(context.Context) LiveEventEncodingResponsePtrOutput
}

type liveEventEncodingResponsePtrType LiveEventEncodingResponseArgs

func LiveEventEncodingResponsePtr(v *LiveEventEncodingResponseArgs) LiveEventEncodingResponsePtrInput {
	return (*liveEventEncodingResponsePtrType)(v)
}

func (*liveEventEncodingResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncodingResponse)(nil)).Elem()
}

func (i *liveEventEncodingResponsePtrType) ToLiveEventEncodingResponsePtrOutput() LiveEventEncodingResponsePtrOutput {
	return i.ToLiveEventEncodingResponsePtrOutputWithContext(context.Background())
}

func (i *liveEventEncodingResponsePtrType) ToLiveEventEncodingResponsePtrOutputWithContext(ctx context.Context) LiveEventEncodingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingResponsePtrOutput)
}

type LiveEventEncodingResponseOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncodingResponse)(nil)).Elem()
}

func (o LiveEventEncodingResponseOutput) ToLiveEventEncodingResponseOutput() LiveEventEncodingResponseOutput {
	return o
}

func (o LiveEventEncodingResponseOutput) ToLiveEventEncodingResponseOutputWithContext(ctx context.Context) LiveEventEncodingResponseOutput {
	return o
}

func (o LiveEventEncodingResponseOutput) ToLiveEventEncodingResponsePtrOutput() LiveEventEncodingResponsePtrOutput {
	return o.ToLiveEventEncodingResponsePtrOutputWithContext(context.Background())
}

func (o LiveEventEncodingResponseOutput) ToLiveEventEncodingResponsePtrOutputWithContext(ctx context.Context) LiveEventEncodingResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventEncodingResponse) *LiveEventEncodingResponse {
		return &v
	}).(LiveEventEncodingResponsePtrOutput)
}

func (o LiveEventEncodingResponseOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncodingResponse) *string { return v.EncodingType }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponseOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncodingResponse) *string { return v.PresetName }).(pulumi.StringPtrOutput)
}

type LiveEventEncodingResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncodingResponse)(nil)).Elem()
}

func (o LiveEventEncodingResponsePtrOutput) ToLiveEventEncodingResponsePtrOutput() LiveEventEncodingResponsePtrOutput {
	return o
}

func (o LiveEventEncodingResponsePtrOutput) ToLiveEventEncodingResponsePtrOutputWithContext(ctx context.Context) LiveEventEncodingResponsePtrOutput {
	return o
}

func (o LiveEventEncodingResponsePtrOutput) Elem() LiveEventEncodingResponseOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) LiveEventEncodingResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventEncodingResponse
		return ret
	}).(LiveEventEncodingResponseOutput)
}

func (o LiveEventEncodingResponsePtrOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) *string {
		if v == nil {
			return nil
		}
		return v.EncodingType
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponsePtrOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) *string {
		if v == nil {
			return nil
		}
		return v.PresetName
	}).(pulumi.StringPtrOutput)
}

type LiveEventEndpoint struct {
	Protocol *string `pulumi:"protocol"`
	Url      *string `pulumi:"url"`
}





type LiveEventEndpointInput interface {
	pulumi.Input

	ToLiveEventEndpointOutput() LiveEventEndpointOutput
	ToLiveEventEndpointOutputWithContext(context.Context) LiveEventEndpointOutput
}

type LiveEventEndpointArgs struct {
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	Url      pulumi.StringPtrInput `pulumi:"url"`
}

func (LiveEventEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEndpoint)(nil)).Elem()
}

func (i LiveEventEndpointArgs) ToLiveEventEndpointOutput() LiveEventEndpointOutput {
	return i.ToLiveEventEndpointOutputWithContext(context.Background())
}

func (i LiveEventEndpointArgs) ToLiveEventEndpointOutputWithContext(ctx context.Context) LiveEventEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEndpointOutput)
}





type LiveEventEndpointArrayInput interface {
	pulumi.Input

	ToLiveEventEndpointArrayOutput() LiveEventEndpointArrayOutput
	ToLiveEventEndpointArrayOutputWithContext(context.Context) LiveEventEndpointArrayOutput
}

type LiveEventEndpointArray []LiveEventEndpointInput

func (LiveEventEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventEndpoint)(nil)).Elem()
}

func (i LiveEventEndpointArray) ToLiveEventEndpointArrayOutput() LiveEventEndpointArrayOutput {
	return i.ToLiveEventEndpointArrayOutputWithContext(context.Background())
}

func (i LiveEventEndpointArray) ToLiveEventEndpointArrayOutputWithContext(ctx context.Context) LiveEventEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEndpointArrayOutput)
}

type LiveEventEndpointOutput struct{ *pulumi.OutputState }

func (LiveEventEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEndpoint)(nil)).Elem()
}

func (o LiveEventEndpointOutput) ToLiveEventEndpointOutput() LiveEventEndpointOutput {
	return o
}

func (o LiveEventEndpointOutput) ToLiveEventEndpointOutputWithContext(ctx context.Context) LiveEventEndpointOutput {
	return o
}

func (o LiveEventEndpointOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEndpoint) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o LiveEventEndpointOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEndpoint) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type LiveEventEndpointArrayOutput struct{ *pulumi.OutputState }

func (LiveEventEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventEndpoint)(nil)).Elem()
}

func (o LiveEventEndpointArrayOutput) ToLiveEventEndpointArrayOutput() LiveEventEndpointArrayOutput {
	return o
}

func (o LiveEventEndpointArrayOutput) ToLiveEventEndpointArrayOutputWithContext(ctx context.Context) LiveEventEndpointArrayOutput {
	return o
}

func (o LiveEventEndpointArrayOutput) Index(i pulumi.IntInput) LiveEventEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventEndpoint {
		return vs[0].([]LiveEventEndpoint)[vs[1].(int)]
	}).(LiveEventEndpointOutput)
}

type LiveEventEndpointResponse struct {
	Protocol *string `pulumi:"protocol"`
	Url      *string `pulumi:"url"`
}





type LiveEventEndpointResponseInput interface {
	pulumi.Input

	ToLiveEventEndpointResponseOutput() LiveEventEndpointResponseOutput
	ToLiveEventEndpointResponseOutputWithContext(context.Context) LiveEventEndpointResponseOutput
}

type LiveEventEndpointResponseArgs struct {
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	Url      pulumi.StringPtrInput `pulumi:"url"`
}

func (LiveEventEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEndpointResponse)(nil)).Elem()
}

func (i LiveEventEndpointResponseArgs) ToLiveEventEndpointResponseOutput() LiveEventEndpointResponseOutput {
	return i.ToLiveEventEndpointResponseOutputWithContext(context.Background())
}

func (i LiveEventEndpointResponseArgs) ToLiveEventEndpointResponseOutputWithContext(ctx context.Context) LiveEventEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEndpointResponseOutput)
}





type LiveEventEndpointResponseArrayInput interface {
	pulumi.Input

	ToLiveEventEndpointResponseArrayOutput() LiveEventEndpointResponseArrayOutput
	ToLiveEventEndpointResponseArrayOutputWithContext(context.Context) LiveEventEndpointResponseArrayOutput
}

type LiveEventEndpointResponseArray []LiveEventEndpointResponseInput

func (LiveEventEndpointResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventEndpointResponse)(nil)).Elem()
}

func (i LiveEventEndpointResponseArray) ToLiveEventEndpointResponseArrayOutput() LiveEventEndpointResponseArrayOutput {
	return i.ToLiveEventEndpointResponseArrayOutputWithContext(context.Background())
}

func (i LiveEventEndpointResponseArray) ToLiveEventEndpointResponseArrayOutputWithContext(ctx context.Context) LiveEventEndpointResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEndpointResponseArrayOutput)
}

type LiveEventEndpointResponseOutput struct{ *pulumi.OutputState }

func (LiveEventEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEndpointResponse)(nil)).Elem()
}

func (o LiveEventEndpointResponseOutput) ToLiveEventEndpointResponseOutput() LiveEventEndpointResponseOutput {
	return o
}

func (o LiveEventEndpointResponseOutput) ToLiveEventEndpointResponseOutputWithContext(ctx context.Context) LiveEventEndpointResponseOutput {
	return o
}

func (o LiveEventEndpointResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEndpointResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o LiveEventEndpointResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEndpointResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type LiveEventEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (LiveEventEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventEndpointResponse)(nil)).Elem()
}

func (o LiveEventEndpointResponseArrayOutput) ToLiveEventEndpointResponseArrayOutput() LiveEventEndpointResponseArrayOutput {
	return o
}

func (o LiveEventEndpointResponseArrayOutput) ToLiveEventEndpointResponseArrayOutputWithContext(ctx context.Context) LiveEventEndpointResponseArrayOutput {
	return o
}

func (o LiveEventEndpointResponseArrayOutput) Index(i pulumi.IntInput) LiveEventEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventEndpointResponse {
		return vs[0].([]LiveEventEndpointResponse)[vs[1].(int)]
	}).(LiveEventEndpointResponseOutput)
}

type LiveEventInputType struct {
	AccessControl            *LiveEventInputAccessControl `pulumi:"accessControl"`
	AccessToken              *string                      `pulumi:"accessToken"`
	Endpoints                []LiveEventEndpoint          `pulumi:"endpoints"`
	KeyFrameIntervalDuration *string                      `pulumi:"keyFrameIntervalDuration"`
	StreamingProtocol        string                       `pulumi:"streamingProtocol"`
}





type LiveEventInputTypeInput interface {
	pulumi.Input

	ToLiveEventInputTypeOutput() LiveEventInputTypeOutput
	ToLiveEventInputTypeOutputWithContext(context.Context) LiveEventInputTypeOutput
}

type LiveEventInputTypeArgs struct {
	AccessControl            LiveEventInputAccessControlPtrInput `pulumi:"accessControl"`
	AccessToken              pulumi.StringPtrInput               `pulumi:"accessToken"`
	Endpoints                LiveEventEndpointArrayInput         `pulumi:"endpoints"`
	KeyFrameIntervalDuration pulumi.StringPtrInput               `pulumi:"keyFrameIntervalDuration"`
	StreamingProtocol        pulumi.StringInput                  `pulumi:"streamingProtocol"`
}

func (LiveEventInputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputType)(nil)).Elem()
}

func (i LiveEventInputTypeArgs) ToLiveEventInputTypeOutput() LiveEventInputTypeOutput {
	return i.ToLiveEventInputTypeOutputWithContext(context.Background())
}

func (i LiveEventInputTypeArgs) ToLiveEventInputTypeOutputWithContext(ctx context.Context) LiveEventInputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTypeOutput)
}

func (i LiveEventInputTypeArgs) ToLiveEventInputTypePtrOutput() LiveEventInputTypePtrOutput {
	return i.ToLiveEventInputTypePtrOutputWithContext(context.Background())
}

func (i LiveEventInputTypeArgs) ToLiveEventInputTypePtrOutputWithContext(ctx context.Context) LiveEventInputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTypeOutput).ToLiveEventInputTypePtrOutputWithContext(ctx)
}









type LiveEventInputTypePtrInput interface {
	pulumi.Input

	ToLiveEventInputTypePtrOutput() LiveEventInputTypePtrOutput
	ToLiveEventInputTypePtrOutputWithContext(context.Context) LiveEventInputTypePtrOutput
}

type liveEventInputTypePtrType LiveEventInputTypeArgs

func LiveEventInputTypePtr(v *LiveEventInputTypeArgs) LiveEventInputTypePtrInput {
	return (*liveEventInputTypePtrType)(v)
}

func (*liveEventInputTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputType)(nil)).Elem()
}

func (i *liveEventInputTypePtrType) ToLiveEventInputTypePtrOutput() LiveEventInputTypePtrOutput {
	return i.ToLiveEventInputTypePtrOutputWithContext(context.Background())
}

func (i *liveEventInputTypePtrType) ToLiveEventInputTypePtrOutputWithContext(ctx context.Context) LiveEventInputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTypePtrOutput)
}

type LiveEventInputTypeOutput struct{ *pulumi.OutputState }

func (LiveEventInputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputType)(nil)).Elem()
}

func (o LiveEventInputTypeOutput) ToLiveEventInputTypeOutput() LiveEventInputTypeOutput {
	return o
}

func (o LiveEventInputTypeOutput) ToLiveEventInputTypeOutputWithContext(ctx context.Context) LiveEventInputTypeOutput {
	return o
}

func (o LiveEventInputTypeOutput) ToLiveEventInputTypePtrOutput() LiveEventInputTypePtrOutput {
	return o.ToLiveEventInputTypePtrOutputWithContext(context.Background())
}

func (o LiveEventInputTypeOutput) ToLiveEventInputTypePtrOutputWithContext(ctx context.Context) LiveEventInputTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventInputType) *LiveEventInputType {
		return &v
	}).(LiveEventInputTypePtrOutput)
}

func (o LiveEventInputTypeOutput) AccessControl() LiveEventInputAccessControlPtrOutput {
	return o.ApplyT(func(v LiveEventInputType) *LiveEventInputAccessControl { return v.AccessControl }).(LiveEventInputAccessControlPtrOutput)
}

func (o LiveEventInputTypeOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputType) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTypeOutput) Endpoints() LiveEventEndpointArrayOutput {
	return o.ApplyT(func(v LiveEventInputType) []LiveEventEndpoint { return v.Endpoints }).(LiveEventEndpointArrayOutput)
}

func (o LiveEventInputTypeOutput) KeyFrameIntervalDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputType) *string { return v.KeyFrameIntervalDuration }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTypeOutput) StreamingProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LiveEventInputType) string { return v.StreamingProtocol }).(pulumi.StringOutput)
}

type LiveEventInputTypePtrOutput struct{ *pulumi.OutputState }

func (LiveEventInputTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputType)(nil)).Elem()
}

func (o LiveEventInputTypePtrOutput) ToLiveEventInputTypePtrOutput() LiveEventInputTypePtrOutput {
	return o
}

func (o LiveEventInputTypePtrOutput) ToLiveEventInputTypePtrOutputWithContext(ctx context.Context) LiveEventInputTypePtrOutput {
	return o
}

func (o LiveEventInputTypePtrOutput) Elem() LiveEventInputTypeOutput {
	return o.ApplyT(func(v *LiveEventInputType) LiveEventInputType {
		if v != nil {
			return *v
		}
		var ret LiveEventInputType
		return ret
	}).(LiveEventInputTypeOutput)
}

func (o LiveEventInputTypePtrOutput) AccessControl() LiveEventInputAccessControlPtrOutput {
	return o.ApplyT(func(v *LiveEventInputType) *LiveEventInputAccessControl {
		if v == nil {
			return nil
		}
		return v.AccessControl
	}).(LiveEventInputAccessControlPtrOutput)
}

func (o LiveEventInputTypePtrOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventInputType) *string {
		if v == nil {
			return nil
		}
		return v.AccessToken
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTypePtrOutput) Endpoints() LiveEventEndpointArrayOutput {
	return o.ApplyT(func(v *LiveEventInputType) []LiveEventEndpoint {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(LiveEventEndpointArrayOutput)
}

func (o LiveEventInputTypePtrOutput) KeyFrameIntervalDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventInputType) *string {
		if v == nil {
			return nil
		}
		return v.KeyFrameIntervalDuration
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTypePtrOutput) StreamingProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventInputType) *string {
		if v == nil {
			return nil
		}
		return &v.StreamingProtocol
	}).(pulumi.StringPtrOutput)
}

type LiveEventInputAccessControl struct {
	Ip *IPAccessControl `pulumi:"ip"`
}





type LiveEventInputAccessControlInput interface {
	pulumi.Input

	ToLiveEventInputAccessControlOutput() LiveEventInputAccessControlOutput
	ToLiveEventInputAccessControlOutputWithContext(context.Context) LiveEventInputAccessControlOutput
}

type LiveEventInputAccessControlArgs struct {
	Ip IPAccessControlPtrInput `pulumi:"ip"`
}

func (LiveEventInputAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputAccessControl)(nil)).Elem()
}

func (i LiveEventInputAccessControlArgs) ToLiveEventInputAccessControlOutput() LiveEventInputAccessControlOutput {
	return i.ToLiveEventInputAccessControlOutputWithContext(context.Background())
}

func (i LiveEventInputAccessControlArgs) ToLiveEventInputAccessControlOutputWithContext(ctx context.Context) LiveEventInputAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputAccessControlOutput)
}

func (i LiveEventInputAccessControlArgs) ToLiveEventInputAccessControlPtrOutput() LiveEventInputAccessControlPtrOutput {
	return i.ToLiveEventInputAccessControlPtrOutputWithContext(context.Background())
}

func (i LiveEventInputAccessControlArgs) ToLiveEventInputAccessControlPtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputAccessControlOutput).ToLiveEventInputAccessControlPtrOutputWithContext(ctx)
}









type LiveEventInputAccessControlPtrInput interface {
	pulumi.Input

	ToLiveEventInputAccessControlPtrOutput() LiveEventInputAccessControlPtrOutput
	ToLiveEventInputAccessControlPtrOutputWithContext(context.Context) LiveEventInputAccessControlPtrOutput
}

type liveEventInputAccessControlPtrType LiveEventInputAccessControlArgs

func LiveEventInputAccessControlPtr(v *LiveEventInputAccessControlArgs) LiveEventInputAccessControlPtrInput {
	return (*liveEventInputAccessControlPtrType)(v)
}

func (*liveEventInputAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputAccessControl)(nil)).Elem()
}

func (i *liveEventInputAccessControlPtrType) ToLiveEventInputAccessControlPtrOutput() LiveEventInputAccessControlPtrOutput {
	return i.ToLiveEventInputAccessControlPtrOutputWithContext(context.Background())
}

func (i *liveEventInputAccessControlPtrType) ToLiveEventInputAccessControlPtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputAccessControlPtrOutput)
}

type LiveEventInputAccessControlOutput struct{ *pulumi.OutputState }

func (LiveEventInputAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputAccessControl)(nil)).Elem()
}

func (o LiveEventInputAccessControlOutput) ToLiveEventInputAccessControlOutput() LiveEventInputAccessControlOutput {
	return o
}

func (o LiveEventInputAccessControlOutput) ToLiveEventInputAccessControlOutputWithContext(ctx context.Context) LiveEventInputAccessControlOutput {
	return o
}

func (o LiveEventInputAccessControlOutput) ToLiveEventInputAccessControlPtrOutput() LiveEventInputAccessControlPtrOutput {
	return o.ToLiveEventInputAccessControlPtrOutputWithContext(context.Background())
}

func (o LiveEventInputAccessControlOutput) ToLiveEventInputAccessControlPtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventInputAccessControl) *LiveEventInputAccessControl {
		return &v
	}).(LiveEventInputAccessControlPtrOutput)
}

func (o LiveEventInputAccessControlOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v LiveEventInputAccessControl) *IPAccessControl { return v.Ip }).(IPAccessControlPtrOutput)
}

type LiveEventInputAccessControlPtrOutput struct{ *pulumi.OutputState }

func (LiveEventInputAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputAccessControl)(nil)).Elem()
}

func (o LiveEventInputAccessControlPtrOutput) ToLiveEventInputAccessControlPtrOutput() LiveEventInputAccessControlPtrOutput {
	return o
}

func (o LiveEventInputAccessControlPtrOutput) ToLiveEventInputAccessControlPtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlPtrOutput {
	return o
}

func (o LiveEventInputAccessControlPtrOutput) Elem() LiveEventInputAccessControlOutput {
	return o.ApplyT(func(v *LiveEventInputAccessControl) LiveEventInputAccessControl {
		if v != nil {
			return *v
		}
		var ret LiveEventInputAccessControl
		return ret
	}).(LiveEventInputAccessControlOutput)
}

func (o LiveEventInputAccessControlPtrOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v *LiveEventInputAccessControl) *IPAccessControl {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlPtrOutput)
}

type LiveEventInputAccessControlResponse struct {
	Ip *IPAccessControlResponse `pulumi:"ip"`
}





type LiveEventInputAccessControlResponseInput interface {
	pulumi.Input

	ToLiveEventInputAccessControlResponseOutput() LiveEventInputAccessControlResponseOutput
	ToLiveEventInputAccessControlResponseOutputWithContext(context.Context) LiveEventInputAccessControlResponseOutput
}

type LiveEventInputAccessControlResponseArgs struct {
	Ip IPAccessControlResponsePtrInput `pulumi:"ip"`
}

func (LiveEventInputAccessControlResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputAccessControlResponse)(nil)).Elem()
}

func (i LiveEventInputAccessControlResponseArgs) ToLiveEventInputAccessControlResponseOutput() LiveEventInputAccessControlResponseOutput {
	return i.ToLiveEventInputAccessControlResponseOutputWithContext(context.Background())
}

func (i LiveEventInputAccessControlResponseArgs) ToLiveEventInputAccessControlResponseOutputWithContext(ctx context.Context) LiveEventInputAccessControlResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputAccessControlResponseOutput)
}

func (i LiveEventInputAccessControlResponseArgs) ToLiveEventInputAccessControlResponsePtrOutput() LiveEventInputAccessControlResponsePtrOutput {
	return i.ToLiveEventInputAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i LiveEventInputAccessControlResponseArgs) ToLiveEventInputAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputAccessControlResponseOutput).ToLiveEventInputAccessControlResponsePtrOutputWithContext(ctx)
}









type LiveEventInputAccessControlResponsePtrInput interface {
	pulumi.Input

	ToLiveEventInputAccessControlResponsePtrOutput() LiveEventInputAccessControlResponsePtrOutput
	ToLiveEventInputAccessControlResponsePtrOutputWithContext(context.Context) LiveEventInputAccessControlResponsePtrOutput
}

type liveEventInputAccessControlResponsePtrType LiveEventInputAccessControlResponseArgs

func LiveEventInputAccessControlResponsePtr(v *LiveEventInputAccessControlResponseArgs) LiveEventInputAccessControlResponsePtrInput {
	return (*liveEventInputAccessControlResponsePtrType)(v)
}

func (*liveEventInputAccessControlResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputAccessControlResponse)(nil)).Elem()
}

func (i *liveEventInputAccessControlResponsePtrType) ToLiveEventInputAccessControlResponsePtrOutput() LiveEventInputAccessControlResponsePtrOutput {
	return i.ToLiveEventInputAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i *liveEventInputAccessControlResponsePtrType) ToLiveEventInputAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputAccessControlResponsePtrOutput)
}

type LiveEventInputAccessControlResponseOutput struct{ *pulumi.OutputState }

func (LiveEventInputAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputAccessControlResponse)(nil)).Elem()
}

func (o LiveEventInputAccessControlResponseOutput) ToLiveEventInputAccessControlResponseOutput() LiveEventInputAccessControlResponseOutput {
	return o
}

func (o LiveEventInputAccessControlResponseOutput) ToLiveEventInputAccessControlResponseOutputWithContext(ctx context.Context) LiveEventInputAccessControlResponseOutput {
	return o
}

func (o LiveEventInputAccessControlResponseOutput) ToLiveEventInputAccessControlResponsePtrOutput() LiveEventInputAccessControlResponsePtrOutput {
	return o.ToLiveEventInputAccessControlResponsePtrOutputWithContext(context.Background())
}

func (o LiveEventInputAccessControlResponseOutput) ToLiveEventInputAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventInputAccessControlResponse) *LiveEventInputAccessControlResponse {
		return &v
	}).(LiveEventInputAccessControlResponsePtrOutput)
}

func (o LiveEventInputAccessControlResponseOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v LiveEventInputAccessControlResponse) *IPAccessControlResponse { return v.Ip }).(IPAccessControlResponsePtrOutput)
}

type LiveEventInputAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventInputAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputAccessControlResponse)(nil)).Elem()
}

func (o LiveEventInputAccessControlResponsePtrOutput) ToLiveEventInputAccessControlResponsePtrOutput() LiveEventInputAccessControlResponsePtrOutput {
	return o
}

func (o LiveEventInputAccessControlResponsePtrOutput) ToLiveEventInputAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlResponsePtrOutput {
	return o
}

func (o LiveEventInputAccessControlResponsePtrOutput) Elem() LiveEventInputAccessControlResponseOutput {
	return o.ApplyT(func(v *LiveEventInputAccessControlResponse) LiveEventInputAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventInputAccessControlResponse
		return ret
	}).(LiveEventInputAccessControlResponseOutput)
}

func (o LiveEventInputAccessControlResponsePtrOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *LiveEventInputAccessControlResponse) *IPAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlResponsePtrOutput)
}

type LiveEventInputResponse struct {
	AccessControl            *LiveEventInputAccessControlResponse `pulumi:"accessControl"`
	AccessToken              *string                              `pulumi:"accessToken"`
	Endpoints                []LiveEventEndpointResponse          `pulumi:"endpoints"`
	KeyFrameIntervalDuration *string                              `pulumi:"keyFrameIntervalDuration"`
	StreamingProtocol        string                               `pulumi:"streamingProtocol"`
}





type LiveEventInputResponseInput interface {
	pulumi.Input

	ToLiveEventInputResponseOutput() LiveEventInputResponseOutput
	ToLiveEventInputResponseOutputWithContext(context.Context) LiveEventInputResponseOutput
}

type LiveEventInputResponseArgs struct {
	AccessControl            LiveEventInputAccessControlResponsePtrInput `pulumi:"accessControl"`
	AccessToken              pulumi.StringPtrInput                       `pulumi:"accessToken"`
	Endpoints                LiveEventEndpointResponseArrayInput         `pulumi:"endpoints"`
	KeyFrameIntervalDuration pulumi.StringPtrInput                       `pulumi:"keyFrameIntervalDuration"`
	StreamingProtocol        pulumi.StringInput                          `pulumi:"streamingProtocol"`
}

func (LiveEventInputResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputResponse)(nil)).Elem()
}

func (i LiveEventInputResponseArgs) ToLiveEventInputResponseOutput() LiveEventInputResponseOutput {
	return i.ToLiveEventInputResponseOutputWithContext(context.Background())
}

func (i LiveEventInputResponseArgs) ToLiveEventInputResponseOutputWithContext(ctx context.Context) LiveEventInputResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputResponseOutput)
}

func (i LiveEventInputResponseArgs) ToLiveEventInputResponsePtrOutput() LiveEventInputResponsePtrOutput {
	return i.ToLiveEventInputResponsePtrOutputWithContext(context.Background())
}

func (i LiveEventInputResponseArgs) ToLiveEventInputResponsePtrOutputWithContext(ctx context.Context) LiveEventInputResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputResponseOutput).ToLiveEventInputResponsePtrOutputWithContext(ctx)
}









type LiveEventInputResponsePtrInput interface {
	pulumi.Input

	ToLiveEventInputResponsePtrOutput() LiveEventInputResponsePtrOutput
	ToLiveEventInputResponsePtrOutputWithContext(context.Context) LiveEventInputResponsePtrOutput
}

type liveEventInputResponsePtrType LiveEventInputResponseArgs

func LiveEventInputResponsePtr(v *LiveEventInputResponseArgs) LiveEventInputResponsePtrInput {
	return (*liveEventInputResponsePtrType)(v)
}

func (*liveEventInputResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputResponse)(nil)).Elem()
}

func (i *liveEventInputResponsePtrType) ToLiveEventInputResponsePtrOutput() LiveEventInputResponsePtrOutput {
	return i.ToLiveEventInputResponsePtrOutputWithContext(context.Background())
}

func (i *liveEventInputResponsePtrType) ToLiveEventInputResponsePtrOutputWithContext(ctx context.Context) LiveEventInputResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputResponsePtrOutput)
}

type LiveEventInputResponseOutput struct{ *pulumi.OutputState }

func (LiveEventInputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputResponse)(nil)).Elem()
}

func (o LiveEventInputResponseOutput) ToLiveEventInputResponseOutput() LiveEventInputResponseOutput {
	return o
}

func (o LiveEventInputResponseOutput) ToLiveEventInputResponseOutputWithContext(ctx context.Context) LiveEventInputResponseOutput {
	return o
}

func (o LiveEventInputResponseOutput) ToLiveEventInputResponsePtrOutput() LiveEventInputResponsePtrOutput {
	return o.ToLiveEventInputResponsePtrOutputWithContext(context.Background())
}

func (o LiveEventInputResponseOutput) ToLiveEventInputResponsePtrOutputWithContext(ctx context.Context) LiveEventInputResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventInputResponse) *LiveEventInputResponse {
		return &v
	}).(LiveEventInputResponsePtrOutput)
}

func (o LiveEventInputResponseOutput) AccessControl() LiveEventInputAccessControlResponsePtrOutput {
	return o.ApplyT(func(v LiveEventInputResponse) *LiveEventInputAccessControlResponse { return v.AccessControl }).(LiveEventInputAccessControlResponsePtrOutput)
}

func (o LiveEventInputResponseOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputResponse) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputResponseOutput) Endpoints() LiveEventEndpointResponseArrayOutput {
	return o.ApplyT(func(v LiveEventInputResponse) []LiveEventEndpointResponse { return v.Endpoints }).(LiveEventEndpointResponseArrayOutput)
}

func (o LiveEventInputResponseOutput) KeyFrameIntervalDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputResponse) *string { return v.KeyFrameIntervalDuration }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputResponseOutput) StreamingProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LiveEventInputResponse) string { return v.StreamingProtocol }).(pulumi.StringOutput)
}

type LiveEventInputResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventInputResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputResponse)(nil)).Elem()
}

func (o LiveEventInputResponsePtrOutput) ToLiveEventInputResponsePtrOutput() LiveEventInputResponsePtrOutput {
	return o
}

func (o LiveEventInputResponsePtrOutput) ToLiveEventInputResponsePtrOutputWithContext(ctx context.Context) LiveEventInputResponsePtrOutput {
	return o
}

func (o LiveEventInputResponsePtrOutput) Elem() LiveEventInputResponseOutput {
	return o.ApplyT(func(v *LiveEventInputResponse) LiveEventInputResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventInputResponse
		return ret
	}).(LiveEventInputResponseOutput)
}

func (o LiveEventInputResponsePtrOutput) AccessControl() LiveEventInputAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *LiveEventInputResponse) *LiveEventInputAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.AccessControl
	}).(LiveEventInputAccessControlResponsePtrOutput)
}

func (o LiveEventInputResponsePtrOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventInputResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccessToken
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventInputResponsePtrOutput) Endpoints() LiveEventEndpointResponseArrayOutput {
	return o.ApplyT(func(v *LiveEventInputResponse) []LiveEventEndpointResponse {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(LiveEventEndpointResponseArrayOutput)
}

func (o LiveEventInputResponsePtrOutput) KeyFrameIntervalDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventInputResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyFrameIntervalDuration
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventInputResponsePtrOutput) StreamingProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventInputResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StreamingProtocol
	}).(pulumi.StringPtrOutput)
}

type LiveEventInputTrackSelection struct {
	Operation *string `pulumi:"operation"`
	Property  *string `pulumi:"property"`
	Value     *string `pulumi:"value"`
}





type LiveEventInputTrackSelectionInput interface {
	pulumi.Input

	ToLiveEventInputTrackSelectionOutput() LiveEventInputTrackSelectionOutput
	ToLiveEventInputTrackSelectionOutputWithContext(context.Context) LiveEventInputTrackSelectionOutput
}

type LiveEventInputTrackSelectionArgs struct {
	Operation pulumi.StringPtrInput `pulumi:"operation"`
	Property  pulumi.StringPtrInput `pulumi:"property"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (LiveEventInputTrackSelectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputTrackSelection)(nil)).Elem()
}

func (i LiveEventInputTrackSelectionArgs) ToLiveEventInputTrackSelectionOutput() LiveEventInputTrackSelectionOutput {
	return i.ToLiveEventInputTrackSelectionOutputWithContext(context.Background())
}

func (i LiveEventInputTrackSelectionArgs) ToLiveEventInputTrackSelectionOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTrackSelectionOutput)
}





type LiveEventInputTrackSelectionArrayInput interface {
	pulumi.Input

	ToLiveEventInputTrackSelectionArrayOutput() LiveEventInputTrackSelectionArrayOutput
	ToLiveEventInputTrackSelectionArrayOutputWithContext(context.Context) LiveEventInputTrackSelectionArrayOutput
}

type LiveEventInputTrackSelectionArray []LiveEventInputTrackSelectionInput

func (LiveEventInputTrackSelectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventInputTrackSelection)(nil)).Elem()
}

func (i LiveEventInputTrackSelectionArray) ToLiveEventInputTrackSelectionArrayOutput() LiveEventInputTrackSelectionArrayOutput {
	return i.ToLiveEventInputTrackSelectionArrayOutputWithContext(context.Background())
}

func (i LiveEventInputTrackSelectionArray) ToLiveEventInputTrackSelectionArrayOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTrackSelectionArrayOutput)
}

type LiveEventInputTrackSelectionOutput struct{ *pulumi.OutputState }

func (LiveEventInputTrackSelectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputTrackSelection)(nil)).Elem()
}

func (o LiveEventInputTrackSelectionOutput) ToLiveEventInputTrackSelectionOutput() LiveEventInputTrackSelectionOutput {
	return o
}

func (o LiveEventInputTrackSelectionOutput) ToLiveEventInputTrackSelectionOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionOutput {
	return o
}

func (o LiveEventInputTrackSelectionOutput) Operation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputTrackSelection) *string { return v.Operation }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTrackSelectionOutput) Property() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputTrackSelection) *string { return v.Property }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTrackSelectionOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputTrackSelection) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type LiveEventInputTrackSelectionArrayOutput struct{ *pulumi.OutputState }

func (LiveEventInputTrackSelectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventInputTrackSelection)(nil)).Elem()
}

func (o LiveEventInputTrackSelectionArrayOutput) ToLiveEventInputTrackSelectionArrayOutput() LiveEventInputTrackSelectionArrayOutput {
	return o
}

func (o LiveEventInputTrackSelectionArrayOutput) ToLiveEventInputTrackSelectionArrayOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionArrayOutput {
	return o
}

func (o LiveEventInputTrackSelectionArrayOutput) Index(i pulumi.IntInput) LiveEventInputTrackSelectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventInputTrackSelection {
		return vs[0].([]LiveEventInputTrackSelection)[vs[1].(int)]
	}).(LiveEventInputTrackSelectionOutput)
}

type LiveEventInputTrackSelectionResponse struct {
	Operation *string `pulumi:"operation"`
	Property  *string `pulumi:"property"`
	Value     *string `pulumi:"value"`
}





type LiveEventInputTrackSelectionResponseInput interface {
	pulumi.Input

	ToLiveEventInputTrackSelectionResponseOutput() LiveEventInputTrackSelectionResponseOutput
	ToLiveEventInputTrackSelectionResponseOutputWithContext(context.Context) LiveEventInputTrackSelectionResponseOutput
}

type LiveEventInputTrackSelectionResponseArgs struct {
	Operation pulumi.StringPtrInput `pulumi:"operation"`
	Property  pulumi.StringPtrInput `pulumi:"property"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (LiveEventInputTrackSelectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputTrackSelectionResponse)(nil)).Elem()
}

func (i LiveEventInputTrackSelectionResponseArgs) ToLiveEventInputTrackSelectionResponseOutput() LiveEventInputTrackSelectionResponseOutput {
	return i.ToLiveEventInputTrackSelectionResponseOutputWithContext(context.Background())
}

func (i LiveEventInputTrackSelectionResponseArgs) ToLiveEventInputTrackSelectionResponseOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTrackSelectionResponseOutput)
}





type LiveEventInputTrackSelectionResponseArrayInput interface {
	pulumi.Input

	ToLiveEventInputTrackSelectionResponseArrayOutput() LiveEventInputTrackSelectionResponseArrayOutput
	ToLiveEventInputTrackSelectionResponseArrayOutputWithContext(context.Context) LiveEventInputTrackSelectionResponseArrayOutput
}

type LiveEventInputTrackSelectionResponseArray []LiveEventInputTrackSelectionResponseInput

func (LiveEventInputTrackSelectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventInputTrackSelectionResponse)(nil)).Elem()
}

func (i LiveEventInputTrackSelectionResponseArray) ToLiveEventInputTrackSelectionResponseArrayOutput() LiveEventInputTrackSelectionResponseArrayOutput {
	return i.ToLiveEventInputTrackSelectionResponseArrayOutputWithContext(context.Background())
}

func (i LiveEventInputTrackSelectionResponseArray) ToLiveEventInputTrackSelectionResponseArrayOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTrackSelectionResponseArrayOutput)
}

type LiveEventInputTrackSelectionResponseOutput struct{ *pulumi.OutputState }

func (LiveEventInputTrackSelectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputTrackSelectionResponse)(nil)).Elem()
}

func (o LiveEventInputTrackSelectionResponseOutput) ToLiveEventInputTrackSelectionResponseOutput() LiveEventInputTrackSelectionResponseOutput {
	return o
}

func (o LiveEventInputTrackSelectionResponseOutput) ToLiveEventInputTrackSelectionResponseOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionResponseOutput {
	return o
}

func (o LiveEventInputTrackSelectionResponseOutput) Operation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputTrackSelectionResponse) *string { return v.Operation }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTrackSelectionResponseOutput) Property() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputTrackSelectionResponse) *string { return v.Property }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTrackSelectionResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputTrackSelectionResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type LiveEventInputTrackSelectionResponseArrayOutput struct{ *pulumi.OutputState }

func (LiveEventInputTrackSelectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventInputTrackSelectionResponse)(nil)).Elem()
}

func (o LiveEventInputTrackSelectionResponseArrayOutput) ToLiveEventInputTrackSelectionResponseArrayOutput() LiveEventInputTrackSelectionResponseArrayOutput {
	return o
}

func (o LiveEventInputTrackSelectionResponseArrayOutput) ToLiveEventInputTrackSelectionResponseArrayOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionResponseArrayOutput {
	return o
}

func (o LiveEventInputTrackSelectionResponseArrayOutput) Index(i pulumi.IntInput) LiveEventInputTrackSelectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventInputTrackSelectionResponse {
		return vs[0].([]LiveEventInputTrackSelectionResponse)[vs[1].(int)]
	}).(LiveEventInputTrackSelectionResponseOutput)
}

type LiveEventOutputTranscriptionTrack struct {
	TrackName string `pulumi:"trackName"`
}





type LiveEventOutputTranscriptionTrackInput interface {
	pulumi.Input

	ToLiveEventOutputTranscriptionTrackOutput() LiveEventOutputTranscriptionTrackOutput
	ToLiveEventOutputTranscriptionTrackOutputWithContext(context.Context) LiveEventOutputTranscriptionTrackOutput
}

type LiveEventOutputTranscriptionTrackArgs struct {
	TrackName pulumi.StringInput `pulumi:"trackName"`
}

func (LiveEventOutputTranscriptionTrackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventOutputTranscriptionTrack)(nil)).Elem()
}

func (i LiveEventOutputTranscriptionTrackArgs) ToLiveEventOutputTranscriptionTrackOutput() LiveEventOutputTranscriptionTrackOutput {
	return i.ToLiveEventOutputTranscriptionTrackOutputWithContext(context.Background())
}

func (i LiveEventOutputTranscriptionTrackArgs) ToLiveEventOutputTranscriptionTrackOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutputTranscriptionTrackOutput)
}

func (i LiveEventOutputTranscriptionTrackArgs) ToLiveEventOutputTranscriptionTrackPtrOutput() LiveEventOutputTranscriptionTrackPtrOutput {
	return i.ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(context.Background())
}

func (i LiveEventOutputTranscriptionTrackArgs) ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutputTranscriptionTrackOutput).ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(ctx)
}









type LiveEventOutputTranscriptionTrackPtrInput interface {
	pulumi.Input

	ToLiveEventOutputTranscriptionTrackPtrOutput() LiveEventOutputTranscriptionTrackPtrOutput
	ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(context.Context) LiveEventOutputTranscriptionTrackPtrOutput
}

type liveEventOutputTranscriptionTrackPtrType LiveEventOutputTranscriptionTrackArgs

func LiveEventOutputTranscriptionTrackPtr(v *LiveEventOutputTranscriptionTrackArgs) LiveEventOutputTranscriptionTrackPtrInput {
	return (*liveEventOutputTranscriptionTrackPtrType)(v)
}

func (*liveEventOutputTranscriptionTrackPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventOutputTranscriptionTrack)(nil)).Elem()
}

func (i *liveEventOutputTranscriptionTrackPtrType) ToLiveEventOutputTranscriptionTrackPtrOutput() LiveEventOutputTranscriptionTrackPtrOutput {
	return i.ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(context.Background())
}

func (i *liveEventOutputTranscriptionTrackPtrType) ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutputTranscriptionTrackPtrOutput)
}

type LiveEventOutputTranscriptionTrackOutput struct{ *pulumi.OutputState }

func (LiveEventOutputTranscriptionTrackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventOutputTranscriptionTrack)(nil)).Elem()
}

func (o LiveEventOutputTranscriptionTrackOutput) ToLiveEventOutputTranscriptionTrackOutput() LiveEventOutputTranscriptionTrackOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackOutput) ToLiveEventOutputTranscriptionTrackOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackOutput) ToLiveEventOutputTranscriptionTrackPtrOutput() LiveEventOutputTranscriptionTrackPtrOutput {
	return o.ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(context.Background())
}

func (o LiveEventOutputTranscriptionTrackOutput) ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventOutputTranscriptionTrack) *LiveEventOutputTranscriptionTrack {
		return &v
	}).(LiveEventOutputTranscriptionTrackPtrOutput)
}

func (o LiveEventOutputTranscriptionTrackOutput) TrackName() pulumi.StringOutput {
	return o.ApplyT(func(v LiveEventOutputTranscriptionTrack) string { return v.TrackName }).(pulumi.StringOutput)
}

type LiveEventOutputTranscriptionTrackPtrOutput struct{ *pulumi.OutputState }

func (LiveEventOutputTranscriptionTrackPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventOutputTranscriptionTrack)(nil)).Elem()
}

func (o LiveEventOutputTranscriptionTrackPtrOutput) ToLiveEventOutputTranscriptionTrackPtrOutput() LiveEventOutputTranscriptionTrackPtrOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackPtrOutput) ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackPtrOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackPtrOutput) Elem() LiveEventOutputTranscriptionTrackOutput {
	return o.ApplyT(func(v *LiveEventOutputTranscriptionTrack) LiveEventOutputTranscriptionTrack {
		if v != nil {
			return *v
		}
		var ret LiveEventOutputTranscriptionTrack
		return ret
	}).(LiveEventOutputTranscriptionTrackOutput)
}

func (o LiveEventOutputTranscriptionTrackPtrOutput) TrackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventOutputTranscriptionTrack) *string {
		if v == nil {
			return nil
		}
		return &v.TrackName
	}).(pulumi.StringPtrOutput)
}

type LiveEventOutputTranscriptionTrackResponse struct {
	TrackName string `pulumi:"trackName"`
}





type LiveEventOutputTranscriptionTrackResponseInput interface {
	pulumi.Input

	ToLiveEventOutputTranscriptionTrackResponseOutput() LiveEventOutputTranscriptionTrackResponseOutput
	ToLiveEventOutputTranscriptionTrackResponseOutputWithContext(context.Context) LiveEventOutputTranscriptionTrackResponseOutput
}

type LiveEventOutputTranscriptionTrackResponseArgs struct {
	TrackName pulumi.StringInput `pulumi:"trackName"`
}

func (LiveEventOutputTranscriptionTrackResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventOutputTranscriptionTrackResponse)(nil)).Elem()
}

func (i LiveEventOutputTranscriptionTrackResponseArgs) ToLiveEventOutputTranscriptionTrackResponseOutput() LiveEventOutputTranscriptionTrackResponseOutput {
	return i.ToLiveEventOutputTranscriptionTrackResponseOutputWithContext(context.Background())
}

func (i LiveEventOutputTranscriptionTrackResponseArgs) ToLiveEventOutputTranscriptionTrackResponseOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutputTranscriptionTrackResponseOutput)
}

func (i LiveEventOutputTranscriptionTrackResponseArgs) ToLiveEventOutputTranscriptionTrackResponsePtrOutput() LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return i.ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(context.Background())
}

func (i LiveEventOutputTranscriptionTrackResponseArgs) ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutputTranscriptionTrackResponseOutput).ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(ctx)
}









type LiveEventOutputTranscriptionTrackResponsePtrInput interface {
	pulumi.Input

	ToLiveEventOutputTranscriptionTrackResponsePtrOutput() LiveEventOutputTranscriptionTrackResponsePtrOutput
	ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(context.Context) LiveEventOutputTranscriptionTrackResponsePtrOutput
}

type liveEventOutputTranscriptionTrackResponsePtrType LiveEventOutputTranscriptionTrackResponseArgs

func LiveEventOutputTranscriptionTrackResponsePtr(v *LiveEventOutputTranscriptionTrackResponseArgs) LiveEventOutputTranscriptionTrackResponsePtrInput {
	return (*liveEventOutputTranscriptionTrackResponsePtrType)(v)
}

func (*liveEventOutputTranscriptionTrackResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventOutputTranscriptionTrackResponse)(nil)).Elem()
}

func (i *liveEventOutputTranscriptionTrackResponsePtrType) ToLiveEventOutputTranscriptionTrackResponsePtrOutput() LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return i.ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(context.Background())
}

func (i *liveEventOutputTranscriptionTrackResponsePtrType) ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutputTranscriptionTrackResponsePtrOutput)
}

type LiveEventOutputTranscriptionTrackResponseOutput struct{ *pulumi.OutputState }

func (LiveEventOutputTranscriptionTrackResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventOutputTranscriptionTrackResponse)(nil)).Elem()
}

func (o LiveEventOutputTranscriptionTrackResponseOutput) ToLiveEventOutputTranscriptionTrackResponseOutput() LiveEventOutputTranscriptionTrackResponseOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackResponseOutput) ToLiveEventOutputTranscriptionTrackResponseOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackResponseOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackResponseOutput) ToLiveEventOutputTranscriptionTrackResponsePtrOutput() LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return o.ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(context.Background())
}

func (o LiveEventOutputTranscriptionTrackResponseOutput) ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventOutputTranscriptionTrackResponse) *LiveEventOutputTranscriptionTrackResponse {
		return &v
	}).(LiveEventOutputTranscriptionTrackResponsePtrOutput)
}

func (o LiveEventOutputTranscriptionTrackResponseOutput) TrackName() pulumi.StringOutput {
	return o.ApplyT(func(v LiveEventOutputTranscriptionTrackResponse) string { return v.TrackName }).(pulumi.StringOutput)
}

type LiveEventOutputTranscriptionTrackResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventOutputTranscriptionTrackResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventOutputTranscriptionTrackResponse)(nil)).Elem()
}

func (o LiveEventOutputTranscriptionTrackResponsePtrOutput) ToLiveEventOutputTranscriptionTrackResponsePtrOutput() LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackResponsePtrOutput) ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackResponsePtrOutput) Elem() LiveEventOutputTranscriptionTrackResponseOutput {
	return o.ApplyT(func(v *LiveEventOutputTranscriptionTrackResponse) LiveEventOutputTranscriptionTrackResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventOutputTranscriptionTrackResponse
		return ret
	}).(LiveEventOutputTranscriptionTrackResponseOutput)
}

func (o LiveEventOutputTranscriptionTrackResponsePtrOutput) TrackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventOutputTranscriptionTrackResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TrackName
	}).(pulumi.StringPtrOutput)
}

type LiveEventPreview struct {
	AccessControl       *LiveEventPreviewAccessControl `pulumi:"accessControl"`
	AlternativeMediaId  *string                        `pulumi:"alternativeMediaId"`
	Endpoints           []LiveEventEndpoint            `pulumi:"endpoints"`
	PreviewLocator      *string                        `pulumi:"previewLocator"`
	StreamingPolicyName *string                        `pulumi:"streamingPolicyName"`
}





type LiveEventPreviewInput interface {
	pulumi.Input

	ToLiveEventPreviewOutput() LiveEventPreviewOutput
	ToLiveEventPreviewOutputWithContext(context.Context) LiveEventPreviewOutput
}

type LiveEventPreviewArgs struct {
	AccessControl       LiveEventPreviewAccessControlPtrInput `pulumi:"accessControl"`
	AlternativeMediaId  pulumi.StringPtrInput                 `pulumi:"alternativeMediaId"`
	Endpoints           LiveEventEndpointArrayInput           `pulumi:"endpoints"`
	PreviewLocator      pulumi.StringPtrInput                 `pulumi:"previewLocator"`
	StreamingPolicyName pulumi.StringPtrInput                 `pulumi:"streamingPolicyName"`
}

func (LiveEventPreviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreview)(nil)).Elem()
}

func (i LiveEventPreviewArgs) ToLiveEventPreviewOutput() LiveEventPreviewOutput {
	return i.ToLiveEventPreviewOutputWithContext(context.Background())
}

func (i LiveEventPreviewArgs) ToLiveEventPreviewOutputWithContext(ctx context.Context) LiveEventPreviewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewOutput)
}

func (i LiveEventPreviewArgs) ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput {
	return i.ToLiveEventPreviewPtrOutputWithContext(context.Background())
}

func (i LiveEventPreviewArgs) ToLiveEventPreviewPtrOutputWithContext(ctx context.Context) LiveEventPreviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewOutput).ToLiveEventPreviewPtrOutputWithContext(ctx)
}









type LiveEventPreviewPtrInput interface {
	pulumi.Input

	ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput
	ToLiveEventPreviewPtrOutputWithContext(context.Context) LiveEventPreviewPtrOutput
}

type liveEventPreviewPtrType LiveEventPreviewArgs

func LiveEventPreviewPtr(v *LiveEventPreviewArgs) LiveEventPreviewPtrInput {
	return (*liveEventPreviewPtrType)(v)
}

func (*liveEventPreviewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreview)(nil)).Elem()
}

func (i *liveEventPreviewPtrType) ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput {
	return i.ToLiveEventPreviewPtrOutputWithContext(context.Background())
}

func (i *liveEventPreviewPtrType) ToLiveEventPreviewPtrOutputWithContext(ctx context.Context) LiveEventPreviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewPtrOutput)
}

type LiveEventPreviewOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreview)(nil)).Elem()
}

func (o LiveEventPreviewOutput) ToLiveEventPreviewOutput() LiveEventPreviewOutput {
	return o
}

func (o LiveEventPreviewOutput) ToLiveEventPreviewOutputWithContext(ctx context.Context) LiveEventPreviewOutput {
	return o
}

func (o LiveEventPreviewOutput) ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput {
	return o.ToLiveEventPreviewPtrOutputWithContext(context.Background())
}

func (o LiveEventPreviewOutput) ToLiveEventPreviewPtrOutputWithContext(ctx context.Context) LiveEventPreviewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventPreview) *LiveEventPreview {
		return &v
	}).(LiveEventPreviewPtrOutput)
}

func (o LiveEventPreviewOutput) AccessControl() LiveEventPreviewAccessControlPtrOutput {
	return o.ApplyT(func(v LiveEventPreview) *LiveEventPreviewAccessControl { return v.AccessControl }).(LiveEventPreviewAccessControlPtrOutput)
}

func (o LiveEventPreviewOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreview) *string { return v.AlternativeMediaId }).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewOutput) Endpoints() LiveEventEndpointArrayOutput {
	return o.ApplyT(func(v LiveEventPreview) []LiveEventEndpoint { return v.Endpoints }).(LiveEventEndpointArrayOutput)
}

func (o LiveEventPreviewOutput) PreviewLocator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreview) *string { return v.PreviewLocator }).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewOutput) StreamingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreview) *string { return v.StreamingPolicyName }).(pulumi.StringPtrOutput)
}

type LiveEventPreviewPtrOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreview)(nil)).Elem()
}

func (o LiveEventPreviewPtrOutput) ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput {
	return o
}

func (o LiveEventPreviewPtrOutput) ToLiveEventPreviewPtrOutputWithContext(ctx context.Context) LiveEventPreviewPtrOutput {
	return o
}

func (o LiveEventPreviewPtrOutput) Elem() LiveEventPreviewOutput {
	return o.ApplyT(func(v *LiveEventPreview) LiveEventPreview {
		if v != nil {
			return *v
		}
		var ret LiveEventPreview
		return ret
	}).(LiveEventPreviewOutput)
}

func (o LiveEventPreviewPtrOutput) AccessControl() LiveEventPreviewAccessControlPtrOutput {
	return o.ApplyT(func(v *LiveEventPreview) *LiveEventPreviewAccessControl {
		if v == nil {
			return nil
		}
		return v.AccessControl
	}).(LiveEventPreviewAccessControlPtrOutput)
}

func (o LiveEventPreviewPtrOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreview) *string {
		if v == nil {
			return nil
		}
		return v.AlternativeMediaId
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewPtrOutput) Endpoints() LiveEventEndpointArrayOutput {
	return o.ApplyT(func(v *LiveEventPreview) []LiveEventEndpoint {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(LiveEventEndpointArrayOutput)
}

func (o LiveEventPreviewPtrOutput) PreviewLocator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreview) *string {
		if v == nil {
			return nil
		}
		return v.PreviewLocator
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewPtrOutput) StreamingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreview) *string {
		if v == nil {
			return nil
		}
		return v.StreamingPolicyName
	}).(pulumi.StringPtrOutput)
}

type LiveEventPreviewAccessControl struct {
	Ip *IPAccessControl `pulumi:"ip"`
}





type LiveEventPreviewAccessControlInput interface {
	pulumi.Input

	ToLiveEventPreviewAccessControlOutput() LiveEventPreviewAccessControlOutput
	ToLiveEventPreviewAccessControlOutputWithContext(context.Context) LiveEventPreviewAccessControlOutput
}

type LiveEventPreviewAccessControlArgs struct {
	Ip IPAccessControlPtrInput `pulumi:"ip"`
}

func (LiveEventPreviewAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewAccessControl)(nil)).Elem()
}

func (i LiveEventPreviewAccessControlArgs) ToLiveEventPreviewAccessControlOutput() LiveEventPreviewAccessControlOutput {
	return i.ToLiveEventPreviewAccessControlOutputWithContext(context.Background())
}

func (i LiveEventPreviewAccessControlArgs) ToLiveEventPreviewAccessControlOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlOutput)
}

func (i LiveEventPreviewAccessControlArgs) ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput {
	return i.ToLiveEventPreviewAccessControlPtrOutputWithContext(context.Background())
}

func (i LiveEventPreviewAccessControlArgs) ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlOutput).ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx)
}









type LiveEventPreviewAccessControlPtrInput interface {
	pulumi.Input

	ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput
	ToLiveEventPreviewAccessControlPtrOutputWithContext(context.Context) LiveEventPreviewAccessControlPtrOutput
}

type liveEventPreviewAccessControlPtrType LiveEventPreviewAccessControlArgs

func LiveEventPreviewAccessControlPtr(v *LiveEventPreviewAccessControlArgs) LiveEventPreviewAccessControlPtrInput {
	return (*liveEventPreviewAccessControlPtrType)(v)
}

func (*liveEventPreviewAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewAccessControl)(nil)).Elem()
}

func (i *liveEventPreviewAccessControlPtrType) ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput {
	return i.ToLiveEventPreviewAccessControlPtrOutputWithContext(context.Background())
}

func (i *liveEventPreviewAccessControlPtrType) ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlPtrOutput)
}

type LiveEventPreviewAccessControlOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewAccessControl)(nil)).Elem()
}

func (o LiveEventPreviewAccessControlOutput) ToLiveEventPreviewAccessControlOutput() LiveEventPreviewAccessControlOutput {
	return o
}

func (o LiveEventPreviewAccessControlOutput) ToLiveEventPreviewAccessControlOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlOutput {
	return o
}

func (o LiveEventPreviewAccessControlOutput) ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput {
	return o.ToLiveEventPreviewAccessControlPtrOutputWithContext(context.Background())
}

func (o LiveEventPreviewAccessControlOutput) ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventPreviewAccessControl) *LiveEventPreviewAccessControl {
		return &v
	}).(LiveEventPreviewAccessControlPtrOutput)
}

func (o LiveEventPreviewAccessControlOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v LiveEventPreviewAccessControl) *IPAccessControl { return v.Ip }).(IPAccessControlPtrOutput)
}

type LiveEventPreviewAccessControlPtrOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewAccessControl)(nil)).Elem()
}

func (o LiveEventPreviewAccessControlPtrOutput) ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput {
	return o
}

func (o LiveEventPreviewAccessControlPtrOutput) ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlPtrOutput {
	return o
}

func (o LiveEventPreviewAccessControlPtrOutput) Elem() LiveEventPreviewAccessControlOutput {
	return o.ApplyT(func(v *LiveEventPreviewAccessControl) LiveEventPreviewAccessControl {
		if v != nil {
			return *v
		}
		var ret LiveEventPreviewAccessControl
		return ret
	}).(LiveEventPreviewAccessControlOutput)
}

func (o LiveEventPreviewAccessControlPtrOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewAccessControl) *IPAccessControl {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlPtrOutput)
}

type LiveEventPreviewAccessControlResponse struct {
	Ip *IPAccessControlResponse `pulumi:"ip"`
}





type LiveEventPreviewAccessControlResponseInput interface {
	pulumi.Input

	ToLiveEventPreviewAccessControlResponseOutput() LiveEventPreviewAccessControlResponseOutput
	ToLiveEventPreviewAccessControlResponseOutputWithContext(context.Context) LiveEventPreviewAccessControlResponseOutput
}

type LiveEventPreviewAccessControlResponseArgs struct {
	Ip IPAccessControlResponsePtrInput `pulumi:"ip"`
}

func (LiveEventPreviewAccessControlResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewAccessControlResponse)(nil)).Elem()
}

func (i LiveEventPreviewAccessControlResponseArgs) ToLiveEventPreviewAccessControlResponseOutput() LiveEventPreviewAccessControlResponseOutput {
	return i.ToLiveEventPreviewAccessControlResponseOutputWithContext(context.Background())
}

func (i LiveEventPreviewAccessControlResponseArgs) ToLiveEventPreviewAccessControlResponseOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlResponseOutput)
}

func (i LiveEventPreviewAccessControlResponseArgs) ToLiveEventPreviewAccessControlResponsePtrOutput() LiveEventPreviewAccessControlResponsePtrOutput {
	return i.ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i LiveEventPreviewAccessControlResponseArgs) ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlResponseOutput).ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(ctx)
}









type LiveEventPreviewAccessControlResponsePtrInput interface {
	pulumi.Input

	ToLiveEventPreviewAccessControlResponsePtrOutput() LiveEventPreviewAccessControlResponsePtrOutput
	ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(context.Context) LiveEventPreviewAccessControlResponsePtrOutput
}

type liveEventPreviewAccessControlResponsePtrType LiveEventPreviewAccessControlResponseArgs

func LiveEventPreviewAccessControlResponsePtr(v *LiveEventPreviewAccessControlResponseArgs) LiveEventPreviewAccessControlResponsePtrInput {
	return (*liveEventPreviewAccessControlResponsePtrType)(v)
}

func (*liveEventPreviewAccessControlResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewAccessControlResponse)(nil)).Elem()
}

func (i *liveEventPreviewAccessControlResponsePtrType) ToLiveEventPreviewAccessControlResponsePtrOutput() LiveEventPreviewAccessControlResponsePtrOutput {
	return i.ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i *liveEventPreviewAccessControlResponsePtrType) ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlResponsePtrOutput)
}

type LiveEventPreviewAccessControlResponseOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewAccessControlResponse)(nil)).Elem()
}

func (o LiveEventPreviewAccessControlResponseOutput) ToLiveEventPreviewAccessControlResponseOutput() LiveEventPreviewAccessControlResponseOutput {
	return o
}

func (o LiveEventPreviewAccessControlResponseOutput) ToLiveEventPreviewAccessControlResponseOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponseOutput {
	return o
}

func (o LiveEventPreviewAccessControlResponseOutput) ToLiveEventPreviewAccessControlResponsePtrOutput() LiveEventPreviewAccessControlResponsePtrOutput {
	return o.ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(context.Background())
}

func (o LiveEventPreviewAccessControlResponseOutput) ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventPreviewAccessControlResponse) *LiveEventPreviewAccessControlResponse {
		return &v
	}).(LiveEventPreviewAccessControlResponsePtrOutput)
}

func (o LiveEventPreviewAccessControlResponseOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v LiveEventPreviewAccessControlResponse) *IPAccessControlResponse { return v.Ip }).(IPAccessControlResponsePtrOutput)
}

type LiveEventPreviewAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewAccessControlResponse)(nil)).Elem()
}

func (o LiveEventPreviewAccessControlResponsePtrOutput) ToLiveEventPreviewAccessControlResponsePtrOutput() LiveEventPreviewAccessControlResponsePtrOutput {
	return o
}

func (o LiveEventPreviewAccessControlResponsePtrOutput) ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponsePtrOutput {
	return o
}

func (o LiveEventPreviewAccessControlResponsePtrOutput) Elem() LiveEventPreviewAccessControlResponseOutput {
	return o.ApplyT(func(v *LiveEventPreviewAccessControlResponse) LiveEventPreviewAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventPreviewAccessControlResponse
		return ret
	}).(LiveEventPreviewAccessControlResponseOutput)
}

func (o LiveEventPreviewAccessControlResponsePtrOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewAccessControlResponse) *IPAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlResponsePtrOutput)
}

type LiveEventPreviewResponse struct {
	AccessControl       *LiveEventPreviewAccessControlResponse `pulumi:"accessControl"`
	AlternativeMediaId  *string                                `pulumi:"alternativeMediaId"`
	Endpoints           []LiveEventEndpointResponse            `pulumi:"endpoints"`
	PreviewLocator      *string                                `pulumi:"previewLocator"`
	StreamingPolicyName *string                                `pulumi:"streamingPolicyName"`
}





type LiveEventPreviewResponseInput interface {
	pulumi.Input

	ToLiveEventPreviewResponseOutput() LiveEventPreviewResponseOutput
	ToLiveEventPreviewResponseOutputWithContext(context.Context) LiveEventPreviewResponseOutput
}

type LiveEventPreviewResponseArgs struct {
	AccessControl       LiveEventPreviewAccessControlResponsePtrInput `pulumi:"accessControl"`
	AlternativeMediaId  pulumi.StringPtrInput                         `pulumi:"alternativeMediaId"`
	Endpoints           LiveEventEndpointResponseArrayInput           `pulumi:"endpoints"`
	PreviewLocator      pulumi.StringPtrInput                         `pulumi:"previewLocator"`
	StreamingPolicyName pulumi.StringPtrInput                         `pulumi:"streamingPolicyName"`
}

func (LiveEventPreviewResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewResponse)(nil)).Elem()
}

func (i LiveEventPreviewResponseArgs) ToLiveEventPreviewResponseOutput() LiveEventPreviewResponseOutput {
	return i.ToLiveEventPreviewResponseOutputWithContext(context.Background())
}

func (i LiveEventPreviewResponseArgs) ToLiveEventPreviewResponseOutputWithContext(ctx context.Context) LiveEventPreviewResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewResponseOutput)
}

func (i LiveEventPreviewResponseArgs) ToLiveEventPreviewResponsePtrOutput() LiveEventPreviewResponsePtrOutput {
	return i.ToLiveEventPreviewResponsePtrOutputWithContext(context.Background())
}

func (i LiveEventPreviewResponseArgs) ToLiveEventPreviewResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewResponseOutput).ToLiveEventPreviewResponsePtrOutputWithContext(ctx)
}









type LiveEventPreviewResponsePtrInput interface {
	pulumi.Input

	ToLiveEventPreviewResponsePtrOutput() LiveEventPreviewResponsePtrOutput
	ToLiveEventPreviewResponsePtrOutputWithContext(context.Context) LiveEventPreviewResponsePtrOutput
}

type liveEventPreviewResponsePtrType LiveEventPreviewResponseArgs

func LiveEventPreviewResponsePtr(v *LiveEventPreviewResponseArgs) LiveEventPreviewResponsePtrInput {
	return (*liveEventPreviewResponsePtrType)(v)
}

func (*liveEventPreviewResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewResponse)(nil)).Elem()
}

func (i *liveEventPreviewResponsePtrType) ToLiveEventPreviewResponsePtrOutput() LiveEventPreviewResponsePtrOutput {
	return i.ToLiveEventPreviewResponsePtrOutputWithContext(context.Background())
}

func (i *liveEventPreviewResponsePtrType) ToLiveEventPreviewResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewResponsePtrOutput)
}

type LiveEventPreviewResponseOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewResponse)(nil)).Elem()
}

func (o LiveEventPreviewResponseOutput) ToLiveEventPreviewResponseOutput() LiveEventPreviewResponseOutput {
	return o
}

func (o LiveEventPreviewResponseOutput) ToLiveEventPreviewResponseOutputWithContext(ctx context.Context) LiveEventPreviewResponseOutput {
	return o
}

func (o LiveEventPreviewResponseOutput) ToLiveEventPreviewResponsePtrOutput() LiveEventPreviewResponsePtrOutput {
	return o.ToLiveEventPreviewResponsePtrOutputWithContext(context.Background())
}

func (o LiveEventPreviewResponseOutput) ToLiveEventPreviewResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventPreviewResponse) *LiveEventPreviewResponse {
		return &v
	}).(LiveEventPreviewResponsePtrOutput)
}

func (o LiveEventPreviewResponseOutput) AccessControl() LiveEventPreviewAccessControlResponsePtrOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) *LiveEventPreviewAccessControlResponse { return v.AccessControl }).(LiveEventPreviewAccessControlResponsePtrOutput)
}

func (o LiveEventPreviewResponseOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) *string { return v.AlternativeMediaId }).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewResponseOutput) Endpoints() LiveEventEndpointResponseArrayOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) []LiveEventEndpointResponse { return v.Endpoints }).(LiveEventEndpointResponseArrayOutput)
}

func (o LiveEventPreviewResponseOutput) PreviewLocator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) *string { return v.PreviewLocator }).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewResponseOutput) StreamingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) *string { return v.StreamingPolicyName }).(pulumi.StringPtrOutput)
}

type LiveEventPreviewResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewResponse)(nil)).Elem()
}

func (o LiveEventPreviewResponsePtrOutput) ToLiveEventPreviewResponsePtrOutput() LiveEventPreviewResponsePtrOutput {
	return o
}

func (o LiveEventPreviewResponsePtrOutput) ToLiveEventPreviewResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewResponsePtrOutput {
	return o
}

func (o LiveEventPreviewResponsePtrOutput) Elem() LiveEventPreviewResponseOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) LiveEventPreviewResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventPreviewResponse
		return ret
	}).(LiveEventPreviewResponseOutput)
}

func (o LiveEventPreviewResponsePtrOutput) AccessControl() LiveEventPreviewAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) *LiveEventPreviewAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.AccessControl
	}).(LiveEventPreviewAccessControlResponsePtrOutput)
}

func (o LiveEventPreviewResponsePtrOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) *string {
		if v == nil {
			return nil
		}
		return v.AlternativeMediaId
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewResponsePtrOutput) Endpoints() LiveEventEndpointResponseArrayOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) []LiveEventEndpointResponse {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(LiveEventEndpointResponseArrayOutput)
}

func (o LiveEventPreviewResponsePtrOutput) PreviewLocator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) *string {
		if v == nil {
			return nil
		}
		return v.PreviewLocator
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewResponsePtrOutput) StreamingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) *string {
		if v == nil {
			return nil
		}
		return v.StreamingPolicyName
	}).(pulumi.StringPtrOutput)
}

type LiveEventTranscription struct {
	InputTrackSelection      []LiveEventInputTrackSelection     `pulumi:"inputTrackSelection"`
	Language                 *string                            `pulumi:"language"`
	OutputTranscriptionTrack *LiveEventOutputTranscriptionTrack `pulumi:"outputTranscriptionTrack"`
}





type LiveEventTranscriptionInput interface {
	pulumi.Input

	ToLiveEventTranscriptionOutput() LiveEventTranscriptionOutput
	ToLiveEventTranscriptionOutputWithContext(context.Context) LiveEventTranscriptionOutput
}

type LiveEventTranscriptionArgs struct {
	InputTrackSelection      LiveEventInputTrackSelectionArrayInput    `pulumi:"inputTrackSelection"`
	Language                 pulumi.StringPtrInput                     `pulumi:"language"`
	OutputTranscriptionTrack LiveEventOutputTranscriptionTrackPtrInput `pulumi:"outputTranscriptionTrack"`
}

func (LiveEventTranscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventTranscription)(nil)).Elem()
}

func (i LiveEventTranscriptionArgs) ToLiveEventTranscriptionOutput() LiveEventTranscriptionOutput {
	return i.ToLiveEventTranscriptionOutputWithContext(context.Background())
}

func (i LiveEventTranscriptionArgs) ToLiveEventTranscriptionOutputWithContext(ctx context.Context) LiveEventTranscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventTranscriptionOutput)
}





type LiveEventTranscriptionArrayInput interface {
	pulumi.Input

	ToLiveEventTranscriptionArrayOutput() LiveEventTranscriptionArrayOutput
	ToLiveEventTranscriptionArrayOutputWithContext(context.Context) LiveEventTranscriptionArrayOutput
}

type LiveEventTranscriptionArray []LiveEventTranscriptionInput

func (LiveEventTranscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventTranscription)(nil)).Elem()
}

func (i LiveEventTranscriptionArray) ToLiveEventTranscriptionArrayOutput() LiveEventTranscriptionArrayOutput {
	return i.ToLiveEventTranscriptionArrayOutputWithContext(context.Background())
}

func (i LiveEventTranscriptionArray) ToLiveEventTranscriptionArrayOutputWithContext(ctx context.Context) LiveEventTranscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventTranscriptionArrayOutput)
}

type LiveEventTranscriptionOutput struct{ *pulumi.OutputState }

func (LiveEventTranscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventTranscription)(nil)).Elem()
}

func (o LiveEventTranscriptionOutput) ToLiveEventTranscriptionOutput() LiveEventTranscriptionOutput {
	return o
}

func (o LiveEventTranscriptionOutput) ToLiveEventTranscriptionOutputWithContext(ctx context.Context) LiveEventTranscriptionOutput {
	return o
}

func (o LiveEventTranscriptionOutput) InputTrackSelection() LiveEventInputTrackSelectionArrayOutput {
	return o.ApplyT(func(v LiveEventTranscription) []LiveEventInputTrackSelection { return v.InputTrackSelection }).(LiveEventInputTrackSelectionArrayOutput)
}

func (o LiveEventTranscriptionOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventTranscription) *string { return v.Language }).(pulumi.StringPtrOutput)
}

func (o LiveEventTranscriptionOutput) OutputTranscriptionTrack() LiveEventOutputTranscriptionTrackPtrOutput {
	return o.ApplyT(func(v LiveEventTranscription) *LiveEventOutputTranscriptionTrack { return v.OutputTranscriptionTrack }).(LiveEventOutputTranscriptionTrackPtrOutput)
}

type LiveEventTranscriptionArrayOutput struct{ *pulumi.OutputState }

func (LiveEventTranscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventTranscription)(nil)).Elem()
}

func (o LiveEventTranscriptionArrayOutput) ToLiveEventTranscriptionArrayOutput() LiveEventTranscriptionArrayOutput {
	return o
}

func (o LiveEventTranscriptionArrayOutput) ToLiveEventTranscriptionArrayOutputWithContext(ctx context.Context) LiveEventTranscriptionArrayOutput {
	return o
}

func (o LiveEventTranscriptionArrayOutput) Index(i pulumi.IntInput) LiveEventTranscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventTranscription {
		return vs[0].([]LiveEventTranscription)[vs[1].(int)]
	}).(LiveEventTranscriptionOutput)
}

type LiveEventTranscriptionResponse struct {
	InputTrackSelection      []LiveEventInputTrackSelectionResponse     `pulumi:"inputTrackSelection"`
	Language                 *string                                    `pulumi:"language"`
	OutputTranscriptionTrack *LiveEventOutputTranscriptionTrackResponse `pulumi:"outputTranscriptionTrack"`
}





type LiveEventTranscriptionResponseInput interface {
	pulumi.Input

	ToLiveEventTranscriptionResponseOutput() LiveEventTranscriptionResponseOutput
	ToLiveEventTranscriptionResponseOutputWithContext(context.Context) LiveEventTranscriptionResponseOutput
}

type LiveEventTranscriptionResponseArgs struct {
	InputTrackSelection      LiveEventInputTrackSelectionResponseArrayInput    `pulumi:"inputTrackSelection"`
	Language                 pulumi.StringPtrInput                             `pulumi:"language"`
	OutputTranscriptionTrack LiveEventOutputTranscriptionTrackResponsePtrInput `pulumi:"outputTranscriptionTrack"`
}

func (LiveEventTranscriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventTranscriptionResponse)(nil)).Elem()
}

func (i LiveEventTranscriptionResponseArgs) ToLiveEventTranscriptionResponseOutput() LiveEventTranscriptionResponseOutput {
	return i.ToLiveEventTranscriptionResponseOutputWithContext(context.Background())
}

func (i LiveEventTranscriptionResponseArgs) ToLiveEventTranscriptionResponseOutputWithContext(ctx context.Context) LiveEventTranscriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventTranscriptionResponseOutput)
}





type LiveEventTranscriptionResponseArrayInput interface {
	pulumi.Input

	ToLiveEventTranscriptionResponseArrayOutput() LiveEventTranscriptionResponseArrayOutput
	ToLiveEventTranscriptionResponseArrayOutputWithContext(context.Context) LiveEventTranscriptionResponseArrayOutput
}

type LiveEventTranscriptionResponseArray []LiveEventTranscriptionResponseInput

func (LiveEventTranscriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventTranscriptionResponse)(nil)).Elem()
}

func (i LiveEventTranscriptionResponseArray) ToLiveEventTranscriptionResponseArrayOutput() LiveEventTranscriptionResponseArrayOutput {
	return i.ToLiveEventTranscriptionResponseArrayOutputWithContext(context.Background())
}

func (i LiveEventTranscriptionResponseArray) ToLiveEventTranscriptionResponseArrayOutputWithContext(ctx context.Context) LiveEventTranscriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventTranscriptionResponseArrayOutput)
}

type LiveEventTranscriptionResponseOutput struct{ *pulumi.OutputState }

func (LiveEventTranscriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventTranscriptionResponse)(nil)).Elem()
}

func (o LiveEventTranscriptionResponseOutput) ToLiveEventTranscriptionResponseOutput() LiveEventTranscriptionResponseOutput {
	return o
}

func (o LiveEventTranscriptionResponseOutput) ToLiveEventTranscriptionResponseOutputWithContext(ctx context.Context) LiveEventTranscriptionResponseOutput {
	return o
}

func (o LiveEventTranscriptionResponseOutput) InputTrackSelection() LiveEventInputTrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v LiveEventTranscriptionResponse) []LiveEventInputTrackSelectionResponse {
		return v.InputTrackSelection
	}).(LiveEventInputTrackSelectionResponseArrayOutput)
}

func (o LiveEventTranscriptionResponseOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventTranscriptionResponse) *string { return v.Language }).(pulumi.StringPtrOutput)
}

func (o LiveEventTranscriptionResponseOutput) OutputTranscriptionTrack() LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return o.ApplyT(func(v LiveEventTranscriptionResponse) *LiveEventOutputTranscriptionTrackResponse {
		return v.OutputTranscriptionTrack
	}).(LiveEventOutputTranscriptionTrackResponsePtrOutput)
}

type LiveEventTranscriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (LiveEventTranscriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventTranscriptionResponse)(nil)).Elem()
}

func (o LiveEventTranscriptionResponseArrayOutput) ToLiveEventTranscriptionResponseArrayOutput() LiveEventTranscriptionResponseArrayOutput {
	return o
}

func (o LiveEventTranscriptionResponseArrayOutput) ToLiveEventTranscriptionResponseArrayOutputWithContext(ctx context.Context) LiveEventTranscriptionResponseArrayOutput {
	return o
}

func (o LiveEventTranscriptionResponseArrayOutput) Index(i pulumi.IntInput) LiveEventTranscriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventTranscriptionResponse {
		return vs[0].([]LiveEventTranscriptionResponse)[vs[1].(int)]
	}).(LiveEventTranscriptionResponseOutput)
}

type StreamingEndpointAccessControl struct {
	Akamai *AkamaiAccessControl `pulumi:"akamai"`
	Ip     *IPAccessControl     `pulumi:"ip"`
}





type StreamingEndpointAccessControlInput interface {
	pulumi.Input

	ToStreamingEndpointAccessControlOutput() StreamingEndpointAccessControlOutput
	ToStreamingEndpointAccessControlOutputWithContext(context.Context) StreamingEndpointAccessControlOutput
}

type StreamingEndpointAccessControlArgs struct {
	Akamai AkamaiAccessControlPtrInput `pulumi:"akamai"`
	Ip     IPAccessControlPtrInput     `pulumi:"ip"`
}

func (StreamingEndpointAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpointAccessControl)(nil)).Elem()
}

func (i StreamingEndpointAccessControlArgs) ToStreamingEndpointAccessControlOutput() StreamingEndpointAccessControlOutput {
	return i.ToStreamingEndpointAccessControlOutputWithContext(context.Background())
}

func (i StreamingEndpointAccessControlArgs) ToStreamingEndpointAccessControlOutputWithContext(ctx context.Context) StreamingEndpointAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlOutput)
}

func (i StreamingEndpointAccessControlArgs) ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput {
	return i.ToStreamingEndpointAccessControlPtrOutputWithContext(context.Background())
}

func (i StreamingEndpointAccessControlArgs) ToStreamingEndpointAccessControlPtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlOutput).ToStreamingEndpointAccessControlPtrOutputWithContext(ctx)
}









type StreamingEndpointAccessControlPtrInput interface {
	pulumi.Input

	ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput
	ToStreamingEndpointAccessControlPtrOutputWithContext(context.Context) StreamingEndpointAccessControlPtrOutput
}

type streamingEndpointAccessControlPtrType StreamingEndpointAccessControlArgs

func StreamingEndpointAccessControlPtr(v *StreamingEndpointAccessControlArgs) StreamingEndpointAccessControlPtrInput {
	return (*streamingEndpointAccessControlPtrType)(v)
}

func (*streamingEndpointAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpointAccessControl)(nil)).Elem()
}

func (i *streamingEndpointAccessControlPtrType) ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput {
	return i.ToStreamingEndpointAccessControlPtrOutputWithContext(context.Background())
}

func (i *streamingEndpointAccessControlPtrType) ToStreamingEndpointAccessControlPtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlPtrOutput)
}

type StreamingEndpointAccessControlOutput struct{ *pulumi.OutputState }

func (StreamingEndpointAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpointAccessControl)(nil)).Elem()
}

func (o StreamingEndpointAccessControlOutput) ToStreamingEndpointAccessControlOutput() StreamingEndpointAccessControlOutput {
	return o
}

func (o StreamingEndpointAccessControlOutput) ToStreamingEndpointAccessControlOutputWithContext(ctx context.Context) StreamingEndpointAccessControlOutput {
	return o
}

func (o StreamingEndpointAccessControlOutput) ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput {
	return o.ToStreamingEndpointAccessControlPtrOutputWithContext(context.Background())
}

func (o StreamingEndpointAccessControlOutput) ToStreamingEndpointAccessControlPtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingEndpointAccessControl) *StreamingEndpointAccessControl {
		return &v
	}).(StreamingEndpointAccessControlPtrOutput)
}

func (o StreamingEndpointAccessControlOutput) Akamai() AkamaiAccessControlPtrOutput {
	return o.ApplyT(func(v StreamingEndpointAccessControl) *AkamaiAccessControl { return v.Akamai }).(AkamaiAccessControlPtrOutput)
}

func (o StreamingEndpointAccessControlOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v StreamingEndpointAccessControl) *IPAccessControl { return v.Ip }).(IPAccessControlPtrOutput)
}

type StreamingEndpointAccessControlPtrOutput struct{ *pulumi.OutputState }

func (StreamingEndpointAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpointAccessControl)(nil)).Elem()
}

func (o StreamingEndpointAccessControlPtrOutput) ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput {
	return o
}

func (o StreamingEndpointAccessControlPtrOutput) ToStreamingEndpointAccessControlPtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlPtrOutput {
	return o
}

func (o StreamingEndpointAccessControlPtrOutput) Elem() StreamingEndpointAccessControlOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControl) StreamingEndpointAccessControl {
		if v != nil {
			return *v
		}
		var ret StreamingEndpointAccessControl
		return ret
	}).(StreamingEndpointAccessControlOutput)
}

func (o StreamingEndpointAccessControlPtrOutput) Akamai() AkamaiAccessControlPtrOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControl) *AkamaiAccessControl {
		if v == nil {
			return nil
		}
		return v.Akamai
	}).(AkamaiAccessControlPtrOutput)
}

func (o StreamingEndpointAccessControlPtrOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControl) *IPAccessControl {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlPtrOutput)
}

type StreamingEndpointAccessControlResponse struct {
	Akamai *AkamaiAccessControlResponse `pulumi:"akamai"`
	Ip     *IPAccessControlResponse     `pulumi:"ip"`
}





type StreamingEndpointAccessControlResponseInput interface {
	pulumi.Input

	ToStreamingEndpointAccessControlResponseOutput() StreamingEndpointAccessControlResponseOutput
	ToStreamingEndpointAccessControlResponseOutputWithContext(context.Context) StreamingEndpointAccessControlResponseOutput
}

type StreamingEndpointAccessControlResponseArgs struct {
	Akamai AkamaiAccessControlResponsePtrInput `pulumi:"akamai"`
	Ip     IPAccessControlResponsePtrInput     `pulumi:"ip"`
}

func (StreamingEndpointAccessControlResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpointAccessControlResponse)(nil)).Elem()
}

func (i StreamingEndpointAccessControlResponseArgs) ToStreamingEndpointAccessControlResponseOutput() StreamingEndpointAccessControlResponseOutput {
	return i.ToStreamingEndpointAccessControlResponseOutputWithContext(context.Background())
}

func (i StreamingEndpointAccessControlResponseArgs) ToStreamingEndpointAccessControlResponseOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlResponseOutput)
}

func (i StreamingEndpointAccessControlResponseArgs) ToStreamingEndpointAccessControlResponsePtrOutput() StreamingEndpointAccessControlResponsePtrOutput {
	return i.ToStreamingEndpointAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i StreamingEndpointAccessControlResponseArgs) ToStreamingEndpointAccessControlResponsePtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlResponseOutput).ToStreamingEndpointAccessControlResponsePtrOutputWithContext(ctx)
}









type StreamingEndpointAccessControlResponsePtrInput interface {
	pulumi.Input

	ToStreamingEndpointAccessControlResponsePtrOutput() StreamingEndpointAccessControlResponsePtrOutput
	ToStreamingEndpointAccessControlResponsePtrOutputWithContext(context.Context) StreamingEndpointAccessControlResponsePtrOutput
}

type streamingEndpointAccessControlResponsePtrType StreamingEndpointAccessControlResponseArgs

func StreamingEndpointAccessControlResponsePtr(v *StreamingEndpointAccessControlResponseArgs) StreamingEndpointAccessControlResponsePtrInput {
	return (*streamingEndpointAccessControlResponsePtrType)(v)
}

func (*streamingEndpointAccessControlResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpointAccessControlResponse)(nil)).Elem()
}

func (i *streamingEndpointAccessControlResponsePtrType) ToStreamingEndpointAccessControlResponsePtrOutput() StreamingEndpointAccessControlResponsePtrOutput {
	return i.ToStreamingEndpointAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i *streamingEndpointAccessControlResponsePtrType) ToStreamingEndpointAccessControlResponsePtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlResponsePtrOutput)
}

type StreamingEndpointAccessControlResponseOutput struct{ *pulumi.OutputState }

func (StreamingEndpointAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpointAccessControlResponse)(nil)).Elem()
}

func (o StreamingEndpointAccessControlResponseOutput) ToStreamingEndpointAccessControlResponseOutput() StreamingEndpointAccessControlResponseOutput {
	return o
}

func (o StreamingEndpointAccessControlResponseOutput) ToStreamingEndpointAccessControlResponseOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponseOutput {
	return o
}

func (o StreamingEndpointAccessControlResponseOutput) ToStreamingEndpointAccessControlResponsePtrOutput() StreamingEndpointAccessControlResponsePtrOutput {
	return o.ToStreamingEndpointAccessControlResponsePtrOutputWithContext(context.Background())
}

func (o StreamingEndpointAccessControlResponseOutput) ToStreamingEndpointAccessControlResponsePtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingEndpointAccessControlResponse) *StreamingEndpointAccessControlResponse {
		return &v
	}).(StreamingEndpointAccessControlResponsePtrOutput)
}

func (o StreamingEndpointAccessControlResponseOutput) Akamai() AkamaiAccessControlResponsePtrOutput {
	return o.ApplyT(func(v StreamingEndpointAccessControlResponse) *AkamaiAccessControlResponse { return v.Akamai }).(AkamaiAccessControlResponsePtrOutput)
}

func (o StreamingEndpointAccessControlResponseOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v StreamingEndpointAccessControlResponse) *IPAccessControlResponse { return v.Ip }).(IPAccessControlResponsePtrOutput)
}

type StreamingEndpointAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (StreamingEndpointAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpointAccessControlResponse)(nil)).Elem()
}

func (o StreamingEndpointAccessControlResponsePtrOutput) ToStreamingEndpointAccessControlResponsePtrOutput() StreamingEndpointAccessControlResponsePtrOutput {
	return o
}

func (o StreamingEndpointAccessControlResponsePtrOutput) ToStreamingEndpointAccessControlResponsePtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponsePtrOutput {
	return o
}

func (o StreamingEndpointAccessControlResponsePtrOutput) Elem() StreamingEndpointAccessControlResponseOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControlResponse) StreamingEndpointAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret StreamingEndpointAccessControlResponse
		return ret
	}).(StreamingEndpointAccessControlResponseOutput)
}

func (o StreamingEndpointAccessControlResponsePtrOutput) Akamai() AkamaiAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControlResponse) *AkamaiAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Akamai
	}).(AkamaiAccessControlResponsePtrOutput)
}

func (o StreamingEndpointAccessControlResponsePtrOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControlResponse) *IPAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlResponsePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiAccessControlInput)(nil)).Elem(), AkamaiAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiAccessControlPtrInput)(nil)).Elem(), AkamaiAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiAccessControlResponseInput)(nil)).Elem(), AkamaiAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiAccessControlResponsePtrInput)(nil)).Elem(), AkamaiAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyInput)(nil)).Elem(), AkamaiSignatureHeaderAuthenticationKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyArrayInput)(nil)).Elem(), AkamaiSignatureHeaderAuthenticationKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyResponseInput)(nil)).Elem(), AkamaiSignatureHeaderAuthenticationKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyResponseArrayInput)(nil)).Elem(), AkamaiSignatureHeaderAuthenticationKeyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CrossSiteAccessPoliciesInput)(nil)).Elem(), CrossSiteAccessPoliciesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CrossSiteAccessPoliciesPtrInput)(nil)).Elem(), CrossSiteAccessPoliciesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CrossSiteAccessPoliciesResponseInput)(nil)).Elem(), CrossSiteAccessPoliciesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CrossSiteAccessPoliciesResponsePtrInput)(nil)).Elem(), CrossSiteAccessPoliciesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HlsInput)(nil)).Elem(), HlsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HlsPtrInput)(nil)).Elem(), HlsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HlsResponseInput)(nil)).Elem(), HlsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HlsResponsePtrInput)(nil)).Elem(), HlsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPAccessControlInput)(nil)).Elem(), IPAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPAccessControlPtrInput)(nil)).Elem(), IPAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPAccessControlResponseInput)(nil)).Elem(), IPAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPAccessControlResponsePtrInput)(nil)).Elem(), IPAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPRangeInput)(nil)).Elem(), IPRangeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPRangeArrayInput)(nil)).Elem(), IPRangeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPRangeResponseInput)(nil)).Elem(), IPRangeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPRangeResponseArrayInput)(nil)).Elem(), IPRangeResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEncodingInput)(nil)).Elem(), LiveEventEncodingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEncodingPtrInput)(nil)).Elem(), LiveEventEncodingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEncodingResponseInput)(nil)).Elem(), LiveEventEncodingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEncodingResponsePtrInput)(nil)).Elem(), LiveEventEncodingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEndpointInput)(nil)).Elem(), LiveEventEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEndpointArrayInput)(nil)).Elem(), LiveEventEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEndpointResponseInput)(nil)).Elem(), LiveEventEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEndpointResponseArrayInput)(nil)).Elem(), LiveEventEndpointResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputTypeInput)(nil)).Elem(), LiveEventInputTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputTypePtrInput)(nil)).Elem(), LiveEventInputTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputAccessControlInput)(nil)).Elem(), LiveEventInputAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputAccessControlPtrInput)(nil)).Elem(), LiveEventInputAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputAccessControlResponseInput)(nil)).Elem(), LiveEventInputAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputAccessControlResponsePtrInput)(nil)).Elem(), LiveEventInputAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputResponseInput)(nil)).Elem(), LiveEventInputResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputResponsePtrInput)(nil)).Elem(), LiveEventInputResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputTrackSelectionInput)(nil)).Elem(), LiveEventInputTrackSelectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputTrackSelectionArrayInput)(nil)).Elem(), LiveEventInputTrackSelectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputTrackSelectionResponseInput)(nil)).Elem(), LiveEventInputTrackSelectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputTrackSelectionResponseArrayInput)(nil)).Elem(), LiveEventInputTrackSelectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventOutputTranscriptionTrackInput)(nil)).Elem(), LiveEventOutputTranscriptionTrackArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventOutputTranscriptionTrackPtrInput)(nil)).Elem(), LiveEventOutputTranscriptionTrackArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventOutputTranscriptionTrackResponseInput)(nil)).Elem(), LiveEventOutputTranscriptionTrackResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventOutputTranscriptionTrackResponsePtrInput)(nil)).Elem(), LiveEventOutputTranscriptionTrackResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewInput)(nil)).Elem(), LiveEventPreviewArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewPtrInput)(nil)).Elem(), LiveEventPreviewArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewAccessControlInput)(nil)).Elem(), LiveEventPreviewAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewAccessControlPtrInput)(nil)).Elem(), LiveEventPreviewAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewAccessControlResponseInput)(nil)).Elem(), LiveEventPreviewAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewAccessControlResponsePtrInput)(nil)).Elem(), LiveEventPreviewAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewResponseInput)(nil)).Elem(), LiveEventPreviewResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewResponsePtrInput)(nil)).Elem(), LiveEventPreviewResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventTranscriptionInput)(nil)).Elem(), LiveEventTranscriptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventTranscriptionArrayInput)(nil)).Elem(), LiveEventTranscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventTranscriptionResponseInput)(nil)).Elem(), LiveEventTranscriptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventTranscriptionResponseArrayInput)(nil)).Elem(), LiveEventTranscriptionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingEndpointAccessControlInput)(nil)).Elem(), StreamingEndpointAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingEndpointAccessControlPtrInput)(nil)).Elem(), StreamingEndpointAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingEndpointAccessControlResponseInput)(nil)).Elem(), StreamingEndpointAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingEndpointAccessControlResponsePtrInput)(nil)).Elem(), StreamingEndpointAccessControlResponseArgs{})
	pulumi.RegisterOutputType(AkamaiAccessControlOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlPtrOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlResponseOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyArrayOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyResponseOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(CrossSiteAccessPoliciesOutput{})
	pulumi.RegisterOutputType(CrossSiteAccessPoliciesPtrOutput{})
	pulumi.RegisterOutputType(CrossSiteAccessPoliciesResponseOutput{})
	pulumi.RegisterOutputType(CrossSiteAccessPoliciesResponsePtrOutput{})
	pulumi.RegisterOutputType(HlsOutput{})
	pulumi.RegisterOutputType(HlsPtrOutput{})
	pulumi.RegisterOutputType(HlsResponseOutput{})
	pulumi.RegisterOutputType(HlsResponsePtrOutput{})
	pulumi.RegisterOutputType(IPAccessControlOutput{})
	pulumi.RegisterOutputType(IPAccessControlPtrOutput{})
	pulumi.RegisterOutputType(IPAccessControlResponseOutput{})
	pulumi.RegisterOutputType(IPAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(IPRangeOutput{})
	pulumi.RegisterOutputType(IPRangeArrayOutput{})
	pulumi.RegisterOutputType(IPRangeResponseOutput{})
	pulumi.RegisterOutputType(IPRangeResponseArrayOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingPtrOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingResponseOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventEndpointOutput{})
	pulumi.RegisterOutputType(LiveEventEndpointArrayOutput{})
	pulumi.RegisterOutputType(LiveEventEndpointResponseOutput{})
	pulumi.RegisterOutputType(LiveEventEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(LiveEventInputTypeOutput{})
	pulumi.RegisterOutputType(LiveEventInputTypePtrOutput{})
	pulumi.RegisterOutputType(LiveEventInputAccessControlOutput{})
	pulumi.RegisterOutputType(LiveEventInputAccessControlPtrOutput{})
	pulumi.RegisterOutputType(LiveEventInputAccessControlResponseOutput{})
	pulumi.RegisterOutputType(LiveEventInputAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventInputResponseOutput{})
	pulumi.RegisterOutputType(LiveEventInputResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventInputTrackSelectionOutput{})
	pulumi.RegisterOutputType(LiveEventInputTrackSelectionArrayOutput{})
	pulumi.RegisterOutputType(LiveEventInputTrackSelectionResponseOutput{})
	pulumi.RegisterOutputType(LiveEventInputTrackSelectionResponseArrayOutput{})
	pulumi.RegisterOutputType(LiveEventOutputTranscriptionTrackOutput{})
	pulumi.RegisterOutputType(LiveEventOutputTranscriptionTrackPtrOutput{})
	pulumi.RegisterOutputType(LiveEventOutputTranscriptionTrackResponseOutput{})
	pulumi.RegisterOutputType(LiveEventOutputTranscriptionTrackResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewPtrOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewAccessControlOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewAccessControlPtrOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewAccessControlResponseOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewResponseOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventTranscriptionOutput{})
	pulumi.RegisterOutputType(LiveEventTranscriptionArrayOutput{})
	pulumi.RegisterOutputType(LiveEventTranscriptionResponseOutput{})
	pulumi.RegisterOutputType(LiveEventTranscriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(StreamingEndpointAccessControlOutput{})
	pulumi.RegisterOutputType(StreamingEndpointAccessControlPtrOutput{})
	pulumi.RegisterOutputType(StreamingEndpointAccessControlResponseOutput{})
	pulumi.RegisterOutputType(StreamingEndpointAccessControlResponsePtrOutput{})
}

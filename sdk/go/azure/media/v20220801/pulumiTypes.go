


package v20220801

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

type ArmStreamingEndpointCurrentSku struct {
	Capacity *int `pulumi:"capacity"`
}





type ArmStreamingEndpointCurrentSkuInput interface {
	pulumi.Input

	ToArmStreamingEndpointCurrentSkuOutput() ArmStreamingEndpointCurrentSkuOutput
	ToArmStreamingEndpointCurrentSkuOutputWithContext(context.Context) ArmStreamingEndpointCurrentSkuOutput
}

type ArmStreamingEndpointCurrentSkuArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
}

func (ArmStreamingEndpointCurrentSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmStreamingEndpointCurrentSku)(nil)).Elem()
}

func (i ArmStreamingEndpointCurrentSkuArgs) ToArmStreamingEndpointCurrentSkuOutput() ArmStreamingEndpointCurrentSkuOutput {
	return i.ToArmStreamingEndpointCurrentSkuOutputWithContext(context.Background())
}

func (i ArmStreamingEndpointCurrentSkuArgs) ToArmStreamingEndpointCurrentSkuOutputWithContext(ctx context.Context) ArmStreamingEndpointCurrentSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmStreamingEndpointCurrentSkuOutput)
}

func (i ArmStreamingEndpointCurrentSkuArgs) ToArmStreamingEndpointCurrentSkuPtrOutput() ArmStreamingEndpointCurrentSkuPtrOutput {
	return i.ToArmStreamingEndpointCurrentSkuPtrOutputWithContext(context.Background())
}

func (i ArmStreamingEndpointCurrentSkuArgs) ToArmStreamingEndpointCurrentSkuPtrOutputWithContext(ctx context.Context) ArmStreamingEndpointCurrentSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmStreamingEndpointCurrentSkuOutput).ToArmStreamingEndpointCurrentSkuPtrOutputWithContext(ctx)
}









type ArmStreamingEndpointCurrentSkuPtrInput interface {
	pulumi.Input

	ToArmStreamingEndpointCurrentSkuPtrOutput() ArmStreamingEndpointCurrentSkuPtrOutput
	ToArmStreamingEndpointCurrentSkuPtrOutputWithContext(context.Context) ArmStreamingEndpointCurrentSkuPtrOutput
}

type armStreamingEndpointCurrentSkuPtrType ArmStreamingEndpointCurrentSkuArgs

func ArmStreamingEndpointCurrentSkuPtr(v *ArmStreamingEndpointCurrentSkuArgs) ArmStreamingEndpointCurrentSkuPtrInput {
	return (*armStreamingEndpointCurrentSkuPtrType)(v)
}

func (*armStreamingEndpointCurrentSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmStreamingEndpointCurrentSku)(nil)).Elem()
}

func (i *armStreamingEndpointCurrentSkuPtrType) ToArmStreamingEndpointCurrentSkuPtrOutput() ArmStreamingEndpointCurrentSkuPtrOutput {
	return i.ToArmStreamingEndpointCurrentSkuPtrOutputWithContext(context.Background())
}

func (i *armStreamingEndpointCurrentSkuPtrType) ToArmStreamingEndpointCurrentSkuPtrOutputWithContext(ctx context.Context) ArmStreamingEndpointCurrentSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmStreamingEndpointCurrentSkuPtrOutput)
}

type ArmStreamingEndpointCurrentSkuOutput struct{ *pulumi.OutputState }

func (ArmStreamingEndpointCurrentSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmStreamingEndpointCurrentSku)(nil)).Elem()
}

func (o ArmStreamingEndpointCurrentSkuOutput) ToArmStreamingEndpointCurrentSkuOutput() ArmStreamingEndpointCurrentSkuOutput {
	return o
}

func (o ArmStreamingEndpointCurrentSkuOutput) ToArmStreamingEndpointCurrentSkuOutputWithContext(ctx context.Context) ArmStreamingEndpointCurrentSkuOutput {
	return o
}

func (o ArmStreamingEndpointCurrentSkuOutput) ToArmStreamingEndpointCurrentSkuPtrOutput() ArmStreamingEndpointCurrentSkuPtrOutput {
	return o.ToArmStreamingEndpointCurrentSkuPtrOutputWithContext(context.Background())
}

func (o ArmStreamingEndpointCurrentSkuOutput) ToArmStreamingEndpointCurrentSkuPtrOutputWithContext(ctx context.Context) ArmStreamingEndpointCurrentSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArmStreamingEndpointCurrentSku) *ArmStreamingEndpointCurrentSku {
		return &v
	}).(ArmStreamingEndpointCurrentSkuPtrOutput)
}

func (o ArmStreamingEndpointCurrentSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArmStreamingEndpointCurrentSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

type ArmStreamingEndpointCurrentSkuPtrOutput struct{ *pulumi.OutputState }

func (ArmStreamingEndpointCurrentSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmStreamingEndpointCurrentSku)(nil)).Elem()
}

func (o ArmStreamingEndpointCurrentSkuPtrOutput) ToArmStreamingEndpointCurrentSkuPtrOutput() ArmStreamingEndpointCurrentSkuPtrOutput {
	return o
}

func (o ArmStreamingEndpointCurrentSkuPtrOutput) ToArmStreamingEndpointCurrentSkuPtrOutputWithContext(ctx context.Context) ArmStreamingEndpointCurrentSkuPtrOutput {
	return o
}

func (o ArmStreamingEndpointCurrentSkuPtrOutput) Elem() ArmStreamingEndpointCurrentSkuOutput {
	return o.ApplyT(func(v *ArmStreamingEndpointCurrentSku) ArmStreamingEndpointCurrentSku {
		if v != nil {
			return *v
		}
		var ret ArmStreamingEndpointCurrentSku
		return ret
	}).(ArmStreamingEndpointCurrentSkuOutput)
}

func (o ArmStreamingEndpointCurrentSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArmStreamingEndpointCurrentSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

type ArmStreamingEndpointCurrentSkuResponse struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
}

type ArmStreamingEndpointCurrentSkuResponseOutput struct{ *pulumi.OutputState }

func (ArmStreamingEndpointCurrentSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmStreamingEndpointCurrentSkuResponse)(nil)).Elem()
}

func (o ArmStreamingEndpointCurrentSkuResponseOutput) ToArmStreamingEndpointCurrentSkuResponseOutput() ArmStreamingEndpointCurrentSkuResponseOutput {
	return o
}

func (o ArmStreamingEndpointCurrentSkuResponseOutput) ToArmStreamingEndpointCurrentSkuResponseOutputWithContext(ctx context.Context) ArmStreamingEndpointCurrentSkuResponseOutput {
	return o
}

func (o ArmStreamingEndpointCurrentSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArmStreamingEndpointCurrentSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ArmStreamingEndpointCurrentSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ArmStreamingEndpointCurrentSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ArmStreamingEndpointCurrentSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmStreamingEndpointCurrentSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmStreamingEndpointCurrentSkuResponse)(nil)).Elem()
}

func (o ArmStreamingEndpointCurrentSkuResponsePtrOutput) ToArmStreamingEndpointCurrentSkuResponsePtrOutput() ArmStreamingEndpointCurrentSkuResponsePtrOutput {
	return o
}

func (o ArmStreamingEndpointCurrentSkuResponsePtrOutput) ToArmStreamingEndpointCurrentSkuResponsePtrOutputWithContext(ctx context.Context) ArmStreamingEndpointCurrentSkuResponsePtrOutput {
	return o
}

func (o ArmStreamingEndpointCurrentSkuResponsePtrOutput) Elem() ArmStreamingEndpointCurrentSkuResponseOutput {
	return o.ApplyT(func(v *ArmStreamingEndpointCurrentSkuResponse) ArmStreamingEndpointCurrentSkuResponse {
		if v != nil {
			return *v
		}
		var ret ArmStreamingEndpointCurrentSkuResponse
		return ret
	}).(ArmStreamingEndpointCurrentSkuResponseOutput)
}

func (o ArmStreamingEndpointCurrentSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArmStreamingEndpointCurrentSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ArmStreamingEndpointCurrentSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmStreamingEndpointCurrentSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
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
	EncodingType     *string `pulumi:"encodingType"`
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	PresetName       *string `pulumi:"presetName"`
	StretchMode      *string `pulumi:"stretchMode"`
}





type LiveEventEncodingInput interface {
	pulumi.Input

	ToLiveEventEncodingOutput() LiveEventEncodingOutput
	ToLiveEventEncodingOutputWithContext(context.Context) LiveEventEncodingOutput
}

type LiveEventEncodingArgs struct {
	EncodingType     pulumi.StringPtrInput `pulumi:"encodingType"`
	KeyFrameInterval pulumi.StringPtrInput `pulumi:"keyFrameInterval"`
	PresetName       pulumi.StringPtrInput `pulumi:"presetName"`
	StretchMode      pulumi.StringPtrInput `pulumi:"stretchMode"`
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

func (o LiveEventEncodingOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncoding) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncoding) *string { return v.PresetName }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncoding) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
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

func (o LiveEventEncodingPtrOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncoding) *string {
		if v == nil {
			return nil
		}
		return v.KeyFrameInterval
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

func (o LiveEventEncodingPtrOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncoding) *string {
		if v == nil {
			return nil
		}
		return v.StretchMode
	}).(pulumi.StringPtrOutput)
}

type LiveEventEncodingResponse struct {
	EncodingType     *string `pulumi:"encodingType"`
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	PresetName       *string `pulumi:"presetName"`
	StretchMode      *string `pulumi:"stretchMode"`
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

func (o LiveEventEncodingResponseOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncodingResponse) *string { return v.EncodingType }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponseOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncodingResponse) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponseOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncodingResponse) *string { return v.PresetName }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponseOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncodingResponse) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
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

func (o LiveEventEncodingResponsePtrOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyFrameInterval
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

func (o LiveEventEncodingResponsePtrOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) *string {
		if v == nil {
			return nil
		}
		return v.StretchMode
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
	pulumi.RegisterOutputType(AkamaiAccessControlOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlPtrOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlResponseOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyArrayOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyResponseOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(ArmStreamingEndpointCurrentSkuOutput{})
	pulumi.RegisterOutputType(ArmStreamingEndpointCurrentSkuPtrOutput{})
	pulumi.RegisterOutputType(ArmStreamingEndpointCurrentSkuResponseOutput{})
	pulumi.RegisterOutputType(ArmStreamingEndpointCurrentSkuResponsePtrOutput{})
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
	pulumi.RegisterOutputType(LiveEventInputAccessControlOutput{})
	pulumi.RegisterOutputType(LiveEventInputAccessControlPtrOutput{})
	pulumi.RegisterOutputType(LiveEventInputAccessControlResponseOutput{})
	pulumi.RegisterOutputType(LiveEventInputAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventInputResponseOutput{})
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
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}

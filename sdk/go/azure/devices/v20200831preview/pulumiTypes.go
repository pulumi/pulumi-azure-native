


package v20200831preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ArmIdentity struct {
	IdentityType           *string                `pulumi:"identityType"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ArmIdentityInput interface {
	pulumi.Input

	ToArmIdentityOutput() ArmIdentityOutput
	ToArmIdentityOutputWithContext(context.Context) ArmIdentityOutput
}

type ArmIdentityArgs struct {
	IdentityType           pulumi.StringPtrInput `pulumi:"identityType"`
	UserAssignedIdentities pulumi.MapInput       `pulumi:"userAssignedIdentities"`
}

func (ArmIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmIdentity)(nil)).Elem()
}

func (i ArmIdentityArgs) ToArmIdentityOutput() ArmIdentityOutput {
	return i.ToArmIdentityOutputWithContext(context.Background())
}

func (i ArmIdentityArgs) ToArmIdentityOutputWithContext(ctx context.Context) ArmIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmIdentityOutput)
}

func (i ArmIdentityArgs) ToArmIdentityPtrOutput() ArmIdentityPtrOutput {
	return i.ToArmIdentityPtrOutputWithContext(context.Background())
}

func (i ArmIdentityArgs) ToArmIdentityPtrOutputWithContext(ctx context.Context) ArmIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmIdentityOutput).ToArmIdentityPtrOutputWithContext(ctx)
}









type ArmIdentityPtrInput interface {
	pulumi.Input

	ToArmIdentityPtrOutput() ArmIdentityPtrOutput
	ToArmIdentityPtrOutputWithContext(context.Context) ArmIdentityPtrOutput
}

type armIdentityPtrType ArmIdentityArgs

func ArmIdentityPtr(v *ArmIdentityArgs) ArmIdentityPtrInput {
	return (*armIdentityPtrType)(v)
}

func (*armIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmIdentity)(nil)).Elem()
}

func (i *armIdentityPtrType) ToArmIdentityPtrOutput() ArmIdentityPtrOutput {
	return i.ToArmIdentityPtrOutputWithContext(context.Background())
}

func (i *armIdentityPtrType) ToArmIdentityPtrOutputWithContext(ctx context.Context) ArmIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmIdentityPtrOutput)
}

type ArmIdentityOutput struct{ *pulumi.OutputState }

func (ArmIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmIdentity)(nil)).Elem()
}

func (o ArmIdentityOutput) ToArmIdentityOutput() ArmIdentityOutput {
	return o
}

func (o ArmIdentityOutput) ToArmIdentityOutputWithContext(ctx context.Context) ArmIdentityOutput {
	return o
}

func (o ArmIdentityOutput) ToArmIdentityPtrOutput() ArmIdentityPtrOutput {
	return o.ToArmIdentityPtrOutputWithContext(context.Background())
}

func (o ArmIdentityOutput) ToArmIdentityPtrOutputWithContext(ctx context.Context) ArmIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArmIdentity) *ArmIdentity {
		return &v
	}).(ArmIdentityPtrOutput)
}

func (o ArmIdentityOutput) IdentityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmIdentity) *string { return v.IdentityType }).(pulumi.StringPtrOutput)
}

func (o ArmIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ArmIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ArmIdentityPtrOutput struct{ *pulumi.OutputState }

func (ArmIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmIdentity)(nil)).Elem()
}

func (o ArmIdentityPtrOutput) ToArmIdentityPtrOutput() ArmIdentityPtrOutput {
	return o
}

func (o ArmIdentityPtrOutput) ToArmIdentityPtrOutputWithContext(ctx context.Context) ArmIdentityPtrOutput {
	return o
}

func (o ArmIdentityPtrOutput) Elem() ArmIdentityOutput {
	return o.ApplyT(func(v *ArmIdentity) ArmIdentity {
		if v != nil {
			return *v
		}
		var ret ArmIdentity
		return ret
	}).(ArmIdentityOutput)
}

func (o ArmIdentityPtrOutput) IdentityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmIdentity) *string {
		if v == nil {
			return nil
		}
		return v.IdentityType
	}).(pulumi.StringPtrOutput)
}

func (o ArmIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ArmIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ArmIdentityResponse struct {
	IdentityType           *string                            `pulumi:"identityType"`
	PrincipalId            string                             `pulumi:"principalId"`
	TenantId               string                             `pulumi:"tenantId"`
	UserAssignedIdentities map[string]ArmUserIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ArmIdentityResponseOutput struct{ *pulumi.OutputState }

func (ArmIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmIdentityResponse)(nil)).Elem()
}

func (o ArmIdentityResponseOutput) ToArmIdentityResponseOutput() ArmIdentityResponseOutput {
	return o
}

func (o ArmIdentityResponseOutput) ToArmIdentityResponseOutputWithContext(ctx context.Context) ArmIdentityResponseOutput {
	return o
}

func (o ArmIdentityResponseOutput) IdentityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmIdentityResponse) *string { return v.IdentityType }).(pulumi.StringPtrOutput)
}

func (o ArmIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ArmIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ArmIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ArmIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ArmIdentityResponseOutput) UserAssignedIdentities() ArmUserIdentityResponseMapOutput {
	return o.ApplyT(func(v ArmIdentityResponse) map[string]ArmUserIdentityResponse { return v.UserAssignedIdentities }).(ArmUserIdentityResponseMapOutput)
}

type ArmIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmIdentityResponse)(nil)).Elem()
}

func (o ArmIdentityResponsePtrOutput) ToArmIdentityResponsePtrOutput() ArmIdentityResponsePtrOutput {
	return o
}

func (o ArmIdentityResponsePtrOutput) ToArmIdentityResponsePtrOutputWithContext(ctx context.Context) ArmIdentityResponsePtrOutput {
	return o
}

func (o ArmIdentityResponsePtrOutput) Elem() ArmIdentityResponseOutput {
	return o.ApplyT(func(v *ArmIdentityResponse) ArmIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ArmIdentityResponse
		return ret
	}).(ArmIdentityResponseOutput)
}

func (o ArmIdentityResponsePtrOutput) IdentityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.IdentityType
	}).(pulumi.StringPtrOutput)
}

func (o ArmIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ArmIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ArmIdentityResponsePtrOutput) UserAssignedIdentities() ArmUserIdentityResponseMapOutput {
	return o.ApplyT(func(v *ArmIdentityResponse) map[string]ArmUserIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(ArmUserIdentityResponseMapOutput)
}

type ArmUserIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type ArmUserIdentityResponseOutput struct{ *pulumi.OutputState }

func (ArmUserIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmUserIdentityResponse)(nil)).Elem()
}

func (o ArmUserIdentityResponseOutput) ToArmUserIdentityResponseOutput() ArmUserIdentityResponseOutput {
	return o
}

func (o ArmUserIdentityResponseOutput) ToArmUserIdentityResponseOutputWithContext(ctx context.Context) ArmUserIdentityResponseOutput {
	return o
}

func (o ArmUserIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ArmUserIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ArmUserIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ArmUserIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ArmUserIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (ArmUserIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ArmUserIdentityResponse)(nil)).Elem()
}

func (o ArmUserIdentityResponseMapOutput) ToArmUserIdentityResponseMapOutput() ArmUserIdentityResponseMapOutput {
	return o
}

func (o ArmUserIdentityResponseMapOutput) ToArmUserIdentityResponseMapOutputWithContext(ctx context.Context) ArmUserIdentityResponseMapOutput {
	return o
}

func (o ArmUserIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) ArmUserIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ArmUserIdentityResponse {
		return vs[0].(map[string]ArmUserIdentityResponse)[vs[1].(string)]
	}).(ArmUserIdentityResponseOutput)
}

type CertificateProperties struct {
	Certificate *string `pulumi:"certificate"`
}





type CertificatePropertiesInput interface {
	pulumi.Input

	ToCertificatePropertiesOutput() CertificatePropertiesOutput
	ToCertificatePropertiesOutputWithContext(context.Context) CertificatePropertiesOutput
}

type CertificatePropertiesArgs struct {
	Certificate pulumi.StringPtrInput `pulumi:"certificate"`
}

func (CertificatePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateProperties)(nil)).Elem()
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesOutput() CertificatePropertiesOutput {
	return i.ToCertificatePropertiesOutputWithContext(context.Background())
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesOutputWithContext(ctx context.Context) CertificatePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesOutput)
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return i.ToCertificatePropertiesPtrOutputWithContext(context.Background())
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesOutput).ToCertificatePropertiesPtrOutputWithContext(ctx)
}









type CertificatePropertiesPtrInput interface {
	pulumi.Input

	ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput
	ToCertificatePropertiesPtrOutputWithContext(context.Context) CertificatePropertiesPtrOutput
}

type certificatePropertiesPtrType CertificatePropertiesArgs

func CertificatePropertiesPtr(v *CertificatePropertiesArgs) CertificatePropertiesPtrInput {
	return (*certificatePropertiesPtrType)(v)
}

func (*certificatePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateProperties)(nil)).Elem()
}

func (i *certificatePropertiesPtrType) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return i.ToCertificatePropertiesPtrOutputWithContext(context.Background())
}

func (i *certificatePropertiesPtrType) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesPtrOutput)
}

type CertificatePropertiesOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateProperties)(nil)).Elem()
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesOutput() CertificatePropertiesOutput {
	return o
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesOutputWithContext(ctx context.Context) CertificatePropertiesOutput {
	return o
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return o.ToCertificatePropertiesPtrOutputWithContext(context.Background())
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateProperties) *CertificateProperties {
		return &v
	}).(CertificatePropertiesPtrOutput)
}

func (o CertificatePropertiesOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateProperties) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

type CertificatePropertiesPtrOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateProperties)(nil)).Elem()
}

func (o CertificatePropertiesPtrOutput) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return o
}

func (o CertificatePropertiesPtrOutput) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return o
}

func (o CertificatePropertiesPtrOutput) Elem() CertificatePropertiesOutput {
	return o.ApplyT(func(v *CertificateProperties) CertificateProperties {
		if v != nil {
			return *v
		}
		var ret CertificateProperties
		return ret
	}).(CertificatePropertiesOutput)
}

func (o CertificatePropertiesPtrOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateProperties) *string {
		if v == nil {
			return nil
		}
		return v.Certificate
	}).(pulumi.StringPtrOutput)
}

type CertificatePropertiesResponse struct {
	Certificate *string `pulumi:"certificate"`
	Created     string  `pulumi:"created"`
	Expiry      string  `pulumi:"expiry"`
	IsVerified  bool    `pulumi:"isVerified"`
	Subject     string  `pulumi:"subject"`
	Thumbprint  string  `pulumi:"thumbprint"`
	Updated     string  `pulumi:"updated"`
}

type CertificatePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificatePropertiesResponse)(nil)).Elem()
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput {
	return o
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponseOutputWithContext(ctx context.Context) CertificatePropertiesResponseOutput {
	return o
}

func (o CertificatePropertiesResponseOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponseOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Created }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Expiry() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Expiry }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) IsVerified() pulumi.BoolOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) bool { return v.IsVerified }).(pulumi.BoolOutput)
}

func (o CertificatePropertiesResponseOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Subject }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Updated }).(pulumi.StringOutput)
}

type CloudToDeviceProperties struct {
	DefaultTtlAsIso8601 *string             `pulumi:"defaultTtlAsIso8601"`
	Feedback            *FeedbackProperties `pulumi:"feedback"`
	MaxDeliveryCount    *int                `pulumi:"maxDeliveryCount"`
}





type CloudToDevicePropertiesInput interface {
	pulumi.Input

	ToCloudToDevicePropertiesOutput() CloudToDevicePropertiesOutput
	ToCloudToDevicePropertiesOutputWithContext(context.Context) CloudToDevicePropertiesOutput
}

type CloudToDevicePropertiesArgs struct {
	DefaultTtlAsIso8601 pulumi.StringPtrInput      `pulumi:"defaultTtlAsIso8601"`
	Feedback            FeedbackPropertiesPtrInput `pulumi:"feedback"`
	MaxDeliveryCount    pulumi.IntPtrInput         `pulumi:"maxDeliveryCount"`
}

func (CloudToDevicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudToDeviceProperties)(nil)).Elem()
}

func (i CloudToDevicePropertiesArgs) ToCloudToDevicePropertiesOutput() CloudToDevicePropertiesOutput {
	return i.ToCloudToDevicePropertiesOutputWithContext(context.Background())
}

func (i CloudToDevicePropertiesArgs) ToCloudToDevicePropertiesOutputWithContext(ctx context.Context) CloudToDevicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudToDevicePropertiesOutput)
}

func (i CloudToDevicePropertiesArgs) ToCloudToDevicePropertiesPtrOutput() CloudToDevicePropertiesPtrOutput {
	return i.ToCloudToDevicePropertiesPtrOutputWithContext(context.Background())
}

func (i CloudToDevicePropertiesArgs) ToCloudToDevicePropertiesPtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudToDevicePropertiesOutput).ToCloudToDevicePropertiesPtrOutputWithContext(ctx)
}









type CloudToDevicePropertiesPtrInput interface {
	pulumi.Input

	ToCloudToDevicePropertiesPtrOutput() CloudToDevicePropertiesPtrOutput
	ToCloudToDevicePropertiesPtrOutputWithContext(context.Context) CloudToDevicePropertiesPtrOutput
}

type cloudToDevicePropertiesPtrType CloudToDevicePropertiesArgs

func CloudToDevicePropertiesPtr(v *CloudToDevicePropertiesArgs) CloudToDevicePropertiesPtrInput {
	return (*cloudToDevicePropertiesPtrType)(v)
}

func (*cloudToDevicePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudToDeviceProperties)(nil)).Elem()
}

func (i *cloudToDevicePropertiesPtrType) ToCloudToDevicePropertiesPtrOutput() CloudToDevicePropertiesPtrOutput {
	return i.ToCloudToDevicePropertiesPtrOutputWithContext(context.Background())
}

func (i *cloudToDevicePropertiesPtrType) ToCloudToDevicePropertiesPtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudToDevicePropertiesPtrOutput)
}

type CloudToDevicePropertiesOutput struct{ *pulumi.OutputState }

func (CloudToDevicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudToDeviceProperties)(nil)).Elem()
}

func (o CloudToDevicePropertiesOutput) ToCloudToDevicePropertiesOutput() CloudToDevicePropertiesOutput {
	return o
}

func (o CloudToDevicePropertiesOutput) ToCloudToDevicePropertiesOutputWithContext(ctx context.Context) CloudToDevicePropertiesOutput {
	return o
}

func (o CloudToDevicePropertiesOutput) ToCloudToDevicePropertiesPtrOutput() CloudToDevicePropertiesPtrOutput {
	return o.ToCloudToDevicePropertiesPtrOutputWithContext(context.Background())
}

func (o CloudToDevicePropertiesOutput) ToCloudToDevicePropertiesPtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudToDeviceProperties) *CloudToDeviceProperties {
		return &v
	}).(CloudToDevicePropertiesPtrOutput)
}

func (o CloudToDevicePropertiesOutput) DefaultTtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudToDeviceProperties) *string { return v.DefaultTtlAsIso8601 }).(pulumi.StringPtrOutput)
}

func (o CloudToDevicePropertiesOutput) Feedback() FeedbackPropertiesPtrOutput {
	return o.ApplyT(func(v CloudToDeviceProperties) *FeedbackProperties { return v.Feedback }).(FeedbackPropertiesPtrOutput)
}

func (o CloudToDevicePropertiesOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CloudToDeviceProperties) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

type CloudToDevicePropertiesPtrOutput struct{ *pulumi.OutputState }

func (CloudToDevicePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudToDeviceProperties)(nil)).Elem()
}

func (o CloudToDevicePropertiesPtrOutput) ToCloudToDevicePropertiesPtrOutput() CloudToDevicePropertiesPtrOutput {
	return o
}

func (o CloudToDevicePropertiesPtrOutput) ToCloudToDevicePropertiesPtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesPtrOutput {
	return o
}

func (o CloudToDevicePropertiesPtrOutput) Elem() CloudToDevicePropertiesOutput {
	return o.ApplyT(func(v *CloudToDeviceProperties) CloudToDeviceProperties {
		if v != nil {
			return *v
		}
		var ret CloudToDeviceProperties
		return ret
	}).(CloudToDevicePropertiesOutput)
}

func (o CloudToDevicePropertiesPtrOutput) DefaultTtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudToDeviceProperties) *string {
		if v == nil {
			return nil
		}
		return v.DefaultTtlAsIso8601
	}).(pulumi.StringPtrOutput)
}

func (o CloudToDevicePropertiesPtrOutput) Feedback() FeedbackPropertiesPtrOutput {
	return o.ApplyT(func(v *CloudToDeviceProperties) *FeedbackProperties {
		if v == nil {
			return nil
		}
		return v.Feedback
	}).(FeedbackPropertiesPtrOutput)
}

func (o CloudToDevicePropertiesPtrOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudToDeviceProperties) *int {
		if v == nil {
			return nil
		}
		return v.MaxDeliveryCount
	}).(pulumi.IntPtrOutput)
}

type CloudToDevicePropertiesResponse struct {
	DefaultTtlAsIso8601 *string                     `pulumi:"defaultTtlAsIso8601"`
	Feedback            *FeedbackPropertiesResponse `pulumi:"feedback"`
	MaxDeliveryCount    *int                        `pulumi:"maxDeliveryCount"`
}

type CloudToDevicePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CloudToDevicePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudToDevicePropertiesResponse)(nil)).Elem()
}

func (o CloudToDevicePropertiesResponseOutput) ToCloudToDevicePropertiesResponseOutput() CloudToDevicePropertiesResponseOutput {
	return o
}

func (o CloudToDevicePropertiesResponseOutput) ToCloudToDevicePropertiesResponseOutputWithContext(ctx context.Context) CloudToDevicePropertiesResponseOutput {
	return o
}

func (o CloudToDevicePropertiesResponseOutput) DefaultTtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudToDevicePropertiesResponse) *string { return v.DefaultTtlAsIso8601 }).(pulumi.StringPtrOutput)
}

func (o CloudToDevicePropertiesResponseOutput) Feedback() FeedbackPropertiesResponsePtrOutput {
	return o.ApplyT(func(v CloudToDevicePropertiesResponse) *FeedbackPropertiesResponse { return v.Feedback }).(FeedbackPropertiesResponsePtrOutput)
}

func (o CloudToDevicePropertiesResponseOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CloudToDevicePropertiesResponse) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

type CloudToDevicePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudToDevicePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudToDevicePropertiesResponse)(nil)).Elem()
}

func (o CloudToDevicePropertiesResponsePtrOutput) ToCloudToDevicePropertiesResponsePtrOutput() CloudToDevicePropertiesResponsePtrOutput {
	return o
}

func (o CloudToDevicePropertiesResponsePtrOutput) ToCloudToDevicePropertiesResponsePtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesResponsePtrOutput {
	return o
}

func (o CloudToDevicePropertiesResponsePtrOutput) Elem() CloudToDevicePropertiesResponseOutput {
	return o.ApplyT(func(v *CloudToDevicePropertiesResponse) CloudToDevicePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CloudToDevicePropertiesResponse
		return ret
	}).(CloudToDevicePropertiesResponseOutput)
}

func (o CloudToDevicePropertiesResponsePtrOutput) DefaultTtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudToDevicePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultTtlAsIso8601
	}).(pulumi.StringPtrOutput)
}

func (o CloudToDevicePropertiesResponsePtrOutput) Feedback() FeedbackPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CloudToDevicePropertiesResponse) *FeedbackPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Feedback
	}).(FeedbackPropertiesResponsePtrOutput)
}

func (o CloudToDevicePropertiesResponsePtrOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudToDevicePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxDeliveryCount
	}).(pulumi.IntPtrOutput)
}

type EncryptionPropertiesDescription struct {
	KeySource          *string                 `pulumi:"keySource"`
	KeyVaultProperties []KeyVaultKeyProperties `pulumi:"keyVaultProperties"`
}





type EncryptionPropertiesDescriptionInput interface {
	pulumi.Input

	ToEncryptionPropertiesDescriptionOutput() EncryptionPropertiesDescriptionOutput
	ToEncryptionPropertiesDescriptionOutputWithContext(context.Context) EncryptionPropertiesDescriptionOutput
}

type EncryptionPropertiesDescriptionArgs struct {
	KeySource          pulumi.StringPtrInput           `pulumi:"keySource"`
	KeyVaultProperties KeyVaultKeyPropertiesArrayInput `pulumi:"keyVaultProperties"`
}

func (EncryptionPropertiesDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesDescription)(nil)).Elem()
}

func (i EncryptionPropertiesDescriptionArgs) ToEncryptionPropertiesDescriptionOutput() EncryptionPropertiesDescriptionOutput {
	return i.ToEncryptionPropertiesDescriptionOutputWithContext(context.Background())
}

func (i EncryptionPropertiesDescriptionArgs) ToEncryptionPropertiesDescriptionOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesDescriptionOutput)
}

func (i EncryptionPropertiesDescriptionArgs) ToEncryptionPropertiesDescriptionPtrOutput() EncryptionPropertiesDescriptionPtrOutput {
	return i.ToEncryptionPropertiesDescriptionPtrOutputWithContext(context.Background())
}

func (i EncryptionPropertiesDescriptionArgs) ToEncryptionPropertiesDescriptionPtrOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesDescriptionOutput).ToEncryptionPropertiesDescriptionPtrOutputWithContext(ctx)
}









type EncryptionPropertiesDescriptionPtrInput interface {
	pulumi.Input

	ToEncryptionPropertiesDescriptionPtrOutput() EncryptionPropertiesDescriptionPtrOutput
	ToEncryptionPropertiesDescriptionPtrOutputWithContext(context.Context) EncryptionPropertiesDescriptionPtrOutput
}

type encryptionPropertiesDescriptionPtrType EncryptionPropertiesDescriptionArgs

func EncryptionPropertiesDescriptionPtr(v *EncryptionPropertiesDescriptionArgs) EncryptionPropertiesDescriptionPtrInput {
	return (*encryptionPropertiesDescriptionPtrType)(v)
}

func (*encryptionPropertiesDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesDescription)(nil)).Elem()
}

func (i *encryptionPropertiesDescriptionPtrType) ToEncryptionPropertiesDescriptionPtrOutput() EncryptionPropertiesDescriptionPtrOutput {
	return i.ToEncryptionPropertiesDescriptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPropertiesDescriptionPtrType) ToEncryptionPropertiesDescriptionPtrOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesDescriptionPtrOutput)
}

type EncryptionPropertiesDescriptionOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesDescription)(nil)).Elem()
}

func (o EncryptionPropertiesDescriptionOutput) ToEncryptionPropertiesDescriptionOutput() EncryptionPropertiesDescriptionOutput {
	return o
}

func (o EncryptionPropertiesDescriptionOutput) ToEncryptionPropertiesDescriptionOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionOutput {
	return o
}

func (o EncryptionPropertiesDescriptionOutput) ToEncryptionPropertiesDescriptionPtrOutput() EncryptionPropertiesDescriptionPtrOutput {
	return o.ToEncryptionPropertiesDescriptionPtrOutputWithContext(context.Background())
}

func (o EncryptionPropertiesDescriptionOutput) ToEncryptionPropertiesDescriptionPtrOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionPropertiesDescription) *EncryptionPropertiesDescription {
		return &v
	}).(EncryptionPropertiesDescriptionPtrOutput)
}

func (o EncryptionPropertiesDescriptionOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesDescription) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesDescriptionOutput) KeyVaultProperties() KeyVaultKeyPropertiesArrayOutput {
	return o.ApplyT(func(v EncryptionPropertiesDescription) []KeyVaultKeyProperties { return v.KeyVaultProperties }).(KeyVaultKeyPropertiesArrayOutput)
}

type EncryptionPropertiesDescriptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesDescription)(nil)).Elem()
}

func (o EncryptionPropertiesDescriptionPtrOutput) ToEncryptionPropertiesDescriptionPtrOutput() EncryptionPropertiesDescriptionPtrOutput {
	return o
}

func (o EncryptionPropertiesDescriptionPtrOutput) ToEncryptionPropertiesDescriptionPtrOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionPtrOutput {
	return o
}

func (o EncryptionPropertiesDescriptionPtrOutput) Elem() EncryptionPropertiesDescriptionOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescription) EncryptionPropertiesDescription {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertiesDescription
		return ret
	}).(EncryptionPropertiesDescriptionOutput)
}

func (o EncryptionPropertiesDescriptionPtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescription) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesDescriptionPtrOutput) KeyVaultProperties() KeyVaultKeyPropertiesArrayOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescription) []KeyVaultKeyProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultKeyPropertiesArrayOutput)
}

type EncryptionPropertiesDescriptionResponse struct {
	KeySource          *string                         `pulumi:"keySource"`
	KeyVaultProperties []KeyVaultKeyPropertiesResponse `pulumi:"keyVaultProperties"`
}

type EncryptionPropertiesDescriptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesDescriptionResponse)(nil)).Elem()
}

func (o EncryptionPropertiesDescriptionResponseOutput) ToEncryptionPropertiesDescriptionResponseOutput() EncryptionPropertiesDescriptionResponseOutput {
	return o
}

func (o EncryptionPropertiesDescriptionResponseOutput) ToEncryptionPropertiesDescriptionResponseOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionResponseOutput {
	return o
}

func (o EncryptionPropertiesDescriptionResponseOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesDescriptionResponse) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesDescriptionResponseOutput) KeyVaultProperties() KeyVaultKeyPropertiesResponseArrayOutput {
	return o.ApplyT(func(v EncryptionPropertiesDescriptionResponse) []KeyVaultKeyPropertiesResponse {
		return v.KeyVaultProperties
	}).(KeyVaultKeyPropertiesResponseArrayOutput)
}

type EncryptionPropertiesDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesDescriptionResponse)(nil)).Elem()
}

func (o EncryptionPropertiesDescriptionResponsePtrOutput) ToEncryptionPropertiesDescriptionResponsePtrOutput() EncryptionPropertiesDescriptionResponsePtrOutput {
	return o
}

func (o EncryptionPropertiesDescriptionResponsePtrOutput) ToEncryptionPropertiesDescriptionResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionResponsePtrOutput {
	return o
}

func (o EncryptionPropertiesDescriptionResponsePtrOutput) Elem() EncryptionPropertiesDescriptionResponseOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescriptionResponse) EncryptionPropertiesDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertiesDescriptionResponse
		return ret
	}).(EncryptionPropertiesDescriptionResponseOutput)
}

func (o EncryptionPropertiesDescriptionResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesDescriptionResponsePtrOutput) KeyVaultProperties() KeyVaultKeyPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescriptionResponse) []KeyVaultKeyPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultKeyPropertiesResponseArrayOutput)
}

type EnrichmentProperties struct {
	EndpointNames []string `pulumi:"endpointNames"`
	Key           string   `pulumi:"key"`
	Value         string   `pulumi:"value"`
}





type EnrichmentPropertiesInput interface {
	pulumi.Input

	ToEnrichmentPropertiesOutput() EnrichmentPropertiesOutput
	ToEnrichmentPropertiesOutputWithContext(context.Context) EnrichmentPropertiesOutput
}

type EnrichmentPropertiesArgs struct {
	EndpointNames pulumi.StringArrayInput `pulumi:"endpointNames"`
	Key           pulumi.StringInput      `pulumi:"key"`
	Value         pulumi.StringInput      `pulumi:"value"`
}

func (EnrichmentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnrichmentProperties)(nil)).Elem()
}

func (i EnrichmentPropertiesArgs) ToEnrichmentPropertiesOutput() EnrichmentPropertiesOutput {
	return i.ToEnrichmentPropertiesOutputWithContext(context.Background())
}

func (i EnrichmentPropertiesArgs) ToEnrichmentPropertiesOutputWithContext(ctx context.Context) EnrichmentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnrichmentPropertiesOutput)
}





type EnrichmentPropertiesArrayInput interface {
	pulumi.Input

	ToEnrichmentPropertiesArrayOutput() EnrichmentPropertiesArrayOutput
	ToEnrichmentPropertiesArrayOutputWithContext(context.Context) EnrichmentPropertiesArrayOutput
}

type EnrichmentPropertiesArray []EnrichmentPropertiesInput

func (EnrichmentPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnrichmentProperties)(nil)).Elem()
}

func (i EnrichmentPropertiesArray) ToEnrichmentPropertiesArrayOutput() EnrichmentPropertiesArrayOutput {
	return i.ToEnrichmentPropertiesArrayOutputWithContext(context.Background())
}

func (i EnrichmentPropertiesArray) ToEnrichmentPropertiesArrayOutputWithContext(ctx context.Context) EnrichmentPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnrichmentPropertiesArrayOutput)
}

type EnrichmentPropertiesOutput struct{ *pulumi.OutputState }

func (EnrichmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnrichmentProperties)(nil)).Elem()
}

func (o EnrichmentPropertiesOutput) ToEnrichmentPropertiesOutput() EnrichmentPropertiesOutput {
	return o
}

func (o EnrichmentPropertiesOutput) ToEnrichmentPropertiesOutputWithContext(ctx context.Context) EnrichmentPropertiesOutput {
	return o
}

func (o EnrichmentPropertiesOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EnrichmentProperties) []string { return v.EndpointNames }).(pulumi.StringArrayOutput)
}

func (o EnrichmentPropertiesOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v EnrichmentProperties) string { return v.Key }).(pulumi.StringOutput)
}

func (o EnrichmentPropertiesOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v EnrichmentProperties) string { return v.Value }).(pulumi.StringOutput)
}

type EnrichmentPropertiesArrayOutput struct{ *pulumi.OutputState }

func (EnrichmentPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnrichmentProperties)(nil)).Elem()
}

func (o EnrichmentPropertiesArrayOutput) ToEnrichmentPropertiesArrayOutput() EnrichmentPropertiesArrayOutput {
	return o
}

func (o EnrichmentPropertiesArrayOutput) ToEnrichmentPropertiesArrayOutputWithContext(ctx context.Context) EnrichmentPropertiesArrayOutput {
	return o
}

func (o EnrichmentPropertiesArrayOutput) Index(i pulumi.IntInput) EnrichmentPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnrichmentProperties {
		return vs[0].([]EnrichmentProperties)[vs[1].(int)]
	}).(EnrichmentPropertiesOutput)
}

type EnrichmentPropertiesResponse struct {
	EndpointNames []string `pulumi:"endpointNames"`
	Key           string   `pulumi:"key"`
	Value         string   `pulumi:"value"`
}

type EnrichmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EnrichmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnrichmentPropertiesResponse)(nil)).Elem()
}

func (o EnrichmentPropertiesResponseOutput) ToEnrichmentPropertiesResponseOutput() EnrichmentPropertiesResponseOutput {
	return o
}

func (o EnrichmentPropertiesResponseOutput) ToEnrichmentPropertiesResponseOutputWithContext(ctx context.Context) EnrichmentPropertiesResponseOutput {
	return o
}

func (o EnrichmentPropertiesResponseOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EnrichmentPropertiesResponse) []string { return v.EndpointNames }).(pulumi.StringArrayOutput)
}

func (o EnrichmentPropertiesResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v EnrichmentPropertiesResponse) string { return v.Key }).(pulumi.StringOutput)
}

func (o EnrichmentPropertiesResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v EnrichmentPropertiesResponse) string { return v.Value }).(pulumi.StringOutput)
}

type EnrichmentPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (EnrichmentPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnrichmentPropertiesResponse)(nil)).Elem()
}

func (o EnrichmentPropertiesResponseArrayOutput) ToEnrichmentPropertiesResponseArrayOutput() EnrichmentPropertiesResponseArrayOutput {
	return o
}

func (o EnrichmentPropertiesResponseArrayOutput) ToEnrichmentPropertiesResponseArrayOutputWithContext(ctx context.Context) EnrichmentPropertiesResponseArrayOutput {
	return o
}

func (o EnrichmentPropertiesResponseArrayOutput) Index(i pulumi.IntInput) EnrichmentPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnrichmentPropertiesResponse {
		return vs[0].([]EnrichmentPropertiesResponse)[vs[1].(int)]
	}).(EnrichmentPropertiesResponseOutput)
}

type EventHubConsumerGroupName struct {
	Name *string `pulumi:"name"`
}





type EventHubConsumerGroupNameInput interface {
	pulumi.Input

	ToEventHubConsumerGroupNameOutput() EventHubConsumerGroupNameOutput
	ToEventHubConsumerGroupNameOutputWithContext(context.Context) EventHubConsumerGroupNameOutput
}

type EventHubConsumerGroupNameArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (EventHubConsumerGroupNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubConsumerGroupName)(nil)).Elem()
}

func (i EventHubConsumerGroupNameArgs) ToEventHubConsumerGroupNameOutput() EventHubConsumerGroupNameOutput {
	return i.ToEventHubConsumerGroupNameOutputWithContext(context.Background())
}

func (i EventHubConsumerGroupNameArgs) ToEventHubConsumerGroupNameOutputWithContext(ctx context.Context) EventHubConsumerGroupNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubConsumerGroupNameOutput)
}

func (i EventHubConsumerGroupNameArgs) ToEventHubConsumerGroupNamePtrOutput() EventHubConsumerGroupNamePtrOutput {
	return i.ToEventHubConsumerGroupNamePtrOutputWithContext(context.Background())
}

func (i EventHubConsumerGroupNameArgs) ToEventHubConsumerGroupNamePtrOutputWithContext(ctx context.Context) EventHubConsumerGroupNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubConsumerGroupNameOutput).ToEventHubConsumerGroupNamePtrOutputWithContext(ctx)
}









type EventHubConsumerGroupNamePtrInput interface {
	pulumi.Input

	ToEventHubConsumerGroupNamePtrOutput() EventHubConsumerGroupNamePtrOutput
	ToEventHubConsumerGroupNamePtrOutputWithContext(context.Context) EventHubConsumerGroupNamePtrOutput
}

type eventHubConsumerGroupNamePtrType EventHubConsumerGroupNameArgs

func EventHubConsumerGroupNamePtr(v *EventHubConsumerGroupNameArgs) EventHubConsumerGroupNamePtrInput {
	return (*eventHubConsumerGroupNamePtrType)(v)
}

func (*eventHubConsumerGroupNamePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubConsumerGroupName)(nil)).Elem()
}

func (i *eventHubConsumerGroupNamePtrType) ToEventHubConsumerGroupNamePtrOutput() EventHubConsumerGroupNamePtrOutput {
	return i.ToEventHubConsumerGroupNamePtrOutputWithContext(context.Background())
}

func (i *eventHubConsumerGroupNamePtrType) ToEventHubConsumerGroupNamePtrOutputWithContext(ctx context.Context) EventHubConsumerGroupNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubConsumerGroupNamePtrOutput)
}

type EventHubConsumerGroupNameOutput struct{ *pulumi.OutputState }

func (EventHubConsumerGroupNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubConsumerGroupName)(nil)).Elem()
}

func (o EventHubConsumerGroupNameOutput) ToEventHubConsumerGroupNameOutput() EventHubConsumerGroupNameOutput {
	return o
}

func (o EventHubConsumerGroupNameOutput) ToEventHubConsumerGroupNameOutputWithContext(ctx context.Context) EventHubConsumerGroupNameOutput {
	return o
}

func (o EventHubConsumerGroupNameOutput) ToEventHubConsumerGroupNamePtrOutput() EventHubConsumerGroupNamePtrOutput {
	return o.ToEventHubConsumerGroupNamePtrOutputWithContext(context.Background())
}

func (o EventHubConsumerGroupNameOutput) ToEventHubConsumerGroupNamePtrOutputWithContext(ctx context.Context) EventHubConsumerGroupNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventHubConsumerGroupName) *EventHubConsumerGroupName {
		return &v
	}).(EventHubConsumerGroupNamePtrOutput)
}

func (o EventHubConsumerGroupNameOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubConsumerGroupName) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type EventHubConsumerGroupNamePtrOutput struct{ *pulumi.OutputState }

func (EventHubConsumerGroupNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubConsumerGroupName)(nil)).Elem()
}

func (o EventHubConsumerGroupNamePtrOutput) ToEventHubConsumerGroupNamePtrOutput() EventHubConsumerGroupNamePtrOutput {
	return o
}

func (o EventHubConsumerGroupNamePtrOutput) ToEventHubConsumerGroupNamePtrOutputWithContext(ctx context.Context) EventHubConsumerGroupNamePtrOutput {
	return o
}

func (o EventHubConsumerGroupNamePtrOutput) Elem() EventHubConsumerGroupNameOutput {
	return o.ApplyT(func(v *EventHubConsumerGroupName) EventHubConsumerGroupName {
		if v != nil {
			return *v
		}
		var ret EventHubConsumerGroupName
		return ret
	}).(EventHubConsumerGroupNameOutput)
}

func (o EventHubConsumerGroupNamePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubConsumerGroupName) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type EventHubProperties struct {
	PartitionCount      *int     `pulumi:"partitionCount"`
	RetentionTimeInDays *float64 `pulumi:"retentionTimeInDays"`
}





type EventHubPropertiesInput interface {
	pulumi.Input

	ToEventHubPropertiesOutput() EventHubPropertiesOutput
	ToEventHubPropertiesOutputWithContext(context.Context) EventHubPropertiesOutput
}

type EventHubPropertiesArgs struct {
	PartitionCount      pulumi.IntPtrInput     `pulumi:"partitionCount"`
	RetentionTimeInDays pulumi.Float64PtrInput `pulumi:"retentionTimeInDays"`
}

func (EventHubPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubProperties)(nil)).Elem()
}

func (i EventHubPropertiesArgs) ToEventHubPropertiesOutput() EventHubPropertiesOutput {
	return i.ToEventHubPropertiesOutputWithContext(context.Background())
}

func (i EventHubPropertiesArgs) ToEventHubPropertiesOutputWithContext(ctx context.Context) EventHubPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubPropertiesOutput)
}





type EventHubPropertiesMapInput interface {
	pulumi.Input

	ToEventHubPropertiesMapOutput() EventHubPropertiesMapOutput
	ToEventHubPropertiesMapOutputWithContext(context.Context) EventHubPropertiesMapOutput
}

type EventHubPropertiesMap map[string]EventHubPropertiesInput

func (EventHubPropertiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventHubProperties)(nil)).Elem()
}

func (i EventHubPropertiesMap) ToEventHubPropertiesMapOutput() EventHubPropertiesMapOutput {
	return i.ToEventHubPropertiesMapOutputWithContext(context.Background())
}

func (i EventHubPropertiesMap) ToEventHubPropertiesMapOutputWithContext(ctx context.Context) EventHubPropertiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubPropertiesMapOutput)
}

type EventHubPropertiesOutput struct{ *pulumi.OutputState }

func (EventHubPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubProperties)(nil)).Elem()
}

func (o EventHubPropertiesOutput) ToEventHubPropertiesOutput() EventHubPropertiesOutput {
	return o
}

func (o EventHubPropertiesOutput) ToEventHubPropertiesOutputWithContext(ctx context.Context) EventHubPropertiesOutput {
	return o
}

func (o EventHubPropertiesOutput) PartitionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EventHubProperties) *int { return v.PartitionCount }).(pulumi.IntPtrOutput)
}

func (o EventHubPropertiesOutput) RetentionTimeInDays() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EventHubProperties) *float64 { return v.RetentionTimeInDays }).(pulumi.Float64PtrOutput)
}

type EventHubPropertiesMapOutput struct{ *pulumi.OutputState }

func (EventHubPropertiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventHubProperties)(nil)).Elem()
}

func (o EventHubPropertiesMapOutput) ToEventHubPropertiesMapOutput() EventHubPropertiesMapOutput {
	return o
}

func (o EventHubPropertiesMapOutput) ToEventHubPropertiesMapOutputWithContext(ctx context.Context) EventHubPropertiesMapOutput {
	return o
}

func (o EventHubPropertiesMapOutput) MapIndex(k pulumi.StringInput) EventHubPropertiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventHubProperties {
		return vs[0].(map[string]EventHubProperties)[vs[1].(string)]
	}).(EventHubPropertiesOutput)
}

type EventHubPropertiesResponse struct {
	Endpoint            string   `pulumi:"endpoint"`
	PartitionCount      *int     `pulumi:"partitionCount"`
	PartitionIds        []string `pulumi:"partitionIds"`
	Path                string   `pulumi:"path"`
	RetentionTimeInDays *float64 `pulumi:"retentionTimeInDays"`
}

type EventHubPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EventHubPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubPropertiesResponse)(nil)).Elem()
}

func (o EventHubPropertiesResponseOutput) ToEventHubPropertiesResponseOutput() EventHubPropertiesResponseOutput {
	return o
}

func (o EventHubPropertiesResponseOutput) ToEventHubPropertiesResponseOutputWithContext(ctx context.Context) EventHubPropertiesResponseOutput {
	return o
}

func (o EventHubPropertiesResponseOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubPropertiesResponse) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o EventHubPropertiesResponseOutput) PartitionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EventHubPropertiesResponse) *int { return v.PartitionCount }).(pulumi.IntPtrOutput)
}

func (o EventHubPropertiesResponseOutput) PartitionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventHubPropertiesResponse) []string { return v.PartitionIds }).(pulumi.StringArrayOutput)
}

func (o EventHubPropertiesResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubPropertiesResponse) string { return v.Path }).(pulumi.StringOutput)
}

func (o EventHubPropertiesResponseOutput) RetentionTimeInDays() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EventHubPropertiesResponse) *float64 { return v.RetentionTimeInDays }).(pulumi.Float64PtrOutput)
}

type EventHubPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (EventHubPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventHubPropertiesResponse)(nil)).Elem()
}

func (o EventHubPropertiesResponseMapOutput) ToEventHubPropertiesResponseMapOutput() EventHubPropertiesResponseMapOutput {
	return o
}

func (o EventHubPropertiesResponseMapOutput) ToEventHubPropertiesResponseMapOutputWithContext(ctx context.Context) EventHubPropertiesResponseMapOutput {
	return o
}

func (o EventHubPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) EventHubPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventHubPropertiesResponse {
		return vs[0].(map[string]EventHubPropertiesResponse)[vs[1].(string)]
	}).(EventHubPropertiesResponseOutput)
}

type FallbackRouteProperties struct {
	Condition     *string  `pulumi:"condition"`
	EndpointNames []string `pulumi:"endpointNames"`
	IsEnabled     bool     `pulumi:"isEnabled"`
	Name          *string  `pulumi:"name"`
	Source        string   `pulumi:"source"`
}





type FallbackRoutePropertiesInput interface {
	pulumi.Input

	ToFallbackRoutePropertiesOutput() FallbackRoutePropertiesOutput
	ToFallbackRoutePropertiesOutputWithContext(context.Context) FallbackRoutePropertiesOutput
}

type FallbackRoutePropertiesArgs struct {
	Condition     pulumi.StringPtrInput   `pulumi:"condition"`
	EndpointNames pulumi.StringArrayInput `pulumi:"endpointNames"`
	IsEnabled     pulumi.BoolInput        `pulumi:"isEnabled"`
	Name          pulumi.StringPtrInput   `pulumi:"name"`
	Source        pulumi.StringInput      `pulumi:"source"`
}

func (FallbackRoutePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FallbackRouteProperties)(nil)).Elem()
}

func (i FallbackRoutePropertiesArgs) ToFallbackRoutePropertiesOutput() FallbackRoutePropertiesOutput {
	return i.ToFallbackRoutePropertiesOutputWithContext(context.Background())
}

func (i FallbackRoutePropertiesArgs) ToFallbackRoutePropertiesOutputWithContext(ctx context.Context) FallbackRoutePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRoutePropertiesOutput)
}

func (i FallbackRoutePropertiesArgs) ToFallbackRoutePropertiesPtrOutput() FallbackRoutePropertiesPtrOutput {
	return i.ToFallbackRoutePropertiesPtrOutputWithContext(context.Background())
}

func (i FallbackRoutePropertiesArgs) ToFallbackRoutePropertiesPtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRoutePropertiesOutput).ToFallbackRoutePropertiesPtrOutputWithContext(ctx)
}









type FallbackRoutePropertiesPtrInput interface {
	pulumi.Input

	ToFallbackRoutePropertiesPtrOutput() FallbackRoutePropertiesPtrOutput
	ToFallbackRoutePropertiesPtrOutputWithContext(context.Context) FallbackRoutePropertiesPtrOutput
}

type fallbackRoutePropertiesPtrType FallbackRoutePropertiesArgs

func FallbackRoutePropertiesPtr(v *FallbackRoutePropertiesArgs) FallbackRoutePropertiesPtrInput {
	return (*fallbackRoutePropertiesPtrType)(v)
}

func (*fallbackRoutePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FallbackRouteProperties)(nil)).Elem()
}

func (i *fallbackRoutePropertiesPtrType) ToFallbackRoutePropertiesPtrOutput() FallbackRoutePropertiesPtrOutput {
	return i.ToFallbackRoutePropertiesPtrOutputWithContext(context.Background())
}

func (i *fallbackRoutePropertiesPtrType) ToFallbackRoutePropertiesPtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRoutePropertiesPtrOutput)
}

type FallbackRoutePropertiesOutput struct{ *pulumi.OutputState }

func (FallbackRoutePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FallbackRouteProperties)(nil)).Elem()
}

func (o FallbackRoutePropertiesOutput) ToFallbackRoutePropertiesOutput() FallbackRoutePropertiesOutput {
	return o
}

func (o FallbackRoutePropertiesOutput) ToFallbackRoutePropertiesOutputWithContext(ctx context.Context) FallbackRoutePropertiesOutput {
	return o
}

func (o FallbackRoutePropertiesOutput) ToFallbackRoutePropertiesPtrOutput() FallbackRoutePropertiesPtrOutput {
	return o.ToFallbackRoutePropertiesPtrOutputWithContext(context.Background())
}

func (o FallbackRoutePropertiesOutput) ToFallbackRoutePropertiesPtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FallbackRouteProperties) *FallbackRouteProperties {
		return &v
	}).(FallbackRoutePropertiesPtrOutput)
}

func (o FallbackRoutePropertiesOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FallbackRouteProperties) *string { return v.Condition }).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FallbackRouteProperties) []string { return v.EndpointNames }).(pulumi.StringArrayOutput)
}

func (o FallbackRoutePropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v FallbackRouteProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o FallbackRoutePropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FallbackRouteProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v FallbackRouteProperties) string { return v.Source }).(pulumi.StringOutput)
}

type FallbackRoutePropertiesPtrOutput struct{ *pulumi.OutputState }

func (FallbackRoutePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FallbackRouteProperties)(nil)).Elem()
}

func (o FallbackRoutePropertiesPtrOutput) ToFallbackRoutePropertiesPtrOutput() FallbackRoutePropertiesPtrOutput {
	return o
}

func (o FallbackRoutePropertiesPtrOutput) ToFallbackRoutePropertiesPtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesPtrOutput {
	return o
}

func (o FallbackRoutePropertiesPtrOutput) Elem() FallbackRoutePropertiesOutput {
	return o.ApplyT(func(v *FallbackRouteProperties) FallbackRouteProperties {
		if v != nil {
			return *v
		}
		var ret FallbackRouteProperties
		return ret
	}).(FallbackRoutePropertiesOutput)
}

func (o FallbackRoutePropertiesPtrOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FallbackRouteProperties) *string {
		if v == nil {
			return nil
		}
		return v.Condition
	}).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesPtrOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FallbackRouteProperties) []string {
		if v == nil {
			return nil
		}
		return v.EndpointNames
	}).(pulumi.StringArrayOutput)
}

func (o FallbackRoutePropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FallbackRouteProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o FallbackRoutePropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FallbackRouteProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesPtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FallbackRouteProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Source
	}).(pulumi.StringPtrOutput)
}

type FallbackRoutePropertiesResponse struct {
	Condition     *string  `pulumi:"condition"`
	EndpointNames []string `pulumi:"endpointNames"`
	IsEnabled     bool     `pulumi:"isEnabled"`
	Name          *string  `pulumi:"name"`
	Source        string   `pulumi:"source"`
}

type FallbackRoutePropertiesResponseOutput struct{ *pulumi.OutputState }

func (FallbackRoutePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FallbackRoutePropertiesResponse)(nil)).Elem()
}

func (o FallbackRoutePropertiesResponseOutput) ToFallbackRoutePropertiesResponseOutput() FallbackRoutePropertiesResponseOutput {
	return o
}

func (o FallbackRoutePropertiesResponseOutput) ToFallbackRoutePropertiesResponseOutputWithContext(ctx context.Context) FallbackRoutePropertiesResponseOutput {
	return o
}

func (o FallbackRoutePropertiesResponseOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FallbackRoutePropertiesResponse) *string { return v.Condition }).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesResponseOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FallbackRoutePropertiesResponse) []string { return v.EndpointNames }).(pulumi.StringArrayOutput)
}

func (o FallbackRoutePropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v FallbackRoutePropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o FallbackRoutePropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FallbackRoutePropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesResponseOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v FallbackRoutePropertiesResponse) string { return v.Source }).(pulumi.StringOutput)
}

type FallbackRoutePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (FallbackRoutePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FallbackRoutePropertiesResponse)(nil)).Elem()
}

func (o FallbackRoutePropertiesResponsePtrOutput) ToFallbackRoutePropertiesResponsePtrOutput() FallbackRoutePropertiesResponsePtrOutput {
	return o
}

func (o FallbackRoutePropertiesResponsePtrOutput) ToFallbackRoutePropertiesResponsePtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesResponsePtrOutput {
	return o
}

func (o FallbackRoutePropertiesResponsePtrOutput) Elem() FallbackRoutePropertiesResponseOutput {
	return o.ApplyT(func(v *FallbackRoutePropertiesResponse) FallbackRoutePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret FallbackRoutePropertiesResponse
		return ret
	}).(FallbackRoutePropertiesResponseOutput)
}

func (o FallbackRoutePropertiesResponsePtrOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FallbackRoutePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Condition
	}).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesResponsePtrOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FallbackRoutePropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.EndpointNames
	}).(pulumi.StringArrayOutput)
}

func (o FallbackRoutePropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FallbackRoutePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o FallbackRoutePropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FallbackRoutePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesResponsePtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FallbackRoutePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Source
	}).(pulumi.StringPtrOutput)
}

type FeedbackProperties struct {
	LockDurationAsIso8601 *string `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      *int    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          *string `pulumi:"ttlAsIso8601"`
}





type FeedbackPropertiesInput interface {
	pulumi.Input

	ToFeedbackPropertiesOutput() FeedbackPropertiesOutput
	ToFeedbackPropertiesOutputWithContext(context.Context) FeedbackPropertiesOutput
}

type FeedbackPropertiesArgs struct {
	LockDurationAsIso8601 pulumi.StringPtrInput `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      pulumi.IntPtrInput    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          pulumi.StringPtrInput `pulumi:"ttlAsIso8601"`
}

func (FeedbackPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FeedbackProperties)(nil)).Elem()
}

func (i FeedbackPropertiesArgs) ToFeedbackPropertiesOutput() FeedbackPropertiesOutput {
	return i.ToFeedbackPropertiesOutputWithContext(context.Background())
}

func (i FeedbackPropertiesArgs) ToFeedbackPropertiesOutputWithContext(ctx context.Context) FeedbackPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedbackPropertiesOutput)
}

func (i FeedbackPropertiesArgs) ToFeedbackPropertiesPtrOutput() FeedbackPropertiesPtrOutput {
	return i.ToFeedbackPropertiesPtrOutputWithContext(context.Background())
}

func (i FeedbackPropertiesArgs) ToFeedbackPropertiesPtrOutputWithContext(ctx context.Context) FeedbackPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedbackPropertiesOutput).ToFeedbackPropertiesPtrOutputWithContext(ctx)
}









type FeedbackPropertiesPtrInput interface {
	pulumi.Input

	ToFeedbackPropertiesPtrOutput() FeedbackPropertiesPtrOutput
	ToFeedbackPropertiesPtrOutputWithContext(context.Context) FeedbackPropertiesPtrOutput
}

type feedbackPropertiesPtrType FeedbackPropertiesArgs

func FeedbackPropertiesPtr(v *FeedbackPropertiesArgs) FeedbackPropertiesPtrInput {
	return (*feedbackPropertiesPtrType)(v)
}

func (*feedbackPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FeedbackProperties)(nil)).Elem()
}

func (i *feedbackPropertiesPtrType) ToFeedbackPropertiesPtrOutput() FeedbackPropertiesPtrOutput {
	return i.ToFeedbackPropertiesPtrOutputWithContext(context.Background())
}

func (i *feedbackPropertiesPtrType) ToFeedbackPropertiesPtrOutputWithContext(ctx context.Context) FeedbackPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedbackPropertiesPtrOutput)
}

type FeedbackPropertiesOutput struct{ *pulumi.OutputState }

func (FeedbackPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FeedbackProperties)(nil)).Elem()
}

func (o FeedbackPropertiesOutput) ToFeedbackPropertiesOutput() FeedbackPropertiesOutput {
	return o
}

func (o FeedbackPropertiesOutput) ToFeedbackPropertiesOutputWithContext(ctx context.Context) FeedbackPropertiesOutput {
	return o
}

func (o FeedbackPropertiesOutput) ToFeedbackPropertiesPtrOutput() FeedbackPropertiesPtrOutput {
	return o.ToFeedbackPropertiesPtrOutputWithContext(context.Background())
}

func (o FeedbackPropertiesOutput) ToFeedbackPropertiesPtrOutputWithContext(ctx context.Context) FeedbackPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FeedbackProperties) *FeedbackProperties {
		return &v
	}).(FeedbackPropertiesPtrOutput)
}

func (o FeedbackPropertiesOutput) LockDurationAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FeedbackProperties) *string { return v.LockDurationAsIso8601 }).(pulumi.StringPtrOutput)
}

func (o FeedbackPropertiesOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FeedbackProperties) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

func (o FeedbackPropertiesOutput) TtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FeedbackProperties) *string { return v.TtlAsIso8601 }).(pulumi.StringPtrOutput)
}

type FeedbackPropertiesPtrOutput struct{ *pulumi.OutputState }

func (FeedbackPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeedbackProperties)(nil)).Elem()
}

func (o FeedbackPropertiesPtrOutput) ToFeedbackPropertiesPtrOutput() FeedbackPropertiesPtrOutput {
	return o
}

func (o FeedbackPropertiesPtrOutput) ToFeedbackPropertiesPtrOutputWithContext(ctx context.Context) FeedbackPropertiesPtrOutput {
	return o
}

func (o FeedbackPropertiesPtrOutput) Elem() FeedbackPropertiesOutput {
	return o.ApplyT(func(v *FeedbackProperties) FeedbackProperties {
		if v != nil {
			return *v
		}
		var ret FeedbackProperties
		return ret
	}).(FeedbackPropertiesOutput)
}

func (o FeedbackPropertiesPtrOutput) LockDurationAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeedbackProperties) *string {
		if v == nil {
			return nil
		}
		return v.LockDurationAsIso8601
	}).(pulumi.StringPtrOutput)
}

func (o FeedbackPropertiesPtrOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FeedbackProperties) *int {
		if v == nil {
			return nil
		}
		return v.MaxDeliveryCount
	}).(pulumi.IntPtrOutput)
}

func (o FeedbackPropertiesPtrOutput) TtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeedbackProperties) *string {
		if v == nil {
			return nil
		}
		return v.TtlAsIso8601
	}).(pulumi.StringPtrOutput)
}

type FeedbackPropertiesResponse struct {
	LockDurationAsIso8601 *string `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      *int    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          *string `pulumi:"ttlAsIso8601"`
}

type FeedbackPropertiesResponseOutput struct{ *pulumi.OutputState }

func (FeedbackPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FeedbackPropertiesResponse)(nil)).Elem()
}

func (o FeedbackPropertiesResponseOutput) ToFeedbackPropertiesResponseOutput() FeedbackPropertiesResponseOutput {
	return o
}

func (o FeedbackPropertiesResponseOutput) ToFeedbackPropertiesResponseOutputWithContext(ctx context.Context) FeedbackPropertiesResponseOutput {
	return o
}

func (o FeedbackPropertiesResponseOutput) LockDurationAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FeedbackPropertiesResponse) *string { return v.LockDurationAsIso8601 }).(pulumi.StringPtrOutput)
}

func (o FeedbackPropertiesResponseOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FeedbackPropertiesResponse) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

func (o FeedbackPropertiesResponseOutput) TtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FeedbackPropertiesResponse) *string { return v.TtlAsIso8601 }).(pulumi.StringPtrOutput)
}

type FeedbackPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (FeedbackPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeedbackPropertiesResponse)(nil)).Elem()
}

func (o FeedbackPropertiesResponsePtrOutput) ToFeedbackPropertiesResponsePtrOutput() FeedbackPropertiesResponsePtrOutput {
	return o
}

func (o FeedbackPropertiesResponsePtrOutput) ToFeedbackPropertiesResponsePtrOutputWithContext(ctx context.Context) FeedbackPropertiesResponsePtrOutput {
	return o
}

func (o FeedbackPropertiesResponsePtrOutput) Elem() FeedbackPropertiesResponseOutput {
	return o.ApplyT(func(v *FeedbackPropertiesResponse) FeedbackPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret FeedbackPropertiesResponse
		return ret
	}).(FeedbackPropertiesResponseOutput)
}

func (o FeedbackPropertiesResponsePtrOutput) LockDurationAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeedbackPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LockDurationAsIso8601
	}).(pulumi.StringPtrOutput)
}

func (o FeedbackPropertiesResponsePtrOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FeedbackPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxDeliveryCount
	}).(pulumi.IntPtrOutput)
}

func (o FeedbackPropertiesResponsePtrOutput) TtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeedbackPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TtlAsIso8601
	}).(pulumi.StringPtrOutput)
}

type IotHubLocationDescriptionResponse struct {
	Location *string `pulumi:"location"`
	Role     *string `pulumi:"role"`
}

type IotHubLocationDescriptionResponseOutput struct{ *pulumi.OutputState }

func (IotHubLocationDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubLocationDescriptionResponse)(nil)).Elem()
}

func (o IotHubLocationDescriptionResponseOutput) ToIotHubLocationDescriptionResponseOutput() IotHubLocationDescriptionResponseOutput {
	return o
}

func (o IotHubLocationDescriptionResponseOutput) ToIotHubLocationDescriptionResponseOutputWithContext(ctx context.Context) IotHubLocationDescriptionResponseOutput {
	return o
}

func (o IotHubLocationDescriptionResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubLocationDescriptionResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o IotHubLocationDescriptionResponseOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubLocationDescriptionResponse) *string { return v.Role }).(pulumi.StringPtrOutput)
}

type IotHubLocationDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (IotHubLocationDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubLocationDescriptionResponse)(nil)).Elem()
}

func (o IotHubLocationDescriptionResponseArrayOutput) ToIotHubLocationDescriptionResponseArrayOutput() IotHubLocationDescriptionResponseArrayOutput {
	return o
}

func (o IotHubLocationDescriptionResponseArrayOutput) ToIotHubLocationDescriptionResponseArrayOutputWithContext(ctx context.Context) IotHubLocationDescriptionResponseArrayOutput {
	return o
}

func (o IotHubLocationDescriptionResponseArrayOutput) Index(i pulumi.IntInput) IotHubLocationDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IotHubLocationDescriptionResponse {
		return vs[0].([]IotHubLocationDescriptionResponse)[vs[1].(int)]
	}).(IotHubLocationDescriptionResponseOutput)
}

type IotHubProperties struct {
	AuthorizationPolicies         []SharedAccessSignatureAuthorizationRule `pulumi:"authorizationPolicies"`
	CloudToDevice                 *CloudToDeviceProperties                 `pulumi:"cloudToDevice"`
	Comments                      *string                                  `pulumi:"comments"`
	DeviceStreams                 *IotHubPropertiesDeviceStreams           `pulumi:"deviceStreams"`
	EnableFileUploadNotifications *bool                                    `pulumi:"enableFileUploadNotifications"`
	Encryption                    *EncryptionPropertiesDescription         `pulumi:"encryption"`
	EventHubEndpoints             map[string]EventHubProperties            `pulumi:"eventHubEndpoints"`
	Features                      *string                                  `pulumi:"features"`
	IpFilterRules                 []IpFilterRule                           `pulumi:"ipFilterRules"`
	MessagingEndpoints            map[string]MessagingEndpointProperties   `pulumi:"messagingEndpoints"`
	MinTlsVersion                 *string                                  `pulumi:"minTlsVersion"`
	NetworkRuleSets               *NetworkRuleSetProperties                `pulumi:"networkRuleSets"`
	PrivateEndpointConnections    []PrivateEndpointConnectionType          `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess           *string                                  `pulumi:"publicNetworkAccess"`
	Routing                       *RoutingProperties                       `pulumi:"routing"`
	StorageEndpoints              map[string]StorageEndpointProperties     `pulumi:"storageEndpoints"`
}


func (val *IotHubProperties) Defaults() *IotHubProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkRuleSets = tmp.NetworkRuleSets.Defaults()

	return &tmp
}





type IotHubPropertiesInput interface {
	pulumi.Input

	ToIotHubPropertiesOutput() IotHubPropertiesOutput
	ToIotHubPropertiesOutputWithContext(context.Context) IotHubPropertiesOutput
}

type IotHubPropertiesArgs struct {
	AuthorizationPolicies         SharedAccessSignatureAuthorizationRuleArrayInput `pulumi:"authorizationPolicies"`
	CloudToDevice                 CloudToDevicePropertiesPtrInput                  `pulumi:"cloudToDevice"`
	Comments                      pulumi.StringPtrInput                            `pulumi:"comments"`
	DeviceStreams                 IotHubPropertiesDeviceStreamsPtrInput            `pulumi:"deviceStreams"`
	EnableFileUploadNotifications pulumi.BoolPtrInput                              `pulumi:"enableFileUploadNotifications"`
	Encryption                    EncryptionPropertiesDescriptionPtrInput          `pulumi:"encryption"`
	EventHubEndpoints             EventHubPropertiesMapInput                       `pulumi:"eventHubEndpoints"`
	Features                      pulumi.StringPtrInput                            `pulumi:"features"`
	IpFilterRules                 IpFilterRuleArrayInput                           `pulumi:"ipFilterRules"`
	MessagingEndpoints            MessagingEndpointPropertiesMapInput              `pulumi:"messagingEndpoints"`
	MinTlsVersion                 pulumi.StringPtrInput                            `pulumi:"minTlsVersion"`
	NetworkRuleSets               NetworkRuleSetPropertiesPtrInput                 `pulumi:"networkRuleSets"`
	PrivateEndpointConnections    PrivateEndpointConnectionTypeArrayInput          `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess           pulumi.StringPtrInput                            `pulumi:"publicNetworkAccess"`
	Routing                       RoutingPropertiesPtrInput                        `pulumi:"routing"`
	StorageEndpoints              StorageEndpointPropertiesMapInput                `pulumi:"storageEndpoints"`
}

func (IotHubPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubProperties)(nil)).Elem()
}

func (i IotHubPropertiesArgs) ToIotHubPropertiesOutput() IotHubPropertiesOutput {
	return i.ToIotHubPropertiesOutputWithContext(context.Background())
}

func (i IotHubPropertiesArgs) ToIotHubPropertiesOutputWithContext(ctx context.Context) IotHubPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubPropertiesOutput)
}

func (i IotHubPropertiesArgs) ToIotHubPropertiesPtrOutput() IotHubPropertiesPtrOutput {
	return i.ToIotHubPropertiesPtrOutputWithContext(context.Background())
}

func (i IotHubPropertiesArgs) ToIotHubPropertiesPtrOutputWithContext(ctx context.Context) IotHubPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubPropertiesOutput).ToIotHubPropertiesPtrOutputWithContext(ctx)
}









type IotHubPropertiesPtrInput interface {
	pulumi.Input

	ToIotHubPropertiesPtrOutput() IotHubPropertiesPtrOutput
	ToIotHubPropertiesPtrOutputWithContext(context.Context) IotHubPropertiesPtrOutput
}

type iotHubPropertiesPtrType IotHubPropertiesArgs

func IotHubPropertiesPtr(v *IotHubPropertiesArgs) IotHubPropertiesPtrInput {
	return (*iotHubPropertiesPtrType)(v)
}

func (*iotHubPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubProperties)(nil)).Elem()
}

func (i *iotHubPropertiesPtrType) ToIotHubPropertiesPtrOutput() IotHubPropertiesPtrOutput {
	return i.ToIotHubPropertiesPtrOutputWithContext(context.Background())
}

func (i *iotHubPropertiesPtrType) ToIotHubPropertiesPtrOutputWithContext(ctx context.Context) IotHubPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubPropertiesPtrOutput)
}

type IotHubPropertiesOutput struct{ *pulumi.OutputState }

func (IotHubPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubProperties)(nil)).Elem()
}

func (o IotHubPropertiesOutput) ToIotHubPropertiesOutput() IotHubPropertiesOutput {
	return o
}

func (o IotHubPropertiesOutput) ToIotHubPropertiesOutputWithContext(ctx context.Context) IotHubPropertiesOutput {
	return o
}

func (o IotHubPropertiesOutput) ToIotHubPropertiesPtrOutput() IotHubPropertiesPtrOutput {
	return o.ToIotHubPropertiesPtrOutputWithContext(context.Background())
}

func (o IotHubPropertiesOutput) ToIotHubPropertiesPtrOutputWithContext(ctx context.Context) IotHubPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotHubProperties) *IotHubProperties {
		return &v
	}).(IotHubPropertiesPtrOutput)
}

func (o IotHubPropertiesOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleArrayOutput {
	return o.ApplyT(func(v IotHubProperties) []SharedAccessSignatureAuthorizationRule { return v.AuthorizationPolicies }).(SharedAccessSignatureAuthorizationRuleArrayOutput)
}

func (o IotHubPropertiesOutput) CloudToDevice() CloudToDevicePropertiesPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *CloudToDeviceProperties { return v.CloudToDevice }).(CloudToDevicePropertiesPtrOutput)
}

func (o IotHubPropertiesOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *string { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesOutput) DeviceStreams() IotHubPropertiesDeviceStreamsPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *IotHubPropertiesDeviceStreams { return v.DeviceStreams }).(IotHubPropertiesDeviceStreamsPtrOutput)
}

func (o IotHubPropertiesOutput) EnableFileUploadNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *bool { return v.EnableFileUploadNotifications }).(pulumi.BoolPtrOutput)
}

func (o IotHubPropertiesOutput) Encryption() EncryptionPropertiesDescriptionPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *EncryptionPropertiesDescription { return v.Encryption }).(EncryptionPropertiesDescriptionPtrOutput)
}

func (o IotHubPropertiesOutput) EventHubEndpoints() EventHubPropertiesMapOutput {
	return o.ApplyT(func(v IotHubProperties) map[string]EventHubProperties { return v.EventHubEndpoints }).(EventHubPropertiesMapOutput)
}

func (o IotHubPropertiesOutput) Features() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *string { return v.Features }).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesOutput) IpFilterRules() IpFilterRuleArrayOutput {
	return o.ApplyT(func(v IotHubProperties) []IpFilterRule { return v.IpFilterRules }).(IpFilterRuleArrayOutput)
}

func (o IotHubPropertiesOutput) MessagingEndpoints() MessagingEndpointPropertiesMapOutput {
	return o.ApplyT(func(v IotHubProperties) map[string]MessagingEndpointProperties { return v.MessagingEndpoints }).(MessagingEndpointPropertiesMapOutput)
}

func (o IotHubPropertiesOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *string { return v.MinTlsVersion }).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesOutput) NetworkRuleSets() NetworkRuleSetPropertiesPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *NetworkRuleSetProperties { return v.NetworkRuleSets }).(NetworkRuleSetPropertiesPtrOutput)
}

func (o IotHubPropertiesOutput) PrivateEndpointConnections() PrivateEndpointConnectionTypeArrayOutput {
	return o.ApplyT(func(v IotHubProperties) []PrivateEndpointConnectionType { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionTypeArrayOutput)
}

func (o IotHubPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesOutput) Routing() RoutingPropertiesPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *RoutingProperties { return v.Routing }).(RoutingPropertiesPtrOutput)
}

func (o IotHubPropertiesOutput) StorageEndpoints() StorageEndpointPropertiesMapOutput {
	return o.ApplyT(func(v IotHubProperties) map[string]StorageEndpointProperties { return v.StorageEndpoints }).(StorageEndpointPropertiesMapOutput)
}

type IotHubPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IotHubPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubProperties)(nil)).Elem()
}

func (o IotHubPropertiesPtrOutput) ToIotHubPropertiesPtrOutput() IotHubPropertiesPtrOutput {
	return o
}

func (o IotHubPropertiesPtrOutput) ToIotHubPropertiesPtrOutputWithContext(ctx context.Context) IotHubPropertiesPtrOutput {
	return o
}

func (o IotHubPropertiesPtrOutput) Elem() IotHubPropertiesOutput {
	return o.ApplyT(func(v *IotHubProperties) IotHubProperties {
		if v != nil {
			return *v
		}
		var ret IotHubProperties
		return ret
	}).(IotHubPropertiesOutput)
}

func (o IotHubPropertiesPtrOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleArrayOutput {
	return o.ApplyT(func(v *IotHubProperties) []SharedAccessSignatureAuthorizationRule {
		if v == nil {
			return nil
		}
		return v.AuthorizationPolicies
	}).(SharedAccessSignatureAuthorizationRuleArrayOutput)
}

func (o IotHubPropertiesPtrOutput) CloudToDevice() CloudToDevicePropertiesPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *CloudToDeviceProperties {
		if v == nil {
			return nil
		}
		return v.CloudToDevice
	}).(CloudToDevicePropertiesPtrOutput)
}

func (o IotHubPropertiesPtrOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *string {
		if v == nil {
			return nil
		}
		return v.Comments
	}).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesPtrOutput) DeviceStreams() IotHubPropertiesDeviceStreamsPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *IotHubPropertiesDeviceStreams {
		if v == nil {
			return nil
		}
		return v.DeviceStreams
	}).(IotHubPropertiesDeviceStreamsPtrOutput)
}

func (o IotHubPropertiesPtrOutput) EnableFileUploadNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableFileUploadNotifications
	}).(pulumi.BoolPtrOutput)
}

func (o IotHubPropertiesPtrOutput) Encryption() EncryptionPropertiesDescriptionPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *EncryptionPropertiesDescription {
		if v == nil {
			return nil
		}
		return v.Encryption
	}).(EncryptionPropertiesDescriptionPtrOutput)
}

func (o IotHubPropertiesPtrOutput) EventHubEndpoints() EventHubPropertiesMapOutput {
	return o.ApplyT(func(v *IotHubProperties) map[string]EventHubProperties {
		if v == nil {
			return nil
		}
		return v.EventHubEndpoints
	}).(EventHubPropertiesMapOutput)
}

func (o IotHubPropertiesPtrOutput) Features() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *string {
		if v == nil {
			return nil
		}
		return v.Features
	}).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesPtrOutput) IpFilterRules() IpFilterRuleArrayOutput {
	return o.ApplyT(func(v *IotHubProperties) []IpFilterRule {
		if v == nil {
			return nil
		}
		return v.IpFilterRules
	}).(IpFilterRuleArrayOutput)
}

func (o IotHubPropertiesPtrOutput) MessagingEndpoints() MessagingEndpointPropertiesMapOutput {
	return o.ApplyT(func(v *IotHubProperties) map[string]MessagingEndpointProperties {
		if v == nil {
			return nil
		}
		return v.MessagingEndpoints
	}).(MessagingEndpointPropertiesMapOutput)
}

func (o IotHubPropertiesPtrOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *string {
		if v == nil {
			return nil
		}
		return v.MinTlsVersion
	}).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesPtrOutput) NetworkRuleSets() NetworkRuleSetPropertiesPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *NetworkRuleSetProperties {
		if v == nil {
			return nil
		}
		return v.NetworkRuleSets
	}).(NetworkRuleSetPropertiesPtrOutput)
}

func (o IotHubPropertiesPtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionTypeArrayOutput {
	return o.ApplyT(func(v *IotHubProperties) []PrivateEndpointConnectionType {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionTypeArrayOutput)
}

func (o IotHubPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesPtrOutput) Routing() RoutingPropertiesPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *RoutingProperties {
		if v == nil {
			return nil
		}
		return v.Routing
	}).(RoutingPropertiesPtrOutput)
}

func (o IotHubPropertiesPtrOutput) StorageEndpoints() StorageEndpointPropertiesMapOutput {
	return o.ApplyT(func(v *IotHubProperties) map[string]StorageEndpointProperties {
		if v == nil {
			return nil
		}
		return v.StorageEndpoints
	}).(StorageEndpointPropertiesMapOutput)
}

type IotHubPropertiesDeviceStreams struct {
	StreamingEndpoints []string `pulumi:"streamingEndpoints"`
}





type IotHubPropertiesDeviceStreamsInput interface {
	pulumi.Input

	ToIotHubPropertiesDeviceStreamsOutput() IotHubPropertiesDeviceStreamsOutput
	ToIotHubPropertiesDeviceStreamsOutputWithContext(context.Context) IotHubPropertiesDeviceStreamsOutput
}

type IotHubPropertiesDeviceStreamsArgs struct {
	StreamingEndpoints pulumi.StringArrayInput `pulumi:"streamingEndpoints"`
}

func (IotHubPropertiesDeviceStreamsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubPropertiesDeviceStreams)(nil)).Elem()
}

func (i IotHubPropertiesDeviceStreamsArgs) ToIotHubPropertiesDeviceStreamsOutput() IotHubPropertiesDeviceStreamsOutput {
	return i.ToIotHubPropertiesDeviceStreamsOutputWithContext(context.Background())
}

func (i IotHubPropertiesDeviceStreamsArgs) ToIotHubPropertiesDeviceStreamsOutputWithContext(ctx context.Context) IotHubPropertiesDeviceStreamsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubPropertiesDeviceStreamsOutput)
}

func (i IotHubPropertiesDeviceStreamsArgs) ToIotHubPropertiesDeviceStreamsPtrOutput() IotHubPropertiesDeviceStreamsPtrOutput {
	return i.ToIotHubPropertiesDeviceStreamsPtrOutputWithContext(context.Background())
}

func (i IotHubPropertiesDeviceStreamsArgs) ToIotHubPropertiesDeviceStreamsPtrOutputWithContext(ctx context.Context) IotHubPropertiesDeviceStreamsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubPropertiesDeviceStreamsOutput).ToIotHubPropertiesDeviceStreamsPtrOutputWithContext(ctx)
}









type IotHubPropertiesDeviceStreamsPtrInput interface {
	pulumi.Input

	ToIotHubPropertiesDeviceStreamsPtrOutput() IotHubPropertiesDeviceStreamsPtrOutput
	ToIotHubPropertiesDeviceStreamsPtrOutputWithContext(context.Context) IotHubPropertiesDeviceStreamsPtrOutput
}

type iotHubPropertiesDeviceStreamsPtrType IotHubPropertiesDeviceStreamsArgs

func IotHubPropertiesDeviceStreamsPtr(v *IotHubPropertiesDeviceStreamsArgs) IotHubPropertiesDeviceStreamsPtrInput {
	return (*iotHubPropertiesDeviceStreamsPtrType)(v)
}

func (*iotHubPropertiesDeviceStreamsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubPropertiesDeviceStreams)(nil)).Elem()
}

func (i *iotHubPropertiesDeviceStreamsPtrType) ToIotHubPropertiesDeviceStreamsPtrOutput() IotHubPropertiesDeviceStreamsPtrOutput {
	return i.ToIotHubPropertiesDeviceStreamsPtrOutputWithContext(context.Background())
}

func (i *iotHubPropertiesDeviceStreamsPtrType) ToIotHubPropertiesDeviceStreamsPtrOutputWithContext(ctx context.Context) IotHubPropertiesDeviceStreamsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubPropertiesDeviceStreamsPtrOutput)
}

type IotHubPropertiesDeviceStreamsOutput struct{ *pulumi.OutputState }

func (IotHubPropertiesDeviceStreamsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubPropertiesDeviceStreams)(nil)).Elem()
}

func (o IotHubPropertiesDeviceStreamsOutput) ToIotHubPropertiesDeviceStreamsOutput() IotHubPropertiesDeviceStreamsOutput {
	return o
}

func (o IotHubPropertiesDeviceStreamsOutput) ToIotHubPropertiesDeviceStreamsOutputWithContext(ctx context.Context) IotHubPropertiesDeviceStreamsOutput {
	return o
}

func (o IotHubPropertiesDeviceStreamsOutput) ToIotHubPropertiesDeviceStreamsPtrOutput() IotHubPropertiesDeviceStreamsPtrOutput {
	return o.ToIotHubPropertiesDeviceStreamsPtrOutputWithContext(context.Background())
}

func (o IotHubPropertiesDeviceStreamsOutput) ToIotHubPropertiesDeviceStreamsPtrOutputWithContext(ctx context.Context) IotHubPropertiesDeviceStreamsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotHubPropertiesDeviceStreams) *IotHubPropertiesDeviceStreams {
		return &v
	}).(IotHubPropertiesDeviceStreamsPtrOutput)
}

func (o IotHubPropertiesDeviceStreamsOutput) StreamingEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IotHubPropertiesDeviceStreams) []string { return v.StreamingEndpoints }).(pulumi.StringArrayOutput)
}

type IotHubPropertiesDeviceStreamsPtrOutput struct{ *pulumi.OutputState }

func (IotHubPropertiesDeviceStreamsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubPropertiesDeviceStreams)(nil)).Elem()
}

func (o IotHubPropertiesDeviceStreamsPtrOutput) ToIotHubPropertiesDeviceStreamsPtrOutput() IotHubPropertiesDeviceStreamsPtrOutput {
	return o
}

func (o IotHubPropertiesDeviceStreamsPtrOutput) ToIotHubPropertiesDeviceStreamsPtrOutputWithContext(ctx context.Context) IotHubPropertiesDeviceStreamsPtrOutput {
	return o
}

func (o IotHubPropertiesDeviceStreamsPtrOutput) Elem() IotHubPropertiesDeviceStreamsOutput {
	return o.ApplyT(func(v *IotHubPropertiesDeviceStreams) IotHubPropertiesDeviceStreams {
		if v != nil {
			return *v
		}
		var ret IotHubPropertiesDeviceStreams
		return ret
	}).(IotHubPropertiesDeviceStreamsOutput)
}

func (o IotHubPropertiesDeviceStreamsPtrOutput) StreamingEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IotHubPropertiesDeviceStreams) []string {
		if v == nil {
			return nil
		}
		return v.StreamingEndpoints
	}).(pulumi.StringArrayOutput)
}

type IotHubPropertiesResponse struct {
	AuthorizationPolicies         []SharedAccessSignatureAuthorizationRuleResponse `pulumi:"authorizationPolicies"`
	CloudToDevice                 *CloudToDevicePropertiesResponse                 `pulumi:"cloudToDevice"`
	Comments                      *string                                          `pulumi:"comments"`
	DeviceStreams                 *IotHubPropertiesResponseDeviceStreams           `pulumi:"deviceStreams"`
	EnableFileUploadNotifications *bool                                            `pulumi:"enableFileUploadNotifications"`
	Encryption                    *EncryptionPropertiesDescriptionResponse         `pulumi:"encryption"`
	EventHubEndpoints             map[string]EventHubPropertiesResponse            `pulumi:"eventHubEndpoints"`
	Features                      *string                                          `pulumi:"features"`
	HostName                      string                                           `pulumi:"hostName"`
	IpFilterRules                 []IpFilterRuleResponse                           `pulumi:"ipFilterRules"`
	Locations                     []IotHubLocationDescriptionResponse              `pulumi:"locations"`
	MessagingEndpoints            map[string]MessagingEndpointPropertiesResponse   `pulumi:"messagingEndpoints"`
	MinTlsVersion                 *string                                          `pulumi:"minTlsVersion"`
	NetworkRuleSets               *NetworkRuleSetPropertiesResponse                `pulumi:"networkRuleSets"`
	PrivateEndpointConnections    []PrivateEndpointConnectionResponse              `pulumi:"privateEndpointConnections"`
	ProvisioningState             string                                           `pulumi:"provisioningState"`
	PublicNetworkAccess           *string                                          `pulumi:"publicNetworkAccess"`
	Routing                       *RoutingPropertiesResponse                       `pulumi:"routing"`
	State                         string                                           `pulumi:"state"`
	StorageEndpoints              map[string]StorageEndpointPropertiesResponse     `pulumi:"storageEndpoints"`
}


func (val *IotHubPropertiesResponse) Defaults() *IotHubPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkRuleSets = tmp.NetworkRuleSets.Defaults()

	return &tmp
}

type IotHubPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IotHubPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubPropertiesResponse)(nil)).Elem()
}

func (o IotHubPropertiesResponseOutput) ToIotHubPropertiesResponseOutput() IotHubPropertiesResponseOutput {
	return o
}

func (o IotHubPropertiesResponseOutput) ToIotHubPropertiesResponseOutputWithContext(ctx context.Context) IotHubPropertiesResponseOutput {
	return o
}

func (o IotHubPropertiesResponseOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleResponseArrayOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) []SharedAccessSignatureAuthorizationRuleResponse {
		return v.AuthorizationPolicies
	}).(SharedAccessSignatureAuthorizationRuleResponseArrayOutput)
}

func (o IotHubPropertiesResponseOutput) CloudToDevice() CloudToDevicePropertiesResponsePtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *CloudToDevicePropertiesResponse { return v.CloudToDevice }).(CloudToDevicePropertiesResponsePtrOutput)
}

func (o IotHubPropertiesResponseOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *string { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesResponseOutput) DeviceStreams() IotHubPropertiesResponseDeviceStreamsPtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *IotHubPropertiesResponseDeviceStreams { return v.DeviceStreams }).(IotHubPropertiesResponseDeviceStreamsPtrOutput)
}

func (o IotHubPropertiesResponseOutput) EnableFileUploadNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *bool { return v.EnableFileUploadNotifications }).(pulumi.BoolPtrOutput)
}

func (o IotHubPropertiesResponseOutput) Encryption() EncryptionPropertiesDescriptionResponsePtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *EncryptionPropertiesDescriptionResponse { return v.Encryption }).(EncryptionPropertiesDescriptionResponsePtrOutput)
}

func (o IotHubPropertiesResponseOutput) EventHubEndpoints() EventHubPropertiesResponseMapOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) map[string]EventHubPropertiesResponse { return v.EventHubEndpoints }).(EventHubPropertiesResponseMapOutput)
}

func (o IotHubPropertiesResponseOutput) Features() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *string { return v.Features }).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesResponseOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) string { return v.HostName }).(pulumi.StringOutput)
}

func (o IotHubPropertiesResponseOutput) IpFilterRules() IpFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) []IpFilterRuleResponse { return v.IpFilterRules }).(IpFilterRuleResponseArrayOutput)
}

func (o IotHubPropertiesResponseOutput) Locations() IotHubLocationDescriptionResponseArrayOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) []IotHubLocationDescriptionResponse { return v.Locations }).(IotHubLocationDescriptionResponseArrayOutput)
}

func (o IotHubPropertiesResponseOutput) MessagingEndpoints() MessagingEndpointPropertiesResponseMapOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) map[string]MessagingEndpointPropertiesResponse {
		return v.MessagingEndpoints
	}).(MessagingEndpointPropertiesResponseMapOutput)
}

func (o IotHubPropertiesResponseOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *string { return v.MinTlsVersion }).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesResponseOutput) NetworkRuleSets() NetworkRuleSetPropertiesResponsePtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *NetworkRuleSetPropertiesResponse { return v.NetworkRuleSets }).(NetworkRuleSetPropertiesResponsePtrOutput)
}

func (o IotHubPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o IotHubPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o IotHubPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesResponseOutput) Routing() RoutingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *RoutingPropertiesResponse { return v.Routing }).(RoutingPropertiesResponsePtrOutput)
}

func (o IotHubPropertiesResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o IotHubPropertiesResponseOutput) StorageEndpoints() StorageEndpointPropertiesResponseMapOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) map[string]StorageEndpointPropertiesResponse {
		return v.StorageEndpoints
	}).(StorageEndpointPropertiesResponseMapOutput)
}

type IotHubPropertiesResponseDeviceStreams struct {
	StreamingEndpoints []string `pulumi:"streamingEndpoints"`
}

type IotHubPropertiesResponseDeviceStreamsOutput struct{ *pulumi.OutputState }

func (IotHubPropertiesResponseDeviceStreamsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubPropertiesResponseDeviceStreams)(nil)).Elem()
}

func (o IotHubPropertiesResponseDeviceStreamsOutput) ToIotHubPropertiesResponseDeviceStreamsOutput() IotHubPropertiesResponseDeviceStreamsOutput {
	return o
}

func (o IotHubPropertiesResponseDeviceStreamsOutput) ToIotHubPropertiesResponseDeviceStreamsOutputWithContext(ctx context.Context) IotHubPropertiesResponseDeviceStreamsOutput {
	return o
}

func (o IotHubPropertiesResponseDeviceStreamsOutput) StreamingEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IotHubPropertiesResponseDeviceStreams) []string { return v.StreamingEndpoints }).(pulumi.StringArrayOutput)
}

type IotHubPropertiesResponseDeviceStreamsPtrOutput struct{ *pulumi.OutputState }

func (IotHubPropertiesResponseDeviceStreamsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubPropertiesResponseDeviceStreams)(nil)).Elem()
}

func (o IotHubPropertiesResponseDeviceStreamsPtrOutput) ToIotHubPropertiesResponseDeviceStreamsPtrOutput() IotHubPropertiesResponseDeviceStreamsPtrOutput {
	return o
}

func (o IotHubPropertiesResponseDeviceStreamsPtrOutput) ToIotHubPropertiesResponseDeviceStreamsPtrOutputWithContext(ctx context.Context) IotHubPropertiesResponseDeviceStreamsPtrOutput {
	return o
}

func (o IotHubPropertiesResponseDeviceStreamsPtrOutput) Elem() IotHubPropertiesResponseDeviceStreamsOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponseDeviceStreams) IotHubPropertiesResponseDeviceStreams {
		if v != nil {
			return *v
		}
		var ret IotHubPropertiesResponseDeviceStreams
		return ret
	}).(IotHubPropertiesResponseDeviceStreamsOutput)
}

func (o IotHubPropertiesResponseDeviceStreamsPtrOutput) StreamingEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponseDeviceStreams) []string {
		if v == nil {
			return nil
		}
		return v.StreamingEndpoints
	}).(pulumi.StringArrayOutput)
}

type IotHubSkuInfo struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     string   `pulumi:"name"`
}





type IotHubSkuInfoInput interface {
	pulumi.Input

	ToIotHubSkuInfoOutput() IotHubSkuInfoOutput
	ToIotHubSkuInfoOutputWithContext(context.Context) IotHubSkuInfoOutput
}

type IotHubSkuInfoArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringInput     `pulumi:"name"`
}

func (IotHubSkuInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSkuInfo)(nil)).Elem()
}

func (i IotHubSkuInfoArgs) ToIotHubSkuInfoOutput() IotHubSkuInfoOutput {
	return i.ToIotHubSkuInfoOutputWithContext(context.Background())
}

func (i IotHubSkuInfoArgs) ToIotHubSkuInfoOutputWithContext(ctx context.Context) IotHubSkuInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSkuInfoOutput)
}

type IotHubSkuInfoOutput struct{ *pulumi.OutputState }

func (IotHubSkuInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSkuInfo)(nil)).Elem()
}

func (o IotHubSkuInfoOutput) ToIotHubSkuInfoOutput() IotHubSkuInfoOutput {
	return o
}

func (o IotHubSkuInfoOutput) ToIotHubSkuInfoOutputWithContext(ctx context.Context) IotHubSkuInfoOutput {
	return o
}

func (o IotHubSkuInfoOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v IotHubSkuInfo) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o IotHubSkuInfoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubSkuInfo) string { return v.Name }).(pulumi.StringOutput)
}

type IotHubSkuInfoResponse struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     string   `pulumi:"name"`
	Tier     string   `pulumi:"tier"`
}

type IotHubSkuInfoResponseOutput struct{ *pulumi.OutputState }

func (IotHubSkuInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSkuInfoResponse)(nil)).Elem()
}

func (o IotHubSkuInfoResponseOutput) ToIotHubSkuInfoResponseOutput() IotHubSkuInfoResponseOutput {
	return o
}

func (o IotHubSkuInfoResponseOutput) ToIotHubSkuInfoResponseOutputWithContext(ctx context.Context) IotHubSkuInfoResponseOutput {
	return o
}

func (o IotHubSkuInfoResponseOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v IotHubSkuInfoResponse) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o IotHubSkuInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubSkuInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o IotHubSkuInfoResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubSkuInfoResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type IpFilterRule struct {
	Action     IpFilterActionType `pulumi:"action"`
	FilterName string             `pulumi:"filterName"`
	IpMask     string             `pulumi:"ipMask"`
}





type IpFilterRuleInput interface {
	pulumi.Input

	ToIpFilterRuleOutput() IpFilterRuleOutput
	ToIpFilterRuleOutputWithContext(context.Context) IpFilterRuleOutput
}

type IpFilterRuleArgs struct {
	Action     IpFilterActionTypeInput `pulumi:"action"`
	FilterName pulumi.StringInput      `pulumi:"filterName"`
	IpMask     pulumi.StringInput      `pulumi:"ipMask"`
}

func (IpFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterRule)(nil)).Elem()
}

func (i IpFilterRuleArgs) ToIpFilterRuleOutput() IpFilterRuleOutput {
	return i.ToIpFilterRuleOutputWithContext(context.Background())
}

func (i IpFilterRuleArgs) ToIpFilterRuleOutputWithContext(ctx context.Context) IpFilterRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpFilterRuleOutput)
}





type IpFilterRuleArrayInput interface {
	pulumi.Input

	ToIpFilterRuleArrayOutput() IpFilterRuleArrayOutput
	ToIpFilterRuleArrayOutputWithContext(context.Context) IpFilterRuleArrayOutput
}

type IpFilterRuleArray []IpFilterRuleInput

func (IpFilterRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpFilterRule)(nil)).Elem()
}

func (i IpFilterRuleArray) ToIpFilterRuleArrayOutput() IpFilterRuleArrayOutput {
	return i.ToIpFilterRuleArrayOutputWithContext(context.Background())
}

func (i IpFilterRuleArray) ToIpFilterRuleArrayOutputWithContext(ctx context.Context) IpFilterRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpFilterRuleArrayOutput)
}

type IpFilterRuleOutput struct{ *pulumi.OutputState }

func (IpFilterRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterRule)(nil)).Elem()
}

func (o IpFilterRuleOutput) ToIpFilterRuleOutput() IpFilterRuleOutput {
	return o
}

func (o IpFilterRuleOutput) ToIpFilterRuleOutputWithContext(ctx context.Context) IpFilterRuleOutput {
	return o
}

func (o IpFilterRuleOutput) Action() IpFilterActionTypeOutput {
	return o.ApplyT(func(v IpFilterRule) IpFilterActionType { return v.Action }).(IpFilterActionTypeOutput)
}

func (o IpFilterRuleOutput) FilterName() pulumi.StringOutput {
	return o.ApplyT(func(v IpFilterRule) string { return v.FilterName }).(pulumi.StringOutput)
}

func (o IpFilterRuleOutput) IpMask() pulumi.StringOutput {
	return o.ApplyT(func(v IpFilterRule) string { return v.IpMask }).(pulumi.StringOutput)
}

type IpFilterRuleArrayOutput struct{ *pulumi.OutputState }

func (IpFilterRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpFilterRule)(nil)).Elem()
}

func (o IpFilterRuleArrayOutput) ToIpFilterRuleArrayOutput() IpFilterRuleArrayOutput {
	return o
}

func (o IpFilterRuleArrayOutput) ToIpFilterRuleArrayOutputWithContext(ctx context.Context) IpFilterRuleArrayOutput {
	return o
}

func (o IpFilterRuleArrayOutput) Index(i pulumi.IntInput) IpFilterRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpFilterRule {
		return vs[0].([]IpFilterRule)[vs[1].(int)]
	}).(IpFilterRuleOutput)
}

type IpFilterRuleResponse struct {
	Action     string `pulumi:"action"`
	FilterName string `pulumi:"filterName"`
	IpMask     string `pulumi:"ipMask"`
}

type IpFilterRuleResponseOutput struct{ *pulumi.OutputState }

func (IpFilterRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterRuleResponse)(nil)).Elem()
}

func (o IpFilterRuleResponseOutput) ToIpFilterRuleResponseOutput() IpFilterRuleResponseOutput {
	return o
}

func (o IpFilterRuleResponseOutput) ToIpFilterRuleResponseOutputWithContext(ctx context.Context) IpFilterRuleResponseOutput {
	return o
}

func (o IpFilterRuleResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v IpFilterRuleResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o IpFilterRuleResponseOutput) FilterName() pulumi.StringOutput {
	return o.ApplyT(func(v IpFilterRuleResponse) string { return v.FilterName }).(pulumi.StringOutput)
}

func (o IpFilterRuleResponseOutput) IpMask() pulumi.StringOutput {
	return o.ApplyT(func(v IpFilterRuleResponse) string { return v.IpMask }).(pulumi.StringOutput)
}

type IpFilterRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IpFilterRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpFilterRuleResponse)(nil)).Elem()
}

func (o IpFilterRuleResponseArrayOutput) ToIpFilterRuleResponseArrayOutput() IpFilterRuleResponseArrayOutput {
	return o
}

func (o IpFilterRuleResponseArrayOutput) ToIpFilterRuleResponseArrayOutputWithContext(ctx context.Context) IpFilterRuleResponseArrayOutput {
	return o
}

func (o IpFilterRuleResponseArrayOutput) Index(i pulumi.IntInput) IpFilterRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpFilterRuleResponse {
		return vs[0].([]IpFilterRuleResponse)[vs[1].(int)]
	}).(IpFilterRuleResponseOutput)
}

type KEKIdentity struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}





type KEKIdentityInput interface {
	pulumi.Input

	ToKEKIdentityOutput() KEKIdentityOutput
	ToKEKIdentityOutputWithContext(context.Context) KEKIdentityOutput
}

type KEKIdentityArgs struct {
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (KEKIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KEKIdentity)(nil)).Elem()
}

func (i KEKIdentityArgs) ToKEKIdentityOutput() KEKIdentityOutput {
	return i.ToKEKIdentityOutputWithContext(context.Background())
}

func (i KEKIdentityArgs) ToKEKIdentityOutputWithContext(ctx context.Context) KEKIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KEKIdentityOutput)
}

func (i KEKIdentityArgs) ToKEKIdentityPtrOutput() KEKIdentityPtrOutput {
	return i.ToKEKIdentityPtrOutputWithContext(context.Background())
}

func (i KEKIdentityArgs) ToKEKIdentityPtrOutputWithContext(ctx context.Context) KEKIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KEKIdentityOutput).ToKEKIdentityPtrOutputWithContext(ctx)
}









type KEKIdentityPtrInput interface {
	pulumi.Input

	ToKEKIdentityPtrOutput() KEKIdentityPtrOutput
	ToKEKIdentityPtrOutputWithContext(context.Context) KEKIdentityPtrOutput
}

type kekidentityPtrType KEKIdentityArgs

func KEKIdentityPtr(v *KEKIdentityArgs) KEKIdentityPtrInput {
	return (*kekidentityPtrType)(v)
}

func (*kekidentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KEKIdentity)(nil)).Elem()
}

func (i *kekidentityPtrType) ToKEKIdentityPtrOutput() KEKIdentityPtrOutput {
	return i.ToKEKIdentityPtrOutputWithContext(context.Background())
}

func (i *kekidentityPtrType) ToKEKIdentityPtrOutputWithContext(ctx context.Context) KEKIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KEKIdentityPtrOutput)
}

type KEKIdentityOutput struct{ *pulumi.OutputState }

func (KEKIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KEKIdentity)(nil)).Elem()
}

func (o KEKIdentityOutput) ToKEKIdentityOutput() KEKIdentityOutput {
	return o
}

func (o KEKIdentityOutput) ToKEKIdentityOutputWithContext(ctx context.Context) KEKIdentityOutput {
	return o
}

func (o KEKIdentityOutput) ToKEKIdentityPtrOutput() KEKIdentityPtrOutput {
	return o.ToKEKIdentityPtrOutputWithContext(context.Background())
}

func (o KEKIdentityOutput) ToKEKIdentityPtrOutputWithContext(ctx context.Context) KEKIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KEKIdentity) *KEKIdentity {
		return &v
	}).(KEKIdentityPtrOutput)
}

func (o KEKIdentityOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KEKIdentity) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type KEKIdentityPtrOutput struct{ *pulumi.OutputState }

func (KEKIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KEKIdentity)(nil)).Elem()
}

func (o KEKIdentityPtrOutput) ToKEKIdentityPtrOutput() KEKIdentityPtrOutput {
	return o
}

func (o KEKIdentityPtrOutput) ToKEKIdentityPtrOutputWithContext(ctx context.Context) KEKIdentityPtrOutput {
	return o
}

func (o KEKIdentityPtrOutput) Elem() KEKIdentityOutput {
	return o.ApplyT(func(v *KEKIdentity) KEKIdentity {
		if v != nil {
			return *v
		}
		var ret KEKIdentity
		return ret
	}).(KEKIdentityOutput)
}

func (o KEKIdentityPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KEKIdentity) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type KEKIdentityResponse struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}

type KEKIdentityResponseOutput struct{ *pulumi.OutputState }

func (KEKIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KEKIdentityResponse)(nil)).Elem()
}

func (o KEKIdentityResponseOutput) ToKEKIdentityResponseOutput() KEKIdentityResponseOutput {
	return o
}

func (o KEKIdentityResponseOutput) ToKEKIdentityResponseOutputWithContext(ctx context.Context) KEKIdentityResponseOutput {
	return o
}

func (o KEKIdentityResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KEKIdentityResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type KEKIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (KEKIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KEKIdentityResponse)(nil)).Elem()
}

func (o KEKIdentityResponsePtrOutput) ToKEKIdentityResponsePtrOutput() KEKIdentityResponsePtrOutput {
	return o
}

func (o KEKIdentityResponsePtrOutput) ToKEKIdentityResponsePtrOutputWithContext(ctx context.Context) KEKIdentityResponsePtrOutput {
	return o
}

func (o KEKIdentityResponsePtrOutput) Elem() KEKIdentityResponseOutput {
	return o.ApplyT(func(v *KEKIdentityResponse) KEKIdentityResponse {
		if v != nil {
			return *v
		}
		var ret KEKIdentityResponse
		return ret
	}).(KEKIdentityResponseOutput)
}

func (o KEKIdentityResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KEKIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type KeyVaultKeyProperties struct {
	Identity      *KEKIdentity `pulumi:"identity"`
	KeyIdentifier *string      `pulumi:"keyIdentifier"`
}





type KeyVaultKeyPropertiesInput interface {
	pulumi.Input

	ToKeyVaultKeyPropertiesOutput() KeyVaultKeyPropertiesOutput
	ToKeyVaultKeyPropertiesOutputWithContext(context.Context) KeyVaultKeyPropertiesOutput
}

type KeyVaultKeyPropertiesArgs struct {
	Identity      KEKIdentityPtrInput   `pulumi:"identity"`
	KeyIdentifier pulumi.StringPtrInput `pulumi:"keyIdentifier"`
}

func (KeyVaultKeyPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyProperties)(nil)).Elem()
}

func (i KeyVaultKeyPropertiesArgs) ToKeyVaultKeyPropertiesOutput() KeyVaultKeyPropertiesOutput {
	return i.ToKeyVaultKeyPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultKeyPropertiesArgs) ToKeyVaultKeyPropertiesOutputWithContext(ctx context.Context) KeyVaultKeyPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyPropertiesOutput)
}





type KeyVaultKeyPropertiesArrayInput interface {
	pulumi.Input

	ToKeyVaultKeyPropertiesArrayOutput() KeyVaultKeyPropertiesArrayOutput
	ToKeyVaultKeyPropertiesArrayOutputWithContext(context.Context) KeyVaultKeyPropertiesArrayOutput
}

type KeyVaultKeyPropertiesArray []KeyVaultKeyPropertiesInput

func (KeyVaultKeyPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultKeyProperties)(nil)).Elem()
}

func (i KeyVaultKeyPropertiesArray) ToKeyVaultKeyPropertiesArrayOutput() KeyVaultKeyPropertiesArrayOutput {
	return i.ToKeyVaultKeyPropertiesArrayOutputWithContext(context.Background())
}

func (i KeyVaultKeyPropertiesArray) ToKeyVaultKeyPropertiesArrayOutputWithContext(ctx context.Context) KeyVaultKeyPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyPropertiesArrayOutput)
}

type KeyVaultKeyPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyProperties)(nil)).Elem()
}

func (o KeyVaultKeyPropertiesOutput) ToKeyVaultKeyPropertiesOutput() KeyVaultKeyPropertiesOutput {
	return o
}

func (o KeyVaultKeyPropertiesOutput) ToKeyVaultKeyPropertiesOutputWithContext(ctx context.Context) KeyVaultKeyPropertiesOutput {
	return o
}

func (o KeyVaultKeyPropertiesOutput) Identity() KEKIdentityPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyProperties) *KEKIdentity { return v.Identity }).(KEKIdentityPtrOutput)
}

func (o KeyVaultKeyPropertiesOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyProperties) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyPropertiesArrayOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultKeyProperties)(nil)).Elem()
}

func (o KeyVaultKeyPropertiesArrayOutput) ToKeyVaultKeyPropertiesArrayOutput() KeyVaultKeyPropertiesArrayOutput {
	return o
}

func (o KeyVaultKeyPropertiesArrayOutput) ToKeyVaultKeyPropertiesArrayOutputWithContext(ctx context.Context) KeyVaultKeyPropertiesArrayOutput {
	return o
}

func (o KeyVaultKeyPropertiesArrayOutput) Index(i pulumi.IntInput) KeyVaultKeyPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyVaultKeyProperties {
		return vs[0].([]KeyVaultKeyProperties)[vs[1].(int)]
	}).(KeyVaultKeyPropertiesOutput)
}

type KeyVaultKeyPropertiesResponse struct {
	Identity      *KEKIdentityResponse `pulumi:"identity"`
	KeyIdentifier *string              `pulumi:"keyIdentifier"`
}

type KeyVaultKeyPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultKeyPropertiesResponseOutput) ToKeyVaultKeyPropertiesResponseOutput() KeyVaultKeyPropertiesResponseOutput {
	return o
}

func (o KeyVaultKeyPropertiesResponseOutput) ToKeyVaultKeyPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultKeyPropertiesResponseOutput {
	return o
}

func (o KeyVaultKeyPropertiesResponseOutput) Identity() KEKIdentityResponsePtrOutput {
	return o.ApplyT(func(v KeyVaultKeyPropertiesResponse) *KEKIdentityResponse { return v.Identity }).(KEKIdentityResponsePtrOutput)
}

func (o KeyVaultKeyPropertiesResponseOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyPropertiesResponse) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultKeyPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultKeyPropertiesResponseArrayOutput) ToKeyVaultKeyPropertiesResponseArrayOutput() KeyVaultKeyPropertiesResponseArrayOutput {
	return o
}

func (o KeyVaultKeyPropertiesResponseArrayOutput) ToKeyVaultKeyPropertiesResponseArrayOutputWithContext(ctx context.Context) KeyVaultKeyPropertiesResponseArrayOutput {
	return o
}

func (o KeyVaultKeyPropertiesResponseArrayOutput) Index(i pulumi.IntInput) KeyVaultKeyPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyVaultKeyPropertiesResponse {
		return vs[0].([]KeyVaultKeyPropertiesResponse)[vs[1].(int)]
	}).(KeyVaultKeyPropertiesResponseOutput)
}

type MessagingEndpointProperties struct {
	LockDurationAsIso8601 *string `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      *int    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          *string `pulumi:"ttlAsIso8601"`
}





type MessagingEndpointPropertiesInput interface {
	pulumi.Input

	ToMessagingEndpointPropertiesOutput() MessagingEndpointPropertiesOutput
	ToMessagingEndpointPropertiesOutputWithContext(context.Context) MessagingEndpointPropertiesOutput
}

type MessagingEndpointPropertiesArgs struct {
	LockDurationAsIso8601 pulumi.StringPtrInput `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      pulumi.IntPtrInput    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          pulumi.StringPtrInput `pulumi:"ttlAsIso8601"`
}

func (MessagingEndpointPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MessagingEndpointProperties)(nil)).Elem()
}

func (i MessagingEndpointPropertiesArgs) ToMessagingEndpointPropertiesOutput() MessagingEndpointPropertiesOutput {
	return i.ToMessagingEndpointPropertiesOutputWithContext(context.Background())
}

func (i MessagingEndpointPropertiesArgs) ToMessagingEndpointPropertiesOutputWithContext(ctx context.Context) MessagingEndpointPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessagingEndpointPropertiesOutput)
}





type MessagingEndpointPropertiesMapInput interface {
	pulumi.Input

	ToMessagingEndpointPropertiesMapOutput() MessagingEndpointPropertiesMapOutput
	ToMessagingEndpointPropertiesMapOutputWithContext(context.Context) MessagingEndpointPropertiesMapOutput
}

type MessagingEndpointPropertiesMap map[string]MessagingEndpointPropertiesInput

func (MessagingEndpointPropertiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MessagingEndpointProperties)(nil)).Elem()
}

func (i MessagingEndpointPropertiesMap) ToMessagingEndpointPropertiesMapOutput() MessagingEndpointPropertiesMapOutput {
	return i.ToMessagingEndpointPropertiesMapOutputWithContext(context.Background())
}

func (i MessagingEndpointPropertiesMap) ToMessagingEndpointPropertiesMapOutputWithContext(ctx context.Context) MessagingEndpointPropertiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessagingEndpointPropertiesMapOutput)
}

type MessagingEndpointPropertiesOutput struct{ *pulumi.OutputState }

func (MessagingEndpointPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessagingEndpointProperties)(nil)).Elem()
}

func (o MessagingEndpointPropertiesOutput) ToMessagingEndpointPropertiesOutput() MessagingEndpointPropertiesOutput {
	return o
}

func (o MessagingEndpointPropertiesOutput) ToMessagingEndpointPropertiesOutputWithContext(ctx context.Context) MessagingEndpointPropertiesOutput {
	return o
}

func (o MessagingEndpointPropertiesOutput) LockDurationAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MessagingEndpointProperties) *string { return v.LockDurationAsIso8601 }).(pulumi.StringPtrOutput)
}

func (o MessagingEndpointPropertiesOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MessagingEndpointProperties) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

func (o MessagingEndpointPropertiesOutput) TtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MessagingEndpointProperties) *string { return v.TtlAsIso8601 }).(pulumi.StringPtrOutput)
}

type MessagingEndpointPropertiesMapOutput struct{ *pulumi.OutputState }

func (MessagingEndpointPropertiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MessagingEndpointProperties)(nil)).Elem()
}

func (o MessagingEndpointPropertiesMapOutput) ToMessagingEndpointPropertiesMapOutput() MessagingEndpointPropertiesMapOutput {
	return o
}

func (o MessagingEndpointPropertiesMapOutput) ToMessagingEndpointPropertiesMapOutputWithContext(ctx context.Context) MessagingEndpointPropertiesMapOutput {
	return o
}

func (o MessagingEndpointPropertiesMapOutput) MapIndex(k pulumi.StringInput) MessagingEndpointPropertiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MessagingEndpointProperties {
		return vs[0].(map[string]MessagingEndpointProperties)[vs[1].(string)]
	}).(MessagingEndpointPropertiesOutput)
}

type MessagingEndpointPropertiesResponse struct {
	LockDurationAsIso8601 *string `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      *int    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          *string `pulumi:"ttlAsIso8601"`
}

type MessagingEndpointPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MessagingEndpointPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessagingEndpointPropertiesResponse)(nil)).Elem()
}

func (o MessagingEndpointPropertiesResponseOutput) ToMessagingEndpointPropertiesResponseOutput() MessagingEndpointPropertiesResponseOutput {
	return o
}

func (o MessagingEndpointPropertiesResponseOutput) ToMessagingEndpointPropertiesResponseOutputWithContext(ctx context.Context) MessagingEndpointPropertiesResponseOutput {
	return o
}

func (o MessagingEndpointPropertiesResponseOutput) LockDurationAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MessagingEndpointPropertiesResponse) *string { return v.LockDurationAsIso8601 }).(pulumi.StringPtrOutput)
}

func (o MessagingEndpointPropertiesResponseOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MessagingEndpointPropertiesResponse) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

func (o MessagingEndpointPropertiesResponseOutput) TtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MessagingEndpointPropertiesResponse) *string { return v.TtlAsIso8601 }).(pulumi.StringPtrOutput)
}

type MessagingEndpointPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (MessagingEndpointPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MessagingEndpointPropertiesResponse)(nil)).Elem()
}

func (o MessagingEndpointPropertiesResponseMapOutput) ToMessagingEndpointPropertiesResponseMapOutput() MessagingEndpointPropertiesResponseMapOutput {
	return o
}

func (o MessagingEndpointPropertiesResponseMapOutput) ToMessagingEndpointPropertiesResponseMapOutputWithContext(ctx context.Context) MessagingEndpointPropertiesResponseMapOutput {
	return o
}

func (o MessagingEndpointPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) MessagingEndpointPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MessagingEndpointPropertiesResponse {
		return vs[0].(map[string]MessagingEndpointPropertiesResponse)[vs[1].(string)]
	}).(MessagingEndpointPropertiesResponseOutput)
}

type NetworkRuleSetIpRule struct {
	Action     *string `pulumi:"action"`
	FilterName string  `pulumi:"filterName"`
	IpMask     string  `pulumi:"ipMask"`
}


func (val *NetworkRuleSetIpRule) Defaults() *NetworkRuleSetIpRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type NetworkRuleSetIpRuleInput interface {
	pulumi.Input

	ToNetworkRuleSetIpRuleOutput() NetworkRuleSetIpRuleOutput
	ToNetworkRuleSetIpRuleOutputWithContext(context.Context) NetworkRuleSetIpRuleOutput
}

type NetworkRuleSetIpRuleArgs struct {
	Action     pulumi.StringPtrInput `pulumi:"action"`
	FilterName pulumi.StringInput    `pulumi:"filterName"`
	IpMask     pulumi.StringInput    `pulumi:"ipMask"`
}

func (NetworkRuleSetIpRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetIpRule)(nil)).Elem()
}

func (i NetworkRuleSetIpRuleArgs) ToNetworkRuleSetIpRuleOutput() NetworkRuleSetIpRuleOutput {
	return i.ToNetworkRuleSetIpRuleOutputWithContext(context.Background())
}

func (i NetworkRuleSetIpRuleArgs) ToNetworkRuleSetIpRuleOutputWithContext(ctx context.Context) NetworkRuleSetIpRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetIpRuleOutput)
}





type NetworkRuleSetIpRuleArrayInput interface {
	pulumi.Input

	ToNetworkRuleSetIpRuleArrayOutput() NetworkRuleSetIpRuleArrayOutput
	ToNetworkRuleSetIpRuleArrayOutputWithContext(context.Context) NetworkRuleSetIpRuleArrayOutput
}

type NetworkRuleSetIpRuleArray []NetworkRuleSetIpRuleInput

func (NetworkRuleSetIpRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkRuleSetIpRule)(nil)).Elem()
}

func (i NetworkRuleSetIpRuleArray) ToNetworkRuleSetIpRuleArrayOutput() NetworkRuleSetIpRuleArrayOutput {
	return i.ToNetworkRuleSetIpRuleArrayOutputWithContext(context.Background())
}

func (i NetworkRuleSetIpRuleArray) ToNetworkRuleSetIpRuleArrayOutputWithContext(ctx context.Context) NetworkRuleSetIpRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetIpRuleArrayOutput)
}

type NetworkRuleSetIpRuleOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetIpRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetIpRule)(nil)).Elem()
}

func (o NetworkRuleSetIpRuleOutput) ToNetworkRuleSetIpRuleOutput() NetworkRuleSetIpRuleOutput {
	return o
}

func (o NetworkRuleSetIpRuleOutput) ToNetworkRuleSetIpRuleOutputWithContext(ctx context.Context) NetworkRuleSetIpRuleOutput {
	return o
}

func (o NetworkRuleSetIpRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetIpRule) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetIpRuleOutput) FilterName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSetIpRule) string { return v.FilterName }).(pulumi.StringOutput)
}

func (o NetworkRuleSetIpRuleOutput) IpMask() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSetIpRule) string { return v.IpMask }).(pulumi.StringOutput)
}

type NetworkRuleSetIpRuleArrayOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetIpRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkRuleSetIpRule)(nil)).Elem()
}

func (o NetworkRuleSetIpRuleArrayOutput) ToNetworkRuleSetIpRuleArrayOutput() NetworkRuleSetIpRuleArrayOutput {
	return o
}

func (o NetworkRuleSetIpRuleArrayOutput) ToNetworkRuleSetIpRuleArrayOutputWithContext(ctx context.Context) NetworkRuleSetIpRuleArrayOutput {
	return o
}

func (o NetworkRuleSetIpRuleArrayOutput) Index(i pulumi.IntInput) NetworkRuleSetIpRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkRuleSetIpRule {
		return vs[0].([]NetworkRuleSetIpRule)[vs[1].(int)]
	}).(NetworkRuleSetIpRuleOutput)
}

type NetworkRuleSetIpRuleResponse struct {
	Action     *string `pulumi:"action"`
	FilterName string  `pulumi:"filterName"`
	IpMask     string  `pulumi:"ipMask"`
}


func (val *NetworkRuleSetIpRuleResponse) Defaults() *NetworkRuleSetIpRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}

type NetworkRuleSetIpRuleResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetIpRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetIpRuleResponse)(nil)).Elem()
}

func (o NetworkRuleSetIpRuleResponseOutput) ToNetworkRuleSetIpRuleResponseOutput() NetworkRuleSetIpRuleResponseOutput {
	return o
}

func (o NetworkRuleSetIpRuleResponseOutput) ToNetworkRuleSetIpRuleResponseOutputWithContext(ctx context.Context) NetworkRuleSetIpRuleResponseOutput {
	return o
}

func (o NetworkRuleSetIpRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetIpRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetIpRuleResponseOutput) FilterName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSetIpRuleResponse) string { return v.FilterName }).(pulumi.StringOutput)
}

func (o NetworkRuleSetIpRuleResponseOutput) IpMask() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSetIpRuleResponse) string { return v.IpMask }).(pulumi.StringOutput)
}

type NetworkRuleSetIpRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetIpRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkRuleSetIpRuleResponse)(nil)).Elem()
}

func (o NetworkRuleSetIpRuleResponseArrayOutput) ToNetworkRuleSetIpRuleResponseArrayOutput() NetworkRuleSetIpRuleResponseArrayOutput {
	return o
}

func (o NetworkRuleSetIpRuleResponseArrayOutput) ToNetworkRuleSetIpRuleResponseArrayOutputWithContext(ctx context.Context) NetworkRuleSetIpRuleResponseArrayOutput {
	return o
}

func (o NetworkRuleSetIpRuleResponseArrayOutput) Index(i pulumi.IntInput) NetworkRuleSetIpRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkRuleSetIpRuleResponse {
		return vs[0].([]NetworkRuleSetIpRuleResponse)[vs[1].(int)]
	}).(NetworkRuleSetIpRuleResponseOutput)
}

type NetworkRuleSetProperties struct {
	ApplyToBuiltInEventHubEndpoint bool                   `pulumi:"applyToBuiltInEventHubEndpoint"`
	DefaultAction                  *string                `pulumi:"defaultAction"`
	IpRules                        []NetworkRuleSetIpRule `pulumi:"ipRules"`
}


func (val *NetworkRuleSetProperties) Defaults() *NetworkRuleSetProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultAction) {
		defaultAction_ := "Deny"
		tmp.DefaultAction = &defaultAction_
	}
	return &tmp
}





type NetworkRuleSetPropertiesInput interface {
	pulumi.Input

	ToNetworkRuleSetPropertiesOutput() NetworkRuleSetPropertiesOutput
	ToNetworkRuleSetPropertiesOutputWithContext(context.Context) NetworkRuleSetPropertiesOutput
}

type NetworkRuleSetPropertiesArgs struct {
	ApplyToBuiltInEventHubEndpoint pulumi.BoolInput               `pulumi:"applyToBuiltInEventHubEndpoint"`
	DefaultAction                  pulumi.StringPtrInput          `pulumi:"defaultAction"`
	IpRules                        NetworkRuleSetIpRuleArrayInput `pulumi:"ipRules"`
}

func (NetworkRuleSetPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetProperties)(nil)).Elem()
}

func (i NetworkRuleSetPropertiesArgs) ToNetworkRuleSetPropertiesOutput() NetworkRuleSetPropertiesOutput {
	return i.ToNetworkRuleSetPropertiesOutputWithContext(context.Background())
}

func (i NetworkRuleSetPropertiesArgs) ToNetworkRuleSetPropertiesOutputWithContext(ctx context.Context) NetworkRuleSetPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPropertiesOutput)
}

func (i NetworkRuleSetPropertiesArgs) ToNetworkRuleSetPropertiesPtrOutput() NetworkRuleSetPropertiesPtrOutput {
	return i.ToNetworkRuleSetPropertiesPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetPropertiesArgs) ToNetworkRuleSetPropertiesPtrOutputWithContext(ctx context.Context) NetworkRuleSetPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPropertiesOutput).ToNetworkRuleSetPropertiesPtrOutputWithContext(ctx)
}









type NetworkRuleSetPropertiesPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPropertiesPtrOutput() NetworkRuleSetPropertiesPtrOutput
	ToNetworkRuleSetPropertiesPtrOutputWithContext(context.Context) NetworkRuleSetPropertiesPtrOutput
}

type networkRuleSetPropertiesPtrType NetworkRuleSetPropertiesArgs

func NetworkRuleSetPropertiesPtr(v *NetworkRuleSetPropertiesArgs) NetworkRuleSetPropertiesPtrInput {
	return (*networkRuleSetPropertiesPtrType)(v)
}

func (*networkRuleSetPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetProperties)(nil)).Elem()
}

func (i *networkRuleSetPropertiesPtrType) ToNetworkRuleSetPropertiesPtrOutput() NetworkRuleSetPropertiesPtrOutput {
	return i.ToNetworkRuleSetPropertiesPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPropertiesPtrType) ToNetworkRuleSetPropertiesPtrOutputWithContext(ctx context.Context) NetworkRuleSetPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPropertiesPtrOutput)
}

type NetworkRuleSetPropertiesOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetProperties)(nil)).Elem()
}

func (o NetworkRuleSetPropertiesOutput) ToNetworkRuleSetPropertiesOutput() NetworkRuleSetPropertiesOutput {
	return o
}

func (o NetworkRuleSetPropertiesOutput) ToNetworkRuleSetPropertiesOutputWithContext(ctx context.Context) NetworkRuleSetPropertiesOutput {
	return o
}

func (o NetworkRuleSetPropertiesOutput) ToNetworkRuleSetPropertiesPtrOutput() NetworkRuleSetPropertiesPtrOutput {
	return o.ToNetworkRuleSetPropertiesPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetPropertiesOutput) ToNetworkRuleSetPropertiesPtrOutputWithContext(ctx context.Context) NetworkRuleSetPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSetProperties) *NetworkRuleSetProperties {
		return &v
	}).(NetworkRuleSetPropertiesPtrOutput)
}

func (o NetworkRuleSetPropertiesOutput) ApplyToBuiltInEventHubEndpoint() pulumi.BoolOutput {
	return o.ApplyT(func(v NetworkRuleSetProperties) bool { return v.ApplyToBuiltInEventHubEndpoint }).(pulumi.BoolOutput)
}

func (o NetworkRuleSetPropertiesOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetProperties) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPropertiesOutput) IpRules() NetworkRuleSetIpRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetProperties) []NetworkRuleSetIpRule { return v.IpRules }).(NetworkRuleSetIpRuleArrayOutput)
}

type NetworkRuleSetPropertiesPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetProperties)(nil)).Elem()
}

func (o NetworkRuleSetPropertiesPtrOutput) ToNetworkRuleSetPropertiesPtrOutput() NetworkRuleSetPropertiesPtrOutput {
	return o
}

func (o NetworkRuleSetPropertiesPtrOutput) ToNetworkRuleSetPropertiesPtrOutputWithContext(ctx context.Context) NetworkRuleSetPropertiesPtrOutput {
	return o
}

func (o NetworkRuleSetPropertiesPtrOutput) Elem() NetworkRuleSetPropertiesOutput {
	return o.ApplyT(func(v *NetworkRuleSetProperties) NetworkRuleSetProperties {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetProperties
		return ret
	}).(NetworkRuleSetPropertiesOutput)
}

func (o NetworkRuleSetPropertiesPtrOutput) ApplyToBuiltInEventHubEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.ApplyToBuiltInEventHubEndpoint
	}).(pulumi.BoolPtrOutput)
}

func (o NetworkRuleSetPropertiesPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetProperties) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPropertiesPtrOutput) IpRules() NetworkRuleSetIpRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetProperties) []NetworkRuleSetIpRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(NetworkRuleSetIpRuleArrayOutput)
}

type NetworkRuleSetPropertiesResponse struct {
	ApplyToBuiltInEventHubEndpoint bool                           `pulumi:"applyToBuiltInEventHubEndpoint"`
	DefaultAction                  *string                        `pulumi:"defaultAction"`
	IpRules                        []NetworkRuleSetIpRuleResponse `pulumi:"ipRules"`
}


func (val *NetworkRuleSetPropertiesResponse) Defaults() *NetworkRuleSetPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultAction) {
		defaultAction_ := "Deny"
		tmp.DefaultAction = &defaultAction_
	}
	return &tmp
}

type NetworkRuleSetPropertiesResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetPropertiesResponse)(nil)).Elem()
}

func (o NetworkRuleSetPropertiesResponseOutput) ToNetworkRuleSetPropertiesResponseOutput() NetworkRuleSetPropertiesResponseOutput {
	return o
}

func (o NetworkRuleSetPropertiesResponseOutput) ToNetworkRuleSetPropertiesResponseOutputWithContext(ctx context.Context) NetworkRuleSetPropertiesResponseOutput {
	return o
}

func (o NetworkRuleSetPropertiesResponseOutput) ApplyToBuiltInEventHubEndpoint() pulumi.BoolOutput {
	return o.ApplyT(func(v NetworkRuleSetPropertiesResponse) bool { return v.ApplyToBuiltInEventHubEndpoint }).(pulumi.BoolOutput)
}

func (o NetworkRuleSetPropertiesResponseOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetPropertiesResponse) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPropertiesResponseOutput) IpRules() NetworkRuleSetIpRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetPropertiesResponse) []NetworkRuleSetIpRuleResponse { return v.IpRules }).(NetworkRuleSetIpRuleResponseArrayOutput)
}

type NetworkRuleSetPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetPropertiesResponse)(nil)).Elem()
}

func (o NetworkRuleSetPropertiesResponsePtrOutput) ToNetworkRuleSetPropertiesResponsePtrOutput() NetworkRuleSetPropertiesResponsePtrOutput {
	return o
}

func (o NetworkRuleSetPropertiesResponsePtrOutput) ToNetworkRuleSetPropertiesResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetPropertiesResponsePtrOutput {
	return o
}

func (o NetworkRuleSetPropertiesResponsePtrOutput) Elem() NetworkRuleSetPropertiesResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetPropertiesResponse) NetworkRuleSetPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetPropertiesResponse
		return ret
	}).(NetworkRuleSetPropertiesResponseOutput)
}

func (o NetworkRuleSetPropertiesResponsePtrOutput) ApplyToBuiltInEventHubEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ApplyToBuiltInEventHubEndpoint
	}).(pulumi.BoolPtrOutput)
}

func (o NetworkRuleSetPropertiesResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPropertiesResponsePtrOutput) IpRules() NetworkRuleSetIpRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetPropertiesResponse) []NetworkRuleSetIpRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(NetworkRuleSetIpRuleResponseArrayOutput)
}

type PrivateEndpointConnectionType struct {
	Properties PrivateEndpointConnectionProperties `pulumi:"properties"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	Properties PrivateEndpointConnectionPropertiesInput `pulumi:"properties"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}





type PrivateEndpointConnectionTypeArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput
	ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Context) PrivateEndpointConnectionTypeArrayOutput
}

type PrivateEndpointConnectionTypeArray []PrivateEndpointConnectionTypeInput

func (PrivateEndpointConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return i.ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeArrayOutput)
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) Properties() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) PrivateEndpointConnectionProperties { return v.Properties }).(PrivateEndpointConnectionPropertiesOutput)
}

type PrivateEndpointConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionType {
		return vs[0].([]PrivateEndpointConnectionType)[vs[1].(int)]
	}).(PrivateEndpointConnectionTypeOutput)
}

type PrivateEndpointConnectionProperties struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return i.ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput)
}

type PrivateEndpointConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	Type       string                                      `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     string  `pulumi:"description"`
	Status          string  `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput    `pulumi:"description"`
	Status          pulumi.StringInput    `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     string  `pulumi:"description"`
	Status          string  `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) string { return v.Status }).(pulumi.StringOutput)
}

type RouteProperties struct {
	Condition     *string  `pulumi:"condition"`
	EndpointNames []string `pulumi:"endpointNames"`
	IsEnabled     bool     `pulumi:"isEnabled"`
	Name          string   `pulumi:"name"`
	Source        string   `pulumi:"source"`
}





type RoutePropertiesInput interface {
	pulumi.Input

	ToRoutePropertiesOutput() RoutePropertiesOutput
	ToRoutePropertiesOutputWithContext(context.Context) RoutePropertiesOutput
}

type RoutePropertiesArgs struct {
	Condition     pulumi.StringPtrInput   `pulumi:"condition"`
	EndpointNames pulumi.StringArrayInput `pulumi:"endpointNames"`
	IsEnabled     pulumi.BoolInput        `pulumi:"isEnabled"`
	Name          pulumi.StringInput      `pulumi:"name"`
	Source        pulumi.StringInput      `pulumi:"source"`
}

func (RoutePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteProperties)(nil)).Elem()
}

func (i RoutePropertiesArgs) ToRoutePropertiesOutput() RoutePropertiesOutput {
	return i.ToRoutePropertiesOutputWithContext(context.Background())
}

func (i RoutePropertiesArgs) ToRoutePropertiesOutputWithContext(ctx context.Context) RoutePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePropertiesOutput)
}





type RoutePropertiesArrayInput interface {
	pulumi.Input

	ToRoutePropertiesArrayOutput() RoutePropertiesArrayOutput
	ToRoutePropertiesArrayOutputWithContext(context.Context) RoutePropertiesArrayOutput
}

type RoutePropertiesArray []RoutePropertiesInput

func (RoutePropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteProperties)(nil)).Elem()
}

func (i RoutePropertiesArray) ToRoutePropertiesArrayOutput() RoutePropertiesArrayOutput {
	return i.ToRoutePropertiesArrayOutputWithContext(context.Background())
}

func (i RoutePropertiesArray) ToRoutePropertiesArrayOutputWithContext(ctx context.Context) RoutePropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePropertiesArrayOutput)
}

type RoutePropertiesOutput struct{ *pulumi.OutputState }

func (RoutePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteProperties)(nil)).Elem()
}

func (o RoutePropertiesOutput) ToRoutePropertiesOutput() RoutePropertiesOutput {
	return o
}

func (o RoutePropertiesOutput) ToRoutePropertiesOutputWithContext(ctx context.Context) RoutePropertiesOutput {
	return o
}

func (o RoutePropertiesOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteProperties) *string { return v.Condition }).(pulumi.StringPtrOutput)
}

func (o RoutePropertiesOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RouteProperties) []string { return v.EndpointNames }).(pulumi.StringArrayOutput)
}

func (o RoutePropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RouteProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o RoutePropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RouteProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutePropertiesOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v RouteProperties) string { return v.Source }).(pulumi.StringOutput)
}

type RoutePropertiesArrayOutput struct{ *pulumi.OutputState }

func (RoutePropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteProperties)(nil)).Elem()
}

func (o RoutePropertiesArrayOutput) ToRoutePropertiesArrayOutput() RoutePropertiesArrayOutput {
	return o
}

func (o RoutePropertiesArrayOutput) ToRoutePropertiesArrayOutputWithContext(ctx context.Context) RoutePropertiesArrayOutput {
	return o
}

func (o RoutePropertiesArrayOutput) Index(i pulumi.IntInput) RoutePropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteProperties {
		return vs[0].([]RouteProperties)[vs[1].(int)]
	}).(RoutePropertiesOutput)
}

type RoutePropertiesResponse struct {
	Condition     *string  `pulumi:"condition"`
	EndpointNames []string `pulumi:"endpointNames"`
	IsEnabled     bool     `pulumi:"isEnabled"`
	Name          string   `pulumi:"name"`
	Source        string   `pulumi:"source"`
}

type RoutePropertiesResponseOutput struct{ *pulumi.OutputState }

func (RoutePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutePropertiesResponse)(nil)).Elem()
}

func (o RoutePropertiesResponseOutput) ToRoutePropertiesResponseOutput() RoutePropertiesResponseOutput {
	return o
}

func (o RoutePropertiesResponseOutput) ToRoutePropertiesResponseOutputWithContext(ctx context.Context) RoutePropertiesResponseOutput {
	return o
}

func (o RoutePropertiesResponseOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutePropertiesResponse) *string { return v.Condition }).(pulumi.StringPtrOutput)
}

func (o RoutePropertiesResponseOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoutePropertiesResponse) []string { return v.EndpointNames }).(pulumi.StringArrayOutput)
}

func (o RoutePropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RoutePropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o RoutePropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutePropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutePropertiesResponseOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v RoutePropertiesResponse) string { return v.Source }).(pulumi.StringOutput)
}

type RoutePropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutePropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutePropertiesResponse)(nil)).Elem()
}

func (o RoutePropertiesResponseArrayOutput) ToRoutePropertiesResponseArrayOutput() RoutePropertiesResponseArrayOutput {
	return o
}

func (o RoutePropertiesResponseArrayOutput) ToRoutePropertiesResponseArrayOutputWithContext(ctx context.Context) RoutePropertiesResponseArrayOutput {
	return o
}

func (o RoutePropertiesResponseArrayOutput) Index(i pulumi.IntInput) RoutePropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutePropertiesResponse {
		return vs[0].([]RoutePropertiesResponse)[vs[1].(int)]
	}).(RoutePropertiesResponseOutput)
}

type RoutingEndpoints struct {
	EventHubs         []RoutingEventHubProperties                `pulumi:"eventHubs"`
	ServiceBusQueues  []RoutingServiceBusQueueEndpointProperties `pulumi:"serviceBusQueues"`
	ServiceBusTopics  []RoutingServiceBusTopicEndpointProperties `pulumi:"serviceBusTopics"`
	StorageContainers []RoutingStorageContainerProperties        `pulumi:"storageContainers"`
}





type RoutingEndpointsInput interface {
	pulumi.Input

	ToRoutingEndpointsOutput() RoutingEndpointsOutput
	ToRoutingEndpointsOutputWithContext(context.Context) RoutingEndpointsOutput
}

type RoutingEndpointsArgs struct {
	EventHubs         RoutingEventHubPropertiesArrayInput                `pulumi:"eventHubs"`
	ServiceBusQueues  RoutingServiceBusQueueEndpointPropertiesArrayInput `pulumi:"serviceBusQueues"`
	ServiceBusTopics  RoutingServiceBusTopicEndpointPropertiesArrayInput `pulumi:"serviceBusTopics"`
	StorageContainers RoutingStorageContainerPropertiesArrayInput        `pulumi:"storageContainers"`
}

func (RoutingEndpointsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEndpoints)(nil)).Elem()
}

func (i RoutingEndpointsArgs) ToRoutingEndpointsOutput() RoutingEndpointsOutput {
	return i.ToRoutingEndpointsOutputWithContext(context.Background())
}

func (i RoutingEndpointsArgs) ToRoutingEndpointsOutputWithContext(ctx context.Context) RoutingEndpointsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEndpointsOutput)
}

func (i RoutingEndpointsArgs) ToRoutingEndpointsPtrOutput() RoutingEndpointsPtrOutput {
	return i.ToRoutingEndpointsPtrOutputWithContext(context.Background())
}

func (i RoutingEndpointsArgs) ToRoutingEndpointsPtrOutputWithContext(ctx context.Context) RoutingEndpointsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEndpointsOutput).ToRoutingEndpointsPtrOutputWithContext(ctx)
}









type RoutingEndpointsPtrInput interface {
	pulumi.Input

	ToRoutingEndpointsPtrOutput() RoutingEndpointsPtrOutput
	ToRoutingEndpointsPtrOutputWithContext(context.Context) RoutingEndpointsPtrOutput
}

type routingEndpointsPtrType RoutingEndpointsArgs

func RoutingEndpointsPtr(v *RoutingEndpointsArgs) RoutingEndpointsPtrInput {
	return (*routingEndpointsPtrType)(v)
}

func (*routingEndpointsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingEndpoints)(nil)).Elem()
}

func (i *routingEndpointsPtrType) ToRoutingEndpointsPtrOutput() RoutingEndpointsPtrOutput {
	return i.ToRoutingEndpointsPtrOutputWithContext(context.Background())
}

func (i *routingEndpointsPtrType) ToRoutingEndpointsPtrOutputWithContext(ctx context.Context) RoutingEndpointsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEndpointsPtrOutput)
}

type RoutingEndpointsOutput struct{ *pulumi.OutputState }

func (RoutingEndpointsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEndpoints)(nil)).Elem()
}

func (o RoutingEndpointsOutput) ToRoutingEndpointsOutput() RoutingEndpointsOutput {
	return o
}

func (o RoutingEndpointsOutput) ToRoutingEndpointsOutputWithContext(ctx context.Context) RoutingEndpointsOutput {
	return o
}

func (o RoutingEndpointsOutput) ToRoutingEndpointsPtrOutput() RoutingEndpointsPtrOutput {
	return o.ToRoutingEndpointsPtrOutputWithContext(context.Background())
}

func (o RoutingEndpointsOutput) ToRoutingEndpointsPtrOutputWithContext(ctx context.Context) RoutingEndpointsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoutingEndpoints) *RoutingEndpoints {
		return &v
	}).(RoutingEndpointsPtrOutput)
}

func (o RoutingEndpointsOutput) EventHubs() RoutingEventHubPropertiesArrayOutput {
	return o.ApplyT(func(v RoutingEndpoints) []RoutingEventHubProperties { return v.EventHubs }).(RoutingEventHubPropertiesArrayOutput)
}

func (o RoutingEndpointsOutput) ServiceBusQueues() RoutingServiceBusQueueEndpointPropertiesArrayOutput {
	return o.ApplyT(func(v RoutingEndpoints) []RoutingServiceBusQueueEndpointProperties { return v.ServiceBusQueues }).(RoutingServiceBusQueueEndpointPropertiesArrayOutput)
}

func (o RoutingEndpointsOutput) ServiceBusTopics() RoutingServiceBusTopicEndpointPropertiesArrayOutput {
	return o.ApplyT(func(v RoutingEndpoints) []RoutingServiceBusTopicEndpointProperties { return v.ServiceBusTopics }).(RoutingServiceBusTopicEndpointPropertiesArrayOutput)
}

func (o RoutingEndpointsOutput) StorageContainers() RoutingStorageContainerPropertiesArrayOutput {
	return o.ApplyT(func(v RoutingEndpoints) []RoutingStorageContainerProperties { return v.StorageContainers }).(RoutingStorageContainerPropertiesArrayOutput)
}

type RoutingEndpointsPtrOutput struct{ *pulumi.OutputState }

func (RoutingEndpointsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingEndpoints)(nil)).Elem()
}

func (o RoutingEndpointsPtrOutput) ToRoutingEndpointsPtrOutput() RoutingEndpointsPtrOutput {
	return o
}

func (o RoutingEndpointsPtrOutput) ToRoutingEndpointsPtrOutputWithContext(ctx context.Context) RoutingEndpointsPtrOutput {
	return o
}

func (o RoutingEndpointsPtrOutput) Elem() RoutingEndpointsOutput {
	return o.ApplyT(func(v *RoutingEndpoints) RoutingEndpoints {
		if v != nil {
			return *v
		}
		var ret RoutingEndpoints
		return ret
	}).(RoutingEndpointsOutput)
}

func (o RoutingEndpointsPtrOutput) EventHubs() RoutingEventHubPropertiesArrayOutput {
	return o.ApplyT(func(v *RoutingEndpoints) []RoutingEventHubProperties {
		if v == nil {
			return nil
		}
		return v.EventHubs
	}).(RoutingEventHubPropertiesArrayOutput)
}

func (o RoutingEndpointsPtrOutput) ServiceBusQueues() RoutingServiceBusQueueEndpointPropertiesArrayOutput {
	return o.ApplyT(func(v *RoutingEndpoints) []RoutingServiceBusQueueEndpointProperties {
		if v == nil {
			return nil
		}
		return v.ServiceBusQueues
	}).(RoutingServiceBusQueueEndpointPropertiesArrayOutput)
}

func (o RoutingEndpointsPtrOutput) ServiceBusTopics() RoutingServiceBusTopicEndpointPropertiesArrayOutput {
	return o.ApplyT(func(v *RoutingEndpoints) []RoutingServiceBusTopicEndpointProperties {
		if v == nil {
			return nil
		}
		return v.ServiceBusTopics
	}).(RoutingServiceBusTopicEndpointPropertiesArrayOutput)
}

func (o RoutingEndpointsPtrOutput) StorageContainers() RoutingStorageContainerPropertiesArrayOutput {
	return o.ApplyT(func(v *RoutingEndpoints) []RoutingStorageContainerProperties {
		if v == nil {
			return nil
		}
		return v.StorageContainers
	}).(RoutingStorageContainerPropertiesArrayOutput)
}

type RoutingEndpointsResponse struct {
	EventHubs         []RoutingEventHubPropertiesResponse                `pulumi:"eventHubs"`
	ServiceBusQueues  []RoutingServiceBusQueueEndpointPropertiesResponse `pulumi:"serviceBusQueues"`
	ServiceBusTopics  []RoutingServiceBusTopicEndpointPropertiesResponse `pulumi:"serviceBusTopics"`
	StorageContainers []RoutingStorageContainerPropertiesResponse        `pulumi:"storageContainers"`
}

type RoutingEndpointsResponseOutput struct{ *pulumi.OutputState }

func (RoutingEndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEndpointsResponse)(nil)).Elem()
}

func (o RoutingEndpointsResponseOutput) ToRoutingEndpointsResponseOutput() RoutingEndpointsResponseOutput {
	return o
}

func (o RoutingEndpointsResponseOutput) ToRoutingEndpointsResponseOutputWithContext(ctx context.Context) RoutingEndpointsResponseOutput {
	return o
}

func (o RoutingEndpointsResponseOutput) EventHubs() RoutingEventHubPropertiesResponseArrayOutput {
	return o.ApplyT(func(v RoutingEndpointsResponse) []RoutingEventHubPropertiesResponse { return v.EventHubs }).(RoutingEventHubPropertiesResponseArrayOutput)
}

func (o RoutingEndpointsResponseOutput) ServiceBusQueues() RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput {
	return o.ApplyT(func(v RoutingEndpointsResponse) []RoutingServiceBusQueueEndpointPropertiesResponse {
		return v.ServiceBusQueues
	}).(RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput)
}

func (o RoutingEndpointsResponseOutput) ServiceBusTopics() RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput {
	return o.ApplyT(func(v RoutingEndpointsResponse) []RoutingServiceBusTopicEndpointPropertiesResponse {
		return v.ServiceBusTopics
	}).(RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput)
}

func (o RoutingEndpointsResponseOutput) StorageContainers() RoutingStorageContainerPropertiesResponseArrayOutput {
	return o.ApplyT(func(v RoutingEndpointsResponse) []RoutingStorageContainerPropertiesResponse {
		return v.StorageContainers
	}).(RoutingStorageContainerPropertiesResponseArrayOutput)
}

type RoutingEndpointsResponsePtrOutput struct{ *pulumi.OutputState }

func (RoutingEndpointsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingEndpointsResponse)(nil)).Elem()
}

func (o RoutingEndpointsResponsePtrOutput) ToRoutingEndpointsResponsePtrOutput() RoutingEndpointsResponsePtrOutput {
	return o
}

func (o RoutingEndpointsResponsePtrOutput) ToRoutingEndpointsResponsePtrOutputWithContext(ctx context.Context) RoutingEndpointsResponsePtrOutput {
	return o
}

func (o RoutingEndpointsResponsePtrOutput) Elem() RoutingEndpointsResponseOutput {
	return o.ApplyT(func(v *RoutingEndpointsResponse) RoutingEndpointsResponse {
		if v != nil {
			return *v
		}
		var ret RoutingEndpointsResponse
		return ret
	}).(RoutingEndpointsResponseOutput)
}

func (o RoutingEndpointsResponsePtrOutput) EventHubs() RoutingEventHubPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *RoutingEndpointsResponse) []RoutingEventHubPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.EventHubs
	}).(RoutingEventHubPropertiesResponseArrayOutput)
}

func (o RoutingEndpointsResponsePtrOutput) ServiceBusQueues() RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *RoutingEndpointsResponse) []RoutingServiceBusQueueEndpointPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ServiceBusQueues
	}).(RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput)
}

func (o RoutingEndpointsResponsePtrOutput) ServiceBusTopics() RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *RoutingEndpointsResponse) []RoutingServiceBusTopicEndpointPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ServiceBusTopics
	}).(RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput)
}

func (o RoutingEndpointsResponsePtrOutput) StorageContainers() RoutingStorageContainerPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *RoutingEndpointsResponse) []RoutingStorageContainerPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.StorageContainers
	}).(RoutingStorageContainerPropertiesResponseArrayOutput)
}

type RoutingEventHubProperties struct {
	AuthenticationType *string `pulumi:"authenticationType"`
	ConnectionString   *string `pulumi:"connectionString"`
	EndpointUri        *string `pulumi:"endpointUri"`
	EntityPath         *string `pulumi:"entityPath"`
	Id                 *string `pulumi:"id"`
	Name               string  `pulumi:"name"`
	ResourceGroup      *string `pulumi:"resourceGroup"`
	SubscriptionId     *string `pulumi:"subscriptionId"`
}





type RoutingEventHubPropertiesInput interface {
	pulumi.Input

	ToRoutingEventHubPropertiesOutput() RoutingEventHubPropertiesOutput
	ToRoutingEventHubPropertiesOutputWithContext(context.Context) RoutingEventHubPropertiesOutput
}

type RoutingEventHubPropertiesArgs struct {
	AuthenticationType pulumi.StringPtrInput `pulumi:"authenticationType"`
	ConnectionString   pulumi.StringPtrInput `pulumi:"connectionString"`
	EndpointUri        pulumi.StringPtrInput `pulumi:"endpointUri"`
	EntityPath         pulumi.StringPtrInput `pulumi:"entityPath"`
	Id                 pulumi.StringPtrInput `pulumi:"id"`
	Name               pulumi.StringInput    `pulumi:"name"`
	ResourceGroup      pulumi.StringPtrInput `pulumi:"resourceGroup"`
	SubscriptionId     pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (RoutingEventHubPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEventHubProperties)(nil)).Elem()
}

func (i RoutingEventHubPropertiesArgs) ToRoutingEventHubPropertiesOutput() RoutingEventHubPropertiesOutput {
	return i.ToRoutingEventHubPropertiesOutputWithContext(context.Background())
}

func (i RoutingEventHubPropertiesArgs) ToRoutingEventHubPropertiesOutputWithContext(ctx context.Context) RoutingEventHubPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEventHubPropertiesOutput)
}





type RoutingEventHubPropertiesArrayInput interface {
	pulumi.Input

	ToRoutingEventHubPropertiesArrayOutput() RoutingEventHubPropertiesArrayOutput
	ToRoutingEventHubPropertiesArrayOutputWithContext(context.Context) RoutingEventHubPropertiesArrayOutput
}

type RoutingEventHubPropertiesArray []RoutingEventHubPropertiesInput

func (RoutingEventHubPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingEventHubProperties)(nil)).Elem()
}

func (i RoutingEventHubPropertiesArray) ToRoutingEventHubPropertiesArrayOutput() RoutingEventHubPropertiesArrayOutput {
	return i.ToRoutingEventHubPropertiesArrayOutputWithContext(context.Background())
}

func (i RoutingEventHubPropertiesArray) ToRoutingEventHubPropertiesArrayOutputWithContext(ctx context.Context) RoutingEventHubPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEventHubPropertiesArrayOutput)
}

type RoutingEventHubPropertiesOutput struct{ *pulumi.OutputState }

func (RoutingEventHubPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEventHubProperties)(nil)).Elem()
}

func (o RoutingEventHubPropertiesOutput) ToRoutingEventHubPropertiesOutput() RoutingEventHubPropertiesOutput {
	return o
}

func (o RoutingEventHubPropertiesOutput) ToRoutingEventHubPropertiesOutputWithContext(ctx context.Context) RoutingEventHubPropertiesOutput {
	return o
}

func (o RoutingEventHubPropertiesOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubProperties) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubProperties) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubProperties) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesOutput) EntityPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubProperties) *string { return v.EntityPath }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingEventHubProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingEventHubPropertiesOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubProperties) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingEventHubPropertiesArrayOutput struct{ *pulumi.OutputState }

func (RoutingEventHubPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingEventHubProperties)(nil)).Elem()
}

func (o RoutingEventHubPropertiesArrayOutput) ToRoutingEventHubPropertiesArrayOutput() RoutingEventHubPropertiesArrayOutput {
	return o
}

func (o RoutingEventHubPropertiesArrayOutput) ToRoutingEventHubPropertiesArrayOutputWithContext(ctx context.Context) RoutingEventHubPropertiesArrayOutput {
	return o
}

func (o RoutingEventHubPropertiesArrayOutput) Index(i pulumi.IntInput) RoutingEventHubPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingEventHubProperties {
		return vs[0].([]RoutingEventHubProperties)[vs[1].(int)]
	}).(RoutingEventHubPropertiesOutput)
}

type RoutingEventHubPropertiesResponse struct {
	AuthenticationType *string `pulumi:"authenticationType"`
	ConnectionString   *string `pulumi:"connectionString"`
	EndpointUri        *string `pulumi:"endpointUri"`
	EntityPath         *string `pulumi:"entityPath"`
	Id                 *string `pulumi:"id"`
	Name               string  `pulumi:"name"`
	ResourceGroup      *string `pulumi:"resourceGroup"`
	SubscriptionId     *string `pulumi:"subscriptionId"`
}

type RoutingEventHubPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RoutingEventHubPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEventHubPropertiesResponse)(nil)).Elem()
}

func (o RoutingEventHubPropertiesResponseOutput) ToRoutingEventHubPropertiesResponseOutput() RoutingEventHubPropertiesResponseOutput {
	return o
}

func (o RoutingEventHubPropertiesResponseOutput) ToRoutingEventHubPropertiesResponseOutputWithContext(ctx context.Context) RoutingEventHubPropertiesResponseOutput {
	return o
}

func (o RoutingEventHubPropertiesResponseOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubPropertiesResponse) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubPropertiesResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesResponseOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubPropertiesResponse) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesResponseOutput) EntityPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubPropertiesResponse) *string { return v.EntityPath }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingEventHubPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingEventHubPropertiesResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubPropertiesResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubPropertiesResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingEventHubPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutingEventHubPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingEventHubPropertiesResponse)(nil)).Elem()
}

func (o RoutingEventHubPropertiesResponseArrayOutput) ToRoutingEventHubPropertiesResponseArrayOutput() RoutingEventHubPropertiesResponseArrayOutput {
	return o
}

func (o RoutingEventHubPropertiesResponseArrayOutput) ToRoutingEventHubPropertiesResponseArrayOutputWithContext(ctx context.Context) RoutingEventHubPropertiesResponseArrayOutput {
	return o
}

func (o RoutingEventHubPropertiesResponseArrayOutput) Index(i pulumi.IntInput) RoutingEventHubPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingEventHubPropertiesResponse {
		return vs[0].([]RoutingEventHubPropertiesResponse)[vs[1].(int)]
	}).(RoutingEventHubPropertiesResponseOutput)
}

type RoutingProperties struct {
	Endpoints     *RoutingEndpoints        `pulumi:"endpoints"`
	Enrichments   []EnrichmentProperties   `pulumi:"enrichments"`
	FallbackRoute *FallbackRouteProperties `pulumi:"fallbackRoute"`
	Routes        []RouteProperties        `pulumi:"routes"`
}





type RoutingPropertiesInput interface {
	pulumi.Input

	ToRoutingPropertiesOutput() RoutingPropertiesOutput
	ToRoutingPropertiesOutputWithContext(context.Context) RoutingPropertiesOutput
}

type RoutingPropertiesArgs struct {
	Endpoints     RoutingEndpointsPtrInput        `pulumi:"endpoints"`
	Enrichments   EnrichmentPropertiesArrayInput  `pulumi:"enrichments"`
	FallbackRoute FallbackRoutePropertiesPtrInput `pulumi:"fallbackRoute"`
	Routes        RoutePropertiesArrayInput       `pulumi:"routes"`
}

func (RoutingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingProperties)(nil)).Elem()
}

func (i RoutingPropertiesArgs) ToRoutingPropertiesOutput() RoutingPropertiesOutput {
	return i.ToRoutingPropertiesOutputWithContext(context.Background())
}

func (i RoutingPropertiesArgs) ToRoutingPropertiesOutputWithContext(ctx context.Context) RoutingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPropertiesOutput)
}

func (i RoutingPropertiesArgs) ToRoutingPropertiesPtrOutput() RoutingPropertiesPtrOutput {
	return i.ToRoutingPropertiesPtrOutputWithContext(context.Background())
}

func (i RoutingPropertiesArgs) ToRoutingPropertiesPtrOutputWithContext(ctx context.Context) RoutingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPropertiesOutput).ToRoutingPropertiesPtrOutputWithContext(ctx)
}









type RoutingPropertiesPtrInput interface {
	pulumi.Input

	ToRoutingPropertiesPtrOutput() RoutingPropertiesPtrOutput
	ToRoutingPropertiesPtrOutputWithContext(context.Context) RoutingPropertiesPtrOutput
}

type routingPropertiesPtrType RoutingPropertiesArgs

func RoutingPropertiesPtr(v *RoutingPropertiesArgs) RoutingPropertiesPtrInput {
	return (*routingPropertiesPtrType)(v)
}

func (*routingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingProperties)(nil)).Elem()
}

func (i *routingPropertiesPtrType) ToRoutingPropertiesPtrOutput() RoutingPropertiesPtrOutput {
	return i.ToRoutingPropertiesPtrOutputWithContext(context.Background())
}

func (i *routingPropertiesPtrType) ToRoutingPropertiesPtrOutputWithContext(ctx context.Context) RoutingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPropertiesPtrOutput)
}

type RoutingPropertiesOutput struct{ *pulumi.OutputState }

func (RoutingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingProperties)(nil)).Elem()
}

func (o RoutingPropertiesOutput) ToRoutingPropertiesOutput() RoutingPropertiesOutput {
	return o
}

func (o RoutingPropertiesOutput) ToRoutingPropertiesOutputWithContext(ctx context.Context) RoutingPropertiesOutput {
	return o
}

func (o RoutingPropertiesOutput) ToRoutingPropertiesPtrOutput() RoutingPropertiesPtrOutput {
	return o.ToRoutingPropertiesPtrOutputWithContext(context.Background())
}

func (o RoutingPropertiesOutput) ToRoutingPropertiesPtrOutputWithContext(ctx context.Context) RoutingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoutingProperties) *RoutingProperties {
		return &v
	}).(RoutingPropertiesPtrOutput)
}

func (o RoutingPropertiesOutput) Endpoints() RoutingEndpointsPtrOutput {
	return o.ApplyT(func(v RoutingProperties) *RoutingEndpoints { return v.Endpoints }).(RoutingEndpointsPtrOutput)
}

func (o RoutingPropertiesOutput) Enrichments() EnrichmentPropertiesArrayOutput {
	return o.ApplyT(func(v RoutingProperties) []EnrichmentProperties { return v.Enrichments }).(EnrichmentPropertiesArrayOutput)
}

func (o RoutingPropertiesOutput) FallbackRoute() FallbackRoutePropertiesPtrOutput {
	return o.ApplyT(func(v RoutingProperties) *FallbackRouteProperties { return v.FallbackRoute }).(FallbackRoutePropertiesPtrOutput)
}

func (o RoutingPropertiesOutput) Routes() RoutePropertiesArrayOutput {
	return o.ApplyT(func(v RoutingProperties) []RouteProperties { return v.Routes }).(RoutePropertiesArrayOutput)
}

type RoutingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (RoutingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingProperties)(nil)).Elem()
}

func (o RoutingPropertiesPtrOutput) ToRoutingPropertiesPtrOutput() RoutingPropertiesPtrOutput {
	return o
}

func (o RoutingPropertiesPtrOutput) ToRoutingPropertiesPtrOutputWithContext(ctx context.Context) RoutingPropertiesPtrOutput {
	return o
}

func (o RoutingPropertiesPtrOutput) Elem() RoutingPropertiesOutput {
	return o.ApplyT(func(v *RoutingProperties) RoutingProperties {
		if v != nil {
			return *v
		}
		var ret RoutingProperties
		return ret
	}).(RoutingPropertiesOutput)
}

func (o RoutingPropertiesPtrOutput) Endpoints() RoutingEndpointsPtrOutput {
	return o.ApplyT(func(v *RoutingProperties) *RoutingEndpoints {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(RoutingEndpointsPtrOutput)
}

func (o RoutingPropertiesPtrOutput) Enrichments() EnrichmentPropertiesArrayOutput {
	return o.ApplyT(func(v *RoutingProperties) []EnrichmentProperties {
		if v == nil {
			return nil
		}
		return v.Enrichments
	}).(EnrichmentPropertiesArrayOutput)
}

func (o RoutingPropertiesPtrOutput) FallbackRoute() FallbackRoutePropertiesPtrOutput {
	return o.ApplyT(func(v *RoutingProperties) *FallbackRouteProperties {
		if v == nil {
			return nil
		}
		return v.FallbackRoute
	}).(FallbackRoutePropertiesPtrOutput)
}

func (o RoutingPropertiesPtrOutput) Routes() RoutePropertiesArrayOutput {
	return o.ApplyT(func(v *RoutingProperties) []RouteProperties {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(RoutePropertiesArrayOutput)
}

type RoutingPropertiesResponse struct {
	Endpoints     *RoutingEndpointsResponse        `pulumi:"endpoints"`
	Enrichments   []EnrichmentPropertiesResponse   `pulumi:"enrichments"`
	FallbackRoute *FallbackRoutePropertiesResponse `pulumi:"fallbackRoute"`
	Routes        []RoutePropertiesResponse        `pulumi:"routes"`
}

type RoutingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RoutingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingPropertiesResponse)(nil)).Elem()
}

func (o RoutingPropertiesResponseOutput) ToRoutingPropertiesResponseOutput() RoutingPropertiesResponseOutput {
	return o
}

func (o RoutingPropertiesResponseOutput) ToRoutingPropertiesResponseOutputWithContext(ctx context.Context) RoutingPropertiesResponseOutput {
	return o
}

func (o RoutingPropertiesResponseOutput) Endpoints() RoutingEndpointsResponsePtrOutput {
	return o.ApplyT(func(v RoutingPropertiesResponse) *RoutingEndpointsResponse { return v.Endpoints }).(RoutingEndpointsResponsePtrOutput)
}

func (o RoutingPropertiesResponseOutput) Enrichments() EnrichmentPropertiesResponseArrayOutput {
	return o.ApplyT(func(v RoutingPropertiesResponse) []EnrichmentPropertiesResponse { return v.Enrichments }).(EnrichmentPropertiesResponseArrayOutput)
}

func (o RoutingPropertiesResponseOutput) FallbackRoute() FallbackRoutePropertiesResponsePtrOutput {
	return o.ApplyT(func(v RoutingPropertiesResponse) *FallbackRoutePropertiesResponse { return v.FallbackRoute }).(FallbackRoutePropertiesResponsePtrOutput)
}

func (o RoutingPropertiesResponseOutput) Routes() RoutePropertiesResponseArrayOutput {
	return o.ApplyT(func(v RoutingPropertiesResponse) []RoutePropertiesResponse { return v.Routes }).(RoutePropertiesResponseArrayOutput)
}

type RoutingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (RoutingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingPropertiesResponse)(nil)).Elem()
}

func (o RoutingPropertiesResponsePtrOutput) ToRoutingPropertiesResponsePtrOutput() RoutingPropertiesResponsePtrOutput {
	return o
}

func (o RoutingPropertiesResponsePtrOutput) ToRoutingPropertiesResponsePtrOutputWithContext(ctx context.Context) RoutingPropertiesResponsePtrOutput {
	return o
}

func (o RoutingPropertiesResponsePtrOutput) Elem() RoutingPropertiesResponseOutput {
	return o.ApplyT(func(v *RoutingPropertiesResponse) RoutingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret RoutingPropertiesResponse
		return ret
	}).(RoutingPropertiesResponseOutput)
}

func (o RoutingPropertiesResponsePtrOutput) Endpoints() RoutingEndpointsResponsePtrOutput {
	return o.ApplyT(func(v *RoutingPropertiesResponse) *RoutingEndpointsResponse {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(RoutingEndpointsResponsePtrOutput)
}

func (o RoutingPropertiesResponsePtrOutput) Enrichments() EnrichmentPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *RoutingPropertiesResponse) []EnrichmentPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Enrichments
	}).(EnrichmentPropertiesResponseArrayOutput)
}

func (o RoutingPropertiesResponsePtrOutput) FallbackRoute() FallbackRoutePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *RoutingPropertiesResponse) *FallbackRoutePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.FallbackRoute
	}).(FallbackRoutePropertiesResponsePtrOutput)
}

func (o RoutingPropertiesResponsePtrOutput) Routes() RoutePropertiesResponseArrayOutput {
	return o.ApplyT(func(v *RoutingPropertiesResponse) []RoutePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(RoutePropertiesResponseArrayOutput)
}

type RoutingServiceBusQueueEndpointProperties struct {
	AuthenticationType *string `pulumi:"authenticationType"`
	ConnectionString   *string `pulumi:"connectionString"`
	EndpointUri        *string `pulumi:"endpointUri"`
	EntityPath         *string `pulumi:"entityPath"`
	Id                 *string `pulumi:"id"`
	Name               string  `pulumi:"name"`
	ResourceGroup      *string `pulumi:"resourceGroup"`
	SubscriptionId     *string `pulumi:"subscriptionId"`
}





type RoutingServiceBusQueueEndpointPropertiesInput interface {
	pulumi.Input

	ToRoutingServiceBusQueueEndpointPropertiesOutput() RoutingServiceBusQueueEndpointPropertiesOutput
	ToRoutingServiceBusQueueEndpointPropertiesOutputWithContext(context.Context) RoutingServiceBusQueueEndpointPropertiesOutput
}

type RoutingServiceBusQueueEndpointPropertiesArgs struct {
	AuthenticationType pulumi.StringPtrInput `pulumi:"authenticationType"`
	ConnectionString   pulumi.StringPtrInput `pulumi:"connectionString"`
	EndpointUri        pulumi.StringPtrInput `pulumi:"endpointUri"`
	EntityPath         pulumi.StringPtrInput `pulumi:"entityPath"`
	Id                 pulumi.StringPtrInput `pulumi:"id"`
	Name               pulumi.StringInput    `pulumi:"name"`
	ResourceGroup      pulumi.StringPtrInput `pulumi:"resourceGroup"`
	SubscriptionId     pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (RoutingServiceBusQueueEndpointPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusQueueEndpointProperties)(nil)).Elem()
}

func (i RoutingServiceBusQueueEndpointPropertiesArgs) ToRoutingServiceBusQueueEndpointPropertiesOutput() RoutingServiceBusQueueEndpointPropertiesOutput {
	return i.ToRoutingServiceBusQueueEndpointPropertiesOutputWithContext(context.Background())
}

func (i RoutingServiceBusQueueEndpointPropertiesArgs) ToRoutingServiceBusQueueEndpointPropertiesOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingServiceBusQueueEndpointPropertiesOutput)
}





type RoutingServiceBusQueueEndpointPropertiesArrayInput interface {
	pulumi.Input

	ToRoutingServiceBusQueueEndpointPropertiesArrayOutput() RoutingServiceBusQueueEndpointPropertiesArrayOutput
	ToRoutingServiceBusQueueEndpointPropertiesArrayOutputWithContext(context.Context) RoutingServiceBusQueueEndpointPropertiesArrayOutput
}

type RoutingServiceBusQueueEndpointPropertiesArray []RoutingServiceBusQueueEndpointPropertiesInput

func (RoutingServiceBusQueueEndpointPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusQueueEndpointProperties)(nil)).Elem()
}

func (i RoutingServiceBusQueueEndpointPropertiesArray) ToRoutingServiceBusQueueEndpointPropertiesArrayOutput() RoutingServiceBusQueueEndpointPropertiesArrayOutput {
	return i.ToRoutingServiceBusQueueEndpointPropertiesArrayOutputWithContext(context.Background())
}

func (i RoutingServiceBusQueueEndpointPropertiesArray) ToRoutingServiceBusQueueEndpointPropertiesArrayOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingServiceBusQueueEndpointPropertiesArrayOutput)
}

type RoutingServiceBusQueueEndpointPropertiesOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusQueueEndpointPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusQueueEndpointProperties)(nil)).Elem()
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) ToRoutingServiceBusQueueEndpointPropertiesOutput() RoutingServiceBusQueueEndpointPropertiesOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) ToRoutingServiceBusQueueEndpointPropertiesOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointProperties) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointProperties) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointProperties) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) EntityPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointProperties) *string { return v.EntityPath }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointProperties) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingServiceBusQueueEndpointPropertiesArrayOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusQueueEndpointPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusQueueEndpointProperties)(nil)).Elem()
}

func (o RoutingServiceBusQueueEndpointPropertiesArrayOutput) ToRoutingServiceBusQueueEndpointPropertiesArrayOutput() RoutingServiceBusQueueEndpointPropertiesArrayOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesArrayOutput) ToRoutingServiceBusQueueEndpointPropertiesArrayOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesArrayOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesArrayOutput) Index(i pulumi.IntInput) RoutingServiceBusQueueEndpointPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingServiceBusQueueEndpointProperties {
		return vs[0].([]RoutingServiceBusQueueEndpointProperties)[vs[1].(int)]
	}).(RoutingServiceBusQueueEndpointPropertiesOutput)
}

type RoutingServiceBusQueueEndpointPropertiesResponse struct {
	AuthenticationType *string `pulumi:"authenticationType"`
	ConnectionString   *string `pulumi:"connectionString"`
	EndpointUri        *string `pulumi:"endpointUri"`
	EntityPath         *string `pulumi:"entityPath"`
	Id                 *string `pulumi:"id"`
	Name               string  `pulumi:"name"`
	ResourceGroup      *string `pulumi:"resourceGroup"`
	SubscriptionId     *string `pulumi:"subscriptionId"`
}

type RoutingServiceBusQueueEndpointPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusQueueEndpointPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusQueueEndpointPropertiesResponse)(nil)).Elem()
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) ToRoutingServiceBusQueueEndpointPropertiesResponseOutput() RoutingServiceBusQueueEndpointPropertiesResponseOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) ToRoutingServiceBusQueueEndpointPropertiesResponseOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesResponseOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointPropertiesResponse) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointPropertiesResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointPropertiesResponse) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) EntityPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointPropertiesResponse) *string { return v.EntityPath }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointPropertiesResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointPropertiesResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusQueueEndpointPropertiesResponse)(nil)).Elem()
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput) ToRoutingServiceBusQueueEndpointPropertiesResponseArrayOutput() RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput) ToRoutingServiceBusQueueEndpointPropertiesResponseArrayOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput) Index(i pulumi.IntInput) RoutingServiceBusQueueEndpointPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingServiceBusQueueEndpointPropertiesResponse {
		return vs[0].([]RoutingServiceBusQueueEndpointPropertiesResponse)[vs[1].(int)]
	}).(RoutingServiceBusQueueEndpointPropertiesResponseOutput)
}

type RoutingServiceBusTopicEndpointProperties struct {
	AuthenticationType *string `pulumi:"authenticationType"`
	ConnectionString   *string `pulumi:"connectionString"`
	EndpointUri        *string `pulumi:"endpointUri"`
	EntityPath         *string `pulumi:"entityPath"`
	Id                 *string `pulumi:"id"`
	Name               string  `pulumi:"name"`
	ResourceGroup      *string `pulumi:"resourceGroup"`
	SubscriptionId     *string `pulumi:"subscriptionId"`
}





type RoutingServiceBusTopicEndpointPropertiesInput interface {
	pulumi.Input

	ToRoutingServiceBusTopicEndpointPropertiesOutput() RoutingServiceBusTopicEndpointPropertiesOutput
	ToRoutingServiceBusTopicEndpointPropertiesOutputWithContext(context.Context) RoutingServiceBusTopicEndpointPropertiesOutput
}

type RoutingServiceBusTopicEndpointPropertiesArgs struct {
	AuthenticationType pulumi.StringPtrInput `pulumi:"authenticationType"`
	ConnectionString   pulumi.StringPtrInput `pulumi:"connectionString"`
	EndpointUri        pulumi.StringPtrInput `pulumi:"endpointUri"`
	EntityPath         pulumi.StringPtrInput `pulumi:"entityPath"`
	Id                 pulumi.StringPtrInput `pulumi:"id"`
	Name               pulumi.StringInput    `pulumi:"name"`
	ResourceGroup      pulumi.StringPtrInput `pulumi:"resourceGroup"`
	SubscriptionId     pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (RoutingServiceBusTopicEndpointPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusTopicEndpointProperties)(nil)).Elem()
}

func (i RoutingServiceBusTopicEndpointPropertiesArgs) ToRoutingServiceBusTopicEndpointPropertiesOutput() RoutingServiceBusTopicEndpointPropertiesOutput {
	return i.ToRoutingServiceBusTopicEndpointPropertiesOutputWithContext(context.Background())
}

func (i RoutingServiceBusTopicEndpointPropertiesArgs) ToRoutingServiceBusTopicEndpointPropertiesOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingServiceBusTopicEndpointPropertiesOutput)
}





type RoutingServiceBusTopicEndpointPropertiesArrayInput interface {
	pulumi.Input

	ToRoutingServiceBusTopicEndpointPropertiesArrayOutput() RoutingServiceBusTopicEndpointPropertiesArrayOutput
	ToRoutingServiceBusTopicEndpointPropertiesArrayOutputWithContext(context.Context) RoutingServiceBusTopicEndpointPropertiesArrayOutput
}

type RoutingServiceBusTopicEndpointPropertiesArray []RoutingServiceBusTopicEndpointPropertiesInput

func (RoutingServiceBusTopicEndpointPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusTopicEndpointProperties)(nil)).Elem()
}

func (i RoutingServiceBusTopicEndpointPropertiesArray) ToRoutingServiceBusTopicEndpointPropertiesArrayOutput() RoutingServiceBusTopicEndpointPropertiesArrayOutput {
	return i.ToRoutingServiceBusTopicEndpointPropertiesArrayOutputWithContext(context.Background())
}

func (i RoutingServiceBusTopicEndpointPropertiesArray) ToRoutingServiceBusTopicEndpointPropertiesArrayOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingServiceBusTopicEndpointPropertiesArrayOutput)
}

type RoutingServiceBusTopicEndpointPropertiesOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusTopicEndpointPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusTopicEndpointProperties)(nil)).Elem()
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) ToRoutingServiceBusTopicEndpointPropertiesOutput() RoutingServiceBusTopicEndpointPropertiesOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) ToRoutingServiceBusTopicEndpointPropertiesOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointProperties) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointProperties) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointProperties) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) EntityPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointProperties) *string { return v.EntityPath }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointProperties) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingServiceBusTopicEndpointPropertiesArrayOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusTopicEndpointPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusTopicEndpointProperties)(nil)).Elem()
}

func (o RoutingServiceBusTopicEndpointPropertiesArrayOutput) ToRoutingServiceBusTopicEndpointPropertiesArrayOutput() RoutingServiceBusTopicEndpointPropertiesArrayOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesArrayOutput) ToRoutingServiceBusTopicEndpointPropertiesArrayOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesArrayOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesArrayOutput) Index(i pulumi.IntInput) RoutingServiceBusTopicEndpointPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingServiceBusTopicEndpointProperties {
		return vs[0].([]RoutingServiceBusTopicEndpointProperties)[vs[1].(int)]
	}).(RoutingServiceBusTopicEndpointPropertiesOutput)
}

type RoutingServiceBusTopicEndpointPropertiesResponse struct {
	AuthenticationType *string `pulumi:"authenticationType"`
	ConnectionString   *string `pulumi:"connectionString"`
	EndpointUri        *string `pulumi:"endpointUri"`
	EntityPath         *string `pulumi:"entityPath"`
	Id                 *string `pulumi:"id"`
	Name               string  `pulumi:"name"`
	ResourceGroup      *string `pulumi:"resourceGroup"`
	SubscriptionId     *string `pulumi:"subscriptionId"`
}

type RoutingServiceBusTopicEndpointPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusTopicEndpointPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusTopicEndpointPropertiesResponse)(nil)).Elem()
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) ToRoutingServiceBusTopicEndpointPropertiesResponseOutput() RoutingServiceBusTopicEndpointPropertiesResponseOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) ToRoutingServiceBusTopicEndpointPropertiesResponseOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesResponseOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointPropertiesResponse) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointPropertiesResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointPropertiesResponse) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) EntityPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointPropertiesResponse) *string { return v.EntityPath }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointPropertiesResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointPropertiesResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusTopicEndpointPropertiesResponse)(nil)).Elem()
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput) ToRoutingServiceBusTopicEndpointPropertiesResponseArrayOutput() RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput) ToRoutingServiceBusTopicEndpointPropertiesResponseArrayOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput) Index(i pulumi.IntInput) RoutingServiceBusTopicEndpointPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingServiceBusTopicEndpointPropertiesResponse {
		return vs[0].([]RoutingServiceBusTopicEndpointPropertiesResponse)[vs[1].(int)]
	}).(RoutingServiceBusTopicEndpointPropertiesResponseOutput)
}

type RoutingStorageContainerProperties struct {
	AuthenticationType      *string `pulumi:"authenticationType"`
	BatchFrequencyInSeconds *int    `pulumi:"batchFrequencyInSeconds"`
	ConnectionString        *string `pulumi:"connectionString"`
	ContainerName           string  `pulumi:"containerName"`
	Encoding                *string `pulumi:"encoding"`
	EndpointUri             *string `pulumi:"endpointUri"`
	FileNameFormat          *string `pulumi:"fileNameFormat"`
	Id                      *string `pulumi:"id"`
	MaxChunkSizeInBytes     *int    `pulumi:"maxChunkSizeInBytes"`
	Name                    string  `pulumi:"name"`
	ResourceGroup           *string `pulumi:"resourceGroup"`
	SubscriptionId          *string `pulumi:"subscriptionId"`
}





type RoutingStorageContainerPropertiesInput interface {
	pulumi.Input

	ToRoutingStorageContainerPropertiesOutput() RoutingStorageContainerPropertiesOutput
	ToRoutingStorageContainerPropertiesOutputWithContext(context.Context) RoutingStorageContainerPropertiesOutput
}

type RoutingStorageContainerPropertiesArgs struct {
	AuthenticationType      pulumi.StringPtrInput `pulumi:"authenticationType"`
	BatchFrequencyInSeconds pulumi.IntPtrInput    `pulumi:"batchFrequencyInSeconds"`
	ConnectionString        pulumi.StringPtrInput `pulumi:"connectionString"`
	ContainerName           pulumi.StringInput    `pulumi:"containerName"`
	Encoding                pulumi.StringPtrInput `pulumi:"encoding"`
	EndpointUri             pulumi.StringPtrInput `pulumi:"endpointUri"`
	FileNameFormat          pulumi.StringPtrInput `pulumi:"fileNameFormat"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	MaxChunkSizeInBytes     pulumi.IntPtrInput    `pulumi:"maxChunkSizeInBytes"`
	Name                    pulumi.StringInput    `pulumi:"name"`
	ResourceGroup           pulumi.StringPtrInput `pulumi:"resourceGroup"`
	SubscriptionId          pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (RoutingStorageContainerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingStorageContainerProperties)(nil)).Elem()
}

func (i RoutingStorageContainerPropertiesArgs) ToRoutingStorageContainerPropertiesOutput() RoutingStorageContainerPropertiesOutput {
	return i.ToRoutingStorageContainerPropertiesOutputWithContext(context.Background())
}

func (i RoutingStorageContainerPropertiesArgs) ToRoutingStorageContainerPropertiesOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingStorageContainerPropertiesOutput)
}





type RoutingStorageContainerPropertiesArrayInput interface {
	pulumi.Input

	ToRoutingStorageContainerPropertiesArrayOutput() RoutingStorageContainerPropertiesArrayOutput
	ToRoutingStorageContainerPropertiesArrayOutputWithContext(context.Context) RoutingStorageContainerPropertiesArrayOutput
}

type RoutingStorageContainerPropertiesArray []RoutingStorageContainerPropertiesInput

func (RoutingStorageContainerPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingStorageContainerProperties)(nil)).Elem()
}

func (i RoutingStorageContainerPropertiesArray) ToRoutingStorageContainerPropertiesArrayOutput() RoutingStorageContainerPropertiesArrayOutput {
	return i.ToRoutingStorageContainerPropertiesArrayOutputWithContext(context.Background())
}

func (i RoutingStorageContainerPropertiesArray) ToRoutingStorageContainerPropertiesArrayOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingStorageContainerPropertiesArrayOutput)
}

type RoutingStorageContainerPropertiesOutput struct{ *pulumi.OutputState }

func (RoutingStorageContainerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingStorageContainerProperties)(nil)).Elem()
}

func (o RoutingStorageContainerPropertiesOutput) ToRoutingStorageContainerPropertiesOutput() RoutingStorageContainerPropertiesOutput {
	return o
}

func (o RoutingStorageContainerPropertiesOutput) ToRoutingStorageContainerPropertiesOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesOutput {
	return o
}

func (o RoutingStorageContainerPropertiesOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) BatchFrequencyInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *int { return v.BatchFrequencyInSeconds }).(pulumi.IntPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o RoutingStorageContainerPropertiesOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *string { return v.Encoding }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) FileNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *string { return v.FileNameFormat }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) MaxChunkSizeInBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *int { return v.MaxChunkSizeInBytes }).(pulumi.IntPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingStorageContainerPropertiesOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingStorageContainerPropertiesArrayOutput struct{ *pulumi.OutputState }

func (RoutingStorageContainerPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingStorageContainerProperties)(nil)).Elem()
}

func (o RoutingStorageContainerPropertiesArrayOutput) ToRoutingStorageContainerPropertiesArrayOutput() RoutingStorageContainerPropertiesArrayOutput {
	return o
}

func (o RoutingStorageContainerPropertiesArrayOutput) ToRoutingStorageContainerPropertiesArrayOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesArrayOutput {
	return o
}

func (o RoutingStorageContainerPropertiesArrayOutput) Index(i pulumi.IntInput) RoutingStorageContainerPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingStorageContainerProperties {
		return vs[0].([]RoutingStorageContainerProperties)[vs[1].(int)]
	}).(RoutingStorageContainerPropertiesOutput)
}

type RoutingStorageContainerPropertiesResponse struct {
	AuthenticationType      *string `pulumi:"authenticationType"`
	BatchFrequencyInSeconds *int    `pulumi:"batchFrequencyInSeconds"`
	ConnectionString        *string `pulumi:"connectionString"`
	ContainerName           string  `pulumi:"containerName"`
	Encoding                *string `pulumi:"encoding"`
	EndpointUri             *string `pulumi:"endpointUri"`
	FileNameFormat          *string `pulumi:"fileNameFormat"`
	Id                      *string `pulumi:"id"`
	MaxChunkSizeInBytes     *int    `pulumi:"maxChunkSizeInBytes"`
	Name                    string  `pulumi:"name"`
	ResourceGroup           *string `pulumi:"resourceGroup"`
	SubscriptionId          *string `pulumi:"subscriptionId"`
}

type RoutingStorageContainerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RoutingStorageContainerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingStorageContainerPropertiesResponse)(nil)).Elem()
}

func (o RoutingStorageContainerPropertiesResponseOutput) ToRoutingStorageContainerPropertiesResponseOutput() RoutingStorageContainerPropertiesResponseOutput {
	return o
}

func (o RoutingStorageContainerPropertiesResponseOutput) ToRoutingStorageContainerPropertiesResponseOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesResponseOutput {
	return o
}

func (o RoutingStorageContainerPropertiesResponseOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) BatchFrequencyInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *int { return v.BatchFrequencyInSeconds }).(pulumi.IntPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *string { return v.Encoding }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) FileNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *string { return v.FileNameFormat }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) MaxChunkSizeInBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *int { return v.MaxChunkSizeInBytes }).(pulumi.IntPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingStorageContainerPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutingStorageContainerPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingStorageContainerPropertiesResponse)(nil)).Elem()
}

func (o RoutingStorageContainerPropertiesResponseArrayOutput) ToRoutingStorageContainerPropertiesResponseArrayOutput() RoutingStorageContainerPropertiesResponseArrayOutput {
	return o
}

func (o RoutingStorageContainerPropertiesResponseArrayOutput) ToRoutingStorageContainerPropertiesResponseArrayOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesResponseArrayOutput {
	return o
}

func (o RoutingStorageContainerPropertiesResponseArrayOutput) Index(i pulumi.IntInput) RoutingStorageContainerPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingStorageContainerPropertiesResponse {
		return vs[0].([]RoutingStorageContainerPropertiesResponse)[vs[1].(int)]
	}).(RoutingStorageContainerPropertiesResponseOutput)
}

type SharedAccessSignatureAuthorizationRule struct {
	KeyName      string       `pulumi:"keyName"`
	PrimaryKey   *string      `pulumi:"primaryKey"`
	Rights       AccessRights `pulumi:"rights"`
	SecondaryKey *string      `pulumi:"secondaryKey"`
}





type SharedAccessSignatureAuthorizationRuleInput interface {
	pulumi.Input

	ToSharedAccessSignatureAuthorizationRuleOutput() SharedAccessSignatureAuthorizationRuleOutput
	ToSharedAccessSignatureAuthorizationRuleOutputWithContext(context.Context) SharedAccessSignatureAuthorizationRuleOutput
}

type SharedAccessSignatureAuthorizationRuleArgs struct {
	KeyName      pulumi.StringInput    `pulumi:"keyName"`
	PrimaryKey   pulumi.StringPtrInput `pulumi:"primaryKey"`
	Rights       AccessRightsInput     `pulumi:"rights"`
	SecondaryKey pulumi.StringPtrInput `pulumi:"secondaryKey"`
}

func (SharedAccessSignatureAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRule)(nil)).Elem()
}

func (i SharedAccessSignatureAuthorizationRuleArgs) ToSharedAccessSignatureAuthorizationRuleOutput() SharedAccessSignatureAuthorizationRuleOutput {
	return i.ToSharedAccessSignatureAuthorizationRuleOutputWithContext(context.Background())
}

func (i SharedAccessSignatureAuthorizationRuleArgs) ToSharedAccessSignatureAuthorizationRuleOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessSignatureAuthorizationRuleOutput)
}





type SharedAccessSignatureAuthorizationRuleArrayInput interface {
	pulumi.Input

	ToSharedAccessSignatureAuthorizationRuleArrayOutput() SharedAccessSignatureAuthorizationRuleArrayOutput
	ToSharedAccessSignatureAuthorizationRuleArrayOutputWithContext(context.Context) SharedAccessSignatureAuthorizationRuleArrayOutput
}

type SharedAccessSignatureAuthorizationRuleArray []SharedAccessSignatureAuthorizationRuleInput

func (SharedAccessSignatureAuthorizationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRule)(nil)).Elem()
}

func (i SharedAccessSignatureAuthorizationRuleArray) ToSharedAccessSignatureAuthorizationRuleArrayOutput() SharedAccessSignatureAuthorizationRuleArrayOutput {
	return i.ToSharedAccessSignatureAuthorizationRuleArrayOutputWithContext(context.Background())
}

func (i SharedAccessSignatureAuthorizationRuleArray) ToSharedAccessSignatureAuthorizationRuleArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessSignatureAuthorizationRuleArrayOutput)
}

type SharedAccessSignatureAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRule)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleOutput) ToSharedAccessSignatureAuthorizationRuleOutput() SharedAccessSignatureAuthorizationRuleOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleOutput) ToSharedAccessSignatureAuthorizationRuleOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRule) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRule) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o SharedAccessSignatureAuthorizationRuleOutput) Rights() AccessRightsOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRule) AccessRights { return v.Rights }).(AccessRightsOutput)
}

func (o SharedAccessSignatureAuthorizationRuleOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRule) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

type SharedAccessSignatureAuthorizationRuleArrayOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRule)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleArrayOutput) ToSharedAccessSignatureAuthorizationRuleArrayOutput() SharedAccessSignatureAuthorizationRuleArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleArrayOutput) ToSharedAccessSignatureAuthorizationRuleArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleArrayOutput) Index(i pulumi.IntInput) SharedAccessSignatureAuthorizationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedAccessSignatureAuthorizationRule {
		return vs[0].([]SharedAccessSignatureAuthorizationRule)[vs[1].(int)]
	}).(SharedAccessSignatureAuthorizationRuleOutput)
}

type SharedAccessSignatureAuthorizationRuleResponse struct {
	KeyName      string  `pulumi:"keyName"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	Rights       string  `pulumi:"rights"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}

type SharedAccessSignatureAuthorizationRuleResponseOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleResponse)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleResponseOutput) ToSharedAccessSignatureAuthorizationRuleResponseOutput() SharedAccessSignatureAuthorizationRuleResponseOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleResponseOutput) ToSharedAccessSignatureAuthorizationRuleResponseOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleResponseOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleResponseOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleResponse) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleResponseOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleResponse) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o SharedAccessSignatureAuthorizationRuleResponseOutput) Rights() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleResponse) string { return v.Rights }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleResponseOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleResponse) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

type SharedAccessSignatureAuthorizationRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRuleResponse)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleResponseArrayOutput) ToSharedAccessSignatureAuthorizationRuleResponseArrayOutput() SharedAccessSignatureAuthorizationRuleResponseArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleResponseArrayOutput) ToSharedAccessSignatureAuthorizationRuleResponseArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleResponseArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleResponseArrayOutput) Index(i pulumi.IntInput) SharedAccessSignatureAuthorizationRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedAccessSignatureAuthorizationRuleResponse {
		return vs[0].([]SharedAccessSignatureAuthorizationRuleResponse)[vs[1].(int)]
	}).(SharedAccessSignatureAuthorizationRuleResponseOutput)
}

type StorageEndpointProperties struct {
	AuthenticationType *string `pulumi:"authenticationType"`
	ConnectionString   string  `pulumi:"connectionString"`
	ContainerName      string  `pulumi:"containerName"`
	SasTtlAsIso8601    *string `pulumi:"sasTtlAsIso8601"`
}





type StorageEndpointPropertiesInput interface {
	pulumi.Input

	ToStorageEndpointPropertiesOutput() StorageEndpointPropertiesOutput
	ToStorageEndpointPropertiesOutputWithContext(context.Context) StorageEndpointPropertiesOutput
}

type StorageEndpointPropertiesArgs struct {
	AuthenticationType pulumi.StringPtrInput `pulumi:"authenticationType"`
	ConnectionString   pulumi.StringInput    `pulumi:"connectionString"`
	ContainerName      pulumi.StringInput    `pulumi:"containerName"`
	SasTtlAsIso8601    pulumi.StringPtrInput `pulumi:"sasTtlAsIso8601"`
}

func (StorageEndpointPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageEndpointProperties)(nil)).Elem()
}

func (i StorageEndpointPropertiesArgs) ToStorageEndpointPropertiesOutput() StorageEndpointPropertiesOutput {
	return i.ToStorageEndpointPropertiesOutputWithContext(context.Background())
}

func (i StorageEndpointPropertiesArgs) ToStorageEndpointPropertiesOutputWithContext(ctx context.Context) StorageEndpointPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageEndpointPropertiesOutput)
}





type StorageEndpointPropertiesMapInput interface {
	pulumi.Input

	ToStorageEndpointPropertiesMapOutput() StorageEndpointPropertiesMapOutput
	ToStorageEndpointPropertiesMapOutputWithContext(context.Context) StorageEndpointPropertiesMapOutput
}

type StorageEndpointPropertiesMap map[string]StorageEndpointPropertiesInput

func (StorageEndpointPropertiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StorageEndpointProperties)(nil)).Elem()
}

func (i StorageEndpointPropertiesMap) ToStorageEndpointPropertiesMapOutput() StorageEndpointPropertiesMapOutput {
	return i.ToStorageEndpointPropertiesMapOutputWithContext(context.Background())
}

func (i StorageEndpointPropertiesMap) ToStorageEndpointPropertiesMapOutputWithContext(ctx context.Context) StorageEndpointPropertiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageEndpointPropertiesMapOutput)
}

type StorageEndpointPropertiesOutput struct{ *pulumi.OutputState }

func (StorageEndpointPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageEndpointProperties)(nil)).Elem()
}

func (o StorageEndpointPropertiesOutput) ToStorageEndpointPropertiesOutput() StorageEndpointPropertiesOutput {
	return o
}

func (o StorageEndpointPropertiesOutput) ToStorageEndpointPropertiesOutputWithContext(ctx context.Context) StorageEndpointPropertiesOutput {
	return o
}

func (o StorageEndpointPropertiesOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageEndpointProperties) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o StorageEndpointPropertiesOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v StorageEndpointProperties) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o StorageEndpointPropertiesOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v StorageEndpointProperties) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o StorageEndpointPropertiesOutput) SasTtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageEndpointProperties) *string { return v.SasTtlAsIso8601 }).(pulumi.StringPtrOutput)
}

type StorageEndpointPropertiesMapOutput struct{ *pulumi.OutputState }

func (StorageEndpointPropertiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StorageEndpointProperties)(nil)).Elem()
}

func (o StorageEndpointPropertiesMapOutput) ToStorageEndpointPropertiesMapOutput() StorageEndpointPropertiesMapOutput {
	return o
}

func (o StorageEndpointPropertiesMapOutput) ToStorageEndpointPropertiesMapOutputWithContext(ctx context.Context) StorageEndpointPropertiesMapOutput {
	return o
}

func (o StorageEndpointPropertiesMapOutput) MapIndex(k pulumi.StringInput) StorageEndpointPropertiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StorageEndpointProperties {
		return vs[0].(map[string]StorageEndpointProperties)[vs[1].(string)]
	}).(StorageEndpointPropertiesOutput)
}

type StorageEndpointPropertiesResponse struct {
	AuthenticationType *string `pulumi:"authenticationType"`
	ConnectionString   string  `pulumi:"connectionString"`
	ContainerName      string  `pulumi:"containerName"`
	SasTtlAsIso8601    *string `pulumi:"sasTtlAsIso8601"`
}

type StorageEndpointPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StorageEndpointPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageEndpointPropertiesResponse)(nil)).Elem()
}

func (o StorageEndpointPropertiesResponseOutput) ToStorageEndpointPropertiesResponseOutput() StorageEndpointPropertiesResponseOutput {
	return o
}

func (o StorageEndpointPropertiesResponseOutput) ToStorageEndpointPropertiesResponseOutputWithContext(ctx context.Context) StorageEndpointPropertiesResponseOutput {
	return o
}

func (o StorageEndpointPropertiesResponseOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageEndpointPropertiesResponse) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o StorageEndpointPropertiesResponseOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v StorageEndpointPropertiesResponse) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o StorageEndpointPropertiesResponseOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v StorageEndpointPropertiesResponse) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o StorageEndpointPropertiesResponseOutput) SasTtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageEndpointPropertiesResponse) *string { return v.SasTtlAsIso8601 }).(pulumi.StringPtrOutput)
}

type StorageEndpointPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (StorageEndpointPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StorageEndpointPropertiesResponse)(nil)).Elem()
}

func (o StorageEndpointPropertiesResponseMapOutput) ToStorageEndpointPropertiesResponseMapOutput() StorageEndpointPropertiesResponseMapOutput {
	return o
}

func (o StorageEndpointPropertiesResponseMapOutput) ToStorageEndpointPropertiesResponseMapOutputWithContext(ctx context.Context) StorageEndpointPropertiesResponseMapOutput {
	return o
}

func (o StorageEndpointPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) StorageEndpointPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StorageEndpointPropertiesResponse {
		return vs[0].(map[string]StorageEndpointPropertiesResponse)[vs[1].(string)]
	}).(StorageEndpointPropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ArmIdentityOutput{})
	pulumi.RegisterOutputType(ArmIdentityPtrOutput{})
	pulumi.RegisterOutputType(ArmIdentityResponseOutput{})
	pulumi.RegisterOutputType(ArmIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ArmUserIdentityResponseOutput{})
	pulumi.RegisterOutputType(ArmUserIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesPtrOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesResponseOutput{})
	pulumi.RegisterOutputType(CloudToDevicePropertiesOutput{})
	pulumi.RegisterOutputType(CloudToDevicePropertiesPtrOutput{})
	pulumi.RegisterOutputType(CloudToDevicePropertiesResponseOutput{})
	pulumi.RegisterOutputType(CloudToDevicePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesDescriptionOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesDescriptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesDescriptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(EnrichmentPropertiesOutput{})
	pulumi.RegisterOutputType(EnrichmentPropertiesArrayOutput{})
	pulumi.RegisterOutputType(EnrichmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EnrichmentPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(EventHubConsumerGroupNameOutput{})
	pulumi.RegisterOutputType(EventHubConsumerGroupNamePtrOutput{})
	pulumi.RegisterOutputType(EventHubPropertiesOutput{})
	pulumi.RegisterOutputType(EventHubPropertiesMapOutput{})
	pulumi.RegisterOutputType(EventHubPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EventHubPropertiesResponseMapOutput{})
	pulumi.RegisterOutputType(FallbackRoutePropertiesOutput{})
	pulumi.RegisterOutputType(FallbackRoutePropertiesPtrOutput{})
	pulumi.RegisterOutputType(FallbackRoutePropertiesResponseOutput{})
	pulumi.RegisterOutputType(FallbackRoutePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(FeedbackPropertiesOutput{})
	pulumi.RegisterOutputType(FeedbackPropertiesPtrOutput{})
	pulumi.RegisterOutputType(FeedbackPropertiesResponseOutput{})
	pulumi.RegisterOutputType(FeedbackPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IotHubLocationDescriptionResponseOutput{})
	pulumi.RegisterOutputType(IotHubLocationDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(IotHubPropertiesOutput{})
	pulumi.RegisterOutputType(IotHubPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IotHubPropertiesDeviceStreamsOutput{})
	pulumi.RegisterOutputType(IotHubPropertiesDeviceStreamsPtrOutput{})
	pulumi.RegisterOutputType(IotHubPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IotHubPropertiesResponseDeviceStreamsOutput{})
	pulumi.RegisterOutputType(IotHubPropertiesResponseDeviceStreamsPtrOutput{})
	pulumi.RegisterOutputType(IotHubSkuInfoOutput{})
	pulumi.RegisterOutputType(IotHubSkuInfoResponseOutput{})
	pulumi.RegisterOutputType(IpFilterRuleOutput{})
	pulumi.RegisterOutputType(IpFilterRuleArrayOutput{})
	pulumi.RegisterOutputType(IpFilterRuleResponseOutput{})
	pulumi.RegisterOutputType(IpFilterRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(KEKIdentityOutput{})
	pulumi.RegisterOutputType(KEKIdentityPtrOutput{})
	pulumi.RegisterOutputType(KEKIdentityResponseOutput{})
	pulumi.RegisterOutputType(KEKIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyPropertiesArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(MessagingEndpointPropertiesOutput{})
	pulumi.RegisterOutputType(MessagingEndpointPropertiesMapOutput{})
	pulumi.RegisterOutputType(MessagingEndpointPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MessagingEndpointPropertiesResponseMapOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetIpRuleOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetIpRuleArrayOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetIpRuleResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetIpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPropertiesOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPropertiesPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPropertiesResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(RoutePropertiesOutput{})
	pulumi.RegisterOutputType(RoutePropertiesArrayOutput{})
	pulumi.RegisterOutputType(RoutePropertiesResponseOutput{})
	pulumi.RegisterOutputType(RoutePropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(RoutingEndpointsOutput{})
	pulumi.RegisterOutputType(RoutingEndpointsPtrOutput{})
	pulumi.RegisterOutputType(RoutingEndpointsResponseOutput{})
	pulumi.RegisterOutputType(RoutingEndpointsResponsePtrOutput{})
	pulumi.RegisterOutputType(RoutingEventHubPropertiesOutput{})
	pulumi.RegisterOutputType(RoutingEventHubPropertiesArrayOutput{})
	pulumi.RegisterOutputType(RoutingEventHubPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RoutingEventHubPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(RoutingPropertiesOutput{})
	pulumi.RegisterOutputType(RoutingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(RoutingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RoutingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusQueueEndpointPropertiesOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusQueueEndpointPropertiesArrayOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusQueueEndpointPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusTopicEndpointPropertiesOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusTopicEndpointPropertiesArrayOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusTopicEndpointPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(RoutingStorageContainerPropertiesOutput{})
	pulumi.RegisterOutputType(RoutingStorageContainerPropertiesArrayOutput{})
	pulumi.RegisterOutputType(RoutingStorageContainerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RoutingStorageContainerPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleArrayOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleResponseOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageEndpointPropertiesOutput{})
	pulumi.RegisterOutputType(StorageEndpointPropertiesMapOutput{})
	pulumi.RegisterOutputType(StorageEndpointPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageEndpointPropertiesResponseMapOutput{})
}

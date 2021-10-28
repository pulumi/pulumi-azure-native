


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssessmentStatusCode string

const (
	// The resource is healthy
	AssessmentStatusCodeHealthy = AssessmentStatusCode("Healthy")
	// The resource has a security issue that needs to be addressed
	AssessmentStatusCodeUnhealthy = AssessmentStatusCode("Unhealthy")
	// Assessment for this resource did not happen
	AssessmentStatusCodeNotApplicable = AssessmentStatusCode("NotApplicable")
)

func (AssessmentStatusCode) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatusCode)(nil)).Elem()
}

func (e AssessmentStatusCode) ToAssessmentStatusCodeOutput() AssessmentStatusCodeOutput {
	return pulumi.ToOutput(e).(AssessmentStatusCodeOutput)
}

func (e AssessmentStatusCode) ToAssessmentStatusCodeOutputWithContext(ctx context.Context) AssessmentStatusCodeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssessmentStatusCodeOutput)
}

func (e AssessmentStatusCode) ToAssessmentStatusCodePtrOutput() AssessmentStatusCodePtrOutput {
	return e.ToAssessmentStatusCodePtrOutputWithContext(context.Background())
}

func (e AssessmentStatusCode) ToAssessmentStatusCodePtrOutputWithContext(ctx context.Context) AssessmentStatusCodePtrOutput {
	return AssessmentStatusCode(e).ToAssessmentStatusCodeOutputWithContext(ctx).ToAssessmentStatusCodePtrOutputWithContext(ctx)
}

func (e AssessmentStatusCode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentStatusCode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentStatusCode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssessmentStatusCode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssessmentStatusCodeOutput struct{ *pulumi.OutputState }

func (AssessmentStatusCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatusCode)(nil)).Elem()
}

func (o AssessmentStatusCodeOutput) ToAssessmentStatusCodeOutput() AssessmentStatusCodeOutput {
	return o
}

func (o AssessmentStatusCodeOutput) ToAssessmentStatusCodeOutputWithContext(ctx context.Context) AssessmentStatusCodeOutput {
	return o
}

func (o AssessmentStatusCodeOutput) ToAssessmentStatusCodePtrOutput() AssessmentStatusCodePtrOutput {
	return o.ToAssessmentStatusCodePtrOutputWithContext(context.Background())
}

func (o AssessmentStatusCodeOutput) ToAssessmentStatusCodePtrOutputWithContext(ctx context.Context) AssessmentStatusCodePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentStatusCode) *AssessmentStatusCode {
		return &v
	}).(AssessmentStatusCodePtrOutput)
}

func (o AssessmentStatusCodeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssessmentStatusCodeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentStatusCode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssessmentStatusCodeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentStatusCodeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentStatusCode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssessmentStatusCodePtrOutput struct{ *pulumi.OutputState }

func (AssessmentStatusCodePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentStatusCode)(nil)).Elem()
}

func (o AssessmentStatusCodePtrOutput) ToAssessmentStatusCodePtrOutput() AssessmentStatusCodePtrOutput {
	return o
}

func (o AssessmentStatusCodePtrOutput) ToAssessmentStatusCodePtrOutputWithContext(ctx context.Context) AssessmentStatusCodePtrOutput {
	return o
}

func (o AssessmentStatusCodePtrOutput) Elem() AssessmentStatusCodeOutput {
	return o.ApplyT(func(v *AssessmentStatusCode) AssessmentStatusCode {
		if v != nil {
			return *v
		}
		var ret AssessmentStatusCode
		return ret
	}).(AssessmentStatusCodeOutput)
}

func (o AssessmentStatusCodePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentStatusCodePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssessmentStatusCode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AssessmentStatusCodeInput interface {
	pulumi.Input

	ToAssessmentStatusCodeOutput() AssessmentStatusCodeOutput
	ToAssessmentStatusCodeOutputWithContext(context.Context) AssessmentStatusCodeOutput
}

var assessmentStatusCodePtrType = reflect.TypeOf((**AssessmentStatusCode)(nil)).Elem()

type AssessmentStatusCodePtrInput interface {
	pulumi.Input

	ToAssessmentStatusCodePtrOutput() AssessmentStatusCodePtrOutput
	ToAssessmentStatusCodePtrOutputWithContext(context.Context) AssessmentStatusCodePtrOutput
}

type assessmentStatusCodePtr string

func AssessmentStatusCodePtr(v string) AssessmentStatusCodePtrInput {
	return (*assessmentStatusCodePtr)(&v)
}

func (*assessmentStatusCodePtr) ElementType() reflect.Type {
	return assessmentStatusCodePtrType
}

func (in *assessmentStatusCodePtr) ToAssessmentStatusCodePtrOutput() AssessmentStatusCodePtrOutput {
	return pulumi.ToOutput(in).(AssessmentStatusCodePtrOutput)
}

func (in *assessmentStatusCodePtr) ToAssessmentStatusCodePtrOutputWithContext(ctx context.Context) AssessmentStatusCodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssessmentStatusCodePtrOutput)
}

type AssessmentType string

const (
	// Azure Security Center managed assessments
	AssessmentTypeBuiltIn = AssessmentType("BuiltIn")
	// User defined policies that are automatically ingested from Azure Policy to Azure Security Center
	AssessmentTypeCustomPolicy = AssessmentType("CustomPolicy")
	// User assessments pushed directly by the user or other third party to Azure Security Center
	AssessmentTypeCustomerManaged = AssessmentType("CustomerManaged")
	// An assessment that was created by a verified 3rd party if the user connected it to ASC
	AssessmentTypeVerifiedPartner = AssessmentType("VerifiedPartner")
)

func (AssessmentType) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentType)(nil)).Elem()
}

func (e AssessmentType) ToAssessmentTypeOutput() AssessmentTypeOutput {
	return pulumi.ToOutput(e).(AssessmentTypeOutput)
}

func (e AssessmentType) ToAssessmentTypeOutputWithContext(ctx context.Context) AssessmentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssessmentTypeOutput)
}

func (e AssessmentType) ToAssessmentTypePtrOutput() AssessmentTypePtrOutput {
	return e.ToAssessmentTypePtrOutputWithContext(context.Background())
}

func (e AssessmentType) ToAssessmentTypePtrOutputWithContext(ctx context.Context) AssessmentTypePtrOutput {
	return AssessmentType(e).ToAssessmentTypeOutputWithContext(ctx).ToAssessmentTypePtrOutputWithContext(ctx)
}

func (e AssessmentType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssessmentType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssessmentTypeOutput struct{ *pulumi.OutputState }

func (AssessmentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentType)(nil)).Elem()
}

func (o AssessmentTypeOutput) ToAssessmentTypeOutput() AssessmentTypeOutput {
	return o
}

func (o AssessmentTypeOutput) ToAssessmentTypeOutputWithContext(ctx context.Context) AssessmentTypeOutput {
	return o
}

func (o AssessmentTypeOutput) ToAssessmentTypePtrOutput() AssessmentTypePtrOutput {
	return o.ToAssessmentTypePtrOutputWithContext(context.Background())
}

func (o AssessmentTypeOutput) ToAssessmentTypePtrOutputWithContext(ctx context.Context) AssessmentTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentType) *AssessmentType {
		return &v
	}).(AssessmentTypePtrOutput)
}

func (o AssessmentTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssessmentTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssessmentTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssessmentTypePtrOutput struct{ *pulumi.OutputState }

func (AssessmentTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentType)(nil)).Elem()
}

func (o AssessmentTypePtrOutput) ToAssessmentTypePtrOutput() AssessmentTypePtrOutput {
	return o
}

func (o AssessmentTypePtrOutput) ToAssessmentTypePtrOutputWithContext(ctx context.Context) AssessmentTypePtrOutput {
	return o
}

func (o AssessmentTypePtrOutput) Elem() AssessmentTypeOutput {
	return o.ApplyT(func(v *AssessmentType) AssessmentType {
		if v != nil {
			return *v
		}
		var ret AssessmentType
		return ret
	}).(AssessmentTypeOutput)
}

func (o AssessmentTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssessmentType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AssessmentTypeInput interface {
	pulumi.Input

	ToAssessmentTypeOutput() AssessmentTypeOutput
	ToAssessmentTypeOutputWithContext(context.Context) AssessmentTypeOutput
}

var assessmentTypePtrType = reflect.TypeOf((**AssessmentType)(nil)).Elem()

type AssessmentTypePtrInput interface {
	pulumi.Input

	ToAssessmentTypePtrOutput() AssessmentTypePtrOutput
	ToAssessmentTypePtrOutputWithContext(context.Context) AssessmentTypePtrOutput
}

type assessmentTypePtr string

func AssessmentTypePtr(v string) AssessmentTypePtrInput {
	return (*assessmentTypePtr)(&v)
}

func (*assessmentTypePtr) ElementType() reflect.Type {
	return assessmentTypePtrType
}

func (in *assessmentTypePtr) ToAssessmentTypePtrOutput() AssessmentTypePtrOutput {
	return pulumi.ToOutput(in).(AssessmentTypePtrOutput)
}

func (in *assessmentTypePtr) ToAssessmentTypePtrOutputWithContext(ctx context.Context) AssessmentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssessmentTypePtrOutput)
}

type Categories string

const (
	CategoriesCompute           = Categories("Compute")
	CategoriesNetworking        = Categories("Networking")
	CategoriesData              = Categories("Data")
	CategoriesIdentityAndAccess = Categories("IdentityAndAccess")
	CategoriesIoT               = Categories("IoT")
)

func (Categories) ElementType() reflect.Type {
	return reflect.TypeOf((*Categories)(nil)).Elem()
}

func (e Categories) ToCategoriesOutput() CategoriesOutput {
	return pulumi.ToOutput(e).(CategoriesOutput)
}

func (e Categories) ToCategoriesOutputWithContext(ctx context.Context) CategoriesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CategoriesOutput)
}

func (e Categories) ToCategoriesPtrOutput() CategoriesPtrOutput {
	return e.ToCategoriesPtrOutputWithContext(context.Background())
}

func (e Categories) ToCategoriesPtrOutputWithContext(ctx context.Context) CategoriesPtrOutput {
	return Categories(e).ToCategoriesOutputWithContext(ctx).ToCategoriesPtrOutputWithContext(ctx)
}

func (e Categories) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Categories) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Categories) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Categories) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CategoriesOutput struct{ *pulumi.OutputState }

func (CategoriesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Categories)(nil)).Elem()
}

func (o CategoriesOutput) ToCategoriesOutput() CategoriesOutput {
	return o
}

func (o CategoriesOutput) ToCategoriesOutputWithContext(ctx context.Context) CategoriesOutput {
	return o
}

func (o CategoriesOutput) ToCategoriesPtrOutput() CategoriesPtrOutput {
	return o.ToCategoriesPtrOutputWithContext(context.Background())
}

func (o CategoriesOutput) ToCategoriesPtrOutputWithContext(ctx context.Context) CategoriesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Categories) *Categories {
		return &v
	}).(CategoriesPtrOutput)
}

func (o CategoriesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CategoriesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Categories) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CategoriesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CategoriesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Categories) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CategoriesPtrOutput struct{ *pulumi.OutputState }

func (CategoriesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Categories)(nil)).Elem()
}

func (o CategoriesPtrOutput) ToCategoriesPtrOutput() CategoriesPtrOutput {
	return o
}

func (o CategoriesPtrOutput) ToCategoriesPtrOutputWithContext(ctx context.Context) CategoriesPtrOutput {
	return o
}

func (o CategoriesPtrOutput) Elem() CategoriesOutput {
	return o.ApplyT(func(v *Categories) Categories {
		if v != nil {
			return *v
		}
		var ret Categories
		return ret
	}).(CategoriesOutput)
}

func (o CategoriesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CategoriesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Categories) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CategoriesInput interface {
	pulumi.Input

	ToCategoriesOutput() CategoriesOutput
	ToCategoriesOutputWithContext(context.Context) CategoriesOutput
}

var categoriesPtrType = reflect.TypeOf((**Categories)(nil)).Elem()

type CategoriesPtrInput interface {
	pulumi.Input

	ToCategoriesPtrOutput() CategoriesPtrOutput
	ToCategoriesPtrOutputWithContext(context.Context) CategoriesPtrOutput
}

type categoriesPtr string

func CategoriesPtr(v string) CategoriesPtrInput {
	return (*categoriesPtr)(&v)
}

func (*categoriesPtr) ElementType() reflect.Type {
	return categoriesPtrType
}

func (in *categoriesPtr) ToCategoriesPtrOutput() CategoriesPtrOutput {
	return pulumi.ToOutput(in).(CategoriesPtrOutput)
}

func (in *categoriesPtr) ToCategoriesPtrOutputWithContext(ctx context.Context) CategoriesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CategoriesPtrOutput)
}

type ImplementationEffort string

const (
	ImplementationEffortLow      = ImplementationEffort("Low")
	ImplementationEffortModerate = ImplementationEffort("Moderate")
	ImplementationEffortHigh     = ImplementationEffort("High")
)

func (ImplementationEffort) ElementType() reflect.Type {
	return reflect.TypeOf((*ImplementationEffort)(nil)).Elem()
}

func (e ImplementationEffort) ToImplementationEffortOutput() ImplementationEffortOutput {
	return pulumi.ToOutput(e).(ImplementationEffortOutput)
}

func (e ImplementationEffort) ToImplementationEffortOutputWithContext(ctx context.Context) ImplementationEffortOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ImplementationEffortOutput)
}

func (e ImplementationEffort) ToImplementationEffortPtrOutput() ImplementationEffortPtrOutput {
	return e.ToImplementationEffortPtrOutputWithContext(context.Background())
}

func (e ImplementationEffort) ToImplementationEffortPtrOutputWithContext(ctx context.Context) ImplementationEffortPtrOutput {
	return ImplementationEffort(e).ToImplementationEffortOutputWithContext(ctx).ToImplementationEffortPtrOutputWithContext(ctx)
}

func (e ImplementationEffort) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ImplementationEffort) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ImplementationEffort) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ImplementationEffort) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ImplementationEffortOutput struct{ *pulumi.OutputState }

func (ImplementationEffortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImplementationEffort)(nil)).Elem()
}

func (o ImplementationEffortOutput) ToImplementationEffortOutput() ImplementationEffortOutput {
	return o
}

func (o ImplementationEffortOutput) ToImplementationEffortOutputWithContext(ctx context.Context) ImplementationEffortOutput {
	return o
}

func (o ImplementationEffortOutput) ToImplementationEffortPtrOutput() ImplementationEffortPtrOutput {
	return o.ToImplementationEffortPtrOutputWithContext(context.Background())
}

func (o ImplementationEffortOutput) ToImplementationEffortPtrOutputWithContext(ctx context.Context) ImplementationEffortPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImplementationEffort) *ImplementationEffort {
		return &v
	}).(ImplementationEffortPtrOutput)
}

func (o ImplementationEffortOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ImplementationEffortOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ImplementationEffort) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ImplementationEffortOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ImplementationEffortOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ImplementationEffort) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ImplementationEffortPtrOutput struct{ *pulumi.OutputState }

func (ImplementationEffortPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImplementationEffort)(nil)).Elem()
}

func (o ImplementationEffortPtrOutput) ToImplementationEffortPtrOutput() ImplementationEffortPtrOutput {
	return o
}

func (o ImplementationEffortPtrOutput) ToImplementationEffortPtrOutputWithContext(ctx context.Context) ImplementationEffortPtrOutput {
	return o
}

func (o ImplementationEffortPtrOutput) Elem() ImplementationEffortOutput {
	return o.ApplyT(func(v *ImplementationEffort) ImplementationEffort {
		if v != nil {
			return *v
		}
		var ret ImplementationEffort
		return ret
	}).(ImplementationEffortOutput)
}

func (o ImplementationEffortPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ImplementationEffortPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ImplementationEffort) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ImplementationEffortInput interface {
	pulumi.Input

	ToImplementationEffortOutput() ImplementationEffortOutput
	ToImplementationEffortOutputWithContext(context.Context) ImplementationEffortOutput
}

var implementationEffortPtrType = reflect.TypeOf((**ImplementationEffort)(nil)).Elem()

type ImplementationEffortPtrInput interface {
	pulumi.Input

	ToImplementationEffortPtrOutput() ImplementationEffortPtrOutput
	ToImplementationEffortPtrOutputWithContext(context.Context) ImplementationEffortPtrOutput
}

type implementationEffortPtr string

func ImplementationEffortPtr(v string) ImplementationEffortPtrInput {
	return (*implementationEffortPtr)(&v)
}

func (*implementationEffortPtr) ElementType() reflect.Type {
	return implementationEffortPtrType
}

func (in *implementationEffortPtr) ToImplementationEffortPtrOutput() ImplementationEffortPtrOutput {
	return pulumi.ToOutput(in).(ImplementationEffortPtrOutput)
}

func (in *implementationEffortPtr) ToImplementationEffortPtrOutputWithContext(ctx context.Context) ImplementationEffortPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ImplementationEffortPtrOutput)
}

type Protocol string

const (
	ProtocolTCP = Protocol("TCP")
	ProtocolUDP = Protocol("UDP")
	ProtocolAll = Protocol("*")
)

func (Protocol) ElementType() reflect.Type {
	return reflect.TypeOf((*Protocol)(nil)).Elem()
}

func (e Protocol) ToProtocolOutput() ProtocolOutput {
	return pulumi.ToOutput(e).(ProtocolOutput)
}

func (e Protocol) ToProtocolOutputWithContext(ctx context.Context) ProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtocolOutput)
}

func (e Protocol) ToProtocolPtrOutput() ProtocolPtrOutput {
	return e.ToProtocolPtrOutputWithContext(context.Background())
}

func (e Protocol) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return Protocol(e).ToProtocolOutputWithContext(ctx).ToProtocolPtrOutputWithContext(ctx)
}

func (e Protocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Protocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Protocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Protocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtocolOutput struct{ *pulumi.OutputState }

func (ProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Protocol)(nil)).Elem()
}

func (o ProtocolOutput) ToProtocolOutput() ProtocolOutput {
	return o
}

func (o ProtocolOutput) ToProtocolOutputWithContext(ctx context.Context) ProtocolOutput {
	return o
}

func (o ProtocolOutput) ToProtocolPtrOutput() ProtocolPtrOutput {
	return o.ToProtocolPtrOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Protocol) *Protocol {
		return &v
	}).(ProtocolPtrOutput)
}

func (o ProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Protocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Protocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtocolPtrOutput struct{ *pulumi.OutputState }

func (ProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Protocol)(nil)).Elem()
}

func (o ProtocolPtrOutput) ToProtocolPtrOutput() ProtocolPtrOutput {
	return o
}

func (o ProtocolPtrOutput) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return o
}

func (o ProtocolPtrOutput) Elem() ProtocolOutput {
	return o.ApplyT(func(v *Protocol) Protocol {
		if v != nil {
			return *v
		}
		var ret Protocol
		return ret
	}).(ProtocolOutput)
}

func (o ProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Protocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProtocolInput interface {
	pulumi.Input

	ToProtocolOutput() ProtocolOutput
	ToProtocolOutputWithContext(context.Context) ProtocolOutput
}

var protocolPtrType = reflect.TypeOf((**Protocol)(nil)).Elem()

type ProtocolPtrInput interface {
	pulumi.Input

	ToProtocolPtrOutput() ProtocolPtrOutput
	ToProtocolPtrOutputWithContext(context.Context) ProtocolPtrOutput
}

type protocolPtr string

func ProtocolPtr(v string) ProtocolPtrInput {
	return (*protocolPtr)(&v)
}

func (*protocolPtr) ElementType() reflect.Type {
	return protocolPtrType
}

func (in *protocolPtr) ToProtocolPtrOutput() ProtocolPtrOutput {
	return pulumi.ToOutput(in).(ProtocolPtrOutput)
}

func (in *protocolPtr) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtocolPtrOutput)
}

type Severity string

const (
	SeverityLow    = Severity("Low")
	SeverityMedium = Severity("Medium")
	SeverityHigh   = Severity("High")
)

func (Severity) ElementType() reflect.Type {
	return reflect.TypeOf((*Severity)(nil)).Elem()
}

func (e Severity) ToSeverityOutput() SeverityOutput {
	return pulumi.ToOutput(e).(SeverityOutput)
}

func (e Severity) ToSeverityOutputWithContext(ctx context.Context) SeverityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SeverityOutput)
}

func (e Severity) ToSeverityPtrOutput() SeverityPtrOutput {
	return e.ToSeverityPtrOutputWithContext(context.Background())
}

func (e Severity) ToSeverityPtrOutputWithContext(ctx context.Context) SeverityPtrOutput {
	return Severity(e).ToSeverityOutputWithContext(ctx).ToSeverityPtrOutputWithContext(ctx)
}

func (e Severity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Severity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Severity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Severity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SeverityOutput struct{ *pulumi.OutputState }

func (SeverityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Severity)(nil)).Elem()
}

func (o SeverityOutput) ToSeverityOutput() SeverityOutput {
	return o
}

func (o SeverityOutput) ToSeverityOutputWithContext(ctx context.Context) SeverityOutput {
	return o
}

func (o SeverityOutput) ToSeverityPtrOutput() SeverityPtrOutput {
	return o.ToSeverityPtrOutputWithContext(context.Background())
}

func (o SeverityOutput) ToSeverityPtrOutputWithContext(ctx context.Context) SeverityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Severity) *Severity {
		return &v
	}).(SeverityPtrOutput)
}

func (o SeverityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SeverityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Severity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SeverityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SeverityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Severity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SeverityPtrOutput struct{ *pulumi.OutputState }

func (SeverityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Severity)(nil)).Elem()
}

func (o SeverityPtrOutput) ToSeverityPtrOutput() SeverityPtrOutput {
	return o
}

func (o SeverityPtrOutput) ToSeverityPtrOutputWithContext(ctx context.Context) SeverityPtrOutput {
	return o
}

func (o SeverityPtrOutput) Elem() SeverityOutput {
	return o.ApplyT(func(v *Severity) Severity {
		if v != nil {
			return *v
		}
		var ret Severity
		return ret
	}).(SeverityOutput)
}

func (o SeverityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SeverityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Severity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SeverityInput interface {
	pulumi.Input

	ToSeverityOutput() SeverityOutput
	ToSeverityOutputWithContext(context.Context) SeverityOutput
}

var severityPtrType = reflect.TypeOf((**Severity)(nil)).Elem()

type SeverityPtrInput interface {
	pulumi.Input

	ToSeverityPtrOutput() SeverityPtrOutput
	ToSeverityPtrOutputWithContext(context.Context) SeverityPtrOutput
}

type severityPtr string

func SeverityPtr(v string) SeverityPtrInput {
	return (*severityPtr)(&v)
}

func (*severityPtr) ElementType() reflect.Type {
	return severityPtrType
}

func (in *severityPtr) ToSeverityPtrOutput() SeverityPtrOutput {
	return pulumi.ToOutput(in).(SeverityPtrOutput)
}

func (in *severityPtr) ToSeverityPtrOutputWithContext(ctx context.Context) SeverityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SeverityPtrOutput)
}

type Source string

const (
	// Resource is in Azure
	SourceAzure = Source("Azure")
	// Resource in an on premise machine connected to Azure cloud
	SourceOnPremise = Source("OnPremise")
	// SQL Resource in an on premise machine connected to Azure cloud
	SourceOnPremiseSql = Source("OnPremiseSql")
)

func (Source) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (e Source) ToSourceOutput() SourceOutput {
	return pulumi.ToOutput(e).(SourceOutput)
}

func (e Source) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceOutput)
}

func (e Source) ToSourcePtrOutput() SourcePtrOutput {
	return e.ToSourcePtrOutputWithContext(context.Background())
}

func (e Source) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return Source(e).ToSourceOutputWithContext(ctx).ToSourcePtrOutputWithContext(ctx)
}

func (e Source) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Source) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Source) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Source) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceOutput struct{ *pulumi.OutputState }

func (SourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (o SourceOutput) ToSourceOutput() SourceOutput {
	return o
}

func (o SourceOutput) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return o
}

func (o SourceOutput) ToSourcePtrOutput() SourcePtrOutput {
	return o.ToSourcePtrOutputWithContext(context.Background())
}

func (o SourceOutput) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Source) *Source {
		return &v
	}).(SourcePtrOutput)
}

func (o SourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Source) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Source) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourcePtrOutput struct{ *pulumi.OutputState }

func (SourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Source)(nil)).Elem()
}

func (o SourcePtrOutput) ToSourcePtrOutput() SourcePtrOutput {
	return o
}

func (o SourcePtrOutput) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return o
}

func (o SourcePtrOutput) Elem() SourceOutput {
	return o.ApplyT(func(v *Source) Source {
		if v != nil {
			return *v
		}
		var ret Source
		return ret
	}).(SourceOutput)
}

func (o SourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Source) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SourceInput interface {
	pulumi.Input

	ToSourceOutput() SourceOutput
	ToSourceOutputWithContext(context.Context) SourceOutput
}

var sourcePtrType = reflect.TypeOf((**Source)(nil)).Elem()

type SourcePtrInput interface {
	pulumi.Input

	ToSourcePtrOutput() SourcePtrOutput
	ToSourcePtrOutputWithContext(context.Context) SourcePtrOutput
}

type sourcePtr string

func SourcePtr(v string) SourcePtrInput {
	return (*sourcePtr)(&v)
}

func (*sourcePtr) ElementType() reflect.Type {
	return sourcePtrType
}

func (in *sourcePtr) ToSourcePtrOutput() SourcePtrOutput {
	return pulumi.ToOutput(in).(SourcePtrOutput)
}

func (in *sourcePtr) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourcePtrOutput)
}

type Status string

const (
	StatusRevoked   = Status("Revoked")
	StatusInitiated = Status("Initiated")
)

func (Status) ElementType() reflect.Type {
	return reflect.TypeOf((*Status)(nil)).Elem()
}

func (e Status) ToStatusOutput() StatusOutput {
	return pulumi.ToOutput(e).(StatusOutput)
}

func (e Status) ToStatusOutputWithContext(ctx context.Context) StatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StatusOutput)
}

func (e Status) ToStatusPtrOutput() StatusPtrOutput {
	return e.ToStatusPtrOutputWithContext(context.Background())
}

func (e Status) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return Status(e).ToStatusOutputWithContext(ctx).ToStatusPtrOutputWithContext(ctx)
}

func (e Status) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Status) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Status) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Status) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StatusOutput struct{ *pulumi.OutputState }

func (StatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Status)(nil)).Elem()
}

func (o StatusOutput) ToStatusOutput() StatusOutput {
	return o
}

func (o StatusOutput) ToStatusOutputWithContext(ctx context.Context) StatusOutput {
	return o
}

func (o StatusOutput) ToStatusPtrOutput() StatusPtrOutput {
	return o.ToStatusPtrOutputWithContext(context.Background())
}

func (o StatusOutput) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Status) *Status {
		return &v
	}).(StatusPtrOutput)
}

func (o StatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Status) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Status) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatusPtrOutput struct{ *pulumi.OutputState }

func (StatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Status)(nil)).Elem()
}

func (o StatusPtrOutput) ToStatusPtrOutput() StatusPtrOutput {
	return o
}

func (o StatusPtrOutput) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return o
}

func (o StatusPtrOutput) Elem() StatusOutput {
	return o.ApplyT(func(v *Status) Status {
		if v != nil {
			return *v
		}
		var ret Status
		return ret
	}).(StatusOutput)
}

func (o StatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Status) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StatusInput interface {
	pulumi.Input

	ToStatusOutput() StatusOutput
	ToStatusOutputWithContext(context.Context) StatusOutput
}

var statusPtrType = reflect.TypeOf((**Status)(nil)).Elem()

type StatusPtrInput interface {
	pulumi.Input

	ToStatusPtrOutput() StatusPtrOutput
	ToStatusPtrOutputWithContext(context.Context) StatusPtrOutput
}

type statusPtr string

func StatusPtr(v string) StatusPtrInput {
	return (*statusPtr)(&v)
}

func (*statusPtr) ElementType() reflect.Type {
	return statusPtrType
}

func (in *statusPtr) ToStatusPtrOutput() StatusPtrOutput {
	return pulumi.ToOutput(in).(StatusPtrOutput)
}

func (in *statusPtr) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatusPtrOutput)
}

type StatusReason string

const (
	StatusReasonExpired               = StatusReason("Expired")
	StatusReasonUserRequested         = StatusReason("UserRequested")
	StatusReasonNewerRequestInitiated = StatusReason("NewerRequestInitiated")
)

func (StatusReason) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusReason)(nil)).Elem()
}

func (e StatusReason) ToStatusReasonOutput() StatusReasonOutput {
	return pulumi.ToOutput(e).(StatusReasonOutput)
}

func (e StatusReason) ToStatusReasonOutputWithContext(ctx context.Context) StatusReasonOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StatusReasonOutput)
}

func (e StatusReason) ToStatusReasonPtrOutput() StatusReasonPtrOutput {
	return e.ToStatusReasonPtrOutputWithContext(context.Background())
}

func (e StatusReason) ToStatusReasonPtrOutputWithContext(ctx context.Context) StatusReasonPtrOutput {
	return StatusReason(e).ToStatusReasonOutputWithContext(ctx).ToStatusReasonPtrOutputWithContext(ctx)
}

func (e StatusReason) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusReason) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusReason) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StatusReason) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StatusReasonOutput struct{ *pulumi.OutputState }

func (StatusReasonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusReason)(nil)).Elem()
}

func (o StatusReasonOutput) ToStatusReasonOutput() StatusReasonOutput {
	return o
}

func (o StatusReasonOutput) ToStatusReasonOutputWithContext(ctx context.Context) StatusReasonOutput {
	return o
}

func (o StatusReasonOutput) ToStatusReasonPtrOutput() StatusReasonPtrOutput {
	return o.ToStatusReasonPtrOutputWithContext(context.Background())
}

func (o StatusReasonOutput) ToStatusReasonPtrOutputWithContext(ctx context.Context) StatusReasonPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StatusReason) *StatusReason {
		return &v
	}).(StatusReasonPtrOutput)
}

func (o StatusReasonOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StatusReasonOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusReason) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StatusReasonOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusReasonOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusReason) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatusReasonPtrOutput struct{ *pulumi.OutputState }

func (StatusReasonPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatusReason)(nil)).Elem()
}

func (o StatusReasonPtrOutput) ToStatusReasonPtrOutput() StatusReasonPtrOutput {
	return o
}

func (o StatusReasonPtrOutput) ToStatusReasonPtrOutputWithContext(ctx context.Context) StatusReasonPtrOutput {
	return o
}

func (o StatusReasonPtrOutput) Elem() StatusReasonOutput {
	return o.ApplyT(func(v *StatusReason) StatusReason {
		if v != nil {
			return *v
		}
		var ret StatusReason
		return ret
	}).(StatusReasonOutput)
}

func (o StatusReasonPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusReasonPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StatusReason) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StatusReasonInput interface {
	pulumi.Input

	ToStatusReasonOutput() StatusReasonOutput
	ToStatusReasonOutputWithContext(context.Context) StatusReasonOutput
}

var statusReasonPtrType = reflect.TypeOf((**StatusReason)(nil)).Elem()

type StatusReasonPtrInput interface {
	pulumi.Input

	ToStatusReasonPtrOutput() StatusReasonPtrOutput
	ToStatusReasonPtrOutputWithContext(context.Context) StatusReasonPtrOutput
}

type statusReasonPtr string

func StatusReasonPtr(v string) StatusReasonPtrInput {
	return (*statusReasonPtr)(&v)
}

func (*statusReasonPtr) ElementType() reflect.Type {
	return statusReasonPtrType
}

func (in *statusReasonPtr) ToStatusReasonPtrOutput() StatusReasonPtrOutput {
	return pulumi.ToOutput(in).(StatusReasonPtrOutput)
}

func (in *statusReasonPtr) ToStatusReasonPtrOutputWithContext(ctx context.Context) StatusReasonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatusReasonPtrOutput)
}

type Threats string

const (
	ThreatsAccountBreach        = Threats("accountBreach")
	ThreatsDataExfiltration     = Threats("dataExfiltration")
	ThreatsDataSpillage         = Threats("dataSpillage")
	ThreatsMaliciousInsider     = Threats("maliciousInsider")
	ThreatsElevationOfPrivilege = Threats("elevationOfPrivilege")
	ThreatsThreatResistance     = Threats("threatResistance")
	ThreatsMissingCoverage      = Threats("missingCoverage")
	ThreatsDenialOfService      = Threats("denialOfService")
)

func (Threats) ElementType() reflect.Type {
	return reflect.TypeOf((*Threats)(nil)).Elem()
}

func (e Threats) ToThreatsOutput() ThreatsOutput {
	return pulumi.ToOutput(e).(ThreatsOutput)
}

func (e Threats) ToThreatsOutputWithContext(ctx context.Context) ThreatsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ThreatsOutput)
}

func (e Threats) ToThreatsPtrOutput() ThreatsPtrOutput {
	return e.ToThreatsPtrOutputWithContext(context.Background())
}

func (e Threats) ToThreatsPtrOutputWithContext(ctx context.Context) ThreatsPtrOutput {
	return Threats(e).ToThreatsOutputWithContext(ctx).ToThreatsPtrOutputWithContext(ctx)
}

func (e Threats) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Threats) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Threats) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Threats) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ThreatsOutput struct{ *pulumi.OutputState }

func (ThreatsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Threats)(nil)).Elem()
}

func (o ThreatsOutput) ToThreatsOutput() ThreatsOutput {
	return o
}

func (o ThreatsOutput) ToThreatsOutputWithContext(ctx context.Context) ThreatsOutput {
	return o
}

func (o ThreatsOutput) ToThreatsPtrOutput() ThreatsPtrOutput {
	return o.ToThreatsPtrOutputWithContext(context.Background())
}

func (o ThreatsOutput) ToThreatsPtrOutputWithContext(ctx context.Context) ThreatsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Threats) *Threats {
		return &v
	}).(ThreatsPtrOutput)
}

func (o ThreatsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ThreatsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Threats) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ThreatsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ThreatsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Threats) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ThreatsPtrOutput struct{ *pulumi.OutputState }

func (ThreatsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Threats)(nil)).Elem()
}

func (o ThreatsPtrOutput) ToThreatsPtrOutput() ThreatsPtrOutput {
	return o
}

func (o ThreatsPtrOutput) ToThreatsPtrOutputWithContext(ctx context.Context) ThreatsPtrOutput {
	return o
}

func (o ThreatsPtrOutput) Elem() ThreatsOutput {
	return o.ApplyT(func(v *Threats) Threats {
		if v != nil {
			return *v
		}
		var ret Threats
		return ret
	}).(ThreatsOutput)
}

func (o ThreatsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ThreatsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Threats) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ThreatsInput interface {
	pulumi.Input

	ToThreatsOutput() ThreatsOutput
	ToThreatsOutputWithContext(context.Context) ThreatsOutput
}

var threatsPtrType = reflect.TypeOf((**Threats)(nil)).Elem()

type ThreatsPtrInput interface {
	pulumi.Input

	ToThreatsPtrOutput() ThreatsPtrOutput
	ToThreatsPtrOutputWithContext(context.Context) ThreatsPtrOutput
}

type threatsPtr string

func ThreatsPtr(v string) ThreatsPtrInput {
	return (*threatsPtr)(&v)
}

func (*threatsPtr) ElementType() reflect.Type {
	return threatsPtrType
}

func (in *threatsPtr) ToThreatsPtrOutput() ThreatsPtrOutput {
	return pulumi.ToOutput(in).(ThreatsPtrOutput)
}

func (in *threatsPtr) ToThreatsPtrOutputWithContext(ctx context.Context) ThreatsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ThreatsPtrOutput)
}

type UserImpact string

const (
	UserImpactLow      = UserImpact("Low")
	UserImpactModerate = UserImpact("Moderate")
	UserImpactHigh     = UserImpact("High")
)

func (UserImpact) ElementType() reflect.Type {
	return reflect.TypeOf((*UserImpact)(nil)).Elem()
}

func (e UserImpact) ToUserImpactOutput() UserImpactOutput {
	return pulumi.ToOutput(e).(UserImpactOutput)
}

func (e UserImpact) ToUserImpactOutputWithContext(ctx context.Context) UserImpactOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UserImpactOutput)
}

func (e UserImpact) ToUserImpactPtrOutput() UserImpactPtrOutput {
	return e.ToUserImpactPtrOutputWithContext(context.Background())
}

func (e UserImpact) ToUserImpactPtrOutputWithContext(ctx context.Context) UserImpactPtrOutput {
	return UserImpact(e).ToUserImpactOutputWithContext(ctx).ToUserImpactPtrOutputWithContext(ctx)
}

func (e UserImpact) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserImpact) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserImpact) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UserImpact) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UserImpactOutput struct{ *pulumi.OutputState }

func (UserImpactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserImpact)(nil)).Elem()
}

func (o UserImpactOutput) ToUserImpactOutput() UserImpactOutput {
	return o
}

func (o UserImpactOutput) ToUserImpactOutputWithContext(ctx context.Context) UserImpactOutput {
	return o
}

func (o UserImpactOutput) ToUserImpactPtrOutput() UserImpactPtrOutput {
	return o.ToUserImpactPtrOutputWithContext(context.Background())
}

func (o UserImpactOutput) ToUserImpactPtrOutputWithContext(ctx context.Context) UserImpactPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserImpact) *UserImpact {
		return &v
	}).(UserImpactPtrOutput)
}

func (o UserImpactOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UserImpactOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserImpact) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UserImpactOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserImpactOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserImpact) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UserImpactPtrOutput struct{ *pulumi.OutputState }

func (UserImpactPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserImpact)(nil)).Elem()
}

func (o UserImpactPtrOutput) ToUserImpactPtrOutput() UserImpactPtrOutput {
	return o
}

func (o UserImpactPtrOutput) ToUserImpactPtrOutputWithContext(ctx context.Context) UserImpactPtrOutput {
	return o
}

func (o UserImpactPtrOutput) Elem() UserImpactOutput {
	return o.ApplyT(func(v *UserImpact) UserImpact {
		if v != nil {
			return *v
		}
		var ret UserImpact
		return ret
	}).(UserImpactOutput)
}

func (o UserImpactPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserImpactPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UserImpact) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UserImpactInput interface {
	pulumi.Input

	ToUserImpactOutput() UserImpactOutput
	ToUserImpactOutputWithContext(context.Context) UserImpactOutput
}

var userImpactPtrType = reflect.TypeOf((**UserImpact)(nil)).Elem()

type UserImpactPtrInput interface {
	pulumi.Input

	ToUserImpactPtrOutput() UserImpactPtrOutput
	ToUserImpactPtrOutputWithContext(context.Context) UserImpactPtrOutput
}

type userImpactPtr string

func UserImpactPtr(v string) UserImpactPtrInput {
	return (*userImpactPtr)(&v)
}

func (*userImpactPtr) ElementType() reflect.Type {
	return userImpactPtrType
}

func (in *userImpactPtr) ToUserImpactPtrOutput() UserImpactPtrOutput {
	return pulumi.ToOutput(in).(UserImpactPtrOutput)
}

func (in *userImpactPtr) ToUserImpactPtrOutputWithContext(ctx context.Context) UserImpactPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UserImpactPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentStatusCodeInput)(nil)).Elem(), AssessmentStatusCode("Healthy"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentStatusCodePtrInput)(nil)).Elem(), AssessmentStatusCode("Healthy"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentTypeInput)(nil)).Elem(), AssessmentType("BuiltIn"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentTypePtrInput)(nil)).Elem(), AssessmentType("BuiltIn"))
	pulumi.RegisterInputType(reflect.TypeOf((*CategoriesInput)(nil)).Elem(), Categories("Compute"))
	pulumi.RegisterInputType(reflect.TypeOf((*CategoriesPtrInput)(nil)).Elem(), Categories("Compute"))
	pulumi.RegisterInputType(reflect.TypeOf((*ImplementationEffortInput)(nil)).Elem(), ImplementationEffort("Low"))
	pulumi.RegisterInputType(reflect.TypeOf((*ImplementationEffortPtrInput)(nil)).Elem(), ImplementationEffort("Low"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProtocolInput)(nil)).Elem(), Protocol("TCP"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProtocolPtrInput)(nil)).Elem(), Protocol("TCP"))
	pulumi.RegisterInputType(reflect.TypeOf((*SeverityInput)(nil)).Elem(), Severity("Low"))
	pulumi.RegisterInputType(reflect.TypeOf((*SeverityPtrInput)(nil)).Elem(), Severity("Low"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourceInput)(nil)).Elem(), Source("Azure"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePtrInput)(nil)).Elem(), Source("Azure"))
	pulumi.RegisterInputType(reflect.TypeOf((*StatusInput)(nil)).Elem(), Status("Revoked"))
	pulumi.RegisterInputType(reflect.TypeOf((*StatusPtrInput)(nil)).Elem(), Status("Revoked"))
	pulumi.RegisterInputType(reflect.TypeOf((*StatusReasonInput)(nil)).Elem(), StatusReason("Expired"))
	pulumi.RegisterInputType(reflect.TypeOf((*StatusReasonPtrInput)(nil)).Elem(), StatusReason("Expired"))
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatsInput)(nil)).Elem(), Threats("accountBreach"))
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatsPtrInput)(nil)).Elem(), Threats("accountBreach"))
	pulumi.RegisterInputType(reflect.TypeOf((*UserImpactInput)(nil)).Elem(), UserImpact("Low"))
	pulumi.RegisterInputType(reflect.TypeOf((*UserImpactPtrInput)(nil)).Elem(), UserImpact("Low"))
	pulumi.RegisterOutputType(AssessmentStatusCodeOutput{})
	pulumi.RegisterOutputType(AssessmentStatusCodePtrOutput{})
	pulumi.RegisterOutputType(AssessmentTypeOutput{})
	pulumi.RegisterOutputType(AssessmentTypePtrOutput{})
	pulumi.RegisterOutputType(CategoriesOutput{})
	pulumi.RegisterOutputType(CategoriesPtrOutput{})
	pulumi.RegisterOutputType(ImplementationEffortOutput{})
	pulumi.RegisterOutputType(ImplementationEffortPtrOutput{})
	pulumi.RegisterOutputType(ProtocolOutput{})
	pulumi.RegisterOutputType(ProtocolPtrOutput{})
	pulumi.RegisterOutputType(SeverityOutput{})
	pulumi.RegisterOutputType(SeverityPtrOutput{})
	pulumi.RegisterOutputType(SourceOutput{})
	pulumi.RegisterOutputType(SourcePtrOutput{})
	pulumi.RegisterOutputType(StatusOutput{})
	pulumi.RegisterOutputType(StatusPtrOutput{})
	pulumi.RegisterOutputType(StatusReasonOutput{})
	pulumi.RegisterOutputType(StatusReasonPtrOutput{})
	pulumi.RegisterOutputType(ThreatsOutput{})
	pulumi.RegisterOutputType(ThreatsPtrOutput{})
	pulumi.RegisterOutputType(UserImpactOutput{})
	pulumi.RegisterOutputType(UserImpactPtrOutput{})
}

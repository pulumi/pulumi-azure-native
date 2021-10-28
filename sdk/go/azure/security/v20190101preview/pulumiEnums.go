


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionType string

const (
	ActionTypeLogicApp  = ActionType("LogicApp")
	ActionTypeEventHub  = ActionType("EventHub")
	ActionTypeWorkspace = ActionType("Workspace")
)

func (ActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionType)(nil)).Elem()
}

func (e ActionType) ToActionTypeOutput() ActionTypeOutput {
	return pulumi.ToOutput(e).(ActionTypeOutput)
}

func (e ActionType) ToActionTypeOutputWithContext(ctx context.Context) ActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ActionTypeOutput)
}

func (e ActionType) ToActionTypePtrOutput() ActionTypePtrOutput {
	return e.ToActionTypePtrOutputWithContext(context.Background())
}

func (e ActionType) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return ActionType(e).ToActionTypeOutputWithContext(ctx).ToActionTypePtrOutputWithContext(ctx)
}

func (e ActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ActionTypeOutput struct{ *pulumi.OutputState }

func (ActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionType)(nil)).Elem()
}

func (o ActionTypeOutput) ToActionTypeOutput() ActionTypeOutput {
	return o
}

func (o ActionTypeOutput) ToActionTypeOutputWithContext(ctx context.Context) ActionTypeOutput {
	return o
}

func (o ActionTypeOutput) ToActionTypePtrOutput() ActionTypePtrOutput {
	return o.ToActionTypePtrOutputWithContext(context.Background())
}

func (o ActionTypeOutput) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionType) *ActionType {
		return &v
	}).(ActionTypePtrOutput)
}

func (o ActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ActionTypePtrOutput struct{ *pulumi.OutputState }

func (ActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionType)(nil)).Elem()
}

func (o ActionTypePtrOutput) ToActionTypePtrOutput() ActionTypePtrOutput {
	return o
}

func (o ActionTypePtrOutput) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return o
}

func (o ActionTypePtrOutput) Elem() ActionTypeOutput {
	return o.ApplyT(func(v *ActionType) ActionType {
		if v != nil {
			return *v
		}
		var ret ActionType
		return ret
	}).(ActionTypeOutput)
}

func (o ActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ActionTypeInput interface {
	pulumi.Input

	ToActionTypeOutput() ActionTypeOutput
	ToActionTypeOutputWithContext(context.Context) ActionTypeOutput
}

var actionTypePtrType = reflect.TypeOf((**ActionType)(nil)).Elem()

type ActionTypePtrInput interface {
	pulumi.Input

	ToActionTypePtrOutput() ActionTypePtrOutput
	ToActionTypePtrOutputWithContext(context.Context) ActionTypePtrOutput
}

type actionTypePtr string

func ActionTypePtr(v string) ActionTypePtrInput {
	return (*actionTypePtr)(&v)
}

func (*actionTypePtr) ElementType() reflect.Type {
	return actionTypePtrType
}

func (in *actionTypePtr) ToActionTypePtrOutput() ActionTypePtrOutput {
	return pulumi.ToOutput(in).(ActionTypePtrOutput)
}

func (in *actionTypePtr) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ActionTypePtrOutput)
}

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

type EventSource string

const (
	EventSourceAssessments                            = EventSource("Assessments")
	EventSourceAssessmentsSnapshot                    = EventSource("AssessmentsSnapshot")
	EventSourceSubAssessments                         = EventSource("SubAssessments")
	EventSourceSubAssessmentsSnapshot                 = EventSource("SubAssessmentsSnapshot")
	EventSourceAlerts                                 = EventSource("Alerts")
	EventSourceSecureScores                           = EventSource("SecureScores")
	EventSourceSecureScoresSnapshot                   = EventSource("SecureScoresSnapshot")
	EventSourceSecureScoreControls                    = EventSource("SecureScoreControls")
	EventSourceSecureScoreControlsSnapshot            = EventSource("SecureScoreControlsSnapshot")
	EventSourceRegulatoryComplianceAssessment         = EventSource("RegulatoryComplianceAssessment")
	EventSourceRegulatoryComplianceAssessmentSnapshot = EventSource("RegulatoryComplianceAssessmentSnapshot")
)

func (EventSource) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSource)(nil)).Elem()
}

func (e EventSource) ToEventSourceOutput() EventSourceOutput {
	return pulumi.ToOutput(e).(EventSourceOutput)
}

func (e EventSource) ToEventSourceOutputWithContext(ctx context.Context) EventSourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EventSourceOutput)
}

func (e EventSource) ToEventSourcePtrOutput() EventSourcePtrOutput {
	return e.ToEventSourcePtrOutputWithContext(context.Background())
}

func (e EventSource) ToEventSourcePtrOutputWithContext(ctx context.Context) EventSourcePtrOutput {
	return EventSource(e).ToEventSourceOutputWithContext(ctx).ToEventSourcePtrOutputWithContext(ctx)
}

func (e EventSource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EventSource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EventSourceOutput struct{ *pulumi.OutputState }

func (EventSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSource)(nil)).Elem()
}

func (o EventSourceOutput) ToEventSourceOutput() EventSourceOutput {
	return o
}

func (o EventSourceOutput) ToEventSourceOutputWithContext(ctx context.Context) EventSourceOutput {
	return o
}

func (o EventSourceOutput) ToEventSourcePtrOutput() EventSourcePtrOutput {
	return o.ToEventSourcePtrOutputWithContext(context.Background())
}

func (o EventSourceOutput) ToEventSourcePtrOutputWithContext(ctx context.Context) EventSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSource) *EventSource {
		return &v
	}).(EventSourcePtrOutput)
}

func (o EventSourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EventSourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EventSourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EventSourcePtrOutput struct{ *pulumi.OutputState }

func (EventSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSource)(nil)).Elem()
}

func (o EventSourcePtrOutput) ToEventSourcePtrOutput() EventSourcePtrOutput {
	return o
}

func (o EventSourcePtrOutput) ToEventSourcePtrOutputWithContext(ctx context.Context) EventSourcePtrOutput {
	return o
}

func (o EventSourcePtrOutput) Elem() EventSourceOutput {
	return o.ApplyT(func(v *EventSource) EventSource {
		if v != nil {
			return *v
		}
		var ret EventSource
		return ret
	}).(EventSourceOutput)
}

func (o EventSourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EventSource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EventSourceInput interface {
	pulumi.Input

	ToEventSourceOutput() EventSourceOutput
	ToEventSourceOutputWithContext(context.Context) EventSourceOutput
}

var eventSourcePtrType = reflect.TypeOf((**EventSource)(nil)).Elem()

type EventSourcePtrInput interface {
	pulumi.Input

	ToEventSourcePtrOutput() EventSourcePtrOutput
	ToEventSourcePtrOutputWithContext(context.Context) EventSourcePtrOutput
}

type eventSourcePtr string

func EventSourcePtr(v string) EventSourcePtrInput {
	return (*eventSourcePtr)(&v)
}

func (*eventSourcePtr) ElementType() reflect.Type {
	return eventSourcePtrType
}

func (in *eventSourcePtr) ToEventSourcePtrOutput() EventSourcePtrOutput {
	return pulumi.ToOutput(in).(EventSourcePtrOutput)
}

func (in *eventSourcePtr) ToEventSourcePtrOutputWithContext(ctx context.Context) EventSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EventSourcePtrOutput)
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

type Operator string

const (
	// Applies for decimal and non-decimal operands
	OperatorEquals = Operator("Equals")
	// Applies only for decimal operands
	OperatorGreaterThan = Operator("GreaterThan")
	// Applies only for decimal operands
	OperatorGreaterThanOrEqualTo = Operator("GreaterThanOrEqualTo")
	// Applies only for decimal operands
	OperatorLesserThan = Operator("LesserThan")
	// Applies only for decimal operands
	OperatorLesserThanOrEqualTo = Operator("LesserThanOrEqualTo")
	// Applies  for decimal and non-decimal operands
	OperatorNotEquals = Operator("NotEquals")
	// Applies only for non-decimal operands
	OperatorContains = Operator("Contains")
	// Applies only for non-decimal operands
	OperatorStartsWith = Operator("StartsWith")
	// Applies only for non-decimal operands
	OperatorEndsWith = Operator("EndsWith")
)

func (Operator) ElementType() reflect.Type {
	return reflect.TypeOf((*Operator)(nil)).Elem()
}

func (e Operator) ToOperatorOutput() OperatorOutput {
	return pulumi.ToOutput(e).(OperatorOutput)
}

func (e Operator) ToOperatorOutputWithContext(ctx context.Context) OperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatorOutput)
}

func (e Operator) ToOperatorPtrOutput() OperatorPtrOutput {
	return e.ToOperatorPtrOutputWithContext(context.Background())
}

func (e Operator) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return Operator(e).ToOperatorOutputWithContext(ctx).ToOperatorPtrOutputWithContext(ctx)
}

func (e Operator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Operator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Operator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Operator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatorOutput struct{ *pulumi.OutputState }

func (OperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Operator)(nil)).Elem()
}

func (o OperatorOutput) ToOperatorOutput() OperatorOutput {
	return o
}

func (o OperatorOutput) ToOperatorOutputWithContext(ctx context.Context) OperatorOutput {
	return o
}

func (o OperatorOutput) ToOperatorPtrOutput() OperatorPtrOutput {
	return o.ToOperatorPtrOutputWithContext(context.Background())
}

func (o OperatorOutput) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Operator) *Operator {
		return &v
	}).(OperatorPtrOutput)
}

func (o OperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Operator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Operator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatorPtrOutput struct{ *pulumi.OutputState }

func (OperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Operator)(nil)).Elem()
}

func (o OperatorPtrOutput) ToOperatorPtrOutput() OperatorPtrOutput {
	return o
}

func (o OperatorPtrOutput) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return o
}

func (o OperatorPtrOutput) Elem() OperatorOutput {
	return o.ApplyT(func(v *Operator) Operator {
		if v != nil {
			return *v
		}
		var ret Operator
		return ret
	}).(OperatorOutput)
}

func (o OperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Operator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatorInput interface {
	pulumi.Input

	ToOperatorOutput() OperatorOutput
	ToOperatorOutputWithContext(context.Context) OperatorOutput
}

var operatorPtrType = reflect.TypeOf((**Operator)(nil)).Elem()

type OperatorPtrInput interface {
	pulumi.Input

	ToOperatorPtrOutput() OperatorPtrOutput
	ToOperatorPtrOutputWithContext(context.Context) OperatorPtrOutput
}

type operatorPtr string

func OperatorPtr(v string) OperatorPtrInput {
	return (*operatorPtr)(&v)
}

func (*operatorPtr) ElementType() reflect.Type {
	return operatorPtrType
}

func (in *operatorPtr) ToOperatorPtrOutput() OperatorPtrOutput {
	return pulumi.ToOutput(in).(OperatorPtrOutput)
}

func (in *operatorPtr) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatorPtrOutput)
}

type PropertyType string

const (
	PropertyTypeString  = PropertyType("String")
	PropertyTypeInteger = PropertyType("Integer")
	PropertyTypeNumber  = PropertyType("Number")
	PropertyTypeBoolean = PropertyType("Boolean")
)

func (PropertyType) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertyType)(nil)).Elem()
}

func (e PropertyType) ToPropertyTypeOutput() PropertyTypeOutput {
	return pulumi.ToOutput(e).(PropertyTypeOutput)
}

func (e PropertyType) ToPropertyTypeOutputWithContext(ctx context.Context) PropertyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PropertyTypeOutput)
}

func (e PropertyType) ToPropertyTypePtrOutput() PropertyTypePtrOutput {
	return e.ToPropertyTypePtrOutputWithContext(context.Background())
}

func (e PropertyType) ToPropertyTypePtrOutputWithContext(ctx context.Context) PropertyTypePtrOutput {
	return PropertyType(e).ToPropertyTypeOutputWithContext(ctx).ToPropertyTypePtrOutputWithContext(ctx)
}

func (e PropertyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PropertyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PropertyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PropertyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PropertyTypeOutput struct{ *pulumi.OutputState }

func (PropertyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertyType)(nil)).Elem()
}

func (o PropertyTypeOutput) ToPropertyTypeOutput() PropertyTypeOutput {
	return o
}

func (o PropertyTypeOutput) ToPropertyTypeOutputWithContext(ctx context.Context) PropertyTypeOutput {
	return o
}

func (o PropertyTypeOutput) ToPropertyTypePtrOutput() PropertyTypePtrOutput {
	return o.ToPropertyTypePtrOutputWithContext(context.Background())
}

func (o PropertyTypeOutput) ToPropertyTypePtrOutputWithContext(ctx context.Context) PropertyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PropertyType) *PropertyType {
		return &v
	}).(PropertyTypePtrOutput)
}

func (o PropertyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PropertyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PropertyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PropertyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PropertyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PropertyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PropertyTypePtrOutput struct{ *pulumi.OutputState }

func (PropertyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyType)(nil)).Elem()
}

func (o PropertyTypePtrOutput) ToPropertyTypePtrOutput() PropertyTypePtrOutput {
	return o
}

func (o PropertyTypePtrOutput) ToPropertyTypePtrOutputWithContext(ctx context.Context) PropertyTypePtrOutput {
	return o
}

func (o PropertyTypePtrOutput) Elem() PropertyTypeOutput {
	return o.ApplyT(func(v *PropertyType) PropertyType {
		if v != nil {
			return *v
		}
		var ret PropertyType
		return ret
	}).(PropertyTypeOutput)
}

func (o PropertyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PropertyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PropertyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PropertyTypeInput interface {
	pulumi.Input

	ToPropertyTypeOutput() PropertyTypeOutput
	ToPropertyTypeOutputWithContext(context.Context) PropertyTypeOutput
}

var propertyTypePtrType = reflect.TypeOf((**PropertyType)(nil)).Elem()

type PropertyTypePtrInput interface {
	pulumi.Input

	ToPropertyTypePtrOutput() PropertyTypePtrOutput
	ToPropertyTypePtrOutputWithContext(context.Context) PropertyTypePtrOutput
}

type propertyTypePtr string

func PropertyTypePtr(v string) PropertyTypePtrInput {
	return (*propertyTypePtr)(&v)
}

func (*propertyTypePtr) ElementType() reflect.Type {
	return propertyTypePtrType
}

func (in *propertyTypePtr) ToPropertyTypePtrOutput() PropertyTypePtrOutput {
	return pulumi.ToOutput(in).(PropertyTypePtrOutput)
}

func (in *propertyTypePtr) ToPropertyTypePtrOutputWithContext(ctx context.Context) PropertyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PropertyTypePtrOutput)
}

type RuleState string

const (
	RuleStateEnabled  = RuleState("Enabled")
	RuleStateDisabled = RuleState("Disabled")
	RuleStateExpired  = RuleState("Expired")
)

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleState)(nil)).Elem()
}

func (e RuleState) ToRuleStateOutput() RuleStateOutput {
	return pulumi.ToOutput(e).(RuleStateOutput)
}

func (e RuleState) ToRuleStateOutputWithContext(ctx context.Context) RuleStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RuleStateOutput)
}

func (e RuleState) ToRuleStatePtrOutput() RuleStatePtrOutput {
	return e.ToRuleStatePtrOutputWithContext(context.Background())
}

func (e RuleState) ToRuleStatePtrOutputWithContext(ctx context.Context) RuleStatePtrOutput {
	return RuleState(e).ToRuleStateOutputWithContext(ctx).ToRuleStatePtrOutputWithContext(ctx)
}

func (e RuleState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RuleState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RuleState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RuleState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RuleStateOutput struct{ *pulumi.OutputState }

func (RuleStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleState)(nil)).Elem()
}

func (o RuleStateOutput) ToRuleStateOutput() RuleStateOutput {
	return o
}

func (o RuleStateOutput) ToRuleStateOutputWithContext(ctx context.Context) RuleStateOutput {
	return o
}

func (o RuleStateOutput) ToRuleStatePtrOutput() RuleStatePtrOutput {
	return o.ToRuleStatePtrOutputWithContext(context.Background())
}

func (o RuleStateOutput) ToRuleStatePtrOutputWithContext(ctx context.Context) RuleStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RuleState) *RuleState {
		return &v
	}).(RuleStatePtrOutput)
}

func (o RuleStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RuleStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RuleState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RuleStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RuleStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RuleState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RuleStatePtrOutput struct{ *pulumi.OutputState }

func (RuleStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleState)(nil)).Elem()
}

func (o RuleStatePtrOutput) ToRuleStatePtrOutput() RuleStatePtrOutput {
	return o
}

func (o RuleStatePtrOutput) ToRuleStatePtrOutputWithContext(ctx context.Context) RuleStatePtrOutput {
	return o
}

func (o RuleStatePtrOutput) Elem() RuleStateOutput {
	return o.ApplyT(func(v *RuleState) RuleState {
		if v != nil {
			return *v
		}
		var ret RuleState
		return ret
	}).(RuleStateOutput)
}

func (o RuleStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RuleStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RuleState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RuleStateInput interface {
	pulumi.Input

	ToRuleStateOutput() RuleStateOutput
	ToRuleStateOutputWithContext(context.Context) RuleStateOutput
}

var ruleStatePtrType = reflect.TypeOf((**RuleState)(nil)).Elem()

type RuleStatePtrInput interface {
	pulumi.Input

	ToRuleStatePtrOutput() RuleStatePtrOutput
	ToRuleStatePtrOutputWithContext(context.Context) RuleStatePtrOutput
}

type ruleStatePtr string

func RuleStatePtr(v string) RuleStatePtrInput {
	return (*ruleStatePtr)(&v)
}

func (*ruleStatePtr) ElementType() reflect.Type {
	return ruleStatePtrType
}

func (in *ruleStatePtr) ToRuleStatePtrOutput() RuleStatePtrOutput {
	return pulumi.ToOutput(in).(RuleStatePtrOutput)
}

func (in *ruleStatePtr) ToRuleStatePtrOutputWithContext(ctx context.Context) RuleStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RuleStatePtrOutput)
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
	pulumi.RegisterOutputType(ActionTypeOutput{})
	pulumi.RegisterOutputType(ActionTypePtrOutput{})
	pulumi.RegisterOutputType(AssessmentStatusCodeOutput{})
	pulumi.RegisterOutputType(AssessmentStatusCodePtrOutput{})
	pulumi.RegisterOutputType(AssessmentTypeOutput{})
	pulumi.RegisterOutputType(AssessmentTypePtrOutput{})
	pulumi.RegisterOutputType(CategoriesOutput{})
	pulumi.RegisterOutputType(CategoriesPtrOutput{})
	pulumi.RegisterOutputType(EventSourceOutput{})
	pulumi.RegisterOutputType(EventSourcePtrOutput{})
	pulumi.RegisterOutputType(ImplementationEffortOutput{})
	pulumi.RegisterOutputType(ImplementationEffortPtrOutput{})
	pulumi.RegisterOutputType(OperatorOutput{})
	pulumi.RegisterOutputType(OperatorPtrOutput{})
	pulumi.RegisterOutputType(PropertyTypeOutput{})
	pulumi.RegisterOutputType(PropertyTypePtrOutput{})
	pulumi.RegisterOutputType(RuleStateOutput{})
	pulumi.RegisterOutputType(RuleStatePtrOutput{})
	pulumi.RegisterOutputType(SeverityOutput{})
	pulumi.RegisterOutputType(SeverityPtrOutput{})
	pulumi.RegisterOutputType(SourceOutput{})
	pulumi.RegisterOutputType(SourcePtrOutput{})
	pulumi.RegisterOutputType(ThreatsOutput{})
	pulumi.RegisterOutputType(ThreatsPtrOutput{})
	pulumi.RegisterOutputType(UserImpactOutput{})
	pulumi.RegisterOutputType(UserImpactPtrOutput{})
}

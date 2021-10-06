


package v20200804preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HealthAlertAction struct {
	ActionGroupId     *string           `pulumi:"actionGroupId"`
	WebHookProperties map[string]string `pulumi:"webHookProperties"`
}





type HealthAlertActionInput interface {
	pulumi.Input

	ToHealthAlertActionOutput() HealthAlertActionOutput
	ToHealthAlertActionOutputWithContext(context.Context) HealthAlertActionOutput
}

type HealthAlertActionArgs struct {
	ActionGroupId     pulumi.StringPtrInput `pulumi:"actionGroupId"`
	WebHookProperties pulumi.StringMapInput `pulumi:"webHookProperties"`
}

func (HealthAlertActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertAction)(nil)).Elem()
}

func (i HealthAlertActionArgs) ToHealthAlertActionOutput() HealthAlertActionOutput {
	return i.ToHealthAlertActionOutputWithContext(context.Background())
}

func (i HealthAlertActionArgs) ToHealthAlertActionOutputWithContext(ctx context.Context) HealthAlertActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertActionOutput)
}





type HealthAlertActionArrayInput interface {
	pulumi.Input

	ToHealthAlertActionArrayOutput() HealthAlertActionArrayOutput
	ToHealthAlertActionArrayOutputWithContext(context.Context) HealthAlertActionArrayOutput
}

type HealthAlertActionArray []HealthAlertActionInput

func (HealthAlertActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertAction)(nil)).Elem()
}

func (i HealthAlertActionArray) ToHealthAlertActionArrayOutput() HealthAlertActionArrayOutput {
	return i.ToHealthAlertActionArrayOutputWithContext(context.Background())
}

func (i HealthAlertActionArray) ToHealthAlertActionArrayOutputWithContext(ctx context.Context) HealthAlertActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertActionArrayOutput)
}

type HealthAlertActionOutput struct{ *pulumi.OutputState }

func (HealthAlertActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertAction)(nil)).Elem()
}

func (o HealthAlertActionOutput) ToHealthAlertActionOutput() HealthAlertActionOutput {
	return o
}

func (o HealthAlertActionOutput) ToHealthAlertActionOutputWithContext(ctx context.Context) HealthAlertActionOutput {
	return o
}

func (o HealthAlertActionOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthAlertAction) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

func (o HealthAlertActionOutput) WebHookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v HealthAlertAction) map[string]string { return v.WebHookProperties }).(pulumi.StringMapOutput)
}

type HealthAlertActionArrayOutput struct{ *pulumi.OutputState }

func (HealthAlertActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertAction)(nil)).Elem()
}

func (o HealthAlertActionArrayOutput) ToHealthAlertActionArrayOutput() HealthAlertActionArrayOutput {
	return o
}

func (o HealthAlertActionArrayOutput) ToHealthAlertActionArrayOutputWithContext(ctx context.Context) HealthAlertActionArrayOutput {
	return o
}

func (o HealthAlertActionArrayOutput) Index(i pulumi.IntInput) HealthAlertActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthAlertAction {
		return vs[0].([]HealthAlertAction)[vs[1].(int)]
	}).(HealthAlertActionOutput)
}

type HealthAlertActionResponse struct {
	ActionGroupId     *string           `pulumi:"actionGroupId"`
	WebHookProperties map[string]string `pulumi:"webHookProperties"`
}





type HealthAlertActionResponseInput interface {
	pulumi.Input

	ToHealthAlertActionResponseOutput() HealthAlertActionResponseOutput
	ToHealthAlertActionResponseOutputWithContext(context.Context) HealthAlertActionResponseOutput
}

type HealthAlertActionResponseArgs struct {
	ActionGroupId     pulumi.StringPtrInput `pulumi:"actionGroupId"`
	WebHookProperties pulumi.StringMapInput `pulumi:"webHookProperties"`
}

func (HealthAlertActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertActionResponse)(nil)).Elem()
}

func (i HealthAlertActionResponseArgs) ToHealthAlertActionResponseOutput() HealthAlertActionResponseOutput {
	return i.ToHealthAlertActionResponseOutputWithContext(context.Background())
}

func (i HealthAlertActionResponseArgs) ToHealthAlertActionResponseOutputWithContext(ctx context.Context) HealthAlertActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertActionResponseOutput)
}





type HealthAlertActionResponseArrayInput interface {
	pulumi.Input

	ToHealthAlertActionResponseArrayOutput() HealthAlertActionResponseArrayOutput
	ToHealthAlertActionResponseArrayOutputWithContext(context.Context) HealthAlertActionResponseArrayOutput
}

type HealthAlertActionResponseArray []HealthAlertActionResponseInput

func (HealthAlertActionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertActionResponse)(nil)).Elem()
}

func (i HealthAlertActionResponseArray) ToHealthAlertActionResponseArrayOutput() HealthAlertActionResponseArrayOutput {
	return i.ToHealthAlertActionResponseArrayOutputWithContext(context.Background())
}

func (i HealthAlertActionResponseArray) ToHealthAlertActionResponseArrayOutputWithContext(ctx context.Context) HealthAlertActionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertActionResponseArrayOutput)
}

type HealthAlertActionResponseOutput struct{ *pulumi.OutputState }

func (HealthAlertActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertActionResponse)(nil)).Elem()
}

func (o HealthAlertActionResponseOutput) ToHealthAlertActionResponseOutput() HealthAlertActionResponseOutput {
	return o
}

func (o HealthAlertActionResponseOutput) ToHealthAlertActionResponseOutputWithContext(ctx context.Context) HealthAlertActionResponseOutput {
	return o
}

func (o HealthAlertActionResponseOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthAlertActionResponse) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

func (o HealthAlertActionResponseOutput) WebHookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v HealthAlertActionResponse) map[string]string { return v.WebHookProperties }).(pulumi.StringMapOutput)
}

type HealthAlertActionResponseArrayOutput struct{ *pulumi.OutputState }

func (HealthAlertActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertActionResponse)(nil)).Elem()
}

func (o HealthAlertActionResponseArrayOutput) ToHealthAlertActionResponseArrayOutput() HealthAlertActionResponseArrayOutput {
	return o
}

func (o HealthAlertActionResponseArrayOutput) ToHealthAlertActionResponseArrayOutputWithContext(ctx context.Context) HealthAlertActionResponseArrayOutput {
	return o
}

func (o HealthAlertActionResponseArrayOutput) Index(i pulumi.IntInput) HealthAlertActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthAlertActionResponse {
		return vs[0].([]HealthAlertActionResponse)[vs[1].(int)]
	}).(HealthAlertActionResponseOutput)
}

type HealthAlertCriteria struct {
	AllOf []VmGuestHealthAlertCriterion `pulumi:"allOf"`
}





type HealthAlertCriteriaInput interface {
	pulumi.Input

	ToHealthAlertCriteriaOutput() HealthAlertCriteriaOutput
	ToHealthAlertCriteriaOutputWithContext(context.Context) HealthAlertCriteriaOutput
}

type HealthAlertCriteriaArgs struct {
	AllOf VmGuestHealthAlertCriterionArrayInput `pulumi:"allOf"`
}

func (HealthAlertCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertCriteria)(nil)).Elem()
}

func (i HealthAlertCriteriaArgs) ToHealthAlertCriteriaOutput() HealthAlertCriteriaOutput {
	return i.ToHealthAlertCriteriaOutputWithContext(context.Background())
}

func (i HealthAlertCriteriaArgs) ToHealthAlertCriteriaOutputWithContext(ctx context.Context) HealthAlertCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaOutput)
}

func (i HealthAlertCriteriaArgs) ToHealthAlertCriteriaPtrOutput() HealthAlertCriteriaPtrOutput {
	return i.ToHealthAlertCriteriaPtrOutputWithContext(context.Background())
}

func (i HealthAlertCriteriaArgs) ToHealthAlertCriteriaPtrOutputWithContext(ctx context.Context) HealthAlertCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaOutput).ToHealthAlertCriteriaPtrOutputWithContext(ctx)
}









type HealthAlertCriteriaPtrInput interface {
	pulumi.Input

	ToHealthAlertCriteriaPtrOutput() HealthAlertCriteriaPtrOutput
	ToHealthAlertCriteriaPtrOutputWithContext(context.Context) HealthAlertCriteriaPtrOutput
}

type healthAlertCriteriaPtrType HealthAlertCriteriaArgs

func HealthAlertCriteriaPtr(v *HealthAlertCriteriaArgs) HealthAlertCriteriaPtrInput {
	return (*healthAlertCriteriaPtrType)(v)
}

func (*healthAlertCriteriaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthAlertCriteria)(nil)).Elem()
}

func (i *healthAlertCriteriaPtrType) ToHealthAlertCriteriaPtrOutput() HealthAlertCriteriaPtrOutput {
	return i.ToHealthAlertCriteriaPtrOutputWithContext(context.Background())
}

func (i *healthAlertCriteriaPtrType) ToHealthAlertCriteriaPtrOutputWithContext(ctx context.Context) HealthAlertCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaPtrOutput)
}

type HealthAlertCriteriaOutput struct{ *pulumi.OutputState }

func (HealthAlertCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertCriteria)(nil)).Elem()
}

func (o HealthAlertCriteriaOutput) ToHealthAlertCriteriaOutput() HealthAlertCriteriaOutput {
	return o
}

func (o HealthAlertCriteriaOutput) ToHealthAlertCriteriaOutputWithContext(ctx context.Context) HealthAlertCriteriaOutput {
	return o
}

func (o HealthAlertCriteriaOutput) ToHealthAlertCriteriaPtrOutput() HealthAlertCriteriaPtrOutput {
	return o.ToHealthAlertCriteriaPtrOutputWithContext(context.Background())
}

func (o HealthAlertCriteriaOutput) ToHealthAlertCriteriaPtrOutputWithContext(ctx context.Context) HealthAlertCriteriaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HealthAlertCriteria) *HealthAlertCriteria {
		return &v
	}).(HealthAlertCriteriaPtrOutput)
}

func (o HealthAlertCriteriaOutput) AllOf() VmGuestHealthAlertCriterionArrayOutput {
	return o.ApplyT(func(v HealthAlertCriteria) []VmGuestHealthAlertCriterion { return v.AllOf }).(VmGuestHealthAlertCriterionArrayOutput)
}

type HealthAlertCriteriaPtrOutput struct{ *pulumi.OutputState }

func (HealthAlertCriteriaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthAlertCriteria)(nil)).Elem()
}

func (o HealthAlertCriteriaPtrOutput) ToHealthAlertCriteriaPtrOutput() HealthAlertCriteriaPtrOutput {
	return o
}

func (o HealthAlertCriteriaPtrOutput) ToHealthAlertCriteriaPtrOutputWithContext(ctx context.Context) HealthAlertCriteriaPtrOutput {
	return o
}

func (o HealthAlertCriteriaPtrOutput) Elem() HealthAlertCriteriaOutput {
	return o.ApplyT(func(v *HealthAlertCriteria) HealthAlertCriteria {
		if v != nil {
			return *v
		}
		var ret HealthAlertCriteria
		return ret
	}).(HealthAlertCriteriaOutput)
}

func (o HealthAlertCriteriaPtrOutput) AllOf() VmGuestHealthAlertCriterionArrayOutput {
	return o.ApplyT(func(v *HealthAlertCriteria) []VmGuestHealthAlertCriterion {
		if v == nil {
			return nil
		}
		return v.AllOf
	}).(VmGuestHealthAlertCriterionArrayOutput)
}

type HealthAlertCriteriaResponse struct {
	AllOf []VmGuestHealthAlertCriterionResponse `pulumi:"allOf"`
}





type HealthAlertCriteriaResponseInput interface {
	pulumi.Input

	ToHealthAlertCriteriaResponseOutput() HealthAlertCriteriaResponseOutput
	ToHealthAlertCriteriaResponseOutputWithContext(context.Context) HealthAlertCriteriaResponseOutput
}

type HealthAlertCriteriaResponseArgs struct {
	AllOf VmGuestHealthAlertCriterionResponseArrayInput `pulumi:"allOf"`
}

func (HealthAlertCriteriaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertCriteriaResponse)(nil)).Elem()
}

func (i HealthAlertCriteriaResponseArgs) ToHealthAlertCriteriaResponseOutput() HealthAlertCriteriaResponseOutput {
	return i.ToHealthAlertCriteriaResponseOutputWithContext(context.Background())
}

func (i HealthAlertCriteriaResponseArgs) ToHealthAlertCriteriaResponseOutputWithContext(ctx context.Context) HealthAlertCriteriaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaResponseOutput)
}

func (i HealthAlertCriteriaResponseArgs) ToHealthAlertCriteriaResponsePtrOutput() HealthAlertCriteriaResponsePtrOutput {
	return i.ToHealthAlertCriteriaResponsePtrOutputWithContext(context.Background())
}

func (i HealthAlertCriteriaResponseArgs) ToHealthAlertCriteriaResponsePtrOutputWithContext(ctx context.Context) HealthAlertCriteriaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaResponseOutput).ToHealthAlertCriteriaResponsePtrOutputWithContext(ctx)
}









type HealthAlertCriteriaResponsePtrInput interface {
	pulumi.Input

	ToHealthAlertCriteriaResponsePtrOutput() HealthAlertCriteriaResponsePtrOutput
	ToHealthAlertCriteriaResponsePtrOutputWithContext(context.Context) HealthAlertCriteriaResponsePtrOutput
}

type healthAlertCriteriaResponsePtrType HealthAlertCriteriaResponseArgs

func HealthAlertCriteriaResponsePtr(v *HealthAlertCriteriaResponseArgs) HealthAlertCriteriaResponsePtrInput {
	return (*healthAlertCriteriaResponsePtrType)(v)
}

func (*healthAlertCriteriaResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthAlertCriteriaResponse)(nil)).Elem()
}

func (i *healthAlertCriteriaResponsePtrType) ToHealthAlertCriteriaResponsePtrOutput() HealthAlertCriteriaResponsePtrOutput {
	return i.ToHealthAlertCriteriaResponsePtrOutputWithContext(context.Background())
}

func (i *healthAlertCriteriaResponsePtrType) ToHealthAlertCriteriaResponsePtrOutputWithContext(ctx context.Context) HealthAlertCriteriaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaResponsePtrOutput)
}

type HealthAlertCriteriaResponseOutput struct{ *pulumi.OutputState }

func (HealthAlertCriteriaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertCriteriaResponse)(nil)).Elem()
}

func (o HealthAlertCriteriaResponseOutput) ToHealthAlertCriteriaResponseOutput() HealthAlertCriteriaResponseOutput {
	return o
}

func (o HealthAlertCriteriaResponseOutput) ToHealthAlertCriteriaResponseOutputWithContext(ctx context.Context) HealthAlertCriteriaResponseOutput {
	return o
}

func (o HealthAlertCriteriaResponseOutput) ToHealthAlertCriteriaResponsePtrOutput() HealthAlertCriteriaResponsePtrOutput {
	return o.ToHealthAlertCriteriaResponsePtrOutputWithContext(context.Background())
}

func (o HealthAlertCriteriaResponseOutput) ToHealthAlertCriteriaResponsePtrOutputWithContext(ctx context.Context) HealthAlertCriteriaResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HealthAlertCriteriaResponse) *HealthAlertCriteriaResponse {
		return &v
	}).(HealthAlertCriteriaResponsePtrOutput)
}

func (o HealthAlertCriteriaResponseOutput) AllOf() VmGuestHealthAlertCriterionResponseArrayOutput {
	return o.ApplyT(func(v HealthAlertCriteriaResponse) []VmGuestHealthAlertCriterionResponse { return v.AllOf }).(VmGuestHealthAlertCriterionResponseArrayOutput)
}

type HealthAlertCriteriaResponsePtrOutput struct{ *pulumi.OutputState }

func (HealthAlertCriteriaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthAlertCriteriaResponse)(nil)).Elem()
}

func (o HealthAlertCriteriaResponsePtrOutput) ToHealthAlertCriteriaResponsePtrOutput() HealthAlertCriteriaResponsePtrOutput {
	return o
}

func (o HealthAlertCriteriaResponsePtrOutput) ToHealthAlertCriteriaResponsePtrOutputWithContext(ctx context.Context) HealthAlertCriteriaResponsePtrOutput {
	return o
}

func (o HealthAlertCriteriaResponsePtrOutput) Elem() HealthAlertCriteriaResponseOutput {
	return o.ApplyT(func(v *HealthAlertCriteriaResponse) HealthAlertCriteriaResponse {
		if v != nil {
			return *v
		}
		var ret HealthAlertCriteriaResponse
		return ret
	}).(HealthAlertCriteriaResponseOutput)
}

func (o HealthAlertCriteriaResponsePtrOutput) AllOf() VmGuestHealthAlertCriterionResponseArrayOutput {
	return o.ApplyT(func(v *HealthAlertCriteriaResponse) []VmGuestHealthAlertCriterionResponse {
		if v == nil {
			return nil
		}
		return v.AllOf
	}).(VmGuestHealthAlertCriterionResponseArrayOutput)
}

type HealthState struct {
	HealthStateName string  `pulumi:"healthStateName"`
	Severity        float64 `pulumi:"severity"`
}





type HealthStateInput interface {
	pulumi.Input

	ToHealthStateOutput() HealthStateOutput
	ToHealthStateOutputWithContext(context.Context) HealthStateOutput
}

type HealthStateArgs struct {
	HealthStateName pulumi.StringInput  `pulumi:"healthStateName"`
	Severity        pulumi.Float64Input `pulumi:"severity"`
}

func (HealthStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthState)(nil)).Elem()
}

func (i HealthStateArgs) ToHealthStateOutput() HealthStateOutput {
	return i.ToHealthStateOutputWithContext(context.Background())
}

func (i HealthStateArgs) ToHealthStateOutputWithContext(ctx context.Context) HealthStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthStateOutput)
}





type HealthStateArrayInput interface {
	pulumi.Input

	ToHealthStateArrayOutput() HealthStateArrayOutput
	ToHealthStateArrayOutputWithContext(context.Context) HealthStateArrayOutput
}

type HealthStateArray []HealthStateInput

func (HealthStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthState)(nil)).Elem()
}

func (i HealthStateArray) ToHealthStateArrayOutput() HealthStateArrayOutput {
	return i.ToHealthStateArrayOutputWithContext(context.Background())
}

func (i HealthStateArray) ToHealthStateArrayOutputWithContext(ctx context.Context) HealthStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthStateArrayOutput)
}

type HealthStateOutput struct{ *pulumi.OutputState }

func (HealthStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthState)(nil)).Elem()
}

func (o HealthStateOutput) ToHealthStateOutput() HealthStateOutput {
	return o
}

func (o HealthStateOutput) ToHealthStateOutputWithContext(ctx context.Context) HealthStateOutput {
	return o
}

func (o HealthStateOutput) HealthStateName() pulumi.StringOutput {
	return o.ApplyT(func(v HealthState) string { return v.HealthStateName }).(pulumi.StringOutput)
}

func (o HealthStateOutput) Severity() pulumi.Float64Output {
	return o.ApplyT(func(v HealthState) float64 { return v.Severity }).(pulumi.Float64Output)
}

type HealthStateArrayOutput struct{ *pulumi.OutputState }

func (HealthStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthState)(nil)).Elem()
}

func (o HealthStateArrayOutput) ToHealthStateArrayOutput() HealthStateArrayOutput {
	return o
}

func (o HealthStateArrayOutput) ToHealthStateArrayOutputWithContext(ctx context.Context) HealthStateArrayOutput {
	return o
}

func (o HealthStateArrayOutput) Index(i pulumi.IntInput) HealthStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthState {
		return vs[0].([]HealthState)[vs[1].(int)]
	}).(HealthStateOutput)
}

type HealthStateResponse struct {
	HealthStateName string  `pulumi:"healthStateName"`
	Severity        float64 `pulumi:"severity"`
}





type HealthStateResponseInput interface {
	pulumi.Input

	ToHealthStateResponseOutput() HealthStateResponseOutput
	ToHealthStateResponseOutputWithContext(context.Context) HealthStateResponseOutput
}

type HealthStateResponseArgs struct {
	HealthStateName pulumi.StringInput  `pulumi:"healthStateName"`
	Severity        pulumi.Float64Input `pulumi:"severity"`
}

func (HealthStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthStateResponse)(nil)).Elem()
}

func (i HealthStateResponseArgs) ToHealthStateResponseOutput() HealthStateResponseOutput {
	return i.ToHealthStateResponseOutputWithContext(context.Background())
}

func (i HealthStateResponseArgs) ToHealthStateResponseOutputWithContext(ctx context.Context) HealthStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthStateResponseOutput)
}





type HealthStateResponseArrayInput interface {
	pulumi.Input

	ToHealthStateResponseArrayOutput() HealthStateResponseArrayOutput
	ToHealthStateResponseArrayOutputWithContext(context.Context) HealthStateResponseArrayOutput
}

type HealthStateResponseArray []HealthStateResponseInput

func (HealthStateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthStateResponse)(nil)).Elem()
}

func (i HealthStateResponseArray) ToHealthStateResponseArrayOutput() HealthStateResponseArrayOutput {
	return i.ToHealthStateResponseArrayOutputWithContext(context.Background())
}

func (i HealthStateResponseArray) ToHealthStateResponseArrayOutputWithContext(ctx context.Context) HealthStateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthStateResponseArrayOutput)
}

type HealthStateResponseOutput struct{ *pulumi.OutputState }

func (HealthStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthStateResponse)(nil)).Elem()
}

func (o HealthStateResponseOutput) ToHealthStateResponseOutput() HealthStateResponseOutput {
	return o
}

func (o HealthStateResponseOutput) ToHealthStateResponseOutputWithContext(ctx context.Context) HealthStateResponseOutput {
	return o
}

func (o HealthStateResponseOutput) HealthStateName() pulumi.StringOutput {
	return o.ApplyT(func(v HealthStateResponse) string { return v.HealthStateName }).(pulumi.StringOutput)
}

func (o HealthStateResponseOutput) Severity() pulumi.Float64Output {
	return o.ApplyT(func(v HealthStateResponse) float64 { return v.Severity }).(pulumi.Float64Output)
}

type HealthStateResponseArrayOutput struct{ *pulumi.OutputState }

func (HealthStateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthStateResponse)(nil)).Elem()
}

func (o HealthStateResponseArrayOutput) ToHealthStateResponseArrayOutput() HealthStateResponseArrayOutput {
	return o
}

func (o HealthStateResponseArrayOutput) ToHealthStateResponseArrayOutputWithContext(ctx context.Context) HealthStateResponseArrayOutput {
	return o
}

func (o HealthStateResponseArrayOutput) Index(i pulumi.IntInput) HealthStateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthStateResponse {
		return vs[0].([]HealthStateResponse)[vs[1].(int)]
	}).(HealthStateResponseOutput)
}

type VmGuestHealthAlertCriterion struct {
	HealthStates []HealthState `pulumi:"healthStates"`
	MonitorNames []string      `pulumi:"monitorNames"`
	MonitorTypes []string      `pulumi:"monitorTypes"`
	Namespace    string        `pulumi:"namespace"`
}





type VmGuestHealthAlertCriterionInput interface {
	pulumi.Input

	ToVmGuestHealthAlertCriterionOutput() VmGuestHealthAlertCriterionOutput
	ToVmGuestHealthAlertCriterionOutputWithContext(context.Context) VmGuestHealthAlertCriterionOutput
}

type VmGuestHealthAlertCriterionArgs struct {
	HealthStates HealthStateArrayInput   `pulumi:"healthStates"`
	MonitorNames pulumi.StringArrayInput `pulumi:"monitorNames"`
	MonitorTypes pulumi.StringArrayInput `pulumi:"monitorTypes"`
	Namespace    pulumi.StringInput      `pulumi:"namespace"`
}

func (VmGuestHealthAlertCriterionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmGuestHealthAlertCriterion)(nil)).Elem()
}

func (i VmGuestHealthAlertCriterionArgs) ToVmGuestHealthAlertCriterionOutput() VmGuestHealthAlertCriterionOutput {
	return i.ToVmGuestHealthAlertCriterionOutputWithContext(context.Background())
}

func (i VmGuestHealthAlertCriterionArgs) ToVmGuestHealthAlertCriterionOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmGuestHealthAlertCriterionOutput)
}





type VmGuestHealthAlertCriterionArrayInput interface {
	pulumi.Input

	ToVmGuestHealthAlertCriterionArrayOutput() VmGuestHealthAlertCriterionArrayOutput
	ToVmGuestHealthAlertCriterionArrayOutputWithContext(context.Context) VmGuestHealthAlertCriterionArrayOutput
}

type VmGuestHealthAlertCriterionArray []VmGuestHealthAlertCriterionInput

func (VmGuestHealthAlertCriterionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmGuestHealthAlertCriterion)(nil)).Elem()
}

func (i VmGuestHealthAlertCriterionArray) ToVmGuestHealthAlertCriterionArrayOutput() VmGuestHealthAlertCriterionArrayOutput {
	return i.ToVmGuestHealthAlertCriterionArrayOutputWithContext(context.Background())
}

func (i VmGuestHealthAlertCriterionArray) ToVmGuestHealthAlertCriterionArrayOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmGuestHealthAlertCriterionArrayOutput)
}

type VmGuestHealthAlertCriterionOutput struct{ *pulumi.OutputState }

func (VmGuestHealthAlertCriterionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmGuestHealthAlertCriterion)(nil)).Elem()
}

func (o VmGuestHealthAlertCriterionOutput) ToVmGuestHealthAlertCriterionOutput() VmGuestHealthAlertCriterionOutput {
	return o
}

func (o VmGuestHealthAlertCriterionOutput) ToVmGuestHealthAlertCriterionOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionOutput {
	return o
}

func (o VmGuestHealthAlertCriterionOutput) HealthStates() HealthStateArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterion) []HealthState { return v.HealthStates }).(HealthStateArrayOutput)
}

func (o VmGuestHealthAlertCriterionOutput) MonitorNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterion) []string { return v.MonitorNames }).(pulumi.StringArrayOutput)
}

func (o VmGuestHealthAlertCriterionOutput) MonitorTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterion) []string { return v.MonitorTypes }).(pulumi.StringArrayOutput)
}

func (o VmGuestHealthAlertCriterionOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterion) string { return v.Namespace }).(pulumi.StringOutput)
}

type VmGuestHealthAlertCriterionArrayOutput struct{ *pulumi.OutputState }

func (VmGuestHealthAlertCriterionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmGuestHealthAlertCriterion)(nil)).Elem()
}

func (o VmGuestHealthAlertCriterionArrayOutput) ToVmGuestHealthAlertCriterionArrayOutput() VmGuestHealthAlertCriterionArrayOutput {
	return o
}

func (o VmGuestHealthAlertCriterionArrayOutput) ToVmGuestHealthAlertCriterionArrayOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionArrayOutput {
	return o
}

func (o VmGuestHealthAlertCriterionArrayOutput) Index(i pulumi.IntInput) VmGuestHealthAlertCriterionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmGuestHealthAlertCriterion {
		return vs[0].([]VmGuestHealthAlertCriterion)[vs[1].(int)]
	}).(VmGuestHealthAlertCriterionOutput)
}

type VmGuestHealthAlertCriterionResponse struct {
	HealthStates []HealthStateResponse `pulumi:"healthStates"`
	MonitorNames []string              `pulumi:"monitorNames"`
	MonitorTypes []string              `pulumi:"monitorTypes"`
	Namespace    string                `pulumi:"namespace"`
}





type VmGuestHealthAlertCriterionResponseInput interface {
	pulumi.Input

	ToVmGuestHealthAlertCriterionResponseOutput() VmGuestHealthAlertCriterionResponseOutput
	ToVmGuestHealthAlertCriterionResponseOutputWithContext(context.Context) VmGuestHealthAlertCriterionResponseOutput
}

type VmGuestHealthAlertCriterionResponseArgs struct {
	HealthStates HealthStateResponseArrayInput `pulumi:"healthStates"`
	MonitorNames pulumi.StringArrayInput       `pulumi:"monitorNames"`
	MonitorTypes pulumi.StringArrayInput       `pulumi:"monitorTypes"`
	Namespace    pulumi.StringInput            `pulumi:"namespace"`
}

func (VmGuestHealthAlertCriterionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmGuestHealthAlertCriterionResponse)(nil)).Elem()
}

func (i VmGuestHealthAlertCriterionResponseArgs) ToVmGuestHealthAlertCriterionResponseOutput() VmGuestHealthAlertCriterionResponseOutput {
	return i.ToVmGuestHealthAlertCriterionResponseOutputWithContext(context.Background())
}

func (i VmGuestHealthAlertCriterionResponseArgs) ToVmGuestHealthAlertCriterionResponseOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmGuestHealthAlertCriterionResponseOutput)
}





type VmGuestHealthAlertCriterionResponseArrayInput interface {
	pulumi.Input

	ToVmGuestHealthAlertCriterionResponseArrayOutput() VmGuestHealthAlertCriterionResponseArrayOutput
	ToVmGuestHealthAlertCriterionResponseArrayOutputWithContext(context.Context) VmGuestHealthAlertCriterionResponseArrayOutput
}

type VmGuestHealthAlertCriterionResponseArray []VmGuestHealthAlertCriterionResponseInput

func (VmGuestHealthAlertCriterionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmGuestHealthAlertCriterionResponse)(nil)).Elem()
}

func (i VmGuestHealthAlertCriterionResponseArray) ToVmGuestHealthAlertCriterionResponseArrayOutput() VmGuestHealthAlertCriterionResponseArrayOutput {
	return i.ToVmGuestHealthAlertCriterionResponseArrayOutputWithContext(context.Background())
}

func (i VmGuestHealthAlertCriterionResponseArray) ToVmGuestHealthAlertCriterionResponseArrayOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmGuestHealthAlertCriterionResponseArrayOutput)
}

type VmGuestHealthAlertCriterionResponseOutput struct{ *pulumi.OutputState }

func (VmGuestHealthAlertCriterionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmGuestHealthAlertCriterionResponse)(nil)).Elem()
}

func (o VmGuestHealthAlertCriterionResponseOutput) ToVmGuestHealthAlertCriterionResponseOutput() VmGuestHealthAlertCriterionResponseOutput {
	return o
}

func (o VmGuestHealthAlertCriterionResponseOutput) ToVmGuestHealthAlertCriterionResponseOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionResponseOutput {
	return o
}

func (o VmGuestHealthAlertCriterionResponseOutput) HealthStates() HealthStateResponseArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterionResponse) []HealthStateResponse { return v.HealthStates }).(HealthStateResponseArrayOutput)
}

func (o VmGuestHealthAlertCriterionResponseOutput) MonitorNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterionResponse) []string { return v.MonitorNames }).(pulumi.StringArrayOutput)
}

func (o VmGuestHealthAlertCriterionResponseOutput) MonitorTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterionResponse) []string { return v.MonitorTypes }).(pulumi.StringArrayOutput)
}

func (o VmGuestHealthAlertCriterionResponseOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterionResponse) string { return v.Namespace }).(pulumi.StringOutput)
}

type VmGuestHealthAlertCriterionResponseArrayOutput struct{ *pulumi.OutputState }

func (VmGuestHealthAlertCriterionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmGuestHealthAlertCriterionResponse)(nil)).Elem()
}

func (o VmGuestHealthAlertCriterionResponseArrayOutput) ToVmGuestHealthAlertCriterionResponseArrayOutput() VmGuestHealthAlertCriterionResponseArrayOutput {
	return o
}

func (o VmGuestHealthAlertCriterionResponseArrayOutput) ToVmGuestHealthAlertCriterionResponseArrayOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionResponseArrayOutput {
	return o
}

func (o VmGuestHealthAlertCriterionResponseArrayOutput) Index(i pulumi.IntInput) VmGuestHealthAlertCriterionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmGuestHealthAlertCriterionResponse {
		return vs[0].([]VmGuestHealthAlertCriterionResponse)[vs[1].(int)]
	}).(VmGuestHealthAlertCriterionResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(HealthAlertActionOutput{})
	pulumi.RegisterOutputType(HealthAlertActionArrayOutput{})
	pulumi.RegisterOutputType(HealthAlertActionResponseOutput{})
	pulumi.RegisterOutputType(HealthAlertActionResponseArrayOutput{})
	pulumi.RegisterOutputType(HealthAlertCriteriaOutput{})
	pulumi.RegisterOutputType(HealthAlertCriteriaPtrOutput{})
	pulumi.RegisterOutputType(HealthAlertCriteriaResponseOutput{})
	pulumi.RegisterOutputType(HealthAlertCriteriaResponsePtrOutput{})
	pulumi.RegisterOutputType(HealthStateOutput{})
	pulumi.RegisterOutputType(HealthStateArrayOutput{})
	pulumi.RegisterOutputType(HealthStateResponseOutput{})
	pulumi.RegisterOutputType(HealthStateResponseArrayOutput{})
	pulumi.RegisterOutputType(VmGuestHealthAlertCriterionOutput{})
	pulumi.RegisterOutputType(VmGuestHealthAlertCriterionArrayOutput{})
	pulumi.RegisterOutputType(VmGuestHealthAlertCriterionResponseOutput{})
	pulumi.RegisterOutputType(VmGuestHealthAlertCriterionResponseArrayOutput{})
}

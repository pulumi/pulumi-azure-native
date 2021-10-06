


package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LocationThresholdRuleCondition struct {
	DataSource          interface{} `pulumi:"dataSource"`
	FailedLocationCount int         `pulumi:"failedLocationCount"`
	OdataType           string      `pulumi:"odataType"`
	WindowSize          *string     `pulumi:"windowSize"`
}





type LocationThresholdRuleConditionInput interface {
	pulumi.Input

	ToLocationThresholdRuleConditionOutput() LocationThresholdRuleConditionOutput
	ToLocationThresholdRuleConditionOutputWithContext(context.Context) LocationThresholdRuleConditionOutput
}

type LocationThresholdRuleConditionArgs struct {
	DataSource          pulumi.Input          `pulumi:"dataSource"`
	FailedLocationCount pulumi.IntInput       `pulumi:"failedLocationCount"`
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
	WindowSize          pulumi.StringPtrInput `pulumi:"windowSize"`
}

func (LocationThresholdRuleConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationThresholdRuleCondition)(nil)).Elem()
}

func (i LocationThresholdRuleConditionArgs) ToLocationThresholdRuleConditionOutput() LocationThresholdRuleConditionOutput {
	return i.ToLocationThresholdRuleConditionOutputWithContext(context.Background())
}

func (i LocationThresholdRuleConditionArgs) ToLocationThresholdRuleConditionOutputWithContext(ctx context.Context) LocationThresholdRuleConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationThresholdRuleConditionOutput)
}

type LocationThresholdRuleConditionOutput struct{ *pulumi.OutputState }

func (LocationThresholdRuleConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationThresholdRuleCondition)(nil)).Elem()
}

func (o LocationThresholdRuleConditionOutput) ToLocationThresholdRuleConditionOutput() LocationThresholdRuleConditionOutput {
	return o
}

func (o LocationThresholdRuleConditionOutput) ToLocationThresholdRuleConditionOutputWithContext(ctx context.Context) LocationThresholdRuleConditionOutput {
	return o
}

func (o LocationThresholdRuleConditionOutput) DataSource() pulumi.AnyOutput {
	return o.ApplyT(func(v LocationThresholdRuleCondition) interface{} { return v.DataSource }).(pulumi.AnyOutput)
}

func (o LocationThresholdRuleConditionOutput) FailedLocationCount() pulumi.IntOutput {
	return o.ApplyT(func(v LocationThresholdRuleCondition) int { return v.FailedLocationCount }).(pulumi.IntOutput)
}

func (o LocationThresholdRuleConditionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v LocationThresholdRuleCondition) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o LocationThresholdRuleConditionOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationThresholdRuleCondition) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
}

type LocationThresholdRuleConditionResponse struct {
	DataSource          interface{} `pulumi:"dataSource"`
	FailedLocationCount int         `pulumi:"failedLocationCount"`
	OdataType           string      `pulumi:"odataType"`
	WindowSize          *string     `pulumi:"windowSize"`
}





type LocationThresholdRuleConditionResponseInput interface {
	pulumi.Input

	ToLocationThresholdRuleConditionResponseOutput() LocationThresholdRuleConditionResponseOutput
	ToLocationThresholdRuleConditionResponseOutputWithContext(context.Context) LocationThresholdRuleConditionResponseOutput
}

type LocationThresholdRuleConditionResponseArgs struct {
	DataSource          pulumi.Input          `pulumi:"dataSource"`
	FailedLocationCount pulumi.IntInput       `pulumi:"failedLocationCount"`
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
	WindowSize          pulumi.StringPtrInput `pulumi:"windowSize"`
}

func (LocationThresholdRuleConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationThresholdRuleConditionResponse)(nil)).Elem()
}

func (i LocationThresholdRuleConditionResponseArgs) ToLocationThresholdRuleConditionResponseOutput() LocationThresholdRuleConditionResponseOutput {
	return i.ToLocationThresholdRuleConditionResponseOutputWithContext(context.Background())
}

func (i LocationThresholdRuleConditionResponseArgs) ToLocationThresholdRuleConditionResponseOutputWithContext(ctx context.Context) LocationThresholdRuleConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationThresholdRuleConditionResponseOutput)
}

type LocationThresholdRuleConditionResponseOutput struct{ *pulumi.OutputState }

func (LocationThresholdRuleConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationThresholdRuleConditionResponse)(nil)).Elem()
}

func (o LocationThresholdRuleConditionResponseOutput) ToLocationThresholdRuleConditionResponseOutput() LocationThresholdRuleConditionResponseOutput {
	return o
}

func (o LocationThresholdRuleConditionResponseOutput) ToLocationThresholdRuleConditionResponseOutputWithContext(ctx context.Context) LocationThresholdRuleConditionResponseOutput {
	return o
}

func (o LocationThresholdRuleConditionResponseOutput) DataSource() pulumi.AnyOutput {
	return o.ApplyT(func(v LocationThresholdRuleConditionResponse) interface{} { return v.DataSource }).(pulumi.AnyOutput)
}

func (o LocationThresholdRuleConditionResponseOutput) FailedLocationCount() pulumi.IntOutput {
	return o.ApplyT(func(v LocationThresholdRuleConditionResponse) int { return v.FailedLocationCount }).(pulumi.IntOutput)
}

func (o LocationThresholdRuleConditionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v LocationThresholdRuleConditionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o LocationThresholdRuleConditionResponseOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationThresholdRuleConditionResponse) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
}

type ManagementEventAggregationCondition struct {
	Operator   *ConditionOperator `pulumi:"operator"`
	Threshold  *float64           `pulumi:"threshold"`
	WindowSize *string            `pulumi:"windowSize"`
}





type ManagementEventAggregationConditionInput interface {
	pulumi.Input

	ToManagementEventAggregationConditionOutput() ManagementEventAggregationConditionOutput
	ToManagementEventAggregationConditionOutputWithContext(context.Context) ManagementEventAggregationConditionOutput
}

type ManagementEventAggregationConditionArgs struct {
	Operator   ConditionOperatorPtrInput `pulumi:"operator"`
	Threshold  pulumi.Float64PtrInput    `pulumi:"threshold"`
	WindowSize pulumi.StringPtrInput     `pulumi:"windowSize"`
}

func (ManagementEventAggregationConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventAggregationCondition)(nil)).Elem()
}

func (i ManagementEventAggregationConditionArgs) ToManagementEventAggregationConditionOutput() ManagementEventAggregationConditionOutput {
	return i.ToManagementEventAggregationConditionOutputWithContext(context.Background())
}

func (i ManagementEventAggregationConditionArgs) ToManagementEventAggregationConditionOutputWithContext(ctx context.Context) ManagementEventAggregationConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventAggregationConditionOutput)
}

func (i ManagementEventAggregationConditionArgs) ToManagementEventAggregationConditionPtrOutput() ManagementEventAggregationConditionPtrOutput {
	return i.ToManagementEventAggregationConditionPtrOutputWithContext(context.Background())
}

func (i ManagementEventAggregationConditionArgs) ToManagementEventAggregationConditionPtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventAggregationConditionOutput).ToManagementEventAggregationConditionPtrOutputWithContext(ctx)
}









type ManagementEventAggregationConditionPtrInput interface {
	pulumi.Input

	ToManagementEventAggregationConditionPtrOutput() ManagementEventAggregationConditionPtrOutput
	ToManagementEventAggregationConditionPtrOutputWithContext(context.Context) ManagementEventAggregationConditionPtrOutput
}

type managementEventAggregationConditionPtrType ManagementEventAggregationConditionArgs

func ManagementEventAggregationConditionPtr(v *ManagementEventAggregationConditionArgs) ManagementEventAggregationConditionPtrInput {
	return (*managementEventAggregationConditionPtrType)(v)
}

func (*managementEventAggregationConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementEventAggregationCondition)(nil)).Elem()
}

func (i *managementEventAggregationConditionPtrType) ToManagementEventAggregationConditionPtrOutput() ManagementEventAggregationConditionPtrOutput {
	return i.ToManagementEventAggregationConditionPtrOutputWithContext(context.Background())
}

func (i *managementEventAggregationConditionPtrType) ToManagementEventAggregationConditionPtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventAggregationConditionPtrOutput)
}

type ManagementEventAggregationConditionOutput struct{ *pulumi.OutputState }

func (ManagementEventAggregationConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventAggregationCondition)(nil)).Elem()
}

func (o ManagementEventAggregationConditionOutput) ToManagementEventAggregationConditionOutput() ManagementEventAggregationConditionOutput {
	return o
}

func (o ManagementEventAggregationConditionOutput) ToManagementEventAggregationConditionOutputWithContext(ctx context.Context) ManagementEventAggregationConditionOutput {
	return o
}

func (o ManagementEventAggregationConditionOutput) ToManagementEventAggregationConditionPtrOutput() ManagementEventAggregationConditionPtrOutput {
	return o.ToManagementEventAggregationConditionPtrOutputWithContext(context.Background())
}

func (o ManagementEventAggregationConditionOutput) ToManagementEventAggregationConditionPtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementEventAggregationCondition) *ManagementEventAggregationCondition {
		return &v
	}).(ManagementEventAggregationConditionPtrOutput)
}

func (o ManagementEventAggregationConditionOutput) Operator() ConditionOperatorPtrOutput {
	return o.ApplyT(func(v ManagementEventAggregationCondition) *ConditionOperator { return v.Operator }).(ConditionOperatorPtrOutput)
}

func (o ManagementEventAggregationConditionOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ManagementEventAggregationCondition) *float64 { return v.Threshold }).(pulumi.Float64PtrOutput)
}

func (o ManagementEventAggregationConditionOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementEventAggregationCondition) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
}

type ManagementEventAggregationConditionPtrOutput struct{ *pulumi.OutputState }

func (ManagementEventAggregationConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementEventAggregationCondition)(nil)).Elem()
}

func (o ManagementEventAggregationConditionPtrOutput) ToManagementEventAggregationConditionPtrOutput() ManagementEventAggregationConditionPtrOutput {
	return o
}

func (o ManagementEventAggregationConditionPtrOutput) ToManagementEventAggregationConditionPtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionPtrOutput {
	return o
}

func (o ManagementEventAggregationConditionPtrOutput) Elem() ManagementEventAggregationConditionOutput {
	return o.ApplyT(func(v *ManagementEventAggregationCondition) ManagementEventAggregationCondition {
		if v != nil {
			return *v
		}
		var ret ManagementEventAggregationCondition
		return ret
	}).(ManagementEventAggregationConditionOutput)
}

func (o ManagementEventAggregationConditionPtrOutput) Operator() ConditionOperatorPtrOutput {
	return o.ApplyT(func(v *ManagementEventAggregationCondition) *ConditionOperator {
		if v == nil {
			return nil
		}
		return v.Operator
	}).(ConditionOperatorPtrOutput)
}

func (o ManagementEventAggregationConditionPtrOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ManagementEventAggregationCondition) *float64 {
		if v == nil {
			return nil
		}
		return v.Threshold
	}).(pulumi.Float64PtrOutput)
}

func (o ManagementEventAggregationConditionPtrOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementEventAggregationCondition) *string {
		if v == nil {
			return nil
		}
		return v.WindowSize
	}).(pulumi.StringPtrOutput)
}

type ManagementEventAggregationConditionResponse struct {
	Operator   *string  `pulumi:"operator"`
	Threshold  *float64 `pulumi:"threshold"`
	WindowSize *string  `pulumi:"windowSize"`
}





type ManagementEventAggregationConditionResponseInput interface {
	pulumi.Input

	ToManagementEventAggregationConditionResponseOutput() ManagementEventAggregationConditionResponseOutput
	ToManagementEventAggregationConditionResponseOutputWithContext(context.Context) ManagementEventAggregationConditionResponseOutput
}

type ManagementEventAggregationConditionResponseArgs struct {
	Operator   pulumi.StringPtrInput  `pulumi:"operator"`
	Threshold  pulumi.Float64PtrInput `pulumi:"threshold"`
	WindowSize pulumi.StringPtrInput  `pulumi:"windowSize"`
}

func (ManagementEventAggregationConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventAggregationConditionResponse)(nil)).Elem()
}

func (i ManagementEventAggregationConditionResponseArgs) ToManagementEventAggregationConditionResponseOutput() ManagementEventAggregationConditionResponseOutput {
	return i.ToManagementEventAggregationConditionResponseOutputWithContext(context.Background())
}

func (i ManagementEventAggregationConditionResponseArgs) ToManagementEventAggregationConditionResponseOutputWithContext(ctx context.Context) ManagementEventAggregationConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventAggregationConditionResponseOutput)
}

func (i ManagementEventAggregationConditionResponseArgs) ToManagementEventAggregationConditionResponsePtrOutput() ManagementEventAggregationConditionResponsePtrOutput {
	return i.ToManagementEventAggregationConditionResponsePtrOutputWithContext(context.Background())
}

func (i ManagementEventAggregationConditionResponseArgs) ToManagementEventAggregationConditionResponsePtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventAggregationConditionResponseOutput).ToManagementEventAggregationConditionResponsePtrOutputWithContext(ctx)
}









type ManagementEventAggregationConditionResponsePtrInput interface {
	pulumi.Input

	ToManagementEventAggregationConditionResponsePtrOutput() ManagementEventAggregationConditionResponsePtrOutput
	ToManagementEventAggregationConditionResponsePtrOutputWithContext(context.Context) ManagementEventAggregationConditionResponsePtrOutput
}

type managementEventAggregationConditionResponsePtrType ManagementEventAggregationConditionResponseArgs

func ManagementEventAggregationConditionResponsePtr(v *ManagementEventAggregationConditionResponseArgs) ManagementEventAggregationConditionResponsePtrInput {
	return (*managementEventAggregationConditionResponsePtrType)(v)
}

func (*managementEventAggregationConditionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementEventAggregationConditionResponse)(nil)).Elem()
}

func (i *managementEventAggregationConditionResponsePtrType) ToManagementEventAggregationConditionResponsePtrOutput() ManagementEventAggregationConditionResponsePtrOutput {
	return i.ToManagementEventAggregationConditionResponsePtrOutputWithContext(context.Background())
}

func (i *managementEventAggregationConditionResponsePtrType) ToManagementEventAggregationConditionResponsePtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventAggregationConditionResponsePtrOutput)
}

type ManagementEventAggregationConditionResponseOutput struct{ *pulumi.OutputState }

func (ManagementEventAggregationConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventAggregationConditionResponse)(nil)).Elem()
}

func (o ManagementEventAggregationConditionResponseOutput) ToManagementEventAggregationConditionResponseOutput() ManagementEventAggregationConditionResponseOutput {
	return o
}

func (o ManagementEventAggregationConditionResponseOutput) ToManagementEventAggregationConditionResponseOutputWithContext(ctx context.Context) ManagementEventAggregationConditionResponseOutput {
	return o
}

func (o ManagementEventAggregationConditionResponseOutput) ToManagementEventAggregationConditionResponsePtrOutput() ManagementEventAggregationConditionResponsePtrOutput {
	return o.ToManagementEventAggregationConditionResponsePtrOutputWithContext(context.Background())
}

func (o ManagementEventAggregationConditionResponseOutput) ToManagementEventAggregationConditionResponsePtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementEventAggregationConditionResponse) *ManagementEventAggregationConditionResponse {
		return &v
	}).(ManagementEventAggregationConditionResponsePtrOutput)
}

func (o ManagementEventAggregationConditionResponseOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementEventAggregationConditionResponse) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o ManagementEventAggregationConditionResponseOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ManagementEventAggregationConditionResponse) *float64 { return v.Threshold }).(pulumi.Float64PtrOutput)
}

func (o ManagementEventAggregationConditionResponseOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementEventAggregationConditionResponse) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
}

type ManagementEventAggregationConditionResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementEventAggregationConditionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementEventAggregationConditionResponse)(nil)).Elem()
}

func (o ManagementEventAggregationConditionResponsePtrOutput) ToManagementEventAggregationConditionResponsePtrOutput() ManagementEventAggregationConditionResponsePtrOutput {
	return o
}

func (o ManagementEventAggregationConditionResponsePtrOutput) ToManagementEventAggregationConditionResponsePtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionResponsePtrOutput {
	return o
}

func (o ManagementEventAggregationConditionResponsePtrOutput) Elem() ManagementEventAggregationConditionResponseOutput {
	return o.ApplyT(func(v *ManagementEventAggregationConditionResponse) ManagementEventAggregationConditionResponse {
		if v != nil {
			return *v
		}
		var ret ManagementEventAggregationConditionResponse
		return ret
	}).(ManagementEventAggregationConditionResponseOutput)
}

func (o ManagementEventAggregationConditionResponsePtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementEventAggregationConditionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o ManagementEventAggregationConditionResponsePtrOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ManagementEventAggregationConditionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Threshold
	}).(pulumi.Float64PtrOutput)
}

func (o ManagementEventAggregationConditionResponsePtrOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementEventAggregationConditionResponse) *string {
		if v == nil {
			return nil
		}
		return v.WindowSize
	}).(pulumi.StringPtrOutput)
}

type ManagementEventRuleCondition struct {
	Aggregation *ManagementEventAggregationCondition `pulumi:"aggregation"`
	DataSource  interface{}                          `pulumi:"dataSource"`
	OdataType   string                               `pulumi:"odataType"`
}





type ManagementEventRuleConditionInput interface {
	pulumi.Input

	ToManagementEventRuleConditionOutput() ManagementEventRuleConditionOutput
	ToManagementEventRuleConditionOutputWithContext(context.Context) ManagementEventRuleConditionOutput
}

type ManagementEventRuleConditionArgs struct {
	Aggregation ManagementEventAggregationConditionPtrInput `pulumi:"aggregation"`
	DataSource  pulumi.Input                                `pulumi:"dataSource"`
	OdataType   pulumi.StringInput                          `pulumi:"odataType"`
}

func (ManagementEventRuleConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventRuleCondition)(nil)).Elem()
}

func (i ManagementEventRuleConditionArgs) ToManagementEventRuleConditionOutput() ManagementEventRuleConditionOutput {
	return i.ToManagementEventRuleConditionOutputWithContext(context.Background())
}

func (i ManagementEventRuleConditionArgs) ToManagementEventRuleConditionOutputWithContext(ctx context.Context) ManagementEventRuleConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventRuleConditionOutput)
}

type ManagementEventRuleConditionOutput struct{ *pulumi.OutputState }

func (ManagementEventRuleConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventRuleCondition)(nil)).Elem()
}

func (o ManagementEventRuleConditionOutput) ToManagementEventRuleConditionOutput() ManagementEventRuleConditionOutput {
	return o
}

func (o ManagementEventRuleConditionOutput) ToManagementEventRuleConditionOutputWithContext(ctx context.Context) ManagementEventRuleConditionOutput {
	return o
}

func (o ManagementEventRuleConditionOutput) Aggregation() ManagementEventAggregationConditionPtrOutput {
	return o.ApplyT(func(v ManagementEventRuleCondition) *ManagementEventAggregationCondition { return v.Aggregation }).(ManagementEventAggregationConditionPtrOutput)
}

func (o ManagementEventRuleConditionOutput) DataSource() pulumi.AnyOutput {
	return o.ApplyT(func(v ManagementEventRuleCondition) interface{} { return v.DataSource }).(pulumi.AnyOutput)
}

func (o ManagementEventRuleConditionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementEventRuleCondition) string { return v.OdataType }).(pulumi.StringOutput)
}

type ManagementEventRuleConditionResponse struct {
	Aggregation *ManagementEventAggregationConditionResponse `pulumi:"aggregation"`
	DataSource  interface{}                                  `pulumi:"dataSource"`
	OdataType   string                                       `pulumi:"odataType"`
}





type ManagementEventRuleConditionResponseInput interface {
	pulumi.Input

	ToManagementEventRuleConditionResponseOutput() ManagementEventRuleConditionResponseOutput
	ToManagementEventRuleConditionResponseOutputWithContext(context.Context) ManagementEventRuleConditionResponseOutput
}

type ManagementEventRuleConditionResponseArgs struct {
	Aggregation ManagementEventAggregationConditionResponsePtrInput `pulumi:"aggregation"`
	DataSource  pulumi.Input                                        `pulumi:"dataSource"`
	OdataType   pulumi.StringInput                                  `pulumi:"odataType"`
}

func (ManagementEventRuleConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventRuleConditionResponse)(nil)).Elem()
}

func (i ManagementEventRuleConditionResponseArgs) ToManagementEventRuleConditionResponseOutput() ManagementEventRuleConditionResponseOutput {
	return i.ToManagementEventRuleConditionResponseOutputWithContext(context.Background())
}

func (i ManagementEventRuleConditionResponseArgs) ToManagementEventRuleConditionResponseOutputWithContext(ctx context.Context) ManagementEventRuleConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventRuleConditionResponseOutput)
}

type ManagementEventRuleConditionResponseOutput struct{ *pulumi.OutputState }

func (ManagementEventRuleConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventRuleConditionResponse)(nil)).Elem()
}

func (o ManagementEventRuleConditionResponseOutput) ToManagementEventRuleConditionResponseOutput() ManagementEventRuleConditionResponseOutput {
	return o
}

func (o ManagementEventRuleConditionResponseOutput) ToManagementEventRuleConditionResponseOutputWithContext(ctx context.Context) ManagementEventRuleConditionResponseOutput {
	return o
}

func (o ManagementEventRuleConditionResponseOutput) Aggregation() ManagementEventAggregationConditionResponsePtrOutput {
	return o.ApplyT(func(v ManagementEventRuleConditionResponse) *ManagementEventAggregationConditionResponse {
		return v.Aggregation
	}).(ManagementEventAggregationConditionResponsePtrOutput)
}

func (o ManagementEventRuleConditionResponseOutput) DataSource() pulumi.AnyOutput {
	return o.ApplyT(func(v ManagementEventRuleConditionResponse) interface{} { return v.DataSource }).(pulumi.AnyOutput)
}

func (o ManagementEventRuleConditionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementEventRuleConditionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type RetentionPolicy struct {
	Days    int  `pulumi:"days"`
	Enabled bool `pulumi:"enabled"`
}





type RetentionPolicyInput interface {
	pulumi.Input

	ToRetentionPolicyOutput() RetentionPolicyOutput
	ToRetentionPolicyOutputWithContext(context.Context) RetentionPolicyOutput
}

type RetentionPolicyArgs struct {
	Days    pulumi.IntInput  `pulumi:"days"`
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (RetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return i.ToRetentionPolicyOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput)
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput).ToRetentionPolicyPtrOutputWithContext(ctx)
}









type RetentionPolicyPtrInput interface {
	pulumi.Input

	ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput
	ToRetentionPolicyPtrOutputWithContext(context.Context) RetentionPolicyPtrOutput
}

type retentionPolicyPtrType RetentionPolicyArgs

func RetentionPolicyPtr(v *RetentionPolicyArgs) RetentionPolicyPtrInput {
	return (*retentionPolicyPtrType)(v)
}

func (*retentionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyPtrOutput)
}

type RetentionPolicyOutput struct{ *pulumi.OutputState }

func (RetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionPolicy) *RetentionPolicy {
		return &v
	}).(RetentionPolicyPtrOutput)
}

func (o RetentionPolicyOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func(v RetentionPolicy) int { return v.Days }).(pulumi.IntOutput)
}

func (o RetentionPolicyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RetentionPolicy) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type RetentionPolicyPtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) Elem() RetentionPolicyOutput {
	return o.ApplyT(func(v *RetentionPolicy) RetentionPolicy {
		if v != nil {
			return *v
		}
		var ret RetentionPolicy
		return ret
	}).(RetentionPolicyOutput)
}

func (o RetentionPolicyPtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type RetentionPolicyResponse struct {
	Days    int  `pulumi:"days"`
	Enabled bool `pulumi:"enabled"`
}





type RetentionPolicyResponseInput interface {
	pulumi.Input

	ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput
	ToRetentionPolicyResponseOutputWithContext(context.Context) RetentionPolicyResponseOutput
}

type RetentionPolicyResponseArgs struct {
	Days    pulumi.IntInput  `pulumi:"days"`
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (RetentionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyResponse)(nil)).Elem()
}

func (i RetentionPolicyResponseArgs) ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput {
	return i.ToRetentionPolicyResponseOutputWithContext(context.Background())
}

func (i RetentionPolicyResponseArgs) ToRetentionPolicyResponseOutputWithContext(ctx context.Context) RetentionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyResponseOutput)
}

func (i RetentionPolicyResponseArgs) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return i.ToRetentionPolicyResponsePtrOutputWithContext(context.Background())
}

func (i RetentionPolicyResponseArgs) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyResponseOutput).ToRetentionPolicyResponsePtrOutputWithContext(ctx)
}









type RetentionPolicyResponsePtrInput interface {
	pulumi.Input

	ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput
	ToRetentionPolicyResponsePtrOutputWithContext(context.Context) RetentionPolicyResponsePtrOutput
}

type retentionPolicyResponsePtrType RetentionPolicyResponseArgs

func RetentionPolicyResponsePtr(v *RetentionPolicyResponseArgs) RetentionPolicyResponsePtrInput {
	return (*retentionPolicyResponsePtrType)(v)
}

func (*retentionPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicyResponse)(nil)).Elem()
}

func (i *retentionPolicyResponsePtrType) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return i.ToRetentionPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *retentionPolicyResponsePtrType) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyResponsePtrOutput)
}

type RetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutputWithContext(ctx context.Context) RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return o.ToRetentionPolicyResponsePtrOutputWithContext(context.Background())
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionPolicyResponse) *RetentionPolicyResponse {
		return &v
	}).(RetentionPolicyResponsePtrOutput)
}

func (o RetentionPolicyResponseOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) int { return v.Days }).(pulumi.IntOutput)
}

func (o RetentionPolicyResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type RetentionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) Elem() RetentionPolicyResponseOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) RetentionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret RetentionPolicyResponse
		return ret
	}).(RetentionPolicyResponseOutput)
}

func (o RetentionPolicyResponsePtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type RuleEmailAction struct {
	CustomEmails        []string `pulumi:"customEmails"`
	OdataType           string   `pulumi:"odataType"`
	SendToServiceOwners *bool    `pulumi:"sendToServiceOwners"`
}





type RuleEmailActionInput interface {
	pulumi.Input

	ToRuleEmailActionOutput() RuleEmailActionOutput
	ToRuleEmailActionOutputWithContext(context.Context) RuleEmailActionOutput
}

type RuleEmailActionArgs struct {
	CustomEmails        pulumi.StringArrayInput `pulumi:"customEmails"`
	OdataType           pulumi.StringInput      `pulumi:"odataType"`
	SendToServiceOwners pulumi.BoolPtrInput     `pulumi:"sendToServiceOwners"`
}

func (RuleEmailActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleEmailAction)(nil)).Elem()
}

func (i RuleEmailActionArgs) ToRuleEmailActionOutput() RuleEmailActionOutput {
	return i.ToRuleEmailActionOutputWithContext(context.Background())
}

func (i RuleEmailActionArgs) ToRuleEmailActionOutputWithContext(ctx context.Context) RuleEmailActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleEmailActionOutput)
}

type RuleEmailActionOutput struct{ *pulumi.OutputState }

func (RuleEmailActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleEmailAction)(nil)).Elem()
}

func (o RuleEmailActionOutput) ToRuleEmailActionOutput() RuleEmailActionOutput {
	return o
}

func (o RuleEmailActionOutput) ToRuleEmailActionOutputWithContext(ctx context.Context) RuleEmailActionOutput {
	return o
}

func (o RuleEmailActionOutput) CustomEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RuleEmailAction) []string { return v.CustomEmails }).(pulumi.StringArrayOutput)
}

func (o RuleEmailActionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleEmailAction) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleEmailActionOutput) SendToServiceOwners() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RuleEmailAction) *bool { return v.SendToServiceOwners }).(pulumi.BoolPtrOutput)
}

type RuleEmailActionResponse struct {
	CustomEmails        []string `pulumi:"customEmails"`
	OdataType           string   `pulumi:"odataType"`
	SendToServiceOwners *bool    `pulumi:"sendToServiceOwners"`
}





type RuleEmailActionResponseInput interface {
	pulumi.Input

	ToRuleEmailActionResponseOutput() RuleEmailActionResponseOutput
	ToRuleEmailActionResponseOutputWithContext(context.Context) RuleEmailActionResponseOutput
}

type RuleEmailActionResponseArgs struct {
	CustomEmails        pulumi.StringArrayInput `pulumi:"customEmails"`
	OdataType           pulumi.StringInput      `pulumi:"odataType"`
	SendToServiceOwners pulumi.BoolPtrInput     `pulumi:"sendToServiceOwners"`
}

func (RuleEmailActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleEmailActionResponse)(nil)).Elem()
}

func (i RuleEmailActionResponseArgs) ToRuleEmailActionResponseOutput() RuleEmailActionResponseOutput {
	return i.ToRuleEmailActionResponseOutputWithContext(context.Background())
}

func (i RuleEmailActionResponseArgs) ToRuleEmailActionResponseOutputWithContext(ctx context.Context) RuleEmailActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleEmailActionResponseOutput)
}

type RuleEmailActionResponseOutput struct{ *pulumi.OutputState }

func (RuleEmailActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleEmailActionResponse)(nil)).Elem()
}

func (o RuleEmailActionResponseOutput) ToRuleEmailActionResponseOutput() RuleEmailActionResponseOutput {
	return o
}

func (o RuleEmailActionResponseOutput) ToRuleEmailActionResponseOutputWithContext(ctx context.Context) RuleEmailActionResponseOutput {
	return o
}

func (o RuleEmailActionResponseOutput) CustomEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RuleEmailActionResponse) []string { return v.CustomEmails }).(pulumi.StringArrayOutput)
}

func (o RuleEmailActionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleEmailActionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleEmailActionResponseOutput) SendToServiceOwners() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RuleEmailActionResponse) *bool { return v.SendToServiceOwners }).(pulumi.BoolPtrOutput)
}

type RuleManagementEventClaimsDataSource struct {
	EmailAddress *string `pulumi:"emailAddress"`
}





type RuleManagementEventClaimsDataSourceInput interface {
	pulumi.Input

	ToRuleManagementEventClaimsDataSourceOutput() RuleManagementEventClaimsDataSourceOutput
	ToRuleManagementEventClaimsDataSourceOutputWithContext(context.Context) RuleManagementEventClaimsDataSourceOutput
}

type RuleManagementEventClaimsDataSourceArgs struct {
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
}

func (RuleManagementEventClaimsDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventClaimsDataSource)(nil)).Elem()
}

func (i RuleManagementEventClaimsDataSourceArgs) ToRuleManagementEventClaimsDataSourceOutput() RuleManagementEventClaimsDataSourceOutput {
	return i.ToRuleManagementEventClaimsDataSourceOutputWithContext(context.Background())
}

func (i RuleManagementEventClaimsDataSourceArgs) ToRuleManagementEventClaimsDataSourceOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventClaimsDataSourceOutput)
}

func (i RuleManagementEventClaimsDataSourceArgs) ToRuleManagementEventClaimsDataSourcePtrOutput() RuleManagementEventClaimsDataSourcePtrOutput {
	return i.ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(context.Background())
}

func (i RuleManagementEventClaimsDataSourceArgs) ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventClaimsDataSourceOutput).ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(ctx)
}









type RuleManagementEventClaimsDataSourcePtrInput interface {
	pulumi.Input

	ToRuleManagementEventClaimsDataSourcePtrOutput() RuleManagementEventClaimsDataSourcePtrOutput
	ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(context.Context) RuleManagementEventClaimsDataSourcePtrOutput
}

type ruleManagementEventClaimsDataSourcePtrType RuleManagementEventClaimsDataSourceArgs

func RuleManagementEventClaimsDataSourcePtr(v *RuleManagementEventClaimsDataSourceArgs) RuleManagementEventClaimsDataSourcePtrInput {
	return (*ruleManagementEventClaimsDataSourcePtrType)(v)
}

func (*ruleManagementEventClaimsDataSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleManagementEventClaimsDataSource)(nil)).Elem()
}

func (i *ruleManagementEventClaimsDataSourcePtrType) ToRuleManagementEventClaimsDataSourcePtrOutput() RuleManagementEventClaimsDataSourcePtrOutput {
	return i.ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(context.Background())
}

func (i *ruleManagementEventClaimsDataSourcePtrType) ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventClaimsDataSourcePtrOutput)
}

type RuleManagementEventClaimsDataSourceOutput struct{ *pulumi.OutputState }

func (RuleManagementEventClaimsDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventClaimsDataSource)(nil)).Elem()
}

func (o RuleManagementEventClaimsDataSourceOutput) ToRuleManagementEventClaimsDataSourceOutput() RuleManagementEventClaimsDataSourceOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourceOutput) ToRuleManagementEventClaimsDataSourceOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourceOutput) ToRuleManagementEventClaimsDataSourcePtrOutput() RuleManagementEventClaimsDataSourcePtrOutput {
	return o.ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(context.Background())
}

func (o RuleManagementEventClaimsDataSourceOutput) ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RuleManagementEventClaimsDataSource) *RuleManagementEventClaimsDataSource {
		return &v
	}).(RuleManagementEventClaimsDataSourcePtrOutput)
}

func (o RuleManagementEventClaimsDataSourceOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventClaimsDataSource) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

type RuleManagementEventClaimsDataSourcePtrOutput struct{ *pulumi.OutputState }

func (RuleManagementEventClaimsDataSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleManagementEventClaimsDataSource)(nil)).Elem()
}

func (o RuleManagementEventClaimsDataSourcePtrOutput) ToRuleManagementEventClaimsDataSourcePtrOutput() RuleManagementEventClaimsDataSourcePtrOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourcePtrOutput) ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourcePtrOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourcePtrOutput) Elem() RuleManagementEventClaimsDataSourceOutput {
	return o.ApplyT(func(v *RuleManagementEventClaimsDataSource) RuleManagementEventClaimsDataSource {
		if v != nil {
			return *v
		}
		var ret RuleManagementEventClaimsDataSource
		return ret
	}).(RuleManagementEventClaimsDataSourceOutput)
}

func (o RuleManagementEventClaimsDataSourcePtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleManagementEventClaimsDataSource) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

type RuleManagementEventClaimsDataSourceResponse struct {
	EmailAddress *string `pulumi:"emailAddress"`
}





type RuleManagementEventClaimsDataSourceResponseInput interface {
	pulumi.Input

	ToRuleManagementEventClaimsDataSourceResponseOutput() RuleManagementEventClaimsDataSourceResponseOutput
	ToRuleManagementEventClaimsDataSourceResponseOutputWithContext(context.Context) RuleManagementEventClaimsDataSourceResponseOutput
}

type RuleManagementEventClaimsDataSourceResponseArgs struct {
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
}

func (RuleManagementEventClaimsDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventClaimsDataSourceResponse)(nil)).Elem()
}

func (i RuleManagementEventClaimsDataSourceResponseArgs) ToRuleManagementEventClaimsDataSourceResponseOutput() RuleManagementEventClaimsDataSourceResponseOutput {
	return i.ToRuleManagementEventClaimsDataSourceResponseOutputWithContext(context.Background())
}

func (i RuleManagementEventClaimsDataSourceResponseArgs) ToRuleManagementEventClaimsDataSourceResponseOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventClaimsDataSourceResponseOutput)
}

func (i RuleManagementEventClaimsDataSourceResponseArgs) ToRuleManagementEventClaimsDataSourceResponsePtrOutput() RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return i.ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(context.Background())
}

func (i RuleManagementEventClaimsDataSourceResponseArgs) ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventClaimsDataSourceResponseOutput).ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(ctx)
}









type RuleManagementEventClaimsDataSourceResponsePtrInput interface {
	pulumi.Input

	ToRuleManagementEventClaimsDataSourceResponsePtrOutput() RuleManagementEventClaimsDataSourceResponsePtrOutput
	ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(context.Context) RuleManagementEventClaimsDataSourceResponsePtrOutput
}

type ruleManagementEventClaimsDataSourceResponsePtrType RuleManagementEventClaimsDataSourceResponseArgs

func RuleManagementEventClaimsDataSourceResponsePtr(v *RuleManagementEventClaimsDataSourceResponseArgs) RuleManagementEventClaimsDataSourceResponsePtrInput {
	return (*ruleManagementEventClaimsDataSourceResponsePtrType)(v)
}

func (*ruleManagementEventClaimsDataSourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleManagementEventClaimsDataSourceResponse)(nil)).Elem()
}

func (i *ruleManagementEventClaimsDataSourceResponsePtrType) ToRuleManagementEventClaimsDataSourceResponsePtrOutput() RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return i.ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(context.Background())
}

func (i *ruleManagementEventClaimsDataSourceResponsePtrType) ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventClaimsDataSourceResponsePtrOutput)
}

type RuleManagementEventClaimsDataSourceResponseOutput struct{ *pulumi.OutputState }

func (RuleManagementEventClaimsDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventClaimsDataSourceResponse)(nil)).Elem()
}

func (o RuleManagementEventClaimsDataSourceResponseOutput) ToRuleManagementEventClaimsDataSourceResponseOutput() RuleManagementEventClaimsDataSourceResponseOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourceResponseOutput) ToRuleManagementEventClaimsDataSourceResponseOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceResponseOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourceResponseOutput) ToRuleManagementEventClaimsDataSourceResponsePtrOutput() RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return o.ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(context.Background())
}

func (o RuleManagementEventClaimsDataSourceResponseOutput) ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RuleManagementEventClaimsDataSourceResponse) *RuleManagementEventClaimsDataSourceResponse {
		return &v
	}).(RuleManagementEventClaimsDataSourceResponsePtrOutput)
}

func (o RuleManagementEventClaimsDataSourceResponseOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventClaimsDataSourceResponse) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

type RuleManagementEventClaimsDataSourceResponsePtrOutput struct{ *pulumi.OutputState }

func (RuleManagementEventClaimsDataSourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleManagementEventClaimsDataSourceResponse)(nil)).Elem()
}

func (o RuleManagementEventClaimsDataSourceResponsePtrOutput) ToRuleManagementEventClaimsDataSourceResponsePtrOutput() RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourceResponsePtrOutput) ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourceResponsePtrOutput) Elem() RuleManagementEventClaimsDataSourceResponseOutput {
	return o.ApplyT(func(v *RuleManagementEventClaimsDataSourceResponse) RuleManagementEventClaimsDataSourceResponse {
		if v != nil {
			return *v
		}
		var ret RuleManagementEventClaimsDataSourceResponse
		return ret
	}).(RuleManagementEventClaimsDataSourceResponseOutput)
}

func (o RuleManagementEventClaimsDataSourceResponsePtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleManagementEventClaimsDataSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

type RuleManagementEventDataSource struct {
	Claims               *RuleManagementEventClaimsDataSource `pulumi:"claims"`
	EventName            *string                              `pulumi:"eventName"`
	EventSource          *string                              `pulumi:"eventSource"`
	LegacyResourceId     *string                              `pulumi:"legacyResourceId"`
	Level                *string                              `pulumi:"level"`
	MetricNamespace      *string                              `pulumi:"metricNamespace"`
	OdataType            string                               `pulumi:"odataType"`
	OperationName        *string                              `pulumi:"operationName"`
	ResourceGroupName    *string                              `pulumi:"resourceGroupName"`
	ResourceLocation     *string                              `pulumi:"resourceLocation"`
	ResourceProviderName *string                              `pulumi:"resourceProviderName"`
	ResourceUri          *string                              `pulumi:"resourceUri"`
	Status               *string                              `pulumi:"status"`
	SubStatus            *string                              `pulumi:"subStatus"`
}





type RuleManagementEventDataSourceInput interface {
	pulumi.Input

	ToRuleManagementEventDataSourceOutput() RuleManagementEventDataSourceOutput
	ToRuleManagementEventDataSourceOutputWithContext(context.Context) RuleManagementEventDataSourceOutput
}

type RuleManagementEventDataSourceArgs struct {
	Claims               RuleManagementEventClaimsDataSourcePtrInput `pulumi:"claims"`
	EventName            pulumi.StringPtrInput                       `pulumi:"eventName"`
	EventSource          pulumi.StringPtrInput                       `pulumi:"eventSource"`
	LegacyResourceId     pulumi.StringPtrInput                       `pulumi:"legacyResourceId"`
	Level                pulumi.StringPtrInput                       `pulumi:"level"`
	MetricNamespace      pulumi.StringPtrInput                       `pulumi:"metricNamespace"`
	OdataType            pulumi.StringInput                          `pulumi:"odataType"`
	OperationName        pulumi.StringPtrInput                       `pulumi:"operationName"`
	ResourceGroupName    pulumi.StringPtrInput                       `pulumi:"resourceGroupName"`
	ResourceLocation     pulumi.StringPtrInput                       `pulumi:"resourceLocation"`
	ResourceProviderName pulumi.StringPtrInput                       `pulumi:"resourceProviderName"`
	ResourceUri          pulumi.StringPtrInput                       `pulumi:"resourceUri"`
	Status               pulumi.StringPtrInput                       `pulumi:"status"`
	SubStatus            pulumi.StringPtrInput                       `pulumi:"subStatus"`
}

func (RuleManagementEventDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventDataSource)(nil)).Elem()
}

func (i RuleManagementEventDataSourceArgs) ToRuleManagementEventDataSourceOutput() RuleManagementEventDataSourceOutput {
	return i.ToRuleManagementEventDataSourceOutputWithContext(context.Background())
}

func (i RuleManagementEventDataSourceArgs) ToRuleManagementEventDataSourceOutputWithContext(ctx context.Context) RuleManagementEventDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventDataSourceOutput)
}

type RuleManagementEventDataSourceOutput struct{ *pulumi.OutputState }

func (RuleManagementEventDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventDataSource)(nil)).Elem()
}

func (o RuleManagementEventDataSourceOutput) ToRuleManagementEventDataSourceOutput() RuleManagementEventDataSourceOutput {
	return o
}

func (o RuleManagementEventDataSourceOutput) ToRuleManagementEventDataSourceOutputWithContext(ctx context.Context) RuleManagementEventDataSourceOutput {
	return o
}

func (o RuleManagementEventDataSourceOutput) Claims() RuleManagementEventClaimsDataSourcePtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *RuleManagementEventClaimsDataSource { return v.Claims }).(RuleManagementEventClaimsDataSourcePtrOutput)
}

func (o RuleManagementEventDataSourceOutput) EventName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.EventName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) EventSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.EventSource }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) LegacyResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.LegacyResourceId }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleManagementEventDataSourceOutput) OperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.OperationName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) ResourceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.ResourceProviderName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) SubStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.SubStatus }).(pulumi.StringPtrOutput)
}

type RuleManagementEventDataSourceResponse struct {
	Claims               *RuleManagementEventClaimsDataSourceResponse `pulumi:"claims"`
	EventName            *string                                      `pulumi:"eventName"`
	EventSource          *string                                      `pulumi:"eventSource"`
	LegacyResourceId     *string                                      `pulumi:"legacyResourceId"`
	Level                *string                                      `pulumi:"level"`
	MetricNamespace      *string                                      `pulumi:"metricNamespace"`
	OdataType            string                                       `pulumi:"odataType"`
	OperationName        *string                                      `pulumi:"operationName"`
	ResourceGroupName    *string                                      `pulumi:"resourceGroupName"`
	ResourceLocation     *string                                      `pulumi:"resourceLocation"`
	ResourceProviderName *string                                      `pulumi:"resourceProviderName"`
	ResourceUri          *string                                      `pulumi:"resourceUri"`
	Status               *string                                      `pulumi:"status"`
	SubStatus            *string                                      `pulumi:"subStatus"`
}





type RuleManagementEventDataSourceResponseInput interface {
	pulumi.Input

	ToRuleManagementEventDataSourceResponseOutput() RuleManagementEventDataSourceResponseOutput
	ToRuleManagementEventDataSourceResponseOutputWithContext(context.Context) RuleManagementEventDataSourceResponseOutput
}

type RuleManagementEventDataSourceResponseArgs struct {
	Claims               RuleManagementEventClaimsDataSourceResponsePtrInput `pulumi:"claims"`
	EventName            pulumi.StringPtrInput                               `pulumi:"eventName"`
	EventSource          pulumi.StringPtrInput                               `pulumi:"eventSource"`
	LegacyResourceId     pulumi.StringPtrInput                               `pulumi:"legacyResourceId"`
	Level                pulumi.StringPtrInput                               `pulumi:"level"`
	MetricNamespace      pulumi.StringPtrInput                               `pulumi:"metricNamespace"`
	OdataType            pulumi.StringInput                                  `pulumi:"odataType"`
	OperationName        pulumi.StringPtrInput                               `pulumi:"operationName"`
	ResourceGroupName    pulumi.StringPtrInput                               `pulumi:"resourceGroupName"`
	ResourceLocation     pulumi.StringPtrInput                               `pulumi:"resourceLocation"`
	ResourceProviderName pulumi.StringPtrInput                               `pulumi:"resourceProviderName"`
	ResourceUri          pulumi.StringPtrInput                               `pulumi:"resourceUri"`
	Status               pulumi.StringPtrInput                               `pulumi:"status"`
	SubStatus            pulumi.StringPtrInput                               `pulumi:"subStatus"`
}

func (RuleManagementEventDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventDataSourceResponse)(nil)).Elem()
}

func (i RuleManagementEventDataSourceResponseArgs) ToRuleManagementEventDataSourceResponseOutput() RuleManagementEventDataSourceResponseOutput {
	return i.ToRuleManagementEventDataSourceResponseOutputWithContext(context.Background())
}

func (i RuleManagementEventDataSourceResponseArgs) ToRuleManagementEventDataSourceResponseOutputWithContext(ctx context.Context) RuleManagementEventDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventDataSourceResponseOutput)
}

type RuleManagementEventDataSourceResponseOutput struct{ *pulumi.OutputState }

func (RuleManagementEventDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventDataSourceResponse)(nil)).Elem()
}

func (o RuleManagementEventDataSourceResponseOutput) ToRuleManagementEventDataSourceResponseOutput() RuleManagementEventDataSourceResponseOutput {
	return o
}

func (o RuleManagementEventDataSourceResponseOutput) ToRuleManagementEventDataSourceResponseOutputWithContext(ctx context.Context) RuleManagementEventDataSourceResponseOutput {
	return o
}

func (o RuleManagementEventDataSourceResponseOutput) Claims() RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *RuleManagementEventClaimsDataSourceResponse {
		return v.Claims
	}).(RuleManagementEventClaimsDataSourceResponsePtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) EventName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.EventName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) EventSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.EventSource }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) LegacyResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.LegacyResourceId }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) OperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.OperationName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) ResourceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.ResourceProviderName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) SubStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.SubStatus }).(pulumi.StringPtrOutput)
}

type RuleMetricDataSource struct {
	LegacyResourceId *string `pulumi:"legacyResourceId"`
	MetricName       *string `pulumi:"metricName"`
	MetricNamespace  *string `pulumi:"metricNamespace"`
	OdataType        string  `pulumi:"odataType"`
	ResourceLocation *string `pulumi:"resourceLocation"`
	ResourceUri      *string `pulumi:"resourceUri"`
}





type RuleMetricDataSourceInput interface {
	pulumi.Input

	ToRuleMetricDataSourceOutput() RuleMetricDataSourceOutput
	ToRuleMetricDataSourceOutputWithContext(context.Context) RuleMetricDataSourceOutput
}

type RuleMetricDataSourceArgs struct {
	LegacyResourceId pulumi.StringPtrInput `pulumi:"legacyResourceId"`
	MetricName       pulumi.StringPtrInput `pulumi:"metricName"`
	MetricNamespace  pulumi.StringPtrInput `pulumi:"metricNamespace"`
	OdataType        pulumi.StringInput    `pulumi:"odataType"`
	ResourceLocation pulumi.StringPtrInput `pulumi:"resourceLocation"`
	ResourceUri      pulumi.StringPtrInput `pulumi:"resourceUri"`
}

func (RuleMetricDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleMetricDataSource)(nil)).Elem()
}

func (i RuleMetricDataSourceArgs) ToRuleMetricDataSourceOutput() RuleMetricDataSourceOutput {
	return i.ToRuleMetricDataSourceOutputWithContext(context.Background())
}

func (i RuleMetricDataSourceArgs) ToRuleMetricDataSourceOutputWithContext(ctx context.Context) RuleMetricDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMetricDataSourceOutput)
}

type RuleMetricDataSourceOutput struct{ *pulumi.OutputState }

func (RuleMetricDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleMetricDataSource)(nil)).Elem()
}

func (o RuleMetricDataSourceOutput) ToRuleMetricDataSourceOutput() RuleMetricDataSourceOutput {
	return o
}

func (o RuleMetricDataSourceOutput) ToRuleMetricDataSourceOutputWithContext(ctx context.Context) RuleMetricDataSourceOutput {
	return o
}

func (o RuleMetricDataSourceOutput) LegacyResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSource) *string { return v.LegacyResourceId }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceOutput) MetricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSource) *string { return v.MetricName }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSource) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleMetricDataSource) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleMetricDataSourceOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSource) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSource) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

type RuleMetricDataSourceResponse struct {
	LegacyResourceId *string `pulumi:"legacyResourceId"`
	MetricName       *string `pulumi:"metricName"`
	MetricNamespace  *string `pulumi:"metricNamespace"`
	OdataType        string  `pulumi:"odataType"`
	ResourceLocation *string `pulumi:"resourceLocation"`
	ResourceUri      *string `pulumi:"resourceUri"`
}





type RuleMetricDataSourceResponseInput interface {
	pulumi.Input

	ToRuleMetricDataSourceResponseOutput() RuleMetricDataSourceResponseOutput
	ToRuleMetricDataSourceResponseOutputWithContext(context.Context) RuleMetricDataSourceResponseOutput
}

type RuleMetricDataSourceResponseArgs struct {
	LegacyResourceId pulumi.StringPtrInput `pulumi:"legacyResourceId"`
	MetricName       pulumi.StringPtrInput `pulumi:"metricName"`
	MetricNamespace  pulumi.StringPtrInput `pulumi:"metricNamespace"`
	OdataType        pulumi.StringInput    `pulumi:"odataType"`
	ResourceLocation pulumi.StringPtrInput `pulumi:"resourceLocation"`
	ResourceUri      pulumi.StringPtrInput `pulumi:"resourceUri"`
}

func (RuleMetricDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleMetricDataSourceResponse)(nil)).Elem()
}

func (i RuleMetricDataSourceResponseArgs) ToRuleMetricDataSourceResponseOutput() RuleMetricDataSourceResponseOutput {
	return i.ToRuleMetricDataSourceResponseOutputWithContext(context.Background())
}

func (i RuleMetricDataSourceResponseArgs) ToRuleMetricDataSourceResponseOutputWithContext(ctx context.Context) RuleMetricDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMetricDataSourceResponseOutput)
}

type RuleMetricDataSourceResponseOutput struct{ *pulumi.OutputState }

func (RuleMetricDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleMetricDataSourceResponse)(nil)).Elem()
}

func (o RuleMetricDataSourceResponseOutput) ToRuleMetricDataSourceResponseOutput() RuleMetricDataSourceResponseOutput {
	return o
}

func (o RuleMetricDataSourceResponseOutput) ToRuleMetricDataSourceResponseOutputWithContext(ctx context.Context) RuleMetricDataSourceResponseOutput {
	return o
}

func (o RuleMetricDataSourceResponseOutput) LegacyResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSourceResponse) *string { return v.LegacyResourceId }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceResponseOutput) MetricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSourceResponse) *string { return v.MetricName }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceResponseOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSourceResponse) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleMetricDataSourceResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleMetricDataSourceResponseOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSourceResponse) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceResponseOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSourceResponse) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

type RuleWebhookAction struct {
	OdataType  string            `pulumi:"odataType"`
	Properties map[string]string `pulumi:"properties"`
	ServiceUri *string           `pulumi:"serviceUri"`
}





type RuleWebhookActionInput interface {
	pulumi.Input

	ToRuleWebhookActionOutput() RuleWebhookActionOutput
	ToRuleWebhookActionOutputWithContext(context.Context) RuleWebhookActionOutput
}

type RuleWebhookActionArgs struct {
	OdataType  pulumi.StringInput    `pulumi:"odataType"`
	Properties pulumi.StringMapInput `pulumi:"properties"`
	ServiceUri pulumi.StringPtrInput `pulumi:"serviceUri"`
}

func (RuleWebhookActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleWebhookAction)(nil)).Elem()
}

func (i RuleWebhookActionArgs) ToRuleWebhookActionOutput() RuleWebhookActionOutput {
	return i.ToRuleWebhookActionOutputWithContext(context.Background())
}

func (i RuleWebhookActionArgs) ToRuleWebhookActionOutputWithContext(ctx context.Context) RuleWebhookActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleWebhookActionOutput)
}

type RuleWebhookActionOutput struct{ *pulumi.OutputState }

func (RuleWebhookActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleWebhookAction)(nil)).Elem()
}

func (o RuleWebhookActionOutput) ToRuleWebhookActionOutput() RuleWebhookActionOutput {
	return o
}

func (o RuleWebhookActionOutput) ToRuleWebhookActionOutputWithContext(ctx context.Context) RuleWebhookActionOutput {
	return o
}

func (o RuleWebhookActionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleWebhookAction) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleWebhookActionOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v RuleWebhookAction) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o RuleWebhookActionOutput) ServiceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleWebhookAction) *string { return v.ServiceUri }).(pulumi.StringPtrOutput)
}

type RuleWebhookActionResponse struct {
	OdataType  string            `pulumi:"odataType"`
	Properties map[string]string `pulumi:"properties"`
	ServiceUri *string           `pulumi:"serviceUri"`
}





type RuleWebhookActionResponseInput interface {
	pulumi.Input

	ToRuleWebhookActionResponseOutput() RuleWebhookActionResponseOutput
	ToRuleWebhookActionResponseOutputWithContext(context.Context) RuleWebhookActionResponseOutput
}

type RuleWebhookActionResponseArgs struct {
	OdataType  pulumi.StringInput    `pulumi:"odataType"`
	Properties pulumi.StringMapInput `pulumi:"properties"`
	ServiceUri pulumi.StringPtrInput `pulumi:"serviceUri"`
}

func (RuleWebhookActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleWebhookActionResponse)(nil)).Elem()
}

func (i RuleWebhookActionResponseArgs) ToRuleWebhookActionResponseOutput() RuleWebhookActionResponseOutput {
	return i.ToRuleWebhookActionResponseOutputWithContext(context.Background())
}

func (i RuleWebhookActionResponseArgs) ToRuleWebhookActionResponseOutputWithContext(ctx context.Context) RuleWebhookActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleWebhookActionResponseOutput)
}

type RuleWebhookActionResponseOutput struct{ *pulumi.OutputState }

func (RuleWebhookActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleWebhookActionResponse)(nil)).Elem()
}

func (o RuleWebhookActionResponseOutput) ToRuleWebhookActionResponseOutput() RuleWebhookActionResponseOutput {
	return o
}

func (o RuleWebhookActionResponseOutput) ToRuleWebhookActionResponseOutputWithContext(ctx context.Context) RuleWebhookActionResponseOutput {
	return o
}

func (o RuleWebhookActionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleWebhookActionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleWebhookActionResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v RuleWebhookActionResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o RuleWebhookActionResponseOutput) ServiceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleWebhookActionResponse) *string { return v.ServiceUri }).(pulumi.StringPtrOutput)
}

type ThresholdRuleCondition struct {
	DataSource      interface{}              `pulumi:"dataSource"`
	OdataType       string                   `pulumi:"odataType"`
	Operator        ConditionOperator        `pulumi:"operator"`
	Threshold       float64                  `pulumi:"threshold"`
	TimeAggregation *TimeAggregationOperator `pulumi:"timeAggregation"`
	WindowSize      *string                  `pulumi:"windowSize"`
}





type ThresholdRuleConditionInput interface {
	pulumi.Input

	ToThresholdRuleConditionOutput() ThresholdRuleConditionOutput
	ToThresholdRuleConditionOutputWithContext(context.Context) ThresholdRuleConditionOutput
}

type ThresholdRuleConditionArgs struct {
	DataSource      pulumi.Input                    `pulumi:"dataSource"`
	OdataType       pulumi.StringInput              `pulumi:"odataType"`
	Operator        ConditionOperatorInput          `pulumi:"operator"`
	Threshold       pulumi.Float64Input             `pulumi:"threshold"`
	TimeAggregation TimeAggregationOperatorPtrInput `pulumi:"timeAggregation"`
	WindowSize      pulumi.StringPtrInput           `pulumi:"windowSize"`
}

func (ThresholdRuleConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdRuleCondition)(nil)).Elem()
}

func (i ThresholdRuleConditionArgs) ToThresholdRuleConditionOutput() ThresholdRuleConditionOutput {
	return i.ToThresholdRuleConditionOutputWithContext(context.Background())
}

func (i ThresholdRuleConditionArgs) ToThresholdRuleConditionOutputWithContext(ctx context.Context) ThresholdRuleConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThresholdRuleConditionOutput)
}

type ThresholdRuleConditionOutput struct{ *pulumi.OutputState }

func (ThresholdRuleConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdRuleCondition)(nil)).Elem()
}

func (o ThresholdRuleConditionOutput) ToThresholdRuleConditionOutput() ThresholdRuleConditionOutput {
	return o
}

func (o ThresholdRuleConditionOutput) ToThresholdRuleConditionOutputWithContext(ctx context.Context) ThresholdRuleConditionOutput {
	return o
}

func (o ThresholdRuleConditionOutput) DataSource() pulumi.AnyOutput {
	return o.ApplyT(func(v ThresholdRuleCondition) interface{} { return v.DataSource }).(pulumi.AnyOutput)
}

func (o ThresholdRuleConditionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdRuleCondition) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ThresholdRuleConditionOutput) Operator() ConditionOperatorOutput {
	return o.ApplyT(func(v ThresholdRuleCondition) ConditionOperator { return v.Operator }).(ConditionOperatorOutput)
}

func (o ThresholdRuleConditionOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v ThresholdRuleCondition) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o ThresholdRuleConditionOutput) TimeAggregation() TimeAggregationOperatorPtrOutput {
	return o.ApplyT(func(v ThresholdRuleCondition) *TimeAggregationOperator { return v.TimeAggregation }).(TimeAggregationOperatorPtrOutput)
}

func (o ThresholdRuleConditionOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThresholdRuleCondition) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
}

type ThresholdRuleConditionResponse struct {
	DataSource      interface{} `pulumi:"dataSource"`
	OdataType       string      `pulumi:"odataType"`
	Operator        string      `pulumi:"operator"`
	Threshold       float64     `pulumi:"threshold"`
	TimeAggregation *string     `pulumi:"timeAggregation"`
	WindowSize      *string     `pulumi:"windowSize"`
}





type ThresholdRuleConditionResponseInput interface {
	pulumi.Input

	ToThresholdRuleConditionResponseOutput() ThresholdRuleConditionResponseOutput
	ToThresholdRuleConditionResponseOutputWithContext(context.Context) ThresholdRuleConditionResponseOutput
}

type ThresholdRuleConditionResponseArgs struct {
	DataSource      pulumi.Input          `pulumi:"dataSource"`
	OdataType       pulumi.StringInput    `pulumi:"odataType"`
	Operator        pulumi.StringInput    `pulumi:"operator"`
	Threshold       pulumi.Float64Input   `pulumi:"threshold"`
	TimeAggregation pulumi.StringPtrInput `pulumi:"timeAggregation"`
	WindowSize      pulumi.StringPtrInput `pulumi:"windowSize"`
}

func (ThresholdRuleConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdRuleConditionResponse)(nil)).Elem()
}

func (i ThresholdRuleConditionResponseArgs) ToThresholdRuleConditionResponseOutput() ThresholdRuleConditionResponseOutput {
	return i.ToThresholdRuleConditionResponseOutputWithContext(context.Background())
}

func (i ThresholdRuleConditionResponseArgs) ToThresholdRuleConditionResponseOutputWithContext(ctx context.Context) ThresholdRuleConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThresholdRuleConditionResponseOutput)
}

type ThresholdRuleConditionResponseOutput struct{ *pulumi.OutputState }

func (ThresholdRuleConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdRuleConditionResponse)(nil)).Elem()
}

func (o ThresholdRuleConditionResponseOutput) ToThresholdRuleConditionResponseOutput() ThresholdRuleConditionResponseOutput {
	return o
}

func (o ThresholdRuleConditionResponseOutput) ToThresholdRuleConditionResponseOutputWithContext(ctx context.Context) ThresholdRuleConditionResponseOutput {
	return o
}

func (o ThresholdRuleConditionResponseOutput) DataSource() pulumi.AnyOutput {
	return o.ApplyT(func(v ThresholdRuleConditionResponse) interface{} { return v.DataSource }).(pulumi.AnyOutput)
}

func (o ThresholdRuleConditionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdRuleConditionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ThresholdRuleConditionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdRuleConditionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ThresholdRuleConditionResponseOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v ThresholdRuleConditionResponse) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o ThresholdRuleConditionResponseOutput) TimeAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThresholdRuleConditionResponse) *string { return v.TimeAggregation }).(pulumi.StringPtrOutput)
}

func (o ThresholdRuleConditionResponseOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThresholdRuleConditionResponse) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LocationThresholdRuleConditionOutput{})
	pulumi.RegisterOutputType(LocationThresholdRuleConditionResponseOutput{})
	pulumi.RegisterOutputType(ManagementEventAggregationConditionOutput{})
	pulumi.RegisterOutputType(ManagementEventAggregationConditionPtrOutput{})
	pulumi.RegisterOutputType(ManagementEventAggregationConditionResponseOutput{})
	pulumi.RegisterOutputType(ManagementEventAggregationConditionResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementEventRuleConditionOutput{})
	pulumi.RegisterOutputType(ManagementEventRuleConditionResponseOutput{})
	pulumi.RegisterOutputType(RetentionPolicyOutput{})
	pulumi.RegisterOutputType(RetentionPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(RuleEmailActionOutput{})
	pulumi.RegisterOutputType(RuleEmailActionResponseOutput{})
	pulumi.RegisterOutputType(RuleManagementEventClaimsDataSourceOutput{})
	pulumi.RegisterOutputType(RuleManagementEventClaimsDataSourcePtrOutput{})
	pulumi.RegisterOutputType(RuleManagementEventClaimsDataSourceResponseOutput{})
	pulumi.RegisterOutputType(RuleManagementEventClaimsDataSourceResponsePtrOutput{})
	pulumi.RegisterOutputType(RuleManagementEventDataSourceOutput{})
	pulumi.RegisterOutputType(RuleManagementEventDataSourceResponseOutput{})
	pulumi.RegisterOutputType(RuleMetricDataSourceOutput{})
	pulumi.RegisterOutputType(RuleMetricDataSourceResponseOutput{})
	pulumi.RegisterOutputType(RuleWebhookActionOutput{})
	pulumi.RegisterOutputType(RuleWebhookActionResponseOutput{})
	pulumi.RegisterOutputType(ThresholdRuleConditionOutput{})
	pulumi.RegisterOutputType(ThresholdRuleConditionResponseOutput{})
}

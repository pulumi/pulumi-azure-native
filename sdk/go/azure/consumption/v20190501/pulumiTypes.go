


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BudgetTimePeriod struct {
	EndDate   *string `pulumi:"endDate"`
	StartDate string  `pulumi:"startDate"`
}





type BudgetTimePeriodInput interface {
	pulumi.Input

	ToBudgetTimePeriodOutput() BudgetTimePeriodOutput
	ToBudgetTimePeriodOutputWithContext(context.Context) BudgetTimePeriodOutput
}

type BudgetTimePeriodArgs struct {
	EndDate   pulumi.StringPtrInput `pulumi:"endDate"`
	StartDate pulumi.StringInput    `pulumi:"startDate"`
}

func (BudgetTimePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetTimePeriod)(nil)).Elem()
}

func (i BudgetTimePeriodArgs) ToBudgetTimePeriodOutput() BudgetTimePeriodOutput {
	return i.ToBudgetTimePeriodOutputWithContext(context.Background())
}

func (i BudgetTimePeriodArgs) ToBudgetTimePeriodOutputWithContext(ctx context.Context) BudgetTimePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetTimePeriodOutput)
}

type BudgetTimePeriodOutput struct{ *pulumi.OutputState }

func (BudgetTimePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetTimePeriod)(nil)).Elem()
}

func (o BudgetTimePeriodOutput) ToBudgetTimePeriodOutput() BudgetTimePeriodOutput {
	return o
}

func (o BudgetTimePeriodOutput) ToBudgetTimePeriodOutputWithContext(ctx context.Context) BudgetTimePeriodOutput {
	return o
}

func (o BudgetTimePeriodOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BudgetTimePeriod) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o BudgetTimePeriodOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v BudgetTimePeriod) string { return v.StartDate }).(pulumi.StringOutput)
}

type BudgetTimePeriodResponse struct {
	EndDate   *string `pulumi:"endDate"`
	StartDate string  `pulumi:"startDate"`
}

type BudgetTimePeriodResponseOutput struct{ *pulumi.OutputState }

func (BudgetTimePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetTimePeriodResponse)(nil)).Elem()
}

func (o BudgetTimePeriodResponseOutput) ToBudgetTimePeriodResponseOutput() BudgetTimePeriodResponseOutput {
	return o
}

func (o BudgetTimePeriodResponseOutput) ToBudgetTimePeriodResponseOutputWithContext(ctx context.Context) BudgetTimePeriodResponseOutput {
	return o
}

func (o BudgetTimePeriodResponseOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BudgetTimePeriodResponse) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o BudgetTimePeriodResponseOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v BudgetTimePeriodResponse) string { return v.StartDate }).(pulumi.StringOutput)
}

type CurrentSpendResponse struct {
	Amount float64 `pulumi:"amount"`
	Unit   string  `pulumi:"unit"`
}

type CurrentSpendResponseOutput struct{ *pulumi.OutputState }

func (CurrentSpendResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CurrentSpendResponse)(nil)).Elem()
}

func (o CurrentSpendResponseOutput) ToCurrentSpendResponseOutput() CurrentSpendResponseOutput {
	return o
}

func (o CurrentSpendResponseOutput) ToCurrentSpendResponseOutputWithContext(ctx context.Context) CurrentSpendResponseOutput {
	return o
}

func (o CurrentSpendResponseOutput) Amount() pulumi.Float64Output {
	return o.ApplyT(func(v CurrentSpendResponse) float64 { return v.Amount }).(pulumi.Float64Output)
}

func (o CurrentSpendResponseOutput) Unit() pulumi.StringOutput {
	return o.ApplyT(func(v CurrentSpendResponse) string { return v.Unit }).(pulumi.StringOutput)
}

type Filter struct {
	Meters         []string            `pulumi:"meters"`
	ResourceGroups []string            `pulumi:"resourceGroups"`
	Resources      []string            `pulumi:"resources"`
	Tags           map[string][]string `pulumi:"tags"`
}





type FilterInput interface {
	pulumi.Input

	ToFilterOutput() FilterOutput
	ToFilterOutputWithContext(context.Context) FilterOutput
}

type FilterArgs struct {
	Meters         pulumi.StringArrayInput    `pulumi:"meters"`
	ResourceGroups pulumi.StringArrayInput    `pulumi:"resourceGroups"`
	Resources      pulumi.StringArrayInput    `pulumi:"resources"`
	Tags           pulumi.StringArrayMapInput `pulumi:"tags"`
}

func (FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Filter)(nil)).Elem()
}

func (i FilterArgs) ToFilterOutput() FilterOutput {
	return i.ToFilterOutputWithContext(context.Background())
}

func (i FilterArgs) ToFilterOutputWithContext(ctx context.Context) FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterOutput)
}

func (i FilterArgs) ToFilterPtrOutput() FilterPtrOutput {
	return i.ToFilterPtrOutputWithContext(context.Background())
}

func (i FilterArgs) ToFilterPtrOutputWithContext(ctx context.Context) FilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterOutput).ToFilterPtrOutputWithContext(ctx)
}









type FilterPtrInput interface {
	pulumi.Input

	ToFilterPtrOutput() FilterPtrOutput
	ToFilterPtrOutputWithContext(context.Context) FilterPtrOutput
}

type filterPtrType FilterArgs

func FilterPtr(v *FilterArgs) FilterPtrInput {
	return (*filterPtrType)(v)
}

func (*filterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Filter)(nil)).Elem()
}

func (i *filterPtrType) ToFilterPtrOutput() FilterPtrOutput {
	return i.ToFilterPtrOutputWithContext(context.Background())
}

func (i *filterPtrType) ToFilterPtrOutputWithContext(ctx context.Context) FilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterPtrOutput)
}

type FilterOutput struct{ *pulumi.OutputState }

func (FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Filter)(nil)).Elem()
}

func (o FilterOutput) ToFilterOutput() FilterOutput {
	return o
}

func (o FilterOutput) ToFilterOutputWithContext(ctx context.Context) FilterOutput {
	return o
}

func (o FilterOutput) ToFilterPtrOutput() FilterPtrOutput {
	return o.ToFilterPtrOutputWithContext(context.Background())
}

func (o FilterOutput) ToFilterPtrOutputWithContext(ctx context.Context) FilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Filter) *Filter {
		return &v
	}).(FilterPtrOutput)
}

func (o FilterOutput) Meters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Filter) []string { return v.Meters }).(pulumi.StringArrayOutput)
}

func (o FilterOutput) ResourceGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Filter) []string { return v.ResourceGroups }).(pulumi.StringArrayOutput)
}

func (o FilterOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Filter) []string { return v.Resources }).(pulumi.StringArrayOutput)
}

func (o FilterOutput) Tags() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v Filter) map[string][]string { return v.Tags }).(pulumi.StringArrayMapOutput)
}

type FilterPtrOutput struct{ *pulumi.OutputState }

func (FilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Filter)(nil)).Elem()
}

func (o FilterPtrOutput) ToFilterPtrOutput() FilterPtrOutput {
	return o
}

func (o FilterPtrOutput) ToFilterPtrOutputWithContext(ctx context.Context) FilterPtrOutput {
	return o
}

func (o FilterPtrOutput) Elem() FilterOutput {
	return o.ApplyT(func(v *Filter) Filter {
		if v != nil {
			return *v
		}
		var ret Filter
		return ret
	}).(FilterOutput)
}

func (o FilterPtrOutput) Meters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Filter) []string {
		if v == nil {
			return nil
		}
		return v.Meters
	}).(pulumi.StringArrayOutput)
}

func (o FilterPtrOutput) ResourceGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Filter) []string {
		if v == nil {
			return nil
		}
		return v.ResourceGroups
	}).(pulumi.StringArrayOutput)
}

func (o FilterPtrOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Filter) []string {
		if v == nil {
			return nil
		}
		return v.Resources
	}).(pulumi.StringArrayOutput)
}

func (o FilterPtrOutput) Tags() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *Filter) map[string][]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringArrayMapOutput)
}

type FilterResponse struct {
	Meters         []string            `pulumi:"meters"`
	ResourceGroups []string            `pulumi:"resourceGroups"`
	Resources      []string            `pulumi:"resources"`
	Tags           map[string][]string `pulumi:"tags"`
}

type FilterResponseOutput struct{ *pulumi.OutputState }

func (FilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterResponse)(nil)).Elem()
}

func (o FilterResponseOutput) ToFilterResponseOutput() FilterResponseOutput {
	return o
}

func (o FilterResponseOutput) ToFilterResponseOutputWithContext(ctx context.Context) FilterResponseOutput {
	return o
}

func (o FilterResponseOutput) Meters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FilterResponse) []string { return v.Meters }).(pulumi.StringArrayOutput)
}

func (o FilterResponseOutput) ResourceGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FilterResponse) []string { return v.ResourceGroups }).(pulumi.StringArrayOutput)
}

func (o FilterResponseOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FilterResponse) []string { return v.Resources }).(pulumi.StringArrayOutput)
}

func (o FilterResponseOutput) Tags() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v FilterResponse) map[string][]string { return v.Tags }).(pulumi.StringArrayMapOutput)
}

type FilterResponsePtrOutput struct{ *pulumi.OutputState }

func (FilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FilterResponse)(nil)).Elem()
}

func (o FilterResponsePtrOutput) ToFilterResponsePtrOutput() FilterResponsePtrOutput {
	return o
}

func (o FilterResponsePtrOutput) ToFilterResponsePtrOutputWithContext(ctx context.Context) FilterResponsePtrOutput {
	return o
}

func (o FilterResponsePtrOutput) Elem() FilterResponseOutput {
	return o.ApplyT(func(v *FilterResponse) FilterResponse {
		if v != nil {
			return *v
		}
		var ret FilterResponse
		return ret
	}).(FilterResponseOutput)
}

func (o FilterResponsePtrOutput) Meters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FilterResponse) []string {
		if v == nil {
			return nil
		}
		return v.Meters
	}).(pulumi.StringArrayOutput)
}

func (o FilterResponsePtrOutput) ResourceGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FilterResponse) []string {
		if v == nil {
			return nil
		}
		return v.ResourceGroups
	}).(pulumi.StringArrayOutput)
}

func (o FilterResponsePtrOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FilterResponse) []string {
		if v == nil {
			return nil
		}
		return v.Resources
	}).(pulumi.StringArrayOutput)
}

func (o FilterResponsePtrOutput) Tags() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *FilterResponse) map[string][]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringArrayMapOutput)
}

type Notification struct {
	ContactEmails []string `pulumi:"contactEmails"`
	ContactGroups []string `pulumi:"contactGroups"`
	ContactRoles  []string `pulumi:"contactRoles"`
	Enabled       bool     `pulumi:"enabled"`
	Operator      string   `pulumi:"operator"`
	Threshold     float64  `pulumi:"threshold"`
	ThresholdType *string  `pulumi:"thresholdType"`
}


func (val *Notification) Defaults() *Notification {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ThresholdType) {
		thresholdType_ := "Actual"
		tmp.ThresholdType = &thresholdType_
	}
	return &tmp
}





type NotificationInput interface {
	pulumi.Input

	ToNotificationOutput() NotificationOutput
	ToNotificationOutputWithContext(context.Context) NotificationOutput
}

type NotificationArgs struct {
	ContactEmails pulumi.StringArrayInput `pulumi:"contactEmails"`
	ContactGroups pulumi.StringArrayInput `pulumi:"contactGroups"`
	ContactRoles  pulumi.StringArrayInput `pulumi:"contactRoles"`
	Enabled       pulumi.BoolInput        `pulumi:"enabled"`
	Operator      pulumi.StringInput      `pulumi:"operator"`
	Threshold     pulumi.Float64Input     `pulumi:"threshold"`
	ThresholdType pulumi.StringPtrInput   `pulumi:"thresholdType"`
}


func (val *NotificationArgs) Defaults() *NotificationArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ThresholdType) {
		tmp.ThresholdType = pulumi.StringPtr("Actual")
	}
	return &tmp
}
func (NotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Notification)(nil)).Elem()
}

func (i NotificationArgs) ToNotificationOutput() NotificationOutput {
	return i.ToNotificationOutputWithContext(context.Background())
}

func (i NotificationArgs) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationOutput)
}





type NotificationMapInput interface {
	pulumi.Input

	ToNotificationMapOutput() NotificationMapOutput
	ToNotificationMapOutputWithContext(context.Context) NotificationMapOutput
}

type NotificationMap map[string]NotificationInput

func (NotificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Notification)(nil)).Elem()
}

func (i NotificationMap) ToNotificationMapOutput() NotificationMapOutput {
	return i.ToNotificationMapOutputWithContext(context.Background())
}

func (i NotificationMap) ToNotificationMapOutputWithContext(ctx context.Context) NotificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationMapOutput)
}

type NotificationOutput struct{ *pulumi.OutputState }

func (NotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Notification)(nil)).Elem()
}

func (o NotificationOutput) ToNotificationOutput() NotificationOutput {
	return o
}

func (o NotificationOutput) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return o
}

func (o NotificationOutput) ContactEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Notification) []string { return v.ContactEmails }).(pulumi.StringArrayOutput)
}

func (o NotificationOutput) ContactGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Notification) []string { return v.ContactGroups }).(pulumi.StringArrayOutput)
}

func (o NotificationOutput) ContactRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Notification) []string { return v.ContactRoles }).(pulumi.StringArrayOutput)
}

func (o NotificationOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v Notification) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o NotificationOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v Notification) string { return v.Operator }).(pulumi.StringOutput)
}

func (o NotificationOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v Notification) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o NotificationOutput) ThresholdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Notification) *string { return v.ThresholdType }).(pulumi.StringPtrOutput)
}

type NotificationMapOutput struct{ *pulumi.OutputState }

func (NotificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Notification)(nil)).Elem()
}

func (o NotificationMapOutput) ToNotificationMapOutput() NotificationMapOutput {
	return o
}

func (o NotificationMapOutput) ToNotificationMapOutputWithContext(ctx context.Context) NotificationMapOutput {
	return o
}

func (o NotificationMapOutput) MapIndex(k pulumi.StringInput) NotificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Notification {
		return vs[0].(map[string]Notification)[vs[1].(string)]
	}).(NotificationOutput)
}

type NotificationResponse struct {
	ContactEmails []string `pulumi:"contactEmails"`
	ContactGroups []string `pulumi:"contactGroups"`
	ContactRoles  []string `pulumi:"contactRoles"`
	Enabled       bool     `pulumi:"enabled"`
	Operator      string   `pulumi:"operator"`
	Threshold     float64  `pulumi:"threshold"`
	ThresholdType *string  `pulumi:"thresholdType"`
}


func (val *NotificationResponse) Defaults() *NotificationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ThresholdType) {
		thresholdType_ := "Actual"
		tmp.ThresholdType = &thresholdType_
	}
	return &tmp
}

type NotificationResponseOutput struct{ *pulumi.OutputState }

func (NotificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationResponse)(nil)).Elem()
}

func (o NotificationResponseOutput) ToNotificationResponseOutput() NotificationResponseOutput {
	return o
}

func (o NotificationResponseOutput) ToNotificationResponseOutputWithContext(ctx context.Context) NotificationResponseOutput {
	return o
}

func (o NotificationResponseOutput) ContactEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NotificationResponse) []string { return v.ContactEmails }).(pulumi.StringArrayOutput)
}

func (o NotificationResponseOutput) ContactGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NotificationResponse) []string { return v.ContactGroups }).(pulumi.StringArrayOutput)
}

func (o NotificationResponseOutput) ContactRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NotificationResponse) []string { return v.ContactRoles }).(pulumi.StringArrayOutput)
}

func (o NotificationResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v NotificationResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o NotificationResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o NotificationResponseOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v NotificationResponse) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o NotificationResponseOutput) ThresholdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationResponse) *string { return v.ThresholdType }).(pulumi.StringPtrOutput)
}

type NotificationResponseMapOutput struct{ *pulumi.OutputState }

func (NotificationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NotificationResponse)(nil)).Elem()
}

func (o NotificationResponseMapOutput) ToNotificationResponseMapOutput() NotificationResponseMapOutput {
	return o
}

func (o NotificationResponseMapOutput) ToNotificationResponseMapOutputWithContext(ctx context.Context) NotificationResponseMapOutput {
	return o
}

func (o NotificationResponseMapOutput) MapIndex(k pulumi.StringInput) NotificationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NotificationResponse {
		return vs[0].(map[string]NotificationResponse)[vs[1].(string)]
	}).(NotificationResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(BudgetTimePeriodOutput{})
	pulumi.RegisterOutputType(BudgetTimePeriodResponseOutput{})
	pulumi.RegisterOutputType(CurrentSpendResponseOutput{})
	pulumi.RegisterOutputType(FilterOutput{})
	pulumi.RegisterOutputType(FilterPtrOutput{})
	pulumi.RegisterOutputType(FilterResponseOutput{})
	pulumi.RegisterOutputType(FilterResponsePtrOutput{})
	pulumi.RegisterOutputType(NotificationOutput{})
	pulumi.RegisterOutputType(NotificationMapOutput{})
	pulumi.RegisterOutputType(NotificationResponseOutput{})
	pulumi.RegisterOutputType(NotificationResponseMapOutput{})
}

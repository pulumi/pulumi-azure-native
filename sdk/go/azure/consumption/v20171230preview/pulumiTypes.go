


package v20171230preview

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

type Notification struct {
	ContactEmails []string `pulumi:"contactEmails"`
	ContactRoles  []string `pulumi:"contactRoles"`
	Enabled       bool     `pulumi:"enabled"`
	Operator      string   `pulumi:"operator"`
	Threshold     float64  `pulumi:"threshold"`
}





type NotificationInput interface {
	pulumi.Input

	ToNotificationOutput() NotificationOutput
	ToNotificationOutputWithContext(context.Context) NotificationOutput
}

type NotificationArgs struct {
	ContactEmails pulumi.StringArrayInput `pulumi:"contactEmails"`
	ContactRoles  pulumi.StringArrayInput `pulumi:"contactRoles"`
	Enabled       pulumi.BoolInput        `pulumi:"enabled"`
	Operator      pulumi.StringInput      `pulumi:"operator"`
	Threshold     pulumi.Float64Input     `pulumi:"threshold"`
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
	ContactRoles  []string `pulumi:"contactRoles"`
	Enabled       bool     `pulumi:"enabled"`
	Operator      string   `pulumi:"operator"`
	Threshold     float64  `pulumi:"threshold"`
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
	pulumi.RegisterOutputType(NotificationOutput{})
	pulumi.RegisterOutputType(NotificationMapOutput{})
	pulumi.RegisterOutputType(NotificationResponseOutput{})
	pulumi.RegisterOutputType(NotificationResponseMapOutput{})
}

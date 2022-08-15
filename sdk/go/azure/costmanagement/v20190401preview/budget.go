


package v20190401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Budget struct {
	pulumi.CustomResourceState

	Amount        pulumi.Float64Output                `pulumi:"amount"`
	Category      pulumi.StringOutput                 `pulumi:"category"`
	CurrentSpend  CurrentSpendResponseOutput          `pulumi:"currentSpend"`
	ETag          pulumi.StringPtrOutput              `pulumi:"eTag"`
	Filter        ReportConfigFilterResponsePtrOutput `pulumi:"filter"`
	Name          pulumi.StringOutput                 `pulumi:"name"`
	Notifications NotificationResponseMapOutput       `pulumi:"notifications"`
	TimeGrain     pulumi.StringOutput                 `pulumi:"timeGrain"`
	TimePeriod    BudgetTimePeriodResponseOutput      `pulumi:"timePeriod"`
	Type          pulumi.StringOutput                 `pulumi:"type"`
}


func NewBudget(ctx *pulumi.Context,
	name string, args *BudgetArgs, opts ...pulumi.ResourceOption) (*Budget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Amount == nil {
		return nil, errors.New("invalid value for required argument 'Amount'")
	}
	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.TimeGrain == nil {
		return nil, errors.New("invalid value for required argument 'TimeGrain'")
	}
	if args.TimePeriod == nil {
		return nil, errors.New("invalid value for required argument 'TimePeriod'")
	}
	var resource Budget
	err := ctx.RegisterResource("azure-native:costmanagement/v20190401preview:Budget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBudget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetState, opts ...pulumi.ResourceOption) (*Budget, error) {
	var resource Budget
	err := ctx.ReadResource("azure-native:costmanagement/v20190401preview:Budget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type budgetState struct {
}

type BudgetState struct {
}

func (BudgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetState)(nil)).Elem()
}

type budgetArgs struct {
	Amount        float64                 `pulumi:"amount"`
	BudgetName    *string                 `pulumi:"budgetName"`
	Category      string                  `pulumi:"category"`
	ETag          *string                 `pulumi:"eTag"`
	Filter        *ReportConfigFilter     `pulumi:"filter"`
	Notifications map[string]Notification `pulumi:"notifications"`
	Scope         string                  `pulumi:"scope"`
	TimeGrain     string                  `pulumi:"timeGrain"`
	TimePeriod    BudgetTimePeriod        `pulumi:"timePeriod"`
}


type BudgetArgs struct {
	Amount        pulumi.Float64Input
	BudgetName    pulumi.StringPtrInput
	Category      pulumi.StringInput
	ETag          pulumi.StringPtrInput
	Filter        ReportConfigFilterPtrInput
	Notifications NotificationMapInput
	Scope         pulumi.StringInput
	TimeGrain     pulumi.StringInput
	TimePeriod    BudgetTimePeriodInput
}

func (BudgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetArgs)(nil)).Elem()
}

type BudgetInput interface {
	pulumi.Input

	ToBudgetOutput() BudgetOutput
	ToBudgetOutputWithContext(ctx context.Context) BudgetOutput
}

func (*Budget) ElementType() reflect.Type {
	return reflect.TypeOf((**Budget)(nil)).Elem()
}

func (i *Budget) ToBudgetOutput() BudgetOutput {
	return i.ToBudgetOutputWithContext(context.Background())
}

func (i *Budget) ToBudgetOutputWithContext(ctx context.Context) BudgetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetOutput)
}

type BudgetOutput struct{ *pulumi.OutputState }

func (BudgetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Budget)(nil)).Elem()
}

func (o BudgetOutput) ToBudgetOutput() BudgetOutput {
	return o
}

func (o BudgetOutput) ToBudgetOutputWithContext(ctx context.Context) BudgetOutput {
	return o
}

func (o BudgetOutput) Amount() pulumi.Float64Output {
	return o.ApplyT(func(v *Budget) pulumi.Float64Output { return v.Amount }).(pulumi.Float64Output)
}

func (o BudgetOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

func (o BudgetOutput) CurrentSpend() CurrentSpendResponseOutput {
	return o.ApplyT(func(v *Budget) CurrentSpendResponseOutput { return v.CurrentSpend }).(CurrentSpendResponseOutput)
}

func (o BudgetOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o BudgetOutput) Filter() ReportConfigFilterResponsePtrOutput {
	return o.ApplyT(func(v *Budget) ReportConfigFilterResponsePtrOutput { return v.Filter }).(ReportConfigFilterResponsePtrOutput)
}

func (o BudgetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BudgetOutput) Notifications() NotificationResponseMapOutput {
	return o.ApplyT(func(v *Budget) NotificationResponseMapOutput { return v.Notifications }).(NotificationResponseMapOutput)
}

func (o BudgetOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringOutput { return v.TimeGrain }).(pulumi.StringOutput)
}

func (o BudgetOutput) TimePeriod() BudgetTimePeriodResponseOutput {
	return o.ApplyT(func(v *Budget) BudgetTimePeriodResponseOutput { return v.TimePeriod }).(BudgetTimePeriodResponseOutput)
}

func (o BudgetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BudgetOutput{})
}

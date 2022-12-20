


package v20190401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupBudget(ctx *pulumi.Context, args *LookupBudgetArgs, opts ...pulumi.InvokeOption) (*LookupBudgetResult, error) {
	var rv LookupBudgetResult
	err := ctx.Invoke("azure-native:consumption/v20190401preview:getBudget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBudgetArgs struct {
	BudgetName string `pulumi:"budgetName"`
	Scope      string `pulumi:"scope"`
}


type LookupBudgetResult struct {
	Amount        float64                         `pulumi:"amount"`
	Category      string                          `pulumi:"category"`
	CurrentSpend  CurrentSpendResponse            `pulumi:"currentSpend"`
	ETag          *string                         `pulumi:"eTag"`
	Filters       *FiltersResponse                `pulumi:"filters"`
	Id            string                          `pulumi:"id"`
	Name          string                          `pulumi:"name"`
	Notifications map[string]NotificationResponse `pulumi:"notifications"`
	TimeGrain     string                          `pulumi:"timeGrain"`
	TimePeriod    BudgetTimePeriodResponse        `pulumi:"timePeriod"`
	Type          string                          `pulumi:"type"`
}

func LookupBudgetOutput(ctx *pulumi.Context, args LookupBudgetOutputArgs, opts ...pulumi.InvokeOption) LookupBudgetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBudgetResult, error) {
			args := v.(LookupBudgetArgs)
			r, err := LookupBudget(ctx, &args, opts...)
			var s LookupBudgetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBudgetResultOutput)
}

type LookupBudgetOutputArgs struct {
	BudgetName pulumi.StringInput `pulumi:"budgetName"`
	Scope      pulumi.StringInput `pulumi:"scope"`
}

func (LookupBudgetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBudgetArgs)(nil)).Elem()
}


type LookupBudgetResultOutput struct{ *pulumi.OutputState }

func (LookupBudgetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBudgetResult)(nil)).Elem()
}

func (o LookupBudgetResultOutput) ToLookupBudgetResultOutput() LookupBudgetResultOutput {
	return o
}

func (o LookupBudgetResultOutput) ToLookupBudgetResultOutputWithContext(ctx context.Context) LookupBudgetResultOutput {
	return o
}

func (o LookupBudgetResultOutput) Amount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupBudgetResult) float64 { return v.Amount }).(pulumi.Float64Output)
}

func (o LookupBudgetResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o LookupBudgetResultOutput) CurrentSpend() CurrentSpendResponseOutput {
	return o.ApplyT(func(v LookupBudgetResult) CurrentSpendResponse { return v.CurrentSpend }).(CurrentSpendResponseOutput)
}

func (o LookupBudgetResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBudgetResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupBudgetResultOutput) Filters() FiltersResponsePtrOutput {
	return o.ApplyT(func(v LookupBudgetResult) *FiltersResponse { return v.Filters }).(FiltersResponsePtrOutput)
}

func (o LookupBudgetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBudgetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBudgetResultOutput) Notifications() NotificationResponseMapOutput {
	return o.ApplyT(func(v LookupBudgetResult) map[string]NotificationResponse { return v.Notifications }).(NotificationResponseMapOutput)
}

func (o LookupBudgetResultOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.TimeGrain }).(pulumi.StringOutput)
}

func (o LookupBudgetResultOutput) TimePeriod() BudgetTimePeriodResponseOutput {
	return o.ApplyT(func(v LookupBudgetResult) BudgetTimePeriodResponse { return v.TimePeriod }).(BudgetTimePeriodResponseOutput)
}

func (o LookupBudgetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBudgetResultOutput{})
}

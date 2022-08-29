


package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupBudgetByResourceGroupName(ctx *pulumi.Context, args *LookupBudgetByResourceGroupNameArgs, opts ...pulumi.InvokeOption) (*LookupBudgetByResourceGroupNameResult, error) {
	var rv LookupBudgetByResourceGroupNameResult
	err := ctx.Invoke("azure-native:consumption/v20181001:getBudgetByResourceGroupName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBudgetByResourceGroupNameArgs struct {
	BudgetName        string `pulumi:"budgetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBudgetByResourceGroupNameResult struct {
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

func LookupBudgetByResourceGroupNameOutput(ctx *pulumi.Context, args LookupBudgetByResourceGroupNameOutputArgs, opts ...pulumi.InvokeOption) LookupBudgetByResourceGroupNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBudgetByResourceGroupNameResult, error) {
			args := v.(LookupBudgetByResourceGroupNameArgs)
			r, err := LookupBudgetByResourceGroupName(ctx, &args, opts...)
			var s LookupBudgetByResourceGroupNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBudgetByResourceGroupNameResultOutput)
}

type LookupBudgetByResourceGroupNameOutputArgs struct {
	BudgetName        pulumi.StringInput `pulumi:"budgetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBudgetByResourceGroupNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBudgetByResourceGroupNameArgs)(nil)).Elem()
}


type LookupBudgetByResourceGroupNameResultOutput struct{ *pulumi.OutputState }

func (LookupBudgetByResourceGroupNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBudgetByResourceGroupNameResult)(nil)).Elem()
}

func (o LookupBudgetByResourceGroupNameResultOutput) ToLookupBudgetByResourceGroupNameResultOutput() LookupBudgetByResourceGroupNameResultOutput {
	return o
}

func (o LookupBudgetByResourceGroupNameResultOutput) ToLookupBudgetByResourceGroupNameResultOutputWithContext(ctx context.Context) LookupBudgetByResourceGroupNameResultOutput {
	return o
}

func (o LookupBudgetByResourceGroupNameResultOutput) Amount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupBudgetByResourceGroupNameResult) float64 { return v.Amount }).(pulumi.Float64Output)
}

func (o LookupBudgetByResourceGroupNameResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetByResourceGroupNameResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o LookupBudgetByResourceGroupNameResultOutput) CurrentSpend() CurrentSpendResponseOutput {
	return o.ApplyT(func(v LookupBudgetByResourceGroupNameResult) CurrentSpendResponse { return v.CurrentSpend }).(CurrentSpendResponseOutput)
}

func (o LookupBudgetByResourceGroupNameResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBudgetByResourceGroupNameResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupBudgetByResourceGroupNameResultOutput) Filters() FiltersResponsePtrOutput {
	return o.ApplyT(func(v LookupBudgetByResourceGroupNameResult) *FiltersResponse { return v.Filters }).(FiltersResponsePtrOutput)
}

func (o LookupBudgetByResourceGroupNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetByResourceGroupNameResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBudgetByResourceGroupNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetByResourceGroupNameResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBudgetByResourceGroupNameResultOutput) Notifications() NotificationResponseMapOutput {
	return o.ApplyT(func(v LookupBudgetByResourceGroupNameResult) map[string]NotificationResponse { return v.Notifications }).(NotificationResponseMapOutput)
}

func (o LookupBudgetByResourceGroupNameResultOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetByResourceGroupNameResult) string { return v.TimeGrain }).(pulumi.StringOutput)
}

func (o LookupBudgetByResourceGroupNameResultOutput) TimePeriod() BudgetTimePeriodResponseOutput {
	return o.ApplyT(func(v LookupBudgetByResourceGroupNameResult) BudgetTimePeriodResponse { return v.TimePeriod }).(BudgetTimePeriodResponseOutput)
}

func (o LookupBudgetByResourceGroupNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetByResourceGroupNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBudgetByResourceGroupNameResultOutput{})
}

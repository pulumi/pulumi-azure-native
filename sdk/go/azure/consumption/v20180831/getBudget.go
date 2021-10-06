


package v20180831

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBudget(ctx *pulumi.Context, args *LookupBudgetArgs, opts ...pulumi.InvokeOption) (*LookupBudgetResult, error) {
	var rv LookupBudgetResult
	err := ctx.Invoke("azure-native:consumption/v20180831:getBudget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBudgetArgs struct {
	BudgetName string `pulumi:"budgetName"`
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

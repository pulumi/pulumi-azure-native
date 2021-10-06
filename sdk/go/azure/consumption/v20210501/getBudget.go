


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBudget(ctx *pulumi.Context, args *LookupBudgetArgs, opts ...pulumi.InvokeOption) (*LookupBudgetResult, error) {
	var rv LookupBudgetResult
	err := ctx.Invoke("azure-native:consumption/v20210501:getBudget", args, &rv, opts...)
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
	Filter        *BudgetFilterResponse           `pulumi:"filter"`
	ForecastSpend ForecastSpendResponse           `pulumi:"forecastSpend"`
	Id            string                          `pulumi:"id"`
	Name          string                          `pulumi:"name"`
	Notifications map[string]NotificationResponse `pulumi:"notifications"`
	TimeGrain     string                          `pulumi:"timeGrain"`
	TimePeriod    BudgetTimePeriodResponse        `pulumi:"timePeriod"`
	Type          string                          `pulumi:"type"`
}

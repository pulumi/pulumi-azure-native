


package v20181001

import (
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




package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBillingHubServiceFreeHourBalance(ctx *pulumi.Context, args *GetBillingHubServiceFreeHourBalanceArgs, opts ...pulumi.InvokeOption) (*GetBillingHubServiceFreeHourBalanceResult, error) {
	var rv GetBillingHubServiceFreeHourBalanceResult
	err := ctx.Invoke("azure-native:testbase/v20220401preview:getBillingHubServiceFreeHourBalance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBillingHubServiceFreeHourBalanceArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}

type GetBillingHubServiceFreeHourBalanceResult struct {
	IncrementEntries        []BillingHubFreeHourIncrementEntryResponse `pulumi:"incrementEntries"`
	TotalRemainingFreeHours *float64                                   `pulumi:"totalRemainingFreeHours"`
}

func GetBillingHubServiceFreeHourBalanceOutput(ctx *pulumi.Context, args GetBillingHubServiceFreeHourBalanceOutputArgs, opts ...pulumi.InvokeOption) GetBillingHubServiceFreeHourBalanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBillingHubServiceFreeHourBalanceResult, error) {
			args := v.(GetBillingHubServiceFreeHourBalanceArgs)
			r, err := GetBillingHubServiceFreeHourBalance(ctx, &args, opts...)
			var s GetBillingHubServiceFreeHourBalanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBillingHubServiceFreeHourBalanceResultOutput)
}

type GetBillingHubServiceFreeHourBalanceOutputArgs struct {
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (GetBillingHubServiceFreeHourBalanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBillingHubServiceFreeHourBalanceArgs)(nil)).Elem()
}

type GetBillingHubServiceFreeHourBalanceResultOutput struct{ *pulumi.OutputState }

func (GetBillingHubServiceFreeHourBalanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBillingHubServiceFreeHourBalanceResult)(nil)).Elem()
}

func (o GetBillingHubServiceFreeHourBalanceResultOutput) ToGetBillingHubServiceFreeHourBalanceResultOutput() GetBillingHubServiceFreeHourBalanceResultOutput {
	return o
}

func (o GetBillingHubServiceFreeHourBalanceResultOutput) ToGetBillingHubServiceFreeHourBalanceResultOutputWithContext(ctx context.Context) GetBillingHubServiceFreeHourBalanceResultOutput {
	return o
}

func (o GetBillingHubServiceFreeHourBalanceResultOutput) IncrementEntries() BillingHubFreeHourIncrementEntryResponseArrayOutput {
	return o.ApplyT(func(v GetBillingHubServiceFreeHourBalanceResult) []BillingHubFreeHourIncrementEntryResponse {
		return v.IncrementEntries
	}).(BillingHubFreeHourIncrementEntryResponseArrayOutput)
}

func (o GetBillingHubServiceFreeHourBalanceResultOutput) TotalRemainingFreeHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetBillingHubServiceFreeHourBalanceResult) *float64 { return v.TotalRemainingFreeHours }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBillingHubServiceFreeHourBalanceResultOutput{})
}




package v20190701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListApplicationAllowedUpgradePlans(ctx *pulumi.Context, args *ListApplicationAllowedUpgradePlansArgs, opts ...pulumi.InvokeOption) (*ListApplicationAllowedUpgradePlansResult, error) {
	var rv ListApplicationAllowedUpgradePlansResult
	err := ctx.Invoke("azure-native:solutions/v20190701:listApplicationAllowedUpgradePlans", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListApplicationAllowedUpgradePlansArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListApplicationAllowedUpgradePlansResult struct {
	Value []PlanResponse `pulumi:"value"`
}

func ListApplicationAllowedUpgradePlansOutput(ctx *pulumi.Context, args ListApplicationAllowedUpgradePlansOutputArgs, opts ...pulumi.InvokeOption) ListApplicationAllowedUpgradePlansResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListApplicationAllowedUpgradePlansResult, error) {
			args := v.(ListApplicationAllowedUpgradePlansArgs)
			r, err := ListApplicationAllowedUpgradePlans(ctx, &args, opts...)
			var s ListApplicationAllowedUpgradePlansResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListApplicationAllowedUpgradePlansResultOutput)
}

type ListApplicationAllowedUpgradePlansOutputArgs struct {
	ApplicationName   pulumi.StringInput `pulumi:"applicationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListApplicationAllowedUpgradePlansOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplicationAllowedUpgradePlansArgs)(nil)).Elem()
}


type ListApplicationAllowedUpgradePlansResultOutput struct{ *pulumi.OutputState }

func (ListApplicationAllowedUpgradePlansResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplicationAllowedUpgradePlansResult)(nil)).Elem()
}

func (o ListApplicationAllowedUpgradePlansResultOutput) ToListApplicationAllowedUpgradePlansResultOutput() ListApplicationAllowedUpgradePlansResultOutput {
	return o
}

func (o ListApplicationAllowedUpgradePlansResultOutput) ToListApplicationAllowedUpgradePlansResultOutputWithContext(ctx context.Context) ListApplicationAllowedUpgradePlansResultOutput {
	return o
}

func (o ListApplicationAllowedUpgradePlansResultOutput) Value() PlanResponseArrayOutput {
	return o.ApplyT(func(v ListApplicationAllowedUpgradePlansResult) []PlanResponse { return v.Value }).(PlanResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListApplicationAllowedUpgradePlansResultOutput{})
}

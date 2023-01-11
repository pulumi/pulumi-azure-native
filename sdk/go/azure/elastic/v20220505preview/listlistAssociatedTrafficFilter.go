


package v20220505preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListlistAssociatedTrafficFilter(ctx *pulumi.Context, args *ListlistAssociatedTrafficFilterArgs, opts ...pulumi.InvokeOption) (*ListlistAssociatedTrafficFilterResult, error) {
	var rv ListlistAssociatedTrafficFilterResult
	err := ctx.Invoke("azure-native:elastic/v20220505preview:listlistAssociatedTrafficFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListlistAssociatedTrafficFilterArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListlistAssociatedTrafficFilterResult struct {
	Rulesets []ElasticTrafficFilterResponse `pulumi:"rulesets"`
}

func ListlistAssociatedTrafficFilterOutput(ctx *pulumi.Context, args ListlistAssociatedTrafficFilterOutputArgs, opts ...pulumi.InvokeOption) ListlistAssociatedTrafficFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListlistAssociatedTrafficFilterResult, error) {
			args := v.(ListlistAssociatedTrafficFilterArgs)
			r, err := ListlistAssociatedTrafficFilter(ctx, &args, opts...)
			var s ListlistAssociatedTrafficFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListlistAssociatedTrafficFilterResultOutput)
}

type ListlistAssociatedTrafficFilterOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListlistAssociatedTrafficFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListlistAssociatedTrafficFilterArgs)(nil)).Elem()
}


type ListlistAssociatedTrafficFilterResultOutput struct{ *pulumi.OutputState }

func (ListlistAssociatedTrafficFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListlistAssociatedTrafficFilterResult)(nil)).Elem()
}

func (o ListlistAssociatedTrafficFilterResultOutput) ToListlistAssociatedTrafficFilterResultOutput() ListlistAssociatedTrafficFilterResultOutput {
	return o
}

func (o ListlistAssociatedTrafficFilterResultOutput) ToListlistAssociatedTrafficFilterResultOutputWithContext(ctx context.Context) ListlistAssociatedTrafficFilterResultOutput {
	return o
}

func (o ListlistAssociatedTrafficFilterResultOutput) Rulesets() ElasticTrafficFilterResponseArrayOutput {
	return o.ApplyT(func(v ListlistAssociatedTrafficFilterResult) []ElasticTrafficFilterResponse { return v.Rulesets }).(ElasticTrafficFilterResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListlistAssociatedTrafficFilterResultOutput{})
}

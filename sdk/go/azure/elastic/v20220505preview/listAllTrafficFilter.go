


package v20220505preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAllTrafficFilter(ctx *pulumi.Context, args *ListAllTrafficFilterArgs, opts ...pulumi.InvokeOption) (*ListAllTrafficFilterResult, error) {
	var rv ListAllTrafficFilterResult
	err := ctx.Invoke("azure-native:elastic/v20220505preview:listAllTrafficFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAllTrafficFilterArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListAllTrafficFilterResult struct {
	Rulesets []ElasticTrafficFilterResponse `pulumi:"rulesets"`
}

func ListAllTrafficFilterOutput(ctx *pulumi.Context, args ListAllTrafficFilterOutputArgs, opts ...pulumi.InvokeOption) ListAllTrafficFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAllTrafficFilterResult, error) {
			args := v.(ListAllTrafficFilterArgs)
			r, err := ListAllTrafficFilter(ctx, &args, opts...)
			var s ListAllTrafficFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAllTrafficFilterResultOutput)
}

type ListAllTrafficFilterOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAllTrafficFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAllTrafficFilterArgs)(nil)).Elem()
}


type ListAllTrafficFilterResultOutput struct{ *pulumi.OutputState }

func (ListAllTrafficFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAllTrafficFilterResult)(nil)).Elem()
}

func (o ListAllTrafficFilterResultOutput) ToListAllTrafficFilterResultOutput() ListAllTrafficFilterResultOutput {
	return o
}

func (o ListAllTrafficFilterResultOutput) ToListAllTrafficFilterResultOutputWithContext(ctx context.Context) ListAllTrafficFilterResultOutput {
	return o
}

func (o ListAllTrafficFilterResultOutput) Rulesets() ElasticTrafficFilterResponseArrayOutput {
	return o.ApplyT(func(v ListAllTrafficFilterResult) []ElasticTrafficFilterResponse { return v.Rulesets }).(ElasticTrafficFilterResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAllTrafficFilterResultOutput{})
}

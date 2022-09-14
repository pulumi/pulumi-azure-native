


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListLabVhds(ctx *pulumi.Context, args *ListLabVhdsArgs, opts ...pulumi.InvokeOption) (*ListLabVhdsResult, error) {
	var rv ListLabVhdsResult
	err := ctx.Invoke("azure-native:devtestlab:listLabVhds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLabVhdsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListLabVhdsResult struct {
	NextLink *string          `pulumi:"nextLink"`
	Value    []LabVhdResponse `pulumi:"value"`
}

func ListLabVhdsOutput(ctx *pulumi.Context, args ListLabVhdsOutputArgs, opts ...pulumi.InvokeOption) ListLabVhdsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListLabVhdsResult, error) {
			args := v.(ListLabVhdsArgs)
			r, err := ListLabVhds(ctx, &args, opts...)
			var s ListLabVhdsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListLabVhdsResultOutput)
}

type ListLabVhdsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListLabVhdsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLabVhdsArgs)(nil)).Elem()
}


type ListLabVhdsResultOutput struct{ *pulumi.OutputState }

func (ListLabVhdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLabVhdsResult)(nil)).Elem()
}

func (o ListLabVhdsResultOutput) ToListLabVhdsResultOutput() ListLabVhdsResultOutput {
	return o
}

func (o ListLabVhdsResultOutput) ToListLabVhdsResultOutputWithContext(ctx context.Context) ListLabVhdsResultOutput {
	return o
}

func (o ListLabVhdsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListLabVhdsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListLabVhdsResultOutput) Value() LabVhdResponseArrayOutput {
	return o.ApplyT(func(v ListLabVhdsResult) []LabVhdResponse { return v.Value }).(LabVhdResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListLabVhdsResultOutput{})
}

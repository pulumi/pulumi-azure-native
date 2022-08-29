


package labservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListGlobalUserLabs(ctx *pulumi.Context, args *ListGlobalUserLabsArgs, opts ...pulumi.InvokeOption) (*ListGlobalUserLabsResult, error) {
	var rv ListGlobalUserLabsResult
	err := ctx.Invoke("azure-native:labservices:listGlobalUserLabs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListGlobalUserLabsArgs struct {
	UserName string `pulumi:"userName"`
}


type ListGlobalUserLabsResult struct {
	Labs []LabDetailsResponse `pulumi:"labs"`
}

func ListGlobalUserLabsOutput(ctx *pulumi.Context, args ListGlobalUserLabsOutputArgs, opts ...pulumi.InvokeOption) ListGlobalUserLabsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListGlobalUserLabsResult, error) {
			args := v.(ListGlobalUserLabsArgs)
			r, err := ListGlobalUserLabs(ctx, &args, opts...)
			var s ListGlobalUserLabsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListGlobalUserLabsResultOutput)
}

type ListGlobalUserLabsOutputArgs struct {
	UserName pulumi.StringInput `pulumi:"userName"`
}

func (ListGlobalUserLabsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGlobalUserLabsArgs)(nil)).Elem()
}


type ListGlobalUserLabsResultOutput struct{ *pulumi.OutputState }

func (ListGlobalUserLabsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGlobalUserLabsResult)(nil)).Elem()
}

func (o ListGlobalUserLabsResultOutput) ToListGlobalUserLabsResultOutput() ListGlobalUserLabsResultOutput {
	return o
}

func (o ListGlobalUserLabsResultOutput) ToListGlobalUserLabsResultOutputWithContext(ctx context.Context) ListGlobalUserLabsResultOutput {
	return o
}

func (o ListGlobalUserLabsResultOutput) Labs() LabDetailsResponseArrayOutput {
	return o.ApplyT(func(v ListGlobalUserLabsResult) []LabDetailsResponse { return v.Labs }).(LabDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListGlobalUserLabsResultOutput{})
}

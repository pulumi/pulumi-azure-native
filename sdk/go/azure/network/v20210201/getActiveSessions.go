


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetActiveSessions(ctx *pulumi.Context, args *GetActiveSessionsArgs, opts ...pulumi.InvokeOption) (*GetActiveSessionsResult, error) {
	var rv GetActiveSessionsResult
	err := ctx.Invoke("azure-native:network/v20210201:getActiveSessions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetActiveSessionsArgs struct {
	BastionHostName   string `pulumi:"bastionHostName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetActiveSessionsResult struct {
	NextLink *string                        `pulumi:"nextLink"`
	Value    []BastionActiveSessionResponse `pulumi:"value"`
}

func GetActiveSessionsOutput(ctx *pulumi.Context, args GetActiveSessionsOutputArgs, opts ...pulumi.InvokeOption) GetActiveSessionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetActiveSessionsResult, error) {
			args := v.(GetActiveSessionsArgs)
			r, err := GetActiveSessions(ctx, &args, opts...)
			var s GetActiveSessionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetActiveSessionsResultOutput)
}

type GetActiveSessionsOutputArgs struct {
	BastionHostName   pulumi.StringInput `pulumi:"bastionHostName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetActiveSessionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActiveSessionsArgs)(nil)).Elem()
}


type GetActiveSessionsResultOutput struct{ *pulumi.OutputState }

func (GetActiveSessionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActiveSessionsResult)(nil)).Elem()
}

func (o GetActiveSessionsResultOutput) ToGetActiveSessionsResultOutput() GetActiveSessionsResultOutput {
	return o
}

func (o GetActiveSessionsResultOutput) ToGetActiveSessionsResultOutputWithContext(ctx context.Context) GetActiveSessionsResultOutput {
	return o
}

func (o GetActiveSessionsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetActiveSessionsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o GetActiveSessionsResultOutput) Value() BastionActiveSessionResponseArrayOutput {
	return o.ApplyT(func(v GetActiveSessionsResult) []BastionActiveSessionResponse { return v.Value }).(BastionActiveSessionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActiveSessionsResultOutput{})
}




package marketplace

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListPrivateStoreSubscriptionsContext(ctx *pulumi.Context, args *ListPrivateStoreSubscriptionsContextArgs, opts ...pulumi.InvokeOption) (*ListPrivateStoreSubscriptionsContextResult, error) {
	var rv ListPrivateStoreSubscriptionsContextResult
	err := ctx.Invoke("azure-native:marketplace:listPrivateStoreSubscriptionsContext", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPrivateStoreSubscriptionsContextArgs struct {
	PrivateStoreId string `pulumi:"privateStoreId"`
}


type ListPrivateStoreSubscriptionsContextResult struct {
	SubscriptionsIds []string `pulumi:"subscriptionsIds"`
}

func ListPrivateStoreSubscriptionsContextOutput(ctx *pulumi.Context, args ListPrivateStoreSubscriptionsContextOutputArgs, opts ...pulumi.InvokeOption) ListPrivateStoreSubscriptionsContextResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPrivateStoreSubscriptionsContextResult, error) {
			args := v.(ListPrivateStoreSubscriptionsContextArgs)
			r, err := ListPrivateStoreSubscriptionsContext(ctx, &args, opts...)
			var s ListPrivateStoreSubscriptionsContextResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPrivateStoreSubscriptionsContextResultOutput)
}

type ListPrivateStoreSubscriptionsContextOutputArgs struct {
	PrivateStoreId pulumi.StringInput `pulumi:"privateStoreId"`
}

func (ListPrivateStoreSubscriptionsContextOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPrivateStoreSubscriptionsContextArgs)(nil)).Elem()
}


type ListPrivateStoreSubscriptionsContextResultOutput struct{ *pulumi.OutputState }

func (ListPrivateStoreSubscriptionsContextResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPrivateStoreSubscriptionsContextResult)(nil)).Elem()
}

func (o ListPrivateStoreSubscriptionsContextResultOutput) ToListPrivateStoreSubscriptionsContextResultOutput() ListPrivateStoreSubscriptionsContextResultOutput {
	return o
}

func (o ListPrivateStoreSubscriptionsContextResultOutput) ToListPrivateStoreSubscriptionsContextResultOutputWithContext(ctx context.Context) ListPrivateStoreSubscriptionsContextResultOutput {
	return o
}

func (o ListPrivateStoreSubscriptionsContextResultOutput) SubscriptionsIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListPrivateStoreSubscriptionsContextResult) []string { return v.SubscriptionsIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPrivateStoreSubscriptionsContextResultOutput{})
}




package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListShareSubscriptionSynchronizations(ctx *pulumi.Context, args *ListShareSubscriptionSynchronizationsArgs, opts ...pulumi.InvokeOption) (*ListShareSubscriptionSynchronizationsResult, error) {
	var rv ListShareSubscriptionSynchronizationsResult
	err := ctx.Invoke("azure-native:datashare/v20191101:listShareSubscriptionSynchronizations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListShareSubscriptionSynchronizationsArgs struct {
	AccountName           string  `pulumi:"accountName"`
	Filter                *string `pulumi:"filter"`
	Orderby               *string `pulumi:"orderby"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	SkipToken             *string `pulumi:"skipToken"`
}


type ListShareSubscriptionSynchronizationsResult struct {
	NextLink *string                                    `pulumi:"nextLink"`
	Value    []ShareSubscriptionSynchronizationResponse `pulumi:"value"`
}

func ListShareSubscriptionSynchronizationsOutput(ctx *pulumi.Context, args ListShareSubscriptionSynchronizationsOutputArgs, opts ...pulumi.InvokeOption) ListShareSubscriptionSynchronizationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListShareSubscriptionSynchronizationsResult, error) {
			args := v.(ListShareSubscriptionSynchronizationsArgs)
			r, err := ListShareSubscriptionSynchronizations(ctx, &args, opts...)
			var s ListShareSubscriptionSynchronizationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListShareSubscriptionSynchronizationsResultOutput)
}

type ListShareSubscriptionSynchronizationsOutputArgs struct {
	AccountName           pulumi.StringInput    `pulumi:"accountName"`
	Filter                pulumi.StringPtrInput `pulumi:"filter"`
	Orderby               pulumi.StringPtrInput `pulumi:"orderby"`
	ResourceGroupName     pulumi.StringInput    `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput    `pulumi:"shareSubscriptionName"`
	SkipToken             pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListShareSubscriptionSynchronizationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListShareSubscriptionSynchronizationsArgs)(nil)).Elem()
}


type ListShareSubscriptionSynchronizationsResultOutput struct{ *pulumi.OutputState }

func (ListShareSubscriptionSynchronizationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListShareSubscriptionSynchronizationsResult)(nil)).Elem()
}

func (o ListShareSubscriptionSynchronizationsResultOutput) ToListShareSubscriptionSynchronizationsResultOutput() ListShareSubscriptionSynchronizationsResultOutput {
	return o
}

func (o ListShareSubscriptionSynchronizationsResultOutput) ToListShareSubscriptionSynchronizationsResultOutputWithContext(ctx context.Context) ListShareSubscriptionSynchronizationsResultOutput {
	return o
}

func (o ListShareSubscriptionSynchronizationsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListShareSubscriptionSynchronizationsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListShareSubscriptionSynchronizationsResultOutput) Value() ShareSubscriptionSynchronizationResponseArrayOutput {
	return o.ApplyT(func(v ListShareSubscriptionSynchronizationsResult) []ShareSubscriptionSynchronizationResponse {
		return v.Value
	}).(ShareSubscriptionSynchronizationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListShareSubscriptionSynchronizationsResultOutput{})
}
